---
page_title: "Unity Catalog set up on Azure"
---

# Deploying pre-requisite resources and enabling Unity Catalog

**Note**
If your workspace was enabled for Unity Catalog automatically, this guide does not apply to you. See [this guide](unity-catalog-default.md) instead.

**Note**
Except for metastore, metastore assignment and storage credential objects, Unity Catalog APIs are accessible via **workspace-level APIs**. This design may change in the future.

Databricks Unity Catalog brings fine-grained governance and security to Lakehouse data using a familiar, open interface. You can use Terraform to deploy the underlying cloud resources and Unity Catalog objects automatically, using a programmatic approach.

This guide creates a metastore without a storage root location or credential to maintain strict separation of storage across catalogs or environments.

This guide uses the following variables in configurations:

- `databricks_resource_id`: The Azure resource ID for the Databricks workspace where Unity Catalog will be deployed. It should be of the format `/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1`. To find the resource ID, navigate to your Databricks workspace in the Azure portal, select the JSON View link on the Overview page.

This guide is provided as-is and you can use this guide as the basis for your custom Terraform module.

To get started with Unity Catalog, this guide takes you through the following high-level steps:

- [Deploying pre-requisite resources and enabling Unity Catalog](#deploying-pre-requisite-resources-and-enabling-unity-catalog)
  - [Provider initialization](#provider-initialization)
  - [Create a Unity Catalog metastore and link it to workspaces](#create-a-unity-catalog-metastore-and-link-it-to-workspaces)
  - [Configure external locations and credentials](#configure-external-locations-and-credentials)
  - [Create Unity Catalog objects in the metastore](#create-unity-catalog-objects-in-the-metastore)
  - [Configure Unity Catalog clusters](#configure-unity-catalog-clusters)

## Provider initialization

Initialize the 3 providers to set up the required resources. See [Databricks provider authentication](../index.md#authenticating-with-databricks-managed-service-principal), [Azure AD provider authentication](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory) and [Azure provider authentication](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs#authenticating-to-azure) for more details.

Define the required variables, and calculate the local values

```hcl
variable "databricks_resource_id" {
  description = "The Azure resource ID for the databricks workspace deployment."
}

locals {
  resource_regex            = "(?i)subscriptions/(.+)/resourceGroups/(.+)/providers/Microsoft.Databricks/workspaces/(.+)"
  subscription_id           = regex(local.resource_regex, var.databricks_resource_id)[0]
  resource_group            = regex(local.resource_regex, var.databricks_resource_id)[1]
  databricks_workspace_name = regex(local.resource_regex, var.databricks_resource_id)[2]
  tenant_id                 = data.azurerm_client_config.current.tenant_id
  databricks_workspace_host = data.azurerm_databricks_workspace.this.workspace_url
  databricks_workspace_id   = data.azurerm_databricks_workspace.this.workspace_id
  prefix                    = replace(replace(lower(data.azurerm_resource_group.this.name), "rg", ""), "-", "")
}

data "azurerm_resource_group" "this" {
  name = local.resource_group
}

data "azurerm_client_config" "current" {
}

data "azurerm_databricks_workspace" "this" {
  name                = local.databricks_workspace_name
  resource_group_name = local.resource_group
}
```

```hcl
terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
    }
    databricks = {
      source = "databricks/databricks"
    }
  }
}

provider "azurerm" {
  subscription_id = local.subscription_id
  features {}
}

provider "databricks" {
  host = local.databricks_workspace_host
}

provider "databricks" {
  alias      = "accounts"
  host       = "https://accounts.azuredatabricks.net"
  account_id = var.databricks_account_id
}
```

## Create a Unity Catalog metastore and link it to workspaces

A [databricks_metastore](../resources/metastore.md) is the top level container for data in Unity Catalog. You can only create a single metastore for each region in which your organization operates, and attach workspaces to the metastore. Each workspace will have the same view of the data you manage in Unity Catalog.

```hcl
resource "databricks_metastore" "this" {
  provider      = databricks.accounts
  name          = "primary"
  force_destroy = true
  region        = data.azurerm_resource_group.this.location
}

resource "databricks_metastore_assignment" "this" {
  provider             = databricks.accounts
  workspace_id         = local.databricks_workspace_id
  metastore_id         = databricks_metastore.this.id
  default_catalog_name = "hive_metastore"
}
```

## Configure external locations and credentials

Unity Catalog introduces two new objects to access and work with external cloud storage:

- [databricks_storage_credential](../resources/storage_credential.md) represent authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a managed identity for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
- [databricks_external_location](../resources/external_location.md) are objects that combine a cloud storage path with a Storage Credential that can be used to access the location.

First, create the required objects in Azure.

```hcl
resource "azurerm_databricks_access_connector" "ext_access_connector" {
  name                = "ext-databricks-mi"
  resource_group_name = data.azurerm_resource_group.this.name
  location            = data.azurerm_resource_group.this.location
  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_storage_account" "ext_storage" {
  name                     = "${local.prefix}extstorage"
  resource_group_name      = data.azurerm_resource_group.this.name
  location                 = data.azurerm_resource_group.this.location
  tags                     = data.azurerm_resource_group.this.tags
  account_tier             = "Standard"
  account_replication_type = "GRS"
  is_hns_enabled           = true
}

resource "azurerm_storage_container" "ext_storage" {
  name                  = "${local.prefix}-ext"
  storage_account_name  = azurerm_storage_account.ext_storage.name
  container_access_type = "private"
}

resource "azurerm_role_assignment" "ext_storage" {
  scope                = azurerm_storage_account.ext_storage.id
  role_definition_name = "Storage Blob Data Contributor"
  principal_id         = azurerm_databricks_access_connector.ext_access_connector.identity[0].principal_id
}
```

Then create the [databricks_storage_credential](../resources/storage_credential.md) and [databricks_external_location](../resources/external_location.md) in Unity Catalog.

```hcl
resource "databricks_storage_credential" "external" {
  name = azurerm_databricks_access_connector.ext_access_connector.name
  azure_managed_identity {
    access_connector_id = azurerm_databricks_access_connector.ext_access_connector.id
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
    privileges = ["CREATE_EXTERNAL_TABLE"]
  }
}

resource "databricks_external_location" "some" {
  name = "external"
  url = format("abfss://%s@%s.dfs.core.windows.net",
    azurerm_storage_container.ext_storage.name,
  azurerm_storage_account.ext_storage.name)

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
    privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
  }
}
```

## Create Unity Catalog objects in the metastore

Each metastore exposes a 3-level namespace (catalog-schema-table) by which data can be organized.

```hcl
resource "databricks_catalog" "sandbox" {
  name = "sandbox"
  storage_root = format("abfss://%s@%s.dfs.core.windows.net",
    azurerm_storage_container.ext_storage.name,
  azurerm_storage_account.ext_storage.name)
  comment = "this catalog is managed by terraform"
  properties = {
    purpose = "testing"
  }
  depends_on = [databricks_metastore_assignment.default_metastore]
}

resource "databricks_grants" "sandbox" {
  catalog = databricks_catalog.sandbox.name
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
    privileges = ["USE_SCHEMA"]
  }
}
```

## Configure Unity Catalog clusters

To ensure the integrity of ACLs, Unity Catalog data can be accessed only through compute resources configured with strong isolation guarantees and other security features. A Unity Catalog [databricks_cluster](../resources/cluster.md) has a  ‘Security Mode’ set to either **User Isolation** or **Single User**.

- **User Isolation** clusters can be shared by multiple users, but has certain [limitations](https://docs.databricks.com/en/compute/access-mode-limitations.html#shared-access-mode-limitations-on-unity-catalog)

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
  num_workers             = 2
  azure_attributes {
    availability = "SPOT_WITH_FALLBACK_AZURE"
  }
  data_security_mode = "SINGLE_USER"
  single_user_name   = each.value.user_name
  # need to wait until the metastore is assigned
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
