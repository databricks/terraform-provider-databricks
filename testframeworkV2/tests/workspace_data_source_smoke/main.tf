# Green-path smoke for `data.databricks_mws_workspaces` against the current branch
# (tested via `version: local`). NOT tied to a specific GitHub issue — this fixture
# guards the happy path so that if the data source ever regresses again post-#5672,
# `tfv2 run -r testframeworkV2/` catches it before any release ships.
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

# Asserting `length(...)` >= 0 is a no-op assertion — the actual smoke is that
# the data source READS without error. The output keeps Terraform from optimising
# the read away (data sources without referenced attributes can be skipped).
output "workspace_count" {
  value = length(data.databricks_mws_workspaces.all.ids)
}
