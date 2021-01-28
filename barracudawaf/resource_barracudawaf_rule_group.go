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

func makeRestAPIPayloadRuleGroup(d *schema.ResourceData, m interface{}, oper string, endpoint string) error {
	payload := map[string]string{
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

func resourceCudaWAFRuleGroupCreate(d *schema.ResourceData, m interface{}) error {
	endpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/content-rules"
	err := makeRestAPIPayloadRuleGroup(d, m, "POST", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
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
	endpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/content-rules/" + name
	err := makeRestAPIPayloadRuleGroup(d, m, "DELETE", endpoint)
	if err != nil {
		return fmt.Errorf("error occurred : %v", err)
	}
	return nil
}
