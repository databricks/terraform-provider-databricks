package workspace_file

import (
	"context"
	"fmt"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	sdkworkspace "github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourceName = "workspace_file"

func DataSourceWorkspaceFilePaths() datasource.DataSource {
	return &WorkspaceFilePathsDataSource{}
}

var _ datasource.DataSourceWithConfigure = &WorkspaceFilePathsDataSource{}

type WorkspaceFilePathsDataSource struct {
	Client *common.DatabricksClient
}

type WorkspaceFileInfo struct {
	Id            types.String `tfsdk:"id"`
	CreatedAt     types.Int64  `tfsdk:"created_at"`
	ModifiedAt    types.Int64  `tfsdk:"modified_at"`
	ObjectId      types.Int64  `tfsdk:"object_id"`
	Path          types.String `tfsdk:"path"`
	ResourceId    types.String `tfsdk:"resource_id"`
	Url           types.String `tfsdk:"url"`
	WorkspacePath types.String `tfsdk:"workspace_path"`
}

func (WorkspaceFileInfo) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["id"] = attrs["id"].SetComputed()
	attrs["created_at"] = attrs["created_at"].SetComputed()
	attrs["modified_at"] = attrs["modified_at"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetComputed()
	attrs["path"] = attrs["path"].SetComputed()
	attrs["resource_id"] = attrs["resource_id"].SetComputed()
	attrs["url"] = attrs["url"].SetComputed()
	attrs["workspace_path"] = attrs["workspace_path"].SetComputed()
	return attrs
}

func (WorkspaceFileInfo) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

type WorkspaceFileDataSource struct {
	Path           types.String `tfsdk:"path"`
	Recursive      types.Bool   `tfsdk:"recursive"`
	WorkspaceFiles types.List   `tfsdk:"workspace_files"`
}

func (WorkspaceFileDataSource) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["path"] = attrs["path"].SetRequired()
	attrs["recursive"] = attrs["recursive"].SetOptional().SetComputed()
	attrs["workspace_files"] = attrs["workspace_files"].SetComputed()
	return attrs
}

func (WorkspaceFileDataSource) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"workspace_files": reflect.TypeOf(WorkspaceFileInfo{}),
	}
}

func (d *WorkspaceFilePathsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(dataSourceName)
}

func (d *WorkspaceFilePathsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceFileDataSource{}, nil)
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

var workspaceFileInfoAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"created_at":     types.Int64Type,
	"modified_at":    types.Int64Type,
	"object_id":      types.Int64Type,
	"path":           types.StringType,
	"resource_id":    types.StringType,
	"url":            types.StringType,
	"workspace_path": types.StringType,
}

func (d *WorkspaceFilePathsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var data WorkspaceFileDataSource
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := d.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get workspace client", err.Error())
		return
	}

	path := data.Path.ValueString()
	recursive := data.Recursive.ValueBool()

	var objects []sdkworkspace.ObjectInfo
	var listErr error
	if recursive {
		objects, listErr = w.Workspace.RecursiveList(ctx, path)
	} else {
		objects, listErr = w.Workspace.ListAll(ctx, sdkworkspace.ListWorkspaceRequest{
			Path: path,
		})
	}
	resp.Diagnostics.Append(checkListError(listErr, path)...)
	if resp.Diagnostics.HasError() {
		return
	}

	fileInfoType := types.ObjectType{AttrTypes: workspaceFileInfoAttrTypes}
	fileValues := make([]attr.Value, 0, len(objects))
	for _, obj := range objects {
		if obj.ObjectType == sdkworkspace.ObjectTypeDirectory {
			continue
		}
		fileValues = append(fileValues, types.ObjectValueMust(
			workspaceFileInfoAttrTypes,
			map[string]attr.Value{
				"id":             types.StringValue(obj.Path),
				"created_at":     types.Int64Value(obj.CreatedAt),
				"modified_at":    types.Int64Value(obj.ModifiedAt),
				"object_id":      types.Int64Value(obj.ObjectId),
				"path":           types.StringValue(obj.Path),
				"resource_id":    types.StringValue(obj.ResourceId),
				"url":            types.StringValue(d.Client.FormatURL("#workspace", obj.Path)),
				"workspace_path": types.StringValue("/Workspace" + obj.Path),
			},
		))
	}

	data.WorkspaceFiles = types.ListValueMust(fileInfoType, fileValues)
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func checkListError(err error, path string) diag.Diagnostics {
	if err == nil {
		return nil
	}
	if apierr.IsMissing(err) {
		return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("path '%s' does not exist", path), "")}
	}
	return diag.Diagnostics{diag.NewErrorDiagnostic(fmt.Sprintf("failed to list workspace files at path: %s", path), err.Error())}
}
