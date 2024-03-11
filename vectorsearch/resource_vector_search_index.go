package vectorsearch

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
)

func ResourceVectorSearchIndex() common.Resource {
	s := common.StructToSchema(
		vectorsearch.CreateVectorIndexRequest{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			common.MustSchemaPath(s, "delta_sync_index_spec", "embedding_vector_columns").MinItems = 1
			exof := []string{"delta_sync_index_spec", "direct_access_index_spec"}
			s["delta_sync_index_spec"].ExactlyOneOf = exof
			s["direct_access_index_spec"].ExactlyOneOf = exof

			return s
		})

	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var req vectorsearch.CreateVectorIndexRequest
			common.DataToStructPointer(d, s, &req)
			_, err = w.VectorSearchIndexes.CreateIndex(ctx, req)
			if err != nil {
				return err
			}
			d.SetId(req.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			index, err := w.VectorSearchIndexes.GetIndexByIndexName(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(*index, s, d)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.VectorSearchIndexes.DeleteIndexByIndexName(ctx, d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  0,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(DefaultProvisionTimeout),
		},
	}
}
