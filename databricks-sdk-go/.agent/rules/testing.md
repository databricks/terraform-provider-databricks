---
description: "Testing standards and anti-patterns for the Databricks SDK."
globs: "**/*_test.go"
alwaysApply: false
---

# Testing

`testify` is **discouraged in new packages** and **tolerated in existing packages**.
The SDK aims to minimize third-party dependencies.

- **New packages**: use Go's standard `testing` package with `google/go-cmp` for complex comparisons
- **Existing packages that already use testify**: keep using it for consistency within that package
- Do NOT introduce testify into a package that doesn't already use it

Table-driven tests are preferred for API client methods:

```go
for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
        // ...
        if diff := cmp.Diff(tc.want, got); diff != "" {
            t.Errorf("mismatch (-want +got):\n%s", diff)
        }
    })
}
```

- Unit tests: `*_test.go` alongside source files
- Tests must be hermetic: no live workspace, account, or cloud credentials
- HTTP-boundary tests: write them in `fixturetests/`, which runs the real clients
  against local servers from the shared HTTP fixture library, rather than with
  `qa.HTTPFixtures`, which the generated tests and older tests next to the code
  use. `fixturetests/` exists only in the internal source repository and isn't
  published; its README describes the helpers
- Tests that need unexported code stay next to it and stub the transport with
  `httpclient/fixtures/`, as the `config/` and `credentials/` tests do
- Generated mocks: `experimental/mocks/` (do not hand-write mocks for generated services)
- Test helpers: `qa/` package (polling and locking helpers)
- Call `t.Helper()` at the start of test helper functions (correct line numbers in failures)
- For configurable options, always test three cases: (1) explicit values, (2) nil/unset default, (3) empty value
- When a feature spans multiple implementations, test the same scenarios consistently in each
- Inject `time.Now` as `func() time.Time` for time-dependent logic — makes tests deterministic
- Always check errors — never use `_` for error returns in tests
- Assert on exact expected values, not partial matches (e.g., full header value, not just prefix)
- Extract repeated test values to shared `const`/`var` to convey intent and avoid silent drift
- Add explicit timeouts to channel/goroutine-based tests to prevent hanging
- Use `defer close(ch)` for channel cleanup to handle error paths safely

## Naming and Style

- Use `want` for expected values and `got` for actual values — not `expected`/`actual`
- Use **named struct fields** in test mock initialization: `&IDToken{Value: tok}` not `&IDToken{tok}`
- Use existing test utilities (e.g., `auth.TokenSourceFn`) rather than creating new
  one-off mock types

## Test Design Philosophy

- **Pragmatic about test value**: don't write tests that can't catch real regressions.
  If removing the tested code wouldn't be caught by the test, the test is not useful
- **Prefer blackbox tests over whitebox tests**: verify observable behavior, not
  implementation details. When structural constraints make blackbox testing impossible,
  consider refactoring for testability — but not at the cost of breaking backward compatibility
- **Test at the right abstraction level**: test behavior through the public interface of
  a component, not through its internals. Tests for `ErrorRetrier` logic belong in
  `errors_test.go`, not in tests for `ApiClient` that happens to use it
- **Mock at the right level**: mock the nearest interface boundary, not deep internals.
  Prefer mocking `PersistentAuth` over the `OAuthClient` it wraps
- **Prefer isolated tests by design**: use dependency injection and mocks rather than
  relying on environment cleanup utilities like `CleanupEnvironment`. Structure tests
  so each validates one layer of logic
- **Don't over-engineer test safety nets**: if the only way a function panics is a
  runtime stack overflow, a recovery wrapper adds no value
- Use mock helpers (e.g., `withMockEnv`) consistently — don't mix `t.Setenv` and
  mock helpers in the same test

## Anti-Patterns (never do these)

- Deleting or skipping flaky tests instead of fixing the root cause
- Tests depending on execution order or shared mutable state
- Non-descriptive test names (`TestFoo`, `TestBasic`) — follow the same pattern as Go
  `Example` functions: `TestFunctionName_description` (e.g., `TestHostType_missingScheme`,
  `TestConfig_HostType_missingScheme` for a method on a type)
- `time.Now` i ntests — use a mock clock
- Testing generated code in `service/` — test your hand-written code that uses it
- Exporting methods solely for test access — use `cmp.Diff` to compare structs,
  or write internal (`_test.go` in same package) tests instead
- Unnecessary concurrency in tests — if goroutines/channels aren't clearly needed, don't use them
- Accepting workarounds without understanding root causes — if `CleanupEnvironment` is
  needed, investigate *why* the environment is dirty
