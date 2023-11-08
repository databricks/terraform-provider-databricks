package acceptance

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func noOpProvider() *schema.Provider {
	s := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"noop_noop": NoOpResource(),
		},
	}
	return s
}

func NoOpResource() *schema.Resource {
	s := &schema.Resource{}
	return s
}

func TestAccProviderPlanShouldSucceedWithoutHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"databricks": func() (*schema.Provider, error) { return provider.DatabricksProvider(), nil },
			"noop":       func() (*schema.Provider, error) { return noOpProvider(), nil },
		},
		Steps: []resource.TestStep{
			{
				Config: `
				provider "databricks" {}

				resource "noop_noop" "this" { }

				data "databricks_spark_version" "latest_lts" {
					long_term_support = true
					# depend on something so that this is known at apply
					depends_on = [noop_noop.this]
				}
				`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
