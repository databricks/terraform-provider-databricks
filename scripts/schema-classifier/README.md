# schema-classifier

Classifies deltas between two `terraform providers schema -json` dumps as **breaking** or **non-breaking** and exits non-zero on any breaking change. Invoked by `scripts/diff-schema.sh` after the existing `jd` JSON diff and gated by the `breaking-schema-check` workflow on every PR.

## Scope

The classifier looks **only** at what Terraform Core sees in the JSON schema dump. Provider-side concepts that are not part of that dump (`ForceNew`, `ConflictsWith`, `ExactlyOneOf`, `ValidateFunc`, `DiffSuppressFunc`, `Default`, etc.) are intentionally out of scope. Catching those would require a different dump source (a custom Go-side dumper) and is tracked as future work.

## Rule taxonomy

Every rule below is encoded in `classify.go` and pinned by a unit test in `classify_test.go`. The rationale column says why we landed on the verdict.

### Breaking — these fail the `breaking-schema-check` job

| Rule | Path-level signal | Rationale |
|---|---|---|
| `ResourceRemoved` / `DataSourceRemoved` | A resource or data source disappears from the dump | Configs referencing it fail to plan |
| `AttributeRemoved` | An attribute disappears from a resource / data source / nested block | Configs setting it fail with `Unsupported argument` |
| `RequiredAttributeAdded` | A new attribute appears with `required: true` | Existing configs that don't set it fail with `Missing required argument` |
| `OptionalToRequired` | Attribute went from `optional: true` to `required: true` | Configs missing it now error |
| `ComputedOnlyToRequired` | Attribute was computed-only (Computed=true, Optional=false, Required=false) and is now required | Users couldn't set it before; now they must. Existing configs that relied on the API default error |
| `BecameComputedOnly` | A previously-settable attribute is now computed-only | Configs that set it fail (the schema rejects writes) |
| `ComputedRemoved` | A still-settable attribute loses `computed: true` | Drift behavior changes — values previously sourced from the API now show as permanent diffs |
| `SensitiveRemoved` | `sensitive: true` → `sensitive: false` | Un-masks previously hidden values in plan output; secrets may leak to CI logs / screenshots / session recordings |
| `TypeChanged` | An attribute's cty `type` JSON differs, OR `nested_type` is added/removed on an existing attribute | HCL value coercion breaks (e.g. `string` → `number`, `list(string)` → `list(number)`). Sub-attribute changes within an existing nested type recurse via the normal rules below — only an overall shape change emits `TypeChanged`. |
| `BlockTypeRemoved` | A nested block disappears | Configs declaring the block fail with `Unsupported block type` |
| `RequiredBlockTypeAdded` | A new nested block appears with `min_items > 0` | Required-by-definition; existing configs without it fail |
| `NestingModeChanged` | `nesting_mode` changes (e.g. `list` → `set`, `single` → `list`) | Plan diffs and/or HCL syntax change |
| `MinItemsIncreased` | A block's `min_items` increased | Previously-valid configs may now fall short |
| `MaxItemsDecreased` | A block's `max_items` decreased (or a previously-unbounded cap was introduced and is lower than reality) | Previously-valid configs may now exceed |

### Non-breaking — surfaced for awareness, not blocking

| Rule | Path-level signal | Rationale |
|---|---|---|
| `ResourceAdded` / `DataSourceAdded` | A new resource or data source appears | Pure addition |
| `OptionalAttributeAdded` | A new attribute appears with `optional: true` | Pure addition |
| `ComputedAttributeAdded` | A new attribute appears as computed-only | Pure addition; user can't set it |
| `RequiredToOptional` | Attribute went from `required: true` to `optional: true` | Strictly looser — every previously-valid config remains valid |
| `ComputedAdded` | A still-settable attribute gained `computed: true` | Strictly looser drift behavior |
| `SensitiveAdded` | `sensitive: false` → `sensitive: true` | Output is now masked; no config validity impact |
| `BlockTypeAdded` | A new nested block with `min_items=0` (or unset) | Pure addition |
| `MaxItemsRelaxed` | A block's `max_items` increased or was removed | Strictly looser |
| `MinItemsDecreased` | A block's `min_items` decreased | Strictly looser |
| _Description-only changes_ | `description` / `description_kind` / `deprecated` toggles | Cosmetic / advisory only |

### What we do not detect (by design — out of scope)

These are real concerns but invisible in `terraform providers schema -json`. Reviewers must catch them manually until/unless we replace the dump source.

| Concern | Why we miss it |
|---|---|
| `ForceNew` flip on an existing attribute | Not in the schema dump; SDKv2 emits it as plan-time `RequiresReplace`, Plugin Framework as a `planmodifier.RequiresReplace()`. Terraform Core never needs the flag, so the dump strips it. |
| `ConflictsWith` / `ExactlyOneOf` / `AtLeastOneOf` / `RequiredWith` changes | Provider-side validation, not in the dump |
| `ValidateFunc` / `ValidateDiagFunc` tightening | Provider-side, not in the dump |
| `Default` value changes | Materialized at apply time on the provider side, not in the dump |
| `DiffSuppressFunc` changes | Provider-side, not in the dump |
| Resource rename | Indistinguishable from delete + add in the dump |

## Bypass

To land an intentional breaking change, add this line anywhere in the PR description and push a new commit so the `synchronize` event re-captures the body:

```
ALLOW_SCHEMA_BREAKING_CHANGE=true
```

The workflow reads the directive from `${{ github.event.pull_request.body }}` (no GitHub API call, so no IP-allowlist issues). Editing the PR description alone does not re-trigger the workflow — a new commit is required.

## Local use

```sh
make diff-schema            # prints the jd JSON diff and the classifier report; exits non-zero on breaking
ALLOW_BREAKING=1 make diff-schema   # same, but always exits 0 (handy for inspecting the report)
```

Or run the binary against pre-dumped schemas:

```sh
go build -mod=mod -o /tmp/schema-classifier ./scripts/schema-classifier
/tmp/schema-classifier --base BASE.json --head HEAD.json [--format text|markdown] [--allow-breaking]
```

Exit codes: `0` = no breaking changes (or `--allow-breaking` set); `1` = breaking change(s) detected; `2` = invocation error.

## Maintaining the taxonomy

When changing rules, update three things in lockstep:

1. The classifier in `classify.go` (the rule itself).
2. A passing test row in `classify_test.go` (so the verdict is pinned).
3. The matching row in this README.

The test file is the executable source of truth; this README is the human-readable spec.
