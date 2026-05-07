# Contributing to testframeworkV2

This guide walks you from "I want to write a regression test for a Databricks
provider bug" to "test passes locally on my profile". For framework internals,
see `DESIGN.md`. For a quick "what is this thing", see `README.md`.

The running example throughout is `issues-repro/issue_5672/`, which
reproduces [issue #5672](https://github.com/databricks/terraform-provider-databricks/issues/5672) end-to-end across 4 provider versions.

---

## Prerequisites

You need:

| Tool | Why | Where |
|---|---|---|
| `terraform` ≥ 1.5.0 on `$PATH` | Framework spawns it as a subprocess; no auto-install. | Override with `--terraform-bin <path>` or `TFV2_TERRAFORM_BIN` env if not on `$PATH`. |
| Go ≥ 1.25 | To build `tfv2` and (for `version: local`) the provider itself. | `go.mod` pins the toolchain. |
| `~/.databrickscfg` with a working profile | Auth flows through `DATABRICKS_CONFIG_PROFILE` only — no inline tokens in HCL. | Run `databricks configure --profile <name>` if you don't have one. |
| Network access to GitHub releases | First test run downloads provider zips into `~/.testframeworkv2/providers/`. Cached after that. | Behind a corporate proxy? `HTTPS_PROXY`/`HTTP_PROXY`/`NO_PROXY` are propagated automatically (DESIGN.md §10/G6). |

For day-to-day test iteration, run from source:

```sh
cd testframeworkV2/
go run ./cmd/tfv2 version     # prints "tfv2 dev"
```

For repeated runs (CI, perf-sensitive loops) install a binary instead — see
README.md "Quickstart". The rest of this guide uses `go run ./cmd/tfv2`.

---

## Step 1 — Pick a directory

Tests are grouped by **intent** (per DESIGN.md v5.0):

| Subdir | Use when… |
|---|---|
| `testframeworkV2/issues-repro/issue_<N>/`  | You are reproducing a specific GitHub issue. The directory name is `issue_<number>` (e.g. `issue_5672`, `issue_5678`). |
| `testframeworkV2/tests/<descriptive-slug>/` | The test is a green-path / smoke / regression-guard fixture NOT tied to a specific issue (e.g. `workspace_data_source_smoke`). |

Profile level (workspace / account / UC) is NOT encoded in the directory
name; it's declared per-test via `requires.level` in `test.yaml`. The
framework reads that field and skips the test cleanly if the named profile
doesn't match.

Issue #5672 is account-only (the post-Read hook only fires against an account
host) → `issues-repro/issue_5672/`. The workspace-level `databricks_token`
regression for issue #5668 lives at `issues-repro/issue_5668/`, and the
`databricks_catalog_workspace_binding` rollback regression for issue #5678
lives at `issues-repro/issue_5678/`. A green-path data-source smoke test
goes under `tests/` (e.g. `tests/workspace_data_source_smoke/`).

Slug-style names: lowercase, digits, `_`, `-` (matches `^[a-z0-9_-]+$`).

```
testframeworkV2/issues-repro/issue_5672/
├── main.tf
└── test.yaml
```

---

## Step 2 — Write `main.tf`

Two rules:

1. **Empty `provider "databricks"` block.** Auth comes from
   `DATABRICKS_CONFIG_PROFILE` (set by the framework from `test.yaml: profile`).
   No `host`, `token`, `client_id`, `google_service_account` etc. inline.
2. **You MAY pin `databricks` in your own `terraform { required_providers {} }` block** (or skip it entirely). If you pin, the framework's
   `_tfv2_versions_override.tf` overrides only the `version` field at test time
   via Terraform's `*_override.tf` per-attribute merge — empirically validated
   in DESIGN.md Appendix A (expA1/A2/A3 + B-COLLISION). Other providers
   (`hashicorp/google`, `hashicorp/random`, etc.) are preserved untouched.

The minimum HCL for the running example:

```hcl
# testframeworkV2/issues-repro/issue_5672/main.tf

provider "databricks" {
  alias = "accounts"            # alias is fine; not required
}

data "databricks_mws_workspaces" "all" {
  provider = databricks.accounts
}

output "workspace_count" {
  value = length(data.databricks_mws_workspaces.all.ids)
}
```

Do NOT add `provider_config { workspace_id = "..." }` blocks — the legacy
helper field is exactly the schema shape that triggers issue #5672. Tests
should exercise the user-shaped configuration.

---

## Step 3 — Write `test.yaml`

The schema (full reference: DESIGN.md §4) has three top-level required fields,
two optional ones, and a list of steps. Walk-through using `issues-repro/issue_5672/test.yaml`:

```yaml
name: issue_5672_mws_workspaces_account_provider_config_regression
profile: ACCOUNT_AWS                 # ~/.databrickscfg section name (any cloud's account-level profile works here)
cleanup: true                        # default true; no-op when all steps are plan-only

requires:                            # skip-on-mismatch declarative gates
  cloud: any                         # one of: aws, azure, gcp, any  (default: any)
  level: account                     # one of: workspace, account, ucws, ucacct  (default: workspace)

steps:                               # ≥1, executed in order, state carries across
  - name: passes_on_1_113_0          # slug; unique within test
    version: "1.113.0"               # semver string OR "local"
    command: plan                    # plan | apply | destroy  (default: apply)
    expect: success                  # success | failure  (default: success)

  - name: fails_on_1_114_0
    version: "1.114.0"
    command: plan
    expect: failure
    # When expect=failure: ≥1 of error_substring (literal) or error_regex (Go RE2) is required.
    # Both can be set (AND semantics). Match is against captured stderr.
    error_regex: 'cannot populate provider_config for mws workspaces.*failed to resolve workspace_id'

  - name: fixed_on_1_114_1
    version: "1.114.1"
    command: plan
    expect: success

  - name: fixed_on_local             # "local" = build from --repo (defaults to provider repo root)
    version: "local"
    command: plan
    expect: success
```

### Picking versions

- Released semver (e.g. `"1.114.0"`) → fetched from GitHub releases, cached at `~/.testframeworkv2/providers/`.
- `"local"` → `go build` of the provider repo at `--repo` (default: parent dir if it looks like a tf-provider checkout).

For a regression test, the standard 4-version shape is:
1. **Last good** (pre-regression baseline) — proves the test setup is sane.
2. **First bad** (the regression) — `expect: failure` + cloud-portable regex.
3. **Last bad in any released tag** OR **patch release** — proves rollback worked.
4. **`local`** — proves the real fix on the current branch.

For #5672, step 3 uses v1.114.1 — a rollback retag of v1.113.0 (same git source
SHA, different binary SHA because goreleaser rebuilds independently per tag;
see DESIGN.md §16 F6 for the empirical sha256s). Step 3 passing only confirms
the rollback worked; it does NOT validate the actual code fix. That's why
step 4 (`local`) is the most important one — it's the only step that exercises
code that isn't in any released tag yet.

### Crafting `error_regex`

Make it cloud-portable. Inner errors often differ by auth method (e.g., AWS
OAuth: `Unable to load OAuth Config`; GCP service account:
`strconv.ParseInt: parsing "": invalid syntax`); anchor on the outer wrapper
that's the same across clouds. Stderr-only matching; multiline with `(?s)`
flag implicit (so `.` matches newlines).

If you need a literal substring match, use `error_substring`. Both can be set
together (AND semantics). At least one is required when `expect: failure`.

---

## Step 4 — Run locally

```sh
cd testframeworkV2/

# Recommended one-liner — wraps `go run ./cmd/tfv2 run` underneath.
make test issues-repro/issue_5672/

# Run every committed fixture (TFV2_RUN=1 set by the Makefile):
make test-all
```

The Makefile's `make test <path>` target is just a wrapper around the
direct CLI invocation; both forms produce identical output. Use
whichever fits your workflow.

```sh
# Direct CLI form — equivalent to the make wrapper, useful when you
# want to pass extra flags Make doesn't pre-wire:
go run ./cmd/tfv2 run issues-repro/issue_5672/

# Override the terraform binary if it's not on PATH:
go run ./cmd/tfv2 run --terraform-bin /opt/terraform/1.5.7/terraform \
  issues-repro/issue_5672/

# Disable cleanup destroy (preserves state for inspection):
go run ./cmd/tfv2 run --no-cleanup issues-repro/issue_5672/

# Verbose framework logs to stderr (does NOT enable terraform's TF_LOG):
go run ./cmd/tfv2 run --verbose issues-repro/issue_5672/

# Run via `go test` — every fixture is also a Go subtest under
# TestFixtures (TFV2_RUN=1 gate keeps `go test ./...` cheap by
# default). Subtest names use the tree-preserving path so the
# `-run` filter is `TestFixtures/<tree>/<fixture-dir>`:
TFV2_RUN=1 go test -run 'TestFixtures/issues-repro/issue_5672' -v ./...
```

`--repo` is **auto-discovered** by walking up from the working
directory looking for the provider repo's `go.mod` (DESIGN.md §12.6).
If you're inside the provider checkout, no flag is needed. Pass
`--repo <path>` or set `TFV2_REPO` only when invoking from outside a
checkout AND your test has at least one step with `version: local`.
Tests pinning only released semvers don't need a repo root at all.

Expected output for the mission test (all four steps green):

```
[PASS] step 1 (passes_on_1_113_0): 1.113.0      plan in 5.1s
[PASS] step 2 (fails_on_1_114_0): 1.114.0       plan in 4.7s     (failure-as-expected)
[PASS] step 3 (fixed_on_1_114_1): 1.114.1       plan in 4.6s
[PASS] step 4 (fixed_on_local):   99.0.0-local  plan in 5.9s
----------------------------------------------------------
issue_5672_mws_workspaces_account_provider_config_regression: PASS (4/4 steps in 22.4s)
run dir: /Users/<you>/.testframeworkv2/runs/issue_5672_..-2026-05-08T08-15-00-a3f2
```

Exit codes:
- `0` — every step passed (or test was skipped due to `requires` mismatch).
- `1` — at least one step did not pass, or the framework hit an error.
- `2` — usage error (bad flags, unknown subcommand, missing `<test-dir>`).

Per-step logs land at
`~/.testframeworkv2/runs/<test>-<ts>-<rand>/step_<n>_<name>.{stdout,stderr}.log`.
The run directory is preserved on disk after the test ends — open it with
your editor to inspect the captured streams, the generated `.terraformrc`,
the generated `_tfv2_versions_override.tf`, and the `terraform.tfstate` that
carried across steps.

---

## Step 5 — Debugging

**Test is being skipped unexpectedly.** Check `requires:` against your
profile (INI section names in `~/.databrickscfg` are case-sensitive — match
exactly):
```sh
grep -A 5 "^\[ACCOUNT_AWS\]" ~/.databrickscfg
```
If your profile has `host = https://accounts.cloud.databricks.com` it's an
AWS account-level profile, which matches `cloud: aws` (or `any`) +
`level: account`.

**Test fails on a step you expected to pass.** Open the per-step stderr log:
```sh
# <test-name> is the `name:` field from your test.yaml (NOT the source-dir name).
# For issues-repro/issue_5672/test.yaml that's
# issue_5672_mws_workspaces_account_provider_config_regression.
ls ~/.testframeworkv2/runs/<test-name>-*/
cat ~/.testframeworkv2/runs/<test-name>-*/step_1_*.stderr.log
```

**Wrong provider version is being served.** Verify the lock file matches the
expected version:
```sh
cat ~/.testframeworkv2/runs/<run-id>/workdir/.terraform.lock.hcl
```
The framework wipes `.terraform/` and `.terraform.lock.hcl` between every
step, so each step's lock should reflect ONLY that step's pinned version.

**Local build (`version: local`) isn't picking up my changes.** Two things
to check:

1. `--repo` is pointing at the right tree. `local-version.json` next to the
   built binary records the git SHA + dirty flag at build time:
   ```sh
   cat ~/.testframeworkv2/providers/registry.terraform.io/databricks/databricks/99.0.0-local/<os>_<arch>/local-version.json
   ```
   A copy is also written into the run dir so each run's provenance is
   self-contained.

2. The framework runs `go build` every step (no source-tree hash caching).
   If the build is silently stale, your `go build` cache is stale —
   `go clean -cache` and retry.

**Custom run-dir for a one-off debug session:**
```sh
go run ./cmd/tfv2 run --run-dir /tmp/myrun issues-repro/issue_5672/
```

---

## Common gotchas

1. **`passthrough_env` containing tfexec-prohibited names.** `tfexec`
   reserves a fixed set of env-var names it manages internally and rejects
   at runtime: `TF_APPEND_USER_AGENT`, `TF_IN_AUTOMATION`, `TF_INPUT`,
   `TF_LOG`, `TF_LOG_PATH`, `TF_REATTACH_PROVIDERS`, `TF_DISABLE_PLUGIN_TLS`,
   `TF_SKIP_PROVIDER_VERIFY`. Listing any of these in your test.yaml's
   `passthrough_env` will make the runner fail when it calls
   `tfexec.SetEnv()`. Don't do it.

2. **`passthrough_env` containing `DATABRICKS_*` names.** Rejected at parse
   time — the profile field is the only sanctioned auth channel. If a test
   genuinely needs different credentials, use a different `~/.databrickscfg`
   profile, not env vars.

3. **Profile must exist in `~/.databrickscfg` (case-sensitive section
   name).** The framework eagerly validates section existence at parse
   time; a typo like `account_aws` instead of `ACCOUNT_AWS` fails fast with
   a clear error. If your `.databrickscfg` lives elsewhere, set
   `DATABRICKS_CONFIG_FILE` in your shell and the framework propagates it.

4. **`version: local` outside a provider checkout, with no `--repo` /
   `TFV2_REPO`.** Auto-discovery walks up from cwd looking for a
   `go.mod` whose `module` line is `github.com/databricks/terraform-provider-databricks`
   (the framework's own `testframeworkV2/go.mod` is intentionally
   skipped — it declares a sub-module path). If the walk hits the
   filesystem root without a match, the framework returns a clear error
   pointing you at `--repo` / `TFV2_REPO`. The check fires only when at
   least one step uses `version: local`; pure released-semver tests
   never look for a repo root.

5. **`provider_config { workspace_id = ... }` block in `main.tf`.** This is
   the exact schema shape that triggered issue #5672. Don't include it
   unless your test specifically intends to exercise it. Plain
   `provider "databricks" {}` covers most regression scenarios.

6. **`DATABRICKS_HOST`/`DATABRICKS_TOKEN` env vars set in your shell.**
   Stripped by default — the framework only propagates
   `DATABRICKS_CONFIG_PROFILE` and `DATABRICKS_CONFIG_FILE`. This is a
   feature, not a bug; profile auth is the only auth channel in tests.

7. **`requires.cloud: gcp` when you mean "this bug only showed up on the GCP
   profile I happened to test against".** Run the test on AWS or Azure too
   first — most bugs are host-agnostic. Use `cloud: any` unless you've
   verified the bug is cloud-specific.

8. **`error_regex` matching only one cloud's inner error.** Anchor on the
   outer error wrapper. Inner SDK errors vary by auth method; outer
   terraform diagnostics are stable across clouds.

9. **Committing `.terraform/`, `.terraform.lock.hcl`, `*.tfstate*`.**
   Already gitignored at the framework level, but worth saying: the
   framework writes these only into the per-run dir under
   `~/.testframeworkv2/runs/...`, never into your test source dir. If they
   show up in `git status`, you've run `terraform` manually in the wrong
   directory.

10. **Wide `include` filter in `.terraformrc`.** Don't write
    `include = ["*/*"]`. The framework's narrow
    `registry.terraform.io/databricks/*` is load-bearing — wider patterns
    break `hashicorp/google` and similar.

11. **`hc-install` / `tfinstall` auto-install of terraform.** Framework
    refuses to use this — the public HashiCorp signing key is currently
    expired in the default flow. Terraform binary discovery is
    user-controlled (`--terraform-bin` / `TFV2_TERRAFORM_BIN` / `$PATH`).
    See DESIGN.md §10/G8.

12. **Running multiple tests in <1 second.** Each run dir gets a 4-char
    `crypto/rand` hex suffix to prevent collisions. The runtime tree at
    `~/.testframeworkv2/runs/` accumulates monotonically — clean it up
    manually when needed.

---

## Adding dependencies

testframeworkV2 has its own `go.mod` (DESIGN.md §12 explains why). Direct
dependencies are kept minimal — currently `gopkg.in/yaml.v3` and
`github.com/hashicorp/terraform-exec`.

**Acceptance policy for new direct deps:**

- Published by **HashiCorp** (`github.com/hashicorp/...`) — accepted without
  further discussion.
- Already a direct dep of the **main provider repo's** `go.mod` (the parent
  `terraform-provider-databricks` module) — accepted. Reusing what the main
  repo already vets keeps the supply-chain footprint flat.
- Anything else — **file an ask via tech-lead first.** Don't slip in
  third-party tools. Transitive deps inherited from sanctioned direct deps
  are fine; we only gate at the direct-dep boundary.

Adding a sanctioned dep is a normal `go get`:

```sh
cd testframeworkV2/
go get github.com/hashicorp/some-package@v1.2.3
go mod tidy
```

**Hard prohibitions:**

- **Never add `hc-install` or `tfinstall`** even though they're HashiCorp.
  See DESIGN.md §10/G8 — these auto-install terraform via a flow whose
  signing key is currently expired. Binary discovery stays user-controlled.
- **Never import `internal/acceptance`** from the main provider repo.
  That package's `init()` calls `os.Setenv("TF_LOG", "DEBUG")` globally and
  has other process-wide side effects. The framework's separate `go.mod`
  exists partly to make this an explicit boundary.
- **Don't import the parent provider's source modules.** The framework
  treats the provider as an opaque binary — the whole point is to test
  multiple versions, including ones from before any `internal/...` symbols
  existed.

---

## Where to ask questions

- **Framework internals**: read `DESIGN.md` cover-to-cover. The 16 sections
  are cross-referenced; the table of contents at the top points everywhere.
- **A specific test misbehaving**: file a Jira issue with the run dir
  contents (logs + workdir) attached.
- **A new design pivot you think is needed**: open a PR with a `DESIGN.md`
  patch alongside the code change. Past pivots (override-merge,
  cloud-portable regex, run-dir suffixing) all started this way.
