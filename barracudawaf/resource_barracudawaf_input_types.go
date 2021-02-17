package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFInputTypes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFInputTypesCreate,
		Read:   resourceCudaWAFInputTypesRead,
		Update: resourceCudaWAFInputTypesUpdate,
		Delete: resourceCudaWAFInputTypesDelete,

		Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}},
	}
}

func makeRestAPIPayloadInputTypes(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{"name": d.Get("name").(string)}

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

func resourceCudaWAFInputTypesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/input-types"
	resourceCreateResponseError := makeRestAPIPayloadInputTypes(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFInputTypesRead(d, m)
}

func resourceCudaWAFInputTypesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFInputTypesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/input-types/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadInputTypes(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFInputTypesRead(d, m)
}

func resourceCudaWAFInputTypesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/input-types/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadInputTypes(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
