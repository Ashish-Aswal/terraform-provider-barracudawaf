package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFLocalUsers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFLocalUsersCreate,
		Read:   resourceCudaWAFLocalUsersRead,
		Update: resourceCudaWAFLocalUsersUpdate,
		Delete: resourceCudaWAFLocalUsersDelete,

		Schema: map[string]*schema.Schema{
			"user_groups": {Type: schema.TypeString, Optional: true},
			"name":        {Type: schema.TypeString, Optional: true},
			"password":    {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadLocalUsers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"user-groups": d.Get("user_groups").(string),
		"name":        d.Get("name").(string),
		"password":    d.Get("password").(string),
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

func resourceCudaWAFLocalUsersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-users"
	resourceCreateResponseError := makeRestAPIPayloadLocalUsers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFLocalUsersRead(d, m)
}

func resourceCudaWAFLocalUsersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFLocalUsersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-users/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadLocalUsers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFLocalUsersRead(d, m)
}

func resourceCudaWAFLocalUsersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-users/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadLocalUsers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
