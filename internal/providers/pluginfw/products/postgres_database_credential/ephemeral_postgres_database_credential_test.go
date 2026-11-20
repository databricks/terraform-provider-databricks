package postgres_database_credential

import (
	"context"
	"errors"
	"testing"
	"time"

	sdktime "github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const testEndpoint = "projects/my-project/branches/production/endpoints/primary"

func TestDatabaseCredentialSchema(t *testing.T) {
	resource := New()
	response := &ephemeral.SchemaResponse{}
	resource.Schema(context.Background(), ephemeral.SchemaRequest{}, response)

	assert.False(t, response.Diagnostics.HasError())
	assert.True(t, response.Schema.Attributes["endpoint"].IsRequired())
	assert.True(t, response.Schema.Attributes["ttl"].IsOptional())
	assert.True(t, response.Schema.Attributes["ttl"].IsComputed())
	assert.True(t, response.Schema.Attributes["token"].IsSensitive())
	assert.True(t, response.Schema.Attributes["token"].IsComputed())
	assert.True(t, response.Schema.Attributes["expire_time"].IsComputed())
}

func TestDatabaseCredentialMetadata(t *testing.T) {
	response := &ephemeral.MetadataResponse{}
	New().Metadata(context.Background(), ephemeral.MetadataRequest{}, response)

	assert.Equal(t, "databricks_postgres_database_credential", response.TypeName)
}

func TestDatabaseCredentialConfigure(t *testing.T) {
	resource := &databaseCredentialResource{}
	client := &common.DatabricksClient{}
	response := &ephemeral.ConfigureResponse{}
	resource.Configure(context.Background(), ephemeral.ConfigureRequest{ProviderData: client}, response)

	require.False(t, response.Diagnostics.HasError(), response.Diagnostics.Errors())
	assert.Same(t, client, resource.client)
}

func TestDatabaseCredentialConfigureRejectsUnexpectedProviderData(t *testing.T) {
	resource := &databaseCredentialResource{}
	response := &ephemeral.ConfigureResponse{}
	resource.Configure(context.Background(), ephemeral.ConfigureRequest{ProviderData: "unexpected"}, response)

	require.True(t, response.Diagnostics.HasError())
	assert.Contains(t, response.Diagnostics.Errors()[0].Detail(), "Expected *common.DatabricksClient")
}

func TestDatabaseCredentialOpen(t *testing.T) {
	expireTime := time.Date(2026, time.September, 7, 23, 0, 0, 0, time.UTC)
	qa.MockWorkspaceApply(t, func(w *mocks.MockWorkspaceClient) {
		w.GetMockPostgresAPI().EXPECT().
			GenerateDatabaseCredential(mock.Anything, mock.MatchedBy(func(request postgres.GenerateDatabaseCredentialRequest) bool {
				return request.Endpoint == testEndpoint && request.Ttl.AsDuration() == defaultTTL
			})).
			Return(&postgres.DatabaseCredential{Token: "test-token", ExpireTime: sdktime.New(expireTime)}, nil)
	}, func(ctx context.Context, client *common.DatabricksClient) {
		resource := &databaseCredentialResource{client: client}
		request, response := openRequest(t, resource, types.StringNull())

		resource.Open(ctx, request, response)

		require.False(t, response.Diagnostics.HasError(), response.Diagnostics.Errors())
		var result databaseCredentialModel
		response.Diagnostics.Append(response.Result.Get(ctx, &result)...)
		require.False(t, response.Diagnostics.HasError(), response.Diagnostics.Errors())
		assert.Equal(t, "test-token", result.Token.ValueString())
		assert.Equal(t, defaultTTL.String(), result.TTL.ValueString())
		assert.Equal(t, expireTime.Format(time.RFC3339), result.ExpireTime.ValueString())
	})
}

func TestDatabaseCredentialOpenReturnsAPIError(t *testing.T) {
	qa.MockWorkspaceApply(t, func(w *mocks.MockWorkspaceClient) {
		w.GetMockPostgresAPI().EXPECT().
			GenerateDatabaseCredential(mock.Anything, mock.Anything).
			Return(nil, errors.New("credential service unavailable"))
	}, func(ctx context.Context, client *common.DatabricksClient) {
		resource := &databaseCredentialResource{client: client}
		request, response := openRequest(t, resource, types.StringValue("10m"))

		resource.Open(ctx, request, response)

		require.True(t, response.Diagnostics.HasError())
		assert.Contains(t, response.Diagnostics.Errors()[0].Detail(), "credential service unavailable")
	})
}

func TestCredentialTTL(t *testing.T) {
	tests := []struct {
		name      string
		value     types.String
		expected  time.Duration
		hasErrors bool
	}{
		{name: "default", value: types.StringNull(), expected: time.Hour},
		{name: "minimum", value: types.StringValue("5m"), expected: 5 * time.Minute},
		{name: "maximum", value: types.StringValue("1h"), expected: time.Hour},
		{name: "too short", value: types.StringValue("299s"), hasErrors: true},
		{name: "too long", value: types.StringValue("3601s"), hasErrors: true},
		{name: "invalid", value: types.StringValue("one hour"), hasErrors: true},
		{name: "unknown", value: types.StringUnknown(), hasErrors: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, diagnostics := credentialTTL(test.value)
			assert.Equal(t, test.hasErrors, diagnostics.HasError())
			assert.Equal(t, test.expected, actual)
		})
	}
}

func openRequest(
	t *testing.T,
	resource *databaseCredentialResource,
	ttl types.String,
) (ephemeral.OpenRequest, *ephemeral.OpenResponse) {
	t.Helper()
	ctx := context.Background()
	schemaResponse := &ephemeral.SchemaResponse{}
	resource.Schema(ctx, ephemeral.SchemaRequest{}, schemaResponse)
	schemaType := schemaResponse.Schema.Type().TerraformType(ctx)
	ttlValue, err := ttl.ToTerraformValue(ctx)
	require.NoError(t, err)
	raw := tftypes.NewValue(schemaType, map[string]tftypes.Value{
		"endpoint":    tftypes.NewValue(tftypes.String, testEndpoint),
		"ttl":         ttlValue,
		"token":       tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		"expire_time": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
	})
	return ephemeral.OpenRequest{Config: tfsdk.Config{Raw: raw, Schema: schemaResponse.Schema}}, &ephemeral.OpenResponse{
		Result: tfsdk.EphemeralResultData{Raw: raw, Schema: schemaResponse.Schema},
	}
}
