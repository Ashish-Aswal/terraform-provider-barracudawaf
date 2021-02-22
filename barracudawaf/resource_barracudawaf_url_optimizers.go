package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlOptimizers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlOptimizersCreate,
		Read:   resourceCudaWAFUrlOptimizersRead,
		Update: resourceCudaWAFUrlOptimizersUpdate,
		Delete: resourceCudaWAFUrlOptimizersDelete,

		Schema: map[string]*schema.Schema{
			"end_delimiter": {Type: schema.TypeString, Optional: true},
			"name":          {Type: schema.TypeString, Required: true},
			"start_token":   {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFUrlOptimizersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-optimizers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFUrlOptimizersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFUrlOptimizersRead(d, m)
}

func resourceCudaWAFUrlOptimizersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUrlOptimizersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-optimizers/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFUrlOptimizersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFUrlOptimizersRead(d, m)
}

func resourceCudaWAFUrlOptimizersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-optimizers/"
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

func hydrateBarracudaWAFUrlOptimizersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"end-delimiter": d.Get("end_delimiter").(string),
		"name":          d.Get("name").(string),
		"start-token":   d.Get("start_token").(string),
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
