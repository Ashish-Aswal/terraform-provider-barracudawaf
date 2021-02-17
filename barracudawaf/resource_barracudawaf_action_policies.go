package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFActionPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFActionPoliciesCreate,
		Read:   resourceCudaWAFActionPoliciesRead,
		Update: resourceCudaWAFActionPoliciesUpdate,
		Delete: resourceCudaWAFActionPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"action":                {Type: schema.TypeString, Optional: true},
			"deny_response":         {Type: schema.TypeString, Optional: true},
			"follow_up_action":      {Type: schema.TypeString, Optional: true},
			"follow_up_action_time": {Type: schema.TypeString, Optional: true},
			"name":                  {Type: schema.TypeString, Required: true},
			"redirect_url":          {Type: schema.TypeString, Optional: true},
			"response_page":         {Type: schema.TypeString, Optional: true},
			"risk_score":            {Type: schema.TypeString, Optional: true},
			"numeric_id":            {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadActionPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"action":                d.Get("action").(string),
		"deny-response":         d.Get("deny_response").(string),
		"follow-up-action":      d.Get("follow_up_action").(string),
		"follow-up-action-time": d.Get("follow_up_action_time").(string),
		"name":                  d.Get("name").(string),
		"redirect-url":          d.Get("redirect_url").(string),
		"response-page":         d.Get("response_page").(string),
		"risk-score":            d.Get("risk_score").(string),
		"numeric-id":            d.Get("numeric_id").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"name", "numeric-id"}
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

func resourceCudaWAFActionPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/action-policies"
	resourceCreateResponseError := makeRestAPIPayloadActionPolicies(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFActionPoliciesRead(d, m)
}

func resourceCudaWAFActionPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFActionPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/action-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadActionPolicies(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFActionPoliciesRead(d, m)
}

func resourceCudaWAFActionPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/action-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadActionPolicies(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
