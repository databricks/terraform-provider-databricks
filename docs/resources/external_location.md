---
subcategory: "Unity Catalog"
---
# databricks_external_location Resource

-> **Public Preview** This feature is in [Public Preview](https://docs.databricks.com/release-notes/release-types.html). Contact your Databricks representative to request access. 

To work with external tables, Unity Catalog introduces two new objects to access and work with external cloud storage:
- [databricks_storage_credential](storage_credential.md) represent authentication methods to access cloud storage (e.g. an IAM role for Amazon S3 or a service principal for Azure Storage). Storage credentials are access-controlled to determine which users can use the credential.
- `databricks_external_location` are objects that combine a cloud storage path with a Storage Credential that can be used to access the location. 

## Example Usage

For AWS

```hcl
resource "databricks_storage_credential" "external" {
  name = aws_iam_role.external_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.external_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_external_location" "some" {
  name            = "external"
  url             = "s3://${aws_s3_bucket.external.id}/some"
  credential_name = databricks_storage_credential.external.id
  comment         = "Managed by TF"
}

resource "databricks_grants" "some" {
  external_location = databricks_external_location.some.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_TABLE", "READ_FILES"]
  }
}
```

For Azure

```hcl
resource "databricks_storage_credential" "external" {
  name = azuread_application.ext_cred.display_name
  azure_service_principal {
    directory_id   = var.tenant_id
    application_id = azuread_application.ext_cred.application_id
    client_secret  = azuread_application_password.ext_cred.value
  }
  comment = "Managed by TF"
  depends_on = [
    databricks_metastore_assignment.this
  ]
}

resource "databricks_external_location" "some" {
  name = "external"
  url = format("abfss://%s@%s.dfs.core.windows.net/",
    azurerm_storage_account.ext_storage.name,
  azurerm_storage_container.ext_storage.name)
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

## Argument Reference

The following arguments are required:

* `name` - Name of External Location, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.
* `url` - Path URL in cloud storage, of the form: `s3://[bucket-host]/[bucket-dir]` (AWS), `abfss://[user]@[host]/[path]` (Azure).
* `credential_name` - Name of the [databricks_storage_credential](storage_credential.md) to use with this External Location.
* `owner` - (Optional) Username/groupname of External Location owner. Currently this field can only be changed after the resource is created.
* `comment` - (Optional) User-supplied free-form text.
* `skip_validation` - (Optional) Suppress validation errors if any & force save the external location

## Import

This resource can be imported by name:

```bash
$ terraform import databricks_external_location.this <name>
```
