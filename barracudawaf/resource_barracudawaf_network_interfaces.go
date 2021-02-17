package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNetworkInterfacesCreate,
		Read:   resourceCudaWAFNetworkInterfacesRead,
		Update: resourceCudaWAFNetworkInterfacesUpdate,
		Delete: resourceCudaWAFNetworkInterfacesDelete,

		Schema: map[string]*schema.Schema{
			"name":                    {Type: schema.TypeString, Required: true},
			"duplexity":               {Type: schema.TypeString, Required: true},
			"auto_negotiation_status": {Type: schema.TypeString, Optional: true},
			"speed":                   {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadNetworkInterfaces(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                    d.Get("name").(string),
		"duplexity":               d.Get("duplexity").(string),
		"auto-negotiation-status": d.Get("auto_negotiation_status").(string),
		"speed":                   d.Get("speed").(string),
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

func resourceCudaWAFNetworkInterfacesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/network-interfaces"
	resourceCreateResponseError := makeRestAPIPayloadNetworkInterfaces(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFNetworkInterfacesRead(d, m)
}

func resourceCudaWAFNetworkInterfacesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNetworkInterfacesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/network-interfaces/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadNetworkInterfaces(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFNetworkInterfacesRead(d, m)
}

func resourceCudaWAFNetworkInterfacesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/network-interfaces/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadNetworkInterfaces(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
