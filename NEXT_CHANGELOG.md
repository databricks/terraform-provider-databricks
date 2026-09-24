# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes
* Fix `databricks_grant` and `databricks_grants` for the AI Gateway securables (`model_service`, `mcp_service`, `model_provider_service`). Wiring `databricks_ai_gateway_*.<x>.name` into the corresponding grant field previously failed on apply with `No API found for 'GET /unity-catalog/permissions/model_provider_service/model-provider-services/<full_name>'`, because that `name` attribute carries a resource-name prefix (e.g. `model-provider-services/`) that the permissions API does not expect. The prefix is now stripped before it reaches the API path and the resource ID, and a `DiffSuppressFunc` treats the prefixed and bare forms as equal, so both the `.name` reference and a bare full name apply cleanly with no perpetual diff.
* Fix `databricks_cluster` dropping a configured `autoscale` block when resizing a running cluster ([#6018](https://github.com/databricks/terraform-provider-databricks/pull/6018)).

  When `num_workers` changed in the plan while the `autoscale` block stayed the same, the provider called `POST /api/2.1/clusters/resize` with `num_workers` and `ForceSendFields: ["NumWorkers"]`, so the API discarded autoscaling and pinned the running cluster to the fixed size - `0` workers when the configuration declares only `autoscale`. The apply reported success. The four predicates that selected the resize path decided from what had changed in the diff rather than from the desired configuration: `isNumWorkersResizeForNonAutoscalingCluster` never checked `cluster.Autoscale == nil` despite its name. They are replaced by a single condition that builds the resize request from the desired configuration, so an `autoscale` block in the configuration always takes precedence over `num_workers`. This also fixes removing `autoscale` while `num_workers` is unchanged, which previously issued a resize request carrying neither a size nor a range.

### Documentation

### Exporter

### Internal Changes
