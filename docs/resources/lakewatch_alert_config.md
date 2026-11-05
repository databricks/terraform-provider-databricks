---
subcategory: "Lakewatch"
---
# databricks_lakewatch_alert_config Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the parent
* `default_security_alert_collation_window` (string, optional) - Window for collating Signals into security alerts. Duration
  string — compact day/hour/minute/second components, e.g. `1d`, `12h`
  (see `Schedule.at_least_every` in common.proto for the grammar). Reads
  return the value as supplied. When unset or empty, the window is `1h`; the window in
  effect is reported in `effective_default_security_alert_collation_window`
* `default_system_alert_collation_window` (string, optional) - Window for collating SystemEvents into system alerts. Duration
  string — compact day/hour/minute/second components, e.g. `1d`, `12h`
  (see `Schedule.at_least_every` in common.proto for the grammar). Reads
  return the value as supplied. When unset or empty, the window is `1d`; the window in
  effect is reported in `effective_default_system_alert_collation_window`
* `security_alert_export_config` (ExportConfig, optional) - Export destination for security alerts. When unset, export is off and reads
  return `{"enabled": false}`
* `system_alert_export_config` (ExportConfig, optional) - Export destination for system alerts. When unset, export is off and reads
  return `{"enabled": false}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### ExportConfig
* `enabled` (boolean, optional) - Whether export is currently active
* `jira` (JiraExportConfig, optional)
* `service_now` (ServiceNowExportConfig, optional)
* `slack` (SlackExportConfig, optional)
* `tines` (TinesExportConfig, optional)
* `webhook` (WebhookExportConfig, optional)
* `workspace_id` (string, optional) - Deprecated and ignored: exports always egress through the request's own
  workspace. Kept so configs stored by older builds still round-trip

### JiraExportConfig
* `api_token_uc_secret` (string, required) - Three-level Unity Catalog name (`catalog.schema.secret`) of the secret holding
  the Jira API token, paired with `user_email` as the Jira Cloud basic-auth
  credential (`email:token`). Only the name is stored; the token is read when an
  alert is exported, so rotating the secret takes effect without editing this config
* `default_issue_type` (string, required) - Default issue type for created tickets (e.g. `Task`, `Incident`). Jira issue creation
  requires an issue type, so it is validated at config-save time rather than failing the
  export at Jira later
* `instance_url` (string, required) - Jira instance URL (e.g. `https://acme.atlassian.net`)
* `project_key` (string, required) - Jira project key the exported issues are created under
* `user_email` (string, required) - User email paired with the API token for Jira Cloud basic auth

### Principal
* `display_name` (string, optional) - Display name (user or SP name)
* `email` (string, optional) - Email (regular users only)
* `id` (string, optional) - User / service-principal ID
* `principal_type` (string, optional) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`

### ServiceNowExportConfig
* `client_id` (string, required) - OAuth client ID of the registered ServiceNow application. An identifier rather
  than secret material — ServiceNow shows it on the application record — so it
  is stored and read back verbatim
* `client_secret_uc_secret` (string, required) - Three-level Unity Catalog name (`catalog.schema.secret`) of the secret holding
  the OAuth client secret. Only the name is stored; the secret is read when an alert
  is exported, so rotating it takes effect without editing this config
* `instance_id` (string, required) - ServiceNow instance ID — the host label only, not the full URL. For an
  instance reachable at `https://acme.service-now.com`, this is `acme`

### SlackExportConfig
* `bot_token_uc_secret` (string, required) - Three-level Unity Catalog name (`catalog.schema.secret`) of the secret holding
  the Slack bot token. Only the name is stored; the token is read when an alert is
  exported, so rotating the secret takes effect without editing this config
* `channel` (string, required) - Channel the alert is posted to: a channel ID (e.g. `C0123456789`) or a channel
  name without the leading `#` (e.g. `soc-alerts`). The value is stored as sent, so
  a leading `#`, surrounding whitespace, or a value naming more than one channel is
  rejected. For private channels, prefer the channel ID — Slack does not guarantee
  name resolution for those.
  
  The Slack app must be a member of the channel and hold the `chat:write` scope

### TinesExportConfig
* `webhook_secret_uc_secret` (string, required) - Three-level Unity Catalog name (`catalog.schema.secret`) of the secret holding the Tines
  webhook secret. Only the name is stored; the secret is read when an alert is exported, so
  rotating it takes effect without editing this config. Required: without it the destination
  would be an unauthenticated endpoint
* `webhook_uri` (string, required) - Tines webhook action URL, typically of the form
  `https://<tenant>.tines.com/webhook/<path>`. Any HTTPS host is accepted, so
  that a Tines tenant on a custom domain works.
  
  The secret is sent as a header, so it must not be embedded in the URL: a
  trailing secret path segment (`/webhook/<path>/<secret>`) and a `secret`
  query parameter are both rejected. API-path-prefix URLs (`/api/public/...`)
  are not supported
* `export_updates` (boolean, optional) - Whether a later change to an already-exported alert is delivered too, as a second
  `lakewatch.alert.updated` event to the same webhook action (a close is such an update,
  carrying a terminal `alert_state`). Creation export is unaffected either way.
  
  Opt-in, and unset means off: Tines is an event transport rather than a system whose record is
  edited in place, so an update is an extra event the receiving story has to correlate to the
  alert it already handled. A destination configured before this event existed must not start
  receiving more without its owner asking

### WebhookExportConfig
* `export_updates` (boolean, optional) - Whether a later change to an already-exported alert is delivered too, as a second
  `lakewatch.alert.updated` event to the same webhook URL (a close is such an update, carrying a
  terminal `alert_state`). Creation export is unaffected either way.
  
  Opt-in, and unset means off: a webhook is an event transport rather than a system whose record
  is edited in place, so an update is an extra event the receiving consumer has to correlate to
  the alert it already handled. A destination configured before this event existed must not start
  receiving more without its owner asking
* `url` (string, optional) - Webhook destination the server POSTs payloads to, in plaintext. Userinfo
  (`user:password@`) is sent as HTTP Basic auth

## Attributes
In addition to the above arguments, the following attributes are exported:
* `effective_default_security_alert_collation_window` (string) - The security alert collation window in effect:
  `default_security_alert_collation_window` when set, otherwise `1h`
* `effective_default_system_alert_collation_window` (string) - The system alert collation window in effect:
  `default_system_alert_collation_window` when set, otherwise `1d`
* `etag` (string) - Opaque entity tag for optimistic concurrency. Returned on every read;
  pass it back as the Update request's `etag` to make the update
  conditional on the resource being unchanged
* `name` (string)
* `update_time` (string)
* `updated_by` (Principal)
* `version` (integer) - Monotonic version, bumped on each update. Informational; use `etag` for
  optimistic concurrency

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_lakewatch_alert_config.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_lakewatch_alert_config.this "name"
```