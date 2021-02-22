package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAutoSystemAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAutoSystemAclsCreate,
		Read:   resourceCudaWAFAutoSystemAclsRead,
		Update: resourceCudaWAFAutoSystemAclsUpdate,
		Delete: resourceCudaWAFAutoSystemAclsDelete,

		Schema: map[string]*schema.Schema{
			"interface":           {Type: schema.TypeString, Optional: true},
			"ip_version":          {Type: schema.TypeString, Optional: true},
			"acl_type":            {Type: schema.TypeString, Optional: true},
			"action":              {Type: schema.TypeString, Optional: true},
			"destination_port":    {Type: schema.TypeString, Required: true},
			"source_address":      {Type: schema.TypeString, Optional: true},
			"source_netmask":      {Type: schema.TypeString, Optional: true},
			"enable_logging":      {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"priority":            {Type: schema.TypeString, Required: true},
			"protocol":            {Type: schema.TypeString, Optional: true},
			"source_port":         {Type: schema.TypeString, Required: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"destination_address": {Type: schema.TypeString, Optional: true},
			"destination_netmask": {Type: schema.TypeString, Optional: true},
			"vsite":               {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFAutoSystemAclsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/auto-system-acls"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAutoSystemAclsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFAutoSystemAclsRead(d, m)
}

func resourceCudaWAFAutoSystemAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAutoSystemAclsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/auto-system-acls/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAutoSystemAclsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAutoSystemAclsRead(d, m)
}

func resourceCudaWAFAutoSystemAclsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/auto-system-acls/"
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

func hydrateBarracudaWAFAutoSystemAclsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"interface":           d.Get("interface").(string),
		"ip-version":          d.Get("ip_version").(string),
		"acl-type":            d.Get("acl_type").(string),
		"action":              d.Get("action").(string),
		"destination-port":    d.Get("destination_port").(string),
		"source-address":      d.Get("source_address").(string),
		"source-netmask":      d.Get("source_netmask").(string),
		"enable-logging":      d.Get("enable_logging").(string),
		"name":                d.Get("name").(string),
		"priority":            d.Get("priority").(string),
		"protocol":            d.Get("protocol").(string),
		"source-port":         d.Get("source_port").(string),
		"status":              d.Get("status").(string),
		"destination-address": d.Get("destination_address").(string),
		"destination-netmask": d.Get("destination_netmask").(string),
		"vsite":               d.Get("vsite").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{
			"interface",
			"ip-version",
			"acl-type",
			"action",
			"destination-port",
			"name",
			"priority",
			"protocol",
			"vsite",
		}
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
