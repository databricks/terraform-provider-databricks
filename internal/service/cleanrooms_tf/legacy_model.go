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

func (to *CleanRoom_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoom_SdkV2) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				// Recursively sync the fields of OutputCatalog
				toOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
	if !from.RemoteDetailedInfo.IsNull() && !from.RemoteDetailedInfo.IsUnknown() {
		if toRemoteDetailedInfo, ok := to.GetRemoteDetailedInfo(ctx); ok {
			if fromRemoteDetailedInfo, ok := from.GetRemoteDetailedInfo(ctx); ok {
				// Recursively sync the fields of RemoteDetailedInfo
				toRemoteDetailedInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromRemoteDetailedInfo)
				to.SetRemoteDetailedInfo(ctx, toRemoteDetailedInfo)
			}
		}
	}
}

func (to *CleanRoom_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoom_SdkV2) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				toOutputCatalog.SyncFieldsDuringRead(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
	if !from.RemoteDetailedInfo.IsNull() && !from.RemoteDetailedInfo.IsUnknown() {
		if toRemoteDetailedInfo, ok := to.GetRemoteDetailedInfo(ctx); ok {
			if fromRemoteDetailedInfo, ok := from.GetRemoteDetailedInfo(ctx); ok {
				toRemoteDetailedInfo.SyncFieldsDuringRead(ctx, fromRemoteDetailedInfo)
				to.SetRemoteDetailedInfo(ctx, toRemoteDetailedInfo)
			}
		}
	}
}

func (m CleanRoom_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoom_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog":       reflect.TypeOf(CleanRoomOutputCatalog_SdkV2{}),
		"remote_detailed_info": reflect.TypeOf(CleanRoomRemoteDetail_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoom_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoom_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_restricted":        m.AccessRestricted,
			"comment":                  m.Comment,
			"created_at":               m.CreatedAt,
			"local_collaborator_alias": m.LocalCollaboratorAlias,
			"name":                     m.Name,
			"output_catalog":           m.OutputCatalog,
			"owner":                    m.Owner,
			"remote_detailed_info":     m.RemoteDetailedInfo,
			"status":                   m.Status,
			"updated_at":               m.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoom_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoom_SdkV2) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog_SdkV2, bool) {
	var e CleanRoomOutputCatalog_SdkV2
	if m.OutputCatalog.IsNull() || m.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog_SdkV2
	d := m.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CleanRoom_SdkV2.
func (m *CleanRoom_SdkV2) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	m.OutputCatalog = types.ListValueMust(t, vs)
}

// GetRemoteDetailedInfo returns the value of the RemoteDetailedInfo field in CleanRoom_SdkV2 as
// a CleanRoomRemoteDetail_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoom_SdkV2) GetRemoteDetailedInfo(ctx context.Context) (CleanRoomRemoteDetail_SdkV2, bool) {
	var e CleanRoomRemoteDetail_SdkV2
	if m.RemoteDetailedInfo.IsNull() || m.RemoteDetailedInfo.IsUnknown() {
		return e, false
	}
	var v []CleanRoomRemoteDetail_SdkV2
	d := m.RemoteDetailedInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRemoteDetailedInfo sets the value of the RemoteDetailedInfo field in CleanRoom_SdkV2.
func (m *CleanRoom_SdkV2) SetRemoteDetailedInfo(ctx context.Context, v CleanRoomRemoteDetail_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["remote_detailed_info"]
	m.RemoteDetailedInfo = types.ListValueMust(t, vs)
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

func (to *CleanRoomAsset_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAsset_SdkV2) {
	if !from.ForeignTable.IsNull() && !from.ForeignTable.IsUnknown() {
		if toForeignTable, ok := to.GetForeignTable(ctx); ok {
			if fromForeignTable, ok := from.GetForeignTable(ctx); ok {
				// Recursively sync the fields of ForeignTable
				toForeignTable.SyncFieldsDuringCreateOrUpdate(ctx, fromForeignTable)
				to.SetForeignTable(ctx, toForeignTable)
			}
		}
	}
	if !from.ForeignTableLocalDetails.IsNull() && !from.ForeignTableLocalDetails.IsUnknown() {
		if toForeignTableLocalDetails, ok := to.GetForeignTableLocalDetails(ctx); ok {
			if fromForeignTableLocalDetails, ok := from.GetForeignTableLocalDetails(ctx); ok {
				// Recursively sync the fields of ForeignTableLocalDetails
				toForeignTableLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromForeignTableLocalDetails)
				to.SetForeignTableLocalDetails(ctx, toForeignTableLocalDetails)
			}
		}
	}
	if !from.Notebook.IsNull() && !from.Notebook.IsUnknown() {
		if toNotebook, ok := to.GetNotebook(ctx); ok {
			if fromNotebook, ok := from.GetNotebook(ctx); ok {
				// Recursively sync the fields of Notebook
				toNotebook.SyncFieldsDuringCreateOrUpdate(ctx, fromNotebook)
				to.SetNotebook(ctx, toNotebook)
			}
		}
	}
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				// Recursively sync the fields of Table
				toTable.SyncFieldsDuringCreateOrUpdate(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
	if !from.TableLocalDetails.IsNull() && !from.TableLocalDetails.IsUnknown() {
		if toTableLocalDetails, ok := to.GetTableLocalDetails(ctx); ok {
			if fromTableLocalDetails, ok := from.GetTableLocalDetails(ctx); ok {
				// Recursively sync the fields of TableLocalDetails
				toTableLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromTableLocalDetails)
				to.SetTableLocalDetails(ctx, toTableLocalDetails)
			}
		}
	}
	if !from.View.IsNull() && !from.View.IsUnknown() {
		if toView, ok := to.GetView(ctx); ok {
			if fromView, ok := from.GetView(ctx); ok {
				// Recursively sync the fields of View
				toView.SyncFieldsDuringCreateOrUpdate(ctx, fromView)
				to.SetView(ctx, toView)
			}
		}
	}
	if !from.ViewLocalDetails.IsNull() && !from.ViewLocalDetails.IsUnknown() {
		if toViewLocalDetails, ok := to.GetViewLocalDetails(ctx); ok {
			if fromViewLocalDetails, ok := from.GetViewLocalDetails(ctx); ok {
				// Recursively sync the fields of ViewLocalDetails
				toViewLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromViewLocalDetails)
				to.SetViewLocalDetails(ctx, toViewLocalDetails)
			}
		}
	}
	if !from.VolumeLocalDetails.IsNull() && !from.VolumeLocalDetails.IsUnknown() {
		if toVolumeLocalDetails, ok := to.GetVolumeLocalDetails(ctx); ok {
			if fromVolumeLocalDetails, ok := from.GetVolumeLocalDetails(ctx); ok {
				// Recursively sync the fields of VolumeLocalDetails
				toVolumeLocalDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromVolumeLocalDetails)
				to.SetVolumeLocalDetails(ctx, toVolumeLocalDetails)
			}
		}
	}
}

func (to *CleanRoomAsset_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAsset_SdkV2) {
	if !from.ForeignTable.IsNull() && !from.ForeignTable.IsUnknown() {
		if toForeignTable, ok := to.GetForeignTable(ctx); ok {
			if fromForeignTable, ok := from.GetForeignTable(ctx); ok {
				toForeignTable.SyncFieldsDuringRead(ctx, fromForeignTable)
				to.SetForeignTable(ctx, toForeignTable)
			}
		}
	}
	if !from.ForeignTableLocalDetails.IsNull() && !from.ForeignTableLocalDetails.IsUnknown() {
		if toForeignTableLocalDetails, ok := to.GetForeignTableLocalDetails(ctx); ok {
			if fromForeignTableLocalDetails, ok := from.GetForeignTableLocalDetails(ctx); ok {
				toForeignTableLocalDetails.SyncFieldsDuringRead(ctx, fromForeignTableLocalDetails)
				to.SetForeignTableLocalDetails(ctx, toForeignTableLocalDetails)
			}
		}
	}
	if !from.Notebook.IsNull() && !from.Notebook.IsUnknown() {
		if toNotebook, ok := to.GetNotebook(ctx); ok {
			if fromNotebook, ok := from.GetNotebook(ctx); ok {
				toNotebook.SyncFieldsDuringRead(ctx, fromNotebook)
				to.SetNotebook(ctx, toNotebook)
			}
		}
	}
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				toTable.SyncFieldsDuringRead(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
	if !from.TableLocalDetails.IsNull() && !from.TableLocalDetails.IsUnknown() {
		if toTableLocalDetails, ok := to.GetTableLocalDetails(ctx); ok {
			if fromTableLocalDetails, ok := from.GetTableLocalDetails(ctx); ok {
				toTableLocalDetails.SyncFieldsDuringRead(ctx, fromTableLocalDetails)
				to.SetTableLocalDetails(ctx, toTableLocalDetails)
			}
		}
	}
	if !from.View.IsNull() && !from.View.IsUnknown() {
		if toView, ok := to.GetView(ctx); ok {
			if fromView, ok := from.GetView(ctx); ok {
				toView.SyncFieldsDuringRead(ctx, fromView)
				to.SetView(ctx, toView)
			}
		}
	}
	if !from.ViewLocalDetails.IsNull() && !from.ViewLocalDetails.IsUnknown() {
		if toViewLocalDetails, ok := to.GetViewLocalDetails(ctx); ok {
			if fromViewLocalDetails, ok := from.GetViewLocalDetails(ctx); ok {
				toViewLocalDetails.SyncFieldsDuringRead(ctx, fromViewLocalDetails)
				to.SetViewLocalDetails(ctx, toViewLocalDetails)
			}
		}
	}
	if !from.VolumeLocalDetails.IsNull() && !from.VolumeLocalDetails.IsUnknown() {
		if toVolumeLocalDetails, ok := to.GetVolumeLocalDetails(ctx); ok {
			if fromVolumeLocalDetails, ok := from.GetVolumeLocalDetails(ctx); ok {
				toVolumeLocalDetails.SyncFieldsDuringRead(ctx, fromVolumeLocalDetails)
				to.SetVolumeLocalDetails(ctx, toVolumeLocalDetails)
			}
		}
	}
}

func (m CleanRoomAsset_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAsset_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CleanRoomAsset_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"added_at":                    m.AddedAt,
			"asset_type":                  m.AssetType,
			"clean_room_name":             m.CleanRoomName,
			"foreign_table":               m.ForeignTable,
			"foreign_table_local_details": m.ForeignTableLocalDetails,
			"name":                        m.Name,
			"notebook":                    m.Notebook,
			"owner_collaborator_alias":    m.OwnerCollaboratorAlias,
			"status":                      m.Status,
			"table":                       m.Table,
			"table_local_details":         m.TableLocalDetails,
			"view":                        m.View,
			"view_local_details":          m.ViewLocalDetails,
			"volume_local_details":        m.VolumeLocalDetails,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAsset_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAsset_SdkV2) GetForeignTable(ctx context.Context) (CleanRoomAssetForeignTable_SdkV2, bool) {
	var e CleanRoomAssetForeignTable_SdkV2
	if m.ForeignTable.IsNull() || m.ForeignTable.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetForeignTable_SdkV2
	d := m.ForeignTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignTable sets the value of the ForeignTable field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetForeignTable(ctx context.Context, v CleanRoomAssetForeignTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_table"]
	m.ForeignTable = types.ListValueMust(t, vs)
}

// GetForeignTableLocalDetails returns the value of the ForeignTableLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetForeignTableLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetForeignTableLocalDetails(ctx context.Context) (CleanRoomAssetForeignTableLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetForeignTableLocalDetails_SdkV2
	if m.ForeignTableLocalDetails.IsNull() || m.ForeignTableLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetForeignTableLocalDetails_SdkV2
	d := m.ForeignTableLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetForeignTableLocalDetails sets the value of the ForeignTableLocalDetails field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetForeignTableLocalDetails(ctx context.Context, v CleanRoomAssetForeignTableLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["foreign_table_local_details"]
	m.ForeignTableLocalDetails = types.ListValueMust(t, vs)
}

// GetNotebook returns the value of the Notebook field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetNotebook_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetNotebook(ctx context.Context) (CleanRoomAssetNotebook_SdkV2, bool) {
	var e CleanRoomAssetNotebook_SdkV2
	if m.Notebook.IsNull() || m.Notebook.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetNotebook_SdkV2
	d := m.Notebook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebook sets the value of the Notebook field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetNotebook(ctx context.Context, v CleanRoomAssetNotebook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook"]
	m.Notebook = types.ListValueMust(t, vs)
}

// GetTable returns the value of the Table field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetTable(ctx context.Context) (CleanRoomAssetTable_SdkV2, bool) {
	var e CleanRoomAssetTable_SdkV2
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetTable_SdkV2
	d := m.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetTable(ctx context.Context, v CleanRoomAssetTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	m.Table = types.ListValueMust(t, vs)
}

// GetTableLocalDetails returns the value of the TableLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetTableLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetTableLocalDetails(ctx context.Context) (CleanRoomAssetTableLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetTableLocalDetails_SdkV2
	if m.TableLocalDetails.IsNull() || m.TableLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetTableLocalDetails_SdkV2
	d := m.TableLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTableLocalDetails sets the value of the TableLocalDetails field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetTableLocalDetails(ctx context.Context, v CleanRoomAssetTableLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table_local_details"]
	m.TableLocalDetails = types.ListValueMust(t, vs)
}

// GetView returns the value of the View field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetView_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetView(ctx context.Context) (CleanRoomAssetView_SdkV2, bool) {
	var e CleanRoomAssetView_SdkV2
	if m.View.IsNull() || m.View.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetView_SdkV2
	d := m.View.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetView sets the value of the View field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetView(ctx context.Context, v CleanRoomAssetView_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["view"]
	m.View = types.ListValueMust(t, vs)
}

// GetViewLocalDetails returns the value of the ViewLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetViewLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetViewLocalDetails(ctx context.Context) (CleanRoomAssetViewLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetViewLocalDetails_SdkV2
	if m.ViewLocalDetails.IsNull() || m.ViewLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetViewLocalDetails_SdkV2
	d := m.ViewLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetViewLocalDetails sets the value of the ViewLocalDetails field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetViewLocalDetails(ctx context.Context, v CleanRoomAssetViewLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["view_local_details"]
	m.ViewLocalDetails = types.ListValueMust(t, vs)
}

// GetVolumeLocalDetails returns the value of the VolumeLocalDetails field in CleanRoomAsset_SdkV2 as
// a CleanRoomAssetVolumeLocalDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset_SdkV2) GetVolumeLocalDetails(ctx context.Context) (CleanRoomAssetVolumeLocalDetails_SdkV2, bool) {
	var e CleanRoomAssetVolumeLocalDetails_SdkV2
	if m.VolumeLocalDetails.IsNull() || m.VolumeLocalDetails.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAssetVolumeLocalDetails_SdkV2
	d := m.VolumeLocalDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetVolumeLocalDetails sets the value of the VolumeLocalDetails field in CleanRoomAsset_SdkV2.
func (m *CleanRoomAsset_SdkV2) SetVolumeLocalDetails(ctx context.Context, v CleanRoomAssetVolumeLocalDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["volume_local_details"]
	m.VolumeLocalDetails = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTable_SdkV2 struct {
	// The metadata information of the columns in the foreign table
	Columns types.List `tfsdk:"columns"`
}

func (to *CleanRoomAssetForeignTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetForeignTable_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *CleanRoomAssetForeignTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetForeignTable_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m CleanRoomAssetForeignTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetForeignTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTable_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetForeignTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetForeignTable_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetForeignTable_SdkV2) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetForeignTable_SdkV2.
func (m *CleanRoomAssetForeignTable_SdkV2) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTableLocalDetails_SdkV2 struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (to *CleanRoomAssetForeignTableLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetForeignTableLocalDetails_SdkV2) {
}

func (to *CleanRoomAssetForeignTableLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetForeignTableLocalDetails_SdkV2) {
}

func (m CleanRoomAssetForeignTableLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetForeignTableLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTableLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetForeignTableLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetForeignTableLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomAssetNotebook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetNotebook_SdkV2) {
	if !from.Reviews.IsNull() && !from.Reviews.IsUnknown() && to.Reviews.IsNull() && len(from.Reviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Reviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Reviews = from.Reviews
	}
	if !from.RunnerCollaboratorAliases.IsNull() && !from.RunnerCollaboratorAliases.IsUnknown() && to.RunnerCollaboratorAliases.IsNull() && len(from.RunnerCollaboratorAliases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RunnerCollaboratorAliases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RunnerCollaboratorAliases = from.RunnerCollaboratorAliases
	}
}

func (to *CleanRoomAssetNotebook_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetNotebook_SdkV2) {
	if !from.Reviews.IsNull() && !from.Reviews.IsUnknown() && to.Reviews.IsNull() && len(from.Reviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Reviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Reviews = from.Reviews
	}
	if !from.RunnerCollaboratorAliases.IsNull() && !from.RunnerCollaboratorAliases.IsUnknown() && to.RunnerCollaboratorAliases.IsNull() && len(from.RunnerCollaboratorAliases.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RunnerCollaboratorAliases, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RunnerCollaboratorAliases = from.RunnerCollaboratorAliases
	}
}

func (m CleanRoomAssetNotebook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetNotebook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"reviews":                     reflect.TypeOf(CleanRoomNotebookReview_SdkV2{}),
		"runner_collaborator_aliases": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetNotebook_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetNotebook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                        m.Etag,
			"notebook_content":            m.NotebookContent,
			"review_state":                m.ReviewState,
			"reviews":                     m.Reviews,
			"runner_collaborator_aliases": m.RunnerCollaboratorAliases,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetNotebook_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetNotebook_SdkV2) GetReviews(ctx context.Context) ([]CleanRoomNotebookReview_SdkV2, bool) {
	if m.Reviews.IsNull() || m.Reviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview_SdkV2
	d := m.Reviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReviews sets the value of the Reviews field in CleanRoomAssetNotebook_SdkV2.
func (m *CleanRoomAssetNotebook_SdkV2) SetReviews(ctx context.Context, v []CleanRoomNotebookReview_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Reviews = types.ListValueMust(t, vs)
}

// GetRunnerCollaboratorAliases returns the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAssetNotebook_SdkV2) GetRunnerCollaboratorAliases(ctx context.Context) ([]types.String, bool) {
	if m.RunnerCollaboratorAliases.IsNull() || m.RunnerCollaboratorAliases.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RunnerCollaboratorAliases.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRunnerCollaboratorAliases sets the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook_SdkV2.
func (m *CleanRoomAssetNotebook_SdkV2) SetRunnerCollaboratorAliases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runner_collaborator_aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RunnerCollaboratorAliases = types.ListValueMust(t, vs)
}

type CleanRoomAssetTable_SdkV2 struct {
	// The metadata information of the columns in the table
	Columns types.List `tfsdk:"columns"`
}

func (to *CleanRoomAssetTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetTable_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *CleanRoomAssetTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetTable_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m CleanRoomAssetTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTable_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetTable_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetTable_SdkV2) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetTable_SdkV2.
func (m *CleanRoomAssetTable_SdkV2) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetTableLocalDetails_SdkV2 struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName types.String `tfsdk:"local_name"`
	// Partition filtering specification for a shared table.
	Partitions types.List `tfsdk:"partitions"`
}

func (to *CleanRoomAssetTableLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetTableLocalDetails_SdkV2) {
	if !from.Partitions.IsNull() && !from.Partitions.IsUnknown() && to.Partitions.IsNull() && len(from.Partitions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Partitions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Partitions = from.Partitions
	}
}

func (to *CleanRoomAssetTableLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetTableLocalDetails_SdkV2) {
	if !from.Partitions.IsNull() && !from.Partitions.IsUnknown() && to.Partitions.IsNull() && len(from.Partitions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Partitions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Partitions = from.Partitions
	}
}

func (m CleanRoomAssetTableLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetTableLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitions": reflect.TypeOf(sharing_tf.Partition_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTableLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetTableLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
			"partitions": m.Partitions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetTableLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetTableLocalDetails_SdkV2) GetPartitions(ctx context.Context) ([]sharing_tf.Partition_SdkV2, bool) {
	if m.Partitions.IsNull() || m.Partitions.IsUnknown() {
		return nil, false
	}
	var v []sharing_tf.Partition_SdkV2
	d := m.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in CleanRoomAssetTableLocalDetails_SdkV2.
func (m *CleanRoomAssetTableLocalDetails_SdkV2) SetPartitions(ctx context.Context, v []sharing_tf.Partition_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["partitions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Partitions = types.ListValueMust(t, vs)
}

type CleanRoomAssetView_SdkV2 struct {
	// The metadata information of the columns in the view
	Columns types.List `tfsdk:"columns"`
}

func (to *CleanRoomAssetView_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetView_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *CleanRoomAssetView_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetView_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m CleanRoomAssetView_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetView_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetView_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetView_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetView_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetView_SdkV2) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo_SdkV2, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo_SdkV2
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetView_SdkV2.
func (m *CleanRoomAssetView_SdkV2) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetViewLocalDetails_SdkV2 struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (to *CleanRoomAssetViewLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetViewLocalDetails_SdkV2) {
}

func (to *CleanRoomAssetViewLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetViewLocalDetails_SdkV2) {
}

func (m CleanRoomAssetViewLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetViewLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetViewLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetViewLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetViewLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomAssetVolumeLocalDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetVolumeLocalDetails_SdkV2) {
}

func (to *CleanRoomAssetVolumeLocalDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetVolumeLocalDetails_SdkV2) {
}

func (m CleanRoomAssetVolumeLocalDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetVolumeLocalDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetVolumeLocalDetails_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetVolumeLocalDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetVolumeLocalDetails_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomAutoApprovalRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAutoApprovalRule_SdkV2) {
}

func (to *CleanRoomAutoApprovalRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAutoApprovalRule_SdkV2) {
}

func (m CleanRoomAutoApprovalRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAutoApprovalRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAutoApprovalRule_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomAutoApprovalRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"author_collaborator_alias":     m.AuthorCollaboratorAlias,
			"author_scope":                  m.AuthorScope,
			"clean_room_name":               m.CleanRoomName,
			"created_at":                    m.CreatedAt,
			"rule_id":                       m.RuleId,
			"rule_owner_collaborator_alias": m.RuleOwnerCollaboratorAlias,
			"runner_collaborator_alias":     m.RunnerCollaboratorAlias,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAutoApprovalRule_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomCollaborator_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomCollaborator_SdkV2) {
	if !from.InviteRecipientWorkspaceId.IsUnknown() && !from.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InviteRecipientWorkspaceId = from.InviteRecipientWorkspaceId
	}
}

func (to *CleanRoomCollaborator_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomCollaborator_SdkV2) {
	if !from.InviteRecipientWorkspaceId.IsUnknown() && !from.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InviteRecipientWorkspaceId = from.InviteRecipientWorkspaceId
	}
}

func (m CleanRoomCollaborator_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomCollaborator_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomCollaborator_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomCollaborator_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_alias":            m.CollaboratorAlias,
			"display_name":                  m.DisplayName,
			"global_metastore_id":           m.GlobalMetastoreId,
			"invite_recipient_email":        m.InviteRecipientEmail,
			"invite_recipient_workspace_id": m.InviteRecipientWorkspaceId,
			"organization_name":             m.OrganizationName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomCollaborator_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomNotebookReview_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomNotebookReview_SdkV2) {
}

func (to *CleanRoomNotebookReview_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomNotebookReview_SdkV2) {
}

func (m CleanRoomNotebookReview_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomNotebookReview_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookReview_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomNotebookReview_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":                     m.Comment,
			"created_at_millis":           m.CreatedAtMillis,
			"review_state":                m.ReviewState,
			"review_sub_reason":           m.ReviewSubReason,
			"reviewer_collaborator_alias": m.ReviewerCollaboratorAlias,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomNotebookReview_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomNotebookTaskRun_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomNotebookTaskRun_SdkV2) {
	if !from.CollaboratorJobRunInfo.IsNull() && !from.CollaboratorJobRunInfo.IsUnknown() {
		if toCollaboratorJobRunInfo, ok := to.GetCollaboratorJobRunInfo(ctx); ok {
			if fromCollaboratorJobRunInfo, ok := from.GetCollaboratorJobRunInfo(ctx); ok {
				// Recursively sync the fields of CollaboratorJobRunInfo
				toCollaboratorJobRunInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromCollaboratorJobRunInfo)
				to.SetCollaboratorJobRunInfo(ctx, toCollaboratorJobRunInfo)
			}
		}
	}
	if !from.NotebookJobRunState.IsNull() && !from.NotebookJobRunState.IsUnknown() {
		if toNotebookJobRunState, ok := to.GetNotebookJobRunState(ctx); ok {
			if fromNotebookJobRunState, ok := from.GetNotebookJobRunState(ctx); ok {
				// Recursively sync the fields of NotebookJobRunState
				toNotebookJobRunState.SyncFieldsDuringCreateOrUpdate(ctx, fromNotebookJobRunState)
				to.SetNotebookJobRunState(ctx, toNotebookJobRunState)
			}
		}
	}
}

func (to *CleanRoomNotebookTaskRun_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomNotebookTaskRun_SdkV2) {
	if !from.CollaboratorJobRunInfo.IsNull() && !from.CollaboratorJobRunInfo.IsUnknown() {
		if toCollaboratorJobRunInfo, ok := to.GetCollaboratorJobRunInfo(ctx); ok {
			if fromCollaboratorJobRunInfo, ok := from.GetCollaboratorJobRunInfo(ctx); ok {
				toCollaboratorJobRunInfo.SyncFieldsDuringRead(ctx, fromCollaboratorJobRunInfo)
				to.SetCollaboratorJobRunInfo(ctx, toCollaboratorJobRunInfo)
			}
		}
	}
	if !from.NotebookJobRunState.IsNull() && !from.NotebookJobRunState.IsUnknown() {
		if toNotebookJobRunState, ok := to.GetNotebookJobRunState(ctx); ok {
			if fromNotebookJobRunState, ok := from.GetNotebookJobRunState(ctx); ok {
				toNotebookJobRunState.SyncFieldsDuringRead(ctx, fromNotebookJobRunState)
				to.SetNotebookJobRunState(ctx, toNotebookJobRunState)
			}
		}
	}
}

func (m CleanRoomNotebookTaskRun_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomNotebookTaskRun_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"collaborator_job_run_info": reflect.TypeOf(CollaboratorJobRunInfo_SdkV2{}),
		"notebook_job_run_state":    reflect.TypeOf(jobs_tf.CleanRoomTaskRunState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookTaskRun_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomNotebookTaskRun_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_job_run_info":     m.CollaboratorJobRunInfo,
			"notebook_etag":                 m.NotebookEtag,
			"notebook_job_run_state":        m.NotebookJobRunState,
			"notebook_name":                 m.NotebookName,
			"notebook_updated_at":           m.NotebookUpdatedAt,
			"output_schema_expiration_time": m.OutputSchemaExpirationTime,
			"output_schema_name":            m.OutputSchemaName,
			"run_duration":                  m.RunDuration,
			"start_time":                    m.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomNotebookTaskRun_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomNotebookTaskRun_SdkV2) GetCollaboratorJobRunInfo(ctx context.Context) (CollaboratorJobRunInfo_SdkV2, bool) {
	var e CollaboratorJobRunInfo_SdkV2
	if m.CollaboratorJobRunInfo.IsNull() || m.CollaboratorJobRunInfo.IsUnknown() {
		return e, false
	}
	var v []CollaboratorJobRunInfo_SdkV2
	d := m.CollaboratorJobRunInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCollaboratorJobRunInfo sets the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun_SdkV2.
func (m *CleanRoomNotebookTaskRun_SdkV2) SetCollaboratorJobRunInfo(ctx context.Context, v CollaboratorJobRunInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborator_job_run_info"]
	m.CollaboratorJobRunInfo = types.ListValueMust(t, vs)
}

// GetNotebookJobRunState returns the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun_SdkV2 as
// a jobs_tf.CleanRoomTaskRunState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomNotebookTaskRun_SdkV2) GetNotebookJobRunState(ctx context.Context) (jobs_tf.CleanRoomTaskRunState_SdkV2, bool) {
	var e jobs_tf.CleanRoomTaskRunState_SdkV2
	if m.NotebookJobRunState.IsNull() || m.NotebookJobRunState.IsUnknown() {
		return e, false
	}
	var v []jobs_tf.CleanRoomTaskRunState_SdkV2
	d := m.NotebookJobRunState.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookJobRunState sets the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun_SdkV2.
func (m *CleanRoomNotebookTaskRun_SdkV2) SetNotebookJobRunState(ctx context.Context, v jobs_tf.CleanRoomTaskRunState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_job_run_state"]
	m.NotebookJobRunState = types.ListValueMust(t, vs)
}

type CleanRoomOutputCatalog_SdkV2 struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName types.String `tfsdk:"catalog_name"`

	Status types.String `tfsdk:"status"`
}

func (to *CleanRoomOutputCatalog_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomOutputCatalog_SdkV2) {
}

func (to *CleanRoomOutputCatalog_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomOutputCatalog_SdkV2) {
}

func (m CleanRoomOutputCatalog_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomOutputCatalog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomOutputCatalog_SdkV2
// only implements ToObjectValue() and Type().
func (m CleanRoomOutputCatalog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": m.CatalogName,
			"status":       m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomOutputCatalog_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomRemoteDetail_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomRemoteDetail_SdkV2) {
	if !from.Collaborators.IsNull() && !from.Collaborators.IsUnknown() && to.Collaborators.IsNull() && len(from.Collaborators.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Collaborators, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Collaborators = from.Collaborators
	}
	if !from.ComplianceSecurityProfile.IsNull() && !from.ComplianceSecurityProfile.IsUnknown() {
		if toComplianceSecurityProfile, ok := to.GetComplianceSecurityProfile(ctx); ok {
			if fromComplianceSecurityProfile, ok := from.GetComplianceSecurityProfile(ctx); ok {
				// Recursively sync the fields of ComplianceSecurityProfile
				toComplianceSecurityProfile.SyncFieldsDuringCreateOrUpdate(ctx, fromComplianceSecurityProfile)
				to.SetComplianceSecurityProfile(ctx, toComplianceSecurityProfile)
			}
		}
	}
	if !from.Creator.IsNull() && !from.Creator.IsUnknown() {
		if toCreator, ok := to.GetCreator(ctx); ok {
			if fromCreator, ok := from.GetCreator(ctx); ok {
				// Recursively sync the fields of Creator
				toCreator.SyncFieldsDuringCreateOrUpdate(ctx, fromCreator)
				to.SetCreator(ctx, toCreator)
			}
		}
	}
	if !from.EgressNetworkPolicy.IsNull() && !from.EgressNetworkPolicy.IsUnknown() {
		if toEgressNetworkPolicy, ok := to.GetEgressNetworkPolicy(ctx); ok {
			if fromEgressNetworkPolicy, ok := from.GetEgressNetworkPolicy(ctx); ok {
				// Recursively sync the fields of EgressNetworkPolicy
				toEgressNetworkPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromEgressNetworkPolicy)
				to.SetEgressNetworkPolicy(ctx, toEgressNetworkPolicy)
			}
		}
	}
}

func (to *CleanRoomRemoteDetail_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CleanRoomRemoteDetail_SdkV2) {
	if !from.Collaborators.IsNull() && !from.Collaborators.IsUnknown() && to.Collaborators.IsNull() && len(from.Collaborators.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Collaborators, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Collaborators = from.Collaborators
	}
	if !from.ComplianceSecurityProfile.IsNull() && !from.ComplianceSecurityProfile.IsUnknown() {
		if toComplianceSecurityProfile, ok := to.GetComplianceSecurityProfile(ctx); ok {
			if fromComplianceSecurityProfile, ok := from.GetComplianceSecurityProfile(ctx); ok {
				toComplianceSecurityProfile.SyncFieldsDuringRead(ctx, fromComplianceSecurityProfile)
				to.SetComplianceSecurityProfile(ctx, toComplianceSecurityProfile)
			}
		}
	}
	if !from.Creator.IsNull() && !from.Creator.IsUnknown() {
		if toCreator, ok := to.GetCreator(ctx); ok {
			if fromCreator, ok := from.GetCreator(ctx); ok {
				toCreator.SyncFieldsDuringRead(ctx, fromCreator)
				to.SetCreator(ctx, toCreator)
			}
		}
	}
	if !from.EgressNetworkPolicy.IsNull() && !from.EgressNetworkPolicy.IsUnknown() {
		if toEgressNetworkPolicy, ok := to.GetEgressNetworkPolicy(ctx); ok {
			if fromEgressNetworkPolicy, ok := from.GetEgressNetworkPolicy(ctx); ok {
				toEgressNetworkPolicy.SyncFieldsDuringRead(ctx, fromEgressNetworkPolicy)
				to.SetEgressNetworkPolicy(ctx, toEgressNetworkPolicy)
			}
		}
	}
}

func (m CleanRoomRemoteDetail_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomRemoteDetail_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CleanRoomRemoteDetail_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"central_clean_room_id":       m.CentralCleanRoomId,
			"cloud_vendor":                m.CloudVendor,
			"collaborators":               m.Collaborators,
			"compliance_security_profile": m.ComplianceSecurityProfile,
			"creator":                     m.Creator,
			"egress_network_policy":       m.EgressNetworkPolicy,
			"region":                      m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomRemoteDetail_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomRemoteDetail_SdkV2) GetCollaborators(ctx context.Context) ([]CleanRoomCollaborator_SdkV2, bool) {
	if m.Collaborators.IsNull() || m.Collaborators.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomCollaborator_SdkV2
	d := m.Collaborators.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaborators sets the value of the Collaborators field in CleanRoomRemoteDetail_SdkV2.
func (m *CleanRoomRemoteDetail_SdkV2) SetCollaborators(ctx context.Context, v []CleanRoomCollaborator_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborators"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Collaborators = types.ListValueMust(t, vs)
}

// GetComplianceSecurityProfile returns the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail_SdkV2 as
// a ComplianceSecurityProfile_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail_SdkV2) GetComplianceSecurityProfile(ctx context.Context) (ComplianceSecurityProfile_SdkV2, bool) {
	var e ComplianceSecurityProfile_SdkV2
	if m.ComplianceSecurityProfile.IsNull() || m.ComplianceSecurityProfile.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfile_SdkV2
	d := m.ComplianceSecurityProfile.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComplianceSecurityProfile sets the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail_SdkV2.
func (m *CleanRoomRemoteDetail_SdkV2) SetComplianceSecurityProfile(ctx context.Context, v ComplianceSecurityProfile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_security_profile"]
	m.ComplianceSecurityProfile = types.ListValueMust(t, vs)
}

// GetCreator returns the value of the Creator field in CleanRoomRemoteDetail_SdkV2 as
// a CleanRoomCollaborator_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail_SdkV2) GetCreator(ctx context.Context) (CleanRoomCollaborator_SdkV2, bool) {
	var e CleanRoomCollaborator_SdkV2
	if m.Creator.IsNull() || m.Creator.IsUnknown() {
		return e, false
	}
	var v []CleanRoomCollaborator_SdkV2
	d := m.Creator.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCreator sets the value of the Creator field in CleanRoomRemoteDetail_SdkV2.
func (m *CleanRoomRemoteDetail_SdkV2) SetCreator(ctx context.Context, v CleanRoomCollaborator_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["creator"]
	m.Creator = types.ListValueMust(t, vs)
}

// GetEgressNetworkPolicy returns the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail_SdkV2 as
// a settings_tf.EgressNetworkPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail_SdkV2) GetEgressNetworkPolicy(ctx context.Context) (settings_tf.EgressNetworkPolicy_SdkV2, bool) {
	var e settings_tf.EgressNetworkPolicy_SdkV2
	if m.EgressNetworkPolicy.IsNull() || m.EgressNetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []settings_tf.EgressNetworkPolicy_SdkV2
	d := m.EgressNetworkPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEgressNetworkPolicy sets the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail_SdkV2.
func (m *CleanRoomRemoteDetail_SdkV2) SetEgressNetworkPolicy(ctx context.Context, v settings_tf.EgressNetworkPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["egress_network_policy"]
	m.EgressNetworkPolicy = types.ListValueMust(t, vs)
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

func (to *CollaboratorJobRunInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CollaboratorJobRunInfo_SdkV2) {
}

func (to *CollaboratorJobRunInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CollaboratorJobRunInfo_SdkV2) {
}

func (m CollaboratorJobRunInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CollaboratorJobRunInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CollaboratorJobRunInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m CollaboratorJobRunInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_alias":        m.CollaboratorAlias,
			"collaborator_job_id":       m.CollaboratorJobId,
			"collaborator_job_run_id":   m.CollaboratorJobRunId,
			"collaborator_task_run_id":  m.CollaboratorTaskRunId,
			"collaborator_workspace_id": m.CollaboratorWorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CollaboratorJobRunInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ComplianceSecurityProfile_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComplianceSecurityProfile_SdkV2) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (to *ComplianceSecurityProfile_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ComplianceSecurityProfile_SdkV2) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (m ComplianceSecurityProfile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ComplianceSecurityProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile_SdkV2
// only implements ToObjectValue() and Type().
func (m ComplianceSecurityProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": m.ComplianceStandards,
			"is_enabled":           m.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComplianceSecurityProfile_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ComplianceSecurityProfile_SdkV2) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
	if m.ComplianceStandards.IsNull() || m.ComplianceStandards.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ComplianceStandards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile_SdkV2.
func (m *ComplianceSecurityProfile_SdkV2) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ComplianceStandards = types.ListValueMust(t, vs)
}

type CreateCleanRoomAssetRequest_SdkV2 struct {
	Asset types.List `tfsdk:"asset"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName types.String `tfsdk:"-"`
}

func (to *CreateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAssetRequest_SdkV2) {
	if !from.Asset.IsNull() && !from.Asset.IsUnknown() {
		if toAsset, ok := to.GetAsset(ctx); ok {
			if fromAsset, ok := from.GetAsset(ctx); ok {
				// Recursively sync the fields of Asset
				toAsset.SyncFieldsDuringCreateOrUpdate(ctx, fromAsset)
				to.SetAsset(ctx, toAsset)
			}
		}
	}
}

func (to *CreateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAssetRequest_SdkV2) {
	if !from.Asset.IsNull() && !from.Asset.IsUnknown() {
		if toAsset, ok := to.GetAsset(ctx); ok {
			if fromAsset, ok := from.GetAsset(ctx); ok {
				toAsset.SyncFieldsDuringRead(ctx, fromAsset)
				to.SetAsset(ctx, toAsset)
			}
		}
	}
}

func (m CreateCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset":           m.Asset,
			"clean_room_name": m.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomAssetRequest_SdkV2) GetAsset(ctx context.Context) (CleanRoomAsset_SdkV2, bool) {
	var e CleanRoomAsset_SdkV2
	if m.Asset.IsNull() || m.Asset.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAsset_SdkV2
	d := m.Asset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAsset sets the value of the Asset field in CreateCleanRoomAssetRequest_SdkV2.
func (m *CreateCleanRoomAssetRequest_SdkV2) SetAsset(ctx context.Context, v CleanRoomAsset_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["asset"]
	m.Asset = types.ListValueMust(t, vs)
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

func (to *CreateCleanRoomAssetReviewRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAssetReviewRequest_SdkV2) {
	if !from.NotebookReview.IsNull() && !from.NotebookReview.IsUnknown() {
		if toNotebookReview, ok := to.GetNotebookReview(ctx); ok {
			if fromNotebookReview, ok := from.GetNotebookReview(ctx); ok {
				// Recursively sync the fields of NotebookReview
				toNotebookReview.SyncFieldsDuringCreateOrUpdate(ctx, fromNotebookReview)
				to.SetNotebookReview(ctx, toNotebookReview)
			}
		}
	}
}

func (to *CreateCleanRoomAssetReviewRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAssetReviewRequest_SdkV2) {
	if !from.NotebookReview.IsNull() && !from.NotebookReview.IsUnknown() {
		if toNotebookReview, ok := to.GetNotebookReview(ctx); ok {
			if fromNotebookReview, ok := from.GetNotebookReview(ctx); ok {
				toNotebookReview.SyncFieldsDuringRead(ctx, fromNotebookReview)
				to.SetNotebookReview(ctx, toNotebookReview)
			}
		}
	}
}

func (m CreateCleanRoomAssetReviewRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomAssetReviewRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_review": reflect.TypeOf(NotebookVersionReview_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAssetReviewRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
			"notebook_review": m.NotebookReview,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAssetReviewRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomAssetReviewRequest_SdkV2) GetNotebookReview(ctx context.Context) (NotebookVersionReview_SdkV2, bool) {
	var e NotebookVersionReview_SdkV2
	if m.NotebookReview.IsNull() || m.NotebookReview.IsUnknown() {
		return e, false
	}
	var v []NotebookVersionReview_SdkV2
	d := m.NotebookReview.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNotebookReview sets the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest_SdkV2.
func (m *CreateCleanRoomAssetReviewRequest_SdkV2) SetNotebookReview(ctx context.Context, v NotebookVersionReview_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_review"]
	m.NotebookReview = types.ListValueMust(t, vs)
}

type CreateCleanRoomAssetReviewResponse_SdkV2 struct {
	// Top-level status derived from all reviews
	NotebookReviewState types.String `tfsdk:"notebook_review_state"`
	// All existing notebook approvals or rejections
	NotebookReviews types.List `tfsdk:"notebook_reviews"`
}

func (to *CreateCleanRoomAssetReviewResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAssetReviewResponse_SdkV2) {
	if !from.NotebookReviews.IsNull() && !from.NotebookReviews.IsUnknown() && to.NotebookReviews.IsNull() && len(from.NotebookReviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for NotebookReviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.NotebookReviews = from.NotebookReviews
	}
}

func (to *CreateCleanRoomAssetReviewResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAssetReviewResponse_SdkV2) {
	if !from.NotebookReviews.IsNull() && !from.NotebookReviews.IsUnknown() && to.NotebookReviews.IsNull() && len(from.NotebookReviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for NotebookReviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.NotebookReviews = from.NotebookReviews
	}
}

func (m CreateCleanRoomAssetReviewResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomAssetReviewResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"notebook_reviews": reflect.TypeOf(CleanRoomNotebookReview_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAssetReviewResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"notebook_review_state": m.NotebookReviewState,
			"notebook_reviews":      m.NotebookReviews,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAssetReviewResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomAssetReviewResponse_SdkV2) GetNotebookReviews(ctx context.Context) ([]CleanRoomNotebookReview_SdkV2, bool) {
	if m.NotebookReviews.IsNull() || m.NotebookReviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview_SdkV2
	d := m.NotebookReviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookReviews sets the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse_SdkV2.
func (m *CreateCleanRoomAssetReviewResponse_SdkV2) SetNotebookReviews(ctx context.Context, v []CleanRoomNotebookReview_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.NotebookReviews = types.ListValueMust(t, vs)
}

type CreateCleanRoomAutoApprovalRuleRequest_SdkV2 struct {
	AutoApprovalRule types.List `tfsdk:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"-"`
}

func (to *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !from.AutoApprovalRule.IsNull() && !from.AutoApprovalRule.IsUnknown() {
		if toAutoApprovalRule, ok := to.GetAutoApprovalRule(ctx); ok {
			if fromAutoApprovalRule, ok := from.GetAutoApprovalRule(ctx); ok {
				// Recursively sync the fields of AutoApprovalRule
				toAutoApprovalRule.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoApprovalRule)
				to.SetAutoApprovalRule(ctx, toAutoApprovalRule)
			}
		}
	}
}

func (to *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !from.AutoApprovalRule.IsNull() && !from.AutoApprovalRule.IsUnknown() {
		if toAutoApprovalRule, ok := to.GetAutoApprovalRule(ctx); ok {
			if fromAutoApprovalRule, ok := from.GetAutoApprovalRule(ctx); ok {
				toAutoApprovalRule.SyncFieldsDuringRead(ctx, fromAutoApprovalRule)
				to.SetAutoApprovalRule(ctx, toAutoApprovalRule)
			}
		}
	}
}

func (m CreateCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": m.AutoApprovalRule,
			"clean_room_name":    m.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule_SdkV2, bool) {
	var e CleanRoomAutoApprovalRule_SdkV2
	if m.AutoApprovalRule.IsNull() || m.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAutoApprovalRule_SdkV2
	d := m.AutoApprovalRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in CreateCleanRoomAutoApprovalRuleRequest_SdkV2.
func (m *CreateCleanRoomAutoApprovalRuleRequest_SdkV2) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_approval_rule"]
	m.AutoApprovalRule = types.ListValueMust(t, vs)
}

type CreateCleanRoomOutputCatalogRequest_SdkV2 struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`

	OutputCatalog types.List `tfsdk:"output_catalog"`
}

func (to *CreateCleanRoomOutputCatalogRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomOutputCatalogRequest_SdkV2) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				// Recursively sync the fields of OutputCatalog
				toOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
}

func (to *CreateCleanRoomOutputCatalogRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomOutputCatalogRequest_SdkV2) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				toOutputCatalog.SyncFieldsDuringRead(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
}

func (m CreateCleanRoomOutputCatalogRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomOutputCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomOutputCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"output_catalog":  m.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomOutputCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomOutputCatalogRequest_SdkV2) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog_SdkV2, bool) {
	var e CleanRoomOutputCatalog_SdkV2
	if m.OutputCatalog.IsNull() || m.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog_SdkV2
	d := m.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest_SdkV2.
func (m *CreateCleanRoomOutputCatalogRequest_SdkV2) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	m.OutputCatalog = types.ListValueMust(t, vs)
}

type CreateCleanRoomOutputCatalogResponse_SdkV2 struct {
	OutputCatalog types.List `tfsdk:"output_catalog"`
}

func (to *CreateCleanRoomOutputCatalogResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomOutputCatalogResponse_SdkV2) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				// Recursively sync the fields of OutputCatalog
				toOutputCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
}

func (to *CreateCleanRoomOutputCatalogResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomOutputCatalogResponse_SdkV2) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				toOutputCatalog.SyncFieldsDuringRead(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
}

func (m CreateCleanRoomOutputCatalogResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomOutputCatalogResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomOutputCatalogResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"output_catalog": m.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomOutputCatalogResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomOutputCatalogResponse_SdkV2) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog_SdkV2, bool) {
	var e CleanRoomOutputCatalog_SdkV2
	if m.OutputCatalog.IsNull() || m.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v []CleanRoomOutputCatalog_SdkV2
	d := m.OutputCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse_SdkV2.
func (m *CreateCleanRoomOutputCatalogResponse_SdkV2) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["output_catalog"]
	m.OutputCatalog = types.ListValueMust(t, vs)
}

type CreateCleanRoomRequest_SdkV2 struct {
	CleanRoom types.List `tfsdk:"clean_room"`
}

func (to *CreateCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomRequest_SdkV2) {
	if !from.CleanRoom.IsNull() && !from.CleanRoom.IsUnknown() {
		if toCleanRoom, ok := to.GetCleanRoom(ctx); ok {
			if fromCleanRoom, ok := from.GetCleanRoom(ctx); ok {
				// Recursively sync the fields of CleanRoom
				toCleanRoom.SyncFieldsDuringCreateOrUpdate(ctx, fromCleanRoom)
				to.SetCleanRoom(ctx, toCleanRoom)
			}
		}
	}
}

func (to *CreateCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomRequest_SdkV2) {
	if !from.CleanRoom.IsNull() && !from.CleanRoom.IsUnknown() {
		if toCleanRoom, ok := to.GetCleanRoom(ctx); ok {
			if fromCleanRoom, ok := from.GetCleanRoom(ctx); ok {
				toCleanRoom.SyncFieldsDuringRead(ctx, fromCleanRoom)
				to.SetCleanRoom(ctx, toCleanRoom)
			}
		}
	}
}

func (m CreateCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": m.CleanRoom,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomRequest_SdkV2) GetCleanRoom(ctx context.Context) (CleanRoom_SdkV2, bool) {
	var e CleanRoom_SdkV2
	if m.CleanRoom.IsNull() || m.CleanRoom.IsUnknown() {
		return e, false
	}
	var v []CleanRoom_SdkV2
	d := m.CleanRoom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoom sets the value of the CleanRoom field in CreateCleanRoomRequest_SdkV2.
func (m *CreateCleanRoomRequest_SdkV2) SetCleanRoom(ctx context.Context, v CleanRoom_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_room"]
	m.CleanRoom = types.ListValueMust(t, vs)
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

func (to *DeleteCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomAssetRequest_SdkV2) {
}

func (to *DeleteCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomAssetRequest_SdkV2) {
}

func (m DeleteCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeleteCleanRoomAssetResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomAssetResponse_SdkV2) {
}

func (to *DeleteCleanRoomAssetResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomAssetResponse_SdkV2) {
}

func (m DeleteCleanRoomAssetResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAssetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCleanRoomAssetResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomAssetResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomAssetResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteCleanRoomAutoApprovalRuleRequest_SdkV2 struct {
	CleanRoomName types.String `tfsdk:"-"`

	RuleId types.String `tfsdk:"-"`
}

func (to *DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (to *DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (m DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"rule_id":         m.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *DeleteCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomRequest_SdkV2) {
}

func (to *DeleteCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomRequest_SdkV2) {
}

func (m DeleteCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomAssetRequest_SdkV2) {
}

func (to *GetCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomAssetRequest_SdkV2) {
}

func (m GetCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomAssetRevisionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomAssetRevisionRequest_SdkV2) {
}

func (to *GetCleanRoomAssetRevisionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomAssetRevisionRequest_SdkV2) {
}

func (m GetCleanRoomAssetRevisionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomAssetRevisionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRevisionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCleanRoomAssetRevisionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"etag":            m.Etag,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomAssetRevisionRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (to *GetCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomAutoApprovalRuleRequest_SdkV2) {
}

func (m GetCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"rule_id":         m.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomRequest_SdkV2) {
}

func (to *GetCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomRequest_SdkV2) {
}

func (m GetCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAssetRevisionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetRevisionsRequest_SdkV2) {
}

func (to *ListCleanRoomAssetRevisionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetRevisionsRequest_SdkV2) {
}

func (m ListCleanRoomAssetRevisionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetRevisionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetRevisionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
			"page_size":       m.PageSize,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetRevisionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAssetRevisionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetRevisionsResponse_SdkV2) {
	if !from.Revisions.IsNull() && !from.Revisions.IsUnknown() && to.Revisions.IsNull() && len(from.Revisions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Revisions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Revisions = from.Revisions
	}
}

func (to *ListCleanRoomAssetRevisionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetRevisionsResponse_SdkV2) {
	if !from.Revisions.IsNull() && !from.Revisions.IsUnknown() && to.Revisions.IsNull() && len(from.Revisions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Revisions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Revisions = from.Revisions
	}
}

func (m ListCleanRoomAssetRevisionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetRevisionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"revisions": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetRevisionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"revisions":       m.Revisions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetRevisionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomAssetRevisionsResponse_SdkV2) GetRevisions(ctx context.Context) ([]CleanRoomAsset_SdkV2, bool) {
	if m.Revisions.IsNull() || m.Revisions.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset_SdkV2
	d := m.Revisions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRevisions sets the value of the Revisions field in ListCleanRoomAssetRevisionsResponse_SdkV2.
func (m *ListCleanRoomAssetRevisionsResponse_SdkV2) SetRevisions(ctx context.Context, v []CleanRoomAsset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["revisions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Revisions = types.ListValueMust(t, vs)
}

type ListCleanRoomAssetsRequest_SdkV2 struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCleanRoomAssetsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetsRequest_SdkV2) {
}

func (to *ListCleanRoomAssetsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetsRequest_SdkV2) {
}

func (m ListCleanRoomAssetsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAssetsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetsResponse_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
}

func (to *ListCleanRoomAssetsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetsResponse_SdkV2) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
}

func (m ListCleanRoomAssetsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":          m.Assets,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomAssetsResponse_SdkV2) GetAssets(ctx context.Context) ([]CleanRoomAsset_SdkV2, bool) {
	if m.Assets.IsNull() || m.Assets.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset_SdkV2
	d := m.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListCleanRoomAssetsResponse_SdkV2.
func (m *ListCleanRoomAssetsResponse_SdkV2) SetAssets(ctx context.Context, v []CleanRoomAsset_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Assets = types.ListValueMust(t, vs)
}

type ListCleanRoomAutoApprovalRulesRequest_SdkV2 struct {
	CleanRoomName types.String `tfsdk:"-"`
	// Maximum number of auto-approval rules to return. Defaults to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCleanRoomAutoApprovalRulesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAutoApprovalRulesRequest_SdkV2) {
}

func (to *ListCleanRoomAutoApprovalRulesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAutoApprovalRulesRequest_SdkV2) {
}

func (m ListCleanRoomAutoApprovalRulesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAutoApprovalRulesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAutoApprovalRulesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"page_size":       m.PageSize,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAutoApprovalRulesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAutoApprovalRulesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAutoApprovalRulesResponse_SdkV2) {
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() && to.Rules.IsNull() && len(from.Rules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Rules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Rules = from.Rules
	}
}

func (to *ListCleanRoomAutoApprovalRulesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAutoApprovalRulesResponse_SdkV2) {
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() && to.Rules.IsNull() && len(from.Rules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Rules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Rules = from.Rules
	}
}

func (m ListCleanRoomAutoApprovalRulesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAutoApprovalRulesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(CleanRoomAutoApprovalRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAutoApprovalRulesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"rules":           m.Rules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAutoApprovalRulesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomAutoApprovalRulesResponse_SdkV2) GetRules(ctx context.Context) ([]CleanRoomAutoApprovalRule_SdkV2, bool) {
	if m.Rules.IsNull() || m.Rules.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAutoApprovalRule_SdkV2
	d := m.Rules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRules sets the value of the Rules field in ListCleanRoomAutoApprovalRulesResponse_SdkV2.
func (m *ListCleanRoomAutoApprovalRulesResponse_SdkV2) SetRules(ctx context.Context, v []CleanRoomAutoApprovalRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Rules = types.ListValueMust(t, vs)
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

func (to *ListCleanRoomNotebookTaskRunsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomNotebookTaskRunsRequest_SdkV2) {
}

func (to *ListCleanRoomNotebookTaskRunsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomNotebookTaskRunsRequest_SdkV2) {
}

func (m ListCleanRoomNotebookTaskRunsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomNotebookTaskRunsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomNotebookTaskRunsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"notebook_name":   m.NotebookName,
			"page_size":       m.PageSize,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomNotebookTaskRunsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomNotebookTaskRunsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomNotebookTaskRunsResponse_SdkV2) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (to *ListCleanRoomNotebookTaskRunsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomNotebookTaskRunsResponse_SdkV2) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (m ListCleanRoomNotebookTaskRunsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomNotebookTaskRunsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(CleanRoomNotebookTaskRun_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomNotebookTaskRunsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"runs":            m.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomNotebookTaskRunsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomNotebookTaskRunsResponse_SdkV2) GetRuns(ctx context.Context) ([]CleanRoomNotebookTaskRun_SdkV2, bool) {
	if m.Runs.IsNull() || m.Runs.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookTaskRun_SdkV2
	d := m.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in ListCleanRoomNotebookTaskRunsResponse_SdkV2.
func (m *ListCleanRoomNotebookTaskRunsResponse_SdkV2) SetRuns(ctx context.Context, v []CleanRoomNotebookTaskRun_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Runs = types.ListValueMust(t, vs)
}

type ListCleanRoomsRequest_SdkV2 struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCleanRoomsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomsRequest_SdkV2) {
}

func (to *ListCleanRoomsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomsRequest_SdkV2) {
}

func (m ListCleanRoomsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomsResponse_SdkV2) {
	if !from.CleanRooms.IsNull() && !from.CleanRooms.IsUnknown() && to.CleanRooms.IsNull() && len(from.CleanRooms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CleanRooms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CleanRooms = from.CleanRooms
	}
}

func (to *ListCleanRoomsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomsResponse_SdkV2) {
	if !from.CleanRooms.IsNull() && !from.CleanRooms.IsUnknown() && to.CleanRooms.IsNull() && len(from.CleanRooms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CleanRooms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CleanRooms = from.CleanRooms
	}
}

func (m ListCleanRoomsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms": reflect.TypeOf(CleanRoom_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCleanRoomsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_rooms":     m.CleanRooms,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomsResponse_SdkV2) GetCleanRooms(ctx context.Context) ([]CleanRoom_SdkV2, bool) {
	if m.CleanRooms.IsNull() || m.CleanRooms.IsUnknown() {
		return nil, false
	}
	var v []CleanRoom_SdkV2
	d := m.CleanRooms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRooms sets the value of the CleanRooms field in ListCleanRoomsResponse_SdkV2.
func (m *ListCleanRoomsResponse_SdkV2) SetCleanRooms(ctx context.Context, v []CleanRoom_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CleanRooms = types.ListValueMust(t, vs)
}

type NotebookVersionReview_SdkV2 struct {
	// Review comment
	Comment types.String `tfsdk:"comment"`
	// Etag identifying the notebook version
	Etag types.String `tfsdk:"etag"`
	// Review outcome
	ReviewState types.String `tfsdk:"review_state"`
}

func (to *NotebookVersionReview_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotebookVersionReview_SdkV2) {
}

func (to *NotebookVersionReview_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NotebookVersionReview_SdkV2) {
}

func (m NotebookVersionReview_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NotebookVersionReview_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookVersionReview_SdkV2
// only implements ToObjectValue() and Type().
func (m NotebookVersionReview_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      m.Comment,
			"etag":         m.Etag,
			"review_state": m.ReviewState,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotebookVersionReview_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *UpdateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCleanRoomAssetRequest_SdkV2) {
	if !from.Asset.IsNull() && !from.Asset.IsUnknown() {
		if toAsset, ok := to.GetAsset(ctx); ok {
			if fromAsset, ok := from.GetAsset(ctx); ok {
				// Recursively sync the fields of Asset
				toAsset.SyncFieldsDuringCreateOrUpdate(ctx, fromAsset)
				to.SetAsset(ctx, toAsset)
			}
		}
	}
}

func (to *UpdateCleanRoomAssetRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCleanRoomAssetRequest_SdkV2) {
	if !from.Asset.IsNull() && !from.Asset.IsUnknown() {
		if toAsset, ok := to.GetAsset(ctx); ok {
			if fromAsset, ok := from.GetAsset(ctx); ok {
				toAsset.SyncFieldsDuringRead(ctx, fromAsset)
				to.SetAsset(ctx, toAsset)
			}
		}
	}
}

func (m UpdateCleanRoomAssetRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateCleanRoomAssetRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAssetRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCleanRoomAssetRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset":           m.Asset,
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCleanRoomAssetRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateCleanRoomAssetRequest_SdkV2) GetAsset(ctx context.Context) (CleanRoomAsset_SdkV2, bool) {
	var e CleanRoomAsset_SdkV2
	if m.Asset.IsNull() || m.Asset.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAsset_SdkV2
	d := m.Asset.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAsset sets the value of the Asset field in UpdateCleanRoomAssetRequest_SdkV2.
func (m *UpdateCleanRoomAssetRequest_SdkV2) SetAsset(ctx context.Context, v CleanRoomAsset_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["asset"]
	m.Asset = types.ListValueMust(t, vs)
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

func (to *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !from.AutoApprovalRule.IsNull() && !from.AutoApprovalRule.IsUnknown() {
		if toAutoApprovalRule, ok := to.GetAutoApprovalRule(ctx); ok {
			if fromAutoApprovalRule, ok := from.GetAutoApprovalRule(ctx); ok {
				// Recursively sync the fields of AutoApprovalRule
				toAutoApprovalRule.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoApprovalRule)
				to.SetAutoApprovalRule(ctx, toAutoApprovalRule)
			}
		}
	}
}

func (to *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) {
	if !from.AutoApprovalRule.IsNull() && !from.AutoApprovalRule.IsUnknown() {
		if toAutoApprovalRule, ok := to.GetAutoApprovalRule(ctx); ok {
			if fromAutoApprovalRule, ok := from.GetAutoApprovalRule(ctx); ok {
				toAutoApprovalRule.SyncFieldsDuringRead(ctx, fromAutoApprovalRule)
				to.SetAutoApprovalRule(ctx, toAutoApprovalRule)
			}
		}
	}
}

func (m UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAutoApprovalRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": m.AutoApprovalRule,
			"clean_room_name":    m.CleanRoomName,
			"rule_id":            m.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule_SdkV2, bool) {
	var e CleanRoomAutoApprovalRule_SdkV2
	if m.AutoApprovalRule.IsNull() || m.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v []CleanRoomAutoApprovalRule_SdkV2
	d := m.AutoApprovalRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in UpdateCleanRoomAutoApprovalRuleRequest_SdkV2.
func (m *UpdateCleanRoomAutoApprovalRuleRequest_SdkV2) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_approval_rule"]
	m.AutoApprovalRule = types.ListValueMust(t, vs)
}

type UpdateCleanRoomRequest_SdkV2 struct {
	CleanRoom types.List `tfsdk:"clean_room"`
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateCleanRoomRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCleanRoomRequest_SdkV2) {
	if !from.CleanRoom.IsNull() && !from.CleanRoom.IsUnknown() {
		if toCleanRoom, ok := to.GetCleanRoom(ctx); ok {
			if fromCleanRoom, ok := from.GetCleanRoom(ctx); ok {
				// Recursively sync the fields of CleanRoom
				toCleanRoom.SyncFieldsDuringCreateOrUpdate(ctx, fromCleanRoom)
				to.SetCleanRoom(ctx, toCleanRoom)
			}
		}
	}
}

func (to *UpdateCleanRoomRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCleanRoomRequest_SdkV2) {
	if !from.CleanRoom.IsNull() && !from.CleanRoom.IsUnknown() {
		if toCleanRoom, ok := to.GetCleanRoom(ctx); ok {
			if fromCleanRoom, ok := from.GetCleanRoom(ctx); ok {
				toCleanRoom.SyncFieldsDuringRead(ctx, fromCleanRoom)
				to.SetCleanRoom(ctx, toCleanRoom)
			}
		}
	}
}

func (m UpdateCleanRoomRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateCleanRoomRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCleanRoomRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": m.CleanRoom,
			"name":       m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCleanRoomRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *UpdateCleanRoomRequest_SdkV2) GetCleanRoom(ctx context.Context) (CleanRoom_SdkV2, bool) {
	var e CleanRoom_SdkV2
	if m.CleanRoom.IsNull() || m.CleanRoom.IsUnknown() {
		return e, false
	}
	var v []CleanRoom_SdkV2
	d := m.CleanRoom.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCleanRoom sets the value of the CleanRoom field in UpdateCleanRoomRequest_SdkV2.
func (m *UpdateCleanRoomRequest_SdkV2) SetCleanRoom(ctx context.Context, v CleanRoom_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_room"]
	m.CleanRoom = types.ListValueMust(t, vs)
}
