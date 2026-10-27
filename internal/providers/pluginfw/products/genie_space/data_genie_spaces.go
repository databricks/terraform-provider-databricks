package genie_space

import (
	"context"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/dashboards_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "genie_spaces"

func DataSourceGenieSpaces() datasource.DataSource {
	return &GenieSpacesDataSource{}
}

var _ datasource.DataSourceWithConfigure = &GenieSpacesDataSource{}

type GenieSpacesDataSource struct {
	Client *common.DatabricksClient
}

// GenieSpacesInfo is the Terraform model for the databricks_genie_spaces
// data source. title_contains is a case-insensitive substring filter on the
// space title (mirroring the dashboards data source's dashboard_name_contains).
type GenieSpacesInfo struct {
	TitleContains types.String `tfsdk:"title_contains"`
	Spaces        types.List   `tfsdk:"spaces"`
	tfschema.Namespace
}

func (GenieSpacesInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["title_contains"] = attrs["title_contains"].SetOptional()
	attrs["spaces"] = attrs["spaces"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (GenieSpacesInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"spaces":          reflect.TypeOf(dashboards_tf.GenieSpace{}),
		"provider_config": reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

func (d *GenieSpacesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *GenieSpacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, GenieSpacesInfo{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *GenieSpacesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *GenieSpacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var info GenieSpacesInfo
	resp.Diagnostics.Append(req.Config.Get(ctx, &info)...)
	if resp.Diagnostics.HasError() {
		return
	}
	workspaceID, diags := tfschema.GetWorkspaceIDDataSource(ctx, info.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	w, diags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	allSpaces, err := listAllGenieSpaces(ctx, w.Genie)
	if err != nil {
		resp.Diagnostics.AddError("failed to list Genie Spaces", err.Error())
		return
	}
	tfSpaces, ok := convertGenieSpacesToTF(ctx, filterSpacesByTitleContains(allSpaces, info.TitleContains.ValueString()), &resp.Diagnostics)
	if !ok {
		return
	}
	info.Spaces = types.ListValueMust(dashboards_tf.GenieSpace{}.Type(ctx), tfSpaces)
	resp.Diagnostics.Append(resp.State.Set(ctx, info)...)
}

// convertGenieSpacesToTF maps the SDK Genie spaces into the Plugin Framework
// representation. Returns the Terraform object values plus a flag indicating
// whether conversion succeeded. Diagnostics are appended to the passed-in sink.
func convertGenieSpacesToTF(ctx context.Context, spaces []dashboards.GenieSpace, diags *diag.Diagnostics) ([]attr.Value, bool) {
	out := make([]attr.Value, 0, len(spaces))
	for _, space := range spaces {
		var tfSpace dashboards_tf.GenieSpace
		diags.Append(converters.GoSdkToTfSdkStruct(ctx, space, &tfSpace)...)
		if diags.HasError() {
			return nil, false
		}
		out = append(out, tfSpace.ToObjectValue(ctx))
	}
	return out, true
}

// filterSpacesByTitleContains returns the spaces whose Title contains the
// query as a case-insensitive substring. An empty query matches everything.
func filterSpacesByTitleContains(spaces []dashboards.GenieSpace, query string) []dashboards.GenieSpace {
	if query == "" {
		return spaces
	}
	lower := strings.ToLower(query)
	out := make([]dashboards.GenieSpace, 0, len(spaces))
	for _, s := range spaces {
		if strings.Contains(strings.ToLower(s.Title), lower) {
			out = append(out, s)
		}
	}
	return out
}

// genieListClient is the subset of dashboards.GenieInterface used by
// listAllGenieSpaces. Defined here (not exported) so tests can substitute
// a fake without depending on the full SDK interface.
type genieListClient interface {
	ListSpaces(ctx context.Context, request dashboards.GenieListSpacesRequest) (*dashboards.GenieListSpacesResponse, error)
}

// listAllGenieSpaces walks every page returned by GenieAPI.ListSpaces and
// returns the union. The SDK does not provide a ListSpacesAll helper for
// Genie (unlike Lakeview.ListAll), so pagination is hand-rolled.
func listAllGenieSpaces(ctx context.Context, client genieListClient) ([]dashboards.GenieSpace, error) {
	var all []dashboards.GenieSpace
	var pageToken string
	for {
		page, err := client.ListSpaces(ctx, dashboards.GenieListSpacesRequest{
			PageToken: pageToken,
		})
		if err != nil {
			return nil, err
		}
		all = append(all, page.Spaces...)
		if page.NextPageToken == "" {
			return all, nil
		}
		pageToken = page.NextPageToken
	}
}
