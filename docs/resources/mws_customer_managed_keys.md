---
subcategory: "AWS"
---
# databricks_mws_customer_managed_keys Resource

-> **Note** This resource has an evolving API, which will change in the upcoming versions of the provider in order to simplify user experience.

This resource to configure KMS keys for new workspaces within AWS. This is to support the following features:
* [Customer-managed keys for managed services](https://docs.databricks.com/security/keys/customer-managed-keys-managed-services-aws.html): Encrypt the workspace’s managed services data in the control plane, including notebooks, secrets, Databricks SQL queries, and Databricks SQL query history  with a CMK.
* [Customer-managed keys for workspace storage](https://docs.databricks.com/security/keys/customer-managed-keys-storage-aws.html): Encrypt the workspace's root S3 bucket and clusters' EBS volumes with a CMK.

It is important to understand that this will require you to configure your provider separately for the multiple workspaces resources. This will point to https://accounts.cloud.databricks.com for the HOST and it will use basic auth as that is the only authentication method available for multiple workspaces API.

Please follow this [complete runnable example](../guides/aws-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.cloud.databricks.com/"`) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one Terraform module for creation of workspace + PAT token and the rest in different modules.


## Example Usage

-> **Note** If you've used the resource before, please add `use_cases = ["MANAGED_SERVICES"]` to keep the previous behaviour.

### Customer-managed key for managed services

You must configure this during workspace creation

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}

data "aws_iam_policy_document" "databricks_managed_services_cmk" {
  version = "2012-10-17"
  statement {
    sid    = "Enable IAM User Permissions"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
    actions   = ["kms:*"]
    resources = ["*"]
  }
  statement {
    sid    = "Allow Databricks to use KMS key for control plane managed services"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::414351767826:root"]
    }
    actions = [
      "kms:Encrypt",
      "kms:Decrypt"
    ]
    resources = ["*"]
  }
}

resource "aws_kms_key" "managed_services_customer_managed_key" {
  policy = data.aws_iam_policy_document.databricks_managed_services_cmk.json
}

resource "aws_kms_alias" "managed_services_customer_managed_key_alias" {
  name          = "alias/managed-services-customer-managed-key-alias"
  target_key_id = aws_kms_key.managed_services_customer_managed_key.key_id
}

resource "databricks_mws_customer_managed_keys" "managed_services" {
    provider     = databricks.mws
    account_id   = var.databricks_account_id
    aws_key_info {
        key_arn   = aws_kms_key.managed_services_customer_managed_key.arn
        key_alias = aws_kms_alias.managed_services_customer_managed_key_alias.name
    }
    use_cases = ["MANAGED_SERVICES"]
}
# supply databricks_mws_customer_managed_keys.managed_services.customer_managed_key_id as managed_services_customer_managed_key_id for databricks_mws_workspaces
```

### Customer-managed key for workspace storage

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the bottom left corner of https://accounts.cloud.databricks.com/"
}

variable "databricks_cross_account_role" {
  description = "AWS ARN for the Databricks cross account role"
}

data "aws_iam_policy_document" "databricks_storage_cmk" {
  version = "2012-10-17"
  statement {
    sid    = "Enable IAM User Permissions"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
    actions   = ["kms:*"]
    resources = ["*"]
  }
  statement {
    sid    = "Allow Databricks to use KMS key for DBFS"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::414351767826:root"]
    }
    actions = [
      "kms:Encrypt",
      "kms:Decrypt",
      "kms:ReEncrypt*",
      "kms:GenerateDataKey*",
      "kms:DescribeKey"
    ]
    resources = ["*"]
  }
  statement {
    sid    = "Allow Databricks to use KMS key for DBFS (Grants)"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::414351767826:root"]
    }
    actions = [
      "kms:CreateGrant",
      "kms:ListGrants",
      "kms:RevokeGrant"
    ]
    resources = ["*"]
    condition {
      test     = "Bool"
      variable = "kms:GrantIsForAWSResource"
      values   = ["true"]
    }
  }
  statement {
    sid    = "Allow Databricks to use KMS key for EBS"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = [var.databricks_cross_account_role]
    }
    actions = [
      "kms:Decrypt",
      "kms:GenerateDataKey*",
      "kms:CreateGrant",
      "kms:DescribeKey"
    ]
    resources = ["*"]
    condition {
      test     = "ForAnyValue:StringLike"
      variable = "kms:ViaService"
      values   = ["ec2.*.amazonaws.com"]
    }
  }
}

resource "aws_kms_key" "storage_customer_managed_key" {
  policy = data.aws_iam_policy_document.databricks_storage_cmk.json
}

resource "aws_kms_alias" "storage_customer_managed_key_alias" {
  name          = "alias/storage-customer-managed-key-alias"
  target_key_id = aws_kms_key.storage_customer_managed_key.key_id
}

resource "databricks_mws_customer_managed_keys" "storage" {
    provider     = databricks.mws
    account_id   = var.databricks_account_id
    aws_key_info {
        key_arn   = aws_kms_key.storage_customer_managed_key.arn
        key_alias = aws_kms_alias.storage_customer_managed_key_alias.name
    }
    use_cases = ["STORAGE"]
}
# supply databricks_mws_customer_managed_keys.storage.customer_managed_key_id as storage_customer_managed_key_id for databricks_mws_workspaces
```

## Argument Reference

The following arguments are required:

* `aws_key_info` - This field is a block and is documented below.
* `account_id` - Account Id that could be found in the bottom left corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `use_cases` - *(since v0.3.4)* List of use cases for which this key will be used. *If you've used the resource before, please add `use_cases = ["MANAGED_SERVICES"]` to keep the previous behaviour.* Possible values are:
  * `MANAGED_SERVICES` - for encryption of the workspace objects (notebooks, secrets) that are stored in the control plane
  * `STORAGE` - for encryption of the DBFS Storage & Cluster EBS Volumes


### aws_key_info Configuration Block

* `key_arn` - The AWS KMS key's Amazon Resource Name (ARN).
* `key_alias` - The AWS KMS key alias.
* `key_region` - (Optional) (Computed) The AWS region in which KMS key is deployed to. This is not required.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws customer managed keys.
* `customer_managed_key_id` - (String) ID of the encryption key configuration object.
* `creation_time` - (Integer) Time in epoch milliseconds when the customer key was created.
