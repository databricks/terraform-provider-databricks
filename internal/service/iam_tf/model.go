// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package iam_tf

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControlRequest) {
}

func (newState *AccessControlRequest) SyncEffectiveFieldsDuringRead(existingState AccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of AccessControlRequest in the Terraform plugin framework type
// system.
func (a AccessControlRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type AccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControlResponse) {
}

func (newState *AccessControlResponse) SyncEffectiveFieldsDuringRead(existingState AccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(Permission{}),
	}
}

// ToObjectType returns the representation of AccessControlResponse in the Terraform plugin framework type
// system.
func (a AccessControlResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: Permission{}.ToObjectType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ComplexValue struct {
	Display types.String `tfsdk:"display" tf:"optional"`

	Primary types.Bool `tfsdk:"primary" tf:"optional"`

	Ref types.String `tfsdk:"$ref" tf:"optional"`

	Type types.String `tfsdk:"type" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *ComplexValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplexValue) {
}

func (newState *ComplexValue) SyncEffectiveFieldsDuringRead(existingState ComplexValue) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplexValue.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ComplexValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ComplexValue in the Terraform plugin framework type
// system.
func (a ComplexValue) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display": types.StringType,
			"primary": types.BoolType,
			"$ref":    types.StringType,
			"type":    types.StringType,
			"value":   types.StringType,
		},
	}
}

// Delete a group
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteAccountGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountGroupRequest) {
}

func (newState *DeleteAccountGroupRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteAccountGroupRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountGroupRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a service principal
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteAccountServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountServicePrincipalRequest) {
}

func (newState *DeleteAccountServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteAccountServicePrincipalRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountServicePrincipalRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a user
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteAccountUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountUserRequest) {
}

func (newState *DeleteAccountUserRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAccountUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteAccountUserRequest in the Terraform plugin framework type
// system.
func (a DeleteAccountUserRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a group
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteGroupRequest) {
}

func (newState *DeleteGroupRequest) SyncEffectiveFieldsDuringRead(existingState DeleteGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteGroupRequest in the Terraform plugin framework type
// system.
func (a DeleteGroupRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteResponse in the Terraform plugin framework type
// system.
func (a DeleteResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a service principal
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteServicePrincipalRequest) {
}

func (newState *DeleteServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState DeleteServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteServicePrincipalRequest in the Terraform plugin framework type
// system.
func (a DeleteServicePrincipalRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete a user
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteUserRequest) {
}

func (newState *DeleteUserRequest) SyncEffectiveFieldsDuringRead(existingState DeleteUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteUserRequest in the Terraform plugin framework type
// system.
func (a DeleteUserRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID for the account.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspaceAssignmentRequest) {
}

func (newState *DeleteWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspaceAssignmentRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteWorkspaceAssignmentRequest in the Terraform plugin framework type
// system.
func (a DeleteWorkspaceAssignmentRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type DeleteWorkspacePermissionAssignmentResponse struct {
}

func (newState *DeleteWorkspacePermissionAssignmentResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWorkspacePermissionAssignmentResponse) {
}

func (newState *DeleteWorkspacePermissionAssignmentResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWorkspacePermissionAssignmentResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteWorkspacePermissionAssignmentResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteWorkspacePermissionAssignmentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteWorkspacePermissionAssignmentResponse in the Terraform plugin framework type
// system.
func (a DeleteWorkspacePermissionAssignmentResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get group details
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *GetAccountGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountGroupRequest) {
}

func (newState *GetAccountGroupRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetAccountGroupRequest in the Terraform plugin framework type
// system.
func (a GetAccountGroupRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get service principal details
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id types.String `tfsdk:"-"`
}

func (newState *GetAccountServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountServicePrincipalRequest) {
}

func (newState *GetAccountServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetAccountServicePrincipalRequest in the Terraform plugin framework type
// system.
func (a GetAccountServicePrincipalRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get user details
type GetAccountUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Unique ID for a user in the Databricks account.
	Id types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *GetAccountUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAccountUserRequest) {
}

func (newState *GetAccountUserRequest) SyncEffectiveFieldsDuringRead(existingState GetAccountUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetAccountUserRequest in the Terraform plugin framework type
// system.
func (a GetAccountUserRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"id":                 types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	Resource types.String `tfsdk:"-"`
}

func (newState *GetAssignableRolesForResourceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAssignableRolesForResourceRequest) {
}

func (newState *GetAssignableRolesForResourceRequest) SyncEffectiveFieldsDuringRead(existingState GetAssignableRolesForResourceRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAssignableRolesForResourceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAssignableRolesForResourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetAssignableRolesForResourceRequest in the Terraform plugin framework type
// system.
func (a GetAssignableRolesForResourceRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resource": types.StringType,
		},
	}
}

type GetAssignableRolesForResourceResponse struct {
	Roles types.List `tfsdk:"roles" tf:"optional"`
}

func (newState *GetAssignableRolesForResourceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAssignableRolesForResourceResponse) {
}

func (newState *GetAssignableRolesForResourceResponse) SyncEffectiveFieldsDuringRead(existingState GetAssignableRolesForResourceResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAssignableRolesForResourceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAssignableRolesForResourceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role{}),
	}
}

// ToObjectType returns the representation of GetAssignableRolesForResourceResponse in the Terraform plugin framework type
// system.
func (a GetAssignableRolesForResourceResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"roles": basetypes.ListType{
				ElemType: Role{}.ToObjectType(ctx),
			},
		},
	}
}

// Get group details
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *GetGroupRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetGroupRequest) {
}

func (newState *GetGroupRequest) SyncEffectiveFieldsDuringRead(existingState GetGroupRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetGroupRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetGroupRequest in the Terraform plugin framework type
// system.
func (a GetGroupRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetPasswordPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPasswordPermissionLevelsResponse) {
}

func (newState *GetPasswordPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPasswordPermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPasswordPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPasswordPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PasswordPermissionsDescription{}),
	}
}

// ToObjectType returns the representation of GetPasswordPermissionLevelsResponse in the Terraform plugin framework type
// system.
func (a GetPasswordPermissionLevelsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PasswordPermissionsDescription{}.ToObjectType(ctx),
			},
		},
	}
}

// Get object permission levels
type GetPermissionLevelsRequest struct {
	// <needs content>
	RequestObjectId types.String `tfsdk:"-"`
	// <needs content>
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *GetPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionLevelsRequest) {
}

func (newState *GetPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetPermissionLevelsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetPermissionLevelsRequest in the Terraform plugin framework type
// system.
func (a GetPermissionLevelsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionLevelsResponse) {
}

func (newState *GetPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetPermissionLevelsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PermissionsDescription{}),
	}
}

// ToObjectType returns the representation of GetPermissionLevelsResponse in the Terraform plugin framework type
// system.
func (a GetPermissionLevelsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PermissionsDescription{}.ToObjectType(ctx),
			},
		},
	}
}

// Get object permissions
type GetPermissionRequest struct {
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *GetPermissionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPermissionRequest) {
}

func (newState *GetPermissionRequest) SyncEffectiveFieldsDuringRead(existingState GetPermissionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPermissionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPermissionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetPermissionRequest in the Terraform plugin framework type
// system.
func (a GetPermissionRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// Get a rule set
type GetRuleSetRequest struct {
	// Etag used for versioning. The response is at least as fresh as the eTag
	// provided. Etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a rule set from overwriting each
	// other. It is strongly suggested that systems make use of the etag in the
	// read -> modify -> write pattern to perform rule set updates in order to
	// avoid race conditions that is get an etag from a GET rule set request,
	// and pass it with the PUT update request to identify the rule set version
	// you are updating.
	Etag types.String `tfsdk:"-"`
	// The ruleset name associated with the request.
	Name types.String `tfsdk:"-"`
}

func (newState *GetRuleSetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetRuleSetRequest) {
}

func (newState *GetRuleSetRequest) SyncEffectiveFieldsDuringRead(existingState GetRuleSetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetRuleSetRequest in the Terraform plugin framework type
// system.
func (a GetRuleSetRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"name": types.StringType,
		},
	}
}

// Get service principal details
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id types.String `tfsdk:"-"`
}

func (newState *GetServicePrincipalRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServicePrincipalRequest) {
}

func (newState *GetServicePrincipalRequest) SyncEffectiveFieldsDuringRead(existingState GetServicePrincipalRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetServicePrincipalRequest in the Terraform plugin framework type
// system.
func (a GetServicePrincipalRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get user details
type GetUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *GetUserRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetUserRequest) {
}

func (newState *GetUserRequest) SyncEffectiveFieldsDuringRead(existingState GetUserRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetUserRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetUserRequest in the Terraform plugin framework type
// system.
func (a GetUserRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"id":                 types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *GetWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceAssignmentRequest) {
}

func (newState *GetWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceAssignmentRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetWorkspaceAssignmentRequest in the Terraform plugin framework type
// system.
func (a GetWorkspaceAssignmentRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type GrantRule struct {
	// Principals this grant rule applies to.
	Principals types.List `tfsdk:"principals" tf:"optional"`
	// Role that is assigned to the list of principals.
	Role types.String `tfsdk:"role" tf:""`
}

func (newState *GrantRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan GrantRule) {
}

func (newState *GrantRule) SyncEffectiveFieldsDuringRead(existingState GrantRule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GrantRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GrantRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"principals": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of GrantRule in the Terraform plugin framework type
// system.
func (a GrantRule) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"principals": basetypes.ListType{
				ElemType: types.StringType,
			},
			"role": types.StringType,
		},
	}
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName types.String `tfsdk:"displayName" tf:"optional"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements" tf:"optional"`

	ExternalId types.String `tfsdk:"externalId" tf:"optional"`

	Groups types.List `tfsdk:"groups" tf:"optional"`
	// Databricks group ID
	Id types.String `tfsdk:"id" tf:"optional"`

	Members types.List `tfsdk:"members" tf:"optional"`
	// Container for the group identifier. Workspace local versus account.
	Meta types.List `tfsdk:"meta" tf:"optional,object"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles" tf:"optional"`
	// The schema of the group.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *Group) SyncEffectiveFieldsDuringCreateOrUpdate(plan Group) {
}

func (newState *Group) SyncEffectiveFieldsDuringRead(existingState Group) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Group.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Group) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"members":      reflect.TypeOf(ComplexValue{}),
		"meta":         reflect.TypeOf(ResourceMeta{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of Group in the Terraform plugin framework type
// system.
func (a Group) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"displayName": types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"id": types.StringType,
			"members": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"meta": basetypes.ListType{
				ElemType: ResourceMeta{}.ToObjectType(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// List group details
type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListAccountGroupsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountGroupsRequest) {
}

func (newState *ListAccountGroupsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountGroupsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListAccountGroupsRequest in the Terraform plugin framework type
// system.
func (a ListAccountGroupsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List service principals
type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListAccountServicePrincipalsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountServicePrincipalsRequest) {
}

func (newState *ListAccountServicePrincipalsRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountServicePrincipalsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListAccountServicePrincipalsRequest in the Terraform plugin framework type
// system.
func (a ListAccountServicePrincipalsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List users
type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page. Default is 10000.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListAccountUsersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAccountUsersRequest) {
}

func (newState *ListAccountUsersRequest) SyncEffectiveFieldsDuringRead(existingState ListAccountUsersRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListAccountUsersRequest in the Terraform plugin framework type
// system.
func (a ListAccountUsersRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List group details
type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListGroupsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGroupsRequest) {
}

func (newState *ListGroupsRequest) SyncEffectiveFieldsDuringRead(existingState ListGroupsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListGroupsRequest in the Terraform plugin framework type
// system.
func (a ListGroupsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage" tf:"optional"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources" tf:"optional"`
	// The schema of the service principal.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex" tf:"optional"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults" tf:"optional"`
}

func (newState *ListGroupsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListGroupsResponse) {
}

func (newState *ListGroupsResponse) SyncEffectiveFieldsDuringRead(existingState ListGroupsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListGroupsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(Group{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ListGroupsResponse in the Terraform plugin framework type
// system.
func (a ListGroupsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: Group{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage" tf:"optional"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources" tf:"optional"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex" tf:"optional"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults" tf:"optional"`
}

func (newState *ListServicePrincipalResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalResponse) {
}

func (newState *ListServicePrincipalResponse) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(ServicePrincipal{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ListServicePrincipalResponse in the Terraform plugin framework type
// system.
func (a ListServicePrincipalResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: ServicePrincipal{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// List service principals
type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListServicePrincipalsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalsRequest) {
}

func (newState *ListServicePrincipalsRequest) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListServicePrincipalsRequest in the Terraform plugin framework type
// system.
func (a ListServicePrincipalsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

// List users
type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes types.String `tfsdk:"-"`
	// Desired number of results per page.
	Count types.Int64 `tfsdk:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes types.String `tfsdk:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter types.String `tfsdk:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy types.String `tfsdk:"-"`
	// The order to sort the results.
	SortOrder types.String `tfsdk:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex types.Int64 `tfsdk:"-"`
}

func (newState *ListUsersRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUsersRequest) {
}

func (newState *ListUsersRequest) SyncEffectiveFieldsDuringRead(existingState ListUsersRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListUsersRequest in the Terraform plugin framework type
// system.
func (a ListUsersRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attributes":         types.StringType,
			"count":              types.Int64Type,
			"excludedAttributes": types.StringType,
			"filter":             types.StringType,
			"sortBy":             types.StringType,
			"sortOrder":          types.StringType,
			"startIndex":         types.Int64Type,
		},
	}
}

type ListUsersResponse struct {
	// Total results returned in the response.
	ItemsPerPage types.Int64 `tfsdk:"itemsPerPage" tf:"optional"`
	// User objects returned in the response.
	Resources types.List `tfsdk:"Resources" tf:"optional"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex types.Int64 `tfsdk:"startIndex" tf:"optional"`
	// Total results that match the request filters.
	TotalResults types.Int64 `tfsdk:"totalResults" tf:"optional"`
}

func (newState *ListUsersResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListUsersResponse) {
}

func (newState *ListUsersResponse) SyncEffectiveFieldsDuringRead(existingState ListUsersResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListUsersResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListUsersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(User{}),
		"schemas":   reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ListUsersResponse in the Terraform plugin framework type
// system.
func (a ListUsersResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: User{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"startIndex":   types.Int64Type,
			"totalResults": types.Int64Type,
		},
	}
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *ListWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWorkspaceAssignmentRequest) {
}

func (newState *ListWorkspaceAssignmentRequest) SyncEffectiveFieldsDuringRead(existingState ListWorkspaceAssignmentRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListWorkspaceAssignmentRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListWorkspaceAssignmentRequest in the Terraform plugin framework type
// system.
func (a ListWorkspaceAssignmentRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type MigratePermissionsRequest struct {
	// The name of the workspace group that permissions will be migrated from.
	FromWorkspaceGroupName types.String `tfsdk:"from_workspace_group_name" tf:""`
	// The maximum number of permissions that will be migrated.
	Size types.Int64 `tfsdk:"size" tf:"optional"`
	// The name of the account group that permissions will be migrated to.
	ToAccountGroupName types.String `tfsdk:"to_account_group_name" tf:""`
	// WorkspaceId of the associated workspace where the permission migration
	// will occur.
	WorkspaceId types.Int64 `tfsdk:"workspace_id" tf:""`
}

func (newState *MigratePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigratePermissionsRequest) {
}

func (newState *MigratePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState MigratePermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of MigratePermissionsRequest in the Terraform plugin framework type
// system.
func (a MigratePermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"from_workspace_group_name": types.StringType,
			"size":                      types.Int64Type,
			"to_account_group_name":     types.StringType,
			"workspace_id":              types.Int64Type,
		},
	}
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	PermissionsMigrated types.Int64 `tfsdk:"permissions_migrated" tf:"optional"`
}

func (newState *MigratePermissionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigratePermissionsResponse) {
}

func (newState *MigratePermissionsResponse) SyncEffectiveFieldsDuringRead(existingState MigratePermissionsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigratePermissionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigratePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of MigratePermissionsResponse in the Terraform plugin framework type
// system.
func (a MigratePermissionsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions_migrated": types.Int64Type,
		},
	}
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName types.String `tfsdk:"familyName" tf:"optional"`
	// Given name of the Databricks user.
	GivenName types.String `tfsdk:"givenName" tf:"optional"`
}

func (newState *Name) SyncEffectiveFieldsDuringCreateOrUpdate(plan Name) {
}

func (newState *Name) SyncEffectiveFieldsDuringRead(existingState Name) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Name.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Name) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Name in the Terraform plugin framework type
// system.
func (a Name) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"familyName": types.StringType,
			"givenName":  types.StringType,
		},
	}
}

type ObjectPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ObjectPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ObjectPermissions) {
}

func (newState *ObjectPermissions) SyncEffectiveFieldsDuringRead(existingState ObjectPermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ObjectPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlResponse{}),
	}
}

// ToObjectType returns the representation of ObjectPermissions in the Terraform plugin framework type
// system.
func (a ObjectPermissions) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlResponse{}.ToObjectType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type PartialUpdate struct {
	// Unique ID for a user in the Databricks workspace.
	Id types.String `tfsdk:"-"`

	Operations types.List `tfsdk:"Operations" tf:"optional"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *PartialUpdate) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartialUpdate) {
}

func (newState *PartialUpdate) SyncEffectiveFieldsDuringRead(existingState PartialUpdate) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartialUpdate.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartialUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Operations": reflect.TypeOf(Patch{}),
		"schemas":    reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of PartialUpdate in the Terraform plugin framework type
// system.
func (a PartialUpdate) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"Operations": basetypes.ListType{
				ElemType: Patch{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type PasswordAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PasswordAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordAccessControlRequest) {
}

func (newState *PasswordAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState PasswordAccessControlRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PasswordAccessControlRequest in the Terraform plugin framework type
// system.
func (a PasswordAccessControlRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PasswordAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PasswordAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordAccessControlResponse) {
}

func (newState *PasswordAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState PasswordAccessControlResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PasswordPermission{}),
	}
}

// ToObjectType returns the representation of PasswordAccessControlResponse in the Terraform plugin framework type
// system.
func (a PasswordAccessControlResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PasswordPermission{}.ToObjectType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type PasswordPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PasswordPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermission) {
}

func (newState *PasswordPermission) SyncEffectiveFieldsDuringRead(existingState PasswordPermission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of PasswordPermission in the Terraform plugin framework type
// system.
func (a PasswordPermission) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

type PasswordPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *PasswordPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissions) {
}

func (newState *PasswordPermissions) SyncEffectiveFieldsDuringRead(existingState PasswordPermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlResponse{}),
	}
}

// ToObjectType returns the representation of PasswordPermissions in the Terraform plugin framework type
// system.
func (a PasswordPermissions) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlResponse{}.ToObjectType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type PasswordPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PasswordPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissionsDescription) {
}

func (newState *PasswordPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PasswordPermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PasswordPermissionsDescription in the Terraform plugin framework type
// system.
func (a PasswordPermissionsDescription) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PasswordPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
}

func (newState *PasswordPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PasswordPermissionsRequest) {
}

func (newState *PasswordPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState PasswordPermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PasswordPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PasswordPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlRequest{}),
	}
}

// ToObjectType returns the representation of PasswordPermissionsRequest in the Terraform plugin framework type
// system.
func (a PasswordPermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlRequest{}.ToObjectType(ctx),
			},
		},
	}
}

type Patch struct {
	// Type of patch operation.
	Op types.String `tfsdk:"op" tf:"optional"`
	// Selection of patch operation
	Path types.String `tfsdk:"path" tf:"optional"`
	// Value to modify
	Value any `tfsdk:"value" tf:"optional"`
}

func (newState *Patch) SyncEffectiveFieldsDuringCreateOrUpdate(plan Patch) {
}

func (newState *Patch) SyncEffectiveFieldsDuringRead(existingState Patch) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Patch.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Patch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Patch in the Terraform plugin framework type
// system.
func (a Patch) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"op":    types.StringType,
			"path":  types.StringType,
			"value": types.ObjectType{},
		},
	}
}

type PatchResponse struct {
}

func (newState *PatchResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PatchResponse) {
}

func (newState *PatchResponse) SyncEffectiveFieldsDuringRead(existingState PatchResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PatchResponse in the Terraform plugin framework type
// system.
func (a PatchResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type Permission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *Permission) SyncEffectiveFieldsDuringCreateOrUpdate(plan Permission) {
}

func (newState *Permission) SyncEffectiveFieldsDuringRead(existingState Permission) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Permission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Permission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of Permission in the Terraform plugin framework type
// system.
func (a Permission) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"inherited": types.BoolType,
			"inherited_from_object": basetypes.ListType{
				ElemType: types.StringType,
			},
			"permission_level": types.StringType,
		},
	}
}

// The output format for existing workspace PermissionAssignment records, which
// contains some info for user consumption.
type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error types.String `tfsdk:"error" tf:"optional"`
	// The permissions level of the principal.
	Permissions types.List `tfsdk:"permissions" tf:"optional"`
	// Information about the principal assigned to the workspace.
	Principal types.List `tfsdk:"principal" tf:"optional,object"`
}

func (newState *PermissionAssignment) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionAssignment) {
}

func (newState *PermissionAssignment) SyncEffectiveFieldsDuringRead(existingState PermissionAssignment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
		"principal":   reflect.TypeOf(PrincipalOutput{}),
	}
}

// ToObjectType returns the representation of PermissionAssignment in the Terraform plugin framework type
// system.
func (a PermissionAssignment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal": basetypes.ListType{
				ElemType: PrincipalOutput{}.ToObjectType(ctx),
			},
		},
	}
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments types.List `tfsdk:"permission_assignments" tf:"optional"`
}

func (newState *PermissionAssignments) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionAssignments) {
}

func (newState *PermissionAssignments) SyncEffectiveFieldsDuringRead(existingState PermissionAssignments) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_assignments": reflect.TypeOf(PermissionAssignment{}),
	}
}

// ToObjectType returns the representation of PermissionAssignments in the Terraform plugin framework type
// system.
func (a PermissionAssignments) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_assignments": basetypes.ListType{
				ElemType: PermissionAssignment{}.ToObjectType(ctx),
			},
		},
	}
}

type PermissionOutput struct {
	// The results of a permissions query.
	Description types.String `tfsdk:"description" tf:"optional"`

	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PermissionOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionOutput) {
}

func (newState *PermissionOutput) SyncEffectiveFieldsDuringRead(existingState PermissionOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PermissionOutput in the Terraform plugin framework type
// system.
func (a PermissionOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *PermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsDescription) {
}

func (newState *PermissionsDescription) SyncEffectiveFieldsDuringRead(existingState PermissionsDescription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PermissionsDescription in the Terraform plugin framework type
// system.
func (a PermissionsDescription) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type PermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The id of the request object.
	RequestObjectId types.String `tfsdk:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType types.String `tfsdk:"-"`
}

func (newState *PermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PermissionsRequest) {
}

func (newState *PermissionsRequest) SyncEffectiveFieldsDuringRead(existingState PermissionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest{}),
	}
}

// ToObjectType returns the representation of PermissionsRequest in the Terraform plugin framework type
// system.
func (a PermissionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest{}.ToObjectType(ctx),
			},
			"request_object_id":   types.StringType,
			"request_object_type": types.StringType,
		},
	}
}

// Information about the principal assigned to the workspace.
type PrincipalOutput struct {
	// The display name of the principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The group name of the group. Present only if the principal is a group.
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// The unique, opaque id of the principal.
	PrincipalId types.Int64 `tfsdk:"principal_id" tf:"optional"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// The username of the user. Present only if the principal is a user.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *PrincipalOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PrincipalOutput) {
}

func (newState *PrincipalOutput) SyncEffectiveFieldsDuringRead(existingState PrincipalOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PrincipalOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PrincipalOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PrincipalOutput in the Terraform plugin framework type
// system.
func (a PrincipalOutput) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"principal_id":           types.Int64Type,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType types.String `tfsdk:"resourceType" tf:"optional"`
}

func (newState *ResourceMeta) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResourceMeta) {
}

func (newState *ResourceMeta) SyncEffectiveFieldsDuringRead(existingState ResourceMeta) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ResourceMeta.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ResourceMeta) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ResourceMeta in the Terraform plugin framework type
// system.
func (a ResourceMeta) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"resourceType": types.StringType,
		},
	}
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *Role) SyncEffectiveFieldsDuringCreateOrUpdate(plan Role) {
}

func (newState *Role) SyncEffectiveFieldsDuringRead(existingState Role) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Role.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Role) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Role in the Terraform plugin framework type
// system.
func (a Role) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned.
	Etag types.String `tfsdk:"etag" tf:"optional"`

	GrantRules types.List `tfsdk:"grant_rules" tf:"optional"`
	// Name of the rule set.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *RuleSetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RuleSetResponse) {
}

func (newState *RuleSetResponse) SyncEffectiveFieldsDuringRead(existingState RuleSetResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSetResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RuleSetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

// ToObjectType returns the representation of RuleSetResponse in the Terraform plugin framework type
// system.
func (a RuleSetResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.ToObjectType(ctx),
			},
			"name": types.StringType,
		},
	}
}

type RuleSetUpdateRequest struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	Etag types.String `tfsdk:"etag" tf:""`

	GrantRules types.List `tfsdk:"grant_rules" tf:"optional"`
	// Name of the rule set.
	Name types.String `tfsdk:"name" tf:""`
}

func (newState *RuleSetUpdateRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RuleSetUpdateRequest) {
}

func (newState *RuleSetUpdateRequest) SyncEffectiveFieldsDuringRead(existingState RuleSetUpdateRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RuleSetUpdateRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RuleSetUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

// ToObjectType returns the representation of RuleSetUpdateRequest in the Terraform plugin framework type
// system.
func (a RuleSetUpdateRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.ToObjectType(ctx),
			},
			"name": types.StringType,
		},
	}
}

type ServicePrincipal struct {
	// If this user is active
	Active types.Bool `tfsdk:"active" tf:"optional"`
	// UUID relating to the service principal
	ApplicationId types.String `tfsdk:"applicationId" tf:"optional"`
	// String that represents a concatenation of given and family names.
	DisplayName types.String `tfsdk:"displayName" tf:"optional"`
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements" tf:"optional"`

	ExternalId types.String `tfsdk:"externalId" tf:"optional"`

	Groups types.List `tfsdk:"groups" tf:"optional"`
	// Databricks service principal ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles" tf:"optional"`
	// The schema of the List response.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
}

func (newState *ServicePrincipal) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServicePrincipal) {
}

func (newState *ServicePrincipal) SyncEffectiveFieldsDuringRead(existingState ServicePrincipal) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServicePrincipal.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of ServicePrincipal in the Terraform plugin framework type
// system.
func (a ServicePrincipal) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":        types.BoolType,
			"applicationId": types.StringType,
			"displayName":   types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"id": types.StringType,
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of UpdateResponse in the Terraform plugin framework type
// system.
func (a UpdateResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name types.String `tfsdk:"name" tf:""`

	RuleSet types.List `tfsdk:"rule_set" tf:"object"`
}

func (newState *UpdateRuleSetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRuleSetRequest) {
}

func (newState *UpdateRuleSetRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRuleSetRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRuleSetRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule_set": reflect.TypeOf(RuleSetUpdateRequest{}),
	}
}

// ToObjectType returns the representation of UpdateRuleSetRequest in the Terraform plugin framework type
// system.
func (a UpdateRuleSetRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rule_set": basetypes.ListType{
				ElemType: RuleSetUpdateRequest{}.ToObjectType(ctx),
			},
		},
	}
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace. Valid values
	// are "USER" and "ADMIN" (case-sensitive). If both "USER" and "ADMIN" are
	// provided, "ADMIN" takes precedence. Other values will be ignored. Note
	// that excluding this field, or providing unsupported values, will have the
	// same effect as providing an empty list, which will result in the deletion
	// of all permissions for the principal.
	Permissions types.List `tfsdk:"permissions" tf:"optional"`
	// The ID of the user, service principal, or group.
	PrincipalId types.Int64 `tfsdk:"-"`
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (newState *UpdateWorkspaceAssignments) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateWorkspaceAssignments) {
}

func (newState *UpdateWorkspaceAssignments) SyncEffectiveFieldsDuringRead(existingState UpdateWorkspaceAssignments) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceAssignments.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of UpdateWorkspaceAssignments in the Terraform plugin framework type
// system.
func (a UpdateWorkspaceAssignments) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal_id": types.Int64Type,
			"workspace_id": types.Int64Type,
		},
	}
}

type User struct {
	// If this user is active
	Active types.Bool `tfsdk:"active" tf:"optional"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	DisplayName types.String `tfsdk:"displayName" tf:"optional"`
	// All the emails associated with the Databricks user.
	Emails types.List `tfsdk:"emails" tf:"optional"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements types.List `tfsdk:"entitlements" tf:"optional"`
	// External ID is not currently supported. It is reserved for future use.
	ExternalId types.String `tfsdk:"externalId" tf:"optional"`

	Groups types.List `tfsdk:"groups" tf:"optional"`
	// Databricks user ID. This is automatically set by Databricks. Any value
	// provided by the client will be ignored.
	Id types.String `tfsdk:"id" tf:"optional"`

	Name types.List `tfsdk:"name" tf:"optional,object"`
	// Corresponds to AWS instance profile/arn role.
	Roles types.List `tfsdk:"roles" tf:"optional"`
	// The schema of the user.
	Schemas types.List `tfsdk:"schemas" tf:"optional"`
	// Email address of the Databricks user.
	UserName types.String `tfsdk:"userName" tf:"optional"`
}

func (newState *User) SyncEffectiveFieldsDuringCreateOrUpdate(plan User) {
}

func (newState *User) SyncEffectiveFieldsDuringRead(existingState User) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in User.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"emails":       reflect.TypeOf(ComplexValue{}),
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"name":         reflect.TypeOf(Name{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.String{}),
	}
}

// ToObjectType returns the representation of User in the Terraform plugin framework type
// system.
func (a User) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":      types.BoolType,
			"displayName": types.StringType,
			"emails": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"id": types.StringType,
			"name": basetypes.ListType{
				ElemType: Name{}.ToObjectType(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.ToObjectType(ctx),
			},
			"schemas": basetypes.ListType{
				ElemType: types.StringType,
			},
			"userName": types.StringType,
		},
	}
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions types.List `tfsdk:"permissions" tf:"optional"`
}

func (newState *WorkspacePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspacePermissions) {
}

func (newState *WorkspacePermissions) SyncEffectiveFieldsDuringRead(existingState WorkspacePermissions) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspacePermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a WorkspacePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(PermissionOutput{}),
	}
}

// ToObjectType returns the representation of WorkspacePermissions in the Terraform plugin framework type
// system.
func (a WorkspacePermissions) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: PermissionOutput{}.ToObjectType(ctx),
			},
		},
	}
}
