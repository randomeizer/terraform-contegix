package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/randomeizer/terraform-provider-contegixclassic/contegixclassic"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: contegixclassic.Provider,
	})
}
