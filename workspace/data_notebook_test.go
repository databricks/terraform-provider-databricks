package workspace

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
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
				Response: NotebookContent{
					Content: "SGVsbG8gd29ybGQK",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceNotebook(),
		ID:          ".",
		State: map[string]interface{}{
			"path":   "/a/b/c",
			"format": "SOURCE",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", d.Id())
	assert.Equal(t, "SGVsbG8gd29ybGQK", d.Get("content"))
}
