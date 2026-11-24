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

Verify that the provider sends the correct data and handles representative API
responses. Omit SDK readback assertions whose only purpose is to check that the
backend persisted a preceding request. HTTP fixtures already check that request;
reading a handwritten response back does not validate the service. Keep checks
of provider-derived values and API-to-state conversion when those are the behavior
under test. Model out-of-band changes with later refresh responses.

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

Use an [import step](README.md#import-steps) to exercise the provider's importer.
Always choose `import.persist` explicitly. Use `false` to verify import against a
previously created resource in separate, discarded state; enable `import.verify`
and ignore only attributes outside the comparison. Use `true` with an explicit
ID to adopt a remote resource into the journey, then plan, update, and destroy it
normally. No create step is needed when fixtures describe an existing resource.
An initial-state snapshot alone does not exercise the importer. Examples show
[group import and update](testdata/imports/group.yaml) and
[separate-state app verification](testdata/imports/app_verification.yaml).

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
- When a volatile value has a format worth pinning, use a `pattern` matcher instead of `any_string`; see [Check formats and absent headers](#check-formats-and-absent-headers).
- For polling/retry scenarios, check meaningful requests and outcomes, not incidental scheduling. Bound elapsed time only when it is the behavior under test, such as a backoff; see [repeated calls and retry timing](#test-repeated-calls-and-retry-timing).
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

## Check formats and absent headers

`any_string` checks only that a value is present. A `pattern` matcher also checks its format: the value must be a present string that the Go regular expression matches in full, so `^` and `$` are implied. Patterns work on headers, query parameters and JSON strings. In `headers` and `query`, write `{pattern: <regular expression>}`; in a `json` entry, write `pattern:` in place of `type:`. The [pattern example](testdata/matchers/group_create_pattern.yml) checks `traceparent` and `User-Agent`:

```yaml
request:
  method: POST
  path: /api/2.0/preview/scim/v2/Groups
  headers:
    # Captured samples; the patterns replace exact comparison.
    traceparent: 00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01
    User-Agent: databricks-tf-provider/1.0.0 databricks-sdk-go/0.200.0 resource/group
  matchers:
    headers:
      traceparent: {pattern: '00-[0-9a-f]{32}-[0-9a-f]{16}-01'}
      User-Agent: {pattern: 'databricks-tf-provider/\S+ databricks-sdk-go/\S+ .*\bresource/group\b.*'}
```

A JSON string uses the same `path` as `any_string`:

```yaml
body:
  json:
    request_id: 0b6d4a3e-5c1f-4e2a-9d7b-8f3a2c1e0d4b
matchers:
  json:
    - path: [request_id]
      pattern: '[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'
```

Write the pattern from the value's contract, such as the W3C `traceparent` format, not from one captured sample. Quote patterns in single quotes: double-quoted YAML treats backslashes as escapes.

An exact header sample of `""` accepts a missing header and a header with an empty value alike. To require that the provider sends no such header, use the `absent` matcher, and leave the header out of `headers`. The [absent example](testdata/matchers/directory_create_absent.yml) pins a read that sends no body and so no `Content-Type`:

```yaml
request:
  method: GET
  path: /api/2.0/workspace/get-status
  query:
    path: [/Shared/replay-matchers]
  matchers:
    headers:
      Content-Type: absent
  body: null
```

`absent` applies only to headers: query parameters and JSON fields not in the sample already fail to match.

## Test calls to several hosts

Replay serves every host from one local server. By default the provider's host is that server, every call expects it, and replay rejects requests to any other host. When the behavior depends on which host a call reaches, such as an account host and a workspace's own host, or a workspace on a given cloud, name the hosts at the test root and pick the provider's host:

```yaml
version: 1
provider_host: account
hosts:
  account:
    url: https://accounts.cloud.databricks.com
    metadata:
      host_type: account
      cloud: AWS
      account_id: replay-account
  workspace:
    url: https://replay-workspace-1002.cloud.databricks.com
    metadata:
      host_type: workspace
      cloud: AWS
      workspace_id: "1002"
steps:
  - name: create
    config: |
      provider "databricks" {
        account_id = "replay-account"
      }

      resource "databricks_directory" "this" {
        path = "/Shared/replay-hosts"
        provider_config {
          workspace_id = "1002"
        }
      }
    calls:
      # No host: the call must reach the provider host, account.
      - request:
          method: GET
          path: /api/2.0/accounts/replay-account/workspaces/1002
          body: null
        response:
          status: 200
          body:
            json:
              account_id: replay-account
              workspace_id: 1002
              deployment_name: replay-workspace-1002
              workspace_status: RUNNING
      - request:
          method: POST
          host: workspace
          path: /api/2.0/workspace/mkdirs
          body:
            json:
              path: /Shared/replay-hosts
        response:
          status: 200
          body:
            json: {}
      # Reads and destroy follow.
```

- `hosts` maps a name to the host's `url`, an `http` or `https` origin without a path, and optional `metadata`, the response to `GET /.well-known/databricks-config` on that host. Like `host_metadata`, metadata answers every discovery call on its host and need not be consumed. No two hosts may share a host and port.
- `provider_host` names the host the provider is configured with. Without it, the provider uses the local server and named hosts are extra hosts, such as a redirect target. `host_metadata` answers the provider host, so do not also set `metadata` on the host that `provider_host` names.
- A call's `host` names the host the request must reach. A call without `host` expects the provider host. A request to the wrong named host fails as an unexpected call, and the error names the host it reached: `unexpected HTTP call: POST /api/2.0/workspace/mkdirs on host "account"`.
- A request to a host that is not named fails before it reaches the server, with `hermetic replay rejected request outside fixture server: <URL>`. Match that diagnostic with `expect_error` to pin a call to a wrong host, as the [unnamed-host example](testdata/hosts/unnamed_host_rejected.yaml) does.
- A provider block may set `host` to a named host's URL, or to its host without `https://` as the SDK allows, to move the provider there for a step, as the [host-change example](testdata/hosts/provider_host_change.yaml) does. Any other `host` fails provider configuration. Without `hosts`, the runner ignores the HCL `host` as before.
- The SDK takes the cloud from a named host's DNS zone, as it does live: `.azuredatabricks.net` is Azure and `.gcp.databricks.com` is GCP. Host metadata without `cloud` resolves to that cloud, and a private-link login redirect names that cloud's private-link service. On the local server both fall back to AWS. The [Azure example](testdata/hosts/azure_private_link.yaml) shows the redirect.
- Recording rejects tests with `hosts`, because a profile has one host. Update their fixtures by hand.

The [workspace-host example](testdata/hosts/workspace_host.yaml) is the complete journey above.

## Choose the credentials

By default the provider authenticates with a fixed PAT, `fixture-token`. To pin the headers and token requests of another credential type, set `auth` at the test root. Its `type` is the SDK auth type, and its other keys are the provider's credential attributes, named as in a provider block. The [OAuth M2M example](testdata/auth/oauth_m2m.yaml) is a complete journey:

```yaml
version: 1
provider_host: workspace
hosts:
  workspace:
    url: https://replay-workspace.cloud.databricks.com
host_metadata:
  host_type: workspace
  cloud: AWS
  workspace_id: "12345"
auth:
  type: oauth-m2m
  client_id: replay-sp
  client_secret: replay-secret
steps:
  - name: create
    config: |
      resource "databricks_directory" "this" {
        path = "/Shared/replay-auth"
      }
    calls:
      # The token endpoint's discovery, sent without credentials.
      - &discovery
        request:
          method: GET
          path: /oidc/.well-known/oauth-authorization-server
          matchers:
            headers:
              Authorization: absent
          body: null
        response:
          status: 200
          body:
            json:
              authorization_endpoint: https://replay-workspace.cloud.databricks.com/oidc/v1/authorize
              token_endpoint: https://replay-workspace.cloud.databricks.com/oidc/v1/token
      # The client ID and secret, sent as basic auth, exchanged for a token.
      - &token
        request:
          method: POST
          path: /oidc/v1/token
          headers:
            Authorization: Basic cmVwbGF5LXNwOnJlcGxheS1zZWNyZXQ=
          body:
            text: grant_type=client_credentials&scope=all-apis
        response:
          status: 200
          body:
            json: {access_token: m2m-token, token_type: Bearer, expires_in: 3600}
      - request:
          method: POST
          path: /api/2.0/workspace/mkdirs
          headers:
            Authorization: Bearer m2m-token
          body:
            json:
              path: /Shared/replay-auth
        response:
          status: 200
          body:
            json: {}
      # Reads, the post-apply plan's token calls and destroy follow.
```

- `type` is required and is one of `pat`, `basic`, `oauth-m2m`, `azure-client-secret` and `azure-msi`. The SDK then configures only that type, as it does for `auth_type` in a provider block. Without a type, the SDK would try each type in turn, including ones that run a CLI, so the result would depend on the machine.
- The other keys are credential attributes, such as `token`, `username`, `client_secret`, `azure_tenant_id` or `azure_use_msi`; their values are strings, and `azure_use_msi` takes `true` or `false`. Other attributes, such as `host` and `account_id`, stay in the provider block. The runner clears the credential attributes of every provider block, so `auth` is their only source. As before, ambient credentials, environment variables and profiles are ignored.
- Token requests are ordinary calls, so a journey pins them and can fail them. Each Terraform operation configures a new provider, which fetches its own token: a create step has one token exchange for the apply and one for Terraform's post-apply plan, and destroy has one for its refresh and one for the destroy.
- Token endpoints are absolute URLs, and the local server's URL changes on every run, so a journey that fetches tokens names its hosts. Azure AD (`https://login.microsoftonline.com`) and the Azure instance metadata service (`http://169.254.169.254`) are named hosts like any other; the provider host must be an Azure workspace for the SDK to use them.
- To pin a token endpoint that fails, answer its call with an error and set `expect_error` on the step, as the [token failure example](testdata/auth/token_endpoint_failure.yaml) does.
- Recording rejects tests with `auth`, because the profile supplies the credentials. Update their fixtures by hand.

Examples in [testdata/auth](testdata/auth) cover basic auth, OAuth M2M, a rejected client secret, an Azure client secret and Azure MSI.

Fixtures cannot serve the other auth types:

- `databricks-cli` and `azure-cli` run a CLI, which fetches its token outside the provider's transport.
- `google-credentials` signs its token request with a service account's private key, which the journey would have to commit.
- `google-id` reads Google application default credentials from the machine, and calls Google's IAM API with a client that does not use the provider's transport.
- `metadata-service`, `github-oidc`, `github-oidc-azure`, `azure-devops-oidc`, `env-oidc` and `file-oidc` read a token, or the URL to fetch one from, out of a variable or file that a CI runner or the machine provides. The runner does not support them.

## Set environment variables

To pin behavior that comes from the provider's environment, such as the User-Agent dimensions for `DATABRICKS_USER_AGENT_EXTRA`, a CI runner, a coding agent or the Databricks Runtime, set `environment` at the test root. The runner sets each variable with `t.Setenv` for that journey only. The [User-Agent extra example](testdata/environment/user_agent_extra.yaml) pins the products' place in every User-Agent:

```yaml
version: 1
host_metadata:
  workspace_id: "12345"
environment:
  DATABRICKS_USER_AGENT_EXTRA: partner-a/1.2.3 team_b/0.0.1
steps:
  - name: create
    config: |
      resource "databricks_directory" "this" {
        path = "/Shared/replay-environment"
      }
    calls:
      - request:
          method: POST
          path: /api/2.0/workspace/mkdirs
          matchers:
            headers:
              User-Agent:
                pattern: 'databricks-tf-provider/\S+ databricks-sdk-go/\S+ go/\S+ os/\S+ partner-a/1\.2\.3 team_b/0\.0\.1 terraform/\S+ sdk/sdkv2 resource/directory auth/pat'
          body:
            json:
              path: /Shared/replay-environment
        response:
          status: 200
          body:
            json: {}
      # Reads and destroy follow.
```

- Before it sets the journey's variables, the runner clears every variable that changes the User-Agent or the client's configuration, so a journey gets the same result under a coding agent, a CI runner or a developer's shell. It clears every `DATABRICKS_*` variable, including `DATABRICKS_RUNTIME_VERSION` and `DATABRICKS_SDK_UPSTREAM`; every variable the SDK reads its configuration from, such as `ARM_CLIENT_ID` and `GOOGLE_CREDENTIALS`; the coding-agent variables the SDK detects, such as `CLAUDECODE`, `CURSOR_AGENT`, `AGENT` and `AI_AGENT`, and the agent meta-harness variable `OMNIGENT`; the CI variables the SDK detects, such as `GITHUB_ACTIONS`, `GITLAB_CI`, `JENKINS_URL` and `TF_BUILD`; and the rollout variables `USE_SDK_V2_RESOURCES` and `USE_SDK_V2_DATA_SOURCES`. `userAgentVariables` in [environment.go](environment.go) lists every agent and CI variable. If the SDK still detects an agent, CI runner or runtime once they are cleared, as it would from a variable a newer SDK reads, the journey fails and names what the SDK detected.
- `environment: {}` clears these variables and sets none, so a journey can pin a whole User-Agent. Without `environment`, the journey keeps the process's variables.
- The provider reads `DATABRICKS_USER_AGENT_EXTRA` once, when its process starts, so `t.Setenv` cannot change it. A journey that sets it runs in a new test process that starts with the journey's environment, as Terraform starts the provider in a new process, and a failure there fails the journey with that process's output. `RunNative` rejects it, because its Go checks run in the calling process. If the test process itself starts with `DATABRICKS_USER_AGENT_EXTRA`, every journey with `environment` fails, because every User-Agent already contains it.
- The SDK reads the agent, CI and runtime variables for each request, so those journeys run in the test process. The [detected User-Agent example](testdata/environment/detected_user_agent.yaml) sets `GITHUB_ACTIONS`, `DATABRICKS_RUNTIME_VERSION` and `CLAUDECODE`.
- A journey cannot set the variables the runner sets, `TF_ACC_*`, `CHECKPOINT_DISABLE`, `TF_IN_AUTOMATION` and `TEST_RECORDER_PROFILE_NAME`, or the SDK's configuration variables, such as `DATABRICKS_HOST` and `DATABRICKS_TOKEN`: replay ignores those, so use `auth` and the provider block instead.
- Recording rejects tests with `environment`. Update their fixtures by hand.

## Write tests that survive redaction

The recorder redacts known secrets in both HTTP requests and responses, including
secret-named fields containing fake test values. `[REDACTED]` is a literal string,
not a wildcard. A live run can pass while replay fails because its outgoing
request still contains the HCL value, or because a redacted response no longer
matches an assertion or a configured value. This is an accepted tradeoff: keeping
secrets out of fixtures takes priority over replaying every capture unchanged.
The test author, including an agent updating a test, must resolve these mismatches.

Choose the approach that preserves the behavior under test:

- For an incidental secret string in a request, keep the redacted sample and add
  `any_string` only at that field. It checks presence and string type while
  allowing the nonsecret HCL input to differ. For example, a secret-write request
  can use:

  ```yaml
  request:
    method: POST
    path: /api/2.0/secrets/put
    body:
      json:
        scope: fixture-scope
        key: fixture-key
        string_value: "[REDACTED]"
    matchers:
      json:
        - path: [string_value]
          type: any_string
  ```

  Keep `scope`, `key`, and other behavior-specific fields exact. Header and query
  string matchers are also supported; JSON matchers require `body.json`, not
  text, base64, or multipart bodies. Matchers apply only to outgoing requests;
  they do not change responses, HCL, or plan/state checks.
- When exact equality matters, use a nonsecret HCL input matching the fixture's
  redacted value, such as `string_value = "[REDACTED]"`, and align relevant
  request/response samples and plan/state `equals` checks. This also avoids a
  configured-value consistency failure if the provider reads the redacted value
  back into state. For secret-read responses, preserve the API's encoding: the
  `/secrets/get` response contains base64-encoded `[REDACTED]`, and the decoded
  value is `[REDACTED]`. Live checks execute too; an expectation of `[REDACTED]`
  must also hold during recording. If the scenario cannot satisfy both modes,
  choose a secret-free scenario or test that behavior separately.

Before recording, add any needed matchers to the existing YAML and use fake
inputs in HCL and local files. The recorder preserves matching annotations but
does not invent them or rewrite HCL and assertions. If verification replay fails,
the source stays unchanged: adjust those inputs, matchers, or expectations and
record again. Never restore the original secret or disable redaction to fix a
failure, and do not mask fields that define the regression being tested.

After recording, review the whole diff for secrets, including text/base64
payloads, local files, and checkpoints; automatic redaction cannot identify every
service-specific secret. Then unset `TEST_RECORDER_PROFILE_NAME` and run the same
exact test path with `-count=1` to verify ordinary replay succeeds. Submit only
secret-free fixtures with passing replay.

## Test local files and uploads

Keep tiny inputs next to the YAML, such as `assets/script.py` containing only
`# comment`. Use the step's [`files` map](README.md#local-input-files) to stage
it, and pass `abspath("${path.module}/script.py")` to a provider's `source` field.
Change its mapping in a later step to test an update without changing the HCL.
Omitted mappings keep their previous contents through destroy.

Use `body.text` for exact text, `body.base64` for binary bytes, and
[`body.multipart`](../httpfixture/README.md#files-and-non-json-bodies) for uploads.
Keep content exact; multipart boundaries and part order are transport details
and are ignored. The [workspace-file showcase](testdata/files/workspace_file.yaml)
exercises staging, upload, replacement, and cleanup with comment-only inputs.

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

## Test slow and broken connections

Optional `response` fields script what happens on the wire, so a journey can pin
timeouts, retries and malformed replies without a hand-written `httptest`
server. The fixture server applies them, so they behave the same whichever SDK
client the provider uses.

- `delay`: a Go duration, such as `2s`, to wait before writing anything. The
  wait ends early when the client closes the connection, so a timed-out request
  does not hold the server.
- `fault`: breaks the connection instead of completing the response.
  - `reset` closes it with linger 0, so the client sees `connection reset by peer`.
  - `close` closes it without writing anything, so the client sees `EOF`.
  - `truncate_body` writes the status line, the headers with the full
    `Content-Length` and the first half of the body, then closes it.
  - `stall_body` writes the same first half, then waits until the client closes
    the connection.
- `status_line`: the literal HTTP/1.0 or HTTP/1.1 status line, such as
  `HTTP/1.1 404` with no reason phrase. Its code must equal `status`; headers
  and body are written as usual, and a `HEAD` response keeps its fixture
  `Content-Length`.

`reset` and `close` write no response, so they omit `status`, `headers`, `body`
and `status_line`. The body faults need a nonempty `body` on a response that can
carry one: not a `HEAD` request or a 1xx, 204 or 304 status, whose bodies
`net/http` never reads. `delay` combines with every other field. Invalid values
fail when the journey loads, naming the file, step and call.

A faulted call is consumed like any other, so `enforce_call_order` and the
unused-fixture check still apply. Follow it with a normal call to pin a retry,
or with nothing to pin that there is none, because a retry then fails as an
unexpected call. This create is retried after a reset:

```yaml
calls:
  - request: &create
      method: POST
      path: /api/2.0/preview/scim/v2/Groups
      body:
        json:
          displayName: Example group
          schemas: [urn:ietf:params:scim:schemas:core:2.0:Group]
          entitlements: [{}]
    response:
      fault: reset
  - request: *create
    response:
      status: 201
      body: {json: {id: group-id}}
  # The read after create and the post-apply refresh follow.
```

Pair `delay` or `stall_body` with a short `http_timeout_seconds` in the provider
block and a delay longer than it. Each example in [testdata/faults](testdata/faults)
fails or retries a group create, with this classic SDK diagnostic:

| Response | Example | Diagnostic |
|---|---|---|
| `delay: 10s`, `http_timeout_seconds = 1` | [delay](testdata/faults/delay.yaml) | `request timed out after 1s of inactivity`, not retried |
| `fault: reset` | [reset](testdata/faults/reset.yaml) | `connection reset by peer`, retried |
| `fault: close` | [close](testdata/faults/close.yaml) | `EOF`, not retried |
| `fault: truncate_body` | [truncate_body](testdata/faults/truncate_body.yaml) | `response body: unexpected EOF`, not retried |
| `fault: stall_body`, `http_timeout_seconds = 1` | [stall_body](testdata/faults/stall_body.yaml) | `response body: context canceled`, not retried |
| `status: 404`, `status_line: HTTP/1.1 404` | [status_line](testdata/faults/status_line.yaml) | `Not Found` |

The replay transport, which `qa/replay/clienttest` shares, opens one connection
per request. Go's HTTP client silently resends a replayable request, such as a
GET, that fails on a reused connection, so a fresh connection is what lets each
fault reach the provider.
Faults and `status_line` close the connection and frame the body by its
`Content-Length`, replacing those headers in the fixture.

Connection refused is not supported: the server has already accepted the
connection by the time it matches a request. Recording rejects journeys that
use these fields, because a live service cannot reproduce them and a refresh
would replace them; edit such journeys by hand.

## Test repeated calls and retry timing

Optional call fields, next to `request` and `response`, say how often a call repeats and how long the client waits between calls, so a retry sequence needs no copies of the call and its backoff can be pinned.

- `times: N` expects the call exactly N times in a row, each with the same response. It is shorthand for N copies of the call.
- `times: {min: 2, max: 3}` expects between `min` and `max` repetitions, for a count that depends on timing, such as retries that end with `retry_timeout_seconds`. The unused-fixture check needs `min`; a repetition past `max` is an unexpected call. A `min` of 0 makes the call optional.
- `min_gap` and `max_gap` are Go durations, such as `1s` or `1500ms`, bounding the time between arrivals at the fixture server, taken before each request's body is read. Without `times`, they bound the gap since the previous matched call in the step. With `times`, they bound the gap before each repetition since the previous one; the first repetition is not bounded, because it usually is the attempt that the retries follow. List it as its own call to bound it too.

A gap out of bounds fails the step with the observed gap, such as `call 2: POST /api/2.0/preview/scim/v2/Groups arrived 2.1s after call 1 (POST /api/2.0/preview/scim/v2/Groups), above max_gap 1s`. The call is still answered, so the failure names the gap rather than whatever the provider does next. A bounded call without `times` that arrives first in its step fails too, since nothing precedes it. Host metadata is not a matched call and does not count. Invalid values fail when the journey loads, naming the file, step and call.

This create gets 503 twice and succeeds on its second retry:

```yaml
calls:
  # The first attempt and the first retry.
  - request: &create
      method: POST
      path: /api/2.0/preview/scim/v2/Groups
      body:
        json:
          displayName: Example group
          schemas: [urn:ietf:params:scim:schemas:core:2.0:Group]
          entitlements: [{}]
    response:
      status: 503
      body: {json: {error_code: TEMPORARILY_UNAVAILABLE, message: Service unavailable from fixture}}
    times: 2
    min_gap: 1s
    max_gap: 3s
  # The second retry.
  - request: *create
    response:
      status: 201
      body: {json: {id: group-id}}
    min_gap: 2s
    max_gap: 4s
  # The read after create and the post-apply refresh follow.
```

This one gets 429 until `retry_timeout_seconds` ends, then fails with the last response:

```yaml
config: |
  provider "databricks" {
    retry_timeout_seconds = 5
  }
  resource "databricks_group" "this" {
    display_name = "Example group"
  }
expect_error: 'cannot create group: Too many requests from fixture'
calls:
  - request:
      method: POST
      path: /api/2.0/preview/scim/v2/Groups
      body:
        json:
          displayName: Example group
          schemas: [urn:ietf:params:scim:schemas:core:2.0:Group]
          entitlements: [{}]
    response:
      status: 429
      body: {json: {error_code: TOO_MANY_REQUESTS, message: Too many requests from fixture}}
    times: {min: 2, max: 3}
    min_gap: 1s
    max_gap: 4s
```

The complete journeys are [retry_503](testdata/repeats/retry_503.yaml) and [retry_429_until_timeout](testdata/repeats/retry_429_until_timeout.yaml). The runner keeps `retry_timeout_seconds`, `http_timeout_seconds` and `rate_limit` from the journey's provider block. Set a short `retry_timeout_seconds` for a journey that retries until the timeout: the provider's default retries indefinitely.

Choose bounds from the client's backoff, with room for a slow CI machine. The classic SDK waits `min(n, 10)` seconds plus 50-750 ms of jitter before retry n: 1.05-1.75 s before the first, 2.05-2.75 s before the second, and 10.05-10.75 s from the tenth on. The observed gap adds the time to answer the previous call and reconnect, so it is never shorter than the wait.

- Set `min_gap` to the shortest wait, rounded down, such as `1s` before the first retry. A slow machine only makes the gap longer.
- Set `max_gap` to the longest wait plus at least 1 s, such as `3s` before the first retry, because a slow machine can stall between calls.
- For a range, attempt n starts after the sum of the earlier waits: attempt 2 at 1.05-1.75 s, attempt 3 at 3.1-4.5 s and attempt 4 at 6.15-8.25 s. Set `max` to the attempts that can start before the timeout, and `min` to those that start well before it on a slow machine.

Repetitions match like copies of the call. Without `enforce_call_order`, other calls can arrive between them. With it, a call must reach its `min` before the next listed call, and cannot repeat once a later call arrives. A range takes every matching request until it reaches `max`, so an identical call listed after a range answers only once the range is full; end a retry sequence with the range, as in the 429 example. Recording rejects journeys with these fields, because a refresh would replace a repeated call with captured copies; edit such journeys by hand.

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

## Native Go checks and provider selection

Use YAML for ordinary lifecycles. For a case that needs per-step provider
factories or a Go check, `RunNative(t, path, configure)` loads one fixture file
and lets Go customize its HashiCorp `resource.TestCase` before execution. It
supplies a config customizer that retains the runner's local HTTP boundary and
isolated authentication. Apply it to every custom provider and SDK client;
keep the generated step and cleanup hooks so fixtures are verified and destroy
still requires empty state.

See `qa/replay/features/provider_migration/replay_test.go` for factory switches
and `qa/replay/resources/mws_workspaces/native_test.go` for SDK checks and an
explicit local workspace-host alias. Native fixture files are registered by
their Go wrappers; do not also discover them through the ordinary YAML wrapper.
