package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSamlIdentityProviders() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSamlIdentityProvidersCreate,
		Read:   resourceCudaWAFSamlIdentityProvidersRead,
		Update: resourceCudaWAFSamlIdentityProvidersUpdate,
		Delete: resourceCudaWAFSamlIdentityProvidersDelete,

		Schema: map[string]*schema.Schema{
			"metadata_file":       {Type: schema.TypeString, Optional: true},
			"metadata_type":       {Type: schema.TypeString, Optional: true},
			"autoupdate_metadata": {Type: schema.TypeString, Optional: true},
			"metadata_url":        {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"metadata_content":    {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFSamlIdentityProvidersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSamlIdentityProvidersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSamlIdentityProvidersRead(d, m)
}

func resourceCudaWAFSamlIdentityProvidersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers"
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

func resourceCudaWAFSamlIdentityProvidersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSamlIdentityProvidersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSamlIdentityProvidersRead(d, m)
}

func resourceCudaWAFSamlIdentityProvidersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers/"
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

func hydrateBarracudaWAFSamlIdentityProvidersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"metadata-file":       d.Get("metadata_file").(string),
		"metadata-type":       d.Get("metadata_type").(string),
		"autoupdate-metadata": d.Get("autoupdate_metadata").(string),
		"metadata-url":        d.Get("metadata_url").(string),
		"name":                d.Get("name").(string),
		"metadata-content":    d.Get("metadata_content").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{}
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
