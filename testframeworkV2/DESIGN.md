# testframeworkV2 — Design Doc

**Status:** Draft v6.0 (implementer-tf11 + researcher-tf11, 2026-05-08)
**Audience:** tech-lead-tf11, reviewer-tf11, tester-tf11, implementer-tf11
**Scope:** Multi-step, multi-version Terraform integration test harness for `databricks/terraform-provider-databricks`. v1 ships small.
**Supersedes:** v5.0, v4.2, v4.1, v4, v3, v2, and v1.
**v6.0 deltas (vs v5.0):** v2-mode opt-in, per-step HCL + structured state assertions (Task #20). New §17 covering the full v2 schema + runner integration. Net behaviour change: tests in `issues-repro/` and `tests/` continue to work as v1 unless they declare `config:` on every step. v2 specs gain (a) per-step `*.tf` file swap (state survives the swap; runner wipes user `*.tf` and copies the per-step file before each `terraform init`), and (b) structured `assert:` blocks with resource presence + attribute equality checks against `terraform show -json` output. Implementation adds `internal/stateassert` package; extends `internal/config` (Step.Config + Step.Assert), `internal/result` (AssertionFailure type + StepResult.Assertions/AssertLog fields), and `internal/runner` (mode-aware prepareRun + per-step swapStepConfig + post-command runStateAssert). New fixture `tests/token_lifecycle_v2/` demonstrates the end-to-end create/modify/destroy lifecycle. v1 backward compat preserved — v1 specs see zero behaviour change; new fields are omitted-by-default in JSON output.
**v5.0 deltas (vs v4.2):** Phase 2 directory restructure to support multiple fixtures (Task #17). (1) `testframeworkV2/account/` → `testframeworkV2/issues-repro/` — the `account/`-as-level-marker convention from v4.x didn't scale to bug fixtures targeting different profile levels (issue #5678 is workspace-level, issue #5668 is account-level). The `issues-repro/` namespace groups all "this commit reproduces a known bug" fixtures regardless of level; per-test `requires.{cloud,level}` does the actual gating. (2) `testframeworkV2/tests/` added as a parallel sibling for green-path / smoke fixtures that aren't tied to a specific issue (e.g. `tests/<workspace-ds-smoke>/` covers the `data.databricks_mws_workspaces` happy path on the current branch). (3) `account/test1_issue_5672/` renamed to `issues-repro/issue_5672/` — drops the `test1_` prefix and aligns with the `issue_<N>` convention used by the new fixtures. (4) §3 directory layout block + §11 worked-example paths + §13 OQ3 updated; OQ3 now CLOSED with the issues-repro / tests split documented as the v5.0 convention. Pure structural change; no schema, runner, or behaviour deltas from v4.2.
**v4.2 deltas (vs v4.1):** picked up reviewer's 3 last-bounce items that the v4.1 sweep missed. (1) §4 schema example `requires.cloud: gcp` → `any` (matches the actual #5672 fixture). (2) §11 in-doc `test.yaml` example block (lines ~810-840) was a v3-era stale copy with `cloud: gcp` + GCP-specific regex + `error_substring` — synced fully with the actual fixture file. (3) §13 open question 2 ("issue #5672 cloud-specificity") marked CLOSED with reference to Appendix A "AWS account full-smoke" entry. Pure consistency cleanup; no behavior or schema changes from v4.1.
**v4.1 deltas (vs v4):** stale-reference sweep — 6 spots that still referenced the dropped preflight rule or the truncated "same SHA" framing. Pure consistency cleanup; no behavior or schema changes. (1) §4 schema example `passthrough_env` no longer lists items already in core (HTTP_PROXY etc.) — replaced with niche cross-cloud examples (AWS_PROFILE, AZURE_CLIENT_ID, GCP_PROJECT). (2) §7.0 pseudocode comment dropped the "validate no `databricks` pin" clause. (3) §8 synthetic-version subsection rewrote the "preflight guarantees no collision" sentence to cite override-merge per-attribute semantics instead. (4) §11 main.tf example block in DESIGN.md sample synced with the actual fixture (allow user pin). (5) test.yaml header comment about v1.114.1 expanded to mention binary-vs-git SHA distinction. (6) test.yaml step 3 inline comment same expansion.
**v4 deltas (vs v3):** mission test regex loosened to outer pattern `cannot populate provider_config for mws workspaces.*failed to resolve workspace_id` (cloud-portable — tester confirmed AWS reproduces with different inner error); `requires.cloud: any` in the #5672 sample (was `gcp`); env allowlist extended with SSL_CERT_FILE/DIR + HTTPS_PROXY/HTTP_PROXY/NO_PROXY (§10/G6); run-dir naming gains a 4-char random hex suffix to prevent same-second collisions; Appendix A renamed `expA` → `expA1`, added `B-COLLISION` and "AWS account full-smoke" entries; §11 walkthrough now shows both AWS and GCP inner-error variants; `error_substring` removed from the #5672 sample test.yaml (kept in schema as a forward-compat option for tests that want literal-byte matching).
**v3 deltas (vs v2):** override-merge pivot (preflight rejection removed — empirically validated by tester expA1/A2/A3 + B-COLLISION); explicit `plugin_cache_dir` hardlink optimization in `.terraformrc`; whole `.terraform/` wipe between steps (was just `providers/`); `passthrough_env` field promoted to v1 schema; v1.114.1 binary-SHA fact correction in §11/§16; tfexec.SetEnv replacement-semantics citation in §10/G6; new "Trade-offs accepted" subsection in §5; new Appendix A with tester's empirical evidence; G2 / open-question-1 closed.

---

## Table of contents
1. [Purpose & non-goals](#1-purpose--non-goals)
2. [Architecture overview](#2-architecture-overview)
3. [Directory layout](#3-directory-layout)
4. [`test.yaml` schema](#4-testyaml-schema)
5. [Generated `.terraformrc` template](#5-generated-terraformrc-template)
6. [Cache layout](#6-cache-layout)
7. [Step execution flow (pseudocode)](#7-step-execution-flow-pseudocode)
8. [Local-build flow](#8-local-build-flow)
9. [Why not `terraform-plugin-testing.ExternalProviders`?](#9-why-not-terraform-plugin-testingexternalproviders)
10. [Failure modes & gotchas](#10-failure-modes--gotchas)
11. [Worked example — `issue_5672`](#11-worked-example--issue_5672)
12. [Public API surface](#12-public-api-surface)
13. [Open questions](#13-open-questions)
14. [Out of scope (v1)](#14-out-of-scope-v1)
15. [Implementation milestones](#15-implementation-milestones)
16. [Footnotes](#16-footnotes)
17. [Appendix A — tester empirical evidence](#appendix-a--tester-empirical-evidence-highlights)

---

## 1. Purpose & non-goals

### Purpose
Run a sequence of `terraform plan`/`apply` steps against a single working directory where **each step pins a different provider version** (a released semver like `1.113.0` or the special string `local` for the current branch). State carries across steps, like running `terraform` from a CLI shell with `init -upgrade` between steps. Auth comes from a `~/.databrickscfg` profile (constant across all steps in a test).

### Mission test (acceptance criterion)
Reproduce GitHub issue [#5672](https://github.com/databricks/terraform-provider-databricks/issues/5672) end-to-end with **4 steps**:

| # | Step name | Provider | Expected | What it proves |
|---|---|---|---|---|
| 1 | `passes_on_1_113_0`  | `1.113.0` | PASS | Pre-regression baseline |
| 2 | `fails_on_1_114_0`  | `1.114.0` | FAIL (matches `cannot populate provider_config for mws workspaces.*failed to resolve workspace_id`) | Bug actually broke things. Outer pattern is cloud-portable — tester verified AWS + GCP both reproduce; the inner error differs by auth method (`Unable to load OAuth Config` on AWS, `strconv.ParseInt: parsing "": invalid syntax` on GCP), so we anchor on the outer wrapper. |
| 3 | `fixed_on_1_114_1`  | `1.114.1` | PASS | Rollback worked. The **git tag** `v1.114.1` points at the same source commit as `v1.113.0` (`7a6b469e`), but the released **binary** was independently rebuilt — different binary SHA, same source. See §16 F6. |
| 4 | `fixed_on_local`    | `local`   | PASS | Validates the framework's local-build pipeline. As of 2026-05-08 main has reverted PR #5667 + PR #5492 + related (see §16 F7), so the `local` source is currently equivalent to v1.113.0 for this code path. |

Step 4 exercises the `local` build pathway — the most interesting capability of the framework. It will graduate from "validates the pipeline" to "proves the regression-fix works" once the next forward-rolled fix lands on main; until then it passes for the same reason step 1 does. See §16/F7 for the revert-chain detail.

### Non-goals (v1)
- Parallel test execution. (Sequential only — tests share a single working dir per run; multiple tests can run sequentially in one invocation.)
- General Go-level `Check func(state)` assertions — `terraform plan` exit code + stderr regex covers the bug class we care about.
- Schema diffing or drift checks beyond what `terraform plan` does naturally.
- Mock/fake `databricks` profile — tests use real `~/.databrickscfg`. Auth-only validation is out of scope.
- Provider zip checksum verification (TLS to GitHub is sufficient for v1; documented as accepted tradeoff in §10).
- Cross-arch / cross-OS builds. Host arch only (`runtime.GOOS`/`runtime.GOARCH`).
- Windows. Darwin + Linux only in v1.
- Multi-provider tests (e.g., `databricks` + `aws` + `random`). v1 only manages `databricks` versions; other providers come from the public registry as normal.

---

## 2. Architecture overview

The framework is a thin orchestration layer between (a) a YAML-described multi-step test, (b) `terraform` invoked as a subprocess, and (c) a per-host filesystem cache of provider binaries. Three principles drive every design choice:

**(i) The user authors only HCL + YAML.** The user's `main.tf` describes resources/data sources/providers as if they'd written them for a real `terraform apply`, with one constraint: no `terraform { required_providers { databricks = ... } }` block — the framework owns version pinning. Auth is empty (`provider "databricks" { alias = "accounts" }`); the framework injects credentials via `DATABRICKS_CONFIG_PROFILE` env var. The `test.yaml` describes the step list, profile, expectations, and skip conditions.

**(ii) The subprocess is hermetic.** Every `terraform` invocation runs with a curated env (allowlist, not `os.Environ()` passthrough), a framework-controlled `TF_CLI_CONFIG_FILE`, no leaked `DATABRICKS_*` or `TF_LOG` from the developer's shell, no `~/.terraformrc` interference, and no auto-installed terraform binary. This is the only way to make tests reproducible across developer machines and CI — tester-tf11 confirmed empirically that any of these leaks silently bypass version pinning or change behavior.

**(iii) One installation method, one cache, one tfrc.** `filesystem_mirror` for both released versions (packed zip layout) and `local` builds (unpacked binary, synthetic version `99.0.0-local`). No `direct{}` for `databricks/databricks`; no `dev_overrides` anywhere; no `terraform-plugin-testing.ExternalProviders`. This rules out an entire class of "version-pinned but actually serving X" bugs.

```
┌──────────────────────────────────────────────────────────────────────────┐
│  $ tfv2 run testframeworkV2/issues-repro/issue_5672/                     │
│                                                                          │
│  ┌──────────┐   ┌────────────────────────────────────────┐               │
│  │ test.yaml│──▶│ config: parse, validate, skip-check    │               │
│  └──────────┘   └─────────┬──────────────────────────────┘               │
│  ┌──────────┐             │                                              │
│  │ main.tf  │──┐          ▼                                              │
│  └──────────┘  │   ┌───────────────────┐                                 │
│  ┌──────────┐  │   │ runner            │                                 │
│  │~/.dbcfg  │──┴──▶│   per step:       │                                 │
│  └──────────┘      │   1. ensure binary cached (download or build)       │
│                    │   2. write _tfv2_versions_override.tf               │
│                    │   3. wipe lock + .terraform/providers/              │
│                    │   4. tf init  (with curated env)                    │
│                    │   5. tf plan/apply (with curated env)               │
│                    │   6. assert exit + error_substring + error_regex    │
│                    │   final: tf destroy with last-successful version    │
│                    └─────┬─────────────────────────────────────┬─────────┘
│                          │                                     │         │
│                          ▼                                     ▼         │
│                  ┌──────────────┐                     ┌──────────────┐   │
│                  │ providercache │                    │ tfexec driver│   │
│                  │   download   │◀── http GET ──┐    │ (CLI wrapper)│   │
│                  │   build      │◀── go build ──┤    │              │   │
│                  └──────┬───────┘                │    └──────┬───────┘   │
│                         ▼                         │           ▼           │
│        ~/.testframeworkv2/providers/              │   /usr/local/bin/    │
│          registry.terraform.io/databricks/        │     terraform        │
│            databricks/                            │     (from $PATH)     │
│              <ver>/<os>_<arch>/...                │                      │
│              terraform-provider-databricks_       │                      │
│                <ver>_<os>_<arch>.zip              │                      │
│                                                                          │
│        ~/.testframeworkv2/runs/<test>-<ts>-<rand>/                              │
│          workdir/                                                         │
│            main.tf (copied)                                              │
│            _tfv2_versions_override.tf (regenerated per step)             │
│            .terraformrc (TF_CLI_CONFIG_FILE)                             │
│            terraform.tfstate (preserved across steps)                    │
│          step-1.{stdout,stderr}.log                                       │
│          step-2.{stdout,stderr}.log ...                                  │
│          local-version.json (when local was used)                        │
│          run.json (final manifest)                                       │
└──────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Directory layout

### Source tree (`testframeworkV2/` in the repo)
```
testframeworkV2/
├── DESIGN.md                              ← this file
├── README.md                              ← user quickstart, after design lands
├── go.mod                                 ← separate module — see §12
├── .gitignore                             ← .terraform/, .terraform.lock.hcl, _tfv2_*.tf, *.tfstate*
├── cmd/
│   └── tfv2/
│       └── main.go                        ← CLI entry point
├── internal/
│   ├── runner/
│   │   ├── runner.go                      ← orchestration: run, step, cleanup
│   │   ├── step.go                        ← single-step execution
│   │   └── runner_test.go
│   ├── providercache/
│   │   ├── cache.go                       ← Get(version, target) → packed-zip path
│   │   ├── download.go                    ← GitHub release fetch (atomic write)
│   │   ├── localbuild.go                  ← go build → unpacked binary
│   │   └── cache_test.go
│   ├── tfrcwriter/
│   │   ├── tfrc.go                        ← generates TF_CLI_CONFIG_FILE content
│   │   ├── overrides.go                   ← writes _tfv2_versions_override.tf
│   │   └── tfrc_test.go
│   ├── profile/
│   │   ├── profile.go                     ← reads ~/.databrickscfg, identifies cloud + level
│   │   └── profile_test.go
│   ├── subprocenv/
│   │   ├── env.go                         ← curated env allowlist for `terraform` subprocess
│   │   └── env_test.go
│   ├── terraform/
│   │   └── locate.go                      ← resolves terraform binary, sanity-checks version
│   ├── config/
│   │   ├── config.go                      ← test.yaml schema + parse + validate
│   │   └── config_test.go
│   └── result/
│       └── result.go                      ← per-step + per-run result types, summary printer
├── issues-repro/                          ← fixtures that reproduce known bugs (one dir per issue)
│   ├── issue_5672/                        ← #5672 mws_workspaces account provider_config regression
│   │   ├── test.yaml
│   │   └── main.tf
│   ├── issue_5678/                        ← researcher-tf11 output (Phase 2)
│   │   ├── test.yaml
│   │   └── main.tf
│   └── issue_5668/                        ← researcher-tf11 output (Phase 2)
│       ├── test.yaml
│       └── main.tf
└── tests/                                 ← green-path / smoke fixtures NOT tied to a specific issue
    └── workspace_data_source_smoke/       ← happy-path data.databricks_mws_workspaces on local build
        ├── test.yaml
        └── main.tf

# Note on directory naming (v5.0):
# Two top-level fixture trees, distinguished by intent rather than profile level:
#   issues-repro/  — each subdir reproduces a SPECIFIC GitHub issue. Naming
#                    convention: issue_<N>/ where N is the issue number.
#                    Profile level (workspace / account / UC) varies per test;
#                    each test.yaml's `requires.level` does the actual gating.
#   tests/         — green-path / regression-guard fixtures NOT tied to a specific
#                    bug. Naming convention: descriptive slug (e.g.
#                    workspace_data_source_smoke/). Same `requires.level` gating.
#
# v4.x used `account/` / `workspace/` / `ucws/` / `ucacct/` as the level marker via
# directory placement. This didn't scale: the new fixtures span multiple levels
# (#5678 is workspace, #5668 is account, the smoke is workspace), and forcing
# bug-fixture organisation through profile-level dirs split related issues across
# different trees. See §13 OQ3 (CLOSED) for the rationale.
```

### Runtime tree (created by the framework)
```
~/.testframeworkv2/                        # framework state, not in repo
├── providers/                             ← shared provider cache (across runs and tests)
│   └── registry.terraform.io/
│       └── databricks/
│           └── databricks/
│               ├── terraform-provider-databricks_1.113.0_darwin_arm64.zip   ← packed
│               ├── terraform-provider-databricks_1.114.0_darwin_arm64.zip   ← packed
│               ├── terraform-provider-databricks_1.114.1_darwin_arm64.zip   ← packed
│               └── 99.0.0-local/
│                   └── darwin_arm64/                                          ← unpacked
│                       └── terraform-provider-databricks_v99.0.0-local
└── runs/                                  ← per-run work dirs (preserved on disk for debugging)
    └── issue_5672_mws_workspaces_account_provider_config_regression-2026-05-07T20-15-00-a3f2/
        ├── workdir/
        │   ├── main.tf                    ← copied from source dir
        │   ├── _tfv2_versions_override.tf ← regenerated each step
        │   ├── .terraformrc               ← framework-generated, points TF_CLI_CONFIG_FILE here
        │   ├── .terraform/                ← per-step (wiped between steps)
        │   ├── .terraform.lock.hcl        ← per-step (deleted before each init)
        │   └── terraform.tfstate          ← preserved across steps
        ├── step-1-passes_on_1_113_0.stdout.log
        ├── step-1-passes_on_1_113_0.stderr.log
        ├── step-2-fails_on_1_114_0.stdout.log
        ├── step-2-fails_on_1_114_0.stderr.log
        ├── step-3-fixed_on_1_114_1.stdout.log
        ├── step-3-fixed_on_1_114_1.stderr.log
        ├── step-4-fixed_on_local.stdout.log
        ├── step-4-fixed_on_local.stderr.log
        ├── local-version.json             ← present when `local` was used (provenance)
        └── run.json                       ← final result manifest (one record per step)
```

**Why two trees?** The provider cache is shared across all runs and tests on the host (downloaded once, reused forever for released versions). Per-run workdirs are throwaway but preserved on disk for forensic debugging. User's source tree (`testframeworkV2/issues-repro/issue_5672/`) is never written to by the framework.

---

## 4. `test.yaml` schema

```yaml
# REQUIRED FIELDS
name: issue_5672_mws_workspaces_account_provider_config_regression
profile: ACCOUNT_GCP                    # ~/.databrickscfg section name; constant across all steps

# OPTIONAL TOP-LEVEL FIELDS
cleanup: true                           # default true; final destroy with last-successful Apply step's version
                                        # (no-op when no Apply steps ran). Override globally with T_NO_CLEANUP=1.

requires:                               # skip-on-mismatch declarative gates
  cloud: any                            # one of: aws, azure, gcp, any  (default: any)
  level: account                        # one of: workspace, account, ucws, ucacct  (default: workspace)

passthrough_env:                        # optional opt-in: extra env vars forwarded to terraform subprocess.
  - AWS_PROFILE                         # Strict allowlist in §10 G6 covers basics, locale, TLS/CA, proxy, framework.
  - AZURE_CLIENT_ID                     # Use passthrough_env for niche cross-cloud needs that DON'T fit the core list.
  - GCP_PROJECT                         # DO NOT add DATABRICKS_* here — that defeats the profile mechanism.

steps:                                  # required, ≥1
  - name: passes_on_1_113_0             # required, slug
    version: "1.113.0"                  # required: semver string OR literal "local"
    command: plan                       # required: plan | apply | destroy  (default: apply)
    expect: success                     # required: success | failure       (default: success)
    timeout: 10m                        # default 10m per step (Go duration)

  - name: fails_on_1_114_0
    version: "1.114.0"
    command: plan
    expect: failure
    # Both `error_substring` (literal) and `error_regex` (Go RE2) are supported in the schema —
    # AND semantics when both set, ≥1 required when expect=failure. For #5672 we use only
    # error_regex anchored on the cloud-portable outer pattern; the inner error varies by
    # auth method (AWS OAuth vs GCP GSA), so a substring match would pin us to one cloud.
    # Example of both-set form (from a hypothetical workspace-level test):
    #   error_substring: 'precondition failed'
    #   error_regex: 'Error:.*precondition failed.*resource X'
    error_regex: 'cannot populate provider_config for mws workspaces.*failed to resolve workspace_id'

  - name: fixed_on_1_114_1
    version: "1.114.1"
    command: plan
    expect: success

  - name: fixed_on_local
    version: "local"                    # framework runs `go build` from --repo / pwd
    command: plan
    expect: success
```

### Schema rules (validated by `internal/config`)

- `name` (test): non-empty, slug (matches `^[a-z0-9_-]+$`).
- `profile`: required; must exist in `~/.databrickscfg` at parse time (eager validation).
- `requires.cloud`: enum; default `any`.
- `requires.level`: enum; default `workspace`.
- `steps`: ≥1.
- `steps[].name`: slug, unique within test.
- `steps[].version`: either a strict semver (`X.Y.Z` or `X.Y.Z-prerelease`) OR exactly the string `local`.
- `steps[].command`: enum {`plan`, `apply`, `destroy`}; default `apply`.
- `steps[].expect`: enum {`success`, `failure`}; default `success`.
- `steps[].error_substring` / `error_regex`: only meaningful when `expect: failure`. **At least one is required** when `expect: failure` — validation fails at parse time if neither is set.
  - `error_substring` = literal byte match against captured stderr (case-sensitive).
  - `error_regex` = Go `regexp` (RE2), evaluated multi-line (`(?s)` flag implicit so `.` matches newlines).
  - When both set: AND semantics — both must match. Stderr-only (§7).
- `steps[].timeout`: Go duration; default 10m.
- `passthrough_env` (top-level): optional list of env var names; the framework adds these to the curated subprocess env (§10 G6). Names that don't exist in the parent environment are silently dropped. NEVER include `DATABRICKS_*` here (that defeats the profile mechanism — use the `profile` field).

### Pre-flight validation (run before any step)
- **Profile existence**: validate that the named `profile` is a section in `~/.databrickscfg` (section-existence check only — do NOT parse cloud-specific fields like `host` / `account_id` / `google_service_account`; defer field parsing to the SDK at terraform-run time).
- **Step name uniqueness**: each `steps[].name` must be unique within the test.
- **Version syntax**: each `steps[].version` is either `local` or matches a strict semver pattern.
- **Regex validity**: each `error_regex` compiles as a Go RE2.
- **Failure-assertion completeness**: when `expect: failure`, **at least one of `error_substring` / `error_regex` must be set** — fail validation with a clear message; we don't accept "fail in any way".
- **`requires` enum validity**: cloud ∈ {aws, azure, gcp, any}; level ∈ {workspace, account, ucws, ucacct}.

**Note on user HCL — no preflight check.** Per tester's empirical confirmation (Appendix A: experiment expA3), Terraform's `*_override.tf` semantics correctly merge per-attribute even for nested `required_providers.<name>.version`. The framework's `_tfv2_versions_override.tf` always wins on the `version` attribute, regardless of what (if anything) the user authored in `main.tf`. Therefore: **users may freely include their own `terraform { required_providers { databricks = { version = "..." } } }` block in `main.tf`** for IDE / standalone-`terraform plan` workflows; the framework's override transparently takes over the version field at test time. No preflight rejection of user HCL is performed. (See §5 for the override mechanism, §10 G2 for empirical evidence, §13 open Q for one residual edge case.)

---

## 5. Generated `.terraformrc` template

The framework writes one `.terraformrc` per run into the run's workdir, then exports `TF_CLI_CONFIG_FILE=<workdir>/.terraformrc` for every `terraform` subprocess. This file is the only `terraform` CLI config the subprocess sees.

```hcl
# Generated by testframeworkV2. Do not edit.
# Run: issue_5672_mws_workspaces_account_provider_config_regression-2026-05-07T20-15-00-a3f2

provider_installation {
  filesystem_mirror {
    path    = "/Users/tanmay.rustagi/.testframeworkv2/providers"
    include = ["registry.terraform.io/databricks/*"]
  }

  direct {
    exclude = ["registry.terraform.io/databricks/*"]
  }
}

# Hardlink optimization: when plugin_cache_dir is set alongside filesystem_mirror,
# Terraform hardlinks (same inode) the resolved provider into .terraform/providers/...
# instead of copying the binary. Tester confirmed empirically (same inode 577301870
# across init invocations). Within-run reuse across the 4 step inits is essentially free.
# Per-run cache (not shared across runs) preserves test isolation.
plugin_cache_dir = "/Users/tanmay.rustagi/.testframeworkv2/runs/<run-id>/plugins"
```

### Why exactly this shape

- **Narrow `include`** — `registry.terraform.io/databricks/*`, NOT `*/*`. If a user's `main.tf` declares `hashicorp/google` (issue #5672's repro does), it must come from the public registry. A wide `*/*` include routes everything through the local mirror and breaks every other provider in the user's HCL with the runtime error `provider hashicorp/google was not found in any of the search locations`. Tester confirmed this empirically (test3-narrow vs test3-broad). **Do NOT widen the `include` pattern.**
- **`direct { exclude = [...] }`** — explicit dual: `databricks` from mirror, everything else from registry direct. The `exclude` list MUST mirror the `include` narrowness exactly, or we end up routing some providers neither through mirror nor direct (terraform errors with "no available release"). Without an explicit `direct{}` block, terraform's default fallback is registry-direct for all — and our mirror would only be consulted as a fallback (wrong direction).
- **Mirror path is absolute** — relative paths are resolved relative to the working directory of the `terraform` subprocess at runtime, which is the per-run workdir. Absolute keeps mirror lookup deterministic.
- **`plugin_cache_dir` is per-run, not shared** — preserving test isolation. Within a single run (4 init invocations across 4 steps), terraform hardlinks resolved provider files instead of re-extracting from the zip each time. F2 in TL's binding constraints. Free speed; no cross-run leakage.
- **No `dev_overrides`** — explicit. If the developer's `~/.terraformrc` has a `dev_overrides` block, our `TF_CLI_CONFIG_FILE` env var **replaces** that file entirely (verified — terraform docs and tester); we don't merge. This is the entire point of B1.

### Override file mechanism — `_tfv2_versions_override.tf`

The framework writes per-step a file named `_tfv2_versions_override.tf` into the workdir:
```hcl
terraform {
  required_providers {
    databricks = {
      source  = "databricks/databricks"
      version = "= 1.114.0"
    }
  }
}
```

**Filename rules** (B13):
- Filename **MUST** end in `_override.tf`. Terraform recognizes the suffix and applies per-attribute override-merge. A non-override `*.tf` file containing the same `terraform {}` block would clash with the user's `terraform {}` block at parse time ("duplicate `terraform` block").
- The `_tfv2_` prefix is a framework convention to namespace generated files (avoids collision with any user-authored `*_override.tf` file).

**Empirically validated semantics** (Appendix A — tester's expA3):
- User's `main.tf` has `terraform { required_providers { databricks = { source = "databricks/databricks", version = ">= 1.0.0" }, google = { source = "hashicorp/google", version = ">= 5.0.0" } } }`.
- Framework writes `_tfv2_versions_override.tf` containing only `terraform { required_providers { databricks = { source = "databricks/databricks", version = "= 1.114.0" } } }`.
- Resulting `.terraform.lock.hcl`: databricks @ `1.114.0` (override won on the `version` field) AND google @ `7.31.0` (preserved untouched — override didn't redeclare it).
- **Override merges per-attribute, even for nested `required_providers.<name>.<field>`. Other providers in `required_providers` are left untouched.**

This means **users may freely include their own `terraform { required_providers { databricks = { ... } } }` block** — the framework's override transparently takes over the `version` attribute at test time. Users get standalone `terraform plan` workflows for IDE/dev (with their own pin) AND framework-driven multi-version regression tests (with framework's override). No preflight rejection needed.

**Future capability (v2 fodder):** the same override mechanism can inject `provider "databricks" { host = ..., google_service_account = ... }` blocks at test time, resolved from the profile. This would enable "test-time auth injection" for tests that want to assert against multiple profiles in sequence — useful for cross-profile regression scenarios. Out of v1 scope.

### Trade-offs accepted

The design makes the following deliberate tradeoffs in v1, all empirically validated by tester. None block correctness; none should surprise reviewers of the implementation.

1. **`(unauthenticated)` mirror warning.** `terraform init` from a `filesystem_mirror` emits warnings like `terraform.lock.hcl is missing the h1: hash and …` and the lock entries carry `(unauthenticated)` markers. This is the price of GPG-free mirror operation. Acceptable: the framework owns binary acquisition end-to-end (downloads from `github.com/databricks/...` over TLS, atomically writes to host-local cache); GPG would be belt-and-suspenders. v2 may add `_SHA256SUMS` validation. See §16 F4.
2. **Single-platform `h1:` hashes in lock files.** A lock file generated on darwin contains only darwin's `h1:` hash; running on linux fails init. Acceptable: locks are per-run, ephemeral, gitignored, regenerated each step. We never depend on lock-file portability. See G4.
3. **Override semantics for nested `required_providers`.** We rely on Terraform's documented `*_override.tf` per-attribute merge — empirically validated by tester (Appendix A) but treated as a stable interface contract. If Terraform ever changes those semantics in a future major version, the framework could fall back to a "stricter preflight + non-override file" mode. v1 does not implement that fallback because tests showed override merge works on terraform 1.5.7 (and is documented behavior in current docs).
4. **`local` build is host-arch only.** No cross-arch/cross-OS builds in v1. A test with `version: local` on a darwin/arm64 host produces a darwin/arm64 binary; the test cannot be run on linux/amd64 in the same invocation. Acceptable for ad-hoc dev; CI matrices can run independently per arch. v3.
5. **Provider cache grows monotonically.** No automatic eviction in v1. `tfv2 cache prune` is the manual safety valve. Disk usage estimate: ~65 MB per (version × target). Tens of versions across multiple targets can exceed 1 GB; not a v1 problem.

---

## 6. Cache layout

### Released versions — packed zip layout

Verified path (downloaded v1.113.0_darwin_arm64.zip during research; HEAD = 200; binary inside is `terraform-provider-databricks_v1.113.0`):
```
~/.testframeworkv2/providers/
└── registry.terraform.io/
    └── databricks/
        └── databricks/
            ├── terraform-provider-databricks_1.113.0_darwin_arm64.zip
            ├── terraform-provider-databricks_1.114.0_darwin_arm64.zip
            └── terraform-provider-databricks_1.114.1_darwin_arm64.zip
```

This is the `filesystem_mirror` packed layout. Terraform discovers versions by scanning files matching `terraform-provider-<TYPE>_<VERSION>_<TARGET>.zip` and extracts to `.terraform/providers/...` during `init`.

**Why packed over unpacked for released versions:**
- 1:1 with the GitHub release asset — download = save = done.
- `_SHA256SUMS` from the release applies to the zip (future v2: verify).
- Trade-off: terraform unzips during init (~1s overhead). Acceptable.

### Local-build — unpacked layout

The local build cannot be a zip without an extra wrap step, so we use the unpacked filesystem_mirror layout for it:
```
~/.testframeworkv2/providers/
└── registry.terraform.io/
    └── databricks/
        └── databricks/
            └── 99.0.0-local/
                └── darwin_arm64/
                    └── terraform-provider-databricks_v99.0.0-local
```

Terraform docs confirm: a single mirror can mix packed and unpacked layouts on a per-version basis. Released versions stay packed; the synthetic local version is unpacked.

**Why not also pack the local build?** Wrapping `go build` output in a zip is 5 lines of Go, but the unpacked path is closer to how `go run` / dev workflows feel and avoids spurious diff noise (zip mtime/permissions). Negligible perf difference.

### Cache atomicity

Every cache write goes through `<path>.partial → fsync → atomic rename`. Two parallel runs racing on the same uncached version both download to distinct `.partial` files; whichever lands first wins via rename, the other overwrites identically. No flock needed.

### Eviction
None in v1. Cache grows monotonically. Sub-command `tfv2 cache prune` (added in M6) clears it manually.

---

## 7. Step execution flow (pseudocode)

```go
// runner.Run is the entry point for a parsed test.
func (r *Runner) Run(ctx context.Context) (RunResult, error) {
    // 7.0 Validate, skip-check, prepare run dir.
    if !r.requiresMatchHostEnv() {
        return RunResult{Skipped: true, Reason: "requires.cloud/level mismatch"}, nil
    }
    runDir := r.makeRunDir()              // ~/.testframeworkv2/runs/<test>-<ts>-<rand>/
    workdir := filepath.Join(runDir, "workdir")
    r.copyUserHCLInto(workdir)            // copy *.tf from source dir; preserve perms (no HCL validation — override-merge wins; see §4)
    r.writeTerraformRC(workdir)           // §5 template, absolute mirror path
    env := subprocenv.Build(r.profile)    // §10/E2 — explicit allowlist

    var lastSuccessfulApply *Step         // for cleanup
    results := []StepResult{}

    // 7.1 For each step.
    for i, step := range r.spec.Steps {
        // 7.1.a Ensure binary cached.
        zipOrBinPath, syntheticVer := r.cache.Resolve(ctx, step.Version, hostTarget)

        // 7.1.b Write per-step _tfv2_versions_override.tf.
        ver := step.Version
        if step.Version == "local" { ver = syntheticVer /* "99.0.0-local" */ }
        tfrcwriter.WriteVersionsOverride(workdir, ver)

        // 7.1.c Wipe lock file AND the entire .terraform/ directory (B2).
        // Order matters: rm -f .terraform.lock.hcl, then rm -rf .terraform.
        // The state file (terraform.tfstate) lives outside .terraform/ at workdir root
        // and is explicitly NOT touched — preserved across all steps.
        os.Remove(filepath.Join(workdir, ".terraform.lock.hcl"))
        os.RemoveAll(filepath.Join(workdir, ".terraform"))

        // 7.1.d terraform init.
        tf := tfexec.NewTerraform(workdir, terraformBin)
        tf.SetEnv(env)                                     // curated; no os.Environ() leak
        var initErr = tf.Init(ctx)
        // (init logs to step-N.{stdout,stderr}.log via tf.SetStdout/Stderr to log files)

        // 7.1.e terraform <command>.
        var cmdErr error
        switch step.Command {
        case "plan":   _, cmdErr = tf.Plan(ctx)            // hasChanges discarded; we only care about error
        case "apply":  cmdErr = tf.Apply(ctx, tfexec.AutoApprove(true))
        case "destroy": cmdErr = tf.Destroy(ctx, tfexec.AutoApprove(true))
        }

        // 7.1.f Assert.
        stepResult := assertStep(step, initErr, cmdErr, capturedStderr)
        results = append(results, stepResult)

        // 7.1.g Track for cleanup.
        if stepResult.Status == StatusPass && step.Command == "apply" {
            lastSuccessfulApply = &step
        }

        // 7.1.h Fail-fast: do we continue on step failure?
        // v1: continue regardless. The point of the framework is multi-step regression tests
        // where a middle step fails-as-expected; downstream steps still run. If the user wants
        // hard-stop on unexpected failure, that's a v2 flag.
    }

    // 7.2 Cleanup.
    if r.spec.Cleanup && os.Getenv("T_NO_CLEANUP") != "1" && lastSuccessfulApply != nil {
        // Re-establish the last good provider version, then destroy.
        zipPath, _ := r.cache.Resolve(ctx, lastSuccessfulApply.Version, hostTarget)
        tfrcwriter.WriteVersionsOverride(workdir, lastSuccessfulApply.Version)
        os.Remove(filepath.Join(workdir, ".terraform.lock.hcl"))
        os.RemoveAll(filepath.Join(workdir, ".terraform"))
        tf := tfexec.NewTerraform(workdir, terraformBin); tf.SetEnv(env)
        if err := tf.Init(ctx); err != nil {
            log.Errorf("cleanup init failed: %v — leaving resources for manual cleanup", err)
        } else if err := tf.Destroy(ctx, tfexec.AutoApprove(true)); err != nil {
            log.Errorf("cleanup destroy failed: %v — leaving resources for manual cleanup", err)
        }
        // No retry loop — destroy failures are loud-logged but never fatal to the framework.
    }

    // 7.3 Persist run.json.
    return RunResult{Steps: results, RunDir: runDir}, nil
}
```

### Assertion semantics (`assertStep`)

```go
// expect=success
if step.Expect == Success {
    if initErr != nil || cmdErr != nil {
        return StepResult{Status: Fail, Reason: "expected success, got error: ..."}
    }
    return StepResult{Status: Pass}
}

// expect=failure
if step.Expect == Failure {
    if cmdErr == nil {
        return StepResult{Status: Fail, Reason: "expected failure, but command succeeded"}
    }
    if step.ErrorSubstring != "" && !bytes.Contains(stderr, []byte(step.ErrorSubstring)) {
        return StepResult{Status: Fail, Reason: "stderr did not contain '...': got <stderr>"}
    }
    if step.ErrorRegex != nil && !step.ErrorRegex.Match(stderr) {
        return StepResult{Status: Fail, Reason: "stderr did not match /.../: got <stderr>"}
    }
    return StepResult{Status: Pass /* failure-as-expected */}
}
```

### Streams matched

`error_substring` and `error_regex` are matched against **stderr only**. Terraform writes errors to stderr (verified: `Error:` blocks and stack traces go to fd 2; `terraform plan -no-color` output to fd 1). Matching stderr-only avoids false positives from `output` blocks containing the literal error text. If a future test ever needs to match stdout, we add an explicit `stream: stderr|stdout|combined` knob.

### Plan-only steps

Issue #5672 reproduces during `ReadDataSource` at `terraform plan` (per tester B9). All 4 mission-test steps use `command: plan`. This means:
- No state is mutated by any step.
- Cleanup destroy is a no-op (no Apply steps → `lastSuccessfulApply == nil` → cleanup skipped per §7.2).
- Total test wall time is fast — no real workspace mutation.

For tests that DO need apply (e.g., a test that creates a resource then asserts schema migration on a different version), the same flow works — cleanup destroy runs against the last successful Apply step's version.

---

## 8. Local-build flow

```go
// providercache.localbuild.Build is invoked when a step's version == "local".
func Build(ctx context.Context, repoRoot string, target Target, cacheRoot string) (binPath, syntheticVer string, err error) {
    syntheticVer = "99.0.0-local"
    binDir := filepath.Join(cacheRoot,
        "registry.terraform.io/databricks/databricks",
        syntheticVer,
        target.String())                               // e.g. "darwin_arm64"
    binPath = filepath.Join(binDir, "terraform-provider-databricks_v"+syntheticVer)

    if err := os.MkdirAll(binDir, 0o755); err != nil { return "", "", err }

    // 8.1 Capture provenance BEFORE building.
    sha, dirty := gitState(repoRoot)                   // exec git rev-parse + git status --porcelain
    provenance := LocalVersion{
        SyntheticVersion: syntheticVer,
        GitSHA:           sha,
        Dirty:            dirty,
        BuiltAt:          time.Now().UTC(),
        GoVersion:        goVersion(),
        OSArch:           target.String(),
    }

    // 8.2 Build.
    cmd := exec.CommandContext(ctx, "go", "build",
        "-o", binPath,
        "./",                                           // build the root package — produces the provider binary
    )
    cmd.Dir = repoRoot
    cmd.Env = append(os.Environ(),                      // for `go build` only — NOT terraform subprocess
        "GOOS="+target.OS, "GOARCH="+target.Arch,
        "CGO_ENABLED=0",                                // matches goreleaser config
    )
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr       // surface compile errors directly
    if err := cmd.Run(); err != nil { return "", "", fmt.Errorf("go build: %w", err) }

    // 8.3 Persist provenance into run dir AND alongside the binary.
    writeJSON(filepath.Join(binDir, "local-version.json"), provenance)

    return binPath, syntheticVer, nil
}
```

### Cache strategy: rebuild every step

Per TL clarification (point 2 in user's must-address list): **rebuild every run is simplest**. Rationale:

- Source-tree hash caching is easy to get wrong (vendor changes, generated files, dirty worktree, untracked `.go` files).
- `go build` is the right primitive for "is anything different?" — Go's incremental cache makes a no-op rebuild ~250-500ms.
- A fresh build per step is dead simple to explain. Output in step logs is unambiguous.

**Force-rebuild env var (forward-compat hook):** `TFV2_FORCE_REBUILD=1` is reserved. v1 rebuilds always, so this is a no-op. v2 might add `--cache-local` flag with `TFV2_FORCE_REBUILD=1` as the override.

### Provenance file (`local-version.json`)

Captured per step that uses `local`:
```json
{
  "synthetic_version": "99.0.0-local",
  "git_sha": "1f62e20c403b5e52febce13fc149be56f2c19f38",
  "dirty": true,
  "built_at": "2026-05-07T20:15:42Z",
  "go_version": "go1.25.7",
  "os_arch": "darwin_arm64"
}
```

Two copies are written:
1. Next to the cached binary at `~/.testframeworkv2/providers/.../99.0.0-local/<target>/local-version.json` — overwritten each rebuild.
2. Into the run dir at `~/.testframeworkv2/runs/<test>-<ts>-<rand>/local-version.json` — preserved with that run's logs.

This makes test results reproducible: if a `local` step passes today, you can later check git SHA + dirty flag from the run dir.

### Synthetic version constraint

The override file uses `version = "99.0.0-local"` (exact equality). Per Terraform's semver matching:
- `99.0.0-local` is a prerelease.
- Operators `>=`, `<=`, `~>` do NOT match prereleases.
- `=` (exact equality) matches.

The framework writes `version = "99.0.0-local"` — exact equality, no operator. This is safe even when user `main.tf` has its own `databricks` version pin: per `*_override.tf` per-attribute merge semantics (G2 / Appendix A expA3 + B-COLLISION), the framework's exact-version always wins on `databricks.version`, transparently replacing whatever operator the user wrote (`>=`, `~>`, `=`, etc.).

---

## 9. Why not `terraform-plugin-testing.ExternalProviders`?

Reviewer raised this as a potential ~50-line glue alternative. We rejected it. Reasons in priority order:

### 9.1 Cannot handle `local` builds (the most interesting capability)

`ExternalProviders{Source, VersionConstraint}` only resolves from a public registry mirror — there is no API to point at a local binary. Issue #5672's mission test has 4 steps; step 4 (`local`) is the only one that exercises the local-build pipeline (and, when a forward-rolled fix is on main, the only one that proves that fix works — see §16/F7 for the current revert-chain caveat). We cannot ship a framework that requires the regression test to lop off its most important step. v2 of the framework would need to implement filesystem_mirror anyway for `local` — at which point we're maintaining two version-resolution paths.

### 9.2 All-or-nothing constraint per `TestCase`

Per plugin-testing godoc: "When providers are specified at the TestStep level, all TestStep within a TestCase must declare providers." So we cannot mix `ExternalProviders` (steps 1-3) with an in-process or filesystem-mirror provider (step 4). Either all four steps go through ExternalProviders (impossible — see 9.1) or none do.

### 9.3 Registry-online required

`ExternalProviders` fetches from `registry.terraform.io` on every run unless terraform's user-level cache is warm. Two consequences:
- CI runs on isolated networks (the kind we run real cloud-auth tests on) need outbound HTTPS to the registry. Our filesystem_mirror approach needs outbound only on cache miss (~once per host per version).
- Registry GPG verification can fail (the `hc-install` precedent — see B7/D and §10/G7) and would block all 4 steps.

### 9.4 Couples to Go-test shape

`ExternalProviders` lives inside `resource.TestCase` / `resource.TestStep`, which means tests are Go source compiled into `_test.go` files. We're going CLI-driven (YAML tests). Wrapping plugin-testing inside our CLI would require `go test -run='ourgenerated'` per test or driving it through `helper/resource.TestCase` programmatically — neither is clean.

### 9.5 Pinning breaks subtly

`ExternalProviders.VersionConstraint` accepts version constraint strings (`= 1.113.0`). plugin-testing then runs `terraform init` against the registry, which produces a `.terraform.lock.hcl` with that version's hashes. State of the lockfile across `TestStep` boundaries (different versions per step) is not documented and tester would have to validate empirically — the same lock-file fragility we hit in tester B2 + B3 still applies, just with less control.

### Bottom line
`ExternalProviders` is the right tool for "test my provider against a separate dependency provider at a fixed version." It's the wrong tool for "exercise the same provider at 4 different versions including a local build." We're using filesystem_mirror — one mechanism, fully under our control.

---

## 10. Failure modes & gotchas

This section catalogs every empirically-confirmed or known-risk pitfall and how the framework handles it. Organized by category. Each has a tracking ID for cross-reference.

### G1. Developer's `~/.terraformrc` silently bypasses pins (B1, A)
**Symptom:** terraform init prints `Installing databricks/databricks v1.114.0 (filesystem mirror)` but actually serves the developer's local `dev_overrides` build. Version-bound bugs become unreproducible.
**Mitigation:** Always set `TF_CLI_CONFIG_FILE=<workdir>/.terraformrc`. `TF_CLI_CONFIG_FILE` REPLACES `~/.terraformrc` in terraform's CLI config resolution (does NOT merge). Verified by tester.

### G2. `_override.tf` semantics for nested `required_providers` — **empirically confirmed**
**Symptom (hypothetical):** Override file's `terraform { required_providers { databricks = { version = "= 1.114.0" } } }` might not merge cleanly into a user-authored `terraform {}` block at the per-attribute level.
**Empirical result (tester, expA3):** Confirmed merge works perfectly. User has `databricks = ">= 1.0.0"` AND `google = ">= 5.0.0"` in main.tf; framework override redeclares ONLY databricks at `= 1.114.0`. Resulting lock: databricks @ `1.114.0` (override won), google @ `7.31.0` (preserved untouched). Override merges per-attribute even for nested `required_providers.<name>.<field>`.
**Consequence for design:** No preflight rejection of user HCL. Users may freely include their own `databricks` pin for IDE/standalone-`terraform plan` workflows. See §4 "Pre-flight validation" and §5 "Override file mechanism".
**Residual risk:** Terraform changing override-merge semantics in a future major version. Treat as a stable interface contract; if the contract breaks, fallback path is documented in §5 "Trade-offs accepted" item 3.

### G3. Lock file pins prior version, init fails on switch (B2)
**Symptom:** Step 1 leaves `.terraform.lock.hcl` pinning 1.113.0. Step 2's init resolves to 1.114.0, lockfile says 1.113.0, init errors `Error: Failed to install provider`.
**Mitigation:** `rm -f .terraform.lock.hcl && rm -rf .terraform/providers/` before EVERY step's init. State (`terraform.tfstate`) preserved.
**Why not `init -upgrade`:** tester reports it's registry-flaky ("intermittent 'no available releases'"). Delete-and-init is deterministic.

### G4. Single-platform `h1:` hashes break cross-platform CI (B3)
**Symptom:** Lock generated on darwin contains only darwin's `h1:` hash; CI on linux fails init.
**Mitigation:** `.terraform.lock.hcl` is per-run, ephemeral, gitignored. Never committed.
**Implementation:** `testframeworkV2/.gitignore` includes `.terraform.lock.hcl`. Each run regenerates fresh.

### G5. Mixed installation methods produce different hashes (B4)
**Symptom:** Same provider version, downloaded from registry vs unpacked from filesystem_mirror, has different `h1:` lock hashes. Switching between them errors.
**Mitigation:** ONE method everywhere — `filesystem_mirror`. Accepted tradeoff: filesystem_mirror provides `(unauthenticated)` `h1:` hashes (no GPG verification). The framework owns binary acquisition end-to-end (downloads from GitHub releases over TLS), so GPG would be belt-and-suspenders.

### G6. `os.Environ()` leaks (B5/E)
**Symptom:** Developer's shell has `DATABRICKS_HOST=...` set for some other reason. terraform subprocess inherits it. SDK auth chain prefers env over profile. Test runs against the wrong host.
**Mitigation:** Build subprocess env from a strict allowlist. **Critical: pass this list to `tfexec.Terraform.SetEnv()`** — per `tfexec` godoc, `SetEnv` REPLACES the subprocess env entirely (it does NOT inherit `os.Environ()`). The framework relies on this replacement semantics for B5 compliance.

```go
// internal/subprocenv/env.go
func Build(profile, tfrcPath, runDir string, passthrough []string) []string {
    env := []string{
        // Basics — required for any subprocess to run.
        "PATH="                       + os.Getenv("PATH"),
        "HOME="                       + os.Getenv("HOME"),
        "USER="                       + os.Getenv("USER"),
        "TMPDIR="                     + os.Getenv("TMPDIR"),                    // CI runners often have custom temp dirs
        // Locale — terraform formats dates / parses strings using locale.
        "LANG="                       + os.Getenv("LANG"),
        "LC_ALL="                     + os.Getenv("LC_ALL"),
        "LC_CTYPE="                   + os.Getenv("LC_CTYPE"),
        // Corporate CA bundle — needed for HTTPS on locked-down dev / CI machines.
        "SSL_CERT_FILE="              + os.Getenv("SSL_CERT_FILE"),
        "SSL_CERT_DIR="               + os.Getenv("SSL_CERT_DIR"),
        // Corporate proxy — terraform init contacts the registry; framework also fetches GitHub.
        "HTTPS_PROXY="                + os.Getenv("HTTPS_PROXY"),
        "HTTP_PROXY="                 + os.Getenv("HTTP_PROXY"),
        "NO_PROXY="                   + os.Getenv("NO_PROXY"),
        // Framework-controlled — terraform-specific.
        "TF_CLI_CONFIG_FILE="         + tfrcPath,
        "TF_PLUGIN_CACHE_DIR="        + filepath.Join(runDir, "plugins"),        // hardlink optimization (F2)
        // NOTE on TF_IN_AUTOMATION (and other tfexec-prohibited vars):
        // We do NOT set TF_IN_AUTOMATION here even though terraform documents it.
        // tfexec auto-manages it (terraform-exec@v0.25.0/tfexec/cmd.go:178
        // unconditionally sets env[automationEnvVar] = "1" in buildEnv) and
        // rejects callers who try to set it via tfexec.Terraform.SetEnv with
        // ErrManualEnvVar: `manual setting of env var "TF_IN_AUTOMATION" detected`.
        // The subprocess still receives it because tfexec sets it; we just must
        // not duplicate. The same prohibition (per cmd.go:24-43 prohibitedEnvVars)
        // applies to TF_CLI_ARGS, TF_CLI_ARGS_*, TF_INPUT, TF_LOG, TF_LOG_CORE,
        // TF_LOG_PATH, TF_LOG_PROVIDER, TF_REATTACH_PROVIDERS, TF_APPEND_USER_AGENT,
        // TF_WORKSPACE, TF_DISABLE_PLUGIN_TLS, TF_SKIP_PROVIDER_VERIFY, TF_VAR_*.
        // TestBuild_NoTfexecProhibitedVars in env_test.go enumerates the full
        // prohibited list as a regression guard against future re-introductions.
        //
        // Framework-controlled — Databricks SDK auth via profile only.
        "DATABRICKS_CONFIG_PROFILE="  + profile,
        "DATABRICKS_CONFIG_FILE="     + filepath.Join(os.Getenv("HOME"), ".databrickscfg"),
        // Explicitly NOT included (stripped — these would defeat the profile mechanism or pollute logs):
        //   DATABRICKS_HOST, DATABRICKS_TOKEN, DATABRICKS_USERNAME, DATABRICKS_PASSWORD,
        //   DATABRICKS_CLIENT_ID, DATABRICKS_CLIENT_SECRET, DATABRICKS_AUTH_TYPE,
        //   DATABRICKS_ACCOUNT_ID, DATABRICKS_AZURE_*, DATABRICKS_GOOGLE_*,
        //   ARM_*, GOOGLE_*, AWS_* (use passthrough_env if a test genuinely needs them),
        //   TF_LOG, TF_LOG_PATH, TF_VAR_* (use passthrough_env if needed).
        // TFV2_TERRAFORM_BIN is read by the framework BEFORE spawning terraform; it is
        // intentionally NOT propagated to the subprocess (terraform doesn't read it).
    }
    // Empty values are silently dropped by os.exec; no need to filter explicitly.
    // Opt-in passthrough (e.g., HTTP_PROXY, AWS_PROFILE for cross-cloud).
    // NEVER add DATABRICKS_* here — that defeats the profile mechanism.
    for _, name := range passthrough {
        if v := os.Getenv(name); v != "" && !strings.HasPrefix(name, "DATABRICKS_") {
            env = append(env, name+"="+v)
        }
    }
    return env
}
```
The `passthrough_env` list is supplied per-test via `test.yaml` (§4). Strict allowlist plus narrow opt-in covers the cases we care about (corporate proxy, cross-cloud env vars) without re-opening the env-leak vulnerability.

### G7. `internal/acceptance/init.go` package-level side effects (B6/F)
**Symptom:** Importing `internal/acceptance` triggers `init()` which calls `os.Setenv("TF_LOG", "DEBUG")` (init.go:39) and writes ~MB of plugin logs per step. Other globals leak similarly.
**Mitigation:** testframeworkV2 does NOT import `internal/acceptance`. Profile loading goes in our own `internal/profile/`. Any helper we'd otherwise grab from `internal/acceptance` is rewritten fresh.

### G8. `hc-install` / `tfinstall` GPG expiry (B7/D)
**Symptom:** Auto-installing terraform via `hc-install` fails with `openpgp: key expired`. Tester confirms `TestAccProviderPlanShouldSucceedWithIncompleteConfiguration` is currently skipped through 2026-05-08 for exactly this reason.
**Mitigation:** Framework does NOT auto-install terraform. Resolution order:
1. `--terraform-bin <path>` CLI flag
2. `TFV2_TERRAFORM_BIN` env var
3. `exec.LookPath("terraform")` in PATH
4. Hard error: `terraform binary not found; install terraform >= 1.5.0 or pass --terraform-bin`

Sanity check at run start: `terraform -version`, parse, fail if < 1.5.0.

### G9. Account-level profile shape varies by cloud (B12)
**Symptom:** GCP account profiles have `host = "https://accounts.gcp.databricks.com"`, `account_id = "..."`, `google_service_account = "..."`. AWS account profiles have different fields. Workspace profiles have yet another shape. Skip-checks (`requires.level: account`) need to read the profile correctly.
**Mitigation:** `internal/profile/` parses `~/.databrickscfg` (INI format) and exposes:
```go
type Profile struct {
    Name       string
    Host       string
    AccountID  string  // empty for workspace profiles
    AuthType   string  // explicit if set; else inferred
    Cloud      Cloud   // aws | azure | gcp — inferred from host
    Level      Level   // workspace | account — inferred from "accounts.*" prefix
    // Raw fields available for debugging.
    Raw        map[string]string
}
```
- Cloud inference: `accounts.cloud.databricks.com` → AWS; `accounts.azuredatabricks.net` → Azure; `accounts.gcp.databricks.com` → GCP; `*.cloud.databricks.com` (no `accounts.` prefix) → AWS workspace; `*.azuredatabricks.net` → Azure workspace; `*.gcp.databricks.com` → GCP workspace.
- Level inference: presence of `accounts.` host prefix OR `account_id` field → `account`; else `workspace`.
- Skip if `requires.cloud != profile.Cloud` (unless `requires.cloud == any`).
- Skip if `requires.level != profile.Level`.
- For GCP, the framework never reads `google_service_account` — it just sets `DATABRICKS_CONFIG_PROFILE` and lets the SDK resolve auth.

### G10. `terraform-plugin-testing.tfinstall` has a similar trap (D, B11)
**Symptom:** Importing `terraform-plugin-testing` doesn't auto-install, but transitively it pulls `hc-install`. We don't use plugin-testing at all (§9), so this doesn't apply, but note: we also don't use `tfexec.WithCheckTerraformVersion` if it ever calls `hc-install`. We use only `tfexec.NewTerraform(workDir, knownPath)`.

### G11. Framework writes into the user's source dir
**Symptom:** Generated `_tfv2_versions_override.tf`, `.terraform/`, `.terraform.lock.hcl`, `terraform.tfstate` accumulate in `testframeworkV2/issues-repro/issue_5672/`. Pollutes git status; tempting to commit by mistake.
**Mitigation:** Framework copies user `.tf` files into `~/.testframeworkv2/runs/<test>-<ts>-<rand>/workdir/` and runs everything there. User's source dir is read-only from the framework's POV. Defense in depth: `testframeworkV2/.gitignore` covers the framework-generated patterns in case anyone runs terraform manually in the source dir.

### G12. Failing destroy retry-loops (per reviewer I)
**Symptom:** Cleanup destroy fails (e.g., transient API error). Framework retries. Each retry takes minutes. Test times out.
**Mitigation:** No retry loop. Cleanup destroy attempts ONCE. On failure: log loudly, write resource list to `run.json` for manual cleanup, exit. Operator handles it.

### G13. Synthetic local version constraint matching
**Symptom:** Per Terraform's semver rules, prereleases (`-local` suffix) are not matched by `>=`/`<=`/`~>` operators — only by exact `=`.
**Mitigation:** Framework's override file always writes `version = "99.0.0-local"` (exact equality). Combined with override-merge semantics (G2 — empirically confirmed), the framework's `version` attribute always wins on the `databricks` provider regardless of what the user wrote in their main.tf. So even if a user has `version = ">= 1.0.0"` in their `terraform { required_providers { databricks = ... } }` block, the override transparently replaces the version field with `= 99.0.0-local` at test time. Confirmed by tester's expA3 + test1/test7 experiments.

### G14. `TF_LOG` from developer shell or leaked from `internal/acceptance`
**Symptom:** `TF_LOG=DEBUG` produces ~MB logs per step. Slows runs, fills disk, dilutes signal.
**Mitigation:** `TF_LOG` is NOT in the subprocess env allowlist. It's stripped automatically. Future debug needs: add `--tf-log <level>` framework flag that injects `TF_LOG` deliberately.

### G15. Race on cache writes
**Symptom:** Two parallel runs (or two steps within a parallel framework — v2) race to download the same uncached version.
**Mitigation:** Atomic write via `<path>.partial → fsync → rename`. v1 doesn't run runs in parallel anyway.

---

## 11. Worked example — `issue_5672`

### 11.1 Files

`testframeworkV2/issues-repro/issue_5672/test.yaml`:
```yaml
name: issue_5672_mws_workspaces_account_provider_config_regression
profile: ACCOUNT_AWS         # any account-level profile works (AWS / Azure / GCP)
cleanup: true                # no Apply steps → effectively no-op; left as default to document intent

requires:
  cloud: any                 # bug fires on AWS, Azure, GCP — confirmed empirically by tester
                             # (Appendix A: AWS account smoke test reproduced with different inner
                             # error than GCP — the outer regex is cloud-portable).
  level: account

steps:
  - name: passes_on_1_113_0
    version: "1.113.0"
    command: plan
    expect: success

  - name: fails_on_1_114_0
    version: "1.114.0"
    command: plan
    expect: failure
    # No error_substring — outer regex matches all auth-method variants. Schema supports both;
    # this test only needs the regex.
    error_regex: 'cannot populate provider_config for mws workspaces.*failed to resolve workspace_id'

  - name: fixed_on_1_114_1
    version: "1.114.1"
    command: plan
    expect: success

  - name: fixed_on_local
    version: "local"
    command: plan
    expect: success
```

`testframeworkV2/issues-repro/issue_5672/main.tf`:
```hcl
# Version pinning: this file does NOT pin a databricks version. The framework writes a
# `_tfv2_versions_override.tf` per step that pins via Terraform's *_override.tf
# per-attribute merge semantics. You MAY optionally add your own
#   terraform { required_providers { databricks = { source = "databricks/databricks", version = ">= 1.113.0" } } }
# block here for IDE / standalone-`terraform plan` workflows; the framework's override
# transparently wins on `version` at test time. Empirically validated — see DESIGN.md
# Appendix A (expA1/A2/A3 + B-COLLISION).

provider "databricks" {
  alias = "accounts"
  # No host, no account_id, no credentials. All come from DATABRICKS_CONFIG_PROFILE.
}

data "databricks_mws_workspaces" "all" {
  provider = databricks.accounts
}

output "workspace_count" {
  value = length(data.databricks_mws_workspaces.all.ids)
}
```

The user's `~/.databrickscfg` must have:
```ini
[ACCOUNT_GCP]
host                   = https://accounts.gcp.databricks.com
account_id             = 12345678-90ab-cdef-1234-567890abcdef
google_service_account = terraform-sa@project.iam.gserviceaccount.com
```
The framework doesn't validate the credential field choice — it just sets `DATABRICKS_CONFIG_PROFILE=ACCOUNT_GCP` and lets the SDK pick it up.

### 11.2 Execution timeline

```
$ tfv2 run testframeworkV2/issues-repro/issue_5672/

▶ issue_5672  profile=ACCOUNT_GCP  cloud=gcp  level=account
  workdir: /Users/tanmay.rustagi/.testframeworkv2/runs/issue_5672_mws_workspaces_account_provider_config_regression-2026-05-07T20-15-00-a3f2/workdir
  terraform: /usr/local/bin/terraform (v1.5.7)

  ▶ step 1/4: passes_on_1_113_0      databricks 1.113.0   command=plan  expect=success
    cache hit: ~/.testframeworkv2/providers/.../terraform-provider-databricks_1.113.0_darwin_arm64.zip
    write _tfv2_versions_override.tf  (= 1.113.0)
    rm -f .terraform.lock.hcl && rm -rf .terraform/
    ▶ terraform init …  ✓ (3.2s)
    ▶ terraform plan …  ✓ (4.1s)  no changes
    PASS

  ▶ step 2/4: fails_on_1_114_0       databricks 1.114.0   command=plan  expect=failure
    cache miss → download v1.114.0  (8.1s)
    write _tfv2_versions_override.tf  (= 1.114.0)
    rm -f .terraform.lock.hcl && rm -rf .terraform/
    ▶ terraform init …  ✓ (1.4s)
    ▶ terraform plan …  ✗ exit 1 (2.8s)
      stderr (excerpt — varies by auth method):
        Error: cannot populate provider_config for mws workspaces:
          failed to resolve workspace_id: failed to get the workspace_id:
          {auth-method-specific inner error — see below}
      Inner error by auth method (tester-empirical):
        AWS / OAuth M2M           : Unable to load OAuth Config
        GCP / google_service_account: strconv.ParseInt: parsing "": invalid syntax
    error_regex 'cannot populate provider_config for mws workspaces.*failed to resolve workspace_id' → MATCH
    PASS (failure-as-expected)

  ▶ step 3/4: fixed_on_1_114_1       databricks 1.114.1   command=plan  expect=success
    cache miss → download v1.114.1  (7.9s)
    write _tfv2_versions_override.tf  (= 1.114.1)
    rm -f .terraform.lock.hcl && rm -rf .terraform/
    ▶ terraform init …  ✓ (1.3s)
    ▶ terraform plan …  ✓ (3.7s)  no changes
    PASS

  ▶ step 4/4: fixed_on_local         databricks local     command=plan  expect=success
    go build ./ → ~/.testframeworkv2/providers/.../99.0.0-local/darwin_arm64/...  (4.2s)
    provenance: git_sha=1f62e20c dirty=true built_at=2026-05-07T20:16:08Z
    write _tfv2_versions_override.tf  (= 99.0.0-local)
    rm -f .terraform.lock.hcl && rm -rf .terraform/
    ▶ terraform init …  ✓ (1.2s)
    ▶ terraform plan …  ✓ (3.5s)  no changes
    PASS

  cleanup: skipped (no apply steps)

PASS  issue_5672  (4/4 steps, 41.6s)
```

### 11.3 What just happened (annotated)

**Step 1** establishes baseline. v1.113.0 is pre-regression; the post-Read hook in `common/resource.go` doesn't exist yet, so `data.databricks_mws_workspaces.all` reads cleanly during plan.

**Step 2** is the regression. v1.114.0 introduced PR #5492's `populateProviderConfigInState` hook. `databricks_mws_workspaces` was a `common.DataResource` (legacy helper that injects `provider_config` into the schema). The hook fires post-Read against the account host, hits `CurrentWorkspaceID()` → `GET /scim/v2/Me` → fails with an auth-method-specific inner error (`strconv.ParseInt("", …)` on GCP; `Unable to load OAuth Config` on AWS) wrapped by the cloud-portable outer message `cannot populate provider_config for mws workspaces: failed to resolve workspace_id: …`. Plan fails. The test's `error_regex` matches against the outer wrapper — step asserts failure-as-expected on AWS, Azure, AND GCP.

**Step 3** uses v1.114.1. Important nuance (per tester's empirical sha256s + see §16 F6):
- The git **tag** `v1.114.1` points at the same source commit as `v1.113.0` (`7a6b469e`). At the source level, they're identical.
- The released **binary** for v1.114.1 has a different SHA than v1.113.0's binary because it was independently rebuilt by goreleaser (different build timestamps / metadata embedded). Same source, different binaries.
- Terraform treats v1.113.0 and v1.114.1 as distinct versions — different filenames in the cache (`terraform-provider-databricks_v1.113.0` vs `terraform-provider-databricks_v1.114.1`), different lock-file entries.
This step proves the rollback was effective at restoring pre-regression behavior, but it does NOT prove that the real code fix works (since the source is identical to v1.113.0, which never had the bug).

**Step 4** uses the current branch via `local` build (synthetic version `99.0.0-local`). It exercises the framework's local-build pipeline end-to-end: `go build` from `--repo`, atomic install into the unpacked filesystem_mirror layout, provenance JSON, override-file rewrite, terraform init + plan against the freshly-compiled provider. **As of 2026-05-08 main has reverted PR #5667 + PR #5492 + related (see §16/F7), so the source under `local` is currently equivalent to v1.113.0 for this code path** — step 4 passes for the same reason step 1 does. Once the next forward-rolled fix lands on main, step 4 graduates from "validates the pipeline" to "proves the regression-fix works" without any test.yaml change required (the synthetic version mechanism does the version-flip transparently).

The framework-pipeline-validation property is independently valuable: it pins the local-build shape (compile flags, layout, override semantics, provenance JSON) so that when the forward-rolled fix DOES land, we already know the local pathway works. A regression in the local pipeline itself would be caught by step 4 today.

State carries across all 4 steps in the same `terraform.tfstate` (data-source-only config — state mostly empty, but conceptually the same model as for resource-based tests). `.terraform/providers/` is wiped between steps to force fresh resolution.

---

## 12. Public API surface

### 12.1 Go module

**Recommendation:** separate `go.mod` under `testframeworkV2/`.

Reasons:
- The provider repo's go.mod has hundreds of indirect deps; the framework's are minimal (`terraform-exec`, `yaml.v3`, `ini.v1` for profile parsing).
- Self-contained module means we can run the framework while bisecting provider source under test — no risk of `go build` deadlocks where the framework's `go build` of provider source depends on the framework itself.
- Trade-off: framework can't `import` provider packages directly. Acceptable; the framework treats provider as opaque binary.

`testframeworkV2/go.mod` skeleton:
```go
module github.com/databricks/terraform-provider-databricks/testframeworkV2

go 1.25

require (
    github.com/hashicorp/terraform-exec v0.25.0    // promoted to direct
    gopkg.in/yaml.v3 v3.0.1
    gopkg.in/ini.v1 v1.67.0
)
```

### 12.2 CLI

```
tfv2 run <test-dir>                             # run a single test
tfv2 run -r <root-dir>                          # recursively run all tests under root
tfv2 cache list                                 # show cached versions
tfv2 cache prune                                # delete provider cache
tfv2 build local --repo <path>                  # eagerly build local provider into cache (debug helper)
```

### 12.3 Common flags
```
--terraform-bin <path>     # override binary discovery (G8)
--cache-dir <path>         # override ~/.testframeworkv2/providers
--run-dir <path>           # override ~/.testframeworkv2/runs/<id>
--repo <path>              # provider repo root for `local` builds (defaults to pwd if it looks like one)
--no-cleanup               # overrides cleanup: true per-test
--keep-run-dir             # always keep run dir (default: keep regardless; this is forward-compat)
-v / --verbose             # framework debug logs
--json                     # emit run.json equivalent to stdout for CI
```

### 12.4 Environment overrides
```
TFV2_TERRAFORM_BIN     # = --terraform-bin
TFV2_CACHE_DIR         # = --cache-dir
TFV2_FORCE_REBUILD     # reserved for v2; v1 always rebuilds local
T_NO_CLEANUP=1         # global cleanup disable (same as --no-cleanup)
```

### 12.5 Programmatic API (advanced — for go test integration in v2)

```go
// internal/runner exposes a small public surface for go-test embeddings.
package runner

type Options struct {
    TerraformBin string
    CacheDir     string
    RunDir       string
    RepoRoot     string
    NoCleanup    bool
    Logger       Logger // optional; defaults to text to os.Stderr
}

type Runner struct{ /* ... */ }
func New(spec config.TestSpec, opts Options) *Runner
func (r *Runner) Run(ctx context.Context) (result.RunResult, error)
```

The CLI is a thin shell over this API (~100 LOC). v2 may add `gotest.Run(t *testing.T, dir string)` helper.

### 12.6 Key types (sketch)

```go
// internal/config
type TestSpec struct {
    Name     string
    Profile  string
    Cleanup  bool
    Requires Requires
    Steps    []Step
}
type Requires struct { Cloud, Level string }
type Step struct {
    Name           string
    Version        string         // "1.113.0" or "local"
    Command        string         // "plan" | "apply" | "destroy"
    Expect         string         // "success" | "failure"
    ErrorSubstring string
    ErrorRegex     *regexp.Regexp
    Timeout        time.Duration
}

// internal/providercache
type Target struct{ OS, Arch string }
func (t Target) String() string { return t.OS + "_" + t.Arch }

type Cache struct{ root string }
func New(root string) *Cache
// Resolve returns a path the filesystem_mirror can serve from.
// For released versions: the .zip path.
// For "local": the unpacked binary's parent dir (and the synthetic version "99.0.0-local").
func (c *Cache) Resolve(ctx context.Context, version string, target Target) (path, syntheticVersion string, err error)

// internal/profile
func Load(name string) (*Profile, error)

// internal/subprocenv
func Build(profileName, tfrcPath string) []string  // explicit allowlist (G6/E)

// internal/tfrcwriter
func WriteTerraformRC(workDir, mirrorRoot string) (path string, err error)
func WriteVersionsOverride(workDir, version string) error  // writes _tfv2_versions_override.tf

// internal/terraform
func Locate(override string) (path, version string, err error)  // G8
```

---

## 13. Open questions

Things to validate during implementation, in priority order. None block this design from being approved as v2; all are either tester-checkable or low-risk to flip.

1. ~~**Override file semantics for nested `terraform { required_providers {} }`**~~ — **CLOSED.** Tester's expA3 empirically confirmed: override file containing only `databricks` correctly merges with a user main.tf that has both `databricks` and `google` pins. Override wins on `databricks.version`; `google` is preserved. See §10 G2 for details.

2. ~~**Issue #5672 cloud-specificity.**~~ — **CLOSED.** Tester ran the full 4-step mission test against an AWS account profile (`https://accounts.cloud.databricks.com`, OAuth M2M) and reproduced the bug with a different inner error (`Unable to load OAuth Config` vs GCP's `strconv.ParseInt`) wrapped by the same outer pattern. `test.yaml` uses `requires.cloud: any` accordingly. See Appendix A "AWS account full-smoke" entry.

3. ~~**Directory naming convention for tests.**~~ — **CLOSED in v5.0.** The earlier `account/` / `workspace/` / `ucws/` / `ucacct/` (one tree per `requires.level` value) didn't scale once Phase 2 added fixtures spanning multiple levels. v5.0 splits along **intent** instead of profile level: `issues-repro/issue_<N>/` for fixtures that reproduce a specific GitHub issue (any level), `tests/<descriptive-slug>/` for green-path / smoke fixtures that aren't tied to a bug. Per-test `requires.{cloud,level}` does the actual host-gating. See §3 "Directory layout" for the convention block and the v5.0 delta entry in the doc header for the rationale.

4. **terraform CLI minimum version.** Picked 1.5.0. Tester to confirm the mission test passes against terraform 1.5.7 (currently on disk) and 1.10+ (likely on CI). If 1.5.0 has any filesystem_mirror edge case, raise the floor.

5. **Lock-file `init -upgrade` fallback worth keeping?** B2 says no — delete is deterministic. If implementer hits an edge case where delete-then-init fails on terraform 1.5.x specifically, fall back to `-upgrade`.

6. **Profile auto-discovery vs explicit `profile` field.** v1 requires explicit `profile:` in test.yaml. Future: read `DATABRICKS_CONFIG_PROFILE` env if `profile:` empty? Probably not — explicit is clearer for tests in CI.

7. **Behavior when `local` build fails at the `go build` step.** Currently surfaces compile errors directly (cmd.Stdout/Stderr to os.Stdout/Stderr) and returns error from runner. Step reports as FAIL with reason "go build failed". Acceptable.

8. **Streaming logs vs end-of-step capture.** Currently writes both stdout/stderr to log files AND tees to terminal at `-v` level. tfexec.Apply doesn't easily stream — we use `tf.SetStdout`/`SetStderr`. Confirm this works for terraform's progress output.

9. **CI integration patterns.** Once the framework lands, where in CI does it run? Likely a separate workflow that selects a profile per cloud env (similar to existing `internal/acceptance`'s CLOUD_ENV pattern). Out of scope for this design; flagging.

---

## 14. Out of scope (v1)

| Item | Why deferred | When |
|---|---|---|
| Parallel runs | `~/.databrickscfg` profile isolation hard, low value yet | v2 if N tests becomes large |
| Provider zip checksum verification | TLS to GitHub sufficient for now | v2 — `_SHA256SUMS` validation |
| Cross-arch / cross-OS builds | Host-arch is fine for ad hoc dev | v3 with build matrix |
| `Check func(state)` Go assertions | YAGNI for #5672 — exit code + stderr regex covers it | v2, alongside `gotest` helper |
| Hash-based local-build cache | `go build` incremental is fast enough | only if benchmarks demand |
| Windows support | Dev machines are darwin/linux | when we have a Windows runner |
| Network mirror / multi-machine cache | Filesystem fine for one-machine dev | v3 |
| Multi-provider tests | Issue #5672 needs only `databricks` | v2 — extend override file per step |
| go test integration | CLI-first ships sooner | v2 (`gotest.Run(t, dir)`) |
| `passthrough_env` per-test | Strict allowlist works for issue #5672 | v2 if a test needs `AWS_PROFILE` etc. |
| `--tf-log <level>` debug flag | Default is no `TF_LOG`; debug is rare | implement when first needed |
| GPG verification of provider zip | filesystem_mirror tradeoff (G5) | v2 footnote on `92A95A66446BCE3F` (B10) |

---

## 15. Implementation milestones

For implementer-tf11. Each milestone is independently shippable to PR; M1-M3 don't need cloud auth.

**M1 — `providercache` standalone.** `Cache.Resolve(version, target) → path`. Download from GitHub releases, atomic write, packed for released versions. Unit tests with `httptest.Server`.

**M2 — `tfrcwriter`.** Generate `.terraformrc` (with absolute mirror path + narrow `include`) + `_tfv2_versions_override.tf`. Pure string templating; trivial unit tests.

**M3 — `subprocenv` + `terraform.Locate`.** Curated env allowlist. Binary discovery with fallthrough. Sanity-check version. Unit tests.

**M4 — `runner` happy path with mocked tfexec.** Single-step PASS case using a `tfexec`-shaped fake. Integration: copy user HCL, write override, init, plan, assert. Test logs verified.

**M5 — End-to-end mission test (4 steps).** Real `terraform` invocation. Issue #5672 reproduced against real account profile (AWS or GCP — bug is host-agnostic per Appendix A AWS smoke). Sequential 4 steps → 4 PASS. Tester signs off.

**M6 — `local` builds.** `99.0.0-local` synthetic, `go build`, unpacked layout, provenance JSON.

**M7 — `tfv2 run -r` + `tfv2 cache` subcommands.** Polish.

---

## 16. Footnotes

**F1. Databricks signing key.** Per tester (B10), Databricks GitHub release zips are signed by `92A95A66446BCE3F` (Serge Smertin), not Hashicorp's signing key. We don't verify any GPG signatures in v1 (G5 — filesystem_mirror tradeoff). When we add verification in v2, this is the key to trust for `_SHA256SUMS.sig`.

**F2. Why we unconditionally strip `TF_LOG`.** `internal/acceptance/init.go:39` calls `os.Setenv("TF_LOG", "DEBUG")` at package-init time. Even though we don't import that package (G7/F), some users set it manually. ~MB-per-step logs from `TRACE`/`DEBUG` are signal-diluting. If a future user needs it, `--tf-log <level>` flag injects deliberately.

**F3. Atomic cache writes.** `<path>.partial → fsync → rename` is sufficient on POSIX. On Windows (out of scope v1), `os.Rename` over an existing file can fail; use `MoveFileEx` with `MOVEFILE_REPLACE_EXISTING`.

**F4. Why we accept filesystem_mirror's `(unauthenticated)` h1: hashes.** From terraform docs: filesystem_mirror produces lock entries marked `(unauthenticated)` because the mirror doesn't carry the same signed metadata as the public registry. The lock still functions for reproducibility within a host, just without signature verification. Since the framework owns binary acquisition end-to-end (downloads from GitHub releases over TLS, atomically writes to host-local cache), GPG verification would be belt-and-suspenders. v2 may add it as defense-in-depth.

**F5. Run dir lifetime + naming.** Run dirs are named `<test>-<ts>-<rand>/` where `<ts>` is `2006-01-02T15-04-05` (RFC3339 with `:` swapped to `-` for filesystem safety) and `<rand>` is a 4-character lowercase hex string from `crypto/rand`. The random suffix prevents collisions when the same test runs twice in <1s (reviewer + tester caught the same-second risk). Run dirs are preserved on disk forever in v1 (intentionally — debugging). v2 may add `tfv2 runs prune` or auto-GC after N days. Disk usage estimate: ~10-50 MB per run (log files + workdir + state). Not a v1 problem.

**F6. v1.114.1 vs v1.113.0 — same source, different binaries.** The git tag `v1.114.1` points at commit `7a6b469e`, identical to `v1.113.0`. But the released binary content differs (tester sha256: 1.113.0 = `d2ee4a9a...9f9f9469`, 1.114.1 = `ddf8cdb0...32d368b`, both 64,601,938 bytes). This is normal: same source rebuilt by goreleaser at different times produces different binaries (build timestamps, environment metadata embedded by Go's linker). Functionally equivalent at runtime; Terraform treats them as distinct versions because the cached filenames differ (`terraform-provider-databricks_v1.113.0` vs `_v1.114.1`) and the lock-file entries differ. The framework caches them independently, by version string. No special handling needed. Worth knowing for anyone debugging "why does v1.114.1 still produce a different lock entry than v1.113.0".

**F7. PR #5492 / #5667 revert chain (2026-05-08).** Between v1.114.0 (which contained PR #5492's "workspace_id support of SDKv2 resources") and the current state of `main`, the following commits reverted the workspace_id machinery and its dependent fixes: `cc89a814`, `f4aebfc2`, `8a4e64ab`, `0d2c487e`, `20108fee`, `a20971ad`, `c707a744`. Net effect: `main`'s source for the affected code paths (the post-Read hook in `common/resource.go`, `validateWorkspaceID` in `common/client.go`, `databricks_mws_workspaces` migration to `common.AccountData`) is currently equivalent to v1.113.0. This means the issue #5672 mission test's step 4 (`local`) currently passes because the source matches v1.113.0, NOT because a forward-rolled fix is in place. Step 4 still validates the framework's local-build pipeline (compile flags, layout, override semantics, provenance JSON) and will graduate to "proves the regression-fix works" automatically once the next forward attempt at workspace_id resolution lands on main — the synthetic-version mechanism handles the version-flip transparently with no test.yaml change required. The new `issues-repro/issue_5678/` fixture (added in v5.0) covers the state-incompatibility consequence of that same revert chain, and the next forward-rolled fix will introduce a new fixture in `issues-repro/` to cover whatever new bug surface (if any) emerges.

---

## 17. v2 mode — per-step HCL + state assertions

Added in v6.0. v2 mode is opt-in: a test stays in v1 unless every step declares `config:`. v1 tests continue to work unchanged — same schema, same runner pipeline, same JSON output shape. v2 tests gain two capabilities that v1 cannot express:

1. **Per-step HCL swap.** Each step has its own `*.tf` file (`config:`); the runner wipes user `*.tf` from the workdir and copies the step's file in before each `terraform init`. State (`terraform.tfstate`) survives the swap. This is the substrate for state-incompatibility regressions like #5678 where the bug only fires when state written by version A is read by version B with a different schema.
2. **Structured state assertions.** Each step can declare an `assert:` block; after the step's command succeeds, the runner runs `terraform show -json` and checks resource presence + attribute values. Failures are surfaced via the existing `StepResult` (new `Assertions []AssertionFailure` field) and a per-step `step_<n>_<name>.assert.log`.

### 17.1 Mode detection

A test is in **v2 mode** iff every step has a non-empty `config:`. v1 mode iff no step has one. Mixed configurations are a parse-time error (`v2 mode requires every step to set config:`). The all-or-none rule lives on `validateV2Consistency`; the runtime accessor is `(*TestSpec).Mode()` returning `ModeV1` / `ModeV2`.

Why all-or-none: a v1-style `main.tf` in the source dir would conflict with a v2 step's per-step swap. Allowing a mix means defining surprising priority rules between user-authored `main.tf` and per-step `config:` files. Easier to require uniformity.

### 17.2 `config:` schema

```yaml
steps:
  - name: setup
    config: 01_initial.tf            # path RELATIVE to the test dir (slug-shaped basename)
    version: "1.114.0"
    command: apply
    expect: success
```

Path validation: the `config:` value must match `^[a-zA-Z0-9_][a-zA-Z0-9_-]*\.(tf|tf\.json)$`. No traversal (`../`), no subdirectories (`subdir/file.tf`), no hidden files (`.foo.tf`), no other extensions (`wrong.txt`). The framework's `_tfv2_` prefix is also forbidden — that's the namespace for runner-generated files.

### 17.3 `assert:` schema

```yaml
assert:
  - resource: databricks_token.pat   # required: <type>.<name> or data.<type>.<name>
    # present: true is the default — omit for "I expect this resource to exist"
    attrs:
      comment: "tfv2-test"            # YAML scalars / lists / maps; numerics are coerced to float64
      lifetime_seconds: 3600

  - resource: databricks_some.removed
    present: false                    # explicit absence assertion; attrs: rejected here
```

Resource address rules: root-module only (`<type>.<name>` for managed, `data.<type>.<name>` for data sources). Module-scoped addresses (`module.X.Y.Z`) are deferred to v3 — v2 launch covers issue #5678 fully without them.

### 17.4 Per-step file flow

In v1 mode, `prepareRun.copyTerraformFiles` runs once at run start; every step shares the same set of `.tf` files. In v2 mode the file flow changes:

- **`prepareRun`** does NOT pre-copy any user `.tf` files. Workdir starts empty (apart from `.terraformrc` written by the framework).
- **Each step** runs `swapStepConfig`:
  1. Remove every user `*.tf` / `*.tfvars` from workdir (preserves `_tfv2_*.tf` and dot-files).
  2. Copy `<sourceDir>/<step.Config>` into workdir under its basename.
  3. State (`terraform.tfstate`) and the framework override are untouched.

Then the existing per-step pipeline takes over: write `_tfv2_versions_override.tf`, wipe `.terraform/` + `.terraform.lock.hcl`, init, command, assert.

### 17.5 State-assertion semantics

After a successful command (only when `expect: success`), the runner spawns `<terraformBin> show -json` in workdir, parses the JSON, and walks `values.root_module.resources[*]` for each assertion:

1. **Presence**: find the resource by canonical address. `present: true` → fail if missing. `present: false` → fail if found.
2. **Attrs** (only when `present: true` and resource found): for each `(key, expected)` pair, dot-walk into the resource's `values` map and `reflect.DeepEqual`-compare. Numeric expected values are coerced to `float64` first because YAML's untagged-int decoder produces `int` while JSON produces `float64`.
3. **Sensitive attrs**: `terraform show -json` reports them as `"(sensitive)"`. The framework surfaces a clear failure ("attribute is marked sensitive in state — assert against a non-sensitive proxy") rather than a value mismatch.
4. **Collect-all-failures**: the evaluator never short-circuits. An assertion with three mismatched attrs produces three failures; the runner surfaces all of them via `StepResult.Assertions`.

`AssertionFailure` shape (in `internal/result`):

```go
type AssertionFailure struct {
    Address  string  // e.g. "databricks_token.pat"
    Reason   string  // "expected present, not found in state" | "value mismatch" | ...
    Field    string  // attribute key (only on per-attr failures)
    Expected any     // YAML-decoded
    Actual   any     // JSON-decoded
}
```

### 17.6 Per-step assert.log

For every step that runs assertions, the framework writes `step_<n>_<name>.assert.log` next to the existing stdout/stderr logs:

```
OK   databricks_token.pat
FAIL databricks_other.x.comment: value mismatch (expected=foo, actual=bar)
```

`StepResult.AssertLog` carries the absolute path; tester's debug grep over `~/.testframeworkv2/runs/<run-id>/` picks it up alongside the existing logs.

### 17.7 Parse-time validation rules (consolidated)

`internal/config.validate` enforces all of:

1. All-or-none `config:` across steps.
2. `config:` shape: slug basename, `.tf` or `.tf.json` extension, no `_tfv2_` prefix.
3. `assert:` requires v2 mode (every step has `config:`).
4. Each assertion has a non-empty `resource:` matching the root-module address shape.
5. `present: false` is incompatible with `attrs:` (logically inconsistent).
6. `assert:` is incompatible with `expect: failure` (no meaningful state to assert against).
7. Sensitive-attribute assertions are surfaced at evaluation time (parse-time gate would require provider schema introspection — out of scope).

### 17.8 Out of scope for v6 launch (v3 fodder)

| Item | Why deferred |
|---|---|
| Module-scoped resource addresses | Needs proper HCL2 address parsing. Root-module covers the bug surface we have. |
| `attrs:` deep-nested matchers via JSON Pointer | Dot-walk handles the common cases. |
| Plan-time assertions (`assert_plan:` matching the plan diff) | Would close #5678's "must be replaced" gap, but plan output is unstructured. Needs separate `tfexec.Plan(... -out ...)` + `ShowPlanFile` flow. |
| Cross-step state diffing | Composite of single-step assertions; library helpers come later. |
| Custom matcher types (regex, `>=`, etc.) | Wait until 1+ test demands it. |

### 17.9 Implementation seam summary

| Package | Change |
|---|---|
| `internal/config` | `Step.Config` + `Step.Assert` fields; `Mode()` accessor; `validateV2Consistency`; address-shape validator. |
| `internal/stateassert` (new) | `Run(ctx, workdir, terraformBin, env, []Assertion) ([]AssertionFailure, error)` — spawns `terraform show -json`, evaluates each assertion, returns structured failures. |
| `internal/result` | `AssertionFailure` type; `StepResult.Assertions` + `StepResult.AssertLog` fields (omitempty). |
| `internal/runner` | `prepareRun` skips bulk-copy in v2; `runStep` calls `swapStepConfig` before init in v2 mode and `runStateAssert` after the command on success-path steps. Per-step `assert.log` written. |
| Fixtures | `tests/token_lifecycle_v2/` — 3-step apply-modify-destroy lifecycle of `databricks_token` exercising all three assertion shapes (present-with-attrs, present-with-attrs after modify, present:false). |

---

## Appendix A — tester empirical evidence (highlights)

Tester ran a battery of experiments to validate this design's mechanism choices. Highlights:

- **expA1**: main.tf pinned to `= 1.113.0`; framework override pinned to `= 1.114.0`. Lock file resolved to `1.114.0`. → override wins.
- **expA2**: main.tf has no `terraform {}` block; override alone provides the full pin. Lock file resolved to overridden version. → override works as a sole source.
- **expA3 (the design's keystone evidence)**: main.tf has `databricks = ">= 1.0.0"` AND `google = ">= 5.0.0"`. Override redeclares only `databricks` at `= 1.114.0`. Lock file: databricks @ `1.114.0` (override won) AND google @ `7.31.0` (preserved). → per-attribute merge, even nested in `required_providers`. **This is the empirical basis for §4's "user main.tf may include their own databricks pin" relaxation.**
- **B-COLLISION**: main.tf pins `databricks = "= 1.113.0"` (an EXACT-equality competing constraint); override pins `= 1.114.0`. Result: no error, lock pinned to 1.114.0 — override wins even against a competing exact pin. → the override-merge contract holds even for the most adversarial user-authored case.
- **test3-narrow vs test3-broad**: confirmed `include = ["registry.terraform.io/databricks/*"]` works; `include = ["registry.terraform.io/*/*"]` breaks `hashicorp/google`. → narrow include is mandatory (§5 / F1).
- **test1 / test7**: synthetic version `99.0.0-local` resolves correctly with exact-match constraint. Operators (`>=`, `~>`) do NOT match prereleases as expected per semver. → exact `=` is the only safe operator for `local` (§8 / G13).
- **plugin_cache_dir**: when set alongside `filesystem_mirror`, terraform hardlinks resolved provider files (same inode `577301870` confirmed across init invocations). → free within-run reuse (§5 / F2).
- **Lock file behavior**: tester-empirical conclusion: `rm -f .terraform.lock.hcl` alone is *sufficient for correctness* (terraform's installer accumulates versions in `.terraform/providers/<v>/` happily and selects via lock). The framework's stronger `rm -rf .terraform` is *recommended for disk hygiene over long-running tests*. Tester signed off on §7.1.c as-written ("researcher's §7.1.c is correct as-written"). `init -upgrade` alone is registry-flaky and rejected. → §7 / G3.
- **AWS account smoke test (full 4-step)**: tester ran the entire mission test against `https://accounts.cloud.databricks.com` with OAuth M2M auth. ALL 4 STEPS PASS, including step 4 with `99.0.0-local` built from the `testframework-v2` branch. State carries across version flips. **The bug reproduces on AWS with a different inner error (`Unable to load OAuth Config`) — confirming host-agnosticism.** → §1 row 2, `requires.cloud: any`, looser regex.
- **Binary SHA fact**: 1.113.0 binary `d2ee4a9a9fff74e9013449caf736ae723430c4313744675855434baa99979469`, 1.114.1 binary `ddf8cdb0c9ccccf6a9aff917f47b6eb5d6e5f4da1d05bdadd8869c7bd32d368b` — same size 64,601,938 bytes, different content. Same git source SHA, independent goreleaser rebuilds. → §11 / F6.

---

*End of design v4.2. Last 3 stale-reference bounces from reviewer's diff-review folded in. Diff vs v4.1 sent to reviewer-tf11 + tester-tf11. Ready for TL Task #5 synthesis.*
