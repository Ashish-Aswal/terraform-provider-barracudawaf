package barracudawaf

import (
	"fmt"
	"log"

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

func resourceCudaWAFActionPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/action-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFActionPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFActionPoliciesRead(d, m)
}

func resourceCudaWAFActionPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFActionPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/action-policies/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFActionPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFActionPoliciesRead(d, m)
}

func resourceCudaWAFActionPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/action-policies/"
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

func hydrateBarracudaWAFActionPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
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

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"name", "numeric-id"}
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
