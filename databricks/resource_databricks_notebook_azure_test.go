package databricks

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAccAzureNotebookResource_multiple_formats(t *testing.T) {
	folderPrefix := acctest.RandString(10)
	rNotebookDataB64, err := notebookToB64("testdata/tf-test-r.r")
	assert.NoError(t, err, err)
	r2NotebookDataB64, err := notebookToB64("testdata/tf-test-r2.r")
	assert.NoError(t, err, err)
	scalaNotebookDataB64, err := notebookToB64("testdata/tf-test-scala.scala")
	assert.NoError(t, err, err)
	scala2NotebookDataB64, err := notebookToB64("testdata/tf-test-scala2.scala")
	assert.NoError(t, err, err)
	sqlNotebookDataB64, err := notebookToB64("testdata/tf-test-sql.sql")
	assert.NoError(t, err, err)
	sql2NotebookDataB64, err := notebookToB64("testdata/tf-test-sql2.sql")
	assert.NoError(t, err, err)
	dbcNotebookDataB64, err := notebookToB64("testdata/tf-test-dbc.dbc")
	assert.NoError(t, err, err)
	dbc2NotebookDataB64, err := notebookToB64("testdata/tf-test-dbc2.dbc")
	assert.NoError(t, err, err)
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	python2NotebookDataB64, err := notebookToB64("testdata/tf-test-python2.py")
	assert.NoError(t, err, err)

	// Future Support currently will be commented out in tests
	jupyterNotebookDataB64, err := notebookToB64("testdata/tf-test-jupyter.ipynb")
	assert.NoError(t, err, err)
	htmlNotebookDataB64, err := notebookToB64("testdata/tf-test-html.html")
	assert.NoError(t, err, err)

	// Future Support will be commented out in tests
	// Please change the location of file to something else.
	jupyter2NotebookDataB64, err := notebookToB64("testdata/tf-test-jupyter.ipynb")
	assert.NoError(t, err, err)
	html2NotebookDataB64, err := notebookToB64("testdata/tf-test-html.html")
	assert.NoError(t, err, err)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAzureNotebookResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAzureNotebookResourceAllNotebookTypes(
					pythonNotebookDataB64,
					sqlNotebookDataB64,
					scalaNotebookDataB64,
					rNotebookDataB64,
					dbcNotebookDataB64,
					jupyterNotebookDataB64,
					htmlNotebookDataB64,
					folderPrefix,
				),
				Destroy: false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAzureNotebookResourceAllNotebookTypes(
					pythonNotebookDataB64,
					sqlNotebookDataB64,
					scalaNotebookDataB64,
					rNotebookDataB64,
					dbcNotebookDataB64,
					jupyterNotebookDataB64,
					htmlNotebookDataB64,
					folderPrefix,
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAzureNotebookResourceAllNotebookTypes(
					python2NotebookDataB64,
					sql2NotebookDataB64,
					scala2NotebookDataB64,
					r2NotebookDataB64,
					dbc2NotebookDataB64,
					jupyter2NotebookDataB64,
					html2NotebookDataB64,
					folderPrefix,
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testAzureNotebookResourceAllNotebookTypes(
					python2NotebookDataB64,
					sql2NotebookDataB64,
					scala2NotebookDataB64,
					r2NotebookDataB64,
					dbc2NotebookDataB64,
					jupyter2NotebookDataB64,
					html2NotebookDataB64,
					folderPrefix,
				),
				Destroy: false,
			},
		},
	})
}

func TestAccAzureNotebookResource_scalability(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAzureNotebookResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config:  testAzureNotebookResourceMultipleNotebooks(pythonNotebookDataB64),
				Destroy: false,
			},
		},
	})
}

func testAzureNotebookResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*service.DBApiClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_notebook" {
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

func testAzureNotebookResourceAllNotebookTypes(pythonContent, sqlContent, scalaContent, rContent, dbcContent, jupyterContent, htmlContent, folderPrefix string) string {
	return fmt.Sprintf(`
								resource "databricks_notebook" "notebook_python" {
								 content = "%[1]s"
								 path = "/Shared/tf-test-notebooks/%[8]s/python"
								 overwrite = "false"
								 mkdirs = "true"
								 format = "SOURCE"
								 language = "PYTHON"
								}
								resource "databricks_notebook" "notebook_sql" {
								 content = "%[2]s"
								 path = "/Shared/tf-test-notebooks/%[8]s/sql"
								 overwrite = "false"
								 mkdirs = "true"
								 format = "SOURCE"
								 language = "SQL"
								}
								resource "databricks_notebook" "notebook_scala" {
								 content = "%[3]s"
								 path = "/Shared/tf-test-notebooks/%[8]s/scala"
								 overwrite = "false"
								 mkdirs = "true"
								 format = "SOURCE"
								 language = "SCALA"
								}
								resource "databricks_notebook" "notebook_r" {
								 content = "%[4]s"
								 path = "/Shared/tf-test-notebooks/%[8]s/r"
								 overwrite = "false"
								 mkdirs = "true"
								 format = "SOURCE"
								 language = "R"
								}
								resource "databricks_notebook" "notebook_dbc" {
								  content = "%[5]s"
								  path = "/Shared/tf-test-notebooks/%[8]s/dbc"
								  overwrite = "false"
								  mkdirs = "true"
								  format = "DBC"
								}
		`, pythonContent, sqlContent, scalaContent, rContent, dbcContent, jupyterContent, htmlContent, folderPrefix)
}

func testAzureNotebookResourceMultipleNotebooks(content string) string {
	var strBuffer bytes.Buffer
	for i := 1; i <= 10; i++ {
		strBuffer.WriteString(fmt.Sprintf(`
								resource "databricks_notebook" "notebook_%[2]v" {
								  content = "%[1]s"
								  path = "/Shared/tf-test/book%[2]v"
								  overwrite = "false"
								  mkdirs = "true"
								  format = "SOURCE"
								  language = "PYTHON"
								}
		`, content, i))
	}
	return strBuffer.String()
}

func notebookToB64(filePath string) (string, error) {
	notebookBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("unable to find notebook to convert to base64; %w", err)
	}
	return base64.StdEncoding.EncodeToString(notebookBytes), nil
}
