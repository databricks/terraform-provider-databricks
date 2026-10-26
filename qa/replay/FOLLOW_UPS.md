# Deferred migration problems

Keep problems found while copying acceptance tests here. Defer provider and
engine fixes until the migration batches finish.
This is a running list of reviewed cases, not a claim that every remaining
acceptance test has been analyzed. Affected scenarios stay pending in
[MIGRATION.md](MIGRATION.md) until their behavior and assertions can be preserved.

| ID | Kind | Problem and affected cases | Follow-up |
| --- | --- | --- | --- |
| P1 | Provider bug | Concurrent cache reads in `common/unified_provider.go` are not protected by the mutex used for writes. A replay of `scim/user_test.go`'s `TestAccUserResource`, which creates three independent users, reports data races under `-race`. | Protect shared client-cache access, then port the original scenario and verify it with `-race`. Keep its independent resources; do not add dependencies to conceal the race. |
| P2 | HTTP fixture limitation | Workspace-file uploads use multipart bodies. The fixture handler currently accepts JSON bodies only, so even `TestAccWorkspaceFileBase64` cannot be copied faithfully. | Support matching the actual upload payload, including content and form fields, without depending on a random multipart boundary. |
| P3 | Local input limitation | Source-file cases need files available at stable paths inside Terraform's isolated working directory. Examples include `TestAccDatabricksDBFSFile_CreateViaSource`, notebook DBC/IPYNB updates, and workspace-file source/empty/ZIP cases. | Define how replay supplies local input files. Preserve the `source` path and file-reading behavior instead of changing those cases to `content_base64`. |
| P4 | Assertion limitation | Some cases make SDK checks that selected Terraform state attributes cannot express: group-member counts, secret-scope ACL counts/current-user ownership, parent-folder existence, and checks after destroying user/service-principal home directories. `TestAccDataSourceClusters_FilterBy` also checks fields through an SDK read. | Provide equivalent replay assertions or Go checks while retaining every original assertion. Fixture responses alone are not evidence that an SDK-only check was preserved. |
| P5 | Test-step limitation | YAML cannot express Terraform import steps and import-state verification. For example, `TestUcAccCatalogForceDestroyConsistentAfterImport` checks consistency after import. | Add import-step support before counting these scenarios as copied. A create/read/destroy port would omit the behavior under test. |
| P6 | Multiple-provider limitation | `internal/acceptance/tf_provider_test.go` tests deferred configuration with an additional noop provider. It is already hermetic, but the YAML runner currently registers only Databricks. | Decide how to retain this existing test in the migration inventory or support its additional provider. Do not replace it with a single-provider scenario. |

When a problem is resolved, link its fix and update the affected inventory counts.
Keep the original integration tests throughout the migration.
