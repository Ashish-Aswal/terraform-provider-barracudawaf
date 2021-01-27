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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vsite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_access_logs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadService(d *schema.ResourceData, m interface{}, oper string, endpoint string) error {
	//Build Payload for the resource
	payload := map[string]string{
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

func resourceCudaWAFServiceCreate(d *schema.ResourceData, m interface{}) error {
	endpoint := "restapi/v3/services"
	err := makeRestAPIPayloadService(d, m, "POST", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourceCudaWAFServiceRead(d, m)
}

func resourceCudaWAFServiceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServiceUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	endpoint := "restapi/v3/services/" + name
	err := makeRestAPIPayloadService(d, m, "PUT", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourceCudaWAFServiceRead(d, m)
}

func resourceCudaWAFServiceDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	endpoint := "restapi/v3/services/" + name
	err := makeRestAPIPayloadService(d, m, "DELETE", endpoint)
	if err != nil {
		return fmt.Errorf("error occurred : %v", err)
	}
	return nil
}
