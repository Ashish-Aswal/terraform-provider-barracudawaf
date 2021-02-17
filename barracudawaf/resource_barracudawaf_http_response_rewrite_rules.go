package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFHttpResponseRewriteRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFHttpResponseRewriteRulesCreate,
		Read:   resourceCudaWAFHttpResponseRewriteRulesRead,
		Update: resourceCudaWAFHttpResponseRewriteRulesUpdate,
		Delete: resourceCudaWAFHttpResponseRewriteRulesDelete,

		Schema: map[string]*schema.Schema{
			"action":              {Type: schema.TypeString, Required: true},
			"comments":            {Type: schema.TypeString, Optional: true},
			"condition":           {Type: schema.TypeString, Optional: true},
			"continue_processing": {Type: schema.TypeString, Optional: true},
			"header":              {Type: schema.TypeString, Optional: true},
			"old_value":           {Type: schema.TypeString, Optional: true},
			"sequence_number":     {Type: schema.TypeString, Required: true},
			"rewrite_value":       {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadHttpResponseRewriteRules(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"action":              d.Get("action").(string),
		"comments":            d.Get("comments").(string),
		"condition":           d.Get("condition").(string),
		"continue-processing": d.Get("continue_processing").(string),
		"header":              d.Get("header").(string),
		"old-value":           d.Get("old_value").(string),
		"sequence-number":     d.Get("sequence_number").(string),
		"rewrite-value":       d.Get("rewrite_value").(string),
		"name":                d.Get("name").(string),
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

func resourceCudaWAFHttpResponseRewriteRulesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/http-response-rewrite-rules"
	resourceCreateResponseError := makeRestAPIPayloadHttpResponseRewriteRules(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFHttpResponseRewriteRulesRead(d, m)
}

func resourceCudaWAFHttpResponseRewriteRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFHttpResponseRewriteRulesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/http-response-rewrite-rules/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadHttpResponseRewriteRules(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFHttpResponseRewriteRulesRead(d, m)
}

func resourceCudaWAFHttpResponseRewriteRulesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/http-response-rewrite-rules/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadHttpResponseRewriteRules(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
