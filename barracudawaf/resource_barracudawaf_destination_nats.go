package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFDestinationNats() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFDestinationNatsCreate,
		Read:   resourceCudaWAFDestinationNatsRead,
		Update: resourceCudaWAFDestinationNatsUpdate,
		Delete: resourceCudaWAFDestinationNatsDelete,

		Schema: map[string]*schema.Schema{
			"comments":                 {Type: schema.TypeString, Optional: true},
			"vsite":                    {Type: schema.TypeString, Required: true},
			"incoming_interface":       {Type: schema.TypeString, Required: true},
			"post_destination_address": {Type: schema.TypeString, Required: true},
			"pre_destination_address":  {Type: schema.TypeString, Required: true},
			"pre_destination_netmask":  {Type: schema.TypeString, Required: true},
			"pre_destination_port":     {Type: schema.TypeString, Optional: true},
			"protocol":                 {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadDestinationNats(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":                 d.Get("comments").(string),
		"vsite":                    d.Get("vsite").(string),
		"incoming-interface":       d.Get("incoming_interface").(string),
		"post-destination-address": d.Get("post_destination_address").(string),
		"pre-destination-address":  d.Get("pre_destination_address").(string),
		"pre-destination-netmask":  d.Get("pre_destination_netmask").(string),
		"pre-destination-port":     d.Get("pre_destination_port").(string),
		"protocol":                 d.Get("protocol").(string),
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

func resourceCudaWAFDestinationNatsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/destination-nats"
	resourceCreateResponseError := makeRestAPIPayloadDestinationNats(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFDestinationNatsRead(d, m)
}

func resourceCudaWAFDestinationNatsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFDestinationNatsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/destination-nats/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadDestinationNats(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFDestinationNatsRead(d, m)
}

func resourceCudaWAFDestinationNatsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/destination-nats/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadDestinationNats(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
