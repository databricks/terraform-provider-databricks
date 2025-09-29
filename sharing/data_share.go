package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
)

type ShareDetail struct {
	Name      string                     `json:"name,omitempty" tf:"computed"`
	Objects   []sharing.SharedDataObject `json:"objects,omitempty" tf:"computed,slice_set,alias:object"`
	CreatedAt int64                      `json:"created_at,omitempty" tf:"computed"`
	CreatedBy string                     `json:"created_by,omitempty" tf:"computed"`
	common.ProviderConfig
}

func (ShareDetail) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	s.SchemaPath("name").SetComputed()
	s.SchemaPath("object", "added_at").SetComputed()
	s.SchemaPath("object", "added_by").SetComputed()
	s.SchemaPath("object", "data_object_type").SetRequired()
	s.SchemaPath("object", "status").SetComputed()
	s.SchemaPath("object", "partition", "value", "op").SetRequired()
	s.SchemaPath("object", "partition", "value", "name").SetRequired()
	s.SchemaPath("object", "partition", "value").SetMinItems(1)

	s.SchemaPath("provider_config").SetOptional()
	s.SchemaPath("provider_config", "workspace_id").SetRequired()

	return s
}

func (ShareDetail) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"sharing.SharedDataObject": {
			"partitions": "partition",
		},
		"sharing.Partition": {
			"values": "value",
		},
	}
}

func DataSourceShare() common.Resource {
	return common.WorkspaceData(func(ctx context.Context, data *ShareDetail, client *databricks.WorkspaceClient) error {
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
