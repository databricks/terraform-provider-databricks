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

// AccessRequestDestinationsData extends the main model with additional fields.
type AccessRequestDestinationsData struct {
	// Indicates whether any destinations are hidden from the caller due to a
	// lack of permissions. This value is true if the caller does not have
	// permission to see all destinations.
	AreAnyDestinationsHidden types.Bool `tfsdk:"are_any_destinations_hidden"`
	// The source securable from which the destinations are inherited. Either
	// the same value as securable (if destination is set directly on the
	// securable) or the nearest parent securable with destinations set.
	DestinationSourceSecurable types.Object `tfsdk:"destination_source_securable"`
	// The access request destinations for the securable.
	Destinations types.List `tfsdk:"destinations"`
	// The full name of the securable. Redundant with the name in the securable
	// object, but necessary for Terraform integration
	FullName types.String `tfsdk:"full_name"`
	// The securable for which the access request destinations are being
	// modified or read.
	Securable types.Object `tfsdk:"securable"`
	// The type of the securable. Redundant with the type in the securable
	// object, but necessary for Terraform integration
	SecurableType types.String `tfsdk:"securable_type"`
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
		"destination_source_securable": reflect.TypeOf(catalog_tf.Securable{}),
		"destinations":                 reflect.TypeOf(catalog_tf.NotificationDestination{}),
		"securable":                    reflect.TypeOf(catalog_tf.Securable{}),
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
		map[string]attr.Value{
			"are_any_destinations_hidden":  m.AreAnyDestinationsHidden,
			"destination_source_securable": m.DestinationSourceSecurable,
			"destinations":                 m.Destinations,
			"full_name":                    m.FullName,
			"securable":                    m.Securable,
			"securable_type":               m.SecurableType,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m AccessRequestDestinationsData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"are_any_destinations_hidden":  types.BoolType,
			"destination_source_securable": catalog_tf.Securable{}.Type(ctx),
			"destinations": basetypes.ListType{
				ElemType: catalog_tf.NotificationDestination{}.Type(ctx),
			},
			"full_name":      types.StringType,
			"securable":      catalog_tf.Securable{}.Type(ctx),
			"securable_type": types.StringType,
		},
	}
}

func (m AccessRequestDestinationsData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["are_any_destinations_hidden"] = attrs["are_any_destinations_hidden"].SetComputed()
	attrs["destination_source_securable"] = attrs["destination_source_securable"].SetComputed()
	attrs["destinations"] = attrs["destinations"].SetComputed()
	attrs["full_name"] = attrs["full_name"].SetRequired()
	attrs["securable"] = attrs["securable"].SetComputed()
	attrs["securable_type"] = attrs["securable_type"].SetRequired()

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

	client, clientDiags := r.Client.GetWorkspaceClient()

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

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
