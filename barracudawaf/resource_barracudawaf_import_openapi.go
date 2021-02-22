package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFImportOpenapi() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFImportOpenapiCreate,
		Read:   resourceCudaWAFImportOpenapiRead,
		Update: resourceCudaWAFImportOpenapiUpdate,
		Delete: resourceCudaWAFImportOpenapiDelete,

		Schema: map[string]*schema.Schema{
			"operation":           {Type: schema.TypeString, Optional: true},
			"import_file_content": {Type: schema.TypeString, Optional: true},
			"advanced_config":     {Type: schema.TypeString, Optional: true},
			"service_name":        {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFImportOpenapiCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/import-openapi"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFImportOpenapiResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFImportOpenapiRead(d, m)
}

func resourceCudaWAFImportOpenapiRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFImportOpenapiUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/import-openapi/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFImportOpenapiResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFImportOpenapiRead(d, m)
}

func resourceCudaWAFImportOpenapiDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/import-openapi/"
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

func hydrateBarracudaWAFImportOpenapiResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"operation":           d.Get("operation").(string),
		"import-file-content": d.Get("import_file_content").(string),
		"advanced-config":     d.Get("advanced_config").(string),
		"service-name":        d.Get("service_name").(string),
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
