package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFLocalHosts() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFLocalHostsCreate,
		Read:   resourceCudaWAFLocalHostsRead,
		Update: resourceCudaWAFLocalHostsUpdate,
		Delete: resourceCudaWAFLocalHostsDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {Type: schema.TypeString, Required: true},
			"hostname":   {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadLocalHosts(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"ip-address": d.Get("ip_address").(string),
		"hostname":   d.Get("hostname").(string),
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

func resourceCudaWAFLocalHostsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-hosts"
	resourceCreateResponseError := makeRestAPIPayloadLocalHosts(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFLocalHostsRead(d, m)
}

func resourceCudaWAFLocalHostsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFLocalHostsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-hosts/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadLocalHosts(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFLocalHostsRead(d, m)
}

func resourceCudaWAFLocalHostsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/local-hosts/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadLocalHosts(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
