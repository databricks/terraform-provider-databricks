# NEXT CHANGELOG

## Release v1.142.0

### Important Changes

### Breaking Changes

### New Features and Improvements

* Add `databricks_app_slack` and `databricks_app_teams` configuration blocks to `databricks_notification_destination`, and accept `DATABRICKS_APP_SLACK` and `DATABRICKS_APP_TEAMS` as `type` in the `databricks_notification_destinations` data source ([#6131](https://github.com/databricks/terraform-provider-databricks/pull/6131)).

### Bug Fixes

* Fix data races when concurrent resource operations access the unified provider's client caches ([#6120](https://github.com/databricks/terraform-provider-databricks/pull/6120)).

* Support importing workspace-level resources through an account-level or unified-host provider by passing the workspace ID in the import block `identity` (Terraform 1.12 or later) ([#6149](https://github.com/databricks/terraform-provider-databricks/pull/6149)).

  Terraform does not pass `provider_config` to the provider during import. Importing with only `id` through such a provider now fails with an error that shows the `identity` import block, instead of reading the resource from the account API.

  On Terraform 1.12 and 1.13, state refreshed by this version cannot be read by older provider versions (`failed to decode identity`). To downgrade the provider, run `terraform apply -refresh-only` once with the older provider and Terraform 1.14 or later, or remove the resource from state and import it again with the older provider.

### Documentation

### Exporter

### Internal Changes
