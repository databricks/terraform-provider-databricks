// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package environments_workspace_base_environment

import (
	"context"
	"reflect"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/environments"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "environments_workspace_base_environment"

var _ datasource.DataSourceWithConfigure = &WorkspaceBaseEnvironmentDataSource{}

func DataSourceWorkspaceBaseEnvironment() datasource.DataSource {
	return &WorkspaceBaseEnvironmentDataSource{}
}

type WorkspaceBaseEnvironmentDataSource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()

	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(
		stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9]\d*$`), "workspace_id must be a positive integer without leading zeros"))
	return attrs
}

// ProviderConfigDataWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigDataWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfigData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfigData
// only implements ToObjectValue() and Type().
func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// WorkspaceBaseEnvironmentData extends the main model with additional fields.
type WorkspaceBaseEnvironmentData struct {
	// The type of base environment (CPU or GPU).
	BaseEnvironmentType          types.String `tfsdk:"base_environment_type"`
	EffectiveBaseEnvironmentType types.String `tfsdk:"effective_base_environment_type"`
	// Timestamp when the environment was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// User ID of the creator.
	CreatorUserId types.String `tfsdk:"creator_user_id"`
	// Human-readable display name for the workspace base environment.
	DisplayName types.String `tfsdk:"display_name"`
	// The WSFS or UC Volumes path to the environment YAML file.
	Filepath types.String `tfsdk:"filepath"`
	// Whether this is the default environment for the workspace.
	IsDefault types.Bool `tfsdk:"is_default"`
	// User ID of the last user who updated the environment.
	LastUpdatedUserId types.String `tfsdk:"last_updated_user_id"`
	// Status message providing additional details about the environment status.
	Message types.String `tfsdk:"message"`
	// The resource name of the workspace base environment. Format:
	// workspace-base-environments/{workspace-base-environment}
	Name types.String `tfsdk:"name"`
	// The status of the materialized workspace base environment.
	Status types.String `tfsdk:"status"`
	// Timestamp when the environment was last updated.
	UpdateTime         timetypes.RFC3339 `tfsdk:"update_time"`
	ProviderConfigData types.Object      `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// WorkspaceBaseEnvironmentData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m WorkspaceBaseEnvironmentData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceBaseEnvironmentData
// only implements ToObjectValue() and Type().
func (m WorkspaceBaseEnvironmentData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"base_environment_type": m.BaseEnvironmentType, "effective_base_environment_type": m.EffectiveBaseEnvironmentType,
			"create_time":          m.CreateTime,
			"creator_user_id":      m.CreatorUserId,
			"display_name":         m.DisplayName,
			"filepath":             m.Filepath,
			"is_default":           m.IsDefault,
			"last_updated_user_id": m.LastUpdatedUserId,
			"message":              m.Message,
			"name":                 m.Name,
			"status":               m.Status,
			"update_time":          m.UpdateTime,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m WorkspaceBaseEnvironmentData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_environment_type":           types.StringType,
			"effective_base_environment_type": types.StringType,
			"create_time":                     timetypes.RFC3339{}.Type(ctx),
			"creator_user_id":                 types.StringType,
			"display_name":                    types.StringType,
			"filepath":                        types.StringType,
			"is_default":                      types.BoolType,
			"last_updated_user_id":            types.StringType,
			"message":                         types.StringType,
			"name":                            types.StringType,
			"status":                          types.StringType,
			"update_time":                     timetypes.RFC3339{}.Type(ctx),

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m WorkspaceBaseEnvironmentData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["base_environment_type"] = attrs["base_environment_type"].SetComputed()
	attrs["effective_base_environment_type"] = attrs["effective_base_environment_type"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator_user_id"] = attrs["creator_user_id"].SetComputed()
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["filepath"] = attrs["filepath"].SetComputed()
	attrs["is_default"] = attrs["is_default"].SetComputed()
	attrs["last_updated_user_id"] = attrs["last_updated_user_id"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *WorkspaceBaseEnvironmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *WorkspaceBaseEnvironmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceBaseEnvironmentData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks WorkspaceBaseEnvironment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceBaseEnvironmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *WorkspaceBaseEnvironmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config WorkspaceBaseEnvironmentData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest environments.GetWorkspaceBaseEnvironmentRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Environments.GetWorkspaceBaseEnvironment(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get environments_workspace_base_environment", err.Error())
		return
	}

	var newState WorkspaceBaseEnvironmentData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config so state.Set has the correct type info
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInStateForDataSource(ctx, r.Client, config.ProviderConfigData, &resp.State)...)
}
