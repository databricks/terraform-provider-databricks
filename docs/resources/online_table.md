---
subcategory: "Unity Catalog"
---
# databricks_online_table (Resource)

This resource allows you to create [Online Table](https://docs.databricks.com/en/machine-learning/feature-store/online-tables.html) in Databricks.  An online table is a read-only copy of a Delta Table that is stored in row-oriented format optimized for online access. Online tables are fully serverless tables that auto-scale throughput capacity with the request load and provide low latency and high throughput access to data of any scale. Online tables are designed to work with Databricks Model Serving, Feature Serving, and retrieval-augmented generation (RAG) applications where they are used for fast data lookups.

-> This resource can only be used with a workspace-level provider!

## Example Usage

```hcl
resource "databricks_online_table" "this" {
  name = "main.default.online_table"
  spec {
    source_table_full_name = "main.default.source_table"
    primary_key_columns = [
      "id"
    ]
    run_triggered {
    }
  }
}
```

## Argument Reference

~> If any parameter changes, online table is recreated.

The following arguments are supported - check [API docs](https://docs.databricks.com/api/workspace/onlinetables/create) for all supported parameters:

* `name` - (Required) 3-level name of the Online Table to create.
* `spec` - (Required) object containing specification of the online table:
  * `source_table_full_name` - (Required) full name of the source table.
  * `primary_key_columns` - (Required) list of the columns comprising the primary key.
  * `timeseries_key` - (Optional) Time series key to deduplicate (tie-break) rows with the same primary key.
  * `perform_full_copy` - (Optional) Whether to create a full-copy pipeline -- a pipeline that stops after creates a full copy of the source table upon initialization and does not process any change data feeds (CDFs) afterwards. The pipeline can still be manually triggered afterwards, but it always perform a full copy of the source table and there are no incremental updates. This mode is useful for syncing views or tables without CDFs to online tables. Note that the full-copy pipeline only supports "triggered" scheduling policy.
  * `run_continuously` - empty block that specifies that pipeline runs continuously after generating the initial data.  Conflicts with `run_triggered`.
  * `run_triggered` - empty block that specifies that pipeline stops after generating the initial data and can be triggered later (manually, through a cron job or through data triggers).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The same as the name of the online table.
* `spec`:
  * `pipeline_id` - ID of the associated Delta Live Table pipeline.
* `status` - object describing status of the online table:
  * `detailed_state` - The state of the online table.
  * `message` - A text description of the current state of the online table.
* `table_serving_url` - Data serving REST API URL for this table.
* `unity_catalog_provisioning_state` - The provisioning state of the online table entity in Unity Catalog. This is distinct from the state of the data synchronization pipeline (i.e. the table may be in "ACTIVE" but the pipeline may be in "PROVISIONING" as it runs asynchronously).

## Import

The resource can be imported using the name of the Online Table:

```hcl
import {
  to = databricks_online_table.this
  id = "<endpoint-name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_online_table.this "<endpoint-name>"
```
