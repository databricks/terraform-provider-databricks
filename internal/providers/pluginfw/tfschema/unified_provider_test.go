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
