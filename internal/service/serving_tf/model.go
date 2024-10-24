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
	"io"

	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type AiGatewayConfig struct {
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails []AiGatewayGuardrails `tfsdk:"guardrails" tf:"optional,object"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig []AiGatewayInferenceTableConfig `tfsdk:"inference_table_config" tf:"optional,object"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit `tfsdk:"rate_limits" tf:"optional"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig []AiGatewayUsageTrackingConfig `tfsdk:"usage_tracking_config" tf:"optional,object"`
}

func (newState *AiGatewayConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayConfig) {
}

func (newState *AiGatewayConfig) SyncEffectiveFieldsDuringRead(existingState AiGatewayConfig) {
}

type AiGatewayGuardrailParameters struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	InvalidKeywords []types.String `tfsdk:"invalid_keywords" tf:"optional"`
	// Configuration for guardrail PII filter.
	Pii []AiGatewayGuardrailPiiBehavior `tfsdk:"pii" tf:"optional,object"`
	// Indicates whether the safety filter is enabled.
	Safety types.Bool `tfsdk:"safety" tf:"optional"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics []types.String `tfsdk:"valid_topics" tf:"optional"`
}

func (newState *AiGatewayGuardrailParameters) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrailParameters) {
}

func (newState *AiGatewayGuardrailParameters) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrailParameters) {
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

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	Input []AiGatewayGuardrailParameters `tfsdk:"input" tf:"optional,object"`
	// Configuration for output guardrail filters.
	Output []AiGatewayGuardrailParameters `tfsdk:"output" tf:"optional,object"`
}

func (newState *AiGatewayGuardrails) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayGuardrails) {
}

func (newState *AiGatewayGuardrails) SyncEffectiveFieldsDuringRead(existingState AiGatewayGuardrails) {
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

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
}

func (newState *AiGatewayUsageTrackingConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan AiGatewayUsageTrackingConfig) {
}

func (newState *AiGatewayUsageTrackingConfig) SyncEffectiveFieldsDuringRead(existingState AiGatewayUsageTrackingConfig) {
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

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the schema in Unity Catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`

	State []AutoCaptureState `tfsdk:"state" tf:"optional,object"`
	// The prefix of the table in Unity Catalog.
	TableNamePrefix types.String `tfsdk:"table_name_prefix" tf:"optional"`
}

func (newState *AutoCaptureConfigOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureConfigOutput) {
}

func (newState *AutoCaptureConfigOutput) SyncEffectiveFieldsDuringRead(existingState AutoCaptureConfigOutput) {
}

type AutoCaptureState struct {
	PayloadTable []PayloadTable `tfsdk:"payload_table" tf:"optional,object"`
}

func (newState *AutoCaptureState) SyncEffectiveFieldsDuringCreateOrUpdate(plan AutoCaptureState) {
}

func (newState *AutoCaptureState) SyncEffectiveFieldsDuringRead(existingState AutoCaptureState) {
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

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs types.String `tfsdk:"logs" tf:""`
}

func (newState *BuildLogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan BuildLogsResponse) {
}

func (newState *BuildLogsResponse) SyncEffectiveFieldsDuringRead(existingState BuildLogsResponse) {
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

type CreateServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: only
	// external model endpoints are supported as of now.
	AiGateway []AiGatewayConfig `tfsdk:"ai_gateway" tf:"optional,object"`
	// The core config of the serving endpoint.
	Config []EndpointCoreConfigInput `tfsdk:"config" tf:"object"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name types.String `tfsdk:"name" tf:""`
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits" tf:"optional"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized types.Bool `tfsdk:"route_optimized" tf:"optional"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `tfsdk:"tags" tf:"optional"`
}

func (newState *CreateServingEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateServingEndpoint) {
}

func (newState *CreateServingEndpoint) SyncEffectiveFieldsDuringRead(existingState CreateServingEndpoint) {
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

type DataframeSplitInput struct {
	Columns []any `tfsdk:"columns" tf:"optional"`

	Data []any `tfsdk:"data" tf:"optional"`

	Index []types.Int64 `tfsdk:"index" tf:"optional"`
}

func (newState *DataframeSplitInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataframeSplitInput) {
}

func (newState *DataframeSplitInput) SyncEffectiveFieldsDuringRead(existingState DataframeSplitInput) {
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
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

type EmbeddingsV1ResponseEmbeddingElement struct {
	Embedding []types.Float64 `tfsdk:"embedding" tf:"optional"`
	// The index of the embedding in the response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// This will always be 'embedding'.
	Object types.String `tfsdk:"object" tf:"optional"`
}

func (newState *EmbeddingsV1ResponseEmbeddingElement) SyncEffectiveFieldsDuringCreateOrUpdate(plan EmbeddingsV1ResponseEmbeddingElement) {
}

func (newState *EmbeddingsV1ResponseEmbeddingElement) SyncEffectiveFieldsDuringRead(existingState EmbeddingsV1ResponseEmbeddingElement) {
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig []AutoCaptureConfigInput `tfsdk:"auto_capture_config" tf:"optional,object"`
	// The name of the serving endpoint to update. This field is required.
	Name types.String `tfsdk:"-"`
	// A list of served entities for the endpoint to serve. A serving endpoint
	// can have up to 15 served entities.
	ServedEntities []ServedEntityInput `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) A list of served models for the
	// endpoint to serve. A serving endpoint can have up to 15 served models.
	ServedModels []ServedModelInput `tfsdk:"served_models" tf:"optional"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig []TrafficConfig `tfsdk:"traffic_config" tf:"optional,object"`
}

func (newState *EndpointCoreConfigInput) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigInput) {
}

func (newState *EndpointCoreConfigInput) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigInput) {
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig []AutoCaptureConfigOutput `tfsdk:"auto_capture_config" tf:"optional,object"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version" tf:"optional"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityOutput `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelOutput `tfsdk:"served_models" tf:"optional"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig []TrafficConfig `tfsdk:"traffic_config" tf:"optional,object"`
}

func (newState *EndpointCoreConfigOutput) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigOutput) {
}

func (newState *EndpointCoreConfigOutput) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigOutput) {
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntitySpec `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelSpec `tfsdk:"served_models" tf:"optional"`
}

func (newState *EndpointCoreConfigSummary) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointCoreConfigSummary) {
}

func (newState *EndpointCoreConfigSummary) SyncEffectiveFieldsDuringRead(existingState EndpointCoreConfigSummary) {
}

type EndpointPendingConfig struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig []AutoCaptureConfigOutput `tfsdk:"auto_capture_config" tf:"optional,object"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version" tf:"optional"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities []ServedEntityOutput `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels []ServedModelOutput `tfsdk:"served_models" tf:"optional"`
	// The timestamp when the update to the pending config started.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig []TrafficConfig `tfsdk:"traffic_config" tf:"optional,object"`
}

func (newState *EndpointPendingConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointPendingConfig) {
}

func (newState *EndpointPendingConfig) SyncEffectiveFieldsDuringRead(existingState EndpointPendingConfig) {
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

type ExportMetricsResponse struct {
	Contents io.ReadCloser `tfsdk:"-"`
}

func (newState *ExportMetricsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExportMetricsResponse) {
}

func (newState *ExportMetricsResponse) SyncEffectiveFieldsDuringRead(existingState ExportMetricsResponse) {
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig []Ai21LabsConfig `tfsdk:"ai21labs_config" tf:"optional,object"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig []AmazonBedrockConfig `tfsdk:"amazon_bedrock_config" tf:"optional,object"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig []AnthropicConfig `tfsdk:"anthropic_config" tf:"optional,object"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig []CohereConfig `tfsdk:"cohere_config" tf:"optional,object"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig []DatabricksModelServingConfig `tfsdk:"databricks_model_serving_config" tf:"optional,object"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig []GoogleCloudVertexAiConfig `tfsdk:"google_cloud_vertex_ai_config" tf:"optional,object"`
	// The name of the external model.
	Name types.String `tfsdk:"name" tf:""`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig []OpenAiConfig `tfsdk:"openai_config" tf:"optional,object"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig []PaLmConfig `tfsdk:"palm_config" tf:"optional,object"`
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

// The response is an OpenAPI spec in JSON format that typically includes fields
// like openapi, info, servers and paths, etc.
type GetOpenApiResponse struct {
}

func (newState *GetOpenApiResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetOpenApiResponse) {
}

func (newState *GetOpenApiResponse) SyncEffectiveFieldsDuringRead(existingState GetOpenApiResponse) {
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

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetServingEndpointPermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointPermissionLevelsResponse) {
}

func (newState *GetServingEndpointPermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointPermissionLevelsResponse) {
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

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-"`
}

func (newState *GetServingEndpointRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetServingEndpointRequest) {
}

func (newState *GetServingEndpointRequest) SyncEffectiveFieldsDuringRead(existingState GetServingEndpointRequest) {
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

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint `tfsdk:"endpoints" tf:"optional"`
}

func (newState *ListEndpointsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListEndpointsResponse) {
}

func (newState *ListEndpointsResponse) SyncEffectiveFieldsDuringRead(existingState ListEndpointsResponse) {
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

type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo oauth2.DataPlaneInfo `tfsdk:"query_info" tf:"optional,object"`
}

func (newState *ModelDataPlaneInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ModelDataPlaneInfo) {
}

func (newState *ModelDataPlaneInfo) SyncEffectiveFieldsDuringRead(existingState ModelDataPlaneInfo) {
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

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag `tfsdk:"add_tags" tf:"optional"`
	// List of tag keys to delete
	DeleteTags []types.String `tfsdk:"delete_tags" tf:"optional"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name types.String `tfsdk:"-"`
}

func (newState *PatchServingEndpointTags) SyncEffectiveFieldsDuringCreateOrUpdate(plan PatchServingEndpointTags) {
}

func (newState *PatchServingEndpointTags) SyncEffectiveFieldsDuringRead(existingState PatchServingEndpointTags) {
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

// Update AI Gateway of a serving endpoint
type PutAiGatewayRequest struct {
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails []AiGatewayGuardrails `tfsdk:"guardrails" tf:"optional,object"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig []AiGatewayInferenceTableConfig `tfsdk:"inference_table_config" tf:"optional,object"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name types.String `tfsdk:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit `tfsdk:"rate_limits" tf:"optional"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig []AiGatewayUsageTrackingConfig `tfsdk:"usage_tracking_config" tf:"optional,object"`
}

func (newState *PutAiGatewayRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAiGatewayRequest) {
}

func (newState *PutAiGatewayRequest) SyncEffectiveFieldsDuringRead(existingState PutAiGatewayRequest) {
}

type PutAiGatewayResponse struct {
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails []AiGatewayGuardrails `tfsdk:"guardrails" tf:"optional,object"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality .
	InferenceTableConfig []AiGatewayInferenceTableConfig `tfsdk:"inference_table_config" tf:"optional,object"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit `tfsdk:"rate_limits" tf:"optional"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig []AiGatewayUsageTrackingConfig `tfsdk:"usage_tracking_config" tf:"optional,object"`
}

func (newState *PutAiGatewayResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutAiGatewayResponse) {
}

func (newState *PutAiGatewayResponse) SyncEffectiveFieldsDuringRead(existingState PutAiGatewayResponse) {
}

// Update rate limits of a serving endpoint
type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name types.String `tfsdk:"-"`
	// The list of endpoint rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits" tf:"optional"`
}

func (newState *PutRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutRequest) {
}

func (newState *PutRequest) SyncEffectiveFieldsDuringRead(existingState PutRequest) {
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits" tf:"optional"`
}

func (newState *PutResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan PutResponse) {
}

func (newState *PutResponse) SyncEffectiveFieldsDuringRead(existingState PutResponse) {
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `tfsdk:"dataframe_records" tf:"optional"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit []DataframeSplitInput `tfsdk:"dataframe_split" tf:"optional,object"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams map[string]types.String `tfsdk:"extra_params" tf:"optional"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input any `tfsdk:"input" tf:"optional"`
	// Tensor-based input in columnar format.
	Inputs any `tfsdk:"inputs" tf:"optional"`
	// Tensor-based input in row format.
	Instances []any `tfsdk:"instances" tf:"optional"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens types.Int64 `tfsdk:"max_tokens" tf:"optional"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	Messages []ChatMessage `tfsdk:"messages" tf:"optional"`
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
	Stop []types.String `tfsdk:"stop" tf:"optional"`
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

type QueryEndpointResponse struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices []V1ResponseChoiceElement `tfsdk:"choices" tf:"optional"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created types.Int64 `tfsdk:"created" tf:"optional"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data []EmbeddingsV1ResponseEmbeddingElement `tfsdk:"data" tf:"optional"`
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
	Predictions []any `tfsdk:"predictions" tf:"optional"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName types.String `tfsdk:"-"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage []ExternalModelUsageElement `tfsdk:"usage" tf:"optional,object"`
}

func (newState *QueryEndpointResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEndpointResponse) {
}

func (newState *QueryEndpointResponse) SyncEffectiveFieldsDuringRead(existingState QueryEndpointResponse) {
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
	EnvironmentVars map[string]types.String `tfsdk:"environment_vars" tf:"optional"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel []ExternalModel `tfsdk:"external_model" tf:"optional,object"`
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
	EnvironmentVars map[string]types.String `tfsdk:"environment_vars" tf:"optional"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	ExternalModel []ExternalModel `tfsdk:"external_model" tf:"optional,object"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	FoundationModel []FoundationModel `tfsdk:"foundation_model" tf:"optional,object"`
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
	State []ServedModelState `tfsdk:"state" tf:"optional,object"`
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
	ExternalModel []ExternalModel `tfsdk:"external_model" tf:"optional,object"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	FoundationModel []FoundationModel `tfsdk:"foundation_model" tf:"optional,object"`
	// The name of the served entity.
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ServedEntitySpec) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServedEntitySpec) {
}

func (newState *ServedEntitySpec) SyncEffectiveFieldsDuringRead(existingState ServedEntitySpec) {
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]types.String `tfsdk:"environment_vars" tf:"optional"`
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
	EnvironmentVars map[string]types.String `tfsdk:"environment_vars" tf:"optional"`
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
	State []ServedModelState `tfsdk:"state" tf:"optional,object"`
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

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs types.String `tfsdk:"logs" tf:""`
}

func (newState *ServerLogsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServerLogsResponse) {
}

func (newState *ServerLogsResponse) SyncEffectiveFieldsDuringRead(existingState ServerLogsResponse) {
}

type ServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: Only
	// external model endpoints are currently supported.
	AiGateway []AiGatewayConfig `tfsdk:"ai_gateway" tf:"optional,object"`
	// The config that is currently being served by the endpoint.
	Config []EndpointCoreConfigSummary `tfsdk:"config" tf:"optional,object"`
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
	State []EndpointState `tfsdk:"state" tf:"optional,object"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `tfsdk:"tags" tf:"optional"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task" tf:"optional"`
}

func (newState *ServingEndpoint) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpoint) {
}

func (newState *ServingEndpoint) SyncEffectiveFieldsDuringRead(existingState ServingEndpoint) {
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

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []ServingEndpointPermission `tfsdk:"all_permissions" tf:"optional"`
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

type ServingEndpointDetailed struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: Only
	// external model endpoints are currently supported.
	AiGateway []AiGatewayConfig `tfsdk:"ai_gateway" tf:"optional,object"`
	// The config that is currently being served by the endpoint.
	Config []EndpointCoreConfigOutput `tfsdk:"config" tf:"optional,object"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo []ModelDataPlaneInfo `tfsdk:"data_plane_info" tf:"optional,object"`
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
	PendingConfig []EndpointPendingConfig `tfsdk:"pending_config" tf:"optional,object"`
	// The permission level of the principal making the request.
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized types.Bool `tfsdk:"route_optimized" tf:"optional"`
	// Information corresponding to the state of the serving endpoint.
	State []EndpointState `tfsdk:"state" tf:"optional,object"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `tfsdk:"tags" tf:"optional"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task" tf:"optional"`
}

func (newState *ServingEndpointDetailed) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointDetailed) {
}

func (newState *ServingEndpointDetailed) SyncEffectiveFieldsDuringRead(existingState ServingEndpointDetailed) {
}

type ServingEndpointPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *ServingEndpointPermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermission) {
}

func (newState *ServingEndpointPermission) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermission) {
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *ServingEndpointPermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissions) {
}

func (newState *ServingEndpointPermissions) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissions) {
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

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-"`
}

func (newState *ServingEndpointPermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServingEndpointPermissionsRequest) {
}

func (newState *ServingEndpointPermissionsRequest) SyncEffectiveFieldsDuringRead(existingState ServingEndpointPermissionsRequest) {
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes []Route `tfsdk:"routes" tf:"optional"`
}

func (newState *TrafficConfig) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrafficConfig) {
}

func (newState *TrafficConfig) SyncEffectiveFieldsDuringRead(existingState TrafficConfig) {
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason types.String `tfsdk:"finishReason" tf:"optional"`
	// The index of the choice in the __chat or completions__ response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs types.Int64 `tfsdk:"logprobs" tf:"optional"`
	// The message response from the __chat__ endpoint.
	Message []ChatMessage `tfsdk:"message" tf:"optional,object"`
	// The text response from the __completions__ endpoint.
	Text types.String `tfsdk:"text" tf:"optional"`
}

func (newState *V1ResponseChoiceElement) SyncEffectiveFieldsDuringCreateOrUpdate(plan V1ResponseChoiceElement) {
}

func (newState *V1ResponseChoiceElement) SyncEffectiveFieldsDuringRead(existingState V1ResponseChoiceElement) {
}
