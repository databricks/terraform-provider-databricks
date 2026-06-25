---
subcategory: "Deployment"
---
# databricks_gcp_crossaccount_policy Data Source

This data source constructs the list of GCP IAM permissions required for the Databricks workspace creator custom role in the **workspace GCP project** (IAM, service accounts, project management, service usage).

For the networking permissions required in the VPC project, use [databricks_gcp_vpc_policy](gcp_vpc_policy.md) instead.

-> This data source has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/gcp/en/admin/cloud-configurations/gcp/permissions) in case of any questions.

## Example Usage

### Single-project workspace (VPC in same project as workspace)

When the VPC resides in the same project as the workspace, merge the permissions from both `databricks_gcp_crossaccount_policy` and `databricks_gcp_vpc_policy` into a single custom role:

```hcl
data "databricks_gcp_crossaccount_policy" "this" {}

data "databricks_gcp_vpc_policy" "this" {
  enable_byovpc = true
}

resource "google_project_iam_custom_role" "workspace_creator" {
  role_id     = "databricks_workspace_creator"
  title       = "Databricks Workspace Creator"
  permissions = tolist(toset(concat(
    data.databricks_gcp_crossaccount_policy.this.permissions,
    data.databricks_gcp_vpc_policy.this.permissions,
  )))
}

resource "google_project_iam_member" "workspace_creator" {
  project = var.google_project
  role    = google_project_iam_custom_role.workspace_creator.id
  member  = "serviceAccount:${var.databricks_google_service_account}"
}
```

### Shared VPC workspace (VPC in a separate host project)

When using a Shared VPC, apply the workspace project role to the service project and the VPC project role to the host project:

```hcl
data "databricks_gcp_crossaccount_policy" "this" {}

resource "google_project_iam_custom_role" "workspace_creator" {
  project     = var.workspace_project_id
  role_id     = "databricks_workspace_creator"
  title       = "Databricks Workspace Creator"
  permissions = data.databricks_gcp_crossaccount_policy.this.permissions
}

resource "google_project_iam_member" "workspace_creator" {
  project = var.workspace_project_id
  role    = google_project_iam_custom_role.workspace_creator.id
  member  = "serviceAccount:${var.databricks_google_service_account}"
}
```

## Argument Reference

This data source takes no arguments. All permissions for the workspace project are always included.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `permissions` - List of GCP IAM permissions for the workspace creator custom role in the workspace project, to be used with [google_project_iam_custom_role](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam_custom_role).

## Related Resources

The following resources are used in the same context:

* [Provisioning GCP Databricks workspaces](../guides/gcp-workspace.md) - guide for setting up workspaces on GCP.
* [databricks_gcp_vpc_policy](gcp_vpc_policy.md) - data source for VPC project permissions on GCP.
* [databricks_gcp_unity_catalog_policy](gcp_unity_catalog_policy.md) - data source for Unity Catalog permissions on GCP.
