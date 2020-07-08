package databricks

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/databrickslabs/databricks-terraform/client/model"
	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/stretchr/testify/assert"
)

func TestDatabricksNotebook_GetDirPath(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		expectedDirPath string
		expectedError   error
	}{
		{
			name:            "basic_path",
			path:            "/test/abc/file.py",
			expectedDirPath: "/test/abc",
			expectedError:   nil,
		},
		{
			name:            "root_path",
			path:            "/file.py",
			expectedDirPath: "",
			expectedError:   notebookDirPathRootDirError,
		},
		{
			name:            "empty_path",
			path:            "",
			expectedDirPath: "",
			expectedError:   notebookPathEmptyError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dirPath, err := getNotebookParentDirPath(tt.path)
			assert.Equal(t, tt.expectedDirPath, dirPath, "dirPath values should match")
			assert.Equal(t, tt.expectedError, err, "err values should match")
		})
	}
}

func TestValidateNotebookPath(t *testing.T) {
	testCases := []struct {
		name         string
		notebookPath string
		errorCount   int
	}{
		{"empty_path",
			"",
			2},
		{"correct_path",
			"/directory",
			0},
		{"path_starts_with_no_slash",
			"directory",
			1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, errs := ValidateNotebookPath(tc.notebookPath, "key")

			assert.Lenf(t, errs, tc.errorCount, "directory '%s' does not generate the expected error count", tc.notebookPath)
		})
	}
}

func TestNotebooksCreate_DirDoesNotExists(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	path := "/test/path.py"
	content := pythonNotebookDataB64
	objectId := 12345

	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/get-status?path=%2Ftest",
			Response: service.APIErrorBody{
				ErrorCode: "NOT_FOUND",
				Message:   "not found",
			},
			Status: 404,
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/workspace/mkdirs",
			Response: model.NotebookImportRequest{
				Content:   content,
				Path:      path,
				Language:  model.Python,
				Overwrite: true,
				Format:    model.Source,
			},
		},
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/workspace/import",
			Response: model.NotebookImportRequest{
				Content:   content,
				Path:      path,
				Language:  model.Python,
				Overwrite: true,
				Format:    model.Source,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
			Response: model.NotebookContent{
				Content: pythonNotebookDataB64,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
			Response: model.WorkspaceObjectStatus{
				ObjectID:   int64(objectId),
				ObjectType: model.Notebook,
				Path:       path,
				Language:   model.Python,
			},
		},
	}, resourceNotebook, map[string]interface{}{
		"path":      path,
		"content":   content,
		"language":  string(model.Python),
		"format":    string(model.Source),
		"overwrite": true,
		"mkdirs":    true,
	}, resourceNotebookCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, string(model.Python), d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestNotebooksCreate_NoMkdirs(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	path := "/test/path.py"
	content := pythonNotebookDataB64
	objectId := 12345

	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodPost,
			Resource: "/api/2.0/workspace/import",
			Response: model.NotebookImportRequest{
				Content:   content,
				Path:      path,
				Language:  model.Python,
				Overwrite: true,
				Format:    model.Source,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
			Response: model.NotebookContent{
				Content: pythonNotebookDataB64,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
			Response: model.WorkspaceObjectStatus{
				ObjectID:   int64(objectId),
				ObjectType: model.Notebook,
				Path:       path,
				Language:   model.Python,
			},
		},
	}, resourceNotebook, map[string]interface{}{
		"path":      path,
		"content":   content,
		"language":  string(model.Python),
		"format":    string(model.Source),
		"overwrite": true,
		"mkdirs":    false,
	}, resourceNotebookCreate)
	assert.NoError(t, err, err)
	assert.Equal(t, path, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, path, d.Get("path"))
	assert.Equal(t, string(model.Python), d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestNotebooksRead(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)
	checkSum, err := convertBase64ToCheckSum(pythonNotebookDataB64)
	assert.NoError(t, err, err)
	exportFormat := model.Source
	testId := "/test/path.py"
	objectId := 12345
	assert.NoError(t, err, err)
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/export?format=SOURCE&path=%2Ftest%2Fpath.py",
			Response: model.NotebookContent{
				Content: pythonNotebookDataB64,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: "/api/2.0/workspace/get-status?path=%2Ftest%2Fpath.py",
			Response: model.WorkspaceObjectStatus{
				ObjectID:   int64(objectId),
				ObjectType: model.Notebook,
				Path:       testId,
				Language:   model.Python,
			},
		},
	}, resourceNotebook, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId(testId)
		err := d.Set("format", exportFormat)
		assert.NoError(t, err, err)
		return resourceNotebookRead(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
	assert.Equal(t, checkSum, d.Get("content"))
	assert.Equal(t, testId, d.Get("path"))
	assert.Equal(t, string(model.Python), d.Get("language"))
	assert.Equal(t, objectId, d.Get("object_id"))
}

func TestNotebooksDelete(t *testing.T) {
	testId := "/test/path.py"
	d, err := ResourceTester(t, []HTTPFixture{
		{
			Method:          http.MethodPost,
			Resource:        "/api/2.0/workspace/delete",
			Status:          http.StatusOK,
			ExpectedRequest: model.NotebookDeleteRequest{Path: testId, Recursive: true},
		},
	}, resourceNotebook, nil, func(d *schema.ResourceData, c interface{}) error {
		d.SetId(testId)
		return resourceNotebookDelete(d, c)
	})
	assert.NoError(t, err, err)
	assert.Equal(t, testId, d.Id())
}

func TestAccAwsNotebookResource_multiple_formats(t *testing.T) {
	testAccNotebookResourceMultipleFormats(t)
}

func TestAccAwsNotebookResource_scalability(t *testing.T) {
	testAccNotebookResourceMultipleFormats(t)
}

func TestAccAzureNotebookResource_multiple_formats(t *testing.T) {
	testAccNotebookResourceMultipleFormats(t)
}

func TestAccAzureNotebookResource_scalability(t *testing.T) {
	testAccNotebookResourceScalability(t)
}

func testAccNotebookResourceScalability(t *testing.T) {
	pythonNotebookDataB64, err := notebookToB64("testdata/tf-test-python.py")
	assert.NoError(t, err, err)

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testNotebookResourceDestroy,
		Steps: []resource.TestStep{
			{
				// use a dynamic configuration with the random name from above
				Config:  testAzureNotebookResourceMultipleNotebooks(pythonNotebookDataB64),
				Destroy: false,
			},
		},
	})
}

func testAccNotebookResourceMultipleFormats(t *testing.T) {
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

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
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

// Only will test source format based content
func testNotebookResourceAllNotebookTypes(pythonContent, sqlContent, scalaContent, rContent, dbcContent, jupyterContent, htmlContent, folderPrefix string) string {
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
