---
subcategory: "Databricks SQL"
---
# databricks_alert_v2 Data Source
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
* `id` (string, required) - UUID identifying the alert

## Attributes
The following attributes are exported:
* `create_time` (string) - The timestamp indicating when the alert was created
* `custom_description` (string) - Custom description for the alert. support mustache template
* `custom_summary` (string) - Custom summary for the alert. support mustache template
* `display_name` (string) - The display name of the alert
* `evaluation` (AlertV2Evaluation) - 
* `id` (string) - UUID identifying the alert
* `lifecycle_state` (string) - Indicates whether the query is trashed. Possible values are: `ACTIVE`, `TRASHED`
* `owner_user_name` (string) - The owner's username. This field is set to "Unavailable" if the user has been deleted
* `parent_path` (string) - The workspace path of the folder containing the alert. Can only be set on create, and cannot be updated
* `query_text` (string) - Text of the query to be run
* `run_as_user_name` (string) - The run as username or application ID of service principal.
  On Create and Update, this field can be set to application ID of an active service principal. Setting this field requires the servicePrincipal/user role
* `schedule` (CronSchedule) - 
* `update_time` (string) - The timestamp indicating when the alert was updated
* `warehouse_id` (string) - ID of the SQL warehouse attached to the alert

### AlertV2Evaluation
* `comparison_operator` (string) - Operator used for comparison in alert evaluation. Possible values are: `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NOT_NULL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`
* `empty_result_state` (string) - Alert state if result is empty. Possible values are: `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`
* `last_evaluated_at` (string) - Timestamp of the last evaluation
* `notification` (AlertV2Notification) - User or Notification Destination to notify when alert is triggered
* `source` (AlertV2OperandColumn) - Source column from result to use to evaluate alert
* `state` (string) - Latest state of alert evaluation. Possible values are: `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`
* `threshold` (AlertV2Operand) - Threshold to user for alert evaluation, can be a column or a value

### AlertV2Notification
* `notify_on_ok` (boolean) - Whether to notify alert subscribers when alert returns back to normal
* `retrigger_seconds` (integer) - Number of seconds an alert must wait after being triggered to rearm itself. After rearming, it can be triggered again. If 0 or not specified, the alert will not be triggered again
* `subscriptions` (list of AlertV2Subscription) - 

### AlertV2Operand
* `column` (AlertV2OperandColumn) - 
* `value` (AlertV2OperandValue) - 

### AlertV2OperandColumn
* `aggregation` (string) - . Possible values are: `AVG`, `COUNT`, `COUNT_DISTINCT`, `MAX`, `MEDIAN`, `MIN`, `STDDEV`, `SUM`
* `display` (string) - 
* `name` (string) - 

### AlertV2OperandValue
* `bool_value` (boolean) - 
* `double_value` (number) - 
* `string_value` (string) - 

### AlertV2Subscription
* `destination_id` (string) - 
* `user_email` (string) - 

### CronSchedule
* `pause_status` (string) - Indicate whether this schedule is paused or not. Possible values are: `PAUSED`, `UNPAUSED`
* `quartz_cron_schedule` (string) - A cron expression using quartz syntax that specifies the schedule for this pipeline.
  Should use the quartz format described here: http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
* `timezone_id` (string) - A Java timezone id. The schedule will be resolved using this timezone.
  This will be combined with the quartz_cron_schedule to determine the schedule.
  See https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html for details