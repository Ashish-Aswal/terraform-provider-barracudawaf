package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlEncryptionRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlEncryptionRulesCreate,
		Read:   resourceCudaWAFUrlEncryptionRulesRead,
		Update: resourceCudaWAFUrlEncryptionRulesUpdate,
		Delete: resourceCudaWAFUrlEncryptionRulesDelete,

		Schema: map[string]*schema.Schema{
			"allow_unencrypted_requests": {Type: schema.TypeString, Optional: true},
			"exclude_urls":               {Type: schema.TypeString, Optional: true},
			"host":                       {Type: schema.TypeString, Required: true},
			"name":                       {Type: schema.TypeString, Required: true},
			"status":                     {Type: schema.TypeString, Optional: true},
			"url":                        {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadUrlEncryptionRules(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"allow-unencrypted-requests": d.Get("allow_unencrypted_requests").(string),
		"exclude-urls":               d.Get("exclude_urls").(string),
		"host":                       d.Get("host").(string),
		"name":                       d.Get("name").(string),
		"status":                     d.Get("status").(string),
		"url":                        d.Get("url").(string),
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

func resourceCudaWAFUrlEncryptionRulesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules"
	resourceCreateResponseError := makeRestAPIPayloadUrlEncryptionRules(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFUrlEncryptionRulesRead(d, m)
}

func resourceCudaWAFUrlEncryptionRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUrlEncryptionRulesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadUrlEncryptionRules(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFUrlEncryptionRulesRead(d, m)
}

func resourceCudaWAFUrlEncryptionRulesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-encryption-rules/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadUrlEncryptionRules(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
