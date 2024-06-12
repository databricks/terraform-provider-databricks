// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21Labs API key.
	Ai21labsApiKey string `tfsdk:"ai21labs_api_key"`
}

type AmazonBedrockConfig struct {
	// The Databricks secret key reference for an AWS Access Key ID with
	// permissions to interact with Bedrock services.
	AwsAccessKeyId string `tfsdk:"aws_access_key_id"`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion string `tfsdk:"aws_region"`
	// The Databricks secret key reference for an AWS Secret Access Key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services.
	AwsSecretAccessKey string `tfsdk:"aws_secret_access_key"`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider AmazonBedrockConfigBedrockProvider `tfsdk:"bedrock_provider"`
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
	AnthropicApiKey string `tfsdk:"anthropic_api_key"`
}

type App struct {
	// The active deployment of the app.
	ActiveDeployment *AppDeployment `tfsdk:"active_deployment"`
	// The creation time of the app. Formatted timestamp in ISO 6801.
	CreateTime string `tfsdk:"create_time"`
	// The email of the user that created the app.
	Creator string `tfsdk:"creator"`
	// The description of the app.
	Description string `tfsdk:"description"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens and be between 2 and 30 characters long. It must
	// be unique within the workspace.
	Name string `tfsdk:"name"`
	// The pending deployment of the app.
	PendingDeployment *AppDeployment `tfsdk:"pending_deployment"`

	Status *AppStatus `tfsdk:"status"`
	// The update time of the app. Formatted timestamp in ISO 6801.
	UpdateTime string `tfsdk:"update_time"`
	// The email of the user that last updated the app.
	Updater string `tfsdk:"updater"`
	// The URL of the app once it is deployed.
	Url string `tfsdk:"url"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *App) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s App) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppDeployment struct {
	// The creation time of the deployment. Formatted timestamp in ISO 6801.
	CreateTime string `tfsdk:"create_time"`
	// The email of the user creates the deployment.
	Creator string `tfsdk:"creator"`
	// The deployment artifacts for an app.
	DeploymentArtifacts *AppDeploymentArtifacts `tfsdk:"deployment_artifacts"`
	// The unique id of the deployment.
	DeploymentId string `tfsdk:"deployment_id"`
	// The source code path of the deployment.
	SourceCodePath string `tfsdk:"source_code_path"`
	// Status and status message of the deployment
	Status *AppDeploymentStatus `tfsdk:"status"`
	// The update time of the deployment. Formatted timestamp in ISO 6801.
	UpdateTime string `tfsdk:"update_time"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AppDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppDeploymentArtifacts struct {
	// The source code of the deployment.
	SourceCodePath string `tfsdk:"source_code_path"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AppDeploymentArtifacts) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeploymentArtifacts) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppDeploymentState string

const AppDeploymentStateCancelled AppDeploymentState = `CANCELLED`

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
	case `CANCELLED`, `FAILED`, `IN_PROGRESS`, `STATE_UNSPECIFIED`, `STOPPED`, `SUCCEEDED`:
		*f = AppDeploymentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "IN_PROGRESS", "STATE_UNSPECIFIED", "STOPPED", "SUCCEEDED"`, v)
	}
}

// Type always returns AppDeploymentState to satisfy [pflag.Value] interface
func (f *AppDeploymentState) Type() string {
	return "AppDeploymentState"
}

type AppDeploymentStatus struct {
	// Message corresponding with the deployment state.
	Message string `tfsdk:"message"`
	// State of the deployment.
	State AppDeploymentState `tfsdk:"state"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AppDeploymentStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppDeploymentStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppEnvironment struct {
	Env []EnvVariable `tfsdk:"env"`
}

type AppState string

const AppStateCreating AppState = `CREATING`

const AppStateDeleted AppState = `DELETED`

const AppStateDeleting AppState = `DELETING`

const AppStateDeployed AppState = `DEPLOYED`

const AppStateDeploying AppState = `DEPLOYING`

const AppStateError AppState = `ERROR`

const AppStateIdle AppState = `IDLE`

const AppStateReady AppState = `READY`

const AppStateRunning AppState = `RUNNING`

const AppStateStarting AppState = `STARTING`

const AppStateStateUnspecified AppState = `STATE_UNSPECIFIED`

const AppStateUpdating AppState = `UPDATING`

// String representation for [fmt.Print]
func (f *AppState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AppState) Set(v string) error {
	switch v {
	case `CREATING`, `DELETED`, `DELETING`, `DEPLOYED`, `DEPLOYING`, `ERROR`, `IDLE`, `READY`, `RUNNING`, `STARTING`, `STATE_UNSPECIFIED`, `UPDATING`:
		*f = AppState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATING", "DELETED", "DELETING", "DEPLOYED", "DEPLOYING", "ERROR", "IDLE", "READY", "RUNNING", "STARTING", "STATE_UNSPECIFIED", "UPDATING"`, v)
	}
}

// Type always returns AppState to satisfy [pflag.Value] interface
func (f *AppState) Type() string {
	return "AppState"
}

type AppStatus struct {
	// Message corresponding with the app state.
	Message string `tfsdk:"message"`
	// State of the app.
	State AppState `tfsdk:"state"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AppStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureConfigInput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if it was already set.
	CatalogName string `tfsdk:"catalog_name"`
	// If inference tables are enabled or not. NOTE: If you have already
	// disabled payload logging once, you cannot enable again.
	Enabled bool `tfsdk:"enabled"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if it was already set.
	SchemaName string `tfsdk:"schema_name"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if it was already set.
	TableNamePrefix string `tfsdk:"table_name_prefix"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AutoCaptureConfigInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog.
	CatalogName string `tfsdk:"catalog_name"`
	// If inference tables are enabled or not.
	Enabled bool `tfsdk:"enabled"`
	// The name of the schema in Unity Catalog.
	SchemaName string `tfsdk:"schema_name"`

	State *AutoCaptureState `tfsdk:"state"`
	// The prefix of the table in Unity Catalog.
	TableNamePrefix string `tfsdk:"table_name_prefix"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *AutoCaptureConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureState struct {
	PayloadTable *PayloadTable `tfsdk:"payload_table"`
}

// Get build logs for a served model
type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `tfsdk:"-" url:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string `tfsdk:"-" url:"-"`
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs string `tfsdk:"logs"`
}

type ChatMessage struct {
	// The content of the message.
	Content string `tfsdk:"content"`
	// The role of the message. One of [system, user, assistant].
	Role ChatMessageRole `tfsdk:"role"`

	ForceSendFields []string `tfsdk:"-"`
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

// Type always returns ChatMessageRole to satisfy [pflag.Value] interface
func (f *ChatMessageRole) Type() string {
	return "ChatMessageRole"
}

type CohereConfig struct {
	// The Databricks secret key reference for a Cohere API key.
	CohereApiKey string `tfsdk:"cohere_api_key"`
}

type CreateAppDeploymentRequest struct {
	// The name of the app.
	AppName string `tfsdk:"-" url:"-"`
	// The source code path of the deployment.
	SourceCodePath string `tfsdk:"source_code_path"`
}

type CreateAppRequest struct {
	// The description of the app.
	Description string `tfsdk:"description"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens and be between 2 and 30 characters long. It must
	// be unique within the workspace.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateAppRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAppRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateServingEndpoint struct {
	// The core config of the serving endpoint.
	Config EndpointCoreConfigInput `tfsdk:"config"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `tfsdk:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: only external
	// and foundation model endpoints are supported as of now.
	RateLimits []RateLimit `tfsdk:"rate_limits"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized bool `tfsdk:"route_optimized"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `tfsdk:"tags"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *CreateServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model.
	DatabricksApiToken string `tfsdk:"databricks_api_token"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl string `tfsdk:"databricks_workspace_url"`
}

type DataframeSplitInput struct {
	Columns []any `tfsdk:"columns"`

	Data []any `tfsdk:"data"`

	Index []int `tfsdk:"index"`
}

// Delete an App
type DeleteAppRequest struct {
	// The name of the app.
	Name string `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `tfsdk:"-" url:"-"`
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	Embedding []float64 `tfsdk:"embedding"`
	// The index of the embedding in the response.
	Index int `tfsdk:"index"`
	// This will always be 'embedding'.
	Object EmbeddingsV1ResponseEmbeddingElementObject `tfsdk:"object"`

	ForceSendFields []string `tfsdk:"-"`
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

// Type always returns EmbeddingsV1ResponseEmbeddingElementObject to satisfy [pflag.Value] interface
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Type() string {
	return "EmbeddingsV1ResponseEmbeddingElementObject"
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigInput `tfsdk:"auto_capture_config"`
	// The name of the serving endpoint to update. This field is required.
	Name string `tfsdk:"-" url:"-"`
	// A list of served entities for the endpoint to serve. A serving endpoint
	// can have up to 15 served entities.
	ServedEntities []ServedEntityInput `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) A list of served models for the
	// endpoint to serve. A serving endpoint can have up to 15 served models.
	ServedModels []ServedModelInput `tfsdk:"served_models"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `tfsdk:"traffic_config"`
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigOutput `tfsdk:"auto_capture_config"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `tfsdk:"config_version"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityOutput `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelOutput `tfsdk:"served_models"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `tfsdk:"traffic_config"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EndpointCoreConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointCoreConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntitySpec `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelSpec `tfsdk:"served_models"`
}

type EndpointPendingConfig struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigOutput `tfsdk:"auto_capture_config"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `tfsdk:"config_version"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities []ServedEntityOutput `tfsdk:"served_entities"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels []ServedModelOutput `tfsdk:"served_models"`
	// The timestamp when the update to the pending config started.
	StartTime int64 `tfsdk:"start_time"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `tfsdk:"traffic_config"`

	ForceSendFields []string `tfsdk:"-"`
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
	// or fails."
	ConfigUpdate EndpointStateConfigUpdate `tfsdk:"config_update"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady `tfsdk:"ready"`
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
	Key string `tfsdk:"key"`
	// Optional value field for a serving endpoint tag.
	Value string `tfsdk:"value"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EndpointTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnvVariable struct {
	Name string `tfsdk:"name"`

	Value string `tfsdk:"value"`

	ValueFrom string `tfsdk:"value_from"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *EnvVariable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnvVariable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get metrics of a serving endpoint
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `tfsdk:"-" url:"-"`
}

type ExportMetricsResponse struct {
	Contents io.ReadCloser `tfsdk:"-"`
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig *Ai21LabsConfig `tfsdk:"ai21labs_config"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig *AmazonBedrockConfig `tfsdk:"amazon_bedrock_config"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig *AnthropicConfig `tfsdk:"anthropic_config"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig *CohereConfig `tfsdk:"cohere_config"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig *DatabricksModelServingConfig `tfsdk:"databricks_model_serving_config"`
	// The name of the external model.
	Name string `tfsdk:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig *OpenAiConfig `tfsdk:"openai_config"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig *PaLmConfig `tfsdk:"palm_config"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'openai', and 'palm'.",
	Provider ExternalModelProvider `tfsdk:"provider"`
	// The task type of the external model.
	Task string `tfsdk:"task"`
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
	CompletionTokens int `tfsdk:"completion_tokens"`
	// The number of tokens in the prompt.
	PromptTokens int `tfsdk:"prompt_tokens"`
	// The total number of tokens in the prompt and response.
	TotalTokens int `tfsdk:"total_tokens"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ExternalModelUsageElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalModelUsageElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FoundationModel struct {
	// The description of the foundation model.
	Description string `tfsdk:"description"`
	// The display name of the foundation model.
	DisplayName string `tfsdk:"display_name"`
	// The URL to the documentation of the foundation model.
	Docs string `tfsdk:"docs"`
	// The name of the foundation model.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *FoundationModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FoundationModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get an App Deployment
type GetAppDeploymentRequest struct {
	// The name of the app.
	AppName string `tfsdk:"-" url:"-"`
	// The unique id of the deployment.
	DeploymentId string `tfsdk:"-" url:"-"`
}

// Get App Environment
type GetAppEnvironmentRequest struct {
	// The name of the app.
	Name string `tfsdk:"-" url:"-"`
}

// Get an App
type GetAppRequest struct {
	// The name of the app.
	Name string `tfsdk:"-" url:"-"`
}

// Get the schema for a serving endpoint
type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `tfsdk:"-" url:"-"`
}

// The response is an OpenAPI spec in JSON format that typically includes fields
// like openapi, info, servers and paths, etc.
type GetOpenApiResponse struct {
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `tfsdk:"-" url:"-"`
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription `tfsdk:"permission_levels"`
}

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `tfsdk:"-" url:"-"`
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `tfsdk:"-" url:"-"`
}

// List App Deployments
type ListAppDeploymentsRequest struct {
	// The name of the app.
	AppName string `tfsdk:"-" url:"-"`
	// Upper bound for items returned.
	PageSize int `tfsdk:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAppDeploymentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppDeploymentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAppDeploymentsResponse struct {
	// Deployment history of the app.
	AppDeployments []AppDeployment `tfsdk:"app_deployments"`
	// Pagination token to request the next page of apps.
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAppDeploymentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppDeploymentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List Apps
type ListAppsRequest struct {
	// Upper bound for items returned.
	PageSize int `tfsdk:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of apps. Requests first page if
	// absent.
	PageToken string `tfsdk:"-" url:"page_token,omitempty"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAppsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAppsResponse struct {
	Apps []App `tfsdk:"apps"`
	// Pagination token to request the next page of apps.
	NextPageToken string `tfsdk:"next_page_token"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ListAppsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint `tfsdk:"endpoints"`
}

// Get the latest logs for a served model
type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `tfsdk:"-" url:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string `tfsdk:"-" url:"-"`
}

type OpenAiConfig struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	MicrosoftEntraClientId string `tfsdk:"microsoft_entra_client_id"`
	// The Databricks secret key reference for the Microsoft Entra Client Secret
	// that is only required for Azure AD OpenAI.
	MicrosoftEntraClientSecret string `tfsdk:"microsoft_entra_client_secret"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId string `tfsdk:"microsoft_entra_tenant_id"`
	// This is the base URL for the OpenAI API (default:
	// "https://api.openai.com/v1"). For Azure OpenAI, this field is required,
	// and is the base URL for the Azure OpenAI API service provided by Azure.
	OpenaiApiBase string `tfsdk:"openai_api_base"`
	// The Databricks secret key reference for an OpenAI or Azure OpenAI API
	// key.
	OpenaiApiKey string `tfsdk:"openai_api_key"`
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	OpenaiApiType string `tfsdk:"openai_api_type"`
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	OpenaiApiVersion string `tfsdk:"openai_api_version"`
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	OpenaiDeploymentName string `tfsdk:"openai_deployment_name"`
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	OpenaiOrganization string `tfsdk:"openai_organization"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *OpenAiConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OpenAiConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key.
	PalmApiKey string `tfsdk:"palm_api_key"`
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag `tfsdk:"add_tags"`
	// List of tag keys to delete
	DeleteTags []string `tfsdk:"delete_tags"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string `tfsdk:"-" url:"-"`
}

type PayloadTable struct {
	// The name of the payload table.
	Name string `tfsdk:"name"`
	// The status of the payload table.
	Status string `tfsdk:"status"`
	// The status message of the payload table.
	StatusMessage string `tfsdk:"status_message"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *PayloadTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PayloadTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update rate limits of a serving endpoint
type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name string `tfsdk:"-" url:"-"`
	// The list of endpoint rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits"`
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits []RateLimit `tfsdk:"rate_limits"`
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `tfsdk:"dataframe_records"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *DataframeSplitInput `tfsdk:"dataframe_split"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams map[string]string `tfsdk:"extra_params"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input any `tfsdk:"input"`
	// Tensor-based input in columnar format.
	Inputs any `tfsdk:"inputs"`
	// Tensor-based input in row format.
	Instances []any `tfsdk:"instances"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens int `tfsdk:"max_tokens"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	Messages []ChatMessage `tfsdk:"messages"`
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	N int `tfsdk:"n"`
	// The name of the serving endpoint. This field is required.
	Name string `tfsdk:"-" url:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	Prompt any `tfsdk:"prompt"`
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	Stop []string `tfsdk:"stop"`
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	Stream bool `tfsdk:"stream"`
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	Temperature float64 `tfsdk:"temperature"`

	ForceSendFields []string `tfsdk:"-"`
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
	Choices []V1ResponseChoiceElement `tfsdk:"choices"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created int64 `tfsdk:"created"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data []EmbeddingsV1ResponseEmbeddingElement `tfsdk:"data"`
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	Id string `tfsdk:"id"`
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	Model string `tfsdk:"model"`
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	Object QueryEndpointResponseObject `tfsdk:"object"`
	// The predictions returned by the serving endpoint.
	Predictions []any `tfsdk:"predictions"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName string `tfsdk:"-" url:"-" header:"served-model-name,omitempty"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage *ExternalModelUsageElement `tfsdk:"usage"`

	ForceSendFields []string `tfsdk:"-"`
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

// Type always returns QueryEndpointResponseObject to satisfy [pflag.Value] interface
func (f *QueryEndpointResponseObject) Type() string {
	return "QueryEndpointResponseObject"
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int `tfsdk:"calls"`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key RateLimitKey `tfsdk:"key"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod RateLimitRenewalPeriod `tfsdk:"renewal_period"`
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
	ServedModelName string `tfsdk:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int `tfsdk:"traffic_percentage"`
}

type ServedEntityInput struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName string `tfsdk:"entity_name"`
	// The version of the model in Databricks Model Registry to be served or
	// empty if the entity is a FEATURE_SPEC.
	EntityVersion string `tfsdk:"entity_version"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `tfsdk:"environment_vars"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. When an
	// external_model is present, the served entities list can only have one
	// served_entity object. For an existing endpoint with external_model, it
	// can not be updated to an endpoint without external_model. If the endpoint
	// is created without external_model, users cannot update it to add
	// external_model later.
	ExternalModel *ExternalModel `tfsdk:"external_model"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `tfsdk:"instance_profile_arn"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `tfsdk:"max_provisioned_throughput"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `tfsdk:"min_provisioned_throughput"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to
	// <entity-name>-<entity-version>.
	Name string `tfsdk:"name"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `tfsdk:"scale_to_zero_enabled"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size is 0.
	WorkloadSize string `tfsdk:"workload_size"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType string `tfsdk:"workload_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedEntityInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntityOutput struct {
	// The creation timestamp of the served entity in Unix time.
	CreationTimestamp int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the served entity.
	Creator string `tfsdk:"creator"`
	// The name of the entity served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName string `tfsdk:"entity_name"`
	// The version of the served entity in Databricks Model Registry or empty if
	// the entity is a FEATURE_SPEC.
	EntityVersion string `tfsdk:"entity_version"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `tfsdk:"environment_vars"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	ExternalModel *ExternalModel `tfsdk:"external_model"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	FoundationModel *FoundationModel `tfsdk:"foundation_model"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `tfsdk:"instance_profile_arn"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `tfsdk:"max_provisioned_throughput"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `tfsdk:"min_provisioned_throughput"`
	// The name of the served entity.
	Name string `tfsdk:"name"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `tfsdk:"scale_to_zero_enabled"`
	// Information corresponding to the state of the served entity.
	State *ServedModelState `tfsdk:"state"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize string `tfsdk:"workload_size"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType string `tfsdk:"workload_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedEntityOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntitySpec struct {
	// The name of the entity served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName string `tfsdk:"entity_name"`
	// The version of the served entity in Databricks Model Registry or empty if
	// the entity is a FEATURE_SPEC.
	EntityVersion string `tfsdk:"entity_version"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	ExternalModel *ExternalModel `tfsdk:"external_model"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	FoundationModel *FoundationModel `tfsdk:"foundation_model"`
	// The name of the served entity.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedEntitySpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntitySpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `tfsdk:"environment_vars"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn string `tfsdk:"instance_profile_arn"`
	// The name of the model in Databricks Model Registry to be served or if the
	// model resides in Unity Catalog, the full name of model, in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	ModelName string `tfsdk:"model_name"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `tfsdk:"model_version"`
	// The name of a served model. It must be unique across an endpoint. If not
	// specified, this field will default to <model-name>-<model-version>. A
	// served model name can consist of alphanumeric characters, dashes, and
	// underscores.
	Name string `tfsdk:"name"`
	// Whether the compute resources for the served model should scale down to
	// zero.
	ScaleToZeroEnabled bool `tfsdk:"scale_to_zero_enabled"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize ServedModelInputWorkloadSize `tfsdk:"workload_size"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServedModelInputWorkloadType `tfsdk:"workload_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedModelInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	CreationTimestamp int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the served model.
	Creator string `tfsdk:"creator"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `tfsdk:"environment_vars"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn string `tfsdk:"instance_profile_arn"`
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName string `tfsdk:"model_name"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `tfsdk:"model_version"`
	// The name of the served model.
	Name string `tfsdk:"name"`
	// Whether the compute resources for the Served Model should scale down to
	// zero.
	ScaleToZeroEnabled bool `tfsdk:"scale_to_zero_enabled"`
	// Information corresponding to the state of the Served Model.
	State *ServedModelState `tfsdk:"state"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize string `tfsdk:"workload_size"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType string `tfsdk:"workload_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedModelOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelSpec struct {
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName string `tfsdk:"model_name"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `tfsdk:"model_version"`
	// The name of the served model.
	Name string `tfsdk:"name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedModelSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Deployment ServedModelStateDeployment `tfsdk:"deployment"`
	// More information about the state of the served entity, if available.
	DeploymentStateMessage string `tfsdk:"deployment_state_message"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServedModelState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Logs string `tfsdk:"logs"`
}

type ServingEndpoint struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigSummary `tfsdk:"config"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the serving endpoint.
	Creator string `tfsdk:"creator"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `tfsdk:"id"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `tfsdk:"last_updated_timestamp"`
	// The name of the serving endpoint.
	Name string `tfsdk:"name"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `tfsdk:"state"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `tfsdk:"tags"`
	// The task type of the serving endpoint.
	Task string `tfsdk:"task"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServingEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []ServingEndpointPermission `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName string `tfsdk:"display_name"`
	// name of the group
	GroupName string `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName string `tfsdk:"service_principal_name"`
	// name of the user
	UserName string `tfsdk:"user_name"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServingEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointDetailed struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigOutput `tfsdk:"config"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `tfsdk:"creation_timestamp"`
	// The email of the user who created the serving endpoint.
	Creator string `tfsdk:"creator"`
	// Endpoint invocation url if route optimization is enabled for endpoint
	EndpointUrl string `tfsdk:"endpoint_url"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `tfsdk:"id"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `tfsdk:"last_updated_timestamp"`
	// The name of the serving endpoint.
	Name string `tfsdk:"name"`
	// The config that the endpoint is attempting to update to.
	PendingConfig *EndpointPendingConfig `tfsdk:"pending_config"`
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel `tfsdk:"permission_level"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized bool `tfsdk:"route_optimized"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `tfsdk:"state"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `tfsdk:"tags"`
	// The task type of the serving endpoint.
	Task string `tfsdk:"task"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServingEndpointDetailed) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointDetailed) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Inherited bool `tfsdk:"inherited"`

	InheritedFromObject []string `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
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

// Type always returns ServingEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointPermissionLevel) Type() string {
	return "ServingEndpointPermissionLevel"
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse `tfsdk:"access_control_list"`

	ObjectId string `tfsdk:"object_id"`

	ObjectType string `tfsdk:"object_type"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServingEndpointPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsDescription struct {
	Description string `tfsdk:"description"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `tfsdk:"permission_level"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *ServingEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest `tfsdk:"access_control_list"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `tfsdk:"-" url:"-"`
}

type StopAppRequest struct {
	// The name of the app.
	Name string `tfsdk:"-" url:"-"`
}

type StopAppResponse struct {
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes []Route `tfsdk:"routes"`
}

type UpdateAppRequest struct {
	// The description of the app.
	Description string `tfsdk:"description"`
	// The name of the app. The name must contain only lowercase alphanumeric
	// characters and hyphens and be between 2 and 30 characters long. It must
	// be unique within the workspace.
	Name string `tfsdk:"name" url:"-"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *UpdateAppRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAppRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason string `tfsdk:"finishReason"`
	// The index of the choice in the __chat or completions__ response.
	Index int `tfsdk:"index"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs int `tfsdk:"logprobs"`
	// The message response from the __chat__ endpoint.
	Message *ChatMessage `tfsdk:"message"`
	// The text response from the __completions__ endpoint.
	Text string `tfsdk:"text"`

	ForceSendFields []string `tfsdk:"-"`
}

func (s *V1ResponseChoiceElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s V1ResponseChoiceElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
