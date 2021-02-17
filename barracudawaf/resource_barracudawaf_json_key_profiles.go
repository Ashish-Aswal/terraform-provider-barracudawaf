package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFJsonKeyProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFJsonKeyProfilesCreate,
		Read:   resourceCudaWAFJsonKeyProfilesRead,
		Update: resourceCudaWAFJsonKeyProfilesUpdate,
		Delete: resourceCudaWAFJsonKeyProfilesDelete,

		Schema: map[string]*schema.Schema{
			"allow_null":                    {Type: schema.TypeString, Optional: true},
			"allowed_metachars":             {Type: schema.TypeString, Optional: true},
			"base64_decode_parameter_value": {Type: schema.TypeString, Optional: true},
			"comments":                      {Type: schema.TypeString, Optional: true},
			"value_class":                   {Type: schema.TypeString, Required: true},
			"exception_patterns":            {Type: schema.TypeString, Optional: true},
			"key":                           {Type: schema.TypeString, Required: true},
			"max_array_elements":            {Type: schema.TypeString, Optional: true},
			"max_length":                    {Type: schema.TypeString, Optional: true},
			"max_number_value":              {Type: schema.TypeString, Optional: true},
			"max_keys":                      {Type: schema.TypeString, Optional: true},
			"name":                          {Type: schema.TypeString, Required: true},
			"status":                        {Type: schema.TypeString, Optional: true},
			"validate_key":                  {Type: schema.TypeString, Optional: true},
			"value_type":                    {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadJsonKeyProfiles(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"allow-null":                    d.Get("allow_null").(string),
		"allowed-metachars":             d.Get("allowed_metachars").(string),
		"base64-decode-parameter-value": d.Get("base64_decode_parameter_value").(string),
		"comments":                      d.Get("comments").(string),
		"value-class":                   d.Get("value_class").(string),
		"exception-patterns":            d.Get("exception_patterns").(string),
		"key":                           d.Get("key").(string),
		"max-array-elements":            d.Get("max_array_elements").(string),
		"max-length":                    d.Get("max_length").(string),
		"max-number-value":              d.Get("max_number_value").(string),
		"max-keys":                      d.Get("max_keys").(string),
		"name":                          d.Get("name").(string),
		"status":                        d.Get("status").(string),
		"validate-key":                  d.Get("validate_key").(string),
		"value-type":                    d.Get("value_type").(string),
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

func resourceCudaWAFJsonKeyProfilesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/json-profiles/" + d.Get("parent.1").(string) + "/json-key-profiles"
	resourceCreateResponseError := makeRestAPIPayloadJsonKeyProfiles(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFJsonKeyProfilesRead(d, m)
}

func resourceCudaWAFJsonKeyProfilesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFJsonKeyProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/json-profiles/" + d.Get("parent.1").(string) + "/json-key-profiles/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadJsonKeyProfiles(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFJsonKeyProfilesRead(d, m)
}

func resourceCudaWAFJsonKeyProfilesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/json-profiles/" + d.Get("parent.1").(string) + "/json-key-profiles/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadJsonKeyProfiles(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
