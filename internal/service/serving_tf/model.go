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
	"io"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/terraform-provider-databricks/internal/service/oauth2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type Ai21LabsConfig struct {
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

func (newState *Ai21LabsConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan Ai21LabsConfig) {
}

func (newState *Ai21LabsConfig) SyncEffectiveFieldsDuringRead(existingState Ai21LabsConfig) {
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

// ToAttrType returns the representation of Ai21LabsConfig in the Terraform plugin framework type
// system.
func (a Ai21LabsConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_api_key":           types.StringType,
			"ai21labs_api_key_plaintext": types.StringType,
		},
	}
}

type AiGatewayConfig struct {
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

func (newState *AiGatewayConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayConfig) {
}

func (newState *AiGatewayConfig) SyncEffectiveFieldsDuringRead(existingState AiGatewayConfig) {
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
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig{}),
	}
}

// ToAttrType returns the representation of AiGatewayConfig in the Terraform plugin framework type
// system.
func (a AiGatewayConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails{}.ToAttrType(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig{}.ToAttrType(ctx),
			},
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.ToAttrType(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type AiGatewayGuardrailParameters struct {
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

func (newState *AiGatewayGuardrailParameters) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrailParameters) {
}

func (newState *AiGatewayGuardrailParameters) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrailParameters) {
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

// ToAttrType returns the representation of AiGatewayGuardrailParameters in the Terraform plugin framework type
// system.
func (a AiGatewayGuardrailParameters) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"invalid_keywords": basetypes.ListType{
				ElemType: types.StringType,
			},
			"pii": basetypes.ListType{
				ElemType: AiGatewayGuardrailPiiBehavior{}.ToAttrType(ctx),
			},
			"safety": types.BoolType,
			"valid_topics": basetypes.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

type AiGatewayGuardrailPiiBehavior struct {
	// Behavior for PII filter. Currently only 'BLOCK' is supported. If 'BLOCK'
	// is set for the input guardrail and the request contains PII, the request
	// is not sent to the model server and 400 status code is returned; if
	// 'BLOCK' is set for the output guardrail and the model response contains
	// PII, the PII info in the response is redacted and 400 status code is
	// returned.
	Behavior types.String `tfsdk:"behavior" tf:""`
}

func (newState *AiGatewayGuardrailPiiBehavior) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrailPiiBehavior) {
}

func (newState *AiGatewayGuardrailPiiBehavior) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrailPiiBehavior) {
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

// ToAttrType returns the representation of AiGatewayGuardrailPiiBehavior in the Terraform plugin framework type
// system.
func (a AiGatewayGuardrailPiiBehavior) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"behavior": types.StringType,
		},
	}
}

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	Input types.List `tfsdk:"input" tf:"optional,object"`
	// Configuration for output guardrail filters.
	Output types.List `tfsdk:"output" tf:"optional,object"`
}

func (newState *AiGatewayGuardrails) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrails) {
}

func (newState *AiGatewayGuardrails) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrails) {
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

// ToAttrType returns the representation of AiGatewayGuardrails in the Terraform plugin framework type
// system.
func (a AiGatewayGuardrails) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"input": basetypes.ListType{
				ElemType: AiGatewayGuardrailParameters{}.ToAttrType(ctx),
			},
			"output": basetypes.ListType{
				ElemType: AiGatewayGuardrailParameters{}.ToAttrType(ctx),
			},
		},
	}
}

type AiGatewayInferenceTableConfig struct {
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

func (newState *AiGatewayInferenceTableConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayInferenceTableConfig) {
}

func (newState *AiGatewayInferenceTableConfig) SyncEffectiveFieldsDuringRead(existingState AiGatewayInferenceTableConfig) {
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

// ToAttrType returns the representation of AiGatewayInferenceTableConfig in the Terraform plugin framework type
// system.
func (a AiGatewayInferenceTableConfig) ToAttrType(ctx context.Context) types.ObjectType {
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
	Calls types.Int64 `tfsdk:"calls" tf:""`
	// Key field for a rate limit. Currently, only 'user' and 'endpoint' are
	// supported, with 'endpoint' being the default if not specified.
	Key types.String `tfsdk:"key" tf:"optional"`
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	RenewalPeriod types.String `tfsdk:"renewal_period" tf:""`
}

func (newState *AiGatewayRateLimit) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayRateLimit) {
}

func (newState *AiGatewayRateLimit) SyncEffectiveFieldsDuringRead(existingState AiGatewayRateLimit) {
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

// ToAttrType returns the representation of AiGatewayRateLimit in the Terraform plugin framework type
// system.
func (a AiGatewayRateLimit) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"renewal_period": types.StringType,
		},
	}
}

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
}

func (newState *AiGatewayUsageTrackingConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayUsageTrackingConfig) {
}

func (newState *AiGatewayUsageTrackingConfig) SyncEffectiveFieldsDuringRead(existingState AiGatewayUsageTrackingConfig) {
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

// ToAttrType returns the representation of AiGatewayUsageTrackingConfig in the Terraform plugin framework type
// system.
func (a AiGatewayUsageTrackingConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

type AmazonBedrockConfig struct {
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

func (newState *AmazonBedrockConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AmazonBedrockConfig) {
}

func (newState *AmazonBedrockConfig) SyncEffectiveFieldsDuringRead(existingState AmazonBedrockConfig) {
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

// ToAttrType returns the representation of AmazonBedrockConfig in the Terraform plugin framework type
// system.
func (a AmazonBedrockConfig) ToAttrType(ctx context.Context) types.ObjectType {
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

type AnthropicConfig struct {
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

func (newState *AnthropicConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AnthropicConfig) {
}

func (newState *AnthropicConfig) SyncEffectiveFieldsDuringRead(existingState AnthropicConfig) {
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

// ToAttrType returns the representation of AnthropicConfig in the Terraform plugin framework type
// system.
func (a AnthropicConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"anthropic_api_key":           types.StringType,
			"anthropic_api_key_plaintext": types.StringType,
		},
	}
}

type AutoCaptureConfigInput struct {
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

func (newState *AutoCaptureConfigInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureConfigInput) {
}

func (newState *AutoCaptureConfigInput) SyncEffectiveFieldsDuringRead(existingState AutoCaptureConfigInput) {
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

// ToAttrType returns the representation of AutoCaptureConfigInput in the Terraform plugin framework type
// system.
func (a AutoCaptureConfigInput) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *AutoCaptureConfigOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureConfigOutput) {
}

func (newState *AutoCaptureConfigOutput) SyncEffectiveFieldsDuringRead(existingState AutoCaptureConfigOutput) {
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

// ToAttrType returns the representation of AutoCaptureConfigOutput in the Terraform plugin framework type
// system.
func (a AutoCaptureConfigOutput) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"catalog_name": types.StringType,
			"enabled":      types.BoolType,
			"schema_name":  types.StringType,
			"state": basetypes.ListType{
				ElemType: AutoCaptureState{}.ToAttrType(ctx),
			},
			"table_name_prefix": types.StringType,
		},
	}
}

type AutoCaptureState struct {
	PayloadTable types.List `tfsdk:"payload_table" tf:"optional,object"`
}

func (newState *AutoCaptureState) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureState) {
}

func (newState *AutoCaptureState) SyncEffectiveFieldsDuringRead(existingState AutoCaptureState) {
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

// ToAttrType returns the representation of AutoCaptureState in the Terraform plugin framework type
// system.
func (a AutoCaptureState) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"payload_table": basetypes.ListType{
				ElemType: PayloadTable{}.ToAttrType(ctx),
			},
		},
	}
}

// Get build logs for a served model
type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (newState *BuildLogsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan BuildLogsRequest) {
}

func (newState *BuildLogsRequest) SyncEffectiveFieldsDuringRead(existingState BuildLogsRequest) {
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

// ToAttrType returns the representation of BuildLogsRequest in the Terraform plugin framework type
// system.
func (a BuildLogsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs types.String `tfsdk:"logs" tf:""`
}

func (newState *BuildLogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan BuildLogsResponse) {
}

func (newState *BuildLogsResponse) SyncEffectiveFieldsDuringRead(existingState BuildLogsResponse) {
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

// ToAttrType returns the representation of BuildLogsResponse in the Terraform plugin framework type
// system.
func (a BuildLogsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ChatMessage struct {
	// The content of the message.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The role of the message. One of [system, user, assistant].
	Role types.String `tfsdk:"role" tf:"optional"`
}

func (newState *ChatMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChatMessage) {
}

func (newState *ChatMessage) SyncEffectiveFieldsDuringRead(existingState ChatMessage) {
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

// ToAttrType returns the representation of ChatMessage in the Terraform plugin framework type
// system.
func (a ChatMessage) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *CohereConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan CohereConfig) {
}

func (newState *CohereConfig) SyncEffectiveFieldsDuringRead(existingState CohereConfig) {
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

// ToAttrType returns the representation of CohereConfig in the Terraform plugin framework type
// system.
func (a CohereConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cohere_api_base":          types.StringType,
			"cohere_api_key":           types.StringType,
			"cohere_api_key_plaintext": types.StringType,
		},
	}
}

type CreateServingEndpoint struct {
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

func (newState *CreateServingEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServingEndpoint) {
}

func (newState *CreateServingEndpoint) SyncEffectiveFieldsDuringRead(existingState CreateServingEndpoint) {
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
		"ai_gateway":  reflect.TypeOf(AiGatewayConfig{}),
		"config":      reflect.TypeOf(EndpointCoreConfigInput{}),
		"rate_limits": reflect.TypeOf(RateLimit{}),
		"tags":        reflect.TypeOf(EndpointTag{}),
	}
}

// ToAttrType returns the representation of CreateServingEndpoint in the Terraform plugin framework type
// system.
func (a CreateServingEndpoint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig{}.ToAttrType(ctx),
			},
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigInput{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.ToAttrType(ctx),
			},
			"route_optimized": types.BoolType,
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.ToAttrType(ctx),
			},
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

func (newState *DatabricksModelServingConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan DatabricksModelServingConfig) {
}

func (newState *DatabricksModelServingConfig) SyncEffectiveFieldsDuringRead(existingState DatabricksModelServingConfig) {
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

// ToAttrType returns the representation of DatabricksModelServingConfig in the Terraform plugin framework type
// system.
func (a DatabricksModelServingConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"databricks_api_token":           types.StringType,
			"databricks_api_token_plaintext": types.StringType,
			"databricks_workspace_url":       types.StringType,
		},
	}
}

type DataframeSplitInput struct {
	Columns types.List `tfsdk:"columns" tf:"optional"`

	Data types.List `tfsdk:"data" tf:"optional"`

	Index types.List `tfsdk:"index" tf:"optional"`
}

func (newState *DataframeSplitInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataframeSplitInput) {
}

func (newState *DataframeSplitInput) SyncEffectiveFieldsDuringRead(existingState DataframeSplitInput) {
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
		"columns": reflect.TypeOf(struct{}{}),
		"data":    reflect.TypeOf(struct{}{}),
		"index":   reflect.TypeOf(types.Int64{}),
	}
}

// ToAttrType returns the representation of DataframeSplitInput in the Terraform plugin framework type
// system.
func (a DataframeSplitInput) ToAttrType(ctx context.Context) types.ObjectType {
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

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
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

// ToAttrType returns the representation of DeleteResponse in the Terraform plugin framework type
// system.
func (a DeleteResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *DeleteServingEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteServingEndpointRequest) {
}

func (newState *DeleteServingEndpointRequest) SyncEffectiveFieldsDuringRead(existingState DeleteServingEndpointRequest) {
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

// ToAttrType returns the representation of DeleteServingEndpointRequest in the Terraform plugin framework type
// system.
func (a DeleteServingEndpointRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	Embedding types.List `tfsdk:"embedding" tf:"optional"`
	// The index of the embedding in the response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// This will always be 'embedding'.
	Object types.String `tfsdk:"object" tf:"optional"`
}

func (newState *EmbeddingsV1ResponseEmbeddingElement) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingsV1ResponseEmbeddingElement) {
}

func (newState *EmbeddingsV1ResponseEmbeddingElement) SyncEffectiveFieldsDuringRead(existingState EmbeddingsV1ResponseEmbeddingElement) {
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

// ToAttrType returns the representation of EmbeddingsV1ResponseEmbeddingElement in the Terraform plugin framework type
// system.
func (a EmbeddingsV1ResponseEmbeddingElement) ToAttrType(ctx context.Context) types.ObjectType {
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

type EndpointCoreConfigInput struct {
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

func (newState *EndpointCoreConfigInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigInput) {
}

func (newState *EndpointCoreConfigInput) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigInput) {
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

// ToAttrType returns the representation of EndpointCoreConfigInput in the Terraform plugin framework type
// system.
func (a EndpointCoreConfigInput) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigInput{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityInput{}.ToAttrType(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelInput{}.ToAttrType(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type EndpointCoreConfigOutput struct {
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

func (newState *EndpointCoreConfigOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigOutput) {
}

func (newState *EndpointCoreConfigOutput) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigOutput) {
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

// ToAttrType returns the representation of EndpointCoreConfigOutput in the Terraform plugin framework type
// system.
func (a EndpointCoreConfigOutput) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigOutput{}.ToAttrType(ctx),
			},
			"config_version": types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput{}.ToAttrType(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput{}.ToAttrType(ctx),
			},
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities types.List `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels types.List `tfsdk:"served_models" tf:"optional"`
}

func (newState *EndpointCoreConfigSummary) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigSummary) {
}

func (newState *EndpointCoreConfigSummary) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigSummary) {
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

// ToAttrType returns the representation of EndpointCoreConfigSummary in the Terraform plugin framework type
// system.
func (a EndpointCoreConfigSummary) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"served_entities": basetypes.ListType{
				ElemType: ServedEntitySpec{}.ToAttrType(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelSpec{}.ToAttrType(ctx),
			},
		},
	}
}

type EndpointPendingConfig struct {
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

func (newState *EndpointPendingConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointPendingConfig) {
}

func (newState *EndpointPendingConfig) SyncEffectiveFieldsDuringRead(existingState EndpointPendingConfig) {
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

// ToAttrType returns the representation of EndpointPendingConfig in the Terraform plugin framework type
// system.
func (a EndpointPendingConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"auto_capture_config": basetypes.ListType{
				ElemType: AutoCaptureConfigOutput{}.ToAttrType(ctx),
			},
			"config_version": types.Int64Type,
			"served_entities": basetypes.ListType{
				ElemType: ServedEntityOutput{}.ToAttrType(ctx),
			},
			"served_models": basetypes.ListType{
				ElemType: ServedModelOutput{}.ToAttrType(ctx),
			},
			"start_time": types.Int64Type,
			"traffic_config": basetypes.ListType{
				ElemType: TrafficConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type EndpointState struct {
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

func (newState *EndpointState) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointState) {
}

func (newState *EndpointState) SyncEffectiveFieldsDuringRead(existingState EndpointState) {
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

// ToAttrType returns the representation of EndpointState in the Terraform plugin framework type
// system.
func (a EndpointState) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config_update": types.StringType,
			"ready":         types.StringType,
		},
	}
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key types.String `tfsdk:"key" tf:""`
	// Optional value field for a serving endpoint tag.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *EndpointTag) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTag) {
}

func (newState *EndpointTag) SyncEffectiveFieldsDuringRead(existingState EndpointTag) {
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

// ToAttrType returns the representation of EndpointTag in the Terraform plugin framework type
// system.
func (a EndpointTag) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}
}

// Get metrics of a serving endpoint
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (newState *ExportMetricsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportMetricsRequest) {
}

func (newState *ExportMetricsRequest) SyncEffectiveFieldsDuringRead(existingState ExportMetricsRequest) {
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

// ToAttrType returns the representation of ExportMetricsRequest in the Terraform plugin framework type
// system.
func (a ExportMetricsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

type ExportMetricsResponse struct {
	Contents io.ReadCloser `tfsdk:"-"`
}

func (newState *ExportMetricsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportMetricsResponse) {
}

func (newState *ExportMetricsResponse) SyncEffectiveFieldsDuringRead(existingState ExportMetricsResponse) {
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

// ToAttrType returns the representation of ExportMetricsResponse in the Terraform plugin framework type
// system.
func (a ExportMetricsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"contents": types.ObjectType{},
		},
	}
}

type ExternalModel struct {
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

func (newState *ExternalModel) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalModel) {
}

func (newState *ExternalModel) SyncEffectiveFieldsDuringRead(existingState ExternalModel) {
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
		"databricks_model_serving_config": reflect.TypeOf(DatabricksModelServingConfig{}),
		"google_cloud_vertex_ai_config":   reflect.TypeOf(GoogleCloudVertexAiConfig{}),
		"openai_config":                   reflect.TypeOf(OpenAiConfig{}),
		"palm_config":                     reflect.TypeOf(PaLmConfig{}),
	}
}

// ToAttrType returns the representation of ExternalModel in the Terraform plugin framework type
// system.
func (a ExternalModel) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai21labs_config": basetypes.ListType{
				ElemType: Ai21LabsConfig{}.ToAttrType(ctx),
			},
			"amazon_bedrock_config": basetypes.ListType{
				ElemType: AmazonBedrockConfig{}.ToAttrType(ctx),
			},
			"anthropic_config": basetypes.ListType{
				ElemType: AnthropicConfig{}.ToAttrType(ctx),
			},
			"cohere_config": basetypes.ListType{
				ElemType: CohereConfig{}.ToAttrType(ctx),
			},
			"databricks_model_serving_config": basetypes.ListType{
				ElemType: DatabricksModelServingConfig{}.ToAttrType(ctx),
			},
			"google_cloud_vertex_ai_config": basetypes.ListType{
				ElemType: GoogleCloudVertexAiConfig{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"openai_config": basetypes.ListType{
				ElemType: OpenAiConfig{}.ToAttrType(ctx),
			},
			"palm_config": basetypes.ListType{
				ElemType: PaLmConfig{}.ToAttrType(ctx),
			},
			"provider": types.StringType,
			"task":     types.StringType,
		},
	}
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens types.Int64 `tfsdk:"completion_tokens" tf:"optional"`
	// The number of tokens in the prompt.
	PromptTokens types.Int64 `tfsdk:"prompt_tokens" tf:"optional"`
	// The total number of tokens in the prompt and response.
	TotalTokens types.Int64 `tfsdk:"total_tokens" tf:"optional"`
}

func (newState *ExternalModelUsageElement) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalModelUsageElement) {
}

func (newState *ExternalModelUsageElement) SyncEffectiveFieldsDuringRead(existingState ExternalModelUsageElement) {
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

// ToAttrType returns the representation of ExternalModelUsageElement in the Terraform plugin framework type
// system.
func (a ExternalModelUsageElement) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"completion_tokens": types.Int64Type,
			"prompt_tokens":     types.Int64Type,
			"total_tokens":      types.Int64Type,
		},
	}
}

type FoundationModel struct {
	// The description of the foundation model.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The display name of the foundation model.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The URL to the documentation of the foundation model.
	Docs types.String `tfsdk:"docs" tf:"optional"`
	// The name of the foundation model.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *FoundationModel) SyncEffectiveFieldsDuringCreateOrUpdate(plan FoundationModel) {
}

func (newState *FoundationModel) SyncEffectiveFieldsDuringRead(existingState FoundationModel) {
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

// ToAttrType returns the representation of FoundationModel in the Terraform plugin framework type
// system.
func (a FoundationModel) ToAttrType(ctx context.Context) types.ObjectType {
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
type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *GetOpenApiRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOpenApiRequest) {
}

func (newState *GetOpenApiRequest) SyncEffectiveFieldsDuringRead(existingState GetOpenApiRequest) {
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

// ToAttrType returns the representation of GetOpenApiRequest in the Terraform plugin framework type
// system.
func (a GetOpenApiRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
		},
	}
}

// The response is an OpenAPI spec in JSON format that typically includes fields
// like openapi, info, servers and paths, etc.
type GetOpenApiResponse struct {
}

func (newState *GetOpenApiResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOpenApiResponse) {
}

func (newState *GetOpenApiResponse) SyncEffectiveFieldsDuringRead(existingState GetOpenApiResponse) {
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

// ToAttrType returns the representation of GetOpenApiResponse in the Terraform plugin framework type
// system.
func (a GetOpenApiResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointPermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionLevelsRequest) {
}

func (newState *GetServingEndpointPermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionLevelsRequest) {
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

// ToAttrType returns the representation of GetServingEndpointPermissionLevelsRequest in the Terraform plugin framework type
// system.
func (a GetServingEndpointPermissionLevelsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetServingEndpointPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionLevelsResponse) {
}

func (newState *GetServingEndpointPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionLevelsResponse) {
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

// ToAttrType returns the representation of GetServingEndpointPermissionLevelsResponse in the Terraform plugin framework type
// system.
func (a GetServingEndpointPermissionLevelsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"permission_levels": basetypes.ListType{
				ElemType: ServingEndpointPermissionsDescription{}.ToAttrType(ctx),
			},
		},
	}
}

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionsRequest) {
}

func (newState *GetServingEndpointPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionsRequest) {
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

// ToAttrType returns the representation of GetServingEndpointPermissionsRequest in the Terraform plugin framework type
// system.
func (a GetServingEndpointPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"serving_endpoint_id": types.StringType,
		},
	}
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointRequest) {
}

func (newState *GetServingEndpointRequest) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointRequest) {
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

// ToAttrType returns the representation of GetServingEndpointRequest in the Terraform plugin framework type
// system.
func (a GetServingEndpointRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *GoogleCloudVertexAiConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan GoogleCloudVertexAiConfig) {
}

func (newState *GoogleCloudVertexAiConfig) SyncEffectiveFieldsDuringRead(existingState GoogleCloudVertexAiConfig) {
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

// ToAttrType returns the representation of GoogleCloudVertexAiConfig in the Terraform plugin framework type
// system.
func (a GoogleCloudVertexAiConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"private_key":           types.StringType,
			"private_key_plaintext": types.StringType,
			"project_id":            types.StringType,
			"region":                types.StringType,
		},
	}
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints types.List `tfsdk:"endpoints" tf:"optional"`
}

func (newState *ListEndpointsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListEndpointsResponse) {
}

func (newState *ListEndpointsResponse) SyncEffectiveFieldsDuringRead(existingState ListEndpointsResponse) {
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

// ToAttrType returns the representation of ListEndpointsResponse in the Terraform plugin framework type
// system.
func (a ListEndpointsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"endpoints": basetypes.ListType{
				ElemType: ServingEndpoint{}.ToAttrType(ctx),
			},
		},
	}
}

// Get the latest logs for a served model
type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName types.String `tfsdk:"-"`
}

func (newState *LogsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan LogsRequest) {
}

func (newState *LogsRequest) SyncEffectiveFieldsDuringRead(existingState LogsRequest) {
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

// ToAttrType returns the representation of LogsRequest in the Terraform plugin framework type
// system.
func (a LogsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":              types.StringType,
			"served_model_name": types.StringType,
		},
	}
}

type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo oauth2.DataPlaneInfo `tfsdk:"query_info" tf:"optional,object"`
}

func (newState *ModelDataPlaneInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelDataPlaneInfo) {
}

func (newState *ModelDataPlaneInfo) SyncEffectiveFieldsDuringRead(existingState ModelDataPlaneInfo) {
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
		"query_info": reflect.TypeOf(oauth2.DataPlaneInfo{}),
	}
}

// ToAttrType returns the representation of ModelDataPlaneInfo in the Terraform plugin framework type
// system.
func (a ModelDataPlaneInfo) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query_info": basetypes.ListType{
				ElemType: oauth2_tf.DataPlaneInfo{}.ToAttrType(ctx),
			},
		},
	}
}

type OpenAiConfig struct {
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

func (newState *OpenAiConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan OpenAiConfig) {
}

func (newState *OpenAiConfig) SyncEffectiveFieldsDuringRead(existingState OpenAiConfig) {
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

// ToAttrType returns the representation of OpenAiConfig in the Terraform plugin framework type
// system.
func (a OpenAiConfig) ToAttrType(ctx context.Context) types.ObjectType {
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
	PalmApiKey types.String `tfsdk:"palm_api_key" tf:"optional"`
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKeyPlaintext types.String `tfsdk:"palm_api_key_plaintext" tf:"optional"`
}

func (newState *PaLmConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan PaLmConfig) {
}

func (newState *PaLmConfig) SyncEffectiveFieldsDuringRead(existingState PaLmConfig) {
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

// ToAttrType returns the representation of PaLmConfig in the Terraform plugin framework type
// system.
func (a PaLmConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"palm_api_key":           types.StringType,
			"palm_api_key_plaintext": types.StringType,
		},
	}
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags types.List `tfsdk:"add_tags" tf:"optional"`
	// List of tag keys to delete
	DeleteTags types.List `tfsdk:"delete_tags" tf:"optional"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (newState *PatchServingEndpointTags) SyncEffectiveFieldsDuringCreateOrUpdate(plan PatchServingEndpointTags) {
}

func (newState *PatchServingEndpointTags) SyncEffectiveFieldsDuringRead(existingState PatchServingEndpointTags) {
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

// ToAttrType returns the representation of PatchServingEndpointTags in the Terraform plugin framework type
// system.
func (a PatchServingEndpointTags) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"add_tags": basetypes.ListType{
				ElemType: EndpointTag{}.ToAttrType(ctx),
			},
			"delete_tags": basetypes.ListType{
				ElemType: types.StringType,
			},
			"name": types.StringType,
		},
	}
}

type PayloadTable struct {
	// The name of the payload table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The status of the payload table.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The status message of the payload table.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
}

func (newState *PayloadTable) SyncEffectiveFieldsDuringCreateOrUpdate(plan PayloadTable) {
}

func (newState *PayloadTable) SyncEffectiveFieldsDuringRead(existingState PayloadTable) {
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

// ToAttrType returns the representation of PayloadTable in the Terraform plugin framework type
// system.
func (a PayloadTable) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name":           types.StringType,
			"status":         types.StringType,
			"status_message": types.StringType,
		},
	}
}

// Update AI Gateway of a serving endpoint
type PutAiGatewayRequest struct {
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

func (newState *PutAiGatewayRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAiGatewayRequest) {
}

func (newState *PutAiGatewayRequest) SyncEffectiveFieldsDuringRead(existingState PutAiGatewayRequest) {
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
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig{}),
	}
}

// ToAttrType returns the representation of PutAiGatewayRequest in the Terraform plugin framework type
// system.
func (a PutAiGatewayRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails{}.ToAttrType(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig{}.ToAttrType(ctx),
			},
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.ToAttrType(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig{}.ToAttrType(ctx),
			},
		},
	}
}

type PutAiGatewayResponse struct {
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

func (newState *PutAiGatewayResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAiGatewayResponse) {
}

func (newState *PutAiGatewayResponse) SyncEffectiveFieldsDuringRead(existingState PutAiGatewayResponse) {
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
		"guardrails":             reflect.TypeOf(AiGatewayGuardrails{}),
		"inference_table_config": reflect.TypeOf(AiGatewayInferenceTableConfig{}),
		"rate_limits":            reflect.TypeOf(AiGatewayRateLimit{}),
		"usage_tracking_config":  reflect.TypeOf(AiGatewayUsageTrackingConfig{}),
	}
}

// ToAttrType returns the representation of PutAiGatewayResponse in the Terraform plugin framework type
// system.
func (a PutAiGatewayResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"guardrails": basetypes.ListType{
				ElemType: AiGatewayGuardrails{}.ToAttrType(ctx),
			},
			"inference_table_config": basetypes.ListType{
				ElemType: AiGatewayInferenceTableConfig{}.ToAttrType(ctx),
			},
			"rate_limits": basetypes.ListType{
				ElemType: AiGatewayRateLimit{}.ToAttrType(ctx),
			},
			"usage_tracking_config": basetypes.ListType{
				ElemType: AiGatewayUsageTrackingConfig{}.ToAttrType(ctx),
			},
		},
	}
}

// Update rate limits of a serving endpoint
type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
}

func (newState *PutRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutRequest) {
}

func (newState *PutRequest) SyncEffectiveFieldsDuringRead(existingState PutRequest) {
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

// ToAttrType returns the representation of PutRequest in the Terraform plugin framework type
// system.
func (a PutRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"name": types.StringType,
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.ToAttrType(ctx),
			},
		},
	}
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits types.List `tfsdk:"rate_limits" tf:"optional"`
}

func (newState *PutResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutResponse) {
}

func (newState *PutResponse) SyncEffectiveFieldsDuringRead(existingState PutResponse) {
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

// ToAttrType returns the representation of PutResponse in the Terraform plugin framework type
// system.
func (a PutResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rate_limits": basetypes.ListType{
				ElemType: RateLimit{}.ToAttrType(ctx),
			},
		},
	}
}

type QueryEndpointInput struct {
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
	Input any `tfsdk:"input" tf:"optional"`
	// Tensor-based input in columnar format.
	Inputs any `tfsdk:"inputs" tf:"optional"`
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
	Prompt any `tfsdk:"prompt" tf:"optional"`
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

func (newState *QueryEndpointInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEndpointInput) {
}

func (newState *QueryEndpointInput) SyncEffectiveFieldsDuringRead(existingState QueryEndpointInput) {
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
		"dataframe_records": reflect.TypeOf(struct{}{}),
		"dataframe_split":   reflect.TypeOf(DataframeSplitInput{}),
		"extra_params":      reflect.TypeOf(types.String{}),
		"instances":         reflect.TypeOf(struct{}{}),
		"messages":          reflect.TypeOf(ChatMessage{}),
		"stop":              reflect.TypeOf(types.String{}),
	}
}

// ToAttrType returns the representation of QueryEndpointInput in the Terraform plugin framework type
// system.
func (a QueryEndpointInput) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dataframe_records": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"dataframe_split": basetypes.ListType{
				ElemType: DataframeSplitInput{}.ToAttrType(ctx),
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
				ElemType: ChatMessage{}.ToAttrType(ctx),
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

type QueryEndpointResponse struct {
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

func (newState *QueryEndpointResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEndpointResponse) {
}

func (newState *QueryEndpointResponse) SyncEffectiveFieldsDuringRead(existingState QueryEndpointResponse) {
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
		"predictions": reflect.TypeOf(struct{}{}),
		"usage":       reflect.TypeOf(ExternalModelUsageElement{}),
	}
}

// ToAttrType returns the representation of QueryEndpointResponse in the Terraform plugin framework type
// system.
func (a QueryEndpointResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"choices": basetypes.ListType{
				ElemType: V1ResponseChoiceElement{}.ToAttrType(ctx),
			},
			"created": types.Int64Type,
			"data": basetypes.ListType{
				ElemType: EmbeddingsV1ResponseEmbeddingElement{}.ToAttrType(ctx),
			},
			"id":     types.StringType,
			"model":  types.StringType,
			"object": types.StringType,
			"predictions": basetypes.ListType{
				ElemType: types.ObjectType{},
			},
			"served-model-name": types.StringType,
			"usage": basetypes.ListType{
				ElemType: ExternalModelUsageElement{}.ToAttrType(ctx),
			},
		},
	}
}

type RateLimit struct {
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

func (newState *RateLimit) SyncEffectiveFieldsDuringCreateOrUpdate(plan RateLimit) {
}

func (newState *RateLimit) SyncEffectiveFieldsDuringRead(existingState RateLimit) {
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

// ToAttrType returns the representation of RateLimit in the Terraform plugin framework type
// system.
func (a RateLimit) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"calls":          types.Int64Type,
			"key":            types.StringType,
			"renewal_period": types.StringType,
		},
	}
}

type Route struct {
	// The name of the served model this route configures traffic for.
	ServedModelName types.String `tfsdk:"served_model_name" tf:""`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage types.Int64 `tfsdk:"traffic_percentage" tf:""`
}

func (newState *Route) SyncEffectiveFieldsDuringCreateOrUpdate(plan Route) {
}

func (newState *Route) SyncEffectiveFieldsDuringRead(existingState Route) {
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

// ToAttrType returns the representation of Route in the Terraform plugin framework type
// system.
func (a Route) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
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

func (newState *ServedEntityInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntityInput) {
}

func (newState *ServedEntityInput) SyncEffectiveFieldsDuringRead(existingState ServedEntityInput) {
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

// ToAttrType returns the representation of ServedEntityInput in the Terraform plugin framework type
// system.
func (a ServedEntityInput) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"environment_vars": basetypes.MapType{
				ElemType: types.StringType,
			},
			"external_model": basetypes.ListType{
				ElemType: ExternalModel{}.ToAttrType(ctx),
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

type ServedEntityOutput struct {
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

func (newState *ServedEntityOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntityOutput) {
}

func (newState *ServedEntityOutput) SyncEffectiveFieldsDuringRead(existingState ServedEntityOutput) {
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

// ToAttrType returns the representation of ServedEntityOutput in the Terraform plugin framework type
// system.
func (a ServedEntityOutput) ToAttrType(ctx context.Context) types.ObjectType {
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
				ElemType: ExternalModel{}.ToAttrType(ctx),
			},
			"foundation_model": basetypes.ListType{
				ElemType: FoundationModel{}.ToAttrType(ctx),
			},
			"instance_profile_arn":       types.StringType,
			"max_provisioned_throughput": types.Int64Type,
			"min_provisioned_throughput": types.Int64Type,
			"name":                       types.StringType,
			"scale_to_zero_enabled":      types.BoolType,
			"state": basetypes.ListType{
				ElemType: ServedModelState{}.ToAttrType(ctx),
			},
			"workload_size": types.StringType,
			"workload_type": types.StringType,
		},
	}
}

type ServedEntitySpec struct {
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

func (newState *ServedEntitySpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntitySpec) {
}

func (newState *ServedEntitySpec) SyncEffectiveFieldsDuringRead(existingState ServedEntitySpec) {
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

// ToAttrType returns the representation of ServedEntitySpec in the Terraform plugin framework type
// system.
func (a ServedEntitySpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"entity_name":    types.StringType,
			"entity_version": types.StringType,
			"external_model": basetypes.ListType{
				ElemType: ExternalModel{}.ToAttrType(ctx),
			},
			"foundation_model": basetypes.ListType{
				ElemType: FoundationModel{}.ToAttrType(ctx),
			},
			"name": types.StringType,
		},
	}
}

type ServedModelInput struct {
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

func (newState *ServedModelInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelInput) {
}

func (newState *ServedModelInput) SyncEffectiveFieldsDuringRead(existingState ServedModelInput) {
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

// ToAttrType returns the representation of ServedModelInput in the Terraform plugin framework type
// system.
func (a ServedModelInput) ToAttrType(ctx context.Context) types.ObjectType {
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

type ServedModelOutput struct {
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

func (newState *ServedModelOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelOutput) {
}

func (newState *ServedModelOutput) SyncEffectiveFieldsDuringRead(existingState ServedModelOutput) {
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

// ToAttrType returns the representation of ServedModelOutput in the Terraform plugin framework type
// system.
func (a ServedModelOutput) ToAttrType(ctx context.Context) types.ObjectType {
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
				ElemType: ServedModelState{}.ToAttrType(ctx),
			},
			"workload_size": types.StringType,
			"workload_type": types.StringType,
		},
	}
}

type ServedModelSpec struct {
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName types.String `tfsdk:"model_name" tf:"optional"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion types.String `tfsdk:"model_version" tf:"optional"`
	// The name of the served model.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ServedModelSpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelSpec) {
}

func (newState *ServedModelSpec) SyncEffectiveFieldsDuringRead(existingState ServedModelSpec) {
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

// ToAttrType returns the representation of ServedModelSpec in the Terraform plugin framework type
// system.
func (a ServedModelSpec) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"model_name":    types.StringType,
			"model_version": types.StringType,
			"name":          types.StringType,
		},
	}
}

type ServedModelState struct {
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

func (newState *ServedModelState) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedModelState) {
}

func (newState *ServedModelState) SyncEffectiveFieldsDuringRead(existingState ServedModelState) {
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

// ToAttrType returns the representation of ServedModelState in the Terraform plugin framework type
// system.
func (a ServedModelState) ToAttrType(ctx context.Context) types.ObjectType {
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
	Logs types.String `tfsdk:"logs" tf:""`
}

func (newState *ServerLogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServerLogsResponse) {
}

func (newState *ServerLogsResponse) SyncEffectiveFieldsDuringRead(existingState ServerLogsResponse) {
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

// ToAttrType returns the representation of ServerLogsResponse in the Terraform plugin framework type
// system.
func (a ServerLogsResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logs": types.StringType,
		},
	}
}

type ServingEndpoint struct {
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

func (newState *ServingEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpoint) {
}

func (newState *ServingEndpoint) SyncEffectiveFieldsDuringRead(existingState ServingEndpoint) {
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

// ToAttrType returns the representation of ServingEndpoint in the Terraform plugin framework type
// system.
func (a ServingEndpoint) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig{}.ToAttrType(ctx),
			},
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigSummary{}.ToAttrType(ctx),
			},
			"creation_timestamp":     types.Int64Type,
			"creator":                types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"state": basetypes.ListType{
				ElemType: EndpointState{}.ToAttrType(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.ToAttrType(ctx),
			},
			"task": types.StringType,
		},
	}
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *ServingEndpointAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointAccessControlRequest) {
}

func (newState *ServingEndpointAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState ServingEndpointAccessControlRequest) {
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

// ToAttrType returns the representation of ServingEndpointAccessControlRequest in the Terraform plugin framework type
// system.
func (a ServingEndpointAccessControlRequest) ToAttrType(ctx context.Context) types.ObjectType {
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

func (newState *ServingEndpointAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointAccessControlResponse) {
}

func (newState *ServingEndpointAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState ServingEndpointAccessControlResponse) {
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

// ToAttrType returns the representation of ServingEndpointAccessControlResponse in the Terraform plugin framework type
// system.
func (a ServingEndpointAccessControlResponse) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"all_permissions": basetypes.ListType{
				ElemType: ServingEndpointPermission{}.ToAttrType(ctx),
			},
			"display_name":           types.StringType,
			"group_name":             types.StringType,
			"service_principal_name": types.StringType,
			"user_name":              types.StringType,
		},
	}
}

type ServingEndpointDetailed struct {
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

func (newState *ServingEndpointDetailed) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointDetailed) {
}

func (newState *ServingEndpointDetailed) SyncEffectiveFieldsDuringRead(existingState ServingEndpointDetailed) {
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
		"ai_gateway":      reflect.TypeOf(AiGatewayConfig{}),
		"config":          reflect.TypeOf(EndpointCoreConfigOutput{}),
		"data_plane_info": reflect.TypeOf(ModelDataPlaneInfo{}),
		"pending_config":  reflect.TypeOf(EndpointPendingConfig{}),
		"state":           reflect.TypeOf(EndpointState{}),
		"tags":            reflect.TypeOf(EndpointTag{}),
	}
}

// ToAttrType returns the representation of ServingEndpointDetailed in the Terraform plugin framework type
// system.
func (a ServingEndpointDetailed) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ai_gateway": basetypes.ListType{
				ElemType: AiGatewayConfig{}.ToAttrType(ctx),
			},
			"config": basetypes.ListType{
				ElemType: EndpointCoreConfigOutput{}.ToAttrType(ctx),
			},
			"creation_timestamp": types.Int64Type,
			"creator":            types.StringType,
			"data_plane_info": basetypes.ListType{
				ElemType: ModelDataPlaneInfo{}.ToAttrType(ctx),
			},
			"endpoint_url":           types.StringType,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"name":                   types.StringType,
			"pending_config": basetypes.ListType{
				ElemType: EndpointPendingConfig{}.ToAttrType(ctx),
			},
			"permission_level": types.StringType,
			"route_optimized":  types.BoolType,
			"state": basetypes.ListType{
				ElemType: EndpointState{}.ToAttrType(ctx),
			},
			"tags": basetypes.ListType{
				ElemType: EndpointTag{}.ToAttrType(ctx),
			},
			"task": types.StringType,
		},
	}
}

type ServingEndpointPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ServingEndpointPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermission) {
}

func (newState *ServingEndpointPermission) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermission) {
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

// ToAttrType returns the representation of ServingEndpointPermission in the Terraform plugin framework type
// system.
func (a ServingEndpointPermission) ToAttrType(ctx context.Context) types.ObjectType {
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

type ServingEndpointPermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ServingEndpointPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissions) {
}

func (newState *ServingEndpointPermissions) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissions) {
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

// ToAttrType returns the representation of ServingEndpointPermissions in the Terraform plugin framework type
// system.
func (a ServingEndpointPermissions) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlResponse{}.ToAttrType(ctx),
			},
			"object_id":   types.StringType,
			"object_type": types.StringType,
		},
	}
}

type ServingEndpointPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ServingEndpointPermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissionsDescription) {
}

func (newState *ServingEndpointPermissionsDescription) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissionsDescription) {
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

// ToAttrType returns the representation of ServingEndpointPermissionsDescription in the Terraform plugin framework type
// system.
func (a ServingEndpointPermissionsDescription) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":      types.StringType,
			"permission_level": types.StringType,
		},
	}
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *ServingEndpointPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissionsRequest) {
}

func (newState *ServingEndpointPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissionsRequest) {
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

// ToAttrType returns the representation of ServingEndpointPermissionsRequest in the Terraform plugin framework type
// system.
func (a ServingEndpointPermissionsRequest) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"access_control_list": basetypes.ListType{
				ElemType: ServingEndpointAccessControlRequest{}.ToAttrType(ctx),
			},
			"serving_endpoint_id": types.StringType,
		},
	}
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes types.List `tfsdk:"routes" tf:"optional"`
}

func (newState *TrafficConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrafficConfig) {
}

func (newState *TrafficConfig) SyncEffectiveFieldsDuringRead(existingState TrafficConfig) {
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

// ToAttrType returns the representation of TrafficConfig in the Terraform plugin framework type
// system.
func (a TrafficConfig) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"routes": basetypes.ListType{
				ElemType: Route{}.ToAttrType(ctx),
			},
		},
	}
}

type V1ResponseChoiceElement struct {
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

func (newState *V1ResponseChoiceElement) SyncEffectiveFieldsDuringCreateOrUpdate(plan V1ResponseChoiceElement) {
}

func (newState *V1ResponseChoiceElement) SyncEffectiveFieldsDuringRead(existingState V1ResponseChoiceElement) {
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

// ToAttrType returns the representation of V1ResponseChoiceElement in the Terraform plugin framework type
// system.
func (a V1ResponseChoiceElement) ToAttrType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"finishReason": types.StringType,
			"index":        types.Int64Type,
			"logprobs":     types.Int64Type,
			"message": basetypes.ListType{
				ElemType: ChatMessage{}.ToAttrType(ctx),
			},
			"text": types.StringType,
		},
	}
}
