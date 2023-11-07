package acceptance

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func checkCurrentConfig(t *testing.T, cloudType string, isAccount bool) func(s *terraform.State) error {
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
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "aws":
		workspaceLevel(t, step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "aws", false),
		})
	case "azure":
		workspaceLevel(t, step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "azure", false),
		})
	case "gcp":
		workspaceLevel(t, step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "gcp", false),
		})
	}
}

func TestMwsAccDataCurrentConfig(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	switch cloudEnv {
	case "MWS":
		accountLevel(t, step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "aws", true),
		})
	case "azure-ucacct":
		accountLevel(t, step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "azure", true),
		})
	case "gcp-accounts":
		accountLevel(t, step{
			Template: `data "databricks_current_config" "this" {}`,
			Check:    checkCurrentConfig(t, "gcp", true),
		})
	}
}
