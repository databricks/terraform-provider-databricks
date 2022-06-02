package catalog

import (
	"context"
	"path"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func updateFunctionFactory(pathPrefix string, updatable []string) func(context.Context, *schema.ResourceData, *common.DatabricksClient) error {
	return func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
		patch := map[string]interface{}{}
		for _, field := range updatable {
			old, new := d.GetChange(field)
			if !d.HasChange(field) {
				continue
			}
			if field == "delta_sharing_enabled" && old != new && new == true &&
				!d.HasChange("delta_sharing_recipient_token_lifetime_in_seconds") {
				patch["delta_sharing_recipient_token_lifetime_in_seconds"] =
					d.Get("delta_sharing_recipient_token_lifetime_in_seconds")
			}
			patch[field] = new
		}
		if len(patch) == 0 {
			return nil
		}

		return c.Patch(ctx, path.Join(pathPrefix, d.Id()), patch)
	}
}
