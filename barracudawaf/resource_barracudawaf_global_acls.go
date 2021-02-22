package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGlobalAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGlobalAclsCreate,
		Read:   resourceCudaWAFGlobalAclsRead,
		Update: resourceCudaWAFGlobalAclsUpdate,
		Delete: resourceCudaWAFGlobalAclsDelete,

		Schema: map[string]*schema.Schema{
			"action":                  {Type: schema.TypeString, Optional: true},
			"comments":                {Type: schema.TypeString, Optional: true},
			"deny_response":           {Type: schema.TypeString, Optional: true},
			"extended_match":          {Type: schema.TypeString, Optional: true},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true},
			"follow_up_action":        {Type: schema.TypeString, Optional: true},
			"follow_up_action_time":   {Type: schema.TypeString, Optional: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"redirect_url":            {Type: schema.TypeString, Optional: true},
			"response_page":           {Type: schema.TypeString, Optional: true},
			"enable":                  {Type: schema.TypeString, Optional: true},
			"url":                     {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFGlobalAclsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/global-acls"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGlobalAclsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFGlobalAclsRead(d, m)
}

func resourceCudaWAFGlobalAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGlobalAclsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/global-acls/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGlobalAclsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFGlobalAclsRead(d, m)
}

func resourceCudaWAFGlobalAclsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/global-acls/"
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

func hydrateBarracudaWAFGlobalAclsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"action":                  d.Get("action").(string),
		"comments":                d.Get("comments").(string),
		"deny-response":           d.Get("deny_response").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"follow-up-action":        d.Get("follow_up_action").(string),
		"follow-up-action-time":   d.Get("follow_up_action_time").(string),
		"name":                    d.Get("name").(string),
		"redirect-url":            d.Get("redirect_url").(string),
		"response-page":           d.Get("response_page").(string),
		"enable":                  d.Get("enable").(string),
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
