# Contributing to testframeworkV2

This guide walks you from "I want to write a regression test for a Databricks
provider bug" to "test passes locally on my profile". For framework internals,
see `DESIGN.md`. For a quick "what is this thing", see `README.md`.

The running example throughout is `account/test1_issue_5672/`, which
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

Build the CLI once:

```bash
cd testframeworkV2/
go build -o tfv2 ./cmd/tfv2/
./tfv2 version            # prints "tfv2 dev"
```

---

## Step 1 — Pick a directory

Tests are grouped by the **profile level** they require. The framework reads
the level from your test's `requires:` block and skips the test cleanly if
the named profile doesn't match.

| Subdir | Use when your bug requires… |
|---|---|
| `testframeworkV2/account/`  | An account-level profile (`host = https://accounts.{cloud}.databricks.com`). |
| `testframeworkV2/workspace/` | A workspace-level profile (any cloud). |
| `testframeworkV2/ucws/`     | Unity Catalog on a workspace profile. |
| `testframeworkV2/ucacct/`   | Unity Catalog on an account profile. |

Issue #5672 is account-only (the post-Read hook only fires against an account
host) → `account/test1_issue_5672/`. If you're testing, say, a workspace-level
`databricks_grant` regression → `workspace/your_test_name/`.

Test directory name convention: `<short_descriptor>` or `<issue_id>_<descriptor>`.
Slug-style, lowercase, underscores: `^[a-z0-9_-]+$`.

```
testframeworkV2/account/test1_issue_5672/
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
# testframeworkV2/account/test1_issue_5672/main.tf

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
two optional ones, and a list of steps. Walk-through using `account/test1_issue_5672/test.yaml`:

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

Step 4 is the most important one. It's the only step that exercises code that
isn't in any released tag yet.

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

```bash
cd testframeworkV2/

# Run a single test:
./tfv2 run account/test1_issue_5672/

# Override the terraform binary if it's not on PATH:
./tfv2 run --terraform-bin /opt/terraform/1.5.7/terraform account/test1_issue_5672/

# Disable cleanup destroy (preserves state for inspection):
./tfv2 run --no-cleanup account/test1_issue_5672/

# Verbose framework logs to stderr (does NOT enable terraform's TF_LOG):
./tfv2 run --verbose account/test1_issue_5672/
```

Exit codes:
- `0` — every step passed (or test was skipped due to `requires` mismatch).
- `1` — at least one step did not pass.
- `2` — bad flags / missing arguments.

Per-step logs land at `~/.testframeworkv2/runs/<test>-<ts>-<rand>/`. The run
directory is preserved on disk after the test ends — open it with your editor
to inspect captured stdout/stderr per step, the generated `.terraformrc`, the
generated `_tfv2_versions_override.tf`, and the `terraform.tfstate` that
carried across steps.

---

## Step 5 — Debugging

**Test is being skipped unexpectedly.** Check `requires:` against your profile:
```bash
# Inspect the named profile's host:
grep -A 5 "^\[ACCOUNT_AWS\]" ~/.databrickscfg
```
If your profile has `host = https://accounts.cloud.databricks.com` it's an AWS
account-level profile → matches `cloud: aws` (or `any`) + `level: account`.

**Test fails on a step you expected to pass.** Open the per-step stderr log:
```bash
ls ~/.testframeworkv2/runs/test1_issue_5672-*/
cat ~/.testframeworkv2/runs/test1_issue_5672-*/step-1-*.stderr.log
```
The framework's curated subprocess env (DESIGN.md §10/G6) intentionally strips
`TF_LOG`. If you need terraform's debug logs, add a temporary `TF_LOG=DEBUG`
to `passthrough_env` in your test.yaml — but DON'T commit that.

**Wrong provider version is being served.** Verify the lock file matches the
expected version:
```bash
cat ~/.testframeworkv2/runs/<run-id>/workdir/.terraform.lock.hcl
```
The framework wipes `.terraform/` and `.terraform.lock.hcl` between every step,
so each step's lock should reflect ONLY that step's pinned version.

**Local build (`version: local`) isn't picking up my changes.** The framework
runs `go build` every step (no source-tree caching). If the build is silently
stale, your `go build` itself is — try `go clean -cache` and retry.

**Custom run-dir for a one-off debug session:**
```bash
./tfv2 run --run-dir /tmp/myrun account/test1_issue_5672/
```

---

## Common gotchas

1. **TF_LOG=DEBUG leaking from your shell.** The curated subprocess env strips
   `TF_LOG`. Good. But if you want it for one debug run, use the
   `passthrough_env` field in your test.yaml — never `os.Setenv` from
   surrounding test infrastructure.

2. **`DATABRICKS_HOST`/`DATABRICKS_TOKEN` env vars in your shell.** Stripped
   by default — the framework only propagates `DATABRICKS_CONFIG_PROFILE` and
   `DATABRICKS_CONFIG_FILE`. If your tests genuinely need different
   credentials, use a different profile in `~/.databrickscfg`, not env vars.

3. **`provider_config { workspace_id = ... }` block in main.tf.** This is the
   exact schema shape that triggered issue #5672. Don't include it unless your
   test specifically intends to exercise it. Plain `provider "databricks" {}`
   covers most regression scenarios.

4. **Wide `include` filter.** Don't be tempted to write `include = ["*/*"]`
   anywhere. The framework's narrow `registry.terraform.io/databricks/*` is
   load-bearing — wider patterns break `hashicorp/google` and similar.

5. **`requires.cloud: gcp` when you mean "this bug only shows up on a GCP
   profile I happened to test against".** Run the test on AWS / Azure too —
   most bugs are host-agnostic. Use `cloud: any` unless you've actually
   verified the bug is cloud-specific.

6. **`error_regex` matching only one cloud's inner error.** Anchor on the
   outer error wrapper. Inner SDK errors vary; outer terraform diagnostics
   are stable.

7. **Committing `.terraform.lock.hcl` from a run dir.** Don't. Lock files
   carry single-platform `h1:` hashes and break cross-arch CI. Per-test
   directories shouldn't have one (the framework writes them only into the
   run dir, which is outside the repo).

8. **`hc-install` / `tfinstall` auto-install of terraform.** Framework refuses
   to use this — the public hashicorp signing key is currently expired in
   their default flow. If you're tempted to add an "automatic terraform
   downloader" feature, please first read DESIGN.md §10/G8 and TL's binding
   constraint B7.

9. **Running multiple tests in <1 second.** Each run gets a 4-char
   `crypto/rand` hex suffix on the run dir. If you somehow generate a
   collision, file an issue with the random number generator company.

---

## Adding dependencies

testframeworkV2 has its own `go.mod` (DESIGN.md §12 explains why). Adding deps
is a normal `go get`:

```bash
cd testframeworkV2/
go get github.com/some/package@v1.2.3
go mod tidy
```

Guidelines:
- Pull in deps sparingly. Right now we have 1 direct dep (`gopkg.in/yaml.v3`)
  + `terraform-exec`. Less is more — every dep is a future supply-chain audit.
- **Never add `hc-install` or `tfinstall`.** Terraform binary discovery is
  user-controlled (`--terraform-bin` / `TFV2_TERRAFORM_BIN` / `$PATH`).
- **Never import `internal/acceptance`** from the parent provider repo.
  That package's `init()` calls `os.Setenv("TF_LOG", "DEBUG")` globally and
  has other process-wide side effects we deliberately do not inherit.
- If you need profile-loading helpers, the framework's
  `internal/profile/` package owns that — extend it rather than reaching into
  the parent provider for similar code.

---

## Where to ask questions

- **Framework internals**: read `DESIGN.md` cover-to-cover. The 16 sections
  are cross-referenced; the table of contents at the top points everywhere.
- **A specific test misbehaving**: file a Jira issue with the run dir
  contents (logs + workdir) attached.
- **A new design pivot you think is needed**: open a PR with a `DESIGN.md`
  patch alongside the code change. Past pivots (override-merge,
  cloud-portable regex, run-dir suffixing) all started this way.
