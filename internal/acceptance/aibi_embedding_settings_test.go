package acceptance

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccAiBiEmbeddings(t *testing.T) {
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
`,
		Check: func(s *terraform.State) error {
			w, err := databricks.NewWorkspaceClient()
			require.NoError(t, err)
			id := s.RootModule().Resources["databricks_aibi_dashboard_embedding_approved_domains_setting.this"].Primary.ID
			require.Equal(t, "global", id)
			setting, err := w.Settings.AibiDashboardEmbeddingApprovedDomains().Get(context.Background(),
				settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{})
			require.NoError(t, err)
			assert.Equal(t, []string{"test.com", "test2.com"}, setting.AibiDashboardEmbeddingApprovedDomains.ApprovedDomains)
			return nil
		},
	})
}
