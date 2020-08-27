# databricks_aws_assume_role_policy Data Source

This data source constructs necessary AWS STS assume role policy for you.

## Example Usage

End-to-end example of provisioning Cross-account IAM role:

```hcl
variable "account_id" {
  type        = string
  description = "External ID you find on https://accounts.cloud.databricks.com/#aws"
}

data "databricks_aws_crossaccount_policy" "this" {}

resource "aws_iam_policy" "cross_account_policy" {
  name   = "${var.prefix}-crossaccount-iam-policy"
  policy = data.databricks_aws_crossaccount_policy.this.json
}

data "databricks_aws_assume_role_policy" "this" {
    external_id = var.account_id
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

// required only in case of multiworkspace setup
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.account_id
  credentials_name = "${var.prefix}-creds"
  role_arn         = aws_iam_role.cross_account.arn
}
```

## Argument Reference

* `external_id` (Required) (String) External ID that can be found at http://accounts.cloud.databricks.com/#aws

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `json` - AWS IAM Policy JSON document