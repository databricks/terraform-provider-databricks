---
subcategory: "Workspace"
---
# databricks_notification_destinations Data Source

This data source allows you to retrieve information about [Notification Destinations](https://docs.databricks.com/api/workspace/notificationdestinations). Notification Destinations are used to send notifications for query alerts and jobs to external systems such as email, Slack, Microsoft Teams, PagerDuty, or generic webhooks. 

## Example Usage


```hcl
resource "databricks_notification_destination" "email" {
  display_name = "Email Destination"
  config {
    email {
      addresses = ["abc@gmail.com"]
    }
  }
}

resource "databricks_notification_destination" "slack" {
  display_name = "Slack Destination"
  config {
    slack {
      url = "https://hooks.slack.com/services/..."
    }
  }
}

data  "databricks_notification_destinations" "this" {} # Lists all notification desitnations

data "databricks_notification_destinations" "filtered_notification" {
    display_name_contains = "Destination"
    type = "EMAIL"
}

```


## Argument Reference

The following arguments are supported:

* `display_name_contains` - (Optional) A **case-insensitive** substring to filter Notification Destinations by their display name. 
* `type` - (Optional) The type of the Notification Destination to filter by. Valid values are: 
  * `EMAIL` - Filters Notification Destinations of type Email. 
  * `MICROSOFT_TEAMS` - Filters Notification Destinations of type Microsoft Teams. 
  * `PAGERDUTY` - Filters Notification Destinations of type PagerDuty.
  * `SLACK` - Filters Notification Destinations of type Slack. 
  * `WEBHOOK` - Filters Notification Destinations of type Webhook. 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `notification_destinations` - A list of Notification Destinations matching the specified criteria. Each element contains the following attributes: 
    * `id` - The unique ID of the Notification Destination.
    * `display_name` - The display name of the Notification Destination.
    * `destination_type` - The type of the notification destination. Possible values are `EMAIL`, `MICROSOFT_TEAMS`, `PAGERDUTY`, `SLACK`, or `WEBHOOK`.

If no matches are found, an empty list will be returned.