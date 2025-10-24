package sharing_test

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/stretchr/testify/require"
)

func dataSourceShareTemplate(provider_config string) string {
	return fmt.Sprintf(`
	resource "databricks_share" "myshare" {
			name  = "{var.STICKY_RANDOM}-share-config"
			object {
				name = databricks_schema.schema1.id
				data_object_type = "SCHEMA"
			}
	}
	data "databricks_share" "this" {
		%s
		name = databricks_share.myshare.name
	}
`, provider_config)
}

func TestAccShareData_ProviderConfig_Invalid(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceShareTemplate(`
			provider_config = {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(
			`Attribute provider_config\.workspace_id\s+workspace_id must be a valid integer`,
		),
		PlanOnly: true,
	})
}

func TestAccShareData_ProviderConfig_Mismatched(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceShareTemplate(`
			provider_config = {
				workspace_id = "123"
			}
		`),
		ExpectError: regexp.MustCompile(
			`(?s)failed to get workspace client.*workspace_id mismatch` +
				`.*please check the workspace_id provided in ` +
				`provider_config`,
		),
	})
}

func TestAccShareData_ProviderConfig_Apply(t *testing.T) {
	acceptance.LoadUcwsEnv(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	workspaceID, err := w.CurrentWorkspaceID(ctx)
	require.NoError(t, err)
	workspaceIDStr := strconv.FormatInt(workspaceID, 10)
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceShareTemplate(``),
	}, acceptance.Step{
		Template: preTestTemplateSchema + dataSourceShareTemplate(fmt.Sprintf(`
			provider_config = {
				workspace_id = "%s"
			}
		`, workspaceIDStr)),
	})
}
