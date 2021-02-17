package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFWebScrapingPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFWebScrapingPoliciesCreate,
		Read:   resourceCudaWAFWebScrapingPoliciesRead,
		Update: resourceCudaWAFWebScrapingPoliciesUpdate,
		Delete: resourceCudaWAFWebScrapingPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"name":                          {Type: schema.TypeString, Required: true},
			"blacklisted_categories":        {Type: schema.TypeString, Optional: true},
			"whitelisted_bots":              {Type: schema.TypeString, Optional: true},
			"comments":                      {Type: schema.TypeString, Optional: true},
			"delay_time":                    {Type: schema.TypeString, Optional: true},
			"insert_delay":                  {Type: schema.TypeString, Optional: true},
			"insert_disallowed_urls":        {Type: schema.TypeString, Optional: true},
			"insert_hidden_links":           {Type: schema.TypeString, Optional: true},
			"insert_javascript_in_response": {Type: schema.TypeString, Optional: true},
			"detect_mouse_event":            {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadWebScrapingPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                          d.Get("name").(string),
		"blacklisted-categories":        d.Get("blacklisted_categories").(string),
		"whitelisted-bots":              d.Get("whitelisted_bots").(string),
		"comments":                      d.Get("comments").(string),
		"delay-time":                    d.Get("delay_time").(string),
		"insert-delay":                  d.Get("insert_delay").(string),
		"insert-disallowed-urls":        d.Get("insert_disallowed_urls").(string),
		"insert-hidden-links":           d.Get("insert_hidden_links").(string),
		"insert-javascript-in-response": d.Get("insert_javascript_in_response").(string),
		"detect-mouse-event":            d.Get("detect_mouse_event").(string),
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

func resourceCudaWAFWebScrapingPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/web-scraping-policies"
	resourceCreateResponseError := makeRestAPIPayloadWebScrapingPolicies(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFWebScrapingPoliciesRead(d, m)
}

func resourceCudaWAFWebScrapingPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFWebScrapingPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/web-scraping-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadWebScrapingPolicies(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFWebScrapingPoliciesRead(d, m)
}

func resourceCudaWAFWebScrapingPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/web-scraping-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadWebScrapingPolicies(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
