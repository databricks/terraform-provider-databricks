// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_search_endpoint

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/aisearch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/aisearch_tf"
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

const dataSourceName = "ai_search_endpoint"

var _ datasource.DataSourceWithConfigure = &EndpointDataSource{}

func DataSourceEndpoint() datasource.DataSource {
	return &EndpointDataSource{}
}

type EndpointDataSource struct {
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

// EndpointData extends the main model with additional fields.
type EndpointData struct {
	// The user-selected budget policy id for the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// Time the endpoint was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Creator of the endpoint
	Creator types.String `tfsdk:"creator"`
	// The custom tags assigned to the endpoint
	CustomTags types.List `tfsdk:"custom_tags"`
	// The budget policy id applied to the endpoint
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// Current status of the endpoint
	EndpointStatus types.Object `tfsdk:"endpoint_status"`
	// Type of endpoint. Required on create and immutable thereafter.
	EndpointType types.String `tfsdk:"endpoint_type"`
	// Unique identifier of the endpoint
	Id types.String `tfsdk:"id"`
	// Number of indexes on the endpoint
	IndexCount types.Int64 `tfsdk:"index_count"`
	// User who last updated the endpoint
	LastUpdatedUser types.String `tfsdk:"last_updated_user"`
	// Name of the AI Search endpoint. Server-assigned full resource path
	// (`workspaces/{workspace}/endpoints/{endpoint}`) on output. On create, the
	// user-supplied short name is conveyed via
	// `CreateEndpointRequest.endpoint_id`; the server composes the full `name`
	// and returns it on the response.
	Name types.String `tfsdk:"name"`
	// The client-supplied desired number of replicas for the endpoint, applied
	// at create/update time. Mutually exclusive with `target_qps`.
	ReplicaCount types.Int64 `tfsdk:"replica_count"`
	// Scaling information for the endpoint
	ScalingInfo types.Object `tfsdk:"scaling_info"`
	// Target QPS for the endpoint. Mutually exclusive with `replica_count`.
	// Best-effort; the system does not guarantee this QPS will be achieved.
	TargetQps types.Int64 `tfsdk:"target_qps"`
	// Throughput information for the endpoint
	ThroughputInfo types.Object `tfsdk:"throughput_info"`
	// Time the endpoint was last updated.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// The usage policy id applied to the endpoint.
	UsagePolicyId      types.String `tfsdk:"usage_policy_id"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// EndpointData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m EndpointData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":     reflect.TypeOf(aisearch_tf.CustomTag{}),
		"endpoint_status": reflect.TypeOf(aisearch_tf.EndpointStatus{}),
		"scaling_info":    reflect.TypeOf(aisearch_tf.EndpointScalingInfo{}),
		"throughput_info": reflect.TypeOf(aisearch_tf.EndpointThroughputInfo{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointData
// only implements ToObjectValue() and Type().
func (m EndpointData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id":           m.BudgetPolicyId,
			"create_time":                m.CreateTime,
			"creator":                    m.Creator,
			"custom_tags":                m.CustomTags,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"endpoint_status":            m.EndpointStatus,
			"endpoint_type":              m.EndpointType,
			"id":                         m.Id,
			"index_count":                m.IndexCount,
			"last_updated_user":          m.LastUpdatedUser,
			"name":                       m.Name,
			"replica_count":              m.ReplicaCount,
			"scaling_info":               m.ScalingInfo,
			"target_qps":                 m.TargetQps,
			"throughput_info":            m.ThroughputInfo,
			"update_time":                m.UpdateTime,
			"usage_policy_id":            m.UsagePolicyId,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m EndpointData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"create_time":      timetypes.RFC3339{}.Type(ctx),
			"creator":          types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: aisearch_tf.CustomTag{}.Type(ctx),
			},
			"effective_budget_policy_id": types.StringType,
			"endpoint_status":            aisearch_tf.EndpointStatus{}.Type(ctx),
			"endpoint_type":              types.StringType,
			"id":                         types.StringType,
			"index_count":                types.Int64Type,
			"last_updated_user":          types.StringType,
			"name":                       types.StringType,
			"replica_count":              types.Int64Type,
			"scaling_info":               aisearch_tf.EndpointScalingInfo{}.Type(ctx),
			"target_qps":                 types.Int64Type,
			"throughput_info":            aisearch_tf.EndpointThroughputInfo{}.Type(ctx),
			"update_time":                timetypes.RFC3339{}.Type(ctx),
			"usage_policy_id":            types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

func (m EndpointData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].SetComputed()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetComputed()
	attrs["endpoint_status"] = attrs["endpoint_status"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index_count"] = attrs["index_count"].SetComputed()
	attrs["last_updated_user"] = attrs["last_updated_user"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["replica_count"] = attrs["replica_count"].SetComputed()
	attrs["scaling_info"] = attrs["scaling_info"].SetComputed()
	attrs["target_qps"] = attrs["target_qps"].SetComputed()
	attrs["throughput_info"] = attrs["throughput_info"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetComputed()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *EndpointDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *EndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, EndpointData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Endpoint",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EndpointDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *EndpointDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config EndpointData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest aisearch.GetEndpointRequest
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

	response, err := client.AiSearch.GetEndpoint(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get ai_search_endpoint", err.Error())
		return
	}

	var newState EndpointData
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
