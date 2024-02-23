package catalog

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// UC catalog resources accept an optional metastore_id parameter. This is required for account-level operations, but it is not used
// for workspace-level. However, to avoid confusion for customers when they specify an id at workspace-level, we validate that the
// id matches the metastore assigned to the workspace.
// This must be done for all operations that modify the resource (create, update, delete) but not when reading the resource to avoid
// breaking the terraform resource if the current compute value is not correct (due to this validation not exiting previously).
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

func supressCaseSensitivity(k, old, new string, d *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}
