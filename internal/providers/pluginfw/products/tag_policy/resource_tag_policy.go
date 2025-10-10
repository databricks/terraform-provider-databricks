// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tag_policy

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/tags"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/tags_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const resourceName = "tag_policy"

var _ resource.ResourceWithConfigure = &TagPolicyResource{}

func ResourceTagPolicy() resource.Resource {
	return &TagPolicyResource{}
}

type TagPolicyResource struct {
	Client *autogen.DatabricksClient
}

// TagPolicy extends the main model with additional fields.
type TagPolicy struct {
	tags_tf.TagPolicy
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// TagPolicy struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m TagPolicy) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.TagPolicy.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagPolicy
// only implements ToObjectValue() and Type().
func (m TagPolicy) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	embeddedObj := m.TagPolicy.ToObjectValue(ctx)
	embeddedAttrs := embeddedObj.Attributes()

	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		embeddedAttrs,
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m TagPolicy) Type(ctx context.Context) attr.Type {
	embeddedType := m.TagPolicy.Type(ctx).(basetypes.ObjectType)
	attrTypes := embeddedType.AttributeTypes()

	return types.ObjectType{AttrTypes: attrTypes}
}

// SyncFieldsDuringCreateOrUpdate copies values from the plan into the receiver,
// including both embedded model fields and additional fields. This method is called
// during create and update.
func (m *TagPolicy) SyncFieldsDuringCreateOrUpdate(ctx context.Context, plan TagPolicy) {
	m.TagPolicy.SyncFieldsDuringCreateOrUpdate(ctx, plan.TagPolicy)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *TagPolicy) SyncFieldsDuringRead(ctx context.Context, existingState TagPolicy) {
	m.TagPolicy.SyncFieldsDuringRead(ctx, existingState.TagPolicy)
}

func (r *TagPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(resourceName)
}

func (r *TagPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs, blocks := tfschema.ResourceStructToSchemaMap(ctx, TagPolicy{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.AddPlanModifier(stringplanmodifier.UseStateForUnknown(), "tag_key")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks tag_policy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *TagPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.Client = autogen.ConfigureResource(req, resp)
}

func (r *TagPolicyResource) update(ctx context.Context, plan TagPolicy, diags *diag.Diagnostics, state *tfsdk.State) {
	client, clientDiags := r.Client.GetWorkspaceClient()
	diags.Append(clientDiags...)
	if diags.HasError() {
		return
	}

	var tag_policy tags.TagPolicy

	diags.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &tag_policy)...)
	if diags.HasError() {
		return
	}

	updateRequest := tags.UpdateTagPolicyRequest{
		TagPolicy:  tag_policy,
		TagKey:     plan.TagKey.ValueString(),
		UpdateMask: "description,values",
	}

	response, err := client.TagPolicies.UpdateTagPolicy(ctx, updateRequest)
	if err != nil {
		diags.AddError("failed to update tag_policy", err.Error())
		return
	}

	var newState TagPolicy
	diags.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if diags.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)
	diags.Append(state.Set(ctx, newState)...)
}

func (r *TagPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var plan TagPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var tag_policy tags.TagPolicy

	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, plan, &tag_policy)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createRequest := tags.CreateTagPolicyRequest{
		TagPolicy: tag_policy,
	}

	response, err := client.TagPolicies.CreateTagPolicy(ctx, createRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to create tag_policy", err.Error())
		return
	}

	var newState TagPolicy

	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringCreateOrUpdate(ctx, plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TagPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var existingState TagPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &existingState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest tags.GetTagPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, existingState, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.TagPolicies.GetTagPolicy(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get tag_policy", err.Error())
		return
	}

	var newState TagPolicy
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, existingState)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *TagPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	var plan TagPolicy
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.update(ctx, plan, &resp.Diagnostics, &resp.State)
}

func (r *TagPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	ctx = pluginfwcontext.SetUserAgentInResourceContext(ctx, resourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TagPolicy
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var deleteRequest tags.DeleteTagPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, state, &deleteRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := client.TagPolicies.DeleteTagPolicy(ctx, deleteRequest)
	if err != nil && !apierr.IsMissing(err) {
		resp.Diagnostics.AddError("failed to delete tag_policy", err.Error())
		return
	}
}

var _ resource.ResourceWithImportState = &TagPolicyResource{}

func (r *TagPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, ",")

	if len(parts) != 1 || parts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf(
				"Expected import identifier with format: tag_key. Got: %q",
				req.ID,
			),
		)
		return
	}

	tagKey := parts[0]
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("tag_key"), tagKey)...)
}
