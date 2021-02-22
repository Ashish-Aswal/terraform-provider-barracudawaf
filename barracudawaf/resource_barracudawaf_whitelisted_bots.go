package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFWhitelistedBots() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFWhitelistedBotsCreate,
		Read:   resourceCudaWAFWhitelistedBotsRead,
		Update: resourceCudaWAFWhitelistedBotsUpdate,
		Delete: resourceCudaWAFWhitelistedBotsDelete,

		Schema: map[string]*schema.Schema{
			"host":       {Type: schema.TypeString, Optional: true},
			"identifier": {Type: schema.TypeString, Optional: true},
			"ip_address": {Type: schema.TypeString, Optional: true},
			"name":       {Type: schema.TypeString, Required: true},
			"user_agent": {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFWhitelistedBotsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/whitelisted-bots"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFWhitelistedBotsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFWhitelistedBotsRead(d, m)
}

func resourceCudaWAFWhitelistedBotsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFWhitelistedBotsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/whitelisted-bots/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFWhitelistedBotsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFWhitelistedBotsRead(d, m)
}

func resourceCudaWAFWhitelistedBotsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/whitelisted-bots/"
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

func hydrateBarracudaWAFWhitelistedBotsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"host":       d.Get("host").(string),
		"identifier": d.Get("identifier").(string),
		"ip-address": d.Get("ip_address").(string),
		"name":       d.Get("name").(string),
		"user-agent": d.Get("user_agent").(string),
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
