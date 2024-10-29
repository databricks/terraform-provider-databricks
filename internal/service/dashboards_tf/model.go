// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package dashboards_tf

import (
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CreateDashboardRequest struct {
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:""`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath          types.String `tfsdk:"parent_path" tf:"optional"`
	EffectiveParentPath types.String `tfsdk:"effective_parent_path" tf:"computed,optional"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *CreateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDashboardRequest) {
	newState.EffectiveParentPath = newState.ParentPath
	newState.ParentPath = plan.ParentPath
}

func (newState *CreateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState CreateDashboardRequest) {
	newState.EffectiveParentPath = existingState.EffectiveParentPath
	if existingState.EffectiveParentPath.ValueString() == newState.ParentPath.ValueString() {
		newState.ParentPath = existingState.ParentPath
	}
}

type CreateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule []CronSchedule `tfsdk:"cron_schedule" tf:"object"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
}

func (newState *CreateScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateScheduleRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
}

func (newState *CreateScheduleRequest) SyncEffectiveFieldsDuringRead(existingState CreateScheduleRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber []Subscriber `tfsdk:"subscriber" tf:"object"`
}

func (newState *CreateSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSubscriptionRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
}

func (newState *CreateSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState CreateSubscriptionRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
}

type CronSchedule struct {
	// A cron expression using quartz syntax. EX: `0 0 8 * * ?` represents
	// everyday at 8am. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression" tf:""`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId types.String `tfsdk:"timezone_id" tf:""`
}

func (newState *CronSchedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan CronSchedule) {
}

func (newState *CronSchedule) SyncEffectiveFieldsDuringRead(existingState CronSchedule) {
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime          types.String `tfsdk:"create_time" tf:"optional"`
	EffectiveCreateTime types.String `tfsdk:"effective_create_time" tf:"computed,optional"`
	// UUID identifying the dashboard.
	DashboardId          types.String `tfsdk:"dashboard_id" tf:"optional"`
	EffectiveDashboardId types.String `tfsdk:"effective_dashboard_id" tf:"computed,optional"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag          types.String `tfsdk:"etag" tf:"optional"`
	EffectiveEtag types.String `tfsdk:"effective_etag" tf:"computed,optional"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath          types.String `tfsdk:"parent_path" tf:"optional"`
	EffectiveParentPath types.String `tfsdk:"effective_parent_path" tf:"computed,optional"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path          types.String `tfsdk:"path" tf:"optional"`
	EffectivePath types.String `tfsdk:"effective_path" tf:"computed,optional"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The timestamp of when the dashboard was last updated by the user. This
	// field is excluded in List Dashboards responses.
	UpdateTime          types.String `tfsdk:"update_time" tf:"optional"`
	EffectiveUpdateTime types.String `tfsdk:"effective_update_time" tf:"computed,optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Dashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard) {
	newState.EffectiveCreateTime = newState.CreateTime
	newState.CreateTime = plan.CreateTime
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
	newState.EffectiveParentPath = newState.ParentPath
	newState.ParentPath = plan.ParentPath
	newState.EffectivePath = newState.Path
	newState.Path = plan.Path
	newState.EffectiveUpdateTime = newState.UpdateTime
	newState.UpdateTime = plan.UpdateTime
}

func (newState *Dashboard) SyncEffectiveFieldsDuringRead(existingState Dashboard) {
	newState.EffectiveCreateTime = existingState.EffectiveCreateTime
	if existingState.EffectiveCreateTime.ValueString() == newState.CreateTime.ValueString() {
		newState.CreateTime = existingState.CreateTime
	}
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
	newState.EffectiveParentPath = existingState.EffectiveParentPath
	if existingState.EffectiveParentPath.ValueString() == newState.ParentPath.ValueString() {
		newState.ParentPath = existingState.ParentPath
	}
	newState.EffectivePath = existingState.EffectivePath
	if existingState.EffectivePath.ValueString() == newState.Path.ValueString() {
		newState.Path = existingState.Path
	}
	newState.EffectiveUpdateTime = existingState.EffectiveUpdateTime
	if existingState.EffectiveUpdateTime.ValueString() == newState.UpdateTime.ValueString() {
		newState.UpdateTime = existingState.UpdateTime
	}
}

// Delete dashboard schedule
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag          types.String `tfsdk:"-"`
	EffectiveEtag types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
}

func (newState *DeleteScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
}

func (newState *DeleteScheduleRequest) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
}

type DeleteScheduleResponse struct {
}

func (newState *DeleteScheduleResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleResponse) {
}

func (newState *DeleteScheduleResponse) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleResponse) {
}

// Delete schedule subscription
type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag          types.String `tfsdk:"-"`
	EffectiveEtag types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId          types.String `tfsdk:"-"`
	EffectiveSubscriptionId types.String `tfsdk:"-"`
}

func (newState *DeleteSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
	newState.EffectiveSubscriptionId = newState.SubscriptionId
	newState.SubscriptionId = plan.SubscriptionId
}

func (newState *DeleteSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
	newState.EffectiveSubscriptionId = existingState.EffectiveSubscriptionId
	if existingState.EffectiveSubscriptionId.ValueString() == newState.SubscriptionId.ValueString() {
		newState.SubscriptionId = existingState.SubscriptionId
	}
}

type DeleteSubscriptionResponse struct {
}

func (newState *DeleteSubscriptionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionResponse) {
}

func (newState *DeleteSubscriptionResponse) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionResponse) {
}

// Execute SQL query in a conversation message
type ExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *ExecuteMessageQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExecuteMessageQueryRequest) {
}

func (newState *ExecuteMessageQueryRequest) SyncEffectiveFieldsDuringRead(existingState ExecuteMessageQueryRequest) {
}

// Genie AI Response
type GenieAttachment struct {
	Query []QueryAttachment `tfsdk:"query" tf:"optional,object"`

	Text []TextAttachment `tfsdk:"text" tf:"optional,object"`
}

func (newState *GenieAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieAttachment) {
}

func (newState *GenieAttachment) SyncEffectiveFieldsDuringRead(existingState GenieAttachment) {
}

type GenieConversation struct {
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp" tf:"optional"`
	// Conversation ID
	Id types.String `tfsdk:"id" tf:""`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id" tf:""`
	// Conversation title
	Title types.String `tfsdk:"title" tf:""`
	// ID of the user who created the conversation
	UserId types.Int64 `tfsdk:"user_id" tf:""`
}

func (newState *GenieConversation) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieConversation) {
}

func (newState *GenieConversation) SyncEffectiveFieldsDuringRead(existingState GenieConversation) {
}

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieCreateConversationMessageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieCreateConversationMessageRequest) {
}

func (newState *GenieCreateConversationMessageRequest) SyncEffectiveFieldsDuringRead(existingState GenieCreateConversationMessageRequest) {
}

// Get conversation message
type GenieGetConversationMessageRequest struct {
	// The ID associated with the target conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the target message from the identified
	// conversation.
	MessageId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the target conversation is
	// located.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieGetConversationMessageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetConversationMessageRequest) {
}

func (newState *GenieGetConversationMessageRequest) SyncEffectiveFieldsDuringRead(existingState GenieGetConversationMessageRequest) {
}

// Get conversation message SQL query result
type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieGetMessageQueryResultRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultRequest) {
}

func (newState *GenieGetMessageQueryResultRequest) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultRequest) {
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse sql.StatementResponse `tfsdk:"statement_response" tf:"optional,object"`
}

func (newState *GenieGetMessageQueryResultResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieGetMessageQueryResultResponse) {
}

func (newState *GenieGetMessageQueryResultResponse) SyncEffectiveFieldsDuringRead(existingState GenieGetMessageQueryResultResponse) {
}

type GenieMessage struct {
	// AI produced response to the message
	Attachments []GenieAttachment `tfsdk:"attachments" tf:"optional"`
	// User message content
	Content types.String `tfsdk:"content" tf:""`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp" tf:"optional"`
	// Error message if AI failed to respond to the message
	Error []MessageError `tfsdk:"error" tf:"optional,object"`
	// Message ID
	Id types.String `tfsdk:"id" tf:""`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The result of SQL query if the message has a query attachment
	QueryResult []Result `tfsdk:"query_result" tf:"optional,object"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id" tf:""`
	// MesssageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
	// metadata from the data sources. * `FILTERING_CONTEXT`: Running smart
	// context step to determine relevant context. * `ASKING_AI`: Waiting for
	// the LLM to respond to the users question. * `EXECUTING_QUERY`: Executing
	// AI provided SQL query. Get the SQL query result by calling
	// [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
	// **Important: The message status will stay in the `EXECUTING_QUERY` until
	// a client calls
	// [getMessageQueryResult](:method:genie/getMessageQueryResult)**. *
	// `FAILED`: Generating a response or the executing the query failed. Please
	// see `error` field. * `COMPLETED`: Message processing is completed.
	// Results are in the `attachments` field. Get the SQL query result by
	// calling [getMessageQueryResult](:method:genie/getMessageQueryResult) API.
	// * `SUBMITTED`: Message has been submitted. * `QUERY_RESULT_EXPIRED`: SQL
	// result is not available anymore. The user needs to execute the query
	// again. * `CANCELLED`: Message has been cancelled.
	Status types.String `tfsdk:"status" tf:"optional"`
	// ID of the user who created the message
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`
}

func (newState *GenieMessage) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieMessage) {
}

func (newState *GenieMessage) SyncEffectiveFieldsDuringRead(existingState GenieMessage) {
}

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieStartConversationMessageRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationMessageRequest) {
}

func (newState *GenieStartConversationMessageRequest) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationMessageRequest) {
}

type GenieStartConversationResponse struct {
	Conversation []GenieConversation `tfsdk:"conversation" tf:"optional,object"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`

	Message []GenieMessage `tfsdk:"message" tf:"optional,object"`
	// Message ID
	MessageId types.String `tfsdk:"message_id" tf:""`
}

func (newState *GenieStartConversationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationResponse) {
}

func (newState *GenieStartConversationResponse) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationResponse) {
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDashboardRequest) {
}

func (newState *GetDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetDashboardRequest) {
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetPublishedDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedDashboardRequest) {
}

func (newState *GetPublishedDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetPublishedDashboardRequest) {
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
}

func (newState *GetScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetScheduleRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
}

func (newState *GetScheduleRequest) SyncEffectiveFieldsDuringRead(existingState GetScheduleRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
}

// Get schedule subscription
type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId          types.String `tfsdk:"-"`
	EffectiveSubscriptionId types.String `tfsdk:"-"`
}

func (newState *GetSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSubscriptionRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
	newState.EffectiveSubscriptionId = newState.SubscriptionId
	newState.SubscriptionId = plan.SubscriptionId
}

func (newState *GetSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState GetSubscriptionRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
	newState.EffectiveSubscriptionId = existingState.EffectiveSubscriptionId
	if existingState.EffectiveSubscriptionId.ValueString() == newState.SubscriptionId.ValueString() {
		newState.SubscriptionId = existingState.SubscriptionId
	}
}

// List dashboards
type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken          types.String `tfsdk:"-"`
	EffectivePageToken types.String `tfsdk:"-"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed types.Bool `tfsdk:"-"`
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	View types.String `tfsdk:"-"`
}

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsRequest) {
	newState.EffectivePageToken = newState.PageToken
	newState.PageToken = plan.PageToken
}

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringRead(existingState ListDashboardsRequest) {
	newState.EffectivePageToken = existingState.EffectivePageToken
	if existingState.EffectivePageToken.ValueString() == newState.PageToken.ValueString() {
		newState.PageToken = existingState.PageToken
	}
}

type ListDashboardsResponse struct {
	Dashboards []Dashboard `tfsdk:"dashboards" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken          types.String `tfsdk:"next_page_token" tf:"optional"`
	EffectiveNextPageToken types.String `tfsdk:"effective_next_page_token" tf:"computed,optional"`
}

func (newState *ListDashboardsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsResponse) {
	newState.EffectiveNextPageToken = newState.NextPageToken
	newState.NextPageToken = plan.NextPageToken
}

func (newState *ListDashboardsResponse) SyncEffectiveFieldsDuringRead(existingState ListDashboardsResponse) {
	newState.EffectiveNextPageToken = existingState.EffectiveNextPageToken
	if existingState.EffectiveNextPageToken.ValueString() == newState.NextPageToken.ValueString() {
		newState.NextPageToken = existingState.NextPageToken
	}
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// The number of schedules to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken          types.String `tfsdk:"-"`
	EffectivePageToken types.String `tfsdk:"-"`
}

func (newState *ListSchedulesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectivePageToken = newState.PageToken
	newState.PageToken = plan.PageToken
}

func (newState *ListSchedulesRequest) SyncEffectiveFieldsDuringRead(existingState ListSchedulesRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectivePageToken = existingState.EffectivePageToken
	if existingState.EffectivePageToken.ValueString() == newState.PageToken.ValueString() {
		newState.PageToken = existingState.PageToken
	}
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken          types.String `tfsdk:"next_page_token" tf:"optional"`
	EffectiveNextPageToken types.String `tfsdk:"effective_next_page_token" tf:"computed,optional"`

	Schedules []Schedule `tfsdk:"schedules" tf:"optional"`
}

func (newState *ListSchedulesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesResponse) {
	newState.EffectiveNextPageToken = newState.NextPageToken
	newState.NextPageToken = plan.NextPageToken
}

func (newState *ListSchedulesResponse) SyncEffectiveFieldsDuringRead(existingState ListSchedulesResponse) {
	newState.EffectiveNextPageToken = existingState.EffectiveNextPageToken
	if existingState.EffectiveNextPageToken.ValueString() == newState.NextPageToken.ValueString() {
		newState.NextPageToken = existingState.NextPageToken
	}
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// The number of subscriptions to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken          types.String `tfsdk:"-"`
	EffectivePageToken types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
}

func (newState *ListSubscriptionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectivePageToken = newState.PageToken
	newState.PageToken = plan.PageToken
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
}

func (newState *ListSubscriptionsRequest) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectivePageToken = existingState.EffectivePageToken
	if existingState.EffectivePageToken.ValueString() == newState.PageToken.ValueString() {
		newState.PageToken = existingState.PageToken
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken          types.String `tfsdk:"next_page_token" tf:"optional"`
	EffectiveNextPageToken types.String `tfsdk:"effective_next_page_token" tf:"computed,optional"`

	Subscriptions []Subscription `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *ListSubscriptionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsResponse) {
	newState.EffectiveNextPageToken = newState.NextPageToken
	newState.NextPageToken = plan.NextPageToken
}

func (newState *ListSubscriptionsResponse) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsResponse) {
	newState.EffectiveNextPageToken = existingState.EffectiveNextPageToken
	if existingState.EffectiveNextPageToken.ValueString() == newState.NextPageToken.ValueString() {
		newState.NextPageToken = existingState.NextPageToken
	}
}

type MessageError struct {
	Error types.String `tfsdk:"error" tf:"optional"`

	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *MessageError) SyncEffectiveFieldsDuringCreateOrUpdate(plan MessageError) {
}

func (newState *MessageError) SyncEffectiveFieldsDuringRead(existingState MessageError) {
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId types.String `tfsdk:"source_dashboard_id" tf:""`
}

func (newState *MigrateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan MigrateDashboardRequest) {
}

func (newState *MigrateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState MigrateDashboardRequest) {
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *PublishRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishRequest) {
}

func (newState *PublishRequest) SyncEffectiveFieldsDuringRead(existingState PublishRequest) {
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName          types.String `tfsdk:"display_name" tf:"optional"`
	EffectiveDisplayName types.String `tfsdk:"effective_display_name" tf:"computed,optional"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime          types.String `tfsdk:"revision_create_time" tf:"optional"`
	EffectiveRevisionCreateTime types.String `tfsdk:"effective_revision_create_time" tf:"computed,optional"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *PublishedDashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedDashboard) {
	newState.EffectiveDisplayName = newState.DisplayName
	newState.DisplayName = plan.DisplayName
	newState.EffectiveRevisionCreateTime = newState.RevisionCreateTime
	newState.RevisionCreateTime = plan.RevisionCreateTime
}

func (newState *PublishedDashboard) SyncEffectiveFieldsDuringRead(existingState PublishedDashboard) {
	newState.EffectiveDisplayName = existingState.EffectiveDisplayName
	if existingState.EffectiveDisplayName.ValueString() == newState.DisplayName.ValueString() {
		newState.DisplayName = existingState.DisplayName
	}
	newState.EffectiveRevisionCreateTime = existingState.EffectiveRevisionCreateTime
	if existingState.EffectiveRevisionCreateTime.ValueString() == newState.RevisionCreateTime.ValueString() {
		newState.RevisionCreateTime = existingState.RevisionCreateTime
	}
}

type QueryAttachment struct {
	// Description of the query
	Description types.String `tfsdk:"description" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
	// If the query was created on an instruction (trusted asset) we link to the
	// id
	InstructionId types.String `tfsdk:"instruction_id" tf:"optional"`
	// Always store the title next to the id in case the original instruction
	// title changes or the instruction is deleted.
	InstructionTitle types.String `tfsdk:"instruction_title" tf:"optional"`
	// Time when the user updated the query last
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// AI generated SQL query
	Query types.String `tfsdk:"query" tf:"optional"`
	// Name of the query
	Title types.String `tfsdk:"title" tf:"optional"`
}

func (newState *QueryAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryAttachment) {
}

func (newState *QueryAttachment) SyncEffectiveFieldsDuringRead(existingState QueryAttachment) {
}

type Result struct {
	// If result is truncated
	IsTruncated types.Bool `tfsdk:"is_truncated" tf:"optional"`
	// Row count of the result
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
}

func (newState *Result) SyncEffectiveFieldsDuringCreateOrUpdate(plan Result) {
}

func (newState *Result) SyncEffectiveFieldsDuringRead(existingState Result) {
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime          types.String `tfsdk:"create_time" tf:"optional"`
	EffectiveCreateTime types.String `tfsdk:"effective_create_time" tf:"computed,optional"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule []CronSchedule `tfsdk:"cron_schedule" tf:"object"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId          types.String `tfsdk:"dashboard_id" tf:"optional"`
	EffectiveDashboardId types.String `tfsdk:"effective_dashboard_id" tf:"computed,optional"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag          types.String `tfsdk:"etag" tf:"optional"`
	EffectiveEtag types.String `tfsdk:"effective_etag" tf:"computed,optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId          types.String `tfsdk:"schedule_id" tf:"optional"`
	EffectiveScheduleId types.String `tfsdk:"effective_schedule_id" tf:"computed,optional"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime          types.String `tfsdk:"update_time" tf:"optional"`
	EffectiveUpdateTime types.String `tfsdk:"effective_update_time" tf:"computed,optional"`
}

func (newState *Schedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan Schedule) {
	newState.EffectiveCreateTime = newState.CreateTime
	newState.CreateTime = plan.CreateTime
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
	newState.EffectiveUpdateTime = newState.UpdateTime
	newState.UpdateTime = plan.UpdateTime
}

func (newState *Schedule) SyncEffectiveFieldsDuringRead(existingState Schedule) {
	newState.EffectiveCreateTime = existingState.EffectiveCreateTime
	if existingState.EffectiveCreateTime.ValueString() == newState.CreateTime.ValueString() {
		newState.CreateTime = existingState.CreateTime
	}
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
	newState.EffectiveUpdateTime = existingState.EffectiveUpdateTime
	if existingState.EffectiveUpdateTime.ValueString() == newState.UpdateTime.ValueString() {
		newState.UpdateTime = existingState.UpdateTime
	}
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber []SubscriptionSubscriberDestination `tfsdk:"destination_subscriber" tf:"optional,object"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber []SubscriptionSubscriberUser `tfsdk:"user_subscriber" tf:"optional,object"`
}

func (newState *Subscriber) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscriber) {
}

func (newState *Subscriber) SyncEffectiveFieldsDuringRead(existingState Subscriber) {
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime          types.String `tfsdk:"create_time" tf:"optional"`
	EffectiveCreateTime types.String `tfsdk:"effective_create_time" tf:"computed,optional"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId          types.Int64 `tfsdk:"created_by_user_id" tf:"optional"`
	EffectiveCreatedByUserId types.Int64 `tfsdk:"effective_created_by_user_id" tf:"computed,optional"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId          types.String `tfsdk:"dashboard_id" tf:"optional"`
	EffectiveDashboardId types.String `tfsdk:"effective_dashboard_id" tf:"computed,optional"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag          types.String `tfsdk:"etag" tf:"optional"`
	EffectiveEtag types.String `tfsdk:"effective_etag" tf:"computed,optional"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId          types.String `tfsdk:"schedule_id" tf:"optional"`
	EffectiveScheduleId types.String `tfsdk:"effective_schedule_id" tf:"computed,optional"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber []Subscriber `tfsdk:"subscriber" tf:"object"`
	// UUID identifying the subscription.
	SubscriptionId          types.String `tfsdk:"subscription_id" tf:"optional"`
	EffectiveSubscriptionId types.String `tfsdk:"effective_subscription_id" tf:"computed,optional"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime          types.String `tfsdk:"update_time" tf:"optional"`
	EffectiveUpdateTime types.String `tfsdk:"effective_update_time" tf:"computed,optional"`
}

func (newState *Subscription) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscription) {
	newState.EffectiveCreateTime = newState.CreateTime
	newState.CreateTime = plan.CreateTime
	newState.EffectiveCreatedByUserId = newState.CreatedByUserId
	newState.CreatedByUserId = plan.CreatedByUserId
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
	newState.EffectiveSubscriptionId = newState.SubscriptionId
	newState.SubscriptionId = plan.SubscriptionId
	newState.EffectiveUpdateTime = newState.UpdateTime
	newState.UpdateTime = plan.UpdateTime
}

func (newState *Subscription) SyncEffectiveFieldsDuringRead(existingState Subscription) {
	newState.EffectiveCreateTime = existingState.EffectiveCreateTime
	if existingState.EffectiveCreateTime.ValueString() == newState.CreateTime.ValueString() {
		newState.CreateTime = existingState.CreateTime
	}
	newState.EffectiveCreatedByUserId = existingState.EffectiveCreatedByUserId
	if existingState.EffectiveCreatedByUserId.ValueInt64() == newState.CreatedByUserId.ValueInt64() {
		newState.CreatedByUserId = existingState.CreatedByUserId
	}
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
	newState.EffectiveSubscriptionId = existingState.EffectiveSubscriptionId
	if existingState.EffectiveSubscriptionId.ValueString() == newState.SubscriptionId.ValueString() {
		newState.SubscriptionId = existingState.SubscriptionId
	}
	newState.EffectiveUpdateTime = existingState.EffectiveUpdateTime
	if existingState.EffectiveUpdateTime.ValueString() == newState.UpdateTime.ValueString() {
		newState.UpdateTime = existingState.UpdateTime
	}
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId          types.String `tfsdk:"destination_id" tf:""`
	EffectiveDestinationId types.String `tfsdk:"effective_destination_id" tf:"computed,optional"`
}

func (newState *SubscriptionSubscriberDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberDestination) {
	newState.EffectiveDestinationId = newState.DestinationId
	newState.DestinationId = plan.DestinationId
}

func (newState *SubscriptionSubscriberDestination) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberDestination) {
	newState.EffectiveDestinationId = existingState.EffectiveDestinationId
	if existingState.EffectiveDestinationId.ValueString() == newState.DestinationId.ValueString() {
		newState.DestinationId = existingState.DestinationId
	}
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId          types.Int64 `tfsdk:"user_id" tf:""`
	EffectiveUserId types.Int64 `tfsdk:"effective_user_id" tf:"computed,optional"`
}

func (newState *SubscriptionSubscriberUser) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberUser) {
	newState.EffectiveUserId = newState.UserId
	newState.UserId = plan.UserId
}

func (newState *SubscriptionSubscriberUser) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberUser) {
	newState.EffectiveUserId = existingState.EffectiveUserId
	if existingState.EffectiveUserId.ValueInt64() == newState.UserId.ValueInt64() {
		newState.UserId = existingState.UserId
	}
}

type TextAttachment struct {
	// AI generated message
	Content types.String `tfsdk:"content" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *TextAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextAttachment) {
}

func (newState *TextAttachment) SyncEffectiveFieldsDuringRead(existingState TextAttachment) {
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *TrashDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardRequest) {
}

func (newState *TrashDashboardRequest) SyncEffectiveFieldsDuringRead(existingState TrashDashboardRequest) {
}

type TrashDashboardResponse struct {
}

func (newState *TrashDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardResponse) {
}

func (newState *TrashDashboardResponse) SyncEffectiveFieldsDuringRead(existingState TrashDashboardResponse) {
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UnpublishDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardRequest) {
}

func (newState *UnpublishDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardRequest) {
}

type UnpublishDashboardResponse struct {
}

func (newState *UnpublishDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardResponse) {
}

func (newState *UnpublishDashboardResponse) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardResponse) {
}

type UpdateDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag          types.String `tfsdk:"etag" tf:"optional"`
	EffectiveEtag types.String `tfsdk:"effective_etag" tf:"computed,optional"`
	// The contents of the dashboard in serialized string form. This field is
	// excluded in List Dashboards responses. Use the [get dashboard API] to
	// retrieve an example response, which includes the `serialized_dashboard`
	// field. This field provides the structure of the JSON string that
	// represents the dashboard's layout and components.
	//
	// [get dashboard API]: https://docs.databricks.com/api/workspace/lakeview/get
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *UpdateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDashboardRequest) {
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
}

func (newState *UpdateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDashboardRequest) {
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
}

type UpdateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule []CronSchedule `tfsdk:"cron_schedule" tf:"object"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId          types.String `tfsdk:"-"`
	EffectiveDashboardId types.String `tfsdk:"-"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag          types.String `tfsdk:"etag" tf:"optional"`
	EffectiveEtag types.String `tfsdk:"effective_etag" tf:"computed,optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId          types.String `tfsdk:"-"`
	EffectiveScheduleId types.String `tfsdk:"-"`
}

func (newState *UpdateScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateScheduleRequest) {
	newState.EffectiveDashboardId = newState.DashboardId
	newState.DashboardId = plan.DashboardId
	newState.EffectiveEtag = newState.Etag
	newState.Etag = plan.Etag
	newState.EffectiveScheduleId = newState.ScheduleId
	newState.ScheduleId = plan.ScheduleId
}

func (newState *UpdateScheduleRequest) SyncEffectiveFieldsDuringRead(existingState UpdateScheduleRequest) {
	newState.EffectiveDashboardId = existingState.EffectiveDashboardId
	if existingState.EffectiveDashboardId.ValueString() == newState.DashboardId.ValueString() {
		newState.DashboardId = existingState.DashboardId
	}
	newState.EffectiveEtag = existingState.EffectiveEtag
	if existingState.EffectiveEtag.ValueString() == newState.Etag.ValueString() {
		newState.Etag = existingState.Etag
	}
	newState.EffectiveScheduleId = existingState.EffectiveScheduleId
	if existingState.EffectiveScheduleId.ValueString() == newState.ScheduleId.ValueString() {
		newState.ScheduleId = existingState.ScheduleId
	}
}
