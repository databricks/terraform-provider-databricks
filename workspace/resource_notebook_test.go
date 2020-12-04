package workspace

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
)

// func TestResourceNotebookCreate_DirDoesNotExists(t *testing.T) {
// 	pythonNotebookDataB64, err := notebookToB64("acceptance/testdata/tf-test-python.py")
// 	assert.NoError(t, err, err)
// 	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
// 	assert.NoError(t, err, err)
// 	path := "/test/path.py"
// 	content := pythonNotebookDataB64
// 	objectId := 12345

// 	d, err := qa.ResourceFixture{
// 		Fixtures: []qa.HTTPFixture{
// 			{
// 				Method:   http.MethodGet,
// 				Resource: "/api/2.0/workspace/get-status?path=%2Ftest",
// 				Response: common.APIErrorBody{
// 					ErrorCode: "NOT_FOUND",
// 					Message:   "not found",
// 				},
// 				Status: 404,
// 			},
// 			{
// 				Method:   http.MethodPost,
// 				Resource: "/api/2.0/workspace/mkdirs",
// 				ExpectedRequest: map[string]string{
// 					"path": "/test",
// 				},
// 			},
// 			{
// 				Method:   http.MethodPost,
// 				Resource: "/api/2.0/workspace/import",
// 				Response: ImportRequest{
// 					Content:   content,
// 					Path:      path,
// 					Language:  "PYTHON",
// 					Overwrite: true,
// 					Format:    "SOURCE",
// 				},
// 			},
// 			{
// 				Method:   http.MethodGet,
// 				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
// 				Response: NotebookContent{
// 					Content: pythonNotebookDataB64,
// 				},
// 			},
// 			{
// 				Method:   http.MethodGet,
// 				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
// 				Response: ObjectStatus{
// 					ObjectID:   int64(objectId),
// 					ObjectType: Notebook,
// 					Path:       path,
// 					Language:   "PYTHON",
// 				},
// 			},
// 		},
// 		Resource: ResourceNotebook(),
// 		State: map[string]interface{}{
// 			"path":     path,
// 			"content":  content,
// 			"language": "PYTHON",
// 		},
// 		Create: true,
// 	}.Apply(t)
// 	assert.NoError(t, err, err)
// 	assert.Equal(t, path, d.Id())
// 	assert.Equal(t, checkSum, d.Get("content"))
// 	assert.Equal(t, path, d.Get("path"))
// 	assert.Equal(t, "PYTHON", d.Get("language"))
// 	assert.Equal(t, objectId, d.Get("object_id"))
// }

func TestResourceNotebookRead(t *testing.T) {
	path := "/test/path.py"
	objectID := 12345
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   int64(objectID),
					ObjectType: Notebook,
					Path:       path,
					Language:   "PYTHON",
				},
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		New:      true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, "PYTHON", d.Get("language"))
	assert.Equal(t, objectID, d.Get("object_id"))
}

func TestResourceNotebookDelete(t *testing.T) {
	path := "/test/path.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: NotebookDeleteRequest{Path: path, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       path,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
}

func TestResourceNotebookDelete_TooManyRequests(t *testing.T) {
	testID := "/test/path.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/delete",
				Status:   http.StatusTooManyRequests,
			},
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: NotebookDeleteRequest{Path: testID, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       testID,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testID, d.Id())
}

func TestResourceNotebookRead_NotFound(t *testing.T) {
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
		Resource: ResourceNotebook(),
		Read:     true,
		Removed:  true,
		ID:       "/test/path",
	}.ApplyNoError(t)
}

func TestResourceNotebookRead_Error(t *testing.T) {
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
		Resource: ResourceNotebook(),
		Read:     true,
		ID:       "/test/path",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "/test/path", d.Id(), "Id should not be empty for error reads")
}

func TestResourceNotebookCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportRequest{
					Content:   "YWJjCg==",
					Path:      "/path.py",
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Fpath.py",
				Response: NotebookContent{
					Content: "YWJjCg==",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: "NOTEBOOK",
					Path:       "/path.py",
					Language:   "PYTHON",
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"content_base64": "YWJjCg==",
			"language":       "PYTHON",
			"path":           "/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/path.py", d.Id())
}

func TestResourceNotebookCreateSource(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportRequest{
					Content: "LS0gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKU0VMRUNUIDEwKjIwC" +
						"gotLSBDT01NQU5EIC0tLS0tLS0tLS0KClNFTEVDVCAyMCoxMDAKCi0tIE" +
						"NPTU1BTkQgLS0tLS0tLS0tLQoKCg==",
					Path:      "/Dashboard",
					Language:  "SQL",
					Overwrite: true,
					Format:    "SOURCE",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FDashboard",
				Response: ObjectStatus{
					ObjectID:   4567,
					ObjectType: "NOTEBOOK",
					Path:       "/Dashboard",
					Language:   "SQL",
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"source": "acceptance/testdata/tf-test-sql.sql",
			"path":   "/Dashboard",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/Dashboard", d.Id())
}

func TestResourceNotebookCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"content_base64": "YWJjCg==",
			"language":       "R",
			"path":           "/path.py",
		},
		Create: true,
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "", d.Id(), "Id should be empty for error creates")
}

func TestResourceNotebookDelete_Error(t *testing.T) {
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
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       "abc",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "abc", d.Id())
}
