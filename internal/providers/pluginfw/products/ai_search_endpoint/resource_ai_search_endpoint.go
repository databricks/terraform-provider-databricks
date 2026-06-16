// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ai_search_endpoint

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/aisearch"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/aisearch_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "ai_search_endpoint"

var _ resource.ResourceWithConfigure = &EndpointResource{}
var _ resource.ResourceWithModifyPlan = &EndpointResource{}

func ResourceEndpoint() resource.Resource {
	return &EndpointResource{}
}

type EndpointResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider.
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

// ApplySchemaCustomizations applies the schema customizations to the ProviderConfig type.
func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(ProviderConfigWorkspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

// ProviderConfigWorkspaceIDPlanModifier is plan modifier for the workspace_id field.
// Resource requires replacement if the workspace_id changes from one non-empty value to another.
func ProviderConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	// Require replacement if workspace_id changes from one non-empty value to another
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()

	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// ProviderConfig struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ProviderConfig
// only implements ToObjectValue() and Type().
func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// Endpoint extends the main model with additional fields.
type Endpoint struct {
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
	// The user-supplied short name for the Endpoint, per AIP-133. The server
	// composes the full `Endpoint.name` as `{parent}/endpoints/{endpoint_id}`.
	// AIP-133 does not list `endpoint_id` as a fields-may-be-required entry, so
	// we annotate it OPTIONAL on the wire; the server still rejects empty
	// values with INVALID_PARAMETER_VALUE.
	EndpointId types.String `tfsdk:"endpoint_id"`
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
	// The Workspace where this Endpoint will be created. Format:
	// `workspaces/{workspace_id}`
	Parent types.String `tfsdk:"parent"`
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
	UsagePolicyId  types.String `tfsdk:"usage_policy_id"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// Endpoint struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m Endpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":     reflect.TypeOf(aisearch_tf.CustomTag{}),
		"endpoint_status": reflect.TypeOf(aisearch_tf.EndpointStatus{}),
		"scaling_info":    reflect.TypeOf(aisearch_tf.EndpointScalingInfo{}),
		"throughput_info": reflect.TypeOf(aisearch_tf.EndpointThroughputInfo{}),
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Endpoint
// only implements ToObjectValue() and Type().
func (m Endpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"budget_policy_id": m.BudgetPolicyId,
			"create_time":                m.CreateTime,
			"creator":                    m.Creator,
			"custom_tags":                m.CustomTags,
			"effective_budget_policy_id": m.EffectiveBudgetPolicyId,
			"endpoint_id":                m.EndpointId,
			"endpoint_status":            m.EndpointStatus,
			"endpoint_type":              m.EndpointType,
			"id":                         m.Id,
			"index_count":                m.IndexCount,
			"last_updated_user":          m.LastUpdatedUser,
			"name":                       m.Name,
			"parent":                     m.Parent,
			"replica_count":              m.ReplicaCount,
			"scaling_info":               m.ScalingInfo,
			"target_qps":                 m.TargetQps,
			"throughput_info":            m.ThroughputInfo,
			"update_time":                m.UpdateTime,
			"usage_policy_id":            m.UsagePolicyId,

			"provider_config": m.ProviderConfig,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m Endpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"budget_policy_id": types.StringType,
			"create_time": timetypes.RFC3339{}.Type(ctx),
			"creator":     types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: aisearch_tf.CustomTag{}.Type(ctx),
			},
			"effective_budget_policy_id": types.StringType,
			"endpoint_id":                types.StringType,
			"endpoint_status":            aisearch_tf.EndpointStatus{}.Type(ctx),
			"endpoint_type":              types.StringType,
			"id":                         types.StringType,
			"index_count":                types.Int64Type,
			"last_updated_user":          types.StringType,
			"name":                       types.StringType,
			"parent":                     types.StringType,
			"replica_count":              types.Int64Type,
			"scaling_info":               aisearch_tf.EndpointScalingInfo{}.Type(ctx),
			"target_qps":                 types.Int64Type,
			"throughput_info":            aisearch_tf.EndpointThroughputInfo{}.Type(ctx),
			"update_time":                timetypes.RFC3339{}.Type(ctx),
			"usage_policy_id":            types.StringType,

			"provider_config": ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *Endpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Endpoint) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EndpointId.IsUnknown() {
		to.EndpointId = from.EndpointId
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				// Recursively sync the fields of EndpointStatus
				toEndpointStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.ScalingInfo.IsNull() && !from.ScalingInfo.IsUnknown() {
		if toScalingInfo, ok := to.GetScalingInfo(ctx); ok {
			if fromScalingInfo, ok := from.GetScalingInfo(ctx); ok {
				// Recursively sync the fields of ScalingInfo
				toScalingInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromScalingInfo)
				to.SetScalingInfo(ctx, toScalingInfo)
			}
		}
	}
	if !from.ThroughputInfo.IsNull() && !from.ThroughputInfo.IsUnknown() {
		if toThroughputInfo, ok := to.GetThroughputInfo(ctx); ok {
			if fromThroughputInfo, ok := from.GetThroughputInfo(ctx); ok {
				// Recursively sync the fields of ThroughputInfo
				toThroughputInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromThroughputInfo)
				to.SetThroughputInfo(ctx, toThroughputInfo)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *Endpoint) SyncFieldsDuringRead(ctx context.Context, from Endpoint) {
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EndpointId.IsUnknown() {
		to.EndpointId = from.EndpointId
	}
	if !from.EndpointStatus.IsNull() && !from.EndpointStatus.IsUnknown() {
		if toEndpointStatus, ok := to.GetEndpointStatus(ctx); ok {
			if fromEndpointStatus, ok := from.GetEndpointStatus(ctx); ok {
				toEndpointStatus.SyncFieldsDuringRead(ctx, fromEndpointStatus)
				to.SetEndpointStatus(ctx, toEndpointStatus)
			}
		}
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.ScalingInfo.IsNull() && !from.ScalingInfo.IsUnknown() {
		if toScalingInfo, ok := to.GetScalingInfo(ctx); ok {
			if fromScalingInfo, ok := from.GetScalingInfo(ctx); ok {
				toScalingInfo.SyncFieldsDuringRead(ctx, fromScalingInfo)
				to.SetScalingInfo(ctx, toScalingInfo)
			}
		}
	}
	if !from.ThroughputInfo.IsNull() && !from.ThroughputInfo.IsUnknown() {
		if toThroughputInfo, ok := to.GetThroughputInfo(ctx); ok {
			if fromThroughputInfo, ok := from.GetThroughputInfo(ctx); ok {
				toThroughputInfo.SyncFieldsDuringRead(ctx, fromThroughputInfo)
				to.SetThroughputInfo(ctx, toThroughputInfo)
			}
		}
	}
	to.ProviderConfig = from.ProviderConfig

}

func (m Endpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetComputed()
	attrs["endpoint_status"] = attrs["endpoint_status"].SetComputed()
	attrs["endpoint_type"] = attrs["endpoint_type"].SetRequired()
	attrs["endpoint_type"] = attrs["endpoint_type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["id"] = attrs["id"].SetComputed()
	attrs["index_count"] = attrs["index_count"].SetComputed()
	attrs["last_updated_user"] = attrs["last_updated_user"].SetComputed()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["replica_count"] = attrs["replica_count"].SetOptional()
	attrs["scaling_info"] = attrs["scaling_info"].SetComputed()
	attrs["target_qps"] = attrs["target_qps"].SetOptional()
	attrs["throughput_info"] = attrs["throughput_info"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetComputed()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetOptional()
	attrs["endpoint_id"] = attrs["endpoint_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["endpoint_id"] = attrs["endpoint_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplaceIf(tfschema.RequiresReplaceIfKnownChange, "", "")).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

// GetCustomTags returns the value of the CustomTags field in Endpoint as
// a slice of aisearch_tf.CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetCustomTags(ctx context.Context) ([]aisearch_tf.CustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []aisearch_tf.CustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in Endpoint.
func (m *Endpoint) SetCustomTags(ctx context.Context, v []aisearch_tf.CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEndpointStatus returns the value of the EndpointStatus field in Endpoint as
// a aisearch_tf.EndpointStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetEndpointStatus(ctx context.Context) (aisearch_tf.EndpointStatus, bool) {
	var e aisearch_tf.EndpointStatus
	if m.EndpointStatus.IsNull() || m.EndpointStatus.IsUnknown() {
		return e, false
	}
	var v aisearch_tf.EndpointStatus
	d := m.EndpointStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpointStatus sets the value of the EndpointStatus field in Endpoint.
func (m *Endpoint) SetEndpointStatus(ctx context.Context, v aisearch_tf.EndpointStatus) {
	vs := v.ToObjectValue(ctx)
	m.EndpointStatus = vs
}

// GetScalingInfo returns the value of the ScalingInfo field in Endpoint as
// a aisearch_tf.EndpointScalingInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetScalingInfo(ctx context.Context) (aisearch_tf.EndpointScalingInfo, bool) {
	var e aisearch_tf.EndpointScalingInfo
	if m.ScalingInfo.IsNull() || m.ScalingInfo.IsUnknown() {
		return e, false
	}
	var v aisearch_tf.EndpointScalingInfo
	d := m.ScalingInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScalingInfo sets the value of the ScalingInfo field in Endpoint.
func (m *Endpoint) SetScalingInfo(ctx context.Context, v aisearch_tf.EndpointScalingInfo) {
	vs := v.ToObjectValue(ctx)
	m.ScalingInfo = vs
}

// GetThroughputInfo returns the value of the ThroughputInfo field in Endpoint as
// a aisearch_tf.EndpointThroughputInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *Endpoint) GetThroughputInfo(ctx context.Context) (aisearch_tf.EndpointThroughputInfo, bool) {
	var e aisearch_tf.EndpointThroughputInfo
	if m.ThroughputInfo.IsNull() || m.ThroughputInfo.IsUnknown() {
		return e, false
	}
	var v aisearch_tf.EndpointThroughputInfo
	d := m.ThroughputInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetThroughputInfo sets the value of the ThroughputInfo field in Endpoint.
func (m *Endpoint) SetThroughputInfo(ctx context.Context, v aisearch_tf.EndpointThroughputInfo) {
	vs := v.ToObjectValue(ctx)
	m.ThroughputInfo = vs
}

func (r *EndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *EndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, Endpoint{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ai_search_endpoint",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *EndpointResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip entirely on destroy (no plan state).
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	tfschema.ValidateWorkspaceID(ctx, r.Client, req, resp)
}

func (r *EndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Endpoint
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var endpoint aisearch.Endpoint

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &endpoint)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := aisearch.CreateEndpointRequest{
		Endpoint:   endpoint,
		Parent:     plan.Parent.ValueString(),
		EndpointId: plan.EndpointId.ValueString(),
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	response, err := client.AiSearch.CreateEndpoint(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create ai_search_endpoint", err.Error())
		return
	}

	var newState Endpoint

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *EndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState Endpoint
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest aisearch.GetEndpointRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get ai_search_endpoint", err.Error())
		return
	}

	var newState Endpoint
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, existingState.ProviderConfig, &resp.State)...)
}

func (r *EndpointResource) update(ctx context.Context, plan Endpoint, diags *diag.Diagnostics, state *tfsdk.State) {
	var endpoint aisearch.Endpoint

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &endpoint)...)
	if diags.HasError() {
		return
	}

	updateRequest := aisearch.UpdateEndpointRequest{
		Endpoint:   endpoint,
		Name:       plan.Name.ValueString(),
		UpdateMask: *fieldmask.New(strings.Split("budget_policy_id,custom_tags,replica_count,target_qps,usage_policy_id", ",")),
	}

	var namespace ProviderConfig
	diags.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.AiSearch.UpdateEndpoint(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update ai_search_endpoint", err.Error())
		return
	}

	var newState Endpoint

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *EndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan Endpoint
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *EndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state Endpoint
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest aisearch.DeleteEndpointRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfig
	resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
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

	err := client.AiSearch.DeleteEndpoint(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete ai_search_endpoint", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &EndpointResource{}

func (r *EndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: name. Got: %q",
				req.ID,
			),
		)
		return
	}

	name := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
}
