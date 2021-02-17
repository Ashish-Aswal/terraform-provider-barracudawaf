package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadNetworkAcls(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{}
		for item := range updatePayloadExceptions {
			delete(resourcePayload, updatePayloadExceptions[item])
		}
	}

	//sanitise the resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	//resourceUpdateData : cudaWAF reource URI update data
	resourceUpdateData := map[string]interface{}{
		"endpoint":  resourceEndpoint,
		"payload":   resourcePayload,
		"operation": resourceOperation,
		"name":      d.Get("name").(string),
	}

	//updateCudaWAFResourceObject : update cudaWAF resource object
	resourceUpdateStatus, resourceUpdateResponseBody := updateCudaWAFResourceObject(
		resourceUpdateData,
	)

	if resourceUpdateStatus == 200 || resourceUpdateStatus == 201 {
		if resourceOperation != "DELETE" {
			d.SetId(resourceUpdateResponseBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", resourceUpdateResponseBody["msg"])
	}

	return nil
}

func resourceCudaWAFNetworkAclsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/network-acls"
	resourceCreateResponseError := makeRestAPIPayloadNetworkAcls(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFNetworkAclsRead(d, m)
}

func resourceCudaWAFNetworkAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNetworkAclsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/network-acls/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadNetworkAcls(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFNetworkAclsRead(d, m)
}

func resourceCudaWAFNetworkAclsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/network-acls/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadNetworkAcls(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
