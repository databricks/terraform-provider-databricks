package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const deploymentResourceName = "app_deployment"

func ResourceAppDeployment() resource.Resource {
	return &resourceAppDeployment{}
}

type appDeploymentModel struct {
	AppName        types.String `tfsdk:"app_name"`
	SourceCodePath types.String `tfsdk:"source_code_path"`
	GitSource      types.Object `tfsdk:"git_source"`
	Mode           types.String `tfsdk:"mode"`
	Triggers       types.Map    `tfsdk:"triggers"`
	DeploymentId   types.String `tfsdk:"deployment_id"`
	Status         types.String `tfsdk:"status"`
	CreateTime     types.String `tfsdk:"create_time"`
}

type gitSourceModel struct {
	Branch         types.String `tfsdk:"branch"`
	Tag            types.String `tfsdk:"tag"`
	Commit         types.String `tfsdk:"commit"`
	SourceCodePath types.String `tfsdk:"source_code_path"`
	GitRepository  types.Object `tfsdk:"git_repository"`
	ResolvedCommit types.String `tfsdk:"resolved_commit"`
}

type gitRepositoryModel struct {
	Provider types.String `tfsdk:"provider"`
	Url      types.String `tfsdk:"url"`
}

type resourceAppDeployment struct {
	client *common.DatabricksClient
}

func (r resourceAppDeployment) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(deploymentResourceName)
}

func (r resourceAppDeployment) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	gitRepositorySchema := schema.SingleNestedAttribute{
		Optional:    true,
		Description: "Git repository configuration. If not specified, uses the app's git_repository configuration.",
		Attributes: map[string]schema.Attribute{
			"provider": schema.StringAttribute{
				Required:    true,
				Description: "Git provider. Supported values: gitHub, gitHubEnterprise, bitbucketCloud, bitbucketServer, azureDevOpsServices, gitLab, gitLabEnterpriseEdition, awsCodeCommit.",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "URL of the Git repository.",
			},
		},
	}

	gitSourceSchema := schema.SingleNestedAttribute{
		Optional:    true,
		Description: "Git source configuration for the deployment. Conflicts with source_code_path.",
		Attributes: map[string]schema.Attribute{
			"branch": schema.StringAttribute{
				Optional:    true,
				Description: "Git branch to checkout. Exactly one of branch, tag, or commit must be specified.",
			},
			"tag": schema.StringAttribute{
				Optional:    true,
				Description: "Git tag to checkout. Exactly one of branch, tag, or commit must be specified.",
			},
			"commit": schema.StringAttribute{
				Optional:    true,
				Description: "Git commit SHA to checkout. Exactly one of branch, tag, or commit must be specified.",
			},
			"source_code_path": schema.StringAttribute{
				Optional:    true,
				Description: "Relative path to the app source code within the Git repository. If not specified, the root of the repository is used.",
			},
			"git_repository": gitRepositorySchema,
			"resolved_commit": schema.StringAttribute{
				Computed:    true,
				Description: "The resolved commit SHA that was actually used for the deployment.",
			},
		},
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.RequiresReplace(),
		},
	}

	resp.Schema = schema.Schema{
		Description: "Deploys code to a Databricks app.",
		Attributes: map[string]schema.Attribute{
			"app_name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the app to deploy to.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"source_code_path": schema.StringAttribute{
				Optional:    true,
				Description: "The workspace filesystem path of the source code to deploy. Conflicts with git_source.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"git_source": gitSourceSchema,
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The deployment mode. Allowed values are SNAPSHOT and AUTO_SYNC. If not specified, defaults to SNAPSHOT.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"triggers": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "A map of arbitrary strings that, when changed, will force a new deployment.",
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
				},
			},
			"deployment_id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the deployment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "The status of the deployment.",
			},
			"create_time": schema.StringAttribute{
				Computed:    true,
				Description: "The creation time of the deployment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *resourceAppDeployment) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if r.client == nil && req.ProviderData != nil {
		r.client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *resourceAppDeployment) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config appDeploymentModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate exactly one of source_code_path or git_source is specified
	// Allow Unknown values (e.g., from local variables) - they'll be validated at plan time
	hasSourceCodePath := !config.SourceCodePath.IsNull()
	hasGitSource := !config.GitSource.IsNull()

	// Skip validation if values are unknown (will be resolved later)
	if config.SourceCodePath.IsUnknown() || config.GitSource.IsUnknown() {
		return
	}

	if !hasSourceCodePath && !hasGitSource {
		resp.Diagnostics.AddError(
			"Missing required field",
			"Exactly one of 'source_code_path' or 'git_source' must be specified",
		)
		return
	}

	if hasSourceCodePath && hasGitSource {
		resp.Diagnostics.AddError(
			"Conflicting fields",
			"Only one of 'source_code_path' or 'git_source' can be specified, not both",
		)
		return
	}

	// If git_source is specified, validate exactly one of branch, tag, or commit
	if hasGitSource && !config.GitSource.IsUnknown() {
		var gitSource gitSourceModel
		diags := config.GitSource.As(ctx, &gitSource, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		// Skip validation if any values are unknown
		if gitSource.Branch.IsUnknown() || gitSource.Tag.IsUnknown() || gitSource.Commit.IsUnknown() {
			return
		}

		hasBranch := !gitSource.Branch.IsNull()
		hasTag := !gitSource.Tag.IsNull()
		hasCommit := !gitSource.Commit.IsNull()

		refCount := 0
		if hasBranch {
			refCount++
		}
		if hasTag {
			refCount++
		}
		if hasCommit {
			refCount++
		}

		if refCount == 0 {
			resp.Diagnostics.AddError(
				"Missing required field in git_source",
				"Exactly one of 'branch', 'tag', or 'commit' must be specified in git_source",
			)
		} else if refCount > 1 {
			resp.Diagnostics.AddError(
				"Conflicting fields in git_source",
				"Only one of 'branch', 'tag', or 'commit' can be specified in git_source",
			)
		}
	}
}

func (r *resourceAppDeployment) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, deploymentResourceName)

	var plan appDeploymentModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get workspace client", err.Error())
		return
	}

	deployReq := apps.CreateAppDeploymentRequest{
		AppName:       plan.AppName.ValueString(),
		AppDeployment: apps.AppDeployment{},
	}

	// Handle source_code_path (workspace filesystem deployment)
	if !plan.SourceCodePath.IsNull() && !plan.SourceCodePath.IsUnknown() {
		deployReq.AppDeployment.SourceCodePath = plan.SourceCodePath.ValueString()
	}

	// Handle git_source (Git repository deployment)
	if !plan.GitSource.IsNull() && !plan.GitSource.IsUnknown() {
		var gitSource gitSourceModel
		diags := plan.GitSource.As(ctx, &gitSource, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		deployReq.AppDeployment.GitSource = &apps.GitSource{}

		if !gitSource.Branch.IsNull() && !gitSource.Branch.IsUnknown() {
			deployReq.AppDeployment.GitSource.Branch = gitSource.Branch.ValueString()
		}
		if !gitSource.Tag.IsNull() && !gitSource.Tag.IsUnknown() {
			deployReq.AppDeployment.GitSource.Tag = gitSource.Tag.ValueString()
		}
		if !gitSource.Commit.IsNull() && !gitSource.Commit.IsUnknown() {
			deployReq.AppDeployment.GitSource.Commit = gitSource.Commit.ValueString()
		}
		if !gitSource.SourceCodePath.IsNull() && !gitSource.SourceCodePath.IsUnknown() {
			deployReq.AppDeployment.GitSource.SourceCodePath = gitSource.SourceCodePath.ValueString()
		}

		// Handle git_repository if specified
		if !gitSource.GitRepository.IsNull() && !gitSource.GitRepository.IsUnknown() {
			var gitRepo gitRepositoryModel
			diags := gitSource.GitRepository.As(ctx, &gitRepo, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				return
			}

			deployReq.AppDeployment.GitSource.GitRepository = &apps.GitRepository{
				Provider: gitRepo.Provider.ValueString(),
				Url:      gitRepo.Url.ValueString(),
			}
		}
	}

	if !plan.Mode.IsNull() && !plan.Mode.IsUnknown() {
		deployReq.AppDeployment.Mode = apps.AppDeploymentMode(plan.Mode.ValueString())
	}

	deployment, err := w.Apps.DeployAndWait(ctx, deployReq)
	if err != nil {
		resp.Diagnostics.AddError("failed to deploy app", err.Error())
		return
	}

	plan.DeploymentId = types.StringValue(deployment.DeploymentId)
	plan.CreateTime = types.StringValue(deployment.CreateTime)
	if deployment.Status != nil {
		plan.Status = types.StringValue(string(deployment.Status.State))
	} else {
		plan.Status = types.StringNull()
	}
	// Populate mode from API response (it has a default even if user didn't specify)
	if deployment.Mode != "" {
		plan.Mode = types.StringValue(string(deployment.Mode))
	}

	// Populate git_source with resolved_commit from deployment response
	if deployment.GitSource != nil && !plan.GitSource.IsNull() {
		var gitSource gitSourceModel
		diags := plan.GitSource.As(ctx, &gitSource, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}

		// Update resolved_commit from API response (computed field)
		if deployment.GitSource.ResolvedCommit != "" {
			gitSource.ResolvedCommit = types.StringValue(deployment.GitSource.ResolvedCommit)
		}

		// Keep git_repository from plan - don't overwrite with API response
		// The API returns the git_repository, but we shouldn't populate it if user didn't specify it

		gitSourceObj, diags := types.ObjectValueFrom(ctx, gitSourceAttrTypes(), gitSource)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		plan.GitSource = gitSourceObj
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *resourceAppDeployment) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, deploymentResourceName)

	var state appDeploymentModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("failed to get workspace client", err.Error())
		return
	}

	deployment, err := w.Apps.GetDeploymentByAppNameAndDeploymentId(ctx, state.AppName.ValueString(), state.DeploymentId.ValueString())
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to read app deployment", err.Error())
		return
	}

	// Populate source_code_path if this is a workspace deployment
	if deployment.SourceCodePath != "" {
		state.SourceCodePath = types.StringValue(deployment.SourceCodePath)
		state.GitSource = types.ObjectNull(gitSourceAttrTypes())
	}

	// Populate git_source if this is a Git deployment
	if deployment.GitSource != nil {
		// Get existing git_source from state to preserve user-specified values
		var existingGitSource gitSourceModel
		if !state.GitSource.IsNull() {
			diags := state.GitSource.As(ctx, &existingGitSource, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				return
			}
		}

		gitSource := gitSourceModel{
			Branch:         types.StringNull(),
			Tag:            types.StringNull(),
			Commit:         types.StringNull(),
			SourceCodePath: types.StringNull(),
			GitRepository:  existingGitSource.GitRepository, // Preserve from state
			ResolvedCommit: types.StringNull(),
		}

		if deployment.GitSource.Branch != "" {
			gitSource.Branch = types.StringValue(deployment.GitSource.Branch)
		}
		if deployment.GitSource.Tag != "" {
			gitSource.Tag = types.StringValue(deployment.GitSource.Tag)
		}
		if deployment.GitSource.Commit != "" {
			gitSource.Commit = types.StringValue(deployment.GitSource.Commit)
		}
		if deployment.GitSource.SourceCodePath != "" {
			gitSource.SourceCodePath = types.StringValue(deployment.GitSource.SourceCodePath)
		}
		if deployment.GitSource.ResolvedCommit != "" {
			gitSource.ResolvedCommit = types.StringValue(deployment.GitSource.ResolvedCommit)
		}

		// Don't populate git_repository from API - keep what was in state
		// The API returns it, but we should preserve user's config (or lack thereof)

		gitSourceObj, diags := types.ObjectValueFrom(ctx, gitSourceAttrTypes(), gitSource)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		state.GitSource = gitSourceObj
		state.SourceCodePath = types.StringNull()
	}

	state.CreateTime = types.StringValue(deployment.CreateTime)
	if deployment.Mode != "" {
		state.Mode = types.StringValue(string(deployment.Mode))
	} else {
		state.Mode = types.StringNull()
	}
	if deployment.Status != nil {
		state.Status = types.StringValue(string(deployment.Status.State))
	} else {
		state.Status = types.StringNull()
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func gitSourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"branch":           types.StringType,
		"tag":              types.StringType,
		"commit":           types.StringType,
		"source_code_path": types.StringType,
		"git_repository":   types.ObjectType{AttrTypes: gitRepositoryAttrTypes()},
		"resolved_commit":  types.StringType,
	}
}

func gitRepositoryAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"provider": types.StringType,
		"url":      types.StringType,
	}
}

func (r *resourceAppDeployment) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	// All user-settable fields are ForceNew, so Update should never be called.
	resp.Diagnostics.AddError("unexpected update", "app_deployment does not support in-place updates")
}

func (r *resourceAppDeployment) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Deployments cannot be deleted via the API. Removing from state is sufficient.
}

func (r *resourceAppDeployment) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.SplitN(req.ID, "/", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("invalid import ID", fmt.Sprintf("expected format: app_name/deployment_id, got: %s", req.ID))
		return
	}

	var state appDeploymentModel
	state.AppName = types.StringValue(parts[0])
	state.DeploymentId = types.StringValue(parts[1])
	state.Triggers = types.MapNull(types.StringType)
	state.GitSource = types.ObjectNull(gitSourceAttrTypes())
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

var _ resource.ResourceWithConfigure = &resourceAppDeployment{}
var _ resource.ResourceWithImportState = &resourceAppDeployment{}
var _ resource.ResourceWithValidateConfig = &resourceAppDeployment{}
