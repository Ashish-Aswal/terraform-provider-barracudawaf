package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAdministratorRoles() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAdministratorRolesCreate,
		Read:   resourceCudaWAFAdministratorRolesRead,
		Update: resourceCudaWAFAdministratorRolesUpdate,
		Delete: resourceCudaWAFAdministratorRolesDelete,

		Schema: map[string]*schema.Schema{
			"api_privilege":           {Type: schema.TypeString, Optional: true},
			"authentication_services": {Type: schema.TypeString, Optional: true},
			"name":                    {Type: schema.TypeString, Required: true},
			"objects":                 {Type: schema.TypeString, Optional: true},
			"operations":              {Type: schema.TypeString, Optional: true},
			"security_policies":       {Type: schema.TypeString, Optional: true},
			"service_groups":          {Type: schema.TypeString, Optional: true},
			"services":                {Type: schema.TypeString, Optional: true},
			"role_type":               {Type: schema.TypeString, Optional: true},
			"vsites":                  {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFAdministratorRolesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/administrator-roles"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAdministratorRolesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFAdministratorRolesRead(d, m)
}

func resourceCudaWAFAdministratorRolesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAdministratorRolesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/administrator-roles/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAdministratorRolesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAdministratorRolesRead(d, m)
}

func resourceCudaWAFAdministratorRolesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/administrator-roles/"
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

func hydrateBarracudaWAFAdministratorRolesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"api-privilege":           d.Get("api_privilege").(string),
		"authentication-services": d.Get("authentication_services").(string),
		"name":                    d.Get("name").(string),
		"objects":                 d.Get("objects").(string),
		"operations":              d.Get("operations").(string),
		"security-policies":       d.Get("security_policies").(string),
		"service-groups":          d.Get("service_groups").(string),
		"services":                d.Get("services").(string),
		"role-type":               d.Get("role_type").(string),
		"vsites":                  d.Get("vsites").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"role-type"}
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
