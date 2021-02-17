package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFModuleLogLevels() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFModuleLogLevelsCreate,
		Read:   resourceCudaWAFModuleLogLevelsRead,
		Update: resourceCudaWAFModuleLogLevelsUpdate,
		Delete: resourceCudaWAFModuleLogLevelsDelete,

		Schema: map[string]*schema.Schema{
			"log_level": {Type: schema.TypeString, Required: true},
			"comments":  {Type: schema.TypeString, Optional: true},
			"name":      {Type: schema.TypeString, Required: true},
			"module":    {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadModuleLogLevels(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"log-level": d.Get("log_level").(string),
		"comments":  d.Get("comments").(string),
		"name":      d.Get("name").(string),
		"module":    d.Get("module").(string),
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

func resourceCudaWAFModuleLogLevelsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/module-log-levels"
	resourceCreateResponseError := makeRestAPIPayloadModuleLogLevels(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFModuleLogLevelsRead(d, m)
}

func resourceCudaWAFModuleLogLevelsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFModuleLogLevelsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/module-log-levels/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadModuleLogLevels(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFModuleLogLevelsRead(d, m)
}

func resourceCudaWAFModuleLogLevelsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/module-log-levels/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadModuleLogLevels(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
