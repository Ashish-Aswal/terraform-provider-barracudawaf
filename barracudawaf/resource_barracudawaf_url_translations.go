package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlTranslations() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlTranslationsCreate,
		Read:   resourceCudaWAFUrlTranslationsRead,
		Update: resourceCudaWAFUrlTranslationsUpdate,
		Delete: resourceCudaWAFUrlTranslationsDelete,

		Schema: map[string]*schema.Schema{
			"comments":       {Type: schema.TypeString, Optional: true},
			"outside_domain": {Type: schema.TypeString, Required: true},
			"outside_prefix": {Type: schema.TypeString, Required: true},
			"inside_domain":  {Type: schema.TypeString, Required: true},
			"inside_prefix":  {Type: schema.TypeString, Required: true},
			"name":           {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadUrlTranslations(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":       d.Get("comments").(string),
		"outside-domain": d.Get("outside_domain").(string),
		"outside-prefix": d.Get("outside_prefix").(string),
		"inside-domain":  d.Get("inside_domain").(string),
		"inside-prefix":  d.Get("inside_prefix").(string),
		"name":           d.Get("name").(string),
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

func resourceCudaWAFUrlTranslationsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-translations"
	resourceCreateResponseError := makeRestAPIPayloadUrlTranslations(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFUrlTranslationsRead(d, m)
}

func resourceCudaWAFUrlTranslationsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUrlTranslationsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-translations/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadUrlTranslations(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFUrlTranslationsRead(d, m)
}

func resourceCudaWAFUrlTranslationsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-translations/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadUrlTranslations(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
