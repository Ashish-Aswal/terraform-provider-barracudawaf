package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSamlServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSamlServicesCreate,
		Read:   resourceCudaWAFSamlServicesRead,
		Update: resourceCudaWAFSamlServicesUpdate,
		Delete: resourceCudaWAFSamlServicesDelete,

		Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}},
	}
}

func resourceCudaWAFSamlServicesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/saml-services"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSamlServicesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSamlServicesRead(d, m)
}

func resourceCudaWAFSamlServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSamlServicesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/saml-services/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSamlServicesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSamlServicesRead(d, m)
}

func resourceCudaWAFSamlServicesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/saml-services/"
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

func hydrateBarracudaWAFSamlServicesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{"name": d.Get("name").(string)}

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
