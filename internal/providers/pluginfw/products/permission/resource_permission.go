package permission

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	resourceName = "permission"
)

var (
	apiPath string
	//_ resource.Resource = &PermissionResource{}
	//_ resource.ResourceWithConfigure = &PermissionResource{}
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
	AccessControlList []permissionAccessControlListModel   `tfsdk:"access_control"`
	LastUpdated       types.String                       `tfsdk:"last_updated"`
}

// accessControlListModel is the same as iam.AccessControlRequest
// was originally just called this way in entity.go
type permissionAccessControlListModel struct {
	ServicePrincipalId types.String `tfsdk:"service_principal_id"`
	GroupName            types.String `tfsdk:"group_name"`
	UserName             types.String `tfsdk:"user_name"`
	PermissionLevel      types.String `tfsdk:"permission_level"`
}

func (permissionResourceModel) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
  attrs["object_id"] = attrs["object_id"].SetOptional().SetComputed()
  attrs["object_type"] = attrs["object_type"].SetOptional().SetComputed()
  attrs["access_control"] = attrs["access_control"].SetRequired().SetComputed()

  return attrs
}


func (permissionResourceModel) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
  return map[string]reflect.Type{
    "access_control" : reflect.TypeOf(permissionAccessControlListModel{}),
  }
}

func (r *PermissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
        resp.Schema = schema.Schema{
                Attributes: map[string]schema.Attribute{
                        "object_id": schema.StringAttribute{
                                Computed: true,
                                Optional: true,
                        },
                        "object_type": schema.StringAttribute{
                                Computed: true,
                                Optional: true,
                        },
                        "last_updated": schema.StringAttribute{
                                Computed: true,
                        },
                },
                Blocks: map[string]schema.Block{
                  "access_control": schema.SetNestedBlock{
                    NestedObject: schema.NestedBlockObject{
                        Attributes: map[string]schema.Attribute{
                              "service_principal_id": schema.StringAttribute{
                                      Computed: true,
                                      Optional: true,
                                      Description: "The service principal ID of the access control entry.",
                              },
                              "group_name": schema.StringAttribute{
                                      Computed: true,
                                      Optional: true,
                                      Description: "The group name of the access control entry.",
                              },
                              "user_name": schema.StringAttribute{
                                      Computed: true,
                                      Optional: true,
                                      Description: "The user name of the access control entry.",
                              },
                              "permission_level": schema.StringAttribute{
                                      Computed: true,
                                      Optional: true,
                                      Description: "The permission level of the access control entry.",
                              },
                      },
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
	var acls []iam.AccessControlRequest
	for _, acl := range plan.AccessControlList {
		acls = append(acls, iam.AccessControlRequest{
			//ServicePrincipalName: strings.Trim(acl.ServicePrincipalId.String(), "\""),
			GroupName:            strings.Trim(acl.GroupName.String(), "\""),
			//UserName:             strings.Trim(acl.UserName.String(), "\""),
			PermissionLevel:      iam.PermissionLevel(strings.Trim(acl.PermissionLevel.String(), "\"")),
		})
	}


	// create the permission
	permission, err := r.workspaceClient.Permissions.Update(ctx, iam.PermissionsRequest{
		RequestObjectId:   plan.ObjectID.ValueString(),
		RequestObjectType: plan.ObjectType.ValueString(),
		AccessControlList: acls,
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
  for permissionAclIndex, permissionAcl := range permission.AccessControlList {
          plan.AccessControlList[permissionAclIndex] = permissionAccessControlListModel{
                ServicePrincipalId: types.StringValue(permissionAcl.ServicePrincipalName),
                GroupName:            types.StringValue(permissionAcl.GroupName),
                UserName:             types.StringValue(permissionAcl.UserName),
                PermissionLevel:      types.StringValue(string(permissionAcl.AllPermissions[permissionAclIndex].PermissionLevel)),
            }
  }


  plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
  
  // Set state to fully populated data
  diags = resp.State.Set(ctx, plan)
  resp.Diagnostics.Append(diags...)
  if resp.Diagnostics.HasError() {
    return
  }
}

	//apiPath = fmt.Sprintf("/api/2.0/permissions/%s/%s", plan.ObjectType.ValueString(), plan.ObjectID.ValueString())

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
	state.AccessControlList = []permissionAccessControlListModel{}
	for aclIndex, acl := range permissions.AccessControlList {
		state.AccessControlList = append(state.AccessControlList, permissionAccessControlListModel{
			ServicePrincipalId: types.StringValue(acl.ServicePrincipalName),
			GroupName:            types.StringValue(acl.GroupName),
			UserName:             types.StringValue(acl.UserName),
			PermissionLevel:      types.StringValue(string(acl.AllPermissions[aclIndex].PermissionLevel)),
    })
  }

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
