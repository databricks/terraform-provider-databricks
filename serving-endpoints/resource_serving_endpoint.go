package serving_endpoints

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/endpoints"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultProvisionTimeout = 45 * time.Minute

func ResourceServingEndpoint() *schema.Resource {
	s := common.StructToSchema(
		endpoints.CreateServingEndpoint{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var e endpoints.CreateServingEndpoint
			common.DataToStructPointer(d, s, &e)
			endpoint, err := w.ServingEndpoints.CreateAndWait(ctx, e, retries.Timeout[endpoints.ServingEndpointDetailed](DefaultProvisionTimeout))
			if err != nil {
				return err
			}
			d.SetId(endpoint.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			endpoint, err := w.ServingEndpoints.GetByName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*endpoint, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var e endpoints.CreateServingEndpoint
			common.DataToStructPointer(d, s, &e)
			e.Config.Name = e.Name
			_, err = w.ServingEndpoints.UpdateConfigAndWait(ctx, e.Config, retries.Timeout[endpoints.ServingEndpointDetailed](DefaultProvisionTimeout))
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.ServingEndpoints.DeleteByName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts:       &schema.ResourceTimeout{},
	}.ToResource()
}
