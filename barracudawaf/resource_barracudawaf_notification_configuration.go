package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNotificationConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNotificationConfigurationCreate,
		Read:   resourceCudaWAFNotificationConfigurationRead,
		Update: resourceCudaWAFNotificationConfigurationUpdate,
		Delete: resourceCudaWAFNotificationConfigurationDelete,

		Schema: map[string]*schema.Schema{"severity": {Type: schema.TypeString, Optional: true}},
	}
}

func makeRestAPIPayloadNotificationConfiguration(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{"severity": d.Get("severity").(string)}

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

func resourceCudaWAFNotificationConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/notification-configuration"
	resourceCreateResponseError := makeRestAPIPayloadNotificationConfiguration(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFNotificationConfigurationRead(d, m)
}

func resourceCudaWAFNotificationConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNotificationConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/notification-configuration/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadNotificationConfiguration(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFNotificationConfigurationRead(d, m)
}

func resourceCudaWAFNotificationConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/notification-configuration/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadNotificationConfiguration(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
