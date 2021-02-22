package barracudawaf

import (
	"fmt"
	"log"

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

func resourceCudaWAFResponseBodyRewriteRulesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/response-body-rewrite-rules"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFResponseBodyRewriteRulesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFResponseBodyRewriteRulesRead(d, m)
}

func resourceCudaWAFResponseBodyRewriteRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFResponseBodyRewriteRulesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/response-body-rewrite-rules/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFResponseBodyRewriteRulesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFResponseBodyRewriteRulesRead(d, m)
}

func resourceCudaWAFResponseBodyRewriteRulesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/response-body-rewrite-rules/"
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

func hydrateBarracudaWAFResponseBodyRewriteRulesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comments":        d.Get("comments").(string),
		"host":            d.Get("host").(string),
		"name":            d.Get("name").(string),
		"replace-string":  d.Get("replace_string").(string),
		"search-string":   d.Get("search_string").(string),
		"sequence-number": d.Get("sequence_number").(string),
		"url":             d.Get("url").(string),
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
