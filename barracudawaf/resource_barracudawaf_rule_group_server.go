package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRuleGroupServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRuleGroupServerCreate,
		Read:   resourceCudaWAFRuleGroupServerRead,
		Update: resourceCudaWAFRuleGroupServerUpdate,
		Delete: resourceCudaWAFRuleGroupServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadRuleGroupServer(d *schema.ResourceData, m interface{}, oper string, endpoint string) error {
	payload := map[string]string{
		"name":            d.Get("name").(string),
		"port":            d.Get("port").(string),
		"status":          d.Get("status").(string),
		"comments":        d.Get("comments").(string),
		"hostname":        d.Get("hostname").(string),
		"identifier":      d.Get("identifier").(string),
		"ip-address":      d.Get("ip_address").(string),
		"address-version": d.Get("address_version").(string),
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

func resourceCudaWAFRuleGroupServerCreate(d *schema.ResourceData, m interface{}) error {
	serviceName := d.Get("service_name").(string)
	ruleGroupName := d.Get("rule_group_name").(string)
	endpoint := "restapi/v3/services/" + serviceName + "/content-rules/" + ruleGroupName + "/content-rule-servers"
	err := makeRestAPIPayloadRuleGroupServer(d, m, "POST", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourceCudaWAFRuleGroupServerRead(d, m)
}

func resourceCudaWAFRuleGroupServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFRuleGroupServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCudaWAFRuleGroupServerRead(d, m)
}

func resourceCudaWAFRuleGroupServerDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	serviceName := d.Get("service_name").(string)
	ruleGroupName := d.Get("rule_group_name").(string)
	endpoint := "restapi/v3/services/" + serviceName + "/content-rules/" + ruleGroupName + "/content-rule-servers/" + name
	err := makeRestAPIPayloadRuleGroupServer(d, m, "DELETE", endpoint)
	if err != nil {
		return fmt.Errorf("error occurred : %v", err)
	}
	return nil
}
