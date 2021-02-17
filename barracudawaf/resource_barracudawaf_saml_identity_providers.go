package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFSamlIdentityProviders() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSamlIdentityProvidersCreate,
		Read:   resourceCudaWAFSamlIdentityProvidersRead,
		Update: resourceCudaWAFSamlIdentityProvidersUpdate,
		Delete: resourceCudaWAFSamlIdentityProvidersDelete,

		Schema: map[string]*schema.Schema{
			"metadata_file":       {Type: schema.TypeString, Optional: true},
			"metadata_type":       {Type: schema.TypeString, Optional: true},
			"autoupdate_metadata": {Type: schema.TypeString, Optional: true},
			"metadata_url":        {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"metadata_content":    {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadSamlIdentityProviders(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"metadata-file":       d.Get("metadata_file").(string),
		"metadata-type":       d.Get("metadata_type").(string),
		"autoupdate-metadata": d.Get("autoupdate_metadata").(string),
		"metadata-url":        d.Get("metadata_url").(string),
		"name":                d.Get("name").(string),
		"metadata-content":    d.Get("metadata_content").(string),
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

func resourceCudaWAFSamlIdentityProvidersCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers"
	resourceCreateResponseError := makeRestAPIPayloadSamlIdentityProviders(
		d,
		"POST",
		resourceEndpoint,
	)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFSamlIdentityProvidersRead(d, m)
}

func resourceCudaWAFSamlIdentityProvidersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSamlIdentityProvidersUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadSamlIdentityProviders(
		d,
		"PUT",
		resourceEndpoint,
	)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFSamlIdentityProvidersRead(d, m)
}

func resourceCudaWAFSamlIdentityProvidersDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/saml-services/" + d.Get("parent.0").(string) + "/saml-identity-providers/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadSamlIdentityProviders(
		d,
		"DELETE",
		resourceEndpoint,
	)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
