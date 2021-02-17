package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCustomIpBlocklist() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCustomIpBlocklistCreate,
		Read:   resourceCudaWAFCustomIpBlocklistRead,
		Update: resourceCudaWAFCustomIpBlocklistUpdate,
		Delete: resourceCudaWAFCustomIpBlocklistDelete,

		Schema: map[string]*schema.Schema{
			"blacklisted_ips":             {Type: schema.TypeString, Optional: true},
			"custom_ip_list":              {Type: schema.TypeString, Optional: true},
			"download_url":                {Type: schema.TypeString, Optional: true},
			"trusted_certificate":         {Type: schema.TypeString, Optional: true},
			"validate_server_certificate": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadCustomIpBlocklist(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"blacklisted-ips":             d.Get("blacklisted_ips").(string),
		"custom-ip-list":              d.Get("custom_ip_list").(string),
		"download-url":                d.Get("download_url").(string),
		"trusted-certificate":         d.Get("trusted_certificate").(string),
		"validate-server-certificate": d.Get("validate_server_certificate").(string),
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

func resourceCudaWAFCustomIpBlocklistCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/custom-ip-blocklist"
	resourceCreateResponseError := makeRestAPIPayloadCustomIpBlocklist(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFCustomIpBlocklistRead(d, m)
}

func resourceCudaWAFCustomIpBlocklistRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFCustomIpBlocklistUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/custom-ip-blocklist/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadCustomIpBlocklist(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFCustomIpBlocklistRead(d, m)
}

func resourceCudaWAFCustomIpBlocklistDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/custom-ip-blocklist/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadCustomIpBlocklist(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
