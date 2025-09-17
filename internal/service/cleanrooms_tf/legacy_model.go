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
	"github.com/databricks/terraform-provider-databricks/internal/service/jobs_tf"     // .tmpl
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf" // .tmpl
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CleanRoom_SdkV2 struct {
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
	OutputCatalog types.List `tfsdk:"output_catalog"`
	// This is the Databricks username of the owner of the local clean room
	// securable for permission management.
	Owner types.String `tfsdk:"owner"`
	// Central clean room details. During creation, users need to specify
	// cloud_vendor, region, and collaborators.global_metastore_id. This field
	// will not be filled in the ListCleanRooms call.
	RemoteDetailedInfo types.List `tfsdk:"remote_detailed_info"`
	// Clean room status.
	Status types.String `tfsdk:"status"`
	// When the clean room was last updated, in epoch milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (toState *CleanRoom_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoom_SdkV2) {
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

func (toState *CleanRoom_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoom_SdkV2) {
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

func (c CleanRoom_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_restricted"] = attrs["access_restricted"].SetComputed()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["local_collaborator_alias"] = attrs["local_collaborator_alias"].SetComputed()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["output_catalog"] = attrs["output_catalog"].SetComputed()
	attrs["output_catalog"] = attrs["output_catalog"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["owner"] = attrs["owner"].SetOptional()
	attrs["remote_detailed_info"] = attrs["remote_detailed_info"].SetOptional()
	attrs["remote_detailed_info"] = attrs["remote_detailed_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a CleanRoom_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog":       reflect.TypeOf(CleanRoomOutputCatalog_SdkV2{}),
		"remote_detailed_info": reflect.TypeOf(CleanRoomRemoteDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoom_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoom_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoom_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_restricted":        types.StringType,
			"comment":                  types.StringType,
			"created_at":               types.Int64Type,
			"local_collaborator_alias": types.StringType,
			"name":                     types.StringType,
			"output_catalog": basetypes.ListType{
				ElemType: CleanRoomOutputCatalog_SdkV2{}.Type(ctx),
			},
			"owner": types.StringType,
			"remote_detailed_info": basetypes.ListType{
				ElemType: CleanRoomRemoteDetail_SdkV2{}.Type(ctx),
			},
			"status":     types.StringType,
			"updated_at": types.Int64Type,
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CleanRoom_SdkV2 as
// a CleanRoomOutputCatalog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoom_SdkV2) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog_SdkV2, bool) {
	var e CleanRoomOutputCatalog_SdkV2
	if o.OutputCatalog.IsNull() || o.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog_SdkV2
	d := o.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CleanRoom_SdkV2.
func (o *CleanRoom_SdkV2) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	o.OutputCatalog = types.ListValueMust(t, vs)
}

// GetRemoteDetailedInfo returns the value of the RemoteDetailedInfo field in CleanRoom_SdkV2 as
// a CleanRoomRemoteDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoom_SdkV2) GetRemoteDetailedInfo(ctx context.Context) (CleanRoomRemoteDetail_SdkV2, bool) {
	var e CleanRoomRemoteDetail_SdkV2
	if o.RemoteDetailedInfo.IsNull() || o.RemoteDetailedInfo.IsUnknown() {
		return e, false
	}
	var v []CleanRoomRemoteDetail_SdkV2
	d := o.RemoteDetailedInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRemoteDetailedInfo sets the value of the RemoteDetailedInfo field in CleanRoom_SdkV2.
func (o *CleanRoom_SdkV2) SetRemoteDetailedInfo(ctx context.Context, v CleanRoomRemoteDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["remote_detailed_info"]
	o.RemoteDetailedInfo = types.ListValueMust(t, vs)
}

// Metadata of the clean room asset
type CleanRoomAsset_SdkV2 struct {
	// When the asset is added to the clean room, in epoch milliseconds.
	AddedAt types.Int64 `tfsdk:"added_at"`
	// The type of the asset.
	AssetType types.String `tfsdk:"asset_type"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName types.String `tfsdk:"clean_room_name"`
	// Foreign table details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTable types.List `tfsdk:"foreign_table"`
	// Local details for a foreign that are only available to its owner. Present
	// if and only if **asset_type** is **FOREIGN_TABLE**
	ForeignTableLocalDetails types.List `tfsdk:"foreign_table_local_details"`
	// A fully qualified name that uniquely identifies the asset within the
	// clean room. This is also the name displayed in the clean room UI.
	//
	// For UC securable assets (tables, volumes, etc.), the format is
	// *shared_catalog*.*shared_schema*.*asset_name*
	//
	// For notebooks, the name is the notebook file name. For jar analyses, the
	// name is the jar analysis name.
	Name types.String `tfsdk:"name"`
	// Notebook details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **NOTEBOOK_FILE**
	Notebook types.List `tfsdk:"notebook"`
	// The alias of the collaborator who owns this asset
	OwnerCollaboratorAlias types.String `tfsdk:"owner_collaborator_alias"`
	// Status of the asset
	Status types.String `tfsdk:"status"`
	// Table details available to all collaborators of the clean room. Present
	// if and only if **asset_type** is **TABLE**
	Table types.List `tfsdk:"table"`
	// Local details for a table that are only available to its owner. Present
	// if and only if **asset_type** is **TABLE**
	TableLocalDetails types.List `tfsdk:"table_local_details"`
	// View details available to all collaborators of the clean room. Present if
	// and only if **asset_type** is **VIEW**
	View types.List `tfsdk:"view"`
	// Local details for a view that are only available to its owner. Present if
	// and only if **asset_type** is **VIEW**
	ViewLocalDetails types.List `tfsdk:"view_local_details"`
	// Local details for a volume that are only available to its owner. Present
	// if and only if **asset_type** is **VOLUME**
	VolumeLocalDetails types.List `tfsdk:"volume_local_details"`
}

func (toState *CleanRoomAsset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAsset_SdkV2) {
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

func (toState *CleanRoomAsset_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAsset_SdkV2) {
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

func (c CleanRoomAsset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["added_at"] = attrs["added_at"].SetComputed()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["clean_room_name"] = attrs["clean_room_name"].SetOptional()
	attrs["foreign_table"] = attrs["foreign_table"].SetOptional()
	attrs["foreign_table"] = attrs["foreign_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["foreign_table_local_details"] = attrs["foreign_table_local_details"].SetOptional()
	attrs["foreign_table_local_details"] = attrs["foreign_table_local_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["notebook"] = attrs["notebook"].SetOptional()
	attrs["notebook"] = attrs["notebook"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["owner_collaborator_alias"] = attrs["owner_collaborator_alias"].SetComputed()
	attrs["status"] = attrs["status"].SetComputed()
	attrs["table"] = attrs["table"].SetOptional()
	attrs["table"] = attrs["table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table_local_details"] = attrs["table_local_details"].SetOptional()
	attrs["table_local_details"] = attrs["table_local_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["view"] = attrs["view"].SetOptional()
	attrs["view"] = attrs["view"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["view_local_details"] = attrs["view_local_details"].SetOptional()
	attrs["view_local_details"] = attrs["view_local_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["volume_local_details"] = attrs["volume_local_details"].SetOptional()
	attrs["volume_local_details"] = attrs["volume_local_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAsset.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CleanRoomAsset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"foreign_table":               reflect.TypeOf(CleanRoomAssetForeignTable_SdkV2{}),
		"foreign_table_local_details": reflect.TypeOf(CleanRoomAssetForeignTableLocalDetails_SdkV2{}),
		"notebook":                    reflect.TypeOf(CleanRoomAssetNotebook_SdkV2{}),
		"table":                       reflect.TypeOf(CleanRoomAssetTable_SdkV2{}),
		"table_local_details":         reflect.TypeOf(CleanRoomAssetTableLocalDetails_SdkV2{}),
		"view":                        reflect.TypeOf(CleanRoomAssetView_SdkV2{}),
		"view_local_details":          reflect.TypeOf(CleanRoomAssetViewLocalDetails_SdkV2{}),
		"volume_local_details":        reflect.TypeOf(CleanRoomAssetVolumeLocalDetails_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAsset_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAsset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomAsset_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"added_at":        types.Int64Type,
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"foreign_table": basetypes.ListType{
				ElemType: CleanRoomAssetForeignTable_SdkV2{}.Type(ctx),
			},
			"foreign_table_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetForeignTableLocalDetails_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"notebook": basetypes.ListType{
				ElemType: CleanRoomAssetNotebook_SdkV2{}.Type(ctx),
			},
			"owner_collaborator_alias": types.StringType,
			"status":                   types.StringType,
			"table": basetypes.ListType{
				ElemType: CleanRoomAssetTable_SdkV2{}.Type(ctx),
			},
			"table_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetTableLocalDetails_SdkV2{}.Type(ctx),
			},
			"view": basetypes.ListType{
				ElemType: CleanRoomAssetView_SdkV2{}.Type(ctx),
			},
			"view_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetViewLocalDetails_SdkV2{}.Type(ctx),
			},
			"volume_local_details": basetypes.ListType{
				ElemType: CleanRoomAssetVolumeLocalDetails_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetForeignTable returns the value of the ForeignTable field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetForeignTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetForeignTable(ctx context.Context) (CleanRoomAssetForeignTable_SdkV2, bool) {
	var e CleanRoomAssetForeignTable_SdkV2
	if o.ForeignTable.IsNull() || o.ForeignTable.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetForeignTable_SdkV2
	d := o.ForeignTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignTable sets the value of the ForeignTable field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetForeignTable(ctx context.Context, v CleanRoomAssetForeignTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_table"]
	o.ForeignTable = types.ListValueMust(t, vs)
}

// GetForeignTableLocalDetails returns the value of the ForeignTableLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetForeignTableLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetForeignTableLocalDetails(ctx context.Context) (CleanRoomAssetForeignTableLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetForeignTableLocalDetails_SdkV2
	if o.ForeignTableLocalDetails.IsNull() || o.ForeignTableLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetForeignTableLocalDetails_SdkV2
	d := o.ForeignTableLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignTableLocalDetails sets the value of the ForeignTableLocalDetails field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetForeignTableLocalDetails(ctx context.Context, v CleanRoomAssetForeignTableLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_table_local_details"]
	o.ForeignTableLocalDetails = types.ListValueMust(t, vs)
}

// GetNotebook returns the value of the Notebook field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetNotebook_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetNotebook(ctx context.Context) (CleanRoomAssetNotebook_SdkV2, bool) {
	var e CleanRoomAssetNotebook_SdkV2
	if o.Notebook.IsNull() || o.Notebook.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetNotebook_SdkV2
	d := o.Notebook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebook sets the value of the Notebook field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetNotebook(ctx context.Context, v CleanRoomAssetNotebook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook"]
	o.Notebook = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetTable(ctx context.Context) (CleanRoomAssetTable_SdkV2, bool) {
	var e CleanRoomAssetTable_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetTable_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetTable(ctx context.Context, v CleanRoomAssetTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

// GetTableLocalDetails returns the value of the TableLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetTableLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetTableLocalDetails(ctx context.Context) (CleanRoomAssetTableLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetTableLocalDetails_SdkV2
	if o.TableLocalDetails.IsNull() || o.TableLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetTableLocalDetails_SdkV2
	d := o.TableLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableLocalDetails sets the value of the TableLocalDetails field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetTableLocalDetails(ctx context.Context, v CleanRoomAssetTableLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table_local_details"]
	o.TableLocalDetails = types.ListValueMust(t, vs)
}

// GetView returns the value of the View field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetView_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetView(ctx context.Context) (CleanRoomAssetView_SdkV2, bool) {
	var e CleanRoomAssetView_SdkV2
	if o.View.IsNull() || o.View.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetView_SdkV2
	d := o.View.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetView sets the value of the View field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetView(ctx context.Context, v CleanRoomAssetView_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["view"]
	o.View = types.ListValueMust(t, vs)
}

// GetViewLocalDetails returns the value of the ViewLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetViewLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetViewLocalDetails(ctx context.Context) (CleanRoomAssetViewLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetViewLocalDetails_SdkV2
	if o.ViewLocalDetails.IsNull() || o.ViewLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetViewLocalDetails_SdkV2
	d := o.ViewLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetViewLocalDetails sets the value of the ViewLocalDetails field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetViewLocalDetails(ctx context.Context, v CleanRoomAssetViewLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["view_local_details"]
	o.ViewLocalDetails = types.ListValueMust(t, vs)
}

// GetVolumeLocalDetails returns the value of the VolumeLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetVolumeLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAsset_SdkV2) GetVolumeLocalDetails(ctx context.Context) (CleanRoomAssetVolumeLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetVolumeLocalDetails_SdkV2
	if o.VolumeLocalDetails.IsNull() || o.VolumeLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetVolumeLocalDetails_SdkV2
	d := o.VolumeLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumeLocalDetails sets the value of the VolumeLocalDetails field in CleanRoomAsset_SdkV2.
func (o *CleanRoomAsset_SdkV2) SetVolumeLocalDetails(ctx context.Context, v CleanRoomAssetVolumeLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["volume_local_details"]
	o.VolumeLocalDetails = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTable_SdkV2 struct {
	// The metadata information of the columns in the foreign table
	Columns types.List `tfsdk:"columns"`
}

func (toState *CleanRoomAssetForeignTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetForeignTable_SdkV2) {
}

func (toState *CleanRoomAssetForeignTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetForeignTable_SdkV2) {
}

func (c CleanRoomAssetForeignTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetForeignTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTable_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetForeignTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetForeignTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: catalog_tf.ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in CleanRoomAssetForeignTable_SdkV2 as
// a slice of catalog_tf.ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetForeignTable_SdkV2) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetForeignTable_SdkV2.
func (o *CleanRoomAssetForeignTable_SdkV2) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTableLocalDetails_SdkV2 struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (toState *CleanRoomAssetForeignTableLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetForeignTableLocalDetails_SdkV2) {
}

func (toState *CleanRoomAssetForeignTableLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetForeignTableLocalDetails_SdkV2) {
}

func (c CleanRoomAssetForeignTableLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetForeignTableLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTableLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetForeignTableLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetForeignTableLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAssetNotebook_SdkV2 struct {
	// Server generated etag that represents the notebook version.
	Etag types.String `tfsdk:"etag"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by :method:workspace/export with the format of **HTML**.
	NotebookContent types.String `tfsdk:"notebook_content"`
	// Top-level status derived from all reviews
	ReviewState types.String `tfsdk:"review_state"`
	// All existing approvals or rejections
	Reviews types.List `tfsdk:"reviews"`
	// Aliases of collaborators that can run the notebook.
	RunnerCollaboratorAliases types.List `tfsdk:"runner_collaborator_aliases"`
}

func (toState *CleanRoomAssetNotebook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetNotebook_SdkV2) {
}

func (toState *CleanRoomAssetNotebook_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetNotebook_SdkV2) {
}

func (c CleanRoomAssetNotebook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetNotebook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"reviews":                     reflect.TypeOf(CleanRoomNotebookReview_SdkV2{}),
		"runner_collaborator_aliases": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetNotebook_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetNotebook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomAssetNotebook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":             types.StringType,
			"notebook_content": types.StringType,
			"review_state":     types.StringType,
			"reviews": basetypes.ListType{
				ElemType: CleanRoomNotebookReview_SdkV2{}.Type(ctx),
			},
			"runner_collaborator_aliases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetReviews returns the value of the Reviews field in CleanRoomAssetNotebook_SdkV2 as
// a slice of CleanRoomNotebookReview_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetNotebook_SdkV2) GetReviews(ctx context.Context) ([]CleanRoomNotebookReview_SdkV2, bool) {
	if o.Reviews.IsNull() || o.Reviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview_SdkV2
	d := o.Reviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReviews sets the value of the Reviews field in CleanRoomAssetNotebook_SdkV2.
func (o *CleanRoomAssetNotebook_SdkV2) SetReviews(ctx context.Context, v []CleanRoomNotebookReview_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Reviews = types.ListValueMust(t, vs)
}

// GetRunnerCollaboratorAliases returns the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetNotebook_SdkV2) GetRunnerCollaboratorAliases(ctx context.Context) ([]types.String, bool) {
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

// SetRunnerCollaboratorAliases sets the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook_SdkV2.
func (o *CleanRoomAssetNotebook_SdkV2) SetRunnerCollaboratorAliases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runner_collaborator_aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RunnerCollaboratorAliases = types.ListValueMust(t, vs)
}

type CleanRoomAssetTable_SdkV2 struct {
	// The metadata information of the columns in the table
	Columns types.List `tfsdk:"columns"`
}

func (toState *CleanRoomAssetTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetTable_SdkV2) {
}

func (toState *CleanRoomAssetTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetTable_SdkV2) {
}

func (c CleanRoomAssetTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTable_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: catalog_tf.ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in CleanRoomAssetTable_SdkV2 as
// a slice of catalog_tf.ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetTable_SdkV2) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetTable_SdkV2.
func (o *CleanRoomAssetTable_SdkV2) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetTableLocalDetails_SdkV2 struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName types.String `tfsdk:"local_name"`
	// Partition filtering specification for a shared table.
	Partitions types.List `tfsdk:"partitions"`
}

func (toState *CleanRoomAssetTableLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetTableLocalDetails_SdkV2) {
}

func (toState *CleanRoomAssetTableLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetTableLocalDetails_SdkV2) {
}

func (c CleanRoomAssetTableLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetTableLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitions": reflect.TypeOf(sharing_tf.Partition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTableLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetTableLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
			"partitions": o.Partitions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetTableLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
			"partitions": basetypes.ListType{
				ElemType: sharing_tf.Partition_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPartitions returns the value of the Partitions field in CleanRoomAssetTableLocalDetails_SdkV2 as
// a slice of sharing_tf.Partition_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetTableLocalDetails_SdkV2) GetPartitions(ctx context.Context) ([]sharing_tf.Partition_SdkV2, bool) {
	if o.Partitions.IsNull() || o.Partitions.IsUnknown() {
		return nil, false
	}
	var v []sharing_tf.Partition_SdkV2
	d := o.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in CleanRoomAssetTableLocalDetails_SdkV2.
func (o *CleanRoomAssetTableLocalDetails_SdkV2) SetPartitions(ctx context.Context, v []sharing_tf.Partition_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["partitions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Partitions = types.ListValueMust(t, vs)
}

type CleanRoomAssetView_SdkV2 struct {
	// The metadata information of the columns in the view
	Columns types.List `tfsdk:"columns"`
}

func (toState *CleanRoomAssetView_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetView_SdkV2) {
}

func (toState *CleanRoomAssetView_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetView_SdkV2) {
}

func (c CleanRoomAssetView_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetView_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetView_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetView_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetView_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: catalog_tf.ColumnInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetColumns returns the value of the Columns field in CleanRoomAssetView_SdkV2 as
// a slice of catalog_tf.ColumnInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomAssetView_SdkV2) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo_SdkV2, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo_SdkV2
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetView_SdkV2.
func (o *CleanRoomAssetView_SdkV2) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetViewLocalDetails_SdkV2 struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (toState *CleanRoomAssetViewLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetViewLocalDetails_SdkV2) {
}

func (toState *CleanRoomAssetViewLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetViewLocalDetails_SdkV2) {
}

func (c CleanRoomAssetViewLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetViewLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetViewLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetViewLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetViewLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAssetVolumeLocalDetails_SdkV2 struct {
	// The fully qualified name of the volume in its owner's local metastore, in
	// the format of *catalog*.*schema*.*volume_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (toState *CleanRoomAssetVolumeLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAssetVolumeLocalDetails_SdkV2) {
}

func (toState *CleanRoomAssetVolumeLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAssetVolumeLocalDetails_SdkV2) {
}

func (c CleanRoomAssetVolumeLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAssetVolumeLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetVolumeLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAssetVolumeLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": o.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomAssetVolumeLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAutoApprovalRule_SdkV2 struct {
	// Collaborator alias of the author covered by the rule. Only one of
	// `author_collaborator_alias` and `author_scope` can be set.
	AuthorCollaboratorAlias types.String `tfsdk:"author_collaborator_alias"`
	// Scope of authors covered by the rule. Only one of
	// `author_collaborator_alias` and `author_scope` can be set.
	AuthorScope types.String `tfsdk:"author_scope"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"clean_room_name"`
	// Timestamp of when the rule was created, in epoch milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// A generated UUID identifying the rule.
	RuleId types.String `tfsdk:"rule_id"`
	// The owner of the rule to whom the rule applies.
	RuleOwnerCollaboratorAlias types.String `tfsdk:"rule_owner_collaborator_alias"`
	// Collaborator alias of the runner covered by the rule.
	RunnerCollaboratorAlias types.String `tfsdk:"runner_collaborator_alias"`
}

func (toState *CleanRoomAutoApprovalRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomAutoApprovalRule_SdkV2) {
}

func (toState *CleanRoomAutoApprovalRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomAutoApprovalRule_SdkV2) {
}

func (c CleanRoomAutoApprovalRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomAutoApprovalRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAutoApprovalRule_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomAutoApprovalRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomAutoApprovalRule_SdkV2) Type(ctx context.Context) attr.Type {
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
type CleanRoomCollaborator_SdkV2 struct {
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
	// The global Unity Catalog metastore ID of the collaborator. The identifier
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

func (toState *CleanRoomCollaborator_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomCollaborator_SdkV2) {
	if !fromPlan.InviteRecipientWorkspaceId.IsUnknown() && !fromPlan.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the plan.
		toState.InviteRecipientWorkspaceId = fromPlan.InviteRecipientWorkspaceId
	}
}

func (toState *CleanRoomCollaborator_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomCollaborator_SdkV2) {
	if !fromState.InviteRecipientWorkspaceId.IsUnknown() && !fromState.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.InviteRecipientWorkspaceId = fromState.InviteRecipientWorkspaceId
	}
}

func (c CleanRoomCollaborator_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomCollaborator_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomCollaborator_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomCollaborator_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomCollaborator_SdkV2) Type(ctx context.Context) attr.Type {
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

type CleanRoomNotebookReview_SdkV2 struct {
	// Review comment
	Comment types.String `tfsdk:"comment"`
	// When the review was submitted, in epoch milliseconds
	CreatedAtMillis types.Int64 `tfsdk:"created_at_millis"`
	// Review outcome
	ReviewState types.String `tfsdk:"review_state"`
	// Specified when the review was not explicitly made by a user
	ReviewSubReason types.String `tfsdk:"review_sub_reason"`
	// Collaborator alias of the reviewer
	ReviewerCollaboratorAlias types.String `tfsdk:"reviewer_collaborator_alias"`
}

func (toState *CleanRoomNotebookReview_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomNotebookReview_SdkV2) {
}

func (toState *CleanRoomNotebookReview_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomNotebookReview_SdkV2) {
}

func (c CleanRoomNotebookReview_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomNotebookReview_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookReview_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomNotebookReview_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomNotebookReview_SdkV2) Type(ctx context.Context) attr.Type {
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
type CleanRoomNotebookTaskRun_SdkV2 struct {
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API. if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo types.List `tfsdk:"collaborator_job_run_info"`
	// Etag of the notebook executed in this task run, used to identify the
	// notebook version.
	NotebookEtag types.String `tfsdk:"notebook_etag"`
	// State of the task run.
	NotebookJobRunState types.List `tfsdk:"notebook_job_run_state"`
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

func (toState *CleanRoomNotebookTaskRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomNotebookTaskRun_SdkV2) {
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

func (toState *CleanRoomNotebookTaskRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomNotebookTaskRun_SdkV2) {
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

func (c CleanRoomNotebookTaskRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["collaborator_job_run_info"] = attrs["collaborator_job_run_info"].SetOptional()
	attrs["collaborator_job_run_info"] = attrs["collaborator_job_run_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["notebook_etag"] = attrs["notebook_etag"].SetOptional()
	attrs["notebook_job_run_state"] = attrs["notebook_job_run_state"].SetOptional()
	attrs["notebook_job_run_state"] = attrs["notebook_job_run_state"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a CleanRoomNotebookTaskRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"collaborator_job_run_info": reflect.TypeOf(CollaboratorJobRunInfo_SdkV2{}),
		"notebook_job_run_state":    reflect.TypeOf(jobs_tf.CleanRoomTaskRunState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookTaskRun_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomNotebookTaskRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomNotebookTaskRun_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"collaborator_job_run_info": basetypes.ListType{
				ElemType: CollaboratorJobRunInfo_SdkV2{}.Type(ctx),
			},
			"notebook_etag": types.StringType,
			"notebook_job_run_state": basetypes.ListType{
				ElemType: jobs_tf.CleanRoomTaskRunState_SdkV2{}.Type(ctx),
			},
			"notebook_name":                 types.StringType,
			"notebook_updated_at":           types.Int64Type,
			"output_schema_expiration_time": types.Int64Type,
			"output_schema_name":            types.StringType,
			"run_duration":                  types.Int64Type,
			"start_time":                    types.Int64Type,
		},
	}
}

// GetCollaboratorJobRunInfo returns the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun_SdkV2 as
// a CollaboratorJobRunInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomNotebookTaskRun_SdkV2) GetCollaboratorJobRunInfo(ctx context.Context) (CollaboratorJobRunInfo_SdkV2, bool) {
	var e CollaboratorJobRunInfo_SdkV2
	if o.CollaboratorJobRunInfo.IsNull() || o.CollaboratorJobRunInfo.IsUnknown() {
		return e, false
	}
	var v []CollaboratorJobRunInfo_SdkV2
	d := o.CollaboratorJobRunInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCollaboratorJobRunInfo sets the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun_SdkV2.
func (o *CleanRoomNotebookTaskRun_SdkV2) SetCollaboratorJobRunInfo(ctx context.Context, v CollaboratorJobRunInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborator_job_run_info"]
	o.CollaboratorJobRunInfo = types.ListValueMust(t, vs)
}

// GetNotebookJobRunState returns the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun_SdkV2 as
// a jobs_tf.CleanRoomTaskRunState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomNotebookTaskRun_SdkV2) GetNotebookJobRunState(ctx context.Context) (jobs_tf.CleanRoomTaskRunState_SdkV2, bool) {
	var e jobs_tf.CleanRoomTaskRunState_SdkV2
	if o.NotebookJobRunState.IsNull() || o.NotebookJobRunState.IsUnknown() {
		return e, false
	}
	var v []jobs_tf.CleanRoomTaskRunState_SdkV2
	d := o.NotebookJobRunState.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookJobRunState sets the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun_SdkV2.
func (o *CleanRoomNotebookTaskRun_SdkV2) SetNotebookJobRunState(ctx context.Context, v jobs_tf.CleanRoomTaskRunState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_job_run_state"]
	o.NotebookJobRunState = types.ListValueMust(t, vs)
}

type CleanRoomOutputCatalog_SdkV2 struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName types.String `tfsdk:"catalog_name"`

	Status types.String `tfsdk:"status"`
}

func (toState *CleanRoomOutputCatalog_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomOutputCatalog_SdkV2) {
}

func (toState *CleanRoomOutputCatalog_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomOutputCatalog_SdkV2) {
}

func (c CleanRoomOutputCatalog_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CleanRoomOutputCatalog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomOutputCatalog_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomOutputCatalog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": o.CatalogName,
			"status":       o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CleanRoomOutputCatalog_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"status":       types.StringType,
		},
	}
}

// Publicly visible central clean room details.
type CleanRoomRemoteDetail_SdkV2 struct {
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

	ComplianceSecurityProfile types.List `tfsdk:"compliance_security_profile"`
	// Collaborator who creates the clean room.
	Creator types.List `tfsdk:"creator"`
	// Egress network policy to apply to the central clean room workspace.
	EgressNetworkPolicy types.List `tfsdk:"egress_network_policy"`
	// Region of the central clean room.
	Region types.String `tfsdk:"region"`
}

func (toState *CleanRoomRemoteDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CleanRoomRemoteDetail_SdkV2) {
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

func (toState *CleanRoomRemoteDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CleanRoomRemoteDetail_SdkV2) {
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

func (c CleanRoomRemoteDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["central_clean_room_id"] = attrs["central_clean_room_id"].SetComputed()
	attrs["cloud_vendor"] = attrs["cloud_vendor"].SetOptional()
	attrs["cloud_vendor"] = attrs["cloud_vendor"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["collaborators"] = attrs["collaborators"].SetOptional()
	attrs["collaborators"] = attrs["collaborators"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["compliance_security_profile"] = attrs["compliance_security_profile"].SetComputed()
	attrs["compliance_security_profile"] = attrs["compliance_security_profile"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["creator"] = attrs["creator"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["egress_network_policy"] = attrs["egress_network_policy"].SetOptional()
	attrs["egress_network_policy"] = attrs["egress_network_policy"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["egress_network_policy"] = attrs["egress_network_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a CleanRoomRemoteDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"collaborators":               reflect.TypeOf(CleanRoomCollaborator_SdkV2{}),
		"compliance_security_profile": reflect.TypeOf(ComplianceSecurityProfile_SdkV2{}),
		"creator":                     reflect.TypeOf(CleanRoomCollaborator_SdkV2{}),
		"egress_network_policy":       reflect.TypeOf(settings_tf.EgressNetworkPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomRemoteDetail_SdkV2
// only implements ToObjectValue() and Type().
func (o CleanRoomRemoteDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CleanRoomRemoteDetail_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"central_clean_room_id": types.StringType,
			"cloud_vendor":          types.StringType,
			"collaborators": basetypes.ListType{
				ElemType: CleanRoomCollaborator_SdkV2{}.Type(ctx),
			},
			"compliance_security_profile": basetypes.ListType{
				ElemType: ComplianceSecurityProfile_SdkV2{}.Type(ctx),
			},
			"creator": basetypes.ListType{
				ElemType: CleanRoomCollaborator_SdkV2{}.Type(ctx),
			},
			"egress_network_policy": basetypes.ListType{
				ElemType: settings_tf.EgressNetworkPolicy_SdkV2{}.Type(ctx),
			},
			"region": types.StringType,
		},
	}
}

// GetCollaborators returns the value of the Collaborators field in CleanRoomRemoteDetail_SdkV2 as
// a slice of CleanRoomCollaborator_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail_SdkV2) GetCollaborators(ctx context.Context) ([]CleanRoomCollaborator_SdkV2, bool) {
	if o.Collaborators.IsNull() || o.Collaborators.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomCollaborator_SdkV2
	d := o.Collaborators.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaborators sets the value of the Collaborators field in CleanRoomRemoteDetail_SdkV2.
func (o *CleanRoomRemoteDetail_SdkV2) SetCollaborators(ctx context.Context, v []CleanRoomCollaborator_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborators"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Collaborators = types.ListValueMust(t, vs)
}

// GetComplianceSecurityProfile returns the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail_SdkV2 as
// a ComplianceSecurityProfile_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail_SdkV2) GetComplianceSecurityProfile(ctx context.Context) (ComplianceSecurityProfile_SdkV2, bool) {
	var e ComplianceSecurityProfile_SdkV2
	if o.ComplianceSecurityProfile.IsNull() || o.ComplianceSecurityProfile.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfile_SdkV2
	d := o.ComplianceSecurityProfile.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComplianceSecurityProfile sets the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail_SdkV2.
func (o *CleanRoomRemoteDetail_SdkV2) SetComplianceSecurityProfile(ctx context.Context, v ComplianceSecurityProfile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_security_profile"]
	o.ComplianceSecurityProfile = types.ListValueMust(t, vs)
}

// GetCreator returns the value of the Creator field in CleanRoomRemoteDetail_SdkV2 as
// a CleanRoomCollaborator_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail_SdkV2) GetCreator(ctx context.Context) (CleanRoomCollaborator_SdkV2, bool) {
	var e CleanRoomCollaborator_SdkV2
	if o.Creator.IsNull() || o.Creator.IsUnknown() {
		return e, false
	}
	var v []CleanRoomCollaborator_SdkV2
	d := o.Creator.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCreator sets the value of the Creator field in CleanRoomRemoteDetail_SdkV2.
func (o *CleanRoomRemoteDetail_SdkV2) SetCreator(ctx context.Context, v CleanRoomCollaborator_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["creator"]
	o.Creator = types.ListValueMust(t, vs)
}

// GetEgressNetworkPolicy returns the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail_SdkV2 as
// a settings_tf.EgressNetworkPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CleanRoomRemoteDetail_SdkV2) GetEgressNetworkPolicy(ctx context.Context) (settings_tf.EgressNetworkPolicy_SdkV2, bool) {
	var e settings_tf.EgressNetworkPolicy_SdkV2
	if o.EgressNetworkPolicy.IsNull() || o.EgressNetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []settings_tf.EgressNetworkPolicy_SdkV2
	d := o.EgressNetworkPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEgressNetworkPolicy sets the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail_SdkV2.
func (o *CleanRoomRemoteDetail_SdkV2) SetEgressNetworkPolicy(ctx context.Context, v settings_tf.EgressNetworkPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["egress_network_policy"]
	o.EgressNetworkPolicy = types.ListValueMust(t, vs)
}

type CollaboratorJobRunInfo_SdkV2 struct {
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

func (toState *CollaboratorJobRunInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CollaboratorJobRunInfo_SdkV2) {
}

func (toState *CollaboratorJobRunInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CollaboratorJobRunInfo_SdkV2) {
}

func (c CollaboratorJobRunInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CollaboratorJobRunInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CollaboratorJobRunInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o CollaboratorJobRunInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CollaboratorJobRunInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type ComplianceSecurityProfile_SdkV2 struct {
	// The list of compliance standards that the compliance security profile is
	// configured to enforce.
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Whether the compliance security profile is enabled.
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (toState *ComplianceSecurityProfile_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ComplianceSecurityProfile_SdkV2) {
}

func (toState *ComplianceSecurityProfile_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ComplianceSecurityProfile_SdkV2) {
}

func (c ComplianceSecurityProfile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComplianceSecurityProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile_SdkV2
// only implements ToObjectValue() and Type().
func (o ComplianceSecurityProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": o.ComplianceStandards,
			"is_enabled":           o.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplianceSecurityProfile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enabled": types.BoolType,
		},
	}
}

// GetComplianceStandards returns the value of the ComplianceStandards field in ComplianceSecurityProfile_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ComplianceSecurityProfile_SdkV2) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
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

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile_SdkV2.
func (o *ComplianceSecurityProfile_SdkV2) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ComplianceStandards = types.ListValueMust(t, vs)
}

type CreateCleanRoomAssetRequest_SdkV2 struct {
	Asset types.List `tfsdk:"asset"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName types.String `tfsdk:"-"`
}

func (toState *CreateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomAssetRequest_SdkV2) {
	if !fromPlan.Asset.IsNull() && !fromPlan.Asset.IsUnknown() {
		if toStateAsset, ok := toState.GetAsset(ctx); ok {
			if fromPlanAsset, ok := fromPlan.GetAsset(ctx); ok {
				toStateAsset.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAsset)
				toState.SetAsset(ctx, toStateAsset)
			}
		}
	}
}

func (toState *CreateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomAssetRequest_SdkV2) {
	if !fromState.Asset.IsNull() && !fromState.Asset.IsUnknown() {
		if toStateAsset, ok := toState.GetAsset(ctx); ok {
			if fromStateAsset, ok := fromState.GetAsset(ctx); ok {
				toStateAsset.SyncFieldsDuringRead(ctx, fromStateAsset)
				toState.SetAsset(ctx, toStateAsset)
			}
		}
	}
}

func (c CreateCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["asset"] = attrs["asset"].SetRequired()
	attrs["asset"] = attrs["asset"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset":           o.Asset,
			"clean_room_name": o.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset": basetypes.ListType{
				ElemType: CleanRoomAsset_SdkV2{}.Type(ctx),
			},
			"clean_room_name": types.StringType,
		},
	}
}

// GetAsset returns the value of the Asset field in CreateCleanRoomAssetRequest_SdkV2 as
// a CleanRoomAsset_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAssetRequest_SdkV2) GetAsset(ctx context.Context) (CleanRoomAsset_SdkV2, bool) {
	var e CleanRoomAsset_SdkV2
	if o.Asset.IsNull() || o.Asset.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAsset_SdkV2
	d := o.Asset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAsset sets the value of the Asset field in CreateCleanRoomAssetRequest_SdkV2.
func (o *CreateCleanRoomAssetRequest_SdkV2) SetAsset(ctx context.Context, v CleanRoomAsset_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["asset"]
	o.Asset = types.ListValueMust(t, vs)
}

type CreateCleanRoomAssetReviewRequest_SdkV2 struct {
	// Asset type. Can either be NOTEBOOK_FILE or JAR_ANALYSIS.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room
	CleanRoomName types.String `tfsdk:"-"`
	// Name of the asset
	Name types.String `tfsdk:"-"`

	NotebookReview types.List `tfsdk:"notebook_review"`
}

func (toState *CreateCleanRoomAssetReviewRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomAssetReviewRequest_SdkV2) {
	if !fromPlan.NotebookReview.IsNull() && !fromPlan.NotebookReview.IsUnknown() {
		if toStateNotebookReview, ok := toState.GetNotebookReview(ctx); ok {
			if fromPlanNotebookReview, ok := fromPlan.GetNotebookReview(ctx); ok {
				toStateNotebookReview.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanNotebookReview)
				toState.SetNotebookReview(ctx, toStateNotebookReview)
			}
		}
	}
}

func (toState *CreateCleanRoomAssetReviewRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomAssetReviewRequest_SdkV2) {
	if !fromState.NotebookReview.IsNull() && !fromState.NotebookReview.IsUnknown() {
		if toStateNotebookReview, ok := toState.GetNotebookReview(ctx); ok {
			if fromStateNotebookReview, ok := fromState.GetNotebookReview(ctx); ok {
				toStateNotebookReview.SyncFieldsDuringRead(ctx, fromStateNotebookReview)
				toState.SetNotebookReview(ctx, toStateNotebookReview)
			}
		}
	}
}

func (c CreateCleanRoomAssetReviewRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["notebook_review"] = attrs["notebook_review"].SetOptional()
	attrs["notebook_review"] = attrs["notebook_review"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAssetReviewRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAssetReviewRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_review": reflect.TypeOf(NotebookVersionReview_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAssetReviewRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o CreateCleanRoomAssetReviewRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
			"notebook_review": basetypes.ListType{
				ElemType: NotebookVersionReview_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNotebookReview returns the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest_SdkV2 as
// a NotebookVersionReview_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAssetReviewRequest_SdkV2) GetNotebookReview(ctx context.Context) (NotebookVersionReview_SdkV2, bool) {
	var e NotebookVersionReview_SdkV2
	if o.NotebookReview.IsNull() || o.NotebookReview.IsUnknown() {
		return e, false
	}
	var v []NotebookVersionReview_SdkV2
	d := o.NotebookReview.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookReview sets the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest_SdkV2.
func (o *CreateCleanRoomAssetReviewRequest_SdkV2) SetNotebookReview(ctx context.Context, v NotebookVersionReview_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_review"]
	o.NotebookReview = types.ListValueMust(t, vs)
}

type CreateCleanRoomAssetReviewResponse_SdkV2 struct {
	// Top-level status derived from all reviews
	NotebookReviewState types.String `tfsdk:"notebook_review_state"`
	// All existing notebook approvals or rejections
	NotebookReviews types.List `tfsdk:"notebook_reviews"`
}

func (toState *CreateCleanRoomAssetReviewResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomAssetReviewResponse_SdkV2) {
}

func (toState *CreateCleanRoomAssetReviewResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomAssetReviewResponse_SdkV2) {
}

func (c CreateCleanRoomAssetReviewResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateCleanRoomAssetReviewResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_reviews": reflect.TypeOf(CleanRoomNotebookReview_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAssetReviewResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebook_review_state": o.NotebookReviewState,
			"notebook_reviews":      o.NotebookReviews,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAssetReviewResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"notebook_review_state": types.StringType,
			"notebook_reviews": basetypes.ListType{
				ElemType: CleanRoomNotebookReview_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNotebookReviews returns the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse_SdkV2 as
// a slice of CleanRoomNotebookReview_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAssetReviewResponse_SdkV2) GetNotebookReviews(ctx context.Context) ([]CleanRoomNotebookReview_SdkV2, bool) {
	if o.NotebookReviews.IsNull() || o.NotebookReviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview_SdkV2
	d := o.NotebookReviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookReviews sets the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse_SdkV2.
func (o *CreateCleanRoomAssetReviewResponse_SdkV2) SetNotebookReviews(ctx context.Context, v []CleanRoomNotebookReview_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.NotebookReviews = types.ListValueMust(t, vs)
}

type CreateCleanRoomAutoApprovalRuleRequest_SdkV2 struct {
	AutoApprovalRule types.List `tfsdk:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"-"`
}

func (toState *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !fromPlan.AutoApprovalRule.IsNull() && !fromPlan.AutoApprovalRule.IsUnknown() {
		if toStateAutoApprovalRule, ok := toState.GetAutoApprovalRule(ctx); ok {
			if fromPlanAutoApprovalRule, ok := fromPlan.GetAutoApprovalRule(ctx); ok {
				toStateAutoApprovalRule.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAutoApprovalRule)
				toState.SetAutoApprovalRule(ctx, toStateAutoApprovalRule)
			}
		}
	}
}

func (toState *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !fromState.AutoApprovalRule.IsNull() && !fromState.AutoApprovalRule.IsUnknown() {
		if toStateAutoApprovalRule, ok := toState.GetAutoApprovalRule(ctx); ok {
			if fromStateAutoApprovalRule, ok := fromState.GetAutoApprovalRule(ctx); ok {
				toStateAutoApprovalRule.SyncFieldsDuringRead(ctx, fromStateAutoApprovalRule)
				toState.SetAutoApprovalRule(ctx, toStateAutoApprovalRule)
			}
		}
	}
}

func (c CreateCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_approval_rule"] = attrs["auto_approval_rule"].SetRequired()
	attrs["auto_approval_rule"] = attrs["auto_approval_rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": o.AutoApprovalRule,
			"clean_room_name":    o.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_approval_rule": basetypes.ListType{
				ElemType: CleanRoomAutoApprovalRule_SdkV2{}.Type(ctx),
			},
			"clean_room_name": types.StringType,
		},
	}
}

// GetAutoApprovalRule returns the value of the AutoApprovalRule field in CreateCleanRoomAutoApprovalRuleRequest_SdkV2 as
// a CleanRoomAutoApprovalRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule_SdkV2, bool) {
	var e CleanRoomAutoApprovalRule_SdkV2
	if o.AutoApprovalRule.IsNull() || o.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAutoApprovalRule_SdkV2
	d := o.AutoApprovalRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in CreateCleanRoomAutoApprovalRuleRequest_SdkV2.
func (o *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_approval_rule"]
	o.AutoApprovalRule = types.ListValueMust(t, vs)
}

type CreateCleanRoomOutputCatalogRequest_SdkV2 struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`

	OutputCatalog types.List `tfsdk:"output_catalog"`
}

func (toState *CreateCleanRoomOutputCatalogRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomOutputCatalogRequest_SdkV2) {
	if !fromPlan.OutputCatalog.IsNull() && !fromPlan.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromPlanOutputCatalog, ok := fromPlan.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
}

func (toState *CreateCleanRoomOutputCatalogRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomOutputCatalogRequest_SdkV2) {
	if !fromState.OutputCatalog.IsNull() && !fromState.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromStateOutputCatalog, ok := fromState.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringRead(ctx, fromStateOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
}

func (c CreateCleanRoomOutputCatalogRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["output_catalog"] = attrs["output_catalog"].SetRequired()
	attrs["output_catalog"] = attrs["output_catalog"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomOutputCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomOutputCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomOutputCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"output_catalog":  o.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomOutputCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"output_catalog": basetypes.ListType{
				ElemType: CleanRoomOutputCatalog_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest_SdkV2 as
// a CleanRoomOutputCatalog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomOutputCatalogRequest_SdkV2) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog_SdkV2, bool) {
	var e CleanRoomOutputCatalog_SdkV2
	if o.OutputCatalog.IsNull() || o.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog_SdkV2
	d := o.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest_SdkV2.
func (o *CreateCleanRoomOutputCatalogRequest_SdkV2) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	o.OutputCatalog = types.ListValueMust(t, vs)
}

type CreateCleanRoomOutputCatalogResponse_SdkV2 struct {
	OutputCatalog types.List `tfsdk:"output_catalog"`
}

func (toState *CreateCleanRoomOutputCatalogResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomOutputCatalogResponse_SdkV2) {
	if !fromPlan.OutputCatalog.IsNull() && !fromPlan.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromPlanOutputCatalog, ok := fromPlan.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
}

func (toState *CreateCleanRoomOutputCatalogResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomOutputCatalogResponse_SdkV2) {
	if !fromState.OutputCatalog.IsNull() && !fromState.OutputCatalog.IsUnknown() {
		if toStateOutputCatalog, ok := toState.GetOutputCatalog(ctx); ok {
			if fromStateOutputCatalog, ok := fromState.GetOutputCatalog(ctx); ok {
				toStateOutputCatalog.SyncFieldsDuringRead(ctx, fromStateOutputCatalog)
				toState.SetOutputCatalog(ctx, toStateOutputCatalog)
			}
		}
	}
}

func (c CreateCleanRoomOutputCatalogResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["output_catalog"] = attrs["output_catalog"].SetOptional()
	attrs["output_catalog"] = attrs["output_catalog"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomOutputCatalogResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomOutputCatalogResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomOutputCatalogResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"output_catalog": o.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomOutputCatalogResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"output_catalog": basetypes.ListType{
				ElemType: CleanRoomOutputCatalog_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse_SdkV2 as
// a CleanRoomOutputCatalog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomOutputCatalogResponse_SdkV2) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog_SdkV2, bool) {
	var e CleanRoomOutputCatalog_SdkV2
	if o.OutputCatalog.IsNull() || o.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog_SdkV2
	d := o.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse_SdkV2.
func (o *CreateCleanRoomOutputCatalogResponse_SdkV2) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	o.OutputCatalog = types.ListValueMust(t, vs)
}

type CreateCleanRoomRequest_SdkV2 struct {
	CleanRoom types.List `tfsdk:"clean_room"`
}

func (toState *CreateCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan CreateCleanRoomRequest_SdkV2) {
	if !fromPlan.CleanRoom.IsNull() && !fromPlan.CleanRoom.IsUnknown() {
		if toStateCleanRoom, ok := toState.GetCleanRoom(ctx); ok {
			if fromPlanCleanRoom, ok := fromPlan.GetCleanRoom(ctx); ok {
				toStateCleanRoom.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCleanRoom)
				toState.SetCleanRoom(ctx, toStateCleanRoom)
			}
		}
	}
}

func (toState *CreateCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState CreateCleanRoomRequest_SdkV2) {
	if !fromState.CleanRoom.IsNull() && !fromState.CleanRoom.IsUnknown() {
		if toStateCleanRoom, ok := toState.GetCleanRoom(ctx); ok {
			if fromStateCleanRoom, ok := fromState.GetCleanRoom(ctx); ok {
				toStateCleanRoom.SyncFieldsDuringRead(ctx, fromStateCleanRoom)
				toState.SetCleanRoom(ctx, toStateCleanRoom)
			}
		}
	}
}

func (c CreateCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room"] = attrs["clean_room"].SetRequired()
	attrs["clean_room"] = attrs["clean_room"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": o.CleanRoom,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room": basetypes.ListType{
				ElemType: CleanRoom_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCleanRoom returns the value of the CleanRoom field in CreateCleanRoomRequest_SdkV2 as
// a CleanRoom_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCleanRoomRequest_SdkV2) GetCleanRoom(ctx context.Context) (CleanRoom_SdkV2, bool) {
	var e CleanRoom_SdkV2
	if o.CleanRoom.IsNull() || o.CleanRoom.IsUnknown() {
		return e, false
	}
	var v []CleanRoom_SdkV2
	d := o.CleanRoom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoom sets the value of the CleanRoom field in CreateCleanRoomRequest_SdkV2.
func (o *CreateCleanRoomRequest_SdkV2) SetCleanRoom(ctx context.Context, v CleanRoom_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_room"]
	o.CleanRoom = types.ListValueMust(t, vs)
}

type DeleteCleanRoomAssetRequest_SdkV2 struct {
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name types.String `tfsdk:"-"`
}

func (toState *DeleteCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCleanRoomAssetRequest_SdkV2) {
}

func (toState *DeleteCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCleanRoomAssetRequest_SdkV2) {
}

func (c DeleteCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
type DeleteCleanRoomAssetResponse_SdkV2 struct {
}

func (toState *DeleteCleanRoomAssetResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCleanRoomAssetResponse_SdkV2) {
}

func (toState *DeleteCleanRoomAssetResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCleanRoomAssetResponse_SdkV2) {
}

func (c DeleteCleanRoomAssetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAssetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomAssetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomAssetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAssetResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteCleanRoomAutoApprovalRuleRequest_SdkV2 struct {
	CleanRoomName types.String `tfsdk:"-"`

	RuleId types.String `tfsdk:"-"`
}

func (toState *DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (toState *DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (c DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["rule_id"] = attrs["rule_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"rule_id":         o.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"rule_id":         types.StringType,
		},
	}
}

type DeleteCleanRoomRequest_SdkV2 struct {
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (toState *DeleteCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeleteCleanRoomRequest_SdkV2) {
}

func (toState *DeleteCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState DeleteCleanRoomRequest_SdkV2) {
}

func (c DeleteCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetCleanRoomAssetRequest_SdkV2 struct {
	// The type of the asset.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// The fully qualified name of the asset, it is same as the name field in
	// CleanRoomAsset.
	Name types.String `tfsdk:"-"`
}

func (toState *GetCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetCleanRoomAssetRequest_SdkV2) {
}

func (toState *GetCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetCleanRoomAssetRequest_SdkV2) {
}

func (c GetCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      o.AssetType,
			"clean_room_name": o.CleanRoomName,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
		},
	}
}

type GetCleanRoomAssetRevisionRequest_SdkV2 struct {
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

func (toState *GetCleanRoomAssetRevisionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetCleanRoomAssetRevisionRequest_SdkV2) {
}

func (toState *GetCleanRoomAssetRevisionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetCleanRoomAssetRevisionRequest_SdkV2) {
}

func (c GetCleanRoomAssetRevisionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomAssetRevisionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomAssetRevisionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRevisionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCleanRoomAssetRevisionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GetCleanRoomAssetRevisionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"etag":            types.StringType,
			"name":            types.StringType,
		},
	}
}

type GetCleanRoomAutoApprovalRuleRequest_SdkV2 struct {
	CleanRoomName types.String `tfsdk:"-"`

	RuleId types.String `tfsdk:"-"`
}

func (toState *GetCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (toState *GetCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (c GetCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["rule_id"] = attrs["rule_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"rule_id":         o.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"rule_id":         types.StringType,
		},
	}
}

type GetCleanRoomRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (toState *GetCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan GetCleanRoomRequest_SdkV2) {
}

func (toState *GetCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState GetCleanRoomRequest_SdkV2) {
}

func (c GetCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListCleanRoomAssetRevisionsRequest_SdkV2 struct {
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

func (toState *ListCleanRoomAssetRevisionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAssetRevisionsRequest_SdkV2) {
}

func (toState *ListCleanRoomAssetRevisionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAssetRevisionsRequest_SdkV2) {
}

func (c ListCleanRoomAssetRevisionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAssetRevisionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAssetRevisionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetRevisionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListCleanRoomAssetRevisionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

type ListCleanRoomAssetRevisionsResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Revisions types.List `tfsdk:"revisions"`
}

func (toState *ListCleanRoomAssetRevisionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAssetRevisionsResponse_SdkV2) {
}

func (toState *ListCleanRoomAssetRevisionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAssetRevisionsResponse_SdkV2) {
}

func (c ListCleanRoomAssetRevisionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCleanRoomAssetRevisionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"revisions": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetRevisionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"revisions":       o.Revisions,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetRevisionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"revisions": basetypes.ListType{
				ElemType: CleanRoomAsset_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRevisions returns the value of the Revisions field in ListCleanRoomAssetRevisionsResponse_SdkV2 as
// a slice of CleanRoomAsset_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomAssetRevisionsResponse_SdkV2) GetRevisions(ctx context.Context) ([]CleanRoomAsset_SdkV2, bool) {
	if o.Revisions.IsNull() || o.Revisions.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset_SdkV2
	d := o.Revisions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRevisions sets the value of the Revisions field in ListCleanRoomAssetRevisionsResponse_SdkV2.
func (o *ListCleanRoomAssetRevisionsResponse_SdkV2) SetRevisions(ctx context.Context, v []CleanRoomAsset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["revisions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Revisions = types.ListValueMust(t, vs)
}

type ListCleanRoomAssetsRequest_SdkV2 struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (toState *ListCleanRoomAssetsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAssetsRequest_SdkV2) {
}

func (toState *ListCleanRoomAssetsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAssetsRequest_SdkV2) {
}

func (c ListCleanRoomAssetsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAssetsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAssetsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomAssetsResponse_SdkV2 struct {
	// Assets in the clean room.
	Assets types.List `tfsdk:"assets"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListCleanRoomAssetsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAssetsResponse_SdkV2) {
}

func (toState *ListCleanRoomAssetsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAssetsResponse_SdkV2) {
}

func (c ListCleanRoomAssetsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCleanRoomAssetsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAssetsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":          o.Assets,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAssetsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"assets": basetypes.ListType{
				ElemType: CleanRoomAsset_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetAssets returns the value of the Assets field in ListCleanRoomAssetsResponse_SdkV2 as
// a slice of CleanRoomAsset_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomAssetsResponse_SdkV2) GetAssets(ctx context.Context) ([]CleanRoomAsset_SdkV2, bool) {
	if o.Assets.IsNull() || o.Assets.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset_SdkV2
	d := o.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListCleanRoomAssetsResponse_SdkV2.
func (o *ListCleanRoomAssetsResponse_SdkV2) SetAssets(ctx context.Context, v []CleanRoomAsset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Assets = types.ListValueMust(t, vs)
}

type ListCleanRoomAutoApprovalRulesRequest_SdkV2 struct {
	CleanRoomName types.String `tfsdk:"-"`
	// Maximum number of auto-approval rules to return. Defaults to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (toState *ListCleanRoomAutoApprovalRulesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAutoApprovalRulesRequest_SdkV2) {
}

func (toState *ListCleanRoomAutoApprovalRulesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAutoApprovalRulesRequest_SdkV2) {
}

func (c ListCleanRoomAutoApprovalRulesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomAutoApprovalRulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomAutoApprovalRulesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAutoApprovalRulesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": o.CleanRoomName,
			"page_size":       o.PageSize,
			"page_token":      o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAutoApprovalRulesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomAutoApprovalRulesResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`

	Rules types.List `tfsdk:"rules"`
}

func (toState *ListCleanRoomAutoApprovalRulesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomAutoApprovalRulesResponse_SdkV2) {
}

func (toState *ListCleanRoomAutoApprovalRulesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomAutoApprovalRulesResponse_SdkV2) {
}

func (c ListCleanRoomAutoApprovalRulesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCleanRoomAutoApprovalRulesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(CleanRoomAutoApprovalRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomAutoApprovalRulesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"rules":           o.Rules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomAutoApprovalRulesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"rules": basetypes.ListType{
				ElemType: CleanRoomAutoApprovalRule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRules returns the value of the Rules field in ListCleanRoomAutoApprovalRulesResponse_SdkV2 as
// a slice of CleanRoomAutoApprovalRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomAutoApprovalRulesResponse_SdkV2) GetRules(ctx context.Context) ([]CleanRoomAutoApprovalRule_SdkV2, bool) {
	if o.Rules.IsNull() || o.Rules.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAutoApprovalRule_SdkV2
	d := o.Rules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRules sets the value of the Rules field in ListCleanRoomAutoApprovalRulesResponse_SdkV2.
func (o *ListCleanRoomAutoApprovalRulesResponse_SdkV2) SetRules(ctx context.Context, v []CleanRoomAutoApprovalRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Rules = types.ListValueMust(t, vs)
}

type ListCleanRoomNotebookTaskRunsRequest_SdkV2 struct {
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

func (toState *ListCleanRoomNotebookTaskRunsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomNotebookTaskRunsRequest_SdkV2) {
}

func (toState *ListCleanRoomNotebookTaskRunsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomNotebookTaskRunsRequest_SdkV2) {
}

func (c ListCleanRoomNotebookTaskRunsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["notebook_name"] = attrs["notebook_name"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomNotebookTaskRunsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomNotebookTaskRunsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomNotebookTaskRunsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ListCleanRoomNotebookTaskRunsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"notebook_name":   types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
		},
	}
}

type ListCleanRoomNotebookTaskRunsResponse_SdkV2 struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Name of the clean room.
	Runs types.List `tfsdk:"runs"`
}

func (toState *ListCleanRoomNotebookTaskRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomNotebookTaskRunsResponse_SdkV2) {
}

func (toState *ListCleanRoomNotebookTaskRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomNotebookTaskRunsResponse_SdkV2) {
}

func (c ListCleanRoomNotebookTaskRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCleanRoomNotebookTaskRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(CleanRoomNotebookTaskRun_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomNotebookTaskRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"runs":            o.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomNotebookTaskRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: CleanRoomNotebookTaskRun_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRuns returns the value of the Runs field in ListCleanRoomNotebookTaskRunsResponse_SdkV2 as
// a slice of CleanRoomNotebookTaskRun_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomNotebookTaskRunsResponse_SdkV2) GetRuns(ctx context.Context) ([]CleanRoomNotebookTaskRun_SdkV2, bool) {
	if o.Runs.IsNull() || o.Runs.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookTaskRun_SdkV2
	d := o.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in ListCleanRoomNotebookTaskRunsResponse_SdkV2.
func (o *ListCleanRoomNotebookTaskRunsResponse_SdkV2) SetRuns(ctx context.Context, v []CleanRoomNotebookTaskRun_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Runs = types.ListValueMust(t, vs)
}

type ListCleanRoomsRequest_SdkV2 struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (toState *ListCleanRoomsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomsRequest_SdkV2) {
}

func (toState *ListCleanRoomsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomsRequest_SdkV2) {
}

func (c ListCleanRoomsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCleanRoomsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListCleanRoomsResponse_SdkV2 struct {
	CleanRooms types.List `tfsdk:"clean_rooms"`
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (toState *ListCleanRoomsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListCleanRoomsResponse_SdkV2) {
}

func (toState *ListCleanRoomsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState ListCleanRoomsResponse_SdkV2) {
}

func (c ListCleanRoomsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListCleanRoomsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms": reflect.TypeOf(CleanRoom_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListCleanRoomsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_rooms":     o.CleanRooms,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCleanRoomsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_rooms": basetypes.ListType{
				ElemType: CleanRoom_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetCleanRooms returns the value of the CleanRooms field in ListCleanRoomsResponse_SdkV2 as
// a slice of CleanRoom_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListCleanRoomsResponse_SdkV2) GetCleanRooms(ctx context.Context) ([]CleanRoom_SdkV2, bool) {
	if o.CleanRooms.IsNull() || o.CleanRooms.IsUnknown() {
		return nil, false
	}
	var v []CleanRoom_SdkV2
	d := o.CleanRooms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRooms sets the value of the CleanRooms field in ListCleanRoomsResponse_SdkV2.
func (o *ListCleanRoomsResponse_SdkV2) SetCleanRooms(ctx context.Context, v []CleanRoom_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CleanRooms = types.ListValueMust(t, vs)
}

type NotebookVersionReview_SdkV2 struct {
	// Review comment
	Comment types.String `tfsdk:"comment"`
	// Etag identifying the notebook version
	Etag types.String `tfsdk:"etag"`
	// Review outcome
	ReviewState types.String `tfsdk:"review_state"`
}

func (toState *NotebookVersionReview_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NotebookVersionReview_SdkV2) {
}

func (toState *NotebookVersionReview_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState NotebookVersionReview_SdkV2) {
}

func (c NotebookVersionReview_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NotebookVersionReview_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookVersionReview_SdkV2
// only implements ToObjectValue() and Type().
func (o NotebookVersionReview_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      o.Comment,
			"etag":         o.Etag,
			"review_state": o.ReviewState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotebookVersionReview_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"etag":         types.StringType,
			"review_state": types.StringType,
		},
	}
}

type UpdateCleanRoomAssetRequest_SdkV2 struct {
	// The asset to update. The asset's `name` and `asset_type` fields are used
	// to identify the asset to update.
	Asset types.List `tfsdk:"asset"`
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
	// For notebooks, the name is the notebook file name. For jar analyses, the
	// name is the jar analysis name.
	Name types.String `tfsdk:"-"`
}

func (toState *UpdateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateCleanRoomAssetRequest_SdkV2) {
	if !fromPlan.Asset.IsNull() && !fromPlan.Asset.IsUnknown() {
		if toStateAsset, ok := toState.GetAsset(ctx); ok {
			if fromPlanAsset, ok := fromPlan.GetAsset(ctx); ok {
				toStateAsset.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAsset)
				toState.SetAsset(ctx, toStateAsset)
			}
		}
	}
}

func (toState *UpdateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateCleanRoomAssetRequest_SdkV2) {
	if !fromState.Asset.IsNull() && !fromState.Asset.IsUnknown() {
		if toStateAsset, ok := toState.GetAsset(ctx); ok {
			if fromStateAsset, ok := fromState.GetAsset(ctx); ok {
				toStateAsset.SyncFieldsDuringRead(ctx, fromStateAsset)
				toState.SetAsset(ctx, toStateAsset)
			}
		}
	}
}

func (c UpdateCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["asset"] = attrs["asset"].SetRequired()
	attrs["asset"] = attrs["asset"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCleanRoomAssetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset": basetypes.ListType{
				ElemType: CleanRoomAsset_SdkV2{}.Type(ctx),
			},
			"asset_type":      types.StringType,
			"clean_room_name": types.StringType,
			"name":            types.StringType,
		},
	}
}

// GetAsset returns the value of the Asset field in UpdateCleanRoomAssetRequest_SdkV2 as
// a CleanRoomAsset_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCleanRoomAssetRequest_SdkV2) GetAsset(ctx context.Context) (CleanRoomAsset_SdkV2, bool) {
	var e CleanRoomAsset_SdkV2
	if o.Asset.IsNull() || o.Asset.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAsset_SdkV2
	d := o.Asset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAsset sets the value of the Asset field in UpdateCleanRoomAssetRequest_SdkV2.
func (o *UpdateCleanRoomAssetRequest_SdkV2) SetAsset(ctx context.Context, v CleanRoomAsset_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["asset"]
	o.Asset = types.ListValueMust(t, vs)
}

type UpdateCleanRoomAutoApprovalRuleRequest_SdkV2 struct {
	// The auto-approval rule to update. The rule_id field is used to identify
	// the rule to update.
	AutoApprovalRule types.List `tfsdk:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"-"`
	// A generated UUID identifying the rule.
	RuleId types.String `tfsdk:"-"`
}

func (toState *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !fromPlan.AutoApprovalRule.IsNull() && !fromPlan.AutoApprovalRule.IsUnknown() {
		if toStateAutoApprovalRule, ok := toState.GetAutoApprovalRule(ctx); ok {
			if fromPlanAutoApprovalRule, ok := fromPlan.GetAutoApprovalRule(ctx); ok {
				toStateAutoApprovalRule.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAutoApprovalRule)
				toState.SetAutoApprovalRule(ctx, toStateAutoApprovalRule)
			}
		}
	}
}

func (toState *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !fromState.AutoApprovalRule.IsNull() && !fromState.AutoApprovalRule.IsUnknown() {
		if toStateAutoApprovalRule, ok := toState.GetAutoApprovalRule(ctx); ok {
			if fromStateAutoApprovalRule, ok := fromState.GetAutoApprovalRule(ctx); ok {
				toStateAutoApprovalRule.SyncFieldsDuringRead(ctx, fromStateAutoApprovalRule)
				toState.SetAutoApprovalRule(ctx, toStateAutoApprovalRule)
			}
		}
	}
}

func (c UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_approval_rule"] = attrs["auto_approval_rule"].SetRequired()
	attrs["auto_approval_rule"] = attrs["auto_approval_rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["rule_id"] = attrs["rule_id"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCleanRoomAutoApprovalRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": o.AutoApprovalRule,
			"clean_room_name":    o.CleanRoomName,
			"rule_id":            o.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_approval_rule": basetypes.ListType{
				ElemType: CleanRoomAutoApprovalRule_SdkV2{}.Type(ctx),
			},
			"clean_room_name": types.StringType,
			"rule_id":         types.StringType,
		},
	}
}

// GetAutoApprovalRule returns the value of the AutoApprovalRule field in UpdateCleanRoomAutoApprovalRuleRequest_SdkV2 as
// a CleanRoomAutoApprovalRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule_SdkV2, bool) {
	var e CleanRoomAutoApprovalRule_SdkV2
	if o.AutoApprovalRule.IsNull() || o.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAutoApprovalRule_SdkV2
	d := o.AutoApprovalRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in UpdateCleanRoomAutoApprovalRuleRequest_SdkV2.
func (o *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_approval_rule"]
	o.AutoApprovalRule = types.ListValueMust(t, vs)
}

type UpdateCleanRoomRequest_SdkV2 struct {
	CleanRoom types.List `tfsdk:"clean_room"`
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (toState *UpdateCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan UpdateCleanRoomRequest_SdkV2) {
	if !fromPlan.CleanRoom.IsNull() && !fromPlan.CleanRoom.IsUnknown() {
		if toStateCleanRoom, ok := toState.GetCleanRoom(ctx); ok {
			if fromPlanCleanRoom, ok := fromPlan.GetCleanRoom(ctx); ok {
				toStateCleanRoom.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanCleanRoom)
				toState.SetCleanRoom(ctx, toStateCleanRoom)
			}
		}
	}
}

func (toState *UpdateCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, fromState UpdateCleanRoomRequest_SdkV2) {
	if !fromState.CleanRoom.IsNull() && !fromState.CleanRoom.IsUnknown() {
		if toStateCleanRoom, ok := toState.GetCleanRoom(ctx); ok {
			if fromStateCleanRoom, ok := fromState.GetCleanRoom(ctx); ok {
				toStateCleanRoom.SyncFieldsDuringRead(ctx, fromStateCleanRoom)
				toState.SetCleanRoom(ctx, toStateCleanRoom)
			}
		}
	}
}

func (c UpdateCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room"] = attrs["clean_room"].SetOptional()
	attrs["clean_room"] = attrs["clean_room"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": o.CleanRoom,
			"name":       o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room": basetypes.ListType{
				ElemType: CleanRoom_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetCleanRoom returns the value of the CleanRoom field in UpdateCleanRoomRequest_SdkV2 as
// a CleanRoom_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCleanRoomRequest_SdkV2) GetCleanRoom(ctx context.Context) (CleanRoom_SdkV2, bool) {
	var e CleanRoom_SdkV2
	if o.CleanRoom.IsNull() || o.CleanRoom.IsUnknown() {
		return e, false
	}
	var v []CleanRoom_SdkV2
	d := o.CleanRoom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoom sets the value of the CleanRoom field in UpdateCleanRoomRequest_SdkV2.
func (o *UpdateCleanRoomRequest_SdkV2) SetCleanRoom(ctx context.Context, v CleanRoom_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_room"]
	o.CleanRoom = types.ListValueMust(t, vs)
}
