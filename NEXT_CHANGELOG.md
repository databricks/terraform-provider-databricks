# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

### Bug Fixes
* Fixed `databricks_share` failing with `produced an unexpected new value: .comment` when a share's comment is set or removed outside of Terraform (e.g. in the UI). The `comment` attribute is now `Optional`+`Computed`, so an out-of-band value is adopted into state instead of erroring, while `comment = ""` still explicitly clears the description.

### Documentation

### Exporter

### Internal Changes
