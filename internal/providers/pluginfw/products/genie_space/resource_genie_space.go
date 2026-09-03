package genie_space

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "genie_space"

var _ resource.ResourceWithConfigure = &GenieSpaceResource{}
var _ resource.ResourceWithModifyPlan = &GenieSpaceResource{}
var _ resource.ResourceWithImportState = &GenieSpaceResource{}

func ResourceGenieSpace() resource.Resource {
	return &GenieSpaceResource{}
}

type GenieSpaceResource struct {
	Client *autogen.DatabricksClient
}

// ProviderConfig contains the fields to configure the provider for this
// resource. Mirrors the autogen ProviderConfig in other unified-provider
// Plugin Framework resources (e.g. knowledge_assistant, app_space).
type ProviderConfig struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (r ProviderConfig) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetOptional()
	attrs["workspace_id"] = attrs["workspace_id"].SetComputed()
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddPlanModifier(
		stringplanmodifier.RequiresReplaceIf(providerConfigWorkspaceIDPlanModifier, "", ""))
	attrs["workspace_id"] = attrs["workspace_id"].(tfschema.StringAttributeBuilder).AddValidator(stringvalidator.LengthAtLeast(1))
	return attrs
}

// providerConfigWorkspaceIDPlanModifier triggers replacement if workspace_id
// changes from one non-empty value to another.
func providerConfigWorkspaceIDPlanModifier(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
	oldValue := req.StateValue.ValueString()
	newValue := req.PlanValue.ValueString()
	if oldValue != "" && newValue != "" && oldValue != newValue {
		resp.RequiresReplace = true
	}
}

func (r ProviderConfig) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (r ProviderConfig) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

func (r ProviderConfig) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// GenieSpace is the Terraform model for the databricks_genie_space resource.
// It maps to dashboards.GenieSpace plus the provider_config block used by
// the unified provider.
type GenieSpace struct {
	// Genie space ID (set by the API on create).
	SpaceId types.String `tfsdk:"space_id"`
	// Title of the Genie space.
	Title types.String `tfsdk:"title"`
	// SQL warehouse associated with the Genie space.
	WarehouseId types.String `tfsdk:"warehouse_id"`
	// Serialized contents of the Genie space as a JSON string.
	SerializedSpace types.String `tfsdk:"serialized_space"`
	// Optional description.
	Description types.String `tfsdk:"description"`
	// Parent workspace folder where the space is registered.
	// Create-only: changes force replacement.
	ParentPath types.String `tfsdk:"parent_path"`
	// ETag returned by the API on every read; informational.
	// Not sent on update (PATCH semantics use last-writer-wins).
	Etag           types.String `tfsdk:"etag"`
	ProviderConfig types.Object `tfsdk:"provider_config"`
}

func (m GenieSpace) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"provider_config": reflect.TypeOf(ProviderConfig{}),
	}
}

func (m GenieSpace) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"space_id":         m.SpaceId,
			"title":            m.Title,
			"warehouse_id":     m.WarehouseId,
			"serialized_space": m.SerializedSpace,
			"description":      m.Description,
			"parent_path":      m.ParentPath,
			"etag":             m.Etag,
			"provider_config":  m.ProviderConfig,
		},
	)
}

func (m GenieSpace) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"space_id":         types.StringType,
			"title":            types.StringType,
			"warehouse_id":     types.StringType,
			"serialized_space": types.StringType,
			"description":      types.StringType,
			"parent_path":      types.StringType,
			"etag":             types.StringType,
			"provider_config":  ProviderConfig{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringCreateOrUpdate keeps user-controlled fields aligned between
// the plan and the post-API state. The API may re-serialize serialized_space
// with different whitespace/key-order; if it is semantically equal to the
// plan, preserve the plan value to avoid spurious drift.
func (to *GenieSpace) SyncFieldsDuringCreateOrUpdate(ctx context.Context, from GenieSpace) {
	to.ProviderConfig = from.ProviderConfig
	if jsonSemanticallyEqual(from.SerializedSpace.ValueString(), to.SerializedSpace.ValueString()) {
		to.SerializedSpace = from.SerializedSpace
	}
}

// SyncFieldsDuringRead is invoked after a Read with the prior state as
// `from`. Same JSON-normalization rationale as during create/update.
func (to *GenieSpace) SyncFieldsDuringRead(ctx context.Context, from GenieSpace) {
	to.ProviderConfig = from.ProviderConfig
	if jsonSemanticallyEqual(from.SerializedSpace.ValueString(), to.SerializedSpace.ValueString()) {
		to.SerializedSpace = from.SerializedSpace
	}
}

func (m GenieSpace) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["space_id"] = attrs["space_id"].SetComputed()
	attrs["title"] = attrs["title"].SetRequired()
	attrs["warehouse_id"] = attrs["warehouse_id"].SetRequired()
	attrs["serialized_space"] = attrs["serialized_space"].SetRequired()
	attrs["description"] = attrs["description"].SetOptional()
	// parent_path is Required to mirror databricks_dashboard. Missing
	// folders are auto-created by createSpaceWithParentRetry.
	attrs["parent_path"] = attrs["parent_path"].SetRequired()
	attrs["etag"] = attrs["etag"].SetComputed()

	// space_id is the primary identifier — preserve across plans.
	attrs["space_id"] = attrs["space_id"].(tfschema.StringAttributeBuilder).
		AddPlanModifier(stringplanmodifier.UseStateForUnknown())
	// etag is API-managed; preserve across plans to avoid spurious drift.
	attrs["etag"] = attrs["etag"].(tfschema.StringAttributeBuilder).
		AddPlanModifier(stringplanmodifier.UseStateForUnknown())
	// parent_path is create-only.
	attrs["parent_path"] = attrs["parent_path"].(tfschema.StringAttributeBuilder).
		AddPlanModifier(stringplanmodifier.RequiresReplace())
	// serialized_space uses JSON semantic equality to suppress whitespace/
	// key-order churn introduced by the API on round-trip.
	attrs["serialized_space"] = attrs["serialized_space"].(tfschema.StringAttributeBuilder).
		AddPlanModifier(jsonSerializedSpacePlanModifier{})

	attrs["provider_config"] = attrs["provider_config"].SetOptional()
	attrs["provider_config"] = attrs["provider_config"].SetComputed()
	attrs["provider_config"] = attrs["provider_config"].(tfschema.SingleNestedAttributeBuilder).
		AddPlanModifier(tfschema.ProviderConfigPlanModifier{})

	return attrs
}

func (r *GenieSpaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *GenieSpaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, GenieSpace{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for the Databricks Genie Space resource",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *GenieSpaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *GenieSpaceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
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

// workspaceClientFor resolves the WorkspaceClient for the given provider_config
// object. Centralizes the boilerplate that otherwise repeats across every CRUD
// method (object parsing + client lookup) so each CRUD body stays under the
// CONTRIBUTING.md 40-line limit.
func (r *GenieSpaceResource) workspaceClientFor(ctx context.Context, providerConfig types.Object) (*databricks.WorkspaceClient, diag.Diagnostics) {
	var diags diag.Diagnostics
	var namespace ProviderConfig
	diags.Append(providerConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if diags.HasError() {
		return nil, diags
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())
	diags.Append(clientDiags...)
	return client, diags
}

// createSpaceWithParentRetry calls CreateSpace and, if the parent folder is
// missing, creates it via MkdirsByPath and retries once. Mirrors the auto-mkdir
// behavior of the SDKv2 dashboards resource.
func createSpaceWithParentRetry(ctx context.Context, client *databricks.WorkspaceClient, req dashboards.GenieCreateSpaceRequest) (*dashboards.GenieSpace, error) {
	created, err := client.Genie.CreateSpace(ctx, req)
	if err == nil || req.ParentPath == "" {
		return created, err
	}
	if _, statusErr := client.Workspace.GetStatusByPath(ctx, req.ParentPath); statusErr != nil && apierr.IsMissing(statusErr) {
		if mkErr := client.Workspace.MkdirsByPath(ctx, req.ParentPath); mkErr != nil {
			return nil, fmt.Errorf("failed to create parent folder %q: %w", req.ParentPath, mkErr)
		}
		return client.Genie.CreateSpace(ctx, req)
	}
	return nil, err
}

// getSpaceWithSerialized re-fetches a space with include_serialized_space=true
// so that etag/serialized_space reflect the latest server state. Used after
// Create and Update.
func getSpaceWithSerialized(ctx context.Context, client *databricks.WorkspaceClient, spaceId string) (*dashboards.GenieSpace, error) {
	return client.Genie.GetSpace(ctx, dashboards.GenieGetSpaceRequest{
		SpaceId:                spaceId,
		IncludeSerializedSpace: true,
	})
}

func (r *GenieSpaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan GenieSpace
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, diags := r.workspaceClientFor(ctx, plan.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	created, err := createSpaceWithParentRetry(ctx, client, dashboards.GenieCreateSpaceRequest{
		Title:           plan.Title.ValueString(),
		WarehouseId:     plan.WarehouseId.ValueString(),
		SerializedSpace: plan.SerializedSpace.ValueString(),
		Description:     plan.Description.ValueString(),
		ParentPath:      plan.ParentPath.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("failed to create genie_space", err.Error())
		return
	}
	fetched, err := getSpaceWithSerialized(ctx, client, created.SpaceId)
	if err != nil {
		resp.Diagnostics.AddError("failed to read genie_space after create", err.Error())
		return
	}
	newState := newStateFromGenieSpace(fetched, plan.ParentPath)
	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, plan.ProviderConfig, &resp.State)...)
}

func (r *GenieSpaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState GenieSpace
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, diags := r.workspaceClientFor(ctx, existingState.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	fetched, err := getSpaceWithSerialized(ctx, client, existingState.SpaceId.ValueString())
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get genie_space", err.Error())
		return
	}
	newState := newStateFromGenieSpace(fetched, existingState.ParentPath)
	newState.SyncFieldsDuringRead(ctx, existingState)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(tfschema.PopulateProviderConfigInState(ctx, r.Client, existingState.ProviderConfig, &resp.State)...)
}

func (r *GenieSpaceResource) update(ctx context.Context, plan GenieSpace, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.workspaceClientFor(ctx, plan.ProviderConfig)
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}
	updated, err := client.Genie.UpdateSpace(ctx, dashboards.GenieUpdateSpaceRequest{
		SpaceId:         plan.SpaceId.ValueString(),
		Title:           plan.Title.ValueString(),
		WarehouseId:     plan.WarehouseId.ValueString(),
		SerializedSpace: plan.SerializedSpace.ValueString(),
		Description:     plan.Description.ValueString(),
	})
	if err != nil {
		diags.AddError("failed to update genie_space", err.Error())
		return
	}
	fetched, err := getSpaceWithSerialized(ctx, client, updated.SpaceId)
	if err != nil {
		diags.AddError("failed to read genie_space after update", err.Error())
		return
	}
	newState := postUpdateState(plan, updated, fetched)
	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

// postUpdateState composes the state to persist after a successful UpdateSpace.
// The Genie GET API can lag behind UpdateSpace and return pre-update
// user-controlled fields (title, description, warehouse_id), which would
// trigger "Provider produced inconsistent result after apply" in the framework.
// We therefore trust the plan for user inputs and only take computed fields
// (space_id, etag) from the update response and the canonical serialized_space
// from the post-update GET so JSON normalization can still suppress diffs.
func postUpdateState(plan GenieSpace, updated, fetched *dashboards.GenieSpace) GenieSpace {
	out := plan
	out.SpaceId = types.StringValue(updated.SpaceId)
	out.Etag = types.StringValue(fetched.Etag)
	out.SerializedSpace = types.StringValue(fetched.SerializedSpace)
	return out
}

func (r *GenieSpaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan GenieSpace
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *GenieSpaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state GenieSpace
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, diags := r.workspaceClientFor(ctx, state.ProviderConfig)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	spaceId := state.SpaceId.ValueString()
	err := client.Genie.TrashSpace(ctx, dashboards.GenieTrashSpaceRequest{SpaceId: spaceId})
	if err == nil || apierr.IsMissing(err) {
		return
	}
	// An already-trashed space returns 403; confirm that's the case by
	// re-fetching and treating IsMissing as success. Mirrors the SDKv2
	// dashboard resource's behavior.
	if errors.Is(err, apierr.ErrPermissionDenied) {
		if _, getErr := client.Genie.GetSpace(ctx, dashboards.GenieGetSpaceRequest{SpaceId: spaceId}); getErr != nil && apierr.IsMissing(getErr) {
			return
		}
	}
	resp.Diagnostics.AddError("failed to delete genie_space", err.Error())
}

func (r *GenieSpaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")
	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: space_id. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("space_id"), parts[0])...)
}

// newStateFromGenieSpace converts a dashboards.GenieSpace response into the
// Terraform state model. parent_path is not returned by the API but is
// preserved from the prior plan/state because it is create-only.
func newStateFromGenieSpace(s *dashboards.GenieSpace, parentPath types.String) GenieSpace {
	return GenieSpace{
		SpaceId:         types.StringValue(s.SpaceId),
		Title:           types.StringValue(s.Title),
		WarehouseId:     types.StringValue(s.WarehouseId),
		SerializedSpace: types.StringValue(s.SerializedSpace),
		Description:     optionalString(s.Description),
		ParentPath:      parentPath,
		Etag:            types.StringValue(s.Etag),
	}
}

// optionalString returns a null string when the API field is empty,
// matching Terraform optional-attribute semantics.
func optionalString(value string) types.String {
	if value == "" {
		return types.StringNull()
	}
	return types.StringValue(value)
}

// jsonSerializedSpacePlanModifier suppresses plan diffs on serialized_space
// when the JSON values are semantically equal (whitespace and key-order
// differences only). Mirrors Henkel's _normalized_serialized_space approach.
type jsonSerializedSpacePlanModifier struct{}

func (jsonSerializedSpacePlanModifier) Description(context.Context) string {
	return "Suppresses plan diffs when serialized_space differs only by JSON whitespace or key ordering."
}

func (m jsonSerializedSpacePlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (jsonSerializedSpacePlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	if jsonSemanticallyEqual(req.StateValue.ValueString(), req.PlanValue.ValueString()) {
		resp.PlanValue = req.StateValue
	}
}

// jsonSemanticallyEqual reports whether two JSON strings encode the same
// value (ignoring whitespace and key order). Returns false if either input
// is not valid JSON so the original diff is preserved for the user.
func jsonSemanticallyEqual(a, b string) bool {
	if a == b {
		return true
	}
	normA, errA := normalizeJSON(a)
	normB, errB := normalizeJSON(b)
	if errA != nil || errB != nil {
		return false
	}
	return normA == normB
}

func normalizeJSON(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	var v any
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return "", err
	}
	// json.Marshal already sorts map keys lexicographically, so re-marshaling
	// any decoded JSON yields a canonical representation.
	out, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(out), nil
}
