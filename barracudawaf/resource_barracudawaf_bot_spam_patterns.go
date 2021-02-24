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
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns"
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

func resourceCudaWAFBotSpamPatternsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns/"
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
		return fmt.Errorf("Unable to delete the Barracuda WAF resource (%s) (%v)", name, err)
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
