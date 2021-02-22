package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSecurityPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSecurityPoliciesCreate,
		Read:   resourceCudaWAFSecurityPoliciesRead,
		Update: resourceCudaWAFSecurityPoliciesUpdate,
		Delete: resourceCudaWAFSecurityPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"based_on": {Type: schema.TypeString, Optional: true},
			"name":     {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFSecurityPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSecurityPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSecurityPoliciesRead(d, m)
}

func resourceCudaWAFSecurityPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSecurityPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/security-policies/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSecurityPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSecurityPoliciesRead(d, m)
}

func resourceCudaWAFSecurityPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies/"
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

func hydrateBarracudaWAFSecurityPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"based-on": d.Get("based_on").(string),
		"name":     d.Get("name").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"based-on"}
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
