package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFVsites() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFVsitesCreate,
		Read:   resourceCudaWAFVsitesRead,
		Update: resourceCudaWAFVsitesUpdate,
		Delete: resourceCudaWAFVsitesDelete,

		Schema: map[string]*schema.Schema{
			"comments":  {Type: schema.TypeString, Optional: true},
			"active_on": {Type: schema.TypeString, Optional: true},
			"interface": {Type: schema.TypeString, Required: true},
			"name":      {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadVsites(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":  d.Get("comments").(string),
		"active-on": d.Get("active_on").(string),
		"interface": d.Get("interface").(string),
		"name":      d.Get("name").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"interface", "name"}
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

func resourceCudaWAFVsitesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vsites"
	resourceCreateResponseError := makeRestAPIPayloadVsites(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFVsitesRead(d, m)
}

func resourceCudaWAFVsitesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFVsitesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vsites/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadVsites(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFVsitesRead(d, m)
}

func resourceCudaWAFVsitesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vsites/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadVsites(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
