package tfschema

import (
	"context"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestWorkspaceIDRegexValidation(t *testing.T) {
	// This regex must match the one used in ProviderConfig.ApplySchemaCustomizations
	// and ProviderConfigData.ApplySchemaCustomizations.
	workspaceIDRegex := regexp.MustCompile(`^[1-9]\d*$`)

	reject := []struct {
		name  string
		input string
	}{
		{"single zero", "0"},
		{"double zero", "00"},
		{"leading zero short", "007"},
		{"leading zero long", "0123"},
		{"alphabetic", "abc"},
		{"negative number", "-1"},
		{"decimal number", "123.456"},
		{"empty string", ""},
		{"string with spaces", "123 456"},
		{"string with hyphens", "123-456"},
	}
	for _, tc := range reject {
		t.Run("reject_"+tc.name, func(t *testing.T) {
			assert.False(t, workspaceIDRegex.MatchString(tc.input),
				"expected regex to reject %q", tc.input)
		})
	}

	accept := []struct {
		name  string
		input string
	}{
		{"single digit", "1"},
		{"two digits", "42"},
		{"three digits", "100"},
		{"large number", "123456789"},
		{"very large number", "999999999999999"},
	}
	for _, tc := range accept {
		t.Run("accept_"+tc.name, func(t *testing.T) {
			assert.True(t, workspaceIDRegex.MatchString(tc.input),
				"expected regex to accept %q", tc.input)
		})
	}
}

// TestWorkspaceIDValidatorPlanTime tests the validators configured by
// ApplySchemaCustomizations — the same validators that run during terraform plan.
// This invokes each validator's ValidateString method directly with test inputs,
// simulating what the Plugin Framework does at plan time.
func TestWorkspaceIDValidatorPlanTime(t *testing.T) {
	// Build the schema attrs through ApplySchemaCustomizations for ProviderConfig (resources)
	resourceAttrs := map[string]AttributeBuilder{
		"workspace_id": StringAttributeBuilder{},
	}
	resourceAttrs = ProviderConfig{}.ApplySchemaCustomizations(resourceAttrs)
	resourceValidators := resourceAttrs["workspace_id"].(StringAttributeBuilder).Validators

	// Build the schema attrs through ApplySchemaCustomizations for ProviderConfigData (data sources)
	dataSourceAttrs := map[string]AttributeBuilder{
		"workspace_id": StringAttributeBuilder{},
	}
	dataSourceAttrs = ProviderConfigData{}.ApplySchemaCustomizations(dataSourceAttrs)
	dataSourceValidators := dataSourceAttrs["workspace_id"].(StringAttributeBuilder).Validators

	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		// Should reject
		{"reject single zero", "0", true},
		{"reject double zero", "00", true},
		{"reject leading zero short", "007", true},
		{"reject leading zero long", "0123", true},
		{"reject alphabetic", "abc", true},
		{"reject negative number", "-1", true},
		{"reject decimal number", "123.456", true},
		{"reject string with spaces", "123 456", true},
		{"reject string with hyphens", "123-456", true},
		// Should accept
		{"accept single digit", "1", false},
		{"accept two digits", "42", false},
		{"accept three digits", "100", false},
		{"accept large number", "123456789", false},
		{"accept very large number", "999999999999999", false},
	}

	for _, tc := range tests {
		t.Run("resource_"+tc.name, func(t *testing.T) {
			hasError := runStringValidators(t, resourceValidators, tc.input)
			if tc.expectError {
				assert.True(t, hasError, "expected resource validators to reject %q", tc.input)
			} else {
				assert.False(t, hasError, "expected resource validators to accept %q", tc.input)
			}
		})
		t.Run("data_source_"+tc.name, func(t *testing.T) {
			hasError := runStringValidators(t, dataSourceValidators, tc.input)
			if tc.expectError {
				assert.True(t, hasError, "expected data source validators to reject %q", tc.input)
			} else {
				assert.False(t, hasError, "expected data source validators to accept %q", tc.input)
			}
		})
	}
}

// runStringValidators invokes all validators on the given input and returns true if any produced errors.
func runStringValidators(t *testing.T, validators []validator.String, input string) bool {
	t.Helper()
	ctx := context.Background()
	for _, v := range validators {
		req := validator.StringRequest{
			Path:        path.Root("workspace_id"),
			ConfigValue: types.StringValue(input),
		}
		resp := &validator.StringResponse{}
		v.ValidateString(ctx, req, resp)
		if resp.Diagnostics.HasError() {
			return true
		}
	}
	return false
}

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
