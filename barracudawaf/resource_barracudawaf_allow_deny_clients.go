package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAllowDenyClients() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAllowDenyClientsCreate,
		Read:   resourceCudaWAFAllowDenyClientsRead,
		Update: resourceCudaWAFAllowDenyClientsUpdate,
		Delete: resourceCudaWAFAllowDenyClientsDelete,

		Schema: map[string]*schema.Schema{
			"name":                {Type: schema.TypeString, Required: true},
			"action":              {Type: schema.TypeString, Optional: true},
			"sequence":            {Type: schema.TypeString, Optional: true},
			"certificate_serial":  {Type: schema.TypeString, Optional: true},
			"common_name":         {Type: schema.TypeString, Optional: true},
			"country":             {Type: schema.TypeString, Optional: true},
			"locality":            {Type: schema.TypeString, Optional: true},
			"organization":        {Type: schema.TypeString, Optional: true},
			"organizational_unit": {Type: schema.TypeString, Optional: true},
			"state":               {Type: schema.TypeString, Optional: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFAllowDenyClientsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAllowDenyClientsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFAllowDenyClientsRead(d, m)
}

func resourceCudaWAFAllowDenyClientsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAllowDenyClientsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAllowDenyClientsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAllowDenyClientsRead(d, m)
}

func resourceCudaWAFAllowDenyClientsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients/"
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

func hydrateBarracudaWAFAllowDenyClientsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":                d.Get("name").(string),
		"action":              d.Get("action").(string),
		"sequence":            d.Get("sequence").(string),
		"certificate-serial":  d.Get("certificate_serial").(string),
		"common-name":         d.Get("common_name").(string),
		"country":             d.Get("country").(string),
		"locality":            d.Get("locality").(string),
		"organization":        d.Get("organization").(string),
		"organizational-unit": d.Get("organizational_unit").(string),
		"state":               d.Get("state").(string),
		"status":              d.Get("status").(string),
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
