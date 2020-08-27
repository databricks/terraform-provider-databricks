# databricks_cluster_policy Resource

This resource creates a cluster policy, which limits the ability to create clusters based on a set of rules. The policy rules limit the attributes or attribute values available for cluster creation. Cluster policies have ACLs that limit their use to specific users and groups. Only admin users can create, edit, and delete policies. Admin users also have access to all policies.

Cluster policies let you:

* Limit users to create clusters with prescribed settings.
* Simplify the user interface and enable more users to create their own clusters (by fixing and hiding some values).
* Control cost by limiting per cluster maximum cost (by setting limits on attributes whose values contribute to hourly price).

Cluster policy permissions limit which policies a user can select in the Policy drop-down when the user creates a cluster:

* If no policies have been created in the workspace, the Policy drop-down does not display.
* A user who has cluster create permission can select the Free form policy and create fully-configurable clusters.
* A user who has both cluster create permission and access to cluster policies can select the Free form policy and policies they have access to.
* A user that has access to only cluster policies, can select the policies they have access to.

## Example Usage

Let us take a look at an example of how you can manage two teams: Marketing and Data Engineering. In the following scenario we want the marketing team to have a really good query experience, so we enabled delta cache for them. On the other hand we want the data engineering team to be able to utilize bigger clusters so we increased the dbus per hour that they can spend. This strategy allows your marketing users and data engineering users to use Databricks in a self service manner but have a different experience in regards to security and performance. And down the line if you need to add more global settings you can propagate them through the “base cluster policy”.

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
        "autotermination_minutes": {
            "type": "fixed",
            "value": 20,
            "hidden": true
        },
        "custom_tags.Team" : {
            "type" : "fixed",
            "value" : var.team
        }
    }
}

resource "databricks_cluster_policy" "fair_use" {
    name = "${var.team} cluster policy"
    definition = jsonencode(merge(local.default_policy, var.policy_overrides))
}

resource "databricks_permissions" "can_use_cluster_policyinstance_profile" {
    cluster_policy_id = databricks_cluster_policy.fair_use.id
    access_control {
        group_name       = var.team
        permission_level = "CAN_USE"
    }
}
```

And custom instances of that base policy module for our marketing and data engineering teams woud look like:

```hcl
module "marketing_compute_policy" {
    source = "../modules/databricks-cluster-policy"
    team = "marketing"
    policy_overrides = {
        // only marketing guys will benefit from delta cache this way
        "spark_conf.spark.databricks.io.cache.enabled": {
            "value": "true"
        },
    }
}

module "engineering_compute_policy" {
    source = "../modules/databricks-cluster-policy"
    team = "engineering"
    policy_overrides = {
        "dbus_per_hour" : {
            "type" : "range",
            // only engineering guys can spin up big clusters
            "maxValue" : 50
        },
    }
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Cluster policy name. This must be unique. Length must be between 1 and 100 characters.
* `definition` - (Required) Policy definition JSON document expressed in [Databricks Policy Definition Language](https://docs.databricks.com/administration-guide/clusters/policies.html#cluster-policy-definition).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the cluster policy. This equal to policy_id.
* `policy_id` - Canonical unique identifier for the cluster policy.

## Import

The resource cluster policy can be imported using the policy id:

```bash
$ terraform import databricks_cluster_policy.this <cluster-policy-id>
```