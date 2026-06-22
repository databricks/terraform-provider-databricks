package pools

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stretchr/testify/assert"
)

func TestInstancePoolDataSource_Metadata(t *testing.T) {
	d := &InstancePoolDataSource{}
	ctx := context.Background()

	req := datasource.MetadataRequest{
		ProviderTypeName: "databricks",
	}
	resp := &datasource.MetadataResponse{}

	d.Metadata(ctx, req, resp)

	// Verify the staging name is set correctly
	assert.Contains(t, resp.TypeName, "instance_pool")
}

func TestInstancePoolDataSource_SchemaProviderConfig(t *testing.T) {
	d := &InstancePoolDataSource{}
	ctx := context.Background()

	resp := &datasource.SchemaResponse{}
	d.Schema(ctx, datasource.SchemaRequest{}, resp)

	providerConfig, ok := resp.Schema.Attributes["provider_config"].(schema.SingleNestedAttribute)
	assert.True(t, ok)
	assert.True(t, providerConfig.IsOptional())
	assert.True(t, providerConfig.Attributes["workspace_id"].IsOptional())
	assert.True(t, providerConfig.Attributes["workspace_id"].IsComputed())
}
