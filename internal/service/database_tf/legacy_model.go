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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Create a Database Catalog
type CreateDatabaseCatalogRequest_SdkV2 struct {
	Catalog types.List `tfsdk:"catalog"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(DatabaseCatalog_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateDatabaseCatalogRequest_SdkV2) GetCatalog(ctx context.Context) (DatabaseCatalog_SdkV2, bool) {
	var e DatabaseCatalog_SdkV2
	if o.Catalog.IsNull() || o.Catalog.IsUnknown() {
		return e, false
	}
	var v []DatabaseCatalog_SdkV2
	d := o.Catalog.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCatalog sets the value of the Catalog field in CreateDatabaseCatalogRequest_SdkV2.
func (o *CreateDatabaseCatalogRequest_SdkV2) SetCatalog(ctx context.Context, v DatabaseCatalog_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["catalog"]
	o.Catalog = types.ListValueMust(t, vs)
}

// Create a Database Instance
type CreateDatabaseInstanceRequest_SdkV2 struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance types.List `tfsdk:"database_instance"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": o.DatabaseInstance,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateDatabaseInstanceRequest_SdkV2) GetDatabaseInstance(ctx context.Context) (DatabaseInstance_SdkV2, bool) {
	var e DatabaseInstance_SdkV2
	if o.DatabaseInstance.IsNull() || o.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance_SdkV2
	d := o.DatabaseInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in CreateDatabaseInstanceRequest_SdkV2.
func (o *CreateDatabaseInstanceRequest_SdkV2) SetDatabaseInstance(ctx context.Context, v DatabaseInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance"]
	o.DatabaseInstance = types.ListValueMust(t, vs)
}

// Create a Database Table
type CreateDatabaseTableRequest_SdkV2 struct {
	// Next field marker: 13
	Table types.List `tfsdk:"table"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(DatabaseTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateDatabaseTableRequest_SdkV2) GetTable(ctx context.Context) (DatabaseTable_SdkV2, bool) {
	var e DatabaseTable_SdkV2
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []DatabaseTable_SdkV2
	d := o.Table.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CreateDatabaseTableRequest_SdkV2.
func (o *CreateDatabaseTableRequest_SdkV2) SetTable(ctx context.Context, v DatabaseTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["table"]
	o.Table = types.ListValueMust(t, vs)
}

// Create a Synced Database Table
type CreateSyncedDatabaseTableRequest_SdkV2 struct {
	// Next field marker: 12
	SyncedTable types.List `tfsdk:"synced_table"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"synced_table": o.SyncedTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateSyncedDatabaseTableRequest_SdkV2) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable_SdkV2, bool) {
	var e SyncedDatabaseTable_SdkV2
	if o.SyncedTable.IsNull() || o.SyncedTable.IsUnknown() {
		return e, false
	}
	var v []SyncedDatabaseTable_SdkV2
	d := o.SyncedTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSyncedTable sets the value of the SyncedTable field in CreateSyncedDatabaseTableRequest_SdkV2.
func (o *CreateSyncedDatabaseTableRequest_SdkV2) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["synced_table"]
	o.SyncedTable = types.ListValueMust(t, vs)
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

func (newState *DatabaseCatalog_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseCatalog_SdkV2) {
}

func (newState *DatabaseCatalog_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabaseCatalog_SdkV2) {
}

func (c DatabaseCatalog_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetOptional()
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
func (a DatabaseCatalog_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCatalog_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabaseCatalog_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_if_not_exists": o.CreateDatabaseIfNotExists,
			"database_instance_name":        o.DatabaseInstanceName,
			"database_name":                 o.DatabaseName,
			"name":                          o.Name,
			"uid":                           o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseCatalog_SdkV2) Type(ctx context.Context) attr.Type {
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
	Token types.String `tfsdk:"token"`
}

func (newState *DatabaseCredential_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseCredential_SdkV2) {
}

func (newState *DatabaseCredential_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabaseCredential_SdkV2) {
}

func (c DatabaseCredential_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatabaseCredential_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCredential_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabaseCredential_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token": o.Token,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseCredential_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token": types.StringType,
		},
	}
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance_SdkV2 struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	Capacity types.String `tfsdk:"capacity"`
	// The timestamp when the instance was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the instance.
	Creator types.String `tfsdk:"creator"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"name"`
	// The version of Postgres running on the instance.
	PgVersion types.String `tfsdk:"pg_version"`
	// The DNS endpoint to connect to the instance for read+write access.
	ReadWriteDns types.String `tfsdk:"read_write_dns"`
	// The current state of the instance.
	State types.String `tfsdk:"state"`
	// Whether the instance is stopped.
	Stopped types.Bool `tfsdk:"stopped"`
	// An immutable UUID identifier for the instance.
	Uid types.String `tfsdk:"uid"`
}

func (newState *DatabaseInstance_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseInstance_SdkV2) {
}

func (newState *DatabaseInstance_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabaseInstance_SdkV2) {
}

func (c DatabaseInstance_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["read_write_dns"] = attrs["read_write_dns"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["stopped"] = attrs["stopped"].SetOptional()
	attrs["uid"] = attrs["uid"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstance.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabaseInstance_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabaseInstance_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"capacity":       o.Capacity,
			"creation_time":  o.CreationTime,
			"creator":        o.Creator,
			"name":           o.Name,
			"pg_version":     o.PgVersion,
			"read_write_dns": o.ReadWriteDns,
			"state":          o.State,
			"stopped":        o.Stopped,
			"uid":            o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstance_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"capacity":       types.StringType,
			"creation_time":  types.StringType,
			"creator":        types.StringType,
			"name":           types.StringType,
			"pg_version":     types.StringType,
			"read_write_dns": types.StringType,
			"state":          types.StringType,
			"stopped":        types.BoolType,
			"uid":            types.StringType,
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
	// This field is optional in all scenarios.
	//
	// When creating a table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a table in a standard catalog, the target database name is
	// inferred to be that of the standard catalog. In this scenario, specifying
	// this field will allow targeting an arbitrary postgres database. Note that
	// this has implications for the `create_database_objects_is_missing` field
	// in `spec`.
	LogicalDatabaseName types.String `tfsdk:"logical_database_name"`
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name"`
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url"`
}

func (newState *DatabaseTable_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseTable_SdkV2) {
}

func (newState *DatabaseTable_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabaseTable_SdkV2) {
}

func (c DatabaseTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["table_serving_url"] = attrs["table_serving_url"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabaseTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseTable_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabaseTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_name": o.DatabaseInstanceName,
			"logical_database_name":  o.LogicalDatabaseName,
			"name":                   o.Name,
			"table_serving_url":      o.TableServingUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_name": types.StringType,
			"logical_database_name":  types.StringType,
			"name":                   types.StringType,
			"table_serving_url":      types.StringType,
		},
	}
}

// Delete a Database Catalog
type DeleteDatabaseCatalogRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseCatalogResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseCatalogResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseCatalogResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseCatalogResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a Database Instance
type DeleteDatabaseInstanceRequest_SdkV2 struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force types.Bool `tfsdk:"-"`
	// Name of the instance to delete.
	Name types.String `tfsdk:"-"`
	// If false, the database instance is soft deleted. Soft deleted instances
	// behave as if they are deleted, and cannot be used for CRUD operations nor
	// connected to. However they can be undeleted by calling the undelete API
	// for a limited time. If true, the database instance is hard deleted and
	// cannot be undeleted.
	Purge types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
			"purge": o.Purge,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
			"purge": types.BoolType,
		},
	}
}

type DeleteDatabaseInstanceResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseInstanceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a Database Table
type DeleteDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseTableResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseTableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseTableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseTableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseTableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseTableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a Synced Database Table
type DeleteSyncedDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSyncedDatabaseTableResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSyncedDatabaseTableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteSyncedDatabaseTableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSyncedDatabaseTableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Find a Database Instance by uid
type FindDatabaseInstanceByUidRequest_SdkV2 struct {
	// UID of the cluster to get.
	Uid types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FindDatabaseInstanceByUidRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FindDatabaseInstanceByUidRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FindDatabaseInstanceByUidRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o FindDatabaseInstanceByUidRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"uid": o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FindDatabaseInstanceByUidRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uid": types.StringType,
		},
	}
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest_SdkV2 struct {
	// Instances to which the token will be scoped.
	InstanceNames types.List `tfsdk:"instance_names"`

	RequestId types.String `tfsdk:"request_id"`
}

func (newState *GenerateDatabaseCredentialRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateDatabaseCredentialRequest_SdkV2) {
}

func (newState *GenerateDatabaseCredentialRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenerateDatabaseCredentialRequest_SdkV2) {
}

func (c GenerateDatabaseCredentialRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenerateDatabaseCredentialRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateDatabaseCredentialRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GenerateDatabaseCredentialRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_names": o.InstanceNames,
			"request_id":     o.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateDatabaseCredentialRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"request_id": types.StringType,
		},
	}
}

// GetInstanceNames returns the value of the InstanceNames field in GenerateDatabaseCredentialRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateDatabaseCredentialRequest_SdkV2) GetInstanceNames(ctx context.Context) ([]types.String, bool) {
	if o.InstanceNames.IsNull() || o.InstanceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InstanceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstanceNames sets the value of the InstanceNames field in GenerateDatabaseCredentialRequest_SdkV2.
func (o *GenerateDatabaseCredentialRequest_SdkV2) SetInstanceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InstanceNames = types.ListValueMust(t, vs)
}

// Get a Database Catalog
type GetDatabaseCatalogRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseCatalogRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseCatalogRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDatabaseCatalogRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseCatalogRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a Database Instance
type GetDatabaseInstanceRequest_SdkV2 struct {
	// Name of the cluster to get.
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a Database Table
type GetDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// Get a Synced Database Table
type GetSyncedDatabaseTableRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSyncedDatabaseTableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSyncedDatabaseTableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetSyncedDatabaseTableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSyncedDatabaseTableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// List Database Instances
type ListDatabaseInstancesRequest_SdkV2 struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstancesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDatabaseInstancesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstancesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstancesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *ListDatabaseInstancesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDatabaseInstancesResponse_SdkV2) {
}

func (newState *ListDatabaseInstancesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListDatabaseInstancesResponse_SdkV2) {
}

func (c ListDatabaseInstancesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDatabaseInstancesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(DatabaseInstance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstancesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instances": o.DatabaseInstances,
			"next_page_token":    o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstancesResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListDatabaseInstancesResponse_SdkV2) GetDatabaseInstances(ctx context.Context) ([]DatabaseInstance_SdkV2, bool) {
	if o.DatabaseInstances.IsNull() || o.DatabaseInstances.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstance_SdkV2
	d := o.DatabaseInstances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstances sets the value of the DatabaseInstances field in ListDatabaseInstancesResponse_SdkV2.
func (o *ListDatabaseInstancesResponse_SdkV2) SetDatabaseInstances(ctx context.Context, v []DatabaseInstance_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DatabaseInstances = types.ListValueMust(t, vs)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec_SdkV2 struct {
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog types.String `tfsdk:"storage_catalog"`
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	StorageSchema types.String `tfsdk:"storage_schema"`
}

func (newState *NewPipelineSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NewPipelineSpec_SdkV2) {
}

func (newState *NewPipelineSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState NewPipelineSpec_SdkV2) {
}

func (c NewPipelineSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NewPipelineSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NewPipelineSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o NewPipelineSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_catalog": o.StorageCatalog,
			"storage_schema":  o.StorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NewPipelineSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_catalog": types.StringType,
			"storage_schema":  types.StringType,
		},
	}
}

// Next field marker: 12
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
	// Target Postgres database object (logical database) name for this table.
	// This field is optional in all scenarios.
	//
	// When creating a synced table in a registered Postgres catalog, the target
	// Postgres database name is inferred to be that of the registered catalog.
	// If this field is specified in this scenario, the Postgres database name
	// MUST match that of the registered catalog (or the request will be
	// rejected).
	//
	// When creating a synced table in a standard catalog, the target database
	// name is inferred to be that of the standard catalog. In this scenario,
	// specifying this field will allow targeting an arbitrary postgres
	// database.
	LogicalDatabaseName types.String `tfsdk:"logical_database_name"`
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"name"`
	// Specification of a synced database table.
	Spec types.List `tfsdk:"spec"`
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state"`
}

func (newState *SyncedDatabaseTable_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedDatabaseTable_SdkV2) {
}

func (newState *SyncedDatabaseTable_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedDatabaseTable_SdkV2) {
}

func (c SyncedDatabaseTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].SetComputed()
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["spec"] = attrs["spec"].SetOptional()
	attrs["spec"] = attrs["spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table_serving_url"] = attrs["table_serving_url"].SetComputed()
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
func (a SyncedDatabaseTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_synchronization_status": reflect.TypeOf(SyncedTableStatus_SdkV2{}),
		"spec":                        reflect.TypeOf(SyncedTableSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedDatabaseTable_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedDatabaseTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_synchronization_status":      o.DataSynchronizationStatus,
			"database_instance_name":           o.DatabaseInstanceName,
			"logical_database_name":            o.LogicalDatabaseName,
			"name":                             o.Name,
			"spec":                             o.Spec,
			"table_serving_url":                o.TableServingUrl,
			"unity_catalog_provisioning_state": o.UnityCatalogProvisioningState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedDatabaseTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_synchronization_status": basetypes.ListType{
				ElemType: SyncedTableStatus_SdkV2{}.Type(ctx),
			},
			"database_instance_name": types.StringType,
			"logical_database_name":  types.StringType,
			"name":                   types.StringType,
			"spec": basetypes.ListType{
				ElemType: SyncedTableSpec_SdkV2{}.Type(ctx),
			},
			"table_serving_url":                types.StringType,
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetDataSynchronizationStatus returns the value of the DataSynchronizationStatus field in SyncedDatabaseTable_SdkV2 as
// a SyncedTableStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedDatabaseTable_SdkV2) GetDataSynchronizationStatus(ctx context.Context) (SyncedTableStatus_SdkV2, bool) {
	var e SyncedTableStatus_SdkV2
	if o.DataSynchronizationStatus.IsNull() || o.DataSynchronizationStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableStatus_SdkV2
	d := o.DataSynchronizationStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataSynchronizationStatus sets the value of the DataSynchronizationStatus field in SyncedDatabaseTable_SdkV2.
func (o *SyncedDatabaseTable_SdkV2) SetDataSynchronizationStatus(ctx context.Context, v SyncedTableStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_synchronization_status"]
	o.DataSynchronizationStatus = types.ListValueMust(t, vs)
}

// GetSpec returns the value of the Spec field in SyncedDatabaseTable_SdkV2 as
// a SyncedTableSpec_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedDatabaseTable_SdkV2) GetSpec(ctx context.Context) (SyncedTableSpec_SdkV2, bool) {
	var e SyncedTableSpec_SdkV2
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []SyncedTableSpec_SdkV2
	d := o.Spec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in SyncedDatabaseTable_SdkV2.
func (o *SyncedDatabaseTable_SdkV2) SetSpec(ctx context.Context, v SyncedTableSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["spec"]
	o.Spec = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus_SdkV2 struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress"`
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may not be completely synced to the synced
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (newState *SyncedTableContinuousUpdateStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableContinuousUpdateStatus_SdkV2) {
}

func (newState *SyncedTableContinuousUpdateStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTableContinuousUpdateStatus_SdkV2) {
}

func (c SyncedTableContinuousUpdateStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetOptional()
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableContinuousUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTableContinuousUpdateStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableContinuousUpdateStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTableContinuousUpdateStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
			"last_processed_commit_version":  o.LastProcessedCommitVersion,
			"timestamp":                      o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableContinuousUpdateStatus_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SyncedTableContinuousUpdateStatus_SdkV2) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress_SdkV2, bool) {
	var e SyncedTablePipelineProgress_SdkV2
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress_SdkV2
	d := o.InitialPipelineSyncProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus_SdkV2.
func (o *SyncedTableContinuousUpdateStatus_SdkV2) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["initial_pipeline_sync_progress"]
	o.InitialPipelineSyncProgress = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus_SdkV2 struct {
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may only be partially synced to the synced
	// table. Only populated if the table is still synced and available for
	// serving.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table. Only populated if the table is still synced
	// and available for serving.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (newState *SyncedTableFailedStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableFailedStatus_SdkV2) {
}

func (newState *SyncedTableFailedStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTableFailedStatus_SdkV2) {
}

func (c SyncedTableFailedStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableFailedStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTableFailedStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableFailedStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTableFailedStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableFailedStatus_SdkV2) Type(ctx context.Context) attr.Type {
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
	// The completion ratio of this update. This is a number between 0 and 1.
	SyncProgressCompletion types.Float64 `tfsdk:"sync_progress_completion"`
	// The number of rows that have been synced in this update.
	SyncedRowCount types.Int64 `tfsdk:"synced_row_count"`
	// The total number of rows that need to be synced in this update. This
	// number may be an estimate.
	TotalRowCount types.Int64 `tfsdk:"total_row_count"`
}

func (newState *SyncedTablePipelineProgress_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTablePipelineProgress_SdkV2) {
}

func (newState *SyncedTablePipelineProgress_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTablePipelineProgress_SdkV2) {
}

func (c SyncedTablePipelineProgress_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["estimated_completion_time_seconds"] = attrs["estimated_completion_time_seconds"].SetOptional()
	attrs["latest_version_currently_processing"] = attrs["latest_version_currently_processing"].SetOptional()
	attrs["sync_progress_completion"] = attrs["sync_progress_completion"].SetOptional()
	attrs["synced_row_count"] = attrs["synced_row_count"].SetOptional()
	attrs["total_row_count"] = attrs["total_row_count"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTablePipelineProgress.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTablePipelineProgress_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePipelineProgress_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTablePipelineProgress_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"estimated_completion_time_seconds":   o.EstimatedCompletionTimeSeconds,
			"latest_version_currently_processing": o.LatestVersionCurrentlyProcessing,
			"sync_progress_completion":            o.SyncProgressCompletion,
			"synced_row_count":                    o.SyncedRowCount,
			"total_row_count":                     o.TotalRowCount,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTablePipelineProgress_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"estimated_completion_time_seconds":   types.Float64Type,
			"latest_version_currently_processing": types.Int64Type,
			"sync_progress_completion":            types.Float64Type,
			"synced_row_count":                    types.Int64Type,
			"total_row_count":                     types.Int64Type,
		},
	}
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus_SdkV2 struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress types.List `tfsdk:"initial_pipeline_sync_progress"`
}

func (newState *SyncedTableProvisioningStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableProvisioningStatus_SdkV2) {
}

func (newState *SyncedTableProvisioningStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTableProvisioningStatus_SdkV2) {
}

func (c SyncedTableProvisioningStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetOptional()
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
func (a SyncedTableProvisioningStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableProvisioningStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTableProvisioningStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableProvisioningStatus_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SyncedTableProvisioningStatus_SdkV2) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress_SdkV2, bool) {
	var e SyncedTablePipelineProgress_SdkV2
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress_SdkV2
	d := o.InitialPipelineSyncProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus_SdkV2.
func (o *SyncedTableProvisioningStatus_SdkV2) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["initial_pipeline_sync_progress"]
	o.InitialPipelineSyncProgress = types.ListValueMust(t, vs)
}

// Specification of a synced database table.
type SyncedTableSpec_SdkV2 struct {
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist.
	CreateDatabaseObjectsIfMissing types.Bool `tfsdk:"create_database_objects_if_missing"`
	// Spec of new pipeline. Should be empty if pipeline_id is set
	NewPipelineSpec types.List `tfsdk:"new_pipeline_spec"`
	// ID of the associated pipeline. Should be empty if new_pipeline_spec is
	// set
	PipelineId types.String `tfsdk:"pipeline_id"`
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

func (newState *SyncedTableSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableSpec_SdkV2) {
}

func (newState *SyncedTableSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTableSpec_SdkV2) {
}

func (c SyncedTableSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetOptional()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetOptional()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
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
func (a SyncedTableSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_pipeline_spec":   reflect.TypeOf(NewPipelineSpec_SdkV2{}),
		"primary_key_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTableSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_objects_if_missing": o.CreateDatabaseObjectsIfMissing,
			"new_pipeline_spec":                  o.NewPipelineSpec,
			"pipeline_id":                        o.PipelineId,
			"primary_key_columns":                o.PrimaryKeyColumns,
			"scheduling_policy":                  o.SchedulingPolicy,
			"source_table_full_name":             o.SourceTableFullName,
			"timeseries_key":                     o.TimeseriesKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_objects_if_missing": types.BoolType,
			"new_pipeline_spec": basetypes.ListType{
				ElemType: NewPipelineSpec_SdkV2{}.Type(ctx),
			},
			"pipeline_id": types.StringType,
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
func (o *SyncedTableSpec_SdkV2) GetNewPipelineSpec(ctx context.Context) (NewPipelineSpec_SdkV2, bool) {
	var e NewPipelineSpec_SdkV2
	if o.NewPipelineSpec.IsNull() || o.NewPipelineSpec.IsUnknown() {
		return e, false
	}
	var v []NewPipelineSpec_SdkV2
	d := o.NewPipelineSpec.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewPipelineSpec sets the value of the NewPipelineSpec field in SyncedTableSpec_SdkV2.
func (o *SyncedTableSpec_SdkV2) SetNewPipelineSpec(ctx context.Context, v NewPipelineSpec_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["new_pipeline_spec"]
	o.NewPipelineSpec = types.ListValueMust(t, vs)
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in SyncedTableSpec_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableSpec_SdkV2) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
	if o.PrimaryKeyColumns.IsNull() || o.PrimaryKeyColumns.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.PrimaryKeyColumns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in SyncedTableSpec_SdkV2.
func (o *SyncedTableSpec_SdkV2) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeyColumns = types.ListValueMust(t, vs)
}

// Status of a synced table.
type SyncedTableStatus_SdkV2 struct {
	// Detailed status of a synced table. Shown if the synced table is in the
	// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus types.List `tfsdk:"continuous_update_status"`
	// The state of the synced table.
	DetailedState types.String `tfsdk:"detailed_state"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
	FailedStatus types.List `tfsdk:"failed_status"`
	// A text description of the current state of the synced table.
	Message types.String `tfsdk:"message"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus types.List `tfsdk:"provisioning_status"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus types.List `tfsdk:"triggered_update_status"`
}

func (newState *SyncedTableStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableStatus_SdkV2) {
}

func (newState *SyncedTableStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTableStatus_SdkV2) {
}

func (c SyncedTableStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["continuous_update_status"] = attrs["continuous_update_status"].SetOptional()
	attrs["continuous_update_status"] = attrs["continuous_update_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["detailed_state"] = attrs["detailed_state"].SetOptional()
	attrs["failed_status"] = attrs["failed_status"].SetOptional()
	attrs["failed_status"] = attrs["failed_status"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["message"] = attrs["message"].SetOptional()
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
func (a SyncedTableStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous_update_status": reflect.TypeOf(SyncedTableContinuousUpdateStatus_SdkV2{}),
		"failed_status":            reflect.TypeOf(SyncedTableFailedStatus_SdkV2{}),
		"provisioning_status":      reflect.TypeOf(SyncedTableProvisioningStatus_SdkV2{}),
		"triggered_update_status":  reflect.TypeOf(SyncedTableTriggeredUpdateStatus_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTableStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"continuous_update_status": o.ContinuousUpdateStatus,
			"detailed_state":           o.DetailedState,
			"failed_status":            o.FailedStatus,
			"message":                  o.Message,
			"provisioning_status":      o.ProvisioningStatus,
			"triggered_update_status":  o.TriggeredUpdateStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableStatus_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"continuous_update_status": basetypes.ListType{
				ElemType: SyncedTableContinuousUpdateStatus_SdkV2{}.Type(ctx),
			},
			"detailed_state": types.StringType,
			"failed_status": basetypes.ListType{
				ElemType: SyncedTableFailedStatus_SdkV2{}.Type(ctx),
			},
			"message": types.StringType,
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
func (o *SyncedTableStatus_SdkV2) GetContinuousUpdateStatus(ctx context.Context) (SyncedTableContinuousUpdateStatus_SdkV2, bool) {
	var e SyncedTableContinuousUpdateStatus_SdkV2
	if o.ContinuousUpdateStatus.IsNull() || o.ContinuousUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableContinuousUpdateStatus_SdkV2
	d := o.ContinuousUpdateStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContinuousUpdateStatus sets the value of the ContinuousUpdateStatus field in SyncedTableStatus_SdkV2.
func (o *SyncedTableStatus_SdkV2) SetContinuousUpdateStatus(ctx context.Context, v SyncedTableContinuousUpdateStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["continuous_update_status"]
	o.ContinuousUpdateStatus = types.ListValueMust(t, vs)
}

// GetFailedStatus returns the value of the FailedStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableFailedStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus_SdkV2) GetFailedStatus(ctx context.Context) (SyncedTableFailedStatus_SdkV2, bool) {
	var e SyncedTableFailedStatus_SdkV2
	if o.FailedStatus.IsNull() || o.FailedStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableFailedStatus_SdkV2
	d := o.FailedStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailedStatus sets the value of the FailedStatus field in SyncedTableStatus_SdkV2.
func (o *SyncedTableStatus_SdkV2) SetFailedStatus(ctx context.Context, v SyncedTableFailedStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["failed_status"]
	o.FailedStatus = types.ListValueMust(t, vs)
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableProvisioningStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus_SdkV2) GetProvisioningStatus(ctx context.Context) (SyncedTableProvisioningStatus_SdkV2, bool) {
	var e SyncedTableProvisioningStatus_SdkV2
	if o.ProvisioningStatus.IsNull() || o.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableProvisioningStatus_SdkV2
	d := o.ProvisioningStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvisioningStatus sets the value of the ProvisioningStatus field in SyncedTableStatus_SdkV2.
func (o *SyncedTableStatus_SdkV2) SetProvisioningStatus(ctx context.Context, v SyncedTableProvisioningStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["provisioning_status"]
	o.ProvisioningStatus = types.ListValueMust(t, vs)
}

// GetTriggeredUpdateStatus returns the value of the TriggeredUpdateStatus field in SyncedTableStatus_SdkV2 as
// a SyncedTableTriggeredUpdateStatus_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus_SdkV2) GetTriggeredUpdateStatus(ctx context.Context) (SyncedTableTriggeredUpdateStatus_SdkV2, bool) {
	var e SyncedTableTriggeredUpdateStatus_SdkV2
	if o.TriggeredUpdateStatus.IsNull() || o.TriggeredUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableTriggeredUpdateStatus_SdkV2
	d := o.TriggeredUpdateStatus.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in SyncedTableStatus_SdkV2.
func (o *SyncedTableStatus_SdkV2) SetTriggeredUpdateStatus(ctx context.Context, v SyncedTableTriggeredUpdateStatus_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["triggered_update_status"]
	o.TriggeredUpdateStatus = types.ListValueMust(t, vs)
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus_SdkV2 struct {
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may not be completely synced to the synced
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress types.List `tfsdk:"triggered_update_progress"`
}

func (newState *SyncedTableTriggeredUpdateStatus_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableTriggeredUpdateStatus_SdkV2) {
}

func (newState *SyncedTableTriggeredUpdateStatus_SdkV2) SyncEffectiveFieldsDuringRead(existingState SyncedTableTriggeredUpdateStatus_SdkV2) {
}

func (c SyncedTableTriggeredUpdateStatus_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()
	attrs["triggered_update_progress"] = attrs["triggered_update_progress"].SetOptional()
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
func (a SyncedTableTriggeredUpdateStatus_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"triggered_update_progress": reflect.TypeOf(SyncedTablePipelineProgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableTriggeredUpdateStatus_SdkV2
// only implements ToObjectValue() and Type().
func (o SyncedTableTriggeredUpdateStatus_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
			"triggered_update_progress":     o.TriggeredUpdateProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableTriggeredUpdateStatus_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *SyncedTableTriggeredUpdateStatus_SdkV2) GetTriggeredUpdateProgress(ctx context.Context) (SyncedTablePipelineProgress_SdkV2, bool) {
	var e SyncedTablePipelineProgress_SdkV2
	if o.TriggeredUpdateProgress.IsNull() || o.TriggeredUpdateProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress_SdkV2
	d := o.TriggeredUpdateProgress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus_SdkV2.
func (o *SyncedTableTriggeredUpdateStatus_SdkV2) SetTriggeredUpdateProgress(ctx context.Context, v SyncedTablePipelineProgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["triggered_update_progress"]
	o.TriggeredUpdateProgress = types.ListValueMust(t, vs)
}

// Update a Database Instance
type UpdateDatabaseInstanceRequest_SdkV2 struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance types.List `tfsdk:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"-"`
	// The list of fields to update.
	UpdateMask types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDatabaseInstanceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseInstanceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDatabaseInstanceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": o.DatabaseInstance,
			"name":              o.Name,
			"update_mask":       o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDatabaseInstanceRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDatabaseInstanceRequest_SdkV2) GetDatabaseInstance(ctx context.Context) (DatabaseInstance_SdkV2, bool) {
	var e DatabaseInstance_SdkV2
	if o.DatabaseInstance.IsNull() || o.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance_SdkV2
	d := o.DatabaseInstance.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest_SdkV2.
func (o *UpdateDatabaseInstanceRequest_SdkV2) SetDatabaseInstance(ctx context.Context, v DatabaseInstance_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance"]
	o.DatabaseInstance = types.ListValueMust(t, vs)
}
