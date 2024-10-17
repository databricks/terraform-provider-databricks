package scim

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceGroups_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceGroups(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestDataSourceGroups(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockGroupsAPI().EXPECT()
			e.ListAll(mock.Anything, iam.ListGroupsRequest{Attributes: "displayName", Count: 100}).Return(
				[]iam.Group{
					{DisplayName: "ds"}, {DisplayName: "product"},
				}, nil)
		},
		Resource:    DataSourceGroups(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names": []string{
			"ds",
			"product",
		},
	})
}

func TestDataSourceGroups_Filter(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: func(m *mocks.MockWorkspaceClient) {
			e := m.GetMockGroupsAPI().EXPECT()
			e.ListAll(mock.Anything, iam.ListGroupsRequest{Attributes: "displayName", Count: 100, Filter: "displayName sw \"prod\""}).Return(
				[]iam.Group{
					{DisplayName: "product"},
				}, nil)
		},
		Resource:    DataSourceGroups(),
		HCL:         `filter="displayName sw \"prod\""`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"display_names": []string{
			"product",
		},
	})
}
