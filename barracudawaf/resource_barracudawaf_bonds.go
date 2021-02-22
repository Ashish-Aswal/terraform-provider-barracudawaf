package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFBonds() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFBondsCreate,
		Read:   resourceCudaWAFBondsRead,
		Update: resourceCudaWAFBondsUpdate,
		Delete: resourceCudaWAFBondsDelete,

		Schema: map[string]*schema.Schema{
			"duplexity":  {Type: schema.TypeString, Optional: true},
			"name":       {Type: schema.TypeString, Required: true},
			"speed":      {Type: schema.TypeString, Optional: true},
			"bond_ports": {Type: schema.TypeString, Required: true},
			"min_link":   {Type: schema.TypeString, Optional: true},
			"mode":       {Type: schema.TypeString, Optional: true},
			"mtu":        {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFBondsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/bonds"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFBondsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFBondsRead(d, m)
}

func resourceCudaWAFBondsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFBondsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/bonds/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFBondsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFBondsRead(d, m)
}

func resourceCudaWAFBondsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/bonds/"
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

func hydrateBarracudaWAFBondsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"duplexity":  d.Get("duplexity").(string),
		"name":       d.Get("name").(string),
		"speed":      d.Get("speed").(string),
		"bond-ports": d.Get("bond_ports").(string),
		"min-link":   d.Get("min_link").(string),
		"mode":       d.Get("mode").(string),
		"mtu":        d.Get("mtu").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"name"}
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
