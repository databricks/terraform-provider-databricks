package acceptance

import (
	"testing"
	"time"

	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

// This test ensures that, within a single Terraform module, a workspace can be created and that
// the Databricks provider can be configured to use that workspace.
func TestAccProviderPlanShouldSucceedWithIncompleteConfiguration(t *testing.T) {
	// Skip until 2026-05-08: Terraform CLI install fails with
	// "openpgp: key expired" when verifying checksum signatures.
	// We are waiting for the new version of Terraform CLI to be released and
	// be included in the JROG proxy.
	if time.Now().Before(time.Date(2026, 5, 8, 0, 0, 0, 0, time.UTC)) {
		t.Skip("temporarily skipped until 2026-05-08: Terraform CLI install fails due to expired openpgp key")
	}
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"databricks": func() (*schema.Provider, error) { return sdkv2.DatabricksProvider(), nil },
			"noop":       func() (*schema.Provider, error) { return noOpProvider(), nil },
		},
		Steps: []resource.TestStep{
			{
				Config: `
				provider "databricks" {}

				resource "noop_noop" "this" { }

				data "databricks_spark_version" "latest_lts" {
					long_term_support = true
					# depend on something so that we try to configure our provider but do not try to fetch this resource
					depends_on = [noop_noop.this]
				}
				`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
