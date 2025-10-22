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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateAccountFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId types.String `tfsdk:"-"`
}

func (to *CreateAccountFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAccountFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *CreateAccountFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateAccountFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m CreateAccountFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m CreateAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":    m.Policy,
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateAccountFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v FederationPolicy
	d := m.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicy sets the value of the Policy field in CreateAccountFederationPolicyRequest.
func (m *CreateAccountFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	m.Policy = vs
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
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes types.List `tfsdk:"user_authorized_scopes"`
}

func (to *CreateCustomAppIntegration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomAppIntegration) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				// Recursively sync the fields of TokenAccessPolicy
				toTokenAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
	if !from.UserAuthorizedScopes.IsNull() && !from.UserAuthorizedScopes.IsUnknown() && to.UserAuthorizedScopes.IsNull() && len(from.UserAuthorizedScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserAuthorizedScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserAuthorizedScopes = from.UserAuthorizedScopes
	}
}

func (to *CreateCustomAppIntegration) SyncFieldsDuringRead(ctx context.Context, from CreateCustomAppIntegration) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
	if !from.UserAuthorizedScopes.IsNull() && !from.UserAuthorizedScopes.IsUnknown() && to.UserAuthorizedScopes.IsNull() && len(from.UserAuthorizedScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserAuthorizedScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserAuthorizedScopes = from.UserAuthorizedScopes
	}
}

func (m CreateCustomAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["confidential"] = attrs["confidential"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["user_authorized_scopes"] = attrs["user_authorized_scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateCustomAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateCustomAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":          reflect.TypeOf(types.String{}),
		"scopes":                 reflect.TypeOf(types.String{}),
		"token_access_policy":    reflect.TypeOf(TokenAccessPolicy{}),
		"user_authorized_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomAppIntegration
// only implements ToObjectValue() and Type().
func (m CreateCustomAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"confidential":           m.Confidential,
			"name":                   m.Name,
			"redirect_urls":          m.RedirectUrls,
			"scopes":                 m.Scopes,
			"token_access_policy":    m.TokenAccessPolicy,
			"user_authorized_scopes": m.UserAuthorizedScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCustomAppIntegration) Type(ctx context.Context) attr.Type {
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
			"user_authorized_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in CreateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if m.RedirectUrls.IsNull() || m.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in CreateCustomAppIntegration.
func (m *CreateCustomAppIntegration) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in CreateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration) GetScopes(ctx context.Context) ([]types.String, bool) {
	if m.Scopes.IsNull() || m.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in CreateCustomAppIntegration.
func (m *CreateCustomAppIntegration) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in CreateCustomAppIntegration as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v TokenAccessPolicy
	d := m.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreateCustomAppIntegration.
func (m *CreateCustomAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TokenAccessPolicy = vs
}

// GetUserAuthorizedScopes returns the value of the UserAuthorizedScopes field in CreateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration) GetUserAuthorizedScopes(ctx context.Context) ([]types.String, bool) {
	if m.UserAuthorizedScopes.IsNull() || m.UserAuthorizedScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.UserAuthorizedScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserAuthorizedScopes sets the value of the UserAuthorizedScopes field in CreateCustomAppIntegration.
func (m *CreateCustomAppIntegration) SetUserAuthorizedScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_authorized_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserAuthorizedScopes = types.ListValueMust(t, vs)
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

func (to *CreateCustomAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomAppIntegrationOutput) {
}

func (to *CreateCustomAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from CreateCustomAppIntegrationOutput) {
}

func (m CreateCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m CreateCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_id":      m.ClientId,
			"client_secret":  m.ClientSecret,
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
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

func (to *CreatePublishedAppIntegration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePublishedAppIntegration) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				// Recursively sync the fields of TokenAccessPolicy
				toTokenAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (to *CreatePublishedAppIntegration) SyncFieldsDuringRead(ctx context.Context, from CreatePublishedAppIntegration) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (m CreatePublishedAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePublishedAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePublishedAppIntegration
// only implements ToObjectValue() and Type().
func (m CreatePublishedAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":              m.AppId,
			"token_access_policy": m.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePublishedAppIntegration) Type(ctx context.Context) attr.Type {
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
func (m *CreatePublishedAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v TokenAccessPolicy
	d := m.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreatePublishedAppIntegration.
func (m *CreatePublishedAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TokenAccessPolicy = vs
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
}

func (to *CreatePublishedAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePublishedAppIntegrationOutput) {
}

func (to *CreatePublishedAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from CreatePublishedAppIntegrationOutput) {
}

func (m CreatePublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m CreatePublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type CreateServicePrincipalFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *CreateServicePrincipalFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *CreateServicePrincipalFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m CreateServicePrincipalFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":               m.Policy,
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
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
func (m *CreateServicePrincipalFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v FederationPolicy
	d := m.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicy sets the value of the Policy field in CreateServicePrincipalFederationPolicyRequest.
func (m *CreateServicePrincipalFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	m.Policy = vs
}

type CreateServicePrincipalSecretRequest struct {
	// The lifetime of the secret in seconds. If this parameter is not provided,
	// the secret will have a default lifetime of 730 days (63072000s).
	Lifetime types.String `tfsdk:"lifetime"`
	// The service principal ID.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *CreateServicePrincipalSecretRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalSecretRequest) {
}

func (to *CreateServicePrincipalSecretRequest) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalSecretRequest) {
}

func (m CreateServicePrincipalSecretRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["lifetime"] = attrs["lifetime"].SetOptional()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServicePrincipalSecretRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateServicePrincipalSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalSecretRequest
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalSecretRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"lifetime":             m.Lifetime,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalSecretRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"lifetime":             types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	CreateTime types.String `tfsdk:"create_time"`
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	ExpireTime types.String `tfsdk:"expire_time"`
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

func (to *CreateServicePrincipalSecretResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalSecretResponse) {
}

func (to *CreateServicePrincipalSecretResponse) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalSecretResponse) {
}

func (m CreateServicePrincipalSecretResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["expire_time"] = attrs["expire_time"].SetOptional()
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
func (m CreateServicePrincipalSecretResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalSecretResponse
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalSecretResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"expire_time": m.ExpireTime,
			"id":          m.Id,
			"secret":      m.Secret,
			"secret_hash": m.SecretHash,
			"status":      m.Status,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalSecretResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"expire_time": types.StringType,
			"id":          types.StringType,
			"secret":      types.StringType,
			"secret_hash": types.StringType,
			"status":      types.StringType,
			"update_time": types.StringType,
		},
	}
}

type DeleteAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *DeleteAccountFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountFederationPolicyRequest) {
}

func (to *DeleteAccountFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountFederationPolicyRequest) {
}

func (m DeleteAccountFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m DeleteAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DeleteCustomAppIntegrationOutput struct {
}

func (to *DeleteCustomAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCustomAppIntegrationOutput) {
}

func (to *DeleteCustomAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from DeleteCustomAppIntegrationOutput) {
}

func (m DeleteCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m DeleteCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteCustomAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (to *DeleteCustomAppIntegrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCustomAppIntegrationRequest) {
}

func (to *DeleteCustomAppIntegrationRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteCustomAppIntegrationRequest) {
}

func (m DeleteCustomAppIntegrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCustomAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (m DeleteCustomAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCustomAppIntegrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type DeletePublishedAppIntegrationOutput struct {
}

func (to *DeletePublishedAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePublishedAppIntegrationOutput) {
}

func (to *DeletePublishedAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from DeletePublishedAppIntegrationOutput) {
}

func (m DeletePublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m DeletePublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeletePublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (to *DeletePublishedAppIntegrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePublishedAppIntegrationRequest) {
}

func (to *DeletePublishedAppIntegrationRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePublishedAppIntegrationRequest) {
}

func (m DeletePublishedAppIntegrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePublishedAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePublishedAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePublishedAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (m DeletePublishedAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePublishedAppIntegrationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type DeleteServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteServicePrincipalFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalFederationPolicyRequest) {
}

func (to *DeleteServicePrincipalFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalFederationPolicyRequest) {
}

func (m DeleteServicePrincipalFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *DeleteServicePrincipalSecretRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalSecretRequest) {
}

func (to *DeleteServicePrincipalSecretRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalSecretRequest) {
}

func (m DeleteServicePrincipalSecretRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["secret_id"] = attrs["secret_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServicePrincipalSecretRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServicePrincipalSecretRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalSecretRequest
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalSecretRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secret_id":            m.SecretId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalSecretRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"secret_id":            types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

type FederationPolicy struct {
	// Creation time of the federation policy.
	CreateTime types.String `tfsdk:"create_time"`
	// Description of the federation policy.
	Description types.String `tfsdk:"description"`
	// Resource name for the federation policy. Example values include
	// `accounts/<account-id>/federationPolicies/my-federation-policy` for
	// Account Federation Policies, and
	// `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
	// for Service Principal Federation Policies. Typically an output parameter,
	// which does not need to be specified in create or update requests. If
	// specified in a request, must match the value in the request URL.
	Name types.String `tfsdk:"name"`

	OidcPolicy types.Object `tfsdk:"oidc_policy"`
	// The ID of the federation policy. Output only.
	PolicyId types.String `tfsdk:"policy_id"`
	// The service principal ID that this federation policy applies to. Output
	// only. Only set for service principal federation policies.
	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id"`
	// Unique, immutable id of the federation policy.
	Uid types.String `tfsdk:"uid"`
	// Last update time of the federation policy.
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *FederationPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FederationPolicy) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				// Recursively sync the fields of OidcPolicy
				toOidcPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

func (to *FederationPolicy) SyncFieldsDuringRead(ctx context.Context, from FederationPolicy) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				toOidcPolicy.SyncFieldsDuringRead(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

func (m FederationPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["oidc_policy"] = attrs["oidc_policy"].SetOptional()
	attrs["policy_id"] = attrs["policy_id"].SetComputed()
	attrs["policy_id"] = attrs["policy_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["service_principal_id"] = attrs["service_principal_id"].SetComputed()
	attrs["service_principal_id"] = attrs["service_principal_id"].(tfschema.Int64AttributeBuilder).AddPlanModifier(int64planmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["uid"] = attrs["uid"].SetComputed()
	attrs["uid"] = attrs["uid"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
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
func (m FederationPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"oidc_policy": reflect.TypeOf(OidcFederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicy
// only implements ToObjectValue() and Type().
func (m FederationPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":          m.CreateTime,
			"description":          m.Description,
			"name":                 m.Name,
			"oidc_policy":          m.OidcPolicy,
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
			"uid":                  m.Uid,
			"update_time":          m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FederationPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":          types.StringType,
			"description":          types.StringType,
			"name":                 types.StringType,
			"oidc_policy":          OidcFederationPolicy{}.Type(ctx),
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
			"uid":                  types.StringType,
			"update_time":          types.StringType,
		},
	}
}

// GetOidcPolicy returns the value of the OidcPolicy field in FederationPolicy as
// a OidcFederationPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *FederationPolicy) GetOidcPolicy(ctx context.Context) (OidcFederationPolicy, bool) {
	var e OidcFederationPolicy
	if m.OidcPolicy.IsNull() || m.OidcPolicy.IsUnknown() {
		return e, false
	}
	var v OidcFederationPolicy
	d := m.OidcPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOidcPolicy sets the value of the OidcPolicy field in FederationPolicy.
func (m *FederationPolicy) SetOidcPolicy(ctx context.Context, v OidcFederationPolicy) {
	vs := v.ToObjectValue(ctx)
	m.OidcPolicy = vs
}

type GetAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *GetAccountFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAccountFederationPolicyRequest) {
}

func (to *GetAccountFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from GetAccountFederationPolicyRequest) {
}

func (m GetAccountFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m GetAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
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
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes types.List `tfsdk:"user_authorized_scopes"`
}

func (to *GetCustomAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomAppIntegrationOutput) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				// Recursively sync the fields of TokenAccessPolicy
				toTokenAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
	if !from.UserAuthorizedScopes.IsNull() && !from.UserAuthorizedScopes.IsUnknown() && to.UserAuthorizedScopes.IsNull() && len(from.UserAuthorizedScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserAuthorizedScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserAuthorizedScopes = from.UserAuthorizedScopes
	}
}

func (to *GetCustomAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from GetCustomAppIntegrationOutput) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
	if !from.UserAuthorizedScopes.IsNull() && !from.UserAuthorizedScopes.IsUnknown() && to.UserAuthorizedScopes.IsNull() && len(from.UserAuthorizedScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserAuthorizedScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserAuthorizedScopes = from.UserAuthorizedScopes
	}
}

func (m GetCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["user_authorized_scopes"] = attrs["user_authorized_scopes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":          reflect.TypeOf(types.String{}),
		"scopes":                 reflect.TypeOf(types.String{}),
		"token_access_policy":    reflect.TypeOf(TokenAccessPolicy{}),
		"user_authorized_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m GetCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_id":              m.ClientId,
			"confidential":           m.Confidential,
			"create_time":            m.CreateTime,
			"created_by":             m.CreatedBy,
			"creator_username":       m.CreatorUsername,
			"integration_id":         m.IntegrationId,
			"name":                   m.Name,
			"redirect_urls":          m.RedirectUrls,
			"scopes":                 m.Scopes,
			"token_access_policy":    m.TokenAccessPolicy,
			"user_authorized_scopes": m.UserAuthorizedScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
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
			"user_authorized_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in GetCustomAppIntegrationOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if m.RedirectUrls.IsNull() || m.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in GetCustomAppIntegrationOutput.
func (m *GetCustomAppIntegrationOutput) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in GetCustomAppIntegrationOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput) GetScopes(ctx context.Context) ([]types.String, bool) {
	if m.Scopes.IsNull() || m.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in GetCustomAppIntegrationOutput.
func (m *GetCustomAppIntegrationOutput) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v TokenAccessPolicy
	d := m.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput.
func (m *GetCustomAppIntegrationOutput) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TokenAccessPolicy = vs
}

// GetUserAuthorizedScopes returns the value of the UserAuthorizedScopes field in GetCustomAppIntegrationOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput) GetUserAuthorizedScopes(ctx context.Context) ([]types.String, bool) {
	if m.UserAuthorizedScopes.IsNull() || m.UserAuthorizedScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.UserAuthorizedScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserAuthorizedScopes sets the value of the UserAuthorizedScopes field in GetCustomAppIntegrationOutput.
func (m *GetCustomAppIntegrationOutput) SetUserAuthorizedScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_authorized_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserAuthorizedScopes = types.ListValueMust(t, vs)
}

type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId types.String `tfsdk:"-"`
}

func (to *GetCustomAppIntegrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomAppIntegrationRequest) {
}

func (to *GetCustomAppIntegrationRequest) SyncFieldsDuringRead(ctx context.Context, from GetCustomAppIntegrationRequest) {
}

func (m GetCustomAppIntegrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCustomAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCustomAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (m GetCustomAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCustomAppIntegrationRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetCustomAppIntegrationsOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomAppIntegrationsOutput) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *GetCustomAppIntegrationsOutput) SyncFieldsDuringRead(ctx context.Context, from GetCustomAppIntegrationsOutput) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m GetCustomAppIntegrationsOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCustomAppIntegrationsOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(GetCustomAppIntegrationOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationsOutput
// only implements ToObjectValue() and Type().
func (m GetCustomAppIntegrationsOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCustomAppIntegrationsOutput) Type(ctx context.Context) attr.Type {
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
func (m *GetCustomAppIntegrationsOutput) GetApps(ctx context.Context) ([]GetCustomAppIntegrationOutput, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []GetCustomAppIntegrationOutput
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetCustomAppIntegrationsOutput.
func (m *GetCustomAppIntegrationsOutput) SetApps(ctx context.Context, v []GetCustomAppIntegrationOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
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

func (to *GetPublishedAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppIntegrationOutput) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				// Recursively sync the fields of TokenAccessPolicy
				toTokenAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (to *GetPublishedAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppIntegrationOutput) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (m GetPublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m GetPublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":              m.AppId,
			"create_time":         m.CreateTime,
			"created_by":          m.CreatedBy,
			"integration_id":      m.IntegrationId,
			"name":                m.Name,
			"token_access_policy": m.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
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
func (m *GetPublishedAppIntegrationOutput) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v TokenAccessPolicy
	d := m.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetPublishedAppIntegrationOutput.
func (m *GetPublishedAppIntegrationOutput) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TokenAccessPolicy = vs
}

type GetPublishedAppIntegrationRequest struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (to *GetPublishedAppIntegrationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppIntegrationRequest) {
}

func (to *GetPublishedAppIntegrationRequest) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppIntegrationRequest) {
}

func (m GetPublishedAppIntegrationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_id"] = attrs["integration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedAppIntegrationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPublishedAppIntegrationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationRequest
// only implements ToObjectValue() and Type().
func (m GetPublishedAppIntegrationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppIntegrationRequest) Type(ctx context.Context) attr.Type {
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

func (to *GetPublishedAppIntegrationsOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppIntegrationsOutput) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *GetPublishedAppIntegrationsOutput) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppIntegrationsOutput) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m GetPublishedAppIntegrationsOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPublishedAppIntegrationsOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(GetPublishedAppIntegrationOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationsOutput
// only implements ToObjectValue() and Type().
func (m GetPublishedAppIntegrationsOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppIntegrationsOutput) Type(ctx context.Context) attr.Type {
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
func (m *GetPublishedAppIntegrationsOutput) GetApps(ctx context.Context) ([]GetPublishedAppIntegrationOutput, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []GetPublishedAppIntegrationOutput
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetPublishedAppIntegrationsOutput.
func (m *GetPublishedAppIntegrationsOutput) SetApps(ctx context.Context, v []GetPublishedAppIntegrationOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps types.List `tfsdk:"apps"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *GetPublishedAppsOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppsOutput) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *GetPublishedAppsOutput) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppsOutput) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m GetPublishedAppsOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPublishedAppsOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(PublishedAppOutput{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppsOutput
// only implements ToObjectValue() and Type().
func (m GetPublishedAppsOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppsOutput) Type(ctx context.Context) attr.Type {
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
func (m *GetPublishedAppsOutput) GetApps(ctx context.Context) ([]PublishedAppOutput, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []PublishedAppOutput
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetPublishedAppsOutput.
func (m *GetPublishedAppsOutput) SetApps(ctx context.Context, v []PublishedAppOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
}

type GetServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetServicePrincipalFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalFederationPolicyRequest) {
}

func (to *GetServicePrincipalFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalFederationPolicyRequest) {
}

func (m GetServicePrincipalFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m GetServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type ListAccountFederationPoliciesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAccountFederationPoliciesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAccountFederationPoliciesRequest) {
}

func (to *ListAccountFederationPoliciesRequest) SyncFieldsDuringRead(ctx context.Context, from ListAccountFederationPoliciesRequest) {
}

func (m ListAccountFederationPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListAccountFederationPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListAccountFederationPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountFederationPoliciesRequest
// only implements ToObjectValue() and Type().
func (m ListAccountFederationPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAccountFederationPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListCustomAppIntegrationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCustomAppIntegrationsRequest) {
}

func (to *ListCustomAppIntegrationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListCustomAppIntegrationsRequest) {
}

func (m ListCustomAppIntegrationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["include_creator_username"] = attrs["include_creator_username"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListCustomAppIntegrationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListCustomAppIntegrationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomAppIntegrationsRequest
// only implements ToObjectValue() and Type().
func (m ListCustomAppIntegrationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_creator_username": m.IncludeCreatorUsername,
			"page_size":                m.PageSize,
			"page_token":               m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCustomAppIntegrationsRequest) Type(ctx context.Context) attr.Type {
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

func (to *ListFederationPoliciesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFederationPoliciesResponse) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (to *ListFederationPoliciesResponse) SyncFieldsDuringRead(ctx context.Context, from ListFederationPoliciesResponse) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (m ListFederationPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFederationPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFederationPoliciesResponse
// only implements ToObjectValue() and Type().
func (m ListFederationPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"policies":        m.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFederationPoliciesResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListFederationPoliciesResponse) GetPolicies(ctx context.Context) ([]FederationPolicy, bool) {
	if m.Policies.IsNull() || m.Policies.IsUnknown() {
		return nil, false
	}
	var v []FederationPolicy
	d := m.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListFederationPoliciesResponse.
func (m *ListFederationPoliciesResponse) SetPolicies(ctx context.Context, v []FederationPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Policies = types.ListValueMust(t, vs)
}

type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListOAuthPublishedAppsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOAuthPublishedAppsRequest) {
}

func (to *ListOAuthPublishedAppsRequest) SyncFieldsDuringRead(ctx context.Context, from ListOAuthPublishedAppsRequest) {
}

func (m ListOAuthPublishedAppsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListOAuthPublishedAppsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListOAuthPublishedAppsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOAuthPublishedAppsRequest
// only implements ToObjectValue() and Type().
func (m ListOAuthPublishedAppsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOAuthPublishedAppsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListPublishedAppIntegrationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListPublishedAppIntegrationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPublishedAppIntegrationsRequest) {
}

func (to *ListPublishedAppIntegrationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListPublishedAppIntegrationsRequest) {
}

func (m ListPublishedAppIntegrationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPublishedAppIntegrationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPublishedAppIntegrationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublishedAppIntegrationsRequest
// only implements ToObjectValue() and Type().
func (m ListPublishedAppIntegrationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPublishedAppIntegrationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListServicePrincipalFederationPoliciesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *ListServicePrincipalFederationPoliciesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalFederationPoliciesRequest) {
}

func (to *ListServicePrincipalFederationPoliciesRequest) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalFederationPoliciesRequest) {
}

func (m ListServicePrincipalFederationPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalFederationPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalFederationPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalFederationPoliciesRequest
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalFederationPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":            m.PageSize,
			"page_token":           m.PageToken,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalFederationPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":            types.Int64Type,
			"page_token":           types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type ListServicePrincipalSecretsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`
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
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *ListServicePrincipalSecretsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalSecretsRequest) {
}

func (to *ListServicePrincipalSecretsRequest) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalSecretsRequest) {
}

func (m ListServicePrincipalSecretsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServicePrincipalSecretsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServicePrincipalSecretsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalSecretsRequest
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalSecretsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":            m.PageSize,
			"page_token":           m.PageToken,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalSecretsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":            types.Int64Type,
			"page_token":           types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

type ListServicePrincipalSecretsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of the secrets
	Secrets types.List `tfsdk:"secrets"`
}

func (to *ListServicePrincipalSecretsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalSecretsResponse) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (to *ListServicePrincipalSecretsResponse) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalSecretsResponse) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (m ListServicePrincipalSecretsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListServicePrincipalSecretsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalSecretsResponse
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalSecretsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"secrets":         m.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalSecretsResponse) Type(ctx context.Context) attr.Type {
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
func (m *ListServicePrincipalSecretsResponse) GetSecrets(ctx context.Context) ([]SecretInfo, bool) {
	if m.Secrets.IsNull() || m.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretInfo
	d := m.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListServicePrincipalSecretsResponse.
func (m *ListServicePrincipalSecretsResponse) SetSecrets(ctx context.Context, v []SecretInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["secrets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Secrets = types.ListValueMust(t, vs)
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
	// JWKS format. Most use cases should not need to specify this field. If
	// jwks_uri and jwks_json are both unspecified (recommended), Databricks
	// automatically fetches the public keys from your issuer’s well known
	// endpoint. Databricks strongly recommends relying on your issuer’s well
	// known endpoint for discovering public keys.
	JwksJson types.String `tfsdk:"jwks_json"`
	// URL of the public keys used to validate the signature of federated
	// tokens, in JWKS format. Most use cases should not need to specify this
	// field. If jwks_uri and jwks_json are both unspecified (recommended),
	// Databricks automatically fetches the public keys from your issuer’s
	// well known endpoint. Databricks strongly recommends relying on your
	// issuer’s well known endpoint for discovering public keys.
	JwksUri types.String `tfsdk:"jwks_uri"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. Must be specified for service principal federation
	// policies. Must not be specified for account federation policies.
	Subject types.String `tfsdk:"subject"`
	// The claim that contains the subject of the token. If unspecified, the
	// default value is 'sub'.
	SubjectClaim types.String `tfsdk:"subject_claim"`
}

func (to *OidcFederationPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OidcFederationPolicy) {
	if !from.Audiences.IsNull() && !from.Audiences.IsUnknown() && to.Audiences.IsNull() && len(from.Audiences.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Audiences, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Audiences = from.Audiences
	}
}

func (to *OidcFederationPolicy) SyncFieldsDuringRead(ctx context.Context, from OidcFederationPolicy) {
	if !from.Audiences.IsNull() && !from.Audiences.IsUnknown() && to.Audiences.IsNull() && len(from.Audiences.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Audiences, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Audiences = from.Audiences
	}
}

func (m OidcFederationPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["audiences"] = attrs["audiences"].SetOptional()
	attrs["issuer"] = attrs["issuer"].SetOptional()
	attrs["jwks_json"] = attrs["jwks_json"].SetOptional()
	attrs["jwks_uri"] = attrs["jwks_uri"].SetOptional()
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
func (m OidcFederationPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"audiences": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OidcFederationPolicy
// only implements ToObjectValue() and Type().
func (m OidcFederationPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"audiences":     m.Audiences,
			"issuer":        m.Issuer,
			"jwks_json":     m.JwksJson,
			"jwks_uri":      m.JwksUri,
			"subject":       m.Subject,
			"subject_claim": m.SubjectClaim,
		})
}

// Type implements basetypes.ObjectValuable.
func (m OidcFederationPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"audiences": basetypes.ListType{
				ElemType: types.StringType,
			},
			"issuer":        types.StringType,
			"jwks_json":     types.StringType,
			"jwks_uri":      types.StringType,
			"subject":       types.StringType,
			"subject_claim": types.StringType,
		},
	}
}

// GetAudiences returns the value of the Audiences field in OidcFederationPolicy as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *OidcFederationPolicy) GetAudiences(ctx context.Context) ([]types.String, bool) {
	if m.Audiences.IsNull() || m.Audiences.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Audiences.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAudiences sets the value of the Audiences field in OidcFederationPolicy.
func (m *OidcFederationPolicy) SetAudiences(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["audiences"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Audiences = types.ListValueMust(t, vs)
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

func (to *PublishedAppOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishedAppOutput) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (to *PublishedAppOutput) SyncFieldsDuringRead(ctx context.Context, from PublishedAppOutput) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (m PublishedAppOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PublishedAppOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls": reflect.TypeOf(types.String{}),
		"scopes":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishedAppOutput
// only implements ToObjectValue() and Type().
func (m PublishedAppOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":                 m.AppId,
			"client_id":              m.ClientId,
			"description":            m.Description,
			"is_confidential_client": m.IsConfidentialClient,
			"name":                   m.Name,
			"redirect_urls":          m.RedirectUrls,
			"scopes":                 m.Scopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublishedAppOutput) Type(ctx context.Context) attr.Type {
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
func (m *PublishedAppOutput) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if m.RedirectUrls.IsNull() || m.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in PublishedAppOutput.
func (m *PublishedAppOutput) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in PublishedAppOutput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PublishedAppOutput) GetScopes(ctx context.Context) ([]types.String, bool) {
	if m.Scopes.IsNull() || m.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in PublishedAppOutput.
func (m *PublishedAppOutput) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime types.String `tfsdk:"create_time"`
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	ExpireTime types.String `tfsdk:"expire_time"`
	// ID of the secret
	Id types.String `tfsdk:"id"`
	// Secret Hash
	SecretHash types.String `tfsdk:"secret_hash"`
	// Status of the secret
	Status types.String `tfsdk:"status"`
	// UTC time when the secret was updated
	UpdateTime types.String `tfsdk:"update_time"`
}

func (to *SecretInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecretInfo) {
}

func (to *SecretInfo) SyncFieldsDuringRead(ctx context.Context, from SecretInfo) {
}

func (m SecretInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["expire_time"] = attrs["expire_time"].SetOptional()
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
func (m SecretInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretInfo
// only implements ToObjectValue() and Type().
func (m SecretInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"expire_time": m.ExpireTime,
			"id":          m.Id,
			"secret_hash": m.SecretHash,
			"status":      m.Status,
			"update_time": m.UpdateTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SecretInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"expire_time": types.StringType,
			"id":          types.StringType,
			"secret_hash": types.StringType,
			"status":      types.StringType,
			"update_time": types.StringType,
		},
	}
}

type TokenAccessPolicy struct {
	// Absolute OAuth session TTL in minutes. Effective only when the single-use
	// refresh token feature is enabled. This is the absolute TTL of all refresh
	// tokens issued in one OAuth session. When a new refresh token is issued
	// during refresh token rotation, it will inherit the same absolute TTL as
	// the old refresh token. In other words, this represents the maximum amount
	// of time a user can stay logged in without re-authenticating.
	AbsoluteSessionLifetimeInMinutes types.Int64 `tfsdk:"absolute_session_lifetime_in_minutes"`
	// access token time to live in minutes
	AccessTokenTtlInMinutes types.Int64 `tfsdk:"access_token_ttl_in_minutes"`
	// Whether to enable single-use refresh tokens (refresh token rotation). If
	// this feature is enabled, upon successfully getting a new access token
	// using a refresh token, Databricks will issue a new refresh token along
	// with the access token in the response and invalidate the old refresh
	// token. The client should use the new refresh token to get access tokens
	// in future requests.
	EnableSingleUseRefreshTokens types.Bool `tfsdk:"enable_single_use_refresh_tokens"`
	// Refresh token time to live in minutes. When single-use refresh tokens are
	// enabled, this represents the TTL of an individual refresh token. If the
	// refresh token is used before it expires, a new one is issued with a
	// renewed individual TTL.
	RefreshTokenTtlInMinutes types.Int64 `tfsdk:"refresh_token_ttl_in_minutes"`
}

func (to *TokenAccessPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenAccessPolicy) {
}

func (to *TokenAccessPolicy) SyncFieldsDuringRead(ctx context.Context, from TokenAccessPolicy) {
}

func (m TokenAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["absolute_session_lifetime_in_minutes"] = attrs["absolute_session_lifetime_in_minutes"].SetOptional()
	attrs["access_token_ttl_in_minutes"] = attrs["access_token_ttl_in_minutes"].SetOptional()
	attrs["enable_single_use_refresh_tokens"] = attrs["enable_single_use_refresh_tokens"].SetOptional()
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
func (m TokenAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessPolicy
// only implements ToObjectValue() and Type().
func (m TokenAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"absolute_session_lifetime_in_minutes": m.AbsoluteSessionLifetimeInMinutes,
			"access_token_ttl_in_minutes":          m.AccessTokenTtlInMinutes,
			"enable_single_use_refresh_tokens":     m.EnableSingleUseRefreshTokens,
			"refresh_token_ttl_in_minutes":         m.RefreshTokenTtlInMinutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenAccessPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"absolute_session_lifetime_in_minutes": types.Int64Type,
			"access_token_ttl_in_minutes":          types.Int64Type,
			"enable_single_use_refresh_tokens":     types.BoolType,
			"refresh_token_ttl_in_minutes":         types.Int64Type,
		},
	}
}

type UpdateAccountFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateAccountFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAccountFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *UpdateAccountFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAccountFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m UpdateAccountFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAccountFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAccountFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m UpdateAccountFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":      m.Policy,
			"policy_id":   m.PolicyId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAccountFederationPolicyRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateAccountFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v FederationPolicy
	d := m.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicy sets the value of the Policy field in UpdateAccountFederationPolicyRequest.
func (m *UpdateAccountFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	m.Policy = vs
}

type UpdateCustomAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls types.List `tfsdk:"redirect_urls"`
	// List of OAuth scopes to be updated in the custom OAuth app integration,
	// similar to redirect URIs this will fully replace the existing values
	// instead of appending
	Scopes types.List `tfsdk:"scopes"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes types.List `tfsdk:"user_authorized_scopes"`
}

func (to *UpdateCustomAppIntegration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCustomAppIntegration) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				// Recursively sync the fields of TokenAccessPolicy
				toTokenAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
	if !from.UserAuthorizedScopes.IsNull() && !from.UserAuthorizedScopes.IsUnknown() && to.UserAuthorizedScopes.IsNull() && len(from.UserAuthorizedScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserAuthorizedScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserAuthorizedScopes = from.UserAuthorizedScopes
	}
}

func (to *UpdateCustomAppIntegration) SyncFieldsDuringRead(ctx context.Context, from UpdateCustomAppIntegration) {
	if !from.RedirectUrls.IsNull() && !from.RedirectUrls.IsUnknown() && to.RedirectUrls.IsNull() && len(from.RedirectUrls.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RedirectUrls, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RedirectUrls = from.RedirectUrls
	}
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
	if !from.UserAuthorizedScopes.IsNull() && !from.UserAuthorizedScopes.IsUnknown() && to.UserAuthorizedScopes.IsNull() && len(from.UserAuthorizedScopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for UserAuthorizedScopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.UserAuthorizedScopes = from.UserAuthorizedScopes
	}
}

func (m UpdateCustomAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["user_authorized_scopes"] = attrs["user_authorized_scopes"].SetOptional()
	attrs["integration_id"] = attrs["integration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCustomAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":          reflect.TypeOf(types.String{}),
		"scopes":                 reflect.TypeOf(types.String{}),
		"token_access_policy":    reflect.TypeOf(TokenAccessPolicy{}),
		"user_authorized_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomAppIntegration
// only implements ToObjectValue() and Type().
func (m UpdateCustomAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id":         m.IntegrationId,
			"redirect_urls":          m.RedirectUrls,
			"scopes":                 m.Scopes,
			"token_access_policy":    m.TokenAccessPolicy,
			"user_authorized_scopes": m.UserAuthorizedScopes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCustomAppIntegration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
			"redirect_urls": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_access_policy": TokenAccessPolicy{}.Type(ctx),
			"user_authorized_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in UpdateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
	if m.RedirectUrls.IsNull() || m.RedirectUrls.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.RedirectUrls.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRedirectUrls sets the value of the RedirectUrls field in UpdateCustomAppIntegration.
func (m *UpdateCustomAppIntegration) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in UpdateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration) GetScopes(ctx context.Context) ([]types.String, bool) {
	if m.Scopes.IsNull() || m.Scopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Scopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetScopes sets the value of the Scopes field in UpdateCustomAppIntegration.
func (m *UpdateCustomAppIntegration) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in UpdateCustomAppIntegration as
// a TokenAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v TokenAccessPolicy
	d := m.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdateCustomAppIntegration.
func (m *UpdateCustomAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TokenAccessPolicy = vs
}

// GetUserAuthorizedScopes returns the value of the UserAuthorizedScopes field in UpdateCustomAppIntegration as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration) GetUserAuthorizedScopes(ctx context.Context) ([]types.String, bool) {
	if m.UserAuthorizedScopes.IsNull() || m.UserAuthorizedScopes.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.UserAuthorizedScopes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUserAuthorizedScopes sets the value of the UserAuthorizedScopes field in UpdateCustomAppIntegration.
func (m *UpdateCustomAppIntegration) SetUserAuthorizedScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_authorized_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserAuthorizedScopes = types.ListValueMust(t, vs)
}

type UpdateCustomAppIntegrationOutput struct {
}

func (to *UpdateCustomAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCustomAppIntegrationOutput) {
}

func (to *UpdateCustomAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from UpdateCustomAppIntegrationOutput) {
}

func (m UpdateCustomAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCustomAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m UpdateCustomAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCustomAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdatePublishedAppIntegration struct {
	IntegrationId types.String `tfsdk:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy types.Object `tfsdk:"token_access_policy"`
}

func (to *UpdatePublishedAppIntegration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePublishedAppIntegration) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				// Recursively sync the fields of TokenAccessPolicy
				toTokenAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (to *UpdatePublishedAppIntegration) SyncFieldsDuringRead(ctx context.Context, from UpdatePublishedAppIntegration) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (m UpdatePublishedAppIntegration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["integration_id"] = attrs["integration_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePublishedAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePublishedAppIntegration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePublishedAppIntegration
// only implements ToObjectValue() and Type().
func (m UpdatePublishedAppIntegration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id":      m.IntegrationId,
			"token_access_policy": m.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePublishedAppIntegration) Type(ctx context.Context) attr.Type {
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
func (m *UpdatePublishedAppIntegration) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy, bool) {
	var e TokenAccessPolicy
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v TokenAccessPolicy
	d := m.TokenAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdatePublishedAppIntegration.
func (m *UpdatePublishedAppIntegration) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.TokenAccessPolicy = vs
}

type UpdatePublishedAppIntegrationOutput struct {
}

func (to *UpdatePublishedAppIntegrationOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePublishedAppIntegrationOutput) {
}

func (to *UpdatePublishedAppIntegrationOutput) SyncFieldsDuringRead(ctx context.Context, from UpdatePublishedAppIntegrationOutput) {
}

func (m UpdatePublishedAppIntegrationOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePublishedAppIntegrationOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePublishedAppIntegrationOutput
// only implements ToObjectValue() and Type().
func (m UpdatePublishedAppIntegrationOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePublishedAppIntegrationOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateServicePrincipalFederationPolicyRequest struct {
	Policy types.Object `tfsdk:"policy"`
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateServicePrincipalFederationPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateServicePrincipalFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				// Recursively sync the fields of Policy
				toPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (to *UpdateServicePrincipalFederationPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateServicePrincipalFederationPolicyRequest) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m UpdateServicePrincipalFederationPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["service_principal_id"] = attrs["service_principal_id"].SetRequired()
	attrs["policy_id"] = attrs["policy_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateServicePrincipalFederationPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateServicePrincipalFederationPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalFederationPolicyRequest
// only implements ToObjectValue() and Type().
func (m UpdateServicePrincipalFederationPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":               m.Policy,
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
			"update_mask":          m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateServicePrincipalFederationPolicyRequest) Type(ctx context.Context) attr.Type {
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
func (m *UpdateServicePrincipalFederationPolicyRequest) GetPolicy(ctx context.Context) (FederationPolicy, bool) {
	var e FederationPolicy
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v FederationPolicy
	d := m.Policy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicy sets the value of the Policy field in UpdateServicePrincipalFederationPolicyRequest.
func (m *UpdateServicePrincipalFederationPolicyRequest) SetPolicy(ctx context.Context, v FederationPolicy) {
	vs := v.ToObjectValue(ctx)
	m.Policy = vs
}
