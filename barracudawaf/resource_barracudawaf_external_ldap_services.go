package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFExternalLdapServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFExternalLdapServicesCreate,
		Read:   resourceCudaWAFExternalLdapServicesRead,
		Update: resourceCudaWAFExternalLdapServicesUpdate,
		Delete: resourceCudaWAFExternalLdapServicesDelete,

		Schema: map[string]*schema.Schema{
			"ip_address":                  {Type: schema.TypeString, Required: true},
			"search_base":                 {Type: schema.TypeString, Optional: true},
			"bind_dn":                     {Type: schema.TypeString, Required: true},
			"bind_password":               {Type: schema.TypeString, Optional: true},
			"retype_bind_password":        {Type: schema.TypeString, Optional: true},
			"default_role":                {Type: schema.TypeString, Required: true},
			"role_map":                    {Type: schema.TypeString, Optional: true},
			"role_order":                  {Type: schema.TypeString, Optional: true},
			"encryption":                  {Type: schema.TypeString, Required: true},
			"group_filter":                {Type: schema.TypeString, Optional: true},
			"group_member_uid_attribute":  {Type: schema.TypeString, Required: true},
			"group_membership_format":     {Type: schema.TypeString, Optional: true},
			"group_name_attribute":        {Type: schema.TypeString, Optional: true},
			"name":                        {Type: schema.TypeString, Required: true},
			"port":                        {Type: schema.TypeString, Optional: true},
			"uid_attribute":               {Type: schema.TypeString, Optional: true},
			"allow_nested_groups":         {Type: schema.TypeString, Required: true},
			"ldap_server_type":            {Type: schema.TypeString, Optional: true},
			"validate_server_certificate": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFExternalLdapServicesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/external-ldap-services"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFExternalLdapServicesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFExternalLdapServicesRead(d, m)
}

func resourceCudaWAFExternalLdapServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFExternalLdapServicesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/external-ldap-services/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFExternalLdapServicesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFExternalLdapServicesRead(d, m)
}

func resourceCudaWAFExternalLdapServicesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/external-ldap-services/"
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

func hydrateBarracudaWAFExternalLdapServicesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"ip-address":                  d.Get("ip_address").(string),
		"search-base":                 d.Get("search_base").(string),
		"bind-dn":                     d.Get("bind_dn").(string),
		"bind-password":               d.Get("bind_password").(string),
		"retype-bind-password":        d.Get("retype_bind_password").(string),
		"default-role":                d.Get("default_role").(string),
		"role-map":                    d.Get("role_map").(string),
		"role-order":                  d.Get("role_order").(string),
		"encryption":                  d.Get("encryption").(string),
		"group-filter":                d.Get("group_filter").(string),
		"group-member-uid-attribute":  d.Get("group_member_uid_attribute").(string),
		"group-membership-format":     d.Get("group_membership_format").(string),
		"group-name-attribute":        d.Get("group_name_attribute").(string),
		"name":                        d.Get("name").(string),
		"port":                        d.Get("port").(string),
		"uid-attribute":               d.Get("uid_attribute").(string),
		"allow-nested-groups":         d.Get("allow_nested_groups").(string),
		"ldap_server_type":            d.Get("ldap_server_type").(string),
		"validate-server-certificate": d.Get("validate_server_certificate").(string),
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
