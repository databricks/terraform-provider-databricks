---
subcategory: "Compute"
---
# databricks_cluster_policy Resource

This resource creates a [cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules. The policy rules limit the attributes or attribute values available for [cluster](cluster.md) creation. [cluster](cluster.md) policies have ACLs that limit their use to specific users and groups. Only admin users can create, edit, and delete policies. Admin users also have access to all policies.

Cluster policies let you:

* Limit users to create clusters with prescribed settings.
* Simplify the user interface and enable more users to create their own clusters (by fixing and hiding some values).
* Control cost by limiting per [cluster](cluster.md) maximum cost (by setting limits on attributes whose values contribute to hourly price).

Cluster policy permissions limit which policies a user can select in the Policy drop-down when the user creates a cluster:

* If no policies have been created in the workspace, the Policy drop-down does not display.
* A user who has [cluster](cluster.md) create permission can select the `Free form` policy and create fully-configurable clusters.
* A user who has both [cluster](cluster.md) create permission and access to [cluster](cluster.md) policies can select the Free form policy and policies they have access to.
* A user that has access to only [cluster](cluster.md) policies, can select the policies they have access to.

## Example Usage

Let us take a look at an example of how you can manage two teams: Marketing and Data Engineering. In the following scenario we want the marketing team to have a really good query experience, so we enabled delta cache for them. On the other hand we want the data engineering team to be able to utilize bigger clusters so we increased the dbus per hour that they can spend. This strategy allows your marketing users and data engineering users to use Databricks in a self service manner but have a different experience in regards to security and performance. And down the line if you need to add more global settings you can propagate them through the "base [cluster](cluster.md) policy".

`modules/base-cluster-policy/main.tf` could look like:

```hcl
variable "team" {
  description = "Team that performs the work"
}

variable "policy_overrides" {
  description = "Cluster policy overrides"
}

locals {
  default_policy = {
    "dbus_per_hour" : {
      "type" : "range",
      "maxValue" : 10
    },
    "autotermination_minutes" : {
      "type" : "fixed",
      "value" : 20,
      "hidden" : true
    },
    "custom_tags.Team" : {
      "type" : "fixed",
      "value" : var.team
    }
  }
}

resource "databricks_cluster_policy" "fair_use" {
  name       = "${var.team} cluster policy"
  definition = jsonencode(merge(local.default_policy, var.policy_overrides))

  libraries {
    pypi {
      package = "databricks-sdk==0.12.0"
      // repo can also be specified here
    }
  }
  libraries {
    maven {
      coordinates = "com.oracle.database.jdbc:ojdbc8:XXXX"
    }
  }
}

resource "databricks_permissions" "can_use_cluster_policyinstance_profile" {
  cluster_policy_id = databricks_cluster_policy.fair_use.id
  access_control {
    group_name       = var.team
    permission_level = "CAN_USE"
  }
}
```

And custom instances of that base policy module for our marketing and data engineering teams would look like:

```hcl
module "marketing_compute_policy" {
  source = "../modules/databricks-cluster-policy"
  team   = "marketing"
  policy_overrides = {
    // only the marketing team will benefit from delta cache this way
    "spark_conf.spark.databricks.io.cache.enabled" : {
      "type" : "fixed",
      "value" : "true"
    },
  }
}

module "engineering_compute_policy" {
  source = "../modules/databricks-cluster-policy"
  team   = "engineering"
  policy_overrides = {
    "dbus_per_hour" : {
      "type" : "range",
      // only the engineering team are allowed to spin up big clusters
      "maxValue" : 50
    },
  }
}
```

### Overriding the built-in cluster policies

You can override built-in cluster policies by creating a `databricks_cluster_policy` resource with following attributes:

* `name` - the name of the built-in cluster policy.
* `policy_family_id` - the ID of the cluster policy family used for built-in cluster policy.
* `policy_family_definition_overrides` - settings to override in the built-in cluster policy.

You can obtain the list of defined cluster policies families using the `databricks policy-families list` command of the new [Databricks CLI](https://docs.databricks.com/en/dev-tools/cli/index.html), or via [list policy families](https://docs.databricks.com/api/workspace/policyfamilies/list) REST API.

```hcl
locals {
  personal_vm_override = {
    "autotermination_minutes" : {
      "type" : "fixed",
      "value" : 220,
      "hidden" : true
    },
    "custom_tags.Team" : {
      "type" : "fixed",
      "value" : var.team
    }
  }
}

resource "databricks_cluster_policy" "personal_vm" {
  policy_family_id                   = "personal-vm"
  policy_family_definition_overrides = jsonencode(personal_vm_override)
  name                               = "Personal Compute"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Cluster policy name. This must be unique. Length must be between 1 and 100 characters.
* `description` - (Optional) Additional human-readable description of the cluster policy.
* `definition` - Policy definition: JSON document expressed in [Databricks Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definition). Cannot be used with `policy_family_id`
* `max_clusters_per_user` - (Optional, integer) Maximum number of clusters allowed per user. When omitted, there is no limit. If specified, value must be greater than zero.
* `policy_family_definition_overrides`(Optional) Policy definition JSON document expressed in Databricks Policy Definition Language. The JSON document must be passed as a string and cannot be embedded in the requests. You can use this to customize the policy definition inherited from the policy family. Policy rules specified here are merged into the inherited policy definition.
* `policy_family_id` (Optional) ID of the policy family. The cluster policy's policy definition inherits the policy family's policy definition. Cannot be used with `definition`. Use `policy_family_definition_overrides` instead to customize the policy definition.

### libraries Configuration Block (Optional)

One must specify each library in a separate configuration block, that will be installed on the cluster that uses a given cluster policy. See [databricks_cluster](cluster.md#library-configuration-block) for more details about supported library types.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the cluster policy. This is equal to `policy_id`.
* `policy_id` - Canonical unique identifier for the cluster policy.

## Import

The resource cluster policy can be imported using the policy id:

```bash
terraform import databricks_cluster_policy.this <cluster-policy-id>
```

## Related Resources

The following resources are often used in the same context:

* [Dynamic Passthrough Clusters for a Group](../guides/workspace-management.md) guide.
* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_clusters](../data-sources/clusters.md) data to retrieve a list of [databricks_cluster](cluster.md) ids.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_current_user](../data-sources/current_user.md) data to retrieve information about [databricks_user](user.md) or [databricks_service_principal](service_principal.md), that is calling Databricks REST API.
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_instance_pool](instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_ip_access_list](ip_access_list.md) to allow access from [predefined IP ranges](https://docs.databricks.com/security/network/ip-access-list.html).
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
* [databricks_node_type](../data-sources/node_type.md) data to get the smallest node type for [databricks_cluster](cluster.md) that fits search criteria, like amount of RAM or number of cores.
* [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_spark_version](../data-sources/spark_version.md) data to get [Databricks Runtime (DBR)](https://docs.databricks.com/runtime/dbr.html) version that could be used for `spark_version` parameter in [databricks_cluster](cluster.md) and other resources.
* [databricks_user_instance_profile](user_instance_profile.md) to attach [databricks_instance_profile](instance_profile.md) (AWS) to [databricks_user](user.md).
* [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
