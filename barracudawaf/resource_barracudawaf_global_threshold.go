package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGlobalThreshold() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGlobalThresholdCreate,
		Read:   resourceCudaWAFGlobalThresholdRead,
		Update: resourceCudaWAFGlobalThresholdUpdate,
		Delete: resourceCudaWAFGlobalThresholdDelete,

		Schema: map[string]*schema.Schema{"threshold": {Type: schema.TypeString, Optional: true}},
	}
}

func makeRestAPIPayloadGlobalThreshold(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{"threshold": d.Get("threshold").(string)}

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

func resourceCudaWAFGlobalThresholdCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/global-threshold"
	resourceCreateResponseError := makeRestAPIPayloadGlobalThreshold(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFGlobalThresholdRead(d, m)
}

func resourceCudaWAFGlobalThresholdRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGlobalThresholdUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/global-threshold/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadGlobalThreshold(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFGlobalThresholdRead(d, m)
}

func resourceCudaWAFGlobalThresholdDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/global-threshold/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadGlobalThreshold(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
