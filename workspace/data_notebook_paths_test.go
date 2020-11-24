package workspace

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceNotebookPaths(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb%2Fc",
				Response: objectList{
					Objects: []ObjectStatus{
						{
							ObjectID:   987,
							ObjectType: Directory,
							Path:       "/a/b/c/d",
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb%2Fc%2Fd",
				Response: objectList{
					Objects: []ObjectStatus{
						{
							ObjectID:   988,
							ObjectType: Notebook,
							Language:   Python,
							Path:       "/a/b/c/d/e",
						},
						{
							ObjectID:   989,
							ObjectType: Notebook,
							Language:   SQL,
							Path:       "/a/b/c/d/f",
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceNotebookPaths(),
		ID:          ".",
		State: map[string]interface{}{
			"path":      "/a/b/c",
			"recursive": true,
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", d.Id())
}
