// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dashboards

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Genie API methods
type genieImpl struct {
	client *client.DatabricksClient
}

func (a *genieImpl) CreateMessage(ctx context.Context, request GenieCreateConversationMessageRequest) (*GenieMessage, error) {
	var genieMessage GenieMessage
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieMessage)
	return &genieMessage, err
}

func (a *genieImpl) CreateMessageComment(ctx context.Context, request GenieCreateMessageCommentRequest) (*GenieMessageComment, error) {
	var genieMessageComment GenieMessageComment
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/comments", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieMessageComment)
	return &genieMessageComment, err
}

func (a *genieImpl) CreateSpace(ctx context.Context, request GenieCreateSpaceRequest) (*GenieSpace, error) {
	var genieSpace GenieSpace
	path := "/api/2.0/genie/spaces"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieSpace)
	return &genieSpace, err
}

func (a *genieImpl) DeleteConversation(ctx context.Context, request GenieDeleteConversationRequest) error {
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *genieImpl) DeleteConversationMessage(ctx context.Context, request GenieDeleteConversationMessageRequest) error {
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *genieImpl) DownloadMessageAttachmentVisualization(ctx context.Context, request DownloadMessageAttachmentVisualizationRequest) (*DownloadMessageAttachmentVisualizationResponse, error) {
	var downloadMessageAttachmentVisualizationResponse DownloadMessageAttachmentVisualizationResponse
	path := fmt.Sprintf("/api/2.0/genie/%v/download-visualization", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/octet-stream"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &downloadMessageAttachmentVisualizationResponse)
	return &downloadMessageAttachmentVisualizationResponse, err
}

func (a *genieImpl) ExecuteMessageAttachmentQuery(ctx context.Context, request GenieExecuteMessageAttachmentQueryRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/execute-query", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *genieImpl) ExecuteMessageQuery(ctx context.Context, request GenieExecuteMessageQueryRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/execute-query", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *genieImpl) GenerateDownloadFullQueryResult(ctx context.Context, request GenieGenerateDownloadFullQueryResultRequest) (*GenieGenerateDownloadFullQueryResultResponse, error) {
	var genieGenerateDownloadFullQueryResultResponse GenieGenerateDownloadFullQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/downloads", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieGenerateDownloadFullQueryResultResponse)
	return &genieGenerateDownloadFullQueryResultResponse, err
}

func (a *genieImpl) GenieCancelResponse(ctx context.Context, request GenieCancelResponseRequest) (*GenieMessage, error) {
	var genieMessage GenieMessage
	path := fmt.Sprintf("/api/2.0/genie/agents/%v/conversations/%v/responses/%v/cancel", request.AgentId, request.ConversationId, request.ResponseId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieMessage)
	return &genieMessage, err
}

func (a *genieImpl) GenieCreateEvalRun(ctx context.Context, request GenieCreateEvalRunRequest) (*GenieEvalRunResponse, error) {
	var genieEvalRunResponse GenieEvalRunResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/eval-runs", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieEvalRunResponse)
	return &genieEvalRunResponse, err
}

func (a *genieImpl) GenieGetEvalResultDetails(ctx context.Context, request GenieGetEvalResultDetailsRequest) (*GenieEvalResultDetails, error) {
	var genieEvalResultDetails GenieEvalResultDetails
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/eval-runs/%v/results/%v", request.SpaceId, request.EvalRunId, request.ResultId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieEvalResultDetails)
	return &genieEvalResultDetails, err
}

func (a *genieImpl) GenieGetEvalRun(ctx context.Context, request GenieGetEvalRunRequest) (*GenieEvalRunResponse, error) {
	var genieEvalRunResponse GenieEvalRunResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/eval-runs/%v", request.SpaceId, request.EvalRunId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieEvalRunResponse)
	return &genieEvalRunResponse, err
}

func (a *genieImpl) GenieListEvalResults(ctx context.Context, request GenieListEvalResultsRequest) (*GenieListEvalResultsResponse, error) {
	var genieListEvalResultsResponse GenieListEvalResultsResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/eval-runs/%v/results", request.SpaceId, request.EvalRunId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListEvalResultsResponse)
	return &genieListEvalResultsResponse, err
}

func (a *genieImpl) GenieListEvalRuns(ctx context.Context, request GenieListEvalRunsRequest) (*GenieListEvalRunsResponse, error) {
	var genieListEvalRunsResponse GenieListEvalRunsResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/eval-runs", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListEvalRunsResponse)
	return &genieListEvalRunsResponse, err
}

func (a *genieImpl) GetDownloadFullQueryResult(ctx context.Context, request GenieGetDownloadFullQueryResultRequest) (*GenieGetDownloadFullQueryResultResponse, error) {
	var genieGetDownloadFullQueryResultResponse GenieGetDownloadFullQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/downloads/%v", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId, request.DownloadId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieGetDownloadFullQueryResultResponse)
	return &genieGetDownloadFullQueryResultResponse, err
}

func (a *genieImpl) GetMessage(ctx context.Context, request GenieGetConversationMessageRequest) (*GenieMessage, error) {
	var genieMessage GenieMessage
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieMessage)
	return &genieMessage, err
}

func (a *genieImpl) GetMessageAttachmentQueryResult(ctx context.Context, request GenieGetMessageAttachmentQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/attachments/%v/query-result", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *genieImpl) GetMessageQueryResult(ctx context.Context, request GenieGetMessageQueryResultRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/query-result", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *genieImpl) GetMessageQueryResultByAttachment(ctx context.Context, request GenieGetQueryResultByAttachmentRequest) (*GenieGetMessageQueryResultResponse, error) {
	var genieGetMessageQueryResultResponse GenieGetMessageQueryResultResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/query-result/%v", request.SpaceId, request.ConversationId, request.MessageId, request.AttachmentId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieGetMessageQueryResultResponse)
	return &genieGetMessageQueryResultResponse, err
}

func (a *genieImpl) GetSpace(ctx context.Context, request GenieGetSpaceRequest) (*GenieSpace, error) {
	var genieSpace GenieSpace
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieSpace)
	return &genieSpace, err
}

func (a *genieImpl) ListConversationComments(ctx context.Context, request GenieListConversationCommentsRequest) (*GenieListConversationCommentsResponse, error) {
	var genieListConversationCommentsResponse GenieListConversationCommentsResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/list-comments", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListConversationCommentsResponse)
	return &genieListConversationCommentsResponse, err
}

func (a *genieImpl) ListConversationMessages(ctx context.Context, request GenieListConversationMessagesRequest) (*GenieListConversationMessagesResponse, error) {
	var genieListConversationMessagesResponse GenieListConversationMessagesResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages", request.SpaceId, request.ConversationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListConversationMessagesResponse)
	return &genieListConversationMessagesResponse, err
}

func (a *genieImpl) ListConversations(ctx context.Context, request GenieListConversationsRequest) (*GenieListConversationsResponse, error) {
	var genieListConversationsResponse GenieListConversationsResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListConversationsResponse)
	return &genieListConversationsResponse, err
}

func (a *genieImpl) ListMessageComments(ctx context.Context, request GenieListMessageCommentsRequest) (*GenieListMessageCommentsResponse, error) {
	var genieListMessageCommentsResponse GenieListMessageCommentsResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/comments", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListMessageCommentsResponse)
	return &genieListMessageCommentsResponse, err
}

func (a *genieImpl) ListSpaces(ctx context.Context, request GenieListSpacesRequest) (*GenieListSpacesResponse, error) {
	var genieListSpacesResponse GenieListSpacesResponse
	path := "/api/2.0/genie/spaces"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &genieListSpacesResponse)
	return &genieListSpacesResponse, err
}

func (a *genieImpl) SendMessageFeedback(ctx context.Context, request GenieSendMessageFeedbackRequest) error {
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/conversations/%v/messages/%v/feedback", request.SpaceId, request.ConversationId, request.MessageId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, nil)
	return err
}

func (a *genieImpl) StartConversation(ctx context.Context, request GenieStartConversationMessageRequest) (*GenieStartConversationResponse, error) {
	var genieStartConversationResponse GenieStartConversationResponse
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v/start-conversation", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &genieStartConversationResponse)
	return &genieStartConversationResponse, err
}

func (a *genieImpl) TrashSpace(ctx context.Context, request GenieTrashSpaceRequest) error {
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *genieImpl) UpdateSpace(ctx context.Context, request GenieUpdateSpaceRequest) (*GenieSpace, error) {
	var genieSpace GenieSpace
	path := fmt.Sprintf("/api/2.0/genie/spaces/%v", request.SpaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &genieSpace)
	return &genieSpace, err
}

// unexported type that holds implementations of just Lakeview API methods
type lakeviewImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewImpl) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/lakeview/dashboards"
	queryParams := make(map[string]any)

	if request.DatasetCatalog != "" || slices.Contains(request.ForceSendFields, "DatasetCatalog") {
		queryParams["dataset_catalog"] = request.DatasetCatalog
	}

	if request.DatasetSchema != "" || slices.Contains(request.ForceSendFields, "DatasetSchema") {
		queryParams["dataset_schema"] = request.DatasetSchema
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Dashboard, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) CreateSchedule(ctx context.Context, request CreateScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Schedule, &schedule)
	return &schedule, err
}

func (a *lakeviewImpl) CreateSubscription(ctx context.Context, request CreateSubscriptionRequest) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Subscription, &subscription)
	return &subscription, err
}

func (a *lakeviewImpl) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakeviewImpl) DeleteSubscription(ctx context.Context, request DeleteSubscriptionRequest) error {
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakeviewImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error) {
	var publishedDashboard PublishedDashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &publishedDashboard)
	return &publishedDashboard, err
}

func (a *lakeviewImpl) GetSchedule(ctx context.Context, request GetScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &schedule)
	return &schedule, err
}

func (a *lakeviewImpl) GetSubscription(ctx context.Context, request GetSubscriptionRequest) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions/%v", request.DashboardId, request.ScheduleId, request.SubscriptionId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &subscription)
	return &subscription, err
}

// List the dashboards that the caller can view. Each entry carries summary
// metadata only; use Get dashboard to retrieve a dashboard's full definition.
func (a *lakeviewImpl) List(ctx context.Context, request ListDashboardsRequest) listing.Iterator[Dashboard] {

	getNextPage := func(ctx context.Context, req ListDashboardsRequest) (*ListDashboardsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListDashboardsResponse) []Dashboard {
		return resp.Dashboards
	}
	getNextReq := func(resp *ListDashboardsResponse) *ListDashboardsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List the dashboards that the caller can view. Each entry carries summary
// metadata only; use Get dashboard to retrieve a dashboard's full definition.
func (a *lakeviewImpl) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Dashboard](ctx, iterator)
}

func (a *lakeviewImpl) internalList(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error) {
	var listDashboardsResponse ListDashboardsResponse
	path := "/api/2.0/lakeview/dashboards"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDashboardsResponse)
	return &listDashboardsResponse, err
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
func (a *lakeviewImpl) ListSchedules(ctx context.Context, request ListSchedulesRequest) listing.Iterator[Schedule] {

	getNextPage := func(ctx context.Context, req ListSchedulesRequest) (*ListSchedulesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSchedules(ctx, req)
	}
	getItems := func(resp *ListSchedulesResponse) []Schedule {
		return resp.Schedules
	}
	getNextReq := func(resp *ListSchedulesResponse) *ListSchedulesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
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
func (a *lakeviewImpl) ListSchedulesAll(ctx context.Context, request ListSchedulesRequest) ([]Schedule, error) {
	iterator := a.ListSchedules(ctx, request)
	return listing.ToSlice[Schedule](ctx, iterator)
}

func (a *lakeviewImpl) internalListSchedules(ctx context.Context, request ListSchedulesRequest) (*ListSchedulesResponse, error) {
	var listSchedulesResponse ListSchedulesResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSchedulesResponse)
	return &listSchedulesResponse, err
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
func (a *lakeviewImpl) ListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) listing.Iterator[Subscription] {

	getNextPage := func(ctx context.Context, req ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListSubscriptions(ctx, req)
	}
	getItems := func(resp *ListSubscriptionsResponse) []Subscription {
		return resp.Subscriptions
	}
	getNextReq := func(resp *ListSubscriptionsResponse) *ListSubscriptionsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
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
func (a *lakeviewImpl) ListSubscriptionsAll(ctx context.Context, request ListSubscriptionsRequest) ([]Subscription, error) {
	iterator := a.ListSubscriptions(ctx, request)
	return listing.ToSlice[Subscription](ctx, iterator)
}

func (a *lakeviewImpl) internalListSubscriptions(ctx context.Context, request ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	var listSubscriptionsResponse ListSubscriptionsResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v/subscriptions", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listSubscriptionsResponse)
	return &listSubscriptionsResponse, err
}

func (a *lakeviewImpl) Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/lakeview/dashboards/migrate"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error) {
	var publishedDashboard PublishedDashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &publishedDashboard)
	return &publishedDashboard, err
}

func (a *lakeviewImpl) Revert(ctx context.Context, request RevertDashboardRequest) (*RevertDashboardResponse, error) {
	var revertDashboardResponse RevertDashboardResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/revert", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &revertDashboardResponse)
	return &revertDashboardResponse, err
}

func (a *lakeviewImpl) Trash(ctx context.Context, request TrashDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakeviewImpl) Unpublish(ctx context.Context, request UnpublishDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *lakeviewImpl) Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v", request.DashboardId)
	queryParams := make(map[string]any)

	if request.DatasetCatalog != "" || slices.Contains(request.ForceSendFields, "DatasetCatalog") {
		queryParams["dataset_catalog"] = request.DatasetCatalog
	}

	if request.DatasetSchema != "" || slices.Contains(request.ForceSendFields, "DatasetSchema") {
		queryParams["dataset_schema"] = request.DatasetSchema
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Dashboard, &dashboard)
	return &dashboard, err
}

func (a *lakeviewImpl) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*Schedule, error) {
	var schedule Schedule
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/schedules/%v", request.DashboardId, request.ScheduleId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request.Schedule, &schedule)
	return &schedule, err
}

// unexported type that holds implementations of just LakeviewEmbedded API methods
type lakeviewEmbeddedImpl struct {
	client *client.DatabricksClient
}

func (a *lakeviewEmbeddedImpl) GetPublishedDashboardTokenInfo(ctx context.Context, request GetPublishedDashboardTokenInfoRequest) (*GetPublishedDashboardTokenInfoResponse, error) {
	var getPublishedDashboardTokenInfoResponse GetPublishedDashboardTokenInfoResponse
	path := fmt.Sprintf("/api/2.0/lakeview/dashboards/%v/published/tokeninfo", request.DashboardId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getPublishedDashboardTokenInfoResponse)
	return &getPublishedDashboardTokenInfoResponse, err
}
