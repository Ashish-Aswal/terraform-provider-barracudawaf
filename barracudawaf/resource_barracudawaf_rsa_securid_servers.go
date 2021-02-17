package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRsaSecuridServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRsaSecuridServersCreate,
		Read:   resourceCudaWAFRsaSecuridServersRead,
		Update: resourceCudaWAFRsaSecuridServersUpdate,
		Delete: resourceCudaWAFRsaSecuridServersDelete,

		Schema: map[string]*schema.Schema{
			"retries":       {Type: schema.TypeString, Optional: true},
			"shared_secret": {Type: schema.TypeString, Optional: true},
			"ip_address":    {Type: schema.TypeString, Required: true},
			"port":          {Type: schema.TypeString, Optional: true},
			"timeout":       {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadRsaSecuridServers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"retries":       d.Get("retries").(string),
		"shared-secret": d.Get("shared_secret").(string),
		"ip-address":    d.Get("ip_address").(string),
		"port":          d.Get("port").(string),
		"timeout":       d.Get("timeout").(string),
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

func resourceCudaWAFRsaSecuridServersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rsa-securid-services/" + d.Get("parent.0").(string) + "/rsa-securid-servers"
	resourceCreateResponseError := makeRestAPIPayloadRsaSecuridServers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFRsaSecuridServersRead(d, m)
}

func resourceCudaWAFRsaSecuridServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFRsaSecuridServersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rsa-securid-services/" + d.Get("parent.0").(string) + "/rsa-securid-servers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadRsaSecuridServers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFRsaSecuridServersRead(d, m)
}

func resourceCudaWAFRsaSecuridServersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rsa-securid-services/" + d.Get("parent.0").(string) + "/rsa-securid-servers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadRsaSecuridServers(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
