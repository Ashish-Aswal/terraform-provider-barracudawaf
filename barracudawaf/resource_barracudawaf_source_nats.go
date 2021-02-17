package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSourceNats() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSourceNatsCreate,
		Read:   resourceCudaWAFSourceNatsRead,
		Update: resourceCudaWAFSourceNatsUpdate,
		Delete: resourceCudaWAFSourceNatsDelete,

		Schema: map[string]*schema.Schema{
			"outgoing_interface":  {Type: schema.TypeString, Required: true},
			"post_source_address": {Type: schema.TypeString, Required: true},
			"protocol":            {Type: schema.TypeString, Required: true},
			"pre_source_address":  {Type: schema.TypeString, Required: true},
			"pre_source_netmask":  {Type: schema.TypeString, Required: true},
			"destination_port":    {Type: schema.TypeString, Optional: true},
			"comments":            {Type: schema.TypeString, Optional: true},
			"vsite":               {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadSourceNats(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"outgoing-interface":  d.Get("outgoing_interface").(string),
		"post-source-address": d.Get("post_source_address").(string),
		"protocol":            d.Get("protocol").(string),
		"pre-source-address":  d.Get("pre_source_address").(string),
		"pre-source-netmask":  d.Get("pre_source_netmask").(string),
		"destination-port":    d.Get("destination_port").(string),
		"comments":            d.Get("comments").(string),
		"vsite":               d.Get("vsite").(string),
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

func resourceCudaWAFSourceNatsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/source-nats"
	resourceCreateResponseError := makeRestAPIPayloadSourceNats(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSourceNatsRead(d, m)
}

func resourceCudaWAFSourceNatsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSourceNatsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/source-nats/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSourceNats(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSourceNatsRead(d, m)
}

func resourceCudaWAFSourceNatsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/source-nats/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSourceNats(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
