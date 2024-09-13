package scim

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
)

func DataSourceDataUsers() common.Resource {

	type DataUsers struct {
		Id                  string                   `json:"id,omitempty" tf:"computed"`
		DisplayNameContains string                   `json:"display_name_contains,omitempty" tf:"computed"`
		Users               []map[string]interface{} `json:"users,omitempty" tf:"computed"`
	}

	return common.AccountData(func(ctx context.Context, data *DataUsers, acc *databricks.AccountClient) error {
		listRequest := iam.ListAccountUsersRequest{}

		if data.DisplayNameContains != "" {
			listRequest.Filter = fmt.Sprintf("displayName co \"%s\"", data.DisplayNameContains)
		}

		userList, err := acc.Users.ListAll(ctx, listRequest)

		if err != nil {
			return err
		}

		if len(userList) == 0 {
			return fmt.Errorf("cannot find users with display name containing %s", data.DisplayNameContains)
		}

		var users []map[string]interface{}

		for _, u := range userList {
			user := map[string]interface{}{
				"id":           u.Id,
				"user_name":    u.UserName,
				"display_name": u.DisplayName,
			}
			users = append(users, user)
		}

		data.Users = users

		return nil
	})
}
