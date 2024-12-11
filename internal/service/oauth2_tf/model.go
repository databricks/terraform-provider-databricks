// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package oauth2_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential types.Bool `tfsdk:"confidential" tf:"optional"`
	// Name of the custom OAuth app
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of OAuth redirect urls
	RedirectUrls types.List `tfsdk:"redirect_urls" tf:"optional"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes types.List `tfsdk:"scopes" tf:"optional"`
	// Token access policy
	TokenAccessPolicy types.List `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *CreateCustomAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomAppIntegration) {
}

func (newState *CreateCustomAppIntegration) SyncEffectiveFieldsDuringRead(existingState CreateCustomAppIntegration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":       reflect.TypeOf(types.String{}),
		"scopes":              reflect.TypeOf(types.String{}),
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomAppIntegration
// only implements ToObjectValue() and Type().
func (o CreateCustomAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"confidential":        o.Confidential,
			"name":                o.Name,
			"redirect_urls":       o.RedirectUrls,
			"scopes":              o.Scopes,
			"token_access_policy": o.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomAppIntegration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"confidential": types.BoolType,
			"name":         types.StringType,
			"redirect_urls": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_access_policy": basetypes.ListType{ElemType: TokenAccessPolicy{}.Type(ctx)},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in CreateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomAppIntegration) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if o.RedirectUrls.IsNull() || o.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in CreateCustomAppIntegration.
func (o *CreateCustomAppIntegration) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in CreateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomAppIntegration) GetScopes(ctx context.Context) ([]types.String, bool) {
	if o.Scopes.IsNull() || o.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in CreateCustomAppIntegration.
func (o *CreateCustomAppIntegration) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in CreateCustomAppIntegration as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateCustomAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if o.TokenAccessPolicy.IsNull() || o.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy
	d := o.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreateCustomAppIntegration.
func (o *CreateCustomAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	o.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	ClientId types.String `tfsdk:"client_id" tf:"optional"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret types.String `tfsdk:"client_secret" tf:"optional"`
	// Unique integration id for the custom OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
}

func (newState *CreateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomAppIntegrationOutput) {
}

func (newState *CreateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState CreateCustomAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o CreateCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_id":      o.ClientId,
			"client_secret":  o.ClientSecret,
			"integration_id": o.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client_id":      types.StringType,
			"client_secret":  types.StringType,
			"integration_id": types.StringType,
		},
	}
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId types.String `tfsdk:"app_id" tf:"optional"`
	// Token access policy
	TokenAccessPolicy types.List `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *CreatePublishedAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePublishedAppIntegration) {
}

func (newState *CreatePublishedAppIntegration) SyncEffectiveFieldsDuringRead(existingState CreatePublishedAppIntegration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePublishedAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePublishedAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePublishedAppIntegration
// only implements ToObjectValue() and Type().
func (o CreatePublishedAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":              o.AppId,
			"token_access_policy": o.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePublishedAppIntegration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_id":              types.StringType,
			"token_access_policy": basetypes.ListType{ElemType: TokenAccessPolicy{}.Type(ctx)},
		},
	}
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in CreatePublishedAppIntegration as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePublishedAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if o.TokenAccessPolicy.IsNull() || o.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy
	d := o.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreatePublishedAppIntegration.
func (o *CreatePublishedAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	o.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
}

func (newState *CreatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePublishedAppIntegrationOutput) {
}

func (newState *CreatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState CreatePublishedAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o CreatePublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": o.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (newState *CreateServicePrincipalSecretRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServicePrincipalSecretRequest) {
}

func (newState *CreateServicePrincipalSecretRequest) SyncEffectiveFieldsDuringRead(existingState CreateServicePrincipalSecretRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalSecretRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateServicePrincipalSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalSecretRequest
// only implements ToObjectValue() and Type().
func (o CreateServicePrincipalSecretRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateServicePrincipalSecretRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"service_principal_id": types.Int64Type,
		},
	}
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// ID of the secret
	Id types.String `tfsdk:"id" tf:"optional"`
	// Secret Value
	Secret types.String `tfsdk:"secret" tf:"optional"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash" tf:"optional"`
	// Status of the secret
	Status types.String `tfsdk:"status" tf:"optional"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *CreateServicePrincipalSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServicePrincipalSecretResponse) {
}

func (newState *CreateServicePrincipalSecretResponse) SyncEffectiveFieldsDuringRead(existingState CreateServicePrincipalSecretResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalSecretResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateServicePrincipalSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalSecretResponse
// only implements ToObjectValue() and Type().
func (o CreateServicePrincipalSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": o.CreateTime,
			"id":          o.Id,
			"secret":      o.Secret,
			"secret_hash": o.SecretHash,
			"status":      o.Status,
			"update_time": o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateServicePrincipalSecretResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"id":          types.StringType,
			"secret":      types.StringType,
			"secret_hash": types.StringType,
			"status":      types.StringType,
			"update_time": types.StringType,
		},
	}
}

type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails types.String `tfsdk:"authorization_details" tf:"optional"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl types.String `tfsdk:"endpoint_url" tf:"optional"`
}

func (newState *DataPlaneInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneInfo) {
}

func (newState *DataPlaneInfo) SyncEffectiveFieldsDuringRead(existingState DataPlaneInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataPlaneInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataPlaneInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneInfo
// only implements ToObjectValue() and Type().
func (o DataPlaneInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authorization_details": o.AuthorizationDetails,
			"endpoint_url":          o.EndpointUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataPlaneInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authorization_details": types.StringType,
			"endpoint_url":          types.StringType,
		},
	}
}

type DeleteCustomAppIntegrationOutput struct {
}

func (newState *DeleteCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCustomAppIntegrationOutput) {
}

func (newState *DeleteCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState DeleteCustomAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o DeleteCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *DeleteCustomAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCustomAppIntegrationRequest) {
}

func (newState *DeleteCustomAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteCustomAppIntegrationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteCustomAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (o DeleteCustomAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": o.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteCustomAppIntegrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type DeletePublishedAppIntegrationOutput struct {
}

func (newState *DeletePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePublishedAppIntegrationOutput) {
}

func (newState *DeletePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState DeletePublishedAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o DeletePublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *DeletePublishedAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePublishedAppIntegrationRequest) {
}

func (newState *DeletePublishedAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState DeletePublishedAppIntegrationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePublishedAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePublishedAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePublishedAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (o DeletePublishedAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": o.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePublishedAppIntegrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse
// only implements ToObjectValue() and Type().
func (o DeleteResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (newState *DeleteServicePrincipalSecretRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteServicePrincipalSecretRequest) {
}

func (newState *DeleteServicePrincipalSecretRequest) SyncEffectiveFieldsDuringRead(existingState DeleteServicePrincipalSecretRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalSecretRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteServicePrincipalSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalSecretRequest
// only implements ToObjectValue() and Type().
func (o DeleteServicePrincipalSecretRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secret_id":            o.SecretId,
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalSecretRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"secret_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	ClientId types.String `tfsdk:"client_id" tf:"optional"`
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential types.Bool `tfsdk:"confidential" tf:"optional"`

	CreateTime types.String `tfsdk:"create_time" tf:"optional"`

	CreatedBy types.Int64 `tfsdk:"created_by" tf:"optional"`

	CreatorUsername types.String `tfsdk:"creator_username" tf:"optional"`
	// ID of this custom app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
	// The display name of the custom OAuth app
	Name types.String `tfsdk:"name" tf:"optional"`
	// List of OAuth redirect urls
	RedirectUrls types.List `tfsdk:"redirect_urls" tf:"optional"`

	Scopes types.List `tfsdk:"scopes" tf:"optional"`
	// Token access policy
	TokenAccessPolicy types.List `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *GetCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationOutput) {
}

func (newState *GetCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":       reflect.TypeOf(types.String{}),
		"scopes":              reflect.TypeOf(types.String{}),
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o GetCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_id":           o.ClientId,
			"confidential":        o.Confidential,
			"create_time":         o.CreateTime,
			"created_by":          o.CreatedBy,
			"creator_username":    o.CreatorUsername,
			"integration_id":      o.IntegrationId,
			"name":                o.Name,
			"redirect_urls":       o.RedirectUrls,
			"scopes":              o.Scopes,
			"token_access_policy": o.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client_id":        types.StringType,
			"confidential":     types.BoolType,
			"create_time":      types.StringType,
			"created_by":       types.Int64Type,
			"creator_username": types.StringType,
			"integration_id":   types.StringType,
			"name":             types.StringType,
			"redirect_urls": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_access_policy": basetypes.ListType{ElemType: TokenAccessPolicy{}.Type(ctx)},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in GetCustomAppIntegrationOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetCustomAppIntegrationOutput) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if o.RedirectUrls.IsNull() || o.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in GetCustomAppIntegrationOutput.
func (o *GetCustomAppIntegrationOutput) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in GetCustomAppIntegrationOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetCustomAppIntegrationOutput) GetScopes(ctx context.Context) ([]types.String, bool) {
	if o.Scopes.IsNull() || o.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in GetCustomAppIntegrationOutput.
func (o *GetCustomAppIntegrationOutput) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetCustomAppIntegrationOutput) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if o.TokenAccessPolicy.IsNull() || o.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy
	d := o.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput.
func (o *GetCustomAppIntegrationOutput) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	o.TokenAccessPolicy = types.ListValueMust(t, vs)
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *GetCustomAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationRequest) {
}

func (newState *GetCustomAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCustomAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (o GetCustomAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": o.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCustomAppIntegrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps types.List `tfsdk:"apps" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetCustomAppIntegrationsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationsOutput) {
}

func (newState *GetCustomAppIntegrationsOutput) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationsOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomAppIntegrationsOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCustomAppIntegrationsOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(GetCustomAppIntegrationOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationsOutput
// only implements ToObjectValue() and Type().
func (o GetCustomAppIntegrationsOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            o.Apps,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCustomAppIntegrationsOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: GetCustomAppIntegrationOutput{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in GetCustomAppIntegrationsOutput as
// a slice of GetCustomAppIntegrationOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetCustomAppIntegrationsOutput) GetApps(ctx context.Context) ([]GetCustomAppIntegrationOutput, bool) {
	if o.Apps.IsNull() || o.Apps.IsUnknown() {
		return nil, false
	}
	var v []GetCustomAppIntegrationOutput
	d := o.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetCustomAppIntegrationsOutput.
func (o *GetCustomAppIntegrationsOutput) SetApps(ctx context.Context, v []GetCustomAppIntegrationOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Apps = types.ListValueMust(t, vs)
}

type GetPublishedAppIntegrationOutput struct {
	// App-id of the published app integration
	AppId types.String `tfsdk:"app_id" tf:"optional"`

	CreateTime types.String `tfsdk:"create_time" tf:"optional"`

	CreatedBy types.Int64 `tfsdk:"created_by" tf:"optional"`
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id" tf:"optional"`
	// Display name of the published OAuth app
	Name types.String `tfsdk:"name" tf:"optional"`
	// Token access policy
	TokenAccessPolicy types.List `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *GetPublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationOutput) {
}

func (newState *GetPublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o GetPublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":              o.AppId,
			"create_time":         o.CreateTime,
			"created_by":          o.CreatedBy,
			"integration_id":      o.IntegrationId,
			"name":                o.Name,
			"token_access_policy": o.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_id":              types.StringType,
			"create_time":         types.StringType,
			"created_by":          types.Int64Type,
			"integration_id":      types.StringType,
			"name":                types.StringType,
			"token_access_policy": basetypes.ListType{ElemType: TokenAccessPolicy{}.Type(ctx)},
		},
	}
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in GetPublishedAppIntegrationOutput as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPublishedAppIntegrationOutput) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if o.TokenAccessPolicy.IsNull() || o.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy
	d := o.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetPublishedAppIntegrationOutput.
func (o *GetPublishedAppIntegrationOutput) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	o.TokenAccessPolicy = types.ListValueMust(t, vs)
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (newState *GetPublishedAppIntegrationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationRequest) {
}

func (newState *GetPublishedAppIntegrationRequest) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (o GetPublishedAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": o.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedAppIntegrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps types.List `tfsdk:"apps" tf:"optional"`

	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetPublishedAppIntegrationsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationsOutput) {
}

func (newState *GetPublishedAppIntegrationsOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationsOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedAppIntegrationsOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedAppIntegrationsOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(GetPublishedAppIntegrationOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationsOutput
// only implements ToObjectValue() and Type().
func (o GetPublishedAppIntegrationsOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            o.Apps,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedAppIntegrationsOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: GetPublishedAppIntegrationOutput{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in GetPublishedAppIntegrationsOutput as
// a slice of GetPublishedAppIntegrationOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPublishedAppIntegrationsOutput) GetApps(ctx context.Context) ([]GetPublishedAppIntegrationOutput, bool) {
	if o.Apps.IsNull() || o.Apps.IsUnknown() {
		return nil, false
	}
	var v []GetPublishedAppIntegrationOutput
	d := o.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetPublishedAppIntegrationsOutput.
func (o *GetPublishedAppIntegrationsOutput) SetApps(ctx context.Context, v []GetPublishedAppIntegrationOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Apps = types.ListValueMust(t, vs)
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps types.List `tfsdk:"apps" tf:"optional"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

func (newState *GetPublishedAppsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppsOutput) {
}

func (newState *GetPublishedAppsOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppsOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedAppsOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedAppsOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(PublishedAppOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppsOutput
// only implements ToObjectValue() and Type().
func (o GetPublishedAppsOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            o.Apps,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPublishedAppsOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: PublishedAppOutput{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in GetPublishedAppsOutput as
// a slice of PublishedAppOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetPublishedAppsOutput) GetApps(ctx context.Context) ([]PublishedAppOutput, bool) {
	if o.Apps.IsNull() || o.Apps.IsUnknown() {
		return nil, false
	}
	var v []PublishedAppOutput
	d := o.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetPublishedAppsOutput.
func (o *GetPublishedAppsOutput) SetApps(ctx context.Context, v []PublishedAppOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Apps = types.ListValueMust(t, vs)
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListCustomAppIntegrationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListCustomAppIntegrationsRequest) {
}

func (newState *ListCustomAppIntegrationsRequest) SyncEffectiveFieldsDuringRead(existingState ListCustomAppIntegrationsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCustomAppIntegrationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListCustomAppIntegrationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomAppIntegrationsRequest
// only implements ToObjectValue() and Type().
func (o ListCustomAppIntegrationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_creator_username": o.IncludeCreatorUsername,
			"page_size":                o.PageSize,
			"page_token":               o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListCustomAppIntegrationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_creator_username": types.BoolType,
			"page_size":                types.Int64Type,
			"page_token":               types.StringType,
		},
	}
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListOAuthPublishedAppsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListOAuthPublishedAppsRequest) {
}

func (newState *ListOAuthPublishedAppsRequest) SyncEffectiveFieldsDuringRead(existingState ListOAuthPublishedAppsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOAuthPublishedAppsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListOAuthPublishedAppsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOAuthPublishedAppsRequest
// only implements ToObjectValue() and Type().
func (o ListOAuthPublishedAppsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListOAuthPublishedAppsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Get published oauth app integrations
type ListPublishedAppIntegrationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListPublishedAppIntegrationsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPublishedAppIntegrationsRequest) {
}

func (newState *ListPublishedAppIntegrationsRequest) SyncEffectiveFieldsDuringRead(existingState ListPublishedAppIntegrationsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPublishedAppIntegrationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPublishedAppIntegrationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublishedAppIntegrationsRequest
// only implements ToObjectValue() and Type().
func (o ListPublishedAppIntegrationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPublishedAppIntegrationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the secrets for this service principal.
	// Provide this token to retrieve the next page of secret entries. When
	// providing a `page_token`, all other parameters provided to the request
	// must match the previous request. To list all of the secrets for a service
	// principal, it is necessary to continue requesting pages of entries until
	// the response contains no `next_page_token`. Note that the number of
	// entries returned must not be used to determine when the listing is
	// complete.
	PageToken types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (newState *ListServicePrincipalSecretsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalSecretsRequest) {
}

func (newState *ListServicePrincipalSecretsRequest) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalSecretsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalSecretsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalSecretsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalSecretsRequest
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalSecretsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token":           o.PageToken,
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalSecretsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token":           types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type ListServicePrincipalSecretsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of the secrets
	Secrets types.List `tfsdk:"secrets" tf:"optional"`
}

func (newState *ListServicePrincipalSecretsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalSecretsResponse) {
}

func (newState *ListServicePrincipalSecretsResponse) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalSecretsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalSecretsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalSecretsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalSecretsResponse
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalSecretsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"secrets":         o.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalSecretsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"secrets": basetypes.ListType{
				ElemType: SecretInfo{}.Type(ctx),
			},
		},
	}
}

// GetSecrets returns the value of the Secrets field in ListServicePrincipalSecretsResponse as
// a slice of SecretInfo values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListServicePrincipalSecretsResponse) GetSecrets(ctx context.Context) ([]SecretInfo, bool) {
	if o.Secrets.IsNull() || o.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretInfo
	d := o.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListServicePrincipalSecretsResponse.
func (o *ListServicePrincipalSecretsResponse) SetSecrets(ctx context.Context, v []SecretInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["secrets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Secrets = types.ListValueMust(t, vs)
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	AppId types.String `tfsdk:"app_id" tf:"optional"`
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	ClientId types.String `tfsdk:"client_id" tf:"optional"`
	// Description of the published OAuth app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	IsConfidentialClient types.Bool `tfsdk:"is_confidential_client" tf:"optional"`
	// The display name of the published OAuth app.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Redirect URLs of the published OAuth app.
	RedirectUrls types.List `tfsdk:"redirect_urls" tf:"optional"`
	// Required scopes for the published OAuth app.
	Scopes types.List `tfsdk:"scopes" tf:"optional"`
}

func (newState *PublishedAppOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedAppOutput) {
}

func (newState *PublishedAppOutput) SyncEffectiveFieldsDuringRead(existingState PublishedAppOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishedAppOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishedAppOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls": reflect.TypeOf(types.String{}),
		"scopes":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishedAppOutput
// only implements ToObjectValue() and Type().
func (o PublishedAppOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":                 o.AppId,
			"client_id":              o.ClientId,
			"description":            o.Description,
			"is_confidential_client": o.IsConfidentialClient,
			"name":                   o.Name,
			"redirect_urls":          o.RedirectUrls,
			"scopes":                 o.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublishedAppOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_id":                 types.StringType,
			"client_id":              types.StringType,
			"description":            types.StringType,
			"is_confidential_client": types.BoolType,
			"name":                   types.StringType,
			"redirect_urls": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in PublishedAppOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PublishedAppOutput) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if o.RedirectUrls.IsNull() || o.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in PublishedAppOutput.
func (o *PublishedAppOutput) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in PublishedAppOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PublishedAppOutput) GetScopes(ctx context.Context) ([]types.String, bool) {
	if o.Scopes.IsNull() || o.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in PublishedAppOutput.
func (o *PublishedAppOutput) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// ID of the secret
	Id types.String `tfsdk:"id" tf:"optional"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash" tf:"optional"`
	// Status of the secret
	Status types.String `tfsdk:"status" tf:"optional"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *SecretInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretInfo) {
}

func (newState *SecretInfo) SyncEffectiveFieldsDuringRead(existingState SecretInfo) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SecretInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SecretInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretInfo
// only implements ToObjectValue() and Type().
func (o SecretInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": o.CreateTime,
			"id":          o.Id,
			"secret_hash": o.SecretHash,
			"status":      o.Status,
			"update_time": o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SecretInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"id":          types.StringType,
			"secret_hash": types.StringType,
			"status":      types.StringType,
			"update_time": types.StringType,
		},
	}
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes types.Int64 `tfsdk:"access_token_ttl_in_minutes" tf:"optional"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes types.Int64 `tfsdk:"refresh_token_ttl_in_minutes" tf:"optional"`
}

func (newState *TokenAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessPolicy) {
}

func (newState *TokenAccessPolicy) SyncEffectiveFieldsDuringRead(existingState TokenAccessPolicy) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenAccessPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TokenAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessPolicy
// only implements ToObjectValue() and Type().
func (o TokenAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_token_ttl_in_minutes":  o.AccessTokenTtlInMinutes,
			"refresh_token_ttl_in_minutes": o.RefreshTokenTtlInMinutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenAccessPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_token_ttl_in_minutes":  types.Int64Type,
			"refresh_token_ttl_in_minutes": types.Int64Type,
		},
	}
}

type UpdateCustomAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls types.List `tfsdk:"redirect_urls" tf:"optional"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy types.List `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *UpdateCustomAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCustomAppIntegration) {
}

func (newState *UpdateCustomAppIntegration) SyncEffectiveFieldsDuringRead(existingState UpdateCustomAppIntegration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCustomAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":       reflect.TypeOf(types.String{}),
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomAppIntegration
// only implements ToObjectValue() and Type().
func (o UpdateCustomAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id":      o.IntegrationId,
			"redirect_urls":       o.RedirectUrls,
			"token_access_policy": o.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCustomAppIntegration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
			"redirect_urls": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_access_policy": basetypes.ListType{ElemType: TokenAccessPolicy{}.Type(ctx)},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in UpdateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCustomAppIntegration) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if o.RedirectUrls.IsNull() || o.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in UpdateCustomAppIntegration.
func (o *UpdateCustomAppIntegration) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RedirectUrls = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in UpdateCustomAppIntegration as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateCustomAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if o.TokenAccessPolicy.IsNull() || o.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy
	d := o.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdateCustomAppIntegration.
func (o *UpdateCustomAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	o.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type UpdateCustomAppIntegrationOutput struct {
}

func (newState *UpdateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCustomAppIntegrationOutput) {
}

func (newState *UpdateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState UpdateCustomAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o UpdateCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdatePublishedAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy types.List `tfsdk:"token_access_policy" tf:"optional,object"`
}

func (newState *UpdatePublishedAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePublishedAppIntegration) {
}

func (newState *UpdatePublishedAppIntegration) SyncEffectiveFieldsDuringRead(existingState UpdatePublishedAppIntegration) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePublishedAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePublishedAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePublishedAppIntegration
// only implements ToObjectValue() and Type().
func (o UpdatePublishedAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id":      o.IntegrationId,
			"token_access_policy": o.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePublishedAppIntegration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id":      types.StringType,
			"token_access_policy": basetypes.ListType{ElemType: TokenAccessPolicy{}.Type(ctx)},
		},
	}
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in UpdatePublishedAppIntegration as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePublishedAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if o.TokenAccessPolicy.IsNull() || o.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy
	d := o.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdatePublishedAppIntegration.
func (o *UpdatePublishedAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	o.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type UpdatePublishedAppIntegrationOutput struct {
}

func (newState *UpdatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePublishedAppIntegrationOutput) {
}

func (newState *UpdatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState UpdatePublishedAppIntegrationOutput) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (o UpdatePublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}
