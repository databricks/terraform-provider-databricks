# Reproducer HCL for https://github.com/databricks/terraform-provider-databricks/issues/5672.
#
# Version pinning: this file does NOT pin a databricks version. The framework writes a
# `_tfv2_versions_override.tf` file per step that pins to the desired version via Terraform's
# documented override-merge semantics. You may optionally add your own
# `terraform { required_providers { databricks = { source = "databricks/databricks", version = ">= 1.113.0" } } }`
# block here for IDE/standalone-`terraform plan` workflows; the framework's override will
# transparently win on `version` at test time. Empirically validated (see DESIGN.md Appendix A).
#
# Auth: the framework injects DATABRICKS_CONFIG_PROFILE=<test.yaml profile> into the
# terraform subprocess env. The Databricks SDK reads ~/.databrickscfg's section to
# resolve host, account_id, and credentials. The provider block here is intentionally
# empty (auth comes from the profile, not from inline arguments).

provider "databricks" {
  alias = "accounts"
}

data "databricks_mws_workspaces" "all" {
  provider = databricks.accounts
}

output "workspace_count" {
  value = length(data.databricks_mws_workspaces.all.ids)
}
