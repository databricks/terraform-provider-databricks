// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package entity_tag_assignment

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "entity_tag_assignments"

var _ datasource.DataSourceWithConfigure = &EntityTagAssignmentsDataSource{}

func DataSourceEntityTagAssignments() datasource.DataSource {
	return &EntityTagAssignmentsDataSource{}
}

// EntityTagAssignmentsData extends the main model with additional fields.
type EntityTagAssignmentsData struct {
	EntityTagAssignments types.List `tfsdk:"tag_assignments"`
	// Optional. Maximum number of tag assignments to return in a single page
	MaxResults types.Int64  `tfsdk:"max_results"`
	EntityType types.String `tfsdk:"entity_type"`
	EntityName types.String `tfsdk:"entity_name"`
}

func (EntityTagAssignmentsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"tag_assignments": reflect.TypeOf(EntityTagAssignmentData{}),
	}
}

func (m EntityTagAssignmentsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["max_results"] = attrs["max_results"].SetOptional()

	attrs["tag_assignments"] = attrs["tag_assignments"].SetComputed()
	attrs["entity_type"] = attrs["entity_type"].SetRequired()
	attrs["entity_name"] = attrs["entity_name"].SetRequired()
	return attrs
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *EntityTagAssignmentsData) SyncFieldsDuringRead(ctx context.Context, from EntityTagAssignmentsData) {
}

type EntityTagAssignmentsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *EntityTagAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *EntityTagAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, EntityTagAssignmentsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks EntityTagAssignment",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *EntityTagAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *EntityTagAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config EntityTagAssignmentsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest catalog.ListEntityTagAssignmentsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetWorkspaceClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.EntityTagAssignments.ListAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list entity_tag_assignments", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var entity_tag_assignment EntityTagAssignmentData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &entity_tag_assignment)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, entity_tag_assignment.ToObjectValue(ctx))
	}

	var newState EntityTagAssignmentsData
	newState.EntityTagAssignments = types.ListValueMust(EntityTagAssignmentData{}.Type(ctx), results)
	newState.SyncFieldsDuringRead(ctx, config)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
