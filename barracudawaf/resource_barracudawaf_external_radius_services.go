package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFExternalRadiusServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFExternalRadiusServicesCreate,
		Read:   resourceCudaWAFExternalRadiusServicesRead,
		Update: resourceCudaWAFExternalRadiusServicesUpdate,
		Delete: resourceCudaWAFExternalRadiusServicesDelete,

		Schema: map[string]*schema.Schema{
			"server_ip":     {Type: schema.TypeString, Required: true},
			"default_role":  {Type: schema.TypeString, Required: true},
			"name":          {Type: schema.TypeString, Required: true},
			"port":          {Type: schema.TypeString, Optional: true},
			"shared_secret": {Type: schema.TypeString, Optional: true},
			"timeout":       {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFExternalRadiusServicesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/external-radius-services"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFExternalRadiusServicesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFExternalRadiusServicesRead(d, m)
}

func resourceCudaWAFExternalRadiusServicesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/external-radius-services"
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

func resourceCudaWAFExternalRadiusServicesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/external-radius-services/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFExternalRadiusServicesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFExternalRadiusServicesRead(d, m)
}

func resourceCudaWAFExternalRadiusServicesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/external-radius-services/"
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

func hydrateBarracudaWAFExternalRadiusServicesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"server-ip":     d.Get("server_ip").(string),
		"default-role":  d.Get("default_role").(string),
		"name":          d.Get("name").(string),
		"port":          d.Get("port").(string),
		"shared-secret": d.Get("shared_secret").(string),
		"timeout":       d.Get("timeout").(string),
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
