package workspace

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceDirectory ...
func DataSourceDirectory() common.Resource {
	type Directory struct {
		Id            string `json:"id,omitempty" tf:"computed"`
		Path          string `json:"path"`
		ObjectId      int64  `json:"object_id,omitempty" tf:"computed"`
		WorkspacePath string `json:"workspace_path,omitempty" tf:"computed"`
	}
	return common.WorkspaceData(func(ctx context.Context, d *Directory, client *databricks.WorkspaceClient) error {
		data, err := robustGetStatus(ctx, client, d.Path)
		if err != nil {
			return err
		}
		if data.ObjectType != workspace.ObjectTypeDirectory {
			return fmt.Errorf("'%s' isn't a directory", d.Path)
		}
		d.Id = data.Path
		d.ObjectId = data.ObjectId
		d.WorkspacePath = "/Workspace" + data.Path
		return nil
	})
}
