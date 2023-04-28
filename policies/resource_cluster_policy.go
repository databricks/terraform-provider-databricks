package policies

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/clusterpolicies"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceClusterPolicy ...
func ResourceClusterPolicy() *schema.Resource {
	s := common.StructToSchema(
		clusterpolicies.CreatePolicy{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["policy_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			}
			m["definition"].ConflictsWith = []string{"policy_family_definition_overrides", "policy_family_id"}
			m["policy_family_definition_overrides"].ConflictsWith = []string{"definition"}
			m["policy_family_id"].ConflictsWith = []string{"definition"}
			m["policy_family_definition_overrides"].RequiredWith = []string{"policy_family_id"}

			return m
		})

	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var request clusterpolicies.CreatePolicy
			common.DataToStructPointer(d, s, &request)

			clusterPolicy, err := w.ClusterPolicies.Create(ctx, request)
			if err != nil {
				return err
			}

			d.SetId(clusterPolicy.PolicyId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			resp, err := w.ClusterPolicies.GetByPolicyId(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(resp, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}

			var request clusterpolicies.EditPolicy
			common.DataToStructPointer(d, s, &request)
			request.PolicyId = d.Id()

			return w.ClusterPolicies.Edit(ctx, request)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.ClusterPolicies.DeleteByPolicyId(ctx, d.Id())
		},
	}.ToResource()
}
