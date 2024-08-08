package workspace

// Preview feature: https://docs.databricks.com/security/network/ip-access-list.html
// REST API: https://docs.databricks.com/dev-tools/api/latest/ip-access-list.html#operation/create-list

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/docs"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This function applies configuration defined in the resource data to the workspace.
func applyWorkspaceConf(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	o, n := d.GetChange("custom_config")
	old, okOld := o.(map[string]any)
	new, okNew := n.(map[string]any)
	if !okNew || !okOld {
		return fmt.Errorf("internal type casting error")
	}
	log.Printf("[DEBUG] Old workspace config: %v, new: %v", old, new)
	patch := settings.WorkspaceConf{}

	// Add new configuration keys
	for k, v := range new {
		patch[k] = fmt.Sprint(v)
	}

	// Remove old configuration keys, that are no longer present in the new configuration
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

func updateWorkspaceConf(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	err := applyWorkspaceConf(ctx, d, c)
	if err != nil {
		// Update methods from the Terraform SDK persist terraform configuration
		// changes to the state by default, even if update fails.
		// We revert back to the previous version of the configuration to prevent an
		// invalid workspace configuration from being persisted in the terraform state.
		prevConf, _ := d.GetChange("custom_config")
		d.Set("custom_config", prevConf)
		return err
	}
	return nil
}

func deleteWorkspaceConf(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	w, err := c.WorkspaceClient()
	if err != nil {
		return err
	}
	config := d.Get("custom_config").(map[string]any)
	// Delete keys one at a time to reset as many configuration values as possible to their original state.
	// Delete in alphabetical order by key to ensure deterministic behavior.
	keys := make([]string, 0, len(config))
	for k := range config {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := config[k]
		patch := settings.WorkspaceConf{}
		// The default value for all configurations is assumed to be "false" for boolean
		// configurations and an empty string for string configurations.
		switch r := v.(type) {
		default:
			patch[k] = ""
		case string:
			_, err := strconv.ParseBool(strings.ToLower(r))
			if err != nil {
				patch[k] = ""
			} else {
				patch[k] = "false"
			}
		case bool:
			patch[k] = "false"
		}
		err = w.WorkspaceConf.SetStatus(ctx, patch)
		// Tolerate errors like the following on deletion:
		// cannot delete workspace conf: Some values are not allowed: {"enableGp3":"","enableIpAccessLists":""}
		// The API for workspace conf is quite limited and doesn't support a generic "reset to original state"
		// operation. So if our attempted reset fails, don't fail resource deletion.
		var apiErr *apierr.APIError
		if errors.As(err, &apiErr) && strings.HasPrefix(apiErr.Message, "Some values are not allowed") {
			tflog.Warn(ctx, fmt.Sprintf("Encountered error while deleting databricks_workspace_conf: %s. The workspace configuration may not be fully restored to its original state. For more information, see %s", apiErr.Message, docs.ResourceDocumentationUrl("workspace_conf")))
		} else if err != nil {
			return err
		}
	}
	return nil
}

// ResourceWorkspaceConf maintains workspace configuration for specified keys
func ResourceWorkspaceConf() common.Resource {
	return common.Resource{
		Create: applyWorkspaceConf,
		Update: updateWorkspaceConf,
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
		Delete: deleteWorkspaceConf,
		Schema: map[string]*schema.Schema{
			"custom_config": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}
