package scim

import (
	"context"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
)

// DataSourceGroups searches from groups based on filter
func DataSourceGroups() common.Resource {
	type groupData struct {
		DisplayNameFilter string   `json:"filter,omitempty"`
		DisplayNames      []string `json:"display_names,omitempty" tf:"computed,slice_set"`
	}

	return common.WorkspaceData(func(ctx context.Context, data *groupData, w *databricks.WorkspaceClient) error {
		groupSearch := iam.ListGroupsRequest{Attributes: "displayName", Count: 100}
		groupSearch.Filter = data.DisplayNameFilter

		groups, err := w.Groups.ListAll(ctx, groupSearch)
		if err != nil {
			return err
		}
		data.DisplayNames = make([]string, 0, len(groups))
		for _, v := range groups {
			data.DisplayNames = append(data.DisplayNames, v.DisplayName)
		}
		return nil
	})
}
