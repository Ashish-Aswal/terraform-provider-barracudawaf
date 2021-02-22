package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFBotSpamPatterns() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFBotSpamPatternsCreate,
		Read:   resourceCudaWAFBotSpamPatternsRead,
		Update: resourceCudaWAFBotSpamPatternsUpdate,
		Delete: resourceCudaWAFBotSpamPatternsDelete,

		Schema: map[string]*schema.Schema{
			"algorithm":      {Type: schema.TypeString, Optional: true},
			"case_sensitive": {Type: schema.TypeString, Optional: true},
			"description":    {Type: schema.TypeString, Optional: true},
			"mode":           {Type: schema.TypeString, Optional: true},
			"name":           {Type: schema.TypeString, Required: true},
			"regex":          {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFBotSpamPatternsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFBotSpamPatternsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFBotSpamPatternsRead(d, m)
}

func resourceCudaWAFBotSpamPatternsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFBotSpamPatternsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFBotSpamPatternsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFBotSpamPatternsRead(d, m)
}

func resourceCudaWAFBotSpamPatternsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns/"
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

func hydrateBarracudaWAFBotSpamPatternsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"algorithm":      d.Get("algorithm").(string),
		"case-sensitive": d.Get("case_sensitive").(string),
		"description":    d.Get("description").(string),
		"mode":           d.Get("mode").(string),
		"name":           d.Get("name").(string),
		"regex":          d.Get("regex").(string),
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
