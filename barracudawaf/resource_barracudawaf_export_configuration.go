package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFExportConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFExportConfigurationCreate,
		Read:   resourceCudaWAFExportConfigurationRead,
		Update: resourceCudaWAFExportConfigurationUpdate,
		Delete: resourceCudaWAFExportConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"backup_type":    {Type: schema.TypeString, Optional: true},
			"name":           {Type: schema.TypeString, Optional: true},
			"destination":    {Type: schema.TypeString, Optional: true},
			"day_of_week":    {Type: schema.TypeString, Optional: true},
			"hour_of_day":    {Type: schema.TypeString, Optional: true},
			"minute_of_hour": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadExportConfiguration(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"backup-type":    d.Get("backup_type").(string),
		"name":           d.Get("name").(string),
		"destination":    d.Get("destination").(string),
		"day-of-week":    d.Get("day_of_week").(string),
		"hour-of-day":    d.Get("hour_of_day").(string),
		"minute-of-hour": d.Get("minute_of_hour").(string),
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

func resourceCudaWAFExportConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/export-configuration"
	resourceCreateResponseError := makeRestAPIPayloadExportConfiguration(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFExportConfigurationRead(d, m)
}

func resourceCudaWAFExportConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFExportConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/export-configuration/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadExportConfiguration(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFExportConfigurationRead(d, m)
}

func resourceCudaWAFExportConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/export-configuration/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadExportConfiguration(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
