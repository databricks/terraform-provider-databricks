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

func (newState *AccountIpAccessEnable) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountIpAccessEnable) {
}

func (newState *AccountIpAccessEnable) SyncEffectiveFieldsDuringRead(existingState AccountIpAccessEnable) {
}

func (c AccountIpAccessEnable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AccountIpAccessEnable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"acct_ip_acl_enable": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountIpAccessEnable
// only implements ToObjectValue() and Type().
func (o AccountIpAccessEnable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"acct_ip_acl_enable": o.AcctIpAclEnable,
			"etag":               o.Etag,
			"setting_name":       o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountIpAccessEnable) Type(ctx context.Context) attr.Type {
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
func (o *AccountIpAccessEnable) GetAcctIpAclEnable(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.AcctIpAclEnable.IsNull() || o.AcctIpAclEnable.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.AcctIpAclEnable.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAcctIpAclEnable sets the value of the AcctIpAclEnable field in AccountIpAccessEnable.
func (o *AccountIpAccessEnable) SetAcctIpAclEnable(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.AcctIpAclEnable = vs
}

type AccountNetworkPolicy struct {
	// The associated account ID for this Network Policy object.
	AccountId types.String `tfsdk:"account_id"`
	// The network policies applying for egress traffic.
	Egress types.Object `tfsdk:"egress"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"network_policy_id"`
}

func (newState *AccountNetworkPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountNetworkPolicy) {
}

func (newState *AccountNetworkPolicy) SyncEffectiveFieldsDuringRead(existingState AccountNetworkPolicy) {
}

func (c AccountNetworkPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AccountNetworkPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress": reflect.TypeOf(NetworkPolicyEgress{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountNetworkPolicy
// only implements ToObjectValue() and Type().
func (o AccountNetworkPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":        o.AccountId,
			"egress":            o.Egress,
			"network_policy_id": o.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountNetworkPolicy) Type(ctx context.Context) attr.Type {
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
func (o *AccountNetworkPolicy) GetEgress(ctx context.Context) (NetworkPolicyEgress, bool) {
	var e NetworkPolicyEgress
	if o.Egress.IsNull() || o.Egress.IsUnknown() {
		return e, false
	}
	var v []NetworkPolicyEgress
	d := o.Egress.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEgress sets the value of the Egress field in AccountNetworkPolicy.
func (o *AccountNetworkPolicy) SetEgress(ctx context.Context, v NetworkPolicyEgress) {
	vs := v.ToObjectValue(ctx)
	o.Egress = vs
}

type AibiDashboardEmbeddingAccessPolicy struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type"`
}

func (newState *AibiDashboardEmbeddingAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingAccessPolicy) {
}

func (newState *AibiDashboardEmbeddingAccessPolicy) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingAccessPolicy) {
}

func (c AibiDashboardEmbeddingAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicy
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_policy_type": o.AccessPolicyType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingAccessPolicy) Type(ctx context.Context) attr.Type {
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

func (newState *AibiDashboardEmbeddingAccessPolicySetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingAccessPolicySetting) {
}

func (newState *AibiDashboardEmbeddingAccessPolicySetting) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingAccessPolicySetting) {
}

func (c AibiDashboardEmbeddingAccessPolicySetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingAccessPolicySetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicySetting
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingAccessPolicySetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_access_policy": o.AibiDashboardEmbeddingAccessPolicy,
			"etag":                                   o.Etag,
			"setting_name":                           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingAccessPolicySetting) Type(ctx context.Context) attr.Type {
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
func (o *AibiDashboardEmbeddingAccessPolicySetting) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy, bool) {
	var e AibiDashboardEmbeddingAccessPolicy
	if o.AibiDashboardEmbeddingAccessPolicy.IsNull() || o.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicy
	d := o.AibiDashboardEmbeddingAccessPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in AibiDashboardEmbeddingAccessPolicySetting.
func (o *AibiDashboardEmbeddingAccessPolicySetting) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.AibiDashboardEmbeddingAccessPolicy = vs
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains types.List `tfsdk:"approved_domains"`
}

func (newState *AibiDashboardEmbeddingApprovedDomains) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingApprovedDomains) {
}

func (newState *AibiDashboardEmbeddingApprovedDomains) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingApprovedDomains) {
}

func (c AibiDashboardEmbeddingApprovedDomains) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingApprovedDomains) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomains
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingApprovedDomains) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"approved_domains": o.ApprovedDomains,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingApprovedDomains) Type(ctx context.Context) attr.Type {
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
func (o *AibiDashboardEmbeddingApprovedDomains) GetApprovedDomains(ctx context.Context) ([]types.String, bool) {
	if o.ApprovedDomains.IsNull() || o.ApprovedDomains.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApprovedDomains sets the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains.
func (o *AibiDashboardEmbeddingApprovedDomains) SetApprovedDomains(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["approved_domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ApprovedDomains = types.ListValueMust(t, vs)
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

func (newState *AibiDashboardEmbeddingApprovedDomainsSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingApprovedDomainsSetting) {
}

func (newState *AibiDashboardEmbeddingApprovedDomainsSetting) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingApprovedDomainsSetting) {
}

func (c AibiDashboardEmbeddingApprovedDomainsSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingApprovedDomainsSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomainsSetting
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingApprovedDomainsSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_approved_domains": o.AibiDashboardEmbeddingApprovedDomains,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingApprovedDomainsSetting) Type(ctx context.Context) attr.Type {
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
func (o *AibiDashboardEmbeddingApprovedDomainsSetting) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains, bool) {
	var e AibiDashboardEmbeddingApprovedDomains
	if o.AibiDashboardEmbeddingApprovedDomains.IsNull() || o.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomains
	d := o.AibiDashboardEmbeddingApprovedDomains.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in AibiDashboardEmbeddingApprovedDomainsSetting.
func (o *AibiDashboardEmbeddingApprovedDomainsSetting) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains) {
	vs := v.ToObjectValue(ctx)
	o.AibiDashboardEmbeddingApprovedDomains = vs
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

func (newState *AutomaticClusterUpdateSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutomaticClusterUpdateSetting) {
}

func (newState *AutomaticClusterUpdateSetting) SyncEffectiveFieldsDuringRead(existingState AutomaticClusterUpdateSetting) {
}

func (c AutomaticClusterUpdateSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AutomaticClusterUpdateSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"automatic_cluster_update_workspace": reflect.TypeOf(ClusterAutoRestartMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutomaticClusterUpdateSetting
// only implements ToObjectValue() and Type().
func (o AutomaticClusterUpdateSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"automatic_cluster_update_workspace": o.AutomaticClusterUpdateWorkspace,
			"etag":                               o.Etag,
			"setting_name":                       o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutomaticClusterUpdateSetting) Type(ctx context.Context) attr.Type {
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
func (o *AutomaticClusterUpdateSetting) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage, bool) {
	var e ClusterAutoRestartMessage
	if o.AutomaticClusterUpdateWorkspace.IsNull() || o.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessage
	d := o.AutomaticClusterUpdateWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in AutomaticClusterUpdateSetting.
func (o *AutomaticClusterUpdateSetting) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage) {
	vs := v.ToObjectValue(ctx)
	o.AutomaticClusterUpdateWorkspace = vs
}

type BooleanMessage struct {
	Value types.Bool `tfsdk:"value"`
}

func (newState *BooleanMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan BooleanMessage) {
}

func (newState *BooleanMessage) SyncEffectiveFieldsDuringRead(existingState BooleanMessage) {
}

func (c BooleanMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BooleanMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage
// only implements ToObjectValue() and Type().
func (o BooleanMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BooleanMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type ClusterAutoRestartMessage struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails types.Object `tfsdk:"enablement_details"`

	MaintenanceWindow types.Object `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

func (newState *ClusterAutoRestartMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessage) {
}

func (newState *ClusterAutoRestartMessage) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessage) {
}

func (c ClusterAutoRestartMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessage
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"can_toggle":                           o.CanToggle,
			"enabled":                              o.Enabled,
			"enablement_details":                   o.EnablementDetails,
			"maintenance_window":                   o.MaintenanceWindow,
			"restart_even_if_no_updates_available": o.RestartEvenIfNoUpdatesAvailable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessage) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAutoRestartMessage) GetEnablementDetails(ctx context.Context) (ClusterAutoRestartMessageEnablementDetails, bool) {
	var e ClusterAutoRestartMessageEnablementDetails
	if o.EnablementDetails.IsNull() || o.EnablementDetails.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageEnablementDetails
	d := o.EnablementDetails.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEnablementDetails sets the value of the EnablementDetails field in ClusterAutoRestartMessage.
func (o *ClusterAutoRestartMessage) SetEnablementDetails(ctx context.Context, v ClusterAutoRestartMessageEnablementDetails) {
	vs := v.ToObjectValue(ctx)
	o.EnablementDetails = vs
}

// GetMaintenanceWindow returns the value of the MaintenanceWindow field in ClusterAutoRestartMessage as
// a ClusterAutoRestartMessageMaintenanceWindow value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessage) GetMaintenanceWindow(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindow, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindow
	if o.MaintenanceWindow.IsNull() || o.MaintenanceWindow.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindow
	d := o.MaintenanceWindow.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMaintenanceWindow sets the value of the MaintenanceWindow field in ClusterAutoRestartMessage.
func (o *ClusterAutoRestartMessage) SetMaintenanceWindow(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindow) {
	vs := v.ToObjectValue(ctx)
	o.MaintenanceWindow = vs
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

func (newState *ClusterAutoRestartMessageEnablementDetails) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageEnablementDetails) {
}

func (newState *ClusterAutoRestartMessageEnablementDetails) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageEnablementDetails) {
}

func (c ClusterAutoRestartMessageEnablementDetails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageEnablementDetails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageEnablementDetails
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageEnablementDetails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"forced_for_compliance_mode":           o.ForcedForComplianceMode,
			"unavailable_for_disabled_entitlement": o.UnavailableForDisabledEntitlement,
			"unavailable_for_non_enterprise_tier":  o.UnavailableForNonEnterpriseTier,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageEnablementDetails) Type(ctx context.Context) attr.Type {
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

func (newState *ClusterAutoRestartMessageMaintenanceWindow) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindow) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindow) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindow) {
}

func (c ClusterAutoRestartMessageMaintenanceWindow) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindow) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"week_day_based_schedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindow
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindow) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"week_day_based_schedule": o.WeekDayBasedSchedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindow) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"week_day_based_schedule": ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}.Type(ctx),
		},
	}
}

// GetWeekDayBasedSchedule returns the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow as
// a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessageMaintenanceWindow) GetWeekDayBasedSchedule(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
	if o.WeekDayBasedSchedule.IsNull() || o.WeekDayBasedSchedule.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
	d := o.WeekDayBasedSchedule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetWeekDayBasedSchedule sets the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow.
func (o *ClusterAutoRestartMessageMaintenanceWindow) SetWeekDayBasedSchedule(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
	vs := v.ToObjectValue(ctx)
	o.WeekDayBasedSchedule = vs
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`

	Frequency types.String `tfsdk:"frequency"`

	WindowStartTime types.Object `tfsdk:"window_start_time"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) {
}

func (c ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"window_start_time": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"day_of_week":       o.DayOfWeek,
			"frequency":         o.Frequency,
			"window_start_time": o.WindowStartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) GetWindowStartTime(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
	if o.WindowStartTime.IsNull() || o.WindowStartTime.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
	d := o.WindowStartTime.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetWindowStartTime sets the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule.
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) SetWindowStartTime(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
	vs := v.ToObjectValue(ctx)
	o.WindowStartTime = vs
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) {
}

func (c ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hours":   o.Hours,
			"minutes": o.Minutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) Type(ctx context.Context) attr.Type {
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

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfile) {
}

func (newState *ComplianceSecurityProfile) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfile) {
}

func (c ComplianceSecurityProfile) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComplianceSecurityProfile) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile
// only implements ToObjectValue() and Type().
func (o ComplianceSecurityProfile) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": o.ComplianceStandards,
			"is_enabled":           o.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplianceSecurityProfile) Type(ctx context.Context) attr.Type {
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
func (o *ComplianceSecurityProfile) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
	if o.ComplianceStandards.IsNull() || o.ComplianceStandards.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ComplianceStandards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile.
func (o *ComplianceSecurityProfile) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ComplianceStandards = types.ListValueMust(t, vs)
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
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

func (newState *ComplianceSecurityProfileSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfileSetting) {
}

func (newState *ComplianceSecurityProfileSetting) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfileSetting) {
}

func (c ComplianceSecurityProfileSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComplianceSecurityProfileSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_security_profile_workspace": reflect.TypeOf(ComplianceSecurityProfile{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfileSetting
// only implements ToObjectValue() and Type().
func (o ComplianceSecurityProfileSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_security_profile_workspace": o.ComplianceSecurityProfileWorkspace,
			"etag":                                  o.Etag,
			"setting_name":                          o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplianceSecurityProfileSetting) Type(ctx context.Context) attr.Type {
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
func (o *ComplianceSecurityProfileSetting) GetComplianceSecurityProfileWorkspace(ctx context.Context) (ComplianceSecurityProfile, bool) {
	var e ComplianceSecurityProfile
	if o.ComplianceSecurityProfileWorkspace.IsNull() || o.ComplianceSecurityProfileWorkspace.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfile
	d := o.ComplianceSecurityProfileWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetComplianceSecurityProfileWorkspace sets the value of the ComplianceSecurityProfileWorkspace field in ComplianceSecurityProfileSetting.
func (o *ComplianceSecurityProfileSetting) SetComplianceSecurityProfileWorkspace(ctx context.Context, v ComplianceSecurityProfile) {
	vs := v.ToObjectValue(ctx)
	o.ComplianceSecurityProfileWorkspace = vs
}

type Config struct {
	Email types.Object `tfsdk:"email"`

	GenericWebhook types.Object `tfsdk:"generic_webhook"`

	MicrosoftTeams types.Object `tfsdk:"microsoft_teams"`

	Pagerduty types.Object `tfsdk:"pagerduty"`

	Slack types.Object `tfsdk:"slack"`
}

func (newState *Config) SyncEffectiveFieldsDuringCreateOrUpdate(plan Config) {
}

func (newState *Config) SyncEffectiveFieldsDuringRead(existingState Config) {
}

func (c Config) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Config) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o Config) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email":           o.Email,
			"generic_webhook": o.GenericWebhook,
			"microsoft_teams": o.MicrosoftTeams,
			"pagerduty":       o.Pagerduty,
			"slack":           o.Slack,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Config) Type(ctx context.Context) attr.Type {
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
func (o *Config) GetEmail(ctx context.Context) (EmailConfig, bool) {
	var e EmailConfig
	if o.Email.IsNull() || o.Email.IsUnknown() {
		return e, false
	}
	var v []EmailConfig
	d := o.Email.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEmail sets the value of the Email field in Config.
func (o *Config) SetEmail(ctx context.Context, v EmailConfig) {
	vs := v.ToObjectValue(ctx)
	o.Email = vs
}

// GetGenericWebhook returns the value of the GenericWebhook field in Config as
// a GenericWebhookConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config) GetGenericWebhook(ctx context.Context) (GenericWebhookConfig, bool) {
	var e GenericWebhookConfig
	if o.GenericWebhook.IsNull() || o.GenericWebhook.IsUnknown() {
		return e, false
	}
	var v []GenericWebhookConfig
	d := o.GenericWebhook.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetGenericWebhook sets the value of the GenericWebhook field in Config.
func (o *Config) SetGenericWebhook(ctx context.Context, v GenericWebhookConfig) {
	vs := v.ToObjectValue(ctx)
	o.GenericWebhook = vs
}

// GetMicrosoftTeams returns the value of the MicrosoftTeams field in Config as
// a MicrosoftTeamsConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config) GetMicrosoftTeams(ctx context.Context) (MicrosoftTeamsConfig, bool) {
	var e MicrosoftTeamsConfig
	if o.MicrosoftTeams.IsNull() || o.MicrosoftTeams.IsUnknown() {
		return e, false
	}
	var v []MicrosoftTeamsConfig
	d := o.MicrosoftTeams.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetMicrosoftTeams sets the value of the MicrosoftTeams field in Config.
func (o *Config) SetMicrosoftTeams(ctx context.Context, v MicrosoftTeamsConfig) {
	vs := v.ToObjectValue(ctx)
	o.MicrosoftTeams = vs
}

// GetPagerduty returns the value of the Pagerduty field in Config as
// a PagerdutyConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config) GetPagerduty(ctx context.Context) (PagerdutyConfig, bool) {
	var e PagerdutyConfig
	if o.Pagerduty.IsNull() || o.Pagerduty.IsUnknown() {
		return e, false
	}
	var v []PagerdutyConfig
	d := o.Pagerduty.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPagerduty sets the value of the Pagerduty field in Config.
func (o *Config) SetPagerduty(ctx context.Context, v PagerdutyConfig) {
	vs := v.ToObjectValue(ctx)
	o.Pagerduty = vs
}

// GetSlack returns the value of the Slack field in Config as
// a SlackConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config) GetSlack(ctx context.Context) (SlackConfig, bool) {
	var e SlackConfig
	if o.Slack.IsNull() || o.Slack.IsUnknown() {
		return e, false
	}
	var v []SlackConfig
	d := o.Slack.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSlack sets the value of the Slack field in Config.
func (o *Config) SetSlack(ctx context.Context, v SlackConfig) {
	vs := v.ToObjectValue(ctx)
	o.Slack = vs
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses types.List `tfsdk:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label types.String `tfsdk:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type"`
}

func (newState *CreateIpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateIpAccessList) {
}

func (newState *CreateIpAccessList) SyncEffectiveFieldsDuringRead(existingState CreateIpAccessList) {
}

func (c CreateIpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessList
// only implements ToObjectValue() and Type().
func (o CreateIpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_addresses": o.IpAddresses,
			"label":        o.Label,
			"list_type":    o.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateIpAccessList) Type(ctx context.Context) attr.Type {
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
func (o *CreateIpAccessList) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if o.IpAddresses.IsNull() || o.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in CreateIpAccessList.
func (o *CreateIpAccessList) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList types.Object `tfsdk:"ip_access_list"`
}

func (newState *CreateIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateIpAccessListResponse) {
}

func (newState *CreateIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState CreateIpAccessListResponse) {
}

func (c CreateIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessListResponse
// only implements ToObjectValue() and Type().
func (o CreateIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": o.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": IpAccessListInfo{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in CreateIpAccessListResponse as
// a IpAccessListInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateIpAccessListResponse) GetIpAccessList(ctx context.Context) (IpAccessListInfo, bool) {
	var e IpAccessListInfo
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo
	d := o.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIpAccessList sets the value of the IpAccessList field in CreateIpAccessListResponse.
func (o *CreateIpAccessListResponse) SetIpAccessList(ctx context.Context, v IpAccessListInfo) {
	vs := v.ToObjectValue(ctx)
	o.IpAccessList = vs
}

// Create a network connectivity configuration
type CreateNetworkConnectivityConfigRequest struct {
	// Properties of the new network connectivity configuration.
	NetworkConnectivityConfig types.Object `tfsdk:"network_connectivity_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkConnectivityConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkConnectivityConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_connectivity_config": reflect.TypeOf(CreateNetworkConnectivityConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfigRequest
// only implements ToObjectValue() and Type().
func (o CreateNetworkConnectivityConfigRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config": o.NetworkConnectivityConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNetworkConnectivityConfigRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config": CreateNetworkConnectivityConfiguration{}.Type(ctx),
		},
	}
}

// GetNetworkConnectivityConfig returns the value of the NetworkConnectivityConfig field in CreateNetworkConnectivityConfigRequest as
// a CreateNetworkConnectivityConfiguration value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkConnectivityConfigRequest) GetNetworkConnectivityConfig(ctx context.Context) (CreateNetworkConnectivityConfiguration, bool) {
	var e CreateNetworkConnectivityConfiguration
	if o.NetworkConnectivityConfig.IsNull() || o.NetworkConnectivityConfig.IsUnknown() {
		return e, false
	}
	var v []CreateNetworkConnectivityConfiguration
	d := o.NetworkConnectivityConfig.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNetworkConnectivityConfig sets the value of the NetworkConnectivityConfig field in CreateNetworkConnectivityConfigRequest.
func (o *CreateNetworkConnectivityConfigRequest) SetNetworkConnectivityConfig(ctx context.Context, v CreateNetworkConnectivityConfiguration) {
	vs := v.ToObjectValue(ctx)
	o.NetworkConnectivityConfig = vs
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

func (newState *CreateNetworkConnectivityConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNetworkConnectivityConfiguration) {
}

func (newState *CreateNetworkConnectivityConfiguration) SyncEffectiveFieldsDuringRead(existingState CreateNetworkConnectivityConfiguration) {
}

func (c CreateNetworkConnectivityConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateNetworkConnectivityConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfiguration
// only implements ToObjectValue() and Type().
func (o CreateNetworkConnectivityConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   o.Name,
			"region": o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNetworkConnectivityConfiguration) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"region": types.StringType,
		},
	}
}

// Create a network policy
type CreateNetworkPolicyRequest struct {
	NetworkPolicy types.Object `tfsdk:"network_policy"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_policy": reflect.TypeOf(AccountNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (o CreateNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy": o.NetworkPolicy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy": AccountNetworkPolicy{}.Type(ctx),
		},
	}
}

// GetNetworkPolicy returns the value of the NetworkPolicy field in CreateNetworkPolicyRequest as
// a AccountNetworkPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateNetworkPolicyRequest) GetNetworkPolicy(ctx context.Context) (AccountNetworkPolicy, bool) {
	var e AccountNetworkPolicy
	if o.NetworkPolicy.IsNull() || o.NetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []AccountNetworkPolicy
	d := o.NetworkPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNetworkPolicy sets the value of the NetworkPolicy field in CreateNetworkPolicyRequest.
func (o *CreateNetworkPolicyRequest) SetNetworkPolicy(ctx context.Context, v AccountNetworkPolicy) {
	vs := v.ToObjectValue(ctx)
	o.NetworkPolicy = vs
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.Object `tfsdk:"config"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
}

func (newState *CreateNotificationDestinationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNotificationDestinationRequest) {
}

func (newState *CreateNotificationDestinationRequest) SyncEffectiveFieldsDuringRead(existingState CreateNotificationDestinationRequest) {
}

func (c CreateNotificationDestinationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (o CreateNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       o.Config,
			"display_name": o.DisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreateNotificationDestinationRequest) GetConfig(ctx context.Context) (Config, bool) {
	var e Config
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []Config
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetConfig sets the value of the Config field in CreateNotificationDestinationRequest.
func (o *CreateNotificationDestinationRequest) SetConfig(ctx context.Context, v Config) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
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

func (newState *CreateOboTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOboTokenRequest) {
}

func (newState *CreateOboTokenRequest) SyncEffectiveFieldsDuringRead(existingState CreateOboTokenRequest) {
}

func (c CreateOboTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateOboTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenRequest
// only implements ToObjectValue() and Type().
func (o CreateOboTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"application_id":   o.ApplicationId,
			"comment":          o.Comment,
			"lifetime_seconds": o.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOboTokenRequest) Type(ctx context.Context) attr.Type {
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

func (newState *CreateOboTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOboTokenResponse) {
}

func (newState *CreateOboTokenResponse) SyncEffectiveFieldsDuringRead(existingState CreateOboTokenResponse) {
}

func (c CreateOboTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateOboTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenResponse
// only implements ToObjectValue() and Type().
func (o CreateOboTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  o.TokenInfo,
			"token_value": o.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOboTokenResponse) Type(ctx context.Context) attr.Type {
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
func (o *CreateOboTokenResponse) GetTokenInfo(ctx context.Context) (TokenInfo, bool) {
	var e TokenInfo
	if o.TokenInfo.IsNull() || o.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []TokenInfo
	d := o.TokenInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenInfo sets the value of the TokenInfo field in CreateOboTokenResponse.
func (o *CreateOboTokenResponse) SetTokenInfo(ctx context.Context, v TokenInfo) {
	vs := v.ToObjectValue(ctx)
	o.TokenInfo = vs
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

func (newState *CreatePrivateEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePrivateEndpointRule) {
}

func (newState *CreatePrivateEndpointRule) SyncEffectiveFieldsDuringRead(existingState CreatePrivateEndpointRule) {
}

func (c CreatePrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreatePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRule
// only implements ToObjectValue() and Type().
func (o CreatePrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain_names":     o.DomainNames,
			"endpoint_service": o.EndpointService,
			"group_id":         o.GroupId,
			"resource_id":      o.ResourceId,
			"resource_names":   o.ResourceNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePrivateEndpointRule) Type(ctx context.Context) attr.Type {
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
func (o *CreatePrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if o.DomainNames.IsNull() || o.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in CreatePrivateEndpointRule.
func (o *CreatePrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in CreatePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if o.ResourceNames.IsNull() || o.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in CreatePrivateEndpointRule.
func (o *CreatePrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ResourceNames = types.ListValueMust(t, vs)
}

// Create a private endpoint rule
type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	PrivateEndpointRule types.Object `tfsdk:"private_endpoint_rule"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_endpoint_rule": reflect.TypeOf(CreatePrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (o CreatePrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"private_endpoint_rule":          o.PrivateEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
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
func (o *CreatePrivateEndpointRuleRequest) GetPrivateEndpointRule(ctx context.Context) (CreatePrivateEndpointRule, bool) {
	var e CreatePrivateEndpointRule
	if o.PrivateEndpointRule.IsNull() || o.PrivateEndpointRule.IsUnknown() {
		return e, false
	}
	var v []CreatePrivateEndpointRule
	d := o.PrivateEndpointRule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPrivateEndpointRule sets the value of the PrivateEndpointRule field in CreatePrivateEndpointRuleRequest.
func (o *CreatePrivateEndpointRuleRequest) SetPrivateEndpointRule(ctx context.Context, v CreatePrivateEndpointRule) {
	vs := v.ToObjectValue(ctx)
	o.PrivateEndpointRule = vs
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

func (newState *CreateTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTokenRequest) {
}

func (newState *CreateTokenRequest) SyncEffectiveFieldsDuringRead(existingState CreateTokenRequest) {
}

func (c CreateTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenRequest
// only implements ToObjectValue() and Type().
func (o CreateTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":          o.Comment,
			"lifetime_seconds": o.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTokenRequest) Type(ctx context.Context) attr.Type {
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

func (newState *CreateTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTokenResponse) {
}

func (newState *CreateTokenResponse) SyncEffectiveFieldsDuringRead(existingState CreateTokenResponse) {
}

func (c CreateTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(PublicTokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenResponse
// only implements ToObjectValue() and Type().
func (o CreateTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  o.TokenInfo,
			"token_value": o.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTokenResponse) Type(ctx context.Context) attr.Type {
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
func (o *CreateTokenResponse) GetTokenInfo(ctx context.Context) (PublicTokenInfo, bool) {
	var e PublicTokenInfo
	if o.TokenInfo.IsNull() || o.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []PublicTokenInfo
	d := o.TokenInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenInfo sets the value of the TokenInfo field in CreateTokenResponse.
func (o *CreateTokenResponse) SetTokenInfo(ctx context.Context, v PublicTokenInfo) {
	vs := v.ToObjectValue(ctx)
	o.TokenInfo = vs
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (newState *CspEnablementAccount) SyncEffectiveFieldsDuringCreateOrUpdate(plan CspEnablementAccount) {
}

func (newState *CspEnablementAccount) SyncEffectiveFieldsDuringRead(existingState CspEnablementAccount) {
}

func (c CspEnablementAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CspEnablementAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccount
// only implements ToObjectValue() and Type().
func (o CspEnablementAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": o.ComplianceStandards,
			"is_enforced":          o.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CspEnablementAccount) Type(ctx context.Context) attr.Type {
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
func (o *CspEnablementAccount) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
	if o.ComplianceStandards.IsNull() || o.ComplianceStandards.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ComplianceStandards.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetComplianceStandards sets the value of the ComplianceStandards field in CspEnablementAccount.
func (o *CspEnablementAccount) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ComplianceStandards = types.ListValueMust(t, vs)
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
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

func (newState *CspEnablementAccountSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan CspEnablementAccountSetting) {
}

func (newState *CspEnablementAccountSetting) SyncEffectiveFieldsDuringRead(existingState CspEnablementAccountSetting) {
}

func (c CspEnablementAccountSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CspEnablementAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"csp_enablement_account": reflect.TypeOf(CspEnablementAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccountSetting
// only implements ToObjectValue() and Type().
func (o CspEnablementAccountSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"csp_enablement_account": o.CspEnablementAccount,
			"etag":                   o.Etag,
			"setting_name":           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CspEnablementAccountSetting) Type(ctx context.Context) attr.Type {
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
func (o *CspEnablementAccountSetting) GetCspEnablementAccount(ctx context.Context) (CspEnablementAccount, bool) {
	var e CspEnablementAccount
	if o.CspEnablementAccount.IsNull() || o.CspEnablementAccount.IsUnknown() {
		return e, false
	}
	var v []CspEnablementAccount
	d := o.CspEnablementAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetCspEnablementAccount sets the value of the CspEnablementAccount field in CspEnablementAccountSetting.
func (o *CspEnablementAccountSetting) SetCspEnablementAccount(ctx context.Context, v CspEnablementAccount) {
	vs := v.ToObjectValue(ctx)
	o.CspEnablementAccount = vs
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

func (newState *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) {
}

func (newState *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SyncEffectiveFieldsDuringRead(existingState CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) {
}

func (c CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule
// only implements ToObjectValue() and Type().
func (o CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                     o.AccountId,
			"connection_state":               o.ConnectionState,
			"creation_time":                  o.CreationTime,
			"deactivated":                    o.Deactivated,
			"deactivated_at":                 o.DeactivatedAt,
			"domain_names":                   o.DomainNames,
			"enabled":                        o.Enabled,
			"endpoint_service":               o.EndpointService,
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"resource_names":                 o.ResourceNames,
			"rule_id":                        o.RuleId,
			"updated_time":                   o.UpdatedTime,
			"vpc_endpoint_id":                o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) Type(ctx context.Context) attr.Type {
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
func (o *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if o.DomainNames.IsNull() || o.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule.
func (o *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if o.ResourceNames.IsNull() || o.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule.
func (o *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ResourceNames = types.ListValueMust(t, vs)
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

func (newState *DashboardEmailSubscriptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardEmailSubscriptions) {
}

func (newState *DashboardEmailSubscriptions) SyncEffectiveFieldsDuringRead(existingState DashboardEmailSubscriptions) {
}

func (c DashboardEmailSubscriptions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DashboardEmailSubscriptions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DashboardEmailSubscriptions
// only implements ToObjectValue() and Type().
func (o DashboardEmailSubscriptions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DashboardEmailSubscriptions) Type(ctx context.Context) attr.Type {
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
func (o *DashboardEmailSubscriptions) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in DashboardEmailSubscriptions.
func (o *DashboardEmailSubscriptions) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
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

func (newState *DefaultNamespaceSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan DefaultNamespaceSetting) {
}

func (newState *DefaultNamespaceSetting) SyncEffectiveFieldsDuringRead(existingState DefaultNamespaceSetting) {
}

func (c DefaultNamespaceSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DefaultNamespaceSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"namespace": reflect.TypeOf(StringMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultNamespaceSetting
// only implements ToObjectValue() and Type().
func (o DefaultNamespaceSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         o.Etag,
			"namespace":    o.Namespace,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DefaultNamespaceSetting) Type(ctx context.Context) attr.Type {
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
func (o *DefaultNamespaceSetting) GetNamespace(ctx context.Context) (StringMessage, bool) {
	var e StringMessage
	if o.Namespace.IsNull() || o.Namespace.IsUnknown() {
		return e, false
	}
	var v []StringMessage
	d := o.Namespace.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNamespace sets the value of the Namespace field in DefaultNamespaceSetting.
func (o *DefaultNamespaceSetting) SetNamespace(ctx context.Context, v StringMessage) {
	vs := v.ToObjectValue(ctx)
	o.Namespace = vs
}

// Delete the account IP access toggle setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountIpAccessEnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountIpAccessEnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountIpAccessEnableRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteAccountIpAccessEnableResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountIpAccessEnableResponse) {
}

func (newState *DeleteAccountIpAccessEnableResponse) SyncEffectiveFieldsDuringRead(existingState DeleteAccountIpAccessEnableResponse) {
}

func (c DeleteAccountIpAccessEnableResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAccountIpAccessEnableResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableResponse
// only implements ToObjectValue() and Type().
func (o DeleteAccountIpAccessEnableResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountIpAccessEnableResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessListRequest
// only implements ToObjectValue() and Type().
func (o DeleteAccountIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

// Delete the AI/BI dashboard embedding access policy
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingRequest
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) {
}

func (newState *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) {
}

func (c DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingResponse
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete AI/BI dashboard embedding approved domains
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) {
}

func (newState *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) {
}

func (c DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the Dashboard Email Subscriptions setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDashboardEmailSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDashboardEmailSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardEmailSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (o DeleteDashboardEmailSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDashboardEmailSubscriptionsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDashboardEmailSubscriptionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDashboardEmailSubscriptionsResponse) {
}

func (newState *DeleteDashboardEmailSubscriptionsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDashboardEmailSubscriptionsResponse) {
}

func (c DeleteDashboardEmailSubscriptionsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDashboardEmailSubscriptionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDashboardEmailSubscriptionsResponse
// only implements ToObjectValue() and Type().
func (o DeleteDashboardEmailSubscriptionsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDashboardEmailSubscriptionsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the default namespace setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingRequest
// only implements ToObjectValue() and Type().
func (o DeleteDefaultNamespaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDefaultNamespaceSettingRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDefaultNamespaceSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDefaultNamespaceSettingResponse) {
}

func (newState *DeleteDefaultNamespaceSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDefaultNamespaceSettingResponse) {
}

func (c DeleteDefaultNamespaceSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDefaultNamespaceSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingResponse
// only implements ToObjectValue() and Type().
func (o DeleteDefaultNamespaceSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDefaultNamespaceSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete Legacy Access Disablement Status
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessRequest
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyAccessRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDisableLegacyAccessResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyAccessResponse) {
}

func (newState *DeleteDisableLegacyAccessResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyAccessResponse) {
}

func (c DeleteDisableLegacyAccessResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDisableLegacyAccessResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessResponse
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyAccessResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyAccessResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the disable legacy DBFS setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsRequest
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyDbfsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyDbfsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDisableLegacyDbfsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyDbfsResponse) {
}

func (newState *DeleteDisableLegacyDbfsResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyDbfsResponse) {
}

func (c DeleteDisableLegacyDbfsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDisableLegacyDbfsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsResponse
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyDbfsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyDbfsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the disable legacy features setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesRequest
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyFeaturesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDisableLegacyFeaturesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyFeaturesResponse) {
}

func (newState *DeleteDisableLegacyFeaturesResponse) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyFeaturesResponse) {
}

func (c DeleteDisableLegacyFeaturesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDisableLegacyFeaturesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesResponse
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyFeaturesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyFeaturesResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIpAccessListRequest
// only implements ToObjectValue() and Type().
func (o DeleteIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

// Delete the enable partner powered AI features workspace setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteLlmProxyPartnerPoweredWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteLlmProxyPartnerPoweredWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLlmProxyPartnerPoweredWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o DeleteLlmProxyPartnerPoweredWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLlmProxyPartnerPoweredWorkspaceRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteLlmProxyPartnerPoweredWorkspaceResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteLlmProxyPartnerPoweredWorkspaceResponse) {
}

func (newState *DeleteLlmProxyPartnerPoweredWorkspaceResponse) SyncEffectiveFieldsDuringRead(existingState DeleteLlmProxyPartnerPoweredWorkspaceResponse) {
}

func (c DeleteLlmProxyPartnerPoweredWorkspaceResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteLlmProxyPartnerPoweredWorkspaceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteLlmProxyPartnerPoweredWorkspaceResponse
// only implements ToObjectValue() and Type().
func (o DeleteLlmProxyPartnerPoweredWorkspaceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteLlmProxyPartnerPoweredWorkspaceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkConnectivityConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkConnectivityConfigurationRequest
// only implements ToObjectValue() and Type().
func (o DeleteNetworkConnectivityConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkConnectivityConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkConnectivityConfigurationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkConnectivityConfigurationResponse
// only implements ToObjectValue() and Type().
func (o DeleteNetworkConnectivityConfigurationResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkConnectivityConfigurationResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a network policy
type DeleteNetworkPolicyRequest struct {
	// The unique identifier of the network policy to delete.
	NetworkPolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (o DeleteNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": o.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
		},
	}
}

type DeleteNetworkPolicyRpcResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkPolicyRpcResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkPolicyRpcResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkPolicyRpcResponse
// only implements ToObjectValue() and Type().
func (o DeleteNetworkPolicyRpcResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkPolicyRpcResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a notification destination
type DeleteNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (o DeleteNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete Personal Compute setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingRequest
// only implements ToObjectValue() and Type().
func (o DeletePersonalComputeSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePersonalComputeSettingRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeletePersonalComputeSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePersonalComputeSettingResponse) {
}

func (newState *DeletePersonalComputeSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeletePersonalComputeSettingResponse) {
}

func (c DeletePersonalComputeSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeletePersonalComputeSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingResponse
// only implements ToObjectValue() and Type().
func (o DeletePersonalComputeSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePersonalComputeSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (o DeletePrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       o.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
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

// Delete the restrict workspace admins setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingRequest
// only implements ToObjectValue() and Type().
func (o DeleteRestrictWorkspaceAdminsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRestrictWorkspaceAdminsSettingRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteRestrictWorkspaceAdminsSettingResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRestrictWorkspaceAdminsSettingResponse) {
}

func (newState *DeleteRestrictWorkspaceAdminsSettingResponse) SyncEffectiveFieldsDuringRead(existingState DeleteRestrictWorkspaceAdminsSettingResponse) {
}

func (c DeleteRestrictWorkspaceAdminsSettingResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRestrictWorkspaceAdminsSettingResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingResponse
// only implements ToObjectValue() and Type().
func (o DeleteRestrictWorkspaceAdminsSettingResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRestrictWorkspaceAdminsSettingResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the SQL Results Download setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSqlResultsDownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSqlResultsDownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSqlResultsDownloadRequest
// only implements ToObjectValue() and Type().
func (o DeleteSqlResultsDownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSqlResultsDownloadRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteSqlResultsDownloadResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSqlResultsDownloadResponse) {
}

func (newState *DeleteSqlResultsDownloadResponse) SyncEffectiveFieldsDuringRead(existingState DeleteSqlResultsDownloadResponse) {
}

func (c DeleteSqlResultsDownloadResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteSqlResultsDownloadResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteSqlResultsDownloadResponse
// only implements ToObjectValue() and Type().
func (o DeleteSqlResultsDownloadResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteSqlResultsDownloadResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTokenManagementRequest
// only implements ToObjectValue() and Type().
func (o DeleteTokenManagementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTokenManagementRequest) Type(ctx context.Context) attr.Type {
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

func (newState *DisableLegacyAccess) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyAccess) {
}

func (newState *DisableLegacyAccess) SyncEffectiveFieldsDuringRead(existingState DisableLegacyAccess) {
}

func (c DisableLegacyAccess) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DisableLegacyAccess) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_access": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyAccess
// only implements ToObjectValue() and Type().
func (o DisableLegacyAccess) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_access": o.DisableLegacyAccess,
			"etag":                  o.Etag,
			"setting_name":          o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableLegacyAccess) Type(ctx context.Context) attr.Type {
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
func (o *DisableLegacyAccess) GetDisableLegacyAccess(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.DisableLegacyAccess.IsNull() || o.DisableLegacyAccess.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.DisableLegacyAccess.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDisableLegacyAccess sets the value of the DisableLegacyAccess field in DisableLegacyAccess.
func (o *DisableLegacyAccess) SetDisableLegacyAccess(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.DisableLegacyAccess = vs
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

func (newState *DisableLegacyDbfs) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyDbfs) {
}

func (newState *DisableLegacyDbfs) SyncEffectiveFieldsDuringRead(existingState DisableLegacyDbfs) {
}

func (c DisableLegacyDbfs) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DisableLegacyDbfs) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_dbfs": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyDbfs
// only implements ToObjectValue() and Type().
func (o DisableLegacyDbfs) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_dbfs": o.DisableLegacyDbfs,
			"etag":                o.Etag,
			"setting_name":        o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableLegacyDbfs) Type(ctx context.Context) attr.Type {
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
func (o *DisableLegacyDbfs) GetDisableLegacyDbfs(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.DisableLegacyDbfs.IsNull() || o.DisableLegacyDbfs.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.DisableLegacyDbfs.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDisableLegacyDbfs sets the value of the DisableLegacyDbfs field in DisableLegacyDbfs.
func (o *DisableLegacyDbfs) SetDisableLegacyDbfs(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.DisableLegacyDbfs = vs
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

func (newState *DisableLegacyFeatures) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyFeatures) {
}

func (newState *DisableLegacyFeatures) SyncEffectiveFieldsDuringRead(existingState DisableLegacyFeatures) {
}

func (c DisableLegacyFeatures) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DisableLegacyFeatures) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_features": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyFeatures
// only implements ToObjectValue() and Type().
func (o DisableLegacyFeatures) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_features": o.DisableLegacyFeatures,
			"etag":                    o.Etag,
			"setting_name":            o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableLegacyFeatures) Type(ctx context.Context) attr.Type {
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
func (o *DisableLegacyFeatures) GetDisableLegacyFeatures(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.DisableLegacyFeatures.IsNull() || o.DisableLegacyFeatures.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.DisableLegacyFeatures.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDisableLegacyFeatures sets the value of the DisableLegacyFeatures field in DisableLegacyFeatures.
func (o *DisableLegacyFeatures) SetDisableLegacyFeatures(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.DisableLegacyFeatures = vs
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess types.Object `tfsdk:"internet_access"`
}

func (newState *EgressNetworkPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicy) {
}

func (newState *EgressNetworkPolicy) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicy) {
}

func (c EgressNetworkPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internet_access": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicy
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internet_access": o.InternetAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicy) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"internet_access": EgressNetworkPolicyInternetAccessPolicy{}.Type(ctx),
		},
	}
}

// GetInternetAccess returns the value of the InternetAccess field in EgressNetworkPolicy as
// a EgressNetworkPolicyInternetAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicy) GetInternetAccess(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicy, bool) {
	var e EgressNetworkPolicyInternetAccessPolicy
	if o.InternetAccess.IsNull() || o.InternetAccess.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicy
	d := o.InternetAccess.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetInternetAccess sets the value of the InternetAccess field in EgressNetworkPolicy.
func (o *EgressNetworkPolicy) SetInternetAccess(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.InternetAccess = vs
}

type EgressNetworkPolicyInternetAccessPolicy struct {
	AllowedInternetDestinations types.List `tfsdk:"allowed_internet_destinations"`

	AllowedStorageDestinations types.List `tfsdk:"allowed_storage_destinations"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode types.Object `tfsdk:"log_only_mode"`
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	RestrictionMode types.String `tfsdk:"restriction_mode"`
}

func (newState *EgressNetworkPolicyInternetAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicy) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicy) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicy) {
}

func (c EgressNetworkPolicyInternetAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyInternetDestination{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyStorageDestination{}),
		"log_only_mode":                 reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyLogOnlyMode{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicy
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_internet_destinations": o.AllowedInternetDestinations,
			"allowed_storage_destinations":  o.AllowedStorageDestinations,
			"log_only_mode":                 o.LogOnlyMode,
			"restriction_mode":              o.RestrictionMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyInternetAccessPolicy) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyInternetAccessPolicy) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyInternetDestination, bool) {
	if o.AllowedInternetDestinations.IsNull() || o.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyInternetDestination
	d := o.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyInternetAccessPolicy.
func (o *EgressNetworkPolicyInternetAccessPolicy) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyInternetDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy as
// a slice of EgressNetworkPolicyInternetAccessPolicyStorageDestination values.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicyInternetAccessPolicy) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyStorageDestination, bool) {
	if o.AllowedStorageDestinations.IsNull() || o.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyStorageDestination
	d := o.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy.
func (o *EgressNetworkPolicyInternetAccessPolicy) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyStorageDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetLogOnlyMode returns the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy as
// a EgressNetworkPolicyInternetAccessPolicyLogOnlyMode value.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicyInternetAccessPolicy) GetLogOnlyMode(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicyLogOnlyMode, bool) {
	var e EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
	if o.LogOnlyMode.IsNull() || o.LogOnlyMode.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
	d := o.LogOnlyMode.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetLogOnlyMode sets the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy.
func (o *EgressNetworkPolicyInternetAccessPolicy) SetLogOnlyMode(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) {
	vs := v.ToObjectValue(ctx)
	o.LogOnlyMode = vs
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination struct {
	Destination types.String `tfsdk:"destination"`
	// The filtering protocol used by the DP. For private and public preview,
	// SEG will only support TCP filtering (i.e. DNS based filtering, filtering
	// by destination IP address), so protocol will be set to TCP by default and
	// hidden from the user. In the future, users may be able to select HTTP
	// filtering (i.e. SNI based filtering, filtering by FQDN).
	Protocol types.String `tfsdk:"protocol"`

	Type_ types.String `tfsdk:"type"`
}

func (newState *EgressNetworkPolicyInternetAccessPolicyInternetDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicyInternetDestination) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicyInternetDestination) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicyInternetDestination) {
}

func (c EgressNetworkPolicyInternetAccessPolicyInternetDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicyInternetDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyInternetDestination
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicyInternetDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
			"protocol":    o.Protocol,
			"type":        o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyInternetAccessPolicyInternetDestination) Type(ctx context.Context) attr.Type {
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

func (newState *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) {
}

func (c EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workloads": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_only_mode_type": o.LogOnlyModeType,
			"workloads":          o.Workloads,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) GetWorkloads(ctx context.Context) ([]types.String, bool) {
	if o.Workloads.IsNull() || o.Workloads.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Workloads.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetWorkloads sets the value of the Workloads field in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode.
func (o *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) SetWorkloads(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workloads"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Workloads = types.ListValueMust(t, vs)
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

func (newState *EgressNetworkPolicyInternetAccessPolicyStorageDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicyStorageDestination) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicyStorageDestination) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicyStorageDestination) {
}

func (c EgressNetworkPolicyInternetAccessPolicyStorageDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicyStorageDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_paths": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyStorageDestination
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicyStorageDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_paths":         o.AllowedPaths,
			"azure_container":       o.AzureContainer,
			"azure_dns_zone":        o.AzureDnsZone,
			"azure_storage_account": o.AzureStorageAccount,
			"azure_storage_service": o.AzureStorageService,
			"bucket_name":           o.BucketName,
			"region":                o.Region,
			"type":                  o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyInternetAccessPolicyStorageDestination) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyInternetAccessPolicyStorageDestination) GetAllowedPaths(ctx context.Context) ([]types.String, bool) {
	if o.AllowedPaths.IsNull() || o.AllowedPaths.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.AllowedPaths.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedPaths sets the value of the AllowedPaths field in EgressNetworkPolicyInternetAccessPolicyStorageDestination.
func (o *EgressNetworkPolicyInternetAccessPolicyStorageDestination) SetAllowedPaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedPaths = types.ListValueMust(t, vs)
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

func (newState *EgressNetworkPolicyNetworkAccessPolicy) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyNetworkAccessPolicy) {
}

func (newState *EgressNetworkPolicyNetworkAccessPolicy) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyNetworkAccessPolicy) {
}

func (c EgressNetworkPolicyNetworkAccessPolicy) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyNetworkAccessPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyInternetDestination{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyStorageDestination{}),
		"policy_enforcement":            reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicy
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyNetworkAccessPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allowed_internet_destinations": o.AllowedInternetDestinations,
			"allowed_storage_destinations":  o.AllowedStorageDestinations,
			"policy_enforcement":            o.PolicyEnforcement,
			"restriction_mode":              o.RestrictionMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyNetworkAccessPolicy) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyNetworkAccessPolicy) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyNetworkAccessPolicyInternetDestination, bool) {
	if o.AllowedInternetDestinations.IsNull() || o.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyInternetDestination
	d := o.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyNetworkAccessPolicy.
func (o *EgressNetworkPolicyNetworkAccessPolicy) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyNetworkAccessPolicyInternetDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyNetworkAccessPolicy as
// a slice of EgressNetworkPolicyNetworkAccessPolicyStorageDestination values.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicyNetworkAccessPolicy) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyNetworkAccessPolicyStorageDestination, bool) {
	if o.AllowedStorageDestinations.IsNull() || o.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyStorageDestination
	d := o.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyNetworkAccessPolicy.
func (o *EgressNetworkPolicyNetworkAccessPolicy) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyNetworkAccessPolicyStorageDestination) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetPolicyEnforcement returns the value of the PolicyEnforcement field in EgressNetworkPolicyNetworkAccessPolicy as
// a EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement value.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicyNetworkAccessPolicy) GetPolicyEnforcement(ctx context.Context) (EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement, bool) {
	var e EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
	if o.PolicyEnforcement.IsNull() || o.PolicyEnforcement.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
	d := o.PolicyEnforcement.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPolicyEnforcement sets the value of the PolicyEnforcement field in EgressNetworkPolicyNetworkAccessPolicy.
func (o *EgressNetworkPolicyNetworkAccessPolicy) SetPolicyEnforcement(ctx context.Context, v EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) {
	vs := v.ToObjectValue(ctx)
	o.PolicyEnforcement = vs
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

func (newState *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyNetworkAccessPolicyInternetDestination) {
}

func (newState *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyNetworkAccessPolicyInternetDestination) {
}

func (c EgressNetworkPolicyNetworkAccessPolicyInternetDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyNetworkAccessPolicyInternetDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyInternetDestination
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyNetworkAccessPolicyInternetDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination":               o.Destination,
			"internet_destination_type": o.InternetDestinationType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyNetworkAccessPolicyInternetDestination) Type(ctx context.Context) attr.Type {
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

func (newState *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) {
}

func (newState *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) {
}

func (c EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dry_run_mode_product_filter": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dry_run_mode_product_filter": o.DryRunModeProductFilter,
			"enforcement_mode":            o.EnforcementMode,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) GetDryRunModeProductFilter(ctx context.Context) ([]types.String, bool) {
	if o.DryRunModeProductFilter.IsNull() || o.DryRunModeProductFilter.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DryRunModeProductFilter.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDryRunModeProductFilter sets the value of the DryRunModeProductFilter field in EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement.
func (o *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) SetDryRunModeProductFilter(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dry_run_mode_product_filter"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DryRunModeProductFilter = types.ListValueMust(t, vs)
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyNetworkAccessPolicyStorageDestination struct {
	// The Azure storage account name.
	AzureStorageAccount types.String `tfsdk:"azure_storage_account"`
	// The Azure storage service type (blob, dfs, etc.).
	AzureStorageService types.String `tfsdk:"azure_storage_service"`

	BucketName types.String `tfsdk:"bucket_name"`
	// The region of the S3 bucket.
	Region types.String `tfsdk:"region"`
	// The type of storage destination.
	StorageDestinationType types.String `tfsdk:"storage_destination_type"`
}

func (newState *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyNetworkAccessPolicyStorageDestination) {
}

func (newState *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyNetworkAccessPolicyStorageDestination) {
}

func (c EgressNetworkPolicyNetworkAccessPolicyStorageDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyNetworkAccessPolicyStorageDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyNetworkAccessPolicyStorageDestination
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyNetworkAccessPolicyStorageDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_storage_account":    o.AzureStorageAccount,
			"azure_storage_service":    o.AzureStorageService,
			"bucket_name":              o.BucketName,
			"region":                   o.Region,
			"storage_destination_type": o.StorageDestinationType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyNetworkAccessPolicyStorageDestination) Type(ctx context.Context) attr.Type {
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

func (newState *EmailConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmailConfig) {
}

func (newState *EmailConfig) SyncEffectiveFieldsDuringRead(existingState EmailConfig) {
}

func (c EmailConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EmailConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmailConfig
// only implements ToObjectValue() and Type().
func (o EmailConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"addresses": o.Addresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmailConfig) Type(ctx context.Context) attr.Type {
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
func (o *EmailConfig) GetAddresses(ctx context.Context) ([]types.String, bool) {
	if o.Addresses.IsNull() || o.Addresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Addresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAddresses sets the value of the Addresses field in EmailConfig.
func (o *EmailConfig) SetAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Addresses = types.ListValueMust(t, vs)
}

type Empty struct {
}

func (newState *Empty) SyncEffectiveFieldsDuringCreateOrUpdate(plan Empty) {
}

func (newState *Empty) SyncEffectiveFieldsDuringRead(existingState Empty) {
}

func (c Empty) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Empty) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty
// only implements ToObjectValue() and Type().
func (o Empty) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o Empty) Type(ctx context.Context) attr.Type {
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

func (newState *EnableExportNotebook) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableExportNotebook) {
}

func (newState *EnableExportNotebook) SyncEffectiveFieldsDuringRead(existingState EnableExportNotebook) {
}

func (c EnableExportNotebook) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnableExportNotebook) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableExportNotebook
// only implements ToObjectValue() and Type().
func (o EnableExportNotebook) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnableExportNotebook) Type(ctx context.Context) attr.Type {
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
func (o *EnableExportNotebook) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in EnableExportNotebook.
func (o *EnableExportNotebook) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
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

func (newState *EnableNotebookTableClipboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableNotebookTableClipboard) {
}

func (newState *EnableNotebookTableClipboard) SyncEffectiveFieldsDuringRead(existingState EnableNotebookTableClipboard) {
}

func (c EnableNotebookTableClipboard) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnableNotebookTableClipboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableNotebookTableClipboard
// only implements ToObjectValue() and Type().
func (o EnableNotebookTableClipboard) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnableNotebookTableClipboard) Type(ctx context.Context) attr.Type {
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
func (o *EnableNotebookTableClipboard) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in EnableNotebookTableClipboard.
func (o *EnableNotebookTableClipboard) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
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

func (newState *EnableResultsDownloading) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnableResultsDownloading) {
}

func (newState *EnableResultsDownloading) SyncEffectiveFieldsDuringRead(existingState EnableResultsDownloading) {
}

func (c EnableResultsDownloading) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnableResultsDownloading) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnableResultsDownloading
// only implements ToObjectValue() and Type().
func (o EnableResultsDownloading) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnableResultsDownloading) Type(ctx context.Context) attr.Type {
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
func (o *EnableResultsDownloading) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in EnableResultsDownloading.
func (o *EnableResultsDownloading) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (newState *EnhancedSecurityMonitoring) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnhancedSecurityMonitoring) {
}

func (newState *EnhancedSecurityMonitoring) SyncEffectiveFieldsDuringRead(existingState EnhancedSecurityMonitoring) {
}

func (c EnhancedSecurityMonitoring) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnhancedSecurityMonitoring) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoring
// only implements ToObjectValue() and Type().
func (o EnhancedSecurityMonitoring) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enabled": o.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnhancedSecurityMonitoring) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enabled": types.BoolType,
		},
	}
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
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

func (newState *EnhancedSecurityMonitoringSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnhancedSecurityMonitoringSetting) {
}

func (newState *EnhancedSecurityMonitoringSetting) SyncEffectiveFieldsDuringRead(existingState EnhancedSecurityMonitoringSetting) {
}

func (c EnhancedSecurityMonitoringSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnhancedSecurityMonitoringSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enhanced_security_monitoring_workspace": reflect.TypeOf(EnhancedSecurityMonitoring{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoringSetting
// only implements ToObjectValue() and Type().
func (o EnhancedSecurityMonitoringSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enhanced_security_monitoring_workspace": o.EnhancedSecurityMonitoringWorkspace,
			"etag":                                   o.Etag,
			"setting_name":                           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnhancedSecurityMonitoringSetting) Type(ctx context.Context) attr.Type {
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
func (o *EnhancedSecurityMonitoringSetting) GetEnhancedSecurityMonitoringWorkspace(ctx context.Context) (EnhancedSecurityMonitoring, bool) {
	var e EnhancedSecurityMonitoring
	if o.EnhancedSecurityMonitoringWorkspace.IsNull() || o.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		return e, false
	}
	var v []EnhancedSecurityMonitoring
	d := o.EnhancedSecurityMonitoringWorkspace.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEnhancedSecurityMonitoringWorkspace sets the value of the EnhancedSecurityMonitoringWorkspace field in EnhancedSecurityMonitoringSetting.
func (o *EnhancedSecurityMonitoringSetting) SetEnhancedSecurityMonitoringWorkspace(ctx context.Context, v EnhancedSecurityMonitoring) {
	vs := v.ToObjectValue(ctx)
	o.EnhancedSecurityMonitoringWorkspace = vs
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (newState *EsmEnablementAccount) SyncEffectiveFieldsDuringCreateOrUpdate(plan EsmEnablementAccount) {
}

func (newState *EsmEnablementAccount) SyncEffectiveFieldsDuringRead(existingState EsmEnablementAccount) {
}

func (c EsmEnablementAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EsmEnablementAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccount
// only implements ToObjectValue() and Type().
func (o EsmEnablementAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enforced": o.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EsmEnablementAccount) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enforced": types.BoolType,
		},
	}
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
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

func (newState *EsmEnablementAccountSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan EsmEnablementAccountSetting) {
}

func (newState *EsmEnablementAccountSetting) SyncEffectiveFieldsDuringRead(existingState EsmEnablementAccountSetting) {
}

func (c EsmEnablementAccountSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EsmEnablementAccountSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"esm_enablement_account": reflect.TypeOf(EsmEnablementAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccountSetting
// only implements ToObjectValue() and Type().
func (o EsmEnablementAccountSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"esm_enablement_account": o.EsmEnablementAccount,
			"etag":                   o.Etag,
			"setting_name":           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EsmEnablementAccountSetting) Type(ctx context.Context) attr.Type {
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
func (o *EsmEnablementAccountSetting) GetEsmEnablementAccount(ctx context.Context) (EsmEnablementAccount, bool) {
	var e EsmEnablementAccount
	if o.EsmEnablementAccount.IsNull() || o.EsmEnablementAccount.IsUnknown() {
		return e, false
	}
	var v []EsmEnablementAccount
	d := o.EsmEnablementAccount.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEsmEnablementAccount sets the value of the EsmEnablementAccount field in EsmEnablementAccountSetting.
func (o *EsmEnablementAccountSetting) SetEsmEnablementAccount(ctx context.Context, v EsmEnablementAccount) {
	vs := v.ToObjectValue(ctx)
	o.EsmEnablementAccount = vs
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential types.String `tfsdk:"credential"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime types.Int64 `tfsdk:"credentialEolTime"`
	// User ID of the user that owns this token.
	OwnerId types.Int64 `tfsdk:"ownerId"`
	// The scopes of access granted in the token.
	Scopes types.List `tfsdk:"scopes"`
	// The type of this exchange token
	TokenType types.String `tfsdk:"tokenType"`
}

func (newState *ExchangeToken) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeToken) {
}

func (newState *ExchangeToken) SyncEffectiveFieldsDuringRead(existingState ExchangeToken) {
}

func (c ExchangeToken) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["credential"] = attrs["credential"].SetOptional()
	attrs["credentialEolTime"] = attrs["credentialEolTime"].SetOptional()
	attrs["ownerId"] = attrs["ownerId"].SetOptional()
	attrs["scopes"] = attrs["scopes"].SetOptional()
	attrs["tokenType"] = attrs["tokenType"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeToken.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeToken) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeToken
// only implements ToObjectValue() and Type().
func (o ExchangeToken) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"credential":        o.Credential,
			"credentialEolTime": o.CredentialEolTime,
			"ownerId":           o.OwnerId,
			"scopes":            o.Scopes,
			"tokenType":         o.TokenType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeToken) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"credential":        types.StringType,
			"credentialEolTime": types.Int64Type,
			"ownerId":           types.Int64Type,
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"tokenType": types.StringType,
		},
	}
}

// GetScopes returns the value of the Scopes field in ExchangeToken as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeToken) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeToken.
func (o *ExchangeToken) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId types.Object `tfsdk:"partitionId"`
	// Array of scopes for the token request.
	Scopes types.List `tfsdk:"scopes"`
	// A list of token types being requested
	TokenType types.List `tfsdk:"tokenType"`
}

func (newState *ExchangeTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeTokenRequest) {
}

func (newState *ExchangeTokenRequest) SyncEffectiveFieldsDuringRead(existingState ExchangeTokenRequest) {
}

func (c ExchangeTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["partitionId"] = attrs["partitionId"].SetRequired()
	attrs["scopes"] = attrs["scopes"].SetRequired()
	attrs["tokenType"] = attrs["tokenType"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExchangeTokenRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExchangeTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitionId": reflect.TypeOf(PartitionId{}),
		"scopes":      reflect.TypeOf(types.String{}),
		"tokenType":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenRequest
// only implements ToObjectValue() and Type().
func (o ExchangeTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"partitionId": o.PartitionId,
			"scopes":      o.Scopes,
			"tokenType":   o.TokenType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"partitionId": PartitionId{}.Type(ctx),
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"tokenType": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPartitionId returns the value of the PartitionId field in ExchangeTokenRequest as
// a PartitionId value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeTokenRequest) GetPartitionId(ctx context.Context) (PartitionId, bool) {
	var e PartitionId
	if o.PartitionId.IsNull() || o.PartitionId.IsUnknown() {
		return e, false
	}
	var v []PartitionId
	d := o.PartitionId.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPartitionId sets the value of the PartitionId field in ExchangeTokenRequest.
func (o *ExchangeTokenRequest) SetPartitionId(ctx context.Context, v PartitionId) {
	vs := v.ToObjectValue(ctx)
	o.PartitionId = vs
}

// GetScopes returns the value of the Scopes field in ExchangeTokenRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeTokenRequest) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeTokenRequest.
func (o *ExchangeTokenRequest) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

// GetTokenType returns the value of the TokenType field in ExchangeTokenRequest as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeTokenRequest) GetTokenType(ctx context.Context) ([]types.String, bool) {
	if o.TokenType.IsNull() || o.TokenType.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TokenType.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenType sets the value of the TokenType field in ExchangeTokenRequest.
func (o *ExchangeTokenRequest) SetTokenType(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tokenType"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TokenType = types.ListValueMust(t, vs)
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values types.List `tfsdk:"values"`
}

func (newState *ExchangeTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeTokenResponse) {
}

func (newState *ExchangeTokenResponse) SyncEffectiveFieldsDuringRead(existingState ExchangeTokenResponse) {
}

func (c ExchangeTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExchangeTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(ExchangeToken{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenResponse
// only implements ToObjectValue() and Type().
func (o ExchangeTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeTokenResponse) Type(ctx context.Context) attr.Type {
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
func (o *ExchangeTokenResponse) GetValues(ctx context.Context) ([]ExchangeToken, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []ExchangeToken
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in ExchangeTokenResponse.
func (o *ExchangeTokenResponse) SetValues(ctx context.Context, v []ExchangeToken) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList types.Object `tfsdk:"ip_access_list"`
}

func (newState *FetchIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan FetchIpAccessListResponse) {
}

func (newState *FetchIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState FetchIpAccessListResponse) {
}

func (c FetchIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FetchIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FetchIpAccessListResponse
// only implements ToObjectValue() and Type().
func (o FetchIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": o.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FetchIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": IpAccessListInfo{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in FetchIpAccessListResponse as
// a IpAccessListInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *FetchIpAccessListResponse) GetIpAccessList(ctx context.Context) (IpAccessListInfo, bool) {
	var e IpAccessListInfo
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo
	d := o.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIpAccessList sets the value of the IpAccessList field in FetchIpAccessListResponse.
func (o *FetchIpAccessListResponse) SetIpAccessList(ctx context.Context, v IpAccessListInfo) {
	vs := v.ToObjectValue(ctx)
	o.IpAccessList = vs
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

func (newState *GenericWebhookConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenericWebhookConfig) {
}

func (newState *GenericWebhookConfig) SyncEffectiveFieldsDuringRead(existingState GenericWebhookConfig) {
}

func (c GenericWebhookConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenericWebhookConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenericWebhookConfig
// only implements ToObjectValue() and Type().
func (o GenericWebhookConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"password":     o.Password,
			"password_set": o.PasswordSet,
			"url":          o.Url,
			"url_set":      o.UrlSet,
			"username":     o.Username,
			"username_set": o.UsernameSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GenericWebhookConfig) Type(ctx context.Context) attr.Type {
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

// Get the account IP access toggle setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountIpAccessEnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessEnableRequest
// only implements ToObjectValue() and Type().
func (o GetAccountIpAccessEnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountIpAccessEnableRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessListRequest
// only implements ToObjectValue() and Type().
func (o GetAccountIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

// Retrieve the AI/BI dashboard embedding access policy
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingAccessPolicySettingRequest
// only implements ToObjectValue() and Type().
func (o GetAibiDashboardEmbeddingAccessPolicySettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAibiDashboardEmbeddingAccessPolicySettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Retrieve the list of domains approved to host embedded AI/BI dashboards
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingApprovedDomainsSettingRequest
// only implements ToObjectValue() and Type().
func (o GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the automatic cluster update setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAutomaticClusterUpdateSettingRequest
// only implements ToObjectValue() and Type().
func (o GetAutomaticClusterUpdateSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAutomaticClusterUpdateSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the compliance security profile setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetComplianceSecurityProfileSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetComplianceSecurityProfileSettingRequest
// only implements ToObjectValue() and Type().
func (o GetComplianceSecurityProfileSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetComplianceSecurityProfileSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the compliance security profile setting for new workspaces
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCspEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCspEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o GetCspEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCspEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the Dashboard Email Subscriptions setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDashboardEmailSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDashboardEmailSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDashboardEmailSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (o GetDashboardEmailSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDashboardEmailSubscriptionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the default namespace setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultNamespaceSettingRequest
// only implements ToObjectValue() and Type().
func (o GetDefaultNamespaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDefaultNamespaceSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Retrieve Legacy Access Disablement Status
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyAccessRequest
// only implements ToObjectValue() and Type().
func (o GetDisableLegacyAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDisableLegacyAccessRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the disable legacy DBFS setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyDbfsRequest
// only implements ToObjectValue() and Type().
func (o GetDisableLegacyDbfsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDisableLegacyDbfsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the disable legacy features setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyFeaturesRequest
// only implements ToObjectValue() and Type().
func (o GetDisableLegacyFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDisableLegacyFeaturesRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the enhanced security monitoring setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnhancedSecurityMonitoringSettingRequest
// only implements ToObjectValue() and Type().
func (o GetEnhancedSecurityMonitoringSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEnhancedSecurityMonitoringSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the enhanced security monitoring setting for new workspaces
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEsmEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEsmEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o GetEsmEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEsmEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetIpAccessListRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetIpAccessListRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListRequest
// only implements ToObjectValue() and Type().
func (o GetIpAccessListRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIpAccessListRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList types.Object `tfsdk:"ip_access_list"`
}

func (newState *GetIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListResponse) {
}

func (newState *GetIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListResponse) {
}

func (c GetIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListResponse
// only implements ToObjectValue() and Type().
func (o GetIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": o.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIpAccessListResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list": IpAccessListInfo{}.Type(ctx),
		},
	}
}

// GetIpAccessList returns the value of the IpAccessList field in GetIpAccessListResponse as
// a IpAccessListInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetIpAccessListResponse) GetIpAccessList(ctx context.Context) (IpAccessListInfo, bool) {
	var e IpAccessListInfo
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo
	d := o.IpAccessList.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetIpAccessList sets the value of the IpAccessList field in GetIpAccessListResponse.
func (o *GetIpAccessListResponse) SetIpAccessList(ctx context.Context, v IpAccessListInfo) {
	vs := v.ToObjectValue(ctx)
	o.IpAccessList = vs
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (newState *GetIpAccessListsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListsResponse) {
}

func (newState *GetIpAccessListsResponse) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListsResponse) {
}

func (c GetIpAccessListsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetIpAccessListsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListsResponse
// only implements ToObjectValue() and Type().
func (o GetIpAccessListsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": o.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIpAccessListsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetIpAccessListsResponse) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo, bool) {
	if o.IpAccessLists.IsNull() || o.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo
	d := o.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in GetIpAccessListsResponse.
func (o *GetIpAccessListsResponse) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAccessLists = types.ListValueMust(t, vs)
}

// Get the enable partner powered AI features account setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLlmProxyPartnerPoweredAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLlmProxyPartnerPoweredAccountRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredAccountRequest
// only implements ToObjectValue() and Type().
func (o GetLlmProxyPartnerPoweredAccountRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLlmProxyPartnerPoweredAccountRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the enforcement status of partner powered AI features account setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLlmProxyPartnerPoweredEnforceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLlmProxyPartnerPoweredEnforceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredEnforceRequest
// only implements ToObjectValue() and Type().
func (o GetLlmProxyPartnerPoweredEnforceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLlmProxyPartnerPoweredEnforceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the enable partner powered AI features workspace setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetLlmProxyPartnerPoweredWorkspaceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetLlmProxyPartnerPoweredWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetLlmProxyPartnerPoweredWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o GetLlmProxyPartnerPoweredWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetLlmProxyPartnerPoweredWorkspaceRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkConnectivityConfigurationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkConnectivityConfigurationRequest
// only implements ToObjectValue() and Type().
func (o GetNetworkConnectivityConfigurationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkConnectivityConfigurationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

// Get a network policy
type GetNetworkPolicyRequest struct {
	// The unique identifier of the network policy to retrieve.
	NetworkPolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (o GetNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": o.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
		},
	}
}

// Get a notification destination
type GetNotificationDestinationRequest struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (o GetNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get Personal Compute setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalComputeSettingRequest
// only implements ToObjectValue() and Type().
func (o GetPersonalComputeSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPersonalComputeSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Gets a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (o GetPrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       o.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

// Get the restrict workspace admins setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRestrictWorkspaceAdminsSettingRequest
// only implements ToObjectValue() and Type().
func (o GetRestrictWorkspaceAdminsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRestrictWorkspaceAdminsSettingRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the SQL Results Download setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSqlResultsDownloadRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSqlResultsDownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetSqlResultsDownloadRequest
// only implements ToObjectValue() and Type().
func (o GetSqlResultsDownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetSqlResultsDownloadRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Check configuration status
type GetStatusRequest struct {
	Keys types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatusRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest
// only implements ToObjectValue() and Type().
func (o GetStatusRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"keys": o.Keys,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatusRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"keys": types.StringType,
		},
	}
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenManagementRequest
// only implements ToObjectValue() and Type().
func (o GetTokenManagementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTokenManagementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetTokenPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenPermissionLevelsResponse) {
}

func (newState *GetTokenPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetTokenPermissionLevelsResponse) {
}

func (c GetTokenPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetTokenPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(TokenPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetTokenPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTokenPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
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
func (o *GetTokenPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]TokenPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []TokenPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetTokenPermissionLevelsResponse.
func (o *GetTokenPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []TokenPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo types.Object `tfsdk:"token_info"`
}

func (newState *GetTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenResponse) {
}

func (newState *GetTokenResponse) SyncEffectiveFieldsDuringRead(existingState GetTokenResponse) {
}

func (c GetTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenResponse
// only implements ToObjectValue() and Type().
func (o GetTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info": o.TokenInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_info": TokenInfo{}.Type(ctx),
		},
	}
}

// GetTokenInfo returns the value of the TokenInfo field in GetTokenResponse as
// a TokenInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *GetTokenResponse) GetTokenInfo(ctx context.Context) (TokenInfo, bool) {
	var e TokenInfo
	if o.TokenInfo.IsNull() || o.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []TokenInfo
	d := o.TokenInfo.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTokenInfo sets the value of the TokenInfo field in GetTokenResponse.
func (o *GetTokenResponse) SetTokenInfo(ctx context.Context, v TokenInfo) {
	vs := v.ToObjectValue(ctx)
	o.TokenInfo = vs
}

// Get workspace network option
type GetWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetWorkspaceNetworkOptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetWorkspaceNetworkOptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetWorkspaceNetworkOptionRequest
// only implements ToObjectValue() and Type().
func (o GetWorkspaceNetworkOptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetWorkspaceNetworkOptionRequest) Type(ctx context.Context) attr.Type {
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
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type"`
	// Update timestamp in milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at"`
	// User ID of the user who updated this list.
	UpdatedBy types.Int64 `tfsdk:"updated_by"`
}

func (newState *IpAccessListInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessListInfo) {
}

func (newState *IpAccessListInfo) SyncEffectiveFieldsDuringRead(existingState IpAccessListInfo) {
}

func (c IpAccessListInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a IpAccessListInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessListInfo
// only implements ToObjectValue() and Type().
func (o IpAccessListInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"address_count": o.AddressCount,
			"created_at":    o.CreatedAt,
			"created_by":    o.CreatedBy,
			"enabled":       o.Enabled,
			"ip_addresses":  o.IpAddresses,
			"label":         o.Label,
			"list_id":       o.ListId,
			"list_type":     o.ListType,
			"updated_at":    o.UpdatedAt,
			"updated_by":    o.UpdatedBy,
		})
}

// Type implements basetypes.ObjectValuable.
func (o IpAccessListInfo) Type(ctx context.Context) attr.Type {
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
func (o *IpAccessListInfo) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if o.IpAddresses.IsNull() || o.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in IpAccessListInfo.
func (o *IpAccessListInfo) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (newState *ListIpAccessListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListIpAccessListResponse) {
}

func (newState *ListIpAccessListResponse) SyncEffectiveFieldsDuringRead(existingState ListIpAccessListResponse) {
}

func (c ListIpAccessListResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListIpAccessListResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIpAccessListResponse
// only implements ToObjectValue() and Type().
func (o ListIpAccessListResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": o.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListIpAccessListResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListIpAccessListResponse) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo, bool) {
	if o.IpAccessLists.IsNull() || o.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo
	d := o.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in ListIpAccessListResponse.
func (o *ListIpAccessListResponse) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAccessLists = types.ListValueMust(t, vs)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkConnectivityConfigurationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNetworkConnectivityConfigurationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsRequest
// only implements ToObjectValue() and Type().
func (o ListNetworkConnectivityConfigurationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworkConnectivityConfigurationsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListNetworkConnectivityConfigurationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNetworkConnectivityConfigurationsResponse) {
}

func (newState *ListNetworkConnectivityConfigurationsResponse) SyncEffectiveFieldsDuringRead(existingState ListNetworkConnectivityConfigurationsResponse) {
}

func (c ListNetworkConnectivityConfigurationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNetworkConnectivityConfigurationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NetworkConnectivityConfiguration{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsResponse
// only implements ToObjectValue() and Type().
func (o ListNetworkConnectivityConfigurationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           o.Items,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworkConnectivityConfigurationsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListNetworkConnectivityConfigurationsResponse) GetItems(ctx context.Context) ([]NetworkConnectivityConfiguration, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []NetworkConnectivityConfiguration
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkConnectivityConfigurationsResponse.
func (o *ListNetworkConnectivityConfigurationsResponse) SetItems(ctx context.Context, v []NetworkConnectivityConfiguration) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
}

// List network policies
type ListNetworkPoliciesRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNetworkPoliciesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNetworkPoliciesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkPoliciesRequest
// only implements ToObjectValue() and Type().
func (o ListNetworkPoliciesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworkPoliciesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListNetworkPoliciesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNetworkPoliciesResponse) {
}

func (newState *ListNetworkPoliciesResponse) SyncEffectiveFieldsDuringRead(existingState ListNetworkPoliciesResponse) {
}

func (c ListNetworkPoliciesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNetworkPoliciesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(AccountNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkPoliciesResponse
// only implements ToObjectValue() and Type().
func (o ListNetworkPoliciesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           o.Items,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworkPoliciesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListNetworkPoliciesResponse) GetItems(ctx context.Context) ([]AccountNetworkPolicy, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []AccountNetworkPolicy
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkPoliciesResponse.
func (o *ListNetworkPoliciesResponse) SetItems(ctx context.Context, v []AccountNetworkPolicy) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
}

// List notification destinations
type ListNotificationDestinationsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNotificationDestinationsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNotificationDestinationsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsRequest
// only implements ToObjectValue() and Type().
func (o ListNotificationDestinationsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNotificationDestinationsRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListNotificationDestinationsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsResponse) {
}

func (newState *ListNotificationDestinationsResponse) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsResponse) {
}

func (c ListNotificationDestinationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNotificationDestinationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListNotificationDestinationsResult{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResponse
// only implements ToObjectValue() and Type().
func (o ListNotificationDestinationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNotificationDestinationsResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListNotificationDestinationsResponse) GetResults(ctx context.Context) ([]ListNotificationDestinationsResult, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ListNotificationDestinationsResult
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListNotificationDestinationsResponse.
func (o *ListNotificationDestinationsResponse) SetResults(ctx context.Context, v []ListNotificationDestinationsResult) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
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

func (newState *ListNotificationDestinationsResult) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsResult) {
}

func (newState *ListNotificationDestinationsResult) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsResult) {
}

func (c ListNotificationDestinationsResult) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNotificationDestinationsResult) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResult
// only implements ToObjectValue() and Type().
func (o ListNotificationDestinationsResult) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_type": o.DestinationType,
			"display_name":     o.DisplayName,
			"id":               o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNotificationDestinationsResult) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
		},
	}
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListPrivateEndpointRulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListPrivateEndpointRulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesRequest
// only implements ToObjectValue() and Type().
func (o ListPrivateEndpointRulesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"page_token":                     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPrivateEndpointRulesRequest) Type(ctx context.Context) attr.Type {
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

func (newState *ListPrivateEndpointRulesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPrivateEndpointRulesResponse) {
}

func (newState *ListPrivateEndpointRulesResponse) SyncEffectiveFieldsDuringRead(existingState ListPrivateEndpointRulesResponse) {
}

func (c ListPrivateEndpointRulesResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPrivateEndpointRulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NccPrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesResponse
// only implements ToObjectValue() and Type().
func (o ListPrivateEndpointRulesResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           o.Items,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPrivateEndpointRulesResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListPrivateEndpointRulesResponse) GetItems(ctx context.Context) ([]NccPrivateEndpointRule, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []NccPrivateEndpointRule
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListPrivateEndpointRulesResponse.
func (o *ListPrivateEndpointRulesResponse) SetItems(ctx context.Context, v []NccPrivateEndpointRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (newState *ListPublicTokensResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPublicTokensResponse) {
}

func (newState *ListPublicTokensResponse) SyncEffectiveFieldsDuringRead(existingState ListPublicTokensResponse) {
}

func (c ListPublicTokensResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPublicTokensResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(PublicTokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublicTokensResponse
// only implements ToObjectValue() and Type().
func (o ListPublicTokensResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": o.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPublicTokensResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListPublicTokensResponse) GetTokenInfos(ctx context.Context) ([]PublicTokenInfo, bool) {
	if o.TokenInfos.IsNull() || o.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []PublicTokenInfo
	d := o.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListPublicTokensResponse.
func (o *ListPublicTokensResponse) SetTokenInfos(ctx context.Context, v []PublicTokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TokenInfos = types.ListValueMust(t, vs)
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById types.Int64 `tfsdk:"-"`
	// Username of the user that created the token.
	CreatedByUsername types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListTokenManagementRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListTokenManagementRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokenManagementRequest
// only implements ToObjectValue() and Type().
func (o ListTokenManagementRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_by_id":       o.CreatedById,
			"created_by_username": o.CreatedByUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTokenManagementRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
		},
	}
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (newState *ListTokensResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTokensResponse) {
}

func (newState *ListTokensResponse) SyncEffectiveFieldsDuringRead(existingState ListTokensResponse) {
}

func (c ListTokensResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTokensResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(TokenInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokensResponse
// only implements ToObjectValue() and Type().
func (o ListTokensResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": o.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTokensResponse) Type(ctx context.Context) attr.Type {
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
func (o *ListTokensResponse) GetTokenInfos(ctx context.Context) ([]TokenInfo, bool) {
	if o.TokenInfos.IsNull() || o.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo
	d := o.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListTokensResponse.
func (o *ListTokensResponse) SetTokenInfos(ctx context.Context, v []TokenInfo) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TokenInfos = types.ListValueMust(t, vs)
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

func (newState *LlmProxyPartnerPoweredAccount) SyncEffectiveFieldsDuringCreateOrUpdate(plan LlmProxyPartnerPoweredAccount) {
}

func (newState *LlmProxyPartnerPoweredAccount) SyncEffectiveFieldsDuringRead(existingState LlmProxyPartnerPoweredAccount) {
}

func (c LlmProxyPartnerPoweredAccount) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LlmProxyPartnerPoweredAccount) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredAccount
// only implements ToObjectValue() and Type().
func (o LlmProxyPartnerPoweredAccount) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LlmProxyPartnerPoweredAccount) Type(ctx context.Context) attr.Type {
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
func (o *LlmProxyPartnerPoweredAccount) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredAccount.
func (o *LlmProxyPartnerPoweredAccount) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
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

func (newState *LlmProxyPartnerPoweredEnforce) SyncEffectiveFieldsDuringCreateOrUpdate(plan LlmProxyPartnerPoweredEnforce) {
}

func (newState *LlmProxyPartnerPoweredEnforce) SyncEffectiveFieldsDuringRead(existingState LlmProxyPartnerPoweredEnforce) {
}

func (c LlmProxyPartnerPoweredEnforce) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LlmProxyPartnerPoweredEnforce) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredEnforce
// only implements ToObjectValue() and Type().
func (o LlmProxyPartnerPoweredEnforce) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LlmProxyPartnerPoweredEnforce) Type(ctx context.Context) attr.Type {
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
func (o *LlmProxyPartnerPoweredEnforce) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredEnforce.
func (o *LlmProxyPartnerPoweredEnforce) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
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

func (newState *LlmProxyPartnerPoweredWorkspace) SyncEffectiveFieldsDuringCreateOrUpdate(plan LlmProxyPartnerPoweredWorkspace) {
}

func (newState *LlmProxyPartnerPoweredWorkspace) SyncEffectiveFieldsDuringRead(existingState LlmProxyPartnerPoweredWorkspace) {
}

func (c LlmProxyPartnerPoweredWorkspace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LlmProxyPartnerPoweredWorkspace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LlmProxyPartnerPoweredWorkspace
// only implements ToObjectValue() and Type().
func (o LlmProxyPartnerPoweredWorkspace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LlmProxyPartnerPoweredWorkspace) Type(ctx context.Context) attr.Type {
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
func (o *LlmProxyPartnerPoweredWorkspace) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in LlmProxyPartnerPoweredWorkspace.
func (o *LlmProxyPartnerPoweredWorkspace) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
}

func (newState *MicrosoftTeamsConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan MicrosoftTeamsConfig) {
}

func (newState *MicrosoftTeamsConfig) SyncEffectiveFieldsDuringRead(existingState MicrosoftTeamsConfig) {
}

func (c MicrosoftTeamsConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MicrosoftTeamsConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MicrosoftTeamsConfig
// only implements ToObjectValue() and Type().
func (o MicrosoftTeamsConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"url":     o.Url,
			"url_set": o.UrlSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MicrosoftTeamsConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"url":     types.StringType,
			"url_set": types.BoolType,
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

func (newState *NccAwsStableIpRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAwsStableIpRule) {
}

func (newState *NccAwsStableIpRule) SyncEffectiveFieldsDuringRead(existingState NccAwsStableIpRule) {
}

func (c NccAwsStableIpRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccAwsStableIpRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cidr_blocks": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAwsStableIpRule
// only implements ToObjectValue() and Type().
func (o NccAwsStableIpRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cidr_blocks": o.CidrBlocks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccAwsStableIpRule) Type(ctx context.Context) attr.Type {
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
func (o *NccAwsStableIpRule) GetCidrBlocks(ctx context.Context) ([]types.String, bool) {
	if o.CidrBlocks.IsNull() || o.CidrBlocks.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.CidrBlocks.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCidrBlocks sets the value of the CidrBlocks field in NccAwsStableIpRule.
func (o *NccAwsStableIpRule) SetCidrBlocks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cidr_blocks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CidrBlocks = types.ListValueMust(t, vs)
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

func (newState *NccAzurePrivateEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAzurePrivateEndpointRule) {
}

func (newState *NccAzurePrivateEndpointRule) SyncEffectiveFieldsDuringRead(existingState NccAzurePrivateEndpointRule) {
}

func (c NccAzurePrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccAzurePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzurePrivateEndpointRule
// only implements ToObjectValue() and Type().
func (o NccAzurePrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_state":               o.ConnectionState,
			"creation_time":                  o.CreationTime,
			"deactivated":                    o.Deactivated,
			"deactivated_at":                 o.DeactivatedAt,
			"domain_names":                   o.DomainNames,
			"endpoint_name":                  o.EndpointName,
			"group_id":                       o.GroupId,
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"resource_id":                    o.ResourceId,
			"rule_id":                        o.RuleId,
			"updated_time":                   o.UpdatedTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccAzurePrivateEndpointRule) Type(ctx context.Context) attr.Type {
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
func (o *NccAzurePrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if o.DomainNames.IsNull() || o.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in NccAzurePrivateEndpointRule.
func (o *NccAzurePrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DomainNames = types.ListValueMust(t, vs)
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

func (newState *NccAzureServiceEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAzureServiceEndpointRule) {
}

func (newState *NccAzureServiceEndpointRule) SyncEffectiveFieldsDuringRead(existingState NccAzureServiceEndpointRule) {
}

func (c NccAzureServiceEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccAzureServiceEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subnets":         reflect.TypeOf(types.String{}),
		"target_services": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzureServiceEndpointRule
// only implements ToObjectValue() and Type().
func (o NccAzureServiceEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"subnets":         o.Subnets,
			"target_region":   o.TargetRegion,
			"target_services": o.TargetServices,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccAzureServiceEndpointRule) Type(ctx context.Context) attr.Type {
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
func (o *NccAzureServiceEndpointRule) GetSubnets(ctx context.Context) ([]types.String, bool) {
	if o.Subnets.IsNull() || o.Subnets.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Subnets.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetSubnets sets the value of the Subnets field in NccAzureServiceEndpointRule.
func (o *NccAzureServiceEndpointRule) SetSubnets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subnets = types.ListValueMust(t, vs)
}

// GetTargetServices returns the value of the TargetServices field in NccAzureServiceEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NccAzureServiceEndpointRule) GetTargetServices(ctx context.Context) ([]types.String, bool) {
	if o.TargetServices.IsNull() || o.TargetServices.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.TargetServices.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTargetServices sets the value of the TargetServices field in NccAzureServiceEndpointRule.
func (o *NccAzureServiceEndpointRule) SetTargetServices(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["target_services"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TargetServices = types.ListValueMust(t, vs)
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

func (newState *NccEgressConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressConfig) {
}

func (newState *NccEgressConfig) SyncEffectiveFieldsDuringRead(existingState NccEgressConfig) {
}

func (c NccEgressConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccEgressConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_rules": reflect.TypeOf(NccEgressDefaultRules{}),
		"target_rules":  reflect.TypeOf(NccEgressTargetRules{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressConfig
// only implements ToObjectValue() and Type().
func (o NccEgressConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_rules": o.DefaultRules,
			"target_rules":  o.TargetRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccEgressConfig) Type(ctx context.Context) attr.Type {
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
func (o *NccEgressConfig) GetDefaultRules(ctx context.Context) (NccEgressDefaultRules, bool) {
	var e NccEgressDefaultRules
	if o.DefaultRules.IsNull() || o.DefaultRules.IsUnknown() {
		return e, false
	}
	var v []NccEgressDefaultRules
	d := o.DefaultRules.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetDefaultRules sets the value of the DefaultRules field in NccEgressConfig.
func (o *NccEgressConfig) SetDefaultRules(ctx context.Context, v NccEgressDefaultRules) {
	vs := v.ToObjectValue(ctx)
	o.DefaultRules = vs
}

// GetTargetRules returns the value of the TargetRules field in NccEgressConfig as
// a NccEgressTargetRules value.
// If the field is unknown or null, the boolean return value is false.
func (o *NccEgressConfig) GetTargetRules(ctx context.Context) (NccEgressTargetRules, bool) {
	var e NccEgressTargetRules
	if o.TargetRules.IsNull() || o.TargetRules.IsUnknown() {
		return e, false
	}
	var v []NccEgressTargetRules
	d := o.TargetRules.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetTargetRules sets the value of the TargetRules field in NccEgressConfig.
func (o *NccEgressConfig) SetTargetRules(ctx context.Context, v NccEgressTargetRules) {
	vs := v.ToObjectValue(ctx)
	o.TargetRules = vs
}

// Default rules don't have specific targets.
type NccEgressDefaultRules struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule types.Object `tfsdk:"aws_stable_ip_rule"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule types.Object `tfsdk:"azure_service_endpoint_rule"`
}

func (newState *NccEgressDefaultRules) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressDefaultRules) {
}

func (newState *NccEgressDefaultRules) SyncEffectiveFieldsDuringRead(existingState NccEgressDefaultRules) {
}

func (c NccEgressDefaultRules) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccEgressDefaultRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_stable_ip_rule":          reflect.TypeOf(NccAwsStableIpRule{}),
		"azure_service_endpoint_rule": reflect.TypeOf(NccAzureServiceEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressDefaultRules
// only implements ToObjectValue() and Type().
func (o NccEgressDefaultRules) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_stable_ip_rule":          o.AwsStableIpRule,
			"azure_service_endpoint_rule": o.AzureServiceEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccEgressDefaultRules) Type(ctx context.Context) attr.Type {
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
func (o *NccEgressDefaultRules) GetAwsStableIpRule(ctx context.Context) (NccAwsStableIpRule, bool) {
	var e NccAwsStableIpRule
	if o.AwsStableIpRule.IsNull() || o.AwsStableIpRule.IsUnknown() {
		return e, false
	}
	var v []NccAwsStableIpRule
	d := o.AwsStableIpRule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAwsStableIpRule sets the value of the AwsStableIpRule field in NccEgressDefaultRules.
func (o *NccEgressDefaultRules) SetAwsStableIpRule(ctx context.Context, v NccAwsStableIpRule) {
	vs := v.ToObjectValue(ctx)
	o.AwsStableIpRule = vs
}

// GetAzureServiceEndpointRule returns the value of the AzureServiceEndpointRule field in NccEgressDefaultRules as
// a NccAzureServiceEndpointRule value.
// If the field is unknown or null, the boolean return value is false.
func (o *NccEgressDefaultRules) GetAzureServiceEndpointRule(ctx context.Context) (NccAzureServiceEndpointRule, bool) {
	var e NccAzureServiceEndpointRule
	if o.AzureServiceEndpointRule.IsNull() || o.AzureServiceEndpointRule.IsUnknown() {
		return e, false
	}
	var v []NccAzureServiceEndpointRule
	d := o.AzureServiceEndpointRule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetAzureServiceEndpointRule sets the value of the AzureServiceEndpointRule field in NccEgressDefaultRules.
func (o *NccEgressDefaultRules) SetAzureServiceEndpointRule(ctx context.Context, v NccAzureServiceEndpointRule) {
	vs := v.ToObjectValue(ctx)
	o.AzureServiceEndpointRule = vs
}

// Target rule controls the egress rules that are dedicated to specific
// resources.
type NccEgressTargetRules struct {
	// AWS private endpoint rule controls the AWS private endpoint based egress
	// rules.
	AwsPrivateEndpointRules types.List `tfsdk:"aws_private_endpoint_rules"`

	AzurePrivateEndpointRules types.List `tfsdk:"azure_private_endpoint_rules"`
}

func (newState *NccEgressTargetRules) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressTargetRules) {
}

func (newState *NccEgressTargetRules) SyncEffectiveFieldsDuringRead(existingState NccEgressTargetRules) {
}

func (c NccEgressTargetRules) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccEgressTargetRules) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_private_endpoint_rules":   reflect.TypeOf(CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule{}),
		"azure_private_endpoint_rules": reflect.TypeOf(NccAzurePrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressTargetRules
// only implements ToObjectValue() and Type().
func (o NccEgressTargetRules) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_private_endpoint_rules":   o.AwsPrivateEndpointRules,
			"azure_private_endpoint_rules": o.AzurePrivateEndpointRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccEgressTargetRules) Type(ctx context.Context) attr.Type {
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
func (o *NccEgressTargetRules) GetAwsPrivateEndpointRules(ctx context.Context) ([]CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule, bool) {
	if o.AwsPrivateEndpointRules.IsNull() || o.AwsPrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule
	d := o.AwsPrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAwsPrivateEndpointRules sets the value of the AwsPrivateEndpointRules field in NccEgressTargetRules.
func (o *NccEgressTargetRules) SetAwsPrivateEndpointRules(ctx context.Context, v []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_private_endpoint_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AwsPrivateEndpointRules = types.ListValueMust(t, vs)
}

// GetAzurePrivateEndpointRules returns the value of the AzurePrivateEndpointRules field in NccEgressTargetRules as
// a slice of NccAzurePrivateEndpointRule values.
// If the field is unknown or null, the boolean return value is false.
func (o *NccEgressTargetRules) GetAzurePrivateEndpointRules(ctx context.Context) ([]NccAzurePrivateEndpointRule, bool) {
	if o.AzurePrivateEndpointRules.IsNull() || o.AzurePrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []NccAzurePrivateEndpointRule
	d := o.AzurePrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzurePrivateEndpointRules sets the value of the AzurePrivateEndpointRules field in NccEgressTargetRules.
func (o *NccEgressTargetRules) SetAzurePrivateEndpointRules(ctx context.Context, v []NccAzurePrivateEndpointRule) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_private_endpoint_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AzurePrivateEndpointRules = types.ListValueMust(t, vs)
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

func (newState *NccPrivateEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccPrivateEndpointRule) {
}

func (newState *NccPrivateEndpointRule) SyncEffectiveFieldsDuringRead(existingState NccPrivateEndpointRule) {
}

func (c NccPrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccPrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccPrivateEndpointRule
// only implements ToObjectValue() and Type().
func (o NccPrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                     o.AccountId,
			"connection_state":               o.ConnectionState,
			"creation_time":                  o.CreationTime,
			"deactivated":                    o.Deactivated,
			"deactivated_at":                 o.DeactivatedAt,
			"domain_names":                   o.DomainNames,
			"enabled":                        o.Enabled,
			"endpoint_name":                  o.EndpointName,
			"endpoint_service":               o.EndpointService,
			"group_id":                       o.GroupId,
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"resource_id":                    o.ResourceId,
			"resource_names":                 o.ResourceNames,
			"rule_id":                        o.RuleId,
			"updated_time":                   o.UpdatedTime,
			"vpc_endpoint_id":                o.VpcEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccPrivateEndpointRule) Type(ctx context.Context) attr.Type {
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
func (o *NccPrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if o.DomainNames.IsNull() || o.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in NccPrivateEndpointRule.
func (o *NccPrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in NccPrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NccPrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if o.ResourceNames.IsNull() || o.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in NccPrivateEndpointRule.
func (o *NccPrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ResourceNames = types.ListValueMust(t, vs)
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

func (newState *NetworkConnectivityConfiguration) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkConnectivityConfiguration) {
}

func (newState *NetworkConnectivityConfiguration) SyncEffectiveFieldsDuringRead(existingState NetworkConnectivityConfiguration) {
}

func (c NetworkConnectivityConfiguration) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NetworkConnectivityConfiguration) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress_config": reflect.TypeOf(NccEgressConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkConnectivityConfiguration
// only implements ToObjectValue() and Type().
func (o NetworkConnectivityConfiguration) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"account_id":                     o.AccountId,
			"creation_time":                  o.CreationTime,
			"egress_config":                  o.EgressConfig,
			"name":                           o.Name,
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"region":                         o.Region,
			"updated_time":                   o.UpdatedTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkConnectivityConfiguration) Type(ctx context.Context) attr.Type {
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
func (o *NetworkConnectivityConfiguration) GetEgressConfig(ctx context.Context) (NccEgressConfig, bool) {
	var e NccEgressConfig
	if o.EgressConfig.IsNull() || o.EgressConfig.IsUnknown() {
		return e, false
	}
	var v []NccEgressConfig
	d := o.EgressConfig.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetEgressConfig sets the value of the EgressConfig field in NetworkConnectivityConfiguration.
func (o *NetworkConnectivityConfiguration) SetEgressConfig(ctx context.Context, v NccEgressConfig) {
	vs := v.ToObjectValue(ctx)
	o.EgressConfig = vs
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

func (newState *NetworkPolicyEgress) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkPolicyEgress) {
}

func (newState *NetworkPolicyEgress) SyncEffectiveFieldsDuringRead(existingState NetworkPolicyEgress) {
}

func (c NetworkPolicyEgress) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NetworkPolicyEgress) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_access": reflect.TypeOf(EgressNetworkPolicyNetworkAccessPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkPolicyEgress
// only implements ToObjectValue() and Type().
func (o NetworkPolicyEgress) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_access": o.NetworkAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NetworkPolicyEgress) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_access": EgressNetworkPolicyNetworkAccessPolicy{}.Type(ctx),
		},
	}
}

// GetNetworkAccess returns the value of the NetworkAccess field in NetworkPolicyEgress as
// a EgressNetworkPolicyNetworkAccessPolicy value.
// If the field is unknown or null, the boolean return value is false.
func (o *NetworkPolicyEgress) GetNetworkAccess(ctx context.Context) (EgressNetworkPolicyNetworkAccessPolicy, bool) {
	var e EgressNetworkPolicyNetworkAccessPolicy
	if o.NetworkAccess.IsNull() || o.NetworkAccess.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyNetworkAccessPolicy
	d := o.NetworkAccess.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNetworkAccess sets the value of the NetworkAccess field in NetworkPolicyEgress.
func (o *NetworkPolicyEgress) SetNetworkAccess(ctx context.Context, v EgressNetworkPolicyNetworkAccessPolicy) {
	vs := v.ToObjectValue(ctx)
	o.NetworkAccess = vs
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

func (newState *NotificationDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotificationDestination) {
}

func (newState *NotificationDestination) SyncEffectiveFieldsDuringRead(existingState NotificationDestination) {
}

func (c NotificationDestination) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NotificationDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationDestination
// only implements ToObjectValue() and Type().
func (o NotificationDestination) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":           o.Config,
			"destination_type": o.DestinationType,
			"display_name":     o.DisplayName,
			"id":               o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NotificationDestination) Type(ctx context.Context) attr.Type {
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
func (o *NotificationDestination) GetConfig(ctx context.Context) (Config, bool) {
	var e Config
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []Config
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetConfig sets the value of the Config field in NotificationDestination.
func (o *NotificationDestination) SetConfig(ctx context.Context, v Config) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey types.String `tfsdk:"integration_key"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet types.Bool `tfsdk:"integration_key_set"`
}

func (newState *PagerdutyConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan PagerdutyConfig) {
}

func (newState *PagerdutyConfig) SyncEffectiveFieldsDuringRead(existingState PagerdutyConfig) {
}

func (c PagerdutyConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PagerdutyConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PagerdutyConfig
// only implements ToObjectValue() and Type().
func (o PagerdutyConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_key":     o.IntegrationKey,
			"integration_key_set": o.IntegrationKeySet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PagerdutyConfig) Type(ctx context.Context) attr.Type {
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
	WorkspaceId types.Int64 `tfsdk:"workspaceId"`
}

func (newState *PartitionId) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionId) {
}

func (newState *PartitionId) SyncEffectiveFieldsDuringRead(existingState PartitionId) {
}

func (c PartitionId) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspaceId"] = attrs["workspaceId"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PartitionId.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PartitionId) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionId
// only implements ToObjectValue() and Type().
func (o PartitionId) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspaceId": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartitionId) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaceId": types.Int64Type,
		},
	}
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value types.String `tfsdk:"value"`
}

func (newState *PersonalComputeMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalComputeMessage) {
}

func (newState *PersonalComputeMessage) SyncEffectiveFieldsDuringRead(existingState PersonalComputeMessage) {
}

func (c PersonalComputeMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PersonalComputeMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeMessage
// only implements ToObjectValue() and Type().
func (o PersonalComputeMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalComputeMessage) Type(ctx context.Context) attr.Type {
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

func (newState *PersonalComputeSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalComputeSetting) {
}

func (newState *PersonalComputeSetting) SyncEffectiveFieldsDuringRead(existingState PersonalComputeSetting) {
}

func (c PersonalComputeSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PersonalComputeSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personal_compute": reflect.TypeOf(PersonalComputeMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeSetting
// only implements ToObjectValue() and Type().
func (o PersonalComputeSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             o.Etag,
			"personal_compute": o.PersonalCompute,
			"setting_name":     o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalComputeSetting) Type(ctx context.Context) attr.Type {
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
func (o *PersonalComputeSetting) GetPersonalCompute(ctx context.Context) (PersonalComputeMessage, bool) {
	var e PersonalComputeMessage
	if o.PersonalCompute.IsNull() || o.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeMessage
	d := o.PersonalCompute.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPersonalCompute sets the value of the PersonalCompute field in PersonalComputeSetting.
func (o *PersonalComputeSetting) SetPersonalCompute(ctx context.Context, v PersonalComputeMessage) {
	vs := v.ToObjectValue(ctx)
	o.PersonalCompute = vs
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

func (newState *PublicTokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublicTokenInfo) {
}

func (newState *PublicTokenInfo) SyncEffectiveFieldsDuringRead(existingState PublicTokenInfo) {
}

func (c PublicTokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PublicTokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublicTokenInfo
// only implements ToObjectValue() and Type().
func (o PublicTokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":       o.Comment,
			"creation_time": o.CreationTime,
			"expiry_time":   o.ExpiryTime,
			"token_id":      o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PublicTokenInfo) Type(ctx context.Context) attr.Type {
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
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type"`
}

func (newState *ReplaceIpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceIpAccessList) {
}

func (newState *ReplaceIpAccessList) SyncEffectiveFieldsDuringRead(existingState ReplaceIpAccessList) {
}

func (c ReplaceIpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetRequired()
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetRequired()
	attrs["list_type"] = attrs["list_type"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceIpAccessList
// only implements ToObjectValue() and Type().
func (o ReplaceIpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":           o.Enabled,
			"ip_access_list_id": o.IpAccessListId,
			"ip_addresses":      o.IpAddresses,
			"label":             o.Label,
			"list_type":         o.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ReplaceIpAccessList) Type(ctx context.Context) attr.Type {
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
func (o *ReplaceIpAccessList) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if o.IpAddresses.IsNull() || o.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in ReplaceIpAccessList.
func (o *ReplaceIpAccessList) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
}

type ReplaceResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceResponse
// only implements ToObjectValue() and Type().
func (o ReplaceResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ReplaceResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestrictWorkspaceAdminsMessage struct {
	Status types.String `tfsdk:"status"`
}

func (newState *RestrictWorkspaceAdminsMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestrictWorkspaceAdminsMessage) {
}

func (newState *RestrictWorkspaceAdminsMessage) SyncEffectiveFieldsDuringRead(existingState RestrictWorkspaceAdminsMessage) {
}

func (c RestrictWorkspaceAdminsMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestrictWorkspaceAdminsMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsMessage
// only implements ToObjectValue() and Type().
func (o RestrictWorkspaceAdminsMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestrictWorkspaceAdminsMessage) Type(ctx context.Context) attr.Type {
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

func (newState *RestrictWorkspaceAdminsSetting) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestrictWorkspaceAdminsSetting) {
}

func (newState *RestrictWorkspaceAdminsSetting) SyncEffectiveFieldsDuringRead(existingState RestrictWorkspaceAdminsSetting) {
}

func (c RestrictWorkspaceAdminsSetting) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestrictWorkspaceAdminsSetting) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"restrict_workspace_admins": reflect.TypeOf(RestrictWorkspaceAdminsMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsSetting
// only implements ToObjectValue() and Type().
func (o RestrictWorkspaceAdminsSetting) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                      o.Etag,
			"restrict_workspace_admins": o.RestrictWorkspaceAdmins,
			"setting_name":              o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestrictWorkspaceAdminsSetting) Type(ctx context.Context) attr.Type {
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
func (o *RestrictWorkspaceAdminsSetting) GetRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage, bool) {
	var e RestrictWorkspaceAdminsMessage
	if o.RestrictWorkspaceAdmins.IsNull() || o.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsMessage
	d := o.RestrictWorkspaceAdmins.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in RestrictWorkspaceAdminsSetting.
func (o *RestrictWorkspaceAdminsSetting) SetRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage) {
	vs := v.ToObjectValue(ctx)
	o.RestrictWorkspaceAdmins = vs
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id"`
}

func (newState *RevokeTokenRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenRequest) {
}

func (newState *RevokeTokenRequest) SyncEffectiveFieldsDuringRead(existingState RevokeTokenRequest) {
}

func (c RevokeTokenRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RevokeTokenRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenRequest
// only implements ToObjectValue() and Type().
func (o RevokeTokenRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RevokeTokenRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type RevokeTokenResponse struct {
}

func (newState *RevokeTokenResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenResponse) {
}

func (newState *RevokeTokenResponse) SyncEffectiveFieldsDuringRead(existingState RevokeTokenResponse) {
}

func (c RevokeTokenResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RevokeTokenResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenResponse
// only implements ToObjectValue() and Type().
func (o RevokeTokenResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RevokeTokenResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetStatusResponse struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetStatusResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetStatusResponse
// only implements ToObjectValue() and Type().
func (o SetStatusResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetStatusResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
}

func (newState *SlackConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan SlackConfig) {
}

func (newState *SlackConfig) SyncEffectiveFieldsDuringRead(existingState SlackConfig) {
}

func (c SlackConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SlackConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SlackConfig
// only implements ToObjectValue() and Type().
func (o SlackConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"url":     o.Url,
			"url_set": o.UrlSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SlackConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"url":     types.StringType,
			"url_set": types.BoolType,
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

func (newState *SqlResultsDownload) SyncEffectiveFieldsDuringCreateOrUpdate(plan SqlResultsDownload) {
}

func (newState *SqlResultsDownload) SyncEffectiveFieldsDuringRead(existingState SqlResultsDownload) {
}

func (c SqlResultsDownload) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SqlResultsDownload) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"boolean_val": reflect.TypeOf(BooleanMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SqlResultsDownload
// only implements ToObjectValue() and Type().
func (o SqlResultsDownload) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"boolean_val":  o.BooleanVal,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SqlResultsDownload) Type(ctx context.Context) attr.Type {
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
func (o *SqlResultsDownload) GetBooleanVal(ctx context.Context) (BooleanMessage, bool) {
	var e BooleanMessage
	if o.BooleanVal.IsNull() || o.BooleanVal.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage
	d := o.BooleanVal.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetBooleanVal sets the value of the BooleanVal field in SqlResultsDownload.
func (o *SqlResultsDownload) SetBooleanVal(ctx context.Context, v BooleanMessage) {
	vs := v.ToObjectValue(ctx)
	o.BooleanVal = vs
}

type StringMessage struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (newState *StringMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan StringMessage) {
}

func (newState *StringMessage) SyncEffectiveFieldsDuringRead(existingState StringMessage) {
}

func (c StringMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a StringMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage
// only implements ToObjectValue() and Type().
func (o StringMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StringMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *TokenAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessControlRequest) {
}

func (newState *TokenAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState TokenAccessControlRequest) {
}

func (c TokenAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlRequest
// only implements ToObjectValue() and Type().
func (o TokenAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_name":             o.GroupName,
			"permission_level":       o.PermissionLevel,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenAccessControlRequest) Type(ctx context.Context) attr.Type {
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

func (newState *TokenAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessControlResponse) {
}

func (newState *TokenAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState TokenAccessControlResponse) {
}

func (c TokenAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(TokenPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlResponse
// only implements ToObjectValue() and Type().
func (o TokenAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"all_permissions":        o.AllPermissions,
			"display_name":           o.DisplayName,
			"group_name":             o.GroupName,
			"service_principal_name": o.ServicePrincipalName,
			"user_name":              o.UserName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenAccessControlResponse) Type(ctx context.Context) attr.Type {
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
func (o *TokenAccessControlResponse) GetAllPermissions(ctx context.Context) ([]TokenPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []TokenPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in TokenAccessControlResponse.
func (o *TokenAccessControlResponse) SetAllPermissions(ctx context.Context, v []TokenPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
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

func (newState *TokenInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenInfo) {
}

func (newState *TokenInfo) SyncEffectiveFieldsDuringRead(existingState TokenInfo) {
}

func (c TokenInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo
// only implements ToObjectValue() and Type().
func (o TokenInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":             o.Comment,
			"created_by_id":       o.CreatedById,
			"created_by_username": o.CreatedByUsername,
			"creation_time":       o.CreationTime,
			"expiry_time":         o.ExpiryTime,
			"last_used_day":       o.LastUsedDay,
			"owner_id":            o.OwnerId,
			"token_id":            o.TokenId,
			"workspace_id":        o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenInfo) Type(ctx context.Context) attr.Type {
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
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *TokenPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermission) {
}

func (newState *TokenPermission) SyncEffectiveFieldsDuringRead(existingState TokenPermission) {
}

func (c TokenPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermission
// only implements ToObjectValue() and Type().
func (o TokenPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermission) Type(ctx context.Context) attr.Type {
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
func (o *TokenPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
	if o.InheritedFromObject.IsNull() || o.InheritedFromObject.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InheritedFromObject.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInheritedFromObject sets the value of the InheritedFromObject field in TokenPermission.
func (o *TokenPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type TokenPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *TokenPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissions) {
}

func (newState *TokenPermissions) SyncEffectiveFieldsDuringRead(existingState TokenPermissions) {
}

func (c TokenPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissions
// only implements ToObjectValue() and Type().
func (o TokenPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermissions) Type(ctx context.Context) attr.Type {
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
func (o *TokenPermissions) GetAccessControlList(ctx context.Context) ([]TokenAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissions.
func (o *TokenPermissions) SetAccessControlList(ctx context.Context, v []TokenAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type TokenPermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *TokenPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissionsDescription) {
}

func (newState *TokenPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState TokenPermissionsDescription) {
}

func (c TokenPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsDescription
// only implements ToObjectValue() and Type().
func (o TokenPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermissionsDescription) Type(ctx context.Context) attr.Type {
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

func (newState *TokenPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissionsRequest) {
}

func (newState *TokenPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState TokenPermissionsRequest) {
}

func (c TokenPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsRequest
// only implements ToObjectValue() and Type().
func (o TokenPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermissionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *TokenPermissionsRequest) GetAccessControlList(ctx context.Context) ([]TokenAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissionsRequest.
func (o *TokenPermissionsRequest) SetAccessControlList(ctx context.Context, v []TokenAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (newState *UpdateAccountIpAccessEnableRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAccountIpAccessEnableRequest) {
}

func (newState *UpdateAccountIpAccessEnableRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAccountIpAccessEnableRequest) {
}

func (c UpdateAccountIpAccessEnableRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateAccountIpAccessEnableRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AccountIpAccessEnable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountIpAccessEnableRequest
// only implements ToObjectValue() and Type().
func (o UpdateAccountIpAccessEnableRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAccountIpAccessEnableRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAccountIpAccessEnableRequest) GetSetting(ctx context.Context) (AccountIpAccessEnable, bool) {
	var e AccountIpAccessEnable
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AccountIpAccessEnable
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateAccountIpAccessEnableRequest.
func (o *UpdateAccountIpAccessEnableRequest) SetSetting(ctx context.Context, v AccountIpAccessEnable) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (newState *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) {
}

func (c UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicySetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingAccessPolicySettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) GetSetting(ctx context.Context) (AibiDashboardEmbeddingAccessPolicySetting, bool) {
	var e AibiDashboardEmbeddingAccessPolicySetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicySetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest.
func (o *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) SetSetting(ctx context.Context, v AibiDashboardEmbeddingAccessPolicySetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (newState *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) {
}

func (c UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomainsSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) GetSetting(ctx context.Context) (AibiDashboardEmbeddingApprovedDomainsSetting, bool) {
	var e AibiDashboardEmbeddingApprovedDomainsSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomainsSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest.
func (o *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) SetSetting(ctx context.Context, v AibiDashboardEmbeddingApprovedDomainsSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateAutomaticClusterUpdateSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAutomaticClusterUpdateSettingRequest) {
}

func (newState *UpdateAutomaticClusterUpdateSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAutomaticClusterUpdateSettingRequest) {
}

func (c UpdateAutomaticClusterUpdateSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAutomaticClusterUpdateSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AutomaticClusterUpdateSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAutomaticClusterUpdateSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateAutomaticClusterUpdateSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAutomaticClusterUpdateSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAutomaticClusterUpdateSettingRequest) GetSetting(ctx context.Context) (AutomaticClusterUpdateSetting, bool) {
	var e AutomaticClusterUpdateSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AutomaticClusterUpdateSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateAutomaticClusterUpdateSettingRequest.
func (o *UpdateAutomaticClusterUpdateSettingRequest) SetSetting(ctx context.Context, v AutomaticClusterUpdateSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateComplianceSecurityProfileSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateComplianceSecurityProfileSettingRequest) {
}

func (newState *UpdateComplianceSecurityProfileSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateComplianceSecurityProfileSettingRequest) {
}

func (c UpdateComplianceSecurityProfileSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateComplianceSecurityProfileSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(ComplianceSecurityProfileSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComplianceSecurityProfileSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateComplianceSecurityProfileSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateComplianceSecurityProfileSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateComplianceSecurityProfileSettingRequest) GetSetting(ctx context.Context) (ComplianceSecurityProfileSetting, bool) {
	var e ComplianceSecurityProfileSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfileSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateComplianceSecurityProfileSettingRequest.
func (o *UpdateComplianceSecurityProfileSettingRequest) SetSetting(ctx context.Context, v ComplianceSecurityProfileSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateCspEnablementAccountSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCspEnablementAccountSettingRequest) {
}

func (newState *UpdateCspEnablementAccountSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateCspEnablementAccountSettingRequest) {
}

func (c UpdateCspEnablementAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateCspEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(CspEnablementAccountSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCspEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateCspEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCspEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateCspEnablementAccountSettingRequest) GetSetting(ctx context.Context) (CspEnablementAccountSetting, bool) {
	var e CspEnablementAccountSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []CspEnablementAccountSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateCspEnablementAccountSettingRequest.
func (o *UpdateCspEnablementAccountSettingRequest) SetSetting(ctx context.Context, v CspEnablementAccountSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateDashboardEmailSubscriptionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDashboardEmailSubscriptionsRequest) {
}

func (newState *UpdateDashboardEmailSubscriptionsRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDashboardEmailSubscriptionsRequest) {
}

func (c UpdateDashboardEmailSubscriptionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDashboardEmailSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DashboardEmailSubscriptions{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDashboardEmailSubscriptionsRequest
// only implements ToObjectValue() and Type().
func (o UpdateDashboardEmailSubscriptionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDashboardEmailSubscriptionsRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDashboardEmailSubscriptionsRequest) GetSetting(ctx context.Context) (DashboardEmailSubscriptions, bool) {
	var e DashboardEmailSubscriptions
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DashboardEmailSubscriptions
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateDashboardEmailSubscriptionsRequest.
func (o *UpdateDashboardEmailSubscriptionsRequest) SetSetting(ctx context.Context, v DashboardEmailSubscriptions) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	Setting types.Object `tfsdk:"setting"`
}

func (newState *UpdateDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDefaultNamespaceSettingRequest) {
}

func (newState *UpdateDefaultNamespaceSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDefaultNamespaceSettingRequest) {
}

func (c UpdateDefaultNamespaceSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDefaultNamespaceSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultNamespaceSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultNamespaceSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateDefaultNamespaceSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDefaultNamespaceSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDefaultNamespaceSettingRequest) GetSetting(ctx context.Context) (DefaultNamespaceSetting, bool) {
	var e DefaultNamespaceSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DefaultNamespaceSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateDefaultNamespaceSettingRequest.
func (o *UpdateDefaultNamespaceSettingRequest) SetSetting(ctx context.Context, v DefaultNamespaceSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateDisableLegacyAccessRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyAccessRequest) {
}

func (newState *UpdateDisableLegacyAccessRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyAccessRequest) {
}

func (c UpdateDisableLegacyAccessRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDisableLegacyAccessRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyAccess{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyAccessRequest
// only implements ToObjectValue() and Type().
func (o UpdateDisableLegacyAccessRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDisableLegacyAccessRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDisableLegacyAccessRequest) GetSetting(ctx context.Context) (DisableLegacyAccess, bool) {
	var e DisableLegacyAccess
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyAccess
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateDisableLegacyAccessRequest.
func (o *UpdateDisableLegacyAccessRequest) SetSetting(ctx context.Context, v DisableLegacyAccess) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyDbfsRequest) {
}

func (newState *UpdateDisableLegacyDbfsRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyDbfsRequest) {
}

func (c UpdateDisableLegacyDbfsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDisableLegacyDbfsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyDbfs{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyDbfsRequest
// only implements ToObjectValue() and Type().
func (o UpdateDisableLegacyDbfsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDisableLegacyDbfsRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDisableLegacyDbfsRequest) GetSetting(ctx context.Context) (DisableLegacyDbfs, bool) {
	var e DisableLegacyDbfs
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyDbfs
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateDisableLegacyDbfsRequest.
func (o *UpdateDisableLegacyDbfsRequest) SetSetting(ctx context.Context, v DisableLegacyDbfs) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyFeaturesRequest) {
}

func (newState *UpdateDisableLegacyFeaturesRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyFeaturesRequest) {
}

func (c UpdateDisableLegacyFeaturesRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDisableLegacyFeaturesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyFeatures{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyFeaturesRequest
// only implements ToObjectValue() and Type().
func (o UpdateDisableLegacyFeaturesRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDisableLegacyFeaturesRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDisableLegacyFeaturesRequest) GetSetting(ctx context.Context) (DisableLegacyFeatures, bool) {
	var e DisableLegacyFeatures
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyFeatures
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateDisableLegacyFeaturesRequest.
func (o *UpdateDisableLegacyFeaturesRequest) SetSetting(ctx context.Context, v DisableLegacyFeatures) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateEnableExportNotebookRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEnableExportNotebookRequest) {
}

func (newState *UpdateEnableExportNotebookRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEnableExportNotebookRequest) {
}

func (c UpdateEnableExportNotebookRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateEnableExportNotebookRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableExportNotebook{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableExportNotebookRequest
// only implements ToObjectValue() and Type().
func (o UpdateEnableExportNotebookRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEnableExportNotebookRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEnableExportNotebookRequest) GetSetting(ctx context.Context) (EnableExportNotebook, bool) {
	var e EnableExportNotebook
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EnableExportNotebook
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateEnableExportNotebookRequest.
func (o *UpdateEnableExportNotebookRequest) SetSetting(ctx context.Context, v EnableExportNotebook) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateEnableNotebookTableClipboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEnableNotebookTableClipboardRequest) {
}

func (newState *UpdateEnableNotebookTableClipboardRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEnableNotebookTableClipboardRequest) {
}

func (c UpdateEnableNotebookTableClipboardRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateEnableNotebookTableClipboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableNotebookTableClipboard{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableNotebookTableClipboardRequest
// only implements ToObjectValue() and Type().
func (o UpdateEnableNotebookTableClipboardRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEnableNotebookTableClipboardRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEnableNotebookTableClipboardRequest) GetSetting(ctx context.Context) (EnableNotebookTableClipboard, bool) {
	var e EnableNotebookTableClipboard
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EnableNotebookTableClipboard
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateEnableNotebookTableClipboardRequest.
func (o *UpdateEnableNotebookTableClipboardRequest) SetSetting(ctx context.Context, v EnableNotebookTableClipboard) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateEnableResultsDownloadingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEnableResultsDownloadingRequest) {
}

func (newState *UpdateEnableResultsDownloadingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEnableResultsDownloadingRequest) {
}

func (c UpdateEnableResultsDownloadingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateEnableResultsDownloadingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnableResultsDownloading{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnableResultsDownloadingRequest
// only implements ToObjectValue() and Type().
func (o UpdateEnableResultsDownloadingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEnableResultsDownloadingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEnableResultsDownloadingRequest) GetSetting(ctx context.Context) (EnableResultsDownloading, bool) {
	var e EnableResultsDownloading
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EnableResultsDownloading
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateEnableResultsDownloadingRequest.
func (o *UpdateEnableResultsDownloadingRequest) SetSetting(ctx context.Context, v EnableResultsDownloading) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateEnhancedSecurityMonitoringSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEnhancedSecurityMonitoringSettingRequest) {
}

func (newState *UpdateEnhancedSecurityMonitoringSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEnhancedSecurityMonitoringSettingRequest) {
}

func (c UpdateEnhancedSecurityMonitoringSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateEnhancedSecurityMonitoringSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnhancedSecurityMonitoringSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnhancedSecurityMonitoringSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateEnhancedSecurityMonitoringSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEnhancedSecurityMonitoringSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEnhancedSecurityMonitoringSettingRequest) GetSetting(ctx context.Context) (EnhancedSecurityMonitoringSetting, bool) {
	var e EnhancedSecurityMonitoringSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EnhancedSecurityMonitoringSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateEnhancedSecurityMonitoringSettingRequest.
func (o *UpdateEnhancedSecurityMonitoringSettingRequest) SetSetting(ctx context.Context, v EnhancedSecurityMonitoringSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateEsmEnablementAccountSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEsmEnablementAccountSettingRequest) {
}

func (newState *UpdateEsmEnablementAccountSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateEsmEnablementAccountSettingRequest) {
}

func (c UpdateEsmEnablementAccountSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateEsmEnablementAccountSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EsmEnablementAccountSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEsmEnablementAccountSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateEsmEnablementAccountSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEsmEnablementAccountSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEsmEnablementAccountSettingRequest) GetSetting(ctx context.Context) (EsmEnablementAccountSetting, bool) {
	var e EsmEnablementAccountSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EsmEnablementAccountSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateEsmEnablementAccountSettingRequest.
func (o *UpdateEsmEnablementAccountSettingRequest) SetSetting(ctx context.Context, v EsmEnablementAccountSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type"`
}

func (newState *UpdateIpAccessList) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateIpAccessList) {
}

func (newState *UpdateIpAccessList) SyncEffectiveFieldsDuringRead(existingState UpdateIpAccessList) {
}

func (c UpdateIpAccessList) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["ip_access_list_id"] = attrs["ip_access_list_id"].SetRequired()
	attrs["ip_addresses"] = attrs["ip_addresses"].SetOptional()
	attrs["label"] = attrs["label"].SetOptional()
	attrs["list_type"] = attrs["list_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateIpAccessList.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateIpAccessList) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateIpAccessList
// only implements ToObjectValue() and Type().
func (o UpdateIpAccessList) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled":           o.Enabled,
			"ip_access_list_id": o.IpAccessListId,
			"ip_addresses":      o.IpAddresses,
			"label":             o.Label,
			"list_type":         o.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateIpAccessList) Type(ctx context.Context) attr.Type {
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
func (o *UpdateIpAccessList) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
	if o.IpAddresses.IsNull() || o.IpAddresses.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.IpAddresses.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAddresses sets the value of the IpAddresses field in UpdateIpAccessList.
func (o *UpdateIpAccessList) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
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

func (newState *UpdateLlmProxyPartnerPoweredAccountRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateLlmProxyPartnerPoweredAccountRequest) {
}

func (newState *UpdateLlmProxyPartnerPoweredAccountRequest) SyncEffectiveFieldsDuringRead(existingState UpdateLlmProxyPartnerPoweredAccountRequest) {
}

func (c UpdateLlmProxyPartnerPoweredAccountRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredAccountRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateLlmProxyPartnerPoweredAccountRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredAccount{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredAccountRequest
// only implements ToObjectValue() and Type().
func (o UpdateLlmProxyPartnerPoweredAccountRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateLlmProxyPartnerPoweredAccountRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateLlmProxyPartnerPoweredAccountRequest) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredAccount, bool) {
	var e LlmProxyPartnerPoweredAccount
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []LlmProxyPartnerPoweredAccount
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredAccountRequest.
func (o *UpdateLlmProxyPartnerPoweredAccountRequest) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredAccount) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateLlmProxyPartnerPoweredEnforceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateLlmProxyPartnerPoweredEnforceRequest) {
}

func (newState *UpdateLlmProxyPartnerPoweredEnforceRequest) SyncEffectiveFieldsDuringRead(existingState UpdateLlmProxyPartnerPoweredEnforceRequest) {
}

func (c UpdateLlmProxyPartnerPoweredEnforceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateLlmProxyPartnerPoweredEnforceRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateLlmProxyPartnerPoweredEnforceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredEnforce{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredEnforceRequest
// only implements ToObjectValue() and Type().
func (o UpdateLlmProxyPartnerPoweredEnforceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateLlmProxyPartnerPoweredEnforceRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateLlmProxyPartnerPoweredEnforceRequest) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredEnforce, bool) {
	var e LlmProxyPartnerPoweredEnforce
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []LlmProxyPartnerPoweredEnforce
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredEnforceRequest.
func (o *UpdateLlmProxyPartnerPoweredEnforceRequest) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredEnforce) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateLlmProxyPartnerPoweredWorkspaceRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateLlmProxyPartnerPoweredWorkspaceRequest) {
}

func (newState *UpdateLlmProxyPartnerPoweredWorkspaceRequest) SyncEffectiveFieldsDuringRead(existingState UpdateLlmProxyPartnerPoweredWorkspaceRequest) {
}

func (c UpdateLlmProxyPartnerPoweredWorkspaceRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateLlmProxyPartnerPoweredWorkspaceRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(LlmProxyPartnerPoweredWorkspace{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateLlmProxyPartnerPoweredWorkspaceRequest
// only implements ToObjectValue() and Type().
func (o UpdateLlmProxyPartnerPoweredWorkspaceRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateLlmProxyPartnerPoweredWorkspaceRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateLlmProxyPartnerPoweredWorkspaceRequest) GetSetting(ctx context.Context) (LlmProxyPartnerPoweredWorkspace, bool) {
	var e LlmProxyPartnerPoweredWorkspace
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []LlmProxyPartnerPoweredWorkspace
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateLlmProxyPartnerPoweredWorkspaceRequest.
func (o *UpdateLlmProxyPartnerPoweredWorkspaceRequest) SetSetting(ctx context.Context, v LlmProxyPartnerPoweredWorkspace) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

// Update a private endpoint rule
type UpdateNccPrivateEndpointRuleRequest struct {
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateNccPrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateNccPrivateEndpointRuleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"private_endpoint_rule": reflect.TypeOf(UpdatePrivateEndpointRule{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNccPrivateEndpointRuleRequest
// only implements ToObjectValue() and Type().
func (o UpdateNccPrivateEndpointRuleRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"private_endpoint_rule":          o.PrivateEndpointRule,
			"private_endpoint_rule_id":       o.PrivateEndpointRuleId,
			"update_mask":                    o.UpdateMask,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateNccPrivateEndpointRuleRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateNccPrivateEndpointRuleRequest) GetPrivateEndpointRule(ctx context.Context) (UpdatePrivateEndpointRule, bool) {
	var e UpdatePrivateEndpointRule
	if o.PrivateEndpointRule.IsNull() || o.PrivateEndpointRule.IsUnknown() {
		return e, false
	}
	var v []UpdatePrivateEndpointRule
	d := o.PrivateEndpointRule.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetPrivateEndpointRule sets the value of the PrivateEndpointRule field in UpdateNccPrivateEndpointRuleRequest.
func (o *UpdateNccPrivateEndpointRuleRequest) SetPrivateEndpointRule(ctx context.Context, v UpdatePrivateEndpointRule) {
	vs := v.ToObjectValue(ctx)
	o.PrivateEndpointRule = vs
}

// Update a network policy
type UpdateNetworkPolicyRequest struct {
	NetworkPolicy types.Object `tfsdk:"network_policy"`
	// The unique identifier for the network policy.
	NetworkPolicyId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateNetworkPolicyRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateNetworkPolicyRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"network_policy": reflect.TypeOf(AccountNetworkPolicy{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNetworkPolicyRequest
// only implements ToObjectValue() and Type().
func (o UpdateNetworkPolicyRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy":    o.NetworkPolicy,
			"network_policy_id": o.NetworkPolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateNetworkPolicyRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateNetworkPolicyRequest) GetNetworkPolicy(ctx context.Context) (AccountNetworkPolicy, bool) {
	var e AccountNetworkPolicy
	if o.NetworkPolicy.IsNull() || o.NetworkPolicy.IsUnknown() {
		return e, false
	}
	var v []AccountNetworkPolicy
	d := o.NetworkPolicy.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetNetworkPolicy sets the value of the NetworkPolicy field in UpdateNetworkPolicyRequest.
func (o *UpdateNetworkPolicyRequest) SetNetworkPolicy(ctx context.Context, v AccountNetworkPolicy) {
	vs := v.ToObjectValue(ctx)
	o.NetworkPolicy = vs
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

func (newState *UpdateNotificationDestinationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateNotificationDestinationRequest) {
}

func (newState *UpdateNotificationDestinationRequest) SyncEffectiveFieldsDuringRead(existingState UpdateNotificationDestinationRequest) {
}

func (c UpdateNotificationDestinationRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateNotificationDestinationRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNotificationDestinationRequest
// only implements ToObjectValue() and Type().
func (o UpdateNotificationDestinationRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       o.Config,
			"display_name": o.DisplayName,
			"id":           o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateNotificationDestinationRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateNotificationDestinationRequest) GetConfig(ctx context.Context) (Config, bool) {
	var e Config
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []Config
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetConfig sets the value of the Config field in UpdateNotificationDestinationRequest.
func (o *UpdateNotificationDestinationRequest) SetConfig(ctx context.Context, v Config) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
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

func (newState *UpdatePersonalComputeSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalComputeSettingRequest) {
}

func (newState *UpdatePersonalComputeSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalComputeSettingRequest) {
}

func (c UpdatePersonalComputeSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["allow_missing"] = attrs["allow_missing"].SetRequired()
	attrs["field_mask"] = attrs["field_mask"].SetRequired()
	attrs["setting"] = attrs["setting"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdatePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdatePersonalComputeSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(PersonalComputeSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalComputeSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdatePersonalComputeSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePersonalComputeSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdatePersonalComputeSettingRequest) GetSetting(ctx context.Context) (PersonalComputeSetting, bool) {
	var e PersonalComputeSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdatePersonalComputeSettingRequest.
func (o *UpdatePersonalComputeSettingRequest) SetSetting(ctx context.Context, v PersonalComputeSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdatePrivateEndpointRule) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePrivateEndpointRule) {
}

func (newState *UpdatePrivateEndpointRule) SyncEffectiveFieldsDuringRead(existingState UpdatePrivateEndpointRule) {
}

func (c UpdatePrivateEndpointRule) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdatePrivateEndpointRule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"domain_names":   reflect.TypeOf(types.String{}),
		"resource_names": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePrivateEndpointRule
// only implements ToObjectValue() and Type().
func (o UpdatePrivateEndpointRule) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"domain_names":   o.DomainNames,
			"enabled":        o.Enabled,
			"resource_names": o.ResourceNames,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePrivateEndpointRule) Type(ctx context.Context) attr.Type {
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
func (o *UpdatePrivateEndpointRule) GetDomainNames(ctx context.Context) ([]types.String, bool) {
	if o.DomainNames.IsNull() || o.DomainNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DomainNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDomainNames sets the value of the DomainNames field in UpdatePrivateEndpointRule.
func (o *UpdatePrivateEndpointRule) SetDomainNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["domain_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DomainNames = types.ListValueMust(t, vs)
}

// GetResourceNames returns the value of the ResourceNames field in UpdatePrivateEndpointRule as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdatePrivateEndpointRule) GetResourceNames(ctx context.Context) ([]types.String, bool) {
	if o.ResourceNames.IsNull() || o.ResourceNames.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ResourceNames.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResourceNames sets the value of the ResourceNames field in UpdatePrivateEndpointRule.
func (o *UpdatePrivateEndpointRule) SetResourceNames(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["resource_names"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ResourceNames = types.ListValueMust(t, vs)
}

type UpdateResponse struct {
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

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse
// only implements ToObjectValue() and Type().
func (o UpdateResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (newState *UpdateRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRestrictWorkspaceAdminsSettingRequest) {
}

func (newState *UpdateRestrictWorkspaceAdminsSettingRequest) SyncEffectiveFieldsDuringRead(existingState UpdateRestrictWorkspaceAdminsSettingRequest) {
}

func (c UpdateRestrictWorkspaceAdminsSettingRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRestrictWorkspaceAdminsSettingRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(RestrictWorkspaceAdminsSetting{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRestrictWorkspaceAdminsSettingRequest
// only implements ToObjectValue() and Type().
func (o UpdateRestrictWorkspaceAdminsSettingRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRestrictWorkspaceAdminsSettingRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateRestrictWorkspaceAdminsSettingRequest) GetSetting(ctx context.Context) (RestrictWorkspaceAdminsSetting, bool) {
	var e RestrictWorkspaceAdminsSetting
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsSetting
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateRestrictWorkspaceAdminsSettingRequest.
func (o *UpdateRestrictWorkspaceAdminsSettingRequest) SetSetting(ctx context.Context, v RestrictWorkspaceAdminsSetting) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
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

func (newState *UpdateSqlResultsDownloadRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateSqlResultsDownloadRequest) {
}

func (newState *UpdateSqlResultsDownloadRequest) SyncEffectiveFieldsDuringRead(existingState UpdateSqlResultsDownloadRequest) {
}

func (c UpdateSqlResultsDownloadRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateSqlResultsDownloadRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(SqlResultsDownload{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateSqlResultsDownloadRequest
// only implements ToObjectValue() and Type().
func (o UpdateSqlResultsDownloadRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateSqlResultsDownloadRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateSqlResultsDownloadRequest) GetSetting(ctx context.Context) (SqlResultsDownload, bool) {
	var e SqlResultsDownload
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []SqlResultsDownload
	d := o.Setting.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetSetting sets the value of the Setting field in UpdateSqlResultsDownloadRequest.
func (o *UpdateSqlResultsDownloadRequest) SetSetting(ctx context.Context, v SqlResultsDownload) {
	vs := v.ToObjectValue(ctx)
	o.Setting = vs
}

// Update workspace network option
type UpdateWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId types.Int64 `tfsdk:"-"`

	WorkspaceNetworkOption types.Object `tfsdk:"workspace_network_option"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateWorkspaceNetworkOptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateWorkspaceNetworkOptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_network_option": reflect.TypeOf(WorkspaceNetworkOption{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateWorkspaceNetworkOptionRequest
// only implements ToObjectValue() and Type().
func (o UpdateWorkspaceNetworkOptionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id":             o.WorkspaceId,
			"workspace_network_option": o.WorkspaceNetworkOption,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateWorkspaceNetworkOptionRequest) Type(ctx context.Context) attr.Type {
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
func (o *UpdateWorkspaceNetworkOptionRequest) GetWorkspaceNetworkOption(ctx context.Context) (WorkspaceNetworkOption, bool) {
	var e WorkspaceNetworkOption
	if o.WorkspaceNetworkOption.IsNull() || o.WorkspaceNetworkOption.IsUnknown() {
		return e, false
	}
	var v []WorkspaceNetworkOption
	d := o.WorkspaceNetworkOption.As(ctx, &v, basetypes.ObjectAsOptions{
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

// SetWorkspaceNetworkOption sets the value of the WorkspaceNetworkOption field in UpdateWorkspaceNetworkOptionRequest.
func (o *UpdateWorkspaceNetworkOptionRequest) SetWorkspaceNetworkOption(ctx context.Context, v WorkspaceNetworkOption) {
	vs := v.ToObjectValue(ctx)
	o.WorkspaceNetworkOption = vs
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

func (newState *WorkspaceNetworkOption) SyncEffectiveFieldsDuringCreateOrUpdate(plan WorkspaceNetworkOption) {
}

func (newState *WorkspaceNetworkOption) SyncEffectiveFieldsDuringRead(existingState WorkspaceNetworkOption) {
}

func (c WorkspaceNetworkOption) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a WorkspaceNetworkOption) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetworkOption
// only implements ToObjectValue() and Type().
func (o WorkspaceNetworkOption) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_policy_id": o.NetworkPolicyId,
			"workspace_id":      o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o WorkspaceNetworkOption) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_policy_id": types.StringType,
			"workspace_id":      types.Int64Type,
		},
	}
}
