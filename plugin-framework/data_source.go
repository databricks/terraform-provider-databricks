package pluginframework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type DatabricksDataSource struct {
}

func (d *DatabricksDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
}

func (d *DatabricksDataSource) Read(_ context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}

func (d *DatabricksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
}
