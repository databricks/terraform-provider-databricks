// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_setting_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settingsv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settingsv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "workspace_setting_v2"

var _ datasource.DataSourceWithConfigure = &SettingDataSource{}

func DataSourceSetting() datasource.DataSource {
	return &SettingDataSource{}
}

type SettingDataSource struct {
	Client *autogen.DatabricksClient
}

// SettingData extends the main model with additional fields.
type SettingData struct {
	AibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"aibi_dashboard_embedding_access_policy"`

	AibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"aibi_dashboard_embedding_approved_domains"`

	AutomaticClusterUpdateWorkspace types.Object `tfsdk:"automatic_cluster_update_workspace"`

	BooleanVal types.Object `tfsdk:"boolean_val"`

	EffectiveAibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"effective_aibi_dashboard_embedding_access_policy"`

	EffectiveAibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"effective_aibi_dashboard_embedding_approved_domains"`

	EffectiveAutomaticClusterUpdateWorkspace types.Object `tfsdk:"effective_automatic_cluster_update_workspace"`

	EffectiveBooleanVal types.Object `tfsdk:"effective_boolean_val"`

	EffectiveIntegerVal types.Object `tfsdk:"effective_integer_val"`

	EffectivePersonalCompute types.Object `tfsdk:"effective_personal_compute"`

	EffectiveRestrictWorkspaceAdmins types.Object `tfsdk:"effective_restrict_workspace_admins"`

	EffectiveStringVal types.Object `tfsdk:"effective_string_val"`

	IntegerVal types.Object `tfsdk:"integer_val"`
	// Name of the setting.
	Name types.String `tfsdk:"name"`

	PersonalCompute types.Object `tfsdk:"personal_compute"`

	RestrictWorkspaceAdmins types.Object `tfsdk:"restrict_workspace_admins"`

	StringVal types.Object `tfsdk:"string_val"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// SettingData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m SettingData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy":              reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}),
		"aibi_dashboard_embedding_approved_domains":           reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}),
		"automatic_cluster_update_workspace":                  reflect.TypeOf(settingsv2_tf.ClusterAutoRestartMessage{}),
		"boolean_val":                                         reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_aibi_dashboard_embedding_access_policy":    reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}),
		"effective_aibi_dashboard_embedding_approved_domains": reflect.TypeOf(settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}),
		"effective_automatic_cluster_update_workspace":        reflect.TypeOf(settingsv2_tf.ClusterAutoRestartMessage{}),
		"effective_boolean_val":                               reflect.TypeOf(settingsv2_tf.BooleanMessage{}),
		"effective_integer_val":                               reflect.TypeOf(settingsv2_tf.IntegerMessage{}),
		"effective_personal_compute":                          reflect.TypeOf(settingsv2_tf.PersonalComputeMessage{}),
		"effective_restrict_workspace_admins":                 reflect.TypeOf(settingsv2_tf.RestrictWorkspaceAdminsMessage{}),
		"effective_string_val":                                reflect.TypeOf(settingsv2_tf.StringMessage{}),
		"integer_val":                                         reflect.TypeOf(settingsv2_tf.IntegerMessage{}),
		"personal_compute":                                    reflect.TypeOf(settingsv2_tf.PersonalComputeMessage{}),
		"restrict_workspace_admins":                           reflect.TypeOf(settingsv2_tf.RestrictWorkspaceAdminsMessage{}),
		"string_val":                                          reflect.TypeOf(settingsv2_tf.StringMessage{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SettingData
// only implements ToObjectValue() and Type().
func (m SettingData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_access_policy":              m.AibiDashboardEmbeddingAccessPolicy,
			"aibi_dashboard_embedding_approved_domains":           m.AibiDashboardEmbeddingApprovedDomains,
			"automatic_cluster_update_workspace":                  m.AutomaticClusterUpdateWorkspace,
			"boolean_val":                                         m.BooleanVal,
			"effective_aibi_dashboard_embedding_access_policy":    m.EffectiveAibiDashboardEmbeddingAccessPolicy,
			"effective_aibi_dashboard_embedding_approved_domains": m.EffectiveAibiDashboardEmbeddingApprovedDomains,
			"effective_automatic_cluster_update_workspace":        m.EffectiveAutomaticClusterUpdateWorkspace,
			"effective_boolean_val":                               m.EffectiveBooleanVal,
			"effective_integer_val":                               m.EffectiveIntegerVal,
			"effective_personal_compute":                          m.EffectivePersonalCompute,
			"effective_restrict_workspace_admins":                 m.EffectiveRestrictWorkspaceAdmins,
			"effective_string_val":                                m.EffectiveStringVal,
			"integer_val":                                         m.IntegerVal,
			"name":                                                m.Name,
			"personal_compute":                                    m.PersonalCompute,
			"restrict_workspace_admins":                           m.RestrictWorkspaceAdmins,
			"string_val":                                          m.StringVal,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m SettingData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"aibi_dashboard_embedding_access_policy": settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"aibi_dashboard_embedding_approved_domains":           settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"automatic_cluster_update_workspace":                  settingsv2_tf.ClusterAutoRestartMessage{}.Type(ctx),
			"boolean_val":                                         settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_aibi_dashboard_embedding_access_policy":    settingsv2_tf.AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"effective_aibi_dashboard_embedding_approved_domains": settingsv2_tf.AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"effective_automatic_cluster_update_workspace":        settingsv2_tf.ClusterAutoRestartMessage{}.Type(ctx),
			"effective_boolean_val":                               settingsv2_tf.BooleanMessage{}.Type(ctx),
			"effective_integer_val":                               settingsv2_tf.IntegerMessage{}.Type(ctx),
			"effective_personal_compute":                          settingsv2_tf.PersonalComputeMessage{}.Type(ctx),
			"effective_restrict_workspace_admins":                 settingsv2_tf.RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"effective_string_val":                                settingsv2_tf.StringMessage{}.Type(ctx),
			"integer_val":                                         settingsv2_tf.IntegerMessage{}.Type(ctx),
			"name":                                                types.StringType,
			"personal_compute":                                    settingsv2_tf.PersonalComputeMessage{}.Type(ctx),
			"restrict_workspace_admins":                           settingsv2_tf.RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"string_val":                                          settingsv2_tf.StringMessage{}.Type(ctx),
		},
	}
}

func (m SettingData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_access_policy"] = attrs["effective_aibi_dashboard_embedding_access_policy"].SetOptional()
	attrs["effective_aibi_dashboard_embedding_approved_domains"] = attrs["effective_aibi_dashboard_embedding_approved_domains"].SetOptional()
	attrs["effective_automatic_cluster_update_workspace"] = attrs["effective_automatic_cluster_update_workspace"].SetOptional()
	attrs["effective_boolean_val"] = attrs["effective_boolean_val"].SetComputed()
	attrs["effective_integer_val"] = attrs["effective_integer_val"].SetComputed()
	attrs["effective_personal_compute"] = attrs["effective_personal_compute"].SetOptional()
	attrs["effective_restrict_workspace_admins"] = attrs["effective_restrict_workspace_admins"].SetOptional()
	attrs["effective_string_val"] = attrs["effective_string_val"].SetComputed()
	attrs["integer_val"] = attrs["integer_val"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].SetOptional()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetOptional()

	return attrs
}

func (r *SettingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *SettingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, SettingData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Setting",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *SettingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *SettingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config SettingData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settingsv2.GetPublicWorkspaceSettingRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.WorkspaceSettingsV2.GetPublicWorkspaceSetting(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get workspace_setting_v2", err.Error())
		return
	}

	var newState SettingData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
