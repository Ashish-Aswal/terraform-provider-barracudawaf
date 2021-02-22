package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRateControlPools() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRateControlPoolsCreate,
		Read:   resourceCudaWAFRateControlPoolsRead,
		Update: resourceCudaWAFRateControlPoolsUpdate,
		Delete: resourceCudaWAFRateControlPoolsDelete,

		Schema: map[string]*schema.Schema{
			"name":                     {Type: schema.TypeString, Required: true},
			"max_active_requests":      {Type: schema.TypeString, Optional: true},
			"max_per_client_backlog":   {Type: schema.TypeString, Optional: true},
			"max_unconfigured_clients": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFRateControlPoolsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/rate-control-pools"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFRateControlPoolsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFRateControlPoolsRead(d, m)
}

func resourceCudaWAFRateControlPoolsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFRateControlPoolsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/rate-control-pools/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFRateControlPoolsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFRateControlPoolsRead(d, m)
}

func resourceCudaWAFRateControlPoolsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/rate-control-pools/"
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

func hydrateBarracudaWAFRateControlPoolsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":                     d.Get("name").(string),
		"max-active-requests":      d.Get("max_active_requests").(string),
		"max-per-client-backlog":   d.Get("max_per_client_backlog").(string),
		"max-unconfigured-clients": d.Get("max_unconfigured_clients").(string),
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
