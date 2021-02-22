package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFJsonSecurityPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFJsonSecurityPoliciesCreate,
		Read:   resourceCudaWAFJsonSecurityPoliciesRead,
		Update: resourceCudaWAFJsonSecurityPoliciesUpdate,
		Delete: resourceCudaWAFJsonSecurityPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"name":               {Type: schema.TypeString, Required: true},
			"max_array_elements": {Type: schema.TypeString, Optional: true},
			"max_siblings":       {Type: schema.TypeString, Optional: true},
			"max_keys":           {Type: schema.TypeString, Required: true},
			"max_key_length":     {Type: schema.TypeString, Required: true},
			"max_number_value":   {Type: schema.TypeString, Optional: true},
			"max_object_depth":   {Type: schema.TypeString, Optional: true},
			"max_value_length":   {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFJsonSecurityPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/json-security-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFJsonSecurityPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFJsonSecurityPoliciesRead(d, m)
}

func resourceCudaWAFJsonSecurityPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFJsonSecurityPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/json-security-policies/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFJsonSecurityPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFJsonSecurityPoliciesRead(d, m)
}

func resourceCudaWAFJsonSecurityPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/json-security-policies/"
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

func hydrateBarracudaWAFJsonSecurityPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":               d.Get("name").(string),
		"max-array-elements": d.Get("max_array_elements").(string),
		"max-siblings":       d.Get("max_siblings").(string),
		"max-keys":           d.Get("max_keys").(string),
		"max-key-length":     d.Get("max_key_length").(string),
		"max-number-value":   d.Get("max_number_value").(string),
		"max-object-depth":   d.Get("max_object_depth").(string),
		"max-value-length":   d.Get("max_value_length").(string),
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
