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
						WorkspaceName:  "bcd",
						WorkspaceID:    123,
						DeploymentName: "deployment1",
					},
					{
						WorkspaceName:  "def",
						WorkspaceID:    456,
						DeploymentName: "deployment2",
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
		//"mws_workspaces": []any{
		//	map[string]any{"workspace_name": "bcd", "workspace_id": 123, "deployment_name": "deployment1"},
		//	map[string]any{"workspace_name": "def", "workspace_id": 456, "deployment_name": "deployment2"},
		//},
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
		"ids":            map[string]any{},
		"mws_workspaces": []any{},
	})
}
