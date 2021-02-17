package barracudawaf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFServiceGroups() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFServiceGroupsCreate,
		Read:   resourceCudaWAFServiceGroupsRead,
		Update: resourceCudaWAFServiceGroupsUpdate,
		Delete: resourceCudaWAFServiceGroupsDelete,

		Schema: map[string]*schema.Schema{
			"service_group": {Type: schema.TypeString, Required: true},
			"parent": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func makeRestAPIPayloadServiceGroups(
	d *schema.ResourceData,
	resourceOperation string,
	resourceEndpoint string,
) error {

	//resourcePayload : Payload for the resource
	resourcePayload := map[string]string{"service-group": d.Get("service_group").(string)}

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

func resourceCudaWAFServiceGroupsCreate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vsites/" + d.Get("parent.0").(string) + "/service-groups"
	resourceCreateResponseError := makeRestAPIPayloadServiceGroups(d, "POST", resourceEndpoint)

	if resourceCreateResponseError != nil {
		return fmt.Errorf("%v", resourceCreateResponseError)
	}

	return resourceCudaWAFServiceGroupsRead(d, m)
}

func resourceCudaWAFServiceGroupsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFServiceGroupsUpdate(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vsites/" + d.Get("parent.0").(string) + "/service-groups/" + d.Get("name").(string)
	resourceUpdateResponseError := makeRestAPIPayloadServiceGroups(d, "PUT", resourceEndpoint)

	if resourceUpdateResponseError != nil {
		return fmt.Errorf("%v", resourceUpdateResponseError)
	}

	return resourceCudaWAFServiceGroupsRead(d, m)
}

func resourceCudaWAFServiceGroupsDelete(d *schema.ResourceData, m interface{}) error {
	resourceEndpoint := baseURI + "/vsites/" + d.Get("parent.0").(string) + "/service-groups/" + d.Get("name").(string)
	resourceDeleteResponseError := makeRestAPIPayloadServiceGroups(d, "DELETE", resourceEndpoint)

	if resourceDeleteResponseError != nil {
		return fmt.Errorf("%v", resourceDeleteResponseError)
	}

	return nil
}
