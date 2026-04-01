package mws

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceMwsWorkspaces(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces",

				Response: []Workspace{
					{
						WorkspaceName: "bcd",
						WorkspaceID:   123,
					},
					{
						WorkspaceName: "def",
						WorkspaceID:   456,
					},
				},
			},
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

func TestCatalogsData_Error(t *testing.T) {
	qa.ResourceFixture{
		AccountID:   "abc",
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceMwsWorkspaces(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "i'm a teapot")
}

func TestDataSourceMwsWorkspaces_Empty(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/accounts/abc/workspaces",

				Response: []Workspace{},
			},
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
