package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	subResourceSecurityPoliciesParams = map[string][]string{
		"request_limits": {
			"max_cookie_name_length",
			"max_number_of_cookies",
			"max_header_name_length",
			"max_request_line_length",
			"max_number_of_headers",
			"max_query_length",
			"max_cookie_value_length",
			"max_header_value_length",
			"max_request_length",
			"max_url_length",
			"enable",
		},
		"url_normalization": {
			"default_charset",
			"detect_response_charset",
			"normalize_special_chars",
			"apply_double_decoding",
			"parameter_separators",
		},
		"url_protection": {
			"allowed_methods",
			"allowed_content_types",
			"custom_blocked_attack_types",
			"exception_patterns",
			"blocked_attack_types",
			"max_content_length",
			"maximum_parameter_name_length",
			"max_parameters",
			"maximum_upload_files",
			"csrf_prevention",
			"enable",
		},
		"parameter_protection": {
			"blocked_attack_types",
			"custom_blocked_attack_types",
			"base64_decode_parameter_value",
			"allowed_file_upload_type",
			"denied_metacharacters",
			"exception_patterns",
			"file_upload_extensions",
			"file_upload_mime_types",
			"maximum_instances",
			"maximum_parameter_value_length",
			"maximum_upload_file_size",
			"enable",
			"validate_parameter_name",
			"ignore_parameters",
		},
		"cloaking": {
			"return_codes_to_exempt",
			"headers_to_filter",
			"filter_response_header",
			"suppress_return_code",
		},
		"cookie_security": {
			"allow_unrecognized_cookies",
			"days_allowed",
			"cookies_exempted",
			"http_only",
			"cookie_max_age",
			"tamper_proof_mode",
			"secure_cookie",
			"cookie_replay_protection_type",
			"custom_headers",
		},
		"client_profile": {"medium_risk_score", "high_risk_score", "exception_client_fingerprints", "client_profile"},
		"tarpit_profile": {"backlog_requests_limit", "tarpit_inactivity_timeout"},
	}
)

func resourceCudaWAFSecurityPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSecurityPoliciesCreate,
		Read:   resourceCudaWAFSecurityPoliciesRead,
		Update: resourceCudaWAFSecurityPoliciesUpdate,
		Delete: resourceCudaWAFSecurityPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"based_on": {Type: schema.TypeString, Optional: true},
			"name":     {Type: schema.TypeString, Required: true},
			"request_limits": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_cookie_name_length":  {Type: schema.TypeString, Optional: true},
						"max_number_of_cookies":   {Type: schema.TypeString, Optional: true},
						"max_header_name_length":  {Type: schema.TypeString, Optional: true},
						"max_request_line_length": {Type: schema.TypeString, Optional: true},
						"max_number_of_headers":   {Type: schema.TypeString, Optional: true},
						"max_query_length":        {Type: schema.TypeString, Optional: true},
						"max_cookie_value_length": {Type: schema.TypeString, Optional: true},
						"max_header_value_length": {Type: schema.TypeString, Optional: true},
						"max_request_length":      {Type: schema.TypeString, Optional: true},
						"max_url_length":          {Type: schema.TypeString, Optional: true},
						"enable":                  {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"url_normalization": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_charset":         {Type: schema.TypeString, Optional: true},
						"detect_response_charset": {Type: schema.TypeString, Optional: true},
						"normalize_special_chars": {Type: schema.TypeString, Optional: true},
						"apply_double_decoding":   {Type: schema.TypeString, Optional: true},
						"parameter_separators":    {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"url_protection": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_methods":               {Type: schema.TypeString, Required: true},
						"allowed_content_types":         {Type: schema.TypeString, Optional: true},
						"custom_blocked_attack_types":   {Type: schema.TypeString, Optional: true},
						"exception_patterns":            {Type: schema.TypeString, Optional: true},
						"blocked_attack_types":          {Type: schema.TypeString, Optional: true},
						"max_content_length":            {Type: schema.TypeString, Optional: true},
						"maximum_parameter_name_length": {Type: schema.TypeString, Optional: true},
						"max_parameters":                {Type: schema.TypeString, Optional: true},
						"maximum_upload_files":          {Type: schema.TypeString, Optional: true},
						"csrf_prevention":               {Type: schema.TypeString, Optional: true},
						"enable":                        {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"parameter_protection": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"blocked_attack_types":           {Type: schema.TypeString, Optional: true},
						"custom_blocked_attack_types":    {Type: schema.TypeString, Optional: true},
						"base64_decode_parameter_value":  {Type: schema.TypeString, Optional: true},
						"allowed_file_upload_type":       {Type: schema.TypeString, Optional: true},
						"denied_metacharacters":          {Type: schema.TypeString, Optional: true},
						"exception_patterns":             {Type: schema.TypeString, Optional: true},
						"file_upload_extensions":         {Type: schema.TypeString, Optional: true},
						"file_upload_mime_types":         {Type: schema.TypeString, Optional: true},
						"maximum_instances":              {Type: schema.TypeString, Optional: true},
						"maximum_parameter_value_length": {Type: schema.TypeString, Optional: true},
						"maximum_upload_file_size":       {Type: schema.TypeString, Optional: true},
						"enable":                         {Type: schema.TypeString, Optional: true},
						"validate_parameter_name":        {Type: schema.TypeString, Optional: true},
						"ignore_parameters":              {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"cloaking": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"return_codes_to_exempt": {Type: schema.TypeString, Optional: true},
						"headers_to_filter":      {Type: schema.TypeString, Required: true},
						"filter_response_header": {Type: schema.TypeString, Optional: true},
						"suppress_return_code":   {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"cookie_security": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_unrecognized_cookies":    {Type: schema.TypeString, Optional: true},
						"days_allowed":                  {Type: schema.TypeString, Optional: true},
						"cookies_exempted":              {Type: schema.TypeString, Optional: true},
						"http_only":                     {Type: schema.TypeString, Optional: true},
						"cookie_max_age":                {Type: schema.TypeString, Optional: true},
						"tamper_proof_mode":             {Type: schema.TypeString, Optional: true},
						"secure_cookie":                 {Type: schema.TypeString, Optional: true},
						"cookie_replay_protection_type": {Type: schema.TypeString, Optional: true},
						"custom_headers":                {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"client_profile": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"medium_risk_score":             {Type: schema.TypeString, Optional: true},
						"high_risk_score":               {Type: schema.TypeString, Optional: true},
						"exception_client_fingerprints": {Type: schema.TypeString, Optional: true},
						"client_profile":                {Type: schema.TypeString, Optional: true},
					},
				},
			},
			"tarpit_profile": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backlog_requests_limit":    {Type: schema.TypeString, Optional: true},
						"tarpit_inactivity_timeout": {Type: schema.TypeString, Optional: true},
					},
				},
			},
		},
	}
}

func resourceCudaWAFSecurityPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies"
	client.CreateBarracudaWAFResource(name, hydrateBarracudaWAFSecurityPoliciesResource(d, "post", resourceEndpoint))

	client.hydrateBarracudaWAFSecurityPoliciesSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFSecurityPoliciesRead(d, m)
}

func resourceCudaWAFSecurityPoliciesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies"
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

func resourceCudaWAFSecurityPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies"
	err := client.UpdateBarracudaWAFResource(name, hydrateBarracudaWAFSecurityPoliciesResource(d, "put", resourceEndpoint))

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFSecurityPoliciesSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSecurityPoliciesRead(d, m)
}

func resourceCudaWAFSecurityPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/security-policies"
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

func hydrateBarracudaWAFSecurityPoliciesResource(d *schema.ResourceData, method string, endpoint string) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{"based-on": d.Get("based_on").(string), "name": d.Get("name").(string)}

	// parameters not supported for updates
	if method == "put" {
		updatePayloadExceptions := [...]string{"based-on"}
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

func (b *BarracudaWAF) hydrateBarracudaWAFSecurityPoliciesSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {

	for subResource, subResourceParams := range subResourceSecurityPoliciesParams {
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
