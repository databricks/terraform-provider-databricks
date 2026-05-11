package tfschema

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

// fakeUnifiedProviderClient records calls to ValidateWorkspaceAccess. Used to
// assert that ValidateWorkspaceID defers (does not call into the dispatcher)
// for the null/unknown/cross-resource-ref cases.
type fakeUnifiedProviderClient struct {
	providerWorkspaceID string
	currentWorkspaceID  int64
	currentErr          error

	validateCalls []string
	validateDiags diag.Diagnostics
}

func (f *fakeUnifiedProviderClient) GetProviderWorkspaceID() string { return f.providerWorkspaceID }
func (f *fakeUnifiedProviderClient) CurrentWorkspaceID(ctx context.Context) (int64, error) {
	return f.currentWorkspaceID, f.currentErr
}
func (f *fakeUnifiedProviderClient) ValidateWorkspaceAccess(ctx context.Context, workspaceID string) diag.Diagnostics {
	f.validateCalls = append(f.validateCalls, workspaceID)
	return f.validateDiags
}

func validateTestSchema() resourceschema.Schema {
	return resourceschema.Schema{
		Attributes: map[string]resourceschema.Attribute{
			"provider_config": resourceschema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]resourceschema.Attribute{
					"workspace_id": resourceschema.StringAttribute{
						Optional: true,
						Computed: true,
					},
				},
			},
		},
	}
}

func validateTfTypes() (tftypes.Object, tftypes.Object) {
	providerConfigTfType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"workspace_id": tftypes.String,
		},
	}
	rootTfType := tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"provider_config": providerConfigTfType,
		},
	}
	return providerConfigTfType, rootTfType
}

// TestValidateWorkspaceID_DefersWhenConfigNull asserts the validator returns
// without invoking ValidateWorkspaceAccess when the user did not write
// provider_config in HCL. The provider-level workspace_id is intentionally set
// non-empty to prove the validator is not falling back to it.
func TestValidateWorkspaceID_DefersWhenConfigNull(t *testing.T) {
	ctx := context.Background()
	testSchema := validateTestSchema()
	providerConfigTfType, rootTfType := validateTfTypes()

	tfConfig := tfsdk.Config{
		Schema: testSchema,
		Raw: tftypes.NewValue(rootTfType, map[string]tftypes.Value{
			"provider_config": tftypes.NewValue(providerConfigTfType, nil),
		}),
	}
	modifyReq := resource.ModifyPlanRequest{Config: tfConfig}
	resp := &resource.ModifyPlanResponse{}

	fake := &fakeUnifiedProviderClient{providerWorkspaceID: "999"}
	ValidateWorkspaceID(ctx, fake, modifyReq, resp)

	assert.False(t, resp.Diagnostics.HasError(), "expected no diagnostics, got %v", resp.Diagnostics.Errors())
	assert.Empty(t, fake.validateCalls, "ValidateWorkspaceAccess must not be called when provider_config is null")
}

// TestValidateWorkspaceID_DefersWhenInnerUnknown asserts the validator returns
// without invoking ValidateWorkspaceAccess when the user wrote provider_config
// but the workspace_id field is unknown — the cross-resource-reference shape
// (workspace_id = databricks_mws_workspaces.this.workspace_id).
func TestValidateWorkspaceID_DefersWhenInnerUnknown(t *testing.T) {
	ctx := context.Background()
	testSchema := validateTestSchema()
	providerConfigTfType, rootTfType := validateTfTypes()

	tfConfig := tfsdk.Config{
		Schema: testSchema,
		Raw: tftypes.NewValue(rootTfType, map[string]tftypes.Value{
			"provider_config": tftypes.NewValue(providerConfigTfType, map[string]tftypes.Value{
				"workspace_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
		}),
	}
	modifyReq := resource.ModifyPlanRequest{Config: tfConfig}
	resp := &resource.ModifyPlanResponse{}

	fake := &fakeUnifiedProviderClient{providerWorkspaceID: "999"}
	ValidateWorkspaceID(ctx, fake, modifyReq, resp)

	assert.False(t, resp.Diagnostics.HasError(), "expected no diagnostics, got %v", resp.Diagnostics.Errors())
	assert.Empty(t, fake.validateCalls, "ValidateWorkspaceAccess must not be called when inner workspace_id is unknown")
}

// TestValidateWorkspaceID_RunsWhenExplicit asserts the validator invokes
// ValidateWorkspaceAccess with the user-typed workspace_id when it is concrete.
func TestValidateWorkspaceID_RunsWhenExplicit(t *testing.T) {
	ctx := context.Background()
	testSchema := validateTestSchema()
	providerConfigTfType, rootTfType := validateTfTypes()

	tfConfig := tfsdk.Config{
		Schema: testSchema,
		Raw: tftypes.NewValue(rootTfType, map[string]tftypes.Value{
			"provider_config": tftypes.NewValue(providerConfigTfType, map[string]tftypes.Value{
				"workspace_id": tftypes.NewValue(tftypes.String, "12345"),
			}),
		}),
	}
	modifyReq := resource.ModifyPlanRequest{Config: tfConfig}
	resp := &resource.ModifyPlanResponse{}

	fake := &fakeUnifiedProviderClient{}
	ValidateWorkspaceID(ctx, fake, modifyReq, resp)

	assert.False(t, resp.Diagnostics.HasError(), "expected no diagnostics, got %v", resp.Diagnostics.Errors())
	assert.Equal(t, []string{"12345"}, fake.validateCalls,
		"ValidateWorkspaceAccess must be called with the explicit workspace_id")
}

// TestWorkspaceDriftDetection_DefersWhenInnerUnknown asserts that when the user
// wrote `provider_config { workspace_id = some_resource.attr }` and the ref is
// unknown at plan time, the drift detector defers — does not trigger
// RequiresReplace, does not synthesize a new plan value for provider_config,
// does not emit the "Missing workspace_id" error.
//
// Without this guard, the drift detector falls back to client.GetProviderWorkspaceID
// (empty on an account host), then to CurrentWorkspaceID (errors on account
// host), and either flags a phantom replacement or surfaces a misleading error.
// Apply will see the resolved value and act on it.
func TestWorkspaceDriftDetection_DefersWhenInnerUnknown(t *testing.T) {
	ctx := context.Background()
	testSchema := validateTestSchema()
	providerConfigTfType, rootTfType := validateTfTypes()

	// State has a concrete workspace_id "12345" — the old side that
	// WorkspaceDriftDetection compares against.
	state := tfsdk.State{
		Schema: testSchema,
		Raw: tftypes.NewValue(rootTfType, map[string]tftypes.Value{
			"provider_config": tftypes.NewValue(providerConfigTfType, map[string]tftypes.Value{
				"workspace_id": tftypes.NewValue(tftypes.String, "12345"),
			}),
		}),
	}
	// Config has provider_config.workspace_id = unknown (cross-resource ref
	// not yet resolved).
	tfConfig := tfsdk.Config{
		Schema: testSchema,
		Raw: tftypes.NewValue(rootTfType, map[string]tftypes.Value{
			"provider_config": tftypes.NewValue(providerConfigTfType, map[string]tftypes.Value{
				"workspace_id": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
		}),
	}
	// Plan starts as a copy of state (typical post-ProviderConfigPlanModifier shape).
	plan := tfsdk.Plan{
		Schema: testSchema,
		Raw: tftypes.NewValue(rootTfType, map[string]tftypes.Value{
			"provider_config": tftypes.NewValue(providerConfigTfType, map[string]tftypes.Value{
				"workspace_id": tftypes.NewValue(tftypes.String, "12345"),
			}),
		}),
	}
	modifyReq := resource.ModifyPlanRequest{State: state, Config: tfConfig, Plan: plan}
	resp := &resource.ModifyPlanResponse{Plan: plan}

	// Fake account-host client: GetProviderWorkspaceID returns empty, and
	// CurrentWorkspaceID errors (matches reality on an account host).
	fake := &fakeUnifiedProviderClient{
		currentErr: fmt.Errorf("account host has no workspace context"),
	}

	WorkspaceDriftDetection(ctx, fake, modifyReq, resp)

	assert.False(t, resp.Diagnostics.HasError(),
		"WorkspaceDriftDetection must defer when inner workspace_id is unknown; got %v",
		resp.Diagnostics.Errors())
	assert.Empty(t, resp.RequiresReplace,
		"RequiresReplace must not be set for an unknown workspace_id ref")
	assert.Empty(t, fake.validateCalls,
		"ValidateWorkspaceAccess must not be called from drift detection")
}
