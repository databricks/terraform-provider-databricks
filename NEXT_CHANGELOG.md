# NEXT CHANGELOG

## Release v1.115.0

### Breaking Changes

### New Features and Improvements

### Bug Fixes

* Fix `databricks_service_principal` data source failing on account-level provider with `cannot populate provider_config for service principal: failed to resolve workspace_id` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Fix `databricks_service_principals` data source failing on account-level provider with the same `cannot populate provider_config for service principals: failed to resolve workspace_id` regression ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)). The data source now supports the `api` field and skips workspace-tracking when used at account level.
* Remove invalid `provider_config` attribute from account-only data sources `databricks_mws_workspaces` and `databricks_mws_credentials` ([#5664](https://github.com/databricks/terraform-provider-databricks/issues/5664)).

### Documentation

### Exporter

### Internal Changes

* Add `testframeworkV2/`, a multi-version Terraform test harness for the provider ([DESIGN.md](https://github.com/databricks/terraform-provider-databricks/blob/main/testframeworkV2/DESIGN.md)).

  The framework runs a single test definition (`test.yaml` + a directory of `*.tf` files) across N released provider versions plus a fresh `go build` of the current branch — without touching the developer's `~/.terraformrc`, `~/.databrickscfg`, or shell environment. Built around issue [#5672](https://github.com/databricks/terraform-provider-databricks/issues/5672)'s mission test (`testframeworkV2/issues-repro/issue_5672/`), which pins the regression-rollback-fix trajectory across `1.113.0` → `1.114.0` → `1.114.1` → `local` in four steps.

  Fixtures live under two trees: `testframeworkV2/issues-repro/issue_<N>/` for fixtures that reproduce a specific GitHub issue, and `testframeworkV2/tests/<descriptive-slug>/` for green-path / smoke / regression-guard fixtures not tied to a bug. Profile level (workspace / account / UC) is declared per-test via `requires.level`.

  `tfv2` auto-discovers the provider repo root (the `--repo` flag) by walking up from the working directory, and exposes every fixture as a `go test` subtest under `TestFixtures` (gated by `TFV2_RUN=1`) — so IDEs and CI can drive the framework without invoking the CLI.

  Quickstart:
  ```sh
  go build -o ~/.local/bin/tfv2 ./testframeworkV2/cmd/tfv2
  tfv2 run testframeworkV2/issues-repro/issue_5672/    # --repo auto-discovered
  ```

  Subcommands: `tfv2 run [-r] <dir>` (single or recursive), `tfv2 cache list/prune`, `tfv2 build local --repo <path>`. The framework lives in its own Go module (`testframeworkV2/go.mod`) so it can be built and run independently of the provider's transitive deps.
