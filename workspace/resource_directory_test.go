package workspace

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceDirectoryRead(t *testing.T) {
	path := "/test/path"
	objectID := 12345
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: ObjectStatus{
					ObjectID:   int64(objectID),
					ObjectType: Directory,
					Path:       path,
				},
			},
		},
		Resource: ResourceDirectory(),
		Read:     true,
		New:      true,
		ID:       path,
	}.ApplyAndExpectData(t, map[string]any{
		"id": path, "path": path, "workspace_path": "/Workspace" + path, "object_id": objectID,
	})
}

func TestResourceDirectoryDelete(t *testing.T) {
	path := "/test/path"
	delete_recursive := true
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: DeletePath{Path: path, Recursive: delete_recursive},
			},
		},
		Resource: ResourceDirectory(),
		Delete:   true,
		ID:       path,
		State: map[string]any{
			"path":             "/foo/path.py",
			"delete_recursive": delete_recursive,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceDirectoryDelete_NotFound(t *testing.T) {
	path := "/test/path"
	delete_recursive := true
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				ExpectedRequest: DeletePath{Path: path, Recursive: delete_recursive},
				Response: common.APIErrorBody{
					ErrorCode: "RESOURCE_DOES_NOT_EXIST",
					Message:   "Path (/test/path) doesn't exist.",
				},
				Status: 404,
			},
		},
		Resource: ResourceDirectory(),
		Delete:   true,
		ID:       path,
		State: map[string]any{
			"path":             "/foo/path.py",
			"delete_recursive": delete_recursive,
		},
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceDirectoryRead_NotFound(t *testing.T) {
	path := "/test/path"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceDirectory(),
		Read:     true,
		Removed:  true,
		ID:       "/test/path",
	}.ApplyNoError(t)
}

func TestResourceDirectoryRead_Error(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceDirectory(),
		Read:     true,
		ID:       "/test/path",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "/test/path", d.Id(), "Id should not be empty for error reads")
}

func TestResourceDirectoryCreate(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": path,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: "DIRECTORY",
					Path:       path,
				},
			},
		},
		Resource: ResourceDirectory(),
		State: map[string]any{
			"object_id":        4567,
			"path":             path,
			"delete_recursive": false,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceDirectoryCreate_Error(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": path,
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceDirectory(),
		State: map[string]any{
			"path": path,
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceDirectoryDelete_Error(t *testing.T) {
	path := "/test/path"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          "POST",
				Resource:        "/api/2.0/workspace/delete",
				ExpectedRequest: DeletePath{Path: path, Recursive: false},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceDirectory(),
		Delete:   true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, path, d.Id())
}

func TestResourceDirectoryUpdate(t *testing.T) {
	path := "/test/path"
	object_id := 4567
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": path,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: ObjectStatus{
					ObjectID:   int64(object_id),
					ObjectType: "DIRECTORY",
					Path:       path,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: ObjectStatus{
					ObjectID:   int64(object_id),
					ObjectType: "DIRECTORY",
					Path:       path,
				},
			},
		},
		Resource: ResourceDirectory(),
		InstanceState: map[string]string{
			"path":             path,
			"delete_recursive": "true",
		},
		State: map[string]any{
			"object_id": object_id,
			"path":      path,
		},
		ID:     path,
		Update: true,
	}.Apply(t)
	require.NoError(t, err)
}

func TestResourceDirectoryReadNotDirectory(t *testing.T) {
	path := "/test/path"
	objectID := 12345
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/workspace/get-status?path=%s", url.PathEscape(path)),
				Response: ObjectStatus{
					ObjectID:   int64(objectID),
					ObjectType: Notebook,
					Path:       path,
					Language:   Python,
				},
			},
		},
		Resource: ResourceDirectory(),
		Read:     true,
		New:      true,
		ID:       path,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "different object type")
	assert.Equal(t, "", d.Id(), "Id should be empty for different object type read")
}

func TestDirectoryPathSuppressDiff(t *testing.T) {
	assert.True(t, directoryPathSuppressDiff("", "/TF_DIR_WITH_SLASH", "/TF_DIR_WITH_SLASH/", nil))
	assert.False(t, directoryPathSuppressDiff("", "/new_dir", "/TF_DIR_WITH_SLASH/", nil))
}
