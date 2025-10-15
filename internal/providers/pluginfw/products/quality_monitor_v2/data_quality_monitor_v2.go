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
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type ProviderConfigData struct {
	WorkspaceID types.String `tfsdk:"workspace_id"`
}

func (r ProviderConfigData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["workspace_id"] = attrs["workspace_id"].SetRequired()
	return attrs
}

func (r ProviderConfigData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{}
}

func (r ProviderConfigData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		r.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"workspace_id": r.WorkspaceID,
		},
	)
}

func (r ProviderConfigData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"workspace_id": types.StringType,
		},
	}
}

// QualityMonitorData extends the main model with additional fields.
type QualityMonitorData struct {
	AnomalyDetectionConfig types.Object `tfsdk:"anomaly_detection_config"`
	// The uuid of the request object. For example, schema id.
	ObjectId types.String `tfsdk:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType         types.String `tfsdk:"object_type"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// QualityMonitorData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m QualityMonitorData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"anomaly_detection_config": reflect.TypeOf(qualitymonitorv2_tf.AnomalyDetectionConfig{}),
		"provider_config":          reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, QualityMonitorData
// only implements ToObjectValue() and Type().
func (m QualityMonitorData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"anomaly_detection_config": m.AnomalyDetectionConfig,
			"object_id":   m.ObjectId,
			"object_type": m.ObjectType,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m QualityMonitorData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"anomaly_detection_config": qualitymonitorv2_tf.AnomalyDetectionConfig{}.Type(ctx),
			"object_id":   types.StringType,
			"object_type": types.StringType,

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *QualityMonitorData) SyncFieldsDuringRead(ctx context.Context, from QualityMonitorData) {
	to.ProviderConfigData = from.ProviderConfigData

}

func (m QualityMonitorData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["anomaly_detection_config"] = attrs["anomaly_detection_config"].SetComputed()
	attrs["object_id"] = attrs["object_id"].SetRequired()
	attrs["object_type"] = attrs["object_type"].SetRequired()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *QualityMonitorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *QualityMonitorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, QualityMonitorData{}, nil)
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

	var config QualityMonitorData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest qualitymonitorv2.GetQualityMonitorRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var namespace ProviderConfigData
	resp.Diagnostics.Append(config.ProviderConfigData.As(ctx, &namespace, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProvider(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
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

	var newState QualityMonitorData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
