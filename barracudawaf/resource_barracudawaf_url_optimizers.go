package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlOptimizers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlOptimizersCreate,
		Read:   resourceCudaWAFUrlOptimizersRead,
		Update: resourceCudaWAFUrlOptimizersUpdate,
		Delete: resourceCudaWAFUrlOptimizersDelete,

		Schema: map[string]*schema.Schema{
			"end_delimiter": {Type: schema.TypeString, Optional: true},
			"name":          {Type: schema.TypeString, Required: true},
			"start_token":   {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadUrlOptimizers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"end-delimiter": d.Get("end_delimiter").(string),
		"name":          d.Get("name").(string),
		"start-token":   d.Get("start_token").(string),
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

func resourceCudaWAFUrlOptimizersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-optimizers"
	resourceCreateResponseError := makeRestAPIPayloadUrlOptimizers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFUrlOptimizersRead(d, m)
}

func resourceCudaWAFUrlOptimizersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUrlOptimizersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-optimizers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadUrlOptimizers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFUrlOptimizersRead(d, m)
}

func resourceCudaWAFUrlOptimizersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-optimizers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadUrlOptimizers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
