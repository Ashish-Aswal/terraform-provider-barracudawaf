package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFLocalGroups() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFLocalGroupsCreate,
		Read:   resourceCudaWAFLocalGroupsRead,
		Update: resourceCudaWAFLocalGroupsUpdate,
		Delete: resourceCudaWAFLocalGroupsDelete,

		Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}},
	}
}

func makeRestAPIPayloadLocalGroups(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{"name": d.Get("name").(string)}

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

func resourceCudaWAFLocalGroupsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-groups"
	resourceCreateResponseError := makeRestAPIPayloadLocalGroups(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFLocalGroupsRead(d, m)
}

func resourceCudaWAFLocalGroupsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFLocalGroupsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-groups/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadLocalGroups(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFLocalGroupsRead(d, m)
}

func resourceCudaWAFLocalGroupsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-groups/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadLocalGroups(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
