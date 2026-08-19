// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package account_iam_service_principal_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/iamv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "account_iam_service_principals_v2"

var _ datasource.DataSourceWithConfigure = &ServicePrincipalsDataSource{}

func DataSourceServicePrincipals() datasource.DataSource {
	return &ServicePrincipalsDataSource{}
}

// ServicePrincipalsData extends the main model with additional fields.
type ServicePrincipalsData struct {
	AccountIamV2 types.List `tfsdk:"service_principals"`
	// Optional. Allows filtering service principals by application id or
	// external id.
	Filter types.String `tfsdk:"filter"`
	// The maximum number of service principals to return. The service may
	// return fewer than this value.
	PageSize types.Int64 `tfsdk:"page_size"`
}

func (ServicePrincipalsData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"service_principals": reflect.TypeOf(ServicePrincipalData{}),
	}
}

func (m ServicePrincipalsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["page_size"] = attrs["page_size"].SetOptional()
	attrs["filter"] = attrs["filter"].SetOptional()

	attrs["service_principals"] = attrs["service_principals"].SetComputed()
	return attrs
}

type ServicePrincipalsDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *ServicePrincipalsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *ServicePrincipalsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, ServicePrincipalsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks ServicePrincipal",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *ServicePrincipalsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *ServicePrincipalsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	var config ServicePrincipalsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest iamv2.ListServicePrincipalsRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.IamV2.ListServicePrincipalsAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list account_iam_service_principals_v2", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var service_principal ServicePrincipalData
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &service_principal)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, service_principal.ToObjectValue(ctx))
	}

	config.AccountIamV2 = types.ListValueMust(ServicePrincipalData{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, config)...)
}
