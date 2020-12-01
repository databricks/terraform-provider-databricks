package storage

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
)

// func notebookToB64(filePath string) (string, error) {
// 	notebookBytes, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		return "", fmt.Errorf("unable to find notebook to convert to base64; %w", err)
// 	}
// 	return base64.StdEncoding.EncodeToString(notebookBytes), nil
// }
// func getTestDBFSFileData() (string, error) {
// 	return notebookToB64("testdata/tf-test-python.py")
// }

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
		name               string
		fixtures           []qa.HTTPFixture
		source             string
		validateRemoteFile bool
		path               string
		expectedError      string
	}{
		{
			name: "TestDBFSFileCreate_NoMkdirs",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(path),
				getBaseDBFSFileGetStatusFixtures(path, false, false),
				getBaseDBFSFileReadFixtures(path),
			),
			source:             "testdata/tf-test-python.py",
			validateRemoteFile: true,
			path:               path,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_RootDir",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(path),
				getBaseDBFSFileGetStatusFixtures(path, false, false),
				getBaseDBFSFileReadFixtures(path),
			),
			source:             "testdata/tf-test-python.py",
			validateRemoteFile: true,
			path:               path,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_NonRootDir_Exists",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, true, false),
				getBaseDBFSFileCreateFixtures(pathWithDir),
				getBaseDBFSFileGetStatusFixtures(pathWithDir, false, false),
				getBaseDBFSFileReadFixtures(pathWithDir),
			),
			source:             "testdata/tf-test-python.py",
			validateRemoteFile: true,
			path:               pathWithDir,
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
			source:             "testdata/tf-test-python.py",
			validateRemoteFile: true,
			path:               pathWithDir,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_ParentDirIsNotDir",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, false, false),
				getBaseDBFSFileCreateFixtures(pathWithDir),
				getBaseDBFSFileGetStatusFixtures(pathWithDir, false, false),
				getBaseDBFSFileReadFixtures(pathWithDir),
			),
			source:             "testdata/tf-test-python.py",
			validateRemoteFile: true,
			path:               pathWithDir,
			expectedError:      "...",
		},
		{
			name: "TestDBFSFileCreate_NoValidateRemote",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(pathWithDir),
				getBaseDBFSFileGetStatusFixtures(pathWithDir, false, false),
			),
			source:             "testdata/tf-test-python.py",
			validateRemoteFile: false,
			path:               pathWithDir,
			expectedError:      "...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := qa.ResourceFixture{
				Fixtures: tt.fixtures,
				Resource: ResourceDBFSFile(),
				Create:   true,
				State: map[string]interface{}{
					"source":               tt.source,
					"path":                 tt.path,
					"validate_remote_file": tt.validateRemoteFile,
				},
			}.Apply(t)
			if tt.expectedError == "" {
				assert.NoError(t, err, err)
				assert.Equal(t, tt.path, d.Id())
				assert.Equal(t, tt.source, d.Get("content"))
				assert.Equal(t, tt.path, d.Get("path"))
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}

// TODO: via content
// func TestDBFSFileCreate_ViaSource(t *testing.T) {
// 	randomDir := "/abc"
// 	path := "/def"
// 	pathWithDir := randomDir + path
// 	d, err := qa.ResourceFixture{
// 		Fixtures: qa.UnionFixturesLists(
// 			getBaseDBFSFileGetStatusFixtures(randomDir, true, false),
// 			getBaseDBFSFileCreateFixtures(pathWithDir),
// 			getBaseDBFSFileGetStatusFixtures(pathWithDir, false, false),
// 			getBaseDBFSFileReadFixtures(pathWithDir),
// 		),
// 		Resource: ResourceDBFSFile(),
// 		Create:   true,
// 		State: map[string]interface{}{
// 			"source":               source,
// 			"path":                 pathWithDir,
// 			"validate_remote_file": true,
// 		},
// 	}.Apply(t)
// 	assert.NoError(t, err, err)
// 	assert.Equal(t, pathWithDir, d.Id())
// 	assert.Equal(t, source, d.Get("source"))
// 	assert.Equal(t, pathWithDir, d.Get("path"))
// 	assert.Equal(t, sourceMD5, d.Get("content_b64_md5"))
// }

func TestDBFSFileUpdate(t *testing.T) {
	path := "/abc"
	d, err := qa.ResourceFixture{
		Fixtures: qa.UnionFixturesLists(
			getBaseDBFSFileGetStatusFixtures(path, false, false),
			getBaseDBFSFileReadFixtures(path),
		),
		Resource: ResourceDBFSFile(),
		Update:   true,
		ID:       path,
		State: map[string]interface{}{
			"source":               "testdata/tf-test-python.py",
			"path":                 path,
			"validate_remote_file": true,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, true, d.Get("validate_remote_file"))
}

func TestDBFSFileDelete(t *testing.T) {
	path := "/abc"
	d, err := qa.ResourceFixture{
		Fixtures: getBaseDBFSDeleteFixtures(path, false),
		Resource: ResourceDBFSFile(),
		Delete:   true,
		ID:       path,
		State: map[string]interface{}{
			"source":               "testdata/tf-test-python.py",
			"path":                 path,
			"validate_remote_file": true,
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, true, d.Get("validate_remote_file"))
}

func TestDBFSFileRead_IsMissingResource(t *testing.T) {
	path := "/abc"
	d, err := qa.ResourceFixture{
		Fixtures: getBaseDBFSFileGetStatusFixtures(path, false, true),
		Resource: ResourceDBFSFile(),
		Read:     true,
		ID:       path,
		State: map[string]interface{}{
			"source":               "testdata/tf-test-python.py",
			"path":                 path,
			"validate_remote_file": false,
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}
