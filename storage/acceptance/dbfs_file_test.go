package acceptance

import (
	"errors"
	"fmt"
	"os"
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
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_dbfs_file" "file" {
		content = base64encode("{var.RANDOM}")
		content_b64_md5 = md5(base64encode("{var.RANDOM}"))
		path = "/tmp/tf-test/file-content-{var.RANDOM}"
		overwrite = false
		mkdirs = true
		validate_remote_file = true
	}`)
	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testDBFSFileResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:  config,
				Destroy: false,
			},
			{
				//Deleting and recreating the token
				PreConfig: func() {
					client := common.CommonEnvironmentClient()
					err := NewDBFSAPI(client).Delete(qa.FirstKeyValue(t, config, "path"), false)
					assert.NoError(t, err, err)
				},
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config:  config,
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
				Config: qa.EnvironmentTemplate(t, `
				resource "databricks_dbfs_file" "file_1" {
					source = "{var.SOURCE}"
					content_b64_md5 = md5(filebase64("{var.SOURCE}"))
					path = "/tmp/tf-test/file-source-{var.RANDOM}"
					overwrite = "false"
					mkdirs = "true"
					validate_remote_file = "true"
				}`, map[string]string{
					"SOURCE": source,
				}),
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testCheckDBFSFileResourceExists("databricks_dbfs_file.file_1", b64, t),
					resource.TestCheckResourceAttr("databricks_dbfs_file.file_1", "content_b64_md5", md5),
				),
				Destroy: false,
			},
			{
				Config: qa.EnvironmentTemplate(t, `
				resource "databricks_dbfs_file" "file_1" {
					source = "{var.SOURCE}"
					content_b64_md5 = md5(filebase64("{var.SOURCE}"))
					path = "/tmp/tf-test/file-source-{var.RANDOM}"
					overwrite = "false"
					mkdirs = "true"
					validate_remote_file = "true"
				}`, map[string]string{
					"SOURCE": source2,
				}),
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
