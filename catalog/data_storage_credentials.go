package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceStorageCredentials() common.Resource {
	type storageCredentialsData struct {
		Names []string `json:"names,omitempty" tf:"computed"`
	}
	return common.WorkspaceDataWithUnifiedProvider(func(ctx context.Context, data *storageCredentialsData, w *databricks.WorkspaceClient) error {
		credentials, err := w.StorageCredentials.ListAll(ctx, catalog.ListStorageCredentialsRequest{})
		if err != nil {
			return err
		}
		data.Names = []string{}
		for _, v := range credentials {
			data.Names = append(data.Names, v.Name)
		}
		return nil
	})
}
