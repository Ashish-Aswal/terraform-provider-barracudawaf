package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadLdapServers(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{}
		for item := range updatePayloadExceptions {
			delete(resourcePayload, updatePayloadExceptions[item])
		}
	}

	//sanitise the resource payload
	for key, val := range resourcePayload {
		if len(val) <= 0 {
			delete(resourcePayload, key)
		}
	}

	//resourceUpdateData : cudaWAF reource URI update data
	resourceUpdateData := map[string]interface{}{
		"endpoint":  resourceEndpoint,
		"payload":   resourcePayload,
		"operation": resourceOperation,
		"name":      d.Get("name").(string),
	}

	//updateCudaWAFResourceObject : update cudaWAF resource object
	resourceUpdateStatus, resourceUpdateResponseBody := updateCudaWAFResourceObject(
		resourceUpdateData,
	)

	if resourceUpdateStatus == 200 || resourceUpdateStatus == 201 {
		if resourceOperation != "DELETE" {
			d.SetId(resourceUpdateResponseBody["id"].(string))
		}
	} else {
		return fmt.Errorf("some error occurred : %v", resourceUpdateResponseBody["msg"])
	}

	return nil
}

func resourceCudaWAFLdapServersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/ldap-services/" + d.Get("parent.0").(string) + "/ldap-servers"
	resourceCreateResponseError := makeRestAPIPayloadLdapServers(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFLdapServersRead(d, m)
}

func resourceCudaWAFLdapServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFLdapServersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/ldap-services/" + d.Get("parent.0").(string) + "/ldap-servers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadLdapServers(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFLdapServersRead(d, m)
}

func resourceCudaWAFLdapServersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/ldap-services/" + d.Get("parent.0").(string) + "/ldap-servers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadLdapServers(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
