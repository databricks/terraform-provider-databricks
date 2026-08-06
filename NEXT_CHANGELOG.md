# NEXT CHANGELOG

## Release v1.125.0

### Important Changes

### Breaking Changes

### New Features and Improvements
* Add optional `git_credential_id` attribute to `databricks_repo` resource to allow explicit credential selection ([#5877](https://github.com/databricks/terraform-provider-databricks/pull/5877)).

### Bug Fixes

### Documentation
* Fix the VPC endpoint security group rules in the AWS Private Link guide: add the missing inbound ports (5432 and 8443-8451), source the rules from the workspace security group, and drop the unnecessary egress rules ([#5930](https://github.com/databricks/terraform-provider-databricks/pull/5930)).

### Exporter

### Internal Changes
