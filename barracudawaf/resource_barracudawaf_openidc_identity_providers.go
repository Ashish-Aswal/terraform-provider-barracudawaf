package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFOpenidcIdentityProviders() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFOpenidcIdentityProvidersCreate,
		Read:   resourceCudaWAFOpenidcIdentityProvidersRead,
		Update: resourceCudaWAFOpenidcIdentityProvidersUpdate,
		Delete: resourceCudaWAFOpenidcIdentityProvidersDelete,

		Schema: map[string]*schema.Schema{
			"name":                   {Type: schema.TypeString, Optional: true},
			"auth_endpoint":          {Type: schema.TypeString, Optional: true},
			"client_id":              {Type: schema.TypeString, Required: true},
			"client_secret":          {Type: schema.TypeString, Required: true},
			"endpoint_configuration": {Type: schema.TypeString, Optional: true},
			"openidc_issuer":         {Type: schema.TypeString, Optional: true},
			"jwks_url":               {Type: schema.TypeString, Required: true},
			"metadata_url":           {Type: schema.TypeString, Optional: true},
			"scope":                  {Type: schema.TypeString, Optional: true},
			"token_endpoint":         {Type: schema.TypeString, Required: true},
			"userinfo_endpoint":      {Type: schema.TypeString, Optional: true},
			"type_openidc":           {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadOpenidcIdentityProviders(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"name":                   d.Get("name").(string),
		"auth-endpoint":          d.Get("auth_endpoint").(string),
		"client-id":              d.Get("client_id").(string),
		"client-secret":          d.Get("client_secret").(string),
		"endpoint-configuration": d.Get("endpoint_configuration").(string),
		"openidc-issuer":         d.Get("openidc_issuer").(string),
		"jwks-url":               d.Get("jwks_url").(string),
		"metadata-url":           d.Get("metadata_url").(string),
		"scope":                  d.Get("scope").(string),
		"token-endpoint":         d.Get("token_endpoint").(string),
		"userinfo-endpoint":      d.Get("userinfo_endpoint").(string),
		"type-openidc":           d.Get("type_openidc").(string),
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

func resourceCudaWAFOpenidcIdentityProvidersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers"
	resourceCreateResponseError := makeRestAPIPayloadOpenidcIdentityProviders(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFOpenidcIdentityProvidersRead(d, m)
}

func resourceCudaWAFOpenidcIdentityProvidersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFOpenidcIdentityProvidersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadOpenidcIdentityProviders(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFOpenidcIdentityProvidersRead(d, m)
}

func resourceCudaWAFOpenidcIdentityProvidersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadOpenidcIdentityProviders(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
