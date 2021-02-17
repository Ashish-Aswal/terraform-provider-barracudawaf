package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNtpServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNtpServersCreate,
		Read:   resourceCudaWAFNtpServersRead,
		Update: resourceCudaWAFNtpServersUpdate,
		Delete: resourceCudaWAFNtpServersDelete,

		Schema: map[string]*schema.Schema{
			"description": {Type: schema.TypeString, Optional: true},
			"ip_address":  {Type: schema.TypeString, Required: true},
			"name":        {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadNtpServers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"description": d.Get("description").(string),
		"ip-address":  d.Get("ip_address").(string),
		"name":        d.Get("name").(string),
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

func resourceCudaWAFNtpServersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/ntp-servers"
	resourceCreateResponseError := makeRestAPIPayloadNtpServers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFNtpServersRead(d, m)
}

func resourceCudaWAFNtpServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNtpServersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/ntp-servers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadNtpServers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFNtpServersRead(d, m)
}

func resourceCudaWAFNtpServersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/ntp-servers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadNtpServers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
