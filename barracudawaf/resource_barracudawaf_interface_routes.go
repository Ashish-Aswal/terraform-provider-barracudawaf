package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFInterfaceRoutes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFInterfaceRoutesCreate,
		Read:   resourceCudaWAFInterfaceRoutesRead,
		Update: resourceCudaWAFInterfaceRoutesUpdate,
		Delete: resourceCudaWAFInterfaceRoutesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {Type: schema.TypeString, Required: true},
			"comments":   {Type: schema.TypeString, Optional: true},
			"interface":  {Type: schema.TypeString, Required: true},
			"ip_version": {Type: schema.TypeString, Optional: true},
			"netmask":    {Type: schema.TypeString, Required: true},
			"vsite":      {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFInterfaceRoutesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/interface-routes"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFInterfaceRoutesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFInterfaceRoutesRead(d, m)
}

func resourceCudaWAFInterfaceRoutesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFInterfaceRoutesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/interface-routes/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFInterfaceRoutesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFInterfaceRoutesRead(d, m)
}

func resourceCudaWAFInterfaceRoutesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/interface-routes/"
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

func hydrateBarracudaWAFInterfaceRoutesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"ip-address": d.Get("ip_address").(string),
		"comments":   d.Get("comments").(string),
		"interface":  d.Get("interface").(string),
		"ip-version": d.Get("ip_version").(string),
		"netmask":    d.Get("netmask").(string),
		"vsite":      d.Get("vsite").(string),
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
