package workspace

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"context"
	"log"
	"strings"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WorkspaceConfAPI exposes the workspace configurations API
type WorkspaceConfAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// NewWorkspaceConfAPI returns workspace conf API
func NewWorkspaceConfAPI(ctx context.Context, m interface{}) WorkspaceConfAPI {
	return WorkspaceConfAPI{m.(*common.DatabricksClient), ctx}
}

// Update will handle creation of new values as well as deletes. Deleting just implies that a value of "" or
// the appropriate disable string like "false" is sent with the appropriate key
func (a WorkspaceConfAPI) Update(workspaceConfMap map[string]interface{}) error {
	return a.client.Patch(a.context, "/workspace-conf", workspaceConfMap)
}

// Read just returns back a map of keys and values which keys are the configuration items and values are the settings
func (a WorkspaceConfAPI) Read(conf *map[string]interface{}) error {
	keys := []string{}
	for k := range *conf {
		keys = append(keys, k)
	}
	return a.client.Get(a.context, "/workspace-conf", map[string]string{
		"keys": strings.Join(keys, ","),
	}, &conf)
}

// ResourceWorkspaceConf maintains workspace configuration for specified keys
func ResourceWorkspaceConf() *schema.Resource {
	readContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		wsConfAPI := NewWorkspaceConfAPI(ctx, m)
		config := d.Get("custom_config").(map[string]interface{})
		log.Printf("[DEBUG] Config available in state: %v", config)
		err := wsConfAPI.Read(&config)
		if err != nil {
			return diag.FromErr(err)
		}
		log.Printf("[DEBUG] Setting new config to state: %v", config)
		// nolint
		d.Set("custom_config", config)
		return nil
	}

	updateContext := func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		wsConfAPI := NewWorkspaceConfAPI(ctx, m)
		o, n := d.GetChange("custom_config")
		old, okOld := o.(map[string]interface{})
		new, okNew := n.(map[string]interface{})
		if !okNew || !okOld {
			return diag.Errorf("Internal type casting error")
		}
		log.Printf("[DEBUG] Old worspace config: %v, new: %v", old, new)
		patch := map[string]interface{}{}
		for k, v := range new {
			patch[k] = v
		}
		for k := range old {
			_, keep := new[k]
			if keep {
				continue
			}
			log.Printf("[DEBUG] Erasing configuration of %s", k)
			if strings.HasPrefix(k, "enable") ||
				strings.HasPrefix(k, "enforce") ||
				strings.HasSuffix(k, "Enabled") {
				patch[k] = "false"
			} else {
				patch[k] = ""
			}
		}
		err := wsConfAPI.Update(patch)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId("_")
		return readContext(ctx, d, m)
	}

	return &schema.Resource{
		ReadContext:   readContext,
		CreateContext: updateContext,
		UpdateContext: updateContext,
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			config := d.Get("custom_config").(map[string]interface{})
			for k := range config {
				if strings.HasPrefix(k, "enable") ||
					strings.HasPrefix(k, "enforce") ||
					strings.HasSuffix(k, "Enabled") {
					config[k] = "false"
				} else {
					config[k] = ""
				}
			}
			wsConfAPI := NewWorkspaceConfAPI(ctx, m)
			err := wsConfAPI.Update(config)
			if err != nil {
				return diag.FromErr(err)
			}
			return nil
		},
		Schema: map[string]*schema.Schema{
			"custom_config": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
