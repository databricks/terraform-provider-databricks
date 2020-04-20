package service

import (
	"encoding/base64"
	"fmt"
	"github.com/databrickslabs/databricks-terraform/client/model"
	"net/http"
	"testing"
)

var base64String = base64.StdEncoding.EncodeToString([]byte("helloworld"))

func TestDBFSAPI_Create(t *testing.T) {
	type handle struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
	}
	type block struct {
		Data   string `json:"data,omitempty" url:"data,omitempty"`
		Handle int64  `json:"handle,omitempty" url:"handle,omitempty"`
	}
	type close struct {
		Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
	}
	type params struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
		Data      string `json:"data,omitempty" url:"data,omitempty"`
	}
	//var handleInput handle
	//var blockInput block
	//var closeInput close
	tests := []struct {
		params           params
		name             string
		response         []string
		responseStatus   []int
		requestMethod    []string
		postStructExpect []interface{}
		args             []interface{}
		wantUri          []string
		want             interface{}
		wantErr          bool
	}{
		{
			name: "Create test",
			params: params{
				Path:      "my-path",
				Overwrite: true,
				Data:      base64String,
			},
			response: []string{
				`{
					"handle": 1000
				}`, ``, ``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodPost, http.MethodPost},
			args: []interface{}{
				&handle{
					Path:      "my-path",
					Overwrite: true,
				},
				&block{
					Data:   base64String,
					Handle: 1000,
				},
				&close{
					Handle: 1000,
				},
			},
			postStructExpect: []interface{}{
				&handle{},
				&block{},
				&close{},
			},
			wantUri: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/add-block", "/api/2.0/dbfs/close"},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Create bad data encoding failure test",
			params: params{
				Path:      "my-path",
				Overwrite: true,
				Data:      "9",
			},
			response: []string{
				``, ``, ``,
			},
			responseStatus:   []int{},
			requestMethod:    []string{},
			args:             []interface{}{},
			postStructExpect: []interface{}{},
			wantUri:          []string{},
			want:             nil,
			wantErr:          true,
		},
		{
			name: "Create handle failure test",
			params: params{
				Path:      "my-path",
				Overwrite: true,
				Data:      base64String,
			},
			response: []string{
				`{
					"handle": 1000
				}`, ``, ``,
			},
			responseStatus: []int{http.StatusBadRequest},
			requestMethod:  []string{http.MethodPost},
			args: []interface{}{
				&handle{
					Path:      "my-path",
					Overwrite: true,
				},
			},
			postStructExpect: []interface{}{
				&handle{},
			},
			wantUri: []string{"/api/2.0/dbfs/create"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Create add block failure test",
			params: params{
				Path:      "my-path",
				Overwrite: true,
				Data:      base64String,
			},
			response: []string{
				`{
					"handle": 1000
				}`, ``, ``,
			},
			responseStatus: []int{http.StatusOK, http.StatusBadRequest},
			requestMethod:  []string{http.MethodPost, http.MethodPost},
			args: []interface{}{
				&handle{
					Path:      "my-path",
					Overwrite: true,
				},
				&block{
					Data:   base64String,
					Handle: 1000,
				},
			},
			postStructExpect: []interface{}{
				&handle{},
				&block{},
			},
			wantUri: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/add-block"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Create close failure test",
			params: params{
				Path:      "my-path",
				Overwrite: true,
				Data:      base64String,
			},
			response: []string{
				`{
					"handle": 1000
				}`, ``, ``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusBadRequest},
			requestMethod:  []string{http.MethodPost, http.MethodPost, http.MethodPost},
			args: []interface{}{
				&handle{
					Path:      "my-path",
					Overwrite: true,
				},
				&block{
					Data:   base64String,
					Handle: 1000,
				},
				&close{
					Handle: 1000,
				},
			},
			postStructExpect: []interface{}{
				&handle{},
				&block{},
				&close{},
			},
			wantUri: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/add-block", "/api/2.0/dbfs/close"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantUri, tt.postStructExpect, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.DBFS().Create(tt.params.Path, tt.params.Overwrite, tt.params.Data)
			})
		})
	}
}

func TestDBFSAPI_Copy(t *testing.T) {
	type handle struct {
		Path      string `json:"path,omitempty" url:"path,omitempty"`
		Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
	}
	type addBlock struct {
		Data   string `json:"data,omitempty" url:"data,omitempty"`
		Handle int64  `json:"handle,omitempty" url:"handle,omitempty"`
	}
	type closeHandle struct {
		Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
	}
	type read struct {
		Path   string `json:"path,omitempty" url:"path,omitempty"`
		Offset int64  `json:"offset,omitempty" url:"offset,omitempty"`
		Length int64  `json:"length,omitempty" url:"length,omitempty"`
	}
	type params struct {
		Src       string `json:"src,omitempty" url:"src,omitempty"`
		Tgt       string `json:"tgt,omitempty" url:"tgt,omitempty"`
		Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
	}

	tests := []struct {
		params           params
		name             string
		response         []string
		responseStatus   []int
		requestMethod    []string
		postStructExpect []interface{}
		args             []interface{}
		wantUri          []string
		want             interface{}
		wantErr          bool
	}{
		{
			name: "Copy test",
			params: params{
				Src:       "my-path",
				Tgt:       "my-path-tgt",
				Overwrite: true,
			},
			response: []string{
				`{
					"handle": 1000
				}`, fmt.Sprintf(`{
					"bytes_read": 10000,
					"data": "%s"
				}`, base64String), ``, ``,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodPost, http.MethodGet, http.MethodPost, http.MethodPost},
			args: []interface{}{
				&handle{
					Path:      "my-path-tgt",
					Overwrite: true,
				},
				&read{
					Path:   "my-path",
					Offset: 0,
					Length: 1e6,
				},
				&addBlock{
					Data:   base64String,
					Handle: 1000,
				},
				&closeHandle{
					Handle: 1000,
				},
			},
			postStructExpect: []interface{}{
				&handle{},
				nil,
				&addBlock{},
				&closeHandle{},
			},
			wantUri: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/read?length=1000000&path=my-path", "/api/2.0/dbfs/add-block", "/api/2.0/dbfs/close"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantUri, tt.postStructExpect, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.DBFS().Copy(tt.params.Src, tt.params.Tgt, &client, tt.params.Overwrite)
			})
		})
	}
}

func TestDBFSAPI_Read(t *testing.T) {
	type read struct {
		Path   string `json:"path,omitempty" url:"path,omitempty"`
		Offset int64  `json:"offset,omitempty" url:"offset,omitempty"`
		Length int64  `json:"length,omitempty" url:"length,omitempty"`
	}

	type params struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}
	tests := []struct {
		params           params
		name             string
		response         []string
		responseStatus   []int
		requestMethod    []string
		postStructExpect []interface{}
		args             []interface{}
		wantUri          []string
		want             interface{}
		wantErr          bool
	}{
		{
			name: "Read test",
			params: params{
				Path: "my-path",
			},
			response: []string{
				fmt.Sprintf(`{
					"bytes_read": 1000000,
					"data": "%s"
				}`, base64String),
				fmt.Sprintf(`{
					"bytes_read": 0,
					"data": ""
				}`),
			},
			responseStatus: []int{http.StatusOK, http.StatusOK},
			requestMethod:  []string{http.MethodGet, http.MethodGet},
			args: []interface{}{
				&read{
					Path:   "my-path",
					Offset: 0,
					Length: 1e6,
				},
				&read{
					Path:   "my-path",
					Offset: 1e6,
					Length: 1e6,
				},
			},
			postStructExpect: []interface{}{
				&read{},
				&read{},
			},
			wantUri: []string{"/api/2.0/dbfs/read?length=1000000&path=my-path", "/api/2.0/dbfs/read?length=1000000&offset=1000000&path=my-path"},
			want:    base64String,
			wantErr: false,
		},
		{
			name: "Read fetch block failure test",
			params: params{
				Path: "my-path",
			},
			response: []string{
				fmt.Sprintf(`{
					"bytes_read": 1000000,
					"data": "%s"
				}`, base64String),
				fmt.Sprintf(`{
					"bytes_read": 0,
					"data": ""
				}`),
			},
			responseStatus: []int{http.StatusOK, http.StatusBadRequest},
			requestMethod:  []string{http.MethodGet, http.MethodGet},
			args: []interface{}{
				&read{
					Path:   "my-path",
					Offset: 0,
					Length: 1e6,
				},
				&read{
					Path:   "my-path",
					Offset: 1e6,
					Length: 1e6,
				},
			},
			postStructExpect: []interface{}{
				&read{},
				&read{},
			},
			wantUri: []string{"/api/2.0/dbfs/read?length=1000000&path=my-path", "/api/2.0/dbfs/read?length=1000000&offset=1000000&path=my-path"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantUri, tt.postStructExpect, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.DBFS().Read(tt.params.Path)
			})
		})
	}
}

func TestDBFSAPI_Delete(t *testing.T) {
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
		{
			name:           "Delete failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				Path:      "mypath",
				Recursive: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/dbfs/delete", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.DBFS().Delete(tt.args.Path, tt.args.Recursive)
			})
		})
	}
}

func TestDBFSAPI_Move(t *testing.T) {
	type args struct {
		SourcePath      string `json:"source_path,omitempty" url:"source_path,omitempty"`
		DestinationPath string `json:"destination_path,omitempty" url:"destination_path,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Move test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				SourcePath:      "mypath",
				DestinationPath: "targetpath",
			},
			wantErr: false,
		},
		{
			name:           "Move failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				SourcePath:      "mypath",
				DestinationPath: "targetpath",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/dbfs/move", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.DBFS().Move(tt.args.SourcePath, tt.args.DestinationPath)
			})
		})
	}
}

func TestDBFSAPI_Mkdirs(t *testing.T) {
	type args struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantErr        bool
	}{
		{
			name:           "Mkdirs test",
			response:       "",
			responseStatus: http.StatusOK,
			args: args{
				Path: "mypath",
			},
			wantErr: false,
		},
		{
			name:           "Mkdirs failure test",
			response:       "",
			responseStatus: http.StatusBadRequest,
			args: args{
				Path: "mypath",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/dbfs/mkdirs", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return nil, client.DBFS().Mkdirs(tt.args.Path)
			})
		})
	}
}

func TestDBFSAPI_Status(t *testing.T) {
	type args struct {
		Path string `json:"path,omitempty" url:"path,omitempty"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		requestUri     string
		args           args
		want           interface{}
		wantErr        bool
	}{
		{
			name: "Status test",
			response: `{
							  "path": "/a.cpp",
							  "is_dir": false,
							  "file_size": 261
							}`,
			requestUri:     "/api/2.0/dbfs/get-status?path=mypath",
			responseStatus: http.StatusOK,
			args: args{
				Path: "mypath",
			},
			want: model.FileInfo{
				Path:     "/a.cpp",
				IsDir:    false,
				FileSize: 261,
			},
			wantErr: false,
		},
		{
			name:           "Status failure test",
			response:       "",
			requestUri:     "/api/2.0/dbfs/get-status?path=mypath",
			responseStatus: http.StatusBadRequest,
			args: args{
				Path: "mypath",
			},
			want:    model.FileInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.requestUri, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.DBFS().Status(tt.args.Path)
			})
		})
	}
}

func TestDBFSAPI_ListNonRecursive(t *testing.T) {
	type args struct {
		Path      string `json:"path"`
		Recursive bool   `json:"recursive"`
	}
	tests := []struct {
		name           string
		response       string
		responseStatus int
		args           args
		wantUri        string
		want           []model.FileInfo
		wantErr        bool
	}{
		{
			name: "List non recursive test",
			response: `{
						  "files": [
							{
							  "path": "/a.cpp",
							  "is_dir": false,
							  "file_size": 261
							},
							{
							  "path": "/foldera",
							  "is_dir": true,
							  "file_size": 0
							}
						  ]
						}`,
			responseStatus: http.StatusOK,
			args: args{

				Path:      "/",
				Recursive: false,
			},
			wantUri: "/api/2.0/dbfs/list?path=%2F",
			want: []model.FileInfo{
				{
					Path:     "/a.cpp",
					IsDir:    false,
					FileSize: 261,
				},
				{
					Path:     "/foldera",
					IsDir:    true,
					FileSize: 0,
				},
			},
			wantErr: false,
		},
		{
			name:           "List non recursive failure test",
			response:       ``,
			responseStatus: http.StatusBadRequest,
			args: args{

				Path:      "/",
				Recursive: false,
			},
			wantUri: "/api/2.0/dbfs/list?path=%2F",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantUri, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.DBFS().List(tt.args.Path, tt.args.Recursive)
			})
		})
	}
}

func TestDBFSAPI_ListRecursive(t *testing.T) {
	type args struct {
		Path      string `json:"path"`
		Recursive bool   `json:"recursive"`
	}
	tests := []struct {
		name           string
		response       []string
		responseStatus []int
		args           []interface{}
		wantUri        []string
		want           []model.FileInfo
		wantErr        bool
	}{
		{
			name: "List recursive test",
			response: []string{`{
								  "files": [
									{
									  "path": "/a.cpp",
									  "is_dir": false,
									  "file_size": 261
									},
									{
									  "path": "/foldera",
									  "is_dir": true,
									  "file_size": 0
									}
								  ]
								}`,
				`{
								  "files": [
									{
									  "path": "/foldera/b.cpp",
									  "is_dir": false,
									  "file_size": 200
									}
								  ]
								}`,
			},
			responseStatus: []int{http.StatusOK, http.StatusOK},
			args: []interface{}{
				&args{
					Path:      "/",
					Recursive: true,
				},
			},
			wantUri: []string{"/api/2.0/dbfs/list?path=%2F", "/api/2.0/dbfs/list?path=%2Ffoldera"},
			want: []model.FileInfo{
				{
					Path:     "/a.cpp",
					IsDir:    false,
					FileSize: 261,
				},
				{
					Path:     "/foldera/b.cpp",
					IsDir:    false,
					FileSize: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "List recursive test failure",
			response: []string{`{
								  "files": [
									{
									  "path": "/a.cpp",
									  "is_dir": false,
									  "file_size": 261
									},
									{
									  "path": "/foldera",
									  "is_dir": true,
									  "file_size": 0
									}
								  ]
								}`, ``,
			},
			responseStatus: []int{http.StatusOK, http.StatusBadRequest},
			args: []interface{}{
				&args{
					Path:      "/",
					Recursive: true,
				},
			},
			wantUri: []string{"/api/2.0/dbfs/list?path=%2F", "/api/2.0/dbfs/list?path=%2Ffoldera"},
			want:    nil,
			wantErr: true,
		},
		{
			name:           "List first recursive test failure",
			response:       []string{``},
			responseStatus: []int{http.StatusBadRequest},
			args: []interface{}{
				&args{
					Path:      "/",
					Recursive: true,
				},
			},
			wantUri: []string{"/api/2.0/dbfs/list?path=%2F"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodGet}, tt.wantUri, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client DBApiClient) (interface{}, error) {
				return client.DBFS().List(tt.args[0].(*args).Path, tt.args[0].(*args).Recursive)
			})
		})
	}
}
