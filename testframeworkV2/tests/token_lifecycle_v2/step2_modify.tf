# Step 2 — modify the same resource (same address) with new attrs.
# State carries from step 1; provider plans an in-place update.

provider "databricks" {}

resource "databricks_token" "pat" {
  comment          = "tfv2-token-lifecycle-step-2-modified"
  lifetime_seconds = 7200
}
