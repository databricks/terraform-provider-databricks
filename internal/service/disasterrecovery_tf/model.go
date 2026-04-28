// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package disasterrecovery_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateFailoverGroupRequest struct {
	// The failover group to create.
	FailoverGroup types.Object `tfsdk:"failover_group"`
	// Client-provided identifier for the failover group. Used to construct the
	// resource name as {parent}/failover-groups/{failover_group_id}.
	FailoverGroupId types.String `tfsdk:"-"`
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"-"`
	// When true, validates the request without creating the failover group.
	ValidateOnly types.Bool `tfsdk:"-"`
}

func (to *CreateFailoverGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateFailoverGroupRequest) {
	if !from.FailoverGroup.IsNull() && !from.FailoverGroup.IsUnknown() {
		if toFailoverGroup, ok := to.GetFailoverGroup(ctx); ok {
			if fromFailoverGroup, ok := from.GetFailoverGroup(ctx); ok {
				// Recursively sync the fields of FailoverGroup
				toFailoverGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromFailoverGroup)
				to.SetFailoverGroup(ctx, toFailoverGroup)
			}
		}
	}
}

func (to *CreateFailoverGroupRequest) SyncFieldsDuringRead(ctx context.Context, from CreateFailoverGroupRequest) {
	if !from.FailoverGroup.IsNull() && !from.FailoverGroup.IsUnknown() {
		if toFailoverGroup, ok := to.GetFailoverGroup(ctx); ok {
			if fromFailoverGroup, ok := from.GetFailoverGroup(ctx); ok {
				toFailoverGroup.SyncFieldsDuringRead(ctx, fromFailoverGroup)
				to.SetFailoverGroup(ctx, toFailoverGroup)
			}
		}
	}
}

func (m CreateFailoverGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failover_group"] = attrs["failover_group"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["validate_only"] = attrs["validate_only"].SetOptional()
	attrs["failover_group_id"] = attrs["failover_group_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateFailoverGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateFailoverGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failover_group": reflect.TypeOf(FailoverGroup{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateFailoverGroupRequest
// only implements ToObjectValue() and Type().
func (m CreateFailoverGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failover_group":    m.FailoverGroup,
			"failover_group_id": m.FailoverGroupId,
			"parent":            m.Parent,
			"validate_only":     m.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateFailoverGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failover_group":    FailoverGroup{}.Type(ctx),
			"failover_group_id": types.StringType,
			"parent":            types.StringType,
			"validate_only":     types.BoolType,
		},
	}
}

// GetFailoverGroup returns the value of the FailoverGroup field in CreateFailoverGroupRequest as
// a FailoverGroup value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateFailoverGroupRequest) GetFailoverGroup(ctx context.Context) (FailoverGroup, bool) {
	var e FailoverGroup
	if m.FailoverGroup.IsNull() || m.FailoverGroup.IsUnknown() {
		return e, false
	}
	var v FailoverGroup
	d := m.FailoverGroup.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailoverGroup sets the value of the FailoverGroup field in CreateFailoverGroupRequest.
func (m *CreateFailoverGroupRequest) SetFailoverGroup(ctx context.Context, v FailoverGroup) {
	vs := v.ToObjectValue(ctx)
	m.FailoverGroup = vs
}

type CreateStableUrlRequest struct {
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"-"`
	// The stable URL to create.
	StableUrl types.Object `tfsdk:"stable_url"`
	// Client-provided identifier for the stable URL. Used to construct the
	// resource name as {parent}/stable-urls/{stable_url_id}.
	StableUrlId types.String `tfsdk:"-"`
	// When true, validates the request without creating the stable URL.
	ValidateOnly types.Bool `tfsdk:"-"`
}

func (to *CreateStableUrlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateStableUrlRequest) {
	if !from.StableUrl.IsNull() && !from.StableUrl.IsUnknown() {
		if toStableUrl, ok := to.GetStableUrl(ctx); ok {
			if fromStableUrl, ok := from.GetStableUrl(ctx); ok {
				// Recursively sync the fields of StableUrl
				toStableUrl.SyncFieldsDuringCreateOrUpdate(ctx, fromStableUrl)
				to.SetStableUrl(ctx, toStableUrl)
			}
		}
	}
}

func (to *CreateStableUrlRequest) SyncFieldsDuringRead(ctx context.Context, from CreateStableUrlRequest) {
	if !from.StableUrl.IsNull() && !from.StableUrl.IsUnknown() {
		if toStableUrl, ok := to.GetStableUrl(ctx); ok {
			if fromStableUrl, ok := from.GetStableUrl(ctx); ok {
				toStableUrl.SyncFieldsDuringRead(ctx, fromStableUrl)
				to.SetStableUrl(ctx, toStableUrl)
			}
		}
	}
}

func (m CreateStableUrlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["stable_url"] = attrs["stable_url"].SetRequired()
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["validate_only"] = attrs["validate_only"].SetOptional()
	attrs["stable_url_id"] = attrs["stable_url_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateStableUrlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateStableUrlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stable_url": reflect.TypeOf(StableUrl{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateStableUrlRequest
// only implements ToObjectValue() and Type().
func (m CreateStableUrlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"parent":        m.Parent,
			"stable_url":    m.StableUrl,
			"stable_url_id": m.StableUrlId,
			"validate_only": m.ValidateOnly,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateStableUrlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"parent":        types.StringType,
			"stable_url":    StableUrl{}.Type(ctx),
			"stable_url_id": types.StringType,
			"validate_only": types.BoolType,
		},
	}
}

// GetStableUrl returns the value of the StableUrl field in CreateStableUrlRequest as
// a StableUrl value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateStableUrlRequest) GetStableUrl(ctx context.Context) (StableUrl, bool) {
	var e StableUrl
	if m.StableUrl.IsNull() || m.StableUrl.IsUnknown() {
		return e, false
	}
	var v StableUrl
	d := m.StableUrl.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStableUrl sets the value of the StableUrl field in CreateStableUrlRequest.
func (m *CreateStableUrlRequest) SetStableUrl(ctx context.Context, v StableUrl) {
	vs := v.ToObjectValue(ctx)
	m.StableUrl = vs
}

type DeleteFailoverGroupRequest struct {
	// Opaque version string for optimistic locking. If provided, must match the
	// current etag. If omitted, the delete proceeds without an etag check.
	Etag types.String `tfsdk:"-"`
	// The fully qualified resource name of the failover group to delete.
	// Format: accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteFailoverGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteFailoverGroupRequest) {
}

func (to *DeleteFailoverGroupRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteFailoverGroupRequest) {
}

func (m DeleteFailoverGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteFailoverGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteFailoverGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteFailoverGroupRequest
// only implements ToObjectValue() and Type().
func (m DeleteFailoverGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteFailoverGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
		},
	}
}

type DeleteStableUrlRequest struct {
	// The fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name types.String `tfsdk:"-"`
}

func (to *DeleteStableUrlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteStableUrlRequest) {
}

func (to *DeleteStableUrlRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteStableUrlRequest) {
}

func (m DeleteStableUrlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteStableUrlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteStableUrlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteStableUrlRequest
// only implements ToObjectValue() and Type().
func (m DeleteStableUrlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteStableUrlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Request to failover a failover group to a new primary region.
type FailoverFailoverGroupRequest struct {
	// Opaque version string for optimistic locking. If provided, must match the
	// current etag. If omitted, the failover proceeds regardless of current
	// state.
	Etag types.String `tfsdk:"etag"`
	// The type of failover to perform.
	FailoverType types.String `tfsdk:"failover_type"`
	// The fully qualified resource name of the failover group to failover.
	// Format: accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"-"`
	// The target primary region. Must be one of the derived regions and
	// different from the current effective_primary_region. Serves as an
	// idempotency check.
	TargetPrimaryRegion types.String `tfsdk:"target_primary_region"`
}

func (to *FailoverFailoverGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FailoverFailoverGroupRequest) {
}

func (to *FailoverFailoverGroupRequest) SyncFieldsDuringRead(ctx context.Context, from FailoverFailoverGroupRequest) {
}

func (m FailoverFailoverGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["failover_type"] = attrs["failover_type"].SetRequired()
	attrs["target_primary_region"] = attrs["target_primary_region"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FailoverFailoverGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FailoverFailoverGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailoverFailoverGroupRequest
// only implements ToObjectValue() and Type().
func (m FailoverFailoverGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                  m.Etag,
			"failover_type":         m.FailoverType,
			"name":                  m.Name,
			"target_primary_region": m.TargetPrimaryRegion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FailoverFailoverGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":                  types.StringType,
			"failover_type":         types.StringType,
			"name":                  types.StringType,
			"target_primary_region": types.StringType,
		},
	}
}

// A failover group manages disaster recovery across workspace sets,
// coordinating UCDR and CPDR replication.
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
	// Initial primary region. Used only in Create requests to set the starting
	// primary region. Not returned in responses.
	InitialPrimaryRegion types.String `tfsdk:"initial_primary_region"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"name"`
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

func (to *FailoverGroup) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FailoverGroup) {
	if !from.InitialPrimaryRegion.IsUnknown() && !from.InitialPrimaryRegion.IsNull() {
		// InitialPrimaryRegion is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialPrimaryRegion = from.InitialPrimaryRegion
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

func (to *FailoverGroup) SyncFieldsDuringRead(ctx context.Context, from FailoverGroup) {
	if !from.InitialPrimaryRegion.IsUnknown() && !from.InitialPrimaryRegion.IsNull() {
		// InitialPrimaryRegion is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialPrimaryRegion = from.InitialPrimaryRegion
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
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["regions"] = attrs["regions"].SetRequired()
	attrs["regions"] = attrs["regions"].(tfschema.ListAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["replication_point"] = attrs["replication_point"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["unity_catalog_assets"] = attrs["unity_catalog_assets"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_sets"] = attrs["workspace_sets"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FailoverGroup.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FailoverGroup) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"regions":              reflect.TypeOf(types.String{}),
		"unity_catalog_assets": reflect.TypeOf(UcReplicationConfig{}),
		"workspace_sets":       reflect.TypeOf(WorkspaceSet{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailoverGroup
// only implements ToObjectValue() and Type().
func (m FailoverGroup) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":              m.CreateTime,
			"effective_primary_region": m.EffectivePrimaryRegion,
			"etag":                     m.Etag,
			"initial_primary_region":   m.InitialPrimaryRegion,
			"name":                     m.Name,
			"regions":                  m.Regions,
			"replication_point":        m.ReplicationPoint,
			"state":                    m.State,
			"unity_catalog_assets":     m.UnityCatalogAssets,
			"update_time":              m.UpdateTime,
			"workspace_sets":           m.WorkspaceSets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FailoverGroup) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":              timetypes.RFC3339{}.Type(ctx),
			"effective_primary_region": types.StringType,
			"etag":                     types.StringType,
			"initial_primary_region":   types.StringType,
			"name":                     types.StringType,
			"regions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"replication_point":    timetypes.RFC3339{}.Type(ctx),
			"state":                types.StringType,
			"unity_catalog_assets": UcReplicationConfig{}.Type(ctx),
			"update_time":          timetypes.RFC3339{}.Type(ctx),
			"workspace_sets": basetypes.ListType{
				ElemType: WorkspaceSet{}.Type(ctx),
			},
		},
	}
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
// a UcReplicationConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *FailoverGroup) GetUnityCatalogAssets(ctx context.Context) (UcReplicationConfig, bool) {
	var e UcReplicationConfig
	if m.UnityCatalogAssets.IsNull() || m.UnityCatalogAssets.IsUnknown() {
		return e, false
	}
	var v UcReplicationConfig
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
func (m *FailoverGroup) SetUnityCatalogAssets(ctx context.Context, v UcReplicationConfig) {
	vs := v.ToObjectValue(ctx)
	m.UnityCatalogAssets = vs
}

// GetWorkspaceSets returns the value of the WorkspaceSets field in FailoverGroup as
// a slice of WorkspaceSet values.
// If the field is unknown or null, the boolean return value is false.
func (m *FailoverGroup) GetWorkspaceSets(ctx context.Context) ([]WorkspaceSet, bool) {
	if m.WorkspaceSets.IsNull() || m.WorkspaceSets.IsUnknown() {
		return nil, false
	}
	var v []WorkspaceSet
	d := m.WorkspaceSets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceSets sets the value of the WorkspaceSets field in FailoverGroup.
func (m *FailoverGroup) SetWorkspaceSets(ctx context.Context, v []WorkspaceSet) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_sets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceSets = types.ListValueMust(t, vs)
}

type GetFailoverGroupRequest struct {
	// The fully qualified resource name of the failover group. Format:
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"-"`
}

func (to *GetFailoverGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetFailoverGroupRequest) {
}

func (to *GetFailoverGroupRequest) SyncFieldsDuringRead(ctx context.Context, from GetFailoverGroupRequest) {
}

func (m GetFailoverGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetFailoverGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetFailoverGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetFailoverGroupRequest
// only implements ToObjectValue() and Type().
func (m GetFailoverGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetFailoverGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetStableUrlRequest struct {
	// The fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name types.String `tfsdk:"-"`
}

func (to *GetStableUrlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStableUrlRequest) {
}

func (to *GetStableUrlRequest) SyncFieldsDuringRead(ctx context.Context, from GetStableUrlRequest) {
}

func (m GetStableUrlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStableUrlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetStableUrlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStableUrlRequest
// only implements ToObjectValue() and Type().
func (m GetStableUrlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStableUrlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListFailoverGroupsRequest struct {
	// Maximum number of failover groups to return per page. Default: 50,
	// maximum: 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token received from a previous ListFailoverGroups call. Provide this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"-"`
}

func (to *ListFailoverGroupsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFailoverGroupsRequest) {
}

func (to *ListFailoverGroupsRequest) SyncFieldsDuringRead(ctx context.Context, from ListFailoverGroupsRequest) {
}

func (m ListFailoverGroupsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFailoverGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFailoverGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFailoverGroupsRequest
// only implements ToObjectValue() and Type().
func (m ListFailoverGroupsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFailoverGroupsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for listing failover groups.
type ListFailoverGroupsResponse struct {
	// The failover groups for this account.
	FailoverGroups types.List `tfsdk:"failover_groups"`
	// A token that can be sent as page_token to retrieve the next page. If
	// omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListFailoverGroupsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFailoverGroupsResponse) {
	if !from.FailoverGroups.IsNull() && !from.FailoverGroups.IsUnknown() && to.FailoverGroups.IsNull() && len(from.FailoverGroups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailoverGroups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailoverGroups = from.FailoverGroups
	}
}

func (to *ListFailoverGroupsResponse) SyncFieldsDuringRead(ctx context.Context, from ListFailoverGroupsResponse) {
	if !from.FailoverGroups.IsNull() && !from.FailoverGroups.IsUnknown() && to.FailoverGroups.IsNull() && len(from.FailoverGroups.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for FailoverGroups, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.FailoverGroups = from.FailoverGroups
	}
}

func (m ListFailoverGroupsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failover_groups"] = attrs["failover_groups"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFailoverGroupsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListFailoverGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failover_groups": reflect.TypeOf(FailoverGroup{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFailoverGroupsResponse
// only implements ToObjectValue() and Type().
func (m ListFailoverGroupsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failover_groups": m.FailoverGroups,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFailoverGroupsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failover_groups": basetypes.ListType{
				ElemType: FailoverGroup{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetFailoverGroups returns the value of the FailoverGroups field in ListFailoverGroupsResponse as
// a slice of FailoverGroup values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFailoverGroupsResponse) GetFailoverGroups(ctx context.Context) ([]FailoverGroup, bool) {
	if m.FailoverGroups.IsNull() || m.FailoverGroups.IsUnknown() {
		return nil, false
	}
	var v []FailoverGroup
	d := m.FailoverGroups.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailoverGroups sets the value of the FailoverGroups field in ListFailoverGroupsResponse.
func (m *ListFailoverGroupsResponse) SetFailoverGroups(ctx context.Context, v []FailoverGroup) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failover_groups"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.FailoverGroups = types.ListValueMust(t, vs)
}

type ListStableUrlsRequest struct {
	// Maximum number of stable URLs to return per page. Default: 50, maximum:
	// 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Page token received from a previous ListStableUrls call. Provide this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The parent resource. Format: accounts/{account_id}.
	Parent types.String `tfsdk:"-"`
}

func (to *ListStableUrlsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListStableUrlsRequest) {
}

func (to *ListStableUrlsRequest) SyncFieldsDuringRead(ctx context.Context, from ListStableUrlsRequest) {
}

func (m ListStableUrlsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["parent"] = attrs["parent"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStableUrlsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListStableUrlsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStableUrlsRequest
// only implements ToObjectValue() and Type().
func (m ListStableUrlsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"parent":     m.Parent,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListStableUrlsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"parent":     types.StringType,
		},
	}
}

// Response for listing stable URLs.
type ListStableUrlsResponse struct {
	// A token that can be sent as page_token to retrieve the next page. If
	// omitted, there are no subsequent pages.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// The stable URLs for this account.
	StableUrls types.List `tfsdk:"stable_urls"`
}

func (to *ListStableUrlsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListStableUrlsResponse) {
	if !from.StableUrls.IsNull() && !from.StableUrls.IsUnknown() && to.StableUrls.IsNull() && len(from.StableUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StableUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StableUrls = from.StableUrls
	}
}

func (to *ListStableUrlsResponse) SyncFieldsDuringRead(ctx context.Context, from ListStableUrlsResponse) {
	if !from.StableUrls.IsNull() && !from.StableUrls.IsUnknown() && to.StableUrls.IsNull() && len(from.StableUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StableUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StableUrls = from.StableUrls
	}
}

func (m ListStableUrlsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["stable_urls"] = attrs["stable_urls"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListStableUrlsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListStableUrlsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stable_urls": reflect.TypeOf(StableUrl{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListStableUrlsResponse
// only implements ToObjectValue() and Type().
func (m ListStableUrlsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"stable_urls":     m.StableUrls,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListStableUrlsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"stable_urls": basetypes.ListType{
				ElemType: StableUrl{}.Type(ctx),
			},
		},
	}
}

// GetStableUrls returns the value of the StableUrls field in ListStableUrlsResponse as
// a slice of StableUrl values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListStableUrlsResponse) GetStableUrls(ctx context.Context) ([]StableUrl, bool) {
	if m.StableUrls.IsNull() || m.StableUrls.IsUnknown() {
		return nil, false
	}
	var v []StableUrl
	d := m.StableUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStableUrls sets the value of the StableUrls field in ListStableUrlsResponse.
func (m *ListStableUrlsResponse) SetStableUrls(ctx context.Context, v []StableUrl) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["stable_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.StableUrls = types.ListValueMust(t, vs)
}

// A location mapping identified by a name, with URIs per region. The system
// derives replication direction from effective_primary_region.
type LocationMapping struct {
	// Resource name for this location.
	Name types.String `tfsdk:"name"`
	// URI for each region. Each entry maps a region name to a storage URI.
	UriByRegion types.List `tfsdk:"uri_by_region"`
}

func (to *LocationMapping) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LocationMapping) {
}

func (to *LocationMapping) SyncFieldsDuringRead(ctx context.Context, from LocationMapping) {
}

func (m LocationMapping) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["uri_by_region"] = attrs["uri_by_region"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LocationMapping.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LocationMapping) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"uri_by_region": reflect.TypeOf(LocationMappingEntry{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LocationMapping
// only implements ToObjectValue() and Type().
func (m LocationMapping) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":          m.Name,
			"uri_by_region": m.UriByRegion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LocationMapping) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"uri_by_region": basetypes.ListType{
				ElemType: LocationMappingEntry{}.Type(ctx),
			},
		},
	}
}

// GetUriByRegion returns the value of the UriByRegion field in LocationMapping as
// a slice of LocationMappingEntry values.
// If the field is unknown or null, the boolean return value is false.
func (m *LocationMapping) GetUriByRegion(ctx context.Context) ([]LocationMappingEntry, bool) {
	if m.UriByRegion.IsNull() || m.UriByRegion.IsUnknown() {
		return nil, false
	}
	var v []LocationMappingEntry
	d := m.UriByRegion.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUriByRegion sets the value of the UriByRegion field in LocationMapping.
func (m *LocationMapping) SetUriByRegion(ctx context.Context, v []LocationMappingEntry) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["uri_by_region"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UriByRegion = types.ListValueMust(t, vs)
}

// A single entry in a location mapping, mapping a region to a storage URI. Used
// instead of map<string, string> for proto2 compatibility.
type LocationMappingEntry struct {
	// The region name.
	Region types.String `tfsdk:"region"`
	// The storage URI for this region.
	Uri types.String `tfsdk:"uri"`
}

func (to *LocationMappingEntry) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LocationMappingEntry) {
}

func (to *LocationMappingEntry) SyncFieldsDuringRead(ctx context.Context, from LocationMappingEntry) {
}

func (m LocationMappingEntry) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["region"] = attrs["region"].SetRequired()
	attrs["uri"] = attrs["uri"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LocationMappingEntry.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LocationMappingEntry) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LocationMappingEntry
// only implements ToObjectValue() and Type().
func (m LocationMappingEntry) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"region": m.Region,
			"uri":    m.Uri,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LocationMappingEntry) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"region": types.StringType,
			"uri":    types.StringType,
		},
	}
}

// A stable URL provides a failover-aware endpoint for accessing a workspace.
// Its lifecycle is independent of any failover group.
type StableUrl struct {
	// The workspace this stable URL is initially bound to. Used only in Create
	// requests to associate the stable URL with a workspace. Not returned in
	// responses. Mirrors FailoverGroup.initial_primary_region semantics.
	InitialWorkspaceId types.String `tfsdk:"initial_workspace_id"`
	// Fully qualified resource name. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}.
	Name types.String `tfsdk:"name"`
	// The stable URL endpoint. Generated by the backend on creation and
	// immutable thereafter. For non-Private-Link workspaces this is
	// `https://<spog_host>/?c=<connection_id>`. For Private-Link workspaces
	// this is the per-connection hostname.
	Url types.String `tfsdk:"url"`
}

func (to *StableUrl) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StableUrl) {
	if !from.InitialWorkspaceId.IsUnknown() && !from.InitialWorkspaceId.IsNull() {
		// InitialWorkspaceId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialWorkspaceId = from.InitialWorkspaceId
	}
}

func (to *StableUrl) SyncFieldsDuringRead(ctx context.Context, from StableUrl) {
	if !from.InitialWorkspaceId.IsUnknown() && !from.InitialWorkspaceId.IsNull() {
		// InitialWorkspaceId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InitialWorkspaceId = from.InitialWorkspaceId
	}
}

func (m StableUrl) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_workspace_id"] = attrs["initial_workspace_id"].SetRequired()
	attrs["initial_workspace_id"] = attrs["initial_workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["url"] = attrs["url"].SetComputed()
	attrs["url"] = attrs["url"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StableUrl.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StableUrl) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StableUrl
// only implements ToObjectValue() and Type().
func (m StableUrl) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_workspace_id": m.InitialWorkspaceId,
			"name":                 m.Name,
			"url":                  m.Url,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StableUrl) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_workspace_id": types.StringType,
			"name":                 types.StringType,
			"url":                  types.StringType,
		},
	}
}

// A Unity Catalog catalog to replicate.
type UcCatalog struct {
	// The name of the UC catalog to replicate.
	Name types.String `tfsdk:"name"`
}

func (to *UcCatalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UcCatalog) {
}

func (to *UcCatalog) SyncFieldsDuringRead(ctx context.Context, from UcCatalog) {
}

func (m UcCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UcCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UcCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UcCatalog
// only implements ToObjectValue() and Type().
func (m UcCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UcCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Unity Catalog replication configuration (top-level, not per-set).
type UcReplicationConfig struct {
	// UC catalogs to replicate.
	Catalogs types.List `tfsdk:"catalogs"`
	// The workspace set whose workspaces will be used for data replication of
	// all UC catalogs' underlying storage.
	DataReplicationWorkspaceSet types.String `tfsdk:"data_replication_workspace_set"`
	// Location mappings - storage URI per region for each location.
	LocationMappings types.List `tfsdk:"location_mappings"`
}

func (to *UcReplicationConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UcReplicationConfig) {
	if !from.LocationMappings.IsNull() && !from.LocationMappings.IsUnknown() && to.LocationMappings.IsNull() && len(from.LocationMappings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LocationMappings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LocationMappings = from.LocationMappings
	}
}

func (to *UcReplicationConfig) SyncFieldsDuringRead(ctx context.Context, from UcReplicationConfig) {
	if !from.LocationMappings.IsNull() && !from.LocationMappings.IsUnknown() && to.LocationMappings.IsNull() && len(from.LocationMappings.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for LocationMappings, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.LocationMappings = from.LocationMappings
	}
}

func (m UcReplicationConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalogs"] = attrs["catalogs"].SetRequired()
	attrs["catalogs"] = attrs["catalogs"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["data_replication_workspace_set"] = attrs["data_replication_workspace_set"].SetRequired()
	attrs["location_mappings"] = attrs["location_mappings"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UcReplicationConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UcReplicationConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalogs":          reflect.TypeOf(UcCatalog{}),
		"location_mappings": reflect.TypeOf(LocationMapping{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UcReplicationConfig
// only implements ToObjectValue() and Type().
func (m UcReplicationConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalogs":                       m.Catalogs,
			"data_replication_workspace_set": m.DataReplicationWorkspaceSet,
			"location_mappings":              m.LocationMappings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UcReplicationConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalogs": basetypes.ListType{
				ElemType: UcCatalog{}.Type(ctx),
			},
			"data_replication_workspace_set": types.StringType,
			"location_mappings": basetypes.ListType{
				ElemType: LocationMapping{}.Type(ctx),
			},
		},
	}
}

// GetCatalogs returns the value of the Catalogs field in UcReplicationConfig as
// a slice of UcCatalog values.
// If the field is unknown or null, the boolean return value is false.
func (m *UcReplicationConfig) GetCatalogs(ctx context.Context) ([]UcCatalog, bool) {
	if m.Catalogs.IsNull() || m.Catalogs.IsUnknown() {
		return nil, false
	}
	var v []UcCatalog
	d := m.Catalogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalogs sets the value of the Catalogs field in UcReplicationConfig.
func (m *UcReplicationConfig) SetCatalogs(ctx context.Context, v []UcCatalog) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["catalogs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Catalogs = types.ListValueMust(t, vs)
}

// GetLocationMappings returns the value of the LocationMappings field in UcReplicationConfig as
// a slice of LocationMapping values.
// If the field is unknown or null, the boolean return value is false.
func (m *UcReplicationConfig) GetLocationMappings(ctx context.Context) ([]LocationMapping, bool) {
	if m.LocationMappings.IsNull() || m.LocationMappings.IsUnknown() {
		return nil, false
	}
	var v []LocationMapping
	d := m.LocationMappings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLocationMappings sets the value of the LocationMappings field in UcReplicationConfig.
func (m *UcReplicationConfig) SetLocationMappings(ctx context.Context, v []LocationMapping) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["location_mappings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.LocationMappings = types.ListValueMust(t, vs)
}

type UpdateFailoverGroupRequest struct {
	// The failover group with updated fields. The name field identifies the
	// resource and is populated from the URL path.
	FailoverGroup types.Object `tfsdk:"failover_group"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"-"`
	// Comma-separated list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateFailoverGroupRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateFailoverGroupRequest) {
	if !from.FailoverGroup.IsNull() && !from.FailoverGroup.IsUnknown() {
		if toFailoverGroup, ok := to.GetFailoverGroup(ctx); ok {
			if fromFailoverGroup, ok := from.GetFailoverGroup(ctx); ok {
				// Recursively sync the fields of FailoverGroup
				toFailoverGroup.SyncFieldsDuringCreateOrUpdate(ctx, fromFailoverGroup)
				to.SetFailoverGroup(ctx, toFailoverGroup)
			}
		}
	}
}

func (to *UpdateFailoverGroupRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateFailoverGroupRequest) {
	if !from.FailoverGroup.IsNull() && !from.FailoverGroup.IsUnknown() {
		if toFailoverGroup, ok := to.GetFailoverGroup(ctx); ok {
			if fromFailoverGroup, ok := from.GetFailoverGroup(ctx); ok {
				toFailoverGroup.SyncFieldsDuringRead(ctx, fromFailoverGroup)
				to.SetFailoverGroup(ctx, toFailoverGroup)
			}
		}
	}
}

func (m UpdateFailoverGroupRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failover_group"] = attrs["failover_group"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateFailoverGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateFailoverGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"failover_group": reflect.TypeOf(FailoverGroup{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateFailoverGroupRequest
// only implements ToObjectValue() and Type().
func (m UpdateFailoverGroupRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failover_group": m.FailoverGroup,
			"name":           m.Name,
			"update_mask":    m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateFailoverGroupRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failover_group": FailoverGroup{}.Type(ctx),
			"name":           types.StringType,
			"update_mask":    types.StringType,
		},
	}
}

// GetFailoverGroup returns the value of the FailoverGroup field in UpdateFailoverGroupRequest as
// a FailoverGroup value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateFailoverGroupRequest) GetFailoverGroup(ctx context.Context) (FailoverGroup, bool) {
	var e FailoverGroup
	if m.FailoverGroup.IsNull() || m.FailoverGroup.IsUnknown() {
		return e, false
	}
	var v FailoverGroup
	d := m.FailoverGroup.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailoverGroup sets the value of the FailoverGroup field in UpdateFailoverGroupRequest.
func (m *UpdateFailoverGroupRequest) SetFailoverGroup(ctx context.Context, v FailoverGroup) {
	vs := v.ToObjectValue(ctx)
	m.FailoverGroup = vs
}

// A set of workspaces that replicate to each other across regions.
type WorkspaceSet struct {
	// Resource name for this workspace set.
	Name types.String `tfsdk:"name"`
	// Whether to enable control plane DR (notebooks, jobs, clusters, etc.) for
	// this set. Requires all workspaces in the set to be Mission Critical tier.
	ReplicateWorkspaceAssets types.Bool `tfsdk:"replicate_workspace_assets"`
	// Resource names of stable URLs associated with this workspace set. Format:
	// accounts/{account_id}/stable-urls/{stable_url_id}. The referenced stable
	// URLs must already exist (via CreateStableUrl).
	StableUrlNames types.List `tfsdk:"stable_url_names"`
	// Workspace IDs in this set. The system derives and validates regions. EA:
	// exactly 2 workspaces (one per region).
	WorkspaceIds types.List `tfsdk:"workspace_ids"`
}

func (to *WorkspaceSet) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceSet) {
	if !from.StableUrlNames.IsNull() && !from.StableUrlNames.IsUnknown() && to.StableUrlNames.IsNull() && len(from.StableUrlNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StableUrlNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StableUrlNames = from.StableUrlNames
	}
}

func (to *WorkspaceSet) SyncFieldsDuringRead(ctx context.Context, from WorkspaceSet) {
	if !from.StableUrlNames.IsNull() && !from.StableUrlNames.IsUnknown() && to.StableUrlNames.IsNull() && len(from.StableUrlNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for StableUrlNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.StableUrlNames = from.StableUrlNames
	}
}

func (m WorkspaceSet) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["replicate_workspace_assets"] = attrs["replicate_workspace_assets"].SetRequired()
	attrs["stable_url_names"] = attrs["stable_url_names"].SetOptional()
	attrs["workspace_ids"] = attrs["workspace_ids"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceSet.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceSet) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"stable_url_names": reflect.TypeOf(types.String{}),
		"workspace_ids":    reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceSet
// only implements ToObjectValue() and Type().
func (m WorkspaceSet) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":                       m.Name,
			"replicate_workspace_assets": m.ReplicateWorkspaceAssets,
			"stable_url_names":           m.StableUrlNames,
			"workspace_ids":              m.WorkspaceIds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceSet) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":                       types.StringType,
			"replicate_workspace_assets": types.BoolType,
			"stable_url_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"workspace_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetStableUrlNames returns the value of the StableUrlNames field in WorkspaceSet as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceSet) GetStableUrlNames(ctx context.Context) ([]types.String, bool) {
	if m.StableUrlNames.IsNull() || m.StableUrlNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.StableUrlNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStableUrlNames sets the value of the StableUrlNames field in WorkspaceSet.
func (m *WorkspaceSet) SetStableUrlNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["stable_url_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.StableUrlNames = types.ListValueMust(t, vs)
}

// GetWorkspaceIds returns the value of the WorkspaceIds field in WorkspaceSet as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *WorkspaceSet) GetWorkspaceIds(ctx context.Context) ([]types.String, bool) {
	if m.WorkspaceIds.IsNull() || m.WorkspaceIds.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.WorkspaceIds.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceIds sets the value of the WorkspaceIds field in WorkspaceSet.
func (m *WorkspaceSet) SetWorkspaceIds(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_ids"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.WorkspaceIds = types.ListValueMust(t, vs)
}
