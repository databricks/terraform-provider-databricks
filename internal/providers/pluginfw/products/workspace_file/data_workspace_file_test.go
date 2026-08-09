package workspace_file

import (
	"context"
	"fmt"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/assert"
)

func TestDataSourceRegistersCorrectName(t *testing.T) {
	d := &WorkspaceFilePathsDataSource{}
	resp := &datasource.MetadataResponse{}
	d.Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "databricks"}, resp)
	assert.Equal(t, "databricks_workspace_file", resp.TypeName)
}

func TestSchemaContainsExpectedAttributes(t *testing.T) {
	d := &WorkspaceFilePathsDataSource{}
	resp := &datasource.SchemaResponse{}
	d.Schema(context.Background(), datasource.SchemaRequest{}, resp)
	assert.False(t, resp.Diagnostics.HasError())
	assert.Contains(t, resp.Schema.Attributes, "path")
	assert.Contains(t, resp.Schema.Attributes, "recursive")
	assert.Contains(t, resp.Schema.Attributes, "workspace_files")
}

func TestEmptyDirectory(t *testing.T) {
	diags := checkListError(nil, "/test/empty")
	assert.Nil(t, diags)
}

func TestNonExistentDirectory(t *testing.T) {
	err := &apierr.APIError{StatusCode: 404, Message: "Path doesn't exist"}
	diags := checkListError(err, "/test/nonexistent")
	expected := diag.Diagnostics{diag.NewErrorDiagnostic("path '/test/nonexistent' does not exist", "")}
	assert.True(t, diags.HasError())
	assert.Equal(t, expected, diags)
}

func TestListError(t *testing.T) {
	err := fmt.Errorf("some real error")
	diags := checkListError(err, "/test/path")
	expected := diag.Diagnostics{diag.NewErrorDiagnostic("failed to list workspace files at path: /test/path", "some real error")}
	assert.True(t, diags.HasError())
	assert.Equal(t, expected, diags)
}
