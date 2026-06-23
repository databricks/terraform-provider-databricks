package catalog

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceConnection() common.Resource {
	type ConnectionByName struct {
		common.Namespace
		Id         string                  `json:"id,omitempty" tf:"computed"`
		Name       string                  `json:"name"`
		Connection *catalog.ConnectionInfo `json:"connection_info,omitempty" tf:"computed"`
	}
	return common.WorkspaceData(func(ctx context.Context, data *ConnectionByName, w *databricks.WorkspaceClient) error {
		conn, err := w.Connections.GetByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Connection = conn
		data.Id = conn.Name
		return nil
	})
}
