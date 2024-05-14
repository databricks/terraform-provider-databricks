package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

type volumeDataParams struct {
	FullName    string `json:"full_name,omitempty"`
	CatalogName string `json:"catalog_name,omitempty"`
	SchemaName  string `json:"schema_name,omitempty"`
	Name        string `json:"name,omitempty"`
}

func (volumeDataParams) Aliases() map[string]string {
	return map[string]string{"Id": "FullName"}
}
func (volumeDataParams) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.SchemaPath("full_name").SetExactlyOneOf([]string{"catalog_name"}).SetComputed()
	s.SchemaPath("catalog_name").SetRequiredWith([]string{"schema_name", "name"}).SetComputed()
	s.SchemaPath("schema_name").SetRequiredWith([]string{"catalog_name", "name"}).SetComputed()
	s.SchemaPath("name").SetRequiredWith([]string{"catalog_name", "schema_name"}).SetComputed()
	return s
}

func volumeDataRead(ctx context.Context, data volumeDataParams, w *databricks.WorkspaceClient) (*catalog.VolumeInfo, error) {
	volumeRequest := catalog.ReadVolumeRequest{}
	if data.FullName != "" {
		volumeRequest.Name = data.FullName
	} else {
		volumeRequest.Name = fmt.Sprintf("%s.%s.%s", data.CatalogName, data.SchemaName, data.Name)
	}

	volume, err := w.Volumes.Read(ctx, volumeRequest)
	if err != nil {
		return nil, err
	}

	return volume, nil
}

func DataSourceVolume() common.Resource {
	return common.WorkspaceDataWithParams(volumeDataRead)
}
