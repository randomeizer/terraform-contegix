package contegixclassic

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONTEGIX_AUTH_TOKEN", nil),
				Description: "A Contegix Authentication Token.",
			},
			"custom_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONTEGIX_CUSTOM_URL", nil),
				Description: "A Custom base URL.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"contegix_classic_virtual_machine": resourceVirtualMachine(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	customUrl := d.Get("custom_url").(string)

	config := Config{
		AuthToken: d.Get("auth_token").(string),
		CustomURL: &customUrl,
	}

	return config.Client()
}
