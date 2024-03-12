package vectorsearch

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
)

const defaultIndexProvisionTimeout = 5 * time.Minute

func waitForSearchIndex(w *databricks.WorkspaceClient, ctx context.Context, searchIndexName string) error {
	return retry.RetryContext(ctx, defaultIndexProvisionTimeout-deleteCallTimeout, func() *retry.RetryError {
		index, err := w.VectorSearchIndexes.GetIndexByIndexName(ctx, searchIndexName)
		if err != nil {
			return retry.NonRetryableError(err)
		}
		if index.Status.Ready { // We really need to depend on the detailed status of the index, but it's not available in the API yet
			return nil
		}
		return retry.RetryableError(fmt.Errorf("vector search index %s is still pending", searchIndexName))
	})
}

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
			err = waitForSearchIndex(w, ctx, req.Name)
			if err != nil {
				nestedErr := w.VectorSearchIndexes.DeleteIndexByIndexName(ctx, req.Name)
				if nestedErr != nil {
					log.Printf("[ERROR] Error cleaning up search index: %s", nestedErr.Error())
				}
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
			Create: schema.DefaultTimeout(defaultIndexProvisionTimeout),
		},
	}
}
