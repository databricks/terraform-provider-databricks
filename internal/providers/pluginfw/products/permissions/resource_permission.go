package permissions

import (
	"context"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	"github.com/databricks/terraform-provider-databricks/permissions"
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
// Note: Object identifiers are NOT defined in the struct - they are read/written dynamically
// using GetAttribute()/SetAttribute(). This eliminates the need for hardcoded fields
// and makes the code truly generic - new permission types require zero changes here.
type PermissionResourceModel struct {
	ID         types.String `tfsdk:"id"`
	ObjectType types.String `tfsdk:"object_type"`

	// Principal identifiers - exactly one required
	UserName             types.String `tfsdk:"user_name"`
	GroupName            types.String `tfsdk:"group_name"`
	ServicePrincipalName types.String `tfsdk:"service_principal_name"`

	// Permission level
	PermissionLevel types.String `tfsdk:"permission_level"`

	// Note: Object identifiers (cluster_id, job_id, etc.) are NOT defined here.
	// They are accessed dynamically using GetAttribute()/SetAttribute() based on
	// the definitions in permissions/permission_definitions.go
}

func (r *PermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

func (r *PermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	// Collect all object identifier field names for ConflictsWith validators
	allPermissions := permissions.AllResourcePermissions()
	objectFieldPaths := make([]path.Expression, 0, len(allPermissions))
	for _, mapping := range allPermissions {
		objectFieldPaths = append(objectFieldPaths, path.MatchRoot(mapping.GetField()))
	}

	attrs := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"object_type": schema.StringAttribute{
			Computed: true,
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
			Description: "Permission level for the principal on the object (e.g., CAN_MANAGE, CAN_USE, CAN_VIEW).",
			Validators: []validator.String{
				ValidatePermissionLevel(),
			},
		},
	}

	// Dynamically add object identifier attributes from permission definitions
	// Each object identifier is mutually exclusive with all others
	for i, mapping := range allPermissions {
		fieldName := mapping.GetField()

		// Build ConflictsWith list - all other object fields except this one
		conflictPaths := make([]path.Expression, 0, len(objectFieldPaths)-1)
		for j, p := range objectFieldPaths {
			if i != j {
				conflictPaths = append(conflictPaths, p)
			}
		}

		attrs[fieldName] = schema.StringAttribute{
			Optional:    true,
			Description: fmt.Sprintf("ID or path for %s object type. Conflicts with all other object identifier attributes.", mapping.GetObjectType()),
			Validators: []validator.String{
				stringvalidator.ConflictsWith(conflictPaths...),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		}
	}

	resp.Schema = schema.Schema{
		Description: "Manages permissions for a single principal on a Databricks object. " +
			"This resource is authoritative for the specified object-principal pair only. " +
			"Use `databricks_permissions` for managing all principals on an object at once.",
		Attributes: attrs,
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
	// Read principal and permission_level using GetAttribute
	var userName, groupName, servicePrincipalName, permissionLevel types.String

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("user_name"), &userName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("group_name"), &groupName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("service_principal_name"), &servicePrincipalName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("permission_level"), &permissionLevel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Get the object mapping and ID (reads object identifiers dynamically from plan)
	mapping, objectID, objectFieldName, objectFieldValue, err := r.getObjectMappingAndID(ctx, w, req.Plan)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get object mapping", err.Error())
		return
	}

	// Lock the object to prevent concurrent modifications
	lockObject(objectID)
	defer unlockObject(objectID)

	// Determine principal identifier
	var principal string
	if !userName.IsNull() && !userName.IsUnknown() && userName.ValueString() != "" {
		principal = userName.ValueString()
	} else if !groupName.IsNull() && !groupName.IsUnknown() && groupName.ValueString() != "" {
		principal = groupName.ValueString()
	} else if !servicePrincipalName.IsNull() && !servicePrincipalName.IsUnknown() && servicePrincipalName.ValueString() != "" {
		principal = servicePrincipalName.ValueString()
	} else {
		resp.Diagnostics.AddError("Invalid principal configuration", "exactly one of 'user_name', 'group_name', or 'service_principal_name' must be set")
		return
	}

	// Create the permission update request
	permLevel := iam.PermissionLevel(permissionLevel.ValueString())

	// Use Update API (PATCH) to add permissions for this principal only
	idParts := strings.Split(objectID, "/")
	permID := idParts[len(idParts)-1]

	_, err = w.Permissions.Update(ctx, iam.UpdateObjectPermissions{
		RequestObjectId:   permID,
		RequestObjectType: mapping.GetRequestObjectType(),
		AccessControlList: []iam.AccessControlRequest{
			{
				UserName:             userName.ValueString(),
				GroupName:            groupName.ValueString(),
				ServicePrincipalName: servicePrincipalName.ValueString(),
				PermissionLevel:      permLevel,
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to create permission", err.Error())
		return
	}

	// Set the ID, object_type, and all other fields in state
	resourceID := fmt.Sprintf("%s/%s", objectID, principal)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(resourceID))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_type"), types.StringValue(mapping.GetRequestObjectType()))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(objectFieldName), objectFieldValue)...) // Set the object identifier field
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("user_name"), userName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_name"), groupName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("service_principal_name"), servicePrincipalName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission_level"), permissionLevel)...)
}

func (r *PermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Read ID from state using GetAttribute
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Parse the ID to get object ID and principal
	objectID, principal, err := r.parseID(id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse resource ID", err.Error())
		return
	}

	// Get the mapping from the ID
	mapping, err := permissions.GetResourcePermissionsFromId(objectID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get resource permissions mapping", err.Error())
		return
	}

	// Read current permissions
	idParts := strings.Split(objectID, "/")
	permID := idParts[len(idParts)-1]

	perms, err := w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectId:   permID,
		RequestObjectType: mapping.GetRequestObjectType(),
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
	var currentPermissionLevel types.String
	for _, acl := range perms.AccessControlList {
		if r.matchesPrincipal(acl, principal) {
			// Update the state with the current permission level
			if len(acl.AllPermissions) > 0 {
				currentPermissionLevel = types.StringValue(string(acl.AllPermissions[0].PermissionLevel))
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

	// Read the object identifier field from current state to preserve it
	// (It should already be in state, but we need to make sure it stays there)
	var objectFieldValue types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(mapping.GetField()), &objectFieldValue)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update state using SetAttribute
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(mapping.GetField()), objectFieldValue)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission_level"), currentPermissionLevel)...)
}

func (r *PermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Read ID, principals, and permission_level from plan using GetAttribute
	var id, userName, groupName, servicePrincipalName, permissionLevel types.String

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("id"), &id)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("user_name"), &userName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("group_name"), &groupName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("service_principal_name"), &servicePrincipalName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("permission_level"), &permissionLevel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Parse the ID to get object ID and principal
	objectID, _, err := r.parseID(id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse resource ID", err.Error())
		return
	}

	// Lock the object to prevent concurrent modifications
	lockObject(objectID)
	defer unlockObject(objectID)

	// Get the mapping from the ID
	mapping, err := permissions.GetResourcePermissionsFromId(objectID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get resource permissions mapping", err.Error())
		return
	}

	// Update the permission using PATCH
	permLevel := iam.PermissionLevel(permissionLevel.ValueString())
	idParts := strings.Split(objectID, "/")
	permID := idParts[len(idParts)-1]

	_, err = w.Permissions.Update(ctx, iam.UpdateObjectPermissions{
		RequestObjectId:   permID,
		RequestObjectType: mapping.GetRequestObjectType(),
		AccessControlList: []iam.AccessControlRequest{
			{
				UserName:             userName.ValueString(),
				GroupName:            groupName.ValueString(),
				ServicePrincipalName: servicePrincipalName.ValueString(),
				PermissionLevel:      permLevel,
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to update permission", err.Error())
		return
	}

	// Read the object identifier field from current state to preserve it
	var objectFieldValue types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(mapping.GetField()), &objectFieldValue)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update state using SetAttribute - preserve the object identifier field
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(mapping.GetField()), objectFieldValue)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission_level"), permissionLevel)...)
}

func (r *PermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Read ID from state using GetAttribute
	var id types.String
	resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("id"), &id)...)
	if resp.Diagnostics.HasError() {
		return
	}

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Parse the ID to get object ID and principal
	objectID, principal, err := r.parseID(id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to parse resource ID", err.Error())
		return
	}

	// Lock the object to prevent concurrent modifications
	// This is CRITICAL for Delete to avoid race conditions when multiple
	// permission resources for the same object are deleted concurrently
	lockObject(objectID)
	defer unlockObject(objectID)

	// Get the mapping from the ID
	mapping, err := permissions.GetResourcePermissionsFromId(objectID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get resource permissions mapping", err.Error())
		return
	}

	// Read current permissions to see what to remove
	idParts := strings.Split(objectID, "/")
	permID := idParts[len(idParts)-1]

	currentPerms, err := w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectId:   permID,
		RequestObjectType: mapping.GetRequestObjectType(),
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
		RequestObjectId:   permID,
		RequestObjectType: mapping.GetRequestObjectType(),
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
	// Import ID format: /<resource_type>/<id>/<principal>
	// Example: /clusters/cluster-123/user@example.com

	w, err := r.Client.WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get workspace client", err.Error())
		return
	}

	// Parse the import ID
	objectID, principal, err := r.parseID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Import ID Format",
			fmt.Sprintf("Expected format: /<resource_type>/<id>/<principal>. Error: %s", err.Error()),
		)
		return
	}

	// Get the mapping from the ID to determine object type and field
	mapping, err := permissions.GetResourcePermissionsFromId(objectID)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get resource permissions mapping", err.Error())
		return
	}

	// Read current permissions from Databricks
	idParts := strings.Split(objectID, "/")
	permID := idParts[len(idParts)-1]

	perms, err := w.Permissions.Get(ctx, iam.GetPermissionRequest{
		RequestObjectId:   permID,
		RequestObjectType: mapping.GetRequestObjectType(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to read permissions", err.Error())
		return
	}

	// Find the specific principal's permission
	var found bool
	var permissionLevel string
	var userName, groupName, servicePrincipalName string

	for _, acl := range perms.AccessControlList {
		if r.matchesPrincipal(acl, principal) {
			if len(acl.AllPermissions) > 0 {
				permissionLevel = string(acl.AllPermissions[0].PermissionLevel)
				userName = acl.UserName
				groupName = acl.GroupName
				servicePrincipalName = acl.ServicePrincipalName
				found = true
				break
			}
		}
	}

	if !found {
		resp.Diagnostics.AddError(
			"Permission Not Found",
			fmt.Sprintf("No permission found for principal %q on object %q", principal, objectID),
		)
		return
	}

	// Set all attributes in state using SetAttribute
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(req.ID))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("object_type"), types.StringValue(mapping.GetRequestObjectType()))...)

	// Set the object identifier field (e.g., cluster_id, job_id, etc.)
	// Extract the configured value from the objectID
	configuredValue := permID
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(mapping.GetField()), types.StringValue(configuredValue))...)

	// Set principal fields - use null for empty strings to avoid ImportStateVerify failures
	if userName != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("user_name"), types.StringValue(userName))...)
	} else {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("user_name"), types.StringNull())...)
	}

	if groupName != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_name"), types.StringValue(groupName))...)
	} else {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_name"), types.StringNull())...)
	}

	if servicePrincipalName != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("service_principal_name"), types.StringValue(servicePrincipalName))...)
	} else {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("service_principal_name"), types.StringNull())...)
	}

	// Set permission level
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("permission_level"), types.StringValue(permissionLevel))...)
}

// Helper methods

// AttributeGetter is an interface for types that can get attributes (Plan, Config, State)
type AttributeGetter interface {
	GetAttribute(ctx context.Context, path path.Path, target interface{}) diag.Diagnostics
}

// PermissionMapping is an interface that abstracts the permissions mapping operations
type PermissionMapping interface {
	GetRequestObjectType() string
	GetObjectType() string
	GetID(context.Context, *databricks.WorkspaceClient, string) (string, error)
}

func (r *PermissionResource) getObjectMappingAndID(ctx context.Context, w *databricks.WorkspaceClient, getter AttributeGetter) (PermissionMapping, string, string, types.String, error) {
	// Dynamically iterate through all permission definitions to find which object ID is set
	allPermissions := permissions.AllResourcePermissions()

	for _, mapping := range allPermissions {
		var attrValue types.String
		diags := getter.GetAttribute(ctx, path.Root(mapping.GetField()), &attrValue)
		if diags.HasError() {
			continue // Attribute doesn't exist or has errors, try next
		}

		if !attrValue.IsNull() && !attrValue.IsUnknown() && attrValue.ValueString() != "" {
			configuredValue := attrValue.ValueString()

			// Get the object ID (may involve path resolution)
			objectID, err := mapping.GetID(ctx, w, configuredValue)
			if err != nil {
				return nil, "", "", types.String{}, err
			}

			// Return mapping, objectID, field name, and field value
			return mapping, objectID, mapping.GetField(), attrValue, nil
		}
	}

	// No object identifier was set
	return nil, "", "", types.String{}, fmt.Errorf("at least one object identifier must be set")
}

func (r *PermissionResource) matchesPrincipal(acl iam.AccessControlResponse, principal string) bool {
	return acl.UserName == principal ||
		acl.GroupName == principal ||
		acl.ServicePrincipalName == principal
}

func (r *PermissionResource) parseID(id string) (objectID string, principal string, error error) {
	// ID format: /<resource_type>/<id>/<principal>
	parts := strings.Split(id, "/")
	if len(parts) < 4 {
		return "", "", fmt.Errorf("invalid ID format: expected /<resource_type>/<id>/<principal>, got %s", id)
	}

	// Reconstruct object ID and get principal
	principal = parts[len(parts)-1]
	objectID = strings.Join(parts[:len(parts)-1], "/")

	return objectID, principal, nil
}
