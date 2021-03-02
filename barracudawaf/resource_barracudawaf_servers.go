package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServersCreate,
		Read:   resourceCudaWAFServersRead,
		Update: resourceCudaWAFServersUpdate,
		Delete: resourceCudaWAFServersDelete,

		Schema: map[string]*schema.Schema{
			"address_version": {Type: schema.TypeString, Optional: true},
			"comments":        {Type: schema.TypeString, Optional: true},
			"name":            {Type: schema.TypeString, Optional: true},
			"hostname":        {Type: schema.TypeString, Optional: true},
			"identifier":      {Type: schema.TypeString, Optional: true},
			"ip_address":      {Type: schema.TypeString, Optional: true},
			"port":            {Type: schema.TypeString, Optional: true},
			"status":          {Type: schema.TypeString, Optional: true},
			"resolved_ips":    {Type: schema.TypeString, Optional: true},
			"in_band_health_checks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_other_failure":   {Type: schema.TypeString, Optional: true},
						"max_refused":         {Type: schema.TypeString, Optional: true},
						"max_timeout_failure": {Type: schema.TypeString, Optional: true},
						"max_http_errors":     {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"redirect": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
			"advanced_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_impersonation":         {Type: schema.TypeString, Optional: true},
						"source_ip_to_connect":         {Type: schema.TypeString, Optional: true},
						"max_connections":              {Type: schema.TypeString, Optional: true},
						"max_establishing_connections": {Type: schema.TypeString, Optional: true},
						"max_requests":                 {Type: schema.TypeString, Optional: true},
						"max_keepalive_requests":       {Type: schema.TypeString, Optional: true},
						"max_spare_connections":        {Type: schema.TypeString, Optional: true},
						"timeout":                      {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"load_balancing": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_server": {Type: schema.TypeString, Optional: true},
						"weight":        {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"ssl_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_certificate":            {Type: schema.TypeString, Optional: true},
						"enable_ssl_compatibility_mode": {Type: schema.TypeString, Optional: true},
						"validate_certificate":          {Type: schema.TypeString, Optional: true},
						"enable_https":                  {Type: schema.TypeString, Optional: true},
						"enable_sni":                    {Type: schema.TypeString, Optional: true},
						"enable_ssl_3":                  {Type: schema.TypeString, Optional: true},
						"enable_tls_1":                  {Type: schema.TypeString, Optional: true},
						"enable_tls_1_1":                {Type: schema.TypeString, Optional: true},
						"enable_tls_1_2":                {Type: schema.TypeString, Optional: true},
						"enable_tls_1_3":                {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"connection_pooling": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keepalive_timeout":         {Type: schema.TypeString, Optional: true},
						"enable_connection_pooling": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"out_of_band_health_checks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval":                 {Type: schema.TypeString, Optional: true},
						"enable_oob_health_checks": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"application_layer_health_checks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_headers":   {Type: schema.TypeString, Optional: true},
						"match_content_string": {Type: schema.TypeString, Optional: true},
						"method":               {Type: schema.TypeString, Optional: true},
						"domain":               {Type: schema.TypeString, Optional: true},
						"status_code":          {Type: schema.TypeString, Optional: true},
						"url":                  {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/servers"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServersResource(d, "post", resourceEndpoint),
	)

	client.hydrateBarracudaWAFServersSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFServersRead(d, m)
}

func resourceCudaWAFServersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/servers"
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

func resourceCudaWAFServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/servers"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFServersSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFServersRead(d, m)
}

func resourceCudaWAFServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/servers"
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

func hydrateBarracudaWAFServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"address-version": d.Get("address_version").(string),
		"comments":        d.Get("comments").(string),
		"name":            d.Get("name").(string),
		"hostname":        d.Get("hostname").(string),
		"identifier":      d.Get("identifier").(string),
		"ip-address":      d.Get("ip_address").(string),
		"port":            d.Get("port").(string),
		"status":          d.Get("status").(string),
		"resolved-ips":    d.Get("resolved_ips").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"address-version", "resolved-ips"}
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

func (b *BarracudaWAF) hydrateBarracudaWAFServersSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {
	subResourceObjects := map[string][]string{
		"in_band_health_checks": {
			"max_other_failure",
			"max_refused",
			"max_timeout_failure",
			"max_http_errors",
		},
		"redirect": {},
		"advanced_configuration": {
			"client_impersonation",
			"source_ip_to_connect",
			"max_connections",
			"max_establishing_connections",
			"max_requests",
			"max_keepalive_requests",
			"max_spare_connections",
			"timeout",
		},
		"load_balancing": {"backup_server", "weight"},
		"ssl_policy": {
			"client_certificate",
			"enable_ssl_compatibility_mode",
			"validate_certificate",
			"enable_https",
			"enable_sni",
			"enable_ssl_3",
			"enable_tls_1",
			"enable_tls_1_1",
			"enable_tls_1_2",
			"enable_tls_1_3",
		},
		"connection_pooling":        {"keepalive_timeout", "enable_connection_pooling"},
		"out_of_band_health_checks": {"interval", "enable_oob_health_checks"},
		"application_layer_health_checks": {
			"additional_headers",
			"match_content_string",
			"method",
			"domain",
			"status_code",
			"url",
		},
	}

	for subResource, subResourceParams := range subResourceObjects {
		subResourceParamsLength := d.Get(subResource + ".#").(int)

		if subResourceParamsLength > 0 {
			log.Printf("[INFO] Updating Barracuda WAF sub resource (%s) (%s)", name, subResource)

			for i := 0; i < subResourceParamsLength; i++ {
				subResourcePayload := map[string]string{}
				suffix := fmt.Sprintf(".%d", i)

				for _, param := range subResourceParams {
					paramSuffix := fmt.Sprintf(".%s", param)
					paramVaule := d.Get(subResource + suffix + paramSuffix).(string)

					param = strings.Replace(param, "_", "-", -1)
					subResourcePayload[param] = paramVaule
				}

				for key, val := range subResourcePayload {
					if len(val) == 0 {
						delete(subResourcePayload, key)
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
	}

	return nil
}
