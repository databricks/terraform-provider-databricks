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

func (toState *DatabaseCatalog) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseCatalog) {
	if !fromPlan.CreateDatabaseIfNotExists.IsUnknown() && !fromPlan.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the plan.
		toState.CreateDatabaseIfNotExists = fromPlan.CreateDatabaseIfNotExists
	}
}

func (toState *DatabaseCatalog) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseCatalog) {
	if !fromState.CreateDatabaseIfNotExists.IsUnknown() && !fromState.CreateDatabaseIfNotExists.IsNull() {
		// CreateDatabaseIfNotExists is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.CreateDatabaseIfNotExists = fromState.CreateDatabaseIfNotExists
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

func (toState *DatabaseCredential) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseCredential) {
}

func (toState *DatabaseCredential) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseCredential) {
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
	// xref AIP-129. `enable_pg_native_login` is owned by the client, while
	// `effective_enable_pg_native_login` is owned by the server.
	// `enable_pg_native_login` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_enable_pg_native_login` on the other hand will always bet set
	// in all response messages (Create/Update/Get/List).
	EffectiveEnablePgNativeLogin types.Bool `tfsdk:"effective_enable_pg_native_login"`
	// xref AIP-129. `enable_readable_secondaries` is owned by the client, while
	// `effective_enable_readable_secondaries` is owned by the server.
	// `enable_readable_secondaries` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_enable_readable_secondaries` on the other hand will always bet
	// set in all response messages (Create/Update/Get/List).
	EffectiveEnableReadableSecondaries types.Bool `tfsdk:"effective_enable_readable_secondaries"`
	// xref AIP-129. `node_count` is owned by the client, while
	// `effective_node_count` is owned by the server. `node_count` will only be
	// set in Create/Update response messages if and only if the user provides
	// the field via the request. `effective_node_count` on the other hand will
	// always bet set in all response messages (Create/Update/Get/List).
	EffectiveNodeCount types.Int64 `tfsdk:"effective_node_count"`
	// xref AIP-129. `retention_window_in_days` is owned by the client, while
	// `effective_retention_window_in_days` is owned by the server.
	// `retention_window_in_days` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_retention_window_in_days` on the other hand will always bet
	// set in all response messages (Create/Update/Get/List).
	EffectiveRetentionWindowInDays types.Int64 `tfsdk:"effective_retention_window_in_days"`
	// xref AIP-129. `stopped` is owned by the client, while `effective_stopped`
	// is owned by the server. `stopped` will only be set in Create/Update
	// response messages if and only if the user provides the field via the
	// request. `effective_stopped` on the other hand will always bet set in all
	// response messages (Create/Update/Get/List).
	EffectiveStopped types.Bool `tfsdk:"effective_stopped"`
	// Whether the instance has PG native password login enabled. Defaults to
	// true.
	EnablePgNativeLogin types.Bool `tfsdk:"enable_pg_native_login"`
	// Whether to enable secondaries to serve read-only traffic. Defaults to
	// false.
	EnableReadableSecondaries types.Bool `tfsdk:"enable_readable_secondaries"`
	// The name of the instance. This is the unique identifier for the instance.
	Name types.String `tfsdk:"name"`
	// The number of nodes in the instance, composed of 1 primary and 0 or more
	// secondaries. Defaults to 1 primary and 0 secondaries.
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
	// Whether the instance is stopped.
	Stopped types.Bool `tfsdk:"stopped"`
	// An immutable UUID identifier for the instance.
	Uid types.String `tfsdk:"uid"`
}

func (toState *DatabaseInstance) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseInstance) {
	if !fromPlan.EnablePgNativeLogin.IsUnknown() && !fromPlan.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the plan.
		toState.EnablePgNativeLogin = fromPlan.EnablePgNativeLogin
	}
	if !fromPlan.ParentInstanceRef.IsNull() && !fromPlan.ParentInstanceRef.IsUnknown() {
		if toStateParentInstanceRef, ok := toState.GetParentInstanceRef(ctx); ok {
			if fromPlanParentInstanceRef, ok := fromPlan.GetParentInstanceRef(ctx); ok {
				toStateParentInstanceRef.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanParentInstanceRef)
				toState.SetParentInstanceRef(ctx, toStateParentInstanceRef)
			}
		}
	}
}

func (toState *DatabaseInstance) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseInstance) {
	if !fromState.EnablePgNativeLogin.IsUnknown() && !fromState.EnablePgNativeLogin.IsNull() {
		// EnablePgNativeLogin is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.EnablePgNativeLogin = fromState.EnablePgNativeLogin
	}
	if !fromState.ParentInstanceRef.IsNull() && !fromState.ParentInstanceRef.IsUnknown() {
		if toStateParentInstanceRef, ok := toState.GetParentInstanceRef(ctx); ok {
			if fromStateParentInstanceRef, ok := fromState.GetParentInstanceRef(ctx); ok {
				toStateParentInstanceRef.SyncFieldsDuringRead(ctx, fromStateParentInstanceRef)
				toState.SetParentInstanceRef(ctx, toStateParentInstanceRef)
			}
		}
	}
}

func (c DatabaseInstance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["capacity"] = attrs["capacity"].SetOptional()
	attrs["child_instance_refs"] = attrs["child_instance_refs"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["creator"] = attrs["creator"].SetComputed()
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
	// xref AIP-129. `lsn` is owned by the client, while `effective_lsn` is
	// owned by the server. `lsn` will only be set in Create/Update response
	// messages if and only if the user provides the field via the request.
	// `effective_lsn` on the other hand will always bet set in all response
	// messages (Create/Update/Get/List). For a parent ref instance, this is the
	// LSN on the parent instance from which the instance was created. For a
	// child ref instance, this is the LSN on the instance from which the child
	// instance was created.
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

func (toState *DatabaseInstanceRef) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseInstanceRef) {
}

func (toState *DatabaseInstanceRef) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseInstanceRef) {
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

func (toState *DatabaseInstanceRole) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseInstanceRole) {
	if !fromPlan.Attributes.IsNull() && !fromPlan.Attributes.IsUnknown() {
		if toStateAttributes, ok := toState.GetAttributes(ctx); ok {
			if fromPlanAttributes, ok := fromPlan.GetAttributes(ctx); ok {
				toStateAttributes.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanAttributes)
				toState.SetAttributes(ctx, toStateAttributes)
			}
		}
	}
}

func (toState *DatabaseInstanceRole) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseInstanceRole) {
	if !fromState.Attributes.IsNull() && !fromState.Attributes.IsUnknown() {
		if toStateAttributes, ok := toState.GetAttributes(ctx); ok {
			if fromStateAttributes, ok := fromState.GetAttributes(ctx); ok {
				toStateAttributes.SyncFieldsDuringRead(ctx, fromStateAttributes)
				toState.SetAttributes(ctx, toStateAttributes)
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

func (toState *DatabaseInstanceRoleAttributes) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseInstanceRoleAttributes) {
}

func (toState *DatabaseInstanceRoleAttributes) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseInstanceRoleAttributes) {
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

func (toState *DatabaseTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DatabaseTable) {
}

func (toState *DatabaseTable) SyncFieldsDuringRead(ctx context.Context, fromState DatabaseTable) {
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

type DeleteDatabaseInstanceRoleRequest struct {
	// This is the AIP standard name for the equivalent of Postgres' `IF EXISTS`
	// option
	AllowMissing types.Bool `tfsdk:"-"`

	InstanceName types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`

	ReassignOwnedTo types.String `tfsdk:"-"`
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

func (toState *DeltaTableSyncInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan DeltaTableSyncInfo) {
}

func (toState *DeltaTableSyncInfo) SyncFieldsDuringRead(ctx context.Context, fromState DeltaTableSyncInfo) {
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

type GetDatabaseInstanceRoleRequest struct {
	InstanceName types.String `tfsdk:"-"`

	Name types.String `tfsdk:"-"`
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

type ListDatabaseCatalogsRequest struct {
	// Name of the instance to get database catalogs for.
	InstanceName types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of synced database tables.
	// Requests first page if absent.
	PageToken types.String `tfsdk:"-"`
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

func (toState *ListDatabaseCatalogsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDatabaseCatalogsResponse) {
}

func (toState *ListDatabaseCatalogsResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListDatabaseCatalogsResponse) {
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

func (toState *ListDatabaseInstanceRolesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDatabaseInstanceRolesResponse) {
}

func (toState *ListDatabaseInstanceRolesResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListDatabaseInstanceRolesResponse) {
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

func (toState *ListDatabaseInstancesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListDatabaseInstancesResponse) {
}

func (toState *ListDatabaseInstancesResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListDatabaseInstancesResponse) {
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

func (toState *ListSyncedDatabaseTablesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan ListSyncedDatabaseTablesResponse) {
}

func (toState *ListSyncedDatabaseTablesResponse) SyncFieldsDuringRead(ctx context.Context, fromState ListSyncedDatabaseTablesResponse) {
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

func (toState *NewPipelineSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan NewPipelineSpec) {
}

func (toState *NewPipelineSpec) SyncFieldsDuringRead(ctx context.Context, fromState NewPipelineSpec) {
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

func (toState *RequestedClaims) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RequestedClaims) {
}

func (toState *RequestedClaims) SyncFieldsDuringRead(ctx context.Context, fromState RequestedClaims) {
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

func (toState *RequestedResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan RequestedResource) {
}

func (toState *RequestedResource) SyncFieldsDuringRead(ctx context.Context, fromState RequestedResource) {
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

func (toState *SyncedDatabaseTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedDatabaseTable) {
	if !fromPlan.DataSynchronizationStatus.IsNull() && !fromPlan.DataSynchronizationStatus.IsUnknown() {
		if toStateDataSynchronizationStatus, ok := toState.GetDataSynchronizationStatus(ctx); ok {
			if fromPlanDataSynchronizationStatus, ok := fromPlan.GetDataSynchronizationStatus(ctx); ok {
				toStateDataSynchronizationStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDataSynchronizationStatus)
				toState.SetDataSynchronizationStatus(ctx, toStateDataSynchronizationStatus)
			}
		}
	}
	if !fromPlan.DatabaseInstanceName.IsUnknown() && !fromPlan.DatabaseInstanceName.IsNull() {
		// DatabaseInstanceName is an input only field and not returned by the service, so we keep the value from the plan.
		toState.DatabaseInstanceName = fromPlan.DatabaseInstanceName
	}
	if !fromPlan.LogicalDatabaseName.IsUnknown() && !fromPlan.LogicalDatabaseName.IsNull() {
		// LogicalDatabaseName is an input only field and not returned by the service, so we keep the value from the plan.
		toState.LogicalDatabaseName = fromPlan.LogicalDatabaseName
	}
	if !fromPlan.Spec.IsNull() && !fromPlan.Spec.IsUnknown() {
		if toStateSpec, ok := toState.GetSpec(ctx); ok {
			if fromPlanSpec, ok := fromPlan.GetSpec(ctx); ok {
				toStateSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanSpec)
				toState.SetSpec(ctx, toStateSpec)
			}
		}
	}
}

func (toState *SyncedDatabaseTable) SyncFieldsDuringRead(ctx context.Context, fromState SyncedDatabaseTable) {
	if !fromState.DataSynchronizationStatus.IsNull() && !fromState.DataSynchronizationStatus.IsUnknown() {
		if toStateDataSynchronizationStatus, ok := toState.GetDataSynchronizationStatus(ctx); ok {
			if fromStateDataSynchronizationStatus, ok := fromState.GetDataSynchronizationStatus(ctx); ok {
				toStateDataSynchronizationStatus.SyncFieldsDuringRead(ctx, fromStateDataSynchronizationStatus)
				toState.SetDataSynchronizationStatus(ctx, toStateDataSynchronizationStatus)
			}
		}
	}
	if !fromState.DatabaseInstanceName.IsUnknown() && !fromState.DatabaseInstanceName.IsNull() {
		// DatabaseInstanceName is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.DatabaseInstanceName = fromState.DatabaseInstanceName
	}
	if !fromState.LogicalDatabaseName.IsUnknown() && !fromState.LogicalDatabaseName.IsNull() {
		// LogicalDatabaseName is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.LogicalDatabaseName = fromState.LogicalDatabaseName
	}
	if !fromState.Spec.IsNull() && !fromState.Spec.IsUnknown() {
		if toStateSpec, ok := toState.GetSpec(ctx); ok {
			if fromStateSpec, ok := fromState.GetSpec(ctx); ok {
				toStateSpec.SyncFieldsDuringRead(ctx, fromStateSpec)
				toState.SetSpec(ctx, toStateSpec)
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

func (toState *SyncedTableContinuousUpdateStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTableContinuousUpdateStatus) {
	if !fromPlan.InitialPipelineSyncProgress.IsNull() && !fromPlan.InitialPipelineSyncProgress.IsUnknown() {
		if toStateInitialPipelineSyncProgress, ok := toState.GetInitialPipelineSyncProgress(ctx); ok {
			if fromPlanInitialPipelineSyncProgress, ok := fromPlan.GetInitialPipelineSyncProgress(ctx); ok {
				toStateInitialPipelineSyncProgress.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanInitialPipelineSyncProgress)
				toState.SetInitialPipelineSyncProgress(ctx, toStateInitialPipelineSyncProgress)
			}
		}
	}
}

func (toState *SyncedTableContinuousUpdateStatus) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTableContinuousUpdateStatus) {
	if !fromState.InitialPipelineSyncProgress.IsNull() && !fromState.InitialPipelineSyncProgress.IsUnknown() {
		if toStateInitialPipelineSyncProgress, ok := toState.GetInitialPipelineSyncProgress(ctx); ok {
			if fromStateInitialPipelineSyncProgress, ok := fromState.GetInitialPipelineSyncProgress(ctx); ok {
				toStateInitialPipelineSyncProgress.SyncFieldsDuringRead(ctx, fromStateInitialPipelineSyncProgress)
				toState.SetInitialPipelineSyncProgress(ctx, toStateInitialPipelineSyncProgress)
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

func (toState *SyncedTableFailedStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTableFailedStatus) {
}

func (toState *SyncedTableFailedStatus) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTableFailedStatus) {
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

func (toState *SyncedTablePipelineProgress) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTablePipelineProgress) {
}

func (toState *SyncedTablePipelineProgress) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTablePipelineProgress) {
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

func (toState *SyncedTablePosition) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTablePosition) {
	if !fromPlan.DeltaTableSyncInfo.IsNull() && !fromPlan.DeltaTableSyncInfo.IsUnknown() {
		if toStateDeltaTableSyncInfo, ok := toState.GetDeltaTableSyncInfo(ctx); ok {
			if fromPlanDeltaTableSyncInfo, ok := fromPlan.GetDeltaTableSyncInfo(ctx); ok {
				toStateDeltaTableSyncInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanDeltaTableSyncInfo)
				toState.SetDeltaTableSyncInfo(ctx, toStateDeltaTableSyncInfo)
			}
		}
	}
}

func (toState *SyncedTablePosition) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTablePosition) {
	if !fromState.DeltaTableSyncInfo.IsNull() && !fromState.DeltaTableSyncInfo.IsUnknown() {
		if toStateDeltaTableSyncInfo, ok := toState.GetDeltaTableSyncInfo(ctx); ok {
			if fromStateDeltaTableSyncInfo, ok := fromState.GetDeltaTableSyncInfo(ctx); ok {
				toStateDeltaTableSyncInfo.SyncFieldsDuringRead(ctx, fromStateDeltaTableSyncInfo)
				toState.SetDeltaTableSyncInfo(ctx, toStateDeltaTableSyncInfo)
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

func (toState *SyncedTableProvisioningStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTableProvisioningStatus) {
	if !fromPlan.InitialPipelineSyncProgress.IsNull() && !fromPlan.InitialPipelineSyncProgress.IsUnknown() {
		if toStateInitialPipelineSyncProgress, ok := toState.GetInitialPipelineSyncProgress(ctx); ok {
			if fromPlanInitialPipelineSyncProgress, ok := fromPlan.GetInitialPipelineSyncProgress(ctx); ok {
				toStateInitialPipelineSyncProgress.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanInitialPipelineSyncProgress)
				toState.SetInitialPipelineSyncProgress(ctx, toStateInitialPipelineSyncProgress)
			}
		}
	}
}

func (toState *SyncedTableProvisioningStatus) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTableProvisioningStatus) {
	if !fromState.InitialPipelineSyncProgress.IsNull() && !fromState.InitialPipelineSyncProgress.IsUnknown() {
		if toStateInitialPipelineSyncProgress, ok := toState.GetInitialPipelineSyncProgress(ctx); ok {
			if fromStateInitialPipelineSyncProgress, ok := fromState.GetInitialPipelineSyncProgress(ctx); ok {
				toStateInitialPipelineSyncProgress.SyncFieldsDuringRead(ctx, fromStateInitialPipelineSyncProgress)
				toState.SetInitialPipelineSyncProgress(ctx, toStateInitialPipelineSyncProgress)
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

func (toState *SyncedTableSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTableSpec) {
	if !fromPlan.CreateDatabaseObjectsIfMissing.IsUnknown() && !fromPlan.CreateDatabaseObjectsIfMissing.IsNull() {
		// CreateDatabaseObjectsIfMissing is an input only field and not returned by the service, so we keep the value from the plan.
		toState.CreateDatabaseObjectsIfMissing = fromPlan.CreateDatabaseObjectsIfMissing
	}
	if !fromPlan.ExistingPipelineId.IsUnknown() && !fromPlan.ExistingPipelineId.IsNull() {
		// ExistingPipelineId is an input only field and not returned by the service, so we keep the value from the plan.
		toState.ExistingPipelineId = fromPlan.ExistingPipelineId
	}
	if !fromPlan.NewPipelineSpec.IsUnknown() && !fromPlan.NewPipelineSpec.IsNull() {
		// NewPipelineSpec is an input only field and not returned by the service, so we keep the value from the plan.
		toState.NewPipelineSpec = fromPlan.NewPipelineSpec
	}
	if !fromPlan.NewPipelineSpec.IsNull() && !fromPlan.NewPipelineSpec.IsUnknown() {
		if toStateNewPipelineSpec, ok := toState.GetNewPipelineSpec(ctx); ok {
			if fromPlanNewPipelineSpec, ok := fromPlan.GetNewPipelineSpec(ctx); ok {
				toStateNewPipelineSpec.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanNewPipelineSpec)
				toState.SetNewPipelineSpec(ctx, toStateNewPipelineSpec)
			}
		}
	}
}

func (toState *SyncedTableSpec) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTableSpec) {
	if !fromState.CreateDatabaseObjectsIfMissing.IsUnknown() && !fromState.CreateDatabaseObjectsIfMissing.IsNull() {
		// CreateDatabaseObjectsIfMissing is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.CreateDatabaseObjectsIfMissing = fromState.CreateDatabaseObjectsIfMissing
	}
	if !fromState.ExistingPipelineId.IsUnknown() && !fromState.ExistingPipelineId.IsNull() {
		// ExistingPipelineId is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.ExistingPipelineId = fromState.ExistingPipelineId
	}
	if !fromState.NewPipelineSpec.IsUnknown() && !fromState.NewPipelineSpec.IsNull() {
		// NewPipelineSpec is an input only field and not returned by the service, so we keep the value from the existing state.
		toState.NewPipelineSpec = fromState.NewPipelineSpec
	}
	if !fromState.NewPipelineSpec.IsNull() && !fromState.NewPipelineSpec.IsUnknown() {
		if toStateNewPipelineSpec, ok := toState.GetNewPipelineSpec(ctx); ok {
			if fromStateNewPipelineSpec, ok := fromState.GetNewPipelineSpec(ctx); ok {
				toStateNewPipelineSpec.SyncFieldsDuringRead(ctx, fromStateNewPipelineSpec)
				toState.SetNewPipelineSpec(ctx, toStateNewPipelineSpec)
			}
		}
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

func (toState *SyncedTableStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTableStatus) {
	if !fromPlan.ContinuousUpdateStatus.IsNull() && !fromPlan.ContinuousUpdateStatus.IsUnknown() {
		if toStateContinuousUpdateStatus, ok := toState.GetContinuousUpdateStatus(ctx); ok {
			if fromPlanContinuousUpdateStatus, ok := fromPlan.GetContinuousUpdateStatus(ctx); ok {
				toStateContinuousUpdateStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanContinuousUpdateStatus)
				toState.SetContinuousUpdateStatus(ctx, toStateContinuousUpdateStatus)
			}
		}
	}
	if !fromPlan.FailedStatus.IsNull() && !fromPlan.FailedStatus.IsUnknown() {
		if toStateFailedStatus, ok := toState.GetFailedStatus(ctx); ok {
			if fromPlanFailedStatus, ok := fromPlan.GetFailedStatus(ctx); ok {
				toStateFailedStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanFailedStatus)
				toState.SetFailedStatus(ctx, toStateFailedStatus)
			}
		}
	}
	if !fromPlan.LastSync.IsNull() && !fromPlan.LastSync.IsUnknown() {
		if toStateLastSync, ok := toState.GetLastSync(ctx); ok {
			if fromPlanLastSync, ok := fromPlan.GetLastSync(ctx); ok {
				toStateLastSync.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanLastSync)
				toState.SetLastSync(ctx, toStateLastSync)
			}
		}
	}
	if !fromPlan.ProvisioningStatus.IsNull() && !fromPlan.ProvisioningStatus.IsUnknown() {
		if toStateProvisioningStatus, ok := toState.GetProvisioningStatus(ctx); ok {
			if fromPlanProvisioningStatus, ok := fromPlan.GetProvisioningStatus(ctx); ok {
				toStateProvisioningStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanProvisioningStatus)
				toState.SetProvisioningStatus(ctx, toStateProvisioningStatus)
			}
		}
	}
	if !fromPlan.TriggeredUpdateStatus.IsNull() && !fromPlan.TriggeredUpdateStatus.IsUnknown() {
		if toStateTriggeredUpdateStatus, ok := toState.GetTriggeredUpdateStatus(ctx); ok {
			if fromPlanTriggeredUpdateStatus, ok := fromPlan.GetTriggeredUpdateStatus(ctx); ok {
				toStateTriggeredUpdateStatus.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTriggeredUpdateStatus)
				toState.SetTriggeredUpdateStatus(ctx, toStateTriggeredUpdateStatus)
			}
		}
	}
}

func (toState *SyncedTableStatus) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTableStatus) {
	if !fromState.ContinuousUpdateStatus.IsNull() && !fromState.ContinuousUpdateStatus.IsUnknown() {
		if toStateContinuousUpdateStatus, ok := toState.GetContinuousUpdateStatus(ctx); ok {
			if fromStateContinuousUpdateStatus, ok := fromState.GetContinuousUpdateStatus(ctx); ok {
				toStateContinuousUpdateStatus.SyncFieldsDuringRead(ctx, fromStateContinuousUpdateStatus)
				toState.SetContinuousUpdateStatus(ctx, toStateContinuousUpdateStatus)
			}
		}
	}
	if !fromState.FailedStatus.IsNull() && !fromState.FailedStatus.IsUnknown() {
		if toStateFailedStatus, ok := toState.GetFailedStatus(ctx); ok {
			if fromStateFailedStatus, ok := fromState.GetFailedStatus(ctx); ok {
				toStateFailedStatus.SyncFieldsDuringRead(ctx, fromStateFailedStatus)
				toState.SetFailedStatus(ctx, toStateFailedStatus)
			}
		}
	}
	if !fromState.LastSync.IsNull() && !fromState.LastSync.IsUnknown() {
		if toStateLastSync, ok := toState.GetLastSync(ctx); ok {
			if fromStateLastSync, ok := fromState.GetLastSync(ctx); ok {
				toStateLastSync.SyncFieldsDuringRead(ctx, fromStateLastSync)
				toState.SetLastSync(ctx, toStateLastSync)
			}
		}
	}
	if !fromState.ProvisioningStatus.IsNull() && !fromState.ProvisioningStatus.IsUnknown() {
		if toStateProvisioningStatus, ok := toState.GetProvisioningStatus(ctx); ok {
			if fromStateProvisioningStatus, ok := fromState.GetProvisioningStatus(ctx); ok {
				toStateProvisioningStatus.SyncFieldsDuringRead(ctx, fromStateProvisioningStatus)
				toState.SetProvisioningStatus(ctx, toStateProvisioningStatus)
			}
		}
	}
	if !fromState.TriggeredUpdateStatus.IsNull() && !fromState.TriggeredUpdateStatus.IsUnknown() {
		if toStateTriggeredUpdateStatus, ok := toState.GetTriggeredUpdateStatus(ctx); ok {
			if fromStateTriggeredUpdateStatus, ok := fromState.GetTriggeredUpdateStatus(ctx); ok {
				toStateTriggeredUpdateStatus.SyncFieldsDuringRead(ctx, fromStateTriggeredUpdateStatus)
				toState.SetTriggeredUpdateStatus(ctx, toStateTriggeredUpdateStatus)
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

func (toState *SyncedTableTriggeredUpdateStatus) SyncFieldsDuringCreateOrUpdate(ctx context.Context, fromPlan SyncedTableTriggeredUpdateStatus) {
	if !fromPlan.TriggeredUpdateProgress.IsNull() && !fromPlan.TriggeredUpdateProgress.IsUnknown() {
		if toStateTriggeredUpdateProgress, ok := toState.GetTriggeredUpdateProgress(ctx); ok {
			if fromPlanTriggeredUpdateProgress, ok := fromPlan.GetTriggeredUpdateProgress(ctx); ok {
				toStateTriggeredUpdateProgress.SyncFieldsDuringCreateOrUpdate(ctx, fromPlanTriggeredUpdateProgress)
				toState.SetTriggeredUpdateProgress(ctx, toStateTriggeredUpdateProgress)
			}
		}
	}
}

func (toState *SyncedTableTriggeredUpdateStatus) SyncFieldsDuringRead(ctx context.Context, fromState SyncedTableTriggeredUpdateStatus) {
	if !fromState.TriggeredUpdateProgress.IsNull() && !fromState.TriggeredUpdateProgress.IsUnknown() {
		if toStateTriggeredUpdateProgress, ok := toState.GetTriggeredUpdateProgress(ctx); ok {
			if fromStateTriggeredUpdateProgress, ok := fromState.GetTriggeredUpdateProgress(ctx); ok {
				toStateTriggeredUpdateProgress.SyncFieldsDuringRead(ctx, fromStateTriggeredUpdateProgress)
				toState.SetTriggeredUpdateProgress(ctx, toStateTriggeredUpdateProgress)
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
