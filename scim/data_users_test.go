package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceDataUsers(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountUsersAPI()
			e.On("ListAll",
				mock.Anything,
				mock.MatchedBy(func(req iam.ListAccountUsersRequest) bool {
					return req.Filter == `displayName co "testuser"` &&
						req.Attributes == "id,userName,displayName"
				})).Return([]iam.User{
				{
					Id:          "123",
					UserName:    "testuser@example.com",
					DisplayName: "testuser",
				},
				{
					Id:          "456",
					UserName:    "testuser2@example.com",
					DisplayName: "testuser2",
				},
			}, nil)
		},
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		display_name_contains = "testuser"
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"users.#":              2,
		"users.0.id":           "123",
		"users.0.user_name":    "testuser@example.com",
		"users.0.display_name": "testuser",
		"users.1.id":           "456",
		"users.1.user_name":    "testuser2@example.com",
		"users.1.display_name": "testuser2",
	})
}

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}
