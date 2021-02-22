package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGeoipBlockedNetworks() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGeoipBlockedNetworksCreate,
		Read:   resourceCudaWAFGeoipBlockedNetworksRead,
		Update: resourceCudaWAFGeoipBlockedNetworksUpdate,
		Delete: resourceCudaWAFGeoipBlockedNetworksDelete,

		Schema: map[string]*schema.Schema{
			"comment":       {Type: schema.TypeString, Optional: true},
			"block_ip":      {Type: schema.TypeString, Required: true},
			"block_netmask": {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFGeoipBlockedNetworksCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/geoip-blocked-networks"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGeoipBlockedNetworksResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFGeoipBlockedNetworksRead(d, m)
}

func resourceCudaWAFGeoipBlockedNetworksRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGeoipBlockedNetworksUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/geoip-blocked-networks/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFGeoipBlockedNetworksResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFGeoipBlockedNetworksRead(d, m)
}

func resourceCudaWAFGeoipBlockedNetworksDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/geoip-blocked-networks/"
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

func hydrateBarracudaWAFGeoipBlockedNetworksResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comment":       d.Get("comment").(string),
		"block-ip":      d.Get("block_ip").(string),
		"block-netmask": d.Get("block_netmask").(string),
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
