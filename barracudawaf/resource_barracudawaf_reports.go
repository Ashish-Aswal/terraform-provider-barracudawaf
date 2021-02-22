package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFReports() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFReportsCreate,
		Read:   resourceCudaWAFReportsRead,
		Update: resourceCudaWAFReportsUpdate,
		Delete: resourceCudaWAFReportsDelete,

		Schema: map[string]*schema.Schema{
			"report_format":    {Type: schema.TypeString, Optional: true},
			"frequency":        {Type: schema.TypeString, Optional: true},
			"ftp_directory":    {Type: schema.TypeString, Optional: true},
			"ftp_ip_address":   {Type: schema.TypeString, Optional: true},
			"ftp_password":     {Type: schema.TypeString, Optional: true},
			"ftp_port":         {Type: schema.TypeString, Optional: true},
			"ftp_username":     {Type: schema.TypeString, Optional: true},
			"email_id":         {Type: schema.TypeString, Optional: true},
			"name":             {Type: schema.TypeString, Required: true},
			"report_types":     {Type: schema.TypeString, Required: true},
			"delivery_options": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFReportsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/reports"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFReportsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFReportsRead(d, m)
}

func resourceCudaWAFReportsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFReportsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/reports/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFReportsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFReportsRead(d, m)
}

func resourceCudaWAFReportsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/reports/"
	request := &APIRequest{
		Method: "delete",
		URL:    resourceEndpoint,
	}

	err := client.DeleteBarracudaWAFResource(name, request)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func hydrateBarracudaWAFReportsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"report-format":    d.Get("report_format").(string),
		"frequency":        d.Get("frequency").(string),
		"ftp-directory":    d.Get("ftp_directory").(string),
		"ftp-ip-address":   d.Get("ftp_ip_address").(string),
		"ftp-password":     d.Get("ftp_password").(string),
		"ftp-port":         d.Get("ftp_port").(string),
		"ftp-username":     d.Get("ftp_username").(string),
		"email-id":         d.Get("email_id").(string),
		"name":             d.Get("name").(string),
		"report-types":     d.Get("report_types").(string),
		"delivery-options": d.Get("delivery_options").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{}
		for item := range updatePayloadExceptions {
			delete(resourcePayload, updatePayloadExceptions[item])
		}
	}

	// remove empty parameters from resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	return &APIRequest{
		URL:  endpoint,
		Body: resourcePayload,
	}
}
