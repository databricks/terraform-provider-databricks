---
subcategory: "Databricks SQL"
---
# databricks_alert_v2 Resource
[![Public Preview](https://img.shields.io/badge/Release_Stage-Public_Preview-yellowgreen)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
resource "databricks_alert_v2" "basic_alert" {
  display_name = "High Error Rate Alert"
  query_text   = "SELECT count(*) as error_count FROM logs WHERE level = 'ERROR' AND timestamp > now() - interval 1 hour"
  warehouse_id = "a7066a8ef796be84"
  parent_path  = "/Users/user@example.com"

  evaluation = {
    source = {
      name        = "error_count"
      display     = "Error Count"
      aggregation = "COUNT"
    }
    comparison_operator = "GREATER_THAN"
    threshold = {
      value = {
        double_value = 100
      }
    }
    empty_result_state = "OK"

    notification = {
      subscriptions = [
        {
          user_email = "user@example.com"
        }
      ]
      notify_on_ok = true
    }
  }

  schedule = {
    quartz_cron_schedule = "0 0/15 * * * ?" # Every 15 minutes
    timezone_id          = "America/Los_Angeles"
    pause_status         = "UNPAUSED"
  }
}
```


## Arguments
The following arguments are supported:
* `display_name` (string, required) - The display name of the alert
* `evaluation` (AlertV2Evaluation, required)
* `query_text` (string, required) - Text of the query to be run
* `schedule` (CronSchedule, required)
* `warehouse_id` (string, required) - ID of the SQL warehouse attached to the alert
* `custom_description` (string, optional) - Custom description for the alert. support mustache template
* `custom_summary` (string, optional) - Custom summary for the alert. support mustache template
* `parent_path` (string, optional) - The workspace path of the folder containing the alert. Can only be set on create, and cannot be updated
* `run_as` (AlertV2RunAs, optional) - Specifies the identity that will be used to run the alert.
  This field allows you to configure alerts to run as a specific user or service principal.
  - For user identity: Set `user_name` to the email of an active workspace user. Users can only set this to their own email.
  - For service principal: Set `service_principal_name` to the application ID. Requires the `servicePrincipal/user` role.
  If not specified, the alert will run as the request user
* `run_as_user_name` (string, optional, deprecated) - The run as username or application ID of service principal.
  On Create and Update, this field can be set to application ID of an active service principal. Setting this field requires the servicePrincipal/user role.
  Deprecated: Use `run_as` field instead. This field will be removed in a future release

### AlertV2Evaluation
* `comparison_operator` (string, required) - Operator used for comparison in alert evaluation. Possible values are: `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NOT_NULL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`
* `source` (AlertV2OperandColumn, required) - Source column from result to use to evaluate alert
* `empty_result_state` (string, optional) - Alert state if result is empty. Please avoid setting this field to be `UNKNOWN` because `UNKNOWN` state is planned to be deprecated. Possible values are: `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`
* `notification` (AlertV2Notification, optional) - User or Notification Destination to notify when alert is triggered
* `threshold` (AlertV2Operand, optional) - Threshold to user for alert evaluation, can be a column or a value

### AlertV2Notification
* `notify_on_ok` (boolean, optional) - Whether to notify alert subscribers when alert returns back to normal
* `retrigger_seconds` (integer, optional) - Number of seconds an alert waits after being triggered before it is allowed to send another notification.
  If set to 0 or omitted, the alert will not send any further notifications after the first trigger
  Setting this value to 1 allows the alert to send a notification on every evaluation where the condition is met, effectively making it always retrigger for notification purposes
* `subscriptions` (list of AlertV2Subscription, optional)

### AlertV2Operand
* `column` (AlertV2OperandColumn, optional)
* `value` (AlertV2OperandValue, optional)

### AlertV2OperandColumn
* `name` (string, required)
* `aggregation` (string, optional) - If not set, the behavior is equivalent to using `First row` in the UI. Possible values are: `AVG`, `COUNT`, `COUNT_DISTINCT`, `MAX`, `MEDIAN`, `MIN`, `STDDEV`, `SUM`
* `display` (string, optional)

### AlertV2OperandValue
* `bool_value` (boolean, optional)
* `double_value` (number, optional)
* `string_value` (string, optional)

### AlertV2RunAs
* `service_principal_name` (string, optional) - Application ID of an active service principal. Setting this field requires the `servicePrincipal/user` role
* `user_name` (string, optional) - The email of an active workspace user. Can only set this field to their own email

### AlertV2Subscription
* `destination_id` (string, optional)
* `user_email` (string, optional)

### CronSchedule
* `quartz_cron_schedule` (string, required) - A cron expression using quartz syntax that specifies the schedule for this pipeline.
  Should use the quartz format described here: http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
* `timezone_id` (string, required) - A Java timezone id. The schedule will be resolved using this timezone.
  This will be combined with the quartz_cron_schedule to determine the schedule.
  See https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html for details
* `pause_status` (string, optional) - Indicate whether this schedule is paused or not. Possible values are: `PAUSED`, `UNPAUSED`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - The timestamp indicating when the alert was created
* `effective_run_as` (AlertV2RunAs) - The actual identity that will be used to execute the alert.
  This is an output-only field that shows the resolved run-as identity after applying
  permissions and defaults
* `id` (string) - UUID identifying the alert
* `lifecycle_state` (string) - Indicates whether the query is trashed. Possible values are: `ACTIVE`, `DELETED`
* `owner_user_name` (string) - The owner's username. This field is set to "Unavailable" if the user has been deleted
* `update_time` (string) - The timestamp indicating when the alert was updated

### AlertV2Evaluation
* `last_evaluated_at` (string) - Timestamp of the last evaluation
* `state` (string) - Latest state of alert evaluation. Possible values are: `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "id"
  to = databricks_alert_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_alert_v2.this "id"
```
