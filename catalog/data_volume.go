package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceVolume() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *struct {
		Id     string              `json:"id,omitempty" tf:"computed"`
		Name   string              `json:"name"`
		Volume *catalog.VolumeInfo `json:"volume_info,omitempty" tf:"computed"`
	}, w *databricks.WorkspaceClient) error {
		fmt.Println(data)
		volume, err := w.Volumes.ReadByName(ctx, data.Name)
		if err != nil {
			return err
		}
		data.Volume = volume
		data.Id = volume.VolumeId
		return nil
	})
}
