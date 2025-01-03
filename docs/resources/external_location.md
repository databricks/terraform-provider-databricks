---
subcategory: "Unity Catalog"
---
# databricks_external_location Resource

-> This resource can only be used with a workspace-level provider!

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
    privileges = ["CREATE_EXTERNAL_TABLE", "READ_FILES"]
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

For GCP

```hcl
resource "databricks_storage_credential" "ext" {
  name = "the-creds"
  databricks_gcp_service_account {}
}

resource "databricks_external_location" "some" {
  name = "the-ext-location"
  url  = "gs://${google_storage_bucket.ext_bucket.name}"

  credential_name = databricks_storage_credential.ext.id
  comment         = "Managed by TF"
}
```

Example `encryption_details` specifying SSE_S3 encryption:

```hcl
encryption_details {
  sse_encryption_details {
    algorithm = "AWS_SSE_S3"
  }
}
```

Example `encryption_details` specifying SSE_KMS encryption with KMS key that has ID "some_key_arn":

```hcl
encryption_details {
  sse_encryption_details {
    algorithm       = "AWS_SSE_KMS"
    aws_kms_key_arn = "some_key_arn"
  }
}
```

## Argument Reference

The following arguments are required:

- `name` - Name of External Location, which must be unique within the [databricks_metastore](metastore.md). Change forces creation of a new resource.
- `url` - Path URL in cloud storage, of the form: `s3://[bucket-host]/[bucket-dir]` (AWS), `abfss://[user]@[host]/[path]` (Azure), `gs://[bucket-host]/[bucket-dir]` (GCP).
- `credential_name` - Name of the [databricks_storage_credential](storage_credential.md) to use with this external location.
- `owner` - (Optional) Username/groupname/sp application_id of the external location owner.
- `comment` - (Optional) User-supplied free-form text.
- `skip_validation` - (Optional) Suppress validation errors if any & force save the external location
- `fallback` - (Optional) Indicates whether fallback mode is enabled for this external location. When fallback mode is enabled, the access to the location falls back to cluster credentials if UC credentials are not sufficient.
- `read_only` - (Optional) Indicates whether the external location is read-only.
- `force_destroy` - (Optional) Destroy external location regardless of its dependents.
- `force_update` - (Optional) Update external location regardless of its dependents.
- `access_point` - (Optional) The ARN of the s3 access point to use with the external location (AWS).
- `encryption_details` - (Optional) The options for Server-Side Encryption to be used by each Databricks s3 client when connecting to S3 cloud storage (AWS).
- `isolation_mode` - (Optional) Whether the external location is accessible from all workspaces or a specific set of workspaces. Can be `ISOLATION_MODE_ISOLATED` or `ISOLATION_MODE_OPEN`. Setting the external location to `ISOLATION_MODE_ISOLATED` will automatically allow access from the current workspace.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - ID of this external location - same as `name`.
- `created_at` - Time at which this external location was created, in epoch milliseconds.
- `created_by` -  Username of external location creator.
- `credential_id` - Unique ID of the location's storage credential.
- `updated_at` - Time at which external location this was last modified, in epoch milliseconds.
- `updated_by` - Username of user who last modified the external location.

## Import

This resource can be imported by `name`:

```bash
terraform import databricks_external_location.this <name>
```
