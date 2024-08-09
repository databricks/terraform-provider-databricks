package bindings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AddCurrentWorkspaceBindings(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient, securableName string, securableType catalog.UpdateBindingsSecurableType) error {
	if d.Get("isolation_mode") != "ISOLATED" && d.Get("isolation_mode") != "ISOLATION_MODE_ISOLATED" {
		return nil
	}
	// Bind the current workspace if the catalog is isolated, otherwise the read will fail
	currentMetastoreAssignment, err := w.Metastores.Current(ctx)
	if err != nil {
		return err
	}
	_, err = w.WorkspaceBindings.UpdateBindings(ctx, catalog.UpdateWorkspaceBindingsParameters{
		SecurableName: securableName,
		SecurableType: securableType,
		Add: []catalog.WorkspaceBinding{
			{
				BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
				WorkspaceId: currentMetastoreAssignment.WorkspaceId,
			},
		},
	})
	return err
}
