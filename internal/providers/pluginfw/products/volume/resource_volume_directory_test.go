package volume

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVolumeDirectoryResourceMetadata(t *testing.T) {
	r := ResourceVolumeDirectory()
	require.NotNil(t, r)

	var resp resource.MetadataResponse
	r.Metadata(context.Background(), resource.MetadataRequest{
		ProviderTypeName: "databricks",
	}, &resp)

	assert.Equal(t, "databricks_volume_directory", resp.TypeName)
}

func TestVolumeDirectoryResourceSchema(t *testing.T) {
	r := ResourceVolumeDirectory()
	require.NotNil(t, r)

	var resp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &resp)

	assert.NotNil(t, resp.Schema)
	assert.Contains(t, resp.Schema.Attributes, "directory_path")
	assert.Contains(t, resp.Schema.Attributes, "id")

	// Verify directory_path is required
	directoryPathAttr := resp.Schema.Attributes["directory_path"]
	assert.True(t, directoryPathAttr.IsRequired())
	assert.False(t, directoryPathAttr.IsOptional())
	assert.False(t, directoryPathAttr.IsComputed())

	// Verify id is computed
	idAttr := resp.Schema.Attributes["id"]
	assert.False(t, idAttr.IsRequired())
	assert.False(t, idAttr.IsOptional())
	assert.True(t, idAttr.IsComputed())
}

func TestVolumeDirectoryResourceConfigure(t *testing.T) {
	r := &VolumeDirectoryResource{}
	assert.Nil(t, r.Client)

	// Configure with nil provider data should not panic
	var resp resource.ConfigureResponse
	r.Configure(context.Background(), resource.ConfigureRequest{}, &resp)

	// Configure with mock client
	mockClient := &common.DatabricksClient{}
	r.Configure(context.Background(), resource.ConfigureRequest{
		ProviderData: mockClient,
	}, &resp)

	assert.NotNil(t, r.Client)
	assert.Equal(t, mockClient, r.Client)
}

func TestVolumeDirectoryResourceImplementsInterface(t *testing.T) {
	var _ resource.Resource = &VolumeDirectoryResource{}
	var _ resource.ResourceWithConfigure = &VolumeDirectoryResource{}
}
