package acceptance

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUcAccFileDontUpdateIfNoChange(t *testing.T) {
	createdTime := ""
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.True(t, m.LastModified != "")
			createdTime = m.LastModified
			return nil
		}),
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.Equal(t, m.LastModified, createdTime)
			return nil
		}),
	})
}

func TestUcAccFileUpdateOnLocalContentChange(t *testing.T) {
	createdTime := ""
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjZA=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.True(t, m.LastModified != "")
			createdTime = m.LastModified
			return nil
		}),
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.NotEqual(t, m.LastModified, createdTime)
			return nil
		}),
	})
}

func TestUcAccFileUpdateOnLocalFileChange(t *testing.T) {
	createdTime := ""
	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	fileName := tmpDir + "/upload_file"
	template := fmt.Sprintf(`
	resource "databricks_schema" "this" {
		name 		 = "schema-{var.STICKY_RANDOM}"
		catalog_name = "main"
	}

	resource "databricks_volume" "this" {
		name = "name-abc"
		comment = "comment-abc"
		catalog_name = "main"
		schema_name = databricks_schema.this.name 
		volume_type = "MANAGED"
	}
	
	resource "databricks_file" "this" {
		source = "%s"
		path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
	}`, fileName)
	unityWorkspaceLevel(t, step{
		PreConfig: func() {
			os.Mkdir(tmpDir, 0755)
			os.WriteFile(fileName, []byte("abc\n"), 0644)
		},
		Template: template,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.True(t, m.LastModified != "")
			createdTime = m.LastModified
			return nil
		}),
	}, step{
		PreConfig: func() {
			os.WriteFile(fileName, []byte("def\n"), 0644)
		},
		Template: template,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.NotEqual(t, m.LastModified, createdTime)
			return nil
		}),
	})
}

func TestUcAccFileNoUpdateIfFileDoesNotChange(t *testing.T) {
	createdTime := ""
	tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
	fileName := tmpDir + "/upload_file"
	template := fmt.Sprintf(`
	resource "databricks_schema" "this" {
		name 		 = "schema-{var.STICKY_RANDOM}"
		catalog_name = "main"
	}

	resource "databricks_volume" "this" {
		name = "name-abc"
		comment = "comment-abc"
		catalog_name = "main"
		schema_name = databricks_schema.this.name 
		volume_type = "MANAGED"
	}
	
	resource "databricks_file" "this" {
		source = "%s"
		path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
	}`, fileName)
	unityWorkspaceLevel(t, step{
		PreConfig: func() {
			os.Mkdir(tmpDir, 0755)
			os.WriteFile(fileName, []byte("abc\n"), 0644)
		},
		Template: template,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.True(t, m.LastModified != "")
			createdTime = m.LastModified
			return nil
		}),
	}, step{
		Template: template,
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.Equal(t, m.LastModified, createdTime)
			return nil
		}),
	})
}

func TestUcAccFileUpdateServerChange(t *testing.T) {
	createdTime := ""
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
		// We are modifying the resource during the check stage, which causes the TF validation to fail. Ignoring the error.
		ExpectError: regexp.MustCompile(` the refresh plan was not empty.`),
		Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
			w, err := client.WorkspaceClient()
			if err != nil {
				return err
			}
			m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
			if err != nil {
				return err
			}
			require.True(t, m.LastModified != "")
			createdTime = m.LastModified

			// Modify the file manually to test next step
			err = w.Files.Upload(ctx, files.UploadRequest{Contents: io.NopCloser(strings.NewReader("acdc")), FilePath: id})
			if err != nil {
				return err
			}
			return nil
		}),
	},
		step{
			Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
			Check: resourceCheck("databricks_file.this", func(ctx context.Context, client *common.DatabricksClient, id string) error {
				w, err := client.WorkspaceClient()
				if err != nil {
					return err
				}
				m, err := w.Files.GetMetadata(ctx, files.GetMetadataRequest{FilePath: id})
				if err != nil {
					return err
				}
				require.NotEqual(t, m.LastModified, createdTime)

				raw, err := w.Files.DownloadByFilePath(ctx, id)
				require.NoError(t, err)
				contents, err := io.ReadAll(raw.Contents)
				require.NoError(t, err)
				// Check that we updated the file
				assert.Equal(t, "abc\n", string(contents))
				return nil
			}),
		})
}

func TestUcAccFileFullLifeCycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python.py"
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			source = "{var.CWD}/../../storage/testdata/tf-test-python2.py"
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	})
}

func TestUcAccFileBase64FullLifeCycle(t *testing.T) {
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	}, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "MANAGED"
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjDg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	})
}
