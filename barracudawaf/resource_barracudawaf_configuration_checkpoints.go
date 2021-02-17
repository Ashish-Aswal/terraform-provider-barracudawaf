package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFConfigurationCheckpoints() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFConfigurationCheckpointsCreate,
		Read:   resourceCudaWAFConfigurationCheckpointsRead,
		Update: resourceCudaWAFConfigurationCheckpointsUpdate,
		Delete: resourceCudaWAFConfigurationCheckpointsDelete,

		Schema: map[string]*schema.Schema{
			"name":    {Type: schema.TypeString, Required: true},
			"comment": {Type: schema.TypeString, Optional: true},
			"date":    {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadConfigurationCheckpoints(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":    d.Get("name").(string),
		"comment": d.Get("comment").(string),
		"date":    d.Get("date").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"date"}
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

func resourceCudaWAFConfigurationCheckpointsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/configuration-checkpoints"
	resourceCreateResponseError := makeRestAPIPayloadConfigurationCheckpoints(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFConfigurationCheckpointsRead(d, m)
}

func resourceCudaWAFConfigurationCheckpointsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFConfigurationCheckpointsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/configuration-checkpoints/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadConfigurationCheckpoints(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFConfigurationCheckpointsRead(d, m)
}

func resourceCudaWAFConfigurationCheckpointsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/configuration-checkpoints/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadConfigurationCheckpoints(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
