# NEXT CHANGELOG

## Release v1.87.1

### Breaking Changes
Terraform databricks provider version 1.86.0 introduced changes to databricks_share resource leading to panic for some users while applying terraform. We are rolling back to SDKv2 implementation in this version. No change is expected for users who are upgrading to latest version from before 1.86.0. Users on 1.86.0 and not facing issues are suggested to not upgrade their version to this patch release. We are investigating the issue and will release a fix soon.

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

### Internal Changes
* Use SDKv2 Implementation for databricks_share resource as default ([#4931](https://github.com/databricks/terraform-provider-databricks/pull/4931))
