# Deferred migration problems

Keep problems found while copying acceptance tests here. Defer provider and
engine fixes in separate PRs from the replay conversions.
The per-case [pending inventory](PENDING.md) maps the remaining baseline cases
to these blockers. Reproduced failures and risks found by source inspection are
distinguished below; not every affected sibling has been executed. Affected scenarios stay pending in
[MIGRATION.md](MIGRATION.md) until their behavior and assertions can be preserved.

| ID | Kind | Problem and affected cases | Follow-up |
| --- | --- | --- | --- |
| P1 | Provider bug | Concurrent cache reads in `common/unified_provider.go` are not protected by the mutex used for writes. A replay of `scim/user_test.go`'s `TestAccUserResource`, which creates three independent users, reports data races under `-race`. The AWS branch of `TestUcAccStorageCredential`, with independent IAM-role and Cloudflare credentials, also reports cache races and shared SDK configuration races during client initialization. | Cache locking is addressed separately in [PR 6120](https://github.com/databricks/terraform-provider-databricks/pull/6120). SDK configuration/authentication races remain tracked. Copy tests independently of the fix, retain ordinary parallelism, and report race-detector failures without suppressing them. |
| P2 | HTTP fixture limitation | Workspace-file uploads use multipart bodies; volume-file uploads/downloads use raw bytes. The fixture handler currently accepts JSON bodies only, so even `TestAccWorkspaceFileBase64` cannot be copied faithfully. | Support matching the actual upload payload, including content and form fields, without depending on a random multipart boundary. |
| P3 | Local input limitation | Source-file cases need files available at stable paths inside Terraform's isolated working directory. Examples include `TestAccDatabricksDBFSFile_CreateViaSource`, notebook DBC/IPYNB updates, and workspace-file source/empty/ZIP cases. | Define how replay supplies local input files. Preserve the `source` path and file-reading behavior instead of changing those cases to `content_base64`. `TestAccNotebook_ProviderConfig_Mismatched` and `TestAccNotebookData_ProviderConfig_EmptyID` also need their local input files: offline attempts fail on the missing source before reaching the expected diagnostic. |
| P4 | Assertion limitation | Remaining SDK checks inspect secret-scope ACL inventories/current-user ownership, post-destroy settings resets and home-cleanup error handling, continuous-job run histories, or pipeline `ForceSendFields` for explicit zero workers. These values are not available through the existing YAML state/plan assertions or equivalent read-only data sources. | Retain equivalent Go/SDK checks before counting these scenarios as copied. Counts exposed by Terraform outputs, directory/cluster data-source reads, and the eleven helper-level Go replays are now copied; those are no longer blockers. Response fixtures alone do not preserve an SDK-only assertion. |
| P5 | Test-step limitation | YAML cannot express Terraform import steps and import-state verification. For example, `TestUcAccCatalogForceDestroyConsistentAfterImport` checks consistency after import. | Add import-step support before counting these scenarios as copied. A create/read/destroy port would omit the behavior under test. |
| P6 | Multiple-provider limitation | `internal/acceptance/tf_provider_test.go` tests deferred configuration with an additional noop provider. It is already hermetic, but the YAML runner currently registers only Databricks. | Decide how to retain this existing test in the migration inventory or support its additional provider. Do not replace it with a single-provider scenario. |
| P7 | Provider/API mismatch | `TestUcAccMetastoreWithOwnerUpdates` changes the metastore name in its third step. The account update model sends owner and sharing settings but has no name field. With a response that preserves the unchanged server name, Terraform reports a nonempty post-apply plan on AWS, Azure, and GCP. | Resolve rename support before copying the full scenario. Do not return an unrequested new name in the response fixture. |
| P8 | Diagnostic mismatch | `TestMwsAccWorkspaceIDHttp_RemoveDefault` and `TestMwsAccWorkspaceIDHttp_RemoveOverrideNoFallback` expect a missing-workspace diagnostic during planning. The current notebook path instead attempts another workspace import. `TestAccShare_ProviderConfig_Invalid` expects an integer-validator diagnostic, but the provider now reports a workspace-client validation error that also explains connection IDs. | Review the original expectations against current behavior, then preserve the agreed diagnostics in replay. Keep these cases pending rather than accepting any error. |
| P9 | Routing/plan mismatch | `TestAccShare_ProviderConfig_Recreate` expects a replacement plan when changing to another workspace. Current planning rejects the workspace ID as mismatched first. The catalog counterpart has the same replacement expectation and also remains pending; it has not yet been replayed. | Resolve the intended routing and plan behavior before porting these scenarios. Preserve the original plan assertions until that decision is made. |
| P10 | SDKv2 schema race | `TestUcAccSchema` creates multiple grants resources. Under `-race`, concurrent `ValidateResourceConfig` calls report writes to the same `AtLeastOneOf` backing slice inside `validateAtLeastOneAttribute` / `removeDuplicates` in the vendored Terraform Plugin SDK. `ResourceGrants` shares this slice across its securable attributes. | Resolve the shared schema slice mutation and replay the unchanged multi-resource scenario under `-race`. Keep its grants/resources independent and leave the source test pending. |
| P11 | Request determinism risk | `catalog/resource_sql_table.go` serializes properties and options by iterating Go maps into SQL strings. Multiple entries can therefore change order inside the JSON statement value; semantic JSON matching does not normalize SQL. This risk was found by source inspection, without a failing replay reproduction. | Verify deterministic SQL generation or define narrow SQL matching before copying the multi-property/option scenarios. Do not wildcard the entire statement and lose request validation. |
| P12 | Post-apply plan limitation | `TestAccJobClusterPolicySparkVersion` explicitly allows a nonempty post-apply plan because policy defaults cause permanent drift. YAML's `plan.expect` checks the pre-apply plan and does not relax HashiCorp's post-apply consistency check. | Define support for the original post-apply expectation before porting this case. Do not replace it with a pre-apply nonempty-plan assertion. |
| P13 | Account/workspace routing | `TestAccCredentialOwner` configures a workspace-level service credential from `UnityAccountLevel` without a workspace ID. Offline replay fails with `managing a workspace-level resource requires a workspace_id, but none was found`. | Establish the original scenario's intended workspace routing before porting it. Do not invent a workspace configuration just to make replay pass. |
| P14 | Provisioning-state read | `TestMwsAccGcpWorkspacesProvisioningToRunning` and `TestMwsAccGcpWorkspacesUnsetExpectedState` cannot finish their first step with a realistic response that omits input-only `expected_workspace_status`. Replay stalls in `ResourceMwsWorkspaces` Read at `WaitForExpectedStatus` while the fixture remains `PROVISIONING`; Read maps the API response into state before retrieving the configured expected status. | Preserve the configured input-only status through Read, then port both full transitions. Do not echo the Terraform-only expected status in API responses to conceal the problem. |
| P15 | Acceptance/schema mismatch | `TestMwsAccNccPrivateEndpointRule_NoSpuriousAccountIdDrift` configures `account_id` on the endpoint rule in its first step. The current replay provider rejects it as an unconfigurable attribute on both AWS and Azure, before any HTTP call. | Reconcile the original scenario with the intended provider implementation/schema before copying it. Removing the first-step field would remove the regression's precondition. |
| P16 | Provider-selection limitation | Share, quality-monitor, and private-endpoint migration cases select SDKv2/Plugin Framework factories per step; the endpoint-rule opt-in test also requires a specific factory. YAML replay uses the production mux without these per-step overrides. | Support faithful provider selection before copying these cases. Replaying both steps through the same implementation would omit the migration under test. |

| P17 | Inventory exception | `TestAccessControlHashFunction` is a local unit test included by the broad AST inventory because its name begins with `TestAcc`. It tests an unexported hash helper, not an HTTP lifecycle. | Keep its existing hermetic test and resolve the inventory classification separately; no replay-engine feature is needed. |
| P18 | Source-test inconsistency | Both user-list-with-groups tests create a `tf-…@databricks.com` user but query `userName eq "me-…@example.com"`, then expect the created user's membership. The NCC-list test configures nonexistent plural resource `databricks_mws_network_connectivity_configs` while referencing the singular resource. | Reconcile source intent before copying. Do not return a user that contradicts the requested equality filter or silently repair the source HCL and count it as unchanged coverage. |
| P19 | Input-state drift | `TestAccPipelineResourcLastModified` sends `expected_last_modified` on update, but Read resets it to zero. A strict replay with a realistic pipeline response fails the non-refresh post-apply empty-plan check. The API's GET response model has no `expected_last_modified` field. | Preserve this input-only value through Read, then copy the complete timestamp/precondition scenario. Do not fabricate a response field to avoid the drift. |
| P20 | Derived-host routing | AWS workspace/token lifecycles derive a second HTTPS workspace host from the deployment name and use it for token creation/read/update. The replay transport permits only its own fixture-server URL. The token-update case also checks the token list through an SDK client for that derived host. | Define explicit local routing for derived hosts and preserve token-list assertions. Do not allow live network access or substitute an unrelated single-host lifecycle. This limitation was found by source inspection. |

| P21 | Provider cache isolation | The process-global SCIM membership cache is keyed only by group ID. Repeating `TestAccGroupMemberResource`, `TestAccGroupDataSplitMembers`, or `TestAccAwsOboTokenResource` with `go test -count=3` reuses the cached fake group across fresh fixture servers and skips an expected GET. All three pass on their first run; the repeat runs fail with an unconsumed fixture. | Scope membership caching to the provider/client or give it an explicit lifecycle. Keep fixtures strict; do not make the missing reads optional or reset the cache from these YAML tests. Three newly copied functions reproduce it; account counterparts use the same cache. |

Of the 42 source functions originally mapped to P1, 30 now have replay copies.
The remaining 12 also need local files/uploads (four), SDK-only checks (six,
including one import lifecycle), or derived-host routing (two). No scenario is
left pending solely because of the client-cache race.

The repository-permissions copies replace the two SDK group-ACL reads with the
provider's Read and HashiCorp's post-apply empty-plan check: the refreshed ACLs
must match both configured group/permission pairs. The original `object_type`
state check remains explicit. Their strict HTTP fixtures assert outgoing
permissions requests; the inferred responses are test inputs, not independent
server-side verification.

Additional P1 reproductions from the later resource batch:

- `jobs/job_test.go`: `TestAccJobRunAsUser`.
- `mws/mws_log_delivery_test.go`: `TestMwsAccLogDelivery`. Concurrent storage
  configuration reads mutate SDK `Config.Host` during authentication while
  credential reads initialize an account client from that same config.
- `internal/providers/pluginfw/products/user/data_users_acc_test.go`:
  `TestWorkspaceDataSourceDataUsers` (two independent users).
- `tokens/obo_token_test.go`: `TestAccAwsOboTokenResource`.
- `permissions/permissions_test.go`: `TestAccPermissions_Alert`,
  `TestAccPermissions_Authorization_Passwords`, `TestAccPermissions_ClusterPolicy`,
  `TestAccPermissions_Directory_Id`, `TestAccPermissions_Directory_Path`,
  `TestAccPermissions_InstancePool`, `TestAccPermissions_RegisteredModel`, and
  `TestAccPermissions_SqlQuery`.

These reproduce cache or SDK initialization races under `-race`. The remaining permission cases and multi-resource UC/identity cases are listed
in PENDING.md. Some also require local files, import steps or SDK-only checks;
a race warning alone is not a reason to postpone conversion. The shared setup uses independent resources; not every
sibling has a separate race reproduction. A passing run alone does
not establish that setup is race-free. The new replay copies retain independent
resources; adding dependencies to serialize them would conceal the bug.

When a problem is resolved, link its fix and update the affected inventory counts.
Keep the original integration tests throughout the migration.

Group-member replay scenarios use distinct fake group IDs: the provider also
has a process-global membership cache keyed only by group ID. Reusing IDs across
independent fixture servers can skip reads and contaminate later tests. Distinct IDs isolate different scenarios, but repeating the same scenario in
one Go process still reuses the ID and exposes P21.
