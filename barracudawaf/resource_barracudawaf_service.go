package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFService() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServiceCreate,
		Read:   resourceCudaWAFServiceRead,
		Update: resourceCudaWAFServiceUpdate,
		Delete: resourceCudaWAFServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vsite": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_access_logs": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadService(d *schema.ResourceData, resourceOperation string, resourceEndpoint string) error {

	//build Payload for the resource
	resourcePayload := map[string]string{
		"name":               d.Get("name").(string),
		"port":               d.Get("port").(string),
		"type":               d.Get("type").(string),
		"mask":               d.Get("mask").(string),
		"vsite":              d.Get("vsite").(string),
		"group":              d.Get("group").(string),
		"status":             d.Get("status").(string),
		"comments":           d.Get("comments").(string),
		"ip-address":         d.Get("ip_address").(string),
		"certificate":        d.Get("certificate").(string),
		"address-version":    d.Get("address_version").(string),
		"enable-access-logs": d.Get("enable_access_logs").(string),
	}

	//sanitise the payload, removing empty keys
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	resourceUpdateData := map[string]interface{}{
		"endpoint":  resourceEndpoint,
		"payload":   resourcePayload,
		"operation": resourceOperation,
		"name":      d.Get("name").(string),
	}

	resourceUpdateStatus, resourceUpdateResponseBody := updateCudaWAFResourceObject(resourceUpdateData)

	if resourceUpdateStatus == 200 || resourceUpdateStatus == 201 {
		if resourceOperation != "DELETE" {
			d.SetId(resourceUpdateResponseBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", resourceUpdateResponseBody["msg"])
	}

	return nil
}

func resourceCudaWAFServiceCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := "restapi/v3/services"
	resourceCreateResponseError := makeRestAPIPayloadService(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFServiceRead(d, m)
}

func resourceCudaWAFServiceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServiceUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceEndpoint := "restapi/v3/services/" + name
	resourceUpdateResponseError := makeRestAPIPayloadService(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFServiceRead(d, m)
}

func resourceCudaWAFServiceDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceEndpoint := "restapi/v3/services/" + name
	resourceDeleteResponseError := makeRestAPIPayloadService(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
