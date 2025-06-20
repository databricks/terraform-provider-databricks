---
subcategory: "Quality Monitor"
---
# databricks_quality_monitor_v2 Data Source
This data source can be used to fetch a quality monitors v2.

-> **Note** This data source can only be used with an workspace-level provider!


## Example Usage
Referring to a quality monitor by uc object type (currently only support `schema`) and object id:

```hcl
data "databricks_schema" "this" {
  name = "my_catalog.my_schema"
}
data "databricks_quality_monitor_v2" "this" {
  object_type = "schema"
  object_id = data.databricks_schema.this.schema_info.schema_id
}
```


## Arguments
The following arguments are supported:
* `object_id` (string, required) - The uuid of the request object. For example, schema id
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: schema

## Attributes
The following attributes are exported:
* `anomaly_detection_config` (AnomalyDetectionConfig) - 
* `object_id` (string) - The uuid of the request object. For example, schema id
* `object_type` (string) - The type of the monitored object. Can be one of the following: schema

### AnomalyDetectionConfig
* `last_run_id` (string) - Run id of the last run of the workflow
* `latest_run_status` (string) - The status of the last run of the workflow. Possible values are: `ANOMALY_DETECTION_RUN_STATUS_CANCELED`, `ANOMALY_DETECTION_RUN_STATUS_FAILED`, `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`, `ANOMALY_DETECTION_RUN_STATUS_PENDING`, `ANOMALY_DETECTION_RUN_STATUS_RUNNING`, `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`, `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`, `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`