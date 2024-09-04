package acceptance

import (
	"testing"

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
	loadWorkspaceEnv(t)
	if isAws(t) {
		WorkspaceLevel(t, Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "aws", "false"),
		})
	} else if isAzure(t) {
		WorkspaceLevel(t, Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "azure", "false"),
		})
	} else if isGcp(t) {
		WorkspaceLevel(t, Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "gcp", "false"),
		})
	}
}

func TestMwsAccDataCurrentConfig(t *testing.T) {
	loadAccountEnv(t)
	if isAws(t) {
		AccountLevel(t, Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "aws", "true"),
		})
	} else if isAzure(t) {
		AccountLevel(t, Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "azure", "true"),
		})
	} else if isGcp(t) {
		AccountLevel(t, Step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "gcp", "true"),
		})
	}
}
