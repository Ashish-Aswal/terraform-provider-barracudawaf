package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFJsonProfiles() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFJsonProfilesCreate,
		Read:   resourceCudaWAFJsonProfilesRead,
		Update: resourceCudaWAFJsonProfilesUpdate,
		Delete: resourceCudaWAFJsonProfilesDelete,

		Schema: map[string]*schema.Schema{
			"allowed_content_types": {Type: schema.TypeString, Optional: true},
			"host_match":            {Type: schema.TypeString, Required: true},
			"ignore_keys":           {Type: schema.TypeString, Optional: true},
			"method":                {Type: schema.TypeString, Required: true},
			"comment":               {Type: schema.TypeString, Optional: true},
			"name":                  {Type: schema.TypeString, Required: true},
			"mode":                  {Type: schema.TypeString, Optional: true},
			"status":                {Type: schema.TypeString, Optional: true},
			"url_match":             {Type: schema.TypeString, Required: true},
			"exception_patterns":    {Type: schema.TypeString, Optional: true},
			"json_policy":           {Type: schema.TypeString, Optional: true},
			"validate_key":          {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFJsonProfilesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/json-profiles"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFJsonProfilesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFJsonProfilesRead(d, m)
}

func resourceCudaWAFJsonProfilesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFJsonProfilesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/json-profiles/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFJsonProfilesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFJsonProfilesRead(d, m)
}

func resourceCudaWAFJsonProfilesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/json-profiles/"
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

func hydrateBarracudaWAFJsonProfilesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"allowed-content-types": d.Get("allowed_content_types").(string),
		"host_match":            d.Get("host_match").(string),
		"ignore-keys":           d.Get("ignore_keys").(string),
		"method":                d.Get("method").(string),
		"comment":               d.Get("comment").(string),
		"name":                  d.Get("name").(string),
		"mode":                  d.Get("mode").(string),
		"status":                d.Get("status").(string),
		"url_match":             d.Get("url_match").(string),
		"exception-patterns":    d.Get("exception_patterns").(string),
		"json_policy":           d.Get("json_policy").(string),
		"validate_key":          d.Get("validate_key").(string),
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
