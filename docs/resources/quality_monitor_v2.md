---
subcategory: "Quality Monitor"
---
# databricks_quality_monitor_v2 Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `provider_config` (ProviderConfig, optional) - Namespace containing arguments which can be used to configure the provider

### ProviderConfig
* `workspace_id` (string, required) - Workspace ID of the resource

## Attributes
In addition to the above arguments, the following attributes are exported:
* `anomaly_detection_config` (AnomalyDetectionConfig)

### AnomalyDetectionConfig
* `job_type` (string) - The type of the last run of the workflow. Possible values are: `ANOMALY_DETECTION_JOB_TYPE_INTERNAL_HIDDEN`, `ANOMALY_DETECTION_JOB_TYPE_NORMAL`
* `last_run_id` (string) - Run id of the last run of the workflow
* `latest_run_status` (string) - The status of the last run of the workflow. Possible values are: `ANOMALY_DETECTION_RUN_STATUS_CANCELED`, `ANOMALY_DETECTION_RUN_STATUS_FAILED`, `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`, `ANOMALY_DETECTION_RUN_STATUS_PENDING`, `ANOMALY_DETECTION_RUN_STATUS_RUNNING`, `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`, `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`, `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "object_type,object_id"
  to = databricks_quality_monitor_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_quality_monitor_v2 "object_type,object_id"
```