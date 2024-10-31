package sharing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type SharesList struct {
	Shares []types.String `tfsdk:"shares" tf:"computed,optional,slice_set"`
}

func DataSourceShares() datasource.DataSource {
	return &SharesDataSource{}
}

var _ datasource.DataSourceWithConfigure = &SharesDataSource{}

type SharesDataSource struct {
	Client *common.DatabricksClient
}

func (d *SharesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName("shares")
}

func (d *SharesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(SharesList{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *SharesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *SharesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	shares, err := w.Shares.ListAll(ctx, sharing.ListSharesRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch shares", err.Error())
		return
	}

	shareNames := make([]types.String, len(shares))
	for i, share := range shares {
		shareNames[i] = types.StringValue(share.Name)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, SharesList{Shares: shareNames})...)
}
