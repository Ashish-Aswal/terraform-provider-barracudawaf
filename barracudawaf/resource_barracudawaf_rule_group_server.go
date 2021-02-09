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

func makeRestAPIPayloadRuleGroupServer(d *schema.ResourceData, m interface{}, resourceOperation string, resourceEndpoint string) error {

	//build Payload for the resource
	resourcePayload := map[string]string{
		"name":            d.Get("name").(string),
		"port":            d.Get("port").(string),
		"status":          d.Get("status").(string),
		"comments":        d.Get("comments").(string),
		"hostname":        d.Get("hostname").(string),
		"identifier":      d.Get("identifier").(string),
		"ip-address":      d.Get("ip_address").(string),
		"address-version": d.Get("address_version").(string),
	}

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

func resourceCudaWAFRuleGroupServerCreate(d *schema.ResourceData, m interface{}) error {
	serviceName := d.Get("service_name").(string)
	ruleGroupName := d.Get("rule_group_name").(string)
	resourceEndpoint := "restapi/v3/services/" + serviceName + "/content-rules/" + ruleGroupName + "/content-rule-servers"
	resourceCreateResponseError := makeRestAPIPayloadRuleGroupServer(d, m, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
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
	resourceEndpoint := "restapi/v3/services/" + serviceName + "/content-rules/" + ruleGroupName + "/content-rule-servers/" + name
	resourceDeleteResponseError := makeRestAPIPayloadRuleGroupServer(d, m, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
