provider "databricks" {
  azure_workspace_resource_id = azurerm_databricks_workspace.this.id
}

resource "databricks_secret_scope" "terraform" {
  name                     = "terraform"
  initial_manage_principal = "users"
}

resource "databricks_secret" "storage_key" {
  key          = "blob_storage_key"
  string_value = azurerm_storage_account.blobaccount.primary_access_key
  scope        = databricks_secret_scope.terraform.name
}

resource "databricks_azure_blob_mount" "sample" {
  container_name       = azurerm_storage_container.sample.name
  storage_account_name = azurerm_storage_account.blobaccount.name
  mount_name           = "sample"
  auth_type            = "ACCESS_KEY"
  token_secret_scope   = databricks_secret_scope.terraform.name
  token_secret_key     = databricks_secret.storage_key.key
}

resource "databricks_instance_pool" "this" {
  instance_pool_name                    = "Sample pool"
  min_idle_instances                    = 0
  max_capacity                          = 10
  node_type_id                          = "Standard_DS3_v2"
  idle_instance_autotermination_minutes = 10
  preloaded_spark_versions              = ["6.6.x-scala2.11"]

  custom_tags = {
    Team    = "Contributors"
    Aything = "Works"
  }
}

resource "databricks_cluster_policy" "fair_use" {
  name = "Fair Use Policy"
  definition = jsonencode({
    "dbus_per_hour" : {
      "type" : "range",
      "maxValue" : 20
    },
    "autotermination_minutes" : {
      "type" : "fixed",
      "value" : 10,
      "hidden" : true
    },
    "custom_tags.Department" : {
      "type" : "fixed",
      "value" : "Marketing",
      "hidden" : true
    }
  })
}

resource "databricks_permissions" "use_policy" {
    cluster_policy_id = databricks_cluster_policy.fair_use.id

    access_control {
        group_name = databricks_group.marketing.display_name
        permission_level = "CAN_USE"
    }
}

resource "databricks_dbfs_file" "show_variables" {
  source          = "${path.module}/init.sh"
  content_b64_md5 = md5(filebase64("${path.module}/init.sh"))
  path            = "/init-scripts/show-variables.sh"
  overwrite       = true
  mkdirs          = true
}

resource "databricks_cluster" "shared_autoscaling" {
  cluster_name            = "Shared Autoscaling"
  instance_pool_id        = databricks_instance_pool.this.id
  spark_version           = "6.6.x-scala2.11"
  policy_id               = databricks_cluster_policy.fair_use.id
  autotermination_minutes = 10

  autoscale {
    min_workers = 1
    max_workers = 10
  }

  spark_conf = {
    "spark.databricks.delta.preview.enabled": true,
    "spark.databricks.io.cache.enabled" : true,
    "spark.databricks.io.cache.maxDiskUsage" : "50g",
    "spark.databricks.io.cache.maxMetaDataCache" : "1g"
  }

  library {
    pypi {
      package = "fbprophet==0.6"
    }
  }

  library {
    maven {
      coordinates = "com.amazon.deequ:deequ:1.0.4"
    }
  }

  cluster_log_conf {
    dbfs {
      destination = "dbfs:/mnt/sample/cluster-logs"
    }
  }

  init_scripts {
    dbfs {
      destination = databricks_dbfs_file.show_variables.path
    }
  }

  custom_tags = {
    Department = "Marketing"
    Order      = "New"
  }
}

resource "databricks_permissions" "shared_autoscaling" {
    cluster_id = databricks_cluster.shared_autoscaling.id

    access_control {
        group_name = databricks_group.marketing.display_name
        permission_level = "CAN_ATTACH_TO"
    }

    access_control {
        group_name = databricks_group.support.display_name
        permission_level = "CAN_ATTACH_TO"
    }
}

resource "databricks_group" "marketing" {
  display_name = "Marketing"
}

resource "databricks_user" "first" {
  user_name    = "marketing-member@example.com"
  display_name = "Marketing Member"
  default_roles = []
  set_admin    = false
}

resource "databricks_group_member" "first" {
  group_id  = databricks_group.marketing.id
  member_id = databricks_user.first.id
}

resource "databricks_group" "support" {
  display_name = "Support"
}

resource "databricks_user" "second" {
  user_name    = "support-member@example.com"
  display_name = "Support Member"
  default_roles = []
  set_admin    = false
}

resource "databricks_group_member" "second" {
  group_id  = databricks_group.support.id
  member_id = databricks_user.second.id
}

resource "databricks_notebook" "simple" {
  content = filebase64("${path.module}/notebook.py")
  path = "/Production/Something"
  overwrite = true
  mkdirs = true
  language = "PYTHON"
}

resource "databricks_job" "featurization" {
  name                = "Featurization"
  timeout_seconds     = 3600
  max_retries         = 1
  max_concurrent_runs = 1

  new_cluster {
    num_workers      = 3
    instance_pool_id = databricks_instance_pool.this.id
    spark_version    = "6.6.x-scala2.11"

    cluster_log_conf {
      dbfs {
        destination = "dbfs:/mnt/sample/cluster-logs"
      }
    }

    init_scripts {
      dbfs {
        destination = databricks_dbfs_file.show_variables.path
      }
    }
  }

  notebook_task {
    notebook_path = databricks_notebook.simple.path
    base_parameters = {
      "department": "Set From Terraform"
    }
  }

  library {
    pypi {
      package = "fbprophet==0.6"
    }
  }

  email_notifications {
    no_alert_for_skipped_runs = true
  }
}

resource "databricks_permissions" "featurization" {
    job_id = databricks_job.featurization.id

    access_control {
        group_name = databricks_group.support.display_name
        permission_level = "CAN_MANAGE"
    }
}

resource "databricks_secret_acl" "support_reads" {
    principal = databricks_group.support.display_name
    scope = databricks_secret_scope.terraform.name
    permission = "READ"
}
