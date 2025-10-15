// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package rfa_access_request_destinations

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/catalog_tf"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "rfa_access_request_destinations"

var _ datasource.DataSourceWithConfigure = &AccessRequestDestinationDataSource{}

func DataSourceAccessRequestDestination() datasource.DataSource {
	return &AccessRequestDestinationDataSource{}
}

type AccessRequestDestinationDataSource struct {
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

// AccessRequestDestinationsData extends the main model with additional fields.
type AccessRequestDestinationsData struct {
	// Indicates whether any destinations are hidden from the caller due to a
	// lack of permissions. This value is true if the caller does not have
	// permission to see all destinations.
	AreAnyDestinationsHidden types.Bool `tfsdk:"are_any_destinations_hidden"`
	// The access request destinations for the securable.
	Destinations types.List `tfsdk:"destinations"`
	// The securable for which the access request destinations are being
	// retrieved.
	Securable          types.Object `tfsdk:"securable"`
	ProviderConfigData types.Object `tfsdk:"provider_config"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// AccessRequestDestinationsData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m AccessRequestDestinationsData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"destinations":    reflect.TypeOf(catalog_tf.NotificationDestination{}),
		"securable":       reflect.TypeOf(catalog_tf.Securable{}),
		"provider_config": reflect.TypeOf(ProviderConfigData{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, AccessRequestDestinationsData
// only implements ToObjectValue() and Type().
func (m AccessRequestDestinationsData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{"are_any_destinations_hidden": m.AreAnyDestinationsHidden,
			"destinations": m.Destinations,
			"securable":    m.Securable,

			"provider_config": m.ProviderConfigData,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccessRequestDestinationsData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{"are_any_destinations_hidden": types.BoolType,
			"destinations": basetypes.ListType{
				ElemType: catalog_tf.NotificationDestination{}.Type(ctx),
			},
			"securable": catalog_tf.Securable{}.Type(ctx),

			"provider_config": ProviderConfigData{}.Type(ctx),
		},
	}
}

// SyncFieldsDuringRead copies values from the existing state into the receiver,
// including both embedded model fields and additional fields. This method is called
// during read.
func (to *AccessRequestDestinationsData) SyncFieldsDuringRead(ctx context.Context, from AccessRequestDestinationsData) {
	to.ProviderConfigData = from.ProviderConfigData

}

func (m AccessRequestDestinationsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["are_any_destinations_hidden"] = attrs["are_any_destinations_hidden"].SetComputed()
	attrs["destinations"] = attrs["destinations"].SetRequired()
	attrs["securable"] = attrs["securable"].SetRequired()

	attrs["provider_config"] = attrs["provider_config"].SetOptional()

	return attrs
}

func (r *AccessRequestDestinationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *AccessRequestDestinationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, AccessRequestDestinationsData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks AccessRequestDestinations",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *AccessRequestDestinationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *AccessRequestDestinationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config AccessRequestDestinationsData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest catalog.GetAccessRequestDestinationsRequest
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
	client, clientDiags := r.Client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, namespace.WorkspaceID.ValueString())

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.Rfa.GetAccessRequestDestinations(ctx, readRequest)
	if err != nil {
		if apierr.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("failed to get rfa_access_request_destinations", err.Error())
		return
	}

	var newState AccessRequestDestinationsData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newState.SyncFieldsDuringRead(ctx, config)

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
