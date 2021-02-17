package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFWhitelistedBots() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFWhitelistedBotsCreate,
		Read:   resourceCudaWAFWhitelistedBotsRead,
		Update: resourceCudaWAFWhitelistedBotsUpdate,
		Delete: resourceCudaWAFWhitelistedBotsDelete,

		Schema: map[string]*schema.Schema{
			"host":       {Type: schema.TypeString, Optional: true},
			"identifier": {Type: schema.TypeString, Optional: true},
			"ip_address": {Type: schema.TypeString, Optional: true},
			"name":       {Type: schema.TypeString, Required: true},
			"user_agent": {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadWhitelistedBots(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"host":       d.Get("host").(string),
		"identifier": d.Get("identifier").(string),
		"ip-address": d.Get("ip_address").(string),
		"name":       d.Get("name").(string),
		"user-agent": d.Get("user_agent").(string),
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

func resourceCudaWAFWhitelistedBotsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/whitelisted-bots"
	resourceCreateResponseError := makeRestAPIPayloadWhitelistedBots(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFWhitelistedBotsRead(d, m)
}

func resourceCudaWAFWhitelistedBotsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFWhitelistedBotsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/whitelisted-bots/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadWhitelistedBots(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFWhitelistedBotsRead(d, m)
}

func resourceCudaWAFWhitelistedBotsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/whitelisted-bots/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadWhitelistedBots(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
