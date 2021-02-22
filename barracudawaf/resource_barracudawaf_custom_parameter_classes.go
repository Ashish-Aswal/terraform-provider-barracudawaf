package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCustomParameterClasses() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCustomParameterClassesCreate,
		Read:   resourceCudaWAFCustomParameterClassesRead,
		Update: resourceCudaWAFCustomParameterClassesUpdate,
		Delete: resourceCudaWAFCustomParameterClassesDelete,

		Schema: map[string]*schema.Schema{
			"name":                         {Type: schema.TypeString, Required: true},
			"custom_blocked_attack_types":  {Type: schema.TypeString, Optional: true},
			"custom_input_type_validation": {Type: schema.TypeString, Optional: true},
			"denied_metacharacters":        {Type: schema.TypeString, Optional: true},
			"input_type_validation":        {Type: schema.TypeString, Optional: true},
			"blocked_attack_types":         {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFCustomParameterClassesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/custom-parameter-classes"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFCustomParameterClassesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFCustomParameterClassesRead(d, m)
}

func resourceCudaWAFCustomParameterClassesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFCustomParameterClassesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/custom-parameter-classes/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFCustomParameterClassesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFCustomParameterClassesRead(d, m)
}

func resourceCudaWAFCustomParameterClassesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/custom-parameter-classes/"
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

func hydrateBarracudaWAFCustomParameterClassesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":                         d.Get("name").(string),
		"custom-blocked-attack-types":  d.Get("custom_blocked_attack_types").(string),
		"custom-input-type-validation": d.Get("custom_input_type_validation").(string),
		"denied-metacharacters":        d.Get("denied_metacharacters").(string),
		"input-type-validation":        d.Get("input_type_validation").(string),
		"blocked-attack-types":         d.Get("blocked_attack_types").(string),
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
