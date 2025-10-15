package vectorsearch

import (
	"context"
	"log"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
)

const defaultEndpointProvisionTimeout = 75 * time.Minute
const deleteCallTimeout = 10 * time.Second

type EndpointInfo struct {
	vectorsearch.EndpointInfo
	common.Namespace
}

func ResourceVectorSearchEndpoint() common.Resource {
	s := common.StructToSchema(
		EndpointInfo{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(s, "name").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(s, "endpoint_type").SetRequired().SetForceNew()
			delete(s, "id")
			delete(s, "custom_tags")
			for _, field := range []string{"creator", "creation_timestamp", "last_updated_timestamp",
				"last_updated_user", "endpoint_status", "num_indexes", "effective_budget_policy_id"} {
				common.CustomizeSchemaPath(s, field).SetReadOnly()
			}
			common.CustomizeSchemaPath(s).AddNewField("endpoint_id", &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			})
			common.CustomizeSchemaPath(s).AddNewField("budget_policy_id", &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			})
			common.NamespaceCustomizeSchemaMap(s)
			return s
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var req vectorsearch.CreateEndpoint
			common.DataToStructPointer(d, s, &req)
			wait, err := w.VectorSearchEndpoints.CreateEndpoint(ctx, req)
			if err != nil {
				return err
			}
			endpoint, err := wait.GetWithTimeout(d.Timeout(schema.TimeoutCreate) - deleteCallTimeout)
			if err != nil {
				log.Printf("[ERROR] Error waiting for endpoint to be created: %s", err.Error())
				nestedErr := w.VectorSearchEndpoints.DeleteEndpointByEndpointName(ctx, req.Name)
				if nestedErr != nil {
					log.Printf("[ERROR] Error cleaning up endpoint: %s", nestedErr.Error())
				}
				return err
			}
			d.SetId(endpoint.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			endpoint, err := w.VectorSearchEndpoints.GetEndpointByEndpointName(ctx, d.Id())
			if err != nil {
				return err
			}
			err = common.StructToData(*endpoint, s, d)
			if err != nil {
				return err
			}
			d.Set("budget_policy_id", endpoint.EffectiveBudgetPolicyId)
			d.Set("endpoint_id", endpoint.Id)
			return nil
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			if d.HasChange("budget_policy_id") {
				_, err := w.VectorSearchEndpoints.UpdateEndpointBudgetPolicy(ctx, vectorsearch.PatchEndpointBudgetPolicyRequest{
					EndpointName:   d.Id(),
					BudgetPolicyId: d.Get("budget_policy_id").(string),
				})
				if err != nil {
					return err
				}
			}
			return nil
		},

		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.VectorSearchEndpoints.DeleteEndpointByEndpointName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultEndpointProvisionTimeout),
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
	}
}
