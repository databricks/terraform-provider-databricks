# YAML Terraform lifecycle tests

For a complete example and advice on choosing scenarios, fixtures, and state
assertions, see [Writing YAML replay tests](AUTHORING.md).

Use one small Go wrapper per fixture folder:

```go
func TestTerraform(t *testing.T) {
    terraformtest.Run(t, "testdata")
}
```

The runner discovers every `.yaml` and `.yml` file directly in that folder
(non-recursively), relative to the test package. Each file becomes a Go subtest
named after its filename without the extension. For example,
`testdata/TestTerraformGroupLifecycle.yaml` appears as
`TestTerraform/TestTerraformGroupLifecycle` in test output and CI reports.
Adding a YAML file requires no additional Go test function. Empty folders and
duplicate filename stems fail instead of silently skipping tests.

Subtests run sequentially in filename order, each with independent Terraform
state and an HTTP fixture server. A failing lifecycle does not prevent later
files from running. Select a single file with normal Go filtering:

```sh
go test ./qa/terraformtest -run '^TestTerraform$/^TestTerraformGroupLifecycle$'
```

Only selected files are loaded and validated. `go test -list` lists the Go
wrapper, not dynamically discovered subtests; `go test -v` shows each subtest.
Tests run the actual muxed Databricks provider, supporting both SDKv2 and
Plugin Framework resources.

Install Terraform on `PATH`, or set `TF_ACC_TERRAFORM_PATH` to its executable.
The normal GitHub Actions tests job installs Terraform 1.13.5 with its wrapper
disabled so the test runner can invoke the CLI directly.
These hermetic tests run in the normal Go test suite, including `-short`;
they do not require `TF_ACC`, live credentials, or `CLOUD_ENV`. The runner
does not download Terraform. Client configuration ignores ambient credential
profiles and environment defaults; SDK HTTP traffic is restricted to the
test's local fixture server. Do not use `t.Parallel`: Terraform CLI discovery is
configured through process environment variables.

## Test definition

Use `version: 1`, optional `initial_state`, followed by `steps` and `destroy`.
Each step has a human `name`, literal HCL under `config`, `calls`, and optional
`state` expectations. State persists across steps. Destroy always runs using
the last successful configuration;
its HTTP exchanges belong under `destroy.calls`, and empty state is always
checked. The testing framework also attempts cleanup after an interrupted test.

Each exchange contains `request` and `response`. Requests specify `method`,
`path`, optional `query` (values are lists of strings), optional `headers`,
and `body`. Responses specify an explicit HTTP `status`, optional `headers`,
and `body`. Use `body: null` for no HTTP body and `body: {json: null}` for JSON
null. Other JSON values go under `body.json`; IDs are literal, without substitution.

Optional `request.matchers` annotations allow selected volatile strings to
vary without ignoring their presence or type. For example:

```yaml
request:
  method: POST
  path: /example
  body:
    json:
      request_id: captured-id
      name: fixed-name
  matchers:
    headers:
      X-Request-ID: any_string
    query:
      trace_id: any_string
    json:
      - path: [request_id]
        type: any_string
```

`any_string` accepts any present string, including empty, but not null or a
missing field. Header names are case-insensitive; matched headers and query
parameters must each have exactly one value. JSON paths use literal string
keys/nonnegative integer indices and must select existing string/null samples
in `body.json`. Other request values remain exact; unlisted headers are ignored
as usual. No wildcard string has special meaning inside payloads. Responses
remain literal, without capture/substitution. See the
[HTTP matcher reference](../httpfixture/README.md#volatile-request-values) for
root paths, sample overrides, validation, and overlapping signatures.

Runnable examples demonstrate [JSON body](testdata/matchers/group_create_body.yml),
[header](testdata/matchers/group_create_header.yml), and
[query](testdata/matchers/group_create_query.yml) matchers separately. Each creates
and destroys a group, with one request sample deliberately differing from the
actual value to demonstrate that only the annotated field is relaxed.

Matching uses the complete request signature. Independent signatures are
unordered; repeated identical signatures consume their responses FIFO.
Every HTTP call needs a consumable fixture, including Terraform's plan
refreshes and the testing framework's post-apply consistency reads.
Unexpected calls and unused fixtures fail with the step's index and name.

Optional test-level `enforce_call_order: true` requires consumable HTTP calls
in their listed order within every step, including destroy. The default is
`false`: Terraform may interleave independent resources. Ordering covers the
whole step, not per-resource lanes; a matching later call fails immediately
without consuming it. Reusable host metadata may occur between any calls.
Recorded call order can be preserved in the file without enabling enforcement.

Optional `host_metadata` supplies a reusable JSON response for
`GET /.well-known/databricks-config` throughout the test, keeping repeated
client initialization out of the consumable resource exchanges.

State assertions are entirely optional. `state.resources` selects Terraform
instance addresses to check, including data sources, modules, and indexed
resources; it is not a complete inventory. Listed resources must exist, but
unlisted resources are allowed. Each resource's `attributes` is a list of
targeted checks using HashiCorp's `tfjsonpath` and `knownvalue` helpers:

```yaml
state:
  resources:
    databricks_group.this:
      attributes:
        - path: [id]
          equals: group-id
        - path: [settings, enabled]
          equals: true
        - path: [items, 0, name]
          equals: first
        - path: [optional_field]
          equals: null
```

These paths illustrate the syntax, not the schema of `databricks_group`.
Paths start with an attribute name. String segments are literal object/map
keys; nonnegative integer segments are zero-based list indices. Quote numeric
map keys (`[labels, "0"]`) to distinguish them from indices. Dots and brackets
inside string keys are not parsed as path syntax. Index paths are suitable for
lists, not Terraform sets, whose element order is not guaranteed.

`equals` is required for each check and accepts only a string, boolean, finite
number, or null. Values are compared with their types preserved, so `42` is
different from `"42"`. `equals: null` requires a present null value; it does not
match a missing attribute. Whole-object, list, and set assertions are deferred.
Unselected values and extra attributes are allowed. A resource entry with no
attributes checks only that the resource exists.

Omitting `state`, using `state: {}`, or using `state: {resources: {}}` adds no
explicit state checks to a step. Terraform's normal apply and post-apply
plan-consistency checks still run. Following
[HashiCorp's testing guidance](https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns#basic-test-to-verify-attributes),
assert the attributes relevant to the behavior being tested, especially computed
values and response-to-state conversions, rather than snapshotting every field.

Destroy always requires an empty resource inventory, including data sources in
child modules. This is enforced by the engine; `destroy` accepts only `calls`,
with no configurable final-state assertion.

## MVP scope

Tests use handwritten fixtures and start with empty Terraform state unless an
initial checkpoint is supplied. YAML has strict field validation: recording/setup
templates and strict attribute matching are not yet supported. Unsupported fields fail
before any Terraform operations. The step/call/state layout is kept stable
so these capabilities can be added independently.

## Expected errors

An optional step-level `expect_error` is a nonempty Go regular expression
matched against Terraform/provider diagnostics, using HashiCorp's
`TestStep.ExpectError`. The step fails if the operation succeeds or its error
does not match. Use a specific diagnostic rather than a broad pattern:

```yaml
steps:
  - name: reject_reserved_name
    config: |
      resource "databricks_group" "this" {
        display_name = "admins"
      }
    expect_error: 'expected display_name to not be any of'
    calls: []
destroy:
  calls: []
```

Errors before HTTP calls need no resource exchanges. Errors returned by APIs
still use ordinary HTTP fixtures. Unexpected calls and unused fixtures fail
independently, even if their resulting diagnostic matches `expect_error`.
Later steps still run after an expected failure.

Automatic destroy remains mandatory and uses the last step without an expected
error, so an invalid final configuration cannot prevent cleanup. If all steps
expect errors, cleanup uses an empty default provider configuration. Tests that
partially create resources before failing must still supply destroy fixtures.
Provider aliases or routing that exist only in an expected-error configuration
are not available to that fallback cleanup configuration.

State checks on expected-error steps are rejected: HashiCorp skips post-apply
checks when an operation fails. Add a subsequent successful step to inspect
state instead. Invalid or empty error patterns fail before Terraform runs.

## Initial state

Reference a real Terraform-generated checkpoint, relative to the YAML file:

```yaml
initial_state:
  checkpoint: checkpoints/group.tfstate
```

The runner copies the original bytes into an isolated working directory using
HashiCorp's `TestStep.ConfigDirectory` API. It does not construct, decode, rewrite,
or merge internal state JSON. Terraform reads and upgrades state through the
provider as usual. The checkpoint is restored only before the first step; later
steps and automatic destroy use the resulting current state. The original file
is never updated. Missing, unreadable, and empty checkpoints fail before replay.

The checkpoint does not populate the HTTP server. Supply fixtures describing
the already-existing remote resources, including refresh reads. The test owns
the checkpoint's managed resources and must provide destroy fixtures for them,
including when the first step expects an error. Replay runs no setup templates.

Checkpoint tests use the default local backend and the provider address
`registry.terraform.io/databricks/databricks`. Explicit `backend` and `cloud`
blocks are rejected: an alternate state location could bypass the checkpoint
or access remote infrastructure. Provider configuration blocks remain optional;
if HCL includes a `terraform` block, it must include its own provider requirements,
as in ordinary replay tests.

Historical checkpoints work within the installed Terraform version's and current
provider's state-upgrade support; the runner cannot restore a removed resource
implementation. Keep historical fixtures unchanged and record their original
Terraform/provider versions. Review the entire checkpoint, including provider
private data, for secrets before committing it; prefer secret-free test scenarios.
Post-step assertions continue to use Terraform's supported inspection APIs,
not raw checkpoint comparisons.

## Plan assertions

Optional `plan` assertions inspect Terraform's supported plan API. For an apply
step, checks run on the saved pre-apply plan; Terraform then applies that same
plan. The normal post-apply empty-plan consistency checks remain enabled.
Omitting `plan` adds no custom assertions. `plan: {}` adds no expectations.

```yaml
plan:
  expect: non_empty
  changes:
    databricks_group.this:
      actions: [update]
      before:
        - path: [display_name]
          equals: Initial group
      after:
        - path: [display_name]
          equals: Updated group
  outputs:
    group_id:
      after:
        - path: []
          unknown: true
```

`expect` is optional for apply steps: `empty` requires no resource or output
changes; `non_empty` requires at least one. Omission ignores whole-plan emptiness.
`changes` maps full resource instance addresses, including modules, `count`, and
`for_each`. Quote addresses containing string keys. `outputs` maps root output
names. Nonempty maps select checks; unlisted addresses and attributes are allowed.
An explicitly empty `changes: {}` requires no resource changes, but permits output
changes. Likewise, `outputs: {}` requires no output changes. No-op entries do not
count as changes. Omitted maps impose no checks of that kind.

Optional `actions` lists use Terraform action names: `[no-op]`, `[create]`,
`[read]`, `[update]`, `[delete]`, `[forget]`, `[delete, create]`, or
`[create, delete]`. Replacement order matters. If omitted, an address need only
exist in the plan; any selected values still must match.

`before` and `after` are lists of path checks, using the same typed primitive/null
`equals` syntax as state checks. Resource paths start at an attribute; output
paths may use `[]` for the whole scalar output, or start with a map key/list index.
`unknown: true` replaces `equals` for an after value not yet known. Unknown is
not null, and missing is not null. Unknown ancestors also make selected children
unknown. Collection equality and set-element selection are not supported.

### Plan-only steps

Set step-level `plan_only: true` to skip apply. A successful plan-only step must
declare `plan.expect: empty` or `non_empty`, because HashiCorp's native harness
requires an emptiness expectation:

```yaml
plan_only: true
plan:
  expect: empty
  changes: {}
```

The harness runs a non-refresh plan followed by a refresh plan. Custom assertions
inspect the latter; API reads still consume ordinary step fixtures. Repeated
reads require repeated fixtures. With `expect: empty`, the harness also requires
the preceding non-refresh plan to be empty. This restricts drift scenarios where
the plan becomes empty only after refreshing; no hidden refresh/apply is added
to bypass that native restriction.

Optional plan-only `state` checks inspect the refreshed prior state exposed by
the plan API, not planned values. A plan does not persist that refreshed state or
apply proposed changes. Later steps continue from the last persisted state.
Automatic destroy still runs, including for tests containing only plan-only
steps or an initial checkpoint. After an apply, cleanup retains the last applied
configuration rather than a later preview's configuration. With no apply, it uses
the last successful preview configuration for provider routing.

`expect_error` cannot be combined with `plan` or `plan_only`; this prevents a
broad diagnostic pattern from accidentally accepting failed plan assertions or
harness emptiness checks. Use a separate ordinary diagnostic step instead.

Complete examples are in [testdata/plans](testdata/plans). They exercise
checkpoint update/apply, empty and nonempty previews, multiple resource instances,
known/unknown outputs, and automatic cleanup.

## Constraints for future recording

Recording is not implemented. These replay contracts must survive capture and
refresh:

- Keep every consumable exchange, including retries and repeated reads, in
  request-dispatch order. Assign sequence numbers before sending requests;
  response-completion order can differ under concurrency. Do not infer resource
  ownership from HTTP calls. Identical concurrent signatures with different
  responses remain ambiguous under FIFO replay and need explicit handling.
- Preserve whole-step boundaries, including all native consistency plans and
  destroy. Capture plan expectations at the same saved-plan hooks as replay:
  pre-apply for apply steps, post-refresh for plan-only steps.
- Preserve author-selected paths, omitted assertions, explicit empty maps, and
  unknown versus null. Do not turn partial assertions into whole-state snapshots
  or silently change `enforce_call_order` while refreshing fixtures.
- Capture reusable host metadata separately from consumable calls. Keep setup
  and background-environment cleanup outside the test transcript; initial-state
  resources belong to test destroy. Generate matching checkpoints and calls
  together without editing opaque checkpoint JSON or reusing stale live IDs.
- Decide retries/polling, secret and volatile-value handling before enabling
  live capture. Recording must be explicit, single-test, unavailable in CI, and
  must never be activated by a replay failure.
