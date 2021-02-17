package barracudawaf

import (
	"fmt"

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

func makeRestAPIPayloadCluster(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
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

func resourceCudaWAFClusterCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/cluster"
	resourceCreateResponseError := makeRestAPIPayloadCluster(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFClusterRead(d, m)
}

func resourceCudaWAFClusterRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFClusterUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/cluster/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadCluster(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFClusterRead(d, m)
}

func resourceCudaWAFClusterDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/cluster/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadCluster(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
