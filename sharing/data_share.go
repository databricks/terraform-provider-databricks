package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
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

	return common.WorkspaceDataWithParams(func(ctx context.Context, data ShareInfo, c *databricks.WorkspaceClient) (*ShareDetail, error) {
		var shareInfo *ShareDetail = &ShareDetail{}
		share, err := c.Shares.Get(ctx, sharing.GetShareRequest{
			Name:              data.Name,
			IncludeSharedData: true,
		})
		if err != nil {
			return nil, err
		}
		shareInfo.Objects = share.Objects
		shareInfo.CreatedAt = share.CreatedAt
		shareInfo.CreatedBy = share.CreatedBy
		return shareInfo, nil
	})
}
