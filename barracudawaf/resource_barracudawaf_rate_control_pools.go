package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRateControlPools() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRateControlPoolsCreate,
		Read:   resourceCudaWAFRateControlPoolsRead,
		Update: resourceCudaWAFRateControlPoolsUpdate,
		Delete: resourceCudaWAFRateControlPoolsDelete,

		Schema: map[string]*schema.Schema{
			"name":                     {Type: schema.TypeString, Required: true},
			"max_active_requests":      {Type: schema.TypeString, Optional: true},
			"max_per_client_backlog":   {Type: schema.TypeString, Optional: true},
			"max_unconfigured_clients": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadRateControlPools(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                     d.Get("name").(string),
		"max-active-requests":      d.Get("max_active_requests").(string),
		"max-per-client-backlog":   d.Get("max_per_client_backlog").(string),
		"max-unconfigured-clients": d.Get("max_unconfigured_clients").(string),
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

func resourceCudaWAFRateControlPoolsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rate-control-pools"
	resourceCreateResponseError := makeRestAPIPayloadRateControlPools(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFRateControlPoolsRead(d, m)
}

func resourceCudaWAFRateControlPoolsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFRateControlPoolsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rate-control-pools/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadRateControlPools(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFRateControlPoolsRead(d, m)
}

func resourceCudaWAFRateControlPoolsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rate-control-pools/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadRateControlPools(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
