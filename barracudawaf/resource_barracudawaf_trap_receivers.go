package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFTrapReceivers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFTrapReceiversCreate,
		Read:   resourceCudaWAFTrapReceiversRead,
		Update: resourceCudaWAFTrapReceiversUpdate,
		Delete: resourceCudaWAFTrapReceiversDelete,

		Schema: map[string]*schema.Schema{
			"community_string": {Type: schema.TypeString, Required: true},
			"ip_address":       {Type: schema.TypeString, Required: true},
			"port":             {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadTrapReceivers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"community-string": d.Get("community_string").(string),
		"ip-address":       d.Get("ip_address").(string),
		"port":             d.Get("port").(string),
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

func resourceCudaWAFTrapReceiversCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trap-receivers"
	resourceCreateResponseError := makeRestAPIPayloadTrapReceivers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFTrapReceiversRead(d, m)
}

func resourceCudaWAFTrapReceiversRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFTrapReceiversUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trap-receivers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadTrapReceivers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFTrapReceiversRead(d, m)
}

func resourceCudaWAFTrapReceiversDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trap-receivers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadTrapReceivers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
