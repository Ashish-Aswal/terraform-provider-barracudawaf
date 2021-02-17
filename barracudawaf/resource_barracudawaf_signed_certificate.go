package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSignedCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSignedCertificateCreate,
		Read:   resourceCudaWAFSignedCertificateRead,
		Update: resourceCudaWAFSignedCertificateUpdate,
		Delete: resourceCudaWAFSignedCertificateDelete,

		Schema: map[string]*schema.Schema{
			"assign_associated_key":     {Type: schema.TypeString, Optional: true},
			"signed_certificate":        {Type: schema.TypeString, Optional: true},
			"certificate_key":           {Type: schema.TypeString, Optional: true},
			"certificate_password":      {Type: schema.TypeString, Optional: true},
			"certificate_type":          {Type: schema.TypeString, Optional: true},
			"download_type":             {Type: schema.TypeString, Optional: true},
			"encrypt_password":          {Type: schema.TypeString, Optional: true},
			"intermediary_certificates": {Type: schema.TypeString, Optional: true},
			"name":                      {Type: schema.TypeString, Optional: true},
			"auto_renew_cert":           {Type: schema.TypeString, Optional: true},
			"common_name":               {Type: schema.TypeString, Optional: true},
			"expiry":                    {Type: schema.TypeString, Optional: true},
			"key_type":                  {Type: schema.TypeString, Optional: true},
			"allow_private_key_export":  {Type: schema.TypeString, Optional: true},
			"schedule_renewal_day":      {Type: schema.TypeString, Optional: true},
			"serial":                    {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadSignedCertificate(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"assign-associated-key":     d.Get("assign_associated_key").(string),
		"signed-certificate":        d.Get("signed_certificate").(string),
		"certificate-key":           d.Get("certificate_key").(string),
		"certificate-password":      d.Get("certificate_password").(string),
		"certificate-type":          d.Get("certificate_type").(string),
		"download-type":             d.Get("download_type").(string),
		"encrypt-password":          d.Get("encrypt_password").(string),
		"intermediary-certificates": d.Get("intermediary_certificates").(string),
		"name":                      d.Get("name").(string),
		"auto-renew-cert":           d.Get("auto_renew_cert").(string),
		"common-name":               d.Get("common_name").(string),
		"expiry":                    d.Get("expiry").(string),
		"key-type":                  d.Get("key_type").(string),
		"allow-private-key-export":  d.Get("allow_private_key_export").(string),
		"schedule-renewal-day":      d.Get("schedule_renewal_day").(string),
		"serial":                    d.Get("serial").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{
			"assign-associated-key",
			"signed-certificate",
			"certificate-key",
			"certificate-password",
			"certificate-type",
			"intermediary-certificates",
			"name",
			"common-name",
			"expiry",
			"key-type",
			"allow-private-key-export",
			"serial",
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

func resourceCudaWAFSignedCertificateCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/signed-certificate"
	resourceCreateResponseError := makeRestAPIPayloadSignedCertificate(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSignedCertificateRead(d, m)
}

func resourceCudaWAFSignedCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSignedCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/signed-certificate/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSignedCertificate(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSignedCertificateRead(d, m)
}

func resourceCudaWAFSignedCertificateDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/signed-certificate/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSignedCertificate(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
