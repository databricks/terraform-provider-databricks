package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceShare() common.Resource {
	type ShareDetail struct {
		Name      string                     `json:"name,omitempty" tf:"computed"`
		Objects   []sharing.SharedDataObject `json:"objects,omitempty" tf:"computed,slice_set,alias:object"`
		CreatedAt int64                      `json:"created_at,omitempty" tf:"computed"`
		CreatedBy string                     `json:"created_by,omitempty" tf:"computed"`
	}
	return common.DataResource(ShareDetail{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*ShareDetail)
		client, err := c.WorkspaceClient()
		if err != nil {
			return err
		}

		share, err := client.Shares.Get(ctx, sharing.GetShareRequest{
			Name:              data.Name,
			IncludeSharedData: true,
		})
		if err != nil {
			return err
		}
		data.Objects = share.Objects
		data.CreatedAt = share.CreatedAt
		data.CreatedBy = share.CreatedBy
		return nil
	})
}
