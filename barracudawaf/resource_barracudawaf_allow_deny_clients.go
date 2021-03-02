package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAllowDenyClients() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAllowDenyClientsCreate,
		Read:   resourceCudaWAFAllowDenyClientsRead,
		Update: resourceCudaWAFAllowDenyClientsUpdate,
		Delete: resourceCudaWAFAllowDenyClientsDelete,

		Schema: map[string]*schema.Schema{
			"name":                {Type: schema.TypeString, Required: true},
			"action":              {Type: schema.TypeString, Optional: true},
			"sequence":            {Type: schema.TypeString, Optional: true},
			"certificate_serial":  {Type: schema.TypeString, Optional: true},
			"common_name":         {Type: schema.TypeString, Optional: true},
			"country":             {Type: schema.TypeString, Optional: true},
			"locality":            {Type: schema.TypeString, Optional: true},
			"organization":        {Type: schema.TypeString, Optional: true},
			"organizational_unit": {Type: schema.TypeString, Optional: true},
			"state":               {Type: schema.TypeString, Optional: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFAllowDenyClientsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAllowDenyClientsResource(d, "post", resourceEndpoint),
	)

	client.hydrateBarracudaWAFAllowDenyClientsSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFAllowDenyClientsRead(d, m)
}

func resourceCudaWAFAllowDenyClientsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients"
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

func resourceCudaWAFAllowDenyClientsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAllowDenyClientsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFAllowDenyClientsSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAllowDenyClientsRead(d, m)
}

func resourceCudaWAFAllowDenyClientsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/allow-deny-clients"
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

func hydrateBarracudaWAFAllowDenyClientsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"name":                d.Get("name").(string),
		"action":              d.Get("action").(string),
		"sequence":            d.Get("sequence").(string),
		"certificate-serial":  d.Get("certificate_serial").(string),
		"common-name":         d.Get("common_name").(string),
		"country":             d.Get("country").(string),
		"locality":            d.Get("locality").(string),
		"organization":        d.Get("organization").(string),
		"organizational-unit": d.Get("organizational_unit").(string),
		"state":               d.Get("state").(string),
		"status":              d.Get("status").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{}
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

func (b *BarracudaWAF) hydrateBarracudaWAFAllowDenyClientsSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {
	subResourceObjects := map[string][]string{}

	for subResource, subResourceParams := range subResourceObjects {
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
