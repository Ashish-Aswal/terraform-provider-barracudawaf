package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAdminIpRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAdminIpRangeCreate,
		Read:   resourceCudaWAFAdminIpRangeRead,
		Update: resourceCudaWAFAdminIpRangeUpdate,
		Delete: resourceCudaWAFAdminIpRangeDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {Type: schema.TypeString, Required: true},
			"netmask":    {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadAdminIpRange(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"ip-address": d.Get("ip_address").(string),
		"netmask":    d.Get("netmask").(string),
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

func resourceCudaWAFAdminIpRangeCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/admin-ip-range"
	resourceCreateResponseError := makeRestAPIPayloadAdminIpRange(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFAdminIpRangeRead(d, m)
}

func resourceCudaWAFAdminIpRangeRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAdminIpRangeUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/admin-ip-range/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadAdminIpRange(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFAdminIpRangeRead(d, m)
}

func resourceCudaWAFAdminIpRangeDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/admin-ip-range/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadAdminIpRange(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
