package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNodes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNodesCreate,
		Read:   resourceCudaWAFNodesRead,
		Update: resourceCudaWAFNodesUpdate,
		Delete: resourceCudaWAFNodesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {Type: schema.TypeString, Required: true},
			"mode":       {Type: schema.TypeString, Required: true},
			"serial":     {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFNodesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/nodes"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNodesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFNodesRead(d, m)
}

func resourceCudaWAFNodesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNodesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/nodes/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNodesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFNodesRead(d, m)
}

func resourceCudaWAFNodesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/nodes/"
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

func hydrateBarracudaWAFNodesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"ip-address": d.Get("ip_address").(string),
		"mode":       d.Get("mode").(string),
		"serial":     d.Get("serial").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"mode", "serial"}
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
