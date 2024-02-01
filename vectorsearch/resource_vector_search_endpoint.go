package vectorsearch

import (
	"context"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
)

const DefaultProvisionTimeout = 45 * time.Minute

func ResourceVectorSearchEndpoint() *schema.Resource {
	s := common.StructToSchema(
		vectorsearch.EndpointInfo{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(s, "name").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(s, "endpoint_type").SetRequired().SetForceNew()
			delete(s, "id")
			common.CustomizeSchemaPath(s, "creator").SetReadOnly()
			common.CustomizeSchemaPath(s, "creation_timestamp").SetReadOnly()
			common.CustomizeSchemaPath(s, "last_updated_timestamp").SetReadOnly()
			common.CustomizeSchemaPath(s, "last_updated_user").SetReadOnly()
			common.CustomizeSchemaPath(s, "endpoint_status").SetReadOnly()
			common.CustomizeSchemaPath(s, "num_indexes").SetReadOnly()
			common.CustomizeSchemaPath(s).AddNewField("endpoint_id", &schema.Schema{
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
			endpoint, err := w.VectorSearchEndpoints.CreateEndpointAndWait(ctx, req,
				retries.Timeout[vectorsearch.EndpointInfo](d.Timeout(schema.TimeoutCreate)))
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
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}.ToResource()
}
