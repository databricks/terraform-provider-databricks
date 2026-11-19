---
subcategory: "Databricks SQL"
---
# databricks_alert_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/alertsv2)

The SQL Alert v2 data source allows you to retrieve detailed information about a specific alert in Databricks SQL. This data source provides access to all alert properties, including its configuration, evaluation criteria, notification settings, and schedule.

You can use this data source to:
- Retrieve alert details for reference in other resources
- Check the current state and configuration of an alert
- Verify notification settings and subscribers
- Examine the schedule configuration

## Example Usage
### Retrieve Alert by ID
This example retrieves a specific alert by its ID:

```hcl
data "databricks_alert_v2" "this" {
  id = "123"
}
```


## Arguments
The following arguments are supported:
* `id` (string, required) - The canonical identifier of the alert to retrieve information about
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `create_time` (string) - The timestamp indicating when the alert was created
* `custom_description` (string) - Custom description for the alert, used as the notification body. Interpreted according to `custom_template_format`.
  The available template variables are listed at https://docs.databricks.com/aws/en/sql/user/alerts/create#notification-templates
* `custom_summary` (string) - Custom summary for the alert, used as the notification subject. Interpreted according to `custom_template_format`
* `custom_template_format` (string) - The format used to interpret the `custom_summary` and `custom_description` templates.
  
  `HTML` treats both templates as HTML with Mustache `{{VARIABLE_NAME}}` placeholders.
  `MARKDOWN` treats them as Markdown with allowlisted `@VARIABLE_NAME` placeholders and converts
  them to the format each notification destination expects, so a single template renders
  correctly in email, Slack, and Microsoft Teams.
  
  When unset, the API applies no default. Responses omit the field unless a format is stored,
  and a template with no format is rendered as HTML. Possible values are: `HTML`, `MARKDOWN`
* `display_name` (string) - The display name of the alert
* `effective_run_as` (AlertV2RunAs) - The actual identity that will be used to execute the alert.
  This is an output-only field that shows the resolved run-as identity after applying
  permissions and defaults
* `evaluation` (AlertV2Evaluation) - Condition evaluated against the query result, and who is notified when it is met
* `id` (string) - The canonical identifier of the alert to retrieve information about
* `lifecycle_state` (string) - Indicates whether the alert is active (`ACTIVE`) or has been moved to the trash (`DELETED`). Possible values are: `ACTIVE`, `DELETED`
* `owner_user_name` (string) - The owner's username. This field is set to "Unavailable" if the user has been deleted
* `parameters` (list of AlertStatementParameter) - A list of parameters to pass into the alert SQL query statement containing parameter markers. Static values only.
  
  Reference a parameter in the query text as `:name`. Each parameter must have a unique, non-empty name.
  Each parameter consists of a name, a value, and optionally a type. To represent a NULL
  value, the `value` field may be omitted or set to `null` explicitly. If the `type` field
  is omitted, the value is interpreted as a string.
  
  If the type is given, parameters will be checked for type correctness according
  to the given type. A value is correct if the provided string can be converted to
  the requested type using the `cast` function. The exact semantics are described in
  the section [`cast` function](https://docs.databricks.com/sql/language-manual/functions/cast.html) of the SQL language reference
* `parent_path` (string) - The workspace path of the folder containing the alert. Can only be set on create, and cannot be updated
* `query_text` (string) - Text of the query to be run
* `run_as` (AlertV2RunAs) - Specifies the identity that will be used to run the alert.
  This field allows you to configure alerts to run as a specific user or service principal.
  - For user identity: Set `user_name` to the email of an active workspace user. Users can only set this to their own email.
  - For service principal: Set `service_principal_name` to the application ID. Requires the `servicePrincipal/user` role.
  If not specified, the alert will run as the request user
* `run_as_user_name` (string, deprecated) - Deprecated: Use `run_as` field instead. This field will be removed in a future release.
  The run as username or application ID of service principal.
  On Create and Update, this field can be set to application ID of an active service principal. Setting this field requires the servicePrincipal/user role
* `schedule` (CronSchedule) - Schedule on which the alert runs its query and evaluates the condition
* `update_time` (string) - The timestamp indicating when the alert was updated
* `warehouse_id` (string) - ID of the SQL warehouse attached to the alert

### AlertStatementParameter
* `name` (string) - The name of the parameter. Reference it in the query text as `:name`. Required, must be
  non-empty, and must be unique across the alert's parameters
* `type` (string) - The SQL data type of the parameter, for example `STRING`, `INT`, or `DECIMAL(10, 2)`. If no type is given
  the type is assumed to be `STRING`. Complex types such as `ARRAY`, `MAP`, and `STRUCT` are not supported
* `value` (string) - The value bound to the parameter, represented as a string. If omitted, the value is interpreted as NULL

### AlertV2Evaluation
* `comparison_operator` (string) - Operator used for comparison in alert evaluation. The alert is triggered when the comparison of the source value with the threshold is true. `IS_NULL` and `IS_NOT_NULL` take no threshold. Possible values are: `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NOT_NULL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`
* `empty_result_state` (string) - Alert state if result is empty. If not set, an empty result sets the state to `ERROR`. Please avoid setting this field to be `UNKNOWN` because `UNKNOWN` state is planned to be deprecated. Possible values are: `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`
* `last_evaluated_at` (string) - Timestamp of the last evaluation
* `notification` (AlertV2Notification) - User or Notification Destination to notify when alert is triggered
* `source` (AlertV2OperandColumn) - Source column from result to use to evaluate alert
* `state` (string) - Latest state of alert evaluation: `TRIGGERED` when the condition was met, `OK` when it was not, and `ERROR` when the evaluation failed. Possible values are: `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`
* `threshold` (AlertV2Operand) - Threshold to compare the source column against, either another column or a fixed value. Not needed when `comparison_operator` is `IS_NULL` or `IS_NOT_NULL`

### AlertV2Notification
* `notify_on_ok` (boolean) - Whether to notify alert subscribers when alert returns back to normal
* `retrigger_seconds` (integer) - Number of seconds an alert waits after being triggered before it is allowed to send another notification.
  If set to 0 or omitted, the alert will not send any further notifications after the first trigger
  Setting this value to 1 allows the alert to send a notification on every evaluation where the condition is met, effectively making it always retrigger for notification purposes
* `subscriptions` (list of AlertV2Subscription) - Users and notification destinations that receive notifications

### AlertV2Operand
* `column` (AlertV2OperandColumn) - Compare the source against another column of the query result
* `value` (AlertV2OperandValue) - Compare the source against a fixed value

### AlertV2OperandColumn
* `aggregation` (string) - Aggregation applied to the column across all rows of the query result: `SUM`, `COUNT`, `COUNT_DISTINCT`, `AVG`, `MEDIAN`, `MIN`, `MAX` or `STDDEV`.
  If not set, the value in the first row is used, which is equivalent to using `First row` in the UI. Possible values are: `AVG`, `COUNT`, `COUNT_DISTINCT`, `MAX`, `MEDIAN`, `MIN`, `STDDEV`, `SUM`
* `display` (string) - Label for the column in notifications. Defaults to `name`
* `name` (string) - Name of the column in the query result

### AlertV2OperandValue
* `bool_value` (boolean) - Boolean threshold, for boolean columns
* `double_value` (number) - Numeric threshold, for numeric columns
* `string_value` (string) - String threshold, for string or timestamp columns

### AlertV2RunAs
* `service_principal_name` (string) - Application ID of an active service principal. Setting this field requires the `servicePrincipal/user` role
* `user_name` (string) - The email of an active workspace user. Can only set this field to their own email

### AlertV2Subscription
* `destination_id` (string) - ID of the notification destination to notify, as returned by the Notification Destinations API
* `user_email` (string) - Email address of the workspace user to notify

### CronSchedule
* `pause_status` (string) - Indicates whether this schedule is paused. While paused, the alert is not evaluated. Possible values are: `PAUSED`, `UNPAUSED`
* `quartz_cron_schedule` (string) - A cron expression using quartz syntax that specifies the schedule for this alert.
  Should use the quartz format described here: http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
* `timezone_id` (string) - A Java timezone id. The schedule will be resolved using this timezone.
  This will be combined with the quartz_cron_schedule to determine the schedule.
  See https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html for details