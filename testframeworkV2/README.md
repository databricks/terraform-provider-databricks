# testframeworkV2

A multi-version Terraform test harness for the Databricks provider. Define a
test once in YAML; run it across N released provider versions plus a fresh
build of the current branch — without touching the developer's
`~/.terraformrc`, `~/.databrickscfg`, or shell environment.

The framework was built to make issue [#5672][issue-5672] reliably
reproducible. The mission test in [`issues-repro/issue_5672/`](issues-repro/issue_5672)
is the keystone scenario: 4 steps that walk a release-rollback-fix
trajectory end-to-end. Additional regression fixtures live under
`issues-repro/` (one directory per GitHub issue); green-path / smoke
fixtures live under `tests/`.

[issue-5672]: https://github.com/databricks/terraform-provider-databricks/issues/5672

## Why

Authoring a regression test for "this bug fires on v1.114.0 but not on
v1.113.0 or v1.114.1" is awkward in vanilla Go test infrastructure: each
version requires a fresh terraform install, the user's `dev_overrides`
silently bypasses version pinning, and `os.Environ()` leaks
`DATABRICKS_HOST` into subprocesses. testframeworkV2 owns the whole
subprocess pipeline:

* one `filesystem_mirror` mirror per host, populated atomically from
  GitHub releases;
* per-run `.terraformrc` exported via `TF_CLI_CONFIG_FILE` (which
  *replaces* `~/.terraformrc`, not merges);
* per-step `_tfv2_versions_override.tf` pinning the
  `databricks/databricks` provider via Terraform's documented
  `*_override.tf` per-attribute merge — the user's `main.tf` is
  untouched;
* curated subprocess env: `DATABRICKS_*` from the developer's shell is
  stripped and replaced with the test.yaml's `profile`.

See [DESIGN.md](DESIGN.md) for the full rationale.

## Quickstart

```sh
# 1. Make sure you have a working terraform on PATH (or set
#    --terraform-bin / TFV2_TERRAFORM_BIN).
terraform -version

# 2. Add a profile to ~/.databrickscfg matching the test.yaml's `profile` field.
#    The framework reads the section's host to infer cloud / level for the
#    requires-skip-check, then sets DATABRICKS_CONFIG_PROFILE for the SDK.

# 3. Run the issue #5672 mission test.
cd testframeworkV2/
make test issues-repro/issue_5672/

# Expected:
# [PASS] step 1 (passes_on_1_113_0): 1.113.0      plan in 5.1s   no changes
# [PASS] step 2 (fails_on_1_114_0): 1.114.0       plan in 4.7s — failure-as-expected
# [PASS] step 3 (fixed_on_1_114_1): 1.114.1       plan in 4.6s   no changes
# [PASS] step 4 (fixed_on_local):   99.0.0-local  plan in 5.9s   no changes
# ----------------------------------------------------------
# issue_5672_...: PASS (4/4 steps passed in 22.4s)
# run dir: /Users/you/.testframeworkv2/runs/issue_5672_...-2026-05-08T08-15-00-a3f2
```

`make test <path>` is a thin wrapper around the underlying CLI. Full
target list:

```sh
make help        # usage banner
make test <path> # single fixture (recommended)
make test-all    # every fixture via 'go test' (TFV2_RUN=1)
make unit        # unit tests only — no cloud auth
make build       # build ./tfv2 binary
make clean       # remove the binary
```

Alternative invocation forms (same outcome, useful for IDEs / CI):

```sh
# Direct CLI — equivalent to `make test <path>` but skips Make:
go run ./cmd/tfv2 run issues-repro/issue_5672/

# `go test` — every fixture is also a Go subtest under TestFixtures
# (gated by TFV2_RUN=1 so plain `go test ./...` stays cheap):
TFV2_RUN=1 go test -run TestFixtures/issues-repro/issue_5672 -v ./...
```

`--repo` is auto-discovered by walking up from cwd looking for the
provider repo's `go.mod`; pass `--repo <path>` or set `TFV2_REPO` if
you're invoking from outside a checkout.

## Subcommands

```
tfv2 run <test-dir>             run a single test
tfv2 run -r <root>              recursively run every test.yaml under root
tfv2 cache list                 show cached provider versions
tfv2 cache prune                delete the provider cache
tfv2 build local --repo <path>  eagerly build local provider into cache
tfv2 version                    print version
tfv2 help                       show usage banner
```

## Running fixtures via `go test`

Every `test.yaml` under `issues-repro/` and `tests/` is also exposed as
a `go test` subtest, so IDEs and CI can drive the fixtures without
shelling out to `tfv2`:

```sh
cd testframeworkV2/
TFV2_RUN=1 go test -run TestFixtures -v ./...
```

`TFV2_RUN=1` is the gate — without it, `TestFixtures` skips (so plain
`go test ./...` stays cheap and doesn't fire real cloud-auth flows).
Each fixture runs as a separate `t.Run` subtest with the tree-
preserving 3-segment path (`<tree>/<fixture-dir>`), so IDEs render one
green/red dot per fixture and
`-run TestFixtures/issues-repro/issue_5672` filters to a single one.
See DESIGN.md §12.7 for the design rationale.

## test.yaml schema

```yaml
name: my_test                # required slug; matches ^[a-z0-9_-]+$
profile: ACCOUNT_AWS         # required; section name in ~/.databrickscfg
cleanup: true                # default true; final destroy with last-successful Apply step

requires:                    # skip-on-mismatch declarative gates
  cloud: any                 # aws | azure | gcp | any (default: any)
  level: account             # workspace | account | ucws | ucacct (default: workspace)

passthrough_env:             # optional; extra env vars forwarded to terraform
  - AWS_PROFILE              # NEVER include DATABRICKS_* — the profile field
  - GCP_PROJECT              #   is the only sanctioned auth channel.

steps:                       # required, ≥1
  - name: pre_regression
    version: "1.113.0"       # strict semver OR literal "local"
    command: plan            # plan | apply | destroy (default: apply)
    expect: success          # success | failure (default: success)
    timeout: 10m             # Go duration (default: 10m)

  - name: regression
    version: "1.114.0"
    command: plan
    expect: failure
    error_regex: 'failed to resolve workspace_id'
    # error_substring: 'literal'  # also supported; AND semantics with regex

  - name: forced_replace_check
    version: "1.113.0"
    command: plan                  # plan-content matchers below require plan
    expect: success                # and require expect: success
    expect_non_empty_plan: true    # plan stdout MUST NOT contain "No changes."
    plan_match: '# forces replacement'  # Go RE2 against plan stdout (multiline)
```

When `expect: failure`, **at least one of `error_substring` / `error_regex`
is required** — the framework rejects "fail in any way".

When asserting on plan content (e.g. "this rollback should produce a
forced-replacement diff"), use `expect_non_empty_plan` and/or
`plan_match`. Both fields require `command: plan` + `expect: success`;
both AND when combined. See DESIGN.md §17.10.

## Repository layout

```
testframeworkV2/
├── DESIGN.md                              ← detailed rationale (16 sections)
├── README.md                              ← this file
├── go.mod                                 ← separate module (minimal deps)
├── cmd/tfv2/                              ← CLI entry point
├── internal/
│   ├── providercache/                     ← provider zip cache + local-build
│   ├── tfrcwriter/                        ← .terraformrc + override file generation
│   ├── profile/                           ← ~/.databrickscfg parsing + cloud/level inference
│   ├── subprocenv/                        ← curated env-var allowlist
│   ├── terraform/                         ← binary discovery + version sanity check
│   ├── config/                            ← test.yaml schema + parse + validate
│   ├── runner/                            ← orchestration: parse → step loop → cleanup
│   └── result/                            ← per-step + per-run result types
├── issues-repro/                          ← fixtures that reproduce a specific GitHub issue
│   ├── issue_5672/                        ← keystone mws_workspaces regression test (account-level)
│   │   ├── test.yaml
│   │   └── main.tf
│   ├── issue_5678/                        ← catalog_workspace_binding force-replace on rollback
│   └── issue_5668/                        ← databricks_token validate (unassigned-SP profile required)
└── tests/                                 ← green-path / smoke / regression-guard fixtures (no specific issue)
    ├── workspace_data_source_smoke/       ← happy-path data.databricks_mws_workspaces
    │   ├── test.yaml
    │   └── main.tf
    ├── token_lifecycle_v2/                ← v2-mode demo: create/modify/destroy databricks_token
    └── rollback-err/                      ← regression-guard: v1.113 → v1.114 → v1.113 must not force-replace
```

Each `issues-repro/issue_<N>/` and `tests/<slug>/` directory is fully
self-contained: one `test.yaml` plus at least one `*.tf` file. Profile
level (workspace / account / UC) is declared per-test via
`requires.level`, not encoded in the directory tree.

## Runtime tree

The framework's working state lives at `~/.testframeworkv2/`:

```
~/.testframeworkv2/
├── providers/                             ← shared provider cache
│   └── registry.terraform.io/databricks/databricks/
│       ├── terraform-provider-databricks_1.113.0_darwin_arm64.zip
│       ├── terraform-provider-databricks_1.114.0_darwin_arm64.zip
│       └── 99.0.0-local/                  ← local builds (unpacked layout)
│           └── darwin_arm64/
│               ├── terraform-provider-databricks_v99.0.0-local
│               └── local-version.json     ← provenance (git SHA, dirty, etc.)
└── runs/                                  ← per-run workdirs (kept for debugging)
    └── <test>-<ts>-<rand>/
        ├── workdir/                       ← *.tf copied here; terraform runs here
        ├── plugins/                       ← TF_PLUGIN_CACHE_DIR (hardlink reuse)
        ├── step_1_passes_on_1_113_0.{stdout,stderr}.log
        ├── step_2_fails_on_1_114_0.{stdout,stderr}.log
        ├── step_3_fixed_on_1_114_1.{stdout,stderr}.log
        ├── step_4_fixed_on_local.{stdout,stderr}.log
        └── local-version.json             ← provenance copy (when local was used)
```

User source dirs are read-only from the framework's POV — every step works
out of `<run-dir>/workdir/` after a copy.

## CLI flags + env-var equivalents

| Flag | Env var | Purpose |
|---|---|---|
| `--terraform-bin <path>` | `TFV2_TERRAFORM_BIN` | override terraform binary discovery |
| `--cache-dir <path>` | `TFV2_CACHE_DIR` | override `~/.testframeworkv2/providers` |
| `--run-dir <path>` | — | override `~/.testframeworkv2/runs` |
| `--repo <path>` | `TFV2_REPO` | provider repo root for `version: local`. **Auto-discovered** when unset by walking up from cwd looking for the provider repo's go.mod (DESIGN.md §12.6). Required only when auto-discovery fails AND a step uses `version: local`. |
| `--no-cleanup` | `T_NO_CLEANUP=1` | skip final destroy regardless of test.yaml |
| `--verbose` | — | print framework debug logs |
| `-r`, `--recursive` | — | (run only) walk `<test-dir>` for nested test.yaml files |

Flags always win over env vars.

## Exit codes

| Code | Meaning |
|---|---|
| 0 | every step passed (or run was skipped per `requires`) |
| 1 | at least one step did not pass, or a framework error |
| 2 | usage error (bad flags, unknown subcommand, missing `<test-dir>`) |

SIGINT and SIGTERM cancel the runner's context — the in-flight step gets a
chance to abort cleanly via `tfexec`'s context propagation.

## What's intentionally not in v1

* `dev_overrides` support — the framework rejects them by design (they
  silently bypass version pins; see DESIGN.md §5 / G1).
* GPG verification of downloaded zips — accepted tradeoff with the
  filesystem_mirror approach (DESIGN.md §16/F4). v2 fodder.
* Cross-arch local builds — `local` builds for the host target only.
* Parallel test execution — v1 runs tests sequentially; v2 may add a flag.
* Retry on cleanup destroy failure — single attempt, loud-log on failure
  (DESIGN.md §10/G12).

## Contributing

Fixes and additions to the framework go through standard provider PR review.
For substantial behaviour changes, please update `DESIGN.md` first; the
design doc is the source of truth for "why does it behave this way?"

Each new test directory should:

1. live under `issues-repro/issue_<N>/` (when reproducing a specific GitHub
   issue) or `tests/<descriptive-slug>/` (when it's a green-path / smoke
   fixture not tied to an issue);
2. include `test.yaml` + at least one `*.tf` file in the same directory;
3. declare the target profile level via `requires.level` in `test.yaml`,
   not via directory placement (per v5.0 — see DESIGN.md §3 + §13 OQ3);
4. document the regression / behaviour the test pins, ideally with an
   issue or PR link.
