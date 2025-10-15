package workspace

import (
	"context"
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func directoryPathSuppressDiff(k, old, new string, d *schema.ResourceData) bool {
	return new == (old + "/")
}

// ResourceDirectory manages directories
func ResourceDirectory() common.Resource {
	s := map[string]*schema.Schema{
		"path": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			DiffSuppressFunc: directoryPathSuppressDiff,
		},
		"object_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"delete_recursive": {
			Type:     schema.TypeBool,
			Default:  false,
			Optional: true,
		},
		"workspace_path": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	common.NamespaceCustomizeSchemaMap(s)

	directoryRead := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		client, err := c.WorkspaceClientUnifiedProvider(ctx, d)
		if err != nil {
			return err
		}
		objectStatus, err := common.RetryOnTimeout(ctx, func(ctx context.Context) (*workspace.ObjectInfo, error) {
			return client.Workspace.GetStatusByPath(ctx, d.Id())
		})
		if err != nil {
			return err
		}
		if objectStatus.ObjectType != workspace.ObjectTypeDirectory {
			d.SetId("")
			return fmt.Errorf("different object type, %s, on this path other than a directory", objectStatus.ObjectType)
		}
		err = common.StructToData(objectStatus, s, d)
		if err != nil {
			return err
		}
		SetWorkspaceObjectComputedProperties(d, c)
		return nil
	}

	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			path := d.Get("path").(string)
			err = client.Workspace.MkdirsByPath(ctx, path)
			if err != nil {
				// TODO: handle RESOURCE_ALREADY_EXISTS
				return err
			}
			d.SetId(path)
			return nil
		},
		Read:   directoryRead,
		Update: directoryRead,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			client, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			err = client.Workspace.Delete(ctx, workspace.Delete{
				Path:      d.Id(),
				Recursive: d.Get("delete_recursive").(bool),
			})
			if err != nil && apierr.IsMissing(err) {
				log.Printf("[INFO] Error deleting directory: %v", err)
				err = nil
			}
			return err
		},
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
	}
}
