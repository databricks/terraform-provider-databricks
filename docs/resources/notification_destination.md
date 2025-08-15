---
subcategory: "Workspace"
---
# databricks_notification_destination Resource

This resource allows you to manage [Notification Destinations](https://docs.databricks.com/api/workspace/notificationdestinations). Notification destinations are used to send notifications for query alerts and jobs to destinations outside of Databricks. Only workspace admins can create, update, and delete notification destinations.

-> This resource can only be used with a workspace-level provider!

## Example Usage

`Email` notification destination:

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    email {
      addresses = ["abc@gmail.com"]
    }
  }
}
```
`Slack` notification destination:

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    slack {
      url = "https://hooks.slack.com/services/..."
    }
  }
}
```
`PagerDuty` notification destination:

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    pagerduty {
      integration_key = "xxxxxx"
    }
  }
}
```
`Microsoft Teams` notification destination:

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    microsoft_teams {
      url = "https://outlook.office.com/webhook/..."
    }
  }
}
```
`Generic Webhook` notification destination:

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    generic_webhook {
      url      = "https://example.com/webhook"
      username = "username" // Optional
      password = "password" // Optional
    }
  }
}
```


## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the Notification Destination.
* `config` - (Required) The configuration of the Notification Destination. It must contain exactly one of the following blocks:
  * `email` - The email configuration of the Notification Destination. It must contain the following:
    * `addresses` - (Required) The list of email addresses to send notifications to.
  * `slack` - The Slack configuration of the Notification Destination. It must contain the following:
    * `url` - (Required) The Slack webhook URL.
    * `channel_id`  - (Optional) Slack channel ID for notifications.
    * `oauth_token` - (Optional) OAuth token for Slack authentication.
  * `pagerduty` - The PagerDuty configuration of the Notification Destination. It must contain the following:
    * `integration_key` - (Required) The PagerDuty integration key.
  * `microsoft_teams` - The Microsoft Teams configuration of the Notification Destination. It must contain the following:
    * `url` - (Required) The Microsoft Teams webhook URL.
  * `generic_webhook` - The Generic Webhook configuration of the Notification Destination. It must contain the following:
    * `url` - (Required) The Generic Webhook URL.
    * `username` - (Optional) The username for basic authentication.
    * `password` - (Optional) The password for basic authentication.

-> **NOTE** If the type of notification destination is changed, the existing notification destination will be deleted and a new notification destination will be created with the new type.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID of the Notification Destination.
* `destination_type` - the type of Notification Destination.

## Import

This resource can be imported by notification ID:

```hcl
import {
  to = databricks_notification_destination.this
  id = "<notification-id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_notification_destination.this <notification-id>
```
