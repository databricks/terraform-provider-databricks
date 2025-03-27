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

func (newState *AccountIpAccessEnable_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccountIpAccessEnable_SdkV2) {
}

func (newState *AccountIpAccessEnable_SdkV2) SyncEffectiveFieldsDuringRead(existingState AccountIpAccessEnable_SdkV2) {
}

func (c AccountIpAccessEnable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AccountIpAccessEnable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"acct_ip_acl_enable": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccountIpAccessEnable_SdkV2
// only implements ToObjectValue() and Type().
func (o AccountIpAccessEnable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"acct_ip_acl_enable": o.AcctIpAclEnable,
			"etag":               o.Etag,
			"setting_name":       o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AccountIpAccessEnable_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *AccountIpAccessEnable_SdkV2) GetAcctIpAclEnable(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.AcctIpAclEnable.IsNull() || o.AcctIpAclEnable.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.AcctIpAclEnable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAcctIpAclEnable sets the value of the AcctIpAclEnable field in AccountIpAccessEnable_SdkV2.
func (o *AccountIpAccessEnable_SdkV2) SetAcctIpAclEnable(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["acct_ip_acl_enable"]
	o.AcctIpAclEnable = types.ListValueMust(t, vs)
}

type AibiDashboardEmbeddingAccessPolicy_SdkV2 struct {
	AccessPolicyType types.String `tfsdk:"access_policy_type"`
}

func (newState *AibiDashboardEmbeddingAccessPolicy_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingAccessPolicy_SdkV2) {
}

func (newState *AibiDashboardEmbeddingAccessPolicy_SdkV2) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingAccessPolicy_SdkV2) {
}

func (c AibiDashboardEmbeddingAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_policy_type": o.AccessPolicyType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingAccessPolicySetting_SdkV2) {
}

func (newState *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingAccessPolicySetting_SdkV2) {
}

func (c AibiDashboardEmbeddingAccessPolicySetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingAccessPolicySetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_access_policy": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingAccessPolicySetting_SdkV2
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingAccessPolicySetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_access_policy": o.AibiDashboardEmbeddingAccessPolicy,
			"etag":                                   o.Etag,
			"setting_name":                           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingAccessPolicySetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) GetAibiDashboardEmbeddingAccessPolicy(ctx context.Context) (AibiDashboardEmbeddingAccessPolicy_SdkV2, bool) {
	var e AibiDashboardEmbeddingAccessPolicy_SdkV2
	if o.AibiDashboardEmbeddingAccessPolicy.IsNull() || o.AibiDashboardEmbeddingAccessPolicy.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicy_SdkV2
	d := o.AibiDashboardEmbeddingAccessPolicy.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAibiDashboardEmbeddingAccessPolicy sets the value of the AibiDashboardEmbeddingAccessPolicy field in AibiDashboardEmbeddingAccessPolicySetting_SdkV2.
func (o *AibiDashboardEmbeddingAccessPolicySetting_SdkV2) SetAibiDashboardEmbeddingAccessPolicy(ctx context.Context, v AibiDashboardEmbeddingAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aibi_dashboard_embedding_access_policy"]
	o.AibiDashboardEmbeddingAccessPolicy = types.ListValueMust(t, vs)
}

type AibiDashboardEmbeddingApprovedDomains_SdkV2 struct {
	ApprovedDomains types.List `tfsdk:"approved_domains"`
}

func (newState *AibiDashboardEmbeddingApprovedDomains_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingApprovedDomains_SdkV2) {
}

func (newState *AibiDashboardEmbeddingApprovedDomains_SdkV2) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingApprovedDomains_SdkV2) {
}

func (c AibiDashboardEmbeddingApprovedDomains_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingApprovedDomains_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"approved_domains": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomains_SdkV2
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingApprovedDomains_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"approved_domains": o.ApprovedDomains,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingApprovedDomains_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *AibiDashboardEmbeddingApprovedDomains_SdkV2) GetApprovedDomains(ctx context.Context) ([]types.String, bool) {
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

// SetApprovedDomains sets the value of the ApprovedDomains field in AibiDashboardEmbeddingApprovedDomains_SdkV2.
func (o *AibiDashboardEmbeddingApprovedDomains_SdkV2) SetApprovedDomains(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["approved_domains"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ApprovedDomains = types.ListValueMust(t, vs)
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

func (newState *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) {
}

func (newState *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) {
}

func (c AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aibi_dashboard_embedding_approved_domains": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomains_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aibi_dashboard_embedding_approved_domains": o.AibiDashboardEmbeddingApprovedDomains,
			"etag":         o.Etag,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) GetAibiDashboardEmbeddingApprovedDomains(ctx context.Context) (AibiDashboardEmbeddingApprovedDomains_SdkV2, bool) {
	var e AibiDashboardEmbeddingApprovedDomains_SdkV2
	if o.AibiDashboardEmbeddingApprovedDomains.IsNull() || o.AibiDashboardEmbeddingApprovedDomains.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomains_SdkV2
	d := o.AibiDashboardEmbeddingApprovedDomains.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAibiDashboardEmbeddingApprovedDomains sets the value of the AibiDashboardEmbeddingApprovedDomains field in AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2.
func (o *AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) SetAibiDashboardEmbeddingApprovedDomains(ctx context.Context, v AibiDashboardEmbeddingApprovedDomains_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aibi_dashboard_embedding_approved_domains"]
	o.AibiDashboardEmbeddingApprovedDomains = types.ListValueMust(t, vs)
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

func (newState *AutomaticClusterUpdateSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutomaticClusterUpdateSetting_SdkV2) {
}

func (newState *AutomaticClusterUpdateSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState AutomaticClusterUpdateSetting_SdkV2) {
}

func (c AutomaticClusterUpdateSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AutomaticClusterUpdateSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"automatic_cluster_update_workspace": reflect.TypeOf(ClusterAutoRestartMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutomaticClusterUpdateSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o AutomaticClusterUpdateSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"automatic_cluster_update_workspace": o.AutomaticClusterUpdateWorkspace,
			"etag":                               o.Etag,
			"setting_name":                       o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutomaticClusterUpdateSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *AutomaticClusterUpdateSetting_SdkV2) GetAutomaticClusterUpdateWorkspace(ctx context.Context) (ClusterAutoRestartMessage_SdkV2, bool) {
	var e ClusterAutoRestartMessage_SdkV2
	if o.AutomaticClusterUpdateWorkspace.IsNull() || o.AutomaticClusterUpdateWorkspace.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessage_SdkV2
	d := o.AutomaticClusterUpdateWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutomaticClusterUpdateWorkspace sets the value of the AutomaticClusterUpdateWorkspace field in AutomaticClusterUpdateSetting_SdkV2.
func (o *AutomaticClusterUpdateSetting_SdkV2) SetAutomaticClusterUpdateWorkspace(ctx context.Context, v ClusterAutoRestartMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["automatic_cluster_update_workspace"]
	o.AutomaticClusterUpdateWorkspace = types.ListValueMust(t, vs)
}

type BooleanMessage_SdkV2 struct {
	Value types.Bool `tfsdk:"value"`
}

func (newState *BooleanMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BooleanMessage_SdkV2) {
}

func (newState *BooleanMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState BooleanMessage_SdkV2) {
}

func (c BooleanMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BooleanMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BooleanMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o BooleanMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BooleanMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.BoolType,
		},
	}
}

type ClusterAutoRestartMessage_SdkV2 struct {
	CanToggle types.Bool `tfsdk:"can_toggle"`

	Enabled types.Bool `tfsdk:"enabled"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails types.List `tfsdk:"enablement_details"`

	MaintenanceWindow types.List `tfsdk:"maintenance_window"`

	RestartEvenIfNoUpdatesAvailable types.Bool `tfsdk:"restart_even_if_no_updates_available"`
}

func (newState *ClusterAutoRestartMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessage_SdkV2) {
}

func (newState *ClusterAutoRestartMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessage_SdkV2) {
}

func (c ClusterAutoRestartMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enablement_details": reflect.TypeOf(ClusterAutoRestartMessageEnablementDetails_SdkV2{}),
		"maintenance_window": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindow_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ClusterAutoRestartMessage_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAutoRestartMessage_SdkV2) GetEnablementDetails(ctx context.Context) (ClusterAutoRestartMessageEnablementDetails_SdkV2, bool) {
	var e ClusterAutoRestartMessageEnablementDetails_SdkV2
	if o.EnablementDetails.IsNull() || o.EnablementDetails.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageEnablementDetails_SdkV2
	d := o.EnablementDetails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnablementDetails sets the value of the EnablementDetails field in ClusterAutoRestartMessage_SdkV2.
func (o *ClusterAutoRestartMessage_SdkV2) SetEnablementDetails(ctx context.Context, v ClusterAutoRestartMessageEnablementDetails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enablement_details"]
	o.EnablementDetails = types.ListValueMust(t, vs)
}

// GetMaintenanceWindow returns the value of the MaintenanceWindow field in ClusterAutoRestartMessage_SdkV2 as
// a ClusterAutoRestartMessageMaintenanceWindow_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ClusterAutoRestartMessage_SdkV2) GetMaintenanceWindow(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindow_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindow_SdkV2
	if o.MaintenanceWindow.IsNull() || o.MaintenanceWindow.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindow_SdkV2
	d := o.MaintenanceWindow.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMaintenanceWindow sets the value of the MaintenanceWindow field in ClusterAutoRestartMessage_SdkV2.
func (o *ClusterAutoRestartMessage_SdkV2) SetMaintenanceWindow(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["maintenance_window"]
	o.MaintenanceWindow = types.ListValueMust(t, vs)
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

func (newState *ClusterAutoRestartMessageEnablementDetails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageEnablementDetails_SdkV2) {
}

func (newState *ClusterAutoRestartMessageEnablementDetails_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageEnablementDetails_SdkV2) {
}

func (c ClusterAutoRestartMessageEnablementDetails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageEnablementDetails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageEnablementDetails_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageEnablementDetails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"forced_for_compliance_mode":           o.ForcedForComplianceMode,
			"unavailable_for_disabled_entitlement": o.UnavailableForDisabledEntitlement,
			"unavailable_for_non_enterprise_tier":  o.UnavailableForNonEnterpriseTier,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageEnablementDetails_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindow_SdkV2) {
}

func (c ClusterAutoRestartMessageMaintenanceWindow_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindow_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"week_day_based_schedule": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindow_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindow_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"week_day_based_schedule": o.WeekDayBasedSchedule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindow_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) GetWeekDayBasedSchedule(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
	if o.WeekDayBasedSchedule.IsNull() || o.WeekDayBasedSchedule.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
	d := o.WeekDayBasedSchedule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWeekDayBasedSchedule sets the value of the WeekDayBasedSchedule field in ClusterAutoRestartMessageMaintenanceWindow_SdkV2.
func (o *ClusterAutoRestartMessageMaintenanceWindow_SdkV2) SetWeekDayBasedSchedule(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["week_day_based_schedule"]
	o.WeekDayBasedSchedule = types.ListValueMust(t, vs)
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2 struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`

	Frequency types.String `tfsdk:"frequency"`

	WindowStartTime types.List `tfsdk:"window_start_time"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) {
}

func (c ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"window_start_time": reflect.TypeOf(ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"day_of_week":       o.DayOfWeek,
			"frequency":         o.Frequency,
			"window_start_time": o.WindowStartTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) GetWindowStartTime(ctx context.Context) (ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2, bool) {
	var e ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
	if o.WindowStartTime.IsNull() || o.WindowStartTime.IsUnknown() {
		return e, false
	}
	var v []ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
	d := o.WindowStartTime.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetWindowStartTime sets the value of the WindowStartTime field in ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2.
func (o *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule_SdkV2) SetWindowStartTime(ctx context.Context, v ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["window_start_time"]
	o.WindowStartTime = types.ListValueMust(t, vs)
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2 struct {
	Hours types.Int64 `tfsdk:"hours"`

	Minutes types.Int64 `tfsdk:"minutes"`
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
}

func (newState *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) SyncEffectiveFieldsDuringRead(existingState ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) {
}

func (c ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2
// only implements ToObjectValue() and Type().
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"hours":   o.Hours,
			"minutes": o.Minutes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ClusterAutoRestartMessageMaintenanceWindowWindowStartTime_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *ComplianceSecurityProfile_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfile_SdkV2) {
}

func (newState *ComplianceSecurityProfile_SdkV2) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfile_SdkV2) {
}

func (c ComplianceSecurityProfile_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComplianceSecurityProfile_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfile_SdkV2
// only implements ToObjectValue() and Type().
func (o ComplianceSecurityProfile_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": o.ComplianceStandards,
			"is_enabled":           o.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplianceSecurityProfile_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ComplianceSecurityProfile_SdkV2) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
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

// SetComplianceStandards sets the value of the ComplianceStandards field in ComplianceSecurityProfile_SdkV2.
func (o *ComplianceSecurityProfile_SdkV2) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ComplianceStandards = types.ListValueMust(t, vs)
}

type ComplianceSecurityProfileSetting_SdkV2 struct {
	// SHIELD feature: CSP
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

func (newState *ComplianceSecurityProfileSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ComplianceSecurityProfileSetting_SdkV2) {
}

func (newState *ComplianceSecurityProfileSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState ComplianceSecurityProfileSetting_SdkV2) {
}

func (c ComplianceSecurityProfileSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ComplianceSecurityProfileSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_security_profile_workspace": reflect.TypeOf(ComplianceSecurityProfile_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ComplianceSecurityProfileSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o ComplianceSecurityProfileSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_security_profile_workspace": o.ComplianceSecurityProfileWorkspace,
			"etag":                                  o.Etag,
			"setting_name":                          o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ComplianceSecurityProfileSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ComplianceSecurityProfileSetting_SdkV2) GetComplianceSecurityProfileWorkspace(ctx context.Context) (ComplianceSecurityProfile_SdkV2, bool) {
	var e ComplianceSecurityProfile_SdkV2
	if o.ComplianceSecurityProfileWorkspace.IsNull() || o.ComplianceSecurityProfileWorkspace.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfile_SdkV2
	d := o.ComplianceSecurityProfileWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetComplianceSecurityProfileWorkspace sets the value of the ComplianceSecurityProfileWorkspace field in ComplianceSecurityProfileSetting_SdkV2.
func (o *ComplianceSecurityProfileSetting_SdkV2) SetComplianceSecurityProfileWorkspace(ctx context.Context, v ComplianceSecurityProfile_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_security_profile_workspace"]
	o.ComplianceSecurityProfileWorkspace = types.ListValueMust(t, vs)
}

type Config_SdkV2 struct {
	Email types.List `tfsdk:"email"`

	GenericWebhook types.List `tfsdk:"generic_webhook"`

	MicrosoftTeams types.List `tfsdk:"microsoft_teams"`

	Pagerduty types.List `tfsdk:"pagerduty"`

	Slack types.List `tfsdk:"slack"`
}

func (newState *Config_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Config_SdkV2) {
}

func (newState *Config_SdkV2) SyncEffectiveFieldsDuringRead(existingState Config_SdkV2) {
}

func (c Config_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Config_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (o Config_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o Config_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *Config_SdkV2) GetEmail(ctx context.Context) (EmailConfig_SdkV2, bool) {
	var e EmailConfig_SdkV2
	if o.Email.IsNull() || o.Email.IsUnknown() {
		return e, false
	}
	var v []EmailConfig_SdkV2
	d := o.Email.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmail sets the value of the Email field in Config_SdkV2.
func (o *Config_SdkV2) SetEmail(ctx context.Context, v EmailConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["email"]
	o.Email = types.ListValueMust(t, vs)
}

// GetGenericWebhook returns the value of the GenericWebhook field in Config_SdkV2 as
// a GenericWebhookConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config_SdkV2) GetGenericWebhook(ctx context.Context) (GenericWebhookConfig_SdkV2, bool) {
	var e GenericWebhookConfig_SdkV2
	if o.GenericWebhook.IsNull() || o.GenericWebhook.IsUnknown() {
		return e, false
	}
	var v []GenericWebhookConfig_SdkV2
	d := o.GenericWebhook.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGenericWebhook sets the value of the GenericWebhook field in Config_SdkV2.
func (o *Config_SdkV2) SetGenericWebhook(ctx context.Context, v GenericWebhookConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["generic_webhook"]
	o.GenericWebhook = types.ListValueMust(t, vs)
}

// GetMicrosoftTeams returns the value of the MicrosoftTeams field in Config_SdkV2 as
// a MicrosoftTeamsConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config_SdkV2) GetMicrosoftTeams(ctx context.Context) (MicrosoftTeamsConfig_SdkV2, bool) {
	var e MicrosoftTeamsConfig_SdkV2
	if o.MicrosoftTeams.IsNull() || o.MicrosoftTeams.IsUnknown() {
		return e, false
	}
	var v []MicrosoftTeamsConfig_SdkV2
	d := o.MicrosoftTeams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMicrosoftTeams sets the value of the MicrosoftTeams field in Config_SdkV2.
func (o *Config_SdkV2) SetMicrosoftTeams(ctx context.Context, v MicrosoftTeamsConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["microsoft_teams"]
	o.MicrosoftTeams = types.ListValueMust(t, vs)
}

// GetPagerduty returns the value of the Pagerduty field in Config_SdkV2 as
// a PagerdutyConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config_SdkV2) GetPagerduty(ctx context.Context) (PagerdutyConfig_SdkV2, bool) {
	var e PagerdutyConfig_SdkV2
	if o.Pagerduty.IsNull() || o.Pagerduty.IsUnknown() {
		return e, false
	}
	var v []PagerdutyConfig_SdkV2
	d := o.Pagerduty.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPagerduty sets the value of the Pagerduty field in Config_SdkV2.
func (o *Config_SdkV2) SetPagerduty(ctx context.Context, v PagerdutyConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pagerduty"]
	o.Pagerduty = types.ListValueMust(t, vs)
}

// GetSlack returns the value of the Slack field in Config_SdkV2 as
// a SlackConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *Config_SdkV2) GetSlack(ctx context.Context) (SlackConfig_SdkV2, bool) {
	var e SlackConfig_SdkV2
	if o.Slack.IsNull() || o.Slack.IsUnknown() {
		return e, false
	}
	var v []SlackConfig_SdkV2
	d := o.Slack.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSlack sets the value of the Slack field in Config_SdkV2.
func (o *Config_SdkV2) SetSlack(ctx context.Context, v SlackConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["slack"]
	o.Slack = types.ListValueMust(t, vs)
}

// Details required to configure a block list or allow list.
type CreateIpAccessList_SdkV2 struct {
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

func (newState *CreateIpAccessList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateIpAccessList_SdkV2) {
}

func (newState *CreateIpAccessList_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateIpAccessList_SdkV2) {
}

func (c CreateIpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateIpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateIpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_addresses": o.IpAddresses,
			"label":        o.Label,
			"list_type":    o.ListType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateIpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateIpAccessList_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in CreateIpAccessList_SdkV2.
func (o *CreateIpAccessList_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
}

// An IP access list was successfully created.
type CreateIpAccessListResponse_SdkV2 struct {
	// Definition of an IP Access list
	IpAccessList types.List `tfsdk:"ip_access_list"`
}

func (newState *CreateIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateIpAccessListResponse_SdkV2) {
}

func (newState *CreateIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateIpAccessListResponse_SdkV2) {
}

func (c CreateIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": o.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateIpAccessListResponse_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessListInfo_SdkV2, bool) {
	var e IpAccessListInfo_SdkV2
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo_SdkV2
	d := o.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in CreateIpAccessListResponse_SdkV2.
func (o *CreateIpAccessListResponse_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessListInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	o.IpAccessList = types.ListValueMust(t, vs)
}

type CreateNetworkConnectivityConfigRequest_SdkV2 struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name types.String `tfsdk:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region types.String `tfsdk:"region"`
}

func (newState *CreateNetworkConnectivityConfigRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNetworkConnectivityConfigRequest_SdkV2) {
}

func (newState *CreateNetworkConnectivityConfigRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateNetworkConnectivityConfigRequest_SdkV2) {
}

func (c CreateNetworkConnectivityConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["region"] = attrs["region"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateNetworkConnectivityConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateNetworkConnectivityConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNetworkConnectivityConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateNetworkConnectivityConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":   o.Name,
			"region": o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNetworkConnectivityConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":   types.StringType,
			"region": types.StringType,
		},
	}
}

type CreateNotificationDestinationRequest_SdkV2 struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config types.List `tfsdk:"config"`
	// The display name for the notification destination.
	DisplayName types.String `tfsdk:"display_name"`
}

func (newState *CreateNotificationDestinationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateNotificationDestinationRequest_SdkV2) {
}

func (newState *CreateNotificationDestinationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateNotificationDestinationRequest_SdkV2) {
}

func (c CreateNotificationDestinationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       o.Config,
			"display_name": o.DisplayName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateNotificationDestinationRequest_SdkV2) GetConfig(ctx context.Context) (Config_SdkV2, bool) {
	var e Config_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []Config_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in CreateNotificationDestinationRequest_SdkV2.
func (o *CreateNotificationDestinationRequest_SdkV2) SetConfig(ctx context.Context, v Config_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
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

func (newState *CreateOboTokenRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOboTokenRequest_SdkV2) {
}

func (newState *CreateOboTokenRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateOboTokenRequest_SdkV2) {
}

func (c CreateOboTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateOboTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateOboTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"application_id":   o.ApplicationId,
			"comment":          o.Comment,
			"lifetime_seconds": o.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOboTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *CreateOboTokenResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateOboTokenResponse_SdkV2) {
}

func (newState *CreateOboTokenResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateOboTokenResponse_SdkV2) {
}

func (c CreateOboTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateOboTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateOboTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateOboTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  o.TokenInfo,
			"token_value": o.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateOboTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateOboTokenResponse_SdkV2) GetTokenInfo(ctx context.Context) (TokenInfo_SdkV2, bool) {
	var e TokenInfo_SdkV2
	if o.TokenInfo.IsNull() || o.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []TokenInfo_SdkV2
	d := o.TokenInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenInfo sets the value of the TokenInfo field in CreateOboTokenResponse_SdkV2.
func (o *CreateOboTokenResponse_SdkV2) SetTokenInfo(ctx context.Context, v TokenInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_info"]
	o.TokenInfo = types.ListValueMust(t, vs)
}

type CreatePrivateEndpointRuleRequest_SdkV2 struct {
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId types.String `tfsdk:"group_id"`
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
	// The Azure resource ID of the target resource.
	ResourceId types.String `tfsdk:"resource_id"`
}

func (newState *CreatePrivateEndpointRuleRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreatePrivateEndpointRuleRequest_SdkV2) {
}

func (newState *CreatePrivateEndpointRuleRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreatePrivateEndpointRuleRequest_SdkV2) {
}

func (c CreatePrivateEndpointRuleRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_id"] = attrs["group_id"].SetRequired()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetRequired()
	attrs["resource_id"] = attrs["resource_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePrivateEndpointRuleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreatePrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreatePrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"group_id":                       o.GroupId,
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"resource_id":                    o.ResourceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_id":                       types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_id":                    types.StringType,
		},
	}
}

type CreateTokenRequest_SdkV2 struct {
	// Optional description to attach to the token.
	Comment types.String `tfsdk:"comment"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds types.Int64 `tfsdk:"lifetime_seconds"`
}

func (newState *CreateTokenRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTokenRequest_SdkV2) {
}

func (newState *CreateTokenRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateTokenRequest_SdkV2) {
}

func (c CreateTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"comment":          o.Comment,
			"lifetime_seconds": o.LifetimeSeconds,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *CreateTokenResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateTokenResponse_SdkV2) {
}

func (newState *CreateTokenResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateTokenResponse_SdkV2) {
}

func (c CreateTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CreateTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(PublicTokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info":  o.TokenInfo,
			"token_value": o.TokenValue,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CreateTokenResponse_SdkV2) GetTokenInfo(ctx context.Context) (PublicTokenInfo_SdkV2, bool) {
	var e PublicTokenInfo_SdkV2
	if o.TokenInfo.IsNull() || o.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []PublicTokenInfo_SdkV2
	d := o.TokenInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenInfo sets the value of the TokenInfo field in CreateTokenResponse_SdkV2.
func (o *CreateTokenResponse_SdkV2) SetTokenInfo(ctx context.Context, v PublicTokenInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_info"]
	o.TokenInfo = types.ListValueMust(t, vs)
}

// Account level policy for CSP
type CspEnablementAccount_SdkV2 struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards types.List `tfsdk:"compliance_standards"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (newState *CspEnablementAccount_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CspEnablementAccount_SdkV2) {
}

func (newState *CspEnablementAccount_SdkV2) SyncEffectiveFieldsDuringRead(existingState CspEnablementAccount_SdkV2) {
}

func (c CspEnablementAccount_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CspEnablementAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"compliance_standards": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccount_SdkV2
// only implements ToObjectValue() and Type().
func (o CspEnablementAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"compliance_standards": o.ComplianceStandards,
			"is_enforced":          o.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CspEnablementAccount_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CspEnablementAccount_SdkV2) GetComplianceStandards(ctx context.Context) ([]types.String, bool) {
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

// SetComplianceStandards sets the value of the ComplianceStandards field in CspEnablementAccount_SdkV2.
func (o *CspEnablementAccount_SdkV2) SetComplianceStandards(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["compliance_standards"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ComplianceStandards = types.ListValueMust(t, vs)
}

type CspEnablementAccountSetting_SdkV2 struct {
	// Account level policy for CSP
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

func (newState *CspEnablementAccountSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CspEnablementAccountSetting_SdkV2) {
}

func (newState *CspEnablementAccountSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState CspEnablementAccountSetting_SdkV2) {
}

func (c CspEnablementAccountSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CspEnablementAccountSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"csp_enablement_account": reflect.TypeOf(CspEnablementAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CspEnablementAccountSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o CspEnablementAccountSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"csp_enablement_account": o.CspEnablementAccount,
			"etag":                   o.Etag,
			"setting_name":           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CspEnablementAccountSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *CspEnablementAccountSetting_SdkV2) GetCspEnablementAccount(ctx context.Context) (CspEnablementAccount_SdkV2, bool) {
	var e CspEnablementAccount_SdkV2
	if o.CspEnablementAccount.IsNull() || o.CspEnablementAccount.IsUnknown() {
		return e, false
	}
	var v []CspEnablementAccount_SdkV2
	d := o.CspEnablementAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCspEnablementAccount sets the value of the CspEnablementAccount field in CspEnablementAccountSetting_SdkV2.
func (o *CspEnablementAccountSetting_SdkV2) SetCspEnablementAccount(ctx context.Context, v CspEnablementAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["csp_enablement_account"]
	o.CspEnablementAccount = types.ListValueMust(t, vs)
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

func (newState *DefaultNamespaceSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DefaultNamespaceSetting_SdkV2) {
}

func (newState *DefaultNamespaceSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState DefaultNamespaceSetting_SdkV2) {
}

func (c DefaultNamespaceSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DefaultNamespaceSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"namespace": reflect.TypeOf(StringMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DefaultNamespaceSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o DefaultNamespaceSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":         o.Etag,
			"namespace":    o.Namespace,
			"setting_name": o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DefaultNamespaceSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DefaultNamespaceSetting_SdkV2) GetNamespace(ctx context.Context) (StringMessage_SdkV2, bool) {
	var e StringMessage_SdkV2
	if o.Namespace.IsNull() || o.Namespace.IsUnknown() {
		return e, false
	}
	var v []StringMessage_SdkV2
	d := o.Namespace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetNamespace sets the value of the Namespace field in DefaultNamespaceSetting_SdkV2.
func (o *DefaultNamespaceSetting_SdkV2) SetNamespace(ctx context.Context, v StringMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["namespace"]
	o.Namespace = types.ListValueMust(t, vs)
}

// Delete the account IP access toggle setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAccountIpAccessEnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountIpAccessEnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountIpAccessEnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteAccountIpAccessEnableResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAccountIpAccessEnableResponse_SdkV2) {
}

func (newState *DeleteAccountIpAccessEnableResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAccountIpAccessEnableResponse_SdkV2) {
}

func (c DeleteAccountIpAccessEnableResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAccountIpAccessEnableResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessEnableResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountIpAccessEnableResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountIpAccessEnableResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete access list
type DeleteAccountIpAccessListRequest_SdkV2 struct {
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
func (a DeleteAccountIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAccountIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAccountIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAccountIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

// Delete the AI/BI dashboard embedding access policy
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) {
}

func (newState *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) {
}

func (c DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingAccessPolicySettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete AI/BI dashboard embedding approved domains
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) {
}

func (newState *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) {
}

func (c DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the default namespace setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDefaultNamespaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDefaultNamespaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDefaultNamespaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDefaultNamespaceSettingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDefaultNamespaceSettingResponse_SdkV2) {
}

func (newState *DeleteDefaultNamespaceSettingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDefaultNamespaceSettingResponse_SdkV2) {
}

func (c DeleteDefaultNamespaceSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDefaultNamespaceSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDefaultNamespaceSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDefaultNamespaceSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDefaultNamespaceSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete Legacy Access Disablement Status
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDisableLegacyAccessResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyAccessResponse_SdkV2) {
}

func (newState *DeleteDisableLegacyAccessResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyAccessResponse_SdkV2) {
}

func (c DeleteDisableLegacyAccessResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDisableLegacyAccessResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyAccessResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyAccessResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyAccessResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the disable legacy DBFS setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDisableLegacyDbfsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyDbfsResponse_SdkV2) {
}

func (newState *DeleteDisableLegacyDbfsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyDbfsResponse_SdkV2) {
}

func (c DeleteDisableLegacyDbfsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDisableLegacyDbfsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyDbfsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyDbfsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyDbfsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete the disable legacy features setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteDisableLegacyFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteDisableLegacyFeaturesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDisableLegacyFeaturesResponse_SdkV2) {
}

func (newState *DeleteDisableLegacyFeaturesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteDisableLegacyFeaturesResponse_SdkV2) {
}

func (c DeleteDisableLegacyFeaturesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteDisableLegacyFeaturesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteDisableLegacyFeaturesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteDisableLegacyFeaturesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteDisableLegacyFeaturesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete access list
type DeleteIpAccessListRequest_SdkV2 struct {
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
func (a DeleteIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest_SdkV2 struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkConnectivityConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkConnectivityConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteNetworkConnectivityConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkConnectivityConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

type DeleteNetworkConnectivityConfigurationResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNetworkConnectivityConfigurationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNetworkConnectivityConfigurationResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNetworkConnectivityConfigurationResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteNetworkConnectivityConfigurationResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNetworkConnectivityConfigurationResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a notification destination
type DeleteNotificationDestinationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Delete Personal Compute setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeletePersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeletePersonalComputeSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePersonalComputeSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePersonalComputeSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeletePersonalComputeSettingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeletePersonalComputeSettingResponse_SdkV2) {
}

func (newState *DeletePersonalComputeSettingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeletePersonalComputeSettingResponse_SdkV2) {
}

func (c DeletePersonalComputeSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeletePersonalComputeSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePersonalComputeSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePersonalComputeSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePersonalComputeSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest_SdkV2 struct {
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
func (a DeletePrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeletePrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeletePrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       o.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeletePrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

type DeleteResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete the restrict workspace admins setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRestrictWorkspaceAdminsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) {
}

func (newState *DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) {
}

func (c DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteRestrictWorkspaceAdminsSettingResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Delete a token
type DeleteTokenManagementRequest_SdkV2 struct {
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
func (a DeleteTokenManagementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteTokenManagementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteTokenManagementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteTokenManagementRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *DisableLegacyAccess_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyAccess_SdkV2) {
}

func (newState *DisableLegacyAccess_SdkV2) SyncEffectiveFieldsDuringRead(existingState DisableLegacyAccess_SdkV2) {
}

func (c DisableLegacyAccess_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DisableLegacyAccess_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_access": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyAccess_SdkV2
// only implements ToObjectValue() and Type().
func (o DisableLegacyAccess_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_access": o.DisableLegacyAccess,
			"etag":                  o.Etag,
			"setting_name":          o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableLegacyAccess_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DisableLegacyAccess_SdkV2) GetDisableLegacyAccess(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.DisableLegacyAccess.IsNull() || o.DisableLegacyAccess.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.DisableLegacyAccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDisableLegacyAccess sets the value of the DisableLegacyAccess field in DisableLegacyAccess_SdkV2.
func (o *DisableLegacyAccess_SdkV2) SetDisableLegacyAccess(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disable_legacy_access"]
	o.DisableLegacyAccess = types.ListValueMust(t, vs)
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

func (newState *DisableLegacyDbfs_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyDbfs_SdkV2) {
}

func (newState *DisableLegacyDbfs_SdkV2) SyncEffectiveFieldsDuringRead(existingState DisableLegacyDbfs_SdkV2) {
}

func (c DisableLegacyDbfs_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DisableLegacyDbfs_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_dbfs": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyDbfs_SdkV2
// only implements ToObjectValue() and Type().
func (o DisableLegacyDbfs_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_dbfs": o.DisableLegacyDbfs,
			"etag":                o.Etag,
			"setting_name":        o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableLegacyDbfs_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DisableLegacyDbfs_SdkV2) GetDisableLegacyDbfs(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.DisableLegacyDbfs.IsNull() || o.DisableLegacyDbfs.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.DisableLegacyDbfs.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDisableLegacyDbfs sets the value of the DisableLegacyDbfs field in DisableLegacyDbfs_SdkV2.
func (o *DisableLegacyDbfs_SdkV2) SetDisableLegacyDbfs(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disable_legacy_dbfs"]
	o.DisableLegacyDbfs = types.ListValueMust(t, vs)
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

func (newState *DisableLegacyFeatures_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DisableLegacyFeatures_SdkV2) {
}

func (newState *DisableLegacyFeatures_SdkV2) SyncEffectiveFieldsDuringRead(existingState DisableLegacyFeatures_SdkV2) {
}

func (c DisableLegacyFeatures_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DisableLegacyFeatures_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"disable_legacy_features": reflect.TypeOf(BooleanMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DisableLegacyFeatures_SdkV2
// only implements ToObjectValue() and Type().
func (o DisableLegacyFeatures_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"disable_legacy_features": o.DisableLegacyFeatures,
			"etag":                    o.Etag,
			"setting_name":            o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DisableLegacyFeatures_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *DisableLegacyFeatures_SdkV2) GetDisableLegacyFeatures(ctx context.Context) (BooleanMessage_SdkV2, bool) {
	var e BooleanMessage_SdkV2
	if o.DisableLegacyFeatures.IsNull() || o.DisableLegacyFeatures.IsUnknown() {
		return e, false
	}
	var v []BooleanMessage_SdkV2
	d := o.DisableLegacyFeatures.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDisableLegacyFeatures sets the value of the DisableLegacyFeatures field in DisableLegacyFeatures_SdkV2.
func (o *DisableLegacyFeatures_SdkV2) SetDisableLegacyFeatures(ctx context.Context, v BooleanMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["disable_legacy_features"]
	o.DisableLegacyFeatures = types.ListValueMust(t, vs)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy_SdkV2 struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess types.List `tfsdk:"internet_access"`
}

func (newState *EgressNetworkPolicy_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicy_SdkV2) {
}

func (newState *EgressNetworkPolicy_SdkV2) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicy_SdkV2) {
}

func (c EgressNetworkPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"internet_access": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicy_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"internet_access": o.InternetAccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicy_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicy_SdkV2) GetInternetAccess(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicy_SdkV2, bool) {
	var e EgressNetworkPolicyInternetAccessPolicy_SdkV2
	if o.InternetAccess.IsNull() || o.InternetAccess.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicy_SdkV2
	d := o.InternetAccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInternetAccess sets the value of the InternetAccess field in EgressNetworkPolicy_SdkV2.
func (o *EgressNetworkPolicy_SdkV2) SetInternetAccess(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicy_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["internet_access"]
	o.InternetAccess = types.ListValueMust(t, vs)
}

type EgressNetworkPolicyInternetAccessPolicy_SdkV2 struct {
	AllowedInternetDestinations types.List `tfsdk:"allowed_internet_destinations"`

	AllowedStorageDestinations types.List `tfsdk:"allowed_storage_destinations"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode types.List `tfsdk:"log_only_mode"`
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	RestrictionMode types.String `tfsdk:"restriction_mode"`
}

func (newState *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicy_SdkV2) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicy_SdkV2) {
}

func (c EgressNetworkPolicyInternetAccessPolicy_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_internet_destinations": reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2{}),
		"allowed_storage_destinations":  reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2{}),
		"log_only_mode":                 reflect.TypeOf(EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicy_SdkV2
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicy_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EgressNetworkPolicyInternetAccessPolicy_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetAllowedInternetDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2, bool) {
	if o.AllowedInternetDestinations.IsNull() || o.AllowedInternetDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2
	d := o.AllowedInternetDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedInternetDestinations sets the value of the AllowedInternetDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2.
func (o *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SetAllowedInternetDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_internet_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedInternetDestinations = types.ListValueMust(t, vs)
}

// GetAllowedStorageDestinations returns the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2 as
// a slice of EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetAllowedStorageDestinations(ctx context.Context) ([]EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2, bool) {
	if o.AllowedStorageDestinations.IsNull() || o.AllowedStorageDestinations.IsUnknown() {
		return nil, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2
	d := o.AllowedStorageDestinations.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllowedStorageDestinations sets the value of the AllowedStorageDestinations field in EgressNetworkPolicyInternetAccessPolicy_SdkV2.
func (o *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SetAllowedStorageDestinations(ctx context.Context, v []EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_storage_destinations"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedStorageDestinations = types.ListValueMust(t, vs)
}

// GetLogOnlyMode returns the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy_SdkV2 as
// a EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EgressNetworkPolicyInternetAccessPolicy_SdkV2) GetLogOnlyMode(ctx context.Context) (EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2, bool) {
	var e EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2
	if o.LogOnlyMode.IsNull() || o.LogOnlyMode.IsUnknown() {
		return e, false
	}
	var v []EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2
	d := o.LogOnlyMode.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetLogOnlyMode sets the value of the LogOnlyMode field in EgressNetworkPolicyInternetAccessPolicy_SdkV2.
func (o *EgressNetworkPolicyInternetAccessPolicy_SdkV2) SetLogOnlyMode(ctx context.Context, v EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["log_only_mode"]
	o.LogOnlyMode = types.ListValueMust(t, vs)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2 struct {
	Destination types.String `tfsdk:"destination"`
	// The filtering protocol used by the DP. For private and public preview,
	// SEG will only support TCP filtering (i.e. DNS based filtering, filtering
	// by destination IP address), so protocol will be set to TCP by default and
	// hidden from the user. In the future, users may be able to select HTTP
	// filtering (i.e. SNI based filtering, filtering by FQDN).
	Protocol types.String `tfsdk:"protocol"`

	Type_ types.String `tfsdk:"type"`
}

func (newState *EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) {
}

func (c EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination": o.Destination,
			"protocol":    o.Protocol,
			"type":        o.Type_,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyInternetAccessPolicyInternetDestination_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) {
}

func (c EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workloads": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"log_only_mode_type": o.LogOnlyModeType,
			"workloads":          o.Workloads,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) GetWorkloads(ctx context.Context) ([]types.String, bool) {
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

// SetWorkloads sets the value of the Workloads field in EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2.
func (o *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode_SdkV2) SetWorkloads(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["workloads"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Workloads = types.ListValueMust(t, vs)
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

func (newState *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) {
}

func (newState *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) SyncEffectiveFieldsDuringRead(existingState EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) {
}

func (c EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"allowed_paths": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2
// only implements ToObjectValue() and Type().
func (o EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) GetAllowedPaths(ctx context.Context) ([]types.String, bool) {
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

// SetAllowedPaths sets the value of the AllowedPaths field in EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2.
func (o *EgressNetworkPolicyInternetAccessPolicyStorageDestination_SdkV2) SetAllowedPaths(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["allowed_paths"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllowedPaths = types.ListValueMust(t, vs)
}

type EmailConfig_SdkV2 struct {
	// Email addresses to notify.
	Addresses types.List `tfsdk:"addresses"`
}

func (newState *EmailConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmailConfig_SdkV2) {
}

func (newState *EmailConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState EmailConfig_SdkV2) {
}

func (c EmailConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EmailConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmailConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o EmailConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"addresses": o.Addresses,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmailConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EmailConfig_SdkV2) GetAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetAddresses sets the value of the Addresses field in EmailConfig_SdkV2.
func (o *EmailConfig_SdkV2) SetAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Addresses = types.ListValueMust(t, vs)
}

type Empty_SdkV2 struct {
}

func (newState *Empty_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Empty_SdkV2) {
}

func (newState *Empty_SdkV2) SyncEffectiveFieldsDuringRead(existingState Empty_SdkV2) {
}

func (c Empty_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Empty.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Empty_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Empty_SdkV2
// only implements ToObjectValue() and Type().
func (o Empty_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o Empty_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring_SdkV2 struct {
	IsEnabled types.Bool `tfsdk:"is_enabled"`
}

func (newState *EnhancedSecurityMonitoring_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnhancedSecurityMonitoring_SdkV2) {
}

func (newState *EnhancedSecurityMonitoring_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnhancedSecurityMonitoring_SdkV2) {
}

func (c EnhancedSecurityMonitoring_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnhancedSecurityMonitoring_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoring_SdkV2
// only implements ToObjectValue() and Type().
func (o EnhancedSecurityMonitoring_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enabled": o.IsEnabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnhancedSecurityMonitoring_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enabled": types.BoolType,
		},
	}
}

type EnhancedSecurityMonitoringSetting_SdkV2 struct {
	// SHIELD feature: ESM
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

func (newState *EnhancedSecurityMonitoringSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnhancedSecurityMonitoringSetting_SdkV2) {
}

func (newState *EnhancedSecurityMonitoringSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState EnhancedSecurityMonitoringSetting_SdkV2) {
}

func (c EnhancedSecurityMonitoringSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EnhancedSecurityMonitoringSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"enhanced_security_monitoring_workspace": reflect.TypeOf(EnhancedSecurityMonitoring_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EnhancedSecurityMonitoringSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o EnhancedSecurityMonitoringSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enhanced_security_monitoring_workspace": o.EnhancedSecurityMonitoringWorkspace,
			"etag":                                   o.Etag,
			"setting_name":                           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EnhancedSecurityMonitoringSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EnhancedSecurityMonitoringSetting_SdkV2) GetEnhancedSecurityMonitoringWorkspace(ctx context.Context) (EnhancedSecurityMonitoring_SdkV2, bool) {
	var e EnhancedSecurityMonitoring_SdkV2
	if o.EnhancedSecurityMonitoringWorkspace.IsNull() || o.EnhancedSecurityMonitoringWorkspace.IsUnknown() {
		return e, false
	}
	var v []EnhancedSecurityMonitoring_SdkV2
	d := o.EnhancedSecurityMonitoringWorkspace.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEnhancedSecurityMonitoringWorkspace sets the value of the EnhancedSecurityMonitoringWorkspace field in EnhancedSecurityMonitoringSetting_SdkV2.
func (o *EnhancedSecurityMonitoringSetting_SdkV2) SetEnhancedSecurityMonitoringWorkspace(ctx context.Context, v EnhancedSecurityMonitoring_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["enhanced_security_monitoring_workspace"]
	o.EnhancedSecurityMonitoringWorkspace = types.ListValueMust(t, vs)
}

// Account level policy for ESM
type EsmEnablementAccount_SdkV2 struct {
	IsEnforced types.Bool `tfsdk:"is_enforced"`
}

func (newState *EsmEnablementAccount_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EsmEnablementAccount_SdkV2) {
}

func (newState *EsmEnablementAccount_SdkV2) SyncEffectiveFieldsDuringRead(existingState EsmEnablementAccount_SdkV2) {
}

func (c EsmEnablementAccount_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EsmEnablementAccount_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccount_SdkV2
// only implements ToObjectValue() and Type().
func (o EsmEnablementAccount_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"is_enforced": o.IsEnforced,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EsmEnablementAccount_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_enforced": types.BoolType,
		},
	}
}

type EsmEnablementAccountSetting_SdkV2 struct {
	// Account level policy for ESM
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

func (newState *EsmEnablementAccountSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EsmEnablementAccountSetting_SdkV2) {
}

func (newState *EsmEnablementAccountSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState EsmEnablementAccountSetting_SdkV2) {
}

func (c EsmEnablementAccountSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EsmEnablementAccountSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"esm_enablement_account": reflect.TypeOf(EsmEnablementAccount_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EsmEnablementAccountSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o EsmEnablementAccountSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"esm_enablement_account": o.EsmEnablementAccount,
			"etag":                   o.Etag,
			"setting_name":           o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EsmEnablementAccountSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *EsmEnablementAccountSetting_SdkV2) GetEsmEnablementAccount(ctx context.Context) (EsmEnablementAccount_SdkV2, bool) {
	var e EsmEnablementAccount_SdkV2
	if o.EsmEnablementAccount.IsNull() || o.EsmEnablementAccount.IsUnknown() {
		return e, false
	}
	var v []EsmEnablementAccount_SdkV2
	d := o.EsmEnablementAccount.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEsmEnablementAccount sets the value of the EsmEnablementAccount field in EsmEnablementAccountSetting_SdkV2.
func (o *EsmEnablementAccountSetting_SdkV2) SetEsmEnablementAccount(ctx context.Context, v EsmEnablementAccount_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["esm_enablement_account"]
	o.EsmEnablementAccount = types.ListValueMust(t, vs)
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken_SdkV2 struct {
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

func (newState *ExchangeToken_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeToken_SdkV2) {
}

func (newState *ExchangeToken_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExchangeToken_SdkV2) {
}

func (c ExchangeToken_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExchangeToken_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"scopes": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeToken_SdkV2
// only implements ToObjectValue() and Type().
func (o ExchangeToken_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ExchangeToken_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetScopes returns the value of the Scopes field in ExchangeToken_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeToken_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeToken_SdkV2.
func (o *ExchangeToken_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

// Exchange a token with the IdP
type ExchangeTokenRequest_SdkV2 struct {
	// The partition of Credentials store
	PartitionId types.List `tfsdk:"partitionId"`
	// Array of scopes for the token request.
	Scopes types.List `tfsdk:"scopes"`
	// A list of token types being requested
	TokenType types.List `tfsdk:"tokenType"`
}

func (newState *ExchangeTokenRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeTokenRequest_SdkV2) {
}

func (newState *ExchangeTokenRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExchangeTokenRequest_SdkV2) {
}

func (c ExchangeTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["partitionId"] = attrs["partitionId"].SetRequired()
	attrs["partitionId"] = attrs["partitionId"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
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
func (a ExchangeTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"partitionId": reflect.TypeOf(PartitionId_SdkV2{}),
		"scopes":      reflect.TypeOf(types.String{}),
		"tokenType":   reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExchangeTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"partitionId": o.PartitionId,
			"scopes":      o.Scopes,
			"tokenType":   o.TokenType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"partitionId": basetypes.ListType{
				ElemType: PartitionId_SdkV2{}.Type(ctx),
			},
			"scopes": basetypes.ListType{
				ElemType: types.StringType,
			},
			"tokenType": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetPartitionId returns the value of the PartitionId field in ExchangeTokenRequest_SdkV2 as
// a PartitionId_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeTokenRequest_SdkV2) GetPartitionId(ctx context.Context) (PartitionId_SdkV2, bool) {
	var e PartitionId_SdkV2
	if o.PartitionId.IsNull() || o.PartitionId.IsUnknown() {
		return e, false
	}
	var v []PartitionId_SdkV2
	d := o.PartitionId.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPartitionId sets the value of the PartitionId field in ExchangeTokenRequest_SdkV2.
func (o *ExchangeTokenRequest_SdkV2) SetPartitionId(ctx context.Context, v PartitionId_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["partitionId"]
	o.PartitionId = types.ListValueMust(t, vs)
}

// GetScopes returns the value of the Scopes field in ExchangeTokenRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeTokenRequest_SdkV2) GetScopes(ctx context.Context) ([]types.String, bool) {
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

// SetScopes sets the value of the Scopes field in ExchangeTokenRequest_SdkV2.
func (o *ExchangeTokenRequest_SdkV2) SetScopes(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["scopes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Scopes = types.ListValueMust(t, vs)
}

// GetTokenType returns the value of the TokenType field in ExchangeTokenRequest_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ExchangeTokenRequest_SdkV2) GetTokenType(ctx context.Context) ([]types.String, bool) {
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

// SetTokenType sets the value of the TokenType field in ExchangeTokenRequest_SdkV2.
func (o *ExchangeTokenRequest_SdkV2) SetTokenType(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tokenType"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TokenType = types.ListValueMust(t, vs)
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse_SdkV2 struct {
	Values types.List `tfsdk:"values"`
}

func (newState *ExchangeTokenResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExchangeTokenResponse_SdkV2) {
}

func (newState *ExchangeTokenResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExchangeTokenResponse_SdkV2) {
}

func (c ExchangeTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExchangeTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(ExchangeToken_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExchangeTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ExchangeTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"values": o.Values,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExchangeTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ExchangeTokenResponse_SdkV2) GetValues(ctx context.Context) ([]ExchangeToken_SdkV2, bool) {
	if o.Values.IsNull() || o.Values.IsUnknown() {
		return nil, false
	}
	var v []ExchangeToken_SdkV2
	d := o.Values.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValues sets the value of the Values field in ExchangeTokenResponse_SdkV2.
func (o *ExchangeTokenResponse_SdkV2) SetValues(ctx context.Context, v []ExchangeToken_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["values"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Values = types.ListValueMust(t, vs)
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse_SdkV2 struct {
	// Definition of an IP Access list
	IpAccessList types.List `tfsdk:"ip_access_list"`
}

func (newState *FetchIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FetchIpAccessListResponse_SdkV2) {
}

func (newState *FetchIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState FetchIpAccessListResponse_SdkV2) {
}

func (c FetchIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FetchIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FetchIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o FetchIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": o.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FetchIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *FetchIpAccessListResponse_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessListInfo_SdkV2, bool) {
	var e IpAccessListInfo_SdkV2
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo_SdkV2
	d := o.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in FetchIpAccessListResponse_SdkV2.
func (o *FetchIpAccessListResponse_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessListInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	o.IpAccessList = types.ListValueMust(t, vs)
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

func (newState *GenericWebhookConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenericWebhookConfig_SdkV2) {
}

func (newState *GenericWebhookConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState GenericWebhookConfig_SdkV2) {
}

func (c GenericWebhookConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GenericWebhookConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GenericWebhookConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GenericWebhookConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GenericWebhookConfig_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAccountIpAccessEnableRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAccountIpAccessEnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessEnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountIpAccessEnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountIpAccessEnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get IP access list
type GetAccountIpAccessListRequest_SdkV2 struct {
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
func (a GetAccountIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAccountIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAccountIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAccountIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

// Retrieve the AI/BI dashboard embedding access policy
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingAccessPolicySettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Retrieve the list of domains approved to host embedded AI/BI dashboards
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAibiDashboardEmbeddingApprovedDomainsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the automatic cluster update setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetAutomaticClusterUpdateSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetAutomaticClusterUpdateSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetAutomaticClusterUpdateSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetAutomaticClusterUpdateSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetAutomaticClusterUpdateSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the compliance security profile setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetComplianceSecurityProfileSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetComplianceSecurityProfileSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetComplianceSecurityProfileSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetComplianceSecurityProfileSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetComplianceSecurityProfileSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the compliance security profile setting for new workspaces
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetCspEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetCspEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetCspEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetCspEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetCspEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the default namespace setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDefaultNamespaceSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDefaultNamespaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDefaultNamespaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDefaultNamespaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDefaultNamespaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Retrieve Legacy Access Disablement Status
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyAccessRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDisableLegacyAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDisableLegacyAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the disable legacy DBFS setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyDbfsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDisableLegacyDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDisableLegacyDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the disable legacy features setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDisableLegacyFeaturesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDisableLegacyFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetDisableLegacyFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetDisableLegacyFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetDisableLegacyFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the enhanced security monitoring setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEnhancedSecurityMonitoringSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEnhancedSecurityMonitoringSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEnhancedSecurityMonitoringSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEnhancedSecurityMonitoringSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEnhancedSecurityMonitoringSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get the enhanced security monitoring setting for new workspaces
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetEsmEnablementAccountSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetEsmEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetEsmEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetEsmEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetEsmEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get access list
type GetIpAccessListRequest_SdkV2 struct {
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
func (a GetIpAccessListRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetIpAccessListRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list_id": o.IpAccessListId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIpAccessListRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ip_access_list_id": types.StringType,
		},
	}
}

type GetIpAccessListResponse_SdkV2 struct {
	// Definition of an IP Access list
	IpAccessList types.List `tfsdk:"ip_access_list"`
}

func (newState *GetIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListResponse_SdkV2) {
}

func (newState *GetIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListResponse_SdkV2) {
}

func (c GetIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_list": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_list": o.IpAccessList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetIpAccessListResponse_SdkV2) GetIpAccessList(ctx context.Context) (IpAccessListInfo_SdkV2, bool) {
	var e IpAccessListInfo_SdkV2
	if o.IpAccessList.IsNull() || o.IpAccessList.IsUnknown() {
		return e, false
	}
	var v []IpAccessListInfo_SdkV2
	d := o.IpAccessList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetIpAccessList sets the value of the IpAccessList field in GetIpAccessListResponse_SdkV2.
func (o *GetIpAccessListResponse_SdkV2) SetIpAccessList(ctx context.Context, v IpAccessListInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_list"]
	o.IpAccessList = types.ListValueMust(t, vs)
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse_SdkV2 struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (newState *GetIpAccessListsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetIpAccessListsResponse_SdkV2) {
}

func (newState *GetIpAccessListsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetIpAccessListsResponse_SdkV2) {
}

func (c GetIpAccessListsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetIpAccessListsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetIpAccessListsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetIpAccessListsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": o.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetIpAccessListsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetIpAccessListsResponse_SdkV2) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo_SdkV2, bool) {
	if o.IpAccessLists.IsNull() || o.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo_SdkV2
	d := o.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in GetIpAccessListsResponse_SdkV2.
func (o *GetIpAccessListsResponse_SdkV2) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAccessLists = types.ListValueMust(t, vs)
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest_SdkV2 struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNetworkConnectivityConfigurationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNetworkConnectivityConfigurationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNetworkConnectivityConfigurationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetNetworkConnectivityConfigurationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNetworkConnectivityConfigurationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
		},
	}
}

// Get a notification destination
type GetNotificationDestinationRequest_SdkV2 struct {
	Id types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetNotificationDestinationRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"id": o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id": types.StringType,
		},
	}
}

// Get Personal Compute setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPersonalComputeSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPersonalComputeSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPersonalComputeSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPersonalComputeSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPersonalComputeSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest_SdkV2 struct {
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
func (a GetPrivateEndpointRuleRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetPrivateEndpointRuleRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetPrivateEndpointRuleRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"private_endpoint_rule_id":       o.PrivateEndpointRuleId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetPrivateEndpointRuleRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"private_endpoint_rule_id":       types.StringType,
		},
	}
}

// Get the restrict workspace admins setting
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetRestrictWorkspaceAdminsSettingRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetRestrictWorkspaceAdminsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetRestrictWorkspaceAdminsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetRestrictWorkspaceAdminsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag": o.Etag,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetRestrictWorkspaceAdminsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"etag": types.StringType,
		},
	}
}

// Check configuration status
type GetStatusRequest_SdkV2 struct {
	Keys types.String `tfsdk:"-"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetStatusRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetStatusRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetStatusRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetStatusRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"keys": o.Keys,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetStatusRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"keys": types.StringType,
		},
	}
}

// Get token info
type GetTokenManagementRequest_SdkV2 struct {
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
func (a GetTokenManagementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenManagementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetTokenManagementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTokenManagementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type GetTokenPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (newState *GetTokenPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenPermissionLevelsResponse_SdkV2) {
}

func (newState *GetTokenPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetTokenPermissionLevelsResponse_SdkV2) {
}

func (c GetTokenPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetTokenPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(TokenPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetTokenPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTokenPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetTokenPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]TokenPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []TokenPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetTokenPermissionLevelsResponse_SdkV2.
func (o *GetTokenPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []TokenPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse_SdkV2 struct {
	TokenInfo types.List `tfsdk:"token_info"`
}

func (newState *GetTokenResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetTokenResponse_SdkV2) {
}

func (newState *GetTokenResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetTokenResponse_SdkV2) {
}

func (c GetTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_info": reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_info": o.TokenInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *GetTokenResponse_SdkV2) GetTokenInfo(ctx context.Context) (TokenInfo_SdkV2, bool) {
	var e TokenInfo_SdkV2
	if o.TokenInfo.IsNull() || o.TokenInfo.IsUnknown() {
		return e, false
	}
	var v []TokenInfo_SdkV2
	d := o.TokenInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTokenInfo sets the value of the TokenInfo field in GetTokenResponse_SdkV2.
func (o *GetTokenResponse_SdkV2) SetTokenInfo(ctx context.Context, v TokenInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_info"]
	o.TokenInfo = types.ListValueMust(t, vs)
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

func (newState *IpAccessListInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan IpAccessListInfo_SdkV2) {
}

func (newState *IpAccessListInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState IpAccessListInfo_SdkV2) {
}

func (c IpAccessListInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a IpAccessListInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, IpAccessListInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o IpAccessListInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o IpAccessListInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *IpAccessListInfo_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in IpAccessListInfo_SdkV2.
func (o *IpAccessListInfo_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse_SdkV2 struct {
	IpAccessLists types.List `tfsdk:"ip_access_lists"`
}

func (newState *ListIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListIpAccessListResponse_SdkV2) {
}

func (newState *ListIpAccessListResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListIpAccessListResponse_SdkV2) {
}

func (c ListIpAccessListResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListIpAccessListResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_access_lists": reflect.TypeOf(IpAccessListInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListIpAccessListResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListIpAccessListResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ip_access_lists": o.IpAccessLists,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListIpAccessListResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListIpAccessListResponse_SdkV2) GetIpAccessLists(ctx context.Context) ([]IpAccessListInfo_SdkV2, bool) {
	if o.IpAccessLists.IsNull() || o.IpAccessLists.IsUnknown() {
		return nil, false
	}
	var v []IpAccessListInfo_SdkV2
	d := o.IpAccessLists.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIpAccessLists sets the value of the IpAccessLists field in ListIpAccessListResponse_SdkV2.
func (o *ListIpAccessListResponse_SdkV2) SetIpAccessLists(ctx context.Context, v []IpAccessListInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_access_lists"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAccessLists = types.ListValueMust(t, vs)
}

type ListNccAzurePrivateEndpointRulesResponse_SdkV2 struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListNccAzurePrivateEndpointRulesResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNccAzurePrivateEndpointRulesResponse_SdkV2) {
}

func (newState *ListNccAzurePrivateEndpointRulesResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListNccAzurePrivateEndpointRulesResponse_SdkV2) {
}

func (c ListNccAzurePrivateEndpointRulesResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["items"] = attrs["items"].SetOptional()
	attrs["next_page_token"] = attrs["next_page_token"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListNccAzurePrivateEndpointRulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListNccAzurePrivateEndpointRulesResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NccAzurePrivateEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNccAzurePrivateEndpointRulesResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNccAzurePrivateEndpointRulesResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           o.Items,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNccAzurePrivateEndpointRulesResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"items": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule_SdkV2{}.Type(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// GetItems returns the value of the Items field in ListNccAzurePrivateEndpointRulesResponse_SdkV2 as
// a slice of NccAzurePrivateEndpointRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListNccAzurePrivateEndpointRulesResponse_SdkV2) GetItems(ctx context.Context) ([]NccAzurePrivateEndpointRule_SdkV2, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []NccAzurePrivateEndpointRule_SdkV2
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNccAzurePrivateEndpointRulesResponse_SdkV2.
func (o *ListNccAzurePrivateEndpointRulesResponse_SdkV2) SetItems(ctx context.Context, v []NccAzurePrivateEndpointRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest_SdkV2 struct {
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
func (a ListNetworkConnectivityConfigurationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNetworkConnectivityConfigurationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworkConnectivityConfigurationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_token": types.StringType,
		},
	}
}

type ListNetworkConnectivityConfigurationsResponse_SdkV2 struct {
	Items types.List `tfsdk:"items"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

func (newState *ListNetworkConnectivityConfigurationsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNetworkConnectivityConfigurationsResponse_SdkV2) {
}

func (newState *ListNetworkConnectivityConfigurationsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListNetworkConnectivityConfigurationsResponse_SdkV2) {
}

func (c ListNetworkConnectivityConfigurationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNetworkConnectivityConfigurationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"items": reflect.TypeOf(NetworkConnectivityConfiguration_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNetworkConnectivityConfigurationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNetworkConnectivityConfigurationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"items":           o.Items,
			"next_page_token": o.NextPageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNetworkConnectivityConfigurationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListNetworkConnectivityConfigurationsResponse_SdkV2) GetItems(ctx context.Context) ([]NetworkConnectivityConfiguration_SdkV2, bool) {
	if o.Items.IsNull() || o.Items.IsUnknown() {
		return nil, false
	}
	var v []NetworkConnectivityConfiguration_SdkV2
	d := o.Items.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetItems sets the value of the Items field in ListNetworkConnectivityConfigurationsResponse_SdkV2.
func (o *ListNetworkConnectivityConfigurationsResponse_SdkV2) SetItems(ctx context.Context, v []NetworkConnectivityConfiguration_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["items"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Items = types.ListValueMust(t, vs)
}

// List notification destinations
type ListNotificationDestinationsRequest_SdkV2 struct {
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
func (a ListNotificationDestinationsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNotificationDestinationsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"page_size":  o.PageSize,
			"page_token": o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNotificationDestinationsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *ListNotificationDestinationsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsResponse_SdkV2) {
}

func (newState *ListNotificationDestinationsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsResponse_SdkV2) {
}

func (c ListNotificationDestinationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNotificationDestinationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"results": reflect.TypeOf(ListNotificationDestinationsResult_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNotificationDestinationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"next_page_token": o.NextPageToken,
			"results":         o.Results,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNotificationDestinationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListNotificationDestinationsResponse_SdkV2) GetResults(ctx context.Context) ([]ListNotificationDestinationsResult_SdkV2, bool) {
	if o.Results.IsNull() || o.Results.IsUnknown() {
		return nil, false
	}
	var v []ListNotificationDestinationsResult_SdkV2
	d := o.Results.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetResults sets the value of the Results field in ListNotificationDestinationsResponse_SdkV2.
func (o *ListNotificationDestinationsResponse_SdkV2) SetResults(ctx context.Context, v []ListNotificationDestinationsResult_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["results"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Results = types.ListValueMust(t, vs)
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

func (newState *ListNotificationDestinationsResult_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListNotificationDestinationsResult_SdkV2) {
}

func (newState *ListNotificationDestinationsResult_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListNotificationDestinationsResult_SdkV2) {
}

func (c ListNotificationDestinationsResult_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListNotificationDestinationsResult_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListNotificationDestinationsResult_SdkV2
// only implements ToObjectValue() and Type().
func (o ListNotificationDestinationsResult_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"destination_type": o.DestinationType,
			"display_name":     o.DisplayName,
			"id":               o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListNotificationDestinationsResult_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_type": types.StringType,
			"display_name":     types.StringType,
			"id":               types.StringType,
		},
	}
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest_SdkV2 struct {
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
func (a ListPrivateEndpointRulesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPrivateEndpointRulesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPrivateEndpointRulesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"page_token":                     o.PageToken,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPrivateEndpointRulesRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"network_connectivity_config_id": types.StringType,
			"page_token":                     types.StringType,
		},
	}
}

type ListPublicTokensResponse_SdkV2 struct {
	// The information for each token.
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (newState *ListPublicTokensResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListPublicTokensResponse_SdkV2) {
}

func (newState *ListPublicTokensResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListPublicTokensResponse_SdkV2) {
}

func (c ListPublicTokensResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListPublicTokensResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(PublicTokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListPublicTokensResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListPublicTokensResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": o.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListPublicTokensResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListPublicTokensResponse_SdkV2) GetTokenInfos(ctx context.Context) ([]PublicTokenInfo_SdkV2, bool) {
	if o.TokenInfos.IsNull() || o.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []PublicTokenInfo_SdkV2
	d := o.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListPublicTokensResponse_SdkV2.
func (o *ListPublicTokensResponse_SdkV2) SetTokenInfos(ctx context.Context, v []PublicTokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TokenInfos = types.ListValueMust(t, vs)
}

// List all tokens
type ListTokenManagementRequest_SdkV2 struct {
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
func (a ListTokenManagementRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokenManagementRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTokenManagementRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"created_by_id":       o.CreatedById,
			"created_by_username": o.CreatedByUsername,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTokenManagementRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_by_id":       types.Int64Type,
			"created_by_username": types.StringType,
		},
	}
}

// Tokens were successfully returned.
type ListTokensResponse_SdkV2 struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos types.List `tfsdk:"token_infos"`
}

func (newState *ListTokensResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListTokensResponse_SdkV2) {
}

func (newState *ListTokensResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListTokensResponse_SdkV2) {
}

func (c ListTokensResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListTokensResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"token_infos": reflect.TypeOf(TokenInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListTokensResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListTokensResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_infos": o.TokenInfos,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListTokensResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ListTokensResponse_SdkV2) GetTokenInfos(ctx context.Context) ([]TokenInfo_SdkV2, bool) {
	if o.TokenInfos.IsNull() || o.TokenInfos.IsUnknown() {
		return nil, false
	}
	var v []TokenInfo_SdkV2
	d := o.TokenInfos.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTokenInfos sets the value of the TokenInfos field in ListTokensResponse_SdkV2.
func (o *ListTokensResponse_SdkV2) SetTokenInfos(ctx context.Context, v []TokenInfo_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["token_infos"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TokenInfos = types.ListValueMust(t, vs)
}

type MicrosoftTeamsConfig_SdkV2 struct {
	// [Input-Only] URL for Microsoft Teams.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
}

func (newState *MicrosoftTeamsConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan MicrosoftTeamsConfig_SdkV2) {
}

func (newState *MicrosoftTeamsConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState MicrosoftTeamsConfig_SdkV2) {
}

func (c MicrosoftTeamsConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a MicrosoftTeamsConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, MicrosoftTeamsConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o MicrosoftTeamsConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"url":     o.Url,
			"url_set": o.UrlSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o MicrosoftTeamsConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"url":     types.StringType,
			"url_set": types.BoolType,
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

func (newState *NccAwsStableIpRule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAwsStableIpRule_SdkV2) {
}

func (newState *NccAwsStableIpRule_SdkV2) SyncEffectiveFieldsDuringRead(existingState NccAwsStableIpRule_SdkV2) {
}

func (c NccAwsStableIpRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccAwsStableIpRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cidr_blocks": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAwsStableIpRule_SdkV2
// only implements ToObjectValue() and Type().
func (o NccAwsStableIpRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cidr_blocks": o.CidrBlocks,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccAwsStableIpRule_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *NccAwsStableIpRule_SdkV2) GetCidrBlocks(ctx context.Context) ([]types.String, bool) {
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

// SetCidrBlocks sets the value of the CidrBlocks field in NccAwsStableIpRule_SdkV2.
func (o *NccAwsStableIpRule_SdkV2) SetCidrBlocks(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cidr_blocks"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.CidrBlocks = types.ListValueMust(t, vs)
}

type NccAzurePrivateEndpointRule_SdkV2 struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is `ESTABLISHED`. Remember
	// that you must approve new endpoints on your resources in the Azure portal
	// before they take effect.
	//
	// The possible values are: - INIT: (deprecated) The endpoint has been
	// created and pending approval. - PENDING: The endpoint has been created
	// and pending approval. - ESTABLISHED: The endpoint has been approved and
	// is ready to use in your serverless compute resources. - REJECTED:
	// Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up.
	ConnectionState types.String `tfsdk:"connection_state"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// Whether this private endpoint is deactivated.
	Deactivated types.Bool `tfsdk:"deactivated"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt types.Int64 `tfsdk:"deactivated_at"`
	// The name of the Azure private endpoint resource.
	EndpointName types.String `tfsdk:"endpoint_name"`
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
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

func (newState *NccAzurePrivateEndpointRule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAzurePrivateEndpointRule_SdkV2) {
}

func (newState *NccAzurePrivateEndpointRule_SdkV2) SyncEffectiveFieldsDuringRead(existingState NccAzurePrivateEndpointRule_SdkV2) {
}

func (c NccAzurePrivateEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_state"] = attrs["connection_state"].SetComputed()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["deactivated"] = attrs["deactivated"].SetComputed()
	attrs["deactivated_at"] = attrs["deactivated_at"].SetComputed()
	attrs["endpoint_name"] = attrs["endpoint_name"].SetComputed()
	attrs["group_id"] = attrs["group_id"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetOptional()
	attrs["resource_id"] = attrs["resource_id"].SetOptional()
	attrs["rule_id"] = attrs["rule_id"].SetComputed()
	attrs["updated_time"] = attrs["updated_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NccAzurePrivateEndpointRule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NccAzurePrivateEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzurePrivateEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (o NccAzurePrivateEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_state":               o.ConnectionState,
			"creation_time":                  o.CreationTime,
			"deactivated":                    o.Deactivated,
			"deactivated_at":                 o.DeactivatedAt,
			"endpoint_name":                  o.EndpointName,
			"group_id":                       o.GroupId,
			"network_connectivity_config_id": o.NetworkConnectivityConfigId,
			"resource_id":                    o.ResourceId,
			"rule_id":                        o.RuleId,
			"updated_time":                   o.UpdatedTime,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccAzurePrivateEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_state":               types.StringType,
			"creation_time":                  types.Int64Type,
			"deactivated":                    types.BoolType,
			"deactivated_at":                 types.Int64Type,
			"endpoint_name":                  types.StringType,
			"group_id":                       types.StringType,
			"network_connectivity_config_id": types.StringType,
			"resource_id":                    types.StringType,
			"rule_id":                        types.StringType,
			"updated_time":                   types.Int64Type,
		},
	}
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule_SdkV2 struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets types.List `tfsdk:"subnets"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion types.String `tfsdk:"target_region"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices types.List `tfsdk:"target_services"`
}

func (newState *NccAzureServiceEndpointRule_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccAzureServiceEndpointRule_SdkV2) {
}

func (newState *NccAzureServiceEndpointRule_SdkV2) SyncEffectiveFieldsDuringRead(existingState NccAzureServiceEndpointRule_SdkV2) {
}

func (c NccAzureServiceEndpointRule_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccAzureServiceEndpointRule_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subnets":         reflect.TypeOf(types.String{}),
		"target_services": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccAzureServiceEndpointRule_SdkV2
// only implements ToObjectValue() and Type().
func (o NccAzureServiceEndpointRule_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"subnets":         o.Subnets,
			"target_region":   o.TargetRegion,
			"target_services": o.TargetServices,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccAzureServiceEndpointRule_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *NccAzureServiceEndpointRule_SdkV2) GetSubnets(ctx context.Context) ([]types.String, bool) {
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

// SetSubnets sets the value of the Subnets field in NccAzureServiceEndpointRule_SdkV2.
func (o *NccAzureServiceEndpointRule_SdkV2) SetSubnets(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["subnets"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Subnets = types.ListValueMust(t, vs)
}

// GetTargetServices returns the value of the TargetServices field in NccAzureServiceEndpointRule_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *NccAzureServiceEndpointRule_SdkV2) GetTargetServices(ctx context.Context) ([]types.String, bool) {
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

// SetTargetServices sets the value of the TargetServices field in NccAzureServiceEndpointRule_SdkV2.
func (o *NccAzureServiceEndpointRule_SdkV2) SetTargetServices(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["target_services"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.TargetServices = types.ListValueMust(t, vs)
}

// The network connectivity rules that apply to network traffic from your
// serverless compute resources.
type NccEgressConfig_SdkV2 struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules types.List `tfsdk:"default_rules"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules types.List `tfsdk:"target_rules"`
}

func (newState *NccEgressConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressConfig_SdkV2) {
}

func (newState *NccEgressConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState NccEgressConfig_SdkV2) {
}

func (c NccEgressConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["default_rules"] = attrs["default_rules"].SetComputed()
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
func (a NccEgressConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"default_rules": reflect.TypeOf(NccEgressDefaultRules_SdkV2{}),
		"target_rules":  reflect.TypeOf(NccEgressTargetRules_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o NccEgressConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"default_rules": o.DefaultRules,
			"target_rules":  o.TargetRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccEgressConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *NccEgressConfig_SdkV2) GetDefaultRules(ctx context.Context) (NccEgressDefaultRules_SdkV2, bool) {
	var e NccEgressDefaultRules_SdkV2
	if o.DefaultRules.IsNull() || o.DefaultRules.IsUnknown() {
		return e, false
	}
	var v []NccEgressDefaultRules_SdkV2
	d := o.DefaultRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDefaultRules sets the value of the DefaultRules field in NccEgressConfig_SdkV2.
func (o *NccEgressConfig_SdkV2) SetDefaultRules(ctx context.Context, v NccEgressDefaultRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["default_rules"]
	o.DefaultRules = types.ListValueMust(t, vs)
}

// GetTargetRules returns the value of the TargetRules field in NccEgressConfig_SdkV2 as
// a NccEgressTargetRules_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *NccEgressConfig_SdkV2) GetTargetRules(ctx context.Context) (NccEgressTargetRules_SdkV2, bool) {
	var e NccEgressTargetRules_SdkV2
	if o.TargetRules.IsNull() || o.TargetRules.IsUnknown() {
		return e, false
	}
	var v []NccEgressTargetRules_SdkV2
	d := o.TargetRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTargetRules sets the value of the TargetRules field in NccEgressConfig_SdkV2.
func (o *NccEgressConfig_SdkV2) SetTargetRules(ctx context.Context, v NccEgressTargetRules_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["target_rules"]
	o.TargetRules = types.ListValueMust(t, vs)
}

// The network connectivity rules that are applied by default without resource
// specific configurations. You can find the stable network information of your
// serverless compute resources here.
type NccEgressDefaultRules_SdkV2 struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule types.List `tfsdk:"aws_stable_ip_rule"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule types.List `tfsdk:"azure_service_endpoint_rule"`
}

func (newState *NccEgressDefaultRules_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressDefaultRules_SdkV2) {
}

func (newState *NccEgressDefaultRules_SdkV2) SyncEffectiveFieldsDuringRead(existingState NccEgressDefaultRules_SdkV2) {
}

func (c NccEgressDefaultRules_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccEgressDefaultRules_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"aws_stable_ip_rule":          reflect.TypeOf(NccAwsStableIpRule_SdkV2{}),
		"azure_service_endpoint_rule": reflect.TypeOf(NccAzureServiceEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressDefaultRules_SdkV2
// only implements ToObjectValue() and Type().
func (o NccEgressDefaultRules_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_stable_ip_rule":          o.AwsStableIpRule,
			"azure_service_endpoint_rule": o.AzureServiceEndpointRule,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccEgressDefaultRules_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *NccEgressDefaultRules_SdkV2) GetAwsStableIpRule(ctx context.Context) (NccAwsStableIpRule_SdkV2, bool) {
	var e NccAwsStableIpRule_SdkV2
	if o.AwsStableIpRule.IsNull() || o.AwsStableIpRule.IsUnknown() {
		return e, false
	}
	var v []NccAwsStableIpRule_SdkV2
	d := o.AwsStableIpRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAwsStableIpRule sets the value of the AwsStableIpRule field in NccEgressDefaultRules_SdkV2.
func (o *NccEgressDefaultRules_SdkV2) SetAwsStableIpRule(ctx context.Context, v NccAwsStableIpRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["aws_stable_ip_rule"]
	o.AwsStableIpRule = types.ListValueMust(t, vs)
}

// GetAzureServiceEndpointRule returns the value of the AzureServiceEndpointRule field in NccEgressDefaultRules_SdkV2 as
// a NccAzureServiceEndpointRule_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *NccEgressDefaultRules_SdkV2) GetAzureServiceEndpointRule(ctx context.Context) (NccAzureServiceEndpointRule_SdkV2, bool) {
	var e NccAzureServiceEndpointRule_SdkV2
	if o.AzureServiceEndpointRule.IsNull() || o.AzureServiceEndpointRule.IsUnknown() {
		return e, false
	}
	var v []NccAzureServiceEndpointRule_SdkV2
	d := o.AzureServiceEndpointRule.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAzureServiceEndpointRule sets the value of the AzureServiceEndpointRule field in NccEgressDefaultRules_SdkV2.
func (o *NccEgressDefaultRules_SdkV2) SetAzureServiceEndpointRule(ctx context.Context, v NccAzureServiceEndpointRule_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_service_endpoint_rule"]
	o.AzureServiceEndpointRule = types.ListValueMust(t, vs)
}

// The network connectivity rules that configured for each destinations. These
// rules override default rules.
type NccEgressTargetRules_SdkV2 struct {
	AzurePrivateEndpointRules types.List `tfsdk:"azure_private_endpoint_rules"`
}

func (newState *NccEgressTargetRules_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NccEgressTargetRules_SdkV2) {
}

func (newState *NccEgressTargetRules_SdkV2) SyncEffectiveFieldsDuringRead(existingState NccEgressTargetRules_SdkV2) {
}

func (c NccEgressTargetRules_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NccEgressTargetRules_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"azure_private_endpoint_rules": reflect.TypeOf(NccAzurePrivateEndpointRule_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NccEgressTargetRules_SdkV2
// only implements ToObjectValue() and Type().
func (o NccEgressTargetRules_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"azure_private_endpoint_rules": o.AzurePrivateEndpointRules,
		})
}

// Type implements basetypes.ObjectValuable.
func (o NccEgressTargetRules_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"azure_private_endpoint_rules": basetypes.ListType{
				ElemType: NccAzurePrivateEndpointRule_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAzurePrivateEndpointRules returns the value of the AzurePrivateEndpointRules field in NccEgressTargetRules_SdkV2 as
// a slice of NccAzurePrivateEndpointRule_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *NccEgressTargetRules_SdkV2) GetAzurePrivateEndpointRules(ctx context.Context) ([]NccAzurePrivateEndpointRule_SdkV2, bool) {
	if o.AzurePrivateEndpointRules.IsNull() || o.AzurePrivateEndpointRules.IsUnknown() {
		return nil, false
	}
	var v []NccAzurePrivateEndpointRule_SdkV2
	d := o.AzurePrivateEndpointRules.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAzurePrivateEndpointRules sets the value of the AzurePrivateEndpointRules field in NccEgressTargetRules_SdkV2.
func (o *NccEgressTargetRules_SdkV2) SetAzurePrivateEndpointRules(ctx context.Context, v []NccAzurePrivateEndpointRule_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["azure_private_endpoint_rules"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AzurePrivateEndpointRules = types.ListValueMust(t, vs)
}

type NetworkConnectivityConfiguration_SdkV2 struct {
	// The Databricks account ID that hosts the credential.
	AccountId types.String `tfsdk:"account_id"`
	// Time in epoch milliseconds when this object was created.
	CreationTime types.Int64 `tfsdk:"creation_time"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig types.List `tfsdk:"egress_config"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
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

func (newState *NetworkConnectivityConfiguration_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NetworkConnectivityConfiguration_SdkV2) {
}

func (newState *NetworkConnectivityConfiguration_SdkV2) SyncEffectiveFieldsDuringRead(existingState NetworkConnectivityConfiguration_SdkV2) {
}

func (c NetworkConnectivityConfiguration_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["account_id"] = attrs["account_id"].SetOptional()
	attrs["creation_time"] = attrs["creation_time"].SetComputed()
	attrs["egress_config"] = attrs["egress_config"].SetOptional()
	attrs["egress_config"] = attrs["egress_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()
	attrs["network_connectivity_config_id"] = attrs["network_connectivity_config_id"].SetComputed()
	attrs["region"] = attrs["region"].SetOptional()
	attrs["updated_time"] = attrs["updated_time"].SetComputed()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in NetworkConnectivityConfiguration.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a NetworkConnectivityConfiguration_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"egress_config": reflect.TypeOf(NccEgressConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NetworkConnectivityConfiguration_SdkV2
// only implements ToObjectValue() and Type().
func (o NetworkConnectivityConfiguration_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o NetworkConnectivityConfiguration_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"account_id":    types.StringType,
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
func (o *NetworkConnectivityConfiguration_SdkV2) GetEgressConfig(ctx context.Context) (NccEgressConfig_SdkV2, bool) {
	var e NccEgressConfig_SdkV2
	if o.EgressConfig.IsNull() || o.EgressConfig.IsUnknown() {
		return e, false
	}
	var v []NccEgressConfig_SdkV2
	d := o.EgressConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEgressConfig sets the value of the EgressConfig field in NetworkConnectivityConfiguration_SdkV2.
func (o *NetworkConnectivityConfiguration_SdkV2) SetEgressConfig(ctx context.Context, v NccEgressConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["egress_config"]
	o.EgressConfig = types.ListValueMust(t, vs)
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

func (newState *NotificationDestination_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan NotificationDestination_SdkV2) {
}

func (newState *NotificationDestination_SdkV2) SyncEffectiveFieldsDuringRead(existingState NotificationDestination_SdkV2) {
}

func (c NotificationDestination_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a NotificationDestination_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, NotificationDestination_SdkV2
// only implements ToObjectValue() and Type().
func (o NotificationDestination_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o NotificationDestination_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *NotificationDestination_SdkV2) GetConfig(ctx context.Context) (Config_SdkV2, bool) {
	var e Config_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []Config_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in NotificationDestination_SdkV2.
func (o *NotificationDestination_SdkV2) SetConfig(ctx context.Context, v Config_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
}

type PagerdutyConfig_SdkV2 struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey types.String `tfsdk:"integration_key"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet types.Bool `tfsdk:"integration_key_set"`
}

func (newState *PagerdutyConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PagerdutyConfig_SdkV2) {
}

func (newState *PagerdutyConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState PagerdutyConfig_SdkV2) {
}

func (c PagerdutyConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PagerdutyConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PagerdutyConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o PagerdutyConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"integration_key":     o.IntegrationKey,
			"integration_key_set": o.IntegrationKeySet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PagerdutyConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
	WorkspaceId types.Int64 `tfsdk:"workspaceId"`
}

func (newState *PartitionId_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PartitionId_SdkV2) {
}

func (newState *PartitionId_SdkV2) SyncEffectiveFieldsDuringRead(existingState PartitionId_SdkV2) {
}

func (c PartitionId_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PartitionId_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PartitionId_SdkV2
// only implements ToObjectValue() and Type().
func (o PartitionId_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspaceId": o.WorkspaceId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PartitionId_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspaceId": types.Int64Type,
		},
	}
}

type PersonalComputeMessage_SdkV2 struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value types.String `tfsdk:"value"`
}

func (newState *PersonalComputeMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalComputeMessage_SdkV2) {
}

func (newState *PersonalComputeMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState PersonalComputeMessage_SdkV2) {
}

func (c PersonalComputeMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PersonalComputeMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o PersonalComputeMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalComputeMessage_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *PersonalComputeSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PersonalComputeSetting_SdkV2) {
}

func (newState *PersonalComputeSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState PersonalComputeSetting_SdkV2) {
}

func (c PersonalComputeSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PersonalComputeSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"personal_compute": reflect.TypeOf(PersonalComputeMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PersonalComputeSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o PersonalComputeSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":             o.Etag,
			"personal_compute": o.PersonalCompute,
			"setting_name":     o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PersonalComputeSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *PersonalComputeSetting_SdkV2) GetPersonalCompute(ctx context.Context) (PersonalComputeMessage_SdkV2, bool) {
	var e PersonalComputeMessage_SdkV2
	if o.PersonalCompute.IsNull() || o.PersonalCompute.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeMessage_SdkV2
	d := o.PersonalCompute.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPersonalCompute sets the value of the PersonalCompute field in PersonalComputeSetting_SdkV2.
func (o *PersonalComputeSetting_SdkV2) SetPersonalCompute(ctx context.Context, v PersonalComputeMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["personal_compute"]
	o.PersonalCompute = types.ListValueMust(t, vs)
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

func (newState *PublicTokenInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublicTokenInfo_SdkV2) {
}

func (newState *PublicTokenInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState PublicTokenInfo_SdkV2) {
}

func (c PublicTokenInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PublicTokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PublicTokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o PublicTokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o PublicTokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type"`
}

func (newState *ReplaceIpAccessList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ReplaceIpAccessList_SdkV2) {
}

func (newState *ReplaceIpAccessList_SdkV2) SyncEffectiveFieldsDuringRead(existingState ReplaceIpAccessList_SdkV2) {
}

func (c ReplaceIpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ReplaceIpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceIpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (o ReplaceIpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ReplaceIpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *ReplaceIpAccessList_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in ReplaceIpAccessList_SdkV2.
func (o *ReplaceIpAccessList_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
}

type ReplaceResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ReplaceResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ReplaceResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ReplaceResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ReplaceResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ReplaceResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type RestrictWorkspaceAdminsMessage_SdkV2 struct {
	Status types.String `tfsdk:"status"`
}

func (newState *RestrictWorkspaceAdminsMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestrictWorkspaceAdminsMessage_SdkV2) {
}

func (newState *RestrictWorkspaceAdminsMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState RestrictWorkspaceAdminsMessage_SdkV2) {
}

func (c RestrictWorkspaceAdminsMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestrictWorkspaceAdminsMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o RestrictWorkspaceAdminsMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"status": o.Status,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestrictWorkspaceAdminsMessage_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *RestrictWorkspaceAdminsSetting_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestrictWorkspaceAdminsSetting_SdkV2) {
}

func (newState *RestrictWorkspaceAdminsSetting_SdkV2) SyncEffectiveFieldsDuringRead(existingState RestrictWorkspaceAdminsSetting_SdkV2) {
}

func (c RestrictWorkspaceAdminsSetting_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RestrictWorkspaceAdminsSetting_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"restrict_workspace_admins": reflect.TypeOf(RestrictWorkspaceAdminsMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RestrictWorkspaceAdminsSetting_SdkV2
// only implements ToObjectValue() and Type().
func (o RestrictWorkspaceAdminsSetting_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"etag":                      o.Etag,
			"restrict_workspace_admins": o.RestrictWorkspaceAdmins,
			"setting_name":              o.SettingName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RestrictWorkspaceAdminsSetting_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *RestrictWorkspaceAdminsSetting_SdkV2) GetRestrictWorkspaceAdmins(ctx context.Context) (RestrictWorkspaceAdminsMessage_SdkV2, bool) {
	var e RestrictWorkspaceAdminsMessage_SdkV2
	if o.RestrictWorkspaceAdmins.IsNull() || o.RestrictWorkspaceAdmins.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsMessage_SdkV2
	d := o.RestrictWorkspaceAdmins.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetRestrictWorkspaceAdmins sets the value of the RestrictWorkspaceAdmins field in RestrictWorkspaceAdminsSetting_SdkV2.
func (o *RestrictWorkspaceAdminsSetting_SdkV2) SetRestrictWorkspaceAdmins(ctx context.Context, v RestrictWorkspaceAdminsMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["restrict_workspace_admins"]
	o.RestrictWorkspaceAdmins = types.ListValueMust(t, vs)
}

type RevokeTokenRequest_SdkV2 struct {
	// The ID of the token to be revoked.
	TokenId types.String `tfsdk:"token_id"`
}

func (newState *RevokeTokenRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenRequest_SdkV2) {
}

func (newState *RevokeTokenRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState RevokeTokenRequest_SdkV2) {
}

func (c RevokeTokenRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RevokeTokenRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o RevokeTokenRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token_id": o.TokenId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RevokeTokenRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token_id": types.StringType,
		},
	}
}

type RevokeTokenResponse_SdkV2 struct {
}

func (newState *RevokeTokenResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RevokeTokenResponse_SdkV2) {
}

func (newState *RevokeTokenResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState RevokeTokenResponse_SdkV2) {
}

func (c RevokeTokenResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RevokeTokenResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RevokeTokenResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RevokeTokenResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o RevokeTokenResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o RevokeTokenResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SetStatusResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SetStatusResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SetStatusResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SetStatusResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o SetStatusResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o SetStatusResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type SlackConfig_SdkV2 struct {
	// [Input-Only] URL for Slack destination.
	Url types.String `tfsdk:"url"`
	// [Output-Only] Whether URL is set.
	UrlSet types.Bool `tfsdk:"url_set"`
}

func (newState *SlackConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan SlackConfig_SdkV2) {
}

func (newState *SlackConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState SlackConfig_SdkV2) {
}

func (c SlackConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a SlackConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, SlackConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o SlackConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"url":     o.Url,
			"url_set": o.UrlSet,
		})
}

// Type implements basetypes.ObjectValuable.
func (o SlackConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"url":     types.StringType,
			"url_set": types.BoolType,
		},
	}
}

type StringMessage_SdkV2 struct {
	// Represents a generic string value.
	Value types.String `tfsdk:"value"`
}

func (newState *StringMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan StringMessage_SdkV2) {
}

func (newState *StringMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState StringMessage_SdkV2) {
}

func (c StringMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a StringMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, StringMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o StringMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o StringMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"value": types.StringType,
		},
	}
}

type TokenAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (newState *TokenAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessControlRequest_SdkV2) {
}

func (newState *TokenAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenAccessControlRequest_SdkV2) {
}

func (c TokenAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TokenAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *TokenAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenAccessControlResponse_SdkV2) {
}

func (newState *TokenAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenAccessControlResponse_SdkV2) {
}

func (c TokenAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(TokenPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TokenAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *TokenAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]TokenPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []TokenPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in TokenAccessControlResponse_SdkV2.
func (o *TokenAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []TokenPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
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

func (newState *TokenInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenInfo_SdkV2) {
}

func (newState *TokenInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenInfo_SdkV2) {
}

func (c TokenInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o TokenInfo_SdkV2) Type(ctx context.Context) attr.Type {
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
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *TokenPermission_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermission_SdkV2) {
}

func (newState *TokenPermission_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenPermission_SdkV2) {
}

func (c TokenPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *TokenPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in TokenPermission_SdkV2.
func (o *TokenPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type TokenPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (newState *TokenPermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissions_SdkV2) {
}

func (newState *TokenPermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenPermissions_SdkV2) {
}

func (c TokenPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermissions_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *TokenPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]TokenAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissions_SdkV2.
func (o *TokenPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []TokenAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type TokenPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (newState *TokenPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissionsDescription_SdkV2) {
}

func (newState *TokenPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenPermissionsDescription_SdkV2) {
}

func (c TokenPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
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

func (newState *TokenPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TokenPermissionsRequest_SdkV2) {
}

func (newState *TokenPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState TokenPermissionsRequest_SdkV2) {
}

func (c TokenPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TokenPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(TokenAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TokenPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o TokenPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TokenPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *TokenPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]TokenAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []TokenAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in TokenPermissionsRequest_SdkV2.
func (o *TokenPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []TokenAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
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

func (newState *UpdateAccountIpAccessEnableRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAccountIpAccessEnableRequest_SdkV2) {
}

func (newState *UpdateAccountIpAccessEnableRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAccountIpAccessEnableRequest_SdkV2) {
}

func (c UpdateAccountIpAccessEnableRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAccountIpAccessEnableRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AccountIpAccessEnable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAccountIpAccessEnableRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAccountIpAccessEnableRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAccountIpAccessEnableRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAccountIpAccessEnableRequest_SdkV2) GetSetting(ctx context.Context) (AccountIpAccessEnable_SdkV2, bool) {
	var e AccountIpAccessEnable_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AccountIpAccessEnable_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAccountIpAccessEnableRequest_SdkV2.
func (o *UpdateAccountIpAccessEnableRequest_SdkV2) SetSetting(ctx context.Context, v AccountIpAccessEnable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
}

func (newState *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) {
}

func (c UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingAccessPolicySetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) GetSetting(ctx context.Context) (AibiDashboardEmbeddingAccessPolicySetting_SdkV2, bool) {
	var e AibiDashboardEmbeddingAccessPolicySetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingAccessPolicySetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2.
func (o *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest_SdkV2) SetSetting(ctx context.Context, v AibiDashboardEmbeddingAccessPolicySetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
}

func (newState *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) {
}

func (c UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) GetSetting(ctx context.Context) (AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2, bool) {
	var e AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2.
func (o *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest_SdkV2) SetSetting(ctx context.Context, v AibiDashboardEmbeddingApprovedDomainsSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAutomaticClusterUpdateSettingRequest_SdkV2) {
}

func (newState *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateAutomaticClusterUpdateSettingRequest_SdkV2) {
}

func (c UpdateAutomaticClusterUpdateSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateAutomaticClusterUpdateSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(AutomaticClusterUpdateSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateAutomaticClusterUpdateSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateAutomaticClusterUpdateSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateAutomaticClusterUpdateSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) GetSetting(ctx context.Context) (AutomaticClusterUpdateSetting_SdkV2, bool) {
	var e AutomaticClusterUpdateSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []AutomaticClusterUpdateSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateAutomaticClusterUpdateSettingRequest_SdkV2.
func (o *UpdateAutomaticClusterUpdateSettingRequest_SdkV2) SetSetting(ctx context.Context, v AutomaticClusterUpdateSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateComplianceSecurityProfileSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateComplianceSecurityProfileSettingRequest_SdkV2) {
}

func (newState *UpdateComplianceSecurityProfileSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateComplianceSecurityProfileSettingRequest_SdkV2) {
}

func (c UpdateComplianceSecurityProfileSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateComplianceSecurityProfileSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(ComplianceSecurityProfileSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateComplianceSecurityProfileSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateComplianceSecurityProfileSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateComplianceSecurityProfileSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateComplianceSecurityProfileSettingRequest_SdkV2) GetSetting(ctx context.Context) (ComplianceSecurityProfileSetting_SdkV2, bool) {
	var e ComplianceSecurityProfileSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []ComplianceSecurityProfileSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateComplianceSecurityProfileSettingRequest_SdkV2.
func (o *UpdateComplianceSecurityProfileSettingRequest_SdkV2) SetSetting(ctx context.Context, v ComplianceSecurityProfileSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateCspEnablementAccountSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateCspEnablementAccountSettingRequest_SdkV2) {
}

func (newState *UpdateCspEnablementAccountSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateCspEnablementAccountSettingRequest_SdkV2) {
}

func (c UpdateCspEnablementAccountSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateCspEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(CspEnablementAccountSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateCspEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateCspEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateCspEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateCspEnablementAccountSettingRequest_SdkV2) GetSetting(ctx context.Context) (CspEnablementAccountSetting_SdkV2, bool) {
	var e CspEnablementAccountSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []CspEnablementAccountSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateCspEnablementAccountSettingRequest_SdkV2.
func (o *UpdateCspEnablementAccountSettingRequest_SdkV2) SetSetting(ctx context.Context, v CspEnablementAccountSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	Setting types.List `tfsdk:"setting"`
}

func (newState *UpdateDefaultNamespaceSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDefaultNamespaceSettingRequest_SdkV2) {
}

func (newState *UpdateDefaultNamespaceSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateDefaultNamespaceSettingRequest_SdkV2) {
}

func (c UpdateDefaultNamespaceSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDefaultNamespaceSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DefaultNamespaceSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDefaultNamespaceSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDefaultNamespaceSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDefaultNamespaceSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDefaultNamespaceSettingRequest_SdkV2) GetSetting(ctx context.Context) (DefaultNamespaceSetting_SdkV2, bool) {
	var e DefaultNamespaceSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DefaultNamespaceSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDefaultNamespaceSettingRequest_SdkV2.
func (o *UpdateDefaultNamespaceSettingRequest_SdkV2) SetSetting(ctx context.Context, v DefaultNamespaceSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateDisableLegacyAccessRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyAccessRequest_SdkV2) {
}

func (newState *UpdateDisableLegacyAccessRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyAccessRequest_SdkV2) {
}

func (c UpdateDisableLegacyAccessRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDisableLegacyAccessRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyAccess_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyAccessRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDisableLegacyAccessRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDisableLegacyAccessRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDisableLegacyAccessRequest_SdkV2) GetSetting(ctx context.Context) (DisableLegacyAccess_SdkV2, bool) {
	var e DisableLegacyAccess_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyAccess_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyAccessRequest_SdkV2.
func (o *UpdateDisableLegacyAccessRequest_SdkV2) SetSetting(ctx context.Context, v DisableLegacyAccess_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateDisableLegacyDbfsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyDbfsRequest_SdkV2) {
}

func (newState *UpdateDisableLegacyDbfsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyDbfsRequest_SdkV2) {
}

func (c UpdateDisableLegacyDbfsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDisableLegacyDbfsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyDbfs_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyDbfsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDisableLegacyDbfsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDisableLegacyDbfsRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDisableLegacyDbfsRequest_SdkV2) GetSetting(ctx context.Context) (DisableLegacyDbfs_SdkV2, bool) {
	var e DisableLegacyDbfs_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyDbfs_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyDbfsRequest_SdkV2.
func (o *UpdateDisableLegacyDbfsRequest_SdkV2) SetSetting(ctx context.Context, v DisableLegacyDbfs_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateDisableLegacyFeaturesRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDisableLegacyFeaturesRequest_SdkV2) {
}

func (newState *UpdateDisableLegacyFeaturesRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateDisableLegacyFeaturesRequest_SdkV2) {
}

func (c UpdateDisableLegacyFeaturesRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateDisableLegacyFeaturesRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(DisableLegacyFeatures_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateDisableLegacyFeaturesRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateDisableLegacyFeaturesRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateDisableLegacyFeaturesRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateDisableLegacyFeaturesRequest_SdkV2) GetSetting(ctx context.Context) (DisableLegacyFeatures_SdkV2, bool) {
	var e DisableLegacyFeatures_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []DisableLegacyFeatures_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateDisableLegacyFeaturesRequest_SdkV2.
func (o *UpdateDisableLegacyFeaturesRequest_SdkV2) SetSetting(ctx context.Context, v DisableLegacyFeatures_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) {
}

func (newState *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) {
}

func (c UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EnhancedSecurityMonitoringSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) GetSetting(ctx context.Context) (EnhancedSecurityMonitoringSetting_SdkV2, bool) {
	var e EnhancedSecurityMonitoringSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EnhancedSecurityMonitoringSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2.
func (o *UpdateEnhancedSecurityMonitoringSettingRequest_SdkV2) SetSetting(ctx context.Context, v EnhancedSecurityMonitoringSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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

func (newState *UpdateEsmEnablementAccountSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateEsmEnablementAccountSettingRequest_SdkV2) {
}

func (newState *UpdateEsmEnablementAccountSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateEsmEnablementAccountSettingRequest_SdkV2) {
}

func (c UpdateEsmEnablementAccountSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateEsmEnablementAccountSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(EsmEnablementAccountSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateEsmEnablementAccountSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateEsmEnablementAccountSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateEsmEnablementAccountSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateEsmEnablementAccountSettingRequest_SdkV2) GetSetting(ctx context.Context) (EsmEnablementAccountSetting_SdkV2, bool) {
	var e EsmEnablementAccountSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []EsmEnablementAccountSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateEsmEnablementAccountSettingRequest_SdkV2.
func (o *UpdateEsmEnablementAccountSettingRequest_SdkV2) SetSetting(ctx context.Context, v EsmEnablementAccountSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
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
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType types.String `tfsdk:"list_type"`
}

func (newState *UpdateIpAccessList_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateIpAccessList_SdkV2) {
}

func (newState *UpdateIpAccessList_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateIpAccessList_SdkV2) {
}

func (c UpdateIpAccessList_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateIpAccessList_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ip_addresses": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateIpAccessList_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateIpAccessList_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o UpdateIpAccessList_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateIpAccessList_SdkV2) GetIpAddresses(ctx context.Context) ([]types.String, bool) {
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

// SetIpAddresses sets the value of the IpAddresses field in UpdateIpAccessList_SdkV2.
func (o *UpdateIpAccessList_SdkV2) SetIpAddresses(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ip_addresses"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.IpAddresses = types.ListValueMust(t, vs)
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

func (newState *UpdateNotificationDestinationRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateNotificationDestinationRequest_SdkV2) {
}

func (newState *UpdateNotificationDestinationRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateNotificationDestinationRequest_SdkV2) {
}

func (c UpdateNotificationDestinationRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateNotificationDestinationRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(Config_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateNotificationDestinationRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateNotificationDestinationRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config":       o.Config,
			"display_name": o.DisplayName,
			"id":           o.Id,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateNotificationDestinationRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateNotificationDestinationRequest_SdkV2) GetConfig(ctx context.Context) (Config_SdkV2, bool) {
	var e Config_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []Config_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in UpdateNotificationDestinationRequest_SdkV2.
func (o *UpdateNotificationDestinationRequest_SdkV2) SetConfig(ctx context.Context, v Config_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
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

func (newState *UpdatePersonalComputeSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdatePersonalComputeSettingRequest_SdkV2) {
}

func (newState *UpdatePersonalComputeSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdatePersonalComputeSettingRequest_SdkV2) {
}

func (c UpdatePersonalComputeSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdatePersonalComputeSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(PersonalComputeSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdatePersonalComputeSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdatePersonalComputeSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdatePersonalComputeSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdatePersonalComputeSettingRequest_SdkV2) GetSetting(ctx context.Context) (PersonalComputeSetting_SdkV2, bool) {
	var e PersonalComputeSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []PersonalComputeSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdatePersonalComputeSettingRequest_SdkV2.
func (o *UpdatePersonalComputeSettingRequest_SdkV2) SetSetting(ctx context.Context, v PersonalComputeSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}

type UpdateResponse_SdkV2 struct {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (newState *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) {
}

func (newState *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) {
}

func (c UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"setting": reflect.TypeOf(RestrictWorkspaceAdminsSetting_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"allow_missing": o.AllowMissing,
			"field_mask":    o.FieldMask,
			"setting":       o.Setting,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
func (o *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) GetSetting(ctx context.Context) (RestrictWorkspaceAdminsSetting_SdkV2, bool) {
	var e RestrictWorkspaceAdminsSetting_SdkV2
	if o.Setting.IsNull() || o.Setting.IsUnknown() {
		return e, false
	}
	var v []RestrictWorkspaceAdminsSetting_SdkV2
	d := o.Setting.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetSetting sets the value of the Setting field in UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2.
func (o *UpdateRestrictWorkspaceAdminsSettingRequest_SdkV2) SetSetting(ctx context.Context, v RestrictWorkspaceAdminsSetting_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["setting"]
	o.Setting = types.ListValueMust(t, vs)
}
