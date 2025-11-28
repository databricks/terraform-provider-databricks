package permissions

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const resourceName = "permission"

var _ resource.ResourceWithConfigure = &PermissionResource{}
var _ resource.ResourceWithImportState = &PermissionResource{}

func ResourcePermission() resource.Resource {
	return &PermissionResource{}
}

type PermissionResource struct {
	Client *common.DatabricksClient
}

// PermissionResourceModel represents the Terraform resource model
type PermissionResourceModel struct {
	ID         types.String `tfsdk:"id"`
	ObjectType types.String `tfsdk:"object_type"`
	ObjectID   types.String `tfsdk:"object_id"`

	// Principal identifiers - exactly one required
	UserName             types.String `tfsdk:"user_name"`
	GroupName            types.String `tfsdk:"group_name"`
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`

	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`
}

func (r *PermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *PermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages permissions for a single principal on a Databricks object. " +
			"This resource is authoritative for the specified object-principal pair only. " +
			"Use `databricks_permissions` for managing all principals on an object at once.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"object_type": schema.StringAttribute{
				Required:    true,
				Description: "The type of object to manage permissions for (e.g., 'clusters', 'jobs', 'notebooks', 'authorization'). See the Databricks Permissions API documentation for the full list of supported object types.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"object_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the object to manage permissions for. The format depends on the object type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// Principal identifiers - exactly one required, mutually exclusive
			"user_name": schema.StringAttribute{
				Optional:    true,
				Description: "User name of the principal. Conflicts with group_name and service_principal_name.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRoot("group_name"),
						path.MatchRoot("service_principal_name"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Group name of the principal. Conflicts with user_name and service_principal_name.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRoot("user_name"),
						path.MatchRoot("service_principal_name"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"service_principal_name": schema.StringAttribute{
				Optional:    true,
				Description: "Service principal name. Conflicts with user_name and group_name.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRoot("user_name"),
						path.MatchRoot("group_name"),
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// Permission level
			"permission_level": schema.StringAttribute{
				Required:    true,
				Description: "Permission level for the principal on the object (e.g., CAN_MANAGE, CAN_USE, CAN_VIEW). See the Databricks Permissions API documentation for valid permission levels for each object type.",
				Validators: []validator.String{
					ValidatePermissionLevel(),
				},
			},
		},
	}
}

func (r *PermissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*common.DatabricksClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *common.DatabricksClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.Client = client
}

func (r *PermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PermissionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.upsertPermission(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set computed fields
	principal := r.getPrincipalFromModel(&plan)
	plan.ID = types.StringValue(fmt.Sprintf("%s/%s/%s", plan.ObjectType.ValueString(), plan.ObjectID.ValueString(), principal))

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PermissionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Parse the ID to get object type, object ID and principal
	objectType, objectID, principal, err := r.parseID(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse resource ID", err.Error())
		return
	}

	// Read current permissions
	perms, err := w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectId:   objectID,
		RequestObjectType: objectType,
	})
	if err != nil {
		// If the object or permissions are not found, remove from state to trigger recreation
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Failed to read permissions", err.Error())
		return
	}

	// Filter for the specific principal
	found := false
	for _, acl := range perms.AccessControlList {
		if r.matchesPrincipal(acl, principal) {
			// Update the state with the current permission level
			if len(acl.AllPermissions) > 0 {
				state.PermissionLevel = types.StringValue(string(acl.AllPermissions[0].PermissionLevel))
				found = true
				break
			}
		}
	}

	if !found {
		// Permission for this specific principal no longer exists, remove from state to trigger recreation
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PermissionResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.upsertPermission(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PermissionResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	objectType := state.ObjectType.ValueString()
	objectID := state.ObjectID.ValueString()
	principal := r.getPrincipalFromModel(&state)

	// Lock the object to prevent concurrent modifications
	// This is CRITICAL for Delete to avoid race conditions when multiple
	// permission resources for the same object are deleted concurrently
	lockObject(objectType, objectID)
	defer unlockObject(objectType, objectID)

	// Read current permissions to see what to remove
	currentPerms, err := w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectId:   objectID,
		RequestObjectType: objectType,
	})
	if err != nil {
		// If the object or permissions are not found, the permission is already gone
		// This is the desired state, so we can return successfully
		if apierr.IsMissing(err) {
			return
		}
		resp.Diagnostics.AddError("Failed to read current permissions", err.Error())
		return
	}

	// Build a list of all permissions EXCEPT the one we're deleting
	var remainingACLs []iam.AccessControlRequest
	for _, acl := range currentPerms.AccessControlList {
		if !r.matchesPrincipal(acl, principal) {
			// Keep this ACL
			if len(acl.AllPermissions) > 0 {
				remainingACLs = append(remainingACLs, iam.AccessControlRequest{
					UserName:             acl.UserName,
					GroupName:            acl.GroupName,
					ServicePrincipalName: acl.ServicePrincipalName,
					PermissionLevel:      acl.AllPermissions[0].PermissionLevel,
				})
			}
		}
	}

	// Use Set to replace all permissions (effectively removing the specified principal)
	_, err = w.Permissions.Set(ctx, iam.SetObjectPermissions{
		RequestObjectId:   objectID,
		RequestObjectType: objectType,
		AccessControlList: remainingACLs,
	})
	if err != nil {
		// If the object or principal doesn't exist, the permission is already gone
		// This can happen if the underlying object or principal was deleted outside of Terraform
		if apierr.IsMissing(err) {
			return
		}
		resp.Diagnostics.AddError("Failed to delete permission", err.Error())
		return
	}
}

func (r *PermissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import ID format: <object_type>/<object_id>/<principal>
	// Example: clusters/cluster-123/user@example.com

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Parse the import ID
	objectType, objectID, principal, err := r.parseID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Import ID Format",
			fmt.Sprintf("Expected format: <object_type>/<object_id>/<principal>. Error: %s", err.Error()),
		)
		return
	}

	// Read current permissions from Databricks
	perms, err := w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectId:   objectID,
		RequestObjectType: objectType,
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to read permissions", err.Error())
		return
	}

	// Find the specific principal's permission
	var found bool
	var state PermissionResourceModel
	state.ID = types.StringValue(req.ID)
	state.ObjectType = types.StringValue(objectType)
	state.ObjectID = types.StringValue(objectID)

	for _, acl := range perms.AccessControlList {
		if r.matchesPrincipal(acl, principal) {
			if len(acl.AllPermissions) > 0 {
				state.PermissionLevel = types.StringValue(string(acl.AllPermissions[0].PermissionLevel))

				// Set principal fields - use null for empty strings
				if acl.UserName != "" {
					state.UserName = types.StringValue(acl.UserName)
				} else {
					state.UserName = types.StringNull()
				}

				if acl.GroupName != "" {
					state.GroupName = types.StringValue(acl.GroupName)
				} else {
					state.GroupName = types.StringNull()
				}

				if acl.ServicePrincipalName != "" {
					state.ServicePrincipalName = types.StringValue(acl.ServicePrincipalName)
				} else {
					state.ServicePrincipalName = types.StringNull()
				}

				found = true
				break
			}
		}
	}

	if !found {
		resp.Diagnostics.AddError(
			"Permission Not Found",
			fmt.Sprintf("No permission found for principal %q on object %s/%s", principal, objectType, objectID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Helper methods

// upsertPermission creates or updates a permission for a principal on an object
func (r *PermissionResource) upsertPermission(ctx context.Context, model *PermissionResourceModel, diags *diag.Diagnostics) {
	w, err := r.Client.WorkspaceClient()
	if err != nil {
		diags.AddError("Failed to get workspace client", err.Error())
		return
	}

	objectType := model.ObjectType.ValueString()
	objectID := model.ObjectID.ValueString()

	// Lock the object to prevent concurrent modifications
	lockObject(objectType, objectID)
	defer unlockObject(objectType, objectID)

	// Create the permission update request
	permLevel := iam.PermissionLevel(model.PermissionLevel.ValueString())

	_, err = w.Permissions.Update(ctx, iam.UpdateObjectPermissions{
		RequestObjectId:   objectID,
		RequestObjectType: objectType,
		AccessControlList: []iam.AccessControlRequest{
			{
				UserName:             model.UserName.ValueString(),
				GroupName:            model.GroupName.ValueString(),
				ServicePrincipalName: model.ServicePrincipalName.ValueString(),
				PermissionLevel:      permLevel,
			},
		},
	})
	if err != nil {
		diags.AddError("Failed to create/update permission", err.Error())
		return
	}
}

// getPrincipalFromModel extracts the principal identifier from the model
func (r *PermissionResource) getPrincipalFromModel(model *PermissionResourceModel) string {
	if !model.UserName.IsNull() && model.UserName.ValueString() != "" {
		return model.UserName.ValueString()
	}
	if !model.GroupName.IsNull() && model.GroupName.ValueString() != "" {
		return model.GroupName.ValueString()
	}
	if !model.ServicePrincipalName.IsNull() && model.ServicePrincipalName.ValueString() != "" {
		return model.ServicePrincipalName.ValueString()
	}
	return ""
}

func (r *PermissionResource) matchesPrincipal(acl iam.AccessControlResponse, principal string) bool {
	return acl.UserName == principal ||
		acl.GroupName == principal ||
		acl.ServicePrincipalName == principal
}

func (r *PermissionResource) parseID(id string) (objectType string, objectID string, principal string, err error) {
	// ID format: <object_type>/<object_id>/<principal>
	// Example: clusters/cluster-123/user@example.com
	parts := strings.Split(id, "/")
	if len(parts) < 3 {
		return "", "", "", fmt.Errorf("invalid ID format: expected <object_type>/<object_id>/<principal>, got %s", id)
	}

	// Handle cases where object_id might contain slashes (e.g., notebooks paths)
	// The principal is always the last part
	principal = parts[len(parts)-1]
	objectType = parts[0]
	objectID = strings.Join(parts[1:len(parts)-1], "/")

	return objectType, objectID, principal, nil
}
