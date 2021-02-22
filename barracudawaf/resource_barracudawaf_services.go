package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServicesCreate,
		Read:   resourceCudaWAFServicesRead,
		Update: resourceCudaWAFServicesUpdate,
		Delete: resourceCudaWAFServicesDelete,

		Schema: map[string]*schema.Schema{
			"address_version":     {Type: schema.TypeString, Optional: true},
			"dps_enabled":         {Type: schema.TypeString, Optional: true},
			"mask":                {Type: schema.TypeString, Optional: true},
			"session_timeout":     {Type: schema.TypeString, Optional: true},
			"linked_service_name": {Type: schema.TypeString, Optional: true},
			"enable_access_logs":  {Type: schema.TypeString, Optional: true},
			"app_id":              {Type: schema.TypeString, Optional: true},
			"comments":            {Type: schema.TypeString, Optional: true},
			"group":               {Type: schema.TypeString, Optional: true},
			"service_id":          {Type: schema.TypeString, Optional: true},
			"ip_address":          {Type: schema.TypeString, Optional: true},
			"cloud_ip_select":     {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"port":                {Type: schema.TypeString, Optional: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"type":                {Type: schema.TypeString, Optional: true},
			"vsite":               {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFServicesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServicesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFServicesRead(d, m)
}

func resourceCudaWAFServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServicesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServicesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFServicesRead(d, m)
}

func resourceCudaWAFServicesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/"
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

func hydrateBarracudaWAFServicesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"address-version":     d.Get("address_version").(string),
		"dps-enabled":         d.Get("dps_enabled").(string),
		"mask":                d.Get("mask").(string),
		"session-timeout":     d.Get("session_timeout").(string),
		"linked-service-name": d.Get("linked_service_name").(string),
		"enable-access-logs":  d.Get("enable_access_logs").(string),
		"app-id":              d.Get("app_id").(string),
		"comments":            d.Get("comments").(string),
		"group":               d.Get("group").(string),
		"service-id":          d.Get("service_id").(string),
		"ip-address":          d.Get("ip_address").(string),
		"cloud-ip-select":     d.Get("cloud_ip_select").(string),
		"name":                d.Get("name").(string),
		"port":                d.Get("port").(string),
		"status":              d.Get("status").(string),
		"type":                d.Get("type").(string),
		"vsite":               d.Get("vsite").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"address-version", "group", "vsite"}
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
