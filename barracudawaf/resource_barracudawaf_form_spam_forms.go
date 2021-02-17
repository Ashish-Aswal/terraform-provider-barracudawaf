package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFFormSpamForms() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFFormSpamFormsCreate,
		Read:   resourceCudaWAFFormSpamFormsRead,
		Update: resourceCudaWAFFormSpamFormsUpdate,
		Delete: resourceCudaWAFFormSpamFormsDelete,

		Schema: map[string]*schema.Schema{
			"name":                   {Type: schema.TypeString, Required: true},
			"created_by":             {Type: schema.TypeString, Optional: true},
			"status":                 {Type: schema.TypeString, Optional: true},
			"mode":                   {Type: schema.TypeString, Optional: true},
			"action_url":             {Type: schema.TypeString, Required: true},
			"minimum_form_fill_time": {Type: schema.TypeString, Optional: true},
			"parameter_name":         {Type: schema.TypeString, Optional: true},
			"parameter_class":        {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadFormSpamForms(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                   d.Get("name").(string),
		"created-by":             d.Get("created_by").(string),
		"status":                 d.Get("status").(string),
		"mode":                   d.Get("mode").(string),
		"action-url":             d.Get("action_url").(string),
		"minimum-form-fill-time": d.Get("minimum_form_fill_time").(string),
		"parameter-name":         d.Get("parameter_name").(string),
		"parameter-class":        d.Get("parameter_class").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"created-by", "action-url"}
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

func resourceCudaWAFFormSpamFormsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/form-spam-forms"
	resourceCreateResponseError := makeRestAPIPayloadFormSpamForms(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFFormSpamFormsRead(d, m)
}

func resourceCudaWAFFormSpamFormsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFFormSpamFormsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/form-spam-forms/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadFormSpamForms(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFFormSpamFormsRead(d, m)
}

func resourceCudaWAFFormSpamFormsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/form-spam-forms/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadFormSpamForms(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
