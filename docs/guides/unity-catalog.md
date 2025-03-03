---
page_title: "Unity Catalog set up on AWS"
---

# Deploying pre-requisite resources and enabling Unity Catalog

**Note**
If your workspace was enabled for Unity Catalog automatically, this guide does not apply to you. See [this guide](unity-catalog-default.md) instead.

**Note**
Except for metastore, metastore assignment and storage credential objects, Unity Catalog APIs are accessible via **workspace-level APIs**. This design may change in the future.

Databricks Unity Catalog brings fine-grained governance and security to Lakehouse data using a familiar, open interface. You can use Terraform to deploy the underlying cloud resources and Unity Catalog objects automatically, using a programmatic approach.

This guide creates a metastore without a storage root location or credential to maintain strict separation of storage across catalogs or environments.

This guide uses the following variables in configurations:

- `databricks_client_id`: The `client_id` is the `application_id` of a [Service Principal](../resources/service_principal.md) that has account-level admin permission on [https://accounts.cloud.databricks.com](https://accounts.cloud.databricks.com).
- `databricks_client_secret`: The secret of the above service principal.
- `databricks_account_id`: The numeric ID for your Databricks account. When you are logged in, it appears in the top right corner of the [Databricks Account Console](https://accounts.cloud.databricks.com/) or [Azure Databricks Account Console](https://accounts.azuredatabricks.net).
- `databricks_workspace_url`: Value of `workspace_url` attribute from [databricks_mws_workspaces](../resources/mws_workspaces.md#attribute-reference) resource.

This guide is provided as-is and you can use this guide as the basis for your custom Terraform module.

To get started with Unity Catalog, this guide takes you through the following high-level steps:

- [Deploying pre-requisite resources and enabling Unity Catalog](#deploying-pre-requisite-resources-and-enabling-unity-catalog)
  - [Provider initialization](#provider-initialization)
  - [Create users and groups](#create-users-and-groups)
  - [Create a Unity Catalog metastore and link it to workspaces](#create-a-unity-catalog-metastore-and-link-it-to-workspaces)
  - [Configure external locations and credentials](#configure-external-locations-and-credentials)
  - [Create Unity Catalog objects in the metastore](#create-unity-catalog-objects-in-the-metastore)
  - [Configure Unity Catalog clusters](#configure-unity-catalog-clusters)

## Provider initialization

Initialize [provider with `mws` alias](https://www.terraform.io/language/providers/configuration#alias-multiple-provider-configurations) to set up account-level resources. See [provider authentication](../index.md#authenticating-with-databricks-managed-service-principal) for more details.

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
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
  alias         = "mws"
  host          = "https://accounts.cloud.databricks.com"
  account_id    = var.databricks_account_id
  client_id     = var.databricks_client_id
  client_secret = var.databricks_client_secret
}

// initialize provider at workspace level, to create UC resources
provider "databricks" {
  alias         = "workspace"
  host          = var.databricks_workspace_url
  client_id     = var.databricks_client_id
  client_secret = var.databricks_client_secret
}
```

Define the required variables

```hcl
variable "databricks_client_id" {}
variable "databricks_client_secret" {}
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

A [databricks_metastore](../resources/metastore.md) is the top level container for data in Unity Catalog. You can only create a single metastore for each region in which your organization operates, and attach workspaces to the metastore. Each workspace will have the same view of the data you manage in Unity Catalog.

```hcl
resource "databricks_metastore" "this" {
  provider      = databricks.mws
  name          = "primary"
  owner         = var.unity_admin_group
  region        = var.region
  force_destroy = true
}

resource "databricks_metastore_assignment" "default_metastore" {
  provider             = databricks.mws
  for_each             = toset(var.databricks_workspace_ids)
  workspace_id         = each.key
  metastore_id         = databricks_metastore.this.id
}


```

## Configure external locations and credentials

Unity Catalog introduces two new objects to access and work with external cloud storage:

- [databricks_storage_credential](../resources/storage_credential.md) represent authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a service principal for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
- [databricks_external_location](../resources/external_location.md) are objects that combine a cloud storage path with a Storage Credential that can be used to access the location.

First, we need to create the storage credential in Databricks before creating the IAM role in AWS. This is because the external ID of the Databricks storage credential is required in the IAM role trust policy.

```hcl
data "aws_caller_identity" "current" {}
locals {
  uc_iam_role = "${local.prefix}-uc-access"
}

resource "databricks_storage_credential" "external" {
  name = "${local.prefix}-external-access"
  //cannot reference aws_iam_role directly, as it will create circular dependency
  aws_iam_role {
    role_arn = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/${local.uc_iam_role}"
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
```

Then we can create the required objects in AWS

```hcl
resource "aws_s3_bucket" "external" {
  bucket = "${local.prefix}-external"
  // destroy all objects with bucket destroy
  force_destroy = true
  tags = merge(var.tags, {
    Name = "${local.prefix}-external"
  })
}

resource "aws_s3_bucket_versioning" "external_versioning" {
  bucket = aws_s3_bucket.external.id
  versioning_configuration {
    status = "Disabled"
  }
}

data "databricks_aws_unity_catalog_assume_role_policy" "this" {
  aws_account_id = data.aws_caller_identity.current.account_id
  role_name      = local.uc_iam_role
  external_id    = databricks_storage_credential.external.aws_iam_role[0].external_id
}

data "databricks_aws_unity_catalog_policy" "this" {
  aws_account_id = data.aws_caller_identity.current.account_id
  bucket_name    = aws_s3_bucket.external.id
  role_name      = local.uc_iam_role
}

resource "aws_iam_policy" "external_data_access" {
  policy = data.databricks_aws_unity_catalog_policy.this.json
  tags = merge(var.tags, {
    Name = "${local.prefix}-unity-catalog external access IAM policy"
  })
}

resource "aws_iam_role" "external_data_access" {
  name                = local.uc_iam_role
  assume_role_policy  = data.databricks_aws_unity_catalog_assume_role_policy.this.json
  managed_policy_arns = [aws_iam_policy.external_data_access.arn]
  tags = merge(var.tags, {
    Name = "${local.prefix}-unity-catalog external access IAM role"
  })
}
```

Then we can create the [databricks_external_location](../resources/external_location.md) in Unity Catalog.

```hcl

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
    privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
  }
}
```

## Create Unity Catalog objects in the metastore

Each metastore exposes a 3-level namespace (catalog-schema-table) by which data can be organized.

```hcl
resource "databricks_catalog" "sandbox" {
  provider     = databricks.workspace
  storage_root = "s3://${aws_s3_bucket.external.id}/some"
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
    privileges = ["USE_CATALOG", "CREATE"]
  }
  grant {
    principal  = "Data Engineers"
    privileges = ["USE_CATALOG"]
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
    privileges = ["USE_SCHEMA"]
  }
}
```

## Configure Unity Catalog clusters

To ensure the integrity of ACLs, Unity Catalog data can be accessed only through compute resources configured with strong isolation guarantees and other security features. A Unity Catalog [databricks_cluster](../resources/cluster.md) has the access mode set to either **Shared** or **Single User**.

- **Shared** clusters can be shared by multiple users, but has certain [limitations](https://docs.databricks.com/en/compute/access-mode-limitations.html#shared-access-mode-limitations-on-unity-catalog)

```hcl
data "databricks_spark_version" "latest" {
  provider = databricks.workspace
}
data "databricks_node_type" "smallest" {
  provider   = databricks.workspace
  local_disk = true
}

resource "databricks_cluster" "unity_shared" {
  provider                = databricks.workspace
  cluster_name            = "Shared clusters"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 60
  enable_elastic_disk     = false
  num_workers             = 2
  aws_attributes {
    availability = "SPOT"
  }
  data_security_mode = "USER_ISOLATION"
  depends_on = [
    databricks_metastore_assignment.this
  ]
}
```

- To use those advanced cluster features or languages like Machine Learning Runtime and R with Unity Catalog, one must choose **Single User** mode when launching the cluster. The cluster can only be used exclusively by a single user (by default the owner of the cluster); other users are not allowed to attach to the cluster.
The below example will create a collection of single-user [databricks_cluster](../resources/cluster.md) for each user in a group managed through SCIM provisioning. Individual user will be able to restart their cluster, but not anyone else. Terraform's `for_each` meta-attribute will help us achieve this.

First we use [databricks_group](../data-sources/group.md) and [databricks_user](../data-sources/user.md) data resources to get the list of user names that belong to a group.

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
  depends_on = [
    databricks_metastore_assignment.this
  ]
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
