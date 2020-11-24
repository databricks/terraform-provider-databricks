package acceptance

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/acceptance"
	"github.com/databrickslabs/databricks-terraform/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/stretchr/testify/assert"
)

func TestAccNotebookResourceScalability(t *testing.T) {
	// TODO: refactor for common instance pool & AZ CLI
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testNotebookResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config:  testNotebookResourceMultipleNotebooks(pythonNotebookDataB64),
				Destroy: false,
			},
		},
	})
}

func TestAccNotebookResourceMultipleFormats(t *testing.T) {
	if _, ok := os.LookupEnv("CLOUD_ENV"); !ok {
		t.Skip("Acceptance tests skipped unless env 'CLOUD_ENV' is set")
	}
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
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	python2NotebookDataB64, err := notebookToB64("testdata/tf-test-python2.py")
	assert.NoError(t, err, err)

	// Potentially Future Support currently will be commented out in tests
	jupyterNotebookDataB64, err := notebookToB64("testdata/tf-test-jupyter.ipynb")
	assert.NoError(t, err, err)
	htmlNotebookDataB64, err := notebookToB64("testdata/tf-test-html.html")
	assert.NoError(t, err, err)
	dbcNotebookDataB64, err := notebookToB64("testdata/tf-test-dbc.dbc")
	assert.NoError(t, err, err)
	dbc2NotebookDataB64, err := notebookToB64("testdata/tf-test-dbc2.dbc")
	assert.NoError(t, err, err)

	acceptance.AccTest(t, resource.TestCase{
		CheckDestroy: testNotebookResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testNotebookResourceAllNotebookTypes(
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
				Config: testNotebookResourceAllNotebookTypes(
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
				Config: testNotebookResourceAllNotebookTypes(
					python2NotebookDataB64,
					sql2NotebookDataB64,
					scala2NotebookDataB64,
					r2NotebookDataB64,
					dbc2NotebookDataB64,
					jupyterNotebookDataB64,
					htmlNotebookDataB64,
					folderPrefix,
				),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				Destroy:            false,
			},
			{
				// use a dynamic configuration with the random name from above
				Config: testNotebookResourceAllNotebookTypes(
					python2NotebookDataB64,
					sql2NotebookDataB64,
					scala2NotebookDataB64,
					r2NotebookDataB64,
					dbc2NotebookDataB64,
					jupyterNotebookDataB64,
					htmlNotebookDataB64,
					folderPrefix,
				),
				Destroy: false,
			},
		},
	})
}

func testNotebookResourceDestroy(s *terraform.State) error {
	client := common.CommonEnvironmentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "databricks_notebook" {
			continue
		}
		_, err := workspace.NewNotebooksAPI(context.Background(), client).Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return errors.New("resource notebook is not cleaned up")
	}
	return nil
}

// Only will test source format based content
func testNotebookResourceAllNotebookTypes(pythonContent, sqlContent, scalaContent,
	rContent, dbcContent, jupyterContent, htmlContent, folderPrefix string) string {
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
		`, pythonContent, sqlContent, scalaContent, rContent, dbcContent, jupyterContent, htmlContent, folderPrefix)
}

func testNotebookResourceMultipleNotebooks(content string) string {
	var strBuffer bytes.Buffer
	randomName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	for i := 1; i <= 10; i++ {
		strBuffer.WriteString(fmt.Sprintf(`
			resource "databricks_notebook" "notebook_%[2]v" {
				content = "%[1]s"
				path = "/Shared/tf-test-%[3]s/book%[2]v"
				overwrite = "false"
				mkdirs = "true"
				format = "SOURCE"
				language = "PYTHON"
			}`, content, i, randomName))
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
