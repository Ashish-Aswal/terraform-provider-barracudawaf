package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFDdosPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFDdosPoliciesCreate,
		Read:   resourceCudaWAFDdosPoliciesRead,
		Update: resourceCudaWAFDdosPoliciesUpdate,
		Delete: resourceCudaWAFDdosPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"evaluate_clients":        {Type: schema.TypeString, Optional: true},
			"comments":                {Type: schema.TypeString, Optional: true},
			"enforce_captcha":         {Type: schema.TypeString, Optional: true},
			"extended_match":          {Type: schema.TypeString, Optional: true},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true},
			"host":                    {Type: schema.TypeString, Optional: true},
			"expiry_time":             {Type: schema.TypeString, Optional: true},
			"mouse_check":             {Type: schema.TypeString, Optional: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"num_captcha_tries":       {Type: schema.TypeString, Optional: true},
			"num_unanswered_captcha":  {Type: schema.TypeString, Optional: true},
			"url":                     {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadDdosPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"evaluate-clients":        d.Get("evaluate_clients").(string),
		"comments":                d.Get("comments").(string),
		"enforce-captcha":         d.Get("enforce_captcha").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"host":                    d.Get("host").(string),
		"expiry-time":             d.Get("expiry_time").(string),
		"mouse-check":             d.Get("mouse_check").(string),
		"name":                    d.Get("name").(string),
		"num-captcha-tries":       d.Get("num_captcha_tries").(string),
		"num-unanswered-captcha":  d.Get("num_unanswered_captcha").(string),
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

func resourceCudaWAFDdosPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/ddos-policies"
	resourceCreateResponseError := makeRestAPIPayloadDdosPolicies(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFDdosPoliciesRead(d, m)
}

func resourceCudaWAFDdosPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFDdosPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/ddos-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadDdosPolicies(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFDdosPoliciesRead(d, m)
}

func resourceCudaWAFDdosPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/ddos-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadDdosPolicies(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
