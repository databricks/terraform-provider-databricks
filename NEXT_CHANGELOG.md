# NEXT CHANGELOG

## Release v1.105.0

### Breaking Changes

* Return empty string for `data_source_id` in `databricks_sql_warehouse` and `databricks_sql_endpoint` if data source API call fails ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### New Features and Improvements

* Allow the `cluster_size` argument to be specified as "5X-Large" in `databricks_sql_warehouse`
* Update the maximum allowed value for `max_num_clusters` argument in `databricks_sql_warehouse` to 40 (see [API docs](https://docs.databricks.com/api/workspace/warehouses/create#max_num_clusters) for details).

### Bug Fixes

* Dashboard File Content Change Detection When Using `file_path` ([#5359])(https://github.com/databricks/terraform-provider-databricks/pull/5359)
* Fix permanent drift in `databricks_model_serving` when using `*_plaintext` credential fields for external models ([#5125](https://github.com/databricks/terraform-provider-databricks/pull/5125))

### Documentation

* Mark `data_source_id` as deprecated in `databricks_sql_warehouse` and `databricks_sql_endpoint` ([#5312](https://github.com/databricks/terraform-provider-databricks/pull/5312))

### Exporter

* Add exporting for RFA destinations and Delta Sharing providers ([#5337](https://github.com/databricks/terraform-provider-databricks/pull/5337))

### Internal Changes

* Update Go SDK to v0.104.0.
