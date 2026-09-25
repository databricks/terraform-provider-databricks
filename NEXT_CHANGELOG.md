# NEXT CHANGELOG

## Release v1.135.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add `string_value_wo` and `string_value_wo_version` attributes to `databricks_secret` resource ([#5480](https://github.com/databricks/terraform-provider-databricks/pull/5480)).

### Bug Fixes

### Documentation
* Fix the VPC endpoint security group rules in the AWS Private Link guide: add the missing inbound ports (5432 and 8443-8451), source the rules from the workspace security group, and drop the unnecessary egress rules ([#5930](https://github.com/databricks/terraform-provider-databricks/pull/5930)).

### Exporter

### Internal Changes
