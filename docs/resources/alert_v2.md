---
subcategory: "Databricks SQL"
---
# databricks_alert_v2 Resource
The Alert v2 resource allows you to manage SQL alerts in Databricks SQL. Alerts monitor query results and notify you when specific conditions are met.

Alerts run on a schedule and evaluate query results against defined thresholds. When an alert is triggered, notifications can be sent to specified users or destinations.

### Alert Evaluation
Alerts support various comparison operators and can evaluate query results against fixed values or other columns. You can configure alerts to trigger when results are empty or when specific conditions are met.

### Notifications
When an alert is triggered, notifications can be sent to:
- User email addresses
- Notification destinations (configured separately)

You can also configure alerts to notify subscribers when the alert returns to normal state and set a retrigger interval to control how frequently the alert can be triggered.

### Scheduling
Alerts use Quartz cron syntax for scheduling. You can specify the timezone and pause status for the schedule.

## Example Usage
### Basic Alert Example
This example creates a basic alert that monitors a query and sends notifications to a user when the value exceeds a threshold:

```hcl
resource "databricks_sql_alert" "basic_alert" {
  display_name = "High Error Rate Alert"
  query_text   = "SELECT count(*) as error_count FROM logs WHERE level = 'ERROR' AND timestamp > now() - interval 1 hour"
  warehouse_id = "a7066a8ef796be84"
  parent_path  = "/Users/user@example.com"
  
  evaluation {
    source {
      name        = "error_count"
      display     = "Error Count"
      aggregation = "COUNT"
    }
    comparison_operator = "GREATER_THAN"
    threshold {
      value {
        double_value = 100
      }
    }
    empty_result_state = "OK"
    
    notification {
      subscriptions {
        user_email = "user@example.com"
      }
      notify_on_ok = true
    }
  }
  
  schedule {
    quartz_cron_schedule = "0 0/15 * * * ?"  # Every 15 minutes
    timezone_id          = "America/Los_Angeles"
    pause_status         = "UNPAUSED"
  }
}
```


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