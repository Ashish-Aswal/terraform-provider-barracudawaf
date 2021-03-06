package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	subResourceOpenidcIdentityProvidersParams = map[string][]string{}
)

func resourceCudaWAFOpenidcIdentityProviders() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFOpenidcIdentityProvidersCreate,
		Read:   resourceCudaWAFOpenidcIdentityProvidersRead,
		Update: resourceCudaWAFOpenidcIdentityProvidersUpdate,
		Delete: resourceCudaWAFOpenidcIdentityProvidersDelete,

		Schema: map[string]*schema.Schema{
			"name":          {Type: schema.TypeString, Optional: true, Description: "OpenID Connect Alias"},
			"auth_endpoint": {Type: schema.TypeString, Optional: true, Description: "Auth Endpoint"},
			"client_id":     {Type: schema.TypeString, Required: true, Description: "Identity Provider Name"},
			"client_secret": {Type: schema.TypeString, Required: true, Description: "Server IP"},
			"endpoint_configuration": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identity Provider Metadata Type",
			},
			"openidc_issuer": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the metadata file being uploaded.",
			},
			"jwks_url": {Type: schema.TypeString, Required: true, Description: "Type"},
			"metadata_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Please specify the OpenId Connect authorization endpoint",
			},
			"scope": {Type: schema.TypeString, Optional: true, Description: "Metadata URL"},
			"token_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identity Provider Metadata Type",
			},
			"userinfo_endpoint": {Type: schema.TypeString, Optional: true, Description: "Type"},
			"type_openidc":      {Type: schema.TypeString, Optional: true, Description: "Type"},
			"parent":            {Type: schema.TypeList, Elem: &schema.Schema{Type: schema.TypeString}, Required: true},
		},

		Description: "`barracudawaf_openidc_identity_providers` manages `Openidc Identity Providers` on the Barracuda Web Application Firewall.",
	}
}

func resourceCudaWAFOpenidcIdentityProvidersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers"
	err := client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFOpenidcIdentityProvidersResource(d, "post", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to create Barracuda WAF resource (%s) (%v) ", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFOpenidcIdentityProvidersSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to create Barracuda WAF sub resource (%s) (%v) ", name, err)
		return err
	}

	d.SetId(name)
	return resourceCudaWAFOpenidcIdentityProvidersRead(d, m)
}

func resourceCudaWAFOpenidcIdentityProvidersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers"
	request := &APIRequest{
		Method: "get",
		URL:    resourceEndpoint,
	}

	var dataItems map[string]interface{}
	resources, err := client.GetBarracudaWAFResource(name, request)

	if err != nil {
		log.Printf("[ERROR] Unable to Retrieve Barracuda WAF resource (%s) (%v) ", name, err)
		return err
	}

	if resources.Data == nil {
		log.Printf("[WARN] Barracuda WAF resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	for _, dataItems = range resources.Data {
		if dataItems["name"] == name {
			break
		}
	}

	if dataItems["name"] != name {
		return fmt.Errorf("Barracuda WAF resource (%s) not found on the system", name)
	}

	d.Set("name", name)
	return nil
}

func resourceCudaWAFOpenidcIdentityProvidersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFOpenidcIdentityProvidersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFOpenidcIdentityProvidersSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFOpenidcIdentityProvidersRead(d, m)
}

func resourceCudaWAFOpenidcIdentityProvidersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/openidc-services/" + d.Get("parent.0").(string) + "/openidc-identity-providers"
	request := &APIRequest{
		Method: "delete",
		URL:    resourceEndpoint,
	}

	err := client.DeleteBarracudaWAFResource(name, request)

	if err != nil {
		return fmt.Errorf("Unable to delete the Barracuda WAF resource (%s) (%v)", name, err)
	}

	return nil
}

func hydrateBarracudaWAFOpenidcIdentityProvidersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
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

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{}
		for _, param := range updatePayloadExceptions {
			delete(resourcePayload, param)
		}
	}

	// remove empty parameters from resource payload
	for key, val := range resourcePayload {
		if len(val) == 0 {
			delete(resourcePayload, key)
		}
	}

	return &APIRequest{
		URL:  endpoint,
		Body: resourcePayload,
	}
}

func (b *BarracudaWAF) hydrateBarracudaWAFOpenidcIdentityProvidersSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {

	for subResource, subResourceParams := range subResourceOpenidcIdentityProvidersParams {
		subResourceParamsLength := d.Get(subResource + ".#").(int)

		log.Printf("[INFO] Updating Barracuda WAF sub resource (%s) (%s)", name, subResource)

		for i := 0; i < subResourceParamsLength; i++ {
			subResourcePayload := map[string]string{}
			suffix := fmt.Sprintf(".%d", i)

			for _, param := range subResourceParams {
				paramSuffix := fmt.Sprintf(".%s", param)
				paramVaule := d.Get(subResource + suffix + paramSuffix).(string)

				if len(paramVaule) > 0 {
					param = strings.Replace(param, "_", "-", -1)
					subResourcePayload[param] = paramVaule
				}
			}

			err := b.UpdateBarracudaWAFSubResource(name, endpoint, &APIRequest{
				URL:  strings.Replace(subResource, "_", "-", -1),
				Body: subResourcePayload,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}
