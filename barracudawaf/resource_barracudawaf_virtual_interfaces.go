package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFVirtualInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFVirtualInterfacesCreate,
		Read:   resourceCudaWAFVirtualInterfacesRead,
		Update: resourceCudaWAFVirtualInterfacesUpdate,
		Delete: resourceCudaWAFVirtualInterfacesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address":            {Type: schema.TypeString, Required: true},
			"comments":              {Type: schema.TypeString, Optional: true},
			"interface":             {Type: schema.TypeString, Required: true},
			"ip_version":            {Type: schema.TypeString, Optional: true},
			"netmask":               {Type: schema.TypeString, Required: true},
			"virtual_ip_service_id": {Type: schema.TypeString, Optional: true},
			"vsite":                 {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFVirtualInterfacesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/virtual-interfaces"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFVirtualInterfacesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFVirtualInterfacesRead(d, m)
}

func resourceCudaWAFVirtualInterfacesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/virtual-interfaces"
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

func resourceCudaWAFVirtualInterfacesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/virtual-interfaces/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFVirtualInterfacesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFVirtualInterfacesRead(d, m)
}

func resourceCudaWAFVirtualInterfacesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/virtual-interfaces/"
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

func hydrateBarracudaWAFVirtualInterfacesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"ip-address":            d.Get("ip_address").(string),
		"comments":              d.Get("comments").(string),
		"interface":             d.Get("interface").(string),
		"ip-version":            d.Get("ip_version").(string),
		"netmask":               d.Get("netmask").(string),
		"virtual-ip-service-id": d.Get("virtual_ip_service_id").(string),
		"vsite":                 d.Get("vsite").(string),
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
