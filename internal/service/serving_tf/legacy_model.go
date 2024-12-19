// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package serving_tf

import (
	"context"
	"reflect"

	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"

	"github.com/databricks/terraform-provider-databricks/internal/service/oauth2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Ai21LabsConfig_SdkV2 struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKey types.String `tfsdk:"ai21labs_api_key" tf:"optional"`
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKeyPlaintext types.String `tfsdk:"ai21labs_api_key_plaintext" tf:"optional"`
}

func (newState *Ai21LabsConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Ai21LabsConfig_SdkV2) {
}

func (newState *Ai21LabsConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState Ai21LabsConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Ai21LabsConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Ai21LabsConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Ai21LabsConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o Ai21LabsConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai21labs_api_key":           o.Ai21labsApiKey,
			"ai21labs_api_key_plaintext": o.Ai21labsApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Ai21LabsConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_api_key":           types.StringType,
			"ai21labs_api_key_plaintext": types.StringType,
		},
	}
}

type AiGatewayConfig_SdkV2 struct {
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.List `tfsdk:"guardrails" tf:"optional,object"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.List `tfsdk:"inference_table_config" tf:"optional,object"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.List `tfsdk:"usage_tracking_config" tf:"optional,object"`
}

func (newState *AiGatewayConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayConfig_SdkV2) {
}

func (newState *AiGatewayConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails_SdkV2{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig_SdkV2{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit_SdkV2{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"guardrails":             o.Guardrails,
			"inference_table_config": o.InferenceTableConfig,
			"rate_limits":            o.RateLimits,
			"usage_tracking_config":  o.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails{}.Type(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig{}.Type(ctx),
			},
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.Type(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig{}.Type(ctx),
			},
		},
	}
}

// GetGuardrails returns the value of the Guardrails field in AiGatewayConfig_SdkV2 as
// a AiGatewayGuardrails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig_SdkV2) GetGuardrails(ctx context.Context) (AiGatewayGuardrails_SdkV2, bool) {
	var e AiGatewayGuardrails_SdkV2
	if o.Guardrails.IsNull() || o.Guardrails.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrails_SdkV2
	d := o.Guardrails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGuardrails sets the value of the Guardrails field in AiGatewayConfig_SdkV2.
func (o *AiGatewayConfig_SdkV2) SetGuardrails(ctx context.Context, v AiGatewayGuardrails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guardrails"]
	o.Guardrails = types.ListValueMust(t, vs)
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in AiGatewayConfig_SdkV2 as
// a AiGatewayInferenceTableConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig_SdkV2) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig_SdkV2, bool) {
	var e AiGatewayInferenceTableConfig_SdkV2
	if o.InferenceTableConfig.IsNull() || o.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayInferenceTableConfig_SdkV2
	d := o.InferenceTableConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in AiGatewayConfig_SdkV2.
func (o *AiGatewayConfig_SdkV2) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_table_config"]
	o.InferenceTableConfig = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in AiGatewayConfig_SdkV2 as
// a slice of AiGatewayRateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig_SdkV2) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit_SdkV2, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit_SdkV2
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in AiGatewayConfig_SdkV2.
func (o *AiGatewayConfig_SdkV2) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in AiGatewayConfig_SdkV2 as
// a AiGatewayUsageTrackingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig_SdkV2) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig_SdkV2, bool) {
	var e AiGatewayUsageTrackingConfig_SdkV2
	if o.UsageTrackingConfig.IsNull() || o.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayUsageTrackingConfig_SdkV2
	d := o.UsageTrackingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in AiGatewayConfig_SdkV2.
func (o *AiGatewayConfig_SdkV2) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_tracking_config"]
	o.UsageTrackingConfig = types.ListValueMust(t, vs)
}

type AiGatewayGuardrailParameters_SdkV2 struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	InvalidKeywords types.List `tfsdk:"invalid_keywords" tf:"optional"`
	// Configuration for guardrail PII filter.
	Pii types.List `tfsdk:"pii" tf:"optional,object"`
	// Indicates whether the safety filter is enabled.
	Safety types.Bool `tfsdk:"safety" tf:"optional"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics types.List `tfsdk:"valid_topics" tf:"optional"`
}

func (newState *AiGatewayGuardrailParameters_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrailParameters_SdkV2) {
}

func (newState *AiGatewayGuardrailParameters_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrailParameters_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrailParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayGuardrailParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"invalid_keywords": reflect.TypeOf(types.String{}),
		"pii":              reflect.TypeOf(AiGatewayGuardrailPiiBehavior_SdkV2{}),
		"valid_topics":     reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrailParameters_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayGuardrailParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"invalid_keywords": o.InvalidKeywords,
			"pii":              o.Pii,
			"safety":           o.Safety,
			"valid_topics":     o.ValidTopics,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayGuardrailParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"invalid_keywords": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pii": basetypes.ListType{
				ElemType: AiGatewayGuardrailPiiBehavior{}.Type(ctx),
			},
			"safety": types.BoolType,
			"valid_topics": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetInvalidKeywords returns the value of the InvalidKeywords field in AiGatewayGuardrailParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrailParameters_SdkV2) GetInvalidKeywords(ctx context.Context) ([]types.String, bool) {
	if o.InvalidKeywords.IsNull() || o.InvalidKeywords.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.InvalidKeywords.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInvalidKeywords sets the value of the InvalidKeywords field in AiGatewayGuardrailParameters_SdkV2.
func (o *AiGatewayGuardrailParameters_SdkV2) SetInvalidKeywords(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["invalid_keywords"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InvalidKeywords = types.ListValueMust(t, vs)
}

// GetPii returns the value of the Pii field in AiGatewayGuardrailParameters_SdkV2 as
// a AiGatewayGuardrailPiiBehavior_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrailParameters_SdkV2) GetPii(ctx context.Context) (AiGatewayGuardrailPiiBehavior_SdkV2, bool) {
	var e AiGatewayGuardrailPiiBehavior_SdkV2
	if o.Pii.IsNull() || o.Pii.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrailPiiBehavior_SdkV2
	d := o.Pii.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPii sets the value of the Pii field in AiGatewayGuardrailParameters_SdkV2.
func (o *AiGatewayGuardrailParameters_SdkV2) SetPii(ctx context.Context, v AiGatewayGuardrailPiiBehavior_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pii"]
	o.Pii = types.ListValueMust(t, vs)
}

// GetValidTopics returns the value of the ValidTopics field in AiGatewayGuardrailParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrailParameters_SdkV2) GetValidTopics(ctx context.Context) ([]types.String, bool) {
	if o.ValidTopics.IsNull() || o.ValidTopics.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.ValidTopics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValidTopics sets the value of the ValidTopics field in AiGatewayGuardrailParameters_SdkV2.
func (o *AiGatewayGuardrailParameters_SdkV2) SetValidTopics(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["valid_topics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ValidTopics = types.ListValueMust(t, vs)
}

type AiGatewayGuardrailPiiBehavior_SdkV2 struct {
	// Behavior for PII filter. Currently only 'BLOCK' is supported. If 'BLOCK'
	// is set for the input guardrail and the request contains PII, the request
	// is not sent to the model server and 400 status code is returned; if
	// 'BLOCK' is set for the output guardrail and the model response contains
	// PII, the PII info in the response is redacted and 400 status code is
	// returned.
	Behavior types.String `tfsdk:"behavior" tf:""`
}

func (newState *AiGatewayGuardrailPiiBehavior_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrailPiiBehavior_SdkV2) {
}

func (newState *AiGatewayGuardrailPiiBehavior_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrailPiiBehavior_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrailPiiBehavior.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayGuardrailPiiBehavior_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrailPiiBehavior_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayGuardrailPiiBehavior_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"behavior": o.Behavior,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayGuardrailPiiBehavior_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"behavior": types.StringType,
		},
	}
}

type AiGatewayGuardrails_SdkV2 struct {
	// Configuration for input guardrail filters.
	Input types.List `tfsdk:"input" tf:"optional,object"`
	// Configuration for output guardrail filters.
	Output types.List `tfsdk:"output" tf:"optional,object"`
}

func (newState *AiGatewayGuardrails_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrails_SdkV2) {
}

func (newState *AiGatewayGuardrails_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrails_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayGuardrails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input":  reflect.TypeOf(AiGatewayGuardrailParameters_SdkV2{}),
		"output": reflect.TypeOf(AiGatewayGuardrailParameters_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrails_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayGuardrails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"input":  o.Input,
			"output": o.Output,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayGuardrails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"input": basetypes.ListType{
				ElemType: AiGatewayGuardrailParameters{}.Type(ctx),
			},
			"output": basetypes.ListType{
				ElemType: AiGatewayGuardrailParameters{}.Type(ctx),
			},
		},
	}
}

// GetInput returns the value of the Input field in AiGatewayGuardrails_SdkV2 as
// a AiGatewayGuardrailParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrails_SdkV2) GetInput(ctx context.Context) (AiGatewayGuardrailParameters_SdkV2, bool) {
	var e AiGatewayGuardrailParameters_SdkV2
	if o.Input.IsNull() || o.Input.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrailParameters_SdkV2
	d := o.Input.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInput sets the value of the Input field in AiGatewayGuardrails_SdkV2.
func (o *AiGatewayGuardrails_SdkV2) SetInput(ctx context.Context, v AiGatewayGuardrailParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["input"]
	o.Input = types.ListValueMust(t, vs)
}

// GetOutput returns the value of the Output field in AiGatewayGuardrails_SdkV2 as
// a AiGatewayGuardrailParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrails_SdkV2) GetOutput(ctx context.Context) (AiGatewayGuardrailParameters_SdkV2, bool) {
	var e AiGatewayGuardrailParameters_SdkV2
	if o.Output.IsNull() || o.Output.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrailParameters_SdkV2
	d := o.Output.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutput sets the value of the Output field in AiGatewayGuardrails_SdkV2.
func (o *AiGatewayGuardrails_SdkV2) SetOutput(ctx context.Context, v AiGatewayGuardrailParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["output"]
	o.Output = types.ListValueMust(t, vs)
}

type AiGatewayInferenceTableConfig_SdkV2 struct {
	// The name of the catalog in Unity Catalog. Required when enabling
	// inference tables. NOTE: On update, you have to disable inference table
	// first in order to change the catalog name.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	TableNamePrefix types.String `tfsdk:"table_name_prefix" tf:"optional"`
}

func (newState *AiGatewayInferenceTableConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayInferenceTableConfig_SdkV2) {
}

func (newState *AiGatewayInferenceTableConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayInferenceTableConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayInferenceTableConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayInferenceTableConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayInferenceTableConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayInferenceTableConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":      o.CatalogName,
			"enabled":           o.Enabled,
			"schema_name":       o.SchemaName,
			"table_name_prefix": o.TableNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayInferenceTableConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"enabled":           types.BoolType,
			"schema_name":       types.StringType,
			"table_name_prefix": types.StringType,
		},
	}
}

type AiGatewayRateLimit_SdkV2 struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls types.Int64 `tfsdk:"calls" tf:""`
	// Key field for a rate limit. Currently, only 'user' and 'endpoint' are
	// supported, with 'endpoint' being the default if not specified.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	RenewalPeriod types.String `tfsdk:"renewal_period" tf:""`
}

func (newState *AiGatewayRateLimit_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayRateLimit_SdkV2) {
}

func (newState *AiGatewayRateLimit_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayRateLimit_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayRateLimit.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayRateLimit_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayRateLimit_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayRateLimit_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"calls":          o.Calls,
			"key":            o.Key,
			"renewal_period": o.RenewalPeriod,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayRateLimit_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"renewal_period": types.StringType,
		},
	}
}

type AiGatewayUsageTrackingConfig_SdkV2 struct {
	// Whether to enable usage tracking.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
}

func (newState *AiGatewayUsageTrackingConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayUsageTrackingConfig_SdkV2) {
}

func (newState *AiGatewayUsageTrackingConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState AiGatewayUsageTrackingConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayUsageTrackingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayUsageTrackingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayUsageTrackingConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o AiGatewayUsageTrackingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": o.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayUsageTrackingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type AmazonBedrockConfig_SdkV2 struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id`. You must provide an API
	// key using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	AwsAccessKeyId types.String `tfsdk:"aws_access_key_id" tf:"optional"`
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	AwsAccessKeyIdPlaintext types.String `tfsdk:"aws_access_key_id_plaintext" tf:"optional"`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion types.String `tfsdk:"aws_region" tf:""`
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKey types.String `tfsdk:"aws_secret_access_key" tf:"optional"`
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKeyPlaintext types.String `tfsdk:"aws_secret_access_key_plaintext" tf:"optional"`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider types.String `tfsdk:"bedrock_provider" tf:""`
}

func (newState *AmazonBedrockConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AmazonBedrockConfig_SdkV2) {
}

func (newState *AmazonBedrockConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState AmazonBedrockConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AmazonBedrockConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AmazonBedrockConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AmazonBedrockConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o AmazonBedrockConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_access_key_id":               o.AwsAccessKeyId,
			"aws_access_key_id_plaintext":     o.AwsAccessKeyIdPlaintext,
			"aws_region":                      o.AwsRegion,
			"aws_secret_access_key":           o.AwsSecretAccessKey,
			"aws_secret_access_key_plaintext": o.AwsSecretAccessKeyPlaintext,
			"bedrock_provider":                o.BedrockProvider,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AmazonBedrockConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_access_key_id":               types.StringType,
			"aws_access_key_id_plaintext":     types.StringType,
			"aws_region":                      types.StringType,
			"aws_secret_access_key":           types.StringType,
			"aws_secret_access_key_plaintext": types.StringType,
			"bedrock_provider":                types.StringType,
		},
	}
}

type AnthropicConfig_SdkV2 struct {
	// The Databricks secret key reference for an Anthropic API key. If you
	// prefer to paste your API key directly, see `anthropic_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKey types.String `tfsdk:"anthropic_api_key" tf:"optional"`
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKeyPlaintext types.String `tfsdk:"anthropic_api_key_plaintext" tf:"optional"`
}

func (newState *AnthropicConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AnthropicConfig_SdkV2) {
}

func (newState *AnthropicConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState AnthropicConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AnthropicConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AnthropicConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnthropicConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o AnthropicConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anthropic_api_key":           o.AnthropicApiKey,
			"anthropic_api_key_plaintext": o.AnthropicApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AnthropicConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anthropic_api_key":           types.StringType,
			"anthropic_api_key_plaintext": types.StringType,
		},
	}
}

type AutoCaptureConfigInput_SdkV2 struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix types.String `tfsdk:"table_name_prefix" tf:"optional"`
}

func (newState *AutoCaptureConfigInput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureConfigInput_SdkV2) {
}

func (newState *AutoCaptureConfigInput_SdkV2) SyncEffectiveFieldsDuringRead(existingState AutoCaptureConfigInput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureConfigInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AutoCaptureConfigInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureConfigInput_SdkV2
// only implements ToObjectValue() and Type().
func (o AutoCaptureConfigInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":      o.CatalogName,
			"enabled":           o.Enabled,
			"schema_name":       o.SchemaName,
			"table_name_prefix": o.TableNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutoCaptureConfigInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"enabled":           types.BoolType,
			"schema_name":       types.StringType,
			"table_name_prefix": types.StringType,
		},
	}
}

type AutoCaptureConfigOutput_SdkV2 struct {
	// The name of the catalog in Unity Catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the schema in Unity Catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`

	State types.List `tfsdk:"state" tf:"optional,object"`
	// The prefix of the table in Unity Catalog.
	TableNamePrefix types.String `tfsdk:"table_name_prefix" tf:"optional"`
}

func (newState *AutoCaptureConfigOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureConfigOutput_SdkV2) {
}

func (newState *AutoCaptureConfigOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState AutoCaptureConfigOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureConfigOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AutoCaptureConfigOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"state": reflect.TypeOf(AutoCaptureState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureConfigOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o AutoCaptureConfigOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":      o.CatalogName,
			"enabled":           o.Enabled,
			"schema_name":       o.SchemaName,
			"state":             o.State,
			"table_name_prefix": o.TableNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutoCaptureConfigOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"enabled":      types.BoolType,
			"schema_name":  types.StringType,
			"state": basetypes.ListType{
				ElemType: AutoCaptureState{}.Type(ctx),
			},
			"table_name_prefix": types.StringType,
		},
	}
}

// GetState returns the value of the State field in AutoCaptureConfigOutput_SdkV2 as
// a AutoCaptureState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AutoCaptureConfigOutput_SdkV2) GetState(ctx context.Context) (AutoCaptureState_SdkV2, bool) {
	var e AutoCaptureState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in AutoCaptureConfigOutput_SdkV2.
func (o *AutoCaptureConfigOutput_SdkV2) SetState(ctx context.Context, v AutoCaptureState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

type AutoCaptureState_SdkV2 struct {
	PayloadTable types.List `tfsdk:"payload_table" tf:"optional,object"`
}

func (newState *AutoCaptureState_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureState_SdkV2) {
}

func (newState *AutoCaptureState_SdkV2) SyncEffectiveFieldsDuringRead(existingState AutoCaptureState_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AutoCaptureState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"payload_table": reflect.TypeOf(PayloadTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureState_SdkV2
// only implements ToObjectValue() and Type().
func (o AutoCaptureState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"payload_table": o.PayloadTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutoCaptureState_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"payload_table": basetypes.ListType{
				ElemType: PayloadTable{}.Type(ctx),
			},
		},
	}
}

// GetPayloadTable returns the value of the PayloadTable field in AutoCaptureState_SdkV2 as
// a PayloadTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *AutoCaptureState_SdkV2) GetPayloadTable(ctx context.Context) (PayloadTable_SdkV2, bool) {
	var e PayloadTable_SdkV2
	if o.PayloadTable.IsNull() || o.PayloadTable.IsUnknown() {
		return e, false
	}
	var v []PayloadTable_SdkV2
	d := o.PayloadTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPayloadTable sets the value of the PayloadTable field in AutoCaptureState_SdkV2.
func (o *AutoCaptureState_SdkV2) SetPayloadTable(ctx context.Context, v PayloadTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["payload_table"]
	o.PayloadTable = types.ListValueMust(t, vs)
}

// Get build logs for a served model
type BuildLogsRequest_SdkV2 struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (newState *BuildLogsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BuildLogsRequest_SdkV2) {
}

func (newState *BuildLogsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState BuildLogsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BuildLogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BuildLogsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BuildLogsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o BuildLogsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              o.Name,
			"served_model_name": o.ServedModelName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BuildLogsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

type BuildLogsResponse_SdkV2 struct {
	// The logs associated with building the served entity's environment.
	Logs types.String `tfsdk:"logs" tf:""`
}

func (newState *BuildLogsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan BuildLogsResponse_SdkV2) {
}

func (newState *BuildLogsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState BuildLogsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BuildLogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a BuildLogsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BuildLogsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o BuildLogsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"logs": o.Logs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BuildLogsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ChatMessage_SdkV2 struct {
	// The content of the message.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The role of the message. One of [system, user, assistant].
	Role types.String `tfsdk:"role" tf:"optional"`
}

func (newState *ChatMessage_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChatMessage_SdkV2) {
}

func (newState *ChatMessage_SdkV2) SyncEffectiveFieldsDuringRead(existingState ChatMessage_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChatMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ChatMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChatMessage_SdkV2
// only implements ToObjectValue() and Type().
func (o ChatMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": o.Content,
			"role":    o.Role,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ChatMessage_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"role":    types.StringType,
		},
	}
}

type CohereConfig_SdkV2 struct {
	// This is an optional field to provide a customized base URL for the Cohere
	// API. If left unspecified, the standard Cohere base URL is used.
	CohereApiBase types.String `tfsdk:"cohere_api_base" tf:"optional"`
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	CohereApiKey types.String `tfsdk:"cohere_api_key" tf:"optional"`
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	CohereApiKeyPlaintext types.String `tfsdk:"cohere_api_key_plaintext" tf:"optional"`
}

func (newState *CohereConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CohereConfig_SdkV2) {
}

func (newState *CohereConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState CohereConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CohereConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CohereConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CohereConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o CohereConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cohere_api_base":          o.CohereApiBase,
			"cohere_api_key":           o.CohereApiKey,
			"cohere_api_key_plaintext": o.CohereApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CohereConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cohere_api_base":          types.StringType,
			"cohere_api_key":           types.StringType,
			"cohere_api_key_plaintext": types.StringType,
		},
	}
}

type CreateServingEndpoint_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: only
	// external model endpoints are supported as of now.
	AiGateway types.List `tfsdk:"ai_gateway" tf:"optional,object"`
	// The core config of the serving endpoint.
	Config types.List `tfsdk:"config" tf:"object"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name types.String `tfsdk:"name" tf:""`
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized types.Bool `tfsdk:"route_optimized" tf:"optional"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateServingEndpoint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServingEndpoint_SdkV2) {
}

func (newState *CreateServingEndpoint_SdkV2) SyncEffectiveFieldsDuringRead(existingState CreateServingEndpoint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServingEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateServingEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":  reflect.TypeOf(AiGatewayConfig_SdkV2{}),
		"config":      reflect.TypeOf(EndpointCoreConfigInput_SdkV2{}),
		"rate_limits": reflect.TypeOf(RateLimit_SdkV2{}),
		"tags":        reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServingEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (o CreateServingEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":      o.AiGateway,
			"config":          o.Config,
			"name":            o.Name,
			"rate_limits":     o.RateLimits,
			"route_optimized": o.RouteOptimized,
			"tags":            o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateServingEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig{}.Type(ctx),
			},
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigInput{}.Type(ctx),
			},
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.Type(ctx),
			},
			"route_optimized": types.BoolType,
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in CreateServingEndpoint_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := o.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in CreateServingEndpoint_SdkV2.
func (o *CreateServingEndpoint_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	o.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in CreateServingEndpoint_SdkV2 as
// a EndpointCoreConfigInput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint_SdkV2) GetConfig(ctx context.Context) (EndpointCoreConfigInput_SdkV2, bool) {
	var e EndpointCoreConfigInput_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []EndpointCoreConfigInput_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in CreateServingEndpoint_SdkV2.
func (o *CreateServingEndpoint_SdkV2) SetConfig(ctx context.Context, v EndpointCoreConfigInput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in CreateServingEndpoint_SdkV2 as
// a slice of RateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint_SdkV2) GetRateLimits(ctx context.Context) ([]RateLimit_SdkV2, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit_SdkV2
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in CreateServingEndpoint_SdkV2.
func (o *CreateServingEndpoint_SdkV2) SetRateLimits(ctx context.Context, v []RateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateServingEndpoint_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateServingEndpoint_SdkV2.
func (o *CreateServingEndpoint_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type DatabricksModelServingConfig_SdkV2 struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiToken types.String `tfsdk:"databricks_api_token" tf:"optional"`
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiTokenPlaintext types.String `tfsdk:"databricks_api_token_plaintext" tf:"optional"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl types.String `tfsdk:"databricks_workspace_url" tf:""`
}

func (newState *DatabricksModelServingConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksModelServingConfig_SdkV2) {
}

func (newState *DatabricksModelServingConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState DatabricksModelServingConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksModelServingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DatabricksModelServingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksModelServingConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o DatabricksModelServingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"databricks_api_token":           o.DatabricksApiToken,
			"databricks_api_token_plaintext": o.DatabricksApiTokenPlaintext,
			"databricks_workspace_url":       o.DatabricksWorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksModelServingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"databricks_api_token":           types.StringType,
			"databricks_api_token_plaintext": types.StringType,
			"databricks_workspace_url":       types.StringType,
		},
	}
}

type DataframeSplitInput_SdkV2 struct {
	Columns types.List `tfsdk:"columns" tf:"optional"`

	Data types.List `tfsdk:"data" tf:"optional"`

	Index types.List `tfsdk:"index" tf:"optional"`
}

func (newState *DataframeSplitInput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataframeSplitInput_SdkV2) {
}

func (newState *DataframeSplitInput_SdkV2) SyncEffectiveFieldsDuringRead(existingState DataframeSplitInput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataframeSplitInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DataframeSplitInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(types.Object{}),
		"data":    reflect.TypeOf(types.Object{}),
		"index":   reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataframeSplitInput_SdkV2
// only implements ToObjectValue() and Type().
func (o DataframeSplitInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
			"data":    o.Data,
			"index":   o.Index,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataframeSplitInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"columns": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"data": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"index": basetypes.ListType{
				ElemType: types.Int64Type,
			},
		},
	}
}

// GetColumns returns the value of the Columns field in DataframeSplitInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *DataframeSplitInput_SdkV2) GetColumns(ctx context.Context) ([]types.Object, bool) {
	if o.Columns.IsNull() || o.Columns.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := o.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in DataframeSplitInput_SdkV2.
func (o *DataframeSplitInput_SdkV2) SetColumns(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

// GetData returns the value of the Data field in DataframeSplitInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *DataframeSplitInput_SdkV2) GetData(ctx context.Context) ([]types.Object, bool) {
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := o.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in DataframeSplitInput_SdkV2.
func (o *DataframeSplitInput_SdkV2) SetData(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Data = types.ListValueMust(t, vs)
}

// GetIndex returns the value of the Index field in DataframeSplitInput_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DataframeSplitInput_SdkV2) GetIndex(ctx context.Context) ([]types.Int64, bool) {
	if o.Index.IsNull() || o.Index.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := o.Index.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIndex sets the value of the Index field in DataframeSplitInput_SdkV2.
func (o *DataframeSplitInput_SdkV2) SetIndex(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["index"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Index = types.ListValueMust(t, vs)
}

type DeleteResponse_SdkV2 struct {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse_SdkV2) {
}

func (newState *DeleteResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteResponse_SdkV2) {
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

// Delete a serving endpoint
type DeleteServingEndpointRequest_SdkV2 struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteServingEndpointRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteServingEndpointRequest_SdkV2) {
}

func (newState *DeleteServingEndpointRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState DeleteServingEndpointRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServingEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteServingEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServingEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o DeleteServingEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServingEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type EmbeddingsV1ResponseEmbeddingElement_SdkV2 struct {
	Embedding types.List `tfsdk:"embedding" tf:"optional"`
	// The index of the embedding in the response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// This will always be 'embedding'.
	Object types.String `tfsdk:"object" tf:"optional"`
}

func (newState *EmbeddingsV1ResponseEmbeddingElement_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingsV1ResponseEmbeddingElement_SdkV2) {
}

func (newState *EmbeddingsV1ResponseEmbeddingElement_SdkV2) SyncEffectiveFieldsDuringRead(existingState EmbeddingsV1ResponseEmbeddingElement_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingsV1ResponseEmbeddingElement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EmbeddingsV1ResponseEmbeddingElement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding": reflect.TypeOf(types.Float64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingsV1ResponseEmbeddingElement_SdkV2
// only implements ToObjectValue() and Type().
func (o EmbeddingsV1ResponseEmbeddingElement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding": o.Embedding,
			"index":     o.Index,
			"object":    o.Object,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmbeddingsV1ResponseEmbeddingElement_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"embedding": basetypes.ListType{
				ElemType: types.Float64Type,
			},
			"index":  types.Int64Type,
			"object": types.StringType,
		},
	}
}

// GetEmbedding returns the value of the Embedding field in EmbeddingsV1ResponseEmbeddingElement_SdkV2 as
// a slice of types.Float64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EmbeddingsV1ResponseEmbeddingElement_SdkV2) GetEmbedding(ctx context.Context) ([]types.Float64, bool) {
	if o.Embedding.IsNull() || o.Embedding.IsUnknown() {
		return nil, false
	}
	var v []types.Float64
	d := o.Embedding.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbedding sets the value of the Embedding field in EmbeddingsV1ResponseEmbeddingElement_SdkV2.
func (o *EmbeddingsV1ResponseEmbeddingElement_SdkV2) SetEmbedding(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Embedding = types.ListValueMust(t, vs)
}

type EndpointCoreConfigInput_SdkV2 struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig types.List `tfsdk:"auto_capture_config" tf:"optional,object"`
	// The name of the serving endpoint to update. This field is required.
	Name types.String `tfsdk:"-"`
	// A list of served entities for the endpoint to serve. A serving endpoint
	// can have up to 15 served entities.
	ServedEntities types.List `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) A list of served models for the
	// endpoint to serve. A serving endpoint can have up to 15 served models.
	ServedModels types.List `tfsdk:"served_models" tf:"optional"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig types.List `tfsdk:"traffic_config" tf:"optional,object"`
}

func (newState *EndpointCoreConfigInput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigInput_SdkV2) {
}

func (newState *EndpointCoreConfigInput_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigInput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointCoreConfigInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_capture_config": reflect.TypeOf(AutoCaptureConfigInput_SdkV2{}),
		"served_entities":     reflect.TypeOf(ServedEntityInput_SdkV2{}),
		"served_models":       reflect.TypeOf(ServedModelInput_SdkV2{}),
		"traffic_config":      reflect.TypeOf(TrafficConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigInput_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointCoreConfigInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_capture_config": o.AutoCaptureConfig,
			"name":                o.Name,
			"served_entities":     o.ServedEntities,
			"served_models":       o.ServedModels,
			"traffic_config":      o.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointCoreConfigInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigInput{}.Type(ctx),
			},
			"name": types.StringType,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityInput{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelInput{}.Type(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig{}.Type(ctx),
			},
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointCoreConfigInput_SdkV2 as
// a AutoCaptureConfigInput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput_SdkV2) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigInput_SdkV2, bool) {
	var e AutoCaptureConfigInput_SdkV2
	if o.AutoCaptureConfig.IsNull() || o.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureConfigInput_SdkV2
	d := o.AutoCaptureConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointCoreConfigInput_SdkV2.
func (o *EndpointCoreConfigInput_SdkV2) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigInput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_capture_config"]
	o.AutoCaptureConfig = types.ListValueMust(t, vs)
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigInput_SdkV2 as
// a slice of ServedEntityInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntityInput_SdkV2, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityInput_SdkV2
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigInput_SdkV2.
func (o *EndpointCoreConfigInput_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntityInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigInput_SdkV2 as
// a slice of ServedModelInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelInput_SdkV2, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelInput_SdkV2
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigInput_SdkV2.
func (o *EndpointCoreConfigInput_SdkV2) SetServedModels(ctx context.Context, v []ServedModelInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointCoreConfigInput_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := o.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointCoreConfigInput_SdkV2.
func (o *EndpointCoreConfigInput_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	o.TrafficConfig = types.ListValueMust(t, vs)
}

type EndpointCoreConfigOutput_SdkV2 struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig types.List `tfsdk:"auto_capture_config" tf:"optional,object"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version" tf:"optional"`
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models" tf:"optional"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig types.List `tfsdk:"traffic_config" tf:"optional,object"`
}

func (newState *EndpointCoreConfigOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigOutput_SdkV2) {
}

func (newState *EndpointCoreConfigOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointCoreConfigOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_capture_config": reflect.TypeOf(AutoCaptureConfigOutput_SdkV2{}),
		"served_entities":     reflect.TypeOf(ServedEntityOutput_SdkV2{}),
		"served_models":       reflect.TypeOf(ServedModelOutput_SdkV2{}),
		"traffic_config":      reflect.TypeOf(TrafficConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointCoreConfigOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_capture_config": o.AutoCaptureConfig,
			"config_version":      o.ConfigVersion,
			"served_entities":     o.ServedEntities,
			"served_models":       o.ServedModels,
			"traffic_config":      o.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointCoreConfigOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigOutput{}.Type(ctx),
			},
			"config_version": types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput{}.Type(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig{}.Type(ctx),
			},
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointCoreConfigOutput_SdkV2 as
// a AutoCaptureConfigOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput_SdkV2) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigOutput_SdkV2, bool) {
	var e AutoCaptureConfigOutput_SdkV2
	if o.AutoCaptureConfig.IsNull() || o.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureConfigOutput_SdkV2
	d := o.AutoCaptureConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointCoreConfigOutput_SdkV2.
func (o *EndpointCoreConfigOutput_SdkV2) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_capture_config"]
	o.AutoCaptureConfig = types.ListValueMust(t, vs)
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigOutput_SdkV2 as
// a slice of ServedEntityOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntityOutput_SdkV2, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityOutput_SdkV2
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigOutput_SdkV2.
func (o *EndpointCoreConfigOutput_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntityOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigOutput_SdkV2 as
// a slice of ServedModelOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelOutput_SdkV2, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelOutput_SdkV2
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigOutput_SdkV2.
func (o *EndpointCoreConfigOutput_SdkV2) SetServedModels(ctx context.Context, v []ServedModelOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointCoreConfigOutput_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := o.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointCoreConfigOutput_SdkV2.
func (o *EndpointCoreConfigOutput_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	o.TrafficConfig = types.ListValueMust(t, vs)
}

type EndpointCoreConfigSummary_SdkV2 struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models" tf:"optional"`
}

func (newState *EndpointCoreConfigSummary_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigSummary_SdkV2) {
}

func (newState *EndpointCoreConfigSummary_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigSummary_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointCoreConfigSummary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"served_entities": reflect.TypeOf(ServedEntitySpec_SdkV2{}),
		"served_models":   reflect.TypeOf(ServedModelSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigSummary_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointCoreConfigSummary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entities": o.ServedEntities,
			"served_models":   o.ServedModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointCoreConfigSummary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entities": basetypes.ListType{
				ElemType: ServedEntitySpec{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelSpec{}.Type(ctx),
			},
		},
	}
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigSummary_SdkV2 as
// a slice of ServedEntitySpec_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigSummary_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntitySpec_SdkV2, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntitySpec_SdkV2
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigSummary_SdkV2.
func (o *EndpointCoreConfigSummary_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntitySpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigSummary_SdkV2 as
// a slice of ServedModelSpec_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigSummary_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelSpec_SdkV2, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelSpec_SdkV2
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigSummary_SdkV2.
func (o *EndpointCoreConfigSummary_SdkV2) SetServedModels(ctx context.Context, v []ServedModelSpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

type EndpointPendingConfig_SdkV2 struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig types.List `tfsdk:"auto_capture_config" tf:"optional,object"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version" tf:"optional"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities types.List `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels types.List `tfsdk:"served_models" tf:"optional"`
	// The timestamp when the update to the pending config started.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig types.List `tfsdk:"traffic_config" tf:"optional,object"`
}

func (newState *EndpointPendingConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointPendingConfig_SdkV2) {
}

func (newState *EndpointPendingConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointPendingConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointPendingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointPendingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_capture_config": reflect.TypeOf(AutoCaptureConfigOutput_SdkV2{}),
		"served_entities":     reflect.TypeOf(ServedEntityOutput_SdkV2{}),
		"served_models":       reflect.TypeOf(ServedModelOutput_SdkV2{}),
		"traffic_config":      reflect.TypeOf(TrafficConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointPendingConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointPendingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_capture_config": o.AutoCaptureConfig,
			"config_version":      o.ConfigVersion,
			"served_entities":     o.ServedEntities,
			"served_models":       o.ServedModels,
			"start_time":          o.StartTime,
			"traffic_config":      o.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointPendingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigOutput{}.Type(ctx),
			},
			"config_version": types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput{}.Type(ctx),
			},
			"start_time": types.Int64Type,
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig{}.Type(ctx),
			},
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointPendingConfig_SdkV2 as
// a AutoCaptureConfigOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig_SdkV2) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigOutput_SdkV2, bool) {
	var e AutoCaptureConfigOutput_SdkV2
	if o.AutoCaptureConfig.IsNull() || o.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureConfigOutput_SdkV2
	d := o.AutoCaptureConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointPendingConfig_SdkV2.
func (o *EndpointPendingConfig_SdkV2) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_capture_config"]
	o.AutoCaptureConfig = types.ListValueMust(t, vs)
}

// GetServedEntities returns the value of the ServedEntities field in EndpointPendingConfig_SdkV2 as
// a slice of ServedEntityOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntityOutput_SdkV2, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityOutput_SdkV2
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointPendingConfig_SdkV2.
func (o *EndpointPendingConfig_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntityOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointPendingConfig_SdkV2 as
// a slice of ServedModelOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelOutput_SdkV2, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelOutput_SdkV2
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointPendingConfig_SdkV2.
func (o *EndpointPendingConfig_SdkV2) SetServedModels(ctx context.Context, v []ServedModelOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointPendingConfig_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := o.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointPendingConfig_SdkV2.
func (o *EndpointPendingConfig_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	o.TrafficConfig = types.ListValueMust(t, vs)
}

type EndpointState_SdkV2 struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails."
	ConfigUpdate types.String `tfsdk:"config_update" tf:"optional"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready types.String `tfsdk:"ready" tf:"optional"`
}

func (newState *EndpointState_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointState_SdkV2) {
}

func (newState *EndpointState_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointState_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointState_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_update": o.ConfigUpdate,
			"ready":         o.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointState_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_update": types.StringType,
			"ready":         types.StringType,
		},
	}
}

type EndpointTag_SdkV2 struct {
	// Key field for a serving endpoint tag.
	Key types.String `tfsdk:"key" tf:""`
	// Optional value field for a serving endpoint tag.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *EndpointTag_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTag_SdkV2) {
}

func (newState *EndpointTag_SdkV2) SyncEffectiveFieldsDuringRead(existingState EndpointTag_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTag_SdkV2
// only implements ToObjectValue() and Type().
func (o EndpointTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Get metrics of a serving endpoint
type ExportMetricsRequest_SdkV2 struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (newState *ExportMetricsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportMetricsRequest_SdkV2) {
}

func (newState *ExportMetricsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExportMetricsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportMetricsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportMetricsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportMetricsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ExportMetricsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportMetricsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ExportMetricsResponse_SdkV2 struct {
	Contents types.Object `tfsdk:"-"`
}

func (newState *ExportMetricsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportMetricsResponse_SdkV2) {
}

func (newState *ExportMetricsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExportMetricsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportMetricsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExportMetricsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportMetricsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ExportMetricsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": o.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportMetricsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

type ExternalModel_SdkV2 struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig types.List `tfsdk:"ai21labs_config" tf:"optional,object"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig types.List `tfsdk:"amazon_bedrock_config" tf:"optional,object"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig types.List `tfsdk:"anthropic_config" tf:"optional,object"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig types.List `tfsdk:"cohere_config" tf:"optional,object"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig types.List `tfsdk:"databricks_model_serving_config" tf:"optional,object"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig types.List `tfsdk:"google_cloud_vertex_ai_config" tf:"optional,object"`
	// The name of the external model.
	Name types.String `tfsdk:"name" tf:""`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig types.List `tfsdk:"openai_config" tf:"optional,object"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig types.List `tfsdk:"palm_config" tf:"optional,object"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', and
	// 'palm'.",
	Provider types.String `tfsdk:"provider" tf:""`
	// The task type of the external model.
	Task types.String `tfsdk:"task" tf:""`
}

func (newState *ExternalModel_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalModel_SdkV2) {
}

func (newState *ExternalModel_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalModel_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai21labs_config":                 reflect.TypeOf(Ai21LabsConfig_SdkV2{}),
		"amazon_bedrock_config":           reflect.TypeOf(AmazonBedrockConfig_SdkV2{}),
		"anthropic_config":                reflect.TypeOf(AnthropicConfig_SdkV2{}),
		"cohere_config":                   reflect.TypeOf(CohereConfig_SdkV2{}),
		"databricks_model_serving_config": reflect.TypeOf(DatabricksModelServingConfig_SdkV2{}),
		"google_cloud_vertex_ai_config":   reflect.TypeOf(GoogleCloudVertexAiConfig_SdkV2{}),
		"openai_config":                   reflect.TypeOf(OpenAiConfig_SdkV2{}),
		"palm_config":                     reflect.TypeOf(PaLmConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalModel_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai21labs_config":                 o.Ai21labsConfig,
			"amazon_bedrock_config":           o.AmazonBedrockConfig,
			"anthropic_config":                o.AnthropicConfig,
			"cohere_config":                   o.CohereConfig,
			"databricks_model_serving_config": o.DatabricksModelServingConfig,
			"google_cloud_vertex_ai_config":   o.GoogleCloudVertexAiConfig,
			"name":                            o.Name,
			"openai_config":                   o.OpenaiConfig,
			"palm_config":                     o.PalmConfig,
			"provider":                        o.Provider,
			"task":                            o.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_config": basetypes.ListType{
				ElemType: Ai21LabsConfig{}.Type(ctx),
			},
			"amazon_bedrock_config": basetypes.ListType{
				ElemType: AmazonBedrockConfig{}.Type(ctx),
			},
			"anthropic_config": basetypes.ListType{
				ElemType: AnthropicConfig{}.Type(ctx),
			},
			"cohere_config": basetypes.ListType{
				ElemType: CohereConfig{}.Type(ctx),
			},
			"databricks_model_serving_config": basetypes.ListType{
				ElemType: DatabricksModelServingConfig{}.Type(ctx),
			},
			"google_cloud_vertex_ai_config": basetypes.ListType{
				ElemType: GoogleCloudVertexAiConfig{}.Type(ctx),
			},
			"name": types.StringType,
			"openai_config": basetypes.ListType{
				ElemType: OpenAiConfig{}.Type(ctx),
			},
			"palm_config": basetypes.ListType{
				ElemType: PaLmConfig{}.Type(ctx),
			},
			"provider": types.StringType,
			"task":     types.StringType,
		},
	}
}

// GetAi21labsConfig returns the value of the Ai21labsConfig field in ExternalModel_SdkV2 as
// a Ai21LabsConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetAi21labsConfig(ctx context.Context) (Ai21LabsConfig_SdkV2, bool) {
	var e Ai21LabsConfig_SdkV2
	if o.Ai21labsConfig.IsNull() || o.Ai21labsConfig.IsUnknown() {
		return e, false
	}
	var v []Ai21LabsConfig_SdkV2
	d := o.Ai21labsConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAi21labsConfig sets the value of the Ai21labsConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetAi21labsConfig(ctx context.Context, v Ai21LabsConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ai21labs_config"]
	o.Ai21labsConfig = types.ListValueMust(t, vs)
}

// GetAmazonBedrockConfig returns the value of the AmazonBedrockConfig field in ExternalModel_SdkV2 as
// a AmazonBedrockConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetAmazonBedrockConfig(ctx context.Context) (AmazonBedrockConfig_SdkV2, bool) {
	var e AmazonBedrockConfig_SdkV2
	if o.AmazonBedrockConfig.IsNull() || o.AmazonBedrockConfig.IsUnknown() {
		return e, false
	}
	var v []AmazonBedrockConfig_SdkV2
	d := o.AmazonBedrockConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAmazonBedrockConfig sets the value of the AmazonBedrockConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetAmazonBedrockConfig(ctx context.Context, v AmazonBedrockConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["amazon_bedrock_config"]
	o.AmazonBedrockConfig = types.ListValueMust(t, vs)
}

// GetAnthropicConfig returns the value of the AnthropicConfig field in ExternalModel_SdkV2 as
// a AnthropicConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetAnthropicConfig(ctx context.Context) (AnthropicConfig_SdkV2, bool) {
	var e AnthropicConfig_SdkV2
	if o.AnthropicConfig.IsNull() || o.AnthropicConfig.IsUnknown() {
		return e, false
	}
	var v []AnthropicConfig_SdkV2
	d := o.AnthropicConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAnthropicConfig sets the value of the AnthropicConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetAnthropicConfig(ctx context.Context, v AnthropicConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["anthropic_config"]
	o.AnthropicConfig = types.ListValueMust(t, vs)
}

// GetCohereConfig returns the value of the CohereConfig field in ExternalModel_SdkV2 as
// a CohereConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetCohereConfig(ctx context.Context) (CohereConfig_SdkV2, bool) {
	var e CohereConfig_SdkV2
	if o.CohereConfig.IsNull() || o.CohereConfig.IsUnknown() {
		return e, false
	}
	var v []CohereConfig_SdkV2
	d := o.CohereConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCohereConfig sets the value of the CohereConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetCohereConfig(ctx context.Context, v CohereConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["cohere_config"]
	o.CohereConfig = types.ListValueMust(t, vs)
}

// GetDatabricksModelServingConfig returns the value of the DatabricksModelServingConfig field in ExternalModel_SdkV2 as
// a DatabricksModelServingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetDatabricksModelServingConfig(ctx context.Context) (DatabricksModelServingConfig_SdkV2, bool) {
	var e DatabricksModelServingConfig_SdkV2
	if o.DatabricksModelServingConfig.IsNull() || o.DatabricksModelServingConfig.IsUnknown() {
		return e, false
	}
	var v []DatabricksModelServingConfig_SdkV2
	d := o.DatabricksModelServingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksModelServingConfig sets the value of the DatabricksModelServingConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetDatabricksModelServingConfig(ctx context.Context, v DatabricksModelServingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_model_serving_config"]
	o.DatabricksModelServingConfig = types.ListValueMust(t, vs)
}

// GetGoogleCloudVertexAiConfig returns the value of the GoogleCloudVertexAiConfig field in ExternalModel_SdkV2 as
// a GoogleCloudVertexAiConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetGoogleCloudVertexAiConfig(ctx context.Context) (GoogleCloudVertexAiConfig_SdkV2, bool) {
	var e GoogleCloudVertexAiConfig_SdkV2
	if o.GoogleCloudVertexAiConfig.IsNull() || o.GoogleCloudVertexAiConfig.IsUnknown() {
		return e, false
	}
	var v []GoogleCloudVertexAiConfig_SdkV2
	d := o.GoogleCloudVertexAiConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGoogleCloudVertexAiConfig sets the value of the GoogleCloudVertexAiConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetGoogleCloudVertexAiConfig(ctx context.Context, v GoogleCloudVertexAiConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["google_cloud_vertex_ai_config"]
	o.GoogleCloudVertexAiConfig = types.ListValueMust(t, vs)
}

// GetOpenaiConfig returns the value of the OpenaiConfig field in ExternalModel_SdkV2 as
// a OpenAiConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetOpenaiConfig(ctx context.Context) (OpenAiConfig_SdkV2, bool) {
	var e OpenAiConfig_SdkV2
	if o.OpenaiConfig.IsNull() || o.OpenaiConfig.IsUnknown() {
		return e, false
	}
	var v []OpenAiConfig_SdkV2
	d := o.OpenaiConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOpenaiConfig sets the value of the OpenaiConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetOpenaiConfig(ctx context.Context, v OpenAiConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["openai_config"]
	o.OpenaiConfig = types.ListValueMust(t, vs)
}

// GetPalmConfig returns the value of the PalmConfig field in ExternalModel_SdkV2 as
// a PaLmConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel_SdkV2) GetPalmConfig(ctx context.Context) (PaLmConfig_SdkV2, bool) {
	var e PaLmConfig_SdkV2
	if o.PalmConfig.IsNull() || o.PalmConfig.IsUnknown() {
		return e, false
	}
	var v []PaLmConfig_SdkV2
	d := o.PalmConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPalmConfig sets the value of the PalmConfig field in ExternalModel_SdkV2.
func (o *ExternalModel_SdkV2) SetPalmConfig(ctx context.Context, v PaLmConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["palm_config"]
	o.PalmConfig = types.ListValueMust(t, vs)
}

type ExternalModelUsageElement_SdkV2 struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens types.Int64 `tfsdk:"completion_tokens" tf:"optional"`
	// The number of tokens in the prompt.
	PromptTokens types.Int64 `tfsdk:"prompt_tokens" tf:"optional"`
	// The total number of tokens in the prompt and response.
	TotalTokens types.Int64 `tfsdk:"total_tokens" tf:"optional"`
}

func (newState *ExternalModelUsageElement_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalModelUsageElement_SdkV2) {
}

func (newState *ExternalModelUsageElement_SdkV2) SyncEffectiveFieldsDuringRead(existingState ExternalModelUsageElement_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalModelUsageElement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ExternalModelUsageElement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalModelUsageElement_SdkV2
// only implements ToObjectValue() and Type().
func (o ExternalModelUsageElement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"completion_tokens": o.CompletionTokens,
			"prompt_tokens":     o.PromptTokens,
			"total_tokens":      o.TotalTokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalModelUsageElement_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"completion_tokens": types.Int64Type,
			"prompt_tokens":     types.Int64Type,
			"total_tokens":      types.Int64Type,
		},
	}
}

type FoundationModel_SdkV2 struct {
	// The description of the foundation model.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The display name of the foundation model.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The URL to the documentation of the foundation model.
	Docs types.String `tfsdk:"docs" tf:"optional"`
	// The name of the foundation model.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *FoundationModel_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan FoundationModel_SdkV2) {
}

func (newState *FoundationModel_SdkV2) SyncEffectiveFieldsDuringRead(existingState FoundationModel_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FoundationModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a FoundationModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FoundationModel_SdkV2
// only implements ToObjectValue() and Type().
func (o FoundationModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":  o.Description,
			"display_name": o.DisplayName,
			"docs":         o.Docs,
			"name":         o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FoundationModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":  types.StringType,
			"display_name": types.StringType,
			"docs":         types.StringType,
			"name":         types.StringType,
		},
	}
}

// Get the schema for a serving endpoint
type GetOpenApiRequest_SdkV2 struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *GetOpenApiRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOpenApiRequest_SdkV2) {
}

func (newState *GetOpenApiRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetOpenApiRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOpenApiRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetOpenApiRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOpenApiRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetOpenApiRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOpenApiRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// The response is an OpenAPI spec in JSON format that typically includes fields
// like openapi, info, servers and paths, etc.
type GetOpenApiResponse_SdkV2 struct {
}

func (newState *GetOpenApiResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOpenApiResponse_SdkV2) {
}

func (newState *GetOpenApiResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetOpenApiResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOpenApiResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetOpenApiResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOpenApiResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetOpenApiResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o GetOpenApiResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest_SdkV2 struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointPermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionLevelsRequest_SdkV2) {
}

func (newState *GetServingEndpointPermissionLevelsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionLevelsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServingEndpointPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetServingEndpointPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"serving_endpoint_id": o.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

type GetServingEndpointPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetServingEndpointPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionLevelsResponse_SdkV2) {
}

func (newState *GetServingEndpointPermissionLevelsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionLevelsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServingEndpointPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ServingEndpointPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o GetServingEndpointPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ServingEndpointPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetServingEndpointPermissionLevelsResponse_SdkV2 as
// a slice of ServingEndpointPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetServingEndpointPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]ServingEndpointPermissionsDescription_SdkV2, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointPermissionsDescription_SdkV2
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetServingEndpointPermissionLevelsResponse_SdkV2.
func (o *GetServingEndpointPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []ServingEndpointPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest_SdkV2 struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionsRequest_SdkV2) {
}

func (newState *GetServingEndpointPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServingEndpointPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetServingEndpointPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"serving_endpoint_id": o.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

// Get a single serving endpoint
type GetServingEndpointRequest_SdkV2 struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointRequest_SdkV2) {
}

func (newState *GetServingEndpointRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetServingEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o GetServingEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GoogleCloudVertexAiConfig_SdkV2 struct {
	// The Databricks secret key reference for a private key for the service
	// account which has access to the Google Cloud Vertex AI Service. See [Best
	// practices for managing service account keys]. If you prefer to paste your
	// API key directly, see `private_key_plaintext`. You must provide an API
	// key using one of the following fields: `private_key` or
	// `private_key_plaintext`
	//
	// [Best practices for managing service account keys]: https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	PrivateKey types.String `tfsdk:"private_key" tf:"optional"`
	// The private key for the service account which has access to the Google
	// Cloud Vertex AI Service provided as a plaintext secret. See [Best
	// practices for managing service account keys]. If you prefer to reference
	// your key using Databricks Secrets, see `private_key`. You must provide an
	// API key using one of the following fields: `private_key` or
	// `private_key_plaintext`.
	//
	// [Best practices for managing service account keys]: https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	PrivateKeyPlaintext types.String `tfsdk:"private_key_plaintext" tf:"optional"`
	// This is the Google Cloud project id that the service account is
	// associated with.
	ProjectId types.String `tfsdk:"project_id" tf:"optional"`
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]: https://cloud.google.com/vertex-ai/docs/general/locations
	Region types.String `tfsdk:"region" tf:"optional"`
}

func (newState *GoogleCloudVertexAiConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan GoogleCloudVertexAiConfig_SdkV2) {
}

func (newState *GoogleCloudVertexAiConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState GoogleCloudVertexAiConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GoogleCloudVertexAiConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GoogleCloudVertexAiConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GoogleCloudVertexAiConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o GoogleCloudVertexAiConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_key":           o.PrivateKey,
			"private_key_plaintext": o.PrivateKeyPlaintext,
			"project_id":            o.ProjectId,
			"region":                o.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GoogleCloudVertexAiConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_key":           types.StringType,
			"private_key_plaintext": types.StringType,
			"project_id":            types.StringType,
			"region":                types.StringType,
		},
	}
}

type ListEndpointsResponse_SdkV2 struct {
	// The list of endpoints.
	Endpoints types.List `tfsdk:"endpoints" tf:"optional"`
}

func (newState *ListEndpointsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListEndpointsResponse_SdkV2) {
}

func (newState *ListEndpointsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ListEndpointsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListEndpointsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(ServingEndpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ListEndpointsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints": o.Endpoints,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListEndpointsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: ServingEndpoint{}.Type(ctx),
			},
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse_SdkV2 as
// a slice of ServingEndpoint_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListEndpointsResponse_SdkV2) GetEndpoints(ctx context.Context) ([]ServingEndpoint_SdkV2, bool) {
	if o.Endpoints.IsNull() || o.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpoint_SdkV2
	d := o.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse_SdkV2.
func (o *ListEndpointsResponse_SdkV2) SetEndpoints(ctx context.Context, v []ServingEndpoint_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Endpoints = types.ListValueMust(t, vs)
}

// Get the latest logs for a served model
type LogsRequest_SdkV2 struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (newState *LogsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogsRequest_SdkV2) {
}

func (newState *LogsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState LogsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a LogsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o LogsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              o.Name,
			"served_model_name": o.ServedModelName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

type ModelDataPlaneInfo_SdkV2 struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo types.List `tfsdk:"query_info" tf:"optional,object"`
}

func (newState *ModelDataPlaneInfo_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelDataPlaneInfo_SdkV2) {
}

func (newState *ModelDataPlaneInfo_SdkV2) SyncEffectiveFieldsDuringRead(existingState ModelDataPlaneInfo_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelDataPlaneInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ModelDataPlaneInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_info": reflect.TypeOf(oauth2_tf.DataPlaneInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDataPlaneInfo_SdkV2
// only implements ToObjectValue() and Type().
func (o ModelDataPlaneInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_info": o.QueryInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelDataPlaneInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_info": basetypes.ListType{
				ElemType: oauth2_tf.DataPlaneInfo{}.Type(ctx),
			},
		},
	}
}

// GetQueryInfo returns the value of the QueryInfo field in ModelDataPlaneInfo_SdkV2 as
// a oauth2_tf.DataPlaneInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelDataPlaneInfo_SdkV2) GetQueryInfo(ctx context.Context) (oauth2_tf.DataPlaneInfo_SdkV2, bool) {
	var e oauth2_tf.DataPlaneInfo_SdkV2
	if o.QueryInfo.IsNull() || o.QueryInfo.IsUnknown() {
		return e, false
	}
	var v []oauth2_tf.DataPlaneInfo_SdkV2
	d := o.QueryInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryInfo sets the value of the QueryInfo field in ModelDataPlaneInfo_SdkV2.
func (o *ModelDataPlaneInfo_SdkV2) SetQueryInfo(ctx context.Context, v oauth2_tf.DataPlaneInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["query_info"]
	o.QueryInfo = types.ListValueMust(t, vs)
}

type OpenAiConfig_SdkV2 struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	MicrosoftEntraClientId types.String `tfsdk:"microsoft_entra_client_id" tf:"optional"`
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecret types.String `tfsdk:"microsoft_entra_client_secret" tf:"optional"`
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecretPlaintext types.String `tfsdk:"microsoft_entra_client_secret_plaintext" tf:"optional"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId types.String `tfsdk:"microsoft_entra_tenant_id" tf:"optional"`
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	OpenaiApiBase types.String `tfsdk:"openai_api_base" tf:"optional"`
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKey types.String `tfsdk:"openai_api_key" tf:"optional"`
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKeyPlaintext types.String `tfsdk:"openai_api_key_plaintext" tf:"optional"`
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	OpenaiApiType types.String `tfsdk:"openai_api_type" tf:"optional"`
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	OpenaiApiVersion types.String `tfsdk:"openai_api_version" tf:"optional"`
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	OpenaiDeploymentName types.String `tfsdk:"openai_deployment_name" tf:"optional"`
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	OpenaiOrganization types.String `tfsdk:"openai_organization" tf:"optional"`
}

func (newState *OpenAiConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan OpenAiConfig_SdkV2) {
}

func (newState *OpenAiConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState OpenAiConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OpenAiConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a OpenAiConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OpenAiConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o OpenAiConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"microsoft_entra_client_id":               o.MicrosoftEntraClientId,
			"microsoft_entra_client_secret":           o.MicrosoftEntraClientSecret,
			"microsoft_entra_client_secret_plaintext": o.MicrosoftEntraClientSecretPlaintext,
			"microsoft_entra_tenant_id":               o.MicrosoftEntraTenantId,
			"openai_api_base":                         o.OpenaiApiBase,
			"openai_api_key":                          o.OpenaiApiKey,
			"openai_api_key_plaintext":                o.OpenaiApiKeyPlaintext,
			"openai_api_type":                         o.OpenaiApiType,
			"openai_api_version":                      o.OpenaiApiVersion,
			"openai_deployment_name":                  o.OpenaiDeploymentName,
			"openai_organization":                     o.OpenaiOrganization,
		})
}

// Type implements basetypes.ObjectValuable.
func (o OpenAiConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"microsoft_entra_client_id":               types.StringType,
			"microsoft_entra_client_secret":           types.StringType,
			"microsoft_entra_client_secret_plaintext": types.StringType,
			"microsoft_entra_tenant_id":               types.StringType,
			"openai_api_base":                         types.StringType,
			"openai_api_key":                          types.StringType,
			"openai_api_key_plaintext":                types.StringType,
			"openai_api_type":                         types.StringType,
			"openai_api_version":                      types.StringType,
			"openai_deployment_name":                  types.StringType,
			"openai_organization":                     types.StringType,
		},
	}
}

type PaLmConfig_SdkV2 struct {
	// The Databricks secret key reference for a PaLM API key. If you prefer to
	// paste your API key directly, see `palm_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKey types.String `tfsdk:"palm_api_key" tf:"optional"`
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKeyPlaintext types.String `tfsdk:"palm_api_key_plaintext" tf:"optional"`
}

func (newState *PaLmConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PaLmConfig_SdkV2) {
}

func (newState *PaLmConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState PaLmConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PaLmConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PaLmConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PaLmConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o PaLmConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"palm_api_key":           o.PalmApiKey,
			"palm_api_key_plaintext": o.PalmApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PaLmConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"palm_api_key":           types.StringType,
			"palm_api_key_plaintext": types.StringType,
		},
	}
}

type PatchServingEndpointTags_SdkV2 struct {
	// List of endpoint tags to add
	AddTags types.List `tfsdk:"add_tags" tf:"optional"`
	// List of tag keys to delete
	DeleteTags types.List `tfsdk:"delete_tags" tf:"optional"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (newState *PatchServingEndpointTags_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PatchServingEndpointTags_SdkV2) {
}

func (newState *PatchServingEndpointTags_SdkV2) SyncEffectiveFieldsDuringRead(existingState PatchServingEndpointTags_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchServingEndpointTags.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PatchServingEndpointTags_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add_tags":    reflect.TypeOf(EndpointTag_SdkV2{}),
		"delete_tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchServingEndpointTags_SdkV2
// only implements ToObjectValue() and Type().
func (o PatchServingEndpointTags_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add_tags":    o.AddTags,
			"delete_tags": o.DeleteTags,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchServingEndpointTags_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add_tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
			"delete_tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name": types.StringType,
		},
	}
}

// GetAddTags returns the value of the AddTags field in PatchServingEndpointTags_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchServingEndpointTags_SdkV2) GetAddTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if o.AddTags.IsNull() || o.AddTags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := o.AddTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAddTags sets the value of the AddTags field in PatchServingEndpointTags_SdkV2.
func (o *PatchServingEndpointTags_SdkV2) SetAddTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["add_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AddTags = types.ListValueMust(t, vs)
}

// GetDeleteTags returns the value of the DeleteTags field in PatchServingEndpointTags_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchServingEndpointTags_SdkV2) GetDeleteTags(ctx context.Context) ([]types.String, bool) {
	if o.DeleteTags.IsNull() || o.DeleteTags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.DeleteTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeleteTags sets the value of the DeleteTags field in PatchServingEndpointTags_SdkV2.
func (o *PatchServingEndpointTags_SdkV2) SetDeleteTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["delete_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DeleteTags = types.ListValueMust(t, vs)
}

type PayloadTable_SdkV2 struct {
	// The name of the payload table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The status of the payload table.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The status message of the payload table.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
}

func (newState *PayloadTable_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PayloadTable_SdkV2) {
}

func (newState *PayloadTable_SdkV2) SyncEffectiveFieldsDuringRead(existingState PayloadTable_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PayloadTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PayloadTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PayloadTable_SdkV2
// only implements ToObjectValue() and Type().
func (o PayloadTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":           o.Name,
			"status":         o.Status,
			"status_message": o.StatusMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PayloadTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":           types.StringType,
			"status":         types.StringType,
			"status_message": types.StringType,
		},
	}
}

// Update AI Gateway of a serving endpoint
type PutAiGatewayRequest_SdkV2 struct {
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.List `tfsdk:"guardrails" tf:"optional,object"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.List `tfsdk:"inference_table_config" tf:"optional,object"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.List `tfsdk:"usage_tracking_config" tf:"optional,object"`
}

func (newState *PutAiGatewayRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAiGatewayRequest_SdkV2) {
}

func (newState *PutAiGatewayRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PutAiGatewayRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAiGatewayRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutAiGatewayRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails_SdkV2{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig_SdkV2{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit_SdkV2{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAiGatewayRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PutAiGatewayRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"guardrails":             o.Guardrails,
			"inference_table_config": o.InferenceTableConfig,
			"name":                   o.Name,
			"rate_limits":            o.RateLimits,
			"usage_tracking_config":  o.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutAiGatewayRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails{}.Type(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig{}.Type(ctx),
			},
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.Type(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig{}.Type(ctx),
			},
		},
	}
}

// GetGuardrails returns the value of the Guardrails field in PutAiGatewayRequest_SdkV2 as
// a AiGatewayGuardrails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest_SdkV2) GetGuardrails(ctx context.Context) (AiGatewayGuardrails_SdkV2, bool) {
	var e AiGatewayGuardrails_SdkV2
	if o.Guardrails.IsNull() || o.Guardrails.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrails_SdkV2
	d := o.Guardrails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGuardrails sets the value of the Guardrails field in PutAiGatewayRequest_SdkV2.
func (o *PutAiGatewayRequest_SdkV2) SetGuardrails(ctx context.Context, v AiGatewayGuardrails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guardrails"]
	o.Guardrails = types.ListValueMust(t, vs)
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in PutAiGatewayRequest_SdkV2 as
// a AiGatewayInferenceTableConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest_SdkV2) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig_SdkV2, bool) {
	var e AiGatewayInferenceTableConfig_SdkV2
	if o.InferenceTableConfig.IsNull() || o.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayInferenceTableConfig_SdkV2
	d := o.InferenceTableConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in PutAiGatewayRequest_SdkV2.
func (o *PutAiGatewayRequest_SdkV2) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_table_config"]
	o.InferenceTableConfig = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in PutAiGatewayRequest_SdkV2 as
// a slice of AiGatewayRateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest_SdkV2) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit_SdkV2, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit_SdkV2
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutAiGatewayRequest_SdkV2.
func (o *PutAiGatewayRequest_SdkV2) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in PutAiGatewayRequest_SdkV2 as
// a AiGatewayUsageTrackingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest_SdkV2) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig_SdkV2, bool) {
	var e AiGatewayUsageTrackingConfig_SdkV2
	if o.UsageTrackingConfig.IsNull() || o.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayUsageTrackingConfig_SdkV2
	d := o.UsageTrackingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in PutAiGatewayRequest_SdkV2.
func (o *PutAiGatewayRequest_SdkV2) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_tracking_config"]
	o.UsageTrackingConfig = types.ListValueMust(t, vs)
}

type PutAiGatewayResponse_SdkV2 struct {
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.List `tfsdk:"guardrails" tf:"optional,object"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality .
	InferenceTableConfig types.List `tfsdk:"inference_table_config" tf:"optional,object"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.List `tfsdk:"usage_tracking_config" tf:"optional,object"`
}

func (newState *PutAiGatewayResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAiGatewayResponse_SdkV2) {
}

func (newState *PutAiGatewayResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState PutAiGatewayResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAiGatewayResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutAiGatewayResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails_SdkV2{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig_SdkV2{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit_SdkV2{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAiGatewayResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PutAiGatewayResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"guardrails":             o.Guardrails,
			"inference_table_config": o.InferenceTableConfig,
			"rate_limits":            o.RateLimits,
			"usage_tracking_config":  o.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutAiGatewayResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails{}.Type(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig{}.Type(ctx),
			},
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.Type(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig{}.Type(ctx),
			},
		},
	}
}

// GetGuardrails returns the value of the Guardrails field in PutAiGatewayResponse_SdkV2 as
// a AiGatewayGuardrails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse_SdkV2) GetGuardrails(ctx context.Context) (AiGatewayGuardrails_SdkV2, bool) {
	var e AiGatewayGuardrails_SdkV2
	if o.Guardrails.IsNull() || o.Guardrails.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrails_SdkV2
	d := o.Guardrails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGuardrails sets the value of the Guardrails field in PutAiGatewayResponse_SdkV2.
func (o *PutAiGatewayResponse_SdkV2) SetGuardrails(ctx context.Context, v AiGatewayGuardrails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["guardrails"]
	o.Guardrails = types.ListValueMust(t, vs)
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in PutAiGatewayResponse_SdkV2 as
// a AiGatewayInferenceTableConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse_SdkV2) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig_SdkV2, bool) {
	var e AiGatewayInferenceTableConfig_SdkV2
	if o.InferenceTableConfig.IsNull() || o.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayInferenceTableConfig_SdkV2
	d := o.InferenceTableConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in PutAiGatewayResponse_SdkV2.
func (o *PutAiGatewayResponse_SdkV2) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_table_config"]
	o.InferenceTableConfig = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in PutAiGatewayResponse_SdkV2 as
// a slice of AiGatewayRateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse_SdkV2) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit_SdkV2, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit_SdkV2
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutAiGatewayResponse_SdkV2.
func (o *PutAiGatewayResponse_SdkV2) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in PutAiGatewayResponse_SdkV2 as
// a AiGatewayUsageTrackingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse_SdkV2) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig_SdkV2, bool) {
	var e AiGatewayUsageTrackingConfig_SdkV2
	if o.UsageTrackingConfig.IsNull() || o.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayUsageTrackingConfig_SdkV2
	d := o.UsageTrackingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in PutAiGatewayResponse_SdkV2.
func (o *PutAiGatewayResponse_SdkV2) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_tracking_config"]
	o.UsageTrackingConfig = types.ListValueMust(t, vs)
}

// Update rate limits of a serving endpoint
type PutRequest_SdkV2 struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
}

func (newState *PutRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutRequest_SdkV2) {
}

func (newState *PutRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState PutRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rate_limits": reflect.TypeOf(RateLimit_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o PutRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        o.Name,
			"rate_limits": o.RateLimits,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.Type(ctx),
			},
		},
	}
}

// GetRateLimits returns the value of the RateLimits field in PutRequest_SdkV2 as
// a slice of RateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutRequest_SdkV2) GetRateLimits(ctx context.Context) ([]RateLimit_SdkV2, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit_SdkV2
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutRequest_SdkV2.
func (o *PutRequest_SdkV2) SetRateLimits(ctx context.Context, v []RateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

type PutResponse_SdkV2 struct {
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
}

func (newState *PutResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutResponse_SdkV2) {
}

func (newState *PutResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState PutResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rate_limits": reflect.TypeOf(RateLimit_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o PutResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"rate_limits": o.RateLimits,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.Type(ctx),
			},
		},
	}
}

// GetRateLimits returns the value of the RateLimits field in PutResponse_SdkV2 as
// a slice of RateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutResponse_SdkV2) GetRateLimits(ctx context.Context) ([]RateLimit_SdkV2, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit_SdkV2
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutResponse_SdkV2.
func (o *PutResponse_SdkV2) SetRateLimits(ctx context.Context, v []RateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

type QueryEndpointInput_SdkV2 struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords types.List `tfsdk:"dataframe_records" tf:"optional"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit types.List `tfsdk:"dataframe_split" tf:"optional,object"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams types.Map `tfsdk:"extra_params" tf:"optional"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input types.Object `tfsdk:"input" tf:"optional"`
	// Tensor-based input in columnar format.
	Inputs types.Object `tfsdk:"inputs" tf:"optional"`
	// Tensor-based input in row format.
	Instances types.List `tfsdk:"instances" tf:"optional"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens types.Int64 `tfsdk:"max_tokens" tf:"optional"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	Messages types.List `tfsdk:"messages" tf:"optional"`
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	N types.Int64 `tfsdk:"n" tf:"optional"`
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	Prompt types.Object `tfsdk:"prompt" tf:"optional"`
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	Stop types.List `tfsdk:"stop" tf:"optional"`
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	Stream types.Bool `tfsdk:"stream" tf:"optional"`
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	Temperature types.Float64 `tfsdk:"temperature" tf:"optional"`
}

func (newState *QueryEndpointInput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEndpointInput_SdkV2) {
}

func (newState *QueryEndpointInput_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryEndpointInput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryEndpointInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryEndpointInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataframe_records": reflect.TypeOf(types.Object{}),
		"dataframe_split":   reflect.TypeOf(DataframeSplitInput_SdkV2{}),
		"extra_params":      reflect.TypeOf(types.String{}),
		"instances":         reflect.TypeOf(types.Object{}),
		"messages":          reflect.TypeOf(ChatMessage_SdkV2{}),
		"stop":              reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEndpointInput_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryEndpointInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"dataframe_records": o.DataframeRecords,
			"dataframe_split":   o.DataframeSplit,
			"extra_params":      o.ExtraParams,
			"input":             o.Input,
			"inputs":            o.Inputs,
			"instances":         o.Instances,
			"max_tokens":        o.MaxTokens,
			"messages":          o.Messages,
			"n":                 o.N,
			"name":              o.Name,
			"prompt":            o.Prompt,
			"stop":              o.Stop,
			"stream":            o.Stream,
			"temperature":       o.Temperature,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryEndpointInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataframe_records": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"dataframe_split": basetypes.ListType{
				ElemType: DataframeSplitInput{}.Type(ctx),
			},
			"extra_params": basetypes.MapType{
				ElemType: types.StringType,
			},
			"input":  types.ObjectType{},
			"inputs": types.ObjectType{},
			"instances": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"max_tokens": types.Int64Type,
			"messages": basetypes.ListType{
				ElemType: ChatMessage{}.Type(ctx),
			},
			"n":      types.Int64Type,
			"name":   types.StringType,
			"prompt": types.ObjectType{},
			"stop": basetypes.ListType{
				ElemType: types.StringType,
			},
			"stream":      types.BoolType,
			"temperature": types.Float64Type,
		},
	}
}

// GetDataframeRecords returns the value of the DataframeRecords field in QueryEndpointInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput_SdkV2) GetDataframeRecords(ctx context.Context) ([]types.Object, bool) {
	if o.DataframeRecords.IsNull() || o.DataframeRecords.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := o.DataframeRecords.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataframeRecords sets the value of the DataframeRecords field in QueryEndpointInput_SdkV2.
func (o *QueryEndpointInput_SdkV2) SetDataframeRecords(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataframe_records"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataframeRecords = types.ListValueMust(t, vs)
}

// GetDataframeSplit returns the value of the DataframeSplit field in QueryEndpointInput_SdkV2 as
// a DataframeSplitInput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput_SdkV2) GetDataframeSplit(ctx context.Context) (DataframeSplitInput_SdkV2, bool) {
	var e DataframeSplitInput_SdkV2
	if o.DataframeSplit.IsNull() || o.DataframeSplit.IsUnknown() {
		return e, false
	}
	var v []DataframeSplitInput_SdkV2
	d := o.DataframeSplit.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataframeSplit sets the value of the DataframeSplit field in QueryEndpointInput_SdkV2.
func (o *QueryEndpointInput_SdkV2) SetDataframeSplit(ctx context.Context, v DataframeSplitInput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataframe_split"]
	o.DataframeSplit = types.ListValueMust(t, vs)
}

// GetExtraParams returns the value of the ExtraParams field in QueryEndpointInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput_SdkV2) GetExtraParams(ctx context.Context) (map[string]types.String, bool) {
	if o.ExtraParams.IsNull() || o.ExtraParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.ExtraParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExtraParams sets the value of the ExtraParams field in QueryEndpointInput_SdkV2.
func (o *QueryEndpointInput_SdkV2) SetExtraParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["extra_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExtraParams = types.MapValueMust(t, vs)
}

// GetInstances returns the value of the Instances field in QueryEndpointInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput_SdkV2) GetInstances(ctx context.Context) ([]types.Object, bool) {
	if o.Instances.IsNull() || o.Instances.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := o.Instances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstances sets the value of the Instances field in QueryEndpointInput_SdkV2.
func (o *QueryEndpointInput_SdkV2) SetInstances(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Instances = types.ListValueMust(t, vs)
}

// GetMessages returns the value of the Messages field in QueryEndpointInput_SdkV2 as
// a slice of ChatMessage_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput_SdkV2) GetMessages(ctx context.Context) ([]ChatMessage_SdkV2, bool) {
	if o.Messages.IsNull() || o.Messages.IsUnknown() {
		return nil, false
	}
	var v []ChatMessage_SdkV2
	d := o.Messages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessages sets the value of the Messages field in QueryEndpointInput_SdkV2.
func (o *QueryEndpointInput_SdkV2) SetMessages(ctx context.Context, v []ChatMessage_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Messages = types.ListValueMust(t, vs)
}

// GetStop returns the value of the Stop field in QueryEndpointInput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput_SdkV2) GetStop(ctx context.Context) ([]types.String, bool) {
	if o.Stop.IsNull() || o.Stop.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.Stop.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStop sets the value of the Stop field in QueryEndpointInput_SdkV2.
func (o *QueryEndpointInput_SdkV2) SetStop(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stop"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Stop = types.ListValueMust(t, vs)
}

type QueryEndpointResponse_SdkV2 struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices types.List `tfsdk:"choices" tf:"optional"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created types.Int64 `tfsdk:"created" tf:"optional"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data types.List `tfsdk:"data" tf:"optional"`
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	Model types.String `tfsdk:"model" tf:"optional"`
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	Object types.String `tfsdk:"object" tf:"optional"`
	// The predictions returned by the serving endpoint.
	Predictions types.List `tfsdk:"predictions" tf:"optional"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName types.String `tfsdk:"-"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage types.List `tfsdk:"usage" tf:"optional,object"`
}

func (newState *QueryEndpointResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEndpointResponse_SdkV2) {
}

func (newState *QueryEndpointResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState QueryEndpointResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryEndpointResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"choices":     reflect.TypeOf(V1ResponseChoiceElement_SdkV2{}),
		"data":        reflect.TypeOf(EmbeddingsV1ResponseEmbeddingElement_SdkV2{}),
		"predictions": reflect.TypeOf(types.Object{}),
		"usage":       reflect.TypeOf(ExternalModelUsageElement_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEndpointResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o QueryEndpointResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"choices":           o.Choices,
			"created":           o.Created,
			"data":              o.Data,
			"id":                o.Id,
			"model":             o.Model,
			"object":            o.Object,
			"predictions":       o.Predictions,
			"served-model-name": o.ServedModelName,
			"usage":             o.Usage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryEndpointResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"choices": basetypes.ListType{
				ElemType: V1ResponseChoiceElement{}.Type(ctx),
			},
			"created": types.Int64Type,
			"data": basetypes.ListType{
				ElemType: EmbeddingsV1ResponseEmbeddingElement{}.Type(ctx),
			},
			"id":     types.StringType,
			"model":  types.StringType,
			"object": types.StringType,
			"predictions": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"served-model-name": types.StringType,
			"usage": basetypes.ListType{
				ElemType: ExternalModelUsageElement{}.Type(ctx),
			},
		},
	}
}

// GetChoices returns the value of the Choices field in QueryEndpointResponse_SdkV2 as
// a slice of V1ResponseChoiceElement_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse_SdkV2) GetChoices(ctx context.Context) ([]V1ResponseChoiceElement_SdkV2, bool) {
	if o.Choices.IsNull() || o.Choices.IsUnknown() {
		return nil, false
	}
	var v []V1ResponseChoiceElement_SdkV2
	d := o.Choices.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChoices sets the value of the Choices field in QueryEndpointResponse_SdkV2.
func (o *QueryEndpointResponse_SdkV2) SetChoices(ctx context.Context, v []V1ResponseChoiceElement_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["choices"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Choices = types.ListValueMust(t, vs)
}

// GetData returns the value of the Data field in QueryEndpointResponse_SdkV2 as
// a slice of EmbeddingsV1ResponseEmbeddingElement_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse_SdkV2) GetData(ctx context.Context) ([]EmbeddingsV1ResponseEmbeddingElement_SdkV2, bool) {
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingsV1ResponseEmbeddingElement_SdkV2
	d := o.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in QueryEndpointResponse_SdkV2.
func (o *QueryEndpointResponse_SdkV2) SetData(ctx context.Context, v []EmbeddingsV1ResponseEmbeddingElement_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Data = types.ListValueMust(t, vs)
}

// GetPredictions returns the value of the Predictions field in QueryEndpointResponse_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse_SdkV2) GetPredictions(ctx context.Context) ([]types.Object, bool) {
	if o.Predictions.IsNull() || o.Predictions.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := o.Predictions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPredictions sets the value of the Predictions field in QueryEndpointResponse_SdkV2.
func (o *QueryEndpointResponse_SdkV2) SetPredictions(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["predictions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Predictions = types.ListValueMust(t, vs)
}

// GetUsage returns the value of the Usage field in QueryEndpointResponse_SdkV2 as
// a ExternalModelUsageElement_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse_SdkV2) GetUsage(ctx context.Context) (ExternalModelUsageElement_SdkV2, bool) {
	var e ExternalModelUsageElement_SdkV2
	if o.Usage.IsNull() || o.Usage.IsUnknown() {
		return e, false
	}
	var v []ExternalModelUsageElement_SdkV2
	d := o.Usage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsage sets the value of the Usage field in QueryEndpointResponse_SdkV2.
func (o *QueryEndpointResponse_SdkV2) SetUsage(ctx context.Context, v ExternalModelUsageElement_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["usage"]
	o.Usage = types.ListValueMust(t, vs)
}

type RateLimit_SdkV2 struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls types.Int64 `tfsdk:"calls" tf:""`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod types.String `tfsdk:"renewal_period" tf:""`
}

func (newState *RateLimit_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan RateLimit_SdkV2) {
}

func (newState *RateLimit_SdkV2) SyncEffectiveFieldsDuringRead(existingState RateLimit_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RateLimit.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a RateLimit_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RateLimit_SdkV2
// only implements ToObjectValue() and Type().
func (o RateLimit_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"calls":          o.Calls,
			"key":            o.Key,
			"renewal_period": o.RenewalPeriod,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RateLimit_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"renewal_period": types.StringType,
		},
	}
}

type Route_SdkV2 struct {
	// The name of the served model this route configures traffic for.
	ServedModelName types.String `tfsdk:"served_model_name" tf:""`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage types.Int64 `tfsdk:"traffic_percentage" tf:""`
}

func (newState *Route_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan Route_SdkV2) {
}

func (newState *Route_SdkV2) SyncEffectiveFieldsDuringRead(existingState Route_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Route.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Route_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Route_SdkV2
// only implements ToObjectValue() and Type().
func (o Route_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_model_name":  o.ServedModelName,
			"traffic_percentage": o.TrafficPercentage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Route_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_model_name":  types.StringType,
			"traffic_percentage": types.Int64Type,
		},
	}
}

type ServedEntityInput_SdkV2 struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName types.String `tfsdk:"entity_name" tf:"optional"`
	// The version of the model in Databricks Model Registry to be served or
	// empty if the entity is a FEATURE_SPEC.
	EntityVersion types.String `tfsdk:"entity_version" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars" tf:"optional"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel types.List `tfsdk:"external_model" tf:"optional,object"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput types.Int64 `tfsdk:"max_provisioned_throughput" tf:"optional"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput types.Int64 `tfsdk:"min_provisioned_throughput" tf:"optional"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to
	// <entity-name>-<entity-version>.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled" tf:"optional"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size is 0.
	WorkloadSize types.String `tfsdk:"workload_size" tf:"optional"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type" tf:"optional"`
}

func (newState *ServedEntityInput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntityInput_SdkV2) {
}

func (newState *ServedEntityInput_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedEntityInput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedEntityInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedEntityInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"external_model":   reflect.TypeOf(ExternalModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntityInput_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedEntityInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":                o.EntityName,
			"entity_version":             o.EntityVersion,
			"environment_vars":           o.EnvironmentVars,
			"external_model":             o.ExternalModel,
			"instance_profile_arn":       o.InstanceProfileArn,
			"max_provisioned_throughput": o.MaxProvisionedThroughput,
			"min_provisioned_throughput": o.MinProvisionedThroughput,
			"name":                       o.Name,
			"scale_to_zero_enabled":      o.ScaleToZeroEnabled,
			"workload_size":              o.WorkloadSize,
			"workload_type":              o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedEntityInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"external_model": basetypes.ListType{
				ElemType: ExternalModel{}.Type(ctx),
			},
			"instance_profile_arn":       types.StringType,
			"max_provisioned_throughput": types.Int64Type,
			"min_provisioned_throughput": types.Int64Type,
			"name":                       types.StringType,
			"scale_to_zero_enabled":      types.BoolType,
			"workload_size":              types.StringType,
			"workload_type":              types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedEntityInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityInput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if o.EnvironmentVars.IsNull() || o.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedEntityInput_SdkV2.
func (o *ServedEntityInput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntityInput_SdkV2 as
// a ExternalModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityInput_SdkV2) GetExternalModel(ctx context.Context) (ExternalModel_SdkV2, bool) {
	var e ExternalModel_SdkV2
	if o.ExternalModel.IsNull() || o.ExternalModel.IsUnknown() {
		return e, false
	}
	var v []ExternalModel_SdkV2
	d := o.ExternalModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntityInput_SdkV2.
func (o *ServedEntityInput_SdkV2) SetExternalModel(ctx context.Context, v ExternalModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_model"]
	o.ExternalModel = types.ListValueMust(t, vs)
}

type ServedEntityOutput_SdkV2 struct {
	// The creation timestamp of the served entity in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// The email of the user who created the served entity.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// The name of the entity served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName types.String `tfsdk:"entity_name" tf:"optional"`
	// The version of the served entity in Databricks Model Registry or empty if
	// the entity is a FEATURE_SPEC.
	EntityVersion types.String `tfsdk:"entity_version" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars" tf:"optional"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	ExternalModel types.List `tfsdk:"external_model" tf:"optional,object"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	FoundationModel types.List `tfsdk:"foundation_model" tf:"optional,object"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput types.Int64 `tfsdk:"max_provisioned_throughput" tf:"optional"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput types.Int64 `tfsdk:"min_provisioned_throughput" tf:"optional"`
	// The name of the served entity.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled" tf:"optional"`
	// Information corresponding to the state of the served entity.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize types.String `tfsdk:"workload_size" tf:"optional"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type" tf:"optional"`
}

func (newState *ServedEntityOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntityOutput_SdkV2) {
}

func (newState *ServedEntityOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedEntityOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedEntityOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedEntityOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"external_model":   reflect.TypeOf(ExternalModel_SdkV2{}),
		"foundation_model": reflect.TypeOf(FoundationModel_SdkV2{}),
		"state":            reflect.TypeOf(ServedModelState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntityOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedEntityOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":         o.CreationTimestamp,
			"creator":                    o.Creator,
			"entity_name":                o.EntityName,
			"entity_version":             o.EntityVersion,
			"environment_vars":           o.EnvironmentVars,
			"external_model":             o.ExternalModel,
			"foundation_model":           o.FoundationModel,
			"instance_profile_arn":       o.InstanceProfileArn,
			"max_provisioned_throughput": o.MaxProvisionedThroughput,
			"min_provisioned_throughput": o.MinProvisionedThroughput,
			"name":                       o.Name,
			"scale_to_zero_enabled":      o.ScaleToZeroEnabled,
			"state":                      o.State,
			"workload_size":              o.WorkloadSize,
			"workload_type":              o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedEntityOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"entity_name":        types.StringType,
			"entity_version":     types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"external_model": basetypes.ListType{
				ElemType: ExternalModel{}.Type(ctx),
			},
			"foundation_model": basetypes.ListType{
				ElemType: FoundationModel{}.Type(ctx),
			},
			"instance_profile_arn":       types.StringType,
			"max_provisioned_throughput": types.Int64Type,
			"min_provisioned_throughput": types.Int64Type,
			"name":                       types.StringType,
			"scale_to_zero_enabled":      types.BoolType,
			"state": basetypes.ListType{
				ElemType: ServedModelState{}.Type(ctx),
			},
			"workload_size": types.StringType,
			"workload_type": types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedEntityOutput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if o.EnvironmentVars.IsNull() || o.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedEntityOutput_SdkV2.
func (o *ServedEntityOutput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntityOutput_SdkV2 as
// a ExternalModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput_SdkV2) GetExternalModel(ctx context.Context) (ExternalModel_SdkV2, bool) {
	var e ExternalModel_SdkV2
	if o.ExternalModel.IsNull() || o.ExternalModel.IsUnknown() {
		return e, false
	}
	var v []ExternalModel_SdkV2
	d := o.ExternalModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntityOutput_SdkV2.
func (o *ServedEntityOutput_SdkV2) SetExternalModel(ctx context.Context, v ExternalModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_model"]
	o.ExternalModel = types.ListValueMust(t, vs)
}

// GetFoundationModel returns the value of the FoundationModel field in ServedEntityOutput_SdkV2 as
// a FoundationModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput_SdkV2) GetFoundationModel(ctx context.Context) (FoundationModel_SdkV2, bool) {
	var e FoundationModel_SdkV2
	if o.FoundationModel.IsNull() || o.FoundationModel.IsUnknown() {
		return e, false
	}
	var v []FoundationModel_SdkV2
	d := o.FoundationModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFoundationModel sets the value of the FoundationModel field in ServedEntityOutput_SdkV2.
func (o *ServedEntityOutput_SdkV2) SetFoundationModel(ctx context.Context, v FoundationModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foundation_model"]
	o.FoundationModel = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in ServedEntityOutput_SdkV2 as
// a ServedModelState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput_SdkV2) GetState(ctx context.Context) (ServedModelState_SdkV2, bool) {
	var e ServedModelState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []ServedModelState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServedEntityOutput_SdkV2.
func (o *ServedEntityOutput_SdkV2) SetState(ctx context.Context, v ServedModelState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

type ServedEntitySpec_SdkV2 struct {
	// The name of the entity served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName types.String `tfsdk:"entity_name" tf:"optional"`
	// The version of the served entity in Databricks Model Registry or empty if
	// the entity is a FEATURE_SPEC.
	EntityVersion types.String `tfsdk:"entity_version" tf:"optional"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	ExternalModel types.List `tfsdk:"external_model" tf:"optional,object"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	FoundationModel types.List `tfsdk:"foundation_model" tf:"optional,object"`
	// The name of the served entity.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ServedEntitySpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntitySpec_SdkV2) {
}

func (newState *ServedEntitySpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedEntitySpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedEntitySpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedEntitySpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_model":   reflect.TypeOf(ExternalModel_SdkV2{}),
		"foundation_model": reflect.TypeOf(FoundationModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntitySpec_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedEntitySpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":      o.EntityName,
			"entity_version":   o.EntityVersion,
			"external_model":   o.ExternalModel,
			"foundation_model": o.FoundationModel,
			"name":             o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedEntitySpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"external_model": basetypes.ListType{
				ElemType: ExternalModel{}.Type(ctx),
			},
			"foundation_model": basetypes.ListType{
				ElemType: FoundationModel{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntitySpec_SdkV2 as
// a ExternalModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntitySpec_SdkV2) GetExternalModel(ctx context.Context) (ExternalModel_SdkV2, bool) {
	var e ExternalModel_SdkV2
	if o.ExternalModel.IsNull() || o.ExternalModel.IsUnknown() {
		return e, false
	}
	var v []ExternalModel_SdkV2
	d := o.ExternalModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntitySpec_SdkV2.
func (o *ServedEntitySpec_SdkV2) SetExternalModel(ctx context.Context, v ExternalModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["external_model"]
	o.ExternalModel = types.ListValueMust(t, vs)
}

// GetFoundationModel returns the value of the FoundationModel field in ServedEntitySpec_SdkV2 as
// a FoundationModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntitySpec_SdkV2) GetFoundationModel(ctx context.Context) (FoundationModel_SdkV2, bool) {
	var e FoundationModel_SdkV2
	if o.FoundationModel.IsNull() || o.FoundationModel.IsUnknown() {
		return e, false
	}
	var v []FoundationModel_SdkV2
	d := o.FoundationModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFoundationModel sets the value of the FoundationModel field in ServedEntitySpec_SdkV2.
func (o *ServedEntitySpec_SdkV2) SetFoundationModel(ctx context.Context, v FoundationModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["foundation_model"]
	o.FoundationModel = types.ListValueMust(t, vs)
}

type ServedModelInput_SdkV2 struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars" tf:"optional"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput types.Int64 `tfsdk:"max_provisioned_throughput" tf:"optional"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput types.Int64 `tfsdk:"min_provisioned_throughput" tf:"optional"`
	// The name of the model in Databricks Model Registry to be served or if the
	// model resides in Unity Catalog, the full name of model, in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	ModelName types.String `tfsdk:"model_name" tf:""`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion types.String `tfsdk:"model_version" tf:""`
	// The name of a served model. It must be unique across an endpoint. If not
	// specified, this field will default to <model-name>-<model-version>. A
	// served model name can consist of alphanumeric characters, dashes, and
	// underscores.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Whether the compute resources for the served model should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled" tf:""`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize types.String `tfsdk:"workload_size" tf:"optional"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type" tf:"optional"`
}

func (newState *ServedModelInput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelInput_SdkV2) {
}

func (newState *ServedModelInput_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedModelInput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedModelInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelInput_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedModelInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"environment_vars":           o.EnvironmentVars,
			"instance_profile_arn":       o.InstanceProfileArn,
			"max_provisioned_throughput": o.MaxProvisionedThroughput,
			"min_provisioned_throughput": o.MinProvisionedThroughput,
			"model_name":                 o.ModelName,
			"model_version":              o.ModelVersion,
			"name":                       o.Name,
			"scale_to_zero_enabled":      o.ScaleToZeroEnabled,
			"workload_size":              o.WorkloadSize,
			"workload_type":              o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"instance_profile_arn":       types.StringType,
			"max_provisioned_throughput": types.Int64Type,
			"min_provisioned_throughput": types.Int64Type,
			"model_name":                 types.StringType,
			"model_version":              types.StringType,
			"name":                       types.StringType,
			"scale_to_zero_enabled":      types.BoolType,
			"workload_size":              types.StringType,
			"workload_type":              types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedModelInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedModelInput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if o.EnvironmentVars.IsNull() || o.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedModelInput_SdkV2.
func (o *ServedModelInput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

type ServedModelOutput_SdkV2 struct {
	// The creation timestamp of the served model in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// The email of the user who created the served model.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars" tf:"optional"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion types.String `tfsdk:"model_version" tf:"optional"`
	// The name of the served model.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Whether the compute resources for the Served Model should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled" tf:"optional"`
	// Information corresponding to the state of the Served Model.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize types.String `tfsdk:"workload_size" tf:"optional"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type" tf:"optional"`
}

func (newState *ServedModelOutput_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelOutput_SdkV2) {
}

func (newState *ServedModelOutput_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedModelOutput_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedModelOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"state":            reflect.TypeOf(ServedModelState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelOutput_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedModelOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":    o.CreationTimestamp,
			"creator":               o.Creator,
			"environment_vars":      o.EnvironmentVars,
			"instance_profile_arn":  o.InstanceProfileArn,
			"model_name":            o.ModelName,
			"model_version":         o.ModelVersion,
			"name":                  o.Name,
			"scale_to_zero_enabled": o.ScaleToZeroEnabled,
			"state":                 o.State,
			"workload_size":         o.WorkloadSize,
			"workload_type":         o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"instance_profile_arn":  types.StringType,
			"model_name":            types.StringType,
			"model_version":         types.StringType,
			"name":                  types.StringType,
			"scale_to_zero_enabled": types.BoolType,
			"state": basetypes.ListType{
				ElemType: ServedModelState{}.Type(ctx),
			},
			"workload_size": types.StringType,
			"workload_type": types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedModelOutput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedModelOutput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if o.EnvironmentVars.IsNull() || o.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedModelOutput_SdkV2.
func (o *ServedModelOutput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetState returns the value of the State field in ServedModelOutput_SdkV2 as
// a ServedModelState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedModelOutput_SdkV2) GetState(ctx context.Context) (ServedModelState_SdkV2, bool) {
	var e ServedModelState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []ServedModelState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServedModelOutput_SdkV2.
func (o *ServedModelOutput_SdkV2) SetState(ctx context.Context, v ServedModelState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

type ServedModelSpec_SdkV2 struct {
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion types.String `tfsdk:"model_version" tf:"optional"`
	// The name of the served model.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ServedModelSpec_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelSpec_SdkV2) {
}

func (newState *ServedModelSpec_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedModelSpec_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedModelSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelSpec_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedModelSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_name":    o.ModelName,
			"model_version": o.ModelVersion,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_name":    types.StringType,
			"model_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

type ServedModelState_SdkV2 struct {
	// The state of the served entity deployment. DEPLOYMENT_CREATING indicates
	// that the served entity is not ready yet because the deployment is still
	// being created (i.e container image is building, model server is deploying
	// for the first time, etc.). DEPLOYMENT_RECOVERING indicates that the
	// served entity was previously in a ready state but no longer is and is
	// attempting to recover. DEPLOYMENT_READY indicates that the served entity
	// is ready to receive traffic. DEPLOYMENT_FAILED indicates that there was
	// an error trying to bring up the served entity (e.g container image build
	// failed, the model server failed to start due to a model loading error,
	// etc.) DEPLOYMENT_ABORTED indicates that the deployment was terminated
	// likely due to a failure in bringing up another served entity under the
	// same endpoint and config version.
	Deployment types.String `tfsdk:"deployment" tf:"optional"`
	// More information about the state of the served entity, if available.
	DeploymentStateMessage types.String `tfsdk:"deployment_state_message" tf:"optional"`
}

func (newState *ServedModelState_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelState_SdkV2) {
}

func (newState *ServedModelState_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServedModelState_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServedModelState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelState_SdkV2
// only implements ToObjectValue() and Type().
func (o ServedModelState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"deployment":               o.Deployment,
			"deployment_state_message": o.DeploymentStateMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelState_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"deployment":               types.StringType,
			"deployment_state_message": types.StringType,
		},
	}
}

type ServerLogsResponse_SdkV2 struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs types.String `tfsdk:"logs" tf:""`
}

func (newState *ServerLogsResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServerLogsResponse_SdkV2) {
}

func (newState *ServerLogsResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServerLogsResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServerLogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServerLogsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServerLogsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ServerLogsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"logs": o.Logs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServerLogsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ServingEndpoint_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: Only
	// external model endpoints are currently supported.
	AiGateway types.List `tfsdk:"ai_gateway" tf:"optional,object"`
	// The config that is currently being served by the endpoint.
	Config types.List `tfsdk:"config" tf:"optional,object"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id types.String `tfsdk:"id" tf:"optional"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The name of the serving endpoint.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Information corresponding to the state of the serving endpoint.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// Tags attached to the serving endpoint.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task" tf:"optional"`
}

func (newState *ServingEndpoint_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpoint_SdkV2) {
}

func (newState *ServingEndpoint_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpoint_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway": reflect.TypeOf(AiGatewayConfig_SdkV2{}),
		"config":     reflect.TypeOf(EndpointCoreConfigSummary_SdkV2{}),
		"state":      reflect.TypeOf(EndpointState_SdkV2{}),
		"tags":       reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":             o.AiGateway,
			"config":                 o.Config,
			"creation_timestamp":     o.CreationTimestamp,
			"creator":                o.Creator,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"name":                   o.Name,
			"state":                  o.State,
			"tags":                   o.Tags,
			"task":                   o.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig{}.Type(ctx),
			},
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigSummary{}.Type(ctx),
			},
			"creation_timestamp":     types.Int64Type,
			"creator":                types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"state": basetypes.ListType{
				ElemType: EndpointState{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
			"task": types.StringType,
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in ServingEndpoint_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := o.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in ServingEndpoint_SdkV2.
func (o *ServingEndpoint_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	o.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in ServingEndpoint_SdkV2 as
// a EndpointCoreConfigSummary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint_SdkV2) GetConfig(ctx context.Context) (EndpointCoreConfigSummary_SdkV2, bool) {
	var e EndpointCoreConfigSummary_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []EndpointCoreConfigSummary_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in ServingEndpoint_SdkV2.
func (o *ServingEndpoint_SdkV2) SetConfig(ctx context.Context, v EndpointCoreConfigSummary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in ServingEndpoint_SdkV2 as
// a EndpointState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint_SdkV2) GetState(ctx context.Context) (EndpointState_SdkV2, bool) {
	var e EndpointState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []EndpointState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServingEndpoint_SdkV2.
func (o *ServingEndpoint_SdkV2) SetState(ctx context.Context, v EndpointState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ServingEndpoint_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ServingEndpoint_SdkV2.
func (o *ServingEndpoint_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ServingEndpointAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ServingEndpointAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointAccessControlRequest_SdkV2) {
}

func (newState *ServingEndpointAccessControlRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointAccessControlRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServingEndpointAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ServingEndpointAccessControlResponse_SdkV2 struct {
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

func (newState *ServingEndpointAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointAccessControlResponse_SdkV2) {
}

func (newState *ServingEndpointAccessControlResponse_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointAccessControlResponse_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ServingEndpointPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServingEndpointAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ServingEndpointPermission{}.Type(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

// GetAllPermissions returns the value of the AllPermissions field in ServingEndpointAccessControlResponse_SdkV2 as
// a slice of ServingEndpointPermission_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]ServingEndpointPermission_SdkV2, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointPermission_SdkV2
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ServingEndpointAccessControlResponse_SdkV2.
func (o *ServingEndpointAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []ServingEndpointPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type ServingEndpointDetailed_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: Only
	// external model endpoints are currently supported.
	AiGateway types.List `tfsdk:"ai_gateway" tf:"optional,object"`
	// The config that is currently being served by the endpoint.
	Config types.List `tfsdk:"config" tf:"optional,object"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo types.List `tfsdk:"data_plane_info" tf:"optional,object"`
	// Endpoint invocation url if route optimization is enabled for endpoint
	EndpointUrl types.String `tfsdk:"endpoint_url" tf:"optional"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id types.String `tfsdk:"id" tf:"optional"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The name of the serving endpoint.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The config that the endpoint is attempting to update to.
	PendingConfig types.List `tfsdk:"pending_config" tf:"optional,object"`
	// The permission level of the principal making the request.
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized types.Bool `tfsdk:"route_optimized" tf:"optional"`
	// Information corresponding to the state of the serving endpoint.
	State types.List `tfsdk:"state" tf:"optional,object"`
	// Tags attached to the serving endpoint.
	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task" tf:"optional"`
}

func (newState *ServingEndpointDetailed_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointDetailed_SdkV2) {
}

func (newState *ServingEndpointDetailed_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointDetailed_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointDetailed.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointDetailed_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":      reflect.TypeOf(AiGatewayConfig_SdkV2{}),
		"config":          reflect.TypeOf(EndpointCoreConfigOutput_SdkV2{}),
		"data_plane_info": reflect.TypeOf(ModelDataPlaneInfo_SdkV2{}),
		"pending_config":  reflect.TypeOf(EndpointPendingConfig_SdkV2{}),
		"state":           reflect.TypeOf(EndpointState_SdkV2{}),
		"tags":            reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointDetailed_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointDetailed_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":             o.AiGateway,
			"config":                 o.Config,
			"creation_timestamp":     o.CreationTimestamp,
			"creator":                o.Creator,
			"data_plane_info":        o.DataPlaneInfo,
			"endpoint_url":           o.EndpointUrl,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"name":                   o.Name,
			"pending_config":         o.PendingConfig,
			"permission_level":       o.PermissionLevel,
			"route_optimized":        o.RouteOptimized,
			"state":                  o.State,
			"tags":                   o.Tags,
			"task":                   o.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointDetailed_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig{}.Type(ctx),
			},
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigOutput{}.Type(ctx),
			},
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"data_plane_info": basetypes.ListType{
				ElemType: ModelDataPlaneInfo{}.Type(ctx),
			},
			"endpoint_url":           types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"pending_config": basetypes.ListType{
				ElemType: EndpointPendingConfig{}.Type(ctx),
			},
			"permission_level": types.StringType,
			"route_optimized":  types.BoolType,
			"state": basetypes.ListType{
				ElemType: EndpointState{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
			"task": types.StringType,
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in ServingEndpointDetailed_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := o.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in ServingEndpointDetailed_SdkV2.
func (o *ServingEndpointDetailed_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	o.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in ServingEndpointDetailed_SdkV2 as
// a EndpointCoreConfigOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed_SdkV2) GetConfig(ctx context.Context) (EndpointCoreConfigOutput_SdkV2, bool) {
	var e EndpointCoreConfigOutput_SdkV2
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v []EndpointCoreConfigOutput_SdkV2
	d := o.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in ServingEndpointDetailed_SdkV2.
func (o *ServingEndpointDetailed_SdkV2) SetConfig(ctx context.Context, v EndpointCoreConfigOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	o.Config = types.ListValueMust(t, vs)
}

// GetDataPlaneInfo returns the value of the DataPlaneInfo field in ServingEndpointDetailed_SdkV2 as
// a ModelDataPlaneInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed_SdkV2) GetDataPlaneInfo(ctx context.Context) (ModelDataPlaneInfo_SdkV2, bool) {
	var e ModelDataPlaneInfo_SdkV2
	if o.DataPlaneInfo.IsNull() || o.DataPlaneInfo.IsUnknown() {
		return e, false
	}
	var v []ModelDataPlaneInfo_SdkV2
	d := o.DataPlaneInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataPlaneInfo sets the value of the DataPlaneInfo field in ServingEndpointDetailed_SdkV2.
func (o *ServingEndpointDetailed_SdkV2) SetDataPlaneInfo(ctx context.Context, v ModelDataPlaneInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data_plane_info"]
	o.DataPlaneInfo = types.ListValueMust(t, vs)
}

// GetPendingConfig returns the value of the PendingConfig field in ServingEndpointDetailed_SdkV2 as
// a EndpointPendingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed_SdkV2) GetPendingConfig(ctx context.Context) (EndpointPendingConfig_SdkV2, bool) {
	var e EndpointPendingConfig_SdkV2
	if o.PendingConfig.IsNull() || o.PendingConfig.IsUnknown() {
		return e, false
	}
	var v []EndpointPendingConfig_SdkV2
	d := o.PendingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPendingConfig sets the value of the PendingConfig field in ServingEndpointDetailed_SdkV2.
func (o *ServingEndpointDetailed_SdkV2) SetPendingConfig(ctx context.Context, v EndpointPendingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["pending_config"]
	o.PendingConfig = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in ServingEndpointDetailed_SdkV2 as
// a EndpointState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed_SdkV2) GetState(ctx context.Context) (EndpointState_SdkV2, bool) {
	var e EndpointState_SdkV2
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v []EndpointState_SdkV2
	d := o.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServingEndpointDetailed_SdkV2.
func (o *ServingEndpointDetailed_SdkV2) SetState(ctx context.Context, v EndpointState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	o.State = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ServingEndpointDetailed_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ServingEndpointDetailed_SdkV2.
func (o *ServingEndpointDetailed_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ServingEndpointPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ServingEndpointPermission_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermission_SdkV2) {
}

func (newState *ServingEndpointPermission_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermission_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermission_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermission_SdkV2) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in ServingEndpointPermission_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ServingEndpointPermission_SdkV2.
func (o *ServingEndpointPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ServingEndpointPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ServingEndpointPermissions_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissions_SdkV2) {
}

func (newState *ServingEndpointPermissions_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissions_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ServingEndpointAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlResponse{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ServingEndpointPermissions_SdkV2 as
// a slice of ServingEndpointAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]ServingEndpointAccessControlResponse_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointAccessControlResponse_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ServingEndpointPermissions_SdkV2.
func (o *ServingEndpointPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []ServingEndpointAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ServingEndpointPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ServingEndpointPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissionsDescription_SdkV2) {
}

func (newState *ServingEndpointPermissionsDescription_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissionsDescription_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ServingEndpointPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *ServingEndpointPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissionsRequest_SdkV2) {
}

func (newState *ServingEndpointPermissionsRequest_SdkV2) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissionsRequest_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ServingEndpointPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ServingEndpointAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"serving_endpoint_id": o.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlRequest{}.Type(ctx),
			},
			"serving_endpoint_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ServingEndpointPermissionsRequest_SdkV2 as
// a slice of ServingEndpointAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]ServingEndpointAccessControlRequest_SdkV2, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointAccessControlRequest_SdkV2
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ServingEndpointPermissionsRequest_SdkV2.
func (o *ServingEndpointPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []ServingEndpointAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type TrafficConfig_SdkV2 struct {
	// The list of routes that define traffic to each served entity.
	Routes types.List `tfsdk:"routes" tf:"optional"`
}

func (newState *TrafficConfig_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrafficConfig_SdkV2) {
}

func (newState *TrafficConfig_SdkV2) SyncEffectiveFieldsDuringRead(existingState TrafficConfig_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrafficConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrafficConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"routes": reflect.TypeOf(Route_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrafficConfig_SdkV2
// only implements ToObjectValue() and Type().
func (o TrafficConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"routes": o.Routes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrafficConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"routes": basetypes.ListType{
				ElemType: Route{}.Type(ctx),
			},
		},
	}
}

// GetRoutes returns the value of the Routes field in TrafficConfig_SdkV2 as
// a slice of Route_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (o *TrafficConfig_SdkV2) GetRoutes(ctx context.Context) ([]Route_SdkV2, bool) {
	if o.Routes.IsNull() || o.Routes.IsUnknown() {
		return nil, false
	}
	var v []Route_SdkV2
	d := o.Routes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoutes sets the value of the Routes field in TrafficConfig_SdkV2.
func (o *TrafficConfig_SdkV2) SetRoutes(ctx context.Context, v []Route_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["routes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Routes = types.ListValueMust(t, vs)
}

type V1ResponseChoiceElement_SdkV2 struct {
	// The finish reason returned by the endpoint.
	FinishReason types.String `tfsdk:"finishReason" tf:"optional"`
	// The index of the choice in the __chat or completions__ response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs types.Int64 `tfsdk:"logprobs" tf:"optional"`
	// The message response from the __chat__ endpoint.
	Message types.List `tfsdk:"message" tf:"optional,object"`
	// The text response from the __completions__ endpoint.
	Text types.String `tfsdk:"text" tf:"optional"`
}

func (newState *V1ResponseChoiceElement_SdkV2) SyncEffectiveFieldsDuringCreateOrUpdate(plan V1ResponseChoiceElement_SdkV2) {
}

func (newState *V1ResponseChoiceElement_SdkV2) SyncEffectiveFieldsDuringRead(existingState V1ResponseChoiceElement_SdkV2) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in V1ResponseChoiceElement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a V1ResponseChoiceElement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"message": reflect.TypeOf(ChatMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, V1ResponseChoiceElement_SdkV2
// only implements ToObjectValue() and Type().
func (o V1ResponseChoiceElement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"finishReason": o.FinishReason,
			"index":        o.Index,
			"logprobs":     o.Logprobs,
			"message":      o.Message,
			"text":         o.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (o V1ResponseChoiceElement_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"finishReason": types.StringType,
			"index":        types.Int64Type,
			"logprobs":     types.Int64Type,
			"message": basetypes.ListType{
				ElemType: ChatMessage{}.Type(ctx),
			},
			"text": types.StringType,
		},
	}
}

// GetMessage returns the value of the Message field in V1ResponseChoiceElement_SdkV2 as
// a ChatMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (o *V1ResponseChoiceElement_SdkV2) GetMessage(ctx context.Context) (ChatMessage_SdkV2, bool) {
	var e ChatMessage_SdkV2
	if o.Message.IsNull() || o.Message.IsUnknown() {
		return e, false
	}
	var v []ChatMessage_SdkV2
	d := o.Message.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMessage sets the value of the Message field in V1ResponseChoiceElement_SdkV2.
func (o *V1ResponseChoiceElement_SdkV2) SetMessage(ctx context.Context, v ChatMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["message"]
	o.Message = types.ListValueMust(t, vs)
}
