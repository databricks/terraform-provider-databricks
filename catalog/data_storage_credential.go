package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceStorageCredential() common.Resource {
	type AccountMetastoreByID struct {
		common.Namespace
		Id                string                         `json:"id,omitempty" tf:"computed"`
		Name              string                         `json:"name"`
		StorageCredential *catalog.StorageCredentialInfo `json:"storage_credential_info,omitempty" tf:"computed" `
	}
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *AccountMetastoreByID, w *databricks.WorkspaceClient) error {
		credential, err := w.StorageCredentials.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.StorageCredential = credential
		data.Id = credential.Id
		return nil
	})
}
