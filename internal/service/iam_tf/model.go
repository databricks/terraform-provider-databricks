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

func (a AccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a AccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a AccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(Permission{}),
	}
}

func (a AccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: Permission{}.ToAttrType(ctx),
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

func (a ComplexValue) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ComplexValue) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteAccountGroupRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteAccountServicePrincipalRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteAccountUserRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteGroupRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteServicePrincipalRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteUserRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteWorkspaceAssignmentRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a DeleteWorkspacePermissionAssignmentResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a DeleteWorkspacePermissionAssignmentResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetAccountGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAccountGroupRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetAccountServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAccountServicePrincipalRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetAccountUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAccountUserRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetAssignableRolesForResourceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetAssignableRolesForResourceRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetAssignableRolesForResourceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"roles": reflect.TypeOf(Role{}),
	}
}

func (a GetAssignableRolesForResourceResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"roles": basetypes.ListType{
				ElemType: Role{}.ToAttrType(ctx),
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

func (a GetGroupRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetGroupRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetPasswordPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PasswordPermissionsDescription{}),
	}
}

func (a GetPasswordPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PasswordPermissionsDescription{}.ToAttrType(ctx),
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

func (a GetPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(PermissionsDescription{}),
	}
}

func (a GetPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: PermissionsDescription{}.ToAttrType(ctx),
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

func (a GetPermissionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetPermissionRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetRuleSetRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetServicePrincipalRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetServicePrincipalRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetUserRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetUserRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GetWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a GetWorkspaceAssignmentRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a GrantRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"principals": reflect.TypeOf(types.StringType),
	}
}

func (a GrantRule) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Group) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"members":      reflect.TypeOf(ComplexValue{}),
		"meta":         reflect.TypeOf(ResourceMeta{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.StringType),
	}
}

func (a Group) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"displayName": types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"id": types.StringType,
			"members": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"meta": basetypes.ListType{
				ElemType: ResourceMeta{}.ToAttrType(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
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

func (a ListAccountGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListAccountGroupsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListAccountServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListAccountServicePrincipalsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListAccountUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListAccountUsersRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListGroupsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListGroupsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListGroupsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(Group{}),
		"schemas":   reflect.TypeOf(types.StringType),
	}
}

func (a ListGroupsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: Group{}.ToAttrType(ctx),
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

func (a ListServicePrincipalResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(ServicePrincipal{}),
		"schemas":   reflect.TypeOf(types.StringType),
	}
}

func (a ListServicePrincipalResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: ServicePrincipal{}.ToAttrType(ctx),
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

func (a ListServicePrincipalsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListServicePrincipalsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListUsersRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListUsersRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ListUsersResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Resources": reflect.TypeOf(User{}),
		"schemas":   reflect.TypeOf(types.StringType),
	}
}

func (a ListUsersResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"itemsPerPage": types.Int64Type,
			"Resources": basetypes.ListType{
				ElemType: User{}.ToAttrType(ctx),
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

func (a ListWorkspaceAssignmentRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ListWorkspaceAssignmentRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a MigratePermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a MigratePermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a MigratePermissionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a MigratePermissionsResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Name) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Name) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ObjectPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlResponse{}),
	}
}

func (a ObjectPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlResponse{}.ToAttrType(ctx),
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

func (a PartialUpdate) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"Operations": reflect.TypeOf(Patch{}),
		"schemas":    reflect.TypeOf(types.StringType),
	}
}

func (a PartialUpdate) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
			"Operations": basetypes.ListType{
				ElemType: Patch{}.ToAttrType(ctx),
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

func (a PasswordAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PasswordAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PasswordAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(PasswordPermission{}),
	}
}

func (a PasswordAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: PasswordPermission{}.ToAttrType(ctx),
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

func (a PasswordPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.StringType),
	}
}

func (a PasswordPermission) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PasswordPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlResponse{}),
	}
}

func (a PasswordPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlResponse{}.ToAttrType(ctx),
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

func (a PasswordPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PasswordPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PasswordPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(PasswordAccessControlRequest{}),
	}
}

func (a PasswordPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: PasswordAccessControlRequest{}.ToAttrType(ctx),
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

func (a Patch) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Patch) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PatchResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PatchResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Permission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.StringType),
	}
}

func (a Permission) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PermissionAssignment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.StringType),
		"principal":   reflect.TypeOf(PrincipalOutput{}),
	}
}

func (a PermissionAssignment) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"permissions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"principal": basetypes.ListType{
				ElemType: PrincipalOutput{}.ToAttrType(ctx),
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

func (a PermissionAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_assignments": reflect.TypeOf(PermissionAssignment{}),
	}
}

func (a PermissionAssignments) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_assignments": basetypes.ListType{
				ElemType: PermissionAssignment{}.ToAttrType(ctx),
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

func (a PermissionOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PermissionOutput) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a PermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(AccessControlRequest{}),
	}
}

func (a PermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: AccessControlRequest{}.ToAttrType(ctx),
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

func (a PrincipalOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a PrincipalOutput) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a ResourceMeta) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a ResourceMeta) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a Role) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a Role) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a RuleSetResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

func (a RuleSetResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.ToAttrType(ctx),
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

func (a RuleSetUpdateRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"grant_rules": reflect.TypeOf(GrantRule{}),
	}
}

func (a RuleSetUpdateRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"grant_rules": basetypes.ListType{
				ElemType: GrantRule{}.ToAttrType(ctx),
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

func (a ServicePrincipal) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.StringType),
	}
}

func (a ServicePrincipal) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":        types.BoolType,
			"applicationId": types.StringType,
			"displayName":   types.StringType,
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"id": types.StringType,
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
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

func (a UpdateResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (a UpdateResponse) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a UpdateRuleSetRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rule_set": reflect.TypeOf(RuleSetUpdateRequest{}),
	}
}

func (a UpdateRuleSetRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rule_set": basetypes.ListType{
				ElemType: RuleSetUpdateRequest{}.ToAttrType(ctx),
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

func (a UpdateWorkspaceAssignments) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(types.StringType),
	}
}

func (a UpdateWorkspaceAssignments) ToAttrType(ctx context.Context) types.ObjectType {
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

func (a User) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"emails":       reflect.TypeOf(ComplexValue{}),
		"entitlements": reflect.TypeOf(ComplexValue{}),
		"groups":       reflect.TypeOf(ComplexValue{}),
		"name":         reflect.TypeOf(Name{}),
		"roles":        reflect.TypeOf(ComplexValue{}),
		"schemas":      reflect.TypeOf(types.StringType),
	}
}

func (a User) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"active":      types.BoolType,
			"displayName": types.StringType,
			"emails": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"entitlements": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"externalId": types.StringType,
			"groups": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
			},
			"id": types.StringType,
			"name": basetypes.ListType{
				ElemType: Name{}.ToAttrType(ctx),
			},
			"roles": basetypes.ListType{
				ElemType: ComplexValue{}.ToAttrType(ctx),
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

func (a WorkspacePermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permissions": reflect.TypeOf(PermissionOutput{}),
	}
}

func (a WorkspacePermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permissions": basetypes.ListType{
				ElemType: PermissionOutput{}.ToAttrType(ctx),
			},
		},
	}
}
