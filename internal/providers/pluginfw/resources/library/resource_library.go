package library

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/databricks-sdk-go"
)

const libraryDefaultInstallationTimeout = 15 * time.Minute

var _ resource.ResourceWithConfigure = &LibraryResource{}

func ResourceLibrary() resource.Resource {
	return &LibraryResource{}
}

func readLibrary(ctx context.Context, w *databricks.WorkspaceClient, waitParams compute.Wait, libraryRep string, libraryExtended *LibraryExtended) diag.Diagnostics {
	res, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, waitParams, libraryDefaultInstallationTimeout)
	if err != nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic("failed to wait for library installation", err.Error())}
	}

	for _, v := range res.LibraryStatuses {
		thisRep := v.Library.String()
		if thisRep == libraryRep {
			// This is not entirely necessary as we can directly write the fields in the config into the state, because there's no computed field.
			diags := converters.GoSdkToTfSdkStruct(ctx, v.Library, libraryExtended)

			if diags.HasError() {
				return diags
			}

			libraryExtended.ClusterId = types.StringValue(waitParams.ClusterID)

			return nil
		}
	}
	return diag.Diagnostics{diag.NewErrorDiagnostic("failed to find the installed library", fmt.Sprintf("failed to find %s on %s", libraryRep, waitParams.ClusterID))}
}

type LibraryExtended struct {
	compute_tf.Library
	ClusterId types.String `tfsdk:"cluster_id"`
}

type LibraryResource struct {
	Client *common.DatabricksClient
}

func (r *LibraryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "databricks_library_pluginframework"
}

func (r *LibraryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Library",
		Attributes: tfschema.ResourceStructToSchemaMap(LibraryExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "cluster_id")
			c.AddPlanModifier(objectplanmodifier.RequiresReplace(), "cran")
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "egg")
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "jar")
			c.AddPlanModifier(objectplanmodifier.RequiresReplace(), "maven")
			c.AddPlanModifier(objectplanmodifier.RequiresReplace(), "pypi")
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "requirements")
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "whl")
			return c
		}),
	}
}

func (r *LibraryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.Client == nil {
		r.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *LibraryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var libraryTfSDK LibraryExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &libraryTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var libGoSDK compute.Library
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, libraryTfSDK, &libGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	installLib := compute.InstallLibraries{
		Libraries: []compute.Library{libGoSDK},
	}
	req.Plan.GetAttribute(ctx, path.Root("cluster_id"), &installLib.ClusterId)
	err := w.Libraries.Install(ctx, installLib)
	if err != nil {
		resp.Diagnostics.AddError("failed to install library", err.Error())
		return
	}
	waitParams := compute.Wait{
		ClusterID: installLib.ClusterId,
		IsRunning: true,
	}
	libraryRep := libGoSDK.String()
	installedLib := LibraryExtended{}

	resp.Diagnostics.Append(readLibrary(ctx, w, waitParams, libraryRep, &installedLib)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, installedLib)...)
}

func (r *LibraryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var libraryTfSDK LibraryExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &libraryTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var libGoSDK compute.Library
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, libraryTfSDK, &libGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	clusterId := libraryTfSDK.ClusterId.ValueString()
	libraryRep := libGoSDK.String()
	installedLib := LibraryExtended{}
	waitParams := compute.Wait{
		ClusterID: clusterId,
		IsRefresh: true,
	}

	resp.Diagnostics.Append(readLibrary(ctx, w, waitParams, libraryRep, &installedLib)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, installedLib)...)
}

func (r *LibraryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("failed to update library", "updating library is not supported")
}

func (r *LibraryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var libraryTfSDK LibraryExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &libraryTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	clusterID := libraryTfSDK.ClusterId.ValueString()
	var libGoSDK compute.Library
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, libraryTfSDK, &libGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	libraryRep := libGoSDK.String()
	_, err := clusters.StartClusterAndGetInfo(ctx, w, clusterID)
	if err != nil {
		resp.Diagnostics.AddError("failed to start and get cluster", err.Error())
		return
	}
	cll, err := w.Libraries.ClusterStatusByClusterId(ctx, clusterID)
	if err != nil {
		resp.Diagnostics.AddError("failed to get libraries", err.Error())
		return
	}
	for _, v := range cll.LibraryStatuses {
		if v.Library.String() != libraryRep {
			continue
		}
		err := w.Libraries.Uninstall(ctx, compute.UninstallLibraries{
			ClusterId: clusterID,
			Libraries: []compute.Library{*v.Library},
		})
		if err != nil {
			resp.Diagnostics.AddError("failed to uninstall library", err.Error())
		}
		return
	}
	// Keeping the implementation to be consistent with the sdk-v2 implementation. Eventually we should update this to be not
	// an error, for cases such as the library being manually uninstalled.
	resp.Diagnostics.AddError("failed to uninstall library", fmt.Sprintf("failed to find %s on %s", libraryRep, clusterID))
}
