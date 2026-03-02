---
subcategory: "Deployment"
---
# databricks_gcp_unity_catalog_policy Data Source

This data source constructs the list of GCP IAM permissions required for the Unity Catalog file events custom role on GCP.

-> This data source has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/gcp/en/admin/cloud-configurations/gcp/permissions) in case of any questions.

## Example Usage

```hcl
resource "databricks_storage_credential" "ext" {
  name = "the-creds"
  databricks_gcp_service_account {}
}

data "databricks_gcp_unity_catalog_policy" "this" {
  databricks_google_service_account = databricks_storage_credential.ext.databricks_gcp_service_account[0].email
}

resource "google_project_iam_custom_role" "uc_file_events" {
  role_id     = "uc_file_events"
  title       = "Unity Catalog File Events"
  permissions = data.databricks_gcp_unity_catalog_policy.this.permissions
}

resource "google_project_iam_member" "uc_file_events" {
  project = var.google_project
  role    = google_project_iam_custom_role.uc_file_events.id
  member  = "serviceAccount:${data.databricks_gcp_unity_catalog_policy.this.databricks_google_service_account}"
}

resource "google_storage_bucket_iam_member" "unity_sa_admin" {
  bucket = google_storage_bucket.this.name
  role   = "roles/storage.objectAdmin"
  member = "serviceAccount:${data.databricks_gcp_unity_catalog_policy.this.databricks_google_service_account}"
}

resource "google_storage_bucket_iam_member" "unity_sa_reader" {
  bucket = google_storage_bucket.this.name
  role   = "roles/storage.legacyBucketReader"
  member = "serviceAccount:${data.databricks_gcp_unity_catalog_policy.this.databricks_google_service_account}"
}
```

## Argument Reference

* `databricks_google_service_account` (Required) - The email of the Databricks-managed GCP service account. Typically obtained from `databricks_storage_credential.*.databricks_gcp_service_account[0].email` or `databricks_metastore_data_access.*.databricks_gcp_service_account[0].email`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `permissions` - List of GCP IAM permissions for the Unity Catalog file events custom role, to be used with [google_project_iam_custom_role](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam_custom_role).
* `databricks_google_service_account` - The email of the Databricks-managed GCP service account (same as the input). Exported for use in `google_storage_bucket_iam_member` and similar resources.

## Related Resources

The following resources are used in the same context:

* [Unity Catalog set up on GCP](../guides/unity-catalog-gcp.md) - guide for setting up Unity Catalog on GCP.
* [databricks_storage_credential](../resources/storage_credential.md) - resource for creating storage credentials.
* [databricks_gcp_crossaccount_policy](gcp_crossaccount_policy.md) - data source for workspace creator permissions on GCP.
