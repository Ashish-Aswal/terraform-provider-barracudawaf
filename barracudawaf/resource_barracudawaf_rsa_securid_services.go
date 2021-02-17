package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRsaSecuridServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRsaSecuridServicesCreate,
		Read:   resourceCudaWAFRsaSecuridServicesRead,
		Update: resourceCudaWAFRsaSecuridServicesUpdate,
		Delete: resourceCudaWAFRsaSecuridServicesDelete,

		Schema: map[string]*schema.Schema{"name": {Type: schema.TypeString, Required: true}},
	}
}

func makeRestAPIPayloadRsaSecuridServices(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{"name": d.Get("name").(string)}

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

func resourceCudaWAFRsaSecuridServicesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rsa-securid-services"
	resourceCreateResponseError := makeRestAPIPayloadRsaSecuridServices(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFRsaSecuridServicesRead(d, m)
}

func resourceCudaWAFRsaSecuridServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFRsaSecuridServicesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rsa-securid-services/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadRsaSecuridServices(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFRsaSecuridServicesRead(d, m)
}

func resourceCudaWAFRsaSecuridServicesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/rsa-securid-services/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadRsaSecuridServices(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
