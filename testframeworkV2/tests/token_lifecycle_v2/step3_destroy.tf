# Step 3 — destroy. The HCL declares no databricks_token block, so terraform
# plans a destroy of the existing resource. After apply, the resource is gone
# from state — verified by `assert: [{resource: databricks_token.pat, present: false}]`
# in test.yaml.
#
# NOTE: terraform requires at least one block in the file (otherwise
# `terraform plan` errors with "no configuration files"). The empty `provider`
# block satisfies that without declaring any resources to manage.

provider "databricks" {}
