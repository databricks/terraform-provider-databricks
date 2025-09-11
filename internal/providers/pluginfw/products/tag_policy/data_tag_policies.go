// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tag_policy

import (
	"context"
	"reflect"

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
)

const dataSourcesName = "tag_policies"

var _ datasource.DataSourceWithConfigure = &TagPoliciesDataSource{}

func DataSourceTagPolicies() datasource.DataSource {
	return &TagPoliciesDataSource{}
}

// TagPoliciesData extends the main model with additional fields.
type TagPoliciesData struct {
	TagPolicies types.List   `tfsdk:"tag_policies"`
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (TagPoliciesData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_policies": reflect.TypeOf(tags_tf.TagPolicy{}),
	}
}

type TagPoliciesDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *TagPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *TagPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, TagPoliciesData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("tag_policies")
		c.SetOptional("workspace_id")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks TagPolicy",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *TagPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *TagPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TagPoliciesData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest tags.ListTagPoliciesRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.TagPolicies.ListTagPoliciesAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list tag_policies", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var tag_policy tags_tf.TagPolicy
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &tag_policy)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, tag_policy.ToObjectValue(ctx))
	}

	var newState TagPoliciesData
	newState.TagPolicies = types.ListValueMust(tags_tf.TagPolicy{}.Type(ctx), results)
	newState.WorkspaceID = config.WorkspaceID
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
