# Regression-guard for the apply-then-downgrade sibling of issue #5678.
#
# When databricks_token is created with v1.113.0 then refreshed under
# v1.114.0, v1.114.0's post-Read hook (common/resource.go::
# populateProviderConfigInState) writes provider_config = [{workspace_id =
# "<resolved>"}] to state. Rolling back to v1.113.0 then re-running plan
# diffs the polluted state against the (provider_config-less) config and
# v1.113.0's auto-ForceNew sweep stamps provider_config.ForceNew=true,
# producing a "# forces replacement" annotation. Applying this plan would
# destroy and recreate the real token.
#
# This fixture is a perpetual canary: step 3's plan_match: 'No changes'
# fails today because the bug is present (PR #5667 + #5492 revert chain
# on main), and turns green once a forward-rolled fix prevents the
# destructive plan diff.

provider "databricks" {}

resource "databricks_token" "pat" {
  comment          = "tfv2-rollback-err-test"
  lifetime_seconds = 3600
}
