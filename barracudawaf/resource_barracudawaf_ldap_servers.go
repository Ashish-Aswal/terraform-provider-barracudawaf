package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFLdapServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFLdapServersCreate,
		Read:   resourceCudaWAFLdapServersRead,
		Update: resourceCudaWAFLdapServersUpdate,
		Delete: resourceCudaWAFLdapServersDelete,

		Schema: map[string]*schema.Schema{
			"domain_alias":           {Type: schema.TypeString, Optional: true},
			"ip_address":             {Type: schema.TypeString, Required: true},
			"port":                   {Type: schema.TypeString, Optional: true},
			"allow_nested_groups":    {Type: schema.TypeString, Optional: true},
			"base_dn":                {Type: schema.TypeString, Required: true},
			"bind_dn":                {Type: schema.TypeString, Required: true},
			"bind_password":          {Type: schema.TypeString, Optional: true},
			"group_filter":           {Type: schema.TypeString, Optional: true},
			"login_attribute":        {Type: schema.TypeString, Optional: true},
			"group_name_attribute":   {Type: schema.TypeString, Optional: true},
			"query_for_group":        {Type: schema.TypeString, Optional: true},
			"secure_connection_type": {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFLdapServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/ldap-services/" + d.Get("parent.0").(string) + "/ldap-servers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFLdapServersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFLdapServersRead(d, m)
}

func resourceCudaWAFLdapServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFLdapServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/ldap-services/" + d.Get("parent.0").(string) + "/ldap-servers/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFLdapServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFLdapServersRead(d, m)
}

func resourceCudaWAFLdapServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/ldap-services/" + d.Get("parent.0").(string) + "/ldap-servers/"
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

func hydrateBarracudaWAFLdapServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"domain-alias":           d.Get("domain_alias").(string),
		"ip-address":             d.Get("ip_address").(string),
		"port":                   d.Get("port").(string),
		"allow-nested-groups":    d.Get("allow_nested_groups").(string),
		"base-dn":                d.Get("base_dn").(string),
		"bind-dn":                d.Get("bind_dn").(string),
		"bind-password":          d.Get("bind_password").(string),
		"group-filter":           d.Get("group_filter").(string),
		"login-attribute":        d.Get("login_attribute").(string),
		"group-name-attribute":   d.Get("group_name_attribute").(string),
		"query-for-group":        d.Get("query_for_group").(string),
		"secure-connection-type": d.Get("secure_connection_type").(string),
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
