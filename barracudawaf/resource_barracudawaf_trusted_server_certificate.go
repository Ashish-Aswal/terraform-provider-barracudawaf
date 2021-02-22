package barracudawaf

import (
	"fmt"
	"log"

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

func resourceCudaWAFTrustedServerCertificateCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/trusted-server-certificate"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFTrustedServerCertificateResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFTrustedServerCertificateRead(d, m)
}

func resourceCudaWAFTrustedServerCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFTrustedServerCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/trusted-server-certificate/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFTrustedServerCertificateResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFTrustedServerCertificateRead(d, m)
}

func resourceCudaWAFTrustedServerCertificateDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/trusted-server-certificate/"
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

func hydrateBarracudaWAFTrustedServerCertificateResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"common-name":   d.Get("common_name").(string),
		"expiry":        d.Get("expiry").(string),
		"serial":        d.Get("serial").(string),
		"name":          d.Get("name").(string),
		"certificate":   d.Get("certificate").(string),
		"download-type": d.Get("download_type").(string),
	}

	// parameters not supported for updates
	if method == "put" {
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
