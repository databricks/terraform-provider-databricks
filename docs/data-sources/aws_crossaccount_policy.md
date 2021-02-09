---
subcategory: "AWS"
---
# databricks_aws_crossaccount_policy Data Source

-> **Note** This resource has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default) in case of any questions.

This data source constructs necessary AWS cross-account policy for you, which is based on [official documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default).

## Example Usage

For more detailed usage please see [databricks_aws_assume_role_policy](aws_assume_role_policy.md) or [databricks_aws_s3_mount](../resources/aws_s3_mount.md) pages.

```hcl
data "databricks_aws_crossaccount_policy" "this" {}
```

## Argument Reference

* `pass_roles` (Optional) (List) List of Data IAM role ARNs that are explicitly granted `iam:PassRole` action.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document
