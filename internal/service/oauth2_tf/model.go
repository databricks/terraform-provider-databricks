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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Create account federation policy
type CreateAccountFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`
	// The identifier for the federation policy. If unspecified, the id will be
	// assigned by Databricks.
	PolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o CreateAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":    o.Policy,
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy":    FederationPolicy{}.Type(ctx),
			"policy_id": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateAccountFederationPolicyRequest as
// a FederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateAccountFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if o.Policy.IsNull() || o.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy
	d := o.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPolicy sets the value of the Policy field in CreateAccountFederationPolicyRequest.
func (o *CreateAccountFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	o.Policy = vs
}

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential types.Bool `tfsdk:"confidential"`
	// Name of the custom OAuth app
	Name types.String `tfsdk:"name"`
	// List of OAuth redirect urls
	RedirectUrls types.List `tfsdk:"redirect_urls"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes types.List `tfsdk:"scopes"`
	// Token access policy
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (newState *CreateCustomAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomAppIntegration) {
}

func (newState *CreateCustomAppIntegration) SyncEffectiveFieldsDuringRead(existingState CreateCustomAppIntegration) {
}

func (c CreateCustomAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["confidential"] = attrs["confidential"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()

	return attrs
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
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
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
	d := o.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreateCustomAppIntegration.
func (o *CreateCustomAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TokenAccessPolicy = vs
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	ClientId types.String `tfsdk:"client_id"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret types.String `tfsdk:"client_secret"`
	// Unique integration id for the custom OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
}

func (newState *CreateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateCustomAppIntegrationOutput) {
}

func (newState *CreateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState CreateCustomAppIntegrationOutput) {
}

func (c CreateCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client_id"] = attrs["client_id"].SetOptional()
	attrs["client_secret"] = attrs["client_secret"].SetOptional()
	attrs["integration_id"] = attrs["integration_id"].SetOptional()

	return attrs
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
	AppId types.String `tfsdk:"app_id"`
	// Token access policy
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (newState *CreatePublishedAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePublishedAppIntegration) {
}

func (newState *CreatePublishedAppIntegration) SyncEffectiveFieldsDuringRead(existingState CreatePublishedAppIntegration) {
}

func (c CreatePublishedAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_id"] = attrs["app_id"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()

	return attrs
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
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
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
	d := o.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreatePublishedAppIntegration.
func (o *CreatePublishedAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TokenAccessPolicy = vs
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
}

func (newState *CreatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePublishedAppIntegrationOutput) {
}

func (newState *CreatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState CreatePublishedAppIntegrationOutput) {
}

func (c CreatePublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetOptional()

	return attrs
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

// Create service principal federation policy
type CreateServicePrincipalFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`
	// The identifier for the federation policy. If unspecified, the id will be
	// assigned by Databricks.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o CreateServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":               o.Policy,
			"policy_id":            o.PolicyId,
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy":               FederationPolicy{}.Type(ctx),
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateServicePrincipalFederationPolicyRequest as
// a FederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServicePrincipalFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if o.Policy.IsNull() || o.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy
	d := o.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPolicy sets the value of the Policy field in CreateServicePrincipalFederationPolicyRequest.
func (o *CreateServicePrincipalFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	o.Policy = vs
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
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
	CreateTime types.String `tfsdk:"create_time"`
	// ID of the secret
	Id types.String `tfsdk:"id"`
	// Secret Value
	Secret types.String `tfsdk:"secret"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash"`
	// Status of the secret
	Status types.String `tfsdk:"status"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *CreateServicePrincipalSecretResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServicePrincipalSecretResponse) {
}

func (newState *CreateServicePrincipalSecretResponse) SyncEffectiveFieldsDuringRead(existingState CreateServicePrincipalSecretResponse) {
}

func (c CreateServicePrincipalSecretResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["secret"] = attrs["secret"].SetOptional()
	attrs["secret_hash"] = attrs["secret_hash"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()

	return attrs
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
	AuthorizationDetails types.String `tfsdk:"authorization_details"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl types.String `tfsdk:"endpoint_url"`
}

func (newState *DataPlaneInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataPlaneInfo) {
}

func (newState *DataPlaneInfo) SyncEffectiveFieldsDuringRead(existingState DataPlaneInfo) {
}

func (c DataPlaneInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["authorization_details"] = attrs["authorization_details"].SetOptional()
	attrs["endpoint_url"] = attrs["endpoint_url"].SetOptional()

	return attrs
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

// Delete account federation policy
type DeleteAccountFederationPolicyRequest struct {
	PolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DeleteCustomAppIntegrationOutput struct {
}

func (newState *DeleteCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteCustomAppIntegrationOutput) {
}

func (newState *DeleteCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState DeleteCustomAppIntegrationOutput) {
}

func (c DeleteCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

func (c DeletePublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// Delete service principal federation policy
type DeleteServicePrincipalFederationPolicyRequest struct {
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o DeleteServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id":            o.PolicyId,
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
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

type FederationPolicy struct {
	// Creation time of the federation policy.
	CreateTime types.String `tfsdk:"create_time"`
	// Description of the federation policy.
	Description types.String `tfsdk:"description"`
	// Name of the federation policy. The name must contain only lowercase
	// alphanumeric characters, numbers, and hyphens. It must be unique within
	// the account.
	Name types.String `tfsdk:"name"`
	// Specifies the policy to use for validating OIDC claims in your federated
	// tokens.
	OidcPolicy types.Object `tfsdk:"oidc_policy"`
	// Unique, immutable id of the federation policy.
	Uid types.String `tfsdk:"uid"`
	// Last update time of the federation policy.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *FederationPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan FederationPolicy) {
}

func (newState *FederationPolicy) SyncEffectiveFieldsDuringRead(existingState FederationPolicy) {
}

func (c FederationPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["oidc_policy"] = attrs["oidc_policy"].SetOptional()
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FederationPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FederationPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"oidc_policy": reflect.TypeOf(OidcFederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicy
// only implements ToObjectValue() and Type().
func (o FederationPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": o.CreateTime,
			"description": o.Description,
			"name":        o.Name,
			"oidc_policy": o.OidcPolicy,
			"uid":         o.Uid,
			"update_time": o.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FederationPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"description": types.StringType,
			"name":        types.StringType,
			"oidc_policy": OidcFederationPolicy{}.Type(ctx),
			"uid":         types.StringType,
			"update_time": types.StringType,
		},
	}
}

// GetOidcPolicy returns the value of the OidcPolicy field in FederationPolicy as
// a OidcFederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *FederationPolicy) GetOidcPolicy(ctx context.Context) (OidcFederationPolicy, bool) {
	var e OidcFederationPolicy
	if o.OidcPolicy.IsNull() || o.OidcPolicy.IsUnknown() {
		return e, false
	}
	var v []OidcFederationPolicy
	d := o.OidcPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetOidcPolicy sets the value of the OidcPolicy field in FederationPolicy.
func (o *FederationPolicy) SetOidcPolicy(ctx context.Context, v OidcFederationPolicy) {
	vs := v.ToObjectValue(ctx)
	o.OidcPolicy = vs
}

// Get account federation policy
type GetAccountFederationPolicyRequest struct {
	PolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o GetAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": o.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	ClientId types.String `tfsdk:"client_id"`
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential types.Bool `tfsdk:"confidential"`

	CreateTime types.String `tfsdk:"create_time"`

	CreatedBy types.Int64 `tfsdk:"created_by"`

	CreatorUsername types.String `tfsdk:"creator_username"`
	// ID of this custom app
	IntegrationId types.String `tfsdk:"integration_id"`
	// The display name of the custom OAuth app
	Name types.String `tfsdk:"name"`
	// List of OAuth redirect urls
	RedirectUrls types.List `tfsdk:"redirect_urls"`

	Scopes types.List `tfsdk:"scopes"`
	// Token access policy
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (newState *GetCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationOutput) {
}

func (newState *GetCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationOutput) {
}

func (c GetCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client_id"] = attrs["client_id"].SetOptional()
	attrs["confidential"] = attrs["confidential"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["creator_username"] = attrs["creator_username"].SetOptional()
	attrs["integration_id"] = attrs["integration_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()

	return attrs
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
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
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
	d := o.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput.
func (o *GetCustomAppIntegrationOutput) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TokenAccessPolicy = vs
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId types.String `tfsdk:"-"`
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
	Apps types.List `tfsdk:"apps"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *GetCustomAppIntegrationsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetCustomAppIntegrationsOutput) {
}

func (newState *GetCustomAppIntegrationsOutput) SyncEffectiveFieldsDuringRead(existingState GetCustomAppIntegrationsOutput) {
}

func (c GetCustomAppIntegrationsOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apps"] = attrs["apps"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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
	AppId types.String `tfsdk:"app_id"`

	CreateTime types.String `tfsdk:"create_time"`

	CreatedBy types.Int64 `tfsdk:"created_by"`
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
	// Display name of the published OAuth app
	Name types.String `tfsdk:"name"`
	// Token access policy
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (newState *GetPublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationOutput) {
}

func (newState *GetPublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationOutput) {
}

func (c GetPublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_id"] = attrs["app_id"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["integration_id"] = attrs["integration_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()

	return attrs
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
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
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
	d := o.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetPublishedAppIntegrationOutput.
func (o *GetPublishedAppIntegrationOutput) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TokenAccessPolicy = vs
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
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
	Apps types.List `tfsdk:"apps"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *GetPublishedAppIntegrationsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppIntegrationsOutput) {
}

func (newState *GetPublishedAppIntegrationsOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppIntegrationsOutput) {
}

func (c GetPublishedAppIntegrationsOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apps"] = attrs["apps"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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
	Apps types.List `tfsdk:"apps"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *GetPublishedAppsOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedAppsOutput) {
}

func (newState *GetPublishedAppsOutput) SyncEffectiveFieldsDuringRead(existingState GetPublishedAppsOutput) {
}

func (c GetPublishedAppsOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["apps"] = attrs["apps"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
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

// Get service principal federation policy
type GetServicePrincipalFederationPolicyRequest struct {
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o GetServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id":            o.PolicyId,
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

// List account federation policies
type ListAccountFederationPoliciesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountFederationPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListAccountFederationPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountFederationPoliciesRequest
// only implements ToObjectValue() and Type().
func (o ListAccountFederationPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListAccountFederationPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
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

type ListFederationPoliciesResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Policies types.List `tfsdk:"policies"`
}

func (newState *ListFederationPoliciesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListFederationPoliciesResponse) {
}

func (newState *ListFederationPoliciesResponse) SyncEffectiveFieldsDuringRead(existingState ListFederationPoliciesResponse) {
}

func (c ListFederationPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["policies"] = attrs["policies"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListFederationPoliciesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListFederationPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFederationPoliciesResponse
// only implements ToObjectValue() and Type().
func (o ListFederationPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"policies":        o.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListFederationPoliciesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policies": basetypes.ListType{
				ElemType: FederationPolicy{}.Type(ctx),
			},
		},
	}
}

// GetPolicies returns the value of the Policies field in ListFederationPoliciesResponse as
// a slice of FederationPolicy values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListFederationPoliciesResponse) GetPolicies(ctx context.Context) ([]FederationPolicy, bool) {
	if o.Policies.IsNull() || o.Policies.IsUnknown() {
		return nil, false
	}
	var v []FederationPolicy
	d := o.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListFederationPoliciesResponse.
func (o *ListFederationPoliciesResponse) SetPolicies(ctx context.Context, v []FederationPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Policies = types.ListValueMust(t, vs)
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
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

// List service principal federation policies
type ListServicePrincipalFederationPoliciesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalFederationPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServicePrincipalFederationPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalFederationPoliciesRequest
// only implements ToObjectValue() and Type().
func (o ListServicePrincipalFederationPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":            o.PageSize,
			"page_token":           o.PageToken,
			"service_principal_id": o.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListServicePrincipalFederationPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":            types.Int64Type,
			"page_token":           types.StringType,
			"service_principal_id": types.Int64Type,
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
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of the secrets
	Secrets types.List `tfsdk:"secrets"`
}

func (newState *ListServicePrincipalSecretsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListServicePrincipalSecretsResponse) {
}

func (newState *ListServicePrincipalSecretsResponse) SyncEffectiveFieldsDuringRead(existingState ListServicePrincipalSecretsResponse) {
}

func (c ListServicePrincipalSecretsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["secrets"] = attrs["secrets"].SetOptional()

	return attrs
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

// Specifies the policy to use for validating OIDC claims in your federated
// tokens.
type OidcFederationPolicy struct {
	// The allowed token audiences, as specified in the 'aud' claim of federated
	// tokens. The audience identifier is intended to represent the recipient of
	// the token. Can be any non-empty string value. As long as the audience in
	// the token matches at least one audience in the policy, the token is
	// considered a match. If audiences is unspecified, defaults to your
	// Databricks account id.
	Audiences types.List `tfsdk:"audiences"`
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	Issuer types.String `tfsdk:"issuer"`
	// The public keys used to validate the signature of federated tokens, in
	// JWKS format. If unspecified (recommended), Databricks automatically
	// fetches the public keys from your issuer’s well known endpoint.
	// Databricks strongly recommends relying on your issuer’s well known
	// endpoint for discovering public keys.
	JwksJson types.String `tfsdk:"jwks_json"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. Must be specified for service principal federation
	// policies. Must not be specified for account federation policies.
	Subject types.String `tfsdk:"subject"`
	// The claim that contains the subject of the token. If unspecified, the
	// default value is 'sub'.
	SubjectClaim types.String `tfsdk:"subject_claim"`
}

func (newState *OidcFederationPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan OidcFederationPolicy) {
}

func (newState *OidcFederationPolicy) SyncEffectiveFieldsDuringRead(existingState OidcFederationPolicy) {
}

func (c OidcFederationPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["audiences"] = attrs["audiences"].SetOptional()
	attrs["issuer"] = attrs["issuer"].SetOptional()
	attrs["jwks_json"] = attrs["jwks_json"].SetOptional()
	attrs["subject"] = attrs["subject"].SetOptional()
	attrs["subject_claim"] = attrs["subject_claim"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OidcFederationPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OidcFederationPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"audiences": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OidcFederationPolicy
// only implements ToObjectValue() and Type().
func (o OidcFederationPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"audiences":     o.Audiences,
			"issuer":        o.Issuer,
			"jwks_json":     o.JwksJson,
			"subject":       o.Subject,
			"subject_claim": o.SubjectClaim,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OidcFederationPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"audiences": basetypes.ListType{
				ElemType: types.StringType,
			},
			"issuer":        types.StringType,
			"jwks_json":     types.StringType,
			"subject":       types.StringType,
			"subject_claim": types.StringType,
		},
	}
}

// GetAudiences returns the value of the Audiences field in OidcFederationPolicy as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *OidcFederationPolicy) GetAudiences(ctx context.Context) ([]types.String, bool) {
	if o.Audiences.IsNull() || o.Audiences.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Audiences.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAudiences sets the value of the Audiences field in OidcFederationPolicy.
func (o *OidcFederationPolicy) SetAudiences(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["audiences"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Audiences = types.ListValueMust(t, vs)
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	AppId types.String `tfsdk:"app_id"`
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	ClientId types.String `tfsdk:"client_id"`
	// Description of the published OAuth app.
	Description types.String `tfsdk:"description"`
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	IsConfidentialClient types.Bool `tfsdk:"is_confidential_client"`
	// The display name of the published OAuth app.
	Name types.String `tfsdk:"name"`
	// Redirect URLs of the published OAuth app.
	RedirectUrls types.List `tfsdk:"redirect_urls"`
	// Required scopes for the published OAuth app.
	Scopes types.List `tfsdk:"scopes"`
}

func (newState *PublishedAppOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedAppOutput) {
}

func (newState *PublishedAppOutput) SyncEffectiveFieldsDuringRead(existingState PublishedAppOutput) {
}

func (c PublishedAppOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_id"] = attrs["app_id"].SetOptional()
	attrs["client_id"] = attrs["client_id"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["is_confidential_client"] = attrs["is_confidential_client"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()

	return attrs
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
	CreateTime types.String `tfsdk:"create_time"`
	// ID of the secret
	Id types.String `tfsdk:"id"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash"`
	// Status of the secret
	Status types.String `tfsdk:"status"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time"`
}

func (newState *SecretInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan SecretInfo) {
}

func (newState *SecretInfo) SyncEffectiveFieldsDuringRead(existingState SecretInfo) {
}

func (c SecretInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["secret_hash"] = attrs["secret_hash"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["update_time"] = attrs["update_time"].SetOptional()

	return attrs
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
	AccessTokenTtlInMinutes types.Int64 `tfsdk:"access_token_ttl_in_minutes"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes types.Int64 `tfsdk:"refresh_token_ttl_in_minutes"`
}

func (newState *TokenAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessPolicy) {
}

func (newState *TokenAccessPolicy) SyncEffectiveFieldsDuringRead(existingState TokenAccessPolicy) {
}

func (c TokenAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_token_ttl_in_minutes"] = attrs["access_token_ttl_in_minutes"].SetOptional()
	attrs["refresh_token_ttl_in_minutes"] = attrs["refresh_token_ttl_in_minutes"].SetOptional()

	return attrs
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

// Update account federation policy
type UpdateAccountFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`

	PolicyId types.String `tfsdk:"-"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o UpdateAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":      o.Policy,
			"policy_id":   o.PolicyId,
			"update_mask": o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy":      FederationPolicy{}.Type(ctx),
			"policy_id":   types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in UpdateAccountFederationPolicyRequest as
// a FederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateAccountFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if o.Policy.IsNull() || o.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy
	d := o.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPolicy sets the value of the Policy field in UpdateAccountFederationPolicyRequest.
func (o *UpdateAccountFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	o.Policy = vs
}

type UpdateCustomAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls types.List `tfsdk:"redirect_urls"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (newState *UpdateCustomAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCustomAppIntegration) {
}

func (newState *UpdateCustomAppIntegration) SyncEffectiveFieldsDuringRead(existingState UpdateCustomAppIntegration) {
}

func (c UpdateCustomAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetRequired()
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()

	return attrs
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
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
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
	d := o.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdateCustomAppIntegration.
func (o *UpdateCustomAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TokenAccessPolicy = vs
}

type UpdateCustomAppIntegrationOutput struct {
}

func (newState *UpdateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCustomAppIntegrationOutput) {
}

func (newState *UpdateCustomAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState UpdateCustomAppIntegrationOutput) {
}

func (c UpdateCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (newState *UpdatePublishedAppIntegration) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePublishedAppIntegration) {
}

func (newState *UpdatePublishedAppIntegration) SyncEffectiveFieldsDuringRead(existingState UpdatePublishedAppIntegration) {
}

func (c UpdatePublishedAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetRequired()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()

	return attrs
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
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
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
	d := o.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdatePublishedAppIntegration.
func (o *UpdatePublishedAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.TokenAccessPolicy = vs
}

type UpdatePublishedAppIntegrationOutput struct {
}

func (newState *UpdatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePublishedAppIntegrationOutput) {
}

func (newState *UpdatePublishedAppIntegrationOutput) SyncEffectiveFieldsDuringRead(existingState UpdatePublishedAppIntegrationOutput) {
}

func (c UpdatePublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
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

// Update service principal federation policy
type UpdateServicePrincipalFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`

	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (o UpdateServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":               o.Policy,
			"policy_id":            o.PolicyId,
			"service_principal_id": o.ServicePrincipalId,
			"update_mask":          o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy":               FederationPolicy{}.Type(ctx),
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
			"update_mask":          types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in UpdateServicePrincipalFederationPolicyRequest as
// a FederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateServicePrincipalFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if o.Policy.IsNull() || o.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy
	d := o.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPolicy sets the value of the Policy field in UpdateServicePrincipalFederationPolicyRequest.
func (o *UpdateServicePrincipalFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	o.Policy = vs
}
