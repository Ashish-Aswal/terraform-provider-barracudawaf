package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServiceGroups() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServiceGroupsCreate,
		Read:   resourceCudaWAFServiceGroupsRead,
		Update: resourceCudaWAFServiceGroupsUpdate,
		Delete: resourceCudaWAFServiceGroupsDelete,

		Schema: map[string]*schema.Schema{
			"service_group": {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFServiceGroupsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/vsites/" + d.Get("parent.0").(string) + "/service-groups"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServiceGroupsResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFServiceGroupsRead(d, m)
}

func resourceCudaWAFServiceGroupsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServiceGroupsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/vsites/" + d.Get("parent.0").(string) + "/service-groups/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServiceGroupsResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFServiceGroupsRead(d, m)
}

func resourceCudaWAFServiceGroupsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/vsites/" + d.Get("parent.0").(string) + "/service-groups/"
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

func hydrateBarracudaWAFServiceGroupsResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{"service-group": d.Get("service_group").(string)}

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
