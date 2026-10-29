# Recorder and fixture changes

- Use replay tests for automated validation. Recorder tests must use local HTTP
  peers and temporary profiles, never a developer's live Databricks credentials.
- Recording is explicitly enabled by `TEST_RECORDER_PROFILE_NAME` and an exact
  anchored `-run` path. Keep the CI and broad-selection guards before profile
  resolution or API calls.
- Workspace/account IDs from the selected profile are translated to the fixture's
  stable IDs. Do not add a general placeholder syntax or edit checkpoint JSON.
- Other resource IDs, volatile fields and assertions can need manual adjustment.
  Keep `any_string` annotations where meaningful; do not weaken unrelated checks
  merely to make a refresh pass.
- Review captured responses, text/base64 payloads and checkpoint data for secrets
  that automatic redaction cannot identify. Use nonsecret test inputs; do not
  commit credentials, tokens, private keys or sensitive user data.
- Preserve HCL, assertion intent and the test's ordering setting when refreshing
  HTTP data. Unexpected failures or incomplete cleanup must leave YAML untouched.
