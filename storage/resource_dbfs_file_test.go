package storage

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func getBaseDBFSMkdirFixtures(path string) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/mkdirs",
			ExpectedRequest: dbfsRequest{
				Path: path,
			},
		},
	}
}

func getBaseDBFSDeleteFixtures(path string, recursive bool) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/delete",
			ExpectedRequest: dbfsRequest{
				Path:      path,
				Recursive: recursive,
			},
		},
	}
}

func getBaseDBFSFileGetStatusFixtures(path string, isDir bool, isMissing bool) []qa.HTTPFixture {
	if isMissing {
		return []qa.HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
				Status:   http.StatusNotFound,
			},
		}
	}
	return []qa.HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
			Response: FileInfo{
				Path:     path,
				IsDir:    isDir,
				FileSize: 1024,
			},
		},
	}
}

func getBaseDBFSFileReadFixtures(path string) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/read?length=1000000&path=%s", url.PathEscape(path)),
			Response: ReadResponse{
				BytesRead: 1024,
				Data:      "...",
			},
		},
	}
}

func getBaseDBFSFileCreateFixtures(path string) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: CreateHandle{
				Path:      path,
				Overwrite: true,
			},
			Response: Handle{329874298374132},
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/add-block",
			// Response: DBFSBlockRequest{
			// 	Data:   source,
			// 	Handle: handleId,
			// },
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/close",
			Response: Handle{329874298374132},
		},
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
			Response: FileInfo{
				Path:     path,
				IsDir:    false,
				FileSize: 1024,
			},
		},
	}
}

func TestDBFSFileCreate(t *testing.T) {
	randomDir := "/abc"
	path := "/def"
	pathWithDir := randomDir + path

	tests := []struct {
		name          string
		fixtures      []qa.HTTPFixture
		source        string
		path          string
		expectedError string
	}{
		{
			name: "TestDBFSFileCreate_NoMkdirs",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(path),
				getBaseDBFSFileGetStatusFixtures(path, false, false),
				getBaseDBFSFileReadFixtures(path),
			),
			source: "testdata/tf-test-python.py",
			path:   path,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_RootDir",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(path),
				getBaseDBFSFileGetStatusFixtures(path, false, false),
				getBaseDBFSFileReadFixtures(path),
			),
			source: "testdata/tf-test-python.py",
			path:   path,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_NonRootDir_Exists",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, true, false),
				getBaseDBFSFileCreateFixtures(pathWithDir),
				getBaseDBFSFileGetStatusFixtures(pathWithDir, false, false),
				getBaseDBFSFileReadFixtures(pathWithDir),
			),
			source: "testdata/tf-test-python.py",
			path:   pathWithDir,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_NonRootDir_DoesNotExist",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, true, true),
				getBaseDBFSMkdirFixtures(randomDir),
				getBaseDBFSFileCreateFixtures(pathWithDir),
				getBaseDBFSFileGetStatusFixtures(pathWithDir, false, false),
				getBaseDBFSFileReadFixtures(pathWithDir),
			),
			source: "testdata/tf-test-python.py",
			path:   pathWithDir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := qa.ResourceFixture{
				Fixtures: tt.fixtures,
				Resource: ResourceDBFSFile(),
				Create:   true,
				State: map[string]interface{}{
					"source": tt.source,
					"path":   tt.path,
				},
			}.Apply(t)
			if tt.expectedError == "" {
				assert.NoError(t, err, err)
				assert.Equal(t, tt.path, d.Id())
				assert.Equal(t, tt.source, d.Get("source"))
				assert.Equal(t, tt.path, d.Get("path"))
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}

func TestDBFSFileDelete(t *testing.T) {
	path := "/abc"
	d, err := qa.ResourceFixture{
		Fixtures: getBaseDBFSDeleteFixtures(path, false),
		Resource: ResourceDBFSFile(),
		Delete:   true,
		ID:       path,
		State: map[string]interface{}{
			"source": "testdata/tf-test-python.py",
			"path":   path,
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
}

func TestDBFSFileRead_IsMissingResource(t *testing.T) {
	path := "/abc"
	qa.ResourceFixture{
		Fixtures: getBaseDBFSFileGetStatusFixtures(path, false, true),
		Resource: ResourceDBFSFile(),
		Read:     true,
		ID:       path,
		Removed:  true,
		State: map[string]interface{}{
			"source": "testdata/tf-test-python.py",
			"path":   path,
		},
	}.ApplyNoError(t)
}
