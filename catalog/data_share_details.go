package catalog

import (
	"context"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceShareDetails() *schema.Resource {
	type sharesData struct {
		Name      string     `json:"name"`
		ShareInfo *ShareInfo `json:"share_info,omitempty" tf:"computed"`
	}
	return common.DataResource(sharesData{}, func(ctx context.Context, e interface{}, c *common.DatabricksClient) error {
		data := e.(*sharesData)
		sharesAPI := NewSharesAPI(ctx, c)
		share, err := sharesAPI.getShare(data.Name)
		if err != nil {
			return err
		}
		data.ShareInfo = &share
		return nil
	})
}
