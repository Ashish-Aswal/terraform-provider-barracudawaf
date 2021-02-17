package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGlobalAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGlobalAclsCreate,
		Read:   resourceCudaWAFGlobalAclsRead,
		Update: resourceCudaWAFGlobalAclsUpdate,
		Delete: resourceCudaWAFGlobalAclsDelete,

		Schema: map[string]*schema.Schema{
			"action":                  {Type: schema.TypeString, Optional: true},
			"comments":                {Type: schema.TypeString, Optional: true},
			"deny_response":           {Type: schema.TypeString, Optional: true},
			"extended_match":          {Type: schema.TypeString, Optional: true},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true},
			"follow_up_action":        {Type: schema.TypeString, Optional: true},
			"follow_up_action_time":   {Type: schema.TypeString, Optional: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"redirect_url":            {Type: schema.TypeString, Optional: true},
			"response_page":           {Type: schema.TypeString, Optional: true},
			"enable":                  {Type: schema.TypeString, Optional: true},
			"url":                     {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadGlobalAcls(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"action":                  d.Get("action").(string),
		"comments":                d.Get("comments").(string),
		"deny-response":           d.Get("deny_response").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"follow-up-action":        d.Get("follow_up_action").(string),
		"follow-up-action-time":   d.Get("follow_up_action_time").(string),
		"name":                    d.Get("name").(string),
		"redirect-url":            d.Get("redirect_url").(string),
		"response-page":           d.Get("response_page").(string),
		"enable":                  d.Get("enable").(string),
		"url":                     d.Get("url").(string),
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

func resourceCudaWAFGlobalAclsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/global-acls"
	resourceCreateResponseError := makeRestAPIPayloadGlobalAcls(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFGlobalAclsRead(d, m)
}

func resourceCudaWAFGlobalAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGlobalAclsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/global-acls/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadGlobalAcls(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFGlobalAclsRead(d, m)
}

func resourceCudaWAFGlobalAclsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/security-policies/" + d.Get("parent.0").(string) + "/global-acls/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadGlobalAcls(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
