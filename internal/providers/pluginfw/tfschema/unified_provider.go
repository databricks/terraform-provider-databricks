package tfschema

import (
	"context"
	"reflect"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Namespace is used to store the namespace for unified terraform provider
// across resources and data sources onboarded to plugin framework.
// Resources and data sources will use the underlying ProviderConfig and ProviderConfigData
// type respectively to store the provider configurations.
type Namespace struct {
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

type Namespace_SdkV2 struct {
	ProviderConfig types.List `tfsdk:"provider_config"`
}

// ProviderConfig is used to store the provider configurations for unified terraform provider
// across resources onboarded to plugin framework.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].(StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(workspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^\d+$`), "workspace_id must be a valid integer"))
	return attrs
}

// workspaceIDPlanModifier is a plan modifier that requires replacement if the
// workspace_id changes from one non-empty value to another
func workspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderConfig.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type for the ProviderConfig type.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// ProviderConfigData is used to store the provider configurations for unified terraform provider
// across data sources onboarded to plugin framework.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfigData type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]AttributeBuilder) map[string]AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].(StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^\d+$`), "workspace_id must be a valid integer"))
	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ProviderConfigData.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the data source
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type for the ProviderConfigData type.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

func GetWorkspaceIDResource(ctx context.Context, providerConfig types.Object) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var workspaceID string

	if providerConfig.IsNull() {
		return workspaceID, diags
	}

	var namespace ProviderConfig
	diags.Append(providerConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return workspaceID, diags
	}

	workspaceID = namespace.WorkspaceID.ValueString()

	return workspaceID, diags
}

func GetWorkspaceIDDataSource(ctx context.Context, providerConfig types.Object) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var workspaceID string

	if providerConfig.IsNull() {
		return workspaceID, diags
	}

	var namespace ProviderConfigData
	diags.Append(providerConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return workspaceID, diags
	}

	workspaceID = namespace.WorkspaceID.ValueString()

	return workspaceID, diags
}
