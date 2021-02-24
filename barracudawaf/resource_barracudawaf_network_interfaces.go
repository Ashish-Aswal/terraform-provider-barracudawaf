package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNetworkInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNetworkInterfacesCreate,
		Read:   resourceCudaWAFNetworkInterfacesRead,
		Update: resourceCudaWAFNetworkInterfacesUpdate,
		Delete: resourceCudaWAFNetworkInterfacesDelete,

		Schema: map[string]*schema.Schema{
			"name":                    {Type: schema.TypeString, Required: true},
			"duplexity":               {Type: schema.TypeString, Required: true},
			"auto_negotiation_status": {Type: schema.TypeString, Optional: true},
			"speed":                   {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFNetworkInterfacesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/network-interfaces"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNetworkInterfacesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFNetworkInterfacesRead(d, m)
}

func resourceCudaWAFNetworkInterfacesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/network-interfaces"
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

func resourceCudaWAFNetworkInterfacesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/network-interfaces/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNetworkInterfacesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFNetworkInterfacesRead(d, m)
}

func resourceCudaWAFNetworkInterfacesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/network-interfaces/"
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

func hydrateBarracudaWAFNetworkInterfacesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":                    d.Get("name").(string),
		"duplexity":               d.Get("duplexity").(string),
		"auto-negotiation-status": d.Get("auto_negotiation_status").(string),
		"speed":                   d.Get("speed").(string),
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
