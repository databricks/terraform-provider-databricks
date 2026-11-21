package workspace_file

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa"
)

func TestWorkspaceFilePathList_FiltersFilesAndSortsByPath(t *testing.T) {
	ctx := context.Background()
	objects := []workspace.ObjectInfo{
		{Path: "/a/z.txt", ObjectType: workspace.ObjectTypeFile, ObjectId: 3, ResourceId: "3", Size: 30, CreatedAt: 1, ModifiedAt: 2},
		{Path: "/a/dir", ObjectType: workspace.ObjectTypeDirectory, ObjectId: 4},
		{Path: "/a/nb", ObjectType: workspace.ObjectTypeNotebook, Language: workspace.LanguagePython, ObjectId: 5},
		{Path: "/a/repo", ObjectType: workspace.ObjectTypeRepo, ObjectId: 6},
		{Path: "/a/b.yml", ObjectType: workspace.ObjectTypeFile, ObjectId: 1, ResourceId: "1", Size: 10, CreatedAt: 3, ModifiedAt: 4},
	}

	list := workspaceFilePathList(ctx, objects)

	var got []WorkspaceFilePath
	require.False(t, list.ElementsAs(ctx, &got, false).HasError())
	require.Len(t, got, 2)
	assert.Equal(t, "/a/b.yml", got[0].Path.ValueString())
	assert.Equal(t, int64(1), got[0].ObjectId.ValueInt64())
	assert.Equal(t, "1", got[0].ResourceId.ValueString())
	assert.Equal(t, int64(10), got[0].Size.ValueInt64())
	assert.Equal(t, int64(3), got[0].CreatedAt.ValueInt64())
	assert.Equal(t, int64(4), got[0].ModifiedAt.ValueInt64())
	assert.Equal(t, "/a/z.txt", got[1].Path.ValueString())
}

func TestWorkspaceFilePathList_Empty(t *testing.T) {
	ctx := context.Background()
	list := workspaceFilePathList(ctx, nil)
	assert.False(t, list.IsNull())
	assert.Empty(t, list.Elements())
}

func readWorkspaceFilePaths(t *testing.T, fixtures []qa.HTTPFixture, path string, recursive *bool) (WorkspaceFilePathsData, diag.Diagnostics) {
	ctx := context.Background()
	client, server, err := qa.HttpFixtureClient(t, fixtures)
	require.NoError(t, err)
	defer server.Close()

	ds := &WorkspaceFilePathsDataSource{Client: client}
	schemaResp := &datasource.SchemaResponse{}
	ds.Schema(ctx, datasource.SchemaRequest{}, schemaResp)
	require.False(t, schemaResp.Diagnostics.HasError())
	s := schemaResp.Schema
	objType := s.Type().TerraformType(ctx).(tftypes.Object)

	values := map[string]tftypes.Value{}
	for name, typ := range objType.AttributeTypes {
		values[name] = tftypes.NewValue(typ, nil)
	}
	values["path"] = tftypes.NewValue(tftypes.String, path)
	if recursive != nil {
		values["recursive"] = tftypes.NewValue(tftypes.Bool, *recursive)
	}

	resp := &datasource.ReadResponse{State: tfsdk.State{Schema: s, Raw: tftypes.NewValue(objType, nil)}}
	ds.Read(ctx, datasource.ReadRequest{Config: tfsdk.Config{Schema: s, Raw: tftypes.NewValue(objType, values)}}, resp)

	var state WorkspaceFilePathsData
	if !resp.Diagnostics.HasError() {
		resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	}
	return state, resp.Diagnostics
}

func statePaths(t *testing.T, state WorkspaceFilePathsData) []string {
	var files []WorkspaceFilePath
	require.False(t, state.WorkspaceFilePathList.ElementsAs(context.Background(), &files, false).HasError())
	paths := []string{}
	for _, f := range files {
		paths = append(paths, f.Path.ValueString())
	}
	return paths
}

func TestDataSourceWorkspaceFilePaths_NonRecursive(t *testing.T) {
	state, diags := readWorkspaceFilePaths(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb",
			Response: workspace.ListResponse{
				Objects: []workspace.ObjectInfo{
					{Path: "/a/b/c", ObjectType: workspace.ObjectTypeDirectory, ObjectId: 1},
					{Path: "/a/b/nb", ObjectType: workspace.ObjectTypeNotebook, ObjectId: 2},
					{Path: "/a/b/file.json", ObjectType: workspace.ObjectTypeFile, ObjectId: 3, Size: 42},
				},
			},
		},
	}, "/a/b", nil)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, []string{"/a/b/file.json"}, statePaths(t, state))
	assert.Equal(t, "/a/b", state.Path.ValueString())
}

func TestDataSourceWorkspaceFilePaths_Recursive(t *testing.T) {
	recursive := true
	state, diags := readWorkspaceFilePaths(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb",
			Response: workspace.ListResponse{
				Objects: []workspace.ObjectInfo{
					{Path: "/a/b/c", ObjectType: workspace.ObjectTypeDirectory, ObjectId: 1},
					{Path: "/a/b/top.sql", ObjectType: workspace.ObjectTypeFile, ObjectId: 2},
				},
			},
		},
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2Fa%2Fb%2Fc",
			Response: workspace.ListResponse{
				Objects: []workspace.ObjectInfo{
					{Path: "/a/b/c/nb", ObjectType: workspace.ObjectTypeNotebook, ObjectId: 3},
					{Path: "/a/b/c/nested.py", ObjectType: workspace.ObjectTypeFile, ObjectId: 4},
				},
			},
		},
	}, "/a/b", &recursive)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, []string{"/a/b/c/nested.py", "/a/b/top.sql"}, statePaths(t, state))
}

func TestDataSourceWorkspaceFilePaths_Error(t *testing.T) {
	_, diags := readWorkspaceFilePaths(t, []qa.HTTPFixture{
		{
			Method:   "GET",
			Resource: "/api/2.0/workspace/list?path=%2Fmissing",
			Status:   404,
			Response: apierr.APIError{
				ErrorCode:  "RESOURCE_DOES_NOT_EXIST",
				StatusCode: 404,
				Message:    "Path (/missing) doesn't exist.",
			},
		},
	}, "/missing", nil)
	require.True(t, diags.HasError())
	assert.Equal(t, "Failed to list workspace files in /missing", diags.Errors()[0].Summary())
	assert.Contains(t, diags.Errors()[0].Detail(), "doesn't exist")
}
