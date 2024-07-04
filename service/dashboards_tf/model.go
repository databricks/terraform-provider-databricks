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
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CreateDashboardRequest struct {
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath types.String `tfsdk:"parent_path"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard types.String `tfsdk:"serialized_dashboard"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type CreateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `tfsdk:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus `tfsdk:"pause_status"`
}

type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-" url:"-"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber `tfsdk:"subscriber"`
}

type CronSchedule struct {
	// A cron expression using quartz syntax. EX: `0 0 8 * * ?` represents
	// everyday at 8am. See [Cron Trigger] for details.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression types.String `tfsdk:"quartz_cron_expression"`
	// A Java timezone id. The schedule will be resolved with respect to this
	// timezone. See [Java TimeZone] for details.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId types.String `tfsdk:"timezone_id"`
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime types.String `tfsdk:"create_time"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag types.String `tfsdk:"etag"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState LifecycleState `tfsdk:"lifecycle_state"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash.
	ParentPath types.String `tfsdk:"parent_path"`
	// The workspace path of the dashboard asset, including the file name.
	Path types.String `tfsdk:"path"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard types.String `tfsdk:"serialized_dashboard"`
	// The timestamp of when the dashboard was last updated by the user.
	UpdateTime types.String `tfsdk:"update_time"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type DashboardView string

const DashboardViewDashboardViewBasic DashboardView = `DASHBOARD_VIEW_BASIC`

const DashboardViewDashboardViewFull DashboardView = `DASHBOARD_VIEW_FULL`

// String representation for [fmt.Print]
func (f *DashboardView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DashboardView) Set(v string) error {
	switch v {
	case `DASHBOARD_VIEW_BASIC`, `DASHBOARD_VIEW_FULL`:
		*f = DashboardView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD_VIEW_BASIC", "DASHBOARD_VIEW_FULL"`, v)
	}
}

// Type always returns DashboardView to satisfy [pflag.Value] interface
func (f *DashboardView) Type() string {
	return "DashboardView"
}

// Delete dashboard schedule
type DeleteScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The etag for the schedule. Optionally, it can be provided to verify that
	// the schedule has not been modified from its last retrieval.
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-" url:"-"`
}

type DeleteScheduleResponse struct {
}

// Delete schedule subscription
type DeleteSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The etag for the subscription. Can be optionally provided to ensure that
	// the subscription has not been modified since the last read.
	Etag types.String `tfsdk:"-" url:"etag,omitempty"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-" url:"-"`
}

type DeleteSubscriptionResponse struct {
}

// Get dashboard
type GetDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-" url:"-"`
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-" url:"-"`
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-" url:"-"`
}

// Get schedule subscription
type GetSubscriptionRequest struct {
	// UUID identifying the dashboard which the subscription belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// UUID identifying the schedule which the subscription belongs.
	ScheduleId types.String `tfsdk:"-" url:"-"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"-" url:"-"`
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

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

// List dashboards
type ListDashboardsRequest struct {
	// The number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDashboards` call. This token
	// can be used to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// The flag to include dashboards located in the trash. If unspecified, only
	// active dashboards will be returned.
	ShowTrashed types.Bool `tfsdk:"-" url:"show_trashed,omitempty"`
	// Indicates whether to include all metadata from the dashboard in the
	// response. If unset, the response defaults to `DASHBOARD_VIEW_BASIC` which
	// only includes summary metadata from the dashboard.
	View DashboardView `tfsdk:"-" url:"view,omitempty"`
}

type ListDashboardsResponse struct {
	Dashboards []Dashboard `tfsdk:"dashboards"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token"`
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The number of schedules to return per page.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Schedules []Schedule `tfsdk:"schedules"`
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The number of subscriptions to return per page.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-" url:"-"`
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Subscriptions []Subscription `tfsdk:"subscriptions"`
}

type MigrateDashboardRequest struct {
	// Display name for the new Lakeview dashboard.
	DisplayName types.String `tfsdk:"display_name"`
	// The workspace path of the folder to contain the migrated Lakeview
	// dashboard.
	ParentPath types.String `tfsdk:"parent_path"`
	// UUID of the dashboard to be migrated.
	SourceDashboardId types.String `tfsdk:"source_dashboard_id"`
}

type PublishRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// Flag to indicate if the publisher's credentials should be embedded in the
	// published dashboard. These embedded credentials will be used to execute
	// the published dashboard's queries.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials"`
	// The ID of the warehouse that can be used to override the warehouse which
	// was set in the draft.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName types.String `tfsdk:"display_name"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime types.String `tfsdk:"create_time"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `tfsdk:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus `tfsdk:"pause_status"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"schedule_id"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
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

// Type always returns SchedulePauseStatus to satisfy [pflag.Value] interface
func (f *SchedulePauseStatus) Type() string {
	return "SchedulePauseStatus"
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber *SubscriptionSubscriberDestination `tfsdk:"destination_subscriber"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber *SubscriptionSubscriberUser `tfsdk:"user_subscriber"`
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime types.String `tfsdk:"create_time"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId types.Int64 `tfsdk:"created_by_user_id"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag types.String `tfsdk:"etag"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"schedule_id"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber Subscriber `tfsdk:"subscriber"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time"`
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId types.String `tfsdk:"destination_id"`
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId types.Int64 `tfsdk:"user_id"`
}

// Trash dashboard
type TrashDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-" url:"-"`
}

type TrashDashboardResponse struct {
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the dashboard to be published.
	DashboardId types.String `tfsdk:"-" url:"-"`
}

type UnpublishDashboardResponse struct {
}

type UpdateDashboardRequest struct {
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read.
	Etag types.String `tfsdk:"etag"`
	// The contents of the dashboard in serialized string form.
	SerializedDashboard types.String `tfsdk:"serialized_dashboard"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type UpdateScheduleRequest struct {
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule CronSchedule `tfsdk:"cron_schedule"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus SchedulePauseStatus `tfsdk:"pause_status"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-" url:"-"`
}
