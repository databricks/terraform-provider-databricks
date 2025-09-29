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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDatabaseCatalogRequest struct {
	Catalog types.Object `tfsdk:"catalog"`
}

func (to *CreateDatabaseCatalogRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseCatalogRequest) {
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

func (to *CreateDatabaseCatalogRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseCatalogRequest) {
	if !from.Catalog.IsNull() && !from.Catalog.IsUnknown() {
		if toCatalog, ok := to.GetCatalog(ctx); ok {
			if fromCatalog, ok := from.GetCatalog(ctx); ok {
				toCatalog.SyncFieldsDuringRead(ctx, fromCatalog)
				to.SetCatalog(ctx, toCatalog)
			}
		}
	}
}

func (c CreateDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog"] = attrs["catalog"].SetRequired()

	return attrs
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
	var v DatabaseCatalog
	d := o.Catalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalog sets the value of the Catalog field in CreateDatabaseCatalogRequest.
func (o *CreateDatabaseCatalogRequest) SetCatalog(ctx context.Context, v DatabaseCatalog) {
	vs := v.ToObjectValue(ctx)
	o.Catalog = vs
}

type CreateDatabaseInstanceRequest struct {
	// Instance to create.
	DatabaseInstance types.Object `tfsdk:"database_instance"`
}

func (to *CreateDatabaseInstanceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseInstanceRequest) {
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

func (to *CreateDatabaseInstanceRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseInstanceRequest) {
	if !from.DatabaseInstance.IsNull() && !from.DatabaseInstance.IsUnknown() {
		if toDatabaseInstance, ok := to.GetDatabaseInstance(ctx); ok {
			if fromDatabaseInstance, ok := from.GetDatabaseInstance(ctx); ok {
				toDatabaseInstance.SyncFieldsDuringRead(ctx, fromDatabaseInstance)
				to.SetDatabaseInstance(ctx, toDatabaseInstance)
			}
		}
	}
}

func (c CreateDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance"] = attrs["database_instance"].SetRequired()

	return attrs
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
	var v DatabaseInstance
	d := o.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in CreateDatabaseInstanceRequest.
func (o *CreateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstance = vs
}

type CreateDatabaseInstanceRoleRequest struct {
	DatabaseInstanceRole types.Object `tfsdk:"database_instance_role"`

	InstanceName types.String `tfsdk:"-"`
}

func (to *CreateDatabaseInstanceRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseInstanceRoleRequest) {
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

func (to *CreateDatabaseInstanceRoleRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseInstanceRoleRequest) {
	if !from.DatabaseInstanceRole.IsNull() && !from.DatabaseInstanceRole.IsUnknown() {
		if toDatabaseInstanceRole, ok := to.GetDatabaseInstanceRole(ctx); ok {
			if fromDatabaseInstanceRole, ok := from.GetDatabaseInstanceRole(ctx); ok {
				toDatabaseInstanceRole.SyncFieldsDuringRead(ctx, fromDatabaseInstanceRole)
				to.SetDatabaseInstanceRole(ctx, toDatabaseInstanceRole)
			}
		}
	}
}

func (c CreateDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_role"] = attrs["database_instance_role"].SetRequired()
	attrs["instance_name"] = attrs["instance_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseInstanceRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_role": reflect.TypeOf(DatabaseInstanceRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (o CreateDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_role": o.DatabaseInstanceRole,
			"instance_name":          o.InstanceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_role": DatabaseInstanceRole{}.Type(ctx),
			"instance_name":          types.StringType,
		},
	}
}

// GetDatabaseInstanceRole returns the value of the DatabaseInstanceRole field in CreateDatabaseInstanceRoleRequest as
// a DatabaseInstanceRole value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateDatabaseInstanceRoleRequest) GetDatabaseInstanceRole(ctx context.Context) (DatabaseInstanceRole, bool) {
	var e DatabaseInstanceRole
	if o.DatabaseInstanceRole.IsNull() || o.DatabaseInstanceRole.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRole
	d := o.DatabaseInstanceRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstanceRole sets the value of the DatabaseInstanceRole field in CreateDatabaseInstanceRoleRequest.
func (o *CreateDatabaseInstanceRoleRequest) SetDatabaseInstanceRole(ctx context.Context, v DatabaseInstanceRole) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstanceRole = vs
}

type CreateDatabaseTableRequest struct {
	Table types.Object `tfsdk:"table"`
}

func (to *CreateDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseTableRequest) {
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

func (to *CreateDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseTableRequest) {
	if !from.Table.IsNull() && !from.Table.IsUnknown() {
		if toTable, ok := to.GetTable(ctx); ok {
			if fromTable, ok := from.GetTable(ctx); ok {
				toTable.SyncFieldsDuringRead(ctx, fromTable)
				to.SetTable(ctx, toTable)
			}
		}
	}
}

func (c CreateDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["table"] = attrs["table"].SetRequired()

	return attrs
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
	var v DatabaseTable
	d := o.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTable sets the value of the Table field in CreateDatabaseTableRequest.
func (o *CreateDatabaseTableRequest) SetTable(ctx context.Context, v DatabaseTable) {
	vs := v.ToObjectValue(ctx)
	o.Table = vs
}

type CreateSyncedDatabaseTableRequest struct {
	SyncedTable types.Object `tfsdk:"synced_table"`
}

func (to *CreateSyncedDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateSyncedDatabaseTableRequest) {
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

func (to *CreateSyncedDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from CreateSyncedDatabaseTableRequest) {
	if !from.SyncedTable.IsNull() && !from.SyncedTable.IsUnknown() {
		if toSyncedTable, ok := to.GetSyncedTable(ctx); ok {
			if fromSyncedTable, ok := from.GetSyncedTable(ctx); ok {
				toSyncedTable.SyncFieldsDuringRead(ctx, fromSyncedTable)
				to.SetSyncedTable(ctx, toSyncedTable)
			}
		}
	}
}

func (c CreateSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["synced_table"] = attrs["synced_table"].SetRequired()

	return attrs
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
	var v SyncedDatabaseTable
	d := o.SyncedTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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

func (to *DatabaseCatalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCatalog) {
	if !from.CreateDatabaseIfNotExists.IsUnknown() && !from.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseIfNotExists = from.CreateDatabaseIfNotExists
	}
}

func (to *DatabaseCatalog) SyncFieldsDuringRead(ctx context.Context, from DatabaseCatalog) {
	if !from.CreateDatabaseIfNotExists.IsUnknown() && !from.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the prior state.
		to.CreateDatabaseIfNotExists = from.CreateDatabaseIfNotExists
	}
}

func (c DatabaseCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *DatabaseCredential) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseCredential) {
}

func (to *DatabaseCredential) SyncFieldsDuringRead(ctx context.Context, from DatabaseCredential) {
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
	// The refs of the child instances. This is only available if the instance
	// is parent instance.
	ChildInstanceRefs types.List `tfsdk:"child_instance_refs"`
	// The timestamp when the instance was created.
	CreationTime types.String `tfsdk:"creation_time"`
	// The email of the creator of the instance.
	Creator types.String `tfsdk:"creator"`
	// Deprecated. The sku of the instance; this field will always match the
	// value of capacity.
	EffectiveCapacity types.String `tfsdk:"effective_capacity"`
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
	ParentInstanceRef types.Object `tfsdk:"parent_instance_ref"`
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
}

func (to *DatabaseInstance) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstance) {
	if !from.ChildInstanceRefs.IsNull() && !from.ChildInstanceRefs.IsUnknown() && to.ChildInstanceRefs.IsNull() && len(from.ChildInstanceRefs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ChildInstanceRefs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ChildInstanceRefs = from.ChildInstanceRefs
	}
	if !from.EnablePgNativeLogin.IsUnknown() && !from.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnablePgNativeLogin = from.EnablePgNativeLogin
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
}

func (to *DatabaseInstance) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstance) {
	if !from.ChildInstanceRefs.IsNull() && !from.ChildInstanceRefs.IsUnknown() && to.ChildInstanceRefs.IsNull() && len(from.ChildInstanceRefs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ChildInstanceRefs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ChildInstanceRefs = from.ChildInstanceRefs
	}
	if !from.EnablePgNativeLogin.IsUnknown() && !from.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the prior state.
		to.EnablePgNativeLogin = from.EnablePgNativeLogin
	}
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				toParentInstanceRef.SyncFieldsDuringRead(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
}

func (c DatabaseInstance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["child_instance_refs"] = attrs["child_instance_refs"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
	attrs["effective_capacity"] = attrs["effective_capacity"].SetComputed()
	attrs["effective_enable_pg_native_login"] = attrs["effective_enable_pg_native_login"].SetComputed()
	attrs["effective_enable_readable_secondaries"] = attrs["effective_enable_readable_secondaries"].SetComputed()
	attrs["effective_node_count"] = attrs["effective_node_count"].SetComputed()
	attrs["effective_retention_window_in_days"] = attrs["effective_retention_window_in_days"].SetComputed()
	attrs["effective_stopped"] = attrs["effective_stopped"].SetComputed()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].SetOptional()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].SetComputed()
	attrs["enable_pg_native_login"] = attrs["enable_pg_native_login"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["enable_readable_secondaries"] = attrs["enable_readable_secondaries"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["node_count"] = attrs["node_count"].SetOptional()
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].SetOptional()
	attrs["parent_instance_ref"] = attrs["parent_instance_ref"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["pg_version"] = attrs["pg_version"].SetComputed()
	attrs["read_only_dns"] = attrs["read_only_dns"].SetComputed()
	attrs["read_write_dns"] = attrs["read_write_dns"].SetComputed()
	attrs["retention_window_in_days"] = attrs["retention_window_in_days"].SetOptional()
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
	return map[string]reflect.Type{
		"child_instance_refs": reflect.TypeOf(DatabaseInstanceRef{}),
		"parent_instance_ref": reflect.TypeOf(DatabaseInstanceRef{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance
// only implements ToObjectValue() and Type().
func (o DatabaseInstance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"capacity":                              o.Capacity,
			"child_instance_refs":                   o.ChildInstanceRefs,
			"creation_time":                         o.CreationTime,
			"creator":                               o.Creator,
			"effective_capacity":                    o.EffectiveCapacity,
			"effective_enable_pg_native_login":      o.EffectiveEnablePgNativeLogin,
			"effective_enable_readable_secondaries": o.EffectiveEnableReadableSecondaries,
			"effective_node_count":                  o.EffectiveNodeCount,
			"effective_retention_window_in_days":    o.EffectiveRetentionWindowInDays,
			"effective_stopped":                     o.EffectiveStopped,
			"enable_pg_native_login":                o.EnablePgNativeLogin,
			"enable_readable_secondaries":           o.EnableReadableSecondaries,
			"name":                                  o.Name,
			"node_count":                            o.NodeCount,
			"parent_instance_ref":                   o.ParentInstanceRef,
			"pg_version":                            o.PgVersion,
			"read_only_dns":                         o.ReadOnlyDns,
			"read_write_dns":                        o.ReadWriteDns,
			"retention_window_in_days":              o.RetentionWindowInDays,
			"state":                                 o.State,
			"stopped":                               o.Stopped,
			"uid":                                   o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstance) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"capacity": types.StringType,
			"child_instance_refs": basetypes.ListType{
				ElemType: DatabaseInstanceRef{}.Type(ctx),
			},
			"creation_time":                         types.StringType,
			"creator":                               types.StringType,
			"effective_capacity":                    types.StringType,
			"effective_enable_pg_native_login":      types.BoolType,
			"effective_enable_readable_secondaries": types.BoolType,
			"effective_node_count":                  types.Int64Type,
			"effective_retention_window_in_days":    types.Int64Type,
			"effective_stopped":                     types.BoolType,
			"enable_pg_native_login":                types.BoolType,
			"enable_readable_secondaries":           types.BoolType,
			"name":                                  types.StringType,
			"node_count":                            types.Int64Type,
			"parent_instance_ref":                   DatabaseInstanceRef{}.Type(ctx),
			"pg_version":                            types.StringType,
			"read_only_dns":                         types.StringType,
			"read_write_dns":                        types.StringType,
			"retention_window_in_days":              types.Int64Type,
			"state":                                 types.StringType,
			"stopped":                               types.BoolType,
			"uid":                                   types.StringType,
		},
	}
}

// GetChildInstanceRefs returns the value of the ChildInstanceRefs field in DatabaseInstance as
// a slice of DatabaseInstanceRef values.
// If the field is unknown or null, the boolean return value is false.
func (o *DatabaseInstance) GetChildInstanceRefs(ctx context.Context) ([]DatabaseInstanceRef, bool) {
	if o.ChildInstanceRefs.IsNull() || o.ChildInstanceRefs.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstanceRef
	d := o.ChildInstanceRefs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChildInstanceRefs sets the value of the ChildInstanceRefs field in DatabaseInstance.
func (o *DatabaseInstance) SetChildInstanceRefs(ctx context.Context, v []DatabaseInstanceRef) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["child_instance_refs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ChildInstanceRefs = types.ListValueMust(t, vs)
}

// GetParentInstanceRef returns the value of the ParentInstanceRef field in DatabaseInstance as
// a DatabaseInstanceRef value.
// If the field is unknown or null, the boolean return value is false.
func (o *DatabaseInstance) GetParentInstanceRef(ctx context.Context) (DatabaseInstanceRef, bool) {
	var e DatabaseInstanceRef
	if o.ParentInstanceRef.IsNull() || o.ParentInstanceRef.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRef
	d := o.ParentInstanceRef.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParentInstanceRef sets the value of the ParentInstanceRef field in DatabaseInstance.
func (o *DatabaseInstance) SetParentInstanceRef(ctx context.Context, v DatabaseInstanceRef) {
	vs := v.ToObjectValue(ctx)
	o.ParentInstanceRef = vs
}

// DatabaseInstanceRef is a reference to a database instance. It is used in the
// DatabaseInstance object to refer to the parent instance of an instance and to
// refer the child instances of an instance. To specify as a parent instance
// during creation of an instance, the lsn and branch_time fields are optional.
// If not specified, the child instance will be created from the latest lsn of
// the parent. If both lsn and branch_time are specified, the lsn will be used
// to create the child instance.
type DatabaseInstanceRef struct {
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

func (to *DatabaseInstanceRef) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstanceRef) {
}

func (to *DatabaseInstanceRef) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstanceRef) {
}

func (c DatabaseInstanceRef) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_time"] = attrs["branch_time"].SetOptional()
	attrs["effective_lsn"] = attrs["effective_lsn"].SetComputed()
	attrs["lsn"] = attrs["lsn"].SetOptional()
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
func (a DatabaseInstanceRef) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRef
// only implements ToObjectValue() and Type().
func (o DatabaseInstanceRef) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_time":   o.BranchTime,
			"effective_lsn": o.EffectiveLsn,
			"lsn":           o.Lsn,
			"name":          o.Name,
			"uid":           o.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstanceRef) Type(ctx context.Context) attr.Type {
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
type DatabaseInstanceRole struct {
	// API-exposed Postgres role attributes
	Attributes types.Object `tfsdk:"attributes"`
	// The type of the role.
	IdentityType types.String `tfsdk:"identity_type"`
	// An enum value for a standard role that this role is a member of.
	MembershipRole types.String `tfsdk:"membership_role"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	Name types.String `tfsdk:"name"`
}

func (to *DatabaseInstanceRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstanceRole) {
	if !from.Attributes.IsNull() && !from.Attributes.IsUnknown() {
		if toAttributes, ok := to.GetAttributes(ctx); ok {
			if fromAttributes, ok := from.GetAttributes(ctx); ok {
				// Recursively sync the fields of Attributes
				toAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromAttributes)
				to.SetAttributes(ctx, toAttributes)
			}
		}
	}
}

func (to *DatabaseInstanceRole) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstanceRole) {
	if !from.Attributes.IsNull() && !from.Attributes.IsUnknown() {
		if toAttributes, ok := to.GetAttributes(ctx); ok {
			if fromAttributes, ok := from.GetAttributes(ctx); ok {
				toAttributes.SyncFieldsDuringRead(ctx, fromAttributes)
				to.SetAttributes(ctx, toAttributes)
			}
		}
	}
}

func (c DatabaseInstanceRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["attributes"] = attrs["attributes"].SetOptional()
	attrs["identity_type"] = attrs["identity_type"].SetOptional()
	attrs["membership_role"] = attrs["membership_role"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseInstanceRole.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabaseInstanceRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes": reflect.TypeOf(DatabaseInstanceRoleAttributes{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRole
// only implements ToObjectValue() and Type().
func (o DatabaseInstanceRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":      o.Attributes,
			"identity_type":   o.IdentityType,
			"membership_role": o.MembershipRole,
			"name":            o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstanceRole) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":      DatabaseInstanceRoleAttributes{}.Type(ctx),
			"identity_type":   types.StringType,
			"membership_role": types.StringType,
			"name":            types.StringType,
		},
	}
}

// GetAttributes returns the value of the Attributes field in DatabaseInstanceRole as
// a DatabaseInstanceRoleAttributes value.
// If the field is unknown or null, the boolean return value is false.
func (o *DatabaseInstanceRole) GetAttributes(ctx context.Context) (DatabaseInstanceRoleAttributes, bool) {
	var e DatabaseInstanceRoleAttributes
	if o.Attributes.IsNull() || o.Attributes.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRoleAttributes
	d := o.Attributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAttributes sets the value of the Attributes field in DatabaseInstanceRole.
func (o *DatabaseInstanceRole) SetAttributes(ctx context.Context, v DatabaseInstanceRoleAttributes) {
	vs := v.ToObjectValue(ctx)
	o.Attributes = vs
}

// Attributes that can be granted to a Postgres role. We are only implementing a
// subset for now, see xref:
// https://www.postgresql.org/docs/16/sql-createrole.html The values follow
// Postgres keyword naming e.g. CREATEDB, BYPASSRLS, etc. which is why they
// don't include typical underscores between words. We were requested to make
// this a nested object/struct representation since these are knobs from an
// external spec.
type DatabaseInstanceRoleAttributes struct {
	Bypassrls types.Bool `tfsdk:"bypassrls"`

	Createdb types.Bool `tfsdk:"createdb"`

	Createrole types.Bool `tfsdk:"createrole"`
}

func (to *DatabaseInstanceRoleAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstanceRoleAttributes) {
}

func (to *DatabaseInstanceRoleAttributes) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstanceRoleAttributes) {
}

func (c DatabaseInstanceRoleAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatabaseInstanceRoleAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRoleAttributes
// only implements ToObjectValue() and Type().
func (o DatabaseInstanceRoleAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bypassrls":  o.Bypassrls,
			"createdb":   o.Createdb,
			"createrole": o.Createrole,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabaseInstanceRoleAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bypassrls":  types.BoolType,
			"createdb":   types.BoolType,
			"createrole": types.BoolType,
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

func (to *DatabaseTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseTable) {
}

func (to *DatabaseTable) SyncFieldsDuringRead(ctx context.Context, from DatabaseTable) {
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

func (to *DeleteDatabaseCatalogRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseCatalogRequest) {
}

func (to *DeleteDatabaseCatalogRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseCatalogRequest) {
}

func (c DeleteDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeleteDatabaseInstanceRequest struct {
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

func (to *DeleteDatabaseInstanceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseInstanceRequest) {
}

func (to *DeleteDatabaseInstanceRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseInstanceRequest) {
}

func (c DeleteDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing types.Bool `tfsdk:"-"`

	InstanceName types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`

	ReassignOwnedTo types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseInstanceRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseInstanceRoleRequest) {
}

func (to *DeleteDatabaseInstanceRoleRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseInstanceRoleRequest) {
}

func (c DeleteDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (o DeleteDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing":     o.AllowMissing,
			"instance_name":     o.InstanceName,
			"name":              o.Name,
			"reassign_owned_to": o.ReassignOwnedTo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing":     types.BoolType,
			"instance_name":     types.StringType,
			"name":              types.StringType,
			"reassign_owned_to": types.StringType,
		},
	}
}

type DeleteDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseTableRequest) {
}

func (to *DeleteDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseTableRequest) {
}

func (c DeleteDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeleteSyncedDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteSyncedDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSyncedDatabaseTableRequest) {
}

func (to *DeleteSyncedDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteSyncedDatabaseTableRequest) {
}

func (c DeleteSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type DeltaTableSyncInfo struct {
	// The timestamp when the above Delta version was committed in the source
	// Delta table. Note: This is the Delta commit time, not the time the data
	// was written to the synced table.
	DeltaCommitTimestamp types.String `tfsdk:"delta_commit_timestamp"`
	// The Delta Lake commit version that was last successfully synced.
	DeltaCommitVersion types.Int64 `tfsdk:"delta_commit_version"`
}

func (to *DeltaTableSyncInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeltaTableSyncInfo) {
}

func (to *DeltaTableSyncInfo) SyncFieldsDuringRead(ctx context.Context, from DeltaTableSyncInfo) {
}

func (c DeltaTableSyncInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeltaTableSyncInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaTableSyncInfo
// only implements ToObjectValue() and Type().
func (o DeltaTableSyncInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_commit_timestamp": o.DeltaCommitTimestamp,
			"delta_commit_version":   o.DeltaCommitVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeltaTableSyncInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_commit_timestamp": types.StringType,
			"delta_commit_version":   types.Int64Type,
		},
	}
}

type FindDatabaseInstanceByUidRequest struct {
	// UID of the cluster to get.
	Uid types.String `tfsdk:"-"`
}

func (to *FindDatabaseInstanceByUidRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FindDatabaseInstanceByUidRequest) {
}

func (to *FindDatabaseInstanceByUidRequest) SyncFieldsDuringRead(ctx context.Context, from FindDatabaseInstanceByUidRequest) {
}

func (c FindDatabaseInstanceByUidRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	// The returned token will be scoped to the union of instance_names and
	// instances containing the specified UC tables, so instance_names is
	// allowed to be empty.
	Claims types.List `tfsdk:"claims"`
	// Instances to which the token will be scoped.
	InstanceNames types.List `tfsdk:"instance_names"`

	RequestId types.String `tfsdk:"request_id"`
}

func (to *GenerateDatabaseCredentialRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenerateDatabaseCredentialRequest) {
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

func (to *GenerateDatabaseCredentialRequest) SyncFieldsDuringRead(ctx context.Context, from GenerateDatabaseCredentialRequest) {
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

func (c GenerateDatabaseCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenerateDatabaseCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"claims":         reflect.TypeOf(RequestedClaims{}),
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
			"claims":         o.Claims,
			"instance_names": o.InstanceNames,
			"request_id":     o.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenerateDatabaseCredentialRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"claims": basetypes.ListType{
				ElemType: RequestedClaims{}.Type(ctx),
			},
			"instance_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"request_id": types.StringType,
		},
	}
}

// GetClaims returns the value of the Claims field in GenerateDatabaseCredentialRequest as
// a slice of RequestedClaims values.
// If the field is unknown or null, the boolean return value is false.
func (o *GenerateDatabaseCredentialRequest) GetClaims(ctx context.Context) ([]RequestedClaims, bool) {
	if o.Claims.IsNull() || o.Claims.IsUnknown() {
		return nil, false
	}
	var v []RequestedClaims
	d := o.Claims.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClaims sets the value of the Claims field in GenerateDatabaseCredentialRequest.
func (o *GenerateDatabaseCredentialRequest) SetClaims(ctx context.Context, v []RequestedClaims) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["claims"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Claims = types.ListValueMust(t, vs)
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

func (to *GetDatabaseCatalogRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseCatalogRequest) {
}

func (to *GetDatabaseCatalogRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseCatalogRequest) {
}

func (c GetDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *GetDatabaseInstanceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseInstanceRequest) {
}

func (to *GetDatabaseInstanceRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseInstanceRequest) {
}

func (c GetDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type GetDatabaseInstanceRoleRequest struct {
	InstanceName types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseInstanceRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseInstanceRoleRequest) {
}

func (to *GetDatabaseInstanceRoleRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseInstanceRoleRequest) {
}

func (c GetDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (o GetDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": o.InstanceName,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"name":          types.StringType,
		},
	}
}

type GetDatabaseTableRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseTableRequest) {
}

func (to *GetDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseTableRequest) {
}

func (c GetDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *GetSyncedDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSyncedDatabaseTableRequest) {
}

func (to *GetSyncedDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from GetSyncedDatabaseTableRequest) {
}

func (c GetSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

type ListDatabaseCatalogsRequest struct {
	// Name of the instance to get database catalogs for.
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseCatalogsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseCatalogsRequest) {
}

func (to *ListDatabaseCatalogsRequest) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseCatalogsRequest) {
}

func (c ListDatabaseCatalogsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDatabaseCatalogsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseCatalogsRequest
// only implements ToObjectValue() and Type().
func (o ListDatabaseCatalogsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": o.InstanceName,
			"page_size":     o.PageSize,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseCatalogsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"page_size":     types.Int64Type,
			"page_token":    types.StringType,
		},
	}
}

type ListDatabaseCatalogsResponse struct {
	DatabaseCatalogs types.List `tfsdk:"database_catalogs"`
	// Pagination token to request the next page of database catalogs.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseCatalogsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseCatalogsResponse) {
	if !from.DatabaseCatalogs.IsNull() && !from.DatabaseCatalogs.IsUnknown() && to.DatabaseCatalogs.IsNull() && len(from.DatabaseCatalogs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseCatalogs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseCatalogs = from.DatabaseCatalogs
	}
}

func (to *ListDatabaseCatalogsResponse) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseCatalogsResponse) {
	if !from.DatabaseCatalogs.IsNull() && !from.DatabaseCatalogs.IsUnknown() && to.DatabaseCatalogs.IsNull() && len(from.DatabaseCatalogs.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseCatalogs, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseCatalogs = from.DatabaseCatalogs
	}
}

func (c ListDatabaseCatalogsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDatabaseCatalogsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalogs": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseCatalogsResponse
// only implements ToObjectValue() and Type().
func (o ListDatabaseCatalogsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_catalogs": o.DatabaseCatalogs,
			"next_page_token":   o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseCatalogsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_catalogs": basetypes.ListType{
				ElemType: DatabaseCatalog{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseCatalogs returns the value of the DatabaseCatalogs field in ListDatabaseCatalogsResponse as
// a slice of DatabaseCatalog values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDatabaseCatalogsResponse) GetDatabaseCatalogs(ctx context.Context) ([]DatabaseCatalog, bool) {
	if o.DatabaseCatalogs.IsNull() || o.DatabaseCatalogs.IsUnknown() {
		return nil, false
	}
	var v []DatabaseCatalog
	d := o.DatabaseCatalogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseCatalogs sets the value of the DatabaseCatalogs field in ListDatabaseCatalogsResponse.
func (o *ListDatabaseCatalogsResponse) SetDatabaseCatalogs(ctx context.Context, v []DatabaseCatalog) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database_catalogs"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DatabaseCatalogs = types.ListValueMust(t, vs)
}

type ListDatabaseInstanceRolesRequest struct {
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseInstanceRolesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstanceRolesRequest) {
}

func (to *ListDatabaseInstanceRolesRequest) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstanceRolesRequest) {
}

func (c ListDatabaseInstanceRolesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListDatabaseInstanceRolesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstanceRolesRequest
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstanceRolesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": o.InstanceName,
			"page_size":     o.PageSize,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstanceRolesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"page_size":     types.Int64Type,
			"page_token":    types.StringType,
		},
	}
}

type ListDatabaseInstanceRolesResponse struct {
	// List of database instance roles.
	DatabaseInstanceRoles types.List `tfsdk:"database_instance_roles"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseInstanceRolesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstanceRolesResponse) {
	if !from.DatabaseInstanceRoles.IsNull() && !from.DatabaseInstanceRoles.IsUnknown() && to.DatabaseInstanceRoles.IsNull() && len(from.DatabaseInstanceRoles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstanceRoles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstanceRoles = from.DatabaseInstanceRoles
	}
}

func (to *ListDatabaseInstanceRolesResponse) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstanceRolesResponse) {
	if !from.DatabaseInstanceRoles.IsNull() && !from.DatabaseInstanceRoles.IsUnknown() && to.DatabaseInstanceRoles.IsNull() && len(from.DatabaseInstanceRoles.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstanceRoles, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstanceRoles = from.DatabaseInstanceRoles
	}
}

func (c ListDatabaseInstanceRolesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_roles"] = attrs["database_instance_roles"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseInstanceRolesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDatabaseInstanceRolesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_roles": reflect.TypeOf(DatabaseInstanceRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstanceRolesResponse
// only implements ToObjectValue() and Type().
func (o ListDatabaseInstanceRolesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_roles": o.DatabaseInstanceRoles,
			"next_page_token":         o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListDatabaseInstanceRolesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_roles": basetypes.ListType{
				ElemType: DatabaseInstanceRole{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseInstanceRoles returns the value of the DatabaseInstanceRoles field in ListDatabaseInstanceRolesResponse as
// a slice of DatabaseInstanceRole values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListDatabaseInstanceRolesResponse) GetDatabaseInstanceRoles(ctx context.Context) ([]DatabaseInstanceRole, bool) {
	if o.DatabaseInstanceRoles.IsNull() || o.DatabaseInstanceRoles.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstanceRole
	d := o.DatabaseInstanceRoles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstanceRoles sets the value of the DatabaseInstanceRoles field in ListDatabaseInstanceRolesResponse.
func (o *ListDatabaseInstanceRolesResponse) SetDatabaseInstanceRoles(ctx context.Context, v []DatabaseInstanceRole) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance_roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DatabaseInstanceRoles = types.ListValueMust(t, vs)
}

type ListDatabaseInstancesRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Instances. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseInstancesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstancesRequest) {
}

func (to *ListDatabaseInstancesRequest) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstancesRequest) {
}

func (c ListDatabaseInstancesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *ListDatabaseInstancesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseInstancesResponse) {
	if !from.DatabaseInstances.IsNull() && !from.DatabaseInstances.IsUnknown() && to.DatabaseInstances.IsNull() && len(from.DatabaseInstances.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstances, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstances = from.DatabaseInstances
	}
}

func (to *ListDatabaseInstancesResponse) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseInstancesResponse) {
	if !from.DatabaseInstances.IsNull() && !from.DatabaseInstances.IsUnknown() && to.DatabaseInstances.IsNull() && len(from.DatabaseInstances.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseInstances, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseInstances = from.DatabaseInstances
	}
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

type ListSyncedDatabaseTablesRequest struct {
	// Name of the instance to get synced tables for.
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListSyncedDatabaseTablesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSyncedDatabaseTablesRequest) {
}

func (to *ListSyncedDatabaseTablesRequest) SyncFieldsDuringRead(ctx context.Context, from ListSyncedDatabaseTablesRequest) {
}

func (c ListSyncedDatabaseTablesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListSyncedDatabaseTablesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSyncedDatabaseTablesRequest
// only implements ToObjectValue() and Type().
func (o ListSyncedDatabaseTablesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": o.InstanceName,
			"page_size":     o.PageSize,
			"page_token":    o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSyncedDatabaseTablesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"page_size":     types.Int64Type,
			"page_token":    types.StringType,
		},
	}
}

type ListSyncedDatabaseTablesResponse struct {
	// Pagination token to request the next page of synced tables.
	NextPageToken types.String `tfsdk:"next_page_token"`

	SyncedTables types.List `tfsdk:"synced_tables"`
}

func (to *ListSyncedDatabaseTablesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListSyncedDatabaseTablesResponse) {
	if !from.SyncedTables.IsNull() && !from.SyncedTables.IsUnknown() && to.SyncedTables.IsNull() && len(from.SyncedTables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SyncedTables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SyncedTables = from.SyncedTables
	}
}

func (to *ListSyncedDatabaseTablesResponse) SyncFieldsDuringRead(ctx context.Context, from ListSyncedDatabaseTablesResponse) {
	if !from.SyncedTables.IsNull() && !from.SyncedTables.IsUnknown() && to.SyncedTables.IsNull() && len(from.SyncedTables.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for SyncedTables, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.SyncedTables = from.SyncedTables
	}
}

func (c ListSyncedDatabaseTablesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListSyncedDatabaseTablesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_tables": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSyncedDatabaseTablesResponse
// only implements ToObjectValue() and Type().
func (o ListSyncedDatabaseTablesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"synced_tables":   o.SyncedTables,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListSyncedDatabaseTablesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"synced_tables": basetypes.ListType{
				ElemType: SyncedDatabaseTable{}.Type(ctx),
			},
		},
	}
}

// GetSyncedTables returns the value of the SyncedTables field in ListSyncedDatabaseTablesResponse as
// a slice of SyncedDatabaseTable values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListSyncedDatabaseTablesResponse) GetSyncedTables(ctx context.Context) ([]SyncedDatabaseTable, bool) {
	if o.SyncedTables.IsNull() || o.SyncedTables.IsUnknown() {
		return nil, false
	}
	var v []SyncedDatabaseTable
	d := o.SyncedTables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSyncedTables sets the value of the SyncedTables field in ListSyncedDatabaseTablesResponse.
func (o *ListSyncedDatabaseTablesResponse) SetSyncedTables(ctx context.Context, v []SyncedDatabaseTable) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["synced_tables"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.SyncedTables = types.ListValueMust(t, vs)
}

// Custom fields that user can set for pipeline while creating
// SyncedDatabaseTable. Note that other fields of pipeline are still inferred by
// table def internally
type NewPipelineSpec struct {
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

func (to *NewPipelineSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NewPipelineSpec) {
}

func (to *NewPipelineSpec) SyncFieldsDuringRead(ctx context.Context, from NewPipelineSpec) {
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

type RequestedClaims struct {
	PermissionSet types.String `tfsdk:"permission_set"`

	Resources types.List `tfsdk:"resources"`
}

func (to *RequestedClaims) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedClaims) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (to *RequestedClaims) SyncFieldsDuringRead(ctx context.Context, from RequestedClaims) {
	if !from.Resources.IsNull() && !from.Resources.IsUnknown() && to.Resources.IsNull() && len(from.Resources.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Resources, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Resources = from.Resources
	}
}

func (c RequestedClaims) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RequestedClaims) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(RequestedResource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedClaims
// only implements ToObjectValue() and Type().
func (o RequestedClaims) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": o.PermissionSet,
			"resources":      o.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RequestedClaims) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_set": types.StringType,
			"resources": basetypes.ListType{
				ElemType: RequestedResource{}.Type(ctx),
			},
		},
	}
}

// GetResources returns the value of the Resources field in RequestedClaims as
// a slice of RequestedResource values.
// If the field is unknown or null, the boolean return value is false.
func (o *RequestedClaims) GetResources(ctx context.Context) ([]RequestedResource, bool) {
	if o.Resources.IsNull() || o.Resources.IsUnknown() {
		return nil, false
	}
	var v []RequestedResource
	d := o.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in RequestedClaims.
func (o *RequestedClaims) SetResources(ctx context.Context, v []RequestedResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Resources = types.ListValueMust(t, vs)
}

type RequestedResource struct {
	TableName types.String `tfsdk:"table_name"`

	UnspecifiedResourceName types.String `tfsdk:"unspecified_resource_name"`
}

func (to *RequestedResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedResource) {
}

func (to *RequestedResource) SyncFieldsDuringRead(ctx context.Context, from RequestedResource) {
}

func (c RequestedResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RequestedResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedResource
// only implements ToObjectValue() and Type().
func (o RequestedResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":                o.TableName,
			"unspecified_resource_name": o.UnspecifiedResourceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RequestedResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":                types.StringType,
			"unspecified_resource_name": types.StringType,
		},
	}
}

// Next field marker: 14
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

	Spec types.Object `tfsdk:"spec"`
	// The provisioning state of the synced table entity in Unity Catalog. This
	// is distinct from the state of the data synchronization pipeline (i.e. the
	// table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it
	// runs asynchronously).
	UnityCatalogProvisioningState types.String `tfsdk:"unity_catalog_provisioning_state"`
}

func (to *SyncedDatabaseTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedDatabaseTable) {
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

func (to *SyncedDatabaseTable) SyncFieldsDuringRead(ctx context.Context, from SyncedDatabaseTable) {
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

func (c SyncedDatabaseTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["data_synchronization_status"] = attrs["data_synchronization_status"].SetComputed()
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
			"effective_database_instance_name": o.EffectiveDatabaseInstanceName,
			"effective_logical_database_name":  o.EffectiveLogicalDatabaseName,
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
			"effective_database_instance_name": types.StringType,
			"effective_logical_database_name":  types.StringType,
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
	var v SyncedTableStatus
	d := o.DataSynchronizationStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v SyncedTableSpec
	d := o.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
}

func (to *SyncedTableContinuousUpdateStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableContinuousUpdateStatus) {
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

func (to *SyncedTableContinuousUpdateStatus) SyncFieldsDuringRead(ctx context.Context, from SyncedTableContinuousUpdateStatus) {
	if !from.InitialPipelineSyncProgress.IsNull() && !from.InitialPipelineSyncProgress.IsUnknown() {
		if toInitialPipelineSyncProgress, ok := to.GetInitialPipelineSyncProgress(ctx); ok {
			if fromInitialPipelineSyncProgress, ok := from.GetInitialPipelineSyncProgress(ctx); ok {
				toInitialPipelineSyncProgress.SyncFieldsDuringRead(ctx, fromInitialPipelineSyncProgress)
				to.SetInitialPipelineSyncProgress(ctx, toInitialPipelineSyncProgress)
			}
		}
	}
}

func (c SyncedTableContinuousUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetComputed()
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
	var v SyncedTablePipelineProgress
	d := o.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus.
func (o *SyncedTableContinuousUpdateStatus) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.InitialPipelineSyncProgress = vs
}

// Detailed status of a synced table. Shown if the synced table is in the
// OFFLINE_FAILED or the SYNCED_PIPELINE_FAILED state.
type SyncedTableFailedStatus struct {
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

func (to *SyncedTableFailedStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableFailedStatus) {
}

func (to *SyncedTableFailedStatus) SyncFieldsDuringRead(ctx context.Context, from SyncedTableFailedStatus) {
}

func (c SyncedTableFailedStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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

func (to *SyncedTablePipelineProgress) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTablePipelineProgress) {
}

func (to *SyncedTablePipelineProgress) SyncFieldsDuringRead(ctx context.Context, from SyncedTablePipelineProgress) {
}

func (c SyncedTablePipelineProgress) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
			"provisioning_phase":                  o.ProvisioningPhase,
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
			"provisioning_phase":                  types.StringType,
			"sync_progress_completion":            types.Float64Type,
			"synced_row_count":                    types.Int64Type,
			"total_row_count":                     types.Int64Type,
		},
	}
}

type SyncedTablePosition struct {
	DeltaTableSyncInfo types.Object `tfsdk:"delta_table_sync_info"`
	// The end timestamp of the most recent successful synchronization. This is
	// the time when the data is available in the synced table.
	SyncEndTimestamp types.String `tfsdk:"sync_end_timestamp"`
	// The starting timestamp of the most recent successful synchronization from
	// the source table to the destination (synced) table. Note this is the
	// starting timestamp of the sync operation, not the end time. E.g., for a
	// batch, this is the time when the sync operation started.
	SyncStartTimestamp types.String `tfsdk:"sync_start_timestamp"`
}

func (to *SyncedTablePosition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTablePosition) {
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

func (to *SyncedTablePosition) SyncFieldsDuringRead(ctx context.Context, from SyncedTablePosition) {
	if !from.DeltaTableSyncInfo.IsNull() && !from.DeltaTableSyncInfo.IsUnknown() {
		if toDeltaTableSyncInfo, ok := to.GetDeltaTableSyncInfo(ctx); ok {
			if fromDeltaTableSyncInfo, ok := from.GetDeltaTableSyncInfo(ctx); ok {
				toDeltaTableSyncInfo.SyncFieldsDuringRead(ctx, fromDeltaTableSyncInfo)
				to.SetDeltaTableSyncInfo(ctx, toDeltaTableSyncInfo)
			}
		}
	}
}

func (c SyncedTablePosition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["delta_table_sync_info"] = attrs["delta_table_sync_info"].SetComputed()
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
func (a SyncedTablePosition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_table_sync_info": reflect.TypeOf(DeltaTableSyncInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePosition
// only implements ToObjectValue() and Type().
func (o SyncedTablePosition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_table_sync_info": o.DeltaTableSyncInfo,
			"sync_end_timestamp":    o.SyncEndTimestamp,
			"sync_start_timestamp":  o.SyncStartTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SyncedTablePosition) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_table_sync_info": DeltaTableSyncInfo{}.Type(ctx),
			"sync_end_timestamp":    types.StringType,
			"sync_start_timestamp":  types.StringType,
		},
	}
}

// GetDeltaTableSyncInfo returns the value of the DeltaTableSyncInfo field in SyncedTablePosition as
// a DeltaTableSyncInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTablePosition) GetDeltaTableSyncInfo(ctx context.Context) (DeltaTableSyncInfo, bool) {
	var e DeltaTableSyncInfo
	if o.DeltaTableSyncInfo.IsNull() || o.DeltaTableSyncInfo.IsUnknown() {
		return e, false
	}
	var v DeltaTableSyncInfo
	d := o.DeltaTableSyncInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaTableSyncInfo sets the value of the DeltaTableSyncInfo field in SyncedTablePosition.
func (o *SyncedTablePosition) SetDeltaTableSyncInfo(ctx context.Context, v DeltaTableSyncInfo) {
	vs := v.ToObjectValue(ctx)
	o.DeltaTableSyncInfo = vs
}

// Detailed status of a synced table. Shown if the synced table is in the
// PROVISIONING_PIPELINE_RESOURCES or the PROVISIONING_INITIAL_SNAPSHOT state.
type SyncedTableProvisioningStatus struct {
	// Details about initial data synchronization. Only populated when in the
	// PROVISIONING_INITIAL_SNAPSHOT state.
	InitialPipelineSyncProgress types.Object `tfsdk:"initial_pipeline_sync_progress"`
}

func (to *SyncedTableProvisioningStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableProvisioningStatus) {
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

func (to *SyncedTableProvisioningStatus) SyncFieldsDuringRead(ctx context.Context, from SyncedTableProvisioningStatus) {
	if !from.InitialPipelineSyncProgress.IsNull() && !from.InitialPipelineSyncProgress.IsUnknown() {
		if toInitialPipelineSyncProgress, ok := to.GetInitialPipelineSyncProgress(ctx); ok {
			if fromInitialPipelineSyncProgress, ok := from.GetInitialPipelineSyncProgress(ctx); ok {
				toInitialPipelineSyncProgress.SyncFieldsDuringRead(ctx, fromInitialPipelineSyncProgress)
				to.SetInitialPipelineSyncProgress(ctx, toInitialPipelineSyncProgress)
			}
		}
	}
}

func (c SyncedTableProvisioningStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["initial_pipeline_sync_progress"] = attrs["initial_pipeline_sync_progress"].SetComputed()

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
	var v SyncedTablePipelineProgress
	d := o.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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

func (to *SyncedTableSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableSpec) {
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

func (to *SyncedTableSpec) SyncFieldsDuringRead(ctx context.Context, from SyncedTableSpec) {
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

func (c SyncedTableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetOptional()
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].SetComputed()
	attrs["create_database_objects_if_missing"] = attrs["create_database_objects_if_missing"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].SetOptional()
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].SetComputed()
	attrs["existing_pipeline_id"] = attrs["existing_pipeline_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetOptional()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].SetComputed()
	attrs["new_pipeline_spec"] = attrs["new_pipeline_spec"].(tfschema.SingleNestedAttributeBuilder).AddPlanModifier(objectplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
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
	var v NewPipelineSpec
	d := o.NewPipelineSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	ContinuousUpdateStatus types.Object `tfsdk:"continuous_update_status"`
	// The state of the synced table.
	DetailedState types.String `tfsdk:"detailed_state"`

	FailedStatus types.Object `tfsdk:"failed_status"`
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
	LastSync types.Object `tfsdk:"last_sync"`
	// A text description of the current state of the synced table.
	Message types.String `tfsdk:"message"`
	// ID of the associated pipeline. The pipeline ID may have been provided by
	// the client (in the case of bin packing), or generated by the server (when
	// creating a new pipeline).
	PipelineId types.String `tfsdk:"pipeline_id"`

	ProvisioningStatus types.Object `tfsdk:"provisioning_status"`

	TriggeredUpdateStatus types.Object `tfsdk:"triggered_update_status"`
}

func (to *SyncedTableStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableStatus) {
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

func (to *SyncedTableStatus) SyncFieldsDuringRead(ctx context.Context, from SyncedTableStatus) {
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

func (c SyncedTableStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["continuous_update_status"] = attrs["continuous_update_status"].SetOptional()
	attrs["detailed_state"] = attrs["detailed_state"].SetComputed()
	attrs["failed_status"] = attrs["failed_status"].SetOptional()
	attrs["last_sync"] = attrs["last_sync"].SetComputed()
	attrs["message"] = attrs["message"].SetComputed()
	attrs["pipeline_id"] = attrs["pipeline_id"].SetComputed()
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
		"last_sync":                reflect.TypeOf(SyncedTablePosition{}),
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
			"last_sync":                o.LastSync,
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
			"last_sync":                SyncedTablePosition{}.Type(ctx),
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
	var v SyncedTableContinuousUpdateStatus
	d := o.ContinuousUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v SyncedTableFailedStatus
	d := o.FailedStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedStatus sets the value of the FailedStatus field in SyncedTableStatus.
func (o *SyncedTableStatus) SetFailedStatus(ctx context.Context, v SyncedTableFailedStatus) {
	vs := v.ToObjectValue(ctx)
	o.FailedStatus = vs
}

// GetLastSync returns the value of the LastSync field in SyncedTableStatus as
// a SyncedTablePosition value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus) GetLastSync(ctx context.Context) (SyncedTablePosition, bool) {
	var e SyncedTablePosition
	if o.LastSync.IsNull() || o.LastSync.IsUnknown() {
		return e, false
	}
	var v SyncedTablePosition
	d := o.LastSync.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLastSync sets the value of the LastSync field in SyncedTableStatus.
func (o *SyncedTableStatus) SetLastSync(ctx context.Context, v SyncedTablePosition) {
	vs := v.ToObjectValue(ctx)
	o.LastSync = vs
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in SyncedTableStatus as
// a SyncedTableProvisioningStatus value.
// If the field is unknown or null, the boolean return value is false.
func (o *SyncedTableStatus) GetProvisioningStatus(ctx context.Context) (SyncedTableProvisioningStatus, bool) {
	var e SyncedTableProvisioningStatus
	if o.ProvisioningStatus.IsNull() || o.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v SyncedTableProvisioningStatus
	d := o.ProvisioningStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
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
	var v SyncedTableTriggeredUpdateStatus
	d := o.TriggeredUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in SyncedTableStatus.
func (o *SyncedTableStatus) SetTriggeredUpdateStatus(ctx context.Context, v SyncedTableTriggeredUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	o.TriggeredUpdateStatus = vs
}

// Detailed status of a synced table. Shown if the synced table is in the
// SYNCED_TRIGGERED_UPDATE or the SYNCED_NO_PENDING_UPDATE state.
type SyncedTableTriggeredUpdateStatus struct {
	// The last source table Delta version that was successfully synced to the
	// synced table.
	LastProcessedCommitVersion types.Int64 `tfsdk:"last_processed_commit_version"`
	// The end timestamp of the last time any data was synchronized from the
	// source table to the synced table. This is when the data is available in
	// the synced table.
	Timestamp types.String `tfsdk:"timestamp"`
	// Progress of the active data synchronization pipeline.
	TriggeredUpdateProgress types.Object `tfsdk:"triggered_update_progress"`
}

func (to *SyncedTableTriggeredUpdateStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SyncedTableTriggeredUpdateStatus) {
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

func (to *SyncedTableTriggeredUpdateStatus) SyncFieldsDuringRead(ctx context.Context, from SyncedTableTriggeredUpdateStatus) {
	if !from.TriggeredUpdateProgress.IsNull() && !from.TriggeredUpdateProgress.IsUnknown() {
		if toTriggeredUpdateProgress, ok := to.GetTriggeredUpdateProgress(ctx); ok {
			if fromTriggeredUpdateProgress, ok := from.GetTriggeredUpdateProgress(ctx); ok {
				toTriggeredUpdateProgress.SyncFieldsDuringRead(ctx, fromTriggeredUpdateProgress)
				to.SetTriggeredUpdateProgress(ctx, toTriggeredUpdateProgress)
			}
		}
	}
}

func (c SyncedTableTriggeredUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["last_processed_commit_version"] = attrs["last_processed_commit_version"].SetComputed()
	attrs["timestamp"] = attrs["timestamp"].SetComputed()
	attrs["triggered_update_progress"] = attrs["triggered_update_progress"].SetComputed()

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
	var v SyncedTablePipelineProgress
	d := o.TriggeredUpdateProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus.
func (o *SyncedTableTriggeredUpdateStatus) SetTriggeredUpdateProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	o.TriggeredUpdateProgress = vs
}

type UpdateDatabaseCatalogRequest struct {
	// Note that updating a database catalog is not yet supported.
	DatabaseCatalog types.Object `tfsdk:"database_catalog"`
	// The name of the catalog in UC.
	Name types.String `tfsdk:"-"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseCatalogRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseCatalogRequest) {
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

func (to *UpdateDatabaseCatalogRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseCatalogRequest) {
	if !from.DatabaseCatalog.IsNull() && !from.DatabaseCatalog.IsUnknown() {
		if toDatabaseCatalog, ok := to.GetDatabaseCatalog(ctx); ok {
			if fromDatabaseCatalog, ok := from.GetDatabaseCatalog(ctx); ok {
				toDatabaseCatalog.SyncFieldsDuringRead(ctx, fromDatabaseCatalog)
				to.SetDatabaseCatalog(ctx, toDatabaseCatalog)
			}
		}
	}
}

func (c UpdateDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_catalog"] = attrs["database_catalog"].SetRequired()
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
func (a UpdateDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalog": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (o UpdateDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_catalog": o.DatabaseCatalog,
			"name":             o.Name,
			"update_mask":      o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_catalog": DatabaseCatalog{}.Type(ctx),
			"name":             types.StringType,
			"update_mask":      types.StringType,
		},
	}
}

// GetDatabaseCatalog returns the value of the DatabaseCatalog field in UpdateDatabaseCatalogRequest as
// a DatabaseCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateDatabaseCatalogRequest) GetDatabaseCatalog(ctx context.Context) (DatabaseCatalog, bool) {
	var e DatabaseCatalog
	if o.DatabaseCatalog.IsNull() || o.DatabaseCatalog.IsUnknown() {
		return e, false
	}
	var v DatabaseCatalog
	d := o.DatabaseCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseCatalog sets the value of the DatabaseCatalog field in UpdateDatabaseCatalogRequest.
func (o *UpdateDatabaseCatalogRequest) SetDatabaseCatalog(ctx context.Context, v DatabaseCatalog) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseCatalog = vs
}

type UpdateDatabaseInstanceRequest struct {
	DatabaseInstance types.Object `tfsdk:"database_instance"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible. To wipe out custom_tags, specify custom_tags in the
	// update_mask with an empty custom_tags map.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseInstanceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseInstanceRequest) {
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

func (to *UpdateDatabaseInstanceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseInstanceRequest) {
	if !from.DatabaseInstance.IsNull() && !from.DatabaseInstance.IsUnknown() {
		if toDatabaseInstance, ok := to.GetDatabaseInstance(ctx); ok {
			if fromDatabaseInstance, ok := from.GetDatabaseInstance(ctx); ok {
				toDatabaseInstance.SyncFieldsDuringRead(ctx, fromDatabaseInstance)
				to.SetDatabaseInstance(ctx, toDatabaseInstance)
			}
		}
	}
}

func (c UpdateDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance"] = attrs["database_instance"].SetRequired()
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
	var v DatabaseInstance
	d := o.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest.
func (o *UpdateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	o.DatabaseInstance = vs
}

type UpdateSyncedDatabaseTableRequest struct {
	// Full three-part (catalog, schema, table) name of the table.
	Name types.String `tfsdk:"-"`
	// Note that updating a synced database table is not yet supported.
	SyncedTable types.Object `tfsdk:"synced_table"`
	// The list of fields to update. Setting this field is not yet supported.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateSyncedDatabaseTableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSyncedDatabaseTableRequest) {
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

func (to *UpdateSyncedDatabaseTableRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateSyncedDatabaseTableRequest) {
	if !from.SyncedTable.IsNull() && !from.SyncedTable.IsUnknown() {
		if toSyncedTable, ok := to.GetSyncedTable(ctx); ok {
			if fromSyncedTable, ok := from.GetSyncedTable(ctx); ok {
				toSyncedTable.SyncFieldsDuringRead(ctx, fromSyncedTable)
				to.SetSyncedTable(ctx, toSyncedTable)
			}
		}
	}
}

func (c UpdateSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["synced_table"] = attrs["synced_table"].SetRequired()
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
func (a UpdateSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (o UpdateSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         o.Name,
			"synced_table": o.SyncedTable,
			"update_mask":  o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":         types.StringType,
			"synced_table": SyncedDatabaseTable{}.Type(ctx),
			"update_mask":  types.StringType,
		},
	}
}

// GetSyncedTable returns the value of the SyncedTable field in UpdateSyncedDatabaseTableRequest as
// a SyncedDatabaseTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateSyncedDatabaseTableRequest) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable, bool) {
	var e SyncedDatabaseTable
	if o.SyncedTable.IsNull() || o.SyncedTable.IsUnknown() {
		return e, false
	}
	var v SyncedDatabaseTable
	d := o.SyncedTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSyncedTable sets the value of the SyncedTable field in UpdateSyncedDatabaseTableRequest.
func (o *UpdateSyncedDatabaseTableRequest) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable) {
	vs := v.ToObjectValue(ctx)
	o.SyncedTable = vs
}
