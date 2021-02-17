package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFJsonSecurityPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFJsonSecurityPoliciesCreate,
		Read:   resourceCudaWAFJsonSecurityPoliciesRead,
		Update: resourceCudaWAFJsonSecurityPoliciesUpdate,
		Delete: resourceCudaWAFJsonSecurityPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"name":               {Type: schema.TypeString, Required: true},
			"max_array_elements": {Type: schema.TypeString, Optional: true},
			"max_siblings":       {Type: schema.TypeString, Optional: true},
			"max_keys":           {Type: schema.TypeString, Required: true},
			"max_key_length":     {Type: schema.TypeString, Required: true},
			"max_number_value":   {Type: schema.TypeString, Optional: true},
			"max_object_depth":   {Type: schema.TypeString, Optional: true},
			"max_value_length":   {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadJsonSecurityPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":               d.Get("name").(string),
		"max-array-elements": d.Get("max_array_elements").(string),
		"max-siblings":       d.Get("max_siblings").(string),
		"max-keys":           d.Get("max_keys").(string),
		"max-key-length":     d.Get("max_key_length").(string),
		"max-number-value":   d.Get("max_number_value").(string),
		"max-object-depth":   d.Get("max_object_depth").(string),
		"max-value-length":   d.Get("max_value_length").(string),
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

func resourceCudaWAFJsonSecurityPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/json-security-policies"
	resourceCreateResponseError := makeRestAPIPayloadJsonSecurityPolicies(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFJsonSecurityPoliciesRead(d, m)
}

func resourceCudaWAFJsonSecurityPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFJsonSecurityPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/json-security-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadJsonSecurityPolicies(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFJsonSecurityPoliciesRead(d, m)
}

func resourceCudaWAFJsonSecurityPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/json-security-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadJsonSecurityPolicies(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
