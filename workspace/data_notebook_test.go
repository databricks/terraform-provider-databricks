package workspace

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceNotebook(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Fa%2Fb%2Fc",
				Response: ObjectStatus{
					ObjectID:   987,
					Language:   "PYTHON",
					ObjectType: "NOTEBOOK",
					Path:       "/a/b/c",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Fa%2Fb%2Fc",
				Response: ExportPath{
					Content: "SGVsbG8gd29ybGQK",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceNotebook(),
		ID:          ".",
		State: map[string]any{
			"path":   "/a/b/c",
			"format": "SOURCE",
		},
	}.ApplyAndExpectData(t, map[string]any{
		"id":             "/a/b/c",
		"content":        "SGVsbG8gd29ybGQK",
		"workspace_path": "/Workspace/a/b/c",
	})
}

func TestDataSourceNotebook_ErrorExport(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceNotebook(),
		ID:          ".",
		State: map[string]any{
			"path":   "/a/b/c",
			"format": "SOURCE",
		},
	}.ExpectError(t, "i'm a teapot")
}

func TestDataSourceNotebook_ErrorStatus(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Fa%2Fb%2Fc",
				Response: ObjectStatus{
					ObjectID:   987,
					Language:   "PYTHON",
					ObjectType: "NOTEBOOK",
					Path:       "/a/b/c",
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Fa%2Fb%2Fc",
				Status:   401,
				Response: apierr.APIError{
					ErrorCode:  "Unauthorized",
					StatusCode: 401,
					Message:    "Unauthorized",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceNotebook(),
		ID:          ".",
		State: map[string]any{
			"path":   "/a/b/c",
			"format": "SOURCE",
		},
	}.ExpectError(t, "Unauthorized")
}
