# Contributing Guide

## This repository is a mirror

This repository is a read-only mirror. The SDK's source of truth lives in
Databricks' internal repository, and changes are published here as part of each
release.

- **Issues:** report bugs and feature requests in
  [GitHub Issues](https://github.com/databricks/databricks-sdk-go/issues).
- **External contributors:** propose improvements in a pull request here.
  Maintainers review these PRs publicly and re-apply approved changes in the
  internal repository for a subsequent release, rather than merging them here.
- **Databricks employees:** make SDK changes in the internal repository.

Only maintainer-approved synchronization PRs carrying the `sync` label are
merged into this mirror.

## Development

Required development tools:

* `go install golang.org/x/tools/cmd/goimports@latest`
* `go install honnef.co/go/tools/cmd/staticcheck@v0.3.2`
* `go install gotest.tools/gotestsum@latest`

## Testing

- All changes to the SDK should be covered by unit tests.
- Tests must be hermetic and must not require a live Databricks workspace or account.
- API behavior should be tested at the HTTP boundary with local fixtures when request routing, pagination, streaming, headers, or response decoding are relevant.
- Unit-level mocks remain appropriate for isolated control-flow behavior.

## Developer Certificate of Origin

To contribute to this repository, you must sign off your commits to certify 
that you have the right to contribute the code and that it complies with the 
open source license. The rules are pretty simple, if you can certify the 
content of [DCO](./DCO), then simply add a "Signed-off-by" line to your 
commit message to certify your compliance. Please use your real name as 
pseudonymous/anonymous contributions are not accepted.

```
Signed-off-by: Joe Smith <joe.smith@email.com>
```

If you set your `user.name` and `user.email` git configs, you can sign your 
commit automatically with `git commit -s`:

```
git commit -s -m "Your commit message"
```
