// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tag_policy

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/tags"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/tags_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "tag_policy"

var _ datasource.DataSourceWithConfigure = &TagPolicyDataSource{}

func DataSourceTagPolicy() datasource.DataSource {
	return &TagPolicyDataSource{}
}

type TagPolicyDataSource struct {
	Client *autogen.DatabricksClient
}

// TagPolicyData extends the main model with additional fields.
type TagPolicyData struct {
	// Timestamp when the tag policy was created
	CreateTime types.String `tfsdk:"create_time"`

	Description types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`

	TagKey types.String `tfsdk:"tag_key"`
	// Timestamp when the tag policy was last updated
	UpdateTime types.String `tfsdk:"update_time"`

	Values types.List `tfsdk:"values"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// TagPolicyData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m TagPolicyData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"values": reflect.TypeOf(tags_tf.Value{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, TagPolicyData
// only implements ToObjectValue() and Type().
func (m TagPolicyData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time": m.CreateTime,
			"description": m.Description,
			"id":          m.Id,
			"tag_key":     m.TagKey,
			"update_time": m.UpdateTime,
			"values":      m.Values,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m TagPolicyData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time": types.StringType,
			"description": types.StringType,
			"id":          types.StringType,
			"tag_key":     types.StringType,
			"update_time": types.StringType,
			"values": basetypes.ListType{
				ElemType: tags_tf.Value{}.Type(ctx),
			},
		},
	}
}

func (m TagPolicyData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["description"] = attrs["description"].SetComputed()
	attrs["id"] = attrs["id"].SetComputed()
	attrs["tag_key"] = attrs["tag_key"].SetRequired()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["values"] = attrs["values"].SetComputed()

	return attrs
}

func (r *TagPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *TagPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, TagPolicyData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks TagPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *TagPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *TagPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config TagPolicyData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest tags.GetTagPolicyRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
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

	var newState TagPolicyData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
