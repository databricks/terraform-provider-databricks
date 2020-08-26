package workspace

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/stretchr/testify/assert"
)

func notebookToB64(filePath string) (string, error) {
	notebookBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to find notebook to convert to base64; %w", err)
	}
	return base64.StdEncoding.EncodeToString(notebookBytes), nil
}

func TestValidateNotebookPath(t *testing.T) {
	testCases := []struct {
		name         string
		notebookPath string
		errorCount   int
	}{
		{"empty_path",
			"",
			2},
		{"correct_path",
			"/directory",
			0},
		{"path_starts_with_no_slash",
			"directory",
			1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, errs := ValidateNotebookPath(tc.notebookPath, "key")

			assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.notebookPath)
		})
	}
}

func TestResourceNotebookCreate_DirDoesNotExists(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("acceptance/testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	path := "/test/path.py"
	content := pythonNotebookDataB64
	objectId := 12345

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "not found",
				},
				Status: 404,
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/mkdirs",
				Response: NotebookImportRequest{
					Content:   content,
					Path:      path,
					Language:  Python,
					Overwrite: true,
					Format:    Source,
				},
			},
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: NotebookImportRequest{
					Content:   content,
					Path:      path,
					Language:  Python,
					Overwrite: true,
					Format:    Source,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: NotebookContent{
					Content: pythonNotebookDataB64,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: WorkspaceObjectStatus{
					ObjectID:   int64(objectId),
					ObjectType: Notebook,
					Path:       path,
					Language:   Python,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"path":      path,
			"content":   content,
			"language":  string(Python),
			"format":    string(Source),
			"overwrite": true,
			"mkdirs":    true,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, string(Python), d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestResourceNotebookCreate_NoMkdirs(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("acceptance/testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	path := "/test/path.py"
	content := pythonNotebookDataB64
	objectId := 12345

	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: NotebookImportRequest{
					Content:   content,
					Path:      path,
					Language:  Python,
					Overwrite: true,
					Format:    Source,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: NotebookContent{
					Content: pythonNotebookDataB64,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: WorkspaceObjectStatus{
					ObjectID:   int64(objectId),
					ObjectType: Notebook,
					Path:       path,
					Language:   Python,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"path":      path,
			"content":   content,
			"language":  string(Python),
			"format":    string(Source),
			"overwrite": true,
			"mkdirs":    false,
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, string(Python), d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestResourceNotebookRead(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("acceptance/testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	exportFormat := Source
	testId := "/test/path.py"
	objectId := 12345
	assert.NoError(t, err, err)
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: NotebookContent{
					Content: pythonNotebookDataB64,
				},
			},
			{
				Method:   http.MethodGet,
				Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
				Response: WorkspaceObjectStatus{
					ObjectID:   int64(objectId),
					ObjectType: Notebook,
					Path:       testId,
					Language:   Python,
				},
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		ID:       testId,
		State: map[string]interface{}{
			"format": exportFormat,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, testId, d.Get("path"))
	assert.Equal(t, string(Python), d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestResourceNotebookDelete(t *testing.T) {
	testId := "/test/path.py"
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:          http.MethodPost,
				Resource:        "/api/2.0/workspace/delete",
				Status:          http.StatusOK,
				ExpectedRequest: NotebookDeleteRequest{Path: testId, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       testId,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
}

func TestResourceNotebookDelete_TooManyRequests(t *testing.T) {
	testId := "/test/path.py"
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
				ExpectedRequest: NotebookDeleteRequest{Path: testId, Recursive: true},
			},
		},
		Resource: ResourceNotebook(),
		Delete:   true,
		ID:       testId,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
}

func TestResourceNotebookRead_NotFound(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: common.APIErrorBody{
					ErrorCode: "NOT_FOUND",
					Message:   "Item not found",
				},
				Status: 404,
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		ID:       "/test/path.py",
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id(), "Id should be empty for missing resources")
}

func TestResourceNotebookRead_Error(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{ // read log output for correct url...
				Method:   "GET",
				Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
				Response: common.APIErrorBody{
					ErrorCode: "INVALID_REQUEST",
					Message:   "Internal error happened",
				},
				Status: 400,
			},
		},
		Resource: ResourceNotebook(),
		Read:     true,
		ID:       "/test/path.py",
	}.Apply(t)
	qa.AssertErrorStartsWith(t, err, "Internal error happened")
	assert.Equal(t, "/test/path.py", d.Id(), "Id should not be empty for error reads")
}

func TestResourceNotebookCreate(t *testing.T) {
	d, err := qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   http.MethodPost,
				Resource: "/api/2.0/workspace/import",
				Response: NotebookImportRequest{
					Content:   "YWJjCg==",
					Path:      "/path.py",
					Language:  Python,
					Overwrite: true,
					Format:    Source,
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
				Response: WorkspaceObjectStatus{
					ObjectID:   4567,
					ObjectType: "NOTEBOOK",
					Path:       "/path.py",
					Language:   Python,
				},
			},
		},
		Resource: ResourceNotebook(),
		State: map[string]interface{}{
			"content":   "YWJjCg==",
			"format":    "SOURCE",
			"language":  "PYTHON",
			"overwrite": true,
			"path":      "/path.py",
		},
		Create: true,
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, "/path.py", d.Id())
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
			"content":   "YWJjCg==",
			"format":    "SOURCE",
			"language":  "PYTHON",
			"overwrite": true,
			"path":      "/path.py",
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

func TestNotebooksAPI_Create(t *testing.T) {
	type args struct {
		Content   string       `json:"content,omitempty"`
		Path      string       `json:"path,omitempty"`
		Language  Language     `json:"language,omitempty"`
		Overwrite bool         `json:"overwrite,omitempty"`
		Format    ExportFormat `json:"format,omitempty"`
	}

	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Create Test",
			response: "",
			args: args{
				Content:   "helloworld",
				Path:      "my-path",
				Language:  Python,
				Overwrite: false,
				Format:    DBC,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/workspace/import", &input, tt.response, http.StatusOK, nil,
				tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
					return nil, NewNotebooksAPI(client).Create(tt.args.Path, tt.args.Content, tt.args.Language, tt.args.Format, tt.args.Overwrite)
				})
		})
	}
}

func TestNotebooksAPI_MkDirs(t *testing.T) {
	type args struct {
		Path string `json:"path,omitempty"`
	}

	tests := []struct {
		name     string
		response string
		args     args
		wantErr  bool
	}{
		{
			name:     "Create Test",
			response: "",
			args: args{
				Path: "/test/path",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/workspace/mkdirs", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewNotebooksAPI(client).Mkdirs(tt.args.Path)
			})
		})
	}
}

func TestNotebooksAPI_Delete(t *testing.T) {
	type args struct {
		Path      string `json:"path,omitempty"`
		Recursive bool   `json:"recursive,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Delete test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				Path:      "mypath",
				Recursive: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/workspace/delete", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewNotebooksAPI(client).Delete(tt.args.Path, tt.args.Recursive)
			})
		})
	}
}

func TestNotebooksAPI_ListNonRecursive(t *testing.T) {
	type args struct {
		Path      string `json:"path"`
		Recursive bool   `json:"recursive"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantURI        string
		want           []WorkspaceObjectStatus
		wantErr        bool
	}{
		{
			name: "List non recursive test",
			response: `{
						  "objects": [
							{
							  "path": "/Users/user@example.com/project",
							  "object_type": "DIRECTORY",
							  "object_id": 123
							},
							{
							  "path": "/Users/user@example.com/PythonExampleNotebook",
							  "language": "PYTHON",
							  "object_type": "NOTEBOOK",
							  "object_id": 456
							}
						  ]
						}`,
			responseStatus: http.StatusOK,
			args: args{

				Path:      "/test/path",
				Recursive: false,
			},
			wantURI: "/api/2.0/workspace/list?path=%2Ftest%2Fpath",
			want: []WorkspaceObjectStatus{
				{
					ObjectID:   123,
					ObjectType: Directory,
					Path:       "/Users/user@example.com/project",
				},
				{
					ObjectID:   456,
					ObjectType: Notebook,
					Language:   Python,
					Path:       "/Users/user@example.com/PythonExampleNotebook",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewNotebooksAPI(client).List(tt.args.Path, tt.args.Recursive)
			})
		})
	}
}

func TestNotebooksAPI_ListRecursive(t *testing.T) {
	type args struct {
		Path      string `json:"path"`
		Recursive bool   `json:"recursive"`
	}
	tests := []struct {
		name           string
		response       []string
		responseStatus []int
		args           []interface{}
		wantURI        []string
		want           []WorkspaceObjectStatus
		wantErr        bool
	}{
		{
			name: "List recursive test",
			response: []string{`{
						  "objects": [
							{
							  "path": "/Users/user@example.com/project",
							  "object_type": "DIRECTORY",
							  "object_id": 123
							},
							{
							  "path": "/Users/user@example.com/PythonExampleNotebook",
							  "language": "PYTHON",
							  "object_type": "NOTEBOOK",
							  "object_id": 456
							}
						  ]
						}`,
				`{
						  "objects": [
							{
							  "path": "/Users/user@example.com/Notebook2",
							  "language": "PYTHON",
							  "object_type": "NOTEBOOK",
							  "object_id": 457
							}
						  ]
						}`,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK},
			args: []interface{}{
				&args{
					Path:      "/test/path",
					Recursive: true,
				},
			},
			wantURI: []string{"/api/2.0/workspace/list?path=%2Ftest%2Fpath", "/api/2.0/workspace/list?path=%2FUsers%2Fuser@example.com%2Fproject"},
			want: []WorkspaceObjectStatus{
				{
					ObjectID:   457,
					ObjectType: Notebook,
					Language:   Python,
					Path:       "/Users/user@example.com/Notebook2",
				},
				{
					ObjectID:   456,
					ObjectType: Notebook,
					Language:   Python,
					Path:       "/Users/user@example.com/PythonExampleNotebook",
				},
			},
			wantErr: false,
		},
		{
			name: "List recursive test failure",
			response: []string{`{
						  "objects": [
							{
							  "path": "/Users/user@example.com/project",
							  "object_type": "DIRECTORY",
							  "object_id": 123
							},
							{
							  "path": "/Users/user@example.com/PythonExampleNotebook",
							  "language": "PYTHON",
							  "object_type": "NOTEBOOK",
							  "object_id": 456
							}
						  ]
						}`,
				``,
			},
			responseStatus: []int{http.StatusOK, http.StatusBadRequest},
			args: []interface{}{
				&args{
					Path:      "/test/path",
					Recursive: true,
				},
			},
			wantURI: []string{"/api/2.0/workspace/list?path=%2Ftest%2Fpath", "/api/2.0/workspace/list?path=%2FUsers%2Fuser@example.com%2Fproject"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodGet}, tt.wantURI, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewNotebooksAPI(client).List(tt.args[0].(*args).Path, tt.args[0].(*args).Recursive)
			})
		})
	}
}

func TestNotebooksAPI_Read(t *testing.T) {
	type args struct {
		Path string `json:"path"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		wantURI        string
		want           WorkspaceObjectStatus
		wantErr        bool
	}{
		{
			name: "Read test",
			response: `{
						  "path": "/Users/user@example.com/project/ScalaExampleNotebook",
						  "language": "SCALA",
						  "object_type": "NOTEBOOK",
						  "object_id": 789
						}`,
			args: args{
				Path: "/test/path",
			},
			responseStatus: http.StatusOK,
			want: WorkspaceObjectStatus{
				ObjectID:   789,
				ObjectType: Notebook,
				Path:       "/Users/user@example.com/project/ScalaExampleNotebook",
				Language:   Scala,
			},
			wantURI: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
			wantErr: false,
		},

		{
			name:     "Read test failure",
			response: ``,
			args: args{
				Path: "/test/path",
			},
			responseStatus: http.StatusBadRequest,
			want:           WorkspaceObjectStatus{},
			wantURI:        "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewNotebooksAPI(client).Read(tt.args.Path)
			})
		})
	}
}

func TestNotebooksAPI_Export(t *testing.T) {
	type args struct {
		Path   string       `json:"path"`
		Format ExportFormat `json:"format"`
	}
	tests := []struct {
		name           string
		response       string
		args           args
		responseStatus int
		wantURI        string
		want           string
		wantErr        bool
	}{
		{
			name: "Export test",
			response: `{
						  "content": "Ly8gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKMSsx"
						}`,
			args: args{
				Path:   "/test/path",
				Format: DBC,
			},
			responseStatus: http.StatusOK,
			want:           "Ly8gRGF0YWJyaWNrcyBub3RlYm9vayBzb3VyY2UKMSsx",
			wantURI:        "/api/2.0/workspace/export?format=DBC&path=%2Ftest%2Fpath",
			wantErr:        false,
		},
		{
			name:     "Export test failure",
			response: ``,
			args: args{
				Path:   "/test/path",
				Format: DBC,
			},
			responseStatus: http.StatusBadRequest,
			want:           "",
			wantURI:        "/api/2.0/workspace/export?format=DBC&path=%2Ftest%2Fpath",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewNotebooksAPI(client).Export(tt.args.Path, tt.args.Format)
			})
		})
	}
}

func TestUri(t *testing.T) {
	uri := "https://sri-e2-test-workspace-3.cloud.databricks.com/api/2.0/workspace/export?format=DBC\u0026path=/demo-notebook-rbc"
	t.Log(url.PathUnescape(uri))
}

// nolint this should be refactored to support dbc files
func convertZipBytesToCRC(b64 []byte) (string, error) {
	r, err := zip.NewReader(bytes.NewReader(b64), int64(len(b64)))
	if err != nil {
		return "0", err
	}
	var totalSum int64
	for _, f := range r.File {
		if f.FileInfo().IsDir() == false {
			file, err := f.Open()
			if err != nil {
				return "", err
			}
			crc, err := getDBCCheckSumForCommands(file)
			if err != nil {
				return "", err
			}
			totalSum += int64(crc)
		}
	}
	return strconv.Itoa(int(totalSum)), nil
}

// nolint this should be refactored to support dbc files
func getDBCCheckSumForCommands(fileIO io.Reader) (int, error) {
	var stringBuff bytes.Buffer
	scanner := bufio.NewScanner(fileIO)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		stringBuff.WriteString(scanner.Text())
	}
	jsonString := stringBuff.Bytes()
	var notebook map[string]interface{}
	err := json.Unmarshal(jsonString, &notebook)
	if err != nil {
		return 0, err
	}
	var commandsBuffer bytes.Buffer
	commandsMap := map[int]string{}
	commands := notebook["commands"].([]interface{})
	for _, command := range commands {
		commandsMap[int(command.(map[string]interface{})["position"].(float64))] = command.(map[string]interface{})["command"].(string)
	}
	keys := make([]int, 0, len(commandsMap))
	for k := range commandsMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		commandsBuffer.WriteString(commandsMap[k])
	}
	return int(crc32.ChecksumIEEE(commandsBuffer.Bytes())), nil
}
