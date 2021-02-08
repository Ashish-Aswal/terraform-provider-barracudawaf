package barracudawaf

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCudaWAFCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCudaWAFCertificateCreate,
		Read:   resourceCudaWAFCertificateRead,
		Update: resourceCudaWAFCertificateUpdate,
		Delete: resourceCudaWAFCertificateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allow_private_key_export": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"common_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"country_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"curve_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_size": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trusted_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"signed_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assign_associated_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"intermediary_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trusted_server_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func makeRestAPIPayloadCertificate(d *schema.ResourceData, oper string, endpoint string) error {
	payload := map[string]string{
		"name":                       d.Get("name").(string),
		"key":                        d.Get("key").(string),
		"city":                       d.Get("city").(string),
		"state":                      d.Get("state").(string),
		"type":                       d.Get("type").(string),
		"upload":                     d.Get("upload").(string),
		"password":                   d.Get("password").(string),
		"key_size":                   d.Get("key_size").(string),
		"key_type":                   d.Get("key_type").(string),
		"curve_type":                 d.Get("curve_type").(string),
		"common_name":                d.Get("common_name").(string),
		"country_code":               d.Get("country_code").(string),
		"organization_name":          d.Get("organization_name").(string),
		"organization_unit":          d.Get("organization_unit").(string),
		"signed_certificate":         d.Get("signed_certificate").(string),
		"trusted_certificate":        d.Get("trusted_certificate").(string),
		"assign_associated_key":      d.Get("assign_associated_key").(string),
		"allow_private_key_export":   d.Get("allow_private_key_export").(string),
		"intermediary_certificate":   d.Get("intermediary_certificate").(string),
		"trusted_server_certificate": d.Get("trusted_server_certificate").(string),
	}

	for key, value := range payload {
		if len(value) > 0 {
			continue
		} else {
			delete(payload, key)
		}
	}

	if oper == "DELETE" {
		callData := map[string]interface{}{
			"endpoint":  endpoint,
			"payload":   payload,
			"operation": oper,
			"name":      d.Get("name").(string),
		}
		callStatus, callRespBody := updateCudaWAFResourceObject(callData)
		if callStatus != 200 && callStatus != 201 {
			return fmt.Errorf("some error occurred : %v", callRespBody["msg"])
		}
		return nil
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	if upload, exists := d.GetOk("upload"); exists {
		upload = upload.(string)
		if upload == "signed" {
			values := map[string]io.Reader{
				"name":                     strings.NewReader(payload["name"]),
				"type":                     strings.NewReader(payload["type"]),
				"key_type":                 strings.NewReader(payload["key_type"]),
				"password":                 strings.NewReader(payload["password"]),
				"assign_associated_key":    strings.NewReader(payload["assign_associated_key"]),
				"allow_private_key_export": strings.NewReader(payload["allow_private_key_export"]),
			}

			if key, exists := d.GetOk("key"); exists {
				key1 := key.(string)
				values["key"] = mustOpen(key1)
			}

			if cert, exists := d.GetOk("signed_certificate"); exists {
				cert1 := cert.(string)
				values["signed_certificate"] = mustOpen(cert1)
			}

			if interCert, exists := d.GetOk("intermediary_certificate"); exists {
				interCert1 := interCert.(string)
				values["intermediary_certificate"] = mustOpen(interCert1)
			}

			status, body := uploadCertificateContent(client, endpoint, values)
			if status == 200 || status == 201 {
				if oper != "DELETE" {
					d.SetId(body["id"].(string))
				}
			} else {
				return fmt.Errorf("some error occurred : %v", body["msg"])
			}
		} else if upload == "trusted" {
			interCert1 := d.Get("trusted_certificate").(string)
			values := map[string]io.Reader{
				"name":                strings.NewReader(payload["name"]),
				"trusted_certificate": mustOpen(interCert1),
			}
			status, body := uploadCertificateContent(client, endpoint, values)
			if status == 200 || status == 201 {
				if oper != "DELETE" {
					d.SetId(body["id"].(string))
				}
			} else {
				return fmt.Errorf("some error occurred : %v", body["msg"])
			}
		} else if upload == "trusted_server" {
			interCert1 := d.Get("trusted_server_certificate").(string)
			values := map[string]io.Reader{
				"name":                       strings.NewReader(payload["name"]),
				"trusted_server_certificate": mustOpen(interCert1),
			}
			status, body := uploadCertificateContent(client, endpoint, values)
			if status == 200 || status == 201 {
				if oper != "DELETE" {
					d.SetId(body["id"].(string))
				}
			} else {
				return fmt.Errorf("some error occurred : %v", body["msg"])
			}
		}
	} else {
		// Create Self-Signed Certificate for Demo
		endpoint1 := "restapi/v3/certificates"
		callData := map[string]interface{}{
			"endpoint":  endpoint1,
			"payload":   payload,
			"operation": oper,
			"name":      d.Get("name").(string),
		}
		callStatus, callRespBody := updateCudaWAFResourceObject(callData)
		if callStatus == 200 || callStatus == 201 {
			if oper != "DELETE" {
				d.SetId(callRespBody["id"].(string))
			}
		} else {
			return fmt.Errorf("some error occurred : %v", callRespBody["msg"])
		}
	}

	return nil
}

func resourceCudaWAFCertificateCreate(d *schema.ResourceData, m interface{}) error {
	remoteURL := "http://" + WAFConfig.IPAddress + ":" + WAFConfig.AdminPort + "/restapi/v3/certificates?upload="
	upload, exists := d.GetOk("upload")
	if exists {
		if upload == "signed" {
			remoteURL = remoteURL + "signed"
		} else if upload == "trusted" {
			remoteURL = remoteURL + "trusted"
		} else if upload == "trusted_server" {
			remoteURL = remoteURL + "trusted_server"
		}
	}
	err := makeRestAPIPayloadCertificate(d, "POST", remoteURL)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourceCudaWAFServerRead(d, m)
}

func resourceCudaWAFCertificateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCudaWAFCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCudaWAFServerRead(d, m)
}

func resourceCudaWAFCertificateDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	endpoint := "restapi/v3/certificates/" + name
	err := makeRestAPIPayloadCertificate(d, "DELETE", endpoint)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
