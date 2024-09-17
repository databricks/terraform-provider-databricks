package scim

import (
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceDataUsers_DisplayNameContains(t *testing.T) {
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

func TestDataSourceDataUsers_UserNameContains(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountUsersAPI()
			e.On("ListAll",
				mock.Anything,
				mock.MatchedBy(func(req iam.ListAccountUsersRequest) bool {
					return req.Filter == `userName co "testuser"` &&
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
		user_name_contains = "testuser"
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

func TestDataSourceDataUsers_BothFiltersSpecified(t *testing.T) {
	qa.ResourceFixture{
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
        display_name_contains = "user"
        user_name_contains    = "example.com"
        `,
	}.ExpectError(t, "exactly one of display_name_contains or user_name_contains should be specified, not both")
}

func TestDataSourceDataUsersNoUsers(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountUsersAPI()
			e.On("ListAll",
				mock.Anything,
				mock.MatchedBy(func(req iam.ListAccountUsersRequest) bool {
					return req.Filter == `displayName co "testuser"` &&
						req.Attributes == "id,userName,displayName"
				})).Return([]iam.User{}, nil)
		},
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		display_name_contains = "testuser"
		`,
	}.ExpectError(t, "cannot find users with display name containing testuser")
}

func TestDataSourceDataUsersNoFilter(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			e := m.GetMockAccountUsersAPI()
			e.On("ListAll",
				mock.Anything,
				mock.MatchedBy(func(req iam.ListAccountUsersRequest) bool {
					return req.Attributes == "id,userName,displayName"
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
				{
					Id:          "789",
					UserName:    "testuser3@example.com",
					DisplayName: "testuser3",
				},
			}, nil)
		},
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"users.#":              3,
		"users.0.id":           "123",
		"users.0.user_name":    "testuser@example.com",
		"users.0.display_name": "testuser",
		"users.1.id":           "456",
		"users.1.user_name":    "testuser2@example.com",
		"users.1.display_name": "testuser2",
		"users.2.id":           "789",
		"users.2.user_name":    "testuser3@example.com",
		"users.2.display_name": "testuser3",
	})
}

func TestDataSourceDataUsers_APIError(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(m *mocks.MockAccountClient) {
			usersAPI := m.GetMockAccountUsersAPI()
			usersAPI.On("ListAll",
				mock.Anything,
				mock.Anything,
			).Return(nil, fmt.Errorf("api error"))
		},
		Resource:    DataSourceDataUsers(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
		HCL: `
		display_name_contains = "testuser"
		`,
	}.ExpectError(t, "api error")
}
