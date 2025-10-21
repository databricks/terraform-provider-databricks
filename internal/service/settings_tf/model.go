// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package settings_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccountIpAccessEnable struct {
	AcctIpAclEnable types.Object `tfsdk:"acct_ip_acl_enable"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *AccountIpAccessEnable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccountIpAccessEnable) {
	if !from.AcctIpAclEnable.IsNull() && !from.AcctIpAclEnable.IsUnknown() {
		if toAcctIpAclEnable, ok := to.GetAcctIpAclEnable(ctx); ok {
			if fromAcctIpAclEnable, ok := from.GetAcctIpAclEnable(ctx); ok {
				// Recursively sync the fields of AcctIpAclEnable
				toAcctIpAclEnable.SyncFieldsDuringCreateOrUpdate(ctx, fromAcctIpAclEnable)
				to.SetAcctIpAclEnable(ctx, toAcctIpAclEnable)
			}
		}
	}
}

func (to *AccountIpAccessEnable) SyncFieldsDuringRead(ctx context.Context, from AccountIpAccessEnable) {
	if !from.AcctIpAclEnable.IsNull() && !from.AcctIpAclEnable.IsUnknown() {
		if toAcctIpAclEnable, ok := to.GetAcctIpAclEnable(ctx); ok {
			if fromAcctIpAclEnable, ok := from.GetAcctIpAclEnable(ctx); ok {
				toAcctIpAclEnable.SyncFieldsDuringRead(ctx, fromAcctIpAclEnable)
				to.SetAcctIpAclEnable(ctx, toAcctIpAclEnable)
			}
		}
	}
}

func (m AccountIpAccessEnable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["acct_ip_acl_enable"] = attrs["acct_ip_acl_enable"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountIpAccessEnable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AccountIpAccessEnable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"acct_ip_acl_enable": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountIpAccessEnable
// only implements ToObjectValue() and Type().
func (m AccountIpAccessEnable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"acct_ip_acl_enable": m.AcctIpAclEnable,
			"etag":               m.Etag,
			"setting_name":       m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AccountIpAccessEnable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"acct_ip_acl_enable": BooleanMessage{}.Type(ctx),
			"etag":               types.StringType,
			"setting_name":       types.StringType,
		},
	}
}

// GetAcctIpAclEnable returns the value of the AcctIpAclEnable field in AccountIpAccessEnable as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *AccountIpAccessEnable) GetAcctIpAclEnable(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.AcctIpAclEnable.IsNull() || m.AcctIpAclEnable.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.AcctIpAclEnable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAcctIpAclEnable sets the value of the AcctIpAclEnable field in AccountIpAccessEnable.
func (m *AccountIpAccessEnable) SetAcctIpAclEnable(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.AcctIpAclEnable = vs
}

type AccountNetworkPolicy struct {
	// The associated account ID for this Network Policy object.
	AccountId types.String `tfsdk:"account_id"`
	// The network policies applying for egress traffic.
	Egress types.Object `tfsdk:"egress"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
}

func (to *AccountNetworkPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccountNetworkPolicy) {
	if !from.Egress.IsNull() && !from.Egress.IsUnknown() {
		if toEgress, ok := to.GetEgress(ctx); ok {
			if fromEgress, ok := from.GetEgress(ctx); ok {
				// Recursively sync the fields of Egress
				toEgress.SyncFieldsDuringCreateOrUpdate(ctx, fromEgress)
				to.SetEgress(ctx, toEgress)
			}
		}
	}
}

func (to *AccountNetworkPolicy) SyncFieldsDuringRead(ctx context.Context, from AccountNetworkPolicy) {
	if !from.Egress.IsNull() && !from.Egress.IsUnknown() {
		if toEgress, ok := to.GetEgress(ctx); ok {
			if fromEgress, ok := from.GetEgress(ctx); ok {
				toEgress.SyncFieldsDuringRead(ctx, fromEgress)
				to.SetEgress(ctx, toEgress)
			}
		}
	}
}

func (m AccountNetworkPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["egress"] = attrs["egress"].SetOptional()
	attrs["network_policy_id"] = attrs["network_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AccountNetworkPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AccountNetworkPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress": reflect.TypeOf(NetworkPolicyEgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountNetworkPolicy
// only implements ToObjectValue() and Type().
func (m AccountNetworkPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":        m.AccountId,
			"egress":            m.Egress,
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AccountNetworkPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":        types.StringType,
			"egress":            NetworkPolicyEgress{}.Type(ctx),
			"network_policy_id": types.StringType,
		},
	}
}

// GetEgress returns the value of the Egress field in AccountNetworkPolicy as
// a NetworkPolicyEgress value.
// If the field is unknown or null, the boolean return value is false.
func (m *AccountNetworkPolicy) GetEgress(ctx context.Context) (NetworkPolicyEgress, bool) {
	var e NetworkPolicyEgress
	if m.Egress.IsNull() || m.Egress.IsUnknown() {
		return e, false
	}
	var v NetworkPolicyEgress
	d := m.Egress.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEgress sets the value of the Egress field in AccountNetworkPolicy.
func (m *AccountNetworkPolicy) SetEgress(ctx context.Context, v NetworkPolicyEgress) {
	vs := v.ToObjectValue(ctx)
	m.Egress = vs
}

type AibiDashboardEmbeddingAccessPolicy struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type"`
}

func (to *AibiDashboardEmbeddingAccessPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingAccessPolicy) {
}

func (to *AibiDashboardEmbeddingAccessPolicy) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingAccessPolicy) {
}

func (m AibiDashboardEmbeddingAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_policy_type"] = attrs["access_policy_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingAccessPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AibiDashboardEmbeddingAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicy
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_policy_type": m.AccessPolicyType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingAccessPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_policy_type": types.StringType,
		},
	}
}

type AibiDashboardEmbeddingAccessPolicySetting struct {
	AibiDashboardEmbeddingAccessPolicy types.Object `tfsdk:"aibi_dashboard_embedding_access_policy"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *AibiDashboardEmbeddingAccessPolicySetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingAccessPolicySetting) {
	if !from.AibiDashboardEmbeddingAccessPolicy.IsNull() && !from.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toAibiDashboardEmbeddingAccessPolicy, ok := to.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromAibiDashboardEmbeddingAccessPolicy, ok := from.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				// Recursively sync the fields of AibiDashboardEmbeddingAccessPolicy
				toAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromAibiDashboardEmbeddingAccessPolicy)
				to.SetAibiDashboardEmbeddingAccessPolicy(ctx, toAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
}

func (to *AibiDashboardEmbeddingAccessPolicySetting) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingAccessPolicySetting) {
	if !from.AibiDashboardEmbeddingAccessPolicy.IsNull() && !from.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toAibiDashboardEmbeddingAccessPolicy, ok := to.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromAibiDashboardEmbeddingAccessPolicy, ok := from.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringRead(ctx, fromAibiDashboardEmbeddingAccessPolicy)
				to.SetAibiDashboardEmbeddingAccessPolicy(ctx, toAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
}

func (m AibiDashboardEmbeddingAccessPolicySetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingAccessPolicySetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AibiDashboardEmbeddingAccessPolicySetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicySetting
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingAccessPolicySetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_access_policy": m.AibiDashboardEmbeddingAccessPolicy,
			"etag":                                   m.Etag,
			"setting_name":                           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingAccessPolicySetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_access_policy": AibiDashboardEmbeddingAccessPolicy{}.Type(ctx),
			"etag":                                   types.StringType,
			"setting_name":                           types.StringType,
		},
	}
}

// GetAibiDashboardEmbeddingAccessPolicy returns the value of the AibiDashboardEmbeddingAccessPolicy field in AibiDashboardEmbeddingAccessPolicySetting as
// a AibiDashboardEmbeddingAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *AibiDashboardEmbeddingAccessPolicySetting) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy, bool) {
	var e AibiDashboardEmbeddingAccessPolicy
	if m.AibiDashboardEmbeddingAccessPolicy.IsNull() || m.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingAccessPolicy
	d := m.AibiDashboardEmbeddingAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in AibiDashboardEmbeddingAccessPolicySetting.
func (m *AibiDashboardEmbeddingAccessPolicySetting) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.AibiDashboardEmbeddingAccessPolicy = vs
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains types.List `tfsdk:"approved_domains"`
}

func (to *AibiDashboardEmbeddingApprovedDomains) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingApprovedDomains) {
	if !from.ApprovedDomains.IsNull() && !from.ApprovedDomains.IsUnknown() && to.ApprovedDomains.IsNull() && len(from.ApprovedDomains.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ApprovedDomains, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ApprovedDomains = from.ApprovedDomains
	}
}

func (to *AibiDashboardEmbeddingApprovedDomains) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingApprovedDomains) {
	if !from.ApprovedDomains.IsNull() && !from.ApprovedDomains.IsUnknown() && to.ApprovedDomains.IsNull() && len(from.ApprovedDomains.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ApprovedDomains, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ApprovedDomains = from.ApprovedDomains
	}
}

func (m AibiDashboardEmbeddingApprovedDomains) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["approved_domains"] = attrs["approved_domains"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingApprovedDomains.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AibiDashboardEmbeddingApprovedDomains) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomains
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingApprovedDomains) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"approved_domains": m.ApprovedDomains,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingApprovedDomains) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"approved_domains": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetApprovedDomains returns the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AibiDashboardEmbeddingApprovedDomains) GetApprovedDomains(ctx context.Context) ([]types.String, bool) {
	if m.ApprovedDomains.IsNull() || m.ApprovedDomains.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApprovedDomains sets the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains.
func (m *AibiDashboardEmbeddingApprovedDomains) SetApprovedDomains(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["approved_domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ApprovedDomains = types.ListValueMust(t, vs)
}

type AibiDashboardEmbeddingApprovedDomainsSetting struct {
	AibiDashboardEmbeddingApprovedDomains types.Object `tfsdk:"aibi_dashboard_embedding_approved_domains"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *AibiDashboardEmbeddingApprovedDomainsSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingApprovedDomainsSetting) {
	if !from.AibiDashboardEmbeddingApprovedDomains.IsNull() && !from.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toAibiDashboardEmbeddingApprovedDomains, ok := to.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromAibiDashboardEmbeddingApprovedDomains, ok := from.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				// Recursively sync the fields of AibiDashboardEmbeddingApprovedDomains
				toAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringCreateOrUpdate(ctx, fromAibiDashboardEmbeddingApprovedDomains)
				to.SetAibiDashboardEmbeddingApprovedDomains(ctx, toAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
}

func (to *AibiDashboardEmbeddingApprovedDomainsSetting) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingApprovedDomainsSetting) {
	if !from.AibiDashboardEmbeddingApprovedDomains.IsNull() && !from.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toAibiDashboardEmbeddingApprovedDomains, ok := to.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromAibiDashboardEmbeddingApprovedDomains, ok := from.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringRead(ctx, fromAibiDashboardEmbeddingApprovedDomains)
				to.SetAibiDashboardEmbeddingApprovedDomains(ctx, toAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
}

func (m AibiDashboardEmbeddingApprovedDomainsSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AibiDashboardEmbeddingApprovedDomainsSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AibiDashboardEmbeddingApprovedDomainsSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomainsSetting
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingApprovedDomainsSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_approved_domains": m.AibiDashboardEmbeddingApprovedDomains,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingApprovedDomainsSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_approved_domains": AibiDashboardEmbeddingApprovedDomains{}.Type(ctx),
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetAibiDashboardEmbeddingApprovedDomains returns the value of the AibiDashboardEmbeddingApprovedDomains field in AibiDashboardEmbeddingApprovedDomainsSetting as
// a AibiDashboardEmbeddingApprovedDomains value.
// If the field is unknown or null, the boolean return value is false.
func (m *AibiDashboardEmbeddingApprovedDomainsSetting) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains, bool) {
	var e AibiDashboardEmbeddingApprovedDomains
	if m.AibiDashboardEmbeddingApprovedDomains.IsNull() || m.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingApprovedDomains
	d := m.AibiDashboardEmbeddingApprovedDomains.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in AibiDashboardEmbeddingApprovedDomainsSetting.
func (m *AibiDashboardEmbeddingApprovedDomainsSetting) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains) {
	vs := v.ToObjectValue(ctx)
	m.AibiDashboardEmbeddingApprovedDomains = vs
}

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace types.Object `tfsdk:"automatic_cluster_update_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *AutomaticClusterUpdateSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutomaticClusterUpdateSetting) {
	if !from.AutomaticClusterUpdateWorkspace.IsNull() && !from.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toAutomaticClusterUpdateWorkspace, ok := to.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromAutomaticClusterUpdateWorkspace, ok := from.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				// Recursively sync the fields of AutomaticClusterUpdateWorkspace
				toAutomaticClusterUpdateWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromAutomaticClusterUpdateWorkspace)
				to.SetAutomaticClusterUpdateWorkspace(ctx, toAutomaticClusterUpdateWorkspace)
			}
		}
	}
}

func (to *AutomaticClusterUpdateSetting) SyncFieldsDuringRead(ctx context.Context, from AutomaticClusterUpdateSetting) {
	if !from.AutomaticClusterUpdateWorkspace.IsNull() && !from.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toAutomaticClusterUpdateWorkspace, ok := to.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromAutomaticClusterUpdateWorkspace, ok := from.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				toAutomaticClusterUpdateWorkspace.SyncFieldsDuringRead(ctx, fromAutomaticClusterUpdateWorkspace)
				to.SetAutomaticClusterUpdateWorkspace(ctx, toAutomaticClusterUpdateWorkspace)
			}
		}
	}
}

func (m AutomaticClusterUpdateSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutomaticClusterUpdateSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AutomaticClusterUpdateSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"automatic_cluster_update_workspace": reflect.TypeOf(ClusterAutoRestartMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutomaticClusterUpdateSetting
// only implements ToObjectValue() and Type().
func (m AutomaticClusterUpdateSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"automatic_cluster_update_workspace": m.AutomaticClusterUpdateWorkspace,
			"etag":                               m.Etag,
			"setting_name":                       m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutomaticClusterUpdateSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"automatic_cluster_update_workspace": ClusterAutoRestartMessage{}.Type(ctx),
			"etag":                               types.StringType,
			"setting_name":                       types.StringType,
		},
	}
}

// GetAutomaticClusterUpdateWorkspace returns the value of the AutomaticClusterUpdateWorkspace field in AutomaticClusterUpdateSetting as
// a ClusterAutoRestartMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutomaticClusterUpdateSetting) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage, bool) {
	var e ClusterAutoRestartMessage
	if m.AutomaticClusterUpdateWorkspace.IsNull() || m.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessage
	d := m.AutomaticClusterUpdateWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in AutomaticClusterUpdateSetting.
func (m *AutomaticClusterUpdateSetting) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage) {
	vs := v.ToObjectValue(ctx)
	m.AutomaticClusterUpdateWorkspace = vs
}

type BooleanMessage struct {
	Value types.Bool `tfsdk:"value"`
}

func (to *BooleanMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BooleanMessage) {
}

func (to *BooleanMessage) SyncFieldsDuringRead(ctx context.Context, from BooleanMessage) {
}

func (m BooleanMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BooleanMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BooleanMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage
// only implements ToObjectValue() and Type().
func (m BooleanMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BooleanMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type ClusterAutoRestartMessage struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`

	EnablementDetails types.Object `tfsdk:"enablement_details"`

	MaintenanceWindow types.Object `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

func (to *ClusterAutoRestartMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessage) {
	if !from.EnablementDetails.IsNull() && !from.EnablementDetails.IsUnknown() {
		if toEnablementDetails, ok := to.GetEnablementDetails(ctx); ok {
			if fromEnablementDetails, ok := from.GetEnablementDetails(ctx); ok {
				// Recursively sync the fields of EnablementDetails
				toEnablementDetails.SyncFieldsDuringCreateOrUpdate(ctx, fromEnablementDetails)
				to.SetEnablementDetails(ctx, toEnablementDetails)
			}
		}
	}
	if !from.MaintenanceWindow.IsNull() && !from.MaintenanceWindow.IsUnknown() {
		if toMaintenanceWindow, ok := to.GetMaintenanceWindow(ctx); ok {
			if fromMaintenanceWindow, ok := from.GetMaintenanceWindow(ctx); ok {
				// Recursively sync the fields of MaintenanceWindow
				toMaintenanceWindow.SyncFieldsDuringCreateOrUpdate(ctx, fromMaintenanceWindow)
				to.SetMaintenanceWindow(ctx, toMaintenanceWindow)
			}
		}
	}
}

func (to *ClusterAutoRestartMessage) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessage) {
	if !from.EnablementDetails.IsNull() && !from.EnablementDetails.IsUnknown() {
		if toEnablementDetails, ok := to.GetEnablementDetails(ctx); ok {
			if fromEnablementDetails, ok := from.GetEnablementDetails(ctx); ok {
				toEnablementDetails.SyncFieldsDuringRead(ctx, fromEnablementDetails)
				to.SetEnablementDetails(ctx, toEnablementDetails)
			}
		}
	}
	if !from.MaintenanceWindow.IsNull() && !from.MaintenanceWindow.IsUnknown() {
		if toMaintenanceWindow, ok := to.GetMaintenanceWindow(ctx); ok {
			if fromMaintenanceWindow, ok := from.GetMaintenanceWindow(ctx); ok {
				toMaintenanceWindow.SyncFieldsDuringRead(ctx, fromMaintenanceWindow)
				to.SetMaintenanceWindow(ctx, toMaintenanceWindow)
			}
		}
	}
}

func (m ClusterAutoRestartMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_toggle"] = attrs["can_toggle"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["enablement_details"] = attrs["enablement_details"].SetOptional()
	attrs["maintenance_window"] = attrs["maintenance_window"].SetOptional()
	attrs["restart_even_if_no_updates_available"] = attrs["restart_even_if_no_updates_available"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessage
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_toggle":                           m.CanToggle,
			"enabled":                              m.Enabled,
			"enablement_details":                   m.EnablementDetails,
			"maintenance_window":                   m.MaintenanceWindow,
			"restart_even_if_no_updates_available": m.RestartEvenIfNoUpdatesAvailable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_toggle":                           types.BoolType,
			"enabled":                              types.BoolType,
			"enablement_details":                   ClusterAutoRestartMessageEnablementDetails{}.Type(ctx),
			"maintenance_window":                   ClusterAutoRestartMessageMaintenanceWindow{}.Type(ctx),
			"restart_even_if_no_updates_available": types.BoolType,
		},
	}
}

// GetEnablementDetails returns the value of the EnablementDetails field in ClusterAutoRestartMessage as
// a ClusterAutoRestartMessageEnablementDetails value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessage) GetEnablementDetails(ctx context.Context) (ClusterAutoRestartMessageEnablementDetails, bool) {
	var e ClusterAutoRestartMessageEnablementDetails
	if m.EnablementDetails.IsNull() || m.EnablementDetails.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageEnablementDetails
	d := m.EnablementDetails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnablementDetails sets the value of the EnablementDetails field in ClusterAutoRestartMessage.
func (m *ClusterAutoRestartMessage) SetEnablementDetails(ctx context.Context, v ClusterAutoRestartMessageEnablementDetails) {
	vs := v.ToObjectValue(ctx)
	m.EnablementDetails = vs
}

// GetMaintenanceWindow returns the value of the MaintenanceWindow field in ClusterAutoRestartMessage as
// a ClusterAutoRestartMessageMaintenanceWindow value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessage) GetMaintenanceWindow(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindow, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindow
	if m.MaintenanceWindow.IsNull() || m.MaintenanceWindow.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageMaintenanceWindow
	d := m.MaintenanceWindow.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMaintenanceWindow sets the value of the MaintenanceWindow field in ClusterAutoRestartMessage.
func (m *ClusterAutoRestartMessage) SetMaintenanceWindow(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindow) {
	vs := v.ToObjectValue(ctx)
	m.MaintenanceWindow = vs
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode types.Bool `tfsdk:"forced_for_compliance_mode"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement types.Bool `tfsdk:"unavailable_for_disabled_entitlement"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier types.Bool `tfsdk:"unavailable_for_non_enterprise_tier"`
}

func (to *ClusterAutoRestartMessageEnablementDetails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageEnablementDetails) {
}

func (to *ClusterAutoRestartMessageEnablementDetails) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageEnablementDetails) {
}

func (m ClusterAutoRestartMessageEnablementDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["forced_for_compliance_mode"] = attrs["forced_for_compliance_mode"].SetOptional()
	attrs["unavailable_for_disabled_entitlement"] = attrs["unavailable_for_disabled_entitlement"].SetOptional()
	attrs["unavailable_for_non_enterprise_tier"] = attrs["unavailable_for_non_enterprise_tier"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageEnablementDetails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessageEnablementDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageEnablementDetails
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageEnablementDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"forced_for_compliance_mode":           m.ForcedForComplianceMode,
			"unavailable_for_disabled_entitlement": m.UnavailableForDisabledEntitlement,
			"unavailable_for_non_enterprise_tier":  m.UnavailableForNonEnterpriseTier,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageEnablementDetails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"forced_for_compliance_mode":           types.BoolType,
			"unavailable_for_disabled_entitlement": types.BoolType,
			"unavailable_for_non_enterprise_tier":  types.BoolType,
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule types.Object `tfsdk:"week_day_based_schedule"`
}

func (to *ClusterAutoRestartMessageMaintenanceWindow) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindow) {
	if !from.WeekDayBasedSchedule.IsNull() && !from.WeekDayBasedSchedule.IsUnknown() {
		if toWeekDayBasedSchedule, ok := to.GetWeekDayBasedSchedule(ctx); ok {
			if fromWeekDayBasedSchedule, ok := from.GetWeekDayBasedSchedule(ctx); ok {
				// Recursively sync the fields of WeekDayBasedSchedule
				toWeekDayBasedSchedule.SyncFieldsDuringCreateOrUpdate(ctx, fromWeekDayBasedSchedule)
				to.SetWeekDayBasedSchedule(ctx, toWeekDayBasedSchedule)
			}
		}
	}
}

func (to *ClusterAutoRestartMessageMaintenanceWindow) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindow) {
	if !from.WeekDayBasedSchedule.IsNull() && !from.WeekDayBasedSchedule.IsUnknown() {
		if toWeekDayBasedSchedule, ok := to.GetWeekDayBasedSchedule(ctx); ok {
			if fromWeekDayBasedSchedule, ok := from.GetWeekDayBasedSchedule(ctx); ok {
				toWeekDayBasedSchedule.SyncFieldsDuringRead(ctx, fromWeekDayBasedSchedule)
				to.SetWeekDayBasedSchedule(ctx, toWeekDayBasedSchedule)
			}
		}
	}
}

func (m ClusterAutoRestartMessageMaintenanceWindow) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["week_day_based_schedule"] = attrs["week_day_based_schedule"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessageMaintenanceWindow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"week_day_based_schedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindow
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageMaintenanceWindow) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"week_day_based_schedule": m.WeekDayBasedSchedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageMaintenanceWindow) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"week_day_based_schedule": ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}.Type(ctx),
		},
	}
}

// GetWeekDayBasedSchedule returns the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow as
// a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessageMaintenanceWindow) GetWeekDayBasedSchedule(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
	if m.WeekDayBasedSchedule.IsNull() || m.WeekDayBasedSchedule.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
	d := m.WeekDayBasedSchedule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWeekDayBasedSchedule sets the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow.
func (m *ClusterAutoRestartMessageMaintenanceWindow) SetWeekDayBasedSchedule(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	vs := v.ToObjectValue(ctx)
	m.WeekDayBasedSchedule = vs
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`

	Frequency types.String `tfsdk:"frequency"`

	WindowStartTime types.Object `tfsdk:"window_start_time"`
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	if !from.WindowStartTime.IsNull() && !from.WindowStartTime.IsUnknown() {
		if toWindowStartTime, ok := to.GetWindowStartTime(ctx); ok {
			if fromWindowStartTime, ok := from.GetWindowStartTime(ctx); ok {
				// Recursively sync the fields of WindowStartTime
				toWindowStartTime.SyncFieldsDuringCreateOrUpdate(ctx, fromWindowStartTime)
				to.SetWindowStartTime(ctx, toWindowStartTime)
			}
		}
	}
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	if !from.WindowStartTime.IsNull() && !from.WindowStartTime.IsUnknown() {
		if toWindowStartTime, ok := to.GetWindowStartTime(ctx); ok {
			if fromWindowStartTime, ok := from.GetWindowStartTime(ctx); ok {
				toWindowStartTime.SyncFieldsDuringRead(ctx, fromWindowStartTime)
				to.SetWindowStartTime(ctx, toWindowStartTime)
			}
		}
	}
}

func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["day_of_week"] = attrs["day_of_week"].SetOptional()
	attrs["frequency"] = attrs["frequency"].SetOptional()
	attrs["window_start_time"] = attrs["window_start_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"window_start_time": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"day_of_week":       m.DayOfWeek,
			"frequency":         m.Frequency,
			"window_start_time": m.WindowStartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"day_of_week":       types.StringType,
			"frequency":         types.StringType,
			"window_start_time": ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}.Type(ctx),
		},
	}
}

// GetWindowStartTime returns the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule as
// a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetWindowStartTime(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
	if m.WindowStartTime.IsNull() || m.WindowStartTime.IsUnknown() {
		return e, false
	}
	var v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
	d := m.WindowStartTime.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWindowStartTime sets the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
func (m *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SetWindowStartTime(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
	vs := v.ToObjectValue(ctx)
	m.WindowStartTime = vs
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["hours"] = attrs["hours"].SetOptional()
	attrs["minutes"] = attrs["minutes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWindowStartTime.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hours":   m.Hours,
			"minutes": m.Minutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"hours":   types.Int64Type,
			"minutes": types.Int64Type,
		},
	}
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards types.List `tfsdk:"compliance_standards"`

	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (to *ComplianceSecurityProfile) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComplianceSecurityProfile) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (to *ComplianceSecurityProfile) SyncFieldsDuringRead(ctx context.Context, from ComplianceSecurityProfile) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (m ComplianceSecurityProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compliance_standards"] = attrs["compliance_standards"].SetOptional()
	attrs["is_enabled"] = attrs["is_enabled"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplianceSecurityProfile.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ComplianceSecurityProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile
// only implements ToObjectValue() and Type().
func (m ComplianceSecurityProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": m.ComplianceStandards,
			"is_enabled":           m.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComplianceSecurityProfile) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enabled": types.BoolType,
		},
	}
}

// GetComplianceStandards returns the value of the ComplianceStandards field in ComplianceSecurityProfile as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ComplianceSecurityProfile) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
	if m.ComplianceStandards.IsNull() || m.ComplianceStandards.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ComplianceStandards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile.
func (m *ComplianceSecurityProfile) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ComplianceStandards = types.ListValueMust(t, vs)
}

type ComplianceSecurityProfileSetting struct {
	ComplianceSecurityProfileWorkspace types.Object `tfsdk:"compliance_security_profile_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *ComplianceSecurityProfileSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComplianceSecurityProfileSetting) {
	if !from.ComplianceSecurityProfileWorkspace.IsNull() && !from.ComplianceSecurityProfileWorkspace.IsUnknown() {
		if toComplianceSecurityProfileWorkspace, ok := to.GetComplianceSecurityProfileWorkspace(ctx); ok {
			if fromComplianceSecurityProfileWorkspace, ok := from.GetComplianceSecurityProfileWorkspace(ctx); ok {
				// Recursively sync the fields of ComplianceSecurityProfileWorkspace
				toComplianceSecurityProfileWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromComplianceSecurityProfileWorkspace)
				to.SetComplianceSecurityProfileWorkspace(ctx, toComplianceSecurityProfileWorkspace)
			}
		}
	}
}

func (to *ComplianceSecurityProfileSetting) SyncFieldsDuringRead(ctx context.Context, from ComplianceSecurityProfileSetting) {
	if !from.ComplianceSecurityProfileWorkspace.IsNull() && !from.ComplianceSecurityProfileWorkspace.IsUnknown() {
		if toComplianceSecurityProfileWorkspace, ok := to.GetComplianceSecurityProfileWorkspace(ctx); ok {
			if fromComplianceSecurityProfileWorkspace, ok := from.GetComplianceSecurityProfileWorkspace(ctx); ok {
				toComplianceSecurityProfileWorkspace.SyncFieldsDuringRead(ctx, fromComplianceSecurityProfileWorkspace)
				to.SetComplianceSecurityProfileWorkspace(ctx, toComplianceSecurityProfileWorkspace)
			}
		}
	}
}

func (m ComplianceSecurityProfileSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compliance_security_profile_workspace"] = attrs["compliance_security_profile_workspace"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ComplianceSecurityProfileSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ComplianceSecurityProfileSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_security_profile_workspace": reflect.TypeOf(ComplianceSecurityProfile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfileSetting
// only implements ToObjectValue() and Type().
func (m ComplianceSecurityProfileSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_security_profile_workspace": m.ComplianceSecurityProfileWorkspace,
			"etag":                                  m.Etag,
			"setting_name":                          m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComplianceSecurityProfileSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_security_profile_workspace": ComplianceSecurityProfile{}.Type(ctx),
			"etag":                                  types.StringType,
			"setting_name":                          types.StringType,
		},
	}
}

// GetComplianceSecurityProfileWorkspace returns the value of the ComplianceSecurityProfileWorkspace field in ComplianceSecurityProfileSetting as
// a ComplianceSecurityProfile value.
// If the field is unknown or null, the boolean return value is false.
func (m *ComplianceSecurityProfileSetting) GetComplianceSecurityProfileWorkspace(ctx context.Context) (ComplianceSecurityProfile, bool) {
	var e ComplianceSecurityProfile
	if m.ComplianceSecurityProfileWorkspace.IsNull() || m.ComplianceSecurityProfileWorkspace.IsUnknown() {
		return e, false
	}
	var v ComplianceSecurityProfile
	d := m.ComplianceSecurityProfileWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceSecurityProfileWorkspace sets the value of the ComplianceSecurityProfileWorkspace field in ComplianceSecurityProfileSetting.
func (m *ComplianceSecurityProfileSetting) SetComplianceSecurityProfileWorkspace(ctx context.Context, v ComplianceSecurityProfile) {
	vs := v.ToObjectValue(ctx)
	m.ComplianceSecurityProfileWorkspace = vs
}

type Config struct {
	Email types.Object `tfsdk:"email"`

	GenericWebhook types.Object `tfsdk:"generic_webhook"`

	MicrosoftTeams types.Object `tfsdk:"microsoft_teams"`

	Pagerduty types.Object `tfsdk:"pagerduty"`

	Slack types.Object `tfsdk:"slack"`
}

func (to *Config) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Config) {
	if !from.Email.IsNull() && !from.Email.IsUnknown() {
		if toEmail, ok := to.GetEmail(ctx); ok {
			if fromEmail, ok := from.GetEmail(ctx); ok {
				// Recursively sync the fields of Email
				toEmail.SyncFieldsDuringCreateOrUpdate(ctx, fromEmail)
				to.SetEmail(ctx, toEmail)
			}
		}
	}
	if !from.GenericWebhook.IsNull() && !from.GenericWebhook.IsUnknown() {
		if toGenericWebhook, ok := to.GetGenericWebhook(ctx); ok {
			if fromGenericWebhook, ok := from.GetGenericWebhook(ctx); ok {
				// Recursively sync the fields of GenericWebhook
				toGenericWebhook.SyncFieldsDuringCreateOrUpdate(ctx, fromGenericWebhook)
				to.SetGenericWebhook(ctx, toGenericWebhook)
			}
		}
	}
	if !from.MicrosoftTeams.IsNull() && !from.MicrosoftTeams.IsUnknown() {
		if toMicrosoftTeams, ok := to.GetMicrosoftTeams(ctx); ok {
			if fromMicrosoftTeams, ok := from.GetMicrosoftTeams(ctx); ok {
				// Recursively sync the fields of MicrosoftTeams
				toMicrosoftTeams.SyncFieldsDuringCreateOrUpdate(ctx, fromMicrosoftTeams)
				to.SetMicrosoftTeams(ctx, toMicrosoftTeams)
			}
		}
	}
	if !from.Pagerduty.IsNull() && !from.Pagerduty.IsUnknown() {
		if toPagerduty, ok := to.GetPagerduty(ctx); ok {
			if fromPagerduty, ok := from.GetPagerduty(ctx); ok {
				// Recursively sync the fields of Pagerduty
				toPagerduty.SyncFieldsDuringCreateOrUpdate(ctx, fromPagerduty)
				to.SetPagerduty(ctx, toPagerduty)
			}
		}
	}
	if !from.Slack.IsNull() && !from.Slack.IsUnknown() {
		if toSlack, ok := to.GetSlack(ctx); ok {
			if fromSlack, ok := from.GetSlack(ctx); ok {
				// Recursively sync the fields of Slack
				toSlack.SyncFieldsDuringCreateOrUpdate(ctx, fromSlack)
				to.SetSlack(ctx, toSlack)
			}
		}
	}
}

func (to *Config) SyncFieldsDuringRead(ctx context.Context, from Config) {
	if !from.Email.IsNull() && !from.Email.IsUnknown() {
		if toEmail, ok := to.GetEmail(ctx); ok {
			if fromEmail, ok := from.GetEmail(ctx); ok {
				toEmail.SyncFieldsDuringRead(ctx, fromEmail)
				to.SetEmail(ctx, toEmail)
			}
		}
	}
	if !from.GenericWebhook.IsNull() && !from.GenericWebhook.IsUnknown() {
		if toGenericWebhook, ok := to.GetGenericWebhook(ctx); ok {
			if fromGenericWebhook, ok := from.GetGenericWebhook(ctx); ok {
				toGenericWebhook.SyncFieldsDuringRead(ctx, fromGenericWebhook)
				to.SetGenericWebhook(ctx, toGenericWebhook)
			}
		}
	}
	if !from.MicrosoftTeams.IsNull() && !from.MicrosoftTeams.IsUnknown() {
		if toMicrosoftTeams, ok := to.GetMicrosoftTeams(ctx); ok {
			if fromMicrosoftTeams, ok := from.GetMicrosoftTeams(ctx); ok {
				toMicrosoftTeams.SyncFieldsDuringRead(ctx, fromMicrosoftTeams)
				to.SetMicrosoftTeams(ctx, toMicrosoftTeams)
			}
		}
	}
	if !from.Pagerduty.IsNull() && !from.Pagerduty.IsUnknown() {
		if toPagerduty, ok := to.GetPagerduty(ctx); ok {
			if fromPagerduty, ok := from.GetPagerduty(ctx); ok {
				toPagerduty.SyncFieldsDuringRead(ctx, fromPagerduty)
				to.SetPagerduty(ctx, toPagerduty)
			}
		}
	}
	if !from.Slack.IsNull() && !from.Slack.IsUnknown() {
		if toSlack, ok := to.GetSlack(ctx); ok {
			if fromSlack, ok := from.GetSlack(ctx); ok {
				toSlack.SyncFieldsDuringRead(ctx, fromSlack)
				to.SetSlack(ctx, toSlack)
			}
		}
	}
}

func (m Config) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email"] = attrs["email"].SetOptional()
	attrs["generic_webhook"] = attrs["generic_webhook"].SetOptional()
	attrs["microsoft_teams"] = attrs["microsoft_teams"].SetOptional()
	attrs["pagerduty"] = attrs["pagerduty"].SetOptional()
	attrs["slack"] = attrs["slack"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Config.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Config) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email":           reflect.TypeOf(EmailConfig{}),
		"generic_webhook": reflect.TypeOf(GenericWebhookConfig{}),
		"microsoft_teams": reflect.TypeOf(MicrosoftTeamsConfig{}),
		"pagerduty":       reflect.TypeOf(PagerdutyConfig{}),
		"slack":           reflect.TypeOf(SlackConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Config
// only implements ToObjectValue() and Type().
func (m Config) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email":           m.Email,
			"generic_webhook": m.GenericWebhook,
			"microsoft_teams": m.MicrosoftTeams,
			"pagerduty":       m.Pagerduty,
			"slack":           m.Slack,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Config) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email":           EmailConfig{}.Type(ctx),
			"generic_webhook": GenericWebhookConfig{}.Type(ctx),
			"microsoft_teams": MicrosoftTeamsConfig{}.Type(ctx),
			"pagerduty":       PagerdutyConfig{}.Type(ctx),
			"slack":           SlackConfig{}.Type(ctx),
		},
	}
}

// GetEmail returns the value of the Email field in Config as
// a EmailConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config) GetEmail(ctx context.Context) (EmailConfig, bool) {
	var e EmailConfig
	if m.Email.IsNull() || m.Email.IsUnknown() {
		return e, false
	}
	var v EmailConfig
	d := m.Email.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmail sets the value of the Email field in Config.
func (m *Config) SetEmail(ctx context.Context, v EmailConfig) {
	vs := v.ToObjectValue(ctx)
	m.Email = vs
}

// GetGenericWebhook returns the value of the GenericWebhook field in Config as
// a GenericWebhookConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config) GetGenericWebhook(ctx context.Context) (GenericWebhookConfig, bool) {
	var e GenericWebhookConfig
	if m.GenericWebhook.IsNull() || m.GenericWebhook.IsUnknown() {
		return e, false
	}
	var v GenericWebhookConfig
	d := m.GenericWebhook.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGenericWebhook sets the value of the GenericWebhook field in Config.
func (m *Config) SetGenericWebhook(ctx context.Context, v GenericWebhookConfig) {
	vs := v.ToObjectValue(ctx)
	m.GenericWebhook = vs
}

// GetMicrosoftTeams returns the value of the MicrosoftTeams field in Config as
// a MicrosoftTeamsConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config) GetMicrosoftTeams(ctx context.Context) (MicrosoftTeamsConfig, bool) {
	var e MicrosoftTeamsConfig
	if m.MicrosoftTeams.IsNull() || m.MicrosoftTeams.IsUnknown() {
		return e, false
	}
	var v MicrosoftTeamsConfig
	d := m.MicrosoftTeams.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMicrosoftTeams sets the value of the MicrosoftTeams field in Config.
func (m *Config) SetMicrosoftTeams(ctx context.Context, v MicrosoftTeamsConfig) {
	vs := v.ToObjectValue(ctx)
	m.MicrosoftTeams = vs
}

// GetPagerduty returns the value of the Pagerduty field in Config as
// a PagerdutyConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config) GetPagerduty(ctx context.Context) (PagerdutyConfig, bool) {
	var e PagerdutyConfig
	if m.Pagerduty.IsNull() || m.Pagerduty.IsUnknown() {
		return e, false
	}
	var v PagerdutyConfig
	d := m.Pagerduty.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPagerduty sets the value of the Pagerduty field in Config.
func (m *Config) SetPagerduty(ctx context.Context, v PagerdutyConfig) {
	vs := v.ToObjectValue(ctx)
	m.Pagerduty = vs
}

// GetSlack returns the value of the Slack field in Config as
// a SlackConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config) GetSlack(ctx context.Context) (SlackConfig, bool) {
	var e SlackConfig
	if m.Slack.IsNull() || m.Slack.IsUnknown() {
		return e, false
	}
	var v SlackConfig
	d := m.Slack.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSlack sets the value of the Slack field in Config.
func (m *Config) SetSlack(ctx context.Context, v SlackConfig) {
	vs := v.ToObjectValue(ctx)
	m.Slack = vs
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`

	ListType types.String `tfsdk:"list_type"`
}

func (to *CreateIpAccessList) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateIpAccessList) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *CreateIpAccessList) SyncFieldsDuringRead(ctx context.Context, from CreateIpAccessList) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m CreateIpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetRequired()
	attrs["list_type"] = attrs["list_type"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessList
// only implements ToObjectValue() and Type().
func (m CreateIpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_addresses": m.IpAddresses,
			"label":        m.Label,
			"list_type":    m.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateIpAccessList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":     types.StringType,
			"list_type": types.StringType,
		},
	}
}

// GetIpAddresses returns the value of the IpAddresses field in CreateIpAccessList as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateIpAccessList) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if m.IpAddresses.IsNull() || m.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in CreateIpAccessList.
func (m *CreateIpAccessList) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	IpAccessList types.Object `tfsdk:"ip_access_list"`
}

func (to *CreateIpAccessListResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateIpAccessListResponse) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				// Recursively sync the fields of IpAccessList
				toIpAccessList.SyncFieldsDuringCreateOrUpdate(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (to *CreateIpAccessListResponse) SyncFieldsDuringRead(ctx context.Context, from CreateIpAccessListResponse) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (m CreateIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessListResponse
// only implements ToObjectValue() and Type().
func (m CreateIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": m.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": IpAccessListInfo{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in CreateIpAccessListResponse as
// a IpAccessListInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateIpAccessListResponse) GetIpAccessList(ctx context.Context) (IpAccessListInfo, bool) {
	var e IpAccessListInfo
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v IpAccessListInfo
	d := m.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessList sets the value of the IpAccessList field in CreateIpAccessListResponse.
func (m *CreateIpAccessListResponse) SetIpAccessList(ctx context.Context, v IpAccessListInfo) {
	vs := v.ToObjectValue(ctx)
	m.IpAccessList = vs
}

type CreateNetworkConnectivityConfigRequest struct {
	NetworkConnectivityConfig types.Object `tfsdk:"network_connectivity_config"`
}

func (to *CreateNetworkConnectivityConfigRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkConnectivityConfigRequest) {
	if !from.NetworkConnectivityConfig.IsNull() && !from.NetworkConnectivityConfig.IsUnknown() {
		if toNetworkConnectivityConfig, ok := to.GetNetworkConnectivityConfig(ctx); ok {
			if fromNetworkConnectivityConfig, ok := from.GetNetworkConnectivityConfig(ctx); ok {
				// Recursively sync the fields of NetworkConnectivityConfig
				toNetworkConnectivityConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromNetworkConnectivityConfig)
				to.SetNetworkConnectivityConfig(ctx, toNetworkConnectivityConfig)
			}
		}
	}
}

func (to *CreateNetworkConnectivityConfigRequest) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkConnectivityConfigRequest) {
	if !from.NetworkConnectivityConfig.IsNull() && !from.NetworkConnectivityConfig.IsUnknown() {
		if toNetworkConnectivityConfig, ok := to.GetNetworkConnectivityConfig(ctx); ok {
			if fromNetworkConnectivityConfig, ok := from.GetNetworkConnectivityConfig(ctx); ok {
				toNetworkConnectivityConfig.SyncFieldsDuringRead(ctx, fromNetworkConnectivityConfig)
				to.SetNetworkConnectivityConfig(ctx, toNetworkConnectivityConfig)
			}
		}
	}
}

func (m CreateNetworkConnectivityConfigRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_connectivity_config"] = attrs["network_connectivity_config"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkConnectivityConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNetworkConnectivityConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_connectivity_config": reflect.TypeOf(CreateNetworkConnectivityConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfigRequest
// only implements ToObjectValue() and Type().
func (m CreateNetworkConnectivityConfigRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config": m.NetworkConnectivityConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkConnectivityConfigRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config": CreateNetworkConnectivityConfiguration{}.Type(ctx),
		},
	}
}

// GetNetworkConnectivityConfig returns the value of the NetworkConnectivityConfig field in CreateNetworkConnectivityConfigRequest as
// a CreateNetworkConnectivityConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkConnectivityConfigRequest) GetNetworkConnectivityConfig(ctx context.Context) (CreateNetworkConnectivityConfiguration, bool) {
	var e CreateNetworkConnectivityConfiguration
	if m.NetworkConnectivityConfig.IsNull() || m.NetworkConnectivityConfig.IsUnknown() {
		return e, false
	}
	var v CreateNetworkConnectivityConfiguration
	d := m.NetworkConnectivityConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNetworkConnectivityConfig sets the value of the NetworkConnectivityConfig field in CreateNetworkConnectivityConfigRequest.
func (m *CreateNetworkConnectivityConfigRequest) SetNetworkConnectivityConfig(ctx context.Context, v CreateNetworkConnectivityConfiguration) {
	vs := v.ToObjectValue(ctx)
	m.NetworkConnectivityConfig = vs
}

// Properties of the new network connectivity configuration.
type CreateNetworkConnectivityConfiguration struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	Name types.String `tfsdk:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region"`
}

func (to *CreateNetworkConnectivityConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkConnectivityConfiguration) {
}

func (to *CreateNetworkConnectivityConfiguration) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkConnectivityConfiguration) {
}

func (m CreateNetworkConnectivityConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["region"] = attrs["region"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkConnectivityConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNetworkConnectivityConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfiguration
// only implements ToObjectValue() and Type().
func (m CreateNetworkConnectivityConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   m.Name,
			"region": m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkConnectivityConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"region": types.StringType,
		},
	}
}

type CreateNetworkPolicyRequest struct {
	// Network policy configuration details.
	NetworkPolicy types.Object `tfsdk:"network_policy"`
}

func (to *CreateNetworkPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkPolicyRequest) {
	if !from.NetworkPolicy.IsNull() && !from.NetworkPolicy.IsUnknown() {
		if toNetworkPolicy, ok := to.GetNetworkPolicy(ctx); ok {
			if fromNetworkPolicy, ok := from.GetNetworkPolicy(ctx); ok {
				// Recursively sync the fields of NetworkPolicy
				toNetworkPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromNetworkPolicy)
				to.SetNetworkPolicy(ctx, toNetworkPolicy)
			}
		}
	}
}

func (to *CreateNetworkPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkPolicyRequest) {
	if !from.NetworkPolicy.IsNull() && !from.NetworkPolicy.IsUnknown() {
		if toNetworkPolicy, ok := to.GetNetworkPolicy(ctx); ok {
			if fromNetworkPolicy, ok := from.GetNetworkPolicy(ctx); ok {
				toNetworkPolicy.SyncFieldsDuringRead(ctx, fromNetworkPolicy)
				to.SetNetworkPolicy(ctx, toNetworkPolicy)
			}
		}
	}
}

func (m CreateNetworkPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_policy"] = attrs["network_policy"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_policy": reflect.TypeOf(AccountNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (m CreateNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy": m.NetworkPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy": AccountNetworkPolicy{}.Type(ctx),
		},
	}
}

// GetNetworkPolicy returns the value of the NetworkPolicy field in CreateNetworkPolicyRequest as
// a AccountNetworkPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkPolicyRequest) GetNetworkPolicy(ctx context.Context) (AccountNetworkPolicy, bool) {
	var e AccountNetworkPolicy
	if m.NetworkPolicy.IsNull() || m.NetworkPolicy.IsUnknown() {
		return e, false
	}
	var v AccountNetworkPolicy
	d := m.NetworkPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNetworkPolicy sets the value of the NetworkPolicy field in CreateNetworkPolicyRequest.
func (m *CreateNetworkPolicyRequest) SetNetworkPolicy(ctx context.Context, v AccountNetworkPolicy) {
	vs := v.ToObjectValue(ctx)
	m.NetworkPolicy = vs
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.Object `tfsdk:"config"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
}

func (to *CreateNotificationDestinationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNotificationDestinationRequest) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (to *CreateNotificationDestinationRequest) SyncFieldsDuringRead(ctx context.Context, from CreateNotificationDestinationRequest) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m CreateNotificationDestinationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (m CreateNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       m.Config,
			"display_name": m.DisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config":       Config{}.Type(ctx),
			"display_name": types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in CreateNotificationDestinationRequest as
// a Config value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNotificationDestinationRequest) GetConfig(ctx context.Context) (Config, bool) {
	var e Config
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v Config
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in CreateNotificationDestinationRequest.
func (m *CreateNotificationDestinationRequest) SetConfig(ctx context.Context, v Config) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId types.String `tfsdk:"application_id"`
	// Comment that describes the purpose of the token.
	Comment types.String `tfsdk:"comment"`
	// The number of seconds before the token expires.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

func (to *CreateOboTokenRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOboTokenRequest) {
}

func (to *CreateOboTokenRequest) SyncFieldsDuringRead(ctx context.Context, from CreateOboTokenRequest) {
}

func (m CreateOboTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["application_id"] = attrs["application_id"].SetRequired()
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["lifetime_seconds"] = attrs["lifetime_seconds"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOboTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateOboTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenRequest
// only implements ToObjectValue() and Type().
func (m CreateOboTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"application_id":   m.ApplicationId,
			"comment":          m.Comment,
			"lifetime_seconds": m.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOboTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"application_id":   types.StringType,
			"comment":          types.StringType,
			"lifetime_seconds": types.Int64Type,
		},
	}
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo types.Object `tfsdk:"token_info"`
	// Value of the token.
	TokenValue types.String `tfsdk:"token_value"`
}

func (to *CreateOboTokenResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOboTokenResponse) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				// Recursively sync the fields of TokenInfo
				toTokenInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (to *CreateOboTokenResponse) SyncFieldsDuringRead(ctx context.Context, from CreateOboTokenResponse) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				toTokenInfo.SyncFieldsDuringRead(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (m CreateOboTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_info"] = attrs["token_info"].SetOptional()
	attrs["token_value"] = attrs["token_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateOboTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateOboTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenResponse
// only implements ToObjectValue() and Type().
func (m CreateOboTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  m.TokenInfo,
			"token_value": m.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOboTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info":  TokenInfo{}.Type(ctx),
			"token_value": types.StringType,
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in CreateOboTokenResponse as
// a TokenInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateOboTokenResponse) GetTokenInfo(ctx context.Context) (TokenInfo, bool) {
	var e TokenInfo
	if m.TokenInfo.IsNull() || m.TokenInfo.IsUnknown() {
		return e, false
	}
	var v TokenInfo
	d := m.TokenInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfo sets the value of the TokenInfo field in CreateOboTokenResponse.
func (m *CreateOboTokenResponse) SetTokenInfo(ctx context.Context, v TokenInfo) {
	vs := v.ToObjectValue(ctx)
	m.TokenInfo = vs
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type CreatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames types.List `tfsdk:"domain_names"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	EndpointService types.String `tfsdk:"endpoint_service"`
	// Not used by customer-managed private endpoint services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId types.String `tfsdk:"group_id"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames types.List `tfsdk:"resource_names"`
}

func (to *CreatePrivateEndpointRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (to *CreatePrivateEndpointRule) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (m CreatePrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["domain_names"] = attrs["domain_names"].SetOptional()
	attrs["endpoint_service"] = attrs["endpoint_service"].SetOptional()
	attrs["group_id"] = attrs["group_id"].SetOptional()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["resource_names"] = attrs["resource_names"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRule
// only implements ToObjectValue() and Type().
func (m CreatePrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain_names":     m.DomainNames,
			"endpoint_service": m.EndpointService,
			"group_id":         m.GroupId,
			"resource_id":      m.ResourceId,
			"resource_names":   m.ResourceNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePrivateEndpointRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"domain_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"endpoint_service": types.StringType,
			"group_id":         types.StringType,
			"resource_id":      types.StringType,
			"resource_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDomainNames returns the value of the DomainNames field in CreatePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if m.DomainNames.IsNull() || m.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in CreatePrivateEndpointRule.
func (m *CreatePrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in CreatePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if m.ResourceNames.IsNull() || m.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in CreatePrivateEndpointRule.
func (m *CreatePrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`

	PrivateEndpointRule types.Object `tfsdk:"private_endpoint_rule"`
}

func (to *CreatePrivateEndpointRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateEndpointRuleRequest) {
	if !from.PrivateEndpointRule.IsNull() && !from.PrivateEndpointRule.IsUnknown() {
		if toPrivateEndpointRule, ok := to.GetPrivateEndpointRule(ctx); ok {
			if fromPrivateEndpointRule, ok := from.GetPrivateEndpointRule(ctx); ok {
				// Recursively sync the fields of PrivateEndpointRule
				toPrivateEndpointRule.SyncFieldsDuringCreateOrUpdate(ctx, fromPrivateEndpointRule)
				to.SetPrivateEndpointRule(ctx, toPrivateEndpointRule)
			}
		}
	}
}

func (to *CreatePrivateEndpointRuleRequest) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateEndpointRuleRequest) {
	if !from.PrivateEndpointRule.IsNull() && !from.PrivateEndpointRule.IsUnknown() {
		if toPrivateEndpointRule, ok := to.GetPrivateEndpointRule(ctx); ok {
			if fromPrivateEndpointRule, ok := from.GetPrivateEndpointRule(ctx); ok {
				toPrivateEndpointRule.SyncFieldsDuringRead(ctx, fromPrivateEndpointRule)
				to.SetPrivateEndpointRule(ctx, toPrivateEndpointRule)
			}
		}
	}
}

func (m CreatePrivateEndpointRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_endpoint_rule"] = attrs["private_endpoint_rule"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_endpoint_rule": reflect.TypeOf(CreatePrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (m CreatePrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule":          m.PrivateEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule":          CreatePrivateEndpointRule{}.Type(ctx),
		},
	}
}

// GetPrivateEndpointRule returns the value of the PrivateEndpointRule field in CreatePrivateEndpointRuleRequest as
// a CreatePrivateEndpointRule value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateEndpointRuleRequest) GetPrivateEndpointRule(ctx context.Context) (CreatePrivateEndpointRule, bool) {
	var e CreatePrivateEndpointRule
	if m.PrivateEndpointRule.IsNull() || m.PrivateEndpointRule.IsUnknown() {
		return e, false
	}
	var v CreatePrivateEndpointRule
	d := m.PrivateEndpointRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateEndpointRule sets the value of the PrivateEndpointRule field in CreatePrivateEndpointRuleRequest.
func (m *CreatePrivateEndpointRuleRequest) SetPrivateEndpointRule(ctx context.Context, v CreatePrivateEndpointRule) {
	vs := v.ToObjectValue(ctx)
	m.PrivateEndpointRule = vs
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

func (to *CreateTokenRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTokenRequest) {
}

func (to *CreateTokenRequest) SyncFieldsDuringRead(ctx context.Context, from CreateTokenRequest) {
}

func (m CreateTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["lifetime_seconds"] = attrs["lifetime_seconds"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenRequest
// only implements ToObjectValue() and Type().
func (m CreateTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":          m.Comment,
			"lifetime_seconds": m.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":          types.StringType,
			"lifetime_seconds": types.Int64Type,
		},
	}
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo types.Object `tfsdk:"token_info"`
	// The value of the new token.
	TokenValue types.String `tfsdk:"token_value"`
}

func (to *CreateTokenResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTokenResponse) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				// Recursively sync the fields of TokenInfo
				toTokenInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (to *CreateTokenResponse) SyncFieldsDuringRead(ctx context.Context, from CreateTokenResponse) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				toTokenInfo.SyncFieldsDuringRead(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (m CreateTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_info"] = attrs["token_info"].SetOptional()
	attrs["token_value"] = attrs["token_value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(PublicTokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenResponse
// only implements ToObjectValue() and Type().
func (m CreateTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  m.TokenInfo,
			"token_value": m.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info":  PublicTokenInfo{}.Type(ctx),
			"token_value": types.StringType,
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in CreateTokenResponse as
// a PublicTokenInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateTokenResponse) GetTokenInfo(ctx context.Context) (PublicTokenInfo, bool) {
	var e PublicTokenInfo
	if m.TokenInfo.IsNull() || m.TokenInfo.IsUnknown() {
		return e, false
	}
	var v PublicTokenInfo
	d := m.TokenInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfo sets the value of the TokenInfo field in CreateTokenResponse.
func (m *CreateTokenResponse) SetTokenInfo(ctx context.Context, v PublicTokenInfo) {
	vs := v.ToObjectValue(ctx)
	m.TokenInfo = vs
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (to *CspEnablementAccount) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CspEnablementAccount) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (to *CspEnablementAccount) SyncFieldsDuringRead(ctx context.Context, from CspEnablementAccount) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (m CspEnablementAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compliance_standards"] = attrs["compliance_standards"].SetOptional()
	attrs["is_enforced"] = attrs["is_enforced"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CspEnablementAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CspEnablementAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccount
// only implements ToObjectValue() and Type().
func (m CspEnablementAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": m.ComplianceStandards,
			"is_enforced":          m.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CspEnablementAccount) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enforced": types.BoolType,
		},
	}
}

// GetComplianceStandards returns the value of the ComplianceStandards field in CspEnablementAccount as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CspEnablementAccount) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
	if m.ComplianceStandards.IsNull() || m.ComplianceStandards.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ComplianceStandards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceStandards sets the value of the ComplianceStandards field in CspEnablementAccount.
func (m *CspEnablementAccount) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ComplianceStandards = types.ListValueMust(t, vs)
}

type CspEnablementAccountSetting struct {
	CspEnablementAccount types.Object `tfsdk:"csp_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *CspEnablementAccountSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CspEnablementAccountSetting) {
	if !from.CspEnablementAccount.IsNull() && !from.CspEnablementAccount.IsUnknown() {
		if toCspEnablementAccount, ok := to.GetCspEnablementAccount(ctx); ok {
			if fromCspEnablementAccount, ok := from.GetCspEnablementAccount(ctx); ok {
				// Recursively sync the fields of CspEnablementAccount
				toCspEnablementAccount.SyncFieldsDuringCreateOrUpdate(ctx, fromCspEnablementAccount)
				to.SetCspEnablementAccount(ctx, toCspEnablementAccount)
			}
		}
	}
}

func (to *CspEnablementAccountSetting) SyncFieldsDuringRead(ctx context.Context, from CspEnablementAccountSetting) {
	if !from.CspEnablementAccount.IsNull() && !from.CspEnablementAccount.IsUnknown() {
		if toCspEnablementAccount, ok := to.GetCspEnablementAccount(ctx); ok {
			if fromCspEnablementAccount, ok := from.GetCspEnablementAccount(ctx); ok {
				toCspEnablementAccount.SyncFieldsDuringRead(ctx, fromCspEnablementAccount)
				to.SetCspEnablementAccount(ctx, toCspEnablementAccount)
			}
		}
	}
}

func (m CspEnablementAccountSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["csp_enablement_account"] = attrs["csp_enablement_account"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CspEnablementAccountSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CspEnablementAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"csp_enablement_account": reflect.TypeOf(CspEnablementAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccountSetting
// only implements ToObjectValue() and Type().
func (m CspEnablementAccountSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"csp_enablement_account": m.CspEnablementAccount,
			"etag":                   m.Etag,
			"setting_name":           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CspEnablementAccountSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"csp_enablement_account": CspEnablementAccount{}.Type(ctx),
			"etag":                   types.StringType,
			"setting_name":           types.StringType,
		},
	}
}

// GetCspEnablementAccount returns the value of the CspEnablementAccount field in CspEnablementAccountSetting as
// a CspEnablementAccount value.
// If the field is unknown or null, the boolean return value is false.
func (m *CspEnablementAccountSetting) GetCspEnablementAccount(ctx context.Context) (CspEnablementAccount, bool) {
	var e CspEnablementAccount
	if m.CspEnablementAccount.IsNull() || m.CspEnablementAccount.IsUnknown() {
		return e, false
	}
	var v CspEnablementAccount
	d := m.CspEnablementAccount.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCspEnablementAccount sets the value of the CspEnablementAccount field in CspEnablementAccountSetting.
func (m *CspEnablementAccountSetting) SetCspEnablementAccount(ctx context.Context, v CspEnablementAccount) {
	vs := v.ToObjectValue(ctx)
	m.CspEnablementAccount = vs
}

// Properties of the new private endpoint rule. Note that for private endpoints
// towards a VPC endpoint service behind a customer-managed NLB, you must
// approve the endpoint in AWS console after initialization.
type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule struct {
	// Databricks account ID. You can find your account ID from the Accounts
	// Console.
	AccountId types.String `tfsdk:"account_id"`
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the AWS console
	// before they take effect. The possible values are: - PENDING: The endpoint
	// has been created and pending approval. - ESTABLISHED: The endpoint has
	// been approved and is ready to use in your serverless compute resources. -
	// REJECTED: Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up. - EXPIRED: If the endpoint is created but not approved in 14
	// days, it is EXPIRED.
	ConnectionState types.String `tfsdk:"connection_state"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at"`
	// Only used by private endpoints towards a VPC endpoint service for
	// customer-managed VPC endpoint service.
	//
	// The target AWS resource FQDNs accessible via the VPC endpoint service.
	// When updating this field, we perform full update on this field. Please
	// ensure a full list of desired domain_names is provided.
	DomainNames types.List `tfsdk:"domain_names"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	Enabled types.Bool `tfsdk:"enabled"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	EndpointService types.String `tfsdk:"endpoint_service"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames types.List `tfsdk:"resource_names"`
	// The ID of a private endpoint rule.
	RuleId types.String `tfsdk:"rule_id"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time"`
	// The AWS VPC endpoint ID. You can use this ID to identify VPC endpoint
	// created by Databricks.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id"`
}

func (to *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (to *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SyncFieldsDuringRead(ctx context.Context, from CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["connection_state"] = attrs["connection_state"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["deactivated"] = attrs["deactivated"].SetOptional()
	attrs["deactivated_at"] = attrs["deactivated_at"].SetOptional()
	attrs["domain_names"] = attrs["domain_names"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["endpoint_service"] = attrs["endpoint_service"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["resource_names"] = attrs["resource_names"].SetOptional()
	attrs["rule_id"] = attrs["rule_id"].SetOptional()
	attrs["updated_time"] = attrs["updated_time"].SetOptional()
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule
// only implements ToObjectValue() and Type().
func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                     m.AccountId,
			"connection_state":               m.ConnectionState,
			"creation_time":                  m.CreationTime,
			"deactivated":                    m.Deactivated,
			"deactivated_at":                 m.DeactivatedAt,
			"domain_names":                   m.DomainNames,
			"enabled":                        m.Enabled,
			"endpoint_service":               m.EndpointService,
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"resource_names":                 m.ResourceNames,
			"rule_id":                        m.RuleId,
			"updated_time":                   m.UpdatedTime,
			"vpc_endpoint_id":                m.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":       types.StringType,
			"connection_state": types.StringType,
			"creation_time":    types.Int64Type,
			"deactivated":      types.BoolType,
			"deactivated_at":   types.Int64Type,
			"domain_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"enabled":                        types.BoolType,
			"endpoint_service":               types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rule_id":         types.StringType,
			"updated_time":    types.Int64Type,
			"vpc_endpoint_id": types.StringType,
		},
	}
}

// GetDomainNames returns the value of the DomainNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if m.DomainNames.IsNull() || m.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if m.ResourceNames.IsNull() || m.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

type DashboardEmailSubscriptions struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *DashboardEmailSubscriptions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DashboardEmailSubscriptions) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *DashboardEmailSubscriptions) SyncFieldsDuringRead(ctx context.Context, from DashboardEmailSubscriptions) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m DashboardEmailSubscriptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DashboardEmailSubscriptions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DashboardEmailSubscriptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardEmailSubscriptions
// only implements ToObjectValue() and Type().
func (m DashboardEmailSubscriptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DashboardEmailSubscriptions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in DashboardEmailSubscriptions as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *DashboardEmailSubscriptions) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in DashboardEmailSubscriptions.
func (m *DashboardEmailSubscriptions) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

// This represents the setting configuration for the default namespace in the
// Databricks workspace. Setting the default catalog for the workspace
// determines the catalog that is used when queries do not reference a fully
// qualified 3 level name. For example, if the default catalog is set to
// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the object
// 'retail_prod.default.myTable' (the schema 'default' is always assumed). This
// setting requires a restart of clusters and SQL warehouses to take effect.
// Additionally, the default namespace only applies when using Unity
// Catalog-enabled compute.
type DefaultNamespaceSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	Namespace types.Object `tfsdk:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *DefaultNamespaceSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultNamespaceSetting) {
	if !from.Namespace.IsNull() && !from.Namespace.IsUnknown() {
		if toNamespace, ok := to.GetNamespace(ctx); ok {
			if fromNamespace, ok := from.GetNamespace(ctx); ok {
				// Recursively sync the fields of Namespace
				toNamespace.SyncFieldsDuringCreateOrUpdate(ctx, fromNamespace)
				to.SetNamespace(ctx, toNamespace)
			}
		}
	}
}

func (to *DefaultNamespaceSetting) SyncFieldsDuringRead(ctx context.Context, from DefaultNamespaceSetting) {
	if !from.Namespace.IsNull() && !from.Namespace.IsUnknown() {
		if toNamespace, ok := to.GetNamespace(ctx); ok {
			if fromNamespace, ok := from.GetNamespace(ctx); ok {
				toNamespace.SyncFieldsDuringRead(ctx, fromNamespace)
				to.SetNamespace(ctx, toNamespace)
			}
		}
	}
}

func (m DefaultNamespaceSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["namespace"] = attrs["namespace"].SetRequired()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultNamespaceSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DefaultNamespaceSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"namespace": reflect.TypeOf(StringMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultNamespaceSetting
// only implements ToObjectValue() and Type().
func (m DefaultNamespaceSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         m.Etag,
			"namespace":    m.Namespace,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultNamespaceSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":         types.StringType,
			"namespace":    StringMessage{}.Type(ctx),
			"setting_name": types.StringType,
		},
	}
}

// GetNamespace returns the value of the Namespace field in DefaultNamespaceSetting as
// a StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultNamespaceSetting) GetNamespace(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if m.Namespace.IsNull() || m.Namespace.IsUnknown() {
		return e, false
	}
	var v StringMessage
	d := m.Namespace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNamespace sets the value of the Namespace field in DefaultNamespaceSetting.
func (m *DefaultNamespaceSetting) SetNamespace(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	m.Namespace = vs
}

type DefaultWarehouseId struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`

	StringVal types.Object `tfsdk:"string_val"`
}

func (to *DefaultWarehouseId) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultWarehouseId) {
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				// Recursively sync the fields of StringVal
				toStringVal.SyncFieldsDuringCreateOrUpdate(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

func (to *DefaultWarehouseId) SyncFieldsDuringRead(ctx context.Context, from DefaultWarehouseId) {
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				toStringVal.SyncFieldsDuringRead(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

func (m DefaultWarehouseId) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultWarehouseId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DefaultWarehouseId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"string_val": reflect.TypeOf(StringMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultWarehouseId
// only implements ToObjectValue() and Type().
func (m DefaultWarehouseId) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         m.Etag,
			"setting_name": m.SettingName,
			"string_val":   m.StringVal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultWarehouseId) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":         types.StringType,
			"setting_name": types.StringType,
			"string_val":   StringMessage{}.Type(ctx),
		},
	}
}

// GetStringVal returns the value of the StringVal field in DefaultWarehouseId as
// a StringMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultWarehouseId) GetStringVal(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if m.StringVal.IsNull() || m.StringVal.IsUnknown() {
		return e, false
	}
	var v StringMessage
	d := m.StringVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStringVal sets the value of the StringVal field in DefaultWarehouseId.
func (m *DefaultWarehouseId) SetStringVal(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	m.StringVal = vs
}

type DeleteAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteAccountIpAccessEnableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountIpAccessEnableRequest) {
}

func (to *DeleteAccountIpAccessEnableRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountIpAccessEnableRequest) {
}

func (m DeleteAccountIpAccessEnableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAccountIpAccessEnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableRequest
// only implements ToObjectValue() and Type().
func (m DeleteAccountIpAccessEnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountIpAccessEnableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteAccountIpAccessEnableResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteAccountIpAccessEnableResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountIpAccessEnableResponse) {
}

func (to *DeleteAccountIpAccessEnableResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountIpAccessEnableResponse) {
}

func (m DeleteAccountIpAccessEnableResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessEnableResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAccountIpAccessEnableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableResponse
// only implements ToObjectValue() and Type().
func (m DeleteAccountIpAccessEnableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountIpAccessEnableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *DeleteAccountIpAccessListRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountIpAccessListRequest) {
}

func (to *DeleteAccountIpAccessListRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountIpAccessListRequest) {
}

func (m DeleteAccountIpAccessListRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAccountIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessListRequest
// only implements ToObjectValue() and Type().
func (m DeleteAccountIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type DeleteAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingRequest
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingAccessPolicySettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) {
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) {
}

func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingAccessPolicySettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingResponse
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) {
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) {
}

func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDashboardEmailSubscriptionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardEmailSubscriptionsRequest) {
}

func (to *DeleteDashboardEmailSubscriptionsRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardEmailSubscriptionsRequest) {
}

func (m DeleteDashboardEmailSubscriptionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardEmailSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDashboardEmailSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardEmailSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (m DeleteDashboardEmailSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDashboardEmailSubscriptionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDashboardEmailSubscriptionsResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDashboardEmailSubscriptionsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardEmailSubscriptionsResponse) {
}

func (to *DeleteDashboardEmailSubscriptionsResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardEmailSubscriptionsResponse) {
}

func (m DeleteDashboardEmailSubscriptionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardEmailSubscriptionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDashboardEmailSubscriptionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardEmailSubscriptionsResponse
// only implements ToObjectValue() and Type().
func (m DeleteDashboardEmailSubscriptionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDashboardEmailSubscriptionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDefaultNamespaceSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultNamespaceSettingRequest) {
}

func (to *DeleteDefaultNamespaceSettingRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultNamespaceSettingRequest) {
}

func (m DeleteDefaultNamespaceSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingRequest
// only implements ToObjectValue() and Type().
func (m DeleteDefaultNamespaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultNamespaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDefaultNamespaceSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDefaultNamespaceSettingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultNamespaceSettingResponse) {
}

func (to *DeleteDefaultNamespaceSettingResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultNamespaceSettingResponse) {
}

func (m DeleteDefaultNamespaceSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultNamespaceSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDefaultNamespaceSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingResponse
// only implements ToObjectValue() and Type().
func (m DeleteDefaultNamespaceSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultNamespaceSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDefaultWarehouseIdRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDefaultWarehouseIdRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultWarehouseIdRequest) {
}

func (to *DeleteDefaultWarehouseIdRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultWarehouseIdRequest) {
}

func (m DeleteDefaultWarehouseIdRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultWarehouseIdRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDefaultWarehouseIdRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultWarehouseIdRequest
// only implements ToObjectValue() and Type().
func (m DeleteDefaultWarehouseIdRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultWarehouseIdRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDefaultWarehouseIdResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDefaultWarehouseIdResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultWarehouseIdResponse) {
}

func (to *DeleteDefaultWarehouseIdResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultWarehouseIdResponse) {
}

func (m DeleteDefaultWarehouseIdResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultWarehouseIdResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDefaultWarehouseIdResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultWarehouseIdResponse
// only implements ToObjectValue() and Type().
func (m DeleteDefaultWarehouseIdResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultWarehouseIdResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDisableLegacyAccessRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyAccessRequest) {
}

func (to *DeleteDisableLegacyAccessRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyAccessRequest) {
}

func (m DeleteDisableLegacyAccessRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessRequest
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDisableLegacyAccessResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDisableLegacyAccessResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyAccessResponse) {
}

func (to *DeleteDisableLegacyAccessResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyAccessResponse) {
}

func (m DeleteDisableLegacyAccessResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyAccessResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDisableLegacyAccessResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessResponse
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyAccessResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyAccessResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDisableLegacyDbfsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyDbfsRequest) {
}

func (to *DeleteDisableLegacyDbfsRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyDbfsRequest) {
}

func (m DeleteDisableLegacyDbfsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsRequest
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyDbfsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyDbfsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDisableLegacyDbfsResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDisableLegacyDbfsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyDbfsResponse) {
}

func (to *DeleteDisableLegacyDbfsResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyDbfsResponse) {
}

func (m DeleteDisableLegacyDbfsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyDbfsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDisableLegacyDbfsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsResponse
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyDbfsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyDbfsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDisableLegacyFeaturesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyFeaturesRequest) {
}

func (to *DeleteDisableLegacyFeaturesRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyFeaturesRequest) {
}

func (m DeleteDisableLegacyFeaturesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesRequest
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyFeaturesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDisableLegacyFeaturesResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDisableLegacyFeaturesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyFeaturesResponse) {
}

func (to *DeleteDisableLegacyFeaturesResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyFeaturesResponse) {
}

func (m DeleteDisableLegacyFeaturesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyFeaturesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteDisableLegacyFeaturesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesResponse
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyFeaturesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyFeaturesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *DeleteIpAccessListRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIpAccessListRequest) {
}

func (to *DeleteIpAccessListRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteIpAccessListRequest) {
}

func (m DeleteIpAccessListRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIpAccessListRequest
// only implements ToObjectValue() and Type().
func (m DeleteIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type DeleteLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceRequest) {
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceRequest) {
}

func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLlmProxyPartnerPoweredWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLlmProxyPartnerPoweredWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteLlmProxyPartnerPoweredWorkspaceResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceResponse) {
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceResponse) {
}

func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLlmProxyPartnerPoweredWorkspaceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLlmProxyPartnerPoweredWorkspaceResponse
// only implements ToObjectValue() and Type().
func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

func (to *DeleteNetworkConnectivityConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkConnectivityConfigurationRequest) {
}

func (to *DeleteNetworkConnectivityConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkConnectivityConfigurationRequest) {
}

func (m DeleteNetworkConnectivityConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteNetworkConnectivityConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkConnectivityConfigurationRequest
// only implements ToObjectValue() and Type().
func (m DeleteNetworkConnectivityConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNetworkConnectivityConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type DeleteNetworkPolicyRequest struct {
	// The unique identifier of the network policy to delete.
	NetworkPolicyId types.String `tfsdk:"-"`
}

func (to *DeleteNetworkPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkPolicyRequest) {
}

func (to *DeleteNetworkPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkPolicyRequest) {
}

func (m DeleteNetworkPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_policy_id"] = attrs["network_policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (m DeleteNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
		},
	}
}

type DeleteNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteNotificationDestinationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNotificationDestinationRequest) {
}

func (to *DeleteNotificationDestinationRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteNotificationDestinationRequest) {
}

func (m DeleteNotificationDestinationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (m DeleteNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeletePersonalComputeSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePersonalComputeSettingRequest) {
}

func (to *DeletePersonalComputeSettingRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePersonalComputeSettingRequest) {
}

func (m DeletePersonalComputeSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingRequest
// only implements ToObjectValue() and Type().
func (m DeletePersonalComputeSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePersonalComputeSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeletePersonalComputeSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeletePersonalComputeSettingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePersonalComputeSettingResponse) {
}

func (to *DeletePersonalComputeSettingResponse) SyncFieldsDuringRead(ctx context.Context, from DeletePersonalComputeSettingResponse) {
}

func (m DeletePersonalComputeSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePersonalComputeSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePersonalComputeSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingResponse
// only implements ToObjectValue() and Type().
func (m DeletePersonalComputeSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePersonalComputeSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

func (to *DeletePrivateEndpointRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePrivateEndpointRuleRequest) {
}

func (to *DeletePrivateEndpointRuleRequest) SyncFieldsDuringRead(ctx context.Context, from DeletePrivateEndpointRuleRequest) {
}

func (m DeletePrivateEndpointRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()
	attrs["private_endpoint_rule_id"] = attrs["private_endpoint_rule_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeletePrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (m DeletePrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       m.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteRestrictWorkspaceAdminsSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingRequest) {
}

func (to *DeleteRestrictWorkspaceAdminsSettingRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingRequest) {
}

func (m DeleteRestrictWorkspaceAdminsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingRequest
// only implements ToObjectValue() and Type().
func (m DeleteRestrictWorkspaceAdminsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRestrictWorkspaceAdminsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteRestrictWorkspaceAdminsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteRestrictWorkspaceAdminsSettingResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingResponse) {
}

func (to *DeleteRestrictWorkspaceAdminsSettingResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingResponse) {
}

func (m DeleteRestrictWorkspaceAdminsSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRestrictWorkspaceAdminsSettingResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteRestrictWorkspaceAdminsSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingResponse
// only implements ToObjectValue() and Type().
func (m DeleteRestrictWorkspaceAdminsSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRestrictWorkspaceAdminsSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteSqlResultsDownloadRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSqlResultsDownloadRequest) {
}

func (to *DeleteSqlResultsDownloadRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteSqlResultsDownloadRequest) {
}

func (m DeleteSqlResultsDownloadRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSqlResultsDownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSqlResultsDownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSqlResultsDownloadRequest
// only implements ToObjectValue() and Type().
func (m DeleteSqlResultsDownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSqlResultsDownloadRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteSqlResultsDownloadResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteSqlResultsDownloadResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSqlResultsDownloadResponse) {
}

func (to *DeleteSqlResultsDownloadResponse) SyncFieldsDuringRead(ctx context.Context, from DeleteSqlResultsDownloadResponse) {
}

func (m DeleteSqlResultsDownloadResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSqlResultsDownloadResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteSqlResultsDownloadResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSqlResultsDownloadResponse
// only implements ToObjectValue() and Type().
func (m DeleteSqlResultsDownloadResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSqlResultsDownloadResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId types.String `tfsdk:"-"`
}

func (to *DeleteTokenManagementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTokenManagementRequest) {
}

func (to *DeleteTokenManagementRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteTokenManagementRequest) {
}

func (m DeleteTokenManagementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_id"] = attrs["token_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTokenManagementRequest
// only implements ToObjectValue() and Type().
func (m DeleteTokenManagementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTokenManagementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type DisableLegacyAccess struct {
	DisableLegacyAccess types.Object `tfsdk:"disable_legacy_access"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *DisableLegacyAccess) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DisableLegacyAccess) {
	if !from.DisableLegacyAccess.IsNull() && !from.DisableLegacyAccess.IsUnknown() {
		if toDisableLegacyAccess, ok := to.GetDisableLegacyAccess(ctx); ok {
			if fromDisableLegacyAccess, ok := from.GetDisableLegacyAccess(ctx); ok {
				// Recursively sync the fields of DisableLegacyAccess
				toDisableLegacyAccess.SyncFieldsDuringCreateOrUpdate(ctx, fromDisableLegacyAccess)
				to.SetDisableLegacyAccess(ctx, toDisableLegacyAccess)
			}
		}
	}
}

func (to *DisableLegacyAccess) SyncFieldsDuringRead(ctx context.Context, from DisableLegacyAccess) {
	if !from.DisableLegacyAccess.IsNull() && !from.DisableLegacyAccess.IsUnknown() {
		if toDisableLegacyAccess, ok := to.GetDisableLegacyAccess(ctx); ok {
			if fromDisableLegacyAccess, ok := from.GetDisableLegacyAccess(ctx); ok {
				toDisableLegacyAccess.SyncFieldsDuringRead(ctx, fromDisableLegacyAccess)
				to.SetDisableLegacyAccess(ctx, toDisableLegacyAccess)
			}
		}
	}
}

func (m DisableLegacyAccess) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disable_legacy_access"] = attrs["disable_legacy_access"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableLegacyAccess.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DisableLegacyAccess) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_access": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyAccess
// only implements ToObjectValue() and Type().
func (m DisableLegacyAccess) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_access": m.DisableLegacyAccess,
			"etag":                  m.Etag,
			"setting_name":          m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DisableLegacyAccess) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_access": BooleanMessage{}.Type(ctx),
			"etag":                  types.StringType,
			"setting_name":          types.StringType,
		},
	}
}

// GetDisableLegacyAccess returns the value of the DisableLegacyAccess field in DisableLegacyAccess as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *DisableLegacyAccess) GetDisableLegacyAccess(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.DisableLegacyAccess.IsNull() || m.DisableLegacyAccess.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.DisableLegacyAccess.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDisableLegacyAccess sets the value of the DisableLegacyAccess field in DisableLegacyAccess.
func (m *DisableLegacyAccess) SetDisableLegacyAccess(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.DisableLegacyAccess = vs
}

type DisableLegacyDbfs struct {
	DisableLegacyDbfs types.Object `tfsdk:"disable_legacy_dbfs"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *DisableLegacyDbfs) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DisableLegacyDbfs) {
	if !from.DisableLegacyDbfs.IsNull() && !from.DisableLegacyDbfs.IsUnknown() {
		if toDisableLegacyDbfs, ok := to.GetDisableLegacyDbfs(ctx); ok {
			if fromDisableLegacyDbfs, ok := from.GetDisableLegacyDbfs(ctx); ok {
				// Recursively sync the fields of DisableLegacyDbfs
				toDisableLegacyDbfs.SyncFieldsDuringCreateOrUpdate(ctx, fromDisableLegacyDbfs)
				to.SetDisableLegacyDbfs(ctx, toDisableLegacyDbfs)
			}
		}
	}
}

func (to *DisableLegacyDbfs) SyncFieldsDuringRead(ctx context.Context, from DisableLegacyDbfs) {
	if !from.DisableLegacyDbfs.IsNull() && !from.DisableLegacyDbfs.IsUnknown() {
		if toDisableLegacyDbfs, ok := to.GetDisableLegacyDbfs(ctx); ok {
			if fromDisableLegacyDbfs, ok := from.GetDisableLegacyDbfs(ctx); ok {
				toDisableLegacyDbfs.SyncFieldsDuringRead(ctx, fromDisableLegacyDbfs)
				to.SetDisableLegacyDbfs(ctx, toDisableLegacyDbfs)
			}
		}
	}
}

func (m DisableLegacyDbfs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disable_legacy_dbfs"] = attrs["disable_legacy_dbfs"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableLegacyDbfs.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DisableLegacyDbfs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_dbfs": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyDbfs
// only implements ToObjectValue() and Type().
func (m DisableLegacyDbfs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_dbfs": m.DisableLegacyDbfs,
			"etag":                m.Etag,
			"setting_name":        m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DisableLegacyDbfs) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_dbfs": BooleanMessage{}.Type(ctx),
			"etag":                types.StringType,
			"setting_name":        types.StringType,
		},
	}
}

// GetDisableLegacyDbfs returns the value of the DisableLegacyDbfs field in DisableLegacyDbfs as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *DisableLegacyDbfs) GetDisableLegacyDbfs(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.DisableLegacyDbfs.IsNull() || m.DisableLegacyDbfs.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.DisableLegacyDbfs.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDisableLegacyDbfs sets the value of the DisableLegacyDbfs field in DisableLegacyDbfs.
func (m *DisableLegacyDbfs) SetDisableLegacyDbfs(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.DisableLegacyDbfs = vs
}

type DisableLegacyFeatures struct {
	DisableLegacyFeatures types.Object `tfsdk:"disable_legacy_features"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *DisableLegacyFeatures) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DisableLegacyFeatures) {
	if !from.DisableLegacyFeatures.IsNull() && !from.DisableLegacyFeatures.IsUnknown() {
		if toDisableLegacyFeatures, ok := to.GetDisableLegacyFeatures(ctx); ok {
			if fromDisableLegacyFeatures, ok := from.GetDisableLegacyFeatures(ctx); ok {
				// Recursively sync the fields of DisableLegacyFeatures
				toDisableLegacyFeatures.SyncFieldsDuringCreateOrUpdate(ctx, fromDisableLegacyFeatures)
				to.SetDisableLegacyFeatures(ctx, toDisableLegacyFeatures)
			}
		}
	}
}

func (to *DisableLegacyFeatures) SyncFieldsDuringRead(ctx context.Context, from DisableLegacyFeatures) {
	if !from.DisableLegacyFeatures.IsNull() && !from.DisableLegacyFeatures.IsUnknown() {
		if toDisableLegacyFeatures, ok := to.GetDisableLegacyFeatures(ctx); ok {
			if fromDisableLegacyFeatures, ok := from.GetDisableLegacyFeatures(ctx); ok {
				toDisableLegacyFeatures.SyncFieldsDuringRead(ctx, fromDisableLegacyFeatures)
				to.SetDisableLegacyFeatures(ctx, toDisableLegacyFeatures)
			}
		}
	}
}

func (m DisableLegacyFeatures) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disable_legacy_features"] = attrs["disable_legacy_features"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DisableLegacyFeatures.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DisableLegacyFeatures) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_features": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyFeatures
// only implements ToObjectValue() and Type().
func (m DisableLegacyFeatures) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_features": m.DisableLegacyFeatures,
			"etag":                    m.Etag,
			"setting_name":            m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DisableLegacyFeatures) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_features": BooleanMessage{}.Type(ctx),
			"etag":                    types.StringType,
			"setting_name":            types.StringType,
		},
	}
}

// GetDisableLegacyFeatures returns the value of the DisableLegacyFeatures field in DisableLegacyFeatures as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *DisableLegacyFeatures) GetDisableLegacyFeatures(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.DisableLegacyFeatures.IsNull() || m.DisableLegacyFeatures.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.DisableLegacyFeatures.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDisableLegacyFeatures sets the value of the DisableLegacyFeatures field in DisableLegacyFeatures.
func (m *DisableLegacyFeatures) SetDisableLegacyFeatures(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.DisableLegacyFeatures = vs
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess types.Object `tfsdk:"internet_access"`
}

func (to *EgressNetworkPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicy) {
	if !from.InternetAccess.IsNull() && !from.InternetAccess.IsUnknown() {
		if toInternetAccess, ok := to.GetInternetAccess(ctx); ok {
			if fromInternetAccess, ok := from.GetInternetAccess(ctx); ok {
				// Recursively sync the fields of InternetAccess
				toInternetAccess.SyncFieldsDuringCreateOrUpdate(ctx, fromInternetAccess)
				to.SetInternetAccess(ctx, toInternetAccess)
			}
		}
	}
}

func (to *EgressNetworkPolicy) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicy) {
	if !from.InternetAccess.IsNull() && !from.InternetAccess.IsUnknown() {
		if toInternetAccess, ok := to.GetInternetAccess(ctx); ok {
			if fromInternetAccess, ok := from.GetInternetAccess(ctx); ok {
				toInternetAccess.SyncFieldsDuringRead(ctx, fromInternetAccess)
				to.SetInternetAccess(ctx, toInternetAccess)
			}
		}
	}
}

func (m EgressNetworkPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internet_access"] = attrs["internet_access"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internet_access": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicy
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internet_access": m.InternetAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internet_access": EgressNetworkPolicyInternetAccessPolicy{}.Type(ctx),
		},
	}
}

// GetInternetAccess returns the value of the InternetAccess field in EgressNetworkPolicy as
// a EgressNetworkPolicyInternetAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicy) GetInternetAccess(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicy, bool) {
	var e EgressNetworkPolicyInternetAccessPolicy
	if m.InternetAccess.IsNull() || m.InternetAccess.IsUnknown() {
		return e, false
	}
	var v EgressNetworkPolicyInternetAccessPolicy
	d := m.InternetAccess.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInternetAccess sets the value of the InternetAccess field in EgressNetworkPolicy.
func (m *EgressNetworkPolicy) SetInternetAccess(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.InternetAccess = vs
}

type EgressNetworkPolicyInternetAccessPolicy struct {
	AllowedInternetDestinations types.List `tfsdk:"allowed_internet_destinations"`

	AllowedStorageDestinations types.List `tfsdk:"allowed_storage_destinations"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode types.Object `tfsdk:"log_only_mode"`

	RestrictionMode types.String `tfsdk:"restriction_mode"`
}

func (to *EgressNetworkPolicyInternetAccessPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicy) {
	if !from.AllowedInternetDestinations.IsNull() && !from.AllowedInternetDestinations.IsUnknown() && to.AllowedInternetDestinations.IsNull() && len(from.AllowedInternetDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedInternetDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedInternetDestinations = from.AllowedInternetDestinations
	}
	if !from.AllowedStorageDestinations.IsNull() && !from.AllowedStorageDestinations.IsUnknown() && to.AllowedStorageDestinations.IsNull() && len(from.AllowedStorageDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedStorageDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedStorageDestinations = from.AllowedStorageDestinations
	}
	if !from.LogOnlyMode.IsNull() && !from.LogOnlyMode.IsUnknown() {
		if toLogOnlyMode, ok := to.GetLogOnlyMode(ctx); ok {
			if fromLogOnlyMode, ok := from.GetLogOnlyMode(ctx); ok {
				// Recursively sync the fields of LogOnlyMode
				toLogOnlyMode.SyncFieldsDuringCreateOrUpdate(ctx, fromLogOnlyMode)
				to.SetLogOnlyMode(ctx, toLogOnlyMode)
			}
		}
	}
}

func (to *EgressNetworkPolicyInternetAccessPolicy) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicy) {
	if !from.AllowedInternetDestinations.IsNull() && !from.AllowedInternetDestinations.IsUnknown() && to.AllowedInternetDestinations.IsNull() && len(from.AllowedInternetDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedInternetDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedInternetDestinations = from.AllowedInternetDestinations
	}
	if !from.AllowedStorageDestinations.IsNull() && !from.AllowedStorageDestinations.IsUnknown() && to.AllowedStorageDestinations.IsNull() && len(from.AllowedStorageDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedStorageDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedStorageDestinations = from.AllowedStorageDestinations
	}
	if !from.LogOnlyMode.IsNull() && !from.LogOnlyMode.IsUnknown() {
		if toLogOnlyMode, ok := to.GetLogOnlyMode(ctx); ok {
			if fromLogOnlyMode, ok := from.GetLogOnlyMode(ctx); ok {
				toLogOnlyMode.SyncFieldsDuringRead(ctx, fromLogOnlyMode)
				to.SetLogOnlyMode(ctx, toLogOnlyMode)
			}
		}
	}
}

func (m EgressNetworkPolicyInternetAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_internet_destinations"] = attrs["allowed_internet_destinations"].SetOptional()
	attrs["allowed_storage_destinations"] = attrs["allowed_storage_destinations"].SetOptional()
	attrs["log_only_mode"] = attrs["log_only_mode"].SetOptional()
	attrs["restriction_mode"] = attrs["restriction_mode"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyInternetAccessPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyInternetAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyInternetDestination{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyStorageDestination{}),
		"log_only_mode":                 reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyLogOnlyMode{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicy
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_internet_destinations": m.AllowedInternetDestinations,
			"allowed_storage_destinations":  m.AllowedStorageDestinations,
			"log_only_mode":                 m.LogOnlyMode,
			"restriction_mode":              m.RestrictionMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyInternetAccessPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_internet_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyInternetAccessPolicyInternetDestination{}.Type(ctx),
			},
			"allowed_storage_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyInternetAccessPolicyStorageDestination{}.Type(ctx),
			},
			"log_only_mode":    EgressNetworkPolicyInternetAccessPolicyLogOnlyMode{}.Type(ctx),
			"restriction_mode": types.StringType,
		},
	}
}

// GetAllowedInternetDestinations returns the value of the AllowedInternetDestinations field in EgressNetworkPolicyInternetAccessPolicy as
// a slice of EgressNetworkPolicyInternetAccessPolicyInternetDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicy) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyInternetDestination, bool) {
	if m.AllowedInternetDestinations.IsNull() || m.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyInternetDestination
	d := m.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyInternetAccessPolicy.
func (m *EgressNetworkPolicyInternetAccessPolicy) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyInternetDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy as
// a slice of EgressNetworkPolicyInternetAccessPolicyStorageDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicy) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyStorageDestination, bool) {
	if m.AllowedStorageDestinations.IsNull() || m.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyStorageDestination
	d := m.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy.
func (m *EgressNetworkPolicyInternetAccessPolicy) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyStorageDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetLogOnlyMode returns the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy as
// a EgressNetworkPolicyInternetAccessPolicyLogOnlyMode value.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicy) GetLogOnlyMode(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicyLogOnlyMode, bool) {
	var e EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
	if m.LogOnlyMode.IsNull() || m.LogOnlyMode.IsUnknown() {
		return e, false
	}
	var v EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
	d := m.LogOnlyMode.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetLogOnlyMode sets the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy.
func (m *EgressNetworkPolicyInternetAccessPolicy) SetLogOnlyMode(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) {
	vs := v.ToObjectValue(ctx)
	m.LogOnlyMode = vs
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination struct {
	Destination types.String `tfsdk:"destination"`

	Protocol types.String `tfsdk:"protocol"`

	Type_ types.String `tfsdk:"type"`
}

func (to *EgressNetworkPolicyInternetAccessPolicyInternetDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyInternetDestination) {
}

func (to *EgressNetworkPolicyInternetAccessPolicyInternetDestination) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyInternetDestination) {
}

func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetOptional()
	attrs["protocol"] = attrs["protocol"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyInternetAccessPolicyInternetDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyInternetDestination
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
			"protocol":    m.Protocol,
			"type":        m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
			"protocol":    types.StringType,
			"type":        types.StringType,
		},
	}
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {
	LogOnlyModeType types.String `tfsdk:"log_only_mode_type"`

	Workloads types.List `tfsdk:"workloads"`
}

func (to *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) {
	if !from.Workloads.IsNull() && !from.Workloads.IsUnknown() && to.Workloads.IsNull() && len(from.Workloads.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Workloads, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Workloads = from.Workloads
	}
}

func (to *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) {
	if !from.Workloads.IsNull() && !from.Workloads.IsUnknown() && to.Workloads.IsNull() && len(from.Workloads.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Workloads, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Workloads = from.Workloads
	}
}

func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["log_only_mode_type"] = attrs["log_only_mode_type"].SetOptional()
	attrs["workloads"] = attrs["workloads"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workloads": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_only_mode_type": m.LogOnlyModeType,
			"workloads":          m.Workloads,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_only_mode_type": types.StringType,
			"workloads": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetWorkloads returns the value of the Workloads field in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) GetWorkloads(ctx context.Context) ([]types.String, bool) {
	if m.Workloads.IsNull() || m.Workloads.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Workloads.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloads sets the value of the Workloads field in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode.
func (m *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) SetWorkloads(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workloads"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Workloads = types.ListValueMust(t, vs)
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination struct {
	AllowedPaths types.List `tfsdk:"allowed_paths"`

	AzureContainer types.String `tfsdk:"azure_container"`

	AzureDnsZone types.String `tfsdk:"azure_dns_zone"`

	AzureStorageAccount types.String `tfsdk:"azure_storage_account"`

	AzureStorageService types.String `tfsdk:"azure_storage_service"`

	BucketName types.String `tfsdk:"bucket_name"`

	Region types.String `tfsdk:"region"`

	Type_ types.String `tfsdk:"type"`
}

func (to *EgressNetworkPolicyInternetAccessPolicyStorageDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyStorageDestination) {
	if !from.AllowedPaths.IsNull() && !from.AllowedPaths.IsUnknown() && to.AllowedPaths.IsNull() && len(from.AllowedPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedPaths = from.AllowedPaths
	}
}

func (to *EgressNetworkPolicyInternetAccessPolicyStorageDestination) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyStorageDestination) {
	if !from.AllowedPaths.IsNull() && !from.AllowedPaths.IsUnknown() && to.AllowedPaths.IsNull() && len(from.AllowedPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedPaths = from.AllowedPaths
	}
}

func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_paths"] = attrs["allowed_paths"].SetOptional()
	attrs["azure_container"] = attrs["azure_container"].SetOptional()
	attrs["azure_dns_zone"] = attrs["azure_dns_zone"].SetOptional()
	attrs["azure_storage_account"] = attrs["azure_storage_account"].SetOptional()
	attrs["azure_storage_service"] = attrs["azure_storage_service"].SetOptional()
	attrs["bucket_name"] = attrs["bucket_name"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["type"] = attrs["type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyInternetAccessPolicyStorageDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_paths": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyStorageDestination
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_paths":         m.AllowedPaths,
			"azure_container":       m.AzureContainer,
			"azure_dns_zone":        m.AzureDnsZone,
			"azure_storage_account": m.AzureStorageAccount,
			"azure_storage_service": m.AzureStorageService,
			"bucket_name":           m.BucketName,
			"region":                m.Region,
			"type":                  m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_paths": basetypes.ListType{
				ElemType: types.StringType,
			},
			"azure_container":       types.StringType,
			"azure_dns_zone":        types.StringType,
			"azure_storage_account": types.StringType,
			"azure_storage_service": types.StringType,
			"bucket_name":           types.StringType,
			"region":                types.StringType,
			"type":                  types.StringType,
		},
	}
}

// GetAllowedPaths returns the value of the AllowedPaths field in EgressNetworkPolicyInternetAccessPolicyStorageDestination as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicyStorageDestination) GetAllowedPaths(ctx context.Context) ([]types.String, bool) {
	if m.AllowedPaths.IsNull() || m.AllowedPaths.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.AllowedPaths.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedPaths sets the value of the AllowedPaths field in EgressNetworkPolicyInternetAccessPolicyStorageDestination.
func (m *EgressNetworkPolicyInternetAccessPolicyStorageDestination) SetAllowedPaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedPaths = types.ListValueMust(t, vs)
}

type EgressNetworkPolicyNetworkAccessPolicy struct {
	// List of internet destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	AllowedInternetDestinations types.List `tfsdk:"allowed_internet_destinations"`
	// List of storage destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	AllowedStorageDestinations types.List `tfsdk:"allowed_storage_destinations"`
	// Optional. When policy_enforcement is not provided, we default to
	// ENFORCE_MODE_ALL_SERVICES
	PolicyEnforcement types.Object `tfsdk:"policy_enforcement"`
	// The restriction mode that controls how serverless workloads can access
	// the internet.
	RestrictionMode types.String `tfsdk:"restriction_mode"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicy) {
	if !from.AllowedInternetDestinations.IsNull() && !from.AllowedInternetDestinations.IsUnknown() && to.AllowedInternetDestinations.IsNull() && len(from.AllowedInternetDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedInternetDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedInternetDestinations = from.AllowedInternetDestinations
	}
	if !from.AllowedStorageDestinations.IsNull() && !from.AllowedStorageDestinations.IsUnknown() && to.AllowedStorageDestinations.IsNull() && len(from.AllowedStorageDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedStorageDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedStorageDestinations = from.AllowedStorageDestinations
	}
	if !from.PolicyEnforcement.IsNull() && !from.PolicyEnforcement.IsUnknown() {
		if toPolicyEnforcement, ok := to.GetPolicyEnforcement(ctx); ok {
			if fromPolicyEnforcement, ok := from.GetPolicyEnforcement(ctx); ok {
				// Recursively sync the fields of PolicyEnforcement
				toPolicyEnforcement.SyncFieldsDuringCreateOrUpdate(ctx, fromPolicyEnforcement)
				to.SetPolicyEnforcement(ctx, toPolicyEnforcement)
			}
		}
	}
}

func (to *EgressNetworkPolicyNetworkAccessPolicy) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicy) {
	if !from.AllowedInternetDestinations.IsNull() && !from.AllowedInternetDestinations.IsUnknown() && to.AllowedInternetDestinations.IsNull() && len(from.AllowedInternetDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedInternetDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedInternetDestinations = from.AllowedInternetDestinations
	}
	if !from.AllowedStorageDestinations.IsNull() && !from.AllowedStorageDestinations.IsUnknown() && to.AllowedStorageDestinations.IsNull() && len(from.AllowedStorageDestinations.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedStorageDestinations, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedStorageDestinations = from.AllowedStorageDestinations
	}
	if !from.PolicyEnforcement.IsNull() && !from.PolicyEnforcement.IsUnknown() {
		if toPolicyEnforcement, ok := to.GetPolicyEnforcement(ctx); ok {
			if fromPolicyEnforcement, ok := from.GetPolicyEnforcement(ctx); ok {
				toPolicyEnforcement.SyncFieldsDuringRead(ctx, fromPolicyEnforcement)
				to.SetPolicyEnforcement(ctx, toPolicyEnforcement)
			}
		}
	}
}

func (m EgressNetworkPolicyNetworkAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_internet_destinations"] = attrs["allowed_internet_destinations"].SetOptional()
	attrs["allowed_storage_destinations"] = attrs["allowed_storage_destinations"].SetOptional()
	attrs["policy_enforcement"] = attrs["policy_enforcement"].SetOptional()
	attrs["restriction_mode"] = attrs["restriction_mode"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyNetworkAccessPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyNetworkAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyInternetDestination{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyStorageDestination{}),
		"policy_enforcement":            reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicy
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_internet_destinations": m.AllowedInternetDestinations,
			"allowed_storage_destinations":  m.AllowedStorageDestinations,
			"policy_enforcement":            m.PolicyEnforcement,
			"restriction_mode":              m.RestrictionMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyNetworkAccessPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_internet_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyNetworkAccessPolicyInternetDestination{}.Type(ctx),
			},
			"allowed_storage_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyNetworkAccessPolicyStorageDestination{}.Type(ctx),
			},
			"policy_enforcement": EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement{}.Type(ctx),
			"restriction_mode":   types.StringType,
		},
	}
}

// GetAllowedInternetDestinations returns the value of the AllowedInternetDestinations field in EgressNetworkPolicyNetworkAccessPolicy as
// a slice of EgressNetworkPolicyNetworkAccessPolicyInternetDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicy) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyNetworkAccessPolicyInternetDestination, bool) {
	if m.AllowedInternetDestinations.IsNull() || m.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyInternetDestination
	d := m.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyNetworkAccessPolicy.
func (m *EgressNetworkPolicyNetworkAccessPolicy) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyNetworkAccessPolicyInternetDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyNetworkAccessPolicy as
// a slice of EgressNetworkPolicyNetworkAccessPolicyStorageDestination values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicy) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyNetworkAccessPolicyStorageDestination, bool) {
	if m.AllowedStorageDestinations.IsNull() || m.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyStorageDestination
	d := m.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyNetworkAccessPolicy.
func (m *EgressNetworkPolicyNetworkAccessPolicy) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyNetworkAccessPolicyStorageDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetPolicyEnforcement returns the value of the PolicyEnforcement field in EgressNetworkPolicyNetworkAccessPolicy as
// a EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement value.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicy) GetPolicyEnforcement(ctx context.Context) (EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement, bool) {
	var e EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
	if m.PolicyEnforcement.IsNull() || m.PolicyEnforcement.IsUnknown() {
		return e, false
	}
	var v EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
	d := m.PolicyEnforcement.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPolicyEnforcement sets the value of the PolicyEnforcement field in EgressNetworkPolicyNetworkAccessPolicy.
func (m *EgressNetworkPolicyNetworkAccessPolicy) SetPolicyEnforcement(ctx context.Context, v EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) {
	vs := v.ToObjectValue(ctx)
	m.PolicyEnforcement = vs
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support DNS_NAME (FQDN format) destinations for the time
// being. Going forward we may extend support to host names and IP addresses.
type EgressNetworkPolicyNetworkAccessPolicyInternetDestination struct {
	// The internet destination to which access will be allowed. Format
	// dependent on the destination type.
	Destination types.String `tfsdk:"destination"`
	// The type of internet destination. Currently only DNS_NAME is supported.
	InternetDestinationType types.String `tfsdk:"internet_destination_type"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyInternetDestination) {
}

func (to *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyInternetDestination) {
}

func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination"] = attrs["destination"].SetOptional()
	attrs["internet_destination_type"] = attrs["internet_destination_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyNetworkAccessPolicyInternetDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyInternetDestination
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination":               m.Destination,
			"internet_destination_type": m.InternetDestinationType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination":               types.StringType,
			"internet_destination_type": types.StringType,
		},
	}
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement struct {
	// When empty, it means dry run for all products. When non-empty, it means
	// dry run for specific products and for the other products, they will run
	// in enforced mode.
	DryRunModeProductFilter types.List `tfsdk:"dry_run_mode_product_filter"`
	// The mode of policy enforcement. ENFORCED blocks traffic that violates
	// policy, while DRY_RUN only logs violations without blocking. When not
	// specified, defaults to ENFORCED.
	EnforcementMode types.String `tfsdk:"enforcement_mode"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) {
	if !from.DryRunModeProductFilter.IsNull() && !from.DryRunModeProductFilter.IsUnknown() && to.DryRunModeProductFilter.IsNull() && len(from.DryRunModeProductFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DryRunModeProductFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DryRunModeProductFilter = from.DryRunModeProductFilter
	}
}

func (to *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) {
	if !from.DryRunModeProductFilter.IsNull() && !from.DryRunModeProductFilter.IsUnknown() && to.DryRunModeProductFilter.IsNull() && len(from.DryRunModeProductFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DryRunModeProductFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DryRunModeProductFilter = from.DryRunModeProductFilter
	}
}

func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["dry_run_mode_product_filter"] = attrs["dry_run_mode_product_filter"].SetOptional()
	attrs["enforcement_mode"] = attrs["enforcement_mode"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dry_run_mode_product_filter": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dry_run_mode_product_filter": m.DryRunModeProductFilter,
			"enforcement_mode":            m.EnforcementMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dry_run_mode_product_filter": basetypes.ListType{
				ElemType: types.StringType,
			},
			"enforcement_mode": types.StringType,
		},
	}
}

// GetDryRunModeProductFilter returns the value of the DryRunModeProductFilter field in EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) GetDryRunModeProductFilter(ctx context.Context) ([]types.String, bool) {
	if m.DryRunModeProductFilter.IsNull() || m.DryRunModeProductFilter.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DryRunModeProductFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDryRunModeProductFilter sets the value of the DryRunModeProductFilter field in EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement.
func (m *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) SetDryRunModeProductFilter(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dry_run_mode_product_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DryRunModeProductFilter = types.ListValueMust(t, vs)
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyNetworkAccessPolicyStorageDestination struct {
	// The Azure storage account name.
	AzureStorageAccount types.String `tfsdk:"azure_storage_account"`
	// The Azure storage service type (blob, dfs, etc.).
	AzureStorageService types.String `tfsdk:"azure_storage_service"`

	BucketName types.String `tfsdk:"bucket_name"`

	Region types.String `tfsdk:"region"`
	// The type of storage destination.
	StorageDestinationType types.String `tfsdk:"storage_destination_type"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyStorageDestination) {
}

func (to *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyStorageDestination) {
}

func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["azure_storage_account"] = attrs["azure_storage_account"].SetOptional()
	attrs["azure_storage_service"] = attrs["azure_storage_service"].SetOptional()
	attrs["bucket_name"] = attrs["bucket_name"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["storage_destination_type"] = attrs["storage_destination_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicyNetworkAccessPolicyStorageDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyStorageDestination
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_storage_account":    m.AzureStorageAccount,
			"azure_storage_service":    m.AzureStorageService,
			"bucket_name":              m.BucketName,
			"region":                   m.Region,
			"storage_destination_type": m.StorageDestinationType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_storage_account":    types.StringType,
			"azure_storage_service":    types.StringType,
			"bucket_name":              types.StringType,
			"region":                   types.StringType,
			"storage_destination_type": types.StringType,
		},
	}
}

type EmailConfig struct {
	// Email addresses to notify.
	Addresses types.List `tfsdk:"addresses"`
}

func (to *EmailConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmailConfig) {
	if !from.Addresses.IsNull() && !from.Addresses.IsUnknown() && to.Addresses.IsNull() && len(from.Addresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Addresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Addresses = from.Addresses
	}
}

func (to *EmailConfig) SyncFieldsDuringRead(ctx context.Context, from EmailConfig) {
	if !from.Addresses.IsNull() && !from.Addresses.IsUnknown() && to.Addresses.IsNull() && len(from.Addresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Addresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Addresses = from.Addresses
	}
}

func (m EmailConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["addresses"] = attrs["addresses"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmailConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmailConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmailConfig
// only implements ToObjectValue() and Type().
func (m EmailConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"addresses": m.Addresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmailConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAddresses returns the value of the Addresses field in EmailConfig as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EmailConfig) GetAddresses(ctx context.Context) ([]types.String, bool) {
	if m.Addresses.IsNull() || m.Addresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Addresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAddresses sets the value of the Addresses field in EmailConfig.
func (m *EmailConfig) SetAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Addresses = types.ListValueMust(t, vs)
}

type Empty struct {
}

func (to *Empty) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Empty) {
}

func (to *Empty) SyncFieldsDuringRead(ctx context.Context, from Empty) {
}

func (m Empty) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Empty) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty
// only implements ToObjectValue() and Type().
func (m Empty) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m Empty) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EnableExportNotebook struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnableExportNotebook) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnableExportNotebook) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *EnableExportNotebook) SyncFieldsDuringRead(ctx context.Context, from EnableExportNotebook) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m EnableExportNotebook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableExportNotebook.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnableExportNotebook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableExportNotebook
// only implements ToObjectValue() and Type().
func (m EnableExportNotebook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnableExportNotebook) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in EnableExportNotebook as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnableExportNotebook) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in EnableExportNotebook.
func (m *EnableExportNotebook) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

type EnableNotebookTableClipboard struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnableNotebookTableClipboard) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnableNotebookTableClipboard) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *EnableNotebookTableClipboard) SyncFieldsDuringRead(ctx context.Context, from EnableNotebookTableClipboard) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m EnableNotebookTableClipboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableNotebookTableClipboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnableNotebookTableClipboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableNotebookTableClipboard
// only implements ToObjectValue() and Type().
func (m EnableNotebookTableClipboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnableNotebookTableClipboard) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in EnableNotebookTableClipboard as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnableNotebookTableClipboard) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in EnableNotebookTableClipboard.
func (m *EnableNotebookTableClipboard) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

type EnableResultsDownloading struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnableResultsDownloading) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnableResultsDownloading) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *EnableResultsDownloading) SyncFieldsDuringRead(ctx context.Context, from EnableResultsDownloading) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m EnableResultsDownloading) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnableResultsDownloading.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnableResultsDownloading) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableResultsDownloading
// only implements ToObjectValue() and Type().
func (m EnableResultsDownloading) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnableResultsDownloading) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in EnableResultsDownloading as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnableResultsDownloading) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in EnableResultsDownloading.
func (m *EnableResultsDownloading) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (to *EnhancedSecurityMonitoring) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnhancedSecurityMonitoring) {
}

func (to *EnhancedSecurityMonitoring) SyncFieldsDuringRead(ctx context.Context, from EnhancedSecurityMonitoring) {
}

func (m EnhancedSecurityMonitoring) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_enabled"] = attrs["is_enabled"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnhancedSecurityMonitoring.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnhancedSecurityMonitoring) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoring
// only implements ToObjectValue() and Type().
func (m EnhancedSecurityMonitoring) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enabled": m.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnhancedSecurityMonitoring) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enabled": types.BoolType,
		},
	}
}

type EnhancedSecurityMonitoringSetting struct {
	EnhancedSecurityMonitoringWorkspace types.Object `tfsdk:"enhanced_security_monitoring_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnhancedSecurityMonitoringSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnhancedSecurityMonitoringSetting) {
	if !from.EnhancedSecurityMonitoringWorkspace.IsNull() && !from.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		if toEnhancedSecurityMonitoringWorkspace, ok := to.GetEnhancedSecurityMonitoringWorkspace(ctx); ok {
			if fromEnhancedSecurityMonitoringWorkspace, ok := from.GetEnhancedSecurityMonitoringWorkspace(ctx); ok {
				// Recursively sync the fields of EnhancedSecurityMonitoringWorkspace
				toEnhancedSecurityMonitoringWorkspace.SyncFieldsDuringCreateOrUpdate(ctx, fromEnhancedSecurityMonitoringWorkspace)
				to.SetEnhancedSecurityMonitoringWorkspace(ctx, toEnhancedSecurityMonitoringWorkspace)
			}
		}
	}
}

func (to *EnhancedSecurityMonitoringSetting) SyncFieldsDuringRead(ctx context.Context, from EnhancedSecurityMonitoringSetting) {
	if !from.EnhancedSecurityMonitoringWorkspace.IsNull() && !from.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		if toEnhancedSecurityMonitoringWorkspace, ok := to.GetEnhancedSecurityMonitoringWorkspace(ctx); ok {
			if fromEnhancedSecurityMonitoringWorkspace, ok := from.GetEnhancedSecurityMonitoringWorkspace(ctx); ok {
				toEnhancedSecurityMonitoringWorkspace.SyncFieldsDuringRead(ctx, fromEnhancedSecurityMonitoringWorkspace)
				to.SetEnhancedSecurityMonitoringWorkspace(ctx, toEnhancedSecurityMonitoringWorkspace)
			}
		}
	}
}

func (m EnhancedSecurityMonitoringSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enhanced_security_monitoring_workspace"] = attrs["enhanced_security_monitoring_workspace"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EnhancedSecurityMonitoringSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EnhancedSecurityMonitoringSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enhanced_security_monitoring_workspace": reflect.TypeOf(EnhancedSecurityMonitoring{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoringSetting
// only implements ToObjectValue() and Type().
func (m EnhancedSecurityMonitoringSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enhanced_security_monitoring_workspace": m.EnhancedSecurityMonitoringWorkspace,
			"etag":                                   m.Etag,
			"setting_name":                           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnhancedSecurityMonitoringSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enhanced_security_monitoring_workspace": EnhancedSecurityMonitoring{}.Type(ctx),
			"etag":                                   types.StringType,
			"setting_name":                           types.StringType,
		},
	}
}

// GetEnhancedSecurityMonitoringWorkspace returns the value of the EnhancedSecurityMonitoringWorkspace field in EnhancedSecurityMonitoringSetting as
// a EnhancedSecurityMonitoring value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnhancedSecurityMonitoringSetting) GetEnhancedSecurityMonitoringWorkspace(ctx context.Context) (EnhancedSecurityMonitoring, bool) {
	var e EnhancedSecurityMonitoring
	if m.EnhancedSecurityMonitoringWorkspace.IsNull() || m.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		return e, false
	}
	var v EnhancedSecurityMonitoring
	d := m.EnhancedSecurityMonitoringWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnhancedSecurityMonitoringWorkspace sets the value of the EnhancedSecurityMonitoringWorkspace field in EnhancedSecurityMonitoringSetting.
func (m *EnhancedSecurityMonitoringSetting) SetEnhancedSecurityMonitoringWorkspace(ctx context.Context, v EnhancedSecurityMonitoring) {
	vs := v.ToObjectValue(ctx)
	m.EnhancedSecurityMonitoringWorkspace = vs
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (to *EsmEnablementAccount) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EsmEnablementAccount) {
}

func (to *EsmEnablementAccount) SyncFieldsDuringRead(ctx context.Context, from EsmEnablementAccount) {
}

func (m EsmEnablementAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["is_enforced"] = attrs["is_enforced"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EsmEnablementAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EsmEnablementAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccount
// only implements ToObjectValue() and Type().
func (m EsmEnablementAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enforced": m.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EsmEnablementAccount) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enforced": types.BoolType,
		},
	}
}

type EsmEnablementAccountSetting struct {
	EsmEnablementAccount types.Object `tfsdk:"esm_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EsmEnablementAccountSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EsmEnablementAccountSetting) {
	if !from.EsmEnablementAccount.IsNull() && !from.EsmEnablementAccount.IsUnknown() {
		if toEsmEnablementAccount, ok := to.GetEsmEnablementAccount(ctx); ok {
			if fromEsmEnablementAccount, ok := from.GetEsmEnablementAccount(ctx); ok {
				// Recursively sync the fields of EsmEnablementAccount
				toEsmEnablementAccount.SyncFieldsDuringCreateOrUpdate(ctx, fromEsmEnablementAccount)
				to.SetEsmEnablementAccount(ctx, toEsmEnablementAccount)
			}
		}
	}
}

func (to *EsmEnablementAccountSetting) SyncFieldsDuringRead(ctx context.Context, from EsmEnablementAccountSetting) {
	if !from.EsmEnablementAccount.IsNull() && !from.EsmEnablementAccount.IsUnknown() {
		if toEsmEnablementAccount, ok := to.GetEsmEnablementAccount(ctx); ok {
			if fromEsmEnablementAccount, ok := from.GetEsmEnablementAccount(ctx); ok {
				toEsmEnablementAccount.SyncFieldsDuringRead(ctx, fromEsmEnablementAccount)
				to.SetEsmEnablementAccount(ctx, toEsmEnablementAccount)
			}
		}
	}
}

func (m EsmEnablementAccountSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["esm_enablement_account"] = attrs["esm_enablement_account"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EsmEnablementAccountSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EsmEnablementAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"esm_enablement_account": reflect.TypeOf(EsmEnablementAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccountSetting
// only implements ToObjectValue() and Type().
func (m EsmEnablementAccountSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"esm_enablement_account": m.EsmEnablementAccount,
			"etag":                   m.Etag,
			"setting_name":           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EsmEnablementAccountSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"esm_enablement_account": EsmEnablementAccount{}.Type(ctx),
			"etag":                   types.StringType,
			"setting_name":           types.StringType,
		},
	}
}

// GetEsmEnablementAccount returns the value of the EsmEnablementAccount field in EsmEnablementAccountSetting as
// a EsmEnablementAccount value.
// If the field is unknown or null, the boolean return value is false.
func (m *EsmEnablementAccountSetting) GetEsmEnablementAccount(ctx context.Context) (EsmEnablementAccount, bool) {
	var e EsmEnablementAccount
	if m.EsmEnablementAccount.IsNull() || m.EsmEnablementAccount.IsUnknown() {
		return e, false
	}
	var v EsmEnablementAccount
	d := m.EsmEnablementAccount.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEsmEnablementAccount sets the value of the EsmEnablementAccount field in EsmEnablementAccountSetting.
func (m *EsmEnablementAccountSetting) SetEsmEnablementAccount(ctx context.Context, v EsmEnablementAccount) {
	vs := v.ToObjectValue(ctx)
	m.EsmEnablementAccount = vs
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential types.String `tfsdk:"credential"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime types.Int64 `tfsdk:"credential_eol_time"`
	// User ID of the user that owns this token.
	OwnerId types.Int64 `tfsdk:"owner_id"`
	// The scopes of access granted in the token.
	Scopes types.List `tfsdk:"scopes"`
	// The type of this exchange token
	TokenType types.String `tfsdk:"token_type"`
}

func (to *ExchangeToken) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeToken) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (to *ExchangeToken) SyncFieldsDuringRead(ctx context.Context, from ExchangeToken) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (m ExchangeToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential"] = attrs["credential"].SetOptional()
	attrs["credential_eol_time"] = attrs["credential_eol_time"].SetOptional()
	attrs["owner_id"] = attrs["owner_id"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["token_type"] = attrs["token_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExchangeToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeToken
// only implements ToObjectValue() and Type().
func (m ExchangeToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential":          m.Credential,
			"credential_eol_time": m.CredentialEolTime,
			"owner_id":            m.OwnerId,
			"scopes":              m.Scopes,
			"token_type":          m.TokenType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential":          types.StringType,
			"credential_eol_time": types.Int64Type,
			"owner_id":            types.Int64Type,
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_type": types.StringType,
		},
	}
}

// GetScopes returns the value of the Scopes field in ExchangeToken as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeToken) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeToken.
func (m *ExchangeToken) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId types.Object `tfsdk:"partition_id"`
	// Array of scopes for the token request.
	Scopes types.List `tfsdk:"scopes"`
	// A list of token types being requested
	TokenType types.List `tfsdk:"token_type"`
}

func (to *ExchangeTokenRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeTokenRequest) {
	if !from.PartitionId.IsNull() && !from.PartitionId.IsUnknown() {
		if toPartitionId, ok := to.GetPartitionId(ctx); ok {
			if fromPartitionId, ok := from.GetPartitionId(ctx); ok {
				// Recursively sync the fields of PartitionId
				toPartitionId.SyncFieldsDuringCreateOrUpdate(ctx, fromPartitionId)
				to.SetPartitionId(ctx, toPartitionId)
			}
		}
	}
}

func (to *ExchangeTokenRequest) SyncFieldsDuringRead(ctx context.Context, from ExchangeTokenRequest) {
	if !from.PartitionId.IsNull() && !from.PartitionId.IsUnknown() {
		if toPartitionId, ok := to.GetPartitionId(ctx); ok {
			if fromPartitionId, ok := from.GetPartitionId(ctx); ok {
				toPartitionId.SyncFieldsDuringRead(ctx, fromPartitionId)
				to.SetPartitionId(ctx, toPartitionId)
			}
		}
	}
}

func (m ExchangeTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["partition_id"] = attrs["partition_id"].SetRequired()
	attrs["scopes"] = attrs["scopes"].SetRequired()
	attrs["token_type"] = attrs["token_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExchangeTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partition_id": reflect.TypeOf(PartitionId{}),
		"scopes":       reflect.TypeOf(types.String{}),
		"token_type":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenRequest
// only implements ToObjectValue() and Type().
func (m ExchangeTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"partition_id": m.PartitionId,
			"scopes":       m.Scopes,
			"token_type":   m.TokenType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"partition_id": PartitionId{}.Type(ctx),
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_type": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPartitionId returns the value of the PartitionId field in ExchangeTokenRequest as
// a PartitionId value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenRequest) GetPartitionId(ctx context.Context) (PartitionId, bool) {
	var e PartitionId
	if m.PartitionId.IsNull() || m.PartitionId.IsUnknown() {
		return e, false
	}
	var v PartitionId
	d := m.PartitionId.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPartitionId sets the value of the PartitionId field in ExchangeTokenRequest.
func (m *ExchangeTokenRequest) SetPartitionId(ctx context.Context, v PartitionId) {
	vs := v.ToObjectValue(ctx)
	m.PartitionId = vs
}

// GetScopes returns the value of the Scopes field in ExchangeTokenRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenRequest) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeTokenRequest.
func (m *ExchangeTokenRequest) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenType returns the value of the TokenType field in ExchangeTokenRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenRequest) GetTokenType(ctx context.Context) ([]types.String, bool) {
	if m.TokenType.IsNull() || m.TokenType.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.TokenType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenType sets the value of the TokenType field in ExchangeTokenRequest.
func (m *ExchangeTokenRequest) SetTokenType(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_type"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TokenType = types.ListValueMust(t, vs)
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values types.List `tfsdk:"values"`
}

func (to *ExchangeTokenResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeTokenResponse) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *ExchangeTokenResponse) SyncFieldsDuringRead(ctx context.Context, from ExchangeTokenResponse) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m ExchangeTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["values"] = attrs["values"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExchangeTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(ExchangeToken{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenResponse
// only implements ToObjectValue() and Type().
func (m ExchangeTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"values": basetypes.ListType{
				ElemType: ExchangeToken{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in ExchangeTokenResponse as
// a slice of ExchangeToken values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenResponse) GetValues(ctx context.Context) ([]ExchangeToken, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []ExchangeToken
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in ExchangeTokenResponse.
func (m *ExchangeTokenResponse) SetValues(ctx context.Context, v []ExchangeToken) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	IpAccessList types.Object `tfsdk:"ip_access_list"`
}

func (to *FetchIpAccessListResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FetchIpAccessListResponse) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				// Recursively sync the fields of IpAccessList
				toIpAccessList.SyncFieldsDuringCreateOrUpdate(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (to *FetchIpAccessListResponse) SyncFieldsDuringRead(ctx context.Context, from FetchIpAccessListResponse) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (m FetchIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FetchIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FetchIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FetchIpAccessListResponse
// only implements ToObjectValue() and Type().
func (m FetchIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": m.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FetchIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": IpAccessListInfo{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in FetchIpAccessListResponse as
// a IpAccessListInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *FetchIpAccessListResponse) GetIpAccessList(ctx context.Context) (IpAccessListInfo, bool) {
	var e IpAccessListInfo
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v IpAccessListInfo
	d := m.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessList sets the value of the IpAccessList field in FetchIpAccessListResponse.
func (m *FetchIpAccessListResponse) SetIpAccessList(ctx context.Context, v IpAccessListInfo) {
	vs := v.ToObjectValue(ctx)
	m.IpAccessList = vs
}

type GenericWebhookConfig struct {
	// [Input-Only][Optional] Password for webhook.
	Password types.String `tfsdk:"password"`
	// [Output-Only] Whether password is set.
	PasswordSet types.Bool `tfsdk:"password_set"`
	// [Input-Only] URL for webhook.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
	// [Input-Only][Optional] Username for webhook.
	Username types.String `tfsdk:"username"`
	// [Output-Only] Whether username is set.
	UsernameSet types.Bool `tfsdk:"username_set"`
}

func (to *GenericWebhookConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenericWebhookConfig) {
}

func (to *GenericWebhookConfig) SyncFieldsDuringRead(ctx context.Context, from GenericWebhookConfig) {
}

func (m GenericWebhookConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["password"] = attrs["password"].SetOptional()
	attrs["password_set"] = attrs["password_set"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()
	attrs["url_set"] = attrs["url_set"].SetOptional()
	attrs["username"] = attrs["username"].SetOptional()
	attrs["username_set"] = attrs["username_set"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenericWebhookConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GenericWebhookConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenericWebhookConfig
// only implements ToObjectValue() and Type().
func (m GenericWebhookConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"password":     m.Password,
			"password_set": m.PasswordSet,
			"url":          m.Url,
			"url_set":      m.UrlSet,
			"username":     m.Username,
			"username_set": m.UsernameSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GenericWebhookConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"password":     types.StringType,
			"password_set": types.BoolType,
			"url":          types.StringType,
			"url_set":      types.BoolType,
			"username":     types.StringType,
			"username_set": types.BoolType,
		},
	}
}

type GetAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAccountIpAccessEnableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAccountIpAccessEnableRequest) {
}

func (to *GetAccountIpAccessEnableRequest) SyncFieldsDuringRead(ctx context.Context, from GetAccountIpAccessEnableRequest) {
}

func (m GetAccountIpAccessEnableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAccountIpAccessEnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessEnableRequest
// only implements ToObjectValue() and Type().
func (m GetAccountIpAccessEnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAccountIpAccessEnableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *GetAccountIpAccessListRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAccountIpAccessListRequest) {
}

func (to *GetAccountIpAccessListRequest) SyncFieldsDuringRead(ctx context.Context, from GetAccountIpAccessListRequest) {
}

func (m GetAccountIpAccessListRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAccountIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessListRequest
// only implements ToObjectValue() and Type().
func (m GetAccountIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAccountIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type GetAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAibiDashboardEmbeddingAccessPolicySettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (to *GetAibiDashboardEmbeddingAccessPolicySettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingAccessPolicySettingRequest
// only implements ToObjectValue() and Type().
func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (to *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingApprovedDomainsSettingRequest
// only implements ToObjectValue() and Type().
func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetAutomaticClusterUpdateSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAutomaticClusterUpdateSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAutomaticClusterUpdateSettingRequest) {
}

func (to *GetAutomaticClusterUpdateSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetAutomaticClusterUpdateSettingRequest) {
}

func (m GetAutomaticClusterUpdateSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAutomaticClusterUpdateSettingRequest
// only implements ToObjectValue() and Type().
func (m GetAutomaticClusterUpdateSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAutomaticClusterUpdateSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetComplianceSecurityProfileSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetComplianceSecurityProfileSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetComplianceSecurityProfileSettingRequest) {
}

func (to *GetComplianceSecurityProfileSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetComplianceSecurityProfileSettingRequest) {
}

func (m GetComplianceSecurityProfileSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetComplianceSecurityProfileSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetComplianceSecurityProfileSettingRequest
// only implements ToObjectValue() and Type().
func (m GetComplianceSecurityProfileSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetComplianceSecurityProfileSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetCspEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetCspEnablementAccountSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCspEnablementAccountSettingRequest) {
}

func (to *GetCspEnablementAccountSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetCspEnablementAccountSettingRequest) {
}

func (m GetCspEnablementAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetCspEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCspEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (m GetCspEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCspEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDashboardEmailSubscriptionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDashboardEmailSubscriptionsRequest) {
}

func (to *GetDashboardEmailSubscriptionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetDashboardEmailSubscriptionsRequest) {
}

func (m GetDashboardEmailSubscriptionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDashboardEmailSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDashboardEmailSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDashboardEmailSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (m GetDashboardEmailSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDashboardEmailSubscriptionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDefaultNamespaceSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDefaultNamespaceSettingRequest) {
}

func (to *GetDefaultNamespaceSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetDefaultNamespaceSettingRequest) {
}

func (m GetDefaultNamespaceSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultNamespaceSettingRequest
// only implements ToObjectValue() and Type().
func (m GetDefaultNamespaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDefaultNamespaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDefaultWarehouseIdRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDefaultWarehouseIdRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDefaultWarehouseIdRequest) {
}

func (to *GetDefaultWarehouseIdRequest) SyncFieldsDuringRead(ctx context.Context, from GetDefaultWarehouseIdRequest) {
}

func (m GetDefaultWarehouseIdRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultWarehouseIdRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDefaultWarehouseIdRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultWarehouseIdRequest
// only implements ToObjectValue() and Type().
func (m GetDefaultWarehouseIdRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDefaultWarehouseIdRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDisableLegacyAccessRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDisableLegacyAccessRequest) {
}

func (to *GetDisableLegacyAccessRequest) SyncFieldsDuringRead(ctx context.Context, from GetDisableLegacyAccessRequest) {
}

func (m GetDisableLegacyAccessRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyAccessRequest
// only implements ToObjectValue() and Type().
func (m GetDisableLegacyAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDisableLegacyAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDisableLegacyDbfsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDisableLegacyDbfsRequest) {
}

func (to *GetDisableLegacyDbfsRequest) SyncFieldsDuringRead(ctx context.Context, from GetDisableLegacyDbfsRequest) {
}

func (m GetDisableLegacyDbfsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyDbfsRequest
// only implements ToObjectValue() and Type().
func (m GetDisableLegacyDbfsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDisableLegacyDbfsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDisableLegacyFeaturesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDisableLegacyFeaturesRequest) {
}

func (to *GetDisableLegacyFeaturesRequest) SyncFieldsDuringRead(ctx context.Context, from GetDisableLegacyFeaturesRequest) {
}

func (m GetDisableLegacyFeaturesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyFeaturesRequest
// only implements ToObjectValue() and Type().
func (m GetDisableLegacyFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDisableLegacyFeaturesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetEnableExportNotebookRequest struct {
}

func (to *GetEnableExportNotebookRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnableExportNotebookRequest) {
}

func (to *GetEnableExportNotebookRequest) SyncFieldsDuringRead(ctx context.Context, from GetEnableExportNotebookRequest) {
}

func (m GetEnableExportNotebookRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnableExportNotebookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnableExportNotebookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnableExportNotebookRequest
// only implements ToObjectValue() and Type().
func (m GetEnableExportNotebookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnableExportNotebookRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetEnableNotebookTableClipboardRequest struct {
}

func (to *GetEnableNotebookTableClipboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnableNotebookTableClipboardRequest) {
}

func (to *GetEnableNotebookTableClipboardRequest) SyncFieldsDuringRead(ctx context.Context, from GetEnableNotebookTableClipboardRequest) {
}

func (m GetEnableNotebookTableClipboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnableNotebookTableClipboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnableNotebookTableClipboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnableNotebookTableClipboardRequest
// only implements ToObjectValue() and Type().
func (m GetEnableNotebookTableClipboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnableNotebookTableClipboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetEnableResultsDownloadingRequest struct {
}

func (to *GetEnableResultsDownloadingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnableResultsDownloadingRequest) {
}

func (to *GetEnableResultsDownloadingRequest) SyncFieldsDuringRead(ctx context.Context, from GetEnableResultsDownloadingRequest) {
}

func (m GetEnableResultsDownloadingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnableResultsDownloadingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnableResultsDownloadingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnableResultsDownloadingRequest
// only implements ToObjectValue() and Type().
func (m GetEnableResultsDownloadingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnableResultsDownloadingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetEnhancedSecurityMonitoringSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetEnhancedSecurityMonitoringSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnhancedSecurityMonitoringSettingRequest) {
}

func (to *GetEnhancedSecurityMonitoringSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetEnhancedSecurityMonitoringSettingRequest) {
}

func (m GetEnhancedSecurityMonitoringSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnhancedSecurityMonitoringSettingRequest
// only implements ToObjectValue() and Type().
func (m GetEnhancedSecurityMonitoringSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnhancedSecurityMonitoringSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetEsmEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetEsmEnablementAccountSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEsmEnablementAccountSettingRequest) {
}

func (to *GetEsmEnablementAccountSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetEsmEnablementAccountSettingRequest) {
}

func (m GetEsmEnablementAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEsmEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEsmEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (m GetEsmEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEsmEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *GetIpAccessListRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIpAccessListRequest) {
}

func (to *GetIpAccessListRequest) SyncFieldsDuringRead(ctx context.Context, from GetIpAccessListRequest) {
}

func (m GetIpAccessListRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListRequest
// only implements ToObjectValue() and Type().
func (m GetIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type GetIpAccessListResponse struct {
	IpAccessList types.Object `tfsdk:"ip_access_list"`
}

func (to *GetIpAccessListResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIpAccessListResponse) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				// Recursively sync the fields of IpAccessList
				toIpAccessList.SyncFieldsDuringCreateOrUpdate(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (to *GetIpAccessListResponse) SyncFieldsDuringRead(ctx context.Context, from GetIpAccessListResponse) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (m GetIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListResponse
// only implements ToObjectValue() and Type().
func (m GetIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": m.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": IpAccessListInfo{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in GetIpAccessListResponse as
// a IpAccessListInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetIpAccessListResponse) GetIpAccessList(ctx context.Context) (IpAccessListInfo, bool) {
	var e IpAccessListInfo
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v IpAccessListInfo
	d := m.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessList sets the value of the IpAccessList field in GetIpAccessListResponse.
func (m *GetIpAccessListResponse) SetIpAccessList(ctx context.Context, v IpAccessListInfo) {
	vs := v.ToObjectValue(ctx)
	m.IpAccessList = vs
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (to *GetIpAccessListsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIpAccessListsResponse) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (to *GetIpAccessListsResponse) SyncFieldsDuringRead(ctx context.Context, from GetIpAccessListsResponse) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (m GetIpAccessListsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_lists"] = attrs["ip_access_lists"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIpAccessListsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListsResponse
// only implements ToObjectValue() and Type().
func (m GetIpAccessListsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": m.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIpAccessListsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_lists": basetypes.ListType{
				ElemType: IpAccessListInfo{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessLists returns the value of the IpAccessLists field in GetIpAccessListsResponse as
// a slice of IpAccessListInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetIpAccessListsResponse) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo, bool) {
	if m.IpAccessLists.IsNull() || m.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo
	d := m.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in GetIpAccessListsResponse.
func (m *GetIpAccessListsResponse) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAccessLists = types.ListValueMust(t, vs)
}

type GetLlmProxyPartnerPoweredAccountRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetLlmProxyPartnerPoweredAccountRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLlmProxyPartnerPoweredAccountRequest) {
}

func (to *GetLlmProxyPartnerPoweredAccountRequest) SyncFieldsDuringRead(ctx context.Context, from GetLlmProxyPartnerPoweredAccountRequest) {
}

func (m GetLlmProxyPartnerPoweredAccountRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLlmProxyPartnerPoweredAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLlmProxyPartnerPoweredAccountRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredAccountRequest
// only implements ToObjectValue() and Type().
func (m GetLlmProxyPartnerPoweredAccountRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLlmProxyPartnerPoweredAccountRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetLlmProxyPartnerPoweredEnforceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetLlmProxyPartnerPoweredEnforceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLlmProxyPartnerPoweredEnforceRequest) {
}

func (to *GetLlmProxyPartnerPoweredEnforceRequest) SyncFieldsDuringRead(ctx context.Context, from GetLlmProxyPartnerPoweredEnforceRequest) {
}

func (m GetLlmProxyPartnerPoweredEnforceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLlmProxyPartnerPoweredEnforceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLlmProxyPartnerPoweredEnforceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredEnforceRequest
// only implements ToObjectValue() and Type().
func (m GetLlmProxyPartnerPoweredEnforceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLlmProxyPartnerPoweredEnforceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetLlmProxyPartnerPoweredWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLlmProxyPartnerPoweredWorkspaceRequest) {
}

func (to *GetLlmProxyPartnerPoweredWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from GetLlmProxyPartnerPoweredWorkspaceRequest) {
}

func (m GetLlmProxyPartnerPoweredWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLlmProxyPartnerPoweredWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetLlmProxyPartnerPoweredWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m GetLlmProxyPartnerPoweredWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLlmProxyPartnerPoweredWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

func (to *GetNetworkConnectivityConfigurationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkConnectivityConfigurationRequest) {
}

func (to *GetNetworkConnectivityConfigurationRequest) SyncFieldsDuringRead(ctx context.Context, from GetNetworkConnectivityConfigurationRequest) {
}

func (m GetNetworkConnectivityConfigurationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetNetworkConnectivityConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkConnectivityConfigurationRequest
// only implements ToObjectValue() and Type().
func (m GetNetworkConnectivityConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNetworkConnectivityConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type GetNetworkPolicyRequest struct {
	// The unique identifier of the network policy to retrieve.
	NetworkPolicyId types.String `tfsdk:"-"`
}

func (to *GetNetworkPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkPolicyRequest) {
}

func (to *GetNetworkPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from GetNetworkPolicyRequest) {
}

func (m GetNetworkPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_policy_id"] = attrs["network_policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (m GetNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
		},
	}
}

type GetNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetNotificationDestinationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNotificationDestinationRequest) {
}

func (to *GetNotificationDestinationRequest) SyncFieldsDuringRead(ctx context.Context, from GetNotificationDestinationRequest) {
}

func (m GetNotificationDestinationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (m GetNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetPersonalComputeSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPersonalComputeSettingRequest) {
}

func (to *GetPersonalComputeSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetPersonalComputeSettingRequest) {
}

func (m GetPersonalComputeSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalComputeSettingRequest
// only implements ToObjectValue() and Type().
func (m GetPersonalComputeSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPersonalComputeSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

func (to *GetPrivateEndpointRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPrivateEndpointRuleRequest) {
}

func (to *GetPrivateEndpointRuleRequest) SyncFieldsDuringRead(ctx context.Context, from GetPrivateEndpointRuleRequest) {
}

func (m GetPrivateEndpointRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()
	attrs["private_endpoint_rule_id"] = attrs["private_endpoint_rule_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetPrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (m GetPrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       m.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetRestrictWorkspaceAdminsSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRestrictWorkspaceAdminsSettingRequest) {
}

func (to *GetRestrictWorkspaceAdminsSettingRequest) SyncFieldsDuringRead(ctx context.Context, from GetRestrictWorkspaceAdminsSettingRequest) {
}

func (m GetRestrictWorkspaceAdminsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRestrictWorkspaceAdminsSettingRequest
// only implements ToObjectValue() and Type().
func (m GetRestrictWorkspaceAdminsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRestrictWorkspaceAdminsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetSqlResultsDownloadRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSqlResultsDownloadRequest) {
}

func (to *GetSqlResultsDownloadRequest) SyncFieldsDuringRead(ctx context.Context, from GetSqlResultsDownloadRequest) {
}

func (m GetSqlResultsDownloadRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSqlResultsDownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetSqlResultsDownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSqlResultsDownloadRequest
// only implements ToObjectValue() and Type().
func (m GetSqlResultsDownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSqlResultsDownloadRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetStatusRequest struct {
	Keys types.String `tfsdk:"-"`
}

func (to *GetStatusRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatusRequest) {
}

func (to *GetStatusRequest) SyncFieldsDuringRead(ctx context.Context, from GetStatusRequest) {
}

func (m GetStatusRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["keys"] = attrs["keys"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest
// only implements ToObjectValue() and Type().
func (m GetStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"keys": m.Keys,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"keys": types.StringType,
		},
	}
}

type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-"`
}

func (to *GetTokenManagementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenManagementRequest) {
}

func (to *GetTokenManagementRequest) SyncFieldsDuringRead(ctx context.Context, from GetTokenManagementRequest) {
}

func (m GetTokenManagementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_id"] = attrs["token_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenManagementRequest
// only implements ToObjectValue() and Type().
func (m GetTokenManagementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenManagementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type GetTokenPermissionLevelsRequest struct {
}

func (to *GetTokenPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenPermissionLevelsRequest) {
}

func (to *GetTokenPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetTokenPermissionLevelsRequest) {
}

func (m GetTokenPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (m GetTokenPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetTokenPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetTokenPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetTokenPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetTokenPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(TokenPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (m GetTokenPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: TokenPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetTokenPermissionLevelsResponse as
// a slice of TokenPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetTokenPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]TokenPermissionsDescription, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []TokenPermissionsDescription
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetTokenPermissionLevelsResponse.
func (m *GetTokenPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []TokenPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetTokenPermissionsRequest struct {
}

func (to *GetTokenPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenPermissionsRequest) {
}

func (to *GetTokenPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetTokenPermissionsRequest) {
}

func (m GetTokenPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionsRequest
// only implements ToObjectValue() and Type().
func (m GetTokenPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo types.Object `tfsdk:"token_info"`
}

func (to *GetTokenResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenResponse) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				// Recursively sync the fields of TokenInfo
				toTokenInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (to *GetTokenResponse) SyncFieldsDuringRead(ctx context.Context, from GetTokenResponse) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				toTokenInfo.SyncFieldsDuringRead(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (m GetTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_info"] = attrs["token_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenResponse
// only implements ToObjectValue() and Type().
func (m GetTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info": m.TokenInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": TokenInfo{}.Type(ctx),
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in GetTokenResponse as
// a TokenInfo value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetTokenResponse) GetTokenInfo(ctx context.Context) (TokenInfo, bool) {
	var e TokenInfo
	if m.TokenInfo.IsNull() || m.TokenInfo.IsUnknown() {
		return e, false
	}
	var v TokenInfo
	d := m.TokenInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfo sets the value of the TokenInfo field in GetTokenResponse.
func (m *GetTokenResponse) SetTokenInfo(ctx context.Context, v TokenInfo) {
	vs := v.ToObjectValue(ctx)
	m.TokenInfo = vs
}

type GetWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceNetworkOptionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceNetworkOptionRequest) {
}

func (to *GetWorkspaceNetworkOptionRequest) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceNetworkOptionRequest) {
}

func (m GetWorkspaceNetworkOptionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceNetworkOptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetWorkspaceNetworkOptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceNetworkOptionRequest
// only implements ToObjectValue() and Type().
func (m GetWorkspaceNetworkOptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceNetworkOptionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount types.Int64 `tfsdk:"address_count"`
	// Creation timestamp in milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at"`
	// User ID of the user who created this list.
	CreatedBy types.Int64 `tfsdk:"created_by"`
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId types.String `tfsdk:"list_id"`

	ListType types.String `tfsdk:"list_type"`
	// Update timestamp in milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// User ID of the user who updated this list.
	UpdatedBy types.Int64 `tfsdk:"updated_by"`
}

func (to *IpAccessListInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IpAccessListInfo) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *IpAccessListInfo) SyncFieldsDuringRead(ctx context.Context, from IpAccessListInfo) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m IpAccessListInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["address_count"] = attrs["address_count"].SetOptional()
	attrs["created_at"] = attrs["created_at"].SetOptional()
	attrs["created_by"] = attrs["created_by"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetOptional()
	attrs["list_id"] = attrs["list_id"].SetOptional()
	attrs["list_type"] = attrs["list_type"].SetOptional()
	attrs["updated_at"] = attrs["updated_at"].SetOptional()
	attrs["updated_by"] = attrs["updated_by"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in IpAccessListInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m IpAccessListInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessListInfo
// only implements ToObjectValue() and Type().
func (m IpAccessListInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"address_count": m.AddressCount,
			"created_at":    m.CreatedAt,
			"created_by":    m.CreatedBy,
			"enabled":       m.Enabled,
			"ip_addresses":  m.IpAddresses,
			"label":         m.Label,
			"list_id":       m.ListId,
			"list_type":     m.ListType,
			"updated_at":    m.UpdatedAt,
			"updated_by":    m.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m IpAccessListInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"address_count": types.Int64Type,
			"created_at":    types.Int64Type,
			"created_by":    types.Int64Type,
			"enabled":       types.BoolType,
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":      types.StringType,
			"list_id":    types.StringType,
			"list_type":  types.StringType,
			"updated_at": types.Int64Type,
			"updated_by": types.Int64Type,
		},
	}
}

// GetIpAddresses returns the value of the IpAddresses field in IpAccessListInfo as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *IpAccessListInfo) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if m.IpAddresses.IsNull() || m.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in IpAccessListInfo.
func (m *IpAccessListInfo) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (to *ListIpAccessListResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIpAccessListResponse) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (to *ListIpAccessListResponse) SyncFieldsDuringRead(ctx context.Context, from ListIpAccessListResponse) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (m ListIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_lists"] = attrs["ip_access_lists"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIpAccessListResponse
// only implements ToObjectValue() and Type().
func (m ListIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": m.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_lists": basetypes.ListType{
				ElemType: IpAccessListInfo{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessLists returns the value of the IpAccessLists field in ListIpAccessListResponse as
// a slice of IpAccessListInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListIpAccessListResponse) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo, bool) {
	if m.IpAccessLists.IsNull() || m.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo
	d := m.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in ListIpAccessListResponse.
func (m *ListIpAccessListResponse) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAccessLists = types.ListValueMust(t, vs)
}

type ListIpAccessLists struct {
}

func (to *ListIpAccessLists) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIpAccessLists) {
}

func (to *ListIpAccessLists) SyncFieldsDuringRead(ctx context.Context, from ListIpAccessLists) {
}

func (m ListIpAccessLists) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIpAccessLists.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIpAccessLists) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIpAccessLists
// only implements ToObjectValue() and Type().
func (m ListIpAccessLists) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListIpAccessLists) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListNetworkConnectivityConfigurationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkConnectivityConfigurationsRequest) {
}

func (to *ListNetworkConnectivityConfigurationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListNetworkConnectivityConfigurationsRequest) {
}

func (m ListNetworkConnectivityConfigurationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkConnectivityConfigurationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNetworkConnectivityConfigurationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsRequest
// only implements ToObjectValue() and Type().
func (m ListNetworkConnectivityConfigurationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkConnectivityConfigurationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

// The network connectivity configuration list was successfully retrieved.
type ListNetworkConnectivityConfigurationsResponse struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListNetworkConnectivityConfigurationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkConnectivityConfigurationsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListNetworkConnectivityConfigurationsResponse) SyncFieldsDuringRead(ctx context.Context, from ListNetworkConnectivityConfigurationsResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListNetworkConnectivityConfigurationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkConnectivityConfigurationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNetworkConnectivityConfigurationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NetworkConnectivityConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsResponse
// only implements ToObjectValue() and Type().
func (m ListNetworkConnectivityConfigurationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkConnectivityConfigurationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NetworkConnectivityConfiguration{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListNetworkConnectivityConfigurationsResponse as
// a slice of NetworkConnectivityConfiguration values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNetworkConnectivityConfigurationsResponse) GetItems(ctx context.Context) ([]NetworkConnectivityConfiguration, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []NetworkConnectivityConfiguration
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkConnectivityConfigurationsResponse.
func (m *ListNetworkConnectivityConfigurationsResponse) SetItems(ctx context.Context, v []NetworkConnectivityConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListNetworkPoliciesRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListNetworkPoliciesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkPoliciesRequest) {
}

func (to *ListNetworkPoliciesRequest) SyncFieldsDuringRead(ctx context.Context, from ListNetworkPoliciesRequest) {
}

func (m ListNetworkPoliciesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNetworkPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkPoliciesRequest
// only implements ToObjectValue() and Type().
func (m ListNetworkPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkPoliciesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListNetworkPoliciesResponse struct {
	// List of network policies.
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListNetworkPoliciesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkPoliciesResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListNetworkPoliciesResponse) SyncFieldsDuringRead(ctx context.Context, from ListNetworkPoliciesResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListNetworkPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkPoliciesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNetworkPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AccountNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkPoliciesResponse
// only implements ToObjectValue() and Type().
func (m ListNetworkPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkPoliciesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: AccountNetworkPolicy{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListNetworkPoliciesResponse as
// a slice of AccountNetworkPolicy values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNetworkPoliciesResponse) GetItems(ctx context.Context) ([]AccountNetworkPolicy, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []AccountNetworkPolicy
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkPoliciesResponse.
func (m *ListNetworkPoliciesResponse) SetItems(ctx context.Context, v []AccountNetworkPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListNotificationDestinationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListNotificationDestinationsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNotificationDestinationsRequest) {
}

func (to *ListNotificationDestinationsRequest) SyncFieldsDuringRead(ctx context.Context, from ListNotificationDestinationsRequest) {
}

func (m ListNotificationDestinationsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_token"] = attrs["page_token"].SetOptional()
	attrs["page_size"] = attrs["page_size"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNotificationDestinationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsRequest
// only implements ToObjectValue() and Type().
func (m ListNotificationDestinationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNotificationDestinationsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (to *ListNotificationDestinationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNotificationDestinationsResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListNotificationDestinationsResponse) SyncFieldsDuringRead(ctx context.Context, from ListNotificationDestinationsResponse) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m ListNotificationDestinationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()
	attrs["results"] = attrs["results"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNotificationDestinationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListNotificationDestinationsResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResponse
// only implements ToObjectValue() and Type().
func (m ListNotificationDestinationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"results":         m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNotificationDestinationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ListNotificationDestinationsResult{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListNotificationDestinationsResponse as
// a slice of ListNotificationDestinationsResult values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNotificationDestinationsResponse) GetResults(ctx context.Context) ([]ListNotificationDestinationsResult, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []ListNotificationDestinationsResult
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListNotificationDestinationsResponse.
func (m *ListNotificationDestinationsResponse) SetResults(ctx context.Context, v []ListNotificationDestinationsResult) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
}

type ListNotificationDestinationsResult struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id"`
}

func (to *ListNotificationDestinationsResult) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNotificationDestinationsResult) {
}

func (to *ListNotificationDestinationsResult) SyncFieldsDuringRead(ctx context.Context, from ListNotificationDestinationsResult) {
}

func (m ListNotificationDestinationsResult) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["destination_type"] = attrs["destination_type"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsResult.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListNotificationDestinationsResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResult
// only implements ToObjectValue() and Type().
func (m ListNotificationDestinationsResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_type": m.DestinationType,
			"display_name":     m.DisplayName,
			"id":               m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNotificationDestinationsResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
		},
	}
}

type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListPrivateEndpointRulesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateEndpointRulesRequest) {
}

func (to *ListPrivateEndpointRulesRequest) SyncFieldsDuringRead(ctx context.Context, from ListPrivateEndpointRulesRequest) {
}

func (m ListPrivateEndpointRulesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()
	attrs["page_token"] = attrs["page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateEndpointRulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPrivateEndpointRulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesRequest
// only implements ToObjectValue() and Type().
func (m ListPrivateEndpointRulesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"page_token":                     m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateEndpointRulesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"page_token":                     types.StringType,
		},
	}
}

// The private endpoint rule list was successfully retrieved.
type ListPrivateEndpointRulesResponse struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListPrivateEndpointRulesResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateEndpointRulesResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListPrivateEndpointRulesResponse) SyncFieldsDuringRead(ctx context.Context, from ListPrivateEndpointRulesResponse) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListPrivateEndpointRulesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateEndpointRulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPrivateEndpointRulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NccPrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesResponse
// only implements ToObjectValue() and Type().
func (m ListPrivateEndpointRulesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateEndpointRulesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NccPrivateEndpointRule{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListPrivateEndpointRulesResponse as
// a slice of NccPrivateEndpointRule values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPrivateEndpointRulesResponse) GetItems(ctx context.Context) ([]NccPrivateEndpointRule, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []NccPrivateEndpointRule
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListPrivateEndpointRulesResponse.
func (m *ListPrivateEndpointRulesResponse) SetItems(ctx context.Context, v []NccPrivateEndpointRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (to *ListPublicTokensResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPublicTokensResponse) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (to *ListPublicTokensResponse) SyncFieldsDuringRead(ctx context.Context, from ListPublicTokensResponse) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (m ListPublicTokensResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_infos"] = attrs["token_infos"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPublicTokensResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListPublicTokensResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(PublicTokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublicTokensResponse
// only implements ToObjectValue() and Type().
func (m ListPublicTokensResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": m.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPublicTokensResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_infos": basetypes.ListType{
				ElemType: PublicTokenInfo{}.Type(ctx),
			},
		},
	}
}

// GetTokenInfos returns the value of the TokenInfos field in ListPublicTokensResponse as
// a slice of PublicTokenInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPublicTokensResponse) GetTokenInfos(ctx context.Context) ([]PublicTokenInfo, bool) {
	if m.TokenInfos.IsNull() || m.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []PublicTokenInfo
	d := m.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListPublicTokensResponse.
func (m *ListPublicTokensResponse) SetTokenInfos(ctx context.Context, v []PublicTokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TokenInfos = types.ListValueMust(t, vs)
}

type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"-"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"-"`
}

func (to *ListTokenManagementRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTokenManagementRequest) {
}

func (to *ListTokenManagementRequest) SyncFieldsDuringRead(ctx context.Context, from ListTokenManagementRequest) {
}

func (m ListTokenManagementRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["created_by_id"] = attrs["created_by_id"].SetOptional()
	attrs["created_by_username"] = attrs["created_by_username"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokenManagementRequest
// only implements ToObjectValue() and Type().
func (m ListTokenManagementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_by_id":       m.CreatedById,
			"created_by_username": m.CreatedByUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTokenManagementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
		},
	}
}

type ListTokens struct {
}

func (to *ListTokens) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTokens) {
}

func (to *ListTokens) SyncFieldsDuringRead(ctx context.Context, from ListTokens) {
}

func (m ListTokens) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokens.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTokens) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokens
// only implements ToObjectValue() and Type().
func (m ListTokens) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListTokens) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (to *ListTokensResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTokensResponse) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (to *ListTokensResponse) SyncFieldsDuringRead(ctx context.Context, from ListTokensResponse) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (m ListTokensResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_infos"] = attrs["token_infos"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokensResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTokensResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokensResponse
// only implements ToObjectValue() and Type().
func (m ListTokensResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": m.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTokensResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_infos": basetypes.ListType{
				ElemType: TokenInfo{}.Type(ctx),
			},
		},
	}
}

// GetTokenInfos returns the value of the TokenInfos field in ListTokensResponse as
// a slice of TokenInfo values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTokensResponse) GetTokenInfos(ctx context.Context) ([]TokenInfo, bool) {
	if m.TokenInfos.IsNull() || m.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo
	d := m.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListTokensResponse.
func (m *ListTokensResponse) SetTokenInfos(ctx context.Context, v []TokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TokenInfos = types.ListValueMust(t, vs)
}

type LlmProxyPartnerPoweredAccount struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *LlmProxyPartnerPoweredAccount) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LlmProxyPartnerPoweredAccount) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *LlmProxyPartnerPoweredAccount) SyncFieldsDuringRead(ctx context.Context, from LlmProxyPartnerPoweredAccount) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m LlmProxyPartnerPoweredAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LlmProxyPartnerPoweredAccount.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LlmProxyPartnerPoweredAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredAccount
// only implements ToObjectValue() and Type().
func (m LlmProxyPartnerPoweredAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LlmProxyPartnerPoweredAccount) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in LlmProxyPartnerPoweredAccount as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *LlmProxyPartnerPoweredAccount) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredAccount.
func (m *LlmProxyPartnerPoweredAccount) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

type LlmProxyPartnerPoweredEnforce struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *LlmProxyPartnerPoweredEnforce) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LlmProxyPartnerPoweredEnforce) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *LlmProxyPartnerPoweredEnforce) SyncFieldsDuringRead(ctx context.Context, from LlmProxyPartnerPoweredEnforce) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m LlmProxyPartnerPoweredEnforce) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LlmProxyPartnerPoweredEnforce.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LlmProxyPartnerPoweredEnforce) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredEnforce
// only implements ToObjectValue() and Type().
func (m LlmProxyPartnerPoweredEnforce) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LlmProxyPartnerPoweredEnforce) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in LlmProxyPartnerPoweredEnforce as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *LlmProxyPartnerPoweredEnforce) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredEnforce.
func (m *LlmProxyPartnerPoweredEnforce) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

type LlmProxyPartnerPoweredWorkspace struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *LlmProxyPartnerPoweredWorkspace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LlmProxyPartnerPoweredWorkspace) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *LlmProxyPartnerPoweredWorkspace) SyncFieldsDuringRead(ctx context.Context, from LlmProxyPartnerPoweredWorkspace) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m LlmProxyPartnerPoweredWorkspace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LlmProxyPartnerPoweredWorkspace.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LlmProxyPartnerPoweredWorkspace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredWorkspace
// only implements ToObjectValue() and Type().
func (m LlmProxyPartnerPoweredWorkspace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LlmProxyPartnerPoweredWorkspace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in LlmProxyPartnerPoweredWorkspace as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *LlmProxyPartnerPoweredWorkspace) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredWorkspace.
func (m *LlmProxyPartnerPoweredWorkspace) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] App ID for Microsoft Teams App.
	AppId types.String `tfsdk:"app_id"`
	// [Output-Only] Whether App ID is set.
	AppIdSet types.Bool `tfsdk:"app_id_set"`
	// [Input-Only] Secret for Microsoft Teams App authentication.
	AuthSecret types.String `tfsdk:"auth_secret"`
	// [Output-Only] Whether secret is set.
	AuthSecretSet types.Bool `tfsdk:"auth_secret_set"`
	// [Input-Only] Channel URL for Microsoft Teams App.
	ChannelUrl types.String `tfsdk:"channel_url"`
	// [Output-Only] Whether Channel URL is set.
	ChannelUrlSet types.Bool `tfsdk:"channel_url_set"`
	// [Input-Only] Tenant ID for Microsoft Teams App.
	TenantId types.String `tfsdk:"tenant_id"`
	// [Output-Only] Whether Tenant ID is set.
	TenantIdSet types.Bool `tfsdk:"tenant_id_set"`
	// [Input-Only] URL for Microsoft Teams webhook.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
}

func (to *MicrosoftTeamsConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MicrosoftTeamsConfig) {
}

func (to *MicrosoftTeamsConfig) SyncFieldsDuringRead(ctx context.Context, from MicrosoftTeamsConfig) {
}

func (m MicrosoftTeamsConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["app_id"] = attrs["app_id"].SetOptional()
	attrs["app_id_set"] = attrs["app_id_set"].SetOptional()
	attrs["auth_secret"] = attrs["auth_secret"].SetOptional()
	attrs["auth_secret_set"] = attrs["auth_secret_set"].SetOptional()
	attrs["channel_url"] = attrs["channel_url"].SetOptional()
	attrs["channel_url_set"] = attrs["channel_url_set"].SetOptional()
	attrs["tenant_id"] = attrs["tenant_id"].SetOptional()
	attrs["tenant_id_set"] = attrs["tenant_id_set"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()
	attrs["url_set"] = attrs["url_set"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MicrosoftTeamsConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m MicrosoftTeamsConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MicrosoftTeamsConfig
// only implements ToObjectValue() and Type().
func (m MicrosoftTeamsConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"app_id":          m.AppId,
			"app_id_set":      m.AppIdSet,
			"auth_secret":     m.AuthSecret,
			"auth_secret_set": m.AuthSecretSet,
			"channel_url":     m.ChannelUrl,
			"channel_url_set": m.ChannelUrlSet,
			"tenant_id":       m.TenantId,
			"tenant_id_set":   m.TenantIdSet,
			"url":             m.Url,
			"url_set":         m.UrlSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (m MicrosoftTeamsConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"app_id":          types.StringType,
			"app_id_set":      types.BoolType,
			"auth_secret":     types.StringType,
			"auth_secret_set": types.BoolType,
			"channel_url":     types.StringType,
			"channel_url_set": types.BoolType,
			"tenant_id":       types.StringType,
			"tenant_id_set":   types.BoolType,
			"url":             types.StringType,
			"url_set":         types.BoolType,
		},
	}
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks types.List `tfsdk:"cidr_blocks"`
}

func (to *NccAwsStableIpRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccAwsStableIpRule) {
	if !from.CidrBlocks.IsNull() && !from.CidrBlocks.IsUnknown() && to.CidrBlocks.IsNull() && len(from.CidrBlocks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CidrBlocks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CidrBlocks = from.CidrBlocks
	}
}

func (to *NccAwsStableIpRule) SyncFieldsDuringRead(ctx context.Context, from NccAwsStableIpRule) {
	if !from.CidrBlocks.IsNull() && !from.CidrBlocks.IsUnknown() && to.CidrBlocks.IsNull() && len(from.CidrBlocks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CidrBlocks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CidrBlocks = from.CidrBlocks
	}
}

func (m NccAwsStableIpRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cidr_blocks"] = attrs["cidr_blocks"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAwsStableIpRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccAwsStableIpRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cidr_blocks": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAwsStableIpRule
// only implements ToObjectValue() and Type().
func (m NccAwsStableIpRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cidr_blocks": m.CidrBlocks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccAwsStableIpRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cidr_blocks": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetCidrBlocks returns the value of the CidrBlocks field in NccAwsStableIpRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAwsStableIpRule) GetCidrBlocks(ctx context.Context) ([]types.String, bool) {
	if m.CidrBlocks.IsNull() || m.CidrBlocks.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.CidrBlocks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCidrBlocks sets the value of the CidrBlocks field in NccAwsStableIpRule.
func (m *NccAwsStableIpRule) SetCidrBlocks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cidr_blocks"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.CidrBlocks = types.ListValueMust(t, vs)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type NccAzurePrivateEndpointRule struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the Azure portal
	// before they take effect. The possible values are: - INIT: (deprecated)
	// The endpoint has been created and pending approval. - PENDING: The
	// endpoint has been created and pending approval. - ESTABLISHED: The
	// endpoint has been approved and is ready to use in your serverless compute
	// resources. - REJECTED: Connection was rejected by the private link
	// resource owner. - DISCONNECTED: Connection was removed by the private
	// link resource owner, the private endpoint becomes informative and should
	// be deleted for clean-up. - EXPIRED: If the endpoint was created but not
	// approved in 14 days, it will be EXPIRED.
	ConnectionState types.String `tfsdk:"connection_state"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at"`
	// Not used by customer-managed private endpoint services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames types.List `tfsdk:"domain_names"`
	// The name of the Azure private endpoint resource.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// Only used by private endpoints to Azure first-party services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId types.String `tfsdk:"group_id"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id"`
	// The ID of a private endpoint rule.
	RuleId types.String `tfsdk:"rule_id"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time"`
}

func (to *NccAzurePrivateEndpointRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccAzurePrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
}

func (to *NccAzurePrivateEndpointRule) SyncFieldsDuringRead(ctx context.Context, from NccAzurePrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
}

func (m NccAzurePrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_state"] = attrs["connection_state"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["deactivated"] = attrs["deactivated"].SetOptional()
	attrs["deactivated_at"] = attrs["deactivated_at"].SetOptional()
	attrs["domain_names"] = attrs["domain_names"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["group_id"] = attrs["group_id"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["rule_id"] = attrs["rule_id"].SetOptional()
	attrs["updated_time"] = attrs["updated_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAzurePrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccAzurePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzurePrivateEndpointRule
// only implements ToObjectValue() and Type().
func (m NccAzurePrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_state":               m.ConnectionState,
			"creation_time":                  m.CreationTime,
			"deactivated":                    m.Deactivated,
			"deactivated_at":                 m.DeactivatedAt,
			"domain_names":                   m.DomainNames,
			"endpoint_name":                  m.EndpointName,
			"group_id":                       m.GroupId,
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"resource_id":                    m.ResourceId,
			"rule_id":                        m.RuleId,
			"updated_time":                   m.UpdatedTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccAzurePrivateEndpointRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_state": types.StringType,
			"creation_time":    types.Int64Type,
			"deactivated":      types.BoolType,
			"deactivated_at":   types.Int64Type,
			"domain_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"endpoint_name":                  types.StringType,
			"group_id":                       types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_id":                    types.StringType,
			"rule_id":                        types.StringType,
			"updated_time":                   types.Int64Type,
		},
	}
}

// GetDomainNames returns the value of the DomainNames field in NccAzurePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAzurePrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if m.DomainNames.IsNull() || m.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in NccAzurePrivateEndpointRule.
func (m *NccAzurePrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets types.List `tfsdk:"subnets"`
	// The Azure region in which this service endpoint rule applies..
	TargetRegion types.String `tfsdk:"target_region"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices types.List `tfsdk:"target_services"`
}

func (to *NccAzureServiceEndpointRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccAzureServiceEndpointRule) {
	if !from.Subnets.IsNull() && !from.Subnets.IsUnknown() && to.Subnets.IsNull() && len(from.Subnets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subnets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subnets = from.Subnets
	}
	if !from.TargetServices.IsNull() && !from.TargetServices.IsUnknown() && to.TargetServices.IsNull() && len(from.TargetServices.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TargetServices, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TargetServices = from.TargetServices
	}
}

func (to *NccAzureServiceEndpointRule) SyncFieldsDuringRead(ctx context.Context, from NccAzureServiceEndpointRule) {
	if !from.Subnets.IsNull() && !from.Subnets.IsUnknown() && to.Subnets.IsNull() && len(from.Subnets.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Subnets, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Subnets = from.Subnets
	}
	if !from.TargetServices.IsNull() && !from.TargetServices.IsUnknown() && to.TargetServices.IsNull() && len(from.TargetServices.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TargetServices, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TargetServices = from.TargetServices
	}
}

func (m NccAzureServiceEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["subnets"] = attrs["subnets"].SetOptional()
	attrs["target_region"] = attrs["target_region"].SetOptional()
	attrs["target_services"] = attrs["target_services"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAzureServiceEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccAzureServiceEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subnets":         reflect.TypeOf(types.String{}),
		"target_services": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzureServiceEndpointRule
// only implements ToObjectValue() and Type().
func (m NccAzureServiceEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"subnets":         m.Subnets,
			"target_region":   m.TargetRegion,
			"target_services": m.TargetServices,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccAzureServiceEndpointRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"subnets": basetypes.ListType{
				ElemType: types.StringType,
			},
			"target_region": types.StringType,
			"target_services": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetSubnets returns the value of the Subnets field in NccAzureServiceEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAzureServiceEndpointRule) GetSubnets(ctx context.Context) ([]types.String, bool) {
	if m.Subnets.IsNull() || m.Subnets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Subnets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnets sets the value of the Subnets field in NccAzureServiceEndpointRule.
func (m *NccAzureServiceEndpointRule) SetSubnets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Subnets = types.ListValueMust(t, vs)
}

// GetTargetServices returns the value of the TargetServices field in NccAzureServiceEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAzureServiceEndpointRule) GetTargetServices(ctx context.Context) ([]types.String, bool) {
	if m.TargetServices.IsNull() || m.TargetServices.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.TargetServices.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTargetServices sets the value of the TargetServices field in NccAzureServiceEndpointRule.
func (m *NccAzureServiceEndpointRule) SetTargetServices(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["target_services"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TargetServices = types.ListValueMust(t, vs)
}

type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules types.Object `tfsdk:"default_rules"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules types.Object `tfsdk:"target_rules"`
}

func (to *NccEgressConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccEgressConfig) {
	if !from.DefaultRules.IsNull() && !from.DefaultRules.IsUnknown() {
		if toDefaultRules, ok := to.GetDefaultRules(ctx); ok {
			if fromDefaultRules, ok := from.GetDefaultRules(ctx); ok {
				// Recursively sync the fields of DefaultRules
				toDefaultRules.SyncFieldsDuringCreateOrUpdate(ctx, fromDefaultRules)
				to.SetDefaultRules(ctx, toDefaultRules)
			}
		}
	}
	if !from.TargetRules.IsNull() && !from.TargetRules.IsUnknown() {
		if toTargetRules, ok := to.GetTargetRules(ctx); ok {
			if fromTargetRules, ok := from.GetTargetRules(ctx); ok {
				// Recursively sync the fields of TargetRules
				toTargetRules.SyncFieldsDuringCreateOrUpdate(ctx, fromTargetRules)
				to.SetTargetRules(ctx, toTargetRules)
			}
		}
	}
}

func (to *NccEgressConfig) SyncFieldsDuringRead(ctx context.Context, from NccEgressConfig) {
	if !from.DefaultRules.IsNull() && !from.DefaultRules.IsUnknown() {
		if toDefaultRules, ok := to.GetDefaultRules(ctx); ok {
			if fromDefaultRules, ok := from.GetDefaultRules(ctx); ok {
				toDefaultRules.SyncFieldsDuringRead(ctx, fromDefaultRules)
				to.SetDefaultRules(ctx, toDefaultRules)
			}
		}
	}
	if !from.TargetRules.IsNull() && !from.TargetRules.IsUnknown() {
		if toTargetRules, ok := to.GetTargetRules(ctx); ok {
			if fromTargetRules, ok := from.GetTargetRules(ctx); ok {
				toTargetRules.SyncFieldsDuringRead(ctx, fromTargetRules)
				to.SetTargetRules(ctx, toTargetRules)
			}
		}
	}
}

func (m NccEgressConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_rules"] = attrs["default_rules"].SetOptional()
	attrs["target_rules"] = attrs["target_rules"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccEgressConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_rules": reflect.TypeOf(NccEgressDefaultRules{}),
		"target_rules":  reflect.TypeOf(NccEgressTargetRules{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressConfig
// only implements ToObjectValue() and Type().
func (m NccEgressConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_rules": m.DefaultRules,
			"target_rules":  m.TargetRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccEgressConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_rules": NccEgressDefaultRules{}.Type(ctx),
			"target_rules":  NccEgressTargetRules{}.Type(ctx),
		},
	}
}

// GetDefaultRules returns the value of the DefaultRules field in NccEgressConfig as
// a NccEgressDefaultRules value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressConfig) GetDefaultRules(ctx context.Context) (NccEgressDefaultRules, bool) {
	var e NccEgressDefaultRules
	if m.DefaultRules.IsNull() || m.DefaultRules.IsUnknown() {
		return e, false
	}
	var v NccEgressDefaultRules
	d := m.DefaultRules.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDefaultRules sets the value of the DefaultRules field in NccEgressConfig.
func (m *NccEgressConfig) SetDefaultRules(ctx context.Context, v NccEgressDefaultRules) {
	vs := v.ToObjectValue(ctx)
	m.DefaultRules = vs
}

// GetTargetRules returns the value of the TargetRules field in NccEgressConfig as
// a NccEgressTargetRules value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressConfig) GetTargetRules(ctx context.Context) (NccEgressTargetRules, bool) {
	var e NccEgressTargetRules
	if m.TargetRules.IsNull() || m.TargetRules.IsUnknown() {
		return e, false
	}
	var v NccEgressTargetRules
	d := m.TargetRules.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTargetRules sets the value of the TargetRules field in NccEgressConfig.
func (m *NccEgressConfig) SetTargetRules(ctx context.Context, v NccEgressTargetRules) {
	vs := v.ToObjectValue(ctx)
	m.TargetRules = vs
}

// Default rules don't have specific targets.
type NccEgressDefaultRules struct {
	AwsStableIpRule types.Object `tfsdk:"aws_stable_ip_rule"`

	AzureServiceEndpointRule types.Object `tfsdk:"azure_service_endpoint_rule"`
}

func (to *NccEgressDefaultRules) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccEgressDefaultRules) {
	if !from.AwsStableIpRule.IsNull() && !from.AwsStableIpRule.IsUnknown() {
		if toAwsStableIpRule, ok := to.GetAwsStableIpRule(ctx); ok {
			if fromAwsStableIpRule, ok := from.GetAwsStableIpRule(ctx); ok {
				// Recursively sync the fields of AwsStableIpRule
				toAwsStableIpRule.SyncFieldsDuringCreateOrUpdate(ctx, fromAwsStableIpRule)
				to.SetAwsStableIpRule(ctx, toAwsStableIpRule)
			}
		}
	}
	if !from.AzureServiceEndpointRule.IsNull() && !from.AzureServiceEndpointRule.IsUnknown() {
		if toAzureServiceEndpointRule, ok := to.GetAzureServiceEndpointRule(ctx); ok {
			if fromAzureServiceEndpointRule, ok := from.GetAzureServiceEndpointRule(ctx); ok {
				// Recursively sync the fields of AzureServiceEndpointRule
				toAzureServiceEndpointRule.SyncFieldsDuringCreateOrUpdate(ctx, fromAzureServiceEndpointRule)
				to.SetAzureServiceEndpointRule(ctx, toAzureServiceEndpointRule)
			}
		}
	}
}

func (to *NccEgressDefaultRules) SyncFieldsDuringRead(ctx context.Context, from NccEgressDefaultRules) {
	if !from.AwsStableIpRule.IsNull() && !from.AwsStableIpRule.IsUnknown() {
		if toAwsStableIpRule, ok := to.GetAwsStableIpRule(ctx); ok {
			if fromAwsStableIpRule, ok := from.GetAwsStableIpRule(ctx); ok {
				toAwsStableIpRule.SyncFieldsDuringRead(ctx, fromAwsStableIpRule)
				to.SetAwsStableIpRule(ctx, toAwsStableIpRule)
			}
		}
	}
	if !from.AzureServiceEndpointRule.IsNull() && !from.AzureServiceEndpointRule.IsUnknown() {
		if toAzureServiceEndpointRule, ok := to.GetAzureServiceEndpointRule(ctx); ok {
			if fromAzureServiceEndpointRule, ok := from.GetAzureServiceEndpointRule(ctx); ok {
				toAzureServiceEndpointRule.SyncFieldsDuringRead(ctx, fromAzureServiceEndpointRule)
				to.SetAzureServiceEndpointRule(ctx, toAzureServiceEndpointRule)
			}
		}
	}
}

func (m NccEgressDefaultRules) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_stable_ip_rule"] = attrs["aws_stable_ip_rule"].SetOptional()
	attrs["azure_service_endpoint_rule"] = attrs["azure_service_endpoint_rule"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressDefaultRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccEgressDefaultRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_stable_ip_rule":          reflect.TypeOf(NccAwsStableIpRule{}),
		"azure_service_endpoint_rule": reflect.TypeOf(NccAzureServiceEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressDefaultRules
// only implements ToObjectValue() and Type().
func (m NccEgressDefaultRules) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_stable_ip_rule":          m.AwsStableIpRule,
			"azure_service_endpoint_rule": m.AzureServiceEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccEgressDefaultRules) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_stable_ip_rule":          NccAwsStableIpRule{}.Type(ctx),
			"azure_service_endpoint_rule": NccAzureServiceEndpointRule{}.Type(ctx),
		},
	}
}

// GetAwsStableIpRule returns the value of the AwsStableIpRule field in NccEgressDefaultRules as
// a NccAwsStableIpRule value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressDefaultRules) GetAwsStableIpRule(ctx context.Context) (NccAwsStableIpRule, bool) {
	var e NccAwsStableIpRule
	if m.AwsStableIpRule.IsNull() || m.AwsStableIpRule.IsUnknown() {
		return e, false
	}
	var v NccAwsStableIpRule
	d := m.AwsStableIpRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsStableIpRule sets the value of the AwsStableIpRule field in NccEgressDefaultRules.
func (m *NccEgressDefaultRules) SetAwsStableIpRule(ctx context.Context, v NccAwsStableIpRule) {
	vs := v.ToObjectValue(ctx)
	m.AwsStableIpRule = vs
}

// GetAzureServiceEndpointRule returns the value of the AzureServiceEndpointRule field in NccEgressDefaultRules as
// a NccAzureServiceEndpointRule value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressDefaultRules) GetAzureServiceEndpointRule(ctx context.Context) (NccAzureServiceEndpointRule, bool) {
	var e NccAzureServiceEndpointRule
	if m.AzureServiceEndpointRule.IsNull() || m.AzureServiceEndpointRule.IsUnknown() {
		return e, false
	}
	var v NccAzureServiceEndpointRule
	d := m.AzureServiceEndpointRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzureServiceEndpointRule sets the value of the AzureServiceEndpointRule field in NccEgressDefaultRules.
func (m *NccEgressDefaultRules) SetAzureServiceEndpointRule(ctx context.Context, v NccAzureServiceEndpointRule) {
	vs := v.ToObjectValue(ctx)
	m.AzureServiceEndpointRule = vs
}

// Target rule controls the egress rules that are dedicated to specific
// resources.
type NccEgressTargetRules struct {
	// AWS private endpoint rule controls the AWS private endpoint based egress
	// rules.
	AwsPrivateEndpointRules types.List `tfsdk:"aws_private_endpoint_rules"`

	AzurePrivateEndpointRules types.List `tfsdk:"azure_private_endpoint_rules"`
}

func (to *NccEgressTargetRules) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccEgressTargetRules) {
	if !from.AwsPrivateEndpointRules.IsNull() && !from.AwsPrivateEndpointRules.IsUnknown() && to.AwsPrivateEndpointRules.IsNull() && len(from.AwsPrivateEndpointRules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AwsPrivateEndpointRules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AwsPrivateEndpointRules = from.AwsPrivateEndpointRules
	}
	if !from.AzurePrivateEndpointRules.IsNull() && !from.AzurePrivateEndpointRules.IsUnknown() && to.AzurePrivateEndpointRules.IsNull() && len(from.AzurePrivateEndpointRules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AzurePrivateEndpointRules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AzurePrivateEndpointRules = from.AzurePrivateEndpointRules
	}
}

func (to *NccEgressTargetRules) SyncFieldsDuringRead(ctx context.Context, from NccEgressTargetRules) {
	if !from.AwsPrivateEndpointRules.IsNull() && !from.AwsPrivateEndpointRules.IsUnknown() && to.AwsPrivateEndpointRules.IsNull() && len(from.AwsPrivateEndpointRules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AwsPrivateEndpointRules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AwsPrivateEndpointRules = from.AwsPrivateEndpointRules
	}
	if !from.AzurePrivateEndpointRules.IsNull() && !from.AzurePrivateEndpointRules.IsUnknown() && to.AzurePrivateEndpointRules.IsNull() && len(from.AzurePrivateEndpointRules.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AzurePrivateEndpointRules, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AzurePrivateEndpointRules = from.AzurePrivateEndpointRules
	}
}

func (m NccEgressTargetRules) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_private_endpoint_rules"] = attrs["aws_private_endpoint_rules"].SetOptional()
	attrs["azure_private_endpoint_rules"] = attrs["azure_private_endpoint_rules"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressTargetRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccEgressTargetRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_private_endpoint_rules":   reflect.TypeOf(CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule{}),
		"azure_private_endpoint_rules": reflect.TypeOf(NccAzurePrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressTargetRules
// only implements ToObjectValue() and Type().
func (m NccEgressTargetRules) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_private_endpoint_rules":   m.AwsPrivateEndpointRules,
			"azure_private_endpoint_rules": m.AzurePrivateEndpointRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccEgressTargetRules) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_private_endpoint_rules": basetypes.ListType{
				ElemType: CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule{}.Type(ctx),
			},
			"azure_private_endpoint_rules": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule{}.Type(ctx),
			},
		},
	}
}

// GetAwsPrivateEndpointRules returns the value of the AwsPrivateEndpointRules field in NccEgressTargetRules as
// a slice of CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressTargetRules) GetAwsPrivateEndpointRules(ctx context.Context) ([]CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule, bool) {
	if m.AwsPrivateEndpointRules.IsNull() || m.AwsPrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule
	d := m.AwsPrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsPrivateEndpointRules sets the value of the AwsPrivateEndpointRules field in NccEgressTargetRules.
func (m *NccEgressTargetRules) SetAwsPrivateEndpointRules(ctx context.Context, v []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_private_endpoint_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AwsPrivateEndpointRules = types.ListValueMust(t, vs)
}

// GetAzurePrivateEndpointRules returns the value of the AzurePrivateEndpointRules field in NccEgressTargetRules as
// a slice of NccAzurePrivateEndpointRule values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressTargetRules) GetAzurePrivateEndpointRules(ctx context.Context) ([]NccAzurePrivateEndpointRule, bool) {
	if m.AzurePrivateEndpointRules.IsNull() || m.AzurePrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []NccAzurePrivateEndpointRule
	d := m.AzurePrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzurePrivateEndpointRules sets the value of the AzurePrivateEndpointRules field in NccEgressTargetRules.
func (m *NccEgressTargetRules) SetAzurePrivateEndpointRules(ctx context.Context, v []NccAzurePrivateEndpointRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_private_endpoint_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AzurePrivateEndpointRules = types.ListValueMust(t, vs)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type NccPrivateEndpointRule struct {
	// Databricks account ID. You can find your account ID from the Accounts
	// Console.
	AccountId types.String `tfsdk:"account_id"`
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the Cloud console
	// before they take effect. The possible values are: - PENDING: The endpoint
	// has been created and pending approval. - ESTABLISHED: The endpoint has
	// been approved and is ready to use in your serverless compute resources. -
	// REJECTED: Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up. - EXPIRED: If the endpoint was created but not approved in 14
	// days, it will be EXPIRED.
	ConnectionState types.String `tfsdk:"connection_state"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at"`
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames types.List `tfsdk:"domain_names"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the Azure private endpoint resource.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	EndpointService types.String `tfsdk:"endpoint_service"`
	// Not used by customer-managed private endpoint services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId types.String `tfsdk:"group_id"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames types.List `tfsdk:"resource_names"`
	// The ID of a private endpoint rule.
	RuleId types.String `tfsdk:"rule_id"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time"`
	// The AWS VPC endpoint ID. You can use this ID to identify the VPC endpoint
	// created by Databricks.
	VpcEndpointId types.String `tfsdk:"vpc_endpoint_id"`
}

func (to *NccPrivateEndpointRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccPrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (to *NccPrivateEndpointRule) SyncFieldsDuringRead(ctx context.Context, from NccPrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (m NccPrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["connection_state"] = attrs["connection_state"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["deactivated"] = attrs["deactivated"].SetOptional()
	attrs["deactivated_at"] = attrs["deactivated_at"].SetOptional()
	attrs["domain_names"] = attrs["domain_names"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetOptional()
	attrs["endpoint_service"] = attrs["endpoint_service"].SetOptional()
	attrs["group_id"] = attrs["group_id"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["resource_names"] = attrs["resource_names"].SetOptional()
	attrs["rule_id"] = attrs["rule_id"].SetOptional()
	attrs["updated_time"] = attrs["updated_time"].SetOptional()
	attrs["vpc_endpoint_id"] = attrs["vpc_endpoint_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccPrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccPrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccPrivateEndpointRule
// only implements ToObjectValue() and Type().
func (m NccPrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                     m.AccountId,
			"connection_state":               m.ConnectionState,
			"creation_time":                  m.CreationTime,
			"deactivated":                    m.Deactivated,
			"deactivated_at":                 m.DeactivatedAt,
			"domain_names":                   m.DomainNames,
			"enabled":                        m.Enabled,
			"endpoint_name":                  m.EndpointName,
			"endpoint_service":               m.EndpointService,
			"group_id":                       m.GroupId,
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"resource_id":                    m.ResourceId,
			"resource_names":                 m.ResourceNames,
			"rule_id":                        m.RuleId,
			"updated_time":                   m.UpdatedTime,
			"vpc_endpoint_id":                m.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccPrivateEndpointRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":       types.StringType,
			"connection_state": types.StringType,
			"creation_time":    types.Int64Type,
			"deactivated":      types.BoolType,
			"deactivated_at":   types.Int64Type,
			"domain_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"enabled":                        types.BoolType,
			"endpoint_name":                  types.StringType,
			"endpoint_service":               types.StringType,
			"group_id":                       types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_id":                    types.StringType,
			"resource_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"rule_id":         types.StringType,
			"updated_time":    types.Int64Type,
			"vpc_endpoint_id": types.StringType,
		},
	}
}

// GetDomainNames returns the value of the DomainNames field in NccPrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccPrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if m.DomainNames.IsNull() || m.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in NccPrivateEndpointRule.
func (m *NccPrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in NccPrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccPrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if m.ResourceNames.IsNull() || m.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in NccPrivateEndpointRule.
func (m *NccPrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

// Properties of the new network connectivity configuration.
type NetworkConnectivityConfiguration struct {
	// Your Databricks account ID. You can find your account ID in your
	// Databricks accounts console.
	AccountId types.String `tfsdk:"account_id"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig types.Object `tfsdk:"egress_config"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	Name types.String `tfsdk:"name"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"network_connectivity_config_id"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime types.Int64 `tfsdk:"updated_time"`
}

func (to *NetworkConnectivityConfiguration) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkConnectivityConfiguration) {
	if !from.EgressConfig.IsNull() && !from.EgressConfig.IsUnknown() {
		if toEgressConfig, ok := to.GetEgressConfig(ctx); ok {
			if fromEgressConfig, ok := from.GetEgressConfig(ctx); ok {
				// Recursively sync the fields of EgressConfig
				toEgressConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromEgressConfig)
				to.SetEgressConfig(ctx, toEgressConfig)
			}
		}
	}
}

func (to *NetworkConnectivityConfiguration) SyncFieldsDuringRead(ctx context.Context, from NetworkConnectivityConfiguration) {
	if !from.EgressConfig.IsNull() && !from.EgressConfig.IsUnknown() {
		if toEgressConfig, ok := to.GetEgressConfig(ctx); ok {
			if fromEgressConfig, ok := from.GetEgressConfig(ctx); ok {
				toEgressConfig.SyncFieldsDuringRead(ctx, fromEgressConfig)
				to.SetEgressConfig(ctx, toEgressConfig)
			}
		}
	}
}

func (m NetworkConnectivityConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["egress_config"] = attrs["egress_config"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["updated_time"] = attrs["updated_time"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkConnectivityConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NetworkConnectivityConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress_config": reflect.TypeOf(NccEgressConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkConnectivityConfiguration
// only implements ToObjectValue() and Type().
func (m NetworkConnectivityConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                     m.AccountId,
			"creation_time":                  m.CreationTime,
			"egress_config":                  m.EgressConfig,
			"name":                           m.Name,
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"region":                         m.Region,
			"updated_time":                   m.UpdatedTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkConnectivityConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":                     types.StringType,
			"creation_time":                  types.Int64Type,
			"egress_config":                  NccEgressConfig{}.Type(ctx),
			"name":                           types.StringType,
			"network_connectivity_config_id": types.StringType,
			"region":                         types.StringType,
			"updated_time":                   types.Int64Type,
		},
	}
}

// GetEgressConfig returns the value of the EgressConfig field in NetworkConnectivityConfiguration as
// a NccEgressConfig value.
// If the field is unknown or null, the boolean return value is false.
func (m *NetworkConnectivityConfiguration) GetEgressConfig(ctx context.Context) (NccEgressConfig, bool) {
	var e NccEgressConfig
	if m.EgressConfig.IsNull() || m.EgressConfig.IsUnknown() {
		return e, false
	}
	var v NccEgressConfig
	d := m.EgressConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEgressConfig sets the value of the EgressConfig field in NetworkConnectivityConfiguration.
func (m *NetworkConnectivityConfiguration) SetEgressConfig(ctx context.Context, v NccEgressConfig) {
	vs := v.ToObjectValue(ctx)
	m.EgressConfig = vs
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto). This policy should be consistent
// with [[com.databricks.api.proto.settingspolicy.EgressNetworkPolicy]]. Details
// see API-design:
// https://docs.google.com/document/d/1DKWO_FpZMCY4cF2O62LpwII1lx8gsnDGG-qgE3t3TOA/
type NetworkPolicyEgress struct {
	// The access policy enforced for egress traffic to the internet.
	NetworkAccess types.Object `tfsdk:"network_access"`
}

func (to *NetworkPolicyEgress) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkPolicyEgress) {
	if !from.NetworkAccess.IsNull() && !from.NetworkAccess.IsUnknown() {
		if toNetworkAccess, ok := to.GetNetworkAccess(ctx); ok {
			if fromNetworkAccess, ok := from.GetNetworkAccess(ctx); ok {
				// Recursively sync the fields of NetworkAccess
				toNetworkAccess.SyncFieldsDuringCreateOrUpdate(ctx, fromNetworkAccess)
				to.SetNetworkAccess(ctx, toNetworkAccess)
			}
		}
	}
}

func (to *NetworkPolicyEgress) SyncFieldsDuringRead(ctx context.Context, from NetworkPolicyEgress) {
	if !from.NetworkAccess.IsNull() && !from.NetworkAccess.IsUnknown() {
		if toNetworkAccess, ok := to.GetNetworkAccess(ctx); ok {
			if fromNetworkAccess, ok := from.GetNetworkAccess(ctx); ok {
				toNetworkAccess.SyncFieldsDuringRead(ctx, fromNetworkAccess)
				to.SetNetworkAccess(ctx, toNetworkAccess)
			}
		}
	}
}

func (m NetworkPolicyEgress) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_access"] = attrs["network_access"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkPolicyEgress.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NetworkPolicyEgress) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_access": reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkPolicyEgress
// only implements ToObjectValue() and Type().
func (m NetworkPolicyEgress) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_access": m.NetworkAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkPolicyEgress) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_access": EgressNetworkPolicyNetworkAccessPolicy{}.Type(ctx),
		},
	}
}

// GetNetworkAccess returns the value of the NetworkAccess field in NetworkPolicyEgress as
// a EgressNetworkPolicyNetworkAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *NetworkPolicyEgress) GetNetworkAccess(ctx context.Context) (EgressNetworkPolicyNetworkAccessPolicy, bool) {
	var e EgressNetworkPolicyNetworkAccessPolicy
	if m.NetworkAccess.IsNull() || m.NetworkAccess.IsUnknown() {
		return e, false
	}
	var v EgressNetworkPolicyNetworkAccessPolicy
	d := m.NetworkAccess.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNetworkAccess sets the value of the NetworkAccess field in NetworkPolicyEgress.
func (m *NetworkPolicyEgress) SetNetworkAccess(ctx context.Context, v EgressNetworkPolicyNetworkAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	m.NetworkAccess = vs
}

type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config types.Object `tfsdk:"config"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id"`
}

func (to *NotificationDestination) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotificationDestination) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (to *NotificationDestination) SyncFieldsDuringRead(ctx context.Context, from NotificationDestination) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m NotificationDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetOptional()
	attrs["destination_type"] = attrs["destination_type"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NotificationDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NotificationDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationDestination
// only implements ToObjectValue() and Type().
func (m NotificationDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":           m.Config,
			"destination_type": m.DestinationType,
			"display_name":     m.DisplayName,
			"id":               m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NotificationDestination) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config":           Config{}.Type(ctx),
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in NotificationDestination as
// a Config value.
// If the field is unknown or null, the boolean return value is false.
func (m *NotificationDestination) GetConfig(ctx context.Context) (Config, bool) {
	var e Config
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v Config
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in NotificationDestination.
func (m *NotificationDestination) SetConfig(ctx context.Context, v Config) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey types.String `tfsdk:"integration_key"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet types.Bool `tfsdk:"integration_key_set"`
}

func (to *PagerdutyConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PagerdutyConfig) {
}

func (to *PagerdutyConfig) SyncFieldsDuringRead(ctx context.Context, from PagerdutyConfig) {
}

func (m PagerdutyConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["integration_key"] = attrs["integration_key"].SetOptional()
	attrs["integration_key_set"] = attrs["integration_key_set"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PagerdutyConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PagerdutyConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PagerdutyConfig
// only implements ToObjectValue() and Type().
func (m PagerdutyConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_key":     m.IntegrationKey,
			"integration_key_set": m.IntegrationKeySet,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PagerdutyConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_key":     types.StringType,
			"integration_key_set": types.BoolType,
		},
	}
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *PartitionId) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PartitionId) {
}

func (to *PartitionId) SyncFieldsDuringRead(ctx context.Context, from PartitionId) {
}

func (m PartitionId) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PartitionId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionId
// only implements ToObjectValue() and Type().
func (m PartitionId) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PartitionId) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type PersonalComputeMessage struct {
	Value types.String `tfsdk:"value"`
}

func (to *PersonalComputeMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PersonalComputeMessage) {
}

func (to *PersonalComputeMessage) SyncFieldsDuringRead(ctx context.Context, from PersonalComputeMessage) {
}

func (m PersonalComputeMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalComputeMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PersonalComputeMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeMessage
// only implements ToObjectValue() and Type().
func (m PersonalComputeMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PersonalComputeMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type PersonalComputeSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	PersonalCompute types.Object `tfsdk:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *PersonalComputeSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PersonalComputeSetting) {
	if !from.PersonalCompute.IsNull() && !from.PersonalCompute.IsUnknown() {
		if toPersonalCompute, ok := to.GetPersonalCompute(ctx); ok {
			if fromPersonalCompute, ok := from.GetPersonalCompute(ctx); ok {
				// Recursively sync the fields of PersonalCompute
				toPersonalCompute.SyncFieldsDuringCreateOrUpdate(ctx, fromPersonalCompute)
				to.SetPersonalCompute(ctx, toPersonalCompute)
			}
		}
	}
}

func (to *PersonalComputeSetting) SyncFieldsDuringRead(ctx context.Context, from PersonalComputeSetting) {
	if !from.PersonalCompute.IsNull() && !from.PersonalCompute.IsUnknown() {
		if toPersonalCompute, ok := to.GetPersonalCompute(ctx); ok {
			if fromPersonalCompute, ok := from.GetPersonalCompute(ctx); ok {
				toPersonalCompute.SyncFieldsDuringRead(ctx, fromPersonalCompute)
				to.SetPersonalCompute(ctx, toPersonalCompute)
			}
		}
	}
}

func (m PersonalComputeSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].SetRequired()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PersonalComputeSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PersonalComputeSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personal_compute": reflect.TypeOf(PersonalComputeMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeSetting
// only implements ToObjectValue() and Type().
func (m PersonalComputeSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             m.Etag,
			"personal_compute": m.PersonalCompute,
			"setting_name":     m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PersonalComputeSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":             types.StringType,
			"personal_compute": PersonalComputeMessage{}.Type(ctx),
			"setting_name":     types.StringType,
		},
	}
}

// GetPersonalCompute returns the value of the PersonalCompute field in PersonalComputeSetting as
// a PersonalComputeMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *PersonalComputeSetting) GetPersonalCompute(ctx context.Context) (PersonalComputeMessage, bool) {
	var e PersonalComputeMessage
	if m.PersonalCompute.IsNull() || m.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v PersonalComputeMessage
	d := m.PersonalCompute.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPersonalCompute sets the value of the PersonalCompute field in PersonalComputeSetting.
func (m *PersonalComputeSetting) SetPersonalCompute(ctx context.Context, v PersonalComputeMessage) {
	vs := v.ToObjectValue(ctx)
	m.PersonalCompute = vs
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment types.String `tfsdk:"comment"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime types.Int64 `tfsdk:"expiry_time"`
	// The ID of this token.
	TokenId types.String `tfsdk:"token_id"`
}

func (to *PublicTokenInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublicTokenInfo) {
}

func (to *PublicTokenInfo) SyncFieldsDuringRead(ctx context.Context, from PublicTokenInfo) {
}

func (m PublicTokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["expiry_time"] = attrs["expiry_time"].SetOptional()
	attrs["token_id"] = attrs["token_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublicTokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PublicTokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublicTokenInfo
// only implements ToObjectValue() and Type().
func (m PublicTokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":       m.Comment,
			"creation_time": m.CreationTime,
			"expiry_time":   m.ExpiryTime,
			"token_id":      m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PublicTokenInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":       types.StringType,
			"creation_time": types.Int64Type,
			"expiry_time":   types.Int64Type,
			"token_id":      types.StringType,
		},
	}
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`

	ListType types.String `tfsdk:"list_type"`
}

func (to *ReplaceIpAccessList) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReplaceIpAccessList) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *ReplaceIpAccessList) SyncFieldsDuringRead(ctx context.Context, from ReplaceIpAccessList) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m ReplaceIpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetRequired()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetRequired()
	attrs["list_type"] = attrs["list_type"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ReplaceIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceIpAccessList
// only implements ToObjectValue() and Type().
func (m ReplaceIpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":           m.Enabled,
			"ip_access_list_id": m.IpAccessListId,
			"ip_addresses":      m.IpAddresses,
			"label":             m.Label,
			"list_type":         m.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ReplaceIpAccessList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":           types.BoolType,
			"ip_access_list_id": types.StringType,
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":     types.StringType,
			"list_type": types.StringType,
		},
	}
}

// GetIpAddresses returns the value of the IpAddresses field in ReplaceIpAccessList as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ReplaceIpAccessList) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if m.IpAddresses.IsNull() || m.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in ReplaceIpAccessList.
func (m *ReplaceIpAccessList) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

type RestrictWorkspaceAdminsMessage struct {
	Status types.String `tfsdk:"status"`
}

func (to *RestrictWorkspaceAdminsMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestrictWorkspaceAdminsMessage) {
}

func (to *RestrictWorkspaceAdminsMessage) SyncFieldsDuringRead(ctx context.Context, from RestrictWorkspaceAdminsMessage) {
}

func (m RestrictWorkspaceAdminsMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["status"] = attrs["status"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestrictWorkspaceAdminsMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestrictWorkspaceAdminsMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsMessage
// only implements ToObjectValue() and Type().
func (m RestrictWorkspaceAdminsMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestrictWorkspaceAdminsMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type RestrictWorkspaceAdminsSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	RestrictWorkspaceAdmins types.Object `tfsdk:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *RestrictWorkspaceAdminsSetting) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestrictWorkspaceAdminsSetting) {
	if !from.RestrictWorkspaceAdmins.IsNull() && !from.RestrictWorkspaceAdmins.IsUnknown() {
		if toRestrictWorkspaceAdmins, ok := to.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromRestrictWorkspaceAdmins, ok := from.GetRestrictWorkspaceAdmins(ctx); ok {
				// Recursively sync the fields of RestrictWorkspaceAdmins
				toRestrictWorkspaceAdmins.SyncFieldsDuringCreateOrUpdate(ctx, fromRestrictWorkspaceAdmins)
				to.SetRestrictWorkspaceAdmins(ctx, toRestrictWorkspaceAdmins)
			}
		}
	}
}

func (to *RestrictWorkspaceAdminsSetting) SyncFieldsDuringRead(ctx context.Context, from RestrictWorkspaceAdminsSetting) {
	if !from.RestrictWorkspaceAdmins.IsNull() && !from.RestrictWorkspaceAdmins.IsUnknown() {
		if toRestrictWorkspaceAdmins, ok := to.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromRestrictWorkspaceAdmins, ok := from.GetRestrictWorkspaceAdmins(ctx); ok {
				toRestrictWorkspaceAdmins.SyncFieldsDuringRead(ctx, fromRestrictWorkspaceAdmins)
				to.SetRestrictWorkspaceAdmins(ctx, toRestrictWorkspaceAdmins)
			}
		}
	}
}

func (m RestrictWorkspaceAdminsSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].SetRequired()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RestrictWorkspaceAdminsSetting.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RestrictWorkspaceAdminsSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"restrict_workspace_admins": reflect.TypeOf(RestrictWorkspaceAdminsMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsSetting
// only implements ToObjectValue() and Type().
func (m RestrictWorkspaceAdminsSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                      m.Etag,
			"restrict_workspace_admins": m.RestrictWorkspaceAdmins,
			"setting_name":              m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestrictWorkspaceAdminsSetting) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":                      types.StringType,
			"restrict_workspace_admins": RestrictWorkspaceAdminsMessage{}.Type(ctx),
			"setting_name":              types.StringType,
		},
	}
}

// GetRestrictWorkspaceAdmins returns the value of the RestrictWorkspaceAdmins field in RestrictWorkspaceAdminsSetting as
// a RestrictWorkspaceAdminsMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *RestrictWorkspaceAdminsSetting) GetRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage, bool) {
	var e RestrictWorkspaceAdminsMessage
	if m.RestrictWorkspaceAdmins.IsNull() || m.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v RestrictWorkspaceAdminsMessage
	d := m.RestrictWorkspaceAdmins.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in RestrictWorkspaceAdminsSetting.
func (m *RestrictWorkspaceAdminsSetting) SetRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage) {
	vs := v.ToObjectValue(ctx)
	m.RestrictWorkspaceAdmins = vs
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id"`
}

func (to *RevokeTokenRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RevokeTokenRequest) {
}

func (to *RevokeTokenRequest) SyncFieldsDuringRead(ctx context.Context, from RevokeTokenRequest) {
}

func (m RevokeTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_id"] = attrs["token_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RevokeTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenRequest
// only implements ToObjectValue() and Type().
func (m RevokeTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RevokeTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type RevokeTokenResponse struct {
}

func (to *RevokeTokenResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RevokeTokenResponse) {
}

func (to *RevokeTokenResponse) SyncFieldsDuringRead(ctx context.Context, from RevokeTokenResponse) {
}

func (m RevokeTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RevokeTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenResponse
// only implements ToObjectValue() and Type().
func (m RevokeTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RevokeTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SlackConfig struct {
	// [Input-Only] Slack channel ID for notifications.
	ChannelId types.String `tfsdk:"channel_id"`
	// [Output-Only] Whether channel ID is set.
	ChannelIdSet types.Bool `tfsdk:"channel_id_set"`
	// [Input-Only] OAuth token for Slack authentication.
	OauthToken types.String `tfsdk:"oauth_token"`
	// [Output-Only] Whether OAuth token is set.
	OauthTokenSet types.Bool `tfsdk:"oauth_token_set"`
	// [Input-Only] URL for Slack destination.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
}

func (to *SlackConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SlackConfig) {
}

func (to *SlackConfig) SyncFieldsDuringRead(ctx context.Context, from SlackConfig) {
}

func (m SlackConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["channel_id"] = attrs["channel_id"].SetOptional()
	attrs["channel_id_set"] = attrs["channel_id_set"].SetOptional()
	attrs["oauth_token"] = attrs["oauth_token"].SetOptional()
	attrs["oauth_token_set"] = attrs["oauth_token_set"].SetOptional()
	attrs["url"] = attrs["url"].SetOptional()
	attrs["url_set"] = attrs["url_set"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SlackConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SlackConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SlackConfig
// only implements ToObjectValue() and Type().
func (m SlackConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"channel_id":      m.ChannelId,
			"channel_id_set":  m.ChannelIdSet,
			"oauth_token":     m.OauthToken,
			"oauth_token_set": m.OauthTokenSet,
			"url":             m.Url,
			"url_set":         m.UrlSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SlackConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"channel_id":      types.StringType,
			"channel_id_set":  types.BoolType,
			"oauth_token":     types.StringType,
			"oauth_token_set": types.BoolType,
			"url":             types.StringType,
			"url_set":         types.BoolType,
		},
	}
}

type SqlResultsDownload struct {
	BooleanVal types.Object `tfsdk:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *SqlResultsDownload) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SqlResultsDownload) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				// Recursively sync the fields of BooleanVal
				toBooleanVal.SyncFieldsDuringCreateOrUpdate(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (to *SqlResultsDownload) SyncFieldsDuringRead(ctx context.Context, from SqlResultsDownload) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m SqlResultsDownload) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SqlResultsDownload.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m SqlResultsDownload) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlResultsDownload
// only implements ToObjectValue() and Type().
func (m SqlResultsDownload) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SqlResultsDownload) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val":  BooleanMessage{}.Type(ctx),
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in SqlResultsDownload as
// a BooleanMessage value.
// If the field is unknown or null, the boolean return value is false.
func (m *SqlResultsDownload) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v BooleanMessage
	d := m.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBooleanVal sets the value of the BooleanVal field in SqlResultsDownload.
func (m *SqlResultsDownload) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	m.BooleanVal = vs
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (to *StringMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StringMessage) {
}

func (to *StringMessage) SyncFieldsDuringRead(ctx context.Context, from StringMessage) {
}

func (m StringMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in StringMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m StringMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage
// only implements ToObjectValue() and Type().
func (m StringMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StringMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *TokenAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenAccessControlRequest) {
}

func (to *TokenAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from TokenAccessControlRequest) {
}

func (m TokenAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlRequest
// only implements ToObjectValue() and Type().
func (m TokenAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             m.GroupName,
			"permission_level":       m.PermissionLevel,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *TokenAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *TokenAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from TokenAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m TokenAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(TokenPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlResponse
// only implements ToObjectValue() and Type().
func (m TokenAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        m.AllPermissions,
			"display_name":           m.DisplayName,
			"group_name":             m.GroupName,
			"service_principal_name": m.ServicePrincipalName,
			"user_name":              m.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenAccessControlResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: TokenPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in TokenAccessControlResponse as
// a slice of TokenPermission values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenAccessControlResponse) GetAllPermissions(ctx context.Context) ([]TokenPermission, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []TokenPermission
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in TokenAccessControlResponse.
func (m *TokenAccessControlResponse) SetAllPermissions(ctx context.Context, v []TokenPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment types.String `tfsdk:"comment"`
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"created_by_id"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"created_by_username"`
	// Timestamp when the token was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Timestamp when the token expires.
	ExpiryTime types.Int64 `tfsdk:"expiry_time"`
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	LastUsedDay types.Int64 `tfsdk:"last_used_day"`
	// User ID of the user that owns the token.
	OwnerId types.Int64 `tfsdk:"owner_id"`
	// ID of the token.
	TokenId types.String `tfsdk:"token_id"`
	// If applicable, the ID of the workspace that the token was created in.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *TokenInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenInfo) {
}

func (to *TokenInfo) SyncFieldsDuringRead(ctx context.Context, from TokenInfo) {
}

func (m TokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["comment"] = attrs["comment"].SetOptional()
	attrs["created_by_id"] = attrs["created_by_id"].SetOptional()
	attrs["created_by_username"] = attrs["created_by_username"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["expiry_time"] = attrs["expiry_time"].SetOptional()
	attrs["last_used_day"] = attrs["last_used_day"].SetOptional()
	attrs["owner_id"] = attrs["owner_id"].SetOptional()
	attrs["token_id"] = attrs["token_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo
// only implements ToObjectValue() and Type().
func (m TokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":             m.Comment,
			"created_by_id":       m.CreatedById,
			"created_by_username": m.CreatedByUsername,
			"creation_time":       m.CreationTime,
			"expiry_time":         m.ExpiryTime,
			"last_used_day":       m.LastUsedDay,
			"owner_id":            m.OwnerId,
			"token_id":            m.TokenId,
			"workspace_id":        m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":             types.StringType,
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
			"creation_time":       types.Int64Type,
			"expiry_time":         types.Int64Type,
			"last_used_day":       types.Int64Type,
			"owner_id":            types.Int64Type,
			"token_id":            types.StringType,
			"workspace_id":        types.Int64Type,
		},
	}
}

type TokenPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *TokenPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *TokenPermission) SyncFieldsDuringRead(ctx context.Context, from TokenPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m TokenPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermission
// only implements ToObjectValue() and Type().
func (m TokenPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in TokenPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if m.InheritedFromObject.IsNull() || m.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in TokenPermission.
func (m *TokenPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type TokenPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *TokenPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *TokenPermissions) SyncFieldsDuringRead(ctx context.Context, from TokenPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m TokenPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissions
// only implements ToObjectValue() and Type().
func (m TokenPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermissions) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: TokenAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in TokenPermissions as
// a slice of TokenAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenPermissions) GetAccessControlList(ctx context.Context) ([]TokenAccessControlResponse, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlResponse
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissions.
func (m *TokenPermissions) SetAccessControlList(ctx context.Context, v []TokenAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type TokenPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *TokenPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermissionsDescription) {
}

func (to *TokenPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from TokenPermissionsDescription) {
}

func (m TokenPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsDescription
// only implements ToObjectValue() and Type().
func (m TokenPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type TokenPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
}

func (to *TokenPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *TokenPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from TokenPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m TokenPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TokenPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TokenPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsRequest
// only implements ToObjectValue() and Type().
func (m TokenPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: TokenAccessControlRequest{}.Type(ctx),
			},
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in TokenPermissionsRequest as
// a slice of TokenAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenPermissionsRequest) GetAccessControlList(ctx context.Context) ([]TokenAccessControlRequest, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlRequest
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissionsRequest.
func (m *TokenPermissionsRequest) SetAccessControlList(ctx context.Context, v []TokenAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateAccountIpAccessEnableRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateAccountIpAccessEnableRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAccountIpAccessEnableRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateAccountIpAccessEnableRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAccountIpAccessEnableRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAccountIpAccessEnableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAccountIpAccessEnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AccountIpAccessEnable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountIpAccessEnableRequest
// only implements ToObjectValue() and Type().
func (m UpdateAccountIpAccessEnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAccountIpAccessEnableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       AccountIpAccessEnable{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAccountIpAccessEnableRequest as
// a AccountIpAccessEnable value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAccountIpAccessEnableRequest) GetSetting(ctx context.Context) (AccountIpAccessEnable, bool) {
	var e AccountIpAccessEnable
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v AccountIpAccessEnable
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateAccountIpAccessEnableRequest.
func (m *UpdateAccountIpAccessEnableRequest) SetSetting(ctx context.Context, v AccountIpAccessEnable) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicySetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingAccessPolicySettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       AibiDashboardEmbeddingAccessPolicySetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest as
// a AibiDashboardEmbeddingAccessPolicySetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) GetSetting(ctx context.Context) (AibiDashboardEmbeddingAccessPolicySetting, bool) {
	var e AibiDashboardEmbeddingAccessPolicySetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingAccessPolicySetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest.
func (m *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SetSetting(ctx context.Context, v AibiDashboardEmbeddingAccessPolicySetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomainsSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       AibiDashboardEmbeddingApprovedDomainsSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest as
// a AibiDashboardEmbeddingApprovedDomainsSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetSetting(ctx context.Context) (AibiDashboardEmbeddingApprovedDomainsSetting, bool) {
	var e AibiDashboardEmbeddingApprovedDomainsSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v AibiDashboardEmbeddingApprovedDomainsSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest.
func (m *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SetSetting(ctx context.Context, v AibiDashboardEmbeddingApprovedDomainsSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateAutomaticClusterUpdateSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAutomaticClusterUpdateSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateAutomaticClusterUpdateSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateAutomaticClusterUpdateSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAutomaticClusterUpdateSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AutomaticClusterUpdateSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAutomaticClusterUpdateSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateAutomaticClusterUpdateSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAutomaticClusterUpdateSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       AutomaticClusterUpdateSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAutomaticClusterUpdateSettingRequest as
// a AutomaticClusterUpdateSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAutomaticClusterUpdateSettingRequest) GetSetting(ctx context.Context) (AutomaticClusterUpdateSetting, bool) {
	var e AutomaticClusterUpdateSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v AutomaticClusterUpdateSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateAutomaticClusterUpdateSettingRequest.
func (m *UpdateAutomaticClusterUpdateSettingRequest) SetSetting(ctx context.Context, v AutomaticClusterUpdateSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateComplianceSecurityProfileSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateComplianceSecurityProfileSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateComplianceSecurityProfileSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateComplianceSecurityProfileSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateComplianceSecurityProfileSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateComplianceSecurityProfileSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(ComplianceSecurityProfileSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComplianceSecurityProfileSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateComplianceSecurityProfileSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateComplianceSecurityProfileSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       ComplianceSecurityProfileSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateComplianceSecurityProfileSettingRequest as
// a ComplianceSecurityProfileSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateComplianceSecurityProfileSettingRequest) GetSetting(ctx context.Context) (ComplianceSecurityProfileSetting, bool) {
	var e ComplianceSecurityProfileSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v ComplianceSecurityProfileSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateComplianceSecurityProfileSettingRequest.
func (m *UpdateComplianceSecurityProfileSettingRequest) SetSetting(ctx context.Context, v ComplianceSecurityProfileSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateCspEnablementAccountSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCspEnablementAccountSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateCspEnablementAccountSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateCspEnablementAccountSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateCspEnablementAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCspEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(CspEnablementAccountSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCspEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateCspEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCspEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       CspEnablementAccountSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateCspEnablementAccountSettingRequest as
// a CspEnablementAccountSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCspEnablementAccountSettingRequest) GetSetting(ctx context.Context) (CspEnablementAccountSetting, bool) {
	var e CspEnablementAccountSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v CspEnablementAccountSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateCspEnablementAccountSettingRequest.
func (m *UpdateCspEnablementAccountSettingRequest) SetSetting(ctx context.Context, v CspEnablementAccountSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateDashboardEmailSubscriptionsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateDashboardEmailSubscriptionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDashboardEmailSubscriptionsRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateDashboardEmailSubscriptionsRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDashboardEmailSubscriptionsRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDashboardEmailSubscriptionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDashboardEmailSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDashboardEmailSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DashboardEmailSubscriptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDashboardEmailSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (m UpdateDashboardEmailSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDashboardEmailSubscriptionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       DashboardEmailSubscriptions{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDashboardEmailSubscriptionsRequest as
// a DashboardEmailSubscriptions value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDashboardEmailSubscriptionsRequest) GetSetting(ctx context.Context) (DashboardEmailSubscriptions, bool) {
	var e DashboardEmailSubscriptions
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v DashboardEmailSubscriptions
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateDashboardEmailSubscriptionsRequest.
func (m *UpdateDashboardEmailSubscriptionsRequest) SetSetting(ctx context.Context, v DashboardEmailSubscriptions) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateDefaultNamespaceSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultNamespaceSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateDefaultNamespaceSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultNamespaceSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDefaultNamespaceSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultNamespaceSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultNamespaceSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateDefaultNamespaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultNamespaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       DefaultNamespaceSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDefaultNamespaceSettingRequest as
// a DefaultNamespaceSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDefaultNamespaceSettingRequest) GetSetting(ctx context.Context) (DefaultNamespaceSetting, bool) {
	var e DefaultNamespaceSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v DefaultNamespaceSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateDefaultNamespaceSettingRequest.
func (m *UpdateDefaultNamespaceSettingRequest) SetSetting(ctx context.Context, v DefaultNamespaceSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateDefaultWarehouseIdRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateDefaultWarehouseIdRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultWarehouseIdRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateDefaultWarehouseIdRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultWarehouseIdRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDefaultWarehouseIdRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultWarehouseIdRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultWarehouseIdRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultWarehouseId{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultWarehouseIdRequest
// only implements ToObjectValue() and Type().
func (m UpdateDefaultWarehouseIdRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultWarehouseIdRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       DefaultWarehouseId{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDefaultWarehouseIdRequest as
// a DefaultWarehouseId value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDefaultWarehouseIdRequest) GetSetting(ctx context.Context) (DefaultWarehouseId, bool) {
	var e DefaultWarehouseId
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v DefaultWarehouseId
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateDefaultWarehouseIdRequest.
func (m *UpdateDefaultWarehouseIdRequest) SetSetting(ctx context.Context, v DefaultWarehouseId) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateDisableLegacyAccessRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateDisableLegacyAccessRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDisableLegacyAccessRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateDisableLegacyAccessRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDisableLegacyAccessRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDisableLegacyAccessRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyAccess{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyAccessRequest
// only implements ToObjectValue() and Type().
func (m UpdateDisableLegacyAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDisableLegacyAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       DisableLegacyAccess{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDisableLegacyAccessRequest as
// a DisableLegacyAccess value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDisableLegacyAccessRequest) GetSetting(ctx context.Context) (DisableLegacyAccess, bool) {
	var e DisableLegacyAccess
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v DisableLegacyAccess
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyAccessRequest.
func (m *UpdateDisableLegacyAccessRequest) SetSetting(ctx context.Context, v DisableLegacyAccess) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateDisableLegacyDbfsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateDisableLegacyDbfsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDisableLegacyDbfsRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateDisableLegacyDbfsRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDisableLegacyDbfsRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDisableLegacyDbfsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyDbfs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyDbfsRequest
// only implements ToObjectValue() and Type().
func (m UpdateDisableLegacyDbfsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDisableLegacyDbfsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       DisableLegacyDbfs{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDisableLegacyDbfsRequest as
// a DisableLegacyDbfs value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDisableLegacyDbfsRequest) GetSetting(ctx context.Context) (DisableLegacyDbfs, bool) {
	var e DisableLegacyDbfs
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v DisableLegacyDbfs
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyDbfsRequest.
func (m *UpdateDisableLegacyDbfsRequest) SetSetting(ctx context.Context, v DisableLegacyDbfs) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateDisableLegacyFeaturesRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateDisableLegacyFeaturesRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDisableLegacyFeaturesRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateDisableLegacyFeaturesRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateDisableLegacyFeaturesRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDisableLegacyFeaturesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyFeatures{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyFeaturesRequest
// only implements ToObjectValue() and Type().
func (m UpdateDisableLegacyFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDisableLegacyFeaturesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       DisableLegacyFeatures{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDisableLegacyFeaturesRequest as
// a DisableLegacyFeatures value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDisableLegacyFeaturesRequest) GetSetting(ctx context.Context) (DisableLegacyFeatures, bool) {
	var e DisableLegacyFeatures
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v DisableLegacyFeatures
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyFeaturesRequest.
func (m *UpdateDisableLegacyFeaturesRequest) SetSetting(ctx context.Context, v DisableLegacyFeatures) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateEnableExportNotebookRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateEnableExportNotebookRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnableExportNotebookRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateEnableExportNotebookRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEnableExportNotebookRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnableExportNotebookRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnableExportNotebookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnableExportNotebookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableExportNotebook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableExportNotebookRequest
// only implements ToObjectValue() and Type().
func (m UpdateEnableExportNotebookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnableExportNotebookRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       EnableExportNotebook{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnableExportNotebookRequest as
// a EnableExportNotebook value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnableExportNotebookRequest) GetSetting(ctx context.Context) (EnableExportNotebook, bool) {
	var e EnableExportNotebook
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v EnableExportNotebook
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateEnableExportNotebookRequest.
func (m *UpdateEnableExportNotebookRequest) SetSetting(ctx context.Context, v EnableExportNotebook) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateEnableNotebookTableClipboardRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateEnableNotebookTableClipboardRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnableNotebookTableClipboardRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateEnableNotebookTableClipboardRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEnableNotebookTableClipboardRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnableNotebookTableClipboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnableNotebookTableClipboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnableNotebookTableClipboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableNotebookTableClipboard{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableNotebookTableClipboardRequest
// only implements ToObjectValue() and Type().
func (m UpdateEnableNotebookTableClipboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnableNotebookTableClipboardRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       EnableNotebookTableClipboard{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnableNotebookTableClipboardRequest as
// a EnableNotebookTableClipboard value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnableNotebookTableClipboardRequest) GetSetting(ctx context.Context) (EnableNotebookTableClipboard, bool) {
	var e EnableNotebookTableClipboard
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v EnableNotebookTableClipboard
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateEnableNotebookTableClipboardRequest.
func (m *UpdateEnableNotebookTableClipboardRequest) SetSetting(ctx context.Context, v EnableNotebookTableClipboard) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateEnableResultsDownloadingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateEnableResultsDownloadingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnableResultsDownloadingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateEnableResultsDownloadingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEnableResultsDownloadingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnableResultsDownloadingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnableResultsDownloadingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnableResultsDownloadingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableResultsDownloading{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableResultsDownloadingRequest
// only implements ToObjectValue() and Type().
func (m UpdateEnableResultsDownloadingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnableResultsDownloadingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       EnableResultsDownloading{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnableResultsDownloadingRequest as
// a EnableResultsDownloading value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnableResultsDownloadingRequest) GetSetting(ctx context.Context) (EnableResultsDownloading, bool) {
	var e EnableResultsDownloading
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v EnableResultsDownloading
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateEnableResultsDownloadingRequest.
func (m *UpdateEnableResultsDownloadingRequest) SetSetting(ctx context.Context, v EnableResultsDownloading) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateEnhancedSecurityMonitoringSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnhancedSecurityMonitoringSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateEnhancedSecurityMonitoringSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEnhancedSecurityMonitoringSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnhancedSecurityMonitoringSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnhancedSecurityMonitoringSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnhancedSecurityMonitoringSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateEnhancedSecurityMonitoringSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnhancedSecurityMonitoringSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       EnhancedSecurityMonitoringSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnhancedSecurityMonitoringSettingRequest as
// a EnhancedSecurityMonitoringSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnhancedSecurityMonitoringSettingRequest) GetSetting(ctx context.Context) (EnhancedSecurityMonitoringSetting, bool) {
	var e EnhancedSecurityMonitoringSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v EnhancedSecurityMonitoringSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateEnhancedSecurityMonitoringSettingRequest.
func (m *UpdateEnhancedSecurityMonitoringSettingRequest) SetSetting(ctx context.Context, v EnhancedSecurityMonitoringSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateEsmEnablementAccountSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEsmEnablementAccountSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateEsmEnablementAccountSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateEsmEnablementAccountSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEsmEnablementAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEsmEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EsmEnablementAccountSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEsmEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateEsmEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEsmEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       EsmEnablementAccountSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEsmEnablementAccountSettingRequest as
// a EsmEnablementAccountSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEsmEnablementAccountSettingRequest) GetSetting(ctx context.Context) (EsmEnablementAccountSetting, bool) {
	var e EsmEnablementAccountSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v EsmEnablementAccountSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateEsmEnablementAccountSettingRequest.
func (m *UpdateEsmEnablementAccountSettingRequest) SetSetting(ctx context.Context, v EsmEnablementAccountSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`

	ListType types.String `tfsdk:"list_type"`
}

func (to *UpdateIpAccessList) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateIpAccessList) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *UpdateIpAccessList) SyncFieldsDuringRead(ctx context.Context, from UpdateIpAccessList) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m UpdateIpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetOptional()
	attrs["list_type"] = attrs["list_type"].SetOptional()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateIpAccessList
// only implements ToObjectValue() and Type().
func (m UpdateIpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":           m.Enabled,
			"ip_access_list_id": m.IpAccessListId,
			"ip_addresses":      m.IpAddresses,
			"label":             m.Label,
			"list_type":         m.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateIpAccessList) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled":           types.BoolType,
			"ip_access_list_id": types.StringType,
			"ip_addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
			"label":     types.StringType,
			"list_type": types.StringType,
		},
	}
}

// GetIpAddresses returns the value of the IpAddresses field in UpdateIpAccessList as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateIpAccessList) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if m.IpAddresses.IsNull() || m.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in UpdateIpAccessList.
func (m *UpdateIpAccessList) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredAccountRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateLlmProxyPartnerPoweredAccountRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLlmProxyPartnerPoweredAccountRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateLlmProxyPartnerPoweredAccountRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateLlmProxyPartnerPoweredAccountRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateLlmProxyPartnerPoweredAccountRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateLlmProxyPartnerPoweredAccountRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredAccountRequest
// only implements ToObjectValue() and Type().
func (m UpdateLlmProxyPartnerPoweredAccountRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLlmProxyPartnerPoweredAccountRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       LlmProxyPartnerPoweredAccount{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateLlmProxyPartnerPoweredAccountRequest as
// a LlmProxyPartnerPoweredAccount value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateLlmProxyPartnerPoweredAccountRequest) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredAccount, bool) {
	var e LlmProxyPartnerPoweredAccount
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v LlmProxyPartnerPoweredAccount
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredAccountRequest.
func (m *UpdateLlmProxyPartnerPoweredAccountRequest) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredAccount) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredEnforceRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateLlmProxyPartnerPoweredEnforceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLlmProxyPartnerPoweredEnforceRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateLlmProxyPartnerPoweredEnforceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateLlmProxyPartnerPoweredEnforceRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateLlmProxyPartnerPoweredEnforceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredEnforceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateLlmProxyPartnerPoweredEnforceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredEnforce{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredEnforceRequest
// only implements ToObjectValue() and Type().
func (m UpdateLlmProxyPartnerPoweredEnforceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLlmProxyPartnerPoweredEnforceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       LlmProxyPartnerPoweredEnforce{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateLlmProxyPartnerPoweredEnforceRequest as
// a LlmProxyPartnerPoweredEnforce value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateLlmProxyPartnerPoweredEnforceRequest) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredEnforce, bool) {
	var e LlmProxyPartnerPoweredEnforce
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v LlmProxyPartnerPoweredEnforce
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredEnforceRequest.
func (m *UpdateLlmProxyPartnerPoweredEnforceRequest) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredEnforce) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredWorkspaceRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateLlmProxyPartnerPoweredWorkspaceRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLlmProxyPartnerPoweredWorkspaceRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateLlmProxyPartnerPoweredWorkspaceRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateLlmProxyPartnerPoweredWorkspaceRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredWorkspace{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredWorkspaceRequest
// only implements ToObjectValue() and Type().
func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       LlmProxyPartnerPoweredWorkspace{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateLlmProxyPartnerPoweredWorkspaceRequest as
// a LlmProxyPartnerPoweredWorkspace value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateLlmProxyPartnerPoweredWorkspaceRequest) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredWorkspace, bool) {
	var e LlmProxyPartnerPoweredWorkspace
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v LlmProxyPartnerPoweredWorkspace
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredWorkspaceRequest.
func (m *UpdateLlmProxyPartnerPoweredWorkspaceRequest) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredWorkspace) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

type UpdateNccPrivateEndpointRuleRequest struct {
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`

	PrivateEndpointRule types.Object `tfsdk:"private_endpoint_rule"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	UpdateMask types.String `tfsdk:"-"`
}

func (to *UpdateNccPrivateEndpointRuleRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateNccPrivateEndpointRuleRequest) {
	if !from.PrivateEndpointRule.IsNull() && !from.PrivateEndpointRule.IsUnknown() {
		if toPrivateEndpointRule, ok := to.GetPrivateEndpointRule(ctx); ok {
			if fromPrivateEndpointRule, ok := from.GetPrivateEndpointRule(ctx); ok {
				// Recursively sync the fields of PrivateEndpointRule
				toPrivateEndpointRule.SyncFieldsDuringCreateOrUpdate(ctx, fromPrivateEndpointRule)
				to.SetPrivateEndpointRule(ctx, toPrivateEndpointRule)
			}
		}
	}
}

func (to *UpdateNccPrivateEndpointRuleRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateNccPrivateEndpointRuleRequest) {
	if !from.PrivateEndpointRule.IsNull() && !from.PrivateEndpointRule.IsUnknown() {
		if toPrivateEndpointRule, ok := to.GetPrivateEndpointRule(ctx); ok {
			if fromPrivateEndpointRule, ok := from.GetPrivateEndpointRule(ctx); ok {
				toPrivateEndpointRule.SyncFieldsDuringRead(ctx, fromPrivateEndpointRule)
				to.SetPrivateEndpointRule(ctx, toPrivateEndpointRule)
			}
		}
	}
}

func (m UpdateNccPrivateEndpointRuleRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_endpoint_rule"] = attrs["private_endpoint_rule"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()
	attrs["private_endpoint_rule_id"] = attrs["private_endpoint_rule_id"].SetRequired()
	attrs["update_mask"] = attrs["update_mask"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateNccPrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateNccPrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_endpoint_rule": reflect.TypeOf(UpdatePrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNccPrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (m UpdateNccPrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule":          m.PrivateEndpointRule,
			"private_endpoint_rule_id":       m.PrivateEndpointRuleId,
			"update_mask":                    m.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateNccPrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule":          UpdatePrivateEndpointRule{}.Type(ctx),
			"private_endpoint_rule_id":       types.StringType,
			"update_mask":                    types.StringType,
		},
	}
}

// GetPrivateEndpointRule returns the value of the PrivateEndpointRule field in UpdateNccPrivateEndpointRuleRequest as
// a UpdatePrivateEndpointRule value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateNccPrivateEndpointRuleRequest) GetPrivateEndpointRule(ctx context.Context) (UpdatePrivateEndpointRule, bool) {
	var e UpdatePrivateEndpointRule
	if m.PrivateEndpointRule.IsNull() || m.PrivateEndpointRule.IsUnknown() {
		return e, false
	}
	var v UpdatePrivateEndpointRule
	d := m.PrivateEndpointRule.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPrivateEndpointRule sets the value of the PrivateEndpointRule field in UpdateNccPrivateEndpointRuleRequest.
func (m *UpdateNccPrivateEndpointRuleRequest) SetPrivateEndpointRule(ctx context.Context, v UpdatePrivateEndpointRule) {
	vs := v.ToObjectValue(ctx)
	m.PrivateEndpointRule = vs
}

type UpdateNetworkPolicyRequest struct {
	// Updated network policy configuration details.
	NetworkPolicy types.Object `tfsdk:"network_policy"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"-"`
}

func (to *UpdateNetworkPolicyRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateNetworkPolicyRequest) {
	if !from.NetworkPolicy.IsNull() && !from.NetworkPolicy.IsUnknown() {
		if toNetworkPolicy, ok := to.GetNetworkPolicy(ctx); ok {
			if fromNetworkPolicy, ok := from.GetNetworkPolicy(ctx); ok {
				// Recursively sync the fields of NetworkPolicy
				toNetworkPolicy.SyncFieldsDuringCreateOrUpdate(ctx, fromNetworkPolicy)
				to.SetNetworkPolicy(ctx, toNetworkPolicy)
			}
		}
	}
}

func (to *UpdateNetworkPolicyRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateNetworkPolicyRequest) {
	if !from.NetworkPolicy.IsNull() && !from.NetworkPolicy.IsUnknown() {
		if toNetworkPolicy, ok := to.GetNetworkPolicy(ctx); ok {
			if fromNetworkPolicy, ok := from.GetNetworkPolicy(ctx); ok {
				toNetworkPolicy.SyncFieldsDuringRead(ctx, fromNetworkPolicy)
				to.SetNetworkPolicy(ctx, toNetworkPolicy)
			}
		}
	}
}

func (m UpdateNetworkPolicyRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_policy"] = attrs["network_policy"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["network_policy_id"] = attrs["network_policy_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_policy": reflect.TypeOf(AccountNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (m UpdateNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy":    m.NetworkPolicy,
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy":    AccountNetworkPolicy{}.Type(ctx),
			"network_policy_id": types.StringType,
		},
	}
}

// GetNetworkPolicy returns the value of the NetworkPolicy field in UpdateNetworkPolicyRequest as
// a AccountNetworkPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateNetworkPolicyRequest) GetNetworkPolicy(ctx context.Context) (AccountNetworkPolicy, bool) {
	var e AccountNetworkPolicy
	if m.NetworkPolicy.IsNull() || m.NetworkPolicy.IsUnknown() {
		return e, false
	}
	var v AccountNetworkPolicy
	d := m.NetworkPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetNetworkPolicy sets the value of the NetworkPolicy field in UpdateNetworkPolicyRequest.
func (m *UpdateNetworkPolicyRequest) SetNetworkPolicy(ctx context.Context, v AccountNetworkPolicy) {
	vs := v.ToObjectValue(ctx)
	m.NetworkPolicy = vs
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.Object `tfsdk:"config"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"-"`
}

func (to *UpdateNotificationDestinationRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateNotificationDestinationRequest) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (to *UpdateNotificationDestinationRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateNotificationDestinationRequest) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m UpdateNotificationDestinationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["id"] = attrs["id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (m UpdateNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       m.Config,
			"display_name": m.DisplayName,
			"id":           m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config":       Config{}.Type(ctx),
			"display_name": types.StringType,
			"id":           types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in UpdateNotificationDestinationRequest as
// a Config value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateNotificationDestinationRequest) GetConfig(ctx context.Context) (Config, bool) {
	var e Config
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v Config
	d := m.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in UpdateNotificationDestinationRequest.
func (m *UpdateNotificationDestinationRequest) SetConfig(ctx context.Context, v Config) {
	vs := v.ToObjectValue(ctx)
	m.Config = vs
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdatePersonalComputeSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePersonalComputeSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdatePersonalComputeSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdatePersonalComputeSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdatePersonalComputeSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(PersonalComputeSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalComputeSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdatePersonalComputeSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePersonalComputeSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       PersonalComputeSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdatePersonalComputeSettingRequest as
// a PersonalComputeSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePersonalComputeSettingRequest) GetSetting(ctx context.Context) (PersonalComputeSetting, bool) {
	var e PersonalComputeSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v PersonalComputeSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdatePersonalComputeSettingRequest.
func (m *UpdatePersonalComputeSettingRequest) SetSetting(ctx context.Context, v PersonalComputeSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type UpdatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames types.List `tfsdk:"domain_names"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	Enabled types.Bool `tfsdk:"enabled"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	ResourceNames types.List `tfsdk:"resource_names"`
}

func (to *UpdatePrivateEndpointRule) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (to *UpdatePrivateEndpointRule) SyncFieldsDuringRead(ctx context.Context, from UpdatePrivateEndpointRule) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
	if !from.ResourceNames.IsNull() && !from.ResourceNames.IsUnknown() && to.ResourceNames.IsNull() && len(from.ResourceNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ResourceNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ResourceNames = from.ResourceNames
	}
}

func (m UpdatePrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["domain_names"] = attrs["domain_names"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["resource_names"] = attrs["resource_names"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePrivateEndpointRule
// only implements ToObjectValue() and Type().
func (m UpdatePrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain_names":   m.DomainNames,
			"enabled":        m.Enabled,
			"resource_names": m.ResourceNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePrivateEndpointRule) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"domain_names": basetypes.ListType{
				ElemType: types.StringType,
			},
			"enabled": types.BoolType,
			"resource_names": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDomainNames returns the value of the DomainNames field in UpdatePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if m.DomainNames.IsNull() || m.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in UpdatePrivateEndpointRule.
func (m *UpdatePrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in UpdatePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if m.ResourceNames.IsNull() || m.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in UpdatePrivateEndpointRule.
func (m *UpdatePrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateRestrictWorkspaceAdminsSettingRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRestrictWorkspaceAdminsSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateRestrictWorkspaceAdminsSettingRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateRestrictWorkspaceAdminsSettingRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateRestrictWorkspaceAdminsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(RestrictWorkspaceAdminsSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRestrictWorkspaceAdminsSettingRequest
// only implements ToObjectValue() and Type().
func (m UpdateRestrictWorkspaceAdminsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRestrictWorkspaceAdminsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       RestrictWorkspaceAdminsSetting{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateRestrictWorkspaceAdminsSettingRequest as
// a RestrictWorkspaceAdminsSetting value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRestrictWorkspaceAdminsSettingRequest) GetSetting(ctx context.Context) (RestrictWorkspaceAdminsSetting, bool) {
	var e RestrictWorkspaceAdminsSetting
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v RestrictWorkspaceAdminsSetting
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateRestrictWorkspaceAdminsSettingRequest.
func (m *UpdateRestrictWorkspaceAdminsSettingRequest) SetSetting(ctx context.Context, v RestrictWorkspaceAdminsSetting) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

// Details required to update a setting.
type UpdateSqlResultsDownloadRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing types.Bool `tfsdk:"allow_missing"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	FieldMask types.String `tfsdk:"field_mask"`

	Setting types.Object `tfsdk:"setting"`
}

func (to *UpdateSqlResultsDownloadRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSqlResultsDownloadRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				// Recursively sync the fields of Setting
				toSetting.SyncFieldsDuringCreateOrUpdate(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (to *UpdateSqlResultsDownloadRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateSqlResultsDownloadRequest) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateSqlResultsDownloadRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSqlResultsDownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSqlResultsDownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(SqlResultsDownload{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSqlResultsDownloadRequest
// only implements ToObjectValue() and Type().
func (m UpdateSqlResultsDownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSqlResultsDownloadRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting":       SqlResultsDownload{}.Type(ctx),
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateSqlResultsDownloadRequest as
// a SqlResultsDownload value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSqlResultsDownloadRequest) GetSetting(ctx context.Context) (SqlResultsDownload, bool) {
	var e SqlResultsDownload
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v SqlResultsDownload
	d := m.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSetting sets the value of the Setting field in UpdateSqlResultsDownloadRequest.
func (m *UpdateSqlResultsDownloadRequest) SetSetting(ctx context.Context, v SqlResultsDownload) {
	vs := v.ToObjectValue(ctx)
	m.Setting = vs
}

type UpdateWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
	// The network option details for the workspace.
	WorkspaceNetworkOption types.Object `tfsdk:"workspace_network_option"`
}

func (to *UpdateWorkspaceNetworkOptionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceNetworkOptionRequest) {
	if !from.WorkspaceNetworkOption.IsNull() && !from.WorkspaceNetworkOption.IsUnknown() {
		if toWorkspaceNetworkOption, ok := to.GetWorkspaceNetworkOption(ctx); ok {
			if fromWorkspaceNetworkOption, ok := from.GetWorkspaceNetworkOption(ctx); ok {
				// Recursively sync the fields of WorkspaceNetworkOption
				toWorkspaceNetworkOption.SyncFieldsDuringCreateOrUpdate(ctx, fromWorkspaceNetworkOption)
				to.SetWorkspaceNetworkOption(ctx, toWorkspaceNetworkOption)
			}
		}
	}
}

func (to *UpdateWorkspaceNetworkOptionRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceNetworkOptionRequest) {
	if !from.WorkspaceNetworkOption.IsNull() && !from.WorkspaceNetworkOption.IsUnknown() {
		if toWorkspaceNetworkOption, ok := to.GetWorkspaceNetworkOption(ctx); ok {
			if fromWorkspaceNetworkOption, ok := from.GetWorkspaceNetworkOption(ctx); ok {
				toWorkspaceNetworkOption.SyncFieldsDuringRead(ctx, fromWorkspaceNetworkOption)
				to.SetWorkspaceNetworkOption(ctx, toWorkspaceNetworkOption)
			}
		}
	}
}

func (m UpdateWorkspaceNetworkOptionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_network_option"] = attrs["workspace_network_option"].SetRequired()
	attrs["account_id"] = attrs["account_id"].SetRequired()
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceNetworkOptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateWorkspaceNetworkOptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_network_option": reflect.TypeOf(WorkspaceNetworkOption{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceNetworkOptionRequest
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceNetworkOptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id":             m.WorkspaceId,
			"workspace_network_option": m.WorkspaceNetworkOption,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceNetworkOptionRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id":             types.Int64Type,
			"workspace_network_option": WorkspaceNetworkOption{}.Type(ctx),
		},
	}
}

// GetWorkspaceNetworkOption returns the value of the WorkspaceNetworkOption field in UpdateWorkspaceNetworkOptionRequest as
// a WorkspaceNetworkOption value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceNetworkOptionRequest) GetWorkspaceNetworkOption(ctx context.Context) (WorkspaceNetworkOption, bool) {
	var e WorkspaceNetworkOption
	if m.WorkspaceNetworkOption.IsNull() || m.WorkspaceNetworkOption.IsUnknown() {
		return e, false
	}
	var v WorkspaceNetworkOption
	d := m.WorkspaceNetworkOption.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkspaceNetworkOption sets the value of the WorkspaceNetworkOption field in UpdateWorkspaceNetworkOptionRequest.
func (m *UpdateWorkspaceNetworkOptionRequest) SetWorkspaceNetworkOption(ctx context.Context, v WorkspaceNetworkOption) {
	vs := v.ToObjectValue(ctx)
	m.WorkspaceNetworkOption = vs
}

type WorkspaceNetworkOption struct {
	// The network policy ID to apply to the workspace. This controls the
	// network access rules for all serverless compute resources in the
	// workspace. Each workspace can only be linked to one policy at a time. If
	// no policy is explicitly assigned, the workspace will use
	// 'default-policy'.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *WorkspaceNetworkOption) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceNetworkOption) {
}

func (to *WorkspaceNetworkOption) SyncFieldsDuringRead(ctx context.Context, from WorkspaceNetworkOption) {
}

func (m WorkspaceNetworkOption) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_policy_id"] = attrs["network_policy_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in WorkspaceNetworkOption.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m WorkspaceNetworkOption) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetworkOption
// only implements ToObjectValue() and Type().
func (m WorkspaceNetworkOption) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": m.NetworkPolicyId,
			"workspace_id":      m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceNetworkOption) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
			"workspace_id":      types.Int64Type,
		},
	}
}
