package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFDdosPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFDdosPoliciesCreate,
		Read:   resourceCudaWAFDdosPoliciesRead,
		Update: resourceCudaWAFDdosPoliciesUpdate,
		Delete: resourceCudaWAFDdosPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"evaluate_clients":        {Type: schema.TypeString, Optional: true},
			"comments":                {Type: schema.TypeString, Optional: true},
			"enforce_captcha":         {Type: schema.TypeString, Optional: true},
			"extended_match":          {Type: schema.TypeString, Optional: true},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true},
			"host":                    {Type: schema.TypeString, Optional: true},
			"expiry_time":             {Type: schema.TypeString, Optional: true},
			"mouse_check":             {Type: schema.TypeString, Optional: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"num_captcha_tries":       {Type: schema.TypeString, Optional: true},
			"num_unanswered_captcha":  {Type: schema.TypeString, Optional: true},
			"url":                     {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFDdosPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFDdosPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFDdosPoliciesRead(d, m)
}

func resourceCudaWAFDdosPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFDdosPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFDdosPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFDdosPoliciesRead(d, m)
}

func resourceCudaWAFDdosPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies/"
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

func hydrateBarracudaWAFDdosPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"evaluate-clients":        d.Get("evaluate_clients").(string),
		"comments":                d.Get("comments").(string),
		"enforce-captcha":         d.Get("enforce_captcha").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"host":                    d.Get("host").(string),
		"expiry-time":             d.Get("expiry_time").(string),
		"mouse-check":             d.Get("mouse_check").(string),
		"name":                    d.Get("name").(string),
		"num-captcha-tries":       d.Get("num_captcha_tries").(string),
		"num-unanswered-captcha":  d.Get("num_unanswered_captcha").(string),
		"url":                     d.Get("url").(string),
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
