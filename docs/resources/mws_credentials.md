# databricks_mws_credentials Resource

This resource to configure the cross-account role for creation of new workspaces within AWS.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces api. 

Please follow this [complete runnable example](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with host=https://accounts.cloud.databricks.com/) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module for creation of workspace + PAT token and the rest in different modules.

## Example Usage

**This resource has evolving API, which may change in future versions of provider.**

```hcl
resource "aws_iam_role" "cross_account_role" {
  name = "${var.prefix}-crossaccount"
  assume_role_policy = data.template_file.cross_account_role_assume_policy.rendered
  tags               = var.tags
}

resource "aws_iam_policy" "cross_account_role_policy" {
  name = "${var.prefix}-policy"
  description = "Workspace Cross account role policy policy"
  policy = data.template_file.cross_account_role_policy.rendered
}

resource "aws_iam_role_policy_attachment" "cross_account_role_policy_attach" {
  role       = aws_iam_role.cross_account_role.name
  policy_arn = aws_iam_policy.cross_account_role_policy.arn
}

provider "databricks" {
  alias = "mws"
  host  = "https://accounts.cloud.databricks.com"
}

// register cross-account ARN
resource "databricks_mws_credentials" "this" {
  provider         = databricks.mws
  account_id       = var.account_id
  credentials_name = "${var.prefix}-creds"
  role_arn         = aws_iam_role.cross_account_role.arn
}
```

## Argument Reference

The following arguments are required:

* `account_id` - (Required) (String) master account id (also used for `sts:ExternalId` of `sts:AssumeRole`)
* `credentials_name` - (Required) (String) name of credentials to register
* `role_arn` - (Required) (String) ARN of cross-account role


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws credentials.
* `creation_time` - (Integer) time of credentials registration
* `external_id` - (String) master account id
* `credentials_id` - (String) identifier of credentials