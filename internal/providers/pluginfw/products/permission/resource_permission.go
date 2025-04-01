package permission

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/permissions/entity"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const resourceName = "permission"

var (
        _ resource.Resource = &permissionResource{}
        _ resource.ResourceWithConfigure = &permissionResource{}
)

func NewPermissionResource() resource.Resource {
    return &permissionResource{}
}

type permissionResource struct {
        client *common.DatabricksClient
        context context.Context
}

func (r *permissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
        if req.ProviderData == nil {
                return
        }

        var provider 
        //TODO CHECK IF ACCOUNTID ATTRIBUTE IS SET IN PROVIDER CONFIG STRUCT,
        // SHOULD BE ABLE TO PULL IF IT'S SET OR NOT FROM THE GOLANG STRUCT ITSELF
        if != "" {
                accountClient, err := req.ProviderData.(*common.DatabricksClient).AccountClient()
                if err != nil {
                        resp.Diagnostics.AddError(
                                //TODO ADD ERROR MESSAGE IN LINE WITH PROVIDER
                                "Unable to configure the Databricks client",
                                fmt.Sprintf("Expected *common.DatabricksClient, got %T", req.ProviderData),
                        )
                r.client = accountClient
                }

        } else {
                workspaceClient, err := req.ProviderData.(*common.DatabricksClient).WorkspaceClient()
                if err != nil {
                        resp.Diagnostics.AddError(
                                //TODO ADD ERROR MESSAGE IN LINE WITH Provider
                                "Unable to configure the Databricks client",
                                fmt.Sprintf("Expected *common.DatabricksClient, got %T", req.ProviderData),
                        )

                r.client = workspaceClient
                }
        }
}

func (r *permissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
        resp.TypeName = pluginfwcommon.GetDatabricksProductionName(resourceName)
}


type permissionResourceModel struct {
        ObjectID types.String `tfsdk:"object_id"`
        ObjectType types.String `tfsdk:"object_type"`
        AccessControlList []permissionAccessControlListModel `tfsdk:"access_control_list"`
        LastUpdated types.String `tfsdk:"last_updated"`
}

//accessControlListModel is the same as iam.AccessControlRequest
// was originally just called this way in entity.go
type permissionAccessControlListModel struct {
        ServicePrincipalName types.String `tfsdk:"service_principal_name"`
        GroupName types.String `tfsdk:"group_name"`
        UserName types.String `tfsdk:"user_name"`
        PermissionLevel types.String `tfsdk:"permission_level"`
}


//TODO set some attributes as optional, required
// see /databricks/permissions/resource_permissions.go line 160
func (r *permissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
        attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, permissionResourceModel{}, nil)
        resp.Schema = schema.Schema{
                Attributes: attrs,
                Blocks: blocks,
        }
}


func (r *permissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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
                        ServicePrincipalName: acl.ServicePrincipalName,
                        GroupName: acl.GroupName,
                        UserName: acl.UserName,
                        PermissionLevel: acl.PermissionLevel,
                })
        }

        // create the permission
        acl, err := r.client.Update.(plan.ObjectID, iam.PermissionsRequest{
                RequestObjectType: plan.ObjectID,
                RequestObjectType: plan.ObjectType,
                AccessControlList: acls,
        })

}

func (r *permissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
        ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
}

func (r *permissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
        ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
}

func (r *permissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
        ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)
}


// safePutWithOwner is a workaround for the limitation where warehouse without owners cannot have IS_OWNER set
func (r *permissionResource) safePutWithOwner(ctx context.Context, objectID string, objectACL []iam.AccessControlRequest, mapping resourcePermissions, ownerOpt string) error {
	w, err := r.client.WorkspaceClient()
	if err != nil {
		return err
	}
	idParts := strings.Split(objectID, "/")
	id := idParts[len(idParts)-1]
	withOwner := mapping.addOwnerPermissionIfNeeded(objectACL, ownerOpt)
	_, err = w.Permissions.Set(ctx, iam.PermissionsRequest{
		RequestObjectId:   id,
		RequestObjectType: mapping.requestObjectType,
		AccessControlList: withOwner,
	})
	if err != nil {
		if strings.Contains(err.Error(), "with no existing owner must provide a new owner") {
			_, err = w.Permissions.Set(ctx, iam.PermissionsRequest{
				RequestObjectId:   id,
				RequestObjectType: mapping.requestObjectType,
				AccessControlList: objectACL,
			})
		}
		return err
	}
	return nil
}

func (r *permissionResource) getCurrentUser(ctx context.Context) (string, error) {
	w, err := r.client.WorkspaceClient()
	if err != nil {
		return "", err
	}
	me, err := w.CurrentUser.Me(ctx)
	if err != nil {
		return "", err
	}
	return me.UserName, nil
}
