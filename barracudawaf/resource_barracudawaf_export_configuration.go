package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFExportConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFExportConfigurationCreate,
		Read:   resourceCudaWAFExportConfigurationRead,
		Update: resourceCudaWAFExportConfigurationUpdate,
		Delete: resourceCudaWAFExportConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"backup_type":    {Type: schema.TypeString, Optional: true},
			"name":           {Type: schema.TypeString, Optional: true},
			"destination":    {Type: schema.TypeString, Optional: true},
			"day_of_week":    {Type: schema.TypeString, Optional: true},
			"hour_of_day":    {Type: schema.TypeString, Optional: true},
			"minute_of_hour": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFExportConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/export-configuration"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFExportConfigurationResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFExportConfigurationRead(d, m)
}

func resourceCudaWAFExportConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFExportConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/export-configuration/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFExportConfigurationResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFExportConfigurationRead(d, m)
}

func resourceCudaWAFExportConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/export-configuration/"
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

func hydrateBarracudaWAFExportConfigurationResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"backup-type":    d.Get("backup_type").(string),
		"name":           d.Get("name").(string),
		"destination":    d.Get("destination").(string),
		"day-of-week":    d.Get("day_of_week").(string),
		"hour-of-day":    d.Get("hour_of_day").(string),
		"minute-of-hour": d.Get("minute_of_hour").(string),
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
