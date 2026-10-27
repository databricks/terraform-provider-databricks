# Reproducer HCL for https://github.com/databricks/terraform-provider-databricks/issues/5678.
#
# Bug: state written by v1.114.0 (which silently injects
#   provider_config { workspace_id = "..." } into resource state via PR #5492's
#   post-Read hook) is incompatible with v1.114.1's rolled-back schema. Reading
#   that state on v1.114.1 produces a "must be replaced" plan diff because
#   provider_config is Optional (not Computed) AND no-Update SDKv2 resources
#   stamp ForceNew on every non-Computed top-level field by default.
#
# Affects 19 no-Update SDKv2 resources (databricks_token, databricks_secret*,
# databricks_workspace_binding, databricks_catalog_workspace_binding,
# databricks_*_role, databricks_metastore_data_access, databricks_online_table,
# databricks_vector_search_index, etc — full list in commit 115199fc's body).
#
# databricks_token is the simplest test target: it's a no-Update resource, has
# provider_config in schema (via common.AddNamespaceInSchema), and doesn't
# require any pre-existing infrastructure (no workspace_id, no catalog).
#
# Auth: framework injects DATABRICKS_CONFIG_PROFILE; provider block stays empty.
# Any workspace-level profile reproduces — bug is at the schema/state level,
# completely cloud-agnostic.

provider "databricks" {}

resource "databricks_token" "pat" {
  comment          = "tfv2-issue-5678-repro"
  lifetime_seconds = 1800   # 30 min — long enough for the 4-step plan flow + cleanup
}
