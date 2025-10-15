package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceExternalLocations() common.Resource {
	type externalLocationsData struct {
		common.Namespace
		Names []string `json:"names,omitempty" tf:"computed"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *externalLocationsData, w *databricks.WorkspaceClient) error {
		locations, err := w.ExternalLocations.ListAll(ctx, catalog.ListExternalLocationsRequest{})
		if err != nil {
			return err
		}
		data.Names = make([]string, 0, len(locations))
		for _, v := range locations {
			data.Names = append(data.Names, v.Name)
		}
		return nil
	})
}
