package acceptance

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	. "github.com/databrickslabs/databricks-terraform/storage"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/stretchr/testify/assert"
)

func TestAccDatabricksDBFSFile_CreateViaContent(t *testing.T) {
	config := qa.EnvironmentTemplate(t, `
	resource "databricks_dbfs_file" "file" {
		content = base64encode("{var.RANDOM}")
		content_b64_md5 = md5(base64encode("{var.RANDOM}"))
		path = "/tmp/tf-test/file-content-{var.RANDOM}"
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
					err := NewDbfsAPI(context.Background(), client).Delete(qa.FirstKeyValue(t, config, "path"), false)
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

	content2 := acctest.RandString(10)
	source2 := qa.TestCreateTempFile(t, content2)
	defer os.Remove(source2)

	// TODO: fix tests
	var b64, md5, b642, md52 string

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testDBFSFileResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: qa.EnvironmentTemplate(t, `
				resource "databricks_dbfs_file" "file_1" {
					source = "{var.SOURCE}"
					content_b64_md5 = md5(filebase64("{var.SOURCE}"))
					path = "/tmp/tf-test/file-source-{var.RANDOM}"
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
				}`, map[string]string{
					"SOURCE": source2,
				}),
				Check: resource.ComposeTestCheckFunc(
					// query the API to retrieve the tokenInfo object
					testCheckDBFSFileResourceExists("databricks_dbfs_file.file_1", b642, t),
					resource.TestCheckResourceAttr("databricks_dbfs_file.file_1", "content_b64_md5", md52),
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
		_, err := NewDbfsAPI(context.Background(), conn).Read(rs.Primary.ID)
		if err != nil {
			return err
		}
		// TODO: rewrite
		// assert.Equal(t, resp, b64)
		// respCheckSum, err := md5.Sum(resp)
		// assert.NoError(t, err, err)
		// expectedCheckSum, err := GetMD5(b64)
		// assert.NoError(t, err, err)
		// assert.Equal(t, expectedCheckSum, respCheckSum)
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
		_, err := NewDbfsAPI(context.Background(), client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource dbfs file is not cleaned up")
	}
	return nil
}
