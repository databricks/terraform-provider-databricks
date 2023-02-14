---
page_title: "Unity Catalog set up on Google Cloud"
---

# Deploying pre-requisite resources and enabling Unity Catalog

Databricks Unity Catalog brings fine-grained governance and security to Lakehouse data using a familiar, open interface. You can use Terraform to deploy the underlying cloud resources and Unity Catalog objects automatically, using a programmatic approach.

This guide uses the following variables in configurations:

- `databricks_workspace_url`: Value of `workspace_url` attribute from [databricks_mws_workspaces](../resources/mws_workspaces.md#attribute-reference) resource.

This guide is provided as-is and you can use this guide as the basis for your custom Terraform module.

To get started with Unity Catalog, this guide takes you throw the following high-level steps:

- [Deploying pre-requisite resources and enabling Unity Catalog](#deploying-pre-requisite-resources-and-enabling-unity-catalog)
  - [Provider initialization](#provider-initialization)
  - [Configure Google Cloud objects](#configure-google-cloud-objects)
  - [Create a Unity Catalog metastore and link it to workspaces](#create-a-unity-catalog-metastore-and-link-it-to-workspaces)
  - [Create Unity Catalog objects in the metastore](#create-unity-catalog-objects-in-the-metastore)
  - [Configure external tables and credentials](#configure-external-tables-and-credentials)
  - [Configure Unity Catalog clusters](#configure-unity-catalog-clusters)

## Provider initialization

Initialize the 3 providers to set up the required resources. See [Databricks provider authentication](../index.md#authenticating-with-hostname,-username,-and-password), [Azure AD provider authentication](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory) and [Azure provider authentication](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure) for more details.

Define the required variables, and calculate the local values

```hcl
variable "databricks_workspace_url" {}

variable "databricks_workspace_id" {}

variable "location" {}

variable "project" {
  type    = string
  default = "<my-project-id>"
}

//generate a random string as the prefix for GCP resources, to ensure uniqueness
resource "random_string" "naming" {
  special = false
  upper   = false
  length  = 6
}

locals {
  prefix = "unity${random_string.naming.result}"
}
```

```hcl
terraform {
  required_providers {
    databricks = {
      source = "databricks/databricks"
    }
    google = {
      source  = "hashicorp/google"
    }
    random = {
      source = "hashicorp/random"
    }    
  }
}

provider "google" {
  project = var.project
}

provider "databricks" {
  host = var.databricks_workspace_url
}
```

## Configure Google Cloud objects

The first step is to create the required Google Cloud objects:

- A GCS bucket which is the default storage location for managed tables in Unity Catalog. Please use a dedicated bucket for each metastore.
- A service account that provides Unity Catalog permissions to access and manage data in the bucket and a service account key.

```hcl
resource "google_storage_bucket" "unity_metastore" {
  name          = "${local.prefix}-metastore"
  location      = var.location
  force_destroy = true
}

resource "google_service_account" "unity_sa" {
  account_id   = "unity-sa"
  display_name = "Service Account for Unity Catalog"
}

resource "google_storage_bucket_iam_member" "unity_sa_admin" {
  bucket = google_storage_bucket.unity_metastore.name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.unity_sa.email}"
}

resource "google_storage_bucket_iam_member" "unity_sa_reader" {
  bucket = google_storage_bucket.unity_metastore.name
  role = "roles/storage.legacyBucketReader"
  member = "serviceAccount:${google_service_account.unity_sa.email}"
}

resource "google_service_account_key" "mykey" {
  service_account_id = google_service_account.unity_sa.name
}
```

## Create a Unity Catalog metastore and link it to workspaces

A [databricks_metastore](../resources/metastore.md) is the top level container for data in Unity Catalog. You can only create a single metastore for each region in which your organization operates, and attach workspaces to the metastore. Each workspace will have the same view of the data you manage in Unity Catalog.

```hcl
resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "gs://${google_storage_bucket.unity_metastore.name}"
  force_destroy = true
}

resource "databricks_metastore_data_access" "first" {
  metastore_id = databricks_metastore.this.id
  name         = "the-keys"
  gcp_service_account_key {
    email          = google_service_account.unity_sa.email
    private_key_id = google_service_account_key.mykey.id
    private_key    = google_service_account_key.mykey.private_key
  }

  is_default = true
}

resource "databricks_metastore_assignment" "this" {
  workspace_id         = var.databricks_workspace_id
  metastore_id         = databricks_metastore.this.id
  default_catalog_name = "hive_metastore"
}
```

## Create Unity Catalog objects in the metastore

Each metastore exposes a 3-level namespace (catalog-schema-table) by which data can be organized.

```hcl
resource "databricks_catalog" "sandbox" {
  metastore_id = databricks_metastore.this.id
  name         = "sandbox"
  comment      = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
  depends_on = [databricks_metastore_assignment.default_metastore]
}

resource "databricks_grants" "sandbox" {
  catalog = databricks_catalog.sandbox.name
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
  catalog_name = databricks_catalog.sandbox.id
  name         = "things"
  comment      = "this database is managed by terraform"
  properties = {
    kind = "various"
  }
}

resource "databricks_grants" "things" {
  schema = databricks_schema.things.id
  grant {
    principal  = "Data Engineers"
    privileges = ["USAGE"]
  }
}
```

## Configure external tables and credentials

To work with external tables, Unity Catalog introduces two new objects to access and work with external cloud storage:

- [databricks_storage_credential](../resources/storage_credential.md) represent authentication methods to access cloud storage. Storage credentials are access-controlled to determine which users can use the credential.
- [databricks_external_location](../resources/external_location.md) are objects that combine a cloud storage path with a Storage Credential that can be used to access the location.

First, create the required objects in GCP.

```hcl
resource "google_storage_bucket" "ext_bucket" {
  name          = "${local.prefix}-ext-bucket"
  location      = var.location
  force_destroy = true
}

resource "google_service_account" "unity_credential" {
  account_id   = "unity-credential"
  display_name = "Service Account for Unity Catalog"
}

resource "google_storage_bucket_iam_member" "unity_cred_admin" {
  bucket = google_storage_bucket.unity_metastore.name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.unity_credential.email}"
}

resource "google_storage_bucket_iam_member" "unity_cred_reader" {
  bucket = google_storage_bucket.unity_metastore.name
  role = "roles/storage.legacyBucketReader"
  member = "serviceAccount:${google_service_account.unity_credential.email}"
}

resource "google_service_account_key" "my_cred_key" {
  service_account_id = google_service_account.unity_credential.name
}
```

Then create the [databricks_storage_credential](../resources/storage_credential.md) and [databricks_external_location](../resources/external_location.md) in Unity Catalog.

```hcl
resource "databricks_storage_credential" "external" {
  name = google_service_account.unity_credential.name

  gcp_service_account_key {
    email          = google_service_account.unity_credential.email
    private_key_id = google_service_account_key.my_cred_key.id
    private_key    = google_service_account_key.my_cred_key.private_key
  }
  comment = "Managed by TF"
  depends_on = [
    databricks_metastore_assignment.this
  ]
}

resource "databricks_grants" "external_creds" {
  storage_credential = databricks_storage_credential.external.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_TABLE"]
  }
}

resource "databricks_external_location" "some" {
  name = "external"
  url = "gs://${google_storage_bucket.ext_bucket.name}"

  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
  depends_on = [
    databricks_metastore_assignment.this
  ]
}

resource "databricks_grants" "some" {
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_TABLE", "READ_FILES"]
  }
}
```

## Configure Unity Catalog clusters

To ensure the integrity of ACLs, Unity Catalog data can be accessed only through compute resources configured with strong isolation guarantees and other security features. A Unity Catalog [databricks_cluster](../resources/cluster.md) has a  ‘Security Mode’ set to either **User Isolation** or **Single User**.

- **User Isolation** clusters can be shared by multiple users, but only Python (using DBR>=11.1) and SQL languages are allowed. Some advanced cluster features such as library installation, init scripts and the DBFS Fuse mount are also disabled in this mode to ensure security isolation among cluster users.

```hcl
data "databricks_spark_version" "latest" {
}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_cluster" "unity_sql" {
  cluster_name            = "Unity SQL"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 60
  enable_elastic_disk     = false
  num_workers             = 2
  azure_attributes {
    availability = "SPOT"
  }
  data_security_mode = "USER_ISOLATION"
  # need to wait until the metastore is assigned
  depends_on = [
    databricks_metastore_assignment.this
  ]
}
```

- To use those advanced cluster features or languages like Machine Learning Runtime, Streaming, Scala and R with Unity Catalog, one must choose **Single User** mode when launching the cluster. The cluster can only be used exclusively by a single user (by default the owner of the cluster); other users are not allowed to attach to the cluster.
The below example will create a collection of single-user [databricks_cluster](../resources/cluster.md) for each user in a group managed through SCIM provisioning. Individual user will be able to restart their cluster, but not anyone else. Terraform's `for_each` meta-attribute will help us achieve this.

First we use [databricks_group](../data-sources/group.md) and [databricks_user](../data-sources/user.md) data resources to get the list of user names that belong to a group.

```hcl
data "databricks_group" "dev" {
  display_name = "dev-clusters"
}

data "databricks_user" "dev" {
  for_each = data.databricks_group.dev.members
  user_id  = each.key
}
```

Once we have a specific list of user resources, we could proceed creating single-user clusters and provide permissions with `for_each = data.databricks_user.dev` to ensure it's done for each user:

```hcl
resource "databricks_cluster" "dev" {
  for_each                = data.databricks_user.dev
  cluster_name            = "${each.value.display_name} unity cluster"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 10
  num_workers             = 2
  data_security_mode = "SINGLE_USER"
  single_user_name   = each.value.user_name
  # need to wait until the metastore is assigned
  depends_on = [
    databricks_metastore_assignment.this
  ]
}

resource "databricks_permissions" "dev_restart" {
  for_each   = data.databricks_user.dev
  cluster_id = databricks_cluster.dev[each.key].cluster_id
  access_control {
    user_name        = each.value.user_name
    permission_level = "CAN_RESTART"
  }
}
```
