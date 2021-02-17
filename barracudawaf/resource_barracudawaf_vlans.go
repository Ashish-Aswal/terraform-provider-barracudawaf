package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFVlans() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFVlansCreate,
		Read:   resourceCudaWAFVlansRead,
		Update: resourceCudaWAFVlansUpdate,
		Delete: resourceCudaWAFVlansDelete,

		Schema: map[string]*schema.Schema{
			"comments":  {Type: schema.TypeString, Optional: true},
			"vlan_id":   {Type: schema.TypeString, Required: true},
			"interface": {Type: schema.TypeString, Required: true},
			"name":      {Type: schema.TypeString, Required: true},
			"vsite":     {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadVlans(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":  d.Get("comments").(string),
		"vlan-id":   d.Get("vlan_id").(string),
		"interface": d.Get("interface").(string),
		"name":      d.Get("name").(string),
		"vsite":     d.Get("vsite").(string),
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

func resourceCudaWAFVlansCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vlans"
	resourceCreateResponseError := makeRestAPIPayloadVlans(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFVlansRead(d, m)
}

func resourceCudaWAFVlansRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFVlansUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vlans/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadVlans(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFVlansRead(d, m)
}

func resourceCudaWAFVlansDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vlans/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadVlans(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
