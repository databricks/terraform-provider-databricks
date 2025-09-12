// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package quality_monitor_v2

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/qualitymonitorv2_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "quality_monitor_v2"

var _ datasource.DataSourceWithConfigure = &QualityMonitorDataSource{}

func DataSourceQualityMonitor() datasource.DataSource {
	return &QualityMonitorDataSource{}
}

type QualityMonitorDataSource struct {
	Client *autogen.DatabricksClient
}

// QualityMonitorDataExtended extends the main model with additional fields.
type QualityMonitorDataExtended struct {
	qualitymonitorv2_tf.QualityMonitor
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// QualityMonitorDataExtended struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m QualityMonitorDataExtended) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return m.QualityMonitor.GetComplexFieldTypes(ctx)
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QualityMonitorDataExtended
// only implements ToObjectValue() and Type().
func (m QualityMonitorDataExtended) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return m.QualityMonitor.ToObjectValue(ctx)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m QualityMonitorDataExtended) Type(ctx context.Context) attr.Type {
	return m.QualityMonitor.Type(ctx)
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (m *QualityMonitorDataExtended) SyncFieldsDuringRead(ctx context.Context, existingState QualityMonitorDataExtended) {
	m.QualityMonitor.SyncFieldsDuringRead(ctx, existingState.QualityMonitor)
}

func (r *QualityMonitorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *QualityMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, QualityMonitorDataExtended{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks QualityMonitor",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *QualityMonitorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *QualityMonitorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	client, diags := r.Client.GetWorkspaceClient()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config QualityMonitorDataExtended
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest qualitymonitorv2.GetQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.QualityMonitorV2.GetQualityMonitor(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get quality_monitor_v2", err.Error())
		return
	}

	var newState QualityMonitorDataExtended
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
