package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
)

func validateMetastoreId(ctx context.Context, w *databricks.WorkspaceClient, metastoreId string) error {
	if metastoreId == "" {
		return nil
	}
	cat, err := w.Metastores.Current(ctx)
	if err != nil {
		return err
	}
	if cat.MetastoreId != metastoreId {
		return fmt.Errorf("metastore_id must be empty or equal to the metastore id assigned to the workspace: %s. "+
			"If the metastore assigned to the workspace has changed, the new metastore id must be explicitly set", cat.MetastoreId)
	}
	return nil
}
