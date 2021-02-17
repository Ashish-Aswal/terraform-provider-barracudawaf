package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadAutoSystemAcls(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
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

func resourceCudaWAFAutoSystemAclsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/auto-system-acls"
	resourceCreateResponseError := makeRestAPIPayloadAutoSystemAcls(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFAutoSystemAclsRead(d, m)
}

func resourceCudaWAFAutoSystemAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAutoSystemAclsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/auto-system-acls/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadAutoSystemAcls(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFAutoSystemAclsRead(d, m)
}

func resourceCudaWAFAutoSystemAclsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/auto-system-acls/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadAutoSystemAcls(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
