# NEXT CHANGELOG

## Release v1.143.0

### Important Changes

* `databricks_app` now updates apps through the asynchronous Apps update API ([#6151](https://github.com/databricks/terraform-provider-databricks/pull/6151)).
  The synchronous update API ignored `compute_size`, `compute_min_instances`, and
  `compute_max_instances`, so changes to them never took effect; they are now applied. As
  before, an update replaces the app's updatable fields with the configuration, so fields
  omitted from it are cleared. That now also covers the fields the synchronous API skipped: an
  omitted `compute_size` resets to `MEDIUM`, omitted `compute_min_instances` and
  `compute_max_instances` disable horizontal scaling, and an omitted `git_source` clears the
  app's default git source. Updates now wait for the app update to reach `SUCCEEDED`.

### Breaking Changes

### New Features and Improvements
* Add resource and data sources for `databricks_assume_group_restricted_sharing_group`.
* Add `personal_access_token_wo` and `personal_access_token_wo_version` write-only attributes to `databricks_git_credential` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Add `options_wo` and `options_wo_version` write-only attributes to `databricks_connection` resource ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Bug Fixes
* Keep the configured values of the `pem_key_passphrase`, `ssl_keystore_key`, `ssl_key_password` and `schema_registry_api_secret` options of `databricks_connection`, which the API doesn't return, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).
* Ignore whitespace-only differences in PEM-encoded options of `databricks_connection`, such as `userProvidedServerCertificate`, to avoid a perpetual diff ([#6088](https://github.com/databricks/terraform-provider-databricks/pull/6088)).

### Documentation

### Exporter

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
