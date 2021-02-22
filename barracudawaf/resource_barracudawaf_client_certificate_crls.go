package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFClientCertificateCrls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFClientCertificateCrlsCreate,
		Read:   resourceCudaWAFClientCertificateCrlsRead,
		Update: resourceCudaWAFClientCertificateCrlsUpdate,
		Delete: resourceCudaWAFClientCertificateCrlsDelete,

		Schema: map[string]*schema.Schema{
			"auto_update_type":  {Type: schema.TypeString, Optional: true},
			"date_of_month":     {Type: schema.TypeString, Optional: true},
			"day_of_week":       {Type: schema.TypeString, Optional: true},
			"time_of_day":       {Type: schema.TypeString, Optional: true},
			"auto_update":       {Type: schema.TypeString, Optional: true},
			"name":              {Type: schema.TypeString, Required: true},
			"number_of_retries": {Type: schema.TypeString, Optional: true},
			"url":               {Type: schema.TypeString, Required: true},
			"enable":            {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFClientCertificateCrlsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/client-certificate-crls"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFClientCertificateCrlsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFClientCertificateCrlsRead(d, m)
}

func resourceCudaWAFClientCertificateCrlsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFClientCertificateCrlsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/client-certificate-crls/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFClientCertificateCrlsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFClientCertificateCrlsRead(d, m)
}

func resourceCudaWAFClientCertificateCrlsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/client-certificate-crls/"
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

func hydrateBarracudaWAFClientCertificateCrlsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"auto-update-type":  d.Get("auto_update_type").(string),
		"date-of-month":     d.Get("date_of_month").(string),
		"day-of-week":       d.Get("day_of_week").(string),
		"time-of-day":       d.Get("time_of_day").(string),
		"auto-update":       d.Get("auto_update").(string),
		"name":              d.Get("name").(string),
		"number-of-retries": d.Get("number_of_retries").(string),
		"url":               d.Get("url").(string),
		"enable":            d.Get("enable").(string),
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
