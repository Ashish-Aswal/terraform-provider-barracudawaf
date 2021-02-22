package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAccessRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAccessRulesCreate,
		Read:   resourceCudaWAFAccessRulesRead,
		Update: resourceCudaWAFAccessRulesUpdate,
		Delete: resourceCudaWAFAccessRulesDelete,

		Schema: map[string]*schema.Schema{
			"attribute_names":  {Type: schema.TypeString, Optional: true},
			"attribute_values": {Type: schema.TypeString, Optional: true},
			"name":             {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFAccessRulesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/access-rules"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAccessRulesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFAccessRulesRead(d, m)
}

func resourceCudaWAFAccessRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAccessRulesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/access-rules/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAccessRulesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAccessRulesRead(d, m)
}

func resourceCudaWAFAccessRulesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/access-rules/"
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

func hydrateBarracudaWAFAccessRulesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"attribute-names":  d.Get("attribute_names").(string),
		"attribute-values": d.Get("attribute_values").(string),
		"name":             d.Get("name").(string),
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
