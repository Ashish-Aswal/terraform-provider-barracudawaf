package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSecurityPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSecurityPoliciesCreate,
		Read:   resourceCudaWAFSecurityPoliciesRead,
		Update: resourceCudaWAFSecurityPoliciesUpdate,
		Delete: resourceCudaWAFSecurityPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"based_on": {Type: schema.TypeString, Optional: true},
			"name":     {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadSecurityPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"based-on": d.Get("based_on").(string),
		"name":     d.Get("name").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"based-on"}
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

func resourceCudaWAFSecurityPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies"
	resourceCreateResponseError := makeRestAPIPayloadSecurityPolicies(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSecurityPoliciesRead(d, m)
}

func resourceCudaWAFSecurityPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSecurityPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSecurityPolicies(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSecurityPoliciesRead(d, m)
}

func resourceCudaWAFSecurityPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSecurityPolicies(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
