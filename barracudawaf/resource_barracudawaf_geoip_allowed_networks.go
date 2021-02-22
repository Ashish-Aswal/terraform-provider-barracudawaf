package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGeoipAllowedNetworks() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGeoipAllowedNetworksCreate,
		Read:   resourceCudaWAFGeoipAllowedNetworksRead,
		Update: resourceCudaWAFGeoipAllowedNetworksUpdate,
		Delete: resourceCudaWAFGeoipAllowedNetworksDelete,

		Schema: map[string]*schema.Schema{
			"comment":       {Type: schema.TypeString, Optional: true},
			"allow_ip":      {Type: schema.TypeString, Required: true},
			"allow_netmask": {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFGeoipAllowedNetworksCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/geoip-allowed-networks"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGeoipAllowedNetworksResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFGeoipAllowedNetworksRead(d, m)
}

func resourceCudaWAFGeoipAllowedNetworksRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGeoipAllowedNetworksUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/geoip-allowed-networks/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGeoipAllowedNetworksResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFGeoipAllowedNetworksRead(d, m)
}

func resourceCudaWAFGeoipAllowedNetworksDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/geoip-allowed-networks/"
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

func hydrateBarracudaWAFGeoipAllowedNetworksResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comment":       d.Get("comment").(string),
		"allow-ip":      d.Get("allow_ip").(string),
		"allow-netmask": d.Get("allow_netmask").(string),
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
