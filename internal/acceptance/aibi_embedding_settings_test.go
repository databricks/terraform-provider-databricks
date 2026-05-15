package acceptance

import (
	"testing"
)

func TestAccAiBiEmbeddings(t *testing.T) {
	if IsGcp(t) {
		t.Skip("Skipping on GCP: workspace-settings estore ramp-up introduces " +
			"~2min eventual-consistency staleness, so GET-after-PATCH returns the " +
			"stale prior value and the post-apply refresh plan is non-empty. " +
			"Re-enable once estore staleness is reduced.")
	}
	WorkspaceLevel(t, Step{
		Template: `
resource "databricks_aibi_dashboard_embedding_access_policy_setting" "this" {
  aibi_dashboard_embedding_access_policy {
    access_policy_type = "ALLOW_APPROVED_DOMAINS"
  }
}

resource "databricks_aibi_dashboard_embedding_approved_domains_setting" "this" {
  aibi_dashboard_embedding_approved_domains {
    approved_domains = ["test.com"]
  }
  depends_on = [databricks_aibi_dashboard_embedding_access_policy_setting.this]
}
`,
	}, Step{
		Template: `
resource "databricks_aibi_dashboard_embedding_access_policy_setting" "this" {
  aibi_dashboard_embedding_access_policy {
    access_policy_type = "ALLOW_APPROVED_DOMAINS"
  }
}

resource "databricks_aibi_dashboard_embedding_approved_domains_setting" "this" {
  aibi_dashboard_embedding_approved_domains {
    approved_domains = ["test.com", "test2.com"]
  }
  depends_on = [databricks_aibi_dashboard_embedding_access_policy_setting.this]
}
`})
}
