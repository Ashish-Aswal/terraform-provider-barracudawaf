package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFVirtualInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFVirtualInterfacesCreate,
		Read:   resourceCudaWAFVirtualInterfacesRead,
		Update: resourceCudaWAFVirtualInterfacesUpdate,
		Delete: resourceCudaWAFVirtualInterfacesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address":            {Type: schema.TypeString, Required: true},
			"comments":              {Type: schema.TypeString, Optional: true},
			"interface":             {Type: schema.TypeString, Required: true},
			"ip_version":            {Type: schema.TypeString, Optional: true},
			"netmask":               {Type: schema.TypeString, Required: true},
			"virtual_ip_service_id": {Type: schema.TypeString, Optional: true},
			"vsite":                 {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadVirtualInterfaces(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"ip-address":            d.Get("ip_address").(string),
		"comments":              d.Get("comments").(string),
		"interface":             d.Get("interface").(string),
		"ip-version":            d.Get("ip_version").(string),
		"netmask":               d.Get("netmask").(string),
		"virtual-ip-service-id": d.Get("virtual_ip_service_id").(string),
		"vsite":                 d.Get("vsite").(string),
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

func resourceCudaWAFVirtualInterfacesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/virtual-interfaces"
	resourceCreateResponseError := makeRestAPIPayloadVirtualInterfaces(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFVirtualInterfacesRead(d, m)
}

func resourceCudaWAFVirtualInterfacesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFVirtualInterfacesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/virtual-interfaces/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadVirtualInterfaces(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFVirtualInterfacesRead(d, m)
}

func resourceCudaWAFVirtualInterfacesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/virtual-interfaces/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadVirtualInterfaces(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
