package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRuleGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRuleGroupCreate,
		Read:   resourceCudaWAFRuleGroupRead,
		Update: resourceCudaWAFRuleGroupUpdate,
		Delete: resourceCudaWAFRuleGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_match": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_log": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_match": {
				Type:     schema.TypeString,
				Required: true,
			},
			"extended_match": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"web_firewall_policy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"extended_match_sequence": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadRuleGroup(d *schema.ResourceData, m interface{}, resourceOperation string, resourceEndpoint string) error {

	//build Payload for the resource
	resourcePayload := map[string]string{
		"name":                    d.Get("name").(string),
		"mode":                    d.Get("mode").(string),
		"status":                  d.Get("status").(string),
		"url-match":               d.Get("url_match").(string),
		"access-log":              d.Get("access_log").(string),
		"host-match":              d.Get("host_match").(string),
		"extended-match":          d.Get("extended_match").(string),
		"web-firewall-policy":     d.Get("web_firewall_policy").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
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

func resourceCudaWAFRuleGroupCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/content-rules"
	resourceUpdateResponseError := makeRestAPIPayloadRuleGroup(d, m, "POST", resourceEndpoint)
	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}
	return resourceCudaWAFRuleGroupRead(d, m)
}

func resourceCudaWAFRuleGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFRuleGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCudaWAFRuleGroupRead(d, m)
}

func resourceCudaWAFRuleGroupDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceEndpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/content-rules/" + name
	resourceDeleteResponseError := makeRestAPIPayloadRuleGroup(d, m, "DELETE", resourceEndpoint)
	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}
	return nil
}
