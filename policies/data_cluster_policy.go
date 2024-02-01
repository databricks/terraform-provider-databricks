package policies

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceClusterPolicy returns information about cluster policy specified by name
func DataSourceClusterPolicy() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Id                              string `json:"id,omitempty" tf:"computed"`
		Name                            string `json:"name,omitempty" tf:"computed"`
		Definition                      string `json:"definition,omitempty" tf:"computed"`
		Description                     string `json:"description,omitempty" tf:"computed"`
		PolicyFamilyId                  string `json:"policy_family_id,omitempty" tf:"computed"`
		PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty" tf:"computed"`
		IsDefault                       bool   `json:"is_default,omitempty" tf:"computed"`
		MaxClustersPerUser              int    `json:"max_clusters_per_user,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		policy, err := w.ClusterPolicies.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Id = policy.PolicyId
		data.Definition = policy.Definition
		data.Description = policy.Description
		data.PolicyFamilyId = policy.PolicyFamilyId
		data.PolicyFamilyDefinitionOverrides = policy.PolicyFamilyDefinitionOverrides
		data.IsDefault = policy.IsDefault
		data.MaxClustersPerUser = int(policy.MaxClustersPerUser)
		return nil
	})
}
