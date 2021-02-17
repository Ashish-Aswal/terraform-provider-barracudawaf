package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadServices(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"address-version", "group", "vsite"}
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

func resourceCudaWAFServicesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services"
	resourceCreateResponseError := makeRestAPIPayloadServices(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFServicesRead(d, m)
}

func resourceCudaWAFServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServicesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadServices(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFServicesRead(d, m)
}

func resourceCudaWAFServicesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadServices(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
