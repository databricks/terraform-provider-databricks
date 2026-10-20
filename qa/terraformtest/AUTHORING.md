# Writing YAML replay tests

Use replay tests to exercise provider behavior through real Terraform and the
production SDKv2/Plugin Framework mux. HTTP fixtures replace Databricks API calls;
the provider and Terraform are not mocked. See the [format reference](README.md)
for supported fields.

## Choose the behavior first

Start with a small, common configuration. Add an update step when the resource
supports updates; a create/update test can cover both scenarios without a
separate duplicate create-only test. Keep one regression or behavior per file,
and include the issue link when reproducing a bug. These follow HashiCorp's
[basic, update, and regression patterns](https://developer.hashicorp.com/terraform/plugin/testing/testing-patterns).

Focus on provider-owned behavior: request construction, resource-specific client
logic, API response conversion, defaults, retries/polling, and cleanup. Use only
the HCL fields needed to exercise it. Do not turn each test into an exhaustive
resource schema or Terraform-internals test.

## Add a test

Use a dedicated fixture folder in the resource's package, for example
`scim/testdata/terraform/`. Add one wrapper for that folder, such as
`scim/terraform_test.go`:

```go
package scim_test

import (
    "testing"

    "github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

func TestTerraform(t *testing.T) {
    terraformtest.Run(t, "testdata/terraform")
}
```

Use the external test package (`scim_test`, not `scim`): the runner imports the
provider, which already imports resource packages. Do not call `t.Parallel` in
the wrapper or its ancestors. Discovery is non-recursive; each `.yaml`/`.yml`
file becomes a subtest, so further files need no new Go functions.

Save this complete example as `scim/testdata/terraform/group_create.yaml`:

```yaml
version: 1
host_metadata:
  workspace_id: "12345"
steps:
  - name: create
    config: |
      resource "databricks_group" "this" {
        display_name = "Example group"
      }
    calls:
      - request:
          method: POST
          path: /api/2.0/preview/scim/v2/Groups
          body:
            json:
              displayName: Example group
              schemas:
                - urn:ietf:params:scim:schemas:core:2.0:Group
              entitlements:
                - {}
        response:
          status: 201
          body:
            json:
              id: group-id
      - request:
          method: GET
          path: /api/2.0/preview/scim/v2/Groups/group-id
          query:
            attributes: ["displayName,externalId,entitlements"]
          body: null
        response:
          status: 200
          body:
            json:
              id: group-id
              displayName: Example group
      - request:
          method: GET
          path: /api/2.0/preview/scim/v2/Groups/group-id
          query:
            attributes: ["displayName,externalId,entitlements"]
          body: null
        response:
          status: 200
          body:
            json:
              id: group-id
              displayName: Example group
    state:
      resources:
        databricks_group.this:
          attributes:
            - path: [id]
              equals: group-id
            - path: [acl_principal_id]
              equals: groups/Example group
destroy:
  calls:
    - request:
        method: GET
        path: /api/2.0/preview/scim/v2/Groups/group-id
        query:
          attributes: ["displayName,externalId,entitlements"]
        body: null
      response:
        status: 200
        body:
          json:
            id: group-id
            displayName: Example group
    - request:
        method: DELETE
        path: /api/2.0/preview/scim/v2/Groups/group-id
        body: null
      response:
        status: 204
        body: null
```

Install Terraform on `PATH` (CI uses 1.13.5), or set `TF_ACC_TERRAFORM_PATH`.
From the repository root, run:

```sh
go test -short -v ./scim -run '^TestTerraform$/^group_create$'
go test -race -short ./scim -run '^TestTerraform$'
```

No provider binary installation, credentials, or `TF_ACC` are needed. For an
update scenario, append another step with the same resource address and changed
HCL, add that step's refresh/update/read fixtures, and update destroy's fixtures
to reflect the last step. Terraform state carries over between steps, not files.

Omit the provider block unless the scenario needs explicit routing settings,
such as `account_id`. The runner supplies the local host and authentication.
For a data source reading an existing API object, supply its GET responses in
fixtures; no server-side setup or initial Terraform state is needed.

## Decide what to assert

Leave HTTP ordering disabled for ordinary scenarios: Terraform may run
independent resources concurrently. For a regression that specifically depends
on call order, set `enforce_call_order: true` at the test root. It applies to
every step and destroy, excluding reusable host metadata. Keep each observed
call in a future recorded transcript even when enforcement is disabled;
identical request signatures still consume responses FIFO.

Use `plan` when the behavior is the proposed diff: create versus update,
replacement order, computed values still unknown, or an unchanged configuration.
Use full instance addresses for multi-resource cases. Select only the relevant
actions and before/after paths; do not copy every provider attribute into an
expectation. The [plan format reference](README.md#plan-assertions) documents
output checks and `plan_only`, including the harness's empty-plan restriction.
Plan+apply assertions inspect the same saved plan that gets applied, without
disabling Terraform's post-apply consistency checks.

Plan-only HTTP reads belong in the step's ordinary `calls`. Plans do not create
the proposed resources, but initial-state or earlier applied resources still
require destroy fixtures. Plan-only state checks inspect the refreshed prior
state, not the unapplied planned values.

HashiCorp's [resource state-check guide](https://developer.hashicorp.com/terraform/plugin/testing/acceptance-tests/state-checks/resource)
focuses state checks on computed attributes. Follow that default: check IDs,
provider-derived values, data-source outputs, and unconfigured optional/computed
values whose defaults or API-to-state conversions matter to the scenario.

In the example, `id` comes from the response and `acl_principal_id` is derived
by the provider. `display_name` is already configured in HCL and checked in the
outgoing HTTP body; repeating every configured value in `state` adds little.
Terraform performs plan/apply consistency checks and the testing framework
requires an empty post-apply plan. HashiCorp documents the Framework's
[configured-value consistency rules](https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#terraform-data-consistency-rules).

This is a testing guideline, not a restriction on YAML paths. A targeted
configured-value check can still be useful for a specific regression, especially
with legacy SDKv2 diff suppression or normalization; explain which behavior it
protects. Do not assume SDKv2 enforces every Framework consistency rule.

State checks are optional and partial. Omit `state` when the HTTP transcript and
normal Terraform checks already prove the behavior. Do not list unrelated
resources, snapshot all attributes, or assert that new optional fields are
absent. Use typed primitive/null expectations and explicit paths; numeric list
indices are not suitable for unordered sets. Destroy's empty-state check is
mandatory and needs no YAML state assertion.

## Keep fixtures trustworthy

- Write requests and responses from the API contract and the scenario, not just
  whatever makes the provider pass. The stateless server does not validate their
  relationship: a PUT fixture verifies what the provider sends, while a later
  GET returns exactly its supplied response. Review both sides independently.
- Use fixed, fake names and IDs. Each file has isolated state and a local server,
  so the random names in HashiCorp's examples are unnecessary here. Never commit
  credentials, real customer data, or machine-specific URLs.
- Keep payloads small but realistic: include required API fields and values
  needed for the behavior. Do not replace all responses with `{}` just to avoid
  state checks. Use `host_metadata` for repeated client initialization calls.
- Include every consumable request, including plan refreshes, post-apply reads,
  retries, and destroy. The two identical create-step GETs above are deliberate.
  Match JSON semantically, not byte-for-byte; whitespace and object-key order
  do not matter, but extra fields and array order do.
- Leave behavior-specific fields exact. For an incidental volatile value such
  as a request ID, use explicit [`request.matchers`](README.md#test-definition)
  with `any_string`. It requires a present string; missing values and JSON null
  still fail. Do not mask names, IDs, or other values that the scenario should
  verify, and do not use `"****"` as a wildcard (it is just a literal string).
- For polling/retry scenarios, check meaningful requests and outcomes, not
  elapsed time or incidental scheduling.
- Independent request signatures can interleave. Repeated identical signatures
  consume responses FIFO, so preserve their response order. Do not assume a
  global or per-resource request order is enforced.
- Overlapping matcher signatures consume the first matching unconsumed fixture,
  not the most specific one. Keep distinct behavior-specific fields exact so a
  broad matcher does not consume the response intended for another request.
- Diagnose failures before editing fixtures. An unexpected call may expose a
  provider regression; an unused call may mean an expected operation vanished.
  Do not automatically bless a changed transcript. Temporarily break the
  behavior-specific fixture or assertion to confirm the test detects it, then
  restore it before submitting a normal test.

Automatic cleanup is part of the test. Normal destroy may refresh first; cleanup
after an interrupted apply can have a different transcript. Missing cleanup
fixtures fail too. Empty Terraform state plus consumed DELETE fixtures verifies
the replayed cleanup flow, not the state of an independently simulated service.

## Current limits

The [format reference](README.md#mvp-scope) lists unsupported features. Recording,
import steps, and strict attribute matching need further engine support. Use the
documented YAML fields rather than copying arbitrary HashiCorp TestStep options.
Do not silently skip an assertion or treat an unexpected failure as a passing test.

## Test expected failures

Use [`expect_error`](README.md#expected-errors) on an individual step for a
specific Terraform/provider diagnostic. Local validation failures can use
`calls: []`; server-side failures still need normal HTTP error responses.
Keep the expression narrow so a different failure cannot accidentally pass.
The engine verifies HTTP consumption independently of the error assertion.

The minimal [validation-error example](testdata/expected_errors/group_validation.yaml)
has one failing configuration, its expected diagnostic, and no HTTP calls.
The [HTTP error example](testdata/expected_errors/api_error.yaml) covers an API
failure instead.

You can create a resource, exercise an invalid update, then let automatic destroy
clean it up using the last successful configuration. To check state after an
expected failure, follow it with a successful step and place state checks there;
state checks directly on the failing step are not supported.

## Start from existing state

Add [`initial_state.checkpoint`](README.md#initial-state) to reference a
Terraform-generated state file next to the test. Keep it as an opaque artifact;
do not hand-build state JSON or edit IDs, schema versions, or provider-private
data. For migration regressions, capture the state with the older provider
version under test and keep that artifact fixed. Record how it was generated.

The server remains stateless: a checkpoint and HTTP fixtures describe different
things. Supply the existing object's read responses, update calls, and destroy
calls. You do not need a create step for a resource already in the checkpoint.
The [update showcase](testdata/initial_state/group_update.yaml) restores a real
group checkpoint, updates it in a single step, and destroys it.
The [expected-error showcase](testdata/initial_state/rejected_update.yaml)
proves cleanup still runs after rejecting the first update.

Review checkpoints for secrets, including encoded private data. Generate
secret-free scenarios rather than editing a checkpoint to sanitize it. These
tests are hermetic; fixture preparation here uses the real provider and Terraform
against local HTTP fixtures, not a live integration test. A recorder will later
handle environment and initial-state templates with separate cleanup ownership.

The repository ignores `*.tfstate` by default to protect ordinary local state.
After reviewing a fixture for secrets, add that exact checkpoint explicitly with
`git add -f path/to/checkpoint.tfstate`; do not relax the global ignore rule.
