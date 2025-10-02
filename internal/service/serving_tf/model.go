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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Ai21LabsConfig struct {
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

func (to *Ai21LabsConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Ai21LabsConfig) {
}

func (to *Ai21LabsConfig) SyncFieldsDuringRead(ctx context.Context, from Ai21LabsConfig) {
}

func (c Ai21LabsConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Ai21LabsConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Ai21LabsConfig
// only implements ToObjectValue() and Type().
func (o Ai21LabsConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai21labs_api_key":           o.Ai21labsApiKey,
			"ai21labs_api_key_plaintext": o.Ai21labsApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Ai21LabsConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_api_key":           types.StringType,
			"ai21labs_api_key_plaintext": types.StringType,
		},
	}
}

type AiGatewayConfig struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig types.Object `tfsdk:"fallback_config"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.Object `tfsdk:"guardrails"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.Object `tfsdk:"inference_table_config"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.Object `tfsdk:"usage_tracking_config"`
}

func (to *AiGatewayConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayConfig) {
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

func (to *AiGatewayConfig) SyncFieldsDuringRead(ctx context.Context, from AiGatewayConfig) {
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

func (c AiGatewayConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fallback_config"] = attrs["fallback_config"].SetOptional()
	attrs["guardrails"] = attrs["guardrails"].SetOptional()
	attrs["inference_table_config"] = attrs["inference_table_config"].SetOptional()
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fallback_config":        reflect.TypeOf(FallbackConfig{}),
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayConfig
// only implements ToObjectValue() and Type().
func (o AiGatewayConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fallback_config":        o.FallbackConfig,
			"guardrails":             o.Guardrails,
			"inference_table_config": o.InferenceTableConfig,
			"rate_limits":            o.RateLimits,
			"usage_tracking_config":  o.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fallback_config":        FallbackConfig{}.Type(ctx),
			"guardrails":             AiGatewayGuardrails{}.Type(ctx),
			"inference_table_config": AiGatewayInferenceTableConfig{}.Type(ctx),
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.Type(ctx),
			},
			"usage_tracking_config": AiGatewayUsageTrackingConfig{}.Type(ctx),
		},
	}
}

// GetFallbackConfig returns the value of the FallbackConfig field in AiGatewayConfig as
// a FallbackConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig) GetFallbackConfig(ctx context.Context) (FallbackConfig, bool) {
	var e FallbackConfig
	if o.FallbackConfig.IsNull() || o.FallbackConfig.IsUnknown() {
		return e, false
	}
	var v FallbackConfig
	d := o.FallbackConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFallbackConfig sets the value of the FallbackConfig field in AiGatewayConfig.
func (o *AiGatewayConfig) SetFallbackConfig(ctx context.Context, v FallbackConfig) {
	vs := v.ToObjectValue(ctx)
	o.FallbackConfig = vs
}

// GetGuardrails returns the value of the Guardrails field in AiGatewayConfig as
// a AiGatewayGuardrails value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig) GetGuardrails(ctx context.Context) (AiGatewayGuardrails, bool) {
	var e AiGatewayGuardrails
	if o.Guardrails.IsNull() || o.Guardrails.IsUnknown() {
		return e, false
	}
	var v AiGatewayGuardrails
	d := o.Guardrails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGuardrails sets the value of the Guardrails field in AiGatewayConfig.
func (o *AiGatewayConfig) SetGuardrails(ctx context.Context, v AiGatewayGuardrails) {
	vs := v.ToObjectValue(ctx)
	o.Guardrails = vs
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in AiGatewayConfig as
// a AiGatewayInferenceTableConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig, bool) {
	var e AiGatewayInferenceTableConfig
	if o.InferenceTableConfig.IsNull() || o.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v AiGatewayInferenceTableConfig
	d := o.InferenceTableConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in AiGatewayConfig.
func (o *AiGatewayConfig) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig) {
	vs := v.ToObjectValue(ctx)
	o.InferenceTableConfig = vs
}

// GetRateLimits returns the value of the RateLimits field in AiGatewayConfig as
// a slice of AiGatewayRateLimit values.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in AiGatewayConfig.
func (o *AiGatewayConfig) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in AiGatewayConfig as
// a AiGatewayUsageTrackingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayConfig) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig, bool) {
	var e AiGatewayUsageTrackingConfig
	if o.UsageTrackingConfig.IsNull() || o.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v AiGatewayUsageTrackingConfig
	d := o.UsageTrackingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in AiGatewayConfig.
func (o *AiGatewayConfig) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig) {
	vs := v.ToObjectValue(ctx)
	o.UsageTrackingConfig = vs
}

type AiGatewayGuardrailParameters struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	InvalidKeywords types.List `tfsdk:"invalid_keywords"`
	// Configuration for guardrail PII filter.
	Pii types.Object `tfsdk:"pii"`
	// Indicates whether the safety filter is enabled.
	Safety types.Bool `tfsdk:"safety"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics types.List `tfsdk:"valid_topics"`
}

func (to *AiGatewayGuardrailParameters) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayGuardrailParameters) {
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

func (to *AiGatewayGuardrailParameters) SyncFieldsDuringRead(ctx context.Context, from AiGatewayGuardrailParameters) {
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

func (c AiGatewayGuardrailParameters) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["invalid_keywords"] = attrs["invalid_keywords"].SetOptional()
	attrs["pii"] = attrs["pii"].SetOptional()
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
func (a AiGatewayGuardrailParameters) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"invalid_keywords": reflect.TypeOf(types.String{}),
		"pii":              reflect.TypeOf(AiGatewayGuardrailPiiBehavior{}),
		"valid_topics":     reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrailParameters
// only implements ToObjectValue() and Type().
func (o AiGatewayGuardrailParameters) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AiGatewayGuardrailParameters) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"invalid_keywords": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pii":    AiGatewayGuardrailPiiBehavior{}.Type(ctx),
			"safety": types.BoolType,
			"valid_topics": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetInvalidKeywords returns the value of the InvalidKeywords field in AiGatewayGuardrailParameters as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrailParameters) GetInvalidKeywords(ctx context.Context) ([]types.String, bool) {
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

// SetInvalidKeywords sets the value of the InvalidKeywords field in AiGatewayGuardrailParameters.
func (o *AiGatewayGuardrailParameters) SetInvalidKeywords(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["invalid_keywords"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InvalidKeywords = types.ListValueMust(t, vs)
}

// GetPii returns the value of the Pii field in AiGatewayGuardrailParameters as
// a AiGatewayGuardrailPiiBehavior value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrailParameters) GetPii(ctx context.Context) (AiGatewayGuardrailPiiBehavior, bool) {
	var e AiGatewayGuardrailPiiBehavior
	if o.Pii.IsNull() || o.Pii.IsUnknown() {
		return e, false
	}
	var v AiGatewayGuardrailPiiBehavior
	d := o.Pii.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPii sets the value of the Pii field in AiGatewayGuardrailParameters.
func (o *AiGatewayGuardrailParameters) SetPii(ctx context.Context, v AiGatewayGuardrailPiiBehavior) {
	vs := v.ToObjectValue(ctx)
	o.Pii = vs
}

// GetValidTopics returns the value of the ValidTopics field in AiGatewayGuardrailParameters as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrailParameters) GetValidTopics(ctx context.Context) ([]types.String, bool) {
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

// SetValidTopics sets the value of the ValidTopics field in AiGatewayGuardrailParameters.
func (o *AiGatewayGuardrailParameters) SetValidTopics(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["valid_topics"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ValidTopics = types.ListValueMust(t, vs)
}

type AiGatewayGuardrailPiiBehavior struct {
	// Configuration for input guardrail filters.
	Behavior types.String `tfsdk:"behavior"`
}

func (to *AiGatewayGuardrailPiiBehavior) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayGuardrailPiiBehavior) {
}

func (to *AiGatewayGuardrailPiiBehavior) SyncFieldsDuringRead(ctx context.Context, from AiGatewayGuardrailPiiBehavior) {
}

func (c AiGatewayGuardrailPiiBehavior) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AiGatewayGuardrailPiiBehavior) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrailPiiBehavior
// only implements ToObjectValue() and Type().
func (o AiGatewayGuardrailPiiBehavior) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"behavior": o.Behavior,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayGuardrailPiiBehavior) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"behavior": types.StringType,
		},
	}
}

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	Input types.Object `tfsdk:"input"`
	// Configuration for output guardrail filters.
	Output types.Object `tfsdk:"output"`
}

func (to *AiGatewayGuardrails) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayGuardrails) {
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

func (to *AiGatewayGuardrails) SyncFieldsDuringRead(ctx context.Context, from AiGatewayGuardrails) {
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

func (c AiGatewayGuardrails) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["input"] = attrs["input"].SetOptional()
	attrs["output"] = attrs["output"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AiGatewayGuardrails.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AiGatewayGuardrails) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"input":  reflect.TypeOf(AiGatewayGuardrailParameters{}),
		"output": reflect.TypeOf(AiGatewayGuardrailParameters{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayGuardrails
// only implements ToObjectValue() and Type().
func (o AiGatewayGuardrails) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"input":  o.Input,
			"output": o.Output,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayGuardrails) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"input":  AiGatewayGuardrailParameters{}.Type(ctx),
			"output": AiGatewayGuardrailParameters{}.Type(ctx),
		},
	}
}

// GetInput returns the value of the Input field in AiGatewayGuardrails as
// a AiGatewayGuardrailParameters value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrails) GetInput(ctx context.Context) (AiGatewayGuardrailParameters, bool) {
	var e AiGatewayGuardrailParameters
	if o.Input.IsNull() || o.Input.IsUnknown() {
		return e, false
	}
	var v AiGatewayGuardrailParameters
	d := o.Input.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInput sets the value of the Input field in AiGatewayGuardrails.
func (o *AiGatewayGuardrails) SetInput(ctx context.Context, v AiGatewayGuardrailParameters) {
	vs := v.ToObjectValue(ctx)
	o.Input = vs
}

// GetOutput returns the value of the Output field in AiGatewayGuardrails as
// a AiGatewayGuardrailParameters value.
// If the field is unknown or null, the boolean return value is false.
func (o *AiGatewayGuardrails) GetOutput(ctx context.Context) (AiGatewayGuardrailParameters, bool) {
	var e AiGatewayGuardrailParameters
	if o.Output.IsNull() || o.Output.IsUnknown() {
		return e, false
	}
	var v AiGatewayGuardrailParameters
	d := o.Output.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOutput sets the value of the Output field in AiGatewayGuardrails.
func (o *AiGatewayGuardrails) SetOutput(ctx context.Context, v AiGatewayGuardrailParameters) {
	vs := v.ToObjectValue(ctx)
	o.Output = vs
}

type AiGatewayInferenceTableConfig struct {
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

func (to *AiGatewayInferenceTableConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayInferenceTableConfig) {
}

func (to *AiGatewayInferenceTableConfig) SyncFieldsDuringRead(ctx context.Context, from AiGatewayInferenceTableConfig) {
}

func (c AiGatewayInferenceTableConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AiGatewayInferenceTableConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayInferenceTableConfig
// only implements ToObjectValue() and Type().
func (o AiGatewayInferenceTableConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AiGatewayInferenceTableConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"enabled":           types.BoolType,
			"schema_name":       types.StringType,
			"table_name_prefix": types.StringType,
		},
	}
}

type AiGatewayRateLimit struct {
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

func (to *AiGatewayRateLimit) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayRateLimit) {
}

func (to *AiGatewayRateLimit) SyncFieldsDuringRead(ctx context.Context, from AiGatewayRateLimit) {
}

func (c AiGatewayRateLimit) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AiGatewayRateLimit) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayRateLimit
// only implements ToObjectValue() and Type().
func (o AiGatewayRateLimit) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"calls":          o.Calls,
			"key":            o.Key,
			"principal":      o.Principal,
			"renewal_period": o.RenewalPeriod,
			"tokens":         o.Tokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayRateLimit) Type(ctx context.Context) attr.Type {
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

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (to *AiGatewayUsageTrackingConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AiGatewayUsageTrackingConfig) {
}

func (to *AiGatewayUsageTrackingConfig) SyncFieldsDuringRead(ctx context.Context, from AiGatewayUsageTrackingConfig) {
}

func (c AiGatewayUsageTrackingConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AiGatewayUsageTrackingConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AiGatewayUsageTrackingConfig
// only implements ToObjectValue() and Type().
func (o AiGatewayUsageTrackingConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": o.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AiGatewayUsageTrackingConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type AmazonBedrockConfig struct {
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

func (to *AmazonBedrockConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AmazonBedrockConfig) {
}

func (to *AmazonBedrockConfig) SyncFieldsDuringRead(ctx context.Context, from AmazonBedrockConfig) {
}

func (c AmazonBedrockConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AmazonBedrockConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AmazonBedrockConfig
// only implements ToObjectValue() and Type().
func (o AmazonBedrockConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"aws_access_key_id":               o.AwsAccessKeyId,
			"aws_access_key_id_plaintext":     o.AwsAccessKeyIdPlaintext,
			"aws_region":                      o.AwsRegion,
			"aws_secret_access_key":           o.AwsSecretAccessKey,
			"aws_secret_access_key_plaintext": o.AwsSecretAccessKeyPlaintext,
			"bedrock_provider":                o.BedrockProvider,
			"instance_profile_arn":            o.InstanceProfileArn,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AmazonBedrockConfig) Type(ctx context.Context) attr.Type {
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

type AnthropicConfig struct {
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

func (to *AnthropicConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AnthropicConfig) {
}

func (to *AnthropicConfig) SyncFieldsDuringRead(ctx context.Context, from AnthropicConfig) {
}

func (c AnthropicConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AnthropicConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AnthropicConfig
// only implements ToObjectValue() and Type().
func (o AnthropicConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"anthropic_api_key":           o.AnthropicApiKey,
			"anthropic_api_key_plaintext": o.AnthropicApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AnthropicConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anthropic_api_key":           types.StringType,
			"anthropic_api_key_plaintext": types.StringType,
		},
	}
}

type ApiKeyAuth struct {
	// The name of the API key parameter used for authentication.
	Key types.String `tfsdk:"key"`
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	Value types.String `tfsdk:"value"`
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	ValuePlaintext types.String `tfsdk:"value_plaintext"`
}

func (to *ApiKeyAuth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ApiKeyAuth) {
}

func (to *ApiKeyAuth) SyncFieldsDuringRead(ctx context.Context, from ApiKeyAuth) {
}

func (c ApiKeyAuth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ApiKeyAuth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ApiKeyAuth
// only implements ToObjectValue() and Type().
func (o ApiKeyAuth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":             o.Key,
			"value":           o.Value,
			"value_plaintext": o.ValuePlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ApiKeyAuth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":             types.StringType,
			"value":           types.StringType,
			"value_plaintext": types.StringType,
		},
	}
}

type AutoCaptureConfigInput struct {
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

func (to *AutoCaptureConfigInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoCaptureConfigInput) {
}

func (to *AutoCaptureConfigInput) SyncFieldsDuringRead(ctx context.Context, from AutoCaptureConfigInput) {
}

func (c AutoCaptureConfigInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a AutoCaptureConfigInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureConfigInput
// only implements ToObjectValue() and Type().
func (o AutoCaptureConfigInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AutoCaptureConfigInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"enabled":           types.BoolType,
			"schema_name":       types.StringType,
			"table_name_prefix": types.StringType,
		},
	}
}

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName types.String `tfsdk:"catalog_name"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName types.String `tfsdk:"schema_name"`

	State types.Object `tfsdk:"state"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix types.String `tfsdk:"table_name_prefix"`
}

func (to *AutoCaptureConfigOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoCaptureConfigOutput) {
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

func (to *AutoCaptureConfigOutput) SyncFieldsDuringRead(ctx context.Context, from AutoCaptureConfigOutput) {
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (c AutoCaptureConfigOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["catalog_name"] = attrs["catalog_name"].SetOptional()
	attrs["enabled"] = attrs["enabled"].SetOptional()
	attrs["schema_name"] = attrs["schema_name"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
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
func (a AutoCaptureConfigOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"state": reflect.TypeOf(AutoCaptureState{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureConfigOutput
// only implements ToObjectValue() and Type().
func (o AutoCaptureConfigOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o AutoCaptureConfigOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name":      types.StringType,
			"enabled":           types.BoolType,
			"schema_name":       types.StringType,
			"state":             AutoCaptureState{}.Type(ctx),
			"table_name_prefix": types.StringType,
		},
	}
}

// GetState returns the value of the State field in AutoCaptureConfigOutput as
// a AutoCaptureState value.
// If the field is unknown or null, the boolean return value is false.
func (o *AutoCaptureConfigOutput) GetState(ctx context.Context) (AutoCaptureState, bool) {
	var e AutoCaptureState
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v AutoCaptureState
	d := o.State.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetState sets the value of the State field in AutoCaptureConfigOutput.
func (o *AutoCaptureConfigOutput) SetState(ctx context.Context, v AutoCaptureState) {
	vs := v.ToObjectValue(ctx)
	o.State = vs
}

type AutoCaptureState struct {
	PayloadTable types.Object `tfsdk:"payload_table"`
}

func (to *AutoCaptureState) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from AutoCaptureState) {
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

func (to *AutoCaptureState) SyncFieldsDuringRead(ctx context.Context, from AutoCaptureState) {
	if !from.PayloadTable.IsNull() && !from.PayloadTable.IsUnknown() {
		if toPayloadTable, ok := to.GetPayloadTable(ctx); ok {
			if fromPayloadTable, ok := from.GetPayloadTable(ctx); ok {
				toPayloadTable.SyncFieldsDuringRead(ctx, fromPayloadTable)
				to.SetPayloadTable(ctx, toPayloadTable)
			}
		}
	}
}

func (c AutoCaptureState) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["payload_table"] = attrs["payload_table"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in AutoCaptureState.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a AutoCaptureState) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"payload_table": reflect.TypeOf(PayloadTable{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AutoCaptureState
// only implements ToObjectValue() and Type().
func (o AutoCaptureState) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"payload_table": o.PayloadTable,
		})
}

// Type implements basetypes.ObjectValuable.
func (o AutoCaptureState) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"payload_table": PayloadTable{}.Type(ctx),
		},
	}
}

// GetPayloadTable returns the value of the PayloadTable field in AutoCaptureState as
// a PayloadTable value.
// If the field is unknown or null, the boolean return value is false.
func (o *AutoCaptureState) GetPayloadTable(ctx context.Context) (PayloadTable, bool) {
	var e PayloadTable
	if o.PayloadTable.IsNull() || o.PayloadTable.IsUnknown() {
		return e, false
	}
	var v PayloadTable
	d := o.PayloadTable.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPayloadTable sets the value of the PayloadTable field in AutoCaptureState.
func (o *AutoCaptureState) SetPayloadTable(ctx context.Context, v PayloadTable) {
	vs := v.ToObjectValue(ctx)
	o.PayloadTable = vs
}

type BearerTokenAuth struct {
	// The Databricks secret key reference for a token. If you prefer to paste
	// your token directly, see `token_plaintext`.
	Token types.String `tfsdk:"token"`
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	TokenPlaintext types.String `tfsdk:"token_plaintext"`
}

func (to *BearerTokenAuth) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BearerTokenAuth) {
}

func (to *BearerTokenAuth) SyncFieldsDuringRead(ctx context.Context, from BearerTokenAuth) {
}

func (c BearerTokenAuth) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BearerTokenAuth) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BearerTokenAuth
// only implements ToObjectValue() and Type().
func (o BearerTokenAuth) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"token":           o.Token,
			"token_plaintext": o.TokenPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BearerTokenAuth) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"token":           types.StringType,
			"token_plaintext": types.StringType,
		},
	}
}

type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (to *BuildLogsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BuildLogsRequest) {
}

func (to *BuildLogsRequest) SyncFieldsDuringRead(ctx context.Context, from BuildLogsRequest) {
}

func (c BuildLogsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BuildLogsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BuildLogsRequest
// only implements ToObjectValue() and Type().
func (o BuildLogsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              o.Name,
			"served_model_name": o.ServedModelName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BuildLogsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs types.String `tfsdk:"logs"`
}

func (to *BuildLogsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from BuildLogsResponse) {
}

func (to *BuildLogsResponse) SyncFieldsDuringRead(ctx context.Context, from BuildLogsResponse) {
}

func (c BuildLogsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a BuildLogsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, BuildLogsResponse
// only implements ToObjectValue() and Type().
func (o BuildLogsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"logs": o.Logs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o BuildLogsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ChatMessage struct {
	// The content of the message.
	Content types.String `tfsdk:"content"`
	// The role of the message. One of [system, user, assistant].
	Role types.String `tfsdk:"role"`
}

func (to *ChatMessage) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ChatMessage) {
}

func (to *ChatMessage) SyncFieldsDuringRead(ctx context.Context, from ChatMessage) {
}

func (c ChatMessage) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ChatMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ChatMessage
// only implements ToObjectValue() and Type().
func (o ChatMessage) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"content": o.Content,
			"role":    o.Role,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ChatMessage) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"role":    types.StringType,
		},
	}
}

type CohereConfig struct {
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

func (to *CohereConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CohereConfig) {
}

func (to *CohereConfig) SyncFieldsDuringRead(ctx context.Context, from CohereConfig) {
}

func (c CohereConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a CohereConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CohereConfig
// only implements ToObjectValue() and Type().
func (o CohereConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"cohere_api_base":          o.CohereApiBase,
			"cohere_api_key":           o.CohereApiKey,
			"cohere_api_key_plaintext": o.CohereApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CohereConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cohere_api_base":          types.StringType,
			"cohere_api_key":           types.StringType,
			"cohere_api_key_plaintext": types.StringType,
		},
	}
}

type CreatePtEndpointRequest struct {
	// The AI Gateway configuration for the serving endpoint.
	AiGateway types.Object `tfsdk:"ai_gateway"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The core config of the serving endpoint.
	Config types.Object `tfsdk:"config"`
	// Email notification settings.
	EmailNotifications types.Object `tfsdk:"email_notifications"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name types.String `tfsdk:"name"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags types.List `tfsdk:"tags"`
}

func (to *CreatePtEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreatePtEndpointRequest) {
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

func (to *CreatePtEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from CreatePtEndpointRequest) {
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

func (c CreatePtEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetRequired()
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
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
func (a CreatePtEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":          reflect.TypeOf(AiGatewayConfig{}),
		"config":              reflect.TypeOf(PtEndpointCoreConfig{}),
		"email_notifications": reflect.TypeOf(EmailNotifications{}),
		"tags":                reflect.TypeOf(EndpointTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreatePtEndpointRequest
// only implements ToObjectValue() and Type().
func (o CreatePtEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":          o.AiGateway,
			"budget_policy_id":    o.BudgetPolicyId,
			"config":              o.Config,
			"email_notifications": o.EmailNotifications,
			"name":                o.Name,
			"tags":                o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreatePtEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway":          AiGatewayConfig{}.Type(ctx),
			"budget_policy_id":    types.StringType,
			"config":              PtEndpointCoreConfig{}.Type(ctx),
			"email_notifications": EmailNotifications{}.Type(ctx),
			"name":                types.StringType,
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in CreatePtEndpointRequest as
// a AiGatewayConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePtEndpointRequest) GetAiGateway(ctx context.Context) (AiGatewayConfig, bool) {
	var e AiGatewayConfig
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v AiGatewayConfig
	d := o.AiGateway.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAiGateway sets the value of the AiGateway field in CreatePtEndpointRequest.
func (o *CreatePtEndpointRequest) SetAiGateway(ctx context.Context, v AiGatewayConfig) {
	vs := v.ToObjectValue(ctx)
	o.AiGateway = vs
}

// GetConfig returns the value of the Config field in CreatePtEndpointRequest as
// a PtEndpointCoreConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePtEndpointRequest) GetConfig(ctx context.Context) (PtEndpointCoreConfig, bool) {
	var e PtEndpointCoreConfig
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v PtEndpointCoreConfig
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in CreatePtEndpointRequest.
func (o *CreatePtEndpointRequest) SetConfig(ctx context.Context, v PtEndpointCoreConfig) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

// GetEmailNotifications returns the value of the EmailNotifications field in CreatePtEndpointRequest as
// a EmailNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePtEndpointRequest) GetEmailNotifications(ctx context.Context) (EmailNotifications, bool) {
	var e EmailNotifications
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v EmailNotifications
	d := o.EmailNotifications.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailNotifications sets the value of the EmailNotifications field in CreatePtEndpointRequest.
func (o *CreatePtEndpointRequest) SetEmailNotifications(ctx context.Context, v EmailNotifications) {
	vs := v.ToObjectValue(ctx)
	o.EmailNotifications = vs
}

// GetTags returns the value of the Tags field in CreatePtEndpointRequest as
// a slice of EndpointTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreatePtEndpointRequest) GetTags(ctx context.Context) ([]EndpointTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreatePtEndpointRequest.
func (o *CreatePtEndpointRequest) SetTags(ctx context.Context, v []EndpointTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type CreateServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway types.Object `tfsdk:"ai_gateway"`
	// The budget policy to be applied to the serving endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The core config of the serving endpoint.
	Config types.Object `tfsdk:"config"`

	Description types.String `tfsdk:"description"`
	// Email notification settings.
	EmailNotifications types.Object `tfsdk:"email_notifications"`
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

func (to *CreateServingEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CreateServingEndpoint) {
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

func (to *CreateServingEndpoint) SyncFieldsDuringRead(ctx context.Context, from CreateServingEndpoint) {
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

func (c CreateServingEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
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
func (a CreateServingEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":          reflect.TypeOf(AiGatewayConfig{}),
		"config":              reflect.TypeOf(EndpointCoreConfigInput{}),
		"email_notifications": reflect.TypeOf(EmailNotifications{}),
		"rate_limits":         reflect.TypeOf(RateLimit{}),
		"tags":                reflect.TypeOf(EndpointTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CreateServingEndpoint
// only implements ToObjectValue() and Type().
func (o CreateServingEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":          o.AiGateway,
			"budget_policy_id":    o.BudgetPolicyId,
			"config":              o.Config,
			"description":         o.Description,
			"email_notifications": o.EmailNotifications,
			"name":                o.Name,
			"rate_limits":         o.RateLimits,
			"route_optimized":     o.RouteOptimized,
			"tags":                o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CreateServingEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway":          AiGatewayConfig{}.Type(ctx),
			"budget_policy_id":    types.StringType,
			"config":              EndpointCoreConfigInput{}.Type(ctx),
			"description":         types.StringType,
			"email_notifications": EmailNotifications{}.Type(ctx),
			"name":                types.StringType,
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

// GetAiGateway returns the value of the AiGateway field in CreateServingEndpoint as
// a AiGatewayConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint) GetAiGateway(ctx context.Context) (AiGatewayConfig, bool) {
	var e AiGatewayConfig
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v AiGatewayConfig
	d := o.AiGateway.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAiGateway sets the value of the AiGateway field in CreateServingEndpoint.
func (o *CreateServingEndpoint) SetAiGateway(ctx context.Context, v AiGatewayConfig) {
	vs := v.ToObjectValue(ctx)
	o.AiGateway = vs
}

// GetConfig returns the value of the Config field in CreateServingEndpoint as
// a EndpointCoreConfigInput value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint) GetConfig(ctx context.Context) (EndpointCoreConfigInput, bool) {
	var e EndpointCoreConfigInput
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v EndpointCoreConfigInput
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in CreateServingEndpoint.
func (o *CreateServingEndpoint) SetConfig(ctx context.Context, v EndpointCoreConfigInput) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

// GetEmailNotifications returns the value of the EmailNotifications field in CreateServingEndpoint as
// a EmailNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint) GetEmailNotifications(ctx context.Context) (EmailNotifications, bool) {
	var e EmailNotifications
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v EmailNotifications
	d := o.EmailNotifications.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailNotifications sets the value of the EmailNotifications field in CreateServingEndpoint.
func (o *CreateServingEndpoint) SetEmailNotifications(ctx context.Context, v EmailNotifications) {
	vs := v.ToObjectValue(ctx)
	o.EmailNotifications = vs
}

// GetRateLimits returns the value of the RateLimits field in CreateServingEndpoint as
// a slice of RateLimit values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint) GetRateLimits(ctx context.Context) ([]RateLimit, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in CreateServingEndpoint.
func (o *CreateServingEndpoint) SetRateLimits(ctx context.Context, v []RateLimit) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetTags returns the value of the Tags field in CreateServingEndpoint as
// a slice of EndpointTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *CreateServingEndpoint) GetTags(ctx context.Context) ([]EndpointTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in CreateServingEndpoint.
func (o *CreateServingEndpoint) SetTags(ctx context.Context, v []EndpointTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

// Configs needed to create a custom provider model route.
type CustomProviderConfig struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	ApiKeyAuth types.Object `tfsdk:"api_key_auth"`
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	BearerTokenAuth types.Object `tfsdk:"bearer_token_auth"`
	// This is a field to provide the URL of the custom provider API.
	CustomProviderUrl types.String `tfsdk:"custom_provider_url"`
}

func (to *CustomProviderConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from CustomProviderConfig) {
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

func (to *CustomProviderConfig) SyncFieldsDuringRead(ctx context.Context, from CustomProviderConfig) {
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

func (c CustomProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["api_key_auth"] = attrs["api_key_auth"].SetOptional()
	attrs["bearer_token_auth"] = attrs["bearer_token_auth"].SetOptional()
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
func (a CustomProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"api_key_auth":      reflect.TypeOf(ApiKeyAuth{}),
		"bearer_token_auth": reflect.TypeOf(BearerTokenAuth{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, CustomProviderConfig
// only implements ToObjectValue() and Type().
func (o CustomProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"api_key_auth":        o.ApiKeyAuth,
			"bearer_token_auth":   o.BearerTokenAuth,
			"custom_provider_url": o.CustomProviderUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o CustomProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"api_key_auth":        ApiKeyAuth{}.Type(ctx),
			"bearer_token_auth":   BearerTokenAuth{}.Type(ctx),
			"custom_provider_url": types.StringType,
		},
	}
}

// GetApiKeyAuth returns the value of the ApiKeyAuth field in CustomProviderConfig as
// a ApiKeyAuth value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomProviderConfig) GetApiKeyAuth(ctx context.Context) (ApiKeyAuth, bool) {
	var e ApiKeyAuth
	if o.ApiKeyAuth.IsNull() || o.ApiKeyAuth.IsUnknown() {
		return e, false
	}
	var v ApiKeyAuth
	d := o.ApiKeyAuth.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetApiKeyAuth sets the value of the ApiKeyAuth field in CustomProviderConfig.
func (o *CustomProviderConfig) SetApiKeyAuth(ctx context.Context, v ApiKeyAuth) {
	vs := v.ToObjectValue(ctx)
	o.ApiKeyAuth = vs
}

// GetBearerTokenAuth returns the value of the BearerTokenAuth field in CustomProviderConfig as
// a BearerTokenAuth value.
// If the field is unknown or null, the boolean return value is false.
func (o *CustomProviderConfig) GetBearerTokenAuth(ctx context.Context) (BearerTokenAuth, bool) {
	var e BearerTokenAuth
	if o.BearerTokenAuth.IsNull() || o.BearerTokenAuth.IsUnknown() {
		return e, false
	}
	var v BearerTokenAuth
	d := o.BearerTokenAuth.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetBearerTokenAuth sets the value of the BearerTokenAuth field in CustomProviderConfig.
func (o *CustomProviderConfig) SetBearerTokenAuth(ctx context.Context, v BearerTokenAuth) {
	vs := v.ToObjectValue(ctx)
	o.BearerTokenAuth = vs
}

// Details necessary to query this object's API through the DataPlane APIs.
type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails types.String `tfsdk:"authorization_details"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl types.String `tfsdk:"endpoint_url"`
}

func (to *DataPlaneInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataPlaneInfo) {
}

func (to *DataPlaneInfo) SyncFieldsDuringRead(ctx context.Context, from DataPlaneInfo) {
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

type DatabricksModelServingConfig struct {
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

func (to *DatabricksModelServingConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DatabricksModelServingConfig) {
}

func (to *DatabricksModelServingConfig) SyncFieldsDuringRead(ctx context.Context, from DatabricksModelServingConfig) {
}

func (c DatabricksModelServingConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DatabricksModelServingConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DatabricksModelServingConfig
// only implements ToObjectValue() and Type().
func (o DatabricksModelServingConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"databricks_api_token":           o.DatabricksApiToken,
			"databricks_api_token_plaintext": o.DatabricksApiTokenPlaintext,
			"databricks_workspace_url":       o.DatabricksWorkspaceUrl,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DatabricksModelServingConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"databricks_api_token":           types.StringType,
			"databricks_api_token_plaintext": types.StringType,
			"databricks_workspace_url":       types.StringType,
		},
	}
}

type DataframeSplitInput struct {
	// Columns array for the dataframe
	Columns types.List `tfsdk:"columns"`
	// Data array for the dataframe
	Data types.List `tfsdk:"data"`
	// Index array for the dataframe
	Index types.List `tfsdk:"index"`
}

func (to *DataframeSplitInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DataframeSplitInput) {
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

func (to *DataframeSplitInput) SyncFieldsDuringRead(ctx context.Context, from DataframeSplitInput) {
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

func (c DataframeSplitInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DataframeSplitInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"columns": reflect.TypeOf(types.Object{}),
		"data":    reflect.TypeOf(types.Object{}),
		"index":   reflect.TypeOf(types.Int64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DataframeSplitInput
// only implements ToObjectValue() and Type().
func (o DataframeSplitInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"columns": o.Columns,
			"data":    o.Data,
			"index":   o.Index,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DataframeSplitInput) Type(ctx context.Context) attr.Type {
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

// GetColumns returns the value of the Columns field in DataframeSplitInput as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *DataframeSplitInput) GetColumns(ctx context.Context) ([]types.Object, bool) {
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

// SetColumns sets the value of the Columns field in DataframeSplitInput.
func (o *DataframeSplitInput) SetColumns(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["columns"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Columns = types.ListValueMust(t, vs)
}

// GetData returns the value of the Data field in DataframeSplitInput as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *DataframeSplitInput) GetData(ctx context.Context) ([]types.Object, bool) {
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

// SetData sets the value of the Data field in DataframeSplitInput.
func (o *DataframeSplitInput) SetData(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Data = types.ListValueMust(t, vs)
}

// GetIndex returns the value of the Index field in DataframeSplitInput as
// a slice of types.Int64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *DataframeSplitInput) GetIndex(ctx context.Context) ([]types.Int64, bool) {
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

// SetIndex sets the value of the Index field in DataframeSplitInput.
func (o *DataframeSplitInput) SetIndex(ctx context.Context, v []types.Int64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["index"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Index = types.ListValueMust(t, vs)
}

type DeleteServingEndpointRequest struct {
	Name types.String `tfsdk:"-"`
}

func (to *DeleteServingEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from DeleteServingEndpointRequest) {
}

func (to *DeleteServingEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from DeleteServingEndpointRequest) {
}

func (c DeleteServingEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a DeleteServingEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, DeleteServingEndpointRequest
// only implements ToObjectValue() and Type().
func (o DeleteServingEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o DeleteServingEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type EmailNotifications struct {
	// A list of email addresses to be notified when an endpoint fails to update
	// its configuration or state.
	OnUpdateFailure types.List `tfsdk:"on_update_failure"`
	// A list of email addresses to be notified when an endpoint successfully
	// updates its configuration or state.
	OnUpdateSuccess types.List `tfsdk:"on_update_success"`
}

func (to *EmailNotifications) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmailNotifications) {
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

func (to *EmailNotifications) SyncFieldsDuringRead(ctx context.Context, from EmailNotifications) {
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

func (c EmailNotifications) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EmailNotifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"on_update_failure": reflect.TypeOf(types.String{}),
		"on_update_success": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmailNotifications
// only implements ToObjectValue() and Type().
func (o EmailNotifications) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"on_update_failure": o.OnUpdateFailure,
			"on_update_success": o.OnUpdateSuccess,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmailNotifications) Type(ctx context.Context) attr.Type {
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

// GetOnUpdateFailure returns the value of the OnUpdateFailure field in EmailNotifications as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EmailNotifications) GetOnUpdateFailure(ctx context.Context) ([]types.String, bool) {
	if o.OnUpdateFailure.IsNull() || o.OnUpdateFailure.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnUpdateFailure.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnUpdateFailure sets the value of the OnUpdateFailure field in EmailNotifications.
func (o *EmailNotifications) SetOnUpdateFailure(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_update_failure"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnUpdateFailure = types.ListValueMust(t, vs)
}

// GetOnUpdateSuccess returns the value of the OnUpdateSuccess field in EmailNotifications as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *EmailNotifications) GetOnUpdateSuccess(ctx context.Context) ([]types.String, bool) {
	if o.OnUpdateSuccess.IsNull() || o.OnUpdateSuccess.IsUnknown() {
		return nil, false
	}
	var v []types.String
	d := o.OnUpdateSuccess.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOnUpdateSuccess sets the value of the OnUpdateSuccess field in EmailNotifications.
func (o *EmailNotifications) SetOnUpdateSuccess(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["on_update_success"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.OnUpdateSuccess = types.ListValueMust(t, vs)
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	// The embedding vector
	Embedding types.List `tfsdk:"embedding"`
	// The index of the embedding in the response.
	Index types.Int64 `tfsdk:"index"`
	// This will always be 'embedding'.
	Object types.String `tfsdk:"object"`
}

func (to *EmbeddingsV1ResponseEmbeddingElement) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EmbeddingsV1ResponseEmbeddingElement) {
	if !from.Embedding.IsNull() && !from.Embedding.IsUnknown() && to.Embedding.IsNull() && len(from.Embedding.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Embedding, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Embedding = from.Embedding
	}
}

func (to *EmbeddingsV1ResponseEmbeddingElement) SyncFieldsDuringRead(ctx context.Context, from EmbeddingsV1ResponseEmbeddingElement) {
	if !from.Embedding.IsNull() && !from.Embedding.IsUnknown() && to.Embedding.IsNull() && len(from.Embedding.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Embedding, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Embedding = from.Embedding
	}
}

func (c EmbeddingsV1ResponseEmbeddingElement) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EmbeddingsV1ResponseEmbeddingElement) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"embedding": reflect.TypeOf(types.Float64{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EmbeddingsV1ResponseEmbeddingElement
// only implements ToObjectValue() and Type().
func (o EmbeddingsV1ResponseEmbeddingElement) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"embedding": o.Embedding,
			"index":     o.Index,
			"object":    o.Object,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EmbeddingsV1ResponseEmbeddingElement) Type(ctx context.Context) attr.Type {
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

// GetEmbedding returns the value of the Embedding field in EmbeddingsV1ResponseEmbeddingElement as
// a slice of types.Float64 values.
// If the field is unknown or null, the boolean return value is false.
func (o *EmbeddingsV1ResponseEmbeddingElement) GetEmbedding(ctx context.Context) ([]types.Float64, bool) {
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

// SetEmbedding sets the value of the Embedding field in EmbeddingsV1ResponseEmbeddingElement.
func (o *EmbeddingsV1ResponseEmbeddingElement) SetEmbedding(ctx context.Context, v []types.Float64) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["embedding"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Embedding = types.ListValueMust(t, vs)
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig types.Object `tfsdk:"auto_capture_config"`
	// The name of the serving endpoint to update. This field is required.
	Name types.String `tfsdk:"-"`
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig types.Object `tfsdk:"traffic_config"`
}

func (to *EndpointCoreConfigInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointCoreConfigInput) {
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

func (to *EndpointCoreConfigInput) SyncFieldsDuringRead(ctx context.Context, from EndpointCoreConfigInput) {
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

func (c EndpointCoreConfigInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_capture_config"] = attrs["auto_capture_config"].SetOptional()
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()
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
func (a EndpointCoreConfigInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_capture_config": reflect.TypeOf(AutoCaptureConfigInput{}),
		"served_entities":     reflect.TypeOf(ServedEntityInput{}),
		"served_models":       reflect.TypeOf(ServedModelInput{}),
		"traffic_config":      reflect.TypeOf(TrafficConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigInput
// only implements ToObjectValue() and Type().
func (o EndpointCoreConfigInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EndpointCoreConfigInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": AutoCaptureConfigInput{}.Type(ctx),
			"name":                types.StringType,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityInput{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelInput{}.Type(ctx),
			},
			"traffic_config": TrafficConfig{}.Type(ctx),
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointCoreConfigInput as
// a AutoCaptureConfigInput value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigInput, bool) {
	var e AutoCaptureConfigInput
	if o.AutoCaptureConfig.IsNull() || o.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v AutoCaptureConfigInput
	d := o.AutoCaptureConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointCoreConfigInput.
func (o *EndpointCoreConfigInput) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigInput) {
	vs := v.ToObjectValue(ctx)
	o.AutoCaptureConfig = vs
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigInput as
// a slice of ServedEntityInput values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput) GetServedEntities(ctx context.Context) ([]ServedEntityInput, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityInput
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigInput.
func (o *EndpointCoreConfigInput) SetServedEntities(ctx context.Context, v []ServedEntityInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigInput as
// a slice of ServedModelInput values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput) GetServedModels(ctx context.Context) ([]ServedModelInput, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelInput
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigInput.
func (o *EndpointCoreConfigInput) SetServedModels(ctx context.Context, v []ServedModelInput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointCoreConfigInput as
// a TrafficConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigInput) GetTrafficConfig(ctx context.Context) (TrafficConfig, bool) {
	var e TrafficConfig
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v TrafficConfig
	d := o.TrafficConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointCoreConfigInput.
func (o *EndpointCoreConfigInput) SetTrafficConfig(ctx context.Context, v TrafficConfig) {
	vs := v.ToObjectValue(ctx)
	o.TrafficConfig = vs
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig types.Object `tfsdk:"auto_capture_config"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version"`
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig types.Object `tfsdk:"traffic_config"`
}

func (to *EndpointCoreConfigOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointCoreConfigOutput) {
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

func (to *EndpointCoreConfigOutput) SyncFieldsDuringRead(ctx context.Context, from EndpointCoreConfigOutput) {
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

func (c EndpointCoreConfigOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_capture_config"] = attrs["auto_capture_config"].SetOptional()
	attrs["config_version"] = attrs["config_version"].SetOptional()
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointCoreConfigOutput.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointCoreConfigOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_capture_config": reflect.TypeOf(AutoCaptureConfigOutput{}),
		"served_entities":     reflect.TypeOf(ServedEntityOutput{}),
		"served_models":       reflect.TypeOf(ServedModelOutput{}),
		"traffic_config":      reflect.TypeOf(TrafficConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigOutput
// only implements ToObjectValue() and Type().
func (o EndpointCoreConfigOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EndpointCoreConfigOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": AutoCaptureConfigOutput{}.Type(ctx),
			"config_version":      types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput{}.Type(ctx),
			},
			"traffic_config": TrafficConfig{}.Type(ctx),
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointCoreConfigOutput as
// a AutoCaptureConfigOutput value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigOutput, bool) {
	var e AutoCaptureConfigOutput
	if o.AutoCaptureConfig.IsNull() || o.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v AutoCaptureConfigOutput
	d := o.AutoCaptureConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointCoreConfigOutput.
func (o *EndpointCoreConfigOutput) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigOutput) {
	vs := v.ToObjectValue(ctx)
	o.AutoCaptureConfig = vs
}

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigOutput as
// a slice of ServedEntityOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput) GetServedEntities(ctx context.Context) ([]ServedEntityOutput, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityOutput
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigOutput.
func (o *EndpointCoreConfigOutput) SetServedEntities(ctx context.Context, v []ServedEntityOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigOutput as
// a slice of ServedModelOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput) GetServedModels(ctx context.Context) ([]ServedModelOutput, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelOutput
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigOutput.
func (o *EndpointCoreConfigOutput) SetServedModels(ctx context.Context, v []ServedModelOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointCoreConfigOutput as
// a TrafficConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigOutput) GetTrafficConfig(ctx context.Context) (TrafficConfig, bool) {
	var e TrafficConfig
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v TrafficConfig
	d := o.TrafficConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointCoreConfigOutput.
func (o *EndpointCoreConfigOutput) SetTrafficConfig(ctx context.Context, v TrafficConfig) {
	vs := v.ToObjectValue(ctx)
	o.TrafficConfig = vs
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models"`
}

func (to *EndpointCoreConfigSummary) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointCoreConfigSummary) {
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

func (to *EndpointCoreConfigSummary) SyncFieldsDuringRead(ctx context.Context, from EndpointCoreConfigSummary) {
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

func (c EndpointCoreConfigSummary) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EndpointCoreConfigSummary) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"served_entities": reflect.TypeOf(ServedEntitySpec{}),
		"served_models":   reflect.TypeOf(ServedModelSpec{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointCoreConfigSummary
// only implements ToObjectValue() and Type().
func (o EndpointCoreConfigSummary) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entities": o.ServedEntities,
			"served_models":   o.ServedModels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointCoreConfigSummary) Type(ctx context.Context) attr.Type {
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

// GetServedEntities returns the value of the ServedEntities field in EndpointCoreConfigSummary as
// a slice of ServedEntitySpec values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigSummary) GetServedEntities(ctx context.Context) ([]ServedEntitySpec, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntitySpec
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointCoreConfigSummary.
func (o *EndpointCoreConfigSummary) SetServedEntities(ctx context.Context, v []ServedEntitySpec) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointCoreConfigSummary as
// a slice of ServedModelSpec values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointCoreConfigSummary) GetServedModels(ctx context.Context) ([]ServedModelSpec, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelSpec
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointCoreConfigSummary.
func (o *EndpointCoreConfigSummary) SetServedModels(ctx context.Context, v []ServedModelSpec) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

type EndpointPendingConfig struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig types.Object `tfsdk:"auto_capture_config"`
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
	TrafficConfig types.Object `tfsdk:"traffic_config"`
}

func (to *EndpointPendingConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointPendingConfig) {
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

func (to *EndpointPendingConfig) SyncFieldsDuringRead(ctx context.Context, from EndpointPendingConfig) {
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

func (c EndpointPendingConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["auto_capture_config"] = attrs["auto_capture_config"].SetOptional()
	attrs["config_version"] = attrs["config_version"].SetOptional()
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["served_models"] = attrs["served_models"].SetOptional()
	attrs["start_time"] = attrs["start_time"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in EndpointPendingConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a EndpointPendingConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"auto_capture_config": reflect.TypeOf(AutoCaptureConfigOutput{}),
		"served_entities":     reflect.TypeOf(ServedEntityOutput{}),
		"served_models":       reflect.TypeOf(ServedModelOutput{}),
		"traffic_config":      reflect.TypeOf(TrafficConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointPendingConfig
// only implements ToObjectValue() and Type().
func (o EndpointPendingConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o EndpointPendingConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": AutoCaptureConfigOutput{}.Type(ctx),
			"config_version":      types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput{}.Type(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput{}.Type(ctx),
			},
			"start_time":     types.Int64Type,
			"traffic_config": TrafficConfig{}.Type(ctx),
		},
	}
}

// GetAutoCaptureConfig returns the value of the AutoCaptureConfig field in EndpointPendingConfig as
// a AutoCaptureConfigOutput value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig) GetAutoCaptureConfig(ctx context.Context) (AutoCaptureConfigOutput, bool) {
	var e AutoCaptureConfigOutput
	if o.AutoCaptureConfig.IsNull() || o.AutoCaptureConfig.IsUnknown() {
		return e, false
	}
	var v AutoCaptureConfigOutput
	d := o.AutoCaptureConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAutoCaptureConfig sets the value of the AutoCaptureConfig field in EndpointPendingConfig.
func (o *EndpointPendingConfig) SetAutoCaptureConfig(ctx context.Context, v AutoCaptureConfigOutput) {
	vs := v.ToObjectValue(ctx)
	o.AutoCaptureConfig = vs
}

// GetServedEntities returns the value of the ServedEntities field in EndpointPendingConfig as
// a slice of ServedEntityOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig) GetServedEntities(ctx context.Context) ([]ServedEntityOutput, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []ServedEntityOutput
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in EndpointPendingConfig.
func (o *EndpointPendingConfig) SetServedEntities(ctx context.Context, v []ServedEntityOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetServedModels returns the value of the ServedModels field in EndpointPendingConfig as
// a slice of ServedModelOutput values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig) GetServedModels(ctx context.Context) ([]ServedModelOutput, bool) {
	if o.ServedModels.IsNull() || o.ServedModels.IsUnknown() {
		return nil, false
	}
	var v []ServedModelOutput
	d := o.ServedModels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedModels sets the value of the ServedModels field in EndpointPendingConfig.
func (o *EndpointPendingConfig) SetServedModels(ctx context.Context, v []ServedModelOutput) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_models"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedModels = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in EndpointPendingConfig as
// a TrafficConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointPendingConfig) GetTrafficConfig(ctx context.Context) (TrafficConfig, bool) {
	var e TrafficConfig
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v TrafficConfig
	d := o.TrafficConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrafficConfig sets the value of the TrafficConfig field in EndpointPendingConfig.
func (o *EndpointPendingConfig) SetTrafficConfig(ctx context.Context, v TrafficConfig) {
	vs := v.ToObjectValue(ctx)
	o.TrafficConfig = vs
}

type EndpointState struct {
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

func (to *EndpointState) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointState) {
}

func (to *EndpointState) SyncFieldsDuringRead(ctx context.Context, from EndpointState) {
}

func (c EndpointState) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EndpointState) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointState
// only implements ToObjectValue() and Type().
func (o EndpointState) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config_update": o.ConfigUpdate,
			"ready":         o.Ready,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointState) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_update": types.StringType,
			"ready":         types.StringType,
		},
	}
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key types.String `tfsdk:"key"`
	// Optional value field for a serving endpoint tag.
	Value types.String `tfsdk:"value"`
}

func (to *EndpointTag) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTag) {
}

func (to *EndpointTag) SyncFieldsDuringRead(ctx context.Context, from EndpointTag) {
}

func (c EndpointTag) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EndpointTag) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTag
// only implements ToObjectValue() and Type().
func (o EndpointTag) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"key":   o.Key,
			"value": o.Value,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTag) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

type EndpointTags struct {
	Tags types.List `tfsdk:"tags"`
}

func (to *EndpointTags) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from EndpointTags) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (to *EndpointTags) SyncFieldsDuringRead(ctx context.Context, from EndpointTags) {
	if !from.Tags.IsNull() && !from.Tags.IsUnknown() && to.Tags.IsNull() && len(from.Tags.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Tags, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Tags = from.Tags
	}
}

func (c EndpointTags) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a EndpointTags) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tags": reflect.TypeOf(EndpointTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, EndpointTags
// only implements ToObjectValue() and Type().
func (o EndpointTags) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"tags": o.Tags,
		})
}

// Type implements basetypes.ObjectValuable.
func (o EndpointTags) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
		},
	}
}

// GetTags returns the value of the Tags field in EndpointTags as
// a slice of EndpointTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *EndpointTags) GetTags(ctx context.Context) ([]EndpointTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in EndpointTags.
func (o *EndpointTags) SetTags(ctx context.Context, v []EndpointTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (to *ExportMetricsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportMetricsRequest) {
}

func (to *ExportMetricsRequest) SyncFieldsDuringRead(ctx context.Context, from ExportMetricsRequest) {
}

func (c ExportMetricsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExportMetricsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportMetricsRequest
// only implements ToObjectValue() and Type().
func (o ExportMetricsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportMetricsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ExportMetricsResponse struct {
	Contents types.Object `tfsdk:"-"`
}

func (to *ExportMetricsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExportMetricsResponse) {
}

func (to *ExportMetricsResponse) SyncFieldsDuringRead(ctx context.Context, from ExportMetricsResponse) {
}

func (c ExportMetricsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExportMetricsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExportMetricsResponse
// only implements ToObjectValue() and Type().
func (o ExportMetricsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": o.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExportMetricsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

// Simple Proto message for testing
type ExternalFunctionRequest struct {
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

func (to *ExternalFunctionRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalFunctionRequest) {
}

func (to *ExternalFunctionRequest) SyncFieldsDuringRead(ctx context.Context, from ExternalFunctionRequest) {
}

func (c ExternalFunctionRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExternalFunctionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalFunctionRequest
// only implements ToObjectValue() and Type().
func (o ExternalFunctionRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"connection_name": o.ConnectionName,
			"headers":         o.Headers,
			"json":            o.Json,
			"method":          o.Method,
			"params":          o.Params,
			"path":            o.Path,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalFunctionRequest) Type(ctx context.Context) attr.Type {
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

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig types.Object `tfsdk:"ai21labs_config"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig types.Object `tfsdk:"amazon_bedrock_config"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig types.Object `tfsdk:"anthropic_config"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig types.Object `tfsdk:"cohere_config"`
	// Custom Provider Config. Only required if the provider is 'custom'.
	CustomProviderConfig types.Object `tfsdk:"custom_provider_config"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig types.Object `tfsdk:"databricks_model_serving_config"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig types.Object `tfsdk:"google_cloud_vertex_ai_config"`
	// The name of the external model.
	Name types.String `tfsdk:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig types.Object `tfsdk:"openai_config"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig types.Object `tfsdk:"palm_config"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	Provider types.String `tfsdk:"provider"`
	// The task type of the external model.
	Task types.String `tfsdk:"task"`
}

func (to *ExternalModel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalModel) {
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

func (to *ExternalModel) SyncFieldsDuringRead(ctx context.Context, from ExternalModel) {
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

func (c ExternalModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai21labs_config"] = attrs["ai21labs_config"].SetOptional()
	attrs["amazon_bedrock_config"] = attrs["amazon_bedrock_config"].SetOptional()
	attrs["anthropic_config"] = attrs["anthropic_config"].SetOptional()
	attrs["cohere_config"] = attrs["cohere_config"].SetOptional()
	attrs["custom_provider_config"] = attrs["custom_provider_config"].SetOptional()
	attrs["databricks_model_serving_config"] = attrs["databricks_model_serving_config"].SetOptional()
	attrs["google_cloud_vertex_ai_config"] = attrs["google_cloud_vertex_ai_config"].SetOptional()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["openai_config"] = attrs["openai_config"].SetOptional()
	attrs["palm_config"] = attrs["palm_config"].SetOptional()
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
func (a ExternalModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai21labs_config":                 reflect.TypeOf(Ai21LabsConfig{}),
		"amazon_bedrock_config":           reflect.TypeOf(AmazonBedrockConfig{}),
		"anthropic_config":                reflect.TypeOf(AnthropicConfig{}),
		"cohere_config":                   reflect.TypeOf(CohereConfig{}),
		"custom_provider_config":          reflect.TypeOf(CustomProviderConfig{}),
		"databricks_model_serving_config": reflect.TypeOf(DatabricksModelServingConfig{}),
		"google_cloud_vertex_ai_config":   reflect.TypeOf(GoogleCloudVertexAiConfig{}),
		"openai_config":                   reflect.TypeOf(OpenAiConfig{}),
		"palm_config":                     reflect.TypeOf(PaLmConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalModel
// only implements ToObjectValue() and Type().
func (o ExternalModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai21labs_config":                 o.Ai21labsConfig,
			"amazon_bedrock_config":           o.AmazonBedrockConfig,
			"anthropic_config":                o.AnthropicConfig,
			"cohere_config":                   o.CohereConfig,
			"custom_provider_config":          o.CustomProviderConfig,
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
func (o ExternalModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_config":                 Ai21LabsConfig{}.Type(ctx),
			"amazon_bedrock_config":           AmazonBedrockConfig{}.Type(ctx),
			"anthropic_config":                AnthropicConfig{}.Type(ctx),
			"cohere_config":                   CohereConfig{}.Type(ctx),
			"custom_provider_config":          CustomProviderConfig{}.Type(ctx),
			"databricks_model_serving_config": DatabricksModelServingConfig{}.Type(ctx),
			"google_cloud_vertex_ai_config":   GoogleCloudVertexAiConfig{}.Type(ctx),
			"name":                            types.StringType,
			"openai_config":                   OpenAiConfig{}.Type(ctx),
			"palm_config":                     PaLmConfig{}.Type(ctx),
			"provider":                        types.StringType,
			"task":                            types.StringType,
		},
	}
}

// GetAi21labsConfig returns the value of the Ai21labsConfig field in ExternalModel as
// a Ai21LabsConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetAi21labsConfig(ctx context.Context) (Ai21LabsConfig, bool) {
	var e Ai21LabsConfig
	if o.Ai21labsConfig.IsNull() || o.Ai21labsConfig.IsUnknown() {
		return e, false
	}
	var v Ai21LabsConfig
	d := o.Ai21labsConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAi21labsConfig sets the value of the Ai21labsConfig field in ExternalModel.
func (o *ExternalModel) SetAi21labsConfig(ctx context.Context, v Ai21LabsConfig) {
	vs := v.ToObjectValue(ctx)
	o.Ai21labsConfig = vs
}

// GetAmazonBedrockConfig returns the value of the AmazonBedrockConfig field in ExternalModel as
// a AmazonBedrockConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetAmazonBedrockConfig(ctx context.Context) (AmazonBedrockConfig, bool) {
	var e AmazonBedrockConfig
	if o.AmazonBedrockConfig.IsNull() || o.AmazonBedrockConfig.IsUnknown() {
		return e, false
	}
	var v AmazonBedrockConfig
	d := o.AmazonBedrockConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAmazonBedrockConfig sets the value of the AmazonBedrockConfig field in ExternalModel.
func (o *ExternalModel) SetAmazonBedrockConfig(ctx context.Context, v AmazonBedrockConfig) {
	vs := v.ToObjectValue(ctx)
	o.AmazonBedrockConfig = vs
}

// GetAnthropicConfig returns the value of the AnthropicConfig field in ExternalModel as
// a AnthropicConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetAnthropicConfig(ctx context.Context) (AnthropicConfig, bool) {
	var e AnthropicConfig
	if o.AnthropicConfig.IsNull() || o.AnthropicConfig.IsUnknown() {
		return e, false
	}
	var v AnthropicConfig
	d := o.AnthropicConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAnthropicConfig sets the value of the AnthropicConfig field in ExternalModel.
func (o *ExternalModel) SetAnthropicConfig(ctx context.Context, v AnthropicConfig) {
	vs := v.ToObjectValue(ctx)
	o.AnthropicConfig = vs
}

// GetCohereConfig returns the value of the CohereConfig field in ExternalModel as
// a CohereConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetCohereConfig(ctx context.Context) (CohereConfig, bool) {
	var e CohereConfig
	if o.CohereConfig.IsNull() || o.CohereConfig.IsUnknown() {
		return e, false
	}
	var v CohereConfig
	d := o.CohereConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCohereConfig sets the value of the CohereConfig field in ExternalModel.
func (o *ExternalModel) SetCohereConfig(ctx context.Context, v CohereConfig) {
	vs := v.ToObjectValue(ctx)
	o.CohereConfig = vs
}

// GetCustomProviderConfig returns the value of the CustomProviderConfig field in ExternalModel as
// a CustomProviderConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetCustomProviderConfig(ctx context.Context) (CustomProviderConfig, bool) {
	var e CustomProviderConfig
	if o.CustomProviderConfig.IsNull() || o.CustomProviderConfig.IsUnknown() {
		return e, false
	}
	var v CustomProviderConfig
	d := o.CustomProviderConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetCustomProviderConfig sets the value of the CustomProviderConfig field in ExternalModel.
func (o *ExternalModel) SetCustomProviderConfig(ctx context.Context, v CustomProviderConfig) {
	vs := v.ToObjectValue(ctx)
	o.CustomProviderConfig = vs
}

// GetDatabricksModelServingConfig returns the value of the DatabricksModelServingConfig field in ExternalModel as
// a DatabricksModelServingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetDatabricksModelServingConfig(ctx context.Context) (DatabricksModelServingConfig, bool) {
	var e DatabricksModelServingConfig
	if o.DatabricksModelServingConfig.IsNull() || o.DatabricksModelServingConfig.IsUnknown() {
		return e, false
	}
	var v DatabricksModelServingConfig
	d := o.DatabricksModelServingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDatabricksModelServingConfig sets the value of the DatabricksModelServingConfig field in ExternalModel.
func (o *ExternalModel) SetDatabricksModelServingConfig(ctx context.Context, v DatabricksModelServingConfig) {
	vs := v.ToObjectValue(ctx)
	o.DatabricksModelServingConfig = vs
}

// GetGoogleCloudVertexAiConfig returns the value of the GoogleCloudVertexAiConfig field in ExternalModel as
// a GoogleCloudVertexAiConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetGoogleCloudVertexAiConfig(ctx context.Context) (GoogleCloudVertexAiConfig, bool) {
	var e GoogleCloudVertexAiConfig
	if o.GoogleCloudVertexAiConfig.IsNull() || o.GoogleCloudVertexAiConfig.IsUnknown() {
		return e, false
	}
	var v GoogleCloudVertexAiConfig
	d := o.GoogleCloudVertexAiConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGoogleCloudVertexAiConfig sets the value of the GoogleCloudVertexAiConfig field in ExternalModel.
func (o *ExternalModel) SetGoogleCloudVertexAiConfig(ctx context.Context, v GoogleCloudVertexAiConfig) {
	vs := v.ToObjectValue(ctx)
	o.GoogleCloudVertexAiConfig = vs
}

// GetOpenaiConfig returns the value of the OpenaiConfig field in ExternalModel as
// a OpenAiConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetOpenaiConfig(ctx context.Context) (OpenAiConfig, bool) {
	var e OpenAiConfig
	if o.OpenaiConfig.IsNull() || o.OpenaiConfig.IsUnknown() {
		return e, false
	}
	var v OpenAiConfig
	d := o.OpenaiConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetOpenaiConfig sets the value of the OpenaiConfig field in ExternalModel.
func (o *ExternalModel) SetOpenaiConfig(ctx context.Context, v OpenAiConfig) {
	vs := v.ToObjectValue(ctx)
	o.OpenaiConfig = vs
}

// GetPalmConfig returns the value of the PalmConfig field in ExternalModel as
// a PaLmConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ExternalModel) GetPalmConfig(ctx context.Context) (PaLmConfig, bool) {
	var e PaLmConfig
	if o.PalmConfig.IsNull() || o.PalmConfig.IsUnknown() {
		return e, false
	}
	var v PaLmConfig
	d := o.PalmConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPalmConfig sets the value of the PalmConfig field in ExternalModel.
func (o *ExternalModel) SetPalmConfig(ctx context.Context, v PaLmConfig) {
	vs := v.ToObjectValue(ctx)
	o.PalmConfig = vs
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens types.Int64 `tfsdk:"completion_tokens"`
	// The number of tokens in the prompt.
	PromptTokens types.Int64 `tfsdk:"prompt_tokens"`
	// The total number of tokens in the prompt and response.
	TotalTokens types.Int64 `tfsdk:"total_tokens"`
}

func (to *ExternalModelUsageElement) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ExternalModelUsageElement) {
}

func (to *ExternalModelUsageElement) SyncFieldsDuringRead(ctx context.Context, from ExternalModelUsageElement) {
}

func (c ExternalModelUsageElement) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ExternalModelUsageElement) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ExternalModelUsageElement
// only implements ToObjectValue() and Type().
func (o ExternalModelUsageElement) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"completion_tokens": o.CompletionTokens,
			"prompt_tokens":     o.PromptTokens,
			"total_tokens":      o.TotalTokens,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ExternalModelUsageElement) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"completion_tokens": types.Int64Type,
			"prompt_tokens":     types.Int64Type,
			"total_tokens":      types.Int64Type,
		},
	}
}

type FallbackConfig struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	Enabled types.Bool `tfsdk:"enabled"`
}

func (to *FallbackConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FallbackConfig) {
}

func (to *FallbackConfig) SyncFieldsDuringRead(ctx context.Context, from FallbackConfig) {
}

func (c FallbackConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FallbackConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FallbackConfig
// only implements ToObjectValue() and Type().
func (o FallbackConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"enabled": o.Enabled,
		})
}

// Type implements basetypes.ObjectValuable.
func (o FallbackConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

// All fields are not sensitive as they are hard-coded in the system and made
// available to customers.
type FoundationModel struct {
	Description types.String `tfsdk:"description"`

	DisplayName types.String `tfsdk:"display_name"`

	Docs types.String `tfsdk:"docs"`

	Name types.String `tfsdk:"name"`
}

func (to *FoundationModel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from FoundationModel) {
}

func (to *FoundationModel) SyncFieldsDuringRead(ctx context.Context, from FoundationModel) {
}

func (c FoundationModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a FoundationModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FoundationModel
// only implements ToObjectValue() and Type().
func (o FoundationModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o FoundationModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":  types.StringType,
			"display_name": types.StringType,
			"docs":         types.StringType,
			"name":         types.StringType,
		},
	}
}

type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
}

func (to *GetOpenApiRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOpenApiRequest) {
}

func (to *GetOpenApiRequest) SyncFieldsDuringRead(ctx context.Context, from GetOpenApiRequest) {
}

func (c GetOpenApiRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetOpenApiRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOpenApiRequest
// only implements ToObjectValue() and Type().
func (o GetOpenApiRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOpenApiRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GetOpenApiResponse struct {
	Contents types.Object `tfsdk:"-"`
}

func (to *GetOpenApiResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetOpenApiResponse) {
}

func (to *GetOpenApiResponse) SyncFieldsDuringRead(ctx context.Context, from GetOpenApiResponse) {
}

func (c GetOpenApiResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetOpenApiResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetOpenApiResponse
// only implements ToObjectValue() and Type().
func (o GetOpenApiResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": o.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetOpenApiResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (to *GetServingEndpointPermissionLevelsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointPermissionLevelsRequest) {
}

func (to *GetServingEndpointPermissionLevelsRequest) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointPermissionLevelsRequest) {
}

func (c GetServingEndpointPermissionLevelsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetServingEndpointPermissionLevelsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionLevelsRequest
// only implements ToObjectValue() and Type().
func (o GetServingEndpointPermissionLevelsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"serving_endpoint_id": o.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointPermissionLevelsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

func (to *GetServingEndpointPermissionLevelsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (to *GetServingEndpointPermissionLevelsResponse) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointPermissionLevelsResponse) {
	if !from.PermissionLevels.IsNull() && !from.PermissionLevels.IsUnknown() && to.PermissionLevels.IsNull() && len(from.PermissionLevels.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for PermissionLevels, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.PermissionLevels = from.PermissionLevels
	}
}

func (c GetServingEndpointPermissionLevelsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetServingEndpointPermissionLevelsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"permission_levels": reflect.TypeOf(ServingEndpointPermissionsDescription{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionLevelsResponse
// only implements ToObjectValue() and Type().
func (o GetServingEndpointPermissionLevelsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"permission_levels": o.PermissionLevels,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointPermissionLevelsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ServingEndpointPermissionsDescription{}.Type(ctx),
			},
		},
	}
}

// GetPermissionLevels returns the value of the PermissionLevels field in GetServingEndpointPermissionLevelsResponse as
// a slice of ServingEndpointPermissionsDescription values.
// If the field is unknown or null, the boolean return value is false.
func (o *GetServingEndpointPermissionLevelsResponse) GetPermissionLevels(ctx context.Context) ([]ServingEndpointPermissionsDescription, bool) {
	if o.PermissionLevels.IsNull() || o.PermissionLevels.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointPermissionsDescription
	d := o.PermissionLevels.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPermissionLevels sets the value of the PermissionLevels field in GetServingEndpointPermissionLevelsResponse.
func (o *GetServingEndpointPermissionLevelsResponse) SetPermissionLevels(ctx context.Context, v []ServingEndpointPermissionsDescription) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["permission_levels"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.PermissionLevels = types.ListValueMust(t, vs)
}

type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (to *GetServingEndpointPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointPermissionsRequest) {
}

func (to *GetServingEndpointPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointPermissionsRequest) {
}

func (c GetServingEndpointPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetServingEndpointPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointPermissionsRequest
// only implements ToObjectValue() and Type().
func (o GetServingEndpointPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"serving_endpoint_id": o.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (to *GetServingEndpointRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GetServingEndpointRequest) {
}

func (to *GetServingEndpointRequest) SyncFieldsDuringRead(ctx context.Context, from GetServingEndpointRequest) {
}

func (c GetServingEndpointRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GetServingEndpointRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GetServingEndpointRequest
// only implements ToObjectValue() and Type().
func (o GetServingEndpointRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name": o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o GetServingEndpointRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type GoogleCloudVertexAiConfig struct {
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

func (to *GoogleCloudVertexAiConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GoogleCloudVertexAiConfig) {
}

func (to *GoogleCloudVertexAiConfig) SyncFieldsDuringRead(ctx context.Context, from GoogleCloudVertexAiConfig) {
}

func (c GoogleCloudVertexAiConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a GoogleCloudVertexAiConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, GoogleCloudVertexAiConfig
// only implements ToObjectValue() and Type().
func (o GoogleCloudVertexAiConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o GoogleCloudVertexAiConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_key":           types.StringType,
			"private_key_plaintext": types.StringType,
			"project_id":            types.StringType,
			"region":                types.StringType,
		},
	}
}

type HttpRequestResponse struct {
	Contents types.Object `tfsdk:"-"`
}

func (to *HttpRequestResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from HttpRequestResponse) {
}

func (to *HttpRequestResponse) SyncFieldsDuringRead(ctx context.Context, from HttpRequestResponse) {
}

func (c HttpRequestResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a HttpRequestResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, HttpRequestResponse
// only implements ToObjectValue() and Type().
func (o HttpRequestResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"contents": o.Contents,
		})
}

// Type implements basetypes.ObjectValuable.
func (o HttpRequestResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints types.List `tfsdk:"endpoints"`
}

func (to *ListEndpointsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListEndpointsResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (to *ListEndpointsResponse) SyncFieldsDuringRead(ctx context.Context, from ListEndpointsResponse) {
	if !from.Endpoints.IsNull() && !from.Endpoints.IsUnknown() && to.Endpoints.IsNull() && len(from.Endpoints.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Endpoints, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Endpoints = from.Endpoints
	}
}

func (c ListEndpointsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ListEndpointsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"endpoints": reflect.TypeOf(ServingEndpoint{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListEndpointsResponse
// only implements ToObjectValue() and Type().
func (o ListEndpointsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"endpoints": o.Endpoints,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ListEndpointsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: ServingEndpoint{}.Type(ctx),
			},
		},
	}
}

// GetEndpoints returns the value of the Endpoints field in ListEndpointsResponse as
// a slice of ServingEndpoint values.
// If the field is unknown or null, the boolean return value is false.
func (o *ListEndpointsResponse) GetEndpoints(ctx context.Context) ([]ServingEndpoint, bool) {
	if o.Endpoints.IsNull() || o.Endpoints.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpoint
	d := o.Endpoints.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEndpoints sets the value of the Endpoints field in ListEndpointsResponse.
func (o *ListEndpointsResponse) SetEndpoints(ctx context.Context, v []ServingEndpoint) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["endpoints"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Endpoints = types.ListValueMust(t, vs)
}

type ListServingEndpointsRequest struct {
}

func (to *ListServingEndpointsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ListServingEndpointsRequest) {
}

func (to *ListServingEndpointsRequest) SyncFieldsDuringRead(ctx context.Context, from ListServingEndpointsRequest) {
}

func (c ListServingEndpointsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListServingEndpointsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListServingEndpointsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ListServingEndpointsRequest
// only implements ToObjectValue() and Type().
func (o ListServingEndpointsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{})
}

// Type implements basetypes.ObjectValuable.
func (o ListServingEndpointsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (to *LogsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from LogsRequest) {
}

func (to *LogsRequest) SyncFieldsDuringRead(ctx context.Context, from LogsRequest) {
}

func (c LogsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a LogsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, LogsRequest
// only implements ToObjectValue() and Type().
func (o LogsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":              o.Name,
			"served_model_name": o.ServedModelName,
		})
}

// Type implements basetypes.ObjectValuable.
func (o LogsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

// A representation of all DataPlaneInfo for operations that can be done on a
// model through Data Plane APIs.
type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo types.Object `tfsdk:"query_info"`
}

func (to *ModelDataPlaneInfo) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ModelDataPlaneInfo) {
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

func (to *ModelDataPlaneInfo) SyncFieldsDuringRead(ctx context.Context, from ModelDataPlaneInfo) {
	if !from.QueryInfo.IsNull() && !from.QueryInfo.IsUnknown() {
		if toQueryInfo, ok := to.GetQueryInfo(ctx); ok {
			if fromQueryInfo, ok := from.GetQueryInfo(ctx); ok {
				toQueryInfo.SyncFieldsDuringRead(ctx, fromQueryInfo)
				to.SetQueryInfo(ctx, toQueryInfo)
			}
		}
	}
}

func (c ModelDataPlaneInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["query_info"] = attrs["query_info"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ModelDataPlaneInfo.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ModelDataPlaneInfo) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query_info": reflect.TypeOf(DataPlaneInfo{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ModelDataPlaneInfo
// only implements ToObjectValue() and Type().
func (o ModelDataPlaneInfo) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"query_info": o.QueryInfo,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ModelDataPlaneInfo) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_info": DataPlaneInfo{}.Type(ctx),
		},
	}
}

// GetQueryInfo returns the value of the QueryInfo field in ModelDataPlaneInfo as
// a DataPlaneInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ModelDataPlaneInfo) GetQueryInfo(ctx context.Context) (DataPlaneInfo, bool) {
	var e DataPlaneInfo
	if o.QueryInfo.IsNull() || o.QueryInfo.IsUnknown() {
		return e, false
	}
	var v DataPlaneInfo
	d := o.QueryInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetQueryInfo sets the value of the QueryInfo field in ModelDataPlaneInfo.
func (o *ModelDataPlaneInfo) SetQueryInfo(ctx context.Context, v DataPlaneInfo) {
	vs := v.ToObjectValue(ctx)
	o.QueryInfo = vs
}

// Configs needed to create an OpenAI model route.
type OpenAiConfig struct {
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

func (to *OpenAiConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from OpenAiConfig) {
}

func (to *OpenAiConfig) SyncFieldsDuringRead(ctx context.Context, from OpenAiConfig) {
}

func (c OpenAiConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a OpenAiConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, OpenAiConfig
// only implements ToObjectValue() and Type().
func (o OpenAiConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o OpenAiConfig) Type(ctx context.Context) attr.Type {
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

type PaLmConfig struct {
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

func (to *PaLmConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PaLmConfig) {
}

func (to *PaLmConfig) SyncFieldsDuringRead(ctx context.Context, from PaLmConfig) {
}

func (c PaLmConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PaLmConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PaLmConfig
// only implements ToObjectValue() and Type().
func (o PaLmConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"palm_api_key":           o.PalmApiKey,
			"palm_api_key_plaintext": o.PalmApiKeyPlaintext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PaLmConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"palm_api_key":           types.StringType,
			"palm_api_key_plaintext": types.StringType,
		},
	}
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags types.List `tfsdk:"add_tags"`
	// List of tag keys to delete
	DeleteTags types.List `tfsdk:"delete_tags"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (to *PatchServingEndpointTags) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PatchServingEndpointTags) {
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

func (to *PatchServingEndpointTags) SyncFieldsDuringRead(ctx context.Context, from PatchServingEndpointTags) {
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

func (c PatchServingEndpointTags) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PatchServingEndpointTags) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"add_tags":    reflect.TypeOf(EndpointTag{}),
		"delete_tags": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PatchServingEndpointTags
// only implements ToObjectValue() and Type().
func (o PatchServingEndpointTags) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"add_tags":    o.AddTags,
			"delete_tags": o.DeleteTags,
			"name":        o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PatchServingEndpointTags) Type(ctx context.Context) attr.Type {
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

// GetAddTags returns the value of the AddTags field in PatchServingEndpointTags as
// a slice of EndpointTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchServingEndpointTags) GetAddTags(ctx context.Context) ([]EndpointTag, bool) {
	if o.AddTags.IsNull() || o.AddTags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag
	d := o.AddTags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAddTags sets the value of the AddTags field in PatchServingEndpointTags.
func (o *PatchServingEndpointTags) SetAddTags(ctx context.Context, v []EndpointTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["add_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AddTags = types.ListValueMust(t, vs)
}

// GetDeleteTags returns the value of the DeleteTags field in PatchServingEndpointTags as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *PatchServingEndpointTags) GetDeleteTags(ctx context.Context) ([]types.String, bool) {
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

// SetDeleteTags sets the value of the DeleteTags field in PatchServingEndpointTags.
func (o *PatchServingEndpointTags) SetDeleteTags(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["delete_tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DeleteTags = types.ListValueMust(t, vs)
}

type PayloadTable struct {
	Name types.String `tfsdk:"name"`

	Status types.String `tfsdk:"status"`

	StatusMessage types.String `tfsdk:"status_message"`
}

func (to *PayloadTable) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PayloadTable) {
}

func (to *PayloadTable) SyncFieldsDuringRead(ctx context.Context, from PayloadTable) {
}

func (c PayloadTable) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PayloadTable) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PayloadTable
// only implements ToObjectValue() and Type().
func (o PayloadTable) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":           o.Name,
			"status":         o.Status,
			"status_message": o.StatusMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PayloadTable) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":           types.StringType,
			"status":         types.StringType,
			"status_message": types.StringType,
		},
	}
}

type PtEndpointCoreConfig struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities"`

	TrafficConfig types.Object `tfsdk:"traffic_config"`
}

func (to *PtEndpointCoreConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PtEndpointCoreConfig) {
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

func (to *PtEndpointCoreConfig) SyncFieldsDuringRead(ctx context.Context, from PtEndpointCoreConfig) {
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

func (c PtEndpointCoreConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["served_entities"] = attrs["served_entities"].SetOptional()
	attrs["traffic_config"] = attrs["traffic_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PtEndpointCoreConfig.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PtEndpointCoreConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"served_entities": reflect.TypeOf(PtServedModel{}),
		"traffic_config":  reflect.TypeOf(TrafficConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PtEndpointCoreConfig
// only implements ToObjectValue() and Type().
func (o PtEndpointCoreConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entities": o.ServedEntities,
			"traffic_config":  o.TrafficConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PtEndpointCoreConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entities": basetypes.ListType{
				ElemType: PtServedModel{}.Type(ctx),
			},
			"traffic_config": TrafficConfig{}.Type(ctx),
		},
	}
}

// GetServedEntities returns the value of the ServedEntities field in PtEndpointCoreConfig as
// a slice of PtServedModel values.
// If the field is unknown or null, the boolean return value is false.
func (o *PtEndpointCoreConfig) GetServedEntities(ctx context.Context) ([]PtServedModel, bool) {
	if o.ServedEntities.IsNull() || o.ServedEntities.IsUnknown() {
		return nil, false
	}
	var v []PtServedModel
	d := o.ServedEntities.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetServedEntities sets the value of the ServedEntities field in PtEndpointCoreConfig.
func (o *PtEndpointCoreConfig) SetServedEntities(ctx context.Context, v []PtServedModel) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["served_entities"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ServedEntities = types.ListValueMust(t, vs)
}

// GetTrafficConfig returns the value of the TrafficConfig field in PtEndpointCoreConfig as
// a TrafficConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PtEndpointCoreConfig) GetTrafficConfig(ctx context.Context) (TrafficConfig, bool) {
	var e TrafficConfig
	if o.TrafficConfig.IsNull() || o.TrafficConfig.IsUnknown() {
		return e, false
	}
	var v TrafficConfig
	d := o.TrafficConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTrafficConfig sets the value of the TrafficConfig field in PtEndpointCoreConfig.
func (o *PtEndpointCoreConfig) SetTrafficConfig(ctx context.Context, v TrafficConfig) {
	vs := v.ToObjectValue(ctx)
	o.TrafficConfig = vs
}

type PtServedModel struct {
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

func (to *PtServedModel) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PtServedModel) {
}

func (to *PtServedModel) SyncFieldsDuringRead(ctx context.Context, from PtServedModel) {
}

func (c PtServedModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PtServedModel) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PtServedModel
// only implements ToObjectValue() and Type().
func (o PtServedModel) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":             o.EntityName,
			"entity_version":          o.EntityVersion,
			"name":                    o.Name,
			"provisioned_model_units": o.ProvisionedModelUnits,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PtServedModel) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":             types.StringType,
			"entity_version":          types.StringType,
			"name":                    types.StringType,
			"provisioned_model_units": types.Int64Type,
		},
	}
}

type PutAiGatewayRequest struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig types.Object `tfsdk:"fallback_config"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.Object `tfsdk:"guardrails"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.Object `tfsdk:"inference_table_config"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.Object `tfsdk:"usage_tracking_config"`
}

func (to *PutAiGatewayRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutAiGatewayRequest) {
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

func (to *PutAiGatewayRequest) SyncFieldsDuringRead(ctx context.Context, from PutAiGatewayRequest) {
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

func (c PutAiGatewayRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fallback_config"] = attrs["fallback_config"].SetOptional()
	attrs["guardrails"] = attrs["guardrails"].SetOptional()
	attrs["inference_table_config"] = attrs["inference_table_config"].SetOptional()
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].SetOptional()
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
func (a PutAiGatewayRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fallback_config":        reflect.TypeOf(FallbackConfig{}),
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAiGatewayRequest
// only implements ToObjectValue() and Type().
func (o PutAiGatewayRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fallback_config":        o.FallbackConfig,
			"guardrails":             o.Guardrails,
			"inference_table_config": o.InferenceTableConfig,
			"name":                   o.Name,
			"rate_limits":            o.RateLimits,
			"usage_tracking_config":  o.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutAiGatewayRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fallback_config":        FallbackConfig{}.Type(ctx),
			"guardrails":             AiGatewayGuardrails{}.Type(ctx),
			"inference_table_config": AiGatewayInferenceTableConfig{}.Type(ctx),
			"name":                   types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.Type(ctx),
			},
			"usage_tracking_config": AiGatewayUsageTrackingConfig{}.Type(ctx),
		},
	}
}

// GetFallbackConfig returns the value of the FallbackConfig field in PutAiGatewayRequest as
// a FallbackConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest) GetFallbackConfig(ctx context.Context) (FallbackConfig, bool) {
	var e FallbackConfig
	if o.FallbackConfig.IsNull() || o.FallbackConfig.IsUnknown() {
		return e, false
	}
	var v FallbackConfig
	d := o.FallbackConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFallbackConfig sets the value of the FallbackConfig field in PutAiGatewayRequest.
func (o *PutAiGatewayRequest) SetFallbackConfig(ctx context.Context, v FallbackConfig) {
	vs := v.ToObjectValue(ctx)
	o.FallbackConfig = vs
}

// GetGuardrails returns the value of the Guardrails field in PutAiGatewayRequest as
// a AiGatewayGuardrails value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest) GetGuardrails(ctx context.Context) (AiGatewayGuardrails, bool) {
	var e AiGatewayGuardrails
	if o.Guardrails.IsNull() || o.Guardrails.IsUnknown() {
		return e, false
	}
	var v AiGatewayGuardrails
	d := o.Guardrails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGuardrails sets the value of the Guardrails field in PutAiGatewayRequest.
func (o *PutAiGatewayRequest) SetGuardrails(ctx context.Context, v AiGatewayGuardrails) {
	vs := v.ToObjectValue(ctx)
	o.Guardrails = vs
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in PutAiGatewayRequest as
// a AiGatewayInferenceTableConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig, bool) {
	var e AiGatewayInferenceTableConfig
	if o.InferenceTableConfig.IsNull() || o.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v AiGatewayInferenceTableConfig
	d := o.InferenceTableConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in PutAiGatewayRequest.
func (o *PutAiGatewayRequest) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig) {
	vs := v.ToObjectValue(ctx)
	o.InferenceTableConfig = vs
}

// GetRateLimits returns the value of the RateLimits field in PutAiGatewayRequest as
// a slice of AiGatewayRateLimit values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutAiGatewayRequest.
func (o *PutAiGatewayRequest) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in PutAiGatewayRequest as
// a AiGatewayUsageTrackingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayRequest) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig, bool) {
	var e AiGatewayUsageTrackingConfig
	if o.UsageTrackingConfig.IsNull() || o.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v AiGatewayUsageTrackingConfig
	d := o.UsageTrackingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in PutAiGatewayRequest.
func (o *PutAiGatewayRequest) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig) {
	vs := v.ToObjectValue(ctx)
	o.UsageTrackingConfig = vs
}

type PutAiGatewayResponse struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig types.Object `tfsdk:"fallback_config"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails types.Object `tfsdk:"guardrails"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig types.Object `tfsdk:"inference_table_config"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits types.List `tfsdk:"rate_limits"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig types.Object `tfsdk:"usage_tracking_config"`
}

func (to *PutAiGatewayResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutAiGatewayResponse) {
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

func (to *PutAiGatewayResponse) SyncFieldsDuringRead(ctx context.Context, from PutAiGatewayResponse) {
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

func (c PutAiGatewayResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["fallback_config"] = attrs["fallback_config"].SetOptional()
	attrs["guardrails"] = attrs["guardrails"].SetOptional()
	attrs["inference_table_config"] = attrs["inference_table_config"].SetOptional()
	attrs["rate_limits"] = attrs["rate_limits"].SetOptional()
	attrs["usage_tracking_config"] = attrs["usage_tracking_config"].SetOptional()

	return attrs
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PutAiGatewayResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PutAiGatewayResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"fallback_config":        reflect.TypeOf(FallbackConfig{}),
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutAiGatewayResponse
// only implements ToObjectValue() and Type().
func (o PutAiGatewayResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"fallback_config":        o.FallbackConfig,
			"guardrails":             o.Guardrails,
			"inference_table_config": o.InferenceTableConfig,
			"rate_limits":            o.RateLimits,
			"usage_tracking_config":  o.UsageTrackingConfig,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutAiGatewayResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"fallback_config":        FallbackConfig{}.Type(ctx),
			"guardrails":             AiGatewayGuardrails{}.Type(ctx),
			"inference_table_config": AiGatewayInferenceTableConfig{}.Type(ctx),
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.Type(ctx),
			},
			"usage_tracking_config": AiGatewayUsageTrackingConfig{}.Type(ctx),
		},
	}
}

// GetFallbackConfig returns the value of the FallbackConfig field in PutAiGatewayResponse as
// a FallbackConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse) GetFallbackConfig(ctx context.Context) (FallbackConfig, bool) {
	var e FallbackConfig
	if o.FallbackConfig.IsNull() || o.FallbackConfig.IsUnknown() {
		return e, false
	}
	var v FallbackConfig
	d := o.FallbackConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFallbackConfig sets the value of the FallbackConfig field in PutAiGatewayResponse.
func (o *PutAiGatewayResponse) SetFallbackConfig(ctx context.Context, v FallbackConfig) {
	vs := v.ToObjectValue(ctx)
	o.FallbackConfig = vs
}

// GetGuardrails returns the value of the Guardrails field in PutAiGatewayResponse as
// a AiGatewayGuardrails value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse) GetGuardrails(ctx context.Context) (AiGatewayGuardrails, bool) {
	var e AiGatewayGuardrails
	if o.Guardrails.IsNull() || o.Guardrails.IsUnknown() {
		return e, false
	}
	var v AiGatewayGuardrails
	d := o.Guardrails.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetGuardrails sets the value of the Guardrails field in PutAiGatewayResponse.
func (o *PutAiGatewayResponse) SetGuardrails(ctx context.Context, v AiGatewayGuardrails) {
	vs := v.ToObjectValue(ctx)
	o.Guardrails = vs
}

// GetInferenceTableConfig returns the value of the InferenceTableConfig field in PutAiGatewayResponse as
// a AiGatewayInferenceTableConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse) GetInferenceTableConfig(ctx context.Context) (AiGatewayInferenceTableConfig, bool) {
	var e AiGatewayInferenceTableConfig
	if o.InferenceTableConfig.IsNull() || o.InferenceTableConfig.IsUnknown() {
		return e, false
	}
	var v AiGatewayInferenceTableConfig
	d := o.InferenceTableConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetInferenceTableConfig sets the value of the InferenceTableConfig field in PutAiGatewayResponse.
func (o *PutAiGatewayResponse) SetInferenceTableConfig(ctx context.Context, v AiGatewayInferenceTableConfig) {
	vs := v.ToObjectValue(ctx)
	o.InferenceTableConfig = vs
}

// GetRateLimits returns the value of the RateLimits field in PutAiGatewayResponse as
// a slice of AiGatewayRateLimit values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse) GetRateLimits(ctx context.Context) ([]AiGatewayRateLimit, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []AiGatewayRateLimit
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutAiGatewayResponse.
func (o *PutAiGatewayResponse) SetRateLimits(ctx context.Context, v []AiGatewayRateLimit) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

// GetUsageTrackingConfig returns the value of the UsageTrackingConfig field in PutAiGatewayResponse as
// a AiGatewayUsageTrackingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *PutAiGatewayResponse) GetUsageTrackingConfig(ctx context.Context) (AiGatewayUsageTrackingConfig, bool) {
	var e AiGatewayUsageTrackingConfig
	if o.UsageTrackingConfig.IsNull() || o.UsageTrackingConfig.IsUnknown() {
		return e, false
	}
	var v AiGatewayUsageTrackingConfig
	d := o.UsageTrackingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsageTrackingConfig sets the value of the UsageTrackingConfig field in PutAiGatewayResponse.
func (o *PutAiGatewayResponse) SetUsageTrackingConfig(ctx context.Context, v AiGatewayUsageTrackingConfig) {
	vs := v.ToObjectValue(ctx)
	o.UsageTrackingConfig = vs
}

type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits"`
}

func (to *PutRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutRequest) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (to *PutRequest) SyncFieldsDuringRead(ctx context.Context, from PutRequest) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (c PutRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PutRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rate_limits": reflect.TypeOf(RateLimit{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutRequest
// only implements ToObjectValue() and Type().
func (o PutRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"name":        o.Name,
			"rate_limits": o.RateLimits,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.Type(ctx),
			},
		},
	}
}

// GetRateLimits returns the value of the RateLimits field in PutRequest as
// a slice of RateLimit values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutRequest) GetRateLimits(ctx context.Context) ([]RateLimit, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutRequest.
func (o *PutRequest) SetRateLimits(ctx context.Context, v []RateLimit) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits"`
}

func (to *PutResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from PutResponse) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (to *PutResponse) SyncFieldsDuringRead(ctx context.Context, from PutResponse) {
	if !from.RateLimits.IsNull() && !from.RateLimits.IsUnknown() && to.RateLimits.IsNull() && len(from.RateLimits.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for RateLimits, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.RateLimits = from.RateLimits
	}
}

func (c PutResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a PutResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"rate_limits": reflect.TypeOf(RateLimit{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, PutResponse
// only implements ToObjectValue() and Type().
func (o PutResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"rate_limits": o.RateLimits,
		})
}

// Type implements basetypes.ObjectValuable.
func (o PutResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.Type(ctx),
			},
		},
	}
}

// GetRateLimits returns the value of the RateLimits field in PutResponse as
// a slice of RateLimit values.
// If the field is unknown or null, the boolean return value is false.
func (o *PutResponse) GetRateLimits(ctx context.Context) ([]RateLimit, bool) {
	if o.RateLimits.IsNull() || o.RateLimits.IsUnknown() {
		return nil, false
	}
	var v []RateLimit
	d := o.RateLimits.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRateLimits sets the value of the RateLimits field in PutResponse.
func (o *PutResponse) SetRateLimits(ctx context.Context, v []RateLimit) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["rate_limits"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.RateLimits = types.ListValueMust(t, vs)
}

type QueryEndpointInput struct {
	// Optional user-provided request identifier that will be recorded in the
	// inference table and the usage tracking table.
	ClientRequestId types.String `tfsdk:"client_request_id"`
	// Pandas Dataframe input in the records orientation.
	DataframeRecords types.List `tfsdk:"dataframe_records"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit types.Object `tfsdk:"dataframe_split"`
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

func (to *QueryEndpointInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryEndpointInput) {
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

func (to *QueryEndpointInput) SyncFieldsDuringRead(ctx context.Context, from QueryEndpointInput) {
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

func (c QueryEndpointInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["client_request_id"] = attrs["client_request_id"].SetOptional()
	attrs["dataframe_records"] = attrs["dataframe_records"].SetOptional()
	attrs["dataframe_split"] = attrs["dataframe_split"].SetOptional()
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
func (a QueryEndpointInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dataframe_records": reflect.TypeOf(types.Object{}),
		"dataframe_split":   reflect.TypeOf(DataframeSplitInput{}),
		"extra_params":      reflect.TypeOf(types.String{}),
		"instances":         reflect.TypeOf(types.Object{}),
		"messages":          reflect.TypeOf(ChatMessage{}),
		"stop":              reflect.TypeOf(types.String{}),
		"usage_context":     reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEndpointInput
// only implements ToObjectValue() and Type().
func (o QueryEndpointInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"client_request_id": o.ClientRequestId,
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
			"usage_context":     o.UsageContext,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryEndpointInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"client_request_id": types.StringType,
			"dataframe_records": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"dataframe_split": DataframeSplitInput{}.Type(ctx),
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
			"usage_context": basetypes.MapType{
				ElemType: types.StringType,
			},
		},
	}
}

// GetDataframeRecords returns the value of the DataframeRecords field in QueryEndpointInput as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetDataframeRecords(ctx context.Context) ([]types.Object, bool) {
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

// SetDataframeRecords sets the value of the DataframeRecords field in QueryEndpointInput.
func (o *QueryEndpointInput) SetDataframeRecords(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["dataframe_records"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.DataframeRecords = types.ListValueMust(t, vs)
}

// GetDataframeSplit returns the value of the DataframeSplit field in QueryEndpointInput as
// a DataframeSplitInput value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetDataframeSplit(ctx context.Context) (DataframeSplitInput, bool) {
	var e DataframeSplitInput
	if o.DataframeSplit.IsNull() || o.DataframeSplit.IsUnknown() {
		return e, false
	}
	var v DataframeSplitInput
	d := o.DataframeSplit.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataframeSplit sets the value of the DataframeSplit field in QueryEndpointInput.
func (o *QueryEndpointInput) SetDataframeSplit(ctx context.Context, v DataframeSplitInput) {
	vs := v.ToObjectValue(ctx)
	o.DataframeSplit = vs
}

// GetExtraParams returns the value of the ExtraParams field in QueryEndpointInput as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetExtraParams(ctx context.Context) (map[string]types.String, bool) {
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

// SetExtraParams sets the value of the ExtraParams field in QueryEndpointInput.
func (o *QueryEndpointInput) SetExtraParams(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["extra_params"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.ExtraParams = types.MapValueMust(t, vs)
}

// GetInstances returns the value of the Instances field in QueryEndpointInput as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetInstances(ctx context.Context) ([]types.Object, bool) {
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

// SetInstances sets the value of the Instances field in QueryEndpointInput.
func (o *QueryEndpointInput) SetInstances(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["instances"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Instances = types.ListValueMust(t, vs)
}

// GetMessages returns the value of the Messages field in QueryEndpointInput as
// a slice of ChatMessage values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetMessages(ctx context.Context) ([]ChatMessage, bool) {
	if o.Messages.IsNull() || o.Messages.IsUnknown() {
		return nil, false
	}
	var v []ChatMessage
	d := o.Messages.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessages sets the value of the Messages field in QueryEndpointInput.
func (o *QueryEndpointInput) SetMessages(ctx context.Context, v []ChatMessage) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["messages"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Messages = types.ListValueMust(t, vs)
}

// GetStop returns the value of the Stop field in QueryEndpointInput as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetStop(ctx context.Context) ([]types.String, bool) {
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

// SetStop sets the value of the Stop field in QueryEndpointInput.
func (o *QueryEndpointInput) SetStop(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["stop"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Stop = types.ListValueMust(t, vs)
}

// GetUsageContext returns the value of the UsageContext field in QueryEndpointInput as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointInput) GetUsageContext(ctx context.Context) (map[string]types.String, bool) {
	if o.UsageContext.IsNull() || o.UsageContext.IsUnknown() {
		return nil, false
	}
	var v map[string]types.String
	d := o.UsageContext.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsageContext sets the value of the UsageContext field in QueryEndpointInput.
func (o *QueryEndpointInput) SetUsageContext(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["usage_context"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.UsageContext = types.MapValueMust(t, vs)
}

type QueryEndpointResponse struct {
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
	Usage types.Object `tfsdk:"usage"`
}

func (to *QueryEndpointResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from QueryEndpointResponse) {
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

func (to *QueryEndpointResponse) SyncFieldsDuringRead(ctx context.Context, from QueryEndpointResponse) {
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

func (c QueryEndpointResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["choices"] = attrs["choices"].SetOptional()
	attrs["created"] = attrs["created"].SetOptional()
	attrs["data"] = attrs["data"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["model"] = attrs["model"].SetOptional()
	attrs["object"] = attrs["object"].SetOptional()
	attrs["predictions"] = attrs["predictions"].SetOptional()
	attrs["usage"] = attrs["usage"].SetOptional()
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
func (a QueryEndpointResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"choices":     reflect.TypeOf(V1ResponseChoiceElement{}),
		"data":        reflect.TypeOf(EmbeddingsV1ResponseEmbeddingElement{}),
		"predictions": reflect.TypeOf(types.Object{}),
		"usage":       reflect.TypeOf(ExternalModelUsageElement{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QueryEndpointResponse
// only implements ToObjectValue() and Type().
func (o QueryEndpointResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
			"served_model_name": o.ServedModelName,
			"usage":             o.Usage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o QueryEndpointResponse) Type(ctx context.Context) attr.Type {
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
			"served_model_name": types.StringType,
			"usage":             ExternalModelUsageElement{}.Type(ctx),
		},
	}
}

// GetChoices returns the value of the Choices field in QueryEndpointResponse as
// a slice of V1ResponseChoiceElement values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse) GetChoices(ctx context.Context) ([]V1ResponseChoiceElement, bool) {
	if o.Choices.IsNull() || o.Choices.IsUnknown() {
		return nil, false
	}
	var v []V1ResponseChoiceElement
	d := o.Choices.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetChoices sets the value of the Choices field in QueryEndpointResponse.
func (o *QueryEndpointResponse) SetChoices(ctx context.Context, v []V1ResponseChoiceElement) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["choices"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Choices = types.ListValueMust(t, vs)
}

// GetData returns the value of the Data field in QueryEndpointResponse as
// a slice of EmbeddingsV1ResponseEmbeddingElement values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse) GetData(ctx context.Context) ([]EmbeddingsV1ResponseEmbeddingElement, bool) {
	if o.Data.IsNull() || o.Data.IsUnknown() {
		return nil, false
	}
	var v []EmbeddingsV1ResponseEmbeddingElement
	d := o.Data.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetData sets the value of the Data field in QueryEndpointResponse.
func (o *QueryEndpointResponse) SetData(ctx context.Context, v []EmbeddingsV1ResponseEmbeddingElement) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["data"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Data = types.ListValueMust(t, vs)
}

// GetPredictions returns the value of the Predictions field in QueryEndpointResponse as
// a slice of types.Object values.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse) GetPredictions(ctx context.Context) ([]types.Object, bool) {
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

// SetPredictions sets the value of the Predictions field in QueryEndpointResponse.
func (o *QueryEndpointResponse) SetPredictions(ctx context.Context, v []types.Object) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["predictions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Predictions = types.ListValueMust(t, vs)
}

// GetUsage returns the value of the Usage field in QueryEndpointResponse as
// a ExternalModelUsageElement value.
// If the field is unknown or null, the boolean return value is false.
func (o *QueryEndpointResponse) GetUsage(ctx context.Context) (ExternalModelUsageElement, bool) {
	var e ExternalModelUsageElement
	if o.Usage.IsNull() || o.Usage.IsUnknown() {
		return e, false
	}
	var v ExternalModelUsageElement
	d := o.Usage.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetUsage sets the value of the Usage field in QueryEndpointResponse.
func (o *QueryEndpointResponse) SetUsage(ctx context.Context, v ExternalModelUsageElement) {
	vs := v.ToObjectValue(ctx)
	o.Usage = vs
}

type RateLimit struct {
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

func (to *RateLimit) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from RateLimit) {
}

func (to *RateLimit) SyncFieldsDuringRead(ctx context.Context, from RateLimit) {
}

func (c RateLimit) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a RateLimit) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, RateLimit
// only implements ToObjectValue() and Type().
func (o RateLimit) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"calls":          o.Calls,
			"key":            o.Key,
			"renewal_period": o.RenewalPeriod,
		})
}

// Type implements basetypes.ObjectValuable.
func (o RateLimit) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"renewal_period": types.StringType,
		},
	}
}

type Route struct {
	ServedEntityName types.String `tfsdk:"served_entity_name"`
	// The name of the served model this route configures traffic for.
	ServedModelName types.String `tfsdk:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage types.Int64 `tfsdk:"traffic_percentage"`
}

func (to *Route) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from Route) {
}

func (to *Route) SyncFieldsDuringRead(ctx context.Context, from Route) {
}

func (c Route) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a Route) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, Route
// only implements ToObjectValue() and Type().
func (o Route) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"served_entity_name": o.ServedEntityName,
			"served_model_name":  o.ServedModelName,
			"traffic_percentage": o.TrafficPercentage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o Route) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entity_name": types.StringType,
			"served_model_name":  types.StringType,
			"traffic_percentage": types.Int64Type,
		},
	}
}

type ServedEntityInput struct {
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
	ExternalModel types.Object `tfsdk:"external_model"`
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

func (to *ServedEntityInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedEntityInput) {
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

func (to *ServedEntityInput) SyncFieldsDuringRead(ctx context.Context, from ServedEntityInput) {
	if !from.ExternalModel.IsNull() && !from.ExternalModel.IsUnknown() {
		if toExternalModel, ok := to.GetExternalModel(ctx); ok {
			if fromExternalModel, ok := from.GetExternalModel(ctx); ok {
				toExternalModel.SyncFieldsDuringRead(ctx, fromExternalModel)
				to.SetExternalModel(ctx, toExternalModel)
			}
		}
	}
}

func (c ServedEntityInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_name"] = attrs["entity_name"].SetOptional()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["environment_vars"] = attrs["environment_vars"].SetOptional()
	attrs["external_model"] = attrs["external_model"].SetOptional()
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
func (a ServedEntityInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"external_model":   reflect.TypeOf(ExternalModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntityInput
// only implements ToObjectValue() and Type().
func (o ServedEntityInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"entity_name":                 o.EntityName,
			"entity_version":              o.EntityVersion,
			"environment_vars":            o.EnvironmentVars,
			"external_model":              o.ExternalModel,
			"instance_profile_arn":        o.InstanceProfileArn,
			"max_provisioned_concurrency": o.MaxProvisionedConcurrency,
			"max_provisioned_throughput":  o.MaxProvisionedThroughput,
			"min_provisioned_concurrency": o.MinProvisionedConcurrency,
			"min_provisioned_throughput":  o.MinProvisionedThroughput,
			"name":                        o.Name,
			"provisioned_model_units":     o.ProvisionedModelUnits,
			"scale_to_zero_enabled":       o.ScaleToZeroEnabled,
			"workload_size":               o.WorkloadSize,
			"workload_type":               o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedEntityInput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"external_model":              ExternalModel{}.Type(ctx),
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

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedEntityInput as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityInput) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedEntityInput.
func (o *ServedEntityInput) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntityInput as
// a ExternalModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityInput) GetExternalModel(ctx context.Context) (ExternalModel, bool) {
	var e ExternalModel
	if o.ExternalModel.IsNull() || o.ExternalModel.IsUnknown() {
		return e, false
	}
	var v ExternalModel
	d := o.ExternalModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntityInput.
func (o *ServedEntityInput) SetExternalModel(ctx context.Context, v ExternalModel) {
	vs := v.ToObjectValue(ctx)
	o.ExternalModel = vs
}

type ServedEntityOutput struct {
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
	ExternalModel types.Object `tfsdk:"external_model"`

	FoundationModel types.Object `tfsdk:"foundation_model"`
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

	State types.Object `tfsdk:"state"`
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

func (to *ServedEntityOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedEntityOutput) {
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

func (to *ServedEntityOutput) SyncFieldsDuringRead(ctx context.Context, from ServedEntityOutput) {
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

func (c ServedEntityOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["entity_name"] = attrs["entity_name"].SetOptional()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["environment_vars"] = attrs["environment_vars"].SetOptional()
	attrs["external_model"] = attrs["external_model"].SetOptional()
	attrs["foundation_model"] = attrs["foundation_model"].SetOptional()
	attrs["instance_profile_arn"] = attrs["instance_profile_arn"].SetOptional()
	attrs["max_provisioned_concurrency"] = attrs["max_provisioned_concurrency"].SetOptional()
	attrs["max_provisioned_throughput"] = attrs["max_provisioned_throughput"].SetOptional()
	attrs["min_provisioned_concurrency"] = attrs["min_provisioned_concurrency"].SetOptional()
	attrs["min_provisioned_throughput"] = attrs["min_provisioned_throughput"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["provisioned_model_units"] = attrs["provisioned_model_units"].SetOptional()
	attrs["scale_to_zero_enabled"] = attrs["scale_to_zero_enabled"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
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
func (a ServedEntityOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"external_model":   reflect.TypeOf(ExternalModel{}),
		"foundation_model": reflect.TypeOf(FoundationModel{}),
		"state":            reflect.TypeOf(ServedModelState{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntityOutput
// only implements ToObjectValue() and Type().
func (o ServedEntityOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":          o.CreationTimestamp,
			"creator":                     o.Creator,
			"entity_name":                 o.EntityName,
			"entity_version":              o.EntityVersion,
			"environment_vars":            o.EnvironmentVars,
			"external_model":              o.ExternalModel,
			"foundation_model":            o.FoundationModel,
			"instance_profile_arn":        o.InstanceProfileArn,
			"max_provisioned_concurrency": o.MaxProvisionedConcurrency,
			"max_provisioned_throughput":  o.MaxProvisionedThroughput,
			"min_provisioned_concurrency": o.MinProvisionedConcurrency,
			"min_provisioned_throughput":  o.MinProvisionedThroughput,
			"name":                        o.Name,
			"provisioned_model_units":     o.ProvisionedModelUnits,
			"scale_to_zero_enabled":       o.ScaleToZeroEnabled,
			"state":                       o.State,
			"workload_size":               o.WorkloadSize,
			"workload_type":               o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedEntityOutput) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"entity_name":        types.StringType,
			"entity_version":     types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"external_model":              ExternalModel{}.Type(ctx),
			"foundation_model":            FoundationModel{}.Type(ctx),
			"instance_profile_arn":        types.StringType,
			"max_provisioned_concurrency": types.Int64Type,
			"max_provisioned_throughput":  types.Int64Type,
			"min_provisioned_concurrency": types.Int64Type,
			"min_provisioned_throughput":  types.Int64Type,
			"name":                        types.StringType,
			"provisioned_model_units":     types.Int64Type,
			"scale_to_zero_enabled":       types.BoolType,
			"state":                       ServedModelState{}.Type(ctx),
			"workload_size":               types.StringType,
			"workload_type":               types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedEntityOutput as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedEntityOutput.
func (o *ServedEntityOutput) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntityOutput as
// a ExternalModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput) GetExternalModel(ctx context.Context) (ExternalModel, bool) {
	var e ExternalModel
	if o.ExternalModel.IsNull() || o.ExternalModel.IsUnknown() {
		return e, false
	}
	var v ExternalModel
	d := o.ExternalModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntityOutput.
func (o *ServedEntityOutput) SetExternalModel(ctx context.Context, v ExternalModel) {
	vs := v.ToObjectValue(ctx)
	o.ExternalModel = vs
}

// GetFoundationModel returns the value of the FoundationModel field in ServedEntityOutput as
// a FoundationModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput) GetFoundationModel(ctx context.Context) (FoundationModel, bool) {
	var e FoundationModel
	if o.FoundationModel.IsNull() || o.FoundationModel.IsUnknown() {
		return e, false
	}
	var v FoundationModel
	d := o.FoundationModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFoundationModel sets the value of the FoundationModel field in ServedEntityOutput.
func (o *ServedEntityOutput) SetFoundationModel(ctx context.Context, v FoundationModel) {
	vs := v.ToObjectValue(ctx)
	o.FoundationModel = vs
}

// GetState returns the value of the State field in ServedEntityOutput as
// a ServedModelState value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntityOutput) GetState(ctx context.Context) (ServedModelState, bool) {
	var e ServedModelState
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v ServedModelState
	d := o.State.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetState sets the value of the State field in ServedEntityOutput.
func (o *ServedEntityOutput) SetState(ctx context.Context, v ServedModelState) {
	vs := v.ToObjectValue(ctx)
	o.State = vs
}

type ServedEntitySpec struct {
	EntityName types.String `tfsdk:"entity_name"`

	EntityVersion types.String `tfsdk:"entity_version"`

	ExternalModel types.Object `tfsdk:"external_model"`

	FoundationModel types.Object `tfsdk:"foundation_model"`

	Name types.String `tfsdk:"name"`
}

func (to *ServedEntitySpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedEntitySpec) {
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

func (to *ServedEntitySpec) SyncFieldsDuringRead(ctx context.Context, from ServedEntitySpec) {
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

func (c ServedEntitySpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["entity_name"] = attrs["entity_name"].SetOptional()
	attrs["entity_version"] = attrs["entity_version"].SetOptional()
	attrs["external_model"] = attrs["external_model"].SetOptional()
	attrs["foundation_model"] = attrs["foundation_model"].SetOptional()
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
func (a ServedEntitySpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"external_model":   reflect.TypeOf(ExternalModel{}),
		"foundation_model": reflect.TypeOf(FoundationModel{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedEntitySpec
// only implements ToObjectValue() and Type().
func (o ServedEntitySpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServedEntitySpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":      types.StringType,
			"entity_version":   types.StringType,
			"external_model":   ExternalModel{}.Type(ctx),
			"foundation_model": FoundationModel{}.Type(ctx),
			"name":             types.StringType,
		},
	}
}

// GetExternalModel returns the value of the ExternalModel field in ServedEntitySpec as
// a ExternalModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntitySpec) GetExternalModel(ctx context.Context) (ExternalModel, bool) {
	var e ExternalModel
	if o.ExternalModel.IsNull() || o.ExternalModel.IsUnknown() {
		return e, false
	}
	var v ExternalModel
	d := o.ExternalModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetExternalModel sets the value of the ExternalModel field in ServedEntitySpec.
func (o *ServedEntitySpec) SetExternalModel(ctx context.Context, v ExternalModel) {
	vs := v.ToObjectValue(ctx)
	o.ExternalModel = vs
}

// GetFoundationModel returns the value of the FoundationModel field in ServedEntitySpec as
// a FoundationModel value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedEntitySpec) GetFoundationModel(ctx context.Context) (FoundationModel, bool) {
	var e FoundationModel
	if o.FoundationModel.IsNull() || o.FoundationModel.IsUnknown() {
		return e, false
	}
	var v FoundationModel
	d := o.FoundationModel.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetFoundationModel sets the value of the FoundationModel field in ServedEntitySpec.
func (o *ServedEntitySpec) SetFoundationModel(ctx context.Context, v FoundationModel) {
	vs := v.ToObjectValue(ctx)
	o.FoundationModel = vs
}

type ServedModelInput struct {
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

func (to *ServedModelInput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelInput) {
}

func (to *ServedModelInput) SyncFieldsDuringRead(ctx context.Context, from ServedModelInput) {
}

func (c ServedModelInput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServedModelInput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelInput
// only implements ToObjectValue() and Type().
func (o ServedModelInput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"environment_vars":            o.EnvironmentVars,
			"instance_profile_arn":        o.InstanceProfileArn,
			"max_provisioned_concurrency": o.MaxProvisionedConcurrency,
			"max_provisioned_throughput":  o.MaxProvisionedThroughput,
			"min_provisioned_concurrency": o.MinProvisionedConcurrency,
			"min_provisioned_throughput":  o.MinProvisionedThroughput,
			"model_name":                  o.ModelName,
			"model_version":               o.ModelVersion,
			"name":                        o.Name,
			"provisioned_model_units":     o.ProvisionedModelUnits,
			"scale_to_zero_enabled":       o.ScaleToZeroEnabled,
			"workload_size":               o.WorkloadSize,
			"workload_type":               o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelInput) Type(ctx context.Context) attr.Type {
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

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedModelInput as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedModelInput) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedModelInput.
func (o *ServedModelInput) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

type ServedModelOutput struct {
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

	State types.Object `tfsdk:"state"`
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

func (to *ServedModelOutput) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelOutput) {
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

func (to *ServedModelOutput) SyncFieldsDuringRead(ctx context.Context, from ServedModelOutput) {
	if !from.State.IsNull() && !from.State.IsUnknown() {
		if toState, ok := to.GetState(ctx); ok {
			if fromState, ok := from.GetState(ctx); ok {
				toState.SyncFieldsDuringRead(ctx, fromState)
				to.SetState(ctx, toState)
			}
		}
	}
}

func (c ServedModelOutput) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServedModelOutput) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"environment_vars": reflect.TypeOf(types.String{}),
		"state":            reflect.TypeOf(ServedModelState{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelOutput
// only implements ToObjectValue() and Type().
func (o ServedModelOutput) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"creation_timestamp":          o.CreationTimestamp,
			"creator":                     o.Creator,
			"environment_vars":            o.EnvironmentVars,
			"instance_profile_arn":        o.InstanceProfileArn,
			"max_provisioned_concurrency": o.MaxProvisionedConcurrency,
			"min_provisioned_concurrency": o.MinProvisionedConcurrency,
			"model_name":                  o.ModelName,
			"model_version":               o.ModelVersion,
			"name":                        o.Name,
			"provisioned_model_units":     o.ProvisionedModelUnits,
			"scale_to_zero_enabled":       o.ScaleToZeroEnabled,
			"state":                       o.State,
			"workload_size":               o.WorkloadSize,
			"workload_type":               o.WorkloadType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelOutput) Type(ctx context.Context) attr.Type {
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
			"state":                       ServedModelState{}.Type(ctx),
			"workload_size":               types.StringType,
			"workload_type":               types.StringType,
		},
	}
}

// GetEnvironmentVars returns the value of the EnvironmentVars field in ServedModelOutput as
// a map of string to types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedModelOutput) GetEnvironmentVars(ctx context.Context) (map[string]types.String, bool) {
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

// SetEnvironmentVars sets the value of the EnvironmentVars field in ServedModelOutput.
func (o *ServedModelOutput) SetEnvironmentVars(ctx context.Context, v map[string]types.String) {
	vs := make(map[string]attr.Value, len(v))
	for k, e := range v {
		vs[k] = e
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["environment_vars"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.EnvironmentVars = types.MapValueMust(t, vs)
}

// GetState returns the value of the State field in ServedModelOutput as
// a ServedModelState value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServedModelOutput) GetState(ctx context.Context) (ServedModelState, bool) {
	var e ServedModelState
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v ServedModelState
	d := o.State.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetState sets the value of the State field in ServedModelOutput.
func (o *ServedModelOutput) SetState(ctx context.Context, v ServedModelState) {
	vs := v.ToObjectValue(ctx)
	o.State = vs
}

type ServedModelSpec struct {
	// Only one of model_name and entity_name should be populated
	ModelName types.String `tfsdk:"model_name"`
	// Only one of model_version and entity_version should be populated
	ModelVersion types.String `tfsdk:"model_version"`

	Name types.String `tfsdk:"name"`
}

func (to *ServedModelSpec) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelSpec) {
}

func (to *ServedModelSpec) SyncFieldsDuringRead(ctx context.Context, from ServedModelSpec) {
}

func (c ServedModelSpec) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServedModelSpec) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelSpec
// only implements ToObjectValue() and Type().
func (o ServedModelSpec) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"model_name":    o.ModelName,
			"model_version": o.ModelVersion,
			"name":          o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelSpec) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_name":    types.StringType,
			"model_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

type ServedModelState struct {
	Deployment types.String `tfsdk:"deployment"`

	DeploymentStateMessage types.String `tfsdk:"deployment_state_message"`
}

func (to *ServedModelState) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServedModelState) {
}

func (to *ServedModelState) SyncFieldsDuringRead(ctx context.Context, from ServedModelState) {
}

func (c ServedModelState) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServedModelState) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServedModelState
// only implements ToObjectValue() and Type().
func (o ServedModelState) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"deployment":               o.Deployment,
			"deployment_state_message": o.DeploymentStateMessage,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServedModelState) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"deployment":               types.StringType,
			"deployment_state_message": types.StringType,
		},
	}
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs types.String `tfsdk:"logs"`
}

func (to *ServerLogsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServerLogsResponse) {
}

func (to *ServerLogsResponse) SyncFieldsDuringRead(ctx context.Context, from ServerLogsResponse) {
}

func (c ServerLogsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServerLogsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServerLogsResponse
// only implements ToObjectValue() and Type().
func (o ServerLogsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"logs": o.Logs,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServerLogsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway types.Object `tfsdk:"ai_gateway"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The config that is currently being served by the endpoint.
	Config types.Object `tfsdk:"config"`
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
	State types.Object `tfsdk:"state"`
	// Tags attached to the serving endpoint.
	Tags types.List `tfsdk:"tags"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task"`
	// The usage policy associated with serving endpoint.
	UsagePolicyId types.String `tfsdk:"usage_policy_id"`
}

func (to *ServingEndpoint) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpoint) {
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

func (to *ServingEndpoint) SyncFieldsDuringRead(ctx context.Context, from ServingEndpoint) {
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

func (c ServingEndpoint) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
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
func (a ServingEndpoint) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway": reflect.TypeOf(AiGatewayConfig{}),
		"config":     reflect.TypeOf(EndpointCoreConfigSummary{}),
		"state":      reflect.TypeOf(EndpointState{}),
		"tags":       reflect.TypeOf(EndpointTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpoint
// only implements ToObjectValue() and Type().
func (o ServingEndpoint) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":             o.AiGateway,
			"budget_policy_id":       o.BudgetPolicyId,
			"config":                 o.Config,
			"creation_timestamp":     o.CreationTimestamp,
			"creator":                o.Creator,
			"description":            o.Description,
			"id":                     o.Id,
			"last_updated_timestamp": o.LastUpdatedTimestamp,
			"name":                   o.Name,
			"state":                  o.State,
			"tags":                   o.Tags,
			"task":                   o.Task,
			"usage_policy_id":        o.UsagePolicyId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpoint) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway":             AiGatewayConfig{}.Type(ctx),
			"budget_policy_id":       types.StringType,
			"config":                 EndpointCoreConfigSummary{}.Type(ctx),
			"creation_timestamp":     types.Int64Type,
			"creator":                types.StringType,
			"description":            types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"state":                  EndpointState{}.Type(ctx),
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
			"task":            types.StringType,
			"usage_policy_id": types.StringType,
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in ServingEndpoint as
// a AiGatewayConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint) GetAiGateway(ctx context.Context) (AiGatewayConfig, bool) {
	var e AiGatewayConfig
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v AiGatewayConfig
	d := o.AiGateway.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAiGateway sets the value of the AiGateway field in ServingEndpoint.
func (o *ServingEndpoint) SetAiGateway(ctx context.Context, v AiGatewayConfig) {
	vs := v.ToObjectValue(ctx)
	o.AiGateway = vs
}

// GetConfig returns the value of the Config field in ServingEndpoint as
// a EndpointCoreConfigSummary value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint) GetConfig(ctx context.Context) (EndpointCoreConfigSummary, bool) {
	var e EndpointCoreConfigSummary
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v EndpointCoreConfigSummary
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in ServingEndpoint.
func (o *ServingEndpoint) SetConfig(ctx context.Context, v EndpointCoreConfigSummary) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

// GetState returns the value of the State field in ServingEndpoint as
// a EndpointState value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint) GetState(ctx context.Context) (EndpointState, bool) {
	var e EndpointState
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v EndpointState
	d := o.State.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetState sets the value of the State field in ServingEndpoint.
func (o *ServingEndpoint) SetState(ctx context.Context, v EndpointState) {
	vs := v.ToObjectValue(ctx)
	o.State = vs
}

// GetTags returns the value of the Tags field in ServingEndpoint as
// a slice of EndpointTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpoint) GetTags(ctx context.Context) ([]EndpointTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ServingEndpoint.
func (o *ServingEndpoint) SetTags(ctx context.Context, v []EndpointTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`

	PermissionLevel types.String `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

func (to *ServingEndpointAccessControlRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointAccessControlRequest) {
}

func (to *ServingEndpointAccessControlRequest) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointAccessControlRequest) {
}

func (c ServingEndpointAccessControlRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServingEndpointAccessControlRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointAccessControlRequest
// only implements ToObjectValue() and Type().
func (o ServingEndpointAccessControlRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServingEndpointAccessControlRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_name":             types.StringType,
			"permission_level":       types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ServingEndpointAccessControlResponse struct {
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

func (to *ServingEndpointAccessControlResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (to *ServingEndpointAccessControlResponse) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointAccessControlResponse) {
	if !from.AllPermissions.IsNull() && !from.AllPermissions.IsUnknown() && to.AllPermissions.IsNull() && len(from.AllPermissions.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AllPermissions, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AllPermissions = from.AllPermissions
	}
}

func (c ServingEndpointAccessControlResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServingEndpointAccessControlResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"all_permissions": reflect.TypeOf(ServingEndpointPermission{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointAccessControlResponse
// only implements ToObjectValue() and Type().
func (o ServingEndpointAccessControlResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
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
func (o ServingEndpointAccessControlResponse) Type(ctx context.Context) attr.Type {
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

// GetAllPermissions returns the value of the AllPermissions field in ServingEndpointAccessControlResponse as
// a slice of ServingEndpointPermission values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointAccessControlResponse) GetAllPermissions(ctx context.Context) ([]ServingEndpointPermission, bool) {
	if o.AllPermissions.IsNull() || o.AllPermissions.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointPermission
	d := o.AllPermissions.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAllPermissions sets the value of the AllPermissions field in ServingEndpointAccessControlResponse.
func (o *ServingEndpointAccessControlResponse) SetAllPermissions(ctx context.Context, v []ServingEndpointPermission) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["all_permissions"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AllPermissions = types.ListValueMust(t, vs)
}

type ServingEndpointDetailed struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway types.Object `tfsdk:"ai_gateway"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId types.String `tfsdk:"budget_policy_id"`
	// The config that is currently being served by the endpoint.
	Config types.Object `tfsdk:"config"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo types.Object `tfsdk:"data_plane_info"`
	// Description of the serving model
	Description types.String `tfsdk:"description"`
	// Email notification settings.
	EmailNotifications types.Object `tfsdk:"email_notifications"`
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
	PendingConfig types.Object `tfsdk:"pending_config"`
	// The permission level of the principal making the request.
	PermissionLevel types.String `tfsdk:"permission_level"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized types.Bool `tfsdk:"route_optimized"`
	// Information corresponding to the state of the serving endpoint.
	State types.Object `tfsdk:"state"`
	// Tags attached to the serving endpoint.
	Tags types.List `tfsdk:"tags"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task"`
}

func (to *ServingEndpointDetailed) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointDetailed) {
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

func (to *ServingEndpointDetailed) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointDetailed) {
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

func (c ServingEndpointDetailed) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["ai_gateway"] = attrs["ai_gateway"].SetOptional()
	attrs["budget_policy_id"] = attrs["budget_policy_id"].SetOptional()
	attrs["config"] = attrs["config"].SetOptional()
	attrs["creation_timestamp"] = attrs["creation_timestamp"].SetOptional()
	attrs["creator"] = attrs["creator"].SetOptional()
	attrs["data_plane_info"] = attrs["data_plane_info"].SetOptional()
	attrs["description"] = attrs["description"].SetOptional()
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
	attrs["endpoint_url"] = attrs["endpoint_url"].SetOptional()
	attrs["id"] = attrs["id"].SetOptional()
	attrs["last_updated_timestamp"] = attrs["last_updated_timestamp"].SetOptional()
	attrs["name"] = attrs["name"].SetOptional()
	attrs["pending_config"] = attrs["pending_config"].SetOptional()
	attrs["permission_level"] = attrs["permission_level"].SetOptional()
	attrs["route_optimized"] = attrs["route_optimized"].SetOptional()
	attrs["state"] = attrs["state"].SetOptional()
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
func (a ServingEndpointDetailed) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"ai_gateway":          reflect.TypeOf(AiGatewayConfig{}),
		"config":              reflect.TypeOf(EndpointCoreConfigOutput{}),
		"data_plane_info":     reflect.TypeOf(ModelDataPlaneInfo{}),
		"email_notifications": reflect.TypeOf(EmailNotifications{}),
		"pending_config":      reflect.TypeOf(EndpointPendingConfig{}),
		"state":               reflect.TypeOf(EndpointState{}),
		"tags":                reflect.TypeOf(EndpointTag{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointDetailed
// only implements ToObjectValue() and Type().
func (o ServingEndpointDetailed) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"ai_gateway":             o.AiGateway,
			"budget_policy_id":       o.BudgetPolicyId,
			"config":                 o.Config,
			"creation_timestamp":     o.CreationTimestamp,
			"creator":                o.Creator,
			"data_plane_info":        o.DataPlaneInfo,
			"description":            o.Description,
			"email_notifications":    o.EmailNotifications,
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
func (o ServingEndpointDetailed) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway":             AiGatewayConfig{}.Type(ctx),
			"budget_policy_id":       types.StringType,
			"config":                 EndpointCoreConfigOutput{}.Type(ctx),
			"creation_timestamp":     types.Int64Type,
			"creator":                types.StringType,
			"data_plane_info":        ModelDataPlaneInfo{}.Type(ctx),
			"description":            types.StringType,
			"email_notifications":    EmailNotifications{}.Type(ctx),
			"endpoint_url":           types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"pending_config":         EndpointPendingConfig{}.Type(ctx),
			"permission_level":       types.StringType,
			"route_optimized":        types.BoolType,
			"state":                  EndpointState{}.Type(ctx),
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.Type(ctx),
			},
			"task": types.StringType,
		},
	}
}

// GetAiGateway returns the value of the AiGateway field in ServingEndpointDetailed as
// a AiGatewayConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetAiGateway(ctx context.Context) (AiGatewayConfig, bool) {
	var e AiGatewayConfig
	if o.AiGateway.IsNull() || o.AiGateway.IsUnknown() {
		return e, false
	}
	var v AiGatewayConfig
	d := o.AiGateway.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAiGateway sets the value of the AiGateway field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetAiGateway(ctx context.Context, v AiGatewayConfig) {
	vs := v.ToObjectValue(ctx)
	o.AiGateway = vs
}

// GetConfig returns the value of the Config field in ServingEndpointDetailed as
// a EndpointCoreConfigOutput value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetConfig(ctx context.Context) (EndpointCoreConfigOutput, bool) {
	var e EndpointCoreConfigOutput
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v EndpointCoreConfigOutput
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetConfig(ctx context.Context, v EndpointCoreConfigOutput) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

// GetDataPlaneInfo returns the value of the DataPlaneInfo field in ServingEndpointDetailed as
// a ModelDataPlaneInfo value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetDataPlaneInfo(ctx context.Context) (ModelDataPlaneInfo, bool) {
	var e ModelDataPlaneInfo
	if o.DataPlaneInfo.IsNull() || o.DataPlaneInfo.IsUnknown() {
		return e, false
	}
	var v ModelDataPlaneInfo
	d := o.DataPlaneInfo.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetDataPlaneInfo sets the value of the DataPlaneInfo field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetDataPlaneInfo(ctx context.Context, v ModelDataPlaneInfo) {
	vs := v.ToObjectValue(ctx)
	o.DataPlaneInfo = vs
}

// GetEmailNotifications returns the value of the EmailNotifications field in ServingEndpointDetailed as
// a EmailNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetEmailNotifications(ctx context.Context) (EmailNotifications, bool) {
	var e EmailNotifications
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v EmailNotifications
	d := o.EmailNotifications.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailNotifications sets the value of the EmailNotifications field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetEmailNotifications(ctx context.Context, v EmailNotifications) {
	vs := v.ToObjectValue(ctx)
	o.EmailNotifications = vs
}

// GetPendingConfig returns the value of the PendingConfig field in ServingEndpointDetailed as
// a EndpointPendingConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetPendingConfig(ctx context.Context) (EndpointPendingConfig, bool) {
	var e EndpointPendingConfig
	if o.PendingConfig.IsNull() || o.PendingConfig.IsUnknown() {
		return e, false
	}
	var v EndpointPendingConfig
	d := o.PendingConfig.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetPendingConfig sets the value of the PendingConfig field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetPendingConfig(ctx context.Context, v EndpointPendingConfig) {
	vs := v.ToObjectValue(ctx)
	o.PendingConfig = vs
}

// GetState returns the value of the State field in ServingEndpointDetailed as
// a EndpointState value.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetState(ctx context.Context) (EndpointState, bool) {
	var e EndpointState
	if o.State.IsNull() || o.State.IsUnknown() {
		return e, false
	}
	var v EndpointState
	d := o.State.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetState sets the value of the State field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetState(ctx context.Context, v EndpointState) {
	vs := v.ToObjectValue(ctx)
	o.State = vs
}

// GetTags returns the value of the Tags field in ServingEndpointDetailed as
// a slice of EndpointTag values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointDetailed) GetTags(ctx context.Context) ([]EndpointTag, bool) {
	if o.Tags.IsNull() || o.Tags.IsUnknown() {
		return nil, false
	}
	var v []EndpointTag
	d := o.Tags.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetTags sets the value of the Tags field in ServingEndpointDetailed.
func (o *ServingEndpointDetailed) SetTags(ctx context.Context, v []EndpointTag) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["tags"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Tags = types.ListValueMust(t, vs)
}

type ServingEndpointPermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ServingEndpointPermission) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (to *ServingEndpointPermission) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermission) {
	if !from.InheritedFromObject.IsNull() && !from.InheritedFromObject.IsUnknown() && to.InheritedFromObject.IsNull() && len(from.InheritedFromObject.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for InheritedFromObject, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.InheritedFromObject = from.InheritedFromObject
	}
}

func (c ServingEndpointPermission) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServingEndpointPermission) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"inherited_from_object": reflect.TypeOf(types.String{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermission
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermission) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"inherited":             o.Inherited,
			"inherited_from_object": o.InheritedFromObject,
			"permission_level":      o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermission) Type(ctx context.Context) attr.Type {
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

// GetInheritedFromObject returns the value of the InheritedFromObject field in ServingEndpointPermission as
// a slice of types.String values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointPermission) GetInheritedFromObject(ctx context.Context) ([]types.String, bool) {
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

// SetInheritedFromObject sets the value of the InheritedFromObject field in ServingEndpointPermission.
func (o *ServingEndpointPermission) SetInheritedFromObject(ctx context.Context, v []types.String) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e)
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["inherited_from_object"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.InheritedFromObject = types.ListValueMust(t, vs)
}

type ServingEndpointPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

func (to *ServingEndpointPermissions) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ServingEndpointPermissions) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermissions) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c ServingEndpointPermissions) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServingEndpointPermissions) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ServingEndpointAccessControlResponse{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissions
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermissions) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"object_id":           o.ObjectId,
			"object_type":         o.ObjectType,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermissions) Type(ctx context.Context) attr.Type {
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

// GetAccessControlList returns the value of the AccessControlList field in ServingEndpointPermissions as
// a slice of ServingEndpointAccessControlResponse values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointPermissions) GetAccessControlList(ctx context.Context) ([]ServingEndpointAccessControlResponse, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointAccessControlResponse
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ServingEndpointPermissions.
func (o *ServingEndpointPermissions) SetAccessControlList(ctx context.Context, v []ServingEndpointAccessControlResponse) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type ServingEndpointPermissionsDescription struct {
	Description types.String `tfsdk:"description"`

	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (to *ServingEndpointPermissionsDescription) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermissionsDescription) {
}

func (to *ServingEndpointPermissionsDescription) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermissionsDescription) {
}

func (c ServingEndpointPermissionsDescription) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServingEndpointPermissionsDescription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissionsDescription
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermissionsDescription) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"description":      o.Description,
			"permission_level": o.PermissionLevel,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermissionsDescription) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (to *ServingEndpointPermissionsRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from ServingEndpointPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (to *ServingEndpointPermissionsRequest) SyncFieldsDuringRead(ctx context.Context, from ServingEndpointPermissionsRequest) {
	if !from.AccessControlList.IsNull() && !from.AccessControlList.IsUnknown() && to.AccessControlList.IsNull() && len(from.AccessControlList.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for AccessControlList, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.AccessControlList = from.AccessControlList
	}
}

func (c ServingEndpointPermissionsRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a ServingEndpointPermissionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"access_control_list": reflect.TypeOf(ServingEndpointAccessControlRequest{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, ServingEndpointPermissionsRequest
// only implements ToObjectValue() and Type().
func (o ServingEndpointPermissionsRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"access_control_list": o.AccessControlList,
			"serving_endpoint_id": o.ServingEndpointId,
		})
}

// Type implements basetypes.ObjectValuable.
func (o ServingEndpointPermissionsRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlRequest{}.Type(ctx),
			},
			"serving_endpoint_id": types.StringType,
		},
	}
}

// GetAccessControlList returns the value of the AccessControlList field in ServingEndpointPermissionsRequest as
// a slice of ServingEndpointAccessControlRequest values.
// If the field is unknown or null, the boolean return value is false.
func (o *ServingEndpointPermissionsRequest) GetAccessControlList(ctx context.Context) ([]ServingEndpointAccessControlRequest, bool) {
	if o.AccessControlList.IsNull() || o.AccessControlList.IsUnknown() {
		return nil, false
	}
	var v []ServingEndpointAccessControlRequest
	d := o.AccessControlList.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetAccessControlList sets the value of the AccessControlList field in ServingEndpointPermissionsRequest.
func (o *ServingEndpointPermissionsRequest) SetAccessControlList(ctx context.Context, v []ServingEndpointAccessControlRequest) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["access_control_list"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.AccessControlList = types.ListValueMust(t, vs)
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes types.List `tfsdk:"routes"`
}

func (to *TrafficConfig) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from TrafficConfig) {
	if !from.Routes.IsNull() && !from.Routes.IsUnknown() && to.Routes.IsNull() && len(from.Routes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Routes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Routes = from.Routes
	}
}

func (to *TrafficConfig) SyncFieldsDuringRead(ctx context.Context, from TrafficConfig) {
	if !from.Routes.IsNull() && !from.Routes.IsUnknown() && to.Routes.IsNull() && len(from.Routes.Elements()) == 0 {
		// The default representation of an empty list for TF autogenerated resources in the resource state is Null.
		// If a user specified a non-Null, empty list for Routes, and the deserialized field value is Null,
		// set the resulting resource state to the empty list to match the planned value.
		to.Routes = from.Routes
	}
}

func (c TrafficConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
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
func (a TrafficConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"routes": reflect.TypeOf(Route{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TrafficConfig
// only implements ToObjectValue() and Type().
func (o TrafficConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"routes": o.Routes,
		})
}

// Type implements basetypes.ObjectValuable.
func (o TrafficConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"routes": basetypes.ListType{
				ElemType: Route{}.Type(ctx),
			},
		},
	}
}

// GetRoutes returns the value of the Routes field in TrafficConfig as
// a slice of Route values.
// If the field is unknown or null, the boolean return value is false.
func (o *TrafficConfig) GetRoutes(ctx context.Context) ([]Route, bool) {
	if o.Routes.IsNull() || o.Routes.IsUnknown() {
		return nil, false
	}
	var v []Route
	d := o.Routes.ElementsAs(ctx, &v, true)
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetRoutes sets the value of the Routes field in TrafficConfig.
func (o *TrafficConfig) SetRoutes(ctx context.Context, v []Route) {
	vs := make([]attr.Value, 0, len(v))
	for _, e := range v {
		vs = append(vs, e.ToObjectValue(ctx))
	}
	t := o.Type(ctx).(basetypes.ObjectType).AttrTypes["routes"]
	t = t.(attr.TypeWithElementType).ElementType()
	o.Routes = types.ListValueMust(t, vs)
}

type UpdateInferenceEndpointNotifications struct {
	// The email notification settings to update. Specify email addresses to
	// notify when endpoint state changes occur.
	EmailNotifications types.Object `tfsdk:"email_notifications"`
	// The name of the serving endpoint whose notifications are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateInferenceEndpointNotifications) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInferenceEndpointNotifications) {
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

func (to *UpdateInferenceEndpointNotifications) SyncFieldsDuringRead(ctx context.Context, from UpdateInferenceEndpointNotifications) {
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
}

func (c UpdateInferenceEndpointNotifications) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
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
func (a UpdateInferenceEndpointNotifications) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_notifications": reflect.TypeOf(EmailNotifications{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInferenceEndpointNotifications
// only implements ToObjectValue() and Type().
func (o UpdateInferenceEndpointNotifications) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_notifications": o.EmailNotifications,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateInferenceEndpointNotifications) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_notifications": EmailNotifications{}.Type(ctx),
			"name":                types.StringType,
		},
	}
}

// GetEmailNotifications returns the value of the EmailNotifications field in UpdateInferenceEndpointNotifications as
// a EmailNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInferenceEndpointNotifications) GetEmailNotifications(ctx context.Context) (EmailNotifications, bool) {
	var e EmailNotifications
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v EmailNotifications
	d := o.EmailNotifications.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailNotifications sets the value of the EmailNotifications field in UpdateInferenceEndpointNotifications.
func (o *UpdateInferenceEndpointNotifications) SetEmailNotifications(ctx context.Context, v EmailNotifications) {
	vs := v.ToObjectValue(ctx)
	o.EmailNotifications = vs
}

type UpdateInferenceEndpointNotificationsResponse struct {
	EmailNotifications types.Object `tfsdk:"email_notifications"`

	Name types.String `tfsdk:"name"`
}

func (to *UpdateInferenceEndpointNotificationsResponse) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateInferenceEndpointNotificationsResponse) {
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

func (to *UpdateInferenceEndpointNotificationsResponse) SyncFieldsDuringRead(ctx context.Context, from UpdateInferenceEndpointNotificationsResponse) {
	if !from.EmailNotifications.IsNull() && !from.EmailNotifications.IsUnknown() {
		if toEmailNotifications, ok := to.GetEmailNotifications(ctx); ok {
			if fromEmailNotifications, ok := from.GetEmailNotifications(ctx); ok {
				toEmailNotifications.SyncFieldsDuringRead(ctx, fromEmailNotifications)
				to.SetEmailNotifications(ctx, toEmailNotifications)
			}
		}
	}
}

func (c UpdateInferenceEndpointNotificationsResponse) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["email_notifications"] = attrs["email_notifications"].SetOptional()
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
func (a UpdateInferenceEndpointNotificationsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"email_notifications": reflect.TypeOf(EmailNotifications{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateInferenceEndpointNotificationsResponse
// only implements ToObjectValue() and Type().
func (o UpdateInferenceEndpointNotificationsResponse) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"email_notifications": o.EmailNotifications,
			"name":                o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateInferenceEndpointNotificationsResponse) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"email_notifications": EmailNotifications{}.Type(ctx),
			"name":                types.StringType,
		},
	}
}

// GetEmailNotifications returns the value of the EmailNotifications field in UpdateInferenceEndpointNotificationsResponse as
// a EmailNotifications value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateInferenceEndpointNotificationsResponse) GetEmailNotifications(ctx context.Context) (EmailNotifications, bool) {
	var e EmailNotifications
	if o.EmailNotifications.IsNull() || o.EmailNotifications.IsUnknown() {
		return e, false
	}
	var v EmailNotifications
	d := o.EmailNotifications.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetEmailNotifications sets the value of the EmailNotifications field in UpdateInferenceEndpointNotificationsResponse.
func (o *UpdateInferenceEndpointNotificationsResponse) SetEmailNotifications(ctx context.Context, v EmailNotifications) {
	vs := v.ToObjectValue(ctx)
	o.EmailNotifications = vs
}

type UpdateProvisionedThroughputEndpointConfigRequest struct {
	Config types.Object `tfsdk:"config"`
	// The name of the pt endpoint to update. This field is required.
	Name types.String `tfsdk:"-"`
}

func (to *UpdateProvisionedThroughputEndpointConfigRequest) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from UpdateProvisionedThroughputEndpointConfigRequest) {
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

func (to *UpdateProvisionedThroughputEndpointConfigRequest) SyncFieldsDuringRead(ctx context.Context, from UpdateProvisionedThroughputEndpointConfigRequest) {
	if !from.Config.IsNull() && !from.Config.IsUnknown() {
		if toConfig, ok := to.GetConfig(ctx); ok {
			if fromConfig, ok := from.GetConfig(ctx); ok {
				toConfig.SyncFieldsDuringRead(ctx, fromConfig)
				to.SetConfig(ctx, toConfig)
			}
		}
	}
}

func (c UpdateProvisionedThroughputEndpointConfigRequest) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["config"] = attrs["config"].SetRequired()
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
func (a UpdateProvisionedThroughputEndpointConfigRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"config": reflect.TypeOf(PtEndpointCoreConfig{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, UpdateProvisionedThroughputEndpointConfigRequest
// only implements ToObjectValue() and Type().
func (o UpdateProvisionedThroughputEndpointConfigRequest) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"config": o.Config,
			"name":   o.Name,
		})
}

// Type implements basetypes.ObjectValuable.
func (o UpdateProvisionedThroughputEndpointConfigRequest) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": PtEndpointCoreConfig{}.Type(ctx),
			"name":   types.StringType,
		},
	}
}

// GetConfig returns the value of the Config field in UpdateProvisionedThroughputEndpointConfigRequest as
// a PtEndpointCoreConfig value.
// If the field is unknown or null, the boolean return value is false.
func (o *UpdateProvisionedThroughputEndpointConfigRequest) GetConfig(ctx context.Context) (PtEndpointCoreConfig, bool) {
	var e PtEndpointCoreConfig
	if o.Config.IsNull() || o.Config.IsUnknown() {
		return e, false
	}
	var v PtEndpointCoreConfig
	d := o.Config.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetConfig sets the value of the Config field in UpdateProvisionedThroughputEndpointConfigRequest.
func (o *UpdateProvisionedThroughputEndpointConfigRequest) SetConfig(ctx context.Context, v PtEndpointCoreConfig) {
	vs := v.ToObjectValue(ctx)
	o.Config = vs
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason types.String `tfsdk:"finish_reason"`
	// The index of the choice in the __chat or completions__ response.
	Index types.Int64 `tfsdk:"index"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs types.Int64 `tfsdk:"logprobs"`
	// The message response from the __chat__ endpoint.
	Message types.Object `tfsdk:"message"`
	// The text response from the __completions__ endpoint.
	Text types.String `tfsdk:"text"`
}

func (to *V1ResponseChoiceElement) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from V1ResponseChoiceElement) {
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

func (to *V1ResponseChoiceElement) SyncFieldsDuringRead(ctx context.Context, from V1ResponseChoiceElement) {
	if !from.Message.IsNull() && !from.Message.IsUnknown() {
		if toMessage, ok := to.GetMessage(ctx); ok {
			if fromMessage, ok := from.GetMessage(ctx); ok {
				toMessage.SyncFieldsDuringRead(ctx, fromMessage)
				to.SetMessage(ctx, toMessage)
			}
		}
	}
}

func (c V1ResponseChoiceElement) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["finish_reason"] = attrs["finish_reason"].SetOptional()
	attrs["index"] = attrs["index"].SetOptional()
	attrs["logprobs"] = attrs["logprobs"].SetOptional()
	attrs["message"] = attrs["message"].SetOptional()
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
func (a V1ResponseChoiceElement) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"message": reflect.TypeOf(ChatMessage{}),
	}
}

// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, V1ResponseChoiceElement
// only implements ToObjectValue() and Type().
func (o V1ResponseChoiceElement) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		o.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"finish_reason": o.FinishReason,
			"index":         o.Index,
			"logprobs":      o.Logprobs,
			"message":       o.Message,
			"text":          o.Text,
		})
}

// Type implements basetypes.ObjectValuable.
func (o V1ResponseChoiceElement) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"finish_reason": types.StringType,
			"index":         types.Int64Type,
			"logprobs":      types.Int64Type,
			"message":       ChatMessage{}.Type(ctx),
			"text":          types.StringType,
		},
	}
}

// GetMessage returns the value of the Message field in V1ResponseChoiceElement as
// a ChatMessage value.
// If the field is unknown or null, the boolean return value is false.
func (o *V1ResponseChoiceElement) GetMessage(ctx context.Context) (ChatMessage, bool) {
	var e ChatMessage
	if o.Message.IsNull() || o.Message.IsUnknown() {
		return e, false
	}
	var v ChatMessage
	d := o.Message.As(ctx, &v, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})
	if d.HasError() {
		panic(pluginfwcommon.DiagToString(d))
	}
	return v, true
}

// SetMessage sets the value of the Message field in V1ResponseChoiceElement.
func (o *V1ResponseChoiceElement) SetMessage(ctx context.Context, v ChatMessage) {
	vs := v.ToObjectValue(ctx)
	o.Message = vs
}
