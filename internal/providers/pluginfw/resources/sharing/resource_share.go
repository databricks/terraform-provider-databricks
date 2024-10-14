package sharing

import (
	"context"
	"reflect"
	"sort"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/sharing_tf"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigure = &ShareResource{}

func ResourceShare() resource.Resource {
	return &ShareResource{}
}

type ShareInfoExtended struct {
	sharing_tf.ShareInfo
}

func sortSharesByName(si *sharing.ShareInfo) {
	sort.Slice(si.Objects, func(i, j int) bool {
		return si.Objects[i].Name < si.Objects[j].Name
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

func Equal(this sharing.SharedDataObject, other sharing.SharedDataObject) bool {
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

func Diff(beforeSi sharing.ShareInfo, afterSi sharing.ShareInfo) []sharing.SharedDataObjectUpdate {
	beforeMap := resourceShareMap(beforeSi)
	afterMap := resourceShareMap(afterSi)
	changes := []sharing.SharedDataObjectUpdate{}
	// not in after so remove
	for _, beforeSdo := range beforeSi.Objects {
		_, exists := afterMap[beforeSdo.Name]
		if exists {
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
		beforeSdo, exists := beforeMap[afterSdo.Name]
		if exists {
			if !Equal(beforeSdo, afterSdo) {
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

func shareChanges(ctx context.Context, si sharing.ShareInfo, action string) sharing.UpdateShare {
	var changes []sharing.SharedDataObjectUpdate
	for _, obj := range si.Objects {
		changes = append(changes, sharing.SharedDataObjectUpdate{
			Action:     sharing.SharedDataObjectUpdateAction(action),
			DataObject: &obj,
		},
		)
	}
	return sharing.UpdateShare{
		Updates: changes,
	}
}

type ShareResource struct {
	Client *common.DatabricksClient
}

func (r *ShareResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "databricks_share_pluginframework"
}

func (r *ShareResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks Share",
		Attributes: tfschema.ResourceStructToSchemaMap(ShareInfoExtended{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
			c.SetRequired("name")
			c.AddPlanModifier(stringplanmodifier.RequiresReplace(), "name") // ForceNew
			// c.SetCustomSuppressDiff("name", common.EqualFoldDiffSuppress)
			c.SetComputed("owner") // was originally SetSuppressDiff

			c.SetComputed("created_at")
			c.SetComputed("created_by")
			c.SetComputed("updated_at")
			c.SetComputed("updated_by")

			c.AddValidator(listvalidator.SizeAtLeast(1), "objects") // MinItems(1)
			c.SetRequired("objects", "data_object_type")
			c.SetComputed("objects", "shared_as")                   // was originally SetSuppressDiff
			c.SetComputed("objects", "cdf_enabled")                 // was originally SetSuppressDiff
			c.SetComputed("objects", "start_version")               // was originally SetSuppressDiff
			c.SetComputed("objects", "history_data_sharing_status") // was originally SetSuppressDiff
			c.SetComputed("objects", "status")
			c.SetComputed("objects", "added_at")
			c.SetComputed("objects", "added_by")
			c.SetRequired("objects", "partitions", "values", "op")
			c.SetRequired("objects", "partitions", "values", "name")
			return c
		}),
	}
}

func (d *ShareResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if d.Client == nil && req.ProviderData != nil {
		d.Client = pluginfwcommon.ConfigureResource(req, resp)
	}
}

func (r *ShareResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var shareInfoTfSDK ShareInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &shareInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createShareGoSDK sharing.CreateShare
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, shareInfoTfSDK, &createShareGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}
	share, err := w.Shares.Create(ctx, createShareGoSDK)
	if err != nil {
		resp.Diagnostics.AddError("failed to create share", err.Error())
		return
	}

	var shareInfoGoSDK sharing.ShareInfo
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, shareInfoTfSDK, &shareInfoGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	shareChanges := shareChanges(ctx, shareInfoGoSDK, string(sharing.SharedDataObjectUpdateActionAdd))
	shareChanges.Name = shareInfoTfSDK.Name.ValueString()
	shareChanges.Owner = shareInfoTfSDK.Owner.ValueString()

	updatedShareInfoGoSDK, err := w.Shares.Update(ctx, shareChanges)
	if err != nil {
		// delete orphaned share if update fails
		if d_err := w.Shares.DeleteByName(ctx, share.Name); d_err != nil {
			resp.Diagnostics.AddError("failed to delete orphaned share", d_err.Error())
			return
		}
		resp.Diagnostics.AddError("failed to update share", err.Error())
		return
	}

	var newShareInfoTfSDK ShareInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, updatedShareInfoGoSDK, &newShareInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newShareInfoTfSDK.Owner = shareInfoTfSDK.Owner

	resp.Diagnostics.Append(resp.State.Set(ctx, newShareInfoTfSDK)...)
}

func (r *ShareResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ShareInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var getShare sharing.GetShareRequest
	getShare.IncludeSharedData = true
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("name"), &getShare.Name)...)
	if resp.Diagnostics.HasError() {
		return
	}

	shareInfo, err := w.Shares.Get(ctx, getShare)
	sortSharesByName(shareInfo)
	suppressCDFEnabledDiff(shareInfo)

	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("failed to get share", err.Error())
		return
	}
	var shareInfoTfSDK ShareInfoExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, shareInfo, &shareInfoTfSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// shareInfoTfSDK.OriginalValues = state.OriginalValues

	resp.Diagnostics.Append(resp.State.Set(ctx, shareInfoTfSDK)...)
}

func (r *ShareResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state ShareInfoExtended
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// get a workspace client
	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan ShareInfoExtended
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plannedShareInfoGoSDK sharing.ShareInfo
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &plannedShareInfoGoSDK)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// get the share into a goSDK struct
	var getShare sharing.GetShareRequest
	getShare.Name = state.Name.ValueString()
	getShare.IncludeSharedData = true

	currentShareInfo, err := client.Shares.Get(ctx, getShare)
	if err != nil {
		return
	}

	sortSharesByName(currentShareInfo)
	suppressCDFEnabledDiff(currentShareInfo)
	changes := Diff(*currentShareInfo, plannedShareInfoGoSDK)

	// if owner has changed, update the share owner
	if !plan.Owner.IsNull() {
		updatedShareInfo, err := client.Shares.Update(ctx, sharing.UpdateShare{
			Name:  state.Name.ValueString(),
			Owner: plan.Owner.ValueString(),
		})
		if err == nil {
			resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, updatedShareInfo, &state)...)
			if resp.Diagnostics.HasError() {
				return
			}
		} else {
			resp.Diagnostics.AddError("failed to share owner", err.Error())
			return
		}
	}

	if len(changes) > 0 {
		// if there are any other changes, update the share with the changes
		updatedShareInfo, err := client.Shares.Update(ctx, sharing.UpdateShare{
			Name:    plan.Name.ValueString(),
			Updates: changes,
		})
		if err == nil {
			resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, updatedShareInfo, &state)...)
			if resp.Diagnostics.HasError() {
				return
			}
		} else {
			rollbackShareInfo, rollbackErr := client.Shares.Update(ctx, sharing.UpdateShare{
				Name:  currentShareInfo.Name,
				Owner: currentShareInfo.Owner,
			})
			if rollbackErr == nil {
				resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, rollbackShareInfo, &state)...)
				if resp.Diagnostics.HasError() {
					return
				}
			} else {
				resp.Diagnostics.AddError("failed to roll back", common.OwnerRollbackError(err, rollbackErr, currentShareInfo.Owner, plan.Owner.ValueString()).Error())
			}

			resp.Diagnostics.AddError("failed to update share", err.Error())
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *ShareResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	w, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest sharing_tf.DeleteShareRequest
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("name"), &deleteRequest.Name)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := w.Shares.DeleteByName(ctx, deleteRequest.Name.ValueString())
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete share", err.Error())
	}
}
