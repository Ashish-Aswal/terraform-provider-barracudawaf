package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSystemCreate,
		Read:   resourceCudaWAFSystemRead,
		Update: resourceCudaWAFSystemUpdate,
		Delete: resourceCudaWAFSystemDelete,

		Schema: map[string]*schema.Schema{
			"attack_definition":             {Type: schema.TypeString, Optional: true},
			"device_name":                   {Type: schema.TypeString, Optional: true},
			"locale":                        {Type: schema.TypeString, Optional: true},
			"enable_multi_ip":               {Type: schema.TypeString, Optional: true},
			"firmware_version":              {Type: schema.TypeString, Optional: true},
			"model":                         {Type: schema.TypeString, Optional: true},
			"operation_mode":                {Type: schema.TypeString, Optional: true},
			"interface_for_system_services": {Type: schema.TypeString, Optional: true},
			"domain":                        {Type: schema.TypeString, Required: true},
			"hostname":                      {Type: schema.TypeString, Optional: true},
			"serial":                        {Type: schema.TypeString, Required: true},
			"time_zone":                     {Type: schema.TypeString, Optional: true},
			"enable_ipv6":                   {Type: schema.TypeString, Optional: true},
			"virus_definition":              {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadSystem(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"attack-definition":             d.Get("attack_definition").(string),
		"device-name":                   d.Get("device_name").(string),
		"locale":                        d.Get("locale").(string),
		"enable-multi-ip":               d.Get("enable_multi_ip").(string),
		"firmware-version":              d.Get("firmware_version").(string),
		"model":                         d.Get("model").(string),
		"operation-mode":                d.Get("operation_mode").(string),
		"interface-for-system-services": d.Get("interface_for_system_services").(string),
		"domain":                        d.Get("domain").(string),
		"hostname":                      d.Get("hostname").(string),
		"serial":                        d.Get("serial").(string),
		"time-zone":                     d.Get("time_zone").(string),
		"enable-ipv6":                   d.Get("enable_ipv6").(string),
		"virus-definition":              d.Get("virus_definition").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{
			"attack-definition",
			"firmware-version",
			"model",
			"serial",
			"virus-definition",
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

func resourceCudaWAFSystemCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/system"
	resourceCreateResponseError := makeRestAPIPayloadSystem(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSystemRead(d, m)
}

func resourceCudaWAFSystemRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSystemUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/system/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSystem(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSystemRead(d, m)
}

func resourceCudaWAFSystemDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/system/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSystem(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
