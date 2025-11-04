package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMwsCredentials() common.Resource {
	type mwsCredentialsData struct {
		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	s := common.StructToSchema(mwsCredentialsData{}, nil)
	common.AddNamespaceInSchema(s)
	common.NamespaceCustomizeSchemaMap(s)
	return common.Resource{
		Schema: s,
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			newClient, err := c.DatabricksClientForUnifiedProvider(ctx, d)
			if err != nil {
				return err
			}
			var data mwsCredentialsData
			common.DataToStructPointer(d, s, &data)
			if newClient.Config.AccountID == "" {
				return fmt.Errorf("provider block is missing `account_id` property")
			}
			credentials, err := NewCredentialsAPI(ctx, newClient).List(newClient.Config.AccountID)
			if err != nil {
				return err
			}
			data.Ids = make(map[string]string)
			for _, v := range credentials {
				data.Ids[v.CredentialsName] = v.CredentialsID
			}
			err = common.StructToData(data, s, d)
			if err != nil {
				return err
			}
			d.SetId("_")
			return nil
		},
	}
}
