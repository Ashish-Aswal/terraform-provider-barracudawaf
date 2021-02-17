package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFGeoipAllowedNetworks() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFGeoipAllowedNetworksCreate,
		Read:   resourceCudaWAFGeoipAllowedNetworksRead,
		Update: resourceCudaWAFGeoipAllowedNetworksUpdate,
		Delete: resourceCudaWAFGeoipAllowedNetworksDelete,

		Schema: map[string]*schema.Schema{
			"comment":       {Type: schema.TypeString, Optional: true},
			"allow_ip":      {Type: schema.TypeString, Required: true},
			"allow_netmask": {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadGeoipAllowedNetworks(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"comment":       d.Get("comment").(string),
		"allow-ip":      d.Get("allow_ip").(string),
		"allow-netmask": d.Get("allow_netmask").(string),
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

func resourceCudaWAFGeoipAllowedNetworksCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/geoip-allowed-networks"
	resourceCreateResponseError := makeRestAPIPayloadGeoipAllowedNetworks(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFGeoipAllowedNetworksRead(d, m)
}

func resourceCudaWAFGeoipAllowedNetworksRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFGeoipAllowedNetworksUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/geoip-allowed-networks/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadGeoipAllowedNetworks(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFGeoipAllowedNetworksRead(d, m)
}

func resourceCudaWAFGeoipAllowedNetworksDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/geoip-allowed-networks/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadGeoipAllowedNetworks(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
