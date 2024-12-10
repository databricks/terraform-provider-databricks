// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package cleanrooms_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/databricks/terraform-provider-databricks/internal/service/jobs_tf"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	AccessRestricted types.String `tfsdk:"access_restricted" tf:"optional"`

	Comment types.String `tfsdk:"comment" tf:"optional"`
	// When the clean room was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The alias of the collaborator tied to the local clean room.
	LocalCollaboratorAlias types.String `tfsdk:"local_collaborator_alias" tf:"optional"`
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	Name types.String `tfsdk:"name" tf:"optional"`
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	OutputCatalog types.List `tfsdk:"output_catalog" tf:"optional,object"`
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner types.String `tfsdk:"owner" tf:"optional"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo types.List `tfsdk:"remote_detailed_info" tf:"optional,object"`
	// Clean room status.
	Status types.String `tfsdk:"status" tf:"optional"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
}

func (newState *CleanRoom) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoom) {
}

func (newState *CleanRoom) SyncEffectiveFieldsDuringRead(existingState CleanRoom) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoom.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoom) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog":       reflect.TypeOf(CleanRoomOutputCatalog{}),
		"remote_detailed_info": reflect.TypeOf(CleanRoomRemoteDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoom
// only implements ToObjectValue() and Type().
func (o CleanRoom) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_restricted":        o.AccessRestricted,
			"comment":                  o.Comment,
			"created_at":               o.CreatedAt,
			"local_collaborator_alias": o.LocalCollaboratorAlias,
			"name":                     o.Name,
			"output_catalog":           o.OutputCatalog,
			"owner":                    o.Owner,
			"remote_detailed_info":     o.RemoteDetailedInfo,
			"status":                   o.Status,
			"updated_at":               o.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoom) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_restricted":        types.StringType,
			"comment":                  types.StringType,
			"created_at":               types.Int64Type,
			"local_collaborator_alias": types.StringType,
			"name":                     types.StringType,
			"output_catalog": basetypes.ListType{
				ElemType: CleanRoomOutputCatalog{}.Type(ctx),
			},
			"owner": types.StringType,
			"remote_detailed_info": basetypes.ListType{
				ElemType: CleanRoomRemoteDetail{}.Type(ctx),
			},
			"status":     types.StringType,
			"updated_at": types.Int64Type,
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CleanRoom as
// a CleanRoomOutputCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoom) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog, bool) {
	var e CleanRoomOutputCatalog
	if o.OutputCatalog.IsNull() || o.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog
	d := o.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CleanRoom.
func (o *CleanRoom) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	o.OutputCatalog = types.ListValueMust(t, vs)
}

// GetRemoteDetailedInfo returns the value of the RemoteDetailedInfo field in CleanRoom as
// a CleanRoomRemoteDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoom) GetRemoteDetailedInfo(ctx context.Context) (CleanRoomRemoteDetail, bool) {
	var e CleanRoomRemoteDetail
	if o.RemoteDetailedInfo.IsNull() || o.RemoteDetailedInfo.IsUnknown() {
		return e, false
	}
	var v []CleanRoomRemoteDetail
	d := o.RemoteDetailedInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRemoteDetailedInfo sets the value of the RemoteDetailedInfo field in CleanRoom.
func (o *CleanRoom) SetRemoteDetailedInfo(ctx context.Context, v CleanRoomRemoteDetail) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["remote_detailed_info"]
	o.RemoteDetailedInfo = types.ListValueMust(t, vs)
}

// Metadata of the clean room asset
type CleanRoomAsset struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at" tf:"optional"`
	// The type of the asset.
	AssetType types.String `tfsdk:"asset_type" tf:"optional"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable types.List `tfsdk:"foreign_table" tf:"optional,object"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails types.List `tfsdk:"foreign_table_local_details" tf:"optional,object"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook types.List `tfsdk:"notebook" tf:"optional,object"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias types.String `tfsdk:"owner_collaborator_alias" tf:"optional"`
	// Status of the asset
	Status types.String `tfsdk:"status" tf:"optional"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table types.List `tfsdk:"table" tf:"optional,object"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails types.List `tfsdk:"table_local_details" tf:"optional,object"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View types.List `tfsdk:"view" tf:"optional,object"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails types.List `tfsdk:"view_local_details" tf:"optional,object"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails types.List `tfsdk:"volume_local_details" tf:"optional,object"`
}

func (newState *CleanRoomAsset) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAsset) {
}

func (newState *CleanRoomAsset) SyncEffectiveFieldsDuringRead(existingState CleanRoomAsset) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAsset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAsset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"foreign_table":               reflect.TypeOf(CleanRoomAssetForeignTable{}),
		"foreign_table_local_details": reflect.TypeOf(CleanRoomAssetForeignTableLocalDetails{}),
		"notebook":                    reflect.TypeOf(CleanRoomAssetNotebook{}),
		"table":                       reflect.TypeOf(CleanRoomAssetTable{}),
		"table_local_details":         reflect.TypeOf(CleanRoomAssetTableLocalDetails{}),
		"view":                        reflect.TypeOf(CleanRoomAssetView{}),
		"view_local_details":          reflect.TypeOf(CleanRoomAssetViewLocalDetails{}),
		"volume_local_details":        reflect.TypeOf(CleanRoomAssetVolumeLocalDetails{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAsset
// only implements ToObjectValue() and Type().
func (o CleanRoomAsset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"added_at":                    o.AddedAt,
			"asset_type":                  o.AssetType,
			"foreign_table":               o.ForeignTable,
			"foreign_table_local_details": o.ForeignTableLocalDetails,
			"name":                        o.Name,
			"notebook":                    o.Notebook,
			"owner_collaborator_alias":    o.OwnerCollaboratorAlias,
			"status":                      o.Status,
			"table":                       o.Table,
			"table_local_details":         o.TableLocalDetails,
			"view":                        o.View,
			"view_local_details":          o.ViewLocalDetails,
			"volume_local_details":        o.VolumeLocalDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAsset) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"added_at":   types.Int64Type,
			"asset_type": types.StringType,
			"foreign_table": basetypes.ListType{
				ElemType: CleanRoomAssetForeignTable{}.Type(ctx),
			},
			"foreign_table_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetForeignTableLocalDetails{}.Type(ctx),
			},
			"name": types.StringType,
			"notebook": basetypes.ListType{
				ElemType: CleanRoomAssetNotebook{}.Type(ctx),
			},
			"owner_collaborator_alias": types.StringType,
			"status":                   types.StringType,
			"table": basetypes.ListType{
				ElemType: CleanRoomAssetTable{}.Type(ctx),
			},
			"table_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetTableLocalDetails{}.Type(ctx),
			},
			"view": basetypes.ListType{
				ElemType: CleanRoomAssetView{}.Type(ctx),
			},
			"view_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetViewLocalDetails{}.Type(ctx),
			},
			"volume_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetVolumeLocalDetails{}.Type(ctx),
			},
		},
	}
}

// GetForeignTable returns the value of the ForeignTable field in CleanRoomAsset as
// a CleanRoomAssetForeignTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetForeignTable(ctx context.Context) (CleanRoomAssetForeignTable, bool) {
	var e CleanRoomAssetForeignTable
	if o.ForeignTable.IsNull() || o.ForeignTable.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetForeignTable
	d := o.ForeignTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignTable sets the value of the ForeignTable field in CleanRoomAsset.
func (o *CleanRoomAsset) SetForeignTable(ctx context.Context, v CleanRoomAssetForeignTable) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_table"]
	o.ForeignTable = types.ListValueMust(t, vs)
}

// GetForeignTableLocalDetails returns the value of the ForeignTableLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetForeignTableLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetForeignTableLocalDetails(ctx context.Context) (CleanRoomAssetForeignTableLocalDetails, bool) {
	var e CleanRoomAssetForeignTableLocalDetails
	if o.ForeignTableLocalDetails.IsNull() || o.ForeignTableLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetForeignTableLocalDetails
	d := o.ForeignTableLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignTableLocalDetails sets the value of the ForeignTableLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetForeignTableLocalDetails(ctx context.Context, v CleanRoomAssetForeignTableLocalDetails) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_table_local_details"]
	o.ForeignTableLocalDetails = types.ListValueMust(t, vs)
}

// GetNotebook returns the value of the Notebook field in CleanRoomAsset as
// a CleanRoomAssetNotebook value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetNotebook(ctx context.Context) (CleanRoomAssetNotebook, bool) {
	var e CleanRoomAssetNotebook
	if o.Notebook.IsNull() || o.Notebook.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetNotebook
	d := o.Notebook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebook sets the value of the Notebook field in CleanRoomAsset.
func (o *CleanRoomAsset) SetNotebook(ctx context.Context, v CleanRoomAssetNotebook) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook"]
	o.Notebook = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in CleanRoomAsset as
// a CleanRoomAssetTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetTable(ctx context.Context) (CleanRoomAssetTable, bool) {
	var e CleanRoomAssetTable
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetTable
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CleanRoomAsset.
func (o *CleanRoomAsset) SetTable(ctx context.Context, v CleanRoomAssetTable) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

// GetTableLocalDetails returns the value of the TableLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetTableLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetTableLocalDetails(ctx context.Context) (CleanRoomAssetTableLocalDetails, bool) {
	var e CleanRoomAssetTableLocalDetails
	if o.TableLocalDetails.IsNull() || o.TableLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetTableLocalDetails
	d := o.TableLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableLocalDetails sets the value of the TableLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetTableLocalDetails(ctx context.Context, v CleanRoomAssetTableLocalDetails) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_local_details"]
	o.TableLocalDetails = types.ListValueMust(t, vs)
}

// GetView returns the value of the View field in CleanRoomAsset as
// a CleanRoomAssetView value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetView(ctx context.Context) (CleanRoomAssetView, bool) {
	var e CleanRoomAssetView
	if o.View.IsNull() || o.View.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetView
	d := o.View.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetView sets the value of the View field in CleanRoomAsset.
func (o *CleanRoomAsset) SetView(ctx context.Context, v CleanRoomAssetView) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["view"]
	o.View = types.ListValueMust(t, vs)
}

// GetViewLocalDetails returns the value of the ViewLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetViewLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetViewLocalDetails(ctx context.Context) (CleanRoomAssetViewLocalDetails, bool) {
	var e CleanRoomAssetViewLocalDetails
	if o.ViewLocalDetails.IsNull() || o.ViewLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetViewLocalDetails
	d := o.ViewLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetViewLocalDetails sets the value of the ViewLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetViewLocalDetails(ctx context.Context, v CleanRoomAssetViewLocalDetails) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["view_local_details"]
	o.ViewLocalDetails = types.ListValueMust(t, vs)
}

// GetVolumeLocalDetails returns the value of the VolumeLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetVolumeLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetVolumeLocalDetails(ctx context.Context) (CleanRoomAssetVolumeLocalDetails, bool) {
	var e CleanRoomAssetVolumeLocalDetails
	if o.VolumeLocalDetails.IsNull() || o.VolumeLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetVolumeLocalDetails
	d := o.VolumeLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumeLocalDetails sets the value of the VolumeLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetVolumeLocalDetails(ctx context.Context, v CleanRoomAssetVolumeLocalDetails) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volume_local_details"]
	o.VolumeLocalDetails = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	Columns types.List `tfsdk:"columns" tf:"optional"`
}

func (newState *CleanRoomAssetForeignTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetForeignTable) {
}

func (newState *CleanRoomAssetForeignTable) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetForeignTable) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetForeignTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetForeignTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTable
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetForeignTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetForeignTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: catalog_tf.ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in CleanRoomAssetForeignTable as
// a slice of catalog_tf.ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetForeignTable) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetForeignTable.
func (o *CleanRoomAssetForeignTable) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
}

func (newState *CleanRoomAssetForeignTableLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetForeignTableLocalDetails) {
}

func (newState *CleanRoomAssetForeignTableLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetForeignTableLocalDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetForeignTableLocalDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetForeignTableLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTableLocalDetails
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetForeignTableLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetForeignTableLocalDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAssetNotebook struct {
	// Server generated checksum that represents the notebook version.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent types.String `tfsdk:"notebook_content" tf:"optional"`
}

func (newState *CleanRoomAssetNotebook) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetNotebook) {
}

func (newState *CleanRoomAssetNotebook) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetNotebook) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetNotebook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetNotebook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetNotebook
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetNotebook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             o.Etag,
			"notebook_content": o.NotebookContent,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetNotebook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":             types.StringType,
			"notebook_content": types.StringType,
		},
	}
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	Columns types.List `tfsdk:"columns" tf:"optional"`
}

func (newState *CleanRoomAssetTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetTable) {
}

func (newState *CleanRoomAssetTable) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetTable) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTable
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: catalog_tf.ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in CleanRoomAssetTable as
// a slice of catalog_tf.ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetTable) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetTable.
func (o *CleanRoomAssetTable) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
	// Partition filtering specification for a shared table.
	Partitions types.List `tfsdk:"partitions" tf:"optional"`
}

func (newState *CleanRoomAssetTableLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetTableLocalDetails) {
}

func (newState *CleanRoomAssetTableLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetTableLocalDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetTableLocalDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetTableLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitions": reflect.TypeOf(sharing_tf.PartitionSpecificationPartition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTableLocalDetails
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetTableLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
			"partitions": o.Partitions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetTableLocalDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
			"partitions": basetypes.ListType{
				ElemType: sharing_tf.PartitionSpecificationPartition{}.Type(ctx),
			},
		},
	}
}

// GetPartitions returns the value of the Partitions field in CleanRoomAssetTableLocalDetails as
// a slice of sharing_tf.PartitionSpecificationPartition values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetTableLocalDetails) GetPartitions(ctx context.Context) ([]sharing_tf.PartitionSpecificationPartition, bool) {
	if o.Partitions.IsNull() || o.Partitions.IsUnknown() {
		return nil, false
	}
	var v []sharing_tf.PartitionSpecificationPartition
	d := o.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in CleanRoomAssetTableLocalDetails.
func (o *CleanRoomAssetTableLocalDetails) SetPartitions(ctx context.Context, v []sharing_tf.PartitionSpecificationPartition) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["partitions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Partitions = types.ListValueMust(t, vs)
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	Columns types.List `tfsdk:"columns" tf:"optional"`
}

func (newState *CleanRoomAssetView) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetView) {
}

func (newState *CleanRoomAssetView) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetView) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetView.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetView) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetView
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetView) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetView) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: catalog_tf.ColumnInfo{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in CleanRoomAssetView as
// a slice of catalog_tf.ColumnInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetView) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetView.
func (o *CleanRoomAssetView) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
}

func (newState *CleanRoomAssetViewLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetViewLocalDetails) {
}

func (newState *CleanRoomAssetViewLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetViewLocalDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetViewLocalDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetViewLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetViewLocalDetails
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetViewLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetViewLocalDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAssetVolumeLocalDetails struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	LocalName types.String `tfsdk:"local_name" tf:"optional"`
}

func (newState *CleanRoomAssetVolumeLocalDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomAssetVolumeLocalDetails) {
}

func (newState *CleanRoomAssetVolumeLocalDetails) SyncEffectiveFieldsDuringRead(existingState CleanRoomAssetVolumeLocalDetails) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetVolumeLocalDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetVolumeLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetVolumeLocalDetails
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetVolumeLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetVolumeLocalDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

// Publicly visible clean room collaborator.
type CleanRoomCollaborator struct {
	// Collaborator alias specified by the clean room creator. It is unique
	// across all collaborators of this clean room, and used to derive multiple
	// values internally such as catalog alias and clean room name for single
	// metastore clean rooms. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CollaboratorAlias types.String `tfsdk:"collaborator_alias" tf:"optional"`
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id" tf:"optional"`
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	InviteRecipientEmail types.String `tfsdk:"invite_recipient_email" tf:"optional"`
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	InviteRecipientWorkspaceId types.Int64 `tfsdk:"invite_recipient_workspace_id" tf:"optional"`
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	OrganizationName types.String `tfsdk:"organization_name" tf:"optional"`
}

func (newState *CleanRoomCollaborator) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomCollaborator) {
}

func (newState *CleanRoomCollaborator) SyncEffectiveFieldsDuringRead(existingState CleanRoomCollaborator) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomCollaborator.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomCollaborator) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomCollaborator
// only implements ToObjectValue() and Type().
func (o CleanRoomCollaborator) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_alias":            o.CollaboratorAlias,
			"display_name":                  o.DisplayName,
			"global_metastore_id":           o.GlobalMetastoreId,
			"invite_recipient_email":        o.InviteRecipientEmail,
			"invite_recipient_workspace_id": o.InviteRecipientWorkspaceId,
			"organization_name":             o.OrganizationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomCollaborator) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"collaborator_alias":            types.StringType,
			"display_name":                  types.StringType,
			"global_metastore_id":           types.StringType,
			"invite_recipient_email":        types.StringType,
			"invite_recipient_workspace_id": types.Int64Type,
			"organization_name":             types.StringType,
		},
	}
}

// Stores information about a single task run.
type CleanRoomNotebookTaskRun struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo types.List `tfsdk:"collaborator_job_run_info" tf:"optional,object"`
	// State of the task run.
	NotebookJobRunState types.List `tfsdk:"notebook_job_run_state" tf:"optional,object"`
	// Asset name of the notebook executed in this task run.
	NotebookName types.String `tfsdk:"notebook_name" tf:"optional"`
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	OutputSchemaExpirationTime types.Int64 `tfsdk:"output_schema_expiration_time" tf:"optional"`
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	OutputSchemaName types.String `tfsdk:"output_schema_name" tf:"optional"`
	// Duration of the task run, in milliseconds.
	RunDuration types.Int64 `tfsdk:"run_duration" tf:"optional"`
	// When the task run started, in epoch milliseconds.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
}

func (newState *CleanRoomNotebookTaskRun) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomNotebookTaskRun) {
}

func (newState *CleanRoomNotebookTaskRun) SyncEffectiveFieldsDuringRead(existingState CleanRoomNotebookTaskRun) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomNotebookTaskRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomNotebookTaskRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"collaborator_job_run_info": reflect.TypeOf(CollaboratorJobRunInfo{}),
		"notebook_job_run_state":    reflect.TypeOf(jobs_tf.CleanRoomTaskRunState{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookTaskRun
// only implements ToObjectValue() and Type().
func (o CleanRoomNotebookTaskRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_job_run_info":     o.CollaboratorJobRunInfo,
			"notebook_job_run_state":        o.NotebookJobRunState,
			"notebook_name":                 o.NotebookName,
			"output_schema_expiration_time": o.OutputSchemaExpirationTime,
			"output_schema_name":            o.OutputSchemaName,
			"run_duration":                  o.RunDuration,
			"start_time":                    o.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomNotebookTaskRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"collaborator_job_run_info": basetypes.ListType{
				ElemType: CollaboratorJobRunInfo{}.Type(ctx),
			},
			"notebook_job_run_state": basetypes.ListType{
				ElemType: jobs_tf.CleanRoomTaskRunState{}.Type(ctx),
			},
			"notebook_name":                 types.StringType,
			"output_schema_expiration_time": types.Int64Type,
			"output_schema_name":            types.StringType,
			"run_duration":                  types.Int64Type,
			"start_time":                    types.Int64Type,
		},
	}
}

// GetCollaboratorJobRunInfo returns the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun as
// a CollaboratorJobRunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomNotebookTaskRun) GetCollaboratorJobRunInfo(ctx context.Context) (CollaboratorJobRunInfo, bool) {
	var e CollaboratorJobRunInfo
	if o.CollaboratorJobRunInfo.IsNull() || o.CollaboratorJobRunInfo.IsUnknown() {
		return e, false
	}
	var v []CollaboratorJobRunInfo
	d := o.CollaboratorJobRunInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCollaboratorJobRunInfo sets the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun.
func (o *CleanRoomNotebookTaskRun) SetCollaboratorJobRunInfo(ctx context.Context, v CollaboratorJobRunInfo) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborator_job_run_info"]
	o.CollaboratorJobRunInfo = types.ListValueMust(t, vs)
}

// GetNotebookJobRunState returns the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun as
// a jobs_tf.CleanRoomTaskRunState value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomNotebookTaskRun) GetNotebookJobRunState(ctx context.Context) (jobs_tf.CleanRoomTaskRunState, bool) {
	var e jobs_tf.CleanRoomTaskRunState
	if o.NotebookJobRunState.IsNull() || o.NotebookJobRunState.IsUnknown() {
		return e, false
	}
	var v []jobs_tf.CleanRoomTaskRunState
	d := o.NotebookJobRunState.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookJobRunState sets the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun.
func (o *CleanRoomNotebookTaskRun) SetNotebookJobRunState(ctx context.Context, v jobs_tf.CleanRoomTaskRunState) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_job_run_state"]
	o.NotebookJobRunState = types.ListValueMust(t, vs)
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`

	Status types.String `tfsdk:"status" tf:"optional"`
}

func (newState *CleanRoomOutputCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomOutputCatalog) {
}

func (newState *CleanRoomOutputCatalog) SyncEffectiveFieldsDuringRead(existingState CleanRoomOutputCatalog) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomOutputCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomOutputCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomOutputCatalog
// only implements ToObjectValue() and Type().
func (o CleanRoomOutputCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": o.CatalogName,
			"status":       o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomOutputCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"status":       types.StringType,
		},
	}
}

// Publicly visible central clean room details.
type CleanRoomRemoteDetail struct {
	// Central clean room ID.
	CentralCleanRoomId types.String `tfsdk:"central_clean_room_id" tf:"optional"`
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	CloudVendor types.String `tfsdk:"cloud_vendor" tf:"optional"`
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	Collaborators types.List `tfsdk:"collaborators" tf:"optional"`
	// The compliance security profile used to process regulated data following
	// compliance standards.
	ComplianceSecurityProfile types.List `tfsdk:"compliance_security_profile" tf:"optional,object"`
	// Collaborator who creates the clean room.
	Creator types.List `tfsdk:"creator" tf:"optional,object"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy types.List `tfsdk:"egress_network_policy" tf:"optional,object"`
	// Region of the central clean room.
	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *CleanRoomRemoteDetail) SyncEffectiveFieldsDuringCreateOrUpdate(plan CleanRoomRemoteDetail) {
}

func (newState *CleanRoomRemoteDetail) SyncEffectiveFieldsDuringRead(existingState CleanRoomRemoteDetail) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomRemoteDetail.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomRemoteDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"collaborators":               reflect.TypeOf(CleanRoomCollaborator{}),
		"compliance_security_profile": reflect.TypeOf(ComplianceSecurityProfile{}),
		"creator":                     reflect.TypeOf(CleanRoomCollaborator{}),
		"egress_network_policy":       reflect.TypeOf(settings_tf.EgressNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomRemoteDetail
// only implements ToObjectValue() and Type().
func (o CleanRoomRemoteDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"central_clean_room_id":       o.CentralCleanRoomId,
			"cloud_vendor":                o.CloudVendor,
			"collaborators":               o.Collaborators,
			"compliance_security_profile": o.ComplianceSecurityProfile,
			"creator":                     o.Creator,
			"egress_network_policy":       o.EgressNetworkPolicy,
			"region":                      o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomRemoteDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"central_clean_room_id": types.StringType,
			"cloud_vendor":          types.StringType,
			"collaborators": basetypes.ListType{
				ElemType: CleanRoomCollaborator{}.Type(ctx),
			},
			"compliance_security_profile": basetypes.ListType{
				ElemType: ComplianceSecurityProfile{}.Type(ctx),
			},
			"creator": basetypes.ListType{
				ElemType: CleanRoomCollaborator{}.Type(ctx),
			},
			"egress_network_policy": basetypes.ListType{
				ElemType: settings_tf.EgressNetworkPolicy{}.Type(ctx),
			},
			"region": types.StringType,
		},
	}
}

// GetCollaborators returns the value of the Collaborators field in CleanRoomRemoteDetail as
// a slice of CleanRoomCollaborator values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail) GetCollaborators(ctx context.Context) ([]CleanRoomCollaborator, bool) {
	if o.Collaborators.IsNull() || o.Collaborators.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomCollaborator
	d := o.Collaborators.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaborators sets the value of the Collaborators field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetCollaborators(ctx context.Context, v []CleanRoomCollaborator) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborators"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Collaborators = types.ListValueMust(t, vs)
}

// GetComplianceSecurityProfile returns the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail as
// a ComplianceSecurityProfile value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail) GetComplianceSecurityProfile(ctx context.Context) (ComplianceSecurityProfile, bool) {
	var e ComplianceSecurityProfile
	if o.ComplianceSecurityProfile.IsNull() || o.ComplianceSecurityProfile.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfile
	d := o.ComplianceSecurityProfile.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComplianceSecurityProfile sets the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetComplianceSecurityProfile(ctx context.Context, v ComplianceSecurityProfile) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_security_profile"]
	o.ComplianceSecurityProfile = types.ListValueMust(t, vs)
}

// GetCreator returns the value of the Creator field in CleanRoomRemoteDetail as
// a CleanRoomCollaborator value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail) GetCreator(ctx context.Context) (CleanRoomCollaborator, bool) {
	var e CleanRoomCollaborator
	if o.Creator.IsNull() || o.Creator.IsUnknown() {
		return e, false
	}
	var v []CleanRoomCollaborator
	d := o.Creator.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCreator sets the value of the Creator field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetCreator(ctx context.Context, v CleanRoomCollaborator) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["creator"]
	o.Creator = types.ListValueMust(t, vs)
}

// GetEgressNetworkPolicy returns the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail as
// a settings_tf.EgressNetworkPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail) GetEgressNetworkPolicy(ctx context.Context) (settings_tf.EgressNetworkPolicy, bool) {
	var e settings_tf.EgressNetworkPolicy
	if o.EgressNetworkPolicy.IsNull() || o.EgressNetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []settings_tf.EgressNetworkPolicy
	d := o.EgressNetworkPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEgressNetworkPolicy sets the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetEgressNetworkPolicy(ctx context.Context, v settings_tf.EgressNetworkPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["egress_network_policy"]
	o.EgressNetworkPolicy = types.ListValueMust(t, vs)
}

type CollaboratorJobRunInfo struct {
	// Alias of the collaborator that triggered the task run.
	CollaboratorAlias types.String `tfsdk:"collaborator_alias" tf:"optional"`
	// Job ID of the task run in the collaborator's workspace.
	CollaboratorJobId types.Int64 `tfsdk:"collaborator_job_id" tf:"optional"`
	// Job run ID of the task run in the collaborator's workspace.
	CollaboratorJobRunId types.Int64 `tfsdk:"collaborator_job_run_id" tf:"optional"`
	// Task run ID of the task run in the collaborator's workspace.
	CollaboratorTaskRunId types.Int64 `tfsdk:"collaborator_task_run_id" tf:"optional"`
	// ID of the collaborator's workspace that triggered the task run.
	CollaboratorWorkspaceId types.Int64 `tfsdk:"collaborator_workspace_id" tf:"optional"`
}

func (newState *CollaboratorJobRunInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan CollaboratorJobRunInfo) {
}

func (newState *CollaboratorJobRunInfo) SyncEffectiveFieldsDuringRead(existingState CollaboratorJobRunInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CollaboratorJobRunInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CollaboratorJobRunInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CollaboratorJobRunInfo
// only implements ToObjectValue() and Type().
func (o CollaboratorJobRunInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_alias":        o.CollaboratorAlias,
			"collaborator_job_id":       o.CollaboratorJobId,
			"collaborator_job_run_id":   o.CollaboratorJobRunId,
			"collaborator_task_run_id":  o.CollaboratorTaskRunId,
			"collaborator_workspace_id": o.CollaboratorWorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CollaboratorJobRunInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"collaborator_alias":        types.StringType,
			"collaborator_job_id":       types.Int64Type,
			"collaborator_job_run_id":   types.Int64Type,
			"collaborator_task_run_id":  types.Int64Type,
			"collaborator_workspace_id": types.Int64Type,
		},
	}
}

// The compliance security profile used to process regulated data following
// compliance standards.
type ComplianceSecurityProfile struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	ComplianceStandards types.List `tfsdk:"compliance_standards" tf:"optional"`
	// Whether the compliance security profile is enabled.
	IsEnabled types.Bool `tfsdk:"is_enabled" tf:"optional"`
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfile) {
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfile) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplianceSecurityProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComplianceSecurityProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile
// only implements ToObjectValue() and Type().
func (o ComplianceSecurityProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": o.ComplianceStandards,
			"is_enabled":           o.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplianceSecurityProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enabled": types.BoolType,
		},
	}
}

// GetComplianceStandards returns the value of the ComplianceStandards field in ComplianceSecurityProfile as
// a slice of settings_tf.ComplianceStandard values.
// If the field is unknown or null, the boolean return value is false.
func (o *ComplianceSecurityProfile) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
	if o.ComplianceStandards.IsNull() || o.ComplianceStandards.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ComplianceStandards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile.
func (o *ComplianceSecurityProfile) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ComplianceStandards = types.ListValueMust(t, vs)
}

// Create an asset
type CreateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	Asset types.List `tfsdk:"asset" tf:"optional,object"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
}

func (newState *CreateCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomAssetRequest) {
}

func (newState *CreateCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomAssetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset":           o.Asset,
			"clean_room_name": o.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset": basetypes.ListType{
				ElemType: CleanRoomAsset{}.Type(ctx),
			},
			"clean_room_name": types.StringType,
		},
	}
}

// GetAsset returns the value of the Asset field in CreateCleanRoomAssetRequest as
// a CleanRoomAsset value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAssetRequest) GetAsset(ctx context.Context) (CleanRoomAsset, bool) {
	var e CleanRoomAsset
	if o.Asset.IsNull() || o.Asset.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAsset
	d := o.Asset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAsset sets the value of the Asset field in CreateCleanRoomAssetRequest.
func (o *CreateCleanRoomAssetRequest) SetAsset(ctx context.Context, v CleanRoomAsset) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["asset"]
	o.Asset = types.ListValueMust(t, vs)
}

// Create an output catalog
type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`

	OutputCatalog types.List `tfsdk:"output_catalog" tf:"optional,object"`
}

func (newState *CreateCleanRoomOutputCatalogRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomOutputCatalogRequest) {
}

func (newState *CreateCleanRoomOutputCatalogRequest) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomOutputCatalogRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomOutputCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomOutputCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogRequest
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomOutputCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"output_catalog":  o.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomOutputCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"output_catalog": basetypes.ListType{
				ElemType: CleanRoomOutputCatalog{}.Type(ctx),
			},
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest as
// a CleanRoomOutputCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomOutputCatalogRequest) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog, bool) {
	var e CleanRoomOutputCatalog
	if o.OutputCatalog.IsNull() || o.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog
	d := o.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest.
func (o *CreateCleanRoomOutputCatalogRequest) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	o.OutputCatalog = types.ListValueMust(t, vs)
}

type CreateCleanRoomOutputCatalogResponse struct {
	OutputCatalog types.List `tfsdk:"output_catalog" tf:"optional,object"`
}

func (newState *CreateCleanRoomOutputCatalogResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomOutputCatalogResponse) {
}

func (newState *CreateCleanRoomOutputCatalogResponse) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomOutputCatalogResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomOutputCatalogResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomOutputCatalogResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogResponse
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomOutputCatalogResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"output_catalog": o.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomOutputCatalogResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"output_catalog": basetypes.ListType{
				ElemType: CleanRoomOutputCatalog{}.Type(ctx),
			},
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse as
// a CleanRoomOutputCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomOutputCatalogResponse) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog, bool) {
	var e CleanRoomOutputCatalog
	if o.OutputCatalog.IsNull() || o.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog
	d := o.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse.
func (o *CreateCleanRoomOutputCatalogResponse) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	o.OutputCatalog = types.ListValueMust(t, vs)
}

// Create a clean room
type CreateCleanRoomRequest struct {
	CleanRoom types.List `tfsdk:"clean_room" tf:"optional,object"`
}

func (newState *CreateCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCleanRoomRequest) {
}

func (newState *CreateCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState CreateCleanRoomRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomRequest
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": o.CleanRoom,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room": basetypes.ListType{
				ElemType: CleanRoom{}.Type(ctx),
			},
		},
	}
}

// GetCleanRoom returns the value of the CleanRoom field in CreateCleanRoomRequest as
// a CleanRoom value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomRequest) GetCleanRoom(ctx context.Context) (CleanRoom, bool) {
	var e CleanRoom
	if o.CleanRoom.IsNull() || o.CleanRoom.IsUnknown() {
		return e, false
	}
	var v []CleanRoom
	d := o.CleanRoom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoom sets the value of the CleanRoom field in CreateCleanRoomRequest.
func (o *CreateCleanRoomRequest) SetCleanRoom(ctx context.Context, v CleanRoom) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_room"]
	o.CleanRoom = types.ListValueMust(t, vs)
}

// Delete an asset
type DeleteCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName types.String `tfsdk:"-"`
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
}

func (newState *DeleteCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCleanRoomAssetRequest) {
}

func (newState *DeleteCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCleanRoomAssetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_full_name": o.AssetFullName,
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_full_name": types.StringType,
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
		},
	}
}

// Response for delete clean room request. Using an empty message since the
// generic Empty proto does not externd UnshadedMessageMarker.
type DeleteCleanRoomAssetResponse struct {
}

func (newState *DeleteCleanRoomAssetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCleanRoomAssetResponse) {
}

func (newState *DeleteCleanRoomAssetResponse) SyncEffectiveFieldsDuringRead(existingState DeleteCleanRoomAssetResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAssetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomAssetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetResponse
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomAssetResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAssetResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a clean room
type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCleanRoomRequest) {
}

func (newState *DeleteCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCleanRoomRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomRequest
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (o DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get an asset
type GetCleanRoomAssetRequest struct {
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	AssetFullName types.String `tfsdk:"-"`
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
}

func (newState *GetCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCleanRoomAssetRequest) {
}

func (newState *GetCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState GetCleanRoomAssetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (o GetCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_full_name": o.AssetFullName,
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_full_name": types.StringType,
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
		},
	}
}

// Get a clean room
type GetCleanRoomRequest struct {
	Name types.String `tfsdk:"-"`
}

func (newState *GetCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCleanRoomRequest) {
}

func (newState *GetCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState GetCleanRoomRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomRequest
// only implements ToObjectValue() and Type().
func (o GetCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// List assets
type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCleanRoomAssetsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomAssetsRequest) {
}

func (newState *ListCleanRoomAssetsRequest) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomAssetsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAssetsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAssetsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsRequest
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomAssetsResponse struct {
	// Assets in the clean room.
	Assets types.List `tfsdk:"assets" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCleanRoomAssetsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomAssetsResponse) {
}

func (newState *ListCleanRoomAssetsResponse) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomAssetsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAssetsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAssetsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsResponse
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":          o.Assets,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: CleanRoomAsset{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetAssets returns the value of the Assets field in ListCleanRoomAssetsResponse as
// a slice of CleanRoomAsset values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomAssetsResponse) GetAssets(ctx context.Context) ([]CleanRoomAsset, bool) {
	if o.Assets.IsNull() || o.Assets.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset
	d := o.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListCleanRoomAssetsResponse.
func (o *ListCleanRoomAssetsResponse) SetAssets(ctx context.Context, v []CleanRoomAsset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

// List notebook task runs
type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Notebook name
	NotebookName types.String `tfsdk:"-"`
	// The maximum number of task runs to return
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCleanRoomNotebookTaskRunsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomNotebookTaskRunsRequest) {
}

func (newState *ListCleanRoomNotebookTaskRunsRequest) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomNotebookTaskRunsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomNotebookTaskRunsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomNotebookTaskRunsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsRequest
// only implements ToObjectValue() and Type().
func (o ListCleanRoomNotebookTaskRunsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"notebook_name":   o.NotebookName,
			"page_size":       o.PageSize,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomNotebookTaskRunsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"notebook_name":   types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomNotebookTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// Name of the clean room.
	Runs types.List `tfsdk:"runs" tf:"optional"`
}

func (newState *ListCleanRoomNotebookTaskRunsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomNotebookTaskRunsResponse) {
}

func (newState *ListCleanRoomNotebookTaskRunsResponse) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomNotebookTaskRunsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomNotebookTaskRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomNotebookTaskRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(CleanRoomNotebookTaskRun{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsResponse
// only implements ToObjectValue() and Type().
func (o ListCleanRoomNotebookTaskRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"runs":            o.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomNotebookTaskRunsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: CleanRoomNotebookTaskRun{}.Type(ctx),
			},
		},
	}
}

// GetRuns returns the value of the Runs field in ListCleanRoomNotebookTaskRunsResponse as
// a slice of CleanRoomNotebookTaskRun values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomNotebookTaskRunsResponse) GetRuns(ctx context.Context) ([]CleanRoomNotebookTaskRun, bool) {
	if o.Runs.IsNull() || o.Runs.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookTaskRun
	d := o.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in ListCleanRoomNotebookTaskRunsResponse.
func (o *ListCleanRoomNotebookTaskRunsResponse) SetRuns(ctx context.Context, v []CleanRoomNotebookTaskRun) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Runs = types.ListValueMust(t, vs)
}

// List clean rooms
type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCleanRoomsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomsRequest) {
}

func (newState *ListCleanRoomsRequest) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsRequest
// only implements ToObjectValue() and Type().
func (o ListCleanRoomsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListCleanRoomsResponse struct {
	CleanRooms types.List `tfsdk:"clean_rooms" tf:"optional"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *ListCleanRoomsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCleanRoomsResponse) {
}

func (newState *ListCleanRoomsResponse) SyncEffectiveFieldsDuringRead(existingState ListCleanRoomsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms": reflect.TypeOf(CleanRoom{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsResponse
// only implements ToObjectValue() and Type().
func (o ListCleanRoomsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_rooms":     o.CleanRooms,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_rooms": basetypes.ListType{
				ElemType: CleanRoom{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetCleanRooms returns the value of the CleanRooms field in ListCleanRoomsResponse as
// a slice of CleanRoom values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomsResponse) GetCleanRooms(ctx context.Context) ([]CleanRoom, bool) {
	if o.CleanRooms.IsNull() || o.CleanRooms.IsUnknown() {
		return nil, false
	}
	var v []CleanRoom
	d := o.CleanRooms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRooms sets the value of the CleanRooms field in ListCleanRoomsResponse.
func (o *ListCleanRoomsResponse) SetCleanRooms(ctx context.Context, v []CleanRoom) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CleanRooms = types.ListValueMust(t, vs)
}

// Update an asset
type UpdateCleanRoomAssetRequest struct {
	// Metadata of the clean room asset
	Asset types.List `tfsdk:"asset" tf:"optional,object"`
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateCleanRoomAssetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCleanRoomAssetRequest) {
}

func (newState *UpdateCleanRoomAssetRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCleanRoomAssetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (o UpdateCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset":           o.Asset,
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset": basetypes.ListType{
				ElemType: CleanRoomAsset{}.Type(ctx),
			},
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
		},
	}
}

// GetAsset returns the value of the Asset field in UpdateCleanRoomAssetRequest as
// a CleanRoomAsset value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCleanRoomAssetRequest) GetAsset(ctx context.Context) (CleanRoomAsset, bool) {
	var e CleanRoomAsset
	if o.Asset.IsNull() || o.Asset.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAsset
	d := o.Asset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAsset sets the value of the Asset field in UpdateCleanRoomAssetRequest.
func (o *UpdateCleanRoomAssetRequest) SetAsset(ctx context.Context, v CleanRoomAsset) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["asset"]
	o.Asset = types.ListValueMust(t, vs)
}

type UpdateCleanRoomRequest struct {
	CleanRoom types.List `tfsdk:"clean_room" tf:"optional,object"`
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (newState *UpdateCleanRoomRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCleanRoomRequest) {
}

func (newState *UpdateCleanRoomRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCleanRoomRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomRequest
// only implements ToObjectValue() and Type().
func (o UpdateCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": o.CleanRoom,
			"name":       o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCleanRoomRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room": basetypes.ListType{
				ElemType: CleanRoom{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetCleanRoom returns the value of the CleanRoom field in UpdateCleanRoomRequest as
// a CleanRoom value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCleanRoomRequest) GetCleanRoom(ctx context.Context) (CleanRoom, bool) {
	var e CleanRoom
	if o.CleanRoom.IsNull() || o.CleanRoom.IsUnknown() {
		return e, false
	}
	var v []CleanRoom
	d := o.CleanRoom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoom sets the value of the CleanRoom field in UpdateCleanRoomRequest.
func (o *UpdateCleanRoomRequest) SetCleanRoom(ctx context.Context, v CleanRoom) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_room"]
	o.CleanRoom = types.ListValueMust(t, vs)
}
