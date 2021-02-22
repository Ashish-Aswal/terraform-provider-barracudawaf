package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFProtectedDataTypes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFProtectedDataTypesCreate,
		Read:   resourceCudaWAFProtectedDataTypesRead,
		Update: resourceCudaWAFProtectedDataTypesUpdate,
		Delete: resourceCudaWAFProtectedDataTypesDelete,

		Schema: map[string]*schema.Schema{
			"action":                      {Type: schema.TypeString, Optional: true},
			"initial_characters_to_keep":  {Type: schema.TypeString, Optional: true},
			"trailing_characters_to_keep": {Type: schema.TypeString, Optional: true},
			"name":                        {Type: schema.TypeString, Required: true},
			"custom_identity_theft_type":  {Type: schema.TypeString, Optional: true},
			"enable":                      {Type: schema.TypeString, Optional: true},
			"identity_theft_type":         {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFProtectedDataTypesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/protected-data-types"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFProtectedDataTypesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFProtectedDataTypesRead(d, m)
}

func resourceCudaWAFProtectedDataTypesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFProtectedDataTypesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/protected-data-types/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFProtectedDataTypesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFProtectedDataTypesRead(d, m)
}

func resourceCudaWAFProtectedDataTypesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/" + d.Get("parent.0").(string) + "/protected-data-types/"
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

func hydrateBarracudaWAFProtectedDataTypesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"action":                      d.Get("action").(string),
		"initial-characters-to-keep":  d.Get("initial_characters_to_keep").(string),
		"trailing-characters-to-keep": d.Get("trailing_characters_to_keep").(string),
		"name":                        d.Get("name").(string),
		"custom-identity-theft-type":  d.Get("custom_identity_theft_type").(string),
		"enable":                      d.Get("enable").(string),
		"identity-theft-type":         d.Get("identity_theft_type").(string),
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
