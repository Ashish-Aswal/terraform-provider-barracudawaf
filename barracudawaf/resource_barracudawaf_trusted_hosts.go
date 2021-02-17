package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadTrustedHosts(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"ip-address":   d.Get("ip_address").(string),
		"ipv6-address": d.Get("ipv6_address").(string),
		"ipv6-mask":    d.Get("ipv6_mask").(string),
		"mask":         d.Get("mask").(string),
		"comments":     d.Get("comments").(string),
		"name":         d.Get("name").(string),
		"version":      d.Get("version").(string),
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

func resourceCudaWAFTrustedHostsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-host-groups/" + d.Get("parent.0").(string) + "/trusted-hosts"
	resourceCreateResponseError := makeRestAPIPayloadTrustedHosts(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFTrustedHostsRead(d, m)
}

func resourceCudaWAFTrustedHostsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFTrustedHostsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-host-groups/" + d.Get("parent.0").(string) + "/trusted-hosts/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadTrustedHosts(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFTrustedHostsRead(d, m)
}

func resourceCudaWAFTrustedHostsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-host-groups/" + d.Get("parent.0").(string) + "/trusted-hosts/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadTrustedHosts(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
