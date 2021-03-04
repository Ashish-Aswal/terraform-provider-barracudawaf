package barracudawaf

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	subResourceSyslogServersParams = map[string][]string{}
)

func resourceCudaWAFSyslogServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFSyslogServersCreate,
		Read:   resourceCudaWAFSyslogServersRead,
		Update: resourceCudaWAFSyslogServersUpdate,
		Delete: resourceCudaWAFSyslogServersDelete,

		Schema: map[string]*schema.Schema{
			"brs_host":                    {Type: schema.TypeString, Optional: true},
			"brs_shared_secret":           {Type: schema.TypeString, Optional: true},
			"brs_system_serial":           {Type: schema.TypeString, Optional: true},
			"cloud_syslog_token":          {Type: schema.TypeString, Optional: true},
			"timestamp_and_hostname":      {Type: schema.TypeString, Optional: true},
			"password":                    {Type: schema.TypeString, Optional: true},
			"username":                    {Type: schema.TypeString, Optional: true},
			"client_certificate":          {Type: schema.TypeString, Optional: true},
			"connection_type":             {Type: schema.TypeString, Optional: true},
			"event_hub_name":              {Type: schema.TypeString, Optional: true},
			"event_queue_name":            {Type: schema.TypeString, Optional: true},
			"oms_custom_log":              {Type: schema.TypeString, Optional: true},
			"oms_govcloud":                {Type: schema.TypeString, Optional: true},
			"oms_key":                     {Type: schema.TypeString, Optional: true},
			"oms_workspace":               {Type: schema.TypeString, Optional: true},
			"policy_sas_key":              {Type: schema.TypeString, Optional: true},
			"policy_name":                 {Type: schema.TypeString, Optional: true},
			"log_group":                   {Type: schema.TypeString, Optional: true},
			"comments":                    {Type: schema.TypeString, Optional: true},
			"ip_address":                  {Type: schema.TypeString, Optional: true},
			"name":                        {Type: schema.TypeString, Required: true},
			"port":                        {Type: schema.TypeString, Optional: true},
			"server_type":                 {Type: schema.TypeString, Required: true},
			"service_bus_name":            {Type: schema.TypeString, Optional: true},
			"validate_server_certificate": {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFSyslogServersCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/syslog-servers"
	client.CreateBarracudaWAFResource(name, hydrateBarracudaWAFSyslogServersResource(d, "post", resourceEndpoint))

	client.hydrateBarracudaWAFSyslogServersSubResource(d, name, resourceEndpoint)

	d.SetId(name)
	return resourceCudaWAFSyslogServersRead(d, m)
}

func resourceCudaWAFSyslogServersRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/syslog-servers"
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

func resourceCudaWAFSyslogServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/syslog-servers"
	err := client.UpdateBarracudaWAFResource(name, hydrateBarracudaWAFSyslogServersResource(d, "put", resourceEndpoint))

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	err = client.hydrateBarracudaWAFSyslogServersSubResource(d, name, resourceEndpoint)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF sub resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSyslogServersRead(d, m)
}

func resourceCudaWAFSyslogServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/syslog-servers"
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

func hydrateBarracudaWAFSyslogServersResource(d *schema.ResourceData, method string, endpoint string) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"brs-host":                    d.Get("brs_host").(string),
		"brs-shared-secret":           d.Get("brs_shared_secret").(string),
		"brs-system-serial":           d.Get("brs_system_serial").(string),
		"cloud-syslog-token":          d.Get("cloud_syslog_token").(string),
		"timestamp-and-hostname":      d.Get("timestamp_and_hostname").(string),
		"password":                    d.Get("password").(string),
		"username":                    d.Get("username").(string),
		"client-certificate":          d.Get("client_certificate").(string),
		"connection-type":             d.Get("connection_type").(string),
		"event-hub-name":              d.Get("event_hub_name").(string),
		"event-queue-name":            d.Get("event_queue_name").(string),
		"oms-custom-log":              d.Get("oms_custom_log").(string),
		"oms-govcloud":                d.Get("oms_govcloud").(string),
		"oms-key":                     d.Get("oms_key").(string),
		"oms-workspace":               d.Get("oms_workspace").(string),
		"policy-sas-key":              d.Get("policy_sas_key").(string),
		"policy-name":                 d.Get("policy_name").(string),
		"log-group":                   d.Get("log_group").(string),
		"comments":                    d.Get("comments").(string),
		"ip-address":                  d.Get("ip_address").(string),
		"name":                        d.Get("name").(string),
		"port":                        d.Get("port").(string),
		"server_type":                 d.Get("server_type").(string),
		"service-bus-name":            d.Get("service_bus_name").(string),
		"validate-server-certificate": d.Get("validate_server_certificate").(string),
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

func (b *BarracudaWAF) hydrateBarracudaWAFSyslogServersSubResource(
	d *schema.ResourceData,
	name string,
	endpoint string,
) error {

	for subResource, subResourceParams := range subResourceSyslogServersParams {
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
