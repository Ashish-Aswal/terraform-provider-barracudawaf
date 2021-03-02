package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServices() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServicesCreate,
		Read:   resourceCudaWAFServicesRead,
		Update: resourceCudaWAFServicesUpdate,
		Delete: resourceCudaWAFServicesDelete,

		Schema: map[string]*schema.Schema{
			"address_version":     {Type: schema.TypeString, Optional: true},
			"dps_enabled":         {Type: schema.TypeString, Optional: true},
			"mask":                {Type: schema.TypeString, Optional: true},
			"session_timeout":     {Type: schema.TypeString, Optional: true},
			"linked_service_name": {Type: schema.TypeString, Optional: true},
			"enable_access_logs":  {Type: schema.TypeString, Optional: true},
			"app_id":              {Type: schema.TypeString, Optional: true},
			"comments":            {Type: schema.TypeString, Optional: true},
			"group":               {Type: schema.TypeString, Optional: true},
			"service_id":          {Type: schema.TypeString, Optional: true},
			"ip_address":          {Type: schema.TypeString, Optional: true},
			"cloud_ip_select":     {Type: schema.TypeString, Optional: true},
			"name":                {Type: schema.TypeString, Required: true},
			"port":                {Type: schema.TypeString, Optional: true},
			"status":              {Type: schema.TypeString, Optional: true},
			"type":                {Type: schema.TypeString, Optional: true},
			"vsite":               {Type: schema.TypeString, Optional: true},
			"authentication": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_domain_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dual_authentication": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secondary_authentication_service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication_service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"access_denied_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_timeout_for_activesync": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cookie_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cookie_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_replay_protection_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"creation_timeout": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dual_login_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"idle_timeout": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_challenge_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_failed_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_processor_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"openidc_redirect_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"openidc_attribute_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"openidc_local_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_successful_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"challenge_prompt_field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"challenge_user_field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_failure_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"login_success_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logout_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logout_processor_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logout_successful_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logout_success_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_expired_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"post_processor_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"saml_logout_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"groups": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sso_cookie_update_interval": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_failed_attempts": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"count_window": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_bruteforce_prevention": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kerberos_debug_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kerberos_enable_delegation": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"kerberos_ldap_authorization": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"krb_authorization_policy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"master_service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"master_service_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_provider_display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_provider_entity_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_provider_org_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_provider_org_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"attribute_format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"attribute_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"attribute_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"attribute_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"encryption_certificate": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"signing_certificate": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"caching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiry_age":              {Type: schema.TypeString, Optional: true},
						"file_extensions":         {Type: schema.TypeString, Optional: true},
						"max_size":                {Type: schema.TypeString, Optional: true},
						"min_size":                {Type: schema.TypeString, Optional: true},
						"cache_negative_response": {Type: schema.TypeString, Optional: true},
						"ignore_request_headers":  {Type: schema.TypeString, Optional: true},
						"ignore_response_headers": {Type: schema.TypeString, Optional: true},
						"status":                  {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"clickjacking": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_origin": {Type: schema.TypeString, Optional: true},
						"options":        {Type: schema.TypeString, Optional: true},
						"status":         {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"compression": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_types":         {Type: schema.TypeString, Optional: true},
						"min_size":              {Type: schema.TypeString, Optional: true},
						"status":                {Type: schema.TypeString, Optional: true},
						"unknown_content_types": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"exception_profiling": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_profiling_trusted_host_group": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"exception_profiling_learn_from_trusted_host": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"exception_profiling_level": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ftp_security": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attack_prevention_status": {Type: schema.TypeString, Optional: true},
						"allowed_verbs":            {Type: schema.TypeString, Optional: true},
						"allowed_verb_status":      {Type: schema.TypeString, Optional: true},
						"pasv_ip_address":          {Type: schema.TypeString, Optional: true},
						"pasv_ports":               {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"instant_ssl": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status":                     {Type: schema.TypeString, Optional: true},
						"sharepoint_rewrite_support": {Type: schema.TypeString, Optional: true},
						"secure_site_domain":         {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"ip_reputation": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"anonymous_proxy":                {Type: schema.TypeString, Optional: true},
						"barracuda_reputation_blocklist": {Type: schema.TypeString, Optional: true},
						"custom_blacklisted_ip_status":   {Type: schema.TypeString, Optional: true},
						"datacenter_ip":                  {Type: schema.TypeString, Optional: true},
						"fake_crawler":                   {Type: schema.TypeString, Optional: true},
						"check_registered_country":       {Type: schema.TypeString, Optional: true},
						"block_unclassified_ips":         {Type: schema.TypeString, Optional: true},
						"apply_policy_at":                {Type: schema.TypeString, Optional: true},
						"geo_pool":                       {Type: schema.TypeString, Optional: true},
						"geoip_action":                   {Type: schema.TypeString, Optional: true},
						"enable_ip_reputation_filter":    {Type: schema.TypeString, Optional: true},
						"geoip_enable_logging":           {Type: schema.TypeString, Optional: true},
						"known_http_attack_sources":      {Type: schema.TypeString, Optional: true},
						"public_proxy":                   {Type: schema.TypeString, Optional: true},
						"satellite_provider":             {Type: schema.TypeString, Optional: true},
						"known_ssh_attack_sources":       {Type: schema.TypeString, Optional: true},
						"tor_nodes":                      {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"adaptive_profiling": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content_types":         {Type: schema.TypeString, Optional: true},
						"ignore_parameters":     {Type: schema.TypeString, Optional: true},
						"navigation_parameters": {Type: schema.TypeString, Optional: true},
						"request_learning":      {Type: schema.TypeString, Optional: true},
						"response_learning":     {Type: schema.TypeString, Optional: true},
						"status":                {Type: schema.TypeString, Optional: true},
						"trusted_host_group":    {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"sensitive_parameter_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sensitive_parameter_names": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"session_tracking": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifiers":         {Type: schema.TypeString, Optional: true},
						"exception_clients":   {Type: schema.TypeString, Optional: true},
						"max_interval":        {Type: schema.TypeString, Optional: true},
						"max_sessions_per_ip": {Type: schema.TypeString, Optional: true},
						"status":              {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"slow_client_attack": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_transfer_rate":           {Type: schema.TypeString, Optional: true},
						"exception_clients":            {Type: schema.TypeString, Optional: true},
						"incremental_request_timeout":  {Type: schema.TypeString, Optional: true},
						"incremental_response_timeout": {Type: schema.TypeString, Optional: true},
						"max_request_timeout":          {Type: schema.TypeString, Optional: true},
						"max_response_timeout":         {Type: schema.TypeString, Optional: true},
						"status":                       {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"website_profile": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"strict_profile_check": {Type: schema.TypeString, Optional: true},
						"allowed_domains":      {Type: schema.TypeString, Optional: true},
						"exclude_url_patterns": {Type: schema.TypeString, Optional: true},
						"include_url_patterns": {Type: schema.TypeString, Optional: true},
						"mode":                 {Type: schema.TypeString, Optional: true},
						"use_profile":          {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"advanced_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_web_application_firewall": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"accept_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"accept_list_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"proxy_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"proxy_list_status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ddos_exception_list": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_fingerprint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_http2": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"keepalive_requests": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ntlm_ignore_extra_data": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_proxy_protocol": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_vdi": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_websocket": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"basic_security": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"web_firewall_log_level": {Type: schema.TypeString, Optional: true},
						"mode":                   {Type: schema.TypeString, Optional: true},
						"trusted_hosts_action":   {Type: schema.TypeString, Optional: true},
						"trusted_hosts_group":    {Type: schema.TypeString, Optional: true},
						"ignore_case":            {Type: schema.TypeString, Optional: true},
						"client_ip_addr_header":  {Type: schema.TypeString, Optional: true},
						"rate_control_pool":      {Type: schema.TypeString, Optional: true},
						"rate_control_status":    {Type: schema.TypeString, Optional: true},
						"web_firewall_policy":    {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"load_balancing": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm":                 {Type: schema.TypeString, Optional: true},
						"persistence_cookie_domain": {Type: schema.TypeString, Optional: true},
						"cookie_age":                {Type: schema.TypeString, Optional: true},
						"persistence_cookie_name":   {Type: schema.TypeString, Optional: true},
						"persistence_cookie_path":   {Type: schema.TypeString, Optional: true},
						"failover_method":           {Type: schema.TypeString, Optional: true},
						"header_name":               {Type: schema.TypeString, Optional: true},
						"persistence_idle_timeout":  {Type: schema.TypeString, Optional: true},
						"persistence_method":        {Type: schema.TypeString, Optional: true},
						"source_ip_netmask":         {Type: schema.TypeString, Optional: true},
						"parameter_name":            {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"ssl_client_authentication": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_certificate_for_rule": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"client_authentication_rule_count": {
							Type:     schema.TypeString,
							Required: true,
						},
						"client_authentication": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enforce_client_certificate": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"trusted_certificates": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ssl_security": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate":                  {Type: schema.TypeString, Optional: true},
						"ciphers":                      {Type: schema.TypeString, Optional: true},
						"ecdsa_certificate":            {Type: schema.TypeString, Optional: true},
						"include_hsts_sub_domains":     {Type: schema.TypeString, Optional: true},
						"hsts_max_age":                 {Type: schema.TypeString, Optional: true},
						"create_hsts_redirect_service": {Type: schema.TypeString, Optional: true},
						"selected_ciphers":             {Type: schema.TypeString, Optional: true},
						"override_ciphers_ssl3":        {Type: schema.TypeString, Optional: true},
						"override_ciphers_tls_1_1":     {Type: schema.TypeString, Optional: true},
						"override_ciphers_tls_1_2":     {Type: schema.TypeString, Optional: true},
						"override_ciphers_tls_1_3":     {Type: schema.TypeString, Optional: true},
						"override_ciphers_tls_1":       {Type: schema.TypeString, Optional: true},
						"enable_pfs":                   {Type: schema.TypeString, Optional: true},
						"enable_ssl_3":                 {Type: schema.TypeString, Optional: true},
						"enable_tls_1":                 {Type: schema.TypeString, Optional: true},
						"enable_tls_1_1":               {Type: schema.TypeString, Optional: true},
						"enable_tls_1_2":               {Type: schema.TypeString, Optional: true},
						"enable_tls_1_3":               {Type: schema.TypeString, Optional: true},
						"enable_hsts":                  {Type: schema.TypeString, Optional: true},
						"enable_ocsp_stapling":         {Type: schema.TypeString, Optional: true},
						"sni_certificate":              {Type: schema.TypeString, Optional: true},
						"domain":                       {Type: schema.TypeString, Optional: true},
						"sni_ecdsa_certificate":        {Type: schema.TypeString, Optional: true},
						"enable_sni":                   {Type: schema.TypeString, Optional: true},
						"enable_strict_sni_check":      {Type: schema.TypeString, Optional: true},
						"status":                       {Type: schema.TypeString, Optional: true},
						"ssl_tls_presets":              {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"captcha_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recaptcha_type":        {Type: schema.TypeString, Optional: true},
						"recaptcha_domain":      {Type: schema.TypeString, Optional: true},
						"recaptcha_site_key":    {Type: schema.TypeString, Optional: true},
						"recaptcha_site_secret": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"ssl_ocsp": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable":        {Type: schema.TypeString, Optional: true},
						"responder_url": {Type: schema.TypeString, Optional: true},
						"certificate":   {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"url_encryption": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"referer_spam": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_patterns":      {Type: schema.TypeString, Optional: true},
						"custom_blocked_patterns": {Type: schema.TypeString, Optional: true},
						"status":                  {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"comment_spam": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exception_patterns": {Type: schema.TypeString, Optional: true},
						"parameter":          {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"advanced_client_analysis": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advanced_analysis":    {Type: schema.TypeString, Optional: true},
						"exclude_url_patterns": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"form_spam": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status":               {Type: schema.TypeString, Optional: true},
						"honeypot_status":      {Type: schema.TypeString, Optional: true},
						"autoconfigure_status": {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"waas_account": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"waas_account_id":     {Type: schema.TypeString, Required: true},
						"waas_account_serial": {Type: schema.TypeString, Optional: true},
					},
				},
			},
		},
	}
}

func resourceCudaWAFServicesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServicesResource(d, "post", resourceEndpoint),
	)

	client.hydrateBarracudaWAFServicesSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFServicesRead(d, m)
}

func resourceCudaWAFServicesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services"
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

func resourceCudaWAFServicesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFServicesResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFServicesSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFServicesRead(d, m)
}

func resourceCudaWAFServicesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services"
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

func hydrateBarracudaWAFServicesResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"address-version":     d.Get("address_version").(string),
		"dps-enabled":         d.Get("dps_enabled").(string),
		"mask":                d.Get("mask").(string),
		"session-timeout":     d.Get("session_timeout").(string),
		"linked-service-name": d.Get("linked_service_name").(string),
		"enable-access-logs":  d.Get("enable_access_logs").(string),
		"app-id":              d.Get("app_id").(string),
		"comments":            d.Get("comments").(string),
		"group":               d.Get("group").(string),
		"service-id":          d.Get("service_id").(string),
		"ip-address":          d.Get("ip_address").(string),
		"cloud-ip-select":     d.Get("cloud_ip_select").(string),
		"name":                d.Get("name").(string),
		"port":                d.Get("port").(string),
		"status":              d.Get("status").(string),
		"type":                d.Get("type").(string),
		"vsite":               d.Get("vsite").(string),
	}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"address-version", "group", "vsite"}
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

func (b *BarracudaWAF) hydrateBarracudaWAFServicesSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {
	subResourceObjects := map[string][]string{
		"authentication": {
			"send_domain_name",
			"dual_authentication",
			"secondary_authentication_service",
			"authentication_service",
			"status",
			"access_denied_page",
			"session_timeout_for_activesync",
			"cookie_domain",
			"cookie_path",
			"session_replay_protection_status",
			"creation_timeout",
			"dual_login_page",
			"idle_timeout",
			"login_challenge_page",
			"login_failed_page",
			"login_page",
			"login_processor_path",
			"openidc_redirect_url",
			"openidc_attribute_name",
			"openidc_local_id",
			"login_successful_page",
			"challenge_prompt_field",
			"challenge_user_field",
			"login_failure_url",
			"login_success_url",
			"logout_page",
			"logout_processor_path",
			"logout_successful_page",
			"logout_success_url",
			"password_expired_url",
			"post_processor_path",
			"saml_logout_url",
			"action",
			"groups",
			"sso_cookie_update_interval",
			"max_failed_attempts",
			"count_window",
			"enable_bruteforce_prevention",
			"kerberos_debug_status",
			"kerberos_enable_delegation",
			"kerberos_ldap_authorization",
			"krb_authorization_policy",
			"master_service",
			"master_service_url",
			"service_provider_display_name",
			"service_provider_entity_id",
			"service_provider_org_name",
			"service_provider_org_url",
			"attribute_format",
			"attribute_id",
			"attribute_name",
			"attribute_type",
			"encryption_certificate",
			"signing_certificate",
		},
		"caching": {
			"expiry_age",
			"file_extensions",
			"max_size",
			"min_size",
			"cache_negative_response",
			"ignore_request_headers",
			"ignore_response_headers",
			"status",
		},
		"clickjacking": {"allowed_origin", "options", "status"},
		"compression": {
			"content_types",
			"min_size",
			"status",
			"unknown_content_types",
		},
		"exception_profiling": {
			"exception_profiling_trusted_host_group",
			"exception_profiling_learn_from_trusted_host",
			"exception_profiling_level",
		},
		"ftp_security": {
			"attack_prevention_status",
			"allowed_verbs",
			"allowed_verb_status",
			"pasv_ip_address",
			"pasv_ports",
		},
		"instant_ssl": {"status", "sharepoint_rewrite_support", "secure_site_domain"},
		"ip_reputation": {
			"anonymous_proxy",
			"barracuda_reputation_blocklist",
			"custom_blacklisted_ip_status",
			"datacenter_ip",
			"fake_crawler",
			"check_registered_country",
			"block_unclassified_ips",
			"apply_policy_at",
			"geo_pool",
			"geoip_action",
			"enable_ip_reputation_filter",
			"geoip_enable_logging",
			"known_http_attack_sources",
			"public_proxy",
			"satellite_provider",
			"known_ssh_attack_sources",
			"tor_nodes",
		},
		"adaptive_profiling": {
			"content_types",
			"ignore_parameters",
			"navigation_parameters",
			"request_learning",
			"response_learning",
			"status",
			"trusted_host_group",
		},
		"sensitive_parameter_names": {"sensitive_parameter_names"},
		"session_tracking": {
			"identifiers",
			"exception_clients",
			"max_interval",
			"max_sessions_per_ip",
			"status",
		},
		"slow_client_attack": {
			"data_transfer_rate",
			"exception_clients",
			"incremental_request_timeout",
			"incremental_response_timeout",
			"max_request_timeout",
			"max_response_timeout",
			"status",
		},
		"website_profile": {
			"strict_profile_check",
			"allowed_domains",
			"exclude_url_patterns",
			"include_url_patterns",
			"mode",
			"use_profile",
		},
		"advanced_configuration": {
			"enable_web_application_firewall",
			"accept_list",
			"accept_list_status",
			"proxy_list",
			"proxy_list_status",
			"ddos_exception_list",
			"enable_fingerprint",
			"enable_http2",
			"keepalive_requests",
			"ntlm_ignore_extra_data",
			"enable_proxy_protocol",
			"enable_vdi",
			"enable_websocket",
		},
		"basic_security": {
			"web_firewall_log_level",
			"mode",
			"trusted_hosts_action",
			"trusted_hosts_group",
			"ignore_case",
			"client_ip_addr_header",
			"rate_control_pool",
			"rate_control_status",
			"web_firewall_policy",
		},
		"load_balancing": {
			"algorithm",
			"persistence_cookie_domain",
			"cookie_age",
			"persistence_cookie_name",
			"persistence_cookie_path",
			"failover_method",
			"header_name",
			"persistence_idle_timeout",
			"persistence_method",
			"source_ip_netmask",
			"parameter_name",
		},
		"ssl_client_authentication": {
			"client_certificate_for_rule",
			"client_authentication_rule_count",
			"client_authentication",
			"enforce_client_certificate",
			"trusted_certificates",
		},
		"ssl_security": {
			"certificate",
			"ciphers",
			"ecdsa_certificate",
			"include_hsts_sub_domains",
			"hsts_max_age",
			"create_hsts_redirect_service",
			"selected_ciphers",
			"override_ciphers_ssl3",
			"override_ciphers_tls_1_1",
			"override_ciphers_tls_1_2",
			"override_ciphers_tls_1_3",
			"override_ciphers_tls_1",
			"enable_pfs",
			"enable_ssl_3",
			"enable_tls_1",
			"enable_tls_1_1",
			"enable_tls_1_2",
			"enable_tls_1_3",
			"enable_hsts",
			"enable_ocsp_stapling",
			"sni_certificate",
			"domain",
			"sni_ecdsa_certificate",
			"enable_sni",
			"enable_strict_sni_check",
			"status",
			"ssl_tls_presets",
		},
		"captcha_settings": {
			"recaptcha_type",
			"recaptcha_domain",
			"recaptcha_site_key",
			"recaptcha_site_secret",
		},
		"ssl_ocsp":                 {"enable", "responder_url", "certificate"},
		"url_encryption":           {"status"},
		"referer_spam":             {"exception_patterns", "custom_blocked_patterns", "status"},
		"comment_spam":             {"exception_patterns", "parameter"},
		"advanced_client_analysis": {"advanced_analysis", "exclude_url_patterns"},
		"form_spam":                {"status", "honeypot_status", "autoconfigure_status"},
		"waas_account":             {"waas_account_id", "waas_account_serial"},
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
