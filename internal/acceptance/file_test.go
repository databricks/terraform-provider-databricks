package acceptance

import (
	"context"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/common"
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

func TestUcAccFileUpdateOnLocalChange(t *testing.T) {
	createdTime := ""
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

func TestUcAccFileUpdateServerChange(t *testing.T) {
	createdTime := ""
	unityWorkspaceLevel(t, step{
		Template: `
		resource "databricks_schema" "this" {
			name 		 = "schema-{var.STICKY_RANDOM}"
			catalog_name = "main"
		}

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjCg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
		// We are modifying the resource during the check stage, which causes the TF validation to fail. Ignoring the error.
		ExpectError: regexp.MustCompile(` the plan was not empty`),
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
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

		resource "databricks_storage_credential" "external" {
			name = "cred-{var.STICKY_RANDOM}"
			aws_iam_role {
				role_arn = "{env.TEST_METASTORE_DATA_ACCESS_ARN}"
			}
			comment = "Managed by TF"
		}

		resource "databricks_external_location" "some" {
			name            = "external-{var.STICKY_RANDOM}"
			url             = "s3://{env.TEST_BUCKET}/somepath-{var.STICKY_RANDOM}"
			credential_name = databricks_storage_credential.external.id
			comment         = "Managed by TF"
		}

		resource "databricks_volume" "this" {
			name = "name-abc"
			comment = "comment-abc"
			catalog_name = "main"
			schema_name = databricks_schema.this.name 
			volume_type = "EXTERNAL"
			storage_location   = databricks_external_location.some.url
		}
		
		resource "databricks_file" "this" {
			content_base64 = "YWJjDg=="
			path = "/Volumes/${databricks_volume.this.catalog_name}/${databricks_volume.this.schema_name}/${databricks_volume.this.name}/abcde"
		}`,
	})
}
