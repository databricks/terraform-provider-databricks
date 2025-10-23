package volume

import (
	"context"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const resourceName = "volume_directory"

var _ resource.ResourceWithConfigure = &VolumeDirectoryResource{}

func ResourceVolumeDirectory() resource.Resource {
	return &VolumeDirectoryResource{}
}

type VolumeDirectoryResource struct {
	Client *common.DatabricksClient
}

type VolumeDirectoryModel struct {
	DirectoryPath types.String `tfsdk:"directory_path"`
	ID            types.String `tfsdk:"id"`
}

func (r *VolumeDirectoryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *VolumeDirectoryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages directories in Unity Catalog volumes using the Files API.",
		Attributes: map[string]schema.Attribute{
			"directory_path": schema.StringAttribute{
				Description: "The absolute path of the directory in a Unity Catalog volume (e.g., `/Volumes/catalog/schema/volume/path/to/dir`).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Description: "The ID of the directory resource, same as directory_path.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *VolumeDirectoryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.Client == nil {
		r.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *VolumeDirectoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var data VolumeDirectoryModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	directoryPath := data.DirectoryPath.ValueString()

	// Create the directory using the Files API
	err := w.Files.CreateDirectory(ctx, files.CreateDirectoryRequest{
		DirectoryPath: directoryPath,
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to create directory", err.Error())
		return
	}

	// Set the ID to the directory path
	data.ID = types.StringValue(directoryPath)

	// Save the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VolumeDirectoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var data VolumeDirectoryModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	directoryPath := data.ID.ValueString()

	// Check if the directory still exists using GetDirectoryMetadata
	err := w.Files.GetDirectoryMetadata(ctx, files.GetDirectoryMetadataRequest{
		DirectoryPath: directoryPath,
	})
	if err != nil {
		if apierr.IsMissing(err) {
			// Directory no longer exists, remove from state
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get directory metadata", err.Error())
		return
	}

	// Directory exists, keep the state as is
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VolumeDirectoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update is a no-op since directories are immutable once created
	// The directory_path has RequiresReplace plan modifier, so any change will trigger recreation
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var data VolumeDirectoryModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Just update the state with the current data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VolumeDirectoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var data VolumeDirectoryModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	directoryPath := data.ID.ValueString()

	// Delete the directory using the Files API
	err := w.Files.DeleteDirectory(ctx, files.DeleteDirectoryRequest{
		DirectoryPath: directoryPath,
	})
	if err != nil {
		// If the directory is already gone, that's okay
		if !apierr.IsMissing(err) {
			resp.Diagnostics.AddError("failed to delete directory", err.Error())
			return
		}
	}
}

// ImportState implements resource.ResourceWithImportState
func (r *VolumeDirectoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	// The import ID is the directory path
	directoryPath := req.ID

	// Set the state with the imported directory path
	data := VolumeDirectoryModel{
		DirectoryPath: types.StringValue(directoryPath),
		ID:            types.StringValue(directoryPath),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
