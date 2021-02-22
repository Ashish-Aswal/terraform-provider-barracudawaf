package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCredentialServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCredentialServersCreate,
		Read:   resourceCudaWAFCredentialServersRead,
		Update: resourceCudaWAFCredentialServersUpdate,
		Delete: resourceCudaWAFCredentialServersDelete,

		Schema: map[string]*schema.Schema{
			"cache_expiry":         {Type: schema.TypeString, Optional: true},
			"cache_valid_sessions": {Type: schema.TypeString, Optional: true},
			"redirect_url":         {Type: schema.TypeString, Optional: true},
			"ip_address":           {Type: schema.TypeString, Required: true},
			"policy_name":          {Type: schema.TypeString, Required: true},
			"port":                 {Type: schema.TypeString, Optional: true},
			"armored_browser_type": {Type: schema.TypeString, Required: true},
			"name":                 {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFCredentialServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/credential-servers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFCredentialServersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFCredentialServersRead(d, m)
}

func resourceCudaWAFCredentialServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFCredentialServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/credential-servers/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFCredentialServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFCredentialServersRead(d, m)
}

func resourceCudaWAFCredentialServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/credential-servers/"
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

func hydrateBarracudaWAFCredentialServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"cache-expiry":         d.Get("cache_expiry").(string),
		"cache-valid-sessions": d.Get("cache_valid_sessions").(string),
		"redirect-url":         d.Get("redirect_url").(string),
		"ip-address":           d.Get("ip_address").(string),
		"policy-name":          d.Get("policy_name").(string),
		"port":                 d.Get("port").(string),
		"armored-browser-type": d.Get("armored_browser_type").(string),
		"name":                 d.Get("name").(string),
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
