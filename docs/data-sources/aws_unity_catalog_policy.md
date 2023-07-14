---
subcategory: "Deployment"
---
# databricks_aws_unity_catalog_policy Data Source

-> **Note** This resource has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default) in case of any questions.

This data source constructs necessary AWS Unity Catalog policy for you, which is based on [official documentation](https://docs.databricks.com/data-governance/unity-catalog/get-started.html#configure-a-storage-bucket-and-iam-role-in-aws).

## Example Usage

```hcl
data "databricks_aws_unity_catalog_policy" "this" {
  aws_account_id = "123456789098"
  bucket_name = "databricks-bucket"
  role_name = "databricks-role"
  kms_name = "databricks-kms"
}

resource "aws_iam_policy" "unity_catalog_policy" {
  name   = "${var.prefix}-unity-catalog-iam-policy"
  policy = data.databricks_aws_unity_catalog_policy.this.json
}
```

## Argument Reference

* `aws_account_id` (Required) The Account ID of the current AWS account (not your Databricks account).
* `bucket_name` (Required) The name of the S3 bucket used as root storage location for [managed tables](https://docs.databricks.com/data-governance/unity-catalog/index.html#managed-table) in Unity Catalog.
* `role_name` (Required) The name of the AWS IAM role that you created in the previous step in the [official documentation](https://docs.databricks.com/data-governance/unity-catalog/get-started.html#configure-a-storage-bucket-and-iam-role-in-aws).
* `kms_name` (Optional) If encryption is enabled, provide the name of the KMS key that encrypts the S3 bucket contents. If encryption is disabled, do not provide this argument.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document
