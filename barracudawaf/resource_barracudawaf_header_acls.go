package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFHeaderAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFHeaderAclsCreate,
		Read:   resourceCudaWAFHeaderAclsRead,
		Update: resourceCudaWAFHeaderAclsUpdate,
		Delete: resourceCudaWAFHeaderAclsDelete,

		Schema: map[string]*schema.Schema{
			"comments":                    {Type: schema.TypeString, Optional: true},
			"max_header_value_length":     {Type: schema.TypeString, Optional: true},
			"header_name":                 {Type: schema.TypeString, Required: true},
			"blocked_attack_types":        {Type: schema.TypeString, Optional: true},
			"denied_metachars":            {Type: schema.TypeString, Optional: true},
			"mode":                        {Type: schema.TypeString, Optional: true},
			"custom_blocked_attack_types": {Type: schema.TypeString, Optional: true},
			"exception_patterns":          {Type: schema.TypeString, Optional: true},
			"status":                      {Type: schema.TypeString, Optional: true},
			"name":                        {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadHeaderAcls(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comments":                    d.Get("comments").(string),
		"max-header-value-length":     d.Get("max_header_value_length").(string),
		"header-name":                 d.Get("header_name").(string),
		"blocked-attack-types":        d.Get("blocked_attack_types").(string),
		"denied-metachars":            d.Get("denied_metachars").(string),
		"mode":                        d.Get("mode").(string),
		"custom-blocked-attack-types": d.Get("custom_blocked_attack_types").(string),
		"exception-patterns":          d.Get("exception_patterns").(string),
		"status":                      d.Get("status").(string),
		"name":                        d.Get("name").(string),
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

func resourceCudaWAFHeaderAclsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/header-acls"
	resourceCreateResponseError := makeRestAPIPayloadHeaderAcls(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFHeaderAclsRead(d, m)
}

func resourceCudaWAFHeaderAclsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFHeaderAclsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/header-acls/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadHeaderAcls(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFHeaderAclsRead(d, m)
}

func resourceCudaWAFHeaderAclsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/header-acls/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadHeaderAcls(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
