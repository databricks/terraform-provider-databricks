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
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21Labs API key.
	Ai21labsApiKey types.String `tfsdk:"ai21labs_api_key" tf:""`
}

type AmazonBedrockConfig struct {
	// The Databricks secret key reference for an AWS Access Key ID with
	// permissions to interact with Bedrock services.
	AwsAccessKeyId types.String `tfsdk:"aws_access_key_id" tf:""`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion types.String `tfsdk:"aws_region" tf:""`
	// The Databricks secret key reference for an AWS Secret Access Key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services.
	AwsSecretAccessKey types.String `tfsdk:"aws_secret_access_key" tf:""`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider AmazonBedrockConfigBedrockProvider `tfsdk:"bedrock_provider" tf:""`
}

// The underlying provider in Amazon Bedrock. Supported values (case
// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
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

// Type always returns AmazonBedrockConfigBedrockProvider to satisfy [pflag.Value] interface
func (f *AmazonBedrockConfigBedrockProvider) Type() string {
	return "AmazonBedrockConfigBedrockProvider"
}

type AnthropicConfig struct {
	// The Databricks secret key reference for an Anthropic API key.
	AnthropicApiKey types.String `tfsdk:"anthropic_api_key" tf:""`
}

type App struct {
	// The active deployment of the app.
	ActiveDeployment *AppDeployment `tfsdk:"active_deployment" tf:"optional"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The email of the user that created the app.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
	// The pending deployment of the app.
	PendingDeployment *AppDeployment `tfsdk:"pending_deployment" tf:"optional"`

	ServicePrincipalId types.Int64 `tfsdk:"service_principal_id" tf:"optional"`

	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`

	Status *AppStatus `tfsdk:"status" tf:"optional"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// The email of the user that last updated the app.
	Updater types.String `tfsdk:"updater" tf:"optional"`
	// The URL of the app once it is deployed.
	Url types.String `tfsdk:"url" tf:"optional"`
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The email of the user creates the deployment.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// The deployment artifacts for an app.
	DeploymentArtifacts *AppDeploymentArtifacts `tfsdk:"deployment_artifacts" tf:"optional"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"deployment_id" tf:"optional"`
	// The mode of which the deployment will manage the source code.
	Mode AppDeploymentMode `tfsdk:"mode" tf:""`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath types.String `tfsdk:"source_code_path" tf:""`
	// Status and status message of the deployment
	Status *AppDeploymentStatus `tfsdk:"status" tf:"optional"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

type AppDeploymentArtifacts struct {
	// The snapshotted workspace file system path of the source code loaded by
	// the deployed app.
	SourceCodePath types.String `tfsdk:"source_code_path" tf:"optional"`
}

type AppDeploymentMode string

const AppDeploymentModeAutoSync AppDeploymentMode = `AUTO_SYNC`

const AppDeploymentModeModeUnspecified AppDeploymentMode = `MODE_UNSPECIFIED`

const AppDeploymentModeSnapshot AppDeploymentMode = `SNAPSHOT`

// String representation for [fmt.Print]
func (f *AppDeploymentMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppDeploymentMode) Set(v string) error {
	switch v {
	case `AUTO_SYNC`, `MODE_UNSPECIFIED`, `SNAPSHOT`:
		*f = AppDeploymentMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO_SYNC", "MODE_UNSPECIFIED", "SNAPSHOT"`, v)
	}
}

// Type always returns AppDeploymentMode to satisfy [pflag.Value] interface
func (f *AppDeploymentMode) Type() string {
	return "AppDeploymentMode"
}

type AppDeploymentState string

const AppDeploymentStateFailed AppDeploymentState = `FAILED`

const AppDeploymentStateInProgress AppDeploymentState = `IN_PROGRESS`

const AppDeploymentStateStateUnspecified AppDeploymentState = `STATE_UNSPECIFIED`

const AppDeploymentStateStopped AppDeploymentState = `STOPPED`

const AppDeploymentStateSucceeded AppDeploymentState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *AppDeploymentState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppDeploymentState) Set(v string) error {
	switch v {
	case `FAILED`, `IN_PROGRESS`, `STATE_UNSPECIFIED`, `STOPPED`, `SUCCEEDED`:
		*f = AppDeploymentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "IN_PROGRESS", "STATE_UNSPECIFIED", "STOPPED", "SUCCEEDED"`, v)
	}
}

// Type always returns AppDeploymentState to satisfy [pflag.Value] interface
func (f *AppDeploymentState) Type() string {
	return "AppDeploymentState"
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message types.String `tfsdk:"message" tf:"optional"`
	// State of the deployment.
	State AppDeploymentState `tfsdk:"state" tf:"optional"`
}

type AppEnvironment struct {
	Env []EnvVariable `tfsdk:"env" tf:"optional"`
}

type AppState string

const AppStateCreating AppState = `CREATING`

const AppStateDeleted AppState = `DELETED`

const AppStateDeleting AppState = `DELETING`

const AppStateError AppState = `ERROR`

const AppStateIdle AppState = `IDLE`

const AppStateRunning AppState = `RUNNING`

const AppStateStarting AppState = `STARTING`

const AppStateStateUnspecified AppState = `STATE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *AppState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppState) Set(v string) error {
	switch v {
	case `CREATING`, `DELETED`, `DELETING`, `ERROR`, `IDLE`, `RUNNING`, `STARTING`, `STATE_UNSPECIFIED`:
		*f = AppState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATING", "DELETED", "DELETING", "ERROR", "IDLE", "RUNNING", "STARTING", "STATE_UNSPECIFIED"`, v)
	}
}

// Type always returns AppState to satisfy [pflag.Value] interface
func (f *AppState) Type() string {
	return "AppState"
}

type AppStatus struct {
	// Message corresponding with the app state.
	Message types.String `tfsdk:"message" tf:"optional"`
	// State of the app.
	State AppState `tfsdk:"state" tf:"optional"`
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

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog.
	CatalogName types.String `tfsdk:"catalog_name" tf:"optional"`
	// Indicates whether the inference table is enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the schema in Unity Catalog.
	SchemaName types.String `tfsdk:"schema_name" tf:"optional"`

	State *AutoCaptureState `tfsdk:"state" tf:"optional"`
	// The prefix of the table in Unity Catalog.
	TableNamePrefix types.String `tfsdk:"table_name_prefix" tf:"optional"`
}

type AutoCaptureState struct {
	PayloadTable *PayloadTable `tfsdk:"payload_table" tf:"optional"`
}

// Get build logs for a served model
type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-" url:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName types.String `tfsdk:"-" url:"-"`
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs types.String `tfsdk:"logs" tf:""`
}

type ChatMessage struct {
	// The content of the message.
	Content types.String `tfsdk:"content" tf:"optional"`
	// The role of the message. One of [system, user, assistant].
	Role ChatMessageRole `tfsdk:"role" tf:"optional"`
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

// Type always returns ChatMessageRole to satisfy [pflag.Value] interface
func (f *ChatMessageRole) Type() string {
	return "ChatMessageRole"
}

type CohereConfig struct {
	// The Databricks secret key reference for a Cohere API key.
	CohereApiKey types.String `tfsdk:"cohere_api_key" tf:""`
}

type CreateAppDeploymentRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-" url:"-"`
	// The mode of which the deployment will manage the source code.
	Mode AppDeploymentMode `tfsdk:"mode" tf:""`
	// The workspace file system path of the source code used to create the app
	// deployment. This is different from
	// `deployment_artifacts.source_code_path`, which is the path used by the
	// deployed app. The former refers to the original source code location of
	// the app in the workspace during deployment creation, whereas the latter
	// provides a system generated stable snapshotted source code path used by
	// the deployment.
	SourceCodePath types.String `tfsdk:"source_code_path" tf:""`
}

type CreateAppRequest struct {
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:""`
}

type CreateServingEndpoint struct {
	// The core config of the serving endpoint.
	Config EndpointCoreConfigInput `tfsdk:"config" tf:""`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name types.String `tfsdk:"name" tf:""`
	// Rate limits to be applied to the serving endpoint. NOTE: only external
	// and foundation model endpoints are supported as of now.
	RateLimits []RateLimit `tfsdk:"rate_limits" tf:"optional"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized types.Bool `tfsdk:"route_optimized" tf:"optional"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `tfsdk:"tags" tf:"optional"`
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model.
	DatabricksApiToken types.String `tfsdk:"databricks_api_token" tf:""`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl types.String `tfsdk:"databricks_workspace_url" tf:""`
}

type DataframeSplitInput struct {
	Columns []any `tfsdk:"columns" tf:"optional"`

	Data []any `tfsdk:"data" tf:"optional"`

	Index []types.Int64 `tfsdk:"index" tf:"optional"`
}

// Delete an app
type DeleteAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-" url:"-"`
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	Embedding []types.Float64 `tfsdk:"embedding" tf:"optional"`
	// The index of the embedding in the response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// This will always be 'embedding'.
	Object EmbeddingsV1ResponseEmbeddingElementObject `tfsdk:"object" tf:"optional"`
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

// Type always returns EmbeddingsV1ResponseEmbeddingElementObject to satisfy [pflag.Value] interface
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Type() string {
	return "EmbeddingsV1ResponseEmbeddingElementObject"
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigInput `tfsdk:"auto_capture_config" tf:"optional"`
	// The name of the serving endpoint to update. This field is required.
	Name types.String `tfsdk:"-" url:"-"`
	// A list of served entities for the endpoint to serve. A serving endpoint
	// can have up to 15 served entities.
	ServedEntities []ServedEntityInput `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) A list of served models for the
	// endpoint to serve. A serving endpoint can have up to 15 served models.
	ServedModels []ServedModelInput `tfsdk:"served_models" tf:"optional"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `tfsdk:"traffic_config" tf:"optional"`
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigOutput `tfsdk:"auto_capture_config" tf:"optional"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion types.Int64 `tfsdk:"config_version" tf:"optional"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityOutput `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelOutput `tfsdk:"served_models" tf:"optional"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `tfsdk:"traffic_config" tf:"optional"`
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntitySpec `tfsdk:"served_entities" tf:"optional"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelSpec `tfsdk:"served_models" tf:"optional"`
}

type EndpointPendingConfig struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigOutput `tfsdk:"auto_capture_config" tf:"optional"`
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
	TrafficConfig *TrafficConfig `tfsdk:"traffic_config" tf:"optional"`
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails."
	ConfigUpdate EndpointStateConfigUpdate `tfsdk:"config_update" tf:"optional"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady `tfsdk:"ready" tf:"optional"`
}

// The state of an endpoint's config update. This informs the user if the
// pending_config is in progress, if the update failed, or if there is no update
// in progress. Note that if the endpoint's config_update state value is
// IN_PROGRESS, another update can not be made until the update completes or
// fails."
type EndpointStateConfigUpdate string

const EndpointStateConfigUpdateInProgress EndpointStateConfigUpdate = `IN_PROGRESS`

const EndpointStateConfigUpdateNotUpdating EndpointStateConfigUpdate = `NOT_UPDATING`

const EndpointStateConfigUpdateUpdateFailed EndpointStateConfigUpdate = `UPDATE_FAILED`

// String representation for [fmt.Print]
func (f *EndpointStateConfigUpdate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateConfigUpdate) Set(v string) error {
	switch v {
	case `IN_PROGRESS`, `NOT_UPDATING`, `UPDATE_FAILED`:
		*f = EndpointStateConfigUpdate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN_PROGRESS", "NOT_UPDATING", "UPDATE_FAILED"`, v)
	}
}

// Type always returns EndpointStateConfigUpdate to satisfy [pflag.Value] interface
func (f *EndpointStateConfigUpdate) Type() string {
	return "EndpointStateConfigUpdate"
}

// The state of an endpoint, indicating whether or not the endpoint is
// queryable. An endpoint is READY if all of the served entities in its active
// configuration are ready. If any of the actively served entities are in a
// non-ready state, the endpoint state will be NOT_READY.
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

// Type always returns EndpointStateReady to satisfy [pflag.Value] interface
func (f *EndpointStateReady) Type() string {
	return "EndpointStateReady"
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key types.String `tfsdk:"key" tf:""`
	// Optional value field for a serving endpoint tag.
	Value types.String `tfsdk:"value" tf:"optional"`
}

type EnvVariable struct {
	Name types.String `tfsdk:"name" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`

	ValueFrom types.String `tfsdk:"value_from" tf:"optional"`
}

// Get metrics of a serving endpoint
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name types.String `tfsdk:"-" url:"-"`
}

type ExportMetricsResponse struct {
	Contents io.ReadCloser `tfsdk:"-"`
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig *Ai21LabsConfig `tfsdk:"ai21labs_config" tf:"optional"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig *AmazonBedrockConfig `tfsdk:"amazon_bedrock_config" tf:"optional"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig *AnthropicConfig `tfsdk:"anthropic_config" tf:"optional"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig *CohereConfig `tfsdk:"cohere_config" tf:"optional"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig *DatabricksModelServingConfig `tfsdk:"databricks_model_serving_config" tf:"optional"`
	// The name of the external model.
	Name types.String `tfsdk:"name" tf:""`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig *OpenAiConfig `tfsdk:"openai_config" tf:"optional"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig *PaLmConfig `tfsdk:"palm_config" tf:"optional"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'openai', and 'palm'.",
	Provider ExternalModelProvider `tfsdk:"provider" tf:""`
	// The task type of the external model.
	Task types.String `tfsdk:"task" tf:""`
}

// The name of the provider for the external model. Currently, the supported
// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
// 'databricks-model-serving', 'openai', and 'palm'.",
type ExternalModelProvider string

const ExternalModelProviderAi21labs ExternalModelProvider = `ai21labs`

const ExternalModelProviderAmazonBedrock ExternalModelProvider = `amazon-bedrock`

const ExternalModelProviderAnthropic ExternalModelProvider = `anthropic`

const ExternalModelProviderCohere ExternalModelProvider = `cohere`

const ExternalModelProviderDatabricksModelServing ExternalModelProvider = `databricks-model-serving`

const ExternalModelProviderOpenai ExternalModelProvider = `openai`

const ExternalModelProviderPalm ExternalModelProvider = `palm`

// String representation for [fmt.Print]
func (f *ExternalModelProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExternalModelProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `amazon-bedrock`, `anthropic`, `cohere`, `databricks-model-serving`, `openai`, `palm`:
		*f = ExternalModelProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "amazon-bedrock", "anthropic", "cohere", "databricks-model-serving", "openai", "palm"`, v)
	}
}

// Type always returns ExternalModelProvider to satisfy [pflag.Value] interface
func (f *ExternalModelProvider) Type() string {
	return "ExternalModelProvider"
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens types.Int64 `tfsdk:"completion_tokens" tf:"optional"`
	// The number of tokens in the prompt.
	PromptTokens types.Int64 `tfsdk:"prompt_tokens" tf:"optional"`
	// The total number of tokens in the prompt and response.
	TotalTokens types.Int64 `tfsdk:"total_tokens" tf:"optional"`
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

// Get an app deployment
type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-" url:"-"`
	// The unique id of the deployment.
	DeploymentId types.String `tfsdk:"-" url:"-"`
}

// Get app environment
type GetAppEnvironmentRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-" url:"-"`
}

// Get an app
type GetAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-" url:"-"`
}

// Get the schema for a serving endpoint
type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-" url:"-"`
}

// The response is an OpenAPI spec in JSON format that typically includes fields
// like openapi, info, servers and paths, etc.
type GetOpenApiResponse struct {
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-" url:"-"`
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-" url:"-"`
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name types.String `tfsdk:"-" url:"-"`
}

// List app deployments
type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName types.String `tfsdk:"-" url:"-"`
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments []AppDeployment `tfsdk:"app_deployments" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListAppsResponse struct {
	Apps []App `tfsdk:"apps" tf:"optional"`
	// Pagination token to request the next page of apps.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint `tfsdk:"endpoints" tf:"optional"`
}

// Get the latest logs for a served model
type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name types.String `tfsdk:"-" url:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName types.String `tfsdk:"-" url:"-"`
}

type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo *oauth2.DataPlaneInfo `tfsdk:"query_info" tf:"optional"`
}

type OpenAiConfig struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	MicrosoftEntraClientId types.String `tfsdk:"microsoft_entra_client_id" tf:"optional"`
	// The Databricks secret key reference for the Microsoft Entra Client Secret
	// that is only required for Azure AD OpenAI.
	MicrosoftEntraClientSecret types.String `tfsdk:"microsoft_entra_client_secret" tf:"optional"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId types.String `tfsdk:"microsoft_entra_tenant_id" tf:"optional"`
	// This is the base URL for the OpenAI API (default:
	// "https://api.openai.com/v1"). For Azure OpenAI, this field is required,
	// and is the base URL for the Azure OpenAI API service provided by Azure.
	OpenaiApiBase types.String `tfsdk:"openai_api_base" tf:"optional"`
	// The Databricks secret key reference for an OpenAI or Azure OpenAI API
	// key.
	OpenaiApiKey types.String `tfsdk:"openai_api_key" tf:"optional"`
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

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key.
	PalmApiKey types.String `tfsdk:"palm_api_key" tf:""`
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag `tfsdk:"add_tags" tf:"optional"`
	// List of tag keys to delete
	DeleteTags []types.String `tfsdk:"delete_tags" tf:"optional"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name types.String `tfsdk:"-" url:"-"`
}

type PayloadTable struct {
	// The name of the payload table.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The status of the payload table.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The status message of the payload table.
	StatusMessage types.String `tfsdk:"status_message" tf:"optional"`
}

// Update rate limits of a serving endpoint
type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name types.String `tfsdk:"-" url:"-"`
	// The list of endpoint rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits" tf:"optional"`
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits" tf:"optional"`
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `tfsdk:"dataframe_records" tf:"optional"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *DataframeSplitInput `tfsdk:"dataframe_split" tf:"optional"`
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
	Name types.String `tfsdk:"-" url:"-"`
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
	Object QueryEndpointResponseObject `tfsdk:"object" tf:"optional"`
	// The predictions returned by the serving endpoint.
	Predictions []any `tfsdk:"predictions" tf:"optional"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName types.String `tfsdk:"-" url:"-" header:"served-model-name,omitempty"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage *ExternalModelUsageElement `tfsdk:"usage" tf:"optional"`
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

// Type always returns QueryEndpointResponseObject to satisfy [pflag.Value] interface
func (f *QueryEndpointResponseObject) Type() string {
	return "QueryEndpointResponseObject"
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls types.Int64 `tfsdk:"calls" tf:""`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key RateLimitKey `tfsdk:"key" tf:"optional"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod RateLimitRenewalPeriod `tfsdk:"renewal_period" tf:""`
}

// Key field for a serving endpoint rate limit. Currently, only 'user' and
// 'endpoint' are supported, with 'endpoint' being the default if not specified.
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

// Type always returns RateLimitKey to satisfy [pflag.Value] interface
func (f *RateLimitKey) Type() string {
	return "RateLimitKey"
}

// Renewal period field for a serving endpoint rate limit. Currently, only
// 'minute' is supported.
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

// Type always returns RateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *RateLimitRenewalPeriod) Type() string {
	return "RateLimitRenewalPeriod"
}

type Route struct {
	// The name of the served model this route configures traffic for.
	ServedModelName types.String `tfsdk:"served_model_name" tf:""`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage types.Int64 `tfsdk:"traffic_percentage" tf:""`
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
	// for custom model serving for a Databricks registered model. When an
	// external_model is present, the served entities list can only have one
	// served_entity object. For an existing endpoint with external_model, it
	// can not be updated to an endpoint without external_model. If the endpoint
	// is created without external_model, users cannot update it to add
	// external_model later.
	ExternalModel *ExternalModel `tfsdk:"external_model" tf:"optional"`
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
	ExternalModel *ExternalModel `tfsdk:"external_model" tf:"optional"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	FoundationModel *FoundationModel `tfsdk:"foundation_model" tf:"optional"`
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
	State *ServedModelState `tfsdk:"state" tf:"optional"`
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
	ExternalModel *ExternalModel `tfsdk:"external_model" tf:"optional"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	FoundationModel *FoundationModel `tfsdk:"foundation_model" tf:"optional"`
	// The name of the served entity.
	Name types.String `tfsdk:"name" tf:"optional"`
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
	WorkloadSize ServedModelInputWorkloadSize `tfsdk:"workload_size" tf:""`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServedModelInputWorkloadType `tfsdk:"workload_type" tf:"optional"`
}

// The workload size of the served model. The workload size corresponds to a
// range of provisioned concurrency that the compute will autoscale between. A
// single unit of provisioned concurrency can process one request at a time.
// Valid workload sizes are "Small" (4 - 4 provisioned concurrency), "Medium" (8
// - 16 provisioned concurrency), and "Large" (16 - 64 provisioned concurrency).
// If scale-to-zero is enabled, the lower bound of the provisioned concurrency
// for each workload size will be 0.
type ServedModelInputWorkloadSize string

const ServedModelInputWorkloadSizeLarge ServedModelInputWorkloadSize = `Large`

const ServedModelInputWorkloadSizeMedium ServedModelInputWorkloadSize = `Medium`

const ServedModelInputWorkloadSizeSmall ServedModelInputWorkloadSize = `Small`

// String representation for [fmt.Print]
func (f *ServedModelInputWorkloadSize) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelInputWorkloadSize) Set(v string) error {
	switch v {
	case `Large`, `Medium`, `Small`:
		*f = ServedModelInputWorkloadSize(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Large", "Medium", "Small"`, v)
	}
}

// Type always returns ServedModelInputWorkloadSize to satisfy [pflag.Value] interface
func (f *ServedModelInputWorkloadSize) Type() string {
	return "ServedModelInputWorkloadSize"
}

// The workload type of the served model. The workload type selects which type
// of compute to use in the endpoint. The default value for this parameter is
// "CPU". For deep learning workloads, GPU acceleration is available by
// selecting workload types like GPU_SMALL and others. See the available [GPU
// types].
//
// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
type ServedModelInputWorkloadType string

const ServedModelInputWorkloadTypeCpu ServedModelInputWorkloadType = `CPU`

const ServedModelInputWorkloadTypeGpuLarge ServedModelInputWorkloadType = `GPU_LARGE`

const ServedModelInputWorkloadTypeGpuMedium ServedModelInputWorkloadType = `GPU_MEDIUM`

const ServedModelInputWorkloadTypeGpuSmall ServedModelInputWorkloadType = `GPU_SMALL`

const ServedModelInputWorkloadTypeMultigpuMedium ServedModelInputWorkloadType = `MULTIGPU_MEDIUM`

// String representation for [fmt.Print]
func (f *ServedModelInputWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelInputWorkloadType) Set(v string) error {
	switch v {
	case `CPU`, `GPU_LARGE`, `GPU_MEDIUM`, `GPU_SMALL`, `MULTIGPU_MEDIUM`:
		*f = ServedModelInputWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CPU", "GPU_LARGE", "GPU_MEDIUM", "GPU_SMALL", "MULTIGPU_MEDIUM"`, v)
	}
}

// Type always returns ServedModelInputWorkloadType to satisfy [pflag.Value] interface
func (f *ServedModelInputWorkloadType) Type() string {
	return "ServedModelInputWorkloadType"
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
	State *ServedModelState `tfsdk:"state" tf:"optional"`
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
	Deployment ServedModelStateDeployment `tfsdk:"deployment" tf:"optional"`
	// More information about the state of the served entity, if available.
	DeploymentStateMessage types.String `tfsdk:"deployment_state_message" tf:"optional"`
}

// The state of the served entity deployment. DEPLOYMENT_CREATING indicates that
// the served entity is not ready yet because the deployment is still being
// created (i.e container image is building, model server is deploying for the
// first time, etc.). DEPLOYMENT_RECOVERING indicates that the served entity was
// previously in a ready state but no longer is and is attempting to recover.
// DEPLOYMENT_READY indicates that the served entity is ready to receive
// traffic. DEPLOYMENT_FAILED indicates that there was an error trying to bring
// up the served entity (e.g container image build failed, the model server
// failed to start due to a model loading error, etc.) DEPLOYMENT_ABORTED
// indicates that the deployment was terminated likely due to a failure in
// bringing up another served entity under the same endpoint and config version.
type ServedModelStateDeployment string

const ServedModelStateDeploymentAborted ServedModelStateDeployment = `DEPLOYMENT_ABORTED`

const ServedModelStateDeploymentCreating ServedModelStateDeployment = `DEPLOYMENT_CREATING`

const ServedModelStateDeploymentFailed ServedModelStateDeployment = `DEPLOYMENT_FAILED`

const ServedModelStateDeploymentReady ServedModelStateDeployment = `DEPLOYMENT_READY`

const ServedModelStateDeploymentRecovering ServedModelStateDeployment = `DEPLOYMENT_RECOVERING`

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

// Type always returns ServedModelStateDeployment to satisfy [pflag.Value] interface
func (f *ServedModelStateDeployment) Type() string {
	return "ServedModelStateDeployment"
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs types.String `tfsdk:"logs" tf:""`
}

type ServingEndpoint struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigSummary `tfsdk:"config" tf:"optional"`
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
	State *EndpointState `tfsdk:"state" tf:"optional"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `tfsdk:"tags" tf:"optional"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task" tf:"optional"`
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
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

type ServingEndpointDetailed struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigOutput `tfsdk:"config" tf:"optional"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp types.Int64 `tfsdk:"creation_timestamp" tf:"optional"`
	// The email of the user who created the serving endpoint.
	Creator types.String `tfsdk:"creator" tf:"optional"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo *ModelDataPlaneInfo `tfsdk:"data_plane_info" tf:"optional"`
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
	PendingConfig *EndpointPendingConfig `tfsdk:"pending_config" tf:"optional"`
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized types.Bool `tfsdk:"route_optimized" tf:"optional"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `tfsdk:"state" tf:"optional"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `tfsdk:"tags" tf:"optional"`
	// The task type of the serving endpoint.
	Task types.String `tfsdk:"task" tf:"optional"`
}

// The permission level of the principal making the request.
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

// Type always returns ServingEndpointDetailedPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointDetailedPermissionLevel) Type() string {
	return "ServingEndpointDetailedPermissionLevel"
}

type ServingEndpointPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `tfsdk:"permission_level" tf:"optional"`
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

// Type always returns ServingEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointPermissionLevel) Type() string {
	return "ServingEndpointPermissionLevel"
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type ServingEndpointPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId types.String `tfsdk:"-" url:"-"`
}

type StartAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-" url:"-"`
}

type StopAppRequest struct {
	// The name of the app.
	Name types.String `tfsdk:"-" url:"-"`
}

type StopAppResponse struct {
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes []Route `tfsdk:"routes" tf:"optional"`
}

type UpdateAppRequest struct {
	// The description of the app.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens. It must be unique within the workspace.
	Name types.String `tfsdk:"name" tf:"" url:"-"`
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason types.String `tfsdk:"finishReason" tf:"optional"`
	// The index of the choice in the __chat or completions__ response.
	Index types.Int64 `tfsdk:"index" tf:"optional"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs types.Int64 `tfsdk:"logprobs" tf:"optional"`
	// The message response from the __chat__ endpoint.
	Message *ChatMessage `tfsdk:"message" tf:"optional"`
	// The text response from the __completions__ endpoint.
	Text types.String `tfsdk:"text" tf:"optional"`
}
