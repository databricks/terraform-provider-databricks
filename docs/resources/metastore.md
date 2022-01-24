---
subcategory: "Unity Catalog"
---
# databricks_metastore Resource

-> **Private Preview** This feature is in [Private Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

A metastore is the top-level container of objects in Unity Catalog. It stores data assets (tables and views) and the permissions that govern access to them. Databricks account admins can create metastores and assign them to Databricks workspaces in order to control which workloads use each metastore.

Unity Catalog offers a new metastore with built in security and auditing. This is distinct to the metastore used in previous versions of Databricks (based on the Hive Metastore).

## Example Usage

```hcl
resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = "uc admins"
  force_destroy = true
}

resource "databricks_metastore_assignment" "this" {
  metastore_id = databricks_metastore.this.id
  workspace_id = local.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of metastore.
* `storage_root` - Path on cloud storage account, where managed [databricks_table](table.md) are stored. Change forces creation of a new resource.
* `owner` - (Optional) Username/groupname of Metastore owner.
* `force_destroy` - (Optional) Destroy metastore regardless of its contents.

## Import

This resource can be imported by ID:

```bash
$ terraform import databricks_metastore.this <id>
```
