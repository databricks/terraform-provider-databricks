package mws

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/mock"
)

func TestDataSourceMwsWorkspaces(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{
				{
					WorkspaceName: "bcd",
					WorkspaceId:   123,
				},
				{
					WorkspaceName: "def",
					WorkspaceId:   456,
				},
			}, nil)
		},
		AccountID:   "abc",
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{
			"bcd": 123,
			"def": 456,
		},
	})
}

func TestDataSourceMwsWorkspaces_Error(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return(nil, errors.New("i'm a teapot"))
		},
		AccountID:   "abc",
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestDataSourceMwsWorkspaces_Empty(t *testing.T) {
	qa.ResourceFixture{
		MockAccountClientFunc: func(a *mocks.MockAccountClient) {
			a.GetMockWorkspacesAPI().EXPECT().List(mock.Anything).Return([]provisioning.Workspace{}, nil)
		},
		AccountID:   "abc",
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyAndExpectData(t, map[string]any{
		"ids": map[string]any{},
	})
}
