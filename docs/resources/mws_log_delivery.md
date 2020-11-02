# databricks_mws_log_delivery Resource

-> **Note** This resource has evolving API, which may change in future versions of provider.

This resource configures delivery of various log types from databricks workspaces. Backend does not support removal of log delivery configuration, so provider handles it through simply disabling the config.

## Example Usage

Configuring usage logs:

```hcl
resource "databricks_mws_log_delivery" "usage_logs" {
    account_id   = var.account_id
    credentials_id = databricks_mws_credentials.log_writer.credentials_id
    storage_configuration_id = databricks_mws_storage_configurations.log_bucket.storage_configuration_id
    config_name = "Usage Logs"
    log_type = "BILLABLE_USAGE"
    output_format = "CSV"
}
```

Configuring audit logs:

```hcl
resource "databricks_mws_log_delivery" "audit_logs" {
    account_id   = var.account_id
    credentials_id = databricks_mws_credentials.log_writer.credentials_id
    storage_configuration_id = databricks_mws_storage_configurations.log_bucket.storage_configuration_id
    config_name = "Audit Logs"
    log_type = "AUDIT_LOGS"
    output_format = "JSON"
}
```

## Argument reference

* `account_id` - The Databricks account ID that hosts the log delivery configuration.
* `config_name`	- The optional human-readable name of the log delivery configuration. Defaults to empty.
* `log_type` - The type of log delivery. [BILLABLE_USAGE](https://docs.databricks.com/administration-guide/account-settings/usage.html#download-usage-as-a-csv-file) and `AUDIT_LOGS` are supported
* `output_format` - The file type of log delivery. Currently `CSV` and `JSON` are supported.
* `credentials_id` - The ID for a Databricks [credential configuration](mws_credentials.md) that represents the AWS IAM role with policy and trust relationship as described in the main billable usage documentation page.
* `storage_configuration_id` - The ID for a Databricks [storage configuration](mws_storage_configurations.md) that represents the S3 bucket with bucket policy as described in the main billable usage documentation page.
* `workspace_ids_filter` - (Optional) The optional filter for workspaces. By default, this log configuration applies to all workspaces associated with your account ID. For some types of deployments there is only one workspace per account ID so this field is unnecessary. If your account is on the E2 version of the platform or on a select custom plan that allows multiple workspaces per account, you may have multiple workspaces associated with your account ID. You can optionally set this field to array of workspace IDs (each one is an int64) that this configuration applies to. If you plan to use different log delivery configurations for different workspaces, set this explicitly rather than leaving it blank. If you leave this blank and your account ID is associated in the future with additional workspaces, this configuration also applies to the new workspaces.
* `delivery_path_prefix` - (Optional) The optional delivery path prefix within AWS S3 storage. Defaults to empty, which means that logs are delivered to the root of the bucket. This must be a valid S3 object key. This must not start or end with a slash character.
* `delivery_start_time`	- (Optional) The optional start month and year for delivery, specified in YYYY-MM format. Defaults to current year and month. Usage is not available before 2019-03.

## Attribute reference

Following attributes are exported

* `config_id` - Databricks log delivery configuration ID.

## Import

This resource can be imported by specifying a combination of an account id and log config id separated by `|`:

```bash
$ terraform import databricks_mws_log_delivery.usage "<account-id>|<log-config-id>"
```