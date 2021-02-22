package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFConfigurationCheckpoints() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFConfigurationCheckpointsCreate,
		Read:   resourceCudaWAFConfigurationCheckpointsRead,
		Update: resourceCudaWAFConfigurationCheckpointsUpdate,
		Delete: resourceCudaWAFConfigurationCheckpointsDelete,

		Schema: map[string]*schema.Schema{
			"name":    {Type: schema.TypeString, Required: true},
			"comment": {Type: schema.TypeString, Optional: true},
			"date":    {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFConfigurationCheckpointsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/configuration-checkpoints"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFConfigurationCheckpointsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFConfigurationCheckpointsRead(d, m)
}

func resourceCudaWAFConfigurationCheckpointsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFConfigurationCheckpointsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/configuration-checkpoints/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFConfigurationCheckpointsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFConfigurationCheckpointsRead(d, m)
}

func resourceCudaWAFConfigurationCheckpointsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/configuration-checkpoints/"
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

func hydrateBarracudaWAFConfigurationCheckpointsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":    d.Get("name").(string),
		"comment": d.Get("comment").(string),
		"date":    d.Get("date").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"date"}
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
