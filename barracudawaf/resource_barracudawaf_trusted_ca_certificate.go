package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFTrustedCaCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFTrustedCaCertificateCreate,
		Read:   resourceCudaWAFTrustedCaCertificateRead,
		Update: resourceCudaWAFTrustedCaCertificateUpdate,
		Delete: resourceCudaWAFTrustedCaCertificateDelete,

		Schema: map[string]*schema.Schema{
			"common_name":   {Type: schema.TypeString, Optional: true},
			"expiry":        {Type: schema.TypeString, Optional: true},
			"name":          {Type: schema.TypeString, Required: true},
			"serial":        {Type: schema.TypeString, Optional: true},
			"certificate":   {Type: schema.TypeString, Optional: true},
			"download_type": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadTrustedCaCertificate(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"common-name":   d.Get("common_name").(string),
		"expiry":        d.Get("expiry").(string),
		"name":          d.Get("name").(string),
		"serial":        d.Get("serial").(string),
		"certificate":   d.Get("certificate").(string),
		"download-type": d.Get("download_type").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{
			"common-name",
			"expiry",
			"name",
			"serial",
			"certificate",
		}
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

func resourceCudaWAFTrustedCaCertificateCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-ca-certificate"
	resourceCreateResponseError := makeRestAPIPayloadTrustedCaCertificate(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFTrustedCaCertificateRead(d, m)
}

func resourceCudaWAFTrustedCaCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFTrustedCaCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-ca-certificate/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadTrustedCaCertificate(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFTrustedCaCertificateRead(d, m)
}

func resourceCudaWAFTrustedCaCertificateDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-ca-certificate/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadTrustedCaCertificate(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
