package workspace

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestDataSourceDirectoryPaths(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb%2Fc",
				Response: ObjectList{
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
				Response: ObjectList{
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
		Resource:    DataSourceDirectoryPaths(),
		ID:          ".",
		State: map[string]any{
			"path":      "/a/b/c",
			"recursive": true,
		},
	}.ApplyNoError(t)
}

func TestDataSourceDirectoryPaths_NoRecursive(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb%2Fc",
				Response: ObjectList{
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
		Resource:    DataSourceDirectoryPaths(),
		ID:          ".",
		State: map[string]any{
			"path":      "/a/b/c",
			"recursive": false,
		},
	}.ApplyNoError(t)
}
