---
subcategory: "Workspace"
---
# databricks_notification_destination Resource

[API Documentation](https://docs.databricks.com/api/workspace/notificationdestinations)

This resource allows you to manage [Notification Destinations](https://docs.databricks.com/api/workspace/notificationdestinations). Notification destinations are used to send notifications for query alerts and jobs to destinations outside of Databricks. Only workspace admins can create, update, and delete notification destinations.

-> This resource can only be used with a workspace-level provider!

-> **Note** Write-only arguments (`url_wo`, `username_wo`, `password_wo`, `oauth_token_wo`, `auth_secret_wo` and `integration_key_wo`) are available to supply webhook URLs, credentials and keys without storing them in the Terraform plan or state. Write-only arguments are supported in HashiCorp Terraform 1.11.0 and later. [Learn more](https://developer.hashicorp.com/terraform/language/manage-sensitive-data/ephemeral#write-only-arguments).

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

`Generic Webhook` notification destination with a write-only URL and password, read from an ephemeral resource such as AWS Secrets Manager, so they're never stored in the Terraform state:

```hcl
ephemeral "aws_secretsmanager_secret_version" "webhook" {
  secret_id = "prod/databricks/job-webhook"
}

locals {
  webhook = jsondecode(ephemeral.aws_secretsmanager_secret_version.webhook.secret_string)
}

resource "databricks_notification_destination" "ndresource" {
  display_name = "Notification Destination"
  config {
    generic_webhook {
      url_wo      = local.webhook.url
      username    = "username"
      password_wo = local.webhook.password
    }
  }
  config_wo_version = 1 // increment to apply a new URL or password
}
```


## Argument Reference

The following arguments are supported:

* `display_name` - (Required) The display name of the Notification Destination.
* `config` - (Required) The configuration of the Notification Destination. It must contain exactly one of the following blocks:
  * `email` - The email configuration of the Notification Destination. It must contain the following:
    * `addresses` - (Required) The list of email addresses to send notifications to.
  * `slack` - The Slack configuration of the Notification Destination. It must contain the following:
    * `url` - (Optional) The Slack webhook URL. Exactly one of `url` or `url_wo` must be set.
    * `url_wo` - (Optional) Write-only Slack webhook URL, not stored in the Terraform state. Requires `config_wo_version`.
    * `channel_id`  - (Optional) Slack channel ID for notifications.
    * `oauth_token` - (Optional) OAuth token for Slack authentication.
    * `oauth_token_wo` - (Optional) Write-only OAuth token for Slack authentication, not stored in the Terraform state. Conflicts with `oauth_token`. Requires `config_wo_version`.
  * `databricks_app_slack` - The configuration for a Slack channel reached through the Databricks app. The app must already be installed in the Slack workspace and connected to the Databricks workspace. It must contain the following:
    * `team_id` - (Optional) Slack workspace (team) ID the Databricks app is installed in.
    * `channel_id` - (Optional) Slack channel ID notifications are posted to.
  * `databricks_app_teams` - The configuration for a Microsoft Teams channel reached through the Databricks app. The app must already be installed in the Microsoft 365 tenant and connected to the Databricks workspace. It must contain the following:
    * `tenant_id` - (Optional) Microsoft Entra tenant ID of the Microsoft 365 tenant the app is installed in.
    * `team_id` - (Optional) Microsoft Entra group ID of the team that owns the channel.
    * `channel_id` - (Optional) Teams channel ID notifications are posted to.
  * `pagerduty` - The PagerDuty configuration of the Notification Destination. It must contain the following:
    * `integration_key` - (Optional) The PagerDuty integration key. Exactly one of `integration_key` or `integration_key_wo` must be set.
    * `integration_key_wo` - (Optional) Write-only PagerDuty integration key, not stored in the Terraform state. Requires `config_wo_version`.
  * `microsoft_teams` - The Microsoft Teams configuration of the Notification Destination. It must contain the following:
    * `url` - (Optional) The Microsoft Teams webhook URL. Exactly one of `url` or `url_wo` must be set.
    * `url_wo` - (Optional) Write-only Microsoft Teams webhook URL, not stored in the Terraform state. Requires `config_wo_version`.
    * `app_id` - (Optional) App ID for Microsoft Teams App.
    * `auth_secret` - (Optional) Secret for Microsoft Teams App authentication.
    * `auth_secret_wo` - (Optional) Write-only secret for Microsoft Teams App authentication, not stored in the Terraform state. Conflicts with `auth_secret`. Requires `config_wo_version`.
    * `channel_url` - (Optional) Channel URL for Microsoft Teams App.
    * `tenant_id` - (Optional) Tenant ID for Microsoft Teams App.
  * `generic_webhook` - The Generic Webhook configuration of the Notification Destination. It must contain the following:
    * `url` - (Optional) The Generic Webhook URL. Exactly one of `url` or `url_wo` must be set.
    * `url_wo` - (Optional) Write-only Generic Webhook URL, not stored in the Terraform state. Use it when the URL contains a token. Requires `config_wo_version`.
    * `username` - (Optional) The username for basic authentication.
    * `username_wo` - (Optional) Write-only username for basic authentication, not stored in the Terraform state. Conflicts with `username`. Requires `config_wo_version`.
    * `password` - (Optional) The password for basic authentication.
    * `password_wo` - (Optional) Write-only password for basic authentication, not stored in the Terraform state. Conflicts with `password`. Requires `config_wo_version`.
* `config_wo_version` - (Optional) Version of the write-only `*_wo` arguments in `config`. Required when any of them is set. Terraform can't detect changes to write-only values, so increment this value (starting at 1) to apply new ones. Changing it updates the notification destination in place, keeping its ID.

  ~> The API replaces the whole configuration block on every update, so the provider sends the current write-only values from the configuration with every update of the notification destination, not only when `config_wo_version` changes. Terraform can't detect changes to these values made outside of Terraform.

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
