package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFKerberosServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFKerberosServersCreate,
		Read:   resourceCudaWAFKerberosServersRead,
		Update: resourceCudaWAFKerberosServersUpdate,
		Delete: resourceCudaWAFKerberosServersDelete,

		Schema: map[string]*schema.Schema{
			"kdc_name":     {Type: schema.TypeString, Required: true},
			"domain_alias": {Type: schema.TypeString, Optional: true},
			"ip_address":   {Type: schema.TypeString, Required: true},
			"port":         {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFKerberosServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFKerberosServersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFKerberosServersRead(d, m)
}

func resourceCudaWAFKerberosServersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers"
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

func resourceCudaWAFKerberosServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFKerberosServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFKerberosServersRead(d, m)
}

func resourceCudaWAFKerberosServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/kerberos-services/" + d.Get("parent.0").(string) + "/kerberos-servers/"
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

func hydrateBarracudaWAFKerberosServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"kdc-name":     d.Get("kdc_name").(string),
		"domain-alias": d.Get("domain_alias").(string),
		"ip-address":   d.Get("ip_address").(string),
		"port":         d.Get("port").(string),
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
