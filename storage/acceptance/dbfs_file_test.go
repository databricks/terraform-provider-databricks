package acceptance

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/storage"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestAccDatabricksDBFSFile_CreateViaContent(t *testing.T) {
	content := acctest.RandString(10000)
	base64Str := base64.StdEncoding.EncodeToString([]byte(content))
	path := "/tmp/tf-test/file-content1"
	// TODO: add random names

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testDBFSFileResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testDBFSFileContentResource(base64Str, 1),
				Destroy: false,
			},
			{
				//Deleting and recreating the token
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewDBFSAPI(client).Delete(path, false)
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
	source := qa.TestCreateTempFile(t, content)
	defer os.Remove(source)

	acceptance.AccTest(t, resource.TestCase{
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
	source := qa.TestCreateTempFile(t, content)
	defer os.Remove(source)
	b64, err := GetLocalFileB64(source)
	assert.NoError(t, err, err)
	md5, err := GetMD5(b64)
	assert.NoError(t, err, err)

	content2 := acctest.RandString(10)
	source2 := qa.TestCreateTempFile(t, content2)
	defer os.Remove(source2)
	source2B64, err := GetLocalFileB64(source2)
	assert.NoError(t, err, err)
	source2Md5, err := GetMD5(source2B64)
	assert.NoError(t, err, err)

	acceptance.AccTest(t, resource.TestCase{
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
		conn := common.CommonEnvironmentClient()
		resp, err := NewDBFSAPI(conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}
		// If no error, assign the response Widget attribute to the widget pointer
		assert.Equal(t, resp, b64)
		// If no error, assign the response Widget attribute to the widget pointer
		respCheckSum, err := GetMD5(resp)
		assert.NoError(t, err, err)
		expectedCheckSum, err := GetMD5(b64)
		assert.NoError(t, err, err)
		assert.Equal(t, expectedCheckSum, respCheckSum)
		return nil
		//return fmt.Errorf("Token (%s) not found", rs.Primary.ID)
	}
}

func testDBFSFileResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_dbfs_file" {
			continue
		}
		_, err := NewDBFSAPI(client).Read(rs.Primary.ID)
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
