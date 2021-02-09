package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSecurityPolicyCreate,
		Read:   resourceCudaWAFSecurityPolicyRead,
		Update: resourceCudaWAFSecurityPolicyUpdate,
		Delete: resourceCudaWAFSecurityPolicyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"based_on": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadSecurityPolicy(d *schema.ResourceData, m interface{}, resourceOperation string, resourceEndpoint string) error {

	//build Payload for the resource
	resourcePayload := map[string]string{
		"name":     d.Get("name").(string),
		"based-on": d.Get("based_on").(string),
	}

	//sanitise the payload, removing empty keys
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	resourceUpdateData := map[string]interface{}{
		"endpoint":  resourceEndpoint,
		"payload":   resourcePayload,
		"operation": resourceOperation,
		"name":      d.Get("name").(string),
	}

	resourceUpdateStatus, resourceUpdateResponseBody := updateCudaWAFResourceObject(resourceUpdateData)

	if resourceUpdateStatus == 200 || resourceUpdateStatus == 201 {
		if resourceOperation != "DELETE" {
			d.SetId(resourceUpdateResponseBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", resourceUpdateResponseBody["msg"])
	}

	return nil
}

func resourceCudaWAFSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := "restapi/v3/security-policies"
	resourceUpdateResponseError := makeRestAPIPayloadSecurityPolicy(d, m, "POST", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSecurityPolicyRead(d, m)
}

func resourceCudaWAFSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSecurityPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCudaWAFSecurityPolicyRead(d, m)
}

func resourceCudaWAFSecurityPolicyDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceEndpoint := "restapi/v3/security-policies/" + name
	resourceDeleteResponseError := makeRestAPIPayloadSecurityPolicy(d, m, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
