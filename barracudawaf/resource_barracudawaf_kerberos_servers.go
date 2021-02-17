package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFKerberosServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFKerberosServersCreate,
		Read:   resourceCudaWAFKerberosServersRead,
		Update: resourceCudaWAFKerberosServersUpdate,
		Delete: resourceCudaWAFKerberosServersDelete,

		Schema: map[string]*schema.Schema{
			"kdc_name":     {Type: schema.TypeString, Required: true},
			"domain_alias": {Type: schema.TypeString, Optional: true},
			"ip_address":   {Type: schema.TypeString, Required: true},
			"port":         {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadKerberosServers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"kdc-name":     d.Get("kdc_name").(string),
		"domain-alias": d.Get("domain_alias").(string),
		"ip-address":   d.Get("ip_address").(string),
		"port":         d.Get("port").(string),
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

func resourceCudaWAFKerberosServersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers"
	resourceCreateResponseError := makeRestAPIPayloadKerberosServers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFKerberosServersRead(d, m)
}

func resourceCudaWAFKerberosServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFKerberosServersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadKerberosServers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFKerberosServersRead(d, m)
}

func resourceCudaWAFKerberosServersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadKerberosServers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
