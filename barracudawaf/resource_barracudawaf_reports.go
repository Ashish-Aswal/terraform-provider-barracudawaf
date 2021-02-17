package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadReports(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{}
		for item := range updatePayloadExceptions {
			delete(resourcePayload, updatePayloadExceptions[item])
		}
	}

	//sanitise the resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	//resourceUpdateData : cudaWAF reource URI update data
	resourceUpdateData := map[string]interface{}{
		"endpoint":  resourceEndpoint,
		"payload":   resourcePayload,
		"operation": resourceOperation,
		"name":      d.Get("name").(string),
	}

	//updateCudaWAFResourceObject : update cudaWAF resource object
	resourceUpdateStatus, resourceUpdateResponseBody := updateCudaWAFResourceObject(
		resourceUpdateData,
	)

	if resourceUpdateStatus == 200 || resourceUpdateStatus == 201 {
		if resourceOperation != "DELETE" {
			d.SetId(resourceUpdateResponseBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", resourceUpdateResponseBody["msg"])
	}

	return nil
}

func resourceCudaWAFReportsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/reports"
	resourceCreateResponseError := makeRestAPIPayloadReports(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFReportsRead(d, m)
}

func resourceCudaWAFReportsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFReportsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/reports/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadReports(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFReportsRead(d, m)
}

func resourceCudaWAFReportsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/reports/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadReports(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
