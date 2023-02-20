package workspace

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceNotebook(t *testing.T) {
	d, err := qa.ResourceFixture{
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
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", d.Id())
	assert.Equal(t, "SGVsbG8gd29ybGQK", d.Get("content"))
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
	}.ExpectError(t, "I'm a teapot")
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
