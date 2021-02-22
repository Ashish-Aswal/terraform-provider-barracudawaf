package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSyslogServersResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFSyslogServersRead(d, m)
}

func resourceCudaWAFSyslogServersRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFSyslogServersUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	resourceEndpoint := "/syslog-servers/"
	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFSyslogServersResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFSyslogServersRead(d, m)
}

func resourceCudaWAFSyslogServersDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/syslog-servers/"
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

func hydrateBarracudaWAFSyslogServersResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

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
