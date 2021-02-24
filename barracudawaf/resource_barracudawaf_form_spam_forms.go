package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFFormSpamForms() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFFormSpamFormsCreate,
		Read:   resourceCudaWAFFormSpamFormsRead,
		Update: resourceCudaWAFFormSpamFormsUpdate,
		Delete: resourceCudaWAFFormSpamFormsDelete,

		Schema: map[string]*schema.Schema{
			"name":                   {Type: schema.TypeString, Required: true},
			"created_by":             {Type: schema.TypeString, Optional: true},
			"status":                 {Type: schema.TypeString, Optional: true},
			"mode":                   {Type: schema.TypeString, Optional: true},
			"action_url":             {Type: schema.TypeString, Required: true},
			"minimum_form_fill_time": {Type: schema.TypeString, Optional: true},
			"parameter_name":         {Type: schema.TypeString, Optional: true},
			"parameter_class":        {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFFormSpamFormsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/form-spam-forms"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFFormSpamFormsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFFormSpamFormsRead(d, m)
}

func resourceCudaWAFFormSpamFormsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/form-spam-forms"
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

func resourceCudaWAFFormSpamFormsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/form-spam-forms/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFFormSpamFormsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFFormSpamFormsRead(d, m)
}

func resourceCudaWAFFormSpamFormsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/form-spam-forms/"
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

func hydrateBarracudaWAFFormSpamFormsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":                   d.Get("name").(string),
		"created-by":             d.Get("created_by").(string),
		"status":                 d.Get("status").(string),
		"mode":                   d.Get("mode").(string),
		"action-url":             d.Get("action_url").(string),
		"minimum-form-fill-time": d.Get("minimum_form_fill_time").(string),
		"parameter-name":         d.Get("parameter_name").(string),
		"parameter-class":        d.Get("parameter_class").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"created-by", "action-url"}
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
