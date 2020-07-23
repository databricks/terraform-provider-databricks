package databricks

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func getTestDBFSFileData() (string, error) {
	return notebookToB64("testdata/tf-test-python.py")
}

func getBaseDBFSMkdirFixtures(path string) []HTTPFixture {
	return []HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/mkdirs",
			ExpectedRequest: model.DBFSMkdirRequest{
				Path: path,
			},
		},
	}
}

func getBaseDBFSDeleteFixtures(path string, recursive bool) []HTTPFixture {
	return []HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/delete",
			ExpectedRequest: model.DBFSDeleteRequest{
				Path:      path,
				Recursive: recursive,
			},
		},
	}
}

func getBaseDBFSFileGetStatusFixtures(path string, fileSize int64, isDir bool, isMissing bool) []HTTPFixture {
	if isMissing {
		return []HTTPFixture{
			{
				Method:   http.MethodGet,
				Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
				Status:   http.StatusNotFound,
			},
		}
	}
	return []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
			Response: model.DBFSFileInfo{
				Path:     path,
				IsDir:    isDir,
				FileSize: fileSize,
			},
		},
	}
}

func getBaseDBFSFileReadFixtures(path string, content string, fileSize int64) []HTTPFixture {
	return []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/read?length=1000000\u0026path=%s", url.PathEscape(path)),
			Response: model.DBFSReadResponse{
				BytesRead: fileSize,
				Data:      content,
			},
		},
	}
}

func getBaseDBFSFileCreateFixtures(path string, overwrite bool, handleId int64, content string, fileSize int64) []HTTPFixture {
	return []HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/create",
			ExpectedRequest: model.DBFSHandleRequest{
				Path:      path,
				Overwrite: overwrite,
			},
			Response: model.DBFSHandleResponse{Handle: handleId},
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/add-block",
			Response: model.DBFSBlockRequest{
				Data:   content,
				Handle: handleId,
			},
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/dbfs/close",
			Response: model.DBFSCloseRequest{
				Handle: handleId,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: fmt.Sprintf("/api/2.0/dbfs/get-status?path=%s", url.PathEscape(path)),
			Response: model.DBFSFileInfo{
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
	checksum, err := getMD5(content)
	assert.NoError(t, err, err)
	fileSize := int64(100)

	tests := []struct {
		name               string
		fixtures           []HTTPFixture
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
			fixtures: UnionFixturesLists(
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
			fixtures: UnionFixturesLists(
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
			fixtures: UnionFixturesLists(
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
			fixtures: UnionFixturesLists(
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
			fixtures: UnionFixturesLists(
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
			fixtures: UnionFixturesLists(
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
			d, err := ResourceTester(t, tt.fixtures, resourceDBFSFile, map[string]interface{}{
				"content":              tt.content,
				"content_b64_md5":      tt.contentB64MD5,
				"path":                 tt.path,
				"overwrite":            tt.overwrite,
				"mkdirs":               tt.mkdirs,
				"validate_remote_file": tt.validateRemoteFile,
			}, resourceDBFSFileCreate)
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
	source := testCreateTempFile(t, content)
	defer os.Remove(source)
	sourceB64, err := getLocalFileB64(source)
	assert.NoError(t, err, err)
	sourceMD5, err := getMD5(sourceB64)
	assert.NoError(t, err, err)
	fileSize := int64(100)

	d, err := ResourceTester(t, UnionFixturesLists(
		getBaseDBFSFileGetStatusFixtures(randomDir, fileSize, true, false),
		getBaseDBFSFileCreateFixtures(randomPathWithDir, true, handleId, sourceB64, fileSize),
		getBaseDBFSFileGetStatusFixtures(randomPathWithDir, fileSize, false, false),
		getBaseDBFSFileReadFixtures(randomPathWithDir, sourceB64, fileSize),
	), resourceDBFSFile, map[string]interface{}{
		"source":               source,
		"content_b64_md5":      sourceMD5,
		"path":                 randomPathWithDir,
		"overwrite":            true,
		"mkdirs":               true,
		"validate_remote_file": true,
	}, resourceDBFSFileCreate)
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
	checksum, err := getMD5(content)
	assert.NoError(t, err, err)
	fileSize := int64(100)
	fixtures := UnionFixturesLists(
		getBaseDBFSFileGetStatusFixtures(randomPath, fileSize, false, false),
		getBaseDBFSFileReadFixtures(randomPath, content, fileSize),
	)
	d, err := ResourceTester(t, fixtures, resourceDBFSFile, map[string]interface{}{
		"content":              content,
		"content_b64_md5":      checksum,
		"path":                 randomPath,
		"overwrite":            true,
		"mkdirs":               true,
		"validate_remote_file": true,
	}, func(d *schema.ResourceData, c interface{}) error {
		d.SetId(randomPath)
		return resourceDBFSFileUpdate(d, c)
	})

	assert.NoError(t, err, err)
	assert.Equal(t, randomPath, d.Id())
	assert.Equal(t, true, d.Get("overwrite"))
	assert.Equal(t, true, d.Get("mkdirs"))
	assert.Equal(t, true, d.Get("validate_remote_file"))
}

func TestDBFSFileDelete(t *testing.T) {
	content, err := getTestDBFSFileData()
	assert.NoError(t, err, err)
	checksum, err := getMD5(content)
	assert.NoError(t, err, err)
	randomPath := "/" + acctest.RandString(5)
	fixtures := UnionFixturesLists(
		getBaseDBFSDeleteFixtures(randomPath, false),
	)
	d, err := ResourceTester(t, fixtures, resourceDBFSFile, map[string]interface{}{
		"content":              content,
		"content_b64_md5":      checksum,
		"path":                 randomPath,
		"overwrite":            true,
		"mkdirs":               true,
		"validate_remote_file": true,
	}, func(d *schema.ResourceData, c interface{}) error {
		d.SetId(randomPath)
		return resourceDBFSFileDelete(d, c)
	})

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
	checksum, err := getMD5(content)
	assert.NoError(t, err, err)
	fileSize := int64(100)
	fixtures := UnionFixturesLists(
		getBaseDBFSFileGetStatusFixtures(randomPath, fileSize, false, true),
	)
	d, err := ResourceTester(t, fixtures, resourceDBFSFile, map[string]interface{}{
		"content":              content,
		"content_b64_md5":      checksum,
		"path":                 randomPath,
		"overwrite":            true,
		"mkdirs":               true,
		"validate_remote_file": false,
	}, func(d *schema.ResourceData, c interface{}) error {
		d.SetId(randomPath)
		return resourceDBFSFileRead(d, c)
	})

	assert.NoError(t, err, err)
	assert.Equal(t, "", d.Id())
}

func TestAccDatabricksDBFSFile_CreateViaContent(t *testing.T) {
	content := acctest.RandString(10000)
	base64Str := base64.StdEncoding.EncodeToString([]byte(content))
	path := "/tmp/tf-test/file-content1"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDBFSFileResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testDBFSFileContentResource(base64Str, 1),
				Destroy: false,
			},
			{
				//Deleting and recreating the token
				PreConfig: func() {
					client := testAccProvider.Meta().(*service.DatabricksClient)
					err := client.DBFS().Delete(path, false)
					assert.NoError(t, err, err)
				},
				Config:             testDBFSFileContentResource(base64Str, 1),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:  testDBFSFileContentResource(base64Str, 1),
				Destroy: false,
			},
		},
	})
}

func TestAccDatabricksDBFSFile_CreateVeryBigFiles(t *testing.T) {
	// Creating via content this will fail, needs to be uploaded via source
	content := acctest.RandString(5000000)
	base64Str := base64.StdEncoding.EncodeToString([]byte(content))
	source := testCreateTempFile(t, content)
	defer os.Remove(source)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDBFSFileResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testDBFSFileContentResource(base64Str, 1),
				Destroy:     false,
				ExpectError: regexp.MustCompile(`.*rpc error.*`),
			},
			{
				Config:  testDBFSFileSourceResource(source, 1),
				Destroy: false,
			},
		},
	})
}

func TestAccDatabricksDBFSFile_CreateViaSource(t *testing.T) {
	content := acctest.RandString(10)
	source := testCreateTempFile(t, content)
	defer os.Remove(source)
	b64, err := getLocalFileB64(source)
	assert.NoError(t, err, err)
	md5, err := getMD5(b64)
	assert.NoError(t, err, err)

	content2 := acctest.RandString(10)
	source2 := testCreateTempFile(t, content2)
	defer os.Remove(source2)
	source2B64, err := getLocalFileB64(source2)
	assert.NoError(t, err, err)
	source2Md5, err := getMD5(source2B64)
	assert.NoError(t, err, err)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDBFSFileResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDBFSFileSourceResource(source, 1),
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testCheckDBFSFileResourceExists("databricks_dbfs_file.file_1", b64, t),
					resource.TestCheckResourceAttr("databricks_dbfs_file.file_1", "content_b64_md5", md5),
				),
				Destroy: false,
			},
			{
				Config: testDBFSFileSourceResource(source2, 1),
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testCheckDBFSFileResourceExists("databricks_dbfs_file.file_1", source2B64, t),
					resource.TestCheckResourceAttr("databricks_dbfs_file.file_1", "content_b64_md5", source2Md5),
				),
				Destroy: false,
			},
		},
	})
}

// testAccAzureCheckTokenResourceExists queries the API and retrieves the matching Widget.
func testCheckDBFSFileResourceExists(n string, b64 string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// find the corresponding state object
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		// retrieve the configured client from the test setup
		conn := testAccProvider.Meta().(*service.DatabricksClient)
		resp, err := conn.DBFS().Read(rs.Primary.ID)
		if err != nil {
			return err
		}
		// If no error, assign the response Widget attribute to the widget pointer
		assert.Equal(t, resp, b64)
		// If no error, assign the response Widget attribute to the widget pointer
		respCheckSum, err := getMD5(resp)
		assert.NoError(t, err, err)
		expectedCheckSum, err := getMD5(b64)
		assert.NoError(t, err, err)
		assert.Equal(t, expectedCheckSum, respCheckSum)
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

func testDBFSFileResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DatabricksClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_dbfs_file" {
			continue
		}
		_, err := client.DBFS().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource dbfs file is not cleaned up")
	}
	return nil
}

func testDBFSFileContentResource(content string, count int) string {
	var strBuffer bytes.Buffer
	for i := 1; i <= count; i++ {
		strBuffer.WriteString(fmt.Sprintf(`
								resource "databricks_dbfs_file" "file_%[2]v" {
								  content = "%[1]s"
								  content_b64_md5 = md5("%[1]s")
								  path = "/tmp/tf-test/file-content%[2]v"
								  overwrite = "false"
								  mkdirs = "true"
								  validate_remote_file = "true"
								}
		`, content, i))
	}
	return strBuffer.String()
}

func testDBFSFileSourceResource(source string, count int) string {
	var strBuffer bytes.Buffer
	for i := 1; i <= count; i++ {
		strBuffer.WriteString(fmt.Sprintf(`
								resource "databricks_dbfs_file" "file_%[2]v" {
								  source = "%[1]s"
								  content_b64_md5 = md5(filebase64("%[1]s"))
								  path = "/tmp/tf-test/file-source%[2]v"
								  overwrite = "false"
								  mkdirs = "true"
								  validate_remote_file = "true"
								}
		`, source, i))
	}
	return strBuffer.String()
}

func testCreateTempFile(t *testing.T, data string) string {
	tmpFile, err := ioutil.TempFile("", "tf-test-create-dbfs-file")
	if err != nil {
		t.Fatal(err)
	}
	filename := tmpFile.Name()

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		os.Remove(filename)
		t.Fatal(err)
	}

	return filename
}

func TestDatabricksFile_Base64(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	client := service.NewClientFromEnvironment()
	pythonNotebookDataB64, err := getLocalFileB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	expected, err := getMD5(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	t.Log(expected)
	t.Log(pythonNotebookDataB64)
	err = client.DBFS().Create("/tmp/tf-test/testfile.txt", true, pythonNotebookDataB64)
	assert.NoError(t, err, err)
	data, err := client.DBFS().Read("/tmp/tf-test/testfile.txt")
	t.Log(data)
	assert.NoError(t, err, err)
	actual, err := getMD5(data)
	assert.NoError(t, err, err)
	assert.Equal(t, expected, actual)
}
