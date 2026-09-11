package iamv2

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func clientWithHost(t *testing.T, host string) *common.DatabricksClient {
	t.Helper()
	return &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{Config: &config.Config{Host: host}},
	}
}

func clientWithUnifiedHost(t *testing.T) *common.DatabricksClient {
	t.Helper()
	cfg := &config.Config{
		Host: "https://unifiedhost.databricks.com",
		HostMetadataResolver: func(ctx context.Context, _ string) (*config.HostMetadata, error) {
			return &config.HostMetadata{HostType: config.UnifiedHost}, nil
		},
	}
	require.NoError(t, cfg.EnsureResolved())
	return &common.DatabricksClient{DatabricksClient: &client.DatabricksClient{Config: cfg}}
}

func TestResolveApiLevel_ExplicitAccount(t *testing.T) {
	isAccount, diags := resolveApiLevel(clientWithHost(t, "https://accounts.cloud.databricks.com"), types.StringValue("account"))
	assert.False(t, diags.HasError())
	assert.True(t, isAccount)
}

func TestResolveApiLevel_ExplicitWorkspace(t *testing.T) {
	isAccount, diags := resolveApiLevel(clientWithHost(t, "https://accounts.cloud.databricks.com"), types.StringValue("workspace"))
	assert.False(t, diags.HasError())
	assert.False(t, isAccount)
}

func TestResolveApiLevel_InvalidValue(t *testing.T) {
	_, diags := resolveApiLevel(clientWithHost(t, "https://accounts.cloud.databricks.com"), types.StringValue("bogus"))
	assert.True(t, diags.HasError())
	assert.Equal(t, diag.Diagnostics{diag.NewErrorDiagnostic("Invalid api value", "api must be either \"account\" or \"workspace\"")}, diags)
}

func TestResolveApiLevel_UnsetAccountHost(t *testing.T) {
	isAccount, diags := resolveApiLevel(clientWithHost(t, "https://accounts.cloud.databricks.com"), types.StringNull())
	assert.False(t, diags.HasError())
	assert.True(t, isAccount)
}

func TestResolveApiLevel_UnsetWorkspaceHost(t *testing.T) {
	isAccount, diags := resolveApiLevel(clientWithHost(t, "https://myworkspace.cloud.databricks.com"), types.StringNull())
	assert.False(t, diags.HasError())
	assert.False(t, isAccount)
}

func TestResolveApiLevel_UnsetUnifiedHost(t *testing.T) {
	_, diags := resolveApiLevel(clientWithUnifiedHost(t), types.StringNull())
	assert.True(t, diags.HasError())
	assert.Equal(t, diag.Diagnostics{diag.NewErrorDiagnostic("Missing api value", "please set api to account or workspace")}, diags)
}

func TestResolveApiLevel_UnknownValue(t *testing.T) {
	isAccount, diags := resolveApiLevel(clientWithHost(t, "https://accounts.cloud.databricks.com"), types.StringUnknown())
	assert.False(t, diags.HasError())
	assert.True(t, isAccount)
}

func TestUserByExternalIdSchema(t *testing.T) {
	attrs, _ := tfschema.DataSourceStructToSchemaMap(context.Background(), UserByExternalIdData{}, nil)
	assertHasAttributes(t, attrs, "external_id", "account_id", "account_user_status", "full_name", "user_id", "username", "api", "provider_config")
}

func TestGroupByExternalIdSchema(t *testing.T) {
	attrs, _ := tfschema.DataSourceStructToSchemaMap(context.Background(), GroupByExternalIdData{}, nil)
	assertHasAttributes(t, attrs, "external_id", "account_id", "group_id", "group_name", "api", "provider_config")
}

func TestServicePrincipalByExternalIdSchema(t *testing.T) {
	attrs, _ := tfschema.DataSourceStructToSchemaMap(context.Background(), ServicePrincipalByExternalIdData{}, nil)
	assertHasAttributes(t, attrs, "external_id", "account_id", "account_sp_status", "application_id", "display_name", "service_principal_id", "api", "provider_config")
}

func assertHasAttributes[V any](t *testing.T, attrs map[string]V, names ...string) {
	t.Helper()
	assert.Len(t, attrs, len(names))
	for _, name := range names {
		assert.Contains(t, attrs, name)
	}
}
