package permission

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
  "github.com/hashicorp/terraform-plugin-framework/path"
)

const (
	resourceName = "permission"
)

var (
	_ resource.Resource = &PermissionResource{}
	_ resource.ResourceWithConfigure = &PermissionResource{}
  _ validator.Object = &accessControlListValidator{}
)

func NewPermissionResource() resource.Resource {
	return &PermissionResource{}
}

type PermissionResource struct {
	client          *common.DatabricksClient
	workspaceClient databricks.WorkspaceClient
	context         context.Context
}

func (r *PermissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	workspaceClient, err := req.ProviderData.(*common.DatabricksClient).WorkspaceClient()
	if err != nil {
		resp.Diagnostics.AddError(
			//TODO ADD ERROR MESSAGE IN LINE WITH Provider
			"Unable to configure the Databricks client",
			fmt.Sprintf("Expected *common.DatabricksClient, got %T", req.ProviderData),
		)

	} else {
		r.workspaceClient = *workspaceClient
	}
}

func (r *PermissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}

type permissionResourceModel struct {
	ObjectID          types.String                       `tfsdk:"object_id"`
	ObjectType        types.String                       `tfsdk:"object_type"`
	AccessControlList permissionAccessControlModel       `tfsdk:"access_control"`
	LastUpdated       types.String                       `tfsdk:"last_updated"`
}

// accessControlListModel is the same as iam.AccessControlRequest
// was originally just called this way in entity.go
type permissionAccessControlModel struct {
	ServicePrincipalId types.String `tfsdk:"service_principal_id"`
	GroupName            types.String `tfsdk:"group_name"`
	UserName             types.String `tfsdk:"user_name"`
	PermissionLevel      types.String `tfsdk:"permission_level"`
}

type accessControlListValidator struct {}

func (v accessControlListValidator) Description(ctx context.Context) string {
	return "Only one of user_name, group_name, or service_principal_id may be set"
}

func (v accessControlListValidator) MarkdownDescription(ctx context.Context) string {
        return v.Description(ctx)
}


func (v accessControlListValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
  var userName, groupName, spID types.String

	// You must check for each field using its path
	if diags := req.Config.GetAttribute(ctx, path.Root("access_control").AtName("user_name"), &userName); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if diags := req.Config.GetAttribute(ctx, path.Root("access_control").AtName("group_name"), &groupName); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if diags := req.Config.GetAttribute(ctx, path.Root("access_control").AtName("service_principal_id"), &spID); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Count how many are set
	count := 0
	if !userName.IsNull() && userName.ValueString() != "" {
		count++
	}
	if !groupName.IsNull() && groupName.ValueString() != "" {
		count++
	}
	if !spID.IsNull() && spID.ValueString() != "" {
		count++
	}

	if count != 1 {
		resp.Diagnostics.AddError(
			"Invalid access control configuration",
			"Exactly one of `user_name`, `group_name`, or `service_principal_id` must be set.",
		)
	}
}



func (r *PermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
        resp.Schema = schema.Schema{
                Attributes: map[string]schema.Attribute{
                        "object_id": schema.StringAttribute{
                                Required: true,
                        },
                        "object_type": schema.StringAttribute{
                                Required: true,
                        },
                        "last_updated": schema.StringAttribute{
                                Computed: true,
                        },
                },
                Blocks: map[string]schema.Block{
                  "access_control": schema.SingleNestedBlock{
                      Attributes: map[string]schema.Attribute{
                            "service_principal_id": schema.StringAttribute{
                                    Optional: true,
                                    Description: "The service principal ID of the access control entry.",
                            },
                            "group_name": schema.StringAttribute{
                                    Optional: true,
                                    Description: "The group name of the access control entry.",
                            },
                            "user_name": schema.StringAttribute{
                                    Optional: true,
                                    Description: "The user name of the access control entry.",
                            },
                            "permission_level": schema.StringAttribute{
                                    Optional: true,
                                    Description: "The permission level of the access control entry.",
                            },
                      },
                      Validators: []validator.Object{
                        accessControlListValidator{},
                      },
                  },
                },
        }
}

// TODO set some attributes as optional, required
// see /databricks/permissions/resource_permissions.go line 160
//func (r *PermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
//	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, permissionResourceModel{}, nil)
//	resp.Schema = schema.Schema{
//		Attributes: attrs,
//		Blocks:     blocks,
//	}
//}

func (r *PermissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan permissionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

  
	// generate API request from plan
  var acl iam.AccessControlRequest
  acl.ServicePrincipalName = strings.Trim(plan.AccessControlList.ServicePrincipalId.String(), "\"")
  acl.GroupName = strings.Trim(plan.AccessControlList.GroupName.String(), "\"")
  acl.UserName = strings.Trim(plan.AccessControlList.UserName.String(), "\"")
  acl.PermissionLevel = iam.PermissionLevel(strings.Trim(plan.AccessControlList.PermissionLevel.String(), "\""))


	// create the permission
	permission, err := r.workspaceClient.Permissions.Update(ctx, iam.PermissionsRequest{
		RequestObjectId:   plan.ObjectID.ValueString(),
		RequestObjectType: plan.ObjectType.ValueString(),
		AccessControlList: []iam.AccessControlRequest{acl},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			//TODO ADD ERROR MESSAGE IN LINE WITH Provider
			"Unable to Create Permission",
			fmt.Sprintf("Error: %s", err.Error()),
		)
  }

  //Map response to to schema and populate Computed attribute values
  plan.ObjectID = types.StringValue(permission.ObjectId)
  plan.ObjectType = types.StringValue(permission.ObjectType)
  plan.AccessControlList = permissionAccessControlModel{
        ServicePrincipalId:   types.StringValue(permission.AccessControlList[0].ServicePrincipalName),
        GroupName:            types.StringValue(permission.AccessControlList[0].GroupName),
        UserName:             types.StringValue(permission.AccessControlList[0].UserName),
        PermissionLevel:      types.StringValue(string(permission.AccessControlList[0].AllPermissions[0].PermissionLevel)),
  }


  plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
  
  // Set state to fully populated data
  diags = resp.State.Set(ctx, plan)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }
}


func (r *PermissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	// Get Current State
	var state permissionResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get refreshed acls from API
  permissions, err := common.RetryOn504(ctx, func(ctx context.Context) (*iam.ObjectPermissions, error) {
    return r.workspaceClient.Permissions.Get(ctx, iam.GetPermissionRequest{
        RequestObjectId:   state.ObjectID.ValueString(),
        RequestObjectType: state.ObjectType.ValueString(),
    })
  })

	var apiErr *apierr.APIError
	// https://github.com/databricks/terraform-provider-databricks/issues/1227
	// platform propagates INVALID_STATE error for auto-purged clusters in
	// the permissions api. this adds "a logical fix" also here, not to introduce
	// cross-package dependency on "clusters".
	if errors.As(err, &apiErr) && strings.Contains(apiErr.Message, "Cannot access cluster") && apiErr.StatusCode == 400 {
		apiErr.StatusCode = 404
		apiErr.ErrorCode = "RESOURCE_DOES_NOT_EXIST"
		err = apiErr
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get permission",
			fmt.Sprintf("Unable to read permission, %s", err.Error()),
		)
		return
	}


	//Overwrite data with refreshed state
	//TODO: parse getResponse to permissionResourceModel
  state.AccessControlList.UserName = types.StringValue(permissions.AccessControlList[0].UserName)
  state.AccessControlList.GroupName = types.StringValue(permissions.AccessControlList[0].GroupName)
  state.AccessControlList.ServicePrincipalId = types.StringValue(permissions.AccessControlList[0].ServicePrincipalName)
  state.AccessControlList.PermissionLevel = types.StringValue(string(permissions.AccessControlList[0].AllPermissions[0].PermissionLevel))
  state.ObjectID = types.StringValue(permissions.ObjectId)
  state.ObjectType = types.StringValue(permissions.ObjectType)

	//Set refreshed State
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PermissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
}

func (r *PermissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
}

// safePutWithOwner is a workaround for the limitation where warehouse without owners cannot have IS_OWNER set
//func (r *PermissionResource) safePutWithOwner(ctx context.Context, objectID string, objectACL []iam.AccessControlRequest, mapping resourcePermissions, ownerOpt string) error {
//	w, err := r.client.WorkspaceClient()
//	if err != nil {
//		return err
//	}
//	idParts := strings.Split(objectID, "/")
//	id := idParts[len(idParts)-1]
//	withOwner := mapping.addOwnerPermissionIfNeeded(objectACL, ownerOpt)
//	_, err = w.Permissions.Set(ctx, iam.PermissionsRequest{
//		RequestObjectId:   id,
//		RequestObjectType: mapping.requestObjectType,
//		AccessControlList: withOwner,
//	})
//	if err != nil {
//		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
//			_, err = w.Permissions.Set(ctx, iam.PermissionsRequest{
//				RequestObjectId:   id,
//				RequestObjectType: mapping.requestObjectType,
//				AccessControlList: objectACL,
//			})
//		}
//		return err
//	}
//	return nil
//}

//func (r *PermissionResource) getCurrentUser(ctx context.Context) (string, error) {
//	w, err := r.client.WorkspaceClient()
//	if err != nil {
//		return "", err
//	}
//	me, err := w.CurrentUser.Me(ctx)
//	if err != nil {
//		return "", err
//	}
//	return me.UserName, nil
//}
