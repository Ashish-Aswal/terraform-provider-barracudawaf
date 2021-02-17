package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFBonds() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFBondsCreate,
		Read:   resourceCudaWAFBondsRead,
		Update: resourceCudaWAFBondsUpdate,
		Delete: resourceCudaWAFBondsDelete,

		Schema: map[string]*schema.Schema{
			"duplexity":  {Type: schema.TypeString, Optional: true},
			"name":       {Type: schema.TypeString, Required: true},
			"speed":      {Type: schema.TypeString, Optional: true},
			"bond_ports": {Type: schema.TypeString, Required: true},
			"min_link":   {Type: schema.TypeString, Optional: true},
			"mode":       {Type: schema.TypeString, Optional: true},
			"mtu":        {Type: schema.TypeString, Optional: true},
		},
	}
}

func makeRestAPIPayloadBonds(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{
		"duplexity":  d.Get("duplexity").(string),
		"name":       d.Get("name").(string),
		"speed":      d.Get("speed").(string),
		"bond-ports": d.Get("bond_ports").(string),
		"min-link":   d.Get("min_link").(string),
		"mode":       d.Get("mode").(string),
		"mtu":        d.Get("mtu").(string),
	}

	//check resourcePayload for updates(modify) on the resource
	if resourceOperation == "PUT" {
		updatePayloadExceptions := [...]string{"name"}
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

func resourceCudaWAFBondsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/bonds"
	resourceCreateResponseError := makeRestAPIPayloadBonds(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFBondsRead(d, m)
}

func resourceCudaWAFBondsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFBondsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/bonds/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadBonds(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFBondsRead(d, m)
}

func resourceCudaWAFBondsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/bonds/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadBonds(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
