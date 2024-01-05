---
subcategory: "Compute"
---

# databricks_cluster_policy Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves information about [databricks_cluster_policy](../resources/cluster_policy.md).

## Example Usage

Referring to a cluster policy by name:

```hcl
data "databricks_cluster_policy" "personal" {
  name = "Personal Compute"
}

resource "databricks_cluster" "my_cluster" {
  policy_id = data.databricks_cluster_policy.personal.id
  # ...
}
```

## Argument Reference

Data source allows you to pick a cluster policy by the following attribute

- `name` - Name of the cluster policy. The cluster policy must exist before this resource can be planned.

## Attribute Reference

Data source exposes the following attributes:

- `id` - The id of the cluster policy.
- `definition` - Policy definition: JSON document expressed in [Databricks Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definition).
- `description` - Additional human-readable description of the cluster policy.
- `policy_family_id` - ID of the policy family.
- `policy_family_definition_overrides` - Policy definition JSON document expressed in Databricks [Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definitions).
- `is_default` - If true, policy is a default policy created and managed by Databricks.
- `max_clusters_per_user` - Max number of clusters per user that can be active using this policy.
