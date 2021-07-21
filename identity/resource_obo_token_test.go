package identity

import (
	"context"
	"os"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAwsAccOboFlow(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	ctx := context.Background()
	client := common.CommonEnvironmentClient()
	tmAPI := NewTokenManagementAPI(ctx, client)
	spAPI := NewServicePrincipalsAPI(ctx, client)
	groupsAPI := NewGroupsAPI(ctx, client)

	sp, err := spAPI.Create(ScimUser{
		DisplayName: qa.RandomName("tf"),
	})
	require.NoError(t, err)
	defer spAPI.Delete(sp.ID)

	admins, err := groupsAPI.ReadByDisplayName("admins")
	require.NoError(t, err)

	// add new SP to admins, as it's the easiest way to give it permissions to do so,
	// because importing access package will create circular import dependency
	err = groupsAPI.Patch(admins.ID, scimPatchRequest("add", "members", sp.ID))
	require.NoError(t, err)

	oboToken, err := tmAPI.CreateTokenOnBehalfOfServicePrincipal(OboToken{
		ApplicationID:   sp.ApplicationID,
		Comment:         qa.RandomLongName(),
		LifetimeSeconds: 60,
	})
	require.NoError(t, err)
	defer tmAPI.Delete(oboToken.TokenInfo.TokenID)

	spClient := &common.DatabricksClient{
		Host:  client.Host,
		Token: oboToken.TokenValue,
	}
	err = spClient.Configure()
	require.NoError(t, err)

	newMe, err := NewUsersAPI(ctx, spClient).Me()
	require.NoError(t, err)
	assert.Equal(t, newMe.DisplayName, sp.DisplayName)

	r, err := tmAPI.Read(oboToken.TokenInfo.TokenID)
	require.NoError(t, err)
	assert.Equal(t, r.TokenInfo.TokenID, oboToken.TokenInfo.TokenID)
}
