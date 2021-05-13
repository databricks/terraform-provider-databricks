---
page_title: "End to end workspace management"
---

# End to end workspace management

Once you have the workspace setup on [Azure](azure-workspace.md) or [AWS](aws-workspace.md), you have to start managing resources within your workspace. The following configuration blocks initializes ths most common variables, [databricks_spark_version](../data-sources/spark_version.md), [databricks_node_type](../data-sources/node_type.md), and [databricks_current_user](../data-sources/current_user.md).

```hcl
terraform {
  required_providers {
    databricks = {
      source  = "databrickslabs/databricks"
      version = "0.3.4"
    }
  }
}

provider "databricks" {}

data "databricks_current_user" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}
```

## Part 1: Standard functionality

These resources do not require administrative priviliges. More documentation is available at dedicated pages for [databricks_secret_scope](../resources/secret_scope.md), [databricks_token](../resources/token.md), [databricks_secret](../resources/secret.md), [databricks_notebook](../data-sources/notebook.md), [databricks_job](../resources/job.md), [databricks_cluster](../resources/cluster.md), [databricks_cluster_policy](../resources/cluster_policy.md), [databricks_instance_pool](../resources/instance_pool.md).

```hcl
resource "databricks_secret_scope" "this" {
  name = "demo-${data.databricks_current_user.me.alphanumeric}"
}

resource "databricks_token" "pat" {
  comment          = "Created from ${abspath(path.module)}"
  lifetime_seconds = 3600
}

resource "databricks_secret" "token" {
  string_value = databricks_token.pat.token_value
  scope        = databricks_secret_scope.this.name
  key          = "token"
}

resource "databricks_notebook" "this" {
  path     = "${data.databricks_current_user.me.home}/Terraform"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    token = dbutils.secrets.get('${databricks_secret_scope.this.name}', '${databricks_secret.token.key}')
    print(f'This should be recacted: {token}')
    EOT
  )
}

resource "databricks_job" "this" {
  name = "Terraform Demo (${data.databricks_current_user.me.alphanumeric})"

  new_cluster {
    num_workers   = 1
    spark_version = data.databricks_spark_version.latest.id
    node_type_id  = data.databricks_node_type.smallest.id
  }

  notebook_task {
    notebook_path = databricks_notebook.this.path
  }

  email_notifications {}
}

resource "databricks_cluster" "this" {
  cluster_name = "Exploration (${data.databricks_current_user.me.alphanumeric})"
  spark_version           = data.databricks_spark_version.latest.id
  instance_pool_id        = databricks_instance_pool.smallest_nodes.id
  autotermination_minutes = 20
  autoscale {
    min_workers = 1
    max_workers = 10
  }
}

resource "databricks_cluster_policy" "this" {
  name = "Minimal (${data.databricks_current_user.me.alphanumeric})"
  definition = jsonencode({
    "dbus_per_hour" : {
      "type" : "range",
      "maxValue" : 10
    },
    "autotermination_minutes" : {
      "type" : "fixed",
      "value" : 20,
      "hidden" : true
    }
  })
}

resource "databricks_instance_pool" "smallest_nodes" {
  instance_pool_name = "Smallest Nodes (${data.databricks_current_user.me.alphanumeric})"
  min_idle_instances = 0
  max_capacity       = 30
  node_type_id       = data.databricks_node_type.smallest.id
  preloaded_spark_versions = [
    data.databricks_spark_version.latest.id
  ]

  idle_instance_autotermination_minutes = 20
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

## Part 2: Workspace security

Managing security requires administrative priviliges. More documentation is available at dedicated pages for [databricks_secret_acl](../resources/secret_acl.md), [databricks_group](../data-sources/group.md), [databricks_user](../resources/user.md), [databricks_group_member](../resources/group_member.md), [databricks_permissions](../resources/permissions.md).

```hcl
resource "databricks_secret_acl" "spectators" {
  principal  = databricks_group.spectators.display_name
  scope      = databricks_secret_scope.this.name
  permission = "READ"
}

resource "databricks_group" "spectators" {
  display_name = "Spectators (by ${data.databricks_current_user.me.alphanumeric})"
}

resource "databricks_user" "dummy" {
  user_name    = "dummy+${data.databricks_current_user.me.alphanumeric}@example.com"
  display_name = "Dummy ${data.databricks_current_user.me.alphanumeric}"
}

resource "databricks_group_member" "a" {
  group_id  = databricks_group.spectators.id
  member_id = databricks_user.dummy.id
}

resource "databricks_permissions" "notebook" {
  notebook_path = databricks_notebook.this.id
  access_control {
    user_name        = databricks_user.dummy.user_name
    permission_level = "CAN_RUN"
  }
  access_control {
    group_name       = databricks_group.spectators.display_name
    permission_level = "CAN_READ"
  }
}

resource "databricks_permissions" "job" {
  job_id = databricks_job.this.id
  access_control {
    user_name        = databricks_user.dummy.user_name
    permission_level = "IS_OWNER"
  }
  access_control {
    group_name       = databricks_group.spectators.display_name
    permission_level = "CAN_MANAGE_RUN"
  }
}

resource "databricks_permissions" "cluster" {
  cluster_id = databricks_cluster.this.id
  access_control {
    user_name        = databricks_user.dummy.user_name
    permission_level = "CAN_RESTART"
  }
  access_control {
    group_name       = databricks_group.spectators.display_name
    permission_level = "CAN_ATTACH_TO"
  }
}

resource "databricks_permissions" "policy" {
  cluster_policy_id = databricks_cluster_policy.this.id
  access_control {
    group_name       = databricks_group.spectators.display_name
    permission_level = "CAN_USE"
  }
}

resource "databricks_permissions" "pool" {
  instance_pool_id = databricks_instance_pool.smallest_nodes.id
  access_control {
    group_name       = databricks_group.spectators.display_name
    permission_level = "CAN_ATTACH_TO"
  }
}
```

## Part 3: Storage

Depending on your preferences and needs, you can

* Manage JAR, Wheel & Egg libraries through [databricks_dbfs_file](../resources/dbfs_file.md).
* List entries on DBFS with [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data source.
* Get contents of small files with [databricks_dbfs_file](../data-sources/dbfs_file.md) data source.
* Mount your AWS storage using [databricks_aws_s3_mount](../resources/aws_s3_mount.md).
* Mount your Azure storage using [databricks_azure_adls_gen1_mount](../resources/azure_adls_gen1_mount.md), [databricks_azure_adls_gen2_mount](../resources/azure_adls_gen2_mount.md), [databricks_azure_blob_mount](../resources/azure_blob_mount.md).

## Part 4: Advanced configuration

More documentation is available at dedicated pages for [databricks_workspace_conf](../resources/workspace_conf.md) and [databricks_ip_access_list](../resources/ip_access_list.md).

```hcl
data "http" "my" {
  url = "https://ifconfig.me"
}

resource "databricks_workspace_conf" "this" {
  custom_config = {
    "enableIpAccessLists": true
  }
}

resource "databricks_ip_access_list" "only_me" {
  label = "only ${data.http.my.body} is allowed to access workspace"
  list_type = "ALLOW"
  ip_addresses = ["${data.http.my.body}/32"]
  depends_on = [databricks_workspace_conf.this]
}
```
