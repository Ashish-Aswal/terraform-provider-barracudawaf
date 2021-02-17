package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFTrustedServerCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFTrustedServerCertificateCreate,
		Read:   resourceCudaWAFTrustedServerCertificateRead,
		Update: resourceCudaWAFTrustedServerCertificateUpdate,
		Delete: resourceCudaWAFTrustedServerCertificateDelete,

		Schema: map[string]*schema.Schema{
			"common_name":   {Type: schema.TypeString, Optional: true},
			"expiry":        {Type: schema.TypeString, Optional: true},
			"serial":        {Type: schema.TypeString, Optional: true},
			"name":          {Type: schema.TypeString, Optional: true},
			"certificate":   {Type: schema.TypeString, Optional: true},
			"download_type": {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadTrustedServerCertificate(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"common-name":   d.Get("common_name").(string),
		"expiry":        d.Get("expiry").(string),
		"serial":        d.Get("serial").(string),
		"name":          d.Get("name").(string),
		"certificate":   d.Get("certificate").(string),
		"download-type": d.Get("download_type").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{
			"common-name",
			"expiry",
			"serial",
			"name",
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

func resourceCudaWAFTrustedServerCertificateCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-server-certificate"
	resourceCreateResponseError := makeRestAPIPayloadTrustedServerCertificate(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFTrustedServerCertificateRead(d, m)
}

func resourceCudaWAFTrustedServerCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFTrustedServerCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-server-certificate/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadTrustedServerCertificate(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFTrustedServerCertificateRead(d, m)
}

func resourceCudaWAFTrustedServerCertificateDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/trusted-server-certificate/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadTrustedServerCertificate(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
