---
subcategory: "Compute"
---

# databricks_job Resource

The `databricks_job` resource allows you to manage [Databricks Jobs](https://docs.databricks.com/jobs.html) to run non-interactive code in a [databricks_cluster](cluster.md).

## Example Usage

-> In Terraform configuration, it is recommended to define tasks in alphabetical order of their `task_key` arguments, so that you get consistent and readable diff. Whenever tasks are added or removed, or `task_key` is renamed, you'll observe a change in the majority of tasks. It's related to the fact that the current version of the provider treats `task` blocks as an ordered list. Alternatively, `task` block could have been an unordered set, though end-users would see the entire block replaced upon a change in single property of the task.

It is possible to create [a Databricks job](https://docs.databricks.com/data-engineering/jobs/jobs-user-guide.html) using `task` blocks. A single task is defined with the `task` block containing one of the `*_task` blocks, `task_key`, and additional arguments described below.

```hcl
resource "databricks_job" "this" {
  name        = "Job with multiple tasks"
  description = "This job executes multiple tasks on a shared job cluster, which will be provisioned as part of execution, and terminated once all tasks are finished."

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

## Argument Reference

The resource supports the following arguments:

* `name` - (Optional) An optional name for the job. The default value is Untitled.
* `description` - (Optional) An optional description for the job. The maximum length is 1024 characters in UTF-8 encoding.
* `task` - (Optional) A list of task specification that the job will execute. See [task Configuration Block](#task-configuration-block) below.
* `job_cluster` - (Optional) A list of job [databricks_cluster](cluster.md) specifications that can be shared and reused by tasks of this job. Libraries cannot be declared in a shared job cluster. You must declare dependent libraries in task settings. *Multi-task syntax*
* `schedule` - (Optional) An optional periodic schedule for this job. The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API request to runNow. See [schedule Configuration Block](#schedule-configuration-block) below.
* `trigger` - (Optional) The conditions that triggers the job to start. See [trigger Configuration Block](#trigger-configuration-block) below.
* `continuous`- (Optional) Configuration block to configure pause status. See [continuous Configuration Block](#continuous-configuration-block).
* `queue` - (Optional) The queue status for the job. See [queue Configuration Block](#queue-configuration-block) below.
* `always_running` - (Optional, Deprecated) (Bool) Whenever the job is always running, like a Spark Streaming application, on every update restart the current active run or start it again, if nothing it is not running. False by default. Any job runs are started with `parameters` specified in `spark_jar_task` or `spark_submit_task` or `spark_python_task` or `notebook_task` blocks.
* `run_as` - (Optional) The user or the service prinicipal the job runs as. See [run_as Configuration Block](#run_as-configuration-block) below.
* `control_run_state` - (Optional) (Bool) If true, the Databricks provider will stop and start the job as needed to ensure that the active run for the job reflects the deployed configuration. For continuous jobs, the provider respects the `pause_status` by stopping the current active run. This flag cannot be set for non-continuous jobs.

  When migrating from `always_running` to `control_run_state`, set `continuous` as follows:

  ```hcl
  continuous { }
  ```

* `library` - (Optional) (List) An optional list of libraries to be installed on the cluster that will execute the job. See [library Configuration Block](#library-configuration-block) below.
* `git_source` - (Optional) Specifices the a Git repository for task source code. See [git_source Configuration Block](#git_source-configuration-block) below.
* `parameter` - (Optional) Specifices job parameter for the job. See [parameter Configuration Block](#parameter-configuration-block)
* `timeout_seconds` - (Optional) (Integer) An optional timeout applied to each run of this job. The default behavior is to have no timeout.
* `min_retry_interval_millis` - (Optional) (Integer) An optional minimal interval in milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that unsuccessful runs are immediately retried.
* `max_concurrent_runs` - (Optional) (Integer) An optional maximum allowed number of concurrent runs of the job. Defaults to *1*.
* `email_notifications` - (Optional) (List) An optional set of email addresses notified when runs of this job begins, completes or fails. The default behavior is to not send any emails. This field is a block and is [documented below](#email_notifications-configuration-block).
* `webhook_notifications` - (Optional) (List) An optional set of system destinations (for example, webhook destinations or Slack) to be notified when runs of this job begins, completes or fails. The default behavior is to not send any notifications. This field is a block and is [documented below](#webhook_notifications-configuration-block).
* `notification_settings` - (Optional) An optional block controlling the notification settings on the job level [documented below](#notification_settings-configuration-block).
* `health` - (Optional) An optional block that specifies the health conditions for the job [documented below](#health-configuration-block).
* `tags` - (Optional) An optional map of the tags associated with the job. See [tags Configuration Map](#tags-configuration-map)
* `budget_policy_id` - (Optional) The ID of the user-specified budget policy to use for this job. If not specified, a default budget policy may be applied when creating or modifying the job.
* `edit_mode` - (Optional) If `"UI_LOCKED"`, the user interface for the job will be locked. If `"EDITABLE"` (the default), the user interface will be editable.

### task Configuration Block

This block describes individual tasks:

* `task_key` - (Required) string specifying an unique key for a given task.
* `*_task` - (Required) one of the specific task blocks described below:
  * `condition_task`
  * `dbt_task`
  * `for_each_task`
  * `notebook_task`
  * `pipeline_task`
  * `python_wheel_task`
  * `run_job_task`
  * `spark_jar_task`
  * `spark_python_task`
  * `spark_submit_task`
  * `sql_task`
* `depends_on` - (Optional) block specifying dependency(-ies) for a given task.
* `description` - (Optional) description for this task.
* `disable_auto_optimization` - (Optional) A flag to disable auto optimization in serverless tasks.
* `email_notifications` - (Optional) An optional block to specify a set of email addresses notified when this task begins, completes or fails. The default behavior is to not send any emails. This block is [documented below](#email_notifications-configuration-block).
* `environment_key` - (Optional) identifier of an `environment` block that is used to specify libraries.  Required for some tasks (`spark_python_task`, `python_wheel_task`, ...) running on serverless compute.
* `existing_cluster_id` - (Optional) Identifier of the [interactive cluster](cluster.md) to run job on.  *Note: running tasks on interactive clusters may lead to increased costs!*
* `health` - (Optional) block described below that specifies health conditions for a given task.
* `job_cluster_key` - (Optional) Identifier of the Job cluster specified in the `job_cluster` block.
* `library` - (Optional) (Set) An optional list of libraries to be installed on the cluster that will execute the job.
* `max_retries` - (Optional) (Integer) An optional maximum number of times to retry an unsuccessful run. A run is considered to be unsuccessful if it completes with a `FAILED` or `INTERNAL_ERROR` lifecycle state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior is to never retry. A run can have the following lifecycle state: `PENDING`, `RUNNING`, `TERMINATING`, `TERMINATED`, `SKIPPED` or `INTERNAL_ERROR`.
* `min_retry_interval_millis` - (Optional) (Integer) An optional minimal interval in milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that unsuccessful runs are immediately retried.
* `new_cluster` - (Optional) Task will run on a dedicated cluster.  See [databricks_cluster](cluster.md) documentation for specification. *Some parameters, such as `autotermination_minutes`, `is_pinned`, `workload_type` aren't supported!*
* `retry_on_timeout` - (Optional) (Bool) An optional policy to specify whether to retry a job when it times out. The default behavior is to not retry on timeout.
* `run_if` - (Optional) An optional value indicating the condition that determines whether the task should be run once its dependencies have been completed. One of `ALL_SUCCESS`, `AT_LEAST_ONE_SUCCESS`, `NONE_FAILED`, `ALL_DONE`, `AT_LEAST_ONE_FAILED` or `ALL_FAILED`. When omitted, defaults to `ALL_SUCCESS`.
* `timeout_seconds` - (Optional) (Integer) An optional timeout applied to each run of this job. The default behavior is to have no timeout.
* `webhook_notifications` - (Optional) (List) An optional set of system destinations (for example, webhook destinations or Slack) to be notified when runs of this task begins, completes or fails. The default behavior is to not send any notifications. This field is a block and is documented below.

-> If no `job_cluster_key`, `existing_cluster_id`, or `new_cluster` were specified in task definition, then task will executed using serverless compute.

#### condition_task Configuration Block

The `condition_task` specifies a condition with an outcome that can be used to control the execution of dependent tasks.

* `left` - The left operand of the condition task. It could be a string value, job state, or a parameter reference.
* `right` - The right operand of the condition task. It could be a string value, job state, or parameter reference.
* `op` - The string specifying the operation used to compare operands.  Currently, following operators are supported: `EQUAL_TO`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`. (Check the [API docs](https://docs.databricks.com/api/workspace/jobs/create) for the latest information).

This task does not require a cluster to execute and does not support retries or notifications.

#### dbt_task Configuration Block

* `commands` - (Required) (Array) Series of dbt commands to execute in sequence. Every command must start with "dbt".
* `source` - (Optional) The source of the project. Possible values are `WORKSPACE` and `GIT`.  Defaults to `GIT` if a `git_source` block is present in the job definition.
* `project_directory` - (Required when `source` is `WORKSPACE`) The path where dbt should look for `dbt_project.yml`. Equivalent to passing `--project-dir` to the dbt CLI.
  * If `source` is `GIT`: Relative path to the directory in the repository specified in the `git_source` block. Defaults to the repository's root directory when not specified.
  * If `source` is `WORKSPACE`: Absolute path to the folder in the workspace.
* `profiles_directory` - (Optional) The relative path to the directory in the repository specified by `git_source` where dbt should look in for the `profiles.yml` file. If not specified, defaults to the repository's root directory. Equivalent to passing `--profile-dir` to a dbt command.
* `catalog` - (Optional) The name of the catalog to use inside Unity Catalog.
* `schema` - (Optional) The name of the schema dbt should run in. Defaults to `default`.
* `warehouse_id` - (Optional) The ID of the SQL warehouse that dbt should execute against.

You also need to include a `git_source` block to configure the repository that contains the dbt project.

#### for_each_task Configuration Block

* `concurrency` - (Optional) Controls the number of active iteration task runs. Default is 20, maximum allowed is 100.
* `inputs` - (Required) (String) Array for task to iterate on. This can be a JSON string or a reference to an array parameter.
* `task` - (Required) Task to run against the `inputs` list.

#### notebook_task Configuration Block

* `notebook_path` - (Required) The path of the [databricks_notebook](notebook.md#path) to be run in the Databricks workspace or remote repository. For notebooks stored in the Databricks workspace, the path must be absolute and begin with a slash. For notebooks stored in a remote repository, the path must be relative. This field is required.
* `source` - (Optional) Location type of the notebook, can only be `WORKSPACE` or `GIT`. When set to `WORKSPACE`, the notebook will be retrieved from the local Databricks workspace. When set to `GIT`, the notebook will be retrieved from a Git repository defined in `git_source`. If the value is empty, the task will use `GIT` if `git_source` is defined and `WORKSPACE` otherwise.
* `base_parameters` - (Optional) (Map) Base parameters to be used for each run of this job. If the run is initiated by a call to run-now with parameters specified, the two parameters maps will be merged. If the same key is specified in base_parameters and in run-now, the value from run-now will be used. If the notebook takes a parameter that is not specified in the job’s base_parameters or the run-now override parameters, the default value from the notebook will be used. Retrieve these parameters in a notebook using `dbutils.widgets.get`.
* `warehouse_id` - (Optional) ID of the (the [databricks_sql_endpoint](sql_endpoint.md)) that will be used to execute the task with SQL notebook.

#### pipeline_task Configuration Block

* `pipeline_id` - (Required) The pipeline's unique ID.
* `full_refresh` - (Optional) (Bool) Specifies if there should be full refresh of the pipeline.

-> The following configuration blocks are only supported inside a `task` block

#### python_wheel_task Configuration Block

* `entry_point` - (Optional) Python function as entry point for the task
* `package_name` - (Optional) Name of Python package
* `parameters` - (Optional) Parameters for the task
* `named_parameters` - (Optional) Named parameters for the task

#### run_job_task Configuration Block

* `job_id` - (Required)(String) ID of the job
* `job_parameters` - (Optional)(Map) Job parameters for the task

#### spark_jar_task Configuration Block

* `parameters` - (Optional) (List) Parameters passed to the main method.
* `main_class_name` - (Optional) The full name of the class containing the main method to be executed. This class must be contained in a JAR provided as a library. The code should use `SparkContext.getOrCreate` to obtain a Spark context; otherwise, runs of the job will fail.

#### spark_python_task Configuration Block

* `python_file` - (Required) The URI of the Python file to be executed. [databricks_dbfs_file](dbfs_file.md#path), cloud file URIs (e.g. `s3:/`, `abfss:/`, `gs:/`), workspace paths and remote repository are supported. For Python files stored in the Databricks workspace, the path must be absolute and begin with `/Repos`. For files stored in a remote repository, the path must be relative. This field is required.
* `source` - (Optional) Location type of the Python file, can only be `GIT`. When set to `GIT`, the Python file will be retrieved from a Git repository defined in `git_source`.
* `parameters` - (Optional) (List) Command line parameters passed to the Python file.

#### spark_submit_task Configuration Block

You can invoke Spark submit tasks only on new clusters. **In the `new_cluster` specification, `libraries` and `spark_conf` are not supported**. Instead, use --jars and --py-files to add Java and Python libraries and `--conf` to set the Spark configuration. By default, the Spark submit job uses all available memory (excluding reserved memory for Databricks services). You can set `--driver-memory`, and `--executor-memory` to a smaller value to leave some room for off-heap usage. **Please use `spark_jar_task`, `spark_python_task` or `notebook_task` wherever possible**.

* `parameters` - (Optional) (List) Command-line parameters passed to spark submit.

#### sql_task Configuration Block

One of the `query`, `dashboard` or `alert` needs to be provided.

* `warehouse_id` - (Required) ID of the (the [databricks_sql_endpoint](sql_endpoint.md)) that will be used to execute the task.  Only Serverless & Pro warehouses are supported right now.
* `parameters` - (Optional) (Map) parameters to be used for each run of this task. The SQL alert task does not support custom parameters.
* `query` - (Optional) block consisting of single string field: `query_id` - identifier of the Databricks Query ([databricks_query](query.md)).
* `dashboard` - (Optional) block consisting of following fields:
  * `dashboard_id` - (Required) (String) identifier of the Databricks SQL Dashboard [databricks_sql_dashboard](sql_dashboard.md).
  * `subscriptions` - (Optional) a list of subscription blocks consisting out of one of the required fields: `user_name` for user emails or `destination_id` - for Alert destination's identifier.
  * `custom_subject` - (Optional) string specifying a custom subject of email sent.
  * `pause_subscriptions` - (Optional) flag that specifies if subscriptions are paused or not.
* `alert` - (Optional) block consisting of following fields:
  * `alert_id` - (Required) (String) identifier of the Databricks Alert ([databricks_alert](alert.md)).
  * `subscriptions` - (Optional) a list of subscription blocks consisting out of one of the required fields: `user_name` for user emails or `destination_id` - for Alert destination's identifier.
  * `pause_subscriptions` - (Optional) flag that specifies if subscriptions are paused or not.
* `file` - (Optional) block consisting of single string fields:
  * `source` - (Optional) The source of the project. Possible values are `WORKSPACE` and `GIT`.
  * `path` - If `source` is `GIT`: Relative path to the file in the repository specified in the `git_source` block with SQL commands to execute. If `source` is `WORKSPACE`: Absolute path to the file in the workspace with SQL commands to execute.

Example

```hcl
resource "databricks_job" "sql_aggregation_job" {
  name = "Example SQL Job"
  task {
    task_key = "run_agg_query"
    sql_task {
      warehouse_id = databricks_sql_endpoint.sql_job_warehouse.id
      query {
        query_id = databricks_sql_query.agg_query.id
      }
    }
  }
  task {
    task_key = "run_dashboard"
    sql_task {
      warehouse_id = databricks_sql_endpoint.sql_job_warehouse.id
      dashboard {
        dashboard_id = databricks_sql_dashboard.dash.id
        subscriptions {
          user_name = "user@domain.com"
        }
      }
    }
  }
  task {
    task_key = "run_alert"
    sql_task {
      warehouse_id = databricks_sql_endpoint.sql_job_warehouse.id
      alert {
        alert_id = databricks_sql_alert.alert.id
        subscriptions {
          user_name = "user@domain.com"
        }
      }
    }
  }
}
```

#### library Configuration Block

This block descripes an optional library to be installed on the cluster that will execute the job. For multiple libraries, use multiple blocks. If the job specifies more than one task, these blocks needs to be placed within the task block. Please consult [libraries section of the databricks_cluster](cluster.md#library-configuration-block) resource for more information.

```hcl
resource "databricks_job" "this" {
  library {
    pypi {
      package = "databricks-mosaic==0.3.14"
    }
  }
}
```

#### environment Configuration Block

This block describes [an Environment](https://docs.databricks.com/en/compute/serverless/dependencies.html) that is used to specify libraries used by the tasks running on serverless compute.  This block contains following attributes:

* `environment_key` - an unique identifier of the Environment.  It will be referenced from `environment_key` attribute of corresponding task.
* `spec` - block describing the Environment. Consists of following attributes:
  * `client` - (Required, string) client version used by the environment.
  * `dependencies` - (list of strings) List of pip dependencies, as supported by the version of pip in this environment. Each dependency is a pip requirement file line.  See [API docs](https://docs.databricks.com/api/workspace/jobs/create#environments-spec-dependencies) for more information.

```hcl
environment {
  spec {
    dependencies = ["foo==0.0.1", "-r /Workspace/test/requirements.txt"]
    client       = "1"
  }
  environment_key = "Default"
}
```

#### depends_on Configuration Block

This block describes upstream dependencies of a given task. For multiple upstream dependencies, use multiple blocks.

* `task_key` - (Required) The name of the task this task depends on.
* `outcome` - (Optional, string) Can only be specified on condition task dependencies. The outcome of the dependent task that must be met for this task to run. Possible values are `"true"` or `"false"`.

-> Similar to the tasks themselves, each dependency inside the task need to be declared in alphabetical order with respect to task_key in order to get consistent Terraform diffs.

### run_as Configuration Block

The `run_as` block allows specifying the user or the service principal that the job runs as. If not specified, the job runs as the user or service
principal that created the job. Only one of `user_name` or `service_principal_name` can be specified.

* `user_name` - (Optional) The email of an active workspace user. Non-admin users can only set this field to their own email.
* `service_principal_name` - (Optional) The application ID of an active service principal. Setting this field requires the `servicePrincipal/user` role.

Example:

```hcl
resource "databricks_job" "this" {
  # ...
  run_as {
    service_principal_name = "8d23ae77-912e-4a19-81e4-b9c3f5cc9349"
  }
}
```

### job_cluster Configuration Block

[Shared job cluster](https://docs.databricks.com/jobs.html#use-shared-job-clusters) specification. Allows multiple tasks in the same job run to reuse the cluster.

* `job_cluster_key` - (Required) Identifier that can be referenced in `task` block, so that cluster is shared between tasks
* `new_cluster` - Block with almost the same set of parameters as for [databricks_cluster](cluster.md) resource, except following (check the [REST API documentation for full list of supported parameters](https://docs.databricks.com/api/workspace/jobs/create#job_clusters-new_cluster)):
  * `autotermination_minutes` - isn't supported
  * `is_pinned` - isn't supported
  * `workload_type` - isn't supported

### schedule Configuration Block

* `quartz_cron_expression` - (Required) A [Cron expression using Quartz syntax](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) that describes the schedule for a job. This field is required.
* `timezone_id` - (Required) A Java timezone ID. The schedule for a job will be resolved with respect to this timezone. See Java TimeZone for details. This field is required.
* `pause_status` - (Optional) Indicate whether this schedule is paused or not. Either `PAUSED` or `UNPAUSED`. When the `pause_status` field is omitted and a schedule is provided, the server will default to using `UNPAUSED` as a value for `pause_status`.

### continuous Configuration Block

* `pause_status` - (Optional) Indicate whether this continuous job is paused or not. Either `PAUSED` or `UNPAUSED`. When the `pause_status` field is omitted in the block, the server will default to using `UNPAUSED` as a value for `pause_status`.

### queue Configuration Block

This block describes the queue settings of the job:

* `enabled` - (Required) If true, enable queueing for the job.

### trigger Configuration Block

* `pause_status` - (Optional) Indicate whether this trigger is paused or not. Either `PAUSED` or `UNPAUSED`. When the `pause_status` field is omitted in the block, the server will default to using `UNPAUSED` as a value for `pause_status`.
* `periodic` - (Optional) configuration block to define a trigger for Periodic Triggers consisting of the following attributes:
  * `interval` - (Required) Specifies the interval at which the job should run. This value is required.
  * `unit` - (Required) Options are {"DAYS", "HOURS", "WEEKS"}.
* `file_arrival` - (Optional) configuration block to define a trigger for [File Arrival events](https://learn.microsoft.com/en-us/azure/databricks/workflows/jobs/file-arrival-triggers) consisting of following attributes:
  * `url` - (Required) URL to be monitored for file arrivals. The path must point to the root or a subpath of the external location. Please note that the URL must have a trailing slash character (`/`).
  * `min_time_between_triggers_seconds` - (Optional) If set, the trigger starts a run only after the specified amount of time passed since the last time the trigger fired. The minimum allowed value is 60 seconds.
  * `wait_after_last_change_seconds` - (Optional) If set, the trigger starts a run only after no file activity has occurred for the specified amount of time. This makes it possible to wait for a batch of incoming files to arrive before triggering a run. The minimum allowed value is 60 seconds.

### git_source Configuration Block

This block is used to specify Git repository information & branch/tag/commit that will be used to pull source code from to execute a job. Supported options are:

* `url` - (Required) URL of the Git repository to use.
* `provider` - (Optional, if it's possible to detect Git provider by host name) case insensitive name of the Git provider.  Following values are supported right now (could be a subject for change, consult [Repos API documentation](https://docs.databricks.com/dev-tools/api/latest/repos.html)): `gitHub`, `gitHubEnterprise`, `bitbucketCloud`, `bitbucketServer`, `azureDevOpsServices`, `gitLab`, `gitLabEnterpriseEdition`.
* `branch` - name of the Git branch to use. Conflicts with `tag` and `commit`.
* `tag` - name of the Git branch to use. Conflicts with `branch` and `commit`.
* `commit` - hash of Git commit to use. Conflicts with `branch` and `tag`.

### parameter Configuration Block

This block defines a job-level parameter for the job. You can define several job-level parameters for the job. Supported options are:

* `name` - (Required) The name of the defined parameter. May only contain alphanumeric characters, `_`, `-`, and `.`.
* `default` - (Required) Default value of the parameter.

*You can use this block only together with `task` blocks, not with the legacy tasks specification!*

### email_notifications Configuration Block

This block can be configured on both job and task levels for corresponding effect.

* `on_start` - (Optional) (List) list of emails to notify when the run starts.
* `on_success` - (Optional) (List) list of emails to notify when the run completes successfully.
* `on_failure` - (Optional) (List) list of emails to notify when the run fails.
* `on_duration_warning_threshold_exceeded` - (Optional) (List) list of emails to notify when the duration of a run exceeds the threshold specified by the `RUN_DURATION_SECONDS` metric in the `health` block.

The following parameter is only available for the job level configuration.

* `no_alert_for_skipped_runs` - (Optional) (Bool) don't send alert for skipped runs. (It's recommended to use the corresponding setting in the `notification_settings` configuration block).

### webhook_notifications Configuration Block

Each entry in `webhook_notification` block takes a list `webhook` blocks. The field is documented below.

* `on_start` - (Optional) (List) list of notification IDs to call when the run starts. A maximum of 3 destinations can be specified.
* `on_success` - (Optional) (List) list of notification IDs to call when the run completes successfully. A maximum of 3 destinations can be specified.
* `on_failure` - (Optional) (List) list of notification IDs to call when the run fails. A maximum of 3 destinations can be specified.
* `on_duration_warning_threshold_exceeded` - (Optional) (List) list of notification IDs to call when the duration of a run exceeds the threshold specified by the `RUN_DURATION_SECONDS` metric in the `health` block.

Note that the `id` is not to be confused with the name of the alert destination. The `id` can be retrieved through the API or the URL of Databricks UI `https://<workspace host>/sql/destinations/<notification id>?o=<workspace id>`

Example

```hcl
webhook_notifications {
  on_failure {
    id = "fb99f3dc-a0a0-11ed-a8fc-0242ac120002"
  }
}
```

### webhook Configuration Block

* `id` - ID of the system notification that is notified when an event defined in `webhook_notifications` is triggered.

-> The following configuration blocks can be standalone or nested inside a `task` block

### notification_settings Configuration Block

This block controls notification settings for both email & webhook notifications.
It can be configured on both job and task level for corresponding effect.

* `no_alert_for_skipped_runs` - (Optional) (Bool) don't send alert for skipped runs.
* `no_alert_for_canceled_runs` - (Optional) (Bool) don't send alert for cancelled runs.

The following parameter is only available on task level.

* `alert_on_last_attempt` - (Optional) (Bool) do not send notifications to recipients specified in `on_start` for the retried runs and do not send notifications to recipients specified in `on_failure` until the last retry of the run.

### health Configuration Block

This block describes health conditions for a given job or an individual task. It consists of the following attributes:

* `rules` - (List) list of rules that are represented as objects with the following attributes:
  * `metric` - (Required) string specifying the metric to check.  The only supported metric is `RUN_DURATION_SECONDS` (check [Jobs REST API documentation](https://docs.databricks.com/api/workspace/jobs/create) for the latest information).
  * `op` - (Required) string specifying the operation used to evaluate the given metric. The only supported operation is `GREATER_THAN`.
  * `value` - (Required) integer value used to compare to the given metric.

### tags Configuration Map

`tags` - (Optional) (Map) An optional map of the tags associated with the job. Specified tags will be used as cluster tags for job clusters.

Example

```hcl
resource "databricks_job" "this" {
  # ...
  tags = {
    environment = "dev"
    owner       = "dream-team"
  }
}
```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the job
* `url` - URL of the job on the given workspace

## Access Control

By default, all users can create and modify jobs unless an administrator [enables jobs access control](https://docs.databricks.com/administration-guide/access-control/jobs-acl.html). With jobs access control, individual permissions determine a user’s abilities.

* [databricks_permissions](permissions.md#Job-usage) can control which groups or individual users can *Can View*, *Can Manage Run*, and *Can Manage*.
* [databricks_cluster_policy](cluster_policy.md) can control which kinds of clusters users can create for jobs.

## Single-task syntax (deprecated)

-> **Deprecated** Please define tasks in a `task` block rather than using single-task syntax.

This syntax uses Jobs API 2.0 to create a job with a single task. Only a subset of arguments above is supported (`name`, `libraries`, `email_notifications`, `webhook_notifications`, `timeout_seconds`, `max_retries`, `min_retry_interval_millis`, `retry_on_timeout`, `schedule`, `max_concurrent_runs`), and only a single block of `notebook_task`, `spark_jar_task`, `spark_python_task`, `spark_submit_task` and `pipeline_task` can be specified.

The job cluster is specified using either of the below argument:

* `new_cluster` - (Optional) Same set of parameters as for [databricks_cluster](cluster.md) resource.
* `existing_cluster_id` - (Optional) If existing_cluster_id, the ID of an existing [cluster](cluster.md) that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually restart the cluster if it stops responding. We strongly suggest to use `new_cluster` for greater reliability.

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

## Timeouts

The `timeouts` block allows you to specify `create` and `update` timeouts if you have an `always_running` job. Please launch `TF_LOG=DEBUG terraform apply` whenever you observe timeout issues.

```hcl
timeouts {
  create = "20m"
  update = "20m"
}
```

## Import

The resource job can be imported using the id of the job

```bash
terraform import databricks_job.this <job-id>
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
* [databricks_jobs](../data-sources/jobs.md) data to get all jobs and their names from a workspace.
* [databricks_library](library.md) to install a [library](https://docs.databricks.com/libraries/index.html) on [databricks_cluster](cluster.md).
* [databricks_node_type](../data-sources/node_type.md) data to get the smallest node type for [databricks_cluster](cluster.md) that fits search criteria, like amount of RAM or number of cores.
* [databricks_notebook](notebook.md) to manage [Databricks Notebooks](https://docs.databricks.com/notebooks/index.html).
* [databricks_pipeline](pipeline.md) to deploy [Delta Live Tables](https://docs.databricks.com/data-engineering/delta-live-tables/index.html).
* [databricks_repo](repo.md) to manage [Databricks Repos](https://docs.databricks.com/repos.html).
* [databricks_spark_version](../data-sources/spark_version.md) data to get [Databricks Runtime (DBR)](https://docs.databricks.com/runtime/dbr.html) version that could be used for `spark_version` parameter in [databricks_cluster](cluster.md) and other resources.
* [databricks_workspace_conf](workspace_conf.md) to manage workspace configuration for expert usage.
