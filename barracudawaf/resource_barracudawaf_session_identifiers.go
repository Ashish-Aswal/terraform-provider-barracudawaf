package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSessionIdentifiers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSessionIdentifiersCreate,
		Read:   resourceCudaWAFSessionIdentifiersRead,
		Update: resourceCudaWAFSessionIdentifiersUpdate,
		Delete: resourceCudaWAFSessionIdentifiersDelete,

		Schema: map[string]*schema.Schema{
			"name":            {Type: schema.TypeString, Required: true},
			"token_name":      {Type: schema.TypeString, Required: true},
			"token_type":      {Type: schema.TypeString, Required: true},
			"end_delimiter":   {Type: schema.TypeString, Optional: true},
			"start_delimiter": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFSessionIdentifiersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/session-identifiers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSessionIdentifiersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSessionIdentifiersRead(d, m)
}

func resourceCudaWAFSessionIdentifiersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSessionIdentifiersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/session-identifiers/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSessionIdentifiersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSessionIdentifiersRead(d, m)
}

func resourceCudaWAFSessionIdentifiersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/session-identifiers/"
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

func hydrateBarracudaWAFSessionIdentifiersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":            d.Get("name").(string),
		"token-name":      d.Get("token_name").(string),
		"token-type":      d.Get("token_type").(string),
		"end-delimiter":   d.Get("end_delimiter").(string),
		"start-delimiter": d.Get("start_delimiter").(string),
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
