package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFDestinationNats() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFDestinationNatsCreate,
		Read:   resourceCudaWAFDestinationNatsRead,
		Update: resourceCudaWAFDestinationNatsUpdate,
		Delete: resourceCudaWAFDestinationNatsDelete,

		Schema: map[string]*schema.Schema{
			"comments":                 {Type: schema.TypeString, Optional: true},
			"vsite":                    {Type: schema.TypeString, Required: true},
			"incoming_interface":       {Type: schema.TypeString, Required: true},
			"post_destination_address": {Type: schema.TypeString, Required: true},
			"pre_destination_address":  {Type: schema.TypeString, Required: true},
			"pre_destination_netmask":  {Type: schema.TypeString, Required: true},
			"pre_destination_port":     {Type: schema.TypeString, Optional: true},
			"protocol":                 {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFDestinationNatsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/destination-nats"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFDestinationNatsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFDestinationNatsRead(d, m)
}

func resourceCudaWAFDestinationNatsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFDestinationNatsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/destination-nats/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFDestinationNatsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFDestinationNatsRead(d, m)
}

func resourceCudaWAFDestinationNatsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/destination-nats/"
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

func hydrateBarracudaWAFDestinationNatsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comments":                 d.Get("comments").(string),
		"vsite":                    d.Get("vsite").(string),
		"incoming-interface":       d.Get("incoming_interface").(string),
		"post-destination-address": d.Get("post_destination_address").(string),
		"pre-destination-address":  d.Get("pre_destination_address").(string),
		"pre-destination-netmask":  d.Get("pre_destination_netmask").(string),
		"pre-destination-port":     d.Get("pre_destination_port").(string),
		"protocol":                 d.Get("protocol").(string),
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
