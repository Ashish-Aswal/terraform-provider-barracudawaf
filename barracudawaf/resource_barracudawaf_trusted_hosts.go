package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFTrustedHosts() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFTrustedHostsCreate,
		Read:   resourceCudaWAFTrustedHostsRead,
		Update: resourceCudaWAFTrustedHostsUpdate,
		Delete: resourceCudaWAFTrustedHostsDelete,

		Schema: map[string]*schema.Schema{
			"ip_address":   {Type: schema.TypeString, Optional: true},
			"ipv6_address": {Type: schema.TypeString, Optional: true},
			"ipv6_mask":    {Type: schema.TypeString, Optional: true},
			"mask":         {Type: schema.TypeString, Optional: true},
			"comments":     {Type: schema.TypeString, Optional: true},
			"name":         {Type: schema.TypeString, Required: true},
			"version":      {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFTrustedHostsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/trusted-host-groups/" + d.Get("parent.0").(string) + "/trusted-hosts"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFTrustedHostsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFTrustedHostsRead(d, m)
}

func resourceCudaWAFTrustedHostsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFTrustedHostsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/trusted-host-groups/" + d.Get("parent.0").(string) + "/trusted-hosts/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFTrustedHostsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFTrustedHostsRead(d, m)
}

func resourceCudaWAFTrustedHostsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/trusted-host-groups/" + d.Get("parent.0").(string) + "/trusted-hosts/"
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

func hydrateBarracudaWAFTrustedHostsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"ip-address":   d.Get("ip_address").(string),
		"ipv6-address": d.Get("ipv6_address").(string),
		"ipv6-mask":    d.Get("ipv6_mask").(string),
		"mask":         d.Get("mask").(string),
		"comments":     d.Get("comments").(string),
		"name":         d.Get("name").(string),
		"version":      d.Get("version").(string),
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
