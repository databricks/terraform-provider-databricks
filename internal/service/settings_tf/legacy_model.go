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

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AccountIpAccessEnable_SdkV2 struct {
	AcctIpAclEnable types.List `tfsdk:"acct_ip_acl_enable"`
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

func (to *AccountIpAccessEnable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccountIpAccessEnable_SdkV2) {
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

func (to *AccountIpAccessEnable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AccountIpAccessEnable_SdkV2) {
	if !from.AcctIpAclEnable.IsNull() && !from.AcctIpAclEnable.IsUnknown() {
		if toAcctIpAclEnable, ok := to.GetAcctIpAclEnable(ctx); ok {
			if fromAcctIpAclEnable, ok := from.GetAcctIpAclEnable(ctx); ok {
				toAcctIpAclEnable.SyncFieldsDuringRead(ctx, fromAcctIpAclEnable)
				to.SetAcctIpAclEnable(ctx, toAcctIpAclEnable)
			}
		}
	}
}

func (m AccountIpAccessEnable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["acct_ip_acl_enable"] = attrs["acct_ip_acl_enable"].SetRequired()
	attrs["acct_ip_acl_enable"] = attrs["acct_ip_acl_enable"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m AccountIpAccessEnable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"acct_ip_acl_enable": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountIpAccessEnable_SdkV2
// only implements ToObjectValue() and Type().
func (m AccountIpAccessEnable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"acct_ip_acl_enable": m.AcctIpAclEnable,
			"etag":               m.Etag,
			"setting_name":       m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AccountIpAccessEnable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"acct_ip_acl_enable": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetAcctIpAclEnable returns the value of the AcctIpAclEnable field in AccountIpAccessEnable_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AccountIpAccessEnable_SdkV2) GetAcctIpAclEnable(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.AcctIpAclEnable.IsNull() || m.AcctIpAclEnable.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.AcctIpAclEnable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAcctIpAclEnable sets the value of the AcctIpAclEnable field in AccountIpAccessEnable_SdkV2.
func (m *AccountIpAccessEnable_SdkV2) SetAcctIpAclEnable(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["acct_ip_acl_enable"]
	m.AcctIpAclEnable = types.ListValueMust(t, vs)
}

type AccountNetworkPolicy_SdkV2 struct {
	// The network policies applying for egress traffic.
	Egress types.List `tfsdk:"egress"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
}

func (to *AccountNetworkPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AccountNetworkPolicy_SdkV2) {
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

func (to *AccountNetworkPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AccountNetworkPolicy_SdkV2) {
	if !from.Egress.IsNull() && !from.Egress.IsUnknown() {
		if toEgress, ok := to.GetEgress(ctx); ok {
			if fromEgress, ok := from.GetEgress(ctx); ok {
				toEgress.SyncFieldsDuringRead(ctx, fromEgress)
				to.SetEgress(ctx, toEgress)
			}
		}
	}
}

func (m AccountNetworkPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["egress"] = attrs["egress"].SetOptional()
	attrs["egress"] = attrs["egress"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m AccountNetworkPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress": reflect.TypeOf(NetworkPolicyEgress_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountNetworkPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m AccountNetworkPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"egress":            m.Egress,
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AccountNetworkPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"egress": basetypes.ListType{
				ElemType: NetworkPolicyEgress_SdkV2{}.Type(ctx),
			},
			"network_policy_id": types.StringType,
		},
	}
}

// GetEgress returns the value of the Egress field in AccountNetworkPolicy_SdkV2 as
// a NetworkPolicyEgress_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AccountNetworkPolicy_SdkV2) GetEgress(ctx context.Context) (NetworkPolicyEgress_SdkV2, bool) {
	var e NetworkPolicyEgress_SdkV2
	if m.Egress.IsNull() || m.Egress.IsUnknown() {
		return e, false
	}
	var v []NetworkPolicyEgress_SdkV2
	d := m.Egress.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEgress sets the value of the Egress field in AccountNetworkPolicy_SdkV2.
func (m *AccountNetworkPolicy_SdkV2) SetEgress(ctx context.Context, v NetworkPolicyEgress_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["egress"]
	m.Egress = types.ListValueMust(t, vs)
}

type AibiDashboardEmbeddingAccessPolicy_SdkV2 struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type"`
}

func (to *AibiDashboardEmbeddingAccessPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingAccessPolicy_SdkV2) {
}

func (to *AibiDashboardEmbeddingAccessPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingAccessPolicy_SdkV2) {
}

func (m AibiDashboardEmbeddingAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AibiDashboardEmbeddingAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_policy_type": m.AccessPolicyType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_policy_type": types.StringType,
		},
	}
}

type AibiDashboardEmbeddingAccessPolicySetting_SdkV2 struct {
	AibiDashboardEmbeddingAccessPolicy types.List `tfsdk:"aibi_dashboard_embedding_access_policy"`
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

func (to *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingAccessPolicySetting_SdkV2) {
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

func (to *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingAccessPolicySetting_SdkV2) {
	if !from.AibiDashboardEmbeddingAccessPolicy.IsNull() && !from.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		if toAibiDashboardEmbeddingAccessPolicy, ok := to.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
			if fromAibiDashboardEmbeddingAccessPolicy, ok := from.GetAibiDashboardEmbeddingAccessPolicy(ctx); ok {
				toAibiDashboardEmbeddingAccessPolicy.SyncFieldsDuringRead(ctx, fromAibiDashboardEmbeddingAccessPolicy)
				to.SetAibiDashboardEmbeddingAccessPolicy(ctx, toAibiDashboardEmbeddingAccessPolicy)
			}
		}
	}
}

func (m AibiDashboardEmbeddingAccessPolicySetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].SetRequired()
	attrs["aibi_dashboard_embedding_access_policy"] = attrs["aibi_dashboard_embedding_access_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m AibiDashboardEmbeddingAccessPolicySetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicySetting_SdkV2
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingAccessPolicySetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_access_policy": m.AibiDashboardEmbeddingAccessPolicy,
			"etag":                                   m.Etag,
			"setting_name":                           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingAccessPolicySetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_access_policy": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingAccessPolicy_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetAibiDashboardEmbeddingAccessPolicy returns the value of the AibiDashboardEmbeddingAccessPolicy field in AibiDashboardEmbeddingAccessPolicySetting_SdkV2 as
// a AibiDashboardEmbeddingAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy_SdkV2, bool) {
	var e AibiDashboardEmbeddingAccessPolicy_SdkV2
	if m.AibiDashboardEmbeddingAccessPolicy.IsNull() || m.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicy_SdkV2
	d := m.AibiDashboardEmbeddingAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in AibiDashboardEmbeddingAccessPolicySetting_SdkV2.
func (m *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aibi_dashboard_embedding_access_policy"]
	m.AibiDashboardEmbeddingAccessPolicy = types.ListValueMust(t, vs)
}

type AibiDashboardEmbeddingApprovedDomains_SdkV2 struct {
	ApprovedDomains types.List `tfsdk:"approved_domains"`
}

func (to *AibiDashboardEmbeddingApprovedDomains_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingApprovedDomains_SdkV2) {
	if !from.ApprovedDomains.IsNull() && !from.ApprovedDomains.IsUnknown() && to.ApprovedDomains.IsNull() && len(from.ApprovedDomains.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ApprovedDomains, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ApprovedDomains = from.ApprovedDomains
	}
}

func (to *AibiDashboardEmbeddingApprovedDomains_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingApprovedDomains_SdkV2) {
	if !from.ApprovedDomains.IsNull() && !from.ApprovedDomains.IsUnknown() && to.ApprovedDomains.IsNull() && len(from.ApprovedDomains.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ApprovedDomains, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ApprovedDomains = from.ApprovedDomains
	}
}

func (m AibiDashboardEmbeddingApprovedDomains_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m AibiDashboardEmbeddingApprovedDomains_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomains_SdkV2
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingApprovedDomains_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"approved_domains": m.ApprovedDomains,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingApprovedDomains_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"approved_domains": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetApprovedDomains returns the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AibiDashboardEmbeddingApprovedDomains_SdkV2) GetApprovedDomains(ctx context.Context) ([]types.String, bool) {
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

// SetApprovedDomains sets the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains_SdkV2.
func (m *AibiDashboardEmbeddingApprovedDomains_SdkV2) SetApprovedDomains(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["approved_domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ApprovedDomains = types.ListValueMust(t, vs)
}

type AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2 struct {
	AibiDashboardEmbeddingApprovedDomains types.List `tfsdk:"aibi_dashboard_embedding_approved_domains"`
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

func (to *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) {
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

func (to *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) {
	if !from.AibiDashboardEmbeddingApprovedDomains.IsNull() && !from.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		if toAibiDashboardEmbeddingApprovedDomains, ok := to.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
			if fromAibiDashboardEmbeddingApprovedDomains, ok := from.GetAibiDashboardEmbeddingApprovedDomains(ctx); ok {
				toAibiDashboardEmbeddingApprovedDomains.SyncFieldsDuringRead(ctx, fromAibiDashboardEmbeddingApprovedDomains)
				to.SetAibiDashboardEmbeddingApprovedDomains(ctx, toAibiDashboardEmbeddingApprovedDomains)
			}
		}
	}
}

func (m AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].SetRequired()
	attrs["aibi_dashboard_embedding_approved_domains"] = attrs["aibi_dashboard_embedding_approved_domains"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_approved_domains": m.AibiDashboardEmbeddingApprovedDomains,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aibi_dashboard_embedding_approved_domains": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingApprovedDomains_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetAibiDashboardEmbeddingApprovedDomains returns the value of the AibiDashboardEmbeddingApprovedDomains field in AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2 as
// a AibiDashboardEmbeddingApprovedDomains_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains_SdkV2, bool) {
	var e AibiDashboardEmbeddingApprovedDomains_SdkV2
	if m.AibiDashboardEmbeddingApprovedDomains.IsNull() || m.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomains_SdkV2
	d := m.AibiDashboardEmbeddingApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2.
func (m *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aibi_dashboard_embedding_approved_domains"]
	m.AibiDashboardEmbeddingApprovedDomains = types.ListValueMust(t, vs)
}

type AutomaticClusterUpdateSetting_SdkV2 struct {
	AutomaticClusterUpdateWorkspace types.List `tfsdk:"automatic_cluster_update_workspace"`
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

func (to *AutomaticClusterUpdateSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutomaticClusterUpdateSetting_SdkV2) {
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

func (to *AutomaticClusterUpdateSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AutomaticClusterUpdateSetting_SdkV2) {
	if !from.AutomaticClusterUpdateWorkspace.IsNull() && !from.AutomaticClusterUpdateWorkspace.IsUnknown() {
		if toAutomaticClusterUpdateWorkspace, ok := to.GetAutomaticClusterUpdateWorkspace(ctx); ok {
			if fromAutomaticClusterUpdateWorkspace, ok := from.GetAutomaticClusterUpdateWorkspace(ctx); ok {
				toAutomaticClusterUpdateWorkspace.SyncFieldsDuringRead(ctx, fromAutomaticClusterUpdateWorkspace)
				to.SetAutomaticClusterUpdateWorkspace(ctx, toAutomaticClusterUpdateWorkspace)
			}
		}
	}
}

func (m AutomaticClusterUpdateSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].SetRequired()
	attrs["automatic_cluster_update_workspace"] = attrs["automatic_cluster_update_workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m AutomaticClusterUpdateSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"automatic_cluster_update_workspace": reflect.TypeOf(ClusterAutoRestartMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutomaticClusterUpdateSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m AutomaticClusterUpdateSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"automatic_cluster_update_workspace": m.AutomaticClusterUpdateWorkspace,
			"etag":                               m.Etag,
			"setting_name":                       m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutomaticClusterUpdateSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"automatic_cluster_update_workspace": basetypes.ListType{
				ElemType: ClusterAutoRestartMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetAutomaticClusterUpdateWorkspace returns the value of the AutomaticClusterUpdateWorkspace field in AutomaticClusterUpdateSetting_SdkV2 as
// a ClusterAutoRestartMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutomaticClusterUpdateSetting_SdkV2) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage_SdkV2, bool) {
	var e ClusterAutoRestartMessage_SdkV2
	if m.AutomaticClusterUpdateWorkspace.IsNull() || m.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessage_SdkV2
	d := m.AutomaticClusterUpdateWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in AutomaticClusterUpdateSetting_SdkV2.
func (m *AutomaticClusterUpdateSetting_SdkV2) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["automatic_cluster_update_workspace"]
	m.AutomaticClusterUpdateWorkspace = types.ListValueMust(t, vs)
}

type BooleanMessage_SdkV2 struct {
	Value types.Bool `tfsdk:"value"`
}

func (to *BooleanMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BooleanMessage_SdkV2) {
}

func (to *BooleanMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BooleanMessage_SdkV2) {
}

func (m BooleanMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m BooleanMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage_SdkV2
// only implements ToObjectValue() and Type().
func (m BooleanMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BooleanMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type ClusterAutoRestartMessage_SdkV2 struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`

	EnablementDetails types.List `tfsdk:"enablement_details"`

	MaintenanceWindow types.List `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

func (to *ClusterAutoRestartMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessage_SdkV2) {
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

func (to *ClusterAutoRestartMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessage_SdkV2) {
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

func (m ClusterAutoRestartMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["can_toggle"] = attrs["can_toggle"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["enablement_details"] = attrs["enablement_details"].SetOptional()
	attrs["enablement_details"] = attrs["enablement_details"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["maintenance_window"] = attrs["maintenance_window"].SetOptional()
	attrs["maintenance_window"] = attrs["maintenance_window"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m ClusterAutoRestartMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails_SdkV2{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessage_SdkV2
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ClusterAutoRestartMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"can_toggle": types.BoolType,
			"enabled":    types.BoolType,
			"enablement_details": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageEnablementDetails_SdkV2{}.Type(ctx),
			},
			"maintenance_window": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindow_SdkV2{}.Type(ctx),
			},
			"restart_even_if_no_updates_available": types.BoolType,
		},
	}
}

// GetEnablementDetails returns the value of the EnablementDetails field in ClusterAutoRestartMessage_SdkV2 as
// a ClusterAutoRestartMessageEnablementDetails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessage_SdkV2) GetEnablementDetails(ctx context.Context) (ClusterAutoRestartMessageEnablementDetails_SdkV2, bool) {
	var e ClusterAutoRestartMessageEnablementDetails_SdkV2
	if m.EnablementDetails.IsNull() || m.EnablementDetails.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageEnablementDetails_SdkV2
	d := m.EnablementDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnablementDetails sets the value of the EnablementDetails field in ClusterAutoRestartMessage_SdkV2.
func (m *ClusterAutoRestartMessage_SdkV2) SetEnablementDetails(ctx context.Context, v ClusterAutoRestartMessageEnablementDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["enablement_details"]
	m.EnablementDetails = types.ListValueMust(t, vs)
}

// GetMaintenanceWindow returns the value of the MaintenanceWindow field in ClusterAutoRestartMessage_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessage_SdkV2) GetMaintenanceWindow(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindow_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindow_SdkV2
	if m.MaintenanceWindow.IsNull() || m.MaintenanceWindow.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindow_SdkV2
	d := m.MaintenanceWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaintenanceWindow sets the value of the MaintenanceWindow field in ClusterAutoRestartMessage_SdkV2.
func (m *ClusterAutoRestartMessage_SdkV2) SetMaintenanceWindow(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["maintenance_window"]
	m.MaintenanceWindow = types.ListValueMust(t, vs)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails_SdkV2 struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode types.Bool `tfsdk:"forced_for_compliance_mode"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement types.Bool `tfsdk:"unavailable_for_disabled_entitlement"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier types.Bool `tfsdk:"unavailable_for_non_enterprise_tier"`
}

func (to *ClusterAutoRestartMessageEnablementDetails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageEnablementDetails_SdkV2) {
}

func (to *ClusterAutoRestartMessageEnablementDetails_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageEnablementDetails_SdkV2) {
}

func (m ClusterAutoRestartMessageEnablementDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ClusterAutoRestartMessageEnablementDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageEnablementDetails_SdkV2
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageEnablementDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"forced_for_compliance_mode":           m.ForcedForComplianceMode,
			"unavailable_for_disabled_entitlement": m.UnavailableForDisabledEntitlement,
			"unavailable_for_non_enterprise_tier":  m.UnavailableForNonEnterpriseTier,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageEnablementDetails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"forced_for_compliance_mode":           types.BoolType,
			"unavailable_for_disabled_entitlement": types.BoolType,
			"unavailable_for_non_enterprise_tier":  types.BoolType,
		},
	}
}

type ClusterAutoRestartMessageMaintenanceWindow_SdkV2 struct {
	WeekDayBasedSchedule types.List `tfsdk:"week_day_based_schedule"`
}

func (to *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
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

func (to *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
	if !from.WeekDayBasedSchedule.IsNull() && !from.WeekDayBasedSchedule.IsUnknown() {
		if toWeekDayBasedSchedule, ok := to.GetWeekDayBasedSchedule(ctx); ok {
			if fromWeekDayBasedSchedule, ok := from.GetWeekDayBasedSchedule(ctx); ok {
				toWeekDayBasedSchedule.SyncFieldsDuringRead(ctx, fromWeekDayBasedSchedule)
				to.SetWeekDayBasedSchedule(ctx, toWeekDayBasedSchedule)
			}
		}
	}
}

func (m ClusterAutoRestartMessageMaintenanceWindow_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["week_day_based_schedule"] = attrs["week_day_based_schedule"].SetOptional()
	attrs["week_day_based_schedule"] = attrs["week_day_based_schedule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindow.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessageMaintenanceWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"week_day_based_schedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindow_SdkV2
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageMaintenanceWindow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"week_day_based_schedule": m.WeekDayBasedSchedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageMaintenanceWindow_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"week_day_based_schedule": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWeekDayBasedSchedule returns the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) GetWeekDayBasedSchedule(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
	if m.WeekDayBasedSchedule.IsNull() || m.WeekDayBasedSchedule.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
	d := m.WeekDayBasedSchedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWeekDayBasedSchedule sets the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow_SdkV2.
func (m *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SetWeekDayBasedSchedule(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["week_day_based_schedule"]
	m.WeekDayBasedSchedule = types.ListValueMust(t, vs)
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`

	Frequency types.String `tfsdk:"frequency"`

	WindowStartTime types.List `tfsdk:"window_start_time"`
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
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

func (to *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
	if !from.WindowStartTime.IsNull() && !from.WindowStartTime.IsUnknown() {
		if toWindowStartTime, ok := to.GetWindowStartTime(ctx); ok {
			if fromWindowStartTime, ok := from.GetWindowStartTime(ctx); ok {
				toWindowStartTime.SyncFieldsDuringRead(ctx, fromWindowStartTime)
				to.SetWindowStartTime(ctx, toWindowStartTime)
			}
		}
	}
}

func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["day_of_week"] = attrs["day_of_week"].SetOptional()
	attrs["frequency"] = attrs["frequency"].SetOptional()
	attrs["window_start_time"] = attrs["window_start_time"].SetOptional()
	attrs["window_start_time"] = attrs["window_start_time"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"window_start_time": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"day_of_week":       m.DayOfWeek,
			"frequency":         m.Frequency,
			"window_start_time": m.WindowStartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"day_of_week": types.StringType,
			"frequency":   types.StringType,
			"window_start_time": basetypes.ListType{
				ElemType: ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWindowStartTime returns the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) GetWindowStartTime(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
	if m.WindowStartTime.IsNull() || m.WindowStartTime.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
	d := m.WindowStartTime.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWindowStartTime sets the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2.
func (m *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SetWindowStartTime(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["window_start_time"]
	m.WindowStartTime = types.ListValueMust(t, vs)
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2 struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
}

func (to *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
}

func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
// only implements ToObjectValue() and Type().
func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hours":   m.Hours,
			"minutes": m.Minutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"hours":   types.Int64Type,
			"minutes": types.Int64Type,
		},
	}
}

// SHIELD feature: CSP
type ComplianceSecurityProfile_SdkV2 struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards types.List `tfsdk:"compliance_standards"`

	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (to *ComplianceSecurityProfile_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComplianceSecurityProfile_SdkV2) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (to *ComplianceSecurityProfile_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ComplianceSecurityProfile_SdkV2) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (m ComplianceSecurityProfile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ComplianceSecurityProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile_SdkV2
// only implements ToObjectValue() and Type().
func (m ComplianceSecurityProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": m.ComplianceStandards,
			"is_enabled":           m.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComplianceSecurityProfile_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enabled": types.BoolType,
		},
	}
}

// GetComplianceStandards returns the value of the ComplianceStandards field in ComplianceSecurityProfile_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ComplianceSecurityProfile_SdkV2) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
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

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile_SdkV2.
func (m *ComplianceSecurityProfile_SdkV2) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ComplianceStandards = types.ListValueMust(t, vs)
}

type ComplianceSecurityProfileSetting_SdkV2 struct {
	ComplianceSecurityProfileWorkspace types.List `tfsdk:"compliance_security_profile_workspace"`
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

func (to *ComplianceSecurityProfileSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ComplianceSecurityProfileSetting_SdkV2) {
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

func (to *ComplianceSecurityProfileSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ComplianceSecurityProfileSetting_SdkV2) {
	if !from.ComplianceSecurityProfileWorkspace.IsNull() && !from.ComplianceSecurityProfileWorkspace.IsUnknown() {
		if toComplianceSecurityProfileWorkspace, ok := to.GetComplianceSecurityProfileWorkspace(ctx); ok {
			if fromComplianceSecurityProfileWorkspace, ok := from.GetComplianceSecurityProfileWorkspace(ctx); ok {
				toComplianceSecurityProfileWorkspace.SyncFieldsDuringRead(ctx, fromComplianceSecurityProfileWorkspace)
				to.SetComplianceSecurityProfileWorkspace(ctx, toComplianceSecurityProfileWorkspace)
			}
		}
	}
}

func (m ComplianceSecurityProfileSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["compliance_security_profile_workspace"] = attrs["compliance_security_profile_workspace"].SetRequired()
	attrs["compliance_security_profile_workspace"] = attrs["compliance_security_profile_workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m ComplianceSecurityProfileSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_security_profile_workspace": reflect.TypeOf(ComplianceSecurityProfile_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfileSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m ComplianceSecurityProfileSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_security_profile_workspace": m.ComplianceSecurityProfileWorkspace,
			"etag":                                  m.Etag,
			"setting_name":                          m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ComplianceSecurityProfileSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_security_profile_workspace": basetypes.ListType{
				ElemType: ComplianceSecurityProfile_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetComplianceSecurityProfileWorkspace returns the value of the ComplianceSecurityProfileWorkspace field in ComplianceSecurityProfileSetting_SdkV2 as
// a ComplianceSecurityProfile_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ComplianceSecurityProfileSetting_SdkV2) GetComplianceSecurityProfileWorkspace(ctx context.Context) (ComplianceSecurityProfile_SdkV2, bool) {
	var e ComplianceSecurityProfile_SdkV2
	if m.ComplianceSecurityProfileWorkspace.IsNull() || m.ComplianceSecurityProfileWorkspace.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfile_SdkV2
	d := m.ComplianceSecurityProfileWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComplianceSecurityProfileWorkspace sets the value of the ComplianceSecurityProfileWorkspace field in ComplianceSecurityProfileSetting_SdkV2.
func (m *ComplianceSecurityProfileSetting_SdkV2) SetComplianceSecurityProfileWorkspace(ctx context.Context, v ComplianceSecurityProfile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_security_profile_workspace"]
	m.ComplianceSecurityProfileWorkspace = types.ListValueMust(t, vs)
}

type Config_SdkV2 struct {
	Email types.List `tfsdk:"email"`

	GenericWebhook types.List `tfsdk:"generic_webhook"`

	MicrosoftTeams types.List `tfsdk:"microsoft_teams"`

	Pagerduty types.List `tfsdk:"pagerduty"`

	Slack types.List `tfsdk:"slack"`
}

func (to *Config_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Config_SdkV2) {
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

func (to *Config_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Config_SdkV2) {
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

func (m Config_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email"] = attrs["email"].SetOptional()
	attrs["email"] = attrs["email"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["generic_webhook"] = attrs["generic_webhook"].SetOptional()
	attrs["generic_webhook"] = attrs["generic_webhook"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["microsoft_teams"] = attrs["microsoft_teams"].SetOptional()
	attrs["microsoft_teams"] = attrs["microsoft_teams"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["pagerduty"] = attrs["pagerduty"].SetOptional()
	attrs["pagerduty"] = attrs["pagerduty"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["slack"] = attrs["slack"].SetOptional()
	attrs["slack"] = attrs["slack"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Config.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Config_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email":           reflect.TypeOf(EmailConfig_SdkV2{}),
		"generic_webhook": reflect.TypeOf(GenericWebhookConfig_SdkV2{}),
		"microsoft_teams": reflect.TypeOf(MicrosoftTeamsConfig_SdkV2{}),
		"pagerduty":       reflect.TypeOf(PagerdutyConfig_SdkV2{}),
		"slack":           reflect.TypeOf(SlackConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Config_SdkV2
// only implements ToObjectValue() and Type().
func (m Config_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m Config_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email": basetypes.ListType{
				ElemType: EmailConfig_SdkV2{}.Type(ctx),
			},
			"generic_webhook": basetypes.ListType{
				ElemType: GenericWebhookConfig_SdkV2{}.Type(ctx),
			},
			"microsoft_teams": basetypes.ListType{
				ElemType: MicrosoftTeamsConfig_SdkV2{}.Type(ctx),
			},
			"pagerduty": basetypes.ListType{
				ElemType: PagerdutyConfig_SdkV2{}.Type(ctx),
			},
			"slack": basetypes.ListType{
				ElemType: SlackConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetEmail returns the value of the Email field in Config_SdkV2 as
// a EmailConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config_SdkV2) GetEmail(ctx context.Context) (EmailConfig_SdkV2, bool) {
	var e EmailConfig_SdkV2
	if m.Email.IsNull() || m.Email.IsUnknown() {
		return e, false
	}
	var v []EmailConfig_SdkV2
	d := m.Email.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmail sets the value of the Email field in Config_SdkV2.
func (m *Config_SdkV2) SetEmail(ctx context.Context, v EmailConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email"]
	m.Email = types.ListValueMust(t, vs)
}

// GetGenericWebhook returns the value of the GenericWebhook field in Config_SdkV2 as
// a GenericWebhookConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config_SdkV2) GetGenericWebhook(ctx context.Context) (GenericWebhookConfig_SdkV2, bool) {
	var e GenericWebhookConfig_SdkV2
	if m.GenericWebhook.IsNull() || m.GenericWebhook.IsUnknown() {
		return e, false
	}
	var v []GenericWebhookConfig_SdkV2
	d := m.GenericWebhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGenericWebhook sets the value of the GenericWebhook field in Config_SdkV2.
func (m *Config_SdkV2) SetGenericWebhook(ctx context.Context, v GenericWebhookConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["generic_webhook"]
	m.GenericWebhook = types.ListValueMust(t, vs)
}

// GetMicrosoftTeams returns the value of the MicrosoftTeams field in Config_SdkV2 as
// a MicrosoftTeamsConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config_SdkV2) GetMicrosoftTeams(ctx context.Context) (MicrosoftTeamsConfig_SdkV2, bool) {
	var e MicrosoftTeamsConfig_SdkV2
	if m.MicrosoftTeams.IsNull() || m.MicrosoftTeams.IsUnknown() {
		return e, false
	}
	var v []MicrosoftTeamsConfig_SdkV2
	d := m.MicrosoftTeams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMicrosoftTeams sets the value of the MicrosoftTeams field in Config_SdkV2.
func (m *Config_SdkV2) SetMicrosoftTeams(ctx context.Context, v MicrosoftTeamsConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["microsoft_teams"]
	m.MicrosoftTeams = types.ListValueMust(t, vs)
}

// GetPagerduty returns the value of the Pagerduty field in Config_SdkV2 as
// a PagerdutyConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config_SdkV2) GetPagerduty(ctx context.Context) (PagerdutyConfig_SdkV2, bool) {
	var e PagerdutyConfig_SdkV2
	if m.Pagerduty.IsNull() || m.Pagerduty.IsUnknown() {
		return e, false
	}
	var v []PagerdutyConfig_SdkV2
	d := m.Pagerduty.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPagerduty sets the value of the Pagerduty field in Config_SdkV2.
func (m *Config_SdkV2) SetPagerduty(ctx context.Context, v PagerdutyConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pagerduty"]
	m.Pagerduty = types.ListValueMust(t, vs)
}

// GetSlack returns the value of the Slack field in Config_SdkV2 as
// a SlackConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *Config_SdkV2) GetSlack(ctx context.Context) (SlackConfig_SdkV2, bool) {
	var e SlackConfig_SdkV2
	if m.Slack.IsNull() || m.Slack.IsUnknown() {
		return e, false
	}
	var v []SlackConfig_SdkV2
	d := m.Slack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSlack sets the value of the Slack field in Config_SdkV2.
func (m *Config_SdkV2) SetSlack(ctx context.Context, v SlackConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["slack"]
	m.Slack = types.ListValueMust(t, vs)
}

// Details required to configure a block list or allow list.
type CreateIpAccessList_SdkV2 struct {
	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`

	ListType types.String `tfsdk:"list_type"`
}

func (to *CreateIpAccessList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateIpAccessList_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *CreateIpAccessList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateIpAccessList_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m CreateIpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetRequired()
	attrs["list_type"] = attrs["list_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateIpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateIpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_addresses": m.IpAddresses,
			"label":        m.Label,
			"list_type":    m.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateIpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetIpAddresses returns the value of the IpAddresses field in CreateIpAccessList_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateIpAccessList_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in CreateIpAccessList_SdkV2.
func (m *CreateIpAccessList_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

// An IP access list was successfully created.
type CreateIpAccessListResponse_SdkV2 struct {
	IpAccessList types.List `tfsdk:"ip_access_list"`
}

func (to *CreateIpAccessListResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateIpAccessListResponse_SdkV2) {
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

func (to *CreateIpAccessListResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateIpAccessListResponse_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (m CreateIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": m.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessListInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in CreateIpAccessListResponse_SdkV2 as
// a IpAccessListInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateIpAccessListResponse_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessListInfo_SdkV2, bool) {
	var e IpAccessListInfo_SdkV2
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo_SdkV2
	d := m.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in CreateIpAccessListResponse_SdkV2.
func (m *CreateIpAccessListResponse_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessListInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	m.IpAccessList = types.ListValueMust(t, vs)
}

type CreateNetworkConnectivityConfigRequest_SdkV2 struct {
	NetworkConnectivityConfig types.List `tfsdk:"network_connectivity_config"`
}

func (to *CreateNetworkConnectivityConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkConnectivityConfigRequest_SdkV2) {
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

func (to *CreateNetworkConnectivityConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkConnectivityConfigRequest_SdkV2) {
	if !from.NetworkConnectivityConfig.IsNull() && !from.NetworkConnectivityConfig.IsUnknown() {
		if toNetworkConnectivityConfig, ok := to.GetNetworkConnectivityConfig(ctx); ok {
			if fromNetworkConnectivityConfig, ok := from.GetNetworkConnectivityConfig(ctx); ok {
				toNetworkConnectivityConfig.SyncFieldsDuringRead(ctx, fromNetworkConnectivityConfig)
				to.SetNetworkConnectivityConfig(ctx, toNetworkConnectivityConfig)
			}
		}
	}
}

func (m CreateNetworkConnectivityConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_connectivity_config"] = attrs["network_connectivity_config"].SetRequired()
	attrs["network_connectivity_config"] = attrs["network_connectivity_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkConnectivityConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNetworkConnectivityConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_connectivity_config": reflect.TypeOf(CreateNetworkConnectivityConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateNetworkConnectivityConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config": m.NetworkConnectivityConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkConnectivityConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config": basetypes.ListType{
				ElemType: CreateNetworkConnectivityConfiguration_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNetworkConnectivityConfig returns the value of the NetworkConnectivityConfig field in CreateNetworkConnectivityConfigRequest_SdkV2 as
// a CreateNetworkConnectivityConfiguration_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkConnectivityConfigRequest_SdkV2) GetNetworkConnectivityConfig(ctx context.Context) (CreateNetworkConnectivityConfiguration_SdkV2, bool) {
	var e CreateNetworkConnectivityConfiguration_SdkV2
	if m.NetworkConnectivityConfig.IsNull() || m.NetworkConnectivityConfig.IsUnknown() {
		return e, false
	}
	var v []CreateNetworkConnectivityConfiguration_SdkV2
	d := m.NetworkConnectivityConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNetworkConnectivityConfig sets the value of the NetworkConnectivityConfig field in CreateNetworkConnectivityConfigRequest_SdkV2.
func (m *CreateNetworkConnectivityConfigRequest_SdkV2) SetNetworkConnectivityConfig(ctx context.Context, v CreateNetworkConnectivityConfiguration_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["network_connectivity_config"]
	m.NetworkConnectivityConfig = types.ListValueMust(t, vs)
}

// Properties of the new network connectivity configuration.
type CreateNetworkConnectivityConfiguration_SdkV2 struct {
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

func (to *CreateNetworkConnectivityConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkConnectivityConfiguration_SdkV2) {
}

func (to *CreateNetworkConnectivityConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkConnectivityConfiguration_SdkV2) {
}

func (m CreateNetworkConnectivityConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateNetworkConnectivityConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateNetworkConnectivityConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   m.Name,
			"region": m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkConnectivityConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"region": types.StringType,
		},
	}
}

type CreateNetworkPolicyRequest_SdkV2 struct {
	// Network policy configuration details.
	NetworkPolicy types.List `tfsdk:"network_policy"`
}

func (to *CreateNetworkPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNetworkPolicyRequest_SdkV2) {
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

func (to *CreateNetworkPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateNetworkPolicyRequest_SdkV2) {
	if !from.NetworkPolicy.IsNull() && !from.NetworkPolicy.IsUnknown() {
		if toNetworkPolicy, ok := to.GetNetworkPolicy(ctx); ok {
			if fromNetworkPolicy, ok := from.GetNetworkPolicy(ctx); ok {
				toNetworkPolicy.SyncFieldsDuringRead(ctx, fromNetworkPolicy)
				to.SetNetworkPolicy(ctx, toNetworkPolicy)
			}
		}
	}
}

func (m CreateNetworkPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_policy"] = attrs["network_policy"].SetRequired()
	attrs["network_policy"] = attrs["network_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateNetworkPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_policy": reflect.TypeOf(AccountNetworkPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateNetworkPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy": m.NetworkPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNetworkPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy": basetypes.ListType{
				ElemType: AccountNetworkPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNetworkPolicy returns the value of the NetworkPolicy field in CreateNetworkPolicyRequest_SdkV2 as
// a AccountNetworkPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNetworkPolicyRequest_SdkV2) GetNetworkPolicy(ctx context.Context) (AccountNetworkPolicy_SdkV2, bool) {
	var e AccountNetworkPolicy_SdkV2
	if m.NetworkPolicy.IsNull() || m.NetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []AccountNetworkPolicy_SdkV2
	d := m.NetworkPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNetworkPolicy sets the value of the NetworkPolicy field in CreateNetworkPolicyRequest_SdkV2.
func (m *CreateNetworkPolicyRequest_SdkV2) SetNetworkPolicy(ctx context.Context, v AccountNetworkPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["network_policy"]
	m.NetworkPolicy = types.ListValueMust(t, vs)
}

type CreateNotificationDestinationRequest_SdkV2 struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.List `tfsdk:"config"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
}

func (to *CreateNotificationDestinationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateNotificationDestinationRequest_SdkV2) {
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

func (to *CreateNotificationDestinationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateNotificationDestinationRequest_SdkV2) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m CreateNotificationDestinationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       m.Config,
			"display_name": m.DisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: Config_SdkV2{}.Type(ctx),
			},
			"display_name": types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in CreateNotificationDestinationRequest_SdkV2 as
// a Config_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateNotificationDestinationRequest_SdkV2) GetConfig(ctx context.Context) (Config_SdkV2, bool) {
	var e Config_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []Config_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in CreateNotificationDestinationRequest_SdkV2.
func (m *CreateNotificationDestinationRequest_SdkV2) SetConfig(ctx context.Context, v Config_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest_SdkV2 struct {
	// Application ID of the service principal.
	ApplicationId types.String `tfsdk:"application_id"`
	// Comment that describes the purpose of the token.
	Comment types.String `tfsdk:"comment"`
	// The number of seconds before the token expires.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

func (to *CreateOboTokenRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOboTokenRequest_SdkV2) {
}

func (to *CreateOboTokenRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateOboTokenRequest_SdkV2) {
}

func (m CreateOboTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateOboTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateOboTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"application_id":   m.ApplicationId,
			"comment":          m.Comment,
			"lifetime_seconds": m.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOboTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"application_id":   types.StringType,
			"comment":          types.StringType,
			"lifetime_seconds": types.Int64Type,
		},
	}
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse_SdkV2 struct {
	TokenInfo types.List `tfsdk:"token_info"`
	// Value of the token.
	TokenValue types.String `tfsdk:"token_value"`
}

func (to *CreateOboTokenResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateOboTokenResponse_SdkV2) {
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

func (to *CreateOboTokenResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateOboTokenResponse_SdkV2) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				toTokenInfo.SyncFieldsDuringRead(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (m CreateOboTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_info"] = attrs["token_info"].SetOptional()
	attrs["token_info"] = attrs["token_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateOboTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateOboTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  m.TokenInfo,
			"token_value": m.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateOboTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": basetypes.ListType{
				ElemType: TokenInfo_SdkV2{}.Type(ctx),
			},
			"token_value": types.StringType,
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in CreateOboTokenResponse_SdkV2 as
// a TokenInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateOboTokenResponse_SdkV2) GetTokenInfo(ctx context.Context) (TokenInfo_SdkV2, bool) {
	var e TokenInfo_SdkV2
	if m.TokenInfo.IsNull() || m.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []TokenInfo_SdkV2
	d := m.TokenInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenInfo sets the value of the TokenInfo field in CreateOboTokenResponse_SdkV2.
func (m *CreateOboTokenResponse_SdkV2) SetTokenInfo(ctx context.Context, v TokenInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_info"]
	m.TokenInfo = types.ListValueMust(t, vs)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type CreatePrivateEndpointRule_SdkV2 struct {
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

func (to *CreatePrivateEndpointRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateEndpointRule_SdkV2) {
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

func (to *CreatePrivateEndpointRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateEndpointRule_SdkV2) {
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

func (m CreatePrivateEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreatePrivateEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (m CreatePrivateEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m CreatePrivateEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetDomainNames returns the value of the DomainNames field in CreatePrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateEndpointRule_SdkV2) GetDomainNames(ctx context.Context) ([]types.String, bool) {
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

// SetDomainNames sets the value of the DomainNames field in CreatePrivateEndpointRule_SdkV2.
func (m *CreatePrivateEndpointRule_SdkV2) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in CreatePrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateEndpointRule_SdkV2) GetResourceNames(ctx context.Context) ([]types.String, bool) {
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

// SetResourceNames sets the value of the ResourceNames field in CreatePrivateEndpointRule_SdkV2.
func (m *CreatePrivateEndpointRule_SdkV2) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

type CreatePrivateEndpointRuleRequest_SdkV2 struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`

	PrivateEndpointRule types.List `tfsdk:"private_endpoint_rule"`
}

func (to *CreatePrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePrivateEndpointRuleRequest_SdkV2) {
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

func (to *CreatePrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePrivateEndpointRuleRequest_SdkV2) {
	if !from.PrivateEndpointRule.IsNull() && !from.PrivateEndpointRule.IsUnknown() {
		if toPrivateEndpointRule, ok := to.GetPrivateEndpointRule(ctx); ok {
			if fromPrivateEndpointRule, ok := from.GetPrivateEndpointRule(ctx); ok {
				toPrivateEndpointRule.SyncFieldsDuringRead(ctx, fromPrivateEndpointRule)
				to.SetPrivateEndpointRule(ctx, toPrivateEndpointRule)
			}
		}
	}
}

func (m CreatePrivateEndpointRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_endpoint_rule"] = attrs["private_endpoint_rule"].SetRequired()
	attrs["private_endpoint_rule"] = attrs["private_endpoint_rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreatePrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_endpoint_rule": reflect.TypeOf(CreatePrivateEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreatePrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule":          m.PrivateEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule": basetypes.ListType{
				ElemType: CreatePrivateEndpointRule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPrivateEndpointRule returns the value of the PrivateEndpointRule field in CreatePrivateEndpointRuleRequest_SdkV2 as
// a CreatePrivateEndpointRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePrivateEndpointRuleRequest_SdkV2) GetPrivateEndpointRule(ctx context.Context) (CreatePrivateEndpointRule_SdkV2, bool) {
	var e CreatePrivateEndpointRule_SdkV2
	if m.PrivateEndpointRule.IsNull() || m.PrivateEndpointRule.IsUnknown() {
		return e, false
	}
	var v []CreatePrivateEndpointRule_SdkV2
	d := m.PrivateEndpointRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPrivateEndpointRule sets the value of the PrivateEndpointRule field in CreatePrivateEndpointRuleRequest_SdkV2.
func (m *CreatePrivateEndpointRuleRequest_SdkV2) SetPrivateEndpointRule(ctx context.Context, v CreatePrivateEndpointRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["private_endpoint_rule"]
	m.PrivateEndpointRule = types.ListValueMust(t, vs)
}

type CreateTokenRequest_SdkV2 struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

func (to *CreateTokenRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTokenRequest_SdkV2) {
}

func (to *CreateTokenRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateTokenRequest_SdkV2) {
}

func (m CreateTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CreateTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":          m.Comment,
			"lifetime_seconds": m.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"comment":          types.StringType,
			"lifetime_seconds": types.Int64Type,
		},
	}
}

type CreateTokenResponse_SdkV2 struct {
	// The information for the new token.
	TokenInfo types.List `tfsdk:"token_info"`
	// The value of the new token.
	TokenValue types.String `tfsdk:"token_value"`
}

func (to *CreateTokenResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateTokenResponse_SdkV2) {
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

func (to *CreateTokenResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateTokenResponse_SdkV2) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				toTokenInfo.SyncFieldsDuringRead(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (m CreateTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_info"] = attrs["token_info"].SetOptional()
	attrs["token_info"] = attrs["token_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CreateTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(PublicTokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  m.TokenInfo,
			"token_value": m.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": basetypes.ListType{
				ElemType: PublicTokenInfo_SdkV2{}.Type(ctx),
			},
			"token_value": types.StringType,
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in CreateTokenResponse_SdkV2 as
// a PublicTokenInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateTokenResponse_SdkV2) GetTokenInfo(ctx context.Context) (PublicTokenInfo_SdkV2, bool) {
	var e PublicTokenInfo_SdkV2
	if m.TokenInfo.IsNull() || m.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []PublicTokenInfo_SdkV2
	d := m.TokenInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenInfo sets the value of the TokenInfo field in CreateTokenResponse_SdkV2.
func (m *CreateTokenResponse_SdkV2) SetTokenInfo(ctx context.Context, v PublicTokenInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_info"]
	m.TokenInfo = types.ListValueMust(t, vs)
}

// Account level policy for CSP
type CspEnablementAccount_SdkV2 struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (to *CspEnablementAccount_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CspEnablementAccount_SdkV2) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (to *CspEnablementAccount_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CspEnablementAccount_SdkV2) {
	if !from.ComplianceStandards.IsNull() && !from.ComplianceStandards.IsUnknown() && to.ComplianceStandards.IsNull() && len(from.ComplianceStandards.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ComplianceStandards, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ComplianceStandards = from.ComplianceStandards
	}
}

func (m CspEnablementAccount_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CspEnablementAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccount_SdkV2
// only implements ToObjectValue() and Type().
func (m CspEnablementAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": m.ComplianceStandards,
			"is_enforced":          m.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CspEnablementAccount_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"compliance_standards": basetypes.ListType{
				ElemType: types.StringType,
			},
			"is_enforced": types.BoolType,
		},
	}
}

// GetComplianceStandards returns the value of the ComplianceStandards field in CspEnablementAccount_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CspEnablementAccount_SdkV2) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
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

// SetComplianceStandards sets the value of the ComplianceStandards field in CspEnablementAccount_SdkV2.
func (m *CspEnablementAccount_SdkV2) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ComplianceStandards = types.ListValueMust(t, vs)
}

type CspEnablementAccountSetting_SdkV2 struct {
	CspEnablementAccount types.List `tfsdk:"csp_enablement_account"`
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

func (to *CspEnablementAccountSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CspEnablementAccountSetting_SdkV2) {
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

func (to *CspEnablementAccountSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CspEnablementAccountSetting_SdkV2) {
	if !from.CspEnablementAccount.IsNull() && !from.CspEnablementAccount.IsUnknown() {
		if toCspEnablementAccount, ok := to.GetCspEnablementAccount(ctx); ok {
			if fromCspEnablementAccount, ok := from.GetCspEnablementAccount(ctx); ok {
				toCspEnablementAccount.SyncFieldsDuringRead(ctx, fromCspEnablementAccount)
				to.SetCspEnablementAccount(ctx, toCspEnablementAccount)
			}
		}
	}
}

func (m CspEnablementAccountSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["csp_enablement_account"] = attrs["csp_enablement_account"].SetRequired()
	attrs["csp_enablement_account"] = attrs["csp_enablement_account"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m CspEnablementAccountSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"csp_enablement_account": reflect.TypeOf(CspEnablementAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccountSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m CspEnablementAccountSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"csp_enablement_account": m.CspEnablementAccount,
			"etag":                   m.Etag,
			"setting_name":           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CspEnablementAccountSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"csp_enablement_account": basetypes.ListType{
				ElemType: CspEnablementAccount_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetCspEnablementAccount returns the value of the CspEnablementAccount field in CspEnablementAccountSetting_SdkV2 as
// a CspEnablementAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CspEnablementAccountSetting_SdkV2) GetCspEnablementAccount(ctx context.Context) (CspEnablementAccount_SdkV2, bool) {
	var e CspEnablementAccount_SdkV2
	if m.CspEnablementAccount.IsNull() || m.CspEnablementAccount.IsUnknown() {
		return e, false
	}
	var v []CspEnablementAccount_SdkV2
	d := m.CspEnablementAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCspEnablementAccount sets the value of the CspEnablementAccount field in CspEnablementAccountSetting_SdkV2.
func (m *CspEnablementAccountSetting_SdkV2) SetCspEnablementAccount(ctx context.Context, v CspEnablementAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["csp_enablement_account"]
	m.CspEnablementAccount = types.ListValueMust(t, vs)
}

// Properties of the new private endpoint rule. Note that for private endpoints
// towards a VPC endpoint service behind a customer-managed NLB, you must
// approve the endpoint in AWS console after initialization.
type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2 struct {
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

func (to *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) {
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

func (to *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) {
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

func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
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
func (m CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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

// GetDomainNames returns the value of the DomainNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) GetDomainNames(ctx context.Context) ([]types.String, bool) {
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

// SetDomainNames sets the value of the DomainNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) GetResourceNames(ctx context.Context) ([]types.String, bool) {
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

// SetResourceNames sets the value of the ResourceNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2.
func (m *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

type DashboardEmailSubscriptions_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
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

func (to *DashboardEmailSubscriptions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DashboardEmailSubscriptions_SdkV2) {
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

func (to *DashboardEmailSubscriptions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DashboardEmailSubscriptions_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m DashboardEmailSubscriptions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m DashboardEmailSubscriptions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardEmailSubscriptions_SdkV2
// only implements ToObjectValue() and Type().
func (m DashboardEmailSubscriptions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DashboardEmailSubscriptions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in DashboardEmailSubscriptions_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DashboardEmailSubscriptions_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in DashboardEmailSubscriptions_SdkV2.
func (m *DashboardEmailSubscriptions_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
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
type DefaultNamespaceSetting_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	Namespace types.List `tfsdk:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *DefaultNamespaceSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultNamespaceSetting_SdkV2) {
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

func (to *DefaultNamespaceSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DefaultNamespaceSetting_SdkV2) {
	if !from.Namespace.IsNull() && !from.Namespace.IsUnknown() {
		if toNamespace, ok := to.GetNamespace(ctx); ok {
			if fromNamespace, ok := from.GetNamespace(ctx); ok {
				toNamespace.SyncFieldsDuringRead(ctx, fromNamespace)
				to.SetNamespace(ctx, toNamespace)
			}
		}
	}
}

func (m DefaultNamespaceSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["namespace"] = attrs["namespace"].SetRequired()
	attrs["namespace"] = attrs["namespace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m DefaultNamespaceSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"namespace": reflect.TypeOf(StringMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultNamespaceSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m DefaultNamespaceSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         m.Etag,
			"namespace":    m.Namespace,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultNamespaceSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"namespace": basetypes.ListType{
				ElemType: StringMessage_SdkV2{}.Type(ctx),
			},
			"setting_name": types.StringType,
		},
	}
}

// GetNamespace returns the value of the Namespace field in DefaultNamespaceSetting_SdkV2 as
// a StringMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultNamespaceSetting_SdkV2) GetNamespace(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if m.Namespace.IsNull() || m.Namespace.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := m.Namespace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNamespace sets the value of the Namespace field in DefaultNamespaceSetting_SdkV2.
func (m *DefaultNamespaceSetting_SdkV2) SetNamespace(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["namespace"]
	m.Namespace = types.ListValueMust(t, vs)
}

type DefaultWarehouseId_SdkV2 struct {
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

	StringVal types.List `tfsdk:"string_val"`
}

func (to *DefaultWarehouseId_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DefaultWarehouseId_SdkV2) {
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

func (to *DefaultWarehouseId_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DefaultWarehouseId_SdkV2) {
	if !from.StringVal.IsNull() && !from.StringVal.IsUnknown() {
		if toStringVal, ok := to.GetStringVal(ctx); ok {
			if fromStringVal, ok := from.GetStringVal(ctx); ok {
				toStringVal.SyncFieldsDuringRead(ctx, fromStringVal)
				to.SetStringVal(ctx, toStringVal)
			}
		}
	}
}

func (m DefaultWarehouseId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["setting_name"] = attrs["setting_name"].SetOptional()
	attrs["string_val"] = attrs["string_val"].SetRequired()
	attrs["string_val"] = attrs["string_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DefaultWarehouseId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DefaultWarehouseId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"string_val": reflect.TypeOf(StringMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultWarehouseId_SdkV2
// only implements ToObjectValue() and Type().
func (m DefaultWarehouseId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         m.Etag,
			"setting_name": m.SettingName,
			"string_val":   m.StringVal,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DefaultWarehouseId_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag":         types.StringType,
			"setting_name": types.StringType,
			"string_val": basetypes.ListType{
				ElemType: StringMessage_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetStringVal returns the value of the StringVal field in DefaultWarehouseId_SdkV2 as
// a StringMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DefaultWarehouseId_SdkV2) GetStringVal(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if m.StringVal.IsNull() || m.StringVal.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := m.StringVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetStringVal sets the value of the StringVal field in DefaultWarehouseId_SdkV2.
func (m *DefaultWarehouseId_SdkV2) SetStringVal(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["string_val"]
	m.StringVal = types.ListValueMust(t, vs)
}

type DeleteAccountIpAccessEnableRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteAccountIpAccessEnableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountIpAccessEnableRequest_SdkV2) {
}

func (to *DeleteAccountIpAccessEnableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountIpAccessEnableRequest_SdkV2) {
}

func (m DeleteAccountIpAccessEnableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAccountIpAccessEnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAccountIpAccessEnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountIpAccessEnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteAccountIpAccessEnableResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteAccountIpAccessEnableResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountIpAccessEnableResponse_SdkV2) {
}

func (to *DeleteAccountIpAccessEnableResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountIpAccessEnableResponse_SdkV2) {
}

func (m DeleteAccountIpAccessEnableResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAccountIpAccessEnableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAccountIpAccessEnableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountIpAccessEnableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteAccountIpAccessListRequest_SdkV2 struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *DeleteAccountIpAccessListRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAccountIpAccessListRequest_SdkV2) {
}

func (to *DeleteAccountIpAccessListRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAccountIpAccessListRequest_SdkV2) {
}

func (m DeleteAccountIpAccessListRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAccountIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAccountIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAccountIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
}

func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) {
}

func (to *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) {
}

func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
}

func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) {
}

func (to *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) {
}

func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDashboardEmailSubscriptionsRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDashboardEmailSubscriptionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardEmailSubscriptionsRequest_SdkV2) {
}

func (to *DeleteDashboardEmailSubscriptionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardEmailSubscriptionsRequest_SdkV2) {
}

func (m DeleteDashboardEmailSubscriptionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDashboardEmailSubscriptionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardEmailSubscriptionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDashboardEmailSubscriptionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDashboardEmailSubscriptionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDashboardEmailSubscriptionsResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDashboardEmailSubscriptionsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDashboardEmailSubscriptionsResponse_SdkV2) {
}

func (to *DeleteDashboardEmailSubscriptionsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDashboardEmailSubscriptionsResponse_SdkV2) {
}

func (m DeleteDashboardEmailSubscriptionsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDashboardEmailSubscriptionsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardEmailSubscriptionsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDashboardEmailSubscriptionsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDashboardEmailSubscriptionsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDefaultNamespaceSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDefaultNamespaceSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultNamespaceSettingRequest_SdkV2) {
}

func (to *DeleteDefaultNamespaceSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultNamespaceSettingRequest_SdkV2) {
}

func (m DeleteDefaultNamespaceSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDefaultNamespaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDefaultNamespaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultNamespaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDefaultNamespaceSettingResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDefaultNamespaceSettingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultNamespaceSettingResponse_SdkV2) {
}

func (to *DeleteDefaultNamespaceSettingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultNamespaceSettingResponse_SdkV2) {
}

func (m DeleteDefaultNamespaceSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDefaultNamespaceSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDefaultNamespaceSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultNamespaceSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDefaultWarehouseIdRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDefaultWarehouseIdRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultWarehouseIdRequest_SdkV2) {
}

func (to *DeleteDefaultWarehouseIdRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultWarehouseIdRequest_SdkV2) {
}

func (m DeleteDefaultWarehouseIdRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDefaultWarehouseIdRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultWarehouseIdRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDefaultWarehouseIdRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultWarehouseIdRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDefaultWarehouseIdResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDefaultWarehouseIdResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDefaultWarehouseIdResponse_SdkV2) {
}

func (to *DeleteDefaultWarehouseIdResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDefaultWarehouseIdResponse_SdkV2) {
}

func (m DeleteDefaultWarehouseIdResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDefaultWarehouseIdResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultWarehouseIdResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDefaultWarehouseIdResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDefaultWarehouseIdResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDisableLegacyAccessRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDisableLegacyAccessRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyAccessRequest_SdkV2) {
}

func (to *DeleteDisableLegacyAccessRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyAccessRequest_SdkV2) {
}

func (m DeleteDisableLegacyAccessRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDisableLegacyAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDisableLegacyAccessResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDisableLegacyAccessResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyAccessResponse_SdkV2) {
}

func (to *DeleteDisableLegacyAccessResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyAccessResponse_SdkV2) {
}

func (m DeleteDisableLegacyAccessResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDisableLegacyAccessResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyAccessResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyAccessResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDisableLegacyDbfsRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDisableLegacyDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyDbfsRequest_SdkV2) {
}

func (to *DeleteDisableLegacyDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyDbfsRequest_SdkV2) {
}

func (m DeleteDisableLegacyDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDisableLegacyDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDisableLegacyDbfsResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDisableLegacyDbfsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyDbfsResponse_SdkV2) {
}

func (to *DeleteDisableLegacyDbfsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyDbfsResponse_SdkV2) {
}

func (m DeleteDisableLegacyDbfsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDisableLegacyDbfsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyDbfsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyDbfsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteDisableLegacyFeaturesRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteDisableLegacyFeaturesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyFeaturesRequest_SdkV2) {
}

func (to *DeleteDisableLegacyFeaturesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyFeaturesRequest_SdkV2) {
}

func (m DeleteDisableLegacyFeaturesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDisableLegacyFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteDisableLegacyFeaturesResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteDisableLegacyFeaturesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteDisableLegacyFeaturesResponse_SdkV2) {
}

func (to *DeleteDisableLegacyFeaturesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteDisableLegacyFeaturesResponse_SdkV2) {
}

func (m DeleteDisableLegacyFeaturesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteDisableLegacyFeaturesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteDisableLegacyFeaturesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteDisableLegacyFeaturesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteIpAccessListRequest_SdkV2 struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *DeleteIpAccessListRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteIpAccessListRequest_SdkV2) {
}

func (to *DeleteIpAccessListRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteIpAccessListRequest_SdkV2) {
}

func (m DeleteIpAccessListRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) {
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) {
}

func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) {
}

func (to *DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) {
}

func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteLlmProxyPartnerPoweredWorkspaceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteNetworkConnectivityConfigurationRequest_SdkV2 struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

func (to *DeleteNetworkConnectivityConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkConnectivityConfigurationRequest_SdkV2) {
}

func (to *DeleteNetworkConnectivityConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkConnectivityConfigurationRequest_SdkV2) {
}

func (m DeleteNetworkConnectivityConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteNetworkConnectivityConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkConnectivityConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteNetworkConnectivityConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNetworkConnectivityConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type DeleteNetworkPolicyRequest_SdkV2 struct {
	// The unique identifier of the network policy to delete.
	NetworkPolicyId types.String `tfsdk:"-"`
}

func (to *DeleteNetworkPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNetworkPolicyRequest_SdkV2) {
}

func (to *DeleteNetworkPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteNetworkPolicyRequest_SdkV2) {
}

func (m DeleteNetworkPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteNetworkPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteNetworkPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNetworkPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
		},
	}
}

type DeleteNotificationDestinationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *DeleteNotificationDestinationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteNotificationDestinationRequest_SdkV2) {
}

func (to *DeleteNotificationDestinationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteNotificationDestinationRequest_SdkV2) {
}

func (m DeleteNotificationDestinationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type DeletePersonalComputeSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeletePersonalComputeSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePersonalComputeSettingRequest_SdkV2) {
}

func (to *DeletePersonalComputeSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePersonalComputeSettingRequest_SdkV2) {
}

func (m DeletePersonalComputeSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeletePersonalComputeSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePersonalComputeSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePersonalComputeSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeletePersonalComputeSettingResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeletePersonalComputeSettingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePersonalComputeSettingResponse_SdkV2) {
}

func (to *DeletePersonalComputeSettingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePersonalComputeSettingResponse_SdkV2) {
}

func (m DeletePersonalComputeSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeletePersonalComputeSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePersonalComputeSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePersonalComputeSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeletePrivateEndpointRuleRequest_SdkV2 struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

func (to *DeletePrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeletePrivateEndpointRuleRequest_SdkV2) {
}

func (to *DeletePrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeletePrivateEndpointRuleRequest_SdkV2) {
}

func (m DeletePrivateEndpointRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeletePrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeletePrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       m.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeletePrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

type DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) {
}

func (to *DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) {
}

func (m DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) {
}

func (to *DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) {
}

func (m DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteSqlResultsDownloadRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *DeleteSqlResultsDownloadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSqlResultsDownloadRequest_SdkV2) {
}

func (to *DeleteSqlResultsDownloadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSqlResultsDownloadRequest_SdkV2) {
}

func (m DeleteSqlResultsDownloadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteSqlResultsDownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSqlResultsDownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSqlResultsDownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSqlResultsDownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// The etag is returned.
type DeleteSqlResultsDownloadResponse_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"etag"`
}

func (to *DeleteSqlResultsDownloadResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteSqlResultsDownloadResponse_SdkV2) {
}

func (to *DeleteSqlResultsDownloadResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteSqlResultsDownloadResponse_SdkV2) {
}

func (m DeleteSqlResultsDownloadResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteSqlResultsDownloadResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSqlResultsDownloadResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteSqlResultsDownloadResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteSqlResultsDownloadResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type DeleteTokenManagementRequest_SdkV2 struct {
	// The ID of the token to revoke.
	TokenId types.String `tfsdk:"-"`
}

func (to *DeleteTokenManagementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteTokenManagementRequest_SdkV2) {
}

func (to *DeleteTokenManagementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteTokenManagementRequest_SdkV2) {
}

func (m DeleteTokenManagementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DeleteTokenManagementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTokenManagementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteTokenManagementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteTokenManagementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type DisableLegacyAccess_SdkV2 struct {
	DisableLegacyAccess types.List `tfsdk:"disable_legacy_access"`
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

func (to *DisableLegacyAccess_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DisableLegacyAccess_SdkV2) {
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

func (to *DisableLegacyAccess_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DisableLegacyAccess_SdkV2) {
	if !from.DisableLegacyAccess.IsNull() && !from.DisableLegacyAccess.IsUnknown() {
		if toDisableLegacyAccess, ok := to.GetDisableLegacyAccess(ctx); ok {
			if fromDisableLegacyAccess, ok := from.GetDisableLegacyAccess(ctx); ok {
				toDisableLegacyAccess.SyncFieldsDuringRead(ctx, fromDisableLegacyAccess)
				to.SetDisableLegacyAccess(ctx, toDisableLegacyAccess)
			}
		}
	}
}

func (m DisableLegacyAccess_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disable_legacy_access"] = attrs["disable_legacy_access"].SetRequired()
	attrs["disable_legacy_access"] = attrs["disable_legacy_access"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m DisableLegacyAccess_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_access": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyAccess_SdkV2
// only implements ToObjectValue() and Type().
func (m DisableLegacyAccess_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_access": m.DisableLegacyAccess,
			"etag":                  m.Etag,
			"setting_name":          m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DisableLegacyAccess_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_access": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetDisableLegacyAccess returns the value of the DisableLegacyAccess field in DisableLegacyAccess_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DisableLegacyAccess_SdkV2) GetDisableLegacyAccess(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.DisableLegacyAccess.IsNull() || m.DisableLegacyAccess.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.DisableLegacyAccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDisableLegacyAccess sets the value of the DisableLegacyAccess field in DisableLegacyAccess_SdkV2.
func (m *DisableLegacyAccess_SdkV2) SetDisableLegacyAccess(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["disable_legacy_access"]
	m.DisableLegacyAccess = types.ListValueMust(t, vs)
}

type DisableLegacyDbfs_SdkV2 struct {
	DisableLegacyDbfs types.List `tfsdk:"disable_legacy_dbfs"`
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

func (to *DisableLegacyDbfs_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DisableLegacyDbfs_SdkV2) {
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

func (to *DisableLegacyDbfs_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DisableLegacyDbfs_SdkV2) {
	if !from.DisableLegacyDbfs.IsNull() && !from.DisableLegacyDbfs.IsUnknown() {
		if toDisableLegacyDbfs, ok := to.GetDisableLegacyDbfs(ctx); ok {
			if fromDisableLegacyDbfs, ok := from.GetDisableLegacyDbfs(ctx); ok {
				toDisableLegacyDbfs.SyncFieldsDuringRead(ctx, fromDisableLegacyDbfs)
				to.SetDisableLegacyDbfs(ctx, toDisableLegacyDbfs)
			}
		}
	}
}

func (m DisableLegacyDbfs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disable_legacy_dbfs"] = attrs["disable_legacy_dbfs"].SetRequired()
	attrs["disable_legacy_dbfs"] = attrs["disable_legacy_dbfs"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m DisableLegacyDbfs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_dbfs": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyDbfs_SdkV2
// only implements ToObjectValue() and Type().
func (m DisableLegacyDbfs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_dbfs": m.DisableLegacyDbfs,
			"etag":                m.Etag,
			"setting_name":        m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DisableLegacyDbfs_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_dbfs": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetDisableLegacyDbfs returns the value of the DisableLegacyDbfs field in DisableLegacyDbfs_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DisableLegacyDbfs_SdkV2) GetDisableLegacyDbfs(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.DisableLegacyDbfs.IsNull() || m.DisableLegacyDbfs.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.DisableLegacyDbfs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDisableLegacyDbfs sets the value of the DisableLegacyDbfs field in DisableLegacyDbfs_SdkV2.
func (m *DisableLegacyDbfs_SdkV2) SetDisableLegacyDbfs(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["disable_legacy_dbfs"]
	m.DisableLegacyDbfs = types.ListValueMust(t, vs)
}

type DisableLegacyFeatures_SdkV2 struct {
	DisableLegacyFeatures types.List `tfsdk:"disable_legacy_features"`
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

func (to *DisableLegacyFeatures_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DisableLegacyFeatures_SdkV2) {
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

func (to *DisableLegacyFeatures_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DisableLegacyFeatures_SdkV2) {
	if !from.DisableLegacyFeatures.IsNull() && !from.DisableLegacyFeatures.IsUnknown() {
		if toDisableLegacyFeatures, ok := to.GetDisableLegacyFeatures(ctx); ok {
			if fromDisableLegacyFeatures, ok := from.GetDisableLegacyFeatures(ctx); ok {
				toDisableLegacyFeatures.SyncFieldsDuringRead(ctx, fromDisableLegacyFeatures)
				to.SetDisableLegacyFeatures(ctx, toDisableLegacyFeatures)
			}
		}
	}
}

func (m DisableLegacyFeatures_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["disable_legacy_features"] = attrs["disable_legacy_features"].SetRequired()
	attrs["disable_legacy_features"] = attrs["disable_legacy_features"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m DisableLegacyFeatures_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_features": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyFeatures_SdkV2
// only implements ToObjectValue() and Type().
func (m DisableLegacyFeatures_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_features": m.DisableLegacyFeatures,
			"etag":                    m.Etag,
			"setting_name":            m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DisableLegacyFeatures_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"disable_legacy_features": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetDisableLegacyFeatures returns the value of the DisableLegacyFeatures field in DisableLegacyFeatures_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *DisableLegacyFeatures_SdkV2) GetDisableLegacyFeatures(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.DisableLegacyFeatures.IsNull() || m.DisableLegacyFeatures.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.DisableLegacyFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDisableLegacyFeatures sets the value of the DisableLegacyFeatures field in DisableLegacyFeatures_SdkV2.
func (m *DisableLegacyFeatures_SdkV2) SetDisableLegacyFeatures(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["disable_legacy_features"]
	m.DisableLegacyFeatures = types.ListValueMust(t, vs)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy_SdkV2 struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess types.List `tfsdk:"internet_access"`
}

func (to *EgressNetworkPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicy_SdkV2) {
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

func (to *EgressNetworkPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicy_SdkV2) {
	if !from.InternetAccess.IsNull() && !from.InternetAccess.IsUnknown() {
		if toInternetAccess, ok := to.GetInternetAccess(ctx); ok {
			if fromInternetAccess, ok := from.GetInternetAccess(ctx); ok {
				toInternetAccess.SyncFieldsDuringRead(ctx, fromInternetAccess)
				to.SetInternetAccess(ctx, toInternetAccess)
			}
		}
	}
}

func (m EgressNetworkPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["internet_access"] = attrs["internet_access"].SetOptional()
	attrs["internet_access"] = attrs["internet_access"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EgressNetworkPolicy.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EgressNetworkPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internet_access": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internet_access": m.InternetAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internet_access": basetypes.ListType{
				ElemType: EgressNetworkPolicyInternetAccessPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInternetAccess returns the value of the InternetAccess field in EgressNetworkPolicy_SdkV2 as
// a EgressNetworkPolicyInternetAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicy_SdkV2) GetInternetAccess(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicy_SdkV2, bool) {
	var e EgressNetworkPolicyInternetAccessPolicy_SdkV2
	if m.InternetAccess.IsNull() || m.InternetAccess.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicy_SdkV2
	d := m.InternetAccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInternetAccess sets the value of the InternetAccess field in EgressNetworkPolicy_SdkV2.
func (m *EgressNetworkPolicy_SdkV2) SetInternetAccess(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["internet_access"]
	m.InternetAccess = types.ListValueMust(t, vs)
}

type EgressNetworkPolicyInternetAccessPolicy_SdkV2 struct {
	AllowedInternetDestinations types.List `tfsdk:"allowed_internet_destinations"`

	AllowedStorageDestinations types.List `tfsdk:"allowed_storage_destinations"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode types.List `tfsdk:"log_only_mode"`

	RestrictionMode types.String `tfsdk:"restriction_mode"`
}

func (to *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicy_SdkV2) {
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

func (to *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicy_SdkV2) {
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

func (m EgressNetworkPolicyInternetAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_internet_destinations"] = attrs["allowed_internet_destinations"].SetOptional()
	attrs["allowed_storage_destinations"] = attrs["allowed_storage_destinations"].SetOptional()
	attrs["log_only_mode"] = attrs["log_only_mode"].SetOptional()
	attrs["log_only_mode"] = attrs["log_only_mode"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2{}),
		"log_only_mode":                 reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EgressNetworkPolicyInternetAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_internet_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2{}.Type(ctx),
			},
			"allowed_storage_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2{}.Type(ctx),
			},
			"log_only_mode": basetypes.ListType{
				ElemType: EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2{}.Type(ctx),
			},
			"restriction_mode": types.StringType,
		},
	}
}

// GetAllowedInternetDestinations returns the value of the AllowedInternetDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2 as
// a slice of EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2, bool) {
	if m.AllowedInternetDestinations.IsNull() || m.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2
	d := m.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2.
func (m *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2 as
// a slice of EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2, bool) {
	if m.AllowedStorageDestinations.IsNull() || m.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2
	d := m.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2.
func (m *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetLogOnlyMode returns the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy_SdkV2 as
// a EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetLogOnlyMode(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2, bool) {
	var e EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2
	if m.LogOnlyMode.IsNull() || m.LogOnlyMode.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2
	d := m.LogOnlyMode.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogOnlyMode sets the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy_SdkV2.
func (m *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SetLogOnlyMode(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["log_only_mode"]
	m.LogOnlyMode = types.ListValueMust(t, vs)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2 struct {
	Destination types.String `tfsdk:"destination"`

	Protocol types.String `tfsdk:"protocol"`

	Type_ types.String `tfsdk:"type"`
}

func (to *EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) {
}

func (to *EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) {
}

func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": m.Destination,
			"protocol":    m.Protocol,
			"type":        m.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination": types.StringType,
			"protocol":    types.StringType,
			"type":        types.StringType,
		},
	}
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2 struct {
	LogOnlyModeType types.String `tfsdk:"log_only_mode_type"`

	Workloads types.List `tfsdk:"workloads"`
}

func (to *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) {
	if !from.Workloads.IsNull() && !from.Workloads.IsUnknown() && to.Workloads.IsNull() && len(from.Workloads.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Workloads, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Workloads = from.Workloads
	}
}

func (to *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) {
	if !from.Workloads.IsNull() && !from.Workloads.IsUnknown() && to.Workloads.IsNull() && len(from.Workloads.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Workloads, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Workloads = from.Workloads
	}
}

func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workloads": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_only_mode_type": m.LogOnlyModeType,
			"workloads":          m.Workloads,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"log_only_mode_type": types.StringType,
			"workloads": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetWorkloads returns the value of the Workloads field in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) GetWorkloads(ctx context.Context) ([]types.String, bool) {
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

// SetWorkloads sets the value of the Workloads field in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2.
func (m *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) SetWorkloads(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workloads"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Workloads = types.ListValueMust(t, vs)
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2 struct {
	AllowedPaths types.List `tfsdk:"allowed_paths"`

	AzureContainer types.String `tfsdk:"azure_container"`

	AzureDnsZone types.String `tfsdk:"azure_dns_zone"`

	AzureStorageAccount types.String `tfsdk:"azure_storage_account"`

	AzureStorageService types.String `tfsdk:"azure_storage_service"`

	BucketName types.String `tfsdk:"bucket_name"`

	Region types.String `tfsdk:"region"`

	Type_ types.String `tfsdk:"type"`
}

func (to *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) {
	if !from.AllowedPaths.IsNull() && !from.AllowedPaths.IsUnknown() && to.AllowedPaths.IsNull() && len(from.AllowedPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedPaths = from.AllowedPaths
	}
}

func (to *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) {
	if !from.AllowedPaths.IsNull() && !from.AllowedPaths.IsUnknown() && to.AllowedPaths.IsNull() && len(from.AllowedPaths.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllowedPaths, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllowedPaths = from.AllowedPaths
	}
}

func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_paths": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetAllowedPaths returns the value of the AllowedPaths field in EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) GetAllowedPaths(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedPaths sets the value of the AllowedPaths field in EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2.
func (m *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) SetAllowedPaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedPaths = types.ListValueMust(t, vs)
}

type EgressNetworkPolicyNetworkAccessPolicy_SdkV2 struct {
	// List of internet destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	AllowedInternetDestinations types.List `tfsdk:"allowed_internet_destinations"`
	// List of storage destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	AllowedStorageDestinations types.List `tfsdk:"allowed_storage_destinations"`
	// Optional. When policy_enforcement is not provided, we default to
	// ENFORCE_MODE_ALL_SERVICES
	PolicyEnforcement types.List `tfsdk:"policy_enforcement"`
	// The restriction mode that controls how serverless workloads can access
	// the internet.
	RestrictionMode types.String `tfsdk:"restriction_mode"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicy_SdkV2) {
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

func (to *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicy_SdkV2) {
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

func (m EgressNetworkPolicyNetworkAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allowed_internet_destinations"] = attrs["allowed_internet_destinations"].SetOptional()
	attrs["allowed_storage_destinations"] = attrs["allowed_storage_destinations"].SetOptional()
	attrs["policy_enforcement"] = attrs["policy_enforcement"].SetOptional()
	attrs["policy_enforcement"] = attrs["policy_enforcement"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EgressNetworkPolicyNetworkAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2{}),
		"policy_enforcement":            reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EgressNetworkPolicyNetworkAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allowed_internet_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2{}.Type(ctx),
			},
			"allowed_storage_destinations": basetypes.ListType{
				ElemType: EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2{}.Type(ctx),
			},
			"policy_enforcement": basetypes.ListType{
				ElemType: EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2{}.Type(ctx),
			},
			"restriction_mode": types.StringType,
		},
	}
}

// GetAllowedInternetDestinations returns the value of the AllowedInternetDestinations field in EgressNetworkPolicyNetworkAccessPolicy_SdkV2 as
// a slice of EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2, bool) {
	if m.AllowedInternetDestinations.IsNull() || m.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2
	d := m.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyNetworkAccessPolicy_SdkV2.
func (m *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyNetworkAccessPolicy_SdkV2 as
// a slice of EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2, bool) {
	if m.AllowedStorageDestinations.IsNull() || m.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2
	d := m.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyNetworkAccessPolicy_SdkV2.
func (m *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetPolicyEnforcement returns the value of the PolicyEnforcement field in EgressNetworkPolicyNetworkAccessPolicy_SdkV2 as
// a EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) GetPolicyEnforcement(ctx context.Context) (EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2, bool) {
	var e EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2
	if m.PolicyEnforcement.IsNull() || m.PolicyEnforcement.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2
	d := m.PolicyEnforcement.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPolicyEnforcement sets the value of the PolicyEnforcement field in EgressNetworkPolicyNetworkAccessPolicy_SdkV2.
func (m *EgressNetworkPolicyNetworkAccessPolicy_SdkV2) SetPolicyEnforcement(ctx context.Context, v EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["policy_enforcement"]
	m.PolicyEnforcement = types.ListValueMust(t, vs)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support DNS_NAME (FQDN format) destinations for the time
// being. Going forward we may extend support to host names and IP addresses.
type EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2 struct {
	// The internet destination to which access will be allowed. Format
	// dependent on the destination type.
	Destination types.String `tfsdk:"destination"`
	// The type of internet destination. Currently only DNS_NAME is supported.
	InternetDestinationType types.String `tfsdk:"internet_destination_type"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) {
}

func (to *EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) {
}

func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination":               m.Destination,
			"internet_destination_type": m.InternetDestinationType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyNetworkAccessPolicyInternetDestination_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination":               types.StringType,
			"internet_destination_type": types.StringType,
		},
	}
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2 struct {
	// When empty, it means dry run for all products. When non-empty, it means
	// dry run for specific products and for the other products, they will run
	// in enforced mode.
	DryRunModeProductFilter types.List `tfsdk:"dry_run_mode_product_filter"`
	// The mode of policy enforcement. ENFORCED blocks traffic that violates
	// policy, while DRY_RUN only logs violations without blocking. When not
	// specified, defaults to ENFORCED.
	EnforcementMode types.String `tfsdk:"enforcement_mode"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) {
	if !from.DryRunModeProductFilter.IsNull() && !from.DryRunModeProductFilter.IsUnknown() && to.DryRunModeProductFilter.IsNull() && len(from.DryRunModeProductFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DryRunModeProductFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DryRunModeProductFilter = from.DryRunModeProductFilter
	}
}

func (to *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) {
	if !from.DryRunModeProductFilter.IsNull() && !from.DryRunModeProductFilter.IsUnknown() && to.DryRunModeProductFilter.IsNull() && len(from.DryRunModeProductFilter.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DryRunModeProductFilter, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DryRunModeProductFilter = from.DryRunModeProductFilter
	}
}

func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dry_run_mode_product_filter": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dry_run_mode_product_filter": m.DryRunModeProductFilter,
			"enforcement_mode":            m.EnforcementMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dry_run_mode_product_filter": basetypes.ListType{
				ElemType: types.StringType,
			},
			"enforcement_mode": types.StringType,
		},
	}
}

// GetDryRunModeProductFilter returns the value of the DryRunModeProductFilter field in EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) GetDryRunModeProductFilter(ctx context.Context) ([]types.String, bool) {
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

// SetDryRunModeProductFilter sets the value of the DryRunModeProductFilter field in EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2.
func (m *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement_SdkV2) SetDryRunModeProductFilter(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dry_run_mode_product_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DryRunModeProductFilter = types.ListValueMust(t, vs)
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2 struct {
	// The Azure storage account name.
	AzureStorageAccount types.String `tfsdk:"azure_storage_account"`
	// The Azure storage service type (blob, dfs, etc.).
	AzureStorageService types.String `tfsdk:"azure_storage_service"`

	BucketName types.String `tfsdk:"bucket_name"`

	Region types.String `tfsdk:"region"`
	// The type of storage destination.
	StorageDestinationType types.String `tfsdk:"storage_destination_type"`
}

func (to *EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) {
}

func (to *EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) {
}

func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2
// only implements ToObjectValue() and Type().
func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m EgressNetworkPolicyNetworkAccessPolicyStorageDestination_SdkV2) Type(ctx context.Context) attr.Type {
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

type EmailConfig_SdkV2 struct {
	// Email addresses to notify.
	Addresses types.List `tfsdk:"addresses"`
}

func (to *EmailConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmailConfig_SdkV2) {
	if !from.Addresses.IsNull() && !from.Addresses.IsUnknown() && to.Addresses.IsNull() && len(from.Addresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Addresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Addresses = from.Addresses
	}
}

func (to *EmailConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmailConfig_SdkV2) {
	if !from.Addresses.IsNull() && !from.Addresses.IsUnknown() && to.Addresses.IsNull() && len(from.Addresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Addresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Addresses = from.Addresses
	}
}

func (m EmailConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EmailConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmailConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m EmailConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"addresses": m.Addresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmailConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"addresses": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetAddresses returns the value of the Addresses field in EmailConfig_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EmailConfig_SdkV2) GetAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetAddresses sets the value of the Addresses field in EmailConfig_SdkV2.
func (m *EmailConfig_SdkV2) SetAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Addresses = types.ListValueMust(t, vs)
}

type Empty_SdkV2 struct {
}

func (to *Empty_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Empty_SdkV2) {
}

func (to *Empty_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Empty_SdkV2) {
}

func (m Empty_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Empty_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty_SdkV2
// only implements ToObjectValue() and Type().
func (m Empty_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m Empty_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type EnableExportNotebook_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnableExportNotebook_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnableExportNotebook_SdkV2) {
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

func (to *EnableExportNotebook_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EnableExportNotebook_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m EnableExportNotebook_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EnableExportNotebook_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableExportNotebook_SdkV2
// only implements ToObjectValue() and Type().
func (m EnableExportNotebook_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnableExportNotebook_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in EnableExportNotebook_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnableExportNotebook_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in EnableExportNotebook_SdkV2.
func (m *EnableExportNotebook_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

type EnableNotebookTableClipboard_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnableNotebookTableClipboard_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnableNotebookTableClipboard_SdkV2) {
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

func (to *EnableNotebookTableClipboard_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EnableNotebookTableClipboard_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m EnableNotebookTableClipboard_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EnableNotebookTableClipboard_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableNotebookTableClipboard_SdkV2
// only implements ToObjectValue() and Type().
func (m EnableNotebookTableClipboard_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnableNotebookTableClipboard_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in EnableNotebookTableClipboard_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnableNotebookTableClipboard_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in EnableNotebookTableClipboard_SdkV2.
func (m *EnableNotebookTableClipboard_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

type EnableResultsDownloading_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *EnableResultsDownloading_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnableResultsDownloading_SdkV2) {
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

func (to *EnableResultsDownloading_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EnableResultsDownloading_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m EnableResultsDownloading_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetOptional()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EnableResultsDownloading_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableResultsDownloading_SdkV2
// only implements ToObjectValue() and Type().
func (m EnableResultsDownloading_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnableResultsDownloading_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in EnableResultsDownloading_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnableResultsDownloading_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in EnableResultsDownloading_SdkV2.
func (m *EnableResultsDownloading_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring_SdkV2 struct {
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (to *EnhancedSecurityMonitoring_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnhancedSecurityMonitoring_SdkV2) {
}

func (to *EnhancedSecurityMonitoring_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EnhancedSecurityMonitoring_SdkV2) {
}

func (m EnhancedSecurityMonitoring_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EnhancedSecurityMonitoring_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoring_SdkV2
// only implements ToObjectValue() and Type().
func (m EnhancedSecurityMonitoring_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enabled": m.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnhancedSecurityMonitoring_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enabled": types.BoolType,
		},
	}
}

type EnhancedSecurityMonitoringSetting_SdkV2 struct {
	EnhancedSecurityMonitoringWorkspace types.List `tfsdk:"enhanced_security_monitoring_workspace"`
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

func (to *EnhancedSecurityMonitoringSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EnhancedSecurityMonitoringSetting_SdkV2) {
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

func (to *EnhancedSecurityMonitoringSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EnhancedSecurityMonitoringSetting_SdkV2) {
	if !from.EnhancedSecurityMonitoringWorkspace.IsNull() && !from.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		if toEnhancedSecurityMonitoringWorkspace, ok := to.GetEnhancedSecurityMonitoringWorkspace(ctx); ok {
			if fromEnhancedSecurityMonitoringWorkspace, ok := from.GetEnhancedSecurityMonitoringWorkspace(ctx); ok {
				toEnhancedSecurityMonitoringWorkspace.SyncFieldsDuringRead(ctx, fromEnhancedSecurityMonitoringWorkspace)
				to.SetEnhancedSecurityMonitoringWorkspace(ctx, toEnhancedSecurityMonitoringWorkspace)
			}
		}
	}
}

func (m EnhancedSecurityMonitoringSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enhanced_security_monitoring_workspace"] = attrs["enhanced_security_monitoring_workspace"].SetRequired()
	attrs["enhanced_security_monitoring_workspace"] = attrs["enhanced_security_monitoring_workspace"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EnhancedSecurityMonitoringSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enhanced_security_monitoring_workspace": reflect.TypeOf(EnhancedSecurityMonitoring_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoringSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m EnhancedSecurityMonitoringSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enhanced_security_monitoring_workspace": m.EnhancedSecurityMonitoringWorkspace,
			"etag":                                   m.Etag,
			"setting_name":                           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EnhancedSecurityMonitoringSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enhanced_security_monitoring_workspace": basetypes.ListType{
				ElemType: EnhancedSecurityMonitoring_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetEnhancedSecurityMonitoringWorkspace returns the value of the EnhancedSecurityMonitoringWorkspace field in EnhancedSecurityMonitoringSetting_SdkV2 as
// a EnhancedSecurityMonitoring_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EnhancedSecurityMonitoringSetting_SdkV2) GetEnhancedSecurityMonitoringWorkspace(ctx context.Context) (EnhancedSecurityMonitoring_SdkV2, bool) {
	var e EnhancedSecurityMonitoring_SdkV2
	if m.EnhancedSecurityMonitoringWorkspace.IsNull() || m.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		return e, false
	}
	var v []EnhancedSecurityMonitoring_SdkV2
	d := m.EnhancedSecurityMonitoringWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnhancedSecurityMonitoringWorkspace sets the value of the EnhancedSecurityMonitoringWorkspace field in EnhancedSecurityMonitoringSetting_SdkV2.
func (m *EnhancedSecurityMonitoringSetting_SdkV2) SetEnhancedSecurityMonitoringWorkspace(ctx context.Context, v EnhancedSecurityMonitoring_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["enhanced_security_monitoring_workspace"]
	m.EnhancedSecurityMonitoringWorkspace = types.ListValueMust(t, vs)
}

// Account level policy for ESM
type EsmEnablementAccount_SdkV2 struct {
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (to *EsmEnablementAccount_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EsmEnablementAccount_SdkV2) {
}

func (to *EsmEnablementAccount_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EsmEnablementAccount_SdkV2) {
}

func (m EsmEnablementAccount_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m EsmEnablementAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccount_SdkV2
// only implements ToObjectValue() and Type().
func (m EsmEnablementAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enforced": m.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EsmEnablementAccount_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enforced": types.BoolType,
		},
	}
}

type EsmEnablementAccountSetting_SdkV2 struct {
	EsmEnablementAccount types.List `tfsdk:"esm_enablement_account"`
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

func (to *EsmEnablementAccountSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EsmEnablementAccountSetting_SdkV2) {
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

func (to *EsmEnablementAccountSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EsmEnablementAccountSetting_SdkV2) {
	if !from.EsmEnablementAccount.IsNull() && !from.EsmEnablementAccount.IsUnknown() {
		if toEsmEnablementAccount, ok := to.GetEsmEnablementAccount(ctx); ok {
			if fromEsmEnablementAccount, ok := from.GetEsmEnablementAccount(ctx); ok {
				toEsmEnablementAccount.SyncFieldsDuringRead(ctx, fromEsmEnablementAccount)
				to.SetEsmEnablementAccount(ctx, toEsmEnablementAccount)
			}
		}
	}
}

func (m EsmEnablementAccountSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["esm_enablement_account"] = attrs["esm_enablement_account"].SetRequired()
	attrs["esm_enablement_account"] = attrs["esm_enablement_account"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m EsmEnablementAccountSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"esm_enablement_account": reflect.TypeOf(EsmEnablementAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccountSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m EsmEnablementAccountSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"esm_enablement_account": m.EsmEnablementAccount,
			"etag":                   m.Etag,
			"setting_name":           m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EsmEnablementAccountSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"esm_enablement_account": basetypes.ListType{
				ElemType: EsmEnablementAccount_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetEsmEnablementAccount returns the value of the EsmEnablementAccount field in EsmEnablementAccountSetting_SdkV2 as
// a EsmEnablementAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EsmEnablementAccountSetting_SdkV2) GetEsmEnablementAccount(ctx context.Context) (EsmEnablementAccount_SdkV2, bool) {
	var e EsmEnablementAccount_SdkV2
	if m.EsmEnablementAccount.IsNull() || m.EsmEnablementAccount.IsUnknown() {
		return e, false
	}
	var v []EsmEnablementAccount_SdkV2
	d := m.EsmEnablementAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEsmEnablementAccount sets the value of the EsmEnablementAccount field in EsmEnablementAccountSetting_SdkV2.
func (m *EsmEnablementAccountSetting_SdkV2) SetEsmEnablementAccount(ctx context.Context, v EsmEnablementAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["esm_enablement_account"]
	m.EsmEnablementAccount = types.ListValueMust(t, vs)
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken_SdkV2 struct {
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

func (to *ExchangeToken_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeToken_SdkV2) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (to *ExchangeToken_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExchangeToken_SdkV2) {
	if !from.Scopes.IsNull() && !from.Scopes.IsUnknown() && to.Scopes.IsNull() && len(from.Scopes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Scopes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Scopes = from.Scopes
	}
}

func (m ExchangeToken_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExchangeToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeToken_SdkV2
// only implements ToObjectValue() and Type().
func (m ExchangeToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ExchangeToken_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetScopes returns the value of the Scopes field in ExchangeToken_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeToken_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeToken_SdkV2.
func (m *ExchangeToken_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// Exchange a token with the IdP
type ExchangeTokenRequest_SdkV2 struct {
	// The partition of Credentials store
	PartitionId types.List `tfsdk:"partition_id"`
	// Array of scopes for the token request.
	Scopes types.List `tfsdk:"scopes"`
	// A list of token types being requested
	TokenType types.List `tfsdk:"token_type"`
}

func (to *ExchangeTokenRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeTokenRequest_SdkV2) {
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

func (to *ExchangeTokenRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExchangeTokenRequest_SdkV2) {
	if !from.PartitionId.IsNull() && !from.PartitionId.IsUnknown() {
		if toPartitionId, ok := to.GetPartitionId(ctx); ok {
			if fromPartitionId, ok := from.GetPartitionId(ctx); ok {
				toPartitionId.SyncFieldsDuringRead(ctx, fromPartitionId)
				to.SetPartitionId(ctx, toPartitionId)
			}
		}
	}
}

func (m ExchangeTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["partition_id"] = attrs["partition_id"].SetRequired()
	attrs["partition_id"] = attrs["partition_id"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m ExchangeTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partition_id": reflect.TypeOf(PartitionId_SdkV2{}),
		"scopes":       reflect.TypeOf(types.String{}),
		"token_type":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExchangeTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"partition_id": m.PartitionId,
			"scopes":       m.Scopes,
			"token_type":   m.TokenType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"partition_id": basetypes.ListType{
				ElemType: PartitionId_SdkV2{}.Type(ctx),
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"token_type": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPartitionId returns the value of the PartitionId field in ExchangeTokenRequest_SdkV2 as
// a PartitionId_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenRequest_SdkV2) GetPartitionId(ctx context.Context) (PartitionId_SdkV2, bool) {
	var e PartitionId_SdkV2
	if m.PartitionId.IsNull() || m.PartitionId.IsUnknown() {
		return e, false
	}
	var v []PartitionId_SdkV2
	d := m.PartitionId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPartitionId sets the value of the PartitionId field in ExchangeTokenRequest_SdkV2.
func (m *ExchangeTokenRequest_SdkV2) SetPartitionId(ctx context.Context, v PartitionId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["partition_id"]
	m.PartitionId = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in ExchangeTokenRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenRequest_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeTokenRequest_SdkV2.
func (m *ExchangeTokenRequest_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Scopes = types.ListValueMust(t, vs)
}

// GetTokenType returns the value of the TokenType field in ExchangeTokenRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenRequest_SdkV2) GetTokenType(ctx context.Context) ([]types.String, bool) {
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

// SetTokenType sets the value of the TokenType field in ExchangeTokenRequest_SdkV2.
func (m *ExchangeTokenRequest_SdkV2) SetTokenType(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_type"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TokenType = types.ListValueMust(t, vs)
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse_SdkV2 struct {
	Values types.List `tfsdk:"values"`
}

func (to *ExchangeTokenResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExchangeTokenResponse_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (to *ExchangeTokenResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExchangeTokenResponse_SdkV2) {
	if !from.Values.IsNull() && !from.Values.IsUnknown() && to.Values.IsNull() && len(from.Values.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Values, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Values = from.Values
	}
}

func (m ExchangeTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ExchangeTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(ExchangeToken_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ExchangeTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": m.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExchangeTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"values": basetypes.ListType{
				ElemType: ExchangeToken_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetValues returns the value of the Values field in ExchangeTokenResponse_SdkV2 as
// a slice of ExchangeToken_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ExchangeTokenResponse_SdkV2) GetValues(ctx context.Context) ([]ExchangeToken_SdkV2, bool) {
	if m.Values.IsNull() || m.Values.IsUnknown() {
		return nil, false
	}
	var v []ExchangeToken_SdkV2
	d := m.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in ExchangeTokenResponse_SdkV2.
func (m *ExchangeTokenResponse_SdkV2) SetValues(ctx context.Context, v []ExchangeToken_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Values = types.ListValueMust(t, vs)
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse_SdkV2 struct {
	IpAccessList types.List `tfsdk:"ip_access_list"`
}

func (to *FetchIpAccessListResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FetchIpAccessListResponse_SdkV2) {
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

func (to *FetchIpAccessListResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FetchIpAccessListResponse_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (m FetchIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FetchIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FetchIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FetchIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m FetchIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": m.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FetchIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessListInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in FetchIpAccessListResponse_SdkV2 as
// a IpAccessListInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *FetchIpAccessListResponse_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessListInfo_SdkV2, bool) {
	var e IpAccessListInfo_SdkV2
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo_SdkV2
	d := m.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in FetchIpAccessListResponse_SdkV2.
func (m *FetchIpAccessListResponse_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessListInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	m.IpAccessList = types.ListValueMust(t, vs)
}

type GenericWebhookConfig_SdkV2 struct {
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

func (to *GenericWebhookConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenericWebhookConfig_SdkV2) {
}

func (to *GenericWebhookConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GenericWebhookConfig_SdkV2) {
}

func (m GenericWebhookConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GenericWebhookConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenericWebhookConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m GenericWebhookConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m GenericWebhookConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

type GetAccountIpAccessEnableRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAccountIpAccessEnableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAccountIpAccessEnableRequest_SdkV2) {
}

func (to *GetAccountIpAccessEnableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAccountIpAccessEnableRequest_SdkV2) {
}

func (m GetAccountIpAccessEnableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAccountIpAccessEnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessEnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAccountIpAccessEnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAccountIpAccessEnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetAccountIpAccessListRequest_SdkV2 struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *GetAccountIpAccessListRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAccountIpAccessListRequest_SdkV2) {
}

func (to *GetAccountIpAccessListRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAccountIpAccessListRequest_SdkV2) {
}

func (m GetAccountIpAccessListRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAccountIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAccountIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAccountIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
}

func (to *GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
}

func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
}

func (to *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
}

func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetAutomaticClusterUpdateSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetAutomaticClusterUpdateSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetAutomaticClusterUpdateSettingRequest_SdkV2) {
}

func (to *GetAutomaticClusterUpdateSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetAutomaticClusterUpdateSettingRequest_SdkV2) {
}

func (m GetAutomaticClusterUpdateSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetAutomaticClusterUpdateSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAutomaticClusterUpdateSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetAutomaticClusterUpdateSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetAutomaticClusterUpdateSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetComplianceSecurityProfileSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetComplianceSecurityProfileSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetComplianceSecurityProfileSettingRequest_SdkV2) {
}

func (to *GetComplianceSecurityProfileSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetComplianceSecurityProfileSettingRequest_SdkV2) {
}

func (m GetComplianceSecurityProfileSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetComplianceSecurityProfileSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetComplianceSecurityProfileSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetComplianceSecurityProfileSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetComplianceSecurityProfileSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetCspEnablementAccountSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetCspEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetCspEnablementAccountSettingRequest_SdkV2) {
}

func (to *GetCspEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetCspEnablementAccountSettingRequest_SdkV2) {
}

func (m GetCspEnablementAccountSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetCspEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCspEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetCspEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetCspEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDashboardEmailSubscriptionsRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDashboardEmailSubscriptionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDashboardEmailSubscriptionsRequest_SdkV2) {
}

func (to *GetDashboardEmailSubscriptionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDashboardEmailSubscriptionsRequest_SdkV2) {
}

func (m GetDashboardEmailSubscriptionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDashboardEmailSubscriptionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDashboardEmailSubscriptionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDashboardEmailSubscriptionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDashboardEmailSubscriptionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDefaultNamespaceSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDefaultNamespaceSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDefaultNamespaceSettingRequest_SdkV2) {
}

func (to *GetDefaultNamespaceSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDefaultNamespaceSettingRequest_SdkV2) {
}

func (m GetDefaultNamespaceSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDefaultNamespaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultNamespaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDefaultNamespaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDefaultNamespaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDefaultWarehouseIdRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDefaultWarehouseIdRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDefaultWarehouseIdRequest_SdkV2) {
}

func (to *GetDefaultWarehouseIdRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDefaultWarehouseIdRequest_SdkV2) {
}

func (m GetDefaultWarehouseIdRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDefaultWarehouseIdRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultWarehouseIdRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDefaultWarehouseIdRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDefaultWarehouseIdRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDisableLegacyAccessRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDisableLegacyAccessRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDisableLegacyAccessRequest_SdkV2) {
}

func (to *GetDisableLegacyAccessRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDisableLegacyAccessRequest_SdkV2) {
}

func (m GetDisableLegacyAccessRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDisableLegacyAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDisableLegacyAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDisableLegacyAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDisableLegacyDbfsRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDisableLegacyDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDisableLegacyDbfsRequest_SdkV2) {
}

func (to *GetDisableLegacyDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDisableLegacyDbfsRequest_SdkV2) {
}

func (m GetDisableLegacyDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDisableLegacyDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDisableLegacyDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDisableLegacyDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetDisableLegacyFeaturesRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetDisableLegacyFeaturesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetDisableLegacyFeaturesRequest_SdkV2) {
}

func (to *GetDisableLegacyFeaturesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetDisableLegacyFeaturesRequest_SdkV2) {
}

func (m GetDisableLegacyFeaturesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetDisableLegacyFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetDisableLegacyFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetDisableLegacyFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetEnableExportNotebookRequest_SdkV2 struct {
}

func (to *GetEnableExportNotebookRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnableExportNotebookRequest_SdkV2) {
}

func (to *GetEnableExportNotebookRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEnableExportNotebookRequest_SdkV2) {
}

func (m GetEnableExportNotebookRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnableExportNotebookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnableExportNotebookRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnableExportNotebookRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEnableExportNotebookRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnableExportNotebookRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetEnableNotebookTableClipboardRequest_SdkV2 struct {
}

func (to *GetEnableNotebookTableClipboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnableNotebookTableClipboardRequest_SdkV2) {
}

func (to *GetEnableNotebookTableClipboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEnableNotebookTableClipboardRequest_SdkV2) {
}

func (m GetEnableNotebookTableClipboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnableNotebookTableClipboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnableNotebookTableClipboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnableNotebookTableClipboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEnableNotebookTableClipboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnableNotebookTableClipboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetEnableResultsDownloadingRequest_SdkV2 struct {
}

func (to *GetEnableResultsDownloadingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnableResultsDownloadingRequest_SdkV2) {
}

func (to *GetEnableResultsDownloadingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEnableResultsDownloadingRequest_SdkV2) {
}

func (m GetEnableResultsDownloadingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnableResultsDownloadingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetEnableResultsDownloadingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnableResultsDownloadingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEnableResultsDownloadingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnableResultsDownloadingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetEnhancedSecurityMonitoringSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetEnhancedSecurityMonitoringSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEnhancedSecurityMonitoringSettingRequest_SdkV2) {
}

func (to *GetEnhancedSecurityMonitoringSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEnhancedSecurityMonitoringSettingRequest_SdkV2) {
}

func (m GetEnhancedSecurityMonitoringSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetEnhancedSecurityMonitoringSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnhancedSecurityMonitoringSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEnhancedSecurityMonitoringSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEnhancedSecurityMonitoringSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetEsmEnablementAccountSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetEsmEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetEsmEnablementAccountSettingRequest_SdkV2) {
}

func (to *GetEsmEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetEsmEnablementAccountSettingRequest_SdkV2) {
}

func (m GetEsmEnablementAccountSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetEsmEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEsmEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetEsmEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetEsmEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetIpAccessListRequest_SdkV2 struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

func (to *GetIpAccessListRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIpAccessListRequest_SdkV2) {
}

func (to *GetIpAccessListRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetIpAccessListRequest_SdkV2) {
}

func (m GetIpAccessListRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": m.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type GetIpAccessListResponse_SdkV2 struct {
	IpAccessList types.List `tfsdk:"ip_access_list"`
}

func (to *GetIpAccessListResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIpAccessListResponse_SdkV2) {
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

func (to *GetIpAccessListResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetIpAccessListResponse_SdkV2) {
	if !from.IpAccessList.IsNull() && !from.IpAccessList.IsUnknown() {
		if toIpAccessList, ok := to.GetIpAccessList(ctx); ok {
			if fromIpAccessList, ok := from.GetIpAccessList(ctx); ok {
				toIpAccessList.SyncFieldsDuringRead(ctx, fromIpAccessList)
				to.SetIpAccessList(ctx, toIpAccessList)
			}
		}
	}
}

func (m GetIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ip_access_list"] = attrs["ip_access_list"].SetOptional()
	attrs["ip_access_list"] = attrs["ip_access_list"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": m.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": basetypes.ListType{
				ElemType: IpAccessListInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in GetIpAccessListResponse_SdkV2 as
// a IpAccessListInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetIpAccessListResponse_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessListInfo_SdkV2, bool) {
	var e IpAccessListInfo_SdkV2
	if m.IpAccessList.IsNull() || m.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo_SdkV2
	d := m.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in GetIpAccessListResponse_SdkV2.
func (m *GetIpAccessListResponse_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessListInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	m.IpAccessList = types.ListValueMust(t, vs)
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse_SdkV2 struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (to *GetIpAccessListsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetIpAccessListsResponse_SdkV2) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (to *GetIpAccessListsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetIpAccessListsResponse_SdkV2) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (m GetIpAccessListsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetIpAccessListsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetIpAccessListsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": m.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetIpAccessListsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_lists": basetypes.ListType{
				ElemType: IpAccessListInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessLists returns the value of the IpAccessLists field in GetIpAccessListsResponse_SdkV2 as
// a slice of IpAccessListInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetIpAccessListsResponse_SdkV2) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo_SdkV2, bool) {
	if m.IpAccessLists.IsNull() || m.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo_SdkV2
	d := m.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in GetIpAccessListsResponse_SdkV2.
func (m *GetIpAccessListsResponse_SdkV2) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAccessLists = types.ListValueMust(t, vs)
}

type GetLlmProxyPartnerPoweredAccountRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetLlmProxyPartnerPoweredAccountRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLlmProxyPartnerPoweredAccountRequest_SdkV2) {
}

func (to *GetLlmProxyPartnerPoweredAccountRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLlmProxyPartnerPoweredAccountRequest_SdkV2) {
}

func (m GetLlmProxyPartnerPoweredAccountRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLlmProxyPartnerPoweredAccountRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredAccountRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLlmProxyPartnerPoweredAccountRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLlmProxyPartnerPoweredAccountRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetLlmProxyPartnerPoweredEnforceRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) {
}

func (to *GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) {
}

func (m GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredEnforceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLlmProxyPartnerPoweredEnforceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) {
}

func (to *GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) {
}

func (m GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetNetworkConnectivityConfigurationRequest_SdkV2 struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

func (to *GetNetworkConnectivityConfigurationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkConnectivityConfigurationRequest_SdkV2) {
}

func (to *GetNetworkConnectivityConfigurationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetNetworkConnectivityConfigurationRequest_SdkV2) {
}

func (m GetNetworkConnectivityConfigurationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetNetworkConnectivityConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkConnectivityConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetNetworkConnectivityConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNetworkConnectivityConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type GetNetworkPolicyRequest_SdkV2 struct {
	// The unique identifier of the network policy to retrieve.
	NetworkPolicyId types.String `tfsdk:"-"`
}

func (to *GetNetworkPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNetworkPolicyRequest_SdkV2) {
}

func (to *GetNetworkPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetNetworkPolicyRequest_SdkV2) {
}

func (m GetNetworkPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetNetworkPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetNetworkPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNetworkPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
		},
	}
}

type GetNotificationDestinationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

func (to *GetNotificationDestinationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetNotificationDestinationRequest_SdkV2) {
}

func (to *GetNotificationDestinationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetNotificationDestinationRequest_SdkV2) {
}

func (m GetNotificationDestinationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

type GetPersonalComputeSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetPersonalComputeSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPersonalComputeSettingRequest_SdkV2) {
}

func (to *GetPersonalComputeSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPersonalComputeSettingRequest_SdkV2) {
}

func (m GetPersonalComputeSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPersonalComputeSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalComputeSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPersonalComputeSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPersonalComputeSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetPrivateEndpointRuleRequest_SdkV2 struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

func (to *GetPrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetPrivateEndpointRuleRequest_SdkV2) {
}

func (to *GetPrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetPrivateEndpointRuleRequest_SdkV2) {
}

func (m GetPrivateEndpointRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetPrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetPrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       m.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetPrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

type GetRestrictWorkspaceAdminsSettingRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetRestrictWorkspaceAdminsSettingRequest_SdkV2) {
}

func (to *GetRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetRestrictWorkspaceAdminsSettingRequest_SdkV2) {
}

func (m GetRestrictWorkspaceAdminsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetRestrictWorkspaceAdminsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRestrictWorkspaceAdminsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetRestrictWorkspaceAdminsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetRestrictWorkspaceAdminsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetSqlResultsDownloadRequest_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag types.String `tfsdk:"-"`
}

func (to *GetSqlResultsDownloadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetSqlResultsDownloadRequest_SdkV2) {
}

func (to *GetSqlResultsDownloadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetSqlResultsDownloadRequest_SdkV2) {
}

func (m GetSqlResultsDownloadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetSqlResultsDownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSqlResultsDownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetSqlResultsDownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": m.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetSqlResultsDownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

type GetStatusRequest_SdkV2 struct {
	Keys types.String `tfsdk:"-"`
}

func (to *GetStatusRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetStatusRequest_SdkV2) {
}

func (to *GetStatusRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetStatusRequest_SdkV2) {
}

func (m GetStatusRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"keys": m.Keys,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"keys": types.StringType,
		},
	}
}

type GetTokenManagementRequest_SdkV2 struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-"`
}

func (to *GetTokenManagementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenManagementRequest_SdkV2) {
}

func (to *GetTokenManagementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetTokenManagementRequest_SdkV2) {
}

func (m GetTokenManagementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetTokenManagementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenManagementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetTokenManagementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenManagementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type GetTokenPermissionLevelsRequest_SdkV2 struct {
}

func (to *GetTokenPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenPermissionLevelsRequest_SdkV2) {
}

func (to *GetTokenPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetTokenPermissionLevelsRequest_SdkV2) {
}

func (m GetTokenPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetTokenPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type GetTokenPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetTokenPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetTokenPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetTokenPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetTokenPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetTokenPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(TokenPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetTokenPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: TokenPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetTokenPermissionLevelsResponse_SdkV2 as
// a slice of TokenPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetTokenPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]TokenPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []TokenPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetTokenPermissionLevelsResponse_SdkV2.
func (m *GetTokenPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []TokenPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetTokenPermissionsRequest_SdkV2 struct {
}

func (to *GetTokenPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenPermissionsRequest_SdkV2) {
}

func (to *GetTokenPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetTokenPermissionsRequest_SdkV2) {
}

func (m GetTokenPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetTokenPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse_SdkV2 struct {
	TokenInfo types.List `tfsdk:"token_info"`
}

func (to *GetTokenResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetTokenResponse_SdkV2) {
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

func (to *GetTokenResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetTokenResponse_SdkV2) {
	if !from.TokenInfo.IsNull() && !from.TokenInfo.IsUnknown() {
		if toTokenInfo, ok := to.GetTokenInfo(ctx); ok {
			if fromTokenInfo, ok := from.GetTokenInfo(ctx); ok {
				toTokenInfo.SyncFieldsDuringRead(ctx, fromTokenInfo)
				to.SetTokenInfo(ctx, toTokenInfo)
			}
		}
	}
}

func (m GetTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token_info"] = attrs["token_info"].SetOptional()
	attrs["token_info"] = attrs["token_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info": m.TokenInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": basetypes.ListType{
				ElemType: TokenInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in GetTokenResponse_SdkV2 as
// a TokenInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *GetTokenResponse_SdkV2) GetTokenInfo(ctx context.Context) (TokenInfo_SdkV2, bool) {
	var e TokenInfo_SdkV2
	if m.TokenInfo.IsNull() || m.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []TokenInfo_SdkV2
	d := m.TokenInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenInfo sets the value of the TokenInfo field in GetTokenResponse_SdkV2.
func (m *GetTokenResponse_SdkV2) SetTokenInfo(ctx context.Context, v TokenInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_info"]
	m.TokenInfo = types.ListValueMust(t, vs)
}

type GetWorkspaceNetworkOptionRequest_SdkV2 struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

func (to *GetWorkspaceNetworkOptionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetWorkspaceNetworkOptionRequest_SdkV2) {
}

func (to *GetWorkspaceNetworkOptionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetWorkspaceNetworkOptionRequest_SdkV2) {
}

func (m GetWorkspaceNetworkOptionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m GetWorkspaceNetworkOptionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceNetworkOptionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetWorkspaceNetworkOptionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetWorkspaceNetworkOptionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

// Definition of an IP Access list
type IpAccessListInfo_SdkV2 struct {
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

func (to *IpAccessListInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from IpAccessListInfo_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *IpAccessListInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from IpAccessListInfo_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m IpAccessListInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m IpAccessListInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessListInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m IpAccessListInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m IpAccessListInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetIpAddresses returns the value of the IpAddresses field in IpAccessListInfo_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *IpAccessListInfo_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in IpAccessListInfo_SdkV2.
func (m *IpAccessListInfo_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse_SdkV2 struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (to *ListIpAccessListResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIpAccessListResponse_SdkV2) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (to *ListIpAccessListResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIpAccessListResponse_SdkV2) {
	if !from.IpAccessLists.IsNull() && !from.IpAccessLists.IsUnknown() && to.IpAccessLists.IsNull() && len(from.IpAccessLists.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAccessLists, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAccessLists = from.IpAccessLists
	}
}

func (m ListIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": m.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_lists": basetypes.ListType{
				ElemType: IpAccessListInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetIpAccessLists returns the value of the IpAccessLists field in ListIpAccessListResponse_SdkV2 as
// a slice of IpAccessListInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListIpAccessListResponse_SdkV2) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo_SdkV2, bool) {
	if m.IpAccessLists.IsNull() || m.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo_SdkV2
	d := m.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in ListIpAccessListResponse_SdkV2.
func (m *ListIpAccessListResponse_SdkV2) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAccessLists = types.ListValueMust(t, vs)
}

type ListIpAccessLists_SdkV2 struct {
}

func (to *ListIpAccessLists_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListIpAccessLists_SdkV2) {
}

func (to *ListIpAccessLists_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListIpAccessLists_SdkV2) {
}

func (m ListIpAccessLists_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListIpAccessLists.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListIpAccessLists_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIpAccessLists_SdkV2
// only implements ToObjectValue() and Type().
func (m ListIpAccessLists_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListIpAccessLists_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type ListNetworkConnectivityConfigurationsRequest_SdkV2 struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListNetworkConnectivityConfigurationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkConnectivityConfigurationsRequest_SdkV2) {
}

func (to *ListNetworkConnectivityConfigurationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNetworkConnectivityConfigurationsRequest_SdkV2) {
}

func (m ListNetworkConnectivityConfigurationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNetworkConnectivityConfigurationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNetworkConnectivityConfigurationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkConnectivityConfigurationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

// The network connectivity configuration list was successfully retrieved.
type ListNetworkConnectivityConfigurationsResponse_SdkV2 struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListNetworkConnectivityConfigurationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkConnectivityConfigurationsResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListNetworkConnectivityConfigurationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNetworkConnectivityConfigurationsResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListNetworkConnectivityConfigurationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNetworkConnectivityConfigurationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NetworkConnectivityConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNetworkConnectivityConfigurationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkConnectivityConfigurationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NetworkConnectivityConfiguration_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListNetworkConnectivityConfigurationsResponse_SdkV2 as
// a slice of NetworkConnectivityConfiguration_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNetworkConnectivityConfigurationsResponse_SdkV2) GetItems(ctx context.Context) ([]NetworkConnectivityConfiguration_SdkV2, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []NetworkConnectivityConfiguration_SdkV2
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkConnectivityConfigurationsResponse_SdkV2.
func (m *ListNetworkConnectivityConfigurationsResponse_SdkV2) SetItems(ctx context.Context, v []NetworkConnectivityConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListNetworkPoliciesRequest_SdkV2 struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListNetworkPoliciesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkPoliciesRequest_SdkV2) {
}

func (to *ListNetworkPoliciesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNetworkPoliciesRequest_SdkV2) {
}

func (m ListNetworkPoliciesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNetworkPoliciesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkPoliciesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNetworkPoliciesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkPoliciesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListNetworkPoliciesResponse_SdkV2 struct {
	// List of network policies.
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListNetworkPoliciesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNetworkPoliciesResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListNetworkPoliciesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNetworkPoliciesResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListNetworkPoliciesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNetworkPoliciesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AccountNetworkPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkPoliciesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNetworkPoliciesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNetworkPoliciesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: AccountNetworkPolicy_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListNetworkPoliciesResponse_SdkV2 as
// a slice of AccountNetworkPolicy_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNetworkPoliciesResponse_SdkV2) GetItems(ctx context.Context) ([]AccountNetworkPolicy_SdkV2, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []AccountNetworkPolicy_SdkV2
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkPoliciesResponse_SdkV2.
func (m *ListNetworkPoliciesResponse_SdkV2) SetItems(ctx context.Context, v []AccountNetworkPolicy_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListNotificationDestinationsRequest_SdkV2 struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (to *ListNotificationDestinationsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNotificationDestinationsRequest_SdkV2) {
}

func (to *ListNotificationDestinationsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNotificationDestinationsRequest_SdkV2) {
}

func (m ListNotificationDestinationsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNotificationDestinationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNotificationDestinationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  m.PageSize,
			"page_token": m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNotificationDestinationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":  types.Int64Type,
			"page_token": types.StringType,
		},
	}
}

type ListNotificationDestinationsResponse_SdkV2 struct {
	// Page token for next of results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Results types.List `tfsdk:"results"`
}

func (to *ListNotificationDestinationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNotificationDestinationsResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (to *ListNotificationDestinationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNotificationDestinationsResponse_SdkV2) {
	if !from.Results.IsNull() && !from.Results.IsUnknown() && to.Results.IsNull() && len(from.Results.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Results, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Results = from.Results
	}
}

func (m ListNotificationDestinationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNotificationDestinationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListNotificationDestinationsResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNotificationDestinationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": m.NextPageToken,
			"results":         m.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNotificationDestinationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"results": basetypes.ListType{
				ElemType: ListNotificationDestinationsResult_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetResults returns the value of the Results field in ListNotificationDestinationsResponse_SdkV2 as
// a slice of ListNotificationDestinationsResult_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListNotificationDestinationsResponse_SdkV2) GetResults(ctx context.Context) ([]ListNotificationDestinationsResult_SdkV2, bool) {
	if m.Results.IsNull() || m.Results.IsUnknown() {
		return nil, false
	}
	var v []ListNotificationDestinationsResult_SdkV2
	d := m.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListNotificationDestinationsResponse_SdkV2.
func (m *ListNotificationDestinationsResponse_SdkV2) SetResults(ctx context.Context, v []ListNotificationDestinationsResult_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Results = types.ListValueMust(t, vs)
}

type ListNotificationDestinationsResult_SdkV2 struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id"`
}

func (to *ListNotificationDestinationsResult_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListNotificationDestinationsResult_SdkV2) {
}

func (to *ListNotificationDestinationsResult_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListNotificationDestinationsResult_SdkV2) {
}

func (m ListNotificationDestinationsResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListNotificationDestinationsResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResult_SdkV2
// only implements ToObjectValue() and Type().
func (m ListNotificationDestinationsResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_type": m.DestinationType,
			"display_name":     m.DisplayName,
			"id":               m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListNotificationDestinationsResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
		},
	}
}

type ListPrivateEndpointRulesRequest_SdkV2 struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

func (to *ListPrivateEndpointRulesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateEndpointRulesRequest_SdkV2) {
}

func (to *ListPrivateEndpointRulesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPrivateEndpointRulesRequest_SdkV2) {
}

func (m ListPrivateEndpointRulesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPrivateEndpointRulesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPrivateEndpointRulesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"page_token":                     m.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateEndpointRulesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"page_token":                     types.StringType,
		},
	}
}

// The private endpoint rule list was successfully retrieved.
type ListPrivateEndpointRulesResponse_SdkV2 struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (to *ListPrivateEndpointRulesResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPrivateEndpointRulesResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (to *ListPrivateEndpointRulesResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPrivateEndpointRulesResponse_SdkV2) {
	if !from.Items.IsNull() && !from.Items.IsUnknown() && to.Items.IsNull() && len(from.Items.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Items, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Items = from.Items
	}
}

func (m ListPrivateEndpointRulesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPrivateEndpointRulesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NccPrivateEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPrivateEndpointRulesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           m.Items,
			"next_page_token": m.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPrivateEndpointRulesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NccPrivateEndpointRule_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListPrivateEndpointRulesResponse_SdkV2 as
// a slice of NccPrivateEndpointRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPrivateEndpointRulesResponse_SdkV2) GetItems(ctx context.Context) ([]NccPrivateEndpointRule_SdkV2, bool) {
	if m.Items.IsNull() || m.Items.IsUnknown() {
		return nil, false
	}
	var v []NccPrivateEndpointRule_SdkV2
	d := m.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListPrivateEndpointRulesResponse_SdkV2.
func (m *ListPrivateEndpointRulesResponse_SdkV2) SetItems(ctx context.Context, v []NccPrivateEndpointRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Items = types.ListValueMust(t, vs)
}

type ListPublicTokensResponse_SdkV2 struct {
	// The information for each token.
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (to *ListPublicTokensResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListPublicTokensResponse_SdkV2) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (to *ListPublicTokensResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListPublicTokensResponse_SdkV2) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (m ListPublicTokensResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListPublicTokensResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(PublicTokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublicTokensResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListPublicTokensResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": m.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListPublicTokensResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_infos": basetypes.ListType{
				ElemType: PublicTokenInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTokenInfos returns the value of the TokenInfos field in ListPublicTokensResponse_SdkV2 as
// a slice of PublicTokenInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListPublicTokensResponse_SdkV2) GetTokenInfos(ctx context.Context) ([]PublicTokenInfo_SdkV2, bool) {
	if m.TokenInfos.IsNull() || m.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []PublicTokenInfo_SdkV2
	d := m.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListPublicTokensResponse_SdkV2.
func (m *ListPublicTokensResponse_SdkV2) SetTokenInfos(ctx context.Context, v []PublicTokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TokenInfos = types.ListValueMust(t, vs)
}

type ListTokenManagementRequest_SdkV2 struct {
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"-"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"-"`
}

func (to *ListTokenManagementRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTokenManagementRequest_SdkV2) {
}

func (to *ListTokenManagementRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTokenManagementRequest_SdkV2) {
}

func (m ListTokenManagementRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTokenManagementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokenManagementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTokenManagementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_by_id":       m.CreatedById,
			"created_by_username": m.CreatedByUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTokenManagementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
		},
	}
}

type ListTokens_SdkV2 struct {
}

func (to *ListTokens_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTokens_SdkV2) {
}

func (to *ListTokens_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTokens_SdkV2) {
}

func (m ListTokens_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokens.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListTokens_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokens_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTokens_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListTokens_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Tokens were successfully returned.
type ListTokensResponse_SdkV2 struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (to *ListTokensResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListTokensResponse_SdkV2) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (to *ListTokensResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListTokensResponse_SdkV2) {
	if !from.TokenInfos.IsNull() && !from.TokenInfos.IsUnknown() && to.TokenInfos.IsNull() && len(from.TokenInfos.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for TokenInfos, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.TokenInfos = from.TokenInfos
	}
}

func (m ListTokensResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m ListTokensResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokensResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListTokensResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": m.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListTokensResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_infos": basetypes.ListType{
				ElemType: TokenInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTokenInfos returns the value of the TokenInfos field in ListTokensResponse_SdkV2 as
// a slice of TokenInfo_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListTokensResponse_SdkV2) GetTokenInfos(ctx context.Context) ([]TokenInfo_SdkV2, bool) {
	if m.TokenInfos.IsNull() || m.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo_SdkV2
	d := m.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListTokensResponse_SdkV2.
func (m *ListTokensResponse_SdkV2) SetTokenInfos(ctx context.Context, v []TokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TokenInfos = types.ListValueMust(t, vs)
}

type LlmProxyPartnerPoweredAccount_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
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

func (to *LlmProxyPartnerPoweredAccount_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LlmProxyPartnerPoweredAccount_SdkV2) {
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

func (to *LlmProxyPartnerPoweredAccount_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LlmProxyPartnerPoweredAccount_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m LlmProxyPartnerPoweredAccount_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m LlmProxyPartnerPoweredAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredAccount_SdkV2
// only implements ToObjectValue() and Type().
func (m LlmProxyPartnerPoweredAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LlmProxyPartnerPoweredAccount_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in LlmProxyPartnerPoweredAccount_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LlmProxyPartnerPoweredAccount_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredAccount_SdkV2.
func (m *LlmProxyPartnerPoweredAccount_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

type LlmProxyPartnerPoweredEnforce_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
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

func (to *LlmProxyPartnerPoweredEnforce_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LlmProxyPartnerPoweredEnforce_SdkV2) {
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

func (to *LlmProxyPartnerPoweredEnforce_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LlmProxyPartnerPoweredEnforce_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m LlmProxyPartnerPoweredEnforce_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m LlmProxyPartnerPoweredEnforce_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredEnforce_SdkV2
// only implements ToObjectValue() and Type().
func (m LlmProxyPartnerPoweredEnforce_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LlmProxyPartnerPoweredEnforce_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in LlmProxyPartnerPoweredEnforce_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LlmProxyPartnerPoweredEnforce_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredEnforce_SdkV2.
func (m *LlmProxyPartnerPoweredEnforce_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

type LlmProxyPartnerPoweredWorkspace_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
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

func (to *LlmProxyPartnerPoweredWorkspace_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LlmProxyPartnerPoweredWorkspace_SdkV2) {
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

func (to *LlmProxyPartnerPoweredWorkspace_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LlmProxyPartnerPoweredWorkspace_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m LlmProxyPartnerPoweredWorkspace_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m LlmProxyPartnerPoweredWorkspace_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredWorkspace_SdkV2
// only implements ToObjectValue() and Type().
func (m LlmProxyPartnerPoweredWorkspace_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LlmProxyPartnerPoweredWorkspace_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in LlmProxyPartnerPoweredWorkspace_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *LlmProxyPartnerPoweredWorkspace_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredWorkspace_SdkV2.
func (m *LlmProxyPartnerPoweredWorkspace_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

type MicrosoftTeamsConfig_SdkV2 struct {
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

func (to *MicrosoftTeamsConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from MicrosoftTeamsConfig_SdkV2) {
}

func (to *MicrosoftTeamsConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from MicrosoftTeamsConfig_SdkV2) {
}

func (m MicrosoftTeamsConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m MicrosoftTeamsConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MicrosoftTeamsConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m MicrosoftTeamsConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m MicrosoftTeamsConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
type NccAwsStableIpRule_SdkV2 struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks types.List `tfsdk:"cidr_blocks"`
}

func (to *NccAwsStableIpRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccAwsStableIpRule_SdkV2) {
	if !from.CidrBlocks.IsNull() && !from.CidrBlocks.IsUnknown() && to.CidrBlocks.IsNull() && len(from.CidrBlocks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CidrBlocks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CidrBlocks = from.CidrBlocks
	}
}

func (to *NccAwsStableIpRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccAwsStableIpRule_SdkV2) {
	if !from.CidrBlocks.IsNull() && !from.CidrBlocks.IsUnknown() && to.CidrBlocks.IsNull() && len(from.CidrBlocks.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for CidrBlocks, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.CidrBlocks = from.CidrBlocks
	}
}

func (m NccAwsStableIpRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NccAwsStableIpRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cidr_blocks": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAwsStableIpRule_SdkV2
// only implements ToObjectValue() and Type().
func (m NccAwsStableIpRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cidr_blocks": m.CidrBlocks,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccAwsStableIpRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cidr_blocks": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetCidrBlocks returns the value of the CidrBlocks field in NccAwsStableIpRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAwsStableIpRule_SdkV2) GetCidrBlocks(ctx context.Context) ([]types.String, bool) {
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

// SetCidrBlocks sets the value of the CidrBlocks field in NccAwsStableIpRule_SdkV2.
func (m *NccAwsStableIpRule_SdkV2) SetCidrBlocks(ctx context.Context, v []types.String) {
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
type NccAzurePrivateEndpointRule_SdkV2 struct {
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

func (to *NccAzurePrivateEndpointRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccAzurePrivateEndpointRule_SdkV2) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
}

func (to *NccAzurePrivateEndpointRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccAzurePrivateEndpointRule_SdkV2) {
	if !from.DomainNames.IsNull() && !from.DomainNames.IsUnknown() && to.DomainNames.IsNull() && len(from.DomainNames.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DomainNames, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DomainNames = from.DomainNames
	}
}

func (m NccAzurePrivateEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NccAzurePrivateEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzurePrivateEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (m NccAzurePrivateEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m NccAzurePrivateEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetDomainNames returns the value of the DomainNames field in NccAzurePrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAzurePrivateEndpointRule_SdkV2) GetDomainNames(ctx context.Context) ([]types.String, bool) {
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

// SetDomainNames sets the value of the DomainNames field in NccAzurePrivateEndpointRule_SdkV2.
func (m *NccAzurePrivateEndpointRule_SdkV2) SetDomainNames(ctx context.Context, v []types.String) {
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
type NccAzureServiceEndpointRule_SdkV2 struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets types.List `tfsdk:"subnets"`
	// The Azure region in which this service endpoint rule applies..
	TargetRegion types.String `tfsdk:"target_region"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices types.List `tfsdk:"target_services"`
}

func (to *NccAzureServiceEndpointRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccAzureServiceEndpointRule_SdkV2) {
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

func (to *NccAzureServiceEndpointRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccAzureServiceEndpointRule_SdkV2) {
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

func (m NccAzureServiceEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NccAzureServiceEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subnets":         reflect.TypeOf(types.String{}),
		"target_services": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzureServiceEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (m NccAzureServiceEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"subnets":         m.Subnets,
			"target_region":   m.TargetRegion,
			"target_services": m.TargetServices,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccAzureServiceEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetSubnets returns the value of the Subnets field in NccAzureServiceEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAzureServiceEndpointRule_SdkV2) GetSubnets(ctx context.Context) ([]types.String, bool) {
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

// SetSubnets sets the value of the Subnets field in NccAzureServiceEndpointRule_SdkV2.
func (m *NccAzureServiceEndpointRule_SdkV2) SetSubnets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Subnets = types.ListValueMust(t, vs)
}

// GetTargetServices returns the value of the TargetServices field in NccAzureServiceEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccAzureServiceEndpointRule_SdkV2) GetTargetServices(ctx context.Context) ([]types.String, bool) {
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

// SetTargetServices sets the value of the TargetServices field in NccAzureServiceEndpointRule_SdkV2.
func (m *NccAzureServiceEndpointRule_SdkV2) SetTargetServices(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["target_services"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.TargetServices = types.ListValueMust(t, vs)
}

type NccEgressConfig_SdkV2 struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules types.List `tfsdk:"default_rules"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules types.List `tfsdk:"target_rules"`
}

func (to *NccEgressConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccEgressConfig_SdkV2) {
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

func (to *NccEgressConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccEgressConfig_SdkV2) {
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

func (m NccEgressConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_rules"] = attrs["default_rules"].SetOptional()
	attrs["default_rules"] = attrs["default_rules"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["target_rules"] = attrs["target_rules"].SetOptional()
	attrs["target_rules"] = attrs["target_rules"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccEgressConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_rules": reflect.TypeOf(NccEgressDefaultRules_SdkV2{}),
		"target_rules":  reflect.TypeOf(NccEgressTargetRules_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m NccEgressConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_rules": m.DefaultRules,
			"target_rules":  m.TargetRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccEgressConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_rules": basetypes.ListType{
				ElemType: NccEgressDefaultRules_SdkV2{}.Type(ctx),
			},
			"target_rules": basetypes.ListType{
				ElemType: NccEgressTargetRules_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetDefaultRules returns the value of the DefaultRules field in NccEgressConfig_SdkV2 as
// a NccEgressDefaultRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressConfig_SdkV2) GetDefaultRules(ctx context.Context) (NccEgressDefaultRules_SdkV2, bool) {
	var e NccEgressDefaultRules_SdkV2
	if m.DefaultRules.IsNull() || m.DefaultRules.IsUnknown() {
		return e, false
	}
	var v []NccEgressDefaultRules_SdkV2
	d := m.DefaultRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultRules sets the value of the DefaultRules field in NccEgressConfig_SdkV2.
func (m *NccEgressConfig_SdkV2) SetDefaultRules(ctx context.Context, v NccEgressDefaultRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["default_rules"]
	m.DefaultRules = types.ListValueMust(t, vs)
}

// GetTargetRules returns the value of the TargetRules field in NccEgressConfig_SdkV2 as
// a NccEgressTargetRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressConfig_SdkV2) GetTargetRules(ctx context.Context) (NccEgressTargetRules_SdkV2, bool) {
	var e NccEgressTargetRules_SdkV2
	if m.TargetRules.IsNull() || m.TargetRules.IsUnknown() {
		return e, false
	}
	var v []NccEgressTargetRules_SdkV2
	d := m.TargetRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTargetRules sets the value of the TargetRules field in NccEgressConfig_SdkV2.
func (m *NccEgressConfig_SdkV2) SetTargetRules(ctx context.Context, v NccEgressTargetRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["target_rules"]
	m.TargetRules = types.ListValueMust(t, vs)
}

// Default rules don't have specific targets.
type NccEgressDefaultRules_SdkV2 struct {
	AwsStableIpRule types.List `tfsdk:"aws_stable_ip_rule"`

	AzureServiceEndpointRule types.List `tfsdk:"azure_service_endpoint_rule"`
}

func (to *NccEgressDefaultRules_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccEgressDefaultRules_SdkV2) {
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

func (to *NccEgressDefaultRules_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccEgressDefaultRules_SdkV2) {
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

func (m NccEgressDefaultRules_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_stable_ip_rule"] = attrs["aws_stable_ip_rule"].SetOptional()
	attrs["aws_stable_ip_rule"] = attrs["aws_stable_ip_rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["azure_service_endpoint_rule"] = attrs["azure_service_endpoint_rule"].SetOptional()
	attrs["azure_service_endpoint_rule"] = attrs["azure_service_endpoint_rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccEgressDefaultRules.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NccEgressDefaultRules_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_stable_ip_rule":          reflect.TypeOf(NccAwsStableIpRule_SdkV2{}),
		"azure_service_endpoint_rule": reflect.TypeOf(NccAzureServiceEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressDefaultRules_SdkV2
// only implements ToObjectValue() and Type().
func (m NccEgressDefaultRules_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_stable_ip_rule":          m.AwsStableIpRule,
			"azure_service_endpoint_rule": m.AzureServiceEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccEgressDefaultRules_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_stable_ip_rule": basetypes.ListType{
				ElemType: NccAwsStableIpRule_SdkV2{}.Type(ctx),
			},
			"azure_service_endpoint_rule": basetypes.ListType{
				ElemType: NccAzureServiceEndpointRule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAwsStableIpRule returns the value of the AwsStableIpRule field in NccEgressDefaultRules_SdkV2 as
// a NccAwsStableIpRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressDefaultRules_SdkV2) GetAwsStableIpRule(ctx context.Context) (NccAwsStableIpRule_SdkV2, bool) {
	var e NccAwsStableIpRule_SdkV2
	if m.AwsStableIpRule.IsNull() || m.AwsStableIpRule.IsUnknown() {
		return e, false
	}
	var v []NccAwsStableIpRule_SdkV2
	d := m.AwsStableIpRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsStableIpRule sets the value of the AwsStableIpRule field in NccEgressDefaultRules_SdkV2.
func (m *NccEgressDefaultRules_SdkV2) SetAwsStableIpRule(ctx context.Context, v NccAwsStableIpRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_stable_ip_rule"]
	m.AwsStableIpRule = types.ListValueMust(t, vs)
}

// GetAzureServiceEndpointRule returns the value of the AzureServiceEndpointRule field in NccEgressDefaultRules_SdkV2 as
// a NccAzureServiceEndpointRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressDefaultRules_SdkV2) GetAzureServiceEndpointRule(ctx context.Context) (NccAzureServiceEndpointRule_SdkV2, bool) {
	var e NccAzureServiceEndpointRule_SdkV2
	if m.AzureServiceEndpointRule.IsNull() || m.AzureServiceEndpointRule.IsUnknown() {
		return e, false
	}
	var v []NccAzureServiceEndpointRule_SdkV2
	d := m.AzureServiceEndpointRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServiceEndpointRule sets the value of the AzureServiceEndpointRule field in NccEgressDefaultRules_SdkV2.
func (m *NccEgressDefaultRules_SdkV2) SetAzureServiceEndpointRule(ctx context.Context, v NccAzureServiceEndpointRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_endpoint_rule"]
	m.AzureServiceEndpointRule = types.ListValueMust(t, vs)
}

// Target rule controls the egress rules that are dedicated to specific
// resources.
type NccEgressTargetRules_SdkV2 struct {
	// AWS private endpoint rule controls the AWS private endpoint based egress
	// rules.
	AwsPrivateEndpointRules types.List `tfsdk:"aws_private_endpoint_rules"`

	AzurePrivateEndpointRules types.List `tfsdk:"azure_private_endpoint_rules"`
}

func (to *NccEgressTargetRules_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccEgressTargetRules_SdkV2) {
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

func (to *NccEgressTargetRules_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccEgressTargetRules_SdkV2) {
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

func (m NccEgressTargetRules_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NccEgressTargetRules_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_private_endpoint_rules":   reflect.TypeOf(CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2{}),
		"azure_private_endpoint_rules": reflect.TypeOf(NccAzurePrivateEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressTargetRules_SdkV2
// only implements ToObjectValue() and Type().
func (m NccEgressTargetRules_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_private_endpoint_rules":   m.AwsPrivateEndpointRules,
			"azure_private_endpoint_rules": m.AzurePrivateEndpointRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NccEgressTargetRules_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_private_endpoint_rules": basetypes.ListType{
				ElemType: CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2{}.Type(ctx),
			},
			"azure_private_endpoint_rules": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAwsPrivateEndpointRules returns the value of the AwsPrivateEndpointRules field in NccEgressTargetRules_SdkV2 as
// a slice of CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressTargetRules_SdkV2) GetAwsPrivateEndpointRules(ctx context.Context) ([]CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2, bool) {
	if m.AwsPrivateEndpointRules.IsNull() || m.AwsPrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2
	d := m.AwsPrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsPrivateEndpointRules sets the value of the AwsPrivateEndpointRules field in NccEgressTargetRules_SdkV2.
func (m *NccEgressTargetRules_SdkV2) SetAwsPrivateEndpointRules(ctx context.Context, v []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_private_endpoint_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AwsPrivateEndpointRules = types.ListValueMust(t, vs)
}

// GetAzurePrivateEndpointRules returns the value of the AzurePrivateEndpointRules field in NccEgressTargetRules_SdkV2 as
// a slice of NccAzurePrivateEndpointRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccEgressTargetRules_SdkV2) GetAzurePrivateEndpointRules(ctx context.Context) ([]NccAzurePrivateEndpointRule_SdkV2, bool) {
	if m.AzurePrivateEndpointRules.IsNull() || m.AzurePrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []NccAzurePrivateEndpointRule_SdkV2
	d := m.AzurePrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzurePrivateEndpointRules sets the value of the AzurePrivateEndpointRules field in NccEgressTargetRules_SdkV2.
func (m *NccEgressTargetRules_SdkV2) SetAzurePrivateEndpointRules(ctx context.Context, v []NccAzurePrivateEndpointRule_SdkV2) {
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
type NccPrivateEndpointRule_SdkV2 struct {
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

func (to *NccPrivateEndpointRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NccPrivateEndpointRule_SdkV2) {
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

func (to *NccPrivateEndpointRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NccPrivateEndpointRule_SdkV2) {
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

func (m NccPrivateEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m NccPrivateEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccPrivateEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (m NccPrivateEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
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
func (m NccPrivateEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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

// GetDomainNames returns the value of the DomainNames field in NccPrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccPrivateEndpointRule_SdkV2) GetDomainNames(ctx context.Context) ([]types.String, bool) {
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

// SetDomainNames sets the value of the DomainNames field in NccPrivateEndpointRule_SdkV2.
func (m *NccPrivateEndpointRule_SdkV2) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in NccPrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *NccPrivateEndpointRule_SdkV2) GetResourceNames(ctx context.Context) ([]types.String, bool) {
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

// SetResourceNames sets the value of the ResourceNames field in NccPrivateEndpointRule_SdkV2.
func (m *NccPrivateEndpointRule_SdkV2) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

// Properties of the new network connectivity configuration.
type NetworkConnectivityConfiguration_SdkV2 struct {
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig types.List `tfsdk:"egress_config"`
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

func (to *NetworkConnectivityConfiguration_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkConnectivityConfiguration_SdkV2) {
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

func (to *NetworkConnectivityConfiguration_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NetworkConnectivityConfiguration_SdkV2) {
	if !from.EgressConfig.IsNull() && !from.EgressConfig.IsUnknown() {
		if toEgressConfig, ok := to.GetEgressConfig(ctx); ok {
			if fromEgressConfig, ok := from.GetEgressConfig(ctx); ok {
				toEgressConfig.SyncFieldsDuringRead(ctx, fromEgressConfig)
				to.SetEgressConfig(ctx, toEgressConfig)
			}
		}
	}
}

func (m NetworkConnectivityConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_time"] = attrs["creation_time"].SetOptional()
	attrs["egress_config"] = attrs["egress_config"].SetOptional()
	attrs["egress_config"] = attrs["egress_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m NetworkConnectivityConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress_config": reflect.TypeOf(NccEgressConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkConnectivityConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (m NetworkConnectivityConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_time":                  m.CreationTime,
			"egress_config":                  m.EgressConfig,
			"name":                           m.Name,
			"network_connectivity_config_id": m.NetworkConnectivityConfigId,
			"region":                         m.Region,
			"updated_time":                   m.UpdatedTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkConnectivityConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_time": types.Int64Type,
			"egress_config": basetypes.ListType{
				ElemType: NccEgressConfig_SdkV2{}.Type(ctx),
			},
			"name":                           types.StringType,
			"network_connectivity_config_id": types.StringType,
			"region":                         types.StringType,
			"updated_time":                   types.Int64Type,
		},
	}
}

// GetEgressConfig returns the value of the EgressConfig field in NetworkConnectivityConfiguration_SdkV2 as
// a NccEgressConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NetworkConnectivityConfiguration_SdkV2) GetEgressConfig(ctx context.Context) (NccEgressConfig_SdkV2, bool) {
	var e NccEgressConfig_SdkV2
	if m.EgressConfig.IsNull() || m.EgressConfig.IsUnknown() {
		return e, false
	}
	var v []NccEgressConfig_SdkV2
	d := m.EgressConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEgressConfig sets the value of the EgressConfig field in NetworkConnectivityConfiguration_SdkV2.
func (m *NetworkConnectivityConfiguration_SdkV2) SetEgressConfig(ctx context.Context, v NccEgressConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["egress_config"]
	m.EgressConfig = types.ListValueMust(t, vs)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto). This policy should be consistent
// with [[com.databricks.api.proto.settingspolicy.EgressNetworkPolicy]]. Details
// see API-design:
// https://docs.google.com/document/d/1DKWO_FpZMCY4cF2O62LpwII1lx8gsnDGG-qgE3t3TOA/
type NetworkPolicyEgress_SdkV2 struct {
	// The access policy enforced for egress traffic to the internet.
	NetworkAccess types.List `tfsdk:"network_access"`
}

func (to *NetworkPolicyEgress_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NetworkPolicyEgress_SdkV2) {
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

func (to *NetworkPolicyEgress_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NetworkPolicyEgress_SdkV2) {
	if !from.NetworkAccess.IsNull() && !from.NetworkAccess.IsUnknown() {
		if toNetworkAccess, ok := to.GetNetworkAccess(ctx); ok {
			if fromNetworkAccess, ok := from.GetNetworkAccess(ctx); ok {
				toNetworkAccess.SyncFieldsDuringRead(ctx, fromNetworkAccess)
				to.SetNetworkAccess(ctx, toNetworkAccess)
			}
		}
	}
}

func (m NetworkPolicyEgress_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_access"] = attrs["network_access"].SetOptional()
	attrs["network_access"] = attrs["network_access"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkPolicyEgress.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m NetworkPolicyEgress_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_access": reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkPolicyEgress_SdkV2
// only implements ToObjectValue() and Type().
func (m NetworkPolicyEgress_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_access": m.NetworkAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (m NetworkPolicyEgress_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_access": basetypes.ListType{
				ElemType: EgressNetworkPolicyNetworkAccessPolicy_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetNetworkAccess returns the value of the NetworkAccess field in NetworkPolicyEgress_SdkV2 as
// a EgressNetworkPolicyNetworkAccessPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NetworkPolicyEgress_SdkV2) GetNetworkAccess(ctx context.Context) (EgressNetworkPolicyNetworkAccessPolicy_SdkV2, bool) {
	var e EgressNetworkPolicyNetworkAccessPolicy_SdkV2
	if m.NetworkAccess.IsNull() || m.NetworkAccess.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicy_SdkV2
	d := m.NetworkAccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNetworkAccess sets the value of the NetworkAccess field in NetworkPolicyEgress_SdkV2.
func (m *NetworkPolicyEgress_SdkV2) SetNetworkAccess(ctx context.Context, v EgressNetworkPolicyNetworkAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["network_access"]
	m.NetworkAccess = types.ListValueMust(t, vs)
}

type NotificationDestination_SdkV2 struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config types.List `tfsdk:"config"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType types.String `tfsdk:"destination_type"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"id"`
}

func (to *NotificationDestination_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from NotificationDestination_SdkV2) {
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

func (to *NotificationDestination_SdkV2) SyncFieldsDuringRead(ctx context.Context, from NotificationDestination_SdkV2) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m NotificationDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m NotificationDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationDestination_SdkV2
// only implements ToObjectValue() and Type().
func (m NotificationDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m NotificationDestination_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: Config_SdkV2{}.Type(ctx),
			},
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in NotificationDestination_SdkV2 as
// a Config_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *NotificationDestination_SdkV2) GetConfig(ctx context.Context) (Config_SdkV2, bool) {
	var e Config_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []Config_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in NotificationDestination_SdkV2.
func (m *NotificationDestination_SdkV2) SetConfig(ctx context.Context, v Config_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

type PagerdutyConfig_SdkV2 struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey types.String `tfsdk:"integration_key"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet types.Bool `tfsdk:"integration_key_set"`
}

func (to *PagerdutyConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PagerdutyConfig_SdkV2) {
}

func (to *PagerdutyConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PagerdutyConfig_SdkV2) {
}

func (m PagerdutyConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PagerdutyConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PagerdutyConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m PagerdutyConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_key":     m.IntegrationKey,
			"integration_key_set": m.IntegrationKeySet,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PagerdutyConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"integration_key":     types.StringType,
			"integration_key_set": types.BoolType,
		},
	}
}

// Partition by workspace or account
type PartitionId_SdkV2 struct {
	// The ID of the workspace.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *PartitionId_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PartitionId_SdkV2) {
}

func (to *PartitionId_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PartitionId_SdkV2) {
}

func (m PartitionId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PartitionId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionId_SdkV2
// only implements ToObjectValue() and Type().
func (m PartitionId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PartitionId_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
		},
	}
}

type PersonalComputeMessage_SdkV2 struct {
	Value types.String `tfsdk:"value"`
}

func (to *PersonalComputeMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PersonalComputeMessage_SdkV2) {
}

func (to *PersonalComputeMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PersonalComputeMessage_SdkV2) {
}

func (m PersonalComputeMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PersonalComputeMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeMessage_SdkV2
// only implements ToObjectValue() and Type().
func (m PersonalComputeMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PersonalComputeMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type PersonalComputeSetting_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	PersonalCompute types.List `tfsdk:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *PersonalComputeSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PersonalComputeSetting_SdkV2) {
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

func (to *PersonalComputeSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PersonalComputeSetting_SdkV2) {
	if !from.PersonalCompute.IsNull() && !from.PersonalCompute.IsUnknown() {
		if toPersonalCompute, ok := to.GetPersonalCompute(ctx); ok {
			if fromPersonalCompute, ok := from.GetPersonalCompute(ctx); ok {
				toPersonalCompute.SyncFieldsDuringRead(ctx, fromPersonalCompute)
				to.SetPersonalCompute(ctx, toPersonalCompute)
			}
		}
	}
}

func (m PersonalComputeSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["personal_compute"] = attrs["personal_compute"].SetRequired()
	attrs["personal_compute"] = attrs["personal_compute"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m PersonalComputeSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personal_compute": reflect.TypeOf(PersonalComputeMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m PersonalComputeSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             m.Etag,
			"personal_compute": m.PersonalCompute,
			"setting_name":     m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PersonalComputeSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"personal_compute": basetypes.ListType{
				ElemType: PersonalComputeMessage_SdkV2{}.Type(ctx),
			},
			"setting_name": types.StringType,
		},
	}
}

// GetPersonalCompute returns the value of the PersonalCompute field in PersonalComputeSetting_SdkV2 as
// a PersonalComputeMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PersonalComputeSetting_SdkV2) GetPersonalCompute(ctx context.Context) (PersonalComputeMessage_SdkV2, bool) {
	var e PersonalComputeMessage_SdkV2
	if m.PersonalCompute.IsNull() || m.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeMessage_SdkV2
	d := m.PersonalCompute.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPersonalCompute sets the value of the PersonalCompute field in PersonalComputeSetting_SdkV2.
func (m *PersonalComputeSetting_SdkV2) SetPersonalCompute(ctx context.Context, v PersonalComputeMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["personal_compute"]
	m.PersonalCompute = types.ListValueMust(t, vs)
}

type PublicTokenInfo_SdkV2 struct {
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

func (to *PublicTokenInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PublicTokenInfo_SdkV2) {
}

func (to *PublicTokenInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PublicTokenInfo_SdkV2) {
}

func (m PublicTokenInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m PublicTokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublicTokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m PublicTokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m PublicTokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
type ReplaceIpAccessList_SdkV2 struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`

	ListType types.String `tfsdk:"list_type"`
}

func (to *ReplaceIpAccessList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ReplaceIpAccessList_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *ReplaceIpAccessList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ReplaceIpAccessList_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m ReplaceIpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetRequired()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetRequired()
	attrs["list_type"] = attrs["list_type"].SetRequired()
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
func (m ReplaceIpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceIpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (m ReplaceIpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ReplaceIpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetIpAddresses returns the value of the IpAddresses field in ReplaceIpAccessList_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ReplaceIpAccessList_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in ReplaceIpAccessList_SdkV2.
func (m *ReplaceIpAccessList_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

type RestrictWorkspaceAdminsMessage_SdkV2 struct {
	Status types.String `tfsdk:"status"`
}

func (to *RestrictWorkspaceAdminsMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestrictWorkspaceAdminsMessage_SdkV2) {
}

func (to *RestrictWorkspaceAdminsMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestrictWorkspaceAdminsMessage_SdkV2) {
}

func (m RestrictWorkspaceAdminsMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RestrictWorkspaceAdminsMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsMessage_SdkV2
// only implements ToObjectValue() and Type().
func (m RestrictWorkspaceAdminsMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": m.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestrictWorkspaceAdminsMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"status": types.StringType,
		},
	}
}

type RestrictWorkspaceAdminsSetting_SdkV2 struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag types.String `tfsdk:"etag"`

	RestrictWorkspaceAdmins types.List `tfsdk:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName types.String `tfsdk:"setting_name"`
}

func (to *RestrictWorkspaceAdminsSetting_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RestrictWorkspaceAdminsSetting_SdkV2) {
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

func (to *RestrictWorkspaceAdminsSetting_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RestrictWorkspaceAdminsSetting_SdkV2) {
	if !from.RestrictWorkspaceAdmins.IsNull() && !from.RestrictWorkspaceAdmins.IsUnknown() {
		if toRestrictWorkspaceAdmins, ok := to.GetRestrictWorkspaceAdmins(ctx); ok {
			if fromRestrictWorkspaceAdmins, ok := from.GetRestrictWorkspaceAdmins(ctx); ok {
				toRestrictWorkspaceAdmins.SyncFieldsDuringRead(ctx, fromRestrictWorkspaceAdmins)
				to.SetRestrictWorkspaceAdmins(ctx, toRestrictWorkspaceAdmins)
			}
		}
	}
}

func (m RestrictWorkspaceAdminsSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["etag"] = attrs["etag"].SetOptional()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].SetRequired()
	attrs["restrict_workspace_admins"] = attrs["restrict_workspace_admins"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m RestrictWorkspaceAdminsSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"restrict_workspace_admins": reflect.TypeOf(RestrictWorkspaceAdminsMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsSetting_SdkV2
// only implements ToObjectValue() and Type().
func (m RestrictWorkspaceAdminsSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                      m.Etag,
			"restrict_workspace_admins": m.RestrictWorkspaceAdmins,
			"setting_name":              m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RestrictWorkspaceAdminsSetting_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
			"restrict_workspace_admins": basetypes.ListType{
				ElemType: RestrictWorkspaceAdminsMessage_SdkV2{}.Type(ctx),
			},
			"setting_name": types.StringType,
		},
	}
}

// GetRestrictWorkspaceAdmins returns the value of the RestrictWorkspaceAdmins field in RestrictWorkspaceAdminsSetting_SdkV2 as
// a RestrictWorkspaceAdminsMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *RestrictWorkspaceAdminsSetting_SdkV2) GetRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage_SdkV2, bool) {
	var e RestrictWorkspaceAdminsMessage_SdkV2
	if m.RestrictWorkspaceAdmins.IsNull() || m.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsMessage_SdkV2
	d := m.RestrictWorkspaceAdmins.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in RestrictWorkspaceAdminsSetting_SdkV2.
func (m *RestrictWorkspaceAdminsSetting_SdkV2) SetRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["restrict_workspace_admins"]
	m.RestrictWorkspaceAdmins = types.ListValueMust(t, vs)
}

type RevokeTokenRequest_SdkV2 struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id"`
}

func (to *RevokeTokenRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RevokeTokenRequest_SdkV2) {
}

func (to *RevokeTokenRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RevokeTokenRequest_SdkV2) {
}

func (m RevokeTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m RevokeTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m RevokeTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": m.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RevokeTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type RevokeTokenResponse_SdkV2 struct {
}

func (to *RevokeTokenResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RevokeTokenResponse_SdkV2) {
}

func (to *RevokeTokenResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RevokeTokenResponse_SdkV2) {
}

func (m RevokeTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RevokeTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m RevokeTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m RevokeTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SlackConfig_SdkV2 struct {
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

func (to *SlackConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SlackConfig_SdkV2) {
}

func (to *SlackConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SlackConfig_SdkV2) {
}

func (m SlackConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m SlackConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SlackConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m SlackConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m SlackConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

type SqlResultsDownload_SdkV2 struct {
	BooleanVal types.List `tfsdk:"boolean_val"`
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

func (to *SqlResultsDownload_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from SqlResultsDownload_SdkV2) {
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

func (to *SqlResultsDownload_SdkV2) SyncFieldsDuringRead(ctx context.Context, from SqlResultsDownload_SdkV2) {
	if !from.BooleanVal.IsNull() && !from.BooleanVal.IsUnknown() {
		if toBooleanVal, ok := to.GetBooleanVal(ctx); ok {
			if fromBooleanVal, ok := from.GetBooleanVal(ctx); ok {
				toBooleanVal.SyncFieldsDuringRead(ctx, fromBooleanVal)
				to.SetBooleanVal(ctx, toBooleanVal)
			}
		}
	}
}

func (m SqlResultsDownload_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["boolean_val"] = attrs["boolean_val"].SetRequired()
	attrs["boolean_val"] = attrs["boolean_val"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m SqlResultsDownload_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlResultsDownload_SdkV2
// only implements ToObjectValue() and Type().
func (m SqlResultsDownload_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  m.BooleanVal,
			"etag":         m.Etag,
			"setting_name": m.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m SqlResultsDownload_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"boolean_val": basetypes.ListType{
				ElemType: BooleanMessage_SdkV2{}.Type(ctx),
			},
			"etag":         types.StringType,
			"setting_name": types.StringType,
		},
	}
}

// GetBooleanVal returns the value of the BooleanVal field in SqlResultsDownload_SdkV2 as
// a BooleanMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *SqlResultsDownload_SdkV2) GetBooleanVal(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if m.BooleanVal.IsNull() || m.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := m.BooleanVal.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBooleanVal sets the value of the BooleanVal field in SqlResultsDownload_SdkV2.
func (m *SqlResultsDownload_SdkV2) SetBooleanVal(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["boolean_val"]
	m.BooleanVal = types.ListValueMust(t, vs)
}

type StringMessage_SdkV2 struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (to *StringMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from StringMessage_SdkV2) {
}

func (to *StringMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from StringMessage_SdkV2) {
}

func (m StringMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m StringMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage_SdkV2
// only implements ToObjectValue() and Type().
func (m StringMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m StringMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type TokenAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *TokenAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenAccessControlRequest_SdkV2) {
}

func (to *TokenAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenAccessControlRequest_SdkV2) {
}

func (m TokenAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TokenAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type TokenAccessControlResponse_SdkV2 struct {
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

func (to *TokenAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *TokenAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m TokenAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(TokenPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TokenAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: TokenPermission_SdkV2{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in TokenAccessControlResponse_SdkV2 as
// a slice of TokenPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]TokenPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []TokenPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in TokenAccessControlResponse_SdkV2.
func (m *TokenAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []TokenPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type TokenInfo_SdkV2 struct {
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

func (to *TokenInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenInfo_SdkV2) {
}

func (to *TokenInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenInfo_SdkV2) {
}

func (m TokenInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m TokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
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

type TokenPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *TokenPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *TokenPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m TokenPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in TokenPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in TokenPermission_SdkV2.
func (m *TokenPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type TokenPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *TokenPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *TokenPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m TokenPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: TokenAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in TokenPermissions_SdkV2 as
// a slice of TokenAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]TokenAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissions_SdkV2.
func (m *TokenPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []TokenAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type TokenPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *TokenPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermissionsDescription_SdkV2) {
}

func (to *TokenPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenPermissionsDescription_SdkV2) {
}

func (m TokenPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type TokenPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
}

func (to *TokenPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TokenPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *TokenPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TokenPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m TokenPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m TokenPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m TokenPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TokenPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: TokenAccessControlRequest_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in TokenPermissionsRequest_SdkV2 as
// a slice of TokenAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *TokenPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]TokenAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissionsRequest_SdkV2.
func (m *TokenPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []TokenAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateAccountIpAccessEnableRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateAccountIpAccessEnableRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAccountIpAccessEnableRequest_SdkV2) {
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

func (to *UpdateAccountIpAccessEnableRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAccountIpAccessEnableRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAccountIpAccessEnableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAccountIpAccessEnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AccountIpAccessEnable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountIpAccessEnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAccountIpAccessEnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAccountIpAccessEnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AccountIpAccessEnable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAccountIpAccessEnableRequest_SdkV2 as
// a AccountIpAccessEnable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAccountIpAccessEnableRequest_SdkV2) GetSetting(ctx context.Context) (AccountIpAccessEnable_SdkV2, bool) {
	var e AccountIpAccessEnable_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []AccountIpAccessEnable_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAccountIpAccessEnableRequest_SdkV2.
func (m *UpdateAccountIpAccessEnableRequest_SdkV2) SetSetting(ctx context.Context, v AccountIpAccessEnable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
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

func (to *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicySetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingAccessPolicySetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2 as
// a AibiDashboardEmbeddingAccessPolicySetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetSetting(ctx context.Context) (AibiDashboardEmbeddingAccessPolicySetting_SdkV2, bool) {
	var e AibiDashboardEmbeddingAccessPolicySetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicySetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2.
func (m *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SetSetting(ctx context.Context, v AibiDashboardEmbeddingAccessPolicySetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
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

func (to *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2 as
// a AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetSetting(ctx context.Context) (AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2, bool) {
	var e AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2.
func (m *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SetSetting(ctx context.Context, v AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateAutomaticClusterUpdateSettingRequest_SdkV2) {
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

func (to *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateAutomaticClusterUpdateSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateAutomaticClusterUpdateSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateAutomaticClusterUpdateSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AutomaticClusterUpdateSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAutomaticClusterUpdateSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateAutomaticClusterUpdateSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateAutomaticClusterUpdateSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: AutomaticClusterUpdateSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateAutomaticClusterUpdateSettingRequest_SdkV2 as
// a AutomaticClusterUpdateSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) GetSetting(ctx context.Context) (AutomaticClusterUpdateSetting_SdkV2, bool) {
	var e AutomaticClusterUpdateSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []AutomaticClusterUpdateSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAutomaticClusterUpdateSettingRequest_SdkV2.
func (m *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) SetSetting(ctx context.Context, v AutomaticClusterUpdateSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateComplianceSecurityProfileSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateComplianceSecurityProfileSettingRequest_SdkV2) {
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

func (to *UpdateComplianceSecurityProfileSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateComplianceSecurityProfileSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateComplianceSecurityProfileSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateComplianceSecurityProfileSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(ComplianceSecurityProfileSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComplianceSecurityProfileSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateComplianceSecurityProfileSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateComplianceSecurityProfileSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: ComplianceSecurityProfileSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateComplianceSecurityProfileSettingRequest_SdkV2 as
// a ComplianceSecurityProfileSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateComplianceSecurityProfileSettingRequest_SdkV2) GetSetting(ctx context.Context) (ComplianceSecurityProfileSetting_SdkV2, bool) {
	var e ComplianceSecurityProfileSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfileSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateComplianceSecurityProfileSettingRequest_SdkV2.
func (m *UpdateComplianceSecurityProfileSettingRequest_SdkV2) SetSetting(ctx context.Context, v ComplianceSecurityProfileSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateCspEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateCspEnablementAccountSettingRequest_SdkV2) {
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

func (to *UpdateCspEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateCspEnablementAccountSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateCspEnablementAccountSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateCspEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(CspEnablementAccountSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCspEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateCspEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateCspEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: CspEnablementAccountSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateCspEnablementAccountSettingRequest_SdkV2 as
// a CspEnablementAccountSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateCspEnablementAccountSettingRequest_SdkV2) GetSetting(ctx context.Context) (CspEnablementAccountSetting_SdkV2, bool) {
	var e CspEnablementAccountSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []CspEnablementAccountSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateCspEnablementAccountSettingRequest_SdkV2.
func (m *UpdateCspEnablementAccountSettingRequest_SdkV2) SetSetting(ctx context.Context, v CspEnablementAccountSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateDashboardEmailSubscriptionsRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateDashboardEmailSubscriptionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDashboardEmailSubscriptionsRequest_SdkV2) {
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

func (to *UpdateDashboardEmailSubscriptionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDashboardEmailSubscriptionsRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDashboardEmailSubscriptionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDashboardEmailSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDashboardEmailSubscriptionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DashboardEmailSubscriptions_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDashboardEmailSubscriptionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDashboardEmailSubscriptionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDashboardEmailSubscriptionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DashboardEmailSubscriptions_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDashboardEmailSubscriptionsRequest_SdkV2 as
// a DashboardEmailSubscriptions_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDashboardEmailSubscriptionsRequest_SdkV2) GetSetting(ctx context.Context) (DashboardEmailSubscriptions_SdkV2, bool) {
	var e DashboardEmailSubscriptions_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []DashboardEmailSubscriptions_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDashboardEmailSubscriptionsRequest_SdkV2.
func (m *UpdateDashboardEmailSubscriptionsRequest_SdkV2) SetSetting(ctx context.Context, v DashboardEmailSubscriptions_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateDefaultNamespaceSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultNamespaceSettingRequest_SdkV2) {
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

func (to *UpdateDefaultNamespaceSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultNamespaceSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDefaultNamespaceSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultNamespaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultNamespaceSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultNamespaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDefaultNamespaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultNamespaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DefaultNamespaceSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDefaultNamespaceSettingRequest_SdkV2 as
// a DefaultNamespaceSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDefaultNamespaceSettingRequest_SdkV2) GetSetting(ctx context.Context) (DefaultNamespaceSetting_SdkV2, bool) {
	var e DefaultNamespaceSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []DefaultNamespaceSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDefaultNamespaceSettingRequest_SdkV2.
func (m *UpdateDefaultNamespaceSettingRequest_SdkV2) SetSetting(ctx context.Context, v DefaultNamespaceSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateDefaultWarehouseIdRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateDefaultWarehouseIdRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDefaultWarehouseIdRequest_SdkV2) {
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

func (to *UpdateDefaultWarehouseIdRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDefaultWarehouseIdRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDefaultWarehouseIdRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDefaultWarehouseIdRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDefaultWarehouseIdRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultWarehouseId_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultWarehouseIdRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDefaultWarehouseIdRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDefaultWarehouseIdRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DefaultWarehouseId_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDefaultWarehouseIdRequest_SdkV2 as
// a DefaultWarehouseId_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDefaultWarehouseIdRequest_SdkV2) GetSetting(ctx context.Context) (DefaultWarehouseId_SdkV2, bool) {
	var e DefaultWarehouseId_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []DefaultWarehouseId_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDefaultWarehouseIdRequest_SdkV2.
func (m *UpdateDefaultWarehouseIdRequest_SdkV2) SetSetting(ctx context.Context, v DefaultWarehouseId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateDisableLegacyAccessRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateDisableLegacyAccessRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDisableLegacyAccessRequest_SdkV2) {
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

func (to *UpdateDisableLegacyAccessRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDisableLegacyAccessRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDisableLegacyAccessRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDisableLegacyAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyAccess_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDisableLegacyAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDisableLegacyAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DisableLegacyAccess_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDisableLegacyAccessRequest_SdkV2 as
// a DisableLegacyAccess_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDisableLegacyAccessRequest_SdkV2) GetSetting(ctx context.Context) (DisableLegacyAccess_SdkV2, bool) {
	var e DisableLegacyAccess_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyAccess_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyAccessRequest_SdkV2.
func (m *UpdateDisableLegacyAccessRequest_SdkV2) SetSetting(ctx context.Context, v DisableLegacyAccess_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateDisableLegacyDbfsRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateDisableLegacyDbfsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDisableLegacyDbfsRequest_SdkV2) {
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

func (to *UpdateDisableLegacyDbfsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDisableLegacyDbfsRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDisableLegacyDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDisableLegacyDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyDbfs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDisableLegacyDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDisableLegacyDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DisableLegacyDbfs_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDisableLegacyDbfsRequest_SdkV2 as
// a DisableLegacyDbfs_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDisableLegacyDbfsRequest_SdkV2) GetSetting(ctx context.Context) (DisableLegacyDbfs_SdkV2, bool) {
	var e DisableLegacyDbfs_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyDbfs_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyDbfsRequest_SdkV2.
func (m *UpdateDisableLegacyDbfsRequest_SdkV2) SetSetting(ctx context.Context, v DisableLegacyDbfs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateDisableLegacyFeaturesRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateDisableLegacyFeaturesRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateDisableLegacyFeaturesRequest_SdkV2) {
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

func (to *UpdateDisableLegacyFeaturesRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateDisableLegacyFeaturesRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateDisableLegacyFeaturesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateDisableLegacyFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyFeatures_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateDisableLegacyFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateDisableLegacyFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: DisableLegacyFeatures_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateDisableLegacyFeaturesRequest_SdkV2 as
// a DisableLegacyFeatures_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateDisableLegacyFeaturesRequest_SdkV2) GetSetting(ctx context.Context) (DisableLegacyFeatures_SdkV2, bool) {
	var e DisableLegacyFeatures_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyFeatures_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyFeaturesRequest_SdkV2.
func (m *UpdateDisableLegacyFeaturesRequest_SdkV2) SetSetting(ctx context.Context, v DisableLegacyFeatures_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateEnableExportNotebookRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateEnableExportNotebookRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnableExportNotebookRequest_SdkV2) {
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

func (to *UpdateEnableExportNotebookRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEnableExportNotebookRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnableExportNotebookRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnableExportNotebookRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnableExportNotebookRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableExportNotebook_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableExportNotebookRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEnableExportNotebookRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnableExportNotebookRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EnableExportNotebook_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnableExportNotebookRequest_SdkV2 as
// a EnableExportNotebook_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnableExportNotebookRequest_SdkV2) GetSetting(ctx context.Context) (EnableExportNotebook_SdkV2, bool) {
	var e EnableExportNotebook_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []EnableExportNotebook_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEnableExportNotebookRequest_SdkV2.
func (m *UpdateEnableExportNotebookRequest_SdkV2) SetSetting(ctx context.Context, v EnableExportNotebook_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateEnableNotebookTableClipboardRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateEnableNotebookTableClipboardRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnableNotebookTableClipboardRequest_SdkV2) {
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

func (to *UpdateEnableNotebookTableClipboardRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEnableNotebookTableClipboardRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnableNotebookTableClipboardRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnableNotebookTableClipboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnableNotebookTableClipboardRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableNotebookTableClipboard_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableNotebookTableClipboardRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEnableNotebookTableClipboardRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnableNotebookTableClipboardRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EnableNotebookTableClipboard_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnableNotebookTableClipboardRequest_SdkV2 as
// a EnableNotebookTableClipboard_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnableNotebookTableClipboardRequest_SdkV2) GetSetting(ctx context.Context) (EnableNotebookTableClipboard_SdkV2, bool) {
	var e EnableNotebookTableClipboard_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []EnableNotebookTableClipboard_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEnableNotebookTableClipboardRequest_SdkV2.
func (m *UpdateEnableNotebookTableClipboardRequest_SdkV2) SetSetting(ctx context.Context, v EnableNotebookTableClipboard_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateEnableResultsDownloadingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateEnableResultsDownloadingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnableResultsDownloadingRequest_SdkV2) {
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

func (to *UpdateEnableResultsDownloadingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEnableResultsDownloadingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnableResultsDownloadingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnableResultsDownloadingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnableResultsDownloadingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableResultsDownloading_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableResultsDownloadingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEnableResultsDownloadingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnableResultsDownloadingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EnableResultsDownloading_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnableResultsDownloadingRequest_SdkV2 as
// a EnableResultsDownloading_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnableResultsDownloadingRequest_SdkV2) GetSetting(ctx context.Context) (EnableResultsDownloading_SdkV2, bool) {
	var e EnableResultsDownloading_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []EnableResultsDownloading_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEnableResultsDownloadingRequest_SdkV2.
func (m *UpdateEnableResultsDownloadingRequest_SdkV2) SetSetting(ctx context.Context, v EnableResultsDownloading_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) {
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

func (to *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnhancedSecurityMonitoringSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EnhancedSecurityMonitoringSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2 as
// a EnhancedSecurityMonitoringSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) GetSetting(ctx context.Context) (EnhancedSecurityMonitoringSetting_SdkV2, bool) {
	var e EnhancedSecurityMonitoringSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []EnhancedSecurityMonitoringSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2.
func (m *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) SetSetting(ctx context.Context, v EnhancedSecurityMonitoringSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateEsmEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateEsmEnablementAccountSettingRequest_SdkV2) {
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

func (to *UpdateEsmEnablementAccountSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateEsmEnablementAccountSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateEsmEnablementAccountSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateEsmEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EsmEnablementAccountSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEsmEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateEsmEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateEsmEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: EsmEnablementAccountSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateEsmEnablementAccountSettingRequest_SdkV2 as
// a EsmEnablementAccountSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateEsmEnablementAccountSettingRequest_SdkV2) GetSetting(ctx context.Context) (EsmEnablementAccountSetting_SdkV2, bool) {
	var e EsmEnablementAccountSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []EsmEnablementAccountSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEsmEnablementAccountSettingRequest_SdkV2.
func (m *UpdateEsmEnablementAccountSettingRequest_SdkV2) SetSetting(ctx context.Context, v EsmEnablementAccountSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update an IP access list.
type UpdateIpAccessList_SdkV2 struct {
	// Specifies whether this IP access list is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`

	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`

	ListType types.String `tfsdk:"list_type"`
}

func (to *UpdateIpAccessList_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateIpAccessList_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (to *UpdateIpAccessList_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateIpAccessList_SdkV2) {
	if !from.IpAddresses.IsNull() && !from.IpAddresses.IsUnknown() && to.IpAddresses.IsNull() && len(from.IpAddresses.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for IpAddresses, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.IpAddresses = from.IpAddresses
	}
}

func (m UpdateIpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetOptional()
	attrs["list_type"] = attrs["list_type"].SetOptional()
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
func (m UpdateIpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateIpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateIpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateIpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetIpAddresses returns the value of the IpAddresses field in UpdateIpAccessList_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateIpAccessList_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in UpdateIpAccessList_SdkV2.
func (m *UpdateIpAccessList_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.IpAddresses = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) {
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

func (to *UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: LlmProxyPartnerPoweredAccount_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2 as
// a LlmProxyPartnerPoweredAccount_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredAccount_SdkV2, bool) {
	var e LlmProxyPartnerPoweredAccount_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []LlmProxyPartnerPoweredAccount_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2.
func (m *UpdateLlmProxyPartnerPoweredAccountRequest_SdkV2) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) {
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

func (to *UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredEnforceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredEnforce_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: LlmProxyPartnerPoweredEnforce_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2 as
// a LlmProxyPartnerPoweredEnforce_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredEnforce_SdkV2, bool) {
	var e LlmProxyPartnerPoweredEnforce_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []LlmProxyPartnerPoweredEnforce_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2.
func (m *UpdateLlmProxyPartnerPoweredEnforceRequest_SdkV2) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredEnforce_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) {
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

func (to *UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredWorkspace_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: LlmProxyPartnerPoweredWorkspace_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2 as
// a LlmProxyPartnerPoweredWorkspace_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredWorkspace_SdkV2, bool) {
	var e LlmProxyPartnerPoweredWorkspace_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []LlmProxyPartnerPoweredWorkspace_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2.
func (m *UpdateLlmProxyPartnerPoweredWorkspaceRequest_SdkV2) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredWorkspace_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

type UpdateNccPrivateEndpointRuleRequest_SdkV2 struct {
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`

	PrivateEndpointRule types.List `tfsdk:"private_endpoint_rule"`
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

func (to *UpdateNccPrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateNccPrivateEndpointRuleRequest_SdkV2) {
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

func (to *UpdateNccPrivateEndpointRuleRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateNccPrivateEndpointRuleRequest_SdkV2) {
	if !from.PrivateEndpointRule.IsNull() && !from.PrivateEndpointRule.IsUnknown() {
		if toPrivateEndpointRule, ok := to.GetPrivateEndpointRule(ctx); ok {
			if fromPrivateEndpointRule, ok := from.GetPrivateEndpointRule(ctx); ok {
				toPrivateEndpointRule.SyncFieldsDuringRead(ctx, fromPrivateEndpointRule)
				to.SetPrivateEndpointRule(ctx, toPrivateEndpointRule)
			}
		}
	}
}

func (m UpdateNccPrivateEndpointRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_endpoint_rule"] = attrs["private_endpoint_rule"].SetRequired()
	attrs["private_endpoint_rule"] = attrs["private_endpoint_rule"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateNccPrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_endpoint_rule": reflect.TypeOf(UpdatePrivateEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNccPrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateNccPrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m UpdateNccPrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule": basetypes.ListType{
				ElemType: UpdatePrivateEndpointRule_SdkV2{}.Type(ctx),
			},
			"private_endpoint_rule_id": types.StringType,
			"update_mask":              types.StringType,
		},
	}
}

// GetPrivateEndpointRule returns the value of the PrivateEndpointRule field in UpdateNccPrivateEndpointRuleRequest_SdkV2 as
// a UpdatePrivateEndpointRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateNccPrivateEndpointRuleRequest_SdkV2) GetPrivateEndpointRule(ctx context.Context) (UpdatePrivateEndpointRule_SdkV2, bool) {
	var e UpdatePrivateEndpointRule_SdkV2
	if m.PrivateEndpointRule.IsNull() || m.PrivateEndpointRule.IsUnknown() {
		return e, false
	}
	var v []UpdatePrivateEndpointRule_SdkV2
	d := m.PrivateEndpointRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPrivateEndpointRule sets the value of the PrivateEndpointRule field in UpdateNccPrivateEndpointRuleRequest_SdkV2.
func (m *UpdateNccPrivateEndpointRuleRequest_SdkV2) SetPrivateEndpointRule(ctx context.Context, v UpdatePrivateEndpointRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["private_endpoint_rule"]
	m.PrivateEndpointRule = types.ListValueMust(t, vs)
}

type UpdateNetworkPolicyRequest_SdkV2 struct {
	// Updated network policy configuration details.
	NetworkPolicy types.List `tfsdk:"network_policy"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"-"`
}

func (to *UpdateNetworkPolicyRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateNetworkPolicyRequest_SdkV2) {
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

func (to *UpdateNetworkPolicyRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateNetworkPolicyRequest_SdkV2) {
	if !from.NetworkPolicy.IsNull() && !from.NetworkPolicy.IsUnknown() {
		if toNetworkPolicy, ok := to.GetNetworkPolicy(ctx); ok {
			if fromNetworkPolicy, ok := from.GetNetworkPolicy(ctx); ok {
				toNetworkPolicy.SyncFieldsDuringRead(ctx, fromNetworkPolicy)
				to.SetNetworkPolicy(ctx, toNetworkPolicy)
			}
		}
	}
}

func (m UpdateNetworkPolicyRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["network_policy"] = attrs["network_policy"].SetRequired()
	attrs["network_policy"] = attrs["network_policy"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateNetworkPolicyRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_policy": reflect.TypeOf(AccountNetworkPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNetworkPolicyRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateNetworkPolicyRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy":    m.NetworkPolicy,
			"network_policy_id": m.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateNetworkPolicyRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy": basetypes.ListType{
				ElemType: AccountNetworkPolicy_SdkV2{}.Type(ctx),
			},
			"network_policy_id": types.StringType,
		},
	}
}

// GetNetworkPolicy returns the value of the NetworkPolicy field in UpdateNetworkPolicyRequest_SdkV2 as
// a AccountNetworkPolicy_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateNetworkPolicyRequest_SdkV2) GetNetworkPolicy(ctx context.Context) (AccountNetworkPolicy_SdkV2, bool) {
	var e AccountNetworkPolicy_SdkV2
	if m.NetworkPolicy.IsNull() || m.NetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []AccountNetworkPolicy_SdkV2
	d := m.NetworkPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNetworkPolicy sets the value of the NetworkPolicy field in UpdateNetworkPolicyRequest_SdkV2.
func (m *UpdateNetworkPolicyRequest_SdkV2) SetNetworkPolicy(ctx context.Context, v AccountNetworkPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["network_policy"]
	m.NetworkPolicy = types.ListValueMust(t, vs)
}

type UpdateNotificationDestinationRequest_SdkV2 struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.List `tfsdk:"config"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
	// UUID identifying notification destination.
	Id types.String `tfsdk:"-"`
}

func (to *UpdateNotificationDestinationRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateNotificationDestinationRequest_SdkV2) {
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

func (to *UpdateNotificationDestinationRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateNotificationDestinationRequest_SdkV2) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m UpdateNotificationDestinationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       m.Config,
			"display_name": m.DisplayName,
			"id":           m.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: Config_SdkV2{}.Type(ctx),
			},
			"display_name": types.StringType,
			"id":           types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in UpdateNotificationDestinationRequest_SdkV2 as
// a Config_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateNotificationDestinationRequest_SdkV2) GetConfig(ctx context.Context) (Config_SdkV2, bool) {
	var e Config_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []Config_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in UpdateNotificationDestinationRequest_SdkV2.
func (m *UpdateNotificationDestinationRequest_SdkV2) SetConfig(ctx context.Context, v Config_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdatePersonalComputeSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePersonalComputeSettingRequest_SdkV2) {
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

func (to *UpdatePersonalComputeSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdatePersonalComputeSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdatePersonalComputeSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdatePersonalComputeSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(PersonalComputeSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalComputeSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdatePersonalComputeSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePersonalComputeSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: PersonalComputeSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdatePersonalComputeSettingRequest_SdkV2 as
// a PersonalComputeSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePersonalComputeSettingRequest_SdkV2) GetSetting(ctx context.Context) (PersonalComputeSetting_SdkV2, bool) {
	var e PersonalComputeSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdatePersonalComputeSettingRequest_SdkV2.
func (m *UpdatePersonalComputeSettingRequest_SdkV2) SetSetting(ctx context.Context, v PersonalComputeSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type UpdatePrivateEndpointRule_SdkV2 struct {
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

func (to *UpdatePrivateEndpointRule_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdatePrivateEndpointRule_SdkV2) {
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

func (to *UpdatePrivateEndpointRule_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdatePrivateEndpointRule_SdkV2) {
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

func (m UpdatePrivateEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m UpdatePrivateEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePrivateEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdatePrivateEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain_names":   m.DomainNames,
			"enabled":        m.Enabled,
			"resource_names": m.ResourceNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdatePrivateEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetDomainNames returns the value of the DomainNames field in UpdatePrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePrivateEndpointRule_SdkV2) GetDomainNames(ctx context.Context) ([]types.String, bool) {
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

// SetDomainNames sets the value of the DomainNames field in UpdatePrivateEndpointRule_SdkV2.
func (m *UpdatePrivateEndpointRule_SdkV2) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in UpdatePrivateEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdatePrivateEndpointRule_SdkV2) GetResourceNames(ctx context.Context) ([]types.String, bool) {
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

// SetResourceNames sets the value of the ResourceNames field in UpdatePrivateEndpointRule_SdkV2.
func (m *UpdatePrivateEndpointRule_SdkV2) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ResourceNames = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) {
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

func (to *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(RestrictWorkspaceAdminsSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: RestrictWorkspaceAdminsSetting_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2 as
// a RestrictWorkspaceAdminsSetting_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) GetSetting(ctx context.Context) (RestrictWorkspaceAdminsSetting_SdkV2, bool) {
	var e RestrictWorkspaceAdminsSetting_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsSetting_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2.
func (m *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) SetSetting(ctx context.Context, v RestrictWorkspaceAdminsSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

// Details required to update a setting.
type UpdateSqlResultsDownloadRequest_SdkV2 struct {
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

	Setting types.List `tfsdk:"setting"`
}

func (to *UpdateSqlResultsDownloadRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateSqlResultsDownloadRequest_SdkV2) {
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

func (to *UpdateSqlResultsDownloadRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateSqlResultsDownloadRequest_SdkV2) {
	if !from.Setting.IsNull() && !from.Setting.IsUnknown() {
		if toSetting, ok := to.GetSetting(ctx); ok {
			if fromSetting, ok := from.GetSetting(ctx); ok {
				toSetting.SyncFieldsDuringRead(ctx, fromSetting)
				to.SetSetting(ctx, toSetting)
			}
		}
	}
}

func (m UpdateSqlResultsDownloadRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()
	attrs["setting"] = attrs["setting"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateSqlResultsDownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateSqlResultsDownloadRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(SqlResultsDownload_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSqlResultsDownloadRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateSqlResultsDownloadRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": m.AllowMissing,
			"field_mask":    m.FieldMask,
			"setting":       m.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateSqlResultsDownloadRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"allow_missing": types.BoolType,
			"field_mask":    types.StringType,
			"setting": basetypes.ListType{
				ElemType: SqlResultsDownload_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetSetting returns the value of the Setting field in UpdateSqlResultsDownloadRequest_SdkV2 as
// a SqlResultsDownload_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateSqlResultsDownloadRequest_SdkV2) GetSetting(ctx context.Context) (SqlResultsDownload_SdkV2, bool) {
	var e SqlResultsDownload_SdkV2
	if m.Setting.IsNull() || m.Setting.IsUnknown() {
		return e, false
	}
	var v []SqlResultsDownload_SdkV2
	d := m.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateSqlResultsDownloadRequest_SdkV2.
func (m *UpdateSqlResultsDownloadRequest_SdkV2) SetSetting(ctx context.Context, v SqlResultsDownload_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	m.Setting = types.ListValueMust(t, vs)
}

type UpdateWorkspaceNetworkOptionRequest_SdkV2 struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
	// The network option details for the workspace.
	WorkspaceNetworkOption types.List `tfsdk:"workspace_network_option"`
}

func (to *UpdateWorkspaceNetworkOptionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateWorkspaceNetworkOptionRequest_SdkV2) {
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

func (to *UpdateWorkspaceNetworkOptionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateWorkspaceNetworkOptionRequest_SdkV2) {
	if !from.WorkspaceNetworkOption.IsNull() && !from.WorkspaceNetworkOption.IsUnknown() {
		if toWorkspaceNetworkOption, ok := to.GetWorkspaceNetworkOption(ctx); ok {
			if fromWorkspaceNetworkOption, ok := from.GetWorkspaceNetworkOption(ctx); ok {
				toWorkspaceNetworkOption.SyncFieldsDuringRead(ctx, fromWorkspaceNetworkOption)
				to.SetWorkspaceNetworkOption(ctx, toWorkspaceNetworkOption)
			}
		}
	}
}

func (m UpdateWorkspaceNetworkOptionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_network_option"] = attrs["workspace_network_option"].SetRequired()
	attrs["workspace_network_option"] = attrs["workspace_network_option"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (m UpdateWorkspaceNetworkOptionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_network_option": reflect.TypeOf(WorkspaceNetworkOption_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceNetworkOptionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateWorkspaceNetworkOptionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id":             m.WorkspaceId,
			"workspace_network_option": m.WorkspaceNetworkOption,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateWorkspaceNetworkOptionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.Int64Type,
			"workspace_network_option": basetypes.ListType{
				ElemType: WorkspaceNetworkOption_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetWorkspaceNetworkOption returns the value of the WorkspaceNetworkOption field in UpdateWorkspaceNetworkOptionRequest_SdkV2 as
// a WorkspaceNetworkOption_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateWorkspaceNetworkOptionRequest_SdkV2) GetWorkspaceNetworkOption(ctx context.Context) (WorkspaceNetworkOption_SdkV2, bool) {
	var e WorkspaceNetworkOption_SdkV2
	if m.WorkspaceNetworkOption.IsNull() || m.WorkspaceNetworkOption.IsUnknown() {
		return e, false
	}
	var v []WorkspaceNetworkOption_SdkV2
	d := m.WorkspaceNetworkOption.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWorkspaceNetworkOption sets the value of the WorkspaceNetworkOption field in UpdateWorkspaceNetworkOptionRequest_SdkV2.
func (m *UpdateWorkspaceNetworkOptionRequest_SdkV2) SetWorkspaceNetworkOption(ctx context.Context, v WorkspaceNetworkOption_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["workspace_network_option"]
	m.WorkspaceNetworkOption = types.ListValueMust(t, vs)
}

type WorkspaceNetworkOption_SdkV2 struct {
	// The network policy ID to apply to the workspace. This controls the
	// network access rules for all serverless compute resources in the
	// workspace. Each workspace can only be linked to one policy at a time. If
	// no policy is explicitly assigned, the workspace will use
	// 'default-policy'.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"workspace_id"`
}

func (to *WorkspaceNetworkOption_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from WorkspaceNetworkOption_SdkV2) {
}

func (to *WorkspaceNetworkOption_SdkV2) SyncFieldsDuringRead(ctx context.Context, from WorkspaceNetworkOption_SdkV2) {
}

func (m WorkspaceNetworkOption_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m WorkspaceNetworkOption_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetworkOption_SdkV2
// only implements ToObjectValue() and Type().
func (m WorkspaceNetworkOption_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": m.NetworkPolicyId,
			"workspace_id":      m.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m WorkspaceNetworkOption_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
			"workspace_id":      types.Int64Type,
		},
	}
}
