# Step 1 — create the token with the initial attrs.

provider "databricks" {}

resource "databricks_token" "pat" {
  comment          = "tfv2-token-lifecycle-step-1"
  lifetime_seconds = 3600
}
