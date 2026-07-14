package exporter

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/databricks/terraform-provider-databricks/scim"
	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("groups,access")
	ic.allGroups = []scim.Group{
		{
			DisplayName: "foo",
			ID:          "123",
			Roles: []scim.ComplexValue{
				{
					Value: "abc",
					Type:  "direct",
				},
			},
			Members: []scim.ComplexValue{
				// this is just for log line printing
				{Value: "a001"},
				{Value: "a002"},
				{Value: "a003"},
				{Value: "a004"},
				{Value: "a005"},
				{Value: "a006"},
				{Value: "a007"},
				{Value: "a008"},
				{Value: "a009"},
				{Value: "a010"},
				{Value: "a011"},
			},
			Groups: []scim.ComplexValue{
				{
					Value: "parent-group",
					Type:  "direct",
				},
			},
		},
	}
	d := scim.ResourceGroup().ToResource().TestResourceData()
	d.Set("display_name", "foo")
	r := &resource{
		Value:     "foo",
		Attribute: "display_name",
		Data:      d,
	}
	err := ic.Importables["databricks_group"].Search(ic, r)
	assert.NoError(t, err)
	assert.Equal(t, "123", r.ID)

	err = ic.Importables["databricks_group"].Import(ic, r)
	assert.NoError(t, err)
	assert.Len(t, ic.testEmits, 4)
	assert.True(t, ic.testEmits["databricks_group_role[<unknown>] (id: 123|abc)"])
	assert.True(t, ic.testEmits["databricks_instance_profile[<unknown>] (id: abc)"])
	assert.True(t, ic.testEmits["databricks_group[<unknown>] (id: parent-group)"])
	assert.True(t, ic.testEmits["databricks_group_member[_parent-group_foo] (id: parent-group|123)"])
}

func TestGroupCacheAndSearchError(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Groups?attributes=id&count=10000&startIndex=1",
			Status:       404,
			Response: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		ic.enableServices("groups")
		err := resourcesMap["databricks_group"].List(ic)
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_group"].Search(ic, &resource{
			Attribute: "display_name",
		})
		assert.EqualError(t, err, "nope")

		err = resourcesMap["databricks_group"].Search(ic, &resource{
			Attribute: "nonsense",
		})
		assert.EqualError(t, err, "wrong search attribute 'nonsense' for databricks_group")
		d := scim.ResourceGroup().ToResource().TestResourceData()
		d.Set("display_name", "nonsense")
		err = resourcesMap["databricks_group"].Import(ic, &resource{
			ID:   "nonsense",
			Data: d,
		})
		assert.EqualError(t, err, "nope")
	})
}

func TestGroupListNoNameMatch(t *testing.T) {
	ic := importContextForTest()
	ic.match = "bcd"
	ic.allGroups = []scim.Group{
		{
			DisplayName: "abc",
		},
	}
	err := resourcesMap["databricks_group"].List(ic)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(ic.testEmits))
}

func TestGroupSearchNoMatch(t *testing.T) {
	ic := importContextForTest()
	ic.allGroups = []scim.Group{
		{
			DisplayName: "abc",
		},
	}
	r := &resource{
		Attribute: "display_name",
		Value:     "dbc",
	}
	err := resourcesMap["databricks_group"].Search(ic, r)
	assert.NoError(t, err)
	assert.Equal(t, "", r.ID)
}

func TestUserSearchFails(t *testing.T) {
	userFixture := qa.ListUsersFixtures([]iam.User{})
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userFixture[0],
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceUser().ToResource().TestResourceData()
		d.Set("user_name", "dbc")
		r := &resource{
			Attribute: "display_name",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_user"].Search(ic, r)
		assert.EqualError(t, err, "there is no user 'dbc'")

		err = resourcesMap["databricks_user"].Import(ic, r)
		assert.EqualError(t, err, "user dbc is not found")
	})
}

func TestSpnSearchFails(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		qa.ListServicePrincipalsFixtures([]iam.ServicePrincipal{})[0],
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceServicePrincipal().ToResource().TestResourceData()
		d.Set("application_id", "dbc")
		r := &resource{
			Attribute: "application_id",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_service_principal"].Search(ic, r)
		assert.EqualError(t, err, "there is no service principal 'dbc'")

		err = resourcesMap["databricks_service_principal"].Import(ic, r)
		assert.EqualError(t, err, "service principal dbc is not found")
	})
}

func TestSpnSearchSuccess(t *testing.T) {
	spFixture := qa.ListServicePrincipalsFixtures([]iam.ServicePrincipal{
		{
			Id: "321", DisplayName: "spn", ApplicationId: "dbc",
		},
	})
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		spFixture[0],
		spFixture[1],
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/ServicePrincipals/321?attributes=userName,displayName,active,externalId,entitlements,groups,roles",
			Response:     scim.User{ID: "321", DisplayName: "spn", ApplicationID: "dbc"},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceServicePrincipal().ToResource().TestResourceData()
		d.Set("application_id", "dbc")
		d.Set("display_name", "dbc")
		r := &resource{
			Attribute: "application_id",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_service_principal"].Search(ic, r)
		assert.NoError(t, err)

		err = resourcesMap["databricks_service_principal"].Import(ic, r)
		assert.NoError(t, err)

		assert.True(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "application_id",
			scim.ResourceServicePrincipal().Schema["application_id"], d, r))
		ic.Client.Config.Host = "https://abc.azuredatabricks.net"
		// We shouldn't omit display_name for Databricks-managed SPs
		assert.False(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "display_name",
			scim.ResourceServicePrincipal().Schema["display_name"], d, r))
		assert.True(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "application_id",
			scim.ResourceServicePrincipal().Schema["application_id"], d, r))
		// We shouldn't omit application_id for Azure-managed SPs, but omit display_name
		d.Set("external_id", "60622399-fd3f-4faf-8810-bf08b225cf3b")
		assert.False(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "application_id",
			scim.ResourceServicePrincipal().Schema["application_id"], d, r))
		assert.True(t, resourcesMap["databricks_service_principal"].ShouldOmitField(ic, "display_name",
			scim.ResourceServicePrincipal().Schema["display_name"], d, r))

		// test for different branches in Name function
		// test for different branches in Name function
		d2 := scim.ResourceServicePrincipal().ToResource().TestResourceData()
		d2.SetId("123")
		d2.Set("application_id", "dbc")
		assert.Equal(t, "dbc_123", resourcesMap["databricks_service_principal"].Name(ic, d2))
		d2.Set("application_id", "60622399-fd3f-4faf-8810-bf08b225cf3b")
		assert.Equal(t, "60622399_123", resourcesMap["databricks_service_principal"].Name(ic, d2))

		d2.Set("display_name", "abc")
		assert.Equal(t, "abc_123", resourcesMap["databricks_service_principal"].Name(ic, d2))
	})
}

func TestShouldOmitForUsers(t *testing.T) {
	d := scim.ResourceUser().ToResource().TestResourceData()
	d.SetId("user1")
	d.Set("user_name", "user@domain.com")
	d.Set("display_name", "")
	r := &resource{
		Attribute: "databricks_user",
		Value:     "user@domain.com",
		Data:      d,
	}
	assert.True(t, resourcesMap["databricks_user"].ShouldOmitField(nil, "display_name",
		scim.ResourceUser().Schema["display_name"], d, r))
	d.Set("display_name", "user@domain.com")
	assert.True(t, resourcesMap["databricks_user"].ShouldOmitField(nil, "display_name",
		scim.ResourceUser().Schema["display_name"], d, r))
	d.Set("display_name", "Some user")
	assert.False(t, resourcesMap["databricks_user"].ShouldOmitField(nil, "display_name",
		scim.ResourceUser().Schema["display_name"], d, r))
}

func TestUserImportSkipNonDirectGroups(t *testing.T) {
	userFixture := qa.ListUsersFixtures([]iam.User{
		{
			UserName: "dbc",
			Id:       "321",
		},
	})
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		userFixture[0],
		userFixture[1],
		{
			ReuseRequest: true,
			Method:       "GET",
			Resource:     "/api/2.0/preview/scim/v2/Users/321?attributes=id,userName,displayName,active,externalId,entitlements,groups,roles",
			Response: scim.UserList{
				Resources: []scim.User{
					{
						Groups: []scim.ComplexValue{
							{
								Display: "x",
								Value:   "y",
							},
						},
					},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		ic := importContextForTestWithClient(ctx, client)
		d := scim.ResourceUser().ToResource().TestResourceData()
		d.Set("user_name", "dbc")
		r := &resource{
			Attribute: "display_name",
			Value:     "dbc",
			Data:      d,
		}
		err := resourcesMap["databricks_user"].Import(ic, r)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(ic.testEmits))
	})
}
