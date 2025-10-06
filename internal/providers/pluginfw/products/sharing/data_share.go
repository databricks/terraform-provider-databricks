package sharing

import (
	"context"
	"reflect"

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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceNameShare = "share"

func DataSourceShare() datasource.DataSource {
	return &ShareDataSource{}
}

var _ datasource.DataSourceWithConfigure = &ShareDataSource{}

type ShareDataSource struct {
	Client *common.DatabricksClient
}

type ShareData struct {
	sharing_tf.ShareInfo
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

func (s ShareData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	types := s.ShareInfo.GetComplexFieldTypes(ctx)
	types["provider_config"] = reflect.TypeOf(tfschema.ProviderConfigData{})
	return types
}

func (s ShareData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	s.ShareInfo.ApplySchemaCustomizations(attrs)
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (d *ShareDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(dataSourceNameShare)
}

func (d *ShareDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ShareData{}, nil)
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

	var config ShareData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !config.ProviderConfigData.IsNull() {
		var namespace tfschema.ProviderConfigData
		resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, diags := d.Client.GetWorkspaceClientForUnifiedProvider(ctx, workspaceID)
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
