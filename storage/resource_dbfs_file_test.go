package storage

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func notebookToB64(filePath string) (string, error) {
	notebookBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to find notebook to convert to base64; %w", err)
	}
	return base64.StdEncoding.EncodeToString(notebookBytes), nil
}

func getTestDBFSFileData() (string, error) {
	return notebookToB64("testdata/tf-test-python.py")
}

func getBaseDBFSMkdirFixtures(path string) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/mkdirs",
			ExpectedRequest: DBFSMkdirRequest{
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
			ExpectedRequest: DBFSDeleteRequest{
				Path:      path,
				Recursive: recursive,
			},
		},
	}
}

func getBaseDBFSFileGetStatusFixtures(path string, fileSize int64, isDir bool, isMissing bool) []qa.HTTPFixture {
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
			Response: DBFSFileInfo{
				Path:     path,
				IsDir:    isDir,
				FileSize: fileSize,
			},
		},
	}
}

func getBaseDBFSFileReadFixtures(path string, content string, fileSize int64) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/read?length=1000000\u0026path=%s", url.PathEscape(path)),
			Response: DBFSReadResponse{
				BytesRead: fileSize,
				Data:      content,
			},
		},
	}
}

func getBaseDBFSFileCreateFixtures(path string, overwrite bool,
	handleId int64, content string, fileSize int64) []qa.HTTPFixture {
	return []qa.HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: DBFSHandleRequest{
				Path:      path,
				Overwrite: overwrite,
			},
			Response: DBFSHandleResponse{Handle: handleId},
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/add-block",
			Response: DBFSBlockRequest{
				Data:   content,
				Handle: handleId,
			},
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/close",
			Response: DBFSCloseRequest{
				Handle: handleId,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
			Response: DBFSFileInfo{
				Path:     path,
				IsDir:    false,
				FileSize: fileSize,
			},
		},
	}
}

func TestDBFSFileCreate(t *testing.T) {
	handleId := int64(acctest.RandInt())

	randomDir := "/" + acctest.RandString(5)
	randomPath := "/" + acctest.RandString(5)
	randomPathWithDir := randomDir + randomPath
	content, err := getTestDBFSFileData()
	assert.NoError(t, err, err)
	checksum, err := GetMD5(content)
	assert.NoError(t, err, err)
	fileSize := int64(100)

	tests := []struct {
		name               string
		fixtures           []qa.HTTPFixture
		content            string
		contentB64MD5      string
		mkdirs             bool
		overwrite          bool
		fileSize           int64
		validateRemoteFile bool
		path               string
		expectedError      error
	}{
		{
			name: "TestDBFSFileCreate_NoMkdirs",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(randomPath, true, handleId, content, fileSize),
				getBaseDBFSFileGetStatusFixtures(randomPath, fileSize, false, false),
				getBaseDBFSFileReadFixtures(randomPath, content, fileSize),
			),
			content:            content,
			mkdirs:             false,
			contentB64MD5:      checksum,
			overwrite:          true,
			fileSize:           fileSize,
			validateRemoteFile: true,
			path:               randomPath,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_RootDir",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(randomPath, true, handleId, content, fileSize),
				getBaseDBFSFileGetStatusFixtures(randomPath, fileSize, false, false),
				getBaseDBFSFileReadFixtures(randomPath, content, fileSize),
			),
			content:            content,
			contentB64MD5:      checksum,
			mkdirs:             true,
			overwrite:          true,
			fileSize:           fileSize,
			validateRemoteFile: true,
			path:               randomPath,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_NonRootDir_Exists",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, fileSize, true, false),
				getBaseDBFSFileCreateFixtures(randomPathWithDir, true, handleId, content, fileSize),
				getBaseDBFSFileGetStatusFixtures(randomPathWithDir, fileSize, false, false),
				getBaseDBFSFileReadFixtures(randomPathWithDir, content, fileSize),
			),
			content:            content,
			contentB64MD5:      checksum,
			mkdirs:             true,
			overwrite:          true,
			fileSize:           fileSize,
			validateRemoteFile: true,
			path:               randomPathWithDir,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_NonRootDir_DoesNotExist",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, fileSize, true, true),
				getBaseDBFSMkdirFixtures(randomDir),
				getBaseDBFSFileCreateFixtures(randomPathWithDir, true, handleId, content, fileSize),
				getBaseDBFSFileGetStatusFixtures(randomPathWithDir, fileSize, false, false),
				getBaseDBFSFileReadFixtures(randomPathWithDir, content, fileSize),
			),
			content:            content,
			contentB64MD5:      checksum,
			mkdirs:             true,
			overwrite:          true,
			fileSize:           fileSize,
			validateRemoteFile: true,
			path:               randomPathWithDir,
		},
		{
			name: "TestDBFSFileCreate_Mkdirs_ParentDirIsNotDir",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileGetStatusFixtures(randomDir, fileSize, false, false),
				getBaseDBFSFileCreateFixtures(randomPathWithDir, true, handleId, content, fileSize),
				getBaseDBFSFileGetStatusFixtures(randomPathWithDir, fileSize, false, false),
				getBaseDBFSFileReadFixtures(randomPathWithDir, content, fileSize),
			),
			content:            content,
			contentB64MD5:      checksum,
			mkdirs:             true,
			overwrite:          true,
			fileSize:           fileSize,
			validateRemoteFile: true,
			path:               randomPathWithDir,
			expectedError:      ParentPathIsFileError,
		},
		{
			name: "TestDBFSFileCreate_NoValidateRemote",
			fixtures: qa.UnionFixturesLists(
				getBaseDBFSFileCreateFixtures(randomPathWithDir, true, handleId, content, fileSize),
				getBaseDBFSFileGetStatusFixtures(randomPathWithDir, fileSize, false, false),
			),
			content:            content,
			contentB64MD5:      checksum,
			mkdirs:             false,
			overwrite:          true,
			fileSize:           fileSize,
			validateRemoteFile: false,
			path:               randomPathWithDir,
			expectedError:      ParentPathIsFileError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := qa.ResourceFixture{
				Fixtures: tt.fixtures,
				Resource: ResourceDBFSFile(),
				Create:   true,
				State: map[string]interface{}{
					"content":              tt.content,
					"content_b64_md5":      tt.contentB64MD5,
					"path":                 tt.path,
					"overwrite":            tt.overwrite,
					"mkdirs":               tt.mkdirs,
					"validate_remote_file": tt.validateRemoteFile,
				},
			}.Apply(t)
			if tt.expectedError == nil {
				assert.NoError(t, err, err)
				assert.Equal(t, tt.path, d.Id())
				assert.Equal(t, tt.content, d.Get("content"))
				assert.Equal(t, tt.path, d.Get("path"))
				assert.Equal(t, int(tt.fileSize), d.Get("file_size"))
			} else {
				assert.Error(t, tt.expectedError, err)
			}
		})
	}
}

func TestDBFSFileCreate_ViaSource(t *testing.T) {
	handleId := int64(acctest.RandInt())
	randomDir := "/" + acctest.RandString(5)
	randomPath := "/" + acctest.RandString(5)
	randomPathWithDir := randomDir + randomPath
	content, err := getTestDBFSFileData()
	assert.NoError(t, err, err)
	source := qa.TestCreateTempFile(t, content)
	defer os.Remove(source)
	sourceB64, err := GetLocalFileB64(source)
	assert.NoError(t, err, err)
	sourceMD5, err := GetMD5(sourceB64)
	assert.NoError(t, err, err)
	fileSize := int64(100)

	d, err := qa.ResourceFixture{
		Fixtures: qa.UnionFixturesLists(
			getBaseDBFSFileGetStatusFixtures(randomDir, fileSize, true, false),
			getBaseDBFSFileCreateFixtures(randomPathWithDir, true, handleId, sourceB64, fileSize),
			getBaseDBFSFileGetStatusFixtures(randomPathWithDir, fileSize, false, false),
			getBaseDBFSFileReadFixtures(randomPathWithDir, sourceB64, fileSize),
		),
		Resource: ResourceDBFSFile(),
		Create:   true,
		State: map[string]interface{}{
			"source":               source,
			"content_b64_md5":      sourceMD5,
			"path":                 randomPathWithDir,
			"overwrite":            true,
			"mkdirs":               true,
			"validate_remote_file": true,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, randomPathWithDir, d.Id())
	assert.Equal(t, source, d.Get("source"))
	assert.Equal(t, randomPathWithDir, d.Get("path"))
	assert.Equal(t, sourceMD5, d.Get("content_b64_md5"))
}

func TestDBFSFileUpdate(t *testing.T) {
	randomPath := "/" + acctest.RandString(5)
	content, err := getTestDBFSFileData()
	assert.NoError(t, err, err)
	checksum, err := GetMD5(content)
	assert.NoError(t, err, err)
	fileSize := int64(100)
	d, err := qa.ResourceFixture{
		Fixtures: qa.UnionFixturesLists(
			getBaseDBFSFileGetStatusFixtures(randomPath, fileSize, false, false),
			getBaseDBFSFileReadFixtures(randomPath, content, fileSize),
		),
		Resource: ResourceDBFSFile(),
		Update:   true,
		ID:       randomPath,
		State: map[string]interface{}{
			"content":              content,
			"content_b64_md5":      checksum,
			"path":                 randomPath,
			"overwrite":            true,
			"mkdirs":               true,
			"validate_remote_file": true,
		},
	}.Apply(t)
	assert.NoError(t, err, err)
	assert.Equal(t, randomPath, d.Id())
	assert.Equal(t, true, d.Get("overwrite"))
	assert.Equal(t, true, d.Get("mkdirs"))
	assert.Equal(t, true, d.Get("validate_remote_file"))
}

func TestDBFSFileDelete(t *testing.T) {
	content, err := getTestDBFSFileData()
	assert.NoError(t, err, err)
	checksum, err := GetMD5(content)
	assert.NoError(t, err, err)
	randomPath := "/" + acctest.RandString(5)
	d, err := qa.ResourceFixture{
		Fixtures: getBaseDBFSDeleteFixtures(randomPath, false),
		Resource: ResourceDBFSFile(),
		Delete:   true,
		ID:       randomPath,
		State: map[string]interface{}{
			"content":              content,
			"content_b64_md5":      checksum,
			"path":                 randomPath,
			"overwrite":            true,
			"mkdirs":               true,
			"validate_remote_file": true,
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, randomPath, d.Id())
	assert.Equal(t, true, d.Get("overwrite"))
	assert.Equal(t, true, d.Get("mkdirs"))
	assert.Equal(t, true, d.Get("validate_remote_file"))
}

func TestDBFSFileRead_IsMissingResource(t *testing.T) {
	randomPath := "/" + acctest.RandString(5)
	content, err := getTestDBFSFileData()
	assert.NoError(t, err, err)
	checksum, err := GetMD5(content)
	assert.NoError(t, err, err)
	fileSize := int64(100)
	d, err := qa.ResourceFixture{
		Fixtures: getBaseDBFSFileGetStatusFixtures(randomPath, fileSize, false, true),
		Resource: ResourceDBFSFile(),
		Read:     true,
		ID:       randomPath,
		State: map[string]interface{}{
			"content":              content,
			"content_b64_md5":      checksum,
			"path":                 randomPath,
			"overwrite":            true,
			"mkdirs":               true,
			"validate_remote_file": false,
		},
	}.Apply(t)

	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestDatabricksFile_Base64(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	client := common.NewClientFromEnvironment()
	pythonNotebookDataB64, err := GetLocalFileB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	expected, err := GetMD5(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	t.Log(expected)
	t.Log(pythonNotebookDataB64)
	err = NewDBFSAPI(client).Create("/tmp/tf-test/testfile.txt", true, pythonNotebookDataB64)
	assert.NoError(t, err, err)
	data, err := NewDBFSAPI(client).Read("/tmp/tf-test/testfile.txt")
	t.Log(data)
	assert.NoError(t, err, err)
	actual, err := GetMD5(data)
	assert.NoError(t, err, err)
	assert.Equal(t, expected, actual)
}
