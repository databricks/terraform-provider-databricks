package workspace

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataSourceDirectory(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Fa%2Fb%2Fc",
				Response: ObjectStatus{
					ObjectID:   987,
					ObjectType: "DIRECTORY",
					Path:       "/a/b/c",
				},
			},
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDirectory(),
		ID:          ".",
		State: map[string]any{
			"path": "/a/b/c",
		},
	}.Apply(t)
	require.NoError(t, err)
	assert.Equal(t, "/a/b/c", d.Id())
	assert.Equal(t, 987, d.Get("object_id").(int))
}

func TestDataSourceDirectory_NotDirectory(t *testing.T) {
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
		},
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDirectory(),
		ID:          ".",
		State: map[string]any{
			"path": "/a/b/c",
		},
	}.ExpectError(t, "'/a/b/c' isn't a directory")
}

func TestDataSourceDirectory_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Read:        true,
		NonWritable: true,
		Resource:    DataSourceDirectory(),
		ID:          ".",
		State: map[string]any{
			"path": "/a/b/c",
		},
	}.ExpectError(t, "I'm a teapot")
}
