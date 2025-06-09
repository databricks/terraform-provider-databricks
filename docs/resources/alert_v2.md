---
subcategory: "Databricks SQL"
---
# databricks_alert_v2 Resource


## Example Usage


## Arguments
The following arguments are supported:
* `custom_description` (string, optional) - Custom description for the alert. support mustache template
* `custom_summary` (string, optional) - Custom summary for the alert. support mustache template
* `display_name` (string, optional) - The display name of the alert
* `evaluation` (AlertV2Evaluation, optional) - 
* `parent_path` (string, optional) - The workspace path of the folder containing the alert. Can only be set on create, and cannot be updated
* `query_text` (string, optional) - Text of the query to be run
* `schedule` (CronSchedule, optional) - 
* `warehouse_id` (string, optional) - ID of the SQL warehouse attached to the alert

### AlertV2Evaluation
* `comparison_operator` (string, optional) - Operator used for comparison in alert evaluation. Possible values are: EQUAL, GREATER_THAN, GREATER_THAN_OR_EQUAL, IS_NOT_NULL, IS_NULL, LESS_THAN, LESS_THAN_OR_EQUAL, NOT_EQUAL
* `empty_result_state` (string, optional) - Alert state if result is empty. Possible values are: ERROR, OK, TRIGGERED, UNKNOWN
* `notification` (AlertV2Notification, optional) - User or Notification Destination to notify when alert is triggered
* `source` (AlertV2OperandColumn, optional) - Source column from result to use to evaluate alert
* `threshold` (AlertV2Operand, optional) - Threshold to user for alert evaluation, can be a column or a value

### AlertV2Notification
* `notify_on_ok` (boolean, optional) - Whether to notify alert subscribers when alert returns back to normal
* `retrigger_seconds` (integer, optional) - Number of seconds an alert must wait after being triggered to rearm itself. After rearming, it can be triggered again. If 0 or not specified, the alert will not be triggered again
* `subscriptions` (list of AlertV2Subscription, optional) - 

### AlertV2Operand
* `column` (AlertV2OperandColumn, optional) - 
* `value` (AlertV2OperandValue, optional) - 

### AlertV2OperandColumn
* `aggregation` (string, optional) - . Possible values are: AVG, COUNT, COUNT_DISTINCT, MAX, MEDIAN, MIN, STDDEV, SUM
* `display` (string, optional) - 
* `name` (string, optional) - 

### AlertV2OperandValue
* `bool_value` (boolean, optional) - 
* `double_value` (number, optional) - 
* `string_value` (string, optional) - 

### AlertV2Subscription
* `destination_id` (string, optional) - 
* `user_email` (string, optional) - 

### CronSchedule
* `pause_status` (string, optional) - Indicate whether this schedule is paused or not. Possible values are: PAUSED, UNPAUSED
* `quartz_cron_schedule` (string, optional) - A cron expression using quartz syntax that specifies the schedule for this pipeline.
  Should use the quartz format described here: http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
* `timezone_id` (string, optional) - A Java timezone id. The schedule will be resolved using this timezone.
  This will be combined with the quartz_cron_schedule to determine the schedule.
  See https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html for details

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - The timestamp indicating when the alert was created
* `id` (string) - UUID identifying the alert
* `lifecycle_state` (string) - Indicates whether the query is trashed. Possible values are: ACTIVE, TRASHED
* `owner_user_name` (string) - The owner's username. This field is set to "Unavailable" if the user has been deleted
* `run_as_user_name` (string) - The run as username. This field is set to "Unavailable" if the user has been deleted
* `update_time` (string) - The timestamp indicating when the alert was updated

### AlertV2Evaluation
* `last_evaluated_at` (string) - Timestamp of the last evaluation
* `state` (string) - Latest state of alert evaluation. Possible values are: ERROR, OK, TRIGGERED, UNKNOWN

## Import
As of terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = id
  to = databricks_alert_v2.this
}
```

If you are using an older version of terraform, you can import the resource using cli as follows:
```sh
$ terraform import databricks_alert_v2 id
```