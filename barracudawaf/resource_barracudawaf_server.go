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

func makeRestAPIPayloadServer(d *schema.ResourceData, m interface{}, oper string, endpoint string) error {
	payload := map[string]string{
		"name":            d.Get("name").(string),
		"ip-address":      d.Get("ip_address").(string),
		"identifier":      d.Get("identifier").(string),
		"address-version": d.Get("address_version").(string),
		"status":          d.Get("status").(string),
		"port":            d.Get("port").(string),
		"comments":        d.Get("comments").(string),
	}

	for key, value := range payload {
		if len(value) > 0 {
			continue
		} else {
			delete(payload, key)
		}
	}

	callData := map[string]interface{}{
		"endpoint":  endpoint,
		"payload":   payload,
		"operation": oper,
		"name":      d.Get("name").(string),
	}

	callStatus, callRespBody := doRestAPICall(callData)
	if callStatus == 200 || callStatus == 201 {
		if oper != "DELETE" {
			d.SetId(callRespBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", callRespBody["msg"])
	}

	return nil
}

func resourceCudaWAFServerCreate(d *schema.ResourceData, m interface{}) error {
	endpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/servers"
	err := makeRestAPIPayloadServer(d, m, "POST", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
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
	endpoint := "restapi/v3/services/" + d.Get("service_name").(string) + "/servers/" + name
	err := makeRestAPIPayloadServer(d, m, "DELETE", endpoint)
	if err != nil {
		return fmt.Errorf("error occurred : %v", err)
	}
	return nil
}
