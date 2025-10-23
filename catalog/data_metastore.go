package catalog

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMetastore() common.Resource {
	type AccountMetastoreByID struct {
		common.Namespace
		Id          string                 `json:"id,omitempty" tf:"computed"`
		MetastoreId string                 `json:"metastore_id,omitempty" tf:"computed"`
		Name        string                 `json:"name,omitempty" tf:"computed"`
		Region      string                 `json:"region,omitempty" tf:"computed"`
		Metastore   *catalog.MetastoreInfo `json:"metastore_info,omitempty" tf:"computed" `
	}
	return common.AccountData(func(ctx context.Context, data *AccountMetastoreByID, acc *databricks.AccountClient) error {
		if data.MetastoreId == "" && data.Name == "" && data.Region == "" {
			return fmt.Errorf("one of metastore_id, name or region must be provided")
		}
		if (data.MetastoreId != "" && data.Name != "") || (data.Region != "" && data.MetastoreId != "") || (data.Region != "" && data.Name != "") {
			return fmt.Errorf("only one of metastore_id, name or region must be provided")
		}
		if data.MetastoreId != "" {
			minfo, err := acc.Metastores.GetByMetastoreId(ctx, data.MetastoreId)
			if err != nil {
				return err
			}
			data.Metastore = minfo.MetastoreInfo
		} else {
			metastores, err := acc.Metastores.ListAll(ctx)
			if err != nil {
				return err
			}
			minfos := []catalog.MetastoreInfo{}
			if data.Name != "" {
				for _, v := range metastores {
					if strings.EqualFold(v.Name, data.Name) {
						minfos = append(minfos, v)
					}
				}
			} else {
				for _, v := range metastores {
					if strings.EqualFold(v.Region, data.Region) {
						minfos = append(minfos, v)
					}
				}
			}
			if len(minfos) == 0 {
				return fmt.Errorf("a metastore with name '%s' or in region '%s' is not found", data.Name, data.Region)
			}
			if len(minfos) > 1 {
				return fmt.Errorf("there are %d metastores with name '%s' in region '%s'", len(minfos), data.Name, data.Region)
			}
			data.Metastore = &minfos[0]
		}
		data.Id = data.Metastore.MetastoreId
		data.MetastoreId = data.Metastore.MetastoreId
		data.Name = data.Metastore.Name
		data.Region = data.Metastore.Region
		return nil
	})
}
