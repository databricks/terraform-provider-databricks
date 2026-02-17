// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps_space

import (
	"context"
	"reflect"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/apps_tf"
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

const dataSourceName = "apps_space"

var _ datasource.DataSourceWithConfigure = &SpaceDataSource{}

func DataSourceSpace() datasource.DataSource {
	return &SpaceDataSource{}
}

type SpaceDataSource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfigData contains the fields to configure the provider.
type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
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

// SpaceData extends the main model with additional fields.
type SpaceData struct {
	// The creation time of the app space. Formatted timestamp in ISO 6801.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// The email of the user that created the app space.
	Creator types.String `tfsdk:"creator"`
	// The description of the app space.
	Description types.String `tfsdk:"description"`
	// The effective usage policy ID used by apps in the space.
	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// The effective api scopes granted to the user access token.
	EffectiveUserApiScopes types.List `tfsdk:"effective_user_api_scopes"`
	// The unique identifier of the app space.
	Id types.String `tfsdk:"id"`
	// The name of the app space. The name must contain only lowercase
	// alphanumeric characters and hyphens. It must be unique within the
	// workspace.
	Name types.String `tfsdk:"name"`
	// The OAuth2 app client ID for the app space.
	Oauth2AppClientId types.String `tfsdk:"oauth2_app_client_id"`
	// The OAuth2 app integration ID for the app space.
	Oauth2AppIntegrationId types.String `tfsdk:"oauth2_app_integration_id"`
	// Resources for the app space. Resources configured at the space level are
	// available to all apps in the space.
	Resources types.List `tfsdk:"resources"`
	// The service principal client ID for the app space.
	ServicePrincipalClientId types.String `tfsdk:"service_principal_client_id"`
	// The service principal ID for the app space.
	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`
	// The service principal name for the app space.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// The status of the app space.
	Status types.Object `tfsdk:"status"`
	// The update time of the app space. Formatted timestamp in ISO 6801.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// The email of the user that last updated the app space.
	Updater types.String `tfsdk:"updater"`
	// The usage policy ID for managing cost at the space level.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
	// OAuth scopes for apps in the space.
	UserApiScopes      types.List   `tfsdk:"user_api_scopes"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// SpaceData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m SpaceData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"effective_user_api_scopes": reflect.TypeOf(types.String{}),
		"resources":                 reflect.TypeOf(apps_tf.AppResource{}),
		"status":                    reflect.TypeOf(apps_tf.SpaceStatus{}),
		"user_api_scopes":           reflect.TypeOf(types.String{}),
		"provider_config":           reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SpaceData
// only implements ToObjectValue() and Type().
func (m SpaceData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":                 m.CreateTime,
			"creator":                     m.Creator,
			"description":                 m.Description,
			"effective_usage_policy_id":   m.EffectiveUsagePolicyId,
			"effective_user_api_scopes":   m.EffectiveUserApiScopes,
			"id":                          m.Id,
			"name":                        m.Name,
			"oauth2_app_client_id":        m.Oauth2AppClientId,
			"oauth2_app_integration_id":   m.Oauth2AppIntegrationId,
			"resources":                   m.Resources,
			"service_principal_client_id": m.ServicePrincipalClientId,
			"service_principal_id":        m.ServicePrincipalId,
			"service_principal_name":      m.ServicePrincipalName,
			"status":                      m.Status,
			"update_time":                 m.UpdateTime,
			"updater":                     m.Updater,
			"usage_policy_id":             m.UsagePolicyId,
			"user_api_scopes":             m.UserApiScopes,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m SpaceData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":               timetypes.RFC3339{}.Type(ctx),
			"creator":                   types.StringType,
			"description":               types.StringType,
			"effective_usage_policy_id": types.StringType,
			"effective_user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"id":                        types.StringType,
			"name":                      types.StringType,
			"oauth2_app_client_id":      types.StringType,
			"oauth2_app_integration_id": types.StringType,
			"resources": basetypes.ListType{
				ElemType: apps_tf.AppResource{}.Type(ctx),
			},
			"service_principal_client_id": types.StringType,
			"service_principal_id":        types.Int64Type,
			"service_principal_name":      types.StringType,
			"status":                      apps_tf.SpaceStatus{}.Type(ctx),
			"update_time":                 timetypes.RFC3339{}.Type(ctx),
			"updater":                     types.StringType,
			"usage_policy_id":             types.StringType,
			"user_api_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m SpaceData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetComputed()
	attrs["effective_user_api_scopes"] = attrs["effective_user_api_scopes"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["oauth2_app_client_id"] = attrs["oauth2_app_client_id"].SetComputed()
	attrs["oauth2_app_integration_id"] = attrs["oauth2_app_integration_id"].SetComputed()
	attrs["resources"] = attrs["resources"].SetComputed()
	attrs["service_principal_client_id"] = attrs["service_principal_client_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["updater"] = attrs["updater"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetComputed()
	attrs["user_api_scopes"] = attrs["user_api_scopes"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *SpaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *SpaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SpaceData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Space",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SpaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SpaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config SpaceData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest apps.GetSpaceRequest
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

	response, err := client.Apps.GetSpace(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get apps_space", err.Error())
		return
	}

	var newState SpaceData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Preserve provider_config from config since it's not part of the API response
	newState.ProviderConfigData = config.ProviderConfigData

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
