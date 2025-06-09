package bindings

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type BindingsSecurableType string

const BindingsSecurableTypeCatalog BindingsSecurableType = `catalog`
const BindingsSecurableTypeCredential BindingsSecurableType = `credential`
const BindingsSecurableTypeExternalLocation BindingsSecurableType = `external_location`
const BindingsSecurableTypeStorageCredential BindingsSecurableType = `storage_credential`

func (f *BindingsSecurableType) String() string {
	return string(*f)
}

func (f *BindingsSecurableType) Set(v string) error {
	switch v {
	case `catalog`, `credential`, `external_location`, `storage_credential`:
		*f = BindingsSecurableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "catalog", "credential", "external_location", "storage_credential"`, v)
	}
}

func (f *BindingsSecurableType) Type() string {
	return "BindingsSecurableType"
}

func AddCurrentWorkspaceBindings(ctx context.Context, d *schema.ResourceData, w *databricks.WorkspaceClient, securableName string, securableType BindingsSecurableType) error {
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
		SecurableType: securableType.String(),
		Add: []catalog.WorkspaceBinding{
			{
				BindingType: catalog.WorkspaceBindingBindingTypeBindingTypeReadWrite,
				WorkspaceId: currentMetastoreAssignment.WorkspaceId,
			},
		},
	})
	return err
}
