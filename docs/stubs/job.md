# databricks_job Resource


## Example Usage
my_usage

## Argument Reference

The following arguments are required:

* `new_cluster` - (Optional) (List) Same set of parameters as for [databricks_cluster](cluster.md) resource. This field is a block and is documented below.

* `library_cran` - (Optional) (Set)  This field is a block and is documented below.

* `name` - (Optional) (String) An optional name for the job. The default value is Untitled.

* `python_file` - (Optional) (String) Deprecated. Please use `spark_python_task`.

* `library_maven` - (Optional) (Set)  This field is a block and is documented below.

* `libraries` - (Optional) (Set) An optional list of libraries to be installed on the cluster that will execute the job. The default value is an empty list. This field is a block and is documented below.

* `retry_on_timeout` - (Optional) (Bool) An optional policy to specify whether to retry a job when it times out. The default behavior is to not retry on timeout.

* `jar_main_class_name` - (Optional) (String) Deprecated. Please use `spark_jar_task`.

* `notebook_task` - (Optional) (List)  This field is a block and is documented below.

* `notebook_path` - (Optional) (String) Deprecated. Please use `notebook_task`.

* `notebook_base_parameters` - (Optional) (Map) Deprecated. Please use `notebook_task`.

* `existing_cluster_id` - (Optional) (String) If existing_cluster_id, the ID of an existing cluster that will be used for all runs of this job. When running jobs on an existing cluster, you may need to manually restart the cluster if it stops responding. We strongly suggest to use `new_cluster` for greater reliability.

* `spark_python_task` - (Optional) (List)  This field is a block and is documented below.

* `max_retries` - (Optional) (Integer) An optional maximum number of times to retry an unsuccessful run. A run is considered to be unsuccessful if it completes with a FAILED result_state or INTERNAL_ERROR life_cycle_state. The value -1 means to retry indefinitely and the value 0 means to never retry. The default behavior is to never retry.

* `jar_uri` - (Optional) (String) 

* `spark_submit_parameters` - (Optional) (List) 

* `timeout_seconds` - (Optional) (Integer) An optional timeout applied to each run of this job. The default behavior is to have no timeout.

* `min_retry_interval_millis` - (Optional) (Integer) An optional minimal interval in milliseconds between the start of the failed run and the subsequent retry run. The default behavior is that unsuccessful runs are immediately retried.

* `max_concurrent_runs` - (Optional) (Integer) An optional maximum allowed number of concurrent runs of the job.

* `email_notifications` - (Optional) (List) An optional set of email addresses notified when runs of this job begin and complete and when this job is deleted. The default behavior is to not send any emails. This field is a block and is documented below.

* `python_parameters` - (Optional) (List) Deprecated. Please use `spark_python_task`.

* `library_pypi` - (Optional) (Set)  This field is a block and is documented below.

* `spark_submit_task` - (Optional) (List)  This field is a block and is documented below.

* `library_jar` - (Optional) (Set)  This field is a block and is documented below.

* `library_whl` - (Optional) (Set)  This field is a block and is documented below.

* `spark_jar_task` - (Optional) (List)  This field is a block and is documented below.

* `schedule` - (Optional) (List) An optional periodic schedule for this job. The default behavior is that the job runs when triggered by clicking Run Now in the Jobs UI or sending an API request to runNow. This field is a block and is documented below.

* `jar_parameters` - (Optional) (List) Deprecated. Please use `spark_jar_task`.

* `library_egg` - (Optional) (Set)  This field is a block and is documented below.



### library_egg Configuration Block


* `status` - (Computed) (String) 

* `path` - (Optional) (String) 

* `messages` - (Computed) (List) 


### schedule Configuration Block


* `quartz_cron_expression` - (Required) (String) 

* `timezone_id` - (Required) (String) 

* `pause_status` - (Optional) (String) 


### spark_jar_task Configuration Block


* `parameters` - (Optional) (List) 

* `jar_uri` - (Optional) (String) 

* `main_class_name` - (Optional) (String) 


### library_whl Configuration Block


* `path` - (Optional) (String) 

* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 


### library_jar Configuration Block


* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `path` - (Optional) (String) 


### spark_submit_task Configuration Block


* `parameters` - (Optional) (List) 


### library_pypi Configuration Block


* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `package` - (Optional) (String) 

* `repo` - (Optional) (String) 


### email_notifications Configuration Block


* `on_failure` - (Optional) (List) 

* `no_alert_for_skipped_runs` - (Optional) (Bool) 

* `on_start` - (Optional) (List) 

* `on_success` - (Optional) (List) 


### spark_python_task Configuration Block


* `python_file` - (Required) (String) 

* `parameters` - (Optional) (List) 


### notebook_task Configuration Block


* `base_parameters` - (Optional) (Map) 

* `notebook_path` - (Required) (String) 


### libraries Configuration Block


* `egg` - (Optional) (String) 

* `whl` - (Optional) (String) 

* `pypi` - (Optional) (List)  This field is a block and is documented below.

* `maven` - (Optional) (List)  This field is a block and is documented below.

* `cran` - (Optional) (List)  This field is a block and is documented below.

* `jar` - (Optional) (String) 


### cran for libraries Configuration Block


* `package` - (Required) (String) 

* `repo` - (Optional) (String) 


### maven for libraries Configuration Block


* `coordinates` - (Required) (String) 

* `repo` - (Optional) (String) 

* `exclusions` - (Optional) (List) 


### pypi for libraries Configuration Block


* `package` - (Required) (String) 

* `repo` - (Optional) (String) 


### library_maven Configuration Block


* `exclusions` - (Optional) (List) 

* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `coordinates` - (Optional) (String) 

* `repo` - (Optional) (String) 


### library_cran Configuration Block


* `messages` - (Computed) (List) 

* `status` - (Computed) (String) 

* `package` - (Optional) (String) 

* `repo` - (Optional) (String) 


### new_cluster Configuration Block


* `cluster_name` - (Optional) (String) 

* `spark_version` - (Required) (String) 

* `node_type_id` - (Computed) (String) 

* `autotermination_minutes` - (Optional) (Integer) 

* `ssh_public_keys` - (Optional) (List) 

* `enable_elastic_disk` - (Computed) (Bool) 

* `driver_node_type_id` - (Computed) (String) 

* `spark_env_vars` - (Optional) (Map) 

* `idempotency_token` - (Optional) (String) 

* `num_workers` - (Optional) (Integer) 

* `autoscale` - (Optional) (List)  This field is a block and is documented below.

* `instance_pool_id` - (Optional) (String) 

* `docker_image` - (Optional) (List)  This field is a block and is documented below.

* `cluster_log_conf` - (Optional) (List)  This field is a block and is documented below.

* `single_user_name` - (Optional) (String) 

* `cluster_id` - (Optional) (String) 

* `policy_id` - (Optional) (String) 

* `aws_attributes` - (Optional) (List)  This field is a block and is documented below.

* `spark_conf` - (Optional) (Map) 

* `custom_tags` - (Optional) (Map) 

* `init_scripts` - (Optional) (List)  This field is a block and is documented below.


### init_scripts for new_cluster Configuration Block


* `s3` - (Optional) (List)  This field is a block and is documented below.

* `dbfs` - (Optional) (List)  This field is a block and is documented below.


### dbfs for init_scripts for new_cluster Configuration Block


* `destination` - (Required) (String) 


### s3 for init_scripts for new_cluster Configuration Block


* `canned_acl` - (Optional) (String) 

* `destination` - (Required) (String) 

* `region` - (Optional) (String) 

* `endpoint` - (Optional) (String) 

* `enable_encryption` - (Optional) (Bool) 

* `encryption_type` - (Optional) (String) 

* `kms_key` - (Optional) (String) 


### aws_attributes for new_cluster Configuration Block


* `availability` - (Computed) (String) 

* `zone_id` - (Computed) (String) 

* `instance_profile_arn` - (Optional) (String) 

* `spot_bid_price_percent` - (Computed) (Integer) 

* `ebs_volume_type` - (Computed) (String) 

* `ebs_volume_count` - (Computed) (Integer) 

* `ebs_volume_size` - (Computed) (Integer) 

* `first_on_demand` - (Computed) (Integer) 


### cluster_log_conf for new_cluster Configuration Block


* `dbfs` - (Optional) (List)  This field is a block and is documented below.

* `s3` - (Optional) (List)  This field is a block and is documented below.


### s3 for cluster_log_conf for new_cluster Configuration Block


* `destination` - (Required) (String) 

* `region` - (Optional) (String) 

* `endpoint` - (Optional) (String) 

* `enable_encryption` - (Optional) (Bool) 

* `encryption_type` - (Optional) (String) 

* `kms_key` - (Optional) (String) 

* `canned_acl` - (Optional) (String) 


### dbfs for cluster_log_conf for new_cluster Configuration Block


* `destination` - (Required) (String) 


### docker_image for new_cluster Configuration Block


* `url` - (Required) (String) 

* `basic_auth` - (Optional) (List)  This field is a block and is documented below.


### basic_auth for docker_image for new_cluster Configuration Block


* `username` - (Required) (String) 

* `password` - (Required) (String) 


### autoscale for new_cluster Configuration Block


* `min_workers` - (Optional) (Integer) 

* `max_workers` - (Optional) (Integer) 






## Import

The resource job can be imported using the `object`, e.g.

```bash
$ terraform import databricks_job.object <job id>
```