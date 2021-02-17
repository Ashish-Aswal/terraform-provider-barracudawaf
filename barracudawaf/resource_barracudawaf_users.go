package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUsers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUsersCreate,
		Read:   resourceCudaWAFUsersRead,
		Update: resourceCudaWAFUsersUpdate,
		Delete: resourceCudaWAFUsersDelete,

		Schema: map[string]*schema.Schema{
			"email_address":     {Type: schema.TypeString, Required: true},
			"name":              {Type: schema.TypeString, Required: true},
			"password":          {Type: schema.TypeString, Optional: true},
			"re_enter_password": {Type: schema.TypeString, Optional: true},
			"role":              {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadUsers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"email-address":     d.Get("email_address").(string),
		"name":              d.Get("name").(string),
		"password":          d.Get("password").(string),
		"re-enter-password": d.Get("re_enter_password").(string),
		"role":              d.Get("role").(string),
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

func resourceCudaWAFUsersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/users"
	resourceCreateResponseError := makeRestAPIPayloadUsers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFUsersRead(d, m)
}

func resourceCudaWAFUsersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUsersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/users/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadUsers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFUsersRead(d, m)
}

func resourceCudaWAFUsersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/users/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadUsers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
