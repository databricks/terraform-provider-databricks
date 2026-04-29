package mws

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceMwsCredentials() common.Resource {
	type mwsCredentialsData struct {
		Ids map[string]string `json:"ids,omitempty" tf:"computed"`
	}
	return common.AccountData(func(ctx context.Context, data *mwsCredentialsData, acc *databricks.AccountClient) error {
		credentials, err := acc.Credentials.List(ctx)
		if err != nil {
			return err
		}
		data.Ids = make(map[string]string)
		for _, v := range credentials {
			data.Ids[v.CredentialsName] = v.CredentialsId
		}
		return nil
	})
}
