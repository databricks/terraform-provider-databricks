# YAML Terraform lifecycle tests

Use one small Go wrapper per fixture folder:

```go
func TestTerraform(t *testing.T) {
    terraformtest.Run(t, "testdata")
}
```

The runner discovers every `.yaml` and `.yml` file directly in that folder
(non-recursively), relative to the test package. Each file becomes a Go subtest
named after its filename without the extension. For example,
`testdata/TestTerraformGroupLifecycle.yaml` appears as
`TestTerraform/TestTerraformGroupLifecycle` in test output and CI reports.
Adding a YAML file requires no additional Go test function. Empty folders and
duplicate filename stems fail instead of silently skipping tests.

Subtests run sequentially in filename order, each with independent Terraform
state and an HTTP fixture server. A failing lifecycle does not prevent later
files from running. Select a single file with normal Go filtering:

```sh
go test ./qa/terraformtest -run '^TestTerraform$/^TestTerraformGroupLifecycle$'
```

Only selected files are loaded and validated. `go test -list` lists the Go
wrapper, not dynamically discovered subtests; `go test -v` shows each subtest.
Tests run the actual muxed Databricks provider, supporting both SDKv2 and
Plugin Framework resources.

Install Terraform on `PATH`, or set `TF_ACC_TERRAFORM_PATH` to its executable.
The normal GitHub Actions tests job installs Terraform 1.13.5 with its wrapper
disabled so the test runner can invoke the CLI directly.
These hermetic tests run in the normal Go test suite, including `-short`;
they do not require `TF_ACC`, live credentials, or `CLOUD_ENV`. The runner
does not download Terraform. Client configuration ignores ambient credential
profiles and environment defaults; SDK HTTP traffic is restricted to the
test's local fixture server. Do not use `t.Parallel`: Terraform CLI discovery is
configured through process environment variables.

## Test definition

Use `version: 1`, followed by `steps` and `destroy`. Each step has a human
`name`, literal `terraform` HCL, and `calls`.
State persists across steps. Destroy always runs using the last configuration;
its HTTP exchanges belong under `destroy.calls`, and empty state is always
checked. The testing framework also attempts cleanup after an interrupted test.

Each exchange contains `request` and `response`. Requests specify `method`,
`path`, optional `query` (values are lists of strings), optional `headers`,
and `body`. Responses specify an explicit HTTP `status`, optional `headers`,
and `body`. Use `body: null` for no HTTP body and `body: {json: null}` for JSON
null. Other JSON values go under `body.json`; IDs are literal, without substitution.

Matching uses the complete request signature. Independent signatures are
unordered; repeated identical signatures consume their responses FIFO.
Every HTTP call needs a consumable fixture, including Terraform's plan
refreshes and the testing framework's post-apply consistency reads.
Unexpected calls and unused fixtures fail with the step's index and name.

Optional `host_metadata` supplies a reusable JSON response for
`GET /.well-known/databricks-config` throughout the test, keeping repeated
client initialization out of the consumable resource exchanges.

Destroy always requires an empty resource inventory, including data sources in
child modules. This is enforced by the engine; `destroy` accepts only `calls`,
with no configurable final-state assertion.

## MVP scope

Tests start with empty Terraform state and handwritten fixtures. YAML has
strict field validation: initial checkpoints, recording/setup templates,
plan assertions, plan-only steps, expected client errors, forced HTTP order,
and optional per-step state assertions are not yet supported. Unsupported fields
fail before any Terraform operations. The step/call layout is kept stable
so these capabilities can be added independently.
