package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFAuthorizationPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFAuthorizationPoliciesCreate,
		Read:   resourceCudaWAFAuthorizationPoliciesRead,
		Update: resourceCudaWAFAuthorizationPoliciesUpdate,
		Delete: resourceCudaWAFAuthorizationPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"allowed_groups":                {Type: schema.TypeString, Optional: true},
			"allowed_users":                 {Type: schema.TypeString, Optional: true},
			"comments":                      {Type: schema.TypeString, Optional: true},
			"auth_context_classref":         {Type: schema.TypeString, Optional: true},
			"name":                          {Type: schema.TypeString, Required: true},
			"host":                          {Type: schema.TypeString, Optional: true},
			"extended_match":                {Type: schema.TypeString, Optional: true},
			"extended_match_sequence":       {Type: schema.TypeString, Optional: true},
			"cookie_timeout":                {Type: schema.TypeString, Optional: true},
			"access_rules":                  {Type: schema.TypeString, Optional: true},
			"enable_signing_on_authrequest": {Type: schema.TypeString, Optional: true},
			"digest_algorithm":              {Type: schema.TypeString, Optional: true},
			"status":                        {Type: schema.TypeString, Optional: true},
			"url":                           {Type: schema.TypeString, Required: true},
			"use_persistent_cookie":         {Type: schema.TypeString, Optional: true},
			"allow_any_authenticated_user":  {Type: schema.TypeString, Optional: true},
			"login_method":                  {Type: schema.TypeString, Optional: true},
			"access_denied_url":             {Type: schema.TypeString, Optional: true},
			"login_url":                     {Type: schema.TypeString, Optional: true},
			"send_basic_auth":               {Type: schema.TypeString, Optional: true},
			"send_domain_in_basic_auth":     {Type: schema.TypeString, Optional: true},
			"kerberos_spn":                  {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFAuthorizationPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/authorization-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAuthorizationPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFAuthorizationPoliciesRead(d, m)
}

func resourceCudaWAFAuthorizationPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFAuthorizationPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/authorization-policies/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFAuthorizationPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFAuthorizationPoliciesRead(d, m)
}

func resourceCudaWAFAuthorizationPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/authorization-policies/"
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

func hydrateBarracudaWAFAuthorizationPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"allowed-groups":                d.Get("allowed_groups").(string),
		"allowed-users":                 d.Get("allowed_users").(string),
		"comments":                      d.Get("comments").(string),
		"auth-context-classref":         d.Get("auth_context_classref").(string),
		"name":                          d.Get("name").(string),
		"host":                          d.Get("host").(string),
		"extended-match":                d.Get("extended_match").(string),
		"extended-match-sequence":       d.Get("extended_match_sequence").(string),
		"cookie-timeout":                d.Get("cookie_timeout").(string),
		"access-rules":                  d.Get("access_rules").(string),
		"enable-signing-on-authrequest": d.Get("enable_signing_on_authrequest").(string),
		"digest-algorithm":              d.Get("digest_algorithm").(string),
		"status":                        d.Get("status").(string),
		"url":                           d.Get("url").(string),
		"use-persistent-cookie":         d.Get("use_persistent_cookie").(string),
		"allow-any-authenticated-user":  d.Get("allow_any_authenticated_user").(string),
		"login-method":                  d.Get("login_method").(string),
		"access-denied-url":             d.Get("access_denied_url").(string),
		"login-url":                     d.Get("login_url").(string),
		"send-basic-auth":               d.Get("send_basic_auth").(string),
		"send-domain-in-basic-auth":     d.Get("send_domain_in_basic_auth").(string),
		"kerberos-spn":                  d.Get("kerberos_spn").(string),
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
