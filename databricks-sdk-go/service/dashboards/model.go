// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql"
)

type AuthorizationDetails struct {
	// Represents downscoped permission rules with specific access rights. This
	// field is specific to the `workspace_rule_set` constraint.
	GrantRules []AuthorizationDetailsGrantRule `json:"grant_rules,omitempty"`
	// The legacy ACL path of the workspace object that the rule applies to.
	ResourceLegacyAclPath string `json:"resource_legacy_acl_path,omitempty"`
	// The resource name to which the authorization rule applies. This field is
	// specific to the `workspace_rule_set` constraint. Format:
	// `workspaces/{workspace_id}` for a workspace-level rule, or
	// `workspaces/{workspace_id}/dashboards/{dashboard_id}` for a rule on a
	// dashboard.
	ResourceName string `json:"resource_name,omitempty"`
	// The type of authorization downscoping policy. A `workspace_rule_set`
	// policy defines access rules for a specific workspace resource.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AuthorizationDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthorizationDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AuthorizationDetailsGrantRule struct {
	// The permission set that the rule grants on the resource.
	PermissionSet string `json:"permission_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AuthorizationDetailsGrantRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AuthorizationDetailsGrantRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateDashboardRequest struct {
	Dashboard Dashboard `json:"dashboard"`
	// Sets the default catalog for all datasets in this dashboard. Does not
	// impact table references that use fully qualified catalog names (ex:
	// samples.nyctaxi.trips). Leave blank to keep each dataset’s existing
	// configuration.
	DatasetCatalog string `json:"-" url:"dataset_catalog,omitempty"`
	// Sets the default schema for all datasets in this dashboard. Does not
	// impact table references that use fully qualified schema names (ex:
	// nyctaxi.trips). Leave blank to keep each dataset’s existing
	// configuration.
	DatasetSchema string `json:"-" url:"dataset_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The schedule to create. A dashboard is limited to 10 schedules.
	Schedule Schedule `json:"schedule"`
}

func (s *CreateScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// The subscription to create. A schedule is limited to 100 subscriptions.
	Subscription Subscription `json:"subscription"`
}

func (s *CreateSubscriptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CronSchedule struct {
	// A cron expression in Quartz syntax that defines when the schedule runs.
	// See the [Quartz cron trigger tutorial] for the syntax.
	//
	// [Quartz cron trigger tutorial]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

func (s *CronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime string `json:"create_time,omitempty"`
	// UUID identifying the dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The display name of the dashboard. It is also the file name of the
	// dashboard asset in the workspace, so it must be unique within the parent
	// folder, cannot contain `/`, and cannot end with `.json`, `.lvdash` or
	// `.yml`. The display name must not exceed 500 characters.
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag string `json:"etag,omitempty"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. The folder must already exist; when
	// omitted on create, the dashboard is created in the caller's home folder.
	// This field is excluded in List Dashboards responses.
	ParentPath string `json:"parent_path,omitempty"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path string `json:"path,omitempty"`
	// The dashboard definition (datasets, pages and widgets) as a JSON string,
	// in the same format as a dashboard file exported from the workspace
	// (`.lvdash.json`). When omitted on create, an empty dashboard is created;
	// when omitted on update, the existing definition is kept and only the
	// other fields change. Use the [get dashboard API] on an existing dashboard
	// to see a complete example. This field is excluded in List Dashboards
	// responses.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard string `json:"serialized_dashboard,omitempty"`
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	UpdateTime string `json:"update_time,omitempty"`
	// The warehouse ID used to run the dashboard.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DashboardView string

const DashboardViewDashboardViewBasic DashboardView = `DASHBOARD_VIEW_BASIC`

// String representation for [fmt.Print]
func (f *DashboardView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardView) Set(v string) error {
	switch v {
	case `DASHBOARD_VIEW_BASIC`:
		*f = DashboardView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD_VIEW_BASIC"`, v)
	}
}

// Values returns all possible values for DashboardView.
//
// There is no guarantee on the order of the values in the slice.
func (f *DashboardView) Values() []DashboardView {
	return []DashboardView{
		DashboardViewDashboardViewBasic,
	}
}

// Type always returns DashboardView to satisfy [pflag.Value] interface
func (f *DashboardView) Type() string {
	return "DashboardView"
}

type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag string `json:"-" url:"etag,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteScheduleRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// The etag of the subscription, as returned when it was read. If it is
	// provided and no longer matches, the request fails with ABORTED and the
	// subscription is not deleted.
	Etag string `json:"-" url:"etag,omitempty"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteSubscriptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DownloadMessageAttachmentVisualizationRequest struct {
	// The resource name of the attachment to render, in the format
	// `spaces/{space_id}/conversations/{conversation_id}/messages/{message_id}/attachments/{attachment_id}`.
	Name string `json:"-" url:"-"`
}

func (s *DownloadMessageAttachmentVisualizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DownloadMessageAttachmentVisualizationResponse struct {
	Contents io.ReadCloser `json:"-"`
}

func (s *DownloadMessageAttachmentVisualizationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EvaluationStatusType string

const EvaluationStatusTypeDone EvaluationStatusType = `DONE`

const EvaluationStatusTypeEvaluationCancelled EvaluationStatusType = `EVALUATION_CANCELLED`

const EvaluationStatusTypeEvaluationFailed EvaluationStatusType = `EVALUATION_FAILED`

const EvaluationStatusTypeEvaluationTimeout EvaluationStatusType = `EVALUATION_TIMEOUT`

const EvaluationStatusTypeNotStarted EvaluationStatusType = `NOT_STARTED`

const EvaluationStatusTypeRunning EvaluationStatusType = `RUNNING`

// String representation for [fmt.Print]
func (f *EvaluationStatusType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EvaluationStatusType) Set(v string) error {
	switch v {
	case `DONE`, `EVALUATION_CANCELLED`, `EVALUATION_FAILED`, `EVALUATION_TIMEOUT`, `NOT_STARTED`, `RUNNING`:
		*f = EvaluationStatusType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DONE", "EVALUATION_CANCELLED", "EVALUATION_FAILED", "EVALUATION_TIMEOUT", "NOT_STARTED", "RUNNING"`, v)
	}
}

// Values returns all possible values for EvaluationStatusType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EvaluationStatusType) Values() []EvaluationStatusType {
	return []EvaluationStatusType{
		EvaluationStatusTypeDone,
		EvaluationStatusTypeEvaluationCancelled,
		EvaluationStatusTypeEvaluationFailed,
		EvaluationStatusTypeEvaluationTimeout,
		EvaluationStatusTypeNotStarted,
		EvaluationStatusTypeRunning,
	}
}

// Type always returns EvaluationStatusType to satisfy [pflag.Value] interface
func (f *EvaluationStatusType) Type() string {
	return "EvaluationStatusType"
}

// One part of Genie's response to a message. Exactly one of the attachment
// fields is set.
type GenieAttachment struct {
	// The ID of the attachment. The attachment query-result, execute-query,
	// download, and visualization endpoints take this ID.
	AttachmentId string `json:"attachment_id,omitempty"`
	// Query Attachment if Genie responds with a SQL query
	Query *GenieQueryAttachment `json:"query,omitempty"`
	// Follow-up questions suggested by Genie
	SuggestedQuestions *GenieSuggestedQuestionsAttachment `json:"suggested_questions,omitempty"`
	// Text Attachment if Genie responds with text This also contains the final
	// summary when available.
	Text *TextAttachment `json:"text,omitempty"`
	// Visualization generated by Genie, if requested via `enable_visualization`
	Viz *GenieVizAttachment `json:"viz,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to cancel an in-flight agent-mode response.
type GenieCancelResponseRequest struct {
	// The ID of the Genie agent (synonymous with the Genie space ID).
	AgentId string `json:"-" url:"-"`
	// The ID of the conversation containing the response.
	ConversationId string `json:"-" url:"-"`
	// The ID of the response to cancel (the id from the `response.created`
	// event).
	ResponseId string `json:"-" url:"-"`
}

func (s *GenieCancelResponseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A Genie conversation. Use chat-mode message endpoints for classic chats and
// agent-mode response and item endpoints for agent conversations. Conversation
// management, feedback, comments, and attachment operations support both modes.
type GenieConversation struct {
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Time when the conversation was created, in milliseconds since the Unix
	// epoch.
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Conversation ID. Legacy identifier, use conversation_id instead
	Id string `json:"id,omitempty"`
	// Time when the conversation was last updated, in milliseconds since the
	// Unix epoch.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// Title of the conversation. Genie sets it to the text of the first
	// message, truncated to 255 characters.
	Title string `json:"title"`
	// ID of the user who created the conversation
	UserId int `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieConversation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieConversation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A summary of a conversation, as returned when listing the conversations in a
// space.
type GenieConversationSummary struct {
	// Whether this is a classic chat or an agent-mode conversation. Allows
	// callers to route message retrieval (chat vs. agent endpoint) without an
	// extra lookup.
	AgentType GenieConversationType `json:"agent_type,omitempty"`
	// The ID of the conversation.
	ConversationId string `json:"conversation_id"`
	// Time when the conversation was created, in milliseconds since the Unix
	// epoch.
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// The title of the conversation.
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieConversationSummary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieConversationSummary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of a Genie conversation. Distinguishes an agent-mode conversation
// from a classic chat conversation so callers can route message retrieval
// accordingly without a per-conversation lookup.
type GenieConversationType string

const GenieConversationTypeGenieConversationTypeAgent GenieConversationType = `GENIE_CONVERSATION_TYPE_AGENT`

const GenieConversationTypeGenieConversationTypeChat GenieConversationType = `GENIE_CONVERSATION_TYPE_CHAT`

// String representation for [fmt.Print]
func (f *GenieConversationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieConversationType) Set(v string) error {
	switch v {
	case `GENIE_CONVERSATION_TYPE_AGENT`, `GENIE_CONVERSATION_TYPE_CHAT`:
		*f = GenieConversationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENIE_CONVERSATION_TYPE_AGENT", "GENIE_CONVERSATION_TYPE_CHAT"`, v)
	}
}

// Values returns all possible values for GenieConversationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieConversationType) Values() []GenieConversationType {
	return []GenieConversationType{
		GenieConversationTypeGenieConversationTypeAgent,
		GenieConversationTypeGenieConversationTypeChat,
	}
}

// Type always returns GenieConversationType to satisfy [pflag.Value] interface
func (f *GenieConversationType) Type() string {
	return "GenieConversationType"
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content string `json:"content"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// Enable visualization generation.
	EnableVisualization bool `json:"enable_visualization,omitempty"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieCreateConversationMessageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieCreateConversationMessageRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieCreateEvalRunRequest struct {
	// List of benchmark question IDs to evaluate. These questions must exist in
	// the specified Genie space. If none are specified, then all benchmark
	// questions are evaluated.
	BenchmarkQuestionIds []string `json:"benchmark_question_ids,omitempty"`
	// The ID associated with the Genie space where the evaluations will be
	// executed.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieCreateEvalRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieCreateMessageCommentRequest struct {
	// Comment text content.
	Content string `json:"content"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the message.
	MessageId string `json:"-" url:"-"`
	// The ID associated with the Genie space.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieCreateMessageCommentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieCreateSpaceRequest struct {
	// Description of the space, up to 16384 bytes.
	Description string `json:"description,omitempty"`
	// Path of the workspace folder in which to create the space.
	ParentPath string `json:"parent_path,omitempty"`
	// Configuration of the new space as a JSON string. The configuration must
	// include at least one table, metric view, volume, MCP connection, or
	// sub-agent data source. To copy an existing space, use the value that [Get
	// Genie Space](:method:genie/getspace) returns when
	// `include_serialized_space` is true. For the JSON format and its
	// validation rules, see
	// https://docs.databricks.com/aws/en/genie-agents/conversation-api.
	SerializedSpace string `json:"serialized_space"`
	// Display name of the new space. A default name is assigned when omitted.
	Title string `json:"title,omitempty"`
	// ID of the SQL warehouse that runs the queries Genie generates for the
	// space, as a 16-character hexadecimal string.
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieCreateSpaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieCreateSpaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieDeleteConversationMessageRequest struct {
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the message to delete.
	MessageId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the message is located.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieDeleteConversationMessageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieDeleteConversationRequest struct {
	// The ID of the conversation to delete.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the conversation is located.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieDeleteConversationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieEvalAssessment string

const GenieEvalAssessmentBad GenieEvalAssessment = `BAD`

const GenieEvalAssessmentGood GenieEvalAssessment = `GOOD`

const GenieEvalAssessmentNeedsReview GenieEvalAssessment = `NEEDS_REVIEW`

// String representation for [fmt.Print]
func (f *GenieEvalAssessment) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieEvalAssessment) Set(v string) error {
	switch v {
	case `BAD`, `GOOD`, `NEEDS_REVIEW`:
		*f = GenieEvalAssessment(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BAD", "GOOD", "NEEDS_REVIEW"`, v)
	}
}

// Values returns all possible values for GenieEvalAssessment.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieEvalAssessment) Values() []GenieEvalAssessment {
	return []GenieEvalAssessment{
		GenieEvalAssessmentBad,
		GenieEvalAssessmentGood,
		GenieEvalAssessmentNeedsReview,
	}
}

// Type always returns GenieEvalAssessment to satisfy [pflag.Value] interface
func (f *GenieEvalAssessment) Type() string {
	return "GenieEvalAssessment"
}

type GenieEvalResponse struct {
	// The content of this entry is the SQL statement when `response_type` is
	// SQL, or Genie's text answer when it is TEXT.
	Response string `json:"response,omitempty"`
	// Whether this entry is a SQL statement or a text answer.
	ResponseType GenieEvalResponseType `json:"response_type,omitempty"`
	// Result of executing the SQL statement, only set when `response_type` is
	// SQL. The value has the structure of a [Statement Execution API response].
	//
	// [Statement Execution API response]: https://docs.databricks.com/api/statement-execution/v1/get-statement-result
	SqlExecutionResult *sql.StatementResponse `json:"sql_execution_result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieEvalResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieEvalResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieEvalResponseType string

const GenieEvalResponseTypeSql GenieEvalResponseType = `SQL`

const GenieEvalResponseTypeText GenieEvalResponseType = `TEXT`

// String representation for [fmt.Print]
func (f *GenieEvalResponseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieEvalResponseType) Set(v string) error {
	switch v {
	case `SQL`, `TEXT`:
		*f = GenieEvalResponseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SQL", "TEXT"`, v)
	}
}

// Values returns all possible values for GenieEvalResponseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieEvalResponseType) Values() []GenieEvalResponseType {
	return []GenieEvalResponseType{
		GenieEvalResponseTypeSql,
		GenieEvalResponseTypeText,
	}
}

// Type always returns GenieEvalResponseType to satisfy [pflag.Value] interface
func (f *GenieEvalResponseType) Type() string {
	return "GenieEvalResponseType"
}

// Shows summary information for an evaluation result. For detailed information
// including SQL execution results, actual/expected responses, and assessment
// scores, use GenieGetEvalResultDetails.
type GenieEvalResult struct {
	// Benchmark SQL statement that defines the expected answer, as it was when
	// the run started.
	BenchmarkAnswer string `json:"benchmark_answer,omitempty"`
	// The ID of the benchmark question that was evaluated.
	BenchmarkQuestionId string `json:"benchmark_question_id"`
	// ID of the user who started the evaluation run.
	CreatedByUser int64 `json:"created_by_user,omitempty"`
	// Text of the benchmark question as it was when the run started. Later
	// edits to the question do not change it.
	Question string `json:"question,omitempty"`
	// Unique identifier for this evaluation result.
	ResultId string `json:"result_id"`
	// The ID of the space the evaluation result belongs to.
	SpaceId string `json:"space_id"`
	// Current status of this result. NOT_STARTED and RUNNING mean the question
	// has not finished. DONE means Genie answered and the response was
	// assessed. EVALUATION_FAILED means Genie could not answer the question or
	// the benchmark SQL could not run. EVALUATION_CANCELLED and
	// EVALUATION_TIMEOUT mean the run was cancelled or its background job
	// stopped before this question finished.
	Status EvaluationStatusType `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieEvalResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieEvalResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Shows detailed information for an evaluation result.
type GenieEvalResultDetails struct {
	// Genie's response to the benchmark question. A SQL entry carries the
	// generated statement and its execution result, and a TEXT entry carries
	// Genie's text answer. The list is empty when Genie produced no response.
	ActualResponse []GenieEvalResponse `json:"actual_response,omitempty"`
	// Outcome of comparing Genie's response with the benchmark. GOOD means the
	// results matched, BAD means they differed, and NEEDS_REVIEW means the
	// comparison was inconclusive and a person should review the result. The
	// value is only meaningful once the result status is DONE.
	Assessment GenieEvalAssessment `json:"assessment,omitempty"`
	// Reasons for the assessment score.
	//
	// Assessment reasons describe why a Genie response was scored as BAD.
	//
	// Deterministic values (compared against the ground truth result): -
	// EMPTY_RESULT: Genie's generated SQL results were empty for this benchmark
	// question. - RESULT_MISSING_ROWS: Genie's generated SQL response is
	// missing rows from the provided ground truth SQL. - RESULT_EXTRA_ROWS:
	// Genie's generated SQL response has more rows than the provided ground
	// truth SQL. - RESULT_MISSING_COLUMNS: Genie's generated SQL response is
	// missing columns from the provided ground truth SQL. -
	// RESULT_EXTRA_COLUMNS: Genie's generated SQL response has more columns
	// than the provided ground truth SQL. - SINGLE_CELL_DIFFERENCE: Single
	// value result was produced but differs from ground truth result. -
	// EMPTY_GOOD_SQL: The benchmark SQL returned an empty result. -
	// COLUMN_TYPE_DIFFERENCE: The values between the results match but the
	// column type is different.
	//
	// LLM judge ratings explain the factors driving BAD results: -
	// LLM_JUDGE_MISSING_OR_INCORRECT_FILTER: Genie's generated SQL is missing a
	// WHERE clause condition or has incorrect filter logic that
	// excludes/includes wrong data. - LLM_JUDGE_INCOMPLETE_OR_PARTIAL_OUTPUT:
	// Genie's generated SQL returns only some of the requested data or columns,
	// missing parts of what the ground truth SQL returns. -
	// LLM_JUDGE_MISINTERPRETATION_OF_USER_REQUEST: Genie's generated SQL
	// fundamentally misunderstands what the user is asking for, addressing the
	// wrong question or goal. -
	// LLM_JUDGE_INSTRUCTION_COMPLIANCE_OR_MISSING_BUSINESS_LOGIC: Genie's
	// generated SQL fails to apply specified instructions or business logic
	// that should be followed. - LLM_JUDGE_INCORRECT_METRIC_CALCULATION:
	// Genie's generated SQL uses incorrect logic or makes wrong assumptions
	// when calculating metrics. - LLM_JUDGE_INCORRECT_TABLE_OR_FIELD_USAGE:
	// Genie's generated SQL references wrong tables, columns, or uses fields
	// that don't match the ground truth SQL's intent. -
	// LLM_JUDGE_INCORRECT_FUNCTION_USAGE: Genie's generated SQL uses SQL
	// functions incorrectly or inappropriately (wrong parameters, wrong
	// function for the task, etc.). - LLM_JUDGE_MISSING_OR_INCORRECT_JOIN:
	// Genie's generated SQL is missing necessary joins between tables or has
	// incorrect join conditions/types that produce wrong results. -
	// LLM_JUDGE_MISSING_OR_INCORRECT_AGGREGATION: Genie's generated SQL is
	// missing GROUP BY clauses or has incorrect grouping that doesn't match the
	// requested aggregation level. - LLM_JUDGE_FORMATTING_ERROR: Genie's
	// generated SQL output has incorrect formatting, ordering (ORDER BY), or
	// presentation issues that don't match expectations. - LLM_JUDGE_OTHER: LLM
	// judge identified an error that doesn't fall into other categories.
	//
	// Deprecated LLM judge values (kept for backward compatibility, do not
	// use): - LLM_JUDGE_MISSING_JOIN (deprecated) - LLM_JUDGE_WRONG_FILTER
	// (deprecated) - LLM_JUDGE_WRONG_AGGREGATION (deprecated) -
	// LLM_JUDGE_WRONG_COLUMNS (deprecated) - LLM_JUDGE_SYNTAX_ERROR
	// (deprecated) - LLM_JUDGE_SEMANTIC_ERROR (deprecated)
	AssessmentReasons []ScoreReason `json:"assessment_reasons,omitempty"`
	// The ID of the benchmark question that was evaluated.
	BenchmarkQuestionId string `json:"benchmark_question_id"`
	// Current status of this evaluation result, not of the whole run. Once the
	// status is DONE, the response and automatic assessment are complete. A
	// user can still change the assessment manually.
	EvalRunStatus EvaluationStatusType `json:"eval_run_status,omitempty"`
	// The expected answer from the benchmark question: one SQL entry with the
	// benchmark SQL statement and its execution result.
	ExpectedResponse []GenieEvalResponse `json:"expected_response,omitempty"`
	// Whether a user set the assessment manually after reviewing the result. A
	// manual assessment replaces the automatic one and clears
	// `assessment_reasons`.
	ManualAssessment bool `json:"manual_assessment,omitempty"`
	// The unique identifier for the evaluation result.
	ResultId string `json:"result_id"`
	// The ID of the space the evaluation result belongs to.
	SpaceId string `json:"space_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieEvalResultDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieEvalResultDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A benchmark evaluation run. A run asks Genie each selected benchmark question
// in a Genie space, compares every response with the benchmark SQL, and records
// one evaluation result per question. Runs execute asynchronously; poll the run
// until its status is terminal, then list its results. All evaluation run
// methods require CAN EDIT permission on the space. The public benchmark API
// currently evaluates chat-mode responses.
type GenieEvalRunResponse struct {
	// Timestamp when the evaluation run was created (milliseconds since epoch).
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// The unique identifier for the evaluation run.
	EvalRunId string `json:"eval_run_id"`
	// Current status of the run, derived from the statuses of its results.
	// Unless a result has timed out or the run has been cancelled, the run is
	// RUNNING while any question has not finished. It is DONE once every
	// question has finished, whether it was scored or failed.
	// EVALUATION_CANCELLED means the run was cancelled before all questions
	// finished, and EVALUATION_TIMEOUT means the background evaluation job
	// stopped before they did. DONE, EVALUATION_CANCELLED and
	// EVALUATION_TIMEOUT are terminal run statuses. NOT_STARTED and
	// EVALUATION_FAILED appear only on individual results.
	EvalRunStatus EvaluationStatusType `json:"eval_run_status,omitempty"`
	// Timestamp when the evaluation run was last updated (milliseconds since
	// epoch).
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Number of questions whose result is assessed GOOD, including results
	// assessed manually.
	NumCorrect int64 `json:"num_correct,omitempty"`
	// Number of questions that have finished, either scored (DONE) or failed
	// (EVALUATION_FAILED). Cancelled and timed-out questions are not counted.
	NumDone int64 `json:"num_done,omitempty"`
	// Number of finished questions whose result is assessed NEEDS_REVIEW and
	// has not been reviewed manually yet.
	NumNeedsReview int64 `json:"num_needs_review,omitempty"`
	// Total number of questions in the evaluation run.
	NumQuestions int64 `json:"num_questions,omitempty"`
	// User ID who initiated the evaluation run.
	RunByUser int64 `json:"run_by_user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieEvalRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieEvalRunResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieExecuteMessageAttachmentQueryRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieExecuteMessageAttachmentQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieExecuteMessageQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Feedback containing rating and optional comment
type GenieFeedback struct {
	// Optional feedback comment text
	Comment string `json:"comment,omitempty"`
	// The feedback rating
	Rating GenieFeedbackRating `json:"rating,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieFeedback) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieFeedback) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Feedback rating for Genie messages
type GenieFeedbackRating string

const GenieFeedbackRatingNegative GenieFeedbackRating = `NEGATIVE`

const GenieFeedbackRatingNone GenieFeedbackRating = `NONE`

const GenieFeedbackRatingPositive GenieFeedbackRating = `POSITIVE`

// String representation for [fmt.Print]
func (f *GenieFeedbackRating) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GenieFeedbackRating) Set(v string) error {
	switch v {
	case `NEGATIVE`, `NONE`, `POSITIVE`:
		*f = GenieFeedbackRating(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NEGATIVE", "NONE", "POSITIVE"`, v)
	}
}

// Values returns all possible values for GenieFeedbackRating.
//
// There is no guarantee on the order of the values in the slice.
func (f *GenieFeedbackRating) Values() []GenieFeedbackRating {
	return []GenieFeedbackRating{
		GenieFeedbackRatingNegative,
		GenieFeedbackRatingNone,
		GenieFeedbackRatingPositive,
	}
}

// Type always returns GenieFeedbackRating to satisfy [pflag.Value] interface
func (f *GenieFeedbackRating) Type() string {
	return "GenieFeedbackRating"
}

type GenieGenerateDownloadFullQueryResultRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGenerateDownloadFullQueryResultRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGenerateDownloadFullQueryResultResponse struct {
	// Download ID. Use this ID to track the download request in subsequent
	// polling calls
	DownloadId string `json:"download_id,omitempty"`
	// JWT signature for the download_id to ensure secure access to query
	// results
	DownloadIdSignature string `json:"download_id_signature,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieGenerateDownloadFullQueryResultResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieGenerateDownloadFullQueryResultResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieGetConversationMessageRequest struct {
	// The ID associated with the target conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetConversationMessageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetDownloadFullQueryResultRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Download ID. This ID is provided by the [Generate Download
	// endpoint](:method:genie/generateDownloadFullQueryResult)
	DownloadId string `json:"-" url:"-"`
	// JWT signature for the download_id to ensure secure access to query
	// results
	DownloadIdSignature string `json:"-" url:"download_id_signature"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetDownloadFullQueryResultRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetDownloadFullQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
}

func (s *GenieGetDownloadFullQueryResultResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetEvalResultDetailsRequest struct {
	// The unique identifier for the evaluation run.
	EvalRunId string `json:"-" url:"-"`
	// The unique identifier for the evaluation result.
	ResultId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the evaluation run is
	// located.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetEvalResultDetailsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetEvalRunRequest struct {
	// The ID of the evaluation run, as returned when the run was created.
	EvalRunId string `json:"-" url:"-"`
	// The ID associated with the Genie space where the evaluation run is
	// located.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetEvalRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetMessageAttachmentQueryResultRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetMessageAttachmentQueryResultRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetMessageQueryResultRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponse `json:"statement_response,omitempty"`
}

func (s *GenieGetMessageQueryResultResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetQueryResultByAttachmentRequest struct {
	// Attachment ID
	AttachmentId string `json:"-" url:"-"`
	// Conversation ID
	ConversationId string `json:"-" url:"-"`
	// Message ID
	MessageId string `json:"-" url:"-"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieGetQueryResultByAttachmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieGetSpaceRequest struct {
	// Whether to include the serialized space export in the response. When
	// true, the response also carries the configuration `etag`. Requires at
	// least CAN EDIT permission on the space.
	IncludeSerializedSpace bool `json:"-" url:"include_serialized_space,omitempty"`
	// The ID associated with the Genie space
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieGetSpaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieGetSpaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationCommentsRequest struct {
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// Maximum number of comments to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token for getting the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID associated with the Genie space.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationCommentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationCommentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationCommentsResponse struct {
	// List of comments in the conversation.
	Comments []GenieMessageComment `json:"comments,omitempty"`
	// Token to get the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationCommentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationCommentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationMessagesRequest struct {
	// The ID of the conversation to list messages from
	ConversationId string `json:"-" url:"-"`
	// Maximum number of messages to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token to get the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID associated with the Genie space where the conversation is located
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationMessagesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationMessagesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationMessagesResponse struct {
	// List of messages in the conversation.
	Messages []GenieMessage `json:"messages,omitempty"`
	// The token to use for retrieving the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationMessagesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationMessagesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationsRequest struct {
	// Include all conversations in the space across all users. Requires at
	// least CAN MANAGE permission on the space.
	IncludeAll bool `json:"-" url:"include_all,omitempty"`
	// Maximum number of conversations to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token to get the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID of the Genie space to retrieve conversations from.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListConversationsResponse struct {
	// List of conversations in the Genie space
	Conversations []GenieConversationSummary `json:"conversations,omitempty"`
	// Token to get the next page of results
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListConversationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListConversationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListEvalResultsRequest struct {
	// The unique identifier for the evaluation run.
	EvalRunId string `json:"-" url:"-"`
	// Maximum number of eval results to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Opaque token to retrieve the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID associated with the Genie space where the evaluation run is
	// located.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListEvalResultsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListEvalResultsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListEvalResultsResponse struct {
	// List of evaluation results for the specified run.
	EvalResults []GenieEvalResult `json:"eval_results,omitempty"`
	// The token to use for retrieving the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListEvalResultsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListEvalResultsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListEvalRunsRequest struct {
	// Maximum number of evaluation runs to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token to get the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID associated with the Genie space where the evaluation run is
	// located.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListEvalRunsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListEvalRunsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListEvalRunsResponse struct {
	// List of evaluation runs for a space on provided page token and page size
	EvalRuns []GenieEvalRunResponse `json:"eval_runs,omitempty"`
	// The token to use for retrieving the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListEvalRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListEvalRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListMessageCommentsRequest struct {
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the message.
	MessageId string `json:"-" url:"-"`
	// Maximum number of comments to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token for getting the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The ID associated with the Genie space.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListMessageCommentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListMessageCommentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListMessageCommentsResponse struct {
	// List of comments on the message.
	Comments []GenieMessageComment `json:"comments,omitempty"`
	// Token to get the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListMessageCommentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListMessageCommentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListSpacesRequest struct {
	// Maximum number of spaces to return per page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token for getting the next page of results
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListSpacesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListSpacesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieListSpacesResponse struct {
	// Token to get the next page of results
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Genie spaces
	Spaces []GenieSpace `json:"spaces,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieListSpacesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieListSpacesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieMessage struct {
	// Genie's response to the message. Each attachment holds one part of the
	// response, such as a text answer or a SQL query.
	Attachments []GenieAttachment `json:"attachments,omitempty"`
	// User message content
	Content string `json:"content"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Time when the message was created, in milliseconds since the Unix epoch.
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Error message if Genie failed to respond to the message
	Error *MessageError `json:"error,omitempty"`
	// User feedback for the message if provided
	Feedback *GenieFeedback `json:"feedback,omitempty"`
	// Message ID. Legacy identifier, use message_id instead
	Id string `json:"id,omitempty"`
	// Time when the message was last updated, in milliseconds since the Unix
	// epoch.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Message ID
	MessageId string `json:"message_id"`
	// The result of SQL query if the message includes a query attachment.
	// Deprecated. Use `query_result_metadata` in `GenieQueryAttachment`
	// instead.
	QueryResult *Result `json:"query_result,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`

	Status MessageStatus `json:"status,omitempty"`
	// ID of the user who created the message
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A comment on a Genie conversation message.
type GenieMessageComment struct {
	// Comment text content
	Content string `json:"content"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// Timestamp when the comment was created
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Comment ID
	MessageCommentId string `json:"message_comment_id"`
	// Message ID
	MessageId string `json:"message_id"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// ID of the user who created the comment
	UserId int64 `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieMessageComment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieMessageComment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieQueryAttachment struct {
	// Description of the query
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`
	// Time when the user updated the query last
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The parameters of the SQL query, when the query is parameterized.
	Parameters []QueryAttachmentParameter `json:"parameters,omitempty"`
	// AI generated SQL query
	Query string `json:"query,omitempty"`
	// Metadata associated with the query result.
	QueryResultMetadata *GenieResultMetadata `json:"query_result_metadata,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`
	// Insights into how Genie came to generate the SQL.
	Thoughts []Thought `json:"thoughts,omitempty"`
	// Name of the query
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieQueryAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieQueryAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieResultMetadata struct {
	// Indicates whether the result set is truncated.
	IsTruncated bool `json:"is_truncated,omitempty"`
	// The number of rows in the result set.
	RowCount int64 `json:"row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieResultMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieResultMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieSendMessageFeedbackRequest struct {
	// Optional text feedback that will be stored as a comment.
	Comment string `json:"comment,omitempty"`
	// The ID associated with the conversation.
	ConversationId string `json:"-" url:"-"`
	// The ID associated with the message to provide feedback for.
	MessageId string `json:"-" url:"-"`
	// The rating (POSITIVE, NEGATIVE, or NONE).
	Rating GenieFeedbackRating `json:"rating"`
	// The ID associated with the Genie space where the message is located.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieSendMessageFeedbackRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieSendMessageFeedbackRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A Genie space that business users query in natural language. Its
// configuration (data sources, instructions, sample questions and benchmarks)
// travels as a JSON string in `serialized_space`; the other fields are metadata
// about the space.
type GenieSpace struct {
	// Time when the Genie space was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Description of the Genie Space
	Description string `json:"description,omitempty"`
	// ETag for this space. Pass this value back in the update request to
	// prevent overwriting concurrent changes.
	Etag string `json:"etag,omitempty"`
	// Path of the workspace folder that contains the space.
	ParentPath string `json:"parent_path,omitempty"`
	// Configuration of the space (data sources, instructions, sample questions
	// and benchmarks) as a JSON string. [Get Genie
	// Space](:method:genie/getspace) returns it only when
	// `include_serialized_space` is true, and List Genie spaces never returns
	// it. For the JSON format and its validation rules, see
	// https://docs.databricks.com/aws/en/genie-agents/conversation-api.
	SerializedSpace string `json:"serialized_space,omitempty"`
	// Genie space ID
	SpaceId string `json:"space_id"`
	// Title of the Genie Space
	Title string `json:"title"`
	// Time when the Genie space was last modified, matching the value shown in
	// the Genie Agent UI.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// ID of the SQL warehouse that runs the queries Genie generates for this
	// space.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieSpace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieSpace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content string `json:"content"`
	// Enable visualization generation.
	EnableVisualization bool `json:"enable_visualization,omitempty"`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieStartConversationMessageRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieStartConversationMessageRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieStartConversationResponse struct {
	// The newly created conversation.
	Conversation *GenieConversation `json:"conversation,omitempty"`
	// Conversation ID
	ConversationId string `json:"conversation_id"`
	// The first message of the conversation, in its initial state.
	Message *GenieMessage `json:"message,omitempty"`
	// Message ID
	MessageId string `json:"message_id"`
}

func (s *GenieStartConversationResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Follow-up questions suggested by Genie
type GenieSuggestedQuestionsAttachment struct {
	// The suggested follow-up questions
	Questions []string `json:"questions,omitempty"`
}

func (s *GenieSuggestedQuestionsAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieTrashSpaceRequest struct {
	// The ID associated with the Genie space to be sent to the trash.
	SpaceId string `json:"-" url:"-"`
}

func (s *GenieTrashSpaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GenieUpdateSpaceRequest struct {
	// New description for the space, up to 16384 bytes.
	Description string `json:"description,omitempty"`
	// ETag returned by a previous GET or UPDATE. When set, the update will fail
	// if the space has been modified since. Omit to apply the update
	// unconditionally.
	Etag string `json:"etag,omitempty"`
	// Parent workspace folder path to move this Genie space under.
	ParentPath string `json:"parent_path,omitempty"`
	// New configuration of the space as a JSON string. It replaces the whole
	// configuration; when omitted, the current configuration is kept. Start
	// from the value that [Get Genie Space](:method:genie/getspace) returns
	// when `include_serialized_space` is true. For the JSON format and its
	// validation rules, see
	// https://docs.databricks.com/aws/en/genie-agents/conversation-api.
	SerializedSpace string `json:"serialized_space,omitempty"`
	// Genie space ID
	SpaceId string `json:"-" url:"-"`
	// New display name for the space. Cannot be empty or contain a slash.
	Title string `json:"title,omitempty"`
	// ID of the SQL warehouse that runs the queries Genie generates for the
	// space.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieUpdateSpaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieUpdateSpaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Visualization generated by Genie for a query result. Use the attachment ID
// with the download visualization API to retrieve the rendered image.
type GenieVizAttachment struct {
	// The ID of the query attachment the visualization was generated from
	QueryAttachmentId string `json:"query_attachment_id,omitempty"`
	// Name of the visualization
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenieVizAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenieVizAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

func (s *GetDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

func (s *GetPublishedDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetPublishedDashboardTokenInfoRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
	// Value to carry in the token for this viewer, such as a tenant or region
	// key. Dashboard dataset queries can reference it as
	// `__aibi_external_value` to filter data per viewer. The combined size of
	// `external_value` and `external_viewer_id` must be less than 1 KB.
	ExternalValue string `json:"-" url:"external_value,omitempty"`
	// Identifier of the viewer in the external application. It is recorded in
	// dashboard audit logs, so it should be unique per viewer and must not
	// contain personally identifiable information.
	ExternalViewerId string `json:"-" url:"external_viewer_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedDashboardTokenInfoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedDashboardTokenInfoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedDashboardTokenInfoResponse struct {
	// Authorization constraints that restrict the minted token to viewing the
	// published dashboard. Pass them to the OAuth token endpoint serialized as
	// a JSON array.
	AuthorizationDetails []AuthorizationDetails `json:"authorization_details,omitempty"`
	// Custom claim that carries `external_value` and `external_viewer_id` into
	// the minted token. It is only present when at least one of those query
	// parameters was provided. Format:
	// `urn:aibi:external_data:<external_value>:<external_viewer_id>:<dashboard_id>`
	CustomClaim string `json:"custom_claim,omitempty"`
	// The space-separated OAuth scopes to request when minting the token.
	Scope string `json:"scope,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedDashboardTokenInfoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedDashboardTokenInfoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

func (s *GetScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"-" url:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"-" url:"-"`
}

func (s *GetSubscriptionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type LifecycleState string

const LifecycleStateActive LifecycleState = `ACTIVE`

const LifecycleStateTrashed LifecycleState = `TRASHED`

// String representation for [fmt.Print]
func (f *LifecycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LifecycleState) Set(v string) error {
	switch v {
	case `ACTIVE`, `TRASHED`:
		*f = LifecycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "TRASHED"`, v)
	}
}

// Values returns all possible values for LifecycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *LifecycleState) Values() []LifecycleState {
	return []LifecycleState{
		LifecycleStateActive,
		LifecycleStateTrashed,
	}
}

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed bool `json:"-" url:"show_trashed,omitempty"`
	// `DASHBOARD_VIEW_BASIC` only includes summary metadata from the dashboard.
	View DashboardView `json:"-" url:"view,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDashboardsResponse struct {
	Dashboards []Dashboard `json:"dashboards,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDashboardsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedules belong.
	DashboardId string `json:"-" url:"-"`
	// The number of schedules to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchedulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken string `json:"next_page_token,omitempty"`

	Schedules []Schedule `json:"schedules,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSchedulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSchedulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard to which the subscriptions belong.
	DashboardId string `json:"-" url:"-"`
	// The number of subscriptions to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// UUID identifying the schedule to which the subscriptions belong.
	ScheduleId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken string `json:"next_page_token,omitempty"`

	Subscriptions []Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MessageError struct {
	// A human-readable description of the error.
	Error string `json:"error,omitempty"`
	// The category of the error.
	Type MessageErrorType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MessageError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MessageError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MessageErrorType string

const MessageErrorTypeBlockMultipleExecutionsException MessageErrorType = `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`

const MessageErrorTypeChatCompletionClientException MessageErrorType = `CHAT_COMPLETION_CLIENT_EXCEPTION`

const MessageErrorTypeChatCompletionClientTimeoutException MessageErrorType = `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`

const MessageErrorTypeChatCompletionNetworkException MessageErrorType = `CHAT_COMPLETION_NETWORK_EXCEPTION`

const MessageErrorTypeContentFilterException MessageErrorType = `CONTENT_FILTER_EXCEPTION`

const MessageErrorTypeContextExceededException MessageErrorType = `CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeCouldNotGetDashboardSchemaException MessageErrorType = `COULD_NOT_GET_DASHBOARD_SCHEMA_EXCEPTION`

const MessageErrorTypeCouldNotGetModelDeploymentsException MessageErrorType = `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`

const MessageErrorTypeCouldNotGetUcSchemaException MessageErrorType = `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`

const MessageErrorTypeDeploymentNotFoundException MessageErrorType = `DEPLOYMENT_NOT_FOUND_EXCEPTION`

const MessageErrorTypeDescribeQueryInvalidSqlError MessageErrorType = `DESCRIBE_QUERY_INVALID_SQL_ERROR`

const MessageErrorTypeDescribeQueryTimeout MessageErrorType = `DESCRIBE_QUERY_TIMEOUT`

const MessageErrorTypeDescribeQueryUnexpectedFailure MessageErrorType = `DESCRIBE_QUERY_UNEXPECTED_FAILURE`

const MessageErrorTypeExceededMaxTokenLengthException MessageErrorType = `EXCEEDED_MAX_TOKEN_LENGTH_EXCEPTION`

const MessageErrorTypeFunctionsNotAvailableException MessageErrorType = `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidJsonException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`

const MessageErrorTypeFunctionArgumentsInvalidTypeException MessageErrorType = `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`

const MessageErrorTypeFunctionCallMissingParameterException MessageErrorType = `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`

const MessageErrorTypeGeneratedSqlQueryTooLongException MessageErrorType = `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`

const MessageErrorTypeGenericChatCompletionException MessageErrorType = `GENERIC_CHAT_COMPLETION_EXCEPTION`

const MessageErrorTypeGenericChatCompletionServiceException MessageErrorType = `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`

const MessageErrorTypeGenericSqlExecApiCallException MessageErrorType = `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`

const MessageErrorTypeIllegalParameterDefinitionException MessageErrorType = `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`

const MessageErrorTypeInternalCatalogAssetCreationFailedException MessageErrorType = `INTERNAL_CATALOG_ASSET_CREATION_FAILED_EXCEPTION`

const MessageErrorTypeInternalCatalogAssetCreationOngoingException MessageErrorType = `INTERNAL_CATALOG_ASSET_CREATION_ONGOING_EXCEPTION`

const MessageErrorTypeInternalCatalogAssetCreationUnsupportedException MessageErrorType = `INTERNAL_CATALOG_ASSET_CREATION_UNSUPPORTED_EXCEPTION`

const MessageErrorTypeInternalCatalogMissingUcPathException MessageErrorType = `INTERNAL_CATALOG_MISSING_UC_PATH_EXCEPTION`

const MessageErrorTypeInternalCatalogPathOverlapException MessageErrorType = `INTERNAL_CATALOG_PATH_OVERLAP_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerFunctionException MessageErrorType = `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`

const MessageErrorTypeInvalidCertifiedAnswerIdentifierException MessageErrorType = `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`

const MessageErrorTypeInvalidChatCompletionJsonException MessageErrorType = `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`

const MessageErrorTypeInvalidCompletionRequestException MessageErrorType = `INVALID_COMPLETION_REQUEST_EXCEPTION`

const MessageErrorTypeInvalidFunctionCallException MessageErrorType = `INVALID_FUNCTION_CALL_EXCEPTION`

const MessageErrorTypeInvalidSqlMultipleDatasetReferencesException MessageErrorType = `INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION`

const MessageErrorTypeInvalidSqlMultipleStatementsException MessageErrorType = `INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION`

const MessageErrorTypeInvalidSqlUnknownTableException MessageErrorType = `INVALID_SQL_UNKNOWN_TABLE_EXCEPTION`

const MessageErrorTypeInvalidTableIdentifierException MessageErrorType = `INVALID_TABLE_IDENTIFIER_EXCEPTION`

const MessageErrorTypeLocalContextExceededException MessageErrorType = `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`

const MessageErrorTypeMessageAttachmentTooLongError MessageErrorType = `MESSAGE_ATTACHMENT_TOO_LONG_ERROR`

const MessageErrorTypeMessageCancelledWhileExecutingException MessageErrorType = `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageDeletedWhileExecutingException MessageErrorType = `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMessageUpdatedWhileExecutingException MessageErrorType = `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`

const MessageErrorTypeMissingSqlQueryException MessageErrorType = `MISSING_SQL_QUERY_EXCEPTION`

const MessageErrorTypeNoDeploymentsAvailableToWorkspace MessageErrorType = `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`

const MessageErrorTypeNoQueryToVisualizeException MessageErrorType = `NO_QUERY_TO_VISUALIZE_EXCEPTION`

const MessageErrorTypeNoTablesToQueryException MessageErrorType = `NO_TABLES_TO_QUERY_EXCEPTION`

const MessageErrorTypeRateLimitExceededGenericException MessageErrorType = `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`

const MessageErrorTypeRateLimitExceededSpecifiedWaitException MessageErrorType = `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`

const MessageErrorTypeReplyProcessTimeoutException MessageErrorType = `REPLY_PROCESS_TIMEOUT_EXCEPTION`

const MessageErrorTypeRetryableProcessingException MessageErrorType = `RETRYABLE_PROCESSING_EXCEPTION`

const MessageErrorTypeSqlExecutionException MessageErrorType = `SQL_EXECUTION_EXCEPTION`

const MessageErrorTypeStopProcessDueToAutoRegenerate MessageErrorType = `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`

const MessageErrorTypeTablesMissingException MessageErrorType = `TABLES_MISSING_EXCEPTION`

const MessageErrorTypeTooManyCertifiedAnswersException MessageErrorType = `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`

const MessageErrorTypeTooManyTablesException MessageErrorType = `TOO_MANY_TABLES_EXCEPTION`

const MessageErrorTypeUnexpectedReplyProcessException MessageErrorType = `UNEXPECTED_REPLY_PROCESS_EXCEPTION`

const MessageErrorTypeUnknownAiModel MessageErrorType = `UNKNOWN_AI_MODEL`

const MessageErrorTypeUnsupportedConversationTypeException MessageErrorType = `UNSUPPORTED_CONVERSATION_TYPE_EXCEPTION`

const MessageErrorTypeWarehouseAccessMissingException MessageErrorType = `WAREHOUSE_ACCESS_MISSING_EXCEPTION`

const MessageErrorTypeWarehouseNotFoundException MessageErrorType = `WAREHOUSE_NOT_FOUND_EXCEPTION`

// String representation for [fmt.Print]
func (f *MessageErrorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MessageErrorType) Set(v string) error {
	switch v {
	case `BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION`, `CHAT_COMPLETION_CLIENT_EXCEPTION`, `CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION`, `CHAT_COMPLETION_NETWORK_EXCEPTION`, `CONTENT_FILTER_EXCEPTION`, `CONTEXT_EXCEEDED_EXCEPTION`, `COULD_NOT_GET_DASHBOARD_SCHEMA_EXCEPTION`, `COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION`, `COULD_NOT_GET_UC_SCHEMA_EXCEPTION`, `DEPLOYMENT_NOT_FOUND_EXCEPTION`, `DESCRIBE_QUERY_INVALID_SQL_ERROR`, `DESCRIBE_QUERY_TIMEOUT`, `DESCRIBE_QUERY_UNEXPECTED_FAILURE`, `EXCEEDED_MAX_TOKEN_LENGTH_EXCEPTION`, `FUNCTIONS_NOT_AVAILABLE_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION`, `FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION`, `FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION`, `GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION`, `GENERIC_CHAT_COMPLETION_EXCEPTION`, `GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION`, `GENERIC_SQL_EXEC_API_CALL_EXCEPTION`, `ILLEGAL_PARAMETER_DEFINITION_EXCEPTION`, `INTERNAL_CATALOG_ASSET_CREATION_FAILED_EXCEPTION`, `INTERNAL_CATALOG_ASSET_CREATION_ONGOING_EXCEPTION`, `INTERNAL_CATALOG_ASSET_CREATION_UNSUPPORTED_EXCEPTION`, `INTERNAL_CATALOG_MISSING_UC_PATH_EXCEPTION`, `INTERNAL_CATALOG_PATH_OVERLAP_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION`, `INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION`, `INVALID_CHAT_COMPLETION_JSON_EXCEPTION`, `INVALID_COMPLETION_REQUEST_EXCEPTION`, `INVALID_FUNCTION_CALL_EXCEPTION`, `INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION`, `INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION`, `INVALID_SQL_UNKNOWN_TABLE_EXCEPTION`, `INVALID_TABLE_IDENTIFIER_EXCEPTION`, `LOCAL_CONTEXT_EXCEEDED_EXCEPTION`, `MESSAGE_ATTACHMENT_TOO_LONG_ERROR`, `MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION`, `MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION`, `MISSING_SQL_QUERY_EXCEPTION`, `NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE`, `NO_QUERY_TO_VISUALIZE_EXCEPTION`, `NO_TABLES_TO_QUERY_EXCEPTION`, `RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION`, `RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION`, `REPLY_PROCESS_TIMEOUT_EXCEPTION`, `RETRYABLE_PROCESSING_EXCEPTION`, `SQL_EXECUTION_EXCEPTION`, `STOP_PROCESS_DUE_TO_AUTO_REGENERATE`, `TABLES_MISSING_EXCEPTION`, `TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION`, `TOO_MANY_TABLES_EXCEPTION`, `UNEXPECTED_REPLY_PROCESS_EXCEPTION`, `UNKNOWN_AI_MODEL`, `UNSUPPORTED_CONVERSATION_TYPE_EXCEPTION`, `WAREHOUSE_ACCESS_MISSING_EXCEPTION`, `WAREHOUSE_NOT_FOUND_EXCEPTION`:
		*f = MessageErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK_MULTIPLE_EXECUTIONS_EXCEPTION", "CHAT_COMPLETION_CLIENT_EXCEPTION", "CHAT_COMPLETION_CLIENT_TIMEOUT_EXCEPTION", "CHAT_COMPLETION_NETWORK_EXCEPTION", "CONTENT_FILTER_EXCEPTION", "CONTEXT_EXCEEDED_EXCEPTION", "COULD_NOT_GET_DASHBOARD_SCHEMA_EXCEPTION", "COULD_NOT_GET_MODEL_DEPLOYMENTS_EXCEPTION", "COULD_NOT_GET_UC_SCHEMA_EXCEPTION", "DEPLOYMENT_NOT_FOUND_EXCEPTION", "DESCRIBE_QUERY_INVALID_SQL_ERROR", "DESCRIBE_QUERY_TIMEOUT", "DESCRIBE_QUERY_UNEXPECTED_FAILURE", "EXCEEDED_MAX_TOKEN_LENGTH_EXCEPTION", "FUNCTIONS_NOT_AVAILABLE_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_JSON_EXCEPTION", "FUNCTION_ARGUMENTS_INVALID_TYPE_EXCEPTION", "FUNCTION_CALL_MISSING_PARAMETER_EXCEPTION", "GENERATED_SQL_QUERY_TOO_LONG_EXCEPTION", "GENERIC_CHAT_COMPLETION_EXCEPTION", "GENERIC_CHAT_COMPLETION_SERVICE_EXCEPTION", "GENERIC_SQL_EXEC_API_CALL_EXCEPTION", "ILLEGAL_PARAMETER_DEFINITION_EXCEPTION", "INTERNAL_CATALOG_ASSET_CREATION_FAILED_EXCEPTION", "INTERNAL_CATALOG_ASSET_CREATION_ONGOING_EXCEPTION", "INTERNAL_CATALOG_ASSET_CREATION_UNSUPPORTED_EXCEPTION", "INTERNAL_CATALOG_MISSING_UC_PATH_EXCEPTION", "INTERNAL_CATALOG_PATH_OVERLAP_EXCEPTION", "INVALID_CERTIFIED_ANSWER_FUNCTION_EXCEPTION", "INVALID_CERTIFIED_ANSWER_IDENTIFIER_EXCEPTION", "INVALID_CHAT_COMPLETION_JSON_EXCEPTION", "INVALID_COMPLETION_REQUEST_EXCEPTION", "INVALID_FUNCTION_CALL_EXCEPTION", "INVALID_SQL_MULTIPLE_DATASET_REFERENCES_EXCEPTION", "INVALID_SQL_MULTIPLE_STATEMENTS_EXCEPTION", "INVALID_SQL_UNKNOWN_TABLE_EXCEPTION", "INVALID_TABLE_IDENTIFIER_EXCEPTION", "LOCAL_CONTEXT_EXCEEDED_EXCEPTION", "MESSAGE_ATTACHMENT_TOO_LONG_ERROR", "MESSAGE_CANCELLED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_DELETED_WHILE_EXECUTING_EXCEPTION", "MESSAGE_UPDATED_WHILE_EXECUTING_EXCEPTION", "MISSING_SQL_QUERY_EXCEPTION", "NO_DEPLOYMENTS_AVAILABLE_TO_WORKSPACE", "NO_QUERY_TO_VISUALIZE_EXCEPTION", "NO_TABLES_TO_QUERY_EXCEPTION", "RATE_LIMIT_EXCEEDED_GENERIC_EXCEPTION", "RATE_LIMIT_EXCEEDED_SPECIFIED_WAIT_EXCEPTION", "REPLY_PROCESS_TIMEOUT_EXCEPTION", "RETRYABLE_PROCESSING_EXCEPTION", "SQL_EXECUTION_EXCEPTION", "STOP_PROCESS_DUE_TO_AUTO_REGENERATE", "TABLES_MISSING_EXCEPTION", "TOO_MANY_CERTIFIED_ANSWERS_EXCEPTION", "TOO_MANY_TABLES_EXCEPTION", "UNEXPECTED_REPLY_PROCESS_EXCEPTION", "UNKNOWN_AI_MODEL", "UNSUPPORTED_CONVERSATION_TYPE_EXCEPTION", "WAREHOUSE_ACCESS_MISSING_EXCEPTION", "WAREHOUSE_NOT_FOUND_EXCEPTION"`, v)
	}
}

// Values returns all possible values for MessageErrorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *MessageErrorType) Values() []MessageErrorType {
	return []MessageErrorType{
		MessageErrorTypeBlockMultipleExecutionsException,
		MessageErrorTypeChatCompletionClientException,
		MessageErrorTypeChatCompletionClientTimeoutException,
		MessageErrorTypeChatCompletionNetworkException,
		MessageErrorTypeContentFilterException,
		MessageErrorTypeContextExceededException,
		MessageErrorTypeCouldNotGetDashboardSchemaException,
		MessageErrorTypeCouldNotGetModelDeploymentsException,
		MessageErrorTypeCouldNotGetUcSchemaException,
		MessageErrorTypeDeploymentNotFoundException,
		MessageErrorTypeDescribeQueryInvalidSqlError,
		MessageErrorTypeDescribeQueryTimeout,
		MessageErrorTypeDescribeQueryUnexpectedFailure,
		MessageErrorTypeExceededMaxTokenLengthException,
		MessageErrorTypeFunctionsNotAvailableException,
		MessageErrorTypeFunctionArgumentsInvalidException,
		MessageErrorTypeFunctionArgumentsInvalidJsonException,
		MessageErrorTypeFunctionArgumentsInvalidTypeException,
		MessageErrorTypeFunctionCallMissingParameterException,
		MessageErrorTypeGeneratedSqlQueryTooLongException,
		MessageErrorTypeGenericChatCompletionException,
		MessageErrorTypeGenericChatCompletionServiceException,
		MessageErrorTypeGenericSqlExecApiCallException,
		MessageErrorTypeIllegalParameterDefinitionException,
		MessageErrorTypeInternalCatalogAssetCreationFailedException,
		MessageErrorTypeInternalCatalogAssetCreationOngoingException,
		MessageErrorTypeInternalCatalogAssetCreationUnsupportedException,
		MessageErrorTypeInternalCatalogMissingUcPathException,
		MessageErrorTypeInternalCatalogPathOverlapException,
		MessageErrorTypeInvalidCertifiedAnswerFunctionException,
		MessageErrorTypeInvalidCertifiedAnswerIdentifierException,
		MessageErrorTypeInvalidChatCompletionJsonException,
		MessageErrorTypeInvalidCompletionRequestException,
		MessageErrorTypeInvalidFunctionCallException,
		MessageErrorTypeInvalidSqlMultipleDatasetReferencesException,
		MessageErrorTypeInvalidSqlMultipleStatementsException,
		MessageErrorTypeInvalidSqlUnknownTableException,
		MessageErrorTypeInvalidTableIdentifierException,
		MessageErrorTypeLocalContextExceededException,
		MessageErrorTypeMessageAttachmentTooLongError,
		MessageErrorTypeMessageCancelledWhileExecutingException,
		MessageErrorTypeMessageDeletedWhileExecutingException,
		MessageErrorTypeMessageUpdatedWhileExecutingException,
		MessageErrorTypeMissingSqlQueryException,
		MessageErrorTypeNoDeploymentsAvailableToWorkspace,
		MessageErrorTypeNoQueryToVisualizeException,
		MessageErrorTypeNoTablesToQueryException,
		MessageErrorTypeRateLimitExceededGenericException,
		MessageErrorTypeRateLimitExceededSpecifiedWaitException,
		MessageErrorTypeReplyProcessTimeoutException,
		MessageErrorTypeRetryableProcessingException,
		MessageErrorTypeSqlExecutionException,
		MessageErrorTypeStopProcessDueToAutoRegenerate,
		MessageErrorTypeTablesMissingException,
		MessageErrorTypeTooManyCertifiedAnswersException,
		MessageErrorTypeTooManyTablesException,
		MessageErrorTypeUnexpectedReplyProcessException,
		MessageErrorTypeUnknownAiModel,
		MessageErrorTypeUnsupportedConversationTypeException,
		MessageErrorTypeWarehouseAccessMissingException,
		MessageErrorTypeWarehouseNotFoundException,
	}
}

// Type always returns MessageErrorType to satisfy [pflag.Value] interface
func (f *MessageErrorType) Type() string {
	return "MessageErrorType"
}

// MessageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart context
// step to determine relevant context. * `ASKING_AI`: Waiting for the LLM to
// respond to the user's question. * `PENDING_WAREHOUSE`: Waiting for warehouse
// before the SQL query can start executing. * `EXECUTING_QUERY`: Executing a
// generated SQL query. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API. * `FAILED`: The response generation or query execution failed. See
// `error` field. * `COMPLETED`: Message processing is completed. Results are in
// the `attachments` field. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API. * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL
// result is not available anymore. The user needs to rerun the query. Rerun the
// SQL query result by calling
// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
// API. * `CANCELLED`: Message has been cancelled.
type MessageStatus string

// Waiting for the LLM to respond to the user's question.
const MessageStatusAskingAi MessageStatus = `ASKING_AI`

// Message has been cancelled.
const MessageStatusCancelled MessageStatus = `CANCELLED`

// Message processing is completed. Results are in the `attachments` field. Get
// the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API.
const MessageStatusCompleted MessageStatus = `COMPLETED`

// Executing a generated SQL query. Get the SQL query result by calling
// [getMessageAttachmentQueryResult](:method:genie/getMessageAttachmentQueryResult)
// API.
const MessageStatusExecutingQuery MessageStatus = `EXECUTING_QUERY`

// The response generation or query execution failed. See `error` field.
const MessageStatusFailed MessageStatus = `FAILED`

// Fetching metadata from the data sources.
const MessageStatusFetchingMetadata MessageStatus = `FETCHING_METADATA`

// Running smart context step to determine relevant context.
const MessageStatusFilteringContext MessageStatus = `FILTERING_CONTEXT`

// Waiting for warehouse before the SQL query can start executing.
const MessageStatusPendingWarehouse MessageStatus = `PENDING_WAREHOUSE`

// SQL result is not available anymore. The user needs to rerun the query. Rerun
// the SQL query result by calling
// [executeMessageAttachmentQuery](:method:genie/executeMessageAttachmentQuery)
// API.
const MessageStatusQueryResultExpired MessageStatus = `QUERY_RESULT_EXPIRED`

// Message has been submitted.
const MessageStatusSubmitted MessageStatus = `SUBMITTED`

// String representation for [fmt.Print]
func (f *MessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MessageStatus) Set(v string) error {
	switch v {
	case `ASKING_AI`, `CANCELLED`, `COMPLETED`, `EXECUTING_QUERY`, `FAILED`, `FETCHING_METADATA`, `FILTERING_CONTEXT`, `PENDING_WAREHOUSE`, `QUERY_RESULT_EXPIRED`, `SUBMITTED`:
		*f = MessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASKING_AI", "CANCELLED", "COMPLETED", "EXECUTING_QUERY", "FAILED", "FETCHING_METADATA", "FILTERING_CONTEXT", "PENDING_WAREHOUSE", "QUERY_RESULT_EXPIRED", "SUBMITTED"`, v)
	}
}

// Values returns all possible values for MessageStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *MessageStatus) Values() []MessageStatus {
	return []MessageStatus{
		MessageStatusAskingAi,
		MessageStatusCancelled,
		MessageStatusCompleted,
		MessageStatusExecutingQuery,
		MessageStatusFailed,
		MessageStatusFetchingMetadata,
		MessageStatusFilteringContext,
		MessageStatusPendingWarehouse,
		MessageStatusQueryResultExpired,
		MessageStatusSubmitted,
	}
}

// Type always returns MessageStatus to satisfy [pflag.Value] interface
func (f *MessageStatus) Type() string {
	return "MessageStatus"
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	DisplayName string `json:"display_name,omitempty"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId string `json:"source_dashboard_id"`
	// Flag to indicate if mustache parameter syntax ({{ param }}) should be
	// auto-updated to named syntax (:param) when converting datasets in the
	// dashboard.
	UpdateParameterSyntax bool `json:"update_parameter_syntax,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigrateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigrateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId string `json:"-" url:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials bool `json:"embed_credentials,omitempty"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName string `json:"display_name,omitempty"`
	// Whether the publisher's credentials are embedded in the published
	// dashboard. When true, viewers run the dashboard's queries with the
	// publisher's data permissions (shared data permissions); when false, each
	// viewer runs them with their own data permissions.
	EmbedCredentials bool `json:"embed_credentials,omitempty"`
	// The timestamp of when the current published version was created, that is,
	// when the dashboard was last published.
	RevisionCreateTime string `json:"revision_create_time,omitempty"`
	// The ID of the SQL warehouse that runs the published dashboard's queries.
	// It is the warehouse given when the dashboard was published, or the
	// draft's warehouse if none was given.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishedDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryAttachmentParameter struct {
	// The name of the parameter as referenced in the SQL query.
	Keyword string `json:"keyword,omitempty"`
	// The SQL type of the parameter value.
	SqlType string `json:"sql_type,omitempty"`
	// The value bound to the parameter.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryAttachmentParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryAttachmentParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Result struct {
	// If result is truncated
	IsTruncated bool `json:"is_truncated,omitempty"`
	// Row count of the result
	RowCount int64 `json:"row_count,omitempty"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId string `json:"statement_id,omitempty"`
	// JWT corresponding to the statement contained in this result
	StatementIdSignature string `json:"statement_id_signature,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Result) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Result) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Request to revert a dashboard draft to its last published state.
type RevertDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
	// The etag of the draft dashboard, as returned by the get dashboard API.
	// When provided, the request fails with ABORTED if the draft has been
	// modified since that etag was read.
	Etag string `json:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RevertDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RevertDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response to revert a dashboard draft to its last published state.
type RevertDashboardResponse struct {
	// The reverted dashboard.
	Dashboard *Dashboard `json:"dashboard,omitempty"`
}

func (s *RevertDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A schedule periodically runs a dashboard on a cron cadence. Each run executes
// the dashboard's queries to refresh its cached data and sends a snapshot of
// the dashboard to the schedule's subscribers. See [Manage scheduled dashboard
// updates and subscriptions] for details.
//
// [Manage scheduled dashboard updates and subscriptions]: https://docs.databricks.com/dashboards/share/schedule-subscribe
type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime string `json:"create_time,omitempty"`
	// The cron expression and time zone that define when the schedule runs.
	CronSchedule CronSchedule `json:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The name of the schedule, as shown in the dashboard's list of schedules.
	DisplayName string `json:"display_name,omitempty"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag string `json:"etag,omitempty"`
	// Whether the schedule is paused. A paused schedule keeps its configuration
	// and subscribers, but the dashboard is not refreshed and no notifications
	// are sent until the schedule is unpaused.
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// UUID identifying the schedule.
	ScheduleId string `json:"schedule_id,omitempty"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime string `json:"update_time,omitempty"`
	// The ID of the SQL warehouse that runs the dashboard's queries during
	// scheduled refreshes. If omitted, the warehouse of the published dashboard
	// is used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Schedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Schedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SchedulePauseStatus string

const SchedulePauseStatusPaused SchedulePauseStatus = `PAUSED`

const SchedulePauseStatusUnpaused SchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *SchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = SchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Values returns all possible values for SchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SchedulePauseStatus) Values() []SchedulePauseStatus {
	return []SchedulePauseStatus{
		SchedulePauseStatusPaused,
		SchedulePauseStatusUnpaused,
	}
}

// Type always returns SchedulePauseStatus to satisfy [pflag.Value] interface
func (f *SchedulePauseStatus) Type() string {
	return "SchedulePauseStatus"
}

type ScoreReason string

const ScoreReasonColumnTypeDifference ScoreReason = `COLUMN_TYPE_DIFFERENCE`

const ScoreReasonEmptyGoodSql ScoreReason = `EMPTY_GOOD_SQL`

const ScoreReasonEmptyResult ScoreReason = `EMPTY_RESULT`

const ScoreReasonLlmJudgeFormattingError ScoreReason = `LLM_JUDGE_FORMATTING_ERROR`

const ScoreReasonLlmJudgeIncompleteOrPartialOutput ScoreReason = `LLM_JUDGE_INCOMPLETE_OR_PARTIAL_OUTPUT`

const ScoreReasonLlmJudgeIncorrectFunctionUsage ScoreReason = `LLM_JUDGE_INCORRECT_FUNCTION_USAGE`

const ScoreReasonLlmJudgeIncorrectMetricCalculation ScoreReason = `LLM_JUDGE_INCORRECT_METRIC_CALCULATION`

const ScoreReasonLlmJudgeIncorrectTableOrFieldUsage ScoreReason = `LLM_JUDGE_INCORRECT_TABLE_OR_FIELD_USAGE`

const ScoreReasonLlmJudgeInstructionComplianceOrMissingBusinessLogic ScoreReason = `LLM_JUDGE_INSTRUCTION_COMPLIANCE_OR_MISSING_BUSINESS_LOGIC`

const ScoreReasonLlmJudgeMisinterpretationOfUserRequest ScoreReason = `LLM_JUDGE_MISINTERPRETATION_OF_USER_REQUEST`

const ScoreReasonLlmJudgeMissingJoin ScoreReason = `LLM_JUDGE_MISSING_JOIN`

const ScoreReasonLlmJudgeMissingOrIncorrectAggregation ScoreReason = `LLM_JUDGE_MISSING_OR_INCORRECT_AGGREGATION`

const ScoreReasonLlmJudgeMissingOrIncorrectFilter ScoreReason = `LLM_JUDGE_MISSING_OR_INCORRECT_FILTER`

const ScoreReasonLlmJudgeMissingOrIncorrectJoin ScoreReason = `LLM_JUDGE_MISSING_OR_INCORRECT_JOIN`

const ScoreReasonLlmJudgeOther ScoreReason = `LLM_JUDGE_OTHER`

const ScoreReasonLlmJudgeSemanticError ScoreReason = `LLM_JUDGE_SEMANTIC_ERROR`

const ScoreReasonLlmJudgeSyntaxError ScoreReason = `LLM_JUDGE_SYNTAX_ERROR`

const ScoreReasonLlmJudgeWrongAggregation ScoreReason = `LLM_JUDGE_WRONG_AGGREGATION`

const ScoreReasonLlmJudgeWrongColumns ScoreReason = `LLM_JUDGE_WRONG_COLUMNS`

const ScoreReasonLlmJudgeWrongFilter ScoreReason = `LLM_JUDGE_WRONG_FILTER`

const ScoreReasonResultExtraColumns ScoreReason = `RESULT_EXTRA_COLUMNS`

const ScoreReasonResultExtraRows ScoreReason = `RESULT_EXTRA_ROWS`

const ScoreReasonResultMissingColumns ScoreReason = `RESULT_MISSING_COLUMNS`

const ScoreReasonResultMissingRows ScoreReason = `RESULT_MISSING_ROWS`

const ScoreReasonSingleCellDifference ScoreReason = `SINGLE_CELL_DIFFERENCE`

// String representation for [fmt.Print]
func (f *ScoreReason) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ScoreReason) Set(v string) error {
	switch v {
	case `COLUMN_TYPE_DIFFERENCE`, `EMPTY_GOOD_SQL`, `EMPTY_RESULT`, `LLM_JUDGE_FORMATTING_ERROR`, `LLM_JUDGE_INCOMPLETE_OR_PARTIAL_OUTPUT`, `LLM_JUDGE_INCORRECT_FUNCTION_USAGE`, `LLM_JUDGE_INCORRECT_METRIC_CALCULATION`, `LLM_JUDGE_INCORRECT_TABLE_OR_FIELD_USAGE`, `LLM_JUDGE_INSTRUCTION_COMPLIANCE_OR_MISSING_BUSINESS_LOGIC`, `LLM_JUDGE_MISINTERPRETATION_OF_USER_REQUEST`, `LLM_JUDGE_MISSING_JOIN`, `LLM_JUDGE_MISSING_OR_INCORRECT_AGGREGATION`, `LLM_JUDGE_MISSING_OR_INCORRECT_FILTER`, `LLM_JUDGE_MISSING_OR_INCORRECT_JOIN`, `LLM_JUDGE_OTHER`, `LLM_JUDGE_SEMANTIC_ERROR`, `LLM_JUDGE_SYNTAX_ERROR`, `LLM_JUDGE_WRONG_AGGREGATION`, `LLM_JUDGE_WRONG_COLUMNS`, `LLM_JUDGE_WRONG_FILTER`, `RESULT_EXTRA_COLUMNS`, `RESULT_EXTRA_ROWS`, `RESULT_MISSING_COLUMNS`, `RESULT_MISSING_ROWS`, `SINGLE_CELL_DIFFERENCE`:
		*f = ScoreReason(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN_TYPE_DIFFERENCE", "EMPTY_GOOD_SQL", "EMPTY_RESULT", "LLM_JUDGE_FORMATTING_ERROR", "LLM_JUDGE_INCOMPLETE_OR_PARTIAL_OUTPUT", "LLM_JUDGE_INCORRECT_FUNCTION_USAGE", "LLM_JUDGE_INCORRECT_METRIC_CALCULATION", "LLM_JUDGE_INCORRECT_TABLE_OR_FIELD_USAGE", "LLM_JUDGE_INSTRUCTION_COMPLIANCE_OR_MISSING_BUSINESS_LOGIC", "LLM_JUDGE_MISINTERPRETATION_OF_USER_REQUEST", "LLM_JUDGE_MISSING_JOIN", "LLM_JUDGE_MISSING_OR_INCORRECT_AGGREGATION", "LLM_JUDGE_MISSING_OR_INCORRECT_FILTER", "LLM_JUDGE_MISSING_OR_INCORRECT_JOIN", "LLM_JUDGE_OTHER", "LLM_JUDGE_SEMANTIC_ERROR", "LLM_JUDGE_SYNTAX_ERROR", "LLM_JUDGE_WRONG_AGGREGATION", "LLM_JUDGE_WRONG_COLUMNS", "LLM_JUDGE_WRONG_FILTER", "RESULT_EXTRA_COLUMNS", "RESULT_EXTRA_ROWS", "RESULT_MISSING_COLUMNS", "RESULT_MISSING_ROWS", "SINGLE_CELL_DIFFERENCE"`, v)
	}
}

// Values returns all possible values for ScoreReason.
//
// There is no guarantee on the order of the values in the slice.
func (f *ScoreReason) Values() []ScoreReason {
	return []ScoreReason{
		ScoreReasonColumnTypeDifference,
		ScoreReasonEmptyGoodSql,
		ScoreReasonEmptyResult,
		ScoreReasonLlmJudgeFormattingError,
		ScoreReasonLlmJudgeIncompleteOrPartialOutput,
		ScoreReasonLlmJudgeIncorrectFunctionUsage,
		ScoreReasonLlmJudgeIncorrectMetricCalculation,
		ScoreReasonLlmJudgeIncorrectTableOrFieldUsage,
		ScoreReasonLlmJudgeInstructionComplianceOrMissingBusinessLogic,
		ScoreReasonLlmJudgeMisinterpretationOfUserRequest,
		ScoreReasonLlmJudgeMissingJoin,
		ScoreReasonLlmJudgeMissingOrIncorrectAggregation,
		ScoreReasonLlmJudgeMissingOrIncorrectFilter,
		ScoreReasonLlmJudgeMissingOrIncorrectJoin,
		ScoreReasonLlmJudgeOther,
		ScoreReasonLlmJudgeSemanticError,
		ScoreReasonLlmJudgeSyntaxError,
		ScoreReasonLlmJudgeWrongAggregation,
		ScoreReasonLlmJudgeWrongColumns,
		ScoreReasonLlmJudgeWrongFilter,
		ScoreReasonResultExtraColumns,
		ScoreReasonResultExtraRows,
		ScoreReasonResultMissingColumns,
		ScoreReasonResultMissingRows,
		ScoreReasonSingleCellDifference,
	}
}

// Type always returns ScoreReason to satisfy [pflag.Value] interface
func (f *ScoreReason) Type() string {
	return "ScoreReason"
}

type Subscriber struct {
	// The notification destination that receives the subscription, such as an
	// email list or a Slack or Microsoft Teams channel. Workspace admins
	// configure notification destinations.
	DestinationSubscriber *SubscriptionSubscriberDestination `json:"destination_subscriber,omitempty"`
	// The workspace user who receives the subscription.
	UserSubscriber *SubscriptionSubscriberUser `json:"user_subscriber,omitempty"`
}

func (s *Subscriber) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A subscription registers a workspace user or a notification destination as a
// subscriber of a dashboard schedule. Each time the schedule runs, the
// subscriber receives a snapshot of the published dashboard, unless the
// subscription is refresh-only. See [Manage scheduled dashboard updates and
// subscriptions].
//
// [Manage scheduled dashboard updates and subscriptions]: https://docs.databricks.com/dashboards/share/schedule-subscribe
type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime string `json:"create_time,omitempty"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId int64 `json:"created_by_user_id,omitempty"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The etag of the subscription. It must be left empty on create. It can be
	// passed as the `etag` parameter of the delete method to make sure the
	// subscription has not changed since it was read.
	Etag string `json:"etag,omitempty"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId string `json:"schedule_id,omitempty"`
	// Whether the subscription is refresh-only. When true, each scheduled run
	// refreshes the dashboard data for the subscriber without sending a
	// notification. Defaults to false. Refresh-only subscriptions are only
	// allowed on dashboards published without embedded credentials.
	SkipNotify bool `json:"skip_notify,omitempty"`
	// The subscriber that receives the dashboard snapshot. Exactly one of
	// `user_subscriber` and `destination_subscriber` must be set.
	Subscriber Subscriber `json:"subscriber"`
	// UUID identifying the subscription.
	SubscriptionId string `json:"subscription_id,omitempty"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Subscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Subscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SubscriptionSubscriberDestination struct {
	// The UUID of the notification destination that receives the subscription,
	// as returned by the [Notification Destinations API].
	//
	// [Notification Destinations API]: https://docs.databricks.com/api/workspace/notificationdestinations
	DestinationId string `json:"destination_id"`
}

func (s *SubscriptionSubscriberDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SubscriptionSubscriberUser struct {
	// The ID of the workspace user who receives the subscription, as returned
	// by the [Workspace Users API].
	//
	// [Workspace Users API]: https://docs.databricks.com/api/workspace/users
	UserId int64 `json:"user_id"`
}

func (s *SubscriptionSubscriberUser) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// A text response on a conversation message: the answer, the final summary, or
// a clarifying follow-up question, along with optional phase and verification
// metadata.
type TextAttachment struct {
	// AI generated message
	Content string `json:"content,omitempty"`

	Id string `json:"id,omitempty"`
	// Purpose of this text attachment. A completed message may contain more
	// than one text attachment (for example a clarifying follow-up question
	// alongside the final answer); use this field to tell them apart.
	// `TEXT_ATTACHMENT_PURPOSE_ANSWER` marks the final answer/summary and
	// `FOLLOW_UP_QUESTION` marks a clarifying question. Agent-mode messages
	// return the answer as their only text attachment.
	Purpose TextAttachmentPurpose `json:"purpose,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TextAttachment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TextAttachment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Purpose/intent of a text attachment
type TextAttachmentPurpose string

const TextAttachmentPurposeFollowUpQuestion TextAttachmentPurpose = `FOLLOW_UP_QUESTION`

const TextAttachmentPurposeTextAttachmentPurposeAnswer TextAttachmentPurpose = `TEXT_ATTACHMENT_PURPOSE_ANSWER`

// String representation for [fmt.Print]
func (f *TextAttachmentPurpose) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TextAttachmentPurpose) Set(v string) error {
	switch v {
	case `FOLLOW_UP_QUESTION`, `TEXT_ATTACHMENT_PURPOSE_ANSWER`:
		*f = TextAttachmentPurpose(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FOLLOW_UP_QUESTION", "TEXT_ATTACHMENT_PURPOSE_ANSWER"`, v)
	}
}

// Values returns all possible values for TextAttachmentPurpose.
//
// There is no guarantee on the order of the values in the slice.
func (f *TextAttachmentPurpose) Values() []TextAttachmentPurpose {
	return []TextAttachmentPurpose{
		TextAttachmentPurposeFollowUpQuestion,
		TextAttachmentPurposeTextAttachmentPurposeAnswer,
	}
}

// Type always returns TextAttachmentPurpose to satisfy [pflag.Value] interface
func (f *TextAttachmentPurpose) Type() string {
	return "TextAttachmentPurpose"
}

// A single thought in the AI's reasoning process for a query.
type Thought struct {
	// The md formatted content for this thought.
	Content string `json:"content,omitempty"`
	// The category of this thought. The possible values are: *
	// `THOUGHT_TYPE_DESCRIPTION`: A high-level description of how the question
	// was interpreted. * `THOUGHT_TYPE_UNDERSTANDING`: How ambiguous parts of
	// the question were resolved. * `THOUGHT_TYPE_DATA_SOURCING`: Which tables
	// or datasets were identified as relevant. * `THOUGHT_TYPE_INSTRUCTIONS`:
	// Which author-defined instructions were referenced. *
	// `THOUGHT_TYPE_STEPS`: The logical steps taken to compute the answer.
	ThoughtType ThoughtType `json:"thought_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Thought) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Thought) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// ThoughtType. The possible values are: * `THOUGHT_TYPE_UNSPECIFIED`: Default
// value that should not be used. * `THOUGHT_TYPE_DESCRIPTION`: A high-level
// description of how the question was interpreted. *
// `THOUGHT_TYPE_UNDERSTANDING`: How ambiguous parts of the question were
// resolved. * `THOUGHT_TYPE_DATA_SOURCING`: Which tables or datasets were
// identified as relevant. * `THOUGHT_TYPE_INSTRUCTIONS`: Which author-defined
// instructions were referenced. * `THOUGHT_TYPE_STEPS`: The logical steps taken
// to compute the answer. The category of a Thought. Additional values may be
// added in the future.
type ThoughtType string

// Which tables or datasets were identified as relevant.
const ThoughtTypeThoughtTypeDataSourcing ThoughtType = `THOUGHT_TYPE_DATA_SOURCING`

// A high-level description of how the question was interpreted.
const ThoughtTypeThoughtTypeDescription ThoughtType = `THOUGHT_TYPE_DESCRIPTION`

// Which author-defined instructions were referenced.
const ThoughtTypeThoughtTypeInstructions ThoughtType = `THOUGHT_TYPE_INSTRUCTIONS`

// The logical steps taken to compute the answer.
const ThoughtTypeThoughtTypeSteps ThoughtType = `THOUGHT_TYPE_STEPS`

// How ambiguous parts of the question were resolved.
const ThoughtTypeThoughtTypeUnderstanding ThoughtType = `THOUGHT_TYPE_UNDERSTANDING`

// String representation for [fmt.Print]
func (f *ThoughtType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ThoughtType) Set(v string) error {
	switch v {
	case `THOUGHT_TYPE_DATA_SOURCING`, `THOUGHT_TYPE_DESCRIPTION`, `THOUGHT_TYPE_INSTRUCTIONS`, `THOUGHT_TYPE_STEPS`, `THOUGHT_TYPE_UNDERSTANDING`:
		*f = ThoughtType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "THOUGHT_TYPE_DATA_SOURCING", "THOUGHT_TYPE_DESCRIPTION", "THOUGHT_TYPE_INSTRUCTIONS", "THOUGHT_TYPE_STEPS", "THOUGHT_TYPE_UNDERSTANDING"`, v)
	}
}

// Values returns all possible values for ThoughtType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ThoughtType) Values() []ThoughtType {
	return []ThoughtType{
		ThoughtTypeThoughtTypeDataSourcing,
		ThoughtTypeThoughtTypeDescription,
		ThoughtTypeThoughtTypeInstructions,
		ThoughtTypeThoughtTypeSteps,
		ThoughtTypeThoughtTypeUnderstanding,
	}
}

// Type always returns ThoughtType to satisfy [pflag.Value] interface
func (f *ThoughtType) Type() string {
	return "ThoughtType"
}

type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
}

func (s *TrashDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId string `json:"-" url:"-"`
}

func (s *UnpublishDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDashboardRequest struct {
	Dashboard Dashboard `json:"dashboard"`
	// UUID identifying the dashboard.
	DashboardId string `json:"-" url:"-"`
	// Sets the default catalog for all datasets in this dashboard. Does not
	// impact table references that use fully qualified catalog names (ex:
	// samples.nyctaxi.trips). Leave blank to keep each dataset’s existing
	// configuration.
	DatasetCatalog string `json:"-" url:"dataset_catalog,omitempty"`
	// Sets the default schema for all datasets in this dashboard. Does not
	// impact table references that use fully qualified schema names (ex:
	// nyctaxi.trips). Leave blank to keep each dataset’s existing
	// configuration.
	DatasetSchema string `json:"-" url:"dataset_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId string `json:"-" url:"-"`
	// The schedule to update.
	Schedule Schedule `json:"schedule"`
	// UUID identifying the schedule.
	ScheduleId string `json:"-" url:"-"`
}

func (s *UpdateScheduleRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}
