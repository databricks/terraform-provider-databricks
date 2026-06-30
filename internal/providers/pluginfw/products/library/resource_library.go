package library

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/compute_tf"
	"github.com/databricks/terraform-provider-databricks/libraries"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/databricks/databricks-sdk-go"
)

const resourceName = "library"
const libraryDefaultInstallationTimeout = 30 * time.Minute

var _ resource.ResourceWithConfigure = &LibraryResource{}
var _ resource.ResourceWithModifyPlan = &LibraryResource{}
var _ resource.ResourceWithUpgradeState = &LibraryResource{}

func ResourceLibrary() resource.Resource {
	return &LibraryResource{}
}

// readLibrary reads the status of the specified library on the specified cluster and returns the library metadata.
// If library cannot be found, either because the cluster doesn't exist, the library is not installed, or some other error, the first return value will be nil.
// The returned diagnostics will contain any errors or warnings that occurred during the operation, and the caller should check for errors before continuing.
func readLibrary(ctx context.Context, w *databricks.WorkspaceClient, waitParams compute.Wait, libraryRep string) (*LibraryExtended, diag.Diagnostics) {
	var d diag.Diagnostics
	res, err := libraries.WaitForLibrariesInstalledSdk(ctx, w, waitParams, libraryDefaultInstallationTimeout)
	if errors.Is(err, databricks.ErrInvalidParameterValue) {
		d.AddWarning("cluster not found", fmt.Sprintf("cluster %s not found", waitParams.ClusterID))
		return nil, d
	}
	if err != nil {
		d.AddError("failed to wait for library installation", err.Error())
		return nil, d
	}

	for _, v := range res.LibraryStatuses {
		thisRep := v.Library.String()
		if thisRep == libraryRep {
			libraryExtended := &LibraryExtended{}
			// This is not entirely necessary as we can directly write the fields in the config into the state, because there's no computed field.
			d.Append(converters.GoSdkToTfSdkStruct(ctx, v.Library, libraryExtended)...)

			if d.HasError() {
				return nil, d
			}

			libraryExtended.ClusterId = types.StringValue(waitParams.ClusterID)

			return libraryExtended, d
		}
	}
	if waitParams.IsRefresh {
		// During Read, the library may have been removed outside of Terraform (e.g. via UI).
		// Return nil without error so the caller can remove it from state and trigger re-creation.
		d.AddWarning("library not found", fmt.Sprintf("library %s not found on cluster %s", libraryRep, waitParams.ClusterID))
	} else {
		d.AddError("failed to find the installed library", fmt.Sprintf("failed to find %s on %s", libraryRep, waitParams.ClusterID))
	}
	return nil, d
}

// LibraryExtended is the schema v1 model: provider_config is a SingleNestedAttribute (types.Object).
type LibraryExtended struct {
	compute_tf.Library_SdkV2
	ClusterId types.String `tfsdk:"cluster_id"`
	ID        types.String `tfsdk:"id"` // Adding ID field to stay compatible with SDKv2
	tfschema.Namespace
}

func (l LibraryExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return tfschema.AddProviderConfigType(l.Library_SdkV2.GetComplexFieldTypes(ctx))
}

// LibraryExtendedV0 is the schema v0 model: provider_config is a ListNestedBlock
// (types.List, MaxItems=1). Used solely as the PriorSchema for the state upgrader.
type LibraryExtendedV0 struct {
	compute_tf.Library_SdkV2
	ClusterId types.String `tfsdk:"cluster_id"`
	ID        types.String `tfsdk:"id"`
	tfschema.Namespace_SdkV2
}

func (l LibraryExtendedV0) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	attrs := l.Library_SdkV2.GetComplexFieldTypes(ctx)
	attrs["provider_config"] = reflect.TypeOf(tfschema.ProviderConfig{})
	return attrs
}

type LibraryResource struct {
	Client *common.DatabricksClient
}

func (r *LibraryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *LibraryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, LibraryExtended{}, libraryCustomizer)
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "Terraform schema for Databricks Library",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

// libraryCustomizer is the schema customizer for the current (v1) schema. It is
// also reused for the v0 PriorSchema (with the extra listvalidator on
// provider_config) via libraryCustomizerV0.
func libraryCustomizer(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
	c.ConfigureAsSdkV2Compatible()
	for field, attribute := range c.ToNestedBlockObject().Attributes {
		// provider_config is a SingleNestedAttribute, but its plan modifier
		// is ProviderConfigPlanModifier (set via ConfigureProviderConfig below),
		// not RequiresReplace.
		if field == "provider_config" {
			continue
		}
		switch attribute.(type) {
		case tfschema.StringAttributeBuilder:
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), field)
		case tfschema.SingleNestedAttributeBuilder:
			c.AddPlanModifier(objectplanmodifier.RequiresReplace(), field)
		}
	}
	for field, block := range c.ToNestedBlockObject().Blocks {
		switch block.(type) {
		case tfschema.ListNestedBlockBuilder:
			c.AddPlanModifier(listplanmodifier.RequiresReplace(), field)
		}
	}
	c.SetRequired("cluster_id")
	c.SetOptional("id")
	c.SetComputed("id")
	c.SetDeprecated(clusters.EggDeprecationWarning, "egg")
	tfschema.ConfigureProviderConfig(&c)
	return c
}

// librarySchemaV0 reconstructs the v0 schema (provider_config as ListNestedBlock)
// so the state upgrader can decode v0 state.
func librarySchemaV0(ctx context.Context) schema.Schema {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, LibraryExtendedV0{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.ConfigureAsSdkV2Compatible()
		for field, attribute := range c.ToNestedBlockObject().Attributes {
			switch attribute.(type) {
			case tfschema.StringAttributeBuilder:
				c.AddPlanModifier(stringplanmodifier.RequiresReplace(), field)
			case tfschema.SingleNestedAttributeBuilder:
				c.AddPlanModifier(objectplanmodifier.RequiresReplace(), field)
			}
		}
		for field, block := range c.ToNestedBlockObject().Blocks {
			switch block.(type) {
			case tfschema.ListNestedBlockBuilder:
				c.AddPlanModifier(listplanmodifier.RequiresReplace(), field)
			}
		}
		c.SetRequired("cluster_id")
		c.SetOptional("id")
		c.SetComputed("id")
		c.SetDeprecated(clusters.EggDeprecationWarning, "egg")
		c.AddValidator(listvalidator.SizeAtMost(1), "provider_config")
		return c
	})
	return schema.Schema{
		Version:     0,
		Description: "Terraform schema for Databricks Library (v0)",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *LibraryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.Client == nil {
		r.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *LibraryResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	if r.Client == nil {
		return
	}
	tfschema.WorkspaceDriftDetection(ctx, r.Client, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}
	tfschema.ValidateWorkspaceID(ctx, r.Client, req, resp)
}

func (r *LibraryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	var libraryTfSDK LibraryExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &libraryTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDResource(ctx, libraryTfSDK.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
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
	_, err := clusters.StartClusterAndGetInfo(ctx, w, installLib.ClusterId)
	if err != nil {
		resp.Diagnostics.AddError("failed to start and get cluster", err.Error())
		return
	}
	err = w.Libraries.Install(ctx, installLib)
	if err != nil {
		resp.Diagnostics.AddError("failed to install library", err.Error())
		return
	}
	waitParams := compute.Wait{
		ClusterID: installLib.ClusterId,
		IsRunning: true,
	}
	libraryRep := libGoSDK.String()

	installedLib, diags := readLibrary(ctx, w, waitParams, libraryRep)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if installedLib == nil {
		resp.Diagnostics.AddError("failed to install library", fmt.Sprintf("the installed library %s was not found. Please report this to the maintainers of terraform-provider-databricks.", libraryRep))
		return
	}

	installedLib.ID = types.StringValue(libGoSDK.String())
	installedLib.ProviderConfig = libraryTfSDK.ProviderConfig
	resp.Diagnostics.Append(resp.State.Set(ctx, installedLib)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, installedLib.ProviderConfig, &resp.State)...)
}

func (r *LibraryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	var libraryTfSDK LibraryExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &libraryTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDResource(ctx, libraryTfSDK.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
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
	waitParams := compute.Wait{
		ClusterID: clusterId,
		IsRefresh: true,
	}

	installedLib, diags := readLibrary(ctx, w, waitParams, libraryRep)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if installedLib == nil {
		resp.Diagnostics.AddWarning("library not found", fmt.Sprintf("library %s not found on cluster %s, marking as deleted", libraryRep, clusterId))
		resp.State.RemoveResource(ctx)
		return
	}

	installedLib.ProviderConfig = libraryTfSDK.ProviderConfig
	resp.Diagnostics.Append(resp.State.Set(ctx, installedLib)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, installedLib.ProviderConfig, &resp.State)...)
}

func (r *LibraryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("failed to update library", "updating library is not supported")
}

// UpgradeState migrates state from v0 (provider_config as ListNestedBlock) to v1
// (provider_config as SingleNestedAttribute). Users who installed databricks_library
// on a release where provider_config was a list need this for their next plan to
// decode against the current Object-shaped schema.
func (r *LibraryResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v0Schema := librarySchemaV0(ctx)
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &v0Schema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior LibraryExtendedV0
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}
				pcObject, diags := tfschema.UpgradeProviderConfigListToObject(ctx, prior.ProviderConfig)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}
				upgraded := LibraryExtended{
					Library_SdkV2: prior.Library_SdkV2,
					ClusterId:     prior.ClusterId,
					ID:            prior.ID,
					Namespace:     tfschema.Namespace{ProviderConfig: pcObject},
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func (r *LibraryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
	var libraryTfSDK LibraryExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &libraryTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceID, diags := tfschema.GetWorkspaceIDResource(ctx, libraryTfSDK.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(diags...)
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
	if apierr.IsMissing(err) {
		resp.Diagnostics.AddWarning("cluster not found", fmt.Sprintf("cluster %s not found, skipping library uninstallation", clusterID))
		return
	}
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
