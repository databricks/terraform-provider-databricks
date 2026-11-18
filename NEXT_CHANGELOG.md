# NEXT CHANGELOG

## Release v1.145.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add the `USE_SDK_GO_RESOURCES` environment variable, which makes the listed resources read and write through the modular `github.com/databricks/sdk-go` clients ([#6229](https://github.com/databricks/terraform-provider-databricks/pull/6229), [#6231](https://github.com/databricks/terraform-provider-databricks/pull/6231), [#6242](https://github.com/databricks/terraform-provider-databricks/pull/6242)).

  `USE_SDK_GO_RESOURCES=databricks_volume` makes `databricks_volume` read and write through the `github.com/databricks/sdk-go/uc/volumes` v1 client. It is unset by default, and an unknown resource name fails provider configuration.

### Bug Fixes
* Escape `#`, `?` and `%` in the full name of a `databricks_volume` when reading, updating and deleting it ([#6206](https://github.com/databricks/terraform-provider-databricks/pull/6206)).

  Without escaping, these calls for the volume `main.replay.raw#2` went to `main.replay.raw`: a refresh could drop the volume from state and plan to create it again, and a destroy could delete a different volume.

### Documentation

### Exporter

### Internal Changes
