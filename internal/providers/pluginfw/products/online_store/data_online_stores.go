// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package online_store

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/ml_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const dataSourcesName = "online_stores"

var _ datasource.DataSourceWithConfigure = &OnlineStoresDataSource{}

func DataSourceOnlineStores() datasource.DataSource {
	return &OnlineStoresDataSource{}
}

// OnlineStoresData extends the main model with additional fields.
type OnlineStoresData struct {
	FeatureStore types.List `tfsdk:"online_stores"`
}

func (OnlineStoresData) GetComplexFieldTypes(context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"online_stores": reflect.TypeOf(ml_tf.OnlineStore{}),
	}
}

type OnlineStoresDataSource struct {
	Client *autogen.DatabricksClient
}

func (r *OnlineStoresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourcesName)
}

func (r *OnlineStoresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, OnlineStoresData{}, func(c tfschema.CustomizableSchema) tfschema.CustomizableSchema {
		c.SetComputed("online_stores")
		return c
	})
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks OnlineStore",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *OnlineStoresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *OnlineStoresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourcesName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config OnlineStoresData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var listRequest ml.ListOnlineStoresRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &listRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.FeatureStore.ListOnlineStoresAll(ctx, listRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to list online_stores", err.Error())
		return
	}

	var results = []attr.Value{}
	for _, item := range response {
		var online_store ml_tf.OnlineStore
		resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, item, &online_store)...)
		if resp.Diagnostics.HasError() {
			return
		}
		results = append(results, online_store.ToObjectValue(ctx))
	}

	var newState OnlineStoresData
	newState.FeatureStore = types.ListValueMust(ml_tf.OnlineStore{}.Type(ctx), results)
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
