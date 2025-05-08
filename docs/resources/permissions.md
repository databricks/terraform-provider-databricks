---
subcategory: "Security"
---

# databricks_permissions Resource

This resource allows you to generically manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspaces. It ensures that only _admins_, _authenticated principal_ and those declared within `access_control` blocks would have specified access. It is not possible to remove management rights from _admins_ group.

-> This resource can only be used with a workspace-level provider!

~> This resource is _authoritative_ for permissions on objects. Configuring this resource for an object will **OVERWRITE** any existing permissions of the same type unless imported, and changes made outside of Terraform will be reset.

-> It is not possible to lower permissions for `admins`, so Databricks Terraform Provider removes those `access_control` blocks automatically.

-> If multiple permission levels are specified for an identity (e.g. `CAN_RESTART` and `CAN_MANAGE` for a cluster), only the highest level permission is returned and will cause permanent drift.

~> To manage access control on service principals, use [databricks_access_control_rule_set](access_control_rule_set.md).

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

data "databricks_spark_version" "latest" {}

data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 60
  autoscale {
    min_workers = 1
    max_workers = 10
  }
}

resource "databricks_permissions" "cluster_usage" {
  cluster_id = databricks_cluster.shared_autoscaling.id

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_ATTACH_TO"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_RESTART"
  }

  access_control {
    group_name       = databricks_group.ds.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Cluster Policy usage

Cluster policies allow creation of [clusters](cluster.md), that match [given policy](https://docs.databricks.com/administration-guide/clusters/policies.html). It's possible to assign `CAN_USE` permission to users and groups:

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
    "spark_conf.spark.hadoop.javax.jdo.option.ConnectionURL" : {
      "type" : "forbidden"
    },
    "spark_conf.spark.secondkey" : {
      "type" : "forbidden"
    }
  })
}

resource "databricks_permissions" "policy_usage" {
  cluster_policy_id = databricks_cluster_policy.something_simple.id

  access_control {
    group_name       = databricks_group.ds.display_name
    permission_level = "CAN_USE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
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

data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_instance_pool" "this" {
  instance_pool_name                    = "Reserved Instances"
  idle_instance_autotermination_minutes = 60
  node_type_id                          = data.databricks_node_type.smallest.id
  min_idle_instances                    = 0
  max_capacity                          = 10
}

resource "databricks_permissions" "pool_usage" {
  instance_pool_id = databricks_instance_pool.this.id

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_ATTACH_TO"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Job usage

There are four assignable [permission levels](https://docs.databricks.com/security/access-control/jobs-acl.html#job-permissions) for [databricks_job](job.md): `CAN_VIEW`, `CAN_MANAGE_RUN`, `IS_OWNER`, and `CAN_MANAGE`. Admins are granted the `CAN_MANAGE` permission by default, and they can assign that permission to non-admin users, and service principals.

- The creator of a job has `IS_OWNER` permission. Destroying `databricks_permissions` resource for a job would revert ownership to the creator.
- A job must have exactly one owner. If a resource is changed and no owner is specified, the currently authenticated principal would become the new owner of the job. Nothing would change, per se, if the job was created through Terraform.
- A job cannot have a group as an owner.
- Jobs triggered through _Run Now_ assume the permissions of the job owner and not the user, and service principal who issued Run Now.
- Read [main documentation](https://docs.databricks.com/security/access-control/jobs-acl.html) for additional detail.

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_service_principal" "aws_principal" {
  display_name = "main"
}

data "databricks_spark_version" "latest" {}

data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_job" "this" {
  name                = "Featurization"
  max_concurrent_runs = 1

  task {
    task_key = "task1"

    new_cluster {
      num_workers   = 300
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }

    notebook_task {
      notebook_path = "/Production/MakeFeatures"
    }
  }
}

resource "databricks_permissions" "job_usage" {
  job_id = databricks_job.this.id

  access_control {
    group_name       = "users"
    permission_level = "CAN_VIEW"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_MANAGE_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }

  access_control {
    service_principal_name = databricks_service_principal.aws_principal.application_id
    permission_level       = "IS_OWNER"
  }
}
```

## Delta Live Tables usage

There are four assignable [permission levels](https://docs.databricks.com/security/access-control/dlt-acl.html#delta-live-tables-permissions) for [databricks_pipeline](pipeline.md): `CAN_VIEW`, `CAN_RUN`, `CAN_MANAGE`, and `IS_OWNER`. Admins are granted the `CAN_MANAGE` permission by default, and they can assign that permission to non-admin users, and service principals.

- The creator of a DLT Pipeline has `IS_OWNER` permission. Destroying `databricks_permissions` resource for a pipeline would revert ownership to the creator.
- A DLT pipeline must have exactly one owner. If a resource is changed and no owner is specified, the currently authenticated principal would become the new owner of the pipeline. Nothing would change, per se, if the pipeline was created through Terraform.
- A DLT pipeline cannot have a group as an owner.
- DLT Pipelines triggered through _Start_ assume the permissions of the pipeline owner and not the user, and service principal who issued Run Now.
- Read [main documentation](https://docs.databricks.com/security/access-control/dlt-acl.html) for additional detail.

```hcl
data "databricks_current_user" "me" {}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_notebook" "dlt_demo" {
  content_base64 = base64encode(<<-EOT
    import dlt
    json_path = "/databricks-datasets/wikipedia-datasets/data-001/clickstream/raw-uncompressed-json/2015_2_clickstream.json"
    @dlt.table(
       comment="The raw wikipedia clickstream dataset, ingested from /databricks-datasets."
    )
    def clickstream_raw():
        return (spark.read.format("json").load(json_path))
    EOT
  )
  language = "PYTHON"
  path     = "${data.databricks_current_user.me.home}/DLT_Demo"
}

resource "databricks_pipeline" "this" {
  name    = "DLT Demo Pipeline (${data.databricks_current_user.me.alphanumeric})"
  storage = "/test/tf-pipeline"
  configuration = {
    key1 = "value1"
    key2 = "value2"
  }

  library {
    notebook {
      path = databricks_notebook.dlt_demo.id
    }
  }

  continuous = false
  filters {
    include = ["com.databricks.include"]
    exclude = ["com.databricks.exclude"]
  }
}

resource "databricks_permissions" "dlt_usage" {
  pipeline_id = databricks_pipeline.this.id

  access_control {
    group_name       = "users"
    permission_level = "CAN_VIEW"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Notebook usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html#notebook-permissions) for [databricks_notebook](notebook.md) are: `CAN_READ`, `CAN_RUN`, `CAN_EDIT`, and `CAN_MANAGE`.

A notebook could be specified by using either `notebook_path` or `notebook_id` attribute.  The value for the `notebook_id` is the object ID of the resource in the Databricks Workspace that is exposed as `object_id` attribute of the `databricks_notebook` resource as shown below.

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_notebook" "this" {
  content_base64 = base64encode("# Welcome to your Python notebook")
  path           = "/Production/ETL/Features"
  language       = "PYTHON"
}

resource "databricks_permissions" "notebook_usage_by_path" {
  notebook_path = databricks_notebook.this.path

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}

resource "databricks_permissions" "notebook_usage_by_id" {
  notebook_id = databricks_notebook.this.object_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}
```

-> when importing a permissions resource, only the `notebook_id` is filled!

## Workspace file usage

Valid permission levels for [databricks_workspace_file](workspace_file.md) are: `CAN_READ`, `CAN_RUN`, `CAN_EDIT`, and `CAN_MANAGE`.

A workspace file could be specified by using either `workspace_file_path` or `workspace_file_id` attribute.  The value for the `workspace_file_id` is the object ID of the resource in the Databricks Workspace that is exposed as `object_id` attribute of the `databricks_workspace_file` resource as shown below.

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_workspace_file" "this" {
  content_base64 = base64encode("print('Hello World')")
  path           = "/Production/ETL/Features.py"
}

resource "databricks_permissions" "workspace_file_usage_by_path" {
  workspace_file_path = databricks_workspace_file.this.path

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}

resource "databricks_permissions" "workspace_file_usage_by_id" {
  workspace_file_id = databricks_workspace_file.this.object_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}
```

-> when importing a permissions resource, only the `workspace_file_id` is filled!

## Folder usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html#folder-permissions) for folders of [databricks_directory](directory.md) are: `CAN_READ`, `CAN_RUN`, `CAN_EDIT`, and `CAN_MANAGE`. Notebooks and experiments in a folder inherit all permissions settings of that folder. For example, a user (or service principal) that has `CAN_RUN` permission on a folder has `CAN_RUN` permission on the notebooks in that folder.

- All users can list items in the folder without any permissions.
- All users (or service principals) have `CAN_MANAGE` permission for items in the Workspace > Shared Icon Shared folder. You can grant `CAN_MANAGE` permission to notebooks and folders by moving them to the Shared Icon Shared folder.
- All users (or service principals) have `CAN_MANAGE` permission for objects the user creates.
- User home directory - The user (or service principal) has `CAN_MANAGE` permission. All other users (or service principals) can list their directories.

A folder could be specified by using either `directory_path` or `directory_id` attribute.  The value for the `directory_id` is the object ID of the resource in the Databricks Workspace that is exposed as `object_id` attribute of the `databricks_directory` resource as shown below.

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_directory" "this" {
  path = "/Production/ETL"
}

resource "databricks_permissions" "folder_usage_by_path" {
  directory_path = databricks_directory.this.path

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}

resource "databricks_permissions" "folder_usage_by_id" {
  directory_id = databricks_directory.this.object_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}
```

-> when importing a permissions resource, only the `directory_id` is filled!

## Repos usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html) for [databricks_repo](repo.md) are: `CAN_READ`, `CAN_RUN`, `CAN_EDIT`, and `CAN_MANAGE`.

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_repo" "this" {
  url = "https://github.com/user/demo.git"
}

resource "databricks_permissions" "repo_usage" {
  repo_id = databricks_repo.this.id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}
```

## MLflow Experiment usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html#mlflow-experiment-permissions-1) for [databricks_mlflow_experiment](mlflow_experiment.md) are: `CAN_READ`, `CAN_EDIT`, and `CAN_MANAGE`.

```hcl
data "databricks_current_user" "me" {}

resource "databricks_mlflow_experiment" "this" {
  name              = "${data.databricks_current_user.me.home}/Sample"
  artifact_location = "dbfs:/tmp/my-experiment"
  description       = "My MLflow experiment description"
}

resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "experiment_usage" {
  experiment_id = databricks_mlflow_experiment.this.id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_MANAGE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_EDIT"
  }
}
```

## MLflow Model usage

Valid [permission levels](https://docs.databricks.com/security/access-control/workspace-acl.html#mlflow-model-permissions-1) for [databricks_mlflow_model](mlflow_model.md) are: `CAN_READ`, `CAN_EDIT`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_MANAGE_PRODUCTION_VERSIONS`, and `CAN_MANAGE`. You can also manage permissions for all MLflow models by `registered_model_id = "root"`.

```hcl
resource "databricks_mlflow_model" "this" {
  name = "SomePredictions"
}

resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "model_usage" {
  registered_model_id = databricks_mlflow_model.this.registered_model_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_MANAGE_PRODUCTION_VERSIONS"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE_STAGING_VERSIONS"
  }
}
```

## Model serving usage

Valid permission levels for [databricks_model_serving](model_serving.md) are: `CAN_VIEW`, `CAN_QUERY`, and `CAN_MANAGE`.

```hcl
resource "databricks_model_serving" "this" {
  name = "tf-test"
  config {
    served_models {
      name                  = "prod_model"
      model_name            = "test"
      model_version         = "1"
      workload_size         = "Small"
      scale_to_zero_enabled = true
    }
  }
}

resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "ml_serving_usage" {
  serving_endpoint_id = databricks_model_serving.this.serving_endpoint_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_VIEW"
  }

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_MANAGE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_QUERY"
  }
}
```

## Mosaic AI Vector Search usage

Valid permission levels for [databricks_vector_search_endpoint](vector_search_endpoint.md) are: `CAN_USE` and `CAN_MANAGE`.

-> You need to use the `endpoint_id` attribute of `databricks_vector_search_endpoint` as value for `vector_search_endpoint_id`, not the `id`!

```hcl
resource "databricks_vector_search_endpoint" "this" {
  name          = "vector-search-test"
  endpoint_type = "STANDARD"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "vector_search_endpoint_usage" {
  vector_search_endpoint_id = databricks_vector_search_endpoint.this.endpoint_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_USE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Passwords usage

By default on AWS deployments, all admin users can sign in to Databricks using either SSO or their username and password, and all API users can authenticate to the Databricks REST APIs using their username and password. As an admin, you [can limit](https://docs.databricks.com/administration-guide/users-groups/single-sign-on/index.html#optional-configure-password-access-control) admin users' and API users' ability to authenticate with their username and password by configuring `CAN_USE` permissions using password access control.

```hcl
resource "databricks_group" "guests" {
  display_name = "Guest Users"
}

resource "databricks_permissions" "password_usage" {
  authorization = "passwords"

  access_control {
    group_name       = databricks_group.guests.display_name
    permission_level = "CAN_USE"
  }
}
```

## Token usage

It is required to have at least 1 personal access token in the workspace before you can manage tokens permissions.

!> **Warning** There can be only one `authorization = "tokens"` permissions resource per workspace, otherwise there'll be a permanent configuration drift. After applying changes, users who previously had either `CAN_USE` or `CAN_MANAGE` permission but no longer have either permission have their access to token-based authentication revoked. Their active tokens are immediately deleted (revoked).

Only [possible permission](https://docs.databricks.com/administration-guide/access-control/tokens.html) to assign to non-admin group is `CAN_USE`, where _admins_ `CAN_MANAGE` all tokens:

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
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_USE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_USE"
  }
}
```

## SQL warehouse usage

[SQL warehouses](https://docs.databricks.com/sql/user/security/access-control/sql-endpoint-acl.html) have five possible permissions: `CAN_USE`, `CAN_MONITOR`, `CAN_MANAGE`, `CAN_VIEW` and `IS_OWNER`:

- The creator of a warehouse has `IS_OWNER` permission. Destroying `databricks_permissions` resource for a warehouse would revert ownership to the creator.
- A warehouse must have exactly one owner. If a resource is changed and no owner is specified, the currently authenticated principal would become the new owner of the warehouse. Nothing would change, per se, if the warehouse was created through Terraform.
- A warehouse cannot have a group as an owner.

```hcl
data "databricks_current_user" "me" {}

resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_sql_endpoint" "this" {
  name             = "Endpoint of ${data.databricks_current_user.me.alphanumeric}"
  cluster_size     = "Small"
  max_num_clusters = 1

  tags {
    custom_tags {
      key   = "City"
      value = "Amsterdam"
    }
  }
}

resource "databricks_permissions" "endpoint_usage" {
  sql_endpoint_id = databricks_sql_endpoint.this.id

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_USE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Dashboard usage

[Dashboards](https://docs.databricks.com/en/dashboards/tutorials/manage-permissions.html) have four possible permissions: `CAN_READ`, `CAN_RUN`, `CAN_EDIT` and `CAN_MANAGE`:

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_dashboard" "dashboard" {
  display_name = "TF New Dashboard"
  # ...
}


resource "databricks_permissions" "dashboard_usage" {
  dashboard_id = databricks_dashboard.dashboard.id

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Legacy SQL Dashboard usage

[Legacy SQL dashboards](https://docs.databricks.com/sql/user/security/access-control/dashboard-acl.html) have three possible permissions: `CAN_VIEW`, `CAN_RUN` and `CAN_MANAGE`:

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "sql_dashboard_usage" {
  sql_dashboard_id = "3244325"

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## SQL Query usage

[SQL queries](https://docs.databricks.com/sql/user/security/access-control/query-acl.html) have three possible permissions: `CAN_VIEW`, `CAN_RUN` and `CAN_MANAGE`:

-> If you do not define an `access_control` block granting `CAN_MANAGE` explictly for the user calling this provider, Databricks Terraform Provider will add `CAN_MANAGE` permission for the caller. This is a failsafe to prevent situations where the caller is locked out from making changes to the targeted `databricks_sql_query` resource when backend API do not apply permission inheritance correctly.

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "query_usage" {
  sql_query_id = "3244325"

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## SQL Alert usage

[SQL alerts](https://docs.databricks.com/sql/user/security/access-control/alert-acl.html) have three possible permissions: `CAN_VIEW`, `CAN_RUN` and `CAN_MANAGE`:

```hcl
resource "databricks_group" "auto" {
  display_name = "Automation"
}

resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "alert_usage" {
  sql_alert_id = "3244325"

  access_control {
    group_name       = databricks_group.auto.display_name
    permission_level = "CAN_RUN"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Databricks Apps usage

[Databricks Apps](https://docs.databricks.com/en/dev-tools/databricks-apps/index.html) have two possible permissions: `CAN_USE` and `CAN_MANAGE`:

```hcl
resource "databricks_group" "eng" {
  display_name = "Engineering"
}

resource "databricks_permissions" "app_usage" {
  app_name = "myapp"

  access_control {
    group_name       = "users"
    permission_level = "CAN_USE"
  }

  access_control {
    group_name       = databricks_group.eng.display_name
    permission_level = "CAN_MANAGE"
  }
}
```

## Instance Profiles

[Instance Profiles](instance_profile.md) are not managed by General Permissions API and therefore [databricks_group_instance_profile](group_instance_profile.md) and [databricks_user_instance_profile](user_instance_profile.md) should be used to allow usage of specific AWS EC2 IAM roles to users or groups.

## Secrets

One can control access to [databricks_secret](secret.md) through `initial_manage_principal` argument on [databricks_secret_scope](secret_scope.md) or [databricks_secret_acl](secret_acl.md), so that users (or service principals) can `READ`, `WRITE` or `MANAGE` entries within secret scope.

## Tables, Views and Databases

General Permissions API does not apply to access control for tables and they have to be managed separately using the [databricks_sql_permissions](sql_permissions.md) resource, though you're encouraged to use Unity Catalog or migrate to it.

## Data Access with Unity Catalog

Initially in Unity Catalog all users have no access to data, which has to be later assigned through [databricks_grants](grants.md) or [databricks_grant](grant.md) resource.

## Argument Reference

One type argument and at least one access control block argument are required.

### Type Argument

Exactly one of the following arguments is required:

- `app_name` - [app](app.md) name
- `cluster_id` - [cluster](cluster.md) id
- `cluster_policy_id` - [cluster policy](cluster_policy.md) id
- `instance_pool_id` - [instance pool](instance_pool.md) id
- `job_id` - [job](job.md) id
- `pipeline_id` - [pipeline](pipeline.md) id
- `notebook_id` - ID of [notebook](notebook.md) within workspace
- `notebook_path` - path of notebook
- `directory_id` - [directory](notebook.md) id
- `directory_path` - path of directory
- `repo_id` - [repo](repo.md) id
- `repo_path` - path of databricks repo directory(`/Repos/<username>/...`)
- `experiment_id` - [MLflow experiment](mlflow_experiment.md) id
- `registered_model_id` - [MLflow registered model](mlflow_model.md) id
- `serving_endpoint_id` - [Model Serving](model_serving.md) endpoint id.
- `vector_search_endpoint_id` - [Vector Search](vector_search_endpoint.md) endpoint id.
- `authorization` - either [`tokens`](https://docs.databricks.com/administration-guide/access-control/tokens.html) or [`passwords`](https://docs.databricks.com/administration-guide/users-groups/single-sign-on/index.html#configure-password-permission).
- `sql_endpoint_id` - [SQL warehouse](sql_endpoint.md) id
- `sql_dashboard_id` - [SQL dashboard](sql_dashboard.md) id
- `sql_query_id` - [SQL query](sql_query.md) id
- `sql_alert_id` - [SQL alert](https://docs.databricks.com/sql/user/security/access-control/alert-acl.html) id

### Access Control Argument

One or more `access_control` blocks are required to actually set the permission levels:

```hcl
access_control {
  group_name       = databricks_group.datascience.display_name
  permission_level = "CAN_USE"
}
```

Arguments for the `access_control` block are:

-> It is not possible to lower permissions for `admins` or your own user anywhere from `CAN_MANAGE` level, so Databricks Terraform Provider [removes](https://github.com/databricks/terraform-provider-databricks/blob/main/permissions/resource_permissions.go#L324-L332) those `access_control` blocks automatically.

- `permission_level` - (Required) permission level according to specific resource. See examples above for the reference.

Exactly one of the below arguments is required:

- `user_name` - (Optional) name of the [user](user.md).
- `service_principal_name` - (Optional) Application ID of the [service_principal](service_principal.md#application_id).
- `group_name` - (Optional) name of the [group](group.md). We recommend setting permissions on groups.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - Canonical unique identifier for the permissions in form of `/<object type>/<object id>`.
- `object_type` - type of permissions.

## Import

The resource permissions can be imported using the object id

```bash
terraform import databricks_permissions.this /<object type>/<object id>
```

### Import Example

Configuration file:

```hcl
resource "databricks_mlflow_model" "model" {
  name        = "example_model"
  description = "MLflow registered model"
}

resource "databricks_permissions" "model_usage" {
  registered_model_id = databricks_mlflow_model.model.registered_model_id

  access_control {
    group_name       = "users"
    permission_level = "CAN_READ"
  }
}
```

Import command:

```bash
terraform import databricks_permissions.model_usage /registered-models/<registered_model_id>
```

```hcl
import {
  to = databricks_permissions.this
  id = "/<object type>/<object id>"
}
```

Alternatively, when using `terraform` version 1.5 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_permissions.this /<object type>/<object id>
```
