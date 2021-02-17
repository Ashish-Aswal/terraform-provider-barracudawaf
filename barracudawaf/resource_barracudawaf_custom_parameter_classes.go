package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCustomParameterClasses() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCustomParameterClassesCreate,
		Read:   resourceCudaWAFCustomParameterClassesRead,
		Update: resourceCudaWAFCustomParameterClassesUpdate,
		Delete: resourceCudaWAFCustomParameterClassesDelete,

		Schema: map[string]*schema.Schema{
			"name":                         {Type: schema.TypeString, Required: true},
			"custom_blocked_attack_types":  {Type: schema.TypeString, Optional: true},
			"custom_input_type_validation": {Type: schema.TypeString, Optional: true},
			"denied_metacharacters":        {Type: schema.TypeString, Optional: true},
			"input_type_validation":        {Type: schema.TypeString, Optional: true},
			"blocked_attack_types":         {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadCustomParameterClasses(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                         d.Get("name").(string),
		"custom-blocked-attack-types":  d.Get("custom_blocked_attack_types").(string),
		"custom-input-type-validation": d.Get("custom_input_type_validation").(string),
		"denied-metacharacters":        d.Get("denied_metacharacters").(string),
		"input-type-validation":        d.Get("input_type_validation").(string),
		"blocked-attack-types":         d.Get("blocked_attack_types").(string),
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

func resourceCudaWAFCustomParameterClassesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/custom-parameter-classes"
	resourceCreateResponseError := makeRestAPIPayloadCustomParameterClasses(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFCustomParameterClassesRead(d, m)
}

func resourceCudaWAFCustomParameterClassesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFCustomParameterClassesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/custom-parameter-classes/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadCustomParameterClasses(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFCustomParameterClassesRead(d, m)
}

func resourceCudaWAFCustomParameterClassesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/custom-parameter-classes/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadCustomParameterClasses(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
