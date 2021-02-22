package barracudawaf

import (
	"fmt"
	"log"

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

func resourceCudaWAFWebScrapingPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/web-scraping-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFWebScrapingPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFWebScrapingPoliciesRead(d, m)
}

func resourceCudaWAFWebScrapingPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFWebScrapingPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/web-scraping-policies/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFWebScrapingPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFWebScrapingPoliciesRead(d, m)
}

func resourceCudaWAFWebScrapingPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/web-scraping-policies/"
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

func hydrateBarracudaWAFWebScrapingPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
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
