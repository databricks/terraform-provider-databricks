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
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/terraform-provider-databricks/internal/service/sql_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Create dashboard
type CreateDashboardRequest struct {
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
}

func (newState *CreateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateDashboardRequest) {
}

func (newState *CreateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState CreateDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboard": reflect.TypeOf(Dashboard{}),
	}
}

// ToObjectType returns the representation of CreateDashboardRequest in the Terraform plugin framework type
// system.
func (a CreateDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{
				ElemType: Dashboard{}.ToObjectType(ctx),
			},
		},
	}
}

// Create dashboard schedule
type CreateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`

	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
}

func (newState *CreateScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateScheduleRequest) {
}

func (newState *CreateScheduleRequest) SyncEffectiveFieldsDuringRead(existingState CreateScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule": reflect.TypeOf(Schedule{}),
	}
}

// ToObjectType returns the representation of CreateScheduleRequest in the Terraform plugin framework type
// system.
func (a CreateScheduleRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: Schedule{}.ToObjectType(ctx),
			},
		},
	}
}

// Create schedule subscription
type CreateSubscriptionRequest struct {
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"-"`

	Subscription types.List `tfsdk:"subscription" tf:"optional,object"`
}

func (newState *CreateSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateSubscriptionRequest) {
}

func (newState *CreateSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState CreateSubscriptionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CreateSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CreateSubscriptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscription": reflect.TypeOf(Subscription{}),
	}
}

// ToObjectType returns the representation of CreateSubscriptionRequest in the Terraform plugin framework type
// system.
func (a CreateSubscriptionRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule_id":  types.StringType,
			"subscription": basetypes.ListType{
				ElemType: Subscription{}.ToObjectType(ctx),
			},
		},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in CronSchedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a CronSchedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of CronSchedule in the Terraform plugin framework type
// system.
func (a CronSchedule) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"quartz_cron_expression": types.StringType,
			"timezone_id":            types.StringType,
		},
	}
}

type Dashboard struct {
	// The timestamp of when the dashboard was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed,optional"`
	// The display name of the dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the dashboard. Can be optionally provided on updates to
	// ensure that the dashboard has not been modified since the last read. This
	// field is excluded in List Dashboards responses.
	Etag types.String `tfsdk:"etag" tf:"computed,optional"`
	// The state of the dashboard resource. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// The workspace path of the folder containing the dashboard. Includes
	// leading slash and no trailing slash. This field is excluded in List
	// Dashboards responses.
	ParentPath types.String `tfsdk:"parent_path" tf:"computed,optional"`
	// The workspace path of the dashboard asset, including the file name.
	// Exported dashboards always have the file extension `.lvdash.json`. This
	// field is excluded in List Dashboards responses.
	Path types.String `tfsdk:"path" tf:"computed,optional"`
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
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
	// The warehouse ID used to run the dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Dashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard) {
}

func (newState *Dashboard) SyncEffectiveFieldsDuringRead(existingState Dashboard) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Dashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Dashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Dashboard in the Terraform plugin framework type
// system.
func (a Dashboard) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":          types.StringType,
			"dashboard_id":         types.StringType,
			"display_name":         types.StringType,
			"etag":                 types.StringType,
			"lifecycle_state":      types.StringType,
			"parent_path":          types.StringType,
			"path":                 types.StringType,
			"serialized_dashboard": types.StringType,
			"update_time":          types.StringType,
			"warehouse_id":         types.StringType,
		},
	}
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

func (newState *DeleteScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleRequest) {
}

func (newState *DeleteScheduleRequest) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteScheduleRequest in the Terraform plugin framework type
// system.
func (a DeleteScheduleRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"etag":         types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

type DeleteScheduleResponse struct {
}

func (newState *DeleteScheduleResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteScheduleResponse) {
}

func (newState *DeleteScheduleResponse) SyncEffectiveFieldsDuringRead(existingState DeleteScheduleResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteScheduleResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteScheduleResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteScheduleResponse in the Terraform plugin framework type
// system.
func (a DeleteScheduleResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
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

func (newState *DeleteSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionRequest) {
}

func (newState *DeleteSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSubscriptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteSubscriptionRequest in the Terraform plugin framework type
// system.
func (a DeleteSubscriptionRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":    types.StringType,
			"etag":            types.StringType,
			"schedule_id":     types.StringType,
			"subscription_id": types.StringType,
		},
	}
}

type DeleteSubscriptionResponse struct {
}

func (newState *DeleteSubscriptionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteSubscriptionResponse) {
}

func (newState *DeleteSubscriptionResponse) SyncEffectiveFieldsDuringRead(existingState DeleteSubscriptionResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in DeleteSubscriptionResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a DeleteSubscriptionResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of DeleteSubscriptionResponse in the Terraform plugin framework type
// system.
func (a DeleteSubscriptionResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Genie AI Response
type GenieAttachment struct {
	Query types.List `tfsdk:"query" tf:"optional,object"`

	Text types.List `tfsdk:"text" tf:"optional,object"`
}

func (newState *GenieAttachment) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieAttachment) {
}

func (newState *GenieAttachment) SyncEffectiveFieldsDuringRead(existingState GenieAttachment) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"query": reflect.TypeOf(QueryAttachment{}),
		"text":  reflect.TypeOf(TextAttachment{}),
	}
}

// ToObjectType returns the representation of GenieAttachment in the Terraform plugin framework type
// system.
func (a GenieAttachment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"query": basetypes.ListType{
				ElemType: QueryAttachment{}.ToObjectType(ctx),
			},
			"text": basetypes.ListType{
				ElemType: TextAttachment{}.ToObjectType(ctx),
			},
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieConversation.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieConversation) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GenieConversation in the Terraform plugin framework type
// system.
func (a GenieConversation) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"created_timestamp":      types.Int64Type,
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"space_id":               types.StringType,
			"title":                  types.StringType,
			"user_id":                types.Int64Type,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieCreateConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieCreateConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GenieCreateConversationMessageRequest in the Terraform plugin framework type
// system.
func (a GenieCreateConversationMessageRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":         types.StringType,
			"conversation_id": types.StringType,
			"space_id":        types.StringType,
		},
	}
}

// Execute SQL query in a conversation message
type GenieExecuteMessageQueryRequest struct {
	// Conversation ID
	ConversationId types.String `tfsdk:"-"`
	// Message ID
	MessageId types.String `tfsdk:"-"`
	// Genie space ID
	SpaceId types.String `tfsdk:"-"`
}

func (newState *GenieExecuteMessageQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieExecuteMessageQueryRequest) {
}

func (newState *GenieExecuteMessageQueryRequest) SyncEffectiveFieldsDuringRead(existingState GenieExecuteMessageQueryRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieExecuteMessageQueryRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieExecuteMessageQueryRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GenieExecuteMessageQueryRequest in the Terraform plugin framework type
// system.
func (a GenieExecuteMessageQueryRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GenieGetConversationMessageRequest in the Terraform plugin framework type
// system.
func (a GenieGetConversationMessageRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageQueryResultRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageQueryResultRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GenieGetMessageQueryResultRequest in the Terraform plugin framework type
// system.
func (a GenieGetMessageQueryResultRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation_id": types.StringType,
			"message_id":      types.StringType,
			"space_id":        types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieGetMessageQueryResultResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieGetMessageQueryResultResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"statement_response": reflect.TypeOf(sql.StatementResponse{}),
	}
}

// ToObjectType returns the representation of GenieGetMessageQueryResultResponse in the Terraform plugin framework type
// system.
func (a GenieGetMessageQueryResultResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"statement_response": basetypes.ListType{
				ElemType: sql_tf.StatementResponse{}.ToObjectType(ctx),
			},
		},
	}
}

type GenieMessage struct {
	// AI produced response to the message
	Attachments types.List `tfsdk:"attachments" tf:"optional"`
	// User message content
	Content types.String `tfsdk:"content" tf:""`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`
	// Timestamp when the message was created
	CreatedTimestamp types.Int64 `tfsdk:"created_timestamp" tf:"optional"`
	// Error message if AI failed to respond to the message
	Error types.List `tfsdk:"error" tf:"optional,object"`
	// Message ID
	Id types.String `tfsdk:"id" tf:""`
	// Timestamp when the message was last updated
	LastUpdatedTimestamp types.Int64 `tfsdk:"last_updated_timestamp" tf:"optional"`
	// The result of SQL query if the message has a query attachment
	QueryResult types.List `tfsdk:"query_result" tf:"optional,object"`
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieMessage.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieMessage) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"attachments":  reflect.TypeOf(GenieAttachment{}),
		"error":        reflect.TypeOf(MessageError{}),
		"query_result": reflect.TypeOf(Result{}),
	}
}

// ToObjectType returns the representation of GenieMessage in the Terraform plugin framework type
// system.
func (a GenieMessage) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attachments": basetypes.ListType{
				ElemType: GenieAttachment{}.ToObjectType(ctx),
			},
			"content":           types.StringType,
			"conversation_id":   types.StringType,
			"created_timestamp": types.Int64Type,
			"error": basetypes.ListType{
				ElemType: MessageError{}.ToObjectType(ctx),
			},
			"id":                     types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query_result": basetypes.ListType{
				ElemType: Result{}.ToObjectType(ctx),
			},
			"space_id": types.StringType,
			"status":   types.StringType,
			"user_id":  types.Int64Type,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieStartConversationMessageRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieStartConversationMessageRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GenieStartConversationMessageRequest in the Terraform plugin framework type
// system.
func (a GenieStartConversationMessageRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content":  types.StringType,
			"space_id": types.StringType,
		},
	}
}

type GenieStartConversationResponse struct {
	Conversation types.List `tfsdk:"conversation" tf:"optional,object"`
	// Conversation ID
	ConversationId types.String `tfsdk:"conversation_id" tf:""`

	Message types.List `tfsdk:"message" tf:"optional,object"`
	// Message ID
	MessageId types.String `tfsdk:"message_id" tf:""`
}

func (newState *GenieStartConversationResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GenieStartConversationResponse) {
}

func (newState *GenieStartConversationResponse) SyncEffectiveFieldsDuringRead(existingState GenieStartConversationResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GenieStartConversationResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GenieStartConversationResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"conversation": reflect.TypeOf(GenieConversation{}),
		"message":      reflect.TypeOf(GenieMessage{}),
	}
}

// ToObjectType returns the representation of GenieStartConversationResponse in the Terraform plugin framework type
// system.
func (a GenieStartConversationResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"conversation": basetypes.ListType{
				ElemType: GenieConversation{}.ToObjectType(ctx),
			},
			"conversation_id": types.StringType,
			"message": basetypes.ListType{
				ElemType: GenieMessage{}.ToObjectType(ctx),
			},
			"message_id": types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetDashboardRequest in the Terraform plugin framework type
// system.
func (a GetDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Get published dashboard
type GetPublishedDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetPublishedDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetPublishedDashboardRequest) {
}

func (newState *GetPublishedDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetPublishedDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetPublishedDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetPublishedDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetPublishedDashboardRequest in the Terraform plugin framework type
// system.
func (a GetPublishedDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

// Get dashboard schedule
type GetScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *GetScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetScheduleRequest) {
}

func (newState *GetScheduleRequest) SyncEffectiveFieldsDuringRead(existingState GetScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetScheduleRequest in the Terraform plugin framework type
// system.
func (a GetScheduleRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule_id":  types.StringType,
		},
	}
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

func (newState *GetSubscriptionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetSubscriptionRequest) {
}

func (newState *GetSubscriptionRequest) SyncEffectiveFieldsDuringRead(existingState GetSubscriptionRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in GetSubscriptionRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a GetSubscriptionRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of GetSubscriptionRequest in the Terraform plugin framework type
// system.
func (a GetSubscriptionRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":    types.StringType,
			"schedule_id":     types.StringType,
			"subscription_id": types.StringType,
		},
	}
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

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsRequest) {
}

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringRead(existingState ListDashboardsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDashboardsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDashboardsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListDashboardsRequest in the Terraform plugin framework type
// system.
func (a ListDashboardsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"show_trashed": types.BoolType,
			"view":         types.StringType,
		},
	}
}

type ListDashboardsResponse struct {
	Dashboards types.List `tfsdk:"dashboards" tf:"optional"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent dashboards.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed,optional"`
}

func (newState *ListDashboardsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsResponse) {
}

func (newState *ListDashboardsResponse) SyncEffectiveFieldsDuringRead(existingState ListDashboardsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListDashboardsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListDashboardsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboards": reflect.TypeOf(Dashboard{}),
	}
}

// ToObjectType returns the representation of ListDashboardsResponse in the Terraform plugin framework type
// system.
func (a ListDashboardsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboards": basetypes.ListType{
				ElemType: Dashboard{}.ToObjectType(ctx),
			},
			"next_page_token": types.StringType,
		},
	}
}

// List dashboard schedules
type ListSchedulesRequest struct {
	// UUID identifying the dashboard to which the schedules belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of schedules to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSchedules` call. Use this to
	// retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListSchedulesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesRequest) {
}

func (newState *ListSchedulesRequest) SyncEffectiveFieldsDuringRead(existingState ListSchedulesRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchedulesRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchedulesRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListSchedulesRequest in the Terraform plugin framework type
// system.
func (a ListSchedulesRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
		},
	}
}

type ListSchedulesResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent schedules.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed,optional"`

	Schedules types.List `tfsdk:"schedules" tf:"optional"`
}

func (newState *ListSchedulesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSchedulesResponse) {
}

func (newState *ListSchedulesResponse) SyncEffectiveFieldsDuringRead(existingState ListSchedulesResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSchedulesResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSchedulesResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedules": reflect.TypeOf(Schedule{}),
	}
}

// ToObjectType returns the representation of ListSchedulesResponse in the Terraform plugin framework type
// system.
func (a ListSchedulesResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"schedules": basetypes.ListType{
				ElemType: Schedule{}.ToObjectType(ctx),
			},
		},
	}
}

// List schedule subscriptions
type ListSubscriptionsRequest struct {
	// UUID identifying the dashboard which the subscriptions belongs.
	DashboardId types.String `tfsdk:"-"`
	// The number of subscriptions to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// A page token, received from a previous `ListSubscriptions` call. Use this
	// to retrieve the subsequent page.
	PageToken types.String `tfsdk:"-"`
	// UUID identifying the schedule which the subscriptions belongs.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *ListSubscriptionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsRequest) {
}

func (newState *ListSubscriptionsRequest) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSubscriptionsRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSubscriptionsRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of ListSubscriptionsRequest in the Terraform plugin framework type
// system.
func (a ListSubscriptionsRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"page_size":    types.Int64Type,
			"page_token":   types.StringType,
			"schedule_id":  types.StringType,
		},
	}
}

type ListSubscriptionsResponse struct {
	// A token that can be used as a `page_token` in subsequent requests to
	// retrieve the next page of results. If this field is omitted, there are no
	// subsequent subscriptions.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"computed,optional"`

	Subscriptions types.List `tfsdk:"subscriptions" tf:"optional"`
}

func (newState *ListSubscriptionsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListSubscriptionsResponse) {
}

func (newState *ListSubscriptionsResponse) SyncEffectiveFieldsDuringRead(existingState ListSubscriptionsResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in ListSubscriptionsResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a ListSubscriptionsResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriptions": reflect.TypeOf(Subscription{}),
	}
}

// ToObjectType returns the representation of ListSubscriptionsResponse in the Terraform plugin framework type
// system.
func (a ListSubscriptionsResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"next_page_token": types.StringType,
			"subscriptions": basetypes.ListType{
				ElemType: Subscription{}.ToObjectType(ctx),
			},
		},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MessageError.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MessageError) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of MessageError in the Terraform plugin framework type
// system.
func (a MessageError) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"error": types.StringType,
			"type":  types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in MigrateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a MigrateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of MigrateDashboardRequest in the Terraform plugin framework type
// system.
func (a MigrateDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":        types.StringType,
			"parent_path":         types.StringType,
			"source_dashboard_id": types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PublishRequest in the Terraform plugin framework type
// system.
func (a PublishRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id":      types.StringType,
			"embed_credentials": types.BoolType,
			"warehouse_id":      types.StringType,
		},
	}
}

type PublishedDashboard struct {
	// The display name of the published dashboard.
	DisplayName types.String `tfsdk:"display_name" tf:"computed,optional"`
	// Indicates whether credentials are embedded in the published dashboard.
	EmbedCredentials types.Bool `tfsdk:"embed_credentials" tf:"optional"`
	// The timestamp of when the published dashboard was last revised.
	RevisionCreateTime types.String `tfsdk:"revision_create_time" tf:"computed,optional"`
	// The warehouse ID used to run the published dashboard.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *PublishedDashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan PublishedDashboard) {
}

func (newState *PublishedDashboard) SyncEffectiveFieldsDuringRead(existingState PublishedDashboard) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in PublishedDashboard.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a PublishedDashboard) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of PublishedDashboard in the Terraform plugin framework type
// system.
func (a PublishedDashboard) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"display_name":         types.StringType,
			"embed_credentials":    types.BoolType,
			"revision_create_time": types.StringType,
			"warehouse_id":         types.StringType,
		},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in QueryAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a QueryAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of QueryAttachment in the Terraform plugin framework type
// system.
func (a QueryAttachment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"description":            types.StringType,
			"id":                     types.StringType,
			"instruction_id":         types.StringType,
			"instruction_title":      types.StringType,
			"last_updated_timestamp": types.Int64Type,
			"query":                  types.StringType,
			"title":                  types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Result.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Result) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of Result in the Terraform plugin framework type
// system.
func (a Result) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"is_truncated": types.BoolType,
			"row_count":    types.Int64Type,
			"statement_id": types.StringType,
		},
	}
}

type Schedule struct {
	// A timestamp indicating when the schedule was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// The cron expression describing the frequency of the periodic refresh for
	// this schedule.
	CronSchedule types.List `tfsdk:"cron_schedule" tf:"object"`
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed,optional"`
	// The display name for schedule.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The etag for the schedule. Must be left empty on create, must be provided
	// on updates to ensure that the schedule has not been modified since the
	// last read, and can be optionally provided on delete.
	Etag types.String `tfsdk:"etag" tf:"computed,optional"`
	// The status indicates whether this schedule is paused or not.
	PauseStatus types.String `tfsdk:"pause_status" tf:"optional"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"computed,optional"`
	// A timestamp indicating when the schedule was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
	// The warehouse id to run the dashboard with for the schedule.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Schedule) SyncEffectiveFieldsDuringCreateOrUpdate(plan Schedule) {
}

func (newState *Schedule) SyncEffectiveFieldsDuringRead(existingState Schedule) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Schedule.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Schedule) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"cron_schedule": reflect.TypeOf(CronSchedule{}),
	}
}

// ToObjectType returns the representation of Schedule in the Terraform plugin framework type
// system.
func (a Schedule) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"cron_schedule": basetypes.ListType{
				ElemType: CronSchedule{}.ToObjectType(ctx),
			},
			"dashboard_id": types.StringType,
			"display_name": types.StringType,
			"etag":         types.StringType,
			"pause_status": types.StringType,
			"schedule_id":  types.StringType,
			"update_time":  types.StringType,
			"warehouse_id": types.StringType,
		},
	}
}

type Subscriber struct {
	// The destination to receive the subscription email. This parameter is
	// mutually exclusive with `user_subscriber`.
	DestinationSubscriber types.List `tfsdk:"destination_subscriber" tf:"optional,object"`
	// The user to receive the subscription email. This parameter is mutually
	// exclusive with `destination_subscriber`.
	UserSubscriber types.List `tfsdk:"user_subscriber" tf:"optional,object"`
}

func (newState *Subscriber) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscriber) {
}

func (newState *Subscriber) SyncEffectiveFieldsDuringRead(existingState Subscriber) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Subscriber.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Subscriber) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"destination_subscriber": reflect.TypeOf(SubscriptionSubscriberDestination{}),
		"user_subscriber":        reflect.TypeOf(SubscriptionSubscriberUser{}),
	}
}

// ToObjectType returns the representation of Subscriber in the Terraform plugin framework type
// system.
func (a Subscriber) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_subscriber": basetypes.ListType{
				ElemType: SubscriptionSubscriberDestination{}.ToObjectType(ctx),
			},
			"user_subscriber": basetypes.ListType{
				ElemType: SubscriptionSubscriberUser{}.ToObjectType(ctx),
			},
		},
	}
}

type Subscription struct {
	// A timestamp indicating when the subscription was created.
	CreateTime types.String `tfsdk:"create_time" tf:"computed,optional"`
	// UserId of the user who adds subscribers (users or notification
	// destinations) to the dashboard's schedule.
	CreatedByUserId types.Int64 `tfsdk:"created_by_user_id" tf:"computed,optional"`
	// UUID identifying the dashboard to which the subscription belongs.
	DashboardId types.String `tfsdk:"dashboard_id" tf:"computed,optional"`
	// The etag for the subscription. Must be left empty on create, can be
	// optionally provided on delete to ensure that the subscription has not
	// been deleted since the last read.
	Etag types.String `tfsdk:"etag" tf:"computed,optional"`
	// UUID identifying the schedule to which the subscription belongs.
	ScheduleId types.String `tfsdk:"schedule_id" tf:"computed,optional"`
	// Subscriber details for users and destinations to be added as subscribers
	// to the schedule.
	Subscriber types.List `tfsdk:"subscriber" tf:"object"`
	// UUID identifying the subscription.
	SubscriptionId types.String `tfsdk:"subscription_id" tf:"computed,optional"`
	// A timestamp indicating when the subscription was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"computed,optional"`
}

func (newState *Subscription) SyncEffectiveFieldsDuringCreateOrUpdate(plan Subscription) {
}

func (newState *Subscription) SyncEffectiveFieldsDuringRead(existingState Subscription) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in Subscription.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a Subscription) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"subscriber": reflect.TypeOf(Subscriber{}),
	}
}

// ToObjectType returns the representation of Subscription in the Terraform plugin framework type
// system.
func (a Subscription) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":        types.StringType,
			"created_by_user_id": types.Int64Type,
			"dashboard_id":       types.StringType,
			"etag":               types.StringType,
			"schedule_id":        types.StringType,
			"subscriber": basetypes.ListType{
				ElemType: Subscriber{}.ToObjectType(ctx),
			},
			"subscription_id": types.StringType,
			"update_time":     types.StringType,
		},
	}
}

type SubscriptionSubscriberDestination struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId types.String `tfsdk:"destination_id" tf:"computed,optional"`
}

func (newState *SubscriptionSubscriberDestination) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberDestination) {
}

func (newState *SubscriptionSubscriberDestination) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberDestination) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubscriptionSubscriberDestination.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubscriptionSubscriberDestination) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SubscriptionSubscriberDestination in the Terraform plugin framework type
// system.
func (a SubscriptionSubscriberDestination) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"destination_id": types.StringType,
		},
	}
}

type SubscriptionSubscriberUser struct {
	// UserId of the subscriber.
	UserId types.Int64 `tfsdk:"user_id" tf:"computed,optional"`
}

func (newState *SubscriptionSubscriberUser) SyncEffectiveFieldsDuringCreateOrUpdate(plan SubscriptionSubscriberUser) {
}

func (newState *SubscriptionSubscriberUser) SyncEffectiveFieldsDuringRead(existingState SubscriptionSubscriberUser) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in SubscriptionSubscriberUser.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a SubscriptionSubscriberUser) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of SubscriptionSubscriberUser in the Terraform plugin framework type
// system.
func (a SubscriptionSubscriberUser) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.Int64Type,
		},
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TextAttachment.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TextAttachment) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TextAttachment in the Terraform plugin framework type
// system.
func (a TextAttachment) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"content": types.StringType,
			"id":      types.StringType,
		},
	}
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

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TrashDashboardRequest in the Terraform plugin framework type
// system.
func (a TrashDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type TrashDashboardResponse struct {
}

func (newState *TrashDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashDashboardResponse) {
}

func (newState *TrashDashboardResponse) SyncEffectiveFieldsDuringRead(existingState TrashDashboardResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in TrashDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a TrashDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of TrashDashboardResponse in the Terraform plugin framework type
// system.
func (a TrashDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Unpublish dashboard
type UnpublishDashboardRequest struct {
	// UUID identifying the published dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UnpublishDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardRequest) {
}

func (newState *UnpublishDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpublishDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpublishDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of UnpublishDashboardRequest in the Terraform plugin framework type
// system.
func (a UnpublishDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
		},
	}
}

type UnpublishDashboardResponse struct {
}

func (newState *UnpublishDashboardResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UnpublishDashboardResponse) {
}

func (newState *UnpublishDashboardResponse) SyncEffectiveFieldsDuringRead(existingState UnpublishDashboardResponse) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UnpublishDashboardResponse.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UnpublishDashboardResponse) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// ToObjectType returns the representation of UnpublishDashboardResponse in the Terraform plugin framework type
// system.
func (a UnpublishDashboardResponse) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{},
	}
}

// Update dashboard
type UpdateDashboardRequest struct {
	Dashboard types.List `tfsdk:"dashboard" tf:"optional,object"`
	// UUID identifying the dashboard.
	DashboardId types.String `tfsdk:"-"`
}

func (newState *UpdateDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateDashboardRequest) {
}

func (newState *UpdateDashboardRequest) SyncEffectiveFieldsDuringRead(existingState UpdateDashboardRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateDashboardRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateDashboardRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"dashboard": reflect.TypeOf(Dashboard{}),
	}
}

// ToObjectType returns the representation of UpdateDashboardRequest in the Terraform plugin framework type
// system.
func (a UpdateDashboardRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard": basetypes.ListType{
				ElemType: Dashboard{}.ToObjectType(ctx),
			},
			"dashboard_id": types.StringType,
		},
	}
}

// Update dashboard schedule
type UpdateScheduleRequest struct {
	// UUID identifying the dashboard to which the schedule belongs.
	DashboardId types.String `tfsdk:"-"`

	Schedule types.List `tfsdk:"schedule" tf:"optional,object"`
	// UUID identifying the schedule.
	ScheduleId types.String `tfsdk:"-"`
}

func (newState *UpdateScheduleRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateScheduleRequest) {
}

func (newState *UpdateScheduleRequest) SyncEffectiveFieldsDuringRead(existingState UpdateScheduleRequest) {
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in UpdateScheduleRequest.
// Container types (types.Map, types.List, types.Set) and object types (types.Object) do not carry
// the type information of their elements in the Go type system. This function provides a way to
// retrieve the type information of the elements in complex fields at runtime. The values of the map
// are the reflected types of the contained elements. They must be either primitive values from the
// plugin framework type system (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF
// SDK values.
func (a UpdateScheduleRequest) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"schedule": reflect.TypeOf(Schedule{}),
	}
}

// ToObjectType returns the representation of UpdateScheduleRequest in the Terraform plugin framework type
// system.
func (a UpdateScheduleRequest) ToObjectType(ctx context.Context) types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"dashboard_id": types.StringType,
			"schedule": basetypes.ListType{
				ElemType: Schedule{}.ToObjectType(ctx),
			},
			"schedule_id": types.StringType,
		},
	}
}
