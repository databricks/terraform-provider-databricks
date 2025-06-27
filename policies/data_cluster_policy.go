package policies

import (
	"context"
	"log"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataSourceClusterPolicy returns information about cluster policy specified by name
func DataSourceClusterPolicy() common.Resource {
	resource := common.WorkspaceData(func(ctx context.Context, data *struct {
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
	resource.SchemaVersion = 1
	resource.StateUpgraders = []schema.StateUpgrader{
		{
			Type:    resource.ToResource().CoreConfigSchema().ImpliedType(),
			Version: 0,
			Upgrade: removeZeroMaxClustersPerUser,
		},
	}
	return resource
}

func removeZeroMaxClustersPerUser(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	newState := map[string]any{}
	for k, v := range rawState {
		switch k {
		case "max_clusters_per_user":
			vv, ok := v.(int)
			if !ok || vv == 0 {
				log.Printf("[INFO] remove zero max_clusters_per_user")
				continue
			}
			newState[k] = v
		default:
			newState[k] = v
		}
	}
	return newState, nil
}
