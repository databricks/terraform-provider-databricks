---
subcategory: "Postgres"
---
# databricks_postgres_catalog Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `catalog_id` (string, required) - The ID in the Unity Catalog.
  It becomes the full resource name, for example "my_catalog" becomes "catalogs/my_catalog"
* `spec` (CatalogCatalogSpec, optional) - The desired state of the Catalog
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,required) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### CatalogCatalogSpec
* `postgres_database` (string, required) - The name of the Postgres database inside the specified Lakebase project and branch to be associated with the UC catalog.
  This database must already exist, unless create_database_if_missing is set to true on creation.
  
  A database can only be registered with one UC catalog at a time.
  To re-register a database with a different catalog, the existing catalog must be deleted first.
  
  A child branch inherits the fact of parent's registration. This means the same-named database
  in a child branch cannot be registered with a second catalog
  while the parent's registration exists. To allow registering the database of a child branch,
  drop and recreate the database on the child branch.
  This removes the fact of parent's registration from this branch only.
  
  Doing Point In Time Restore (PITR) prior to the moment before the Postgres DB was registered
  in the Catalog drops the fact of registration of the database. So the user should avoid doing so
* `branch` (string, optional) - The resource path of the branch associated with the catalog.
  
  Format: projects/{project_id}/branches/{branch_id}
* `create_database_if_missing` (boolean, optional) - If set to true, the specified postgres_database is created on behalf of the calling user
  if it does not already exist. In this case, the calling user has a role created for
  them in Postgres if they do not already have one.
  
  Defaults to false, meaning that the request fails if the specified postgres_database does not already exist

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - A timestamp indicating when the catalog was created
* `name` (string) - Output only. The full resource path of the catalog.
  
  Format: "catalogs/{catalog_id}"
* `status` (CatalogCatalogStatus) - The observed state of the Catalog
* `uid` (string) - System-generated unique identifier for the catalog
* `update_time` (string) - A timestamp indicating when the catalog was last updated

### CatalogCatalogStatus
* `branch` (string) - The resource path of the branch associated with the catalog.
  
  Format: projects/{project_id}/branches/{branch_id}
* `postgres_database` (string) - The name of the Postgres database associated with the catalog
* `project` (string) - The resource path of the project associated with the catalog.
  
  Format: projects/{project_id}

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_postgres_catalog.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_postgres_catalog.this "name"
```