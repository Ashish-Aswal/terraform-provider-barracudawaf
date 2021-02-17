package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadClientCertificateCrls(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

func resourceCudaWAFClientCertificateCrlsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/client-certificate-crls"
	resourceCreateResponseError := makeRestAPIPayloadClientCertificateCrls(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFClientCertificateCrlsRead(d, m)
}

func resourceCudaWAFClientCertificateCrlsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFClientCertificateCrlsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/client-certificate-crls/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadClientCertificateCrls(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFClientCertificateCrlsRead(d, m)
}

func resourceCudaWAFClientCertificateCrlsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/client-certificate-crls/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadClientCertificateCrls(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
