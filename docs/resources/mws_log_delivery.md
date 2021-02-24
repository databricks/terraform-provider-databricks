---
subcategory: "Log Delivery"
---
# databricks_mws_log_delivery Resource

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

Make sure you have authenticated with [username and password for Accounts Console](../guides/aws-workspace.md). This resource configures the delivery of the two supported log types from Databricks workspaces: [billable usage logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html) and [audit logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html). You cannot delete a log delivery configuration, but you can disable it when you no longer need it. This fact is important because there is a limit to the number of enabled log delivery configurations that you can create for an account. You can create a maximum of two enabled using the account level *(without workspace filter)* and two that use the workspace filter. There is an additional uniqueness constraint that two enabled configurations cannot share all their fields (not including the `config_name`). Re-enabling may fail when there's a violation of limit or uniqueness constraints.

## Example Usage

End-to-end example of usage and audit log delivery:

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}

resource "aws_s3_bucket" "logdelivery" {
  bucket = "${var.prefix}-logdelivery"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(var.tags, {
    Name = "${var.prefix}-logdelivery"
  })
}

resource "aws_s3_bucket_public_access_block" "logdelivery" {
  bucket             = aws_s3_bucket.logdelivery.id
  ignore_public_acls = true
}

data "databricks_aws_assume_role_policy" "logdelivery" {
  external_id = var.databricks_account_id
  for_log_delivery = true
}

resource "aws_iam_role" "logdelivery" {
  name               = "${var.prefix}-logdelivery"
  description        = "(${var.prefix}) UsageDelivery role"
  assume_role_policy = data.databricks_aws_assume_role_policy.logdelivery.json
  tags               = var.tags
}

data "databricks_aws_bucket_policy" "logdelivery" {
  full_access_role = aws_iam_role.logdelivery.arn
  bucket           = aws_s3_bucket.logdelivery.bucket
}

resource "aws_s3_bucket_policy" "logdelivery" {
  bucket = aws_s3_bucket.logdelivery.id
  policy = data.databricks_aws_bucket_policy.logdelivery.json
}

resource "databricks_mws_credentials" "log_writer" {
    account_id       = var.databricks_account_id
    credentials_name = "Usage Delivery"
    role_arn         = aws_iam_role.logdelivery.arn
}

resource "databricks_mws_storage_configurations" "log_bucket" {
    account_id                 = var.databricks_account_id
    storage_configuration_name = "Usage Logs"
    bucket_name                = aws_s3_bucket.logdelivery.bucket
}

resource "databricks_mws_log_delivery" "usage_logs" {
    account_id = var.databricks_account_id
    credentials_id = databricks_mws_credentials.log_writer.credentials_id
    storage_configuration_id = databricks_mws_storage_configurations.log_bucket.storage_configuration_id
    delivery_path_prefix = "billable-usage"
    config_name = "Usage Logs"
    log_type = "BILLABLE_USAGE"
    output_format = "CSV"
}

resource "databricks_mws_log_delivery" "audit_logs" {
    account_id = var.databricks_account_id
    credentials_id = databricks_mws_credentials.log_writer.credentials_id
    storage_configuration_id = databricks_mws_storage_configurations.log_bucket.storage_configuration_id
    delivery_path_prefix = "audit-logs"
    config_name = "Audit Logs"
    log_type = "AUDIT_LOGS"
    output_format = "JSON"
}
```

## Billable Usage

CSV files are delivered to `<delivery_path_prefix>/billable-usage/csv/` and are named `workspaceId=<workspace-id>-usageMonth=<month>.csv`, which are delivered daily by overwriting the month's CSV file for each workspace. Format of CSV file, as well as some usage examples, can be found [here](https://docs.databricks.com/administration-guide/account-settings/usage.html#download-usage-as-a-csv-file). 

Common processing scenario is to apply [cost allocation tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html), that could be enforced by setting [custom_tags](cluster.md#custom_tags) on a cluster or through [cluster policy](cluster_policy.md). Report contains `clusterId` field, that could be joined with data from AWS [cost and usage reports](https://docs.aws.amazon.com/cur/latest/userguide/cur-create.html), that can be joined with `user:ClusterId` tag from AWS usage report.

```hcl
resource "databricks_mws_log_delivery" "usage_logs" {
    account_id = var.databricks_account_id
    credentials_id = databricks_mws_credentials.log_writer.credentials_id
    storage_configuration_id = databricks_mws_storage_configurations.log_bucket.storage_configuration_id
    delivery_path_prefix = "billable-usage"
    config_name = "Usage Logs"
    log_type = "BILLABLE_USAGE"
    output_format = "CSV"
}
```

## Audit Logs

JSON files with [static schema](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html#audit-log-schema) are delivered to `<delivery_path_prefix>/workspaceId=<workspaceId>/date=<yyyy-mm-dd>/auditlogs_<internal-id>.json`. Logs are available within 15 minutes of activation for audit logs. New JSON files are delivered every few minutes, potentially overwriting existing files for each workspace. Sometimes data may arrive later than 15 minutes. Databricks can overwrite the delivered log files in your bucket at any time. If a file is overwritten, the existing content remains, but there may be additional lines for more auditable events. Overwriting ensures exactly-once semantics without requiring read or delete access to your account.

```hcl
resource "databricks_mws_log_delivery" "audit_logs" {
    account_id = var.databricks_account_id
    credentials_id = databricks_mws_credentials.log_writer.credentials_id
    storage_configuration_id = databricks_mws_storage_configurations.log_bucket.storage_configuration_id
    delivery_path_prefix = "audit-logs"
    config_name = "Audit Logs"
    log_type = "AUDIT_LOGS"
    output_format = "JSON"
}
```

## Argument reference

* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/).
* `config_name` - The optional human-readable name of the log delivery configuration. Defaults to empty.
* `log_type` - The type of log delivery. `BILLABLE_USAGE` and `AUDIT_LOGS` are supported.
* `output_format` - The file type of log delivery. Currently `CSV` (for `BILLABLE_USAGE`) and `JSON` (for `AUDIT_LOGS`) are supported.
* `credentials_id` - The ID for a Databricks [credential configuration](mws_credentials.md) that represents the AWS IAM role [with policy](../data-sources/aws_assume_role_policy.md) and [trust relationship](../data-sources/aws_assume_role_policy.md) as described in the main billable usage documentation page.
* `storage_configuration_id` - The ID for a Databricks [storage configuration](mws_storage_configurations.md) that represents the S3 bucket with [bucket policy](../data-sources/aws_bucket_policy.md) as described in the main billable usage documentation page.
* `workspace_ids_filter` - (Optional) By default, this log configuration applies to all workspaces associated with your account ID. If your account is on the E2 version of the platform or on a select custom plan that allows multiple workspaces per account, you may have multiple workspaces associated with your account ID. You can optionally set the field as mentioned earlier to an array of workspace IDs. If you plan to use different log delivery configurations for several workspaces, set this explicitly rather than leaving it blank. If you leave this blank and your account ID gets additional workspaces in the future, this configuration will also apply to the new workspaces.
* `delivery_path_prefix` - (Optional) Defaults to empty, which means that logs are delivered to the root of the bucket. The value must be a valid S3 object key. It must not start or end with a slash character.
* `delivery_start_time` - (Optional) The optional start month and year for delivery, specified in YYYY-MM format. Defaults to current year and month. Usage is not available before 2019-03.

## Attribute reference

Resource exports the following attributes:

* `config_id` - Databricks log delivery configuration ID.

## Import

This resource can be imported by specifying a combination of an account id and log config id separated by `|`:

```bash
$ terraform import databricks_mws_log_delivery.usage "<account-id>|<log-config-id>"
```