---
subcategory: "Quality Monitor"
---
# databricks_quality_monitors_v2 Data Source
This data source can be used to fetch the list of quality monitors v2.

-> **Note** This data source can only be used with an workspace-level provider!

## Example Usage
Getting a list of all quality monitors:

```hcl
data "databricks_quality_monitors_v2" "all" {
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional)



## Attributes
This data source exports a single attribute, `quality_monitors`. It is a list of resources, each with the following attributes:
* `anomaly_detection_config` (AnomalyDetectionConfig)
* `object_id` (string) - The uuid of the request object. For example, schema id
* `object_type` (string) - The type of the monitored object. Can be one of the following: schema

### AnomalyDetectionConfig
* `last_run_id` (string) - Run id of the last run of the workflow
* `latest_run_status` (string) - The status of the last run of the workflow. Possible values are: `ANOMALY_DETECTION_RUN_STATUS_CANCELED`, `ANOMALY_DETECTION_RUN_STATUS_FAILED`, `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`, `ANOMALY_DETECTION_RUN_STATUS_PENDING`, `ANOMALY_DETECTION_RUN_STATUS_RUNNING`, `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`, `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`, `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`