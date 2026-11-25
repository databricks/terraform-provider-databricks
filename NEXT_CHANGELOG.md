# NEXT CHANGELOG

## Release v1.146.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Added write-only arguments to `databricks_notification_destination` (`url_wo`, `username_wo` and `password_wo` for generic webhooks, `url_wo` and `oauth_token_wo` for Slack, `url_wo` and `auth_secret_wo` for Microsoft Teams, `integration_key_wo` for PagerDuty) with `config_wo_version`, so webhook URLs, passwords and keys can be supplied without storing them in the Terraform state ([#6306](https://github.com/databricks/terraform-provider-databricks/pull/6306)).

### Bug Fixes

### Documentation

### Exporter

* Added support for exporting Lakewatch resources (`databricks_lakewatch_rule`, `databricks_lakewatch_datasource`, `databricks_lakewatch_parser`, `databricks_lakewatch_parser_version`, and per-catalog configurations) as dependencies of `databricks_catalog` ([#6171](https://github.com/databricks/terraform-provider-databricks/pull/6171)).

### Internal Changes
