package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSelfSignedCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSelfSignedCertificateCreate,
		Read:   resourceCudaWAFSelfSignedCertificateRead,
		Update: resourceCudaWAFSelfSignedCertificateUpdate,
		Delete: resourceCudaWAFSelfSignedCertificateDelete,

		Schema: map[string]*schema.Schema{
			"download_type":            {Type: schema.TypeString, Optional: true},
			"encrypt_password":         {Type: schema.TypeString, Optional: true},
			"city":                     {Type: schema.TypeString, Optional: true},
			"common_name":              {Type: schema.TypeString, Required: true},
			"country_code":             {Type: schema.TypeString, Required: true},
			"elliptic_curve_name":      {Type: schema.TypeString, Optional: true},
			"expiry":                   {Type: schema.TypeString, Optional: true},
			"key_size":                 {Type: schema.TypeString, Optional: true},
			"key_type":                 {Type: schema.TypeString, Optional: true},
			"allow_private_key_export": {Type: schema.TypeString, Optional: true},
			"name":                     {Type: schema.TypeString, Required: true},
			"organization_name":        {Type: schema.TypeString, Optional: true},
			"organizational_unit":      {Type: schema.TypeString, Optional: true},
			"san_certificate":          {Type: schema.TypeString, Optional: true},
			"serial":                   {Type: schema.TypeString, Optional: true},
			"state":                    {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadSelfSignedCertificate(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"download-type":            d.Get("download_type").(string),
		"encrypt-password":         d.Get("encrypt_password").(string),
		"city":                     d.Get("city").(string),
		"common-name":              d.Get("common_name").(string),
		"country-code":             d.Get("country_code").(string),
		"elliptic-curve-name":      d.Get("elliptic_curve_name").(string),
		"expiry":                   d.Get("expiry").(string),
		"key-size":                 d.Get("key_size").(string),
		"key-type":                 d.Get("key_type").(string),
		"allow-private-key-export": d.Get("allow_private_key_export").(string),
		"name":                     d.Get("name").(string),
		"organization-name":        d.Get("organization_name").(string),
		"organizational-unit":      d.Get("organizational_unit").(string),
		"san-certificate":          d.Get("san_certificate").(string),
		"serial":                   d.Get("serial").(string),
		"state":                    d.Get("state").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{
			"city",
			"common-name",
			"country-code",
			"elliptic-curve-name",
			"expiry",
			"key-size",
			"key-type",
			"allow-private-key-export",
			"name",
			"organization-name",
			"organizational-unit",
			"san-certificate",
			"serial",
			"state",
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

func resourceCudaWAFSelfSignedCertificateCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/self-signed-certificate"
	resourceCreateResponseError := makeRestAPIPayloadSelfSignedCertificate(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSelfSignedCertificateRead(d, m)
}

func resourceCudaWAFSelfSignedCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSelfSignedCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/self-signed-certificate/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSelfSignedCertificate(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSelfSignedCertificateRead(d, m)
}

func resourceCudaWAFSelfSignedCertificateDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/self-signed-certificate/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSelfSignedCertificate(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
