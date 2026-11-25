# sdk-go-utility

Package `sdkgoutility` lets an SDKv2 resource call the modular Databricks SDK
(`github.com/databricks/sdk-go`) and behave as it does on the classic one
(`github.com/databricks/databricks-sdk-go`). It has 16 helpers, each in its own file next to
its tests. `databricks_mlflow_model` is the first resource that uses them, and this document
uses it as the example of every step.

## How it works

A resource that moves keeps two implementations of create, read, update and delete:

- **The classic path** is the code the resource had. It is the default.
- **The modular path** makes the same calls with a client of the modular SDK, here
  `github.com/databricks/sdk-go/modelregistry/v1`.

**The opt-in** is the environment variable `USE_SDK_GO_RESOURCES`, read by `Enabled` in
[`enabled.go`](enabled.go). It lists Terraform resource types, separated by commas or white
space: `USE_SDK_GO_RESOURCES="databricks_mlflow_model"`. A name that no resource asks about
enables nothing and is not reported.

- Every operation asks `Enabled` when it runs (`mlflowModelPickPath` in
  [`mlflow/resource_mlflow_model_sdkgo.go`](../mlflow/resource_mlflow_model_sdkgo.go)), so
  the variable can change between two operations of one provider process.
- The schema is built once, when the provider builds the resource (`mlflowModelWithSdkGo`).
  The two schemas must be equal, so that an operation of either path works on either.

One operation on the modular path, `mlflowModelSdkGoCall`:

1. `c.WorkspaceClientUnifiedProvider(ctx, d)` returns the classic workspace client, as on the
   classic path, so `provider_config` routes the call the same way.
2. `NewClient(ctx, w.Config, modelregistry.NewClient)` builds a modular client from the config
   of that client, which sends with the headers, the debug log and the attempt timeout of the
   classic client. It costs no request, so the resource builds one for every operation.
3. `ClassicCall(ctx, w.Config, method)` runs one method of the client with the retries, the
   rate limit and the errors of the classic client.

**To run a resource on the modular path:**

| What | How |
|---|---|
| Unit tests and journeys, both paths | `env -u DATABRICKS_WORKSPACE_ID go test -short ./mlflow/... ./qa/replay/resources/mlflow_model...` with Terraform 1.13.5 on `PATH`. The tests set and clear the variable themselves. |
| Only the modular path | add `-run 'TestReplaySdkGo'` for the journeys, or `-run '/sdkgo'` for the unit tests |
| A real workspace | set `USE_SDK_GO_RESOURCES=databricks_mlflow_model` in the environment that Terraform runs in. `Enabled` reads it in the provider process. Not tried yet. |
| The acceptance test | `USE_SDK_GO_RESOURCES=databricks_mlflow_model go test -v -run '^TestAccMLflowModel$' ./mlflow/`, with the workspace environment loaded as [`AGENTS.md`](../AGENTS.md) describes |

The acceptance test `TestAccMLflowModel` in
[`mlflow/mlflow_model_test.go`](../mlflow/mlflow_model_test.go) has not run on the modular
path. Everything in this document was checked against replay journeys and local test servers,
never against a workspace. Running that command is the check that is left.

## Moving a resource

Do the steps in this order. Each names the file of `databricks_mlflow_model` that shows it.

### 1. Read "Check before you move a resource"

The helpers cover what `databricks_mlflow_model` needs. The section below lists what they
don't, and how to tell whether the resource is affected.

### 2. Pin the classic behavior, before any production change

Write tests that pass on the classic path and assert bytes on the wire, states and plans:
replay journeys ([authoring guide](../qa/terraformtest/AUTHORING.md)) with literal JSON and
`enforce_call_order: true`, and `qa.ResourceFixture` unit tests with literal JSON where a
journey can't express the case. A test built on `MockWorkspaceClientFunc` fails on the modular
path, and a fixture built from classic SDK structs doesn't pin the bytes. Rewrite both first,
as [`mlflow/resource_mlflow_model_test.go`](../mlflow/resource_mlflow_model_test.go) was.

| Theme | What to pin | Why it mattered here | Where |
|---|---|---|---|
| Create, replace and delete | the exact body, query and headers of each request for each shape of the configuration, and where the ID comes from | `DataToStruct` leaves an empty `name` out, which the classic request sends | [`mlflow_model_lifecycle`](../qa/replay/resources/mlflow_model_lifecycle), [`resource_mlflow_model_lifecycle_test.go`](../mlflow/resource_mlflow_model_lifecycle_test.go) |
| Updates and drift | the body of each update, and the state after each refresh | the state depends on the skipping rules of `common.StructToData`, which `StructToData` has to repeat | [`mlflow_model_updates`](../qa/replay/resources/mlflow_model_updates), [`resource_mlflow_model_update_test.go`](../mlflow/resource_mlflow_model_update_test.go) |
| API errors | the diagnostic, the state and the next plan after a failure of each call | a modular error has another text and type until `ClassicError` converts it | [`mlflow_model_errors`](../qa/replay/resources/mlflow_model_errors), [`resource_mlflow_model_errors_test.go`](../mlflow/resource_mlflow_model_errors_test.go) |
| Error bodies | which bodies drop the resource on refresh and count as deleted on delete | the classic parser decides by the body, not by the status, so a conversion by status would be wrong | [`mlflow_model_error_bodies`](../qa/replay/resources/mlflow_model_error_bodies) |
| Response decoding | success responses with missing envelopes, nulls, wrong types, empty and undecodable bodies | the journeys that differ on the modular path are all here: its decoding is stricter | [`mlflow_model_decoding`](../qa/replay/resources/mlflow_model_decoding) |
| Routing by `provider_config` | the host and the `X-Databricks-Workspace-Id` header of every call | a modular client left alone takes host and credentials from the environment | [`mlflow_model_routing`](../qa/replay/resources/mlflow_model_routing) |
| Import, validation and schema | the import, the configurations that fail before a request, and every attribute of the schema | the modular schema is built from another struct | [`mlflow_model_schema`](../qa/replay/resources/mlflow_model_schema), [`resource_mlflow_model_schema_test.go`](../mlflow/resource_mlflow_model_schema_test.go) |
| Retries | which responses each call retries, which it sends once, and what apply reports when the retries run out | a modular client retries by other rules, and not at all without a retrier | [`mlflow_model_retries`](../qa/replay/resources/mlflow_model_retries) |
| `User-Agent` and `traceparent` | the exact `User-Agent` in every detected environment, and `traceparent` on every attempt | usage is attributed through the `User-Agent`, and a modular client sends its own | [`mlflow_model/user_agent_test.go`](../qa/replay/resources/mlflow_model/user_agent_test.go) |

Pin behavior that looks wrong as it is. One package per theme keeps each under the 10 minutes
that `go test` gives a package.

### 3. Add the modular path

| Step | What to write | Where |
|---|---|---|
| The switch file | a file `resource_<name>_sdkgo.go` with the resource type, a function that wraps the classic resource, and one that picks the path for an operation. The classic file changes only where it passes the resource to the switch. | `mlflowModelResourceType`, `mlflowModelWithSdkGo` and `mlflowModelPickPath` in [`resource_mlflow_model_sdkgo.go`](../mlflow/resource_mlflow_model_sdkgo.go); [`resource_mlflow_model.go`](../mlflow/resource_mlflow_model.go) |
| The schema | `StructToSchema` of the modular request, with a customization that adds what the classic struct tags gave: required, `ForceNew`, computed attributes, and `provider_config` through `common.AddNamespaceInSchema` and `common.NamespaceCustomizeSchemaMap` | `mlflowModelSdkGoSchema` |
| The schema test | compare the modular schema with the classic one in every property at every depth | `TestResourceMlflowModelSdkGoSchemaEqualsClassic` in [`resource_mlflow_model_sdkgo_test.go`](../mlflow/resource_mlflow_model_sdkgo_test.go) |
| The requests | `DataToStruct` for a request that the schema describes. Then set every pointer that must be sent although it is zero. Build a request that has only the ID by hand. | `create`, `update` and `mlflowModelSdkGoName`; `read` and `delete` |
| The state | check the response for a nil message, then `StructToData`, then the attributes that have no field of their name | `read` and `mlflowModelSdkGoString` |
| The client | the classic client from the provider, then `NewClient` with its config | `mlflowModelSdkGoCall` |
| The calls | `ClassicCall` around every method of the modular client, one call each | `mlflowModelSdkGoCall` |

### 4. Run every test on both paths

| Step | What to write | Where |
|---|---|---|
| Unit tests | move each test body into a function and run it as the subtests `classic` and `sdkgo`. The `classic` subtest clears the variable, so that the environment of the test process can't switch the path. | `mlflowModelOnBothPaths` in [`resource_mlflow_model_sdkgo_test.go`](../mlflow/resource_mlflow_model_sdkgo_test.go) |
| Journeys | add `TestReplaySdkGo` next to `TestReplay` in every replay package | [`mlflow_model_lifecycle/replay_test.go`](../qa/replay/resources/mlflow_model_lifecycle/replay_test.go) |
| Native tests | run them as `classic` and `sdkgo` as well, and expect the `sdk-go-` segments in the `User-Agent` of the modular path | `wantUserAgent` in [`user_agent_test.go`](../qa/replay/resources/mlflow_model/user_agent_test.go) |
| The twin of a journey that differs | move the journey unchanged to `testdata_classic/`, which only `TestReplay` runs, and write its twin under the same name in `testdata_sdkgo/`, which only `TestReplaySdkGo` runs. The header of the twin says what differs and why it is accepted. | [`mlflow_model_decoding/replay_test.go`](../qa/replay/resources/mlflow_model_decoding/replay_test.go), [`testdata_sdkgo/decode_wrong_types.yaml`](../qa/replay/resources/mlflow_model_decoding/testdata_sdkgo/decode_wrong_types.yaml) |
| The changelog | one entry under `Internal Changes` while the opt-in is not offered to users | [`NEXT_CHANGELOG.md`](../NEXT_CHANGELOG.md) |

A journey that differs in an outcome needs a decision of the owner of the resource, not only a
twin. Two things to know about the twins:

- CI spreads the replay packages over its shards by the number of YAML files directly in
  `testdata/` (`replay_weight` in
  [`.github/scripts/select_test_packages.py`](../.github/scripts/select_test_packages.py)).
  It doesn't count the twins, and every journey now runs twice.
- A pattern on a message of `encoding/json` has to take the spelling of Go 1.26 and of Go 1.27,
  as the twins do.

### 5. Run the acceptance test on the modular path

See "How it works". For `databricks_mlflow_model` this is still open.

## The helpers

`cfg` is the `*config.Config` of the classic SDK, always the config of a classic client that
exists: `w.Config` of the client that `WorkspaceClientUnifiedProvider` returns. The doc comment
of each helper has the details.

| Helper | Signature | What it does | What it leaves to its caller |
|---|---|---|---|
| [`FieldName`](field_name.go) | `FieldName(structType reflect.Type, field reflect.StructField) string` | The API name of a field of a modular SDK struct, read from the JSON that the struct marshals to, then from the `fieldmask` tag, then by snake_case. | Recognizing a field that is not in the JSON at all: a oneof group, a body stream, a field sent in a header. |
| [`StructToSchema`](struct_to_schema.go) | `StructToSchema(v any, customize func(map[string]*schema.Schema) map[string]*schema.Schema) map[string]*schema.Schema` | An SDKv2 schema from a modular SDK struct, every attribute `Optional`. Panics on a field without a schema type. | Required, computed, `ForceNew`, sensitive, defaults, sets and `provider_config`, all in `customize`. Checking that SDKv2 accepts the names. |
| [`DataToStruct`](data_to_struct.go) | `DataToStruct(d *schema.ResourceData, s map[string]*schema.Schema, result any) error` | Fills a modular SDK struct from resource data. A value that is unset or zero stays nil. A timestamp and a duration are read from a string, and a field mask is skipped. | Setting a pointer that must be sent although it is zero, and a field mask. Sets. |
| [`StructToData`](struct_to_data.go) | `StructToData(result any, s map[string]*schema.Schema, d *schema.ResourceData) error` | Writes a modular SDK struct into resource data, by the rules of `common.StructToData`. A timestamp and a duration are written as the string that the SDK marshals, and a field mask is skipped. | A non-nil result. Attributes that have no field of their name. |
| [`Credentials`](credentials.go) | `Credentials(cfg *config.Config) auth.Credentials` | Credentials that call `cfg.Authenticate`, so both clients share one token. | A resolved config. Everything but the authentication headers. |
| [`Transport`](transport.go) | `Transport(cfg *config.Config, wrap func(base http.RoundTripper) (http.RoundTripper, error)) (http.RoundTripper, error)` | A transport that sends with the classic `User-Agent`, `traceparent`, `Accept`, `Headers` hook, proxy and TLS settings. It answers the login page of a private link redirect with the 403 that the classic client reports. `wrap` puts a transport of the caller between these steps and the connection. | Credentials that authenticate `cfg` first. Retries, rate limit, every other error. The HTTP timeout and request logging, for which `wrap` is the place. |
| [`ClientOptions`](client_options.go) | `ClientOptions(cfg *config.Config, credentials auth.Credentials, transport http.RoundTripper) ([]client.Option, error)` | The options that point a modular client at the host and workspace of `cfg`, without environment or config file, bounded by `retry_timeout_seconds`. | A config that has authenticated. Credentials and transport. A retrier and a limiter for each call. A logger. |
| [`AttemptTimeout`](attempt_timeout.go) | `AttemptTimeout(cfg *config.Config, transport http.RoundTripper) (http.RoundTripper, error)` | Ends an attempt after `http_timeout_seconds` without progress, with the error of the classic client. | Closing the response body. Whether a timed-out attempt is retried. The limit on the whole call. |
| [`DebugLog`](debug_log.go) | `DebugLog(cfg *config.Config, transport http.RoundTripper) (http.RoundTripper, error)` | Writes each attempt to the debug log as the classic client does: the method and the URL, the headers with `debug_headers`, the bodies and the status or the error, redacted and cut to `debug_truncate_bytes`. It only passes the request on when debug logging is off. | Its place among the transports: inside the one that sets the headers, outside the one that ends an attempt. Retries. |
| [`ClassicError`](classic_error.go) | `ClassicError(ctx context.Context, err error) error` | Turns the error of a modular call into the classic `*apierr.APIError` for the same response, by parsing the response again. | Retries. The overrides by request path, which it can't see. The private link redirect, which `Transport` answers. |
| [`Retrier`](retrier.go) | `Retrier(ctx context.Context, cfg *config.Config, classic func(error) error) (ops.Retrier, error)` | A retrier for one call that retries what the classic client retries, with its backoff and its retry timeout. | Converting errors. A new retrier for every call. The timeout of the call. Rate limiting. |
| [`Limiter`](limiter.go) | `Limiter(cfg *config.Config) ops.Limiter` | The rate limiter of the classic client, with one budget for each `*config.Config`. | A resolved config, the same pointer every time. Passing it to every call. |
| [`Call`](call.go) | `Call[T any](ctx context.Context, method func(context.Context, ...call.Option) (T, error), retrier ops.Retrier, limiter ops.Limiter, classic func(error) error) (T, error)` | Runs one method with a retrier, a limiter and an error converter, and returns the last error when the context ends during a wait to retry. | The three arguments. The timeout. A method that returns an iterator. |
| [`Enabled`](enabled.go) | `Enabled(resource string) bool` | Whether `USE_SDK_GO_RESOURCES` names the resource. Reads the variable on every call. | The resource type as Terraform spells it. Branching, and asking again for every operation. |
| [`NewClient`](new_client.go) | `NewClient[C any](ctx context.Context, cfg *config.Config, construct func(context.Context, ...client.Option) (C, error)) (C, error)` | Builds a modular client from `Credentials`, `Transport`, `DebugLog`, `AttemptTimeout` and `ClientOptions`. A request goes through `Transport`, then `DebugLog`, then `AttemptTimeout`, then the connection. | A config that has authenticated. A retrier and a limiter for each call. The error of a failed call. |
| [`ClassicCall`](classic_call.go) | `ClassicCall[T any](ctx context.Context, cfg *config.Config, method func(context.Context, ...call.Option) (T, error)) (T, error)` | Runs one method through `Call` with a new `Retrier`, `Limiter(cfg)` and `ClassicError`. | The client. The same config pointer for every call. A method that makes exactly one call. |

A resource calls `Enabled`, `StructToSchema`, `DataToStruct`, `StructToData`, `NewClient` and
`ClassicCall`. The other ten are their parts.

## Differences from the classic client

What a resource on the modular path does differently, with all helpers in place. "Outcome"
means a state, a plan or whether a step fails. The tests are in this package unless a path
says otherwise. `sdkgo twin` is the journey of that name in
[`mlflow_model_decoding/testdata_sdkgo`](../qa/replay/resources/mlflow_model_decoding/testdata_sdkgo).

| # | Difference | Changes | Pinned by |
|---|---|---|---|
| 1 | `User-Agent` ends with the `sdk-go-` segments of the modular SDK, such as ` sdk-go-core/0.1.0 sdk-go-modelregistry/0.7.0` | wire | `TestMlflowModelUserAgentDimensions` and `TestMlflowModelUserAgentExtra` in [`user_agent_test.go`](../qa/replay/resources/mlflow_model/user_agent_test.go), `TestTransport_UserAgent` |
| 2 | The keys of a JSON request body come in another order | wire | `TestNewClient_DifferencesFromClassicClient` |
| 3 | The transport runs again on a redirected request, so it gets a new `traceparent` and the `Headers` hook runs on it once more | wire | no test asserts either. `TestTransport_WrapSeesThePrivateLinkLoginPage` shows that both requests of a redirect get the classic headers. |
| 4 | A `Headers` hook wins over the headers of the call, such as `Content-Type`. The provider sets no hook. | wire | `TestTransport_Headers` for `Authorization`; no test for a header of the call |
| 5 | The credentials strategy is given `GET /` instead of the request, and only the headers it sets are used | wire | `TestCredentials_RequestGivenToStrategy` |
| 6 | Modular clients have connection pools of their own: one for all configs that verify certificates, one for all with `InsecureSkipVerify` | wire | `TestTransport_Config`, `TestTransport_InsecureSkipVerify`, `TestNewClient_KeepsNoState` |
| 7 | The classic and the modular calls of one config have separate rate limit budgets, so together they can send twice the rate. The modular budget is one for each config, where the classic client has one for each client. | wire | `TestLimiterHoldsModularCallsToClassicRate`, `TestLimiterBudgets`, `TestClassicCallSharesTheRateLimitOfAConfig` |
| 8 | With `debug_headers`, the debug log shows the `Host` that is sent, which the classic client logs as empty, and the `Traceparent` and `User-Agent` lines differ as the headers do. Without `debug_headers` the message of a call is byte for byte the classic one. | text | `TestNewClient_WritesDebugLogOfClassicClient` |
| 9 | The classic error parser logs twice about the last response when its body is not JSON | text | `TestClassicCallReturnsWhatTheClassicClientReturns` |
| 10 | An error comes without the `*retries.Err` wrapper, and the result next to it is the zero value instead of an empty response. Its `ResponseWrapper` holds a rebuilt response without the request, and for the private link error the 403 that `Transport` made in place of the login page. | only code that reads these sees it | `TestClassicCallReturnsWhatTheClassicClientReturns`, `TestClassicError_RebuildsResponse`, `TestTransport_PrivateLinkResponse` |
| 11 | A success response that can't be decoded reports the error of `encoding/json` alone, without the request to report a bug and the request log | text | sdkgo twins `decode_timestamps`, `decode_undecodable_bodies` |
| 12 | A response without its message is a clear error (`the response has no registered_model`), where the classic path panics or reports a reflection error. This is in the resource, not in a helper. | text | sdkgo twins `decode_create_without_registered_model`, `decode_read_without_registered_model_databricks` |
| 13 | A value that is no string, sent for a string field, fails the call. The classic client takes a number as its digits. | **outcome** | sdkgo twin `decode_wrong_types` |
| 14 | An empty 200 body where the response has fields fails the call. The classic client takes it as an empty response. | **outcome** | sdkgo twin `decode_unexpected_patch_and_delete_bodies` |
| 15 | A timestamp and a duration are written to the state and sent in the form that the modular SDK marshals: a time in UTC with a `Z`, and both with 0, 3, 6 or 9 fraction digits. The classic struct holds the text as it is, so the state differs where the API sends another form, such as `2026-01-02T03:04:05.000Z`, an offset or `3.5s`. `databricks_mlflow_model` has no such field. | **outcome** for such a field | `TestStructToData_State` and `TestDataToStruct_ReadsTimeAndDuration` pin the form. No test compares it with another form of the API: that was seen in a one-off run when this document was written. |
| 16 | An HTML error with a `<pre>` block gets its error code from the standard text of the status, not from the text that the server sent. For `404` without a text over HTTP/1.x the resource counts as missing, and on the classic path it doesn't. | **outcome** | `TestClassicError_DiffersByStatusText` |
| 17 | With `retry_timeout_seconds` set, the bound also ends an attempt in flight and the wait for the rate limiter. The classic client checks it only between attempts. | **outcome** | `TestClientOptions_BoundsCallLikeClassicClient`, `TestCallWhenTheTimeoutOfTheMethodPasses`, `TestLimiterWaitFails` |
| 18 | Retries that run out return up to one backoff sooner, with the same error | timing | `TestRetrierStopsWhenClassicClientDoes`, `TestClassicCallReturnsWhatTheClassicClientReturns` |
| 19 | The retry rules also see an error from before the request is sent, such as a failed authentication, which the classic client never retries | **outcome**: more attempts | no test |
| 20 | A stall while the response body is read ends with `request timed out after ... of inactivity`, where the classic client reports `response body: context canceled`. Other errors of reading the body come without the `response body: ` prefix. | text | `TestAttemptTimeout_EndsLikeClassicClient`, `TestAttemptTimeout_WaitsLikeClassicClient` |
| 21 | A response that ends just after the timeout is a timeout. The classic client takes the cut body, and succeeds when the body is empty. | **outcome** | `TestAttemptTimeout_WaitsLikeClassicClient` |
| 22 | Reads from a request body that `net/http` rewinds restart the wait. On the classic client a slow upload after a rewind times out. | **outcome** | `TestAttemptTimeout_WaitsLikeClassicClient` |
| 23 | Under a `cfg.HTTPTransport` that wraps `context.Canceled`, a timed-out attempt reports the timeout, where the classic client reports the wrapped cancellation | text | `TestAttemptTimeout_WaitsLikeClassicClient` |
| 24 | A 400 with `INVALID_PARAMETER_VALUE` from a read of a cluster, a job or a job run is not a missing resource. No call of `databricks_mlflow_model` is affected: see the next section. | **outcome** | `TestClassicError_DiffersByRequestOverride` |
| 25 | With a fixture transport of the classic SDK (`SkipRetryOnIO`) and `retry_timeout_seconds` 0, the classic client retries for 20 minutes and the modular call ends after 5. The provider sets -1, so this shows in tests only. | **outcome** in such a test | no test |
| 26 | The debug log has a message for each request of a redirect, where the classic client writes one with the first request and the last response. For the private link login page that is two messages, and the one for the page has no text. | text | `TestNewClient_DebugLogShowsWhatIsSent` |
| 27 | The debug log shows a request body only when it is JSON, and `<io.Reader>` otherwise, where the classic client also logs a string or a byte slice. A response to a request without `Accept` is taken for JSON, so a stream that sets none is copied while the caller reads it and logged at its end. No call of `databricks_mlflow_model` is affected. | text | `TestDebugLog_Message` |
| 28 | The inactivity timeout of an attempt starts after the request visitors have run. On the classic client it starts before them, so a `Headers` hook that takes longer than `http_timeout_seconds` fails the attempt there. | **outcome** with such a hook | `TestNewClient_TimesAttemptFromTheConnection` |

Not in the list, because the helpers remove them together:

- A context that is cancelled during a wait to retry returns the error of the last attempt, as
  on the classic client. `Retrier` alone can't do that, and `Call` does
  (`TestCallWhenTheContextEnds`).
- A stalled attempt ends with the same text on both paths, in the error and in the debug log,
  and is not retried. `TestNewClient_EndsStalledAttemptLikeClassicClient` pins that for a
  client without a retrier. No test runs it through `ClassicCall`.
- After a failed attempt that is not retried, the retry loop of the classic client writes
  `non-retriable error: ...` to the debug log. A client of `NewClient` alone doesn't
  (`TestDebugLog_MatchesClassicClient`), and the `Retrier` of `ClassicCall` does. No test
  compares the log of a call through `ClassicCall` with the classic one: that was seen in a
  one-off run when this document was written.
- A workspace that refuses a request from outside its private link redirects it to its login
  page, which answers 200. A modular client alone fails to decode the page, and succeeds when
  the call decodes no response, as a delete does (`TestClassicError_DiffersForPrivateLinkRedirect`).
  `Transport` answers the page with the 403 and `PRIVATE_LINK_VALIDATION_ERROR` of the classic
  client, and `ClassicError` converts it, so every call fails as on the classic path
  (`TestTransport_PrivateLinkMatchesClassicClient`,
  `TestClassicError_ConvertsPrivateLinkValidationError`).
  `TestNewClient_DebugLogShowsWhatIsSent` runs a delete through `NewClient`. No test runs it
  through `ClassicCall`, or through the resource.

The reflection helpers differ from their classic twins in ways that change what a resource has
to write, which step 3 covers:

- `DataToStruct` leaves a zero value out that a classic field without `omitempty` sends
  (`TestDataToStruct_LeavesOutZeroThatClassicSends`). It returns an error where
  `common.DataToStructPointer` panics or takes the first of two blocks
  (`TestDataToStruct_ValueDoesNotFit`).
- `StructToSchema` marks nothing required (`TestStructToSchema_SameAsClassicTwin`).
- `StructToData` returns the error of writing a default, which `common.StructToData` drops
  (`TestStructToData_Errors`).

## Check before you move a resource

The helpers were built for the four calls of `databricks_mlflow_model`. Check the resource for
each of these first.

| Not covered | What happens | How to tell |
|---|---|---|
| A oneof group or another interface field, such as a body stream | `StructToSchema` panics. `DataToStruct` and `StructToData` return an error once the schema has the name of the field, whatever it holds. | a field of an interface type in the struct of the modular SDK. `TestStructToSchema_UnsupportedFieldTypes` lists the shapes. |
| An embedded struct | `StructToSchema` panics, and the other two don't look into it. A wrapper struct that embeds `common.Namespace`, as `MlflowModelSchemaStruct`, can't build the modular schema. | the classic schema struct embeds another struct |
| A message that contains itself | `StructToSchema` panics, where `common.StructToSchema` cuts at the depth that the resource declares | the classic schema struct has a `MaxDepthForTypes` method, or the panic names the message |
| An API name that is not snake_case | `StructToSchema` keeps the name, and SDKv2 rejects it as an attribute name | `InternalValidate` of the resource fails. `TestStructToSchema_KeepsNameThatIsNotInSnakeCase` shows it. |
| A set | `StructToSchema` never makes one, and `DataToStruct` returns an error for a field whose schema is a set | `schema.TypeSet` in the classic schema, or a `slice_set` tag |
| A field mask (`types.FieldMask` in `core/types` of the modular SDK) | `StructToSchema` gives it no attribute, and `DataToStruct` and `StructToData` skip it, so the request goes out without one. Set it in the resource. | a field of type `types.FieldMask` in the request. `TestDataToStruct_SkipsFieldsAndEntries` shows it. |
| A timestamp or a duration (`types.Time`, `types.Duration`) | both are string attributes, written and sent in the form of the modular SDK, difference 15. `DataToStruct` returns an error for a string that the type doesn't decode, such as `5m` (`TestDataToStruct_ValueDoesNotFit`). | a field of one of the two types in the request or the response. Compare the text that the API sends with the form of the SDK. |
| A zero value that must be sent | `DataToStruct` leaves it nil, so the request doesn't have it | a field without `omitempty` in the classic struct, or a `ForceSendFields` entry in the resource |
| A method that returns an iterator | its requests run after the method has returned, outside `ClassicCall`: no classic retries, rate limit or errors. One retrier would also count the pages together. | the resource lists, or the method of the modular client returns an iterator |
| A long-running operation | no helper waits for one. `ClassicCall` runs exactly one call. | the classic resource calls a method that waits, or polls |
| An account-level client | every helper was tested with a workspace client only. `ClientOptions` passes the account ID of the config, and no test sends a call that uses it. | the resource calls `c.AccountClient()` or has an account-level variant |
| An endpoint that can answer with an empty body, or with a number for a string field | the call fails on the modular path, differences 13 and 14 | pin both with a decoding journey for every call whose response has fields |
| A read of a cluster, a job or a job run | the classic client reports a 400 with `INVALID_PARAMETER_VALUE` as a missing resource for these three paths. `ClassicError` can't, so the resource stays in the state with an error. Recognize the response in the resource. | the resource reads `clusters/get`, `jobs/get` or `jobs/runs/get`. `TestClassicError_DiffersByRequestOverride` shows it. |
| A `Headers` hook, or an authentication type that signs the request | the hook wins over the headers of the call, and one that takes longer than `http_timeout_seconds` no longer fails the attempt, difference 28. A strategy that reads the method, the URL or the body of the request gets `GET /` instead. | `Headers` set on the config, which the provider doesn't do, or a custom credentials strategy |
| A call that streams its response, or sends a body that is not JSON | the debug log differs, difference 27: the request body is logged as `<io.Reader>`, and a streamed response to a request without `Accept` is copied while the caller reads it | the method of the modular client takes or returns a reader |
| A config that is made for every call | `Limiter` keeps one budget for each config pointer, for as long as the process runs, so such calls share no rate limit | the resource gets its client from `ClientForHost` and not from `WorkspaceClientUnifiedProvider` |
| Unit tests on a mocked classic client | on the modular path they fail with `modular SDK client: config is nil`, because the mock has no config | `MockWorkspaceClientFunc` in the tests of the resource |

## Removing the classic path

Nothing has been removed yet, so this list follows from the code and has not been done once.
When a resource no longer needs its classic path:

1. Delete the classic operations and the classic schema: in
   [`resource_mlflow_model.go`](../mlflow/resource_mlflow_model.go) that is the four functions,
   the `common.StructToSchema` call and `MlflowModelSchemaStruct`. Keep `CustomizeDiff`.
2. Delete the switch: `mlflowModelWithSdkGo`, `mlflowModelPickPath` and
   `mlflowModelResourceType`. The resource takes its schema from `mlflowModelSdkGoSchema` and
   its operations from `mlflowModelSdkGo`, and no longer asks `Enabled`.
3. Delete `mlflowModelOnBothPaths` and the `classic` subtests, and call the test bodies
   directly. Delete `TestResourceMlflowModelSdkGoSchemaEqualsClassic`, which has nothing left
   to compare with: `TestResourceMlflowModelSchema` stays as the pin of the schema.
4. In every replay package delete `TestReplay`, which clears the variable, and keep the
   journeys running through one function. Delete `testdata_classic/` and move the twins from
   `testdata_sdkgo/` into `testdata/`, so that CI counts them.
5. Remove the classic branch of `wantUserAgent` in
   [`user_agent_test.go`](../qa/replay/resources/mlflow_model/user_agent_test.go).
6. Write a changelog entry for users. The differences above that change an outcome become the
   behavior of the resource, so the entry names them.
7. Keep the helpers. Only `Enabled` goes, once no resource has two paths.
