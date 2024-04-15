package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceExternalLocation() common.Resource {
	type ExternalLocationByID struct {
		Name             string                        `json:"name"`
		ExternalLocation *catalog.ExternalLocationInfo `json:"external_location_info,omitempty" tf:"computed" `
	}
	return common.WorkspaceData(func(ctx context.Context, data *ExternalLocationByID, w *databricks.WorkspaceClient) error {
		credential, err := w.ExternalLocations.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.ExternalLocation = credential
		return nil
	})
}
