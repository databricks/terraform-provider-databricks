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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type CreateAccountFederationPolicyRequest_SdkV2 struct {
	Policy types.List `tfsdk:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId types.String `tfsdk:"-"`
}

func (to *CreateAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateAccountFederationPolicyRequest_SdkV2) {
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

func (to *CreateAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateAccountFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m CreateAccountFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateAccountFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateAccountFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateAccountFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":    m.Policy,
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateAccountFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
			"policy_id": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateAccountFederationPolicyRequest_SdkV2 as
// a FederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateAccountFederationPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (FederationPolicy_SdkV2, bool) {
	var e FederationPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in CreateAccountFederationPolicyRequest_SdkV2.
func (m *CreateAccountFederationPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v FederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

type CreateCustomAppIntegration_SdkV2 struct {
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
	TokenAccessPolicy types.List `tfsdk:"token_access_policy"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes types.List `tfsdk:"user_authorized_scopes"`
}

func (to *CreateCustomAppIntegration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomAppIntegration_SdkV2) {
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

func (to *CreateCustomAppIntegration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCustomAppIntegration_SdkV2) {
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

func (m CreateCustomAppIntegration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["confidential"] = attrs["confidential"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateCustomAppIntegration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":          reflect.TypeOf(types.String{}),
		"scopes":                 reflect.TypeOf(types.String{}),
		"token_access_policy":    reflect.TypeOf(TokenAccessPolicy_SdkV2{}),
		"user_authorized_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomAppIntegration_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCustomAppIntegration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateCustomAppIntegration_SdkV2) Type(ctx context.Context) attr.Type {
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
			"token_access_policy": basetypes.ListType{
				ElemType: TokenAccessPolicy_SdkV2{}.Type(ctx),
			},
			"user_authorized_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in CreateCustomAppIntegration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration_SdkV2) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
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

// SetRedirectUrls sets the value of the RedirectUrls field in CreateCustomAppIntegration_SdkV2.
func (m *CreateCustomAppIntegration_SdkV2) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in CreateCustomAppIntegration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in CreateCustomAppIntegration_SdkV2.
func (m *CreateCustomAppIntegration_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in CreateCustomAppIntegration_SdkV2 as
// a TokenAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration_SdkV2) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy_SdkV2, bool) {
	var e TokenAccessPolicy_SdkV2
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy_SdkV2
	d := m.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreateCustomAppIntegration_SdkV2.
func (m *CreateCustomAppIntegration_SdkV2) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	m.TokenAccessPolicy = types.ListValueMust(t, vs)
}

// GetUserAuthorizedScopes returns the value of the UserAuthorizedScopes field in CreateCustomAppIntegration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateCustomAppIntegration_SdkV2) GetUserAuthorizedScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserAuthorizedScopes sets the value of the UserAuthorizedScopes field in CreateCustomAppIntegration_SdkV2.
func (m *CreateCustomAppIntegration_SdkV2) SetUserAuthorizedScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_authorized_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserAuthorizedScopes = types.ListValueMust(t, vs)
}

type CreateCustomAppIntegrationOutput_SdkV2 struct {
	// OAuth client-id generated by the Databricks
	ClientId types.String `tfsdk:"client_id"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret types.String `tfsdk:"client_secret"`
	// Unique integration id for the custom OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
}

func (to *CreateCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateCustomAppIntegrationOutput_SdkV2) {
}

func (to *CreateCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateCustomAppIntegrationOutput_SdkV2) {
}

func (m CreateCustomAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateCustomAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateCustomAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateCustomAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_id":      m.ClientId,
			"client_secret":  m.ClientSecret,
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateCustomAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client_id":      types.StringType,
			"client_secret":  types.StringType,
			"integration_id": types.StringType,
		},
	}
}

type CreatePublishedAppIntegration_SdkV2 struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId types.String `tfsdk:"app_id"`
	// Token access policy
	TokenAccessPolicy types.List `tfsdk:"token_access_policy"`
}

func (to *CreatePublishedAppIntegration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePublishedAppIntegration_SdkV2) {
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

func (to *CreatePublishedAppIntegration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePublishedAppIntegration_SdkV2) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (m CreatePublishedAppIntegration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_id"] = attrs["app_id"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePublishedAppIntegration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePublishedAppIntegration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePublishedAppIntegration_SdkV2
// only implements ToObjectValue() and Type().
func (m CreatePublishedAppIntegration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":              m.AppId,
			"token_access_policy": m.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePublishedAppIntegration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_id": types.StringType,
			"token_access_policy": basetypes.ListType{
				ElemType: TokenAccessPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in CreatePublishedAppIntegration_SdkV2 as
// a TokenAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePublishedAppIntegration_SdkV2) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy_SdkV2, bool) {
	var e TokenAccessPolicy_SdkV2
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy_SdkV2
	d := m.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in CreatePublishedAppIntegration_SdkV2.
func (m *CreatePublishedAppIntegration_SdkV2) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	m.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type CreatePublishedAppIntegrationOutput_SdkV2 struct {
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
}

func (to *CreatePublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePublishedAppIntegrationOutput_SdkV2) {
}

func (to *CreatePublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePublishedAppIntegrationOutput_SdkV2) {
}

func (m CreatePublishedAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePublishedAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePublishedAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m CreatePublishedAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePublishedAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type CreateServicePrincipalFederationPolicyRequest_SdkV2 struct {
	Policy types.List `tfsdk:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *CreateServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalFederationPolicyRequest_SdkV2) {
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

func (to *CreateServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m CreateServicePrincipalFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateServicePrincipalFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":               m.Policy,
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

// GetPolicy returns the value of the Policy field in CreateServicePrincipalFederationPolicyRequest_SdkV2 as
// a FederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServicePrincipalFederationPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (FederationPolicy_SdkV2, bool) {
	var e FederationPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in CreateServicePrincipalFederationPolicyRequest_SdkV2.
func (m *CreateServicePrincipalFederationPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v FederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

type CreateServicePrincipalSecretRequest_SdkV2 struct {
	// The lifetime of the secret in seconds. If this parameter is not provided,
	// the secret will have a default lifetime of 730 days (63072000s).
	Lifetime types.String `tfsdk:"lifetime"`
	// The service principal ID.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *CreateServicePrincipalSecretRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalSecretRequest_SdkV2) {
}

func (to *CreateServicePrincipalSecretRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalSecretRequest_SdkV2) {
}

func (m CreateServicePrincipalSecretRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateServicePrincipalSecretRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalSecretRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalSecretRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"lifetime":             m.Lifetime,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServicePrincipalSecretRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"lifetime":             types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

type CreateServicePrincipalSecretResponse_SdkV2 struct {
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

func (to *CreateServicePrincipalSecretResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServicePrincipalSecretResponse_SdkV2) {
}

func (to *CreateServicePrincipalSecretResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateServicePrincipalSecretResponse_SdkV2) {
}

func (m CreateServicePrincipalSecretResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateServicePrincipalSecretResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServicePrincipalSecretResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateServicePrincipalSecretResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreateServicePrincipalSecretResponse_SdkV2) Type(ctx context.Context) attr.Type {
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

type DeleteAccountFederationPolicyRequest_SdkV2 struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *DeleteAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountFederationPolicyRequest_SdkV2) {
}

func (to *DeleteAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountFederationPolicyRequest_SdkV2) {
}

func (m DeleteAccountFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAccountFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAccountFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type DeleteCustomAppIntegrationOutput_SdkV2 struct {
}

func (to *DeleteCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCustomAppIntegrationOutput_SdkV2) {
}

func (to *DeleteCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCustomAppIntegrationOutput_SdkV2) {
}

func (m DeleteCustomAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteCustomAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCustomAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCustomAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeleteCustomAppIntegrationRequest_SdkV2 struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (to *DeleteCustomAppIntegrationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteCustomAppIntegrationRequest_SdkV2) {
}

func (to *DeleteCustomAppIntegrationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteCustomAppIntegrationRequest_SdkV2) {
}

func (m DeleteCustomAppIntegrationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteCustomAppIntegrationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteCustomAppIntegrationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteCustomAppIntegrationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteCustomAppIntegrationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type DeletePublishedAppIntegrationOutput_SdkV2 struct {
}

func (to *DeletePublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePublishedAppIntegrationOutput_SdkV2) {
}

func (to *DeletePublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePublishedAppIntegrationOutput_SdkV2) {
}

func (m DeletePublishedAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePublishedAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePublishedAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePublishedAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePublishedAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type DeletePublishedAppIntegrationRequest_SdkV2 struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (to *DeletePublishedAppIntegrationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePublishedAppIntegrationRequest_SdkV2) {
}

func (to *DeletePublishedAppIntegrationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePublishedAppIntegrationRequest_SdkV2) {
}

func (m DeletePublishedAppIntegrationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeletePublishedAppIntegrationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePublishedAppIntegrationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePublishedAppIntegrationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePublishedAppIntegrationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type DeleteServicePrincipalFederationPolicyRequest_SdkV2 struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *DeleteServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalFederationPolicyRequest_SdkV2) {
}

func (to *DeleteServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalFederationPolicyRequest_SdkV2) {
}

func (m DeleteServicePrincipalFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteServicePrincipalFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type DeleteServicePrincipalSecretRequest_SdkV2 struct {
	// The secret ID.
	SecretId types.String `tfsdk:"-"`
	// The service principal ID.
	ServicePrincipalId types.String `tfsdk:"-"`
}

func (to *DeleteServicePrincipalSecretRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServicePrincipalSecretRequest_SdkV2) {
}

func (to *DeleteServicePrincipalSecretRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServicePrincipalSecretRequest_SdkV2) {
}

func (m DeleteServicePrincipalSecretRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteServicePrincipalSecretRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServicePrincipalSecretRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteServicePrincipalSecretRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"secret_id":            m.SecretId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServicePrincipalSecretRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"secret_id":            types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

type FederationPolicy_SdkV2 struct {
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

	OidcPolicy types.List `tfsdk:"oidc_policy"`
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

func (to *FederationPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FederationPolicy_SdkV2) {
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

func (to *FederationPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FederationPolicy_SdkV2) {
	if !from.OidcPolicy.IsNull() && !from.OidcPolicy.IsUnknown() {
		if toOidcPolicy, ok := to.GetOidcPolicy(ctx); ok {
			if fromOidcPolicy, ok := from.GetOidcPolicy(ctx); ok {
				toOidcPolicy.SyncFieldsDuringRead(ctx, fromOidcPolicy)
				to.SetOidcPolicy(ctx, toOidcPolicy)
			}
		}
	}
}

func (m FederationPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["create_time"] = attrs["create_time"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["name"] = attrs["name"].SetComputed()
	attrs["name"] = attrs["name"].(tfschema.StringAttributeBuilder).AddPlanModifier(stringplanmodifier.UseStateForUnknown()).(tfschema.AttributeBuilder)
	attrs["oidc_policy"] = attrs["oidc_policy"].SetOptional()
	attrs["oidc_policy"] = attrs["oidc_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m FederationPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"oidc_policy": reflect.TypeOf(OidcFederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FederationPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m FederationPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m FederationPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"description": types.StringType,
			"name":        types.StringType,
			"oidc_policy": basetypes.ListType{
				ElemType: OidcFederationPolicy_SdkV2{}.Type(ctx),
			},
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
			"uid":                  types.StringType,
			"update_time":          types.StringType,
		},
	}
}

// GetOidcPolicy returns the value of the OidcPolicy field in FederationPolicy_SdkV2 as
// a OidcFederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *FederationPolicy_SdkV2) GetOidcPolicy(ctx context.Context) (OidcFederationPolicy_SdkV2, bool) {
	var e OidcFederationPolicy_SdkV2
	if m.OidcPolicy.IsNull() || m.OidcPolicy.IsUnknown() {
		return e, false
	}
	var v []OidcFederationPolicy_SdkV2
	d := m.OidcPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOidcPolicy sets the value of the OidcPolicy field in FederationPolicy_SdkV2.
func (m *FederationPolicy_SdkV2) SetOidcPolicy(ctx context.Context, v OidcFederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["oidc_policy"]
	m.OidcPolicy = types.ListValueMust(t, vs)
}

type GetAccountFederationPolicyRequest_SdkV2 struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
}

func (to *GetAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAccountFederationPolicyRequest_SdkV2) {
}

func (to *GetAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAccountFederationPolicyRequest_SdkV2) {
}

func (m GetAccountFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAccountFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAccountFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id": m.PolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAccountFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id": types.StringType,
		},
	}
}

type GetCustomAppIntegrationOutput_SdkV2 struct {
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
	TokenAccessPolicy types.List `tfsdk:"token_access_policy"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes types.List `tfsdk:"user_authorized_scopes"`
}

func (to *GetCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomAppIntegrationOutput_SdkV2) {
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

func (to *GetCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCustomAppIntegrationOutput_SdkV2) {
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

func (m GetCustomAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
	attrs["token_access_policy"] = attrs["token_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m GetCustomAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":          reflect.TypeOf(types.String{}),
		"scopes":                 reflect.TypeOf(types.String{}),
		"token_access_policy":    reflect.TypeOf(TokenAccessPolicy_SdkV2{}),
		"user_authorized_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCustomAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetCustomAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
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
			"token_access_policy": basetypes.ListType{
				ElemType: TokenAccessPolicy_SdkV2{}.Type(ctx),
			},
			"user_authorized_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in GetCustomAppIntegrationOutput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput_SdkV2) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
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

// SetRedirectUrls sets the value of the RedirectUrls field in GetCustomAppIntegrationOutput_SdkV2.
func (m *GetCustomAppIntegrationOutput_SdkV2) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in GetCustomAppIntegrationOutput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in GetCustomAppIntegrationOutput_SdkV2.
func (m *GetCustomAppIntegrationOutput_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput_SdkV2 as
// a TokenAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput_SdkV2) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy_SdkV2, bool) {
	var e TokenAccessPolicy_SdkV2
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy_SdkV2
	d := m.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetCustomAppIntegrationOutput_SdkV2.
func (m *GetCustomAppIntegrationOutput_SdkV2) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	m.TokenAccessPolicy = types.ListValueMust(t, vs)
}

// GetUserAuthorizedScopes returns the value of the UserAuthorizedScopes field in GetCustomAppIntegrationOutput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationOutput_SdkV2) GetUserAuthorizedScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserAuthorizedScopes sets the value of the UserAuthorizedScopes field in GetCustomAppIntegrationOutput_SdkV2.
func (m *GetCustomAppIntegrationOutput_SdkV2) SetUserAuthorizedScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_authorized_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserAuthorizedScopes = types.ListValueMust(t, vs)
}

type GetCustomAppIntegrationRequest_SdkV2 struct {
	// The OAuth app integration ID.
	IntegrationId types.String `tfsdk:"-"`
}

func (to *GetCustomAppIntegrationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomAppIntegrationRequest_SdkV2) {
}

func (to *GetCustomAppIntegrationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCustomAppIntegrationRequest_SdkV2) {
}

func (m GetCustomAppIntegrationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCustomAppIntegrationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCustomAppIntegrationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCustomAppIntegrationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type GetCustomAppIntegrationsOutput_SdkV2 struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps types.List `tfsdk:"apps"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *GetCustomAppIntegrationsOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCustomAppIntegrationsOutput_SdkV2) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *GetCustomAppIntegrationsOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCustomAppIntegrationsOutput_SdkV2) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m GetCustomAppIntegrationsOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCustomAppIntegrationsOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(GetCustomAppIntegrationOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCustomAppIntegrationsOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCustomAppIntegrationsOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCustomAppIntegrationsOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: GetCustomAppIntegrationOutput_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in GetCustomAppIntegrationsOutput_SdkV2 as
// a slice of GetCustomAppIntegrationOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetCustomAppIntegrationsOutput_SdkV2) GetApps(ctx context.Context) ([]GetCustomAppIntegrationOutput_SdkV2, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []GetCustomAppIntegrationOutput_SdkV2
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetCustomAppIntegrationsOutput_SdkV2.
func (m *GetCustomAppIntegrationsOutput_SdkV2) SetApps(ctx context.Context, v []GetCustomAppIntegrationOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
}

type GetPublishedAppIntegrationOutput_SdkV2 struct {
	// App-id of the published app integration
	AppId types.String `tfsdk:"app_id"`

	CreateTime types.String `tfsdk:"create_time"`

	CreatedBy types.Int64 `tfsdk:"created_by"`
	// Unique integration id for the published OAuth app
	IntegrationId types.String `tfsdk:"integration_id"`
	// Display name of the published OAuth app
	Name types.String `tfsdk:"name"`
	// Token access policy
	TokenAccessPolicy types.List `tfsdk:"token_access_policy"`
}

func (to *GetPublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppIntegrationOutput_SdkV2) {
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

func (to *GetPublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppIntegrationOutput_SdkV2) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (m GetPublishedAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_id"] = attrs["app_id"].SetOptional()
	attrs["create_time"] = attrs["create_time"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["integration_id"] = attrs["integration_id"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPublishedAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPublishedAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GetPublishedAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_id":         types.StringType,
			"create_time":    types.StringType,
			"created_by":     types.Int64Type,
			"integration_id": types.StringType,
			"name":           types.StringType,
			"token_access_policy": basetypes.ListType{
				ElemType: TokenAccessPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in GetPublishedAppIntegrationOutput_SdkV2 as
// a TokenAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPublishedAppIntegrationOutput_SdkV2) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy_SdkV2, bool) {
	var e TokenAccessPolicy_SdkV2
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy_SdkV2
	d := m.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in GetPublishedAppIntegrationOutput_SdkV2.
func (m *GetPublishedAppIntegrationOutput_SdkV2) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	m.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type GetPublishedAppIntegrationRequest_SdkV2 struct {
	IntegrationId types.String `tfsdk:"-"`
}

func (to *GetPublishedAppIntegrationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppIntegrationRequest_SdkV2) {
}

func (to *GetPublishedAppIntegrationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppIntegrationRequest_SdkV2) {
}

func (m GetPublishedAppIntegrationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPublishedAppIntegrationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPublishedAppIntegrationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id": m.IntegrationId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppIntegrationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
		},
	}
}

type GetPublishedAppIntegrationsOutput_SdkV2 struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps types.List `tfsdk:"apps"`

	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *GetPublishedAppIntegrationsOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppIntegrationsOutput_SdkV2) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *GetPublishedAppIntegrationsOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppIntegrationsOutput_SdkV2) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m GetPublishedAppIntegrationsOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPublishedAppIntegrationsOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(GetPublishedAppIntegrationOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppIntegrationsOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPublishedAppIntegrationsOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppIntegrationsOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: GetPublishedAppIntegrationOutput_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in GetPublishedAppIntegrationsOutput_SdkV2 as
// a slice of GetPublishedAppIntegrationOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPublishedAppIntegrationsOutput_SdkV2) GetApps(ctx context.Context) ([]GetPublishedAppIntegrationOutput_SdkV2, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []GetPublishedAppIntegrationOutput_SdkV2
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetPublishedAppIntegrationsOutput_SdkV2.
func (m *GetPublishedAppIntegrationsOutput_SdkV2) SetApps(ctx context.Context, v []GetPublishedAppIntegrationOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
}

type GetPublishedAppsOutput_SdkV2 struct {
	// List of Published OAuth Apps.
	Apps types.List `tfsdk:"apps"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *GetPublishedAppsOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPublishedAppsOutput_SdkV2) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (to *GetPublishedAppsOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPublishedAppsOutput_SdkV2) {
	if !from.Apps.IsNull() && !from.Apps.IsUnknown() && to.Apps.IsNull() && len(from.Apps.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Apps, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Apps = from.Apps
	}
}

func (m GetPublishedAppsOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPublishedAppsOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"apps": reflect.TypeOf(PublishedAppOutput_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPublishedAppsOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPublishedAppsOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"apps":            m.Apps,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPublishedAppsOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"apps": basetypes.ListType{
				ElemType: PublishedAppOutput_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetApps returns the value of the Apps field in GetPublishedAppsOutput_SdkV2 as
// a slice of PublishedAppOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetPublishedAppsOutput_SdkV2) GetApps(ctx context.Context) ([]PublishedAppOutput_SdkV2, bool) {
	if m.Apps.IsNull() || m.Apps.IsUnknown() {
		return nil, false
	}
	var v []PublishedAppOutput_SdkV2
	d := m.Apps.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApps sets the value of the Apps field in GetPublishedAppsOutput_SdkV2.
func (m *GetPublishedAppsOutput_SdkV2) SetApps(ctx context.Context, v []PublishedAppOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["apps"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Apps = types.ListValueMust(t, vs)
}

type GetServicePrincipalFederationPolicyRequest_SdkV2 struct {
	// The identifier for the federation policy.
	PolicyId types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *GetServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServicePrincipalFederationPolicyRequest_SdkV2) {
}

func (to *GetServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServicePrincipalFederationPolicyRequest_SdkV2) {
}

func (m GetServicePrincipalFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetServicePrincipalFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServicePrincipalFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServicePrincipalFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy_id":            m.PolicyId,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServicePrincipalFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type ListAccountFederationPoliciesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListAccountFederationPoliciesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListAccountFederationPoliciesRequest_SdkV2) {
}

func (to *ListAccountFederationPoliciesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListAccountFederationPoliciesRequest_SdkV2) {
}

func (m ListAccountFederationPoliciesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListAccountFederationPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListAccountFederationPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListAccountFederationPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListAccountFederationPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListCustomAppIntegrationsRequest_SdkV2 struct {
	IncludeCreatorUsername types.Bool `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListCustomAppIntegrationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListCustomAppIntegrationsRequest_SdkV2) {
}

func (to *ListCustomAppIntegrationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListCustomAppIntegrationsRequest_SdkV2) {
}

func (m ListCustomAppIntegrationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListCustomAppIntegrationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListCustomAppIntegrationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListCustomAppIntegrationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"include_creator_username": m.IncludeCreatorUsername,
			"page_size":                m.PageSize,
			"page_token":               m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListCustomAppIntegrationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"include_creator_username": types.BoolType,
			"page_size":                types.Int64Type,
			"page_token":               types.StringType,
		},
	}
}

type ListFederationPoliciesResponse_SdkV2 struct {
	NextPageToken types.String `tfsdk:"next_page_token"`

	Policies types.List `tfsdk:"policies"`
}

func (to *ListFederationPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListFederationPoliciesResponse_SdkV2) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (to *ListFederationPoliciesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListFederationPoliciesResponse_SdkV2) {
	if !from.Policies.IsNull() && !from.Policies.IsUnknown() && to.Policies.IsNull() && len(from.Policies.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Policies, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Policies = from.Policies
	}
}

func (m ListFederationPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListFederationPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policies": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListFederationPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListFederationPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"policies":        m.Policies,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListFederationPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"policies": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPolicies returns the value of the Policies field in ListFederationPoliciesResponse_SdkV2 as
// a slice of FederationPolicy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListFederationPoliciesResponse_SdkV2) GetPolicies(ctx context.Context) ([]FederationPolicy_SdkV2, bool) {
	if m.Policies.IsNull() || m.Policies.IsUnknown() {
		return nil, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policies.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicies sets the value of the Policies field in ListFederationPoliciesResponse_SdkV2.
func (m *ListFederationPoliciesResponse_SdkV2) SetPolicies(ctx context.Context, v []FederationPolicy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policies"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Policies = types.ListValueMust(t, vs)
}

type ListOAuthPublishedAppsRequest_SdkV2 struct {
	// The max number of OAuth published apps to return in one page.
	PageSize types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListOAuthPublishedAppsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListOAuthPublishedAppsRequest_SdkV2) {
}

func (to *ListOAuthPublishedAppsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListOAuthPublishedAppsRequest_SdkV2) {
}

func (m ListOAuthPublishedAppsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListOAuthPublishedAppsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListOAuthPublishedAppsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListOAuthPublishedAppsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListOAuthPublishedAppsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListPublishedAppIntegrationsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListPublishedAppIntegrationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPublishedAppIntegrationsRequest_SdkV2) {
}

func (to *ListPublishedAppIntegrationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPublishedAppIntegrationsRequest_SdkV2) {
}

func (m ListPublishedAppIntegrationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPublishedAppIntegrationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublishedAppIntegrationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPublishedAppIntegrationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPublishedAppIntegrationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListServicePrincipalFederationPoliciesRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId types.Int64 `tfsdk:"-"`
}

func (to *ListServicePrincipalFederationPoliciesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalFederationPoliciesRequest_SdkV2) {
}

func (to *ListServicePrincipalFederationPoliciesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalFederationPoliciesRequest_SdkV2) {
}

func (m ListServicePrincipalFederationPoliciesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListServicePrincipalFederationPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalFederationPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalFederationPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":            m.PageSize,
			"page_token":           m.PageToken,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalFederationPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":            types.Int64Type,
			"page_token":           types.StringType,
			"service_principal_id": types.Int64Type,
		},
	}
}

type ListServicePrincipalSecretsRequest_SdkV2 struct {
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

func (to *ListServicePrincipalSecretsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalSecretsRequest_SdkV2) {
}

func (to *ListServicePrincipalSecretsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalSecretsRequest_SdkV2) {
}

func (m ListServicePrincipalSecretsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListServicePrincipalSecretsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalSecretsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalSecretsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":            m.PageSize,
			"page_token":           m.PageToken,
			"service_principal_id": m.ServicePrincipalId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalSecretsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":            types.Int64Type,
			"page_token":           types.StringType,
			"service_principal_id": types.StringType,
		},
	}
}

type ListServicePrincipalSecretsResponse_SdkV2 struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken types.String `tfsdk:"next_page_token"`
	// List of the secrets
	Secrets types.List `tfsdk:"secrets"`
}

func (to *ListServicePrincipalSecretsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServicePrincipalSecretsResponse_SdkV2) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (to *ListServicePrincipalSecretsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServicePrincipalSecretsResponse_SdkV2) {
	if !from.Secrets.IsNull() && !from.Secrets.IsUnknown() && to.Secrets.IsNull() && len(from.Secrets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Secrets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Secrets = from.Secrets
	}
}

func (m ListServicePrincipalSecretsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListServicePrincipalSecretsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"secrets": reflect.TypeOf(SecretInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServicePrincipalSecretsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServicePrincipalSecretsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"secrets":         m.Secrets,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListServicePrincipalSecretsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"secrets": basetypes.ListType{
				ElemType: SecretInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSecrets returns the value of the Secrets field in ListServicePrincipalSecretsResponse_SdkV2 as
// a slice of SecretInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListServicePrincipalSecretsResponse_SdkV2) GetSecrets(ctx context.Context) ([]SecretInfo_SdkV2, bool) {
	if m.Secrets.IsNull() || m.Secrets.IsUnknown() {
		return nil, false
	}
	var v []SecretInfo_SdkV2
	d := m.Secrets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSecrets sets the value of the Secrets field in ListServicePrincipalSecretsResponse_SdkV2.
func (m *ListServicePrincipalSecretsResponse_SdkV2) SetSecrets(ctx context.Context, v []SecretInfo_SdkV2) {
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
type OidcFederationPolicy_SdkV2 struct {
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

func (to *OidcFederationPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OidcFederationPolicy_SdkV2) {
	if !from.Audiences.IsNull() && !from.Audiences.IsUnknown() && to.Audiences.IsNull() && len(from.Audiences.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Audiences, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Audiences = from.Audiences
	}
}

func (to *OidcFederationPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from OidcFederationPolicy_SdkV2) {
	if !from.Audiences.IsNull() && !from.Audiences.IsUnknown() && to.Audiences.IsNull() && len(from.Audiences.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Audiences, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Audiences = from.Audiences
	}
}

func (m OidcFederationPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m OidcFederationPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"audiences": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OidcFederationPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m OidcFederationPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m OidcFederationPolicy_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAudiences returns the value of the Audiences field in OidcFederationPolicy_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *OidcFederationPolicy_SdkV2) GetAudiences(ctx context.Context) ([]types.String, bool) {
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

// SetAudiences sets the value of the Audiences field in OidcFederationPolicy_SdkV2.
func (m *OidcFederationPolicy_SdkV2) SetAudiences(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["audiences"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Audiences = types.ListValueMust(t, vs)
}

type PublishedAppOutput_SdkV2 struct {
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

func (to *PublishedAppOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublishedAppOutput_SdkV2) {
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

func (to *PublishedAppOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PublishedAppOutput_SdkV2) {
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

func (m PublishedAppOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PublishedAppOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls": reflect.TypeOf(types.String{}),
		"scopes":        reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublishedAppOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m PublishedAppOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PublishedAppOutput_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetRedirectUrls returns the value of the RedirectUrls field in PublishedAppOutput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PublishedAppOutput_SdkV2) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
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

// SetRedirectUrls sets the value of the RedirectUrls field in PublishedAppOutput_SdkV2.
func (m *PublishedAppOutput_SdkV2) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in PublishedAppOutput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PublishedAppOutput_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in PublishedAppOutput_SdkV2.
func (m *PublishedAppOutput_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

type SecretInfo_SdkV2 struct {
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

func (to *SecretInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SecretInfo_SdkV2) {
}

func (to *SecretInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SecretInfo_SdkV2) {
}

func (m SecretInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SecretInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SecretInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m SecretInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SecretInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type TokenAccessPolicy_SdkV2 struct {
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

func (to *TokenAccessPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenAccessPolicy_SdkV2) {
}

func (to *TokenAccessPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenAccessPolicy_SdkV2) {
}

func (m TokenAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TokenAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"absolute_session_lifetime_in_minutes": types.Int64Type,
			"access_token_ttl_in_minutes":          types.Int64Type,
			"enable_single_use_refresh_tokens":     types.BoolType,
			"refresh_token_ttl_in_minutes":         types.Int64Type,
		},
	}
}

type UpdateAccountFederationPolicyRequest_SdkV2 struct {
	Policy types.List `tfsdk:"policy"`
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

func (to *UpdateAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAccountFederationPolicyRequest_SdkV2) {
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

func (to *UpdateAccountFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAccountFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m UpdateAccountFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateAccountFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAccountFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"policy":      m.Policy,
			"policy_id":   m.PolicyId,
			"update_mask": m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAccountFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
			"policy_id":   types.StringType,
			"update_mask": types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in UpdateAccountFederationPolicyRequest_SdkV2 as
// a FederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAccountFederationPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (FederationPolicy_SdkV2, bool) {
	var e FederationPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in UpdateAccountFederationPolicyRequest_SdkV2.
func (m *UpdateAccountFederationPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v FederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}

type UpdateCustomAppIntegration_SdkV2 struct {
	IntegrationId types.String `tfsdk:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls types.List `tfsdk:"redirect_urls"`
	// List of OAuth scopes to be updated in the custom OAuth app integration,
	// similar to redirect URIs this will fully replace the existing values
	// instead of appending
	Scopes types.List `tfsdk:"scopes"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy types.List `tfsdk:"token_access_policy"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes types.List `tfsdk:"user_authorized_scopes"`
}

func (to *UpdateCustomAppIntegration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCustomAppIntegration_SdkV2) {
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

func (to *UpdateCustomAppIntegration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCustomAppIntegration_SdkV2) {
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

func (m UpdateCustomAppIntegration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["redirect_urls"] = attrs["redirect_urls"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateCustomAppIntegration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"redirect_urls":          reflect.TypeOf(types.String{}),
		"scopes":                 reflect.TypeOf(types.String{}),
		"token_access_policy":    reflect.TypeOf(TokenAccessPolicy_SdkV2{}),
		"user_authorized_scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomAppIntegration_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCustomAppIntegration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateCustomAppIntegration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
			"redirect_urls": basetypes.ListType{
				ElemType: types.StringType,
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_access_policy": basetypes.ListType{
				ElemType: TokenAccessPolicy_SdkV2{}.Type(ctx),
			},
			"user_authorized_scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetRedirectUrls returns the value of the RedirectUrls field in UpdateCustomAppIntegration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration_SdkV2) GetRedirectUrls(ctx context.Context) ([]types.String, bool) {
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

// SetRedirectUrls sets the value of the RedirectUrls field in UpdateCustomAppIntegration_SdkV2.
func (m *UpdateCustomAppIntegration_SdkV2) SetRedirectUrls(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["redirect_urls"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RedirectUrls = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in UpdateCustomAppIntegration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in UpdateCustomAppIntegration_SdkV2.
func (m *UpdateCustomAppIntegration_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in UpdateCustomAppIntegration_SdkV2 as
// a TokenAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration_SdkV2) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy_SdkV2, bool) {
	var e TokenAccessPolicy_SdkV2
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy_SdkV2
	d := m.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdateCustomAppIntegration_SdkV2.
func (m *UpdateCustomAppIntegration_SdkV2) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	m.TokenAccessPolicy = types.ListValueMust(t, vs)
}

// GetUserAuthorizedScopes returns the value of the UserAuthorizedScopes field in UpdateCustomAppIntegration_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCustomAppIntegration_SdkV2) GetUserAuthorizedScopes(ctx context.Context) ([]types.String, bool) {
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

// SetUserAuthorizedScopes sets the value of the UserAuthorizedScopes field in UpdateCustomAppIntegration_SdkV2.
func (m *UpdateCustomAppIntegration_SdkV2) SetUserAuthorizedScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["user_authorized_scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UserAuthorizedScopes = types.ListValueMust(t, vs)
}

type UpdateCustomAppIntegrationOutput_SdkV2 struct {
}

func (to *UpdateCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCustomAppIntegrationOutput_SdkV2) {
}

func (to *UpdateCustomAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCustomAppIntegrationOutput_SdkV2) {
}

func (m UpdateCustomAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCustomAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCustomAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCustomAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCustomAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCustomAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdatePublishedAppIntegration_SdkV2 struct {
	IntegrationId types.String `tfsdk:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy types.List `tfsdk:"token_access_policy"`
}

func (to *UpdatePublishedAppIntegration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePublishedAppIntegration_SdkV2) {
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

func (to *UpdatePublishedAppIntegration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdatePublishedAppIntegration_SdkV2) {
	if !from.TokenAccessPolicy.IsNull() && !from.TokenAccessPolicy.IsUnknown() {
		if toTokenAccessPolicy, ok := to.GetTokenAccessPolicy(ctx); ok {
			if fromTokenAccessPolicy, ok := from.GetTokenAccessPolicy(ctx); ok {
				toTokenAccessPolicy.SyncFieldsDuringRead(ctx, fromTokenAccessPolicy)
				to.SetTokenAccessPolicy(ctx, toTokenAccessPolicy)
			}
		}
	}
}

func (m UpdatePublishedAppIntegration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_access_policy"] = attrs["token_access_policy"].SetOptional()
	attrs["token_access_policy"] = attrs["token_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdatePublishedAppIntegration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_access_policy": reflect.TypeOf(TokenAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePublishedAppIntegration_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdatePublishedAppIntegration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_id":      m.IntegrationId,
			"token_access_policy": m.TokenAccessPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePublishedAppIntegration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_id": types.StringType,
			"token_access_policy": basetypes.ListType{
				ElemType: TokenAccessPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTokenAccessPolicy returns the value of the TokenAccessPolicy field in UpdatePublishedAppIntegration_SdkV2 as
// a TokenAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePublishedAppIntegration_SdkV2) GetTokenAccessPolicy(ctx context.Context) (TokenAccessPolicy_SdkV2, bool) {
	var e TokenAccessPolicy_SdkV2
	if m.TokenAccessPolicy.IsNull() || m.TokenAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []TokenAccessPolicy_SdkV2
	d := m.TokenAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenAccessPolicy sets the value of the TokenAccessPolicy field in UpdatePublishedAppIntegration_SdkV2.
func (m *UpdatePublishedAppIntegration_SdkV2) SetTokenAccessPolicy(ctx context.Context, v TokenAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_access_policy"]
	m.TokenAccessPolicy = types.ListValueMust(t, vs)
}

type UpdatePublishedAppIntegrationOutput_SdkV2 struct {
}

func (to *UpdatePublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePublishedAppIntegrationOutput_SdkV2) {
}

func (to *UpdatePublishedAppIntegrationOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdatePublishedAppIntegrationOutput_SdkV2) {
}

func (m UpdatePublishedAppIntegrationOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePublishedAppIntegrationOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePublishedAppIntegrationOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePublishedAppIntegrationOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdatePublishedAppIntegrationOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePublishedAppIntegrationOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type UpdateServicePrincipalFederationPolicyRequest_SdkV2 struct {
	Policy types.List `tfsdk:"policy"`
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

func (to *UpdateServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateServicePrincipalFederationPolicyRequest_SdkV2) {
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

func (to *UpdateServicePrincipalFederationPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateServicePrincipalFederationPolicyRequest_SdkV2) {
	if !from.Policy.IsNull() && !from.Policy.IsUnknown() {
		if toPolicy, ok := to.GetPolicy(ctx); ok {
			if fromPolicy, ok := from.GetPolicy(ctx); ok {
				toPolicy.SyncFieldsDuringRead(ctx, fromPolicy)
				to.SetPolicy(ctx, toPolicy)
			}
		}
	}
}

func (m UpdateServicePrincipalFederationPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["policy"] = attrs["policy"].SetRequired()
	attrs["policy"] = attrs["policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateServicePrincipalFederationPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"policy": reflect.TypeOf(FederationPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateServicePrincipalFederationPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateServicePrincipalFederationPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateServicePrincipalFederationPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"policy": basetypes.ListType{
				ElemType: FederationPolicy_SdkV2{}.Type(ctx),
			},
			"policy_id":            types.StringType,
			"service_principal_id": types.Int64Type,
			"update_mask":          types.StringType,
		},
	}
}

// GetPolicy returns the value of the Policy field in UpdateServicePrincipalFederationPolicyRequest_SdkV2 as
// a FederationPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateServicePrincipalFederationPolicyRequest_SdkV2) GetPolicy(ctx context.Context) (FederationPolicy_SdkV2, bool) {
	var e FederationPolicy_SdkV2
	if m.Policy.IsNull() || m.Policy.IsUnknown() {
		return e, false
	}
	var v []FederationPolicy_SdkV2
	d := m.Policy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicy sets the value of the Policy field in UpdateServicePrincipalFederationPolicyRequest_SdkV2.
func (m *UpdateServicePrincipalFederationPolicyRequest_SdkV2) SetPolicy(ctx context.Context, v FederationPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy"]
	m.Policy = types.ListValueMust(t, vs)
}
