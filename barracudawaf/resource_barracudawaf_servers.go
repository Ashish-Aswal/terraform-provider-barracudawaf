package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServersCreate,
		Read:   resourceCudaWAFServersRead,
		Update: resourceCudaWAFServersUpdate,
		Delete: resourceCudaWAFServersDelete,

		Schema: map[string]*schema.Schema{
			"address_version": {Type: schema.TypeString, Optional: true},
			"comments":        {Type: schema.TypeString, Optional: true},
			"name":            {Type: schema.TypeString, Optional: true},
			"hostname":        {Type: schema.TypeString, Optional: true},
			"identifier":      {Type: schema.TypeString, Optional: true},
			"ip_address":      {Type: schema.TypeString, Optional: true},
			"port":            {Type: schema.TypeString, Optional: true},
			"status":          {Type: schema.TypeString, Optional: true},
			"resolved_ips":    {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadServers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"address-version": d.Get("address_version").(string),
		"comments":        d.Get("comments").(string),
		"name":            d.Get("name").(string),
		"hostname":        d.Get("hostname").(string),
		"identifier":      d.Get("identifier").(string),
		"ip-address":      d.Get("ip_address").(string),
		"port":            d.Get("port").(string),
		"status":          d.Get("status").(string),
		"resolved-ips":    d.Get("resolved_ips").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"address-version", "resolved-ips"}
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

func resourceCudaWAFServersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/servers"
	resourceCreateResponseError := makeRestAPIPayloadServers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFServersRead(d, m)
}

func resourceCudaWAFServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/servers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadServers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFServersRead(d, m)
}

func resourceCudaWAFServersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/servers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadServers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
