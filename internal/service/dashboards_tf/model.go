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
	// leading slash and no trailing slash.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

type CreateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `tfsdk:"cron_schedule" tf:""`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber `tfsdk:"subscriber" tf:""`
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

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// The workspace path of the dashboard asset, including the file name.
	Path types.String `tfsdk:"path" tf:"optional"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The timestamp of when the dashboard was last updated by the user.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

// Delete dashboard schedule
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

type DeleteScheduleResponse struct {
}

// Delete schedule subscription
type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-"`
}

type DeleteSubscriptionResponse struct {
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

// Genie AI Response
type GenieAttachment struct {
	Query *QueryAttachment `tfsdk:"query" tf:"optional"`

	Text *TextAttachment `tfsdk:"text" tf:"optional"`
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

type GenieCreateConversationMessageRequest struct {
	// User message content.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the conversation.
	ConversationId types.String `tfsdk:"-"`
	// The ID associated with the Genie space where the conversation is started.
	SpaceId types.String `tfsdk:"-"`
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

// Get conversation message SQL query result
type GenieGetMessageQueryResultRequest struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

type GenieGetMessageQueryResultResponse struct {
	// SQL Statement Execution response. See [Get status, manifest, and result
	// first chunk](:method:statementexecution/getstatement) for more details.
	StatementResponse *sql.StatementResponse `tfsdk:"statement_response" tf:"optional"`
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
	Error *MessageError `tfsdk:"error" tf:"optional"`
	// Message ID
	Id types.String `tfsdk:"id" tf:""`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The result of SQL query if the message has a query attachment
	QueryResult *Result `tfsdk:"query_result" tf:"optional"`
	// Genie space ID
	SpaceId types.String `tfsdk:"space_id" tf:""`
	// MesssageStatus. The possible values are: * `FETCHING_METADATA`: Fetching
	// metadata from the data sources. * `ASKING_AI`: Waiting for the LLM to
	// respond to the users question. * `EXECUTING_QUERY`: Executing AI provided
	// SQL query. Get the SQL query result by calling
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

type GenieStartConversationMessageRequest struct {
	// The text of the message that starts the conversation.
	Content types.String `tfsdk:"content" tf:""`
	// The ID associated with the Genie space where you want to start a
	// conversation.
	SpaceId types.String `tfsdk:"-"`
}

type GenieStartConversationResponse struct {
	Conversation *GenieConversation `tfsdk:"conversation" tf:"optional"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`

	Message *GenieMessage `tfsdk:"message" tf:"optional"`
	// Message ID
	MessageId types.String `tfsdk:"message_id" tf:""`
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

// Get schedule subscription
type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-"`
}

// List dashboards
type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed types.Bool `tfsdk:"-"`
	// `DASHBOARD_VIEW_BASIC`only includes summary metadata from the dashboard.
	View types.String `tfsdk:"-"`
}

type ListDashboardsResponse struct {
	Dashboards []Dashboard `tfsdk:"dashboards" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of schedules to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Schedules []Schedule `tfsdk:"schedules" tf:"optional"`
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of subscriptions to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Subscriptions []Subscription `tfsdk:"subscriptions" tf:"optional"`
}

type MessageError struct {
	Error types.String `tfsdk:"error" tf:"optional"`

	Type types.String `tfsdk:"type" tf:"optional"`
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

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time" tf:"optional"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
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

type Result struct {
	// Row count of the result
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// Statement Execution API statement id. Use [Get status, manifest, and
	// result first chunk](:method:statementexecution/getstatement) to get the
	// full result data.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `tfsdk:"cron_schedule" tf:""`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"optional"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber *SubscriptionSubscriberDestination `tfsdk:"destination_subscriber" tf:"optional"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber *SubscriptionSubscriberUser `tfsdk:"user_subscriber" tf:"optional"`
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId types.Int64 `tfsdk:"created_by_user_id" tf:"optional"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"optional"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"optional"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber `tfsdk:"subscriber" tf:""`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"optional"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId types.String `tfsdk:"destination_id" tf:""`
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId types.Int64 `tfsdk:"user_id" tf:""`
}

type TextAttachment struct {
	// AI generated message
	Content types.String `tfsdk:"content" tf:"optional"`

	Id types.String `tfsdk:"id" tf:"optional"`
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

type TrashDashboardResponse struct {
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-"`
}

type UnpublishDashboardResponse struct {
}

type UpdateDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard types.String `tfsdk:"serialized_dashboard" tf:"optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

type UpdateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `tfsdk:"cron_schedule" tf:""`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag" tf:"optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}
