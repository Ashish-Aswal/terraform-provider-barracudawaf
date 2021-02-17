package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFParameterProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFParameterProfilesCreate,
		Read:   resourceCudaWAFParameterProfilesRead,
		Update: resourceCudaWAFParameterProfilesUpdate,
		Delete: resourceCudaWAFParameterProfilesDelete,

		Schema: map[string]*schema.Schema{
			"allowed_metachars":             {Type: schema.TypeString, Optional: true},
			"base64_decode_parameter_value": {Type: schema.TypeString, Optional: true},
			"allowed_file_upload_type":      {Type: schema.TypeString, Optional: true},
			"comments":                      {Type: schema.TypeString, Optional: true},
			"custom_parameter_class":        {Type: schema.TypeString, Optional: true},
			"exception_patterns":            {Type: schema.TypeString, Optional: true},
			"file_upload_extensions":        {Type: schema.TypeString, Optional: true},
			"file_upload_mime_types":        {Type: schema.TypeString, Optional: true},
			"ignore":                        {Type: schema.TypeString, Optional: true},
			"maximum_instances":             {Type: schema.TypeString, Optional: true},
			"max_value_length":              {Type: schema.TypeString, Optional: true},
			"parameter":                     {Type: schema.TypeString, Required: true},
			"parameter_class":               {Type: schema.TypeString, Required: true},
			"required":                      {Type: schema.TypeString, Optional: true},
			"status":                        {Type: schema.TypeString, Optional: true},
			"type":                          {Type: schema.TypeString, Optional: true},
			"validate_parameter_name":       {Type: schema.TypeString, Optional: true},
			"values":                        {Type: schema.TypeString, Optional: true},
			"name":                          {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadParameterProfiles(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"allowed-metachars":             d.Get("allowed_metachars").(string),
		"base64-decode-parameter-value": d.Get("base64_decode_parameter_value").(string),
		"allowed-file-upload-type":      d.Get("allowed_file_upload_type").(string),
		"comments":                      d.Get("comments").(string),
		"custom-parameter-class":        d.Get("custom_parameter_class").(string),
		"exception-patterns":            d.Get("exception_patterns").(string),
		"file-upload-extensions":        d.Get("file_upload_extensions").(string),
		"file-upload-mime-types":        d.Get("file_upload_mime_types").(string),
		"ignore":                        d.Get("ignore").(string),
		"maximum-instances":             d.Get("maximum_instances").(string),
		"max-value-length":              d.Get("max_value_length").(string),
		"parameter":                     d.Get("parameter").(string),
		"parameter-class":               d.Get("parameter_class").(string),
		"required":                      d.Get("required").(string),
		"status":                        d.Get("status").(string),
		"type":                          d.Get("type").(string),
		"validate-parameter-name":       d.Get("validate_parameter_name").(string),
		"values":                        d.Get("values").(string),
		"name":                          d.Get("name").(string),
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

func resourceCudaWAFParameterProfilesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-profiles/" + d.Get("parent.1").(string) + "/parameter-profiles"
	resourceCreateResponseError := makeRestAPIPayloadParameterProfiles(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFParameterProfilesRead(d, m)
}

func resourceCudaWAFParameterProfilesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFParameterProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-profiles/" + d.Get("parent.1").(string) + "/parameter-profiles/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadParameterProfiles(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFParameterProfilesRead(d, m)
}

func resourceCudaWAFParameterProfilesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-profiles/" + d.Get("parent.1").(string) + "/parameter-profiles/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadParameterProfiles(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
