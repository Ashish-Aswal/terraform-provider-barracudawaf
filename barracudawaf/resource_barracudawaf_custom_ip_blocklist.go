package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCustomIpBlocklist() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCustomIpBlocklistCreate,
		Read:   resourceCudaWAFCustomIpBlocklistRead,
		Update: resourceCudaWAFCustomIpBlocklistUpdate,
		Delete: resourceCudaWAFCustomIpBlocklistDelete,

		Schema: map[string]*schema.Schema{
			"blacklisted_ips":             {Type: schema.TypeString, Optional: true},
			"custom_ip_list":              {Type: schema.TypeString, Optional: true},
			"download_url":                {Type: schema.TypeString, Optional: true},
			"trusted_certificate":         {Type: schema.TypeString, Optional: true},
			"validate_server_certificate": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFCustomIpBlocklistCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/custom-ip-blocklist"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFCustomIpBlocklistResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFCustomIpBlocklistRead(d, m)
}

func resourceCudaWAFCustomIpBlocklistRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/custom-ip-blocklist"
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

func resourceCudaWAFCustomIpBlocklistUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/custom-ip-blocklist/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFCustomIpBlocklistResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFCustomIpBlocklistRead(d, m)
}

func resourceCudaWAFCustomIpBlocklistDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/custom-ip-blocklist/"
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

func hydrateBarracudaWAFCustomIpBlocklistResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"blacklisted-ips":             d.Get("blacklisted_ips").(string),
		"custom-ip-list":              d.Get("custom_ip_list").(string),
		"download-url":                d.Get("download_url").(string),
		"trusted-certificate":         d.Get("trusted_certificate").(string),
		"validate-server-certificate": d.Get("validate_server_certificate").(string),
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
