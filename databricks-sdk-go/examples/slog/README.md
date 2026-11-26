# slog

This example shows how to use [slog][slog] with this library.

[slog]: https://pkg.go.dev/golang.org/x/exp/slog

## Output

With JSON output:

```text
{"time":"2023-03-10T08:33:11.318122+01:00","level":"DEBUG","msg":"Loading config via environment","global":true}
{"time":"2023-03-10T08:33:11.318411+01:00","level":"DEBUG","msg":"Loading config via config-file","global":true}
{"time":"2023-03-10T08:33:11.318552+01:00","level":"DEBUG","msg":"Loading DEFAULT profile from [...]/.databrickscfg","global":true}
{"time":"2023-03-10T08:33:11.31862+01:00","level":"DEBUG","msg":"Attempting to configure auth: pat","global":false}
{"time":"2023-03-10T08:33:12.170168+01:00","level":"DEBUG","msg":"GET /api/2.0/preview/scim/v2/Me\n< HTTP/2.0 200 OK\n
```

With text output:
```text
time=2023-03-10T08:33:44.219+01:00 level=DEBUG msg="Loading config via environment" global=true
time=2023-03-10T08:33:44.219+01:00 level=DEBUG msg="Loading config via config-file" global=true
time=2023-03-10T08:33:44.219+01:00 level=DEBUG msg="Loading DEFAULT profile from [...]/.databrickscfg" global=true
time=2023-03-10T08:33:44.220+01:00 level=DEBUG msg="Attempting to configure auth: pat" global=false
time=2023-03-10T08:33:45.077+01:00 level=DEBUG msg="GET /api/2.0/preview/scim/v2/Me\n< HTTP/2.0 200 OK\n
```
