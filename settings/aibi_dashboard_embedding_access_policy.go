package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Automatic Cluster Update setting
var aibiDashboardEmbeddingAccessPolicySettingFieldMask = "aibi_dashboard_embedding_access_policy.access_policy_type"

var aibiDashboardEmbeddingAccessPolicySetting = workspaceSetting[settings.AibiDashboardEmbeddingAccessPolicySetting]{
	settingStruct: settings.AibiDashboardEmbeddingAccessPolicySetting{},
	customizeSchemaFunc: func(s map[string]*schema.Schema) map[string]*schema.Schema {
		common.CustomizeSchemaPath(s, "aibi_dashboard_embedding_access_policy", "access_policy_type").SetRequired()
		return s
	},
	readFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (*settings.AibiDashboardEmbeddingAccessPolicySetting, error) {
		return w.Settings.AibiDashboardEmbeddingAccessPolicy().Get(ctx, settings.GetAibiDashboardEmbeddingAccessPolicySettingRequest{
			Etag: etag,
		})
	},
	updateFunc: func(ctx context.Context, w *databricks.WorkspaceClient, t settings.AibiDashboardEmbeddingAccessPolicySetting) (string, error) {
		t.SettingName = "default"
		res, err := w.Settings.AibiDashboardEmbeddingAccessPolicy().Update(ctx, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
			AllowMissing: true,
			Setting:      t,
			FieldMask:    aibiDashboardEmbeddingAccessPolicySettingFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
	deleteFunc: func(ctx context.Context, w *databricks.WorkspaceClient, etag string) (string, error) {
		res, err := w.Settings.AibiDashboardEmbeddingAccessPolicy().Update(ctx, settings.UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{
			AllowMissing: true,
			Setting: settings.AibiDashboardEmbeddingAccessPolicySetting{
				Etag:        etag,
				SettingName: "default",
				AibiDashboardEmbeddingAccessPolicy: settings.AibiDashboardEmbeddingAccessPolicy{
					AccessPolicyType: settings.AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains,
				},
			},
			FieldMask: aibiDashboardEmbeddingAccessPolicySettingFieldMask,
		})
		if err != nil {
			return "", err
		}
		return res.Etag, err
	},
}
