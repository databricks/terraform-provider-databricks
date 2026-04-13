package tfschema

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// TestWorkspaceIDPlanModifier_WithProviderWorkspaceID tests the plan modifier behavior
// when workspace_id is involved. These tests document the expected behavior
// for ForceNew/RequiresReplace when the effective workspace ID changes due to
// workspace_id interaction.
//
// Currently skipped because the plan modifier only sees raw attribute values and
// doesn't have access to provider-level workspace_id. The implementation
// approach (ModifyPlan at resource level vs plan modifier config) is TBD.
func TestWorkspaceIDPlanModifier_WithProviderWorkspaceID(t *testing.T) {
	tests := []struct {
		name                    string
		stateValue              string
		planValue               string
		providerWorkspaceID     string
		expectedRequiresReplace bool
		description             string
	}{
		{
			name:                    "remove workspace_id with different default - should require replace",
			stateValue:              "100",
			planValue:               "",
			providerWorkspaceID:     "200",
			expectedRequiresReplace: true,
			description:             "Removing workspace_id when default differs should trigger replace",
		},
		{
			name:                    "remove workspace_id with same default - no replace",
			stateValue:              "100",
			planValue:               "",
			providerWorkspaceID:     "100",
			expectedRequiresReplace: false,
			description:             "Removing workspace_id when default matches should not trigger replace",
		},
		{
			name:                    "add workspace_id different from default - should require replace",
			stateValue:              "",
			planValue:               "200",
			providerWorkspaceID:     "100",
			expectedRequiresReplace: true,
			description:             "Adding workspace_id different from default should trigger replace",
		},
		{
			name:                    "add workspace_id same as default - no replace",
			stateValue:              "",
			planValue:               "100",
			providerWorkspaceID:     "100",
			expectedRequiresReplace: false,
			description:             "Adding workspace_id matching default should not trigger replace",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip("Skipped: PF plan modifier does not yet have access to workspace_id from provider config")

			// When the PF implementation is ready, this test should:
			// 1. Create a plan modifier that has access to workspace_id
			// 2. Pass state and plan values
			// 3. Assert RequiresReplace matches expectedRequiresReplace
			req := planmodifier.StringRequest{
				StateValue: types.StringValue(tt.stateValue),
				PlanValue:  types.StringValue(tt.planValue),
			}
			resp := &stringplanmodifier.RequiresReplaceIfFuncResponse{}

			workspaceIDPlanModifier(context.Background(), req, resp)

			assert.Equal(t, tt.expectedRequiresReplace, resp.RequiresReplace, tt.description)
		})
	}
}
