package barracudawaf

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFClusterCreate,
		Read:   resourceCudaWAFClusterRead,
		Update: resourceCudaWAFClusterUpdate,
		Delete: resourceCudaWAFClusterDelete,

		Schema: map[string]*schema.Schema{
			"heartbeat_count_per_interface": {Type: schema.TypeString, Optional: true},
			"heartbeat_frequency":           {Type: schema.TypeString, Optional: true},
			"monitor_link":                  {Type: schema.TypeString, Optional: true},
			"transmit_heartbeat_on":         {Type: schema.TypeString, Optional: true},
			"cluster_shared_secret":         {Type: schema.TypeString, Optional: true},
			"failback_mode":                 {Type: schema.TypeString, Optional: true},
			"cluster_name":                  {Type: schema.TypeString, Required: true},
			"data_path_failure_action":      {Type: schema.TypeString, Optional: true},
			"vx_aa_enable":                  {Type: schema.TypeString, Optional: true},
		},
	}
}

func resourceCudaWAFClusterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Get("name").(string)

	log.Println("[INFO] Creating Barracuda WAF resource " + name)

	resourceEndpoint := "/cluster"
	client.CreateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFClusterResource(d, "post", resourceEndpoint),
	)

	d.SetId(name)
	return resourceCudaWAFClusterRead(d, m)
}

func resourceCudaWAFClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()
	log.Println("[INFO] Fetching Barracuda WAF resource " + name)

	resourceEndpoint := "/cluster"
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

func resourceCudaWAFClusterUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Updating Barracuda WAF resource " + name)

	resourceEndpoint := "/cluster/"
	err := client.UpdateBarracudaWAFResource(
		name,
		hydrateBarracudaWAFClusterResource(d, "put", resourceEndpoint),
	)

	if err != nil {
		log.Printf("[ERROR] Unable to update the Barracuda WAF resource (%s) (%v)", name, err)
		return err
	}

	return resourceCudaWAFClusterRead(d, m)
}

func resourceCudaWAFClusterDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*BarracudaWAF)

	name := d.Id()

	log.Println("[INFO] Deleting Barracuda WAF resource " + name)

	resourceEndpoint := "/cluster/"
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

func hydrateBarracudaWAFClusterResource(
	d *schema.ResourceData,
	method string,
	endpoint string,
) *APIRequest {

	//resourcePayload : payload for the resource
	resourcePayload := map[string]string{
		"heartbeat-count-per-interface": d.Get("heartbeat_count_per_interface").(string),
		"heartbeat-frequency":           d.Get("heartbeat_frequency").(string),
		"monitor-link":                  d.Get("monitor_link").(string),
		"transmit-heartbeat-on":         d.Get("transmit_heartbeat_on").(string),
		"cluster-shared-secret":         d.Get("cluster_shared_secret").(string),
		"failback-mode":                 d.Get("failback_mode").(string),
		"cluster-name":                  d.Get("cluster_name").(string),
		"data-path-failure-action":      d.Get("data_path_failure_action").(string),
		"vx-aa-enable":                  d.Get("vx_aa_enable").(string),
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
