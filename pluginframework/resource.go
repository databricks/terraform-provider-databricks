package pluginframework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = (*DatabricksResource)(nil)

type DatabricksResource struct {
}

func (r *DatabricksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
}

func (r *DatabricksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
}

func (r *DatabricksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

func (r *DatabricksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *DatabricksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *DatabricksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
