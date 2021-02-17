package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCredentialServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCredentialServersCreate,
		Read:   resourceCudaWAFCredentialServersRead,
		Update: resourceCudaWAFCredentialServersUpdate,
		Delete: resourceCudaWAFCredentialServersDelete,

		Schema: map[string]*schema.Schema{
			"cache_expiry":         {Type: schema.TypeString, Optional: true},
			"cache_valid_sessions": {Type: schema.TypeString, Optional: true},
			"redirect_url":         {Type: schema.TypeString, Optional: true},
			"ip_address":           {Type: schema.TypeString, Required: true},
			"policy_name":          {Type: schema.TypeString, Required: true},
			"port":                 {Type: schema.TypeString, Optional: true},
			"armored_browser_type": {Type: schema.TypeString, Required: true},
			"name":                 {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadCredentialServers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"cache-expiry":         d.Get("cache_expiry").(string),
		"cache-valid-sessions": d.Get("cache_valid_sessions").(string),
		"redirect-url":         d.Get("redirect_url").(string),
		"ip-address":           d.Get("ip_address").(string),
		"policy-name":          d.Get("policy_name").(string),
		"port":                 d.Get("port").(string),
		"armored-browser-type": d.Get("armored_browser_type").(string),
		"name":                 d.Get("name").(string),
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

func resourceCudaWAFCredentialServersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/credential-servers"
	resourceCreateResponseError := makeRestAPIPayloadCredentialServers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFCredentialServersRead(d, m)
}

func resourceCudaWAFCredentialServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFCredentialServersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/credential-servers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadCredentialServers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFCredentialServersRead(d, m)
}

func resourceCudaWAFCredentialServersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/credential-servers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadCredentialServers(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
