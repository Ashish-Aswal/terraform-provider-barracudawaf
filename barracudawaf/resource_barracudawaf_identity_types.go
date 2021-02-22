package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFIdentityTypes() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFIdentityTypesCreate,
		Read:   resourceCudaWAFIdentityTypesRead,
		Update: resourceCudaWAFIdentityTypesUpdate,
		Delete: resourceCudaWAFIdentityTypesDelete,

		Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}},
	}
}

func resourceCudaWAFIdentityTypesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/identity-types"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFIdentityTypesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFIdentityTypesRead(d, m)
}

func resourceCudaWAFIdentityTypesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFIdentityTypesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/identity-types/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFIdentityTypesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFIdentityTypesRead(d, m)
}

func resourceCudaWAFIdentityTypesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/identity-types/"
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

func hydrateBarracudaWAFIdentityTypesResource(
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
