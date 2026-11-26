// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Genie, Lakeview, Lakeview Embedded, etc.
package dashboards

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type GenieInterface interface {

	// WaitGetMessageGenieCompleted repeatedly calls [GenieAPI.GetMessage] and waits to reach COMPLETED state
	WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string,
		timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error)

	// Sends a new message in a chat-mode
	// [conversation](:method:genie/startconversation). The AI response uses all
	// previously created messages in the conversation to respond. The message is
	// processed asynchronously. Poll [Get a conversation
	// message](:method:genie/getmessage) until its `status` is `COMPLETED`,
	// `FAILED`, or `CANCELLED`.
	CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieMessage], error)

	// Calls [GenieAPIInterface.CreateMessage] and waits to reach COMPLETED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GenieMessage](60*time.Minute) functional option.
	//
	// Deprecated: use [GenieAPIInterface.CreateMessage].Get() or [GenieAPIInterface.WaitGetMessageGenieCompleted]
	CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error)

	// Create a comment on a conversation message.
	CreateMessageComment(ctx context.Context, request GenieCreateMessageCommentRequest) (*GenieMessageComment, error)

	// Creates a Genie space from a serialized payload.
	CreateSpace(ctx context.Context, request GenieCreateSpaceRequest) (*GenieSpace, error)

	// Deletes a conversation and all of its messages. Users can delete their own
	// conversations. Deleting another user's conversation requires at least CAN
	// MANAGE permission on the space.
	DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error

	// Deletes a conversation and all of its messages. Users can delete their own
	// conversations. Deleting another user's conversation requires at least CAN
	// MANAGE permission on the space.
	DeleteConversationBySpaceIdAndConversationId(ctx context.Context, spaceId string, conversationId string) error

	// Delete a conversation message.
	DeleteConversationMessage(ctx context.Context, request GenieDeleteConversationMessageRequest) error

	// Download a rendered image of a message visualization attachment. The response
	// body is the raw PNG image, not a JSON payload. This is only available if the
	// attachment is a visualization and the message status is `COMPLETED`.
	DownloadMessageAttachmentVisualization(ctx context.Context, request DownloadMessageAttachmentVisualizationRequest) (*DownloadMessageAttachmentVisualizationResponse, error)

	// Execute the SQL for a message query attachment. Use this API when the query
	// attachment has expired and needs to be re-executed.
	ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Execute Message Attachment
	// Query](:method:genie/executemessageattachmentquery) instead.
	ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error)

	// Initiates a new SQL execution and returns a `download_id` and
	// `download_id_signature` that you can use to track the progress of the
	// download. The query result is stored in an external link and can be retrieved
	// using the [Get Download Full Query
	// Result](:method:genie/getdownloadfullqueryresult) API. Both `download_id` and
	// `download_id_signature` must be provided when calling the Get endpoint.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs that
	// are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived cloud-storage
	// URL is generated to download the results. The URL contains temporary access
	// credentials, so protect it and do not set an `Authorization` header in the
	// download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for more
	// details.
	//
	// ----
	GenerateDownloadFullQueryResult(ctx context.Context, request GenieGenerateDownloadFullQueryResultRequest) (*GenieGenerateDownloadFullQueryResultResponse, error)

	// Cancels an in-flight agent-mode response. `response_id` is the id returned in
	// the `response.created` event from the agent-mode responses endpoint. The
	// response stops at the next agent boundary and its terminal state is returned.
	GenieCancelResponse(ctx context.Context, request GenieCancelResponseRequest) (*GenieMessage, error)

	// Starts a benchmark evaluation run that asks Genie each selected benchmark
	// question in a Genie space and compares every response with the benchmark SQL.
	// The run executes asynchronously: the response is the new run with its current
	// status. If it is still running, poll with the get method until the status is
	// terminal.
	GenieCreateEvalRun(ctx context.Context, request GenieCreateEvalRunRequest) (*GenieEvalRunResponse, error)

	// Gets the full result for one benchmark question in an evaluation run: its
	// assessment and the reasons for it, Genie's response with the SQL execution
	// result, and the expected response from the benchmark. The result must belong
	// to the given run.
	GenieGetEvalResultDetails(ctx context.Context, request GenieGetEvalResultDetailsRequest) (*GenieEvalResultDetails, error)

	// Gets the current status and summary counts of a benchmark evaluation run.
	// Poll this method until `eval_run_status` is terminal: DONE,
	// EVALUATION_CANCELLED or EVALUATION_TIMEOUT.
	GenieGetEvalRun(ctx context.Context, request GenieGetEvalRunRequest) (*GenieEvalRunResponse, error)

	// Lists the results of a benchmark evaluation run, one per benchmark question.
	// Each entry carries the result status and a snapshot of the question and
	// benchmark answer; the assessment and Genie's response are available through
	// the result details method.
	GenieListEvalResults(ctx context.Context, request GenieListEvalResultsRequest) (*GenieListEvalResultsResponse, error)

	// Lists the benchmark evaluation runs in a Genie space with the current status
	// and summary counts of each run.
	GenieListEvalRuns(ctx context.Context, request GenieListEvalRunsRequest) (*GenieListEvalRunsResponse, error)

	// After [Generating a Full Query Result
	// Download](:method:genie/generatedownloadfullqueryresult) and successfully
	// receiving a `download_id` and `download_id_signature`, use this API to poll
	// the download progress. Both `download_id` and `download_id_signature` are
	// required to call this endpoint. When the download is complete, the API
	// returns the result in the `EXTERNAL_LINKS` disposition, containing one or
	// more external links to the query result files.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs that
	// are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived cloud-storage
	// URL is generated to download the results. The URL contains temporary access
	// credentials, so protect it and do not set an `Authorization` header in the
	// download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for more
	// details.
	//
	// ----
	GetDownloadFullQueryResult(ctx context.Context, request GenieGetDownloadFullQueryResultRequest) (*GenieGetDownloadFullQueryResultResponse, error)

	// After [Generating a Full Query Result
	// Download](:method:genie/generatedownloadfullqueryresult) and successfully
	// receiving a `download_id` and `download_id_signature`, use this API to poll
	// the download progress. Both `download_id` and `download_id_signature` are
	// required to call this endpoint. When the download is complete, the API
	// returns the result in the `EXTERNAL_LINKS` disposition, containing one or
	// more external links to the query result files.
	//
	// ----
	//
	// ### **Warning: Databricks strongly recommends that you protect the URLs that
	// are returned by the `EXTERNAL_LINKS` disposition.**
	//
	// When you use the `EXTERNAL_LINKS` disposition, a short-lived cloud-storage
	// URL is generated to download the results. The URL contains temporary access
	// credentials, so protect it and do not set an `Authorization` header in the
	// download request.
	//
	// See [Execute Statement](:method:statementexecution/executestatement) for more
	// details.
	//
	// ----
	GetDownloadFullQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentIdAndDownloadId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string, downloadId string) (*GenieGetDownloadFullQueryResultResponse, error)

	// Gets a message from a chat-mode or agent-mode conversation. For a complete
	// agent-mode transcript, use the List conversation items endpoint.
	GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error)

	// Gets a message from a chat-mode or agent-mode conversation. For a complete
	// agent-mode transcript, use the List conversation items endpoint.
	GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error)

	// Gets the result of the SQL query in a message's query attachment. The result
	// is available only while the message status is `EXECUTING_QUERY` or
	// `COMPLETED`.
	GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// Gets the result of the SQL query in a message's query attachment. The result
	// is available only while the message status is `EXECUTING_QUERY` or
	// `COMPLETED`.
	GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error)

	// DEPRECATED: Use [Get Message Attachment Query
	// Result](:method:genie/getmessageattachmentqueryresult) instead.
	GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error)

	// Get details of a Genie Space.
	GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error)

	// Get details of a Genie Space.
	GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error)

	// List all comments across all messages in a conversation.
	ListConversationComments(ctx context.Context, request GenieListConversationCommentsRequest) (*GenieListConversationCommentsResponse, error)

	// Lists messages in a chat-mode or agent-mode conversation. Agent-mode messages
	// are returned as GenieMessage projections. Use the List conversation items
	// endpoint for the complete reasoning and tool-call history.
	ListConversationMessages(ctx context.Context, request GenieListConversationMessagesRequest) (*GenieListConversationMessagesResponse, error)

	// Lists the conversations in a Genie space. By default the list contains only
	// the caller's own conversations, ordered by most recent activity. With
	// `include_all`, it contains every user's conversations in the space, ordered
	// by creation time with the newest first.
	ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error)

	// Lists the conversations in a Genie space. By default the list contains only
	// the caller's own conversations, ordered by most recent activity. With
	// `include_all`, it contains every user's conversations in the space, ordered
	// by creation time with the newest first.
	ListConversationsBySpaceId(ctx context.Context, spaceId string) (*GenieListConversationsResponse, error)

	// List comments on a specific conversation message.
	ListMessageComments(ctx context.Context, request GenieListMessageCommentsRequest) (*GenieListMessageCommentsResponse, error)

	// Lists the Genie spaces that the caller can access, most recently updated
	// first.
	ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error)

	// Sends feedback for a message in a chat-mode or agent-mode conversation.
	SendMessageFeedback(ctx context.Context, request GenieSendMessageFeedbackRequest) error

	// Starts a new chat-mode conversation and sends its first message. The message
	// is processed asynchronously. Poll [Get a conversation
	// message](:method:genie/getmessage) until its `status` is `COMPLETED`,
	// `FAILED`, or `CANCELLED`.
	StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieStartConversationResponse], error)

	// Calls [GenieAPIInterface.StartConversation] and waits to reach COMPLETED state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[GenieMessage](60*time.Minute) functional option.
	//
	// Deprecated: use [GenieAPIInterface.StartConversation].Get() or [GenieAPIInterface.WaitGetMessageGenieCompleted]
	StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error)

	// Move a Genie Space to the trash.
	TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error

	// Move a Genie Space to the trash.
	TrashSpaceBySpaceId(ctx context.Context, spaceId string) error

	// Updates a Genie space. Replaces the whole configuration when
	// `serialized_space` is set and changes the title, description, warehouse or
	// folder when those fields are set; omitted fields keep their current values.
	UpdateSpace(ctx context.Context, request GenieUpdateSpaceRequest) (*GenieSpace, error)
}

func NewGenie(client *client.DatabricksClient) *GenieAPI {
	return &GenieAPI{
		genieImpl: genieImpl{
			client: client,
		},
	}
}

// Genie provides a no-code experience for business users, powered by AI/BI.
// Analysts set up spaces that business users can use to ask questions using
// natural language. Genie uses data registered to Unity Catalog and requires at
// least CAN USE permission on a Pro or Serverless SQL warehouse. Also,
// Databricks Assistant must be enabled.
type GenieAPI struct {
	genieImpl
}

// WaitGetMessageGenieCompleted repeatedly calls [GenieAPI.GetMessage] and waits to reach COMPLETED state
func (a *GenieAPI) WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string,
	timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[GenieMessage](ctx, timeout, func() (*GenieMessage, *retries.Err) {
		genieMessage, err := a.GetMessage(ctx, GenieGetConversationMessageRequest{
			ConversationId: conversationId,
			MessageId:      messageId,
			SpaceId:        spaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(genieMessage)
		}
		status := genieMessage.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case MessageStatusCompleted: // target state
			return genieMessage, nil
		case MessageStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				MessageStatusCompleted, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetMessageGenieCompleted is a wrapper that calls [GenieAPI.WaitGetMessageGenieCompleted] and waits to reach COMPLETED state.
type WaitGetMessageGenieCompleted[R any] struct {
	Response       *R
	ConversationId string `json:"conversation_id"`
	MessageId      string `json:"message_id"`
	SpaceId        string `json:"space_id"`
	Poll           func(time.Duration, func(*GenieMessage)) (*GenieMessage, error)
	callback       func(*GenieMessage)
	timeout        time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetMessageGenieCompleted[R]) OnProgress(callback func(*GenieMessage)) *WaitGetMessageGenieCompleted[R] {
	w.callback = callback
	return w
}

// Get the GenieMessage with the default timeout of 20 minutes.
func (w *WaitGetMessageGenieCompleted[R]) Get() (*GenieMessage, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the GenieMessage with custom timeout.
func (w *WaitGetMessageGenieCompleted[R]) GetWithTimeout(timeout time.Duration) (*GenieMessage, error) {
	return w.Poll(timeout, w.callback)
}

// Sends a new message in a chat-mode
// [conversation](:method:genie/startconversation). The AI response uses all
// previously created messages in the conversation to respond. The message is
// processed asynchronously. Poll [Get a conversation
// message](:method:genie/getmessage) until its `status` is `COMPLETED`,
// `FAILED`, or `CANCELLED`.
func (a *GenieAPI) CreateMessage(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieMessage], error) {
	genieMessage, err := a.genieImpl.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetMessageGenieCompleted[GenieMessage]{
		Response:       genieMessage,
		ConversationId: genieCreateConversationMessageRequest.ConversationId,
		MessageId:      genieMessage.MessageId,
		SpaceId:        genieCreateConversationMessageRequest.SpaceId,
		Poll: func(timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
			return a.WaitGetMessageGenieCompleted(ctx, genieCreateConversationMessageRequest.ConversationId, genieMessage.MessageId, genieCreateConversationMessageRequest.SpaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [GenieAPI.CreateMessage] and waits to reach COMPLETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GenieMessage](60*time.Minute) functional option.
//
// Deprecated: use [GenieAPI.CreateMessage].Get() or [GenieAPI.WaitGetMessageGenieCompleted]
func (a *GenieAPI) CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest GenieCreateConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error) {
	wait, err := a.CreateMessage(ctx, genieCreateConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GenieMessage]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GenieMessage) {
		for _, o := range options {
			o(&retries.Info[GenieMessage]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Deletes a conversation and all of its messages. Users can delete their own
// conversations. Deleting another user's conversation requires at least CAN
// MANAGE permission on the space.
func (a *GenieAPI) DeleteConversationBySpaceIdAndConversationId(ctx context.Context, spaceId string, conversationId string) error {
	return a.genieImpl.DeleteConversation(ctx, GenieDeleteConversationRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
	})
}

// After [Generating a Full Query Result
// Download](:method:genie/generatedownloadfullqueryresult) and successfully
// receiving a `download_id` and `download_id_signature`, use this API to poll
// the download progress. Both `download_id` and `download_id_signature` are
// required to call this endpoint. When the download is complete, the API
// returns the result in the `EXTERNAL_LINKS` disposition, containing one or
// more external links to the query result files.
//
// ----
//
// ### **Warning: Databricks strongly recommends that you protect the URLs that
// are returned by the `EXTERNAL_LINKS` disposition.**
//
// When you use the `EXTERNAL_LINKS` disposition, a short-lived cloud-storage
// URL is generated to download the results. The URL contains temporary access
// credentials, so protect it and do not set an `Authorization` header in the
// download request.
//
// See [Execute Statement](:method:statementexecution/executestatement) for more
// details.
//
// ----
func (a *GenieAPI) GetDownloadFullQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentIdAndDownloadId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string, downloadId string) (*GenieGetDownloadFullQueryResultResponse, error) {
	return a.genieImpl.GetDownloadFullQueryResult(ctx, GenieGetDownloadFullQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
		DownloadId:     downloadId,
	})
}

// Gets a message from a chat-mode or agent-mode conversation. For a complete
// agent-mode transcript, use the List conversation items endpoint.
func (a *GenieAPI) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieMessage, error) {
	return a.genieImpl.GetMessage(ctx, GenieGetConversationMessageRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// Gets the result of the SQL query in a message's query attachment. The result
// is available only while the message status is `EXECUTING_QUERY` or
// `COMPLETED`.
func (a *GenieAPI) GetMessageAttachmentQueryResultBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageAttachmentQueryResult(ctx, GenieGetMessageAttachmentQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

// DEPRECATED: Use [Get Message Attachment Query
// Result](:method:genie/getmessageattachmentqueryresult) instead.
func (a *GenieAPI) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageQueryResult(ctx, GenieGetMessageQueryResultRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
	})
}

// DEPRECATED: Use [Get Message Attachment Query
// Result](:method:genie/getmessageattachmentqueryresult) instead.
func (a *GenieAPI) GetMessageQueryResultByAttachmentBySpaceIdAndConversationIdAndMessageIdAndAttachmentId(ctx context.Context, spaceId string, conversationId string, messageId string, attachmentId string) (*GenieGetMessageQueryResultResponse, error) {
	return a.genieImpl.GetMessageQueryResultByAttachment(ctx, GenieGetQueryResultByAttachmentRequest{
		SpaceId:        spaceId,
		ConversationId: conversationId,
		MessageId:      messageId,
		AttachmentId:   attachmentId,
	})
}

// Get details of a Genie Space.
func (a *GenieAPI) GetSpaceBySpaceId(ctx context.Context, spaceId string) (*GenieSpace, error) {
	return a.genieImpl.GetSpace(ctx, GenieGetSpaceRequest{
		SpaceId: spaceId,
	})
}

// Lists the conversations in a Genie space. By default the list contains only
// the caller's own conversations, ordered by most recent activity. With
// `include_all`, it contains every user's conversations in the space, ordered
// by creation time with the newest first.
func (a *GenieAPI) ListConversationsBySpaceId(ctx context.Context, spaceId string) (*GenieListConversationsResponse, error) {
	return a.genieImpl.ListConversations(ctx, GenieListConversationsRequest{
		SpaceId: spaceId,
	})
}

// Starts a new chat-mode conversation and sends its first message. The message
// is processed asynchronously. Poll [Get a conversation
// message](:method:genie/getmessage) until its `status` is `COMPLETED`,
// `FAILED`, or `CANCELLED`.
func (a *GenieAPI) StartConversation(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest) (*WaitGetMessageGenieCompleted[GenieStartConversationResponse], error) {
	genieStartConversationResponse, err := a.genieImpl.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetMessageGenieCompleted[GenieStartConversationResponse]{
		Response:       genieStartConversationResponse,
		ConversationId: genieStartConversationResponse.ConversationId,
		MessageId:      genieStartConversationResponse.MessageId,
		SpaceId:        genieStartConversationMessageRequest.SpaceId,
		Poll: func(timeout time.Duration, callback func(*GenieMessage)) (*GenieMessage, error) {
			return a.WaitGetMessageGenieCompleted(ctx, genieStartConversationResponse.ConversationId, genieStartConversationResponse.MessageId, genieStartConversationMessageRequest.SpaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [GenieAPI.StartConversation] and waits to reach COMPLETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GenieMessage](60*time.Minute) functional option.
//
// Deprecated: use [GenieAPI.StartConversation].Get() or [GenieAPI.WaitGetMessageGenieCompleted]
func (a *GenieAPI) StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest GenieStartConversationMessageRequest, options ...retries.Option[GenieMessage]) (*GenieMessage, error) {
	wait, err := a.StartConversation(ctx, genieStartConversationMessageRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[GenieMessage]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *GenieMessage) {
		for _, o := range options {
			o(&retries.Info[GenieMessage]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Move a Genie Space to the trash.
func (a *GenieAPI) TrashSpaceBySpaceId(ctx context.Context, spaceId string) error {
	return a.genieImpl.TrashSpace(ctx, GenieTrashSpaceRequest{
		SpaceId: spaceId,
	})
}

type LakeviewInterface interface {

	// Create a draft dashboard.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Creates a schedule for a dashboard. The dashboard must be published before a
	// schedule can be created for it.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL access
	// in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error)

	// Create schedule subscription. Subscribes a workspace user or a notification
	// destination to the schedule, so that it receives a snapshot of the dashboard
	// each time the schedule runs. Subscribing yourself requires CAN VIEW
	// permission on the dashboard; subscribing another user or a notification
	// destination requires CAN EDIT. On dashboards published without embedded
	// credentials, you can only subscribe yourself. If the subscriber is already
	// subscribed to the schedule, the existing subscription is returned.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error)

	// Deletes a schedule together with all of its subscriptions.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL access
	// in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error

	// Deletes a schedule together with all of its subscriptions.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL access
	// in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error

	// Delete schedule subscription. The subscriber receives no further snapshots
	// from the schedule. Deleting your own subscription requires CAN VIEW
	// permission on the dashboard; deleting another user's or a notification
	// destination's subscription requires CAN EDIT.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error

	// Delete schedule subscription. The subscriber receives no further snapshots
	// from the schedule. Deleting your own subscription requires CAN VIEW
	// permission on the dashboard; deleting another user's or a notification
	// destination's subscription requires CAN EDIT.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error

	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get a draft dashboard.
	GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	// Get the current published dashboard. The request fails with NOT_FOUND if the
	// dashboard has never been published or has been unpublished.
	//
	// The caller must be a workspace user with one of the following entitlements:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get the current published dashboard. The request fails with NOT_FOUND if the
	// dashboard has never been published or has been unpublished.
	//
	// The caller must be a workspace user with one of the following entitlements:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error)

	// Get dashboard schedule.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error)

	// Get dashboard schedule.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error)

	// Get schedule subscription.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error)

	// Get schedule subscription.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error)

	// List the dashboards that the caller can view. Each entry carries summary
	// metadata only; use Get dashboard to retrieve a dashboard's full definition.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard]

	// List the dashboards that the caller can view. Each entry carries summary
	// metadata only; use Get dashboard to retrieve a dashboard's full definition.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error)

	// List dashboard schedules.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule]

	// List dashboard schedules.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error)

	// List dashboard schedules.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error)

	// List schedule subscriptions.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription]

	// List schedule subscriptions.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error)

	// List schedule subscriptions.
	//
	// The caller must be a workspace user with one of the following [entitlements]:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [entitlements]: https://docs.databricks.com/security/auth/entitlements
	ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error)

	// Deprecated: Legacy dashboard migration is no longer supported. Use Lakeview
	// (AI/BI) dashboards instead.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Revert a dashboard's definition in draft mode to the last published version.
	// The draft's datasets, pages and widgets are replaced with those of the
	// current published version, and the updated draft is returned. The dashboard
	// must have been published at least once; the published version itself is not
	// changed.
	Revert(ctx context.Context, request RevertDashboardRequest) (*RevertDashboardResponse, error)

	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Trash a dashboard.
	TrashByDashboardId(ctx context.Context, dashboardId string) error

	// Unpublish the dashboard. This deletes the published version, so viewers can
	// no longer open it, and pauses the dashboard's schedules. The draft is not
	// changed; publish it again to make the dashboard available to viewers.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Unpublish the dashboard. This deletes the published version, so viewers can
	// no longer open it, and pauses the dashboard's schedules. The draft is not
	// changed; publish it again to make the dashboard available to viewers.
	UnpublishByDashboardId(ctx context.Context, dashboardId string) error

	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)

	// Updates a schedule. The request body replaces the stored schedule, so fields
	// omitted from the body are reset to their defaults.
	//
	// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL access
	// in addition to Workspace access.
	//
	// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error)
}

func NewLakeview(client *client.DatabricksClient) *LakeviewAPI {
	return &LakeviewAPI{
		lakeviewImpl: lakeviewImpl{
			client: client,
		},
	}
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewAPI struct {
	lakeviewImpl
}

// Deletes a schedule together with all of its subscriptions.
//
// Requires the [Databricks SQL access] entitlement. Grant Databricks SQL access
// in addition to Workspace access.
//
// [Databricks SQL access]: https://docs.databricks.com/security/auth/entitlements
func (a *LakeviewAPI) DeleteScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) error {
	return a.lakeviewImpl.DeleteSchedule(ctx, DeleteScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Delete schedule subscription. The subscriber receives no further snapshots
// from the schedule. Deleting your own subscription requires CAN VIEW
// permission on the dashboard; deleting another user's or a notification
// destination's subscription requires CAN EDIT.
//
// The caller must be a workspace user with one of the following [entitlements]:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
//
// [entitlements]: https://docs.databricks.com/security/auth/entitlements
func (a *LakeviewAPI) DeleteSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) error {
	return a.lakeviewImpl.DeleteSubscription(ctx, DeleteSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// Get a draft dashboard.
func (a *LakeviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.lakeviewImpl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get the current published dashboard. The request fails with NOT_FOUND if the
// dashboard has never been published or has been unpublished.
//
// The caller must be a workspace user with one of the following entitlements:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
func (a *LakeviewAPI) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.lakeviewImpl.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard schedule.
//
// The caller must be a workspace user with one of the following [entitlements]:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
//
// [entitlements]: https://docs.databricks.com/security/auth/entitlements
func (a *LakeviewAPI) GetScheduleByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*Schedule, error) {
	return a.lakeviewImpl.GetSchedule(ctx, GetScheduleRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Get schedule subscription.
//
// The caller must be a workspace user with one of the following [entitlements]:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
//
// [entitlements]: https://docs.databricks.com/security/auth/entitlements
func (a *LakeviewAPI) GetSubscriptionByDashboardIdAndScheduleIdAndSubscriptionId(ctx context.Context, dashboardId string, scheduleId string, subscriptionId string) (*Subscription, error) {
	return a.lakeviewImpl.GetSubscription(ctx, GetSubscriptionRequest{
		DashboardId:    dashboardId,
		ScheduleId:     scheduleId,
		SubscriptionId: subscriptionId,
	})
}

// List dashboard schedules.
//
// The caller must be a workspace user with one of the following [entitlements]:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
//
// [entitlements]: https://docs.databricks.com/security/auth/entitlements
func (a *LakeviewAPI) ListSchedulesByDashboardId(ctx context.Context, dashboardId string) (*ListSchedulesResponse, error) {
	return a.lakeviewImpl.internalListSchedules(ctx, ListSchedulesRequest{
		DashboardId: dashboardId,
	})
}

// List schedule subscriptions.
//
// The caller must be a workspace user with one of the following [entitlements]:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
//
// [entitlements]: https://docs.databricks.com/security/auth/entitlements
func (a *LakeviewAPI) ListSubscriptionsByDashboardIdAndScheduleId(ctx context.Context, dashboardId string, scheduleId string) (*ListSubscriptionsResponse, error) {
	return a.lakeviewImpl.internalListSubscriptions(ctx, ListSubscriptionsRequest{
		DashboardId: dashboardId,
		ScheduleId:  scheduleId,
	})
}

// Trash a dashboard.
func (a *LakeviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewImpl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish the dashboard. This deletes the published version, so viewers can
// no longer open it, and pauses the dashboard's schedules. The draft is not
// changed; publish it again to make the dashboard available to viewers.
func (a *LakeviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) error {
	return a.lakeviewImpl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

type LakeviewEmbeddedInterface interface {

	// Get the OAuth scope and authorization details needed to mint a token that is
	// scoped to viewing the published dashboard, for example to embed it in an
	// external application. Pass the response fields, with `authorization_details`
	// serialized as a JSON string, to the workspace OAuth token endpoint with the
	// `client_credentials` grant to mint the token. See [Embedding for external
	// users].
	//
	// The caller must be a workspace user with one of the following entitlements:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [Embedding for external users]: https://docs.databricks.com/dashboards/share/embedding/external-embed
	GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error)

	// Get the OAuth scope and authorization details needed to mint a token that is
	// scoped to viewing the published dashboard, for example to embed it in an
	// external application. Pass the response fields, with `authorization_details`
	// serialized as a JSON string, to the workspace OAuth token endpoint with the
	// `client_credentials` grant to mint the token. See [Embedding for external
	// users].
	//
	// The caller must be a workspace user with one of the following entitlements:
	// Workspace access, Databricks SQL access, or Consumer access.
	//
	// Account-level users who are not members of the workspace cannot call this
	// endpoint, even if the dashboard has been shared with them.
	//
	// [Embedding for external users]: https://docs.databricks.com/dashboards/share/embedding/external-embed
	GetPublishedDashboardTokenInfoByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardTokenInfoResponse, error)
}

func NewLakeviewEmbedded(client *client.DatabricksClient) *LakeviewEmbeddedAPI {
	return &LakeviewEmbeddedAPI{
		lakeviewEmbeddedImpl: lakeviewEmbeddedImpl{
			client: client,
		},
	}
}

// Token-based Lakeview APIs for embedding dashboards in external applications.
type LakeviewEmbeddedAPI struct {
	lakeviewEmbeddedImpl
}

// Get the OAuth scope and authorization details needed to mint a token that is
// scoped to viewing the published dashboard, for example to embed it in an
// external application. Pass the response fields, with `authorization_details`
// serialized as a JSON string, to the workspace OAuth token endpoint with the
// `client_credentials` grant to mint the token. See [Embedding for external
// users].
//
// The caller must be a workspace user with one of the following entitlements:
// Workspace access, Databricks SQL access, or Consumer access.
//
// Account-level users who are not members of the workspace cannot call this
// endpoint, even if the dashboard has been shared with them.
//
// [Embedding for external users]: https://docs.databricks.com/dashboards/share/embedding/external-embed
func (a *LakeviewEmbeddedAPI) GetPublishedDashboardTokenInfoByDashboardId(ctx context.Context, dashboardId string) (*GetPublishedDashboardTokenInfoResponse, error) {
	return a.lakeviewEmbeddedImpl.GetPublishedDashboardTokenInfo(ctx, GetPublishedDashboardTokenInfoRequest{
		DashboardId: dashboardId,
	})
}
