package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServerCreate,
		Read:   resourceCudaWAFServerRead,
		Update: resourceCudaWAFServerUpdate,
		Delete: resourceCudaWAFServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadServer(d *schema.ResourceData, m interface{}, resourceOperation string, resourceEndpoint string) error {

	//build Payload for the resource
	resourcePayload := map[string]string{
		"name":            d.Get("name").(string),
		"ip-address":      d.Get("ip_address").(string),
		"identifier":      d.Get("identifier").(string),
		"address-version": d.Get("address_version").(string),
		"status":          d.Get("status").(string),
		"port":            d.Get("port").(string),
		"comments":        d.Get("comments").(string),
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

func resourceCudaWAFServerCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/servers"
	resourceCreateResponseError := makeRestAPIPayloadServer(d, m, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFServerRead(d, m)
}

func resourceCudaWAFServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCudaWAFServerRead(d, m)
}

func resourceCudaWAFServerDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceEndpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/servers/" + name
	resourceDeleteResponseError := makeRestAPIPayloadServer(d, m, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
