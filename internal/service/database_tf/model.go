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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDatabaseCatalogRequest struct {
	Catalog types.Object `tfsdk:"catalog"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o CreateDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": o.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": DatabaseCatalog{}.Type(ctx),
		},
	}
}

// GetCatalog returns the value of the Catalog field in CreateDatabaseCatalogRequest as
// a DatabaseCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDatabaseCatalogRequest) GetCatalog(ctx context.Context) (DatabaseCatalog, bool) {
	var e DatabaseCatalog
	if o.Catalog.IsNull() || o.Catalog.IsUnknown() {
		return e, false
	}
	var v []DatabaseCatalog
	d := o.Catalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCatalog sets the value of the Catalog field in CreateDatabaseCatalogRequest.
func (o *CreateDatabaseCatalogRequest) SetCatalog(ctx context.Context, v DatabaseCatalog) {
	vs := v.ToObjectValue(ctx)
	o.Catalog = vs
}

type CreateDatabaseInstanceRequest struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance types.Object `tfsdk:"database_instance"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o CreateDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": o.DatabaseInstance,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": DatabaseInstance{}.Type(ctx),
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in CreateDatabaseInstanceRequest as
// a DatabaseInstance value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDatabaseInstanceRequest) GetDatabaseInstance(ctx context.Context) (DatabaseInstance, bool) {
	var e DatabaseInstance
	if o.DatabaseInstance.IsNull() || o.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance
	d := o.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in CreateDatabaseInstanceRequest.
func (o *CreateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstance = vs
}

type CreateDatabaseTableRequest struct {
	// Next field marker: 13
	Table types.Object `tfsdk:"table"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(DatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o CreateDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": o.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": DatabaseTable{}.Type(ctx),
		},
	}
}

// GetTable returns the value of the Table field in CreateDatabaseTableRequest as
// a DatabaseTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDatabaseTableRequest) GetTable(ctx context.Context) (DatabaseTable, bool) {
	var e DatabaseTable
	if o.Table.IsNull() || o.Table.IsUnknown() {
		return e, false
	}
	var v []DatabaseTable
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTable sets the value of the Table field in CreateDatabaseTableRequest.
func (o *CreateDatabaseTableRequest) SetTable(ctx context.Context, v DatabaseTable) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

type CreateSyncedDatabaseTableRequest struct {
	// Next field marker: 12
	SyncedTable types.Object `tfsdk:"synced_table"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o CreateSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"synced_table": o.SyncedTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"synced_table": SyncedDatabaseTable{}.Type(ctx),
		},
	}
}

// GetSyncedTable returns the value of the SyncedTable field in CreateSyncedDatabaseTableRequest as
// a SyncedDatabaseTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateSyncedDatabaseTableRequest) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable, bool) {
	var e SyncedDatabaseTable
	if o.SyncedTable.IsNull() || o.SyncedTable.IsUnknown() {
		return e, false
	}
	var v []SyncedDatabaseTable
	d := o.SyncedTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSyncedTable sets the value of the SyncedTable field in CreateSyncedDatabaseTableRequest.
func (o *CreateSyncedDatabaseTableRequest) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable) {
	vs := v.ToObjectValue(ctx)
	o.SyncedTable = vs
}

type DatabaseCatalog struct {
	CreateDatabaseIfNotExists types.Bool `tfsdk:"create_database_if_not_exists"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName types.String `tfsdk:"database_name"`
	// The name of the catalog in UC.
	Name types.String `tfsdk:"name"`

	Uid types.String `tfsdk:"uid"`
}

func (newState *DatabaseCatalog) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseCatalog) {
}

func (newState *DatabaseCatalog) SyncEffectiveFieldsDuringRead(existingState DatabaseCatalog) {
}

func (c DatabaseCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatabaseCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCatalog
// only implements ToObjectValue() and Type().
func (o DatabaseCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o DatabaseCatalog) Type(ctx context.Context) attr.Type {
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

type DatabaseCredential struct {
	ExpirationTime types.String `tfsdk:"expiration_time"`

	Token types.String `tfsdk:"token"`
}

func (newState *DatabaseCredential) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseCredential) {
}

func (newState *DatabaseCredential) SyncEffectiveFieldsDuringRead(existingState DatabaseCredential) {
}

func (c DatabaseCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatabaseCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCredential
// only implements ToObjectValue() and Type().
func (o DatabaseCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expiration_time": o.ExpirationTime,
			"token":           o.Token,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expiration_time": types.StringType,
			"token":           types.StringType,
		},
	}
}

// A DatabaseInstance represents a logical Postgres instance, comprised of both
// compute and storage.
type DatabaseInstance struct {
	// The sku of the instance. Valid values are "CU_1", "CU_2", "CU_4", "CU_8".
	Capacity types.String `tfsdk:"capacity"`
	// The timestamp when the instance was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the instance.
	Creator types.String `tfsdk:"creator"`
	// xref AIP-129. `stopped` is owned by the client, while `effective_stopped`
	// is owned by the server. `stopped` will only be set in Create/Update
	// response messages if and only if the user provides the field via the
	// request. `effective_stopped` on the other hand will always bet set in all
	// response messages (Create/Update/Get/List).
	EffectiveStopped types.Bool `tfsdk:"effective_stopped"`
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

func (newState *DatabaseInstance) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseInstance) {
}

func (newState *DatabaseInstance) SyncEffectiveFieldsDuringRead(existingState DatabaseInstance) {
}

func (c DatabaseInstance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["effective_stopped"] = attrs["effective_stopped"].SetComputed()
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
func (a DatabaseInstance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance
// only implements ToObjectValue() and Type().
func (o DatabaseInstance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"capacity":          o.Capacity,
			"creation_time":     o.CreationTime,
			"creator":           o.Creator,
			"effective_stopped": o.EffectiveStopped,
			"name":              o.Name,
			"pg_version":        o.PgVersion,
			"read_write_dns":    o.ReadWriteDns,
			"state":             o.State,
			"stopped":           o.Stopped,
			"uid":               o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstance) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"capacity":          types.StringType,
			"creation_time":     types.StringType,
			"creator":           types.StringType,
			"effective_stopped": types.BoolType,
			"name":              types.StringType,
			"pg_version":        types.StringType,
			"read_write_dns":    types.StringType,
			"state":             types.StringType,
			"stopped":           types.BoolType,
			"uid":               types.StringType,
		},
	}
}

// Next field marker: 13
type DatabaseTable struct {
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
}

func (newState *DatabaseTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabaseTable) {
}

func (newState *DatabaseTable) SyncEffectiveFieldsDuringRead(existingState DatabaseTable) {
}

func (c DatabaseTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatabaseTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseTable
// only implements ToObjectValue() and Type().
func (o DatabaseTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_name": o.DatabaseInstanceName,
			"logical_database_name":  o.LogicalDatabaseName,
			"name":                   o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_name": types.StringType,
			"logical_database_name":  types.StringType,
			"name":                   types.StringType,
		},
	}
}

type DeleteDatabaseCatalogRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseCatalogResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseCatalogResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseCatalogResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogResponse
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseCatalogResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseCatalogResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteDatabaseInstanceRequest struct {
	// By default, a instance cannot be deleted if it has descendant instances
	// created via PITR. If this flag is specified as true, all descendent
	// instances will be deleted as well.
	Force types.Bool `tfsdk:"-"`
	// Name of the instance to delete.
	Name types.String `tfsdk:"-"`
	// Note purge=false is in development. If false, the database instance is
	// soft deleted (implementation pending). Soft deleted instances behave as
	// if they are deleted, and cannot be used for CRUD operations nor connected
	// to. However they can be undeleted by calling the undelete API for a
	// limited time (implementation pending). If true, the database instance is
	// hard deleted and cannot be undeleted. For the time being, setting this
	// value to true is required to delete an instance (soft delete is not yet
	// supported).
	Purge types.Bool `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": o.Force,
			"name":  o.Name,
			"purge": o.Purge,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"force": types.BoolType,
			"name":  types.StringType,
			"purge": types.BoolType,
		},
	}
}

type DeleteDatabaseInstanceResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseInstanceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseInstanceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceResponse
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseTableResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseTableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDatabaseTableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseTableResponse
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseTableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseTableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteSyncedDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o DeleteSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteSyncedDatabaseTableResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSyncedDatabaseTableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSyncedDatabaseTableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableResponse
// only implements ToObjectValue() and Type().
func (o DeleteSyncedDatabaseTableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSyncedDatabaseTableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type FindDatabaseInstanceByUidRequest struct {
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
func (a FindDatabaseInstanceByUidRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FindDatabaseInstanceByUidRequest
// only implements ToObjectValue() and Type().
func (o FindDatabaseInstanceByUidRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"uid": o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FindDatabaseInstanceByUidRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"uid": types.StringType,
		},
	}
}

// Generates a credential that can be used to access database instances
type GenerateDatabaseCredentialRequest struct {
	// Instances to which the token will be scoped.
	InstanceNames types.List `tfsdk:"instance_names"`

	RequestId types.String `tfsdk:"request_id"`
}

func (newState *GenerateDatabaseCredentialRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenerateDatabaseCredentialRequest) {
}

func (newState *GenerateDatabaseCredentialRequest) SyncEffectiveFieldsDuringRead(existingState GenerateDatabaseCredentialRequest) {
}

func (c GenerateDatabaseCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenerateDatabaseCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"instance_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateDatabaseCredentialRequest
// only implements ToObjectValue() and Type().
func (o GenerateDatabaseCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_names": o.InstanceNames,
			"request_id":     o.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateDatabaseCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"request_id": types.StringType,
		},
	}
}

// GetInstanceNames returns the value of the InstanceNames field in GenerateDatabaseCredentialRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateDatabaseCredentialRequest) GetInstanceNames(ctx context.Context) ([]types.String, bool) {
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

// SetInstanceNames sets the value of the InstanceNames field in GenerateDatabaseCredentialRequest.
func (o *GenerateDatabaseCredentialRequest) SetInstanceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InstanceNames = types.ListValueMust(t, vs)
}

type GetDatabaseCatalogRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseCatalogRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o GetDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatabaseInstanceRequest struct {
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
func (a GetDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o GetDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o GetDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetSyncedDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSyncedDatabaseTableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o GetSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListDatabaseInstancesRequest struct {
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
func (a ListDatabaseInstancesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesRequest
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstancesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstancesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDatabaseInstancesResponse struct {
	// List of instances.
	DatabaseInstances types.List `tfsdk:"database_instances"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListDatabaseInstancesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDatabaseInstancesResponse) {
}

func (newState *ListDatabaseInstancesResponse) SyncEffectiveFieldsDuringRead(existingState ListDatabaseInstancesResponse) {
}

func (c ListDatabaseInstancesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDatabaseInstancesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesResponse
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstancesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instances": o.DatabaseInstances,
			"next_page_token":    o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstancesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instances": basetypes.ListType{
				ElemType: DatabaseInstance{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseInstances returns the value of the DatabaseInstances field in ListDatabaseInstancesResponse as
// a slice of DatabaseInstance values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDatabaseInstancesResponse) GetDatabaseInstances(ctx context.Context) ([]DatabaseInstance, bool) {
	if o.DatabaseInstances.IsNull() || o.DatabaseInstances.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstance
	d := o.DatabaseInstances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstances sets the value of the DatabaseInstances field in ListDatabaseInstancesResponse.
func (o *ListDatabaseInstancesResponse) SetDatabaseInstances(ctx context.Context, v []DatabaseInstance) {
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
type NewPipelineSpec struct {
	// UC catalog for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be a standard catalog where the user has
	// permissions to create Delta tables.
	StorageCatalog types.String `tfsdk:"storage_catalog"`
	// UC schema for the pipeline to store intermediate files (checkpoints,
	// event logs etc). This needs to be in the standard catalog where the user
	// has permissions to create Delta tables.
	StorageSchema types.String `tfsdk:"storage_schema"`
}

func (newState *NewPipelineSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan NewPipelineSpec) {
}

func (newState *NewPipelineSpec) SyncEffectiveFieldsDuringRead(existingState NewPipelineSpec) {
}

func (c NewPipelineSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NewPipelineSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NewPipelineSpec
// only implements ToObjectValue() and Type().
func (o NewPipelineSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"storage_catalog": o.StorageCatalog,
			"storage_schema":  o.StorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NewPipelineSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"storage_catalog": types.StringType,
			"storage_schema":  types.StringType,
		},
	}
}

// Next field marker: 12
type SyncedDatabaseTable struct {
	// Synced Table data synchronization status
	DataSynchronizationStatus types.Object `tfsdk:"data_synchronization_status"`
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
	Spec types.Object `tfsdk:"spec"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state"`
}

func (newState *SyncedDatabaseTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedDatabaseTable) {
}

func (newState *SyncedDatabaseTable) SyncEffectiveFieldsDuringRead(existingState SyncedDatabaseTable) {
}

func (c SyncedDatabaseTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].SetComputed()
	attrs["database_instance_name"] = attrs["database_instance_name"].SetOptional()
	attrs["logical_database_name"] = attrs["logical_database_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["spec"] = attrs["spec"].SetOptional()
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
func (a SyncedDatabaseTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_synchronization_status": reflect.TypeOf(SyncedTableStatus{}),
		"spec":                        reflect.TypeOf(SyncedTableSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedDatabaseTable
// only implements ToObjectValue() and Type().
func (o SyncedDatabaseTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"data_synchronization_status":      o.DataSynchronizationStatus,
			"database_instance_name":           o.DatabaseInstanceName,
			"logical_database_name":            o.LogicalDatabaseName,
			"name":                             o.Name,
			"spec":                             o.Spec,
			"unity_catalog_provisioning_state": o.UnityCatalogProvisioningState,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedDatabaseTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_synchronization_status":      SyncedTableStatus{}.Type(ctx),
			"database_instance_name":           types.StringType,
			"logical_database_name":            types.StringType,
			"name":                             types.StringType,
			"spec":                             SyncedTableSpec{}.Type(ctx),
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetDataSynchronizationStatus returns the value of the DataSynchronizationStatus field in SyncedDatabaseTable as
// a SyncedTableStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedDatabaseTable) GetDataSynchronizationStatus(ctx context.Context) (SyncedTableStatus, bool) {
	var e SyncedTableStatus
	if o.DataSynchronizationStatus.IsNull() || o.DataSynchronizationStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableStatus
	d := o.DataSynchronizationStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataSynchronizationStatus sets the value of the DataSynchronizationStatus field in SyncedDatabaseTable.
func (o *SyncedDatabaseTable) SetDataSynchronizationStatus(ctx context.Context, v SyncedTableStatus) {
	vs := v.ToObjectValue(ctx)
	o.DataSynchronizationStatus = vs
}

// GetSpec returns the value of the Spec field in SyncedDatabaseTable as
// a SyncedTableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedDatabaseTable) GetSpec(ctx context.Context) (SyncedTableSpec, bool) {
	var e SyncedTableSpec
	if o.Spec.IsNull() || o.Spec.IsUnknown() {
		return e, false
	}
	var v []SyncedTableSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSpec sets the value of the Spec field in SyncedDatabaseTable.
func (o *SyncedDatabaseTable) SetSpec(ctx context.Context, v SyncedTableSpec) {
	vs := v.ToObjectValue(ctx)
	o.Spec = vs
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
type SyncedTableContinuousUpdateStatus struct {
	// Progress of the initial data synchronization.
	InitialPipelineSyncProgress types.Object `tfsdk:"initial_pipeline_sync_progress"`
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may not be completely synced to the synced
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (newState *SyncedTableContinuousUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableContinuousUpdateStatus) {
}

func (newState *SyncedTableContinuousUpdateStatus) SyncEffectiveFieldsDuringRead(existingState SyncedTableContinuousUpdateStatus) {
}

func (c SyncedTableContinuousUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetOptional()
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
func (a SyncedTableContinuousUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableContinuousUpdateStatus
// only implements ToObjectValue() and Type().
func (o SyncedTableContinuousUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
			"last_processed_commit_version":  o.LastProcessedCommitVersion,
			"timestamp":                      o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableContinuousUpdateStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": SyncedTablePipelineProgress{}.Type(ctx),
			"last_processed_commit_version":  types.Int64Type,
			"timestamp":                      types.StringType,
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus as
// a SyncedTablePipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableContinuousUpdateStatus) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress, bool) {
	var e SyncedTablePipelineProgress
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress
	d := o.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus.
func (o *SyncedTableContinuousUpdateStatus) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.InitialPipelineSyncProgress = vs
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus struct {
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

func (newState *SyncedTableFailedStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableFailedStatus) {
}

func (newState *SyncedTableFailedStatus) SyncEffectiveFieldsDuringRead(existingState SyncedTableFailedStatus) {
}

func (c SyncedTableFailedStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SyncedTableFailedStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableFailedStatus
// only implements ToObjectValue() and Type().
func (o SyncedTableFailedStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableFailedStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
		},
	}
}

// Progress information of the Synced Table data synchronization pipeline.
type SyncedTablePipelineProgress struct {
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

func (newState *SyncedTablePipelineProgress) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTablePipelineProgress) {
}

func (newState *SyncedTablePipelineProgress) SyncEffectiveFieldsDuringRead(existingState SyncedTablePipelineProgress) {
}

func (c SyncedTablePipelineProgress) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SyncedTablePipelineProgress) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePipelineProgress
// only implements ToObjectValue() and Type().
func (o SyncedTablePipelineProgress) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o SyncedTablePipelineProgress) Type(ctx context.Context) attr.Type {
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
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress types.Object `tfsdk:"initial_pipeline_sync_progress"`
}

func (newState *SyncedTableProvisioningStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableProvisioningStatus) {
}

func (newState *SyncedTableProvisioningStatus) SyncEffectiveFieldsDuringRead(existingState SyncedTableProvisioningStatus) {
}

func (c SyncedTableProvisioningStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableProvisioningStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTableProvisioningStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableProvisioningStatus
// only implements ToObjectValue() and Type().
func (o SyncedTableProvisioningStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": o.InitialPipelineSyncProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableProvisioningStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": SyncedTablePipelineProgress{}.Type(ctx),
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus as
// a SyncedTablePipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableProvisioningStatus) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress, bool) {
	var e SyncedTablePipelineProgress
	if o.InitialPipelineSyncProgress.IsNull() || o.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress
	d := o.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus.
func (o *SyncedTableProvisioningStatus) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.InitialPipelineSyncProgress = vs
}

// Specification of a synced database table.
type SyncedTableSpec struct {
	// If true, the synced table's logical database and schema resources in PG
	// will be created if they do not already exist.
	CreateDatabaseObjectsIfMissing types.Bool `tfsdk:"create_database_objects_if_missing"`
	// User-specified ID of a pre-existing pipeline to bin pack. This field is
	// optional, and should be empty if new_pipeline_spec is set. This field
	// will only be set by the server in response messages if it is specified in
	// the request. The SyncedTableStatus message will always contain the
	// effective pipeline ID (either client provided or server generated),
	// however.
	ExistingPipelineId types.String `tfsdk:"existing_pipeline_id"`
	// Spec of new pipeline. Should be empty if pipeline_id /
	// existing_pipeline_id is set
	NewPipelineSpec types.Object `tfsdk:"new_pipeline_spec"`
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

func (newState *SyncedTableSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableSpec) {
}

func (newState *SyncedTableSpec) SyncEffectiveFieldsDuringRead(existingState SyncedTableSpec) {
}

func (c SyncedTableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetOptional()
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].SetOptional()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetOptional()
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
func (a SyncedTableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_pipeline_spec":   reflect.TypeOf(NewPipelineSpec{}),
		"primary_key_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableSpec
// only implements ToObjectValue() and Type().
func (o SyncedTableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_objects_if_missing": o.CreateDatabaseObjectsIfMissing,
			"existing_pipeline_id":               o.ExistingPipelineId,
			"new_pipeline_spec":                  o.NewPipelineSpec,
			"primary_key_columns":                o.PrimaryKeyColumns,
			"scheduling_policy":                  o.SchedulingPolicy,
			"source_table_full_name":             o.SourceTableFullName,
			"timeseries_key":                     o.TimeseriesKey,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_objects_if_missing": types.BoolType,
			"existing_pipeline_id":               types.StringType,
			"new_pipeline_spec":                  NewPipelineSpec{}.Type(ctx),
			"primary_key_columns": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scheduling_policy":      types.StringType,
			"source_table_full_name": types.StringType,
			"timeseries_key":         types.StringType,
		},
	}
}

// GetNewPipelineSpec returns the value of the NewPipelineSpec field in SyncedTableSpec as
// a NewPipelineSpec value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableSpec) GetNewPipelineSpec(ctx context.Context) (NewPipelineSpec, bool) {
	var e NewPipelineSpec
	if o.NewPipelineSpec.IsNull() || o.NewPipelineSpec.IsUnknown() {
		return e, false
	}
	var v []NewPipelineSpec
	d := o.NewPipelineSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNewPipelineSpec sets the value of the NewPipelineSpec field in SyncedTableSpec.
func (o *SyncedTableSpec) SetNewPipelineSpec(ctx context.Context, v NewPipelineSpec) {
	vs := v.ToObjectValue(ctx)
	o.NewPipelineSpec = vs
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in SyncedTableSpec as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableSpec) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
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

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in SyncedTableSpec.
func (o *SyncedTableSpec) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PrimaryKeyColumns = types.ListValueMust(t, vs)
}

// Status of a synced table.
type SyncedTableStatus struct {
	// Detailed status of a synced table. Shown if the synced table is in the
	// SYNCED_CONTINUOUS_UPDATE or the SYNCED_UPDATING_PIPELINE_RESOURCES state.
	ContinuousUpdateStatus types.Object `tfsdk:"continuous_update_status"`
	// The state of the synced table.
	DetailedState types.String `tfsdk:"detailed_state"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
	FailedStatus types.Object `tfsdk:"failed_status"`
	// A text description of the current state of the synced table.
	Message types.String `tfsdk:"message"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	PipelineId types.String `tfsdk:"pipeline_id"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT
	// state.
	ProvisioningStatus types.Object `tfsdk:"provisioning_status"`
	// Detailed status of a synced table. Shown if the synced table is in the
	// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
	TriggeredUpdateStatus types.Object `tfsdk:"triggered_update_status"`
}

func (newState *SyncedTableStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableStatus) {
}

func (newState *SyncedTableStatus) SyncEffectiveFieldsDuringRead(existingState SyncedTableStatus) {
}

func (c SyncedTableStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["continuous_update_status"] = attrs["continuous_update_status"].SetOptional()
	attrs["detailed_state"] = attrs["detailed_state"].SetOptional()
	attrs["failed_status"] = attrs["failed_status"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetOptional()
	attrs["provisioning_status"] = attrs["provisioning_status"].SetOptional()
	attrs["triggered_update_status"] = attrs["triggered_update_status"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTableStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"continuous_update_status": reflect.TypeOf(SyncedTableContinuousUpdateStatus{}),
		"failed_status":            reflect.TypeOf(SyncedTableFailedStatus{}),
		"provisioning_status":      reflect.TypeOf(SyncedTableProvisioningStatus{}),
		"triggered_update_status":  reflect.TypeOf(SyncedTableTriggeredUpdateStatus{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableStatus
// only implements ToObjectValue() and Type().
func (o SyncedTableStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"continuous_update_status": o.ContinuousUpdateStatus,
			"detailed_state":           o.DetailedState,
			"failed_status":            o.FailedStatus,
			"message":                  o.Message,
			"pipeline_id":              o.PipelineId,
			"provisioning_status":      o.ProvisioningStatus,
			"triggered_update_status":  o.TriggeredUpdateStatus,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"continuous_update_status": SyncedTableContinuousUpdateStatus{}.Type(ctx),
			"detailed_state":           types.StringType,
			"failed_status":            SyncedTableFailedStatus{}.Type(ctx),
			"message":                  types.StringType,
			"pipeline_id":              types.StringType,
			"provisioning_status":      SyncedTableProvisioningStatus{}.Type(ctx),
			"triggered_update_status":  SyncedTableTriggeredUpdateStatus{}.Type(ctx),
		},
	}
}

// GetContinuousUpdateStatus returns the value of the ContinuousUpdateStatus field in SyncedTableStatus as
// a SyncedTableContinuousUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus) GetContinuousUpdateStatus(ctx context.Context) (SyncedTableContinuousUpdateStatus, bool) {
	var e SyncedTableContinuousUpdateStatus
	if o.ContinuousUpdateStatus.IsNull() || o.ContinuousUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableContinuousUpdateStatus
	d := o.ContinuousUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetContinuousUpdateStatus sets the value of the ContinuousUpdateStatus field in SyncedTableStatus.
func (o *SyncedTableStatus) SetContinuousUpdateStatus(ctx context.Context, v SyncedTableContinuousUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	o.ContinuousUpdateStatus = vs
}

// GetFailedStatus returns the value of the FailedStatus field in SyncedTableStatus as
// a SyncedTableFailedStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus) GetFailedStatus(ctx context.Context) (SyncedTableFailedStatus, bool) {
	var e SyncedTableFailedStatus
	if o.FailedStatus.IsNull() || o.FailedStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableFailedStatus
	d := o.FailedStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFailedStatus sets the value of the FailedStatus field in SyncedTableStatus.
func (o *SyncedTableStatus) SetFailedStatus(ctx context.Context, v SyncedTableFailedStatus) {
	vs := v.ToObjectValue(ctx)
	o.FailedStatus = vs
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in SyncedTableStatus as
// a SyncedTableProvisioningStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus) GetProvisioningStatus(ctx context.Context) (SyncedTableProvisioningStatus, bool) {
	var e SyncedTableProvisioningStatus
	if o.ProvisioningStatus.IsNull() || o.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableProvisioningStatus
	d := o.ProvisioningStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetProvisioningStatus sets the value of the ProvisioningStatus field in SyncedTableStatus.
func (o *SyncedTableStatus) SetProvisioningStatus(ctx context.Context, v SyncedTableProvisioningStatus) {
	vs := v.ToObjectValue(ctx)
	o.ProvisioningStatus = vs
}

// GetTriggeredUpdateStatus returns the value of the TriggeredUpdateStatus field in SyncedTableStatus as
// a SyncedTableTriggeredUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus) GetTriggeredUpdateStatus(ctx context.Context) (SyncedTableTriggeredUpdateStatus, bool) {
	var e SyncedTableTriggeredUpdateStatus
	if o.TriggeredUpdateStatus.IsNull() || o.TriggeredUpdateStatus.IsUnknown() {
		return e, false
	}
	var v []SyncedTableTriggeredUpdateStatus
	d := o.TriggeredUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in SyncedTableStatus.
func (o *SyncedTableStatus) SetTriggeredUpdateStatus(ctx context.Context, v SyncedTableTriggeredUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	o.TriggeredUpdateStatus = vs
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus struct {
	// The last source table Delta version that was synced to the synced table.
	// Note that this Delta version may not be completely synced to the synced
	// table yet.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The timestamp of the last time any data was synchronized from the source
	// table to the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress types.Object `tfsdk:"triggered_update_progress"`
}

func (newState *SyncedTableTriggeredUpdateStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan SyncedTableTriggeredUpdateStatus) {
}

func (newState *SyncedTableTriggeredUpdateStatus) SyncEffectiveFieldsDuringRead(existingState SyncedTableTriggeredUpdateStatus) {
}

func (c SyncedTableTriggeredUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetOptional()
	attrs["timestamp"] = attrs["timestamp"].SetOptional()
	attrs["triggered_update_progress"] = attrs["triggered_update_progress"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SyncedTableTriggeredUpdateStatus.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SyncedTableTriggeredUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"triggered_update_progress": reflect.TypeOf(SyncedTablePipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableTriggeredUpdateStatus
// only implements ToObjectValue() and Type().
func (o SyncedTableTriggeredUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": o.LastProcessedCommitVersion,
			"timestamp":                     o.Timestamp,
			"triggered_update_progress":     o.TriggeredUpdateProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTableTriggeredUpdateStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"last_processed_commit_version": types.Int64Type,
			"timestamp":                     types.StringType,
			"triggered_update_progress":     SyncedTablePipelineProgress{}.Type(ctx),
		},
	}
}

// GetTriggeredUpdateProgress returns the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus as
// a SyncedTablePipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableTriggeredUpdateStatus) GetTriggeredUpdateProgress(ctx context.Context) (SyncedTablePipelineProgress, bool) {
	var e SyncedTablePipelineProgress
	if o.TriggeredUpdateProgress.IsNull() || o.TriggeredUpdateProgress.IsUnknown() {
		return e, false
	}
	var v []SyncedTablePipelineProgress
	d := o.TriggeredUpdateProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus.
func (o *SyncedTableTriggeredUpdateStatus) SetTriggeredUpdateProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.TriggeredUpdateProgress = vs
}

type UpdateDatabaseInstanceRequest struct {
	// A DatabaseInstance represents a logical Postgres instance, comprised of
	// both compute and storage.
	DatabaseInstance types.Object `tfsdk:"database_instance"`
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
func (a UpdateDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (o UpdateDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": o.DatabaseInstance,
			"name":              o.Name,
			"update_mask":       o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": DatabaseInstance{}.Type(ctx),
			"name":              types.StringType,
			"update_mask":       types.StringType,
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest as
// a DatabaseInstance value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateDatabaseInstanceRequest) GetDatabaseInstance(ctx context.Context) (DatabaseInstance, bool) {
	var e DatabaseInstance
	if o.DatabaseInstance.IsNull() || o.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v []DatabaseInstance
	d := o.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest.
func (o *UpdateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstance = vs
}
