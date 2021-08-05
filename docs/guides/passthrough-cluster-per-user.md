---
page_title: "Dynamic Passthrough Clusters for a Group"
---

# Dynamic Passthrough Clusters

This example addresses a pretty common use-case: data science team, which is managed as a group through SCIM provisioning, needs a collection of individual passthrough [databricks_cluster](../resources/cluster.md), which they should be able to restart. It could simply be achieved by [databricks_group](../data-sources/group.md) and [databricks_user](../data-sources/user.md) data resources to get the list of user names, that belong to a group. Terraform's `for_each` meta-attribute helps to do this easily.

```hcl
data "databricks_group" "dev" {
  display_name = "dev-clusters"
}

data "databricks_user" "dev" {
  for_each = data.databricks_group.dev.members
  user_id = each.key
}
```

Once we have a specific list of user resources, we could proceed creating clusters and permissions with `for_each = data.databricks_user.dev` to ensure it's done for each user:

```hcl
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_cluster" "dev" {
  for_each = data.databricks_user.dev

  cluster_name            = "${each.value.display_name} dev cluster"
  single_user_name        = each.value.user_name

  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 10

  spark_conf = {
    # Single-node
    "spark.databricks.cluster.profile" : "singleNode"
    "spark.master" : "local[*]",

    # Passthrough
    "spark.databricks.passthrough.enabled": "true"
  }

  custom_tags = {
    "ResourceClass" = "SingleNode"
  }
}

resource "databricks_permissions" "dev_restart" {
  for_each = data.databricks_user.dev
  cluster_id = databricks_cluster.dev[each.key].cluster_id
  access_control {
    user_name = each.value.user_name
    permission_level = "CAN_RESTART"
  }
}
```