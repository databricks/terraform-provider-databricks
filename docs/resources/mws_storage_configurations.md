---
subcategory: "Deployment"
---
# databricks_mws_storage_configurations Resource

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.cloud.databricks.com"` and use `provider = databricks.mws`

This resource to configure root bucket new workspaces within AWS.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to <https://accounts.cloud.databricks.com> for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces api.

Please follow this [complete runnable example](../guides/aws-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.cloud.databricks.com/"`) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module for creation of workspace + PAT token and the rest in different modules.

## Example Usage

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

resource "aws_s3_bucket" "root_storage_bucket" {
  bucket = "${var.prefix}-rootbucket"
  acl    = "private"
}

resource "aws_s3_bucket_versioning" "root_versioning" {
  bucket = aws_s3_bucket.root_storage_bucket.id
  versioning_configuration {
    status = "Disabled"
  }
}

resource "databricks_mws_storage_configurations" "this" {
  provider                   = databricks.mws
  account_id                 = var.databricks_account_id
  storage_configuration_name = "${var.prefix}-storage"
  bucket_name                = aws_s3_bucket.root_storage_bucket.bucket
}
```

## Argument Reference

The following arguments are required:

* `bucket_name` - name of AWS S3 bucket
* `account_id` - Account Id that could be found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `storage_configuration_name` - name under which this storage configuration is stored

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws storage configurations.
* `storage_configuration_id` - (String) id of storage config to be used for `databricks_mws_workspace` resource.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [Provisioning Databricks on AWS with Private Link](../guides/aws-private-link-workspace.md) guide.
* [databricks_mws_credentials](mws_credentials.md) to configure the cross-account role for creation of new workspaces within AWS.
* [databricks_mws_customer_managed_keys](mws_customer_managed_keys.md) to configure KMS keys for new workspaces within AWS.
* [databricks_mws_log_delivery](mws_log_delivery.md) to configure delivery of [billable usage logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html) and [audit logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_workspaces](mws_workspaces.md) to set up [AWS and GCP workspaces](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
