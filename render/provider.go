package render

import (
	"context"

	renderRest "github.com/borkod/render-rest-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Your Render API key",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RENDER_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"render_job": resourceJob(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"render_job": dataSourceJob(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	apiKey := d.Get("api_key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client := renderRest.NewClient(apiKey)

	return client, diags
}
