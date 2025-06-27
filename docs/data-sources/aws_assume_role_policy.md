---
subcategory: "Deployment"
---

# databricks_aws_assume_role_policy Data Source

This data source constructs necessary AWS STS assume role policy for you.

-> This data source can be used with an account or workspace-level provider.

## Example Usage

End-to-end example of provisioning Cross-account IAM role with [databricks_mws_credentials](../resources/mws_credentials.md) and [aws_iam_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role):

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

data "databricks_aws_crossaccount_policy" "this" {}

resource "aws_iam_policy" "cross_account_policy" {
  name   = "${var.prefix}-crossaccount-iam-policy"
  policy = data.databricks_aws_crossaccount_policy.this.json
}

data "databricks_aws_assume_role_policy" "this" {
  external_id = var.databricks_account_id
}

resource "aws_iam_role" "cross_account" {
  name               = "${var.prefix}-crossaccount-iam-role"
  assume_role_policy = data.databricks_aws_assume_role_policy.this.json
  description        = "Grants Databricks full access to VPC resources"
}

resource "aws_iam_role_policy_attachment" "cross_account" {
  policy_arn = aws_iam_policy.cross_account_policy.arn
  role       = aws_iam_role.cross_account.name
}

// required only in case of multi-workspace setup
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.databricks_account_id
  credentials_name = "${var.prefix}-creds"
  role_arn         = aws_iam_role.cross_account.arn
}
```

## Argument Reference

* `external_id` (Required) Account Id that could be found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/).
* `aws_partition` - (Optional) AWS partition. The options are `aws`, `aws-us-gov`, or `aws-us-gov-dod`. Defaults to `aws`
* `for_log_delivery` (Optional) Either or not this assume role policy should be created for usage log delivery. Defaults to false.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document

## Related Resources

The following resources are used in the same context:

* [Provisioning AWS Databricks workspaces with a Hub & Spoke firewall for data exfiltration protection](../guides/aws-e2-firewall-hub-and-spoke.md) guide
* [databricks_aws_bucket_policy](aws_bucket_policy.md) data to configure a simple access policy for AWS S3 buckets, so that Databricks can access data in it.
* [databricks_aws_crossaccount_policy](aws_crossaccount_policy.md) data to construct the necessary AWS cross-account policy for you, which is based on [official documentation](https://docs.databricks.com/administration-guide/account-api/iam-role.html#language-Your%C2%A0VPC,%C2%A0default).
