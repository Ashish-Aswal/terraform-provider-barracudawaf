package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFParameterOptimizers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFParameterOptimizersCreate,
		Read:   resourceCudaWAFParameterOptimizersRead,
		Update: resourceCudaWAFParameterOptimizersUpdate,
		Delete: resourceCudaWAFParameterOptimizersDelete,

		Schema: map[string]*schema.Schema{
			"name":        {Type: schema.TypeString, Required: true},
			"start_token": {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFParameterOptimizersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/parameter-optimizers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFParameterOptimizersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFParameterOptimizersRead(d, m)
}

func resourceCudaWAFParameterOptimizersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFParameterOptimizersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/parameter-optimizers/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFParameterOptimizersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFParameterOptimizersRead(d, m)
}

func resourceCudaWAFParameterOptimizersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/parameter-optimizers/"
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

func hydrateBarracudaWAFParameterOptimizersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":        d.Get("name").(string),
		"start-token": d.Get("start_token").(string),
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
