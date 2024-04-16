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

func ResourceVectorSearchEndpoint() common.Resource {
	s := common.StructToSchema(
		vectorsearch.EndpointInfo{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			emptyCtx := common.SchemaPathContext{}
			common.CustomizeSchemaPath(emptyCtx, s, "name").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(emptyCtx, s, "endpoint_type").SetRequired().SetForceNew()
			delete(s, "id")
			common.CustomizeSchemaPath(emptyCtx, s, "creator").SetReadOnly()
			common.CustomizeSchemaPath(emptyCtx, s, "creation_timestamp").SetReadOnly()
			common.CustomizeSchemaPath(emptyCtx, s, "last_updated_timestamp").SetReadOnly()
			common.CustomizeSchemaPath(emptyCtx, s, "last_updated_user").SetReadOnly()
			common.CustomizeSchemaPath(emptyCtx, s, "endpoint_status").SetReadOnly()
			common.CustomizeSchemaPath(emptyCtx, s, "num_indexes").SetReadOnly()
			common.CustomizeSchemaPath(emptyCtx, s).AddNewField("endpoint_id", &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			})

			return s
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
			w, err := c.WorkspaceClient()
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
			d.Set("endpoint_id", endpoint.Id)
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
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
	}
}
