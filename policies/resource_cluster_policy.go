package policies

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func isBuiltinPolicyFamily(ctx context.Context, w *databricks.WorkspaceClient, familyId, familyName string) (bool, error) {
	// Fetch supported policy families, and check against it
	families, err2 := w.PolicyFamilies.ListAll(ctx, compute.ListPolicyFamiliesRequest{})
	if err2 != nil {
		return false, err2
	}
	for _, family := range families {
		if familyId == family.PolicyFamilyId && familyName == family.Name {
			return true, nil
		}
	}
	return false, nil
}

// ResourceClusterPolicy ...
func ResourceClusterPolicy() common.Resource {
	s := common.StructToSchema(
		compute.CreatePolicy{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["policy_id"] = &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			}
			m["definition"].ConflictsWith = []string{"policy_family_definition_overrides", "policy_family_id"}
			m["definition"].Computed = true
			m["definition"].DiffSuppressFunc = common.SuppressDiffWhitespaceChange

			m["policy_family_definition_overrides"].ConflictsWith = []string{"definition"}
			m["policy_family_definition_overrides"].DiffSuppressFunc = common.SuppressDiffWhitespaceChange
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

			var request compute.CreatePolicy
			common.DataToStructPointer(d, s, &request)

			var clusterPolicy *compute.CreatePolicyResponse
			if request.PolicyFamilyId != "" {
				isBuiltin, err2 := isBuiltinPolicyFamily(ctx, w, request.PolicyFamilyId, request.Name)
				if err2 != nil {
					return err2
				}
				if isBuiltin {
					resp, err2 := w.ClusterPolicies.GetByName(ctx, request.Name)
					if err2 != nil {
						return err2
					}
					clusterPolicy = &compute.CreatePolicyResponse{PolicyId: resp.PolicyId}
					var editRequest compute.EditPolicy
					common.DataToStructPointer(d, s, &editRequest)
					editRequest.PolicyId = resp.PolicyId
					err = w.ClusterPolicies.Edit(ctx, editRequest)
				} else {
					clusterPolicy, err = w.ClusterPolicies.Create(ctx, request)
				}
			} else {
				clusterPolicy, err = w.ClusterPolicies.Create(ctx, request)
			}
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

			var request compute.EditPolicy
			common.DataToStructPointer(d, s, &request)
			request.PolicyId = d.Id()
			if request.PolicyFamilyId != "" {
				request.Definition = ""
			}

			return w.ClusterPolicies.Edit(ctx, request)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var request compute.EditPolicy
			common.DataToStructPointer(d, s, &request)
			if request.PolicyFamilyId != "" {
				isBuiltin, err := isBuiltinPolicyFamily(ctx, w, request.PolicyFamilyId, request.Name)
				if err != nil {
					return err
				}
				if isBuiltin {
					request.PolicyId = d.Id()
					request.PolicyFamilyDefinitionOverrides = ""
					request.Definition = ""
					return w.ClusterPolicies.Edit(ctx, request)
				}
			}
			return w.ClusterPolicies.DeleteByPolicyId(ctx, d.Id())
		},
	}
}
