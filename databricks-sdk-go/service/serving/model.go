// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKey string `json:"ai21labs_api_key,omitempty"`
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKeyPlaintext string `json:"ai21labs_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Ai21LabsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Ai21LabsConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiGatewayConfig struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func (s *AiGatewayConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiGatewayGuardrailParameters struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	InvalidKeywords []string `json:"invalid_keywords,omitempty"`
	// Configuration for guardrail PII filter.
	Pii *AiGatewayGuardrailPiiBehavior `json:"pii,omitempty"`
	// Indicates whether the safety filter is enabled.
	Safety bool `json:"safety,omitempty"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics []string `json:"valid_topics,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiGatewayGuardrailParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayGuardrailParameters) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiGatewayGuardrailPiiBehavior struct {
	// Configuration for input guardrail filters.
	Behavior AiGatewayGuardrailPiiBehaviorBehavior `json:"behavior,omitempty"`
}

func (s *AiGatewayGuardrailPiiBehavior) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiGatewayGuardrailPiiBehaviorBehavior string

const AiGatewayGuardrailPiiBehaviorBehaviorBlock AiGatewayGuardrailPiiBehaviorBehavior = `BLOCK`

const AiGatewayGuardrailPiiBehaviorBehaviorMask AiGatewayGuardrailPiiBehaviorBehavior = `MASK`

const AiGatewayGuardrailPiiBehaviorBehaviorNone AiGatewayGuardrailPiiBehaviorBehavior = `NONE`

// String representation for [fmt.Print]
func (f *AiGatewayGuardrailPiiBehaviorBehavior) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Set(v string) error {
	switch v {
	case `BLOCK`, `MASK`, `NONE`:
		*f = AiGatewayGuardrailPiiBehaviorBehavior(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK", "MASK", "NONE"`, v)
	}
}

// Values returns all possible values for AiGatewayGuardrailPiiBehaviorBehavior.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Values() []AiGatewayGuardrailPiiBehaviorBehavior {
	return []AiGatewayGuardrailPiiBehaviorBehavior{
		AiGatewayGuardrailPiiBehaviorBehaviorBlock,
		AiGatewayGuardrailPiiBehaviorBehaviorMask,
		AiGatewayGuardrailPiiBehaviorBehaviorNone,
	}
}

// Type always returns AiGatewayGuardrailPiiBehaviorBehavior to satisfy [pflag.Value] interface
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Type() string {
	return "AiGatewayGuardrailPiiBehaviorBehavior"
}

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	Input *AiGatewayGuardrailParameters `json:"input,omitempty"`
	// Configuration for output guardrail filters.
	Output *AiGatewayGuardrailParameters `json:"output,omitempty"`
}

func (s *AiGatewayGuardrails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AiGatewayInferenceTableConfig struct {
	// The name of the catalog in Unity Catalog. Required when enabling
	// inference tables. NOTE: On update, you have to disable inference table
	// first in order to change the catalog name.
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiGatewayInferenceTableConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayInferenceTableConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiGatewayRateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int64 `json:"calls,omitempty"`
	// Key field for a rate limit. Currently, 'user', 'user_group,
	// 'service_principal', and 'endpoint' are supported, with 'endpoint' being
	// the default if not specified.
	Key AiGatewayRateLimitKey `json:"key,omitempty"`
	// Principal field for a user, user group, or service principal to apply
	// rate limiting to. Accepts a user email, group name, or service principal
	// application ID.
	Principal string `json:"principal,omitempty"`
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	RenewalPeriod AiGatewayRateLimitRenewalPeriod `json:"renewal_period"`
	// Used to specify how many tokens are allowed for a key within the
	// renewal_period.
	Tokens int64 `json:"tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiGatewayRateLimit) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayRateLimit) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AiGatewayRateLimitKey string

const AiGatewayRateLimitKeyEndpoint AiGatewayRateLimitKey = `endpoint`

const AiGatewayRateLimitKeyServicePrincipal AiGatewayRateLimitKey = `service_principal`

const AiGatewayRateLimitKeyUser AiGatewayRateLimitKey = `user`

const AiGatewayRateLimitKeyUserGroup AiGatewayRateLimitKey = `user_group`

// String representation for [fmt.Print]
func (f *AiGatewayRateLimitKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayRateLimitKey) Set(v string) error {
	switch v {
	case `endpoint`, `service_principal`, `user`, `user_group`:
		*f = AiGatewayRateLimitKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "endpoint", "service_principal", "user", "user_group"`, v)
	}
}

// Values returns all possible values for AiGatewayRateLimitKey.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiGatewayRateLimitKey) Values() []AiGatewayRateLimitKey {
	return []AiGatewayRateLimitKey{
		AiGatewayRateLimitKeyEndpoint,
		AiGatewayRateLimitKeyServicePrincipal,
		AiGatewayRateLimitKeyUser,
		AiGatewayRateLimitKeyUserGroup,
	}
}

// Type always returns AiGatewayRateLimitKey to satisfy [pflag.Value] interface
func (f *AiGatewayRateLimitKey) Type() string {
	return "AiGatewayRateLimitKey"
}

type AiGatewayRateLimitRenewalPeriod string

const AiGatewayRateLimitRenewalPeriodMinute AiGatewayRateLimitRenewalPeriod = `minute`

// String representation for [fmt.Print]
func (f *AiGatewayRateLimitRenewalPeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayRateLimitRenewalPeriod) Set(v string) error {
	switch v {
	case `minute`:
		*f = AiGatewayRateLimitRenewalPeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "minute"`, v)
	}
}

// Values returns all possible values for AiGatewayRateLimitRenewalPeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiGatewayRateLimitRenewalPeriod) Values() []AiGatewayRateLimitRenewalPeriod {
	return []AiGatewayRateLimitRenewalPeriod{
		AiGatewayRateLimitRenewalPeriodMinute,
	}
}

// Type always returns AiGatewayRateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *AiGatewayRateLimitRenewalPeriod) Type() string {
	return "AiGatewayRateLimitRenewalPeriod"
}

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AiGatewayUsageTrackingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayUsageTrackingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AmazonBedrockConfig struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id_plaintext`. You must
	// provide an API key using one of the following fields: `aws_access_key_id`
	// or `aws_access_key_id_plaintext`.
	AwsAccessKeyId string `json:"aws_access_key_id,omitempty"`
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	AwsAccessKeyIdPlaintext string `json:"aws_access_key_id_plaintext,omitempty"`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion string `json:"aws_region"`
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKey string `json:"aws_secret_access_key,omitempty"`
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKeyPlaintext string `json:"aws_secret_access_key_plaintext,omitempty"`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider AmazonBedrockConfigBedrockProvider `json:"bedrock_provider"`
	// ARN of the instance profile that the external model will use to access
	// AWS resources. You must authenticate using an instance profile or access
	// keys. If you prefer to authenticate using access keys, see
	// `aws_access_key_id`, `aws_access_key_id_plaintext`,
	// `aws_secret_access_key` and `aws_secret_access_key_plaintext`.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The name of the Unity Catalog service credential that the external model
	// uses to access AWS resources.
	UcServiceCredentialName string `json:"uc_service_credential_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AmazonBedrockConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AmazonBedrockConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AmazonBedrockConfigBedrockProvider string

const AmazonBedrockConfigBedrockProviderAi21labs AmazonBedrockConfigBedrockProvider = `ai21labs`

const AmazonBedrockConfigBedrockProviderAmazon AmazonBedrockConfigBedrockProvider = `amazon`

const AmazonBedrockConfigBedrockProviderAnthropic AmazonBedrockConfigBedrockProvider = `anthropic`

const AmazonBedrockConfigBedrockProviderCohere AmazonBedrockConfigBedrockProvider = `cohere`

// String representation for [fmt.Print]
func (f *AmazonBedrockConfigBedrockProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AmazonBedrockConfigBedrockProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `amazon`, `anthropic`, `cohere`:
		*f = AmazonBedrockConfigBedrockProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "amazon", "anthropic", "cohere"`, v)
	}
}

// Values returns all possible values for AmazonBedrockConfigBedrockProvider.
//
// There is no guarantee on the order of the values in the slice.
func (f *AmazonBedrockConfigBedrockProvider) Values() []AmazonBedrockConfigBedrockProvider {
	return []AmazonBedrockConfigBedrockProvider{
		AmazonBedrockConfigBedrockProviderAi21labs,
		AmazonBedrockConfigBedrockProviderAmazon,
		AmazonBedrockConfigBedrockProviderAnthropic,
		AmazonBedrockConfigBedrockProviderCohere,
	}
}

// Type always returns AmazonBedrockConfigBedrockProvider to satisfy [pflag.Value] interface
func (f *AmazonBedrockConfigBedrockProvider) Type() string {
	return "AmazonBedrockConfigBedrockProvider"
}

type AnthropicConfig struct {
	// The Databricks secret key reference for an Anthropic API key. If you
	// prefer to paste your API key directly, see `anthropic_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKey string `json:"anthropic_api_key,omitempty"`
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKeyPlaintext string `json:"anthropic_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AnthropicConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AnthropicConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ApiKeyAuth struct {
	// The name of the API key parameter used for authentication.
	Key string `json:"key"`
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	Value string `json:"value,omitempty"`
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	ValuePlaintext string `json:"value_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ApiKeyAuth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApiKeyAuth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Deprecated: legacy inference table configuration. Please use AI Gateway
// inference tables instead. See
// https://docs.databricks.com/aws/en/ai-gateway/inference-tables.
type AutoCaptureConfigInput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutoCaptureConfigInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Deprecated: legacy inference table configuration. Please use AI Gateway
// inference tables instead. See
// https://docs.databricks.com/aws/en/ai-gateway/inference-tables.
type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName string `json:"schema_name,omitempty"`
	// The state of the inference table.
	State *AutoCaptureState `json:"state,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutoCaptureConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureState struct {
	// The payload table that stores the logged requests and responses.
	PayloadTable *PayloadTable `json:"payload_table,omitempty"`
}

func (s *AutoCaptureState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BearerTokenAuth struct {
	// The Databricks secret key reference for a token. If you prefer to paste
	// your token directly, see `token_plaintext`.
	Token string `json:"token,omitempty"`
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	TokenPlaintext string `json:"token_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BearerTokenAuth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BearerTokenAuth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string `json:"-" url:"-"`
}

func (s *BuildLogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs string `json:"logs"`
}

func (s *BuildLogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ChatMessage struct {
	// The content of the message.
	Content string `json:"content,omitempty"`
	// The role of the message. One of [system, user, assistant].
	Role ChatMessageRole `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ChatMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ChatMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The role of the message. One of [system, user, assistant].
type ChatMessageRole string

const ChatMessageRoleAssistant ChatMessageRole = `assistant`

const ChatMessageRoleSystem ChatMessageRole = `system`

const ChatMessageRoleUser ChatMessageRole = `user`

// String representation for [fmt.Print]
func (f *ChatMessageRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ChatMessageRole) Set(v string) error {
	switch v {
	case `assistant`, `system`, `user`:
		*f = ChatMessageRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "assistant", "system", "user"`, v)
	}
}

// Values returns all possible values for ChatMessageRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *ChatMessageRole) Values() []ChatMessageRole {
	return []ChatMessageRole{
		ChatMessageRoleAssistant,
		ChatMessageRoleSystem,
		ChatMessageRoleUser,
	}
}

// Type always returns ChatMessageRole to satisfy [pflag.Value] interface
func (f *ChatMessageRole) Type() string {
	return "ChatMessageRole"
}

type CohereConfig struct {
	// This is an optional field to provide a customized base URL for the Cohere
	// API. If left unspecified, the standard Cohere base URL is used.
	CohereApiBase string `json:"cohere_api_base,omitempty"`
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	CohereApiKey string `json:"cohere_api_key,omitempty"`
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	CohereApiKeyPlaintext string `json:"cohere_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CohereConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CohereConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The commitment term length for a Provisioned Throughput pool.
type CommitmentTerm string

const CommitmentTermCommitmentTerm12Hours CommitmentTerm = `COMMITMENT_TERM_12_HOURS`

const CommitmentTermCommitmentTerm1Day CommitmentTerm = `COMMITMENT_TERM_1_DAY`

const CommitmentTermCommitmentTerm1Hour CommitmentTerm = `COMMITMENT_TERM_1_HOUR`

const CommitmentTermCommitmentTerm1Month CommitmentTerm = `COMMITMENT_TERM_1_MONTH`

const CommitmentTermCommitmentTerm1Week CommitmentTerm = `COMMITMENT_TERM_1_WEEK`

const CommitmentTermCommitmentTerm1Year CommitmentTerm = `COMMITMENT_TERM_1_YEAR`

const CommitmentTermCommitmentTerm3Months CommitmentTerm = `COMMITMENT_TERM_3_MONTHS`

const CommitmentTermCommitmentTerm6Months CommitmentTerm = `COMMITMENT_TERM_6_MONTHS`

// String representation for [fmt.Print]
func (f *CommitmentTerm) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommitmentTerm) Set(v string) error {
	switch v {
	case `COMMITMENT_TERM_12_HOURS`, `COMMITMENT_TERM_1_DAY`, `COMMITMENT_TERM_1_HOUR`, `COMMITMENT_TERM_1_MONTH`, `COMMITMENT_TERM_1_WEEK`, `COMMITMENT_TERM_1_YEAR`, `COMMITMENT_TERM_3_MONTHS`, `COMMITMENT_TERM_6_MONTHS`:
		*f = CommitmentTerm(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMITMENT_TERM_12_HOURS", "COMMITMENT_TERM_1_DAY", "COMMITMENT_TERM_1_HOUR", "COMMITMENT_TERM_1_MONTH", "COMMITMENT_TERM_1_WEEK", "COMMITMENT_TERM_1_YEAR", "COMMITMENT_TERM_3_MONTHS", "COMMITMENT_TERM_6_MONTHS"`, v)
	}
}

// Values returns all possible values for CommitmentTerm.
//
// There is no guarantee on the order of the values in the slice.
func (f *CommitmentTerm) Values() []CommitmentTerm {
	return []CommitmentTerm{
		CommitmentTermCommitmentTerm12Hours,
		CommitmentTermCommitmentTerm1Day,
		CommitmentTermCommitmentTerm1Hour,
		CommitmentTermCommitmentTerm1Month,
		CommitmentTermCommitmentTerm1Week,
		CommitmentTermCommitmentTerm1Year,
		CommitmentTermCommitmentTerm3Months,
		CommitmentTermCommitmentTerm6Months,
	}
}

// Type always returns CommitmentTerm to satisfy [pflag.Value] interface
func (f *CommitmentTerm) Type() string {
	return "CommitmentTerm"
}

type CreatePtEndpointRequest struct {
	// The AI Gateway configuration for the serving endpoint.
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The core config of the serving endpoint.
	Config PtEndpointCoreConfig `json:"config"`
	// Email notification settings.
	EmailNotifications *EmailNotifications `json:"email_notifications,omitempty"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `json:"name"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePtEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePtEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy to be applied to the serving endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The core config of the serving endpoint.
	Config *EndpointCoreConfigInput `json:"config,omitempty"`
	// The description of the serving endpoint.
	Description string `json:"description,omitempty"`
	// Email notification settings.
	EmailNotifications *EmailNotifications `json:"email_notifications,omitempty"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `json:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized bool `json:"route_optimized,omitempty"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `json:"tags,omitempty"`
	// Configuration for persisting endpoint telemetry (logs, traces, and
	// metrics) to Unity Catalog tables.
	TelemetryConfig *TelemetryConfig `json:"telemetry_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configs needed to create a custom provider model route.
type CustomProviderConfig struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	ApiKeyAuth *ApiKeyAuth `json:"api_key_auth,omitempty"`
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	BearerTokenAuth *BearerTokenAuth `json:"bearer_token_auth,omitempty"`
	// This is a field to provide the URL of the custom provider API.
	CustomProviderUrl string `json:"custom_provider_url"`
}

func (s *CustomProviderConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Details necessary to query this object's API through the DataPlane APIs.
type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails string `json:"authorization_details,omitempty"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl string `json:"endpoint_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataPlaneInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiToken string `json:"databricks_api_token,omitempty"`
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiTokenPlaintext string `json:"databricks_api_token_plaintext,omitempty"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl string `json:"databricks_workspace_url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DatabricksModelServingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksModelServingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataframeSplitInput struct {
	// Columns array for the dataframe
	Columns []any `json:"columns,omitempty"`
	// Data array for the dataframe
	Data []any `json:"data,omitempty"`
	// Index array for the dataframe
	Index []int `json:"index,omitempty"`
}

func (s *DataframeSplitInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint to delete.
	Name string `json:"-" url:"-"`
}

func (s *DeleteServingEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EmailNotifications struct {
	// A list of email addresses to be notified when an endpoint fails to update
	// its configuration or state.
	OnUpdateFailure []string `json:"on_update_failure,omitempty"`
	// A list of email addresses to be notified when an endpoint successfully
	// updates its configuration or state.
	OnUpdateSuccess []string `json:"on_update_success,omitempty"`
}

func (s *EmailNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	// The embedding vector
	Embedding []float64 `json:"embedding,omitempty"`
	// The index of the embedding in the response.
	Index int `json:"index,omitempty"`
	// This will always be 'embedding'.
	Object EmbeddingsV1ResponseEmbeddingElementObject `json:"object,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EmbeddingsV1ResponseEmbeddingElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingsV1ResponseEmbeddingElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This will always be 'embedding'.
type EmbeddingsV1ResponseEmbeddingElementObject string

const EmbeddingsV1ResponseEmbeddingElementObjectEmbedding EmbeddingsV1ResponseEmbeddingElementObject = `embedding`

// String representation for [fmt.Print]
func (f *EmbeddingsV1ResponseEmbeddingElementObject) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Set(v string) error {
	switch v {
	case `embedding`:
		*f = EmbeddingsV1ResponseEmbeddingElementObject(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "embedding"`, v)
	}
}

// Values returns all possible values for EmbeddingsV1ResponseEmbeddingElementObject.
//
// There is no guarantee on the order of the values in the slice.
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Values() []EmbeddingsV1ResponseEmbeddingElementObject {
	return []EmbeddingsV1ResponseEmbeddingElementObject{
		EmbeddingsV1ResponseEmbeddingElementObjectEmbedding,
	}
}

// Type always returns EmbeddingsV1ResponseEmbeddingElementObject to satisfy [pflag.Value] interface
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Type() string {
	return "EmbeddingsV1ResponseEmbeddingElementObject"
}

type EndpointCoreConfigInput struct {
	// Configuration for legacy Inference Tables which automatically log
	// requests and responses to Unity Catalog. Deprecated: please use AI
	// Gateway inference tables instead. See
	// https://docs.databricks.com/aws/en/ai-gateway/inference-tables.
	AutoCaptureConfig *AutoCaptureConfigInput `json:"auto_capture_config,omitempty"`
	// The name of the serving endpoint to update. This field is required.
	Name string `json:"-" url:"-"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityInput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelInput `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

func (s *EndpointCoreConfigInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EndpointCoreConfigOutput struct {
	// Configuration for legacy Inference Tables which automatically log
	// requests and responses to Unity Catalog. Deprecated: please use AI
	// Gateway inference tables instead. See
	// https://docs.databricks.com/aws/en/ai-gateway/inference-tables.
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int64 `json:"config_version,omitempty"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointCoreConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointCoreConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntitySpec `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelSpec `json:"served_models,omitempty"`
}

func (s *EndpointCoreConfigSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EndpointPendingConfig struct {
	// Configuration for legacy Inference Tables which automatically log
	// requests and responses to Unity Catalog. Deprecated: please use AI
	// Gateway inference tables instead. See
	// https://docs.databricks.com/aws/en/ai-gateway/inference-tables.
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The timestamp when the update to the pending config started, in
	// milliseconds since the Unix epoch.
	StartTime int64 `json:"start_time,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointPendingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointPendingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails.
	ConfigUpdate EndpointStateConfigUpdate `json:"config_update,omitempty"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady `json:"ready,omitempty"`
}

func (s *EndpointState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EndpointStateConfigUpdate string

const EndpointStateConfigUpdateInProgress EndpointStateConfigUpdate = `IN_PROGRESS`

const EndpointStateConfigUpdateNotUpdating EndpointStateConfigUpdate = `NOT_UPDATING`

const EndpointStateConfigUpdateUpdateCanceled EndpointStateConfigUpdate = `UPDATE_CANCELED`

const EndpointStateConfigUpdateUpdateFailed EndpointStateConfigUpdate = `UPDATE_FAILED`

// String representation for [fmt.Print]
func (f *EndpointStateConfigUpdate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateConfigUpdate) Set(v string) error {
	switch v {
	case `IN_PROGRESS`, `NOT_UPDATING`, `UPDATE_CANCELED`, `UPDATE_FAILED`:
		*f = EndpointStateConfigUpdate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN_PROGRESS", "NOT_UPDATING", "UPDATE_CANCELED", "UPDATE_FAILED"`, v)
	}
}

// Values returns all possible values for EndpointStateConfigUpdate.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStateConfigUpdate) Values() []EndpointStateConfigUpdate {
	return []EndpointStateConfigUpdate{
		EndpointStateConfigUpdateInProgress,
		EndpointStateConfigUpdateNotUpdating,
		EndpointStateConfigUpdateUpdateCanceled,
		EndpointStateConfigUpdateUpdateFailed,
	}
}

// Type always returns EndpointStateConfigUpdate to satisfy [pflag.Value] interface
func (f *EndpointStateConfigUpdate) Type() string {
	return "EndpointStateConfigUpdate"
}

type EndpointStateReady string

const EndpointStateReadyNotReady EndpointStateReady = `NOT_READY`

const EndpointStateReadyReady EndpointStateReady = `READY`

// String representation for [fmt.Print]
func (f *EndpointStateReady) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateReady) Set(v string) error {
	switch v {
	case `NOT_READY`, `READY`:
		*f = EndpointStateReady(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NOT_READY", "READY"`, v)
	}
}

// Values returns all possible values for EndpointStateReady.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStateReady) Values() []EndpointStateReady {
	return []EndpointStateReady{
		EndpointStateReadyNotReady,
		EndpointStateReadyReady,
	}
}

// Type always returns EndpointStateReady to satisfy [pflag.Value] interface
func (f *EndpointStateReady) Type() string {
	return "EndpointStateReady"
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key string `json:"key"`
	// Optional value field for a serving endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointTags struct {
	// The tags attached to the serving endpoint after the update.
	Tags []EndpointTag `json:"tags,omitempty"`
}

func (s *EndpointTags) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `json:"-" url:"-"`
}

func (s *ExportMetricsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ExportMetricsResponse struct {
	Contents io.ReadCloser `json:"-"`
}

func (s *ExportMetricsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Simple Proto message for testing
type ExternalFunctionRequest struct {
	// The connection name to use. This is required to identify the external
	// connection.
	ConnectionName string `json:"connection_name"`
	// Additional headers for the request. If not provided, only auth headers
	// from connections would be passed.
	Headers string `json:"headers,omitempty"`
	// The JSON payload to send in the request body.
	Json string `json:"json,omitempty"`
	// The HTTP method to use (e.g., 'GET', 'POST').
	Method ExternalFunctionRequestHttpMethod `json:"method"`
	// Query parameters for the request.
	Params string `json:"params,omitempty"`
	// The relative path for the API endpoint. This is required.
	Path string `json:"path"`
	// Optional subdomain to prepend to the connection URL's host. If provided,
	// this will be added as a prefix to the connection URL's host. For example,
	// if the connection URL is `https://api.example.com/v1` and `sub_domain` is
	// `"custom"`, the resulting URL will be
	// `https://custom.api.example.com/v1`.
	SubDomain string `json:"sub_domain,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExternalFunctionRequestHttpMethod string

const ExternalFunctionRequestHttpMethodDelete ExternalFunctionRequestHttpMethod = `DELETE`

const ExternalFunctionRequestHttpMethodGet ExternalFunctionRequestHttpMethod = `GET`

const ExternalFunctionRequestHttpMethodPatch ExternalFunctionRequestHttpMethod = `PATCH`

const ExternalFunctionRequestHttpMethodPost ExternalFunctionRequestHttpMethod = `POST`

const ExternalFunctionRequestHttpMethodPut ExternalFunctionRequestHttpMethod = `PUT`

// String representation for [fmt.Print]
func (f *ExternalFunctionRequestHttpMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExternalFunctionRequestHttpMethod) Set(v string) error {
	switch v {
	case `DELETE`, `GET`, `PATCH`, `POST`, `PUT`:
		*f = ExternalFunctionRequestHttpMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE", "GET", "PATCH", "POST", "PUT"`, v)
	}
}

// Values returns all possible values for ExternalFunctionRequestHttpMethod.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExternalFunctionRequestHttpMethod) Values() []ExternalFunctionRequestHttpMethod {
	return []ExternalFunctionRequestHttpMethod{
		ExternalFunctionRequestHttpMethodDelete,
		ExternalFunctionRequestHttpMethodGet,
		ExternalFunctionRequestHttpMethodPatch,
		ExternalFunctionRequestHttpMethodPost,
		ExternalFunctionRequestHttpMethodPut,
	}
}

// Type always returns ExternalFunctionRequestHttpMethod to satisfy [pflag.Value] interface
func (f *ExternalFunctionRequestHttpMethod) Type() string {
	return "ExternalFunctionRequestHttpMethod"
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig *Ai21LabsConfig `json:"ai21labs_config,omitempty"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig *AmazonBedrockConfig `json:"amazon_bedrock_config,omitempty"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig *AnthropicConfig `json:"anthropic_config,omitempty"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig *CohereConfig `json:"cohere_config,omitempty"`
	// Custom Provider Config. Only required if the provider is 'custom'.
	CustomProviderConfig *CustomProviderConfig `json:"custom_provider_config,omitempty"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig *DatabricksModelServingConfig `json:"databricks_model_serving_config,omitempty"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig *GoogleCloudVertexAiConfig `json:"google_cloud_vertex_ai_config,omitempty"`
	// The name of the external model.
	Name string `json:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig *OpenAiConfig `json:"openai_config,omitempty"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig *PaLmConfig `json:"palm_config,omitempty"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	Provider ExternalModelProvider `json:"provider"`
	// The task type of the external model.
	Task string `json:"task"`
}

func (s *ExternalModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ExternalModelProvider string

const ExternalModelProviderAi21labs ExternalModelProvider = `ai21labs`

const ExternalModelProviderAmazonBedrock ExternalModelProvider = `amazon-bedrock`

const ExternalModelProviderAnthropic ExternalModelProvider = `anthropic`

const ExternalModelProviderCohere ExternalModelProvider = `cohere`

const ExternalModelProviderCustom ExternalModelProvider = `custom`

const ExternalModelProviderDatabricksModelServing ExternalModelProvider = `databricks-model-serving`

const ExternalModelProviderGoogleCloudVertexAi ExternalModelProvider = `google-cloud-vertex-ai`

const ExternalModelProviderOpenai ExternalModelProvider = `openai`

const ExternalModelProviderPalm ExternalModelProvider = `palm`

// String representation for [fmt.Print]
func (f *ExternalModelProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExternalModelProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `amazon-bedrock`, `anthropic`, `cohere`, `custom`, `databricks-model-serving`, `google-cloud-vertex-ai`, `openai`, `palm`:
		*f = ExternalModelProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "amazon-bedrock", "anthropic", "cohere", "custom", "databricks-model-serving", "google-cloud-vertex-ai", "openai", "palm"`, v)
	}
}

// Values returns all possible values for ExternalModelProvider.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExternalModelProvider) Values() []ExternalModelProvider {
	return []ExternalModelProvider{
		ExternalModelProviderAi21labs,
		ExternalModelProviderAmazonBedrock,
		ExternalModelProviderAnthropic,
		ExternalModelProviderCohere,
		ExternalModelProviderCustom,
		ExternalModelProviderDatabricksModelServing,
		ExternalModelProviderGoogleCloudVertexAi,
		ExternalModelProviderOpenai,
		ExternalModelProviderPalm,
	}
}

// Type always returns ExternalModelProvider to satisfy [pflag.Value] interface
func (f *ExternalModelProvider) Type() string {
	return "ExternalModelProvider"
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// The number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens,omitempty"`
	// The total number of tokens in the prompt and response.
	TotalTokens int `json:"total_tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalModelUsageElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalModelUsageElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FallbackConfig struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	Enabled bool `json:"enabled"`
}

func (s *FallbackConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// All fields are not sensitive as they are hard-coded in the system and made
// available to customers.
type FoundationModel struct {
	// A description of the foundation model.
	Description string `json:"description,omitempty"`
	// The human-readable display name of the foundation model.
	DisplayName string `json:"display_name,omitempty"`
	// The URL of the Databricks documentation page for the foundation model.
	Docs string `json:"docs,omitempty"`
	// The name of the foundation model.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FoundationModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FoundationModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
}

func (s *GetOpenApiRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetOpenApiResponse struct {
	Contents io.ReadCloser `json:"-"`
}

func (s *GetOpenApiResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

func (s *GetServingEndpointPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetServingEndpointPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

func (s *GetServingEndpointPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

func (s *GetServingEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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
	PrivateKey string `json:"private_key,omitempty"`
	// The private key for the service account which has access to the Google
	// Cloud Vertex AI Service provided as a plaintext secret. See [Best
	// practices for managing service account keys]. If you prefer to reference
	// your key using Databricks Secrets, see `private_key`. You must provide an
	// API key using one of the following fields: `private_key` or
	// `private_key_plaintext`.
	//
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	PrivateKeyPlaintext string `json:"private_key_plaintext,omitempty"`
	// This is the Google Cloud project id that the service account is
	// associated with.
	ProjectId string `json:"project_id"`
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]:
	// https://cloud.google.com/vertex-ai/docs/general/locations
	Region string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GoogleCloudVertexAiConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GoogleCloudVertexAiConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type HttpRequestResponse struct {
	Contents io.ReadCloser `json:"-"`
}

func (s *HttpRequestResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint `json:"endpoints,omitempty"`
}

func (s *ListEndpointsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string `json:"-" url:"-"`
}

func (s *LogsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A representation of all DataPlaneInfo for operations that can be done on a
// model through Data Plane APIs.
type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo *DataPlaneInfo `json:"query_info,omitempty"`
}

func (s *ModelDataPlaneInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configs needed to create an OpenAI model route.
type OpenAiConfig struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	MicrosoftEntraClientId string `json:"microsoft_entra_client_id,omitempty"`
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecret string `json:"microsoft_entra_client_secret,omitempty"`
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecretPlaintext string `json:"microsoft_entra_client_secret_plaintext,omitempty"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId string `json:"microsoft_entra_tenant_id,omitempty"`
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	OpenaiApiBase string `json:"openai_api_base,omitempty"`
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKey string `json:"openai_api_key,omitempty"`
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKeyPlaintext string `json:"openai_api_key_plaintext,omitempty"`
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	OpenaiApiType string `json:"openai_api_type,omitempty"`
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	OpenaiApiVersion string `json:"openai_api_version,omitempty"`
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	OpenaiDeploymentName string `json:"openai_deployment_name,omitempty"`
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	OpenaiOrganization string `json:"openai_organization,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OpenAiConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OpenAiConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key. If you prefer to
	// paste your API key directly, see `palm_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKey string `json:"palm_api_key,omitempty"`
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKeyPlaintext string `json:"palm_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PaLmConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PaLmConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag `json:"add_tags,omitempty"`
	// List of tag keys to delete
	DeleteTags []string `json:"delete_tags,omitempty"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string `json:"-" url:"-"`
}

func (s *PatchServingEndpointTags) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Updates the telemetry configuration of a serving endpoint.
type PatchTelemetryConfigRequest struct {
	// The name of the serving endpoint whose telemetry configuration is being
	// updated. This field is required.
	Name string `json:"-" url:"-"`
	// The telemetry configuration to be applied to the serving endpoint. Can
	// specify either a telemetry_profile_id to use an existing profile, or
	// table_names to create a new profile with the specified Unity Catalog
	// tables. If not provided, the telemetry configuration will be removed from
	// the endpoint.
	TelemetryConfig *TelemetryConfig `json:"telemetry_config,omitempty"`
}

func (s *PatchTelemetryConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PayloadTable struct {
	// The name of the payload table within the configured catalog and schema.
	Name string `json:"name,omitempty"`
	// The status of the payload table. One of PENDING, READY, FAILED or
	// DISABLED.
	Status string `json:"status,omitempty"`
	// More information about the status of the payload table, if available.
	StatusMessage string `json:"status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PayloadTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PayloadTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PtEndpointCoreConfig struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []PtServedModel `json:"served_entities,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

func (s *PtEndpointCoreConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PtServedModel struct {
	// Whether burst scaling is enabled. When enabled (default), the endpoint
	// can automatically scale up beyond provisioned capacity to handle traffic
	// spikes. When disabled, the endpoint maintains fixed capacity at
	// provisioned_model_units.
	BurstScalingEnabled bool `json:"burst_scaling_enabled,omitempty"`
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string `json:"entity_name"`
	// The version of the model in Unity Catalog to be served.
	EntityVersion string `json:"entity_version,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// The number of model units to be provisioned. Dedicated inference requires
	// positive fixed capacity and does not support scale-to-zero.
	ProvisionedModelUnits int64 `json:"provisioned_model_units"`
	// Commitment term to purchase for this served entity; its presence marks
	// the entity as reserved (vs on-demand) PT. Valid on create and on update
	// (purchases a commitment for a scale-up).
	PtTerm CommitmentTerm `json:"pt_term,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PtServedModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PtServedModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PutAiGatewayRequest struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name string `json:"-" url:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func (s *PutAiGatewayRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PutAiGatewayResponse struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func (s *PutAiGatewayResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name string `json:"-" url:"-"`
	// The list of endpoint rate limits.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

func (s *PutRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

func (s *PutResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type QueryEndpointInput struct {
	// Optional user-provided request identifier that will be recorded in the
	// inference table and the usage tracking table.
	ClientRequestId string `json:"client_request_id,omitempty"`
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `json:"dataframe_records,omitempty"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *DataframeSplitInput `json:"dataframe_split,omitempty"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input any `json:"input,omitempty"`
	// Tensor-based input in columnar format.
	Inputs any `json:"inputs,omitempty"`
	// Tensor-based input in row format.
	Instances []any `json:"instances,omitempty"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens int `json:"max_tokens,omitempty"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is an array of ChatMessage objects and should
	// only be used with other chat query fields.
	Messages []ChatMessage `json:"messages,omitempty"`
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	N int `json:"n,omitempty"`
	// The name of the serving endpoint. This field is required and is provided
	// via the path parameter.
	Name string `json:"-" url:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	Prompt any `json:"prompt,omitempty"`
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	Stop []string `json:"stop,omitempty"`
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	Stream bool `json:"stream,omitempty"`
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	Temperature float64 `json:"temperature,omitempty"`
	// Optional user-provided context that will be recorded in the usage
	// tracking table.
	UsageContext map[string]string `json:"usage_context,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryEndpointInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEndpointInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryEndpointResponse struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices []V1ResponseChoiceElement `json:"choices,omitempty"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created int64 `json:"created,omitempty"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data []EmbeddingsV1ResponseEmbeddingElement `json:"data,omitempty"`
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	Id string `json:"id,omitempty"`
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	Model string `json:"model,omitempty"`
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	Object QueryEndpointResponseObject `json:"object,omitempty"`
	// The outputs of the feature serving endpoint.
	Outputs []json.RawMessage `json:"outputs,omitempty"`
	// The predictions returned by the serving endpoint.
	Predictions []any `json:"predictions,omitempty"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName string `json:"-" url:"-" header:"served-model-name,omitempty"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage *ExternalModelUsageElement `json:"usage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of object returned by the __external/foundation model__ serving
// endpoint, one of [text_completion, chat.completion, list (of embeddings)].
type QueryEndpointResponseObject string

const QueryEndpointResponseObjectChatCompletion QueryEndpointResponseObject = `chat.completion`

const QueryEndpointResponseObjectList QueryEndpointResponseObject = `list`

const QueryEndpointResponseObjectTextCompletion QueryEndpointResponseObject = `text_completion`

// String representation for [fmt.Print]
func (f *QueryEndpointResponseObject) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryEndpointResponseObject) Set(v string) error {
	switch v {
	case `chat.completion`, `list`, `text_completion`:
		*f = QueryEndpointResponseObject(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "chat.completion", "list", "text_completion"`, v)
	}
}

// Values returns all possible values for QueryEndpointResponseObject.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueryEndpointResponseObject) Values() []QueryEndpointResponseObject {
	return []QueryEndpointResponseObject{
		QueryEndpointResponseObjectChatCompletion,
		QueryEndpointResponseObjectList,
		QueryEndpointResponseObjectTextCompletion,
	}
}

// Type always returns QueryEndpointResponseObject to satisfy [pflag.Value] interface
func (f *QueryEndpointResponseObject) Type() string {
	return "QueryEndpointResponseObject"
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int64 `json:"calls"`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key RateLimitKey `json:"key,omitempty"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod RateLimitRenewalPeriod `json:"renewal_period"`
}

func (s *RateLimit) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RateLimitKey string

const RateLimitKeyEndpoint RateLimitKey = `endpoint`

const RateLimitKeyUser RateLimitKey = `user`

// String representation for [fmt.Print]
func (f *RateLimitKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RateLimitKey) Set(v string) error {
	switch v {
	case `endpoint`, `user`:
		*f = RateLimitKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "endpoint", "user"`, v)
	}
}

// Values returns all possible values for RateLimitKey.
//
// There is no guarantee on the order of the values in the slice.
func (f *RateLimitKey) Values() []RateLimitKey {
	return []RateLimitKey{
		RateLimitKeyEndpoint,
		RateLimitKeyUser,
	}
}

// Type always returns RateLimitKey to satisfy [pflag.Value] interface
func (f *RateLimitKey) Type() string {
	return "RateLimitKey"
}

type RateLimitRenewalPeriod string

const RateLimitRenewalPeriodMinute RateLimitRenewalPeriod = `minute`

// String representation for [fmt.Print]
func (f *RateLimitRenewalPeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RateLimitRenewalPeriod) Set(v string) error {
	switch v {
	case `minute`:
		*f = RateLimitRenewalPeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "minute"`, v)
	}
}

// Values returns all possible values for RateLimitRenewalPeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *RateLimitRenewalPeriod) Values() []RateLimitRenewalPeriod {
	return []RateLimitRenewalPeriod{
		RateLimitRenewalPeriodMinute,
	}
}

// Type always returns RateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *RateLimitRenewalPeriod) Type() string {
	return "RateLimitRenewalPeriod"
}

type Route struct {
	// The name of the served entity this route configures traffic for.
	ServedEntityName string `json:"served_entity_name,omitempty"`
	// The name of the served model this route configures traffic for.
	ServedModelName string `json:"served_model_name,omitempty"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int `json:"traffic_percentage"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Route) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Route) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntityInput struct {
	// Whether burst scaling is enabled. When enabled (default), the endpoint
	// can automatically scale up beyond provisioned capacity to handle traffic
	// spikes. When disabled, the endpoint maintains fixed capacity at
	// provisioned_model_units.
	BurstScalingEnabled bool `json:"burst_scaling_enabled,omitempty"`
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string `json:"entity_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served, or empty if the entity is a FEATURE_SPEC function.
	EntityVersion string `json:"entity_version,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
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
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedEntityInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntityOutput struct {
	// Whether burst scaling is enabled. When enabled (default), the endpoint
	// can automatically scale up beyond provisioned capacity to handle traffic
	// spikes. When disabled, the endpoint maintains fixed capacity at
	// provisioned_model_units.
	BurstScalingEnabled bool `json:"burst_scaling_enabled,omitempty"`
	// The timestamp when the served entity was created, in milliseconds since
	// the Unix epoch.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the served entity.
	Creator string `json:"creator,omitempty"`
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string `json:"entity_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served, or empty if the entity is a FEATURE_SPEC function.
	EntityVersion string `json:"entity_version,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// The foundation model that is served. Present for served entities that
	// serve a Databricks-hosted foundation model.
	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// Information corresponding to the state of the served entity.
	State *ServedModelState `json:"state,omitempty"`
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
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedEntityOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntitySpec struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string `json:"entity_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served, or empty if the entity is a FEATURE_SPEC function.
	EntityVersion string `json:"entity_version,omitempty"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`
	// The name of the served entity.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedEntitySpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntitySpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelInput struct {
	// Whether burst scaling is enabled. When enabled (default), the endpoint
	// can automatically scale up beyond provisioned capacity to handle traffic
	// spikes. When disabled, the endpoint maintains fixed capacity at
	// provisioned_model_units.
	BurstScalingEnabled bool `json:"burst_scaling_enabled,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of the model in Databricks Model Registry to be served or, if
	// the model resides in Unity Catalog, the full name of the model in the
	// form of **catalog_name.schema_name.model_name**.
	ModelName string `json:"model_name"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled"`
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
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServedModelInputWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedModelInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Please keep this in sync with workload types in
// InferenceEndpointEntities.scala.
type ServedModelInputWorkloadType string

const ServedModelInputWorkloadTypeCpu ServedModelInputWorkloadType = `CPU`

const ServedModelInputWorkloadTypeCpuLarge ServedModelInputWorkloadType = `CPU_LARGE`

const ServedModelInputWorkloadTypeCpuMedium ServedModelInputWorkloadType = `CPU_MEDIUM`

const ServedModelInputWorkloadTypeGpuLarge ServedModelInputWorkloadType = `GPU_LARGE`

const ServedModelInputWorkloadTypeGpuMedium ServedModelInputWorkloadType = `GPU_MEDIUM`

const ServedModelInputWorkloadTypeGpuSmall ServedModelInputWorkloadType = `GPU_SMALL`

const ServedModelInputWorkloadTypeGpuXlarge ServedModelInputWorkloadType = `GPU_XLARGE`

const ServedModelInputWorkloadTypeGpuXlarge8 ServedModelInputWorkloadType = `GPU_XLARGE_8`

const ServedModelInputWorkloadTypeMultigpuMedium ServedModelInputWorkloadType = `MULTIGPU_MEDIUM`

// String representation for [fmt.Print]
func (f *ServedModelInputWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelInputWorkloadType) Set(v string) error {
	switch v {
	case `CPU`, `CPU_LARGE`, `CPU_MEDIUM`, `GPU_LARGE`, `GPU_MEDIUM`, `GPU_SMALL`, `GPU_XLARGE`, `GPU_XLARGE_8`, `MULTIGPU_MEDIUM`:
		*f = ServedModelInputWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CPU", "CPU_LARGE", "CPU_MEDIUM", "GPU_LARGE", "GPU_MEDIUM", "GPU_SMALL", "GPU_XLARGE", "GPU_XLARGE_8", "MULTIGPU_MEDIUM"`, v)
	}
}

// Values returns all possible values for ServedModelInputWorkloadType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServedModelInputWorkloadType) Values() []ServedModelInputWorkloadType {
	return []ServedModelInputWorkloadType{
		ServedModelInputWorkloadTypeCpu,
		ServedModelInputWorkloadTypeCpuLarge,
		ServedModelInputWorkloadTypeCpuMedium,
		ServedModelInputWorkloadTypeGpuLarge,
		ServedModelInputWorkloadTypeGpuMedium,
		ServedModelInputWorkloadTypeGpuSmall,
		ServedModelInputWorkloadTypeGpuXlarge,
		ServedModelInputWorkloadTypeGpuXlarge8,
		ServedModelInputWorkloadTypeMultigpuMedium,
	}
}

// Type always returns ServedModelInputWorkloadType to satisfy [pflag.Value] interface
func (f *ServedModelInputWorkloadType) Type() string {
	return "ServedModelInputWorkloadType"
}

type ServedModelOutput struct {
	// Whether burst scaling is enabled. When enabled (default), the endpoint
	// can automatically scale up beyond provisioned capacity to handle traffic
	// spikes. When disabled, the endpoint maintains fixed capacity at
	// provisioned_model_units.
	BurstScalingEnabled bool `json:"burst_scaling_enabled,omitempty"`
	// The timestamp when the served entity was created, in milliseconds since
	// the Unix epoch.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the served entity.
	Creator string `json:"creator,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The name of the model in Databricks Model Registry to be served or, if
	// the model resides in Unity Catalog, the full name of the model in the
	// form of **catalog_name.schema_name.model_name**.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// Information corresponding to the state of the served entity.
	State *ServedModelState `json:"state,omitempty"`
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
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedModelOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelSpec struct {
	// The name of the model in Databricks Model Registry to be served or, if
	// the model resides in Unity Catalog, the full name of the model in the
	// form of **catalog_name.schema_name.model_name**.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of the served entity.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedModelSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelState struct {
	// The state of the deployment of the served entity. DEPLOYMENT_CREATING
	// indicates that the served entity is not ready yet because the deployment
	// is still being created (for example, the container image is building or
	// the model server is starting for the first time). DEPLOYMENT_RECOVERING
	// indicates that the served entity was previously ready but no longer is
	// and is attempting to recover. DEPLOYMENT_READY indicates that the served
	// entity is ready to receive traffic. DEPLOYMENT_FAILED indicates that
	// there was an error bringing up the served entity (for example, the
	// container image build failed or the model server failed to start due to a
	// model loading error). DEPLOYMENT_ABORTED indicates that the deployment
	// was terminated, likely due to a failure in bringing up another served
	// entity under the same endpoint and config version.
	Deployment ServedModelStateDeployment `json:"deployment,omitempty"`
	// More information about the state of the served entity, if available.
	DeploymentStateMessage string `json:"deployment_state_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServedModelState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelStateDeployment string

const ServedModelStateDeploymentDeploymentAborted ServedModelStateDeployment = `DEPLOYMENT_ABORTED`

const ServedModelStateDeploymentDeploymentCreating ServedModelStateDeployment = `DEPLOYMENT_CREATING`

const ServedModelStateDeploymentDeploymentFailed ServedModelStateDeployment = `DEPLOYMENT_FAILED`

const ServedModelStateDeploymentDeploymentReady ServedModelStateDeployment = `DEPLOYMENT_READY`

const ServedModelStateDeploymentDeploymentRecovering ServedModelStateDeployment = `DEPLOYMENT_RECOVERING`

// String representation for [fmt.Print]
func (f *ServedModelStateDeployment) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelStateDeployment) Set(v string) error {
	switch v {
	case `DEPLOYMENT_ABORTED`, `DEPLOYMENT_CREATING`, `DEPLOYMENT_FAILED`, `DEPLOYMENT_READY`, `DEPLOYMENT_RECOVERING`:
		*f = ServedModelStateDeployment(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_ABORTED", "DEPLOYMENT_CREATING", "DEPLOYMENT_FAILED", "DEPLOYMENT_READY", "DEPLOYMENT_RECOVERING"`, v)
	}
}

// Values returns all possible values for ServedModelStateDeployment.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServedModelStateDeployment) Values() []ServedModelStateDeployment {
	return []ServedModelStateDeployment{
		ServedModelStateDeploymentDeploymentAborted,
		ServedModelStateDeploymentDeploymentCreating,
		ServedModelStateDeploymentDeploymentFailed,
		ServedModelStateDeploymentDeploymentReady,
		ServedModelStateDeploymentDeploymentRecovering,
	}
}

// Type always returns ServedModelStateDeployment to satisfy [pflag.Value] interface
func (f *ServedModelStateDeployment) Type() string {
	return "ServedModelStateDeployment"
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs string `json:"logs"`
}

func (s *ServerLogsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigSummary `json:"config,omitempty"`
	// The timestamp when the endpoint was created, in milliseconds since the
	// Unix epoch.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// Description of the endpoint
	Description string `json:"description,omitempty"`
	// System-generated ID of the endpoint, included to be used by the
	// Permissions API.
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user, in
	// milliseconds since the Unix epoch.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	Task string `json:"task,omitempty"`
	// Telemetry configuration for the endpoint, including inference-table
	// payload logging.
	TelemetryConfig *TelemetryConfig `json:"telemetry_config,omitempty"`
	// The usage policy associated with serving endpoint.
	UsagePolicyId string `json:"usage_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []ServingEndpointPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointDetailed struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigOutput `json:"config,omitempty"`
	// The timestamp when the endpoint was created, in milliseconds since the
	// Unix epoch.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo *ModelDataPlaneInfo `json:"data_plane_info,omitempty"`
	// Description of the serving endpoint.
	Description string `json:"description,omitempty"`
	// Email notification settings.
	EmailNotifications *EmailNotifications `json:"email_notifications,omitempty"`
	// Endpoint invocation url if route optimization is enabled for endpoint
	EndpointUrl string `json:"endpoint_url,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user, in
	// milliseconds since the Unix epoch.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// The config that the endpoint is attempting to update to.
	PendingConfig *EndpointPendingConfig `json:"pending_config,omitempty"`
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel `json:"permission_level,omitempty"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized bool `json:"route_optimized,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	Task string `json:"task,omitempty"`
	// Telemetry configuration for the endpoint, including inference-table
	// payload logging.
	TelemetryConfig *TelemetryConfig `json:"telemetry_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpointDetailed) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointDetailed) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointDetailedPermissionLevel string

const ServingEndpointDetailedPermissionLevelCanManage ServingEndpointDetailedPermissionLevel = `CAN_MANAGE`

const ServingEndpointDetailedPermissionLevelCanQuery ServingEndpointDetailedPermissionLevel = `CAN_QUERY`

const ServingEndpointDetailedPermissionLevelCanView ServingEndpointDetailedPermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *ServingEndpointDetailedPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingEndpointDetailedPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = ServingEndpointDetailedPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Values returns all possible values for ServingEndpointDetailedPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServingEndpointDetailedPermissionLevel) Values() []ServingEndpointDetailedPermissionLevel {
	return []ServingEndpointDetailedPermissionLevel{
		ServingEndpointDetailedPermissionLevelCanManage,
		ServingEndpointDetailedPermissionLevelCanQuery,
		ServingEndpointDetailedPermissionLevelCanView,
	}
}

// Type always returns ServingEndpointDetailedPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointDetailedPermissionLevel) Type() string {
	return "ServingEndpointDetailedPermissionLevel"
}

type ServingEndpointPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpointPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ServingEndpointPermissionLevel string

const ServingEndpointPermissionLevelCanManage ServingEndpointPermissionLevel = `CAN_MANAGE`

const ServingEndpointPermissionLevelCanQuery ServingEndpointPermissionLevel = `CAN_QUERY`

const ServingEndpointPermissionLevelCanView ServingEndpointPermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *ServingEndpointPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingEndpointPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = ServingEndpointPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Values returns all possible values for ServingEndpointPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServingEndpointPermissionLevel) Values() []ServingEndpointPermissionLevel {
	return []ServingEndpointPermissionLevel{
		ServingEndpointPermissionLevelCanManage,
		ServingEndpointPermissionLevelCanQuery,
		ServingEndpointPermissionLevelCanView,
	}
}

// Type always returns ServingEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointPermissionLevel) Type() string {
	return "ServingEndpointPermissionLevel"
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpointPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServingEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest `json:"access_control_list,omitempty"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

func (s *ServingEndpointPermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Please keep this in sync with workload types in
// InferenceEndpointEntities.scala.
type ServingModelWorkloadType string

const ServingModelWorkloadTypeCpu ServingModelWorkloadType = `CPU`

const ServingModelWorkloadTypeCpuLarge ServingModelWorkloadType = `CPU_LARGE`

const ServingModelWorkloadTypeCpuMedium ServingModelWorkloadType = `CPU_MEDIUM`

const ServingModelWorkloadTypeGpuLarge ServingModelWorkloadType = `GPU_LARGE`

const ServingModelWorkloadTypeGpuMedium ServingModelWorkloadType = `GPU_MEDIUM`

const ServingModelWorkloadTypeGpuSmall ServingModelWorkloadType = `GPU_SMALL`

const ServingModelWorkloadTypeGpuXlarge ServingModelWorkloadType = `GPU_XLARGE`

const ServingModelWorkloadTypeGpuXlarge8 ServingModelWorkloadType = `GPU_XLARGE_8`

const ServingModelWorkloadTypeMultigpuMedium ServingModelWorkloadType = `MULTIGPU_MEDIUM`

// String representation for [fmt.Print]
func (f *ServingModelWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingModelWorkloadType) Set(v string) error {
	switch v {
	case `CPU`, `CPU_LARGE`, `CPU_MEDIUM`, `GPU_LARGE`, `GPU_MEDIUM`, `GPU_SMALL`, `GPU_XLARGE`, `GPU_XLARGE_8`, `MULTIGPU_MEDIUM`:
		*f = ServingModelWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CPU", "CPU_LARGE", "CPU_MEDIUM", "GPU_LARGE", "GPU_MEDIUM", "GPU_SMALL", "GPU_XLARGE", "GPU_XLARGE_8", "MULTIGPU_MEDIUM"`, v)
	}
}

// Values returns all possible values for ServingModelWorkloadType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServingModelWorkloadType) Values() []ServingModelWorkloadType {
	return []ServingModelWorkloadType{
		ServingModelWorkloadTypeCpu,
		ServingModelWorkloadTypeCpuLarge,
		ServingModelWorkloadTypeCpuMedium,
		ServingModelWorkloadTypeGpuLarge,
		ServingModelWorkloadTypeGpuMedium,
		ServingModelWorkloadTypeGpuSmall,
		ServingModelWorkloadTypeGpuXlarge,
		ServingModelWorkloadTypeGpuXlarge8,
		ServingModelWorkloadTypeMultigpuMedium,
	}
}

// Type always returns ServingModelWorkloadType to satisfy [pflag.Value] interface
func (f *ServingModelWorkloadType) Type() string {
	return "ServingModelWorkloadType"
}

type TelemetryConfig struct {
	// The telemetry signals to enable for this endpoint. If empty or omitted,
	// all signals are enabled; otherwise only the listed signals are enabled.
	EnabledTelemetryFeatures []TelemetryFeature `json:"enabled_telemetry_features,omitempty"`
	// Configuration for inference table payload logging, including sampling.
	InferenceTableConfig *TelemetryInferenceTableConfig `json:"inference_table_config,omitempty"`
	// The Unity Catalog tables to which endpoint telemetry (logs, traces, and
	// metrics) is exported. Provide this to create a new telemetry profile for
	// the endpoint from the given tables.
	TableNames *UnityCatalogTableNames `json:"table_names,omitempty"`
	// The ID of an existing telemetry profile to apply to this endpoint.
	// Provide this to reuse a telemetry profile that has already been created,
	// instead of specifying table_names.
	TelemetryProfileId string `json:"telemetry_profile_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TelemetryConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TelemetryConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A telemetry signal that a serving endpoint can export to Unity Catalog. Use
// these values to select which signals the endpoint exports.
type TelemetryFeature string

const TelemetryFeatureTelemetryFeatureInferenceTable TelemetryFeature = `TELEMETRY_FEATURE_INFERENCE_TABLE`

const TelemetryFeatureTelemetryFeatureLogs TelemetryFeature = `TELEMETRY_FEATURE_LOGS`

const TelemetryFeatureTelemetryFeatureMetrics TelemetryFeature = `TELEMETRY_FEATURE_METRICS`

const TelemetryFeatureTelemetryFeatureTraces TelemetryFeature = `TELEMETRY_FEATURE_TRACES`

// String representation for [fmt.Print]
func (f *TelemetryFeature) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TelemetryFeature) Set(v string) error {
	switch v {
	case `TELEMETRY_FEATURE_INFERENCE_TABLE`, `TELEMETRY_FEATURE_LOGS`, `TELEMETRY_FEATURE_METRICS`, `TELEMETRY_FEATURE_TRACES`:
		*f = TelemetryFeature(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TELEMETRY_FEATURE_INFERENCE_TABLE", "TELEMETRY_FEATURE_LOGS", "TELEMETRY_FEATURE_METRICS", "TELEMETRY_FEATURE_TRACES"`, v)
	}
}

// Values returns all possible values for TelemetryFeature.
//
// There is no guarantee on the order of the values in the slice.
func (f *TelemetryFeature) Values() []TelemetryFeature {
	return []TelemetryFeature{
		TelemetryFeatureTelemetryFeatureInferenceTable,
		TelemetryFeatureTelemetryFeatureLogs,
		TelemetryFeatureTelemetryFeatureMetrics,
		TelemetryFeatureTelemetryFeatureTraces,
	}
}

// Type always returns TelemetryFeature to satisfy [pflag.Value] interface
func (f *TelemetryFeature) Type() string {
	return "TelemetryFeature"
}

// Inference table payload logging configuration
type TelemetryInferenceTableConfig struct {
	// The full name of the inference table created for this endpoint.
	Name string `json:"name,omitempty"`
	// Fraction of requests sampled for payload logging, in the range [0.0,
	// 1.0], where 1.0 logs all requests.
	SamplingFraction float64 `json:"sampling_fraction,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TelemetryInferenceTableConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TelemetryInferenceTableConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes []Route `json:"routes,omitempty"`
}

func (s *TrafficConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UnityCatalogTableNames struct {
	// The full three-level Unity Catalog name (catalog.schema.table) of the
	// table that receives exported annotations.
	AnnotationsTable string `json:"annotations_table,omitempty"`
	// The full three-level Unity Catalog name (catalog.schema.table) of the
	// table that receives exported logs.
	LogsTable string `json:"logs_table,omitempty"`
	// The full three-level Unity Catalog name (catalog.schema.table) of the
	// table that receives exported metrics.
	MetricsTable string `json:"metrics_table,omitempty"`
	// The full three-level Unity Catalog name (catalog.schema.table) of the
	// table that receives exported traces (spans).
	TracesTable string `json:"traces_table,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UnityCatalogTableNames) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UnityCatalogTableNames) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateInferenceEndpointNotifications struct {
	// The email notification settings to update. Specify email addresses to
	// notify when endpoint state changes occur.
	EmailNotifications *EmailNotifications `json:"email_notifications,omitempty"`
	// The name of the serving endpoint whose notifications are being updated.
	// This field is required.
	Name string `json:"-" url:"-"`
}

func (s *UpdateInferenceEndpointNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateInferenceEndpointNotificationsResponse struct {
	// The email notification settings now in effect for the serving endpoint.
	EmailNotifications *EmailNotifications `json:"email_notifications,omitempty"`
	// The name of the serving endpoint whose notification settings were
	// updated.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateInferenceEndpointNotificationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInferenceEndpointNotificationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateProvisionedThroughputEndpointConfigRequest struct {
	// The core config of the serving endpoint.
	Config PtEndpointCoreConfig `json:"config"`
	// The name of the pt endpoint to update. This field is required.
	Name string `json:"-" url:"-"`
}

func (s *UpdateProvisionedThroughputEndpointConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason string `json:"finishReason,omitempty"`
	// The index of the choice in the __chat or completions__ response.
	Index int `json:"index,omitempty"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs int `json:"logprobs,omitempty"`
	// The message response from the __chat__ endpoint.
	Message *ChatMessage `json:"message,omitempty"`
	// The text response from the __completions__ endpoint.
	Text string `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *V1ResponseChoiceElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s V1ResponseChoiceElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
