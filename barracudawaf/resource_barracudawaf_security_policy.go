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

func makeRestAPIPayloadSecurityPolicy(d *schema.ResourceData, m interface{}, oper string, endpoint string) error {
	payload := map[string]string{
		"name":     d.Get("name").(string),
		"based-on": d.Get("based_on").(string),
	}

	for key, value := range payload {
		if len(value) > 0 {
			continue
		} else {
			delete(payload, key)
		}
	}

	callData := map[string]interface{}{
		"endpoint":  endpoint,
		"payload":   payload,
		"operation": oper,
		"name":      d.Get("name").(string),
	}

	callStatus, callRespBody := doRestAPICall(callData)
	if callStatus == 200 || callStatus == 201 {
		if oper != "DELETE" {
			d.SetId(callRespBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", callRespBody["msg"])
	}

	return nil
}

func resourceCudaWAFSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {
	endpoint := "restapi/v3/security-policies"
	err := makeRestAPIPayloadSecurityPolicy(d, m, "POST", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
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
	endpoint := "restapi/v3/security-policies/" + name
	err := makeRestAPIPayloadSecurityPolicy(d, m, "DELETE", endpoint)
	if err != nil {
		return fmt.Errorf("error occurred : %v", err)
	}
	return nil
}
