// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disaster_recovery_failover_group

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/service/disasterrecovery"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/declarative"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/disasterrecovery_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "disaster_recovery_failover_group"

var _ resource.ResourceWithConfigure = &FailoverGroupResource{}

func ResourceFailoverGroup() resource.Resource {
	return &FailoverGroupResource{}
}

type FailoverGroupResource struct {
	Client *autogen.DatabricksClient
}

// FailoverGroup extends the main model with additional fields.
type FailoverGroup struct {
	// Time at which this failover group was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Current effective primary region. Replication flows FROM workspaces in
	// this region. Changes after a successful failover.
	EffectivePrimaryRegion types.String `tfsdk:"effective_primary_region"`
	// Opaque version string for optimistic locking. Server-generated, returned
	// in responses. Must be provided on Update requests to prevent concurrent
	// modifications.
	Etag types.String `tfsdk:"etag"`
	// Client-provided identifier for the failover group. Used to construct the
	// resource name as {parent}/failover-groups/{failover_group_id}.
	FailoverGroupId types.String `tfsdk:"failover_group_id"`
	// Initial primary region. Used only in Create requests to set the starting
	// primary region. Not returned in responses.
	InitialPrimaryRegion types.String `tfsdk:"initial_primary_region"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"name"`
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"parent"`
	// List of all regions participating in this failover group.
	Regions types.List `tfsdk:"regions"`
	// The latest point in time to which data has been replicated.
	ReplicationPoint timetypes.RFC3339 `tfsdk:"replication_point"`
	// Aggregate state of the failover group.
	State types.String `tfsdk:"state"`
	// Unity Catalog replication configuration.
	UnityCatalogAssets types.Object `tfsdk:"unity_catalog_assets"`
	// Time at which this failover group was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Workspace sets, each containing workspaces that replicate to each other.
	WorkspaceSets types.List `tfsdk:"workspace_sets"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// FailoverGroup struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m FailoverGroup) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"regions":              reflect.TypeOf(types.String{}),
		"unity_catalog_assets": reflect.TypeOf(disasterrecovery_tf.UcReplicationConfig{}),
		"workspace_sets":       reflect.TypeOf(disasterrecovery_tf.WorkspaceSet{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailoverGroup
// only implements ToObjectValue() and Type().
func (m FailoverGroup) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"create_time": m.CreateTime,
			"effective_primary_region": m.EffectivePrimaryRegion,
			"etag":                     m.Etag,
			"failover_group_id":        m.FailoverGroupId,
			"initial_primary_region":   m.InitialPrimaryRegion,
			"name":                     m.Name,
			"parent":                   m.Parent,
			"regions":                  m.Regions,
			"replication_point":        m.ReplicationPoint,
			"state":                    m.State,
			"unity_catalog_assets":     m.UnityCatalogAssets,
			"update_time":              m.UpdateTime,
			"workspace_sets":           m.WorkspaceSets,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FailoverGroup) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"create_time": timetypes.RFC3339{}.Type(ctx),
			"effective_primary_region": types.StringType,
			"etag":                     types.StringType,
			"failover_group_id":        types.StringType,
			"initial_primary_region":   types.StringType,
			"name":                     types.StringType,
			"parent":                   types.StringType,
			"regions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"replication_point":    timetypes.RFC3339{}.Type(ctx),
			"state":                types.StringType,
			"unity_catalog_assets": disasterrecovery_tf.UcReplicationConfig{}.Type(ctx),
			"update_time":          timetypes.RFC3339{}.Type(ctx),
			"workspace_sets": basetypes.ListType{
				ElemType: disasterrecovery_tf.WorkspaceSet{}.Type(ctx),
			},
		},
	}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (to *FailoverGroup) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FailoverGroup) {
	if !from.FailoverGroupId.IsUnknown() {
		to.FailoverGroupId = from.FailoverGroupId
	}
	if !from.InitialPrimaryRegion.IsUnknown() && !from.InitialPrimaryRegion.IsNull() {
		// InitialPrimaryRegion is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialPrimaryRegion = from.InitialPrimaryRegion
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.UnityCatalogAssets.IsNull() && !from.UnityCatalogAssets.IsUnknown() {
		if toUnityCatalogAssets, ok := to.GetUnityCatalogAssets(ctx); ok {
			if fromUnityCatalogAssets, ok := from.GetUnityCatalogAssets(ctx); ok {
				// Recursively sync the fields of UnityCatalogAssets
				toUnityCatalogAssets.SyncFieldsDuringCreateOrUpdate(ctx, fromUnityCatalogAssets)
				to.SetUnityCatalogAssets(ctx, toUnityCatalogAssets)
			}
		}
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *FailoverGroup) SyncFieldsDuringRead(ctx context.Context, from FailoverGroup) {
	if !from.FailoverGroupId.IsUnknown() {
		to.FailoverGroupId = from.FailoverGroupId
	}
	if !from.InitialPrimaryRegion.IsUnknown() && !from.InitialPrimaryRegion.IsNull() {
		// InitialPrimaryRegion is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialPrimaryRegion = from.InitialPrimaryRegion
	}
	if !from.Parent.IsUnknown() {
		to.Parent = from.Parent
	}
	if !from.UnityCatalogAssets.IsNull() && !from.UnityCatalogAssets.IsUnknown() {
		if toUnityCatalogAssets, ok := to.GetUnityCatalogAssets(ctx); ok {
			if fromUnityCatalogAssets, ok := from.GetUnityCatalogAssets(ctx); ok {
				toUnityCatalogAssets.SyncFieldsDuringRead(ctx, fromUnityCatalogAssets)
				to.SetUnityCatalogAssets(ctx, toUnityCatalogAssets)
			}
		}
	}
}

func (m FailoverGroup) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["effective_primary_region"] = attrs["effective_primary_region"].SetComputed()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["initial_primary_region"] = attrs["initial_primary_region"].SetRequired()
	attrs["initial_primary_region"] = attrs["initial_primary_region"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["regions"] = attrs["regions"].SetRequired()
	attrs["regions"] = attrs["regions"].(tfschema.ListAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["replication_point"] = attrs["replication_point"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["unity_catalog_assets"] = attrs["unity_catalog_assets"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_sets"] = attrs["workspace_sets"].SetRequired()
	attrs["failover_group_id"] = attrs["failover_group_id"].SetRequired()
	attrs["failover_group_id"] = attrs["failover_group_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["failover_group_id"] = attrs["failover_group_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["parent"] = attrs["parent"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	return attrs
}

// GetRegions returns the value of the Regions field in FailoverGroup as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *FailoverGroup) GetRegions(ctx context.Context) ([]types.String, bool) {
	if m.Regions.IsNull() || m.Regions.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Regions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRegions sets the value of the Regions field in FailoverGroup.
func (m *FailoverGroup) SetRegions(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["regions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Regions = types.ListValueMust(t, vs)
}

// GetUnityCatalogAssets returns the value of the UnityCatalogAssets field in FailoverGroup as
// a disasterrecovery_tf.UcReplicationConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *FailoverGroup) GetUnityCatalogAssets(ctx context.Context) (disasterrecovery_tf.UcReplicationConfig, bool) {
	var e disasterrecovery_tf.UcReplicationConfig
	if m.UnityCatalogAssets.IsNull() || m.UnityCatalogAssets.IsUnknown() {
		return e, false
	}
	var v disasterrecovery_tf.UcReplicationConfig
	d := m.UnityCatalogAssets.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUnityCatalogAssets sets the value of the UnityCatalogAssets field in FailoverGroup.
func (m *FailoverGroup) SetUnityCatalogAssets(ctx context.Context, v disasterrecovery_tf.UcReplicationConfig) {
	vs := v.ToObjectValue(ctx)
	m.UnityCatalogAssets = vs
}

// GetWorkspaceSets returns the value of the WorkspaceSets field in FailoverGroup as
// a slice of disasterrecovery_tf.WorkspaceSet values.
// If the field is unknown or null, the boolean return value is false.
func (m *FailoverGroup) GetWorkspaceSets(ctx context.Context) ([]disasterrecovery_tf.WorkspaceSet, bool) {
	if m.WorkspaceSets.IsNull() || m.WorkspaceSets.IsUnknown() {
		return nil, false
	}
	var v []disasterrecovery_tf.WorkspaceSet
	d := m.WorkspaceSets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceSets sets the value of the WorkspaceSets field in FailoverGroup.
func (m *FailoverGroup) SetWorkspaceSets(ctx context.Context, v []disasterrecovery_tf.WorkspaceSet) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_sets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceSets = types.ListValueMust(t, vs)
}

func (r *FailoverGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *FailoverGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, FailoverGroup{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks disaster_recovery_failover_group",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FailoverGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *FailoverGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan FailoverGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var failover_group disasterrecovery.FailoverGroup

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &failover_group)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := disasterrecovery.CreateFailoverGroupRequest{
		FailoverGroup:   failover_group,
		Parent:          plan.Parent.ValueString(),
		FailoverGroupId: plan.FailoverGroupId.ValueString(),
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DisasterRecovery.CreateFailoverGroup(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create disaster_recovery_failover_group", err.Error())
		return
	}

	var newState FailoverGroup

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *FailoverGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState FailoverGroup
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest disasterrecovery.GetFailoverGroupRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := client.DisasterRecovery.GetFailoverGroup(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get disaster_recovery_failover_group", err.Error())
		return
	}

	var newState FailoverGroup
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *FailoverGroupResource) update(ctx context.Context, plan FailoverGroup, diags *diag.Diagnostics, state *tfsdk.State) {
	var failover_group disasterrecovery.FailoverGroup

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &failover_group)...)
	if diags.HasError() {
		return
	}

	updateRequest := disasterrecovery.UpdateFailoverGroupRequest{
		FailoverGroup: failover_group,
		Name:          plan.Name.ValueString(),
		UpdateMask:    *fieldmask.New(strings.Split("unity_catalog_assets,workspace_sets", ",")),
	}

	client, clientDiags := r.Client.GetAccountClient()

	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	response, err := client.DisasterRecovery.UpdateFailoverGroup(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update disaster_recovery_failover_group", err.Error())
		return
	}

	var newState FailoverGroup

	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *FailoverGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan FailoverGroup
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *FailoverGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state FailoverGroup
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest disasterrecovery.DeleteFailoverGroupRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.DisasterRecovery.DeleteFailoverGroup(ctx, deleteRequest)
	if !declarative.IsDeleteError(err) {
		err = nil
	}
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete disaster_recovery_failover_group", err.Error())
		return
	}

}

var _ resource.ResourceWithImportState = &FailoverGroupResource{}

func (r *FailoverGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
