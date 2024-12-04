// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package sql_tf

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessControl struct {
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`

	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *AccessControl) SyncEffectiveFieldsDuringCreateOrUpdate(plan AccessControl) {
}

func (newState *AccessControl) SyncEffectiveFieldsDuringRead(existingState AccessControl) {
}

func (a AccessControl) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type Alert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition" tf:"optional,object"`
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// The workspace path of the folder containing the alert.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State types.String `tfsdk:"state" tf:"optional"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime types.String `tfsdk:"trigger_time" tf:"optional"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *Alert) SyncEffectiveFieldsDuringCreateOrUpdate(plan Alert) {
}

func (newState *Alert) SyncEffectiveFieldsDuringRead(existingState Alert) {
}

func (a Alert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Condition": reflect.TypeOf(AlertCondition{}),
	}
}

type AlertCondition struct {
	// Alert state if result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state" tf:"optional"`
	// Operator used for comparison in alert evaluation.
	Op types.String `tfsdk:"op" tf:"optional"`
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	Operand types.Object `tfsdk:"operand" tf:"optional,object"`
	// Threshold value used for comparison in alert evaluation.
	Threshold types.Object `tfsdk:"threshold" tf:"optional,object"`
}

func (newState *AlertCondition) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertCondition) {
}

func (newState *AlertCondition) SyncEffectiveFieldsDuringRead(existingState AlertCondition) {
}

func (a AlertCondition) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Operand":   reflect.TypeOf(AlertConditionOperand{}),
		"Threshold": reflect.TypeOf(AlertConditionThreshold{}),
	}
}

type AlertConditionOperand struct {
	Column types.Object `tfsdk:"column" tf:"optional,object"`
}

func (newState *AlertConditionOperand) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConditionOperand) {
}

func (newState *AlertConditionOperand) SyncEffectiveFieldsDuringRead(existingState AlertConditionOperand) {
}

func (a AlertConditionOperand) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Column": reflect.TypeOf(AlertOperandColumn{}),
	}
}

type AlertConditionThreshold struct {
	Value types.Object `tfsdk:"value" tf:"optional,object"`
}

func (newState *AlertConditionThreshold) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertConditionThreshold) {
}

func (newState *AlertConditionThreshold) SyncEffectiveFieldsDuringRead(existingState AlertConditionThreshold) {
}

func (a AlertConditionThreshold) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Value": reflect.TypeOf(AlertOperandValue{}),
	}
}

type AlertOperandColumn struct {
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *AlertOperandColumn) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOperandColumn) {
}

func (newState *AlertOperandColumn) SyncEffectiveFieldsDuringRead(existingState AlertOperandColumn) {
}

func (a AlertOperandColumn) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type AlertOperandValue struct {
	BoolValue types.Bool `tfsdk:"bool_value" tf:"optional"`

	DoubleValue types.Float64 `tfsdk:"double_value" tf:"optional"`

	StringValue types.String `tfsdk:"string_value" tf:"optional"`
}

func (newState *AlertOperandValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOperandValue) {
}

func (newState *AlertOperandValue) SyncEffectiveFieldsDuringRead(existingState AlertOperandValue) {
}

func (a AlertOperandValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	Column types.String `tfsdk:"column" tf:""`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// State that alert evaluates to when query result is empty.
	EmptyResultState types.String `tfsdk:"empty_result_state" tf:"optional"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted types.Bool `tfsdk:"muted" tf:"optional"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op types.String `tfsdk:"op" tf:""`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value any `tfsdk:"value" tf:""`
}

func (newState *AlertOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertOptions) {
}

func (newState *AlertOptions) SyncEffectiveFieldsDuringRead(existingState AlertOptions) {
}

func (a AlertOptions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type AlertQuery struct {
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Query ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived" tf:"optional"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft" tf:"optional"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`

	Options types.Object `tfsdk:"options" tf:"optional,object"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`
}

func (newState *AlertQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan AlertQuery) {
}

func (newState *AlertQuery) SyncEffectiveFieldsDuringRead(existingState AlertQuery) {
}

func (a AlertQuery) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options": reflect.TypeOf(QueryOptions{}),
		"Tags":    reflect.TypeOf(""),
	}
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count" tf:"optional"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index" tf:"optional"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset" tf:"optional"`
}

func (newState *BaseChunkInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan BaseChunkInfo) {
}

func (newState *BaseChunkInfo) SyncEffectiveFieldsDuringRead(existingState BaseChunkInfo) {
}

func (a BaseChunkInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Cancel statement execution
type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (newState *CancelExecutionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelExecutionRequest) {
}

func (newState *CancelExecutionRequest) SyncEffectiveFieldsDuringRead(existingState CancelExecutionRequest) {
}

func (a CancelExecutionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type CancelExecutionResponse struct {
}

func (newState *CancelExecutionResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CancelExecutionResponse) {
}

func (newState *CancelExecutionResponse) SyncEffectiveFieldsDuringRead(existingState CancelExecutionResponse) {
}

func (a CancelExecutionResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {
	DbsqlVersion types.String `tfsdk:"dbsql_version" tf:"optional"`

	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *Channel) SyncEffectiveFieldsDuringCreateOrUpdate(plan Channel) {
}

func (newState *Channel) SyncEffectiveFieldsDuringRead(existingState Channel) {
}

func (a Channel) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	DbsqlVersion types.String `tfsdk:"dbsql_version" tf:"optional"`
	// Name of the channel
	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *ChannelInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ChannelInfo) {
}

func (newState *ChannelInfo) SyncEffectiveFieldsDuringRead(existingState ChannelInfo) {
}

func (a ChannelInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type ColumnInfo struct {
	// The name of the column.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The ordinal position of the column (starting at position 0).
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The format of the interval type.
	TypeIntervalType types.String `tfsdk:"type_interval_type" tf:"optional"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName types.String `tfsdk:"type_name" tf:"optional"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision types.Int64 `tfsdk:"type_precision" tf:"optional"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale types.Int64 `tfsdk:"type_scale" tf:"optional"`
	// The full SQL type specification.
	TypeText types.String `tfsdk:"type_text" tf:"optional"`
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan ColumnInfo) {
}

func (newState *ColumnInfo) SyncEffectiveFieldsDuringRead(existingState ColumnInfo) {
}

func (a ColumnInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type CreateAlert struct {
	// Name of the alert.
	Name types.String `tfsdk:"name" tf:""`
	// Alert configuration options.
	Options types.Object `tfsdk:"options" tf:"object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id" tf:""`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm" tf:"optional"`
}

func (newState *CreateAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlert) {
}

func (newState *CreateAlert) SyncEffectiveFieldsDuringRead(existingState CreateAlert) {
}

func (a CreateAlert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options": reflect.TypeOf(AlertOptions{}),
	}
}

type CreateAlertRequest struct {
	Alert types.Object `tfsdk:"alert" tf:"optional,object"`
}

func (newState *CreateAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlertRequest) {
}

func (newState *CreateAlertRequest) SyncEffectiveFieldsDuringRead(existingState CreateAlertRequest) {
}

func (a CreateAlertRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Alert": reflect.TypeOf(CreateAlertRequestAlert{}),
	}
}

type CreateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition" tf:"optional,object"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The workspace path of the folder containing the alert.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
}

func (newState *CreateAlertRequestAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateAlertRequestAlert) {
}

func (newState *CreateAlertRequestAlert) SyncEffectiveFieldsDuringRead(existingState CreateAlertRequestAlert) {
}

func (a CreateAlertRequestAlert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Condition": reflect.TypeOf(AlertCondition{}),
	}
}

type CreateQueryRequest struct {
	Query types.Object `tfsdk:"query" tf:"optional,object"`
}

func (newState *CreateQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryRequest) {
}

func (newState *CreateQueryRequest) SyncEffectiveFieldsDuringRead(existingState CreateQueryRequest) {
}

func (a CreateQueryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Query": reflect.TypeOf(CreateQueryRequestQuery{}),
	}
}

type CreateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Workspace path of the workspace folder containing the object.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *CreateQueryRequestQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryRequestQuery) {
}

func (newState *CreateQueryRequestQuery) SyncEffectiveFieldsDuringRead(existingState CreateQueryRequestQuery) {
}

func (a CreateQueryRequestQuery) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(QueryParameter{}),
		"Tags":       reflect.TypeOf(""),
	}
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `tfsdk:"options" tf:""`
	// The identifier returned by :method:queries/create
	QueryId types.String `tfsdk:"query_id" tf:""`
	// The type of visualization: chart, table, pivot table, and so on.
	Type types.String `tfsdk:"type" tf:""`
}

func (newState *CreateQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateQueryVisualizationsLegacyRequest) {
}

func (newState *CreateQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState CreateQueryVisualizationsLegacyRequest) {
}

func (a CreateQueryVisualizationsLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type CreateVisualizationRequest struct {
	Visualization types.Object `tfsdk:"visualization" tf:"optional,object"`
}

func (newState *CreateVisualizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVisualizationRequest) {
}

func (newState *CreateVisualizationRequest) SyncEffectiveFieldsDuringRead(existingState CreateVisualizationRequest) {
}

func (a CreateVisualizationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Visualization": reflect.TypeOf(CreateVisualizationRequestVisualization{}),
	}
}

type CreateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID of the query that the visualization is attached to.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options" tf:"optional"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan" tf:"optional"`
	// The type of visualization: counter, table, funnel, and so on.
	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *CreateVisualizationRequestVisualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateVisualizationRequestVisualization) {
}

func (newState *CreateVisualizationRequestVisualization) SyncEffectiveFieldsDuringRead(existingState CreateVisualizationRequestVisualization) {
}

func (a CreateVisualizationRequestVisualization) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type CreateWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be >= 0 mins for serverless warehouses - Must be
	// == 0 or >= 10 mins for non-serverless warehouses - 0 indicates no
	// autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.Object `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.Object `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *CreateWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWarehouseRequest) {
}

func (newState *CreateWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState CreateWarehouseRequest) {
}

func (a CreateWarehouseRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Channel": reflect.TypeOf(Channel{}),
		"Tags":    reflect.TypeOf(EndpointTags{}),
	}
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id types.String `tfsdk:"id" tf:"optional"`
}

func (newState *CreateWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWarehouseResponse) {
}

func (newState *CreateWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState CreateWarehouseResponse) {
}

func (a CreateWarehouseResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId types.String `tfsdk:"dashboard_id" tf:""`
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`

	Options types.Object `tfsdk:"options" tf:"object"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text" tf:"optional"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id" tf:"optional"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width" tf:""`
}

func (newState *CreateWidget) SyncEffectiveFieldsDuringCreateOrUpdate(plan CreateWidget) {
}

func (newState *CreateWidget) SyncEffectiveFieldsDuringRead(existingState CreateWidget) {
}

func (a CreateWidget) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options": reflect.TypeOf(WidgetOptions{}),
	}
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit types.Bool `tfsdk:"can_edit" tf:"optional"`
	// Timestamp when this dashboard was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled" tf:"optional"`
	// The ID for this dashboard.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived" tf:"optional"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft types.Bool `tfsdk:"is_draft" tf:"optional"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite" tf:"optional"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name" tf:"optional"`

	Options types.Object `tfsdk:"options" tf:"optional,object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier types.String `tfsdk:"permission_tier" tf:"optional"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug types.String `tfsdk:"slug" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`

	User types.Object `tfsdk:"user" tf:"optional,object"`
	// The ID of the user who owns the dashboard.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`

	Widgets types.List `tfsdk:"widgets" tf:"optional"`
}

func (newState *Dashboard) SyncEffectiveFieldsDuringCreateOrUpdate(plan Dashboard) {
}

func (newState *Dashboard) SyncEffectiveFieldsDuringRead(existingState Dashboard) {
}

func (a Dashboard) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options": reflect.TypeOf(DashboardOptions{}),
		"Tags":    reflect.TypeOf(""),
		"User":    reflect.TypeOf(User{}),
		"Widgets": reflect.TypeOf(Widget{}),
	}
}

type DashboardEditContent struct {
	DashboardId types.String `tfsdk:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *DashboardEditContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardEditContent) {
}

func (newState *DashboardEditContent) SyncEffectiveFieldsDuringRead(existingState DashboardEditContent) {
}

func (a DashboardEditContent) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(""),
	}
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at" tf:"optional"`
}

func (newState *DashboardOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardOptions) {
}

func (newState *DashboardOptions) SyncEffectiveFieldsDuringRead(existingState DashboardOptions) {
}

func (a DashboardOptions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type DashboardPostContent struct {
	// Indicates whether the dashboard filters are enabled
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled" tf:"optional"`
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	IsFavorite types.Bool `tfsdk:"is_favorite" tf:"optional"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name" tf:""`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *DashboardPostContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan DashboardPostContent) {
}

func (newState *DashboardPostContent) SyncEffectiveFieldsDuringRead(existingState DashboardPostContent) {
}

func (a DashboardPostContent) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(""),
	}
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	Id types.String `tfsdk:"id" tf:"optional"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Reserved for internal use.
	PauseReason types.String `tfsdk:"pause_reason" tf:"optional"`
	// Reserved for internal use.
	Paused types.Int64 `tfsdk:"paused" tf:"optional"`
	// Reserved for internal use.
	SupportsAutoLimit types.Bool `tfsdk:"supports_auto_limit" tf:"optional"`
	// Reserved for internal use.
	Syntax types.String `tfsdk:"syntax" tf:"optional"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type types.String `tfsdk:"type" tf:"optional"`
	// Reserved for internal use.
	ViewOnly types.Bool `tfsdk:"view_only" tf:"optional"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *DataSource) SyncEffectiveFieldsDuringCreateOrUpdate(plan DataSource) {
}

func (newState *DataSource) SyncEffectiveFieldsDuringRead(existingState DataSource) {
}

func (a DataSource) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type DateRange struct {
	End types.String `tfsdk:"end" tf:""`

	Start types.String `tfsdk:"start" tf:""`
}

func (newState *DateRange) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateRange) {
}

func (newState *DateRange) SyncEffectiveFieldsDuringRead(existingState DateRange) {
}

func (a DateRange) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type DateRangeValue struct {
	// Manually specified date-time range value.
	DateRangeValue types.Object `tfsdk:"date_range_value" tf:"optional,object"`
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue types.String `tfsdk:"dynamic_date_range_value" tf:"optional"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision" tf:"optional"`

	StartDayOfWeek types.Int64 `tfsdk:"start_day_of_week" tf:"optional"`
}

func (newState *DateRangeValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateRangeValue) {
}

func (newState *DateRangeValue) SyncEffectiveFieldsDuringRead(existingState DateRangeValue) {
}

func (a DateRangeValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DateRangeValue": reflect.TypeOf(DateRange{}),
	}
}

type DateValue struct {
	// Manually specified date-time value.
	DateValue types.String `tfsdk:"date_value" tf:"optional"`
	// Dynamic date-time value based on current date-time.
	DynamicDateValue types.String `tfsdk:"dynamic_date_value" tf:"optional"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision types.String `tfsdk:"precision" tf:"optional"`
}

func (newState *DateValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan DateValue) {
}

func (newState *DateValue) SyncEffectiveFieldsDuringRead(existingState DateValue) {
}

func (a DateValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Delete an alert
type DeleteAlertsLegacyRequest struct {
	AlertId types.String `tfsdk:"-"`
}

func (newState *DeleteAlertsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteAlertsLegacyRequest) {
}

func (newState *DeleteAlertsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteAlertsLegacyRequest) {
}

func (a DeleteAlertsLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Remove a dashboard
type DeleteDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (newState *DeleteDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDashboardRequest) {
}

func (newState *DeleteDashboardRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDashboardRequest) {
}

func (a DeleteDashboardRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Remove widget
type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteDashboardWidgetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteDashboardWidgetRequest) {
}

func (newState *DeleteDashboardWidgetRequest) SyncEffectiveFieldsDuringRead(existingState DeleteDashboardWidgetRequest) {
}

func (a DeleteDashboardWidgetRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Delete a query
type DeleteQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (newState *DeleteQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQueriesLegacyRequest) {
}

func (newState *DeleteQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteQueriesLegacyRequest) {
}

func (a DeleteQueriesLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Remove visualization
type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvizualisations/create
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteQueryVisualizationsLegacyRequest) {
}

func (newState *DeleteQueryVisualizationsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState DeleteQueryVisualizationsLegacyRequest) {
}

func (a DeleteQueryVisualizationsLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type DeleteResponse struct {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteResponse) {
}

func (newState *DeleteResponse) SyncEffectiveFieldsDuringRead(existingState DeleteResponse) {
}

func (a DeleteResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Remove a visualization
type DeleteVisualizationRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteVisualizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteVisualizationRequest) {
}

func (newState *DeleteVisualizationRequest) SyncEffectiveFieldsDuringRead(existingState DeleteVisualizationRequest) {
}

func (a DeleteVisualizationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Delete a warehouse
type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *DeleteWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWarehouseRequest) {
}

func (newState *DeleteWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState DeleteWarehouseRequest) {
}

func (a DeleteWarehouseRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type DeleteWarehouseResponse struct {
}

func (newState *DeleteWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan DeleteWarehouseResponse) {
}

func (newState *DeleteWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState DeleteWarehouseResponse) {
}

func (a DeleteWarehouseResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type EditAlert struct {
	AlertId types.String `tfsdk:"-"`
	// Name of the alert.
	Name types.String `tfsdk:"name" tf:""`
	// Alert configuration options.
	Options types.Object `tfsdk:"options" tf:"object"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id" tf:""`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm" tf:"optional"`
}

func (newState *EditAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditAlert) {
}

func (newState *EditAlert) SyncEffectiveFieldsDuringRead(existingState EditAlert) {
}

func (a EditAlert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options": reflect.TypeOf(AlertOptions{}),
	}
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.Object `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute.
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Required. Id of the warehouse to configure.
	Id types.String `tfsdk:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.Object `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *EditWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditWarehouseRequest) {
}

func (newState *EditWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState EditWarehouseRequest) {
}

func (a EditWarehouseRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Channel": reflect.TypeOf(Channel{}),
		"Tags":    reflect.TypeOf(EndpointTags{}),
	}
}

type EditWarehouseResponse struct {
}

func (newState *EditWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan EditWarehouseResponse) {
}

func (newState *EditWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState EditWarehouseResponse) {
}

func (a EditWarehouseResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty struct {
}

func (newState *Empty) SyncEffectiveFieldsDuringCreateOrUpdate(plan Empty) {
}

func (newState *Empty) SyncEffectiveFieldsDuringRead(existingState Empty) {
}

func (a Empty) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type EndpointConfPair struct {
	Key types.String `tfsdk:"key" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *EndpointConfPair) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointConfPair) {
}

func (newState *EndpointConfPair) SyncEffectiveFieldsDuringRead(existingState EndpointConfPair) {
}

func (a EndpointConfPair) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	Details types.String `tfsdk:"details" tf:"optional"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason types.Object `tfsdk:"failure_reason" tf:"optional,object"`
	// Deprecated. split into summary and details for security
	Message types.String `tfsdk:"message" tf:"optional"`
	// Health status of the warehouse.
	Status types.String `tfsdk:"status" tf:"optional"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary types.String `tfsdk:"summary" tf:"optional"`
}

func (newState *EndpointHealth) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointHealth) {
}

func (newState *EndpointHealth) SyncEffectiveFieldsDuringRead(existingState EndpointHealth) {
}

func (a EndpointHealth) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FailureReason": reflect.TypeOf(TerminationReason{}),
	}
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.Object `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health types.Object `tfsdk:"health" tf:"optional,object"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id" tf:"optional"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions" tf:"optional"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters" tf:"optional"`
	// ODBC parameters for the SQL warehouse
	OdbcParams types.Object `tfsdk:"odbc_params" tf:"optional,object"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// State of the warehouse
	State types.String `tfsdk:"state" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.Object `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *EndpointInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointInfo) {
}

func (newState *EndpointInfo) SyncEffectiveFieldsDuringRead(existingState EndpointInfo) {
}

func (a EndpointInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Channel":    reflect.TypeOf(Channel{}),
		"Health":     reflect.TypeOf(EndpointHealth{}),
		"OdbcParams": reflect.TypeOf(OdbcParams{}),
		"Tags":       reflect.TypeOf(EndpointTags{}),
	}
}

type EndpointTagPair struct {
	Key types.String `tfsdk:"key" tf:"optional"`

	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *EndpointTagPair) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTagPair) {
}

func (newState *EndpointTagPair) SyncEffectiveFieldsDuringRead(existingState EndpointTagPair) {
}

func (a EndpointTagPair) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type EndpointTags struct {
	CustomTags types.List `tfsdk:"custom_tags" tf:"optional"`
}

func (newState *EndpointTags) SyncEffectiveFieldsDuringCreateOrUpdate(plan EndpointTags) {
}

func (newState *EndpointTags) SyncEffectiveFieldsDuringRead(existingState EndpointTags) {
}

func (a EndpointTags) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"CustomTags": reflect.TypeOf(EndpointTagPair{}),
	}
}

type EnumValue struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions types.String `tfsdk:"enum_options" tf:"optional"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.Object `tfsdk:"multi_values_options" tf:"optional,object"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *EnumValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan EnumValue) {
}

func (newState *EnumValue) SyncEffectiveFieldsDuringRead(existingState EnumValue) {
}

func (a EnumValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"MultiValuesOptions": reflect.TypeOf(MultiValuesOptions{}),
		"Values":             reflect.TypeOf(""),
	}
}

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	ByteLimit types.Int64 `tfsdk:"byte_limit" tf:"optional"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog types.String `tfsdk:"catalog" tf:"optional"`

	Disposition types.String `tfsdk:"disposition" tf:"optional"`
	// Statement execution supports three result formats: `JSON_ARRAY`
	// (default), `ARROW_STREAM`, and `CSV`.
	//
	// Important: The formats `ARROW_STREAM` and `CSV` are supported only with
	// `EXTERNAL_LINKS` disposition. `JSON_ARRAY` is supported in `INLINE` and
	// `EXTERNAL_LINKS` disposition.
	//
	// When specifying `format=JSON_ARRAY`, result data will be formatted as an
	// array of arrays of values, where each value is either the *string
	// representation* of a value, or `null`. For example, the output of `SELECT
	// concat('id-', id) AS strCol, id AS intCol, null AS nullCol FROM range(3)`
	// would look like this:
	//
	// ``` [ [ "id-1", "1", null ], [ "id-2", "2", null ], [ "id-3", "3", null
	// ], ] ```
	//
	// When specifying `format=JSON_ARRAY` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result contains compact JSON with no indentation or
	// extra whitespace.
	//
	// When specifying `format=ARROW_STREAM` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result will be formatted as Apache Arrow Stream. See
	// the [Apache Arrow streaming format].
	//
	// When specifying `format=CSV` and `disposition=EXTERNAL_LINKS`, each chunk
	// in the result will be a CSV according to [RFC 4180] standard. All the
	// columns values will have *string representation* similar to the
	// `JSON_ARRAY` format, and `null` values will be encoded as “null”.
	// Only the first chunk in the result would contain a header row with column
	// names. For example, the output of `SELECT concat('id-', id) AS strCol, id
	// AS intCol, null as nullCol FROM range(3)` would look like this:
	//
	// ``` strCol,intCol,nullCol id-1,1,null id-2,2,null id-3,3,null ```
	//
	// [Apache Arrow streaming format]: https://arrow.apache.org/docs/format/Columnar.html#ipc-streaming-format
	// [RFC 4180]: https://www.rfc-editor.org/rfc/rfc4180
	Format types.String `tfsdk:"format" tf:"optional"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	OnWaitTimeout types.String `tfsdk:"on_wait_timeout" tf:"optional"`
	// A list of parameters to pass into a SQL statement containing parameter
	// markers. A parameter consists of a name, a value, and optionally a type.
	// To represent a NULL value, the `value` field may be omitted or set to
	// `null` explicitly. If the `type` field is omitted, the value is
	// interpreted as a string.
	//
	// If the type is given, parameters will be checked for type correctness
	// according to the given type. A value is correct if the provided string
	// can be converted to the requested type using the `cast` function. The
	// exact semantics are described in the section [`cast` function] of the SQL
	// language reference.
	//
	// For example, the following statement contains two parameters, `my_name`
	// and `my_date`:
	//
	// SELECT * FROM my_table WHERE name = :my_name AND date = :my_date
	//
	// The parameters can be passed in the request body as follows:
	//
	// { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	//
	// Currently, positional parameters denoted by a `?` marker are not
	// supported by the Databricks SQL Statement Execution API.
	//
	// Also see the section [Parameter markers] of the SQL language reference.
	//
	// [Parameter markers]: https://docs.databricks.com/sql/language-manual/sql-ref-parameter-marker.html
	// [`cast` function]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit types.Int64 `tfsdk:"row_limit" tf:"optional"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema types.String `tfsdk:"schema" tf:"optional"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`.
	Statement types.String `tfsdk:"statement" tf:""`
	// The time in seconds the call will wait for the statement's result set as
	// `Ns`, where `N` can be set to 0 or to a value between 5 and 50.
	//
	// When set to `0s`, the statement will execute in asynchronous mode and the
	// call will not wait for the execution to finish. In this case, the call
	// returns directly with `PENDING` state and a statement ID which can be
	// used for polling with :method:statementexecution/getStatement.
	//
	// When set between 5 and 50 seconds, the call will behave synchronously up
	// to this timeout and wait for the statement execution to finish. If the
	// execution finishes within this time, the call returns immediately with a
	// manifest and result data (or a `FAILED` state in case of an execution
	// error). If the statement takes longer to execute, `on_wait_timeout`
	// determines what should happen after the timeout is reached.
	WaitTimeout types.String `tfsdk:"wait_timeout" tf:"optional"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId types.String `tfsdk:"warehouse_id" tf:""`
}

func (newState *ExecuteStatementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExecuteStatementRequest) {
}

func (newState *ExecuteStatementRequest) SyncEffectiveFieldsDuringRead(existingState ExecuteStatementRequest) {
}

func (a ExecuteStatementRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(StatementParameterListItem{}),
	}
}

type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count" tf:"optional"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index" tf:"optional"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration types.String `tfsdk:"expiration" tf:"optional"`

	ExternalLink types.String `tfsdk:"external_link" tf:"optional"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders types.Map `tfsdk:"http_headers" tf:"optional"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index" tf:"optional"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link" tf:"optional"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset" tf:"optional"`
}

func (newState *ExternalLink) SyncEffectiveFieldsDuringCreateOrUpdate(plan ExternalLink) {
}

func (newState *ExternalLink) SyncEffectiveFieldsDuringRead(existingState ExternalLink) {
}

func (a ExternalLink) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"HttpHeaders": reflect.TypeOf(""),
	}
}

// Get an alert
type GetAlertRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAlertRequest) {
}

func (newState *GetAlertRequest) SyncEffectiveFieldsDuringRead(existingState GetAlertRequest) {
}

func (a GetAlertRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get an alert
type GetAlertsLegacyRequest struct {
	AlertId types.String `tfsdk:"-"`
}

func (newState *GetAlertsLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetAlertsLegacyRequest) {
}

func (newState *GetAlertsLegacyRequest) SyncEffectiveFieldsDuringRead(existingState GetAlertsLegacyRequest) {
}

func (a GetAlertsLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Retrieve a definition
type GetDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (newState *GetDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDashboardRequest) {
}

func (newState *GetDashboardRequest) SyncEffectiveFieldsDuringRead(existingState GetDashboardRequest) {
}

func (a GetDashboardRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get object ACL
type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permissions to check.
	ObjectType types.String `tfsdk:"-"`
}

func (newState *GetDbsqlPermissionRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetDbsqlPermissionRequest) {
}

func (newState *GetDbsqlPermissionRequest) SyncEffectiveFieldsDuringRead(existingState GetDbsqlPermissionRequest) {
}

func (a GetDbsqlPermissionRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get a query definition.
type GetQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (newState *GetQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQueriesLegacyRequest) {
}

func (newState *GetQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState GetQueriesLegacyRequest) {
}

func (a GetQueriesLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get a query
type GetQueryRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *GetQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetQueryRequest) {
}

func (newState *GetQueryRequest) SyncEffectiveFieldsDuringRead(existingState GetQueryRequest) {
}

func (a GetQueryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type GetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id" tf:"optional"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *GetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetResponse) {
}

func (newState *GetResponse) SyncEffectiveFieldsDuringRead(existingState GetResponse) {
}

func (a GetResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(AccessControl{}),
	}
}

// Get status, manifest, and result first chunk
type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (newState *GetStatementRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatementRequest) {
}

func (newState *GetStatementRequest) SyncEffectiveFieldsDuringRead(existingState GetStatementRequest) {
}

func (a GetStatementRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get result chunk by index
type GetStatementResultChunkNRequest struct {
	ChunkIndex types.Int64 `tfsdk:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-"`
}

func (newState *GetStatementResultChunkNRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetStatementResultChunkNRequest) {
}

func (newState *GetStatementResultChunkNRequest) SyncEffectiveFieldsDuringRead(existingState GetStatementResultChunkNRequest) {
}

func (a GetStatementResultChunkNRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get SQL warehouse permission levels
type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (newState *GetWarehousePermissionLevelsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionLevelsRequest) {
}

func (newState *GetWarehousePermissionLevelsRequest) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionLevelsRequest) {
}

func (a GetWarehousePermissionLevelsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels" tf:"optional"`
}

func (newState *GetWarehousePermissionLevelsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionLevelsResponse) {
}

func (newState *GetWarehousePermissionLevelsResponse) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionLevelsResponse) {
}

func (a GetWarehousePermissionLevelsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"PermissionLevels": reflect.TypeOf(WarehousePermissionsDescription{}),
	}
}

// Get SQL warehouse permissions
type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (newState *GetWarehousePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehousePermissionsRequest) {
}

func (newState *GetWarehousePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState GetWarehousePermissionsRequest) {
}

func (a GetWarehousePermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get warehouse info
type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *GetWarehouseRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehouseRequest) {
}

func (newState *GetWarehouseRequest) SyncEffectiveFieldsDuringRead(existingState GetWarehouseRequest) {
}

func (a GetWarehouseRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins" tf:"optional"`
	// Channel Details
	Channel types.Object `tfsdk:"channel" tf:"optional,object"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size" tf:"optional"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name" tf:"optional"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon" tf:"optional"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute" tf:"optional"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health types.Object `tfsdk:"health" tf:"optional,object"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id" tf:"optional"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url" tf:"optional"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters" tf:"optional"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters" tf:"optional"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions" tf:"optional"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters" tf:"optional"`
	// ODBC parameters for the SQL warehouse
	OdbcParams types.Object `tfsdk:"odbc_params" tf:"optional,object"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy types.String `tfsdk:"spot_instance_policy" tf:"optional"`
	// State of the warehouse
	State types.String `tfsdk:"state" tf:"optional"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags types.Object `tfsdk:"tags" tf:"optional,object"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *GetWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWarehouseResponse) {
}

func (newState *GetWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState GetWarehouseResponse) {
}

func (a GetWarehouseResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Channel":    reflect.TypeOf(Channel{}),
		"Health":     reflect.TypeOf(EndpointHealth{}),
		"OdbcParams": reflect.TypeOf(OdbcParams{}),
		"Tags":       reflect.TypeOf(EndpointTags{}),
	}
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	Channel types.Object `tfsdk:"channel" tf:"optional,object"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.Object `tfsdk:"config_param" tf:"optional,object"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config" tf:"optional"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types" tf:"optional"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam types.Object `tfsdk:"global_param" tf:"optional,object"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account" tf:"optional"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy" tf:"optional"`
	// SQL configuration parameters
	SqlConfigurationParameters types.Object `tfsdk:"sql_configuration_parameters" tf:"optional,object"`
}

func (newState *GetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan GetWorkspaceWarehouseConfigResponse) {
}

func (newState *GetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringRead(existingState GetWorkspaceWarehouseConfigResponse) {
}

func (a GetWorkspaceWarehouseConfigResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Channel":                    reflect.TypeOf(Channel{}),
		"ConfigParam":                reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"DataAccessConfig":           reflect.TypeOf(EndpointConfPair{}),
		"EnabledWarehouseTypes":      reflect.TypeOf(WarehouseTypePair{}),
		"GlobalParam":                reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"SqlConfigurationParameters": reflect.TypeOf(RepeatedEndpointConfPairs{}),
	}
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Alert ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt types.String `tfsdk:"last_triggered_at" tf:"optional"`
	// Name of the alert.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Alert configuration options.
	Options types.Object `tfsdk:"options" tf:"optional,object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`

	Query types.Object `tfsdk:"query" tf:"optional,object"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm" tf:"optional"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State types.String `tfsdk:"state" tf:"optional"`
	// Timestamp when the alert was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`

	User types.Object `tfsdk:"user" tf:"optional,object"`
}

func (newState *LegacyAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyAlert) {
}

func (newState *LegacyAlert) SyncEffectiveFieldsDuringRead(existingState LegacyAlert) {
}

func (a LegacyAlert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options": reflect.TypeOf(AlertOptions{}),
		"Query":   reflect.TypeOf(AlertQuery{}),
		"User":    reflect.TypeOf(User{}),
	}
}

type LegacyQuery struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit types.Bool `tfsdk:"can_edit" tf:"optional"`
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Query ID.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived" tf:"optional"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft" tf:"optional"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite" tf:"optional"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe" tf:"optional"`

	LastModifiedBy types.Object `tfsdk:"last_modified_by" tf:"optional,object"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById types.Int64 `tfsdk:"last_modified_by_id" tf:"optional"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId types.String `tfsdk:"latest_query_data_id" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`

	Options types.Object `tfsdk:"options" tf:"optional,object"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier types.String `tfsdk:"permission_tier" tf:"optional"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash types.String `tfsdk:"query_hash" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`

	User types.Object `tfsdk:"user" tf:"optional,object"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`

	Visualizations types.List `tfsdk:"visualizations" tf:"optional"`
}

func (newState *LegacyQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyQuery) {
}

func (newState *LegacyQuery) SyncEffectiveFieldsDuringRead(existingState LegacyQuery) {
}

func (a LegacyQuery) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"LastModifiedBy": reflect.TypeOf(User{}),
		"Options":        reflect.TypeOf(QueryOptions{}),
		"Tags":           reflect.TypeOf(""),
		"User":           reflect.TypeOf(User{}),
		"Visualizations": reflect.TypeOf(LegacyVisualization{}),
	}
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The UUID for this visualization.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `tfsdk:"options" tf:"optional"`

	Query types.Object `tfsdk:"query" tf:"optional,object"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type types.String `tfsdk:"type" tf:"optional"`

	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`
}

func (newState *LegacyVisualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan LegacyVisualization) {
}

func (newState *LegacyVisualization) SyncEffectiveFieldsDuringRead(existingState LegacyVisualization) {
}

func (a LegacyVisualization) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Query": reflect.TypeOf(LegacyQuery{}),
	}
}

// List alerts
type ListAlertsRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListAlertsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsRequest) {
}

func (newState *ListAlertsRequest) SyncEffectiveFieldsDuringRead(existingState ListAlertsRequest) {
}

func (a ListAlertsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type ListAlertsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListAlertsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsResponse) {
}

func (newState *ListAlertsResponse) SyncEffectiveFieldsDuringRead(existingState ListAlertsResponse) {
}

func (a ListAlertsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Results": reflect.TypeOf(ListAlertsResponseAlert{}),
	}
}

type ListAlertsResponseAlert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition" tf:"optional,object"`
	// The timestamp indicating when the alert was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the alert.
	Id types.String `tfsdk:"id" tf:"optional"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State types.String `tfsdk:"state" tf:"optional"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime types.String `tfsdk:"trigger_time" tf:"optional"`
	// The timestamp indicating when the alert was updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *ListAlertsResponseAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListAlertsResponseAlert) {
}

func (newState *ListAlertsResponseAlert) SyncEffectiveFieldsDuringRead(existingState ListAlertsResponseAlert) {
}

func (a ListAlertsResponseAlert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Condition": reflect.TypeOf(AlertCondition{}),
	}
}

// Get dashboard objects
type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order types.String `tfsdk:"-"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-"`
	// Number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// Full text search term.
	Q types.String `tfsdk:"-"`
}

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListDashboardsRequest) {
}

func (newState *ListDashboardsRequest) SyncEffectiveFieldsDuringRead(existingState ListDashboardsRequest) {
}

func (a ListDashboardsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Get a list of queries
type ListQueriesLegacyRequest struct {
	// Name of query attribute to order by. Default sort order is ascending.
	// Append a dash (`-`) to order descending instead.
	//
	// - `name`: The name of the query.
	//
	// - `created_at`: The timestamp the query was created.
	//
	// - `runtime`: The time it took to run this query. This is blank for
	// parameterized queries. A blank value is treated as the highest value for
	// sorting.
	//
	// - `executed_at`: The timestamp when the query was last run.
	//
	// - `created_by`: The user name of the user that created the query.
	Order types.String `tfsdk:"-"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-"`
	// Number of queries to return per page.
	PageSize types.Int64 `tfsdk:"-"`
	// Full text search term
	Q types.String `tfsdk:"-"`
}

func (newState *ListQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesLegacyRequest) {
}

func (newState *ListQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState ListQueriesLegacyRequest) {
}

func (a ListQueriesLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// List queries
type ListQueriesRequest struct {
	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQueriesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesRequest) {
}

func (newState *ListQueriesRequest) SyncEffectiveFieldsDuringRead(existingState ListQueriesRequest) {
}

func (a ListQueriesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	HasNextPage types.Bool `tfsdk:"has_next_page" tf:"optional"`
	// A token that can be used to get the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Res types.List `tfsdk:"res" tf:"optional"`
}

func (newState *ListQueriesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueriesResponse) {
}

func (newState *ListQueriesResponse) SyncEffectiveFieldsDuringRead(existingState ListQueriesResponse) {
}

func (a ListQueriesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Res": reflect.TypeOf(QueryInfo{}),
	}
}

// List Queries
type ListQueryHistoryRequest struct {
	// A filter to limit query history results. This field is optional.
	FilterBy types.Object `tfsdk:"-"`
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	IncludeMetrics types.Bool `tfsdk:"-"`
	// Limit the number of results returned in one page. Must be less than 1000
	// and the default is 100.
	MaxResults types.Int64 `tfsdk:"-"`
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	PageToken types.String `tfsdk:"-"`
}

func (newState *ListQueryHistoryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryHistoryRequest) {
}

func (newState *ListQueryHistoryRequest) SyncEffectiveFieldsDuringRead(existingState ListQueryHistoryRequest) {
}

func (a ListQueryHistoryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"FilterBy": reflect.TypeOf(QueryFilter{}),
	}
}

type ListQueryObjectsResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListQueryObjectsResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryObjectsResponse) {
}

func (newState *ListQueryObjectsResponse) SyncEffectiveFieldsDuringRead(existingState ListQueryObjectsResponse) {
}

func (a ListQueryObjectsResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Results": reflect.TypeOf(ListQueryObjectsResponseQuery{}),
	}
}

type ListQueryObjectsResponseQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// Timestamp when this query was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the query.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName types.String `tfsdk:"last_modifier_user_name" tf:"optional"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// Timestamp when this query was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *ListQueryObjectsResponseQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListQueryObjectsResponseQuery) {
}

func (newState *ListQueryObjectsResponseQuery) SyncEffectiveFieldsDuringRead(existingState ListQueryObjectsResponseQuery) {
}

func (a ListQueryObjectsResponseQuery) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(QueryParameter{}),
		"Tags":       reflect.TypeOf(""),
	}
}

type ListResponse struct {
	// The total number of dashboards.
	Count types.Int64 `tfsdk:"count" tf:"optional"`
	// The current page being displayed.
	Page types.Int64 `tfsdk:"page" tf:"optional"`
	// The number of dashboards per page.
	PageSize types.Int64 `tfsdk:"page_size" tf:"optional"`
	// List of dashboards returned.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListResponse) {
}

func (newState *ListResponse) SyncEffectiveFieldsDuringRead(existingState ListResponse) {
}

func (a ListResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Results": reflect.TypeOf(Dashboard{}),
	}
}

// List visualizations on a query
type ListVisualizationsForQueryRequest struct {
	Id types.String `tfsdk:"-"`

	PageSize types.Int64 `tfsdk:"-"`

	PageToken types.String `tfsdk:"-"`
}

func (newState *ListVisualizationsForQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVisualizationsForQueryRequest) {
}

func (newState *ListVisualizationsForQueryRequest) SyncEffectiveFieldsDuringRead(existingState ListVisualizationsForQueryRequest) {
}

func (a ListVisualizationsForQueryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type ListVisualizationsForQueryResponse struct {
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`

	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *ListVisualizationsForQueryResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListVisualizationsForQueryResponse) {
}

func (newState *ListVisualizationsForQueryResponse) SyncEffectiveFieldsDuringRead(existingState ListVisualizationsForQueryResponse) {
}

func (a ListVisualizationsForQueryResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Results": reflect.TypeOf(Visualization{}),
	}
}

// List warehouses
type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId types.Int64 `tfsdk:"-"`
}

func (newState *ListWarehousesRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWarehousesRequest) {
}

func (newState *ListWarehousesRequest) SyncEffectiveFieldsDuringRead(existingState ListWarehousesRequest) {
}

func (a ListWarehousesRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	Warehouses types.List `tfsdk:"warehouses" tf:"optional"`
}

func (newState *ListWarehousesResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan ListWarehousesResponse) {
}

func (newState *ListWarehousesResponse) SyncEffectiveFieldsDuringRead(existingState ListWarehousesResponse) {
}

func (a ListWarehousesResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Warehouses": reflect.TypeOf(EndpointInfo{}),
	}
}

type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	Prefix types.String `tfsdk:"prefix" tf:"optional"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator types.String `tfsdk:"separator" tf:"optional"`
	// Character that suffixes each selected parameter value.
	Suffix types.String `tfsdk:"suffix" tf:"optional"`
}

func (newState *MultiValuesOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan MultiValuesOptions) {
}

func (newState *MultiValuesOptions) SyncEffectiveFieldsDuringRead(existingState MultiValuesOptions) {
}

func (a MultiValuesOptions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type NumericValue struct {
	Value types.Float64 `tfsdk:"value" tf:"optional"`
}

func (newState *NumericValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan NumericValue) {
}

func (newState *NumericValue) SyncEffectiveFieldsDuringRead(existingState NumericValue) {
}

func (a NumericValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type OdbcParams struct {
	Hostname types.String `tfsdk:"hostname" tf:"optional"`

	Path types.String `tfsdk:"path" tf:"optional"`

	Port types.Int64 `tfsdk:"port" tf:"optional"`

	Protocol types.String `tfsdk:"protocol" tf:"optional"`
}

func (newState *OdbcParams) SyncEffectiveFieldsDuringCreateOrUpdate(plan OdbcParams) {
}

func (newState *OdbcParams) SyncEffectiveFieldsDuringRead(existingState OdbcParams) {
}

func (a OdbcParams) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type Parameter struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	EnumOptions types.String `tfsdk:"enumOptions" tf:"optional"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions types.Object `tfsdk:"multiValuesOptions" tf:"optional,object"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name types.String `tfsdk:"name" tf:"optional"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId types.String `tfsdk:"queryId" tf:"optional"`
	// The text displayed in a parameter picking widget.
	Title types.String `tfsdk:"title" tf:"optional"`
	// Parameters can have several different types.
	Type types.String `tfsdk:"type" tf:"optional"`
	// The default value for this parameter.
	Value any `tfsdk:"value" tf:"optional"`
}

func (newState *Parameter) SyncEffectiveFieldsDuringCreateOrUpdate(plan Parameter) {
}

func (newState *Parameter) SyncEffectiveFieldsDuringRead(existingState Parameter) {
}

func (a Parameter) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"MultiValuesOptions": reflect.TypeOf(MultiValuesOptions{}),
	}
}

type Query struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// Timestamp when this query was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the query.
	Id types.String `tfsdk:"id" tf:"optional"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName types.String `tfsdk:"last_modifier_user_name" tf:"optional"`
	// Indicates whether the query is trashed.
	LifecycleState types.String `tfsdk:"lifecycle_state" tf:"optional"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Workspace path of the workspace folder containing the object.
	ParentPath types.String `tfsdk:"parent_path" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// Timestamp when this query was last updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *Query) SyncEffectiveFieldsDuringCreateOrUpdate(plan Query) {
}

func (newState *Query) SyncEffectiveFieldsDuringRead(existingState Query) {
}

func (a Query) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(QueryParameter{}),
		"Tags":       reflect.TypeOf(""),
	}
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions types.Object `tfsdk:"multi_values_options" tf:"optional,object"`
	// UUID of the query that provides the parameter values.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// List of selected query parameter values.
	Values types.List `tfsdk:"values" tf:"optional"`
}

func (newState *QueryBackedValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryBackedValue) {
}

func (newState *QueryBackedValue) SyncEffectiveFieldsDuringRead(existingState QueryBackedValue) {
}

func (a QueryBackedValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"MultiValuesOptions": reflect.TypeOf(MultiValuesOptions{}),
		"Values":             reflect.TypeOf(""),
	}
}

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `tfsdk:"options" tf:"optional"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`

	QueryId types.String `tfsdk:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *QueryEditContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryEditContent) {
}

func (newState *QueryEditContent) SyncEffectiveFieldsDuringRead(existingState QueryEditContent) {
}

func (a QueryEditContent) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(""),
	}
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be <= 30
	// days.
	QueryStartTimeRange types.Object `tfsdk:"query_start_time_range" tf:"optional,object"`
	// A list of statement IDs.
	StatementIds types.List `tfsdk:"statement_ids" tf:"optional"`

	Statuses types.List `tfsdk:"statuses" tf:"optional"`
	// A list of user IDs who ran the queries.
	UserIds types.List `tfsdk:"user_ids" tf:"optional"`
	// A list of warehouse IDs.
	WarehouseIds types.List `tfsdk:"warehouse_ids" tf:"optional"`
}

func (newState *QueryFilter) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryFilter) {
}

func (newState *QueryFilter) SyncEffectiveFieldsDuringRead(existingState QueryFilter) {
}

func (a QueryFilter) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"QueryStartTimeRange": reflect.TypeOf(TimeRange{}),
		"StatementIds":        reflect.TypeOf(""),
		"Statuses":            reflect.TypeOf(""),
		"UserIds":             reflect.TypeOf(0),
		"WarehouseIds":        reflect.TypeOf(""),
	}
}

type QueryInfo struct {
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed types.Object `tfsdk:"channel_used" tf:"optional,object"`
	// Total execution time of the statement ( excluding result fetch time ).
	Duration types.Int64 `tfsdk:"duration" tf:"optional"`
	// Alias for `warehouse_id`.
	EndpointId types.String `tfsdk:"endpoint_id" tf:"optional"`
	// Message describing why the query could not complete.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId types.Int64 `tfsdk:"executed_as_user_id" tf:"optional"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName types.String `tfsdk:"executed_as_user_name" tf:"optional"`
	// The time execution of the query ended.
	ExecutionEndTimeMs types.Int64 `tfsdk:"execution_end_time_ms" tf:"optional"`
	// Whether more updates for the query are expected.
	IsFinal types.Bool `tfsdk:"is_final" tf:"optional"`
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key" tf:"optional"`
	// Metrics about query execution.
	Metrics types.Object `tfsdk:"metrics" tf:"optional,object"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState types.String `tfsdk:"plans_state" tf:"optional"`
	// The time the query ended.
	QueryEndTimeMs types.Int64 `tfsdk:"query_end_time_ms" tf:"optional"`
	// The query ID.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// The time the query started.
	QueryStartTimeMs types.Int64 `tfsdk:"query_start_time_ms" tf:"optional"`
	// The text of the query.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// The number of results returned by the query.
	RowsProduced types.Int64 `tfsdk:"rows_produced" tf:"optional"`
	// URL to the Spark UI query plan.
	SparkUiUrl types.String `tfsdk:"spark_ui_url" tf:"optional"`
	// Type of statement for this query
	StatementType types.String `tfsdk:"statement_type" tf:"optional"`
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	Status types.String `tfsdk:"status" tf:"optional"`
	// The ID of the user who ran the query.
	UserId types.Int64 `tfsdk:"user_id" tf:"optional"`
	// The email address or username of the user who ran the query.
	UserName types.String `tfsdk:"user_name" tf:"optional"`
	// Warehouse ID.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *QueryInfo) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryInfo) {
}

func (newState *QueryInfo) SyncEffectiveFieldsDuringRead(existingState QueryInfo) {
}

func (a QueryInfo) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ChannelUsed": reflect.TypeOf(ChannelInfo{}),
		"Metrics":     reflect.TypeOf(QueryMetrics{}),
	}
}

type QueryList struct {
	// The total number of queries.
	Count types.Int64 `tfsdk:"count" tf:"optional"`
	// The page number that is currently displayed.
	Page types.Int64 `tfsdk:"page" tf:"optional"`
	// The number of queries per page.
	PageSize types.Int64 `tfsdk:"page_size" tf:"optional"`
	// List of queries returned.
	Results types.List `tfsdk:"results" tf:"optional"`
}

func (newState *QueryList) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryList) {
}

func (newState *QueryList) SyncEffectiveFieldsDuringRead(existingState QueryList) {
}

func (a QueryList) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Results": reflect.TypeOf(LegacyQuery{}),
	}
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs types.Int64 `tfsdk:"compilation_time_ms" tf:"optional"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs types.Int64 `tfsdk:"execution_time_ms" tf:"optional"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes types.Int64 `tfsdk:"network_sent_bytes" tf:"optional"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	OverloadingQueueStartTimestamp types.Int64 `tfsdk:"overloading_queue_start_timestamp" tf:"optional"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs types.Int64 `tfsdk:"photon_total_time_ms" tf:"optional"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	ProvisioningQueueStartTimestamp types.Int64 `tfsdk:"provisioning_queue_start_timestamp" tf:"optional"`
	// Total number of bytes in all tables not read due to pruning
	PrunedBytes types.Int64 `tfsdk:"pruned_bytes" tf:"optional"`
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount types.Int64 `tfsdk:"pruned_files_count" tf:"optional"`
	// Timestamp of when the underlying compute started compilation of the
	// query.
	QueryCompilationStartTimestamp types.Int64 `tfsdk:"query_compilation_start_timestamp" tf:"optional"`
	// Total size of data read by the query, in bytes.
	ReadBytes types.Int64 `tfsdk:"read_bytes" tf:"optional"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes types.Int64 `tfsdk:"read_cache_bytes" tf:"optional"`
	// Number of files read after pruning
	ReadFilesCount types.Int64 `tfsdk:"read_files_count" tf:"optional"`
	// Number of partitions read after pruning.
	ReadPartitionsCount types.Int64 `tfsdk:"read_partitions_count" tf:"optional"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes types.Int64 `tfsdk:"read_remote_bytes" tf:"optional"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs types.Int64 `tfsdk:"result_fetch_time_ms" tf:"optional"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache types.Bool `tfsdk:"result_from_cache" tf:"optional"`
	// Total number of rows returned by the query.
	RowsProducedCount types.Int64 `tfsdk:"rows_produced_count" tf:"optional"`
	// Total number of rows read by the query.
	RowsReadCount types.Int64 `tfsdk:"rows_read_count" tf:"optional"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes types.Int64 `tfsdk:"spill_to_disk_bytes" tf:"optional"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs types.Int64 `tfsdk:"task_total_time_ms" tf:"optional"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs types.Int64 `tfsdk:"total_time_ms" tf:"optional"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes types.Int64 `tfsdk:"write_remote_bytes" tf:"optional"`
}

func (newState *QueryMetrics) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryMetrics) {
}

func (newState *QueryMetrics) SyncEffectiveFieldsDuringRead(existingState QueryMetrics) {
}

func (a QueryMetrics) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at" tf:"optional"`

	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// The name of the schema to execute this query in.
	Schema types.String `tfsdk:"schema" tf:"optional"`
}

func (newState *QueryOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryOptions) {
}

func (newState *QueryOptions) SyncEffectiveFieldsDuringRead(existingState QueryOptions) {
}

func (a QueryOptions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(Parameter{}),
	}
}

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	DateRangeValue types.Object `tfsdk:"date_range_value" tf:"optional,object"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue types.Object `tfsdk:"date_value" tf:"optional,object"`
	// Dropdown query parameter value.
	EnumValue types.Object `tfsdk:"enum_value" tf:"optional,object"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Numeric query parameter value.
	NumericValue types.Object `tfsdk:"numeric_value" tf:"optional,object"`
	// Query-based dropdown query parameter value.
	QueryBackedValue types.Object `tfsdk:"query_backed_value" tf:"optional,object"`
	// Text query parameter value.
	TextValue types.Object `tfsdk:"text_value" tf:"optional,object"`
	// Text displayed in the user-facing parameter widget in the UI.
	Title types.String `tfsdk:"title" tf:"optional"`
}

func (newState *QueryParameter) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryParameter) {
}

func (newState *QueryParameter) SyncEffectiveFieldsDuringRead(existingState QueryParameter) {
}

func (a QueryParameter) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DateRangeValue":   reflect.TypeOf(DateRangeValue{}),
		"DateValue":        reflect.TypeOf(DateValue{}),
		"EnumValue":        reflect.TypeOf(EnumValue{}),
		"NumericValue":     reflect.TypeOf(NumericValue{}),
		"QueryBackedValue": reflect.TypeOf(QueryBackedValue{}),
		"TextValue":        reflect.TypeOf(TextValue{}),
	}
}

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `tfsdk:"options" tf:"optional"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent" tf:"optional"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query" tf:"optional"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole types.String `tfsdk:"run_as_role" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
}

func (newState *QueryPostContent) SyncEffectiveFieldsDuringCreateOrUpdate(plan QueryPostContent) {
}

func (newState *QueryPostContent) SyncEffectiveFieldsDuringRead(existingState QueryPostContent) {
}

func (a QueryPostContent) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Tags": reflect.TypeOf(""),
	}
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	ConfigPair types.List `tfsdk:"config_pair" tf:"optional"`

	ConfigurationPairs types.List `tfsdk:"configuration_pairs" tf:"optional"`
}

func (newState *RepeatedEndpointConfPairs) SyncEffectiveFieldsDuringCreateOrUpdate(plan RepeatedEndpointConfPairs) {
}

func (newState *RepeatedEndpointConfPairs) SyncEffectiveFieldsDuringRead(existingState RepeatedEndpointConfPairs) {
}

func (a RepeatedEndpointConfPairs) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ConfigPair":         reflect.TypeOf(EndpointConfPair{}),
		"ConfigurationPairs": reflect.TypeOf(EndpointConfPair{}),
	}
}

// Restore a dashboard
type RestoreDashboardRequest struct {
	DashboardId types.String `tfsdk:"-"`
}

func (newState *RestoreDashboardRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreDashboardRequest) {
}

func (newState *RestoreDashboardRequest) SyncEffectiveFieldsDuringRead(existingState RestoreDashboardRequest) {
}

func (a RestoreDashboardRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Restore a query
type RestoreQueriesLegacyRequest struct {
	QueryId types.String `tfsdk:"-"`
}

func (newState *RestoreQueriesLegacyRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreQueriesLegacyRequest) {
}

func (newState *RestoreQueriesLegacyRequest) SyncEffectiveFieldsDuringRead(existingState RestoreQueriesLegacyRequest) {
}

func (a RestoreQueriesLegacyRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type RestoreResponse struct {
}

func (newState *RestoreResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan RestoreResponse) {
}

func (newState *RestoreResponse) SyncEffectiveFieldsDuringRead(existingState RestoreResponse) {
}

func (a RestoreResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count" tf:"optional"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index" tf:"optional"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray types.List `tfsdk:"data_array" tf:"optional"`

	ExternalLinks types.List `tfsdk:"external_links" tf:"optional"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index" tf:"optional"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link" tf:"optional"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count" tf:"optional"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset" tf:"optional"`
}

func (newState *ResultData) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultData) {
}

func (newState *ResultData) SyncEffectiveFieldsDuringRead(existingState ResultData) {
}

func (a ResultData) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"DataArray":     reflect.TypeOf(""),
		"ExternalLinks": reflect.TypeOf(ExternalLink{}),
	}
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks" tf:"optional"`

	Format types.String `tfsdk:"format" tf:"optional"`
	// The schema is an ordered list of column descriptions.
	Schema types.Object `tfsdk:"schema" tf:"optional,object"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount types.Int64 `tfsdk:"total_byte_count" tf:"optional"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount types.Int64 `tfsdk:"total_chunk_count" tf:"optional"`
	// The total number of rows in the result set.
	TotalRowCount types.Int64 `tfsdk:"total_row_count" tf:"optional"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated types.Bool `tfsdk:"truncated" tf:"optional"`
}

func (newState *ResultManifest) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultManifest) {
}

func (newState *ResultManifest) SyncEffectiveFieldsDuringRead(existingState ResultManifest) {
}

func (a ResultManifest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Chunks": reflect.TypeOf(BaseChunkInfo{}),
		"Schema": reflect.TypeOf(ResultSchema{}),
	}
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {
	ColumnCount types.Int64 `tfsdk:"column_count" tf:"optional"`

	Columns types.List `tfsdk:"columns" tf:"optional"`
}

func (newState *ResultSchema) SyncEffectiveFieldsDuringCreateOrUpdate(plan ResultSchema) {
}

func (newState *ResultSchema) SyncEffectiveFieldsDuringRead(existingState ResultSchema) {
}

func (a ResultSchema) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Columns": reflect.TypeOf(ColumnInfo{}),
	}
}

type ServiceError struct {
	ErrorCode types.String `tfsdk:"error_code" tf:"optional"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *ServiceError) SyncEffectiveFieldsDuringCreateOrUpdate(plan ServiceError) {
}

func (newState *ServiceError) SyncEffectiveFieldsDuringRead(existingState ServiceError) {
}

func (a ServiceError) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Set object ACL
type SetRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId types.String `tfsdk:"-"`
	// The type of object permission to set.
	ObjectType types.String `tfsdk:"-"`
}

func (newState *SetRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetRequest) {
}

func (newState *SetRequest) SyncEffectiveFieldsDuringRead(existingState SetRequest) {
}

func (a SetRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(AccessControl{}),
	}
}

type SetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id" tf:"optional"`
	// A singular noun object type.
	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *SetResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetResponse) {
}

func (newState *SetResponse) SyncEffectiveFieldsDuringRead(existingState SetResponse) {
}

func (a SetResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(AccessControl{}),
	}
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel types.Object `tfsdk:"channel" tf:"optional,object"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam types.Object `tfsdk:"config_param" tf:"optional,object"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config" tf:"optional"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types" tf:"optional"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam types.Object `tfsdk:"global_param" tf:"optional,object"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account" tf:"optional"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// Security policy for warehouses
	SecurityPolicy types.String `tfsdk:"security_policy" tf:"optional"`
	// SQL configuration parameters
	SqlConfigurationParameters types.Object `tfsdk:"sql_configuration_parameters" tf:"optional,object"`
}

func (newState *SetWorkspaceWarehouseConfigRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetWorkspaceWarehouseConfigRequest) {
}

func (newState *SetWorkspaceWarehouseConfigRequest) SyncEffectiveFieldsDuringRead(existingState SetWorkspaceWarehouseConfigRequest) {
}

func (a SetWorkspaceWarehouseConfigRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Channel":                    reflect.TypeOf(Channel{}),
		"ConfigParam":                reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"DataAccessConfig":           reflect.TypeOf(EndpointConfPair{}),
		"EnabledWarehouseTypes":      reflect.TypeOf(WarehouseTypePair{}),
		"GlobalParam":                reflect.TypeOf(RepeatedEndpointConfPairs{}),
		"SqlConfigurationParameters": reflect.TypeOf(RepeatedEndpointConfPairs{}),
	}
}

type SetWorkspaceWarehouseConfigResponse struct {
}

func (newState *SetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan SetWorkspaceWarehouseConfigResponse) {
}

func (newState *SetWorkspaceWarehouseConfigResponse) SyncEffectiveFieldsDuringRead(existingState SetWorkspaceWarehouseConfigResponse) {
}

func (a SetWorkspaceWarehouseConfigResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Start a warehouse
type StartRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *StartRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartRequest) {
}

func (newState *StartRequest) SyncEffectiveFieldsDuringRead(existingState StartRequest) {
}

func (a StartRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type StartWarehouseResponse struct {
}

func (newState *StartWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StartWarehouseResponse) {
}

func (newState *StartWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState StartWarehouseResponse) {
}

func (a StartWarehouseResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	Name types.String `tfsdk:"name" tf:""`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type types.String `tfsdk:"type" tf:"optional"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *StatementParameterListItem) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementParameterListItem) {
}

func (newState *StatementParameterListItem) SyncEffectiveFieldsDuringRead(existingState StatementParameterListItem) {
}

func (a StatementParameterListItem) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type StatementResponse struct {
	// The result manifest provides schema and metadata for the result set.
	Manifest types.Object `tfsdk:"manifest" tf:"optional,object"`

	Result types.Object `tfsdk:"result" tf:"optional,object"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id" tf:"optional"`
	// The status response includes execution state and if relevant, error
	// information.
	Status types.Object `tfsdk:"status" tf:"optional,object"`
}

func (newState *StatementResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementResponse) {
}

func (newState *StatementResponse) SyncEffectiveFieldsDuringRead(existingState StatementResponse) {
}

func (a StatementResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Manifest": reflect.TypeOf(ResultManifest{}),
		"Result":   reflect.TypeOf(ResultData{}),
		"Status":   reflect.TypeOf(StatementStatus{}),
	}
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	Error types.Object `tfsdk:"error" tf:"optional,object"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accomanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State types.String `tfsdk:"state" tf:"optional"`
}

func (newState *StatementStatus) SyncEffectiveFieldsDuringCreateOrUpdate(plan StatementStatus) {
}

func (newState *StatementStatus) SyncEffectiveFieldsDuringRead(existingState StatementStatus) {
}

func (a StatementStatus) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Error": reflect.TypeOf(ServiceError{}),
	}
}

// Stop a warehouse
type StopRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-"`
}

func (newState *StopRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopRequest) {
}

func (newState *StopRequest) SyncEffectiveFieldsDuringRead(existingState StopRequest) {
}

func (a StopRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type StopWarehouseResponse struct {
}

func (newState *StopWarehouseResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan StopWarehouseResponse) {
}

func (newState *StopWarehouseResponse) SyncEffectiveFieldsDuringRead(existingState StopWarehouseResponse) {
}

func (a StopWarehouseResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type Success struct {
	Message types.String `tfsdk:"message" tf:"optional"`
}

func (newState *Success) SyncEffectiveFieldsDuringCreateOrUpdate(plan Success) {
}

func (newState *Success) SyncEffectiveFieldsDuringRead(existingState Success) {
}

func (a Success) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code types.String `tfsdk:"code" tf:"optional"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters" tf:"optional"`
	// type of the termination
	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringCreateOrUpdate(plan TerminationReason) {
}

func (newState *TerminationReason) SyncEffectiveFieldsDuringRead(existingState TerminationReason) {
}

func (a TerminationReason) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(""),
	}
}

type TextValue struct {
	Value types.String `tfsdk:"value" tf:"optional"`
}

func (newState *TextValue) SyncEffectiveFieldsDuringCreateOrUpdate(plan TextValue) {
}

func (newState *TextValue) SyncEffectiveFieldsDuringRead(existingState TextValue) {
}

func (a TextValue) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type TimeRange struct {
	// The end time in milliseconds.
	EndTimeMs types.Int64 `tfsdk:"end_time_ms" tf:"optional"`
	// The start time in milliseconds.
	StartTimeMs types.Int64 `tfsdk:"start_time_ms" tf:"optional"`
}

func (newState *TimeRange) SyncEffectiveFieldsDuringCreateOrUpdate(plan TimeRange) {
}

func (newState *TimeRange) SyncEffectiveFieldsDuringRead(existingState TimeRange) {
}

func (a TimeRange) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner" tf:"optional"`
}

func (newState *TransferOwnershipObjectId) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransferOwnershipObjectId) {
}

func (newState *TransferOwnershipObjectId) SyncEffectiveFieldsDuringRead(existingState TransferOwnershipObjectId) {
}

func (a TransferOwnershipObjectId) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Transfer object ownership
type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner" tf:"optional"`
	// The ID of the object on which to change ownership.
	ObjectId types.Object `tfsdk:"-"`
	// The type of object on which to change ownership.
	ObjectType types.String `tfsdk:"-"`
}

func (newState *TransferOwnershipRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TransferOwnershipRequest) {
}

func (newState *TransferOwnershipRequest) SyncEffectiveFieldsDuringRead(existingState TransferOwnershipRequest) {
}

func (a TransferOwnershipRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"ObjectId": reflect.TypeOf(TransferOwnershipObjectId{}),
	}
}

// Delete an alert
type TrashAlertRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *TrashAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashAlertRequest) {
}

func (newState *TrashAlertRequest) SyncEffectiveFieldsDuringRead(existingState TrashAlertRequest) {
}

func (a TrashAlertRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Delete a query
type TrashQueryRequest struct {
	Id types.String `tfsdk:"-"`
}

func (newState *TrashQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan TrashQueryRequest) {
}

func (newState *TrashQueryRequest) SyncEffectiveFieldsDuringRead(existingState TrashQueryRequest) {
}

func (a TrashQueryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type UpdateAlertRequest struct {
	Alert types.Object `tfsdk:"alert" tf:"optional,object"`

	Id types.String `tfsdk:"-"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"update_mask" tf:""`
}

func (newState *UpdateAlertRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAlertRequest) {
}

func (newState *UpdateAlertRequest) SyncEffectiveFieldsDuringRead(existingState UpdateAlertRequest) {
}

func (a UpdateAlertRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Alert": reflect.TypeOf(UpdateAlertRequestAlert{}),
	}
}

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition types.Object `tfsdk:"condition" tf:"optional,object"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body" tf:"optional"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject" tf:"optional"`
	// The display name of the alert.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk types.Bool `tfsdk:"notify_on_ok" tf:"optional"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// UUID of the query attached to the alert.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger types.Int64 `tfsdk:"seconds_to_retrigger" tf:"optional"`
}

func (newState *UpdateAlertRequestAlert) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateAlertRequestAlert) {
}

func (newState *UpdateAlertRequestAlert) SyncEffectiveFieldsDuringRead(existingState UpdateAlertRequestAlert) {
}

func (a UpdateAlertRequestAlert) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Condition": reflect.TypeOf(AlertCondition{}),
	}
}

type UpdateQueryRequest struct {
	Id types.String `tfsdk:"-"`

	Query types.Object `tfsdk:"query" tf:"optional,object"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"update_mask" tf:""`
}

func (newState *UpdateQueryRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateQueryRequest) {
}

func (newState *UpdateQueryRequest) SyncEffectiveFieldsDuringRead(existingState UpdateQueryRequest) {
}

func (a UpdateQueryRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Query": reflect.TypeOf(UpdateQueryRequestQuery{}),
	}
}

type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit types.Bool `tfsdk:"apply_auto_limit" tf:"optional"`
	// Name of the catalog where this query will be executed.
	Catalog types.String `tfsdk:"catalog" tf:"optional"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description" tf:"optional"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// Username of the user that owns the query.
	OwnerUserName types.String `tfsdk:"owner_user_name" tf:"optional"`
	// List of query parameter definitions.
	Parameters types.List `tfsdk:"parameters" tf:"optional"`
	// Text of the query to be run.
	QueryText types.String `tfsdk:"query_text" tf:"optional"`
	// Sets the "Run as" role for the object.
	RunAsMode types.String `tfsdk:"run_as_mode" tf:"optional"`
	// Name of the schema where this query will be executed.
	Schema types.String `tfsdk:"schema" tf:"optional"`

	Tags types.List `tfsdk:"tags" tf:"optional"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId types.String `tfsdk:"warehouse_id" tf:"optional"`
}

func (newState *UpdateQueryRequestQuery) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateQueryRequestQuery) {
}

func (newState *UpdateQueryRequestQuery) SyncEffectiveFieldsDuringRead(existingState UpdateQueryRequestQuery) {
}

func (a UpdateQueryRequestQuery) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Parameters": reflect.TypeOf(QueryParameter{}),
		"Tags":       reflect.TypeOf(""),
	}
}

type UpdateResponse struct {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateResponse) {
}

func (newState *UpdateResponse) SyncEffectiveFieldsDuringRead(existingState UpdateResponse) {
}

func (a UpdateResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type UpdateVisualizationRequest struct {
	Id types.String `tfsdk:"-"`
	// Field mask is required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. The field
	// mask needs to be supplied as single string. To specify multiple fields in
	// the field mask, use comma as the separator (no space).
	UpdateMask types.String `tfsdk:"update_mask" tf:""`

	Visualization types.Object `tfsdk:"visualization" tf:"optional,object"`
}

func (newState *UpdateVisualizationRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVisualizationRequest) {
}

func (newState *UpdateVisualizationRequest) SyncEffectiveFieldsDuringRead(existingState UpdateVisualizationRequest) {
}

func (a UpdateVisualizationRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Visualization": reflect.TypeOf(UpdateVisualizationRequestVisualization{}),
	}
}

type UpdateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options" tf:"optional"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan" tf:"optional"`
	// The type of visualization: counter, table, funnel, and so on.
	Type types.String `tfsdk:"type" tf:"optional"`
}

func (newState *UpdateVisualizationRequestVisualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan UpdateVisualizationRequestVisualization) {
}

func (newState *UpdateVisualizationRequestVisualization) SyncEffectiveFieldsDuringRead(existingState UpdateVisualizationRequestVisualization) {
}

func (a UpdateVisualizationRequestVisualization) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type User struct {
	Email types.String `tfsdk:"email" tf:"optional"`

	Id types.Int64 `tfsdk:"id" tf:"optional"`

	Name types.String `tfsdk:"name" tf:"optional"`
}

func (newState *User) SyncEffectiveFieldsDuringCreateOrUpdate(plan User) {
}

func (newState *User) SyncEffectiveFieldsDuringRead(existingState User) {
}

func (a User) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	CreateTime types.String `tfsdk:"create_time" tf:"optional"`
	// The display name of the visualization.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// UUID identifying the visualization.
	Id types.String `tfsdk:"id" tf:"optional"`
	// UUID of the query that the visualization is attached to.
	QueryId types.String `tfsdk:"query_id" tf:"optional"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions types.String `tfsdk:"serialized_options" tf:"optional"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan types.String `tfsdk:"serialized_query_plan" tf:"optional"`
	// The type of visualization: counter, table, funnel, and so on.
	Type types.String `tfsdk:"type" tf:"optional"`
	// The timestamp indicating when the visualization was updated.
	UpdateTime types.String `tfsdk:"update_time" tf:"optional"`
}

func (newState *Visualization) SyncEffectiveFieldsDuringCreateOrUpdate(plan Visualization) {
}

func (newState *Visualization) SyncEffectiveFieldsDuringRead(existingState Visualization) {
}

func (a Visualization) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type WarehouseAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *WarehouseAccessControlRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseAccessControlRequest) {
}

func (newState *WarehouseAccessControlRequest) SyncEffectiveFieldsDuringRead(existingState WarehouseAccessControlRequest) {
}

func (a WarehouseAccessControlRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

func (newState *WarehouseAccessControlResponse) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseAccessControlResponse) {
}

func (newState *WarehouseAccessControlResponse) SyncEffectiveFieldsDuringRead(existingState WarehouseAccessControlResponse) {
}

func (a WarehouseAccessControlResponse) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AllPermissions": reflect.TypeOf(WarehousePermission{}),
	}
}

type WarehousePermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *WarehousePermission) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermission) {
}

func (newState *WarehousePermission) SyncEffectiveFieldsDuringRead(existingState WarehousePermission) {
}

func (a WarehousePermission) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"InheritedFromObject": reflect.TypeOf(""),
	}
}

type WarehousePermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

func (newState *WarehousePermissions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissions) {
}

func (newState *WarehousePermissions) SyncEffectiveFieldsDuringRead(existingState WarehousePermissions) {
}

func (a WarehousePermissions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(WarehouseAccessControlResponse{}),
	}
}

type WarehousePermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level" tf:"optional"`
}

func (newState *WarehousePermissionsDescription) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissionsDescription) {
}

func (newState *WarehousePermissionsDescription) SyncEffectiveFieldsDuringRead(existingState WarehousePermissionsDescription) {
}

func (a WarehousePermissionsDescription) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type WarehousePermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list" tf:"optional"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-"`
}

func (newState *WarehousePermissionsRequest) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehousePermissionsRequest) {
}

func (newState *WarehousePermissionsRequest) SyncEffectiveFieldsDuringRead(existingState WarehousePermissionsRequest) {
}

func (a WarehousePermissionsRequest) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"AccessControlList": reflect.TypeOf(WarehouseAccessControlRequest{}),
	}
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType types.String `tfsdk:"warehouse_type" tf:"optional"`
}

func (newState *WarehouseTypePair) SyncEffectiveFieldsDuringCreateOrUpdate(plan WarehouseTypePair) {
}

func (newState *WarehouseTypePair) SyncEffectiveFieldsDuringRead(existingState WarehouseTypePair) {
}

func (a WarehouseTypePair) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type Widget struct {
	// The unique ID for this widget.
	Id types.String `tfsdk:"id" tf:"optional"`

	Options types.Object `tfsdk:"options" tf:"optional,object"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization types.Object `tfsdk:"visualization" tf:"optional,object"`
	// Unused field.
	Width types.Int64 `tfsdk:"width" tf:"optional"`
}

func (newState *Widget) SyncEffectiveFieldsDuringCreateOrUpdate(plan Widget) {
}

func (newState *Widget) SyncEffectiveFieldsDuringRead(existingState Widget) {
}

func (a Widget) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Options":       reflect.TypeOf(WidgetOptions{}),
		"Visualization": reflect.TypeOf(LegacyVisualization{}),
	}
}

type WidgetOptions struct {
	// Timestamp when this object was created
	CreatedAt types.String `tfsdk:"created_at" tf:"optional"`
	// Custom description of the widget
	Description types.String `tfsdk:"description" tf:"optional"`
	// Whether this widget is hidden on the dashboard.
	IsHidden types.Bool `tfsdk:"isHidden" tf:"optional"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings any `tfsdk:"parameterMappings" tf:"optional"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position types.Object `tfsdk:"position" tf:"optional,object"`
	// Custom title of the widget
	Title types.String `tfsdk:"title" tf:"optional"`
	// Timestamp of the last time this object was updated.
	UpdatedAt types.String `tfsdk:"updated_at" tf:"optional"`
}

func (newState *WidgetOptions) SyncEffectiveFieldsDuringCreateOrUpdate(plan WidgetOptions) {
}

func (newState *WidgetOptions) SyncEffectiveFieldsDuringRead(existingState WidgetOptions) {
}

func (a WidgetOptions) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{
		"Position": reflect.TypeOf(WidgetPosition{}),
	}
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	AutoHeight types.Bool `tfsdk:"autoHeight" tf:"optional"`
	// column in the dashboard grid. Values start with 0
	Col types.Int64 `tfsdk:"col" tf:"optional"`
	// row in the dashboard grid. Values start with 0
	Row types.Int64 `tfsdk:"row" tf:"optional"`
	// width of the widget measured in dashboard grid cells
	SizeX types.Int64 `tfsdk:"sizeX" tf:"optional"`
	// height of the widget measured in dashboard grid cells
	SizeY types.Int64 `tfsdk:"sizeY" tf:"optional"`
}

func (newState *WidgetPosition) SyncEffectiveFieldsDuringCreateOrUpdate(plan WidgetPosition) {
}

func (newState *WidgetPosition) SyncEffectiveFieldsDuringRead(existingState WidgetPosition) {
}

func (a WidgetPosition) GetComplexFieldTypes() map[string]reflect.Type {
	return map[string]reflect.Type{}
}
