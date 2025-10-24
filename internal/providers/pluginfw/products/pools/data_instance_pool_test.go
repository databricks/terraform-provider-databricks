package pools

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
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
