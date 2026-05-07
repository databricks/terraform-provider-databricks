package mws

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsWorkspaces() common.Resource {
	type mwsWorkspacesData struct {
		Ids map[string]int64 `json:"ids" tf:"computed"`
	}
	r := common.DataResource(mwsWorkspacesData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*mwsWorkspacesData)
		workspaces, err := NewWorkspacesAPI(ctx, c).List(c.Config.AccountID)
		if err != nil {
			return err
		}
		data.Ids = map[string]int64{}
		for _, v := range workspaces {
			data.Ids[v.WorkspaceName] = v.WorkspaceID
		}
		return nil
	})
	r.SkipProviderConfigStatePopulation = true
	deprecateProviderConfig(r.Schema)
	return r
}

// deprecateProviderConfig marks the auto-injected provider_config block (added
// by common.DataResource via AddNamespaceInSchema) and its nested workspace_id
// as deprecated for account-only data sources. These data sources have no
// workspace context, so the field has never had a meaningful effect.
func deprecateProviderConfig(s map[string]*schema.Schema) {
	pc, ok := s["provider_config"]
	if !ok {
		return
	}
	pc.Deprecated = "provider_config has no effect on this account-only data source and will be removed in a future major release."
	if elem, ok := pc.Elem.(*schema.Resource); ok {
		if ws, ok := elem.Schema["workspace_id"]; ok {
			ws.Deprecated = "workspace_id is ignored for account-only data sources."
		}
	}
}
