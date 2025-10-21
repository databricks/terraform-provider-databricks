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
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Ai21LabsConfig_SdkV2 struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKey types.String `tfsdk:"ai21labs_api_key"`
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKeyPlaintext types.String `tfsdk:"ai21labs_api_key_plaintext"`
}

func (to *Ai21LabsConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Ai21LabsConfig_SdkV2) {
}

func (to *Ai21LabsConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Ai21LabsConfig_SdkV2) {
}

func (m Ai21LabsConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai21labs_api_key"] = attrs["ai21labs_api_key"].SetOptional()
	attrs["ai21labs_api_key_plaintext"] = attrs["ai21labs_api_key_plaintext"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Ai21LabsConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Ai21LabsConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Ai21LabsConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m Ai21LabsConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai21labs_api_key":           m.Ai21labsApiKey,
			"ai21labs_api_key_plaintext": m.Ai21labsApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Ai21LabsConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_api_key":           types.StringType,
			"ai21labs_api_key_plaintext": types.StringType,
		},
	}
}

type AiGatewayConfig_SdkV2 struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig types.List `tfsdk:"fallback_config"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.List `tfsdk:"guardrails"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.List `tfsdk:"inference_table_config"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.List `tfsdk:"usage_tracking_config"`
}

func (to *AiGatewayConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayConfig_SdkV2) {
	if !from.FallbackConfig.IsNull() && !from.FallbackConfig.IsUnknown() {
		if toFallbackConfig, ok := to.GetFallbackConfig(ctx); ok {
			if fromFallbackConfig, ok := from.GetFallbackConfig(ctx); ok {
				// Recursively sync the fields of FallbackConfig
				toFallbackConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromFallbackConfig)
				to.SetFallbackConfig(ctx, toFallbackConfig)
			}
		}
	}
	if !from.Guardrails.IsNull() && !from.Guardrails.IsUnknown() {
		if toGuardrails, ok := to.GetGuardrails(ctx); ok {
			if fromGuardrails, ok := from.GetGuardrails(ctx); ok {
				// Recursively sync the fields of Guardrails
				toGuardrails.SyncFieldsDuringCreateOrUpdate(ctx, fromGuardrails)
				to.SetGuardrails(ctx, toGuardrails)
			}
		}
	}
	if !from.InferenceTableConfig.IsNull() && !from.InferenceTableConfig.IsUnknown() {
		if toInferenceTableConfig, ok := to.GetInferenceTableConfig(ctx); ok {
			if fromInferenceTableConfig, ok := from.GetInferenceTableConfig(ctx); ok {
				// Recursively sync the fields of InferenceTableConfig
				toInferenceTableConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromInferenceTableConfig)
				to.SetInferenceTableConfig(ctx, toInferenceTableConfig)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.UsageTrackingConfig.IsNull() && !from.UsageTrackingConfig.IsUnknown() {
		if toUsageTrackingConfig, ok := to.GetUsageTrackingConfig(ctx); ok {
			if fromUsageTrackingConfig, ok := from.GetUsageTrackingConfig(ctx); ok {
				// Recursively sync the fields of UsageTrackingConfig
				toUsageTrackingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromUsageTrackingConfig)
				to.SetUsageTrackingConfig(ctx, toUsageTrackingConfig)
			}
		}
	}
}

func (to *AiGatewayConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayConfig_SdkV2) {
	if !from.FallbackConfig.IsNull() && !from.FallbackConfig.IsUnknown() {
		if toFallbackConfig, ok := to.GetFallbackConfig(ctx); ok {
			if fromFallbackConfig, ok := from.GetFallbackConfig(ctx); ok {
				toFallbackConfig.SyncFieldsDuringRead(ctx, fromFallbackConfig)
				to.SetFallbackConfig(ctx, toFallbackConfig)
			}
		}
	}
	if !from.Guardrails.IsNull() && !from.Guardrails.IsUnknown() {
		if toGuardrails, ok := to.GetGuardrails(ctx); ok {
			if fromGuardrails, ok := from.GetGuardrails(ctx); ok {
				toGuardrails.SyncFieldsDuringRead(ctx, fromGuardrails)
				to.SetGuardrails(ctx, toGuardrails)
			}
		}
	}
	if !from.InferenceTableConfig.IsNull() && !from.InferenceTableConfig.IsUnknown() {
		if toInferenceTableConfig, ok := to.GetInferenceTableConfig(ctx); ok {
			if fromInferenceTableConfig, ok := from.GetInferenceTableConfig(ctx); ok {
				toInferenceTableConfig.SyncFieldsDuringRead(ctx, fromInferenceTableConfig)
				to.SetInferenceTableConfig(ctx, toInferenceTableConfig)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.UsageTrackingConfig.IsNull() && !from.UsageTrackingConfig.IsUnknown() {
		if toUsageTrackingConfig, ok := to.GetUsageTrackingConfig(ctx); ok {
			if fromUsageTrackingConfig, ok := from.GetUsageTrackingConfig(ctx); ok {
				toUsageTrackingConfig.SyncFieldsDuringRead(ctx, fromUsageTrackingConfig)
				to.SetUsageTrackingConfig(ctx, toUsageTrackingConfig)
			}
		}
	}
}

func (m AiGatewayConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fallback_config"] = attrs["fallback_config"].SetOptional()
	attrs["fallback_config"] = attrs["fallback_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["guardrails"] = attrs["guardrails"].SetOptional()
	attrs["guardrails"] = attrs["guardrails"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["inference_table_config"] = attrs["inference_table_config"].SetOptional()
	attrs["inference_table_config"] = attrs["inference_table_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fallback_config":        reflect.TypeOf(FallbackConfig_SdkV2{}),
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails_SdkV2{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig_SdkV2{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit_SdkV2{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fallback_config":        m.FallbackConfig,
			"guardrails":             m.Guardrails,
			"inference_table_config": m.InferenceTableConfig,
			"rate_limits":            m.RateLimits,
			"usage_tracking_config":  m.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fallback_config": basetypes.ListType{
				ElemType: FallbackConfig_SdkV2{}.Type(ctx),
			},
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails_SdkV2{}.Type(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig_SdkV2{}.Type(ctx),
			},
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit_SdkV2{}.Type(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFallbackConfig returns the value of the FallbackConfig field in AiGatewayConfig_SdkV2 as
// a FallbackConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayConfig_SdkV2) GetFallbackConfig(ctx context.Context) (FallbackConfig_SdkV2, bool) {
	var e FallbackConfig_SdkV2
	if m.FallbackConfig.IsNull() || m.FallbackConfig.IsUnknown() {
		return e, false
	}
	var v []FallbackConfig_SdkV2
	d := m.FallbackConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFallbackConfig sets the value of the FallbackConfig field in AiGatewayConfig_SdkV2.
func (m *AiGatewayConfig_SdkV2) SetFallbackConfig(ctx context.Context, v FallbackConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["fallback_config"]
	m.FallbackConfig = types.ListValueMust(t, vs)
}

// GetGuardrails returns the value of the Guardrails field in AiGatewayConfig_SdkV2 as
// a AiGatewayGuardrails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayConfig_SdkV2) GetGuardrails(ctx context.Context) (AiGatewayGuardrails_SdkV2, bool) {
	var e AiGatewayGuardrails_SdkV2
	if m.Guardrails.IsNull() || m.Guardrails.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrails_SdkV2
	d := m.Guardrails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGuardrails sets the value of the Guardrails field in AiGatewayConfig_SdkV2.
func (m *AiGatewayConfig_SdkV2) SetGuardrails(ctx context.Context, v AiGatewayGuardrails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["guardrails"]
	m.Guardrails = types.ListValueMust(t, vs)
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in AiGatewayConfig_SdkV2 as
// a AiGatewayInferenceTableConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayConfig_SdkV2) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig_SdkV2, bool) {
	var e AiGatewayInferenceTableConfig_SdkV2
	if m.InferenceTableConfig.IsNull() || m.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayInferenceTableConfig_SdkV2
	d := m.InferenceTableConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in AiGatewayConfig_SdkV2.
func (m *AiGatewayConfig_SdkV2) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_table_config"]
	m.InferenceTableConfig = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in AiGatewayConfig_SdkV2 as
// a slice of AiGatewayRateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayConfig_SdkV2) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit_SdkV2, bool) {
	if m.RateLimits.IsNull() || m.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit_SdkV2
	d := m.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in AiGatewayConfig_SdkV2.
func (m *AiGatewayConfig_SdkV2) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in AiGatewayConfig_SdkV2 as
// a AiGatewayUsageTrackingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayConfig_SdkV2) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig_SdkV2, bool) {
	var e AiGatewayUsageTrackingConfig_SdkV2
	if m.UsageTrackingConfig.IsNull() || m.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayUsageTrackingConfig_SdkV2
	d := m.UsageTrackingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in AiGatewayConfig_SdkV2.
func (m *AiGatewayConfig_SdkV2) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_tracking_config"]
	m.UsageTrackingConfig = types.ListValueMust(t, vs)
}

type AiGatewayGuardrailParameters_SdkV2 struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	InvalidKeywords types.List `tfsdk:"invalid_keywords"`
	// Configuration for guardrail PII filter.
	Pii types.List `tfsdk:"pii"`
	// Indicates whether the safety filter is enabled.
	Safety types.Bool `tfsdk:"safety"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics types.List `tfsdk:"valid_topics"`
}

func (to *AiGatewayGuardrailParameters_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayGuardrailParameters_SdkV2) {
	if !from.InvalidKeywords.IsNull() && !from.InvalidKeywords.IsUnknown() && to.InvalidKeywords.IsNull() && len(from.InvalidKeywords.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InvalidKeywords, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InvalidKeywords = from.InvalidKeywords
	}
	if !from.Pii.IsNull() && !from.Pii.IsUnknown() {
		if toPii, ok := to.GetPii(ctx); ok {
			if fromPii, ok := from.GetPii(ctx); ok {
				// Recursively sync the fields of Pii
				toPii.SyncFieldsDuringCreateOrUpdate(ctx, fromPii)
				to.SetPii(ctx, toPii)
			}
		}
	}
	if !from.ValidTopics.IsNull() && !from.ValidTopics.IsUnknown() && to.ValidTopics.IsNull() && len(from.ValidTopics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ValidTopics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ValidTopics = from.ValidTopics
	}
}

func (to *AiGatewayGuardrailParameters_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayGuardrailParameters_SdkV2) {
	if !from.InvalidKeywords.IsNull() && !from.InvalidKeywords.IsUnknown() && to.InvalidKeywords.IsNull() && len(from.InvalidKeywords.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InvalidKeywords, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InvalidKeywords = from.InvalidKeywords
	}
	if !from.Pii.IsNull() && !from.Pii.IsUnknown() {
		if toPii, ok := to.GetPii(ctx); ok {
			if fromPii, ok := from.GetPii(ctx); ok {
				toPii.SyncFieldsDuringRead(ctx, fromPii)
				to.SetPii(ctx, toPii)
			}
		}
	}
	if !from.ValidTopics.IsNull() && !from.ValidTopics.IsUnknown() && to.ValidTopics.IsNull() && len(from.ValidTopics.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ValidTopics, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ValidTopics = from.ValidTopics
	}
}

func (m AiGatewayGuardrailParameters_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["invalid_keywords"] = attrs["invalid_keywords"].SetOptional()
	attrs["pii"] = attrs["pii"].SetOptional()
	attrs["pii"] = attrs["pii"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["safety"] = attrs["safety"].SetOptional()
	attrs["valid_topics"] = attrs["valid_topics"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrailParameters.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayGuardrailParameters_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"invalid_keywords": reflect.TypeOf(types.String{}),
		"pii":              reflect.TypeOf(AiGatewayGuardrailPiiBehavior_SdkV2{}),
		"valid_topics":     reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrailParameters_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayGuardrailParameters_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"invalid_keywords": m.InvalidKeywords,
			"pii":              m.Pii,
			"safety":           m.Safety,
			"valid_topics":     m.ValidTopics,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayGuardrailParameters_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"invalid_keywords": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pii": basetypes.ListType{
				ElemType: AiGatewayGuardrailPiiBehavior_SdkV2{}.Type(ctx),
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
func (m *AiGatewayGuardrailParameters_SdkV2) GetInvalidKeywords(ctx context.Context) ([]types.String, bool) {
	if m.InvalidKeywords.IsNull() || m.InvalidKeywords.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.InvalidKeywords.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInvalidKeywords sets the value of the InvalidKeywords field in AiGatewayGuardrailParameters_SdkV2.
func (m *AiGatewayGuardrailParameters_SdkV2) SetInvalidKeywords(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["invalid_keywords"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InvalidKeywords = types.ListValueMust(t, vs)
}

// GetPii returns the value of the Pii field in AiGatewayGuardrailParameters_SdkV2 as
// a AiGatewayGuardrailPiiBehavior_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayGuardrailParameters_SdkV2) GetPii(ctx context.Context) (AiGatewayGuardrailPiiBehavior_SdkV2, bool) {
	var e AiGatewayGuardrailPiiBehavior_SdkV2
	if m.Pii.IsNull() || m.Pii.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrailPiiBehavior_SdkV2
	d := m.Pii.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPii sets the value of the Pii field in AiGatewayGuardrailParameters_SdkV2.
func (m *AiGatewayGuardrailParameters_SdkV2) SetPii(ctx context.Context, v AiGatewayGuardrailPiiBehavior_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pii"]
	m.Pii = types.ListValueMust(t, vs)
}

// GetValidTopics returns the value of the ValidTopics field in AiGatewayGuardrailParameters_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayGuardrailParameters_SdkV2) GetValidTopics(ctx context.Context) ([]types.String, bool) {
	if m.ValidTopics.IsNull() || m.ValidTopics.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.ValidTopics.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetValidTopics sets the value of the ValidTopics field in AiGatewayGuardrailParameters_SdkV2.
func (m *AiGatewayGuardrailParameters_SdkV2) SetValidTopics(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["valid_topics"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ValidTopics = types.ListValueMust(t, vs)
}

type AiGatewayGuardrailPiiBehavior_SdkV2 struct {
	// Configuration for input guardrail filters.
	Behavior types.String `tfsdk:"behavior"`
}

func (to *AiGatewayGuardrailPiiBehavior_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayGuardrailPiiBehavior_SdkV2) {
}

func (to *AiGatewayGuardrailPiiBehavior_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayGuardrailPiiBehavior_SdkV2) {
}

func (m AiGatewayGuardrailPiiBehavior_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["behavior"] = attrs["behavior"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrailPiiBehavior.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayGuardrailPiiBehavior_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrailPiiBehavior_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayGuardrailPiiBehavior_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"behavior": m.Behavior,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayGuardrailPiiBehavior_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"behavior": types.StringType,
		},
	}
}

type AiGatewayGuardrails_SdkV2 struct {
	// Configuration for input guardrail filters.
	Input types.List `tfsdk:"input"`
	// Configuration for output guardrail filters.
	Output types.List `tfsdk:"output"`
}

func (to *AiGatewayGuardrails_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayGuardrails_SdkV2) {
	if !from.Input.IsNull() && !from.Input.IsUnknown() {
		if toInput, ok := to.GetInput(ctx); ok {
			if fromInput, ok := from.GetInput(ctx); ok {
				// Recursively sync the fields of Input
				toInput.SyncFieldsDuringCreateOrUpdate(ctx, fromInput)
				to.SetInput(ctx, toInput)
			}
		}
	}
	if !from.Output.IsNull() && !from.Output.IsUnknown() {
		if toOutput, ok := to.GetOutput(ctx); ok {
			if fromOutput, ok := from.GetOutput(ctx); ok {
				// Recursively sync the fields of Output
				toOutput.SyncFieldsDuringCreateOrUpdate(ctx, fromOutput)
				to.SetOutput(ctx, toOutput)
			}
		}
	}
}

func (to *AiGatewayGuardrails_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayGuardrails_SdkV2) {
	if !from.Input.IsNull() && !from.Input.IsUnknown() {
		if toInput, ok := to.GetInput(ctx); ok {
			if fromInput, ok := from.GetInput(ctx); ok {
				toInput.SyncFieldsDuringRead(ctx, fromInput)
				to.SetInput(ctx, toInput)
			}
		}
	}
	if !from.Output.IsNull() && !from.Output.IsUnknown() {
		if toOutput, ok := to.GetOutput(ctx); ok {
			if fromOutput, ok := from.GetOutput(ctx); ok {
				toOutput.SyncFieldsDuringRead(ctx, fromOutput)
				to.SetOutput(ctx, toOutput)
			}
		}
	}
}

func (m AiGatewayGuardrails_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["input"] = attrs["input"].SetOptional()
	attrs["input"] = attrs["input"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["output"] = attrs["output"].SetOptional()
	attrs["output"] = attrs["output"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayGuardrails_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input":  reflect.TypeOf(AiGatewayGuardrailParameters_SdkV2{}),
		"output": reflect.TypeOf(AiGatewayGuardrailParameters_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrails_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayGuardrails_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"input":  m.Input,
			"output": m.Output,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayGuardrails_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"input": basetypes.ListType{
				ElemType: AiGatewayGuardrailParameters_SdkV2{}.Type(ctx),
			},
			"output": basetypes.ListType{
				ElemType: AiGatewayGuardrailParameters_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetInput returns the value of the Input field in AiGatewayGuardrails_SdkV2 as
// a AiGatewayGuardrailParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayGuardrails_SdkV2) GetInput(ctx context.Context) (AiGatewayGuardrailParameters_SdkV2, bool) {
	var e AiGatewayGuardrailParameters_SdkV2
	if m.Input.IsNull() || m.Input.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrailParameters_SdkV2
	d := m.Input.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInput sets the value of the Input field in AiGatewayGuardrails_SdkV2.
func (m *AiGatewayGuardrails_SdkV2) SetInput(ctx context.Context, v AiGatewayGuardrailParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["input"]
	m.Input = types.ListValueMust(t, vs)
}

// GetOutput returns the value of the Output field in AiGatewayGuardrails_SdkV2 as
// a AiGatewayGuardrailParameters_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AiGatewayGuardrails_SdkV2) GetOutput(ctx context.Context) (AiGatewayGuardrailParameters_SdkV2, bool) {
	var e AiGatewayGuardrailParameters_SdkV2
	if m.Output.IsNull() || m.Output.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrailParameters_SdkV2
	d := m.Output.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOutput sets the value of the Output field in AiGatewayGuardrails_SdkV2.
func (m *AiGatewayGuardrails_SdkV2) SetOutput(ctx context.Context, v AiGatewayGuardrailParameters_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["output"]
	m.Output = types.ListValueMust(t, vs)
}

type AiGatewayInferenceTableConfig_SdkV2 struct {
	// The name of the catalog in Unity Catalog. Required when enabling
	// inference tables. NOTE: On update, you have to disable inference table
	// first in order to change the catalog name.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	SchemaName types.String `tfsdk:"schema_name"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	TableNamePrefix types.String `tfsdk:"table_name_prefix"`
}

func (to *AiGatewayInferenceTableConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayInferenceTableConfig_SdkV2) {
}

func (to *AiGatewayInferenceTableConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayInferenceTableConfig_SdkV2) {
}

func (m AiGatewayInferenceTableConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["table_name_prefix"] = attrs["table_name_prefix"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayInferenceTableConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayInferenceTableConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayInferenceTableConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayInferenceTableConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":      m.CatalogName,
			"enabled":           m.Enabled,
			"schema_name":       m.SchemaName,
			"table_name_prefix": m.TableNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayInferenceTableConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
	Calls types.Int64 `tfsdk:"calls"`
	// Key field for a rate limit. Currently, 'user', 'user_group,
	// 'service_principal', and 'endpoint' are supported, with 'endpoint' being
	// the default if not specified.
	Key types.String `tfsdk:"key"`
	// Principal field for a user, user group, or service principal to apply
	// rate limiting to. Accepts a user email, group name, or service principal
	// application ID.
	Principal types.String `tfsdk:"principal"`
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	RenewalPeriod types.String `tfsdk:"renewal_period"`
	// Used to specify how many tokens are allowed for a key within the
	// renewal_period.
	Tokens types.Int64 `tfsdk:"tokens"`
}

func (to *AiGatewayRateLimit_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayRateLimit_SdkV2) {
}

func (to *AiGatewayRateLimit_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayRateLimit_SdkV2) {
}

func (m AiGatewayRateLimit_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["calls"] = attrs["calls"].SetOptional()
	attrs["key"] = attrs["key"].SetOptional()
	attrs["principal"] = attrs["principal"].SetOptional()
	attrs["renewal_period"] = attrs["renewal_period"].SetRequired()
	attrs["tokens"] = attrs["tokens"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayRateLimit.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayRateLimit_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayRateLimit_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayRateLimit_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"calls":          m.Calls,
			"key":            m.Key,
			"principal":      m.Principal,
			"renewal_period": m.RenewalPeriod,
			"tokens":         m.Tokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayRateLimit_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"principal":      types.StringType,
			"renewal_period": types.StringType,
			"tokens":         types.Int64Type,
		},
	}
}

type AiGatewayUsageTrackingConfig_SdkV2 struct {
	// Whether to enable usage tracking.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (to *AiGatewayUsageTrackingConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayUsageTrackingConfig_SdkV2) {
}

func (to *AiGatewayUsageTrackingConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AiGatewayUsageTrackingConfig_SdkV2) {
}

func (m AiGatewayUsageTrackingConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayUsageTrackingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AiGatewayUsageTrackingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayUsageTrackingConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AiGatewayUsageTrackingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": m.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AiGatewayUsageTrackingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type AmazonBedrockConfig_SdkV2 struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id_plaintext`. You must
	// provide an API key using one of the following fields: `aws_access_key_id`
	// or `aws_access_key_id_plaintext`.
	AwsAccessKeyId types.String `tfsdk:"aws_access_key_id"`
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	AwsAccessKeyIdPlaintext types.String `tfsdk:"aws_access_key_id_plaintext"`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion types.String `tfsdk:"aws_region"`
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKey types.String `tfsdk:"aws_secret_access_key"`
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKeyPlaintext types.String `tfsdk:"aws_secret_access_key_plaintext"`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider types.String `tfsdk:"bedrock_provider"`
	// ARN of the instance profile that the external model will use to access
	// AWS resources. You must authenticate using an instance profile or access
	// keys. If you prefer to authenticate using access keys, see
	// `aws_access_key_id`, `aws_access_key_id_plaintext`,
	// `aws_secret_access_key` and `aws_secret_access_key_plaintext`.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
}

func (to *AmazonBedrockConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AmazonBedrockConfig_SdkV2) {
}

func (to *AmazonBedrockConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AmazonBedrockConfig_SdkV2) {
}

func (m AmazonBedrockConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["aws_access_key_id"] = attrs["aws_access_key_id"].SetOptional()
	attrs["aws_access_key_id_plaintext"] = attrs["aws_access_key_id_plaintext"].SetOptional()
	attrs["aws_region"] = attrs["aws_region"].SetRequired()
	attrs["aws_secret_access_key"] = attrs["aws_secret_access_key"].SetOptional()
	attrs["aws_secret_access_key_plaintext"] = attrs["aws_secret_access_key_plaintext"].SetOptional()
	attrs["bedrock_provider"] = attrs["bedrock_provider"].SetRequired()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AmazonBedrockConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AmazonBedrockConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AmazonBedrockConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AmazonBedrockConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_access_key_id":               m.AwsAccessKeyId,
			"aws_access_key_id_plaintext":     m.AwsAccessKeyIdPlaintext,
			"aws_region":                      m.AwsRegion,
			"aws_secret_access_key":           m.AwsSecretAccessKey,
			"aws_secret_access_key_plaintext": m.AwsSecretAccessKeyPlaintext,
			"bedrock_provider":                m.BedrockProvider,
			"instance_profile_arn":            m.InstanceProfileArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AmazonBedrockConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws_access_key_id":               types.StringType,
			"aws_access_key_id_plaintext":     types.StringType,
			"aws_region":                      types.StringType,
			"aws_secret_access_key":           types.StringType,
			"aws_secret_access_key_plaintext": types.StringType,
			"bedrock_provider":                types.StringType,
			"instance_profile_arn":            types.StringType,
		},
	}
}

type AnthropicConfig_SdkV2 struct {
	// The Databricks secret key reference for an Anthropic API key. If you
	// prefer to paste your API key directly, see `anthropic_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKey types.String `tfsdk:"anthropic_api_key"`
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKeyPlaintext types.String `tfsdk:"anthropic_api_key_plaintext"`
}

func (to *AnthropicConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AnthropicConfig_SdkV2) {
}

func (to *AnthropicConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AnthropicConfig_SdkV2) {
}

func (m AnthropicConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anthropic_api_key"] = attrs["anthropic_api_key"].SetOptional()
	attrs["anthropic_api_key_plaintext"] = attrs["anthropic_api_key_plaintext"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AnthropicConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AnthropicConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnthropicConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m AnthropicConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anthropic_api_key":           m.AnthropicApiKey,
			"anthropic_api_key_plaintext": m.AnthropicApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AnthropicConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anthropic_api_key":           types.StringType,
			"anthropic_api_key_plaintext": types.StringType,
		},
	}
}

type ApiKeyAuth_SdkV2 struct {
	// The name of the API key parameter used for authentication.
	Key types.String `tfsdk:"key"`
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	Value types.String `tfsdk:"value"`
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	ValuePlaintext types.String `tfsdk:"value_plaintext"`
}

func (to *ApiKeyAuth_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApiKeyAuth_SdkV2) {
}

func (to *ApiKeyAuth_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ApiKeyAuth_SdkV2) {
}

func (m ApiKeyAuth_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()
	attrs["value_plaintext"] = attrs["value_plaintext"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ApiKeyAuth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ApiKeyAuth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApiKeyAuth_SdkV2
// only implements ToObjectValue() and Type().
func (m ApiKeyAuth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":             m.Key,
			"value":           m.Value,
			"value_plaintext": m.ValuePlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ApiKeyAuth_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":             types.StringType,
			"value":           types.StringType,
			"value_plaintext": types.StringType,
		},
	}
}

type AutoCaptureConfigInput_SdkV2 struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName types.String `tfsdk:"schema_name"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix types.String `tfsdk:"table_name_prefix"`
}

func (to *AutoCaptureConfigInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoCaptureConfigInput_SdkV2) {
}

func (to *AutoCaptureConfigInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AutoCaptureConfigInput_SdkV2) {
}

func (m AutoCaptureConfigInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["table_name_prefix"] = attrs["table_name_prefix"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureConfigInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AutoCaptureConfigInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureConfigInput_SdkV2
// only implements ToObjectValue() and Type().
func (m AutoCaptureConfigInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":      m.CatalogName,
			"enabled":           m.Enabled,
			"schema_name":       m.SchemaName,
			"table_name_prefix": m.TableNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutoCaptureConfigInput_SdkV2) Type(ctx context.Context) attr.Type {
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
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName types.String `tfsdk:"schema_name"`

	State types.List `tfsdk:"state"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix types.String `tfsdk:"table_name_prefix"`
}

func (to *AutoCaptureConfigOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoCaptureConfigOutput_SdkV2) {
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				// Recursively sync the fields of State
				toState.SyncFieldsDuringCreateOrUpdate(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (to *AutoCaptureConfigOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AutoCaptureConfigOutput_SdkV2) {
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (m AutoCaptureConfigOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state"] = attrs["state"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["table_name_prefix"] = attrs["table_name_prefix"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureConfigOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AutoCaptureConfigOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"state": reflect.TypeOf(AutoCaptureState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureConfigOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m AutoCaptureConfigOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"catalog_name":      m.CatalogName,
			"enabled":           m.Enabled,
			"schema_name":       m.SchemaName,
			"state":             m.State,
			"table_name_prefix": m.TableNamePrefix,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutoCaptureConfigOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"enabled":      types.BoolType,
			"schema_name":  types.StringType,
			"state": basetypes.ListType{
				ElemType: AutoCaptureState_SdkV2{}.Type(ctx),
			},
			"table_name_prefix": types.StringType,
		},
	}
}

// GetState returns the value of the State field in AutoCaptureConfigOutput_SdkV2 as
// a AutoCaptureState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoCaptureConfigOutput_SdkV2) GetState(ctx context.Context) (AutoCaptureState_SdkV2, bool) {
	var e AutoCaptureState_SdkV2
	if m.State.IsNull() || m.State.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureState_SdkV2
	d := m.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in AutoCaptureConfigOutput_SdkV2.
func (m *AutoCaptureConfigOutput_SdkV2) SetState(ctx context.Context, v AutoCaptureState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	m.State = types.ListValueMust(t, vs)
}

type AutoCaptureState_SdkV2 struct {
	PayloadTable types.List `tfsdk:"payload_table"`
}

func (to *AutoCaptureState_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoCaptureState_SdkV2) {
	if !from.PayloadTable.IsNull() && !from.PayloadTable.IsUnknown() {
		if toPayloadTable, ok := to.GetPayloadTable(ctx); ok {
			if fromPayloadTable, ok := from.GetPayloadTable(ctx); ok {
				// Recursively sync the fields of PayloadTable
				toPayloadTable.SyncFieldsDuringCreateOrUpdate(ctx, fromPayloadTable)
				to.SetPayloadTable(ctx, toPayloadTable)
			}
		}
	}
}

func (to *AutoCaptureState_SdkV2) SyncFieldsDuringRead(ctx context.Context, from AutoCaptureState_SdkV2) {
	if !from.PayloadTable.IsNull() && !from.PayloadTable.IsUnknown() {
		if toPayloadTable, ok := to.GetPayloadTable(ctx); ok {
			if fromPayloadTable, ok := from.GetPayloadTable(ctx); ok {
				toPayloadTable.SyncFieldsDuringRead(ctx, fromPayloadTable)
				to.SetPayloadTable(ctx, toPayloadTable)
			}
		}
	}
}

func (m AutoCaptureState_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["payload_table"] = attrs["payload_table"].SetOptional()
	attrs["payload_table"] = attrs["payload_table"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m AutoCaptureState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"payload_table": reflect.TypeOf(PayloadTable_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureState_SdkV2
// only implements ToObjectValue() and Type().
func (m AutoCaptureState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"payload_table": m.PayloadTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (m AutoCaptureState_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"payload_table": basetypes.ListType{
				ElemType: PayloadTable_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPayloadTable returns the value of the PayloadTable field in AutoCaptureState_SdkV2 as
// a PayloadTable_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *AutoCaptureState_SdkV2) GetPayloadTable(ctx context.Context) (PayloadTable_SdkV2, bool) {
	var e PayloadTable_SdkV2
	if m.PayloadTable.IsNull() || m.PayloadTable.IsUnknown() {
		return e, false
	}
	var v []PayloadTable_SdkV2
	d := m.PayloadTable.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPayloadTable sets the value of the PayloadTable field in AutoCaptureState_SdkV2.
func (m *AutoCaptureState_SdkV2) SetPayloadTable(ctx context.Context, v PayloadTable_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["payload_table"]
	m.PayloadTable = types.ListValueMust(t, vs)
}

type BearerTokenAuth_SdkV2 struct {
	// The Databricks secret key reference for a token. If you prefer to paste
	// your token directly, see `token_plaintext`.
	Token types.String `tfsdk:"token"`
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	TokenPlaintext types.String `tfsdk:"token_plaintext"`
}

func (to *BearerTokenAuth_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BearerTokenAuth_SdkV2) {
}

func (to *BearerTokenAuth_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BearerTokenAuth_SdkV2) {
}

func (m BearerTokenAuth_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["token"] = attrs["token"].SetOptional()
	attrs["token_plaintext"] = attrs["token_plaintext"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BearerTokenAuth.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BearerTokenAuth_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BearerTokenAuth_SdkV2
// only implements ToObjectValue() and Type().
func (m BearerTokenAuth_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token":           m.Token,
			"token_plaintext": m.TokenPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BearerTokenAuth_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token":           types.StringType,
			"token_plaintext": types.StringType,
		},
	}
}

type BuildLogsRequest_SdkV2 struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (to *BuildLogsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BuildLogsRequest_SdkV2) {
}

func (to *BuildLogsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BuildLogsRequest_SdkV2) {
}

func (m BuildLogsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["served_model_name"] = attrs["served_model_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BuildLogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BuildLogsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BuildLogsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m BuildLogsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              m.Name,
			"served_model_name": m.ServedModelName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BuildLogsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

type BuildLogsResponse_SdkV2 struct {
	// The logs associated with building the served entity's environment.
	Logs types.String `tfsdk:"logs"`
}

func (to *BuildLogsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BuildLogsResponse_SdkV2) {
}

func (to *BuildLogsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from BuildLogsResponse_SdkV2) {
}

func (m BuildLogsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["logs"] = attrs["logs"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in BuildLogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m BuildLogsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BuildLogsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m BuildLogsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"logs": m.Logs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m BuildLogsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ChatMessage_SdkV2 struct {
	// The content of the message.
	Content types.String `tfsdk:"content"`
	// The role of the message. One of [system, user, assistant].
	Role types.String `tfsdk:"role"`
}

func (to *ChatMessage_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ChatMessage_SdkV2) {
}

func (to *ChatMessage_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ChatMessage_SdkV2) {
}

func (m ChatMessage_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["content"] = attrs["content"].SetOptional()
	attrs["role"] = attrs["role"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ChatMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ChatMessage_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChatMessage_SdkV2
// only implements ToObjectValue() and Type().
func (m ChatMessage_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": m.Content,
			"role":    m.Role,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ChatMessage_SdkV2) Type(ctx context.Context) attr.Type {
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
	CohereApiBase types.String `tfsdk:"cohere_api_base"`
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	CohereApiKey types.String `tfsdk:"cohere_api_key"`
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	CohereApiKeyPlaintext types.String `tfsdk:"cohere_api_key_plaintext"`
}

func (to *CohereConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CohereConfig_SdkV2) {
}

func (to *CohereConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CohereConfig_SdkV2) {
}

func (m CohereConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["cohere_api_base"] = attrs["cohere_api_base"].SetOptional()
	attrs["cohere_api_key"] = attrs["cohere_api_key"].SetOptional()
	attrs["cohere_api_key_plaintext"] = attrs["cohere_api_key_plaintext"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CohereConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CohereConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CohereConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m CohereConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cohere_api_base":          m.CohereApiBase,
			"cohere_api_key":           m.CohereApiKey,
			"cohere_api_key_plaintext": m.CohereApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CohereConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cohere_api_base":          types.StringType,
			"cohere_api_key":           types.StringType,
			"cohere_api_key_plaintext": types.StringType,
		},
	}
}

type CreatePtEndpointRequest_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint.
	AiGateway types.List `tfsdk:"ai_gateway"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The core config of the serving endpoint.
	Config types.List `tfsdk:"config"`
	// Email notification settings.
	EmailNotifications types.List `tfsdk:"email_notifications"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name types.String `tfsdk:"name"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreatePtEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePtEndpointRequest_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				// Recursively sync the fields of AiGateway
				toAiGateway.SyncFieldsDuringCreateOrUpdate(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				// Recursively sync the fields of EmailNotifications
				toEmailNotifications.SyncFieldsDuringCreateOrUpdate(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreatePtEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreatePtEndpointRequest_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				toAiGateway.SyncFieldsDuringRead(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreatePtEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["ai_gateway"] = attrs["ai_gateway"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetRequired()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreatePtEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreatePtEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":          reflect.TypeOf(AiGatewayConfig_SdkV2{}),
		"config":              reflect.TypeOf(PtEndpointCoreConfig_SdkV2{}),
		"email_notifications": reflect.TypeOf(EmailNotifications_SdkV2{}),
		"tags":                reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePtEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m CreatePtEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":          m.AiGateway,
			"budget_policy_id":    m.BudgetPolicyId,
			"config":              m.Config,
			"email_notifications": m.EmailNotifications,
			"name":                m.Name,
			"tags":                m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreatePtEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig_SdkV2{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"config": basetypes.ListType{
				ElemType: PtEndpointCoreConfig_SdkV2{}.Type(ctx),
			},
			"email_notifications": basetypes.ListType{
				ElemType: EmailNotifications_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in CreatePtEndpointRequest_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePtEndpointRequest_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if m.AiGateway.IsNull() || m.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := m.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in CreatePtEndpointRequest_SdkV2.
func (m *CreatePtEndpointRequest_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	m.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in CreatePtEndpointRequest_SdkV2 as
// a PtEndpointCoreConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePtEndpointRequest_SdkV2) GetConfig(ctx context.Context) (PtEndpointCoreConfig_SdkV2, bool) {
	var e PtEndpointCoreConfig_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []PtEndpointCoreConfig_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in CreatePtEndpointRequest_SdkV2.
func (m *CreatePtEndpointRequest_SdkV2) SetConfig(ctx context.Context, v PtEndpointCoreConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in CreatePtEndpointRequest_SdkV2 as
// a EmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePtEndpointRequest_SdkV2) GetEmailNotifications(ctx context.Context) (EmailNotifications_SdkV2, bool) {
	var e EmailNotifications_SdkV2
	if m.EmailNotifications.IsNull() || m.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []EmailNotifications_SdkV2
	d := m.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in CreatePtEndpointRequest_SdkV2.
func (m *CreatePtEndpointRequest_SdkV2) SetEmailNotifications(ctx context.Context, v EmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	m.EmailNotifications = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreatePtEndpointRequest_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreatePtEndpointRequest_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreatePtEndpointRequest_SdkV2.
func (m *CreatePtEndpointRequest_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type CreateServingEndpoint_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway types.List `tfsdk:"ai_gateway"`
	// The budget policy to be applied to the serving endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The core config of the serving endpoint.
	Config types.List `tfsdk:"config"`

	Description types.String `tfsdk:"description"`
	// Email notification settings.
	EmailNotifications types.List `tfsdk:"email_notifications"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name types.String `tfsdk:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized types.Bool `tfsdk:"route_optimized"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreateServingEndpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServingEndpoint_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				// Recursively sync the fields of AiGateway
				toAiGateway.SyncFieldsDuringCreateOrUpdate(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				// Recursively sync the fields of EmailNotifications
				toEmailNotifications.SyncFieldsDuringCreateOrUpdate(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *CreateServingEndpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CreateServingEndpoint_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				toAiGateway.SyncFieldsDuringRead(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m CreateServingEndpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["ai_gateway"] = attrs["ai_gateway"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["route_optimized"] = attrs["route_optimized"].SetOptional()
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateServingEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CreateServingEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":          reflect.TypeOf(AiGatewayConfig_SdkV2{}),
		"config":              reflect.TypeOf(EndpointCoreConfigInput_SdkV2{}),
		"email_notifications": reflect.TypeOf(EmailNotifications_SdkV2{}),
		"rate_limits":         reflect.TypeOf(RateLimit_SdkV2{}),
		"tags":                reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServingEndpoint_SdkV2
// only implements ToObjectValue() and Type().
func (m CreateServingEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":          m.AiGateway,
			"budget_policy_id":    m.BudgetPolicyId,
			"config":              m.Config,
			"description":         m.Description,
			"email_notifications": m.EmailNotifications,
			"name":                m.Name,
			"rate_limits":         m.RateLimits,
			"route_optimized":     m.RouteOptimized,
			"tags":                m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CreateServingEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig_SdkV2{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigInput_SdkV2{}.Type(ctx),
			},
			"description": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: EmailNotifications_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit_SdkV2{}.Type(ctx),
			},
			"route_optimized": types.BoolType,
			"tags": basetypes.ListType{
				ElemType: EndpointTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in CreateServingEndpoint_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServingEndpoint_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if m.AiGateway.IsNull() || m.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := m.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in CreateServingEndpoint_SdkV2.
func (m *CreateServingEndpoint_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	m.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in CreateServingEndpoint_SdkV2 as
// a EndpointCoreConfigInput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServingEndpoint_SdkV2) GetConfig(ctx context.Context) (EndpointCoreConfigInput_SdkV2, bool) {
	var e EndpointCoreConfigInput_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []EndpointCoreConfigInput_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in CreateServingEndpoint_SdkV2.
func (m *CreateServingEndpoint_SdkV2) SetConfig(ctx context.Context, v EndpointCoreConfigInput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in CreateServingEndpoint_SdkV2 as
// a EmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServingEndpoint_SdkV2) GetEmailNotifications(ctx context.Context) (EmailNotifications_SdkV2, bool) {
	var e EmailNotifications_SdkV2
	if m.EmailNotifications.IsNull() || m.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []EmailNotifications_SdkV2
	d := m.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in CreateServingEndpoint_SdkV2.
func (m *CreateServingEndpoint_SdkV2) SetEmailNotifications(ctx context.Context, v EmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	m.EmailNotifications = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in CreateServingEndpoint_SdkV2 as
// a slice of RateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServingEndpoint_SdkV2) GetRateLimits(ctx context.Context) ([]RateLimit_SdkV2, bool) {
	if m.RateLimits.IsNull() || m.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit_SdkV2
	d := m.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in CreateServingEndpoint_SdkV2.
func (m *CreateServingEndpoint_SdkV2) SetRateLimits(ctx context.Context, v []RateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RateLimits = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateServingEndpoint_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *CreateServingEndpoint_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateServingEndpoint_SdkV2.
func (m *CreateServingEndpoint_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

// Configs needed to create a custom provider model route.
type CustomProviderConfig_SdkV2 struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	ApiKeyAuth types.List `tfsdk:"api_key_auth"`
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	BearerTokenAuth types.List `tfsdk:"bearer_token_auth"`
	// This is a field to provide the URL of the custom provider API.
	CustomProviderUrl types.String `tfsdk:"custom_provider_url"`
}

func (to *CustomProviderConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomProviderConfig_SdkV2) {
	if !from.ApiKeyAuth.IsNull() && !from.ApiKeyAuth.IsUnknown() {
		if toApiKeyAuth, ok := to.GetApiKeyAuth(ctx); ok {
			if fromApiKeyAuth, ok := from.GetApiKeyAuth(ctx); ok {
				// Recursively sync the fields of ApiKeyAuth
				toApiKeyAuth.SyncFieldsDuringCreateOrUpdate(ctx, fromApiKeyAuth)
				to.SetApiKeyAuth(ctx, toApiKeyAuth)
			}
		}
	}
	if !from.BearerTokenAuth.IsNull() && !from.BearerTokenAuth.IsUnknown() {
		if toBearerTokenAuth, ok := to.GetBearerTokenAuth(ctx); ok {
			if fromBearerTokenAuth, ok := from.GetBearerTokenAuth(ctx); ok {
				// Recursively sync the fields of BearerTokenAuth
				toBearerTokenAuth.SyncFieldsDuringCreateOrUpdate(ctx, fromBearerTokenAuth)
				to.SetBearerTokenAuth(ctx, toBearerTokenAuth)
			}
		}
	}
}

func (to *CustomProviderConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from CustomProviderConfig_SdkV2) {
	if !from.ApiKeyAuth.IsNull() && !from.ApiKeyAuth.IsUnknown() {
		if toApiKeyAuth, ok := to.GetApiKeyAuth(ctx); ok {
			if fromApiKeyAuth, ok := from.GetApiKeyAuth(ctx); ok {
				toApiKeyAuth.SyncFieldsDuringRead(ctx, fromApiKeyAuth)
				to.SetApiKeyAuth(ctx, toApiKeyAuth)
			}
		}
	}
	if !from.BearerTokenAuth.IsNull() && !from.BearerTokenAuth.IsUnknown() {
		if toBearerTokenAuth, ok := to.GetBearerTokenAuth(ctx); ok {
			if fromBearerTokenAuth, ok := from.GetBearerTokenAuth(ctx); ok {
				toBearerTokenAuth.SyncFieldsDuringRead(ctx, fromBearerTokenAuth)
				to.SetBearerTokenAuth(ctx, toBearerTokenAuth)
			}
		}
	}
}

func (m CustomProviderConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["api_key_auth"] = attrs["api_key_auth"].SetOptional()
	attrs["api_key_auth"] = attrs["api_key_auth"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["bearer_token_auth"] = attrs["bearer_token_auth"].SetOptional()
	attrs["bearer_token_auth"] = attrs["bearer_token_auth"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_provider_url"] = attrs["custom_provider_url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CustomProviderConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m CustomProviderConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"api_key_auth":      reflect.TypeOf(ApiKeyAuth_SdkV2{}),
		"bearer_token_auth": reflect.TypeOf(BearerTokenAuth_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomProviderConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m CustomProviderConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"api_key_auth":        m.ApiKeyAuth,
			"bearer_token_auth":   m.BearerTokenAuth,
			"custom_provider_url": m.CustomProviderUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m CustomProviderConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"api_key_auth": basetypes.ListType{
				ElemType: ApiKeyAuth_SdkV2{}.Type(ctx),
			},
			"bearer_token_auth": basetypes.ListType{
				ElemType: BearerTokenAuth_SdkV2{}.Type(ctx),
			},
			"custom_provider_url": types.StringType,
		},
	}
}

// GetApiKeyAuth returns the value of the ApiKeyAuth field in CustomProviderConfig_SdkV2 as
// a ApiKeyAuth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomProviderConfig_SdkV2) GetApiKeyAuth(ctx context.Context) (ApiKeyAuth_SdkV2, bool) {
	var e ApiKeyAuth_SdkV2
	if m.ApiKeyAuth.IsNull() || m.ApiKeyAuth.IsUnknown() {
		return e, false
	}
	var v []ApiKeyAuth_SdkV2
	d := m.ApiKeyAuth.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetApiKeyAuth sets the value of the ApiKeyAuth field in CustomProviderConfig_SdkV2.
func (m *CustomProviderConfig_SdkV2) SetApiKeyAuth(ctx context.Context, v ApiKeyAuth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["api_key_auth"]
	m.ApiKeyAuth = types.ListValueMust(t, vs)
}

// GetBearerTokenAuth returns the value of the BearerTokenAuth field in CustomProviderConfig_SdkV2 as
// a BearerTokenAuth_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *CustomProviderConfig_SdkV2) GetBearerTokenAuth(ctx context.Context) (BearerTokenAuth_SdkV2, bool) {
	var e BearerTokenAuth_SdkV2
	if m.BearerTokenAuth.IsNull() || m.BearerTokenAuth.IsUnknown() {
		return e, false
	}
	var v []BearerTokenAuth_SdkV2
	d := m.BearerTokenAuth.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetBearerTokenAuth sets the value of the BearerTokenAuth field in CustomProviderConfig_SdkV2.
func (m *CustomProviderConfig_SdkV2) SetBearerTokenAuth(ctx context.Context, v BearerTokenAuth_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["bearer_token_auth"]
	m.BearerTokenAuth = types.ListValueMust(t, vs)
}

// Details necessary to query this object's API through the DataPlane APIs.
type DataPlaneInfo_SdkV2 struct {
	// Authorization details as a string.
	AuthorizationDetails types.String `tfsdk:"authorization_details"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl types.String `tfsdk:"endpoint_url"`
}

func (to *DataPlaneInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataPlaneInfo_SdkV2) {
}

func (to *DataPlaneInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataPlaneInfo_SdkV2) {
}

func (m DataPlaneInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (m DataPlaneInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataPlaneInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m DataPlaneInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"authorization_details": m.AuthorizationDetails,
			"endpoint_url":          m.EndpointUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataPlaneInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"authorization_details": types.StringType,
			"endpoint_url":          types.StringType,
		},
	}
}

type DatabricksModelServingConfig_SdkV2 struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiToken types.String `tfsdk:"databricks_api_token"`
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiTokenPlaintext types.String `tfsdk:"databricks_api_token_plaintext"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl types.String `tfsdk:"databricks_workspace_url"`
}

func (to *DatabricksModelServingConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksModelServingConfig_SdkV2) {
}

func (to *DatabricksModelServingConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DatabricksModelServingConfig_SdkV2) {
}

func (m DatabricksModelServingConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["databricks_api_token"] = attrs["databricks_api_token"].SetOptional()
	attrs["databricks_api_token_plaintext"] = attrs["databricks_api_token_plaintext"].SetOptional()
	attrs["databricks_workspace_url"] = attrs["databricks_workspace_url"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DatabricksModelServingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DatabricksModelServingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksModelServingConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m DatabricksModelServingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"databricks_api_token":           m.DatabricksApiToken,
			"databricks_api_token_plaintext": m.DatabricksApiTokenPlaintext,
			"databricks_workspace_url":       m.DatabricksWorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DatabricksModelServingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"databricks_api_token":           types.StringType,
			"databricks_api_token_plaintext": types.StringType,
			"databricks_workspace_url":       types.StringType,
		},
	}
}

type DataframeSplitInput_SdkV2 struct {
	// Columns array for the dataframe
	Columns types.List `tfsdk:"columns"`
	// Data array for the dataframe
	Data types.List `tfsdk:"data"`
	// Index array for the dataframe
	Index types.List `tfsdk:"index"`
}

func (to *DataframeSplitInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataframeSplitInput_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() && to.Index.IsNull() && len(from.Index.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Index, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Index = from.Index
	}
}

func (to *DataframeSplitInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DataframeSplitInput_SdkV2) {
	if !from.Columns.IsNull() && !from.Columns.IsUnknown() && to.Columns.IsNull() && len(from.Columns.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Columns, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Columns = from.Columns
	}
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
	if !from.Index.IsNull() && !from.Index.IsUnknown() && to.Index.IsNull() && len(from.Index.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Index, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Index = from.Index
	}
}

func (m DataframeSplitInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["columns"] = attrs["columns"].SetOptional()
	attrs["data"] = attrs["data"].SetOptional()
	attrs["index"] = attrs["index"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DataframeSplitInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DataframeSplitInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(types.Object{}),
		"data":    reflect.TypeOf(types.Object{}),
		"index":   reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataframeSplitInput_SdkV2
// only implements ToObjectValue() and Type().
func (m DataframeSplitInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": m.Columns,
			"data":    m.Data,
			"index":   m.Index,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DataframeSplitInput_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *DataframeSplitInput_SdkV2) GetColumns(ctx context.Context) ([]types.Object, bool) {
	if m.Columns.IsNull() || m.Columns.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Columns.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetColumns sets the value of the Columns field in DataframeSplitInput_SdkV2.
func (m *DataframeSplitInput_SdkV2) SetColumns(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Columns = types.ListValueMust(t, vs)
}

// GetData returns the value of the Data field in DataframeSplitInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataframeSplitInput_SdkV2) GetData(ctx context.Context) ([]types.Object, bool) {
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in DataframeSplitInput_SdkV2.
func (m *DataframeSplitInput_SdkV2) SetData(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Data = types.ListValueMust(t, vs)
}

// GetIndex returns the value of the Index field in DataframeSplitInput_SdkV2 as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (m *DataframeSplitInput_SdkV2) GetIndex(ctx context.Context) ([]types.Int64, bool) {
	if m.Index.IsNull() || m.Index.IsUnknown() {
		return nil, false
	}
	var v []types.Int64
	d := m.Index.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetIndex sets the value of the Index field in DataframeSplitInput_SdkV2.
func (m *DataframeSplitInput_SdkV2) SetIndex(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["index"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Index = types.ListValueMust(t, vs)
}

type DeleteServingEndpointRequest_SdkV2 struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteServingEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServingEndpointRequest_SdkV2) {
}

func (to *DeleteServingEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from DeleteServingEndpointRequest_SdkV2) {
}

func (m DeleteServingEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteServingEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m DeleteServingEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServingEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m DeleteServingEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m DeleteServingEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type EmailNotifications_SdkV2 struct {
	// A list of email addresses to be notified when an endpoint fails to update
	// its configuration or state.
	OnUpdateFailure types.List `tfsdk:"on_update_failure"`
	// A list of email addresses to be notified when an endpoint successfully
	// updates its configuration or state.
	OnUpdateSuccess types.List `tfsdk:"on_update_success"`
}

func (to *EmailNotifications_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmailNotifications_SdkV2) {
	if !from.OnUpdateFailure.IsNull() && !from.OnUpdateFailure.IsUnknown() && to.OnUpdateFailure.IsNull() && len(from.OnUpdateFailure.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnUpdateFailure, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnUpdateFailure = from.OnUpdateFailure
	}
	if !from.OnUpdateSuccess.IsNull() && !from.OnUpdateSuccess.IsUnknown() && to.OnUpdateSuccess.IsNull() && len(from.OnUpdateSuccess.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnUpdateSuccess, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnUpdateSuccess = from.OnUpdateSuccess
	}
}

func (to *EmailNotifications_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmailNotifications_SdkV2) {
	if !from.OnUpdateFailure.IsNull() && !from.OnUpdateFailure.IsUnknown() && to.OnUpdateFailure.IsNull() && len(from.OnUpdateFailure.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnUpdateFailure, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnUpdateFailure = from.OnUpdateFailure
	}
	if !from.OnUpdateSuccess.IsNull() && !from.OnUpdateSuccess.IsUnknown() && to.OnUpdateSuccess.IsNull() && len(from.OnUpdateSuccess.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for OnUpdateSuccess, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.OnUpdateSuccess = from.OnUpdateSuccess
	}
}

func (m EmailNotifications_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["on_update_failure"] = attrs["on_update_failure"].SetOptional()
	attrs["on_update_success"] = attrs["on_update_success"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmailNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmailNotifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_update_failure": reflect.TypeOf(types.String{}),
		"on_update_success": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmailNotifications_SdkV2
// only implements ToObjectValue() and Type().
func (m EmailNotifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_update_failure": m.OnUpdateFailure,
			"on_update_success": m.OnUpdateSuccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmailNotifications_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"on_update_failure": basetypes.ListType{
				ElemType: types.StringType,
			},
			"on_update_success": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetOnUpdateFailure returns the value of the OnUpdateFailure field in EmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EmailNotifications_SdkV2) GetOnUpdateFailure(ctx context.Context) ([]types.String, bool) {
	if m.OnUpdateFailure.IsNull() || m.OnUpdateFailure.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OnUpdateFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnUpdateFailure sets the value of the OnUpdateFailure field in EmailNotifications_SdkV2.
func (m *EmailNotifications_SdkV2) SetOnUpdateFailure(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["on_update_failure"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OnUpdateFailure = types.ListValueMust(t, vs)
}

// GetOnUpdateSuccess returns the value of the OnUpdateSuccess field in EmailNotifications_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *EmailNotifications_SdkV2) GetOnUpdateSuccess(ctx context.Context) ([]types.String, bool) {
	if m.OnUpdateSuccess.IsNull() || m.OnUpdateSuccess.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.OnUpdateSuccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnUpdateSuccess sets the value of the OnUpdateSuccess field in EmailNotifications_SdkV2.
func (m *EmailNotifications_SdkV2) SetOnUpdateSuccess(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["on_update_success"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.OnUpdateSuccess = types.ListValueMust(t, vs)
}

type EmbeddingsV1ResponseEmbeddingElement_SdkV2 struct {
	// The embedding vector
	Embedding types.List `tfsdk:"embedding"`
	// The index of the embedding in the response.
	Index types.Int64 `tfsdk:"index"`
	// This will always be 'embedding'.
	Object types.String `tfsdk:"object"`
}

func (to *EmbeddingsV1ResponseEmbeddingElement_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingsV1ResponseEmbeddingElement_SdkV2) {
	if !from.Embedding.IsNull() && !from.Embedding.IsUnknown() && to.Embedding.IsNull() && len(from.Embedding.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Embedding, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Embedding = from.Embedding
	}
}

func (to *EmbeddingsV1ResponseEmbeddingElement_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EmbeddingsV1ResponseEmbeddingElement_SdkV2) {
	if !from.Embedding.IsNull() && !from.Embedding.IsUnknown() && to.Embedding.IsNull() && len(from.Embedding.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Embedding, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Embedding = from.Embedding
	}
}

func (m EmbeddingsV1ResponseEmbeddingElement_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["embedding"] = attrs["embedding"].SetOptional()
	attrs["index"] = attrs["index"].SetOptional()
	attrs["object"] = attrs["object"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EmbeddingsV1ResponseEmbeddingElement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EmbeddingsV1ResponseEmbeddingElement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding": reflect.TypeOf(types.Float64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingsV1ResponseEmbeddingElement_SdkV2
// only implements ToObjectValue() and Type().
func (m EmbeddingsV1ResponseEmbeddingElement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding": m.Embedding,
			"index":     m.Index,
			"object":    m.Object,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EmbeddingsV1ResponseEmbeddingElement_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *EmbeddingsV1ResponseEmbeddingElement_SdkV2) GetEmbedding(ctx context.Context) ([]types.Float64, bool) {
	if m.Embedding.IsNull() || m.Embedding.IsUnknown() {
		return nil, false
	}
	var v []types.Float64
	d := m.Embedding.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmbedding sets the value of the Embedding field in EmbeddingsV1ResponseEmbeddingElement_SdkV2.
func (m *EmbeddingsV1ResponseEmbeddingElement_SdkV2) SetEmbedding(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Embedding = types.ListValueMust(t, vs)
}

type EndpointCoreConfigInput_SdkV2 struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig types.List `tfsdk:"auto_capture_config"`
	// The name of the serving endpoint to update. This field is required.
	Name types.String `tfsdk:"-"`
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig types.List `tfsdk:"traffic_config"`
}

func (to *EndpointCoreConfigInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointCoreConfigInput_SdkV2) {
	if !from.AutoCaptureConfig.IsNull() && !from.AutoCaptureConfig.IsUnknown() {
		if toAutoCaptureConfig, ok := to.GetAutoCaptureConfig(ctx); ok {
			if fromAutoCaptureConfig, ok := from.GetAutoCaptureConfig(ctx); ok {
				// Recursively sync the fields of AutoCaptureConfig
				toAutoCaptureConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoCaptureConfig)
				to.SetAutoCaptureConfig(ctx, toAutoCaptureConfig)
			}
		}
	}
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				// Recursively sync the fields of TrafficConfig
				toTrafficConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (to *EndpointCoreConfigInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointCoreConfigInput_SdkV2) {
	if !from.AutoCaptureConfig.IsNull() && !from.AutoCaptureConfig.IsUnknown() {
		if toAutoCaptureConfig, ok := to.GetAutoCaptureConfig(ctx); ok {
			if fromAutoCaptureConfig, ok := from.GetAutoCaptureConfig(ctx); ok {
				toAutoCaptureConfig.SyncFieldsDuringRead(ctx, fromAutoCaptureConfig)
				to.SetAutoCaptureConfig(ctx, toAutoCaptureConfig)
			}
		}
	}
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				toTrafficConfig.SyncFieldsDuringRead(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (m EndpointCoreConfigInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_capture_config"] = attrs["auto_capture_config"].SetOptional()
	attrs["auto_capture_config"] = attrs["auto_capture_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointCoreConfigInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m EndpointCoreConfigInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_capture_config": m.AutoCaptureConfig,
			"name":                m.Name,
			"served_entities":     m.ServedEntities,
			"served_models":       m.ServedModels,
			"traffic_config":      m.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointCoreConfigInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigInput_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityInput_SdkV2{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelInput_SdkV2{}.Type(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointCoreConfigInput_SdkV2 as
// a AutoCaptureConfigInput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigInput_SdkV2) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigInput_SdkV2, bool) {
	var e AutoCaptureConfigInput_SdkV2
	if m.AutoCaptureConfig.IsNull() || m.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureConfigInput_SdkV2
	d := m.AutoCaptureConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointCoreConfigInput_SdkV2.
func (m *EndpointCoreConfigInput_SdkV2) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigInput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_capture_config"]
	m.AutoCaptureConfig = types.ListValueMust(t, vs)
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigInput_SdkV2 as
// a slice of ServedEntityInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigInput_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntityInput_SdkV2, bool) {
	if m.ServedEntities.IsNull() || m.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityInput_SdkV2
	d := m.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigInput_SdkV2.
func (m *EndpointCoreConfigInput_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntityInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigInput_SdkV2 as
// a slice of ServedModelInput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigInput_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelInput_SdkV2, bool) {
	if m.ServedModels.IsNull() || m.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelInput_SdkV2
	d := m.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigInput_SdkV2.
func (m *EndpointCoreConfigInput_SdkV2) SetServedModels(ctx context.Context, v []ServedModelInput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointCoreConfigInput_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigInput_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if m.TrafficConfig.IsNull() || m.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := m.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointCoreConfigInput_SdkV2.
func (m *EndpointCoreConfigInput_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	m.TrafficConfig = types.ListValueMust(t, vs)
}

type EndpointCoreConfigOutput_SdkV2 struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig types.List `tfsdk:"auto_capture_config"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version"`
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig types.List `tfsdk:"traffic_config"`
}

func (to *EndpointCoreConfigOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointCoreConfigOutput_SdkV2) {
	if !from.AutoCaptureConfig.IsNull() && !from.AutoCaptureConfig.IsUnknown() {
		if toAutoCaptureConfig, ok := to.GetAutoCaptureConfig(ctx); ok {
			if fromAutoCaptureConfig, ok := from.GetAutoCaptureConfig(ctx); ok {
				// Recursively sync the fields of AutoCaptureConfig
				toAutoCaptureConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoCaptureConfig)
				to.SetAutoCaptureConfig(ctx, toAutoCaptureConfig)
			}
		}
	}
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				// Recursively sync the fields of TrafficConfig
				toTrafficConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (to *EndpointCoreConfigOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointCoreConfigOutput_SdkV2) {
	if !from.AutoCaptureConfig.IsNull() && !from.AutoCaptureConfig.IsUnknown() {
		if toAutoCaptureConfig, ok := to.GetAutoCaptureConfig(ctx); ok {
			if fromAutoCaptureConfig, ok := from.GetAutoCaptureConfig(ctx); ok {
				toAutoCaptureConfig.SyncFieldsDuringRead(ctx, fromAutoCaptureConfig)
				to.SetAutoCaptureConfig(ctx, toAutoCaptureConfig)
			}
		}
	}
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				toTrafficConfig.SyncFieldsDuringRead(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (m EndpointCoreConfigOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_capture_config"] = attrs["auto_capture_config"].SetOptional()
	attrs["auto_capture_config"] = attrs["auto_capture_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["config_version"] = attrs["config_version"].SetOptional()
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointCoreConfigOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m EndpointCoreConfigOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_capture_config": m.AutoCaptureConfig,
			"config_version":      m.ConfigVersion,
			"served_entities":     m.ServedEntities,
			"served_models":       m.ServedModels,
			"traffic_config":      m.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointCoreConfigOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigOutput_SdkV2{}.Type(ctx),
			},
			"config_version": types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput_SdkV2{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput_SdkV2{}.Type(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointCoreConfigOutput_SdkV2 as
// a AutoCaptureConfigOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigOutput_SdkV2) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigOutput_SdkV2, bool) {
	var e AutoCaptureConfigOutput_SdkV2
	if m.AutoCaptureConfig.IsNull() || m.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureConfigOutput_SdkV2
	d := m.AutoCaptureConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointCoreConfigOutput_SdkV2.
func (m *EndpointCoreConfigOutput_SdkV2) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_capture_config"]
	m.AutoCaptureConfig = types.ListValueMust(t, vs)
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigOutput_SdkV2 as
// a slice of ServedEntityOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigOutput_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntityOutput_SdkV2, bool) {
	if m.ServedEntities.IsNull() || m.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityOutput_SdkV2
	d := m.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigOutput_SdkV2.
func (m *EndpointCoreConfigOutput_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntityOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigOutput_SdkV2 as
// a slice of ServedModelOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigOutput_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelOutput_SdkV2, bool) {
	if m.ServedModels.IsNull() || m.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelOutput_SdkV2
	d := m.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigOutput_SdkV2.
func (m *EndpointCoreConfigOutput_SdkV2) SetServedModels(ctx context.Context, v []ServedModelOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointCoreConfigOutput_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigOutput_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if m.TrafficConfig.IsNull() || m.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := m.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointCoreConfigOutput_SdkV2.
func (m *EndpointCoreConfigOutput_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	m.TrafficConfig = types.ListValueMust(t, vs)
}

type EndpointCoreConfigSummary_SdkV2 struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models"`
}

func (to *EndpointCoreConfigSummary_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointCoreConfigSummary_SdkV2) {
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
}

func (to *EndpointCoreConfigSummary_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointCoreConfigSummary_SdkV2) {
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
}

func (m EndpointCoreConfigSummary_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigSummary.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointCoreConfigSummary_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"served_entities": reflect.TypeOf(ServedEntitySpec_SdkV2{}),
		"served_models":   reflect.TypeOf(ServedModelSpec_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigSummary_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointCoreConfigSummary_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entities": m.ServedEntities,
			"served_models":   m.ServedModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointCoreConfigSummary_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entities": basetypes.ListType{
				ElemType: ServedEntitySpec_SdkV2{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelSpec_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigSummary_SdkV2 as
// a slice of ServedEntitySpec_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigSummary_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntitySpec_SdkV2, bool) {
	if m.ServedEntities.IsNull() || m.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntitySpec_SdkV2
	d := m.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigSummary_SdkV2.
func (m *EndpointCoreConfigSummary_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntitySpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigSummary_SdkV2 as
// a slice of ServedModelSpec_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointCoreConfigSummary_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelSpec_SdkV2, bool) {
	if m.ServedModels.IsNull() || m.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelSpec_SdkV2
	d := m.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigSummary_SdkV2.
func (m *EndpointCoreConfigSummary_SdkV2) SetServedModels(ctx context.Context, v []ServedModelSpec_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedModels = types.ListValueMust(t, vs)
}

type EndpointPendingConfig_SdkV2 struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig types.List `tfsdk:"auto_capture_config"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels types.List `tfsdk:"served_models"`
	// The timestamp when the update to the pending config started.
	StartTime types.Int64 `tfsdk:"start_time"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig types.List `tfsdk:"traffic_config"`
}

func (to *EndpointPendingConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointPendingConfig_SdkV2) {
	if !from.AutoCaptureConfig.IsNull() && !from.AutoCaptureConfig.IsUnknown() {
		if toAutoCaptureConfig, ok := to.GetAutoCaptureConfig(ctx); ok {
			if fromAutoCaptureConfig, ok := from.GetAutoCaptureConfig(ctx); ok {
				// Recursively sync the fields of AutoCaptureConfig
				toAutoCaptureConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAutoCaptureConfig)
				to.SetAutoCaptureConfig(ctx, toAutoCaptureConfig)
			}
		}
	}
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				// Recursively sync the fields of TrafficConfig
				toTrafficConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (to *EndpointPendingConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointPendingConfig_SdkV2) {
	if !from.AutoCaptureConfig.IsNull() && !from.AutoCaptureConfig.IsUnknown() {
		if toAutoCaptureConfig, ok := to.GetAutoCaptureConfig(ctx); ok {
			if fromAutoCaptureConfig, ok := from.GetAutoCaptureConfig(ctx); ok {
				toAutoCaptureConfig.SyncFieldsDuringRead(ctx, fromAutoCaptureConfig)
				to.SetAutoCaptureConfig(ctx, toAutoCaptureConfig)
			}
		}
	}
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.ServedModels.IsNull() && !from.ServedModels.IsUnknown() && to.ServedModels.IsNull() && len(from.ServedModels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedModels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedModels = from.ServedModels
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				toTrafficConfig.SyncFieldsDuringRead(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (m EndpointPendingConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_capture_config"] = attrs["auto_capture_config"].SetOptional()
	attrs["auto_capture_config"] = attrs["auto_capture_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["config_version"] = attrs["config_version"].SetOptional()
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointPendingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointPendingConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m EndpointPendingConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"auto_capture_config": m.AutoCaptureConfig,
			"config_version":      m.ConfigVersion,
			"served_entities":     m.ServedEntities,
			"served_models":       m.ServedModels,
			"start_time":          m.StartTime,
			"traffic_config":      m.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointPendingConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigOutput_SdkV2{}.Type(ctx),
			},
			"config_version": types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput_SdkV2{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput_SdkV2{}.Type(ctx),
			},
			"start_time": types.Int64Type,
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointPendingConfig_SdkV2 as
// a AutoCaptureConfigOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointPendingConfig_SdkV2) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigOutput_SdkV2, bool) {
	var e AutoCaptureConfigOutput_SdkV2
	if m.AutoCaptureConfig.IsNull() || m.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v []AutoCaptureConfigOutput_SdkV2
	d := m.AutoCaptureConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointPendingConfig_SdkV2.
func (m *EndpointPendingConfig_SdkV2) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["auto_capture_config"]
	m.AutoCaptureConfig = types.ListValueMust(t, vs)
}

// GetServedEntities returns the value of the ServedEntities field in EndpointPendingConfig_SdkV2 as
// a slice of ServedEntityOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointPendingConfig_SdkV2) GetServedEntities(ctx context.Context) ([]ServedEntityOutput_SdkV2, bool) {
	if m.ServedEntities.IsNull() || m.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityOutput_SdkV2
	d := m.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointPendingConfig_SdkV2.
func (m *EndpointPendingConfig_SdkV2) SetServedEntities(ctx context.Context, v []ServedEntityOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointPendingConfig_SdkV2 as
// a slice of ServedModelOutput_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointPendingConfig_SdkV2) GetServedModels(ctx context.Context) ([]ServedModelOutput_SdkV2, bool) {
	if m.ServedModels.IsNull() || m.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelOutput_SdkV2
	d := m.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointPendingConfig_SdkV2.
func (m *EndpointPendingConfig_SdkV2) SetServedModels(ctx context.Context, v []ServedModelOutput_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointPendingConfig_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointPendingConfig_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if m.TrafficConfig.IsNull() || m.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := m.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointPendingConfig_SdkV2.
func (m *EndpointPendingConfig_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	m.TrafficConfig = types.ListValueMust(t, vs)
}

type EndpointState_SdkV2 struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails.
	ConfigUpdate types.String `tfsdk:"config_update"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready types.String `tfsdk:"ready"`
}

func (to *EndpointState_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointState_SdkV2) {
}

func (to *EndpointState_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointState_SdkV2) {
}

func (m EndpointState_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config_update"] = attrs["config_update"].SetOptional()
	attrs["ready"] = attrs["ready"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointState_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_update": m.ConfigUpdate,
			"ready":         m.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointState_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_update": types.StringType,
			"ready":         types.StringType,
		},
	}
}

type EndpointTag_SdkV2 struct {
	// Key field for a serving endpoint tag.
	Key types.String `tfsdk:"key"`
	// Optional value field for a serving endpoint tag.
	Value types.String `tfsdk:"value"`
}

func (to *EndpointTag_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTag_SdkV2) {
}

func (to *EndpointTag_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointTag_SdkV2) {
}

func (m EndpointTag_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["key"] = attrs["key"].SetRequired()
	attrs["value"] = attrs["value"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTag.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointTag_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTag_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointTag_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   m.Key,
			"value": m.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointTag_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type EndpointTags_SdkV2 struct {
	Tags types.List `tfsdk:"tags"`
}

func (to *EndpointTags_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTags_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *EndpointTags_SdkV2) SyncFieldsDuringRead(ctx context.Context, from EndpointTags_SdkV2) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m EndpointTags_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["tags"] = attrs["tags"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointTags.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m EndpointTags_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTags_SdkV2
// only implements ToObjectValue() and Type().
func (m EndpointTags_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tags": m.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (m EndpointTags_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tags": basetypes.ListType{
				ElemType: EndpointTag_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in EndpointTags_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *EndpointTags_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in EndpointTags_SdkV2.
func (m *EndpointTags_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ExportMetricsRequest_SdkV2 struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (to *ExportMetricsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportMetricsRequest_SdkV2) {
}

func (to *ExportMetricsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExportMetricsRequest_SdkV2) {
}

func (m ExportMetricsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportMetricsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExportMetricsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportMetricsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExportMetricsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportMetricsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ExportMetricsResponse_SdkV2 struct {
	Contents types.Object `tfsdk:"-"`
}

func (to *ExportMetricsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportMetricsResponse_SdkV2) {
}

func (to *ExportMetricsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExportMetricsResponse_SdkV2) {
}

func (m ExportMetricsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExportMetricsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExportMetricsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportMetricsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ExportMetricsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": m.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExportMetricsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

// Simple Proto message for testing
type ExternalFunctionRequest_SdkV2 struct {
	// The connection name to use. This is required to identify the external
	// connection.
	ConnectionName types.String `tfsdk:"connection_name"`
	// Additional headers for the request. If not provided, only auth headers
	// from connections would be passed.
	Headers types.String `tfsdk:"headers"`
	// The JSON payload to send in the request body.
	Json types.String `tfsdk:"json"`
	// The HTTP method to use (e.g., 'GET', 'POST').
	Method types.String `tfsdk:"method"`
	// Query parameters for the request.
	Params types.String `tfsdk:"params"`
	// The relative path for the API endpoint. This is required.
	Path types.String `tfsdk:"path"`
}

func (to *ExternalFunctionRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalFunctionRequest_SdkV2) {
}

func (to *ExternalFunctionRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalFunctionRequest_SdkV2) {
}

func (m ExternalFunctionRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["connection_name"] = attrs["connection_name"].SetRequired()
	attrs["headers"] = attrs["headers"].SetOptional()
	attrs["json"] = attrs["json"].SetOptional()
	attrs["method"] = attrs["method"].SetRequired()
	attrs["params"] = attrs["params"].SetOptional()
	attrs["path"] = attrs["path"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalFunctionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalFunctionRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalFunctionRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalFunctionRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_name": m.ConnectionName,
			"headers":         m.Headers,
			"json":            m.Json,
			"method":          m.Method,
			"params":          m.Params,
			"path":            m.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalFunctionRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"connection_name": types.StringType,
			"headers":         types.StringType,
			"json":            types.StringType,
			"method":          types.StringType,
			"params":          types.StringType,
			"path":            types.StringType,
		},
	}
}

type ExternalModel_SdkV2 struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig types.List `tfsdk:"ai21labs_config"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig types.List `tfsdk:"amazon_bedrock_config"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig types.List `tfsdk:"anthropic_config"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig types.List `tfsdk:"cohere_config"`
	// Custom Provider Config. Only required if the provider is 'custom'.
	CustomProviderConfig types.List `tfsdk:"custom_provider_config"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig types.List `tfsdk:"databricks_model_serving_config"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig types.List `tfsdk:"google_cloud_vertex_ai_config"`
	// The name of the external model.
	Name types.String `tfsdk:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig types.List `tfsdk:"openai_config"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig types.List `tfsdk:"palm_config"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	Provider types.String `tfsdk:"provider"`
	// The task type of the external model.
	Task types.String `tfsdk:"task"`
}

func (to *ExternalModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalModel_SdkV2) {
	if !from.Ai21labsConfig.IsNull() && !from.Ai21labsConfig.IsUnknown() {
		if toAi21labsConfig, ok := to.GetAi21labsConfig(ctx); ok {
			if fromAi21labsConfig, ok := from.GetAi21labsConfig(ctx); ok {
				// Recursively sync the fields of Ai21labsConfig
				toAi21labsConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAi21labsConfig)
				to.SetAi21labsConfig(ctx, toAi21labsConfig)
			}
		}
	}
	if !from.AmazonBedrockConfig.IsNull() && !from.AmazonBedrockConfig.IsUnknown() {
		if toAmazonBedrockConfig, ok := to.GetAmazonBedrockConfig(ctx); ok {
			if fromAmazonBedrockConfig, ok := from.GetAmazonBedrockConfig(ctx); ok {
				// Recursively sync the fields of AmazonBedrockConfig
				toAmazonBedrockConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAmazonBedrockConfig)
				to.SetAmazonBedrockConfig(ctx, toAmazonBedrockConfig)
			}
		}
	}
	if !from.AnthropicConfig.IsNull() && !from.AnthropicConfig.IsUnknown() {
		if toAnthropicConfig, ok := to.GetAnthropicConfig(ctx); ok {
			if fromAnthropicConfig, ok := from.GetAnthropicConfig(ctx); ok {
				// Recursively sync the fields of AnthropicConfig
				toAnthropicConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromAnthropicConfig)
				to.SetAnthropicConfig(ctx, toAnthropicConfig)
			}
		}
	}
	if !from.CohereConfig.IsNull() && !from.CohereConfig.IsUnknown() {
		if toCohereConfig, ok := to.GetCohereConfig(ctx); ok {
			if fromCohereConfig, ok := from.GetCohereConfig(ctx); ok {
				// Recursively sync the fields of CohereConfig
				toCohereConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromCohereConfig)
				to.SetCohereConfig(ctx, toCohereConfig)
			}
		}
	}
	if !from.CustomProviderConfig.IsNull() && !from.CustomProviderConfig.IsUnknown() {
		if toCustomProviderConfig, ok := to.GetCustomProviderConfig(ctx); ok {
			if fromCustomProviderConfig, ok := from.GetCustomProviderConfig(ctx); ok {
				// Recursively sync the fields of CustomProviderConfig
				toCustomProviderConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromCustomProviderConfig)
				to.SetCustomProviderConfig(ctx, toCustomProviderConfig)
			}
		}
	}
	if !from.DatabricksModelServingConfig.IsNull() && !from.DatabricksModelServingConfig.IsUnknown() {
		if toDatabricksModelServingConfig, ok := to.GetDatabricksModelServingConfig(ctx); ok {
			if fromDatabricksModelServingConfig, ok := from.GetDatabricksModelServingConfig(ctx); ok {
				// Recursively sync the fields of DatabricksModelServingConfig
				toDatabricksModelServingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromDatabricksModelServingConfig)
				to.SetDatabricksModelServingConfig(ctx, toDatabricksModelServingConfig)
			}
		}
	}
	if !from.GoogleCloudVertexAiConfig.IsNull() && !from.GoogleCloudVertexAiConfig.IsUnknown() {
		if toGoogleCloudVertexAiConfig, ok := to.GetGoogleCloudVertexAiConfig(ctx); ok {
			if fromGoogleCloudVertexAiConfig, ok := from.GetGoogleCloudVertexAiConfig(ctx); ok {
				// Recursively sync the fields of GoogleCloudVertexAiConfig
				toGoogleCloudVertexAiConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromGoogleCloudVertexAiConfig)
				to.SetGoogleCloudVertexAiConfig(ctx, toGoogleCloudVertexAiConfig)
			}
		}
	}
	if !from.OpenaiConfig.IsNull() && !from.OpenaiConfig.IsUnknown() {
		if toOpenaiConfig, ok := to.GetOpenaiConfig(ctx); ok {
			if fromOpenaiConfig, ok := from.GetOpenaiConfig(ctx); ok {
				// Recursively sync the fields of OpenaiConfig
				toOpenaiConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromOpenaiConfig)
				to.SetOpenaiConfig(ctx, toOpenaiConfig)
			}
		}
	}
	if !from.PalmConfig.IsNull() && !from.PalmConfig.IsUnknown() {
		if toPalmConfig, ok := to.GetPalmConfig(ctx); ok {
			if fromPalmConfig, ok := from.GetPalmConfig(ctx); ok {
				// Recursively sync the fields of PalmConfig
				toPalmConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPalmConfig)
				to.SetPalmConfig(ctx, toPalmConfig)
			}
		}
	}
}

func (to *ExternalModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalModel_SdkV2) {
	if !from.Ai21labsConfig.IsNull() && !from.Ai21labsConfig.IsUnknown() {
		if toAi21labsConfig, ok := to.GetAi21labsConfig(ctx); ok {
			if fromAi21labsConfig, ok := from.GetAi21labsConfig(ctx); ok {
				toAi21labsConfig.SyncFieldsDuringRead(ctx, fromAi21labsConfig)
				to.SetAi21labsConfig(ctx, toAi21labsConfig)
			}
		}
	}
	if !from.AmazonBedrockConfig.IsNull() && !from.AmazonBedrockConfig.IsUnknown() {
		if toAmazonBedrockConfig, ok := to.GetAmazonBedrockConfig(ctx); ok {
			if fromAmazonBedrockConfig, ok := from.GetAmazonBedrockConfig(ctx); ok {
				toAmazonBedrockConfig.SyncFieldsDuringRead(ctx, fromAmazonBedrockConfig)
				to.SetAmazonBedrockConfig(ctx, toAmazonBedrockConfig)
			}
		}
	}
	if !from.AnthropicConfig.IsNull() && !from.AnthropicConfig.IsUnknown() {
		if toAnthropicConfig, ok := to.GetAnthropicConfig(ctx); ok {
			if fromAnthropicConfig, ok := from.GetAnthropicConfig(ctx); ok {
				toAnthropicConfig.SyncFieldsDuringRead(ctx, fromAnthropicConfig)
				to.SetAnthropicConfig(ctx, toAnthropicConfig)
			}
		}
	}
	if !from.CohereConfig.IsNull() && !from.CohereConfig.IsUnknown() {
		if toCohereConfig, ok := to.GetCohereConfig(ctx); ok {
			if fromCohereConfig, ok := from.GetCohereConfig(ctx); ok {
				toCohereConfig.SyncFieldsDuringRead(ctx, fromCohereConfig)
				to.SetCohereConfig(ctx, toCohereConfig)
			}
		}
	}
	if !from.CustomProviderConfig.IsNull() && !from.CustomProviderConfig.IsUnknown() {
		if toCustomProviderConfig, ok := to.GetCustomProviderConfig(ctx); ok {
			if fromCustomProviderConfig, ok := from.GetCustomProviderConfig(ctx); ok {
				toCustomProviderConfig.SyncFieldsDuringRead(ctx, fromCustomProviderConfig)
				to.SetCustomProviderConfig(ctx, toCustomProviderConfig)
			}
		}
	}
	if !from.DatabricksModelServingConfig.IsNull() && !from.DatabricksModelServingConfig.IsUnknown() {
		if toDatabricksModelServingConfig, ok := to.GetDatabricksModelServingConfig(ctx); ok {
			if fromDatabricksModelServingConfig, ok := from.GetDatabricksModelServingConfig(ctx); ok {
				toDatabricksModelServingConfig.SyncFieldsDuringRead(ctx, fromDatabricksModelServingConfig)
				to.SetDatabricksModelServingConfig(ctx, toDatabricksModelServingConfig)
			}
		}
	}
	if !from.GoogleCloudVertexAiConfig.IsNull() && !from.GoogleCloudVertexAiConfig.IsUnknown() {
		if toGoogleCloudVertexAiConfig, ok := to.GetGoogleCloudVertexAiConfig(ctx); ok {
			if fromGoogleCloudVertexAiConfig, ok := from.GetGoogleCloudVertexAiConfig(ctx); ok {
				toGoogleCloudVertexAiConfig.SyncFieldsDuringRead(ctx, fromGoogleCloudVertexAiConfig)
				to.SetGoogleCloudVertexAiConfig(ctx, toGoogleCloudVertexAiConfig)
			}
		}
	}
	if !from.OpenaiConfig.IsNull() && !from.OpenaiConfig.IsUnknown() {
		if toOpenaiConfig, ok := to.GetOpenaiConfig(ctx); ok {
			if fromOpenaiConfig, ok := from.GetOpenaiConfig(ctx); ok {
				toOpenaiConfig.SyncFieldsDuringRead(ctx, fromOpenaiConfig)
				to.SetOpenaiConfig(ctx, toOpenaiConfig)
			}
		}
	}
	if !from.PalmConfig.IsNull() && !from.PalmConfig.IsUnknown() {
		if toPalmConfig, ok := to.GetPalmConfig(ctx); ok {
			if fromPalmConfig, ok := from.GetPalmConfig(ctx); ok {
				toPalmConfig.SyncFieldsDuringRead(ctx, fromPalmConfig)
				to.SetPalmConfig(ctx, toPalmConfig)
			}
		}
	}
}

func (m ExternalModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai21labs_config"] = attrs["ai21labs_config"].SetOptional()
	attrs["ai21labs_config"] = attrs["ai21labs_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["amazon_bedrock_config"] = attrs["amazon_bedrock_config"].SetOptional()
	attrs["amazon_bedrock_config"] = attrs["amazon_bedrock_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["anthropic_config"] = attrs["anthropic_config"].SetOptional()
	attrs["anthropic_config"] = attrs["anthropic_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["cohere_config"] = attrs["cohere_config"].SetOptional()
	attrs["cohere_config"] = attrs["cohere_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["custom_provider_config"] = attrs["custom_provider_config"].SetOptional()
	attrs["custom_provider_config"] = attrs["custom_provider_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["databricks_model_serving_config"] = attrs["databricks_model_serving_config"].SetOptional()
	attrs["databricks_model_serving_config"] = attrs["databricks_model_serving_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["google_cloud_vertex_ai_config"] = attrs["google_cloud_vertex_ai_config"].SetOptional()
	attrs["google_cloud_vertex_ai_config"] = attrs["google_cloud_vertex_ai_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()
	attrs["openai_config"] = attrs["openai_config"].SetOptional()
	attrs["openai_config"] = attrs["openai_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["palm_config"] = attrs["palm_config"].SetOptional()
	attrs["palm_config"] = attrs["palm_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["provider"] = attrs["provider"].SetRequired()
	attrs["task"] = attrs["task"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai21labs_config":                 reflect.TypeOf(Ai21LabsConfig_SdkV2{}),
		"amazon_bedrock_config":           reflect.TypeOf(AmazonBedrockConfig_SdkV2{}),
		"anthropic_config":                reflect.TypeOf(AnthropicConfig_SdkV2{}),
		"cohere_config":                   reflect.TypeOf(CohereConfig_SdkV2{}),
		"custom_provider_config":          reflect.TypeOf(CustomProviderConfig_SdkV2{}),
		"databricks_model_serving_config": reflect.TypeOf(DatabricksModelServingConfig_SdkV2{}),
		"google_cloud_vertex_ai_config":   reflect.TypeOf(GoogleCloudVertexAiConfig_SdkV2{}),
		"openai_config":                   reflect.TypeOf(OpenAiConfig_SdkV2{}),
		"palm_config":                     reflect.TypeOf(PaLmConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalModel_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai21labs_config":                 m.Ai21labsConfig,
			"amazon_bedrock_config":           m.AmazonBedrockConfig,
			"anthropic_config":                m.AnthropicConfig,
			"cohere_config":                   m.CohereConfig,
			"custom_provider_config":          m.CustomProviderConfig,
			"databricks_model_serving_config": m.DatabricksModelServingConfig,
			"google_cloud_vertex_ai_config":   m.GoogleCloudVertexAiConfig,
			"name":                            m.Name,
			"openai_config":                   m.OpenaiConfig,
			"palm_config":                     m.PalmConfig,
			"provider":                        m.Provider,
			"task":                            m.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_config": basetypes.ListType{
				ElemType: Ai21LabsConfig_SdkV2{}.Type(ctx),
			},
			"amazon_bedrock_config": basetypes.ListType{
				ElemType: AmazonBedrockConfig_SdkV2{}.Type(ctx),
			},
			"anthropic_config": basetypes.ListType{
				ElemType: AnthropicConfig_SdkV2{}.Type(ctx),
			},
			"cohere_config": basetypes.ListType{
				ElemType: CohereConfig_SdkV2{}.Type(ctx),
			},
			"custom_provider_config": basetypes.ListType{
				ElemType: CustomProviderConfig_SdkV2{}.Type(ctx),
			},
			"databricks_model_serving_config": basetypes.ListType{
				ElemType: DatabricksModelServingConfig_SdkV2{}.Type(ctx),
			},
			"google_cloud_vertex_ai_config": basetypes.ListType{
				ElemType: GoogleCloudVertexAiConfig_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"openai_config": basetypes.ListType{
				ElemType: OpenAiConfig_SdkV2{}.Type(ctx),
			},
			"palm_config": basetypes.ListType{
				ElemType: PaLmConfig_SdkV2{}.Type(ctx),
			},
			"provider": types.StringType,
			"task":     types.StringType,
		},
	}
}

// GetAi21labsConfig returns the value of the Ai21labsConfig field in ExternalModel_SdkV2 as
// a Ai21LabsConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetAi21labsConfig(ctx context.Context) (Ai21LabsConfig_SdkV2, bool) {
	var e Ai21LabsConfig_SdkV2
	if m.Ai21labsConfig.IsNull() || m.Ai21labsConfig.IsUnknown() {
		return e, false
	}
	var v []Ai21LabsConfig_SdkV2
	d := m.Ai21labsConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAi21labsConfig sets the value of the Ai21labsConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetAi21labsConfig(ctx context.Context, v Ai21LabsConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ai21labs_config"]
	m.Ai21labsConfig = types.ListValueMust(t, vs)
}

// GetAmazonBedrockConfig returns the value of the AmazonBedrockConfig field in ExternalModel_SdkV2 as
// a AmazonBedrockConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetAmazonBedrockConfig(ctx context.Context) (AmazonBedrockConfig_SdkV2, bool) {
	var e AmazonBedrockConfig_SdkV2
	if m.AmazonBedrockConfig.IsNull() || m.AmazonBedrockConfig.IsUnknown() {
		return e, false
	}
	var v []AmazonBedrockConfig_SdkV2
	d := m.AmazonBedrockConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAmazonBedrockConfig sets the value of the AmazonBedrockConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetAmazonBedrockConfig(ctx context.Context, v AmazonBedrockConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["amazon_bedrock_config"]
	m.AmazonBedrockConfig = types.ListValueMust(t, vs)
}

// GetAnthropicConfig returns the value of the AnthropicConfig field in ExternalModel_SdkV2 as
// a AnthropicConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetAnthropicConfig(ctx context.Context) (AnthropicConfig_SdkV2, bool) {
	var e AnthropicConfig_SdkV2
	if m.AnthropicConfig.IsNull() || m.AnthropicConfig.IsUnknown() {
		return e, false
	}
	var v []AnthropicConfig_SdkV2
	d := m.AnthropicConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAnthropicConfig sets the value of the AnthropicConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetAnthropicConfig(ctx context.Context, v AnthropicConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["anthropic_config"]
	m.AnthropicConfig = types.ListValueMust(t, vs)
}

// GetCohereConfig returns the value of the CohereConfig field in ExternalModel_SdkV2 as
// a CohereConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetCohereConfig(ctx context.Context) (CohereConfig_SdkV2, bool) {
	var e CohereConfig_SdkV2
	if m.CohereConfig.IsNull() || m.CohereConfig.IsUnknown() {
		return e, false
	}
	var v []CohereConfig_SdkV2
	d := m.CohereConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCohereConfig sets the value of the CohereConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetCohereConfig(ctx context.Context, v CohereConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["cohere_config"]
	m.CohereConfig = types.ListValueMust(t, vs)
}

// GetCustomProviderConfig returns the value of the CustomProviderConfig field in ExternalModel_SdkV2 as
// a CustomProviderConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetCustomProviderConfig(ctx context.Context) (CustomProviderConfig_SdkV2, bool) {
	var e CustomProviderConfig_SdkV2
	if m.CustomProviderConfig.IsNull() || m.CustomProviderConfig.IsUnknown() {
		return e, false
	}
	var v []CustomProviderConfig_SdkV2
	d := m.CustomProviderConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetCustomProviderConfig sets the value of the CustomProviderConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetCustomProviderConfig(ctx context.Context, v CustomProviderConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["custom_provider_config"]
	m.CustomProviderConfig = types.ListValueMust(t, vs)
}

// GetDatabricksModelServingConfig returns the value of the DatabricksModelServingConfig field in ExternalModel_SdkV2 as
// a DatabricksModelServingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetDatabricksModelServingConfig(ctx context.Context) (DatabricksModelServingConfig_SdkV2, bool) {
	var e DatabricksModelServingConfig_SdkV2
	if m.DatabricksModelServingConfig.IsNull() || m.DatabricksModelServingConfig.IsUnknown() {
		return e, false
	}
	var v []DatabricksModelServingConfig_SdkV2
	d := m.DatabricksModelServingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDatabricksModelServingConfig sets the value of the DatabricksModelServingConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetDatabricksModelServingConfig(ctx context.Context, v DatabricksModelServingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["databricks_model_serving_config"]
	m.DatabricksModelServingConfig = types.ListValueMust(t, vs)
}

// GetGoogleCloudVertexAiConfig returns the value of the GoogleCloudVertexAiConfig field in ExternalModel_SdkV2 as
// a GoogleCloudVertexAiConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetGoogleCloudVertexAiConfig(ctx context.Context) (GoogleCloudVertexAiConfig_SdkV2, bool) {
	var e GoogleCloudVertexAiConfig_SdkV2
	if m.GoogleCloudVertexAiConfig.IsNull() || m.GoogleCloudVertexAiConfig.IsUnknown() {
		return e, false
	}
	var v []GoogleCloudVertexAiConfig_SdkV2
	d := m.GoogleCloudVertexAiConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGoogleCloudVertexAiConfig sets the value of the GoogleCloudVertexAiConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetGoogleCloudVertexAiConfig(ctx context.Context, v GoogleCloudVertexAiConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["google_cloud_vertex_ai_config"]
	m.GoogleCloudVertexAiConfig = types.ListValueMust(t, vs)
}

// GetOpenaiConfig returns the value of the OpenaiConfig field in ExternalModel_SdkV2 as
// a OpenAiConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetOpenaiConfig(ctx context.Context) (OpenAiConfig_SdkV2, bool) {
	var e OpenAiConfig_SdkV2
	if m.OpenaiConfig.IsNull() || m.OpenaiConfig.IsUnknown() {
		return e, false
	}
	var v []OpenAiConfig_SdkV2
	d := m.OpenaiConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetOpenaiConfig sets the value of the OpenaiConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetOpenaiConfig(ctx context.Context, v OpenAiConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["openai_config"]
	m.OpenaiConfig = types.ListValueMust(t, vs)
}

// GetPalmConfig returns the value of the PalmConfig field in ExternalModel_SdkV2 as
// a PaLmConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ExternalModel_SdkV2) GetPalmConfig(ctx context.Context) (PaLmConfig_SdkV2, bool) {
	var e PaLmConfig_SdkV2
	if m.PalmConfig.IsNull() || m.PalmConfig.IsUnknown() {
		return e, false
	}
	var v []PaLmConfig_SdkV2
	d := m.PalmConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPalmConfig sets the value of the PalmConfig field in ExternalModel_SdkV2.
func (m *ExternalModel_SdkV2) SetPalmConfig(ctx context.Context, v PaLmConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["palm_config"]
	m.PalmConfig = types.ListValueMust(t, vs)
}

type ExternalModelUsageElement_SdkV2 struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens types.Int64 `tfsdk:"completion_tokens"`
	// The number of tokens in the prompt.
	PromptTokens types.Int64 `tfsdk:"prompt_tokens"`
	// The total number of tokens in the prompt and response.
	TotalTokens types.Int64 `tfsdk:"total_tokens"`
}

func (to *ExternalModelUsageElement_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalModelUsageElement_SdkV2) {
}

func (to *ExternalModelUsageElement_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ExternalModelUsageElement_SdkV2) {
}

func (m ExternalModelUsageElement_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["completion_tokens"] = attrs["completion_tokens"].SetOptional()
	attrs["prompt_tokens"] = attrs["prompt_tokens"].SetOptional()
	attrs["total_tokens"] = attrs["total_tokens"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ExternalModelUsageElement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ExternalModelUsageElement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalModelUsageElement_SdkV2
// only implements ToObjectValue() and Type().
func (m ExternalModelUsageElement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"completion_tokens": m.CompletionTokens,
			"prompt_tokens":     m.PromptTokens,
			"total_tokens":      m.TotalTokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ExternalModelUsageElement_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"completion_tokens": types.Int64Type,
			"prompt_tokens":     types.Int64Type,
			"total_tokens":      types.Int64Type,
		},
	}
}

type FallbackConfig_SdkV2 struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (to *FallbackConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FallbackConfig_SdkV2) {
}

func (to *FallbackConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FallbackConfig_SdkV2) {
}

func (m FallbackConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["enabled"] = attrs["enabled"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FallbackConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FallbackConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FallbackConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m FallbackConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": m.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FallbackConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

// All fields are not sensitive as they are hard-coded in the system and made
// available to customers.
type FoundationModel_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	DisplayName types.String `tfsdk:"display_name"`

	Docs types.String `tfsdk:"docs"`

	Name types.String `tfsdk:"name"`
}

func (to *FoundationModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FoundationModel_SdkV2) {
}

func (to *FoundationModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from FoundationModel_SdkV2) {
}

func (m FoundationModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["docs"] = attrs["docs"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in FoundationModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m FoundationModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FoundationModel_SdkV2
// only implements ToObjectValue() and Type().
func (m FoundationModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":  m.Description,
			"display_name": m.DisplayName,
			"docs":         m.Docs,
			"name":         m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m FoundationModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":  types.StringType,
			"display_name": types.StringType,
			"docs":         types.StringType,
			"name":         types.StringType,
		},
	}
}

type GetOpenApiRequest_SdkV2 struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
}

func (to *GetOpenApiRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOpenApiRequest_SdkV2) {
}

func (to *GetOpenApiRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetOpenApiRequest_SdkV2) {
}

func (m GetOpenApiRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOpenApiRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOpenApiRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOpenApiRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetOpenApiRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOpenApiRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOpenApiResponse_SdkV2 struct {
	Contents types.Object `tfsdk:"-"`
}

func (to *GetOpenApiResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOpenApiResponse_SdkV2) {
}

func (to *GetOpenApiResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetOpenApiResponse_SdkV2) {
}

func (m GetOpenApiResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetOpenApiResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetOpenApiResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOpenApiResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetOpenApiResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": m.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetOpenApiResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

type GetServingEndpointPermissionLevelsRequest_SdkV2 struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (to *GetServingEndpointPermissionLevelsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointPermissionLevelsRequest_SdkV2) {
}

func (to *GetServingEndpointPermissionLevelsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointPermissionLevelsRequest_SdkV2) {
}

func (m GetServingEndpointPermissionLevelsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["serving_endpoint_id"] = attrs["serving_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointPermissionLevelsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServingEndpointPermissionLevelsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionLevelsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServingEndpointPermissionLevelsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"serving_endpoint_id": m.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServingEndpointPermissionLevelsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

type GetServingEndpointPermissionLevelsResponse_SdkV2 struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetServingEndpointPermissionLevelsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetServingEndpointPermissionLevelsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointPermissionLevelsResponse_SdkV2) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (m GetServingEndpointPermissionLevelsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["permission_levels"] = attrs["permission_levels"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointPermissionLevelsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServingEndpointPermissionLevelsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ServingEndpointPermissionsDescription_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionLevelsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServingEndpointPermissionLevelsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": m.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServingEndpointPermissionLevelsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ServingEndpointPermissionsDescription_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetServingEndpointPermissionLevelsResponse_SdkV2 as
// a slice of ServingEndpointPermissionsDescription_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *GetServingEndpointPermissionLevelsResponse_SdkV2) GetPermissionLevels(ctx context.Context) ([]ServingEndpointPermissionsDescription_SdkV2, bool) {
	if m.PermissionLevels.IsNull() || m.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointPermissionsDescription_SdkV2
	d := m.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetServingEndpointPermissionLevelsResponse_SdkV2.
func (m *GetServingEndpointPermissionLevelsResponse_SdkV2) SetPermissionLevels(ctx context.Context, v []ServingEndpointPermissionsDescription_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.PermissionLevels = types.ListValueMust(t, vs)
}

type GetServingEndpointPermissionsRequest_SdkV2 struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (to *GetServingEndpointPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointPermissionsRequest_SdkV2) {
}

func (to *GetServingEndpointPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointPermissionsRequest_SdkV2) {
}

func (m GetServingEndpointPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["serving_endpoint_id"] = attrs["serving_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServingEndpointPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServingEndpointPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"serving_endpoint_id": m.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServingEndpointPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

type GetServingEndpointRequest_SdkV2 struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (to *GetServingEndpointRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointRequest_SdkV2) {
}

func (to *GetServingEndpointRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointRequest_SdkV2) {
}

func (m GetServingEndpointRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetServingEndpointRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GetServingEndpointRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m GetServingEndpointRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GetServingEndpointRequest_SdkV2) Type(ctx context.Context) attr.Type {
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
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	PrivateKey types.String `tfsdk:"private_key"`
	// The private key for the service account which has access to the Google
	// Cloud Vertex AI Service provided as a plaintext secret. See [Best
	// practices for managing service account keys]. If you prefer to reference
	// your key using Databricks Secrets, see `private_key`. You must provide an
	// API key using one of the following fields: `private_key` or
	// `private_key_plaintext`.
	//
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	PrivateKeyPlaintext types.String `tfsdk:"private_key_plaintext"`
	// This is the Google Cloud project id that the service account is
	// associated with.
	ProjectId types.String `tfsdk:"project_id"`
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]:
	// https://cloud.google.com/vertex-ai/docs/general/locations
	Region types.String `tfsdk:"region"`
}

func (to *GoogleCloudVertexAiConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GoogleCloudVertexAiConfig_SdkV2) {
}

func (to *GoogleCloudVertexAiConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from GoogleCloudVertexAiConfig_SdkV2) {
}

func (m GoogleCloudVertexAiConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["private_key"] = attrs["private_key"].SetOptional()
	attrs["private_key_plaintext"] = attrs["private_key_plaintext"].SetOptional()
	attrs["project_id"] = attrs["project_id"].SetRequired()
	attrs["region"] = attrs["region"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GoogleCloudVertexAiConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m GoogleCloudVertexAiConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GoogleCloudVertexAiConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m GoogleCloudVertexAiConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"private_key":           m.PrivateKey,
			"private_key_plaintext": m.PrivateKeyPlaintext,
			"project_id":            m.ProjectId,
			"region":                m.Region,
		})
}

// Type implements basetypes.ObjectValuable.
func (m GoogleCloudVertexAiConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_key":           types.StringType,
			"private_key_plaintext": types.StringType,
			"project_id":            types.StringType,
			"region":                types.StringType,
		},
	}
}

type HttpRequestResponse_SdkV2 struct {
	Contents types.Object `tfsdk:"-"`
}

func (to *HttpRequestResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HttpRequestResponse_SdkV2) {
}

func (to *HttpRequestResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from HttpRequestResponse_SdkV2) {
}

func (m HttpRequestResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["contents"] = attrs["contents"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in HttpRequestResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m HttpRequestResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpRequestResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m HttpRequestResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": m.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (m HttpRequestResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

type ListEndpointsResponse_SdkV2 struct {
	// The list of endpoints.
	Endpoints types.List `tfsdk:"endpoints"`
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse_SdkV2) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (m ListEndpointsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["endpoints"] = attrs["endpoints"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListEndpointsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListEndpointsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(ServingEndpoint_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ListEndpointsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints": m.Endpoints,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ListEndpointsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: ServingEndpoint_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse_SdkV2 as
// a slice of ServingEndpoint_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ListEndpointsResponse_SdkV2) GetEndpoints(ctx context.Context) ([]ServingEndpoint_SdkV2, bool) {
	if m.Endpoints.IsNull() || m.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpoint_SdkV2
	d := m.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse_SdkV2.
func (m *ListEndpointsResponse_SdkV2) SetEndpoints(ctx context.Context, v []ServingEndpoint_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Endpoints = types.ListValueMust(t, vs)
}

type ListServingEndpointsRequest_SdkV2 struct {
}

func (to *ListServingEndpointsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServingEndpointsRequest_SdkV2) {
}

func (to *ListServingEndpointsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ListServingEndpointsRequest_SdkV2) {
}

func (m ListServingEndpointsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServingEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ListServingEndpointsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServingEndpointsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ListServingEndpointsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (m ListServingEndpointsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogsRequest_SdkV2 struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (to *LogsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogsRequest_SdkV2) {
}

func (to *LogsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from LogsRequest_SdkV2) {
}

func (m LogsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetRequired()
	attrs["served_model_name"] = attrs["served_model_name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in LogsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m LogsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m LogsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              m.Name,
			"served_model_name": m.ServedModelName,
		})
}

// Type implements basetypes.ObjectValuable.
func (m LogsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

// A representation of all DataPlaneInfo for operations that can be done on a
// model through Data Plane APIs.
type ModelDataPlaneInfo_SdkV2 struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo types.List `tfsdk:"query_info"`
}

func (to *ModelDataPlaneInfo_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelDataPlaneInfo_SdkV2) {
	if !from.QueryInfo.IsNull() && !from.QueryInfo.IsUnknown() {
		if toQueryInfo, ok := to.GetQueryInfo(ctx); ok {
			if fromQueryInfo, ok := from.GetQueryInfo(ctx); ok {
				// Recursively sync the fields of QueryInfo
				toQueryInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromQueryInfo)
				to.SetQueryInfo(ctx, toQueryInfo)
			}
		}
	}
}

func (to *ModelDataPlaneInfo_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ModelDataPlaneInfo_SdkV2) {
	if !from.QueryInfo.IsNull() && !from.QueryInfo.IsUnknown() {
		if toQueryInfo, ok := to.GetQueryInfo(ctx); ok {
			if fromQueryInfo, ok := from.GetQueryInfo(ctx); ok {
				toQueryInfo.SyncFieldsDuringRead(ctx, fromQueryInfo)
				to.SetQueryInfo(ctx, toQueryInfo)
			}
		}
	}
}

func (m ModelDataPlaneInfo_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_info"] = attrs["query_info"].SetOptional()
	attrs["query_info"] = attrs["query_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelDataPlaneInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ModelDataPlaneInfo_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_info": reflect.TypeOf(DataPlaneInfo_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDataPlaneInfo_SdkV2
// only implements ToObjectValue() and Type().
func (m ModelDataPlaneInfo_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_info": m.QueryInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ModelDataPlaneInfo_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_info": basetypes.ListType{
				ElemType: DataPlaneInfo_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetQueryInfo returns the value of the QueryInfo field in ModelDataPlaneInfo_SdkV2 as
// a DataPlaneInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ModelDataPlaneInfo_SdkV2) GetQueryInfo(ctx context.Context) (DataPlaneInfo_SdkV2, bool) {
	var e DataPlaneInfo_SdkV2
	if m.QueryInfo.IsNull() || m.QueryInfo.IsUnknown() {
		return e, false
	}
	var v []DataPlaneInfo_SdkV2
	d := m.QueryInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetQueryInfo sets the value of the QueryInfo field in ModelDataPlaneInfo_SdkV2.
func (m *ModelDataPlaneInfo_SdkV2) SetQueryInfo(ctx context.Context, v DataPlaneInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["query_info"]
	m.QueryInfo = types.ListValueMust(t, vs)
}

// Configs needed to create an OpenAI model route.
type OpenAiConfig_SdkV2 struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	MicrosoftEntraClientId types.String `tfsdk:"microsoft_entra_client_id"`
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecret types.String `tfsdk:"microsoft_entra_client_secret"`
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecretPlaintext types.String `tfsdk:"microsoft_entra_client_secret_plaintext"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId types.String `tfsdk:"microsoft_entra_tenant_id"`
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	OpenaiApiBase types.String `tfsdk:"openai_api_base"`
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKey types.String `tfsdk:"openai_api_key"`
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKeyPlaintext types.String `tfsdk:"openai_api_key_plaintext"`
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	OpenaiApiType types.String `tfsdk:"openai_api_type"`
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	OpenaiApiVersion types.String `tfsdk:"openai_api_version"`
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	OpenaiDeploymentName types.String `tfsdk:"openai_deployment_name"`
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	OpenaiOrganization types.String `tfsdk:"openai_organization"`
}

func (to *OpenAiConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OpenAiConfig_SdkV2) {
}

func (to *OpenAiConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from OpenAiConfig_SdkV2) {
}

func (m OpenAiConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["microsoft_entra_client_id"] = attrs["microsoft_entra_client_id"].SetOptional()
	attrs["microsoft_entra_client_secret"] = attrs["microsoft_entra_client_secret"].SetOptional()
	attrs["microsoft_entra_client_secret_plaintext"] = attrs["microsoft_entra_client_secret_plaintext"].SetOptional()
	attrs["microsoft_entra_tenant_id"] = attrs["microsoft_entra_tenant_id"].SetOptional()
	attrs["openai_api_base"] = attrs["openai_api_base"].SetOptional()
	attrs["openai_api_key"] = attrs["openai_api_key"].SetOptional()
	attrs["openai_api_key_plaintext"] = attrs["openai_api_key_plaintext"].SetOptional()
	attrs["openai_api_type"] = attrs["openai_api_type"].SetOptional()
	attrs["openai_api_version"] = attrs["openai_api_version"].SetOptional()
	attrs["openai_deployment_name"] = attrs["openai_deployment_name"].SetOptional()
	attrs["openai_organization"] = attrs["openai_organization"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in OpenAiConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m OpenAiConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OpenAiConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m OpenAiConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"microsoft_entra_client_id":               m.MicrosoftEntraClientId,
			"microsoft_entra_client_secret":           m.MicrosoftEntraClientSecret,
			"microsoft_entra_client_secret_plaintext": m.MicrosoftEntraClientSecretPlaintext,
			"microsoft_entra_tenant_id":               m.MicrosoftEntraTenantId,
			"openai_api_base":                         m.OpenaiApiBase,
			"openai_api_key":                          m.OpenaiApiKey,
			"openai_api_key_plaintext":                m.OpenaiApiKeyPlaintext,
			"openai_api_type":                         m.OpenaiApiType,
			"openai_api_version":                      m.OpenaiApiVersion,
			"openai_deployment_name":                  m.OpenaiDeploymentName,
			"openai_organization":                     m.OpenaiOrganization,
		})
}

// Type implements basetypes.ObjectValuable.
func (m OpenAiConfig_SdkV2) Type(ctx context.Context) attr.Type {
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
	PalmApiKey types.String `tfsdk:"palm_api_key"`
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKeyPlaintext types.String `tfsdk:"palm_api_key_plaintext"`
}

func (to *PaLmConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PaLmConfig_SdkV2) {
}

func (to *PaLmConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PaLmConfig_SdkV2) {
}

func (m PaLmConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["palm_api_key"] = attrs["palm_api_key"].SetOptional()
	attrs["palm_api_key_plaintext"] = attrs["palm_api_key_plaintext"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PaLmConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PaLmConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PaLmConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m PaLmConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"palm_api_key":           m.PalmApiKey,
			"palm_api_key_plaintext": m.PalmApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PaLmConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"palm_api_key":           types.StringType,
			"palm_api_key_plaintext": types.StringType,
		},
	}
}

type PatchServingEndpointTags_SdkV2 struct {
	// List of endpoint tags to add
	AddTags types.List `tfsdk:"add_tags"`
	// List of tag keys to delete
	DeleteTags types.List `tfsdk:"delete_tags"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (to *PatchServingEndpointTags_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchServingEndpointTags_SdkV2) {
	if !from.AddTags.IsNull() && !from.AddTags.IsUnknown() && to.AddTags.IsNull() && len(from.AddTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AddTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AddTags = from.AddTags
	}
	if !from.DeleteTags.IsNull() && !from.DeleteTags.IsUnknown() && to.DeleteTags.IsNull() && len(from.DeleteTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DeleteTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DeleteTags = from.DeleteTags
	}
}

func (to *PatchServingEndpointTags_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PatchServingEndpointTags_SdkV2) {
	if !from.AddTags.IsNull() && !from.AddTags.IsUnknown() && to.AddTags.IsNull() && len(from.AddTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AddTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AddTags = from.AddTags
	}
	if !from.DeleteTags.IsNull() && !from.DeleteTags.IsUnknown() && to.DeleteTags.IsNull() && len(from.DeleteTags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DeleteTags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DeleteTags = from.DeleteTags
	}
}

func (m PatchServingEndpointTags_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["add_tags"] = attrs["add_tags"].SetOptional()
	attrs["delete_tags"] = attrs["delete_tags"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PatchServingEndpointTags.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PatchServingEndpointTags_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add_tags":    reflect.TypeOf(EndpointTag_SdkV2{}),
		"delete_tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchServingEndpointTags_SdkV2
// only implements ToObjectValue() and Type().
func (m PatchServingEndpointTags_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add_tags":    m.AddTags,
			"delete_tags": m.DeleteTags,
			"name":        m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PatchServingEndpointTags_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add_tags": basetypes.ListType{
				ElemType: EndpointTag_SdkV2{}.Type(ctx),
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
func (m *PatchServingEndpointTags_SdkV2) GetAddTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if m.AddTags.IsNull() || m.AddTags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := m.AddTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAddTags sets the value of the AddTags field in PatchServingEndpointTags_SdkV2.
func (m *PatchServingEndpointTags_SdkV2) SetAddTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["add_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AddTags = types.ListValueMust(t, vs)
}

// GetDeleteTags returns the value of the DeleteTags field in PatchServingEndpointTags_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *PatchServingEndpointTags_SdkV2) GetDeleteTags(ctx context.Context) ([]types.String, bool) {
	if m.DeleteTags.IsNull() || m.DeleteTags.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.DeleteTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDeleteTags sets the value of the DeleteTags field in PatchServingEndpointTags_SdkV2.
func (m *PatchServingEndpointTags_SdkV2) SetDeleteTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["delete_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DeleteTags = types.ListValueMust(t, vs)
}

type PayloadTable_SdkV2 struct {
	Name types.String `tfsdk:"name"`

	Status types.String `tfsdk:"status"`

	StatusMessage types.String `tfsdk:"status_message"`
}

func (to *PayloadTable_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PayloadTable_SdkV2) {
}

func (to *PayloadTable_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PayloadTable_SdkV2) {
}

func (m PayloadTable_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["name"] = attrs["name"].SetOptional()
	attrs["status"] = attrs["status"].SetOptional()
	attrs["status_message"] = attrs["status_message"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PayloadTable.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PayloadTable_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PayloadTable_SdkV2
// only implements ToObjectValue() and Type().
func (m PayloadTable_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":           m.Name,
			"status":         m.Status,
			"status_message": m.StatusMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PayloadTable_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":           types.StringType,
			"status":         types.StringType,
			"status_message": types.StringType,
		},
	}
}

type PtEndpointCoreConfig_SdkV2 struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`

	TrafficConfig types.List `tfsdk:"traffic_config"`
}

func (to *PtEndpointCoreConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PtEndpointCoreConfig_SdkV2) {
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				// Recursively sync the fields of TrafficConfig
				toTrafficConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (to *PtEndpointCoreConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PtEndpointCoreConfig_SdkV2) {
	if !from.ServedEntities.IsNull() && !from.ServedEntities.IsUnknown() && to.ServedEntities.IsNull() && len(from.ServedEntities.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for ServedEntities, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.ServedEntities = from.ServedEntities
	}
	if !from.TrafficConfig.IsNull() && !from.TrafficConfig.IsUnknown() {
		if toTrafficConfig, ok := to.GetTrafficConfig(ctx); ok {
			if fromTrafficConfig, ok := from.GetTrafficConfig(ctx); ok {
				toTrafficConfig.SyncFieldsDuringRead(ctx, fromTrafficConfig)
				to.SetTrafficConfig(ctx, toTrafficConfig)
			}
		}
	}
}

func (m PtEndpointCoreConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PtEndpointCoreConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PtEndpointCoreConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"served_entities": reflect.TypeOf(PtServedModel_SdkV2{}),
		"traffic_config":  reflect.TypeOf(TrafficConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PtEndpointCoreConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m PtEndpointCoreConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entities": m.ServedEntities,
			"traffic_config":  m.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PtEndpointCoreConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entities": basetypes.ListType{
				ElemType: PtServedModel_SdkV2{}.Type(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetServedEntities returns the value of the ServedEntities field in PtEndpointCoreConfig_SdkV2 as
// a slice of PtServedModel_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PtEndpointCoreConfig_SdkV2) GetServedEntities(ctx context.Context) ([]PtServedModel_SdkV2, bool) {
	if m.ServedEntities.IsNull() || m.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []PtServedModel_SdkV2
	d := m.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in PtEndpointCoreConfig_SdkV2.
func (m *PtEndpointCoreConfig_SdkV2) SetServedEntities(ctx context.Context, v []PtServedModel_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ServedEntities = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in PtEndpointCoreConfig_SdkV2 as
// a TrafficConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PtEndpointCoreConfig_SdkV2) GetTrafficConfig(ctx context.Context) (TrafficConfig_SdkV2, bool) {
	var e TrafficConfig_SdkV2
	if m.TrafficConfig.IsNull() || m.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v []TrafficConfig_SdkV2
	d := m.TrafficConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetTrafficConfig sets the value of the TrafficConfig field in PtEndpointCoreConfig_SdkV2.
func (m *PtEndpointCoreConfig_SdkV2) SetTrafficConfig(ctx context.Context, v TrafficConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["traffic_config"]
	m.TrafficConfig = types.ListValueMust(t, vs)
}

type PtServedModel_SdkV2 struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName types.String `tfsdk:"entity_name"`

	EntityVersion types.String `tfsdk:"entity_version"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name types.String `tfsdk:"name"`
	// The number of model units to be provisioned.
	ProvisionedModelUnits types.Int64 `tfsdk:"provisioned_model_units"`
}

func (to *PtServedModel_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PtServedModel_SdkV2) {
}

func (to *PtServedModel_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PtServedModel_SdkV2) {
}

func (m PtServedModel_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_name"] = attrs["entity_name"].SetRequired()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["provisioned_model_units"] = attrs["provisioned_model_units"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PtServedModel.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PtServedModel_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PtServedModel_SdkV2
// only implements ToObjectValue() and Type().
func (m PtServedModel_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":             m.EntityName,
			"entity_version":          m.EntityVersion,
			"name":                    m.Name,
			"provisioned_model_units": m.ProvisionedModelUnits,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PtServedModel_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":             types.StringType,
			"entity_version":          types.StringType,
			"name":                    types.StringType,
			"provisioned_model_units": types.Int64Type,
		},
	}
}

type PutAiGatewayRequest_SdkV2 struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig types.List `tfsdk:"fallback_config"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.List `tfsdk:"guardrails"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.List `tfsdk:"inference_table_config"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.List `tfsdk:"usage_tracking_config"`
}

func (to *PutAiGatewayRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutAiGatewayRequest_SdkV2) {
	if !from.FallbackConfig.IsNull() && !from.FallbackConfig.IsUnknown() {
		if toFallbackConfig, ok := to.GetFallbackConfig(ctx); ok {
			if fromFallbackConfig, ok := from.GetFallbackConfig(ctx); ok {
				// Recursively sync the fields of FallbackConfig
				toFallbackConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromFallbackConfig)
				to.SetFallbackConfig(ctx, toFallbackConfig)
			}
		}
	}
	if !from.Guardrails.IsNull() && !from.Guardrails.IsUnknown() {
		if toGuardrails, ok := to.GetGuardrails(ctx); ok {
			if fromGuardrails, ok := from.GetGuardrails(ctx); ok {
				// Recursively sync the fields of Guardrails
				toGuardrails.SyncFieldsDuringCreateOrUpdate(ctx, fromGuardrails)
				to.SetGuardrails(ctx, toGuardrails)
			}
		}
	}
	if !from.InferenceTableConfig.IsNull() && !from.InferenceTableConfig.IsUnknown() {
		if toInferenceTableConfig, ok := to.GetInferenceTableConfig(ctx); ok {
			if fromInferenceTableConfig, ok := from.GetInferenceTableConfig(ctx); ok {
				// Recursively sync the fields of InferenceTableConfig
				toInferenceTableConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromInferenceTableConfig)
				to.SetInferenceTableConfig(ctx, toInferenceTableConfig)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.UsageTrackingConfig.IsNull() && !from.UsageTrackingConfig.IsUnknown() {
		if toUsageTrackingConfig, ok := to.GetUsageTrackingConfig(ctx); ok {
			if fromUsageTrackingConfig, ok := from.GetUsageTrackingConfig(ctx); ok {
				// Recursively sync the fields of UsageTrackingConfig
				toUsageTrackingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromUsageTrackingConfig)
				to.SetUsageTrackingConfig(ctx, toUsageTrackingConfig)
			}
		}
	}
}

func (to *PutAiGatewayRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutAiGatewayRequest_SdkV2) {
	if !from.FallbackConfig.IsNull() && !from.FallbackConfig.IsUnknown() {
		if toFallbackConfig, ok := to.GetFallbackConfig(ctx); ok {
			if fromFallbackConfig, ok := from.GetFallbackConfig(ctx); ok {
				toFallbackConfig.SyncFieldsDuringRead(ctx, fromFallbackConfig)
				to.SetFallbackConfig(ctx, toFallbackConfig)
			}
		}
	}
	if !from.Guardrails.IsNull() && !from.Guardrails.IsUnknown() {
		if toGuardrails, ok := to.GetGuardrails(ctx); ok {
			if fromGuardrails, ok := from.GetGuardrails(ctx); ok {
				toGuardrails.SyncFieldsDuringRead(ctx, fromGuardrails)
				to.SetGuardrails(ctx, toGuardrails)
			}
		}
	}
	if !from.InferenceTableConfig.IsNull() && !from.InferenceTableConfig.IsUnknown() {
		if toInferenceTableConfig, ok := to.GetInferenceTableConfig(ctx); ok {
			if fromInferenceTableConfig, ok := from.GetInferenceTableConfig(ctx); ok {
				toInferenceTableConfig.SyncFieldsDuringRead(ctx, fromInferenceTableConfig)
				to.SetInferenceTableConfig(ctx, toInferenceTableConfig)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.UsageTrackingConfig.IsNull() && !from.UsageTrackingConfig.IsUnknown() {
		if toUsageTrackingConfig, ok := to.GetUsageTrackingConfig(ctx); ok {
			if fromUsageTrackingConfig, ok := from.GetUsageTrackingConfig(ctx); ok {
				toUsageTrackingConfig.SyncFieldsDuringRead(ctx, fromUsageTrackingConfig)
				to.SetUsageTrackingConfig(ctx, toUsageTrackingConfig)
			}
		}
	}
}

func (m PutAiGatewayRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fallback_config"] = attrs["fallback_config"].SetOptional()
	attrs["fallback_config"] = attrs["fallback_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["guardrails"] = attrs["guardrails"].SetOptional()
	attrs["guardrails"] = attrs["guardrails"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["inference_table_config"] = attrs["inference_table_config"].SetOptional()
	attrs["inference_table_config"] = attrs["inference_table_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAiGatewayRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutAiGatewayRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fallback_config":        reflect.TypeOf(FallbackConfig_SdkV2{}),
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails_SdkV2{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig_SdkV2{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit_SdkV2{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAiGatewayRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PutAiGatewayRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fallback_config":        m.FallbackConfig,
			"guardrails":             m.Guardrails,
			"inference_table_config": m.InferenceTableConfig,
			"name":                   m.Name,
			"rate_limits":            m.RateLimits,
			"usage_tracking_config":  m.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutAiGatewayRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fallback_config": basetypes.ListType{
				ElemType: FallbackConfig_SdkV2{}.Type(ctx),
			},
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails_SdkV2{}.Type(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit_SdkV2{}.Type(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFallbackConfig returns the value of the FallbackConfig field in PutAiGatewayRequest_SdkV2 as
// a FallbackConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayRequest_SdkV2) GetFallbackConfig(ctx context.Context) (FallbackConfig_SdkV2, bool) {
	var e FallbackConfig_SdkV2
	if m.FallbackConfig.IsNull() || m.FallbackConfig.IsUnknown() {
		return e, false
	}
	var v []FallbackConfig_SdkV2
	d := m.FallbackConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFallbackConfig sets the value of the FallbackConfig field in PutAiGatewayRequest_SdkV2.
func (m *PutAiGatewayRequest_SdkV2) SetFallbackConfig(ctx context.Context, v FallbackConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["fallback_config"]
	m.FallbackConfig = types.ListValueMust(t, vs)
}

// GetGuardrails returns the value of the Guardrails field in PutAiGatewayRequest_SdkV2 as
// a AiGatewayGuardrails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayRequest_SdkV2) GetGuardrails(ctx context.Context) (AiGatewayGuardrails_SdkV2, bool) {
	var e AiGatewayGuardrails_SdkV2
	if m.Guardrails.IsNull() || m.Guardrails.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrails_SdkV2
	d := m.Guardrails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGuardrails sets the value of the Guardrails field in PutAiGatewayRequest_SdkV2.
func (m *PutAiGatewayRequest_SdkV2) SetGuardrails(ctx context.Context, v AiGatewayGuardrails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["guardrails"]
	m.Guardrails = types.ListValueMust(t, vs)
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in PutAiGatewayRequest_SdkV2 as
// a AiGatewayInferenceTableConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayRequest_SdkV2) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig_SdkV2, bool) {
	var e AiGatewayInferenceTableConfig_SdkV2
	if m.InferenceTableConfig.IsNull() || m.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayInferenceTableConfig_SdkV2
	d := m.InferenceTableConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in PutAiGatewayRequest_SdkV2.
func (m *PutAiGatewayRequest_SdkV2) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_table_config"]
	m.InferenceTableConfig = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in PutAiGatewayRequest_SdkV2 as
// a slice of AiGatewayRateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayRequest_SdkV2) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit_SdkV2, bool) {
	if m.RateLimits.IsNull() || m.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit_SdkV2
	d := m.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutAiGatewayRequest_SdkV2.
func (m *PutAiGatewayRequest_SdkV2) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in PutAiGatewayRequest_SdkV2 as
// a AiGatewayUsageTrackingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayRequest_SdkV2) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig_SdkV2, bool) {
	var e AiGatewayUsageTrackingConfig_SdkV2
	if m.UsageTrackingConfig.IsNull() || m.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayUsageTrackingConfig_SdkV2
	d := m.UsageTrackingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in PutAiGatewayRequest_SdkV2.
func (m *PutAiGatewayRequest_SdkV2) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_tracking_config"]
	m.UsageTrackingConfig = types.ListValueMust(t, vs)
}

type PutAiGatewayResponse_SdkV2 struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig types.List `tfsdk:"fallback_config"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.List `tfsdk:"guardrails"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.List `tfsdk:"inference_table_config"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.List `tfsdk:"usage_tracking_config"`
}

func (to *PutAiGatewayResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutAiGatewayResponse_SdkV2) {
	if !from.FallbackConfig.IsNull() && !from.FallbackConfig.IsUnknown() {
		if toFallbackConfig, ok := to.GetFallbackConfig(ctx); ok {
			if fromFallbackConfig, ok := from.GetFallbackConfig(ctx); ok {
				// Recursively sync the fields of FallbackConfig
				toFallbackConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromFallbackConfig)
				to.SetFallbackConfig(ctx, toFallbackConfig)
			}
		}
	}
	if !from.Guardrails.IsNull() && !from.Guardrails.IsUnknown() {
		if toGuardrails, ok := to.GetGuardrails(ctx); ok {
			if fromGuardrails, ok := from.GetGuardrails(ctx); ok {
				// Recursively sync the fields of Guardrails
				toGuardrails.SyncFieldsDuringCreateOrUpdate(ctx, fromGuardrails)
				to.SetGuardrails(ctx, toGuardrails)
			}
		}
	}
	if !from.InferenceTableConfig.IsNull() && !from.InferenceTableConfig.IsUnknown() {
		if toInferenceTableConfig, ok := to.GetInferenceTableConfig(ctx); ok {
			if fromInferenceTableConfig, ok := from.GetInferenceTableConfig(ctx); ok {
				// Recursively sync the fields of InferenceTableConfig
				toInferenceTableConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromInferenceTableConfig)
				to.SetInferenceTableConfig(ctx, toInferenceTableConfig)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.UsageTrackingConfig.IsNull() && !from.UsageTrackingConfig.IsUnknown() {
		if toUsageTrackingConfig, ok := to.GetUsageTrackingConfig(ctx); ok {
			if fromUsageTrackingConfig, ok := from.GetUsageTrackingConfig(ctx); ok {
				// Recursively sync the fields of UsageTrackingConfig
				toUsageTrackingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromUsageTrackingConfig)
				to.SetUsageTrackingConfig(ctx, toUsageTrackingConfig)
			}
		}
	}
}

func (to *PutAiGatewayResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutAiGatewayResponse_SdkV2) {
	if !from.FallbackConfig.IsNull() && !from.FallbackConfig.IsUnknown() {
		if toFallbackConfig, ok := to.GetFallbackConfig(ctx); ok {
			if fromFallbackConfig, ok := from.GetFallbackConfig(ctx); ok {
				toFallbackConfig.SyncFieldsDuringRead(ctx, fromFallbackConfig)
				to.SetFallbackConfig(ctx, toFallbackConfig)
			}
		}
	}
	if !from.Guardrails.IsNull() && !from.Guardrails.IsUnknown() {
		if toGuardrails, ok := to.GetGuardrails(ctx); ok {
			if fromGuardrails, ok := from.GetGuardrails(ctx); ok {
				toGuardrails.SyncFieldsDuringRead(ctx, fromGuardrails)
				to.SetGuardrails(ctx, toGuardrails)
			}
		}
	}
	if !from.InferenceTableConfig.IsNull() && !from.InferenceTableConfig.IsUnknown() {
		if toInferenceTableConfig, ok := to.GetInferenceTableConfig(ctx); ok {
			if fromInferenceTableConfig, ok := from.GetInferenceTableConfig(ctx); ok {
				toInferenceTableConfig.SyncFieldsDuringRead(ctx, fromInferenceTableConfig)
				to.SetInferenceTableConfig(ctx, toInferenceTableConfig)
			}
		}
	}
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
	if !from.UsageTrackingConfig.IsNull() && !from.UsageTrackingConfig.IsUnknown() {
		if toUsageTrackingConfig, ok := to.GetUsageTrackingConfig(ctx); ok {
			if fromUsageTrackingConfig, ok := from.GetUsageTrackingConfig(ctx); ok {
				toUsageTrackingConfig.SyncFieldsDuringRead(ctx, fromUsageTrackingConfig)
				to.SetUsageTrackingConfig(ctx, toUsageTrackingConfig)
			}
		}
	}
}

func (m PutAiGatewayResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fallback_config"] = attrs["fallback_config"].SetOptional()
	attrs["fallback_config"] = attrs["fallback_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["guardrails"] = attrs["guardrails"].SetOptional()
	attrs["guardrails"] = attrs["guardrails"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["inference_table_config"] = attrs["inference_table_config"].SetOptional()
	attrs["inference_table_config"] = attrs["inference_table_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAiGatewayResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutAiGatewayResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fallback_config":        reflect.TypeOf(FallbackConfig_SdkV2{}),
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails_SdkV2{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig_SdkV2{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit_SdkV2{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAiGatewayResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PutAiGatewayResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fallback_config":        m.FallbackConfig,
			"guardrails":             m.Guardrails,
			"inference_table_config": m.InferenceTableConfig,
			"rate_limits":            m.RateLimits,
			"usage_tracking_config":  m.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutAiGatewayResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fallback_config": basetypes.ListType{
				ElemType: FallbackConfig_SdkV2{}.Type(ctx),
			},
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails_SdkV2{}.Type(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig_SdkV2{}.Type(ctx),
			},
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit_SdkV2{}.Type(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetFallbackConfig returns the value of the FallbackConfig field in PutAiGatewayResponse_SdkV2 as
// a FallbackConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayResponse_SdkV2) GetFallbackConfig(ctx context.Context) (FallbackConfig_SdkV2, bool) {
	var e FallbackConfig_SdkV2
	if m.FallbackConfig.IsNull() || m.FallbackConfig.IsUnknown() {
		return e, false
	}
	var v []FallbackConfig_SdkV2
	d := m.FallbackConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFallbackConfig sets the value of the FallbackConfig field in PutAiGatewayResponse_SdkV2.
func (m *PutAiGatewayResponse_SdkV2) SetFallbackConfig(ctx context.Context, v FallbackConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["fallback_config"]
	m.FallbackConfig = types.ListValueMust(t, vs)
}

// GetGuardrails returns the value of the Guardrails field in PutAiGatewayResponse_SdkV2 as
// a AiGatewayGuardrails_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayResponse_SdkV2) GetGuardrails(ctx context.Context) (AiGatewayGuardrails_SdkV2, bool) {
	var e AiGatewayGuardrails_SdkV2
	if m.Guardrails.IsNull() || m.Guardrails.IsUnknown() {
		return e, false
	}
	var v []AiGatewayGuardrails_SdkV2
	d := m.Guardrails.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetGuardrails sets the value of the Guardrails field in PutAiGatewayResponse_SdkV2.
func (m *PutAiGatewayResponse_SdkV2) SetGuardrails(ctx context.Context, v AiGatewayGuardrails_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["guardrails"]
	m.Guardrails = types.ListValueMust(t, vs)
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in PutAiGatewayResponse_SdkV2 as
// a AiGatewayInferenceTableConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayResponse_SdkV2) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig_SdkV2, bool) {
	var e AiGatewayInferenceTableConfig_SdkV2
	if m.InferenceTableConfig.IsNull() || m.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayInferenceTableConfig_SdkV2
	d := m.InferenceTableConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in PutAiGatewayResponse_SdkV2.
func (m *PutAiGatewayResponse_SdkV2) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inference_table_config"]
	m.InferenceTableConfig = types.ListValueMust(t, vs)
}

// GetRateLimits returns the value of the RateLimits field in PutAiGatewayResponse_SdkV2 as
// a slice of AiGatewayRateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayResponse_SdkV2) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit_SdkV2, bool) {
	if m.RateLimits.IsNull() || m.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit_SdkV2
	d := m.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutAiGatewayResponse_SdkV2.
func (m *PutAiGatewayResponse_SdkV2) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in PutAiGatewayResponse_SdkV2 as
// a AiGatewayUsageTrackingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *PutAiGatewayResponse_SdkV2) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig_SdkV2, bool) {
	var e AiGatewayUsageTrackingConfig_SdkV2
	if m.UsageTrackingConfig.IsNull() || m.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v []AiGatewayUsageTrackingConfig_SdkV2
	d := m.UsageTrackingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in PutAiGatewayResponse_SdkV2.
func (m *PutAiGatewayResponse_SdkV2) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_tracking_config"]
	m.UsageTrackingConfig = types.ListValueMust(t, vs)
}

type PutRequest_SdkV2 struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits"`
}

func (to *PutRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutRequest_SdkV2) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (to *PutRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutRequest_SdkV2) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (m PutRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rate_limits": reflect.TypeOf(RateLimit_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m PutRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        m.Name,
			"rate_limits": m.RateLimits,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRateLimits returns the value of the RateLimits field in PutRequest_SdkV2 as
// a slice of RateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PutRequest_SdkV2) GetRateLimits(ctx context.Context) ([]RateLimit_SdkV2, bool) {
	if m.RateLimits.IsNull() || m.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit_SdkV2
	d := m.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutRequest_SdkV2.
func (m *PutRequest_SdkV2) SetRateLimits(ctx context.Context, v []RateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RateLimits = types.ListValueMust(t, vs)
}

type PutResponse_SdkV2 struct {
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits"`
}

func (to *PutResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutResponse_SdkV2) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (to *PutResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from PutResponse_SdkV2) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (m PutResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m PutResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rate_limits": reflect.TypeOf(RateLimit_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m PutResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"rate_limits": m.RateLimits,
		})
}

// Type implements basetypes.ObjectValuable.
func (m PutResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRateLimits returns the value of the RateLimits field in PutResponse_SdkV2 as
// a slice of RateLimit_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *PutResponse_SdkV2) GetRateLimits(ctx context.Context) ([]RateLimit_SdkV2, bool) {
	if m.RateLimits.IsNull() || m.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit_SdkV2
	d := m.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutResponse_SdkV2.
func (m *PutResponse_SdkV2) SetRateLimits(ctx context.Context, v []RateLimit_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.RateLimits = types.ListValueMust(t, vs)
}

type QueryEndpointInput_SdkV2 struct {
	// Optional user-provided request identifier that will be recorded in the
	// inference table and the usage tracking table.
	ClientRequestId types.String `tfsdk:"client_request_id"`
	// Pandas Dataframe input in the records orientation.
	DataframeRecords types.List `tfsdk:"dataframe_records"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit types.List `tfsdk:"dataframe_split"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams types.Map `tfsdk:"extra_params"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input types.Object `tfsdk:"input"`
	// Tensor-based input in columnar format.
	Inputs types.Object `tfsdk:"inputs"`
	// Tensor-based input in row format.
	Instances types.List `tfsdk:"instances"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens types.Int64 `tfsdk:"max_tokens"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is an array of ChatMessage objects and should
	// only be used with other chat query fields.
	Messages types.List `tfsdk:"messages"`
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	N types.Int64 `tfsdk:"n"`
	// The name of the serving endpoint. This field is required and is provided
	// via the path parameter.
	Name types.String `tfsdk:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	Prompt types.Object `tfsdk:"prompt"`
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	Stop types.List `tfsdk:"stop"`
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	Stream types.Bool `tfsdk:"stream"`
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	Temperature types.Float64 `tfsdk:"temperature"`
	// Optional user-provided context that will be recorded in the usage
	// tracking table.
	UsageContext types.Map `tfsdk:"usage_context"`
}

func (to *QueryEndpointInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryEndpointInput_SdkV2) {
	if !from.DataframeRecords.IsNull() && !from.DataframeRecords.IsUnknown() && to.DataframeRecords.IsNull() && len(from.DataframeRecords.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataframeRecords, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataframeRecords = from.DataframeRecords
	}
	if !from.DataframeSplit.IsNull() && !from.DataframeSplit.IsUnknown() {
		if toDataframeSplit, ok := to.GetDataframeSplit(ctx); ok {
			if fromDataframeSplit, ok := from.GetDataframeSplit(ctx); ok {
				// Recursively sync the fields of DataframeSplit
				toDataframeSplit.SyncFieldsDuringCreateOrUpdate(ctx, fromDataframeSplit)
				to.SetDataframeSplit(ctx, toDataframeSplit)
			}
		}
	}
	if !from.Instances.IsNull() && !from.Instances.IsUnknown() && to.Instances.IsNull() && len(from.Instances.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Instances, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Instances = from.Instances
	}
	if !from.Messages.IsNull() && !from.Messages.IsUnknown() && to.Messages.IsNull() && len(from.Messages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Messages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Messages = from.Messages
	}
	if !from.Stop.IsNull() && !from.Stop.IsUnknown() && to.Stop.IsNull() && len(from.Stop.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stop, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stop = from.Stop
	}
}

func (to *QueryEndpointInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryEndpointInput_SdkV2) {
	if !from.DataframeRecords.IsNull() && !from.DataframeRecords.IsUnknown() && to.DataframeRecords.IsNull() && len(from.DataframeRecords.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for DataframeRecords, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.DataframeRecords = from.DataframeRecords
	}
	if !from.DataframeSplit.IsNull() && !from.DataframeSplit.IsUnknown() {
		if toDataframeSplit, ok := to.GetDataframeSplit(ctx); ok {
			if fromDataframeSplit, ok := from.GetDataframeSplit(ctx); ok {
				toDataframeSplit.SyncFieldsDuringRead(ctx, fromDataframeSplit)
				to.SetDataframeSplit(ctx, toDataframeSplit)
			}
		}
	}
	if !from.Instances.IsNull() && !from.Instances.IsUnknown() && to.Instances.IsNull() && len(from.Instances.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Instances, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Instances = from.Instances
	}
	if !from.Messages.IsNull() && !from.Messages.IsUnknown() && to.Messages.IsNull() && len(from.Messages.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Messages, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Messages = from.Messages
	}
	if !from.Stop.IsNull() && !from.Stop.IsUnknown() && to.Stop.IsNull() && len(from.Stop.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Stop, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Stop = from.Stop
	}
}

func (m QueryEndpointInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client_request_id"] = attrs["client_request_id"].SetOptional()
	attrs["dataframe_records"] = attrs["dataframe_records"].SetOptional()
	attrs["dataframe_split"] = attrs["dataframe_split"].SetOptional()
	attrs["dataframe_split"] = attrs["dataframe_split"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["extra_params"] = attrs["extra_params"].SetOptional()
	attrs["input"] = attrs["input"].SetOptional()
	attrs["inputs"] = attrs["inputs"].SetOptional()
	attrs["instances"] = attrs["instances"].SetOptional()
	attrs["max_tokens"] = attrs["max_tokens"].SetOptional()
	attrs["messages"] = attrs["messages"].SetOptional()
	attrs["n"] = attrs["n"].SetOptional()
	attrs["prompt"] = attrs["prompt"].SetOptional()
	attrs["stop"] = attrs["stop"].SetOptional()
	attrs["stream"] = attrs["stream"].SetOptional()
	attrs["temperature"] = attrs["temperature"].SetOptional()
	attrs["usage_context"] = attrs["usage_context"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryEndpointInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryEndpointInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataframe_records": reflect.TypeOf(types.Object{}),
		"dataframe_split":   reflect.TypeOf(DataframeSplitInput_SdkV2{}),
		"extra_params":      reflect.TypeOf(types.String{}),
		"instances":         reflect.TypeOf(types.Object{}),
		"messages":          reflect.TypeOf(ChatMessage_SdkV2{}),
		"stop":              reflect.TypeOf(types.String{}),
		"usage_context":     reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEndpointInput_SdkV2
// only implements ToObjectValue() and Type().
func (m QueryEndpointInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_request_id": m.ClientRequestId,
			"dataframe_records": m.DataframeRecords,
			"dataframe_split":   m.DataframeSplit,
			"extra_params":      m.ExtraParams,
			"input":             m.Input,
			"inputs":            m.Inputs,
			"instances":         m.Instances,
			"max_tokens":        m.MaxTokens,
			"messages":          m.Messages,
			"n":                 m.N,
			"name":              m.Name,
			"prompt":            m.Prompt,
			"stop":              m.Stop,
			"stream":            m.Stream,
			"temperature":       m.Temperature,
			"usage_context":     m.UsageContext,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryEndpointInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client_request_id": types.StringType,
			"dataframe_records": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"dataframe_split": basetypes.ListType{
				ElemType: DataframeSplitInput_SdkV2{}.Type(ctx),
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
				ElemType: ChatMessage_SdkV2{}.Type(ctx),
			},
			"n":      types.Int64Type,
			"name":   types.StringType,
			"prompt": types.ObjectType{},
			"stop": basetypes.ListType{
				ElemType: types.StringType,
			},
			"stream":      types.BoolType,
			"temperature": types.Float64Type,
			"usage_context": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDataframeRecords returns the value of the DataframeRecords field in QueryEndpointInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetDataframeRecords(ctx context.Context) ([]types.Object, bool) {
	if m.DataframeRecords.IsNull() || m.DataframeRecords.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.DataframeRecords.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataframeRecords sets the value of the DataframeRecords field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetDataframeRecords(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dataframe_records"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.DataframeRecords = types.ListValueMust(t, vs)
}

// GetDataframeSplit returns the value of the DataframeSplit field in QueryEndpointInput_SdkV2 as
// a DataframeSplitInput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetDataframeSplit(ctx context.Context) (DataframeSplitInput_SdkV2, bool) {
	var e DataframeSplitInput_SdkV2
	if m.DataframeSplit.IsNull() || m.DataframeSplit.IsUnknown() {
		return e, false
	}
	var v []DataframeSplitInput_SdkV2
	d := m.DataframeSplit.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataframeSplit sets the value of the DataframeSplit field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetDataframeSplit(ctx context.Context, v DataframeSplitInput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["dataframe_split"]
	m.DataframeSplit = types.ListValueMust(t, vs)
}

// GetExtraParams returns the value of the ExtraParams field in QueryEndpointInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetExtraParams(ctx context.Context) (map[string]types.String, bool) {
	if m.ExtraParams.IsNull() || m.ExtraParams.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.ExtraParams.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExtraParams sets the value of the ExtraParams field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetExtraParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["extra_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.ExtraParams = types.MapValueMust(t, vs)
}

// GetInstances returns the value of the Instances field in QueryEndpointInput_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetInstances(ctx context.Context) ([]types.Object, bool) {
	if m.Instances.IsNull() || m.Instances.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Instances.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInstances sets the value of the Instances field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetInstances(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Instances = types.ListValueMust(t, vs)
}

// GetMessages returns the value of the Messages field in QueryEndpointInput_SdkV2 as
// a slice of ChatMessage_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetMessages(ctx context.Context) ([]ChatMessage_SdkV2, bool) {
	if m.Messages.IsNull() || m.Messages.IsUnknown() {
		return nil, false
	}
	var v []ChatMessage_SdkV2
	d := m.Messages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessages sets the value of the Messages field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetMessages(ctx context.Context, v []ChatMessage_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Messages = types.ListValueMust(t, vs)
}

// GetStop returns the value of the Stop field in QueryEndpointInput_SdkV2 as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetStop(ctx context.Context) ([]types.String, bool) {
	if m.Stop.IsNull() || m.Stop.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := m.Stop.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetStop sets the value of the Stop field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetStop(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["stop"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Stop = types.ListValueMust(t, vs)
}

// GetUsageContext returns the value of the UsageContext field in QueryEndpointInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointInput_SdkV2) GetUsageContext(ctx context.Context) (map[string]types.String, bool) {
	if m.UsageContext.IsNull() || m.UsageContext.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.UsageContext.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsageContext sets the value of the UsageContext field in QueryEndpointInput_SdkV2.
func (m *QueryEndpointInput_SdkV2) SetUsageContext(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_context"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.UsageContext = types.MapValueMust(t, vs)
}

type QueryEndpointResponse_SdkV2 struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices types.List `tfsdk:"choices"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created types.Int64 `tfsdk:"created"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data types.List `tfsdk:"data"`
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	Id types.String `tfsdk:"id"`
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	Model types.String `tfsdk:"model"`
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	Object types.String `tfsdk:"object"`
	// The predictions returned by the serving endpoint.
	Predictions types.List `tfsdk:"predictions"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName types.String `tfsdk:"-"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage types.List `tfsdk:"usage"`
}

func (to *QueryEndpointResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryEndpointResponse_SdkV2) {
	if !from.Choices.IsNull() && !from.Choices.IsUnknown() && to.Choices.IsNull() && len(from.Choices.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Choices, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Choices = from.Choices
	}
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
	if !from.Predictions.IsNull() && !from.Predictions.IsUnknown() && to.Predictions.IsNull() && len(from.Predictions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Predictions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Predictions = from.Predictions
	}
	if !from.Usage.IsNull() && !from.Usage.IsUnknown() {
		if toUsage, ok := to.GetUsage(ctx); ok {
			if fromUsage, ok := from.GetUsage(ctx); ok {
				// Recursively sync the fields of Usage
				toUsage.SyncFieldsDuringCreateOrUpdate(ctx, fromUsage)
				to.SetUsage(ctx, toUsage)
			}
		}
	}
}

func (to *QueryEndpointResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from QueryEndpointResponse_SdkV2) {
	if !from.Choices.IsNull() && !from.Choices.IsUnknown() && to.Choices.IsNull() && len(from.Choices.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Choices, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Choices = from.Choices
	}
	if !from.Data.IsNull() && !from.Data.IsUnknown() && to.Data.IsNull() && len(from.Data.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Data, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Data = from.Data
	}
	if !from.Predictions.IsNull() && !from.Predictions.IsUnknown() && to.Predictions.IsNull() && len(from.Predictions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Predictions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Predictions = from.Predictions
	}
	if !from.Usage.IsNull() && !from.Usage.IsUnknown() {
		if toUsage, ok := to.GetUsage(ctx); ok {
			if fromUsage, ok := from.GetUsage(ctx); ok {
				toUsage.SyncFieldsDuringRead(ctx, fromUsage)
				to.SetUsage(ctx, toUsage)
			}
		}
	}
}

func (m QueryEndpointResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["choices"] = attrs["choices"].SetOptional()
	attrs["created"] = attrs["created"].SetOptional()
	attrs["data"] = attrs["data"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["model"] = attrs["model"].SetOptional()
	attrs["object"] = attrs["object"].SetOptional()
	attrs["predictions"] = attrs["predictions"].SetOptional()
	attrs["usage"] = attrs["usage"].SetOptional()
	attrs["usage"] = attrs["usage"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["served_model_name"] = attrs["served_model_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryEndpointResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m QueryEndpointResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m QueryEndpointResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"choices":           m.Choices,
			"created":           m.Created,
			"data":              m.Data,
			"id":                m.Id,
			"model":             m.Model,
			"object":            m.Object,
			"predictions":       m.Predictions,
			"served_model_name": m.ServedModelName,
			"usage":             m.Usage,
		})
}

// Type implements basetypes.ObjectValuable.
func (m QueryEndpointResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"choices": basetypes.ListType{
				ElemType: V1ResponseChoiceElement_SdkV2{}.Type(ctx),
			},
			"created": types.Int64Type,
			"data": basetypes.ListType{
				ElemType: EmbeddingsV1ResponseEmbeddingElement_SdkV2{}.Type(ctx),
			},
			"id":     types.StringType,
			"model":  types.StringType,
			"object": types.StringType,
			"predictions": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"served_model_name": types.StringType,
			"usage": basetypes.ListType{
				ElemType: ExternalModelUsageElement_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetChoices returns the value of the Choices field in QueryEndpointResponse_SdkV2 as
// a slice of V1ResponseChoiceElement_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointResponse_SdkV2) GetChoices(ctx context.Context) ([]V1ResponseChoiceElement_SdkV2, bool) {
	if m.Choices.IsNull() || m.Choices.IsUnknown() {
		return nil, false
	}
	var v []V1ResponseChoiceElement_SdkV2
	d := m.Choices.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChoices sets the value of the Choices field in QueryEndpointResponse_SdkV2.
func (m *QueryEndpointResponse_SdkV2) SetChoices(ctx context.Context, v []V1ResponseChoiceElement_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["choices"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Choices = types.ListValueMust(t, vs)
}

// GetData returns the value of the Data field in QueryEndpointResponse_SdkV2 as
// a slice of EmbeddingsV1ResponseEmbeddingElement_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointResponse_SdkV2) GetData(ctx context.Context) ([]EmbeddingsV1ResponseEmbeddingElement_SdkV2, bool) {
	if m.Data.IsNull() || m.Data.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingsV1ResponseEmbeddingElement_SdkV2
	d := m.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in QueryEndpointResponse_SdkV2.
func (m *QueryEndpointResponse_SdkV2) SetData(ctx context.Context, v []EmbeddingsV1ResponseEmbeddingElement_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Data = types.ListValueMust(t, vs)
}

// GetPredictions returns the value of the Predictions field in QueryEndpointResponse_SdkV2 as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointResponse_SdkV2) GetPredictions(ctx context.Context) ([]types.Object, bool) {
	if m.Predictions.IsNull() || m.Predictions.IsUnknown() {
		return nil, false
	}
	var v []types.Object
	d := m.Predictions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPredictions sets the value of the Predictions field in QueryEndpointResponse_SdkV2.
func (m *QueryEndpointResponse_SdkV2) SetPredictions(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["predictions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Predictions = types.ListValueMust(t, vs)
}

// GetUsage returns the value of the Usage field in QueryEndpointResponse_SdkV2 as
// a ExternalModelUsageElement_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *QueryEndpointResponse_SdkV2) GetUsage(ctx context.Context) (ExternalModelUsageElement_SdkV2, bool) {
	var e ExternalModelUsageElement_SdkV2
	if m.Usage.IsNull() || m.Usage.IsUnknown() {
		return e, false
	}
	var v []ExternalModelUsageElement_SdkV2
	d := m.Usage.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetUsage sets the value of the Usage field in QueryEndpointResponse_SdkV2.
func (m *QueryEndpointResponse_SdkV2) SetUsage(ctx context.Context, v ExternalModelUsageElement_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["usage"]
	m.Usage = types.ListValueMust(t, vs)
}

type RateLimit_SdkV2 struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls types.Int64 `tfsdk:"calls"`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key types.String `tfsdk:"key"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod types.String `tfsdk:"renewal_period"`
}

func (to *RateLimit_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RateLimit_SdkV2) {
}

func (to *RateLimit_SdkV2) SyncFieldsDuringRead(ctx context.Context, from RateLimit_SdkV2) {
}

func (m RateLimit_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["calls"] = attrs["calls"].SetRequired()
	attrs["key"] = attrs["key"].SetOptional()
	attrs["renewal_period"] = attrs["renewal_period"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in RateLimit.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m RateLimit_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RateLimit_SdkV2
// only implements ToObjectValue() and Type().
func (m RateLimit_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"calls":          m.Calls,
			"key":            m.Key,
			"renewal_period": m.RenewalPeriod,
		})
}

// Type implements basetypes.ObjectValuable.
func (m RateLimit_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"renewal_period": types.StringType,
		},
	}
}

type Route_SdkV2 struct {
	ServedEntityName types.String `tfsdk:"served_entity_name"`
	// The name of the served model this route configures traffic for.
	ServedModelName types.String `tfsdk:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage types.Int64 `tfsdk:"traffic_percentage"`
}

func (to *Route_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Route_SdkV2) {
}

func (to *Route_SdkV2) SyncFieldsDuringRead(ctx context.Context, from Route_SdkV2) {
}

func (m Route_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["served_entity_name"] = attrs["served_entity_name"].SetOptional()
	attrs["served_model_name"] = attrs["served_model_name"].SetOptional()
	attrs["traffic_percentage"] = attrs["traffic_percentage"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Route.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m Route_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Route_SdkV2
// only implements ToObjectValue() and Type().
func (m Route_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entity_name": m.ServedEntityName,
			"served_model_name":  m.ServedModelName,
			"traffic_percentage": m.TrafficPercentage,
		})
}

// Type implements basetypes.ObjectValuable.
func (m Route_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entity_name": types.StringType,
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
	// **catalog_name.schema_name.model_name**.
	EntityName types.String `tfsdk:"entity_name"`

	EntityVersion types.String `tfsdk:"entity_version"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel types.List `tfsdk:"external_model"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency types.Int64 `tfsdk:"max_provisioned_concurrency"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput types.Int64 `tfsdk:"max_provisioned_throughput"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency types.Int64 `tfsdk:"min_provisioned_concurrency"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput types.Int64 `tfsdk:"min_provisioned_throughput"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name types.String `tfsdk:"name"`
	// The number of model units provisioned.
	ProvisionedModelUnits types.Int64 `tfsdk:"provisioned_model_units"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	WorkloadSize types.String `tfsdk:"workload_size"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type"`
}

func (to *ServedEntityInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedEntityInput_SdkV2) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				// Recursively sync the fields of ExternalModel
				toExternalModel.SyncFieldsDuringCreateOrUpdate(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
}

func (to *ServedEntityInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedEntityInput_SdkV2) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				toExternalModel.SyncFieldsDuringRead(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
}

func (m ServedEntityInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_name"] = attrs["entity_name"].SetOptional()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["environment_vars"] = attrs["environment_vars"].SetOptional()
	attrs["external_model"] = attrs["external_model"].SetOptional()
	attrs["external_model"] = attrs["external_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_provisioned_concurrency"] = attrs["max_provisioned_concurrency"].SetOptional()
	attrs["max_provisioned_throughput"] = attrs["max_provisioned_throughput"].SetOptional()
	attrs["min_provisioned_concurrency"] = attrs["min_provisioned_concurrency"].SetOptional()
	attrs["min_provisioned_throughput"] = attrs["min_provisioned_throughput"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["provisioned_model_units"] = attrs["provisioned_model_units"].SetOptional()
	attrs["scale_to_zero_enabled"] = attrs["scale_to_zero_enabled"].SetOptional()
	attrs["workload_size"] = attrs["workload_size"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedEntityInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedEntityInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"external_model":   reflect.TypeOf(ExternalModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntityInput_SdkV2
// only implements ToObjectValue() and Type().
func (m ServedEntityInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":                 m.EntityName,
			"entity_version":              m.EntityVersion,
			"environment_vars":            m.EnvironmentVars,
			"external_model":              m.ExternalModel,
			"instance_profile_arn":        m.InstanceProfileArn,
			"max_provisioned_concurrency": m.MaxProvisionedConcurrency,
			"max_provisioned_throughput":  m.MaxProvisionedThroughput,
			"min_provisioned_concurrency": m.MinProvisionedConcurrency,
			"min_provisioned_throughput":  m.MinProvisionedThroughput,
			"name":                        m.Name,
			"provisioned_model_units":     m.ProvisionedModelUnits,
			"scale_to_zero_enabled":       m.ScaleToZeroEnabled,
			"workload_size":               m.WorkloadSize,
			"workload_type":               m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedEntityInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"external_model": basetypes.ListType{
				ElemType: ExternalModel_SdkV2{}.Type(ctx),
			},
			"instance_profile_arn":        types.StringType,
			"max_provisioned_concurrency": types.Int64Type,
			"max_provisioned_throughput":  types.Int64Type,
			"min_provisioned_concurrency": types.Int64Type,
			"min_provisioned_throughput":  types.Int64Type,
			"name":                        types.StringType,
			"provisioned_model_units":     types.Int64Type,
			"scale_to_zero_enabled":       types.BoolType,
			"workload_size":               types.StringType,
			"workload_type":               types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedEntityInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntityInput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if m.EnvironmentVars.IsNull() || m.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedEntityInput_SdkV2.
func (m *ServedEntityInput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntityInput_SdkV2 as
// a ExternalModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntityInput_SdkV2) GetExternalModel(ctx context.Context) (ExternalModel_SdkV2, bool) {
	var e ExternalModel_SdkV2
	if m.ExternalModel.IsNull() || m.ExternalModel.IsUnknown() {
		return e, false
	}
	var v []ExternalModel_SdkV2
	d := m.ExternalModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntityInput_SdkV2.
func (m *ServedEntityInput_SdkV2) SetExternalModel(ctx context.Context, v ExternalModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["external_model"]
	m.ExternalModel = types.ListValueMust(t, vs)
}

type ServedEntityOutput_SdkV2 struct {
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`

	Creator types.String `tfsdk:"creator"`
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName types.String `tfsdk:"entity_name"`

	EntityVersion types.String `tfsdk:"entity_version"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel types.List `tfsdk:"external_model"`

	FoundationModel types.List `tfsdk:"foundation_model"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency types.Int64 `tfsdk:"max_provisioned_concurrency"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput types.Int64 `tfsdk:"max_provisioned_throughput"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency types.Int64 `tfsdk:"min_provisioned_concurrency"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput types.Int64 `tfsdk:"min_provisioned_throughput"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name types.String `tfsdk:"name"`
	// The number of model units provisioned.
	ProvisionedModelUnits types.Int64 `tfsdk:"provisioned_model_units"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled"`

	State types.List `tfsdk:"state"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	WorkloadSize types.String `tfsdk:"workload_size"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type"`
}

func (to *ServedEntityOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedEntityOutput_SdkV2) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				// Recursively sync the fields of ExternalModel
				toExternalModel.SyncFieldsDuringCreateOrUpdate(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
	if !from.FoundationModel.IsNull() && !from.FoundationModel.IsUnknown() {
		if toFoundationModel, ok := to.GetFoundationModel(ctx); ok {
			if fromFoundationModel, ok := from.GetFoundationModel(ctx); ok {
				// Recursively sync the fields of FoundationModel
				toFoundationModel.SyncFieldsDuringCreateOrUpdate(ctx, fromFoundationModel)
				to.SetFoundationModel(ctx, toFoundationModel)
			}
		}
	}
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				// Recursively sync the fields of State
				toState.SyncFieldsDuringCreateOrUpdate(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (to *ServedEntityOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedEntityOutput_SdkV2) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				toExternalModel.SyncFieldsDuringRead(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
	if !from.FoundationModel.IsNull() && !from.FoundationModel.IsUnknown() {
		if toFoundationModel, ok := to.GetFoundationModel(ctx); ok {
			if fromFoundationModel, ok := from.GetFoundationModel(ctx); ok {
				toFoundationModel.SyncFieldsDuringRead(ctx, fromFoundationModel)
				to.SetFoundationModel(ctx, toFoundationModel)
			}
		}
	}
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (m ServedEntityOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["entity_name"] = attrs["entity_name"].SetOptional()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["environment_vars"] = attrs["environment_vars"].SetOptional()
	attrs["external_model"] = attrs["external_model"].SetOptional()
	attrs["external_model"] = attrs["external_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["foundation_model"] = attrs["foundation_model"].SetOptional()
	attrs["foundation_model"] = attrs["foundation_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_provisioned_concurrency"] = attrs["max_provisioned_concurrency"].SetOptional()
	attrs["max_provisioned_throughput"] = attrs["max_provisioned_throughput"].SetOptional()
	attrs["min_provisioned_concurrency"] = attrs["min_provisioned_concurrency"].SetOptional()
	attrs["min_provisioned_throughput"] = attrs["min_provisioned_throughput"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["provisioned_model_units"] = attrs["provisioned_model_units"].SetOptional()
	attrs["scale_to_zero_enabled"] = attrs["scale_to_zero_enabled"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state"] = attrs["state"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["workload_size"] = attrs["workload_size"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedEntityOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedEntityOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m ServedEntityOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":          m.CreationTimestamp,
			"creator":                     m.Creator,
			"entity_name":                 m.EntityName,
			"entity_version":              m.EntityVersion,
			"environment_vars":            m.EnvironmentVars,
			"external_model":              m.ExternalModel,
			"foundation_model":            m.FoundationModel,
			"instance_profile_arn":        m.InstanceProfileArn,
			"max_provisioned_concurrency": m.MaxProvisionedConcurrency,
			"max_provisioned_throughput":  m.MaxProvisionedThroughput,
			"min_provisioned_concurrency": m.MinProvisionedConcurrency,
			"min_provisioned_throughput":  m.MinProvisionedThroughput,
			"name":                        m.Name,
			"provisioned_model_units":     m.ProvisionedModelUnits,
			"scale_to_zero_enabled":       m.ScaleToZeroEnabled,
			"state":                       m.State,
			"workload_size":               m.WorkloadSize,
			"workload_type":               m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedEntityOutput_SdkV2) Type(ctx context.Context) attr.Type {
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
				ElemType: ExternalModel_SdkV2{}.Type(ctx),
			},
			"foundation_model": basetypes.ListType{
				ElemType: FoundationModel_SdkV2{}.Type(ctx),
			},
			"instance_profile_arn":        types.StringType,
			"max_provisioned_concurrency": types.Int64Type,
			"max_provisioned_throughput":  types.Int64Type,
			"min_provisioned_concurrency": types.Int64Type,
			"min_provisioned_throughput":  types.Int64Type,
			"name":                        types.StringType,
			"provisioned_model_units":     types.Int64Type,
			"scale_to_zero_enabled":       types.BoolType,
			"state": basetypes.ListType{
				ElemType: ServedModelState_SdkV2{}.Type(ctx),
			},
			"workload_size": types.StringType,
			"workload_type": types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedEntityOutput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntityOutput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if m.EnvironmentVars.IsNull() || m.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedEntityOutput_SdkV2.
func (m *ServedEntityOutput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntityOutput_SdkV2 as
// a ExternalModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntityOutput_SdkV2) GetExternalModel(ctx context.Context) (ExternalModel_SdkV2, bool) {
	var e ExternalModel_SdkV2
	if m.ExternalModel.IsNull() || m.ExternalModel.IsUnknown() {
		return e, false
	}
	var v []ExternalModel_SdkV2
	d := m.ExternalModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntityOutput_SdkV2.
func (m *ServedEntityOutput_SdkV2) SetExternalModel(ctx context.Context, v ExternalModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["external_model"]
	m.ExternalModel = types.ListValueMust(t, vs)
}

// GetFoundationModel returns the value of the FoundationModel field in ServedEntityOutput_SdkV2 as
// a FoundationModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntityOutput_SdkV2) GetFoundationModel(ctx context.Context) (FoundationModel_SdkV2, bool) {
	var e FoundationModel_SdkV2
	if m.FoundationModel.IsNull() || m.FoundationModel.IsUnknown() {
		return e, false
	}
	var v []FoundationModel_SdkV2
	d := m.FoundationModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFoundationModel sets the value of the FoundationModel field in ServedEntityOutput_SdkV2.
func (m *ServedEntityOutput_SdkV2) SetFoundationModel(ctx context.Context, v FoundationModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["foundation_model"]
	m.FoundationModel = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in ServedEntityOutput_SdkV2 as
// a ServedModelState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntityOutput_SdkV2) GetState(ctx context.Context) (ServedModelState_SdkV2, bool) {
	var e ServedModelState_SdkV2
	if m.State.IsNull() || m.State.IsUnknown() {
		return e, false
	}
	var v []ServedModelState_SdkV2
	d := m.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServedEntityOutput_SdkV2.
func (m *ServedEntityOutput_SdkV2) SetState(ctx context.Context, v ServedModelState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	m.State = types.ListValueMust(t, vs)
}

type ServedEntitySpec_SdkV2 struct {
	EntityName types.String `tfsdk:"entity_name"`

	EntityVersion types.String `tfsdk:"entity_version"`

	ExternalModel types.List `tfsdk:"external_model"`

	FoundationModel types.List `tfsdk:"foundation_model"`

	Name types.String `tfsdk:"name"`
}

func (to *ServedEntitySpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedEntitySpec_SdkV2) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				// Recursively sync the fields of ExternalModel
				toExternalModel.SyncFieldsDuringCreateOrUpdate(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
	if !from.FoundationModel.IsNull() && !from.FoundationModel.IsUnknown() {
		if toFoundationModel, ok := to.GetFoundationModel(ctx); ok {
			if fromFoundationModel, ok := from.GetFoundationModel(ctx); ok {
				// Recursively sync the fields of FoundationModel
				toFoundationModel.SyncFieldsDuringCreateOrUpdate(ctx, fromFoundationModel)
				to.SetFoundationModel(ctx, toFoundationModel)
			}
		}
	}
}

func (to *ServedEntitySpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedEntitySpec_SdkV2) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				toExternalModel.SyncFieldsDuringRead(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
	if !from.FoundationModel.IsNull() && !from.FoundationModel.IsUnknown() {
		if toFoundationModel, ok := to.GetFoundationModel(ctx); ok {
			if fromFoundationModel, ok := from.GetFoundationModel(ctx); ok {
				toFoundationModel.SyncFieldsDuringRead(ctx, fromFoundationModel)
				to.SetFoundationModel(ctx, toFoundationModel)
			}
		}
	}
}

func (m ServedEntitySpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_name"] = attrs["entity_name"].SetOptional()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["external_model"] = attrs["external_model"].SetOptional()
	attrs["external_model"] = attrs["external_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["foundation_model"] = attrs["foundation_model"].SetOptional()
	attrs["foundation_model"] = attrs["foundation_model"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedEntitySpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedEntitySpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_model":   reflect.TypeOf(ExternalModel_SdkV2{}),
		"foundation_model": reflect.TypeOf(FoundationModel_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntitySpec_SdkV2
// only implements ToObjectValue() and Type().
func (m ServedEntitySpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":      m.EntityName,
			"entity_version":   m.EntityVersion,
			"external_model":   m.ExternalModel,
			"foundation_model": m.FoundationModel,
			"name":             m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedEntitySpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"external_model": basetypes.ListType{
				ElemType: ExternalModel_SdkV2{}.Type(ctx),
			},
			"foundation_model": basetypes.ListType{
				ElemType: FoundationModel_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntitySpec_SdkV2 as
// a ExternalModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntitySpec_SdkV2) GetExternalModel(ctx context.Context) (ExternalModel_SdkV2, bool) {
	var e ExternalModel_SdkV2
	if m.ExternalModel.IsNull() || m.ExternalModel.IsUnknown() {
		return e, false
	}
	var v []ExternalModel_SdkV2
	d := m.ExternalModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntitySpec_SdkV2.
func (m *ServedEntitySpec_SdkV2) SetExternalModel(ctx context.Context, v ExternalModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["external_model"]
	m.ExternalModel = types.ListValueMust(t, vs)
}

// GetFoundationModel returns the value of the FoundationModel field in ServedEntitySpec_SdkV2 as
// a FoundationModel_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedEntitySpec_SdkV2) GetFoundationModel(ctx context.Context) (FoundationModel_SdkV2, bool) {
	var e FoundationModel_SdkV2
	if m.FoundationModel.IsNull() || m.FoundationModel.IsUnknown() {
		return e, false
	}
	var v []FoundationModel_SdkV2
	d := m.FoundationModel.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetFoundationModel sets the value of the FoundationModel field in ServedEntitySpec_SdkV2.
func (m *ServedEntitySpec_SdkV2) SetFoundationModel(ctx context.Context, v FoundationModel_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["foundation_model"]
	m.FoundationModel = types.ListValueMust(t, vs)
}

type ServedModelInput_SdkV2 struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency types.Int64 `tfsdk:"max_provisioned_concurrency"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput types.Int64 `tfsdk:"max_provisioned_throughput"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency types.Int64 `tfsdk:"min_provisioned_concurrency"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput types.Int64 `tfsdk:"min_provisioned_throughput"`

	ModelName types.String `tfsdk:"model_name"`

	ModelVersion types.String `tfsdk:"model_version"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name types.String `tfsdk:"name"`
	// The number of model units provisioned.
	ProvisionedModelUnits types.Int64 `tfsdk:"provisioned_model_units"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	WorkloadSize types.String `tfsdk:"workload_size"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type"`
}

func (to *ServedModelInput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelInput_SdkV2) {
}

func (to *ServedModelInput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedModelInput_SdkV2) {
}

func (m ServedModelInput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["environment_vars"] = attrs["environment_vars"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_provisioned_concurrency"] = attrs["max_provisioned_concurrency"].SetOptional()
	attrs["max_provisioned_throughput"] = attrs["max_provisioned_throughput"].SetOptional()
	attrs["min_provisioned_concurrency"] = attrs["min_provisioned_concurrency"].SetOptional()
	attrs["min_provisioned_throughput"] = attrs["min_provisioned_throughput"].SetOptional()
	attrs["model_name"] = attrs["model_name"].SetRequired()
	attrs["model_version"] = attrs["model_version"].SetRequired()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["provisioned_model_units"] = attrs["provisioned_model_units"].SetOptional()
	attrs["scale_to_zero_enabled"] = attrs["scale_to_zero_enabled"].SetRequired()
	attrs["workload_size"] = attrs["workload_size"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelInput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedModelInput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelInput_SdkV2
// only implements ToObjectValue() and Type().
func (m ServedModelInput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"environment_vars":            m.EnvironmentVars,
			"instance_profile_arn":        m.InstanceProfileArn,
			"max_provisioned_concurrency": m.MaxProvisionedConcurrency,
			"max_provisioned_throughput":  m.MaxProvisionedThroughput,
			"min_provisioned_concurrency": m.MinProvisionedConcurrency,
			"min_provisioned_throughput":  m.MinProvisionedThroughput,
			"model_name":                  m.ModelName,
			"model_version":               m.ModelVersion,
			"name":                        m.Name,
			"provisioned_model_units":     m.ProvisionedModelUnits,
			"scale_to_zero_enabled":       m.ScaleToZeroEnabled,
			"workload_size":               m.WorkloadSize,
			"workload_type":               m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedModelInput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"instance_profile_arn":        types.StringType,
			"max_provisioned_concurrency": types.Int64Type,
			"max_provisioned_throughput":  types.Int64Type,
			"min_provisioned_concurrency": types.Int64Type,
			"min_provisioned_throughput":  types.Int64Type,
			"model_name":                  types.StringType,
			"model_version":               types.StringType,
			"name":                        types.StringType,
			"provisioned_model_units":     types.Int64Type,
			"scale_to_zero_enabled":       types.BoolType,
			"workload_size":               types.StringType,
			"workload_type":               types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedModelInput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedModelInput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if m.EnvironmentVars.IsNull() || m.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedModelInput_SdkV2.
func (m *ServedModelInput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnvironmentVars = types.MapValueMust(t, vs)
}

type ServedModelOutput_SdkV2 struct {
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`

	Creator types.String `tfsdk:"creator"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars types.Map `tfsdk:"environment_vars"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency types.Int64 `tfsdk:"max_provisioned_concurrency"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency types.Int64 `tfsdk:"min_provisioned_concurrency"`

	ModelName types.String `tfsdk:"model_name"`

	ModelVersion types.String `tfsdk:"model_version"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name types.String `tfsdk:"name"`
	// The number of model units provisioned.
	ProvisionedModelUnits types.Int64 `tfsdk:"provisioned_model_units"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled types.Bool `tfsdk:"scale_to_zero_enabled"`

	State types.List `tfsdk:"state"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	WorkloadSize types.String `tfsdk:"workload_size"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType types.String `tfsdk:"workload_type"`
}

func (to *ServedModelOutput_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelOutput_SdkV2) {
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				// Recursively sync the fields of State
				toState.SyncFieldsDuringCreateOrUpdate(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (to *ServedModelOutput_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedModelOutput_SdkV2) {
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (m ServedModelOutput_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["environment_vars"] = attrs["environment_vars"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_provisioned_concurrency"] = attrs["max_provisioned_concurrency"].SetOptional()
	attrs["min_provisioned_concurrency"] = attrs["min_provisioned_concurrency"].SetOptional()
	attrs["model_name"] = attrs["model_name"].SetOptional()
	attrs["model_version"] = attrs["model_version"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["provisioned_model_units"] = attrs["provisioned_model_units"].SetOptional()
	attrs["scale_to_zero_enabled"] = attrs["scale_to_zero_enabled"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state"] = attrs["state"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["workload_size"] = attrs["workload_size"].SetOptional()
	attrs["workload_type"] = attrs["workload_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedModelOutput_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"state":            reflect.TypeOf(ServedModelState_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelOutput_SdkV2
// only implements ToObjectValue() and Type().
func (m ServedModelOutput_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":          m.CreationTimestamp,
			"creator":                     m.Creator,
			"environment_vars":            m.EnvironmentVars,
			"instance_profile_arn":        m.InstanceProfileArn,
			"max_provisioned_concurrency": m.MaxProvisionedConcurrency,
			"min_provisioned_concurrency": m.MinProvisionedConcurrency,
			"model_name":                  m.ModelName,
			"model_version":               m.ModelVersion,
			"name":                        m.Name,
			"provisioned_model_units":     m.ProvisionedModelUnits,
			"scale_to_zero_enabled":       m.ScaleToZeroEnabled,
			"state":                       m.State,
			"workload_size":               m.WorkloadSize,
			"workload_type":               m.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedModelOutput_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"instance_profile_arn":        types.StringType,
			"max_provisioned_concurrency": types.Int64Type,
			"min_provisioned_concurrency": types.Int64Type,
			"model_name":                  types.StringType,
			"model_version":               types.StringType,
			"name":                        types.StringType,
			"provisioned_model_units":     types.Int64Type,
			"scale_to_zero_enabled":       types.BoolType,
			"state": basetypes.ListType{
				ElemType: ServedModelState_SdkV2{}.Type(ctx),
			},
			"workload_size": types.StringType,
			"workload_type": types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedModelOutput_SdkV2 as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedModelOutput_SdkV2) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
	if m.EnvironmentVars.IsNull() || m.EnvironmentVars.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := m.EnvironmentVars.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedModelOutput_SdkV2.
func (m *ServedModelOutput_SdkV2) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetState returns the value of the State field in ServedModelOutput_SdkV2 as
// a ServedModelState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServedModelOutput_SdkV2) GetState(ctx context.Context) (ServedModelState_SdkV2, bool) {
	var e ServedModelState_SdkV2
	if m.State.IsNull() || m.State.IsUnknown() {
		return e, false
	}
	var v []ServedModelState_SdkV2
	d := m.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServedModelOutput_SdkV2.
func (m *ServedModelOutput_SdkV2) SetState(ctx context.Context, v ServedModelState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	m.State = types.ListValueMust(t, vs)
}

type ServedModelSpec_SdkV2 struct {
	// Only one of model_name and entity_name should be populated
	ModelName types.String `tfsdk:"model_name"`
	// Only one of model_version and entity_version should be populated
	ModelVersion types.String `tfsdk:"model_version"`

	Name types.String `tfsdk:"name"`
}

func (to *ServedModelSpec_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelSpec_SdkV2) {
}

func (to *ServedModelSpec_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedModelSpec_SdkV2) {
}

func (m ServedModelSpec_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["model_name"] = attrs["model_name"].SetOptional()
	attrs["model_version"] = attrs["model_version"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelSpec.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedModelSpec_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelSpec_SdkV2
// only implements ToObjectValue() and Type().
func (m ServedModelSpec_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_name":    m.ModelName,
			"model_version": m.ModelVersion,
			"name":          m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedModelSpec_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_name":    types.StringType,
			"model_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

type ServedModelState_SdkV2 struct {
	Deployment types.String `tfsdk:"deployment"`

	DeploymentStateMessage types.String `tfsdk:"deployment_state_message"`
}

func (to *ServedModelState_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelState_SdkV2) {
}

func (to *ServedModelState_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServedModelState_SdkV2) {
}

func (m ServedModelState_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["deployment"] = attrs["deployment"].SetOptional()
	attrs["deployment_state_message"] = attrs["deployment_state_message"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServedModelState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServedModelState_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelState_SdkV2
// only implements ToObjectValue() and Type().
func (m ServedModelState_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"deployment":               m.Deployment,
			"deployment_state_message": m.DeploymentStateMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServedModelState_SdkV2) Type(ctx context.Context) attr.Type {
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
	Logs types.String `tfsdk:"logs"`
}

func (to *ServerLogsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServerLogsResponse_SdkV2) {
}

func (to *ServerLogsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServerLogsResponse_SdkV2) {
}

func (m ServerLogsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["logs"] = attrs["logs"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServerLogsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServerLogsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServerLogsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ServerLogsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"logs": m.Logs,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServerLogsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ServingEndpoint_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway types.List `tfsdk:"ai_gateway"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The config that is currently being served by the endpoint.
	Config types.List `tfsdk:"config"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator"`
	// Description of the endpoint
	Description types.String `tfsdk:"description"`
	// System-generated ID of the endpoint, included to be used by the
	// Permissions API.
	Id types.String `tfsdk:"id"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// The name of the serving endpoint.
	Name types.String `tfsdk:"name"`
	// Information corresponding to the state of the serving endpoint.
	State types.List `tfsdk:"state"`
	// Tags attached to the serving endpoint.
	Tags types.List `tfsdk:"tags"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task"`
	// The usage policy associated with serving endpoint.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *ServingEndpoint_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpoint_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				// Recursively sync the fields of AiGateway
				toAiGateway.SyncFieldsDuringCreateOrUpdate(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				// Recursively sync the fields of State
				toState.SyncFieldsDuringCreateOrUpdate(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ServingEndpoint_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpoint_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				toAiGateway.SyncFieldsDuringRead(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m ServingEndpoint_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["ai_gateway"] = attrs["ai_gateway"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state"] = attrs["state"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["task"] = attrs["task"].SetOptional()
	attrs["usage_policy_id"] = attrs["usage_policy_id"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpoint.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpoint_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
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
func (m ServingEndpoint_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":             m.AiGateway,
			"budget_policy_id":       m.BudgetPolicyId,
			"config":                 m.Config,
			"creation_timestamp":     m.CreationTimestamp,
			"creator":                m.Creator,
			"description":            m.Description,
			"id":                     m.Id,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"name":                   m.Name,
			"state":                  m.State,
			"tags":                   m.Tags,
			"task":                   m.Task,
			"usage_policy_id":        m.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServingEndpoint_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig_SdkV2{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigSummary_SdkV2{}.Type(ctx),
			},
			"creation_timestamp":     types.Int64Type,
			"creator":                types.StringType,
			"description":            types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"state": basetypes.ListType{
				ElemType: EndpointState_SdkV2{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: EndpointTag_SdkV2{}.Type(ctx),
			},
			"task":            types.StringType,
			"usage_policy_id": types.StringType,
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in ServingEndpoint_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpoint_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if m.AiGateway.IsNull() || m.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := m.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in ServingEndpoint_SdkV2.
func (m *ServingEndpoint_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	m.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in ServingEndpoint_SdkV2 as
// a EndpointCoreConfigSummary_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpoint_SdkV2) GetConfig(ctx context.Context) (EndpointCoreConfigSummary_SdkV2, bool) {
	var e EndpointCoreConfigSummary_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []EndpointCoreConfigSummary_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in ServingEndpoint_SdkV2.
func (m *ServingEndpoint_SdkV2) SetConfig(ctx context.Context, v EndpointCoreConfigSummary_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in ServingEndpoint_SdkV2 as
// a EndpointState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpoint_SdkV2) GetState(ctx context.Context) (EndpointState_SdkV2, bool) {
	var e EndpointState_SdkV2
	if m.State.IsNull() || m.State.IsUnknown() {
		return e, false
	}
	var v []EndpointState_SdkV2
	d := m.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServingEndpoint_SdkV2.
func (m *ServingEndpoint_SdkV2) SetState(ctx context.Context, v EndpointState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	m.State = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ServingEndpoint_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpoint_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ServingEndpoint_SdkV2.
func (m *ServingEndpoint_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ServingEndpointAccessControlRequest_SdkV2 struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ServingEndpointAccessControlRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointAccessControlRequest_SdkV2) {
}

func (to *ServingEndpointAccessControlRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointAccessControlRequest_SdkV2) {
}

func (m ServingEndpointAccessControlRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointAccessControlRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointAccessControlRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointAccessControlRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointAccessControlRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ServingEndpointAccessControlRequest_SdkV2) Type(ctx context.Context) attr.Type {
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

func (to *ServingEndpointAccessControlResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *ServingEndpointAccessControlResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointAccessControlResponse_SdkV2) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (m ServingEndpointAccessControlResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["all_permissions"] = attrs["all_permissions"].SetOptional()
	attrs["display_name"] = attrs["display_name"].SetOptional()
	attrs["group_name"] = attrs["group_name"].SetOptional()
	attrs["service_principal_name"] = attrs["service_principal_name"].SetOptional()
	attrs["user_name"] = attrs["user_name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointAccessControlResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointAccessControlResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ServingEndpointPermission_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointAccessControlResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointAccessControlResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (m ServingEndpointAccessControlResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ServingEndpointPermission_SdkV2{}.Type(ctx),
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
func (m *ServingEndpointAccessControlResponse_SdkV2) GetAllPermissions(ctx context.Context) ([]ServingEndpointPermission_SdkV2, bool) {
	if m.AllPermissions.IsNull() || m.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointPermission_SdkV2
	d := m.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ServingEndpointAccessControlResponse_SdkV2.
func (m *ServingEndpointAccessControlResponse_SdkV2) SetAllPermissions(ctx context.Context, v []ServingEndpointPermission_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AllPermissions = types.ListValueMust(t, vs)
}

type ServingEndpointDetailed_SdkV2 struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway types.List `tfsdk:"ai_gateway"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The config that is currently being served by the endpoint.
	Config types.List `tfsdk:"config"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo types.List `tfsdk:"data_plane_info"`
	// Description of the serving model
	Description types.String `tfsdk:"description"`
	// Email notification settings.
	EmailNotifications types.List `tfsdk:"email_notifications"`
	// Endpoint invocation url if route optimization is enabled for endpoint
	EndpointUrl types.String `tfsdk:"endpoint_url"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id types.String `tfsdk:"id"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp"`
	// The name of the serving endpoint.
	Name types.String `tfsdk:"name"`
	// The config that the endpoint is attempting to update to.
	PendingConfig types.List `tfsdk:"pending_config"`
	// The permission level of the principal making the request.
	PermissionLevel types.String `tfsdk:"permission_level"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized types.Bool `tfsdk:"route_optimized"`
	// Information corresponding to the state of the serving endpoint.
	State types.List `tfsdk:"state"`
	// Tags attached to the serving endpoint.
	Tags types.List `tfsdk:"tags"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task"`
}

func (to *ServingEndpointDetailed_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointDetailed_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				// Recursively sync the fields of AiGateway
				toAiGateway.SyncFieldsDuringCreateOrUpdate(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				// Recursively sync the fields of Config
				toConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.DataPlaneInfo.IsNull() && !from.DataPlaneInfo.IsUnknown() {
		if toDataPlaneInfo, ok := to.GetDataPlaneInfo(ctx); ok {
			if fromDataPlaneInfo, ok := from.GetDataPlaneInfo(ctx); ok {
				// Recursively sync the fields of DataPlaneInfo
				toDataPlaneInfo.SyncFieldsDuringCreateOrUpdate(ctx, fromDataPlaneInfo)
				to.SetDataPlaneInfo(ctx, toDataPlaneInfo)
			}
		}
	}
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				// Recursively sync the fields of EmailNotifications
				toEmailNotifications.SyncFieldsDuringCreateOrUpdate(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
	if !from.PendingConfig.IsNull() && !from.PendingConfig.IsUnknown() {
		if toPendingConfig, ok := to.GetPendingConfig(ctx); ok {
			if fromPendingConfig, ok := from.GetPendingConfig(ctx); ok {
				// Recursively sync the fields of PendingConfig
				toPendingConfig.SyncFieldsDuringCreateOrUpdate(ctx, fromPendingConfig)
				to.SetPendingConfig(ctx, toPendingConfig)
			}
		}
	}
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				// Recursively sync the fields of State
				toState.SyncFieldsDuringCreateOrUpdate(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *ServingEndpointDetailed_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointDetailed_SdkV2) {
	if !from.AiGateway.IsNull() && !from.AiGateway.IsUnknown() {
		if toAiGateway, ok := to.GetAiGateway(ctx); ok {
			if fromAiGateway, ok := from.GetAiGateway(ctx); ok {
				toAiGateway.SyncFieldsDuringRead(ctx, fromAiGateway)
				to.SetAiGateway(ctx, toAiGateway)
			}
		}
	}
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
	if !from.DataPlaneInfo.IsNull() && !from.DataPlaneInfo.IsUnknown() {
		if toDataPlaneInfo, ok := to.GetDataPlaneInfo(ctx); ok {
			if fromDataPlaneInfo, ok := from.GetDataPlaneInfo(ctx); ok {
				toDataPlaneInfo.SyncFieldsDuringRead(ctx, fromDataPlaneInfo)
				to.SetDataPlaneInfo(ctx, toDataPlaneInfo)
			}
		}
	}
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
	if !from.PendingConfig.IsNull() && !from.PendingConfig.IsUnknown() {
		if toPendingConfig, ok := to.GetPendingConfig(ctx); ok {
			if fromPendingConfig, ok := from.GetPendingConfig(ctx); ok {
				toPendingConfig.SyncFieldsDuringRead(ctx, fromPendingConfig)
				to.SetPendingConfig(ctx, toPendingConfig)
			}
		}
	}
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (m ServingEndpointDetailed_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["ai_gateway"] = attrs["ai_gateway"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["data_plane_info"] = attrs["data_plane_info"].SetOptional()
	attrs["data_plane_info"] = attrs["data_plane_info"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["description"] = attrs["description"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["endpoint_url"] = attrs["endpoint_url"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pending_config"] = attrs["pending_config"].SetOptional()
	attrs["pending_config"] = attrs["pending_config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["route_optimized"] = attrs["route_optimized"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
	attrs["state"] = attrs["state"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["tags"] = attrs["tags"].SetOptional()
	attrs["task"] = attrs["task"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointDetailed.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointDetailed_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":          reflect.TypeOf(AiGatewayConfig_SdkV2{}),
		"config":              reflect.TypeOf(EndpointCoreConfigOutput_SdkV2{}),
		"data_plane_info":     reflect.TypeOf(ModelDataPlaneInfo_SdkV2{}),
		"email_notifications": reflect.TypeOf(EmailNotifications_SdkV2{}),
		"pending_config":      reflect.TypeOf(EndpointPendingConfig_SdkV2{}),
		"state":               reflect.TypeOf(EndpointState_SdkV2{}),
		"tags":                reflect.TypeOf(EndpointTag_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointDetailed_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointDetailed_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":             m.AiGateway,
			"budget_policy_id":       m.BudgetPolicyId,
			"config":                 m.Config,
			"creation_timestamp":     m.CreationTimestamp,
			"creator":                m.Creator,
			"data_plane_info":        m.DataPlaneInfo,
			"description":            m.Description,
			"email_notifications":    m.EmailNotifications,
			"endpoint_url":           m.EndpointUrl,
			"id":                     m.Id,
			"last_updated_timestamp": m.LastUpdatedTimestamp,
			"name":                   m.Name,
			"pending_config":         m.PendingConfig,
			"permission_level":       m.PermissionLevel,
			"route_optimized":        m.RouteOptimized,
			"state":                  m.State,
			"tags":                   m.Tags,
			"task":                   m.Task,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServingEndpointDetailed_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig_SdkV2{}.Type(ctx),
			},
			"budget_policy_id": types.StringType,
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigOutput_SdkV2{}.Type(ctx),
			},
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"data_plane_info": basetypes.ListType{
				ElemType: ModelDataPlaneInfo_SdkV2{}.Type(ctx),
			},
			"description": types.StringType,
			"email_notifications": basetypes.ListType{
				ElemType: EmailNotifications_SdkV2{}.Type(ctx),
			},
			"endpoint_url":           types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"pending_config": basetypes.ListType{
				ElemType: EndpointPendingConfig_SdkV2{}.Type(ctx),
			},
			"permission_level": types.StringType,
			"route_optimized":  types.BoolType,
			"state": basetypes.ListType{
				ElemType: EndpointState_SdkV2{}.Type(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: EndpointTag_SdkV2{}.Type(ctx),
			},
			"task": types.StringType,
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in ServingEndpointDetailed_SdkV2 as
// a AiGatewayConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetAiGateway(ctx context.Context) (AiGatewayConfig_SdkV2, bool) {
	var e AiGatewayConfig_SdkV2
	if m.AiGateway.IsNull() || m.AiGateway.IsUnknown() {
		return e, false
	}
	var v []AiGatewayConfig_SdkV2
	d := m.AiGateway.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetAiGateway sets the value of the AiGateway field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetAiGateway(ctx context.Context, v AiGatewayConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["ai_gateway"]
	m.AiGateway = types.ListValueMust(t, vs)
}

// GetConfig returns the value of the Config field in ServingEndpointDetailed_SdkV2 as
// a EndpointCoreConfigOutput_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetConfig(ctx context.Context) (EndpointCoreConfigOutput_SdkV2, bool) {
	var e EndpointCoreConfigOutput_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []EndpointCoreConfigOutput_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetConfig(ctx context.Context, v EndpointCoreConfigOutput_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

// GetDataPlaneInfo returns the value of the DataPlaneInfo field in ServingEndpointDetailed_SdkV2 as
// a ModelDataPlaneInfo_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetDataPlaneInfo(ctx context.Context) (ModelDataPlaneInfo_SdkV2, bool) {
	var e ModelDataPlaneInfo_SdkV2
	if m.DataPlaneInfo.IsNull() || m.DataPlaneInfo.IsUnknown() {
		return e, false
	}
	var v []ModelDataPlaneInfo_SdkV2
	d := m.DataPlaneInfo.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetDataPlaneInfo sets the value of the DataPlaneInfo field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetDataPlaneInfo(ctx context.Context, v ModelDataPlaneInfo_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["data_plane_info"]
	m.DataPlaneInfo = types.ListValueMust(t, vs)
}

// GetEmailNotifications returns the value of the EmailNotifications field in ServingEndpointDetailed_SdkV2 as
// a EmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetEmailNotifications(ctx context.Context) (EmailNotifications_SdkV2, bool) {
	var e EmailNotifications_SdkV2
	if m.EmailNotifications.IsNull() || m.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []EmailNotifications_SdkV2
	d := m.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetEmailNotifications(ctx context.Context, v EmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	m.EmailNotifications = types.ListValueMust(t, vs)
}

// GetPendingConfig returns the value of the PendingConfig field in ServingEndpointDetailed_SdkV2 as
// a EndpointPendingConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetPendingConfig(ctx context.Context) (EndpointPendingConfig_SdkV2, bool) {
	var e EndpointPendingConfig_SdkV2
	if m.PendingConfig.IsNull() || m.PendingConfig.IsUnknown() {
		return e, false
	}
	var v []EndpointPendingConfig_SdkV2
	d := m.PendingConfig.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetPendingConfig sets the value of the PendingConfig field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetPendingConfig(ctx context.Context, v EndpointPendingConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["pending_config"]
	m.PendingConfig = types.ListValueMust(t, vs)
}

// GetState returns the value of the State field in ServingEndpointDetailed_SdkV2 as
// a EndpointState_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetState(ctx context.Context) (EndpointState_SdkV2, bool) {
	var e EndpointState_SdkV2
	if m.State.IsNull() || m.State.IsUnknown() {
		return e, false
	}
	var v []EndpointState_SdkV2
	d := m.State.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetState sets the value of the State field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetState(ctx context.Context, v EndpointState_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["state"]
	m.State = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in ServingEndpointDetailed_SdkV2 as
// a slice of EndpointTag_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointDetailed_SdkV2) GetTags(ctx context.Context) ([]EndpointTag_SdkV2, bool) {
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag_SdkV2
	d := m.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ServingEndpointDetailed_SdkV2.
func (m *ServingEndpointDetailed_SdkV2) SetTags(ctx context.Context, v []EndpointTag_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Tags = types.ListValueMust(t, vs)
}

type ServingEndpointPermission_SdkV2 struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ServingEndpointPermission_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *ServingEndpointPermission_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermission_SdkV2) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (m ServingEndpointPermission_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["inherited"] = attrs["inherited"].SetOptional()
	attrs["inherited_from_object"] = attrs["inherited_from_object"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermission.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointPermission_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermission_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointPermission_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             m.Inherited,
			"inherited_from_object": m.InheritedFromObject,
			"permission_level":      m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServingEndpointPermission_SdkV2) Type(ctx context.Context) attr.Type {
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
func (m *ServingEndpointPermission_SdkV2) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ServingEndpointPermission_SdkV2.
func (m *ServingEndpointPermission_SdkV2) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.InheritedFromObject = types.ListValueMust(t, vs)
}

type ServingEndpointPermissions_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *ServingEndpointPermissions_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ServingEndpointPermissions_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermissions_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ServingEndpointPermissions_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["object_id"] = attrs["object_id"].SetOptional()
	attrs["object_type"] = attrs["object_type"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermissions.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointPermissions_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ServingEndpointAccessControlResponse_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissions_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointPermissions_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"object_id":           m.ObjectId,
			"object_type":         m.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServingEndpointPermissions_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlResponse_SdkV2{}.Type(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ServingEndpointPermissions_SdkV2 as
// a slice of ServingEndpointAccessControlResponse_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointPermissions_SdkV2) GetAccessControlList(ctx context.Context) ([]ServingEndpointAccessControlResponse_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointAccessControlResponse_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ServingEndpointPermissions_SdkV2.
func (m *ServingEndpointPermissions_SdkV2) SetAccessControlList(ctx context.Context, v []ServingEndpointAccessControlResponse_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type ServingEndpointPermissionsDescription_SdkV2 struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ServingEndpointPermissionsDescription_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermissionsDescription_SdkV2) {
}

func (to *ServingEndpointPermissionsDescription_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermissionsDescription_SdkV2) {
}

func (m ServingEndpointPermissionsDescription_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["description"] = attrs["description"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermissionsDescription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointPermissionsDescription_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissionsDescription_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointPermissionsDescription_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      m.Description,
			"permission_level": m.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServingEndpointPermissionsDescription_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ServingEndpointPermissionsRequest_SdkV2 struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (to *ServingEndpointPermissionsRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ServingEndpointPermissionsRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermissionsRequest_SdkV2) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (m ServingEndpointPermissionsRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["access_control_list"] = attrs["access_control_list"].SetOptional()
	attrs["serving_endpoint_id"] = attrs["serving_endpoint_id"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ServingEndpointPermissionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m ServingEndpointPermissionsRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ServingEndpointAccessControlRequest_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissionsRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m ServingEndpointPermissionsRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": m.AccessControlList,
			"serving_endpoint_id": m.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (m ServingEndpointPermissionsRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlRequest_SdkV2{}.Type(ctx),
			},
			"serving_endpoint_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ServingEndpointPermissionsRequest_SdkV2 as
// a slice of ServingEndpointAccessControlRequest_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *ServingEndpointPermissionsRequest_SdkV2) GetAccessControlList(ctx context.Context) ([]ServingEndpointAccessControlRequest_SdkV2, bool) {
	if m.AccessControlList.IsNull() || m.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointAccessControlRequest_SdkV2
	d := m.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ServingEndpointPermissionsRequest_SdkV2.
func (m *ServingEndpointPermissionsRequest_SdkV2) SetAccessControlList(ctx context.Context, v []ServingEndpointAccessControlRequest_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.AccessControlList = types.ListValueMust(t, vs)
}

type TrafficConfig_SdkV2 struct {
	// The list of routes that define traffic to each served entity.
	Routes types.List `tfsdk:"routes"`
}

func (to *TrafficConfig_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrafficConfig_SdkV2) {
	if !from.Routes.IsNull() && !from.Routes.IsUnknown() && to.Routes.IsNull() && len(from.Routes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Routes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Routes = from.Routes
	}
}

func (to *TrafficConfig_SdkV2) SyncFieldsDuringRead(ctx context.Context, from TrafficConfig_SdkV2) {
	if !from.Routes.IsNull() && !from.Routes.IsUnknown() && to.Routes.IsNull() && len(from.Routes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Routes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Routes = from.Routes
	}
}

func (m TrafficConfig_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["routes"] = attrs["routes"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrafficConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m TrafficConfig_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"routes": reflect.TypeOf(Route_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrafficConfig_SdkV2
// only implements ToObjectValue() and Type().
func (m TrafficConfig_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"routes": m.Routes,
		})
}

// Type implements basetypes.ObjectValuable.
func (m TrafficConfig_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"routes": basetypes.ListType{
				ElemType: Route_SdkV2{}.Type(ctx),
			},
		},
	}
}

// GetRoutes returns the value of the Routes field in TrafficConfig_SdkV2 as
// a slice of Route_SdkV2 values.
// If the field is unknown or null, the boolean return value is false.
func (m *TrafficConfig_SdkV2) GetRoutes(ctx context.Context) ([]Route_SdkV2, bool) {
	if m.Routes.IsNull() || m.Routes.IsUnknown() {
		return nil, false
	}
	var v []Route_SdkV2
	d := m.Routes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoutes sets the value of the Routes field in TrafficConfig_SdkV2.
func (m *TrafficConfig_SdkV2) SetRoutes(ctx context.Context, v []Route_SdkV2) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["routes"]
	t = t.(attr.TypeWithElementType).ElementType()
	m.Routes = types.ListValueMust(t, vs)
}

type UpdateInferenceEndpointNotifications_SdkV2 struct {
	// The email notification settings to update. Specify email addresses to
	// notify when endpoint state changes occur.
	EmailNotifications types.List `tfsdk:"email_notifications"`
	// The name of the serving endpoint whose notifications are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateInferenceEndpointNotifications_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInferenceEndpointNotifications_SdkV2) {
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				// Recursively sync the fields of EmailNotifications
				toEmailNotifications.SyncFieldsDuringCreateOrUpdate(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
}

func (to *UpdateInferenceEndpointNotifications_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateInferenceEndpointNotifications_SdkV2) {
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
}

func (m UpdateInferenceEndpointNotifications_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInferenceEndpointNotifications.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateInferenceEndpointNotifications_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_notifications": reflect.TypeOf(EmailNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInferenceEndpointNotifications_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateInferenceEndpointNotifications_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_notifications": m.EmailNotifications,
			"name":                m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateInferenceEndpointNotifications_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_notifications": basetypes.ListType{
				ElemType: EmailNotifications_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetEmailNotifications returns the value of the EmailNotifications field in UpdateInferenceEndpointNotifications_SdkV2 as
// a EmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInferenceEndpointNotifications_SdkV2) GetEmailNotifications(ctx context.Context) (EmailNotifications_SdkV2, bool) {
	var e EmailNotifications_SdkV2
	if m.EmailNotifications.IsNull() || m.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []EmailNotifications_SdkV2
	d := m.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in UpdateInferenceEndpointNotifications_SdkV2.
func (m *UpdateInferenceEndpointNotifications_SdkV2) SetEmailNotifications(ctx context.Context, v EmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	m.EmailNotifications = types.ListValueMust(t, vs)
}

type UpdateInferenceEndpointNotificationsResponse_SdkV2 struct {
	EmailNotifications types.List `tfsdk:"email_notifications"`

	Name types.String `tfsdk:"name"`
}

func (to *UpdateInferenceEndpointNotificationsResponse_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInferenceEndpointNotificationsResponse_SdkV2) {
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				// Recursively sync the fields of EmailNotifications
				toEmailNotifications.SyncFieldsDuringCreateOrUpdate(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
}

func (to *UpdateInferenceEndpointNotificationsResponse_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateInferenceEndpointNotificationsResponse_SdkV2) {
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
}

func (m UpdateInferenceEndpointNotificationsResponse_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateInferenceEndpointNotificationsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateInferenceEndpointNotificationsResponse_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_notifications": reflect.TypeOf(EmailNotifications_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInferenceEndpointNotificationsResponse_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateInferenceEndpointNotificationsResponse_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_notifications": m.EmailNotifications,
			"name":                m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateInferenceEndpointNotificationsResponse_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_notifications": basetypes.ListType{
				ElemType: EmailNotifications_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetEmailNotifications returns the value of the EmailNotifications field in UpdateInferenceEndpointNotificationsResponse_SdkV2 as
// a EmailNotifications_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateInferenceEndpointNotificationsResponse_SdkV2) GetEmailNotifications(ctx context.Context) (EmailNotifications_SdkV2, bool) {
	var e EmailNotifications_SdkV2
	if m.EmailNotifications.IsNull() || m.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v []EmailNotifications_SdkV2
	d := m.EmailNotifications.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetEmailNotifications sets the value of the EmailNotifications field in UpdateInferenceEndpointNotificationsResponse_SdkV2.
func (m *UpdateInferenceEndpointNotificationsResponse_SdkV2) SetEmailNotifications(ctx context.Context, v EmailNotifications_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["email_notifications"]
	m.EmailNotifications = types.ListValueMust(t, vs)
}

type UpdateProvisionedThroughputEndpointConfigRequest_SdkV2 struct {
	Config types.List `tfsdk:"config"`
	// The name of the pt endpoint to update. This field is required.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) {
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

func (to *UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) SyncFieldsDuringRead(ctx context.Context, from UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (m UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetRequired()
	attrs["config"] = attrs["config"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["name"] = attrs["name"].SetRequired()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateProvisionedThroughputEndpointConfigRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(PtEndpointCoreConfig_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProvisionedThroughputEndpointConfigRequest_SdkV2
// only implements ToObjectValue() and Type().
func (m UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config": m.Config,
			"name":   m.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (m UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": basetypes.ListType{
				ElemType: PtEndpointCoreConfig_SdkV2{}.Type(ctx),
			},
			"name": types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in UpdateProvisionedThroughputEndpointConfigRequest_SdkV2 as
// a PtEndpointCoreConfig_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) GetConfig(ctx context.Context) (PtEndpointCoreConfig_SdkV2, bool) {
	var e PtEndpointCoreConfig_SdkV2
	if m.Config.IsNull() || m.Config.IsUnknown() {
		return e, false
	}
	var v []PtEndpointCoreConfig_SdkV2
	d := m.Config.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetConfig sets the value of the Config field in UpdateProvisionedThroughputEndpointConfigRequest_SdkV2.
func (m *UpdateProvisionedThroughputEndpointConfigRequest_SdkV2) SetConfig(ctx context.Context, v PtEndpointCoreConfig_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["config"]
	m.Config = types.ListValueMust(t, vs)
}

type V1ResponseChoiceElement_SdkV2 struct {
	// The finish reason returned by the endpoint.
	FinishReason types.String `tfsdk:"finish_reason"`
	// The index of the choice in the __chat or completions__ response.
	Index types.Int64 `tfsdk:"index"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs types.Int64 `tfsdk:"logprobs"`
	// The message response from the __chat__ endpoint.
	Message types.List `tfsdk:"message"`
	// The text response from the __completions__ endpoint.
	Text types.String `tfsdk:"text"`
}

func (to *V1ResponseChoiceElement_SdkV2) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from V1ResponseChoiceElement_SdkV2) {
	if !from.Message.IsNull() && !from.Message.IsUnknown() {
		if toMessage, ok := to.GetMessage(ctx); ok {
			if fromMessage, ok := from.GetMessage(ctx); ok {
				// Recursively sync the fields of Message
				toMessage.SyncFieldsDuringCreateOrUpdate(ctx, fromMessage)
				to.SetMessage(ctx, toMessage)
			}
		}
	}
}

func (to *V1ResponseChoiceElement_SdkV2) SyncFieldsDuringRead(ctx context.Context, from V1ResponseChoiceElement_SdkV2) {
	if !from.Message.IsNull() && !from.Message.IsUnknown() {
		if toMessage, ok := to.GetMessage(ctx); ok {
			if fromMessage, ok := from.GetMessage(ctx); ok {
				toMessage.SyncFieldsDuringRead(ctx, fromMessage)
				to.SetMessage(ctx, toMessage)
			}
		}
	}
}

func (m V1ResponseChoiceElement_SdkV2) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["finish_reason"] = attrs["finish_reason"].SetOptional()
	attrs["index"] = attrs["index"].SetOptional()
	attrs["logprobs"] = attrs["logprobs"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
	attrs["message"] = attrs["message"].(tfschema.ListNestedAttributeBuilder).AddValidator(listvalidator.SizeAtMost(1)).(tfschema.AttributeBuilder)
	attrs["text"] = attrs["text"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in V1ResponseChoiceElement.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (m V1ResponseChoiceElement_SdkV2) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"message": reflect.TypeOf(ChatMessage_SdkV2{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, V1ResponseChoiceElement_SdkV2
// only implements ToObjectValue() and Type().
func (m V1ResponseChoiceElement_SdkV2) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"finish_reason": m.FinishReason,
			"index":         m.Index,
			"logprobs":      m.Logprobs,
			"message":       m.Message,
			"text":          m.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (m V1ResponseChoiceElement_SdkV2) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"finish_reason": types.StringType,
			"index":         types.Int64Type,
			"logprobs":      types.Int64Type,
			"message": basetypes.ListType{
				ElemType: ChatMessage_SdkV2{}.Type(ctx),
			},
			"text": types.StringType,
		},
	}
}

// GetMessage returns the value of the Message field in V1ResponseChoiceElement_SdkV2 as
// a ChatMessage_SdkV2 value.
// If the field is unknown or null, the boolean return value is false.
func (m *V1ResponseChoiceElement_SdkV2) GetMessage(ctx context.Context) (ChatMessage_SdkV2, bool) {
	var e ChatMessage_SdkV2
	if m.Message.IsNull() || m.Message.IsUnknown() {
		return e, false
	}
	var v []ChatMessage_SdkV2
	d := m.Message.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	if len(v) == 0 {
		return e, false
	}
	return v[0], true
}

// SetMessage sets the value of the Message field in V1ResponseChoiceElement_SdkV2.
func (m *V1ResponseChoiceElement_SdkV2) SetMessage(ctx context.Context, v ChatMessage_SdkV2) {
	vs := []attr.Value{v.ToObjectValue(ctx)}
	t := m.Type(ctx).(basetypes.ObjectType).AttrTypes["message"]
	m.Message = types.ListValueMust(t, vs)
}
