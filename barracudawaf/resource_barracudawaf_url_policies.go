package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadUrlPolicies(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

func resourceCudaWAFUrlPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-policies"
	resourceCreateResponseError := makeRestAPIPayloadUrlPolicies(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFUrlPoliciesRead(d, m)
}

func resourceCudaWAFUrlPoliciesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFUrlPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-policies/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadUrlPolicies(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFUrlPoliciesRead(d, m)
}

func resourceCudaWAFUrlPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/services/" + d.Get("parent.0").(string) + "/url-policies/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadUrlPolicies(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
