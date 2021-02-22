package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFIdentityTheftPatterns() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFIdentityTheftPatternsCreate,
		Read:   resourceCudaWAFIdentityTheftPatternsRead,
		Update: resourceCudaWAFIdentityTheftPatternsUpdate,
		Delete: resourceCudaWAFIdentityTheftPatternsDelete,

		Schema: map[string]*schema.Schema{
			"algorithm":      {Type: schema.TypeString, Optional: true},
			"case_sensitive": {Type: schema.TypeString, Optional: true},
			"description":    {Type: schema.TypeString, Optional: true},
			"name":           {Type: schema.TypeString, Required: true},
			"regex":          {Type: schema.TypeString, Required: true},
			"status":         {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFIdentityTheftPatternsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/identity-types/" + d.Get("parent.0").(string) + "/identity-theft-patterns"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFIdentityTheftPatternsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFIdentityTheftPatternsRead(d, m)
}

func resourceCudaWAFIdentityTheftPatternsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFIdentityTheftPatternsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/identity-types/" + d.Get("parent.0").(string) + "/identity-theft-patterns/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFIdentityTheftPatternsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFIdentityTheftPatternsRead(d, m)
}

func resourceCudaWAFIdentityTheftPatternsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/identity-types/" + d.Get("parent.0").(string) + "/identity-theft-patterns/"
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

func hydrateBarracudaWAFIdentityTheftPatternsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"algorithm":      d.Get("algorithm").(string),
		"case-sensitive": d.Get("case_sensitive").(string),
		"description":    d.Get("description").(string),
		"name":           d.Get("name").(string),
		"regex":          d.Get("regex").(string),
		"status":         d.Get("status").(string),
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
