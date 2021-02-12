package workspace

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	defaultPosition = 10000
	maxScriptSize   = 64 * 1024
)

// ResourceGlobalInitScript manages notebooks
func ResourceGlobalInitScript() *schema.Resource {
	// TODO: move this into a common piece, in the file_resource, and merge with "path" entry
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
				validation.StringMatch(regexp.MustCompile("^[-a-zA-Z0-9_.]{1,100}$"), "Name should match regex!")),
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
				return fmt.Errorf("Size of the global init script (%d bytes) exceeds maximal allowed (%d bytes)",
					contentLen, maxScriptSize)
			}
			globalInitScriptsAPI := NewGlobalInitScriptsAPI(ctx, c)
			scriptID, err := globalInitScriptsAPI.Create(GlobalInitScriptPayload{
				ContentBase64: base64.StdEncoding.EncodeToString(content),
				Enabled:       d.Get("enabled").(bool),
				Position:      int32(d.Get("position").(int)),
				Name:          d.Get("name").(string),
			})
			if err != nil {
				return err
			}
			d.SetId(scriptID)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			globalInitScriptsAPI := NewGlobalInitScriptsAPI(ctx, c)
			scriptStatus, err := globalInitScriptsAPI.Get(d.Id())
			if err != nil {
				return err
			}
			return internal.StructToData(scriptStatus, s, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			content, err := ReadContent(d)
			if err != nil {
				return err
			}
			if contentLen := len(content); contentLen > maxScriptSize {
				return fmt.Errorf("Size of the global init script (%d bytes) exceeds maximal allowed (%d bytes)",
					contentLen, maxScriptSize)
			}
			globalInitScriptsAPI := NewGlobalInitScriptsAPI(ctx, c)
			return globalInitScriptsAPI.Update(d.Id(), GlobalInitScriptPayload{
				ContentBase64: base64.StdEncoding.EncodeToString(content),
				Enabled:       d.Get("enabled").(bool),
				Position:      int32(d.Get("position").(int)),
				Name:          d.Get("name").(string),
			})
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewGlobalInitScriptsAPI(ctx, c).Delete(d.Id())
		},
		StateUpgraders: []schema.StateUpgrader{},
		Schema:         s,
		SchemaVersion:  1,
		Timeouts:       &schema.ResourceTimeout{},
	}.ToResource()
}
