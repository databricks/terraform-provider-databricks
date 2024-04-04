package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
func (volumeDataParams) CustomizeSchema(s map[string]*schema.Schema) map[string]*schema.Schema {
	common.CustomizeSchemaPath(s, "full_name").SetExactlyOneOf([]string{"catalog_name"}).SetComputed()
	common.CustomizeSchemaPath(s, "catalog_name").SetRequiredWith([]string{"schema_name", "name"}).SetComputed()
	common.CustomizeSchemaPath(s, "schema_name").SetRequiredWith([]string{"catalog_name", "name"}).SetComputed()
	common.CustomizeSchemaPath(s, "name").SetRequiredWith([]string{"catalog_name", "schema_name"}).SetComputed()
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
