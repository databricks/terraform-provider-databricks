---
subcategory: "Deployment"
---
# databricks_aws_unity_catalog_assume_role_policy Data Source

-> **Note** This resource has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/data-governance/unity-catalog/get-started.html#configure-a-storage-bucket-and-iam-role-in-aws) in case of any questions.

This data source constructs necessary AWS Unity Catalog assume role policy for you.

## Example Usage

```hcl
data "databricks_aws_unity_catalog_policy" "this" {
  aws_account_id = var.aws_account_id
  bucket_name    = "databricks-bucket"
  role_name      = "${var.prefix}-uc-access"
  kms_name       = "databricks-kms"
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
  assume_role_policy  = data.aws_iam_policy_document.this.json
  managed_policy_arns = [aws_iam_policy.unity_metastore.arn]
}
```

## Argument Reference

* `aws_account_id` (Required) The Account ID of the current AWS account (not your Databricks account).
* `external_id` (Required) The [storage credential](../resources/storage_credential.md) external id.
* `role_name` (Required) The name of the AWS IAM role that you created in the previous step in the [official documentation](https://docs.databricks.com/data-governance/unity-catalog/get-started.html#configure-a-storage-bucket-and-iam-role-in-aws).
* `unity_catalog_iam_arn` (Optional) The Databricks Unity Catalog IAM Role ARN. Defaults to `arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL`

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document for assume role
