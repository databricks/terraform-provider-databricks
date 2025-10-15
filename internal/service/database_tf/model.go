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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateDatabaseBranchRequest struct {
	DatabaseBranch types.Object `tfsdk:"database_branch"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *CreateDatabaseBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseBranchRequest) {
	if !from.DatabaseBranch.IsNull() && !from.DatabaseBranch.IsUnknown() {
		if toDatabaseBranch, ok := to.GetDatabaseBranch(ctx); ok {
			if fromDatabaseBranch, ok := from.GetDatabaseBranch(ctx); ok {
				// Recursively sync the fields of DatabaseBranch
				toDatabaseBranch.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseBranch)
				to.SetDatabaseBranch(ctx, toDatabaseBranch)
			}
		}
	}
}

func (to *CreateDatabaseBranchRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseBranchRequest) {
	if !from.DatabaseBranch.IsNull() && !from.DatabaseBranch.IsUnknown() {
		if toDatabaseBranch, ok := to.GetDatabaseBranch(ctx); ok {
			if fromDatabaseBranch, ok := from.GetDatabaseBranch(ctx); ok {
				toDatabaseBranch.SyncFieldsDuringRead(ctx, fromDatabaseBranch)
				to.SetDatabaseBranch(ctx, toDatabaseBranch)
			}
		}
	}
}

func (m CreateDatabaseBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_branch"] = attrs["database_branch"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_branch": reflect.TypeOf(DatabaseBranch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseBranchRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_branch": m.DatabaseBranch,
			"project_id":      m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_branch": DatabaseBranch{}.Type(ctx),
			"project_id":      types.StringType,
		},
	}
}

// GetDatabaseBranch returns the value of the DatabaseBranch field in CreateDatabaseBranchRequest as
// a DatabaseBranch value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseBranchRequest) GetDatabaseBranch(ctx context.Context) (DatabaseBranch, bool) {
	var e DatabaseBranch
	if m.DatabaseBranch.IsNull() || m.DatabaseBranch.IsUnknown() {
		return e, false
	}
	var v DatabaseBranch
	d := m.DatabaseBranch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseBranch sets the value of the DatabaseBranch field in CreateDatabaseBranchRequest.
func (m *CreateDatabaseBranchRequest) SetDatabaseBranch(ctx context.Context, v DatabaseBranch) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseBranch = vs
}

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

func (m CreateDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"catalog": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog": m.Catalog,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog": DatabaseCatalog{}.Type(ctx),
		},
	}
}

// GetCatalog returns the value of the Catalog field in CreateDatabaseCatalogRequest as
// a DatabaseCatalog value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseCatalogRequest) GetCatalog(ctx context.Context) (DatabaseCatalog, bool) {
	var e DatabaseCatalog
	if m.Catalog.IsNull() || m.Catalog.IsUnknown() {
		return e, false
	}
	var v DatabaseCatalog
	d := m.Catalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCatalog sets the value of the Catalog field in CreateDatabaseCatalogRequest.
func (m *CreateDatabaseCatalogRequest) SetCatalog(ctx context.Context, v DatabaseCatalog) {
	vs := v.ToObjectValue(ctx)
	m.Catalog = vs
}

type CreateDatabaseEndpointRequest struct {
	BranchId types.String `tfsdk:"-"`

	DatabaseEndpoint types.Object `tfsdk:"database_endpoint"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *CreateDatabaseEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseEndpointRequest) {
	if !from.DatabaseEndpoint.IsNull() && !from.DatabaseEndpoint.IsUnknown() {
		if toDatabaseEndpoint, ok := to.GetDatabaseEndpoint(ctx); ok {
			if fromDatabaseEndpoint, ok := from.GetDatabaseEndpoint(ctx); ok {
				// Recursively sync the fields of DatabaseEndpoint
				toDatabaseEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseEndpoint)
				to.SetDatabaseEndpoint(ctx, toDatabaseEndpoint)
			}
		}
	}
}

func (to *CreateDatabaseEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseEndpointRequest) {
	if !from.DatabaseEndpoint.IsNull() && !from.DatabaseEndpoint.IsUnknown() {
		if toDatabaseEndpoint, ok := to.GetDatabaseEndpoint(ctx); ok {
			if fromDatabaseEndpoint, ok := from.GetDatabaseEndpoint(ctx); ok {
				toDatabaseEndpoint.SyncFieldsDuringRead(ctx, fromDatabaseEndpoint)
				to.SetDatabaseEndpoint(ctx, toDatabaseEndpoint)
			}
		}
	}
}

func (m CreateDatabaseEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_endpoint"] = attrs["database_endpoint"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_endpoint": reflect.TypeOf(DatabaseEndpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseEndpointRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":         m.BranchId,
			"database_endpoint": m.DatabaseEndpoint,
			"project_id":        m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":         types.StringType,
			"database_endpoint": DatabaseEndpoint{}.Type(ctx),
			"project_id":        types.StringType,
		},
	}
}

// GetDatabaseEndpoint returns the value of the DatabaseEndpoint field in CreateDatabaseEndpointRequest as
// a DatabaseEndpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseEndpointRequest) GetDatabaseEndpoint(ctx context.Context) (DatabaseEndpoint, bool) {
	var e DatabaseEndpoint
	if m.DatabaseEndpoint.IsNull() || m.DatabaseEndpoint.IsUnknown() {
		return e, false
	}
	var v DatabaseEndpoint
	d := m.DatabaseEndpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseEndpoint sets the value of the DatabaseEndpoint field in CreateDatabaseEndpointRequest.
func (m *CreateDatabaseEndpointRequest) SetDatabaseEndpoint(ctx context.Context, v DatabaseEndpoint) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseEndpoint = vs
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

func (m CreateDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": m.DatabaseInstance,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance": DatabaseInstance{}.Type(ctx),
		},
	}
}

// GetDatabaseInstance returns the value of the DatabaseInstance field in CreateDatabaseInstanceRequest as
// a DatabaseInstance value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseInstanceRequest) GetDatabaseInstance(ctx context.Context) (DatabaseInstance, bool) {
	var e DatabaseInstance
	if m.DatabaseInstance.IsNull() || m.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v DatabaseInstance
	d := m.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in CreateDatabaseInstanceRequest.
func (m *CreateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseInstance = vs
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

func (m CreateDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_role": reflect.TypeOf(DatabaseInstanceRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_role": m.DatabaseInstanceRole,
			"instance_name":          m.InstanceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateDatabaseInstanceRoleRequest) GetDatabaseInstanceRole(ctx context.Context) (DatabaseInstanceRole, bool) {
	var e DatabaseInstanceRole
	if m.DatabaseInstanceRole.IsNull() || m.DatabaseInstanceRole.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRole
	d := m.DatabaseInstanceRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstanceRole sets the value of the DatabaseInstanceRole field in CreateDatabaseInstanceRoleRequest.
func (m *CreateDatabaseInstanceRoleRequest) SetDatabaseInstanceRole(ctx context.Context, v DatabaseInstanceRole) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseInstanceRole = vs
}

type CreateDatabaseProjectRequest struct {
	DatabaseProject types.Object `tfsdk:"database_project"`
}

func (to *CreateDatabaseProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateDatabaseProjectRequest) {
	if !from.DatabaseProject.IsNull() && !from.DatabaseProject.IsUnknown() {
		if toDatabaseProject, ok := to.GetDatabaseProject(ctx); ok {
			if fromDatabaseProject, ok := from.GetDatabaseProject(ctx); ok {
				// Recursively sync the fields of DatabaseProject
				toDatabaseProject.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseProject)
				to.SetDatabaseProject(ctx, toDatabaseProject)
			}
		}
	}
}

func (to *CreateDatabaseProjectRequest) SyncFieldsDuringRead(ctx context.Context, from CreateDatabaseProjectRequest) {
	if !from.DatabaseProject.IsNull() && !from.DatabaseProject.IsUnknown() {
		if toDatabaseProject, ok := to.GetDatabaseProject(ctx); ok {
			if fromDatabaseProject, ok := from.GetDatabaseProject(ctx); ok {
				toDatabaseProject.SyncFieldsDuringRead(ctx, fromDatabaseProject)
				to.SetDatabaseProject(ctx, toDatabaseProject)
			}
		}
	}
}

func (m CreateDatabaseProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_project"] = attrs["database_project"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDatabaseProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateDatabaseProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_project": reflect.TypeOf(DatabaseProject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseProjectRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_project": m.DatabaseProject,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_project": DatabaseProject{}.Type(ctx),
		},
	}
}

// GetDatabaseProject returns the value of the DatabaseProject field in CreateDatabaseProjectRequest as
// a DatabaseProject value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseProjectRequest) GetDatabaseProject(ctx context.Context) (DatabaseProject, bool) {
	var e DatabaseProject
	if m.DatabaseProject.IsNull() || m.DatabaseProject.IsUnknown() {
		return e, false
	}
	var v DatabaseProject
	d := m.DatabaseProject.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseProject sets the value of the DatabaseProject field in CreateDatabaseProjectRequest.
func (m *CreateDatabaseProjectRequest) SetDatabaseProject(ctx context.Context, v DatabaseProject) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseProject = vs
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

func (m CreateDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"table": reflect.TypeOf(DatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m CreateDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table": m.Table,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table": DatabaseTable{}.Type(ctx),
		},
	}
}

// GetTable returns the value of the Table field in CreateDatabaseTableRequest as
// a DatabaseTable value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateDatabaseTableRequest) GetTable(ctx context.Context) (DatabaseTable, bool) {
	var e DatabaseTable
	if m.Table.IsNull() || m.Table.IsUnknown() {
		return e, false
	}
	var v DatabaseTable
	d := m.Table.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTable sets the value of the Table field in CreateDatabaseTableRequest.
func (m *CreateDatabaseTableRequest) SetTable(ctx context.Context, v DatabaseTable) {
	vs := v.ToObjectValue(ctx)
	m.Table = vs
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

func (m CreateSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m CreateSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"synced_table": m.SyncedTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"synced_table": SyncedDatabaseTable{}.Type(ctx),
		},
	}
}

// GetSyncedTable returns the value of the SyncedTable field in CreateSyncedDatabaseTableRequest as
// a SyncedDatabaseTable value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateSyncedDatabaseTableRequest) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable, bool) {
	var e SyncedDatabaseTable
	if m.SyncedTable.IsNull() || m.SyncedTable.IsUnknown() {
		return e, false
	}
	var v SyncedDatabaseTable
	d := m.SyncedTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSyncedTable sets the value of the SyncedTable field in CreateSyncedDatabaseTableRequest.
func (m *CreateSyncedDatabaseTableRequest) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable) {
	vs := v.ToObjectValue(ctx)
	m.SyncedTable = vs
}

type CustomTag struct {
	// The key of the custom tag.
	Key types.String `tfsdk:"key"`
	// The value of the custom tag.
	Value types.String `tfsdk:"value"`
}

func (to *CustomTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomTag) {
}

func (to *CustomTag) SyncFieldsDuringRead(ctx context.Context, from CustomTag) {
}

func (m CustomTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CustomTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomTag
// only implements ToObjectValue() and Type().
func (m CustomTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type DatabaseBranch struct {
	BranchId types.String `tfsdk:"branch_id"`
	// A timestamp indicating when the branch was created.
	CreateTime types.String `tfsdk:"create_time"`
	// The branch’s state, indicating if it is initializing, ready for use, or
	// archived.
	CurrentState types.String `tfsdk:"current_state"`
	// Whether the branch is the project's default branch. This field is only
	// returned on create/update responses. See effective_default for the value
	// that is actually applied to the database branch.
	Default types.Bool `tfsdk:"default"`
	// Whether the branch is the project's default branch.
	EffectiveDefault types.Bool `tfsdk:"effective_default"`
	// The logical size of the branch.
	LogicalSizeBytes types.Int64 `tfsdk:"logical_size_bytes"`
	// The id of the parent branch
	ParentId types.String `tfsdk:"parent_id"`
	// The Log Sequence Number (LSN) on the parent branch from which this branch
	// was created. When restoring a branch using the Restore Database Branch
	// endpoint, this value isn’t finalized until all operations related to
	// the restore have completed successfully.
	ParentLsn types.String `tfsdk:"parent_lsn"`
	// The point in time on the parent branch from which this branch was
	// created.
	ParentTime types.String `tfsdk:"parent_time"`

	PendingState types.String `tfsdk:"pending_state"`

	ProjectId types.String `tfsdk:"project_id"`
	// Whether the branch is protected.
	Protected types.Bool `tfsdk:"protected"`
	// A timestamp indicating when the `current_state` began.
	StateChangeTime types.String `tfsdk:"state_change_time"`
	// A timestamp indicating when the branch was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *DatabaseBranch) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseBranch) {
	if !from.Default.IsUnknown() && !from.Default.IsNull() {
		// Default is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Default = from.Default
	}
}

func (to *DatabaseBranch) SyncFieldsDuringRead(ctx context.Context, from DatabaseBranch) {
	if !from.Default.IsUnknown() && !from.Default.IsNull() {
		// Default is an input only field and not returned by the service, so we keep the value from the prior state.
		to.Default = from.Default
	}
}

func (m DatabaseBranch) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_id"] = attrs["branch_id"].SetOptional()
	attrs["branch_id"] = attrs["branch_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["default"] = attrs["default"].SetOptional()
	attrs["default"] = attrs["default"].SetComputed()
	attrs["default"] = attrs["default"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["effective_default"] = attrs["effective_default"].SetComputed()
	attrs["logical_size_bytes"] = attrs["logical_size_bytes"].SetComputed()
	attrs["parent_id"] = attrs["parent_id"].SetOptional()
	attrs["parent_lsn"] = attrs["parent_lsn"].SetOptional()
	attrs["parent_time"] = attrs["parent_time"].SetOptional()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["protected"] = attrs["protected"].SetOptional()
	attrs["state_change_time"] = attrs["state_change_time"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseBranch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseBranch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseBranch
// only implements ToObjectValue() and Type().
func (m DatabaseBranch) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":          m.BranchId,
			"create_time":        m.CreateTime,
			"current_state":      m.CurrentState,
			"default":            m.Default,
			"effective_default":  m.EffectiveDefault,
			"logical_size_bytes": m.LogicalSizeBytes,
			"parent_id":          m.ParentId,
			"parent_lsn":         m.ParentLsn,
			"parent_time":        m.ParentTime,
			"pending_state":      m.PendingState,
			"project_id":         m.ProjectId,
			"protected":          m.Protected,
			"state_change_time":  m.StateChangeTime,
			"update_time":        m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseBranch) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":          types.StringType,
			"create_time":        types.StringType,
			"current_state":      types.StringType,
			"default":            types.BoolType,
			"effective_default":  types.BoolType,
			"logical_size_bytes": types.Int64Type,
			"parent_id":          types.StringType,
			"parent_lsn":         types.StringType,
			"parent_time":        types.StringType,
			"pending_state":      types.StringType,
			"project_id":         types.StringType,
			"protected":          types.BoolType,
			"state_change_time":  types.StringType,
			"update_time":        types.StringType,
		},
	}
}

type DatabaseCatalog struct {
	CreateDatabaseIfNotExists types.Bool `tfsdk:"create_database_if_not_exists"`
	// The branch_id of the database branch associated with the catalog.
	DatabaseBranchId types.String `tfsdk:"database_branch_id"`
	// The name of the DatabaseInstance housing the database.
	DatabaseInstanceName types.String `tfsdk:"database_instance_name"`
	// The name of the database (in a instance) associated with the catalog.
	DatabaseName types.String `tfsdk:"database_name"`
	// The project_id of the database project associated with the catalog.
	DatabaseProjectId types.String `tfsdk:"database_project_id"`
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

func (m DatabaseCatalog) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetOptional()
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].SetComputed()
	attrs["create_database_if_not_exists"] = attrs["create_database_if_not_exists"].(tfschema.BoolAttributeBuilder).AddPlanModifier(boolplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["database_branch_id"] = attrs["database_branch_id"].SetOptional()
	attrs["database_instance_name"] = attrs["database_instance_name"].SetRequired()
	attrs["database_name"] = attrs["database_name"].SetRequired()
	attrs["database_project_id"] = attrs["database_project_id"].SetOptional()
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
func (m DatabaseCatalog) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCatalog
// only implements ToObjectValue() and Type().
func (m DatabaseCatalog) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_database_if_not_exists": m.CreateDatabaseIfNotExists,
			"database_branch_id":            m.DatabaseBranchId,
			"database_instance_name":        m.DatabaseInstanceName,
			"database_name":                 m.DatabaseName,
			"database_project_id":           m.DatabaseProjectId,
			"name":                          m.Name,
			"uid":                           m.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseCatalog) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_database_if_not_exists": types.BoolType,
			"database_branch_id":            types.StringType,
			"database_instance_name":        types.StringType,
			"database_name":                 types.StringType,
			"database_project_id":           types.StringType,
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

func (m DatabaseCredential) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseCredential) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseCredential
// only implements ToObjectValue() and Type().
func (m DatabaseCredential) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"expiration_time": m.ExpirationTime,
			"token":           m.Token,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseCredential) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"expiration_time": types.StringType,
			"token":           types.StringType,
		},
	}
}

type DatabaseEndpoint struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`

	BranchId types.String `tfsdk:"branch_id"`
	// A timestamp indicating when the compute endpoint was created.
	CreateTime types.String `tfsdk:"create_time"`

	CurrentState types.String `tfsdk:"current_state"`
	// Whether to restrict connections to the compute endpoint. Enabling this
	// option schedules a suspend compute operation. A disabled compute endpoint
	// cannot be enabled by a connection or console action.
	Disabled types.Bool `tfsdk:"disabled"`

	EndpointId types.String `tfsdk:"endpoint_id"`
	// The hostname of the compute endpoint. This is the hostname specified when
	// connecting to a database.
	Host types.String `tfsdk:"host"`
	// A timestamp indicating when the compute endpoint was last active.
	LastActiveTime types.String `tfsdk:"last_active_time"`

	PendingState types.String `tfsdk:"pending_state"`

	PoolerMode types.String `tfsdk:"pooler_mode"`

	ProjectId types.String `tfsdk:"project_id"`

	Settings types.Object `tfsdk:"settings"`
	// A timestamp indicating when the compute endpoint was last started.
	StartTime types.String `tfsdk:"start_time"`
	// A timestamp indicating when the compute endpoint was last suspended.
	SuspendTime types.String `tfsdk:"suspend_time"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration types.String `tfsdk:"suspend_timeout_duration"`
	// NOTE: if want type to default to some value set the server then an
	// effective_type field OR make this field REQUIRED
	Type_ types.String `tfsdk:"type"`
	// A timestamp indicating when the compute endpoint was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *DatabaseEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseEndpoint) {
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (to *DatabaseEndpoint) SyncFieldsDuringRead(ctx context.Context, from DatabaseEndpoint) {
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m DatabaseEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["branch_id"] = attrs["branch_id"].SetOptional()
	attrs["branch_id"] = attrs["branch_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["current_state"] = attrs["current_state"].SetComputed()
	attrs["disabled"] = attrs["disabled"].SetOptional()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()
	attrs["host"] = attrs["host"].SetComputed()
	attrs["last_active_time"] = attrs["last_active_time"].SetComputed()
	attrs["pending_state"] = attrs["pending_state"].SetComputed()
	attrs["pooler_mode"] = attrs["pooler_mode"].SetOptional()
	attrs["project_id"] = attrs["project_id"].SetOptional()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetComputed()
	attrs["suspend_time"] = attrs["suspend_time"].SetComputed()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()
	attrs["type"] = attrs["type"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"settings": reflect.TypeOf(DatabaseEndpointSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseEndpoint
// only implements ToObjectValue() and Type().
func (m DatabaseEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu": m.AutoscalingLimitMinCu,
			"branch_id":                m.BranchId,
			"create_time":              m.CreateTime,
			"current_state":            m.CurrentState,
			"disabled":                 m.Disabled,
			"endpoint_id":              m.EndpointId,
			"host":                     m.Host,
			"last_active_time":         m.LastActiveTime,
			"pending_state":            m.PendingState,
			"pooler_mode":              m.PoolerMode,
			"project_id":               m.ProjectId,
			"settings":                 m.Settings,
			"start_time":               m.StartTime,
			"suspend_time":             m.SuspendTime,
			"suspend_timeout_duration": m.SuspendTimeoutDuration,
			"type":                     m.Type_,
			"update_time":              m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"branch_id":                types.StringType,
			"create_time":              types.StringType,
			"current_state":            types.StringType,
			"disabled":                 types.BoolType,
			"endpoint_id":              types.StringType,
			"host":                     types.StringType,
			"last_active_time":         types.StringType,
			"pending_state":            types.StringType,
			"pooler_mode":              types.StringType,
			"project_id":               types.StringType,
			"settings":                 DatabaseEndpointSettings{}.Type(ctx),
			"start_time":               types.StringType,
			"suspend_time":             types.StringType,
			"suspend_timeout_duration": types.StringType,
			"type":                     types.StringType,
			"update_time":              types.StringType,
		},
	}
}

// GetSettings returns the value of the Settings field in DatabaseEndpoint as
// a DatabaseEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseEndpoint) GetSettings(ctx context.Context) (DatabaseEndpointSettings, bool) {
	var e DatabaseEndpointSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v DatabaseEndpointSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in DatabaseEndpoint.
func (m *DatabaseEndpoint) SetSettings(ctx context.Context, v DatabaseEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

// A collection of settings for a compute endpoint
type DatabaseEndpointSettings struct {
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings types.Map `tfsdk:"pgbouncer_settings"`
}

func (to *DatabaseEndpointSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseEndpointSettings) {
}

func (to *DatabaseEndpointSettings) SyncFieldsDuringRead(ctx context.Context, from DatabaseEndpointSettings) {
}

func (m DatabaseEndpointSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["pg_settings"] = attrs["pg_settings"].SetOptional()
	attrs["pgbouncer_settings"] = attrs["pgbouncer_settings"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseEndpointSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseEndpointSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings":        reflect.TypeOf(types.String{}),
		"pgbouncer_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseEndpointSettings
// only implements ToObjectValue() and Type().
func (m DatabaseEndpointSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"pg_settings":        m.PgSettings,
			"pgbouncer_settings": m.PgbouncerSettings,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseEndpointSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pgbouncer_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in DatabaseEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseEndpointSettings) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgSettings.IsNull() || m.PgSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgSettings sets the value of the PgSettings field in DatabaseEndpointSettings.
func (m *DatabaseEndpointSettings) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

// GetPgbouncerSettings returns the value of the PgbouncerSettings field in DatabaseEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseEndpointSettings) GetPgbouncerSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgbouncerSettings.IsNull() || m.PgbouncerSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgbouncerSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgbouncerSettings sets the value of the PgbouncerSettings field in DatabaseEndpointSettings.
func (m *DatabaseEndpointSettings) SetPgbouncerSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pgbouncer_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgbouncerSettings = types.MapValueMust(t, vs)
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
	// The desired usage policy to associate with the instance.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *DatabaseInstance) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseInstance) {
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
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				// Recursively sync the fields of ParentInstanceRef
				toParentInstanceRef.SyncFieldsDuringCreateOrUpdate(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
	if !from.UsagePolicyId.IsUnknown() && !from.UsagePolicyId.IsNull() {
		// UsagePolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.UsagePolicyId = from.UsagePolicyId
	}
}

func (to *DatabaseInstance) SyncFieldsDuringRead(ctx context.Context, from DatabaseInstance) {
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
	if !from.ParentInstanceRef.IsNull() && !from.ParentInstanceRef.IsUnknown() {
		if toParentInstanceRef, ok := to.GetParentInstanceRef(ctx); ok {
			if fromParentInstanceRef, ok := from.GetParentInstanceRef(ctx); ok {
				toParentInstanceRef.SyncFieldsDuringRead(ctx, fromParentInstanceRef)
				to.SetParentInstanceRef(ctx, toParentInstanceRef)
			}
		}
	}
	if !from.UsagePolicyId.IsUnknown() && !from.UsagePolicyId.IsNull() {
		// UsagePolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.UsagePolicyId = from.UsagePolicyId
	}
}

func (m DatabaseInstance) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseInstance) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"child_instance_refs":   reflect.TypeOf(DatabaseInstanceRef{}),
		"custom_tags":           reflect.TypeOf(CustomTag{}),
		"effective_custom_tags": reflect.TypeOf(CustomTag{}),
		"parent_instance_ref":   reflect.TypeOf(DatabaseInstanceRef{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstance
// only implements ToObjectValue() and Type().
func (m DatabaseInstance) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DatabaseInstance) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"capacity": types.StringType,
			"child_instance_refs": basetypes.ListType{
				ElemType: DatabaseInstanceRef{}.Type(ctx),
			},
			"creation_time": types.StringType,
			"creator":       types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: CustomTag{}.Type(ctx),
			},
			"effective_capacity": types.StringType,
			"effective_custom_tags": basetypes.ListType{
				ElemType: CustomTag{}.Type(ctx),
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
			"parent_instance_ref":                   DatabaseInstanceRef{}.Type(ctx),
			"pg_version":                            types.StringType,
			"read_only_dns":                         types.StringType,
			"read_write_dns":                        types.StringType,
			"retention_window_in_days":              types.Int64Type,
			"state":                                 types.StringType,
			"stopped":                               types.BoolType,
			"uid":                                   types.StringType,
			"usage_policy_id":                       types.StringType,
		},
	}
}

// GetChildInstanceRefs returns the value of the ChildInstanceRefs field in DatabaseInstance as
// a slice of DatabaseInstanceRef values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetChildInstanceRefs(ctx context.Context) ([]DatabaseInstanceRef, bool) {
	if m.ChildInstanceRefs.IsNull() || m.ChildInstanceRefs.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstanceRef
	d := m.ChildInstanceRefs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChildInstanceRefs sets the value of the ChildInstanceRefs field in DatabaseInstance.
func (m *DatabaseInstance) SetChildInstanceRefs(ctx context.Context, v []DatabaseInstanceRef) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["child_instance_refs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ChildInstanceRefs = types.ListValueMust(t, vs)
}

// GetCustomTags returns the value of the CustomTags field in DatabaseInstance as
// a slice of CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetCustomTags(ctx context.Context) ([]CustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in DatabaseInstance.
func (m *DatabaseInstance) SetCustomTags(ctx context.Context, v []CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetEffectiveCustomTags returns the value of the EffectiveCustomTags field in DatabaseInstance as
// a slice of CustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetEffectiveCustomTags(ctx context.Context) ([]CustomTag, bool) {
	if m.EffectiveCustomTags.IsNull() || m.EffectiveCustomTags.IsUnknown() {
		return nil, false
	}
	var v []CustomTag
	d := m.EffectiveCustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEffectiveCustomTags sets the value of the EffectiveCustomTags field in DatabaseInstance.
func (m *DatabaseInstance) SetEffectiveCustomTags(ctx context.Context, v []CustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["effective_custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EffectiveCustomTags = types.ListValueMust(t, vs)
}

// GetParentInstanceRef returns the value of the ParentInstanceRef field in DatabaseInstance as
// a DatabaseInstanceRef value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseInstance) GetParentInstanceRef(ctx context.Context) (DatabaseInstanceRef, bool) {
	var e DatabaseInstanceRef
	if m.ParentInstanceRef.IsNull() || m.ParentInstanceRef.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRef
	d := m.ParentInstanceRef.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetParentInstanceRef sets the value of the ParentInstanceRef field in DatabaseInstance.
func (m *DatabaseInstance) SetParentInstanceRef(ctx context.Context, v DatabaseInstanceRef) {
	vs := v.ToObjectValue(ctx)
	m.ParentInstanceRef = vs
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

func (m DatabaseInstanceRef) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseInstanceRef) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRef
// only implements ToObjectValue() and Type().
func (m DatabaseInstanceRef) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DatabaseInstanceRef) Type(ctx context.Context) attr.Type {
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

func (m DatabaseInstanceRole) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseInstanceRole) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attributes": reflect.TypeOf(DatabaseInstanceRoleAttributes{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRole
// only implements ToObjectValue() and Type().
func (m DatabaseInstanceRole) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"attributes":      m.Attributes,
			"identity_type":   m.IdentityType,
			"membership_role": m.MembershipRole,
			"name":            m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseInstanceRole) Type(ctx context.Context) attr.Type {
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
func (m *DatabaseInstanceRole) GetAttributes(ctx context.Context) (DatabaseInstanceRoleAttributes, bool) {
	var e DatabaseInstanceRoleAttributes
	if m.Attributes.IsNull() || m.Attributes.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRoleAttributes
	d := m.Attributes.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAttributes sets the value of the Attributes field in DatabaseInstanceRole.
func (m *DatabaseInstanceRole) SetAttributes(ctx context.Context, v DatabaseInstanceRoleAttributes) {
	vs := v.ToObjectValue(ctx)
	m.Attributes = vs
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

func (m DatabaseInstanceRoleAttributes) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseInstanceRoleAttributes) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseInstanceRoleAttributes
// only implements ToObjectValue() and Type().
func (m DatabaseInstanceRoleAttributes) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"bypassrls":  m.Bypassrls,
			"createdb":   m.Createdb,
			"createrole": m.Createrole,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseInstanceRoleAttributes) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"bypassrls":  types.BoolType,
			"createdb":   types.BoolType,
			"createrole": types.BoolType,
		},
	}
}

type DatabaseProject struct {
	// The logical size limit for a branch.
	BranchLogicalSizeLimitBytes types.Int64 `tfsdk:"branch_logical_size_limit_bytes"`
	// The desired budget policy to associate with the instance. This field is
	// only returned on create/update responses, and represents the customer
	// provided budget policy. See effective_budget_policy_id for the policy
	// that is actually applied to the instance.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The most recent time when any endpoint of this project was active.
	ComputeLastActiveTime types.String `tfsdk:"compute_last_active_time"`
	// A timestamp indicating when the project was created.
	CreateTime types.String `tfsdk:"create_time"`
	// Custom tags associated with the instance.
	CustomTags types.List `tfsdk:"custom_tags"`

	DefaultEndpointSettings types.Object `tfsdk:"default_endpoint_settings"`
	// Human-readable project name.
	DisplayName types.String `tfsdk:"display_name"`
	// The policy that is applied to the instance.
	EffectiveBudgetPolicyId types.String `tfsdk:"effective_budget_policy_id"`
	// The number of seconds to retain the shared history for point in time
	// recovery for all branches in this project.
	HistoryRetentionDuration types.String `tfsdk:"history_retention_duration"`
	// The major Postgres version number. NOTE: fields could be either user-set
	// or server-set. we can't have fields that are optionally user-provided and
	// server-set to default value. TODO: this needs an effective variant or
	// make REQUIRED
	PgVersion types.Int64 `tfsdk:"pg_version"`

	ProjectId types.String `tfsdk:"project_id"`

	Settings types.Object `tfsdk:"settings"`
	// The current space occupied by the project in storage. Synthetic storage
	// size combines the logical data size and Write-Ahead Log (WAL) size for
	// all branches in a project.
	SyntheticStorageSizeBytes types.Int64 `tfsdk:"synthetic_storage_size_bytes"`
	// A timestamp indicating when the project was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *DatabaseProject) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseProject) {
	if !from.BudgetPolicyId.IsUnknown() && !from.BudgetPolicyId.IsNull() {
		// BudgetPolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.BudgetPolicyId = from.BudgetPolicyId
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				// Recursively sync the fields of DefaultEndpointSettings
				toDefaultEndpointSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				// Recursively sync the fields of Settings
				toSettings.SyncFieldsDuringCreateOrUpdate(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (to *DatabaseProject) SyncFieldsDuringRead(ctx context.Context, from DatabaseProject) {
	if !from.BudgetPolicyId.IsUnknown() && !from.BudgetPolicyId.IsNull() {
		// BudgetPolicyId is an input only field and not returned by the service, so we keep the value from the prior state.
		to.BudgetPolicyId = from.BudgetPolicyId
	}
	if !from.CustomTags.IsNull() && !from.CustomTags.IsUnknown() && to.CustomTags.IsNull() && len(from.CustomTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CustomTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CustomTags = from.CustomTags
	}
	if !from.DefaultEndpointSettings.IsNull() && !from.DefaultEndpointSettings.IsUnknown() {
		if toDefaultEndpointSettings, ok := to.GetDefaultEndpointSettings(ctx); ok {
			if fromDefaultEndpointSettings, ok := from.GetDefaultEndpointSettings(ctx); ok {
				toDefaultEndpointSettings.SyncFieldsDuringRead(ctx, fromDefaultEndpointSettings)
				to.SetDefaultEndpointSettings(ctx, toDefaultEndpointSettings)
			}
		}
	}
	if !from.Settings.IsNull() && !from.Settings.IsUnknown() {
		if toSettings, ok := to.GetSettings(ctx); ok {
			if fromSettings, ok := from.GetSettings(ctx); ok {
				toSettings.SyncFieldsDuringRead(ctx, fromSettings)
				to.SetSettings(ctx, toSettings)
			}
		}
	}
}

func (m DatabaseProject) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["branch_logical_size_limit_bytes"] = attrs["branch_logical_size_limit_bytes"].SetComputed()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetComputed()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["compute_last_active_time"] = attrs["compute_last_active_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["custom_tags"] = attrs["custom_tags"].SetOptional()
	attrs["default_endpoint_settings"] = attrs["default_endpoint_settings"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["effective_budget_policy_id"] = attrs["effective_budget_policy_id"].SetComputed()
	attrs["history_retention_duration"] = attrs["history_retention_duration"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].SetOptional()
	attrs["pg_version"] = attrs["pg_version"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["project_id"] = attrs["project_id"].SetOptional()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["settings"] = attrs["settings"].SetOptional()
	attrs["synthetic_storage_size_bytes"] = attrs["synthetic_storage_size_bytes"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseProject.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseProject) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"custom_tags":               reflect.TypeOf(DatabaseProjectCustomTag{}),
		"default_endpoint_settings": reflect.TypeOf(DatabaseProjectDefaultEndpointSettings{}),
		"settings":                  reflect.TypeOf(DatabaseProjectSettings{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseProject
// only implements ToObjectValue() and Type().
func (m DatabaseProject) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_logical_size_limit_bytes": m.BranchLogicalSizeLimitBytes,
			"budget_policy_id":                m.BudgetPolicyId,
			"compute_last_active_time":        m.ComputeLastActiveTime,
			"create_time":                     m.CreateTime,
			"custom_tags":                     m.CustomTags,
			"default_endpoint_settings":       m.DefaultEndpointSettings,
			"display_name":                    m.DisplayName,
			"effective_budget_policy_id":      m.EffectiveBudgetPolicyId,
			"history_retention_duration":      m.HistoryRetentionDuration,
			"pg_version":                      m.PgVersion,
			"project_id":                      m.ProjectId,
			"settings":                        m.Settings,
			"synthetic_storage_size_bytes":    m.SyntheticStorageSizeBytes,
			"update_time":                     m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseProject) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_logical_size_limit_bytes": types.Int64Type,
			"budget_policy_id":                types.StringType,
			"compute_last_active_time":        types.StringType,
			"create_time":                     types.StringType,
			"custom_tags": basetypes.ListType{
				ElemType: DatabaseProjectCustomTag{}.Type(ctx),
			},
			"default_endpoint_settings":    DatabaseProjectDefaultEndpointSettings{}.Type(ctx),
			"display_name":                 types.StringType,
			"effective_budget_policy_id":   types.StringType,
			"history_retention_duration":   types.StringType,
			"pg_version":                   types.Int64Type,
			"project_id":                   types.StringType,
			"settings":                     DatabaseProjectSettings{}.Type(ctx),
			"synthetic_storage_size_bytes": types.Int64Type,
			"update_time":                  types.StringType,
		},
	}
}

// GetCustomTags returns the value of the CustomTags field in DatabaseProject as
// a slice of DatabaseProjectCustomTag values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseProject) GetCustomTags(ctx context.Context) ([]DatabaseProjectCustomTag, bool) {
	if m.CustomTags.IsNull() || m.CustomTags.IsUnknown() {
		return nil, false
	}
	var v []DatabaseProjectCustomTag
	d := m.CustomTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomTags sets the value of the CustomTags field in DatabaseProject.
func (m *DatabaseProject) SetCustomTags(ctx context.Context, v []DatabaseProjectCustomTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CustomTags = types.ListValueMust(t, vs)
}

// GetDefaultEndpointSettings returns the value of the DefaultEndpointSettings field in DatabaseProject as
// a DatabaseProjectDefaultEndpointSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseProject) GetDefaultEndpointSettings(ctx context.Context) (DatabaseProjectDefaultEndpointSettings, bool) {
	var e DatabaseProjectDefaultEndpointSettings
	if m.DefaultEndpointSettings.IsNull() || m.DefaultEndpointSettings.IsUnknown() {
		return e, false
	}
	var v DatabaseProjectDefaultEndpointSettings
	d := m.DefaultEndpointSettings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultEndpointSettings sets the value of the DefaultEndpointSettings field in DatabaseProject.
func (m *DatabaseProject) SetDefaultEndpointSettings(ctx context.Context, v DatabaseProjectDefaultEndpointSettings) {
	vs := v.ToObjectValue(ctx)
	m.DefaultEndpointSettings = vs
}

// GetSettings returns the value of the Settings field in DatabaseProject as
// a DatabaseProjectSettings value.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseProject) GetSettings(ctx context.Context) (DatabaseProjectSettings, bool) {
	var e DatabaseProjectSettings
	if m.Settings.IsNull() || m.Settings.IsUnknown() {
		return e, false
	}
	var v DatabaseProjectSettings
	d := m.Settings.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSettings sets the value of the Settings field in DatabaseProject.
func (m *DatabaseProject) SetSettings(ctx context.Context, v DatabaseProjectSettings) {
	vs := v.ToObjectValue(ctx)
	m.Settings = vs
}

type DatabaseProjectCustomTag struct {
	// The key of the custom tag.
	Key types.String `tfsdk:"key"`
	// The value of the custom tag.
	Value types.String `tfsdk:"value"`
}

func (to *DatabaseProjectCustomTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseProjectCustomTag) {
}

func (to *DatabaseProjectCustomTag) SyncFieldsDuringRead(ctx context.Context, from DatabaseProjectCustomTag) {
}

func (m DatabaseProjectCustomTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetOptional()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseProjectCustomTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseProjectCustomTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseProjectCustomTag
// only implements ToObjectValue() and Type().
func (m DatabaseProjectCustomTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseProjectCustomTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// A collection of settings for a database endpoint.
type DatabaseProjectDefaultEndpointSettings struct {
	// The maximum number of Compute Units.
	AutoscalingLimitMaxCu types.Float64 `tfsdk:"autoscaling_limit_max_cu"`
	// The minimum number of Compute Units.
	AutoscalingLimitMinCu types.Float64 `tfsdk:"autoscaling_limit_min_cu"`
	// A raw representation of Postgres settings.
	PgSettings types.Map `tfsdk:"pg_settings"`
	// A raw representation of PgBouncer settings.
	PgbouncerSettings types.Map `tfsdk:"pgbouncer_settings"`
	// Duration of inactivity after which the compute endpoint is automatically
	// suspended.
	SuspendTimeoutDuration types.String `tfsdk:"suspend_timeout_duration"`
}

func (to *DatabaseProjectDefaultEndpointSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseProjectDefaultEndpointSettings) {
}

func (to *DatabaseProjectDefaultEndpointSettings) SyncFieldsDuringRead(ctx context.Context, from DatabaseProjectDefaultEndpointSettings) {
}

func (m DatabaseProjectDefaultEndpointSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["autoscaling_limit_max_cu"] = attrs["autoscaling_limit_max_cu"].SetOptional()
	attrs["autoscaling_limit_min_cu"] = attrs["autoscaling_limit_min_cu"].SetOptional()
	attrs["pg_settings"] = attrs["pg_settings"].SetOptional()
	attrs["pgbouncer_settings"] = attrs["pgbouncer_settings"].SetOptional()
	attrs["suspend_timeout_duration"] = attrs["suspend_timeout_duration"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseProjectDefaultEndpointSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseProjectDefaultEndpointSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"pg_settings":        reflect.TypeOf(types.String{}),
		"pgbouncer_settings": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseProjectDefaultEndpointSettings
// only implements ToObjectValue() and Type().
func (m DatabaseProjectDefaultEndpointSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"autoscaling_limit_max_cu": m.AutoscalingLimitMaxCu,
			"autoscaling_limit_min_cu": m.AutoscalingLimitMinCu,
			"pg_settings":              m.PgSettings,
			"pgbouncer_settings":       m.PgbouncerSettings,
			"suspend_timeout_duration": m.SuspendTimeoutDuration,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseProjectDefaultEndpointSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"autoscaling_limit_max_cu": types.Float64Type,
			"autoscaling_limit_min_cu": types.Float64Type,
			"pg_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"pgbouncer_settings": basetypes.MapType{
				ElemType: types.StringType,
			},
			"suspend_timeout_duration": types.StringType,
		},
	}
}

// GetPgSettings returns the value of the PgSettings field in DatabaseProjectDefaultEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseProjectDefaultEndpointSettings) GetPgSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgSettings.IsNull() || m.PgSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgSettings sets the value of the PgSettings field in DatabaseProjectDefaultEndpointSettings.
func (m *DatabaseProjectDefaultEndpointSettings) SetPgSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pg_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgSettings = types.MapValueMust(t, vs)
}

// GetPgbouncerSettings returns the value of the PgbouncerSettings field in DatabaseProjectDefaultEndpointSettings as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *DatabaseProjectDefaultEndpointSettings) GetPgbouncerSettings(ctx context.Context) (map[string]types.String, bool) {
	if m.PgbouncerSettings.IsNull() || m.PgbouncerSettings.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.PgbouncerSettings.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPgbouncerSettings sets the value of the PgbouncerSettings field in DatabaseProjectDefaultEndpointSettings.
func (m *DatabaseProjectDefaultEndpointSettings) SetPgbouncerSettings(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pgbouncer_settings"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PgbouncerSettings = types.MapValueMust(t, vs)
}

type DatabaseProjectSettings struct {
	// Sets wal_level=logical for all compute endpoints in this project. All
	// active endpoints will be suspended. Once enabled, logical replication
	// cannot be disabled.
	EnableLogicalReplication types.Bool `tfsdk:"enable_logical_replication"`
}

func (to *DatabaseProjectSettings) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseProjectSettings) {
}

func (to *DatabaseProjectSettings) SyncFieldsDuringRead(ctx context.Context, from DatabaseProjectSettings) {
}

func (m DatabaseProjectSettings) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enable_logical_replication"] = attrs["enable_logical_replication"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabaseProjectSettings.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabaseProjectSettings) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseProjectSettings
// only implements ToObjectValue() and Type().
func (m DatabaseProjectSettings) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enable_logical_replication": m.EnableLogicalReplication,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseProjectSettings) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enable_logical_replication": types.BoolType,
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
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url"`
}

func (to *DatabaseTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabaseTable) {
}

func (to *DatabaseTable) SyncFieldsDuringRead(ctx context.Context, from DatabaseTable) {
}

func (m DatabaseTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DatabaseTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabaseTable
// only implements ToObjectValue() and Type().
func (m DatabaseTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_name": m.DatabaseInstanceName,
			"logical_database_name":  m.LogicalDatabaseName,
			"name":                   m.Name,
			"table_serving_url":      m.TableServingUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabaseTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_name": types.StringType,
			"logical_database_name":  types.StringType,
			"name":                   types.StringType,
			"table_serving_url":      types.StringType,
		},
	}
}

type DeleteDatabaseBranchRequest struct {
	BranchId types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseBranchRequest) {
}

func (to *DeleteDatabaseBranchRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseBranchRequest) {
}

func (m DeleteDatabaseBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseBranchRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":  m.BranchId,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":  types.StringType,
			"project_id": types.StringType,
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

func (m DeleteDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type DeleteDatabaseEndpointRequest struct {
	BranchId types.String `tfsdk:"-"`

	EndpointId types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseEndpointRequest) {
}

func (to *DeleteDatabaseEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseEndpointRequest) {
}

func (m DeleteDatabaseEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseEndpointRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":   m.BranchId,
			"endpoint_id": m.EndpointId,
			"project_id":  m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":   types.StringType,
			"endpoint_id": types.StringType,
			"project_id":  types.StringType,
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

func (m DeleteDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"force": m.Force,
			"name":  m.Name,
			"purge": m.Purge,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
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

func (m DeleteDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m DeleteDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing":     types.BoolType,
			"instance_name":     types.StringType,
			"name":              types.StringType,
			"reassign_owned_to": types.StringType,
		},
	}
}

type DeleteDatabaseProjectRequest struct {
	ProjectId types.String `tfsdk:"-"`
}

func (to *DeleteDatabaseProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDatabaseProjectRequest) {
}

func (to *DeleteDatabaseProjectRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDatabaseProjectRequest) {
}

func (m DeleteDatabaseProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDatabaseProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDatabaseProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseProjectRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
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

func (m DeleteDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m DeleteDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDatabaseTableRequest) Type(ctx context.Context) attr.Type {
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

func (m DeleteSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m DeleteSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
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

func (m DeltaTableSyncInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeltaTableSyncInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeltaTableSyncInfo
// only implements ToObjectValue() and Type().
func (m DeltaTableSyncInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_commit_timestamp": m.DeltaCommitTimestamp,
			"delta_commit_version":   m.DeltaCommitVersion,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeltaTableSyncInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"delta_commit_timestamp": types.StringType,
			"delta_commit_version":   types.Int64Type,
		},
	}
}

type FailoverDatabaseInstanceRequest struct {
	FailoverTargetDatabaseInstanceName types.String `tfsdk:"failover_target_database_instance_name"`
	// Name of the instance to failover.
	Name types.String `tfsdk:"-"`
}

func (to *FailoverDatabaseInstanceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FailoverDatabaseInstanceRequest) {
}

func (to *FailoverDatabaseInstanceRequest) SyncFieldsDuringRead(ctx context.Context, from FailoverDatabaseInstanceRequest) {
}

func (m FailoverDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["failover_target_database_instance_name"] = attrs["failover_target_database_instance_name"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FailoverDatabaseInstanceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FailoverDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailoverDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (m FailoverDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"failover_target_database_instance_name": m.FailoverTargetDatabaseInstanceName,
			"name":                                   m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FailoverDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"failover_target_database_instance_name": types.StringType,
			"name":                                   types.StringType,
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

func (m FindDatabaseInstanceByUidRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m FindDatabaseInstanceByUidRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FindDatabaseInstanceByUidRequest
// only implements ToObjectValue() and Type().
func (m FindDatabaseInstanceByUidRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"uid": m.Uid,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FindDatabaseInstanceByUidRequest) Type(ctx context.Context) attr.Type {
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

func (m GenerateDatabaseCredentialRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GenerateDatabaseCredentialRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"claims":         reflect.TypeOf(RequestedClaims{}),
		"instance_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenerateDatabaseCredentialRequest
// only implements ToObjectValue() and Type().
func (m GenerateDatabaseCredentialRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"claims":         m.Claims,
			"instance_names": m.InstanceNames,
			"request_id":     m.RequestId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenerateDatabaseCredentialRequest) Type(ctx context.Context) attr.Type {
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
func (m *GenerateDatabaseCredentialRequest) GetClaims(ctx context.Context) ([]RequestedClaims, bool) {
	if m.Claims.IsNull() || m.Claims.IsUnknown() {
		return nil, false
	}
	var v []RequestedClaims
	d := m.Claims.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetClaims sets the value of the Claims field in GenerateDatabaseCredentialRequest.
func (m *GenerateDatabaseCredentialRequest) SetClaims(ctx context.Context, v []RequestedClaims) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["claims"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Claims = types.ListValueMust(t, vs)
}

// GetInstanceNames returns the value of the InstanceNames field in GenerateDatabaseCredentialRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GenerateDatabaseCredentialRequest) GetInstanceNames(ctx context.Context) ([]types.String, bool) {
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

// SetInstanceNames sets the value of the InstanceNames field in GenerateDatabaseCredentialRequest.
func (m *GenerateDatabaseCredentialRequest) SetInstanceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["instance_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InstanceNames = types.ListValueMust(t, vs)
}

type GetDatabaseBranchRequest struct {
	BranchId types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *GetDatabaseBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseBranchRequest) {
}

func (to *GetDatabaseBranchRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseBranchRequest) {
}

func (m GetDatabaseBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseBranchRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":  m.BranchId,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":  types.StringType,
			"project_id": types.StringType,
		},
	}
}

type GetDatabaseCatalogRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *GetDatabaseCatalogRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseCatalogRequest) {
}

func (to *GetDatabaseCatalogRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseCatalogRequest) {
}

func (m GetDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetDatabaseEndpointRequest struct {
	BranchId types.String `tfsdk:"-"`

	EndpointId types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *GetDatabaseEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseEndpointRequest) {
}

func (to *GetDatabaseEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseEndpointRequest) {
}

func (m GetDatabaseEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseEndpointRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":   m.BranchId,
			"endpoint_id": m.EndpointId,
			"project_id":  m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":   types.StringType,
			"endpoint_id": types.StringType,
			"project_id":  types.StringType,
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

func (m GetDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
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

func (m GetDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"name":          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"instance_name": types.StringType,
			"name":          types.StringType,
		},
	}
}

type GetDatabaseProjectRequest struct {
	ProjectId types.String `tfsdk:"-"`
}

func (to *GetDatabaseProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDatabaseProjectRequest) {
}

func (to *GetDatabaseProjectRequest) SyncFieldsDuringRead(ctx context.Context, from GetDatabaseProjectRequest) {
}

func (m GetDatabaseProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDatabaseProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDatabaseProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseProjectRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"project_id": types.StringType,
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

func (m GetDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m GetDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDatabaseTableRequest) Type(ctx context.Context) attr.Type {
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

func (m GetSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m GetSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ListDatabaseBranchesRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Branches. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *ListDatabaseBranchesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseBranchesRequest) {
}

func (to *ListDatabaseBranchesRequest) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseBranchesRequest) {
}

func (m ListDatabaseBranchesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseBranchesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseBranchesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseBranchesRequest
// only implements ToObjectValue() and Type().
func (m ListDatabaseBranchesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseBranchesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"project_id": types.StringType,
		},
	}
}

type ListDatabaseBranchesResponse struct {
	// List of branches.
	DatabaseBranches types.List `tfsdk:"database_branches"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseBranchesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseBranchesResponse) {
	if !from.DatabaseBranches.IsNull() && !from.DatabaseBranches.IsUnknown() && to.DatabaseBranches.IsNull() && len(from.DatabaseBranches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseBranches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseBranches = from.DatabaseBranches
	}
}

func (to *ListDatabaseBranchesResponse) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseBranchesResponse) {
	if !from.DatabaseBranches.IsNull() && !from.DatabaseBranches.IsUnknown() && to.DatabaseBranches.IsNull() && len(from.DatabaseBranches.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseBranches, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseBranches = from.DatabaseBranches
	}
}

func (m ListDatabaseBranchesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_branches"] = attrs["database_branches"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseBranchesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseBranchesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_branches": reflect.TypeOf(DatabaseBranch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseBranchesResponse
// only implements ToObjectValue() and Type().
func (m ListDatabaseBranchesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_branches": m.DatabaseBranches,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseBranchesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_branches": basetypes.ListType{
				ElemType: DatabaseBranch{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseBranches returns the value of the DatabaseBranches field in ListDatabaseBranchesResponse as
// a slice of DatabaseBranch values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatabaseBranchesResponse) GetDatabaseBranches(ctx context.Context) ([]DatabaseBranch, bool) {
	if m.DatabaseBranches.IsNull() || m.DatabaseBranches.IsUnknown() {
		return nil, false
	}
	var v []DatabaseBranch
	d := m.DatabaseBranches.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseBranches sets the value of the DatabaseBranches field in ListDatabaseBranchesResponse.
func (m *ListDatabaseBranchesResponse) SetDatabaseBranches(ctx context.Context, v []DatabaseBranch) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_branches"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseBranches = types.ListValueMust(t, vs)
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

func (m ListDatabaseCatalogsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDatabaseCatalogsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseCatalogsRequest
// only implements ToObjectValue() and Type().
func (m ListDatabaseCatalogsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"page_size":     m.PageSize,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseCatalogsRequest) Type(ctx context.Context) attr.Type {
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

func (m ListDatabaseCatalogsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDatabaseCatalogsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalogs": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseCatalogsResponse
// only implements ToObjectValue() and Type().
func (m ListDatabaseCatalogsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_catalogs": m.DatabaseCatalogs,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseCatalogsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListDatabaseCatalogsResponse) GetDatabaseCatalogs(ctx context.Context) ([]DatabaseCatalog, bool) {
	if m.DatabaseCatalogs.IsNull() || m.DatabaseCatalogs.IsUnknown() {
		return nil, false
	}
	var v []DatabaseCatalog
	d := m.DatabaseCatalogs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseCatalogs sets the value of the DatabaseCatalogs field in ListDatabaseCatalogsResponse.
func (m *ListDatabaseCatalogsResponse) SetDatabaseCatalogs(ctx context.Context, v []DatabaseCatalog) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_catalogs"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseCatalogs = types.ListValueMust(t, vs)
}

type ListDatabaseEndpointsRequest struct {
	BranchId types.String `tfsdk:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Endpoints. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *ListDatabaseEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseEndpointsRequest) {
}

func (to *ListDatabaseEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseEndpointsRequest) {
}

func (m ListDatabaseEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseEndpointsRequest
// only implements ToObjectValue() and Type().
func (m ListDatabaseEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":  m.BranchId,
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
			"project_id": m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":  types.StringType,
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
			"project_id": types.StringType,
		},
	}
}

type ListDatabaseEndpointsResponse struct {
	// List of endpoints.
	DatabaseEndpoints types.List `tfsdk:"database_endpoints"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseEndpointsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseEndpointsResponse) {
	if !from.DatabaseEndpoints.IsNull() && !from.DatabaseEndpoints.IsUnknown() && to.DatabaseEndpoints.IsNull() && len(from.DatabaseEndpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseEndpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseEndpoints = from.DatabaseEndpoints
	}
}

func (to *ListDatabaseEndpointsResponse) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseEndpointsResponse) {
	if !from.DatabaseEndpoints.IsNull() && !from.DatabaseEndpoints.IsUnknown() && to.DatabaseEndpoints.IsNull() && len(from.DatabaseEndpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseEndpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseEndpoints = from.DatabaseEndpoints
	}
}

func (m ListDatabaseEndpointsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_endpoints"] = attrs["database_endpoints"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseEndpointsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_endpoints": reflect.TypeOf(DatabaseEndpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseEndpointsResponse
// only implements ToObjectValue() and Type().
func (m ListDatabaseEndpointsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_endpoints": m.DatabaseEndpoints,
			"next_page_token":    m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseEndpointsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_endpoints": basetypes.ListType{
				ElemType: DatabaseEndpoint{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseEndpoints returns the value of the DatabaseEndpoints field in ListDatabaseEndpointsResponse as
// a slice of DatabaseEndpoint values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatabaseEndpointsResponse) GetDatabaseEndpoints(ctx context.Context) ([]DatabaseEndpoint, bool) {
	if m.DatabaseEndpoints.IsNull() || m.DatabaseEndpoints.IsUnknown() {
		return nil, false
	}
	var v []DatabaseEndpoint
	d := m.DatabaseEndpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseEndpoints sets the value of the DatabaseEndpoints field in ListDatabaseEndpointsResponse.
func (m *ListDatabaseEndpointsResponse) SetDatabaseEndpoints(ctx context.Context, v []DatabaseEndpoint) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseEndpoints = types.ListValueMust(t, vs)
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

func (m ListDatabaseInstanceRolesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDatabaseInstanceRolesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstanceRolesRequest
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstanceRolesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"page_size":     m.PageSize,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstanceRolesRequest) Type(ctx context.Context) attr.Type {
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

func (m ListDatabaseInstanceRolesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDatabaseInstanceRolesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_roles": reflect.TypeOf(DatabaseInstanceRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstanceRolesResponse
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstanceRolesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_roles": m.DatabaseInstanceRoles,
			"next_page_token":         m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstanceRolesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListDatabaseInstanceRolesResponse) GetDatabaseInstanceRoles(ctx context.Context) ([]DatabaseInstanceRole, bool) {
	if m.DatabaseInstanceRoles.IsNull() || m.DatabaseInstanceRoles.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstanceRole
	d := m.DatabaseInstanceRoles.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstanceRoles sets the value of the DatabaseInstanceRoles field in ListDatabaseInstanceRolesResponse.
func (m *ListDatabaseInstanceRolesResponse) SetDatabaseInstanceRoles(ctx context.Context, v []DatabaseInstanceRole) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instance_roles"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseInstanceRoles = types.ListValueMust(t, vs)
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

func (m ListDatabaseInstancesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDatabaseInstancesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesRequest
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstancesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstancesRequest) Type(ctx context.Context) attr.Type {
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

func (m ListDatabaseInstancesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListDatabaseInstancesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instances": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseInstancesResponse
// only implements ToObjectValue() and Type().
func (m ListDatabaseInstancesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instances": m.DatabaseInstances,
			"next_page_token":    m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseInstancesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListDatabaseInstancesResponse) GetDatabaseInstances(ctx context.Context) ([]DatabaseInstance, bool) {
	if m.DatabaseInstances.IsNull() || m.DatabaseInstances.IsUnknown() {
		return nil, false
	}
	var v []DatabaseInstance
	d := m.DatabaseInstances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstances sets the value of the DatabaseInstances field in ListDatabaseInstancesResponse.
func (m *ListDatabaseInstancesResponse) SetDatabaseInstances(ctx context.Context, v []DatabaseInstance) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseInstances = types.ListValueMust(t, vs)
}

type ListDatabaseProjectsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-"`
	// Pagination token to go to the next page of Database Projects. Requests
	// first page if absent.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListDatabaseProjectsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseProjectsRequest) {
}

func (to *ListDatabaseProjectsRequest) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseProjectsRequest) {
}

func (m ListDatabaseProjectsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseProjectsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseProjectsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseProjectsRequest
// only implements ToObjectValue() and Type().
func (m ListDatabaseProjectsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseProjectsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListDatabaseProjectsResponse struct {
	// List of projects.
	DatabaseProjects types.List `tfsdk:"database_projects"`
	// Pagination token to request the next page of instances.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListDatabaseProjectsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListDatabaseProjectsResponse) {
	if !from.DatabaseProjects.IsNull() && !from.DatabaseProjects.IsUnknown() && to.DatabaseProjects.IsNull() && len(from.DatabaseProjects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseProjects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseProjects = from.DatabaseProjects
	}
}

func (to *ListDatabaseProjectsResponse) SyncFieldsDuringRead(ctx context.Context, from ListDatabaseProjectsResponse) {
	if !from.DatabaseProjects.IsNull() && !from.DatabaseProjects.IsUnknown() && to.DatabaseProjects.IsNull() && len(from.DatabaseProjects.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DatabaseProjects, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DatabaseProjects = from.DatabaseProjects
	}
}

func (m ListDatabaseProjectsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_projects"] = attrs["database_projects"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDatabaseProjectsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListDatabaseProjectsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_projects": reflect.TypeOf(DatabaseProject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListDatabaseProjectsResponse
// only implements ToObjectValue() and Type().
func (m ListDatabaseProjectsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_projects": m.DatabaseProjects,
			"next_page_token":   m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListDatabaseProjectsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_projects": basetypes.ListType{
				ElemType: DatabaseProject{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetDatabaseProjects returns the value of the DatabaseProjects field in ListDatabaseProjectsResponse as
// a slice of DatabaseProject values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListDatabaseProjectsResponse) GetDatabaseProjects(ctx context.Context) ([]DatabaseProject, bool) {
	if m.DatabaseProjects.IsNull() || m.DatabaseProjects.IsUnknown() {
		return nil, false
	}
	var v []DatabaseProject
	d := m.DatabaseProjects.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseProjects sets the value of the DatabaseProjects field in ListDatabaseProjectsResponse.
func (m *ListDatabaseProjectsResponse) SetDatabaseProjects(ctx context.Context, v []DatabaseProject) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["database_projects"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DatabaseProjects = types.ListValueMust(t, vs)
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

func (m ListSyncedDatabaseTablesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSyncedDatabaseTablesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSyncedDatabaseTablesRequest
// only implements ToObjectValue() and Type().
func (m ListSyncedDatabaseTablesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"instance_name": m.InstanceName,
			"page_size":     m.PageSize,
			"page_token":    m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSyncedDatabaseTablesRequest) Type(ctx context.Context) attr.Type {
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

func (m ListSyncedDatabaseTablesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListSyncedDatabaseTablesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_tables": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListSyncedDatabaseTablesResponse
// only implements ToObjectValue() and Type().
func (m ListSyncedDatabaseTablesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"synced_tables":   m.SyncedTables,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListSyncedDatabaseTablesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListSyncedDatabaseTablesResponse) GetSyncedTables(ctx context.Context) ([]SyncedDatabaseTable, bool) {
	if m.SyncedTables.IsNull() || m.SyncedTables.IsUnknown() {
		return nil, false
	}
	var v []SyncedDatabaseTable
	d := m.SyncedTables.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSyncedTables sets the value of the SyncedTables field in ListSyncedDatabaseTablesResponse.
func (m *ListSyncedDatabaseTablesResponse) SetSyncedTables(ctx context.Context, v []SyncedDatabaseTable) {
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
type NewPipelineSpec struct {
	// Budget policy of this pipeline.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
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

func (m NewPipelineSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
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
func (m NewPipelineSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NewPipelineSpec
// only implements ToObjectValue() and Type().
func (m NewPipelineSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"budget_policy_id": m.BudgetPolicyId,
			"storage_catalog":  m.StorageCatalog,
			"storage_schema":   m.StorageSchema,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NewPipelineSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"budget_policy_id": types.StringType,
			"storage_catalog":  types.StringType,
			"storage_schema":   types.StringType,
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

func (m RequestedClaims) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RequestedClaims) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"resources": reflect.TypeOf(RequestedResource{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedClaims
// only implements ToObjectValue() and Type().
func (m RequestedClaims) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_set": m.PermissionSet,
			"resources":      m.Resources,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedClaims) Type(ctx context.Context) attr.Type {
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
func (m *RequestedClaims) GetResources(ctx context.Context) ([]RequestedResource, bool) {
	if m.Resources.IsNull() || m.Resources.IsUnknown() {
		return nil, false
	}
	var v []RequestedResource
	d := m.Resources.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResources sets the value of the Resources field in RequestedClaims.
func (m *RequestedClaims) SetResources(ctx context.Context, v []RequestedResource) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resources"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Resources = types.ListValueMust(t, vs)
}

type RequestedResource struct {
	TableName types.String `tfsdk:"table_name"`

	UnspecifiedResourceName types.String `tfsdk:"unspecified_resource_name"`
}

func (to *RequestedResource) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RequestedResource) {
}

func (to *RequestedResource) SyncFieldsDuringRead(ctx context.Context, from RequestedResource) {
}

func (m RequestedResource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RequestedResource) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RequestedResource
// only implements ToObjectValue() and Type().
func (m RequestedResource) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"table_name":                m.TableName,
			"unspecified_resource_name": m.UnspecifiedResourceName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RequestedResource) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"table_name":                types.StringType,
			"unspecified_resource_name": types.StringType,
		},
	}
}

type RestartDatabaseEndpointRequest struct {
	BranchId types.String `tfsdk:"-"`

	EndpointId types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
}

func (to *RestartDatabaseEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestartDatabaseEndpointRequest) {
}

func (to *RestartDatabaseEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from RestartDatabaseEndpointRequest) {
}

func (m RestartDatabaseEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestartDatabaseEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestartDatabaseEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestartDatabaseEndpointRequest
// only implements ToObjectValue() and Type().
func (m RestartDatabaseEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":   m.BranchId,
			"endpoint_id": m.EndpointId,
			"project_id":  m.ProjectId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestartDatabaseEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":   types.StringType,
			"endpoint_id": types.StringType,
			"project_id":  types.StringType,
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
	// Data serving REST API URL for this table
	TableServingUrl types.String `tfsdk:"table_serving_url"`
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

func (m SyncedDatabaseTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedDatabaseTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"data_synchronization_status": reflect.TypeOf(SyncedTableStatus{}),
		"spec":                        reflect.TypeOf(SyncedTableSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedDatabaseTable
// only implements ToObjectValue() and Type().
func (m SyncedDatabaseTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
			"table_serving_url":                m.TableServingUrl,
			"unity_catalog_provisioning_state": m.UnityCatalogProvisioningState,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedDatabaseTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"data_synchronization_status":      SyncedTableStatus{}.Type(ctx),
			"database_instance_name":           types.StringType,
			"effective_database_instance_name": types.StringType,
			"effective_logical_database_name":  types.StringType,
			"logical_database_name":            types.StringType,
			"name":                             types.StringType,
			"spec":                             SyncedTableSpec{}.Type(ctx),
			"table_serving_url":                types.StringType,
			"unity_catalog_provisioning_state": types.StringType,
		},
	}
}

// GetDataSynchronizationStatus returns the value of the DataSynchronizationStatus field in SyncedDatabaseTable as
// a SyncedTableStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedDatabaseTable) GetDataSynchronizationStatus(ctx context.Context) (SyncedTableStatus, bool) {
	var e SyncedTableStatus
	if m.DataSynchronizationStatus.IsNull() || m.DataSynchronizationStatus.IsUnknown() {
		return e, false
	}
	var v SyncedTableStatus
	d := m.DataSynchronizationStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataSynchronizationStatus sets the value of the DataSynchronizationStatus field in SyncedDatabaseTable.
func (m *SyncedDatabaseTable) SetDataSynchronizationStatus(ctx context.Context, v SyncedTableStatus) {
	vs := v.ToObjectValue(ctx)
	m.DataSynchronizationStatus = vs
}

// GetSpec returns the value of the Spec field in SyncedDatabaseTable as
// a SyncedTableSpec value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedDatabaseTable) GetSpec(ctx context.Context) (SyncedTableSpec, bool) {
	var e SyncedTableSpec
	if m.Spec.IsNull() || m.Spec.IsUnknown() {
		return e, false
	}
	var v SyncedTableSpec
	d := m.Spec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSpec sets the value of the Spec field in SyncedDatabaseTable.
func (m *SyncedDatabaseTable) SetSpec(ctx context.Context, v SyncedTableSpec) {
	vs := v.ToObjectValue(ctx)
	m.Spec = vs
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

func (m SyncedTableContinuousUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTableContinuousUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableContinuousUpdateStatus
// only implements ToObjectValue() and Type().
func (m SyncedTableContinuousUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": m.InitialPipelineSyncProgress,
			"last_processed_commit_version":  m.LastProcessedCommitVersion,
			"timestamp":                      m.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableContinuousUpdateStatus) Type(ctx context.Context) attr.Type {
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
func (m *SyncedTableContinuousUpdateStatus) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress, bool) {
	var e SyncedTablePipelineProgress
	if m.InitialPipelineSyncProgress.IsNull() || m.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v SyncedTablePipelineProgress
	d := m.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableContinuousUpdateStatus.
func (m *SyncedTableContinuousUpdateStatus) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	m.InitialPipelineSyncProgress = vs
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

func (m SyncedTableFailedStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTableFailedStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableFailedStatus
// only implements ToObjectValue() and Type().
func (m SyncedTableFailedStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": m.LastProcessedCommitVersion,
			"timestamp":                     m.Timestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableFailedStatus) Type(ctx context.Context) attr.Type {
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

func (m SyncedTablePipelineProgress) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTablePipelineProgress) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePipelineProgress
// only implements ToObjectValue() and Type().
func (m SyncedTablePipelineProgress) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SyncedTablePipelineProgress) Type(ctx context.Context) attr.Type {
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

func (m SyncedTablePosition) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTablePosition) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"delta_table_sync_info": reflect.TypeOf(DeltaTableSyncInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTablePosition
// only implements ToObjectValue() and Type().
func (m SyncedTablePosition) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"delta_table_sync_info": m.DeltaTableSyncInfo,
			"sync_end_timestamp":    m.SyncEndTimestamp,
			"sync_start_timestamp":  m.SyncStartTimestamp,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTablePosition) Type(ctx context.Context) attr.Type {
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
func (m *SyncedTablePosition) GetDeltaTableSyncInfo(ctx context.Context) (DeltaTableSyncInfo, bool) {
	var e DeltaTableSyncInfo
	if m.DeltaTableSyncInfo.IsNull() || m.DeltaTableSyncInfo.IsUnknown() {
		return e, false
	}
	var v DeltaTableSyncInfo
	d := m.DeltaTableSyncInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeltaTableSyncInfo sets the value of the DeltaTableSyncInfo field in SyncedTablePosition.
func (m *SyncedTablePosition) SetDeltaTableSyncInfo(ctx context.Context, v DeltaTableSyncInfo) {
	vs := v.ToObjectValue(ctx)
	m.DeltaTableSyncInfo = vs
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

func (m SyncedTableProvisioningStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTableProvisioningStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"initial_pipeline_sync_progress": reflect.TypeOf(SyncedTablePipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableProvisioningStatus
// only implements ToObjectValue() and Type().
func (m SyncedTableProvisioningStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"initial_pipeline_sync_progress": m.InitialPipelineSyncProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableProvisioningStatus) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"initial_pipeline_sync_progress": SyncedTablePipelineProgress{}.Type(ctx),
		},
	}
}

// GetInitialPipelineSyncProgress returns the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus as
// a SyncedTablePipelineProgress value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableProvisioningStatus) GetInitialPipelineSyncProgress(ctx context.Context) (SyncedTablePipelineProgress, bool) {
	var e SyncedTablePipelineProgress
	if m.InitialPipelineSyncProgress.IsNull() || m.InitialPipelineSyncProgress.IsUnknown() {
		return e, false
	}
	var v SyncedTablePipelineProgress
	d := m.InitialPipelineSyncProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInitialPipelineSyncProgress sets the value of the InitialPipelineSyncProgress field in SyncedTableProvisioningStatus.
func (m *SyncedTableProvisioningStatus) SetInitialPipelineSyncProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	m.InitialPipelineSyncProgress = vs
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

func (m SyncedTableSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTableSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"new_pipeline_spec":   reflect.TypeOf(NewPipelineSpec{}),
		"primary_key_columns": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableSpec
// only implements ToObjectValue() and Type().
func (m SyncedTableSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SyncedTableSpec) Type(ctx context.Context) attr.Type {
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
func (m *SyncedTableSpec) GetNewPipelineSpec(ctx context.Context) (NewPipelineSpec, bool) {
	var e NewPipelineSpec
	if m.NewPipelineSpec.IsNull() || m.NewPipelineSpec.IsUnknown() {
		return e, false
	}
	var v NewPipelineSpec
	d := m.NewPipelineSpec.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNewPipelineSpec sets the value of the NewPipelineSpec field in SyncedTableSpec.
func (m *SyncedTableSpec) SetNewPipelineSpec(ctx context.Context, v NewPipelineSpec) {
	vs := v.ToObjectValue(ctx)
	m.NewPipelineSpec = vs
}

// GetPrimaryKeyColumns returns the value of the PrimaryKeyColumns field in SyncedTableSpec as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableSpec) GetPrimaryKeyColumns(ctx context.Context) ([]types.String, bool) {
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

// SetPrimaryKeyColumns sets the value of the PrimaryKeyColumns field in SyncedTableSpec.
func (m *SyncedTableSpec) SetPrimaryKeyColumns(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["primary_key_columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PrimaryKeyColumns = types.ListValueMust(t, vs)
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

func (m SyncedTableStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTableStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m SyncedTableStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SyncedTableStatus) Type(ctx context.Context) attr.Type {
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
func (m *SyncedTableStatus) GetContinuousUpdateStatus(ctx context.Context) (SyncedTableContinuousUpdateStatus, bool) {
	var e SyncedTableContinuousUpdateStatus
	if m.ContinuousUpdateStatus.IsNull() || m.ContinuousUpdateStatus.IsUnknown() {
		return e, false
	}
	var v SyncedTableContinuousUpdateStatus
	d := m.ContinuousUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetContinuousUpdateStatus sets the value of the ContinuousUpdateStatus field in SyncedTableStatus.
func (m *SyncedTableStatus) SetContinuousUpdateStatus(ctx context.Context, v SyncedTableContinuousUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	m.ContinuousUpdateStatus = vs
}

// GetFailedStatus returns the value of the FailedStatus field in SyncedTableStatus as
// a SyncedTableFailedStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus) GetFailedStatus(ctx context.Context) (SyncedTableFailedStatus, bool) {
	var e SyncedTableFailedStatus
	if m.FailedStatus.IsNull() || m.FailedStatus.IsUnknown() {
		return e, false
	}
	var v SyncedTableFailedStatus
	d := m.FailedStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFailedStatus sets the value of the FailedStatus field in SyncedTableStatus.
func (m *SyncedTableStatus) SetFailedStatus(ctx context.Context, v SyncedTableFailedStatus) {
	vs := v.ToObjectValue(ctx)
	m.FailedStatus = vs
}

// GetLastSync returns the value of the LastSync field in SyncedTableStatus as
// a SyncedTablePosition value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus) GetLastSync(ctx context.Context) (SyncedTablePosition, bool) {
	var e SyncedTablePosition
	if m.LastSync.IsNull() || m.LastSync.IsUnknown() {
		return e, false
	}
	var v SyncedTablePosition
	d := m.LastSync.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLastSync sets the value of the LastSync field in SyncedTableStatus.
func (m *SyncedTableStatus) SetLastSync(ctx context.Context, v SyncedTablePosition) {
	vs := v.ToObjectValue(ctx)
	m.LastSync = vs
}

// GetProvisioningStatus returns the value of the ProvisioningStatus field in SyncedTableStatus as
// a SyncedTableProvisioningStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus) GetProvisioningStatus(ctx context.Context) (SyncedTableProvisioningStatus, bool) {
	var e SyncedTableProvisioningStatus
	if m.ProvisioningStatus.IsNull() || m.ProvisioningStatus.IsUnknown() {
		return e, false
	}
	var v SyncedTableProvisioningStatus
	d := m.ProvisioningStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetProvisioningStatus sets the value of the ProvisioningStatus field in SyncedTableStatus.
func (m *SyncedTableStatus) SetProvisioningStatus(ctx context.Context, v SyncedTableProvisioningStatus) {
	vs := v.ToObjectValue(ctx)
	m.ProvisioningStatus = vs
}

// GetTriggeredUpdateStatus returns the value of the TriggeredUpdateStatus field in SyncedTableStatus as
// a SyncedTableTriggeredUpdateStatus value.
// If the field is unknown or null, the boolean return value is false.
func (m *SyncedTableStatus) GetTriggeredUpdateStatus(ctx context.Context) (SyncedTableTriggeredUpdateStatus, bool) {
	var e SyncedTableTriggeredUpdateStatus
	if m.TriggeredUpdateStatus.IsNull() || m.TriggeredUpdateStatus.IsUnknown() {
		return e, false
	}
	var v SyncedTableTriggeredUpdateStatus
	d := m.TriggeredUpdateStatus.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTriggeredUpdateStatus sets the value of the TriggeredUpdateStatus field in SyncedTableStatus.
func (m *SyncedTableStatus) SetTriggeredUpdateStatus(ctx context.Context, v SyncedTableTriggeredUpdateStatus) {
	vs := v.ToObjectValue(ctx)
	m.TriggeredUpdateStatus = vs
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

func (m SyncedTableTriggeredUpdateStatus) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SyncedTableTriggeredUpdateStatus) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"triggered_update_progress": reflect.TypeOf(SyncedTablePipelineProgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SyncedTableTriggeredUpdateStatus
// only implements ToObjectValue() and Type().
func (m SyncedTableTriggeredUpdateStatus) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"last_processed_commit_version": m.LastProcessedCommitVersion,
			"timestamp":                     m.Timestamp,
			"triggered_update_progress":     m.TriggeredUpdateProgress,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SyncedTableTriggeredUpdateStatus) Type(ctx context.Context) attr.Type {
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
func (m *SyncedTableTriggeredUpdateStatus) GetTriggeredUpdateProgress(ctx context.Context) (SyncedTablePipelineProgress, bool) {
	var e SyncedTablePipelineProgress
	if m.TriggeredUpdateProgress.IsNull() || m.TriggeredUpdateProgress.IsUnknown() {
		return e, false
	}
	var v SyncedTablePipelineProgress
	d := m.TriggeredUpdateProgress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTriggeredUpdateProgress sets the value of the TriggeredUpdateProgress field in SyncedTableTriggeredUpdateStatus.
func (m *SyncedTableTriggeredUpdateStatus) SetTriggeredUpdateProgress(ctx context.Context, v SyncedTablePipelineProgress) {
	vs := v.ToObjectValue(ctx)
	m.TriggeredUpdateProgress = vs
}

type UpdateDatabaseBranchRequest struct {
	BranchId types.String `tfsdk:"-"`

	DatabaseBranch types.Object `tfsdk:"database_branch"`

	ProjectId types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseBranchRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseBranchRequest) {
	if !from.DatabaseBranch.IsNull() && !from.DatabaseBranch.IsUnknown() {
		if toDatabaseBranch, ok := to.GetDatabaseBranch(ctx); ok {
			if fromDatabaseBranch, ok := from.GetDatabaseBranch(ctx); ok {
				// Recursively sync the fields of DatabaseBranch
				toDatabaseBranch.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseBranch)
				to.SetDatabaseBranch(ctx, toDatabaseBranch)
			}
		}
	}
}

func (to *UpdateDatabaseBranchRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseBranchRequest) {
	if !from.DatabaseBranch.IsNull() && !from.DatabaseBranch.IsUnknown() {
		if toDatabaseBranch, ok := to.GetDatabaseBranch(ctx); ok {
			if fromDatabaseBranch, ok := from.GetDatabaseBranch(ctx); ok {
				toDatabaseBranch.SyncFieldsDuringRead(ctx, fromDatabaseBranch)
				to.SetDatabaseBranch(ctx, toDatabaseBranch)
			}
		}
	}
}

func (m UpdateDatabaseBranchRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_branch"] = attrs["database_branch"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseBranchRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatabaseBranchRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_branch": reflect.TypeOf(DatabaseBranch{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseBranchRequest
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseBranchRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":       m.BranchId,
			"database_branch": m.DatabaseBranch,
			"project_id":      m.ProjectId,
			"update_mask":     m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseBranchRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":       types.StringType,
			"database_branch": DatabaseBranch{}.Type(ctx),
			"project_id":      types.StringType,
			"update_mask":     types.StringType,
		},
	}
}

// GetDatabaseBranch returns the value of the DatabaseBranch field in UpdateDatabaseBranchRequest as
// a DatabaseBranch value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatabaseBranchRequest) GetDatabaseBranch(ctx context.Context) (DatabaseBranch, bool) {
	var e DatabaseBranch
	if m.DatabaseBranch.IsNull() || m.DatabaseBranch.IsUnknown() {
		return e, false
	}
	var v DatabaseBranch
	d := m.DatabaseBranch.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseBranch sets the value of the DatabaseBranch field in UpdateDatabaseBranchRequest.
func (m *UpdateDatabaseBranchRequest) SetDatabaseBranch(ctx context.Context, v DatabaseBranch) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseBranch = vs
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

func (m UpdateDatabaseCatalogRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateDatabaseCatalogRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_catalog": reflect.TypeOf(DatabaseCatalog{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseCatalogRequest
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseCatalogRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_catalog": m.DatabaseCatalog,
			"name":             m.Name,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseCatalogRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateDatabaseCatalogRequest) GetDatabaseCatalog(ctx context.Context) (DatabaseCatalog, bool) {
	var e DatabaseCatalog
	if m.DatabaseCatalog.IsNull() || m.DatabaseCatalog.IsUnknown() {
		return e, false
	}
	var v DatabaseCatalog
	d := m.DatabaseCatalog.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseCatalog sets the value of the DatabaseCatalog field in UpdateDatabaseCatalogRequest.
func (m *UpdateDatabaseCatalogRequest) SetDatabaseCatalog(ctx context.Context, v DatabaseCatalog) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseCatalog = vs
}

type UpdateDatabaseEndpointRequest struct {
	BranchId types.String `tfsdk:"-"`

	DatabaseEndpoint types.Object `tfsdk:"database_endpoint"`

	EndpointId types.String `tfsdk:"-"`

	ProjectId types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseEndpointRequest) {
	if !from.DatabaseEndpoint.IsNull() && !from.DatabaseEndpoint.IsUnknown() {
		if toDatabaseEndpoint, ok := to.GetDatabaseEndpoint(ctx); ok {
			if fromDatabaseEndpoint, ok := from.GetDatabaseEndpoint(ctx); ok {
				// Recursively sync the fields of DatabaseEndpoint
				toDatabaseEndpoint.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseEndpoint)
				to.SetDatabaseEndpoint(ctx, toDatabaseEndpoint)
			}
		}
	}
}

func (to *UpdateDatabaseEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseEndpointRequest) {
	if !from.DatabaseEndpoint.IsNull() && !from.DatabaseEndpoint.IsUnknown() {
		if toDatabaseEndpoint, ok := to.GetDatabaseEndpoint(ctx); ok {
			if fromDatabaseEndpoint, ok := from.GetDatabaseEndpoint(ctx); ok {
				toDatabaseEndpoint.SyncFieldsDuringRead(ctx, fromDatabaseEndpoint)
				to.SetDatabaseEndpoint(ctx, toDatabaseEndpoint)
			}
		}
	}
}

func (m UpdateDatabaseEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_endpoint"] = attrs["database_endpoint"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["branch_id"] = attrs["branch_id"].SetRequired()
	attrs["branch_id"] = attrs["branch_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["endpoint_id"] = attrs["endpoint_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatabaseEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_endpoint": reflect.TypeOf(DatabaseEndpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseEndpointRequest
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"branch_id":         m.BranchId,
			"database_endpoint": m.DatabaseEndpoint,
			"endpoint_id":       m.EndpointId,
			"project_id":        m.ProjectId,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"branch_id":         types.StringType,
			"database_endpoint": DatabaseEndpoint{}.Type(ctx),
			"endpoint_id":       types.StringType,
			"project_id":        types.StringType,
			"update_mask":       types.StringType,
		},
	}
}

// GetDatabaseEndpoint returns the value of the DatabaseEndpoint field in UpdateDatabaseEndpointRequest as
// a DatabaseEndpoint value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatabaseEndpointRequest) GetDatabaseEndpoint(ctx context.Context) (DatabaseEndpoint, bool) {
	var e DatabaseEndpoint
	if m.DatabaseEndpoint.IsNull() || m.DatabaseEndpoint.IsUnknown() {
		return e, false
	}
	var v DatabaseEndpoint
	d := m.DatabaseEndpoint.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseEndpoint sets the value of the DatabaseEndpoint field in UpdateDatabaseEndpointRequest.
func (m *UpdateDatabaseEndpointRequest) SetDatabaseEndpoint(ctx context.Context, v DatabaseEndpoint) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseEndpoint = vs
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

func (m UpdateDatabaseInstanceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateDatabaseInstanceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance": reflect.TypeOf(DatabaseInstance{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseInstanceRequest
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseInstanceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance": m.DatabaseInstance,
			"name":              m.Name,
			"update_mask":       m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseInstanceRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateDatabaseInstanceRequest) GetDatabaseInstance(ctx context.Context) (DatabaseInstance, bool) {
	var e DatabaseInstance
	if m.DatabaseInstance.IsNull() || m.DatabaseInstance.IsUnknown() {
		return e, false
	}
	var v DatabaseInstance
	d := m.DatabaseInstance.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstance sets the value of the DatabaseInstance field in UpdateDatabaseInstanceRequest.
func (m *UpdateDatabaseInstanceRequest) SetDatabaseInstance(ctx context.Context, v DatabaseInstance) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseInstance = vs
}

type UpdateDatabaseInstanceRoleRequest struct {
	DatabaseInstanceRole types.Object `tfsdk:"database_instance_role"`

	InstanceName types.String `tfsdk:"-"`
	// The name of the role. This is the unique identifier for the role in an
	// instance.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseInstanceRoleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseInstanceRoleRequest) {
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

func (to *UpdateDatabaseInstanceRoleRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseInstanceRoleRequest) {
	if !from.DatabaseInstanceRole.IsNull() && !from.DatabaseInstanceRole.IsUnknown() {
		if toDatabaseInstanceRole, ok := to.GetDatabaseInstanceRole(ctx); ok {
			if fromDatabaseInstanceRole, ok := from.GetDatabaseInstanceRole(ctx); ok {
				toDatabaseInstanceRole.SyncFieldsDuringRead(ctx, fromDatabaseInstanceRole)
				to.SetDatabaseInstanceRole(ctx, toDatabaseInstanceRole)
			}
		}
	}
}

func (m UpdateDatabaseInstanceRoleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_instance_role"] = attrs["database_instance_role"].SetRequired()
	attrs["instance_name"] = attrs["instance_name"].SetRequired()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseInstanceRoleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatabaseInstanceRoleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_instance_role": reflect.TypeOf(DatabaseInstanceRole{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseInstanceRoleRequest
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseInstanceRoleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_instance_role": m.DatabaseInstanceRole,
			"instance_name":          m.InstanceName,
			"name":                   m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseInstanceRoleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_instance_role": DatabaseInstanceRole{}.Type(ctx),
			"instance_name":          types.StringType,
			"name":                   types.StringType,
		},
	}
}

// GetDatabaseInstanceRole returns the value of the DatabaseInstanceRole field in UpdateDatabaseInstanceRoleRequest as
// a DatabaseInstanceRole value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatabaseInstanceRoleRequest) GetDatabaseInstanceRole(ctx context.Context) (DatabaseInstanceRole, bool) {
	var e DatabaseInstanceRole
	if m.DatabaseInstanceRole.IsNull() || m.DatabaseInstanceRole.IsUnknown() {
		return e, false
	}
	var v DatabaseInstanceRole
	d := m.DatabaseInstanceRole.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseInstanceRole sets the value of the DatabaseInstanceRole field in UpdateDatabaseInstanceRoleRequest.
func (m *UpdateDatabaseInstanceRoleRequest) SetDatabaseInstanceRole(ctx context.Context, v DatabaseInstanceRole) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseInstanceRole = vs
}

type UpdateDatabaseProjectRequest struct {
	DatabaseProject types.Object `tfsdk:"database_project"`

	ProjectId types.String `tfsdk:"-"`
	// The list of fields to update. If unspecified, all fields will be updated
	// when possible.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateDatabaseProjectRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDatabaseProjectRequest) {
	if !from.DatabaseProject.IsNull() && !from.DatabaseProject.IsUnknown() {
		if toDatabaseProject, ok := to.GetDatabaseProject(ctx); ok {
			if fromDatabaseProject, ok := from.GetDatabaseProject(ctx); ok {
				// Recursively sync the fields of DatabaseProject
				toDatabaseProject.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabaseProject)
				to.SetDatabaseProject(ctx, toDatabaseProject)
			}
		}
	}
}

func (to *UpdateDatabaseProjectRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDatabaseProjectRequest) {
	if !from.DatabaseProject.IsNull() && !from.DatabaseProject.IsUnknown() {
		if toDatabaseProject, ok := to.GetDatabaseProject(ctx); ok {
			if fromDatabaseProject, ok := from.GetDatabaseProject(ctx); ok {
				toDatabaseProject.SyncFieldsDuringRead(ctx, fromDatabaseProject)
				to.SetDatabaseProject(ctx, toDatabaseProject)
			}
		}
	}
}

func (m UpdateDatabaseProjectRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["database_project"] = attrs["database_project"].SetRequired()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["project_id"] = attrs["project_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.RequiresReplace()).(tfschema.AttributeBuilder)
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDatabaseProjectRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDatabaseProjectRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"database_project": reflect.TypeOf(DatabaseProject{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDatabaseProjectRequest
// only implements ToObjectValue() and Type().
func (m UpdateDatabaseProjectRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"database_project": m.DatabaseProject,
			"project_id":       m.ProjectId,
			"update_mask":      m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDatabaseProjectRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"database_project": DatabaseProject{}.Type(ctx),
			"project_id":       types.StringType,
			"update_mask":      types.StringType,
		},
	}
}

// GetDatabaseProject returns the value of the DatabaseProject field in UpdateDatabaseProjectRequest as
// a DatabaseProject value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDatabaseProjectRequest) GetDatabaseProject(ctx context.Context) (DatabaseProject, bool) {
	var e DatabaseProject
	if m.DatabaseProject.IsNull() || m.DatabaseProject.IsUnknown() {
		return e, false
	}
	var v DatabaseProject
	d := m.DatabaseProject.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabaseProject sets the value of the DatabaseProject field in UpdateDatabaseProjectRequest.
func (m *UpdateDatabaseProjectRequest) SetDatabaseProject(ctx context.Context, v DatabaseProject) {
	vs := v.ToObjectValue(ctx)
	m.DatabaseProject = vs
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

func (m UpdateSyncedDatabaseTableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdateSyncedDatabaseTableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"synced_table": reflect.TypeOf(SyncedDatabaseTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSyncedDatabaseTableRequest
// only implements ToObjectValue() and Type().
func (m UpdateSyncedDatabaseTableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":         m.Name,
			"synced_table": m.SyncedTable,
			"update_mask":  m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSyncedDatabaseTableRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateSyncedDatabaseTableRequest) GetSyncedTable(ctx context.Context) (SyncedDatabaseTable, bool) {
	var e SyncedDatabaseTable
	if m.SyncedTable.IsNull() || m.SyncedTable.IsUnknown() {
		return e, false
	}
	var v SyncedDatabaseTable
	d := m.SyncedTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSyncedTable sets the value of the SyncedTable field in UpdateSyncedDatabaseTableRequest.
func (m *UpdateSyncedDatabaseTableRequest) SetSyncedTable(ctx context.Context, v SyncedDatabaseTable) {
	vs := v.ToObjectValue(ctx)
	m.SyncedTable = vs
}
