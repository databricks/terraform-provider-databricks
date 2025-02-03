package qa

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/iam"
)

func ListUsersFixtures(users []iam.User) []HTTPFixture {
	return genListFixtures(userResponse{
		&iam.ListUsersResponse{
			StartIndex: 1,
			Resources:  users,
		},
	}, "/api/2.0/preview/scim/v2/Users?attributes=id%2CuserName")
}

func ListGroupsFixtures(groups []iam.Group) []HTTPFixture {
	return genListFixtures(groupsResponse{
		&iam.ListGroupsResponse{
			StartIndex: 1,
			Resources:  groups,
		},
	}, "/api/2.0/preview/scim/v2/Groups?attributes=id")
}

func ListServicePrincipalsFixtures(sps []iam.ServicePrincipal) []HTTPFixture {
	return genListFixtures(servicePrincipalResponse{
		&iam.ListServicePrincipalResponse{
			StartIndex: 1,
			Resources:  sps,
		},
	}, "/api/2.0/preview/scim/v2/ServicePrincipals?attributes=id%2CuserName")
}

type scimResponse interface {
	NumResources() int
	Empty() any
}

type userResponse struct {
	*iam.ListUsersResponse
}

func (u userResponse) NumResources() int {
	return len(u.Resources)
}

func (u userResponse) Empty() any {
	return new(iam.ListUsersResponse)
}

type groupsResponse struct {
	*iam.ListGroupsResponse
}

func (g groupsResponse) NumResources() int {
	return len(g.Resources)
}

func (u groupsResponse) Empty() any {
	return new(iam.ListGroupsResponse)
}

type servicePrincipalResponse struct {
	*iam.ListServicePrincipalResponse
}

func (sp servicePrincipalResponse) NumResources() int {
	return len(sp.Resources)
}

func (sp servicePrincipalResponse) Empty() any {
	return new(iam.ListServicePrincipalResponse)
}

func genListFixtures[R scimResponse](resp R, path string) []HTTPFixture {
	res := make([]HTTPFixture, 0)
	if resp.NumResources() > 0 {
		res = append(res, HTTPFixture{
			Method:   "GET",
			Resource: fmt.Sprintf("%s&count=100&startIndex=1", path),
			Response: resp,
		})
	}
	res = append(res, HTTPFixture{
		Method:   "GET",
		Resource: fmt.Sprintf("%s&count=100&startIndex=%d", path, resp.NumResources()+1),
		Response: resp.Empty(),
	})
	return res
}
