package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFExternalRadiusServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFExternalRadiusServicesCreate,
		Read:   resourceCudaWAFExternalRadiusServicesRead,
		Update: resourceCudaWAFExternalRadiusServicesUpdate,
		Delete: resourceCudaWAFExternalRadiusServicesDelete,

		Schema: map[string]*schema.Schema{
			"server_ip":     {Type: schema.TypeString, Required: true},
			"default_role":  {Type: schema.TypeString, Required: true},
			"name":          {Type: schema.TypeString, Required: true},
			"port":          {Type: schema.TypeString, Optional: true},
			"shared_secret": {Type: schema.TypeString, Optional: true},
			"timeout":       {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadExternalRadiusServices(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"server-ip":     d.Get("server_ip").(string),
		"default-role":  d.Get("default_role").(string),
		"name":          d.Get("name").(string),
		"port":          d.Get("port").(string),
		"shared-secret": d.Get("shared_secret").(string),
		"timeout":       d.Get("timeout").(string),
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

func resourceCudaWAFExternalRadiusServicesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/external-radius-services"
	resourceCreateResponseError := makeRestAPIPayloadExternalRadiusServices(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFExternalRadiusServicesRead(d, m)
}

func resourceCudaWAFExternalRadiusServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFExternalRadiusServicesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/external-radius-services/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadExternalRadiusServices(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFExternalRadiusServicesRead(d, m)
}

func resourceCudaWAFExternalRadiusServicesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/external-radius-services/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadExternalRadiusServices(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
