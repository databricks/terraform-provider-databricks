// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
)

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GenieService interface {

	// Sends a new message in a chat-mode
	// [conversation](:method:genie/startconversation). The AI response uses all
	// previously created messages in the conversation to respond. The message
	// is processed asynchronously. Poll [Get a conversation
	// message](:method:genie/getmessage) until its `status` is `COMPLETED`,
	// `FAILED`, or `CANCELLED`.
	CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error)

	// Create a comment on a conversation message.
	CreateMessageComment(ctx context.Context, request GenieCreateMessageCommentRequest) (*GenieMessageComment, error)

	// Creates a Genie space from a serialized payload.
	CreateSpace(ctx context.Context, request GenieCreateSpaceRequest) (*GenieSpace, error)

	// Deletes a conversation and all of its messages. Users can delete their
	// own conversations. Deleting another user's conversation requires at least
	// CAN MANAGE permission on the space.
	DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error

	// Delete a conversation message.
	DeleteConversationMessage(ctx context.Context, request GenieDeleteConversationMessageRequest) error

	// Download a rendered image of a message visualization attachment. The
	// response body is the raw PNG image, not a JSON payload. This is only
	// available if the attachment is a visualization and the message status is
	// `COMPLETED`.
	DownloadMessageAttachmentVisualization(ctx context.Context, request DownloadMessageAttachmentVisualizationRequest) (*DownloadMessageAttachmentVisualizationResponse, error)

	// Execute the SQL for a message query attachment. Use this API when the
	// query attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Execute Message Attachment
	// Query](:method:genie/executemessageattachmentquery) instead.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Initiates a new SQL execution and returns a `download_id` and
	// `download_id_signature` that you can use to track the progress of the
	// download. The query result is stored in an external link and can be
	// retrieved using the [Get Download Full Query
	// Result](:method:genie/getdownloadfullqueryresult) API. Both `download_id`
	// and `download_id_signature` must be provided when calling the Get
	// endpoint.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs
	// that are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived
	// cloud-storage URL is generated to download the results. The URL contains
	// temporary access credentials, so protect it and do not set an
	// `Authorization` header in the download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for
	// more details.
	//
	// ----
	GenerateDownloadFullQueryResult(ctx context.Context, request GenieGenerateDownloadFullQueryResultRequest) (*GenieGenerateDownloadFullQueryResultResponse, error)

	// Cancels an in-flight agent-mode response. `response_id` is the id
	// returned in the `response.created` event from the agent-mode responses
	// endpoint. The response stops at the next agent boundary and its terminal
	// state is returned.
	GenieCancelResponse(ctx context.Context, request GenieCancelResponseRequest) (*GenieMessage, error)

	// Starts a benchmark evaluation run that asks Genie each selected benchmark
	// question in a Genie space and compares every response with the benchmark
	// SQL. The run executes asynchronously: the response is the new run with
	// its current status. If it is still running, poll with the get method
	// until the status is terminal.
	GenieCreateEvalRun(ctx context.Context, request GenieCreateEvalRunRequest) (*GenieEvalRunResponse, error)

	// Gets the full result for one benchmark question in an evaluation run: its
	// assessment and the reasons for it, Genie's response with the SQL
	// execution result, and the expected response from the benchmark. The
	// result must belong to the given run.
	GenieGetEvalResultDetails(ctx context.Context, request GenieGetEvalResultDetailsRequest) (*GenieEvalResultDetails, error)

	// Gets the current status and summary counts of a benchmark evaluation run.
	// Poll this method until `eval_run_status` is terminal: DONE,
	// EVALUATION_CANCELLED or EVALUATION_TIMEOUT.
	GenieGetEvalRun(ctx context.Context, request GenieGetEvalRunRequest) (*GenieEvalRunResponse, error)

	// Lists the results of a benchmark evaluation run, one per benchmark
	// question. Each entry carries the result status and a snapshot of the
	// question and benchmark answer; the assessment and Genie's response are
	// available through the result details method.
	GenieListEvalResults(ctx context.Context, request GenieListEvalResultsRequest) (*GenieListEvalResultsResponse, error)

	// Lists the benchmark evaluation runs in a Genie space with the current
	// status and summary counts of each run.
	GenieListEvalRuns(ctx context.Context, request GenieListEvalRunsRequest) (*GenieListEvalRunsResponse, error)

	// After [Generating a Full Query Result
	// Download](:method:genie/generatedownloadfullqueryresult) and successfully
	// receiving a `download_id` and `download_id_signature`, use this API to
	// poll the download progress. Both `download_id` and
	// `download_id_signature` are required to call this endpoint. When the
	// download is complete, the API returns the result in the `EXTERNAL_LINKS`
	// disposition, containing one or more external links to the query result
	// files.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs
	// that are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived
	// cloud-storage URL is generated to download the results. The URL contains
	// temporary access credentials, so protect it and do not set an
	// `Authorization` header in the download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for
	// more details.
	//
	// ----
	GetDownloadFullQueryResult(ctx context.Context, request GenieGetDownloadFullQueryResultRequest) (*GenieGetDownloadFullQueryResultResponse, error)

	// Gets a message from a chat-mode or agent-mode conversation. For a
	// complete agent-mode transcript, use the List conversation items endpoint.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Gets the result of the SQL query in a message's query attachment. The
	// result is available only while the message status is `EXECUTING_QUERY` or
	// `COMPLETED`.
	GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// Get details of a Genie Space.
	GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error)

	// List all comments across all messages in a conversation.
	ListConversationComments(ctx context.Context, request GenieListConversationCommentsRequest) (*GenieListConversationCommentsResponse, error)

	// Lists messages in a chat-mode or agent-mode conversation. Agent-mode
	// messages are returned as GenieMessage projections. Use the List
	// conversation items endpoint for the complete reasoning and tool-call
	// history.
	ListConversationMessages(ctx context.Context, request GenieListConversationMessagesRequest) (*GenieListConversationMessagesResponse, error)

	// Lists the conversations in a Genie space. By default the list contains
	// only the caller's own conversations, ordered by most recent activity.
	// With `include_all`, it contains every user's conversations in the space,
	// ordered by creation time with the newest first.
	ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error)

	// List comments on a specific conversation message.
	ListMessageComments(ctx context.Context, request GenieListMessageCommentsRequest) (*GenieListMessageCommentsResponse, error)

	// Lists the Genie spaces that the caller can access, most recently updated
	// first.
	ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error)

	// Sends feedback for a message in a chat-mode or agent-mode conversation.
	SendMessageFeedback(ctx context.Context, request GenieSendMessageFeedbackRequest) error

	// Starts a new chat-mode conversation and sends its first message. The
	// message is processed asynchronously. Poll [Get a conversation
	// message](:method:genie/getmessage) until its `status` is `COMPLETED`,
	// `FAILED`, or `CANCELLED`.
	StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error)

	// Move a Genie Space to the trash.
	TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error

	// Updates a Genie space. Replaces the whole configuration when
	// `serialized_space` is set and changes the title, description, warehouse
	// or folder when those fields are set; omitted fields keep their current
	// values.
	UpdateSpace(ctx context.Context, request GenieUpdateSpaceRequest) (*GenieSpace, error)
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakeviewService interface {

	// Create a draft dashboard.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Creates a schedule for a dashboard. The dashboard must be published
	// before a schedule can be created for it.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error)

	// Create schedule subscription. Subscribes a workspace user or a
	// notification destination to the schedule, so that it receives a snapshot
	// of the dashboard each time the schedule runs. Subscribing yourself
	// requires CAN VIEW permission on the dashboard; subscribing another user
	// or a notification destination requires CAN EDIT. On dashboards published
	// without embedded credentials, you can only subscribe yourself. If the
	// subscriber is already subscribed to the schedule, the existing
	// subscription is returned.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error)

	// Deletes a schedule together with all of its subscriptions.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error

	// Delete schedule subscription. The subscriber receives no further
	// snapshots from the schedule. Deleting your own subscription requires CAN
	// VIEW permission on the dashboard; deleting another user's or a
	// notification destination's subscription requires CAN EDIT.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error

	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get the current published dashboard. The request fails with NOT_FOUND if
	// the dashboard has never been published or has been unpublished.
	//
	// The caller must be a workspace user with one of the following
	// entitlements: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get dashboard schedule.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get schedule subscription.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// List the dashboards that the caller can view. Each entry carries summary
	// metadata only; use Get dashboard to retrieve a dashboard's full
	// definition.
	List(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error)

	// List dashboard schedules.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	ListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	//
	// The caller must be a workspace user with one of the following
	// [entitlements]: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)

	// Deprecated: Legacy dashboard migration is no longer supported. Use
	// Lakeview (AI/BI) dashboards instead.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Revert a dashboard's definition in draft mode to the last published
	// version. The draft's datasets, pages and widgets are replaced with those
	// of the current published version, and the updated draft is returned. The
	// dashboard must have been published at least once; the published version
	// itself is not changed.
	Revert(ctx context.Context, request RevertDashboardRequest) (*RevertDashboardResponse, error)

	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Unpublish the dashboard. This deletes the published version, so viewers
	// can no longer open it, and pauses the dashboard's schedules. The draft is
	// not changed; publish it again to make the dashboard available to viewers.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Updates a schedule. The request body replaces the stored schedule, so
	// fields omitted from the body are reset to their defaults.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL
	// access in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LakeviewEmbeddedService interface {

	// Get the OAuth scope and authorization details needed to mint a token that
	// is scoped to viewing the published dashboard, for example to embed it in
	// an external application. Pass the response fields, with
	// `authorization_details` serialized as a JSON string, to the workspace
	// OAuth token endpoint with the `client_credentials` grant to mint the
	// token. See [Embedding for external users].
	//
	// The caller must be a workspace user with one of the following
	// entitlements: Workspace access, Databricks SQL access, or Consumer
	// access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [Embedding for external users]: https://docs.databricks.com/dashboards/share/embedding/external-embed
	GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error)
}
