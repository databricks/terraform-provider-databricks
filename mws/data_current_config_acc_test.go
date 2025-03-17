package mws_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkCurrentConfig(t *testing.T, cloudType string, isAccount string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		r, ok := s.Modules[0].Resources["data.databricks_current_config.this"]
		require.True(t, ok, "data.databricks_current_config.this has to be there")

		attr := r.Primary.Attributes

		assert.Equal(t, attr["cloud_type"], cloudType)
		assert.Equal(t, attr["is_account"], isAccount)
		return nil
	}
}

func TestAccDataCurrentConfig(t *testing.T) {
	acceptance.LoadWorkspaceEnv(t)
	if acceptance.IsAws(t) {
		acceptance.WorkspaceLevel(t, acceptance.Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "aws", "false"),
		})
	} else if acceptance.IsAzure(t) {
		acceptance.WorkspaceLevel(t, acceptance.Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "azure", "false"),
		})
	} else if acceptance.IsGcp(t) {
		acceptance.WorkspaceLevel(t, acceptance.Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "gcp", "false"),
		})
	}
}

func TestMwsAccDataCurrentConfig(t *testing.T) {
	acceptance.LoadAccountEnv(t)
	if acceptance.IsAws(t) {
		acceptance.AccountLevel(t, acceptance.Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "aws", "true"),
		})
	} else if acceptance.IsAzure(t) {
		acceptance.AccountLevel(t, acceptance.Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "azure", "true"),
		})
	} else if acceptance.IsGcp(t) {
		acceptance.AccountLevel(t, acceptance.Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "gcp", "true"),
		})
	}
}
