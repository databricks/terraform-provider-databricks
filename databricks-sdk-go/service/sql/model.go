// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccessControl struct {
	GroupName string `json:"group_name,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccessControl) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessControl) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Aggregation string

const AggregationAvg Aggregation = `AVG`

const AggregationCount Aggregation = `COUNT`

const AggregationCountDistinct Aggregation = `COUNT_DISTINCT`

const AggregationMax Aggregation = `MAX`

const AggregationMedian Aggregation = `MEDIAN`

const AggregationMin Aggregation = `MIN`

const AggregationStddev Aggregation = `STDDEV`

const AggregationSum Aggregation = `SUM`

// String representation for [fmt.Print]
func (f *Aggregation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Aggregation) Set(v string) error {
	switch v {
	case `AVG`, `COUNT`, `COUNT_DISTINCT`, `MAX`, `MEDIAN`, `MIN`, `STDDEV`, `SUM`:
		*f = Aggregation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVG", "COUNT", "COUNT_DISTINCT", "MAX", "MEDIAN", "MIN", "STDDEV", "SUM"`, v)
	}
}

// Values returns all possible values for Aggregation.
//
// There is no guarantee on the order of the values in the slice.
func (f *Aggregation) Values() []Aggregation {
	return []Aggregation{
		AggregationAvg,
		AggregationCount,
		AggregationCountDistinct,
		AggregationMax,
		AggregationMedian,
		AggregationMin,
		AggregationStddev,
		AggregationSum,
	}
}

// Type always returns Aggregation to satisfy [pflag.Value] interface
func (f *Aggregation) Type() string {
	return "Aggregation"
}

// A legacy Databricks SQL alert. An alert runs a saved query, compares a column
// of the result against a threshold, and notifies its subscribers when the
// condition is met. See
// https://docs.databricks.com/aws/en/sql/user/alerts/legacy for details.
type Alert struct {
	// Condition that triggers the alert: a column of the query result, a
	// comparison operator, and a threshold value. The value in the first row of
	// the column is evaluated; aggregations are not supported through this API.
	Condition *AlertCondition `json:"condition,omitempty"`
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom body of the alert notification. It can use the template variables
	// and HTML described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default notification body is used.
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of the alert notification, used as the email subject or
	// the Slack message header. It can use the template variables described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default subject is used.
	CustomSubject string `json:"custom_subject,omitempty"`
	// Display name of the alert, shown in list views and notifications. It must
	// not be empty.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the alert.
	Id string `json:"id,omitempty"`
	// Lifecycle state of the alert: `ACTIVE`, or `TRASHED` after it has been
	// moved to the trash.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Whether to notify subscribers when the alert state changes from
	// `TRIGGERED` back to `OK`. Defaults to true.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// Workspace path of the folder containing the alert. Defaults to the
	// creator's home folder when omitted on create. It cannot be changed after
	// creation.
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the saved query whose result the alert evaluates. The caller must
	// have permission to view the query.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds the alert waits after a notification before it can
	// notify again while its state stays `TRIGGERED`. If 0 or not specified,
	// the alert notifies once when its state changes to `TRIGGERED` and does
	// not notify again until its state has changed.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`
	// Current state of the alert's trigger status: `TRIGGERED` if the condition
	// was met on the last evaluation, `OK` if it was not, and `UNKNOWN` if the
	// alert has not yet been evaluated or ran into an error during the last
	// evaluation.
	State AlertState `json:"state,omitempty"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime string `json:"trigger_time,omitempty"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Alert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Alert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertCondition struct {
	// State the alert takes when the query returns no rows. Defaults to
	// `UNKNOWN` when not set.
	EmptyResultState AlertState `json:"empty_result_state,omitempty"`
	// Comparison operator applied between the operand column value and the
	// threshold. The alert is triggered when the comparison is true. `IS_NULL`
	// triggers when the column value is null; a threshold must still be
	// supplied but its value is ignored.
	Op AlertOperator `json:"op,omitempty"`
	// Left-hand side of the comparison. Through this API it must be a `column`
	// naming a column of the query result; the value in the first row of that
	// column is compared against the threshold.
	Operand *AlertConditionOperand `json:"operand,omitempty"`
	// Right-hand side of the comparison. Through this API it must be a fixed
	// `value` whose type matches the operand column.
	Threshold *AlertConditionThreshold `json:"threshold,omitempty"`
}

func (s *AlertCondition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AlertConditionOperand struct {
	// A column of the query result. Used for the `operand` of a condition.
	Column *AlertOperandColumn `json:"column,omitempty"`
}

func (s *AlertConditionOperand) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AlertConditionThreshold struct {
	// A fixed value. Used for the `threshold` of a condition.
	Value *AlertOperandValue `json:"value,omitempty"`
}

func (s *AlertConditionThreshold) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// UNSPECIFIED - default unspecify value for proto enum, do not use it in the
// code UNKNOWN - alert not yet evaluated TRIGGERED - alert is triggered OK -
// alert is not triggered ERROR - alert evaluation failed
type AlertEvaluationState string

const AlertEvaluationStateError AlertEvaluationState = `ERROR`

const AlertEvaluationStateOk AlertEvaluationState = `OK`

const AlertEvaluationStateTriggered AlertEvaluationState = `TRIGGERED`

const AlertEvaluationStateUnknown AlertEvaluationState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *AlertEvaluationState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertEvaluationState) Set(v string) error {
	switch v {
	case `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = AlertEvaluationState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for AlertEvaluationState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertEvaluationState) Values() []AlertEvaluationState {
	return []AlertEvaluationState{
		AlertEvaluationStateError,
		AlertEvaluationStateOk,
		AlertEvaluationStateTriggered,
		AlertEvaluationStateUnknown,
	}
}

// Type always returns AlertEvaluationState to satisfy [pflag.Value] interface
func (f *AlertEvaluationState) Type() string {
	return "AlertEvaluationState"
}

type AlertLifecycleState string

const AlertLifecycleStateActive AlertLifecycleState = `ACTIVE`

const AlertLifecycleStateDeleted AlertLifecycleState = `DELETED`

// String representation for [fmt.Print]
func (f *AlertLifecycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertLifecycleState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETED`:
		*f = AlertLifecycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETED"`, v)
	}
}

// Values returns all possible values for AlertLifecycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertLifecycleState) Values() []AlertLifecycleState {
	return []AlertLifecycleState{
		AlertLifecycleStateActive,
		AlertLifecycleStateDeleted,
	}
}

// Type always returns AlertLifecycleState to satisfy [pflag.Value] interface
func (f *AlertLifecycleState) Type() string {
	return "AlertLifecycleState"
}

type AlertOperandColumn struct {
	// Name of the column in the query result.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertOperandColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertOperandColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertOperandValue struct {
	// Boolean threshold, for boolean columns.
	BoolValue bool `json:"bool_value,omitempty"`
	// Numeric threshold, for numeric columns.
	DoubleValue float64 `json:"double_value,omitempty"`
	// String threshold, for string or timestamp columns.
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertOperandValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertOperandValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertOperator string

const AlertOperatorEqual AlertOperator = `EQUAL`

const AlertOperatorGreaterThan AlertOperator = `GREATER_THAN`

const AlertOperatorGreaterThanOrEqual AlertOperator = `GREATER_THAN_OR_EQUAL`

const AlertOperatorIsNull AlertOperator = `IS_NULL`

const AlertOperatorLessThan AlertOperator = `LESS_THAN`

const AlertOperatorLessThanOrEqual AlertOperator = `LESS_THAN_OR_EQUAL`

const AlertOperatorNotEqual AlertOperator = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *AlertOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertOperator) Set(v string) error {
	switch v {
	case `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = AlertOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "IS_NULL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Values returns all possible values for AlertOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertOperator) Values() []AlertOperator {
	return []AlertOperator{
		AlertOperatorEqual,
		AlertOperatorGreaterThan,
		AlertOperatorGreaterThanOrEqual,
		AlertOperatorIsNull,
		AlertOperatorLessThan,
		AlertOperatorLessThanOrEqual,
		AlertOperatorNotEqual,
	}
}

// Type always returns AlertOperator to satisfy [pflag.Value] interface
func (f *AlertOperator) Type() string {
	return "AlertOperator"
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	Column string `json:"column"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string `json:"custom_subject,omitempty"`
	// State that alert evaluates to when query result is empty.
	EmptyResultState AlertOptionsEmptyResultState `json:"empty_result_state,omitempty"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted bool `json:"muted,omitempty"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op string `json:"op"`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value any `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for AlertOptionsEmptyResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertOptionsEmptyResultState) Values() []AlertOptionsEmptyResultState {
	return []AlertOptionsEmptyResultState{
		AlertOptionsEmptyResultStateOk,
		AlertOptionsEmptyResultStateTriggered,
		AlertOptionsEmptyResultStateUnknown,
	}
}

// Type always returns AlertOptionsEmptyResultState to satisfy [pflag.Value] interface
func (f *AlertOptionsEmptyResultState) Type() string {
	return "AlertOptionsEmptyResultState"
}

type AlertQuery struct {
	// The timestamp when this query was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Query ID.
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool `json:"is_draft,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool `json:"is_safe,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`

	Options *QueryOptions `json:"options,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the user who owns the query.
	UserId int `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertState string

const AlertStateOk AlertState = `OK`

const AlertStateTriggered AlertState = `TRIGGERED`

const AlertStateUnknown AlertState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *AlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertState) Set(v string) error {
	switch v {
	case `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = AlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for AlertState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertState) Values() []AlertState {
	return []AlertState{
		AlertStateOk,
		AlertStateTriggered,
		AlertStateUnknown,
	}
}

// Type always returns AlertState to satisfy [pflag.Value] interface
func (f *AlertState) Type() string {
	return "AlertState"
}

// A named parameter bound to the alert query. Only flat, named scalar
// parameters are supported; complex types such as ARRAY, MAP, and STRUCT are
// not.
type AlertStatementParameter struct {
	// The name of the parameter. Reference it in the query text as `:name`.
	// Required, must be non-empty, and must be unique across the alert's
	// parameters.
	Name string `json:"name"`
	// The SQL data type of the parameter, for example `STRING`, `INT`, or
	// `DECIMAL(10, 2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	Type string `json:"type,omitempty"`
	// The value bound to the parameter, represented as a string. If omitted,
	// the value is interpreted as NULL.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertStatementParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertStatementParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A SQL alert. An alert runs a query on a schedule, evaluates the result
// against a condition, and notifies users or notification destinations when the
// condition is met.
type AlertV2 struct {
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom description for the alert, used as the notification body.
	// Interpreted according to `custom_template_format`. The available template
	// variables are listed at
	// https://docs.databricks.com/aws/en/sql/user/alerts/create#notification-templates.
	CustomDescription string `json:"custom_description,omitempty"`
	// Custom summary for the alert, used as the notification subject.
	// Interpreted according to `custom_template_format`.
	CustomSummary string `json:"custom_summary,omitempty"`
	// The format used to interpret the `custom_summary` and
	// `custom_description` templates.
	//
	// `HTML` treats both templates as HTML with Mustache `{{VARIABLE_NAME}}`
	// placeholders. `MARKDOWN` treats them as Markdown with allowlisted
	// `@VARIABLE_NAME` placeholders and converts them to the format each
	// notification destination expects, so a single template renders correctly
	// in email, Slack, and Microsoft Teams.
	//
	// When unset, the API applies no default. Responses omit the field unless a
	// format is stored, and a template with no format is rendered as HTML.
	CustomTemplateFormat CustomTemplateFormat `json:"custom_template_format,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name"`
	// The actual identity that will be used to execute the alert. This is an
	// output-only field that shows the resolved run-as identity after applying
	// permissions and defaults.
	EffectiveRunAs *AlertV2RunAs `json:"effective_run_as,omitempty"`
	// Condition evaluated against the query result, and who is notified when it
	// is met.
	Evaluation AlertV2Evaluation `json:"evaluation"`
	// The canonical identifier of the alert to retrieve information about.
	Id string `json:"id,omitempty"`
	// Indicates whether the alert is active (`ACTIVE`) or has been moved to the
	// trash (`DELETED`).
	LifecycleState AlertLifecycleState `json:"lifecycle_state,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// A list of parameters to pass into the alert SQL query statement
	// containing parameter markers. Static values only.
	//
	// Reference a parameter in the query text as `:name`. Each parameter must
	// have a unique, non-empty name. Each parameter consists of a name, a
	// value, and optionally a type. To represent a NULL value, the `value`
	// field may be omitted or set to `null` explicitly. If the `type` field is
	// omitted, the value is interpreted as a string.
	//
	// If the type is given, parameters will be checked for type correctness
	// according to the given type. A value is correct if the provided string
	// can be converted to the requested type using the `cast` function. The
	// exact semantics are described in the section [`cast` function] of the SQL
	// language reference.
	//
	// [`cast` function]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Parameters []AlertStatementParameter `json:"parameters,omitempty"`
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text"`
	// Specifies the identity that will be used to run the alert. This field
	// allows you to configure alerts to run as a specific user or service
	// principal. - For user identity: Set `user_name` to the email of an active
	// workspace user. Users can only set this to their own email. - For service
	// principal: Set `service_principal_name` to the application ID. Requires
	// the `servicePrincipal/user` role. If not specified, the alert will run as
	// the request user.
	RunAs *AlertV2RunAs `json:"run_as,omitempty"`
	// Deprecated: Use `run_as` field instead. This field will be removed in a
	// future release. The run as username or application ID of service
	// principal. On Create and Update, this field can be set to application ID
	// of an active service principal. Setting this field requires the
	// servicePrincipal/user role.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// Schedule on which the alert runs its query and evaluates the condition.
	Schedule CronSchedule `json:"schedule"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertV2Evaluation struct {
	// Operator used for comparison in alert evaluation. The alert is triggered
	// when the comparison of the source value with the threshold is true.
	// `IS_NULL` and `IS_NOT_NULL` take no threshold.
	ComparisonOperator ComparisonOperator `json:"comparison_operator"`
	// Alert state if result is empty. If not set, an empty result sets the
	// state to `ERROR`. Please avoid setting this field to be `UNKNOWN` because
	// `UNKNOWN` state is planned to be deprecated.
	EmptyResultState AlertEvaluationState `json:"empty_result_state,omitempty"`
	// Timestamp of the last evaluation.
	LastEvaluatedAt string `json:"last_evaluated_at,omitempty"`
	// User or Notification Destination to notify when alert is triggered.
	Notification *AlertV2Notification `json:"notification,omitempty"`
	// Source column from result to use to evaluate alert
	Source AlertV2OperandColumn `json:"source"`
	// Latest state of alert evaluation: `TRIGGERED` when the condition was met,
	// `OK` when it was not, and `ERROR` when the evaluation failed.
	State AlertEvaluationState `json:"state,omitempty"`
	// Threshold to compare the source column against, either another column or
	// a fixed value. Not needed when `comparison_operator` is `IS_NULL` or
	// `IS_NOT_NULL`.
	Threshold *AlertV2Operand `json:"threshold,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2Evaluation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2Evaluation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertV2Notification struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// Number of seconds an alert waits after being triggered before it is
	// allowed to send another notification. If set to 0 or omitted, the alert
	// will not send any further notifications after the first trigger Setting
	// this value to 1 allows the alert to send a notification on every
	// evaluation where the condition is met, effectively making it always
	// retrigger for notification purposes.
	RetriggerSeconds int `json:"retrigger_seconds,omitempty"`
	// Users and notification destinations that receive notifications.
	Subscriptions []AlertV2Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2Notification) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2Notification) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertV2Operand struct {
	// Compare the source against another column of the query result.
	Column *AlertV2OperandColumn `json:"column,omitempty"`
	// Compare the source against a fixed value.
	Value *AlertV2OperandValue `json:"value,omitempty"`
}

func (s *AlertV2Operand) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type AlertV2OperandColumn struct {
	// Aggregation applied to the column across all rows of the query result:
	// `SUM`, `COUNT`, `COUNT_DISTINCT`, `AVG`, `MEDIAN`, `MIN`, `MAX` or
	// `STDDEV`. If not set, the value in the first row is used, which is
	// equivalent to using `First row` in the UI.
	Aggregation Aggregation `json:"aggregation,omitempty"`
	// Label for the column in notifications. Defaults to `name`.
	Display string `json:"display,omitempty"`
	// Name of the column in the query result.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2OperandColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2OperandColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertV2OperandValue struct {
	// Boolean threshold, for boolean columns.
	BoolValue bool `json:"bool_value,omitempty"`
	// Numeric threshold, for numeric columns.
	DoubleValue float64 `json:"double_value,omitempty"`
	// String threshold, for string or timestamp columns.
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2OperandValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2OperandValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertV2RunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Can only set this field to their
	// own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2RunAs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2RunAs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AlertV2Subscription struct {
	// ID of the notification destination to notify, as returned by the
	// Notification Destinations API.
	DestinationId string `json:"destination_id,omitempty"`
	// Email address of the workspace user to notify.
	UserEmail string `json:"user_email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AlertV2Subscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2Subscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BaseChunkInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseChunkInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
}

func (s *CancelExecutionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {
	// Databricks SQL version to pin to. Set it if and only if name is
	// CHANNEL_NAME_CUSTOM.
	DbsqlVersion string `json:"dbsql_version,omitempty"`
	// Channel to follow. CHANNEL_NAME_CURRENT is the default, and
	// CHANNEL_NAME_PREVIEW carries the upcoming Databricks SQL release for
	// testing.
	Name ChannelName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Channel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Channel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	DbsqlVersion string `json:"dbsql_version,omitempty"`
	// Name of the channel
	Name ChannelName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ChannelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ChannelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ChannelName string

const ChannelNameChannelNameCurrent ChannelName = `CHANNEL_NAME_CURRENT`

const ChannelNameChannelNameCustom ChannelName = `CHANNEL_NAME_CUSTOM`

const ChannelNameChannelNamePreview ChannelName = `CHANNEL_NAME_PREVIEW`

const ChannelNameChannelNamePrevious ChannelName = `CHANNEL_NAME_PREVIOUS`

// String representation for [fmt.Print]
func (f *ChannelName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ChannelName) Set(v string) error {
	switch v {
	case `CHANNEL_NAME_CURRENT`, `CHANNEL_NAME_CUSTOM`, `CHANNEL_NAME_PREVIEW`, `CHANNEL_NAME_PREVIOUS`:
		*f = ChannelName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CHANNEL_NAME_CURRENT", "CHANNEL_NAME_CUSTOM", "CHANNEL_NAME_PREVIEW", "CHANNEL_NAME_PREVIOUS"`, v)
	}
}

// Values returns all possible values for ChannelName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ChannelName) Values() []ChannelName {
	return []ChannelName{
		ChannelNameChannelNameCurrent,
		ChannelNameChannelNameCustom,
		ChannelNameChannelNamePreview,
		ChannelNameChannelNamePrevious,
	}
}

// Type always returns ChannelName to satisfy [pflag.Value] interface
func (f *ChannelName) Type() string {
	return "ChannelName"
}

type ClientConfig struct {
	// allow_custom_js_visualizations is not supported/implemneted.
	AllowCustomJsVisualizations bool `json:"allow_custom_js_visualizations,omitempty"`

	AllowDownloads bool `json:"allow_downloads,omitempty"`

	AllowExternalShares bool `json:"allow_external_shares,omitempty"`

	AllowSubscriptions bool `json:"allow_subscriptions,omitempty"`

	DateFormat string `json:"date_format,omitempty"`

	DateTimeFormat string `json:"date_time_format,omitempty"`

	DisablePublish bool `json:"disable_publish,omitempty"`

	EnableLegacyAutodetectTypes bool `json:"enable_legacy_autodetect_types,omitempty"`

	FeatureShowPermissionsControl bool `json:"feature_show_permissions_control,omitempty"`

	HidePlotlyModeBar bool `json:"hide_plotly_mode_bar,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClientConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClientConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ColumnInfo struct {
	// The name of the column.
	Name string `json:"name,omitempty"`
	// The ordinal position of the column (starting at position 0).
	Position int `json:"position,omitempty"`
	// The format of the interval type.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName ColumnInfoTypeName `json:"type_name,omitempty"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision int `json:"type_precision,omitempty"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale int `json:"type_scale,omitempty"`
	// The full SQL type specification.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ColumnInfoTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnInfoTypeName) Values() []ColumnInfoTypeName {
	return []ColumnInfoTypeName{
		ColumnInfoTypeNameArray,
		ColumnInfoTypeNameBinary,
		ColumnInfoTypeNameBoolean,
		ColumnInfoTypeNameByte,
		ColumnInfoTypeNameChar,
		ColumnInfoTypeNameDate,
		ColumnInfoTypeNameDecimal,
		ColumnInfoTypeNameDouble,
		ColumnInfoTypeNameFloat,
		ColumnInfoTypeNameInt,
		ColumnInfoTypeNameInterval,
		ColumnInfoTypeNameLong,
		ColumnInfoTypeNameMap,
		ColumnInfoTypeNameNull,
		ColumnInfoTypeNameShort,
		ColumnInfoTypeNameString,
		ColumnInfoTypeNameStruct,
		ColumnInfoTypeNameTimestamp,
		ColumnInfoTypeNameUserDefinedType,
	}
}

// Type always returns ColumnInfoTypeName to satisfy [pflag.Value] interface
func (f *ColumnInfoTypeName) Type() string {
	return "ColumnInfoTypeName"
}

type ComparisonOperator string

const ComparisonOperatorEqual ComparisonOperator = `EQUAL`

const ComparisonOperatorGreaterThan ComparisonOperator = `GREATER_THAN`

const ComparisonOperatorGreaterThanOrEqual ComparisonOperator = `GREATER_THAN_OR_EQUAL`

const ComparisonOperatorIsNotNull ComparisonOperator = `IS_NOT_NULL`

const ComparisonOperatorIsNull ComparisonOperator = `IS_NULL`

const ComparisonOperatorLessThan ComparisonOperator = `LESS_THAN`

const ComparisonOperatorLessThanOrEqual ComparisonOperator = `LESS_THAN_OR_EQUAL`

const ComparisonOperatorNotEqual ComparisonOperator = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *ComparisonOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComparisonOperator) Set(v string) error {
	switch v {
	case `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NOT_NULL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = ComparisonOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "IS_NOT_NULL", "IS_NULL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Values returns all possible values for ComparisonOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		ComparisonOperatorEqual,
		ComparisonOperatorGreaterThan,
		ComparisonOperatorGreaterThanOrEqual,
		ComparisonOperatorIsNotNull,
		ComparisonOperatorIsNull,
		ComparisonOperatorLessThan,
		ComparisonOperatorLessThanOrEqual,
		ComparisonOperatorNotEqual,
	}
}

// Type always returns ComparisonOperator to satisfy [pflag.Value] interface
func (f *ComparisonOperator) Type() string {
	return "ComparisonOperator"
}

type CreateAlert struct {
	// Name of the alert.
	Name string `json:"name"`
	// Alert configuration options.
	Options AlertOptions `json:"options"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// Query ID.
	QueryId string `json:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateAlertRequest struct {
	// The alert to create. A display name, a `query_id`, and a `condition` are
	// required.
	Alert *CreateAlertRequestAlert `json:"alert,omitempty"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAlertRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAlertRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A legacy Databricks SQL alert. An alert runs a saved query, compares a column
// of the result against a threshold, and notifies its subscribers when the
// condition is met. See
// https://docs.databricks.com/aws/en/sql/user/alerts/legacy for details.
type CreateAlertRequestAlert struct {
	// Condition that triggers the alert: a column of the query result, a
	// comparison operator, and a threshold value. The value in the first row of
	// the column is evaluated; aggregations are not supported through this API.
	Condition *AlertCondition `json:"condition,omitempty"`
	// Custom body of the alert notification. It can use the template variables
	// and HTML described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default notification body is used.
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of the alert notification, used as the email subject or
	// the Slack message header. It can use the template variables described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default subject is used.
	CustomSubject string `json:"custom_subject,omitempty"`
	// Display name of the alert, shown in list views and notifications. It must
	// not be empty.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to notify subscribers when the alert state changes from
	// `TRIGGERED` back to `OK`. Defaults to true.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// Workspace path of the folder containing the alert. Defaults to the
	// creator's home folder when omitted on create. It cannot be changed after
	// creation.
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the saved query whose result the alert evaluates. The caller must
	// have permission to view the query.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds the alert waits after a notification before it can
	// notify again while its state stays `TRIGGERED`. If 0 or not specified,
	// the alert notifies once when its state changes to `TRIGGERED` and does
	// not notify again until its state has changed.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAlertRequestAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAlertRequestAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateAlertV2Request struct {
	Alert AlertV2 `json:"alert"`
}

func (s *CreateAlertV2Request) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateDefaultWarehouseOverrideRequest struct {
	// Required. The default warehouse override to create.
	DefaultWarehouseOverride DefaultWarehouseOverride `json:"default_warehouse_override"`
	// Required. The ID to use for the override, which will become the final
	// component of the override's resource name. Can be a numeric user ID or
	// the literal string "me" for the current user.
	DefaultWarehouseOverrideId string `json:"-" url:"default_warehouse_override_id"`
}

func (s *CreateDefaultWarehouseOverrideRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateQueryRequest struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`
	// The query to create. A display name and query text are required.
	Query *CreateQueryRequestQuery `json:"query,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateQueryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A saved Databricks SQL query. A query stores the SQL text, the SQL warehouse
// it runs on, its parameters, and metadata such as the display name,
// description, owner, and tags.
type CreateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// List of query parameter definitions.
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Workspace path of the workspace folder containing the object.
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Whose credentials the query runs with when it is shared. `VIEWER` uses
	// the credentials of the user running the query. `OWNER` uses the owner's
	// credentials where the product supports that, such as for legacy alerts
	// and jobs. When omitted on create, the workspace default applies. A
	// workspace admin can disallow `OWNER`.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`
	// Tags attached to the query. Tags appear in the query list and can be used
	// to filter it.
	Tags []string `json:"tags,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateQueryRequestQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateQueryRequestQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description string `json:"description,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `json:"options"`
	// The identifier returned by :method:queries/create
	QueryId string `json:"query_id"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type string `json:"type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateQueryVisualizationsLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateQueryVisualizationsLegacyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateVisualizationRequest struct {
	Visualization *CreateVisualizationRequestVisualization `json:"visualization,omitempty"`
}

func (s *CreateVisualizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	DisplayName string `json:"display_name,omitempty"`
	// UUID of the query that the visualization is attached to.
	QueryId string `json:"query_id,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateVisualizationRequestVisualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVisualizationRequestVisualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Creates a new SQL warehouse.
type CreateWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 1 min for serverless warehouses. -
	// Must be == 0 or >= 10 mins for classic and pro warehouses. - 0 indicates
	// no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel that determines which Databricks SQL version the warehouse runs.
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large - 5X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// Name of the user who created the warehouse. Set by the server and ignored
	// in requests.
	CreatorName string `json:"creator_name,omitempty"`
	// Compute version label for a Real-Time warehouse. This field is output
	// only.
	CustomerVersionLabel string `json:"customer_version_label,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to true.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum autoscaling capacity of a Real-Time SQL warehouse, in DBU per
	// hour. Must be at least the minimum DBU for the selected warehouse size
	// and no greater than the configured DBU limit. Actual capacity may be
	// below this ceiling.
	MaxDbu int `json:"max_dbu,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the warehouse.
	//
	// Supported values: - Must be unique within an org. - Must be 1 to 100
	// characters.
	Name string `json:"name,omitempty"`
	// Cluster size multiplier for Reyden warehouses. Scales the effective
	// compute capacity relative to the nominal virtual cluster size. Only
	// meaningful when the warehouse is of Reyden type.
	//
	// Accepted range: 1.0 to the workspace-configured maximum (decimals
	// allowed). The default maximum is 100.0. - Values below 1.0 are normalized
	// to 1.0. - Values above the configured maximum are capped at that maximum.
	// - When absent, the warehouse is treated as having a multiplier of 1.0.
	ReydenClusterSizeMultiplier float64 `json:"reyden_cluster_size_multiplier,omitempty"`
	// Configures whether the warehouse should use spot instances.
	// COST_OPTIMIZED uses spot instances where the cloud provider supports
	// them, and RELIABILITY_OPTIMIZED uses on-demand instances only.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// Maximum time in seconds that a statement can run on this warehouse before
	// it is cancelled. Set to 0 to remove the warehouse-level timeout. Changing
	// the value of a running warehouse restarts it. See
	// https://docs.databricks.com/aws/en/sql/language-manual/parameters/statement_timeout.
	StatementTimeout int `json:"statement_timeout,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type: `PRO`, `CLASSIC`, or `REALTIME`. For serverless compute,
	// use `PRO` or `REALTIME` with `enable_serverless_compute=true`. `REALTIME`
	// requires serverless compute and Photon; `enable_photon` defaults to
	// `true`.
	WarehouseType CreateWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWarehouseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWarehouseRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateWarehouseRequestWarehouseType string

const CreateWarehouseRequestWarehouseTypeClassic CreateWarehouseRequestWarehouseType = `CLASSIC`

const CreateWarehouseRequestWarehouseTypePro CreateWarehouseRequestWarehouseType = `PRO`

const CreateWarehouseRequestWarehouseTypeRealtime CreateWarehouseRequestWarehouseType = `REALTIME`

const CreateWarehouseRequestWarehouseTypeTypeUnspecified CreateWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *CreateWarehouseRequestWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `REALTIME`, `TYPE_UNSPECIFIED`:
		*f = CreateWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "REALTIME", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for CreateWarehouseRequestWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateWarehouseRequestWarehouseType) Values() []CreateWarehouseRequestWarehouseType {
	return []CreateWarehouseRequestWarehouseType{
		CreateWarehouseRequestWarehouseTypeClassic,
		CreateWarehouseRequestWarehouseTypePro,
		CreateWarehouseRequestWarehouseTypeRealtime,
		CreateWarehouseRequestWarehouseTypeTypeUnspecified,
	}
}

// Type always returns CreateWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *CreateWarehouseRequestWarehouseType) Type() string {
	return "CreateWarehouseRequestWarehouseType"
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWarehouseResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWarehouseResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId string `json:"dashboard_id"`

	Options WidgetOptions `json:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text string `json:"text,omitempty"`
	// Query Visualization ID returned by :method:queryvisualizations/create.
	VisualizationId string `json:"visualization_id,omitempty"`
	// Width of a widget
	Width int `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWidget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWidget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CronSchedule struct {
	// Indicates whether this schedule is paused. While paused, the alert is not
	// evaluated.
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// A cron expression using quartz syntax that specifies the schedule for
	// this alert. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	QuartzCronSchedule string `json:"quartz_cron_schedule"`
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	TimezoneId string `json:"timezone_id"`
}

func (s *CronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Rendering format for the `custom_summary` and `custom_description` templates.
type CustomTemplateFormat string

const CustomTemplateFormatHtml CustomTemplateFormat = `HTML`

const CustomTemplateFormatMarkdown CustomTemplateFormat = `MARKDOWN`

// String representation for [fmt.Print]
func (f *CustomTemplateFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomTemplateFormat) Set(v string) error {
	switch v {
	case `HTML`, `MARKDOWN`:
		*f = CustomTemplateFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HTML", "MARKDOWN"`, v)
	}
}

// Values returns all possible values for CustomTemplateFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomTemplateFormat) Values() []CustomTemplateFormat {
	return []CustomTemplateFormat{
		CustomTemplateFormatHtml,
		CustomTemplateFormatMarkdown,
	}
}

// Type always returns CustomTemplateFormat to satisfy [pflag.Value] interface
func (f *CustomTemplateFormat) Type() string {
	return "CustomTemplateFormat"
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit bool `json:"can_edit,omitempty"`
	// Timestamp when this dashboard was created.
	CreatedAt string `json:"created_at,omitempty"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// The ID for this dashboard.
	Id string `json:"id,omitempty"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft bool `json:"is_draft,omitempty"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name,omitempty"`

	Options *DashboardOptions `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug string `json:"slug,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *User `json:"user,omitempty"`
	// The ID of the user who owns the dashboard.
	UserId int `json:"user_id,omitempty"`

	Widgets []Widget `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DashboardEditContent struct {
	DashboardId string `json:"-" url:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardEditContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardEditContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DashboardOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	Id string `json:"id,omitempty"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name string `json:"name,omitempty"`
	// Reserved for internal use.
	PauseReason string `json:"pause_reason,omitempty"`
	// Reserved for internal use.
	Paused int `json:"paused,omitempty"`
	// Reserved for internal use.
	SupportsAutoLimit bool `json:"supports_auto_limit,omitempty"`
	// Reserved for internal use.
	Syntax string `json:"syntax,omitempty"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type string `json:"type,omitempty"`
	// Reserved for internal use.
	ViewOnly bool `json:"view_only,omitempty"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DatePrecision string

const DatePrecisionDayPrecision DatePrecision = `DAY_PRECISION`

const DatePrecisionMinutePrecision DatePrecision = `MINUTE_PRECISION`

const DatePrecisionSecondPrecision DatePrecision = `SECOND_PRECISION`

// String representation for [fmt.Print]
func (f *DatePrecision) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatePrecision) Set(v string) error {
	switch v {
	case `DAY_PRECISION`, `MINUTE_PRECISION`, `SECOND_PRECISION`:
		*f = DatePrecision(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAY_PRECISION", "MINUTE_PRECISION", "SECOND_PRECISION"`, v)
	}
}

// Values returns all possible values for DatePrecision.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatePrecision) Values() []DatePrecision {
	return []DatePrecision{
		DatePrecisionDayPrecision,
		DatePrecisionMinutePrecision,
		DatePrecisionSecondPrecision,
	}
}

// Type always returns DatePrecision to satisfy [pflag.Value] interface
func (f *DatePrecision) Type() string {
	return "DatePrecision"
}

type DateRange struct {
	// End of the range, in the same format as `start`.
	End string `json:"end"`
	// Start of the range in the format `yyyy-MM-dd`, optionally followed by a
	// time as ` HH:mm` or ` HH:mm:ss`.
	Start string `json:"start"`
}

func (s *DateRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DateRangeValue struct {
	// Manually specified date-time range value.
	DateRangeValue *DateRange `json:"date_range_value,omitempty"`
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue DateRangeValueDynamicDateRange `json:"dynamic_date_range_value,omitempty"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision DatePrecision `json:"precision,omitempty"`
	// Day the week starts on for week-based dynamic ranges such as `THIS_WEEK`
	// and `LAST_WEEK`, from 0 (Sunday) to 6 (Saturday). Defaults to 0.
	StartDayOfWeek int `json:"start_day_of_week,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DateRangeValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DateRangeValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DateRangeValueDynamicDateRange string

const DateRangeValueDynamicDateRangeLast12Months DateRangeValueDynamicDateRange = `LAST_12_MONTHS`

const DateRangeValueDynamicDateRangeLast14Days DateRangeValueDynamicDateRange = `LAST_14_DAYS`

const DateRangeValueDynamicDateRangeLast24Hours DateRangeValueDynamicDateRange = `LAST_24_HOURS`

const DateRangeValueDynamicDateRangeLast30Days DateRangeValueDynamicDateRange = `LAST_30_DAYS`

const DateRangeValueDynamicDateRangeLast60Days DateRangeValueDynamicDateRange = `LAST_60_DAYS`

const DateRangeValueDynamicDateRangeLast7Days DateRangeValueDynamicDateRange = `LAST_7_DAYS`

const DateRangeValueDynamicDateRangeLast8Hours DateRangeValueDynamicDateRange = `LAST_8_HOURS`

const DateRangeValueDynamicDateRangeLast90Days DateRangeValueDynamicDateRange = `LAST_90_DAYS`

const DateRangeValueDynamicDateRangeLastHour DateRangeValueDynamicDateRange = `LAST_HOUR`

const DateRangeValueDynamicDateRangeLastMonth DateRangeValueDynamicDateRange = `LAST_MONTH`

const DateRangeValueDynamicDateRangeLastWeek DateRangeValueDynamicDateRange = `LAST_WEEK`

const DateRangeValueDynamicDateRangeLastYear DateRangeValueDynamicDateRange = `LAST_YEAR`

const DateRangeValueDynamicDateRangeThisMonth DateRangeValueDynamicDateRange = `THIS_MONTH`

const DateRangeValueDynamicDateRangeThisWeek DateRangeValueDynamicDateRange = `THIS_WEEK`

const DateRangeValueDynamicDateRangeThisYear DateRangeValueDynamicDateRange = `THIS_YEAR`

const DateRangeValueDynamicDateRangeToday DateRangeValueDynamicDateRange = `TODAY`

const DateRangeValueDynamicDateRangeYesterday DateRangeValueDynamicDateRange = `YESTERDAY`

// String representation for [fmt.Print]
func (f *DateRangeValueDynamicDateRange) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DateRangeValueDynamicDateRange) Set(v string) error {
	switch v {
	case `LAST_12_MONTHS`, `LAST_14_DAYS`, `LAST_24_HOURS`, `LAST_30_DAYS`, `LAST_60_DAYS`, `LAST_7_DAYS`, `LAST_8_HOURS`, `LAST_90_DAYS`, `LAST_HOUR`, `LAST_MONTH`, `LAST_WEEK`, `LAST_YEAR`, `THIS_MONTH`, `THIS_WEEK`, `THIS_YEAR`, `TODAY`, `YESTERDAY`:
		*f = DateRangeValueDynamicDateRange(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LAST_12_MONTHS", "LAST_14_DAYS", "LAST_24_HOURS", "LAST_30_DAYS", "LAST_60_DAYS", "LAST_7_DAYS", "LAST_8_HOURS", "LAST_90_DAYS", "LAST_HOUR", "LAST_MONTH", "LAST_WEEK", "LAST_YEAR", "THIS_MONTH", "THIS_WEEK", "THIS_YEAR", "TODAY", "YESTERDAY"`, v)
	}
}

// Values returns all possible values for DateRangeValueDynamicDateRange.
//
// There is no guarantee on the order of the values in the slice.
func (f *DateRangeValueDynamicDateRange) Values() []DateRangeValueDynamicDateRange {
	return []DateRangeValueDynamicDateRange{
		DateRangeValueDynamicDateRangeLast12Months,
		DateRangeValueDynamicDateRangeLast14Days,
		DateRangeValueDynamicDateRangeLast24Hours,
		DateRangeValueDynamicDateRangeLast30Days,
		DateRangeValueDynamicDateRangeLast60Days,
		DateRangeValueDynamicDateRangeLast7Days,
		DateRangeValueDynamicDateRangeLast8Hours,
		DateRangeValueDynamicDateRangeLast90Days,
		DateRangeValueDynamicDateRangeLastHour,
		DateRangeValueDynamicDateRangeLastMonth,
		DateRangeValueDynamicDateRangeLastWeek,
		DateRangeValueDynamicDateRangeLastYear,
		DateRangeValueDynamicDateRangeThisMonth,
		DateRangeValueDynamicDateRangeThisWeek,
		DateRangeValueDynamicDateRangeThisYear,
		DateRangeValueDynamicDateRangeToday,
		DateRangeValueDynamicDateRangeYesterday,
	}
}

// Type always returns DateRangeValueDynamicDateRange to satisfy [pflag.Value] interface
func (f *DateRangeValueDynamicDateRange) Type() string {
	return "DateRangeValueDynamicDateRange"
}

type DateValue struct {
	// Manually specified date-time value in the format `yyyy-MM-dd`, optionally
	// followed by a time as ` HH:mm` or ` HH:mm:ss`.
	DateValue string `json:"date_value,omitempty"`
	// Dynamic date-time value based on current date-time.
	DynamicDateValue DateValueDynamicDate `json:"dynamic_date_value,omitempty"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision DatePrecision `json:"precision,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DateValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DateValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DateValueDynamicDate string

const DateValueDynamicDateNow DateValueDynamicDate = `NOW`

const DateValueDynamicDateYesterday DateValueDynamicDate = `YESTERDAY`

// String representation for [fmt.Print]
func (f *DateValueDynamicDate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DateValueDynamicDate) Set(v string) error {
	switch v {
	case `NOW`, `YESTERDAY`:
		*f = DateValueDynamicDate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NOW", "YESTERDAY"`, v)
	}
}

// Values returns all possible values for DateValueDynamicDate.
//
// There is no guarantee on the order of the values in the slice.
func (f *DateValueDynamicDate) Values() []DateValueDynamicDate {
	return []DateValueDynamicDate{
		DateValueDynamicDateNow,
		DateValueDynamicDateYesterday,
	}
}

// Type always returns DateValueDynamicDate to satisfy [pflag.Value] interface
func (f *DateValueDynamicDate) Type() string {
	return "DateValueDynamicDate"
}

// Represents a per-user default warehouse override configuration. This resource
// allows users or administrators to customize how a user's default warehouse is
// selected for SQL operations. If no override exists for a user, the workspace
// default warehouse will be used.
type DefaultWarehouseOverride struct {
	// The ID component of the resource name (user ID).
	DefaultWarehouseOverrideId string `json:"default_warehouse_override_id,omitempty"`
	// The resource name of the default warehouse override. Format:
	// default-warehouse-overrides/{default_warehouse_override_id}
	Name string `json:"name,omitempty"`
	// The type of override behavior.
	Type DefaultWarehouseOverrideType `json:"type"`
	// The specific warehouse ID when type is CUSTOM. Not set for LAST_SELECTED
	// type.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DefaultWarehouseOverride) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultWarehouseOverride) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Type of default warehouse override behavior.
type DefaultWarehouseOverrideType string

const DefaultWarehouseOverrideTypeCustom DefaultWarehouseOverrideType = `CUSTOM`

const DefaultWarehouseOverrideTypeLastSelected DefaultWarehouseOverrideType = `LAST_SELECTED`

// String representation for [fmt.Print]
func (f *DefaultWarehouseOverrideType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DefaultWarehouseOverrideType) Set(v string) error {
	switch v {
	case `CUSTOM`, `LAST_SELECTED`:
		*f = DefaultWarehouseOverrideType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUSTOM", "LAST_SELECTED"`, v)
	}
}

// Values returns all possible values for DefaultWarehouseOverrideType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DefaultWarehouseOverrideType) Values() []DefaultWarehouseOverrideType {
	return []DefaultWarehouseOverrideType{
		DefaultWarehouseOverrideTypeCustom,
		DefaultWarehouseOverrideTypeLastSelected,
	}
}

// Type always returns DefaultWarehouseOverrideType to satisfy [pflag.Value] interface
func (f *DefaultWarehouseOverrideType) Type() string {
	return "DefaultWarehouseOverrideType"
}

type DeleteAlertsLegacyRequest struct {
	AlertId string `json:"-" url:"-"`
}

func (s *DeleteAlertsLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDashboardRequest struct {
	DashboardId string `json:"-" url:"-"`
}

func (s *DeleteDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id string `json:"-" url:"-"`
}

func (s *DeleteDashboardWidgetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDefaultWarehouseOverrideRequest struct {
	// Required. The resource name of the default warehouse override to delete.
	// Format: default-warehouse-overrides/{default_warehouse_override_id} The
	// default_warehouse_override_id can be a numeric user ID or the literal
	// string "me" for the current user.
	Name string `json:"-" url:"-"`
}

func (s *DeleteDefaultWarehouseOverrideRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteQueriesLegacyRequest struct {
	QueryId string `json:"-" url:"-"`
}

func (s *DeleteQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvisualizations/create
	Id string `json:"-" url:"-"`
}

func (s *DeleteQueryVisualizationsLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteVisualizationRequest struct {
	Id string `json:"-" url:"-"`
}

func (s *DeleteVisualizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func (s *DeleteWarehouseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

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

// Values returns all possible values for Disposition.
//
// There is no guarantee on the order of the values in the slice.
func (f *Disposition) Values() []Disposition {
	return []Disposition{
		DispositionExternalLinks,
		DispositionInline,
	}
}

// Type always returns Disposition to satisfy [pflag.Value] interface
func (f *Disposition) Type() string {
	return "Disposition"
}

type EditAlert struct {
	AlertId string `json:"-" url:"-"`
	// Name of the alert.
	Name string `json:"name"`
	// Alert configuration options.
	Options AlertOptions `json:"options"`
	// Query ID.
	QueryId string `json:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This is an incremental edit functionality, so all fields except id are
// optional. If a field is set, the corresponding configuration in the SQL
// warehouse is modified. If a field is unset, the existing configuration value
// in the SQL warehouse is retained. Thus, this API is not idempotent.
type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 1 min for serverless warehouses. -
	// Must be == 0 or >= 10 mins for classic and pro warehouses. - 0 indicates
	// no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel that determines which Databricks SQL version the warehouse runs.
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large - 5X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// Name of the user who created the warehouse. Set by the server and ignored
	// in requests.
	CreatorName string `json:"creator_name,omitempty"`
	// Compute version label for a Real-Time warehouse. This field is output
	// only.
	CustomerVersionLabel string `json:"customer_version_label,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to true.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Required. Id of the warehouse to configure.
	Id string `json:"-" url:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum autoscaling capacity of a Real-Time SQL warehouse, in DBU per
	// hour. Must be at least the minimum DBU for the selected warehouse size
	// and no greater than the configured DBU limit. Actual capacity may be
	// below this ceiling.
	MaxDbu int `json:"max_dbu,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the warehouse.
	//
	// Supported values: - Must be unique within an org. - Must be 1 to 100
	// characters.
	Name string `json:"name,omitempty"`
	// Cluster size multiplier for Reyden warehouses. Scales the effective
	// compute capacity relative to the nominal virtual cluster size. Only
	// meaningful when the warehouse is of Reyden type.
	//
	// Accepted range: 1.0 to the workspace-configured maximum (decimals
	// allowed). The default maximum is 100.0. - Values below 1.0 are normalized
	// to 1.0. - Values above the configured maximum are capped at that maximum.
	// - When absent, the warehouse is treated as having a multiplier of 1.0.
	ReydenClusterSizeMultiplier float64 `json:"reyden_cluster_size_multiplier,omitempty"`
	// Configures whether the warehouse should use spot instances.
	// COST_OPTIMIZED uses spot instances where the cloud provider supports
	// them, and RELIABILITY_OPTIMIZED uses on-demand instances only.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// Maximum time in seconds that a statement can run on this warehouse before
	// it is cancelled. Set to 0 to remove the warehouse-level timeout. Changing
	// the value of a running warehouse restarts it. See
	// https://docs.databricks.com/aws/en/sql/language-manual/parameters/statement_timeout.
	StatementTimeout int `json:"statement_timeout,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type: `PRO`, `CLASSIC`, or `REALTIME`. For serverless compute,
	// use `PRO` or `REALTIME` with `enable_serverless_compute=true`. `REALTIME`
	// requires serverless compute and Photon; `enable_photon` defaults to
	// `true`.
	WarehouseType EditWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditWarehouseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditWarehouseRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EditWarehouseRequestWarehouseType string

const EditWarehouseRequestWarehouseTypeClassic EditWarehouseRequestWarehouseType = `CLASSIC`

const EditWarehouseRequestWarehouseTypePro EditWarehouseRequestWarehouseType = `PRO`

const EditWarehouseRequestWarehouseTypeRealtime EditWarehouseRequestWarehouseType = `REALTIME`

const EditWarehouseRequestWarehouseTypeTypeUnspecified EditWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *EditWarehouseRequestWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EditWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `REALTIME`, `TYPE_UNSPECIFIED`:
		*f = EditWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "REALTIME", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for EditWarehouseRequestWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EditWarehouseRequestWarehouseType) Values() []EditWarehouseRequestWarehouseType {
	return []EditWarehouseRequestWarehouseType{
		EditWarehouseRequestWarehouseTypeClassic,
		EditWarehouseRequestWarehouseTypePro,
		EditWarehouseRequestWarehouseTypeRealtime,
		EditWarehouseRequestWarehouseTypeTypeUnspecified,
	}
}

// Type always returns EditWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *EditWarehouseRequestWarehouseType) Type() string {
	return "EditWarehouseRequestWarehouseType"
}

type EndpointConfPair struct {
	// Name of the configuration parameter.
	Key string `json:"key,omitempty"`
	// Value of the configuration parameter.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointConfPair) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointConfPair) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	Details string `json:"details,omitempty"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason *TerminationReason `json:"failure_reason,omitempty"`
	// Deprecated. split into summary and details for security
	Message string `json:"message,omitempty"`
	// Health status of the endpoint.
	Status Status `json:"status,omitempty"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary string `json:"summary,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointHealth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointHealth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 1 min for serverless warehouses. -
	// Must be == 0 or >= 10 mins for classic and pro warehouses. - 0 indicates
	// no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel that determines which Databricks SQL version the warehouse runs.
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large - 5X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// Name of the user who created the warehouse. Set by the server and ignored
	// in requests.
	CreatorName string `json:"creator_name,omitempty"`
	// Compute version label for a Real-Time warehouse. This field is output
	// only.
	CustomerVersionLabel string `json:"customer_version_label,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to true.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *EndpointHealth `json:"health,omitempty"`
	// Unique identifier of the warehouse.
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// JDBC connection string for the warehouse.
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum autoscaling capacity of a Real-Time SQL warehouse, in DBU per
	// hour. Must be at least the minimum DBU for the selected warehouse size
	// and no greater than the configured DBU limit. Actual capacity may be
	// below this ceiling.
	MaxDbu int `json:"max_dbu,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the warehouse.
	//
	// Supported values: - Must be unique within an org. - Must be 1 to 100
	// characters.
	Name string `json:"name,omitempty"`
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// Number of clusters currently allocated to the warehouse, including
	// clusters that are still starting.
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the SQL warehouse
	OdbcParams *OdbcParams `json:"odbc_params,omitempty"`
	// Cluster size multiplier for Reyden warehouses. Scales the effective
	// compute capacity relative to the nominal virtual cluster size. Only
	// meaningful when the warehouse is of Reyden type.
	//
	// Accepted range: 1.0 to the workspace-configured maximum (decimals
	// allowed). The default maximum is 100.0. - Values below 1.0 are normalized
	// to 1.0. - Values above the configured maximum are capped at that maximum.
	// - When absent, the warehouse is treated as having a multiplier of 1.0.
	ReydenClusterSizeMultiplier float64 `json:"reyden_cluster_size_multiplier,omitempty"`
	// Configures whether the warehouse should use spot instances.
	// COST_OPTIMIZED uses spot instances where the cloud provider supports
	// them, and RELIABILITY_OPTIMIZED uses on-demand instances only.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// Current state of the warehouse.
	State State `json:"state,omitempty"`
	// Maximum time in seconds that a statement can run on this warehouse before
	// it is cancelled. Set to 0 to remove the warehouse-level timeout. Changing
	// the value of a running warehouse restarts it. See
	// https://docs.databricks.com/aws/en/sql/language-manual/parameters/statement_timeout.
	StatementTimeout int `json:"statement_timeout,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type: `PRO`, `CLASSIC`, or `REALTIME`. For serverless compute,
	// use `PRO` or `REALTIME` with `enable_serverless_compute=true`. `REALTIME`
	// requires serverless compute and Photon; `enable_photon` defaults to
	// `true`.
	WarehouseType EndpointInfoWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointInfoWarehouseType string

const EndpointInfoWarehouseTypeClassic EndpointInfoWarehouseType = `CLASSIC`

const EndpointInfoWarehouseTypePro EndpointInfoWarehouseType = `PRO`

const EndpointInfoWarehouseTypeRealtime EndpointInfoWarehouseType = `REALTIME`

const EndpointInfoWarehouseTypeTypeUnspecified EndpointInfoWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *EndpointInfoWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointInfoWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `REALTIME`, `TYPE_UNSPECIFIED`:
		*f = EndpointInfoWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "REALTIME", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for EndpointInfoWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointInfoWarehouseType) Values() []EndpointInfoWarehouseType {
	return []EndpointInfoWarehouseType{
		EndpointInfoWarehouseTypeClassic,
		EndpointInfoWarehouseTypePro,
		EndpointInfoWarehouseTypeRealtime,
		EndpointInfoWarehouseTypeTypeUnspecified,
	}
}

// Type always returns EndpointInfoWarehouseType to satisfy [pflag.Value] interface
func (f *EndpointInfoWarehouseType) Type() string {
	return "EndpointInfoWarehouseType"
}

type EndpointTagPair struct {
	// Tag key. Keys must be unique within the warehouse and cannot use the keys
	// that Databricks reserves for the tags it sets automatically.
	Key string `json:"key,omitempty"`
	// Tag value. Keys and values must follow the tagging rules of the cloud
	// provider.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EndpointTagPair) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTagPair) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointTags struct {
	// User-defined tags, applied in addition to the tags Databricks sets
	// automatically.
	CustomTags []EndpointTagPair `json:"custom_tags,omitempty"`
}

func (s *EndpointTags) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type EnumValue struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions string `json:"enum_options,omitempty"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions *MultiValuesOptions `json:"multi_values_options,omitempty"`
	// List of selected query parameter values.
	Values []string `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnumValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnumValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explicitly set.
	ByteLimit int64 `json:"byte_limit,omitempty"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog string `json:"catalog,omitempty"`
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
	// 2. These are short-lived cloud-storage URLs with a specific expiration,
	// indicated in the response. The behavior when attempting to use an expired
	// link is cloud specific.
	Disposition Disposition `json:"disposition,omitempty"`
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
	Format Format `json:"format,omitempty"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeout `json:"on_wait_timeout,omitempty"`
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
	// ``` SELECT * FROM my_table WHERE name = :my_name AND date = :my_date ```
	//
	// The parameters can be passed in the request body as follows:
	//
	// ``` { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	// ```
	//
	// Currently, positional parameters denoted by a `?` marker are not
	// supported by the Databricks SQL Statement Execution API.
	//
	// Also see the section [Parameter markers] of the SQL language reference.
	//
	// [Parameter markers]: https://docs.databricks.com/sql/language-manual/sql-ref-parameter-marker.html
	// [`cast` function]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Parameters []StatementParameterListItem `json:"parameters,omitempty"`
	// An array of query tags to annotate a SQL statement. A query tag consists
	// of a non-empty key and, optionally, a value. To represent a NULL value,
	// either omit the `value` field or manually set it to `null` or white
	// space. Refer to the SQL language reference for the format specification
	// of query tags. There's no significance to the order of tags. Only one
	// value per key will be recorded. A sequence in excess of 20 query tags
	// will be coerced to 20. Example:
	//
	// { ..., "query_tags": [ { "key": "team", "value": "eng" }, { "key": "some
	// key only tag" } ] }
	QueryTags []QueryTag `json:"query_tags,omitempty"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit int64 `json:"row_limit,omitempty"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema string `json:"schema,omitempty"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`. The maximum query text size is 16 MiB.
	Statement string `json:"statement"`
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
	WaitTimeout string `json:"wait_timeout,omitempty"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExecuteStatementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteStatementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ExecuteStatementRequestOnWaitTimeout.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExecuteStatementRequestOnWaitTimeout) Values() []ExecuteStatementRequestOnWaitTimeout {
	return []ExecuteStatementRequestOnWaitTimeout{
		ExecuteStatementRequestOnWaitTimeoutCancel,
		ExecuteStatementRequestOnWaitTimeoutContinue,
	}
}

// Type always returns ExecuteStatementRequestOnWaitTimeout to satisfy [pflag.Value] interface
func (f *ExecuteStatementRequestOnWaitTimeout) Type() string {
	return "ExecuteStatementRequestOnWaitTimeout"
}

type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration string `json:"expiration,omitempty"`
	// A short-lived cloud-storage URL pointing to a chunk of result data,
	// hosted by an external service, with a short expiration time (<= 15
	// minutes). As this URL contains a temporary credential, it should be
	// considered sensitive and the client should not expose this URL in a log.
	ExternalLink string `json:"external_link,omitempty"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders map[string]string `json:"http_headers,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalLink) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLink) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExternalQuerySource struct {
	// The canonical identifier for this SQL alert
	AlertId string `json:"alert_id,omitempty"`
	// The canonical identifier for this Lakeview dashboard
	DashboardId string `json:"dashboard_id,omitempty"`
	// The canonical identifier for this Genie space
	GenieSpaceId string `json:"genie_space_id,omitempty"`
	// Identifiers of the job, job run, and task run that executed this
	// statement, when it ran as part of a job.
	JobInfo *ExternalQuerySourceJobInfo `json:"job_info,omitempty"`
	// The canonical identifier for this legacy dashboard
	LegacyDashboardId string `json:"legacy_dashboard_id,omitempty"`
	// The canonical identifier for this notebook
	NotebookId string `json:"notebook_id,omitempty"`
	// The canonical identifier for this SQL query
	SqlQueryId string `json:"sql_query_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalQuerySource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalQuerySource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExternalQuerySourceJobInfo struct {
	// The canonical identifier for this job.
	JobId string `json:"job_id,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	JobRunId string `json:"job_run_id,omitempty"`
	// The canonical identifier of the task run.
	JobTaskRunId string `json:"job_task_run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalQuerySourceJobInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalQuerySourceJobInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for Format.
//
// There is no guarantee on the order of the values in the slice.
func (f *Format) Values() []Format {
	return []Format{
		FormatArrowStream,
		FormatCsv,
		FormatJsonArray,
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

type GetAlertRequest struct {
	// UUID of the alert to fetch.
	Id string `json:"-" url:"-"`
}

func (s *GetAlertRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetAlertV2Request struct {
	// ID of the alert to retrieve, as returned in the alert's `id` field.
	Id string `json:"-" url:"-"`
}

func (s *GetAlertV2Request) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetAlertsLegacyRequest struct {
	AlertId string `json:"-" url:"-"`
}

func (s *GetAlertsLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDashboardRequest struct {
	DashboardId string `json:"-" url:"-"`
}

func (s *GetDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId string `json:"-" url:"-"`
	// The type of object permissions to check.
	ObjectType ObjectTypePlural `json:"-" url:"-"`
}

func (s *GetDbsqlPermissionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDefaultWarehouseOverrideRequest struct {
	// Required. The resource name of the default warehouse override to
	// retrieve. Format:
	// default-warehouse-overrides/{default_warehouse_override_id} The
	// default_warehouse_override_id can be a numeric user ID or the literal
	// string "me" for the current user.
	Name string `json:"-" url:"-"`
}

func (s *GetDefaultWarehouseOverrideRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetQueriesLegacyRequest struct {
	QueryId string `json:"-" url:"-"`
}

func (s *GetQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetQueryRequest struct {
	// UUID of the query to fetch.
	Id string `json:"-" url:"-"`
}

func (s *GetQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetResponse struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId string `json:"object_id,omitempty"`
	// A singular noun object type.
	ObjectType ObjectType `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
}

func (s *GetStatementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetStatementResultChunkNRequest struct {
	// The zero-based index of the result chunk to fetch.
	ChunkIndex int `json:"-" url:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
}

func (s *GetStatementResultChunkNRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" url:"-"`
}

func (s *GetWarehousePermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []WarehousePermissionsDescription `json:"permission_levels,omitempty"`
}

func (s *GetWarehousePermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" url:"-"`
}

func (s *GetWarehousePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func (s *GetWarehouseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 1 min for serverless warehouses. -
	// Must be == 0 or >= 10 mins for classic and pro warehouses. - 0 indicates
	// no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel that determines which Databricks SQL version the warehouse runs.
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large - 5X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// Name of the user who created the warehouse. Set by the server and ignored
	// in requests.
	CreatorName string `json:"creator_name,omitempty"`
	// Compute version label for a Real-Time warehouse. This field is output
	// only.
	CustomerVersionLabel string `json:"customer_version_label,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to true.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *EndpointHealth `json:"health,omitempty"`
	// Unique identifier of the warehouse.
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// JDBC connection string for the warehouse.
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum autoscaling capacity of a Real-Time SQL warehouse, in DBU per
	// hour. Must be at least the minimum DBU for the selected warehouse size
	// and no greater than the configured DBU limit. Actual capacity may be
	// below this ceiling.
	MaxDbu int `json:"max_dbu,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 40.
	//
	// Defaults to min_clusters if unset.
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the warehouse.
	//
	// Supported values: - Must be unique within an org. - Must be 1 to 100
	// characters.
	Name string `json:"name,omitempty"`
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// Number of clusters currently allocated to the warehouse, including
	// clusters that are still starting.
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the SQL warehouse
	OdbcParams *OdbcParams `json:"odbc_params,omitempty"`
	// Cluster size multiplier for Reyden warehouses. Scales the effective
	// compute capacity relative to the nominal virtual cluster size. Only
	// meaningful when the warehouse is of Reyden type.
	//
	// Accepted range: 1.0 to the workspace-configured maximum (decimals
	// allowed). The default maximum is 100.0. - Values below 1.0 are normalized
	// to 1.0. - Values above the configured maximum are capped at that maximum.
	// - When absent, the warehouse is treated as having a multiplier of 1.0.
	ReydenClusterSizeMultiplier float64 `json:"reyden_cluster_size_multiplier,omitempty"`
	// Configures whether the warehouse should use spot instances.
	// COST_OPTIMIZED uses spot instances where the cloud provider supports
	// them, and RELIABILITY_OPTIMIZED uses on-demand instances only.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// Current state of the warehouse.
	State State `json:"state,omitempty"`
	// Maximum time in seconds that a statement can run on this warehouse before
	// it is cancelled. Set to 0 to remove the warehouse-level timeout. Changing
	// the value of a running warehouse restarts it. See
	// https://docs.databricks.com/aws/en/sql/language-manual/parameters/statement_timeout.
	StatementTimeout int `json:"statement_timeout,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type: `PRO`, `CLASSIC`, or `REALTIME`. For serverless compute,
	// use `PRO` or `REALTIME` with `enable_serverless_compute=true`. `REALTIME`
	// requires serverless compute and Photon; `enable_photon` defaults to
	// `true`.
	WarehouseType GetWarehouseResponseWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetWarehouseResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetWarehouseResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetWarehouseResponseWarehouseType string

const GetWarehouseResponseWarehouseTypeClassic GetWarehouseResponseWarehouseType = `CLASSIC`

const GetWarehouseResponseWarehouseTypePro GetWarehouseResponseWarehouseType = `PRO`

const GetWarehouseResponseWarehouseTypeRealtime GetWarehouseResponseWarehouseType = `REALTIME`

const GetWarehouseResponseWarehouseTypeTypeUnspecified GetWarehouseResponseWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *GetWarehouseResponseWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetWarehouseResponseWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `REALTIME`, `TYPE_UNSPECIFIED`:
		*f = GetWarehouseResponseWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "REALTIME", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for GetWarehouseResponseWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetWarehouseResponseWarehouseType) Values() []GetWarehouseResponseWarehouseType {
	return []GetWarehouseResponseWarehouseType{
		GetWarehouseResponseWarehouseTypeClassic,
		GetWarehouseResponseWarehouseTypePro,
		GetWarehouseResponseWarehouseTypeRealtime,
		GetWarehouseResponseWarehouseTypeTypeUnspecified,
	}
}

// Type always returns GetWarehouseResponseWarehouseType to satisfy [pflag.Value] interface
func (f *GetWarehouseResponseWarehouseType) Type() string {
	return "GetWarehouseResponseWarehouseType"
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	Channel *Channel `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs `json:"config_param,omitempty"`
	// Spark configuration applied to every warehouse in the workspace, for
	// example to connect to an external Hive metastore or to configure access
	// to cloud storage. Only allowlisted keys are accepted, and the
	// JSON-serialized size must be at most 512 KB. See
	// https://docs.databricks.com/aws/en/admin/sql/data-access-configuration.
	DataAccessConfig []EndpointConfPair `json:"data_access_config,omitempty"`
	// Deprecated: only setting this to true is allowed.
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type.
	EnabledWarehouseTypes []WarehouseTypePair `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs `json:"global_param,omitempty"`
	// GCP only: Google service account that the warehouse clusters use to
	// access Google Cloud Storage.
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: The instance profile used to pass an IAM role to the SQL
	// warehouses. This configuration is also applied to the workspace's
	// serverless compute for notebooks and jobs.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for warehouses
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters applied to every warehouse in the workspace.
	// Only parameters that can be set at the workspace scope are accepted. See
	// https://docs.databricks.com/aws/en/sql/language-manual/sql-ref-parameters.
	SqlConfigurationParameters *RepeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetWorkspaceWarehouseConfigResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetWorkspaceWarehouseConfigResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Security policy to be used for warehouses
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

// Values returns all possible values for GetWorkspaceWarehouseConfigResponseSecurityPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Values() []GetWorkspaceWarehouseConfigResponseSecurityPolicy {
	return []GetWorkspaceWarehouseConfigResponseSecurityPolicy{
		GetWorkspaceWarehouseConfigResponseSecurityPolicyDataAccessControl,
		GetWorkspaceWarehouseConfigResponseSecurityPolicyNone,
		GetWorkspaceWarehouseConfigResponseSecurityPolicyPassthrough,
	}
}

// Type always returns GetWorkspaceWarehouseConfigResponseSecurityPolicy to satisfy [pflag.Value] interface
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Type() string {
	return "GetWorkspaceWarehouseConfigResponseSecurityPolicy"
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Alert ID.
	Id string `json:"id,omitempty"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt string `json:"last_triggered_at,omitempty"`
	// Name of the alert.
	Name string `json:"name,omitempty"`
	// Alert configuration options.
	Options *AlertOptions `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`

	Query *AlertQuery `json:"query,omitempty"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int `json:"rearm,omitempty"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State LegacyAlertState `json:"state,omitempty"`
	// Timestamp when the alert was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *User `json:"user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LegacyAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LegacyAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type LegacyAlertState string

const LegacyAlertStateOk LegacyAlertState = `ok`

const LegacyAlertStateTriggered LegacyAlertState = `triggered`

const LegacyAlertStateUnknown LegacyAlertState = `unknown`

// String representation for [fmt.Print]
func (f *LegacyAlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LegacyAlertState) Set(v string) error {
	switch v {
	case `ok`, `triggered`, `unknown`:
		*f = LegacyAlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ok", "triggered", "unknown"`, v)
	}
}

// Values returns all possible values for LegacyAlertState.
//
// There is no guarantee on the order of the values in the slice.
func (f *LegacyAlertState) Values() []LegacyAlertState {
	return []LegacyAlertState{
		LegacyAlertStateOk,
		LegacyAlertStateTriggered,
		LegacyAlertStateUnknown,
	}
}

// Type always returns LegacyAlertState to satisfy [pflag.Value] interface
func (f *LegacyAlertState) Type() string {
	return "LegacyAlertState"
}

type LegacyQuery struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit bool `json:"can_edit,omitempty"`
	// The timestamp when this query was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Query ID.
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool `json:"is_draft,omitempty"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool `json:"is_safe,omitempty"`

	LastModifiedBy *User `json:"last_modified_by,omitempty"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById int `json:"last_modified_by_id,omitempty"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId string `json:"latest_query_data_id,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`

	Options *QueryOptions `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash string `json:"query_hash,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *User `json:"user,omitempty"`
	// The ID of the user who owns the query.
	UserId int `json:"user_id,omitempty"`

	Visualizations []LegacyVisualization `json:"visualizations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LegacyQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LegacyQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {
	CreatedAt string `json:"created_at,omitempty"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description string `json:"description,omitempty"`
	// The UUID for this visualization.
	Id string `json:"id,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `json:"options,omitempty"`

	Query *LegacyQuery `json:"query,omitempty"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type string `json:"type,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LegacyVisualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LegacyVisualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

type ListAlertsRequest struct {
	// Maximum number of alerts to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token from the `next_page_token` of a previous response. Omit it to fetch
	// the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAlertsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAlertsResponse struct {
	// Token to pass as `page_token` to fetch the next page. Omitted when there
	// are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Alerts on this page.
	Results []ListAlertsResponseAlert `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAlertsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A legacy Databricks SQL alert. An alert runs a saved query, compares a column
// of the result against a threshold, and notifies its subscribers when the
// condition is met. See
// https://docs.databricks.com/aws/en/sql/user/alerts/legacy for details.
type ListAlertsResponseAlert struct {
	// Condition that triggers the alert: a column of the query result, a
	// comparison operator, and a threshold value. The value in the first row of
	// the column is evaluated; aggregations are not supported through this API.
	Condition *AlertCondition `json:"condition,omitempty"`
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom body of the alert notification. It can use the template variables
	// and HTML described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default notification body is used.
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of the alert notification, used as the email subject or
	// the Slack message header. It can use the template variables described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default subject is used.
	CustomSubject string `json:"custom_subject,omitempty"`
	// Display name of the alert, shown in list views and notifications. It must
	// not be empty.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the alert.
	Id string `json:"id,omitempty"`
	// Lifecycle state of the alert: `ACTIVE`, or `TRASHED` after it has been
	// moved to the trash.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Whether to notify subscribers when the alert state changes from
	// `TRIGGERED` back to `OK`. Defaults to true.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// UUID of the saved query whose result the alert evaluates. The caller must
	// have permission to view the query.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds the alert waits after a notification before it can
	// notify again while its state stays `TRIGGERED`. If 0 or not specified,
	// the alert notifies once when its state changes to `TRIGGERED` and does
	// not notify again until its state has changed.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`
	// Current state of the alert's trigger status: `TRIGGERED` if the condition
	// was met on the last evaluation, `OK` if it was not, and `UNKNOWN` if the
	// alert has not yet been evaluated or ran into an error during the last
	// evaluation.
	State AlertState `json:"state,omitempty"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime string `json:"trigger_time,omitempty"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAlertsResponseAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsResponseAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAlertsV2Request struct {
	// Maximum number of alerts to return per page. Values above 100 are coerced
	// to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Page token returned as `next_page_token` by a previous call. Omit it to
	// retrieve the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAlertsV2Request) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsV2Request) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAlertsV2Response struct {
	Alerts []AlertV2 `json:"alerts,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAlertsV2Response) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsV2Response) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order ListOrder `json:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page int `json:"-" url:"page,omitempty"`
	// Number of dashboards to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Full text search term.
	Q string `json:"-" url:"q,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDefaultWarehouseOverridesRequest struct {
	// The maximum number of overrides to return. The service may return fewer
	// than this value. If unspecified, at most 100 overrides will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListDefaultWarehouseOverrides`
	// call. Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListDefaultWarehouseOverrides` must match the call that provided the
	// page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDefaultWarehouseOverridesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDefaultWarehouseOverridesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for ListDefaultWarehouseOverrides.
type ListDefaultWarehouseOverridesResponse struct {
	// The default warehouse overrides in the workspace.
	DefaultWarehouseOverrides []DefaultWarehouseOverride `json:"default_warehouse_overrides,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDefaultWarehouseOverridesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDefaultWarehouseOverridesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ListOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListOrder) Values() []ListOrder {
	return []ListOrder{
		ListOrderCreatedAt,
		ListOrderName,
	}
}

// Type always returns ListOrder to satisfy [pflag.Value] interface
func (f *ListOrder) Type() string {
	return "ListOrder"
}

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
	Order string `json:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page int `json:"-" url:"page,omitempty"`
	// Number of queries to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Full text search term
	Q string `json:"-" url:"q,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueriesLegacyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQueriesRequest struct {
	// Maximum number of queries to return per page. Values above 100 are
	// coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Token from the `next_page_token` of a previous response. Omit it to fetch
	// the first page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQueriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	HasNextPage bool `json:"has_next_page,omitempty"`
	// A token that can be used to get the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Res []QueryInfo `json:"res,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQueriesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueriesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQueryHistoryRequest struct {
	// An optional filter object to limit query history results. Accepts
	// parameters such as user IDs, warehouse IDs, and statuses to narrow the
	// returned data. In a URL, the parameters of this filter are specified with
	// dot notation. For example: `filter_by.statement_ids`.
	FilterBy *QueryFilter `json:"-" url:"filter_by,omitempty"`
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	IncludeMetrics bool `json:"-" url:"include_metrics,omitempty"`
	// Limit the number of results returned in one page. Must be no more than
	// 1000; the default is 100.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results. The token can
	// contain characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQueryHistoryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueryHistoryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQueryObjectsResponse struct {
	// Token to pass as `page_token` to fetch the next page. Omitted when there
	// are no more pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Queries on this page.
	Results []ListQueryObjectsResponseQuery `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQueryObjectsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueryObjectsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A saved Databricks SQL query. A query stores the SQL text, the SQL warehouse
// it runs on, its parameters, and metadata such as the display name,
// description, owner, and tags.
type ListQueryObjectsResponseQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// Timestamp when this query was created.
	CreateTime string `json:"create_time,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the query.
	Id string `json:"id,omitempty"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName string `json:"last_modifier_user_name,omitempty"`
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Username of the user that owns the query.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Whose credentials the query runs with when it is shared. `VIEWER` uses
	// the credentials of the user running the query. `OWNER` uses the owner's
	// credentials where the product supports that, such as for legacy alerts
	// and jobs. When omitted on create, the workspace default applies. A
	// workspace admin can disallow `OWNER`.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`
	// Tags attached to the query. Tags appear in the query list and can be used
	// to filter it.
	Tags []string `json:"tags,omitempty"`
	// Timestamp when this query was last updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQueryObjectsResponseQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueryObjectsResponseQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListResponse struct {
	// The total number of dashboards.
	Count int `json:"count,omitempty"`
	// The current page being displayed.
	Page int `json:"page,omitempty"`
	// The number of dashboards per page.
	PageSize int `json:"page_size,omitempty"`
	// List of dashboards returned.
	Results []Dashboard `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVisualizationsForQueryRequest struct {
	Id string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVisualizationsForQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVisualizationsForQueryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListVisualizationsForQueryResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []Visualization `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListVisualizationsForQueryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVisualizationsForQueryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWarehousesRequest struct {
	// The max number of warehouses to return.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListWarehouses` call. Provide
	// this to retrieve the subsequent page; otherwise the first will be
	// retrieved.
	//
	// When paginating, all other parameters provided to `ListWarehouses` must
	// match the call that provided the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Deprecated: this field is ignored by the server. Service Principal which
	// will be used to fetch the list of endpoints. If not specified, SQL
	// Gateway will use the user from the session header.
	RunAsUserId int `json:"-" url:"run_as_user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWarehousesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWarehousesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWarehousesResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// A list of warehouses and their configurations.
	Warehouses []EndpointInfo `json:"warehouses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWarehousesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWarehousesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	Prefix string `json:"prefix,omitempty"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator string `json:"separator,omitempty"`
	// Character that suffixes each selected parameter value.
	Suffix string `json:"suffix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MultiValuesOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MultiValuesOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NumericValue struct {
	// Number substituted for the parameter marker when the query runs.
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NumericValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NumericValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ObjectType) Values() []ObjectType {
	return []ObjectType{
		ObjectTypeAlert,
		ObjectTypeDashboard,
		ObjectTypeDataSource,
		ObjectTypeQuery,
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

// Values returns all possible values for ObjectTypePlural.
//
// There is no guarantee on the order of the values in the slice.
func (f *ObjectTypePlural) Values() []ObjectTypePlural {
	return []ObjectTypePlural{
		ObjectTypePluralAlerts,
		ObjectTypePluralDashboards,
		ObjectTypePluralDataSources,
		ObjectTypePluralQueries,
	}
}

// Type always returns ObjectTypePlural to satisfy [pflag.Value] interface
func (f *ObjectTypePlural) Type() string {
	return "ObjectTypePlural"
}

type OdbcParams struct {
	// Hostname of the workspace to connect to.
	Hostname string `json:"hostname,omitempty"`
	// HTTP path of the warehouse.
	Path string `json:"path,omitempty"`
	// The connection port is always 443.
	Port int `json:"port,omitempty"`
	// The transport protocol is always https.
	Protocol string `json:"protocol,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OdbcParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OdbcParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for OwnableObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *OwnableObjectType) Values() []OwnableObjectType {
	return []OwnableObjectType{
		OwnableObjectTypeAlert,
		OwnableObjectTypeDashboard,
		OwnableObjectTypeQuery,
	}
}

// Type always returns OwnableObjectType to satisfy [pflag.Value] interface
func (f *OwnableObjectType) Type() string {
	return "OwnableObjectType"
}

type Parameter struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	EnumOptions string `json:"enumOptions,omitempty"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions *MultiValuesOptions `json:"multiValuesOptions,omitempty"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name string `json:"name,omitempty"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId string `json:"queryId,omitempty"`
	// The text displayed in a parameter picking widget.
	Title string `json:"title,omitempty"`
	// Parameters can have several different types.
	Type ParameterType `json:"type,omitempty"`
	// The default value for this parameter.
	Value any `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Parameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Parameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for ParameterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ParameterType) Values() []ParameterType {
	return []ParameterType{
		ParameterTypeDatetime,
		ParameterTypeEnum,
		ParameterTypeNumber,
		ParameterTypeQuery,
		ParameterTypeText,
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

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelCanEdit,
		PermissionLevelCanManage,
		PermissionLevelCanRun,
		PermissionLevelCanView,
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

// Possible Reasons for which we have not saved plans in the database
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

// Values returns all possible values for PlansState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PlansState) Values() []PlansState {
	return []PlansState{
		PlansStateEmpty,
		PlansStateExists,
		PlansStateIgnoredLargePlansSize,
		PlansStateIgnoredSmallDuration,
		PlansStateIgnoredSparkPlanType,
		PlansStateUnknown,
	}
}

// Type always returns PlansState to satisfy [pflag.Value] interface
func (f *PlansState) Type() string {
	return "PlansState"
}

// A saved Databricks SQL query. A query stores the SQL text, the SQL warehouse
// it runs on, its parameters, and metadata such as the display name,
// description, owner, and tags.
type Query struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// Timestamp when this query was created.
	CreateTime string `json:"create_time,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the query.
	Id string `json:"id,omitempty"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName string `json:"last_modifier_user_name,omitempty"`
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Username of the user that owns the query.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Workspace path of the workspace folder containing the object.
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Whose credentials the query runs with when it is shared. `VIEWER` uses
	// the credentials of the user running the query. `OWNER` uses the owner's
	// credentials where the product supports that, such as for legacy alerts
	// and jobs. When omitted on create, the workspace default applies. A
	// workspace admin can disallow `OWNER`.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`
	// Tags attached to the query. Tags appear in the query list and can be used
	// to filter it.
	Tags []string `json:"tags,omitempty"`
	// Timestamp when this query was last updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Query) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Query) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions *MultiValuesOptions `json:"multi_values_options,omitempty"`
	// UUID of the query that provides the parameter values.
	QueryId string `json:"query_id,omitempty"`
	// List of selected query parameter values.
	Values []string `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryBackedValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryBackedValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `json:"options,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`

	QueryId string `json:"-" url:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryEditContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEditContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be less than
	// or equal to 30 days.
	QueryStartTimeRange *TimeRange `json:"query_start_time_range,omitempty" url:"query_start_time_range,omitempty"`
	// A list of statement IDs.
	StatementIds []string `json:"statement_ids,omitempty" url:"statement_ids,omitempty"`
	// A list of statuses (QUEUED, RUNNING, CANCELED, FAILED, FINISHED) to match
	// query results. Corresponds to the `status` field in the response.
	// Filtering for multiple statuses is not recommended. Instead, opt to
	// filter by a single status multiple times and then combine the results.
	Statuses []QueryStatus `json:"statuses,omitempty" url:"statuses,omitempty"`
	// A list of user IDs who ran the queries.
	UserIds []int64 `json:"user_ids,omitempty" url:"user_ids,omitempty"`
	// A list of warehouse IDs.
	WarehouseIds []string `json:"warehouse_ids,omitempty" url:"warehouse_ids,omitempty"`
}

func (s *QueryFilter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type QueryInfo struct {
	// The ID of the query whose cached result was returned, if this result was
	// retrieved from the cache.
	CacheQueryId string `json:"cache_query_id,omitempty"`
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed *ChannelInfo `json:"channel_used,omitempty"`
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	ClientApplication string `json:"client_application,omitempty"`
	// Total time of the statement execution, in milliseconds. This value does
	// not include the time taken to retrieve the results, which can result in a
	// discrepancy between this value and the start-to-finish wall-clock time.
	Duration int64 `json:"duration,omitempty"`
	// Alias for `warehouse_id`.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Message describing why the query could not complete.
	ErrorMessage string `json:"error_message,omitempty"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId int64 `json:"executed_as_user_id,omitempty"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName string `json:"executed_as_user_name,omitempty"`
	// The time execution of the query ended, in milliseconds since the Unix
	// epoch.
	ExecutionEndTimeMs int64 `json:"execution_end_time_ms,omitempty"`
	// Whether the query has reached a final state, so no further updates to
	// this record are expected.
	IsFinal bool `json:"is_final,omitempty"`
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`
	// Metrics about query execution. Only returned when `include_metrics` is
	// true in the request.
	Metrics *QueryMetrics `json:"metrics,omitempty"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState PlansState `json:"plans_state,omitempty"`
	// The time the query ended, in milliseconds since the Unix epoch.
	QueryEndTimeMs int64 `json:"query_end_time_ms,omitempty"`
	// The UUID that uniquely identifies the query execution. It matches the
	// statement ID shown in the Query History UI and the values accepted by the
	// `statement_ids` filter.
	QueryId string `json:"query_id,omitempty"`
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	QuerySource *ExternalQuerySource `json:"query_source,omitempty"`
	// The time the query started, in milliseconds since the Unix epoch.
	QueryStartTimeMs int64 `json:"query_start_time_ms,omitempty"`
	// Key-value tags attached to the query execution, for example with the `SET
	// QUERY_TAGS` statement, for grouping and filtering. See
	// https://docs.databricks.com/aws/en/sql/user/queries/query-tags.
	QueryTags []QueryTag `json:"query_tags,omitempty"`
	// The text of the query.
	QueryText string `json:"query_text,omitempty"`
	// The number of results returned by the query.
	RowsProduced int64 `json:"rows_produced,omitempty"`
	// The spark session UUID that query ran on. This is either the Spark
	// Connect, DBSQL, or SDP session ID.
	SessionId string `json:"session_id,omitempty"`
	// URL to the Spark UI query plan.
	SparkUiUrl string `json:"spark_ui_url,omitempty"`
	// Type of statement for this query
	StatementType QueryStatementType `json:"statement_type,omitempty"`
	// Query status with one of the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	Status QueryStatus `json:"status,omitempty"`
	// The ID of the user who ran the query.
	UserId int64 `json:"user_id,omitempty"`
	// The email address or username of the user who ran the query.
	UserName string `json:"user_name,omitempty"`
	// ID of the warehouse that ran the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryList struct {
	// The total number of queries.
	Count int `json:"count,omitempty"`
	// The page number that is currently displayed.
	Page int `json:"page,omitempty"`
	// The number of queries per page.
	PageSize int `json:"page_size,omitempty"`
	// List of queries returned.
	Results []LegacyQuery `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs int64 `json:"compilation_time_ms,omitempty"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs int64 `json:"execution_time_ms,omitempty"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes int64 `json:"network_sent_bytes,omitempty"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load, in milliseconds since the Unix epoch. This field is optional
	// and will not appear if the query skipped the overloading queue.
	OverloadingQueueStartTimestamp int64 `json:"overloading_queue_start_timestamp,omitempty"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs int64 `json:"photon_total_time_ms,omitempty"`
	// projected remaining work to be done aggregated across all stages in the
	// query, in milliseconds
	ProjectedRemainingTaskTotalTimeMs int64 `json:"projected_remaining_task_total_time_ms,omitempty"`
	// projected lower bound on remaining total task time based on
	// projected_remaining_task_total_time_ms / maximum concurrency
	ProjectedRemainingWallclockTimeMs int64 `json:"projected_remaining_wallclock_time_ms,omitempty"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse, in milliseconds since the Unix epoch. This
	// field is optional and will not appear if the query skipped the
	// provisioning queue.
	ProvisioningQueueStartTimestamp int64 `json:"provisioning_queue_start_timestamp,omitempty"`
	// Total number of file bytes in all tables not read due to pruning
	PrunedBytes int64 `json:"pruned_bytes,omitempty"`
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount int64 `json:"pruned_files_count,omitempty"`
	// Timestamp of when the underlying compute started compilation of the
	// query, in milliseconds since the Unix epoch.
	QueryCompilationStartTimestamp int64 `json:"query_compilation_start_timestamp,omitempty"`
	// Total size of data read by the query, in bytes.
	ReadBytes int64 `json:"read_bytes,omitempty"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes int64 `json:"read_cache_bytes,omitempty"`
	// Total number of file bytes in all tables read
	ReadFilesBytes int64 `json:"read_files_bytes,omitempty"`
	// Number of files read after pruning
	ReadFilesCount int64 `json:"read_files_count,omitempty"`
	// Number of partitions read after pruning.
	ReadPartitionsCount int64 `json:"read_partitions_count,omitempty"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes int64 `json:"read_remote_bytes,omitempty"`
	// number of remaining tasks to complete this is based on the current status
	// and could be bigger or smaller in the future based on future updates
	RemainingTaskCount int64 `json:"remaining_task_count,omitempty"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs int64 `json:"result_fetch_time_ms,omitempty"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache bool `json:"result_from_cache,omitempty"`
	// Total number of rows returned by the query.
	RowsProducedCount int64 `json:"rows_produced_count,omitempty"`
	// Total number of rows read by the query.
	RowsReadCount int64 `json:"rows_read_count,omitempty"`
	// Deprecated: use `remaining_task_count` instead. Number of remaining tasks
	// to complete.
	RunnableTasks int64 `json:"runnable_tasks,omitempty"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes int64 `json:"spill_to_disk_bytes,omitempty"`
	// sum of task times completed in a range of wall clock time, approximated
	// to a configurable number of points aggregated over all stages and jobs in
	// the query (based on task_total_time_ms)
	TaskTimeOverTimeRange *TaskTimeOverRange `json:"task_time_over_time_range,omitempty"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs int64 `json:"task_total_time_ms,omitempty"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs int64 `json:"total_time_ms,omitempty"`
	// Deprecated: use `projected_remaining_task_total_time_ms` instead.
	// Remaining work to be done across all stages in the query, in
	// milliseconds.
	WorkToBeDone int64 `json:"work_to_be_done,omitempty"`
	// Size of persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes int64 `json:"write_remote_bytes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryMetrics) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryMetrics) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	Catalog string `json:"catalog,omitempty"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	Parameters []Parameter `json:"parameters,omitempty"`
	// The name of the schema to execute this query in.
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	DateRangeValue *DateRangeValue `json:"date_range_value,omitempty"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue *DateValue `json:"date_value,omitempty"`
	// Dropdown query parameter value.
	EnumValue *EnumValue `json:"enum_value,omitempty"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name string `json:"name,omitempty"`
	// Numeric query parameter value.
	NumericValue *NumericValue `json:"numeric_value,omitempty"`
	// Query-based dropdown query parameter value.
	QueryBackedValue *QueryBackedValue `json:"query_backed_value,omitempty"`
	// Text query parameter value.
	TextValue *TextValue `json:"text_value,omitempty"`
	// Text displayed in the user-facing parameter widget in the UI.
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryPostContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryPostContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

// Values returns all possible values for QueryStatementType.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueryStatementType) Values() []QueryStatementType {
	return []QueryStatementType{
		QueryStatementTypeAlter,
		QueryStatementTypeAnalyze,
		QueryStatementTypeCopy,
		QueryStatementTypeCreate,
		QueryStatementTypeDelete,
		QueryStatementTypeDescribe,
		QueryStatementTypeDrop,
		QueryStatementTypeExplain,
		QueryStatementTypeGrant,
		QueryStatementTypeInsert,
		QueryStatementTypeMerge,
		QueryStatementTypeOptimize,
		QueryStatementTypeOther,
		QueryStatementTypeRefresh,
		QueryStatementTypeReplace,
		QueryStatementTypeRevoke,
		QueryStatementTypeSelect,
		QueryStatementTypeSet,
		QueryStatementTypeShow,
		QueryStatementTypeTruncate,
		QueryStatementTypeUpdate,
		QueryStatementTypeUse,
	}
}

// Type always returns QueryStatementType to satisfy [pflag.Value] interface
func (f *QueryStatementType) Type() string {
	return "QueryStatementType"
}

// Statuses which are also used by OperationStatus in runtime. When adding a new
// QueryStatus, make sure to update
// com.databricks.sqlgateway.history.QueryStatusOrdering
type QueryStatus string

const QueryStatusCanceled QueryStatus = `CANCELED`

const QueryStatusCompiled QueryStatus = `COMPILED`

const QueryStatusCompiling QueryStatus = `COMPILING`

const QueryStatusFailed QueryStatus = `FAILED`

const QueryStatusFinished QueryStatus = `FINISHED`

const QueryStatusQueued QueryStatus = `QUEUED`

const QueryStatusRunning QueryStatus = `RUNNING`

const QueryStatusStarted QueryStatus = `STARTED`

// String representation for [fmt.Print]
func (f *QueryStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryStatus) Set(v string) error {
	switch v {
	case `CANCELED`, `COMPILED`, `COMPILING`, `FAILED`, `FINISHED`, `QUEUED`, `RUNNING`, `STARTED`:
		*f = QueryStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "COMPILED", "COMPILING", "FAILED", "FINISHED", "QUEUED", "RUNNING", "STARTED"`, v)
	}
}

// Values returns all possible values for QueryStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueryStatus) Values() []QueryStatus {
	return []QueryStatus{
		QueryStatusCanceled,
		QueryStatusCompiled,
		QueryStatusCompiling,
		QueryStatusFailed,
		QueryStatusFinished,
		QueryStatusQueued,
		QueryStatusRunning,
		QueryStatusStarted,
	}
}

// Type always returns QueryStatus to satisfy [pflag.Value] interface
func (f *QueryStatus) Type() string {
	return "QueryStatus"
}

// * A query execution can be annotated with an optional key-value pair to allow
// users to attribute the executions by key and optional value to filter by.
// QueryTag is the user-facing representation.
type QueryTag struct {
	// The key of the tag.
	Key string `json:"key"`
	// The value of the tag. Key-only tags have no value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *QueryTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	ConfigPair []EndpointConfPair `json:"config_pair,omitempty"`
	// Configuration parameters, each a key-value pair.
	ConfigurationPairs []EndpointConfPair `json:"configuration_pairs,omitempty"`
}

func (s *RepeatedEndpointConfPairs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RestoreDashboardRequest struct {
	DashboardId string `json:"-" url:"-"`
}

func (s *RestoreDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type RestoreQueriesLegacyRequest struct {
	QueryId string `json:"-" url:"-"`
}

func (s *RestoreQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Contains the result data of a single chunk when using `INLINE` disposition.
// When using `EXTERNAL_LINKS` disposition, the array `external_links` is used
// instead to provide short-lived cloud-storage URLs to the result data in cloud
// storage. Exactly one of these alternatives is used. (While the
// `external_links` array prepares the API to return multiple links in a single
// response. Currently only a single link is returned.)
type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray [][]string `json:"data_array,omitempty"`
	// Short-lived cloud-storage links to the result data, used with the
	// `EXTERNAL_LINKS` disposition.
	ExternalLinks []ExternalLink `json:"external_links,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	Chunks []BaseChunkInfo `json:"chunks,omitempty"`
	// The format of the result data, as requested with `format`.
	Format Format `json:"format,omitempty"`
	// The schema of the result set: the ordered list of column descriptions.
	Schema *ResultSchema `json:"schema,omitempty"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount int64 `json:"total_byte_count,omitempty"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount int `json:"total_chunk_count,omitempty"`
	// The total number of rows in the result set.
	TotalRowCount int64 `json:"total_row_count,omitempty"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {
	// The number of columns in the result set.
	ColumnCount int `json:"column_count,omitempty"`
	// The column descriptions, in result set order.
	Columns []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResultSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunAsMode string

const RunAsModeOwner RunAsMode = `OWNER`

const RunAsModeViewer RunAsMode = `VIEWER`

// String representation for [fmt.Print]
func (f *RunAsMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunAsMode) Set(v string) error {
	switch v {
	case `OWNER`, `VIEWER`:
		*f = RunAsMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OWNER", "VIEWER"`, v)
	}
}

// Values returns all possible values for RunAsMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunAsMode) Values() []RunAsMode {
	return []RunAsMode{
		RunAsModeOwner,
		RunAsModeViewer,
	}
}

// Type always returns RunAsMode to satisfy [pflag.Value] interface
func (f *RunAsMode) Type() string {
	return "RunAsMode"
}

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

// Values returns all possible values for RunAsRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunAsRole) Values() []RunAsRole {
	return []RunAsRole{
		RunAsRoleOwner,
		RunAsRoleViewer,
	}
}

// Type always returns RunAsRole to satisfy [pflag.Value] interface
func (f *RunAsRole) Type() string {
	return "RunAsRole"
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

type ServiceError struct {
	// The error code that classifies the failure.
	ErrorCode ServiceErrorCode `json:"error_code,omitempty"`
	// A brief summary of the error condition.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServiceError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServiceError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ServiceErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServiceErrorCode) Values() []ServiceErrorCode {
	return []ServiceErrorCode{
		ServiceErrorCodeAborted,
		ServiceErrorCodeAlreadyExists,
		ServiceErrorCodeBadRequest,
		ServiceErrorCodeCancelled,
		ServiceErrorCodeDeadlineExceeded,
		ServiceErrorCodeInternalError,
		ServiceErrorCodeIoError,
		ServiceErrorCodeNotFound,
		ServiceErrorCodeResourceExhausted,
		ServiceErrorCodeServiceUnderMaintenance,
		ServiceErrorCodeTemporarilyUnavailable,
		ServiceErrorCodeUnauthenticated,
		ServiceErrorCodeUnknown,
		ServiceErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ServiceErrorCode to satisfy [pflag.Value] interface
func (f *ServiceErrorCode) Type() string {
	return "ServiceErrorCode"
}

// Set object ACL
type SetRequest struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId string `json:"-" url:"-"`
	// The type of object permission to set.
	ObjectType ObjectTypePlural `json:"-" url:"-"`
}

func (s *SetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type SetResponse struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId string `json:"object_id,omitempty"`
	// A singular noun object type.
	ObjectType ObjectType `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SetResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Sets the workspace level warehouse configuration that is shared by all SQL
// warehouses in this workspace.
//
// This is idempotent.
type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel *Channel `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs `json:"config_param,omitempty"`
	// Spark configuration applied to every warehouse in the workspace, for
	// example to connect to an external Hive metastore or to configure access
	// to cloud storage. Only allowlisted keys are accepted, and the
	// JSON-serialized size must be at most 512 KB. See
	// https://docs.databricks.com/aws/en/admin/sql/data-access-configuration.
	DataAccessConfig []EndpointConfPair `json:"data_access_config,omitempty"`
	// Deprecated: only setting this to true is allowed.
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type.
	EnabledWarehouseTypes []WarehouseTypePair `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs `json:"global_param,omitempty"`
	// GCP only: Google service account that the warehouse clusters use to
	// access Google Cloud Storage.
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: The instance profile used to pass an IAM role to the SQL
	// warehouses. This configuration is also applied to the workspace's
	// serverless compute for notebooks and jobs.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for warehouses
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters applied to every warehouse in the workspace.
	// Only parameters that can be set at the workspace scope are accepted. See
	// https://docs.databricks.com/aws/en/sql/language-manual/sql-ref-parameters.
	SqlConfigurationParameters *RepeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SetWorkspaceWarehouseConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetWorkspaceWarehouseConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Security policy to be used for warehouses
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

// Values returns all possible values for SetWorkspaceWarehouseConfigRequestSecurityPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Values() []SetWorkspaceWarehouseConfigRequestSecurityPolicy {
	return []SetWorkspaceWarehouseConfigRequestSecurityPolicy{
		SetWorkspaceWarehouseConfigRequestSecurityPolicyDataAccessControl,
		SetWorkspaceWarehouseConfigRequestSecurityPolicyNone,
		SetWorkspaceWarehouseConfigRequestSecurityPolicyPassthrough,
	}
}

// Type always returns SetWorkspaceWarehouseConfigRequestSecurityPolicy to satisfy [pflag.Value] interface
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Type() string {
	return "SetWorkspaceWarehouseConfigRequestSecurityPolicy"
}

// EndpointSpotInstancePolicy configures whether the endpoint should use spot
// instances.
//
// The breakdown of how the EndpointSpotInstancePolicy converts to per cloud
// configurations is:
//
// - AWS, COST_OPTIMIZED: On Demand Driver with Spot Executors. - AWS,
// RELIABILITY_OPTIMIZED: On Demand Driver and Executors. - AZURE,
// COST_OPTIMIZED: On Demand Driver and Executors. - AZURE,
// RELIABILITY_OPTIMIZED: On Demand Driver and Executors.
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

// Values returns all possible values for SpotInstancePolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SpotInstancePolicy) Values() []SpotInstancePolicy {
	return []SpotInstancePolicy{
		SpotInstancePolicyCostOptimized,
		SpotInstancePolicyPolicyUnspecified,
		SpotInstancePolicyReliabilityOptimized,
	}
}

// Type always returns SpotInstancePolicy to satisfy [pflag.Value] interface
func (f *SpotInstancePolicy) Type() string {
	return "SpotInstancePolicy"
}

// Starts a SQL warehouse. This API is idempotent.
type StartRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func (s *StartRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// * State of a warehouse.
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

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateDeleted,
		StateDeleting,
		StateRunning,
		StateStarting,
		StateStopped,
		StateStopping,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	Name string `json:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type string `json:"type,omitempty"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StatementParameterListItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementParameterListItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Statement Execution works with all SQL warehouse types, including
// Lakehouse//RT. Statement Execution (Lakehouse//RT) is optimized for
// Lakehouse//RT warehouses, with lower latency and higher concurrency, and
// takes `warehouse_id` as a path parameter.
type StatementResponse struct {
	// The schema and metadata of the result set. Present when `status.state` is
	// `SUCCEEDED`.
	Manifest *ResultManifest `json:"manifest,omitempty"`
	// Contains the result data of a single chunk when using `INLINE`
	// disposition. When using `EXTERNAL_LINKS` disposition, the array
	// `external_links` is used instead to provide short-lived cloud-storage
	// URLs to the result data in cloud storage. Exactly one of these
	// alternatives is used. Calls to `getResultData` return the link for the
	// requested chunk; `executeStatement` and `getStatementResult` responses
	// can contain links for multiple chunks.
	Result *ResultData `json:"result,omitempty"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"statement_id,omitempty"`
	// The execution state of the statement and, if it failed, the error
	// details.
	Status *StatementStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StatementResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StatementState string

const StatementStateCanceled StatementState = `CANCELED`

const StatementStateClosed StatementState = `CLOSED`

const StatementStateFailed StatementState = `FAILED`

const StatementStatePending StatementState = `PENDING`

const StatementStateRunning StatementState = `RUNNING`

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

// Values returns all possible values for StatementState.
//
// There is no guarantee on the order of the values in the slice.
func (f *StatementState) Values() []StatementState {
	return []StatementState{
		StatementStateCanceled,
		StatementStateClosed,
		StatementStateFailed,
		StatementStatePending,
		StatementStateRunning,
		StatementStateSucceeded,
	}
}

// Type always returns StatementState to satisfy [pflag.Value] interface
func (f *StatementState) Type() string {
	return "StatementState"
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {
	// The error that caused the statement to fail. Only populated when `state`
	// is `FAILED`.
	Error *ServiceError `json:"error,omitempty"`
	// SQLSTATE error code returned when the statement execution fails. Only
	// populated when the statement status is `FAILED`.
	SqlState string `json:"sql_state,omitempty"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accompanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State StatementState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StatementStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Status string

const StatusDegraded Status = `DEGRADED`

const StatusFailed Status = `FAILED`

const StatusHealthy Status = `HEALTHY`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `DEGRADED`, `FAILED`, `HEALTHY`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEGRADED", "FAILED", "HEALTHY"`, v)
	}
}

// Values returns all possible values for Status.
//
// There is no guarantee on the order of the values in the slice.
func (f *Status) Values() []Status {
	return []Status{
		StatusDegraded,
		StatusFailed,
		StatusHealthy,
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

// Stops a SQL warehouse. This API is idempotent.
type StopRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func (s *StopRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type Success struct {
	Message SuccessMessage `json:"message,omitempty"`
}

func (s *Success) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
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

// Values returns all possible values for SuccessMessage.
//
// There is no guarantee on the order of the values in the slice.
func (f *SuccessMessage) Values() []SuccessMessage {
	return []SuccessMessage{
		SuccessMessageSuccess,
	}
}

// Type always returns SuccessMessage to satisfy [pflag.Value] interface
func (f *SuccessMessage) Type() string {
	return "SuccessMessage"
}

type TaskTimeOverRange struct {
	// Each entry covers one consecutive time bucket of length `interval`,
	// starting at `query_start_time_ms`.
	Entries []TaskTimeOverRangeEntry `json:"entries,omitempty"`
	// interval length for all entries (difference in start time and end time of
	// an entry range) the same for all entries start time of first interval is
	// query_start_time_ms
	Interval int64 `json:"interval,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TaskTimeOverRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskTimeOverRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TaskTimeOverRangeEntry struct {
	// total task completion time in this time range, aggregated over all stages
	// and jobs in the query
	TaskCompletedTimeMs int64 `json:"task_completed_time_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TaskTimeOverRangeEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskTimeOverRangeEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code TerminationReasonCode `json:"code,omitempty"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters map[string]string `json:"parameters,omitempty"`
	// type of the termination
	Type TerminationReasonType `json:"type,omitempty"`
}

func (s *TerminationReason) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The status code indicating why the cluster was terminated.
type TerminationReasonCode string

const TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`

const TerminationReasonCodeAccessTokenFailure TerminationReasonCode = `ACCESS_TOKEN_FAILURE`

const TerminationReasonCodeAllocationTimeout TerminationReasonCode = `ALLOCATION_TIMEOUT`

const TerminationReasonCodeAllocationTimeoutNodeDaemonNotReady TerminationReasonCode = `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`

const TerminationReasonCodeAllocationTimeoutNoHealthyAndWarmedUpClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoHealthyClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoMatchedClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoReadyClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoUnallocatedClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoWarmedUpClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInaccessibleKmsKeyFailure TerminationReasonCode = `AWS_INACCESSIBLE_KMS_KEY_FAILURE`

const TerminationReasonCodeAwsInstanceProfileUpdateFailure TerminationReasonCode = `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsInvalidKeyPair TerminationReasonCode = `AWS_INVALID_KEY_PAIR`

const TerminationReasonCodeAwsInvalidKmsKeyState TerminationReasonCode = `AWS_INVALID_KMS_KEY_STATE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsResourceQuotaExceeded TerminationReasonCode = `AWS_RESOURCE_QUOTA_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzurePackedDeploymentPartialFailure TerminationReasonCode = `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeBootstrapTimeoutDueToMisconfig TerminationReasonCode = `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`

const TerminationReasonCodeBudgetPolicyLimitEnforcementActivated TerminationReasonCode = `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`

const TerminationReasonCodeBudgetPolicyResolutionFailure TerminationReasonCode = `BUDGET_POLICY_RESOLUTION_FAILURE`

const TerminationReasonCodeCloudAccountPodQuotaExceeded TerminationReasonCode = `CLOUD_ACCOUNT_POD_QUOTA_EXCEEDED`

const TerminationReasonCodeCloudAccountSetupFailure TerminationReasonCode = `CLOUD_ACCOUNT_SETUP_FAILURE`

const TerminationReasonCodeCloudOperationCancelled TerminationReasonCode = `CLOUD_OPERATION_CANCELLED`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderInstanceNotLaunched TerminationReasonCode = `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailureDueToMisconfig TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderResourceStockoutDueToMisconfig TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeClusterOperationThrottled TerminationReasonCode = `CLUSTER_OPERATION_THROTTLED`

const TerminationReasonCodeClusterOperationTimeout TerminationReasonCode = `CLUSTER_OPERATION_TIMEOUT`

const TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneConnectionFailure TerminationReasonCode = `CONTROL_PLANE_CONNECTION_FAILURE`

const TerminationReasonCodeControlPlaneConnectionFailureDueToMisconfig TerminationReasonCode = `CONTROL_PLANE_CONNECTION_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailureDueToMisconfig TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDataAccessConfigChanged TerminationReasonCode = `DATA_ACCESS_CONFIG_CHANGED`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDbrImageResolutionFailure TerminationReasonCode = `DBR_IMAGE_RESOLUTION_FAILURE`

const TerminationReasonCodeDisasterRecoveryReplication TerminationReasonCode = `DISASTER_RECOVERY_REPLICATION`

const TerminationReasonCodeDnsResolutionError TerminationReasonCode = `DNS_RESOLUTION_ERROR`

const TerminationReasonCodeDockerContainerCreationException TerminationReasonCode = `DOCKER_CONTAINER_CREATION_EXCEPTION`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDockerImageTooLargeForInstanceException TerminationReasonCode = `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`

const TerminationReasonCodeDockerInvalidOsException TerminationReasonCode = `DOCKER_INVALID_OS_EXCEPTION`

const TerminationReasonCodeDriverEviction TerminationReasonCode = `DRIVER_EVICTION`

const TerminationReasonCodeDriverLaunchTimeout TerminationReasonCode = `DRIVER_LAUNCH_TIMEOUT`

const TerminationReasonCodeDriverNodeUnreachable TerminationReasonCode = `DRIVER_NODE_UNREACHABLE`

const TerminationReasonCodeDriverOutOfDisk TerminationReasonCode = `DRIVER_OUT_OF_DISK`

const TerminationReasonCodeDriverOutOfMemory TerminationReasonCode = `DRIVER_OUT_OF_MEMORY`

const TerminationReasonCodeDriverPodCreationFailure TerminationReasonCode = `DRIVER_POD_CREATION_FAILURE`

const TerminationReasonCodeDriverUnexpectedFailure TerminationReasonCode = `DRIVER_UNEXPECTED_FAILURE`

const TerminationReasonCodeDriverUnhealthy TerminationReasonCode = `DRIVER_UNHEALTHY`

const TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeDynamicSparkConfSizeExceeded TerminationReasonCode = `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`

const TerminationReasonCodeEosSparkImage TerminationReasonCode = `EOS_SPARK_IMAGE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeExecutorPodUnscheduled TerminationReasonCode = `EXECUTOR_POD_UNSCHEDULED`

const TerminationReasonCodeGcpApiRateQuotaExceeded TerminationReasonCode = `GCP_API_RATE_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpDeniedByOrgPolicy TerminationReasonCode = `GCP_DENIED_BY_ORG_POLICY`

const TerminationReasonCodeGcpForbidden TerminationReasonCode = `GCP_FORBIDDEN`

const TerminationReasonCodeGcpIamTimeout TerminationReasonCode = `GCP_IAM_TIMEOUT`

const TerminationReasonCodeGcpInaccessibleKmsKeyFailure TerminationReasonCode = `GCP_INACCESSIBLE_KMS_KEY_FAILURE`

const TerminationReasonCodeGcpInsufficientCapacity TerminationReasonCode = `GCP_INSUFFICIENT_CAPACITY`

const TerminationReasonCodeGcpIpSpaceExhausted TerminationReasonCode = `GCP_IP_SPACE_EXHAUSTED`

const TerminationReasonCodeGcpKmsKeyPermissionDenied TerminationReasonCode = `GCP_KMS_KEY_PERMISSION_DENIED`

const TerminationReasonCodeGcpNotFound TerminationReasonCode = `GCP_NOT_FOUND`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpResourceQuotaExceeded TerminationReasonCode = `GCP_RESOURCE_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountAccessDenied TerminationReasonCode = `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGcpServiceAccountNotFound TerminationReasonCode = `GCP_SERVICE_ACCOUNT_NOT_FOUND`

const TerminationReasonCodeGcpSubnetNotReady TerminationReasonCode = `GCP_SUBNET_NOT_READY`

const TerminationReasonCodeGcpTrustedImageProjectsViolated TerminationReasonCode = `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`

const TerminationReasonCodeGkeBasedClusterTermination TerminationReasonCode = `GKE_BASED_CLUSTER_TERMINATION`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHivemetastoreConnectivityFailure TerminationReasonCode = `HIVEMETASTORE_CONNECTIVITY_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`

const TerminationReasonCodeInitContainerNotFinished TerminationReasonCode = `INIT_CONTAINER_NOT_FINISHED`

const TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstancePoolMaxCapacityReached TerminationReasonCode = `INSTANCE_POOL_MAX_CAPACITY_REACHED`

const TerminationReasonCodeInstancePoolNotFound TerminationReasonCode = `INSTANCE_POOL_NOT_FOUND`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInstanceUnreachableDueToMisconfig TerminationReasonCode = `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`

const TerminationReasonCodeInternalCapacityFailure TerminationReasonCode = `INTERNAL_CAPACITY_FAILURE`

const TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidAwsParameter TerminationReasonCode = `INVALID_AWS_PARAMETER`

const TerminationReasonCodeInvalidInstancePlacementProtocol TerminationReasonCode = `INVALID_INSTANCE_PLACEMENT_PROTOCOL`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeInvalidWorkerImageFailure TerminationReasonCode = `INVALID_WORKER_IMAGE_FAILURE`

const TerminationReasonCodeInPenaltyBox TerminationReasonCode = `IN_PENALTY_BOX`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`

const TerminationReasonCodeK8sActivePodQuotaExceeded TerminationReasonCode = `K8S_ACTIVE_POD_QUOTA_EXCEEDED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeLazyAllocationTimeout TerminationReasonCode = `LAZY_ALLOCATION_TIMEOUT`

const TerminationReasonCodeMaintenanceMode TerminationReasonCode = `MAINTENANCE_MODE`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeMtlsPortConnectivityFailure TerminationReasonCode = `MTLS_PORT_CONNECTIVITY_FAILURE`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetvisorSetupTimeout TerminationReasonCode = `NETVISOR_SETUP_TIMEOUT`

const TerminationReasonCodeNetworkCheckControlPlaneFailure TerminationReasonCode = `NETWORK_CHECK_CONTROL_PLANE_FAILURE`

const TerminationReasonCodeNetworkCheckControlPlaneFailureDueToMisconfig TerminationReasonCode = `NETWORK_CHECK_CONTROL_PLANE_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeNetworkCheckDnsServerFailure TerminationReasonCode = `NETWORK_CHECK_DNS_SERVER_FAILURE`

const TerminationReasonCodeNetworkCheckDnsServerFailureDueToMisconfig TerminationReasonCode = `NETWORK_CHECK_DNS_SERVER_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeNetworkCheckMetadataEndpointFailure TerminationReasonCode = `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE`

const TerminationReasonCodeNetworkCheckMetadataEndpointFailureDueToMisconfig TerminationReasonCode = `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeNetworkCheckMultipleComponentsFailure TerminationReasonCode = `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE`

const TerminationReasonCodeNetworkCheckMultipleComponentsFailureDueToMisconfig TerminationReasonCode = `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeNetworkCheckNicFailure TerminationReasonCode = `NETWORK_CHECK_NIC_FAILURE`

const TerminationReasonCodeNetworkCheckNicFailureDueToMisconfig TerminationReasonCode = `NETWORK_CHECK_NIC_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeNetworkCheckStorageFailure TerminationReasonCode = `NETWORK_CHECK_STORAGE_FAILURE`

const TerminationReasonCodeNetworkCheckStorageFailureDueToMisconfig TerminationReasonCode = `NETWORK_CHECK_STORAGE_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNoMatchedK8s TerminationReasonCode = `NO_MATCHED_K8S`

const TerminationReasonCodeNoMatchedK8sTestingTag TerminationReasonCode = `NO_MATCHED_K8S_TESTING_TAG`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodePodAssignmentFailure TerminationReasonCode = `POD_ASSIGNMENT_FAILURE`

const TerminationReasonCodePodSchedulingFailure TerminationReasonCode = `POD_SCHEDULING_FAILURE`

const TerminationReasonCodeRateLimited TerminationReasonCode = `RATE_LIMITED`

const TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`

const TerminationReasonCodeResourceUsageBlocked TerminationReasonCode = `RESOURCE_USAGE_BLOCKED`

const TerminationReasonCodeSecretCreationFailure TerminationReasonCode = `SECRET_CREATION_FAILURE`

const TerminationReasonCodeSecretPermissionDenied TerminationReasonCode = `SECRET_PERMISSION_DENIED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeServerlessLongRunningTerminated TerminationReasonCode = `SERVERLESS_LONG_RUNNING_TERMINATED`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkImageDownloadThrottled TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_THROTTLED`

const TerminationReasonCodeSparkImageNotFound TerminationReasonCode = `SPARK_IMAGE_NOT_FOUND`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeSshBootstrapFailure TerminationReasonCode = `SSH_BOOTSTRAP_FAILURE`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStorageDownloadFailureDueToMisconfig TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeStorageDownloadFailureSlow TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE_SLOW`

const TerminationReasonCodeStorageDownloadFailureThrottled TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE_THROTTLED`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnexpectedPodRecreation TerminationReasonCode = `UNEXPECTED_POD_RECREATION`

const TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUsagePolicyEntitlementDenied TerminationReasonCode = `USAGE_POLICY_ENTITLEMENT_DENIED`

const TerminationReasonCodeUserInitiatedVmTermination TerminationReasonCode = `USER_INITIATED_VM_TERMINATION`

const TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`

const TerminationReasonCodeWorkspaceUpdate TerminationReasonCode = `WORKSPACE_UPDATE`

// String representation for [fmt.Print]
func (f *TerminationReasonCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonCode) Set(v string) error {
	switch v {
	case `ABUSE_DETECTED`, `ACCESS_TOKEN_FAILURE`, `ALLOCATION_TIMEOUT`, `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`, `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INACCESSIBLE_KMS_KEY_FAILURE`, `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_INVALID_KEY_PAIR`, `AWS_INVALID_KMS_KEY_STATE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_RESOURCE_QUOTA_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`, `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`, `BUDGET_POLICY_RESOLUTION_FAILURE`, `CLOUD_ACCOUNT_POD_QUOTA_EXCEEDED`, `CLOUD_ACCOUNT_SETUP_FAILURE`, `CLOUD_OPERATION_CANCELLED`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_SHUTDOWN`, `CLUSTER_OPERATION_THROTTLED`, `CLUSTER_OPERATION_TIMEOUT`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_CONNECTION_FAILURE`, `CONTROL_PLANE_CONNECTION_FAILURE_DUE_TO_MISCONFIG`, `CONTROL_PLANE_REQUEST_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`, `DATABASE_CONNECTION_FAILURE`, `DATA_ACCESS_CONFIG_CHANGED`, `DBFS_COMPONENT_UNHEALTHY`, `DBR_IMAGE_RESOLUTION_FAILURE`, `DISASTER_RECOVERY_REPLICATION`, `DNS_RESOLUTION_ERROR`, `DOCKER_CONTAINER_CREATION_EXCEPTION`, `DOCKER_IMAGE_PULL_FAILURE`, `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`, `DOCKER_INVALID_OS_EXCEPTION`, `DRIVER_EVICTION`, `DRIVER_LAUNCH_TIMEOUT`, `DRIVER_NODE_UNREACHABLE`, `DRIVER_OUT_OF_DISK`, `DRIVER_OUT_OF_MEMORY`, `DRIVER_POD_CREATION_FAILURE`, `DRIVER_UNEXPECTED_FAILURE`, `DRIVER_UNHEALTHY`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`, `EOS_SPARK_IMAGE`, `EXECUTION_COMPONENT_UNHEALTHY`, `EXECUTOR_POD_UNSCHEDULED`, `GCP_API_RATE_QUOTA_EXCEEDED`, `GCP_DENIED_BY_ORG_POLICY`, `GCP_FORBIDDEN`, `GCP_IAM_TIMEOUT`, `GCP_INACCESSIBLE_KMS_KEY_FAILURE`, `GCP_INSUFFICIENT_CAPACITY`, `GCP_IP_SPACE_EXHAUSTED`, `GCP_KMS_KEY_PERMISSION_DENIED`, `GCP_NOT_FOUND`, `GCP_QUOTA_EXCEEDED`, `GCP_RESOURCE_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GCP_SERVICE_ACCOUNT_NOT_FOUND`, `GCP_SUBNET_NOT_READY`, `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`, `GKE_BASED_CLUSTER_TERMINATION`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVEMETASTORE_CONNECTIVITY_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_CONTAINER_NOT_FINISHED`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_POOL_MAX_CAPACITY_REACHED`, `INSTANCE_POOL_NOT_FOUND`, `INSTANCE_UNREACHABLE`, `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`, `INTERNAL_CAPACITY_FAILURE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_AWS_PARAMETER`, `INVALID_INSTANCE_PLACEMENT_PROTOCOL`, `INVALID_SPARK_IMAGE`, `INVALID_WORKER_IMAGE_FAILURE`, `IN_PENALTY_BOX`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_ACTIVE_POD_QUOTA_EXCEEDED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `LAZY_ALLOCATION_TIMEOUT`, `MAINTENANCE_MODE`, `METASTORE_COMPONENT_UNHEALTHY`, `MTLS_PORT_CONNECTIVITY_FAILURE`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETVISOR_SETUP_TIMEOUT`, `NETWORK_CHECK_CONTROL_PLANE_FAILURE`, `NETWORK_CHECK_CONTROL_PLANE_FAILURE_DUE_TO_MISCONFIG`, `NETWORK_CHECK_DNS_SERVER_FAILURE`, `NETWORK_CHECK_DNS_SERVER_FAILURE_DUE_TO_MISCONFIG`, `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE`, `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE_DUE_TO_MISCONFIG`, `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE`, `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE_DUE_TO_MISCONFIG`, `NETWORK_CHECK_NIC_FAILURE`, `NETWORK_CHECK_NIC_FAILURE_DUE_TO_MISCONFIG`, `NETWORK_CHECK_STORAGE_FAILURE`, `NETWORK_CHECK_STORAGE_FAILURE_DUE_TO_MISCONFIG`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NO_MATCHED_K8S`, `NO_MATCHED_K8S_TESTING_TAG`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `POD_ASSIGNMENT_FAILURE`, `POD_SCHEDULING_FAILURE`, `RATE_LIMITED`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `RESOURCE_USAGE_BLOCKED`, `SECRET_CREATION_FAILURE`, `SECRET_PERMISSION_DENIED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SERVERLESS_LONG_RUNNING_TERMINATED`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_IMAGE_DOWNLOAD_THROTTLED`, `SPARK_IMAGE_NOT_FOUND`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `SSH_BOOTSTRAP_FAILURE`, `STORAGE_DOWNLOAD_FAILURE`, `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`, `STORAGE_DOWNLOAD_FAILURE_SLOW`, `STORAGE_DOWNLOAD_FAILURE_THROTTLED`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNEXPECTED_POD_RECREATION`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USAGE_POLICY_ENTITLEMENT_DENIED`, `USER_INITIATED_VM_TERMINATION`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`, `WORKSPACE_UPDATE`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ACCESS_TOKEN_FAILURE", "ALLOCATION_TIMEOUT", "ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY", "ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS", "ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_READY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INACCESSIBLE_KMS_KEY_FAILURE", "AWS_INSTANCE_PROFILE_UPDATE_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_INVALID_KEY_PAIR", "AWS_INVALID_KMS_KEY_STATE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_RESOURCE_QUOTA_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG", "BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED", "BUDGET_POLICY_RESOLUTION_FAILURE", "CLOUD_ACCOUNT_POD_QUOTA_EXCEEDED", "CLOUD_ACCOUNT_SETUP_FAILURE", "CLOUD_OPERATION_CANCELLED", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_SHUTDOWN", "CLUSTER_OPERATION_THROTTLED", "CLUSTER_OPERATION_TIMEOUT", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_CONNECTION_FAILURE", "CONTROL_PLANE_CONNECTION_FAILURE_DUE_TO_MISCONFIG", "CONTROL_PLANE_REQUEST_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG", "DATABASE_CONNECTION_FAILURE", "DATA_ACCESS_CONFIG_CHANGED", "DBFS_COMPONENT_UNHEALTHY", "DBR_IMAGE_RESOLUTION_FAILURE", "DISASTER_RECOVERY_REPLICATION", "DNS_RESOLUTION_ERROR", "DOCKER_CONTAINER_CREATION_EXCEPTION", "DOCKER_IMAGE_PULL_FAILURE", "DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION", "DOCKER_INVALID_OS_EXCEPTION", "DRIVER_EVICTION", "DRIVER_LAUNCH_TIMEOUT", "DRIVER_NODE_UNREACHABLE", "DRIVER_OUT_OF_DISK", "DRIVER_OUT_OF_MEMORY", "DRIVER_POD_CREATION_FAILURE", "DRIVER_UNEXPECTED_FAILURE", "DRIVER_UNHEALTHY", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "DYNAMIC_SPARK_CONF_SIZE_EXCEEDED", "EOS_SPARK_IMAGE", "EXECUTION_COMPONENT_UNHEALTHY", "EXECUTOR_POD_UNSCHEDULED", "GCP_API_RATE_QUOTA_EXCEEDED", "GCP_DENIED_BY_ORG_POLICY", "GCP_FORBIDDEN", "GCP_IAM_TIMEOUT", "GCP_INACCESSIBLE_KMS_KEY_FAILURE", "GCP_INSUFFICIENT_CAPACITY", "GCP_IP_SPACE_EXHAUSTED", "GCP_KMS_KEY_PERMISSION_DENIED", "GCP_NOT_FOUND", "GCP_QUOTA_EXCEEDED", "GCP_RESOURCE_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_ACCESS_DENIED", "GCP_SERVICE_ACCOUNT_DELETED", "GCP_SERVICE_ACCOUNT_NOT_FOUND", "GCP_SUBNET_NOT_READY", "GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED", "GKE_BASED_CLUSTER_TERMINATION", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVEMETASTORE_CONNECTIVITY_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_CONTAINER_NOT_FINISHED", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_POOL_MAX_CAPACITY_REACHED", "INSTANCE_POOL_NOT_FOUND", "INSTANCE_UNREACHABLE", "INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG", "INTERNAL_CAPACITY_FAILURE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_AWS_PARAMETER", "INVALID_INSTANCE_PLACEMENT_PROTOCOL", "INVALID_SPARK_IMAGE", "INVALID_WORKER_IMAGE_FAILURE", "IN_PENALTY_BOX", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_ACTIVE_POD_QUOTA_EXCEEDED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "LAZY_ALLOCATION_TIMEOUT", "MAINTENANCE_MODE", "METASTORE_COMPONENT_UNHEALTHY", "MTLS_PORT_CONNECTIVITY_FAILURE", "NEPHOS_RESOURCE_MANAGEMENT", "NETVISOR_SETUP_TIMEOUT", "NETWORK_CHECK_CONTROL_PLANE_FAILURE", "NETWORK_CHECK_CONTROL_PLANE_FAILURE_DUE_TO_MISCONFIG", "NETWORK_CHECK_DNS_SERVER_FAILURE", "NETWORK_CHECK_DNS_SERVER_FAILURE_DUE_TO_MISCONFIG", "NETWORK_CHECK_METADATA_ENDPOINT_FAILURE", "NETWORK_CHECK_METADATA_ENDPOINT_FAILURE_DUE_TO_MISCONFIG", "NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE", "NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE_DUE_TO_MISCONFIG", "NETWORK_CHECK_NIC_FAILURE", "NETWORK_CHECK_NIC_FAILURE_DUE_TO_MISCONFIG", "NETWORK_CHECK_STORAGE_FAILURE", "NETWORK_CHECK_STORAGE_FAILURE_DUE_TO_MISCONFIG", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NO_MATCHED_K8S", "NO_MATCHED_K8S_TESTING_TAG", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "POD_ASSIGNMENT_FAILURE", "POD_SCHEDULING_FAILURE", "RATE_LIMITED", "REQUEST_REJECTED", "REQUEST_THROTTLED", "RESOURCE_USAGE_BLOCKED", "SECRET_CREATION_FAILURE", "SECRET_PERMISSION_DENIED", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SERVERLESS_LONG_RUNNING_TERMINATED", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_IMAGE_DOWNLOAD_THROTTLED", "SPARK_IMAGE_NOT_FOUND", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "SSH_BOOTSTRAP_FAILURE", "STORAGE_DOWNLOAD_FAILURE", "STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG", "STORAGE_DOWNLOAD_FAILURE_SLOW", "STORAGE_DOWNLOAD_FAILURE_THROTTLED", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNEXPECTED_POD_RECREATION", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USAGE_POLICY_ENTITLEMENT_DENIED", "USER_INITIATED_VM_TERMINATION", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR", "WORKSPACE_UPDATE"`, v)
	}
}

// Values returns all possible values for TerminationReasonCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationReasonCode) Values() []TerminationReasonCode {
	return []TerminationReasonCode{
		TerminationReasonCodeAbuseDetected,
		TerminationReasonCodeAccessTokenFailure,
		TerminationReasonCodeAllocationTimeout,
		TerminationReasonCodeAllocationTimeoutNodeDaemonNotReady,
		TerminationReasonCodeAllocationTimeoutNoHealthyAndWarmedUpClusters,
		TerminationReasonCodeAllocationTimeoutNoHealthyClusters,
		TerminationReasonCodeAllocationTimeoutNoMatchedClusters,
		TerminationReasonCodeAllocationTimeoutNoReadyClusters,
		TerminationReasonCodeAllocationTimeoutNoUnallocatedClusters,
		TerminationReasonCodeAllocationTimeoutNoWarmedUpClusters,
		TerminationReasonCodeAttachProjectFailure,
		TerminationReasonCodeAwsAuthorizationFailure,
		TerminationReasonCodeAwsInaccessibleKmsKeyFailure,
		TerminationReasonCodeAwsInstanceProfileUpdateFailure,
		TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure,
		TerminationReasonCodeAwsInsufficientInstanceCapacityFailure,
		TerminationReasonCodeAwsInvalidKeyPair,
		TerminationReasonCodeAwsInvalidKmsKeyState,
		TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure,
		TerminationReasonCodeAwsRequestLimitExceeded,
		TerminationReasonCodeAwsResourceQuotaExceeded,
		TerminationReasonCodeAwsUnsupportedFailure,
		TerminationReasonCodeAzureByokKeyPermissionFailure,
		TerminationReasonCodeAzureEphemeralDiskFailure,
		TerminationReasonCodeAzureInvalidDeploymentTemplate,
		TerminationReasonCodeAzureOperationNotAllowedException,
		TerminationReasonCodeAzurePackedDeploymentPartialFailure,
		TerminationReasonCodeAzureQuotaExceededException,
		TerminationReasonCodeAzureResourceManagerThrottling,
		TerminationReasonCodeAzureResourceProviderThrottling,
		TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure,
		TerminationReasonCodeAzureVmExtensionFailure,
		TerminationReasonCodeAzureVnetConfigurationFailure,
		TerminationReasonCodeBootstrapTimeout,
		TerminationReasonCodeBootstrapTimeoutCloudProviderException,
		TerminationReasonCodeBootstrapTimeoutDueToMisconfig,
		TerminationReasonCodeBudgetPolicyLimitEnforcementActivated,
		TerminationReasonCodeBudgetPolicyResolutionFailure,
		TerminationReasonCodeCloudAccountPodQuotaExceeded,
		TerminationReasonCodeCloudAccountSetupFailure,
		TerminationReasonCodeCloudOperationCancelled,
		TerminationReasonCodeCloudProviderDiskSetupFailure,
		TerminationReasonCodeCloudProviderInstanceNotLaunched,
		TerminationReasonCodeCloudProviderLaunchFailure,
		TerminationReasonCodeCloudProviderLaunchFailureDueToMisconfig,
		TerminationReasonCodeCloudProviderResourceStockout,
		TerminationReasonCodeCloudProviderResourceStockoutDueToMisconfig,
		TerminationReasonCodeCloudProviderShutdown,
		TerminationReasonCodeClusterOperationThrottled,
		TerminationReasonCodeClusterOperationTimeout,
		TerminationReasonCodeCommunicationLost,
		TerminationReasonCodeContainerLaunchFailure,
		TerminationReasonCodeControlPlaneConnectionFailure,
		TerminationReasonCodeControlPlaneConnectionFailureDueToMisconfig,
		TerminationReasonCodeControlPlaneRequestFailure,
		TerminationReasonCodeControlPlaneRequestFailureDueToMisconfig,
		TerminationReasonCodeDatabaseConnectionFailure,
		TerminationReasonCodeDataAccessConfigChanged,
		TerminationReasonCodeDbfsComponentUnhealthy,
		TerminationReasonCodeDbrImageResolutionFailure,
		TerminationReasonCodeDisasterRecoveryReplication,
		TerminationReasonCodeDnsResolutionError,
		TerminationReasonCodeDockerContainerCreationException,
		TerminationReasonCodeDockerImagePullFailure,
		TerminationReasonCodeDockerImageTooLargeForInstanceException,
		TerminationReasonCodeDockerInvalidOsException,
		TerminationReasonCodeDriverEviction,
		TerminationReasonCodeDriverLaunchTimeout,
		TerminationReasonCodeDriverNodeUnreachable,
		TerminationReasonCodeDriverOutOfDisk,
		TerminationReasonCodeDriverOutOfMemory,
		TerminationReasonCodeDriverPodCreationFailure,
		TerminationReasonCodeDriverUnexpectedFailure,
		TerminationReasonCodeDriverUnhealthy,
		TerminationReasonCodeDriverUnreachable,
		TerminationReasonCodeDriverUnresponsive,
		TerminationReasonCodeDynamicSparkConfSizeExceeded,
		TerminationReasonCodeEosSparkImage,
		TerminationReasonCodeExecutionComponentUnhealthy,
		TerminationReasonCodeExecutorPodUnscheduled,
		TerminationReasonCodeGcpApiRateQuotaExceeded,
		TerminationReasonCodeGcpDeniedByOrgPolicy,
		TerminationReasonCodeGcpForbidden,
		TerminationReasonCodeGcpIamTimeout,
		TerminationReasonCodeGcpInaccessibleKmsKeyFailure,
		TerminationReasonCodeGcpInsufficientCapacity,
		TerminationReasonCodeGcpIpSpaceExhausted,
		TerminationReasonCodeGcpKmsKeyPermissionDenied,
		TerminationReasonCodeGcpNotFound,
		TerminationReasonCodeGcpQuotaExceeded,
		TerminationReasonCodeGcpResourceQuotaExceeded,
		TerminationReasonCodeGcpServiceAccountAccessDenied,
		TerminationReasonCodeGcpServiceAccountDeleted,
		TerminationReasonCodeGcpServiceAccountNotFound,
		TerminationReasonCodeGcpSubnetNotReady,
		TerminationReasonCodeGcpTrustedImageProjectsViolated,
		TerminationReasonCodeGkeBasedClusterTermination,
		TerminationReasonCodeGlobalInitScriptFailure,
		TerminationReasonCodeHivemetastoreConnectivityFailure,
		TerminationReasonCodeHiveMetastoreProvisioningFailure,
		TerminationReasonCodeImagePullPermissionDenied,
		TerminationReasonCodeInactivity,
		TerminationReasonCodeInitContainerNotFinished,
		TerminationReasonCodeInitScriptFailure,
		TerminationReasonCodeInstancePoolClusterFailure,
		TerminationReasonCodeInstancePoolMaxCapacityReached,
		TerminationReasonCodeInstancePoolNotFound,
		TerminationReasonCodeInstanceUnreachable,
		TerminationReasonCodeInstanceUnreachableDueToMisconfig,
		TerminationReasonCodeInternalCapacityFailure,
		TerminationReasonCodeInternalError,
		TerminationReasonCodeInvalidArgument,
		TerminationReasonCodeInvalidAwsParameter,
		TerminationReasonCodeInvalidInstancePlacementProtocol,
		TerminationReasonCodeInvalidSparkImage,
		TerminationReasonCodeInvalidWorkerImageFailure,
		TerminationReasonCodeInPenaltyBox,
		TerminationReasonCodeIpExhaustionFailure,
		TerminationReasonCodeJobFinished,
		TerminationReasonCodeK8sActivePodQuotaExceeded,
		TerminationReasonCodeK8sAutoscalingFailure,
		TerminationReasonCodeK8sDbrClusterLaunchTimeout,
		TerminationReasonCodeLazyAllocationTimeout,
		TerminationReasonCodeMaintenanceMode,
		TerminationReasonCodeMetastoreComponentUnhealthy,
		TerminationReasonCodeMtlsPortConnectivityFailure,
		TerminationReasonCodeNephosResourceManagement,
		TerminationReasonCodeNetvisorSetupTimeout,
		TerminationReasonCodeNetworkCheckControlPlaneFailure,
		TerminationReasonCodeNetworkCheckControlPlaneFailureDueToMisconfig,
		TerminationReasonCodeNetworkCheckDnsServerFailure,
		TerminationReasonCodeNetworkCheckDnsServerFailureDueToMisconfig,
		TerminationReasonCodeNetworkCheckMetadataEndpointFailure,
		TerminationReasonCodeNetworkCheckMetadataEndpointFailureDueToMisconfig,
		TerminationReasonCodeNetworkCheckMultipleComponentsFailure,
		TerminationReasonCodeNetworkCheckMultipleComponentsFailureDueToMisconfig,
		TerminationReasonCodeNetworkCheckNicFailure,
		TerminationReasonCodeNetworkCheckNicFailureDueToMisconfig,
		TerminationReasonCodeNetworkCheckStorageFailure,
		TerminationReasonCodeNetworkCheckStorageFailureDueToMisconfig,
		TerminationReasonCodeNetworkConfigurationFailure,
		TerminationReasonCodeNfsMountFailure,
		TerminationReasonCodeNoMatchedK8s,
		TerminationReasonCodeNoMatchedK8sTestingTag,
		TerminationReasonCodeNpipTunnelSetupFailure,
		TerminationReasonCodeNpipTunnelTokenFailure,
		TerminationReasonCodePodAssignmentFailure,
		TerminationReasonCodePodSchedulingFailure,
		TerminationReasonCodeRateLimited,
		TerminationReasonCodeRequestRejected,
		TerminationReasonCodeRequestThrottled,
		TerminationReasonCodeResourceUsageBlocked,
		TerminationReasonCodeSecretCreationFailure,
		TerminationReasonCodeSecretPermissionDenied,
		TerminationReasonCodeSecretResolutionError,
		TerminationReasonCodeSecurityDaemonRegistrationException,
		TerminationReasonCodeSelfBootstrapFailure,
		TerminationReasonCodeServerlessLongRunningTerminated,
		TerminationReasonCodeSkippedSlowNodes,
		TerminationReasonCodeSlowImageDownload,
		TerminationReasonCodeSparkError,
		TerminationReasonCodeSparkImageDownloadFailure,
		TerminationReasonCodeSparkImageDownloadThrottled,
		TerminationReasonCodeSparkImageNotFound,
		TerminationReasonCodeSparkStartupFailure,
		TerminationReasonCodeSpotInstanceTermination,
		TerminationReasonCodeSshBootstrapFailure,
		TerminationReasonCodeStorageDownloadFailure,
		TerminationReasonCodeStorageDownloadFailureDueToMisconfig,
		TerminationReasonCodeStorageDownloadFailureSlow,
		TerminationReasonCodeStorageDownloadFailureThrottled,
		TerminationReasonCodeStsClientSetupFailure,
		TerminationReasonCodeSubnetExhaustedFailure,
		TerminationReasonCodeTemporarilyUnavailable,
		TerminationReasonCodeTrialExpired,
		TerminationReasonCodeUnexpectedLaunchFailure,
		TerminationReasonCodeUnexpectedPodRecreation,
		TerminationReasonCodeUnknown,
		TerminationReasonCodeUnsupportedInstanceType,
		TerminationReasonCodeUpdateInstanceProfileFailure,
		TerminationReasonCodeUsagePolicyEntitlementDenied,
		TerminationReasonCodeUserInitiatedVmTermination,
		TerminationReasonCodeUserRequest,
		TerminationReasonCodeWorkerSetupFailure,
		TerminationReasonCodeWorkspaceCancelledError,
		TerminationReasonCodeWorkspaceConfigurationError,
		TerminationReasonCodeWorkspaceUpdate,
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

// Values returns all possible values for TerminationReasonType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationReasonType) Values() []TerminationReasonType {
	return []TerminationReasonType{
		TerminationReasonTypeClientError,
		TerminationReasonTypeCloudFailure,
		TerminationReasonTypeServiceFault,
		TerminationReasonTypeSuccess,
	}
}

// Type always returns TerminationReasonType to satisfy [pflag.Value] interface
func (f *TerminationReasonType) Type() string {
	return "TerminationReasonType"
}

type TextValue struct {
	// Text substituted for the parameter marker when the query runs. Quotation
	// marks are added around the value automatically.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TextValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TextValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TimeRange struct {
	// End of the range, in milliseconds since the Unix epoch.
	EndTimeMs int64 `json:"end_time_ms,omitempty" url:"end_time_ms,omitempty"`
	// Start of the range, in milliseconds since the Unix epoch.
	StartTimeMs int64 `json:"start_time_ms,omitempty" url:"start_time_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TimeRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TimeRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner string `json:"new_owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransferOwnershipObjectId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransferOwnershipObjectId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner string `json:"new_owner,omitempty"`
	// The ID of the object on which to change ownership.
	ObjectId TransferOwnershipObjectId `json:"-" url:"-"`
	// The type of object on which to change ownership.
	ObjectType OwnableObjectType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransferOwnershipRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransferOwnershipRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TrashAlertRequest struct {
	// UUID of the alert to move to the trash.
	Id string `json:"-" url:"-"`
}

func (s *TrashAlertRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type TrashAlertV2Request struct {
	// ID of the alert to move to the trash.
	Id string `json:"-" url:"-"`
	// Whether to permanently delete the alert. If not set, the alert will only
	// be soft deleted.
	Purge bool `json:"-" url:"purge,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TrashAlertV2Request) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TrashAlertV2Request) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TrashQueryRequest struct {
	// UUID of the query to move to the trash.
	Id string `json:"-" url:"-"`
}

func (s *TrashQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateAlertRequest struct {
	// Alert fields to update. Only the fields named in `update_mask` are
	// changed.
	Alert *UpdateAlertRequestAlert `json:"alert,omitempty"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`
	// UUID of the alert to update.
	Id string `json:"-" url:"-"`
	// Comma-separated list of the alert fields to update, such as
	// `display_name` or `condition`. Fields not listed keep their current
	// values.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateAlertRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAlertRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A legacy Databricks SQL alert. An alert runs a saved query, compares a column
// of the result against a threshold, and notifies its subscribers when the
// condition is met. See
// https://docs.databricks.com/aws/en/sql/user/alerts/legacy for details.
type UpdateAlertRequestAlert struct {
	// Condition that triggers the alert: a column of the query result, a
	// comparison operator, and a threshold value. The value in the first row of
	// the column is evaluated; aggregations are not supported through this API.
	Condition *AlertCondition `json:"condition,omitempty"`
	// Custom body of the alert notification. It can use the template variables
	// and HTML described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default notification body is used.
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of the alert notification, used as the email subject or
	// the Slack message header. It can use the template variables described at
	// https://docs.databricks.com/aws/en/sql/user/alerts/legacy. If omitted,
	// the default subject is used.
	CustomSubject string `json:"custom_subject,omitempty"`
	// Display name of the alert, shown in list views and notifications. It must
	// not be empty.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to notify subscribers when the alert state changes from
	// `TRIGGERED` back to `OK`. Defaults to true.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// UUID of the saved query whose result the alert evaluates. The caller must
	// have permission to view the query.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds the alert waits after a notification before it can
	// notify again while its state stays `TRIGGERED`. If 0 or not specified,
	// the alert notifies once when its state changes to `TRIGGERED` and does
	// not notify again until its state has changed.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateAlertRequestAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAlertRequestAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateAlertV2Request struct {
	Alert AlertV2 `json:"alert"`
	// The canonical identifier of the alert to retrieve information about.
	Id string `json:"-" url:"-"`
	// Comma-separated list of the alert fields to update, for example
	// `display_name,query_text` or `schedule.pause_status`. Only the listed
	// fields are changed; all other fields keep their current values.
	// `parent_path` cannot be updated.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateAlertV2Request) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateDefaultWarehouseOverrideRequest struct {
	// If set to true, and the override is not found, a new override will be
	// created. In this situation, `update_mask` is ignored and all fields are
	// applied. Defaults to false.
	AllowMissing bool `json:"-" url:"allow_missing,omitempty"`
	// Required. The default warehouse override to update. The name field must
	// be set in the format:
	// default-warehouse-overrides/{default_warehouse_override_id} The
	// default_warehouse_override_id can be a numeric user ID or the literal
	// string "me" for the current user.
	DefaultWarehouseOverride DefaultWarehouseOverride `json:"default_warehouse_override"`
	// The resource name of the default warehouse override. Format:
	// default-warehouse-overrides/{default_warehouse_override_id}
	Name string `json:"-" url:"-"`
	// Required. Field mask specifying which fields to update. Only the fields
	// specified in the mask will be updated. Use "*" to update all fields. When
	// allow_missing is true, this field is ignored and all fields are applied.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateDefaultWarehouseOverrideRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDefaultWarehouseOverrideRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateQueryRequest struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`
	// UUID of the query to update.
	Id string `json:"-" url:"-"`
	// Query fields to update. Only the fields named in `update_mask` are
	// changed.
	Query *UpdateQueryRequestQuery `json:"query,omitempty"`
	// Comma-separated list of the query fields to update, such as the display
	// name or query text. Fields not listed keep their current values.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateQueryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A saved Databricks SQL query. A query stores the SQL text, the SQL warehouse
// it runs on, its parameters, and metadata such as the display name,
// description, owner, and tags.
type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// Username of the user that owns the query.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Whose credentials the query runs with when it is shared. `VIEWER` uses
	// the credentials of the user running the query. `OWNER` uses the owner's
	// credentials where the product supports that, such as for legacy alerts
	// and jobs. When omitted on create, the workspace default applies. A
	// workspace admin can disallow `OWNER`.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`
	// Tags attached to the query. Tags appear in the query list and can be used
	// to filter it.
	Tags []string `json:"tags,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateQueryRequestQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateQueryRequestQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateVisualizationRequest struct {
	Id string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"update_mask"`

	Visualization *UpdateVisualizationRequestVisualization `json:"visualization,omitempty"`
}

func (s *UpdateVisualizationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	DisplayName string `json:"display_name,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateVisualizationRequestVisualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateVisualizationRequestVisualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWidgetRequest struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId string `json:"dashboard_id"`
	// Widget ID returned by :method:dashboardwidgets/create
	Id string `json:"-" url:"-"`

	Options WidgetOptions `json:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text string `json:"text,omitempty"`
	// Query Visualization ID returned by :method:queryvisualizations/create.
	VisualizationId string `json:"visualization_id,omitempty"`
	// Width of a widget
	Width int `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateWidgetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateWidgetRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type User struct {
	Email string `json:"email,omitempty"`

	Id int `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s User) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	CreateTime string `json:"create_time,omitempty"`
	// The display name of the visualization.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the visualization.
	Id string `json:"id,omitempty"`
	// UUID of the query that the visualization is attached to.
	QueryId string `json:"query_id,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	Type string `json:"type,omitempty"`
	// The timestamp indicating when the visualization was updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Visualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Visualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WarehouseAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WarehouseAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehouseAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	AllPermissions []WarehousePermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WarehouseAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehouseAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WarehousePermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WarehousePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehousePermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type WarehousePermissionLevel string

const WarehousePermissionLevelCanManage WarehousePermissionLevel = `CAN_MANAGE`

const WarehousePermissionLevelCanMonitor WarehousePermissionLevel = `CAN_MONITOR`

const WarehousePermissionLevelCanUse WarehousePermissionLevel = `CAN_USE`

const WarehousePermissionLevelCanView WarehousePermissionLevel = `CAN_VIEW`

const WarehousePermissionLevelIsOwner WarehousePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *WarehousePermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarehousePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MONITOR`, `CAN_USE`, `CAN_VIEW`, `IS_OWNER`:
		*f = WarehousePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MONITOR", "CAN_USE", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for WarehousePermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarehousePermissionLevel) Values() []WarehousePermissionLevel {
	return []WarehousePermissionLevel{
		WarehousePermissionLevelCanManage,
		WarehousePermissionLevelCanMonitor,
		WarehousePermissionLevelCanUse,
		WarehousePermissionLevelCanView,
		WarehousePermissionLevelIsOwner,
	}
}

// Type always returns WarehousePermissionLevel to satisfy [pflag.Value] interface
func (f *WarehousePermissionLevel) Type() string {
	return "WarehousePermissionLevel"
}

type WarehousePermissions struct {
	AccessControlList []WarehouseAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WarehousePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehousePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WarehousePermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WarehousePermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehousePermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WarehousePermissionsRequest struct {
	AccessControlList []WarehouseAccessControlRequest `json:"access_control_list,omitempty"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" url:"-"`
}

func (s *WarehousePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// * Configuration values to enable or disable the access to specific warehouse
// types in the workspace.
type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled bool `json:"enabled,omitempty"`
	// Warehouse type that this entry enables or disables.
	WarehouseType WarehouseTypePairWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WarehouseTypePair) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehouseTypePair) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WarehouseTypePairWarehouseType string

const WarehouseTypePairWarehouseTypeClassic WarehouseTypePairWarehouseType = `CLASSIC`

const WarehouseTypePairWarehouseTypePro WarehouseTypePairWarehouseType = `PRO`

const WarehouseTypePairWarehouseTypeRealtime WarehouseTypePairWarehouseType = `REALTIME`

const WarehouseTypePairWarehouseTypeTypeUnspecified WarehouseTypePairWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *WarehouseTypePairWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarehouseTypePairWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `REALTIME`, `TYPE_UNSPECIFIED`:
		*f = WarehouseTypePairWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "REALTIME", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for WarehouseTypePairWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarehouseTypePairWarehouseType) Values() []WarehouseTypePairWarehouseType {
	return []WarehouseTypePairWarehouseType{
		WarehouseTypePairWarehouseTypeClassic,
		WarehouseTypePairWarehouseTypePro,
		WarehouseTypePairWarehouseTypeRealtime,
		WarehouseTypePairWarehouseTypeTypeUnspecified,
	}
}

// Type always returns WarehouseTypePairWarehouseType to satisfy [pflag.Value] interface
func (f *WarehouseTypePairWarehouseType) Type() string {
	return "WarehouseTypePairWarehouseType"
}

type Widget struct {
	// The unique ID for this widget.
	Id string `json:"id,omitempty"`

	Options *WidgetOptions `json:"options,omitempty"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization *LegacyVisualization `json:"visualization,omitempty"`
	// Unused field.
	Width int `json:"width,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Widget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Widget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WidgetOptions struct {
	// Timestamp when this object was created
	CreatedAt string `json:"created_at,omitempty"`
	// Custom description of the widget
	Description string `json:"description,omitempty"`
	// Whether this widget is hidden on the dashboard.
	IsHidden bool `json:"isHidden,omitempty"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings any `json:"parameterMappings,omitempty"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position *WidgetPosition `json:"position,omitempty"`
	// Custom title of the widget
	Title string `json:"title,omitempty"`
	// Timestamp of the last time this object was updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WidgetOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WidgetOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	AutoHeight bool `json:"autoHeight,omitempty"`
	// column in the dashboard grid. Values start with 0
	Col int `json:"col,omitempty"`
	// row in the dashboard grid. Values start with 0
	Row int `json:"row,omitempty"`
	// width of the widget measured in dashboard grid cells
	SizeX int `json:"sizeX,omitempty"`
	// height of the widget measured in dashboard grid cells
	SizeY int `json:"sizeY,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WidgetPosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WidgetPosition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
