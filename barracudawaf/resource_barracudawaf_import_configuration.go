package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFImportConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFImportConfigurationCreate,
		Read:   resourceCudaWAFImportConfigurationRead,
		Update: resourceCudaWAFImportConfigurationUpdate,
		Delete: resourceCudaWAFImportConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"operation":         {Type: schema.TypeString, Optional: true},
			"json_file_content": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadImportConfiguration(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"operation":         d.Get("operation").(string),
		"json-file-content": d.Get("json_file_content").(string),
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

func resourceCudaWAFImportConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/import-configuration"
	resourceCreateResponseError := makeRestAPIPayloadImportConfiguration(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFImportConfigurationRead(d, m)
}

func resourceCudaWAFImportConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFImportConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/import-configuration/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadImportConfiguration(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFImportConfigurationRead(d, m)
}

func resourceCudaWAFImportConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/import-configuration/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadImportConfiguration(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
