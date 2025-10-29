---
subcategory: "Security"
---

# databricks_permission Resource

This resource allows you to manage permissions for a single principal on a Databricks workspace object. Unlike `databricks_permissions`, which manages all principals' permissions for an object at once, this resource is authoritative for a specific object-principal pair only.

-> This resource can only be used with a workspace-level provider!

~> This resource is _authoritative_ for the specified object-principal pair. Configuring this resource will manage the permission for the specified principal only, without affecting permissions for other principals.

-> Use `databricks_permissions` when you need to manage all permissions for an object in a single resource. Use `databricks_permission` (singular) when you want to manage permissions for individual principals independently.

## Example Usage

### Cluster Permissions

```hcl
resource "databricks_cluster" "shared" {
  cluster_name            = "Shared Analytics"
  spark_version           = data.databricks_spark_version.latest.id
  node_type_id            = data.databricks_node_type.smallest.id
  autotermination_minutes = 20

  autoscale {
    min_workers = 1
    max_workers = 10
  }
}

resource "databricks_group" "data_engineers" {
  display_name = "Data Engineers"
}

# Grant CAN_RESTART permission to a group
resource "databricks_permission" "cluster_de" {
  cluster_id       = databricks_cluster.shared.id
  group_name       = databricks_group.data_engineers.display_name
  permission_level = "CAN_RESTART"
}

# Grant CAN_ATTACH_TO permission to a user
resource "databricks_permission" "cluster_analyst" {
  cluster_id       = databricks_cluster.shared.id
  user_name        = "analyst@company.com"
  permission_level = "CAN_ATTACH_TO"
}
```

### Job Permissions

```hcl
resource "databricks_job" "etl" {
  name = "ETL Pipeline"

  task {
    task_key = "process_data"

    notebook_task {
      notebook_path = "/Shared/ETL"
    }

    new_cluster {
      num_workers   = 2
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }
  }
}

# Grant CAN_MANAGE to a service principal
resource "databricks_permission" "job_sp" {
  job_id                   = databricks_job.etl.id
  service_principal_name   = databricks_service_principal.automation.application_id
  permission_level         = "CAN_MANAGE"
}

# Grant CAN_VIEW to a group
resource "databricks_permission" "job_viewers" {
  job_id           = databricks_job.etl.id
  group_name       = "Data Viewers"
  permission_level = "CAN_VIEW"
}
```

### Notebook Permissions

```hcl
resource "databricks_notebook" "analysis" {
  path     = "/Shared/Analysis"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    # Analysis Notebook
    print("Hello, World!")
  EOT
  )
}

# Grant CAN_RUN to a user
resource "databricks_permission" "notebook_user" {
  notebook_path    = databricks_notebook.analysis.path
  user_name        = "data.scientist@company.com"
  permission_level = "CAN_RUN"
}

# Grant CAN_EDIT to a group
resource "databricks_permission" "notebook_editors" {
  notebook_path    = databricks_notebook.analysis.path
  group_name       = "Notebook Editors"
  permission_level = "CAN_EDIT"
}
```

### Token Permissions

This resource solves the limitation where all token permissions must be defined in a single `databricks_permissions` resource:

```hcl
# Multiple resources can now manage different principals independently
resource "databricks_permission" "tokens_team_a" {
  authorization    = "tokens"
  group_name       = "Team A"
  permission_level = "CAN_USE"
}

resource "databricks_permission" "tokens_team_b" {
  authorization    = "tokens"
  group_name       = "Team B"
  permission_level = "CAN_USE"
}

resource "databricks_permission" "tokens_service_account" {
  authorization          = "tokens"
  service_principal_name = databricks_service_principal.ci_cd.application_id
  permission_level       = "CAN_USE"
}
```

### SQL Endpoint Permissions

```hcl
resource "databricks_sql_endpoint" "analytics" {
  name             = "Analytics Warehouse"
  cluster_size     = "Small"
  max_num_clusters = 1
}

resource "databricks_permission" "warehouse_users" {
  sql_endpoint_id  = databricks_sql_endpoint.analytics.id
  group_name       = "SQL Users"
  permission_level = "CAN_USE"
}
```

## Argument Reference

The following arguments are required:

* `permission_level` - (Required) The permission level to grant. The available permission levels depend on the object type. Common values include `CAN_MANAGE`, `CAN_USE`, `CAN_VIEW`, `CAN_RUN`, `CAN_EDIT`, `CAN_READ`, `CAN_RESTART`, `CAN_ATTACH_TO`.

Exactly one of the following principal identifiers must be specified:

* `user_name` - (Optional) User email address to grant permissions to. Conflicts with `group_name` and `service_principal_name`.
* `group_name` - (Optional) Group name to grant permissions to. Conflicts with `user_name` and `service_principal_name`.
* `service_principal_name` - (Optional) Application ID of the service principal. Conflicts with `user_name` and `group_name`.

Exactly one of the following object identifiers must be specified:

* `cluster_id` - (Optional) ID of the [databricks_cluster](cluster.md).
* `cluster_policy_id` - (Optional) ID of the [databricks_cluster_policy](cluster_policy.md).
* `instance_pool_id` - (Optional) ID of the [databricks_instance_pool](instance_pool.md).
* `job_id` - (Optional) ID of the [databricks_job](job.md).
* `pipeline_id` - (Optional) ID of the [databricks_pipeline](pipeline.md).
* `notebook_id` - (Optional) ID of the [databricks_notebook](notebook.md). Can be used when the notebook is referenced by ID.
* `notebook_path` - (Optional) Path to the [databricks_notebook](notebook.md).
* `directory_id` - (Optional) ID of the [databricks_directory](directory.md).
* `directory_path` - (Optional) Path to the [databricks_directory](directory.md).
* `workspace_file_id` - (Optional) ID of the [databricks_workspace_file](workspace_file.md).
* `workspace_file_path` - (Optional) Path to the [databricks_workspace_file](workspace_file.md).
* `registered_model_id` - (Optional) ID of the [databricks_mlflow_model](mlflow_model.md).
* `experiment_id` - (Optional) ID of the [databricks_mlflow_experiment](mlflow_experiment.md).
* `sql_dashboard_id` - (Optional) ID of the legacy [databricks_sql_dashboard](sql_dashboard.md).
* `sql_endpoint_id` - (Optional) ID of the [databricks_sql_endpoint](sql_endpoint.md).
* `sql_query_id` - (Optional) ID of the [databricks_query](query.md).
* `sql_alert_id` - (Optional) ID of the [databricks_alert](alert.md).
* `dashboard_id` - (Optional) ID of the [databricks_dashboard](dashboard.md) (Lakeview).
* `repo_id` - (Optional) ID of the [databricks_repo](repo.md).
* `repo_path` - (Optional) Path to the [databricks_repo](repo.md).
* `authorization` - (Optional) Type of authorization. Currently supports `tokens` and `passwords`.
* `serving_endpoint_id` - (Optional) ID of the [databricks_model_serving](model_serving.md) endpoint.
* `vector_search_endpoint_id` - (Optional) ID of the [databricks_vector_search_endpoint](vector_search_endpoint.md).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the permission in the format `/<object_type>/<object_id>/<principal>`.
* `object_type` - The type of object (e.g., `clusters`, `jobs`, `notebooks`).

## Import

Permissions can be imported using the format `/<object_type>/<object_id>/<principal>`. For example:

```bash
terraform import databricks_permission.cluster_analyst /clusters/0123-456789-abc12345/analyst@company.com
```

## Comparison with databricks_permissions

### When to use `databricks_permission` (singular)

* You want to manage permissions for individual principals independently
* Different teams manage permissions for different principals on the same object
* You need to avoid the "all-or-nothing" approach of `databricks_permissions`
* You want to add/remove principals without affecting others
* Special cases like token permissions where multiple independent configurations are needed

### When to use `databricks_permissions` (plural)

* You want to manage all permissions for an object in one place
* You have a small, stable set of principals for an object
* You want to ensure no unexpected permissions exist on the object
* You prefer a single source of truth for all permissions on an object

### Example Comparison

**Using `databricks_permissions` (manages ALL principals)**:

```hcl
resource "databricks_permissions" "cluster_all" {
  cluster_id = databricks_cluster.shared.id

  access_control {
    group_name       = "Data Engineers"
    permission_level = "CAN_RESTART"
  }

  access_control {
    user_name        = "analyst@company.com"
    permission_level = "CAN_ATTACH_TO"
  }

  # Adding a third principal requires modifying this resource
}
```

**Using `databricks_permission` (manages ONE principal per resource)**:

```hcl
resource "databricks_permission" "cluster_de" {
  cluster_id       = databricks_cluster.shared.id
  group_name       = "Data Engineers"
  permission_level = "CAN_RESTART"
}

resource "databricks_permission" "cluster_analyst" {
  cluster_id       = databricks_cluster.shared.id
  user_name        = "analyst@company.com"
  permission_level = "CAN_ATTACH_TO"
}

# Adding a third principal is a separate resource
# No need to modify existing resources
resource "databricks_permission" "cluster_viewer" {
  cluster_id       = databricks_cluster.shared.id
  group_name       = "Viewers"
  permission_level = "CAN_ATTACH_TO"
}
```

## Related Resources

The following resources are used in the same context:

* [databricks_permissions](permissions.md) - Manage all permissions for an object at once
* [databricks_cluster](cluster.md) - Create Databricks clusters
* [databricks_job](job.md) - Create Databricks jobs
* [databricks_notebook](notebook.md) - Manage Databricks notebooks
* [databricks_group](group.md) - Manage Databricks groups
* [databricks_user](user.md) - Manage Databricks users
* [databricks_service_principal](service_principal.md) - Manage Databricks service principals
