package barracudawaf

import (
	"fmt"
	"log"

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

func resourceCudaWAFSignedCertificateCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/signed-certificate"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSignedCertificateResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSignedCertificateRead(d, m)
}

func resourceCudaWAFSignedCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSignedCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/signed-certificate/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSignedCertificateResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSignedCertificateRead(d, m)
}

func resourceCudaWAFSignedCertificateDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/signed-certificate/"
	request := &APIRequest{
		Method: "delete",
		URL:    resourceEndpoint,
	}

	err := client.DeleteBarracudaWAFResource(name, request)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func hydrateBarracudaWAFSignedCertificateResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
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

	// parameters not supported for updates
	if method == "put" {
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

	// remove empty parameters from resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	return &APIRequest{
		URL:  endpoint,
		Body: resourcePayload,
	}
}
