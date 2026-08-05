package iamv2_test

import (
	"fmt"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/require"
)

// These tests require the workspace/account to be onboarded to Automatic
// Identity Management (AIM) and a real external ID from the configured IdP,
// so they cannot provision their own fixture data and are gated on an env var.
func userByExternalIdTemplate(externalId string) string {
	return fmt.Sprintf(`
		data "databricks_user_by_external_id" "this" {
			external_id = "%s"
		}
	`, externalId)
}

func checkUserByExternalIdPopulated(t *testing.T, externalId string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.Modules[0].Resources["data.databricks_user_by_external_id.this"]
		require.True(t, ok, "data.databricks_user_by_external_id.this has to be there")
		require.Equal(t, externalId, ds.Primary.Attributes["external_id"])
		require.NotEmpty(t, ds.Primary.Attributes["user_id"])
		require.NotEmpty(t, ds.Primary.Attributes["username"])
		return nil
	}
}

func TestAccDataSourceUserByExternalId(t *testing.T) {
	externalId := acceptance.GetEnvOrSkipTest(t, "TEST_IDP_EXTERNAL_USER_ID")
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: userByExternalIdTemplate(externalId),
		Check:    checkUserByExternalIdPopulated(t, externalId),
	})
}

func TestMwsAccDataSourceUserByExternalId(t *testing.T) {
	externalId := acceptance.GetEnvOrSkipTest(t, "TEST_IDP_EXTERNAL_USER_ID")
	acceptance.AccountLevel(t, acceptance.Step{
		Template: userByExternalIdTemplate(externalId),
		Check:    checkUserByExternalIdPopulated(t, externalId),
	})
}
