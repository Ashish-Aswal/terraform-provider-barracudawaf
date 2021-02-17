package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFResponseBodyRewriteRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFResponseBodyRewriteRulesCreate,
		Read:   resourceCudaWAFResponseBodyRewriteRulesRead,
		Update: resourceCudaWAFResponseBodyRewriteRulesUpdate,
		Delete: resourceCudaWAFResponseBodyRewriteRulesDelete,

		Schema: map[string]*schema.Schema{
			"comments":        {Type: schema.TypeString, Optional: true},
			"host":            {Type: schema.TypeString, Required: true},
			"name":            {Type: schema.TypeString, Required: true},
			"replace_string":  {Type: schema.TypeString, Optional: true},
			"search_string":   {Type: schema.TypeString, Required: true},
			"sequence_number": {Type: schema.TypeString, Required: true},
			"url":             {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadResponseBodyRewriteRules(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":        d.Get("comments").(string),
		"host":            d.Get("host").(string),
		"name":            d.Get("name").(string),
		"replace-string":  d.Get("replace_string").(string),
		"search-string":   d.Get("search_string").(string),
		"sequence-number": d.Get("sequence_number").(string),
		"url":             d.Get("url").(string),
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

func resourceCudaWAFResponseBodyRewriteRulesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/response-body-rewrite-rules"
	resourceCreateResponseError := makeRestAPIPayloadResponseBodyRewriteRules(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFResponseBodyRewriteRulesRead(d, m)
}

func resourceCudaWAFResponseBodyRewriteRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFResponseBodyRewriteRulesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/response-body-rewrite-rules/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadResponseBodyRewriteRules(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFResponseBodyRewriteRulesRead(d, m)
}

func resourceCudaWAFResponseBodyRewriteRulesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/response-body-rewrite-rules/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadResponseBodyRewriteRules(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
