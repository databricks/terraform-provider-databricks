// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessControl struct {
	GroupName types.String `tfsdk:"group_name"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel PermissionLevel `tfsdk:"permission_level"`

	UserName types.String `tfsdk:"user_name"`
}

type Alert struct {
	// Timestamp when the alert was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Alert ID.
	Id types.String `tfsdk:"id"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt types.String `tfsdk:"last_triggered_at"`
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options *AlertOptions `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`

	Query *AlertQuery `tfsdk:"query"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State AlertState `tfsdk:"state"`
	// Timestamp when the alert was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`

	User *User `tfsdk:"user"`
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	Column types.String `tfsdk:"column"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody types.String `tfsdk:"custom_body"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject types.String `tfsdk:"custom_subject"`
	// State that alert evaluates to when query result is empty.
	EmptyResultState AlertOptionsEmptyResultState `tfsdk:"empty_result_state"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted types.Bool `tfsdk:"muted"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op types.String `tfsdk:"op"`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value any `tfsdk:"value"`
}

// State that alert evaluates to when query result is empty.
type AlertOptionsEmptyResultState string

const AlertOptionsEmptyResultStateOk AlertOptionsEmptyResultState = `ok`

const AlertOptionsEmptyResultStateTriggered AlertOptionsEmptyResultState = `triggered`

const AlertOptionsEmptyResultStateUnknown AlertOptionsEmptyResultState = `unknown`

// String representation for [fmt.Print]
func (f *AlertOptionsEmptyResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertOptionsEmptyResultState) Set(v string) error {
	switch v {
	case `ok`, `triggered`, `unknown`:
		*f = AlertOptionsEmptyResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ok", "triggered", "unknown"`, v)
	}
}

// Type always returns AlertOptionsEmptyResultState to satisfy [pflag.Value] interface
func (f *AlertOptionsEmptyResultState) Type() string {
	return "AlertOptionsEmptyResultState"
}

type AlertQuery struct {
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Query ID.
	Id types.String `tfsdk:"id"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`

	Options *QueryOptions `tfsdk:"options"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`

	Tags types.List `tfsdk:"tags"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id"`
}

// State of the alert. Possible values are: `unknown` (yet to be evaluated),
// `triggered` (evaluated and fulfilled trigger conditions), or `ok` (evaluated
// and did not fulfill trigger conditions).
type AlertState string

const AlertStateOk AlertState = `ok`

const AlertStateTriggered AlertState = `triggered`

const AlertStateUnknown AlertState = `unknown`

// String representation for [fmt.Print]
func (f *AlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertState) Set(v string) error {
	switch v {
	case `ok`, `triggered`, `unknown`:
		*f = AlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ok", "triggered", "unknown"`, v)
	}
}

// Type always returns AlertState to satisfy [pflag.Value] interface
func (f *AlertState) Type() string {
	return "AlertState"
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

// Cancel statement execution
type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-" url:"-"`
}

type CancelExecutionResponse struct {
}

type Channel struct {
	DbsqlVersion types.String `tfsdk:"dbsql_version"`

	Name ChannelName `tfsdk:"name"`
}

// Channel information for the SQL warehouse at the time of query execution
type ChannelInfo struct {
	// DBSQL Version the channel is mapped to
	DbsqlVersion types.String `tfsdk:"dbsql_version"`
	// Name of the channel
	Name ChannelName `tfsdk:"name"`
}

type ChannelName string

const ChannelNameChannelNameCurrent ChannelName = `CHANNEL_NAME_CURRENT`

const ChannelNameChannelNameCustom ChannelName = `CHANNEL_NAME_CUSTOM`

const ChannelNameChannelNamePreview ChannelName = `CHANNEL_NAME_PREVIEW`

const ChannelNameChannelNamePrevious ChannelName = `CHANNEL_NAME_PREVIOUS`

const ChannelNameChannelNameUnspecified ChannelName = `CHANNEL_NAME_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *ChannelName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ChannelName) Set(v string) error {
	switch v {
	case `CHANNEL_NAME_CURRENT`, `CHANNEL_NAME_CUSTOM`, `CHANNEL_NAME_PREVIEW`, `CHANNEL_NAME_PREVIOUS`, `CHANNEL_NAME_UNSPECIFIED`:
		*f = ChannelName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CHANNEL_NAME_CURRENT", "CHANNEL_NAME_CUSTOM", "CHANNEL_NAME_PREVIEW", "CHANNEL_NAME_PREVIOUS", "CHANNEL_NAME_UNSPECIFIED"`, v)
	}
}

// Type always returns ChannelName to satisfy [pflag.Value] interface
func (f *ChannelName) Type() string {
	return "ChannelName"
}

type ColumnInfo struct {
	// The name of the column.
	Name types.String `tfsdk:"name"`
	// The ordinal position of the column (starting at position 0).
	Position types.Int64 `tfsdk:"position"`
	// The format of the interval type.
	TypeIntervalType types.String `tfsdk:"type_interval_type"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName ColumnInfoTypeName `tfsdk:"type_name"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision types.Int64 `tfsdk:"type_precision"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale types.Int64 `tfsdk:"type_scale"`
	// The full SQL type specification.
	TypeText types.String `tfsdk:"type_text"`
}

// The name of the base data type. This doesn't include details for complex
// types such as STRUCT, MAP or ARRAY.
type ColumnInfoTypeName string

const ColumnInfoTypeNameArray ColumnInfoTypeName = `ARRAY`

const ColumnInfoTypeNameBinary ColumnInfoTypeName = `BINARY`

const ColumnInfoTypeNameBoolean ColumnInfoTypeName = `BOOLEAN`

const ColumnInfoTypeNameByte ColumnInfoTypeName = `BYTE`

const ColumnInfoTypeNameChar ColumnInfoTypeName = `CHAR`

const ColumnInfoTypeNameDate ColumnInfoTypeName = `DATE`

const ColumnInfoTypeNameDecimal ColumnInfoTypeName = `DECIMAL`

const ColumnInfoTypeNameDouble ColumnInfoTypeName = `DOUBLE`

const ColumnInfoTypeNameFloat ColumnInfoTypeName = `FLOAT`

const ColumnInfoTypeNameInt ColumnInfoTypeName = `INT`

const ColumnInfoTypeNameInterval ColumnInfoTypeName = `INTERVAL`

const ColumnInfoTypeNameLong ColumnInfoTypeName = `LONG`

const ColumnInfoTypeNameMap ColumnInfoTypeName = `MAP`

const ColumnInfoTypeNameNull ColumnInfoTypeName = `NULL`

const ColumnInfoTypeNameShort ColumnInfoTypeName = `SHORT`

const ColumnInfoTypeNameString ColumnInfoTypeName = `STRING`

const ColumnInfoTypeNameStruct ColumnInfoTypeName = `STRUCT`

const ColumnInfoTypeNameTimestamp ColumnInfoTypeName = `TIMESTAMP`

const ColumnInfoTypeNameUserDefinedType ColumnInfoTypeName = `USER_DEFINED_TYPE`

// String representation for [fmt.Print]
func (f *ColumnInfoTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnInfoTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TIMESTAMP`, `USER_DEFINED_TYPE`:
		*f = ColumnInfoTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TIMESTAMP", "USER_DEFINED_TYPE"`, v)
	}
}

// Type always returns ColumnInfoTypeName to satisfy [pflag.Value] interface
func (f *ColumnInfoTypeName) Type() string {
	return "ColumnInfoTypeName"
}

type CreateAlert struct {
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options AlertOptions `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
}

// Add visualization to a query
type CreateQueryVisualizationRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `tfsdk:"options"`
	// The identifier returned by :method:queries/create
	QueryId types.String `tfsdk:"query_id"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type types.String `tfsdk:"type"`
}

type CreateWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel *Channel `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `tfsdk:"spot_instance_policy"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `tfsdk:"tags"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType CreateWarehouseRequestWarehouseType `tfsdk:"warehouse_type"`
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type CreateWarehouseRequestWarehouseType string

const CreateWarehouseRequestWarehouseTypeClassic CreateWarehouseRequestWarehouseType = `CLASSIC`

const CreateWarehouseRequestWarehouseTypePro CreateWarehouseRequestWarehouseType = `PRO`

const CreateWarehouseRequestWarehouseTypeTypeUnspecified CreateWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *CreateWarehouseRequestWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = CreateWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns CreateWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *CreateWarehouseRequestWarehouseType) Type() string {
	return "CreateWarehouseRequestWarehouseType"
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id types.String `tfsdk:"id"`
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId types.String `tfsdk:"dashboard_id"`
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-" url:"-"`

	Options WidgetOptions `tfsdk:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text types.String `tfsdk:"text"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId types.String `tfsdk:"visualization_id"`
	// Width of a widget
	Width types.Int64 `tfsdk:"width"`
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit types.Bool `tfsdk:"can_edit"`
	// Timestamp when this dashboard was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled"`
	// The ID for this dashboard.
	Id types.String `tfsdk:"id"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft types.Bool `tfsdk:"is_draft"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name"`

	Options *DashboardOptions `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel `tfsdk:"permission_tier"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug types.String `tfsdk:"slug"`

	Tags types.List `tfsdk:"tags"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`

	User *User `tfsdk:"user"`
	// The ID of the user who owns the dashboard.
	UserId types.Int64 `tfsdk:"user_id"`

	Widgets types.List `tfsdk:"widgets"`
}

type DashboardEditContent struct {
	DashboardId types.String `tfsdk:"-" url:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at"`
}

type DashboardPostContent struct {
	// Indicates whether the dashboard filters are enabled
	DashboardFiltersEnabled types.Bool `tfsdk:"dashboard_filters_enabled"`
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	IsFavorite types.Bool `tfsdk:"is_favorite"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name types.String `tfsdk:"name"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	Id types.String `tfsdk:"id"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name types.String `tfsdk:"name"`
	// Reserved for internal use.
	PauseReason types.String `tfsdk:"pause_reason"`
	// Reserved for internal use.
	Paused types.Int64 `tfsdk:"paused"`
	// Reserved for internal use.
	SupportsAutoLimit types.Bool `tfsdk:"supports_auto_limit"`
	// Reserved for internal use.
	Syntax types.String `tfsdk:"syntax"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type types.String `tfsdk:"type"`
	// Reserved for internal use.
	ViewOnly types.Bool `tfsdk:"view_only"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

// Delete an alert
type DeleteAlertRequest struct {
	AlertId types.String `tfsdk:"-" url:"-"`
}

// Remove a dashboard
type DeleteDashboardRequest struct {
	DashboardId types.String `tfsdk:"-" url:"-"`
}

// Remove widget
type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id types.String `tfsdk:"-" url:"-"`
}

// Delete a query
type DeleteQueryRequest struct {
	QueryId types.String `tfsdk:"-" url:"-"`
}

// Remove visualization
type DeleteQueryVisualizationRequest struct {
	// Widget ID returned by :method:queryvizualisations/create
	Id types.String `tfsdk:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete a warehouse
type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-" url:"-"`
}

type DeleteWarehouseResponse struct {
}

// The fetch disposition provides two modes of fetching results: `INLINE` and
// `EXTERNAL_LINKS`.
//
// Statements executed with `INLINE` disposition will return result data inline,
// in `JSON_ARRAY` format, in a series of chunks. If a given statement produces
// a result set with a size larger than 25 MiB, that statement execution is
// aborted, and no result set will be available.
//
// **NOTE** Byte limits are computed based upon internal representations of the
// result set data, and might not match the sizes visible in JSON responses.
//
// Statements executed with `EXTERNAL_LINKS` disposition will return result data
// as external links: URLs that point to cloud storage internal to the
// workspace. Using `EXTERNAL_LINKS` disposition allows statements to generate
// arbitrarily sized result sets for fetching up to 100 GiB. The resulting links
// have two important properties:
//
// 1. They point to resources _external_ to the Databricks compute; therefore
// any associated authentication information (typically a personal access token,
// OAuth token, or similar) _must be removed_ when fetching from these links.
//
// 2. These are presigned URLs with a specific expiration, indicated in the
// response. The behavior when attempting to use an expired link is cloud
// specific.
type Disposition string

const DispositionExternalLinks Disposition = `EXTERNAL_LINKS`

const DispositionInline Disposition = `INLINE`

// String representation for [fmt.Print]
func (f *Disposition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Disposition) Set(v string) error {
	switch v {
	case `EXTERNAL_LINKS`, `INLINE`:
		*f = Disposition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL_LINKS", "INLINE"`, v)
	}
}

// Type always returns Disposition to satisfy [pflag.Value] interface
func (f *Disposition) Type() string {
	return "Disposition"
}

type EditAlert struct {
	AlertId types.String `tfsdk:"-" url:"-"`
	// Name of the alert.
	Name types.String `tfsdk:"name"`
	// Alert configuration options.
	Options AlertOptions `tfsdk:"options"`
	// Query ID.
	QueryId types.String `tfsdk:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm types.Int64 `tfsdk:"rearm"`
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel *Channel `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute.
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Required. Id of the warehouse to configure.
	Id types.String `tfsdk:"-" url:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `tfsdk:"spot_instance_policy"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `tfsdk:"tags"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType EditWarehouseRequestWarehouseType `tfsdk:"warehouse_type"`
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type EditWarehouseRequestWarehouseType string

const EditWarehouseRequestWarehouseTypeClassic EditWarehouseRequestWarehouseType = `CLASSIC`

const EditWarehouseRequestWarehouseTypePro EditWarehouseRequestWarehouseType = `PRO`

const EditWarehouseRequestWarehouseTypeTypeUnspecified EditWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *EditWarehouseRequestWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EditWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = EditWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns EditWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *EditWarehouseRequestWarehouseType) Type() string {
	return "EditWarehouseRequestWarehouseType"
}

type EditWarehouseResponse struct {
}

type EndpointConfPair struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	Details types.String `tfsdk:"details"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason *TerminationReason `tfsdk:"failure_reason"`
	// Deprecated. split into summary and details for security
	Message types.String `tfsdk:"message"`
	// Health status of the warehouse.
	Status Status `tfsdk:"status"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary types.String `tfsdk:"summary"`
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel *Channel `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *EndpointHealth `tfsdk:"health"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`
	// current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters"`
	// ODBC parameters for the SQL warehouse
	OdbcParams *OdbcParams `tfsdk:"odbc_params"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `tfsdk:"spot_instance_policy"`
	// State of the warehouse
	State State `tfsdk:"state"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `tfsdk:"tags"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType EndpointInfoWarehouseType `tfsdk:"warehouse_type"`
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type EndpointInfoWarehouseType string

const EndpointInfoWarehouseTypeClassic EndpointInfoWarehouseType = `CLASSIC`

const EndpointInfoWarehouseTypePro EndpointInfoWarehouseType = `PRO`

const EndpointInfoWarehouseTypeTypeUnspecified EndpointInfoWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *EndpointInfoWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointInfoWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = EndpointInfoWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns EndpointInfoWarehouseType to satisfy [pflag.Value] interface
func (f *EndpointInfoWarehouseType) Type() string {
	return "EndpointInfoWarehouseType"
}

type EndpointTagPair struct {
	Key types.String `tfsdk:"key"`

	Value types.String `tfsdk:"value"`
}

type EndpointTags struct {
	CustomTags types.List `tfsdk:"custom_tags"`
}

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	ByteLimit types.Int64 `tfsdk:"byte_limit"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog types.String `tfsdk:"catalog"`
	// The fetch disposition provides two modes of fetching results: `INLINE`
	// and `EXTERNAL_LINKS`.
	//
	// Statements executed with `INLINE` disposition will return result data
	// inline, in `JSON_ARRAY` format, in a series of chunks. If a given
	// statement produces a result set with a size larger than 25 MiB, that
	// statement execution is aborted, and no result set will be available.
	//
	// **NOTE** Byte limits are computed based upon internal representations of
	// the result set data, and might not match the sizes visible in JSON
	// responses.
	//
	// Statements executed with `EXTERNAL_LINKS` disposition will return result
	// data as external links: URLs that point to cloud storage internal to the
	// workspace. Using `EXTERNAL_LINKS` disposition allows statements to
	// generate arbitrarily sized result sets for fetching up to 100 GiB. The
	// resulting links have two important properties:
	//
	// 1. They point to resources _external_ to the Databricks compute;
	// therefore any associated authentication information (typically a personal
	// access token, OAuth token, or similar) _must be removed_ when fetching
	// from these links.
	//
	// 2. These are presigned URLs with a specific expiration, indicated in the
	// response. The behavior when attempting to use an expired link is cloud
	// specific.
	Disposition Disposition `tfsdk:"disposition"`
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
	Format Format `tfsdk:"format"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeout `tfsdk:"on_wait_timeout"`
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
	Parameters types.List `tfsdk:"parameters"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit types.Int64 `tfsdk:"row_limit"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema types.String `tfsdk:"schema"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`.
	Statement types.String `tfsdk:"statement"`
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
	WaitTimeout types.String `tfsdk:"wait_timeout"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

// When `wait_timeout > 0s`, the call will block up to the specified time. If
// the statement execution doesn't finish within this time, `on_wait_timeout`
// determines whether the execution should continue or be canceled. When set to
// `CONTINUE`, the statement execution continues asynchronously and the call
// returns a statement ID which can be used for polling with
// :method:statementexecution/getStatement. When set to `CANCEL`, the statement
// execution is canceled and the call returns with a `CANCELED` state.
type ExecuteStatementRequestOnWaitTimeout string

const ExecuteStatementRequestOnWaitTimeoutCancel ExecuteStatementRequestOnWaitTimeout = `CANCEL`

const ExecuteStatementRequestOnWaitTimeoutContinue ExecuteStatementRequestOnWaitTimeout = `CONTINUE`

// String representation for [fmt.Print]
func (f *ExecuteStatementRequestOnWaitTimeout) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExecuteStatementRequestOnWaitTimeout) Set(v string) error {
	switch v {
	case `CANCEL`, `CONTINUE`:
		*f = ExecuteStatementRequestOnWaitTimeout(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCEL", "CONTINUE"`, v)
	}
}

// Type always returns ExecuteStatementRequestOnWaitTimeout to satisfy [pflag.Value] interface
func (f *ExecuteStatementRequestOnWaitTimeout) Type() string {
	return "ExecuteStatementRequestOnWaitTimeout"
}

type ExecuteStatementResponse struct {
	// The result manifest provides schema and metadata for the result set.
	Manifest *ResultManifest `tfsdk:"manifest"`
	// Contains the result data of a single chunk when using `INLINE`
	// disposition. When using `EXTERNAL_LINKS` disposition, the array
	// `external_links` is used instead to provide presigned URLs to the result
	// data in cloud storage. Exactly one of these alternatives is used. (While
	// the `external_links` array prepares the API to return multiple links in a
	// single response. Currently only a single link is returned.)
	Result *ResultData `tfsdk:"result"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id"`
	// The status response includes execution state and if relevant, error
	// information.
	Status *StatementStatus `tfsdk:"status"`
}

type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration types.String `tfsdk:"expiration"`
	// A presigned URL pointing to a chunk of result data, hosted by an external
	// service, with a short expiration time (<= 15 minutes). As this URL
	// contains a temporary credential, it should be considered sensitive and
	// the client should not expose this URL in a log.
	ExternalLink types.String `tfsdk:"external_link"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders types.Map `tfsdk:"http_headers"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

type Format string

const FormatArrowStream Format = `ARROW_STREAM`

const FormatCsv Format = `CSV`

const FormatJsonArray Format = `JSON_ARRAY`

// String representation for [fmt.Print]
func (f *Format) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Format) Set(v string) error {
	switch v {
	case `ARROW_STREAM`, `CSV`, `JSON_ARRAY`:
		*f = Format(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARROW_STREAM", "CSV", "JSON_ARRAY"`, v)
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

// Get an alert
type GetAlertRequest struct {
	AlertId types.String `tfsdk:"-" url:"-"`
}

// Retrieve a definition
type GetDashboardRequest struct {
	DashboardId types.String `tfsdk:"-" url:"-"`
}

// Get object ACL
type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId types.String `tfsdk:"-" url:"-"`
	// The type of object permissions to check.
	ObjectType ObjectTypePlural `tfsdk:"-" url:"-"`
}

// Get a query definition.
type GetQueryRequest struct {
	QueryId types.String `tfsdk:"-" url:"-"`
}

type GetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType ObjectType `tfsdk:"object_type"`
}

// Get status, manifest, and result first chunk
type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-" url:"-"`
}

type GetStatementResponse struct {
	// The result manifest provides schema and metadata for the result set.
	Manifest *ResultManifest `tfsdk:"manifest"`
	// Contains the result data of a single chunk when using `INLINE`
	// disposition. When using `EXTERNAL_LINKS` disposition, the array
	// `external_links` is used instead to provide presigned URLs to the result
	// data in cloud storage. Exactly one of these alternatives is used. (While
	// the `external_links` array prepares the API to return multiple links in a
	// single response. Currently only a single link is returned.)
	Result *ResultData `tfsdk:"result"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"statement_id"`
	// The status response includes execution state and if relevant, error
	// information.
	Status *StatementStatus `tfsdk:"status"`
}

// Get result chunk by index
type GetStatementResultChunkNRequest struct {
	ChunkIndex types.Int64 `tfsdk:"-" url:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId types.String `tfsdk:"-" url:"-"`
}

// Get SQL warehouse permission levels
type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-" url:"-"`
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels types.List `tfsdk:"permission_levels"`
}

// Get SQL warehouse permissions
type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-" url:"-"`
}

// Get warehouse info
type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-" url:"-"`
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins types.Int64 `tfsdk:"auto_stop_mins"`
	// Channel Details
	Channel *Channel `tfsdk:"channel"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize types.String `tfsdk:"cluster_size"`
	// warehouse creator name
	CreatorName types.String `tfsdk:"creator_name"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton types.Bool `tfsdk:"enable_photon"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute types.Bool `tfsdk:"enable_serverless_compute"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *EndpointHealth `tfsdk:"health"`
	// unique identifier for warehouse
	Id types.String `tfsdk:"id"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// the jdbc connection string for this warehouse
	JdbcUrl types.String `tfsdk:"jdbc_url"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters types.Int64 `tfsdk:"max_num_clusters"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters types.Int64 `tfsdk:"min_num_clusters"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name types.String `tfsdk:"name"`
	// current number of active sessions for the warehouse
	NumActiveSessions types.Int64 `tfsdk:"num_active_sessions"`
	// current number of clusters running for the service
	NumClusters types.Int64 `tfsdk:"num_clusters"`
	// ODBC parameters for the SQL warehouse
	OdbcParams *OdbcParams `tfsdk:"odbc_params"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `tfsdk:"spot_instance_policy"`
	// State of the warehouse
	State State `tfsdk:"state"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `tfsdk:"tags"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType GetWarehouseResponseWarehouseType `tfsdk:"warehouse_type"`
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type GetWarehouseResponseWarehouseType string

const GetWarehouseResponseWarehouseTypeClassic GetWarehouseResponseWarehouseType = `CLASSIC`

const GetWarehouseResponseWarehouseTypePro GetWarehouseResponseWarehouseType = `PRO`

const GetWarehouseResponseWarehouseTypeTypeUnspecified GetWarehouseResponseWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *GetWarehouseResponseWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetWarehouseResponseWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = GetWarehouseResponseWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns GetWarehouseResponseWarehouseType to satisfy [pflag.Value] interface
func (f *GetWarehouseResponseWarehouseType) Type() string {
	return "GetWarehouseResponseWarehouseType"
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	Channel *Channel `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs `tfsdk:"config_param"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs `tfsdk:"global_param"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters *RepeatedEndpointConfPairs `tfsdk:"sql_configuration_parameters"`
}

// Security policy for warehouses
type GetWorkspaceWarehouseConfigResponseSecurityPolicy string

const GetWorkspaceWarehouseConfigResponseSecurityPolicyDataAccessControl GetWorkspaceWarehouseConfigResponseSecurityPolicy = `DATA_ACCESS_CONTROL`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyNone GetWorkspaceWarehouseConfigResponseSecurityPolicy = `NONE`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyPassthrough GetWorkspaceWarehouseConfigResponseSecurityPolicy = `PASSTHROUGH`

// String representation for [fmt.Print]
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Set(v string) error {
	switch v {
	case `DATA_ACCESS_CONTROL`, `NONE`, `PASSTHROUGH`:
		*f = GetWorkspaceWarehouseConfigResponseSecurityPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_ACCESS_CONTROL", "NONE", "PASSTHROUGH"`, v)
	}
}

// Type always returns GetWorkspaceWarehouseConfigResponseSecurityPolicy to satisfy [pflag.Value] interface
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Type() string {
	return "GetWorkspaceWarehouseConfigResponseSecurityPolicy"
}

// Get dashboard objects
type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order ListOrder `tfsdk:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-" url:"page,omitempty"`
	// Number of dashboards to return per page.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// Full text search term.
	Q types.String `tfsdk:"-" url:"q,omitempty"`
}

type ListOrder string

const ListOrderCreatedAt ListOrder = `created_at`

const ListOrderName ListOrder = `name`

// String representation for [fmt.Print]
func (f *ListOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListOrder) Set(v string) error {
	switch v {
	case `created_at`, `name`:
		*f = ListOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "created_at", "name"`, v)
	}
}

// Type always returns ListOrder to satisfy [pflag.Value] interface
func (f *ListOrder) Type() string {
	return "ListOrder"
}

// Get a list of queries
type ListQueriesRequest struct {
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
	Order types.String `tfsdk:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page types.Int64 `tfsdk:"-" url:"page,omitempty"`
	// Number of queries to return per page.
	PageSize types.Int64 `tfsdk:"-" url:"page_size,omitempty"`
	// Full text search term
	Q types.String `tfsdk:"-" url:"q,omitempty"`
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	HasNextPage types.Bool `tfsdk:"has_next_page"`
	// A token that can be used to get the next page of results.
	NextPageToken types.String `tfsdk:"next_page_token"`

	Res types.List `tfsdk:"res"`
}

// List Queries
type ListQueryHistoryRequest struct {
	// A filter to limit query history results. This field is optional.
	FilterBy *QueryFilter `tfsdk:"-" url:"filter_by,omitempty"`
	// Whether to include metrics about query.
	IncludeMetrics types.Bool `tfsdk:"-" url:"include_metrics,omitempty"`
	// Limit the number of results returned in one page. The default is 100.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListResponse struct {
	// The total number of dashboards.
	Count types.Int64 `tfsdk:"count"`
	// The current page being displayed.
	Page types.Int64 `tfsdk:"page"`
	// The number of dashboards per page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// List of dashboards returned.
	Results types.List `tfsdk:"results"`
}

// List warehouses
type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId types.Int64 `tfsdk:"-" url:"run_as_user_id,omitempty"`
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	Warehouses types.List `tfsdk:"warehouses"`
}

// If specified, allows multiple values to be selected for this parameter. Only
// applies to dropdown list and query-based dropdown list parameters.
type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	Prefix types.String `tfsdk:"prefix"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator types.String `tfsdk:"separator"`
	// Character that suffixes each selected parameter value.
	Suffix types.String `tfsdk:"suffix"`
}

// A singular noun object type.
type ObjectType string

const ObjectTypeAlert ObjectType = `alert`

const ObjectTypeDashboard ObjectType = `dashboard`

const ObjectTypeDataSource ObjectType = `data_source`

const ObjectTypeQuery ObjectType = `query`

// String representation for [fmt.Print]
func (f *ObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ObjectType) Set(v string) error {
	switch v {
	case `alert`, `dashboard`, `data_source`, `query`:
		*f = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "alert", "dashboard", "data_source", "query"`, v)
	}
}

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (f *ObjectType) Type() string {
	return "ObjectType"
}

// Always a plural of the object type.
type ObjectTypePlural string

const ObjectTypePluralAlerts ObjectTypePlural = `alerts`

const ObjectTypePluralDashboards ObjectTypePlural = `dashboards`

const ObjectTypePluralDataSources ObjectTypePlural = `data_sources`

const ObjectTypePluralQueries ObjectTypePlural = `queries`

// String representation for [fmt.Print]
func (f *ObjectTypePlural) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ObjectTypePlural) Set(v string) error {
	switch v {
	case `alerts`, `dashboards`, `data_sources`, `queries`:
		*f = ObjectTypePlural(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "alerts", "dashboards", "data_sources", "queries"`, v)
	}
}

// Type always returns ObjectTypePlural to satisfy [pflag.Value] interface
func (f *ObjectTypePlural) Type() string {
	return "ObjectTypePlural"
}

type OdbcParams struct {
	Hostname types.String `tfsdk:"hostname"`

	Path types.String `tfsdk:"path"`

	Port types.Int64 `tfsdk:"port"`

	Protocol types.String `tfsdk:"protocol"`
}

// The singular form of the type of object which can be owned.
type OwnableObjectType string

const OwnableObjectTypeAlert OwnableObjectType = `alert`

const OwnableObjectTypeDashboard OwnableObjectType = `dashboard`

const OwnableObjectTypeQuery OwnableObjectType = `query`

// String representation for [fmt.Print]
func (f *OwnableObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OwnableObjectType) Set(v string) error {
	switch v {
	case `alert`, `dashboard`, `query`:
		*f = OwnableObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "alert", "dashboard", "query"`, v)
	}
}

// Type always returns OwnableObjectType to satisfy [pflag.Value] interface
func (f *OwnableObjectType) Type() string {
	return "OwnableObjectType"
}

type Parameter struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	EnumOptions types.String `tfsdk:"enumOptions"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions *MultiValuesOptions `tfsdk:"multiValuesOptions"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name types.String `tfsdk:"name"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId types.String `tfsdk:"queryId"`
	// The text displayed in a parameter picking widget.
	Title types.String `tfsdk:"title"`
	// Parameters can have several different types.
	Type ParameterType `tfsdk:"type"`
	// The default value for this parameter.
	Value any `tfsdk:"value"`
}

// Parameters can have several different types.
type ParameterType string

const ParameterTypeDatetime ParameterType = `datetime`

const ParameterTypeEnum ParameterType = `enum`

const ParameterTypeNumber ParameterType = `number`

const ParameterTypeQuery ParameterType = `query`

const ParameterTypeText ParameterType = `text`

// String representation for [fmt.Print]
func (f *ParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ParameterType) Set(v string) error {
	switch v {
	case `datetime`, `enum`, `number`, `query`, `text`:
		*f = ParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "datetime", "enum", "number", "query", "text"`, v)
	}
}

// Type always returns ParameterType to satisfy [pflag.Value] interface
func (f *ParameterType) Type() string {
	return "ParameterType"
}

// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query * `CAN_EDIT`:
// Can edit the query * `CAN_MANAGE`: Can manage the query
type PermissionLevel string

// Can edit the query
const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

// Can manage the query
const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

// Can run the query
const PermissionLevelCanRun PermissionLevel = `CAN_RUN`

// Can view the query
const PermissionLevelCanView PermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_RUN", "CAN_VIEW"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

// Whether plans exist for the execution, or the reason why they are missing
type PlansState string

const PlansStateEmpty PlansState = `EMPTY`

const PlansStateExists PlansState = `EXISTS`

const PlansStateIgnoredLargePlansSize PlansState = `IGNORED_LARGE_PLANS_SIZE`

const PlansStateIgnoredSmallDuration PlansState = `IGNORED_SMALL_DURATION`

const PlansStateIgnoredSparkPlanType PlansState = `IGNORED_SPARK_PLAN_TYPE`

const PlansStateUnknown PlansState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *PlansState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PlansState) Set(v string) error {
	switch v {
	case `EMPTY`, `EXISTS`, `IGNORED_LARGE_PLANS_SIZE`, `IGNORED_SMALL_DURATION`, `IGNORED_SPARK_PLAN_TYPE`, `UNKNOWN`:
		*f = PlansState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMPTY", "EXISTS", "IGNORED_LARGE_PLANS_SIZE", "IGNORED_SMALL_DURATION", "IGNORED_SPARK_PLAN_TYPE", "UNKNOWN"`, v)
	}
}

// Type always returns PlansState to satisfy [pflag.Value] interface
func (f *PlansState) Type() string {
	return "PlansState"
}

type Query struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit types.Bool `tfsdk:"can_edit"`
	// The timestamp when this query was created.
	CreatedAt types.String `tfsdk:"created_at"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// Query ID.
	Id types.String `tfsdk:"id"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived types.Bool `tfsdk:"is_archived"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft types.Bool `tfsdk:"is_draft"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite types.Bool `tfsdk:"is_favorite"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe types.Bool `tfsdk:"is_safe"`

	LastModifiedBy *User `tfsdk:"last_modified_by"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById types.Int64 `tfsdk:"last_modified_by_id"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId types.String `tfsdk:"latest_query_data_id"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`

	Options *QueryOptions `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel `tfsdk:"permission_tier"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash types.String `tfsdk:"query_hash"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
	// The timestamp at which this query was last updated.
	UpdatedAt types.String `tfsdk:"updated_at"`

	User *User `tfsdk:"user"`
	// The ID of the user who owns the query.
	UserId types.Int64 `tfsdk:"user_id"`

	Visualizations types.List `tfsdk:"visualizations"`
}

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `tfsdk:"options"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`

	QueryId types.String `tfsdk:"-" url:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// A filter to limit query history results. This field is optional.
type QueryFilter struct {
	QueryStartTimeRange *TimeRange `tfsdk:"query_start_time_range" url:"query_start_time_range,omitempty"`
	// A list of statement IDs.
	StatementIds types.List `tfsdk:"statement_ids" url:"statement_ids,omitempty"`

	Statuses types.List `tfsdk:"statuses" url:"statuses,omitempty"`
	// A list of user IDs who ran the queries.
	UserIds types.List `tfsdk:"user_ids" url:"user_ids,omitempty"`
	// A list of warehouse IDs.
	WarehouseIds types.List `tfsdk:"warehouse_ids" url:"warehouse_ids,omitempty"`
}

type QueryInfo struct {
	// Reserved for internal use.
	CanSubscribeToLiveQuery types.Bool `tfsdk:"canSubscribeToLiveQuery"`
	// Channel information for the SQL warehouse at the time of query execution
	ChannelUsed *ChannelInfo `tfsdk:"channel_used"`
	// Total execution time of the statement ( excluding result fetch time ).
	Duration types.Int64 `tfsdk:"duration"`
	// Alias for `warehouse_id`.
	EndpointId types.String `tfsdk:"endpoint_id"`
	// Message describing why the query could not complete.
	ErrorMessage types.String `tfsdk:"error_message"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId types.Int64 `tfsdk:"executed_as_user_id"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName types.String `tfsdk:"executed_as_user_name"`
	// The time execution of the query ended.
	ExecutionEndTimeMs types.Int64 `tfsdk:"execution_end_time_ms"`
	// Whether more updates for the query are expected.
	IsFinal types.Bool `tfsdk:"is_final"`
	// A key that can be used to look up query details.
	LookupKey types.String `tfsdk:"lookup_key"`
	// Metrics about query execution.
	Metrics *QueryMetrics `tfsdk:"metrics"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState PlansState `tfsdk:"plans_state"`
	// The time the query ended.
	QueryEndTimeMs types.Int64 `tfsdk:"query_end_time_ms"`
	// The query ID.
	QueryId types.String `tfsdk:"query_id"`
	// The time the query started.
	QueryStartTimeMs types.Int64 `tfsdk:"query_start_time_ms"`
	// The text of the query.
	QueryText types.String `tfsdk:"query_text"`
	// The number of results returned by the query.
	RowsProduced types.Int64 `tfsdk:"rows_produced"`
	// URL to the query plan.
	SparkUiUrl types.String `tfsdk:"spark_ui_url"`
	// Type of statement for this query
	StatementType QueryStatementType `tfsdk:"statement_type"`
	// Query status with one the following values: * `QUEUED`: Query has been
	// received and queued. * `RUNNING`: Query has started. * `CANCELED`: Query
	// has been cancelled by the user. * `FAILED`: Query has failed. *
	// `FINISHED`: Query has completed.
	Status QueryStatus `tfsdk:"status"`
	// The ID of the user who ran the query.
	UserId types.Int64 `tfsdk:"user_id"`
	// The email address or username of the user who ran the query.
	UserName types.String `tfsdk:"user_name"`
	// Warehouse ID.
	WarehouseId types.String `tfsdk:"warehouse_id"`
}

type QueryList struct {
	// The total number of queries.
	Count types.Int64 `tfsdk:"count"`
	// The page number that is currently displayed.
	Page types.Int64 `tfsdk:"page"`
	// The number of queries per page.
	PageSize types.Int64 `tfsdk:"page_size"`
	// List of queries returned.
	Results types.List `tfsdk:"results"`
}

// Metrics about query execution.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs types.Int64 `tfsdk:"compilation_time_ms"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs types.Int64 `tfsdk:"execution_time_ms"`
	// Reserved for internal use.
	MetadataTimeMs types.Int64 `tfsdk:"metadata_time_ms"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes types.Int64 `tfsdk:"network_sent_bytes"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	OverloadingQueueStartTimestamp types.Int64 `tfsdk:"overloading_queue_start_timestamp"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs types.Int64 `tfsdk:"photon_total_time_ms"`
	// Reserved for internal use.
	PlanningTimeMs types.Int64 `tfsdk:"planning_time_ms"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	ProvisioningQueueStartTimestamp types.Int64 `tfsdk:"provisioning_queue_start_timestamp"`
	// Total number of bytes in all tables not read due to pruning
	PrunedBytes types.Int64 `tfsdk:"pruned_bytes"`
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount types.Int64 `tfsdk:"pruned_files_count"`
	// Timestamp of when the underlying compute started compilation of the
	// query.
	QueryCompilationStartTimestamp types.Int64 `tfsdk:"query_compilation_start_timestamp"`
	// Reserved for internal use.
	QueryExecutionTimeMs types.Int64 `tfsdk:"query_execution_time_ms"`
	// Total size of data read by the query, in bytes.
	ReadBytes types.Int64 `tfsdk:"read_bytes"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes types.Int64 `tfsdk:"read_cache_bytes"`
	// Number of files read after pruning.
	ReadFilesCount types.Int64 `tfsdk:"read_files_count"`
	// Number of partitions read after pruning.
	ReadPartitionsCount types.Int64 `tfsdk:"read_partitions_count"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes types.Int64 `tfsdk:"read_remote_bytes"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs types.Int64 `tfsdk:"result_fetch_time_ms"`
	// true if the query result was fetched from cache, false otherwise.
	ResultFromCache types.Bool `tfsdk:"result_from_cache"`
	// Total number of rows returned by the query.
	RowsProducedCount types.Int64 `tfsdk:"rows_produced_count"`
	// Total number of rows read by the query.
	RowsReadCount types.Int64 `tfsdk:"rows_read_count"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes types.Int64 `tfsdk:"spill_to_disk_bytes"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs types.Int64 `tfsdk:"task_total_time_ms"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs types.Int64 `tfsdk:"total_time_ms"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes types.Int64 `tfsdk:"write_remote_bytes"`
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	Catalog types.String `tfsdk:"catalog"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt types.String `tfsdk:"moved_to_trash_at"`

	Parameters types.List `tfsdk:"parameters"`
	// The name of the schema to execute this query in.
	Schema types.String `tfsdk:"schema"`
}

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId types.String `tfsdk:"data_source_id"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description types.String `tfsdk:"description"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name types.String `tfsdk:"name"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `tfsdk:"options"`
	// The identifier of the workspace folder containing the object.
	Parent types.String `tfsdk:"parent"`
	// The text of the query to be run.
	Query types.String `tfsdk:"query"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `tfsdk:"run_as_role"`

	Tags types.List `tfsdk:"tags"`
}

// Type of statement for this query
type QueryStatementType string

const QueryStatementTypeAlter QueryStatementType = `ALTER`

const QueryStatementTypeAnalyze QueryStatementType = `ANALYZE`

const QueryStatementTypeCopy QueryStatementType = `COPY`

const QueryStatementTypeCreate QueryStatementType = `CREATE`

const QueryStatementTypeDelete QueryStatementType = `DELETE`

const QueryStatementTypeDescribe QueryStatementType = `DESCRIBE`

const QueryStatementTypeDrop QueryStatementType = `DROP`

const QueryStatementTypeExplain QueryStatementType = `EXPLAIN`

const QueryStatementTypeGrant QueryStatementType = `GRANT`

const QueryStatementTypeInsert QueryStatementType = `INSERT`

const QueryStatementTypeMerge QueryStatementType = `MERGE`

const QueryStatementTypeOptimize QueryStatementType = `OPTIMIZE`

const QueryStatementTypeOther QueryStatementType = `OTHER`

const QueryStatementTypeRefresh QueryStatementType = `REFRESH`

const QueryStatementTypeReplace QueryStatementType = `REPLACE`

const QueryStatementTypeRevoke QueryStatementType = `REVOKE`

const QueryStatementTypeSelect QueryStatementType = `SELECT`

const QueryStatementTypeSet QueryStatementType = `SET`

const QueryStatementTypeShow QueryStatementType = `SHOW`

const QueryStatementTypeTruncate QueryStatementType = `TRUNCATE`

const QueryStatementTypeUpdate QueryStatementType = `UPDATE`

const QueryStatementTypeUse QueryStatementType = `USE`

// String representation for [fmt.Print]
func (f *QueryStatementType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryStatementType) Set(v string) error {
	switch v {
	case `ALTER`, `ANALYZE`, `COPY`, `CREATE`, `DELETE`, `DESCRIBE`, `DROP`, `EXPLAIN`, `GRANT`, `INSERT`, `MERGE`, `OPTIMIZE`, `OTHER`, `REFRESH`, `REPLACE`, `REVOKE`, `SELECT`, `SET`, `SHOW`, `TRUNCATE`, `UPDATE`, `USE`:
		*f = QueryStatementType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALTER", "ANALYZE", "COPY", "CREATE", "DELETE", "DESCRIBE", "DROP", "EXPLAIN", "GRANT", "INSERT", "MERGE", "OPTIMIZE", "OTHER", "REFRESH", "REPLACE", "REVOKE", "SELECT", "SET", "SHOW", "TRUNCATE", "UPDATE", "USE"`, v)
	}
}

// Type always returns QueryStatementType to satisfy [pflag.Value] interface
func (f *QueryStatementType) Type() string {
	return "QueryStatementType"
}

// Query status with one the following values: * `QUEUED`: Query has been
// received and queued. * `RUNNING`: Query has started. * `CANCELED`: Query has
// been cancelled by the user. * `FAILED`: Query has failed. * `FINISHED`: Query
// has completed.
type QueryStatus string

// Query has been cancelled by the user.
const QueryStatusCanceled QueryStatus = `CANCELED`

// Query has failed.
const QueryStatusFailed QueryStatus = `FAILED`

// Query has completed.
const QueryStatusFinished QueryStatus = `FINISHED`

// Query has been received and queued.
const QueryStatusQueued QueryStatus = `QUEUED`

// Query has started.
const QueryStatusRunning QueryStatus = `RUNNING`

// String representation for [fmt.Print]
func (f *QueryStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryStatus) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `FINISHED`, `QUEUED`, `RUNNING`:
		*f = QueryStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "FINISHED", "QUEUED", "RUNNING"`, v)
	}
}

// Type always returns QueryStatus to satisfy [pflag.Value] interface
func (f *QueryStatus) Type() string {
	return "QueryStatus"
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	ConfigPair types.List `tfsdk:"config_pair"`

	ConfigurationPairs types.List `tfsdk:"configuration_pairs"`
}

// Restore a dashboard
type RestoreDashboardRequest struct {
	DashboardId types.String `tfsdk:"-" url:"-"`
}

// Restore a query
type RestoreQueryRequest struct {
	QueryId types.String `tfsdk:"-" url:"-"`
}

type RestoreResponse struct {
}

// Contains the result data of a single chunk when using `INLINE` disposition.
// When using `EXTERNAL_LINKS` disposition, the array `external_links` is used
// instead to provide presigned URLs to the result data in cloud storage.
// Exactly one of these alternatives is used. (While the `external_links` array
// prepares the API to return multiple links in a single response. Currently
// only a single link is returned.)
type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount types.Int64 `tfsdk:"byte_count"`
	// The position within the sequence of result set chunks.
	ChunkIndex types.Int64 `tfsdk:"chunk_index"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray types.List `tfsdk:"data_array"`

	ExternalLinks types.List `tfsdk:"external_links"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex types.Int64 `tfsdk:"next_chunk_index"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink types.String `tfsdk:"next_chunk_internal_link"`
	// The number of rows within the result chunk.
	RowCount types.Int64 `tfsdk:"row_count"`
	// The starting row offset within the result set.
	RowOffset types.Int64 `tfsdk:"row_offset"`
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks types.List `tfsdk:"chunks"`

	Format Format `tfsdk:"format"`
	// The schema is an ordered list of column descriptions.
	Schema *ResultSchema `tfsdk:"schema"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount types.Int64 `tfsdk:"total_byte_count"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount types.Int64 `tfsdk:"total_chunk_count"`
	// The total number of rows in the result set.
	TotalRowCount types.Int64 `tfsdk:"total_row_count"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated types.Bool `tfsdk:"truncated"`
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {
	ColumnCount types.Int64 `tfsdk:"column_count"`

	Columns types.List `tfsdk:"columns"`
}

// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as owner"
// behavior)
type RunAsRole string

const RunAsRoleOwner RunAsRole = `owner`

const RunAsRoleViewer RunAsRole = `viewer`

// String representation for [fmt.Print]
func (f *RunAsRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunAsRole) Set(v string) error {
	switch v {
	case `owner`, `viewer`:
		*f = RunAsRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "owner", "viewer"`, v)
	}
}

// Type always returns RunAsRole to satisfy [pflag.Value] interface
func (f *RunAsRole) Type() string {
	return "RunAsRole"
}

type ServiceError struct {
	ErrorCode ServiceErrorCode `tfsdk:"error_code"`
	// A brief summary of the error condition.
	Message types.String `tfsdk:"message"`
}

type ServiceErrorCode string

const ServiceErrorCodeAborted ServiceErrorCode = `ABORTED`

const ServiceErrorCodeAlreadyExists ServiceErrorCode = `ALREADY_EXISTS`

const ServiceErrorCodeBadRequest ServiceErrorCode = `BAD_REQUEST`

const ServiceErrorCodeCancelled ServiceErrorCode = `CANCELLED`

const ServiceErrorCodeDeadlineExceeded ServiceErrorCode = `DEADLINE_EXCEEDED`

const ServiceErrorCodeInternalError ServiceErrorCode = `INTERNAL_ERROR`

const ServiceErrorCodeIoError ServiceErrorCode = `IO_ERROR`

const ServiceErrorCodeNotFound ServiceErrorCode = `NOT_FOUND`

const ServiceErrorCodeResourceExhausted ServiceErrorCode = `RESOURCE_EXHAUSTED`

const ServiceErrorCodeServiceUnderMaintenance ServiceErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ServiceErrorCodeTemporarilyUnavailable ServiceErrorCode = `TEMPORARILY_UNAVAILABLE`

const ServiceErrorCodeUnauthenticated ServiceErrorCode = `UNAUTHENTICATED`

const ServiceErrorCodeUnknown ServiceErrorCode = `UNKNOWN`

const ServiceErrorCodeWorkspaceTemporarilyUnavailable ServiceErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ServiceErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServiceErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `DEADLINE_EXCEEDED`, `INTERNAL_ERROR`, `IO_ERROR`, `NOT_FOUND`, `RESOURCE_EXHAUSTED`, `SERVICE_UNDER_MAINTENANCE`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNKNOWN`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ServiceErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "DEADLINE_EXCEEDED", "INTERNAL_ERROR", "IO_ERROR", "NOT_FOUND", "RESOURCE_EXHAUSTED", "SERVICE_UNDER_MAINTENANCE", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNKNOWN", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Type always returns ServiceErrorCode to satisfy [pflag.Value] interface
func (f *ServiceErrorCode) Type() string {
	return "ServiceErrorCode"
}

// Set object ACL
type SetRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId types.String `tfsdk:"-" url:"-"`
	// The type of object permission to set.
	ObjectType ObjectTypePlural `tfsdk:"-" url:"-"`
}

type SetResponse struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId types.String `tfsdk:"object_id"`
	// A singular noun object type.
	ObjectType ObjectType `tfsdk:"object_type"`
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel *Channel `tfsdk:"channel"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs `tfsdk:"config_param"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig types.List `tfsdk:"data_access_config"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes types.List `tfsdk:"enabled_warehouse_types"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs `tfsdk:"global_param"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount types.String `tfsdk:"google_service_account"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn"`
	// Security policy for warehouses
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy `tfsdk:"security_policy"`
	// SQL configuration parameters
	SqlConfigurationParameters *RepeatedEndpointConfPairs `tfsdk:"sql_configuration_parameters"`
}

// Security policy for warehouses
type SetWorkspaceWarehouseConfigRequestSecurityPolicy string

const SetWorkspaceWarehouseConfigRequestSecurityPolicyDataAccessControl SetWorkspaceWarehouseConfigRequestSecurityPolicy = `DATA_ACCESS_CONTROL`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyNone SetWorkspaceWarehouseConfigRequestSecurityPolicy = `NONE`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyPassthrough SetWorkspaceWarehouseConfigRequestSecurityPolicy = `PASSTHROUGH`

// String representation for [fmt.Print]
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Set(v string) error {
	switch v {
	case `DATA_ACCESS_CONTROL`, `NONE`, `PASSTHROUGH`:
		*f = SetWorkspaceWarehouseConfigRequestSecurityPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_ACCESS_CONTROL", "NONE", "PASSTHROUGH"`, v)
	}
}

// Type always returns SetWorkspaceWarehouseConfigRequestSecurityPolicy to satisfy [pflag.Value] interface
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Type() string {
	return "SetWorkspaceWarehouseConfigRequestSecurityPolicy"
}

type SetWorkspaceWarehouseConfigResponse struct {
}

// Configurations whether the warehouse should use spot instances.
type SpotInstancePolicy string

const SpotInstancePolicyCostOptimized SpotInstancePolicy = `COST_OPTIMIZED`

const SpotInstancePolicyPolicyUnspecified SpotInstancePolicy = `POLICY_UNSPECIFIED`

const SpotInstancePolicyReliabilityOptimized SpotInstancePolicy = `RELIABILITY_OPTIMIZED`

// String representation for [fmt.Print]
func (f *SpotInstancePolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SpotInstancePolicy) Set(v string) error {
	switch v {
	case `COST_OPTIMIZED`, `POLICY_UNSPECIFIED`, `RELIABILITY_OPTIMIZED`:
		*f = SpotInstancePolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COST_OPTIMIZED", "POLICY_UNSPECIFIED", "RELIABILITY_OPTIMIZED"`, v)
	}
}

// Type always returns SpotInstancePolicy to satisfy [pflag.Value] interface
func (f *SpotInstancePolicy) Type() string {
	return "SpotInstancePolicy"
}

// Start a warehouse
type StartRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-" url:"-"`
}

type StartWarehouseResponse struct {
}

// State of the warehouse
type State string

const StateDeleted State = `DELETED`

const StateDeleting State = `DELETING`

const StateRunning State = `RUNNING`

const StateStarting State = `STARTING`

const StateStopped State = `STOPPED`

const StateStopping State = `STOPPING`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `DELETED`, `DELETING`, `RUNNING`, `STARTING`, `STOPPED`, `STOPPING`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "DELETING", "RUNNING", "STARTING", "STOPPED", "STOPPING"`, v)
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	Name types.String `tfsdk:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type types.String `tfsdk:"type"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value types.String `tfsdk:"value"`
}

// Statement execution state: - `PENDING`: waiting for warehouse - `RUNNING`:
// running - `SUCCEEDED`: execution was successful, result data available for
// fetch - `FAILED`: execution failed; reason for failure described in
// accomanying error message - `CANCELED`: user canceled; can come from explicit
// cancel call, or timeout with `on_wait_timeout=CANCEL` - `CLOSED`: execution
// successful, and statement closed; result no longer available for fetch
type StatementState string

// user canceled; can come from explicit cancel call, or timeout with
// `on_wait_timeout=CANCEL`
const StatementStateCanceled StatementState = `CANCELED`

// execution successful, and statement closed; result no longer available for
// fetch
const StatementStateClosed StatementState = `CLOSED`

// execution failed; reason for failure described in accomanying error message
const StatementStateFailed StatementState = `FAILED`

// waiting for warehouse
const StatementStatePending StatementState = `PENDING`

// running
const StatementStateRunning StatementState = `RUNNING`

// execution was successful, result data available for fetch
const StatementStateSucceeded StatementState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *StatementState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StatementState) Set(v string) error {
	switch v {
	case `CANCELED`, `CLOSED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = StatementState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "CLOSED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Type always returns StatementState to satisfy [pflag.Value] interface
func (f *StatementState) Type() string {
	return "StatementState"
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	Error *ServiceError `tfsdk:"error"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accomanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State StatementState `tfsdk:"state"`
}

// Health status of the warehouse.
type Status string

const StatusDegraded Status = `DEGRADED`

const StatusFailed Status = `FAILED`

const StatusHealthy Status = `HEALTHY`

const StatusStatusUnspecified Status = `STATUS_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `DEGRADED`, `FAILED`, `HEALTHY`, `STATUS_UNSPECIFIED`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEGRADED", "FAILED", "HEALTHY", "STATUS_UNSPECIFIED"`, v)
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

// Stop a warehouse
type StopRequest struct {
	// Required. Id of the SQL warehouse.
	Id types.String `tfsdk:"-" url:"-"`
}

type StopWarehouseResponse struct {
}

type Success struct {
	Message SuccessMessage `tfsdk:"message"`
}

type SuccessMessage string

const SuccessMessageSuccess SuccessMessage = `Success`

// String representation for [fmt.Print]
func (f *SuccessMessage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SuccessMessage) Set(v string) error {
	switch v {
	case `Success`:
		*f = SuccessMessage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Success"`, v)
	}
}

// Type always returns SuccessMessage to satisfy [pflag.Value] interface
func (f *SuccessMessage) Type() string {
	return "SuccessMessage"
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code TerminationReasonCode `tfsdk:"code"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters types.Map `tfsdk:"parameters"`
	// type of the termination
	Type TerminationReasonType `tfsdk:"type"`
}

// status code indicating why the cluster was terminated
type TerminationReasonCode string

const TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`

const TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`

// String representation for [fmt.Print]
func (f *TerminationReasonCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonCode) Set(v string) error {
	switch v {
	case `ABUSE_DETECTED`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_SHUTDOWN`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `DATABASE_CONNECTION_FAILURE`, `DBFS_COMPONENT_UNHEALTHY`, `DOCKER_IMAGE_PULL_FAILURE`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `EXECUTION_COMPONENT_UNHEALTHY`, `GCP_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_UNREACHABLE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_SPARK_IMAGE`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `STORAGE_DOWNLOAD_FAILURE`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_SHUTDOWN", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "DATABASE_CONNECTION_FAILURE", "DBFS_COMPONENT_UNHEALTHY", "DOCKER_IMAGE_PULL_FAILURE", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "EXECUTION_COMPONENT_UNHEALTHY", "GCP_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_DELETED", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_UNREACHABLE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_SPARK_IMAGE", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "STORAGE_DOWNLOAD_FAILURE", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR"`, v)
	}
}

// Type always returns TerminationReasonCode to satisfy [pflag.Value] interface
func (f *TerminationReasonCode) Type() string {
	return "TerminationReasonCode"
}

// type of the termination
type TerminationReasonType string

const TerminationReasonTypeClientError TerminationReasonType = `CLIENT_ERROR`

const TerminationReasonTypeCloudFailure TerminationReasonType = `CLOUD_FAILURE`

const TerminationReasonTypeServiceFault TerminationReasonType = `SERVICE_FAULT`

const TerminationReasonTypeSuccess TerminationReasonType = `SUCCESS`

// String representation for [fmt.Print]
func (f *TerminationReasonType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `SERVICE_FAULT`, `SUCCESS`:
		*f = TerminationReasonType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "SERVICE_FAULT", "SUCCESS"`, v)
	}
}

// Type always returns TerminationReasonType to satisfy [pflag.Value] interface
func (f *TerminationReasonType) Type() string {
	return "TerminationReasonType"
}

type TimeRange struct {
	// Limit results to queries that started before this time.
	EndTimeMs types.Int64 `tfsdk:"end_time_ms" url:"end_time_ms,omitempty"`
	// Limit results to queries that started after this time.
	StartTimeMs types.Int64 `tfsdk:"start_time_ms" url:"start_time_ms,omitempty"`
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner"`
}

// Transfer object ownership
type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner types.String `tfsdk:"new_owner"`
	// The ID of the object on which to change ownership.
	ObjectId TransferOwnershipObjectId `tfsdk:"-" url:"-"`
	// The type of object on which to change ownership.
	ObjectType OwnableObjectType `tfsdk:"-" url:"-"`
}

type UpdateResponse struct {
}

type User struct {
	Email types.String `tfsdk:"email"`

	Id types.Int64 `tfsdk:"id"`

	Name types.String `tfsdk:"name"`
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type Visualization struct {
	CreatedAt types.String `tfsdk:"created_at"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description types.String `tfsdk:"description"`
	// The UUID for this visualization.
	Id types.String `tfsdk:"id"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name types.String `tfsdk:"name"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `tfsdk:"options"`

	Query *Query `tfsdk:"query"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type types.String `tfsdk:"type"`

	UpdatedAt types.String `tfsdk:"updated_at"`
}

type WarehouseAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Permission level
	PermissionLevel WarehousePermissionLevel `tfsdk:"permission_level"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	AllPermissions types.List `tfsdk:"all_permissions"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name"`
	// name of the group
	GroupName types.String `tfsdk:"group_name"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`
	// name of the user
	UserName types.String `tfsdk:"user_name"`
}

type WarehousePermission struct {
	Inherited types.Bool `tfsdk:"inherited"`

	InheritedFromObject types.List `tfsdk:"inherited_from_object"`
	// Permission level
	PermissionLevel WarehousePermissionLevel `tfsdk:"permission_level"`
}

// Permission level
type WarehousePermissionLevel string

const WarehousePermissionLevelCanManage WarehousePermissionLevel = `CAN_MANAGE`

const WarehousePermissionLevelCanUse WarehousePermissionLevel = `CAN_USE`

const WarehousePermissionLevelIsOwner WarehousePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *WarehousePermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarehousePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`:
		*f = WarehousePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE", "IS_OWNER"`, v)
	}
}

// Type always returns WarehousePermissionLevel to satisfy [pflag.Value] interface
func (f *WarehousePermissionLevel) Type() string {
	return "WarehousePermissionLevel"
}

type WarehousePermissions struct {
	AccessControlList types.List `tfsdk:"access_control_list"`

	ObjectId types.String `tfsdk:"object_id"`

	ObjectType types.String `tfsdk:"object_type"`
}

type WarehousePermissionsDescription struct {
	Description types.String `tfsdk:"description"`
	// Permission level
	PermissionLevel WarehousePermissionLevel `tfsdk:"permission_level"`
}

type WarehousePermissionsRequest struct {
	AccessControlList types.List `tfsdk:"access_control_list"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId types.String `tfsdk:"-" url:"-"`
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled types.Bool `tfsdk:"enabled"`
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType WarehouseTypePairWarehouseType `tfsdk:"warehouse_type"`
}

// Warehouse type: `PRO` or `CLASSIC`.
type WarehouseTypePairWarehouseType string

const WarehouseTypePairWarehouseTypeClassic WarehouseTypePairWarehouseType = `CLASSIC`

const WarehouseTypePairWarehouseTypePro WarehouseTypePairWarehouseType = `PRO`

const WarehouseTypePairWarehouseTypeTypeUnspecified WarehouseTypePairWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *WarehouseTypePairWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarehouseTypePairWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = WarehouseTypePairWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Type always returns WarehouseTypePairWarehouseType to satisfy [pflag.Value] interface
func (f *WarehouseTypePairWarehouseType) Type() string {
	return "WarehouseTypePairWarehouseType"
}

type Widget struct {
	// The unique ID for this widget.
	Id types.String `tfsdk:"id"`

	Options *WidgetOptions `tfsdk:"options"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization *Visualization `tfsdk:"visualization"`
	// Unused field.
	Width types.Int64 `tfsdk:"width"`
}

type WidgetOptions struct {
	// Timestamp when this object was created
	CreatedAt types.String `tfsdk:"created_at"`
	// Custom description of the widget
	Description types.String `tfsdk:"description"`
	// Whether this widget is hidden on the dashboard.
	IsHidden types.Bool `tfsdk:"isHidden"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings any `tfsdk:"parameterMappings"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position *WidgetPosition `tfsdk:"position"`
	// Custom title of the widget
	Title types.String `tfsdk:"title"`
	// Timestamp of the last time this object was updated.
	UpdatedAt types.String `tfsdk:"updated_at"`
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	AutoHeight types.Bool `tfsdk:"autoHeight"`
	// column in the dashboard grid. Values start with 0
	Col types.Int64 `tfsdk:"col"`
	// row in the dashboard grid. Values start with 0
	Row types.Int64 `tfsdk:"row"`
	// width of the widget measured in dashboard grid cells
	SizeX types.Int64 `tfsdk:"sizeX"`
	// height of the widget measured in dashboard grid cells
	SizeY types.Int64 `tfsdk:"sizeY"`
}
