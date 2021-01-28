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
			"barracudawaf_service":           resourceCudaWAFService(),
			"barracudawaf_server":            resourceCudaWAFServer(),
			"barracudawaf_rule_group":        resourceCudaWAFRuleGroup(),
			"barracudawaf_certificate":       resourceCudaWAFCertificate(),
			"barracudawaf_security_policy":   resourceCudaWAFSecurityPolicy(),
			"barracudawaf_rule_group_server": resourceCudaWAFRuleGroupServer(),
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
