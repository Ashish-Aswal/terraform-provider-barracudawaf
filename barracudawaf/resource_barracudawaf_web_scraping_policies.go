package barracudawaf

import (
	"fmt"
	"log"
	"strings"

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

	client.hydrateBarracudaWAFWebScrapingPoliciesSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFWebScrapingPoliciesRead(d, m)
}

func resourceCudaWAFWebScrapingPoliciesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/web-scraping-policies"
	request := &APIRequest{
		Method: "get",
		URL:    resourceEndpoint,
	}

	var dataItems map[string]interface{}
	resources, err := client.GetBarracudaWAFResource(name, request)

	if err != nil {
		log.Printf("[ERROR] Unable to Retrieve Barracuda WAF resource (%s) (%v) ", name, err)
		return err
	}

	if resources.Data == nil {
		log.Printf("[WARN] Barracuda WAF resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	for _, dataItems = range resources.Data {
		if dataItems["name"] == name {
			break
		}
	}

	if dataItems["name"] != name {
		return fmt.Errorf("Barracuda WAF resource (%s) not found on the system", name)
	}

	d.Set("name", name)
	return nil
}

func resourceCudaWAFWebScrapingPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/web-scraping-policies"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFWebScrapingPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFWebScrapingPoliciesSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFWebScrapingPoliciesRead(d, m)
}

func resourceCudaWAFWebScrapingPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/web-scraping-policies"
	request := &APIRequest{
		Method: "delete",
		URL:    resourceEndpoint,
	}

	err := client.DeleteBarracudaWAFResource(name, request)

	if err != nil {
		return fmt.Errorf("Unable to delete the Barracuda WAF resource (%s) (%v)", name, err)
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
		for _, param := range updatePayloadExceptions {
			delete(resourcePayload, param)
		}
	}

	// remove empty parameters from resource payload
	for key, val := range resourcePayload {
		if len(val) == 0 {
			delete(resourcePayload, key)
		}
	}

	return &APIRequest{
		URL:  endpoint,
		Body: resourcePayload,
	}
}

func (b *BarracudaWAF) hydrateBarracudaWAFWebScrapingPoliciesSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {
	subResourceObjects := map[string][]string{}

	for subResource, subResourceParams := range subResourceObjects {
		subResourceParamsLength := d.Get(subResource + ".#").(int)

		if subResourceParamsLength > 0 {
			log.Printf("[INFO] Updating Barracuda WAF sub resource (%s) (%s)", name, subResource)

			for i := 0; i < subResourceParamsLength; i++ {
				subResourcePayload := map[string]string{}
				suffix := fmt.Sprintf(".%d", i)

				for _, param := range subResourceParams {
					paramSuffix := fmt.Sprintf(".%s", param)
					paramVaule := d.Get(subResource + suffix + paramSuffix).(string)

					param = strings.Replace(param, "_", "-", -1)
					subResourcePayload[param] = paramVaule
				}

				for key, val := range subResourcePayload {
					if len(val) == 0 {
						delete(subResourcePayload, key)
					}
				}

				err := b.UpdateBarracudaWAFSubResource(name, endpoint, &APIRequest{
					URL:  strings.Replace(subResource, "_", "-", -1),
					Body: subResourcePayload,
				})

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
