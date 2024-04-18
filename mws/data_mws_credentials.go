package mws

import (
	"context"
	"fmt"

	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsCredentials() common.Resource {
	type mwsCredentialsData struct {
		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.DataResource(mwsCredentialsData{}, func(ctx context.Context, e any, c *common.DatabricksClient) error {
		data := e.(*mwsCredentialsData)
		if c.Config.AccountID == "" {
			return fmt.Errorf("provider block is missing `account_id` property")
		}
		credentials, err := NewCredentialsAPI(ctx, c).List(c.Config.AccountID)
		if err != nil {
			return err
		}
		data.Ids = make(map[string]string)
		for _, v := range credentials {
			data.Ids[v.CredentialsName] = v.CredentialsID
		}
		return nil
	})
}
