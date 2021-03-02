package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFDestinationNats() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFDestinationNatsCreate,
		Read:   resourceCudaWAFDestinationNatsRead,
		Update: resourceCudaWAFDestinationNatsUpdate,
		Delete: resourceCudaWAFDestinationNatsDelete,

		Schema: map[string]*schema.Schema{
			"comments":                 {Type: schema.TypeString, Optional: true},
			"vsite":                    {Type: schema.TypeString, Required: true},
			"incoming_interface":       {Type: schema.TypeString, Required: true},
			"post_destination_address": {Type: schema.TypeString, Required: true},
			"pre_destination_address":  {Type: schema.TypeString, Required: true},
			"pre_destination_netmask":  {Type: schema.TypeString, Required: true},
			"pre_destination_port":     {Type: schema.TypeString, Optional: true},
			"protocol":                 {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFDestinationNatsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/destination-nats"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFDestinationNatsResource(d, "post", resourceEndpoint),
	)

	client.hydrateBarracudaWAFDestinationNatsSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFDestinationNatsRead(d, m)
}

func resourceCudaWAFDestinationNatsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/destination-nats"
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

func resourceCudaWAFDestinationNatsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/destination-nats"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFDestinationNatsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFDestinationNatsSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFDestinationNatsRead(d, m)
}

func resourceCudaWAFDestinationNatsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/destination-nats"
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

func hydrateBarracudaWAFDestinationNatsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"comments":                 d.Get("comments").(string),
		"vsite":                    d.Get("vsite").(string),
		"incoming-interface":       d.Get("incoming_interface").(string),
		"post-destination-address": d.Get("post_destination_address").(string),
		"pre-destination-address":  d.Get("pre_destination_address").(string),
		"pre-destination-netmask":  d.Get("pre_destination_netmask").(string),
		"pre-destination-port":     d.Get("pre_destination_port").(string),
		"protocol":                 d.Get("protocol").(string),
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

func (b *BarracudaWAF) hydrateBarracudaWAFDestinationNatsSubResource(
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
