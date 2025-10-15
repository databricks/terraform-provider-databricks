package sharing

import (
	"context"
	"reflect"
	"sort"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "share"

var _ resource.ResourceWithConfigure = &ShareResource{}

func ResourceShare() resource.Resource {
	return &ShareResource{}
}

type ShareInfoExtended struct {
	sharing_tf.ShareInfo_SdkV2
	tfschema.Namespace
	ID types.String `tfsdk:"id"` // Adding ID field to stay compatible with SDKv2
}

var _ pluginfwcommon.ComplexFieldTypeProvider = ShareInfoExtended{}

func (s ShareInfoExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	types := s.ShareInfo_SdkV2.GetComplexFieldTypes(ctx)
	types["provider_config"] = reflect.TypeOf(tfschema.ProviderConfig{})
	return types
}

func matchOrder[T any, K comparable](target, reference []T, keyFunc func(T) K) {
	// Create a map to store the index positions of each key in the reference slice.
	orderMap := make(map[K]int)
	for index, item := range reference {
		orderMap[keyFunc(item)] = index
	}

	// Sort the target slice based on the order defined in the orderMap.
	sort.Slice(target, func(i, j int) bool {
		return orderMap[keyFunc(target[i])] < orderMap[keyFunc(target[j])]
	})
}

func suppressCDFEnabledDiff(si *sharing.ShareInfo) {
	//suppress diff for CDF Enabled if HistoryDataSharingStatus is enabled , as API does not accept both fields to be set
	for i := range si.Objects {
		if si.Objects[i].HistoryDataSharingStatus == "ENABLED" {
			si.Objects[i].CdfEnabled = false
		}
	}
}

func resourceShareMap(si sharing.ShareInfo) map[string]sharing.SharedDataObject {
	m := make(map[string]sharing.SharedDataObject, len(si.Objects))
	for _, sdo := range si.Objects {
		m[sdo.Name] = sdo
	}
	return m
}

func equal(this sharing.SharedDataObject, other sharing.SharedDataObject) bool {
	if other.SharedAs == "" {
		other.SharedAs = this.SharedAs
	}
	//don't compare computed fields
	other.AddedAt = this.AddedAt
	other.AddedBy = this.AddedBy
	other.Status = this.Status
	other.HistoryDataSharingStatus = this.HistoryDataSharingStatus
	other.ForceSendFields = this.ForceSendFields // TODO: is this the right thing to do?
	return reflect.DeepEqual(this, other)
}

func diff(beforeSi sharing.ShareInfo, afterSi sharing.ShareInfo) []sharing.SharedDataObjectUpdate {
	beforeMap := resourceShareMap(beforeSi)
	afterMap := resourceShareMap(afterSi)
	changes := []sharing.SharedDataObjectUpdate{}
	// not in after so remove
	for _, beforeSdo := range beforeSi.Objects {
		if _, ok := afterMap[beforeSdo.Name]; ok {
			continue
		}
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateActionRemove,
			DataObject: &beforeSdo,
		})
	}

	// not in before so add
	// if in before but diff then update
	for _, afterSdo := range afterSi.Objects {
		if beforeSdo, ok := beforeMap[afterSdo.Name]; ok {
			if !equal(beforeSdo, afterSdo) {
				// do not send SharedAs
				afterSdo.SharedAs = ""
				changes = append(changes, sharing.SharedDataObjectUpdate{
					Action:     sharing.SharedDataObjectUpdateActionUpdate,
					DataObject: &afterSdo,
				})
			}
			continue
		}
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateActionAdd,
			DataObject: &afterSdo,
		})
	}
	return changes
}

func shareChanges(si sharing.ShareInfo, action string) sharing.UpdateShare {
	var changes []sharing.SharedDataObjectUpdate
	for _, obj := range si.Objects {
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateAction(action),
			DataObject: &obj,
		},
		)
	}
	return sharing.UpdateShare{
		Name:    si.Name,
		Owner:   si.Owner,
		Updates: changes,
	}
}

type ShareResource struct {
	Client *common.DatabricksClient
}

func (r *ShareResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksStagingName(resourceName)
}

func (r *ShareResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, ShareInfoExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.ConfigureAsSdkV2Compatible()
		c.SetRequired("name")

		c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "name") // ForceNew
		c.AddPlanModifier(int64planmodifier.UseStateForUnknown(), "created_at")
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "created_by")

		c.SetRequired("object", "data_object_type")
		c.SetRequired("object", "partition", "value", "op")
		c.SetRequired("object", "partition", "value", "name")

		c.SetComputed("id")
		c.SetOptional("provider_config")

		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Share",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (d *ShareResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if d.Client == nil && req.ProviderData != nil {
		d.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *ShareResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan ShareInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var planGoSDK sharing.ShareInfo
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &planGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createShare sharing.CreateShare
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &createShare)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !plan.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfig
		resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	shareInfo, err := w.Shares.Create(ctx, createShare)
	if err != nil {
		resp.Diagnostics.AddError("failed to create share", err.Error())
		return
	}

	shareChanges := shareChanges(planGoSDK, string(sharing.SharedDataObjectUpdateActionAdd))

	updatedShareInfo, err := w.Shares.Update(ctx, shareChanges)
	if err != nil {
		// delete orphaned share if update fails
		if d_err := w.Shares.DeleteByName(ctx, shareInfo.Name); d_err != nil {
			resp.Diagnostics.AddError("failed to delete orphaned share", d_err.Error())
			return
		}
		resp.Diagnostics.AddError("failed to update share", err.Error())
		return
	}

	matchOrder(updatedShareInfo.Objects, planGoSDK.Objects, func(obj sharing.SharedDataObject) string { return obj.Name })

	var newState ShareInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, updatedShareInfo, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState, d := r.syncEffectiveFields(ctx, plan, newState, effectiveFieldsActionCreateOrUpdate{})
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.ID = newState.Name

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ShareResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var existingState ShareInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateGoSDK sharing.ShareInfo
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &stateGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var getShareRequest sharing.GetShareRequest
	getShareRequest.IncludeSharedData = true
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("name"), &getShareRequest.Name)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !existingState.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfig
		resp.Diagnostics.Append(existingState.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	shareInfo, err := w.Shares.Get(ctx, getShareRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get share", err.Error())
		return
	}

	matchOrder(shareInfo.Objects, stateGoSDK.Objects, func(obj sharing.SharedDataObject) string { return obj.Name })
	suppressCDFEnabledDiff(shareInfo)

	var newState ShareInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, shareInfo, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState, d := r.syncEffectiveFields(ctx, existingState, newState, effectiveFieldsActionRead{})
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *ShareResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ShareInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan ShareInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var planGoSDK sharing.ShareInfo
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &planGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var getShareRequest sharing.GetShareRequest
	getShareRequest.Name = state.Name.ValueString()
	getShareRequest.IncludeSharedData = true

	var workspaceID string
	if !plan.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfig
		resp.Diagnostics.Append(plan.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	currentShareInfo, err := w.Shares.Get(ctx, getShareRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get current share info", err.Error())
		return
	}

	matchOrder(currentShareInfo.Objects, planGoSDK.Objects, func(obj sharing.SharedDataObject) string { return obj.Name })
	suppressCDFEnabledDiff(currentShareInfo)

	changes := diff(*currentShareInfo, planGoSDK)

	// if owner has changed, update the share owner
	if !plan.Owner.IsNull() {
		updatedShareInfo, err := w.Shares.Update(ctx, sharing.UpdateShare{
			Name:  state.Name.ValueString(),
			Owner: plan.Owner.ValueString(),
		})
		if err == nil {
			resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, updatedShareInfo, &state)...)
			if resp.Diagnostics.HasError() {
				return
			}
		} else {
			resp.Diagnostics.AddError("failed to update share owner", err.Error())
			return
		}
	}

	upToDateShareInfo := currentShareInfo
	if len(changes) > 0 || !plan.Comment.IsNull() {
		// if there are any other changes, update the share with the changes
		update := sharing.UpdateShare{
			Name:    plan.Name.ValueString(),
			Updates: changes,
		}
		if !plan.Comment.IsNull() {
			update.Comment = plan.Comment.ValueString()
		}
		upToDateShareInfo, err = w.Shares.Update(ctx, update)

		if err != nil {
			resp.Diagnostics.AddError("failed to update share", err.Error())

			rollbackShareInfo, rollbackErr := w.Shares.Update(ctx, sharing.UpdateShare{
				Name:  currentShareInfo.Name,
				Owner: currentShareInfo.Owner,
			})
			if rollbackErr != nil {
				resp.Diagnostics.AddError("failed to roll back", common.OwnerRollbackError(err, rollbackErr, currentShareInfo.Owner, plan.Owner.ValueString()).Error())
				return
			}

			resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, rollbackShareInfo, &state)...)
			if resp.Diagnostics.HasError() {
				return
			}
		}

	}

	matchOrder(upToDateShareInfo.Objects, planGoSDK.Objects, func(obj sharing.SharedDataObject) string { return obj.Name })
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, upToDateShareInfo, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, d := r.syncEffectiveFields(ctx, plan, state, effectiveFieldsActionCreateOrUpdate{})
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *ShareResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var state ShareInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteShareRequest sharing_tf.DeleteShareRequest
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("name"), &deleteShareRequest.Name)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var workspaceID string
	if !state.ProviderConfig.IsNull() {
		var namespace tfschema.ProviderConfig
		resp.Diagnostics.Append(state.ProviderConfig.As(ctx, &namespace, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: true,
		})...)
		if resp.Diagnostics.HasError() {
			return
		}
		workspaceID = namespace.WorkspaceID.ValueString()
	}
	w, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, workspaceID)
	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := w.Shares.DeleteByName(ctx, deleteShareRequest.Name.ValueString())
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete share", err.Error())
		return
	}
}

type effectiveFieldsAction interface {
	resourceLevel(ctx context.Context, state *ShareInfoExtended, plan sharing_tf.ShareInfo_SdkV2)
	objectLevel(ctx context.Context, state *sharing_tf.SharedDataObject_SdkV2, plan sharing_tf.SharedDataObject_SdkV2)
}

type effectiveFieldsActionCreateOrUpdate struct{}

func (effectiveFieldsActionCreateOrUpdate) resourceLevel(ctx context.Context, state *ShareInfoExtended, plan sharing_tf.ShareInfo_SdkV2) {
	state.SyncFieldsDuringCreateOrUpdate(ctx, plan)
}

func (effectiveFieldsActionCreateOrUpdate) objectLevel(ctx context.Context, state *sharing_tf.SharedDataObject_SdkV2, plan sharing_tf.SharedDataObject_SdkV2) {
	state.SyncFieldsDuringCreateOrUpdate(ctx, plan)
}

type effectiveFieldsActionRead struct{}

func (effectiveFieldsActionRead) resourceLevel(ctx context.Context, state *ShareInfoExtended, plan sharing_tf.ShareInfo_SdkV2) {
	state.SyncFieldsDuringRead(ctx, plan)
}

func (effectiveFieldsActionRead) objectLevel(ctx context.Context, state *sharing_tf.SharedDataObject_SdkV2, plan sharing_tf.SharedDataObject_SdkV2) {
	state.SyncFieldsDuringRead(ctx, plan)
}

// syncEffectiveFields syncs the effective fields between existingState and newState
// and returns the newState
//
// existingState: infrastructure values that are recorded in the existing terraform state.
// newState: latest infrastructure values that are returned by the CRUD API calls.
//
// HCL config is compared with this newState to determine what changes are to be made
// to the infrastructure and then the newState values are recorded in the terraform state.
// Hence we ignore the values in existingState which are not present in newState.
func (r *ShareResource) syncEffectiveFields(ctx context.Context, existingState, newState ShareInfoExtended, mode effectiveFieldsAction) (ShareInfoExtended, diag.Diagnostics) {
	var d diag.Diagnostics
	mode.resourceLevel(ctx, &newState, existingState.ShareInfo_SdkV2)
	existingStateObjects, _ := existingState.GetObjects(ctx)
	newStateObjects, _ := newState.GetObjects(ctx)
	finalObjects := []sharing_tf.SharedDataObject_SdkV2{}
	for i := range newStateObjects {
		// For each object in the new state, we check if it exists in the existing state
		// and if it does, we sync the effective fields.
		// If it does not exist, we keep the new state object as is.
		for j := range existingStateObjects {
			if newStateObjects[i].Name == existingStateObjects[j].Name {
				mode.objectLevel(ctx, &newStateObjects[i], existingStateObjects[j])
				break
			}
		}
		finalObjects = append(finalObjects, newStateObjects[i])
	}
	newState.SetObjects(ctx, finalObjects)
	newState.ProviderConfig = existingState.ProviderConfig
	return newState, d
}
