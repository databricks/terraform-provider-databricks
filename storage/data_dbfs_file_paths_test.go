package storage

import (
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceFilePaths(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/dbfs/list?path=%2Fa%2Fb%2Fc",
				Response: FileList{
					[]FileInfo{
						{
							Path:  "/a/b/c/d",
							IsDir: true,
						},
					},
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/dbfs/list?path=%2Fa%2Fb%2Fc%2Fd",
				Response: FileList{
					[]FileInfo{
						{
							Path:     "/a/b/c/d/e",
							FileSize: 1024,
							IsDir:    false,
						},
						{
							Path:     "/a/b/c/d/f",
							FileSize: 1025,
							IsDir:    false,
						},
					},
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDBFSFilePaths(),
		ID:          ".",
		State: map[string]interface{}{
			"path":      "/a/b/c",
			"recursive": true,
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", d.Id())
}
