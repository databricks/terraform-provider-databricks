package service

import (
	"net/http"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

func TestNotebooksAPI_Create(t *testing.T) {
	type args struct {
		Content   string             `json:"content,omitempty"`
		Path      string             `json:"path,omitempty"`
		Language  model.Language     `json:"language,omitempty"`
		Overwrite bool               `json:"overwrite,omitempty"`
		Format    model.ExportFormat `json:"format,omitempty"`
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
				Language:  model.Python,
				Overwrite: false,
				Format:    model.DBC,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/workspace/import", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Notebooks().Create(tt.args.Path, tt.args.Content, tt.args.Language, tt.args.Format, tt.args.Overwrite)
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/workspace/mkdirs", &input, tt.response, http.StatusOK, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Notebooks().Mkdirs(tt.args.Path)
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/workspace/delete", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.Notebooks().Delete(tt.args.Path, tt.args.Recursive)
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
		want           []model.NotebookInfo
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
			want: []model.NotebookInfo{
				{
					ObjectID:   123,
					ObjectType: model.Directory,
					Path:       "/Users/user@example.com/project",
				},
				{
					ObjectID:   456,
					ObjectType: model.Notebook,
					Language:   model.Python,
					Path:       "/Users/user@example.com/PythonExampleNotebook",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Notebooks().List(tt.args.Path, tt.args.Recursive)
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
		want           []model.NotebookInfo
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
			wantURI: []string{"/api/2.0/workspace/list?path=%2Ftest%2Fpath", "/api/2.0/workspace/list?path=%2FUsers%2Fuser%40example.com%2Fproject"},
			want: []model.NotebookInfo{
				{
					ObjectID:   457,
					ObjectType: model.Notebook,
					Language:   model.Python,
					Path:       "/Users/user@example.com/Notebook2",
				},
				{
					ObjectID:   456,
					ObjectType: model.Notebook,
					Language:   model.Python,
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
			wantURI: []string{"/api/2.0/workspace/list?path=%2Ftest%2Fpath", "/api/2.0/workspace/list?path=%2FUsers%2Fuser%40example.com%2Fproject"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodGet}, tt.wantURI, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Notebooks().List(tt.args[0].(*args).Path, tt.args[0].(*args).Recursive)
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
		want           model.NotebookInfo
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
			want: model.NotebookInfo{
				ObjectID:   789,
				ObjectType: model.Notebook,
				Path:       "/Users/user@example.com/project/ScalaExampleNotebook",
				Language:   model.Scala,
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
			want:           model.NotebookInfo{},
			wantURI:        "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Notebooks().Read(tt.args.Path)
			})
		})
	}
}

func TestNotebooksAPI_Export(t *testing.T) {
	type args struct {
		Path   string             `json:"path"`
		Format model.ExportFormat `json:"format"`
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
				Format: model.DBC,
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
				Format: model.DBC,
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
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.Notebooks().Export(tt.args.Path, tt.args.Format)
			})
		})
	}
}
