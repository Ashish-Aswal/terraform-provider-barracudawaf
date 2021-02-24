package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFStaticRoutes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFStaticRoutesCreate,
		Read:   resourceCudaWAFStaticRoutesRead,
		Update: resourceCudaWAFStaticRoutesUpdate,
		Delete: resourceCudaWAFStaticRoutesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {Type: schema.TypeString, Required: true},
			"comments":   {Type: schema.TypeString, Optional: true},
			"gateway":    {Type: schema.TypeString, Required: true},
			"ip_version": {Type: schema.TypeString, Optional: true},
			"netmask":    {Type: schema.TypeString, Required: true},
			"vsite":      {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFStaticRoutesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/static-routes"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFStaticRoutesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFStaticRoutesRead(d, m)
}

func resourceCudaWAFStaticRoutesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/static-routes"
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

func resourceCudaWAFStaticRoutesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/static-routes/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFStaticRoutesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFStaticRoutesRead(d, m)
}

func resourceCudaWAFStaticRoutesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/static-routes/"
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

func hydrateBarracudaWAFStaticRoutesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"ip-address": d.Get("ip_address").(string),
		"comments":   d.Get("comments").(string),
		"gateway":    d.Get("gateway").(string),
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
