package sharing

import (
	"context"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

const dataSourceNameShare = "share"

func DataSourceShare() datasource.DataSource {
	return &ShareDataSource{}
}

var _ datasource.DataSourceWithConfigure = &ShareDataSource{}

type ShareDataSource struct {
	Client *common.DatabricksClient
}

func (d *ShareDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(dataSourceNameShare)
}

func (d *ShareDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(sharing_tf.ShareInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *ShareDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *ShareDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceNameShare)
	w, diags := d.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config sharing_tf.ShareInfo
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	share, err := w.Shares.Get(ctx, sharing.GetShareRequest{
		Name:              config.Name.ValueString(),
		IncludeSharedData: true,
	})
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
		}

		resp.Diagnostics.AddError("Failed to fetch share", err.Error())
		return
	}

	var shareInfoTfSdk sharing_tf.ShareInfo
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, share, &shareInfoTfSdk)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, shareInfoTfSdk)...)
}
