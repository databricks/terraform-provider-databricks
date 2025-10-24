// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package database_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDatabaseCatalogRequest_SdkV2 struct {
	Catalog types.List `tfsdk:"catalog"`
}

func (to *CreateDatabaseCatalogRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseCatalogRequest_SdkV2) {
	if !from.Catalog.IsNull() && !from.Catalog.IsUnknown() {
		if toCatalog, ok := to.GetCatalog(ctx); ok {
			if fromCatalog, ok := from.GetCatalog(ctx); ok {
				// Recursively sync the fields of Catalog
				toCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromCatalog)
				to.SetCatalog(ctx, toCatalog)
			}
		}
	}
}

func (to *CreateDatabaseCatalogRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseCatalogRequest_SdkV2) {
	if !from.Catalog.IsNull() && !from.Catalog.IsUnknown() {
		if toCatalog, ok := to.GetCatalog(ctx); ok {
			if fromCatalog, ok := from.GetCatalog(ctx); ok {
				toCatalog.SyncFieldsDuringRead(ctx, fromCatalog)
				to.SetCatalog(ctx, toCatalog)
			}
		}
	}
}

func (m CreateDatabaseCatalogRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetRequired()
	attrs["catalog"] = attrs["catalog"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(DatabaseCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": m.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": basetypes.ListType{
				ElemType: DatabaseCatalog_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetCatalog returns the value of the Catalog field in CreateDatabaseCatalogRequest_SdkV2 as
// a DatabaseCatalog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseCatalogRequest_SdkV2) GetCatalog(ctx context.Context) (DatabaseCatalog_SdkV2, bool) {
	var e DatabaseCatalog_SdkV2
	if m.Catalog.IsNull() || m.Catalog.IsUnknown() {
		return e, false
	}
	var v []DatabaseCatalog_SdkV2
	d := m.Catalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCatalog sets the value of the Catalog field in CreateDatabaseCatalogRequest_SdkV2.
func (m *CreateDatabaseCatalogRequest_SdkV2) SetCatalog(ctx context.Context, v DatabaseCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["catalog"]
	m.Catalog = types.ListValueMust(t, vs)
}

type CreateDatabaseInstanceRequest_SdkV2 struct {
	// Instance to create.
	DatabaseInstance types.List `tfsdk:"database_instance"`
}

func (to *CreateDatabaseInstanceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseInstanceRequest_SdkV2) {
	if !from.DatabaseInstance.IsNull() && !from.DatabaseInstance.IsUnknown() {
		if toDatabaseInstance, ok := to.GetDatabaseInstance(ctx); ok {
			if fromDatabaseInstance, ok := from.GetDatabaseInstance(ctx); ok {
				// Recursively sync the fields of DatabaseInstance
				toDatabaseInstance.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseInstance)
				to.SetDatabaseInstance(ctx, toDatabaseInstance)
			}
		}
	}
}

func (to *CreateDatabaseInstanceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseInstanceRequest_SdkV2) {
	if !from.DatabaseInstance.IsNull() && !from.DatabaseInstance.IsUnknown() {
		if toDatabaseInstance, ok := to.GetDatabaseInstance(ctx); ok {
			if fromDatabaseInstance, ok := from.GetDatabaseInstance(ctx); ok {
				toDatabaseInstance.SyncFieldsDuringRead(ctx, fromDatabaseInstance)
				to.SetDatabaseInstance(ctx, toDatabaseInstance)
			}
		}
	}
}

func (m CreateDatabaseInstanceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance"] = attrs["database_instance"].SetRequired()
	attrs["database_instance"] = attrs["database_instance"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": m.DatabaseInstance,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": basetypes.ListType{
				ElemType: DatabaseInstance_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in CreateDatabaseInstanceRequest_SdkV2 as
// a DatabaseInstance_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseInstanceRequest_SdkV2) GetDatabaseInstance(ctx context.Context) (DatabaseInstance_SdkV2, bool) {
	var e DatabaseInstance_SdkV2
	if m.DatabaseInstance.IsNull() || m.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance_SdkV2
	d := m.DatabaseInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in CreateDatabaseInstanceRequest_SdkV2.
func (m *CreateDatabaseInstanceRequest_SdkV2) SetDatabaseInstance(ctx context.Context, v DatabaseInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance"]
	m.DatabaseInstance = types.ListValueMust(t, vs)
}

type CreateDatabaseInstanceRoleRequest_SdkV2 struct {
	DatabaseInstanceName types.String `tfsdk:"-"`

	DatabaseInstanceRole types.List `tfsdk:"database_instance_role"`

	InstanceName types.String `tfsdk:"-"`
}

func (to *CreateDatabaseInstanceRoleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseInstanceRoleRequest_SdkV2) {
	if !from.DatabaseInstanceRole.IsNull() && !from.DatabaseInstanceRole.IsUnknown() {
		if toDatabaseInstanceRole, ok := to.GetDatabaseInstanceRole(ctx); ok {
			if fromDatabaseInstanceRole, ok := from.GetDatabaseInstanceRole(ctx); ok {
				// Recursively sync the fields of DatabaseInstanceRole
				toDatabaseInstanceRole.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseInstanceRole)
				to.SetDatabaseInstanceRole(ctx, toDatabaseInstanceRole)
			}
		}
	}
}

func (to *CreateDatabaseInstanceRoleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseInstanceRoleRequest_SdkV2) {
	if !from.DatabaseInstanceRole.IsNull() && !from.DatabaseInstanceRole.IsUnknown() {
		if toDatabaseInstanceRole, ok := to.GetDatabaseInstanceRole(ctx); ok {
			if fromDatabaseInstanceRole, ok := from.GetDatabaseInstanceRole(ctx); ok {
				toDatabaseInstanceRole.SyncFieldsDuringRead(ctx, fromDatabaseInstanceRole)
				to.SetDatabaseInstanceRole(ctx, toDatabaseInstanceRole)
			}
		}
	}
}

func (m CreateDatabaseInstanceRoleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_role"] = attrs["database_instance_role"].SetRequired()
	attrs["database_instance_role"] = attrs["database_instance_role"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseInstanceRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseInstanceRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_role": reflect.TypeOf(DatabaseInstanceRole_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDatabaseInstanceRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_name": m.DatabaseInstanceName,
			"database_instance_role": m.DatabaseInstanceRole,
			"instance_name":          m.InstanceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseInstanceRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_name": types.StringType,
			"database_instance_role": basetypes.ListType{
				ElemType: DatabaseInstanceRole_SdkV2{}.Type(ctx),
			},
			"instance_name": types.StringType,
		},
	}
}

// GetDatabaseInstanceRole returns the value of the DatabaseInstanceRole field in CreateDatabaseInstanceRoleRequest_SdkV2 as
// a DatabaseInstanceRole_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseInstanceRoleRequest_SdkV2) GetDatabaseInstanceRole(ctx context.Context) (DatabaseInstanceRole_SdkV2, bool) {
	var e DatabaseInstanceRole_SdkV2
	if m.DatabaseInstanceRole.IsNull() || m.DatabaseInstanceRole.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstanceRole_SdkV2
	d := m.DatabaseInstanceRole.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstanceRole sets the value of the DatabaseInstanceRole field in CreateDatabaseInstanceRoleRequest_SdkV2.
func (m *CreateDatabaseInstanceRoleRequest_SdkV2) SetDatabaseInstanceRole(ctx context.Context, v DatabaseInstanceRole_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance_role"]
	m.DatabaseInstanceRole = types.ListValueMust(t, vs)
}

type CreateDatabaseTableRequest_SdkV2 struct {
	Table types.List `tfsdk:"table"`
}

func (to *CreateDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseTableRequest_SdkV2) {
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				// Recursively sync the fields of Table
				toTable.SyncFieldsDuringCreateOrUpdate(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (to *CreateDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseTableRequest_SdkV2) {
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				toTable.SyncFieldsDuringRead(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (m CreateDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table"] = attrs["table"].SetRequired()
	attrs["table"] = attrs["table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(DatabaseTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": m.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": basetypes.ListType{
				ElemType: DatabaseTable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTable returns the value of the Table field in CreateDatabaseTableRequest_SdkV2 as
// a DatabaseTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseTableRequest_SdkV2) GetTable(ctx context.Context) (DatabaseTable_SdkV2, bool) {
	var e DatabaseTable_SdkV2
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v []DatabaseTable_SdkV2
	d := m.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CreateDatabaseTableRequest_SdkV2.
func (m *CreateDatabaseTableRequest_SdkV2) SetTable(ctx context.Context, v DatabaseTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	m.Table = types.ListValueMust(t, vs)
}

type CreateSyncedDatabaseTableRequest_SdkV2 struct {
	SyncedTable types.List `tfsdk:"synced_table"`
}

func (to *CreateSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSyncedDatabaseTableRequest_SdkV2) {
	if !from.SyncedTable.IsNull() && !from.SyncedTable.IsUnknown() {
		if toSyncedTable, ok := to.GetSyncedTable(ctx); ok {
			if fromSyncedTable, ok := from.GetSyncedTable(ctx); ok {
				// Recursively sync the fields of SyncedTable
				toSyncedTable.SyncFieldsDuringCreateOrUpdate(ctx, fromSyncedTable)
				to.SetSyncedTable(ctx, toSyncedTable)
			}
		}
	}
}

func (to *CreateSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateSyncedDatabaseTableRequest_SdkV2) {
	if !from.SyncedTable.IsNull() && !from.SyncedTable.IsUnknown() {
		if toSyncedTable, ok := to.GetSyncedTable(ctx); ok {
			if fromSyncedTable, ok := from.GetSyncedTable(ctx); ok {
				toSyncedTable.SyncFieldsDuringRead(ctx, fromSyncedTable)
				to.SetSyncedTable(ctx, toSyncedTable)
			}
		}
	}
}

func (m CreateSyncedDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["synced_table"] = attrs["synced_table"].SetRequired()
	attrs["synced_table"] = attrs["synced_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"synced_table": m.SyncedTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"synced_table": basetypes.ListType{
				ElemType: SyncedDatabaseTable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSyncedTable returns the value of the SyncedTable field in CreateSyncedDatabaseTableRequest_SdkV2 as
// a SyncedDatabaseTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSyncedDatabaseTableRequest_SdkV2) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable_SdkV2, bool) {
	var e SyncedDatabaseTable_SdkV2
	if m.SyncedTable.IsNull() || m.SyncedTable.IsUnknown() {
		return e, false
	}
	var v []SyncedDatabaseTable_SdkV2
	d := m.SyncedTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSyncedTable sets the value of the SyncedTable field in CreateSyncedDatabaseTableRequest_SdkV2.
func (m *CreateSyncedDatabaseTableRequest_SdkV2) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synced_table"]
	m.SyncedTable = types.ListValueMust(t, vs)
}

type CustomTag_SdkV2 struct {
	// The key of the custom tag.
	Key types.String `tfsdk:"key"`
	// The value of the custom tag.
	Value types.String `tfsdk:"value"`
}

func (to *CustomTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTag_SdkV2) {
}

func (to *CustomTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomTag_SdkV2) {
}

func (m CustomTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTag_SdkV2
// only implements ToObjectValue() and Type().
func (m CustomTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type DatabaseCatalog_SdkV2 struct {
	CreateDatabaseIfNotExists types.Bool `tfsdk:"create_database_if_not_exists"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName types.String `tfsdk:"database_name"`
	// The name of the catalog in UC.
	Name types.String `tfsdk:"name"`

	Uid types.String `tfsdk:"uid"`
}

func (to *DatabaseCatalog_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCatalog_SdkV2) {
	if !from.CreateDatabaseIfNotExists.IsUnknown() && !from.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseIfNotExists = from.CreateDatabaseIfNotExists
	}
}

func (to *DatabaseCatalog_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseCatalog_SdkV2) {
	if !from.CreateDatabaseIfNotExists.IsUnknown() && !from.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseIfNotExists = from.CreateDatabaseIfNotExists
	}
}

func (m DatabaseCatalog_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetOptional()
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetComputed()
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["database_instance_name"] = attrs["database_instance_name"].SetRequired()
	attrs["database_name"] = attrs["database_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["uid"] = attrs["uid"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseCatalog.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseCatalog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCatalog_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseCatalog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_if_not_exists": m.CreateDatabaseIfNotExists,
			"database_instance_name":        m.DatabaseInstanceName,
			"database_name":                 m.DatabaseName,
			"name":                          m.Name,
			"uid":                           m.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseCatalog_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_if_not_exists": types.BoolType,
			"database_instance_name":        types.StringType,
			"database_name":                 types.StringType,
			"name":                          types.StringType,
			"uid":                           types.StringType,
		},
	}
}

type DatabaseCredential_SdkV2 struct {
	ExpirationTime types.String `tfsdk:"expiration_time"`

	Token types.String `tfsdk:"token"`
}

func (to *DatabaseCredential_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCredential_SdkV2) {
}

func (to *DatabaseCredential_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseCredential_SdkV2) {
}

func (m DatabaseCredential_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["expiration_time"] = attrs["expiration_time"].SetOptional()
	attrs["token"] = attrs["token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseCredential.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCredential_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expiration_time": m.ExpirationTime,
			"token":           m.Token,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expiration_time": types.StringType,
			"token":           types.StringType,
		},
	}
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance_SdkV2 struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	Capacity types.String `tfsdk:"capacity"`
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	ChildInstanceRefs types.List `tfsdk:"child_instance_refs"`
	// The timestamp when the instance was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the instance.
	Creator types.String `tfsdk:"creator"`
	// Custom tags associated with the instance. This field is only included on
	// create and update responses.
	CustomTags types.List `tfsdk:"custom_tags"`
	// Deprecated. The sku of the instance; this field will always match the
	// value of capacity.
	EffectiveCapacity types.String `tfsdk:"effective_capacity"`
	// The recorded custom tags associated with the instance.
	EffectiveCustomTags types.List `tfsdk:"effective_custom_tags"`
	// Whether the instance has PG native password login enabled.
	EffectiveEnablePgNativeLogin types.Bool `tfsdk:"effective_enable_pg_native_login"`
	// Whether secondaries serving read-only traffic are enabled. Defaults to
	// false.
	EffectiveEnableReadableSecondaries types.Bool `tfsdk:"effective_enable_readable_secondaries"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries.
	EffectiveNodeCount types.Int64 `tfsdk:"effective_node_count"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained.
	EffectiveRetentionWindowInDays types.Int64 `tfsdk:"effective_retention_window_in_days"`
	// Whether the instance is stopped.
	EffectiveStopped types.Bool `tfsdk:"effective_stopped"`
	// The policy that is applied to the instance.
	EffectiveUsagePolicyId types.String `tfsdk:"effective_usage_policy_id"`
	// Whether to enable PG native password login on the instance. Defaults to
	// false.
	EnablePgNativeLogin types.Bool `tfsdk:"enable_pg_native_login"`
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false.
	EnableReadableSecondaries types.Bool `tfsdk:"enable_readable_secondaries"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"name"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries. This field is input
	// only, see effective_node_count for the output.
	NodeCount types.Int64 `tfsdk:"node_count"`
	// The ref of the parent instance. This is only available if the instance is
	// child instance. Input: For specifying the parent instance to create a
	// child instance. Optional. Output: Only populated if provided as input to
	// create a child instance.
	ParentInstanceRef types.List `tfsdk:"parent_instance_ref"`
	// The version of Postgres running on the instance.
	PgVersion types.String `tfsdk:"pg_version"`
	// The DNS endpoint to connect to the instance for read only access. This is
	// only available if enable_readable_secondaries is true.
	ReadOnlyDns types.String `tfsdk:"read_only_dns"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns types.String `tfsdk:"read_write_dns"`
	// The retention window for the instance. This is the time window in days
	// for which the historical data is retained. The default value is 7 days.
	// Valid values are 2 to 35 days.
	RetentionWindowInDays types.Int64 `tfsdk:"retention_window_in_days"`
	// The current state of the instance.
	State types.String `tfsdk:"state"`
	// Whether to stop the instance. An input only param, see effective_stopped
	// for the output.
	Stopped types.Bool `tfsdk:"stopped"`
	// An immutable UUID identifier for the instance.
	Uid types.String `tfsdk:"uid"`
	// The desired usage policy to associate with the instance.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *DatabaseInstance_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstance_SdkV2) {
	if !from.ChildInstanceRefs.IsNull() && !from.ChildInstanceRefs.IsUnknown() && to.ChildInstanceRefs.IsNull() && len(from.ChildInstanceRefs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ChildInstanceRefs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ChildInstanceRefs = from.ChildInstanceRefs
	}
	if !from.CustomTags.IsUnknown() && !from.CustomTags.IsNull() {
		// CustomTags is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EffectiveCustomTags.IsNull() && !from.EffectiveCustomTags.IsUnknown() && to.EffectiveCustomTags.IsNull() && len(from.EffectiveCustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveCustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveCustomTags = from.EffectiveCustomTags
	}
	if !from.EnablePgNativeLogin.IsUnknown() && !from.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnablePgNativeLogin = from.EnablePgNativeLogin
	}
	if !from.EnableReadableSecondaries.IsUnknown() && !from.EnableReadableSecondaries.IsNull() {
		// EnableReadableSecondaries is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnableReadableSecondaries = from.EnableReadableSecondaries
	}
	if !from.NodeCount.IsUnknown() && !from.NodeCount.IsNull() {
		// NodeCount is an input only field and not returned by the service, so we keep the value from the prior state.
		to.NodeCount = from.NodeCount
	}
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				// Recursively sync the fields of ParentInstanceRef
				toParentInstanceRef.SyncFieldsDuringCreateOrUpdate(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
	if !from.RetentionWindowInDays.IsUnknown() && !from.RetentionWindowInDays.IsNull() {
		// RetentionWindowInDays is an input only field and not returned by the service, so we keep the value from the prior state.
		to.RetentionWindowInDays = from.RetentionWindowInDays
	}
	if !from.Stopped.IsUnknown() && !from.Stopped.IsNull() {
		// Stopped is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Stopped = from.Stopped
	}
	if !from.UsagePolicyId.IsUnknown() && !from.UsagePolicyId.IsNull() {
		// UsagePolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.UsagePolicyId = from.UsagePolicyId
	}
}

func (to *DatabaseInstance_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstance_SdkV2) {
	if !from.ChildInstanceRefs.IsNull() && !from.ChildInstanceRefs.IsUnknown() && to.ChildInstanceRefs.IsNull() && len(from.ChildInstanceRefs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ChildInstanceRefs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ChildInstanceRefs = from.ChildInstanceRefs
	}
	if !from.CustomTags.IsUnknown() && !from.CustomTags.IsNull() {
		// CustomTags is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CustomTags = from.CustomTags
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.EffectiveCustomTags.IsNull() && !from.EffectiveCustomTags.IsUnknown() && to.EffectiveCustomTags.IsNull() && len(from.EffectiveCustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for EffectiveCustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.EffectiveCustomTags = from.EffectiveCustomTags
	}
	if !from.EnablePgNativeLogin.IsUnknown() && !from.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnablePgNativeLogin = from.EnablePgNativeLogin
	}
	if !from.EnableReadableSecondaries.IsUnknown() && !from.EnableReadableSecondaries.IsNull() {
		// EnableReadableSecondaries is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnableReadableSecondaries = from.EnableReadableSecondaries
	}
	if !from.NodeCount.IsUnknown() && !from.NodeCount.IsNull() {
		// NodeCount is an input only field and not returned by the service, so we keep the value from the prior state.
		to.NodeCount = from.NodeCount
	}
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				toParentInstanceRef.SyncFieldsDuringRead(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
	if !from.RetentionWindowInDays.IsUnknown() && !from.RetentionWindowInDays.IsNull() {
		// RetentionWindowInDays is an input only field and not returned by the service, so we keep the value from the prior state.
		to.RetentionWindowInDays = from.RetentionWindowInDays
	}
	if !from.Stopped.IsUnknown() && !from.Stopped.IsNull() {
		// Stopped is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Stopped = from.Stopped
	}
	if !from.UsagePolicyId.IsUnknown() && !from.UsagePolicyId.IsNull() {
		// UsagePolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.UsagePolicyId = from.UsagePolicyId
	}
}

func (m DatabaseInstance_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["child_instance_refs"] = attrs["child_instance_refs"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["custom_tags"] = attrs["custom_tags"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_capacity"] = attrs["effective_capacity"].SetComputed()
	attrs["effective_custom_tags"] = attrs["effective_custom_tags"].SetComputed()
	attrs["effective_enable_pg_native_login"] = attrs["effective_enable_pg_native_login"].SetComputed()
	attrs["effective_enable_readable_secondaries"] = attrs["effective_enable_readable_secondaries"].SetComputed()
	attrs["effective_node_count"] = attrs["effective_node_count"].SetComputed()
	attrs["effective_retention_window_in_days"] = attrs["effective_retention_window_in_days"].SetComputed()
	attrs["effective_stopped"] = attrs["effective_stopped"].SetComputed()
	attrs["effective_usage_policy_id"] = attrs["effective_usage_policy_id"].SetComputed()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].SetOptional()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].SetComputed()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].SetOptional()
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].SetComputed()
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["node_count"] = attrs["node_count"].SetOptional()
	attrs["node_count"] = attrs["node_count"].SetComputed()
	attrs["node_count"] = attrs["node_count"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].SetOptional()
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["read_only_dns"] = attrs["read_only_dns"].SetComputed()
	attrs["read_write_dns"] = attrs["read_write_dns"].SetComputed()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].SetOptional()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].SetComputed()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["state"] = attrs["state"].SetComputed()
	attrs["stopped"] = attrs["stopped"].SetOptional()
	attrs["stopped"] = attrs["stopped"].SetComputed()
	attrs["stopped"] = attrs["stopped"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetComputed()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseInstance_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_instance_refs":   reflect.TypeOf(DatabaseInstanceRef_SdkV2{}),
		"custom_tags":           reflect.TypeOf(CustomTag_SdkV2{}),
		"effective_custom_tags": reflect.TypeOf(CustomTag_SdkV2{}),
		"parent_instance_ref":   reflect.TypeOf(DatabaseInstanceRef_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseInstance_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"capacity":                              m.Capacity,
			"child_instance_refs":                   m.ChildInstanceRefs,
			"creation_time":                         m.CreationTime,
			"creator":                               m.Creator,
			"custom_tags":                           m.CustomTags,
			"effective_capacity":                    m.EffectiveCapacity,
			"effective_custom_tags":                 m.EffectiveCustomTags,
			"effective_enable_pg_native_login":      m.EffectiveEnablePgNativeLogin,
			"effective_enable_readable_secondaries": m.EffectiveEnableReadableSecondaries,
			"effective_node_count":                  m.EffectiveNodeCount,
			"effective_retention_window_in_days":    m.EffectiveRetentionWindowInDays,
			"effective_stopped":                     m.EffectiveStopped,
			"effective_usage_policy_id":             m.EffectiveUsagePolicyId,
			"enable_pg_native_login":                m.EnablePgNativeLogin,
			"enable_readable_secondaries":           m.EnableReadableSecondaries,
			"name":                                  m.Name,
			"node_count":                            m.NodeCount,
			"parent_instance_ref":                   m.ParentInstanceRef,
			"pg_version":                            m.PgVersion,
			"read_only_dns":                         m.ReadOnlyDns,
			"read_write_dns":                        m.ReadWriteDns,
			"retention_window_in_days":              m.RetentionWindowInDays,
			"state":                                 m.State,
			"stopped":                               m.Stopped,
			"uid":                                   m.Uid,
			"usage_policy_id":                       m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseInstance_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"capacity": types.StringType,
			"child_instance_refs": basetypes.ListType{
				ElemType: DatabaseInstanceRef_SdkV2{}.Type(ctx),
			},
			"creation_time": types.StringType,
			"creator":       types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag_SdkV2{}.Type(ctx),
			},
			"effective_capacity": types.StringType,
			"effective_custom_tags": basetypes.ListType{
				ElemType: CustomTag_SdkV2{}.Type(ctx),
			},
			"effective_enable_pg_native_login":      types.BoolType,
			"effective_enable_readable_secondaries": types.BoolType,
			"effective_node_count":                  types.Int64Type,
			"effective_retention_window_in_days":    types.Int64Type,
			"effective_stopped":                     types.BoolType,
			"effective_usage_policy_id":             types.StringType,
			"enable_pg_native_login":                types.BoolType,
			"enable_readable_secondaries":           types.BoolType,
			"name":                                  types.StringType,
			"node_count":                            types.Int64Type,
			"parent_instance_ref": basetypes.ListType{
				ElemType: DatabaseInstanceRef_SdkV2{}.Type(ctx),
			},
			"pg_version":               types.StringType,
			"read_only_dns":            types.StringType,
			"read_write_dns":           types.StringType,
			"retention_window_in_days": types.Int64Type,
			"state":                    types.StringType,
			"stopped":                  types.BoolType,
			"uid":                      types.StringType,
			"usage_policy_id":          types.StringType,
		},
	}
}

// GetChildInstanceRefs returns the value of the ChildInstanceRefs field in DatabaseInstance_SdkV2 as
// a slice of DatabaseInstanceRef_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance_SdkV2) GetChildInstanceRefs(ctx context.Context) ([]DatabaseInstanceRef_SdkV2, bool) {
	if m.ChildInstanceRefs.IsNull() || m.ChildInstanceRefs.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstanceRef_SdkV2
	d := m.ChildInstanceRefs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChildInstanceRefs sets the value of the ChildInstanceRefs field in DatabaseInstance_SdkV2.
func (m *DatabaseInstance_SdkV2) SetChildInstanceRefs(ctx context.Context, v []DatabaseInstanceRef_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["child_instance_refs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ChildInstanceRefs = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in DatabaseInstance_SdkV2 as
// a slice of CustomTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance_SdkV2) GetCustomTags(ctx context.Context) ([]CustomTag_SdkV2, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag_SdkV2
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in DatabaseInstance_SdkV2.
func (m *DatabaseInstance_SdkV2) SetCustomTags(ctx context.Context, v []CustomTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEffectiveCustomTags returns the value of the EffectiveCustomTags field in DatabaseInstance_SdkV2 as
// a slice of CustomTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance_SdkV2) GetEffectiveCustomTags(ctx context.Context) ([]CustomTag_SdkV2, bool) {
	if m.EffectiveCustomTags.IsNull() || m.EffectiveCustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag_SdkV2
	d := m.EffectiveCustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveCustomTags sets the value of the EffectiveCustomTags field in DatabaseInstance_SdkV2.
func (m *DatabaseInstance_SdkV2) SetEffectiveCustomTags(ctx context.Context, v []CustomTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveCustomTags = types.ListValueMust(t, vs)
}

// GetParentInstanceRef returns the value of the ParentInstanceRef field in DatabaseInstance_SdkV2 as
// a DatabaseInstanceRef_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance_SdkV2) GetParentInstanceRef(ctx context.Context) (DatabaseInstanceRef_SdkV2, bool) {
	var e DatabaseInstanceRef_SdkV2
	if m.ParentInstanceRef.IsNull() || m.ParentInstanceRef.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstanceRef_SdkV2
	d := m.ParentInstanceRef.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetParentInstanceRef sets the value of the ParentInstanceRef field in DatabaseInstance_SdkV2.
func (m *DatabaseInstance_SdkV2) SetParentInstanceRef(ctx context.Context, v DatabaseInstanceRef_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["parent_instance_ref"]
	m.ParentInstanceRef = types.ListValueMust(t, vs)
}

// DatabaseInstanceRef is a reference to a database instance. It is used in the
// DatabaseInstance object to refer to the parent instance of an instance and to
// refer the child instances of an instance. To specify as a parent instance
// during creation of an instance, the lsn and branch_time fields are optional.
// If not specified, the child instance will be created from the latest lsn of
// the parent. If both lsn and branch_time are specified, the lsn will be used
// to create the child instance.
type DatabaseInstanceRef_SdkV2 struct {
	// Branch time of the ref database instance. For a parent ref instance, this
	// is the point in time on the parent instance from which the instance was
	// created. For a child ref instance, this is the point in time on the
	// instance from which the child instance was created. Input: For specifying
	// the point in time to create a child instance. Optional. Output: Only
	// populated if provided as input to create a child instance.
	BranchTime types.String `tfsdk:"branch_time"`
	// For a parent ref instance, this is the LSN on the parent instance from
	// which the instance was created. For a child ref instance, this is the LSN
	// on the instance from which the child instance was created.
	EffectiveLsn types.String `tfsdk:"effective_lsn"`
	// User-specified WAL LSN of the ref database instance.
	//
	// Input: For specifying the WAL LSN to create a child instance. Optional.
	// Output: Only populated if provided as input to create a child instance.
	Lsn types.String `tfsdk:"lsn"`
	// Name of the ref database instance.
	Name types.String `tfsdk:"name"`
	// Id of the ref database instance.
	Uid types.String `tfsdk:"uid"`
}

func (to *DatabaseInstanceRef_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstanceRef_SdkV2) {
	if !from.Lsn.IsUnknown() && !from.Lsn.IsNull() {
		// Lsn is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Lsn = from.Lsn
	}
}

func (to *DatabaseInstanceRef_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstanceRef_SdkV2) {
	if !from.Lsn.IsUnknown() && !from.Lsn.IsNull() {
		// Lsn is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Lsn = from.Lsn
	}
}

func (m DatabaseInstanceRef_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_time"] = attrs["branch_time"].SetOptional()
	attrs["effective_lsn"] = attrs["effective_lsn"].SetComputed()
	attrs["lsn"] = attrs["lsn"].SetOptional()
	attrs["lsn"] = attrs["lsn"].SetComputed()
	attrs["lsn"] = attrs["lsn"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["uid"] = attrs["uid"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstanceRef.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseInstanceRef_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRef_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseInstanceRef_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_time":   m.BranchTime,
			"effective_lsn": m.EffectiveLsn,
			"lsn":           m.Lsn,
			"name":          m.Name,
			"uid":           m.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseInstanceRef_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_time":   types.StringType,
			"effective_lsn": types.StringType,
			"lsn":           types.StringType,
			"name":          types.StringType,
			"uid":           types.StringType,
		},
	}
}

// A DatabaseInstanceRole represents a Postgres role in a database instance.
type DatabaseInstanceRole_SdkV2 struct {
	// The desired API-exposed Postgres role attribute to associate with the
	// role. Optional.
	Attributes types.List `tfsdk:"attributes"`
	// The attributes that are applied to the role.
	EffectiveAttributes types.List `tfsdk:"effective_attributes"`
	// The type of the role.
	IdentityType types.String `tfsdk:"identity_type"`

	InstanceName types.String `tfsdk:"instance_name"`
	// An enum value for a standard role that this role is a member of.
	MembershipRole types.String `tfsdk:"membership_role"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	Name types.String `tfsdk:"name"`
}

func (to *DatabaseInstanceRole_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstanceRole_SdkV2) {
	if !from.Attributes.IsNull() && !from.Attributes.IsUnknown() {
		if toAttributes, ok := to.GetAttributes(ctx); ok {
			if fromAttributes, ok := from.GetAttributes(ctx); ok {
				// Recursively sync the fields of Attributes
				toAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAttributes)
				to.SetAttributes(ctx, toAttributes)
			}
		}
	}
	if !from.EffectiveAttributes.IsNull() && !from.EffectiveAttributes.IsUnknown() {
		if toEffectiveAttributes, ok := to.GetEffectiveAttributes(ctx); ok {
			if fromEffectiveAttributes, ok := from.GetEffectiveAttributes(ctx); ok {
				// Recursively sync the fields of EffectiveAttributes
				toEffectiveAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromEffectiveAttributes)
				to.SetEffectiveAttributes(ctx, toEffectiveAttributes)
			}
		}
	}
}

func (to *DatabaseInstanceRole_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstanceRole_SdkV2) {
	if !from.Attributes.IsNull() && !from.Attributes.IsUnknown() {
		if toAttributes, ok := to.GetAttributes(ctx); ok {
			if fromAttributes, ok := from.GetAttributes(ctx); ok {
				toAttributes.SyncFieldsDuringRead(ctx, fromAttributes)
				to.SetAttributes(ctx, toAttributes)
			}
		}
	}
	if !from.EffectiveAttributes.IsNull() && !from.EffectiveAttributes.IsUnknown() {
		if toEffectiveAttributes, ok := to.GetEffectiveAttributes(ctx); ok {
			if fromEffectiveAttributes, ok := from.GetEffectiveAttributes(ctx); ok {
				toEffectiveAttributes.SyncFieldsDuringRead(ctx, fromEffectiveAttributes)
				to.SetEffectiveAttributes(ctx, toEffectiveAttributes)
			}
		}
	}
}

func (m DatabaseInstanceRole_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attributes"] = attrs["attributes"].SetOptional()
	attrs["attributes"] = attrs["attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["effective_attributes"] = attrs["effective_attributes"].SetComputed()
	attrs["effective_attributes"] = attrs["effective_attributes"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["identity_type"] = attrs["identity_type"].SetOptional()
	attrs["instance_name"] = attrs["instance_name"].SetOptional()
	attrs["membership_role"] = attrs["membership_role"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstanceRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseInstanceRole_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes":           reflect.TypeOf(DatabaseInstanceRoleAttributes_SdkV2{}),
		"effective_attributes": reflect.TypeOf(DatabaseInstanceRoleAttributes_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRole_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseInstanceRole_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":           m.Attributes,
			"effective_attributes": m.EffectiveAttributes,
			"identity_type":        m.IdentityType,
			"instance_name":        m.InstanceName,
			"membership_role":      m.MembershipRole,
			"name":                 m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseInstanceRole_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes": basetypes.ListType{
				ElemType: DatabaseInstanceRoleAttributes_SdkV2{}.Type(ctx),
			},
			"effective_attributes": basetypes.ListType{
				ElemType: DatabaseInstanceRoleAttributes_SdkV2{}.Type(ctx),
			},
			"identity_type":   types.StringType,
			"instance_name":   types.StringType,
			"membership_role": types.StringType,
			"name":            types.StringType,
		},
	}
}

// GetAttributes returns the value of the Attributes field in DatabaseInstanceRole_SdkV2 as
// a DatabaseInstanceRoleAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstanceRole_SdkV2) GetAttributes(ctx context.Context) (DatabaseInstanceRoleAttributes_SdkV2, bool) {
	var e DatabaseInstanceRoleAttributes_SdkV2
	if m.Attributes.IsNull() || m.Attributes.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstanceRoleAttributes_SdkV2
	d := m.Attributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAttributes sets the value of the Attributes field in DatabaseInstanceRole_SdkV2.
func (m *DatabaseInstanceRole_SdkV2) SetAttributes(ctx context.Context, v DatabaseInstanceRoleAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["attributes"]
	m.Attributes = types.ListValueMust(t, vs)
}

// GetEffectiveAttributes returns the value of the EffectiveAttributes field in DatabaseInstanceRole_SdkV2 as
// a DatabaseInstanceRoleAttributes_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstanceRole_SdkV2) GetEffectiveAttributes(ctx context.Context) (DatabaseInstanceRoleAttributes_SdkV2, bool) {
	var e DatabaseInstanceRoleAttributes_SdkV2
	if m.EffectiveAttributes.IsNull() || m.EffectiveAttributes.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstanceRoleAttributes_SdkV2
	d := m.EffectiveAttributes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEffectiveAttributes sets the value of the EffectiveAttributes field in DatabaseInstanceRole_SdkV2.
func (m *DatabaseInstanceRole_SdkV2) SetEffectiveAttributes(ctx context.Context, v DatabaseInstanceRoleAttributes_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_attributes"]
	m.EffectiveAttributes = types.ListValueMust(t, vs)
}

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words. We were requested to make
// this a nested object/struct representation since these are knobs from an
// external spec.
type DatabaseInstanceRoleAttributes_SdkV2 struct {
	Bypassrls types.Bool `tfsdk:"bypassrls"`

	Createdb types.Bool `tfsdk:"createdb"`

	Createrole types.Bool `tfsdk:"createrole"`
}

func (to *DatabaseInstanceRoleAttributes_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstanceRoleAttributes_SdkV2) {
}

func (to *DatabaseInstanceRoleAttributes_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstanceRoleAttributes_SdkV2) {
}

func (m DatabaseInstanceRoleAttributes_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["bypassrls"] = attrs["bypassrls"].SetOptional()
	attrs["createdb"] = attrs["createdb"].SetOptional()
	attrs["createrole"] = attrs["createrole"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstanceRoleAttributes.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseInstanceRoleAttributes_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRoleAttributes_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseInstanceRoleAttributes_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bypassrls":  m.Bypassrls,
			"createdb":   m.Createdb,
			"createrole": m.Createrole,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseInstanceRoleAttributes_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bypassrls":  types.BoolType,
			"createdb":   types.BoolType,
			"createrole": types.BoolType,
		},
	}
}

// Next field marker: 13
type DatabaseTable_SdkV2 struct {
	// Name of the target database instance. This is required when creating
	// database tables in standard catalogs. This is optional when creating
	// database tables in registered catalogs. If this field is specified when
	// creating database tables in registered catalogs, the database instance
	// name MUST match that of the registered catalog (or the request will be
	// rejected).
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// Target Postgres database object (logical database) name for this table.
	//
	// When creating a table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a table in a standard catalog, this field is required. In
	// this scenario, specifying this field will allow targeting an arbitrary
	// postgres database.
	LogicalDatabaseName types.String `tfsdk:"logical_database_name"`
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name"`
}

func (to *DatabaseTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseTable_SdkV2) {
}

func (to *DatabaseTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabaseTable_SdkV2) {
}

func (m DatabaseTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseTable_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabaseTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_name": m.DatabaseInstanceName,
			"logical_database_name":  m.LogicalDatabaseName,
			"name":                   m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_name": types.StringType,
			"logical_database_name":  types.StringType,
			"name":                   types.StringType,
		},
	}
}

type DeleteDatabaseCatalogRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseCatalogRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseCatalogRequest_SdkV2) {
}

func (to *DeleteDatabaseCatalogRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseCatalogRequest_SdkV2) {
}

func (m DeleteDatabaseCatalogRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseInstanceRequest_SdkV2 struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force types.Bool `tfsdk:"-"`
	// Name of the instance to delete.
	Name types.String `tfsdk:"-"`
	// Deprecated. Omitting the field or setting it to true will result in the
	// field being hard deleted. Setting a value of false will throw a bad
	// request.
	Purge types.Bool `tfsdk:"-"`
}

func (to *DeleteDatabaseInstanceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseInstanceRequest_SdkV2) {
}

func (to *DeleteDatabaseInstanceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseInstanceRequest_SdkV2) {
}

func (m DeleteDatabaseInstanceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["force"] = attrs["force"].SetOptional()
	attrs["purge"] = attrs["purge"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": m.Force,
			"name":  m.Name,
			"purge": m.Purge,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
			"purge": types.BoolType,
		},
	}
}

type DeleteDatabaseInstanceRoleRequest_SdkV2 struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing types.Bool `tfsdk:"-"`

	InstanceName types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`

	ReassignOwnedTo types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseInstanceRoleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseInstanceRoleRequest_SdkV2) {
}

func (to *DeleteDatabaseInstanceRoleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseInstanceRoleRequest_SdkV2) {
}

func (m DeleteDatabaseInstanceRoleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["reassign_owned_to"] = attrs["reassign_owned_to"].SetOptional()
	attrs["allow_missing"] = attrs["allow_missing"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseInstanceRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseInstanceRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing":     m.AllowMissing,
			"instance_name":     m.InstanceName,
			"name":              m.Name,
			"reassign_owned_to": m.ReassignOwnedTo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseInstanceRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing":     types.BoolType,
			"instance_name":     types.StringType,
			"name":              types.StringType,
			"reassign_owned_to": types.StringType,
		},
	}
}

type DeleteDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseTableRequest_SdkV2) {
}

func (to *DeleteDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseTableRequest_SdkV2) {
}

func (m DeleteDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSyncedDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSyncedDatabaseTableRequest_SdkV2) {
}

func (to *DeleteSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSyncedDatabaseTableRequest_SdkV2) {
}

func (m DeleteSyncedDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeltaTableSyncInfo_SdkV2 struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	DeltaCommitTimestamp types.String `tfsdk:"delta_commit_timestamp"`
	// The Delta Lake commit version that was last successfully synced.
	DeltaCommitVersion types.Int64 `tfsdk:"delta_commit_version"`
}

func (to *DeltaTableSyncInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaTableSyncInfo_SdkV2) {
}

func (to *DeltaTableSyncInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeltaTableSyncInfo_SdkV2) {
}

func (m DeltaTableSyncInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_commit_timestamp"] = attrs["delta_commit_timestamp"].SetComputed()
	attrs["delta_commit_version"] = attrs["delta_commit_version"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeltaTableSyncInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeltaTableSyncInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaTableSyncInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m DeltaTableSyncInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_commit_timestamp": m.DeltaCommitTimestamp,
			"delta_commit_version":   m.DeltaCommitVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaTableSyncInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_commit_timestamp": types.StringType,
			"delta_commit_version":   types.Int64Type,
		},
	}
}

type FindDatabaseInstanceByUidRequest_SdkV2 struct {
	// UID of the cluster to get.
	Uid types.String `tfsdk:"-"`
}

func (to *FindDatabaseInstanceByUidRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FindDatabaseInstanceByUidRequest_SdkV2) {
}

func (to *FindDatabaseInstanceByUidRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FindDatabaseInstanceByUidRequest_SdkV2) {
}

func (m FindDatabaseInstanceByUidRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["uid"] = attrs["uid"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FindDatabaseInstanceByUidRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FindDatabaseInstanceByUidRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FindDatabaseInstanceByUidRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m FindDatabaseInstanceByUidRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"uid": m.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FindDatabaseInstanceByUidRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uid": types.StringType,
		},
	}
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest_SdkV2 struct {
	// The returned token will be scoped to the union of instance_names and
	// instances containing the specified UC tables, so instance_names is
	// allowed to be empty.
	Claims types.List `tfsdk:"claims"`
	// Instances to which the token will be scoped.
	InstanceNames types.List `tfsdk:"instance_names"`

	RequestId types.String `tfsdk:"request_id"`
}

func (to *GenerateDatabaseCredentialRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenerateDatabaseCredentialRequest_SdkV2) {
	if !from.Claims.IsNull() && !from.Claims.IsUnknown() && to.Claims.IsNull() && len(from.Claims.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Claims, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Claims = from.Claims
	}
	if !from.InstanceNames.IsNull() && !from.InstanceNames.IsUnknown() && to.InstanceNames.IsNull() && len(from.InstanceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InstanceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InstanceNames = from.InstanceNames
	}
}

func (to *GenerateDatabaseCredentialRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GenerateDatabaseCredentialRequest_SdkV2) {
	if !from.Claims.IsNull() && !from.Claims.IsUnknown() && to.Claims.IsNull() && len(from.Claims.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Claims, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Claims = from.Claims
	}
	if !from.InstanceNames.IsNull() && !from.InstanceNames.IsUnknown() && to.InstanceNames.IsNull() && len(from.InstanceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InstanceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InstanceNames = from.InstanceNames
	}
}

func (m GenerateDatabaseCredentialRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["claims"] = attrs["claims"].SetOptional()
	attrs["instance_names"] = attrs["instance_names"].SetOptional()
	attrs["request_id"] = attrs["request_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenerateDatabaseCredentialRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenerateDatabaseCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"claims":         reflect.TypeOf(RequestedClaims_SdkV2{}),
		"instance_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateDatabaseCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GenerateDatabaseCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"claims":         m.Claims,
			"instance_names": m.InstanceNames,
			"request_id":     m.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenerateDatabaseCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"claims": basetypes.ListType{
				ElemType: RequestedClaims_SdkV2{}.Type(ctx),
			},
			"instance_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"request_id": types.StringType,
		},
	}
}

// GetClaims returns the value of the Claims field in GenerateDatabaseCredentialRequest_SdkV2 as
// a slice of RequestedClaims_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenerateDatabaseCredentialRequest_SdkV2) GetClaims(ctx context.Context) ([]RequestedClaims_SdkV2, bool) {
	if m.Claims.IsNull() || m.Claims.IsUnknown() {
		return nil, false
	}
	var v []RequestedClaims_SdkV2
	d := m.Claims.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClaims sets the value of the Claims field in GenerateDatabaseCredentialRequest_SdkV2.
func (m *GenerateDatabaseCredentialRequest_SdkV2) SetClaims(ctx context.Context, v []RequestedClaims_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["claims"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Claims = types.ListValueMust(t, vs)
}

// GetInstanceNames returns the value of the InstanceNames field in GenerateDatabaseCredentialRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenerateDatabaseCredentialRequest_SdkV2) GetInstanceNames(ctx context.Context) ([]types.String, bool) {
	if m.InstanceNames.IsNull() || m.InstanceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InstanceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstanceNames sets the value of the InstanceNames field in GenerateDatabaseCredentialRequest_SdkV2.
func (m *GenerateDatabaseCredentialRequest_SdkV2) SetInstanceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InstanceNames = types.ListValueMust(t, vs)
}

type GetDatabaseCatalogRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseCatalogRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseCatalogRequest_SdkV2) {
}

func (to *GetDatabaseCatalogRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseCatalogRequest_SdkV2) {
}

func (m GetDatabaseCatalogRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatabaseInstanceRequest_SdkV2 struct {
	// Name of the cluster to get.
	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseInstanceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseInstanceRequest_SdkV2) {
}

func (to *GetDatabaseInstanceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseInstanceRequest_SdkV2) {
}

func (m GetDatabaseInstanceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatabaseInstanceRoleRequest_SdkV2 struct {
	InstanceName types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseInstanceRoleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseInstanceRoleRequest_SdkV2) {
}

func (to *GetDatabaseInstanceRoleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseInstanceRoleRequest_SdkV2) {
}

func (m GetDatabaseInstanceRoleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseInstanceRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseInstanceRoleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRoleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDatabaseInstanceRoleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"name":          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseInstanceRoleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"name":          types.StringType,
		},
	}
}

type GetDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseTableRequest_SdkV2) {
}

func (to *GetDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseTableRequest_SdkV2) {
}

func (m GetDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSyncedDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSyncedDatabaseTableRequest_SdkV2) {
}

func (to *GetSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSyncedDatabaseTableRequest_SdkV2) {
}

func (m GetSyncedDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListDatabaseCatalogsRequest_SdkV2 struct {
	// Name of the instance to get database catalogs for.
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseCatalogsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseCatalogsRequest_SdkV2) {
}

func (to *ListDatabaseCatalogsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseCatalogsRequest_SdkV2) {
}

func (m ListDatabaseCatalogsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseCatalogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseCatalogsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseCatalogsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatabaseCatalogsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"page_size":     m.PageSize,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseCatalogsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"page_size":     types.Int64Type,
			"page_token":    types.StringType,
		},
	}
}

type ListDatabaseCatalogsResponse_SdkV2 struct {
	DatabaseCatalogs types.List `tfsdk:"database_catalogs"`
	// Pagination token to request the next page of database catalogs.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseCatalogsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseCatalogsResponse_SdkV2) {
	if !from.DatabaseCatalogs.IsNull() && !from.DatabaseCatalogs.IsUnknown() && to.DatabaseCatalogs.IsNull() && len(from.DatabaseCatalogs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseCatalogs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseCatalogs = from.DatabaseCatalogs
	}
}

func (to *ListDatabaseCatalogsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseCatalogsResponse_SdkV2) {
	if !from.DatabaseCatalogs.IsNull() && !from.DatabaseCatalogs.IsUnknown() && to.DatabaseCatalogs.IsNull() && len(from.DatabaseCatalogs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseCatalogs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseCatalogs = from.DatabaseCatalogs
	}
}

func (m ListDatabaseCatalogsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_catalogs"] = attrs["database_catalogs"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseCatalogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseCatalogsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalogs": reflect.TypeOf(DatabaseCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseCatalogsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatabaseCatalogsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_catalogs": m.DatabaseCatalogs,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseCatalogsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_catalogs": basetypes.ListType{
				ElemType: DatabaseCatalog_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseCatalogs returns the value of the DatabaseCatalogs field in ListDatabaseCatalogsResponse_SdkV2 as
// a slice of DatabaseCatalog_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatabaseCatalogsResponse_SdkV2) GetDatabaseCatalogs(ctx context.Context) ([]DatabaseCatalog_SdkV2, bool) {
	if m.DatabaseCatalogs.IsNull() || m.DatabaseCatalogs.IsUnknown() {
		return nil, false
	}
	var v []DatabaseCatalog_SdkV2
	d := m.DatabaseCatalogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseCatalogs sets the value of the DatabaseCatalogs field in ListDatabaseCatalogsResponse_SdkV2.
func (m *ListDatabaseCatalogsResponse_SdkV2) SetDatabaseCatalogs(ctx context.Context, v []DatabaseCatalog_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_catalogs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseCatalogs = types.ListValueMust(t, vs)
}

type ListDatabaseInstanceRolesRequest_SdkV2 struct {
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseInstanceRolesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstanceRolesRequest_SdkV2) {
}

func (to *ListDatabaseInstanceRolesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstanceRolesRequest_SdkV2) {
}

func (m ListDatabaseInstanceRolesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstanceRolesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseInstanceRolesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstanceRolesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstanceRolesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"page_size":     m.PageSize,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstanceRolesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"page_size":     types.Int64Type,
			"page_token":    types.StringType,
		},
	}
}

type ListDatabaseInstanceRolesResponse_SdkV2 struct {
	// List of database instance roles.
	DatabaseInstanceRoles types.List `tfsdk:"database_instance_roles"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseInstanceRolesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstanceRolesResponse_SdkV2) {
	if !from.DatabaseInstanceRoles.IsNull() && !from.DatabaseInstanceRoles.IsUnknown() && to.DatabaseInstanceRoles.IsNull() && len(from.DatabaseInstanceRoles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstanceRoles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstanceRoles = from.DatabaseInstanceRoles
	}
}

func (to *ListDatabaseInstanceRolesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstanceRolesResponse_SdkV2) {
	if !from.DatabaseInstanceRoles.IsNull() && !from.DatabaseInstanceRoles.IsUnknown() && to.DatabaseInstanceRoles.IsNull() && len(from.DatabaseInstanceRoles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstanceRoles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstanceRoles = from.DatabaseInstanceRoles
	}
}

func (m ListDatabaseInstanceRolesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_roles"] = attrs["database_instance_roles"].SetComputed()
	attrs["next_page_token"] = attrs["next_page_token"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstanceRolesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseInstanceRolesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_roles": reflect.TypeOf(DatabaseInstanceRole_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstanceRolesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstanceRolesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_roles": m.DatabaseInstanceRoles,
			"next_page_token":         m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstanceRolesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_roles": basetypes.ListType{
				ElemType: DatabaseInstanceRole_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseInstanceRoles returns the value of the DatabaseInstanceRoles field in ListDatabaseInstanceRolesResponse_SdkV2 as
// a slice of DatabaseInstanceRole_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatabaseInstanceRolesResponse_SdkV2) GetDatabaseInstanceRoles(ctx context.Context) ([]DatabaseInstanceRole_SdkV2, bool) {
	if m.DatabaseInstanceRoles.IsNull() || m.DatabaseInstanceRoles.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstanceRole_SdkV2
	d := m.DatabaseInstanceRoles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstanceRoles sets the value of the DatabaseInstanceRoles field in ListDatabaseInstanceRolesResponse_SdkV2.
func (m *ListDatabaseInstanceRolesResponse_SdkV2) SetDatabaseInstanceRoles(ctx context.Context, v []DatabaseInstanceRole_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance_roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseInstanceRoles = types.ListValueMust(t, vs)
}

type ListDatabaseInstancesRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseInstancesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstancesRequest_SdkV2) {
}

func (to *ListDatabaseInstancesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstancesRequest_SdkV2) {
}

func (m ListDatabaseInstancesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstancesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseInstancesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstancesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstancesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDatabaseInstancesResponse_SdkV2 struct {
	// List of instances.
	DatabaseInstances types.List `tfsdk:"database_instances"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseInstancesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstancesResponse_SdkV2) {
	if !from.DatabaseInstances.IsNull() && !from.DatabaseInstances.IsUnknown() && to.DatabaseInstances.IsNull() && len(from.DatabaseInstances.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstances, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstances = from.DatabaseInstances
	}
}

func (to *ListDatabaseInstancesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstancesResponse_SdkV2) {
	if !from.DatabaseInstances.IsNull() && !from.DatabaseInstances.IsUnknown() && to.DatabaseInstances.IsNull() && len(from.DatabaseInstances.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstances, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstances = from.DatabaseInstances
	}
}

func (m ListDatabaseInstancesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instances"] = attrs["database_instances"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstancesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseInstancesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(DatabaseInstance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstancesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instances": m.DatabaseInstances,
			"next_page_token":    m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstancesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instances": basetypes.ListType{
				ElemType: DatabaseInstance_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseInstances returns the value of the DatabaseInstances field in ListDatabaseInstancesResponse_SdkV2 as
// a slice of DatabaseInstance_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatabaseInstancesResponse_SdkV2) GetDatabaseInstances(ctx context.Context) ([]DatabaseInstance_SdkV2, bool) {
	if m.DatabaseInstances.IsNull() || m.DatabaseInstances.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstance_SdkV2
	d := m.DatabaseInstances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstances sets the value of the DatabaseInstances field in ListDatabaseInstancesResponse_SdkV2.
func (m *ListDatabaseInstancesResponse_SdkV2) SetDatabaseInstances(ctx context.Context, v []DatabaseInstance_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseInstances = types.ListValueMust(t, vs)
}

type ListSyncedDatabaseTablesRequest_SdkV2 struct {
	// Name of the instance to get synced tables for.
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSyncedDatabaseTablesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSyncedDatabaseTablesRequest_SdkV2) {
}

func (to *ListSyncedDatabaseTablesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSyncedDatabaseTablesRequest_SdkV2) {
}

func (m ListSyncedDatabaseTablesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSyncedDatabaseTablesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSyncedDatabaseTablesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSyncedDatabaseTablesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSyncedDatabaseTablesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"page_size":     m.PageSize,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSyncedDatabaseTablesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"page_size":     types.Int64Type,
			"page_token":    types.StringType,
		},
	}
}

type ListSyncedDatabaseTablesResponse_SdkV2 struct {
	// Pagination token to request the next page of synced tables.
	NextPageToken types.String `tfsdk:"next_page_token"`

	SyncedTables types.List `tfsdk:"synced_tables"`
}

func (to *ListSyncedDatabaseTablesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSyncedDatabaseTablesResponse_SdkV2) {
	if !from.SyncedTables.IsNull() && !from.SyncedTables.IsUnknown() && to.SyncedTables.IsNull() && len(from.SyncedTables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SyncedTables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SyncedTables = from.SyncedTables
	}
}

func (to *ListSyncedDatabaseTablesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListSyncedDatabaseTablesResponse_SdkV2) {
	if !from.SyncedTables.IsNull() && !from.SyncedTables.IsUnknown() && to.SyncedTables.IsNull() && len(from.SyncedTables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SyncedTables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SyncedTables = from.SyncedTables
	}
}

func (m ListSyncedDatabaseTablesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["synced_tables"] = attrs["synced_tables"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSyncedDatabaseTablesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListSyncedDatabaseTablesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_tables": reflect.TypeOf(SyncedDatabaseTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSyncedDatabaseTablesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListSyncedDatabaseTablesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"synced_tables":   m.SyncedTables,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSyncedDatabaseTablesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"synced_tables": basetypes.ListType{
				ElemType: SyncedDatabaseTable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSyncedTables returns the value of the SyncedTables field in ListSyncedDatabaseTablesResponse_SdkV2 as
// a slice of SyncedDatabaseTable_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListSyncedDatabaseTablesResponse_SdkV2) GetSyncedTables(ctx context.Context) ([]SyncedDatabaseTable_SdkV2, bool) {
	if m.SyncedTables.IsNull() || m.SyncedTables.IsUnknown() {
		return nil, false
	}
	var v []SyncedDatabaseTable_SdkV2
	d := m.SyncedTables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSyncedTables sets the value of the SyncedTables field in ListSyncedDatabaseTablesResponse_SdkV2.
func (m *ListSyncedDatabaseTablesResponse_SdkV2) SetSyncedTables(ctx context.Context, v []SyncedDatabaseTable_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synced_tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.SyncedTables = types.ListValueMust(t, vs)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec_SdkV2 struct {
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog types.String `tfsdk:"storage_catalog"`
	// This field needs to be specified if the destination catalog is a managed
	// postgres catalog.
	//
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	StorageSchema types.String `tfsdk:"storage_schema"`
}

func (to *NewPipelineSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NewPipelineSpec_SdkV2) {
}

func (to *NewPipelineSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NewPipelineSpec_SdkV2) {
}

func (m NewPipelineSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["storage_catalog"] = attrs["storage_catalog"].SetOptional()
	attrs["storage_schema"] = attrs["storage_schema"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NewPipelineSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NewPipelineSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NewPipelineSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m NewPipelineSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_catalog": m.StorageCatalog,
			"storage_schema":  m.StorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NewPipelineSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_catalog": types.StringType,
			"storage_schema":  types.StringType,
		},
	}
}

type RequestedClaims_SdkV2 struct {
	PermissionSet types.String `tfsdk:"permission_set"`

	Resources types.List `tfsdk:"resources"`
}

func (to *RequestedClaims_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedClaims_SdkV2) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (to *RequestedClaims_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RequestedClaims_SdkV2) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (m RequestedClaims_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_set"] = attrs["permission_set"].SetOptional()
	attrs["resources"] = attrs["resources"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RequestedClaims.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RequestedClaims_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(RequestedResource_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedClaims_SdkV2
// only implements ToObjectValue() and Type().
func (m RequestedClaims_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": m.PermissionSet,
			"resources":      m.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedClaims_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_set": types.StringType,
			"resources": basetypes.ListType{
				ElemType: RequestedResource_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResources returns the value of the Resources field in RequestedClaims_SdkV2 as
// a slice of RequestedResource_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *RequestedClaims_SdkV2) GetResources(ctx context.Context) ([]RequestedResource_SdkV2, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []RequestedResource_SdkV2
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in RequestedClaims_SdkV2.
func (m *RequestedClaims_SdkV2) SetResources(ctx context.Context, v []RequestedResource_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

type RequestedResource_SdkV2 struct {
	TableName types.String `tfsdk:"table_name"`

	UnspecifiedResourceName types.String `tfsdk:"unspecified_resource_name"`
}

func (to *RequestedResource_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedResource_SdkV2) {
}

func (to *RequestedResource_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RequestedResource_SdkV2) {
}

func (m RequestedResource_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table_name"] = attrs["table_name"].SetOptional()
	attrs["unspecified_resource_name"] = attrs["unspecified_resource_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RequestedResource.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RequestedResource_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedResource_SdkV2
// only implements ToObjectValue() and Type().
func (m RequestedResource_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":                m.TableName,
			"unspecified_resource_name": m.UnspecifiedResourceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedResource_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":                types.StringType,
			"unspecified_resource_name": types.StringType,
		},
	}
}

// Next field marker: 18
type SyncedDatabaseTable_SdkV2 struct {
	// Synced Table data synchronization status
	DataSynchronizationStatus types.List `tfsdk:"data_synchronization_status"`
	// Name of the target database instance. This is required when creating
	// synced database tables in standard catalogs. This is optional when
	// creating synced database tables in registered catalogs. If this field is
	// specified when creating synced database tables in registered catalogs,
	// the database instance name MUST match that of the registered catalog (or
	// the request will be rejected).
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// The name of the database instance that this table is registered to. This
	// field is always returned, and for tables inside database catalogs is
	// inferred database instance associated with the catalog.
	EffectiveDatabaseInstanceName types.String `tfsdk:"effective_database_instance_name"`
	// The name of the logical database that this table is registered to.
	EffectiveLogicalDatabaseName types.String `tfsdk:"effective_logical_database_name"`
	// Target Postgres database object (logical database) name for this table.
	//
	// When creating a synced table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a synced table in a standard catalog, this field is
	// required. In this scenario, specifying this field will allow targeting an
	// arbitrary postgres database. Note that this has implications for the
	// `create_database_objects_is_missing` field in `spec`.
	LogicalDatabaseName types.String `tfsdk:"logical_database_name"`
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name"`

	Spec types.List `tfsdk:"spec"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state"`
}

func (to *SyncedDatabaseTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedDatabaseTable_SdkV2) {
	if !from.DataSynchronizationStatus.IsNull() && !from.DataSynchronizationStatus.IsUnknown() {
		if toDataSynchronizationStatus, ok := to.GetDataSynchronizationStatus(ctx); ok {
			if fromDataSynchronizationStatus, ok := from.GetDataSynchronizationStatus(ctx); ok {
				// Recursively sync the fields of DataSynchronizationStatus
				toDataSynchronizationStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromDataSynchronizationStatus)
				to.SetDataSynchronizationStatus(ctx, toDataSynchronizationStatus)
			}
		}
	}
	if !from.DatabaseInstanceName.IsUnknown() && !from.DatabaseInstanceName.IsNull() {
		// DatabaseInstanceName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DatabaseInstanceName = from.DatabaseInstanceName
	}
	if !from.LogicalDatabaseName.IsUnknown() && !from.LogicalDatabaseName.IsNull() {
		// LogicalDatabaseName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.LogicalDatabaseName = from.LogicalDatabaseName
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				// Recursively sync the fields of Spec
				toSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
}

func (to *SyncedDatabaseTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedDatabaseTable_SdkV2) {
	if !from.DataSynchronizationStatus.IsNull() && !from.DataSynchronizationStatus.IsUnknown() {
		if toDataSynchronizationStatus, ok := to.GetDataSynchronizationStatus(ctx); ok {
			if fromDataSynchronizationStatus, ok := from.GetDataSynchronizationStatus(ctx); ok {
				toDataSynchronizationStatus.SyncFieldsDuringRead(ctx, fromDataSynchronizationStatus)
				to.SetDataSynchronizationStatus(ctx, toDataSynchronizationStatus)
			}
		}
	}
	if !from.DatabaseInstanceName.IsUnknown() && !from.DatabaseInstanceName.IsNull() {
		// DatabaseInstanceName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.DatabaseInstanceName = from.DatabaseInstanceName
	}
	if !from.LogicalDatabaseName.IsUnknown() && !from.LogicalDatabaseName.IsNull() {
		// LogicalDatabaseName is an input only field and not returned by the service, so we keep the value from the prior state.
		to.LogicalDatabaseName = from.LogicalDatabaseName
	}
	if !from.Spec.IsNull() && !from.Spec.IsUnknown() {
		if toSpec, ok := to.GetSpec(ctx); ok {
			if fromSpec, ok := from.GetSpec(ctx); ok {
				toSpec.SyncFieldsDuringRead(ctx, fromSpec)
				to.SetSpec(ctx, toSpec)
			}
		}
	}
}

func (m SyncedDatabaseTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].SetComputed()
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()
	attrs["database_instance_name"] = attrs["database_instance_name"].SetComputed()
	attrs["database_instance_name"] = attrs["database_instance_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_database_instance_name"] = attrs["effective_database_instance_name"].SetComputed()
	attrs["effective_logical_database_name"] = attrs["effective_logical_database_name"].SetComputed()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetOptional()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetComputed()
	attrs["logical_database_name"] = attrs["logical_database_name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["unity_catalog_provisioning_state"] = attrs["unity_catalog_provisioning_state"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedDatabaseTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedDatabaseTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_synchronization_status": reflect.TypeOf(SyncedTableStatus_SdkV2{}),
		"spec":                        reflect.TypeOf(SyncedTableSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedDatabaseTable_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedDatabaseTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_synchronization_status":      m.DataSynchronizationStatus,
			"database_instance_name":           m.DatabaseInstanceName,
			"effective_database_instance_name": m.EffectiveDatabaseInstanceName,
			"effective_logical_database_name":  m.EffectiveLogicalDatabaseName,
			"logical_database_name":            m.LogicalDatabaseName,
			"name":                             m.Name,
			"spec":                             m.Spec,
			"unity_catalog_provisioning_state": m.UnityCatalogProvisioningState,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedDatabaseTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_synchronization_status": basetypes.ListType{
				ElemType: SyncedTableStatus_SdkV2{}.Type(ctx),
			},
			"database_instance_name":           types.StringType,
			"effective_database_instance_name": types.StringType,
			"effective_logical_database_name":  types.StringType,
			"logical_database_name":            types.StringType,
			"name":                             types.StringType,
			"spec": basetypes.ListType{
				ElemType: SyncedTableSpec_SdkV2{}.Type(ctx),
			},
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetDataSynchronizationStatus returns the value of the DataSynchronizationStatus field in SyncedDatabaseTable_SdkV2 as
// a SyncedTableStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedDatabaseTable_SdkV2) GetDataSynchronizationStatus(ctx context.Context) (SyncedTableStatus_SdkV2, bool) {
	var e SyncedTableStatus_SdkV2
	if m.DataSynchronizationStatus.IsNull() || m.DataSynchronizationStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableStatus_SdkV2
	d := m.DataSynchronizationStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataSynchronizationStatus sets the value of the DataSynchronizationStatus field in SyncedDatabaseTable_SdkV2.
func (m *SyncedDatabaseTable_SdkV2) SetDataSynchronizationStatus(ctx context.Context, v SyncedTableStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_synchronization_status"]
	m.DataSynchronizationStatus = types.ListValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in SyncedDatabaseTable_SdkV2 as
// a SyncedTableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedDatabaseTable_SdkV2) GetSpec(ctx context.Context) (SyncedTableSpec_SdkV2, bool) {
	var e SyncedTableSpec_SdkV2
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v []SyncedTableSpec_SdkV2
	d := m.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in SyncedDatabaseTable_SdkV2.
func (m *SyncedDatabaseTable_SdkV2) SetSpec(ctx context.Context, v SyncedTableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	m.Spec = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus_SdkV2 struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress"`
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (to *SyncedTableContinuousUpdateStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableContinuousUpdateStatus_SdkV2) {
	if !from.InitialPipelineSyncProgress.IsNull() && !from.InitialPipelineSyncProgress.IsUnknown() {
		if toInitialPipelineSyncProgress, ok := to.GetInitialPipelineSyncProgress(ctx); ok {
			if fromInitialPipelineSyncProgress, ok := from.GetInitialPipelineSyncProgress(ctx); ok {
				// Recursively sync the fields of InitialPipelineSyncProgress
				toInitialPipelineSyncProgress.SyncFieldsDuringCreateOrUpdate(ctx, fromInitialPipelineSyncProgress)
				to.SetInitialPipelineSyncProgress(ctx, toInitialPipelineSyncProgress)
			}
		}
	}
}

func (to *SyncedTableContinuousUpdateStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTableContinuousUpdateStatus_SdkV2) {
	if !from.InitialPipelineSyncProgress.IsNull() && !from.InitialPipelineSyncProgress.IsUnknown() {
		if toInitialPipelineSyncProgress, ok := to.GetInitialPipelineSyncProgress(ctx); ok {
			if fromInitialPipelineSyncProgress, ok := from.GetInitialPipelineSyncProgress(ctx); ok {
				toInitialPipelineSyncProgress.SyncFieldsDuringRead(ctx, fromInitialPipelineSyncProgress)
				to.SetInitialPipelineSyncProgress(ctx, toInitialPipelineSyncProgress)
			}
		}
	}
}

func (m SyncedTableContinuousUpdateStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetComputed()
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetComputed()
	attrs["timestamp"] = attrs["timestamp"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableContinuousUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTableContinuousUpdateStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableContinuousUpdateStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTableContinuousUpdateStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": m.InitialPipelineSyncProgress,
			"last_processed_commit_version":  m.LastProcessedCommitVersion,
			"timestamp":                      m.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableContinuousUpdateStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": basetypes.ListType{
				ElemType: SyncedTablePipelineProgress_SdkV2{}.Type(ctx),
			},
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus_SdkV2 as
// a SyncedTablePipelineProgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableContinuousUpdateStatus_SdkV2) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress_SdkV2, bool) {
	var e SyncedTablePipelineProgress_SdkV2
	if m.InitialPipelineSyncProgress.IsNull() || m.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress_SdkV2
	d := m.InitialPipelineSyncProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus_SdkV2.
func (m *SyncedTableContinuousUpdateStatus_SdkV2) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["initial_pipeline_sync_progress"]
	m.InitialPipelineSyncProgress = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus_SdkV2 struct {
	// The last source table Delta version that was successfully synced to the
	// synced table. The last source table Delta version that was synced to the
	// synced table. Only populated if the table is still synced and available
	// for serving.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. Only populated if the table is still
	// synced and available for serving.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (to *SyncedTableFailedStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableFailedStatus_SdkV2) {
}

func (to *SyncedTableFailedStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTableFailedStatus_SdkV2) {
}

func (m SyncedTableFailedStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetComputed()
	attrs["timestamp"] = attrs["timestamp"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableFailedStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTableFailedStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableFailedStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTableFailedStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": m.LastProcessedCommitVersion,
			"timestamp":                     m.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableFailedStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress_SdkV2 struct {
	// The estimated time remaining to complete this update in seconds.
	EstimatedCompletionTimeSeconds types.Float64 `tfsdk:"estimated_completion_time_seconds"`
	// The source table Delta version that was last processed by the pipeline.
	// The pipeline may not have completely processed this version yet.
	LatestVersionCurrentlyProcessing types.Int64 `tfsdk:"latest_version_currently_processing"`
	// The current phase of the data synchronization pipeline.
	ProvisioningPhase types.String `tfsdk:"provisioning_phase"`
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion types.Float64 `tfsdk:"sync_progress_completion"`
	// The number of rows that have been synced in this update.
	SyncedRowCount types.Int64 `tfsdk:"synced_row_count"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount types.Int64 `tfsdk:"total_row_count"`
}

func (to *SyncedTablePipelineProgress_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTablePipelineProgress_SdkV2) {
}

func (to *SyncedTablePipelineProgress_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTablePipelineProgress_SdkV2) {
}

func (m SyncedTablePipelineProgress_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["estimated_completion_time_seconds"] = attrs["estimated_completion_time_seconds"].SetComputed()
	attrs["latest_version_currently_processing"] = attrs["latest_version_currently_processing"].SetComputed()
	attrs["provisioning_phase"] = attrs["provisioning_phase"].SetComputed()
	attrs["sync_progress_completion"] = attrs["sync_progress_completion"].SetComputed()
	attrs["synced_row_count"] = attrs["synced_row_count"].SetComputed()
	attrs["total_row_count"] = attrs["total_row_count"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTablePipelineProgress.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTablePipelineProgress_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePipelineProgress_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTablePipelineProgress_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"estimated_completion_time_seconds":   m.EstimatedCompletionTimeSeconds,
			"latest_version_currently_processing": m.LatestVersionCurrentlyProcessing,
			"provisioning_phase":                  m.ProvisioningPhase,
			"sync_progress_completion":            m.SyncProgressCompletion,
			"synced_row_count":                    m.SyncedRowCount,
			"total_row_count":                     m.TotalRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTablePipelineProgress_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"estimated_completion_time_seconds":   types.Float64Type,
			"latest_version_currently_processing": types.Int64Type,
			"provisioning_phase":                  types.StringType,
			"sync_progress_completion":            types.Float64Type,
			"synced_row_count":                    types.Int64Type,
			"total_row_count":                     types.Int64Type,
		},
	}
}

type SyncedTablePosition_SdkV2 struct {
	DeltaTableSyncInfo types.List `tfsdk:"delta_table_sync_info"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	SyncEndTimestamp types.String `tfsdk:"sync_end_timestamp"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	SyncStartTimestamp types.String `tfsdk:"sync_start_timestamp"`
}

func (to *SyncedTablePosition_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTablePosition_SdkV2) {
	if !from.DeltaTableSyncInfo.IsNull() && !from.DeltaTableSyncInfo.IsUnknown() {
		if toDeltaTableSyncInfo, ok := to.GetDeltaTableSyncInfo(ctx); ok {
			if fromDeltaTableSyncInfo, ok := from.GetDeltaTableSyncInfo(ctx); ok {
				// Recursively sync the fields of DeltaTableSyncInfo
				toDeltaTableSyncInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromDeltaTableSyncInfo)
				to.SetDeltaTableSyncInfo(ctx, toDeltaTableSyncInfo)
			}
		}
	}
}

func (to *SyncedTablePosition_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTablePosition_SdkV2) {
	if !from.DeltaTableSyncInfo.IsNull() && !from.DeltaTableSyncInfo.IsUnknown() {
		if toDeltaTableSyncInfo, ok := to.GetDeltaTableSyncInfo(ctx); ok {
			if fromDeltaTableSyncInfo, ok := from.GetDeltaTableSyncInfo(ctx); ok {
				toDeltaTableSyncInfo.SyncFieldsDuringRead(ctx, fromDeltaTableSyncInfo)
				to.SetDeltaTableSyncInfo(ctx, toDeltaTableSyncInfo)
			}
		}
	}
}

func (m SyncedTablePosition_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_table_sync_info"] = attrs["delta_table_sync_info"].SetComputed()
	attrs["delta_table_sync_info"] = attrs["delta_table_sync_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["sync_end_timestamp"] = attrs["sync_end_timestamp"].SetComputed()
	attrs["sync_start_timestamp"] = attrs["sync_start_timestamp"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTablePosition.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTablePosition_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_table_sync_info": reflect.TypeOf(DeltaTableSyncInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePosition_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTablePosition_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_table_sync_info": m.DeltaTableSyncInfo,
			"sync_end_timestamp":    m.SyncEndTimestamp,
			"sync_start_timestamp":  m.SyncStartTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTablePosition_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_table_sync_info": basetypes.ListType{
				ElemType: DeltaTableSyncInfo_SdkV2{}.Type(ctx),
			},
			"sync_end_timestamp":   types.StringType,
			"sync_start_timestamp": types.StringType,
		},
	}
}

// GetDeltaTableSyncInfo returns the value of the DeltaTableSyncInfo field in SyncedTablePosition_SdkV2 as
// a DeltaTableSyncInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTablePosition_SdkV2) GetDeltaTableSyncInfo(ctx context.Context) (DeltaTableSyncInfo_SdkV2, bool) {
	var e DeltaTableSyncInfo_SdkV2
	if m.DeltaTableSyncInfo.IsNull() || m.DeltaTableSyncInfo.IsUnknown() {
		return e, false
	}
	var v []DeltaTableSyncInfo_SdkV2
	d := m.DeltaTableSyncInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDeltaTableSyncInfo sets the value of the DeltaTableSyncInfo field in SyncedTablePosition_SdkV2.
func (m *SyncedTablePosition_SdkV2) SetDeltaTableSyncInfo(ctx context.Context, v DeltaTableSyncInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["delta_table_sync_info"]
	m.DeltaTableSyncInfo = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus_SdkV2 struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress"`
}

func (to *SyncedTableProvisioningStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableProvisioningStatus_SdkV2) {
	if !from.InitialPipelineSyncProgress.IsNull() && !from.InitialPipelineSyncProgress.IsUnknown() {
		if toInitialPipelineSyncProgress, ok := to.GetInitialPipelineSyncProgress(ctx); ok {
			if fromInitialPipelineSyncProgress, ok := from.GetInitialPipelineSyncProgress(ctx); ok {
				// Recursively sync the fields of InitialPipelineSyncProgress
				toInitialPipelineSyncProgress.SyncFieldsDuringCreateOrUpdate(ctx, fromInitialPipelineSyncProgress)
				to.SetInitialPipelineSyncProgress(ctx, toInitialPipelineSyncProgress)
			}
		}
	}
}

func (to *SyncedTableProvisioningStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTableProvisioningStatus_SdkV2) {
	if !from.InitialPipelineSyncProgress.IsNull() && !from.InitialPipelineSyncProgress.IsUnknown() {
		if toInitialPipelineSyncProgress, ok := to.GetInitialPipelineSyncProgress(ctx); ok {
			if fromInitialPipelineSyncProgress, ok := from.GetInitialPipelineSyncProgress(ctx); ok {
				toInitialPipelineSyncProgress.SyncFieldsDuringRead(ctx, fromInitialPipelineSyncProgress)
				to.SetInitialPipelineSyncProgress(ctx, toInitialPipelineSyncProgress)
			}
		}
	}
}

func (m SyncedTableProvisioningStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetComputed()
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableProvisioningStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTableProvisioningStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableProvisioningStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTableProvisioningStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": m.InitialPipelineSyncProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableProvisioningStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": basetypes.ListType{
				ElemType: SyncedTablePipelineProgress_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus_SdkV2 as
// a SyncedTablePipelineProgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableProvisioningStatus_SdkV2) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress_SdkV2, bool) {
	var e SyncedTablePipelineProgress_SdkV2
	if m.InitialPipelineSyncProgress.IsNull() || m.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress_SdkV2
	d := m.InitialPipelineSyncProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus_SdkV2.
func (m *SyncedTableProvisioningStatus_SdkV2) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["initial_pipeline_sync_progress"]
	m.InitialPipelineSyncProgress = types.ListValueMust(t, vs)
}

// Specification of a synced database table.
type SyncedTableSpec_SdkV2 struct {
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist.
	CreateDatabaseObjectsIfMissing types.Bool `tfsdk:"create_database_objects_if_missing"`
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If existing_pipeline_id is defined, the synced table will be bin packed
	// into the existing pipeline referenced. This avoids creating a new
	// pipeline and allows sharing existing compute. In this case, the
	// scheduling_policy of this synced table must match the scheduling policy
	// of the existing pipeline.
	ExistingPipelineId types.String `tfsdk:"existing_pipeline_id"`
	// At most one of existing_pipeline_id and new_pipeline_spec should be
	// defined.
	//
	// If new_pipeline_spec is defined, a new pipeline is created for this
	// synced table. The location pointed to is used to store intermediate files
	// (checkpoints, event logs etc). The caller must have write permissions to
	// create Delta tables in the specified catalog and schema. Again, note this
	// requires write permissions, whereas the source table only requires read
	// permissions.
	NewPipelineSpec types.List `tfsdk:"new_pipeline_spec"`
	// Primary Key columns to be used for data insert/update in the destination.
	PrimaryKeyColumns types.List `tfsdk:"primary_key_columns"`
	// Scheduling policy of the underlying pipeline.
	SchedulingPolicy types.String `tfsdk:"scheduling_policy"`
	// Three-part (catalog, schema, table) name of the source Delta table.
	SourceTableFullName types.String `tfsdk:"source_table_full_name"`
	// Time series key to deduplicate (tie-break) rows with the same primary
	// key.
	TimeseriesKey types.String `tfsdk:"timeseries_key"`
}

func (to *SyncedTableSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableSpec_SdkV2) {
	if !from.CreateDatabaseObjectsIfMissing.IsUnknown() && !from.CreateDatabaseObjectsIfMissing.IsNull() {
		// CreateDatabaseObjectsIfMissing is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseObjectsIfMissing = from.CreateDatabaseObjectsIfMissing
	}
	if !from.ExistingPipelineId.IsUnknown() && !from.ExistingPipelineId.IsNull() {
		// ExistingPipelineId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.ExistingPipelineId = from.ExistingPipelineId
	}
	if !from.NewPipelineSpec.IsUnknown() && !from.NewPipelineSpec.IsNull() {
		// NewPipelineSpec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.NewPipelineSpec = from.NewPipelineSpec
	}
	if !from.NewPipelineSpec.IsNull() && !from.NewPipelineSpec.IsUnknown() {
		if toNewPipelineSpec, ok := to.GetNewPipelineSpec(ctx); ok {
			if fromNewPipelineSpec, ok := from.GetNewPipelineSpec(ctx); ok {
				// Recursively sync the fields of NewPipelineSpec
				toNewPipelineSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromNewPipelineSpec)
				to.SetNewPipelineSpec(ctx, toNewPipelineSpec)
			}
		}
	}
	if !from.PrimaryKeyColumns.IsNull() && !from.PrimaryKeyColumns.IsUnknown() && to.PrimaryKeyColumns.IsNull() && len(from.PrimaryKeyColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrimaryKeyColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrimaryKeyColumns = from.PrimaryKeyColumns
	}
}

func (to *SyncedTableSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTableSpec_SdkV2) {
	if !from.CreateDatabaseObjectsIfMissing.IsUnknown() && !from.CreateDatabaseObjectsIfMissing.IsNull() {
		// CreateDatabaseObjectsIfMissing is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseObjectsIfMissing = from.CreateDatabaseObjectsIfMissing
	}
	if !from.ExistingPipelineId.IsUnknown() && !from.ExistingPipelineId.IsNull() {
		// ExistingPipelineId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.ExistingPipelineId = from.ExistingPipelineId
	}
	if !from.NewPipelineSpec.IsUnknown() && !from.NewPipelineSpec.IsNull() {
		// NewPipelineSpec is an input only field and not returned by the service, so we keep the value from the prior state.
		to.NewPipelineSpec = from.NewPipelineSpec
	}
	if !from.NewPipelineSpec.IsNull() && !from.NewPipelineSpec.IsUnknown() {
		if toNewPipelineSpec, ok := to.GetNewPipelineSpec(ctx); ok {
			if fromNewPipelineSpec, ok := from.GetNewPipelineSpec(ctx); ok {
				toNewPipelineSpec.SyncFieldsDuringRead(ctx, fromNewPipelineSpec)
				to.SetNewPipelineSpec(ctx, toNewPipelineSpec)
			}
		}
	}
	if !from.PrimaryKeyColumns.IsNull() && !from.PrimaryKeyColumns.IsUnknown() && to.PrimaryKeyColumns.IsNull() && len(from.PrimaryKeyColumns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PrimaryKeyColumns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PrimaryKeyColumns = from.PrimaryKeyColumns
	}
}

func (m SyncedTableSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetOptional()
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetComputed()
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].SetOptional()
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].SetComputed()
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetOptional()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetComputed()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].(tfschema.ListNestedAttributeBuilder).AddPlanModifier(listplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["primary_key_columns"] = attrs["primary_key_columns"].SetOptional()
	attrs["scheduling_policy"] = attrs["scheduling_policy"].SetOptional()
	attrs["source_table_full_name"] = attrs["source_table_full_name"].SetOptional()
	attrs["timeseries_key"] = attrs["timeseries_key"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_pipeline_spec":   reflect.TypeOf(NewPipelineSpec_SdkV2{}),
		"primary_key_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_objects_if_missing": m.CreateDatabaseObjectsIfMissing,
			"existing_pipeline_id":               m.ExistingPipelineId,
			"new_pipeline_spec":                  m.NewPipelineSpec,
			"primary_key_columns":                m.PrimaryKeyColumns,
			"scheduling_policy":                  m.SchedulingPolicy,
			"source_table_full_name":             m.SourceTableFullName,
			"timeseries_key":                     m.TimeseriesKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_objects_if_missing": types.BoolType,
			"existing_pipeline_id":               types.StringType,
			"new_pipeline_spec": basetypes.ListType{
				ElemType: NewPipelineSpec_SdkV2{}.Type(ctx),
			},
			"primary_key_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scheduling_policy":      types.StringType,
			"source_table_full_name": types.StringType,
			"timeseries_key":         types.StringType,
		},
	}
}

// GetNewPipelineSpec returns the value of the NewPipelineSpec field in SyncedTableSpec_SdkV2 as
// a NewPipelineSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableSpec_SdkV2) GetNewPipelineSpec(ctx context.Context) (NewPipelineSpec_SdkV2, bool) {
	var e NewPipelineSpec_SdkV2
	if m.NewPipelineSpec.IsNull() || m.NewPipelineSpec.IsUnknown() {
		return e, false
	}
	var v []NewPipelineSpec_SdkV2
	d := m.NewPipelineSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewPipelineSpec sets the value of the NewPipelineSpec field in SyncedTableSpec_SdkV2.
func (m *SyncedTableSpec_SdkV2) SetNewPipelineSpec(ctx context.Context, v NewPipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["new_pipeline_spec"]
	m.NewPipelineSpec = types.ListValueMust(t, vs)
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in SyncedTableSpec_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableSpec_SdkV2) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
	if m.PrimaryKeyColumns.IsNull() || m.PrimaryKeyColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.PrimaryKeyColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in SyncedTableSpec_SdkV2.
func (m *SyncedTableSpec_SdkV2) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeyColumns = types.ListValueMust(t, vs)
}

// Status of a synced table.
type SyncedTableStatus_SdkV2 struct {
	ContinuousUpdateStatus types.List `tfsdk:"continuous_update_status"`
	// The state of the synced table.
	DetailedState types.String `tfsdk:"detailed_state"`

	FailedStatus types.List `tfsdk:"failed_status"`
	// Summary of the last successful synchronization from source to
	// destination.
	//
	// Will always be present if there has been a successful sync. Even if the
	// most recent syncs have failed.
	//
	// Limitation: The only exception is if the synced table is doing a FULL
	// REFRESH, then the last sync information will not be available until the
	// full refresh is complete. This limitation will be addressed in a future
	// version.
	//
	// This top-level field is a convenience for consumers who want easy access
	// to last sync information without having to traverse detailed_status.
	LastSync types.List `tfsdk:"last_sync"`
	// A text description of the current state of the synced table.
	Message types.String `tfsdk:"message"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	PipelineId types.String `tfsdk:"pipeline_id"`

	ProvisioningStatus types.List `tfsdk:"provisioning_status"`

	TriggeredUpdateStatus types.List `tfsdk:"triggered_update_status"`
}

func (to *SyncedTableStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableStatus_SdkV2) {
	if !from.ContinuousUpdateStatus.IsNull() && !from.ContinuousUpdateStatus.IsUnknown() {
		if toContinuousUpdateStatus, ok := to.GetContinuousUpdateStatus(ctx); ok {
			if fromContinuousUpdateStatus, ok := from.GetContinuousUpdateStatus(ctx); ok {
				// Recursively sync the fields of ContinuousUpdateStatus
				toContinuousUpdateStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromContinuousUpdateStatus)
				to.SetContinuousUpdateStatus(ctx, toContinuousUpdateStatus)
			}
		}
	}
	if !from.FailedStatus.IsNull() && !from.FailedStatus.IsUnknown() {
		if toFailedStatus, ok := to.GetFailedStatus(ctx); ok {
			if fromFailedStatus, ok := from.GetFailedStatus(ctx); ok {
				// Recursively sync the fields of FailedStatus
				toFailedStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromFailedStatus)
				to.SetFailedStatus(ctx, toFailedStatus)
			}
		}
	}
	if !from.LastSync.IsNull() && !from.LastSync.IsUnknown() {
		if toLastSync, ok := to.GetLastSync(ctx); ok {
			if fromLastSync, ok := from.GetLastSync(ctx); ok {
				// Recursively sync the fields of LastSync
				toLastSync.SyncFieldsDuringCreateOrUpdate(ctx, fromLastSync)
				to.SetLastSync(ctx, toLastSync)
			}
		}
	}
	if !from.ProvisioningStatus.IsNull() && !from.ProvisioningStatus.IsUnknown() {
		if toProvisioningStatus, ok := to.GetProvisioningStatus(ctx); ok {
			if fromProvisioningStatus, ok := from.GetProvisioningStatus(ctx); ok {
				// Recursively sync the fields of ProvisioningStatus
				toProvisioningStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromProvisioningStatus)
				to.SetProvisioningStatus(ctx, toProvisioningStatus)
			}
		}
	}
	if !from.TriggeredUpdateStatus.IsNull() && !from.TriggeredUpdateStatus.IsUnknown() {
		if toTriggeredUpdateStatus, ok := to.GetTriggeredUpdateStatus(ctx); ok {
			if fromTriggeredUpdateStatus, ok := from.GetTriggeredUpdateStatus(ctx); ok {
				// Recursively sync the fields of TriggeredUpdateStatus
				toTriggeredUpdateStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromTriggeredUpdateStatus)
				to.SetTriggeredUpdateStatus(ctx, toTriggeredUpdateStatus)
			}
		}
	}
}

func (to *SyncedTableStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTableStatus_SdkV2) {
	if !from.ContinuousUpdateStatus.IsNull() && !from.ContinuousUpdateStatus.IsUnknown() {
		if toContinuousUpdateStatus, ok := to.GetContinuousUpdateStatus(ctx); ok {
			if fromContinuousUpdateStatus, ok := from.GetContinuousUpdateStatus(ctx); ok {
				toContinuousUpdateStatus.SyncFieldsDuringRead(ctx, fromContinuousUpdateStatus)
				to.SetContinuousUpdateStatus(ctx, toContinuousUpdateStatus)
			}
		}
	}
	if !from.FailedStatus.IsNull() && !from.FailedStatus.IsUnknown() {
		if toFailedStatus, ok := to.GetFailedStatus(ctx); ok {
			if fromFailedStatus, ok := from.GetFailedStatus(ctx); ok {
				toFailedStatus.SyncFieldsDuringRead(ctx, fromFailedStatus)
				to.SetFailedStatus(ctx, toFailedStatus)
			}
		}
	}
	if !from.LastSync.IsNull() && !from.LastSync.IsUnknown() {
		if toLastSync, ok := to.GetLastSync(ctx); ok {
			if fromLastSync, ok := from.GetLastSync(ctx); ok {
				toLastSync.SyncFieldsDuringRead(ctx, fromLastSync)
				to.SetLastSync(ctx, toLastSync)
			}
		}
	}
	if !from.ProvisioningStatus.IsNull() && !from.ProvisioningStatus.IsUnknown() {
		if toProvisioningStatus, ok := to.GetProvisioningStatus(ctx); ok {
			if fromProvisioningStatus, ok := from.GetProvisioningStatus(ctx); ok {
				toProvisioningStatus.SyncFieldsDuringRead(ctx, fromProvisioningStatus)
				to.SetProvisioningStatus(ctx, toProvisioningStatus)
			}
		}
	}
	if !from.TriggeredUpdateStatus.IsNull() && !from.TriggeredUpdateStatus.IsUnknown() {
		if toTriggeredUpdateStatus, ok := to.GetTriggeredUpdateStatus(ctx); ok {
			if fromTriggeredUpdateStatus, ok := from.GetTriggeredUpdateStatus(ctx); ok {
				toTriggeredUpdateStatus.SyncFieldsDuringRead(ctx, fromTriggeredUpdateStatus)
				to.SetTriggeredUpdateStatus(ctx, toTriggeredUpdateStatus)
			}
		}
	}
}

func (m SyncedTableStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["continuous_update_status"] = attrs["continuous_update_status"].SetOptional()
	attrs["continuous_update_status"] = attrs["continuous_update_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["detailed_state"] = attrs["detailed_state"].SetComputed()
	attrs["failed_status"] = attrs["failed_status"].SetOptional()
	attrs["failed_status"] = attrs["failed_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["last_sync"] = attrs["last_sync"].SetComputed()
	attrs["last_sync"] = attrs["last_sync"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["message"] = attrs["message"].SetComputed()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetComputed()
	attrs["provisioning_status"] = attrs["provisioning_status"].SetOptional()
	attrs["provisioning_status"] = attrs["provisioning_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["triggered_update_status"] = attrs["triggered_update_status"].SetOptional()
	attrs["triggered_update_status"] = attrs["triggered_update_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTableStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous_update_status": reflect.TypeOf(SyncedTableContinuousUpdateStatus_SdkV2{}),
		"failed_status":            reflect.TypeOf(SyncedTableFailedStatus_SdkV2{}),
		"last_sync":                reflect.TypeOf(SyncedTablePosition_SdkV2{}),
		"provisioning_status":      reflect.TypeOf(SyncedTableProvisioningStatus_SdkV2{}),
		"triggered_update_status":  reflect.TypeOf(SyncedTableTriggeredUpdateStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTableStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"continuous_update_status": m.ContinuousUpdateStatus,
			"detailed_state":           m.DetailedState,
			"failed_status":            m.FailedStatus,
			"last_sync":                m.LastSync,
			"message":                  m.Message,
			"pipeline_id":              m.PipelineId,
			"provisioning_status":      m.ProvisioningStatus,
			"triggered_update_status":  m.TriggeredUpdateStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"continuous_update_status": basetypes.ListType{
				ElemType: SyncedTableContinuousUpdateStatus_SdkV2{}.Type(ctx),
			},
			"detailed_state": types.StringType,
			"failed_status": basetypes.ListType{
				ElemType: SyncedTableFailedStatus_SdkV2{}.Type(ctx),
			},
			"last_sync": basetypes.ListType{
				ElemType: SyncedTablePosition_SdkV2{}.Type(ctx),
			},
			"message":     types.StringType,
			"pipeline_id": types.StringType,
			"provisioning_status": basetypes.ListType{
				ElemType: SyncedTableProvisioningStatus_SdkV2{}.Type(ctx),
			},
			"triggered_update_status": basetypes.ListType{
				ElemType: SyncedTableTriggeredUpdateStatus_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetContinuousUpdateStatus returns the value of the ContinuousUpdateStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableContinuousUpdateStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus_SdkV2) GetContinuousUpdateStatus(ctx context.Context) (SyncedTableContinuousUpdateStatus_SdkV2, bool) {
	var e SyncedTableContinuousUpdateStatus_SdkV2
	if m.ContinuousUpdateStatus.IsNull() || m.ContinuousUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableContinuousUpdateStatus_SdkV2
	d := m.ContinuousUpdateStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContinuousUpdateStatus sets the value of the ContinuousUpdateStatus field in SyncedTableStatus_SdkV2.
func (m *SyncedTableStatus_SdkV2) SetContinuousUpdateStatus(ctx context.Context, v SyncedTableContinuousUpdateStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["continuous_update_status"]
	m.ContinuousUpdateStatus = types.ListValueMust(t, vs)
}

// GetFailedStatus returns the value of the FailedStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableFailedStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus_SdkV2) GetFailedStatus(ctx context.Context) (SyncedTableFailedStatus_SdkV2, bool) {
	var e SyncedTableFailedStatus_SdkV2
	if m.FailedStatus.IsNull() || m.FailedStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableFailedStatus_SdkV2
	d := m.FailedStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailedStatus sets the value of the FailedStatus field in SyncedTableStatus_SdkV2.
func (m *SyncedTableStatus_SdkV2) SetFailedStatus(ctx context.Context, v SyncedTableFailedStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_status"]
	m.FailedStatus = types.ListValueMust(t, vs)
}

// GetLastSync returns the value of the LastSync field in SyncedTableStatus_SdkV2 as
// a SyncedTablePosition_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus_SdkV2) GetLastSync(ctx context.Context) (SyncedTablePosition_SdkV2, bool) {
	var e SyncedTablePosition_SdkV2
	if m.LastSync.IsNull() || m.LastSync.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePosition_SdkV2
	d := m.LastSync.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLastSync sets the value of the LastSync field in SyncedTableStatus_SdkV2.
func (m *SyncedTableStatus_SdkV2) SetLastSync(ctx context.Context, v SyncedTablePosition_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["last_sync"]
	m.LastSync = types.ListValueMust(t, vs)
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableProvisioningStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus_SdkV2) GetProvisioningStatus(ctx context.Context) (SyncedTableProvisioningStatus_SdkV2, bool) {
	var e SyncedTableProvisioningStatus_SdkV2
	if m.ProvisioningStatus.IsNull() || m.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableProvisioningStatus_SdkV2
	d := m.ProvisioningStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvisioningStatus sets the value of the ProvisioningStatus field in SyncedTableStatus_SdkV2.
func (m *SyncedTableStatus_SdkV2) SetProvisioningStatus(ctx context.Context, v SyncedTableProvisioningStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["provisioning_status"]
	m.ProvisioningStatus = types.ListValueMust(t, vs)
}

// GetTriggeredUpdateStatus returns the value of the TriggeredUpdateStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableTriggeredUpdateStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus_SdkV2) GetTriggeredUpdateStatus(ctx context.Context) (SyncedTableTriggeredUpdateStatus_SdkV2, bool) {
	var e SyncedTableTriggeredUpdateStatus_SdkV2
	if m.TriggeredUpdateStatus.IsNull() || m.TriggeredUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableTriggeredUpdateStatus_SdkV2
	d := m.TriggeredUpdateStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in SyncedTableStatus_SdkV2.
func (m *SyncedTableStatus_SdkV2) SetTriggeredUpdateStatus(ctx context.Context, v SyncedTableTriggeredUpdateStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["triggered_update_status"]
	m.TriggeredUpdateStatus = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus_SdkV2 struct {
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress types.List `tfsdk:"triggered_update_progress"`
}

func (to *SyncedTableTriggeredUpdateStatus_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableTriggeredUpdateStatus_SdkV2) {
	if !from.TriggeredUpdateProgress.IsNull() && !from.TriggeredUpdateProgress.IsUnknown() {
		if toTriggeredUpdateProgress, ok := to.GetTriggeredUpdateProgress(ctx); ok {
			if fromTriggeredUpdateProgress, ok := from.GetTriggeredUpdateProgress(ctx); ok {
				// Recursively sync the fields of TriggeredUpdateProgress
				toTriggeredUpdateProgress.SyncFieldsDuringCreateOrUpdate(ctx, fromTriggeredUpdateProgress)
				to.SetTriggeredUpdateProgress(ctx, toTriggeredUpdateProgress)
			}
		}
	}
}

func (to *SyncedTableTriggeredUpdateStatus_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SyncedTableTriggeredUpdateStatus_SdkV2) {
	if !from.TriggeredUpdateProgress.IsNull() && !from.TriggeredUpdateProgress.IsUnknown() {
		if toTriggeredUpdateProgress, ok := to.GetTriggeredUpdateProgress(ctx); ok {
			if fromTriggeredUpdateProgress, ok := from.GetTriggeredUpdateProgress(ctx); ok {
				toTriggeredUpdateProgress.SyncFieldsDuringRead(ctx, fromTriggeredUpdateProgress)
				to.SetTriggeredUpdateProgress(ctx, toTriggeredUpdateProgress)
			}
		}
	}
}

func (m SyncedTableTriggeredUpdateStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetComputed()
	attrs["timestamp"] = attrs["timestamp"].SetComputed()
	attrs["triggered_update_progress"] = attrs["triggered_update_progress"].SetComputed()
	attrs["triggered_update_progress"] = attrs["triggered_update_progress"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableTriggeredUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SyncedTableTriggeredUpdateStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"triggered_update_progress": reflect.TypeOf(SyncedTablePipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableTriggeredUpdateStatus_SdkV2
// only implements ToObjectValue() and Type().
func (m SyncedTableTriggeredUpdateStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": m.LastProcessedCommitVersion,
			"timestamp":                     m.Timestamp,
			"triggered_update_progress":     m.TriggeredUpdateProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableTriggeredUpdateStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
			"triggered_update_progress": basetypes.ListType{
				ElemType: SyncedTablePipelineProgress_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTriggeredUpdateProgress returns the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus_SdkV2 as
// a SyncedTablePipelineProgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableTriggeredUpdateStatus_SdkV2) GetTriggeredUpdateProgress(ctx context.Context) (SyncedTablePipelineProgress_SdkV2, bool) {
	var e SyncedTablePipelineProgress_SdkV2
	if m.TriggeredUpdateProgress.IsNull() || m.TriggeredUpdateProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress_SdkV2
	d := m.TriggeredUpdateProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus_SdkV2.
func (m *SyncedTableTriggeredUpdateStatus_SdkV2) SetTriggeredUpdateProgress(ctx context.Context, v SyncedTablePipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["triggered_update_progress"]
	m.TriggeredUpdateProgress = types.ListValueMust(t, vs)
}

type UpdateDatabaseCatalogRequest_SdkV2 struct {
	// Note that updating a database catalog is not yet supported.
	DatabaseCatalog types.List `tfsdk:"database_catalog"`
	// The name of the catalog in UC.
	Name types.String `tfsdk:"-"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseCatalogRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseCatalogRequest_SdkV2) {
	if !from.DatabaseCatalog.IsNull() && !from.DatabaseCatalog.IsUnknown() {
		if toDatabaseCatalog, ok := to.GetDatabaseCatalog(ctx); ok {
			if fromDatabaseCatalog, ok := from.GetDatabaseCatalog(ctx); ok {
				// Recursively sync the fields of DatabaseCatalog
				toDatabaseCatalog.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseCatalog)
				to.SetDatabaseCatalog(ctx, toDatabaseCatalog)
			}
		}
	}
}

func (to *UpdateDatabaseCatalogRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseCatalogRequest_SdkV2) {
	if !from.DatabaseCatalog.IsNull() && !from.DatabaseCatalog.IsUnknown() {
		if toDatabaseCatalog, ok := to.GetDatabaseCatalog(ctx); ok {
			if fromDatabaseCatalog, ok := from.GetDatabaseCatalog(ctx); ok {
				toDatabaseCatalog.SyncFieldsDuringRead(ctx, fromDatabaseCatalog)
				to.SetDatabaseCatalog(ctx, toDatabaseCatalog)
			}
		}
	}
}

func (m UpdateDatabaseCatalogRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_catalog"] = attrs["database_catalog"].SetRequired()
	attrs["database_catalog"] = attrs["database_catalog"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalog": reflect.TypeOf(DatabaseCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_catalog": m.DatabaseCatalog,
			"name":             m.Name,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_catalog": basetypes.ListType{
				ElemType: DatabaseCatalog_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetDatabaseCatalog returns the value of the DatabaseCatalog field in UpdateDatabaseCatalogRequest_SdkV2 as
// a DatabaseCatalog_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatabaseCatalogRequest_SdkV2) GetDatabaseCatalog(ctx context.Context) (DatabaseCatalog_SdkV2, bool) {
	var e DatabaseCatalog_SdkV2
	if m.DatabaseCatalog.IsNull() || m.DatabaseCatalog.IsUnknown() {
		return e, false
	}
	var v []DatabaseCatalog_SdkV2
	d := m.DatabaseCatalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseCatalog sets the value of the DatabaseCatalog field in UpdateDatabaseCatalogRequest_SdkV2.
func (m *UpdateDatabaseCatalogRequest_SdkV2) SetDatabaseCatalog(ctx context.Context, v DatabaseCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_catalog"]
	m.DatabaseCatalog = types.ListValueMust(t, vs)
}

type UpdateDatabaseInstanceRequest_SdkV2 struct {
	DatabaseInstance types.List `tfsdk:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible. To wipe out custom_tags, specify custom_tags in the
	// update_mask with an empty custom_tags map.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseInstanceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseInstanceRequest_SdkV2) {
	if !from.DatabaseInstance.IsNull() && !from.DatabaseInstance.IsUnknown() {
		if toDatabaseInstance, ok := to.GetDatabaseInstance(ctx); ok {
			if fromDatabaseInstance, ok := from.GetDatabaseInstance(ctx); ok {
				// Recursively sync the fields of DatabaseInstance
				toDatabaseInstance.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseInstance)
				to.SetDatabaseInstance(ctx, toDatabaseInstance)
			}
		}
	}
}

func (to *UpdateDatabaseInstanceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseInstanceRequest_SdkV2) {
	if !from.DatabaseInstance.IsNull() && !from.DatabaseInstance.IsUnknown() {
		if toDatabaseInstance, ok := to.GetDatabaseInstance(ctx); ok {
			if fromDatabaseInstance, ok := from.GetDatabaseInstance(ctx); ok {
				toDatabaseInstance.SyncFieldsDuringRead(ctx, fromDatabaseInstance)
				to.SetDatabaseInstance(ctx, toDatabaseInstance)
			}
		}
	}
}

func (m UpdateDatabaseInstanceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance"] = attrs["database_instance"].SetRequired()
	attrs["database_instance"] = attrs["database_instance"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": m.DatabaseInstance,
			"name":              m.Name,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": basetypes.ListType{
				ElemType: DatabaseInstance_SdkV2{}.Type(ctx),
			},
			"name":        types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest_SdkV2 as
// a DatabaseInstance_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatabaseInstanceRequest_SdkV2) GetDatabaseInstance(ctx context.Context) (DatabaseInstance_SdkV2, bool) {
	var e DatabaseInstance_SdkV2
	if m.DatabaseInstance.IsNull() || m.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance_SdkV2
	d := m.DatabaseInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest_SdkV2.
func (m *UpdateDatabaseInstanceRequest_SdkV2) SetDatabaseInstance(ctx context.Context, v DatabaseInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance"]
	m.DatabaseInstance = types.ListValueMust(t, vs)
}

type UpdateSyncedDatabaseTableRequest_SdkV2 struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
	// Note that updating a synced database table is not yet supported.
	SyncedTable types.List `tfsdk:"synced_table"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSyncedDatabaseTableRequest_SdkV2) {
	if !from.SyncedTable.IsNull() && !from.SyncedTable.IsUnknown() {
		if toSyncedTable, ok := to.GetSyncedTable(ctx); ok {
			if fromSyncedTable, ok := from.GetSyncedTable(ctx); ok {
				// Recursively sync the fields of SyncedTable
				toSyncedTable.SyncFieldsDuringCreateOrUpdate(ctx, fromSyncedTable)
				to.SetSyncedTable(ctx, toSyncedTable)
			}
		}
	}
}

func (to *UpdateSyncedDatabaseTableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSyncedDatabaseTableRequest_SdkV2) {
	if !from.SyncedTable.IsNull() && !from.SyncedTable.IsUnknown() {
		if toSyncedTable, ok := to.GetSyncedTable(ctx); ok {
			if fromSyncedTable, ok := from.GetSyncedTable(ctx); ok {
				toSyncedTable.SyncFieldsDuringRead(ctx, fromSyncedTable)
				to.SetSyncedTable(ctx, toSyncedTable)
			}
		}
	}
}

func (m UpdateSyncedDatabaseTableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["synced_table"] = attrs["synced_table"].SetRequired()
	attrs["synced_table"] = attrs["synced_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         m.Name,
			"synced_table": m.SyncedTable,
			"update_mask":  m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"synced_table": basetypes.ListType{
				ElemType: SyncedDatabaseTable_SdkV2{}.Type(ctx),
			},
			"update_mask": types.StringType,
		},
	}
}

// GetSyncedTable returns the value of the SyncedTable field in UpdateSyncedDatabaseTableRequest_SdkV2 as
// a SyncedDatabaseTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSyncedDatabaseTableRequest_SdkV2) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable_SdkV2, bool) {
	var e SyncedDatabaseTable_SdkV2
	if m.SyncedTable.IsNull() || m.SyncedTable.IsUnknown() {
		return e, false
	}
	var v []SyncedDatabaseTable_SdkV2
	d := m.SyncedTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSyncedTable sets the value of the SyncedTable field in UpdateSyncedDatabaseTableRequest_SdkV2.
func (m *UpdateSyncedDatabaseTableRequest_SdkV2) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["synced_table"]
	m.SyncedTable = types.ListValueMust(t, vs)
}
