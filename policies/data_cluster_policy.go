package policies

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceClusterPolicy returns information about cluster policy specified by name
func DataSourceClusterPolicy() *schema.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Id                        string `json:"id,omitempty" tf:"computed"`
		Name                      string `json:"name,omitempty" tf:"computed"`
		Definition                string `json:"definition,omitempty" tf:"computed"`
		Description               string `json:"description,omitempty" tf:"computed"`
		FamilyId                  string `json:"family_id,omitempty" tf:"computed"`
		FamilyDefinitionOverrides string `json:"family_definition_overrides,omitempty" tf:"computed"`
		IsDefault                 bool   `json:"is_default,omitempty" tf:"computed"`
		MaxClustersPerUser        int    `json:"max_clusters_per_user,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		policy, err := w.ClusterPolicies.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Id = policy.PolicyId
		data.Definition = policy.Definition
		data.Description = policy.Description
		data.FamilyId = policy.PolicyFamilyId
		data.FamilyDefinitionOverrides = policy.PolicyFamilyDefinitionOverrides
		data.IsDefault = policy.IsDefault
		data.MaxClustersPerUser = int(policy.MaxClustersPerUser)
		return nil
	})
}
