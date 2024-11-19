package workspace

import (
	"net/http"
	"testing"

	ws_api "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

var (
	dummyWorkspaceFilePath    = "/foo/path.py"
	dummyWorkspaceFilePathUrl = "path=%2Ffoo%2Fpath.py"
	dummyWorkspaceFilePayload = "YWJjCg=="
)

func TestResourceWorkspaceFileRead(t *testing.T) {
	objectID := 12345
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?" + dummyWorkspaceFilePathUrl,
				Response: ObjectStatus{
					ObjectID:   int64(objectID),
					ObjectType: File,
					Path:       dummyWorkspaceFilePath,
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		Read:     true,
		New:      true,
		ID:       dummyWorkspaceFilePath,
	}.ApplyAndExpectData(t, map[string]any{
		"id":             dummyWorkspaceFilePath,
		"path":           dummyWorkspaceFilePath,
		"workspace_path": "/Workspace" + dummyWorkspaceFilePath,
		"object_id":      objectID,
	})
}

func TestResourceWorkspaceFileDelete(t *testing.T) {
	path := "/test/path.py"
	qa.ResourceFixture{
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
	}.ApplyAndExpectData(t, map[string]any{
		"id": path,
	})
}

func TestResourceWorkspaceFileRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
				Response: common.APIErrorBody{
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
				Response: common.APIErrorBody{
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
	qa.ResourceFixture{
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
					Content:   dummyWorkspaceFilePayload,
					Path:      dummyWorkspaceFilePath,
					Overwrite: true,
					Format:    "RAW",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&" + dummyWorkspaceFilePathUrl,
				Response: ExportPath{
					Content: dummyWorkspaceFilePayload,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?" + dummyWorkspaceFilePathUrl,
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       dummyWorkspaceFilePath,
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           dummyWorkspaceFilePath,
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id": dummyWorkspaceFilePath,
	})
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
					Content:   dummyWorkspaceFilePayload,
					Path:      dummyWorkspaceFilePath,
					Overwrite: true,
					Format:    "RAW",
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
					Content:   dummyWorkspaceFilePayload,
					Path:      dummyWorkspaceFilePath,
					Overwrite: true,
					Format:    "RAW",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&" + dummyWorkspaceFilePathUrl,
				Response: ExportPath{
					Content: dummyWorkspaceFilePayload,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?" + dummyWorkspaceFilePathUrl,
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       dummyWorkspaceFilePath,
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
			"path":           dummyWorkspaceFilePath,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, dummyWorkspaceFilePath, d.Id())
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
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content:   dummyWorkspaceFilePayload,
					Path:      dummyWorkspaceFilePath,
					Overwrite: true,
					Format:    "RAW",
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
			"content_base64": dummyWorkspaceFilePayload,
			"path":           dummyWorkspaceFilePath,
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
					Format:    "RAW",
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

func TestResourceWorkspaceFileCreateEmptyFileSource(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ws_api.Import{
					Content:         "",
					Path:            "/__init__.py",
					Overwrite:       true,
					Format:          "RAW",
					ForceSendFields: []string{"Content"},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2F__init__.py",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: File,
					Path:       "/__init__.py",
				},
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"source": "acceptance/testdata/empty_file",
			"path":   "/__init__.py",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, "/__init__.py", d.Id())
}

func TestResourceWorkspaceFileCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: map[string]interface{}{
					"content":   dummyWorkspaceFilePayload,
					"format":    "RAW",
					"overwrite": true,
					"path":      "/path.py",
				},
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceWorkspaceFile(),
		State: map[string]any{
			"content_base64": dummyWorkspaceFilePayload,
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
				Response: common.APIErrorBody{
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
					Format:    "RAW",
					Overwrite: true,
					Content:   dummyWorkspaceFilePayload,
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
			"content_base64": dummyWorkspaceFilePayload,
			"path":           "/path.py",
		},
		ID:          "abc",
		RequiresNew: true,
		Update:      true,
	}.ApplyNoError(t)
}
