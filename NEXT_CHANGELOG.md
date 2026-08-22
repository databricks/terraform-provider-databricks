# NEXT CHANGELOG

## Release v1.129.0

### Important Changes

### Breaking Changes

### New Features and Improvements

### Bug Fixes

### Documentation

### Exporter

* Added support for exporting `databricks_data_classification_catalog_config` resource ([#5961](https://github.com/databricks/terraform-provider-databricks/pull/5961)).
* Added support for exporting `databricks_warehouses_default_warehouse_override` resource ([#5962](https://github.com/databricks/terraform-provider-databricks/pull/5962)).
* Added support for exporting `databricks_secret_uc` resource ([#5963](https://github.com/databricks/terraform-provider-databricks/pull/5963)).
* Skip system-managed jobs during export and add missing file references for job task parameters ([#5956](https://github.com/databricks/terraform-provider-databricks/issues/5956)).
* Added support for exporting `databricks_endpoint` resource ([#5951](https://github.com/databricks/terraform-provider-databricks/pull/5951)).
* Added support for exporting `databricks_environments_workspace_base_environment` and `databricks_environments_default_workspace_base_environment` resources ([#5960](https://github.com/databricks/terraform-provider-databricks/pull/5960)).
* Add an `exporter` dimension to the user agent ([#5954](https://github.com/databricks/terraform-provider-databricks/pull/5954)).
* Allow to generate named variables from references; introduce `databricks_account_id` variable for account-level exports; bug fixes ([#5952](https://github.com/databricks/terraform-provider-databricks/pull/5952)).
* Preserve zero `value` fields when exporting `databricks_workspace_setting_v2` and `databricks_account_setting_v2` ([#5955](https://github.com/databricks/terraform-provider-databricks/issues/5955)).
* Resolve references embedded in `databricks_cluster_policy` definitions instead of emitting hardcoded values ([#5953](https://github.com/databricks/terraform-provider-databricks/issues/5953)).

### Internal Changes

* Add `testframeworkV2/`, a multi-version Terraform test harness for the provider.

  The framework runs a single test definition (`test.yaml` + a directory of `*.tf` files) across N released provider versions plus a fresh `go build` of the current branch — without touching the developer's `~/.terraformrc`, `~/.databrickscfg`, or shell environment. Built around issue [#5672](https://github.com/databricks/terraform-provider-databricks/issues/5672)'s mission test (`testframeworkV2/issues-repro/issue_5672/`), which pins the regression-rollback-fix trajectory across `1.113.0` → `1.114.0` → `1.114.1` → `local` in four steps.

  Fixtures live under two trees: `testframeworkV2/issues-repro/issue_<N>/` for fixtures that reproduce a specific GitHub issue, and `testframeworkV2/tests/<descriptive-slug>/` for green-path / smoke / regression-guard fixtures not tied to a bug. Profile level (workspace / account / UC) is declared per-test via `requires.level`.

  `tfv2` auto-discovers the provider repo root (the `--repo` flag) by walking up from the working directory, and exposes every fixture as a `go test` subtest under `TestFixtures` (gated by `TFV2_RUN=1`) — so IDEs and CI can drive the framework without invoking the CLI. A `testframeworkV2/Makefile` wraps both entry points behind a `make test <path>` shortcut. Test specs can also assert against `terraform plan` stdout via `expect_non_empty_plan: true` and `plan_match: <regex>` for regressions that surface as a destructive plan diff rather than a non-zero exit. Plan-assertion failures now include the last 15 lines of plan stdout inline under the FAIL line — no manual `cat` of the per-step stdout log to see what the plan was.

  Quickstart:
  ```sh
  cd testframeworkV2/
  make test issues-repro/issue_5672/    # --repo auto-discovered
  ```

  Make targets: `make test <path>` (single fixture), `make test-all` (every fixture via `go test -run TestFixtures`), `make unit` (unit tests only — no cloud auth), `make build` (build `./tfv2`), `make clean`, `make help`. CLI subcommands: `tfv2 run [-r] <dir>`, `tfv2 cache list/prune`, `tfv2 build local --repo <path>`. The framework lives in its own Go module (`testframeworkV2/go.mod`) so it can be built and run independently of the provider's transitive deps.

  Shipped fixtures: `issues-repro/issue_5672` (the keystone mws_workspaces regression), `issues-repro/issue_5678` (catalog_workspace_binding force-replace on rollback), `issues-repro/issue_5668` (databricks_token validate, requires unassigned-SP profile), `tests/workspace_data_source_smoke` (data.databricks_mws_workspaces happy path), `tests/token_lifecycle_v2` (v2-mode demo: create/modify/destroy databricks_token with state assertions), and `tests/rollback-err` (regression-guard: v1.113.0 → v1.114.0 → v1.113.0 must not destructively replace databricks_token — the apply-then-downgrade sibling of #5678's apply-then-rollback-tag scenario).
