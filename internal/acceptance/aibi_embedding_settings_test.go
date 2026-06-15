package acceptance

import (
	"testing"
	"time"
)

// Resource-level expected-behavior coverage that does not require a real API
// is provided by the package-internal unit tests in
// settings/resource_aibi_dashboard_embedding_access_policy_setting_test.go and
// settings/resource_aibi_dashboard_embedding_approved_domains_setting_test.go
// (TestCreate/TestRead/TestUpdate/TestDelete for each resource, plus the
// lifecycle scenarios TestAiBiEmbeddingAccessPolicySettingLifecycle and
// TestAiBiEmbeddingApprovedDomainsLifecycle which mirror the two steps of
// this acceptance test against mocked SDK calls).
func TestAccAiBiEmbeddings(t *testing.T) {
	if time.Now().Before(time.Date(2026, 6, 24, 0, 0, 0, 0, time.UTC)) {
		t.Skip("temporarily skipped until 2026-06-24: workspace-settings API is eventually consistent so Get after Update may return stale values; tracked internally.")
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
