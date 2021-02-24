package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFUrlPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFUrlPoliciesCreate,
		Read:   resourceCudaWAFUrlPoliciesRead,
		Update: resourceCudaWAFUrlPoliciesUpdate,
		Delete: resourceCudaWAFUrlPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"enable_data_theft_protection":           {Type: schema.TypeString, Optional: true},
			"enable_batd_scan":                       {Type: schema.TypeString, Optional: true},
			"comments":                               {Type: schema.TypeString, Optional: true},
			"host":                                   {Type: schema.TypeString, Optional: true},
			"extended_match":                         {Type: schema.TypeString, Optional: true},
			"extended_match_sequence":                {Type: schema.TypeString, Optional: true},
			"mode":                                   {Type: schema.TypeString, Optional: true},
			"name":                                   {Type: schema.TypeString, Required: true},
			"parse_urls_in_scripts":                  {Type: schema.TypeString, Optional: true},
			"rate_control_pool":                      {Type: schema.TypeString, Optional: true},
			"response_charset":                       {Type: schema.TypeString, Optional: true},
			"status":                                 {Type: schema.TypeString, Optional: true},
			"url":                                    {Type: schema.TypeString, Required: true},
			"enable_virus_scan":                      {Type: schema.TypeString, Optional: true},
			"web_scraping_policy":                    {Type: schema.TypeString, Optional: true},
			"counting_criterion":                     {Type: schema.TypeString, Optional: true},
			"enable_count_auth_resp_code":            {Type: schema.TypeString, Optional: true},
			"exception_clients":                      {Type: schema.TypeString, Optional: true},
			"exception_fingerprints":                 {Type: schema.TypeString, Optional: true},
			"max_allowed_accesses_per_ip":            {Type: schema.TypeString, Optional: true},
			"max_allowed_accesses_per_fingerprint":   {Type: schema.TypeString, Optional: true},
			"max_allowed_accesses_from_all_sources":  {Type: schema.TypeString, Optional: true},
			"max_bandwidth_per_ip":                   {Type: schema.TypeString, Required: true},
			"max_bandwidth_per_fingerprint":          {Type: schema.TypeString, Required: true},
			"max_bandwidth_from_all_sources":         {Type: schema.TypeString, Required: true},
			"max_failed_accesses_per_ip":             {Type: schema.TypeString, Optional: true},
			"max_failed_accesses_per_fingerprint":    {Type: schema.TypeString, Optional: true},
			"max_failed_accesses_from_all_sources":   {Type: schema.TypeString, Optional: true},
			"count_window":                           {Type: schema.TypeString, Optional: true},
			"enable_bruteforce_prevention":           {Type: schema.TypeString, Optional: true},
			"credential_stuffing_username_field":     {Type: schema.TypeString, Optional: true},
			"credential_stuffing_password_field":     {Type: schema.TypeString, Optional: true},
			"credential_spraying_blocking_threshold": {Type: schema.TypeString, Optional: true},
			"credential_protection_type":             {Type: schema.TypeString, Optional: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceCudaWAFUrlPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-policies"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFUrlPoliciesResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFUrlPoliciesRead(d, m)
}

func resourceCudaWAFUrlPoliciesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-policies"
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

func resourceCudaWAFUrlPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-policies/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFUrlPoliciesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFUrlPoliciesRead(d, m)
}

func resourceCudaWAFUrlPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/url-policies/"
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

func hydrateBarracudaWAFUrlPoliciesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"enable-data-theft-protection":           d.Get("enable_data_theft_protection").(string),
		"enable-batd-scan":                       d.Get("enable_batd_scan").(string),
		"comments":                               d.Get("comments").(string),
		"host":                                   d.Get("host").(string),
		"extended-match":                         d.Get("extended_match").(string),
		"extended-match-sequence":                d.Get("extended_match_sequence").(string),
		"mode":                                   d.Get("mode").(string),
		"name":                                   d.Get("name").(string),
		"parse-urls-in-scripts":                  d.Get("parse_urls_in_scripts").(string),
		"rate-control-pool":                      d.Get("rate_control_pool").(string),
		"response-charset":                       d.Get("response_charset").(string),
		"status":                                 d.Get("status").(string),
		"url":                                    d.Get("url").(string),
		"enable-virus-scan":                      d.Get("enable_virus_scan").(string),
		"web-scraping-policy":                    d.Get("web_scraping_policy").(string),
		"counting-criterion":                     d.Get("counting_criterion").(string),
		"enable-count-auth-resp-code":            d.Get("enable_count_auth_resp_code").(string),
		"exception-clients":                      d.Get("exception_clients").(string),
		"exception-fingerprints":                 d.Get("exception_fingerprints").(string),
		"max-allowed-accesses-per-ip":            d.Get("max_allowed_accesses_per_ip").(string),
		"max-allowed-accesses-per-fingerprint":   d.Get("max_allowed_accesses_per_fingerprint").(string),
		"max-allowed-accesses-from-all-sources":  d.Get("max_allowed_accesses_from_all_sources").(string),
		"max-bandwidth-per-ip":                   d.Get("max_bandwidth_per_ip").(string),
		"max-bandwidth-per-fingerprint":          d.Get("max_bandwidth_per_fingerprint").(string),
		"max-bandwidth-from-all-sources":         d.Get("max_bandwidth_from_all_sources").(string),
		"max-failed-accesses-per-ip":             d.Get("max_failed_accesses_per_ip").(string),
		"max-failed-accesses-per-fingerprint":    d.Get("max_failed_accesses_per_fingerprint").(string),
		"max-failed-accesses-from-all-sources":   d.Get("max_failed_accesses_from_all_sources").(string),
		"count-window":                           d.Get("count_window").(string),
		"enable-bruteforce-prevention":           d.Get("enable_bruteforce_prevention").(string),
		"credential-stuffing-username-field":     d.Get("credential_stuffing_username_field").(string),
		"credential-stuffing-password-field":     d.Get("credential_stuffing_password_field").(string),
		"credential-spraying-blocking-threshold": d.Get("credential_spraying_blocking_threshold").(string),
		"credential-protection-type":             d.Get("credential_protection_type").(string),
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
