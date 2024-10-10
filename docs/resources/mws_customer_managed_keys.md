---
subcategory: "Deployment"
---
# databricks_mws_customer_managed_keys Resource

-> Initialize provider with `alias = "mws"`, `host  = "https://accounts.cloud.databricks.com"` and use `provider = databricks.mws`

This resource to configure KMS keys for new workspaces within AWS or GCP. This is to support the following features:

* [Customer-managed keys for managed services](https://docs.databricks.com/security/keys/customer-managed-keys-managed-services-aws.html): Encrypt the workspace’s managed services data in the control plane, including notebooks, secrets, Databricks SQL queries, and Databricks SQL query history  with a CMK.
* [Customer-managed keys for workspace storage](https://docs.databricks.com/security/keys/customer-managed-keys-storage-aws.html): Encrypt the workspace's root S3 bucket and clusters' EBS volumes with a CMK.

Please follow this [complete runnable example](../guides/aws-workspace.md) with new VPC and new workspace setup. Please pay special attention to the fact that there you have two different instances of a databricks provider - one for deploying workspaces (with `host="https://accounts.cloud.databricks.com/"`) and another for the workspace you've created with databricks_mws_workspaces resource. If you want both creation of workspaces & clusters within workspace within the same terraform module (essentially same directory), you should use the provider aliasing feature of Terraform. We strongly recommend having one Terraform module for creation of workspace + PAT token and the rest in different modules.

## Example Usage

-> If you've used the resource before, please add `use_cases = ["MANAGED_SERVICES"]` to keep the previous behaviour.

### Customer-managed key for managed services

You must configure this during workspace creation

#### For AWS

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "databricks_managed_services_cmk" {
  version = "2012-10-17"
  statement {
    sid    = "Enable IAM User Permissions"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
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
  account_id = var.databricks_account_id
  aws_key_info {
    key_arn   = aws_kms_key.managed_services_customer_managed_key.arn
    key_alias = aws_kms_alias.managed_services_customer_managed_key_alias.name
  }
  use_cases = ["MANAGED_SERVICES"]
}
# supply databricks_mws_customer_managed_keys.managed_services.customer_managed_key_id as managed_services_customer_managed_key_id for databricks_mws_workspaces
```

#### For GCP

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.gcp.databricks.com/"
}

variable "cmek_resource_id" {
  description = "Id of a google_kms_crypto_key"
}

resource "databricks_mws_customer_managed_keys" "managed_services" {
  account_id = var.databricks_account_id
  gcp_key_info {
    kms_key_id = var.cmek_resource_id
  }
  use_cases = ["MANAGED_SERVICES"]
}
# supply databricks_mws_customer_managed_keys.managed_services.customer_managed_key_id as managed_services_customer_managed_key_id for databricks_mws_workspaces
```

### Customer-managed key for workspace storage

#### For AWS

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.cloud.databricks.com/"
}

variable "databricks_cross_account_role" {
  description = "AWS ARN for the Databricks cross account role"
}

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "databricks_storage_cmk" {
  version = "2012-10-17"
  statement {
    sid    = "Enable IAM User Permissions"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
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
  account_id = var.databricks_account_id
  aws_key_info {
    key_arn   = aws_kms_key.storage_customer_managed_key.arn
    key_alias = aws_kms_alias.storage_customer_managed_key_alias.name
  }
  use_cases = ["STORAGE"]
}
# supply databricks_mws_customer_managed_keys.storage.customer_managed_key_id as storage_customer_managed_key_id for databricks_mws_workspaces
```

#### For GCP

```hcl
variable "databricks_account_id" {
  description = "Account Id that could be found in the top right corner of https://accounts.gcp.databricks.com/"
}

variable "cmek_resource_id" {
  description = "Id of a google_kms_crypto_key"
}

resource "databricks_mws_customer_managed_keys" "storage" {
  account_id = var.databricks_account_id
  gcp_key_info {
    kms_key_id = var.cmek_resource_id
  }
  use_cases = ["STORAGE"]
}
# supply databricks_mws_customer_managed_keys.storage.customer_managed_key_id as storage_customer_managed_key_id for databricks_mws_workspaces
```

## Argument Reference

The following arguments are required:

* `aws_key_info` - This field is a block and is documented below. This conflicts with `gcp_key_info`
* `gcp_key_info` - This field is a block and is documented below. This conflicts with `aws_key_info`
* `account_id` - Account Id that could be found in the top right corner of [Accounts Console](https://accounts.cloud.databricks.com/)
* `use_cases` - *(since v0.3.4)* List of use cases for which this key will be used. *If you've used the resource before, please add `use_cases = ["MANAGED_SERVICES"]` to keep the previous behaviour.* Possible values are:
  * `MANAGED_SERVICES` - for encryption of the workspace objects (notebooks, secrets) that are stored in the control plane
  * `STORAGE` - for encryption of the DBFS Storage & Cluster EBS Volumes

### aws_key_info Configuration Block

* `key_arn` - The AWS KMS key's Amazon Resource Name (ARN).
* `key_alias` - (Optional) The AWS KMS key alias.
* `key_region` - (Optional) (Computed) The AWS region in which KMS key is deployed to. This is not required.

### gcp_key_info Configuration Block

* `kms_key_id` - The GCP KMS key's resource name.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the mws customer managed keys.
* `customer_managed_key_id` - (String) ID of the encryption key configuration object.
* `creation_time` - (Integer) Time in epoch milliseconds when the customer key was created.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are used in the same context:

* [Provisioning Databricks on AWS](../guides/aws-workspace.md) guide.
* [databricks_mws_credentials](mws_credentials.md) to configure the cross-account role for creation of new workspaces within AWS.
* [databricks_mws_log_delivery](mws_log_delivery.md) to configure delivery of [billable usage logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html) and [audit logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
* [databricks_mws_networks](mws_networks.md) to [configure VPC](https://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html) & subnets for new workspaces within AWS.
* [databricks_mws_storage_configurations](mws_storage_configurations.md) to configure root bucket new workspaces within AWS.
* [databricks_mws_workspaces](mws_workspaces.md) to set up [AWS and GCP workspaces](https://docs.databricks.com/getting-started/overview.html#e2-architecture-1).
