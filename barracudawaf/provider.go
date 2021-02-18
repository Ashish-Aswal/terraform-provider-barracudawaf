package barracudawaf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Config : init config struct
type Config struct {
	IPAddress string
	Username  string
	Password  string
	AdminPort string
}

// WAFConfig : Provider Config struct
var WAFConfig Config

//Provider : Schema definition for barracudawaf provider
func Provider() *schema.Provider {

	// The actual provider
	provider := &schema.Provider{

		Schema: map[string]*schema.Schema{
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP Address of the WAF to be configured",
			},
			"admin_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Admin port on the WAF to be configured",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username of the WAF to be configured",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password of the WAF to be configured",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"barracudawaf_form_spam_forms":             resourceCudaWAFFormSpamForms(),
			"barracudawaf_trusted_ca_certificate":      resourceCudaWAFTrustedCaCertificate(),
			"barracudawaf_json_key_profiles":           resourceCudaWAFJsonKeyProfiles(),
			"barracudawaf_users":                       resourceCudaWAFUsers(),
			"barracudawaf_http_response_rewrite_rules": resourceCudaWAFHttpResponseRewriteRules(),
			"barracudawaf_radius_servers":              resourceCudaWAFRadiusServers(),
			"barracudawaf_local_groups":                resourceCudaWAFLocalGroups(),
			"barracudawaf_protected_data_types":        resourceCudaWAFProtectedDataTypes(),
			"barracudawaf_local_users":                 resourceCudaWAFLocalUsers(),
			"barracudawaf_content_rules":               resourceCudaWAFContentRules(),
			"barracudawaf_saml_services":               resourceCudaWAFSamlServices(),
			"barracudawaf_response_body_rewrite_rules": resourceCudaWAFResponseBodyRewriteRules(),
			"barracudawaf_network_acls":                resourceCudaWAFNetworkAcls(),
			"barracudawaf_attack_patterns":             resourceCudaWAFAttackPatterns(),
			"barracudawaf_trusted_server_certificate":  resourceCudaWAFTrustedServerCertificate(),
			"barracudawaf_secure_browsing_policies":    resourceCudaWAFSecureBrowsingPolicies(),
			"barracudawaf_global_threshold":            resourceCudaWAFGlobalThreshold(),
			"barracudawaf_external_radius_services":    resourceCudaWAFExternalRadiusServices(),
			"barracudawaf_url_profiles":                resourceCudaWAFUrlProfiles(),
			"barracudawaf_syslog_servers":              resourceCudaWAFSyslogServers(),
			"barracudawaf_services":                    resourceCudaWAFServices(),
			"barracudawaf_url_acls":                    resourceCudaWAFUrlAcls(),
			"barracudawaf_http_request_rewrite_rules":  resourceCudaWAFHttpRequestRewriteRules(),
			"barracudawaf_bonds":                       resourceCudaWAFBonds(),
			"barracudawaf_input_patterns":              resourceCudaWAFInputPatterns(),
			"barracudawaf_radius_services":             resourceCudaWAFRadiusServices(),
			"barracudawaf_web_scraping_policies":       resourceCudaWAFWebScrapingPolicies(),
			"barracudawaf_action_policies":             resourceCudaWAFActionPolicies(),
			"barracudawaf_external_ldap_services":      resourceCudaWAFExternalLdapServices(),
			"barracudawaf_export_configuration":        resourceCudaWAFExportConfiguration(),
			"barracudawaf_kerberos_services":           resourceCudaWAFKerberosServices(),
			"barracudawaf_notification_configuration":  resourceCudaWAFNotificationConfiguration(),
			"barracudawaf_header_acls":                 resourceCudaWAFHeaderAcls(),
			"barracudawaf_adaptive_profiling_rules":    resourceCudaWAFAdaptiveProfilingRules(),
			"barracudawaf_authorization_policies":      resourceCudaWAFAuthorizationPolicies(),
			"barracudawaf_virtual_interfaces":          resourceCudaWAFVirtualInterfaces(),
			"barracudawaf_global_acls":                 resourceCudaWAFGlobalAcls(),
			"barracudawaf_parameter_profiles":          resourceCudaWAFParameterProfiles(),
			"barracudawaf_configuration_checkpoints":   resourceCudaWAFConfigurationCheckpoints(),
			"barracudawaf_cluster":                     resourceCudaWAFCluster(),
			"barracudawaf_import_configuration":        resourceCudaWAFImportConfiguration(),
			"barracudawaf_identity_types":              resourceCudaWAFIdentityTypes(),
			"barracudawaf_ldap_services":               resourceCudaWAFLdapServices(),
			"barracudawaf_nodes":                       resourceCudaWAFNodes(),
			"barracudawaf_geoip_allowed_networks":      resourceCudaWAFGeoipAllowedNetworks(),
			"barracudawaf_service_groups":              resourceCudaWAFServiceGroups(),
			"barracudawaf_rsa_securid_services":        resourceCudaWAFRsaSecuridServices(),
			"barracudawaf_vlans":                       resourceCudaWAFVlans(),
			"barracudawaf_bot_spam_patterns":           resourceCudaWAFBotSpamPatterns(),
			"barracudawaf_interface_routes":            resourceCudaWAFInterfaceRoutes(),
			"barracudawaf_openidc_identity_providers":  resourceCudaWAFOpenidcIdentityProviders(),
			"barracudawaf_bot_spam_types":              resourceCudaWAFBotSpamTypes(),
			"barracudawaf_administrator_roles":         resourceCudaWAFAdministratorRoles(),
			"barracudawaf_content_rule_servers":        resourceCudaWAFContentRuleServers(),
			"barracudawaf_parameter_optimizers":        resourceCudaWAFParameterOptimizers(),
			"barracudawaf_url_optimizers":              resourceCudaWAFUrlOptimizers(),
			"barracudawaf_rate_control_pools":          resourceCudaWAFRateControlPools(),
			"barracudawaf_whitelisted_bots":            resourceCudaWAFWhitelistedBots(),
			"barracudawaf_ddos_policies":               resourceCudaWAFDdosPolicies(),
			"barracudawaf_geoip_blocked_networks":      resourceCudaWAFGeoipBlockedNetworks(),
			"barracudawaf_identity_theft_patterns":     resourceCudaWAFIdentityTheftPatterns(),
			"barracudawaf_destination_nats":            resourceCudaWAFDestinationNats(),
			"barracudawaf_preferred_clients":           resourceCudaWAFPreferredClients(),
			"barracudawaf_network_interfaces":          resourceCudaWAFNetworkInterfaces(),
			"barracudawaf_trusted_host_groups":         resourceCudaWAFTrustedHostGroups(),
			"barracudawaf_backup":                      resourceCudaWAFBackup(),
			"barracudawaf_rsa_securid_servers":         resourceCudaWAFRsaSecuridServers(),
			"barracudawaf_reports":                     resourceCudaWAFReports(),
			"barracudawaf_input_types":                 resourceCudaWAFInputTypes(),
			"barracudawaf_security_policies":           resourceCudaWAFSecurityPolicies(),
			"barracudawaf_auto_system_acls":            resourceCudaWAFAutoSystemAcls(),
			"barracudawaf_geo_pools":                   resourceCudaWAFGeoPools(),
			"barracudawaf_source_nats":                 resourceCudaWAFSourceNats(),
			"barracudawaf_client_certificate_crls":     resourceCudaWAFClientCertificateCrls(),
			"barracudawaf_signed_certificate":          resourceCudaWAFSignedCertificate(),
			"barracudawaf_static_routes":               resourceCudaWAFStaticRoutes(),
			"barracudawaf_trusted_hosts":               resourceCudaWAFTrustedHosts(),
			"barracudawaf_openidc_services":            resourceCudaWAFOpenidcServices(),
			"barracudawaf_kerberos_servers":            resourceCudaWAFKerberosServers(),
			"barracudawaf_session_identifiers":         resourceCudaWAFSessionIdentifiers(),
			"barracudawaf_credential_servers":          resourceCudaWAFCredentialServers(),
			"barracudawaf_json_profiles":               resourceCudaWAFJsonProfiles(),
			"barracudawaf_access_rules":                resourceCudaWAFAccessRules(),
			"barracudawaf_url_encryption_rules":        resourceCudaWAFUrlEncryptionRules(),
			"barracudawaf_attack_types":                resourceCudaWAFAttackTypes(),
			"barracudawaf_url_policies":                resourceCudaWAFUrlPolicies(),
			"barracudawaf_self_signed_certificate":     resourceCudaWAFSelfSignedCertificate(),
			"barracudawaf_json_security_policies":      resourceCudaWAFJsonSecurityPolicies(),
			"barracudawaf_custom_parameter_classes":    resourceCudaWAFCustomParameterClasses(),
			"barracudawaf_import_openapi":              resourceCudaWAFImportOpenapi(),
			"barracudawaf_servers":                     resourceCudaWAFServers(),
			"barracudawaf_response_pages":              resourceCudaWAFResponsePages(),
			"barracudawaf_saml_identity_providers":     resourceCudaWAFSamlIdentityProviders(),
			"barracudawaf_custom_ip_blocklist":         resourceCudaWAFCustomIpBlocklist(),
			"barracudawaf_url_translations":            resourceCudaWAFUrlTranslations(),
			"barracudawaf_allow_deny_clients":          resourceCudaWAFAllowDenyClients(),
			"barracudawaf_vsites":                      resourceCudaWAFVsites(),
			"barracudawaf_ldap_servers":                resourceCudaWAFLdapServers(),
		},

		ConfigureFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	WAFConfig = Config{
		IPAddress: d.Get("ip").(string),
		AdminPort: d.Get("admin_port").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
	}

	return &WAFConfig, nil
}
