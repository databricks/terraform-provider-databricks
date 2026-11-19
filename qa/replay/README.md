# Replay test suites

These suites exercise provider behavior against local HTTP fixtures without live
Databricks credentials. See [authoring guidance](../terraformtest/AUTHORING.md)
and the [YAML reference](../terraformtest/README.md).

## Organization

- `features/<feature>/`: shared provider behavior, such as authentication,
  configuration validation, and workspace routing.
- `resources/<resource>/`: resource behavior and related data sources.
- Each Terraform suite has a small Go discovery wrapper and a `testdata/` folder.
  Each YAML filename becomes a subtest name.
- CI runs each package under `go test`'s default 10-minute timeout, so a resource whose journeys take more than a few minutes splits them by theme into sibling packages, such as `resources/volume_routing/` next to `resources/volume/`. Add a journey to the package of its theme.

Helper-only Go tests use YAML HTTP transcripts through
`clienttest.WorkspaceClient`. Keep their transcripts in a nested folder when
sharing a suite with Terraform discovery.

## Running suites

Install Terraform on PATH or set `TF_ACC_TERRAFORM_PATH`, then run:

```sh
go test -short -race -v ./qa/replay/...
go test -short -v ./qa/replay/features/provider_config -run '/TestAccToken_ProviderConfig_EmptyID$'
```

No `TF_ACC`, Databricks credentials, provider installation, or live environment
is required. Suites are also discovered by the repository-wide Go tests.
