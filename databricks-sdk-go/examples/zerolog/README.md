# zerolog

This example shows how to use [zerolog][zerolog] with this SDK.

[zerolog]: https://pkg.go.dev/github.com/rs/zerolog

## Output

With JSON output:

```text
{"level":"trace","global":true,"time":"2023-03-10T08:34:56+01:00","message":"Loading config via environment"}
{"level":"trace","global":true,"time":"2023-03-10T08:34:56+01:00","message":"Loading config via config-file"}
{"level":"debug","global":true,"time":"2023-03-10T08:34:56+01:00","message":"Loading DEFAULT profile from [...]/.databrickscfg"}
{"level":"trace","global":false,"time":"2023-03-10T08:34:56+01:00","message":"Attempting to configure auth: pat"}
{"level":"debug","global":false,"time":"2023-03-10T08:34:57+01:00","message":"GET /api/2.0/preview/scim/v2/Me\n< HTTP/2.0 200 OK\n...
```

With text output:
```text
8:35AM TRC Loading config via environment global=true
8:35AM TRC Loading config via config-file global=true
8:35AM DBG Loading DEFAULT profile from [...]/.databrickscfg global=true
8:35AM TRC Attempting to configure auth: pat global=false
8:35AM DBG GET /api/2.0/preview/scim/v2/Me
< HTTP/2.0 200 OK
...
```
