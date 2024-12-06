package workspace

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// DataSourceNotebook ...
func DataSourceNotebook() common.Resource {
	s := map[string]*schema.Schema{
		"path": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"format": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			ValidateFunc: validation.StringInSlice([]string{
				"DBC",
				"SOURCE",
				"HTML",
			}, false),
		},
		"content": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"language": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"object_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"object_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"workspace_path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, m *common.DatabricksClient) error {
			w, err := m.WorkspaceClient()
			if err != nil {
				return err
			}
			notebooksAPI := NewNotebooksAPI(ctx, m)
			path := d.Get("path").(string)
			format := d.Get("format").(string)
			notebookContent, err := notebooksAPI.Export(path, format)
			if err != nil {
				return err
			}
			d.SetId(path)
			// nolint
			d.Set("content", notebookContent)
			objectStatus, err := common.RetryOnTimeout(ctx, func(ctx context.Context) (*workspace.ObjectInfo, error) {
				return w.Workspace.GetStatusByPath(ctx, d.Id())
			})
			if err != nil {
				return err
			}
			err = common.StructToData(objectStatus, s, d)
			if err != nil {
				return err
			}
			d.Set("workspace_path", "/Workspace"+objectStatus.Path)
			return nil
		},
	}
}
