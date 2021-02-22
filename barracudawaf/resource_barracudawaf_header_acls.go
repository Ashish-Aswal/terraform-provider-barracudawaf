package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFHeaderAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFHeaderAclsCreate,
		Read:   resourceCudaWAFHeaderAclsRead,
		Update: resourceCudaWAFHeaderAclsUpdate,
		Delete: resourceCudaWAFHeaderAclsDelete,

		Schema: map[string]*schema.Schema{
			"comments":                    {Type: schema.TypeString, Optional: true},
			"max_header_value_length":     {Type: schema.TypeString, Optional: true},
			"header_name":                 {Type: schema.TypeString, Required: true},
			"blocked_attack_types":        {Type: schema.TypeString, Optional: true},
			"denied_metachars":            {Type: schema.TypeString, Optional: true},
			"mode":                        {Type: schema.TypeString, Optional: true},
			"custom_blocked_attack_types": {Type: schema.TypeString, Optional: true},
			"exception_patterns":          {Type: schema.TypeString, Optional: true},
			"status":                      {Type: schema.TypeString, Optional: true},
			"name":                        {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFHeaderAclsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/header-acls"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFHeaderAclsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFHeaderAclsRead(d, m)
}

func resourceCudaWAFHeaderAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFHeaderAclsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/header-acls/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFHeaderAclsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFHeaderAclsRead(d, m)
}

func resourceCudaWAFHeaderAclsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/header-acls/"
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

func hydrateBarracudaWAFHeaderAclsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comments":                    d.Get("comments").(string),
		"max-header-value-length":     d.Get("max_header_value_length").(string),
		"header-name":                 d.Get("header_name").(string),
		"blocked-attack-types":        d.Get("blocked_attack_types").(string),
		"denied-metachars":            d.Get("denied_metachars").(string),
		"mode":                        d.Get("mode").(string),
		"custom-blocked-attack-types": d.Get("custom_blocked_attack_types").(string),
		"exception-patterns":          d.Get("exception_patterns").(string),
		"status":                      d.Get("status").(string),
		"name":                        d.Get("name").(string),
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
