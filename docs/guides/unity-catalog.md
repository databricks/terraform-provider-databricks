---
page_title: "Unity Catalog set up on AWS"
---

# Deploying pre-requisite resources and enabling Unity Catalog (AWS Preview)

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/data-governance/unity-catalog/index.html). Contact your Databricks representative to request access. 

Databricks Unity Catalog brings fine-grained governance and security to Lakehouse data using a familiar, open interface. You can use Terraform to deploy the underlying cloud resources and Unity Catalog objects automatically, using a programmatic approach.

This guide uses the following variables in configurations:

- `databricks_account_username`: The username an account-level admin uses to log in to  [https://accounts.cloud.databricks.com](https://accounts.cloud.databricks.com).
- `databricks_account_password`: The password for `databricks_account_username`.
- `databricks_account_id`: The numeric ID for your Databricks account. When you are logged in, it appears in the bottom left corner of the page.
- `databricks_workspace_url`: Value of `workspace_url` attribute from [databricks_mws_workspaces](../resources/mws_workspaces.md#attribute-reference) resource.

This guide is provided as-is and you can use this guide as the basis for your custom Terraform module.

To get started with Unity Catalog, this guide takes you throw the following high-level steps:
- [Deploying pre-requisite resources and enabling Unity Catalog (AWS Preview)](#deploying-pre-requisite-resources-and-enabling-unity-catalog-aws-preview)
  - [Provider initialization](#provider-initialization)
  - [Configure AWS objects](#configure-aws-objects)
  - [Create users and groups](#create-users-and-groups)
  - [Create a Unity Catalog metastore and link it to workspaces](#create-a-unity-catalog-metastore-and-link-it-to-workspaces)
  - [Create Unity Catalog objects in the metastore](#create-unity-catalog-objects-in-the-metastore)
  - [Configure external tables and credentials](#configure-external-tables-and-credentials)
  - [Configure Unity Catalog clusters](#configure-unity-catalog-clusters)

## Provider initialization

Initialize [provider with `mws` alias](https://www.terraform.io/language/providers/configuration#alias-multiple-provider-configurations) to set up account-level resources. See [provider authentication](../index.md#authenticating-with-hostname,-username,-and-password) for more details.

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databrickslabs/databricks"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "3.49.0"
    }
  }
}

provider "aws" {
  region = var.region
}

// initialize provider in "MWS" mode for account-level resources
provider "databricks" {
  alias      = "mws"
  host       = "https://accounts.cloud.databricks.com"
  account_id = var.databricks_account_id
  username   = var.databricks_account_username
  password   = var.databricks_account_password
}

// initialize provider at workspace level, to create UC resources
provider "databricks" {
  alias    = "workspace"
  host     = var.databricks_workspace_url
  username = var.databricks_account_username
  password = var.databricks_account_password
}
```

Define the required variables

```hcl
variable "databricks_account_username" {}
variable "databricks_account_password" {}
variable "databricks_account_id" {}
variable "databricks_workspace_url" {}

variable "tags" {
  default = {}
}

variable "region" {
  default = "eu-west-1"
}

variable "databricks_workspace_ids" {
  description = <<EOT
  List of Databricks workspace IDs to be enabled with Unity Catalog.
  Enter with square brackets and double quotes
  e.g. ["111111111", "222222222"]
  EOT
  type        = list(string)
}

variable "databricks_users" {
  description = <<EOT
  List of Databricks users to be added at account-level for Unity Catalog.
  Enter with square brackets and double quotes
  e.g ["first.last@domain.com", "second.last@domain.com"]
  EOT
  type        = list(string)
}

variable "databricks_metastore_admins" {
  description = <<EOT
  List of Admins to be added at account-level for Unity Catalog.
  Enter with square brackets and double quotes
  e.g ["first.admin@domain.com", "second.admin@domain.com"]
  EOT
  type        = list(string)
}

variable "unity_admin_group" {
  description = "Name of the admin group. This group will be set as the owner of the Unity Catalog metastore"
  type        = string
}

//generate a random string as the prefix for AWS resources, to ensure uniqueness
resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix = "demo${random_string.naming.result}"
}
```

## Configure AWS objects
The first step is to create the required AWS objects:
- An S3 bucket, which is the default storage location for managed tables in Unity Catalog. Please use a dedicated bucket for each metastore.
- An IAM policy that provides Unity Catalog permissions to access and manage data in the bucket.
- An IAM role that is associated with the IAM policy and will be assumed by Unity Catalog.

```hcl
resource "aws_s3_bucket" "metastore" {
  bucket = "${local.prefix}-metastore"
  acl    = "private"
  versioning {
    enabled = false
  }
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-metastore"
  })
}

resource "aws_s3_bucket_public_access_block" "metastore" {
  bucket             = aws_s3_bucket.metastore.id
  ignore_public_acls = true
  depends_on         = [aws_s3_bucket.metastore]
}

resource "aws_s3_bucket_public_access_block" "metastore" {
  bucket                  = aws_s3_bucket.metastore.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
  depends_on              = [aws_s3_bucket.metastore]
}

data "aws_iam_policy_document" "passrole_for_uc" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["arn:aws:iam::414351767826:role/unity-catalog-prod-UCMasterRole-14S5ZJVKOTYTL"]
      type        = "AWS"
    }
    condition {
      test     = "StringEquals"
      variable = "sts:ExternalId"
      values   = [var.databricks_account_id]
    }
  }
}

resource "aws_iam_policy" "unity_metastore" {
  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "${local.prefix}-databricks-unity-metastore"
    Statement = [
      {
        "Action" : [
          "s3:GetObject",
          "s3:GetObjectVersion",
          "s3:PutObject",
          "s3:PutObjectAcl",
          "s3:DeleteObject",
          "s3:ListBucket",
          "s3:GetBucketLocation"
        ],
        "Resource" : [
          aws_s3_bucket.unity_metastore.arn,
          "${aws_s3_bucket.unity_metastore.arn}/*"
        ],
        "Effect" : "Allow"
      }
    ]
  })
  tags = merge(local.tags, {
    Name = "${local.prefix}-unity-catalog IAM policy"
  })
}

// Required, in case https://docs.databricks.com/data/databricks-datasets.html are needed
resource "aws_iam_policy" "sample_data" {
  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "${local.prefix}-databricks-sample-data"
    Statement = [
      {
        "Action" : [
          "s3:GetObject",
          "s3:GetObjectVersion",
          "s3:ListBucket",
          "s3:GetBucketLocation"
        ],
        "Resource" : [
          "arn:aws:s3:::databricks-datasets-oregon/*",
          "arn:aws:s3:::databricks-datasets-oregon"

        ],
        "Effect" : "Allow"
      }
    ]
  })
  tags = merge(local.tags, {
    Name = "${local.prefix}-unity-catalog IAM policy"
  })
}

resource "aws_iam_role" "metastore_data_access" {
  name                = "${local.prefix}-uc-access"
  assume_role_policy  = data.aws_iam_policy_document.passrole_for_uc.json
  managed_policy_arns = [aws_iam_policy.unity_metastore.arn, aws_iam_policy.sample_data.arn]
  tags = merge(local.tags, {
    Name = "${local.prefix}-unity-catalog IAM role"
  })
}
```

## Create users and groups

A Unity Catalog [databricks_metastore](../resources/metastore.md) can be shared across multiple Databricks workspaces. To enable this, Databricks must have a consistent view of users and groups across all workspaces, and has introduced features within the account console to manage this. Users and groups that wish to use Unity Catalog must be created as account level identities and as workspace-level identities. All users are added to the `account users` group by default.

-> **Note** Databricks does not allow a single user to be added to more than one Databricks account. You will receive the error `User already exists in another account`.

```hcl
resource "databricks_user" "unity_users" {
  provider  = databricks.mws
  for_each  = toset(concat(var.databricks_users, var.databricks_metastore_admins))
  user_name = each.key
  force     = true
}

resource "databricks_group" "admin_group" {
  provider     = databricks.mws
  display_name = var.unity_admin_group
}

resource "databricks_group_member" "admin_group_member" {
  provider  = databricks.mws
  for_each  = toset(var.databricks_metastore_admins)
  group_id  = databricks_group.admin_group.id
  member_id = databricks_user.unity_users[each.value].id
}

resource "databricks_user_role" "metastore_admin" {
  provider = databricks.mws
  for_each = toset(var.databricks_metastore_admins)
  user_id  = databricks_user.unity_users[each.value].id
  role     = "account_admin"
}
```
## Create a Unity Catalog metastore and link it to workspaces

A [databricks_metastore](../resources/metastore.md) is the top level container for data in Unity Catalog. A single metastore can be shared across Databricks workspaces, and each linked workspace has a consistent view of the data and a single set of access policies. Databricks recommends using a small number of metastores, except when organizations wish to have hard isolation boundaries between data. Data cannot be easily joined/queried across metastores.

```hcl
resource "databricks_metastore" "this" {
  provider      = databricks.workspace
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = var.unity_admin_group
  force_destroy = true
}

resource "databricks_metastore_data_access" "this" {
  provider     = databricks.workspace
  metastore_id = databricks_metastore.this.id
  name         = aws_iam_role.metastore_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.metastore_data_access.arn
  }
  is_default = true
}

resource "databricks_metastore_assignment" "default_metastore" {
  provider             = databricks.workspace
  for_each             = toset(var.databricks_workspace_ids)
  workspace_id         = each.key
  metastore_id         = databricks_metastore.unity.id
  default_catalog_name = "hive_metastore"
}
```

## Create Unity Catalog objects in the metastore

Each metastore exposes a 3-level namespace (catalog-schema-table) by which data can be organized. 

```hcl
resource "databricks_catalog" "sandbox" {
  provider     = databricks.workspace
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
  depends_on = [databricks_metastore_assignment.default_metastore]
}

resource "databricks_grants" "sandbox" {
  provider = databricks.workspace
  catalog  = databricks_catalog.sandbox.name
  grant {
    principal  = "Data Scientists"
    privileges = ["USAGE", "CREATE"]
  }
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}

resource "databricks_schema" "things" {
  provider     = databricks.workspace
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grants" "things" {
  provider = databricks.workspace
  schema   = databricks_schema.things.id
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}
```

## Configure external tables and credentials

To work with external tables, Unity Catalog introduces two new objects to access and work with external cloud storage:
- [databricks_storage_credential](../resources/storage_credential.md) represent authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a service principal for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
- [databricks_external_location](../resources/external_location.md) are objects that combine a cloud storage path with a Storage Credential that can be used to access the location. 

First, create the required objects in AWS.

```hcl
resource "aws_s3_bucket" "external" {
  bucket = "${local.prefix}-external"
  acl    = "private"
  versioning {
    enabled = false
  }
  // destroy all objects with bucket destroy
  force_destroy = true
  tags = merge(local.tags, {
    Name = "${local.prefix}-external"
  })
}

resource "aws_s3_bucket_public_access_block" "external" {
  bucket             = aws_s3_bucket.external.id
  ignore_public_acls = true
  depends_on         = [aws_s3_bucket.external]
}

resource "aws_iam_policy" "external_data_access" {
  // Terraform's "jsonencode" function converts a
  // Terraform expression's result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "${aws_s3_bucket.external.id}-access"
    Statement = [
      {
        "Action" : [
          "s3:GetObject",
          "s3:GetObjectVersion",
          "s3:PutObject",
          "s3:PutObjectAcl",
          "s3:DeleteObject",
          "s3:ListBucket",
          "s3:GetBucketLocation"
        ],
        "Resource" : [
          aws_s3_bucket.external.arn,
          "${aws_s3_bucket.external.arn}/*"
        ],
        "Effect" : "Allow"
      }
    ]
  })
  tags = merge(local.tags, {
    Name = "${local.prefix}-unity-catalog external access IAM policy"
  })
}

resource "aws_iam_role" "external_data_access" {
  name                = "${local.prefix}-external-access"
  assume_role_policy  = data.aws_iam_policy_document.passrole_for_uc.json
  managed_policy_arns = [aws_iam_policy.external_data_access.arn]
  tags = merge(local.tags, {
    Name = "${local.prefix}-unity-catalog external access IAM role"
  })
}
```

Then create the [databricks_storage_credential](../resources/storage_credential.md) and [databricks_external_location](../resources/external_location.md) in Unity Catalog.

```hcl
resource "databricks_storage_credential" "external" {
  provider = databricks.workspace
  name     = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_grants" "external_creds" {
  provider           = databricks.workspace
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_TABLE"]
  }
}

resource "databricks_external_location" "some" {
  provider        = databricks.workspace
  name            = "external"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
}

resource "databricks_grants" "some" {
  provider          = databricks.workspace
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_TABLE", "READ_FILES"]
  }
}
```

## Configure Unity Catalog clusters

To ensure the integrity of ACLs, Unity Catalog data can be accessed only through compute resources configured with strong isolation guarantees and other security features. A Unity Catalog [databricks_cluster](../resources/cluster.md) has a  ‘Security Mode’ set to either **User Isolation** or **Single User**.

- **User Isolation** clusters can be shared by multiple users, but only SQL language is allowed. Some advanced cluster features such as library installation, init scripts and the DBFS Fuse mount are also disabled in this mode to ensure security isolation among cluster users.

```hcl
data "databricks_spark_version" "latest" {
  provider = databricks.workspace
}
data "databricks_node_type" "smallest" {
  provider   = databricks.workspace
  local_disk = true
}

resource "databricks_cluster" "unity_sql" {
  provider                = databricks.workspace
  cluster_name            = "Unity SQL"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 60
  enable_elastic_disk     = false
  num_workers             = 2
  aws_attributes {
    availability = "SPOT"
  }
  data_security_mode = "USER_ISOLATION"
}
```

- To use those advanced cluster features or languages like Python, Scala and R with Unity Catalog, one must choose **Single User** mode when launching the cluster. The cluster can only be used exclusively by a single user (by default the owner of the cluster); other users are not allowed to attach to the cluster.
The below example will create a collection of single-user [databricks_cluster](../resources/cluster.md) for each user in a group managed through SCIM provisioning. Individual user will be able to restart their cluster, but not anyone else. Terraform's `for_each` meta-attribute will help us achieve this.

First we use [databricks_group](../data-sources/group.md) and [databricks_user](../data-sources/user.md) data resources to get the list of user names, that belong to a group.

```hcl
data "databricks_group" "dev" {
  provider     = databricks.workspace
  display_name = "dev-clusters"
}

data "databricks_user" "dev" {
  provider = databricks.workspace
  for_each = data.databricks_group.dev.members
  user_id  = each.key
}
```

Once we have a specific list of user resources, we could proceed creating single-user clusters and provide permissions with `for_each = data.databricks_user.dev` to ensure it's done for each user:

```hcl
resource "databricks_cluster" "dev" {
  for_each                = data.databricks_user.dev
  provider                = databricks.workspace
  cluster_name            = "${each.value.display_name} unity cluster"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 10
  enable_elastic_disk     = false
  num_workers             = 2
  aws_attributes {
    availability = "SPOT"
  }
  data_security_mode = "SINGLE_USER"
  single_user_name   = each.value.user_name
}

resource "databricks_permissions" "dev_restart" {
  for_each   = data.databricks_user.dev
  provider   = databricks.workspace
  cluster_id = databricks_cluster.dev[each.key].cluster_id
  access_control {
    user_name        = each.value.user_name
    permission_level = "CAN_RESTART"
  }
}
```
