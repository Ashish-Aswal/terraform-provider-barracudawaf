package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSourceNats() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSourceNatsCreate,
		Read:   resourceCudaWAFSourceNatsRead,
		Update: resourceCudaWAFSourceNatsUpdate,
		Delete: resourceCudaWAFSourceNatsDelete,

		Schema: map[string]*schema.Schema{
			"outgoing_interface":  {Type: schema.TypeString, Required: true},
			"post_source_address": {Type: schema.TypeString, Required: true},
			"protocol":            {Type: schema.TypeString, Required: true},
			"pre_source_address":  {Type: schema.TypeString, Required: true},
			"pre_source_netmask":  {Type: schema.TypeString, Required: true},
			"destination_port":    {Type: schema.TypeString, Optional: true},
			"comments":            {Type: schema.TypeString, Optional: true},
			"vsite":               {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFSourceNatsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/source-nats"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSourceNatsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSourceNatsRead(d, m)
}

func resourceCudaWAFSourceNatsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/source-nats"
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

func resourceCudaWAFSourceNatsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/source-nats/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSourceNatsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSourceNatsRead(d, m)
}

func resourceCudaWAFSourceNatsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/source-nats/"
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

func hydrateBarracudaWAFSourceNatsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"outgoing-interface":  d.Get("outgoing_interface").(string),
		"post-source-address": d.Get("post_source_address").(string),
		"protocol":            d.Get("protocol").(string),
		"pre-source-address":  d.Get("pre_source_address").(string),
		"pre-source-netmask":  d.Get("pre_source_netmask").(string),
		"destination-port":    d.Get("destination_port").(string),
		"comments":            d.Get("comments").(string),
		"vsite":               d.Get("vsite").(string),
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
