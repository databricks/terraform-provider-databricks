# NEXT CHANGELOG

## Release v1.138.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fix `databricks_cluster` dropping a configured `autoscale` block when resizing a running cluster ([#6018](https://github.com/databricks/terraform-provider-databricks/pull/6018)).

  When `num_workers` changed in the plan while the `autoscale` block stayed the same, the provider called `POST /api/2.1/clusters/resize` with `num_workers` and `ForceSendFields: ["NumWorkers"]`, so the API discarded autoscaling and pinned the running cluster to the fixed size - `0` workers when the configuration declares only `autoscale`. The apply reported success. The four predicates that selected the resize path decided from what had changed in the diff rather than from the desired configuration: `isNumWorkersResizeForNonAutoscalingCluster` never checked `cluster.Autoscale == nil` despite its name. They are replaced by a single condition that builds the resize request from the desired configuration, so an `autoscale` block in the configuration always takes precedence over `num_workers`. This also fixes removing `autoscale` while `num_workers` is unchanged, which previously issued a resize request carrying neither a size nor a range.

### Documentation

### Exporter

### Internal Changes
