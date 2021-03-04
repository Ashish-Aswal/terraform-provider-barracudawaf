package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	subResourceAutoSystemAclsParams = map[string][]string{}
)

func resourceCudaWAFAutoSystemAcls() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAutoSystemAclsCreate,
		Read:   resourceCudaWAFAutoSystemAclsRead,
		Update: resourceCudaWAFAutoSystemAclsUpdate,
		Delete: resourceCudaWAFAutoSystemAclsDelete,

		Schema: map[string]*schema.Schema{
			"interface":           {Type: schema.TypeString, Optional: true},
			"ip_version":          {Type: schema.TypeString, Optional: true},
			"acl_type":            {Type: schema.TypeString, Optional: true},
			"action":              {Type: schema.TypeString, Optional: true},
			"destination_port":    {Type: schema.TypeString, Required: true},
			"source_address":      {Type: schema.TypeString, Optional: true},
			"source_netmask":      {Type: schema.TypeString, Optional: true},
			"enable_logging":      {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"priority":            {Type: schema.TypeString, Required: true},
			"protocol":            {Type: schema.TypeString, Optional: true},
			"source_port":         {Type: schema.TypeString, Required: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"destination_address": {Type: schema.TypeString, Optional: true},
			"destination_netmask": {Type: schema.TypeString, Optional: true},
			"vsite":               {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFAutoSystemAclsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/auto-system-acls"
	client.CreateBarracudaWAFResource(name, hydrateBarracudaWAFAutoSystemAclsResource(d, "post", resourceEndpoint))

	client.hydrateBarracudaWAFAutoSystemAclsSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFAutoSystemAclsRead(d, m)
}

func resourceCudaWAFAutoSystemAclsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/auto-system-acls"
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

func resourceCudaWAFAutoSystemAclsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/auto-system-acls"
	err := client.UpdateBarracudaWAFResource(name, hydrateBarracudaWAFAutoSystemAclsResource(d, "put", resourceEndpoint))

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFAutoSystemAclsSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAutoSystemAclsRead(d, m)
}

func resourceCudaWAFAutoSystemAclsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/auto-system-acls"
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

func hydrateBarracudaWAFAutoSystemAclsResource(d *schema.ResourceData, method string, endpoint string) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"interface":           d.Get("interface").(string),
		"ip-version":          d.Get("ip_version").(string),
		"acl-type":            d.Get("acl_type").(string),
		"action":              d.Get("action").(string),
		"destination-port":    d.Get("destination_port").(string),
		"source-address":      d.Get("source_address").(string),
		"source-netmask":      d.Get("source_netmask").(string),
		"enable-logging":      d.Get("enable_logging").(string),
		"name":                d.Get("name").(string),
		"priority":            d.Get("priority").(string),
		"protocol":            d.Get("protocol").(string),
		"source-port":         d.Get("source_port").(string),
		"status":              d.Get("status").(string),
		"destination-address": d.Get("destination_address").(string),
		"destination-netmask": d.Get("destination_netmask").(string),
		"vsite":               d.Get("vsite").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{
			"interface",
			"ip-version",
			"acl-type",
			"action",
			"destination-port",
			"name",
			"priority",
			"protocol",
			"vsite",
		}
		for _, param := range updatePayloadExceptions {
			delete(resourcePayload, param)
		}
	}

	// remove empty parameters from resource payload
	for key, val := range resourcePayload {
		if len(val) == 0 {
			delete(resourcePayload, key)
		}
	}

	return &APIRequest{
		URL:  endpoint,
		Body: resourcePayload,
	}
}

func (b *BarracudaWAF) hydrateBarracudaWAFAutoSystemAclsSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {

	for subResource, subResourceParams := range subResourceAutoSystemAclsParams {
		subResourceParamsLength := d.Get(subResource + ".#").(int)

		if subResourceParamsLength > 0 {
			log.Printf("[INFO] Updating Barracuda WAF sub resource (%s) (%s)", name, subResource)

			for i := 0; i < subResourceParamsLength; i++ {
				subResourcePayload := map[string]string{}
				suffix := fmt.Sprintf(".%d", i)

				for _, param := range subResourceParams {
					paramSuffix := fmt.Sprintf(".%s", param)
					paramVaule := d.Get(subResource + suffix + paramSuffix).(string)

					param = strings.Replace(param, "_", "-", -1)
					subResourcePayload[param] = paramVaule
				}

				for key, val := range subResourcePayload {
					if len(val) == 0 {
						delete(subResourcePayload, key)
					}
				}

				err := b.UpdateBarracudaWAFSubResource(name, endpoint, &APIRequest{
					URL:  strings.Replace(subResource, "_", "-", -1),
					Body: subResourcePayload,
				})

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
