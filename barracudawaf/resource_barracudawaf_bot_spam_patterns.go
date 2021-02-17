package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFBotSpamPatterns() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFBotSpamPatternsCreate,
		Read:   resourceCudaWAFBotSpamPatternsRead,
		Update: resourceCudaWAFBotSpamPatternsUpdate,
		Delete: resourceCudaWAFBotSpamPatternsDelete,

		Schema: map[string]*schema.Schema{
			"algorithm":      {Type: schema.TypeString, Optional: true},
			"case_sensitive": {Type: schema.TypeString, Optional: true},
			"description":    {Type: schema.TypeString, Optional: true},
			"mode":           {Type: schema.TypeString, Optional: true},
			"name":           {Type: schema.TypeString, Required: true},
			"regex":          {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadBotSpamPatterns(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"algorithm":      d.Get("algorithm").(string),
		"case-sensitive": d.Get("case_sensitive").(string),
		"description":    d.Get("description").(string),
		"mode":           d.Get("mode").(string),
		"name":           d.Get("name").(string),
		"regex":          d.Get("regex").(string),
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

func resourceCudaWAFBotSpamPatternsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns"
	resourceCreateResponseError := makeRestAPIPayloadBotSpamPatterns(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFBotSpamPatternsRead(d, m)
}

func resourceCudaWAFBotSpamPatternsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFBotSpamPatternsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadBotSpamPatterns(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFBotSpamPatternsRead(d, m)
}

func resourceCudaWAFBotSpamPatternsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/bot-spam-types/" + d.Get("parent.0").(string) + "/bot-spam-patterns/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadBotSpamPatterns(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
