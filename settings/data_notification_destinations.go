package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceNotificationDestinations() common.Resource {
	type notificationDestinationsData struct {
		Ids []string `json:"ids,omitempty" tf:"computed,slice_set"`
	}

	return common.WorkspaceData(func(ctx context.Context, data *notificationDestinationsData, w *databricks.WorkspaceClient) error {
		return nil
	})
}
