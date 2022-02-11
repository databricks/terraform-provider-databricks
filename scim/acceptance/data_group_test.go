package acceptance

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/internal/acceptance"
	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/databrickslabs/terraform-provider-databricks/scim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "10000000-2000-3000-4000-500000000000"
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func TestAccGroupDataSplitMembers(t *testing.T) {
	if cloudEnv, ok := os.LookupEnv("CLOUD_ENV"); !ok || cloudEnv != "azure" {
		t.Skip("This test will only run on Azure. For simplicity.")
	}

	ctx := context.Background()
	client := common.CommonEnvironmentClient()

	usersAPI := scim.NewUsersAPI(ctx, client)
	groupsAPI := scim.NewGroupsAPI(ctx, client)
	spAPI := scim.NewServicePrincipalsAPI(ctx, client)

	user, err := usersAPI.Create(scim.User{
		UserName: fmt.Sprintf("%s@example.com", qa.RandomName("tfuser-")),
	})
	assert.NoError(t, err)
	defer usersAPI.Delete(user.ID)

	sp, err := spAPI.Create(scim.User{
		ApplicationID: createUuid(),
		DisplayName:   qa.RandomName("spn-"),
	})
	assert.NoError(t, err)
	defer spAPI.Delete(sp.ID)

	childGroup, err := groupsAPI.Create(scim.Group{
		DisplayName: qa.RandomName("child-"),
	})
	assert.NoError(t, err)
	defer groupsAPI.Delete(childGroup.ID)

	parentGroup, err := groupsAPI.Create(scim.Group{
		DisplayName: qa.RandomName("parent-"),
		Members: []scim.ComplexValue{
			{Value: user.ID},
			{Value: sp.ID},
			{Value: childGroup.ID},
		},
	})
	assert.NoError(t, err)
	defer groupsAPI.Delete(parentGroup.ID)

	acceptance.Test(t, []acceptance.Step{
		{
			Template: `data "databricks_group" "this" {
				display_name = "` + parentGroup.DisplayName + `"
			}`,
			Check: func(s *terraform.State) error {
				r, ok := s.Modules[0].Resources["data.databricks_group.this"]
				require.True(t, ok, "data.databricks_group.this has to be there")
				attr := r.Primary.Attributes
				assert.Equal(t, user.ID, attr["users.0"])
				assert.Equal(t, sp.ID, attr["service_principals.0"])
				assert.Equal(t, childGroup.ID, attr["child_groups.0"])
				return nil
			},
		},
	})
}
