package storage

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
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
		wantURI          []string
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
			wantURI: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/add-block", "/api/2.0/dbfs/close"},
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
			wantURI:          []string{},
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
			wantURI: []string{"/api/2.0/dbfs/create"},
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
			responseStatus: []int{http.StatusOK, http.StatusBadRequest, http.StatusOK},
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
			wantURI: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/add-block", "/api/2.0/dbfs/close"},
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
			wantURI: []string{"/api/2.0/dbfs/create", "/api/2.0/dbfs/add-block", "/api/2.0/dbfs/close"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI,
				tt.postStructExpect, tt.response, tt.responseStatus, nil, tt.wantErr,
				func(client *common.DatabricksClient) (interface{}, error) {
					return nil, NewDBFSAPI(client).Create(tt.params.Path, tt.params.Overwrite, tt.params.Data)
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
		wantURI          []string
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
				`{
					"bytes_read": 0,
					"data": ""
				}`,
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
			wantURI: []string{"/api/2.0/dbfs/read?length=1000000&path=my-path", "/api/2.0/dbfs/read?length=1000000&offset=1000000&path=my-path"},
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
				`{
					"bytes_read": 0,
					"data": ""
				}`,
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
			wantURI: []string{"/api/2.0/dbfs/read?length=1000000&path=my-path", "/api/2.0/dbfs/read?length=1000000&offset=1000000&path=my-path"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertMultipleRequestsWithMockServer(t, tt.args, tt.requestMethod, tt.wantURI, tt.postStructExpect, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewDBFSAPI(client).Read(tt.params.Path)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/dbfs/delete", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewDBFSAPI(client).Delete(tt.args.Path, tt.args.Recursive)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/dbfs/move", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewDBFSAPI(client).Move(tt.args.SourcePath, tt.args.DestinationPath)
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
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodPost, "/api/2.0/dbfs/mkdirs", &input, tt.response, tt.responseStatus, nil, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return nil, NewDBFSAPI(client).Mkdirs(tt.args.Path)
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
		requestURI     string
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
			requestURI:     "/api/2.0/dbfs/get-status?path=mypath",
			responseStatus: http.StatusOK,
			args: args{
				Path: "mypath",
			},
			want: DBFSFileInfo{
				Path:     "/a.cpp",
				IsDir:    false,
				FileSize: 261,
			},
			wantErr: false,
		},
		{
			name:           "Status failure test",
			response:       "",
			requestURI:     "/api/2.0/dbfs/get-status?path=mypath",
			responseStatus: http.StatusBadRequest,
			args: args{
				Path: "mypath",
			},
			want:    DBFSFileInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, &tt.args, http.MethodGet, tt.requestURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewDBFSAPI(client).Status(tt.args.Path)
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
		wantURI        string
		want           []DBFSFileInfo
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
			wantURI: "/api/2.0/dbfs/list?path=%2F",
			want: []DBFSFileInfo{
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
			wantURI: "/api/2.0/dbfs/list?path=%2F",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input args
			qa.AssertRequestWithMockServer(t, tt.args, http.MethodGet, tt.wantURI, &input, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewDBFSAPI(client).List(tt.args.Path, tt.args.Recursive)
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
		wantURI        []string
		want           []DBFSFileInfo
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
			wantURI: []string{"/api/2.0/dbfs/list?path=%2F", "/api/2.0/dbfs/list?path=%2Ffoldera"},
			want: []DBFSFileInfo{
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
			wantURI: []string{"/api/2.0/dbfs/list?path=%2F", "/api/2.0/dbfs/list?path=%2Ffoldera"},
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
			wantURI: []string{"/api/2.0/dbfs/list?path=%2F"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qa.AssertMultipleRequestsWithMockServer(t, tt.args, []string{http.MethodGet, http.MethodGet}, tt.wantURI, []interface{}{&args{}}, tt.response, tt.responseStatus, tt.want, tt.wantErr, func(client *common.DatabricksClient) (interface{}, error) {
				return NewDBFSAPI(client).List(tt.args[0].(*args).Path, tt.args[0].(*args).Recursive)
			})
		})
	}
}

func GenString(times int) []byte {
	var buf bytes.Buffer
	for i := 0; i < times; i++ {
		buf.WriteString("Hello world how are you doing?\n")
	}
	return buf.Bytes()
}

func TestAccCreateFile(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}

	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	dir := "/client-test/" + randomName
	dir2 := dir + "/dir2/"
	path := dir + "/randomfile"
	path2 := dir + "/dir2/randomfile"
	path3 := dir + "/dir2/randomfile2"

	randomStr := GenString(500)
	t.Log(len(randomStr))
	t.Log(len(base64.StdEncoding.EncodeToString(randomStr)))

	client := common.NewClientFromEnvironment()

	err := NewDBFSAPI(client).Mkdirs(dir)
	assert.NoError(t, err, err)

	err = NewDBFSAPI(client).Mkdirs(dir2)
	assert.NoError(t, err, err)

	inputData := base64.StdEncoding.EncodeToString(randomStr)
	err = NewDBFSAPI(client).Create(path, true, inputData)
	assert.NoError(t, err, err)

	err = NewDBFSAPI(client).Create(path2, true, inputData)
	assert.NoError(t, err, err)

	err = NewDBFSAPI(client).Create(path3, true, inputData)
	assert.NoError(t, err, err)

	defer func() {
		err := NewDBFSAPI(client).Delete(dir, true)
		assert.NoError(t, err, err)
	}()

	base64Resp, err := NewDBFSAPI(client).Read(path)
	assert.NoError(t, err, err)
	assert.True(t, inputData == base64Resp)

	items, err := NewDBFSAPI(client).List(dir, false)
	assert.NoError(t, err, err)
	assert.Len(t, items, 2)

	items, err = NewDBFSAPI(client).List(dir, true)
	assert.NoError(t, err, err)
	assert.Len(t, items, 3)
}
