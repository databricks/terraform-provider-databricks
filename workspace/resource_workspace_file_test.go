package workspace

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	ws_api "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestResourceWorkspaceFileRead(t *testing.T) {
	path := "/test/path.py"
	objectID := 12345
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   int64(objectID),
					ObjectType: File,
					Path:       path,
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		New:      true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, objectID, d.Get("object_id"))
}

func TestResourceWorkspaceFileDelete(t *testing.T) {
	path := "/test/path.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: DeletePath{Path: path, Recursive: false},
			},
		},
		Resource: ResourceWorkspaceFile(),
		Delete:   true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceWorkspaceFileRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
				Response: apierr.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		Removed:  true,
		ID:       "/test/path",
	}.ApplyNoError(t)
}

func TestResourceWorkspaceFileRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		ID:       "/test/path",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "/test/path", d.Id(), "Id should not be empty for error reads")
}

func TestResourceWorkspaceFileCreate_DirectoryExist(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/foo",
				},
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Overwrite: true,
					Format:    "AUTO",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ffoo%2Fpath.py",
				Response: ExportPath{
					Content: "YWJjCg==",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ffoo%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       "/foo/path.py",
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           "/foo/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/foo/path.py", d.Id())
}

func TestResourceWorkspaceFileCreate_DirectoryDoesntExist(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/foo",
				},
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Overwrite: true,
					Format:    "AUTO",
				},
				Response: map[string]string{
					"error_code": "RESOURCE_DOES_NOT_EXIST",
					"message":    "The parent folder (/foo) does not exist.",
				},
				Status: 404,
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Overwrite: true,
					Format:    "AUTO",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ffoo%2Fpath.py",
				Response: ExportPath{
					Content: "YWJjCg==",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ffoo%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       "/foo/path.py",
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           "/foo/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/foo/path.py", d.Id())
}

func TestResourceWorkspaceFileCreate_DirectoryCreateError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/foo",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Overwrite: true,
					Format:    "AUTO",
				},
				Response: map[string]string{
					"error_code": "RESOURCE_DOES_NOT_EXIST",
					"message":    "The parent folder (/foo) does not exist.",
				},
				Status: 404,
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           "/foo/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.Error(t, err, "Internal error happened")
}

func TestResourceWorkspaceFileCreateSource(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content: "LS0gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKU0VMRUNUIDEwKjIwC" +
						"gotLSBDT01NQU5EIC0tLS0tLS0tLS0KClNFTEVDVCAyMCoxMDAKCi0tIE" +
						"NPTU1BTkQgLS0tLS0tLS0tLQoKCg==",
					Path:      "/Dashboard",
					Overwrite: true,
					Format:    "AUTO",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDashboard",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       "/Dashboard",
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"source": "acceptance/testdata/tf-test-sql.sql",
			"path":   "/Dashboard",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/Dashboard", d.Id())
}

func TestResourceWorkspaceFileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: map[string]interface{}{
					"content":   "YWJjCg==",
					"format":    "AUTO",
					"overwrite": true,
					"path":      "/path.py",
				},
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           "/path.py",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceWorkspaceFileDelete_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				Response: apierr.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceFile(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}

func TestResourceWorkspaceFileUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Format:    "AUTO",
					Overwrite: true,
					Content:   "YWJjCg==",
					Path:      "abc",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=abc",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       "abc",
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"path":           "/path.py",
		},
		ID:          "abc",
		RequiresNew: true,
		Update:      true,
	}.ApplyNoError(t)
}
