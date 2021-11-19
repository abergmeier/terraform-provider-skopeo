package provider

import (
	"github.com/abergmeier/terraform-provider-skopeo/internal/providerlog"
	"github.com/abergmeier/terraform-provider-skopeo/internal/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	providerlog.SetDefault()
	provider := &schema.Provider{
		/*DataSourcesMap: map[string]*schema.Resource{
			"skopeo_copy": datasource.CopyResource(),
		},*/
		ResourcesMap: map[string]*schema.Resource{
			"skopeo_copy": resources.CopyResource(),
		},
		Schema: map[string]*schema.Schema{},
	}
	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		// Shameless plug from https://github.com/terraform-providers/terraform-provider-aws/blob/d51784148586f605ab30ecea268e80fe83d415a9/aws/provider.go
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}
	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	return nil, nil
}
