# Building a Plugin Framework Resource

This guide walks through writing a new resource under
`internal/providers/pluginfw/products/`. It captures the patterns the team has
converged on and the reasoning behind them, so the next resource you migrate
or write from scratch can skip the dead ends.

The worked example throughout is
`internal/providers/pluginfw/products/networks/privateendpointrule/`. Skim it
first; the rest of this document explains why it looks the way it does.

## Why Plugin Framework, not SDKv2

SDKv2 is in maintenance mode at HashiCorp. New resources go into Plugin
Framework (PF). PF gives you:

- First-class `Null` / `Unknown` semantics instead of ambiguous Go zero values.
- Isolated `Plan`, `State`, and `Config` objects instead of a merged
  `ResourceData` blob.
- Strongly typed Go structs mapped via `tfsdk:` tags, no `interface{}` round-trips.
- Native nested attributes instead of Blocks.

The cost is more ceremony per resource. The patterns in this guide exist to
keep that ceremony bounded.

## Project layout

A resource lives in its own directory under
`internal/providers/pluginfw/products/<service>/<resource>/` with these files:

```
resource.go        CRUD methods and the local SDK interface
model.go           tfsdk model struct and translation methods
schema.go          PF schema literal
resource_test.go   unit tests for CRUD and translation
resource_acc_test.go acceptance tests (cloud-touching)
```

Keep each under ~250 lines. If a file grows past that, the resource is doing
too much.

## The model

### One Go struct per resource

The model is your Terraform-side representation. It carries PF types
(`types.String`, `types.List`, etc.), and has `tfsdk:` tags that match the
schema. It is distinct from the SDK's request/response structs.

```go
type model struct {
    ID                          types.String       `tfsdk:"id"`
    NetworkConnectivityConfigId types.String       `tfsdk:"network_connectivity_config_id"`
    RuleId                      types.String       `tfsdk:"rule_id"`
    DomainNames                 types.List         `tfsdk:"domain_names"`
    GcpEndpoint                 []gcpEndpointModel `tfsdk:"gcp_endpoint"`
    ...
}

type gcpEndpointModel struct {
    PscEndpointUri    types.String `tfsdk:"psc_endpoint_uri"`
    ServiceAttachment types.String `tfsdk:"service_attachment"`
}
```

Do **not** put `tfsdk:` tags on the SDK's request/response types. The SDK
should have zero knowledge of Terraform. Translation is your problem.

### Translation methods on the model

The model owns its own translation, with composition for nested types:

```go
func (m *model) toCreateRequest(ctx context.Context) (*settings.CreatePrivateEndpointRuleRequest, diag.Diagnostics)
func (m *model) toUpdateRequest(ctx context.Context, prev model) (*settings.UpdateNccPrivateEndpointRuleRequest, diag.Diagnostics)
func (m *model) fromAPI(ctx context.Context, rule *settings.NccPrivateEndpointRule) diag.Diagnostics

func gcpEndpointFromAPI(gcp *settings.GcpEndpoint) []gcpEndpointModel { ... }
```

The CRUD methods call `plan.toCreateRequest(ctx)` and `state.fromAPI(ctx, rule)`.
They never plumb raw fields through. This keeps the resource code readable and
isolates the "Terraform type to Go type" boundary in one place.

Return `diag.Diagnostics`, not `error`, from any helper that may produce
user-facing diagnostics. PF's diagnostics carry richer information (summary,
detail, severity, multiple per response) that `error.Error()` loses.

### emptyModel() for partial-construction paths

PF's reflection requires list-typed fields to declare their element type. A
zero `types.List{}` has `DynamicPseudoType` and fails the round-trip through
`State.Set`. Any code path that constructs a model from scratch (not via
`req.Plan.Get` or `req.State.Get`, which populate from the schema) must
initialize list fields:

```go
func emptyModel() model {
    return model{
        DomainNames:   types.ListNull(types.StringType),
        ResourceNames: types.ListNull(types.StringType),
    }
}
```

Use this in `ImportState` and in the test harness. One source of truth for
"a model with valid framework-typed defaults."

### Do not use the `converters` package

`internal/providers/pluginfw/converters` provides reflective `GoSdkToTfSdkStruct`
and `TfSdkToGoSdkStruct` helpers. They look like they save you the `fromAPI`
boilerplate. Don't use them for hand-curated resources. Reasons:

1. **Compile-time safety vanishes.** A field rename in the SDK produces a
   runtime "destination struct does not have field X" error at apply time, not
   a build failure. Errors end with "please report this to the provider
   developers", admitting the package is brittle.
2. **Hidden behavioral rules.** Zero-value structs silently become `null`.
   Pointer-to-struct wraps in a single-element list. The same special case
   for `*json.RawMessage` appears in two branches of the same function.
3. **Boilerplate moves, doesn't disappear.** You still implement
   `ComplexFieldTypeProvider` to declare nested types.

The converters are meant for autogenerated resources where the model and the
SDK are regenerated together. For hand-curated resources, write `fromAPI` by
hand. It's mechanical, typesafe, and ~30 lines of one-time work.

## The schema

Use native nested attributes:

```go
"os_profile": schema.SingleNestedAttribute{...}
"network_interfaces": schema.ListNestedAttribute{...}
```

Not Blocks. Blocks are SDKv2 legacy; they exist in PF only for backward
compatibility with existing state files. For a new resource, never reach for
`schema.Block`. For a resource migrated from SDKv2 with existing user HCL,
you may need to keep Blocks for state-file compatibility, but flag it as a
follow-up to migrate via a state upgrader.

Mark fields honestly:

- `Required` for inputs the user must provide.
- `Optional` for inputs the user may provide.
- `Computed` for outputs the server sets.
- `Optional + Computed` only when the user may set it and the server may also
  default it.

Do not mark server-set fields as `Optional` "just in case." If the user can't
meaningfully set the value, it's `Computed` only.

`RequiresReplace` plan modifiers belong on fields that the API cannot update;
the framework uses them to plan a destroy + create.

## Dependencies

A resource has two kinds of dependencies: the SDK methods it calls, and the
provider's configured state. Both get hidden behind narrow seams.

### Define a narrow API interface locally

Go's "interfaces belong to the consumer" idiom: name the methods you call,
not the type that implements them. Define a small interface in `resource.go`:

```go
type api interface {
    CreatePrivateEndpointRule(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
    GetPrivateEndpointRule(ctx context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
    UpdatePrivateEndpointRule(ctx context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
    DeletePrivateEndpointRule(ctx context.Context, req settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
}

type resourcePrivateEndpointRule struct {
    api     api
    backoff retrier.BackoffPolicy
}
```

The real SDK type and the test fake both satisfy this interface structurally.
The resource never imports `settings.NetworkConnectivityInterface` or
`*databricks.AccountClient`.

Trade-off: defining the interface costs you a few lines per resource. The
payoff is that future additions to the SDK's interface cannot leak into your
resource's contract, and the test fake implements only what's actually used.

### Use the basic SDK methods, not the path-encoded shortcuts

The SDK exposes both forms:

```go
// Basic
DeletePrivateEndpointRule(ctx, settings.DeletePrivateEndpointRuleRequest{
    NetworkConnectivityConfigId: nccId,
    PrivateEndpointRuleId:       ruleId,
})

// Shortcut (avoid)
DeletePrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(ctx, nccId, ruleId)
```

Use the basic form. The shortcuts are autogenerated convenience that pollute
the interface with very long method names, and they don't carry forward to
methods that take additional fields beyond path parameters.

### Resolve provider state via the `provider` package

Resources never import `common.DatabricksClient`. Instead, the `provider`
package exposes dependency-getter functions that take `req.ProviderData` and
return the concrete dependency:

```go
ac, d := provider.AccountClient(req.ProviderData)
```

The signature accepts `any` because both `resource.ConfigureRequest` and
`datasource.ConfigureRequest` expose `ProviderData` as `any`; one function
works for both contexts.

The contract:

- `(client, nil-diags)` on success.
- `(nil, nil-diags)` when `ProviderData` is nil; acceptance-test setup path.
  Caller bails without error.
- `(nil, diags-with-error)` on a real failure.

To add a new dependency-getter (e.g., `WorkspaceClient`, `AccountID`),
extend `internal/providers/pluginfw/provider/provider.go`. The provider
package is the only place that knows about `common.DatabricksClient`.

## The CRUD methods

### The resource struct

```go
type resourcePrivateEndpointRule struct {
    api     api
    backoff retrier.BackoffPolicy
}
```

Two fields:

- The narrow API interface (see above).
- The retry backoff policy, exposed as a field so in-package tests can inject
  a tight policy. Production resolves it from the package default.

Always declare the interface assertions:

```go
var (
    _ resource.Resource                = &resourcePrivateEndpointRule{}
    _ resource.ResourceWithConfigure   = &resourcePrivateEndpointRule{}
    _ resource.ResourceWithImportState = &resourcePrivateEndpointRule{}
)
```

These catch interface drift at compile time.

### Metadata and Schema

```go
const (
    resourceName     = "mws_ncc_private_endpoint_rule"
    fullResourceName = "databricks_" + resourceName
)

func (r *resourcePrivateEndpointRule) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = fullResourceName
}

func (r *resourcePrivateEndpointRule) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = resourceSchema()
}
```

Two constants tied by compile-time string concat, not two unrelated literals.
The short form goes into the User-Agent tag (`resource=<name>`); the full form
becomes the Terraform resource type.

### Configure

Re-resolve the dependency on every call. Do not cache on the receiver with a
nil-guard:

```go
func (r *resourcePrivateEndpointRule) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
    ac, d := provider.AccountClient(req.ProviderData)
    resp.Diagnostics.Append(d...)
    if ac == nil {
        return
    }
    r.api = ac.NetworkConnectivity
}
```

PF may call `Configure` more than once on the same resource value if the
underlying provider data is updated. A `if r.api != nil { return }` guard
locks in stale data. The work is cheap (the SDK client mutex-caches under the
hood), so just re-resolve.

The `ac == nil` check covers both "acceptance test setup with nil
`ProviderData`" (no diags) and "real error" (diags already appended). Either
way, bail without touching `r.api`.

### Create

Five steps, in order:

1. Read the plan.
2. Build the request via `plan.toCreateRequest(ctx)`.
3. Call the API.
4. Persist intermediate state with the just-returned rule.
5. Poll until the rule leaves CREATING; persist the final state.

The intermediate `State.Set` is what makes the resource destroyable after a
polling failure. Without it, a failed apply leaves an orphaned cloud
resource that Terraform doesn't know how to clean up. Test the invariant:
see `TestCreate_PersistsStateBeforePolling`.

```go
plan.NetworkConnectivityConfigId = types.StringValue(apiReq.NetworkConnectivityConfigId)
plan.ID = types.StringValue(packID(apiReq.NetworkConnectivityConfigId, pendingRule.RuleId))
resp.Diagnostics.Append(plan.fromAPI(ctx, pendingRule)...)
resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
if resp.Diagnostics.HasError() {
    return
}

rule, err := retrier.Run(ctx, retrier.RetryIf(r.backoff, isStillCreating), ...)
```

Don't return from a partial Create with `resp.State` still null. Either set
state or expect a stuck resource.

### Polling

Use `internal/retrier`. The Go SDK's `retries.New` is also available but is
being phased out.

The predicate inspects the value directly. It does **not** smuggle "keep
polling" through a sentinel error:

```go
func isStillCreating(rule *settings.NccPrivateEndpointRule, err error) bool {
    return err == nil && rule.ConnectionState == settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreating
}

rule, err := retrier.Run(ctx, retrier.RetryIf(r.backoff, isStillCreating), func(ctx context.Context) (*settings.NccPrivateEndpointRule, error) {
    return r.api.GetPrivateEndpointRule(ctx, ...)
})
```

After `retrier.Run` returns, interpret terminal states in a switch in the
calling code. The polling function itself stays trivial: a single API call,
no state translation.

Backoff lives as a field on the resource struct, not a package-level
constant. Tests inject a tight policy; production uses the default.

### Polling-state semantics: read the actual contract

Don't guess transient vs. terminal from the state name. Check the API doc or
the proto. For NCC private endpoint rules: CREATING is the only transient
state; PENDING is terminal-success (the rule needs out-of-band customer
approval to reach ESTABLISHED, which Terraform cannot drive). Polling past
PENDING would deadlock `terraform apply` waiting on a customer action that
must happen in their cloud console.

When in doubt, look at `~/universe/estore/namespaces/<service>/latest.proto`
for the authoritative state enum and comments.

### Read

```go
rule, err := r.api.GetPrivateEndpointRule(ctx, ...)
if err != nil {
    if apierr.IsMissing(err) {
        resp.State.RemoveResource(ctx)
        return
    }
    resp.Diagnostics.AddError("failed to read private endpoint rule", err.Error())
    return
}
```

`apierr.IsMissing(err)` is the canonical 404 check. On 404, remove the
resource from state; do **not** return an error. Anything else is a real
failure.

### Update

Skip the API call when nothing updatable changed:

```go
apiReq, d := plan.toUpdateRequest(ctx, state)
resp.Diagnostics.Append(d...)
if resp.Diagnostics.HasError() {
    return
}
if apiReq.UpdateMask != "" {
    updated, err := r.api.UpdatePrivateEndpointRule(ctx, *apiReq)
    ...
}
```

`toUpdateRequest` computes the update_mask by comparing plan to state. An
empty mask means "no updatable field changed"; the API call would be a no-op
but cost a round trip. Test both branches.

### Delete

404 is idempotent; treat it as success:

```go
_, err = r.api.DeletePrivateEndpointRule(ctx, ...)
if err != nil && !apierr.IsMissing(err) {
    resp.Diagnostics.AddError("failed to delete private endpoint rule", err.Error())
}
```

### ImportState

Build the state via a typed `model` literal, not a sequence of
`SetAttribute(path.Root(...))` calls:

```go
func (r *resourcePrivateEndpointRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    nccId, ruleId, err := unpackID(req.ID)
    if err != nil {
        resp.Diagnostics.AddError("invalid import id", err.Error())
        return
    }
    m := emptyModel()
    m.ID = types.StringValue(packID(nccId, ruleId))
    m.NetworkConnectivityConfigId = types.StringValue(nccId)
    m.RuleId = types.StringValue(ruleId)
    resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
```

`emptyModel()` (not a partial literal) because PF requires list-typed fields
to declare their element type. A bare `&model{ID: ..., RuleId: ...}` literal
fails round-trip with "Received framework type from provider logic:
DynamicPseudoType".

ImportState sets just enough state for the subsequent Read to take over.
Read fetches the rule and fills in the rest.

For simple resources where the import ID equals the resource ID, PF ships
`resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)`. Use it
if you can. We can't here because the ID is composite.

## Errors and diagnostics

PF's diagnostics are richer than Go errors (summary, detail, severity,
multiple per response). The CRUD method body is the translation boundary.

### From Go errors at API call sites

```go
if err != nil {
    resp.Diagnostics.AddError("failed to create private endpoint rule", err.Error())
    return
}
```

Lowercase summaries that name the operation. The detail is the wrapped error
text.

### From helpers that produce user-facing errors

Return `diag.Diagnostics`, not `error`:

```go
apiReq, d := plan.toCreateRequest(ctx)
resp.Diagnostics.Append(d...)
if resp.Diagnostics.HasError() {
    return
}
```

Batch multiple `Append` calls before a single `HasError` check. The user
sees all decode errors in one apply, not just the first.

### What not to do

- Don't return `error` from a helper that already handles PF-typed inputs.
  You lose the summary/detail/severity structure PF wants for the response.
- Don't wrap errors with `fmt.Errorf("create: %w", err)` expecting PF to
  render the chain nicely. `AddError` calls `err.Error()` once; the chain
  flattens.
- Don't write a `mustNotErr(diags, err)` helper that panics. PF doesn't model
  panics, and a panic during Create takes down the plugin process.

## Testing

### What to test

Behavior, not field-by-field struct echoes:

- **CRUD branches**: each `if err != nil`, each `apierr.IsMissing`, each
  early-return path. These are the bugs that ship.
- **Polling integration**: the predicate + retrier.Run + post-Run state
  interpretation, end to end.
- **ImportState**: happy path and invalid-ID rejection.
- **Translation logic that has branches**: e.g., `toUpdateRequest` mask
  computation. Straight field copies don't need a test.
- **ID parsing**: round-trip fuzz, garbage rejection.

### What not to test

- **Struct-literal echoes**: a test that asserts "every field of fromAPI maps
  the SDK field of the same name" is testing that you typed it twice. Delete.
- **The framework**: don't test that `types.StringValue("x").ValueString() == "x"`.
- **The SDK**: don't write a unit test for the production SDK behavior.

### Testify is an anti-pattern

`stretchr/testify` (`assert`, `require`, `mock`) is a parallel assertion DSL
that fights idiomatic Go. Use stdlib `testing` plus `github.com/google/go-cmp/cmp`:

- `if got != want { t.Errorf(...) }` for simple equality.
- `if diff := cmp.Diff(want, got); diff != "" { t.Errorf("...:\n%s", diff) }`
  for structural comparison. PF types implement `Equal(attr.Value) bool`,
  which cmp picks up automatically. No `AllowUnexported` ceremony needed.
- For mocks: hand-rolled fakes with closure fields, not `mock.Mock`.

### The fake API

Build a fake that satisfies the local `api` interface with closure fields:

```go
type fakeAPI struct {
    create func(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
    get    func(ctx context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
    update func(ctx context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
    del    func(ctx context.Context, req settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
}

func (f *fakeAPI) CreatePrivateEndpointRule(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
    return f.create(ctx, req)
}
// ... etc
```

Each test wires only the methods it exercises. Unset closure fields nil-panic
if the resource calls them unexpectedly: more diagnostic than a mock
framework's "call was not expected".

### The harness

Drive the CRUD methods directly. Build `tfsdk.Plan` / `tfsdk.State` from a
model:

```go
func rawPlan(t *testing.T, ctx context.Context, m model) tfsdk.Plan {
    t.Helper()
    fillListDefaults(&m)
    p := tfsdk.Plan{Schema: resourceSchema()}
    if diags := p.Set(ctx, &m); diags.HasError() {
        t.Fatalf("rawPlan: %v", diags)
    }
    return p
}
```

`fillListDefaults` backfills list fields from `emptyModel()`. The same factory
production code uses, so tests and production are aligned.

`rawState`, `emptyState`, and `readModel` are siblings. The harness is
~30 lines and unlocks the CRUD-method tests.

### Behavioural tests

Construct the resource directly with the fake and exercise the CRUD method:

```go
func TestCreate_PollsUntilNotCreating_ThenPersistsFinalState(t *testing.T) {
    ctx := context.Background()
    var getCalls int
    api := &fakeAPI{
        create: func(...) (*settings.NccPrivateEndpointRule, error) {
            return &settings.NccPrivateEndpointRule{RuleId: "rule-1", ConnectionState: "CREATING"}, nil
        },
        get: func(...) (*settings.NccPrivateEndpointRule, error) {
            getCalls++
            if getCalls < 2 {
                return &settings.NccPrivateEndpointRule{ConnectionState: "CREATING"}, nil
            }
            return &settings.NccPrivateEndpointRule{ConnectionState: "PENDING", VpcEndpointId: "vpce-abc"}, nil
        },
    }
    r := &resourcePrivateEndpointRule{api: api, backoff: tightBackoff}

    req := resource.CreateRequest{Plan: rawPlan(t, ctx, model{...})}
    resp := resource.CreateResponse{State: emptyState()}
    r.Create(ctx, req, &resp)

    // assert on resp.Diagnostics, getCalls, and readModel(t, ctx, resp.State).
}
```

Each test runs in milliseconds. `tightBackoff` (1ms initial, 1ms max) makes
polling tests fast.

### Acceptance tests

Acceptance tests live in `resource_acc_test.go`. They exercise the resource
end to end against real Databricks accounts. They run when `CLOUD_ENV` is
set; otherwise they silently `t.Skip()`.

The silent-skip behavior is a known wart. Default `go test ./...` from a
laptop or unconfigured CI prints `--- SKIP` and exits 0, giving positive
signal for zero work done. A future cleanup will gate acceptance tests behind
a `//go:build acceptance` tag so they don't compile into the default test
binary. Until then, treat the unit suite as your only safety net in CI.

## Common gotchas

### types.List zero-value has no element type

PF reflection requires list-typed fields to declare their element type at
conversion time. The zero `types.List{}` value has `DynamicPseudoType`, which
fails the round-trip through `State.Set` / `Plan.Set` with a Value Conversion
Error.

Production code paths that construct a model from scratch (rather than
reading via `req.Plan.Get` or `req.State.Get`, which populate from the
schema) must initialize list fields. Use `emptyModel()` as the factory.

### Configure may be called more than once

PF can call Configure repeatedly on the same resource value if the underlying
provider data updates. Don't cache the resolved dependency behind a nil-guard
on the receiver. Re-resolve on every call.

### The schema-parity test is a trap

A test that asserts "the new PF schema matches the legacy SDKv2 schema
attribute-by-attribute" is testing parity with the thing you're replacing,
not the design you want. For migrations where state-file compatibility
matters, keep the parity for the launch and add a TODO to tighten the schema
later (Optional+Computed to Computed-only, Block to SingleNestedAttribute)
via state upgraders.

### CLOUD_ENV gates acceptance tests via skip, not via build tags

See above. Unit tests are the only thing that runs reliably in CI today.
Don't rely on acceptance tests catching CRUD regressions.

## Anti-patterns

1. **Reflection-based conversion** (`internal/providers/pluginfw/converters`).
   Use hand-written `fromAPI` / `toCreateRequest` for hand-curated resources.
2. **Testify** (`assert`, `require`, `mock`). Use stdlib + go-cmp + closure
   fakes.
3. **Caching guard in Configure**. Re-resolve every call.
4. **Direct `common.DatabricksClient` dependency in resources**. Go through
   `provider.AccountClient` (or sibling helpers).
5. **`FromResource` / `FromDataSource` function suffixes**. Take
   `req.ProviderData any` instead.
6. **Long shortcut SDK methods** like
   `GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId`.
   Use the basic forms that take a request struct.
7. **Returning `error` from PF helpers**. Return `diag.Diagnostics`.
8. **Struct-literal mirror tests**. Test behavior, not assignment.
9. **`schema.Block`** for a new resource. Use native nested attributes.
10. **Sentinel-error retry signaling**. The retrier predicate inspects the
    value directly.

## Design principles

1. **Resources depend on narrow consumer interfaces, not concrete SDK or
   provider types.** The local `api` interface and the `provider` package
   helpers are the only seams the resource sees.
2. **Translation lives on the model, with composition for nested types.**
   No reflective converters, no inline literals scattered through CRUD
   methods.
3. **Tests drive the CRUD methods directly through a small harness.**
   Hand-rolled fakes, no mock DSL, no acceptance-test crutch.
4. **The framework's diagnostics model is the boundary; everything inside is
   typed.** Helpers return `diag.Diagnostics`; CRUD methods translate Go
   errors at the API call site.
5. **Re-resolve, don't cache.** Configure runs every time; the underlying
   SDK client already handles caching.
