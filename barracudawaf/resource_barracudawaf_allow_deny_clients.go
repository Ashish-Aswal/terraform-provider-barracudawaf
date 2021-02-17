package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAllowDenyClients() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAllowDenyClientsCreate,
		Read:   resourceCudaWAFAllowDenyClientsRead,
		Update: resourceCudaWAFAllowDenyClientsUpdate,
		Delete: resourceCudaWAFAllowDenyClientsDelete,

		Schema: map[string]*schema.Schema{
			"name":                {Type: schema.TypeString, Required: true},
			"action":              {Type: schema.TypeString, Optional: true},
			"sequence":            {Type: schema.TypeString, Optional: true},
			"certificate_serial":  {Type: schema.TypeString, Optional: true},
			"common_name":         {Type: schema.TypeString, Optional: true},
			"country":             {Type: schema.TypeString, Optional: true},
			"locality":            {Type: schema.TypeString, Optional: true},
			"organization":        {Type: schema.TypeString, Optional: true},
			"organizational_unit": {Type: schema.TypeString, Optional: true},
			"state":               {Type: schema.TypeString, Optional: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadAllowDenyClients(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                d.Get("name").(string),
		"action":              d.Get("action").(string),
		"sequence":            d.Get("sequence").(string),
		"certificate-serial":  d.Get("certificate_serial").(string),
		"common-name":         d.Get("common_name").(string),
		"country":             d.Get("country").(string),
		"locality":            d.Get("locality").(string),
		"organization":        d.Get("organization").(string),
		"organizational-unit": d.Get("organizational_unit").(string),
		"state":               d.Get("state").(string),
		"status":              d.Get("status").(string),
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

func resourceCudaWAFAllowDenyClientsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients"
	resourceCreateResponseError := makeRestAPIPayloadAllowDenyClients(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFAllowDenyClientsRead(d, m)
}

func resourceCudaWAFAllowDenyClientsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAllowDenyClientsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadAllowDenyClients(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFAllowDenyClientsRead(d, m)
}

func resourceCudaWAFAllowDenyClientsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadAllowDenyClients(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
