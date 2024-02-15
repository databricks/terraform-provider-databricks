package vectorsearch

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
)

func ResourceVectorSearchIndex() common.Resource {
	s := common.StructToSchema(
		vectorsearch.CreateVectorIndexRequest{},
		func(s map[string]*schema.Schema) map[string]*schema.Schema {
			common.CustomizeSchemaPath(s, "name").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(s, "primary_key").SetRequired().SetForceNew()
			common.CustomizeSchemaPath(s, "index_type").SetRequired().SetForceNew()

			common.CustomizeSchemaPath(s, "status").SetReadOnly()
			common.CustomizeSchemaPath(s, "creator").SetReadOnly()

			common.CustomizeSchemaPath(s, "delta_sync_index_spec", "pipeline_id").SetReadOnly()
			// common.MustSchemaPath(s, "delta_sync_vector_index_spec", "embedding_vector_columns").MinItems = 1

			s["delta_sync_index_spec"].ExactlyOneOf = []string{"delta_sync_index_spec", "direct_access_index_spec"}
			s["direct_access_index_spec"].ExactlyOneOf = []string{"delta_sync_index_spec", "direct_access_index_spec"}

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
			index, err := w.VectorSearchIndexes.CreateIndex(ctx, req)
			if err != nil {
				return err
			}
			if index.VectorIndex == nil {
				return fmt.Errorf("vector index information is nil")
			}
			d.SetId(index.VectorIndex.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			index, err := w.VectorSearchIndexes.GetIndex(ctx, vectorsearch.GetIndexRequest{
				IndexName: d.Id(),
			})
			if err != nil {
				return err
			}
			err = common.StructToData(*index, s, d)
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
