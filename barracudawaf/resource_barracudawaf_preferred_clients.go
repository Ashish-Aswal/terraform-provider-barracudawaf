package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFPreferredClients() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFPreferredClientsCreate,
		Read:   resourceCudaWAFPreferredClientsRead,
		Update: resourceCudaWAFPreferredClientsUpdate,
		Delete: resourceCudaWAFPreferredClientsDelete,

		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"ip_range": {Type: schema.TypeString, Required: true},
			"status":   {Type: schema.TypeString, Optional: true},
			"weight":   {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFPreferredClientsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFPreferredClientsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFPreferredClientsRead(d, m)
}

func resourceCudaWAFPreferredClientsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients"
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

func resourceCudaWAFPreferredClientsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFPreferredClientsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFPreferredClientsRead(d, m)
}

func resourceCudaWAFPreferredClientsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/rate-control-pools/" + d.Get("parent.0").(string) + "/preferred-clients/"
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

func hydrateBarracudaWAFPreferredClientsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":     d.Get("name").(string),
		"ip-range": d.Get("ip_range").(string),
		"status":   d.Get("status").(string),
		"weight":   d.Get("weight").(string),
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
