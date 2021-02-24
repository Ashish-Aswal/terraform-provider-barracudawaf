package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlEncryptionRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlEncryptionRulesCreate,
		Read:   resourceCudaWAFUrlEncryptionRulesRead,
		Update: resourceCudaWAFUrlEncryptionRulesUpdate,
		Delete: resourceCudaWAFUrlEncryptionRulesDelete,

		Schema: map[string]*schema.Schema{
			"allow_unencrypted_requests": {Type: schema.TypeString, Optional: true},
			"exclude_urls":               {Type: schema.TypeString, Optional: true},
			"host":                       {Type: schema.TypeString, Required: true},
			"name":                       {Type: schema.TypeString, Required: true},
			"status":                     {Type: schema.TypeString, Optional: true},
			"url":                        {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFUrlEncryptionRulesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFUrlEncryptionRulesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFUrlEncryptionRulesRead(d, m)
}

func resourceCudaWAFUrlEncryptionRulesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules"
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

func resourceCudaWAFUrlEncryptionRulesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFUrlEncryptionRulesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFUrlEncryptionRulesRead(d, m)
}

func resourceCudaWAFUrlEncryptionRulesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules/"
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

func hydrateBarracudaWAFUrlEncryptionRulesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"allow-unencrypted-requests": d.Get("allow_unencrypted_requests").(string),
		"exclude-urls":               d.Get("exclude_urls").(string),
		"host":                       d.Get("host").(string),
		"name":                       d.Get("name").(string),
		"status":                     d.Get("status").(string),
		"url":                        d.Get("url").(string),
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
