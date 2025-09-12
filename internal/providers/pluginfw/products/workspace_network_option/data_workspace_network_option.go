// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace_network_option

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/settings_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "workspace_network_option"

var _ datasource.DataSourceWithConfigure = &WorkspaceNetworkOptionDataSource{}

func DataSourceWorkspaceNetworkOption() datasource.DataSource {
	return &WorkspaceNetworkOptionDataSource{}
}

type WorkspaceNetworkOptionDataSource struct {
	Client *autogen.DatabricksClient
}

// WorkspaceNetworkOptionDataExtended extends the main model with additional fields.
type WorkspaceNetworkOptionDataExtended struct {
	settings_tf.WorkspaceNetworkOption
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// WorkspaceNetworkOptionDataExtended struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m WorkspaceNetworkOptionDataExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.WorkspaceNetworkOption.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, WorkspaceNetworkOptionDataExtended
// only implements ToObjectValue() and Type().
func (m WorkspaceNetworkOptionDataExtended) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return m.WorkspaceNetworkOption.ToObjectValue(ctx)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m WorkspaceNetworkOptionDataExtended) Type(ctx context.Context) attr.Type {
	return m.WorkspaceNetworkOption.Type(ctx)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *WorkspaceNetworkOptionDataExtended) SyncFieldsDuringRead(ctx context.Context, existingState WorkspaceNetworkOptionDataExtended) {
	m.WorkspaceNetworkOption.SyncFieldsDuringRead(ctx, existingState.WorkspaceNetworkOption)
}

func (r *WorkspaceNetworkOptionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *WorkspaceNetworkOptionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, WorkspaceNetworkOptionDataExtended{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks WorkspaceNetworkOption",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *WorkspaceNetworkOptionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *WorkspaceNetworkOptionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetAccountClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config WorkspaceNetworkOptionDataExtended
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest settings.GetWorkspaceNetworkOptionRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.WorkspaceNetworkConfiguration.GetWorkspaceNetworkOptionRpc(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get workspace_network_option", err.Error())
		return
	}

	var newState WorkspaceNetworkOptionDataExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
