package storage

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceFile(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/dbfs/get-status?path=%2Fa%2Fb%2Fc",
				Response: FileInfo{
					Path:     "/a/b/c",
					FileSize: 1024,
					IsDir:    false,
				},
			},
			{
				Method:   "GET",
				Resource: "/api/2.0/dbfs/read?length=1000000&path=%2Fa%2Fb%2Fc",
				Response: map[string]interface{}{
					"bytes_read": 1024,
					"data":       "SGVsbG8gd29ybGQK",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDBFSFile(),
		ID:          ".",
		State: map[string]interface{}{
			"path":            "/a/b/c",
			"limit_file_size": true,
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", d.Id())
	assert.Equal(t, "SGVsbG8gd29ybGQK", d.Get("content"))
}
