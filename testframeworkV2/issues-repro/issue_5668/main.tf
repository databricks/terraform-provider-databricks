# Reproducer HCL for https://github.com/databricks/terraform-provider-databricks/issues/5668.
#
# Bug: workspace-level resources fail at plan with
#   Error: failed to validate workspace_id: failed to get the workspace_id: User not authorized.
# in v1.114.0, when the auth identity cannot reach SCIM /Me (e.g., an Azure SP
# that has not been assigned to the workspace as a user, an AWS OAuth M2M
# identity without workspace user access, etc).
#
# Profile requirement (CRITICAL): the named `profile` MUST be a workspace-level
# profile whose auth identity LACKS /Me access. A standard PAT or admin profile
# will NOT reproduce — the bug fires only when /Me returns "User not authorized".
# In practice this means: an Azure SP / AWS OAuth M2M / GCP SA configured with
# host + workspace credentials but no user-level workspace assignment yet.
#
# If your test environment doesn't have such a profile, the test will fail
# step 2 with "test setup error: profile X authenticated successfully" rather
# than reproducing the bug. Provision a dedicated unassigned-SP profile in
# ~/.databrickscfg, OR skip this test in environments without one.
#
# Auth: framework injects DATABRICKS_CONFIG_PROFILE; provider block stays empty.

provider "databricks" {}

resource "databricks_token" "pat" {
  comment          = "tfv2-issue-5668-repro"
  lifetime_seconds = 600    # 10 min — stays alive for the test, auto-expires after
}
