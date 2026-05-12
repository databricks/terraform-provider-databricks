// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disaster_recovery_failover_group

import (
	"context"
	"reflect"

	"github.com/databricks/databricks-sdk-go/service/disasterrecovery"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/autogen"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/converters"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/tfschema"
	"github.com/databricks/terraform-provider-databricks/internal/service/disasterrecovery_tf"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const dataSourceName = "disaster_recovery_failover_group"

var _ datasource.DataSourceWithConfigure = &FailoverGroupDataSource{}

func DataSourceFailoverGroup() datasource.DataSource {
	return &FailoverGroupDataSource{}
}

type FailoverGroupDataSource struct {
	Client *autogen.DatabricksClient
}

// FailoverGroupData extends the main model with additional fields.
type FailoverGroupData struct {
	// Time at which this failover group was created.
	CreateTime timetypes.RFC3339 `tfsdk:"create_time"`
	// Current effective primary region. Replication flows FROM workspaces in
	// this region. Changes after a successful failover.
	EffectivePrimaryRegion types.String `tfsdk:"effective_primary_region"`
	// Opaque version string for optimistic locking. Server-generated, returned
	// in responses. Must be provided on Update requests to prevent concurrent
	// modifications.
	Etag types.String `tfsdk:"etag"`
	// Initial primary region. Used only in Create requests to set the starting
	// primary region. Not returned in responses.
	InitialPrimaryRegion types.String `tfsdk:"initial_primary_region"`
	// Fully qualified resource name in the format
	// accounts/{account_id}/failover-groups/{failover_group_id}.
	Name types.String `tfsdk:"name"`
	// List of all regions participating in this failover group.
	Regions types.List `tfsdk:"regions"`
	// The latest point in time to which data has been replicated.
	ReplicationPoint timetypes.RFC3339 `tfsdk:"replication_point"`
	// Aggregate state of the failover group.
	State types.String `tfsdk:"state"`
	// Unity Catalog replication configuration.
	UnityCatalogAssets types.Object `tfsdk:"unity_catalog_assets"`
	// Time at which this failover group was last modified.
	UpdateTime timetypes.RFC3339 `tfsdk:"update_time"`
	// Workspace sets, each containing workspaces that replicate to each other.
	WorkspaceSets types.List `tfsdk:"workspace_sets"`
}

// GetComplexFieldTypes returns a map of the types of elements in complex fields in the extended
// FailoverGroupData struct. Container types (types.Map, types.List, types.Set) and
// object types (types.Object) do not carry the type information of their elements in the Go
// type system. This function provides a way to retrieve the type information of the elements in
// complex fields at runtime. The values of the map are the reflected types of the contained elements.
// They must be either primitive values from the plugin framework type system
// (types.String{}, types.Bool{}, types.Int64{}, types.Float64{}) or TF SDK values.
func (m FailoverGroupData) GetComplexFieldTypes(ctx context.Context) map[string]reflect.Type {
	return map[string]reflect.Type{
		"regions":              reflect.TypeOf(types.String{}),
		"unity_catalog_assets": reflect.TypeOf(disasterrecovery_tf.UcReplicationConfig{}),
		"workspace_sets":       reflect.TypeOf(disasterrecovery_tf.WorkspaceSet{}),
	}
}

// ToObjectValue returns the object value for the resource, combining attributes from the
// embedded TFSDK model and contains additional fields.
//
// TFSDK types cannot implement the ObjectValuable interface directly, as it would otherwise
// interfere with how the plugin framework retrieves and sets values in state. Thus, FailoverGroupData
// only implements ToObjectValue() and Type().
func (m FailoverGroupData) ToObjectValue(ctx context.Context) basetypes.ObjectValue {
	return types.ObjectValueMust(
		m.Type(ctx).(basetypes.ObjectType).AttrTypes,
		map[string]attr.Value{
			"create_time":              m.CreateTime,
			"effective_primary_region": m.EffectivePrimaryRegion,
			"etag":                     m.Etag,
			"initial_primary_region":   m.InitialPrimaryRegion,
			"name":                     m.Name,
			"regions":                  m.Regions,
			"replication_point":        m.ReplicationPoint,
			"state":                    m.State,
			"unity_catalog_assets":     m.UnityCatalogAssets,
			"update_time":              m.UpdateTime,
			"workspace_sets":           m.WorkspaceSets,
		},
	)
}

// Type returns the object type with attributes from both the embedded TFSDK model
// and contains additional fields.
func (m FailoverGroupData) Type(ctx context.Context) attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"create_time":              timetypes.RFC3339{}.Type(ctx),
			"effective_primary_region": types.StringType,
			"etag":                     types.StringType,
			"initial_primary_region":   types.StringType,
			"name":                     types.StringType,
			"regions": basetypes.ListType{
				ElemType: types.StringType,
			},
			"replication_point":    timetypes.RFC3339{}.Type(ctx),
			"state":                types.StringType,
			"unity_catalog_assets": disasterrecovery_tf.UcReplicationConfig{}.Type(ctx),
			"update_time":          timetypes.RFC3339{}.Type(ctx),
			"workspace_sets": basetypes.ListType{
				ElemType: disasterrecovery_tf.WorkspaceSet{}.Type(ctx),
			},
		},
	}
}

func (m FailoverGroupData) ApplySchemaCustomizations(attrs map[string]tfschema.AttributeBuilder) map[string]tfschema.AttributeBuilder {
	attrs["create_time"] = attrs["create_time"].SetComputed()
	attrs["effective_primary_region"] = attrs["effective_primary_region"].SetComputed()
	attrs["etag"] = attrs["etag"].SetComputed()
	attrs["initial_primary_region"] = attrs["initial_primary_region"].SetComputed()
	attrs["name"] = attrs["name"].SetRequired()
	attrs["regions"] = attrs["regions"].SetComputed()
	attrs["replication_point"] = attrs["replication_point"].SetComputed()
	attrs["state"] = attrs["state"].SetComputed()
	attrs["unity_catalog_assets"] = attrs["unity_catalog_assets"].SetComputed()
	attrs["update_time"] = attrs["update_time"].SetComputed()
	attrs["workspace_sets"] = attrs["workspace_sets"].SetComputed()

	return attrs
}

func (r *FailoverGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = autogen.GetDatabricksProductionName(dataSourceName)
}

func (r *FailoverGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs, blocks := tfschema.DataSourceStructToSchemaMap(ctx, FailoverGroupData{}, nil)
	resp.Schema = schema.Schema{
		Description: "Terraform schema for Databricks FailoverGroup",
		Attributes:  attrs,
		Blocks:      blocks,
	}
}

func (r *FailoverGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.Client = autogen.ConfigureDataSource(req, resp)
}

func (r *FailoverGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	ctx = pluginfwcontext.SetUserAgentInDataSourceContext(ctx, dataSourceName)

	var config FailoverGroupData
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var readRequest disasterrecovery.GetFailoverGroupRequest
	resp.Diagnostics.Append(converters.TfSdkToGoSdkStruct(ctx, config, &readRequest)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, clientDiags := r.Client.GetAccountClient()

	resp.Diagnostics.Append(clientDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response, err := client.DisasterRecovery.GetFailoverGroup(ctx, readRequest)
	if err != nil {
		resp.Diagnostics.AddError("failed to get disaster_recovery_failover_group", err.Error())
		return
	}

	var newState FailoverGroupData
	resp.Diagnostics.Append(converters.GoSdkToTfSdkStruct(ctx, response, &newState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
