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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/databricks/terraform-provider-databricks/internal/service/jobs_tf"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CleanRoom struct {
	// Whether clean room access is restricted due to [CSP]
	//
	// [CSP]: https://docs.databricks.com/en/security/privacy/security-profile.html
	AccessRestricted types.String `tfsdk:"access_restricted"`

	Comment types.String `tfsdk:"comment"`
	// When the clean room was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// The alias of the collaborator tied to the local clean room.
	LocalCollaboratorAlias types.String `tfsdk:"local_collaborator_alias"`
	// The name of the clean room. It should follow [UC securable naming
	// requirements].
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	Name types.String `tfsdk:"name"`
	// Output catalog of the clean room. It is an output only field. Output
	// catalog is manipulated using the separate CreateCleanRoomOutputCatalog
	// API.
	OutputCatalog types.Object `tfsdk:"output_catalog"`
	// This is Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner types.String `tfsdk:"owner"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo types.Object `tfsdk:"remote_detailed_info"`
	// Clean room status.
	Status types.String `tfsdk:"status"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (toState *CleanRoom) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoom) {
	if !fromPlan.OutputCatalog.IsNull() && !fromPlan.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromPlanOutputCatalog, ok := fromPlan.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
	if !fromPlan.RemoteDetailedInfo.IsNull() && !fromPlan.RemoteDetailedInfo.IsUnknown() {
		if toStateRemoteDetailedInfo, ok := toState.GetRemoteDetailedInfo(ctx); ok {
			if fromPlanRemoteDetailedInfo, ok := fromPlan.GetRemoteDetailedInfo(ctx); ok {
				toStateRemoteDetailedInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanRemoteDetailedInfo)
				toState.SetRemoteDetailedInfo(ctx, toStateRemoteDetailedInfo)
			}
		}
	}
}

func (toState *CleanRoom) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoom) {
	if !fromState.OutputCatalog.IsNull() && !fromState.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromStateOutputCatalog, ok := fromState.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringRead(ctx, fromStateOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
	if !fromState.RemoteDetailedInfo.IsNull() && !fromState.RemoteDetailedInfo.IsUnknown() {
		if toStateRemoteDetailedInfo, ok := toState.GetRemoteDetailedInfo(ctx); ok {
			if fromStateRemoteDetailedInfo, ok := fromState.GetRemoteDetailedInfo(ctx); ok {
				toStateRemoteDetailedInfo.SyncFieldsDuringRead(ctx, fromStateRemoteDetailedInfo)
				toState.SetRemoteDetailedInfo(ctx, toStateRemoteDetailedInfo)
			}
		}
	}
}

func (c CleanRoom) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_restricted"] = attrs["access_restricted"].SetComputed()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["local_collaborator_alias"] = attrs["local_collaborator_alias"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["output_catalog"] = attrs["output_catalog"].SetComputed()
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["remote_detailed_info"] = attrs["remote_detailed_info"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["updated_at"] = attrs["updated_at"].SetComputed()

	return attrs
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
			"output_catalog":           CleanRoomOutputCatalog{}.Type(ctx),
			"owner":                    types.StringType,
			"remote_detailed_info":     CleanRoomRemoteDetail{}.Type(ctx),
			"status":                   types.StringType,
			"updated_at":               types.Int64Type,
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
	var v CleanRoomOutputCatalog
	d := o.OutputCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CleanRoom.
func (o *CleanRoom) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := v.ToObjectValue(ctx)
	o.OutputCatalog = vs
}

// GetRemoteDetailedInfo returns the value of the RemoteDetailedInfo field in CleanRoom as
// a CleanRoomRemoteDetail value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoom) GetRemoteDetailedInfo(ctx context.Context) (CleanRoomRemoteDetail, bool) {
	var e CleanRoomRemoteDetail
	if o.RemoteDetailedInfo.IsNull() || o.RemoteDetailedInfo.IsUnknown() {
		return e, false
	}
	var v CleanRoomRemoteDetail
	d := o.RemoteDetailedInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRemoteDetailedInfo sets the value of the RemoteDetailedInfo field in CleanRoom.
func (o *CleanRoom) SetRemoteDetailedInfo(ctx context.Context, v CleanRoomRemoteDetail) {
	vs := v.ToObjectValue(ctx)
	o.RemoteDetailedInfo = vs
}

// Metadata of the clean room asset
type CleanRoomAsset struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at"`
	// The type of the asset.
	AssetType types.String `tfsdk:"asset_type"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName types.String `tfsdk:"clean_room_name"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable types.Object `tfsdk:"foreign_table"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails types.Object `tfsdk:"foreign_table_local_details"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name.
	Name types.String `tfsdk:"name"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook types.Object `tfsdk:"notebook"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias types.String `tfsdk:"owner_collaborator_alias"`
	// Status of the asset
	Status types.String `tfsdk:"status"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table types.Object `tfsdk:"table"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails types.Object `tfsdk:"table_local_details"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View types.Object `tfsdk:"view"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails types.Object `tfsdk:"view_local_details"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails types.Object `tfsdk:"volume_local_details"`
}

func (toState *CleanRoomAsset) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAsset) {
	if !fromPlan.ForeignTable.IsNull() && !fromPlan.ForeignTable.IsUnknown() {
		if toStateForeignTable, ok := toState.GetForeignTable(ctx); ok {
			if fromPlanForeignTable, ok := fromPlan.GetForeignTable(ctx); ok {
				toStateForeignTable.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanForeignTable)
				toState.SetForeignTable(ctx, toStateForeignTable)
			}
		}
	}
	if !fromPlan.ForeignTableLocalDetails.IsNull() && !fromPlan.ForeignTableLocalDetails.IsUnknown() {
		if toStateForeignTableLocalDetails, ok := toState.GetForeignTableLocalDetails(ctx); ok {
			if fromPlanForeignTableLocalDetails, ok := fromPlan.GetForeignTableLocalDetails(ctx); ok {
				toStateForeignTableLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanForeignTableLocalDetails)
				toState.SetForeignTableLocalDetails(ctx, toStateForeignTableLocalDetails)
			}
		}
	}
	if !fromPlan.Notebook.IsNull() && !fromPlan.Notebook.IsUnknown() {
		if toStateNotebook, ok := toState.GetNotebook(ctx); ok {
			if fromPlanNotebook, ok := fromPlan.GetNotebook(ctx); ok {
				toStateNotebook.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanNotebook)
				toState.SetNotebook(ctx, toStateNotebook)
			}
		}
	}
	if !fromPlan.Table.IsNull() && !fromPlan.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromPlanTable, ok := fromPlan.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
	if !fromPlan.TableLocalDetails.IsNull() && !fromPlan.TableLocalDetails.IsUnknown() {
		if toStateTableLocalDetails, ok := toState.GetTableLocalDetails(ctx); ok {
			if fromPlanTableLocalDetails, ok := fromPlan.GetTableLocalDetails(ctx); ok {
				toStateTableLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTableLocalDetails)
				toState.SetTableLocalDetails(ctx, toStateTableLocalDetails)
			}
		}
	}
	if !fromPlan.View.IsNull() && !fromPlan.View.IsUnknown() {
		if toStateView, ok := toState.GetView(ctx); ok {
			if fromPlanView, ok := fromPlan.GetView(ctx); ok {
				toStateView.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanView)
				toState.SetView(ctx, toStateView)
			}
		}
	}
	if !fromPlan.ViewLocalDetails.IsNull() && !fromPlan.ViewLocalDetails.IsUnknown() {
		if toStateViewLocalDetails, ok := toState.GetViewLocalDetails(ctx); ok {
			if fromPlanViewLocalDetails, ok := fromPlan.GetViewLocalDetails(ctx); ok {
				toStateViewLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanViewLocalDetails)
				toState.SetViewLocalDetails(ctx, toStateViewLocalDetails)
			}
		}
	}
	if !fromPlan.VolumeLocalDetails.IsNull() && !fromPlan.VolumeLocalDetails.IsUnknown() {
		if toStateVolumeLocalDetails, ok := toState.GetVolumeLocalDetails(ctx); ok {
			if fromPlanVolumeLocalDetails, ok := fromPlan.GetVolumeLocalDetails(ctx); ok {
				toStateVolumeLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanVolumeLocalDetails)
				toState.SetVolumeLocalDetails(ctx, toStateVolumeLocalDetails)
			}
		}
	}
}

func (toState *CleanRoomAsset) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAsset) {
	if !fromState.ForeignTable.IsNull() && !fromState.ForeignTable.IsUnknown() {
		if toStateForeignTable, ok := toState.GetForeignTable(ctx); ok {
			if fromStateForeignTable, ok := fromState.GetForeignTable(ctx); ok {
				toStateForeignTable.SyncFieldsDuringRead(ctx, fromStateForeignTable)
				toState.SetForeignTable(ctx, toStateForeignTable)
			}
		}
	}
	if !fromState.ForeignTableLocalDetails.IsNull() && !fromState.ForeignTableLocalDetails.IsUnknown() {
		if toStateForeignTableLocalDetails, ok := toState.GetForeignTableLocalDetails(ctx); ok {
			if fromStateForeignTableLocalDetails, ok := fromState.GetForeignTableLocalDetails(ctx); ok {
				toStateForeignTableLocalDetails.SyncFieldsDuringRead(ctx, fromStateForeignTableLocalDetails)
				toState.SetForeignTableLocalDetails(ctx, toStateForeignTableLocalDetails)
			}
		}
	}
	if !fromState.Notebook.IsNull() && !fromState.Notebook.IsUnknown() {
		if toStateNotebook, ok := toState.GetNotebook(ctx); ok {
			if fromStateNotebook, ok := fromState.GetNotebook(ctx); ok {
				toStateNotebook.SyncFieldsDuringRead(ctx, fromStateNotebook)
				toState.SetNotebook(ctx, toStateNotebook)
			}
		}
	}
	if !fromState.Table.IsNull() && !fromState.Table.IsUnknown() {
		if toStateTable, ok := toState.GetTable(ctx); ok {
			if fromStateTable, ok := fromState.GetTable(ctx); ok {
				toStateTable.SyncFieldsDuringRead(ctx, fromStateTable)
				toState.SetTable(ctx, toStateTable)
			}
		}
	}
	if !fromState.TableLocalDetails.IsNull() && !fromState.TableLocalDetails.IsUnknown() {
		if toStateTableLocalDetails, ok := toState.GetTableLocalDetails(ctx); ok {
			if fromStateTableLocalDetails, ok := fromState.GetTableLocalDetails(ctx); ok {
				toStateTableLocalDetails.SyncFieldsDuringRead(ctx, fromStateTableLocalDetails)
				toState.SetTableLocalDetails(ctx, toStateTableLocalDetails)
			}
		}
	}
	if !fromState.View.IsNull() && !fromState.View.IsUnknown() {
		if toStateView, ok := toState.GetView(ctx); ok {
			if fromStateView, ok := fromState.GetView(ctx); ok {
				toStateView.SyncFieldsDuringRead(ctx, fromStateView)
				toState.SetView(ctx, toStateView)
			}
		}
	}
	if !fromState.ViewLocalDetails.IsNull() && !fromState.ViewLocalDetails.IsUnknown() {
		if toStateViewLocalDetails, ok := toState.GetViewLocalDetails(ctx); ok {
			if fromStateViewLocalDetails, ok := fromState.GetViewLocalDetails(ctx); ok {
				toStateViewLocalDetails.SyncFieldsDuringRead(ctx, fromStateViewLocalDetails)
				toState.SetViewLocalDetails(ctx, toStateViewLocalDetails)
			}
		}
	}
	if !fromState.VolumeLocalDetails.IsNull() && !fromState.VolumeLocalDetails.IsUnknown() {
		if toStateVolumeLocalDetails, ok := toState.GetVolumeLocalDetails(ctx); ok {
			if fromStateVolumeLocalDetails, ok := fromState.GetVolumeLocalDetails(ctx); ok {
				toStateVolumeLocalDetails.SyncFieldsDuringRead(ctx, fromStateVolumeLocalDetails)
				toState.SetVolumeLocalDetails(ctx, toStateVolumeLocalDetails)
			}
		}
	}
}

func (c CleanRoomAsset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["added_at"] = attrs["added_at"].SetComputed()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["clean_room_name"] = attrs["clean_room_name"].SetOptional()
	attrs["foreign_table"] = attrs["foreign_table"].SetOptional()
	attrs["foreign_table_local_details"] = attrs["foreign_table_local_details"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["notebook"] = attrs["notebook"].SetOptional()
	attrs["owner_collaborator_alias"] = attrs["owner_collaborator_alias"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["table"] = attrs["table"].SetOptional()
	attrs["table_local_details"] = attrs["table_local_details"].SetOptional()
	attrs["view"] = attrs["view"].SetOptional()
	attrs["view_local_details"] = attrs["view_local_details"].SetOptional()
	attrs["volume_local_details"] = attrs["volume_local_details"].SetOptional()

	return attrs
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
			"clean_room_name":             o.CleanRoomName,
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
			"added_at":                    types.Int64Type,
			"asset_type":                  types.StringType,
			"clean_room_name":             types.StringType,
			"foreign_table":               CleanRoomAssetForeignTable{}.Type(ctx),
			"foreign_table_local_details": CleanRoomAssetForeignTableLocalDetails{}.Type(ctx),
			"name":                        types.StringType,
			"notebook":                    CleanRoomAssetNotebook{}.Type(ctx),
			"owner_collaborator_alias":    types.StringType,
			"status":                      types.StringType,
			"table":                       CleanRoomAssetTable{}.Type(ctx),
			"table_local_details":         CleanRoomAssetTableLocalDetails{}.Type(ctx),
			"view":                        CleanRoomAssetView{}.Type(ctx),
			"view_local_details":          CleanRoomAssetViewLocalDetails{}.Type(ctx),
			"volume_local_details":        CleanRoomAssetVolumeLocalDetails{}.Type(ctx),
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
	var v CleanRoomAssetForeignTable
	d := o.ForeignTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetForeignTable sets the value of the ForeignTable field in CleanRoomAsset.
func (o *CleanRoomAsset) SetForeignTable(ctx context.Context, v CleanRoomAssetForeignTable) {
	vs := v.ToObjectValue(ctx)
	o.ForeignTable = vs
}

// GetForeignTableLocalDetails returns the value of the ForeignTableLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetForeignTableLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetForeignTableLocalDetails(ctx context.Context) (CleanRoomAssetForeignTableLocalDetails, bool) {
	var e CleanRoomAssetForeignTableLocalDetails
	if o.ForeignTableLocalDetails.IsNull() || o.ForeignTableLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetForeignTableLocalDetails
	d := o.ForeignTableLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetForeignTableLocalDetails sets the value of the ForeignTableLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetForeignTableLocalDetails(ctx context.Context, v CleanRoomAssetForeignTableLocalDetails) {
	vs := v.ToObjectValue(ctx)
	o.ForeignTableLocalDetails = vs
}

// GetNotebook returns the value of the Notebook field in CleanRoomAsset as
// a CleanRoomAssetNotebook value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetNotebook(ctx context.Context) (CleanRoomAssetNotebook, bool) {
	var e CleanRoomAssetNotebook
	if o.Notebook.IsNull() || o.Notebook.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetNotebook
	d := o.Notebook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebook sets the value of the Notebook field in CleanRoomAsset.
func (o *CleanRoomAsset) SetNotebook(ctx context.Context, v CleanRoomAssetNotebook) {
	vs := v.ToObjectValue(ctx)
	o.Notebook = vs
}

// GetTable returns the value of the Table field in CleanRoomAsset as
// a CleanRoomAssetTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetTable(ctx context.Context) (CleanRoomAssetTable, bool) {
	var e CleanRoomAssetTable
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetTable
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTable sets the value of the Table field in CleanRoomAsset.
func (o *CleanRoomAsset) SetTable(ctx context.Context, v CleanRoomAssetTable) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

// GetTableLocalDetails returns the value of the TableLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetTableLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetTableLocalDetails(ctx context.Context) (CleanRoomAssetTableLocalDetails, bool) {
	var e CleanRoomAssetTableLocalDetails
	if o.TableLocalDetails.IsNull() || o.TableLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetTableLocalDetails
	d := o.TableLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableLocalDetails sets the value of the TableLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetTableLocalDetails(ctx context.Context, v CleanRoomAssetTableLocalDetails) {
	vs := v.ToObjectValue(ctx)
	o.TableLocalDetails = vs
}

// GetView returns the value of the View field in CleanRoomAsset as
// a CleanRoomAssetView value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetView(ctx context.Context) (CleanRoomAssetView, bool) {
	var e CleanRoomAssetView
	if o.View.IsNull() || o.View.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetView
	d := o.View.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetView sets the value of the View field in CleanRoomAsset.
func (o *CleanRoomAsset) SetView(ctx context.Context, v CleanRoomAssetView) {
	vs := v.ToObjectValue(ctx)
	o.View = vs
}

// GetViewLocalDetails returns the value of the ViewLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetViewLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetViewLocalDetails(ctx context.Context) (CleanRoomAssetViewLocalDetails, bool) {
	var e CleanRoomAssetViewLocalDetails
	if o.ViewLocalDetails.IsNull() || o.ViewLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetViewLocalDetails
	d := o.ViewLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViewLocalDetails sets the value of the ViewLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetViewLocalDetails(ctx context.Context, v CleanRoomAssetViewLocalDetails) {
	vs := v.ToObjectValue(ctx)
	o.ViewLocalDetails = vs
}

// GetVolumeLocalDetails returns the value of the VolumeLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetVolumeLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset) GetVolumeLocalDetails(ctx context.Context) (CleanRoomAssetVolumeLocalDetails, bool) {
	var e CleanRoomAssetVolumeLocalDetails
	if o.VolumeLocalDetails.IsNull() || o.VolumeLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetVolumeLocalDetails
	d := o.VolumeLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumeLocalDetails sets the value of the VolumeLocalDetails field in CleanRoomAsset.
func (o *CleanRoomAsset) SetVolumeLocalDetails(ctx context.Context, v CleanRoomAssetVolumeLocalDetails) {
	vs := v.ToObjectValue(ctx)
	o.VolumeLocalDetails = vs
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	Columns types.List `tfsdk:"columns"`
}

func (toState *CleanRoomAssetForeignTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetForeignTable) {
}

func (toState *CleanRoomAssetForeignTable) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetForeignTable) {
}

func (c CleanRoomAssetForeignTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetComputed()

	return attrs
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
	LocalName types.String `tfsdk:"local_name"`
}

func (toState *CleanRoomAssetForeignTableLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetForeignTableLocalDetails) {
}

func (toState *CleanRoomAssetForeignTableLocalDetails) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetForeignTableLocalDetails) {
}

func (c CleanRoomAssetForeignTableLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["local_name"] = attrs["local_name"].SetRequired()

	return attrs
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
	// Server generated etag that represents the notebook version.
	Etag types.String `tfsdk:"etag"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent types.String `tfsdk:"notebook_content"`
	// top-level status derived from all reviews
	ReviewState types.String `tfsdk:"review_state"`
	// All existing approvals or rejections
	Reviews types.List `tfsdk:"reviews"`
	// collaborators that can run the notebook
	RunnerCollaboratorAliases types.List `tfsdk:"runner_collaborator_aliases"`
}

func (toState *CleanRoomAssetNotebook) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetNotebook) {
}

func (toState *CleanRoomAssetNotebook) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetNotebook) {
}

func (c CleanRoomAssetNotebook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["notebook_content"] = attrs["notebook_content"].SetRequired()
	attrs["review_state"] = attrs["review_state"].SetComputed()
	attrs["reviews"] = attrs["reviews"].SetComputed()
	attrs["runner_collaborator_aliases"] = attrs["runner_collaborator_aliases"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetNotebook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAssetNotebook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"reviews":                     reflect.TypeOf(CleanRoomNotebookReview{}),
		"runner_collaborator_aliases": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetNotebook
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetNotebook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                        o.Etag,
			"notebook_content":            o.NotebookContent,
			"review_state":                o.ReviewState,
			"reviews":                     o.Reviews,
			"runner_collaborator_aliases": o.RunnerCollaboratorAliases,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetNotebook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":             types.StringType,
			"notebook_content": types.StringType,
			"review_state":     types.StringType,
			"reviews": basetypes.ListType{
				ElemType: CleanRoomNotebookReview{}.Type(ctx),
			},
			"runner_collaborator_aliases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetReviews returns the value of the Reviews field in CleanRoomAssetNotebook as
// a slice of CleanRoomNotebookReview values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetNotebook) GetReviews(ctx context.Context) ([]CleanRoomNotebookReview, bool) {
	if o.Reviews.IsNull() || o.Reviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview
	d := o.Reviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReviews sets the value of the Reviews field in CleanRoomAssetNotebook.
func (o *CleanRoomAssetNotebook) SetReviews(ctx context.Context, v []CleanRoomNotebookReview) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Reviews = types.ListValueMust(t, vs)
}

// GetRunnerCollaboratorAliases returns the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetNotebook) GetRunnerCollaboratorAliases(ctx context.Context) ([]types.String, bool) {
	if o.RunnerCollaboratorAliases.IsNull() || o.RunnerCollaboratorAliases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RunnerCollaboratorAliases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunnerCollaboratorAliases sets the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook.
func (o *CleanRoomAssetNotebook) SetRunnerCollaboratorAliases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runner_collaborator_aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RunnerCollaboratorAliases = types.ListValueMust(t, vs)
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	Columns types.List `tfsdk:"columns"`
}

func (toState *CleanRoomAssetTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetTable) {
}

func (toState *CleanRoomAssetTable) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetTable) {
}

func (c CleanRoomAssetTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetComputed()

	return attrs
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
	LocalName types.String `tfsdk:"local_name"`
	// Partition filtering specification for a shared table.
	Partitions types.List `tfsdk:"partitions"`
}

func (toState *CleanRoomAssetTableLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetTableLocalDetails) {
}

func (toState *CleanRoomAssetTableLocalDetails) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetTableLocalDetails) {
}

func (c CleanRoomAssetTableLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["local_name"] = attrs["local_name"].SetRequired()
	attrs["partitions"] = attrs["partitions"].SetOptional()

	return attrs
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
		"partitions": reflect.TypeOf(sharing_tf.Partition{}),
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
				ElemType: sharing_tf.Partition{}.Type(ctx),
			},
		},
	}
}

// GetPartitions returns the value of the Partitions field in CleanRoomAssetTableLocalDetails as
// a slice of sharing_tf.Partition values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetTableLocalDetails) GetPartitions(ctx context.Context) ([]sharing_tf.Partition, bool) {
	if o.Partitions.IsNull() || o.Partitions.IsUnknown() {
		return nil, false
	}
	var v []sharing_tf.Partition
	d := o.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in CleanRoomAssetTableLocalDetails.
func (o *CleanRoomAssetTableLocalDetails) SetPartitions(ctx context.Context, v []sharing_tf.Partition) {
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
	Columns types.List `tfsdk:"columns"`
}

func (toState *CleanRoomAssetView) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetView) {
}

func (toState *CleanRoomAssetView) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetView) {
}

func (c CleanRoomAssetView) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetComputed()

	return attrs
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
	LocalName types.String `tfsdk:"local_name"`
}

func (toState *CleanRoomAssetViewLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetViewLocalDetails) {
}

func (toState *CleanRoomAssetViewLocalDetails) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetViewLocalDetails) {
}

func (c CleanRoomAssetViewLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["local_name"] = attrs["local_name"].SetRequired()

	return attrs
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
	LocalName types.String `tfsdk:"local_name"`
}

func (toState *CleanRoomAssetVolumeLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetVolumeLocalDetails) {
}

func (toState *CleanRoomAssetVolumeLocalDetails) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetVolumeLocalDetails) {
}

func (c CleanRoomAssetVolumeLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["local_name"] = attrs["local_name"].SetRequired()

	return attrs
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

type CleanRoomAutoApprovalRule struct {
	AuthorCollaboratorAlias types.String `tfsdk:"author_collaborator_alias"`

	AuthorScope types.String `tfsdk:"author_scope"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"clean_room_name"`
	// Timestamp of when the rule was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// A generated UUID identifying the rule.
	RuleId types.String `tfsdk:"rule_id"`
	// The owner of the rule to whom the rule applies.
	RuleOwnerCollaboratorAlias types.String `tfsdk:"rule_owner_collaborator_alias"`

	RunnerCollaboratorAlias types.String `tfsdk:"runner_collaborator_alias"`
}

func (toState *CleanRoomAutoApprovalRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAutoApprovalRule) {
}

func (toState *CleanRoomAutoApprovalRule) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAutoApprovalRule) {
}

func (c CleanRoomAutoApprovalRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["author_collaborator_alias"] = attrs["author_collaborator_alias"].SetOptional()
	attrs["author_scope"] = attrs["author_scope"].SetOptional()
	attrs["clean_room_name"] = attrs["clean_room_name"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["rule_id"] = attrs["rule_id"].SetComputed()
	attrs["rule_owner_collaborator_alias"] = attrs["rule_owner_collaborator_alias"].SetComputed()
	attrs["runner_collaborator_alias"] = attrs["runner_collaborator_alias"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAutoApprovalRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAutoApprovalRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAutoApprovalRule
// only implements ToObjectValue() and Type().
func (o CleanRoomAutoApprovalRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"author_collaborator_alias":     o.AuthorCollaboratorAlias,
			"author_scope":                  o.AuthorScope,
			"clean_room_name":               o.CleanRoomName,
			"created_at":                    o.CreatedAt,
			"rule_id":                       o.RuleId,
			"rule_owner_collaborator_alias": o.RuleOwnerCollaboratorAlias,
			"runner_collaborator_alias":     o.RunnerCollaboratorAlias,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAutoApprovalRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"author_collaborator_alias":     types.StringType,
			"author_scope":                  types.StringType,
			"clean_room_name":               types.StringType,
			"created_at":                    types.Int64Type,
			"rule_id":                       types.StringType,
			"rule_owner_collaborator_alias": types.StringType,
			"runner_collaborator_alias":     types.StringType,
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
	CollaboratorAlias types.String `tfsdk:"collaborator_alias"`
	// Generated display name for the collaborator. In the case of a single
	// metastore clean room, it is the clean room name. For x-metastore clean
	// rooms, it is the organization name of the metastore. It is not restricted
	// to these values and could change in the future
	DisplayName types.String `tfsdk:"display_name"`
	// The global Unity Catalog metastore id of the collaborator. The identifier
	// is of format cloud:region:metastore-uuid.
	GlobalMetastoreId types.String `tfsdk:"global_metastore_id"`
	// Email of the user who is receiving the clean room "invitation". It should
	// be empty for the creator of the clean room, and non-empty for the
	// invitees of the clean room. It is only returned in the output when clean
	// room creator calls GET
	InviteRecipientEmail types.String `tfsdk:"invite_recipient_email"`
	// Workspace ID of the user who is receiving the clean room "invitation".
	// Must be specified if invite_recipient_email is specified. It should be
	// empty when the collaborator is the creator of the clean room.
	InviteRecipientWorkspaceId types.Int64 `tfsdk:"invite_recipient_workspace_id"`
	// [Organization
	// name](:method:metastores/list#metastores-delta_sharing_organization_name)
	// configured in the metastore
	OrganizationName types.String `tfsdk:"organization_name"`
}

func (toState *CleanRoomCollaborator) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomCollaborator) {
	if !fromPlan.InviteRecipientWorkspaceId.IsUnknown() && !fromPlan.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the plan.
		toState.InviteRecipientWorkspaceId = fromPlan.InviteRecipientWorkspaceId
	}
}

func (toState *CleanRoomCollaborator) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomCollaborator) {
	if !fromState.InviteRecipientWorkspaceId.IsUnknown() && !fromState.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.InviteRecipientWorkspaceId = fromState.InviteRecipientWorkspaceId
	}
}

func (c CleanRoomCollaborator) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["collaborator_alias"] = attrs["collaborator_alias"].SetRequired()
	attrs["collaborator_alias"] = attrs["collaborator_alias"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["display_name"] = attrs["display_name"].SetComputed()
	attrs["global_metastore_id"] = attrs["global_metastore_id"].SetOptional()
	attrs["invite_recipient_email"] = attrs["invite_recipient_email"].SetOptional()
	attrs["invite_recipient_email"] = attrs["invite_recipient_email"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["invite_recipient_workspace_id"] = attrs["invite_recipient_workspace_id"].SetOptional()
	attrs["invite_recipient_workspace_id"] = attrs["invite_recipient_workspace_id"].SetComputed()
	attrs["invite_recipient_workspace_id"] = attrs["invite_recipient_workspace_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["organization_name"] = attrs["organization_name"].SetComputed()

	return attrs
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

type CleanRoomNotebookReview struct {
	// review comment
	Comment types.String `tfsdk:"comment"`
	// timestamp of when the review was submitted
	CreatedAtMillis types.Int64 `tfsdk:"created_at_millis"`
	// review outcome
	ReviewState types.String `tfsdk:"review_state"`
	// specified when the review was not explicitly made by a user
	ReviewSubReason types.String `tfsdk:"review_sub_reason"`
	// collaborator alias of the reviewer
	ReviewerCollaboratorAlias types.String `tfsdk:"reviewer_collaborator_alias"`
}

func (toState *CleanRoomNotebookReview) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomNotebookReview) {
}

func (toState *CleanRoomNotebookReview) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomNotebookReview) {
}

func (c CleanRoomNotebookReview) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at_millis"] = attrs["created_at_millis"].SetOptional()
	attrs["review_state"] = attrs["review_state"].SetOptional()
	attrs["review_sub_reason"] = attrs["review_sub_reason"].SetOptional()
	attrs["reviewer_collaborator_alias"] = attrs["reviewer_collaborator_alias"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomNotebookReview.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomNotebookReview) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookReview
// only implements ToObjectValue() and Type().
func (o CleanRoomNotebookReview) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                     o.Comment,
			"created_at_millis":           o.CreatedAtMillis,
			"review_state":                o.ReviewState,
			"review_sub_reason":           o.ReviewSubReason,
			"reviewer_collaborator_alias": o.ReviewerCollaboratorAlias,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomNotebookReview) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":                     types.StringType,
			"created_at_millis":           types.Int64Type,
			"review_state":                types.StringType,
			"review_sub_reason":           types.StringType,
			"reviewer_collaborator_alias": types.StringType,
		},
	}
}

// Stores information about a single task run.
type CleanRoomNotebookTaskRun struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo types.Object `tfsdk:"collaborator_job_run_info"`
	// Etag of the notebook executed in this task run, used to identify the
	// notebook version.
	NotebookEtag types.String `tfsdk:"notebook_etag"`
	// State of the task run.
	NotebookJobRunState types.Object `tfsdk:"notebook_job_run_state"`
	// Asset name of the notebook executed in this task run.
	NotebookName types.String `tfsdk:"notebook_name"`
	// The timestamp of when the notebook was last updated.
	NotebookUpdatedAt types.Int64 `tfsdk:"notebook_updated_at"`
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	OutputSchemaExpirationTime types.Int64 `tfsdk:"output_schema_expiration_time"`
	// Name of the output schema associated with the clean rooms notebook task
	// run.
	OutputSchemaName types.String `tfsdk:"output_schema_name"`
	// Duration of the task run, in milliseconds.
	RunDuration types.Int64 `tfsdk:"run_duration"`
	// When the task run started, in epoch milliseconds.
	StartTime types.Int64 `tfsdk:"start_time"`
}

func (toState *CleanRoomNotebookTaskRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomNotebookTaskRun) {
	if !fromPlan.CollaboratorJobRunInfo.IsNull() && !fromPlan.CollaboratorJobRunInfo.IsUnknown() {
		if toStateCollaboratorJobRunInfo, ok := toState.GetCollaboratorJobRunInfo(ctx); ok {
			if fromPlanCollaboratorJobRunInfo, ok := fromPlan.GetCollaboratorJobRunInfo(ctx); ok {
				toStateCollaboratorJobRunInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCollaboratorJobRunInfo)
				toState.SetCollaboratorJobRunInfo(ctx, toStateCollaboratorJobRunInfo)
			}
		}
	}
	if !fromPlan.NotebookJobRunState.IsNull() && !fromPlan.NotebookJobRunState.IsUnknown() {
		if toStateNotebookJobRunState, ok := toState.GetNotebookJobRunState(ctx); ok {
			if fromPlanNotebookJobRunState, ok := fromPlan.GetNotebookJobRunState(ctx); ok {
				toStateNotebookJobRunState.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanNotebookJobRunState)
				toState.SetNotebookJobRunState(ctx, toStateNotebookJobRunState)
			}
		}
	}
}

func (toState *CleanRoomNotebookTaskRun) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomNotebookTaskRun) {
	if !fromState.CollaboratorJobRunInfo.IsNull() && !fromState.CollaboratorJobRunInfo.IsUnknown() {
		if toStateCollaboratorJobRunInfo, ok := toState.GetCollaboratorJobRunInfo(ctx); ok {
			if fromStateCollaboratorJobRunInfo, ok := fromState.GetCollaboratorJobRunInfo(ctx); ok {
				toStateCollaboratorJobRunInfo.SyncFieldsDuringRead(ctx, fromStateCollaboratorJobRunInfo)
				toState.SetCollaboratorJobRunInfo(ctx, toStateCollaboratorJobRunInfo)
			}
		}
	}
	if !fromState.NotebookJobRunState.IsNull() && !fromState.NotebookJobRunState.IsUnknown() {
		if toStateNotebookJobRunState, ok := toState.GetNotebookJobRunState(ctx); ok {
			if fromStateNotebookJobRunState, ok := fromState.GetNotebookJobRunState(ctx); ok {
				toStateNotebookJobRunState.SyncFieldsDuringRead(ctx, fromStateNotebookJobRunState)
				toState.SetNotebookJobRunState(ctx, toStateNotebookJobRunState)
			}
		}
	}
}

func (c CleanRoomNotebookTaskRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["collaborator_job_run_info"] = attrs["collaborator_job_run_info"].SetOptional()
	attrs["notebook_etag"] = attrs["notebook_etag"].SetOptional()
	attrs["notebook_job_run_state"] = attrs["notebook_job_run_state"].SetOptional()
	attrs["notebook_name"] = attrs["notebook_name"].SetOptional()
	attrs["notebook_updated_at"] = attrs["notebook_updated_at"].SetOptional()
	attrs["output_schema_expiration_time"] = attrs["output_schema_expiration_time"].SetOptional()
	attrs["output_schema_name"] = attrs["output_schema_name"].SetOptional()
	attrs["run_duration"] = attrs["run_duration"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()

	return attrs
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
			"notebook_etag":                 o.NotebookEtag,
			"notebook_job_run_state":        o.NotebookJobRunState,
			"notebook_name":                 o.NotebookName,
			"notebook_updated_at":           o.NotebookUpdatedAt,
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
			"collaborator_job_run_info":     CollaboratorJobRunInfo{}.Type(ctx),
			"notebook_etag":                 types.StringType,
			"notebook_job_run_state":        jobs_tf.CleanRoomTaskRunState{}.Type(ctx),
			"notebook_name":                 types.StringType,
			"notebook_updated_at":           types.Int64Type,
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
	var v CollaboratorJobRunInfo
	d := o.CollaboratorJobRunInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaboratorJobRunInfo sets the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun.
func (o *CleanRoomNotebookTaskRun) SetCollaboratorJobRunInfo(ctx context.Context, v CollaboratorJobRunInfo) {
	vs := v.ToObjectValue(ctx)
	o.CollaboratorJobRunInfo = vs
}

// GetNotebookJobRunState returns the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun as
// a jobs_tf.CleanRoomTaskRunState value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomNotebookTaskRun) GetNotebookJobRunState(ctx context.Context) (jobs_tf.CleanRoomTaskRunState, bool) {
	var e jobs_tf.CleanRoomTaskRunState
	if o.NotebookJobRunState.IsNull() || o.NotebookJobRunState.IsUnknown() {
		return e, false
	}
	var v jobs_tf.CleanRoomTaskRunState
	d := o.NotebookJobRunState.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookJobRunState sets the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun.
func (o *CleanRoomNotebookTaskRun) SetNotebookJobRunState(ctx context.Context, v jobs_tf.CleanRoomTaskRunState) {
	vs := v.ToObjectValue(ctx)
	o.NotebookJobRunState = vs
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName types.String `tfsdk:"catalog_name"`

	Status types.String `tfsdk:"status"`
}

func (toState *CleanRoomOutputCatalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomOutputCatalog) {
}

func (toState *CleanRoomOutputCatalog) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomOutputCatalog) {
}

func (c CleanRoomOutputCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["status"] = attrs["status"].SetComputed()

	return attrs
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
	CentralCleanRoomId types.String `tfsdk:"central_clean_room_id"`
	// Cloud vendor (aws,azure,gcp) of the central clean room.
	CloudVendor types.String `tfsdk:"cloud_vendor"`
	// Collaborators in the central clean room. There should one and only one
	// collaborator in the list that satisfies the owner condition:
	//
	// 1. It has the creator's global_metastore_id (determined by caller of
	// CreateCleanRoom).
	//
	// 2. Its invite_recipient_email is empty.
	Collaborators types.List `tfsdk:"collaborators"`

	ComplianceSecurityProfile types.Object `tfsdk:"compliance_security_profile"`
	// Collaborator who creates the clean room.
	Creator types.Object `tfsdk:"creator"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy types.Object `tfsdk:"egress_network_policy"`
	// Region of the central clean room.
	Region types.String `tfsdk:"region"`
}

func (toState *CleanRoomRemoteDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomRemoteDetail) {
	if !fromPlan.ComplianceSecurityProfile.IsNull() && !fromPlan.ComplianceSecurityProfile.IsUnknown() {
		if toStateComplianceSecurityProfile, ok := toState.GetComplianceSecurityProfile(ctx); ok {
			if fromPlanComplianceSecurityProfile, ok := fromPlan.GetComplianceSecurityProfile(ctx); ok {
				toStateComplianceSecurityProfile.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanComplianceSecurityProfile)
				toState.SetComplianceSecurityProfile(ctx, toStateComplianceSecurityProfile)
			}
		}
	}
	if !fromPlan.Creator.IsNull() && !fromPlan.Creator.IsUnknown() {
		if toStateCreator, ok := toState.GetCreator(ctx); ok {
			if fromPlanCreator, ok := fromPlan.GetCreator(ctx); ok {
				toStateCreator.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCreator)
				toState.SetCreator(ctx, toStateCreator)
			}
		}
	}
	if !fromPlan.EgressNetworkPolicy.IsNull() && !fromPlan.EgressNetworkPolicy.IsUnknown() {
		if toStateEgressNetworkPolicy, ok := toState.GetEgressNetworkPolicy(ctx); ok {
			if fromPlanEgressNetworkPolicy, ok := fromPlan.GetEgressNetworkPolicy(ctx); ok {
				toStateEgressNetworkPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanEgressNetworkPolicy)
				toState.SetEgressNetworkPolicy(ctx, toStateEgressNetworkPolicy)
			}
		}
	}
}

func (toState *CleanRoomRemoteDetail) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomRemoteDetail) {
	if !fromState.ComplianceSecurityProfile.IsNull() && !fromState.ComplianceSecurityProfile.IsUnknown() {
		if toStateComplianceSecurityProfile, ok := toState.GetComplianceSecurityProfile(ctx); ok {
			if fromStateComplianceSecurityProfile, ok := fromState.GetComplianceSecurityProfile(ctx); ok {
				toStateComplianceSecurityProfile.SyncFieldsDuringRead(ctx, fromStateComplianceSecurityProfile)
				toState.SetComplianceSecurityProfile(ctx, toStateComplianceSecurityProfile)
			}
		}
	}
	if !fromState.Creator.IsNull() && !fromState.Creator.IsUnknown() {
		if toStateCreator, ok := toState.GetCreator(ctx); ok {
			if fromStateCreator, ok := fromState.GetCreator(ctx); ok {
				toStateCreator.SyncFieldsDuringRead(ctx, fromStateCreator)
				toState.SetCreator(ctx, toStateCreator)
			}
		}
	}
	if !fromState.EgressNetworkPolicy.IsNull() && !fromState.EgressNetworkPolicy.IsUnknown() {
		if toStateEgressNetworkPolicy, ok := toState.GetEgressNetworkPolicy(ctx); ok {
			if fromStateEgressNetworkPolicy, ok := fromState.GetEgressNetworkPolicy(ctx); ok {
				toStateEgressNetworkPolicy.SyncFieldsDuringRead(ctx, fromStateEgressNetworkPolicy)
				toState.SetEgressNetworkPolicy(ctx, toStateEgressNetworkPolicy)
			}
		}
	}
}

func (c CleanRoomRemoteDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["central_clean_room_id"] = attrs["central_clean_room_id"].SetComputed()
	attrs["cloud_vendor"] = attrs["cloud_vendor"].SetOptional()
	attrs["cloud_vendor"] = attrs["cloud_vendor"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["collaborators"] = attrs["collaborators"].SetOptional()
	attrs["collaborators"] = attrs["collaborators"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["compliance_security_profile"] = attrs["compliance_security_profile"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["egress_network_policy"] = attrs["egress_network_policy"].SetOptional()
	attrs["egress_network_policy"] = attrs["egress_network_policy"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["region"] = attrs["region"].SetOptional()
	attrs["region"] = attrs["region"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
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
			"compliance_security_profile": ComplianceSecurityProfile{}.Type(ctx),
			"creator":                     CleanRoomCollaborator{}.Type(ctx),
			"egress_network_policy":       settings_tf.EgressNetworkPolicy{}.Type(ctx),
			"region":                      types.StringType,
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
	var v ComplianceSecurityProfile
	d := o.ComplianceSecurityProfile.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceSecurityProfile sets the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetComplianceSecurityProfile(ctx context.Context, v ComplianceSecurityProfile) {
	vs := v.ToObjectValue(ctx)
	o.ComplianceSecurityProfile = vs
}

// GetCreator returns the value of the Creator field in CleanRoomRemoteDetail as
// a CleanRoomCollaborator value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail) GetCreator(ctx context.Context) (CleanRoomCollaborator, bool) {
	var e CleanRoomCollaborator
	if o.Creator.IsNull() || o.Creator.IsUnknown() {
		return e, false
	}
	var v CleanRoomCollaborator
	d := o.Creator.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCreator sets the value of the Creator field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetCreator(ctx context.Context, v CleanRoomCollaborator) {
	vs := v.ToObjectValue(ctx)
	o.Creator = vs
}

// GetEgressNetworkPolicy returns the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail as
// a settings_tf.EgressNetworkPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail) GetEgressNetworkPolicy(ctx context.Context) (settings_tf.EgressNetworkPolicy, bool) {
	var e settings_tf.EgressNetworkPolicy
	if o.EgressNetworkPolicy.IsNull() || o.EgressNetworkPolicy.IsUnknown() {
		return e, false
	}
	var v settings_tf.EgressNetworkPolicy
	d := o.EgressNetworkPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEgressNetworkPolicy sets the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail.
func (o *CleanRoomRemoteDetail) SetEgressNetworkPolicy(ctx context.Context, v settings_tf.EgressNetworkPolicy) {
	vs := v.ToObjectValue(ctx)
	o.EgressNetworkPolicy = vs
}

type CollaboratorJobRunInfo struct {
	// Alias of the collaborator that triggered the task run.
	CollaboratorAlias types.String `tfsdk:"collaborator_alias"`
	// Job ID of the task run in the collaborator's workspace.
	CollaboratorJobId types.Int64 `tfsdk:"collaborator_job_id"`
	// Job run ID of the task run in the collaborator's workspace.
	CollaboratorJobRunId types.Int64 `tfsdk:"collaborator_job_run_id"`
	// Task run ID of the task run in the collaborator's workspace.
	CollaboratorTaskRunId types.Int64 `tfsdk:"collaborator_task_run_id"`
	// ID of the collaborator's workspace that triggered the task run.
	CollaboratorWorkspaceId types.Int64 `tfsdk:"collaborator_workspace_id"`
}

func (toState *CollaboratorJobRunInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CollaboratorJobRunInfo) {
}

func (toState *CollaboratorJobRunInfo) SyncFieldsDuringRead(ctx context.Context, fromState CollaboratorJobRunInfo) {
}

func (c CollaboratorJobRunInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["collaborator_alias"] = attrs["collaborator_alias"].SetOptional()
	attrs["collaborator_job_id"] = attrs["collaborator_job_id"].SetOptional()
	attrs["collaborator_job_run_id"] = attrs["collaborator_job_run_id"].SetOptional()
	attrs["collaborator_task_run_id"] = attrs["collaborator_task_run_id"].SetOptional()
	attrs["collaborator_workspace_id"] = attrs["collaborator_workspace_id"].SetOptional()

	return attrs
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
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Whether the compliance security profile is enabled.
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (toState *ComplianceSecurityProfile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ComplianceSecurityProfile) {
}

func (toState *ComplianceSecurityProfile) SyncFieldsDuringRead(ctx context.Context, fromState ComplianceSecurityProfile) {
}

func (c ComplianceSecurityProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compliance_standards"] = attrs["compliance_standards"].SetOptional()
	attrs["is_enabled"] = attrs["is_enabled"].SetOptional()

	return attrs
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
// a slice of types.String values.
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

type CreateCleanRoomAssetRequest struct {
	Asset types.Object `tfsdk:"asset"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName types.String `tfsdk:"-"`
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
			"asset":           CleanRoomAsset{}.Type(ctx),
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
	var v CleanRoomAsset
	d := o.Asset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAsset sets the value of the Asset field in CreateCleanRoomAssetRequest.
func (o *CreateCleanRoomAssetRequest) SetAsset(ctx context.Context, v CleanRoomAsset) {
	vs := v.ToObjectValue(ctx)
	o.Asset = vs
}

type CreateCleanRoomAssetReviewRequest struct {
	// can only be NOTEBOOK_FILE for now
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room
	CleanRoomName types.String `tfsdk:"-"`
	// Name of the asset
	Name types.String `tfsdk:"-"`

	NotebookReview types.Object `tfsdk:"notebook_review"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAssetReviewRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAssetReviewRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_review": reflect.TypeOf(NotebookVersionReview{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewRequest
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAssetReviewRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
			"notebook_review": o.NotebookReview,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAssetReviewRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
			"notebook_review": NotebookVersionReview{}.Type(ctx),
		},
	}
}

// GetNotebookReview returns the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest as
// a NotebookVersionReview value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAssetReviewRequest) GetNotebookReview(ctx context.Context) (NotebookVersionReview, bool) {
	var e NotebookVersionReview
	if o.NotebookReview.IsNull() || o.NotebookReview.IsUnknown() {
		return e, false
	}
	var v NotebookVersionReview
	d := o.NotebookReview.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookReview sets the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest.
func (o *CreateCleanRoomAssetReviewRequest) SetNotebookReview(ctx context.Context, v NotebookVersionReview) {
	vs := v.ToObjectValue(ctx)
	o.NotebookReview = vs
}

type CreateCleanRoomAssetReviewResponse struct {
	// top-level status derived from all reviews
	NotebookReviewState types.String `tfsdk:"notebook_review_state"`
	// All existing notebook approvals or rejections
	NotebookReviews types.List `tfsdk:"notebook_reviews"`
}

func (toState *CreateCleanRoomAssetReviewResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomAssetReviewResponse) {
}

func (toState *CreateCleanRoomAssetReviewResponse) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomAssetReviewResponse) {
}

func (c CreateCleanRoomAssetReviewResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["notebook_review_state"] = attrs["notebook_review_state"].SetOptional()
	attrs["notebook_reviews"] = attrs["notebook_reviews"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAssetReviewResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAssetReviewResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_reviews": reflect.TypeOf(CleanRoomNotebookReview{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewResponse
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAssetReviewResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebook_review_state": o.NotebookReviewState,
			"notebook_reviews":      o.NotebookReviews,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAssetReviewResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notebook_review_state": types.StringType,
			"notebook_reviews": basetypes.ListType{
				ElemType: CleanRoomNotebookReview{}.Type(ctx),
			},
		},
	}
}

// GetNotebookReviews returns the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse as
// a slice of CleanRoomNotebookReview values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAssetReviewResponse) GetNotebookReviews(ctx context.Context) ([]CleanRoomNotebookReview, bool) {
	if o.NotebookReviews.IsNull() || o.NotebookReviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview
	d := o.NotebookReviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookReviews sets the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse.
func (o *CreateCleanRoomAssetReviewResponse) SetNotebookReviews(ctx context.Context, v []CleanRoomNotebookReview) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookReviews = types.ListValueMust(t, vs)
}

type CreateCleanRoomAutoApprovalRuleRequest struct {
	AutoApprovalRule types.Object `tfsdk:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": o.AutoApprovalRule,
			"clean_room_name":    o.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_approval_rule": CleanRoomAutoApprovalRule{}.Type(ctx),
			"clean_room_name":    types.StringType,
		},
	}
}

// GetAutoApprovalRule returns the value of the AutoApprovalRule field in CreateCleanRoomAutoApprovalRuleRequest as
// a CleanRoomAutoApprovalRule value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAutoApprovalRuleRequest) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule, bool) {
	var e CleanRoomAutoApprovalRule
	if o.AutoApprovalRule.IsNull() || o.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v CleanRoomAutoApprovalRule
	d := o.AutoApprovalRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in CreateCleanRoomAutoApprovalRuleRequest.
func (o *CreateCleanRoomAutoApprovalRuleRequest) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule) {
	vs := v.ToObjectValue(ctx)
	o.AutoApprovalRule = vs
}

type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`

	OutputCatalog types.Object `tfsdk:"output_catalog"`
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
			"output_catalog":  CleanRoomOutputCatalog{}.Type(ctx),
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
	var v CleanRoomOutputCatalog
	d := o.OutputCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest.
func (o *CreateCleanRoomOutputCatalogRequest) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := v.ToObjectValue(ctx)
	o.OutputCatalog = vs
}

type CreateCleanRoomOutputCatalogResponse struct {
	OutputCatalog types.Object `tfsdk:"output_catalog"`
}

func (toState *CreateCleanRoomOutputCatalogResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomOutputCatalogResponse) {
	if !fromPlan.OutputCatalog.IsNull() && !fromPlan.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromPlanOutputCatalog, ok := fromPlan.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
}

func (toState *CreateCleanRoomOutputCatalogResponse) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomOutputCatalogResponse) {
	if !fromState.OutputCatalog.IsNull() && !fromState.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromStateOutputCatalog, ok := fromState.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringRead(ctx, fromStateOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
}

func (c CreateCleanRoomOutputCatalogResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["output_catalog"] = attrs["output_catalog"].SetOptional()

	return attrs
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
			"output_catalog": CleanRoomOutputCatalog{}.Type(ctx),
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
	var v CleanRoomOutputCatalog
	d := o.OutputCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse.
func (o *CreateCleanRoomOutputCatalogResponse) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := v.ToObjectValue(ctx)
	o.OutputCatalog = vs
}

type CreateCleanRoomRequest struct {
	CleanRoom types.Object `tfsdk:"clean_room"`
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
			"clean_room": CleanRoom{}.Type(ctx),
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
	var v CleanRoom
	d := o.CleanRoom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRoom sets the value of the CleanRoom field in CreateCleanRoomRequest.
func (o *CreateCleanRoomRequest) SetCleanRoom(ctx context.Context, v CleanRoom) {
	vs := v.ToObjectValue(ctx)
	o.CleanRoom = vs
}

type DeleteCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name types.String `tfsdk:"-"`
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
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
		},
	}
}

// Response for delete clean room request. Using an empty message since the
// generic Empty proto does not externd UnshadedMessageMarker.
type DeleteCleanRoomAssetResponse struct {
}

func (toState *DeleteCleanRoomAssetResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCleanRoomAssetResponse) {
}

func (toState *DeleteCleanRoomAssetResponse) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCleanRoomAssetResponse) {
}

func (c DeleteCleanRoomAssetResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

type DeleteCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName types.String `tfsdk:"-"`

	RuleId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"rule_id":         o.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"rule_id":         types.StringType,
		},
	}
}

type DeleteCleanRoomRequest struct {
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
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

type GetCleanRoomAssetRequest struct {
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name types.String `tfsdk:"-"`
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
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
		},
	}
}

type GetCleanRoomAssetRevisionRequest struct {
	// Asset type. Only NOTEBOOK_FILE is supported.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Revision etag to fetch. If not provided, the latest revision will be
	// returned.
	Etag types.String `tfsdk:"-"`
	// Name of the asset.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomAssetRevisionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomAssetRevisionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRevisionRequest
// only implements ToObjectValue() and Type().
func (o GetCleanRoomAssetRevisionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"etag":            o.Etag,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomAssetRevisionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"etag":            types.StringType,
			"name":            types.StringType,
		},
	}
}

type GetCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName types.String `tfsdk:"-"`

	RuleId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (o GetCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"rule_id":         o.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"rule_id":         types.StringType,
		},
	}
}

type GetCleanRoomRequest struct {
	Name types.String `tfsdk:"-"`
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

type ListCleanRoomAssetRevisionsRequest struct {
	// Asset type. Only NOTEBOOK_FILE is supported.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Name of the asset.
	Name types.String `tfsdk:"-"`
	// Maximum number of asset revisions to return. Defaults to 10.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on the previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAssetRevisionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAssetRevisionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsRequest
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetRevisionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
			"page_size":       o.PageSize,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetRevisionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomAssetRevisionsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Revisions types.List `tfsdk:"revisions"`
}

func (toState *ListCleanRoomAssetRevisionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAssetRevisionsResponse) {
}

func (toState *ListCleanRoomAssetRevisionsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAssetRevisionsResponse) {
}

func (c ListCleanRoomAssetRevisionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["revisions"] = attrs["revisions"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAssetRevisionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAssetRevisionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"revisions": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsResponse
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetRevisionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"revisions":       o.Revisions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetRevisionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"revisions": basetypes.ListType{
				ElemType: CleanRoomAsset{}.Type(ctx),
			},
		},
	}
}

// GetRevisions returns the value of the Revisions field in ListCleanRoomAssetRevisionsResponse as
// a slice of CleanRoomAsset values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomAssetRevisionsResponse) GetRevisions(ctx context.Context) ([]CleanRoomAsset, bool) {
	if o.Revisions.IsNull() || o.Revisions.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset
	d := o.Revisions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRevisions sets the value of the Revisions field in ListCleanRoomAssetRevisionsResponse.
func (o *ListCleanRoomAssetRevisionsResponse) SetRevisions(ctx context.Context, v []CleanRoomAsset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["revisions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Revisions = types.ListValueMust(t, vs)
}

type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
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
	Assets types.List `tfsdk:"assets"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListCleanRoomAssetsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAssetsResponse) {
}

func (toState *ListCleanRoomAssetsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAssetsResponse) {
}

func (c ListCleanRoomAssetsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["assets"] = attrs["assets"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

type ListCleanRoomAutoApprovalRulesRequest struct {
	CleanRoomName types.String `tfsdk:"-"`
	// Maximum number of auto-approval rules to return. Defaults to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAutoApprovalRulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAutoApprovalRulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesRequest
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAutoApprovalRulesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"page_size":       o.PageSize,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAutoApprovalRulesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomAutoApprovalRulesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`

	Rules types.List `tfsdk:"rules"`
}

func (toState *ListCleanRoomAutoApprovalRulesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAutoApprovalRulesResponse) {
}

func (toState *ListCleanRoomAutoApprovalRulesResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAutoApprovalRulesResponse) {
}

func (c ListCleanRoomAutoApprovalRulesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["rules"] = attrs["rules"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAutoApprovalRulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAutoApprovalRulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(CleanRoomAutoApprovalRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesResponse
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAutoApprovalRulesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"rules":           o.Rules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAutoApprovalRulesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"rules": basetypes.ListType{
				ElemType: CleanRoomAutoApprovalRule{}.Type(ctx),
			},
		},
	}
}

// GetRules returns the value of the Rules field in ListCleanRoomAutoApprovalRulesResponse as
// a slice of CleanRoomAutoApprovalRule values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomAutoApprovalRulesResponse) GetRules(ctx context.Context) ([]CleanRoomAutoApprovalRule, bool) {
	if o.Rules.IsNull() || o.Rules.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAutoApprovalRule
	d := o.Rules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRules sets the value of the Rules field in ListCleanRoomAutoApprovalRulesResponse.
func (o *ListCleanRoomAutoApprovalRulesResponse) SetRules(ctx context.Context, v []CleanRoomAutoApprovalRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Rules = types.ListValueMust(t, vs)
}

type ListCleanRoomNotebookTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Notebook name
	NotebookName types.String `tfsdk:"-"`
	// The maximum number of task runs to return. Currently ignored - all runs
	// will be returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Name of the clean room.
	Runs types.List `tfsdk:"runs"`
}

func (toState *ListCleanRoomNotebookTaskRunsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomNotebookTaskRunsResponse) {
}

func (toState *ListCleanRoomNotebookTaskRunsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomNotebookTaskRunsResponse) {
}

func (c ListCleanRoomNotebookTaskRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["runs"] = attrs["runs"].SetOptional()

	return attrs
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

type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
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
	CleanRooms types.List `tfsdk:"clean_rooms"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListCleanRoomsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomsResponse) {
}

func (toState *ListCleanRoomsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomsResponse) {
}

func (c ListCleanRoomsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_rooms"] = attrs["clean_rooms"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

type NotebookVersionReview struct {
	// review comment
	Comment types.String `tfsdk:"comment"`
	// etag that identifies the notebook version
	Etag types.String `tfsdk:"etag"`
	// review outcome
	ReviewState types.String `tfsdk:"review_state"`
}

func (toState *NotebookVersionReview) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NotebookVersionReview) {
}

func (toState *NotebookVersionReview) SyncFieldsDuringRead(ctx context.Context, fromState NotebookVersionReview) {
}

func (c NotebookVersionReview) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["etag"] = attrs["etag"].SetRequired()
	attrs["review_state"] = attrs["review_state"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotebookVersionReview.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NotebookVersionReview) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookVersionReview
// only implements ToObjectValue() and Type().
func (o NotebookVersionReview) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      o.Comment,
			"etag":         o.Etag,
			"review_state": o.ReviewState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookVersionReview) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"etag":         types.StringType,
			"review_state": types.StringType,
		},
	}
}

type UpdateCleanRoomAssetRequest struct {
	// The asset to update. The asset's `name` and `asset_type` fields are used
	// to identify the asset to update.
	Asset types.Object `tfsdk:"asset"`
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
			"asset":           CleanRoomAsset{}.Type(ctx),
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
	var v CleanRoomAsset
	d := o.Asset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAsset sets the value of the Asset field in UpdateCleanRoomAssetRequest.
func (o *UpdateCleanRoomAssetRequest) SetAsset(ctx context.Context, v CleanRoomAsset) {
	vs := v.ToObjectValue(ctx)
	o.Asset = vs
}

type UpdateCleanRoomAutoApprovalRuleRequest struct {
	// The auto-approval rule to update. The rule_id field is used to identify
	// the rule to update.
	AutoApprovalRule types.Object `tfsdk:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"-"`
	// A generated UUID identifying the rule.
	RuleId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (o UpdateCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": o.AutoApprovalRule,
			"clean_room_name":    o.CleanRoomName,
			"rule_id":            o.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_approval_rule": CleanRoomAutoApprovalRule{}.Type(ctx),
			"clean_room_name":    types.StringType,
			"rule_id":            types.StringType,
		},
	}
}

// GetAutoApprovalRule returns the value of the AutoApprovalRule field in UpdateCleanRoomAutoApprovalRuleRequest as
// a CleanRoomAutoApprovalRule value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCleanRoomAutoApprovalRuleRequest) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule, bool) {
	var e CleanRoomAutoApprovalRule
	if o.AutoApprovalRule.IsNull() || o.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v CleanRoomAutoApprovalRule
	d := o.AutoApprovalRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in UpdateCleanRoomAutoApprovalRuleRequest.
func (o *UpdateCleanRoomAutoApprovalRuleRequest) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule) {
	vs := v.ToObjectValue(ctx)
	o.AutoApprovalRule = vs
}

type UpdateCleanRoomRequest struct {
	CleanRoom types.Object `tfsdk:"clean_room"`
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
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
			"clean_room": CleanRoom{}.Type(ctx),
			"name":       types.StringType,
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
	var v CleanRoom
	d := o.CleanRoom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRoom sets the value of the CleanRoom field in UpdateCleanRoomRequest.
func (o *UpdateCleanRoomRequest) SetCleanRoom(ctx context.Context, v CleanRoom) {
	vs := v.ToObjectValue(ctx)
	o.CleanRoom = vs
}
