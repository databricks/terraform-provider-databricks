package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Automatic Cluster Update setting
var aibiDashboardEmbeddingApprovedDomainsSettingFieldMask = "aibi_dashboard_embedding_approved_domains.approved_domains"

var aibiDashboardEmbeddingApprovedDomainsSetting = workspaceSetting[settings.AibiDashboardEmbeddingApprovedDomainsSetting]{
	settingStruct: settings.AibiDashboardEmbeddingApprovedDomainsSetting{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "aibi_dashboard_embedding_approved_domains", "approved_domains").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.AibiDashboardEmbeddingApprovedDomainsSetting, error) {
		return w.Settings.AibiDashboardEmbeddingApprovedDomains().Get(ctx, settings.GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.AibiDashboardEmbeddingApprovedDomainsSetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.AibiDashboardEmbeddingApprovedDomains().Update(ctx, settings.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    aibiDashboardEmbeddingApprovedDomainsSettingFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.AibiDashboardEmbeddingApprovedDomains().Delete(ctx, settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{
			Etag: etag,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
