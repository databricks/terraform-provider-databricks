package workspace

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	defaultPosition = 10000
	maxScriptSize   = 64 * 1024
)

// ResourceGlobalInitScript manages notebooks
func ResourceGlobalInitScript() common.Resource {
	extra := map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ValidateDiagFunc: validation.ToDiagFunc(
				validation.StringMatch(regexp.MustCompile("^[-a-zA-Z0-9_. ]{1,100}$"), "Name should match regex!")),
		},
		"position": {
			Type:             schema.TypeInt,
			Optional:         true,
			Computed:         true,
			ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(0)),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				defaultPosString := fmt.Sprintf("%d", defaultPosition)
				return (old == new) || (old != "" && new == defaultPosString && old != defaultPosString)
			},
		},
	}
	s := FileContentSchemaWithoutPath(extra)
	return common.Resource{
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			if contentLen := len(content); contentLen > maxScriptSize {
				return fmt.Errorf("size of the global init script (%d bytes) exceeds maximal allowed (%d bytes)",
					contentLen, maxScriptSize)
			}
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			created, err := w.GlobalInitScripts.Create(ctx, compute.GlobalInitScriptCreateRequest{
				Script:   base64.StdEncoding.EncodeToString(content),
				Enabled:  d.Get("enabled").(bool),
				Position: d.Get("position").(int),
				Name:     d.Get("name").(string),
			})
			if err != nil {
				return err
			}
			d.SetId(created.ScriptId)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			scriptStatus, err := w.GlobalInitScripts.GetByScriptId(ctx, d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(scriptStatus, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			if contentLen := len(content); contentLen > maxScriptSize {
				return fmt.Errorf("size of the global init script (%d bytes) exceeds maximal allowed (%d bytes)",
					contentLen, maxScriptSize)
			}
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.GlobalInitScripts.Update(ctx, compute.GlobalInitScriptUpdateRequest{
				ScriptId: d.Id(),
				Script:   base64.StdEncoding.EncodeToString(content),
				Enabled:  d.Get("enabled").(bool),
				Position: d.Get("position").(int),
				Name:     d.Get("name").(string),
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClientUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			return w.GlobalInitScripts.DeleteByScriptId(ctx, d.Id())
		},
		Schema:        s,
		SchemaVersion: 1,
		Timeouts:      &schema.ResourceTimeout{},
	}
}
