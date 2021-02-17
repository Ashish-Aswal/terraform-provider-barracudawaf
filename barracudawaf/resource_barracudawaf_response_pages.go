package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFResponsePages() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFResponsePagesCreate,
		Read:   resourceCudaWAFResponsePagesRead,
		Update: resourceCudaWAFResponsePagesUpdate,
		Delete: resourceCudaWAFResponsePagesDelete,

		Schema: map[string]*schema.Schema{
			"name":        {Type: schema.TypeString, Required: true},
			"body":        {Type: schema.TypeString, Optional: true},
			"headers":     {Type: schema.TypeString, Optional: true},
			"status_code": {Type: schema.TypeString, Required: true},
			"type":        {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadResponsePages(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":        d.Get("name").(string),
		"body":        d.Get("body").(string),
		"headers":     d.Get("headers").(string),
		"status-code": d.Get("status_code").(string),
		"type":        d.Get("type").(string),
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

func resourceCudaWAFResponsePagesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/response-pages"
	resourceCreateResponseError := makeRestAPIPayloadResponsePages(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFResponsePagesRead(d, m)
}

func resourceCudaWAFResponsePagesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFResponsePagesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/response-pages/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadResponsePages(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFResponsePagesRead(d, m)
}

func resourceCudaWAFResponsePagesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/response-pages/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadResponsePages(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
