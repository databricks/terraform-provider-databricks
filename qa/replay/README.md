# Replay test suites

These suites copy existing acceptance scenarios without removing the originals.
The originals remain nightly alarms; replay runs without live credentials.
See [authoring guidance](../terraformtest/AUTHORING.md) and the
[YAML reference](../terraformtest/README.md).

## Organization

- `features/<feature>/`: shared provider behavior. A resource can be a carrier
  for routing, configuration validation, or another generic feature.
- `resources/<resource>/`: behavior specific to a resource or data source.
  Keep related data-source cases in that resource's suite.
- Each suite has a tiny Go discovery wrapper and a `testdata/` folder.
  Preserve the source function name as the filename for easy comparison.
  Track migration sources in the inventory rather than comments in the YAML.

A source function can contribute to more than one suite when it covers both
generic and resource-specific behavior. Keep the assertions relevant to each
suite; duplication is preferable to losing coverage.

## Migrating a scenario

Preserve the original steps, configured values, expected errors, and meaningful
checks. Replace random names with fixed fake values: replay files have isolated
Terraform state and a local server. Explain any adaptation beside the step.

Every hand-authored transcript must say that its responses are **inferred, not
recorded**. Build those responses from the source scenario and API response model:
include realistic IDs, configured values, and fields needed for existing checks.
Mark dummy host metadata too. Never run live acceptance tests to obtain them.
A future recorder can replace these responses with sanitized recordings.

Request fixtures assert what the provider sends. Response fixtures are inputs,
not independent evidence of server behavior. A copied SDK read assertion can be
represented by a provider read plus a state assertion when it checks the same
value; document that adaptation and retain the exact outgoing request assertion.
It does not become an independent server-side check.

Represent an out-of-band server change with the next step's refresh response.
Do not add the original test's setup SDK calls as though the provider made them.
Do not drop an SDK-only assertion, import, or unsupported check silently: leave
the scenario pending until it can be represented faithfully.

Keep ordinary request ordering disabled. Do not make resource responses reusable
to hide missing or extra calls. Include automatic destroy's refresh and deletion;
the runner always checks empty state afterward.

## Running suites

Install Terraform on PATH or set `TF_ACC_TERRAFORM_PATH`, then run:

```sh
go test -short -race -v ./qa/replay/...
go test -short -v ./qa/replay/features/provider_config -run '/TestAccToken_ProviderConfig_EmptyID$'
```

No `TF_ACC`, Databricks credentials, provider installation, or live environment
is required. Suites are also discovered by the existing repository-wide Go tests.

## Coverage tracking

[MIGRATION.md](MIGRATION.md) tracks source files from the initial acceptance
inventory. It distinguishes copied scenarios from pending ones; it is not a claim
that every acceptance test is live or that the inventory is exhaustive.
[FOLLOW_UPS.md](FOLLOW_UPS.md) tracks provider bugs and replay limitations found
while porting tests. Their fixes are deferred until the migration batches finish.

The provider-configuration suite copies all 43 cases in the dedicated
`*_provider_config_test.go` files: 42 validation errors and the successful
account-provider/workspace-entitlements regression. That success case uses
step-local metadata fixtures for the account and derived workspace clients.
Other provider-configuration and workspace-routing scenarios remain pending.
