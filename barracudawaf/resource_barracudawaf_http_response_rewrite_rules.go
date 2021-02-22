package barracudawaf

import (
	"fmt"
	"log"

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

func resourceCudaWAFHttpResponseRewriteRulesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/http-response-rewrite-rules"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFHttpResponseRewriteRulesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFHttpResponseRewriteRulesRead(d, m)
}

func resourceCudaWAFHttpResponseRewriteRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFHttpResponseRewriteRulesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/http-response-rewrite-rules/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFHttpResponseRewriteRulesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFHttpResponseRewriteRulesRead(d, m)
}

func resourceCudaWAFHttpResponseRewriteRulesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/http-response-rewrite-rules/"
	request := &APIRequest{
		Method: "delete",
		URL:    resourceEndpoint,
	}

	err := client.DeleteBarracudaWAFResource(name, request)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func hydrateBarracudaWAFHttpResponseRewriteRulesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
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

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{}
		for item := range updatePayloadExceptions {
			delete(resourcePayload, updatePayloadExceptions[item])
		}
	}

	// remove empty parameters from resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	return &APIRequest{
		URL:  endpoint,
		Body: resourcePayload,
	}
}
