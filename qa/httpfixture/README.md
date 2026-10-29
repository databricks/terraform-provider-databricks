# HTTP fixture replay

`NewCaptureTransport(base)` collects the same exchange format from an HTTP
transport. Call `StartStep(label)` before dispatch, then `Snapshot()` after
requests finish. Concurrent calls retain dispatch order, without enabling
ordered replay. Transport/body-read failures invalidate the entire recording.
Bodies remain readable by the SDK; multipart boundaries are omitted, and binary
content uses base64. Authentication and cookie headers are excluded. API
payloads are retained in memory: the recorder must redact secrets and normalize
profile-specific values before saving them. This transport does not load
profiles, enable live requests, or write fixtures itself.

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

## Files and non-JSON bodies

A non-null body selects exactly one representation:

```yaml
body:
  text: |
    # comment
# Or: body: {base64: "AP8K"} for the bytes 00 ff 0a.
```

`json` compares parsed JSON; `text` and `base64` compare exact bytes, including
whitespace. An empty text/binary body matches zero bytes. Use base64 for bytes
that cannot be represented as UTF-8. Responses support these same three modes,
with default content types `application/json`, `text/plain; charset=utf-8`, and
`application/octet-stream`, respectively. Explicit response headers override them.

Multipart requests describe the decoded form rather than its random boundary:

```yaml
body:
  multipart:
    - name: path
      body: {text: /Shared/script.py}
    - name: content
      filename: content
      headers: {Content-Type: application/octet-stream}
      body:
        text: |
          # comment
```

Every part needs a name and a `body.text` or `body.base64`. Matchers compare
names, filenames, bytes, multiplicity, and explicitly listed part headers.
Additional part headers are ignored. Part order and the transport boundary are
ignored; missing or extra parts fail. The request must use `multipart/form-data`.
Multipart responses and JSON matchers on non-JSON bodies are unsupported.
Header and query matchers work with every body representation.
