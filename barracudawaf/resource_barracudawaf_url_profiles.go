package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlProfilesCreate,
		Read:   resourceCudaWAFUrlProfilesRead,
		Update: resourceCudaWAFUrlProfilesUpdate,
		Delete: resourceCudaWAFUrlProfilesDelete,

		Schema: map[string]*schema.Schema{
			"allowed_content_types":         {Type: schema.TypeString, Optional: true},
			"allowed_methods":               {Type: schema.TypeString, Optional: true},
			"custom_blocked_attack_types":   {Type: schema.TypeString, Optional: true},
			"comment":                       {Type: schema.TypeString, Optional: true},
			"display_name":                  {Type: schema.TypeString, Optional: true},
			"exception_patterns":            {Type: schema.TypeString, Optional: true},
			"extended_match":                {Type: schema.TypeString, Optional: true},
			"extended_match_sequence":       {Type: schema.TypeString, Optional: true},
			"hidden_parameter_protection":   {Type: schema.TypeString, Optional: true},
			"blocked_attack_types":          {Type: schema.TypeString, Optional: true},
			"max_content_length":            {Type: schema.TypeString, Optional: true},
			"maximum_parameter_name_length": {Type: schema.TypeString, Optional: true},
			"maximum_upload_files":          {Type: schema.TypeString, Optional: true},
			"minimum_form_fill_time":        {Type: schema.TypeString, Optional: true},
			"name":                          {Type: schema.TypeString, Required: true},
			"csrf_prevention":               {Type: schema.TypeString, Optional: true},
			"allow_query_string":            {Type: schema.TypeString, Optional: true},
			"referrers_for_the_url_profile": {Type: schema.TypeString, Optional: true},
			"mode":                          {Type: schema.TypeString, Optional: true},
			"status":                        {Type: schema.TypeString, Optional: true},
			"url":                           {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadUrlProfiles(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"allowed-content-types":         d.Get("allowed_content_types").(string),
		"allowed-methods":               d.Get("allowed_methods").(string),
		"custom-blocked-attack-types":   d.Get("custom_blocked_attack_types").(string),
		"comment":                       d.Get("comment").(string),
		"display-name":                  d.Get("display_name").(string),
		"exception-patterns":            d.Get("exception_patterns").(string),
		"extended-match":                d.Get("extended_match").(string),
		"extended-match-sequence":       d.Get("extended_match_sequence").(string),
		"hidden-parameter-protection":   d.Get("hidden_parameter_protection").(string),
		"blocked-attack-types":          d.Get("blocked_attack_types").(string),
		"max-content-length":            d.Get("max_content_length").(string),
		"maximum-parameter-name-length": d.Get("maximum_parameter_name_length").(string),
		"maximum-upload-files":          d.Get("maximum_upload_files").(string),
		"minimum-form-fill-time":        d.Get("minimum_form_fill_time").(string),
		"name":                          d.Get("name").(string),
		"csrf-prevention":               d.Get("csrf_prevention").(string),
		"allow-query-string":            d.Get("allow_query_string").(string),
		"referrers-for-the-url-profile": d.Get("referrers_for_the_url_profile").(string),
		"mode":                          d.Get("mode").(string),
		"status":                        d.Get("status").(string),
		"url":                           d.Get("url").(string),
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

func resourceCudaWAFUrlProfilesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-profiles"
	resourceCreateResponseError := makeRestAPIPayloadUrlProfiles(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFUrlProfilesRead(d, m)
}

func resourceCudaWAFUrlProfilesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUrlProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-profiles/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadUrlProfiles(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFUrlProfilesRead(d, m)
}

func resourceCudaWAFUrlProfilesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-profiles/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadUrlProfiles(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
