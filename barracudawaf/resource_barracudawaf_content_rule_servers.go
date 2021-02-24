package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFContentRuleServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFContentRuleServersCreate,
		Read:   resourceCudaWAFContentRuleServersRead,
		Update: resourceCudaWAFContentRuleServersUpdate,
		Delete: resourceCudaWAFContentRuleServersDelete,

		Schema: map[string]*schema.Schema{
			"comments":        {Type: schema.TypeString, Optional: true},
			"name":            {Type: schema.TypeString, Optional: true},
			"hostname":        {Type: schema.TypeString, Optional: true},
			"identifier":      {Type: schema.TypeString, Optional: true},
			"ip_address":      {Type: schema.TypeString, Optional: true},
			"address_version": {Type: schema.TypeString, Optional: true},
			"port":            {Type: schema.TypeString, Optional: true},
			"status":          {Type: schema.TypeString, Optional: true},
			"resolved_ips":    {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFContentRuleServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/content-rules/" + d.Get("parent.1").(string) + "/content-rule-servers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFContentRuleServersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFContentRuleServersRead(d, m)
}

func resourceCudaWAFContentRuleServersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/content-rules/" + d.Get("parent.1").(string) + "/content-rule-servers"
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

func resourceCudaWAFContentRuleServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/content-rules/" + d.Get("parent.1").(string) + "/content-rule-servers/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFContentRuleServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFContentRuleServersRead(d, m)
}

func resourceCudaWAFContentRuleServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/content-rules/" + d.Get("parent.1").(string) + "/content-rule-servers/"
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

func hydrateBarracudaWAFContentRuleServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comments":        d.Get("comments").(string),
		"name":            d.Get("name").(string),
		"hostname":        d.Get("hostname").(string),
		"identifier":      d.Get("identifier").(string),
		"ip-address":      d.Get("ip_address").(string),
		"address-version": d.Get("address_version").(string),
		"port":            d.Get("port").(string),
		"status":          d.Get("status").(string),
		"resolved-ips":    d.Get("resolved_ips").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"address-version", "resolved-ips"}
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
