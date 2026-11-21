package workspace_file

import (
	"context"
	"reflect"
	"sort"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceNameWorkspaceFilePaths = "workspace_file_paths"

// WorkspaceFilePathsData is the schema of the databricks_workspace_file_paths data source.
type WorkspaceFilePathsData struct {
	Path                  types.String `tfsdk:"path"`
	Recursive             types.Bool   `tfsdk:"recursive"`
	WorkspaceFilePathList types.List   `tfsdk:"workspace_file_path_list"`
	tfschema.Namespace
}

func (WorkspaceFilePathsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["recursive"] = attrs["recursive"].SetOptional()
	attrs["workspace_file_path_list"] = attrs["workspace_file_path_list"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	return attrs
}

func (WorkspaceFilePathsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_file_path_list": reflect.TypeOf(WorkspaceFilePath{}),
		"provider_config":          reflect.TypeOf(tfschema.ProviderConfigData{}),
	}
}

// WorkspaceFilePath describes a single workspace file returned by the data source.
type WorkspaceFilePath struct {
	Path       types.String `tfsdk:"path"`
	ObjectId   types.Int64  `tfsdk:"object_id"`
	ResourceId types.String `tfsdk:"resource_id"`
	Size       types.Int64  `tfsdk:"size"`
	CreatedAt  types.Int64  `tfsdk:"created_at"`
	ModifiedAt types.Int64  `tfsdk:"modified_at"`
}

func (WorkspaceFilePath) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	for _, name := range []string{"path", "object_id", "resource_id", "size", "created_at", "modified_at"} {
		attrs[name] = attrs[name].SetComputed()
	}
	return attrs
}

func (WorkspaceFilePath) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (m WorkspaceFilePath) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"path":        m.Path,
			"object_id":   m.ObjectId,
			"resource_id": m.ResourceId,
			"size":        m.Size,
			"created_at":  m.CreatedAt,
			"modified_at": m.ModifiedAt,
		})
}

func (m WorkspaceFilePath) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"path":        types.StringType,
			"object_id":   types.Int64Type,
			"resource_id": types.StringType,
			"size":        types.Int64Type,
			"created_at":  types.Int64Type,
			"modified_at": types.Int64Type,
		},
	}
}

// workspaceFilePathList keeps only objects of type FILE and returns them sorted by
// path, so that the result is stable across reads regardless of API ordering.
func workspaceFilePathList(ctx context.Context, objects []workspace.ObjectInfo) basetypes.ListValue {
	files := make([]workspace.ObjectInfo, 0, len(objects))
	for _, o := range objects {
		if o.ObjectType == workspace.ObjectTypeFile {
			files = append(files, o)
		}
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})
	values := make([]attr.Value, 0, len(files))
	for _, f := range files {
		values = append(values, WorkspaceFilePath{
			Path:       types.StringValue(f.Path),
			ObjectId:   types.Int64Value(f.ObjectId),
			ResourceId: types.StringValue(f.ResourceId),
			Size:       types.Int64Value(f.Size),
			CreatedAt:  types.Int64Value(f.CreatedAt),
			ModifiedAt: types.Int64Value(f.ModifiedAt),
		}.ToObjectValue(ctx))
	}
	return types.ListValueMust(WorkspaceFilePath{}.Type(ctx), values)
}

func DataSourceWorkspaceFilePaths() datasource.DataSource {
	return &WorkspaceFilePathsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &WorkspaceFilePathsDataSource{}

type WorkspaceFilePathsDataSource struct {
	Client *common.DatabricksClient
}

func (d *WorkspaceFilePathsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceNameWorkspaceFilePaths)
}

func (d *WorkspaceFilePathsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceFilePathsData{}, nil)
	resp.Schema = schema.Schema{
		Attributes: attrs,
		Blocks:     blocks,
	}
}

func (d *WorkspaceFilePathsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if d.Client == nil {
		d.Client = pluginfwcommon.ConfigureDataSource(req, resp)
	}
}

func (d *WorkspaceFilePathsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceNameWorkspaceFilePaths)

	var config WorkspaceFilePathsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDDataSource(ctx, config.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := d.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := config.Path.ValueString()
	var objects []workspace.ObjectInfo
	var err error
	if config.Recursive.ValueBool() {
		objects, err = w.Workspace.RecursiveList(ctx, path)
	} else {
		objects, err = w.Workspace.ListAll(ctx, workspace.ListWorkspaceRequest{Path: path})
	}
	if err != nil {
		resp.Diagnostics.AddError("Failed to list workspace files in "+path, err.Error())
		return
	}

	config.WorkspaceFilePathList = workspaceFilePathList(ctx, objects)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
