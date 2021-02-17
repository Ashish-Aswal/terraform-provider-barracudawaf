package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSessionIdentifiers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSessionIdentifiersCreate,
		Read:   resourceCudaWAFSessionIdentifiersRead,
		Update: resourceCudaWAFSessionIdentifiersUpdate,
		Delete: resourceCudaWAFSessionIdentifiersDelete,

		Schema: map[string]*schema.Schema{
			"name":            {Type: schema.TypeString, Required: true},
			"token_name":      {Type: schema.TypeString, Required: true},
			"token_type":      {Type: schema.TypeString, Required: true},
			"end_delimiter":   {Type: schema.TypeString, Optional: true},
			"start_delimiter": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadSessionIdentifiers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":            d.Get("name").(string),
		"token-name":      d.Get("token_name").(string),
		"token-type":      d.Get("token_type").(string),
		"end-delimiter":   d.Get("end_delimiter").(string),
		"start-delimiter": d.Get("start_delimiter").(string),
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

func resourceCudaWAFSessionIdentifiersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/session-identifiers"
	resourceCreateResponseError := makeRestAPIPayloadSessionIdentifiers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSessionIdentifiersRead(d, m)
}

func resourceCudaWAFSessionIdentifiersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSessionIdentifiersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/session-identifiers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSessionIdentifiers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSessionIdentifiersRead(d, m)
}

func resourceCudaWAFSessionIdentifiersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/session-identifiers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSessionIdentifiers(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
