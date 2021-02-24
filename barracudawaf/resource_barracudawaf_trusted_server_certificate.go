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
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/trusted-server-certificate"
	request := &APIRequest{
		Method: "get",
		URL:    resourceEndpoint,
	}

	var dataItems map[string]interface{}
	resources, err := client.GetBarracudaWAFResource(name, request)

	if err != nil {
		log.Printf("[ERROR] Unable to Retrieve Barracuda WAF resource (%s) (%v) ", name, err)
		return err
	}

	if resources.Data == nil {
		log.Printf("[WARN] Barracuda WAF resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	for _, dataItems = range resources.Data {
		if dataItems["name"] == name {
			break
		}
	}

	if dataItems["name"] != name {
		return fmt.Errorf("Barracuda WAF resource (%s) not found on the system", name)
	}

	d.Set("name", name)
	return nil
}

func resourceCudaWAFTrustedServerCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/trusted-server-certificate/"
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
		return fmt.Errorf("Unable to delete the Barracuda WAF resource (%s) (%v)", name, err)
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
