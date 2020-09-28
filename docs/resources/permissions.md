# databricks_permissions Resource

-> **Note** This feature is not available to all customers. Please contact [sales@databricks.com](mailto:sales@databricks.com) in order to enable this feature. This resource has evolving API, which may change in future versions of provider.

This resource allows you to generically manage permissions for other resources in Databricks workspace. 

## Example Usage

```hcl
resource "databricks_group" "datascience" {
    display_name = "Data scientists"
    allow_cluster_create = false
    allow_instance_pool_create = false
}

resource "databricks_cluster_policy" "something_simple" {
    name = "Some simple policy"
    definition = jsonencode({
        "spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL": {
            "type": "forbidden"
        },
        "spark_conf.spark.secondkey": {
            "type": "forbidden"
        }
    })
}

resource "databricks_permissions" "grant policy usage" {
    cluster_policy_id = databricks_cluster_policy.something_simple.id

    access_control {
        group_name = databricks_scim_group.datascience.display_name
        permission_level = "CAN_USE"
    }
}
```

## Argument Reference

Exactly one of the following attribues is required:

* `cluster_id` - [cluster](cluster.md) id
* `job_id` - [job](job.md) id
* `directory_id` - [directory](notebook.md) id
* `directory_path` - path of directory
* `notebook_id` - ID of [notebook](notebook.md) within workspace
* `notebook_path` - path of notebook
* `cluster_policy_id` - [cluster policy](cluster_policy.md) id
* `instance_pool_id` - [instance pool](instance_pool.md) id
* `authorization` - either [`tokens`](https://docs.databricks.com/administration-guide/access-control/tokens.html) or [`passwords`](https://docs.databricks.com/administration-guide/users-groups/single-sign-on/index.html#configure-password-permission).

One or more `access_control` blocks are required to actually set the permission levels:

```hcl
access_control {
    group_name = databricks_scim_group.datascience.display_name
    permission_level = "CAN_USE"
}
```

Attributes are:

* `permission_level` - (Required) (String) permission level according to [specific resource](https://docs.databricks.com/security/access-control/workspace-acl.html) 
* `user_name` - (Optional) (String) name of the user, which should be used if group name is not used
* `group_name` - (Optional) (String) name of the group, which should be used if user name is not used. We recommend setting permissions on groups.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the permissions.
* `object_type` - (String) type of permissions.