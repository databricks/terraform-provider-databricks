---
subcategory: "Databricks SQL"
---
# databricks_alerts_v2 Data Source
The SQL Alerts v2 data source allows you to retrieve a list of alerts in Databricks SQL that are accessible to the current user. This data source returns alerts ordered by their creation time.

You can use this data source to:
- Get a comprehensive list of all alerts in your workspace
- Monitor and audit alert configurations across your workspace

### Pagination
The data source supports pagination to efficiently retrieve alerts. You can control the page size to limit the number of results returned in a single request.

## Example Usage
### List All Alerts
This example retrieves all alerts accessible to the current user:

```hcl
data "databricks_alert_v2" "all" {}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - 



## Attributes
This data source exports a single attribute, `results`. It is a list of resources, each with the following attributes:
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