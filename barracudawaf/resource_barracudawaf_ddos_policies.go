package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	subResourceDdosPoliciesParams = map[string][]string{}
)

func resourceCudaWAFDdosPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFDdosPoliciesCreate,
		Read:   resourceCudaWAFDdosPoliciesRead,
		Update: resourceCudaWAFDdosPoliciesUpdate,
		Delete: resourceCudaWAFDdosPoliciesDelete,

		Schema: map[string]*schema.Schema{
			"evaluate_clients":        {Type: schema.TypeString, Optional: true, Description: "Enable URL ACL"},
			"comments":                {Type: schema.TypeString, Optional: true, Description: "Comments"},
			"enforce_captcha":         {Type: schema.TypeString, Optional: true, Description: "Enforce CAPTCHA"},
			"extended_match":          {Type: schema.TypeString, Optional: true, Description: "Extended Match"},
			"extended_match_sequence": {Type: schema.TypeString, Optional: true, Description: "Extended Match Sequence"},
			"host":                    {Type: schema.TypeString, Optional: true, Description: "Host Match"},
			"expiry_time":             {Type: schema.TypeString, Optional: true, Description: "Expiry time"},
			"mouse_check":             {Type: schema.TypeString, Optional: true, Description: "Detect Mouse Event"},
			"name":                    {Type: schema.TypeString, Required: true, Description: "DDos Policy Name"},
			"num_captcha_tries":       {Type: schema.TypeString, Optional: true, Description: "Max CAPTCHA Attempts"},
			"num_unanswered_captcha":  {Type: schema.TypeString, Optional: true, Description: "Max Unanswered CAPTCHA"},
			"url":                     {Type: schema.TypeString, Optional: true, Description: "URL Match"},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},

		Description: "`barracudawaf_ddos_policies` manages `Ddos Policies` on the Barracuda Web Application Firewall.",
	}
}

func resourceCudaWAFDdosPoliciesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies"
	err := client.CreateBarracudaWAFResource(name, hydrateBarracudaWAFDdosPoliciesResource(d, "post", resourceEndpoint))

	if err != nil {
		log.Printf("[ERROR] Unable to create Barracuda WAF resource (%s) (%v) ", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFDdosPoliciesSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to create Barracuda WAF sub resource (%s) (%v) ", name, err)
		return err
	}

	d.SetId(name)
	return resourceCudaWAFDdosPoliciesRead(d, m)
}

func resourceCudaWAFDdosPoliciesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies"
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

func resourceCudaWAFDdosPoliciesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies"
	err := client.UpdateBarracudaWAFResource(name, hydrateBarracudaWAFDdosPoliciesResource(d, "put", resourceEndpoint))

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFDdosPoliciesSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFDdosPoliciesRead(d, m)
}

func resourceCudaWAFDdosPoliciesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/services/" + d.Get("parent.0").(string) + "/ddos-policies"
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

func hydrateBarracudaWAFDdosPoliciesResource(d *schema.ResourceData, method string, endpoint string) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"evaluate-clients":        d.Get("evaluate_clients").(string),
		"comments":                d.Get("comments").(string),
		"enforce-captcha":         d.Get("enforce_captcha").(string),
		"extended-match":          d.Get("extended_match").(string),
		"extended-match-sequence": d.Get("extended_match_sequence").(string),
		"host":                    d.Get("host").(string),
		"expiry-time":             d.Get("expiry_time").(string),
		"mouse-check":             d.Get("mouse_check").(string),
		"name":                    d.Get("name").(string),
		"num-captcha-tries":       d.Get("num_captcha_tries").(string),
		"num-unanswered-captcha":  d.Get("num_unanswered_captcha").(string),
		"url":                     d.Get("url").(string),
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

func (b *BarracudaWAF) hydrateBarracudaWAFDdosPoliciesSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {

	for subResource, subResourceParams := range subResourceDdosPoliciesParams {
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
