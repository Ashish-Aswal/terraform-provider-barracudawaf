package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFLocalUsers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFLocalUsersCreate,
		Read:   resourceCudaWAFLocalUsersRead,
		Update: resourceCudaWAFLocalUsersUpdate,
		Delete: resourceCudaWAFLocalUsersDelete,

		Schema: map[string]*schema.Schema{
			"user_groups": {Type: schema.TypeString, Optional: true},
			"name":        {Type: schema.TypeString, Optional: true},
			"password":    {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceCudaWAFLocalUsersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/local-users"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFLocalUsersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFLocalUsersRead(d, m)
}

func resourceCudaWAFLocalUsersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFLocalUsersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/local-users/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFLocalUsersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFLocalUsersRead(d, m)
}

func resourceCudaWAFLocalUsersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/local-users/"
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

func hydrateBarracudaWAFLocalUsersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"user-groups": d.Get("user_groups").(string),
		"name":        d.Get("name").(string),
		"password":    d.Get("password").(string),
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
