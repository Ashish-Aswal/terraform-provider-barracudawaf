package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAccessRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAccessRulesCreate,
		Read:   resourceCudaWAFAccessRulesRead,
		Update: resourceCudaWAFAccessRulesUpdate,
		Delete: resourceCudaWAFAccessRulesDelete,

		Schema: map[string]*schema.Schema{
			"attribute_names":  {Type: schema.TypeString, Optional: true},
			"attribute_values": {Type: schema.TypeString, Optional: true},
			"name":             {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadAccessRules(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"attribute-names":  d.Get("attribute_names").(string),
		"attribute-values": d.Get("attribute_values").(string),
		"name":             d.Get("name").(string),
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

func resourceCudaWAFAccessRulesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/access-rules"
	resourceCreateResponseError := makeRestAPIPayloadAccessRules(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFAccessRulesRead(d, m)
}

func resourceCudaWAFAccessRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAccessRulesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/access-rules/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadAccessRules(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFAccessRulesRead(d, m)
}

func resourceCudaWAFAccessRulesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/access-rules/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadAccessRules(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
