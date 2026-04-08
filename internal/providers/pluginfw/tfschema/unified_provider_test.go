package tfschema

import (
	"context"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resource_schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceIDPlanModifier(t *testing.T) {
	tests := []struct {
		name                    string
		stateValue              string
		planValue               string
		expectedRequiresReplace bool
	}{
		{
			name:                    "both non-empty and different - requires replace",
			stateValue:              "workspace-123",
			planValue:               "workspace-456",
			expectedRequiresReplace: true,
		},
		{
			name:                    "both non-empty and same - no replace",
			stateValue:              "workspace-123",
			planValue:               "workspace-123",
			expectedRequiresReplace: false,
		},
		{
			name:                    "old empty, new non-empty - no replace",
			stateValue:              "",
			planValue:               "workspace-123",
			expectedRequiresReplace: false,
		},
		{
			name:                    "old non-empty, new empty - no replace",
			stateValue:              "workspace-123",
			planValue:               "",
			expectedRequiresReplace: false,
		},
		{
			name:                    "both empty - no replace",
			stateValue:              "",
			planValue:               "",
			expectedRequiresReplace: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := planmodifier.StringRequest{
				StateValue: types.StringValue(tt.stateValue),
				PlanValue:  types.StringValue(tt.planValue),
			}
			resp := &stringplanmodifier.RequiresReplaceIfFuncResponse{}

			workspaceIDPlanModifier(context.Background(), req, resp)

			assert.Equal(t, tt.expectedRequiresReplace, resp.RequiresReplace,
				"RequiresReplace mismatch for state '%s' -> plan '%s'",
				tt.stateValue, tt.planValue)
		})
	}
}

func TestGetWorkspaceID_SdkV2(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                string
		setupProviderConfig func() types.List
		expectedWorkspaceID string
		expectError         bool
	}{
		{
			name: "valid workspace ID",
			setupProviderConfig: func() types.List {
				providerConfig := ProviderConfig{
					WorkspaceID: types.StringValue("123456789"),
				}
				return types.ListValueMust(
					ProviderConfig{}.Type(ctx),
					[]attr.Value{providerConfig.ToObjectValue(ctx)},
				)
			},
			expectedWorkspaceID: "123456789",
			expectError:         false,
		},
		{
			name: "null provider_config",
			setupProviderConfig: func() types.List {
				return types.ListNull(ProviderConfig{}.Type(ctx))
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "unknown provider_config",
			setupProviderConfig: func() types.List {
				return types.ListUnknown(ProviderConfig{}.Type(ctx))
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "empty list",
			setupProviderConfig: func() types.List {
				return types.ListValueMust(
					ProviderConfig{}.Type(ctx),
					[]attr.Value{},
				)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			providerConfigList := tt.setupProviderConfig()
			workspaceID, diags := GetWorkspaceID_SdkV2(ctx, providerConfigList)

			if tt.expectError {
				assert.True(t, diags.HasError(), "Expected diagnostics error")
			} else {
				assert.False(t, diags.HasError(), "Expected no diagnostics error")
			}
			assert.Equal(t, tt.expectedWorkspaceID, workspaceID, "Workspace ID mismatch")
		})
	}
}

func TestGetWorkspaceIDResource(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                string
		setupProviderConfig func() types.Object
		expectedWorkspaceID string
		expectError         bool
	}{
		{
			name: "valid workspace ID",
			setupProviderConfig: func() types.Object {
				providerConfig := ProviderConfig{
					WorkspaceID: types.StringValue("123456789"),
				}
				return providerConfig.ToObjectValue(ctx)
			},
			expectedWorkspaceID: "123456789",
			expectError:         false,
		},
		{
			name: "null provider_config",
			setupProviderConfig: func() types.Object {
				return types.ObjectNull(ProviderConfig{}.Type(ctx).(types.ObjectType).AttrTypes)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "unknown provider_config",
			setupProviderConfig: func() types.Object {
				return types.ObjectUnknown(ProviderConfig{}.Type(ctx).(types.ObjectType).AttrTypes)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "empty workspace ID string",
			setupProviderConfig: func() types.Object {
				providerConfig := ProviderConfig{
					WorkspaceID: types.StringValue(""),
				}
				return providerConfig.ToObjectValue(ctx)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "null workspace ID in object",
			setupProviderConfig: func() types.Object {
				providerConfig := ProviderConfig{
					WorkspaceID: types.StringNull(),
				}
				return providerConfig.ToObjectValue(ctx)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "incompatible object structure",
			setupProviderConfig: func() types.Object {
				// Create an object with wrong attribute types to trigger conversion error
				attrTypes := map[string]attr.Type{
					"workspace_id": types.Int64Type,
				}
				attrValues := map[string]attr.Value{
					"workspace_id": types.Int64Value(123),
				}
				return types.ObjectValueMust(attrTypes, attrValues)
			},
			expectedWorkspaceID: "",
			expectError:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			providerConfigObject := tt.setupProviderConfig()
			workspaceID, diags := GetWorkspaceIDResource(ctx, providerConfigObject)

			if tt.expectError {
				assert.True(t, diags.HasError(), "Expected diagnostics error")
			} else {
				assert.False(t, diags.HasError(), "Expected no diagnostics error")
			}
			assert.Equal(t, tt.expectedWorkspaceID, workspaceID, "Workspace ID mismatch")
		})
	}
}

func TestGetWorkspaceIDDataSource(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                string
		setupProviderConfig func() types.Object
		expectedWorkspaceID string
		expectError         bool
	}{
		{
			name: "valid workspace ID",
			setupProviderConfig: func() types.Object {
				providerConfig := ProviderConfigData{
					WorkspaceID: types.StringValue("123456789"),
				}
				return providerConfig.ToObjectValue(ctx)
			},
			expectedWorkspaceID: "123456789",
			expectError:         false,
		},
		{
			name: "null provider_config",
			setupProviderConfig: func() types.Object {
				return types.ObjectNull(ProviderConfigData{}.Type(ctx).(types.ObjectType).AttrTypes)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "unknown provider_config",
			setupProviderConfig: func() types.Object {
				return types.ObjectUnknown(ProviderConfigData{}.Type(ctx).(types.ObjectType).AttrTypes)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "empty workspace ID string",
			setupProviderConfig: func() types.Object {
				providerConfig := ProviderConfigData{
					WorkspaceID: types.StringValue(""),
				}
				return providerConfig.ToObjectValue(ctx)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "null workspace ID in object",
			setupProviderConfig: func() types.Object {
				providerConfig := ProviderConfigData{
					WorkspaceID: types.StringNull(),
				}
				return providerConfig.ToObjectValue(ctx)
			},
			expectedWorkspaceID: "",
			expectError:         false,
		},
		{
			name: "incompatible object structure",
			setupProviderConfig: func() types.Object {
				// Create an object with wrong attribute types to trigger conversion error
				attrTypes := map[string]attr.Type{
					"workspace_id": types.Int64Type,
				}
				attrValues := map[string]attr.Value{
					"workspace_id": types.Int64Value(123),
				}
				return types.ObjectValueMust(attrTypes, attrValues)
			},
			expectedWorkspaceID: "",
			expectError:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			providerConfigObject := tt.setupProviderConfig()
			workspaceID, diags := GetWorkspaceIDDataSource(ctx, providerConfigObject)

			if tt.expectError {
				assert.True(t, diags.HasError(), "Expected diagnostics error")
			} else {
				assert.False(t, diags.HasError(), "Expected no diagnostics error")
			}
			assert.Equal(t, tt.expectedWorkspaceID, workspaceID, "Workspace ID mismatch")
		})
	}
}

// TestValidateWorkspaceID verifies that ValidateWorkspaceID reads the workspace_id
// from resp.Plan (which WorkspaceDriftDetection may have updated) rather than
// req.Plan (which still has the old value preserved by the plan modifier).
func TestValidateWorkspaceID(t *testing.T) {
	ctx := context.Background()

	testSchema := resource_schema.Schema{
		Attributes: map[string]resource_schema.Attribute{
			"provider_config": resource_schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]resource_schema.Attribute{
					"workspace_id": resource_schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}

	pcObjectType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"workspace_id": tftypes.String,
		},
	}
	rootType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"provider_config": pcObjectType,
		},
	}
	makePC := func(wsID string) tftypes.Value {
		return tftypes.NewValue(pcObjectType, map[string]tftypes.Value{
			"workspace_id": tftypes.NewValue(tftypes.String, wsID),
		})
	}
	makeRoot := func(pc tftypes.Value) tftypes.Value {
		return tftypes.NewValue(rootType, map[string]tftypes.Value{
			"provider_config": pc,
		})
	}
	nullPC := tftypes.NewValue(pcObjectType, nil)

	newMockClient := func(t *testing.T) *common.DatabricksClient {
		cfg := &config.Config{
			Host:      "https://accounts.cloud.databricks.com",
			AccountID: "test-account",
			Token:     "test-token",
		}
		c, err := client.New(cfg)
		assert.NoError(t, err)
		return &common.DatabricksClient{DatabricksClient: c}
	}

	tests := []struct {
		name                     string
		reqPlanWSID              string
		respPlanWSID             string
		stateWSID                string
		expectError              bool
		expectedErrorContains    string
		expectedErrorNotContains string
	}{
		{
			name:                     "reads resp.Plan not req.Plan when they differ",
			reqPlanWSID:              "111",
			respPlanWSID:             "999",
			stateWSID:                "111",
			expectError:              true,
			expectedErrorContains:    "999",
			expectedErrorNotContains: "111",
		},
		{
			name:                  "validates workspace_id when both plans agree",
			reqPlanWSID:           "555",
			respPlanWSID:          "555",
			stateWSID:             "555",
			expectError:           true,
			expectedErrorContains: "555",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := resource.ModifyPlanRequest{
				Plan:   tfsdk.Plan{Schema: testSchema, Raw: makeRoot(makePC(tt.reqPlanWSID))},
				State:  tfsdk.State{Schema: testSchema, Raw: makeRoot(makePC(tt.stateWSID))},
				Config: tfsdk.Config{Schema: testSchema, Raw: makeRoot(nullPC)},
			}
			resp := &resource.ModifyPlanResponse{
				Plan: tfsdk.Plan{Schema: testSchema, Raw: makeRoot(makePC(tt.respPlanWSID))},
			}

			ValidateWorkspaceID(ctx, newMockClient(t), req, resp)

			if tt.expectError {
				assert.True(t, resp.Diagnostics.HasError(), "Expected error from ValidateWorkspaceID")
			} else {
				assert.False(t, resp.Diagnostics.HasError(), "Expected no error from ValidateWorkspaceID")
			}

			if tt.expectedErrorContains != "" {
				found := false
				for _, d := range resp.Diagnostics.Errors() {
					if strings.Contains(d.Detail(), tt.expectedErrorContains) {
						found = true
						break
					}
				}
				assert.True(t, found,
					"Expected error containing %q, got: %v", tt.expectedErrorContains, resp.Diagnostics.Errors())
			}

			if tt.expectedErrorNotContains != "" {
				for _, d := range resp.Diagnostics.Errors() {
					assert.False(t, strings.Contains(d.Detail(), tt.expectedErrorNotContains),
						"Error should not contain %q, got: %s", tt.expectedErrorNotContains, d.Detail())
				}
			}
		})
	}
}
