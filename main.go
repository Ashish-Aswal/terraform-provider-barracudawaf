package main

import (
	"github.com/Ashish-Aswal/terraform-provider-barracudawaf/barracudawaf"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: barracudawaf.Provider,
	})
}
