package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGeoPools() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGeoPoolsCreate,
		Read:   resourceCudaWAFGeoPoolsRead,
		Update: resourceCudaWAFGeoPoolsUpdate,
		Delete: resourceCudaWAFGeoPoolsDelete,

		Schema: map[string]*schema.Schema{
			"region": {Type: schema.TypeString, Optional: true},
			"name":   {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFGeoPoolsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/geo-pools"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGeoPoolsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFGeoPoolsRead(d, m)
}

func resourceCudaWAFGeoPoolsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGeoPoolsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/geo-pools/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGeoPoolsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFGeoPoolsRead(d, m)
}

func resourceCudaWAFGeoPoolsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/geo-pools/"
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

func hydrateBarracudaWAFGeoPoolsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"region": d.Get("region").(string),
		"name":   d.Get("name").(string),
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
