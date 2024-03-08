package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceStorageCredential() common.Resource {
	type AccountMetastoreByID struct {
		Name              string                         `json:"name"`
		StorageCredential *catalog.StorageCredentialInfo `json:"storage_credential_info,omitempty" tf:"computed" `
	}
	return common.WorkspaceData(func(ctx context.Context, data *AccountMetastoreByID, w *databricks.WorkspaceClient) error {
		credential, err := w.StorageCredentials.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.StorageCredential = credential
		return nil
	})
}
