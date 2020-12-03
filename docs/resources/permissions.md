# databricks_permissions Resource

This resource allows you to generically manage permissions for other resources in Databricks workspace. It would guarantee, that only *admins*, *authenticated principal* and those declared within `access_control` blocks would have specified access. It is not possible to remove management rights from *admins* group.

## Cluster usage

It's possible to separate [cluster access control](https://docs.databricks.com/security/access-control/cluster-acl.html) to three different permission levels: `CAN_ATTACH_TO`, `CAN_RESTART` and `CAN_MANAGE`:

```hcl
resource "databricks_group" "auto" {
    display_name = "Automation"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
}

resource "databricks_group" "ds" {
    display_name = "Data Science"
}

resource "databricks_cluster" "shared_autoscaling" {
    cluster_name            = "Shared Autoscaling"
    spark_version           = "6.6.x-scala2.11"
    node_type_id            = "Standard_DS3_v2"
    autotermination_minutes = 60
    autoscale {
        min_workers = 1
        max_workers = 10
    }
}

resource "databricks_permissions" "cluster_usage" {
    cluster_id = databricks_cluster.shared_autoscaling.cluster_id

    access_control {
        group_name = databricks_group.auto.display_name
        permission_level = "CAN_ATTACH_TO"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_RESTART"
    }

    access_control {
        group_name = databricks_group.ds.display_name
        permission_level = "CAN_MANAGE"
    }
}
```

## Cluster Policy usage

Cluster policies allow creation of [clusters](cluster.md), that match [given policy](https://docs.databricks.com/administration-guide/clusters/policies.html). It's possible to assign `CAN_USE` permission to users, service principals, and groups:

```hcl
resource "databricks_group" "ds" {
    display_name = "Data Science"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
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

resource "databricks_permissions" "policy_usage" {
    cluster_policy_id = databricks_cluster_policy.something_simple.id

    access_control {
        group_name = databricks_group.ds.display_name
        permission_level = "CAN_USE"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_USE"
    }
}
```

## Instance Pool usage

[Instance Pools](instance_pool.md) access control [allows to](https://docs.databricks.com/security/access-control/pool-acl.html) assign `CAN_ATTACH_TO` and `CAN_MANAGE` permissions to users, service principals, and groups. It's also possible to grant creation of Instance Pools to individual [groups](group.md#allow_instance_pool_create) and [users](user.md#allow_instance_pool_create), [service principals](service_principal.md#allow_instance_pool_create).

```hcl
resource "databricks_group" "auto" {
    display_name = "Automation"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
}

resource "databricks_instance_pool" "this" {
    instance_pool_name = "Reserved Instances"
    idle_instance_autotermination_minutes = 60
    node_type_id = "i3.xlarge"
    min_idle_instances = 0
    max_capacity = 10
}

resource "databricks_permissions" "pool_usage" {
    instance_pool_id = databricks_instance_pool.this.id

    access_control {
        group_name = databricks_group.auto.display_name
        permission_level = "CAN_ATTACH_TO"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_MANAGE"
    }
}
```

## Job usage

There are four assignable [permission levels](https://docs.databricks.com/security/access-control/jobs-acl.html#job-permissions) for [databricks_job](job.md): `CAN_VIEW`, `CAN_MANAGE_RUN`, `IS_OWNER`, and `CAN_MANAGE`. Admins are granted the `CAN_MANAGE` permission by default, and they can assign that permission to non-admin users, and service principals.

* The creator of a job has `IS_OWNER` permission. Destroying `databricks_permissions` resource for a job would revert ownership to the creator.
* A job must have exactly one owner. If resource is changed and no owner is specified, currently authenticated principal would become new owner of the job. Nothing would change, per se, if the job was created through Terraform.
* A job cannot have a group as an owner.
* Jobs triggered through *Run Now* assume the permissions of the job owner and not the user, and service principal who issued Run Now.
* Read [main documentation](https://docs.databricks.com/security/access-control/jobs-acl.html) for additional detail.

```hcl
resource "databricks_group" "auto" {
    display_name = "Automation"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
}

resource "databricks_job" "this" {
    name = "Featurization"
    max_concurrent_runs = 1
    email_notifications {}

    new_cluster  {
        num_workers   = 300
        spark_version = "6.6.x-scala2.11"
        node_type_id  = "Standard_DS3_v2"
    }
    
    notebook_task {
        notebook_path = "/Production/MakeFeatures"
    }
}

resource "databricks_permissions" "job_usage" {
    job_id = databricks_job.this.id

    access_control {
        group_name = "users"
        permission_level = "CAN_VIEW"
    }

    access_control {
        group_name = databricks_group.auto.display_name
        permission_level = "CAN_MANAGE_RUN"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_MANAGE"
    }
}
```

## Notebook usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html#notebook-permissions) for [databricks_notebook](notebook.md) are: `CAN_READ`, `CAN_RUN`, `CAN_EDIT`, and `CAN_MANAGE`.

```hcl
resource "databricks_group" "auto" {
    display_name = "Automation"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
}

resource "databricks_notebook" "this" {
    content = base64encode("# Welcome to your Python notebook")
    path = "/Production/ETL/Features"
    overwrite = true
    mkdirs = true
    language = "PYTHON"
    format = "SOURCE"
}

resource "databricks_permissions" "notebook_usage" {
    notebook_path = databricks_notebook.this.path

    access_control {
        group_name = "users"
        permission_level = "CAN_READ"
    }

    access_control {
        group_name = databricks_group.auto.display_name
        permission_level = "CAN_RUN"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_EDIT"
    }
}
```

## Folder usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html#folder-permissions) for folders of [databricks_notebook](notebook.md) are: `CAN_READ`, `CAN_RUN`, `CAN_EDIT`, and `CAN_MANAGE`. Notebooks and experiments in a folder inherit all permissions settings of that folder. For example, a user (or service principal) that has `CAN_RUN` permission on a folder has `CAN_RUN` permission on the notebooks in that folder.

* All users can list items in the folder without any permissions.
* All users (or service principals) have `CAN_MANAGE` permission for items in the Workspace > Shared Icon Shared folder. You can grant `CAN_MANAGE` permission to notebooks and folders by moving them to the Shared Icon Shared folder.
* All users (or service principals) have `CAN_MANAGE` permission for objects the user creates.
* User home directory - The user (or service principal) has `CAN_MANAGE` permission. All other users (or service principals) can list their directories.

```hcl
resource "databricks_group" "auto" {
    display_name = "Automation"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
}

resource "databricks_notebook" "this" {
    content = base64encode("# Welcome to your Python notebook")
    path = "/Production/ETL/Features"
    overwrite = true
    mkdirs = true
    language = "PYTHON"
    format = "SOURCE"
}

resource "databricks_permissions" "folder_usage" {
    directory_path = "/Production/ETL"
    depends_on = [databricks_notebook.this]

    access_control {
        group_name = "users"
        permission_level = "CAN_READ"
    }

    access_control {
        group_name = databricks_group.auto.display_name
        permission_level = "CAN_RUN"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_EDIT"
    }
}
```

## Passwords usage

By default on AWS deployments, all admin users (or service principals) can sign in to Databricks using either SSO or their username and password, and all API users (or service principals) can authenticate to the Databricks REST APIs using their username and password. As an admin, you [can limit](https://docs.databricks.com/administration-guide/users-groups/single-sign-on/index.html#optional-configure-password-access-control) admin users’ and API users’ ability to authenticate with their username and password by configuring `CAN_USE` permissions using password access control.
ign-On.

```hcl
resource "databricks_group" "guests" {
    display_name = "Guest Users"
}

resource "databricks_permissions" "password_usage" {
    authorization = "passwords"

    access_control {
        group_name = databricks_group.guests.display_name
        permission_level = "CAN_USE"
    }
}
```

## Token usage

Only [possible permission](https://docs.databricks.com/administration-guide/access-control/tokens.html) to assign to non-admin group is `CAN_USE`, where *admins* `CAN_MANAGE` all tokens:

```hcl
resource "databricks_group" "auto" {
    display_name = "Automation"
}

resource "databricks_group" "eng" {
    display_name = "Engineering"
}

resource "databricks_permissions" "token_usage" {
    authorization = "tokens"

    access_control {
        group_name = databricks_group.auto.display_name
        permission_level = "CAN_USE"
    }

    access_control {
        group_name = databricks_group.eng.display_name
        permission_level = "CAN_USE"
    }
}
```

## Instance Profiles

[Instance Profiles](instance_profile.md) are not managed by General Permissions API and therefore [databricks_group_instance_profile](group_instance_profile.md) and [databricks_user_instance_profile](user_instance_profile.md) should be used to allow usage of specific AWS EC2 IAM roles to users or groups.

## Secrets

One can control access to [databricks_secret](secret.md) through `initial_manage_principal` argument on [databricks_secret_scope](secret_scope.md) or [databricks_secret_acl](secret_acl.md), so that users (or service principals) can `READ`, `WRITE` or `MANAGE` entries within secret scope.

## Tables, Views and Databases

General Permissions API does not apply to access control for tables and permissions have to be managed separately. Though, terraform integration is coming in the future versions.

## Argument Reference

Exactly one of the following attributes is required:

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

* `permission_level` - (Required) permission level according to specific resource. See examples above for the reference.
* `user_name` - (Optional) name of the [user](user.md), which should be used if group name is not used
* `group_name` - (Optional) name of the [group](group.md), which should be used if the user name is not used. We recommend setting permissions on groups.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the permissions.
* `object_type` - type of permissions.

## Import

The resource permissions can be imported using the object id

```bash
$ terraform import databricks_permissions.this /<object type>/<object id>
```
