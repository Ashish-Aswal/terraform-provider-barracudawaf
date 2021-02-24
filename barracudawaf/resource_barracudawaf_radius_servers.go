package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFRadiusServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFRadiusServersCreate,
		Read:   resourceCudaWAFRadiusServersRead,
		Update: resourceCudaWAFRadiusServersUpdate,
		Delete: resourceCudaWAFRadiusServersDelete,

		Schema: map[string]*schema.Schema{
			"shared_secret": {Type: schema.TypeString, Optional: true},
			"ip_address":    {Type: schema.TypeString, Required: true},
			"port":          {Type: schema.TypeString, Optional: true},
			"timeout":       {Type: schema.TypeString, Optional: true},
			"retries":       {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFRadiusServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/radius-services/" + d.Get("parent.0").(string) + "/radius-servers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFRadiusServersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFRadiusServersRead(d, m)
}

func resourceCudaWAFRadiusServersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/radius-services/" + d.Get("parent.0").(string) + "/radius-servers"
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

func resourceCudaWAFRadiusServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/radius-services/" + d.Get("parent.0").(string) + "/radius-servers/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFRadiusServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFRadiusServersRead(d, m)
}

func resourceCudaWAFRadiusServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/radius-services/" + d.Get("parent.0").(string) + "/radius-servers/"
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

func hydrateBarracudaWAFRadiusServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"shared-secret": d.Get("shared_secret").(string),
		"ip-address":    d.Get("ip_address").(string),
		"port":          d.Get("port").(string),
		"timeout":       d.Get("timeout").(string),
		"retries":       d.Get("retries").(string),
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
