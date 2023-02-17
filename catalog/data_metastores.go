package catalog

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMetastores() *schema.Resource {
	return common.DataResource(MetastoreInfo{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*MetastoreInfo)
		metastoresAPI := NewMetastoresAPI(ctx, c)
		metastores, err := metastoresAPI.listMetastores()
		if err != nil {
			return err
		}
		for _, v := range metastores.Metastores {
			data.Name = append(data.Name, v.name)
			data.StorageRoot = append(data.StorageRoot, v.storage_root)
			data.DefaultDacID = append(data.DefaultDacID, v.default_data_access_config_id)
			data.Owner = append(data.Owner, v.owner)
			data.MetastoreID = append(data.MetastoreID, v.metastore_id)
			data.Region = append(data.Region, v.region)
			data.Cloud = append(data.Cloud, v.cloud)
			data.GlobalMetastoreId = append(data.GlobalMetastoreId, v.global_metastore_id)
			data.CreatedAt = append(data.CreatedAt, v.created_at)
			data.CreatedBy = append(data.CreatedBy, v.created_by)
			data.UpdatedAt = append(data.UpdatedAt, v.updated_at)
			data.UpdatedBy = append(data.UpdatedBy, v.updated_by)
			data.DeltaSharingScope = append(data.DeltaSharingScope, v.delta_sharing_scope)
			data.DeltaSharingRecipientTokenLifetimeInSeconds = append(data.DeltaSharingRecipientTokenLifetimeInSeconds, v.delta_sharing_recipient_token_lifetime_in_seconds)
			data.DeltaSharingOrganizationName = (data.DeltaSharingOrganizationName, v.delta_sharing_organization_name)
		}
		return nil
	})
}