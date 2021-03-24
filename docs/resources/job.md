---
subcategory: "Compute"
---
# databricks_job Resource

The `databricks_job` resource allows you to create, edit, and delete jobs, which run on either new or existing [clusters](cluster.md).

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

  email_notifications {}
}

output "notebook_url" {
  value = databricks_notebook.this.url
}

output "job_url" {
  value = databricks_job.this.url
}
```

## Argument Reference

The following arguments are required:

* `name` - (Optional) An optional name for the job. The default value is Untitled.
* `new_cluster` - (Optional) Same set of parameters as for [databricks_cluster](cluster.md) resource.
* `existing_cluster_id` - (Optional) If existing_cluster_id, the ID of an existing [cluster](cluster.md) that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually restart the cluster if it stops responding. We strongly suggest to use `new_cluster` for greater reliability.
* `library` - (Optional) (Set) An optional list of libraries to be installed on the cluster that will execute the job. Please consult [libraries section](cluster.md#libraries) for [databricks_cluster](cluster.md) resource.
* `retry_on_timeout` - (Optional) (Bool) An optional policy to specify whether to retry a job when it times out. The default behavior is to not retry on timeout.
* `max_retries` - (Optional) (Integer) An optional maximum number of times to retry an unsuccessful run. A run is considered to be unsuccessful if it completes with a FAILED result_state or INTERNAL_ERROR life_cycle_state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior is to never retry.
* `timeout_seconds` - (Optional) (Integer) An optional timeout applied to each run of this job. The default behavior is to have no timeout.
* `min_retry_interval_millis` - (Optional) (Integer) An optional minimal interval in milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that unsuccessful runs are immediately retried.
* `max_concurrent_runs` - (Optional) (Integer) An optional maximum allowed number of concurrent runs of the job.
* `email_notifications` - (Optional) (List) An optional set of email addresses notified when runs of this job begin and complete and when this job is deleted. The default behavior is to not send any emails. This field is a block and is documented below.
* `schedule` - (Optional) (List) An optional periodic schedule for this job. The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API request to runNow. This field is a block and is documented below.

### schedule Configuration Block

* `quartz_cron_expression` - (Required) A [Cron expression using Quartz syntax](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html) that describes the schedule for a job. This field is required.
* `timezone_id` - (Required) A Java timezone ID. The schedule for a job will be resolved with respect to this timezone. See Java TimeZone for details. This field is required.
* `pause_status` - (Optional) Indicate whether this schedule is paused or not. Either “PAUSED” or “UNPAUSED”.

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

### email_notifications Configuration Block

* `on_failure` - (Optional) (List) list of emails to notify on failure
* `no_alert_for_skipped_runs` - (Optional) (Bool) don't send alert for skipped runs
* `on_start` - (Optional) (List) list of emails to notify on failure
* `on_success` - (Optional) (List) list of emails to notify on failure

## Access Control

By default, all users can create and modify jobs unless an administrator [enables jobs access control](https://docs.databricks.com/administration-guide/access-control/jobs-acl.html). With jobs access control, individual permissions determine a user’s abilities. 

* [databricks_permissions](permissions.md#Job-usage) can control which groups or individual users can *Can View*, *Can Manage Run*, and *Can Manage*.
* [databricks_cluster_policy](cluster_policy.md) can control which kinds of clusters users can create for jobs.

## Import

The resource job can be imported using the id of the job

```bash
$ terraform import databricks_job.this <job-id>
```
