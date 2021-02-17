package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFContentRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFContentRulesCreate,
		Read:   resourceCudaWAFContentRulesRead,
		Update: resourceCudaWAFContentRulesUpdate,
		Delete: resourceCudaWAFContentRulesDelete,

		Schema: map[string]*schema.Schema{
			"access_log":              {Type: schema.TypeString, Optional: true},
			"app_id":                  {Type: schema.TypeString, Optional: true},
			"comments":                {Type: schema.TypeString, Optional: true},
			"host_match":              {Type: schema.TypeString, Required: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"status":                  {Type: schema.TypeString, Optional: true},
			"extended_match":          {Type: schema.TypeString, Optional: true},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true},
			"mode":                    {Type: schema.TypeString, Optional: true},
			"url_match":               {Type: schema.TypeString, Required: true},
			"web_firewall_policy":     {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadContentRules(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"access-log":              d.Get("access_log").(string),
		"app-id":                  d.Get("app_id").(string),
		"comments":                d.Get("comments").(string),
		"host-match":              d.Get("host_match").(string),
		"name":                    d.Get("name").(string),
		"status":                  d.Get("status").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"mode":                    d.Get("mode").(string),
		"url-match":               d.Get("url_match").(string),
		"web-firewall-policy":     d.Get("web_firewall_policy").(string),
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

func resourceCudaWAFContentRulesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/content-rules"
	resourceCreateResponseError := makeRestAPIPayloadContentRules(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFContentRulesRead(d, m)
}

func resourceCudaWAFContentRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFContentRulesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/content-rules/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadContentRules(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFContentRulesRead(d, m)
}

func resourceCudaWAFContentRulesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/content-rules/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadContentRules(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
