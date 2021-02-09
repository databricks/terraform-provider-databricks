---
subcategory: "E2"
---
# databricks_mws_customer_managed_keys Resource

-> **Note** This resource has an evolving API, which may change in future versions of the provider.

This resource to configure AWS KMS key for new workspaces within AWS. This KMS key will be used to encrypt your notebooks in your workspace.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces api. 

Please follow this [complete runnable example](https://github.com/databrickslabs/terraform-provider-databricks/blob/master/scripts/awsmt-integration/main.tf) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with host=https://accounts.cloud.databricks.com/) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one terraform module for creation of workspace + PAT token and the rest in different modules.


## Example Usage

```hcl
resource "aws_kms_key" "customer_managed_key" {
}

resource "aws_kms_grant" "databricks-grant" {
  name = "databricks-grant"
  key_id  = aws_kms_key.customer_managed_key.key_id
  grantee_principal = "arn:aws:iam::414351767826:root"
  operations = ["Encrypt", "Decrypt"]
}

resource "aws_kms_alias" "customer_managed_key_alias" {
  name          = "alias/customer-managed-key-alias"
  target_key_id = aws_kms_key.customer_managed_key.key_id
}

resource "databricks_mws_customer_managed_keys" "my_cmk" {
    account_id   = var.account_id
    aws_key_info {
        key_arn   = aws_kms_key.customer_managed_key.arn
        key_alias = aws_kms_alias.customer_managed_key_alias.name
    }
}
```


## Argument Reference

The following arguments are required:

* `aws_key_info` - (Required) (List) This field is a block and is documented below.
* `account_id` - (Required) (String) The Databricks account ID that holds the customer-managed key.


### aws_key_info Configuration Block

* `key_arn` - (Required) (String) The AWS KMS key's Amazon Resource Name (ARN).
* `key_alias` - (Required) (String) The AWS KMS key alias.
* `key_region` - (Optional) (Computed) (String) The AWS region in which KMS key is deployed to. This is not required.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws customer managed keys.
* `customer_managed_key_id` - (String) ID of the notebook encryption key configuration object.
* `creation_time` - (Integer) Time in epoch milliseconds when the customer key was created.
