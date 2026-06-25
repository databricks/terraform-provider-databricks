# Fakebricks acceptance-test tracking

Tracks how every Terraform acceptance test behaves against the in-process
**fakebricks** backend, so that after fixing something in fakebricks you can
retest exactly the affected set and have passing tests marked done automatically.

## Files
- **`tests.jsonl`** — source of truth, one record per acceptance test:
  `id, package, test, level (WS|UC|ACCT), status (PASS|FAIL|SKIP|TIMEOUT),
  cause, endpoint, detail, fix_status (open|fixed|pass|skip), last_run, fixed_at`.
- **`report.md`** — generated human view: open failures grouped by cause (with
  checkboxes), the fixed list, and skipped (gated) tests. Regenerate with `report`.
- **`retest.py`** — the tool (build / list / run / report).

## How a fix cycle works
1. Implement something on the `fakebricks-tf` branch (the repo's `go.mod` already
   `replace`s `github.com/databricks/cli` with that worktree).
2. Retest just the tests blocked by what you fixed — select by endpoint, cause, or package:
   ```
   python3 fakebricks-acc/retest.py run --match "/scim/v2/Users"
   python3 fakebricks-acc/retest.py run --cause WORKSPACE_ID
   python3 fakebricks-acc/retest.py run --package catalog
   python3 fakebricks-acc/retest.py run --test workspace/TestAccGlobalInitScriptResource_Create
   ```
   Each selected test is run in isolation (own process → fresh fake → no
   cross-test state). Any that now pass flip to `status=PASS`, `fix_status=fixed`
   (with `fixed_at`); ones still failing get their latest cause/detail.
3. `python3 fakebricks-acc/retest.py report` to refresh `report.md`.

Preview a selection without running: add `--dry-run` to `run`, or use `list`:
```
python3 fakebricks-acc/retest.py list --cause 501_NOT_IMPLEMENTED --match scim
```

## Selectors (AND-combined)
`--match SUBSTR` (over `id cause endpoint detail`), `--package PKG`, `--cause CAUSE`,
`--status STATUS`, `--test ID`, `--open` (only not-yet-fixed). `run` with no
selector retests everything still open.

## Cause buckets
`501_NOT_IMPLEMENTED` (endpoint is a placeholder — see `endpoint`), `404_NO_HANDLER`,
`NOT_SEEDED` (resource/fixture not present), `WORKSPACE_ID` (workspace_id
resolution/mismatch semantics), `REQUEST_PARSE`, `POST_APPLY_DRIFT`,
`NOTEBOOK_DBC`, `GROUP_ROLES`, `CHECK_ASSERTION`, `OTHER`.

`SKIP` records aren't fakebricks failures — the test self-skips because it needs
external fixture env vars (e.g. `TEST_DEFAULT_WAREHOUSE_ID`, `TEST_WORKSPACE_ID`)
or a different provider level (account/UC-account).

## Rebuilding from a fresh full sweep
```
python3 fakebricks-acc/retest.py build --from-tsv <sweep>/results_clean.tsv --results-dir <sweep>/results.d
```

## Env
Runs set `CLOUD_ENV=aws`, `DATABRICKS_FAKE=1`, `TEST_ENVIRONMENT_TYPE=UC_WORKSPACE`
(or `ACCOUNT` for `ACCT` level), `TF_ACC_TERRAFORM_PATH` (default
`/usr/local/bin/terraform`), and **`GOFLAGS=-mod=mod`**. The fake server + env
wiring lives in `internal/acceptance/fake.go`.

## Important: vendoring vs. live fakebricks edits
`go.mod` has `replace github.com/databricks/cli => <abs path>/.worktrees/fakebricks-tf`.
This repo also keeps a (git-ignored) `vendor/` that `make build`/`make lint`
populate. With a populated `vendor/`, a bare `go test`/`make build` uses the
**vendored snapshot** and will NOT see edits you make in the fakebricks-tf
worktree. `retest.py` sets `GOFLAGS=-mod=mod` to read the replace live — so after
editing fakebricks, just rerun `retest.py`. If you run `go test`/`make` directly,
either pass `-mod=mod` or re-run `make vendor` after each fakebricks change.

To point at a different fakebricks worktree, change the `replace` path in `go.mod`.
