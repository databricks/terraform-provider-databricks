---
subcategory: "Workspace"
---
# databricks_notification_destination Resource

[API Documentation](https://docs.databricks.com/api/workspace/notificationdestinations)

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
`Databricks app in Slack` notification destination (the Databricks app must already be installed in your Slack workspace and connected to your Databricks workspace):

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    databricks_app_slack {
      team_id    = "T0123456789"
      channel_id = "C0123456789"
    }
  }
}
```
`Databricks app in Microsoft Teams` notification destination (the Databricks app must already be installed in your Microsoft 365 tenant and connected to your Databricks workspace):

```hcl
resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    databricks_app_teams {
      tenant_id  = "00000000-0000-0000-0000-000000000000"
      team_id    = "00000000-0000-0000-0000-000000000000"
      channel_id = "19:abc123@thread.tacv2"
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
    * `url` - (Optional) The Slack webhook URL.
    * `channel_id`  - (Optional) Slack channel ID for notifications.
    * `oauth_token` - (Optional) OAuth token for Slack authentication.
  * `databricks_app_slack` - The configuration for a Slack channel reached through the Databricks app. The app must already be installed in the Slack workspace and connected to the Databricks workspace. It must contain the following:
    * `team_id` - (Optional) Slack workspace (team) ID the Databricks app is installed in.
    * `channel_id` - (Optional) Slack channel ID notifications are posted to.
  * `databricks_app_teams` - The configuration for a Microsoft Teams channel reached through the Databricks app. The app must already be installed in the Microsoft 365 tenant and connected to the Databricks workspace. It must contain the following:
    * `tenant_id` - (Optional) Microsoft Entra tenant ID of the Microsoft 365 tenant the app is installed in.
    * `team_id` - (Optional) Microsoft Entra group ID of the team that owns the channel.
    * `channel_id` - (Optional) Teams channel ID notifications are posted to.
  * `pagerduty` - The PagerDuty configuration of the Notification Destination. It must contain the following:
    * `integration_key` - (Required) The PagerDuty integration key.
  * `microsoft_teams` - The Microsoft Teams configuration of the Notification Destination. It must contain the following:
    * `url` - (Optional) The Microsoft Teams webhook URL.
    * `app_id` - (Optional) App ID for Microsoft Teams App.
    * `auth_secret` - (Optional) Secret for Microsoft Teams App authentication.
    * `channel_url` - (Optional) Channel URL for Microsoft Teams App.
    * `tenant_id` - (Optional) Tenant ID for Microsoft Teams App.
  * `generic_webhook` - The Generic Webhook configuration of the Notification Destination. It must contain the following:
    * `url` - (Required) The Generic Webhook URL.
    * `username` - (Optional) The username for basic authentication.
    * `password` - (Optional) The password for basic authentication.

* `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  * `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

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
