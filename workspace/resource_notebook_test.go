package workspace

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/qa"

	"github.com/stretchr/testify/assert"
)

func TestResourceNotebookRead(t *testing.T) {
	path := "/test/path.py"
	objectID := 12345
	qa.ResourceFixture{
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
	}.ApplyAndExpectData(t, map[string]any{
		"path":           path,
		"object_id":      objectID,
		"language":       "PYTHON",
		"id":             path,
		"workspace_path": "/Workspace" + path,
		"format":         "SOURCE",
	})
}

func TestResourceNotebookReadWithState(t *testing.T) {
	path := "/test/path.py"
	objectID := 12345
	qa.ResourceFixture{
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
		State: map[string]any{
			"source": "acceptance/testdata/tf-test-jupyter.ipynb",
			"format": "JUPYTER",
			"path":   "/foo/path.py",
		},
		ID: path,
	}.ApplyAndExpectData(t, map[string]any{
		"path":           path,
		"object_id":      objectID,
		"language":       "PYTHON",
		"id":             path,
		"workspace_path": "/Workspace" + path,
		"format":         "JUPYTER",
	})
}

func TestResourceNotebookDelete(t *testing.T) {
	path := "/test/path.py"
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: DeletePath{Path: path, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       path,
	}.ApplyAndExpectData(t, map[string]any{
		"id": path,
	})
}

func TestResourceNotebookRead_NotFound(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
				Response: apierr.APIError{
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
				Response: apierr.APIError{
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

func TestResourceNotebookCreate_DirectoryExist(t *testing.T) {
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
				ExpectedRequest: ImportPath{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
				},
				Response: ImportResponse{
					ObjectID: 12345,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"language":       "PYTHON",
			"path":           "/foo/path.py",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"path":      "/foo/path.py",
		"id":        "/foo/path.py",
		"object_id": 12345,
		"format":    "SOURCE",
	})
}

func TestResourceNotebookCreate_DirectoryDoesntExist(t *testing.T) {
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
				ExpectedRequest: ImportPath{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
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
				ExpectedRequest: ImportPath{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
				},
				Response: ImportResponse{
					ObjectID: 12345,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"language":       "PYTHON",
			"path":           "/foo/path.py",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"path":      "/foo/path.py",
		"id":        "/foo/path.py",
		"object_id": 12345,
	})
}

func TestResourceNotebookCreate_DirectoryCreateError(t *testing.T) {
	_, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/mkdirs",
				ExpectedRequest: map[string]string{
					"path": "/foo",
				},
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportPath{
					Content:   "YWJjCg==",
					Path:      "/foo/path.py",
					Language:  "PYTHON",
					Overwrite: true,
					Format:    "SOURCE",
				},
				Response: map[string]string{
					"error_code": "RESOURCE_DOES_NOT_EXIST",
					"message":    "The parent folder (/foo) does not exist.",
				},
				Status: 404,
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"language":       "PYTHON",
			"path":           "/foo/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.Error(t, err, "Internal error happened")
}

func TestResourceNotebookCreateSource_Jupyter(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportPath{
					Content: "eyJjZWxscyI6W3siY2VsbF90eXBlIjoiY29kZSIsInNvdXJjZSI6WyJwc" +
						"mludChcImhlbGxvIHdvcmxkXCIpIl0sIm1ldGFkYXRhIjp7fSwib3V0cHV" +
						"0cyI6W10sImV4ZWN1dGlvbl9jb3VudCI6MX0seyJjZWxsX3R5cGUiOiJjb" +
						"2RlIiwic291cmNlIjpbInByaW50KFwiaG93IGFyZSB5b3VcIikiXSwibWV" +
						"0YWRhdGEiOnt9LCJvdXRwdXRzIjpbeyJtZXRhZGF0YSI6e30sIm91dHB1d" +
						"F90eXBlIjoiZGlzcGxheV9kYXRhIiwiZGF0YSI6eyJ0ZXh0L2h0bWwiOls" +
						"iPHN0eWxlIHNjb3BlZD5cbiAgLmFuc2lvdXQge1xuICAgIGRpc3BsYXk6I" +
						"GJsb2NrO1xuICAgIHVuaWNvZGUtYmlkaTogZW1iZWQ7XG4gICAgd2hpdGU" +
						"tc3BhY2U6IHByZS13cmFwO1xuICAgIHdvcmQtd3JhcDogYnJlYWstd29yZ" +
						"DtcbiAgICB3b3JkLWJyZWFrOiBicmVhay1hbGw7XG4gICAgZm9udC1mYW1" +
						"pbHk6IFwiU291cmNlIENvZGUgUHJvXCIsIFwiTWVubG9cIiwgbW9ub3NwY" +
						"WNlOztcbiAgICBmb250LXNpemU6IDEzcHg7XG4gICAgY29sb3I6ICM1NTU" +
						"7XG4gICAgbWFyZ2luLWxlZnQ6IDRweDtcbiAgICBsaW5lLWhlaWdodDogM" +
						"TlweDtcbiAgfVxuPC9zdHlsZT5cbjxkaXYgY2xhc3M9XCJhbnNpb3V0XCI" +
						"+aG93IGFyZSB5b3VcbjwvZGl2PiJdfX1dLCJleGVjdXRpb25fY291bnQiO" +
						"jJ9LHsiY2VsbF90eXBlIjoiY29kZSIsInNvdXJjZSI6WyIiXSwibWV0YWR" +
						"hdGEiOnt9LCJvdXRwdXRzIjpbXSwiZXhlY3V0aW9uX2NvdW50IjozfV0sI" +
						"m1ldGFkYXRhIjp7Im5hbWUiOiJ0ZXN0X2p1cHl0ZXIiLCJub3RlYm9va0l" +
						"kIjoxMjc1OTg0MjQzMjkzMDI4fSwibmJmb3JtYXQiOjQsIm5iZm9ybWF0X" +
						"21pbm9yIjowfQo=",
					Path:      "/Mars",
					Language:  "",
					Overwrite: true,
					Format:    "JUPYTER",
				},
				Response: ImportResponse{
					ObjectID: 12345,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2FMars",
				Response: ObjectStatus{
					ObjectID:   12345,
					ObjectType: "NOTEBOOK",
					Language:   "PYTHON",
					Path:       "/Mars",
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"source": "acceptance/testdata/tf-test-jupyter.ipynb",
			"path":   "/Mars",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":        "/Mars",
		"object_id": 12345,
		"language":  "PYTHON",
	})
}

func TestResourceNotebookCreateSource(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportPath{
					Content: "LS0gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKU0VMRUNUIDEwKjIwC" +
						"gotLSBDT01NQU5EIC0tLS0tLS0tLS0KClNFTEVDVCAyMCoxMDAKCi0tIE" +
						"NPTU1BTkQgLS0tLS0tLS0tLQoKCg==",
					Path:      "/Dashboard",
					Language:  "SQL",
					Overwrite: true,
					Format:    "SOURCE",
				},
				Response: ImportResponse{
					ObjectID: 12345,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"source": "acceptance/testdata/tf-test-sql.sql",
			"path":   "/Dashboard",
		},
		Create: true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":        "/Dashboard",
		"object_id": 12345,
		"language":  "SQL",
		"format":    "SOURCE",
	})
}

func TestResourceNotebookCreate_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: apierr.APIError{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
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
				Response: apierr.APIError{
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

func TestResourceNotebookUpdate(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportPath{
					Format:    "SOURCE",
					Overwrite: true,
					Content:   "YWJjCg==",
					Path:      "/path.py",
					Language:  "R",
				},
				Response: ImportResponse{
					ObjectID: 12345,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			"language":       "R",
			"path":           "/path.py",
			"format":         "SOURCE",
		},
		ID:          "/path.py",
		RequiresNew: true,
		Update:      true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":        "/path.py",
		"object_id": 12345,
		"language":  "R",
	})
}

func TestResourceNotebookUpdate_DBC(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/delete",
				ExpectedRequest: DeletePath{
					Recursive: true,
					Path:      "/path.py",
				},
			},
			{
				Method:   "POST",
				Resource: "/api/2.0/workspace/import",
				ExpectedRequest: ImportPath{
					Format:  "DBC",
					Content: "YWJjCg==",
					Path:    "/path.py",
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Fpath.py",
				Response: ObjectStatus{
					ObjectID:   12345,
					ObjectType: Directory,
					Path:       "/path.py",
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]any{
			"content_base64": "YWJjCg==",
			// technically language is not needed, but makes the test simpler
			"language":  "PYTHON",
			"format":    "DBC",
			"path":      "/path.py",
			"object_id": 45678,
		},
		ID:          "/path.py",
		RequiresNew: true,
		Update:      true,
	}.ApplyAndExpectData(t, map[string]any{
		"id":        "/path.py",
		"object_id": 12345,
	})
}
