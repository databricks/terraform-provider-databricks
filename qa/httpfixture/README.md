# HTTP fixture replay

`NewHandler()` creates an empty handler for stateless request/response exchanges.
`SetReusableFixtures(calls)` replaces the test-wide reusable fallbacks; `nil`
clears them. Invalid replacements leave the old fixtures intact. Replacing
reusable fixtures does not reset the current step or its failures.
`StartStep` installs a step's consumable fixtures and returns any failures from
the previous step.
`Verify` checks the current step for unexpected calls and unused fixtures.

Matching includes the method, escaped path, query parameters, JSON body, and
listed request headers. JSON object key order and whitespace are ignored;
array order and values matter. An absent body differs from JSON `null`.

## Volatile request values

Use explicit `request.matchers` annotations rather than wildcard strings:

```yaml
request:
  method: POST
  path: /requests
  headers:
    X-Request-ID: captured-header-id
  query:
    request_id: [captured-query-id]
  body:
    json:
      request_id: captured-body-id
      name: fixed-name
  matchers:
    headers:
      X-Request-ID: any_string
    query:
      request_id: any_string
    json:
      - path: [request_id]
        type: any_string
response:
  status: 200
```

`any_string` requires a present string, including an empty string; it does not
match a missing value, JSON null, or another JSON type. Header names are
case-insensitive. Header/query matchers require exactly one value and override
the corresponding sample if supplied; the sample may also be omitted.
Unlisted headers remain ignored. Unmatched query keys/values remain exact,
including multiplicity and value order; extra query keys fail.

JSON paths use literal string keys and zero-based integer array indices, for
example `[items, 0, request_id]`. An explicit `path: []` selects a scalar JSON
body. The target must exist as a string or null sample in `body.json`; null is
only a placeholder and does not make null match. Other fields, object keys,
array lengths/order, and numeric values remain exact. Invalid matcher types,
duplicate targets, and paths absent from the sample fail when fixtures are loaded.

In Go, use `Request.Matchers` with `AnyString` and `JSONMatcher` paths. Matchers
do not modify the original samples, so a future recorder can keep captured
values alongside these annotations. Response payloads remain literal; there
are no captures, substitutions, regexes, or magic payload strings/objects.
`"any_string"` and `"****"` without annotations still match only those strings.

## Consumption and ordering

Independent request signatures may interleave. Identical signatures consume
responses FIFO. Requests and fixture consumption are protected by a mutex.
If matcher signatures overlap, the first matching unconsumed fixture wins;
there is no preference for a more specific fixture. Different captured samples
overridden by the same matchers therefore consume responses FIFO too.
`SetEnforceCallOrder(true)` optionally requires all consumable calls in their
listed order within each step. Out-of-order requests fail without consuming
their fixture. The setting survives step transitions; changing it does not
clear consumption or failures. Reusable fallbacks remain outside ordering,
but cannot bypass an out-of-order matching consumable call.
Test-wide reusable exchanges are intended for invariant client initialization,
such as host metadata; they are not required to be consumed.
