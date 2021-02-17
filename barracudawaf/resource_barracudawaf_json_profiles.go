package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFJsonProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFJsonProfilesCreate,
		Read:   resourceCudaWAFJsonProfilesRead,
		Update: resourceCudaWAFJsonProfilesUpdate,
		Delete: resourceCudaWAFJsonProfilesDelete,

		Schema: map[string]*schema.Schema{
			"allowed_content_types": {Type: schema.TypeString, Optional: true},
			"host_match":            {Type: schema.TypeString, Required: true},
			"ignore_keys":           {Type: schema.TypeString, Optional: true},
			"method":                {Type: schema.TypeString, Required: true},
			"comment":               {Type: schema.TypeString, Optional: true},
			"name":                  {Type: schema.TypeString, Required: true},
			"mode":                  {Type: schema.TypeString, Optional: true},
			"status":                {Type: schema.TypeString, Optional: true},
			"url_match":             {Type: schema.TypeString, Required: true},
			"exception_patterns":    {Type: schema.TypeString, Optional: true},
			"json_policy":           {Type: schema.TypeString, Optional: true},
			"validate_key":          {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadJsonProfiles(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"allowed-content-types": d.Get("allowed_content_types").(string),
		"host_match":            d.Get("host_match").(string),
		"ignore-keys":           d.Get("ignore_keys").(string),
		"method":                d.Get("method").(string),
		"comment":               d.Get("comment").(string),
		"name":                  d.Get("name").(string),
		"mode":                  d.Get("mode").(string),
		"status":                d.Get("status").(string),
		"url_match":             d.Get("url_match").(string),
		"exception-patterns":    d.Get("exception_patterns").(string),
		"json_policy":           d.Get("json_policy").(string),
		"validate_key":          d.Get("validate_key").(string),
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

func resourceCudaWAFJsonProfilesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/json-profiles"
	resourceCreateResponseError := makeRestAPIPayloadJsonProfiles(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFJsonProfilesRead(d, m)
}

func resourceCudaWAFJsonProfilesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFJsonProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/json-profiles/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadJsonProfiles(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFJsonProfilesRead(d, m)
}

func resourceCudaWAFJsonProfilesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/json-profiles/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadJsonProfiles(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
