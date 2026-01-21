package scim

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestReadByDisplayName_CaseSensitive(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%22Name%22",
			Response: GroupList{
				Resources: []Group{
					{
						DisplayName: "name",
						ID:          "1",
					},
					{
						DisplayName: "Name",
						ID:          "2",
					},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/preview/scim/v2/Groups?filter=displayName%20eq%20%22name%22",
			Response: GroupList{
				Resources: []Group{
					{
						DisplayName: "name",
						ID:          "1",
					},
					{
						DisplayName: "Name",
						ID:          "2",
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		api := NewGroupsAPI(ctx, client)

		// Test case-sensitive fetch for upper-case "Name"
		group, err := api.ReadByDisplayName("Name", "")
		assert.NoError(t, err)
		assert.Equal(t, "Name", group.DisplayName)
		assert.Equal(t, "2", group.ID)

		// Test case-sensitive fetch for lower-case "name"
		group, err = api.ReadByDisplayName("name", "")
		assert.NoError(t, err)
		assert.Equal(t, "name", group.DisplayName)
		assert.Equal(t, "1", group.ID)
	})
}
