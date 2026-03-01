# NEXT CHANGELOG

## Release v1.111.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Mark plaintext credential fields in `databricks_model_serving` as sensitive to prevent them from being displayed in plan/apply output ([#5409](https://github.com/databricks/terraform-provider-databricks/pull/5409)).
* Mark `personal_access_token` as sensitive in `databricks_git_credential` to prevent the value from being displayed in plan and apply output ([#5395](https://github.com/databricks/terraform-provider-databricks/pull/5395)).

### Documentation

* Added documentation note about whitespace handling in `MAP` column types for `databricks_sql_table`.

### Exporter

### Internal Changes

* Add support for host agnostic SQL global config resource.
