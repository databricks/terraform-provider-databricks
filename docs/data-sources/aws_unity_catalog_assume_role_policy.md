---
subcategory: "Deployment"
---
# databricks_aws_unity_catalog_assume_role_policy Data Source

-> **Note** This resource has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/data-governance/unity-catalog/get-started.html#configure-a-storage-bucket-and-iam-role-in-aws) in case of any questions.

This data source constructs the necessary AWS Unity Catalog assume role policy for you.

## Example Usage

```hcl
data "databricks_aws_unity_catalog_policy" "this" {
  aws_account_id = var.aws_account_id
  bucket_name    = "databricks-bucket"
  role_name      = "${var.prefix}-uc-access"
  kms_name       = "arn:aws:kms:us-west-2:111122223333:key/databricks-kms"
}

data "databricks_aws_unity_catalog_assume_role_policy" "this" {
  aws_account_id = var.aws_account_id
  role_name      = "${var.prefix}-uc-access"
  external_id    = "12345"
}

resource "aws_iam_policy" "unity_metastore" {
  name   = "${var.prefix}-unity-catalog-metastore-access-iam-policy"
  policy = data.databricks_aws_unity_catalog_policy.this.json
}

resource "aws_iam_role" "metastore_data_access" {
  name                = "${var.prefix}-uc-access"
  assume_role_policy  = data.databricks_aws_unity_catalog_assume_role_policy.this.json
  managed_policy_arns = [aws_iam_policy.unity_metastore.arn]
}
```

## Argument Reference

* `aws_account_id` (Required) The Account ID of the current AWS account (not your Databricks account).
* `aws_partition` - (Optional) AWS partition. The options are `aws`,`aws-us-gov` or `aws-us-gov-dod`. Defaults to `aws`
* `external_id` (Required) The [storage credential](../resources/storage_credential.md) external id.
* `role_name` (Required) The name of the AWS IAM role to be created for Unity Catalog.
* `unity_catalog_iam_arn` (Optional) The Databricks Unity Catalog IAM Role ARN. Defaults to `arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL` on standard AWS partition selection, `arn:aws-us-gov:iam::044793339203:role/unity-catalog-prod-UCMasterRole-1QRFA8SGY15OJ` on GovCloud partition selection, and `arn:aws-us-gov:iam::170661010020:role/unity-catalog-prod-UCMasterRole-1DI6DL6ZP26AS` on GovCloud DoD partition selection

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document for assume role
