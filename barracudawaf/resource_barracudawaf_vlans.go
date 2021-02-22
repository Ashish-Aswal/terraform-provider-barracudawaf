package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFVlans() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFVlansCreate,
		Read:   resourceCudaWAFVlansRead,
		Update: resourceCudaWAFVlansUpdate,
		Delete: resourceCudaWAFVlansDelete,

		Schema: map[string]*schema.Schema{
			"comments":  {Type: schema.TypeString, Optional: true},
			"vlan_id":   {Type: schema.TypeString, Required: true},
			"interface": {Type: schema.TypeString, Required: true},
			"name":      {Type: schema.TypeString, Required: true},
			"vsite":     {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFVlansCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/vlans"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFVlansResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFVlansRead(d, m)
}

func resourceCudaWAFVlansRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFVlansUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/vlans/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFVlansResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFVlansRead(d, m)
}

func resourceCudaWAFVlansDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/vlans/"
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

func hydrateBarracudaWAFVlansResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comments":  d.Get("comments").(string),
		"vlan-id":   d.Get("vlan_id").(string),
		"interface": d.Get("interface").(string),
		"name":      d.Get("name").(string),
		"vsite":     d.Get("vsite").(string),
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
