package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAdaptiveProfilingRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAdaptiveProfilingRulesCreate,
		Read:   resourceCudaWAFAdaptiveProfilingRulesRead,
		Update: resourceCudaWAFAdaptiveProfilingRulesUpdate,
		Delete: resourceCudaWAFAdaptiveProfilingRulesDelete,

		Schema: map[string]*schema.Schema{
			"host":                {Type: schema.TypeString, Required: true},
			"learn_from_request":  {Type: schema.TypeString, Optional: true},
			"learn_from_response": {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"url":                 {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadAdaptiveProfilingRules(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"host":                d.Get("host").(string),
		"learn-from-request":  d.Get("learn_from_request").(string),
		"learn-from-response": d.Get("learn_from_response").(string),
		"name":                d.Get("name").(string),
		"status":              d.Get("status").(string),
		"url":                 d.Get("url").(string),
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

func resourceCudaWAFAdaptiveProfilingRulesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/adaptive-profiling-rules"
	resourceCreateResponseError := makeRestAPIPayloadAdaptiveProfilingRules(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFAdaptiveProfilingRulesRead(d, m)
}

func resourceCudaWAFAdaptiveProfilingRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAdaptiveProfilingRulesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/adaptive-profiling-rules/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadAdaptiveProfilingRules(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFAdaptiveProfilingRulesRead(d, m)
}

func resourceCudaWAFAdaptiveProfilingRulesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/adaptive-profiling-rules/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadAdaptiveProfilingRules(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}