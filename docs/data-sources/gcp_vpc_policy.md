---
subcategory: "Deployment"
---
# databricks_gcp_vpc_policy Data Source

This data source constructs the list of GCP IAM permissions required for the Databricks workspace creator custom role in the **VPC GCP project** (compute, networking, firewall management).

In a [Shared VPC](https://docs.databricks.com/gcp/en/security/network/classic/customer-managed-vpc) setup, the VPC resides in a separate host project. This data source covers the permissions needed in that host project. When the VPC is in the same project as the workspace, combine these permissions with [databricks_gcp_crossaccount_policy](gcp_crossaccount_policy.md).

-> This data source has an evolving API, which may change in future versions of the provider. Please always consult [latest documentation](https://docs.databricks.com/gcp/en/admin/cloud-configurations/gcp/permissions) in case of any questions.

## Example Usage

### Shared VPC workspace with all features enabled

```hcl
data "databricks_gcp_vpc_policy" "this" {
  enable_byovpc = true
  enable_cmk    = true
  enable_psc    = true
}

resource "google_project_iam_custom_role" "vpc_creator" {
  project     = var.vpc_project_id
  role_id     = "databricks_vpc_creator"
  title       = "Databricks Workspace Creator Role for VPC Project"
  permissions = data.databricks_gcp_vpc_policy.this.permissions
}

resource "google_project_iam_member" "vpc_creator" {
  project = var.vpc_project_id
  role    = google_project_iam_custom_role.vpc_creator.id
  member  = "serviceAccount:${var.databricks_google_service_account}"
}
```

### Single-project workspace (VPC in same project as workspace)

When the VPC resides in the same project, merge these permissions with `databricks_gcp_crossaccount_policy`. See the [databricks_gcp_crossaccount_policy](gcp_crossaccount_policy.md) documentation for a complete example.

## Argument Reference

* `enable_byovpc` - (Optional) Set to `true` to include additional permissions required for [customer-managed VPC (Bring Your Own VPC)](https://docs.databricks.com/gcp/en/security/network/classic/customer-managed-vpc). Adds `compute.subnetworks.*` permissions. Defaults to `false`.
* `enable_cmk` - (Optional) Set to `true` to include additional permissions required for [Customer-Managed Keys](https://docs.databricks.com/gcp/en/security/keys/customer-managed-keys). Adds `cloudkms.cryptoKeys.*` permissions. Defaults to `false`.
* `enable_psc` - (Optional) Set to `true` to include additional permissions required for [Private Service Connect](https://docs.databricks.com/gcp/en/security/network/classic/private-service-connect). Adds `compute.forwardingRules.*` permissions. Defaults to `false`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `permissions` - List of GCP IAM permissions for the workspace creator custom role in the VPC project, to be used with [google_project_iam_custom_role](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam_custom_role).

## Related Resources

The following resources are used in the same context:

* [Provisioning GCP Databricks workspaces](../guides/gcp-workspace.md) - guide for setting up workspaces on GCP.
* [Provisioning GCP Databricks workspaces with PSC](../guides/gcp-private-service-connect-workspace.md) - guide for setting up workspaces with Private Service Connect.
* [databricks_gcp_crossaccount_policy](gcp_crossaccount_policy.md) - data source for workspace project permissions on GCP.
* [databricks_gcp_unity_catalog_policy](gcp_unity_catalog_policy.md) - data source for Unity Catalog permissions on GCP.
