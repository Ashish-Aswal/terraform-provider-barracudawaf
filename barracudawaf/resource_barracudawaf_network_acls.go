package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNetworkAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNetworkAclsCreate,
		Read:   resourceCudaWAFNetworkAclsRead,
		Update: resourceCudaWAFNetworkAclsUpdate,
		Delete: resourceCudaWAFNetworkAclsDelete,

		Schema: map[string]*schema.Schema{
			"interface":                 {Type: schema.TypeString, Optional: true},
			"ip_version":                {Type: schema.TypeString, Optional: true},
			"action":                    {Type: schema.TypeString, Optional: true},
			"comments":                  {Type: schema.TypeString, Optional: true},
			"destination_port":          {Type: schema.TypeString, Optional: true},
			"source_address":            {Type: schema.TypeString, Optional: true},
			"ipv6_source_address":       {Type: schema.TypeString, Optional: true},
			"source_netmask":            {Type: schema.TypeString, Optional: true},
			"ipv6_source_netmask":       {Type: schema.TypeString, Optional: true},
			"icmp_response":             {Type: schema.TypeString, Optional: true},
			"enable_logging":            {Type: schema.TypeString, Optional: true},
			"max_connections":           {Type: schema.TypeString, Optional: true},
			"max_half_open_connections": {Type: schema.TypeString, Optional: true},
			"name":                      {Type: schema.TypeString, Required: true},
			"priority":                  {Type: schema.TypeString, Required: true},
			"protocol":                  {Type: schema.TypeString, Optional: true},
			"source_port":               {Type: schema.TypeString, Optional: true},
			"status":                    {Type: schema.TypeString, Optional: true},
			"destination_address":       {Type: schema.TypeString, Optional: true},
			"ipv6_destination_address":  {Type: schema.TypeString, Optional: true},
			"destination_netmask":       {Type: schema.TypeString, Optional: true},
			"ipv6_destination_netmask":  {Type: schema.TypeString, Optional: true},
			"vsite":                     {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFNetworkAclsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/network-acls"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNetworkAclsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFNetworkAclsRead(d, m)
}

func resourceCudaWAFNetworkAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNetworkAclsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/network-acls/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNetworkAclsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFNetworkAclsRead(d, m)
}

func resourceCudaWAFNetworkAclsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/network-acls/"
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

func hydrateBarracudaWAFNetworkAclsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"interface":                 d.Get("interface").(string),
		"ip-version":                d.Get("ip_version").(string),
		"action":                    d.Get("action").(string),
		"comments":                  d.Get("comments").(string),
		"destination-port":          d.Get("destination_port").(string),
		"source-address":            d.Get("source_address").(string),
		"ipv6-source-address":       d.Get("ipv6_source_address").(string),
		"source-netmask":            d.Get("source_netmask").(string),
		"ipv6-source-netmask":       d.Get("ipv6_source_netmask").(string),
		"icmp-response":             d.Get("icmp_response").(string),
		"enable-logging":            d.Get("enable_logging").(string),
		"max-connections":           d.Get("max_connections").(string),
		"max-half-open-connections": d.Get("max_half_open_connections").(string),
		"name":                      d.Get("name").(string),
		"priority":                  d.Get("priority").(string),
		"protocol":                  d.Get("protocol").(string),
		"source-port":               d.Get("source_port").(string),
		"status":                    d.Get("status").(string),
		"destination-address":       d.Get("destination_address").(string),
		"ipv6-destination-address":  d.Get("ipv6_destination_address").(string),
		"destination-netmask":       d.Get("destination_netmask").(string),
		"ipv6-destination-netmask":  d.Get("ipv6_destination_netmask").(string),
		"vsite":                     d.Get("vsite").(string),
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
