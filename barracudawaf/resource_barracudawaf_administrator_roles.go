package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadAdministratorRoles(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"role-type"}
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

func resourceCudaWAFAdministratorRolesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/administrator-roles"
	resourceCreateResponseError := makeRestAPIPayloadAdministratorRoles(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFAdministratorRolesRead(d, m)
}

func resourceCudaWAFAdministratorRolesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAdministratorRolesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/administrator-roles/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadAdministratorRoles(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFAdministratorRolesRead(d, m)
}

func resourceCudaWAFAdministratorRolesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/administrator-roles/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadAdministratorRoles(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
