---
subcategory: "Compute"
---
# databricks_job Resource

The `databricks_job` resource allows you to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).

## Example Usage

```hcl
data "databricks_current_user" "me" {}
data "databricks_spark_version" "latest" {}
data "databricks_node_type" "smallest" {
  local_disk = true
}

resource "databricks_notebook" "this" {
  path     = "${data.databricks_current_user.me.home}/Terraform"
  language = "PYTHON"
  content_base64 = base64encode(<<-EOT
    # created from ${abspath(path.module)}
    display(spark.range(10))
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
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

## Jobs with Multiple Tasks

-> **Note** In terraform configuration, it is recommended to define tasks in alphabetical order of their `task_key` arguments, so that you get consistent and readable diff. Whenever tasks are added or removed, or `task_key` is renamed, you'll observe a change in the majority of tasks. It's related to the fact that the current version of the provider treats `task` blocks as an ordered list. Alternatively, `task` block could have been an unordered set, though end-users would see the entire block replaced upon a change in single property of the task.

It is possible to create [jobs with multiple tasks](https://docs.databricks.com/data-engineering/jobs/jobs-user-guide.html) using `task` blocks:

```hcl
resource "databricks_job" "this" {
  name = "Job with multiple tasks"

  job_cluster {
    job_cluster_key = "j"
    new_cluster {
      num_workers   = 2
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }
  }

  task {
    task_key = "a"

    new_cluster {
      num_workers   = 1
      spark_version = data.databricks_spark_version.latest.id
      node_type_id  = data.databricks_node_type.smallest.id
    }

    notebook_task {
      notebook_path = databricks_notebook.this.path
    }
  }

  task {
    task_key = "b"
    //this task will only run after task a
    depends_on {
      task_key = "a"
    }

    existing_cluster_id = databricks_cluster.shared.id

    spark_jar_task {
      main_class_name = "com.acme.data.Main"
    }
  }

  task {
    task_key = "c"

    job_cluster_key = "j"

    notebook_task {
      notebook_path = databricks_notebook.this.path
    }
  }
  //this task starts a Delta Live Tables pipline update
  task {
    task_key = "d"

    pipeline_task {
      pipeline_id = databricks_pipeline.this.id
    }
  }
}
```

Every `task` block can have almost all available arguments with the addition of `task_key` attribute and `depends_on` blocks to define cross-task dependencies.

## Argument Reference

The following arguments are required:

* `name` - (Optional) An optional name for the job. The default value is Untitled.
* `new_cluster` - (Optional) Same set of parameters as for [databricks_cluster](cluster.md) resource.
* `existing_cluster_id` - (Optional) If existing_cluster_id, the ID of an existing [cluster](cluster.md) that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually restart the cluster if it stops responding. We strongly suggest to use `new_cluster` for greater reliability.
* `always_running` - (Optional) (Bool) Whenever the job is always running, like a Spark Streaming application, on every update restart the current active run or start it again, if nothing it is not running. False by default. Any job runs are started with `parameters` specified in `spark_jar_task` or `spark_submit_task` or `spark_python_task` or `notebook_task` blocks.
* `library` - (Optional) (Set) An optional list of libraries to be installed on the cluster that will execute the job. Please consult [libraries section](cluster.md#libraries) for [databricks_cluster](cluster.md) resource.
* `retry_on_timeout` - (Optional) (Bool) An optional policy to specify whether to retry a job when it times out. The default behavior is to not retry on timeout.
* `max_retries` - (Optional) (Integer) An optional maximum number of times to retry an unsuccessful run. A run is considered to be unsuccessful if it completes with a FAILED result_state or INTERNAL_ERROR life_cycle_state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior is to never retry.
* `timeout_seconds` - (Optional) (Integer) An optional timeout applied to each run of this job. The default behavior is to have no timeout.
* `min_retry_interval_millis` - (Optional) (Integer) An optional minimal interval in milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that unsuccessful runs are immediately retried.
* `max_concurrent_runs` - (Optional) (Integer) An optional maximum allowed number of concurrent runs of the job. Defaults to *1*.
* `email_notifications` - (Optional) (List) An optional set of email addresses notified when runs of this job begin and complete and when this job is deleted. The default behavior is to not send any emails. This field is a block and is documented below.
* `schedule` - (Optional) (List) An optional periodic schedule for this job. The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API request to runNow. This field is a block and is documented below.
* `tags` - (Optional) (Map) An optional map of the tags associated with the job. Specified tags will be used as cluster tags for job clusters.

### job_cluster Configuration Block
[Shared job cluster](https://docs.databricks.com/jobs.html#use-shared-job-clusters) specification. Allows multiple tasks in the same job run to reuse the cluster. 
* `job_cluster_key` - (Required) Identifier that can be referenced in `task` block, so that cluster is shared between tasks
* `new_cluster` - Same set of parameters as for [databricks_cluster](cluster.md) resource.

### schedule Configuration Block

* `quartz_cron_expression` - (Required) A [Cron expression using Quartz syntax](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) that describes the schedule for a job. This field is required.
* `timezone_id` - (Required) A Java timezone ID. The schedule for a job will be resolved with respect to this timezone. See Java TimeZone for details. This field is required.
* `pause_status` - (Optional) Indicate whether this schedule is paused or not. Either “PAUSED” or “UNPAUSED”. When the pause_status field is omitted and a schedule is provided, the server will default to using "UNPAUSED" as a value for pause_status.

### spark_jar_task Configuration Block

* `parameters` - (Optional) (List) Parameters passed to the main method.
* `main_class_name` - (Optional) The full name of the class containing the main method to be executed. This class must be contained in a JAR provided as a library. The code should use `SparkContext.getOrCreate` to obtain a Spark context; otherwise, runs of the job will fail.

### spark_submit_task Configuration Block

You can invoke Spark submit tasks only on new clusters. **In the `new_cluster` specification, `libraries` and `spark_conf` are not supported**. Instead, use --jars and --py-files to add Java and Python libraries and `--conf` to set the Spark configuration. By default, the Spark submit job uses all available memory (excluding reserved memory for Databricks services). You can set `--driver-memory`, and `--executor-memory` to a smaller value to leave some room for off-heap usage. **Please use `spark_jar_task`, `spark_python_task` or `notebook_task` wherever possible**.

* `parameters` - (Optional) (List) Command-line parameters passed to spark submit.

### spark_python_task Configuration Block

* `python_file` - (Required) The URI of the Python file to be executed. [databricks_dbfs_file](dbfs_file.md#path) and S3 paths are supported. This field is required.
* `parameters` - (Optional) (List) Command line parameters passed to the Python file.

### notebook_task Configuration Block

* `base_parameters` - (Optional) (Map) Base parameters to be used for each run of this job. If the run is initiated by a call to run-now with parameters specified, the two parameters maps will be merged. If the same key is specified in base_parameters and in run-now, the value from run-now will be used. If the notebook takes a parameter that is not specified in the job’s base_parameters or the run-now override parameters, the default value from the notebook will be used. Retrieve these parameters in a notebook using `dbutils.widgets.get`.
* `notebook_path` - (Required) The absolute path of the [databricks_notebook](notebook.md#path) to be run in the Databricks workspace. This path must begin with a slash. This field is required.

### pipeline_task Configuration Block

* `pipeline_id` - (Required) The pipeline's unique ID.

### python_wheel_task Configuration Block

* `entry_point` - (Optional) Python function as entry point for the task
* `package_name` - (Optional) Name of Python package
* `parameters` - (Optional) Parameters for the task
* `named_parameters` - (Optional) Named parameters for the task

### dbt_task Configuration Block

* `commands` - (Required) (Array) Series of dbt commands to execute in sequence. Every command must start with "dbt".
* `project_directory` - (Optional) The relative path to the directory in the repository specified in `git_source` where dbt should look in for the `dbt_project.yml` file. If not specified, defaults to the repository's root directory. Equivalent to passing `--project-dir` to a dbt command.
* `profiles_directory` - (Optional) The relative path to the directory in the repository specified by `git_source` where dbt should look in for the `profiles.yml` file. If not specified, defaults to the repository's root directory. Equivalent to passing `--profile-dir` to a dbt command.
* `catalog` - (Optional) The name of the catalog to use inside Unity Catalog.
* `schema` - (Optional) The name of the schema dbt should run in. Defaults to `default`.
* `warehouse_id` - (Optional) The ID of the SQL warehouse that dbt should execute against.

You also need to include a `git_source` block to configure the repository that contains the dbt project.

### sql_task Configuration Block

One of the `query`, `dashboard` or `alert` needs to be provided.

* `warehouse_id` - (Required) ID of the (the [databricks_sql_endpoint](sql_endpoint.md)) that will be used to execute the task.  Only serverless warehouses are supported right now.
* `parameters` - (Optional) (Map) parameters to be used for each run of this task. The SQL alert task does not support custom parameters.
* `query` - (Optional) block consisting of single string field: `query_id` - identifier of the Databricks SQL Query ([databricks_sql_query](sql_query.md)).
* `dashboard` - (Optional) block consisting of single string field: `dashboard_id` - identifier of the Databricks SQL Dashboard [databricks_sql_dashboard](sql_dashboard.md).
* `alert` - (Optional) block consisting of single string field: `alert_id` - identifier of the Databricks SQL Alert.

### email_notifications Configuration Block

* `on_failure` - (Optional) (List) list of emails to notify on failure
* `no_alert_for_skipped_runs` - (Optional) (Bool) don't send alert for skipped runs
* `on_start` - (Optional) (List) list of emails to notify on failure
* `on_success` - (Optional) (List) list of emails to notify on failure

### git_source Configuration Block

This block is used to specify Git repository information & branch/tag/commit that will be used to pull source code from to execute a job. Supported options are:

* `url` - (Required) URL of the Git repository to use.
* `provider` - (Optional, if it's possible to detect Git provider by host name) case insensitive name of the Git provider.  Following values are supported right now (could be a subject for change, consult [Repos API documentation](https://docs.databricks.com/dev-tools/api/latest/repos.html)): `gitHub`, `gitHubEnterprise`, `bitbucketCloud`, `bitbucketServer`, `azureDevOpsServices`, `gitLab`, `gitLabEnterpriseEdition`.
* `branch` - name of the Git branch to use. Conflicts with `tag` and `commit`.
* `tag` - name of the Git branch to use. Conflicts with `branch` and `commit`.
* `commit` - hash of Git commit to use. Conflicts with `branch` and `tag`.

### Exported attributes

In addition to all arguments above, the following attributes are exported:

* `url` - URL of the job on the given workspace

## Access Control

By default, all users can create and modify jobs unless an administrator [enables jobs access control](https://docs.databricks.com/administration-guide/access-control/jobs-acl.html). With jobs access control, individual permissions determine a user’s abilities. 

* [databricks_permissions](permissions.md#Job-usage) can control which groups or individual users can *Can View*, *Can Manage Run*, and *Can Manage*.
* [databricks_cluster_policy](cluster_policy.md) can control which kinds of clusters users can create for jobs.

## Timeouts

The `timeouts` block allows you to specify `create` and `update` timeouts if you have an `always_running` job. Please launch `TF_LOG=DEBUG terraform apply` whenever you observe timeout issues.

```
timeouts {
  create = "20m"
  update = "20m
}
```

## Import

The resource job can be imported using the id of the job

```bash
$ terraform import databricks_job.this <job-id>
```

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_cluster](cluster.md) to create [Databricks Clusters](https://docs.databricks.com/clusters/index.html).
* [databricks_cluster_policy](cluster_policy.md) to create a [databricks_cluster](cluster.md) policy, which limits the ability to create clusters based on a set of rules.
* [databricks_current_user](../data-sources/current_user.md) data to retrieve information about [databricks_user](user.md) or [databricks_service_principal](service_principal.md), that is calling Databricks REST API.
* [databricks_dbfs_file](../data-sources/dbfs_file.md) data to get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file_paths](../data-sources/dbfs_file_paths.md) data to get list of file names from get file content from [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_dbfs_file](dbfs_file.md) to manage relatively small files on [Databricks File System (DBFS)](https://docs.databricks.com/data/databricks-file-system.html).
* [databricks_global_init_script](global_init_script.md) to manage [global init scripts](https://docs.databricks.com/clusters/init-scripts.html#global-init-scripts), which are run on all [databricks_cluster](cluster.md#init_scripts) and [databricks_job](job.md#new_cluster).
* [databricks_instance_pool](instance_pool.md) to manage [instance pools](https://docs.databricks.com/clusters/instance-pools/index.html) to reduce [cluster](cluster.md) start and auto-scaling times by maintaining a set of idle, ready-to-use instances.
* [databricks_instance_profile](instance_profile.md) to manage AWS EC2 instance profiles that users can launch [databricks_cluster](cluster.md) and access data, like [databricks_mount](mount.md).
* [databricks_jobs] data to get all jobs and their names from a workspace.
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
* [databricks_node_type](../data-sources/node_type.md) data to get the smallest node type for [databricks_cluster](cluster.md) that fits search criteria, like amount of RAM or number of cores.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_pipeline](pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html). 
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
* [databricks_spark_version](../data-sources/spark_version.md) data to get [Databricks Runtime (DBR)](https://docs.databricks.com/runtime/dbr.html) version that could be used for `spark_version` parameter in [databricks_cluster](cluster.md) and other resources.
* [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
