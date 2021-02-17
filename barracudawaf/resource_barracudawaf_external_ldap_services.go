package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadExternalLdapServices(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

func resourceCudaWAFExternalLdapServicesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/external-ldap-services"
	resourceCreateResponseError := makeRestAPIPayloadExternalLdapServices(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFExternalLdapServicesRead(d, m)
}

func resourceCudaWAFExternalLdapServicesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFExternalLdapServicesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/external-ldap-services/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadExternalLdapServices(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFExternalLdapServicesRead(d, m)
}

func resourceCudaWAFExternalLdapServicesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/external-ldap-services/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadExternalLdapServices(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
