package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFImportOpenapi() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFImportOpenapiCreate,
		Read:   resourceCudaWAFImportOpenapiRead,
		Update: resourceCudaWAFImportOpenapiUpdate,
		Delete: resourceCudaWAFImportOpenapiDelete,

		Schema: map[string]*schema.Schema{
			"operation":           {Type: schema.TypeString, Optional: true},
			"import_file_content": {Type: schema.TypeString, Optional: true},
			"advanced_config":     {Type: schema.TypeString, Optional: true},
			"service_name":        {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadImportOpenapi(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"operation":           d.Get("operation").(string),
		"import-file-content": d.Get("import_file_content").(string),
		"advanced-config":     d.Get("advanced_config").(string),
		"service-name":        d.Get("service_name").(string),
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

func resourceCudaWAFImportOpenapiCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/import-openapi"
	resourceCreateResponseError := makeRestAPIPayloadImportOpenapi(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFImportOpenapiRead(d, m)
}

func resourceCudaWAFImportOpenapiRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFImportOpenapiUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/import-openapi/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadImportOpenapi(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFImportOpenapiRead(d, m)
}

func resourceCudaWAFImportOpenapiDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/import-openapi/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadImportOpenapi(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
