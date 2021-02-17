package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNodes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNodesCreate,
		Read:   resourceCudaWAFNodesRead,
		Update: resourceCudaWAFNodesUpdate,
		Delete: resourceCudaWAFNodesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {Type: schema.TypeString, Required: true},
			"mode":       {Type: schema.TypeString, Required: true},
			"serial":     {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadNodes(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"ip-address": d.Get("ip_address").(string),
		"mode":       d.Get("mode").(string),
		"serial":     d.Get("serial").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"mode", "serial"}
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

func resourceCudaWAFNodesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/nodes"
	resourceCreateResponseError := makeRestAPIPayloadNodes(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFNodesRead(d, m)
}

func resourceCudaWAFNodesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNodesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/nodes/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadNodes(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFNodesRead(d, m)
}

func resourceCudaWAFNodesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/nodes/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadNodes(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
