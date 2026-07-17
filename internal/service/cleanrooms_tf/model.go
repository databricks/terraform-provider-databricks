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

	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"  // .tmpl
	"github.com/databricks/terraform-provider-databricks/internal/service/jobs_tf"     // .tmpl
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf" // .tmpl
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"  // .tmpl
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
	// Whether allow task to write to shared output schema. When enabled, clean
	// room task runs triggered by the current collaborator can write to the
	// run-scoped shared output schema which is accessible by all collaborators.
	EnableSharedOutput types.Bool `tfsdk:"enable_shared_output"`
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
	// This is the Databricks username of the owner of the local clean room
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

func (to *CleanRoom) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoom) {
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

func (to *CleanRoom) SyncFieldsDuringRead(ctx context.Context, from CleanRoom) {
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

func (m CleanRoom) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_restricted"] = attrs["access_restricted"].SetComputed()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["enable_shared_output"] = attrs["enable_shared_output"].SetOptional()
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
func (m CleanRoom) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog":       reflect.TypeOf(CleanRoomOutputCatalog{}),
		"remote_detailed_info": reflect.TypeOf(CleanRoomRemoteDetail{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoom
// only implements ToObjectValue() and Type().
func (m CleanRoom) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_restricted":        m.AccessRestricted,
			"comment":                  m.Comment,
			"created_at":               m.CreatedAt,
			"enable_shared_output":     m.EnableSharedOutput,
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
func (m CleanRoom) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_restricted":        types.StringType,
			"comment":                  types.StringType,
			"created_at":               types.Int64Type,
			"enable_shared_output":     types.BoolType,
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
func (m *CleanRoom) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog, bool) {
	var e CleanRoomOutputCatalog
	if m.OutputCatalog.IsNull() || m.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v CleanRoomOutputCatalog
	d := m.OutputCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CleanRoom.
func (m *CleanRoom) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := v.ToObjectValue(ctx)
	m.OutputCatalog = vs
}

// GetRemoteDetailedInfo returns the value of the RemoteDetailedInfo field in CleanRoom as
// a CleanRoomRemoteDetail value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoom) GetRemoteDetailedInfo(ctx context.Context) (CleanRoomRemoteDetail, bool) {
	var e CleanRoomRemoteDetail
	if m.RemoteDetailedInfo.IsNull() || m.RemoteDetailedInfo.IsUnknown() {
		return e, false
	}
	var v CleanRoomRemoteDetail
	d := m.RemoteDetailedInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRemoteDetailedInfo sets the value of the RemoteDetailedInfo field in CleanRoom.
func (m *CleanRoom) SetRemoteDetailedInfo(ctx context.Context, v CleanRoomRemoteDetail) {
	vs := v.ToObjectValue(ctx)
	m.RemoteDetailedInfo = vs
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
	// Jar analysis details available to all collaborators of the clean room.
	// Present if and only if **asset_type** is **JAR_ANALYSIS**
	JarAnalysis types.Object `tfsdk:"jar_analysis"`
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

func (to *CleanRoomAsset) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAsset) {
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
	if !from.JarAnalysis.IsNull() && !from.JarAnalysis.IsUnknown() {
		if toJarAnalysis, ok := to.GetJarAnalysis(ctx); ok {
			if fromJarAnalysis, ok := from.GetJarAnalysis(ctx); ok {
				// Recursively sync the fields of JarAnalysis
				toJarAnalysis.SyncFieldsDuringCreateOrUpdate(ctx, fromJarAnalysis)
				to.SetJarAnalysis(ctx, toJarAnalysis)
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

func (to *CleanRoomAsset) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAsset) {
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
	if !from.JarAnalysis.IsNull() && !from.JarAnalysis.IsUnknown() {
		if toJarAnalysis, ok := to.GetJarAnalysis(ctx); ok {
			if fromJarAnalysis, ok := from.GetJarAnalysis(ctx); ok {
				toJarAnalysis.SyncFieldsDuringRead(ctx, fromJarAnalysis)
				to.SetJarAnalysis(ctx, toJarAnalysis)
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

func (m CleanRoomAsset) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["added_at"] = attrs["added_at"].SetComputed()
	attrs["asset_type"] = attrs["asset_type"].SetRequired()
	attrs["clean_room_name"] = attrs["clean_room_name"].SetOptional()
	attrs["foreign_table"] = attrs["foreign_table"].SetOptional()
	attrs["foreign_table_local_details"] = attrs["foreign_table_local_details"].SetOptional()
	attrs["jar_analysis"] = attrs["jar_analysis"].SetOptional()
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
func (m CleanRoomAsset) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"foreign_table":               reflect.TypeOf(CleanRoomAssetForeignTable{}),
		"foreign_table_local_details": reflect.TypeOf(CleanRoomAssetForeignTableLocalDetails{}),
		"jar_analysis":                reflect.TypeOf(CleanRoomAssetJarAnalysis{}),
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
func (m CleanRoomAsset) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"added_at":                    m.AddedAt,
			"asset_type":                  m.AssetType,
			"clean_room_name":             m.CleanRoomName,
			"foreign_table":               m.ForeignTable,
			"foreign_table_local_details": m.ForeignTableLocalDetails,
			"jar_analysis":                m.JarAnalysis,
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
func (m CleanRoomAsset) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"added_at":                    types.Int64Type,
			"asset_type":                  types.StringType,
			"clean_room_name":             types.StringType,
			"foreign_table":               CleanRoomAssetForeignTable{}.Type(ctx),
			"foreign_table_local_details": CleanRoomAssetForeignTableLocalDetails{}.Type(ctx),
			"jar_analysis":                CleanRoomAssetJarAnalysis{}.Type(ctx),
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
func (m *CleanRoomAsset) GetForeignTable(ctx context.Context) (CleanRoomAssetForeignTable, bool) {
	var e CleanRoomAssetForeignTable
	if m.ForeignTable.IsNull() || m.ForeignTable.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetForeignTable
	d := m.ForeignTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetForeignTable sets the value of the ForeignTable field in CleanRoomAsset.
func (m *CleanRoomAsset) SetForeignTable(ctx context.Context, v CleanRoomAssetForeignTable) {
	vs := v.ToObjectValue(ctx)
	m.ForeignTable = vs
}

// GetForeignTableLocalDetails returns the value of the ForeignTableLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetForeignTableLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetForeignTableLocalDetails(ctx context.Context) (CleanRoomAssetForeignTableLocalDetails, bool) {
	var e CleanRoomAssetForeignTableLocalDetails
	if m.ForeignTableLocalDetails.IsNull() || m.ForeignTableLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetForeignTableLocalDetails
	d := m.ForeignTableLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetForeignTableLocalDetails sets the value of the ForeignTableLocalDetails field in CleanRoomAsset.
func (m *CleanRoomAsset) SetForeignTableLocalDetails(ctx context.Context, v CleanRoomAssetForeignTableLocalDetails) {
	vs := v.ToObjectValue(ctx)
	m.ForeignTableLocalDetails = vs
}

// GetJarAnalysis returns the value of the JarAnalysis field in CleanRoomAsset as
// a CleanRoomAssetJarAnalysis value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetJarAnalysis(ctx context.Context) (CleanRoomAssetJarAnalysis, bool) {
	var e CleanRoomAssetJarAnalysis
	if m.JarAnalysis.IsNull() || m.JarAnalysis.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetJarAnalysis
	d := m.JarAnalysis.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarAnalysis sets the value of the JarAnalysis field in CleanRoomAsset.
func (m *CleanRoomAsset) SetJarAnalysis(ctx context.Context, v CleanRoomAssetJarAnalysis) {
	vs := v.ToObjectValue(ctx)
	m.JarAnalysis = vs
}

// GetNotebook returns the value of the Notebook field in CleanRoomAsset as
// a CleanRoomAssetNotebook value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetNotebook(ctx context.Context) (CleanRoomAssetNotebook, bool) {
	var e CleanRoomAssetNotebook
	if m.Notebook.IsNull() || m.Notebook.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetNotebook
	d := m.Notebook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebook sets the value of the Notebook field in CleanRoomAsset.
func (m *CleanRoomAsset) SetNotebook(ctx context.Context, v CleanRoomAssetNotebook) {
	vs := v.ToObjectValue(ctx)
	m.Notebook = vs
}

// GetTable returns the value of the Table field in CleanRoomAsset as
// a CleanRoomAssetTable value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetTable(ctx context.Context) (CleanRoomAssetTable, bool) {
	var e CleanRoomAssetTable
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetTable
	d := m.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTable sets the value of the Table field in CleanRoomAsset.
func (m *CleanRoomAsset) SetTable(ctx context.Context, v CleanRoomAssetTable) {
	vs := v.ToObjectValue(ctx)
	m.Table = vs
}

// GetTableLocalDetails returns the value of the TableLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetTableLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetTableLocalDetails(ctx context.Context) (CleanRoomAssetTableLocalDetails, bool) {
	var e CleanRoomAssetTableLocalDetails
	if m.TableLocalDetails.IsNull() || m.TableLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetTableLocalDetails
	d := m.TableLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTableLocalDetails sets the value of the TableLocalDetails field in CleanRoomAsset.
func (m *CleanRoomAsset) SetTableLocalDetails(ctx context.Context, v CleanRoomAssetTableLocalDetails) {
	vs := v.ToObjectValue(ctx)
	m.TableLocalDetails = vs
}

// GetView returns the value of the View field in CleanRoomAsset as
// a CleanRoomAssetView value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetView(ctx context.Context) (CleanRoomAssetView, bool) {
	var e CleanRoomAssetView
	if m.View.IsNull() || m.View.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetView
	d := m.View.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetView sets the value of the View field in CleanRoomAsset.
func (m *CleanRoomAsset) SetView(ctx context.Context, v CleanRoomAssetView) {
	vs := v.ToObjectValue(ctx)
	m.View = vs
}

// GetViewLocalDetails returns the value of the ViewLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetViewLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetViewLocalDetails(ctx context.Context) (CleanRoomAssetViewLocalDetails, bool) {
	var e CleanRoomAssetViewLocalDetails
	if m.ViewLocalDetails.IsNull() || m.ViewLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetViewLocalDetails
	d := m.ViewLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetViewLocalDetails sets the value of the ViewLocalDetails field in CleanRoomAsset.
func (m *CleanRoomAsset) SetViewLocalDetails(ctx context.Context, v CleanRoomAssetViewLocalDetails) {
	vs := v.ToObjectValue(ctx)
	m.ViewLocalDetails = vs
}

// GetVolumeLocalDetails returns the value of the VolumeLocalDetails field in CleanRoomAsset as
// a CleanRoomAssetVolumeLocalDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAsset) GetVolumeLocalDetails(ctx context.Context) (CleanRoomAssetVolumeLocalDetails, bool) {
	var e CleanRoomAssetVolumeLocalDetails
	if m.VolumeLocalDetails.IsNull() || m.VolumeLocalDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomAssetVolumeLocalDetails
	d := m.VolumeLocalDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetVolumeLocalDetails sets the value of the VolumeLocalDetails field in CleanRoomAsset.
func (m *CleanRoomAsset) SetVolumeLocalDetails(ctx context.Context, v CleanRoomAssetVolumeLocalDetails) {
	vs := v.ToObjectValue(ctx)
	m.VolumeLocalDetails = vs
}

type CleanRoomAssetForeignTable struct {
	// The metadata information of the columns in the foreign table
	Columns types.List `tfsdk:"columns"`
}

func (to *CleanRoomAssetForeignTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetForeignTable) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *CleanRoomAssetForeignTable) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetForeignTable) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m CleanRoomAssetForeignTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetForeignTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTable
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetForeignTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetForeignTable) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetForeignTable) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetForeignTable.
func (m *CleanRoomAssetForeignTable) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetForeignTableLocalDetails struct {
	// The fully qualified name of the foreign table in its owner's local
	// metastore, in the format of *catalog*.*schema*.*foreign_table_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (to *CleanRoomAssetForeignTableLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetForeignTableLocalDetails) {
}

func (to *CleanRoomAssetForeignTableLocalDetails) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetForeignTableLocalDetails) {
}

func (m CleanRoomAssetForeignTableLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetForeignTableLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetForeignTableLocalDetails
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetForeignTableLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetForeignTableLocalDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAssetJarAnalysis struct {
	// The full paths in central to the jar files that are added to the library
	// during execution (e.g.
	// /Volumes/creator/schema/volume/folder/my_jar_file.jar) Only returned for
	// the owner collaborator.
	CentralJarFilePaths types.List `tfsdk:"central_jar_file_paths"`
	// Optional description of the jar analysis shown to all collaborators.
	Description types.String `tfsdk:"description"`
	// The serverless environment version used to execute the JAR analysis (e.g.
	// "4"). Defaults to "4-scala-preview" if not specified.
	EnvironmentVersion types.String `tfsdk:"environment_version"`
	// Server generated etag that represents the jar analysis version.
	Etag types.String `tfsdk:"etag"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library The code must
	// use `SparkContext.getOrCreate` to obtain a Spark context; otherwise, runs
	// of the job fail
	MainClassName types.String `tfsdk:"main_class_name"`
	// Top-level status derived from all reviews.
	ReviewState types.String `tfsdk:"review_state"`
	// All existing approvals or rejections.
	Reviews types.List `tfsdk:"reviews"`
	// Collaborators that can run the jar.
	RunnerCollaboratorAliases types.List `tfsdk:"runner_collaborator_aliases"`
}

func (to *CleanRoomAssetJarAnalysis) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetJarAnalysis) {
	if !from.CentralJarFilePaths.IsNull() && !from.CentralJarFilePaths.IsUnknown() && to.CentralJarFilePaths.IsNull() && len(from.CentralJarFilePaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CentralJarFilePaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CentralJarFilePaths = from.CentralJarFilePaths
	}
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

func (to *CleanRoomAssetJarAnalysis) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetJarAnalysis) {
	if !from.CentralJarFilePaths.IsNull() && !from.CentralJarFilePaths.IsUnknown() && to.CentralJarFilePaths.IsNull() && len(from.CentralJarFilePaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CentralJarFilePaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CentralJarFilePaths = from.CentralJarFilePaths
	}
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

func (m CleanRoomAssetJarAnalysis) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["central_jar_file_paths"] = attrs["central_jar_file_paths"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["environment_version"] = attrs["environment_version"].SetOptional()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["main_class_name"] = attrs["main_class_name"].SetOptional()
	attrs["review_state"] = attrs["review_state"].SetOptional()
	attrs["reviews"] = attrs["reviews"].SetOptional()
	attrs["runner_collaborator_aliases"] = attrs["runner_collaborator_aliases"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomAssetJarAnalysis.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CleanRoomAssetJarAnalysis) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"central_jar_file_paths":      reflect.TypeOf(types.String{}),
		"reviews":                     reflect.TypeOf(CleanRoomJarAnalysisReview{}),
		"runner_collaborator_aliases": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetJarAnalysis
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetJarAnalysis) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"central_jar_file_paths":      m.CentralJarFilePaths,
			"description":                 m.Description,
			"environment_version":         m.EnvironmentVersion,
			"etag":                        m.Etag,
			"main_class_name":             m.MainClassName,
			"review_state":                m.ReviewState,
			"reviews":                     m.Reviews,
			"runner_collaborator_aliases": m.RunnerCollaboratorAliases,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetJarAnalysis) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"central_jar_file_paths": basetypes.ListType{
				ElemType: types.StringType,
			},
			"description":         types.StringType,
			"environment_version": types.StringType,
			"etag":                types.StringType,
			"main_class_name":     types.StringType,
			"review_state":        types.StringType,
			"reviews": basetypes.ListType{
				ElemType: CleanRoomJarAnalysisReview{}.Type(ctx),
			},
			"runner_collaborator_aliases": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetCentralJarFilePaths returns the value of the CentralJarFilePaths field in CleanRoomAssetJarAnalysis as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAssetJarAnalysis) GetCentralJarFilePaths(ctx context.Context) ([]types.String, bool) {
	if m.CentralJarFilePaths.IsNull() || m.CentralJarFilePaths.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.CentralJarFilePaths.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCentralJarFilePaths sets the value of the CentralJarFilePaths field in CleanRoomAssetJarAnalysis.
func (m *CleanRoomAssetJarAnalysis) SetCentralJarFilePaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["central_jar_file_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CentralJarFilePaths = types.ListValueMust(t, vs)
}

// GetReviews returns the value of the Reviews field in CleanRoomAssetJarAnalysis as
// a slice of CleanRoomJarAnalysisReview values.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAssetJarAnalysis) GetReviews(ctx context.Context) ([]CleanRoomJarAnalysisReview, bool) {
	if m.Reviews.IsNull() || m.Reviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomJarAnalysisReview
	d := m.Reviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReviews sets the value of the Reviews field in CleanRoomAssetJarAnalysis.
func (m *CleanRoomAssetJarAnalysis) SetReviews(ctx context.Context, v []CleanRoomJarAnalysisReview) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Reviews = types.ListValueMust(t, vs)
}

// GetRunnerCollaboratorAliases returns the value of the RunnerCollaboratorAliases field in CleanRoomAssetJarAnalysis as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAssetJarAnalysis) GetRunnerCollaboratorAliases(ctx context.Context) ([]types.String, bool) {
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

// SetRunnerCollaboratorAliases sets the value of the RunnerCollaboratorAliases field in CleanRoomAssetJarAnalysis.
func (m *CleanRoomAssetJarAnalysis) SetRunnerCollaboratorAliases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runner_collaborator_aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RunnerCollaboratorAliases = types.ListValueMust(t, vs)
}

type CleanRoomAssetNotebook struct {
	// Optional description of the notebook shown to all collaborators.
	Description types.String `tfsdk:"description"`
	// The serverless environment version used to execute the notebook (e.g.
	// "4"). Defaults to "2" if not specified.
	EnvironmentVersion types.String `tfsdk:"environment_version"`
	// Server generated etag that represents the notebook version.
	Etag types.String `tfsdk:"etag"`
	// Base 64 representation of the notebook contents. This is the same format
	// as returned by [workspace/export] with the format of **HTML**.
	//
	// [workspace/export]: https://docs.databricks.com/api/workspace/workspace/export
	NotebookContent types.String `tfsdk:"notebook_content"`
	// Top-level status derived from all reviews
	ReviewState types.String `tfsdk:"review_state"`
	// All existing approvals or rejections
	Reviews types.List `tfsdk:"reviews"`
	// Aliases of collaborators that can run the notebook.
	RunnerCollaboratorAliases types.List `tfsdk:"runner_collaborator_aliases"`
}

func (to *CleanRoomAssetNotebook) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetNotebook) {
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

func (to *CleanRoomAssetNotebook) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetNotebook) {
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

func (m CleanRoomAssetNotebook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["environment_version"] = attrs["environment_version"].SetOptional()
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
func (m CleanRoomAssetNotebook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"reviews":                     reflect.TypeOf(CleanRoomNotebookReview{}),
		"runner_collaborator_aliases": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetNotebook
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetNotebook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":                 m.Description,
			"environment_version":         m.EnvironmentVersion,
			"etag":                        m.Etag,
			"notebook_content":            m.NotebookContent,
			"review_state":                m.ReviewState,
			"reviews":                     m.Reviews,
			"runner_collaborator_aliases": m.RunnerCollaboratorAliases,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetNotebook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":         types.StringType,
			"environment_version": types.StringType,
			"etag":                types.StringType,
			"notebook_content":    types.StringType,
			"review_state":        types.StringType,
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
func (m *CleanRoomAssetNotebook) GetReviews(ctx context.Context) ([]CleanRoomNotebookReview, bool) {
	if m.Reviews.IsNull() || m.Reviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview
	d := m.Reviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetReviews sets the value of the Reviews field in CleanRoomAssetNotebook.
func (m *CleanRoomAssetNotebook) SetReviews(ctx context.Context, v []CleanRoomNotebookReview) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Reviews = types.ListValueMust(t, vs)
}

// GetRunnerCollaboratorAliases returns the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomAssetNotebook) GetRunnerCollaboratorAliases(ctx context.Context) ([]types.String, bool) {
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

// SetRunnerCollaboratorAliases sets the value of the RunnerCollaboratorAliases field in CleanRoomAssetNotebook.
func (m *CleanRoomAssetNotebook) SetRunnerCollaboratorAliases(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runner_collaborator_aliases"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RunnerCollaboratorAliases = types.ListValueMust(t, vs)
}

type CleanRoomAssetTable struct {
	// The metadata information of the columns in the table
	Columns types.List `tfsdk:"columns"`
}

func (to *CleanRoomAssetTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetTable) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *CleanRoomAssetTable) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetTable) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m CleanRoomAssetTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTable
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetTable) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetTable) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetTable.
func (m *CleanRoomAssetTable) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetTableLocalDetails struct {
	// The fully qualified name of the table in its owner's local metastore, in
	// the format of *catalog*.*schema*.*table_name*
	LocalName types.String `tfsdk:"local_name"`
	// Partition filtering specification for a shared table.
	Partitions types.List `tfsdk:"partitions"`
}

func (to *CleanRoomAssetTableLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetTableLocalDetails) {
	if !from.Partitions.IsNull() && !from.Partitions.IsUnknown() && to.Partitions.IsNull() && len(from.Partitions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Partitions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Partitions = from.Partitions
	}
}

func (to *CleanRoomAssetTableLocalDetails) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetTableLocalDetails) {
	if !from.Partitions.IsNull() && !from.Partitions.IsUnknown() && to.Partitions.IsNull() && len(from.Partitions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Partitions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Partitions = from.Partitions
	}
}

func (m CleanRoomAssetTableLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetTableLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitions": reflect.TypeOf(sharing_tf.Partition{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetTableLocalDetails
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetTableLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
			"partitions": m.Partitions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetTableLocalDetails) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetTableLocalDetails) GetPartitions(ctx context.Context) ([]sharing_tf.Partition, bool) {
	if m.Partitions.IsNull() || m.Partitions.IsUnknown() {
		return nil, false
	}
	var v []sharing_tf.Partition
	d := m.Partitions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitions sets the value of the Partitions field in CleanRoomAssetTableLocalDetails.
func (m *CleanRoomAssetTableLocalDetails) SetPartitions(ctx context.Context, v []sharing_tf.Partition) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["partitions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Partitions = types.ListValueMust(t, vs)
}

type CleanRoomAssetView struct {
	// The metadata information of the columns in the view
	Columns types.List `tfsdk:"columns"`
}

func (to *CleanRoomAssetView) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetView) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (to *CleanRoomAssetView) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetView) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
}

func (m CleanRoomAssetView) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetView) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(catalog_tf.ColumnInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetView
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetView) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetView) Type(ctx context.Context) attr.Type {
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
func (m *CleanRoomAssetView) GetColumns(ctx context.Context) ([]catalog_tf.ColumnInfo, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []catalog_tf.ColumnInfo
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in CleanRoomAssetView.
func (m *CleanRoomAssetView) SetColumns(ctx context.Context, v []catalog_tf.ColumnInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

type CleanRoomAssetViewLocalDetails struct {
	// The fully qualified name of the view in its owner's local metastore, in
	// the format of *catalog*.*schema*.*view_name*
	LocalName types.String `tfsdk:"local_name"`
}

func (to *CleanRoomAssetViewLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetViewLocalDetails) {
}

func (to *CleanRoomAssetViewLocalDetails) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetViewLocalDetails) {
}

func (m CleanRoomAssetViewLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetViewLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetViewLocalDetails
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetViewLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetViewLocalDetails) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomAssetVolumeLocalDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAssetVolumeLocalDetails) {
}

func (to *CleanRoomAssetVolumeLocalDetails) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAssetVolumeLocalDetails) {
}

func (m CleanRoomAssetVolumeLocalDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAssetVolumeLocalDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAssetVolumeLocalDetails
// only implements ToObjectValue() and Type().
func (m CleanRoomAssetVolumeLocalDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"local_name": m.LocalName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomAssetVolumeLocalDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"local_name": types.StringType,
		},
	}
}

type CleanRoomAutoApprovalRule struct {
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

func (to *CleanRoomAutoApprovalRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomAutoApprovalRule) {
}

func (to *CleanRoomAutoApprovalRule) SyncFieldsDuringRead(ctx context.Context, from CleanRoomAutoApprovalRule) {
}

func (m CleanRoomAutoApprovalRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomAutoApprovalRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomAutoApprovalRule
// only implements ToObjectValue() and Type().
func (m CleanRoomAutoApprovalRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CleanRoomAutoApprovalRule) Type(ctx context.Context) attr.Type {
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

func (to *CleanRoomCollaborator) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomCollaborator) {
	if !from.InviteRecipientWorkspaceId.IsUnknown() && !from.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InviteRecipientWorkspaceId = from.InviteRecipientWorkspaceId
	}
}

func (to *CleanRoomCollaborator) SyncFieldsDuringRead(ctx context.Context, from CleanRoomCollaborator) {
	if !from.InviteRecipientWorkspaceId.IsUnknown() && !from.InviteRecipientWorkspaceId.IsNull() {
		// InviteRecipientWorkspaceId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.InviteRecipientWorkspaceId = from.InviteRecipientWorkspaceId
	}
}

func (m CleanRoomCollaborator) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomCollaborator) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomCollaborator
// only implements ToObjectValue() and Type().
func (m CleanRoomCollaborator) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CleanRoomCollaborator) Type(ctx context.Context) attr.Type {
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

// This only applies to a JAR Analysis as a first-class asset in the Clean Room,
// and not to Volumes
type CleanRoomJarAnalysisReview struct {
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

func (to *CleanRoomJarAnalysisReview) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomJarAnalysisReview) {
}

func (to *CleanRoomJarAnalysisReview) SyncFieldsDuringRead(ctx context.Context, from CleanRoomJarAnalysisReview) {
}

func (m CleanRoomJarAnalysisReview) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_at_millis"] = attrs["created_at_millis"].SetOptional()
	attrs["review_state"] = attrs["review_state"].SetOptional()
	attrs["review_sub_reason"] = attrs["review_sub_reason"].SetOptional()
	attrs["reviewer_collaborator_alias"] = attrs["reviewer_collaborator_alias"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomJarAnalysisReview.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CleanRoomJarAnalysisReview) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomJarAnalysisReview
// only implements ToObjectValue() and Type().
func (m CleanRoomJarAnalysisReview) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CleanRoomJarAnalysisReview) Type(ctx context.Context) attr.Type {
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

type CleanRoomNotebookReview struct {
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

func (to *CleanRoomNotebookReview) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomNotebookReview) {
}

func (to *CleanRoomNotebookReview) SyncFieldsDuringRead(ctx context.Context, from CleanRoomNotebookReview) {
}

func (m CleanRoomNotebookReview) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomNotebookReview) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookReview
// only implements ToObjectValue() and Type().
func (m CleanRoomNotebookReview) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CleanRoomNotebookReview) Type(ctx context.Context) attr.Type {
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
	// Expiration time of the shared output schema of the task run (if any), in
	// epoch milliseconds.
	SharedOutputSchemaExpirationTime types.Int64 `tfsdk:"shared_output_schema_expiration_time"`
	// Name of the shared output schema associated with the clean rooms notebook
	// task run. This schema is accessible by all collaborators when
	// enable_shared_output is true.
	SharedOutputSchemaName types.String `tfsdk:"shared_output_schema_name"`
	// When the task run started, in epoch milliseconds.
	StartTime types.Int64 `tfsdk:"start_time"`
}

func (to *CleanRoomNotebookTaskRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomNotebookTaskRun) {
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

func (to *CleanRoomNotebookTaskRun) SyncFieldsDuringRead(ctx context.Context, from CleanRoomNotebookTaskRun) {
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

func (m CleanRoomNotebookTaskRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["collaborator_job_run_info"] = attrs["collaborator_job_run_info"].SetOptional()
	attrs["notebook_etag"] = attrs["notebook_etag"].SetOptional()
	attrs["notebook_job_run_state"] = attrs["notebook_job_run_state"].SetOptional()
	attrs["notebook_name"] = attrs["notebook_name"].SetOptional()
	attrs["notebook_updated_at"] = attrs["notebook_updated_at"].SetOptional()
	attrs["output_schema_expiration_time"] = attrs["output_schema_expiration_time"].SetOptional()
	attrs["output_schema_name"] = attrs["output_schema_name"].SetOptional()
	attrs["run_duration"] = attrs["run_duration"].SetOptional()
	attrs["shared_output_schema_expiration_time"] = attrs["shared_output_schema_expiration_time"].SetOptional()
	attrs["shared_output_schema_name"] = attrs["shared_output_schema_name"].SetOptional()
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
func (m CleanRoomNotebookTaskRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"collaborator_job_run_info": reflect.TypeOf(CollaboratorJobRunInfo{}),
		"notebook_job_run_state":    reflect.TypeOf(jobs_tf.CleanRoomTaskRunState{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomNotebookTaskRun
// only implements ToObjectValue() and Type().
func (m CleanRoomNotebookTaskRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"collaborator_job_run_info":            m.CollaboratorJobRunInfo,
			"notebook_etag":                        m.NotebookEtag,
			"notebook_job_run_state":               m.NotebookJobRunState,
			"notebook_name":                        m.NotebookName,
			"notebook_updated_at":                  m.NotebookUpdatedAt,
			"output_schema_expiration_time":        m.OutputSchemaExpirationTime,
			"output_schema_name":                   m.OutputSchemaName,
			"run_duration":                         m.RunDuration,
			"shared_output_schema_expiration_time": m.SharedOutputSchemaExpirationTime,
			"shared_output_schema_name":            m.SharedOutputSchemaName,
			"start_time":                           m.StartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomNotebookTaskRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"collaborator_job_run_info":            CollaboratorJobRunInfo{}.Type(ctx),
			"notebook_etag":                        types.StringType,
			"notebook_job_run_state":               jobs_tf.CleanRoomTaskRunState{}.Type(ctx),
			"notebook_name":                        types.StringType,
			"notebook_updated_at":                  types.Int64Type,
			"output_schema_expiration_time":        types.Int64Type,
			"output_schema_name":                   types.StringType,
			"run_duration":                         types.Int64Type,
			"shared_output_schema_expiration_time": types.Int64Type,
			"shared_output_schema_name":            types.StringType,
			"start_time":                           types.Int64Type,
		},
	}
}

// GetCollaboratorJobRunInfo returns the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun as
// a CollaboratorJobRunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomNotebookTaskRun) GetCollaboratorJobRunInfo(ctx context.Context) (CollaboratorJobRunInfo, bool) {
	var e CollaboratorJobRunInfo
	if m.CollaboratorJobRunInfo.IsNull() || m.CollaboratorJobRunInfo.IsUnknown() {
		return e, false
	}
	var v CollaboratorJobRunInfo
	d := m.CollaboratorJobRunInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaboratorJobRunInfo sets the value of the CollaboratorJobRunInfo field in CleanRoomNotebookTaskRun.
func (m *CleanRoomNotebookTaskRun) SetCollaboratorJobRunInfo(ctx context.Context, v CollaboratorJobRunInfo) {
	vs := v.ToObjectValue(ctx)
	m.CollaboratorJobRunInfo = vs
}

// GetNotebookJobRunState returns the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun as
// a jobs_tf.CleanRoomTaskRunState value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomNotebookTaskRun) GetNotebookJobRunState(ctx context.Context) (jobs_tf.CleanRoomTaskRunState, bool) {
	var e jobs_tf.CleanRoomTaskRunState
	if m.NotebookJobRunState.IsNull() || m.NotebookJobRunState.IsUnknown() {
		return e, false
	}
	var v jobs_tf.CleanRoomTaskRunState
	d := m.NotebookJobRunState.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookJobRunState sets the value of the NotebookJobRunState field in CleanRoomNotebookTaskRun.
func (m *CleanRoomNotebookTaskRun) SetNotebookJobRunState(ctx context.Context, v jobs_tf.CleanRoomTaskRunState) {
	vs := v.ToObjectValue(ctx)
	m.NotebookJobRunState = vs
}

type CleanRoomOutputCatalog struct {
	// The name of the output catalog in UC. It should follow [UC securable
	// naming requirements]. The field will always exist if status is CREATED.
	//
	// [UC securable naming requirements]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#securable-object-naming-requirements
	CatalogName types.String `tfsdk:"catalog_name"`

	Status types.String `tfsdk:"status"`
}

func (to *CleanRoomOutputCatalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomOutputCatalog) {
}

func (to *CleanRoomOutputCatalog) SyncFieldsDuringRead(ctx context.Context, from CleanRoomOutputCatalog) {
}

func (m CleanRoomOutputCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CleanRoomOutputCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomOutputCatalog
// only implements ToObjectValue() and Type().
func (m CleanRoomOutputCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name": m.CatalogName,
			"status":       m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomOutputCatalog) Type(ctx context.Context) attr.Type {
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
	// Whether to enable shared output for the central clean room. When enabled,
	// clean room task runs can write to the run-scoped shared output schema
	// which is accessible by all collaborators.
	EnableSharedOutput types.Bool `tfsdk:"enable_shared_output"`
	// Alias of the provider collaborator. If set, packaged clean rooms mode is
	// enabled. The consumer's experience is restricted: they can view notebook
	// names and READMEs, add their own data assets, and trigger runs, but
	// cannot view notebook code, provider data assets, or notebook run output.
	PackageProviderCollaboratorAlias types.String `tfsdk:"package_provider_collaborator_alias"`
	// Region of the central clean room.
	Region types.String `tfsdk:"region"`
}

func (to *CleanRoomRemoteDetail) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomRemoteDetail) {
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

func (to *CleanRoomRemoteDetail) SyncFieldsDuringRead(ctx context.Context, from CleanRoomRemoteDetail) {
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

func (m CleanRoomRemoteDetail) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["central_clean_room_id"] = attrs["central_clean_room_id"].SetComputed()
	attrs["cloud_vendor"] = attrs["cloud_vendor"].SetOptional()
	attrs["cloud_vendor"] = attrs["cloud_vendor"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["collaborators"] = attrs["collaborators"].SetOptional()
	attrs["collaborators"] = attrs["collaborators"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["compliance_security_profile"] = attrs["compliance_security_profile"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["egress_network_policy"] = attrs["egress_network_policy"].SetOptional()
	attrs["egress_network_policy"] = attrs["egress_network_policy"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["enable_shared_output"] = attrs["enable_shared_output"].SetOptional()
	attrs["enable_shared_output"] = attrs["enable_shared_output"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["package_provider_collaborator_alias"] = attrs["package_provider_collaborator_alias"].SetOptional()
	attrs["package_provider_collaborator_alias"] = attrs["package_provider_collaborator_alias"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
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
func (m CleanRoomRemoteDetail) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m CleanRoomRemoteDetail) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"central_clean_room_id":               m.CentralCleanRoomId,
			"cloud_vendor":                        m.CloudVendor,
			"collaborators":                       m.Collaborators,
			"compliance_security_profile":         m.ComplianceSecurityProfile,
			"creator":                             m.Creator,
			"egress_network_policy":               m.EgressNetworkPolicy,
			"enable_shared_output":                m.EnableSharedOutput,
			"package_provider_collaborator_alias": m.PackageProviderCollaboratorAlias,
			"region":                              m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomRemoteDetail) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"central_clean_room_id": types.StringType,
			"cloud_vendor":          types.StringType,
			"collaborators": basetypes.ListType{
				ElemType: CleanRoomCollaborator{}.Type(ctx),
			},
			"compliance_security_profile":         ComplianceSecurityProfile{}.Type(ctx),
			"creator":                             CleanRoomCollaborator{}.Type(ctx),
			"egress_network_policy":               settings_tf.EgressNetworkPolicy{}.Type(ctx),
			"enable_shared_output":                types.BoolType,
			"package_provider_collaborator_alias": types.StringType,
			"region":                              types.StringType,
		},
	}
}

// GetCollaborators returns the value of the Collaborators field in CleanRoomRemoteDetail as
// a slice of CleanRoomCollaborator values.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail) GetCollaborators(ctx context.Context) ([]CleanRoomCollaborator, bool) {
	if m.Collaborators.IsNull() || m.Collaborators.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomCollaborator
	d := m.Collaborators.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaborators sets the value of the Collaborators field in CleanRoomRemoteDetail.
func (m *CleanRoomRemoteDetail) SetCollaborators(ctx context.Context, v []CleanRoomCollaborator) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["collaborators"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Collaborators = types.ListValueMust(t, vs)
}

// GetComplianceSecurityProfile returns the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail as
// a ComplianceSecurityProfile value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail) GetComplianceSecurityProfile(ctx context.Context) (ComplianceSecurityProfile, bool) {
	var e ComplianceSecurityProfile
	if m.ComplianceSecurityProfile.IsNull() || m.ComplianceSecurityProfile.IsUnknown() {
		return e, false
	}
	var v ComplianceSecurityProfile
	d := m.ComplianceSecurityProfile.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceSecurityProfile sets the value of the ComplianceSecurityProfile field in CleanRoomRemoteDetail.
func (m *CleanRoomRemoteDetail) SetComplianceSecurityProfile(ctx context.Context, v ComplianceSecurityProfile) {
	vs := v.ToObjectValue(ctx)
	m.ComplianceSecurityProfile = vs
}

// GetCreator returns the value of the Creator field in CleanRoomRemoteDetail as
// a CleanRoomCollaborator value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail) GetCreator(ctx context.Context) (CleanRoomCollaborator, bool) {
	var e CleanRoomCollaborator
	if m.Creator.IsNull() || m.Creator.IsUnknown() {
		return e, false
	}
	var v CleanRoomCollaborator
	d := m.Creator.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCreator sets the value of the Creator field in CleanRoomRemoteDetail.
func (m *CleanRoomRemoteDetail) SetCreator(ctx context.Context, v CleanRoomCollaborator) {
	vs := v.ToObjectValue(ctx)
	m.Creator = vs
}

// GetEgressNetworkPolicy returns the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail as
// a settings_tf.EgressNetworkPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomRemoteDetail) GetEgressNetworkPolicy(ctx context.Context) (settings_tf.EgressNetworkPolicy, bool) {
	var e settings_tf.EgressNetworkPolicy
	if m.EgressNetworkPolicy.IsNull() || m.EgressNetworkPolicy.IsUnknown() {
		return e, false
	}
	var v settings_tf.EgressNetworkPolicy
	d := m.EgressNetworkPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEgressNetworkPolicy sets the value of the EgressNetworkPolicy field in CleanRoomRemoteDetail.
func (m *CleanRoomRemoteDetail) SetEgressNetworkPolicy(ctx context.Context, v settings_tf.EgressNetworkPolicy) {
	vs := v.ToObjectValue(ctx)
	m.EgressNetworkPolicy = vs
}

// Stores information about a single task run.
type CleanRoomTaskRun struct {
	// Information about the analysis run (etag, updated at)
	AnalysisDetails types.Object `tfsdk:"analysis_details"`
	// Job run info of the task in the runner's local workspace. This field is
	// only included in the LIST API if the task was run within the same
	// workspace the API is being called. If the task run was in a different
	// workspace under the same metastore, only the workspace_id is included.
	CollaboratorJobRunInfo types.Object `tfsdk:"collaborator_job_run_info"`
	// Name of the executable.
	Name types.String `tfsdk:"name"`
	// Information about run output
	OutputInfo types.Object `tfsdk:"output_info"`
	// Duration of the task run, in milliseconds.
	RunDuration types.Int64 `tfsdk:"run_duration"`
	// Information about shared output accessible by all collaborators. This
	// field is only populated when enable_shared_output is true.
	SharedOutputInfo types.Object `tfsdk:"shared_output_info"`
	// When the task run started, in epoch milliseconds.
	StartTime types.Int64 `tfsdk:"start_time"`
	// State of the task run.
	TaskRunState types.Object `tfsdk:"task_run_state"`
	// The type of Clean Room task.
	TaskType types.String `tfsdk:"task_type"`
}

func (to *CleanRoomTaskRun) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomTaskRun) {
	if !from.AnalysisDetails.IsNull() && !from.AnalysisDetails.IsUnknown() {
		if toAnalysisDetails, ok := to.GetAnalysisDetails(ctx); ok {
			if fromAnalysisDetails, ok := from.GetAnalysisDetails(ctx); ok {
				// Recursively sync the fields of AnalysisDetails
				toAnalysisDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromAnalysisDetails)
				to.SetAnalysisDetails(ctx, toAnalysisDetails)
			}
		}
	}
	if !from.CollaboratorJobRunInfo.IsNull() && !from.CollaboratorJobRunInfo.IsUnknown() {
		if toCollaboratorJobRunInfo, ok := to.GetCollaboratorJobRunInfo(ctx); ok {
			if fromCollaboratorJobRunInfo, ok := from.GetCollaboratorJobRunInfo(ctx); ok {
				// Recursively sync the fields of CollaboratorJobRunInfo
				toCollaboratorJobRunInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromCollaboratorJobRunInfo)
				to.SetCollaboratorJobRunInfo(ctx, toCollaboratorJobRunInfo)
			}
		}
	}
	if !from.OutputInfo.IsNull() && !from.OutputInfo.IsUnknown() {
		if toOutputInfo, ok := to.GetOutputInfo(ctx); ok {
			if fromOutputInfo, ok := from.GetOutputInfo(ctx); ok {
				// Recursively sync the fields of OutputInfo
				toOutputInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromOutputInfo)
				to.SetOutputInfo(ctx, toOutputInfo)
			}
		}
	}
	if !from.SharedOutputInfo.IsNull() && !from.SharedOutputInfo.IsUnknown() {
		if toSharedOutputInfo, ok := to.GetSharedOutputInfo(ctx); ok {
			if fromSharedOutputInfo, ok := from.GetSharedOutputInfo(ctx); ok {
				// Recursively sync the fields of SharedOutputInfo
				toSharedOutputInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromSharedOutputInfo)
				to.SetSharedOutputInfo(ctx, toSharedOutputInfo)
			}
		}
	}
	if !from.TaskRunState.IsNull() && !from.TaskRunState.IsUnknown() {
		if toTaskRunState, ok := to.GetTaskRunState(ctx); ok {
			if fromTaskRunState, ok := from.GetTaskRunState(ctx); ok {
				// Recursively sync the fields of TaskRunState
				toTaskRunState.SyncFieldsDuringCreateOrUpdate(ctx, fromTaskRunState)
				to.SetTaskRunState(ctx, toTaskRunState)
			}
		}
	}
}

func (to *CleanRoomTaskRun) SyncFieldsDuringRead(ctx context.Context, from CleanRoomTaskRun) {
	if !from.AnalysisDetails.IsNull() && !from.AnalysisDetails.IsUnknown() {
		if toAnalysisDetails, ok := to.GetAnalysisDetails(ctx); ok {
			if fromAnalysisDetails, ok := from.GetAnalysisDetails(ctx); ok {
				toAnalysisDetails.SyncFieldsDuringRead(ctx, fromAnalysisDetails)
				to.SetAnalysisDetails(ctx, toAnalysisDetails)
			}
		}
	}
	if !from.CollaboratorJobRunInfo.IsNull() && !from.CollaboratorJobRunInfo.IsUnknown() {
		if toCollaboratorJobRunInfo, ok := to.GetCollaboratorJobRunInfo(ctx); ok {
			if fromCollaboratorJobRunInfo, ok := from.GetCollaboratorJobRunInfo(ctx); ok {
				toCollaboratorJobRunInfo.SyncFieldsDuringRead(ctx, fromCollaboratorJobRunInfo)
				to.SetCollaboratorJobRunInfo(ctx, toCollaboratorJobRunInfo)
			}
		}
	}
	if !from.OutputInfo.IsNull() && !from.OutputInfo.IsUnknown() {
		if toOutputInfo, ok := to.GetOutputInfo(ctx); ok {
			if fromOutputInfo, ok := from.GetOutputInfo(ctx); ok {
				toOutputInfo.SyncFieldsDuringRead(ctx, fromOutputInfo)
				to.SetOutputInfo(ctx, toOutputInfo)
			}
		}
	}
	if !from.SharedOutputInfo.IsNull() && !from.SharedOutputInfo.IsUnknown() {
		if toSharedOutputInfo, ok := to.GetSharedOutputInfo(ctx); ok {
			if fromSharedOutputInfo, ok := from.GetSharedOutputInfo(ctx); ok {
				toSharedOutputInfo.SyncFieldsDuringRead(ctx, fromSharedOutputInfo)
				to.SetSharedOutputInfo(ctx, toSharedOutputInfo)
			}
		}
	}
	if !from.TaskRunState.IsNull() && !from.TaskRunState.IsUnknown() {
		if toTaskRunState, ok := to.GetTaskRunState(ctx); ok {
			if fromTaskRunState, ok := from.GetTaskRunState(ctx); ok {
				toTaskRunState.SyncFieldsDuringRead(ctx, fromTaskRunState)
				to.SetTaskRunState(ctx, toTaskRunState)
			}
		}
	}
}

func (m CleanRoomTaskRun) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["analysis_details"] = attrs["analysis_details"].SetOptional()
	attrs["collaborator_job_run_info"] = attrs["collaborator_job_run_info"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["output_info"] = attrs["output_info"].SetOptional()
	attrs["run_duration"] = attrs["run_duration"].SetOptional()
	attrs["shared_output_info"] = attrs["shared_output_info"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["task_run_state"] = attrs["task_run_state"].SetOptional()
	attrs["task_type"] = attrs["task_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomTaskRun.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CleanRoomTaskRun) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"analysis_details":          reflect.TypeOf(CleanRoomTaskRunCleanRoomTaskAnalysisDetails{}),
		"collaborator_job_run_info": reflect.TypeOf(CollaboratorJobRunInfo{}),
		"output_info":               reflect.TypeOf(CleanRoomTaskRunOutputInfo{}),
		"shared_output_info":        reflect.TypeOf(CleanRoomTaskRunOutputInfo{}),
		"task_run_state":            reflect.TypeOf(jobs_tf.CleanRoomTaskRunState{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomTaskRun
// only implements ToObjectValue() and Type().
func (m CleanRoomTaskRun) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"analysis_details":          m.AnalysisDetails,
			"collaborator_job_run_info": m.CollaboratorJobRunInfo,
			"name":                      m.Name,
			"output_info":               m.OutputInfo,
			"run_duration":              m.RunDuration,
			"shared_output_info":        m.SharedOutputInfo,
			"start_time":                m.StartTime,
			"task_run_state":            m.TaskRunState,
			"task_type":                 m.TaskType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomTaskRun) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"analysis_details":          CleanRoomTaskRunCleanRoomTaskAnalysisDetails{}.Type(ctx),
			"collaborator_job_run_info": CollaboratorJobRunInfo{}.Type(ctx),
			"name":                      types.StringType,
			"output_info":               CleanRoomTaskRunOutputInfo{}.Type(ctx),
			"run_duration":              types.Int64Type,
			"shared_output_info":        CleanRoomTaskRunOutputInfo{}.Type(ctx),
			"start_time":                types.Int64Type,
			"task_run_state":            jobs_tf.CleanRoomTaskRunState{}.Type(ctx),
			"task_type":                 types.StringType,
		},
	}
}

// GetAnalysisDetails returns the value of the AnalysisDetails field in CleanRoomTaskRun as
// a CleanRoomTaskRunCleanRoomTaskAnalysisDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomTaskRun) GetAnalysisDetails(ctx context.Context) (CleanRoomTaskRunCleanRoomTaskAnalysisDetails, bool) {
	var e CleanRoomTaskRunCleanRoomTaskAnalysisDetails
	if m.AnalysisDetails.IsNull() || m.AnalysisDetails.IsUnknown() {
		return e, false
	}
	var v CleanRoomTaskRunCleanRoomTaskAnalysisDetails
	d := m.AnalysisDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnalysisDetails sets the value of the AnalysisDetails field in CleanRoomTaskRun.
func (m *CleanRoomTaskRun) SetAnalysisDetails(ctx context.Context, v CleanRoomTaskRunCleanRoomTaskAnalysisDetails) {
	vs := v.ToObjectValue(ctx)
	m.AnalysisDetails = vs
}

// GetCollaboratorJobRunInfo returns the value of the CollaboratorJobRunInfo field in CleanRoomTaskRun as
// a CollaboratorJobRunInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomTaskRun) GetCollaboratorJobRunInfo(ctx context.Context) (CollaboratorJobRunInfo, bool) {
	var e CollaboratorJobRunInfo
	if m.CollaboratorJobRunInfo.IsNull() || m.CollaboratorJobRunInfo.IsUnknown() {
		return e, false
	}
	var v CollaboratorJobRunInfo
	d := m.CollaboratorJobRunInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCollaboratorJobRunInfo sets the value of the CollaboratorJobRunInfo field in CleanRoomTaskRun.
func (m *CleanRoomTaskRun) SetCollaboratorJobRunInfo(ctx context.Context, v CollaboratorJobRunInfo) {
	vs := v.ToObjectValue(ctx)
	m.CollaboratorJobRunInfo = vs
}

// GetOutputInfo returns the value of the OutputInfo field in CleanRoomTaskRun as
// a CleanRoomTaskRunOutputInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomTaskRun) GetOutputInfo(ctx context.Context) (CleanRoomTaskRunOutputInfo, bool) {
	var e CleanRoomTaskRunOutputInfo
	if m.OutputInfo.IsNull() || m.OutputInfo.IsUnknown() {
		return e, false
	}
	var v CleanRoomTaskRunOutputInfo
	d := m.OutputInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputInfo sets the value of the OutputInfo field in CleanRoomTaskRun.
func (m *CleanRoomTaskRun) SetOutputInfo(ctx context.Context, v CleanRoomTaskRunOutputInfo) {
	vs := v.ToObjectValue(ctx)
	m.OutputInfo = vs
}

// GetSharedOutputInfo returns the value of the SharedOutputInfo field in CleanRoomTaskRun as
// a CleanRoomTaskRunOutputInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomTaskRun) GetSharedOutputInfo(ctx context.Context) (CleanRoomTaskRunOutputInfo, bool) {
	var e CleanRoomTaskRunOutputInfo
	if m.SharedOutputInfo.IsNull() || m.SharedOutputInfo.IsUnknown() {
		return e, false
	}
	var v CleanRoomTaskRunOutputInfo
	d := m.SharedOutputInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSharedOutputInfo sets the value of the SharedOutputInfo field in CleanRoomTaskRun.
func (m *CleanRoomTaskRun) SetSharedOutputInfo(ctx context.Context, v CleanRoomTaskRunOutputInfo) {
	vs := v.ToObjectValue(ctx)
	m.SharedOutputInfo = vs
}

// GetTaskRunState returns the value of the TaskRunState field in CleanRoomTaskRun as
// a jobs_tf.CleanRoomTaskRunState value.
// If the field is unknown or null, the boolean return value is false.
func (m *CleanRoomTaskRun) GetTaskRunState(ctx context.Context) (jobs_tf.CleanRoomTaskRunState, bool) {
	var e jobs_tf.CleanRoomTaskRunState
	if m.TaskRunState.IsNull() || m.TaskRunState.IsUnknown() {
		return e, false
	}
	var v jobs_tf.CleanRoomTaskRunState
	d := m.TaskRunState.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTaskRunState sets the value of the TaskRunState field in CleanRoomTaskRun.
func (m *CleanRoomTaskRun) SetTaskRunState(ctx context.Context, v jobs_tf.CleanRoomTaskRunState) {
	vs := v.ToObjectValue(ctx)
	m.TaskRunState = vs
}

type CleanRoomTaskRunCleanRoomTaskAnalysisDetails struct {
	// Etag of the asset executed in this task run, used to identify the asset
	// version.
	Etag types.String `tfsdk:"etag"`
	// The timestamp of when the asset was last updated.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
}

func (to *CleanRoomTaskRunCleanRoomTaskAnalysisDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomTaskRunCleanRoomTaskAnalysisDetails) {
}

func (to *CleanRoomTaskRunCleanRoomTaskAnalysisDetails) SyncFieldsDuringRead(ctx context.Context, from CleanRoomTaskRunCleanRoomTaskAnalysisDetails) {
}

func (m CleanRoomTaskRunCleanRoomTaskAnalysisDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomTaskRunCleanRoomTaskAnalysisDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CleanRoomTaskRunCleanRoomTaskAnalysisDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomTaskRunCleanRoomTaskAnalysisDetails
// only implements ToObjectValue() and Type().
func (m CleanRoomTaskRunCleanRoomTaskAnalysisDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":       m.Etag,
			"updated_at": m.UpdatedAt,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomTaskRunCleanRoomTaskAnalysisDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":       types.StringType,
			"updated_at": types.Int64Type,
		},
	}
}

type CleanRoomTaskRunOutputInfo struct {
	// Expiration time of the output schema of the task run (if any), in epoch
	// milliseconds.
	OutputSchemaExpirationTime types.Int64 `tfsdk:"output_schema_expiration_time"`
	// Name of the output schema associated with the clean room task run.
	OutputSchemaName types.String `tfsdk:"output_schema_name"`
}

func (to *CleanRoomTaskRunOutputInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CleanRoomTaskRunOutputInfo) {
}

func (to *CleanRoomTaskRunOutputInfo) SyncFieldsDuringRead(ctx context.Context, from CleanRoomTaskRunOutputInfo) {
}

func (m CleanRoomTaskRunOutputInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["output_schema_expiration_time"] = attrs["output_schema_expiration_time"].SetOptional()
	attrs["output_schema_name"] = attrs["output_schema_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CleanRoomTaskRunOutputInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CleanRoomTaskRunOutputInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CleanRoomTaskRunOutputInfo
// only implements ToObjectValue() and Type().
func (m CleanRoomTaskRunOutputInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"output_schema_expiration_time": m.OutputSchemaExpirationTime,
			"output_schema_name":            m.OutputSchemaName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CleanRoomTaskRunOutputInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"output_schema_expiration_time": types.Int64Type,
			"output_schema_name":            types.StringType,
		},
	}
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

func (to *CollaboratorJobRunInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CollaboratorJobRunInfo) {
}

func (to *CollaboratorJobRunInfo) SyncFieldsDuringRead(ctx context.Context, from CollaboratorJobRunInfo) {
}

func (m CollaboratorJobRunInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CollaboratorJobRunInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CollaboratorJobRunInfo
// only implements ToObjectValue() and Type().
func (m CollaboratorJobRunInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CollaboratorJobRunInfo) Type(ctx context.Context) attr.Type {
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

func (to *ComplianceSecurityProfile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComplianceSecurityProfile) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (to *ComplianceSecurityProfile) SyncFieldsDuringRead(ctx context.Context, from ComplianceSecurityProfile) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (m ComplianceSecurityProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ComplianceSecurityProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile
// only implements ToObjectValue() and Type().
func (m ComplianceSecurityProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": m.ComplianceStandards,
			"is_enabled":           m.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComplianceSecurityProfile) Type(ctx context.Context) attr.Type {
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
func (m *ComplianceSecurityProfile) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
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

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile.
func (m *ComplianceSecurityProfile) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ComplianceStandards = types.ListValueMust(t, vs)
}

type CreateCleanRoomAssetRequest struct {
	Asset types.Object `tfsdk:"asset"`
	// The name of the clean room this asset belongs to. This field is required
	// for create operations and populated by the server for responses.
	CleanRoomName types.String `tfsdk:"-"`
}

func (to *CreateCleanRoomAssetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAssetRequest) {
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

func (to *CreateCleanRoomAssetRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAssetRequest) {
	if !from.Asset.IsNull() && !from.Asset.IsUnknown() {
		if toAsset, ok := to.GetAsset(ctx); ok {
			if fromAsset, ok := from.GetAsset(ctx); ok {
				toAsset.SyncFieldsDuringRead(ctx, fromAsset)
				to.SetAsset(ctx, toAsset)
			}
		}
	}
}

func (m CreateCleanRoomAssetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["asset"] = attrs["asset"].SetRequired()
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
func (m CreateCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset":           m.Asset,
			"clean_room_name": m.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomAssetRequest) GetAsset(ctx context.Context) (CleanRoomAsset, bool) {
	var e CleanRoomAsset
	if m.Asset.IsNull() || m.Asset.IsUnknown() {
		return e, false
	}
	var v CleanRoomAsset
	d := m.Asset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAsset sets the value of the Asset field in CreateCleanRoomAssetRequest.
func (m *CreateCleanRoomAssetRequest) SetAsset(ctx context.Context, v CleanRoomAsset) {
	vs := v.ToObjectValue(ctx)
	m.Asset = vs
}

type CreateCleanRoomAssetReviewRequest struct {
	// Asset type. Can either be NOTEBOOK_FILE or JAR_ANALYSIS.
	AssetType types.String `tfsdk:"-"`
	// Name of the clean room
	CleanRoomName types.String `tfsdk:"-"`

	JarAnalysisReview types.Object `tfsdk:"jar_analysis_review"`
	// Name of the asset
	Name types.String `tfsdk:"-"`

	NotebookReview types.Object `tfsdk:"notebook_review"`
}

func (to *CreateCleanRoomAssetReviewRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAssetReviewRequest) {
	if !from.JarAnalysisReview.IsNull() && !from.JarAnalysisReview.IsUnknown() {
		if toJarAnalysisReview, ok := to.GetJarAnalysisReview(ctx); ok {
			if fromJarAnalysisReview, ok := from.GetJarAnalysisReview(ctx); ok {
				// Recursively sync the fields of JarAnalysisReview
				toJarAnalysisReview.SyncFieldsDuringCreateOrUpdate(ctx, fromJarAnalysisReview)
				to.SetJarAnalysisReview(ctx, toJarAnalysisReview)
			}
		}
	}
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

func (to *CreateCleanRoomAssetReviewRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAssetReviewRequest) {
	if !from.JarAnalysisReview.IsNull() && !from.JarAnalysisReview.IsUnknown() {
		if toJarAnalysisReview, ok := to.GetJarAnalysisReview(ctx); ok {
			if fromJarAnalysisReview, ok := from.GetJarAnalysisReview(ctx); ok {
				toJarAnalysisReview.SyncFieldsDuringRead(ctx, fromJarAnalysisReview)
				to.SetJarAnalysisReview(ctx, toJarAnalysisReview)
			}
		}
	}
	if !from.NotebookReview.IsNull() && !from.NotebookReview.IsUnknown() {
		if toNotebookReview, ok := to.GetNotebookReview(ctx); ok {
			if fromNotebookReview, ok := from.GetNotebookReview(ctx); ok {
				toNotebookReview.SyncFieldsDuringRead(ctx, fromNotebookReview)
				to.SetNotebookReview(ctx, toNotebookReview)
			}
		}
	}
}

func (m CreateCleanRoomAssetReviewRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["jar_analysis_review"] = attrs["jar_analysis_review"].SetOptional()
	attrs["notebook_review"] = attrs["notebook_review"].SetOptional()
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
func (m CreateCleanRoomAssetReviewRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jar_analysis_review": reflect.TypeOf(JarAnalysisVersionReview{}),
		"notebook_review":     reflect.TypeOf(NotebookVersionReview{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewRequest
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAssetReviewRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":          m.AssetType,
			"clean_room_name":     m.CleanRoomName,
			"jar_analysis_review": m.JarAnalysisReview,
			"name":                m.Name,
			"notebook_review":     m.NotebookReview,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAssetReviewRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"asset_type":          types.StringType,
			"clean_room_name":     types.StringType,
			"jar_analysis_review": JarAnalysisVersionReview{}.Type(ctx),
			"name":                types.StringType,
			"notebook_review":     NotebookVersionReview{}.Type(ctx),
		},
	}
}

// GetJarAnalysisReview returns the value of the JarAnalysisReview field in CreateCleanRoomAssetReviewRequest as
// a JarAnalysisVersionReview value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCleanRoomAssetReviewRequest) GetJarAnalysisReview(ctx context.Context) (JarAnalysisVersionReview, bool) {
	var e JarAnalysisVersionReview
	if m.JarAnalysisReview.IsNull() || m.JarAnalysisReview.IsUnknown() {
		return e, false
	}
	var v JarAnalysisVersionReview
	d := m.JarAnalysisReview.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarAnalysisReview sets the value of the JarAnalysisReview field in CreateCleanRoomAssetReviewRequest.
func (m *CreateCleanRoomAssetReviewRequest) SetJarAnalysisReview(ctx context.Context, v JarAnalysisVersionReview) {
	vs := v.ToObjectValue(ctx)
	m.JarAnalysisReview = vs
}

// GetNotebookReview returns the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest as
// a NotebookVersionReview value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCleanRoomAssetReviewRequest) GetNotebookReview(ctx context.Context) (NotebookVersionReview, bool) {
	var e NotebookVersionReview
	if m.NotebookReview.IsNull() || m.NotebookReview.IsUnknown() {
		return e, false
	}
	var v NotebookVersionReview
	d := m.NotebookReview.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookReview sets the value of the NotebookReview field in CreateCleanRoomAssetReviewRequest.
func (m *CreateCleanRoomAssetReviewRequest) SetNotebookReview(ctx context.Context, v NotebookVersionReview) {
	vs := v.ToObjectValue(ctx)
	m.NotebookReview = vs
}

type CreateCleanRoomAssetReviewResponse struct {
	// top-level status derived from all reviews
	JarAnalysisReviewState types.String `tfsdk:"jar_analysis_review_state"`
	// All existing jar analysis approvals or rejections
	JarAnalysisReviews types.List `tfsdk:"jar_analysis_reviews"`
	// Top-level status derived from all reviews
	NotebookReviewState types.String `tfsdk:"notebook_review_state"`
	// All existing notebook approvals or rejections
	NotebookReviews types.List `tfsdk:"notebook_reviews"`
}

func (to *CreateCleanRoomAssetReviewResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAssetReviewResponse) {
	if !from.JarAnalysisReviews.IsNull() && !from.JarAnalysisReviews.IsUnknown() && to.JarAnalysisReviews.IsNull() && len(from.JarAnalysisReviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JarAnalysisReviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JarAnalysisReviews = from.JarAnalysisReviews
	}
	if !from.NotebookReviews.IsNull() && !from.NotebookReviews.IsUnknown() && to.NotebookReviews.IsNull() && len(from.NotebookReviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for NotebookReviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.NotebookReviews = from.NotebookReviews
	}
}

func (to *CreateCleanRoomAssetReviewResponse) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAssetReviewResponse) {
	if !from.JarAnalysisReviews.IsNull() && !from.JarAnalysisReviews.IsUnknown() && to.JarAnalysisReviews.IsNull() && len(from.JarAnalysisReviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for JarAnalysisReviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.JarAnalysisReviews = from.JarAnalysisReviews
	}
	if !from.NotebookReviews.IsNull() && !from.NotebookReviews.IsUnknown() && to.NotebookReviews.IsNull() && len(from.NotebookReviews.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for NotebookReviews, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.NotebookReviews = from.NotebookReviews
	}
}

func (m CreateCleanRoomAssetReviewResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["jar_analysis_review_state"] = attrs["jar_analysis_review_state"].SetOptional()
	attrs["jar_analysis_reviews"] = attrs["jar_analysis_reviews"].SetOptional()
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
func (m CreateCleanRoomAssetReviewResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"jar_analysis_reviews": reflect.TypeOf(CleanRoomJarAnalysisReview{}),
		"notebook_reviews":     reflect.TypeOf(CleanRoomNotebookReview{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAssetReviewResponse
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAssetReviewResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"jar_analysis_review_state": m.JarAnalysisReviewState,
			"jar_analysis_reviews":      m.JarAnalysisReviews,
			"notebook_review_state":     m.NotebookReviewState,
			"notebook_reviews":          m.NotebookReviews,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAssetReviewResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"jar_analysis_review_state": types.StringType,
			"jar_analysis_reviews": basetypes.ListType{
				ElemType: CleanRoomJarAnalysisReview{}.Type(ctx),
			},
			"notebook_review_state": types.StringType,
			"notebook_reviews": basetypes.ListType{
				ElemType: CleanRoomNotebookReview{}.Type(ctx),
			},
		},
	}
}

// GetJarAnalysisReviews returns the value of the JarAnalysisReviews field in CreateCleanRoomAssetReviewResponse as
// a slice of CleanRoomJarAnalysisReview values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCleanRoomAssetReviewResponse) GetJarAnalysisReviews(ctx context.Context) ([]CleanRoomJarAnalysisReview, bool) {
	if m.JarAnalysisReviews.IsNull() || m.JarAnalysisReviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomJarAnalysisReview
	d := m.JarAnalysisReviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetJarAnalysisReviews sets the value of the JarAnalysisReviews field in CreateCleanRoomAssetReviewResponse.
func (m *CreateCleanRoomAssetReviewResponse) SetJarAnalysisReviews(ctx context.Context, v []CleanRoomJarAnalysisReview) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["jar_analysis_reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.JarAnalysisReviews = types.ListValueMust(t, vs)
}

// GetNotebookReviews returns the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse as
// a slice of CleanRoomNotebookReview values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCleanRoomAssetReviewResponse) GetNotebookReviews(ctx context.Context) ([]CleanRoomNotebookReview, bool) {
	if m.NotebookReviews.IsNull() || m.NotebookReviews.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookReview
	d := m.NotebookReviews.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNotebookReviews sets the value of the NotebookReviews field in CreateCleanRoomAssetReviewResponse.
func (m *CreateCleanRoomAssetReviewResponse) SetNotebookReviews(ctx context.Context, v []CleanRoomNotebookReview) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["notebook_reviews"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.NotebookReviews = types.ListValueMust(t, vs)
}

type CreateCleanRoomAutoApprovalRuleRequest struct {
	AutoApprovalRule types.Object `tfsdk:"auto_approval_rule"`
	// The name of the clean room this auto-approval rule belongs to.
	CleanRoomName types.String `tfsdk:"-"`
}

func (to *CreateCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomAutoApprovalRuleRequest) {
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

func (to *CreateCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomAutoApprovalRuleRequest) {
	if !from.AutoApprovalRule.IsNull() && !from.AutoApprovalRule.IsUnknown() {
		if toAutoApprovalRule, ok := to.GetAutoApprovalRule(ctx); ok {
			if fromAutoApprovalRule, ok := from.GetAutoApprovalRule(ctx); ok {
				toAutoApprovalRule.SyncFieldsDuringRead(ctx, fromAutoApprovalRule)
				to.SetAutoApprovalRule(ctx, toAutoApprovalRule)
			}
		}
	}
}

func (m CreateCleanRoomAutoApprovalRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_approval_rule"] = attrs["auto_approval_rule"].SetRequired()
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
func (m CreateCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": m.AutoApprovalRule,
			"clean_room_name":    m.CleanRoomName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomAutoApprovalRuleRequest) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule, bool) {
	var e CleanRoomAutoApprovalRule
	if m.AutoApprovalRule.IsNull() || m.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v CleanRoomAutoApprovalRule
	d := m.AutoApprovalRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in CreateCleanRoomAutoApprovalRuleRequest.
func (m *CreateCleanRoomAutoApprovalRuleRequest) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule) {
	vs := v.ToObjectValue(ctx)
	m.AutoApprovalRule = vs
}

type CreateCleanRoomOutputCatalogRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`

	OutputCatalog types.Object `tfsdk:"output_catalog"`
}

func (to *CreateCleanRoomOutputCatalogRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomOutputCatalogRequest) {
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

func (to *CreateCleanRoomOutputCatalogRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomOutputCatalogRequest) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				toOutputCatalog.SyncFieldsDuringRead(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
}

func (m CreateCleanRoomOutputCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["output_catalog"] = attrs["output_catalog"].SetRequired()
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
func (m CreateCleanRoomOutputCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogRequest
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomOutputCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"output_catalog":  m.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomOutputCatalogRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateCleanRoomOutputCatalogRequest) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog, bool) {
	var e CleanRoomOutputCatalog
	if m.OutputCatalog.IsNull() || m.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v CleanRoomOutputCatalog
	d := m.OutputCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogRequest.
func (m *CreateCleanRoomOutputCatalogRequest) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := v.ToObjectValue(ctx)
	m.OutputCatalog = vs
}

type CreateCleanRoomOutputCatalogResponse struct {
	OutputCatalog types.Object `tfsdk:"output_catalog"`
}

func (to *CreateCleanRoomOutputCatalogResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomOutputCatalogResponse) {
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

func (to *CreateCleanRoomOutputCatalogResponse) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomOutputCatalogResponse) {
	if !from.OutputCatalog.IsNull() && !from.OutputCatalog.IsUnknown() {
		if toOutputCatalog, ok := to.GetOutputCatalog(ctx); ok {
			if fromOutputCatalog, ok := from.GetOutputCatalog(ctx); ok {
				toOutputCatalog.SyncFieldsDuringRead(ctx, fromOutputCatalog)
				to.SetOutputCatalog(ctx, toOutputCatalog)
			}
		}
	}
}

func (m CreateCleanRoomOutputCatalogResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["output_catalog"] = attrs["output_catalog"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomOutputCatalogResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCleanRoomOutputCatalogResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"output_catalog": reflect.TypeOf(CleanRoomOutputCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomOutputCatalogResponse
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomOutputCatalogResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"output_catalog": m.OutputCatalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomOutputCatalogResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"output_catalog": CleanRoomOutputCatalog{}.Type(ctx),
		},
	}
}

// GetOutputCatalog returns the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse as
// a CleanRoomOutputCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCleanRoomOutputCatalogResponse) GetOutputCatalog(ctx context.Context) (CleanRoomOutputCatalog, bool) {
	var e CleanRoomOutputCatalog
	if m.OutputCatalog.IsNull() || m.OutputCatalog.IsUnknown() {
		return e, false
	}
	var v CleanRoomOutputCatalog
	d := m.OutputCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutputCatalog sets the value of the OutputCatalog field in CreateCleanRoomOutputCatalogResponse.
func (m *CreateCleanRoomOutputCatalogResponse) SetOutputCatalog(ctx context.Context, v CleanRoomOutputCatalog) {
	vs := v.ToObjectValue(ctx)
	m.OutputCatalog = vs
}

type CreateCleanRoomRequest struct {
	CleanRoom types.Object `tfsdk:"clean_room"`
}

func (to *CreateCleanRoomRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCleanRoomRequest) {
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

func (to *CreateCleanRoomRequest) SyncFieldsDuringRead(ctx context.Context, from CreateCleanRoomRequest) {
	if !from.CleanRoom.IsNull() && !from.CleanRoom.IsUnknown() {
		if toCleanRoom, ok := to.GetCleanRoom(ctx); ok {
			if fromCleanRoom, ok := from.GetCleanRoom(ctx); ok {
				toCleanRoom.SyncFieldsDuringRead(ctx, fromCleanRoom)
				to.SetCleanRoom(ctx, toCleanRoom)
			}
		}
	}
}

func (m CreateCleanRoomRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room"] = attrs["clean_room"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCleanRoomRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCleanRoomRequest
// only implements ToObjectValue() and Type().
func (m CreateCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": m.CleanRoom,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCleanRoomRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room": CleanRoom{}.Type(ctx),
		},
	}
}

// GetCleanRoom returns the value of the CleanRoom field in CreateCleanRoomRequest as
// a CleanRoom value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCleanRoomRequest) GetCleanRoom(ctx context.Context) (CleanRoom, bool) {
	var e CleanRoom
	if m.CleanRoom.IsNull() || m.CleanRoom.IsUnknown() {
		return e, false
	}
	var v CleanRoom
	d := m.CleanRoom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRoom sets the value of the CleanRoom field in CreateCleanRoomRequest.
func (m *CreateCleanRoomRequest) SetCleanRoom(ctx context.Context, v CleanRoom) {
	vs := v.ToObjectValue(ctx)
	m.CleanRoom = vs
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

func (to *DeleteCleanRoomAssetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomAssetRequest) {
}

func (to *DeleteCleanRoomAssetRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomAssetRequest) {
}

func (m DeleteCleanRoomAssetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
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

func (to *DeleteCleanRoomAssetResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomAssetResponse) {
}

func (to *DeleteCleanRoomAssetResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomAssetResponse) {
}

func (m DeleteCleanRoomAssetResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCleanRoomAssetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCleanRoomAssetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAssetResponse
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomAssetResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomAssetResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteCleanRoomAutoApprovalRuleRequest struct {
	CleanRoomName types.String `tfsdk:"-"`

	RuleId types.String `tfsdk:"-"`
}

func (to *DeleteCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomAutoApprovalRuleRequest) {
}

func (to *DeleteCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomAutoApprovalRuleRequest) {
}

func (m DeleteCleanRoomAutoApprovalRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"rule_id":         m.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
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

func (to *DeleteCleanRoomRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCleanRoomRequest) {
}

func (to *DeleteCleanRoomRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCleanRoomRequest) {
}

func (m DeleteCleanRoomRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCleanRoomRequest
// only implements ToObjectValue() and Type().
func (m DeleteCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCleanRoomRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomAssetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomAssetRequest) {
}

func (to *GetCleanRoomAssetRequest) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomAssetRequest) {
}

func (m GetCleanRoomAssetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (m GetCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"asset_type":      m.AssetType,
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomAssetRevisionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomAssetRevisionRequest) {
}

func (to *GetCleanRoomAssetRevisionRequest) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomAssetRevisionRequest) {
}

func (m GetCleanRoomAssetRevisionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomAssetRevisionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAssetRevisionRequest
// only implements ToObjectValue() and Type().
func (m GetCleanRoomAssetRevisionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetCleanRoomAssetRevisionRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomAutoApprovalRuleRequest) {
}

func (to *GetCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomAutoApprovalRuleRequest) {
}

func (m GetCleanRoomAutoApprovalRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (m GetCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"rule_id":         m.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetCleanRoomRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCleanRoomRequest) {
}

func (to *GetCleanRoomRequest) SyncFieldsDuringRead(ctx context.Context, from GetCleanRoomRequest) {
}

func (m GetCleanRoomRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCleanRoomRequest
// only implements ToObjectValue() and Type().
func (m GetCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCleanRoomRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type JarAnalysisVersionReview struct {
	// Review comment
	Comment types.String `tfsdk:"comment"`
	// Etag identifying the jar analysis version, with its value being a hash of
	// an internally-generated UUID
	Etag types.String `tfsdk:"etag"`
	// Review outcome
	ReviewState types.String `tfsdk:"review_state"`
}

func (to *JarAnalysisVersionReview) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from JarAnalysisVersionReview) {
}

func (to *JarAnalysisVersionReview) SyncFieldsDuringRead(ctx context.Context, from JarAnalysisVersionReview) {
}

func (m JarAnalysisVersionReview) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["etag"] = attrs["etag"].SetRequired()
	attrs["review_state"] = attrs["review_state"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in JarAnalysisVersionReview.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m JarAnalysisVersionReview) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, JarAnalysisVersionReview
// only implements ToObjectValue() and Type().
func (m JarAnalysisVersionReview) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      m.Comment,
			"etag":         m.Etag,
			"review_state": m.ReviewState,
		})
}

// Type implements basetypes.ObjectValuable.
func (m JarAnalysisVersionReview) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":      types.StringType,
			"etag":         types.StringType,
			"review_state": types.StringType,
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

func (to *ListCleanRoomAssetRevisionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetRevisionsRequest) {
}

func (to *ListCleanRoomAssetRevisionsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetRevisionsRequest) {
}

func (m ListCleanRoomAssetRevisionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetRevisionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsRequest
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetRevisionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListCleanRoomAssetRevisionsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAssetRevisionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetRevisionsResponse) {
	if !from.Revisions.IsNull() && !from.Revisions.IsUnknown() && to.Revisions.IsNull() && len(from.Revisions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Revisions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Revisions = from.Revisions
	}
}

func (to *ListCleanRoomAssetRevisionsResponse) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetRevisionsResponse) {
	if !from.Revisions.IsNull() && !from.Revisions.IsUnknown() && to.Revisions.IsNull() && len(from.Revisions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Revisions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Revisions = from.Revisions
	}
}

func (m ListCleanRoomAssetRevisionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetRevisionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"revisions": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetRevisionsResponse
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetRevisionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"revisions":       m.Revisions,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetRevisionsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomAssetRevisionsResponse) GetRevisions(ctx context.Context) ([]CleanRoomAsset, bool) {
	if m.Revisions.IsNull() || m.Revisions.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset
	d := m.Revisions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRevisions sets the value of the Revisions field in ListCleanRoomAssetRevisionsResponse.
func (m *ListCleanRoomAssetRevisionsResponse) SetRevisions(ctx context.Context, v []CleanRoomAsset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["revisions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Revisions = types.ListValueMust(t, vs)
}

type ListCleanRoomAssetsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCleanRoomAssetsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetsRequest) {
}

func (to *ListCleanRoomAssetsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetsRequest) {
}

func (m ListCleanRoomAssetsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsRequest
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAssetsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAssetsResponse) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
}

func (to *ListCleanRoomAssetsResponse) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAssetsResponse) {
	if !from.Assets.IsNull() && !from.Assets.IsUnknown() && to.Assets.IsNull() && len(from.Assets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Assets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Assets = from.Assets
	}
}

func (m ListCleanRoomAssetsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAssetsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"assets": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAssetsResponse
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAssetsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"assets":          m.Assets,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAssetsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomAssetsResponse) GetAssets(ctx context.Context) ([]CleanRoomAsset, bool) {
	if m.Assets.IsNull() || m.Assets.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAsset
	d := m.Assets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAssets sets the value of the Assets field in ListCleanRoomAssetsResponse.
func (m *ListCleanRoomAssetsResponse) SetAssets(ctx context.Context, v []CleanRoomAsset) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["assets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Assets = types.ListValueMust(t, vs)
}

type ListCleanRoomAutoApprovalRulesRequest struct {
	CleanRoomName types.String `tfsdk:"-"`
	// Maximum number of auto-approval rules to return. Defaults to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCleanRoomAutoApprovalRulesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAutoApprovalRulesRequest) {
}

func (to *ListCleanRoomAutoApprovalRulesRequest) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAutoApprovalRulesRequest) {
}

func (m ListCleanRoomAutoApprovalRulesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAutoApprovalRulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesRequest
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAutoApprovalRulesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"page_size":       m.PageSize,
			"page_token":      m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAutoApprovalRulesRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomAutoApprovalRulesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomAutoApprovalRulesResponse) {
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() && to.Rules.IsNull() && len(from.Rules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Rules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Rules = from.Rules
	}
}

func (to *ListCleanRoomAutoApprovalRulesResponse) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomAutoApprovalRulesResponse) {
	if !from.Rules.IsNull() && !from.Rules.IsUnknown() && to.Rules.IsNull() && len(from.Rules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Rules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Rules = from.Rules
	}
}

func (m ListCleanRoomAutoApprovalRulesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomAutoApprovalRulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rules": reflect.TypeOf(CleanRoomAutoApprovalRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomAutoApprovalRulesResponse
// only implements ToObjectValue() and Type().
func (m ListCleanRoomAutoApprovalRulesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"rules":           m.Rules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomAutoApprovalRulesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomAutoApprovalRulesResponse) GetRules(ctx context.Context) ([]CleanRoomAutoApprovalRule, bool) {
	if m.Rules.IsNull() || m.Rules.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomAutoApprovalRule
	d := m.Rules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRules sets the value of the Rules field in ListCleanRoomAutoApprovalRulesResponse.
func (m *ListCleanRoomAutoApprovalRulesResponse) SetRules(ctx context.Context, v []CleanRoomAutoApprovalRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Rules = types.ListValueMust(t, vs)
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

func (to *ListCleanRoomNotebookTaskRunsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomNotebookTaskRunsRequest) {
}

func (to *ListCleanRoomNotebookTaskRunsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomNotebookTaskRunsRequest) {
}

func (m ListCleanRoomNotebookTaskRunsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomNotebookTaskRunsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsRequest
// only implements ToObjectValue() and Type().
func (m ListCleanRoomNotebookTaskRunsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ListCleanRoomNotebookTaskRunsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomNotebookTaskRunsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomNotebookTaskRunsResponse) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (to *ListCleanRoomNotebookTaskRunsResponse) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomNotebookTaskRunsResponse) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (m ListCleanRoomNotebookTaskRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomNotebookTaskRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(CleanRoomNotebookTaskRun{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomNotebookTaskRunsResponse
// only implements ToObjectValue() and Type().
func (m ListCleanRoomNotebookTaskRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"runs":            m.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomNotebookTaskRunsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomNotebookTaskRunsResponse) GetRuns(ctx context.Context) ([]CleanRoomNotebookTaskRun, bool) {
	if m.Runs.IsNull() || m.Runs.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomNotebookTaskRun
	d := m.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in ListCleanRoomNotebookTaskRunsResponse.
func (m *ListCleanRoomNotebookTaskRunsResponse) SetRuns(ctx context.Context, v []CleanRoomNotebookTaskRun) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Runs = types.ListValueMust(t, vs)
}

type ListCleanRoomTaskRunsRequest struct {
	// Name of the clean room.
	CleanRoomName types.String `tfsdk:"-"`
	// Executable name.
	Name types.String `tfsdk:"-"`
	// The maximum number of task runs to return. Maximum value of 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
	// Filter by the type of Clean Room task.
	TaskType types.String `tfsdk:"-"`
}

func (to *ListCleanRoomTaskRunsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomTaskRunsRequest) {
}

func (to *ListCleanRoomTaskRunsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomTaskRunsRequest) {
}

func (m ListCleanRoomTaskRunsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room_name"] = attrs["clean_room_name"].SetRequired()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["task_type"] = attrs["task_type"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomTaskRunsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCleanRoomTaskRunsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomTaskRunsRequest
// only implements ToObjectValue() and Type().
func (m ListCleanRoomTaskRunsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room_name": m.CleanRoomName,
			"name":            m.Name,
			"page_size":       m.PageSize,
			"page_token":      m.PageToken,
			"task_type":       m.TaskType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomTaskRunsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"clean_room_name": types.StringType,
			"name":            types.StringType,
			"page_size":       types.Int64Type,
			"page_token":      types.StringType,
			"task_type":       types.StringType,
		},
	}
}

type ListCleanRoomTaskRunsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. page_token should be set to this value for the next request
	// (for the next page of results).
	NextPageToken types.String `tfsdk:"next_page_token"`
	// Task runs in the clean room.
	Runs types.List `tfsdk:"runs"`
}

func (to *ListCleanRoomTaskRunsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomTaskRunsResponse) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (to *ListCleanRoomTaskRunsResponse) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomTaskRunsResponse) {
	if !from.Runs.IsNull() && !from.Runs.IsUnknown() && to.Runs.IsNull() && len(from.Runs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Runs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Runs = from.Runs
	}
}

func (m ListCleanRoomTaskRunsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["runs"] = attrs["runs"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCleanRoomTaskRunsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCleanRoomTaskRunsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"runs": reflect.TypeOf(CleanRoomTaskRun{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomTaskRunsResponse
// only implements ToObjectValue() and Type().
func (m ListCleanRoomTaskRunsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"runs":            m.Runs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomTaskRunsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"runs": basetypes.ListType{
				ElemType: CleanRoomTaskRun{}.Type(ctx),
			},
		},
	}
}

// GetRuns returns the value of the Runs field in ListCleanRoomTaskRunsResponse as
// a slice of CleanRoomTaskRun values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListCleanRoomTaskRunsResponse) GetRuns(ctx context.Context) ([]CleanRoomTaskRun, bool) {
	if m.Runs.IsNull() || m.Runs.IsUnknown() {
		return nil, false
	}
	var v []CleanRoomTaskRun
	d := m.Runs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRuns sets the value of the Runs field in ListCleanRoomTaskRunsResponse.
func (m *ListCleanRoomTaskRunsResponse) SetRuns(ctx context.Context, v []CleanRoomTaskRun) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["runs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Runs = types.ListValueMust(t, vs)
}

type ListCleanRoomsRequest struct {
	// Maximum number of clean rooms to return (i.e., the page length). Defaults
	// to 100.
	PageSize types.Int64 `tfsdk:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListCleanRoomsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomsRequest) {
}

func (to *ListCleanRoomsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomsRequest) {
}

func (m ListCleanRoomsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsRequest
// only implements ToObjectValue() and Type().
func (m ListCleanRoomsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListCleanRoomsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCleanRoomsResponse) {
	if !from.CleanRooms.IsNull() && !from.CleanRooms.IsUnknown() && to.CleanRooms.IsNull() && len(from.CleanRooms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CleanRooms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CleanRooms = from.CleanRooms
	}
}

func (to *ListCleanRoomsResponse) SyncFieldsDuringRead(ctx context.Context, from ListCleanRoomsResponse) {
	if !from.CleanRooms.IsNull() && !from.CleanRooms.IsUnknown() && to.CleanRooms.IsNull() && len(from.CleanRooms.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CleanRooms, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CleanRooms = from.CleanRooms
	}
}

func (m ListCleanRoomsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCleanRoomsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_rooms": reflect.TypeOf(CleanRoom{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCleanRoomsResponse
// only implements ToObjectValue() and Type().
func (m ListCleanRoomsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_rooms":     m.CleanRooms,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCleanRoomsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListCleanRoomsResponse) GetCleanRooms(ctx context.Context) ([]CleanRoom, bool) {
	if m.CleanRooms.IsNull() || m.CleanRooms.IsUnknown() {
		return nil, false
	}
	var v []CleanRoom
	d := m.CleanRooms.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRooms sets the value of the CleanRooms field in ListCleanRoomsResponse.
func (m *ListCleanRoomsResponse) SetCleanRooms(ctx context.Context, v []CleanRoom) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["clean_rooms"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CleanRooms = types.ListValueMust(t, vs)
}

type NotebookVersionReview struct {
	// Review comment
	Comment types.String `tfsdk:"comment"`
	// Etag identifying the notebook version
	Etag types.String `tfsdk:"etag"`
	// Review outcome
	ReviewState types.String `tfsdk:"review_state"`
}

func (to *NotebookVersionReview) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotebookVersionReview) {
}

func (to *NotebookVersionReview) SyncFieldsDuringRead(ctx context.Context, from NotebookVersionReview) {
}

func (m NotebookVersionReview) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NotebookVersionReview) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotebookVersionReview
// only implements ToObjectValue() and Type().
func (m NotebookVersionReview) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":      m.Comment,
			"etag":         m.Etag,
			"review_state": m.ReviewState,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotebookVersionReview) Type(ctx context.Context) attr.Type {
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
	// For notebooks, the name is the notebook file name. For jar analyses, the
	// name is the jar analysis name.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateCleanRoomAssetRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCleanRoomAssetRequest) {
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

func (to *UpdateCleanRoomAssetRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateCleanRoomAssetRequest) {
	if !from.Asset.IsNull() && !from.Asset.IsUnknown() {
		if toAsset, ok := to.GetAsset(ctx); ok {
			if fromAsset, ok := from.GetAsset(ctx); ok {
				toAsset.SyncFieldsDuringRead(ctx, fromAsset)
				to.SetAsset(ctx, toAsset)
			}
		}
	}
}

func (m UpdateCleanRoomAssetRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["asset"] = attrs["asset"].SetRequired()
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
func (m UpdateCleanRoomAssetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"asset": reflect.TypeOf(CleanRoomAsset{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAssetRequest
// only implements ToObjectValue() and Type().
func (m UpdateCleanRoomAssetRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateCleanRoomAssetRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateCleanRoomAssetRequest) GetAsset(ctx context.Context) (CleanRoomAsset, bool) {
	var e CleanRoomAsset
	if m.Asset.IsNull() || m.Asset.IsUnknown() {
		return e, false
	}
	var v CleanRoomAsset
	d := m.Asset.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAsset sets the value of the Asset field in UpdateCleanRoomAssetRequest.
func (m *UpdateCleanRoomAssetRequest) SetAsset(ctx context.Context, v CleanRoomAsset) {
	vs := v.ToObjectValue(ctx)
	m.Asset = vs
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

func (to *UpdateCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCleanRoomAutoApprovalRuleRequest) {
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

func (to *UpdateCleanRoomAutoApprovalRuleRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateCleanRoomAutoApprovalRuleRequest) {
	if !from.AutoApprovalRule.IsNull() && !from.AutoApprovalRule.IsUnknown() {
		if toAutoApprovalRule, ok := to.GetAutoApprovalRule(ctx); ok {
			if fromAutoApprovalRule, ok := from.GetAutoApprovalRule(ctx); ok {
				toAutoApprovalRule.SyncFieldsDuringRead(ctx, fromAutoApprovalRule)
				to.SetAutoApprovalRule(ctx, toAutoApprovalRule)
			}
		}
	}
}

func (m UpdateCleanRoomAutoApprovalRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_approval_rule"] = attrs["auto_approval_rule"].SetRequired()
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
func (m UpdateCleanRoomAutoApprovalRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_approval_rule": reflect.TypeOf(CleanRoomAutoApprovalRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomAutoApprovalRuleRequest
// only implements ToObjectValue() and Type().
func (m UpdateCleanRoomAutoApprovalRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_approval_rule": m.AutoApprovalRule,
			"clean_room_name":    m.CleanRoomName,
			"rule_id":            m.RuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCleanRoomAutoApprovalRuleRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateCleanRoomAutoApprovalRuleRequest) GetAutoApprovalRule(ctx context.Context) (CleanRoomAutoApprovalRule, bool) {
	var e CleanRoomAutoApprovalRule
	if m.AutoApprovalRule.IsNull() || m.AutoApprovalRule.IsUnknown() {
		return e, false
	}
	var v CleanRoomAutoApprovalRule
	d := m.AutoApprovalRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoApprovalRule sets the value of the AutoApprovalRule field in UpdateCleanRoomAutoApprovalRuleRequest.
func (m *UpdateCleanRoomAutoApprovalRuleRequest) SetAutoApprovalRule(ctx context.Context, v CleanRoomAutoApprovalRule) {
	vs := v.ToObjectValue(ctx)
	m.AutoApprovalRule = vs
}

type UpdateCleanRoomRequest struct {
	CleanRoom types.Object `tfsdk:"clean_room"`
	// Name of the clean room.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateCleanRoomRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCleanRoomRequest) {
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

func (to *UpdateCleanRoomRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateCleanRoomRequest) {
	if !from.CleanRoom.IsNull() && !from.CleanRoom.IsUnknown() {
		if toCleanRoom, ok := to.GetCleanRoom(ctx); ok {
			if fromCleanRoom, ok := from.GetCleanRoom(ctx); ok {
				toCleanRoom.SyncFieldsDuringRead(ctx, fromCleanRoom)
				to.SetCleanRoom(ctx, toCleanRoom)
			}
		}
	}
}

func (m UpdateCleanRoomRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["clean_room"] = attrs["clean_room"].SetOptional()
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
func (m UpdateCleanRoomRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"clean_room": reflect.TypeOf(CleanRoom{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCleanRoomRequest
// only implements ToObjectValue() and Type().
func (m UpdateCleanRoomRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"clean_room": m.CleanRoom,
			"name":       m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCleanRoomRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateCleanRoomRequest) GetCleanRoom(ctx context.Context) (CleanRoom, bool) {
	var e CleanRoom
	if m.CleanRoom.IsNull() || m.CleanRoom.IsUnknown() {
		return e, false
	}
	var v CleanRoom
	d := m.CleanRoom.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCleanRoom sets the value of the CleanRoom field in UpdateCleanRoomRequest.
func (m *UpdateCleanRoomRequest) SetCleanRoom(ctx context.Context, v CleanRoom) {
	vs := v.ToObjectValue(ctx)
	m.CleanRoom = vs
}
