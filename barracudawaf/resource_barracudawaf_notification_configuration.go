package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFNotificationConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFNotificationConfigurationCreate,
		Read:   resourceCudaWAFNotificationConfigurationRead,
		Update: resourceCudaWAFNotificationConfigurationUpdate,
		Delete: resourceCudaWAFNotificationConfigurationDelete,

		Schema: map[string]*schema.Schema{"severity": {Type: schema.TypeString, Optional: true}},
	}
}

func resourceCudaWAFNotificationConfigurationCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/notification-configuration"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNotificationConfigurationResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFNotificationConfigurationRead(d, m)
}

func resourceCudaWAFNotificationConfigurationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFNotificationConfigurationUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/notification-configuration/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFNotificationConfigurationResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFNotificationConfigurationRead(d, m)
}

func resourceCudaWAFNotificationConfigurationDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/notification-configuration/"
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

func hydrateBarracudaWAFNotificationConfigurationResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{"severity": d.Get("severity").(string)}

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
