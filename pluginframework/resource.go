package pluginframework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = (*DatabricksResource)(nil)

type DatabricksResource struct {
}

func (r *DatabricksResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
}

func (r *DatabricksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
}

func (r *DatabricksResource) Create(_ context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

func (r *DatabricksResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *DatabricksResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *DatabricksResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type ABC struct{}
