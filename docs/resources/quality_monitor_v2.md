---
subcategory: "Quality Monitor"
---
# databricks_quality_monitor_v2 Resource
Users with MANAGE Schema can use quality monitor v2 to set up data quality monitoring checks for UC objects, currently support schema. 


-> **Note** This resource can only be used with an workspace-level provider!


## Example Usage
```hcl
resource "databricks_schema" "this" {
  catalog_name = "my_catalog"
  name = "my_schema"
}
resource "databricks_quality_monitor_v2" "this" {
  object_type = "schema"
  object_id = databricks_schema.this.schema_id
}
```


## Arguments
The following arguments are supported:
* `object_id` (string, required) - The uuid of the request object. For example, schema id
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: schema

## Attributes
In addition to the above arguments, the following attributes are exported:
* `anomaly_detection_config` (AnomalyDetectionConfig)

### AnomalyDetectionConfig
* `last_run_id` (string) - Run id of the last run of the workflow
* `latest_run_status` (string) - The status of the last run of the workflow. Possible values are: `ANOMALY_DETECTION_RUN_STATUS_CANCELED`, `ANOMALY_DETECTION_RUN_STATUS_FAILED`, `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`, `ANOMALY_DETECTION_RUN_STATUS_PENDING`, `ANOMALY_DETECTION_RUN_STATUS_RUNNING`, `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`, `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`, `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = object_type,object_id
  to = databricks_quality_monitor_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_quality_monitor_v2 object_type,object_id
```