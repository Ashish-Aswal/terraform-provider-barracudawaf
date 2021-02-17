package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFProtectedDataTypes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFProtectedDataTypesCreate,
		Read:   resourceCudaWAFProtectedDataTypesRead,
		Update: resourceCudaWAFProtectedDataTypesUpdate,
		Delete: resourceCudaWAFProtectedDataTypesDelete,

		Schema: map[string]*schema.Schema{
			"action":                      {Type: schema.TypeString, Optional: true},
			"initial_characters_to_keep":  {Type: schema.TypeString, Optional: true},
			"trailing_characters_to_keep": {Type: schema.TypeString, Optional: true},
			"name":                        {Type: schema.TypeString, Required: true},
			"custom_identity_theft_type":  {Type: schema.TypeString, Optional: true},
			"enable":                      {Type: schema.TypeString, Optional: true},
			"identity_theft_type":         {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadProtectedDataTypes(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"action":                      d.Get("action").(string),
		"initial-characters-to-keep":  d.Get("initial_characters_to_keep").(string),
		"trailing-characters-to-keep": d.Get("trailing_characters_to_keep").(string),
		"name":                        d.Get("name").(string),
		"custom-identity-theft-type":  d.Get("custom_identity_theft_type").(string),
		"enable":                      d.Get("enable").(string),
		"identity-theft-type":         d.Get("identity_theft_type").(string),
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

func resourceCudaWAFProtectedDataTypesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/protected-data-types"
	resourceCreateResponseError := makeRestAPIPayloadProtectedDataTypes(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFProtectedDataTypesRead(d, m)
}

func resourceCudaWAFProtectedDataTypesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFProtectedDataTypesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/protected-data-types/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadProtectedDataTypes(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFProtectedDataTypesRead(d, m)
}

func resourceCudaWAFProtectedDataTypesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/protected-data-types/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadProtectedDataTypes(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
