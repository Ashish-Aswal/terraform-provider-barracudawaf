package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFPreferredClients() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFPreferredClientsCreate,
		Read:   resourceCudaWAFPreferredClientsRead,
		Update: resourceCudaWAFPreferredClientsUpdate,
		Delete: resourceCudaWAFPreferredClientsDelete,

		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"ip_range": {Type: schema.TypeString, Required: true},
			"status":   {Type: schema.TypeString, Optional: true},
			"weight":   {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadPreferredClients(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":     d.Get("name").(string),
		"ip-range": d.Get("ip_range").(string),
		"status":   d.Get("status").(string),
		"weight":   d.Get("weight").(string),
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

func resourceCudaWAFPreferredClientsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients"
	resourceCreateResponseError := makeRestAPIPayloadPreferredClients(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFPreferredClientsRead(d, m)
}

func resourceCudaWAFPreferredClientsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFPreferredClientsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadPreferredClients(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFPreferredClientsRead(d, m)
}

func resourceCudaWAFPreferredClientsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadPreferredClients(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
