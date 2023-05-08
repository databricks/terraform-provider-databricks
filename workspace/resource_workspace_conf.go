package workspace

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceWorkspaceConf maintains workspace configuration for specified keys
func ResourceWorkspaceConf() *schema.Resource {
	create := func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		o, n := d.GetChange("custom_config")
		old, okOld := o.(map[string]any)
		new, okNew := n.(map[string]any)
		if !okNew || !okOld {
			return fmt.Errorf("internal type casting error")
		}
		log.Printf("[DEBUG] Old workspace config: %v, new: %v", old, new)
		patch := settings.WorkspaceConf{}
		for k, v := range new {
			patch[k] = fmt.Sprint(v)
		}
		for k, v := range old {
			_, keep := new[k]
			if keep {
				continue
			}
			log.Printf("[DEBUG] Erasing configuration of %s", k)
			switch r := v.(type) {
			default:
				patch[k] = ""
			case string:
				_, err := strconv.ParseBool(r)
				if err != nil {
					patch[k] = ""
				} else {
					patch[k] = "false"
				}
			case bool:
				patch[k] = "false"
			}
		}
		w, err := c.WorkspaceClient()
		if err != nil {
			return err
		}
		err = w.WorkspaceConf.SetStatus(ctx, patch)
		if err != nil {
			return err
		}
		newConfig := map[string]any{}
		for k, v := range patch {
			newConfig[k] = v
		}
		d.SetId("_")
		return nil
	}
	return common.Resource{
		Create: create,
		Update: create,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			config := d.Get("custom_config").(map[string]any)
			log.Printf("[DEBUG] Config available in state: %v", config)
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var keys []string
			for k := range config {
				keys = append(keys, k)
			}
			if len(keys) == 0 {
				return nil
			}
			remote, err := w.WorkspaceConf.GetStatus(ctx, settings.GetStatusRequest{
				Keys: strings.Join(keys, ","),
			})
			if err != nil {
				return err
			}
			for k, v := range *remote {
				config[k] = v
			}
			log.Printf("[DEBUG] Setting new config to state: %v", config)
			return d.Set("custom_config", config)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			patch := settings.WorkspaceConf{}
			config := d.Get("custom_config").(map[string]any)
			for k, v := range config {
				switch r := v.(type) {
				default:
					patch[k] = ""
				case string:
					_, err := strconv.ParseBool(r)
					if err != nil {
						patch[k] = ""
					} else {
						patch[k] = "false"
					}
				case bool:
					patch[k] = "false"
				}
			}
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.WorkspaceConf.SetStatus(ctx, patch)
		},
		Schema: map[string]*schema.Schema{
			"custom_config": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}.ToResource()
}
