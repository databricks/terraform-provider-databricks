package databricks

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAwsFolderResource(t *testing.T) {
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest
	//scope := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAwsFolderResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config:  testAwsFolderResourceMultipleFolders(),
				Destroy: false,
			},
		},
	})
}

func TestAccAwsFolderResource_fail_if_already_exists(t *testing.T) {
	// the acctest package includes many helpers such as RandStringFromCharSet
	// See https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/acctest

	if _, ok := os.LookupEnv("TF_ACC"); !ok {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set")
	}

	foldername := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	path := fmt.Sprintf("/tf-test/folder%[1]v/book%[1]v", foldername)
	client := getTokenBasedClient()
	err := client.Notebooks().Mkdirs(path)
	assert.NoError(t, err, err)
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAwsFolderResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config:      testAwsFolderResource(foldername),
				ExpectError: regexp.MustCompile(`.*object already exists in path.*`),
				Destroy:     false,
			},
		},
	})
	err = client.Notebooks().Delete(path, true)
	assert.NoError(t, err, err)
}

func testAwsFolderResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_folder" {
			continue
		}
		_, err := client.Notebooks().Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource Scim Group is not cleaned up")
	}
	return nil
}

func testAwsFolderResource(name string) string {
	return fmt.Sprintf(`
								resource "databricks_folder" "folder_%[1]v" {
								  path = "/tf-test/folder%[1]v/book%[1]v"
								  recursive_delete = "false"
								}
		`, name)
}

func testAwsFolderResourceMultipleFolders() string {
	var strBuffer bytes.Buffer
	for i := 1; i <= 10; i++ {
		strBuffer.WriteString(fmt.Sprintf(`
								resource "databricks_folder" "folder_%[1]v" {
								  path = "/tf-test/folder%[1]v/book%[1]v"
								  recursive_delete = "false"
								}
		`, i))
	}
	return strBuffer.String()
}
