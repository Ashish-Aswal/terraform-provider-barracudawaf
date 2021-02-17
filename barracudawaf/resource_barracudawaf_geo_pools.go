package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGeoPools() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGeoPoolsCreate,
		Read:   resourceCudaWAFGeoPoolsRead,
		Update: resourceCudaWAFGeoPoolsUpdate,
		Delete: resourceCudaWAFGeoPoolsDelete,

		Schema: map[string]*schema.Schema{
			"region": {Type: schema.TypeString, Optional: true},
			"name":   {Type: schema.TypeString, Required: true},
		},
	}
}

func makeRestAPIPayloadGeoPools(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"region": d.Get("region").(string),
		"name":   d.Get("name").(string),
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

func resourceCudaWAFGeoPoolsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/geo-pools"
	resourceCreateResponseError := makeRestAPIPayloadGeoPools(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFGeoPoolsRead(d, m)
}

func resourceCudaWAFGeoPoolsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGeoPoolsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/geo-pools/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadGeoPools(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFGeoPoolsRead(d, m)
}

func resourceCudaWAFGeoPoolsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/geo-pools/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadGeoPools(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
