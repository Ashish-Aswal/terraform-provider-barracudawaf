package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSecureBrowsingPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSecureBrowsingPoliciesCreate,
		Read:   resourceCudaWAFSecureBrowsingPoliciesRead,
		Update: resourceCudaWAFSecureBrowsingPoliciesUpdate,
		Delete: resourceCudaWAFSecureBrowsingPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"comments":                {Type: schema.TypeString, Optional: true},
			"host":                    {Type: schema.TypeString, Optional: true},
			"extended_match":          {Type: schema.TypeString, Optional: true},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"credential_server":       {Type: schema.TypeString, Required: true},
			"status":                  {Type: schema.TypeString, Optional: true},
			"url":                     {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadSecureBrowsingPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":                d.Get("comments").(string),
		"host":                    d.Get("host").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"name":                    d.Get("name").(string),
		"credential-server":       d.Get("credential_server").(string),
		"status":                  d.Get("status").(string),
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

func resourceCudaWAFSecureBrowsingPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/secure-browsing-policies"
	resourceCreateResponseError := makeRestAPIPayloadSecureBrowsingPolicies(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSecureBrowsingPoliciesRead(d, m)
}

func resourceCudaWAFSecureBrowsingPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSecureBrowsingPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/secure-browsing-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSecureBrowsingPolicies(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSecureBrowsingPoliciesRead(d, m)
}

func resourceCudaWAFSecureBrowsingPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/secure-browsing-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSecureBrowsingPolicies(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
